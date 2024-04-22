// Copyright 2024 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/northes/go-moonshot"
)

type MoonshotModelProvider struct {
	temperature float64
	subType     string
	secretKey   string
}

func NewMoonshotModelProvider(subType string, secretKey string, temperature float64) (*MoonshotModelProvider, error) {
	client := &MoonshotModelProvider{
		subType:     subType,
		secretKey:   secretKey,
		temperature: temperature,
	}
	return client, nil
}

func (p *MoonshotModelProvider) GetPricing() string {
	return `URL: 
	https://api.moonshot.cn

Model

| Model           | Unit Of Charge | Price  |
|-----------------|----------------|--------|
| moonshot-v1-8k  | 1M tokens      | 12yuan |
| moonshot-v1-32k | 1M tokens      | 24yuan |
| moonshot-v1-128k| 1M tokens      | 60yuan |
`
}

func (p *MoonshotModelProvider) calculatePrice(modelResult *ModelResult) error {
	price := 0.0
	switch p.subType {
	case "moonshot-v1-8k":
		price = getPrice(modelResult.TotalTokenCount, 0.012)
	case "moonshot-v1-32k":
		price = getPrice(modelResult.TotalTokenCount, 0.024)
	case "moonshot-v1-128k":
		price = getPrice(modelResult.TotalTokenCount, 0.06)
	default:
		return fmt.Errorf("calculatePrice() error: unknown model type: %s", p.subType)
	}
	modelResult.TotalPrice = price
	modelResult.Currency = "CNY"
	return nil
}

func (p *MoonshotModelProvider) QueryText(question string, writer io.Writer, history []*RawMessage, prompt string, knowledgeMessages []*RawMessage) (*ModelResult, error) {
	if p.secretKey == "" {
		return nil, errors.New("missing moonshot_key")
	}
	cli, err := moonshot.NewClientWithConfig(
		moonshot.NewConfig(
			moonshot.WithAPIKey(p.secretKey),
		),
	)
	if err != nil {
		return nil, err
	}
	// Chat completions
	resp, err := cli.Chat().Completions(context.Background(), &moonshot.ChatCompletionsRequest{
		Model: moonshot.ChatCompletionsModelID(p.subType),
		Messages: []*moonshot.ChatCompletionsMessage{
			{
				Role:    moonshot.RoleUser,
				Content: question,
			},
		},
		Temperature: p.temperature,
	})
	if err != nil {
		return nil, err
	}

	flusher, ok := writer.(http.Flusher)
	if !ok {
		return nil, fmt.Errorf("writer does not implement http.Flusher")
	}

	flushData := func(data string) error {
		if _, err = fmt.Fprintf(writer, "event: message\ndata: %s\n\n", data); err != nil {
			return err
		}
		flusher.Flush()
		return nil
	}

	err = flushData(resp.Choices[0].Message.Content)
	if err != nil {
		return nil, err
	}

	modelResult := &ModelResult{
		PromptTokenCount:   resp.Usage.PromptTokens,
		ResponseTokenCount: resp.Usage.CompletionTokens,
		TotalTokenCount:    resp.Usage.TotalTokens,
	}

	err = p.calculatePrice(modelResult)
	if err != nil {
		return nil, err
	}

	return modelResult, nil
}
