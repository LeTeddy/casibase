// Copyright 2023 The Casibase Authors. All Rights Reserved.
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

// Package routers
// @APIVersion 1.70.0
// @Title Casibase RESTful API
// @Description Swagger Docs of Casibase Backend API
// @Contact admin@casibase.org
// @SecurityDefinition AccessToken apiKey Authorization header
// @Schemes https,http
// @ExternalDocs Find out more about Casibase
// @ExternalDocsUrl https://casibase.org/
package routers

import (
	"github.com/beego/beego"
	"github.com/casibase/casibase/controllers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	initAPI()
}

func initAPI() {
	ns := beego.NewNamespace("/api",
		beego.NSInclude(
			&controllers.ApiController{},
		),
	)
	beego.AddNamespace(ns)

	beego.Router("/api/signin", &controllers.ApiController{}, "POST:Signin")
	beego.Router("/api/signout", &controllers.ApiController{}, "POST:Signout")
	beego.Router("/api/get-account", &controllers.ApiController{}, "GET:GetAccount")

	beego.Router("/api/get-global-videos", &controllers.ApiController{}, "GET:GetGlobalVideos")
	beego.Router("/api/get-videos", &controllers.ApiController{}, "GET:GetVideos")
	beego.Router("/api/get-video", &controllers.ApiController{}, "GET:GetVideo")
	beego.Router("/api/update-video", &controllers.ApiController{}, "POST:UpdateVideo")
	beego.Router("/api/add-video", &controllers.ApiController{}, "POST:AddVideo")
	beego.Router("/api/delete-video", &controllers.ApiController{}, "POST:DeleteVideo")
	beego.Router("/api/upload-video", &controllers.ApiController{}, "POST:UploadVideo")

	beego.Router("/api/get-global-stores", &controllers.ApiController{}, "GET:GetGlobalStores")
	beego.Router("/api/get-stores", &controllers.ApiController{}, "GET:GetStores")
	beego.Router("/api/get-store", &controllers.ApiController{}, "GET:GetStore")
	beego.Router("/api/update-store", &controllers.ApiController{}, "POST:UpdateStore")
	beego.Router("/api/add-store", &controllers.ApiController{}, "POST:AddStore")
	beego.Router("/api/delete-store", &controllers.ApiController{}, "POST:DeleteStore")
	beego.Router("/api/refresh-store-vectors", &controllers.ApiController{}, "POST:RefreshStoreVectors")
	beego.Router("/api/get-storage-providers", &controllers.ApiController{}, "GET:GetStorageProviders")

	beego.Router("/api/get-global-providers", &controllers.ApiController{}, "GET:GetGlobalProviders")
	beego.Router("/api/get-providers", &controllers.ApiController{}, "GET:GetProviders")
	beego.Router("/api/get-provider", &controllers.ApiController{}, "GET:GetProvider")
	beego.Router("/api/update-provider", &controllers.ApiController{}, "POST:UpdateProvider")
	beego.Router("/api/add-provider", &controllers.ApiController{}, "POST:AddProvider")
	beego.Router("/api/delete-provider", &controllers.ApiController{}, "POST:DeleteProvider")
	beego.Router("/api/refresh-mcp-tools", &controllers.ApiController{}, "POST:RefreshMcpTools")

	beego.Router("/api/get-global-vectors", &controllers.ApiController{}, "GET:GetGlobalVectors")
	beego.Router("/api/get-vectors", &controllers.ApiController{}, "GET:GetVectors")
	beego.Router("/api/get-vector", &controllers.ApiController{}, "GET:GetVector")
	beego.Router("/api/update-vector", &controllers.ApiController{}, "POST:UpdateVector")
	beego.Router("/api/add-vector", &controllers.ApiController{}, "POST:AddVector")
	beego.Router("/api/delete-vector", &controllers.ApiController{}, "POST:DeleteVector")
	beego.Router("/api/delete-all-vectors", &controllers.ApiController{}, "POST:DeleteAllVectors")

	beego.Router("/api/generate-text-to-speech-audio", &controllers.ApiController{}, "POST:GenerateTextToSpeechAudio")
	beego.Router("/api/generate-text-to-speech-audio-stream", &controllers.ApiController{}, "GET:GenerateTextToSpeechAudioStream")
	beego.Router("/api/process-speech-to-text", &controllers.ApiController{}, "POST:ProcessSpeechToText")

	beego.Router("/api/get-global-chats", &controllers.ApiController{}, "GET:GetGlobalChats")
	beego.Router("/api/get-chats", &controllers.ApiController{}, "GET:GetChats")
	beego.Router("/api/get-chat", &controllers.ApiController{}, "GET:GetChat")
	beego.Router("/api/update-chat", &controllers.ApiController{}, "POST:UpdateChat")
	beego.Router("/api/add-chat", &controllers.ApiController{}, "POST:AddChat")
	beego.Router("/api/delete-chat", &controllers.ApiController{}, "POST:DeleteChat")

	beego.Router("/api/get-global-messages", &controllers.ApiController{}, "GET:GetGlobalMessages")
	beego.Router("/api/get-messages", &controllers.ApiController{}, "GET:GetMessages")
	beego.Router("/api/get-message", &controllers.ApiController{}, "GET:GetMessage")
	beego.Router("/api/get-message-answer", &controllers.ApiController{}, "GET:GetMessageAnswer")
	beego.Router("/api/get-answer", &controllers.ApiController{}, "GET:GetAnswer")
	beego.Router("/api/update-message", &controllers.ApiController{}, "POST:UpdateMessage")
	beego.Router("/api/add-message", &controllers.ApiController{}, "POST:AddMessage")
	beego.Router("/api/delete-message", &controllers.ApiController{}, "POST:DeleteMessage")
	beego.Router("/api/delete-welcome-message", &controllers.ApiController{}, "POST:DeleteWelcomeMessage")

	beego.Router("/api/get-templates", &controllers.ApiController{}, "GET:GetTemplates")
	beego.Router("/api/get-template", &controllers.ApiController{}, "GET:GetTemplate")
	beego.Router("/api/update-template", &controllers.ApiController{}, "POST:UpdateTemplate")
	beego.Router("/api/add-template", &controllers.ApiController{}, "POST:AddTemplate")
	beego.Router("/api/delete-template", &controllers.ApiController{}, "POST:DeleteTemplate")
	beego.Router("/api/get-k8s-status", &controllers.ApiController{}, "GET:GetK8sStatus")

	beego.Router("/api/get-applications", &controllers.ApiController{}, "GET:GetApplications")
	beego.Router("/api/get-application", &controllers.ApiController{}, "GET:GetApplication")
	beego.Router("/api/update-application", &controllers.ApiController{}, "POST:UpdateApplication")
	beego.Router("/api/add-application", &controllers.ApiController{}, "POST:AddApplication")
	beego.Router("/api/delete-application", &controllers.ApiController{}, "POST:DeleteApplication")

	beego.Router("/api/deploy-application", &controllers.ApiController{}, "POST:DeployApplication")
	beego.Router("/api/undeploy-application", &controllers.ApiController{}, "POST:UndeployApplication")
	beego.Router("/api/get-application-status", &controllers.ApiController{}, "GET:GetApplicationStatus")

	beego.Router("/api/get-usages", &controllers.ApiController{}, "GET:GetUsages")
	beego.Router("/api/get-range-usages", &controllers.ApiController{}, "GET:GetRangeUsages")
	beego.Router("/api/get-users", &controllers.ApiController{}, "GET:GetUsers")
	beego.Router("/api/get-user-table-infos", &controllers.ApiController{}, "GET:GetUserTableInfos")

	beego.Router("/api/get-activities", &controllers.ApiController{}, "GET:GetActivities")
	// beego.Router("/api/get-range-activities", &controllers.ApiController{}, "GET:GetRangeActivities")

	beego.Router("/api/get-global-workflows", &controllers.ApiController{}, "GET:GetGlobalWorkflows")
	beego.Router("/api/get-workflows", &controllers.ApiController{}, "GET:GetWorkflows")
	beego.Router("/api/get-workflow", &controllers.ApiController{}, "GET:GetWorkflow")
	beego.Router("/api/update-workflow", &controllers.ApiController{}, "POST:UpdateWorkflow")
	beego.Router("/api/add-workflow", &controllers.ApiController{}, "POST:AddWorkflow")
	beego.Router("/api/delete-workflow", &controllers.ApiController{}, "POST:DeleteWorkflow")

	beego.Router("/api/get-global-tasks", &controllers.ApiController{}, "GET:GetGlobalTasks")
	beego.Router("/api/get-tasks", &controllers.ApiController{}, "GET:GetTasks")
	beego.Router("/api/get-task", &controllers.ApiController{}, "GET:GetTask")
	beego.Router("/api/update-task", &controllers.ApiController{}, "POST:UpdateTask")
	beego.Router("/api/add-task", &controllers.ApiController{}, "POST:AddTask")
	beego.Router("/api/delete-task", &controllers.ApiController{}, "POST:DeleteTask")

	beego.Router("/api/get-global-forms", &controllers.ApiController{}, "GET:GetGlobalForms")
	beego.Router("/api/get-forms", &controllers.ApiController{}, "GET:GetForms")
	beego.Router("/api/get-form", &controllers.ApiController{}, "GET:GetForm")
	beego.Router("/api/update-form", &controllers.ApiController{}, "POST:UpdateForm")
	beego.Router("/api/add-form", &controllers.ApiController{}, "POST:AddForm")
	beego.Router("/api/delete-form", &controllers.ApiController{}, "POST:DeleteForm")

	beego.Router("/api/get-form-data", &controllers.ApiController{}, "GET:GetFormData")

	beego.Router("/api/get-global-articles", &controllers.ApiController{}, "GET:GetGlobalArticles")
	beego.Router("/api/get-articles", &controllers.ApiController{}, "GET:GetArticles")
	beego.Router("/api/get-article", &controllers.ApiController{}, "GET:GetArticle")
	beego.Router("/api/update-article", &controllers.ApiController{}, "POST:UpdateArticle")
	beego.Router("/api/add-article", &controllers.ApiController{}, "POST:AddArticle")
	beego.Router("/api/delete-article", &controllers.ApiController{}, "POST:DeleteArticle")

	beego.Router("/api/update-file", &controllers.ApiController{}, "POST:UpdateFile")
	beego.Router("/api/add-file", &controllers.ApiController{}, "POST:AddFile")
	beego.Router("/api/delete-file", &controllers.ApiController{}, "POST:DeleteFile")
	beego.Router("/api/activate-file", &controllers.ApiController{}, "POST:ActivateFile")
	beego.Router("/api/get-active-file", &controllers.ApiController{}, "GET:GetActiveFile")

	beego.Router("/api/upload-file", &controllers.ApiController{}, "POST:UploadFile")

	beego.Router("/api/get-permissions", &controllers.ApiController{}, "GET:GetPermissions")
	beego.Router("/api/get-permission", &controllers.ApiController{}, "GET:GetPermission")
	beego.Router("/api/update-permission", &controllers.ApiController{}, "POST:UpdatePermission")
	beego.Router("/api/add-permission", &controllers.ApiController{}, "POST:AddPermission")
	beego.Router("/api/delete-permission", &controllers.ApiController{}, "POST:DeletePermission")

	beego.Router("/api/get-nodes", &controllers.ApiController{}, "GET:GetNodes")
	beego.Router("/api/get-node", &controllers.ApiController{}, "GET:GetNode")
	beego.Router("/api/update-node", &controllers.ApiController{}, "POST:UpdateNode")
	beego.Router("/api/add-node", &controllers.ApiController{}, "POST:AddNode")
	beego.Router("/api/delete-node", &controllers.ApiController{}, "POST:DeleteNode")

	beego.Router("/api/get-machines", &controllers.ApiController{}, "GET:GetMachines")
	beego.Router("/api/get-machine", &controllers.ApiController{}, "GET:GetMachine")
	beego.Router("/api/update-machine", &controllers.ApiController{}, "POST:UpdateMachine")
	beego.Router("/api/add-machine", &controllers.ApiController{}, "POST:AddMachine")
	beego.Router("/api/delete-machine", &controllers.ApiController{}, "POST:DeleteMachine")

	beego.Router("/api/get-images", &controllers.ApiController{}, "GET:GetImages")
	beego.Router("/api/get-image", &controllers.ApiController{}, "GET:GetImage")
	beego.Router("/api/update-image", &controllers.ApiController{}, "POST:UpdateImage")
	beego.Router("/api/add-image", &controllers.ApiController{}, "POST:AddImage")
	beego.Router("/api/delete-image", &controllers.ApiController{}, "POST:DeleteImage")

	beego.Router("/api/get-containers", &controllers.ApiController{}, "GET:GetContainers")
	beego.Router("/api/get-container", &controllers.ApiController{}, "GET:GetContainer")
	beego.Router("/api/update-container", &controllers.ApiController{}, "POST:UpdateContainer")
	beego.Router("/api/add-container", &controllers.ApiController{}, "POST:AddContainer")
	beego.Router("/api/delete-container", &controllers.ApiController{}, "POST:DeleteContainer")

	beego.Router("/api/get-pods", &controllers.ApiController{}, "GET:GetPods")
	beego.Router("/api/get-pod", &controllers.ApiController{}, "GET:GetPod")
	beego.Router("/api/update-pod", &controllers.ApiController{}, "POST:UpdatePod")
	beego.Router("/api/add-pod", &controllers.ApiController{}, "POST:AddPod")
	beego.Router("/api/delete-pod", &controllers.ApiController{}, "POST:DeletePod")

	beego.Router("/api/add-node-tunnel", &controllers.ApiController{}, "POST:AddNodeTunnel")
	beego.Router("/api/get-node-tunnel", &controllers.ApiController{}, "GET:GetNodeTunnel")

	beego.Router("/api/get-connections", &controllers.ApiController{}, "GET:GetConnections")
	beego.Router("/api/get-connection", &controllers.ApiController{}, "GET:GetConnection")
	beego.Router("/api/update-connection", &controllers.ApiController{}, "POST:UpdateConnection")
	beego.Router("/api/add-connection", &controllers.ApiController{}, "POST:AddConnection")
	beego.Router("/api/delete-connection", &controllers.ApiController{}, "POST:DeleteConnection")
	beego.Router("/api/start-connection", &controllers.ApiController{}, "POST:StartConnection")
	beego.Router("/api/stop-connection", &controllers.ApiController{}, "POST:StopConnection")

	beego.Router("/api/get-records", &controllers.ApiController{}, "GET:GetRecords")
	beego.Router("/api/get-record", &controllers.ApiController{}, "GET:GetRecord")
	beego.Router("/api/update-record", &controllers.ApiController{}, "POST:UpdateRecord")
	beego.Router("/api/add-record", &controllers.ApiController{}, "POST:AddRecord")
	beego.Router("/api/add-records", &controllers.ApiController{}, "POST:AddRecords")
	beego.Router("/api/delete-record", &controllers.ApiController{}, "POST:DeleteRecord")

	beego.Router("/api/commit-record", &controllers.ApiController{}, "POST:CommitRecord")
	beego.Router("/api/commit-record-second", &controllers.ApiController{}, "POST:CommitRecordSecond")
	beego.Router("/api/query-record", &controllers.ApiController{}, "GET:QueryRecord")
	beego.Router("/api/query-record-second", &controllers.ApiController{}, "GET:QueryRecordSecond")

	beego.Router("/api/get-system-info", &controllers.ApiController{}, "GET:GetSystemInfo")
	beego.Router("/api/get-version-info", &controllers.ApiController{}, "GET:GetVersionInfo")
	beego.Router("/api/health", &controllers.ApiController{}, "GET:Health")
	beego.Router("/api/get-prometheus-info", &controllers.ApiController{}, "GET:GetPrometheusInfo")
	beego.Handler("/api/metrics", promhttp.Handler())

	beego.Router("/api/chat/completions", &controllers.ApiController{}, "POST:ChatCompletions")
}
