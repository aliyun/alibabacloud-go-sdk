// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateWorkflowRequest struct {
	// example:
	//
	// ag-v7n2gp3vv3j****
	AlertGroupId *string `json:"alertGroupId,omitempty" xml:"alertGroupId,omitempty"`
	// example:
	//
	// NONE
	AlertStrategy *string `json:"alertStrategy,omitempty" xml:"alertStrategy,omitempty"`
	// example:
	//
	// 0 0 	- 	- 	- ? *
	CronExpr *string `json:"cronExpr,omitempty" xml:"cronExpr,omitempty"`
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// PARALLEL
	ExecutionType *string `json:"executionType,omitempty" xml:"executionType,omitempty"`
	// example:
	//
	// END
	FailureStrategy *string `json:"failureStrategy,omitempty" xml:"failureStrategy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// wd-v7n2gp3vv3j****
	ParentDirectoryId *string `json:"parentDirectoryId,omitempty" xml:"parentDirectoryId,omitempty"`
	// example:
	//
	// wg-acfmv4opbs****
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	ScheduleEndTime *string `json:"scheduleEndTime,omitempty" xml:"scheduleEndTime,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	ScheduleStartTime *string `json:"scheduleStartTime,omitempty" xml:"scheduleStartTime,omitempty"`
	// example:
	//
	// OFFLINE
	ScheduleState *string `json:"scheduleState,omitempty" xml:"scheduleState,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [{"taskId":"t1","name":"t1","taskParams":{"rawScript":"echo 1"},"taskType":"SHELL"}]
	TaskDefinitionJson *string `json:"taskDefinitionJson,omitempty" xml:"taskDefinitionJson,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [{"preTaskId":"0", "postTaskId":"t1"}]
	TaskRelationJson *string `json:"taskRelationJson,omitempty" xml:"taskRelationJson,omitempty"`
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
	// example:
	//
	// 10
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// example:
	//
	// MEDIUM
	WorkflowInstancePriority *string `json:"workflowInstancePriority,omitempty" xml:"workflowInstancePriority,omitempty"`
	// example:
	//
	// [{"prop":"key1","value":"value1"}]
	WorkflowParams *string `json:"workflowParams,omitempty" xml:"workflowParams,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreateWorkflowRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkflowRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkflowRequest) SetAlertGroupId(v string) *CreateWorkflowRequest {
	s.AlertGroupId = &v
	return s
}

func (s *CreateWorkflowRequest) SetAlertStrategy(v string) *CreateWorkflowRequest {
	s.AlertStrategy = &v
	return s
}

func (s *CreateWorkflowRequest) SetCronExpr(v string) *CreateWorkflowRequest {
	s.CronExpr = &v
	return s
}

func (s *CreateWorkflowRequest) SetDescription(v string) *CreateWorkflowRequest {
	s.Description = &v
	return s
}

func (s *CreateWorkflowRequest) SetExecutionType(v string) *CreateWorkflowRequest {
	s.ExecutionType = &v
	return s
}

func (s *CreateWorkflowRequest) SetFailureStrategy(v string) *CreateWorkflowRequest {
	s.FailureStrategy = &v
	return s
}

func (s *CreateWorkflowRequest) SetName(v string) *CreateWorkflowRequest {
	s.Name = &v
	return s
}

func (s *CreateWorkflowRequest) SetParentDirectoryId(v string) *CreateWorkflowRequest {
	s.ParentDirectoryId = &v
	return s
}

func (s *CreateWorkflowRequest) SetResourceGroupId(v string) *CreateWorkflowRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateWorkflowRequest) SetScheduleEndTime(v string) *CreateWorkflowRequest {
	s.ScheduleEndTime = &v
	return s
}

func (s *CreateWorkflowRequest) SetScheduleStartTime(v string) *CreateWorkflowRequest {
	s.ScheduleStartTime = &v
	return s
}

func (s *CreateWorkflowRequest) SetScheduleState(v string) *CreateWorkflowRequest {
	s.ScheduleState = &v
	return s
}

func (s *CreateWorkflowRequest) SetTaskDefinitionJson(v string) *CreateWorkflowRequest {
	s.TaskDefinitionJson = &v
	return s
}

func (s *CreateWorkflowRequest) SetTaskRelationJson(v string) *CreateWorkflowRequest {
	s.TaskRelationJson = &v
	return s
}

func (s *CreateWorkflowRequest) SetTimeZone(v string) *CreateWorkflowRequest {
	s.TimeZone = &v
	return s
}

func (s *CreateWorkflowRequest) SetTimeout(v int32) *CreateWorkflowRequest {
	s.Timeout = &v
	return s
}

func (s *CreateWorkflowRequest) SetWorkflowInstancePriority(v string) *CreateWorkflowRequest {
	s.WorkflowInstancePriority = &v
	return s
}

func (s *CreateWorkflowRequest) SetWorkflowParams(v string) *CreateWorkflowRequest {
	s.WorkflowParams = &v
	return s
}

func (s *CreateWorkflowRequest) SetWorkspaceId(v string) *CreateWorkflowRequest {
	s.WorkspaceId = &v
	return s
}

type CreateWorkflowResponseBody struct {
	Data *CreateWorkflowResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 9E3A7161-EB7B-172B-8D18-FFB06BA38***
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateWorkflowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkflowResponseBody) SetData(v *CreateWorkflowResponseBodyData) *CreateWorkflowResponseBody {
	s.Data = v
	return s
}

func (s *CreateWorkflowResponseBody) SetRequestId(v string) *CreateWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkflowResponseBody) SetSuccess(v bool) *CreateWorkflowResponseBody {
	s.Success = &v
	return s
}

type CreateWorkflowResponseBodyData struct {
	// example:
	//
	// w-acfmv4opbs****
	WorkflowId *string `json:"workflowId,omitempty" xml:"workflowId,omitempty"`
}

func (s CreateWorkflowResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkflowResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateWorkflowResponseBodyData) SetWorkflowId(v string) *CreateWorkflowResponseBodyData {
	s.WorkflowId = &v
	return s
}

type CreateWorkflowResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWorkflowResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkflowResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkflowResponse) SetHeaders(v map[string]*string) *CreateWorkflowResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkflowResponse) SetStatusCode(v int32) *CreateWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkflowResponse) SetBody(v *CreateWorkflowResponseBody) *CreateWorkflowResponse {
	s.Body = v
	return s
}

type DeleteWorkflowRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s DeleteWorkflowRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkflowRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowRequest) SetWorkspaceId(v string) *DeleteWorkflowRequest {
	s.WorkspaceId = &v
	return s
}

type DeleteWorkflowResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteWorkflowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowResponseBody) SetRequestId(v string) *DeleteWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWorkflowResponseBody) SetSuccess(v bool) *DeleteWorkflowResponseBody {
	s.Success = &v
	return s
}

type DeleteWorkflowResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWorkflowResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkflowResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowResponse) SetHeaders(v map[string]*string) *DeleteWorkflowResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkflowResponse) SetStatusCode(v int32) *DeleteWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWorkflowResponse) SetBody(v *DeleteWorkflowResponseBody) *DeleteWorkflowResponse {
	s.Body = v
	return s
}

type DescribeManualTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123***
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s DescribeManualTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeManualTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeManualTaskRequest) SetWorkspaceId(v string) *DescribeManualTaskRequest {
	s.WorkspaceId = &v
	return s
}

type DescribeManualTaskResponseBody struct {
	// example:
	//
	// 2024-03-27 00:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// mt-3q9jo749ne5****
	ManualTaskId *string `json:"ManualTaskId,omitempty" xml:"ManualTaskId,omitempty"`
	// example:
	//
	// test
	ManualTaskName *string `json:"ManualTaskName,omitempty" xml:"ManualTaskName,omitempty"`
	// example:
	//
	// mtd-oy98v7n43el****
	ParentDirectoryId *string `json:"ParentDirectoryId,omitempty" xml:"ParentDirectoryId,omitempty"`
	// example:
	//
	// p-3q9jo749ne5****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// r-oy98v7n43el****
	ResourceIds *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// example:
	//
	// {
	//
	//     "yarnUser": "",
	//
	//     "conditionResult": "null",
	//
	//     "rawScript": "sleep 300",
	//
	//     "submitOnYarnFlag": false,
	//
	//     "emrClusterId": "",
	//
	//     "yarnPriority": "",
	//
	//     "dependence": "null",
	//
	//     "yarnMemory": "",
	//
	//     "localParams": [],
	//
	//     "switchResult": "null",
	//
	//     "resourceList": [],
	//
	//     "yarnQueue": "",
	//
	//     "yarnVCores": "",
	//
	//     "associateManualTaskFlag": false
	//
	// }
	TaskParams *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty"`
	// example:
	//
	// SHELL
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// example:
	//
	// 2024-03-27 00:00:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// Id of the request
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DescribeManualTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeManualTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeManualTaskResponseBody) SetCreateTime(v string) *DescribeManualTaskResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeManualTaskResponseBody) SetDescription(v string) *DescribeManualTaskResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeManualTaskResponseBody) SetManualTaskId(v string) *DescribeManualTaskResponseBody {
	s.ManualTaskId = &v
	return s
}

func (s *DescribeManualTaskResponseBody) SetManualTaskName(v string) *DescribeManualTaskResponseBody {
	s.ManualTaskName = &v
	return s
}

func (s *DescribeManualTaskResponseBody) SetParentDirectoryId(v string) *DescribeManualTaskResponseBody {
	s.ParentDirectoryId = &v
	return s
}

func (s *DescribeManualTaskResponseBody) SetProjectId(v string) *DescribeManualTaskResponseBody {
	s.ProjectId = &v
	return s
}

func (s *DescribeManualTaskResponseBody) SetResourceIds(v string) *DescribeManualTaskResponseBody {
	s.ResourceIds = &v
	return s
}

func (s *DescribeManualTaskResponseBody) SetTaskParams(v string) *DescribeManualTaskResponseBody {
	s.TaskParams = &v
	return s
}

func (s *DescribeManualTaskResponseBody) SetTaskType(v string) *DescribeManualTaskResponseBody {
	s.TaskType = &v
	return s
}

func (s *DescribeManualTaskResponseBody) SetUpdateTime(v string) *DescribeManualTaskResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeManualTaskResponseBody) SetRequestId(v string) *DescribeManualTaskResponseBody {
	s.RequestId = &v
	return s
}

type DescribeManualTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeManualTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeManualTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeManualTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeManualTaskResponse) SetHeaders(v map[string]*string) *DescribeManualTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeManualTaskResponse) SetStatusCode(v int32) *DescribeManualTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeManualTaskResponse) SetBody(v *DescribeManualTaskResponseBody) *DescribeManualTaskResponse {
	s.Body = v
	return s
}

type DescribeManualTaskInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123***
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s DescribeManualTaskInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeManualTaskInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeManualTaskInstanceRequest) SetWorkspaceId(v string) *DescribeManualTaskInstanceRequest {
	s.WorkspaceId = &v
	return s
}

type DescribeManualTaskInstanceResponseBody struct {
	// example:
	//
	// c-b933c5aac7f7***
	EmrClusterId *string `json:"EmrClusterId,omitempty" xml:"EmrClusterId,omitempty"`
	// example:
	//
	// 2024-03-27 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// application_123_***
	ExternalAppId *string `json:"ExternalAppId,omitempty" xml:"ExternalAppId,omitempty"`
	// example:
	//
	// mti-0k5vype05xm****
	ManualTaskInstanceId *string `json:"ManualTaskInstanceId,omitempty" xml:"ManualTaskInstanceId,omitempty"`
	// example:
	//
	// test
	ManualTaskInstanceName *string `json:"ManualTaskInstanceName,omitempty" xml:"ManualTaskInstanceName,omitempty"`
	// example:
	//
	// wg-123abc***
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 2024-03-27 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2024-03-27 00:00:00
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	// example:
	//
	// {
	//
	//     "yarnUser": "",
	//
	//     "conditionResult": "null",
	//
	//     "rawScript": "sleep 300",
	//
	//     "submitOnYarnFlag": false,
	//
	//     "emrClusterId": "",
	//
	//     "yarnPriority": "",
	//
	//     "dependence": "null",
	//
	//     "yarnMemory": "",
	//
	//     "localParams": [],
	//
	//     "switchResult": "null",
	//
	//     "resourceList": [],
	//
	//     "yarnQueue": "",
	//
	//     "yarnVCores": "",
	//
	//     "associateManualTaskFlag": false
	//
	// }
	TaskParams *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty"`
	// example:
	//
	// SHELL
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// Id of the request
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DescribeManualTaskInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeManualTaskInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeManualTaskInstanceResponseBody) SetEmrClusterId(v string) *DescribeManualTaskInstanceResponseBody {
	s.EmrClusterId = &v
	return s
}

func (s *DescribeManualTaskInstanceResponseBody) SetEndTime(v string) *DescribeManualTaskInstanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeManualTaskInstanceResponseBody) SetExternalAppId(v string) *DescribeManualTaskInstanceResponseBody {
	s.ExternalAppId = &v
	return s
}

func (s *DescribeManualTaskInstanceResponseBody) SetManualTaskInstanceId(v string) *DescribeManualTaskInstanceResponseBody {
	s.ManualTaskInstanceId = &v
	return s
}

func (s *DescribeManualTaskInstanceResponseBody) SetManualTaskInstanceName(v string) *DescribeManualTaskInstanceResponseBody {
	s.ManualTaskInstanceName = &v
	return s
}

func (s *DescribeManualTaskInstanceResponseBody) SetResourceGroupId(v string) *DescribeManualTaskInstanceResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeManualTaskInstanceResponseBody) SetStartTime(v string) *DescribeManualTaskInstanceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeManualTaskInstanceResponseBody) SetStatus(v string) *DescribeManualTaskInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeManualTaskInstanceResponseBody) SetSubmitTime(v string) *DescribeManualTaskInstanceResponseBody {
	s.SubmitTime = &v
	return s
}

func (s *DescribeManualTaskInstanceResponseBody) SetTaskParams(v string) *DescribeManualTaskInstanceResponseBody {
	s.TaskParams = &v
	return s
}

func (s *DescribeManualTaskInstanceResponseBody) SetTaskType(v string) *DescribeManualTaskInstanceResponseBody {
	s.TaskType = &v
	return s
}

func (s *DescribeManualTaskInstanceResponseBody) SetRequestId(v string) *DescribeManualTaskInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeManualTaskInstanceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeManualTaskInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeManualTaskInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeManualTaskInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeManualTaskInstanceResponse) SetHeaders(v map[string]*string) *DescribeManualTaskInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeManualTaskInstanceResponse) SetStatusCode(v int32) *DescribeManualTaskInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeManualTaskInstanceResponse) SetBody(v *DescribeManualTaskInstanceResponseBody) *DescribeManualTaskInstanceResponse {
	s.Body = v
	return s
}

type DescribeProjectRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s DescribeProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectRequest) GoString() string {
	return s.String()
}

func (s *DescribeProjectRequest) SetWorkspaceId(v string) *DescribeProjectRequest {
	s.WorkspaceId = &v
	return s
}

type DescribeProjectResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// this is a project description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_project_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// p-3q9jo749ne5****
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// B897B94B-6754-5D09-AB8C-2E8186CCADC0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DescribeProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProjectResponseBody) SetDescription(v string) *DescribeProjectResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeProjectResponseBody) SetName(v string) *DescribeProjectResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeProjectResponseBody) SetProjectId(v string) *DescribeProjectResponseBody {
	s.ProjectId = &v
	return s
}

func (s *DescribeProjectResponseBody) SetRequestId(v string) *DescribeProjectResponseBody {
	s.RequestId = &v
	return s
}

type DescribeProjectResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectResponse) GoString() string {
	return s.String()
}

func (s *DescribeProjectResponse) SetHeaders(v map[string]*string) *DescribeProjectResponse {
	s.Headers = v
	return s
}

func (s *DescribeProjectResponse) SetStatusCode(v int32) *DescribeProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProjectResponse) SetBody(v *DescribeProjectResponseBody) *DescribeProjectResponse {
	s.Body = v
	return s
}

type DescribeTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123***
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s DescribeTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeTaskRequest) SetWorkspaceId(v string) *DescribeTaskRequest {
	s.WorkspaceId = &v
	return s
}

type DescribeTaskResponseBody struct {
	// example:
	//
	// 2024-03-27 00:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 0
	DelayTime *int32 `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1
	FailRetryInterval *int32 `json:"FailRetryInterval,omitempty" xml:"FailRetryInterval,omitempty"`
	// example:
	//
	// 0
	FailRetryTimes *int32 `json:"FailRetryTimes,omitempty" xml:"FailRetryTimes,omitempty"`
	// example:
	//
	// YES
	Flag *string `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// example:
	//
	// p-3q9jo749ne5****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// r-oy98v7n43el****
	ResourceIds *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// example:
	//
	// t-3q9jo749ne5****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// test
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// {
	//
	//     "yarnUser": "",
	//
	//     "conditionResult": "null",
	//
	//     "rawScript": "sleep 300",
	//
	//     "submitOnYarnFlag": false,
	//
	//     "emrClusterId": "",
	//
	//     "yarnPriority": "",
	//
	//     "dependence": "null",
	//
	//     "yarnMemory": "",
	//
	//     "localParams": [],
	//
	//     "switchResult": "null",
	//
	//     "resourceList": [],
	//
	//     "yarnQueue": "",
	//
	//     "yarnVCores": "",
	//
	//     "associateManualTaskFlag": false
	//
	// }
	TaskParams *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty"`
	// example:
	//
	// MEDIUM
	TaskPriority *string `json:"TaskPriority,omitempty" xml:"TaskPriority,omitempty"`
	// example:
	//
	// SHELL
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// example:
	//
	// 10
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// example:
	//
	// CLOSE
	TimeoutFlag *string `json:"TimeoutFlag,omitempty" xml:"TimeoutFlag,omitempty"`
	// example:
	//
	// WARN
	TimeoutNotifyStrategy *string `json:"TimeoutNotifyStrategy,omitempty" xml:"TimeoutNotifyStrategy,omitempty"`
	// example:
	//
	// 2024-03-27 00:00:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// Id of the request
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DescribeTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTaskResponseBody) SetCreateTime(v string) *DescribeTaskResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeTaskResponseBody) SetDelayTime(v int32) *DescribeTaskResponseBody {
	s.DelayTime = &v
	return s
}

func (s *DescribeTaskResponseBody) SetDescription(v string) *DescribeTaskResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeTaskResponseBody) SetFailRetryInterval(v int32) *DescribeTaskResponseBody {
	s.FailRetryInterval = &v
	return s
}

func (s *DescribeTaskResponseBody) SetFailRetryTimes(v int32) *DescribeTaskResponseBody {
	s.FailRetryTimes = &v
	return s
}

func (s *DescribeTaskResponseBody) SetFlag(v string) *DescribeTaskResponseBody {
	s.Flag = &v
	return s
}

func (s *DescribeTaskResponseBody) SetProjectId(v string) *DescribeTaskResponseBody {
	s.ProjectId = &v
	return s
}

func (s *DescribeTaskResponseBody) SetResourceIds(v string) *DescribeTaskResponseBody {
	s.ResourceIds = &v
	return s
}

func (s *DescribeTaskResponseBody) SetTaskId(v string) *DescribeTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeTaskResponseBody) SetTaskName(v string) *DescribeTaskResponseBody {
	s.TaskName = &v
	return s
}

func (s *DescribeTaskResponseBody) SetTaskParams(v string) *DescribeTaskResponseBody {
	s.TaskParams = &v
	return s
}

func (s *DescribeTaskResponseBody) SetTaskPriority(v string) *DescribeTaskResponseBody {
	s.TaskPriority = &v
	return s
}

func (s *DescribeTaskResponseBody) SetTaskType(v string) *DescribeTaskResponseBody {
	s.TaskType = &v
	return s
}

func (s *DescribeTaskResponseBody) SetTimeout(v int32) *DescribeTaskResponseBody {
	s.Timeout = &v
	return s
}

func (s *DescribeTaskResponseBody) SetTimeoutFlag(v string) *DescribeTaskResponseBody {
	s.TimeoutFlag = &v
	return s
}

func (s *DescribeTaskResponseBody) SetTimeoutNotifyStrategy(v string) *DescribeTaskResponseBody {
	s.TimeoutNotifyStrategy = &v
	return s
}

func (s *DescribeTaskResponseBody) SetUpdateTime(v string) *DescribeTaskResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeTaskResponseBody) SetVersion(v string) *DescribeTaskResponseBody {
	s.Version = &v
	return s
}

func (s *DescribeTaskResponseBody) SetRequestId(v string) *DescribeTaskResponseBody {
	s.RequestId = &v
	return s
}

type DescribeTaskResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeTaskResponse) SetHeaders(v map[string]*string) *DescribeTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeTaskResponse) SetStatusCode(v int32) *DescribeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTaskResponse) SetBody(v *DescribeTaskResponseBody) *DescribeTaskResponse {
	s.Body = v
	return s
}

type DescribeTaskInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123***
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s DescribeTaskInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeTaskInstanceRequest) SetWorkspaceId(v string) *DescribeTaskInstanceRequest {
	s.WorkspaceId = &v
	return s
}

type DescribeTaskInstanceResponseBody struct {
	// example:
	//
	// 0
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// c-b933c5aac7f7***
	EmrClusterId *string `json:"EmrClusterId,omitempty" xml:"EmrClusterId,omitempty"`
	// example:
	//
	// 2024-03-27 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// application_123_***
	ExternalAppId *string `json:"ExternalAppId,omitempty" xml:"ExternalAppId,omitempty"`
	// example:
	//
	// wg-123abc***
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 0
	RetryTimes *int32 `json:"RetryTimes,omitempty" xml:"RetryTimes,omitempty"`
	// example:
	//
	// 2024-03-27 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2024-03-27 00:00:00
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	// example:
	//
	// t-3q9jo749ne5****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// ti-3q9jo749ne5****
	TaskInstanceId *string `json:"TaskInstanceId,omitempty" xml:"TaskInstanceId,omitempty"`
	// example:
	//
	// test
	TaskInstanceName *string `json:"TaskInstanceName,omitempty" xml:"TaskInstanceName,omitempty"`
	// example:
	//
	// {
	//
	//     "yarnUser": "",
	//
	//     "conditionResult": "null",
	//
	//     "rawScript": "sleep 300",
	//
	//     "submitOnYarnFlag": false,
	//
	//     "emrClusterId": "",
	//
	//     "yarnPriority": "",
	//
	//     "dependence": "null",
	//
	//     "yarnMemory": "",
	//
	//     "localParams": [],
	//
	//     "switchResult": "null",
	//
	//     "resourceList": [],
	//
	//     "yarnQueue": "",
	//
	//     "yarnVCores": "",
	//
	//     "associateManualTaskFlag": false
	//
	// }
	TaskParams *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty"`
	// example:
	//
	// SHELL
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// example:
	//
	// 1
	TaskVersion *string `json:"TaskVersion,omitempty" xml:"TaskVersion,omitempty"`
	// example:
	//
	// wi-3q9jo749ne5****
	WorkflowInstanceId *string `json:"WorkflowInstanceId,omitempty" xml:"WorkflowInstanceId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DescribeTaskInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTaskInstanceResponseBody) SetDryRun(v string) *DescribeTaskInstanceResponseBody {
	s.DryRun = &v
	return s
}

func (s *DescribeTaskInstanceResponseBody) SetEmrClusterId(v string) *DescribeTaskInstanceResponseBody {
	s.EmrClusterId = &v
	return s
}

func (s *DescribeTaskInstanceResponseBody) SetEndTime(v string) *DescribeTaskInstanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeTaskInstanceResponseBody) SetExternalAppId(v string) *DescribeTaskInstanceResponseBody {
	s.ExternalAppId = &v
	return s
}

func (s *DescribeTaskInstanceResponseBody) SetResourceGroupId(v string) *DescribeTaskInstanceResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeTaskInstanceResponseBody) SetRetryTimes(v int32) *DescribeTaskInstanceResponseBody {
	s.RetryTimes = &v
	return s
}

func (s *DescribeTaskInstanceResponseBody) SetStartTime(v string) *DescribeTaskInstanceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeTaskInstanceResponseBody) SetStatus(v string) *DescribeTaskInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeTaskInstanceResponseBody) SetSubmitTime(v string) *DescribeTaskInstanceResponseBody {
	s.SubmitTime = &v
	return s
}

func (s *DescribeTaskInstanceResponseBody) SetTaskId(v string) *DescribeTaskInstanceResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeTaskInstanceResponseBody) SetTaskInstanceId(v string) *DescribeTaskInstanceResponseBody {
	s.TaskInstanceId = &v
	return s
}

func (s *DescribeTaskInstanceResponseBody) SetTaskInstanceName(v string) *DescribeTaskInstanceResponseBody {
	s.TaskInstanceName = &v
	return s
}

func (s *DescribeTaskInstanceResponseBody) SetTaskParams(v string) *DescribeTaskInstanceResponseBody {
	s.TaskParams = &v
	return s
}

func (s *DescribeTaskInstanceResponseBody) SetTaskType(v string) *DescribeTaskInstanceResponseBody {
	s.TaskType = &v
	return s
}

func (s *DescribeTaskInstanceResponseBody) SetTaskVersion(v string) *DescribeTaskInstanceResponseBody {
	s.TaskVersion = &v
	return s
}

func (s *DescribeTaskInstanceResponseBody) SetWorkflowInstanceId(v string) *DescribeTaskInstanceResponseBody {
	s.WorkflowInstanceId = &v
	return s
}

func (s *DescribeTaskInstanceResponseBody) SetRequestId(v string) *DescribeTaskInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeTaskInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTaskInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTaskInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeTaskInstanceResponse) SetHeaders(v map[string]*string) *DescribeTaskInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeTaskInstanceResponse) SetStatusCode(v int32) *DescribeTaskInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTaskInstanceResponse) SetBody(v *DescribeTaskInstanceResponseBody) *DescribeTaskInstanceResponse {
	s.Body = v
	return s
}

type DescribeWorkflowRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s DescribeWorkflowRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWorkflowRequest) GoString() string {
	return s.String()
}

func (s *DescribeWorkflowRequest) SetWorkspaceId(v string) *DescribeWorkflowRequest {
	s.WorkspaceId = &v
	return s
}

type DescribeWorkflowResponseBody struct {
	// example:
	//
	// 611AD6E6-BFE3-5897-AA12-569F79DBAF9B
	RequestId     *string                                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Schedule      *DescribeWorkflowResponseBodySchedule        `json:"schedule,omitempty" xml:"schedule,omitempty" type:"Struct"`
	TaskRelations []*DescribeWorkflowResponseBodyTaskRelations `json:"taskRelations,omitempty" xml:"taskRelations,omitempty" type:"Repeated"`
	Tasks         []*DescribeWorkflowResponseBodyTasks         `json:"tasks,omitempty" xml:"tasks,omitempty" type:"Repeated"`
	Workflow      *DescribeWorkflowResponseBodyWorkflow        `json:"workflow,omitempty" xml:"workflow,omitempty" type:"Struct"`
}

func (s DescribeWorkflowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWorkflowResponseBody) SetRequestId(v string) *DescribeWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWorkflowResponseBody) SetSchedule(v *DescribeWorkflowResponseBodySchedule) *DescribeWorkflowResponseBody {
	s.Schedule = v
	return s
}

func (s *DescribeWorkflowResponseBody) SetTaskRelations(v []*DescribeWorkflowResponseBodyTaskRelations) *DescribeWorkflowResponseBody {
	s.TaskRelations = v
	return s
}

func (s *DescribeWorkflowResponseBody) SetTasks(v []*DescribeWorkflowResponseBodyTasks) *DescribeWorkflowResponseBody {
	s.Tasks = v
	return s
}

func (s *DescribeWorkflowResponseBody) SetWorkflow(v *DescribeWorkflowResponseBodyWorkflow) *DescribeWorkflowResponseBody {
	s.Workflow = v
	return s
}

type DescribeWorkflowResponseBodySchedule struct {
	// example:
	//
	// ag-n72kong0832****
	AlertGroupId *string `json:"alertGroupId,omitempty" xml:"alertGroupId,omitempty"`
	// example:
	//
	// NONE
	AlertStrategy *string `json:"alertStrategy,omitempty" xml:"alertStrategy,omitempty"`
	// example:
	//
	// 0 0 	- 	- 	- ? *
	CronExpr *string `json:"cronExpr,omitempty" xml:"cronExpr,omitempty"`
	// example:
	//
	// C-15F7AB9B53F1****
	EmrClusterId *string `json:"emrClusterId,omitempty" xml:"emrClusterId,omitempty"`
	// example:
	//
	// END
	FailureStrategy *string `json:"failureStrategy,omitempty" xml:"failureStrategy,omitempty"`
	// example:
	//
	// wg-susqimrr649x****
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	ScheduleEndTime *string `json:"scheduleEndTime,omitempty" xml:"scheduleEndTime,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	ScheduleStartTime *string `json:"scheduleStartTime,omitempty" xml:"scheduleStartTime,omitempty"`
	// example:
	//
	// OFFLINE
	ScheduleState *string `json:"scheduleState,omitempty" xml:"scheduleState,omitempty"`
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
	// example:
	//
	// MEDIUM
	WorkflowInstancePriority *string `json:"workflowInstancePriority,omitempty" xml:"workflowInstancePriority,omitempty"`
}

func (s DescribeWorkflowResponseBodySchedule) String() string {
	return tea.Prettify(s)
}

func (s DescribeWorkflowResponseBodySchedule) GoString() string {
	return s.String()
}

func (s *DescribeWorkflowResponseBodySchedule) SetAlertGroupId(v string) *DescribeWorkflowResponseBodySchedule {
	s.AlertGroupId = &v
	return s
}

func (s *DescribeWorkflowResponseBodySchedule) SetAlertStrategy(v string) *DescribeWorkflowResponseBodySchedule {
	s.AlertStrategy = &v
	return s
}

func (s *DescribeWorkflowResponseBodySchedule) SetCronExpr(v string) *DescribeWorkflowResponseBodySchedule {
	s.CronExpr = &v
	return s
}

func (s *DescribeWorkflowResponseBodySchedule) SetEmrClusterId(v string) *DescribeWorkflowResponseBodySchedule {
	s.EmrClusterId = &v
	return s
}

func (s *DescribeWorkflowResponseBodySchedule) SetFailureStrategy(v string) *DescribeWorkflowResponseBodySchedule {
	s.FailureStrategy = &v
	return s
}

func (s *DescribeWorkflowResponseBodySchedule) SetResourceGroupId(v string) *DescribeWorkflowResponseBodySchedule {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeWorkflowResponseBodySchedule) SetScheduleEndTime(v string) *DescribeWorkflowResponseBodySchedule {
	s.ScheduleEndTime = &v
	return s
}

func (s *DescribeWorkflowResponseBodySchedule) SetScheduleStartTime(v string) *DescribeWorkflowResponseBodySchedule {
	s.ScheduleStartTime = &v
	return s
}

func (s *DescribeWorkflowResponseBodySchedule) SetScheduleState(v string) *DescribeWorkflowResponseBodySchedule {
	s.ScheduleState = &v
	return s
}

func (s *DescribeWorkflowResponseBodySchedule) SetTimeZone(v string) *DescribeWorkflowResponseBodySchedule {
	s.TimeZone = &v
	return s
}

func (s *DescribeWorkflowResponseBodySchedule) SetWorkflowInstancePriority(v string) *DescribeWorkflowResponseBodySchedule {
	s.WorkflowInstancePriority = &v
	return s
}

type DescribeWorkflowResponseBodyTaskRelations struct {
	// example:
	//
	// t-n72kong0832****
	PostTaskId *string `json:"postTaskId,omitempty" xml:"postTaskId,omitempty"`
	// example:
	//
	// t-n72kong0832****
	PreTaskId *string `json:"preTaskId,omitempty" xml:"preTaskId,omitempty"`
}

func (s DescribeWorkflowResponseBodyTaskRelations) String() string {
	return tea.Prettify(s)
}

func (s DescribeWorkflowResponseBodyTaskRelations) GoString() string {
	return s.String()
}

func (s *DescribeWorkflowResponseBodyTaskRelations) SetPostTaskId(v string) *DescribeWorkflowResponseBodyTaskRelations {
	s.PostTaskId = &v
	return s
}

func (s *DescribeWorkflowResponseBodyTaskRelations) SetPreTaskId(v string) *DescribeWorkflowResponseBodyTaskRelations {
	s.PreTaskId = &v
	return s
}

type DescribeWorkflowResponseBodyTasks struct {
	// example:
	//
	// task description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// task_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// t-n72kong0832****
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 1
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s DescribeWorkflowResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeWorkflowResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DescribeWorkflowResponseBodyTasks) SetDescription(v string) *DescribeWorkflowResponseBodyTasks {
	s.Description = &v
	return s
}

func (s *DescribeWorkflowResponseBodyTasks) SetName(v string) *DescribeWorkflowResponseBodyTasks {
	s.Name = &v
	return s
}

func (s *DescribeWorkflowResponseBodyTasks) SetTaskId(v string) *DescribeWorkflowResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *DescribeWorkflowResponseBodyTasks) SetVersion(v int32) *DescribeWorkflowResponseBodyTasks {
	s.Version = &v
	return s
}

type DescribeWorkflowResponseBodyWorkflow struct {
	// example:
	//
	// 2024-01-01 00:00:00
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// PARALLEL
	ExecutionType *string `json:"executionType,omitempty" xml:"executionType,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// wd-n72kong0832****
	ParentDirectoryId *string `json:"parentDirectoryId,omitempty" xml:"parentDirectoryId,omitempty"`
	// example:
	//
	// 0
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// example:
	//
	// w-n72kong0832****
	WorkflowId *string `json:"workflowId,omitempty" xml:"workflowId,omitempty"`
	// example:
	//
	// [{"prop":"key1","value":"value1"}]
	WorkflowParams *string `json:"workflowParams,omitempty" xml:"workflowParams,omitempty"`
}

func (s DescribeWorkflowResponseBodyWorkflow) String() string {
	return tea.Prettify(s)
}

func (s DescribeWorkflowResponseBodyWorkflow) GoString() string {
	return s.String()
}

func (s *DescribeWorkflowResponseBodyWorkflow) SetCreateTime(v string) *DescribeWorkflowResponseBodyWorkflow {
	s.CreateTime = &v
	return s
}

func (s *DescribeWorkflowResponseBodyWorkflow) SetDescription(v string) *DescribeWorkflowResponseBodyWorkflow {
	s.Description = &v
	return s
}

func (s *DescribeWorkflowResponseBodyWorkflow) SetExecutionType(v string) *DescribeWorkflowResponseBodyWorkflow {
	s.ExecutionType = &v
	return s
}

func (s *DescribeWorkflowResponseBodyWorkflow) SetName(v string) *DescribeWorkflowResponseBodyWorkflow {
	s.Name = &v
	return s
}

func (s *DescribeWorkflowResponseBodyWorkflow) SetParentDirectoryId(v string) *DescribeWorkflowResponseBodyWorkflow {
	s.ParentDirectoryId = &v
	return s
}

func (s *DescribeWorkflowResponseBodyWorkflow) SetTimeout(v int32) *DescribeWorkflowResponseBodyWorkflow {
	s.Timeout = &v
	return s
}

func (s *DescribeWorkflowResponseBodyWorkflow) SetUpdateTime(v string) *DescribeWorkflowResponseBodyWorkflow {
	s.UpdateTime = &v
	return s
}

func (s *DescribeWorkflowResponseBodyWorkflow) SetWorkflowId(v string) *DescribeWorkflowResponseBodyWorkflow {
	s.WorkflowId = &v
	return s
}

func (s *DescribeWorkflowResponseBodyWorkflow) SetWorkflowParams(v string) *DescribeWorkflowResponseBodyWorkflow {
	s.WorkflowParams = &v
	return s
}

type DescribeWorkflowResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWorkflowResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWorkflowResponse) GoString() string {
	return s.String()
}

func (s *DescribeWorkflowResponse) SetHeaders(v map[string]*string) *DescribeWorkflowResponse {
	s.Headers = v
	return s
}

func (s *DescribeWorkflowResponse) SetStatusCode(v int32) *DescribeWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWorkflowResponse) SetBody(v *DescribeWorkflowResponseBody) *DescribeWorkflowResponse {
	s.Body = v
	return s
}

type DescribeWorkflowInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s DescribeWorkflowInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWorkflowInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeWorkflowInstanceRequest) SetWorkspaceId(v string) *DescribeWorkflowInstanceRequest {
	s.WorkspaceId = &v
	return s
}

type DescribeWorkflowInstanceResponseBody struct {
	// example:
	//
	// ag-n72kong0832****
	AlertGroupId *string `json:"alertGroupId,omitempty" xml:"alertGroupId,omitempty"`
	// example:
	//
	// NONE
	AlertStrategy *string `json:"alertStrategy,omitempty" xml:"alertStrategy,omitempty"`
	// example:
	//
	// c-047fa6bbe732****
	EmrClusterId *string `json:"emrClusterId,omitempty" xml:"emrClusterId,omitempty"`
	// example:
	//
	// 2024-03-27 00:00:00
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// END
	FailureStrategy *string `json:"failureStrategy,omitempty" xml:"failureStrategy,omitempty"`
	// example:
	//
	// false
	IsComplementData *bool `json:"isComplementData,omitempty" xml:"isComplementData,omitempty"`
	// example:
	//
	// workflow_instance_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// wg-susqimrr649x****
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// example:
	//
	// 2024-03-27 00:00:00
	RestartTime *string `json:"restartTime,omitempty" xml:"restartTime,omitempty"`
	// example:
	//
	// 1
	RunTimes *int32 `json:"runTimes,omitempty" xml:"runTimes,omitempty"`
	// example:
	//
	// 2024-03-27 00:00:00
	ScheduleTime *string `json:"scheduleTime,omitempty" xml:"scheduleTime,omitempty"`
	// example:
	//
	// 2024-03-27 00:00:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 0
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// example:
	//
	// w-3q9jo749ne5****
	WorkflowId *string `json:"workflowId,omitempty" xml:"workflowId,omitempty"`
	// example:
	//
	// wi-3q9jo749ne5****
	WorkflowInstanceId *string `json:"workflowInstanceId,omitempty" xml:"workflowInstanceId,omitempty"`
	// example:
	//
	// MEDIUM
	WorkflowInstancePriority *string `json:"workflowInstancePriority,omitempty" xml:"workflowInstancePriority,omitempty"`
	// example:
	//
	// 1
	WorkflowVersion *int32 `json:"workflowVersion,omitempty" xml:"workflowVersion,omitempty"`
}

func (s DescribeWorkflowInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWorkflowInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWorkflowInstanceResponseBody) SetAlertGroupId(v string) *DescribeWorkflowInstanceResponseBody {
	s.AlertGroupId = &v
	return s
}

func (s *DescribeWorkflowInstanceResponseBody) SetAlertStrategy(v string) *DescribeWorkflowInstanceResponseBody {
	s.AlertStrategy = &v
	return s
}

func (s *DescribeWorkflowInstanceResponseBody) SetEmrClusterId(v string) *DescribeWorkflowInstanceResponseBody {
	s.EmrClusterId = &v
	return s
}

func (s *DescribeWorkflowInstanceResponseBody) SetEndTime(v string) *DescribeWorkflowInstanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeWorkflowInstanceResponseBody) SetFailureStrategy(v string) *DescribeWorkflowInstanceResponseBody {
	s.FailureStrategy = &v
	return s
}

func (s *DescribeWorkflowInstanceResponseBody) SetIsComplementData(v bool) *DescribeWorkflowInstanceResponseBody {
	s.IsComplementData = &v
	return s
}

func (s *DescribeWorkflowInstanceResponseBody) SetName(v string) *DescribeWorkflowInstanceResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeWorkflowInstanceResponseBody) SetRequestId(v string) *DescribeWorkflowInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWorkflowInstanceResponseBody) SetResourceGroupId(v string) *DescribeWorkflowInstanceResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeWorkflowInstanceResponseBody) SetRestartTime(v string) *DescribeWorkflowInstanceResponseBody {
	s.RestartTime = &v
	return s
}

func (s *DescribeWorkflowInstanceResponseBody) SetRunTimes(v int32) *DescribeWorkflowInstanceResponseBody {
	s.RunTimes = &v
	return s
}

func (s *DescribeWorkflowInstanceResponseBody) SetScheduleTime(v string) *DescribeWorkflowInstanceResponseBody {
	s.ScheduleTime = &v
	return s
}

func (s *DescribeWorkflowInstanceResponseBody) SetStartTime(v string) *DescribeWorkflowInstanceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeWorkflowInstanceResponseBody) SetStatus(v string) *DescribeWorkflowInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeWorkflowInstanceResponseBody) SetTimeout(v int32) *DescribeWorkflowInstanceResponseBody {
	s.Timeout = &v
	return s
}

func (s *DescribeWorkflowInstanceResponseBody) SetWorkflowId(v string) *DescribeWorkflowInstanceResponseBody {
	s.WorkflowId = &v
	return s
}

func (s *DescribeWorkflowInstanceResponseBody) SetWorkflowInstanceId(v string) *DescribeWorkflowInstanceResponseBody {
	s.WorkflowInstanceId = &v
	return s
}

func (s *DescribeWorkflowInstanceResponseBody) SetWorkflowInstancePriority(v string) *DescribeWorkflowInstanceResponseBody {
	s.WorkflowInstancePriority = &v
	return s
}

func (s *DescribeWorkflowInstanceResponseBody) SetWorkflowVersion(v int32) *DescribeWorkflowInstanceResponseBody {
	s.WorkflowVersion = &v
	return s
}

type DescribeWorkflowInstanceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWorkflowInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWorkflowInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWorkflowInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeWorkflowInstanceResponse) SetHeaders(v map[string]*string) *DescribeWorkflowInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeWorkflowInstanceResponse) SetStatusCode(v int32) *DescribeWorkflowInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWorkflowInstanceResponse) SetBody(v *DescribeWorkflowInstanceResponseBody) *DescribeWorkflowInstanceResponse {
	s.Body = v
	return s
}

type ListManualTaskInstancesRequest struct {
	// example:
	//
	// 2024-03-27 00:00:00
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 123abc***
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// test
	SearchVal *string `json:"searchVal,omitempty" xml:"searchVal,omitempty"`
	// example:
	//
	// 2024-03-27 00:00:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123***
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListManualTaskInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListManualTaskInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListManualTaskInstancesRequest) SetEndTime(v string) *ListManualTaskInstancesRequest {
	s.EndTime = &v
	return s
}

func (s *ListManualTaskInstancesRequest) SetMaxResults(v int32) *ListManualTaskInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListManualTaskInstancesRequest) SetNextToken(v string) *ListManualTaskInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *ListManualTaskInstancesRequest) SetSearchVal(v string) *ListManualTaskInstancesRequest {
	s.SearchVal = &v
	return s
}

func (s *ListManualTaskInstancesRequest) SetStartTime(v string) *ListManualTaskInstancesRequest {
	s.StartTime = &v
	return s
}

func (s *ListManualTaskInstancesRequest) SetStatus(v string) *ListManualTaskInstancesRequest {
	s.Status = &v
	return s
}

func (s *ListManualTaskInstancesRequest) SetWorkspaceId(v string) *ListManualTaskInstancesRequest {
	s.WorkspaceId = &v
	return s
}

type ListManualTaskInstancesResponseBody struct {
	Data []*ListManualTaskInstancesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 123abc***
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 10
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListManualTaskInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListManualTaskInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListManualTaskInstancesResponseBody) SetData(v []*ListManualTaskInstancesResponseBodyData) *ListManualTaskInstancesResponseBody {
	s.Data = v
	return s
}

func (s *ListManualTaskInstancesResponseBody) SetNextToken(v string) *ListManualTaskInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListManualTaskInstancesResponseBody) SetRequestId(v string) *ListManualTaskInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListManualTaskInstancesResponseBody) SetTotalSize(v int32) *ListManualTaskInstancesResponseBody {
	s.TotalSize = &v
	return s
}

type ListManualTaskInstancesResponseBodyData struct {
	// example:
	//
	// c-b933c5aac7f7***
	EmrClusterId *string `json:"EmrClusterId,omitempty" xml:"EmrClusterId,omitempty"`
	// example:
	//
	// 2024-03-27 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// application_123_***
	ExternalAppId *string `json:"ExternalAppId,omitempty" xml:"ExternalAppId,omitempty"`
	// example:
	//
	// mti-3q9jo749ne5****
	ManualTaskInstanceId *string `json:"ManualTaskInstanceId,omitempty" xml:"ManualTaskInstanceId,omitempty"`
	// example:
	//
	// test
	ManualTaskInstanceName *string `json:"ManualTaskInstanceName,omitempty" xml:"ManualTaskInstanceName,omitempty"`
	// example:
	//
	// wg-3q9jo749ne5****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 2024-03-27 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2024-03-27 00:00:00
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	// example:
	//
	// {
	//
	//     "yarnUser": "",
	//
	//     "conditionResult": "null",
	//
	//     "rawScript": "sleep 300",
	//
	//     "submitOnYarnFlag": false,
	//
	//     "emrClusterId": "",
	//
	//     "yarnPriority": "",
	//
	//     "dependence": "null",
	//
	//     "yarnMemory": "",
	//
	//     "localParams": [],
	//
	//     "switchResult": "null",
	//
	//     "resourceList": [],
	//
	//     "yarnQueue": "",
	//
	//     "yarnVCores": "",
	//
	//     "associateManualTaskFlag": false
	//
	// }
	TaskParams *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty"`
	// example:
	//
	// SHELL
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s ListManualTaskInstancesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListManualTaskInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListManualTaskInstancesResponseBodyData) SetEmrClusterId(v string) *ListManualTaskInstancesResponseBodyData {
	s.EmrClusterId = &v
	return s
}

func (s *ListManualTaskInstancesResponseBodyData) SetEndTime(v string) *ListManualTaskInstancesResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListManualTaskInstancesResponseBodyData) SetExternalAppId(v string) *ListManualTaskInstancesResponseBodyData {
	s.ExternalAppId = &v
	return s
}

func (s *ListManualTaskInstancesResponseBodyData) SetManualTaskInstanceId(v string) *ListManualTaskInstancesResponseBodyData {
	s.ManualTaskInstanceId = &v
	return s
}

func (s *ListManualTaskInstancesResponseBodyData) SetManualTaskInstanceName(v string) *ListManualTaskInstancesResponseBodyData {
	s.ManualTaskInstanceName = &v
	return s
}

func (s *ListManualTaskInstancesResponseBodyData) SetResourceGroupId(v string) *ListManualTaskInstancesResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *ListManualTaskInstancesResponseBodyData) SetStartTime(v string) *ListManualTaskInstancesResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListManualTaskInstancesResponseBodyData) SetStatus(v string) *ListManualTaskInstancesResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListManualTaskInstancesResponseBodyData) SetSubmitTime(v string) *ListManualTaskInstancesResponseBodyData {
	s.SubmitTime = &v
	return s
}

func (s *ListManualTaskInstancesResponseBodyData) SetTaskParams(v string) *ListManualTaskInstancesResponseBodyData {
	s.TaskParams = &v
	return s
}

func (s *ListManualTaskInstancesResponseBodyData) SetTaskType(v string) *ListManualTaskInstancesResponseBodyData {
	s.TaskType = &v
	return s
}

type ListManualTaskInstancesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListManualTaskInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListManualTaskInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListManualTaskInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListManualTaskInstancesResponse) SetHeaders(v map[string]*string) *ListManualTaskInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListManualTaskInstancesResponse) SetStatusCode(v int32) *ListManualTaskInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListManualTaskInstancesResponse) SetBody(v *ListManualTaskInstancesResponseBody) *ListManualTaskInstancesResponse {
	s.Body = v
	return s
}

type ListManualTasksRequest struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 123abc***
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// test
	SearchVal *string `json:"searchVal,omitempty" xml:"searchVal,omitempty"`
	// example:
	//
	// SHELL
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123***
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListManualTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListManualTasksRequest) GoString() string {
	return s.String()
}

func (s *ListManualTasksRequest) SetMaxResults(v int32) *ListManualTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListManualTasksRequest) SetNextToken(v string) *ListManualTasksRequest {
	s.NextToken = &v
	return s
}

func (s *ListManualTasksRequest) SetSearchVal(v string) *ListManualTasksRequest {
	s.SearchVal = &v
	return s
}

func (s *ListManualTasksRequest) SetTaskType(v string) *ListManualTasksRequest {
	s.TaskType = &v
	return s
}

func (s *ListManualTasksRequest) SetWorkspaceId(v string) *ListManualTasksRequest {
	s.WorkspaceId = &v
	return s
}

type ListManualTasksResponseBody struct {
	Data []*ListManualTasksResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 123abc***
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 10
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListManualTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListManualTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListManualTasksResponseBody) SetData(v []*ListManualTasksResponseBodyData) *ListManualTasksResponseBody {
	s.Data = v
	return s
}

func (s *ListManualTasksResponseBody) SetNextToken(v string) *ListManualTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListManualTasksResponseBody) SetRequestId(v string) *ListManualTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListManualTasksResponseBody) SetTotalSize(v int32) *ListManualTasksResponseBody {
	s.TotalSize = &v
	return s
}

type ListManualTasksResponseBodyData struct {
	// example:
	//
	// 2024-03-27 00:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// mt-3q9jo749ne5****
	ManualTaskId *string `json:"ManualTaskId,omitempty" xml:"ManualTaskId,omitempty"`
	// example:
	//
	// test
	ManualTaskName *string `json:"ManualTaskName,omitempty" xml:"ManualTaskName,omitempty"`
	// example:
	//
	// mtd-oy98v7n43el****
	ParentDirectoryId *string `json:"ParentDirectoryId,omitempty" xml:"ParentDirectoryId,omitempty"`
	// example:
	//
	// p-3q9jo749ne5****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// r-oy98v7n43el****
	ResourceIds *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// example:
	//
	// {
	//
	//     "yarnUser": "",
	//
	//     "conditionResult": "null",
	//
	//     "rawScript": "sleep 300",
	//
	//     "submitOnYarnFlag": false,
	//
	//     "emrClusterId": "",
	//
	//     "yarnPriority": "",
	//
	//     "dependence": "null",
	//
	//     "yarnMemory": "",
	//
	//     "localParams": [],
	//
	//     "switchResult": "null",
	//
	//     "resourceList": [],
	//
	//     "yarnQueue": "",
	//
	//     "yarnVCores": "",
	//
	//     "associateManualTaskFlag": false
	//
	// }
	TaskParams *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty"`
	// example:
	//
	// SHELL
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// example:
	//
	// 2024-03-27 00:00:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListManualTasksResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListManualTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListManualTasksResponseBodyData) SetCreateTime(v string) *ListManualTasksResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListManualTasksResponseBodyData) SetDescription(v string) *ListManualTasksResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListManualTasksResponseBodyData) SetManualTaskId(v string) *ListManualTasksResponseBodyData {
	s.ManualTaskId = &v
	return s
}

func (s *ListManualTasksResponseBodyData) SetManualTaskName(v string) *ListManualTasksResponseBodyData {
	s.ManualTaskName = &v
	return s
}

func (s *ListManualTasksResponseBodyData) SetParentDirectoryId(v string) *ListManualTasksResponseBodyData {
	s.ParentDirectoryId = &v
	return s
}

func (s *ListManualTasksResponseBodyData) SetProjectId(v string) *ListManualTasksResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *ListManualTasksResponseBodyData) SetResourceIds(v string) *ListManualTasksResponseBodyData {
	s.ResourceIds = &v
	return s
}

func (s *ListManualTasksResponseBodyData) SetTaskParams(v string) *ListManualTasksResponseBodyData {
	s.TaskParams = &v
	return s
}

func (s *ListManualTasksResponseBodyData) SetTaskType(v string) *ListManualTasksResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *ListManualTasksResponseBodyData) SetUpdateTime(v string) *ListManualTasksResponseBodyData {
	s.UpdateTime = &v
	return s
}

type ListManualTasksResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListManualTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListManualTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListManualTasksResponse) GoString() string {
	return s.String()
}

func (s *ListManualTasksResponse) SetHeaders(v map[string]*string) *ListManualTasksResponse {
	s.Headers = v
	return s
}

func (s *ListManualTasksResponse) SetStatusCode(v int32) *ListManualTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListManualTasksResponse) SetBody(v *ListManualTasksResponseBody) *ListManualTasksResponse {
	s.Body = v
	return s
}

type ListProjectsRequest struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 123abc***
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// test
	SearchVal *string `json:"searchVal,omitempty" xml:"searchVal,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListProjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsRequest) SetMaxResults(v int32) *ListProjectsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProjectsRequest) SetNextToken(v string) *ListProjectsRequest {
	s.NextToken = &v
	return s
}

func (s *ListProjectsRequest) SetSearchVal(v string) *ListProjectsRequest {
	s.SearchVal = &v
	return s
}

func (s *ListProjectsRequest) SetWorkspaceId(v string) *ListProjectsRequest {
	s.WorkspaceId = &v
	return s
}

type ListProjectsResponseBody struct {
	Data []*ListProjectsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 123abc***
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 10
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListProjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBody) SetData(v []*ListProjectsResponseBodyData) *ListProjectsResponseBody {
	s.Data = v
	return s
}

func (s *ListProjectsResponseBody) SetNextToken(v string) *ListProjectsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListProjectsResponseBody) SetRequestId(v string) *ListProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectsResponseBody) SetTotalSize(v int32) *ListProjectsResponseBody {
	s.TotalSize = &v
	return s
}

type ListProjectsResponseBodyData struct {
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// project_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// p-3q9jo749ne5****
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
}

func (s ListProjectsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyData) SetDescription(v string) *ListProjectsResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListProjectsResponseBodyData) SetName(v string) *ListProjectsResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListProjectsResponseBodyData) SetProjectId(v string) *ListProjectsResponseBodyData {
	s.ProjectId = &v
	return s
}

type ListProjectsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponse) GoString() string {
	return s.String()
}

func (s *ListProjectsResponse) SetHeaders(v map[string]*string) *ListProjectsResponse {
	s.Headers = v
	return s
}

func (s *ListProjectsResponse) SetStatusCode(v int32) *ListProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectsResponse) SetBody(v *ListProjectsResponseBody) *ListProjectsResponse {
	s.Body = v
	return s
}

type ListTaskInstancesRequest struct {
	// example:
	//
	// 2024-03-27 00:00:00
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 123abc***
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// test
	SearchVal *string `json:"searchVal,omitempty" xml:"searchVal,omitempty"`
	// example:
	//
	// 2024-03-27 00:00:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// wi-3q9jo749ne5****
	WorkflowInstanceId *string `json:"workflowInstanceId,omitempty" xml:"workflowInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123***
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListTaskInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTaskInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListTaskInstancesRequest) SetEndTime(v string) *ListTaskInstancesRequest {
	s.EndTime = &v
	return s
}

func (s *ListTaskInstancesRequest) SetMaxResults(v int32) *ListTaskInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTaskInstancesRequest) SetNextToken(v string) *ListTaskInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTaskInstancesRequest) SetSearchVal(v string) *ListTaskInstancesRequest {
	s.SearchVal = &v
	return s
}

func (s *ListTaskInstancesRequest) SetStartTime(v string) *ListTaskInstancesRequest {
	s.StartTime = &v
	return s
}

func (s *ListTaskInstancesRequest) SetStatus(v string) *ListTaskInstancesRequest {
	s.Status = &v
	return s
}

func (s *ListTaskInstancesRequest) SetWorkflowInstanceId(v string) *ListTaskInstancesRequest {
	s.WorkflowInstanceId = &v
	return s
}

func (s *ListTaskInstancesRequest) SetWorkspaceId(v string) *ListTaskInstancesRequest {
	s.WorkspaceId = &v
	return s
}

type ListTaskInstancesResponseBody struct {
	Data []*ListTaskInstancesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 123abc***
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 9E3A7161-EB7B-172B-8D18-FFB06BA38***
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 10
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListTaskInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTaskInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTaskInstancesResponseBody) SetData(v []*ListTaskInstancesResponseBodyData) *ListTaskInstancesResponseBody {
	s.Data = v
	return s
}

func (s *ListTaskInstancesResponseBody) SetNextToken(v string) *ListTaskInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTaskInstancesResponseBody) SetRequestId(v string) *ListTaskInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTaskInstancesResponseBody) SetTotalSize(v int32) *ListTaskInstancesResponseBody {
	s.TotalSize = &v
	return s
}

type ListTaskInstancesResponseBodyData struct {
	// example:
	//
	// 0
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// c-b933c5aac7f7***
	EmrClusterId *string `json:"EmrClusterId,omitempty" xml:"EmrClusterId,omitempty"`
	// example:
	//
	// 2024-03-27 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// application_123_***
	ExternalAppId *string `json:"ExternalAppId,omitempty" xml:"ExternalAppId,omitempty"`
	// example:
	//
	// wg-3q9jo749ne5****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 0
	RetryTimes *int32 `json:"RetryTimes,omitempty" xml:"RetryTimes,omitempty"`
	// example:
	//
	// 2024-03-27 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2024-03-27 00:00:00
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	// example:
	//
	// t-3q9jo749ne5****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// ti-3q9jo749ne5****
	TaskInstanceId *string `json:"TaskInstanceId,omitempty" xml:"TaskInstanceId,omitempty"`
	// example:
	//
	// test
	TaskInstanceName *string `json:"TaskInstanceName,omitempty" xml:"TaskInstanceName,omitempty"`
	// example:
	//
	// {
	//
	//     "yarnUser": "",
	//
	//     "conditionResult": "null",
	//
	//     "rawScript": "sleep 300",
	//
	//     "submitOnYarnFlag": false,
	//
	//     "emrClusterId": "",
	//
	//     "yarnPriority": "",
	//
	//     "dependence": "null",
	//
	//     "yarnMemory": "",
	//
	//     "localParams": [],
	//
	//     "switchResult": "null",
	//
	//     "resourceList": [],
	//
	//     "yarnQueue": "",
	//
	//     "yarnVCores": "",
	//
	//     "associateManualTaskFlag": false
	//
	// }
	TaskParams *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty"`
	// example:
	//
	// SHELL
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// example:
	//
	// 1
	TaskVersion *string `json:"TaskVersion,omitempty" xml:"TaskVersion,omitempty"`
	// example:
	//
	// wi-3q9jo749ne5****
	WorkflowInstanceId *string `json:"WorkflowInstanceId,omitempty" xml:"WorkflowInstanceId,omitempty"`
}

func (s ListTaskInstancesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTaskInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTaskInstancesResponseBodyData) SetDryRun(v string) *ListTaskInstancesResponseBodyData {
	s.DryRun = &v
	return s
}

func (s *ListTaskInstancesResponseBodyData) SetEmrClusterId(v string) *ListTaskInstancesResponseBodyData {
	s.EmrClusterId = &v
	return s
}

func (s *ListTaskInstancesResponseBodyData) SetEndTime(v string) *ListTaskInstancesResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListTaskInstancesResponseBodyData) SetExternalAppId(v string) *ListTaskInstancesResponseBodyData {
	s.ExternalAppId = &v
	return s
}

func (s *ListTaskInstancesResponseBodyData) SetResourceGroupId(v string) *ListTaskInstancesResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTaskInstancesResponseBodyData) SetRetryTimes(v int32) *ListTaskInstancesResponseBodyData {
	s.RetryTimes = &v
	return s
}

func (s *ListTaskInstancesResponseBodyData) SetStartTime(v string) *ListTaskInstancesResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListTaskInstancesResponseBodyData) SetStatus(v string) *ListTaskInstancesResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListTaskInstancesResponseBodyData) SetSubmitTime(v string) *ListTaskInstancesResponseBodyData {
	s.SubmitTime = &v
	return s
}

func (s *ListTaskInstancesResponseBodyData) SetTaskId(v string) *ListTaskInstancesResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ListTaskInstancesResponseBodyData) SetTaskInstanceId(v string) *ListTaskInstancesResponseBodyData {
	s.TaskInstanceId = &v
	return s
}

func (s *ListTaskInstancesResponseBodyData) SetTaskInstanceName(v string) *ListTaskInstancesResponseBodyData {
	s.TaskInstanceName = &v
	return s
}

func (s *ListTaskInstancesResponseBodyData) SetTaskParams(v string) *ListTaskInstancesResponseBodyData {
	s.TaskParams = &v
	return s
}

func (s *ListTaskInstancesResponseBodyData) SetTaskType(v string) *ListTaskInstancesResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *ListTaskInstancesResponseBodyData) SetTaskVersion(v string) *ListTaskInstancesResponseBodyData {
	s.TaskVersion = &v
	return s
}

func (s *ListTaskInstancesResponseBodyData) SetWorkflowInstanceId(v string) *ListTaskInstancesResponseBodyData {
	s.WorkflowInstanceId = &v
	return s
}

type ListTaskInstancesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTaskInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTaskInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTaskInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListTaskInstancesResponse) SetHeaders(v map[string]*string) *ListTaskInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListTaskInstancesResponse) SetStatusCode(v int32) *ListTaskInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTaskInstancesResponse) SetBody(v *ListTaskInstancesResponseBody) *ListTaskInstancesResponse {
	s.Body = v
	return s
}

type ListTasksRequest struct {
	// example:
	//
	// 10
	MaxResults *string `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 123abc***
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// test
	SearchVal *string `json:"searchVal,omitempty" xml:"searchVal,omitempty"`
	// example:
	//
	// SHELL
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	// example:
	//
	// w-n72kong0832****
	WorkflowId *string `json:"workflowId,omitempty" xml:"workflowId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123***
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTasksRequest) GoString() string {
	return s.String()
}

func (s *ListTasksRequest) SetMaxResults(v string) *ListTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTasksRequest) SetNextToken(v string) *ListTasksRequest {
	s.NextToken = &v
	return s
}

func (s *ListTasksRequest) SetSearchVal(v string) *ListTasksRequest {
	s.SearchVal = &v
	return s
}

func (s *ListTasksRequest) SetTaskType(v string) *ListTasksRequest {
	s.TaskType = &v
	return s
}

func (s *ListTasksRequest) SetWorkflowId(v string) *ListTasksRequest {
	s.WorkflowId = &v
	return s
}

func (s *ListTasksRequest) SetWorkspaceId(v string) *ListTasksRequest {
	s.WorkspaceId = &v
	return s
}

type ListTasksResponseBody struct {
	Data []*ListTasksResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 123abc***
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 10
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBody) SetData(v []*ListTasksResponseBodyData) *ListTasksResponseBody {
	s.Data = v
	return s
}

func (s *ListTasksResponseBody) SetNextToken(v string) *ListTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTasksResponseBody) SetRequestId(v string) *ListTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTasksResponseBody) SetTotalSize(v int32) *ListTasksResponseBody {
	s.TotalSize = &v
	return s
}

type ListTasksResponseBodyData struct {
	// example:
	//
	// 2024-03-27 00:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 0
	DelayTime *int32 `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1
	FailRetryInterval *int32 `json:"FailRetryInterval,omitempty" xml:"FailRetryInterval,omitempty"`
	// example:
	//
	// 0
	FailRetryTimes *int32 `json:"FailRetryTimes,omitempty" xml:"FailRetryTimes,omitempty"`
	// example:
	//
	// YES
	Flag *string `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// example:
	//
	// p-3q9jo749ne5****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// r-3q9jo749ne5****
	ResourceIds *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// example:
	//
	// t-3q9jo749ne5****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// test
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// {
	//
	//     "yarnUser": "",
	//
	//     "conditionResult": "null",
	//
	//     "rawScript": "sleep 300",
	//
	//     "submitOnYarnFlag": false,
	//
	//     "emrClusterId": "",
	//
	//     "yarnPriority": "",
	//
	//     "dependence": "null",
	//
	//     "yarnMemory": "",
	//
	//     "localParams": [],
	//
	//     "switchResult": "null",
	//
	//     "resourceList": [],
	//
	//     "yarnQueue": "",
	//
	//     "yarnVCores": "",
	//
	//     "associateManualTaskFlag": false
	//
	// }
	TaskParams *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty"`
	// example:
	//
	// MEDIUM
	TaskPriority *string `json:"TaskPriority,omitempty" xml:"TaskPriority,omitempty"`
	// example:
	//
	// SHELL
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// example:
	//
	// 0
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// example:
	//
	// CLOSE
	TimeoutFlag *string `json:"TimeoutFlag,omitempty" xml:"TimeoutFlag,omitempty"`
	// example:
	//
	// WARN
	TimeoutNotifyStrategy *string `json:"TimeoutNotifyStrategy,omitempty" xml:"TimeoutNotifyStrategy,omitempty"`
	// example:
	//
	// 2024-03-27 00:00:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListTasksResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyData) SetCreateTime(v string) *ListTasksResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListTasksResponseBodyData) SetDelayTime(v int32) *ListTasksResponseBodyData {
	s.DelayTime = &v
	return s
}

func (s *ListTasksResponseBodyData) SetDescription(v string) *ListTasksResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListTasksResponseBodyData) SetFailRetryInterval(v int32) *ListTasksResponseBodyData {
	s.FailRetryInterval = &v
	return s
}

func (s *ListTasksResponseBodyData) SetFailRetryTimes(v int32) *ListTasksResponseBodyData {
	s.FailRetryTimes = &v
	return s
}

func (s *ListTasksResponseBodyData) SetFlag(v string) *ListTasksResponseBodyData {
	s.Flag = &v
	return s
}

func (s *ListTasksResponseBodyData) SetProjectId(v string) *ListTasksResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *ListTasksResponseBodyData) SetResourceIds(v string) *ListTasksResponseBodyData {
	s.ResourceIds = &v
	return s
}

func (s *ListTasksResponseBodyData) SetTaskId(v string) *ListTasksResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ListTasksResponseBodyData) SetTaskName(v string) *ListTasksResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *ListTasksResponseBodyData) SetTaskParams(v string) *ListTasksResponseBodyData {
	s.TaskParams = &v
	return s
}

func (s *ListTasksResponseBodyData) SetTaskPriority(v string) *ListTasksResponseBodyData {
	s.TaskPriority = &v
	return s
}

func (s *ListTasksResponseBodyData) SetTaskType(v string) *ListTasksResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *ListTasksResponseBodyData) SetTimeout(v int32) *ListTasksResponseBodyData {
	s.Timeout = &v
	return s
}

func (s *ListTasksResponseBodyData) SetTimeoutFlag(v string) *ListTasksResponseBodyData {
	s.TimeoutFlag = &v
	return s
}

func (s *ListTasksResponseBodyData) SetTimeoutNotifyStrategy(v string) *ListTasksResponseBodyData {
	s.TimeoutNotifyStrategy = &v
	return s
}

func (s *ListTasksResponseBodyData) SetUpdateTime(v string) *ListTasksResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListTasksResponseBodyData) SetVersion(v string) *ListTasksResponseBodyData {
	s.Version = &v
	return s
}

type ListTasksResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponse) GoString() string {
	return s.String()
}

func (s *ListTasksResponse) SetHeaders(v map[string]*string) *ListTasksResponse {
	s.Headers = v
	return s
}

func (s *ListTasksResponse) SetStatusCode(v int32) *ListTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTasksResponse) SetBody(v *ListTasksResponseBody) *ListTasksResponse {
	s.Body = v
	return s
}

type ListWorkflowInstancesRequest struct {
	// example:
	//
	// 2024-03-27 00:00:00
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 123abc***
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 2024-03-27 00:00:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// w-3q9jo749ne5****
	WorkflowId *string `json:"workflowId,omitempty" xml:"workflowId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListWorkflowInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkflowInstancesRequest) SetEndTime(v string) *ListWorkflowInstancesRequest {
	s.EndTime = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetMaxResults(v int32) *ListWorkflowInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetNextToken(v string) *ListWorkflowInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetStartTime(v string) *ListWorkflowInstancesRequest {
	s.StartTime = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetStatus(v string) *ListWorkflowInstancesRequest {
	s.Status = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetWorkflowId(v string) *ListWorkflowInstancesRequest {
	s.WorkflowId = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetWorkspaceId(v string) *ListWorkflowInstancesRequest {
	s.WorkspaceId = &v
	return s
}

type ListWorkflowInstancesResponseBody struct {
	Data []*ListWorkflowInstancesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 123abc***
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 10
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListWorkflowInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkflowInstancesResponseBody) SetData(v []*ListWorkflowInstancesResponseBodyData) *ListWorkflowInstancesResponseBody {
	s.Data = v
	return s
}

func (s *ListWorkflowInstancesResponseBody) SetNextToken(v string) *ListWorkflowInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWorkflowInstancesResponseBody) SetRequestId(v string) *ListWorkflowInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkflowInstancesResponseBody) SetTotalSize(v int32) *ListWorkflowInstancesResponseBody {
	s.TotalSize = &v
	return s
}

type ListWorkflowInstancesResponseBodyData struct {
	// example:
	//
	// 2024-01-01 00:00:00
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// workflow_instance_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	ScheduleTime *string `json:"scheduleTime,omitempty" xml:"scheduleTime,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// w-3q9jo749ne5****
	WorkflowId *string `json:"workflowId,omitempty" xml:"workflowId,omitempty"`
	// example:
	//
	// wi-3q9jo749ne5****
	WorkflowInstanceId *string `json:"workflowInstanceId,omitempty" xml:"workflowInstanceId,omitempty"`
	// example:
	//
	// 1
	WorkflowVersion *int32 `json:"workflowVersion,omitempty" xml:"workflowVersion,omitempty"`
}

func (s ListWorkflowInstancesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWorkflowInstancesResponseBodyData) SetEndTime(v string) *ListWorkflowInstancesResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyData) SetName(v string) *ListWorkflowInstancesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyData) SetScheduleTime(v string) *ListWorkflowInstancesResponseBodyData {
	s.ScheduleTime = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyData) SetStartTime(v string) *ListWorkflowInstancesResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyData) SetStatus(v string) *ListWorkflowInstancesResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyData) SetWorkflowId(v string) *ListWorkflowInstancesResponseBodyData {
	s.WorkflowId = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyData) SetWorkflowInstanceId(v string) *ListWorkflowInstancesResponseBodyData {
	s.WorkflowInstanceId = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyData) SetWorkflowVersion(v int32) *ListWorkflowInstancesResponseBodyData {
	s.WorkflowVersion = &v
	return s
}

type ListWorkflowInstancesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkflowInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkflowInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListWorkflowInstancesResponse) SetHeaders(v map[string]*string) *ListWorkflowInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListWorkflowInstancesResponse) SetStatusCode(v int32) *ListWorkflowInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkflowInstancesResponse) SetBody(v *ListWorkflowInstancesResponseBody) *ListWorkflowInstancesResponse {
	s.Body = v
	return s
}

type ListWorkflowsRequest struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 123abc***
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// test
	SearchVal *string `json:"searchVal,omitempty" xml:"searchVal,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListWorkflowsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowsRequest) GoString() string {
	return s.String()
}

func (s *ListWorkflowsRequest) SetMaxResults(v int32) *ListWorkflowsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkflowsRequest) SetNextToken(v string) *ListWorkflowsRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkflowsRequest) SetSearchVal(v string) *ListWorkflowsRequest {
	s.SearchVal = &v
	return s
}

func (s *ListWorkflowsRequest) SetWorkspaceId(v string) *ListWorkflowsRequest {
	s.WorkspaceId = &v
	return s
}

type ListWorkflowsResponseBody struct {
	Data []*ListWorkflowsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 123abc****
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 10
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListWorkflowsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkflowsResponseBody) SetData(v []*ListWorkflowsResponseBodyData) *ListWorkflowsResponseBody {
	s.Data = v
	return s
}

func (s *ListWorkflowsResponseBody) SetNextToken(v string) *ListWorkflowsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWorkflowsResponseBody) SetRequestId(v string) *ListWorkflowsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkflowsResponseBody) SetTotalSize(v int32) *ListWorkflowsResponseBody {
	s.TotalSize = &v
	return s
}

type ListWorkflowsResponseBodyData struct {
	// example:
	//
	// 2024-01-01 00:00:00
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// workflow description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// workflow_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// wd-3q9jo749ne5****
	ParentDirectoryId *string `json:"parentDirectoryId,omitempty" xml:"parentDirectoryId,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// example:
	//
	// w-3q9jo749ne5****
	WorkflowId *string `json:"workflowId,omitempty" xml:"workflowId,omitempty"`
}

func (s ListWorkflowsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWorkflowsResponseBodyData) SetCreateTime(v string) *ListWorkflowsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListWorkflowsResponseBodyData) SetDescription(v string) *ListWorkflowsResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListWorkflowsResponseBodyData) SetName(v string) *ListWorkflowsResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListWorkflowsResponseBodyData) SetParentDirectoryId(v string) *ListWorkflowsResponseBodyData {
	s.ParentDirectoryId = &v
	return s
}

func (s *ListWorkflowsResponseBodyData) SetUpdateTime(v string) *ListWorkflowsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListWorkflowsResponseBodyData) SetWorkflowId(v string) *ListWorkflowsResponseBodyData {
	s.WorkflowId = &v
	return s
}

type ListWorkflowsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkflowsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkflowsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowsResponse) GoString() string {
	return s.String()
}

func (s *ListWorkflowsResponse) SetHeaders(v map[string]*string) *ListWorkflowsResponse {
	s.Headers = v
	return s
}

func (s *ListWorkflowsResponse) SetStatusCode(v int32) *ListWorkflowsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkflowsResponse) SetBody(v *ListWorkflowsResponseBody) *ListWorkflowsResponse {
	s.Body = v
	return s
}

type RunWorkflowRequest struct {
	// example:
	//
	// ag-n72kong0832****
	AlertGroupId *string `json:"alertGroupId,omitempty" xml:"alertGroupId,omitempty"`
	// example:
	//
	// NONE
	AlertStrategy *string `json:"alertStrategy,omitempty" xml:"alertStrategy,omitempty"`
	// example:
	//
	// OFF_MODE
	ComplementDependentMode *string `json:"complementDependentMode,omitempty" xml:"complementDependentMode,omitempty"`
	// example:
	//
	// 0
	DryRun *string `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// example:
	//
	// START_PROCESS
	ExecType *string `json:"execType,omitempty" xml:"execType,omitempty"`
	// example:
	//
	// 1
	ExpectedParallelismNumber *string `json:"expectedParallelismNumber,omitempty" xml:"expectedParallelismNumber,omitempty"`
	// example:
	//
	// END
	FailureStrategy *string `json:"failureStrategy,omitempty" xml:"failureStrategy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// wg-acfmv4opbs****
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// example:
	//
	// RUN_MODE_PARALLEL
	RunMode *string `json:"runMode,omitempty" xml:"runMode,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00,2024-01-02 00:00:00
	ScheduleTime *string `json:"scheduleTime,omitempty" xml:"scheduleTime,omitempty"`
	// example:
	//
	// {"key1":"value1"}
	StartParams *string `json:"startParams,omitempty" xml:"startParams,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// w-3q9jo749ne5****
	WorkflowId *string `json:"workflowId,omitempty" xml:"workflowId,omitempty"`
	// example:
	//
	// MEDIUM
	WorkflowInstancePriority *string `json:"workflowInstancePriority,omitempty" xml:"workflowInstancePriority,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s RunWorkflowRequest) String() string {
	return tea.Prettify(s)
}

func (s RunWorkflowRequest) GoString() string {
	return s.String()
}

func (s *RunWorkflowRequest) SetAlertGroupId(v string) *RunWorkflowRequest {
	s.AlertGroupId = &v
	return s
}

func (s *RunWorkflowRequest) SetAlertStrategy(v string) *RunWorkflowRequest {
	s.AlertStrategy = &v
	return s
}

func (s *RunWorkflowRequest) SetComplementDependentMode(v string) *RunWorkflowRequest {
	s.ComplementDependentMode = &v
	return s
}

func (s *RunWorkflowRequest) SetDryRun(v string) *RunWorkflowRequest {
	s.DryRun = &v
	return s
}

func (s *RunWorkflowRequest) SetExecType(v string) *RunWorkflowRequest {
	s.ExecType = &v
	return s
}

func (s *RunWorkflowRequest) SetExpectedParallelismNumber(v string) *RunWorkflowRequest {
	s.ExpectedParallelismNumber = &v
	return s
}

func (s *RunWorkflowRequest) SetFailureStrategy(v string) *RunWorkflowRequest {
	s.FailureStrategy = &v
	return s
}

func (s *RunWorkflowRequest) SetResourceGroupId(v string) *RunWorkflowRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *RunWorkflowRequest) SetRunMode(v string) *RunWorkflowRequest {
	s.RunMode = &v
	return s
}

func (s *RunWorkflowRequest) SetScheduleTime(v string) *RunWorkflowRequest {
	s.ScheduleTime = &v
	return s
}

func (s *RunWorkflowRequest) SetStartParams(v string) *RunWorkflowRequest {
	s.StartParams = &v
	return s
}

func (s *RunWorkflowRequest) SetWorkflowId(v string) *RunWorkflowRequest {
	s.WorkflowId = &v
	return s
}

func (s *RunWorkflowRequest) SetWorkflowInstancePriority(v string) *RunWorkflowRequest {
	s.WorkflowInstancePriority = &v
	return s
}

func (s *RunWorkflowRequest) SetWorkspaceId(v string) *RunWorkflowRequest {
	s.WorkspaceId = &v
	return s
}

type RunWorkflowResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RunWorkflowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *RunWorkflowResponseBody) SetRequestId(v string) *RunWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunWorkflowResponseBody) SetSuccess(v bool) *RunWorkflowResponseBody {
	s.Success = &v
	return s
}

type RunWorkflowResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunWorkflowResponse) String() string {
	return tea.Prettify(s)
}

func (s RunWorkflowResponse) GoString() string {
	return s.String()
}

func (s *RunWorkflowResponse) SetHeaders(v map[string]*string) *RunWorkflowResponse {
	s.Headers = v
	return s
}

func (s *RunWorkflowResponse) SetStatusCode(v int32) *RunWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *RunWorkflowResponse) SetBody(v *RunWorkflowResponseBody) *RunWorkflowResponse {
	s.Body = v
	return s
}

type UpdateWorkflowRequest struct {
	// example:
	//
	// ag-n72kong0832****
	AlertGroupId *string `json:"alertGroupId,omitempty" xml:"alertGroupId,omitempty"`
	// example:
	//
	// NONE
	AlertStrategy *string `json:"alertStrategy,omitempty" xml:"alertStrategy,omitempty"`
	// example:
	//
	// 0 0 	- 	- 	- ? *
	CronExpr *string `json:"cronExpr,omitempty" xml:"cronExpr,omitempty"`
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// PARALLEL
	ExecutionType *string `json:"executionType,omitempty" xml:"executionType,omitempty"`
	// example:
	//
	// END
	FailureStrategy *string `json:"failureStrategy,omitempty" xml:"failureStrategy,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// wd-n72kong0832****
	ParentDirectoryId *string `json:"parentDirectoryId,omitempty" xml:"parentDirectoryId,omitempty"`
	// example:
	//
	// wg-acfmv4opbs****
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	ScheduleEndTime *string `json:"scheduleEndTime,omitempty" xml:"scheduleEndTime,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	ScheduleStartTime *string `json:"scheduleStartTime,omitempty" xml:"scheduleStartTime,omitempty"`
	// example:
	//
	// OFFLINE
	ScheduleState *string `json:"scheduleState,omitempty" xml:"scheduleState,omitempty"`
	// example:
	//
	// [{"taskId":"t1","name":"t1","taskParams":{"rawScript":"echo 1"},"taskType":"SHELL"}]
	TaskDefinitionJson *string `json:"taskDefinitionJson,omitempty" xml:"taskDefinitionJson,omitempty"`
	// example:
	//
	// [{"preTaskId":"0", "postTaskId":"t1"}]
	TaskRelationJson *string `json:"taskRelationJson,omitempty" xml:"taskRelationJson,omitempty"`
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
	// example:
	//
	// 10
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// example:
	//
	// MEDIUM
	WorkflowInstancePriority *string `json:"workflowInstancePriority,omitempty" xml:"workflowInstancePriority,omitempty"`
	// example:
	//
	// [{"prop":"key1","value":"value1"}]
	WorkflowParams *string `json:"workflowParams,omitempty" xml:"workflowParams,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123***
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s UpdateWorkflowRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkflowRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowRequest) SetAlertGroupId(v string) *UpdateWorkflowRequest {
	s.AlertGroupId = &v
	return s
}

func (s *UpdateWorkflowRequest) SetAlertStrategy(v string) *UpdateWorkflowRequest {
	s.AlertStrategy = &v
	return s
}

func (s *UpdateWorkflowRequest) SetCronExpr(v string) *UpdateWorkflowRequest {
	s.CronExpr = &v
	return s
}

func (s *UpdateWorkflowRequest) SetDescription(v string) *UpdateWorkflowRequest {
	s.Description = &v
	return s
}

func (s *UpdateWorkflowRequest) SetExecutionType(v string) *UpdateWorkflowRequest {
	s.ExecutionType = &v
	return s
}

func (s *UpdateWorkflowRequest) SetFailureStrategy(v string) *UpdateWorkflowRequest {
	s.FailureStrategy = &v
	return s
}

func (s *UpdateWorkflowRequest) SetName(v string) *UpdateWorkflowRequest {
	s.Name = &v
	return s
}

func (s *UpdateWorkflowRequest) SetParentDirectoryId(v string) *UpdateWorkflowRequest {
	s.ParentDirectoryId = &v
	return s
}

func (s *UpdateWorkflowRequest) SetResourceGroupId(v string) *UpdateWorkflowRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateWorkflowRequest) SetScheduleEndTime(v string) *UpdateWorkflowRequest {
	s.ScheduleEndTime = &v
	return s
}

func (s *UpdateWorkflowRequest) SetScheduleStartTime(v string) *UpdateWorkflowRequest {
	s.ScheduleStartTime = &v
	return s
}

func (s *UpdateWorkflowRequest) SetScheduleState(v string) *UpdateWorkflowRequest {
	s.ScheduleState = &v
	return s
}

func (s *UpdateWorkflowRequest) SetTaskDefinitionJson(v string) *UpdateWorkflowRequest {
	s.TaskDefinitionJson = &v
	return s
}

func (s *UpdateWorkflowRequest) SetTaskRelationJson(v string) *UpdateWorkflowRequest {
	s.TaskRelationJson = &v
	return s
}

func (s *UpdateWorkflowRequest) SetTimeZone(v string) *UpdateWorkflowRequest {
	s.TimeZone = &v
	return s
}

func (s *UpdateWorkflowRequest) SetTimeout(v int32) *UpdateWorkflowRequest {
	s.Timeout = &v
	return s
}

func (s *UpdateWorkflowRequest) SetWorkflowInstancePriority(v string) *UpdateWorkflowRequest {
	s.WorkflowInstancePriority = &v
	return s
}

func (s *UpdateWorkflowRequest) SetWorkflowParams(v string) *UpdateWorkflowRequest {
	s.WorkflowParams = &v
	return s
}

func (s *UpdateWorkflowRequest) SetWorkspaceId(v string) *UpdateWorkflowRequest {
	s.WorkspaceId = &v
	return s
}

type UpdateWorkflowResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateWorkflowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowResponseBody) SetRequestId(v string) *UpdateWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkflowResponseBody) SetSuccess(v bool) *UpdateWorkflowResponseBody {
	s.Success = &v
	return s
}

type UpdateWorkflowResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkflowResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkflowResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowResponse) SetHeaders(v map[string]*string) *UpdateWorkflowResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkflowResponse) SetStatusCode(v int32) *UpdateWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkflowResponse) SetBody(v *UpdateWorkflowResponseBody) *UpdateWorkflowResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("emrstudio"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建工作流
//
// @param request - CreateWorkflowRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkflowResponse
func (client *Client) CreateWorkflowWithOptions(projectId *string, request *CreateWorkflowRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateWorkflowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertGroupId)) {
		query["alertGroupId"] = request.AlertGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.AlertStrategy)) {
		query["alertStrategy"] = request.AlertStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.CronExpr)) {
		query["cronExpr"] = request.CronExpr
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutionType)) {
		query["executionType"] = request.ExecutionType
	}

	if !tea.BoolValue(util.IsUnset(request.FailureStrategy)) {
		query["failureStrategy"] = request.FailureStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ParentDirectoryId)) {
		query["parentDirectoryId"] = request.ParentDirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleEndTime)) {
		query["scheduleEndTime"] = request.ScheduleEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleStartTime)) {
		query["scheduleStartTime"] = request.ScheduleStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleState)) {
		query["scheduleState"] = request.ScheduleState
	}

	if !tea.BoolValue(util.IsUnset(request.TaskDefinitionJson)) {
		query["taskDefinitionJson"] = request.TaskDefinitionJson
	}

	if !tea.BoolValue(util.IsUnset(request.TaskRelationJson)) {
		query["taskRelationJson"] = request.TaskRelationJson
	}

	if !tea.BoolValue(util.IsUnset(request.TimeZone)) {
		query["timeZone"] = request.TimeZone
	}

	if !tea.BoolValue(util.IsUnset(request.Timeout)) {
		query["timeout"] = request.Timeout
	}

	if !tea.BoolValue(util.IsUnset(request.WorkflowInstancePriority)) {
		query["workflowInstancePriority"] = request.WorkflowInstancePriority
	}

	if !tea.BoolValue(util.IsUnset(request.WorkflowParams)) {
		query["workflowParams"] = request.WorkflowParams
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWorkflow"),
		Version:     tea.String("2024-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dolphinscheduler/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectId)) + "/workflows"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWorkflowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建工作流
//
// @param request - CreateWorkflowRequest
//
// @return CreateWorkflowResponse
func (client *Client) CreateWorkflow(projectId *string, request *CreateWorkflowRequest) (_result *CreateWorkflowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateWorkflowResponse{}
	_body, _err := client.CreateWorkflowWithOptions(projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除工作流
//
// @param request - DeleteWorkflowRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkflowResponse
func (client *Client) DeleteWorkflowWithOptions(projectId *string, workflowId *string, request *DeleteWorkflowRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteWorkflowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteWorkflow"),
		Version:     tea.String("2024-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dolphinscheduler/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectId)) + "/workflows/" + tea.StringValue(openapiutil.GetEncodeParam(workflowId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteWorkflowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除工作流
//
// @param request - DeleteWorkflowRequest
//
// @return DeleteWorkflowResponse
func (client *Client) DeleteWorkflow(projectId *string, workflowId *string, request *DeleteWorkflowRequest) (_result *DeleteWorkflowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteWorkflowResponse{}
	_body, _err := client.DeleteWorkflowWithOptions(projectId, workflowId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取手动任务定义
//
// @param request - DescribeManualTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeManualTaskResponse
func (client *Client) DescribeManualTaskWithOptions(projectId *string, manualTaskId *string, request *DescribeManualTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeManualTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeManualTask"),
		Version:     tea.String("2024-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dolphinscheduler/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectId)) + "/manualTasks/" + tea.StringValue(openapiutil.GetEncodeParam(manualTaskId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeManualTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取手动任务定义
//
// @param request - DescribeManualTaskRequest
//
// @return DescribeManualTaskResponse
func (client *Client) DescribeManualTask(projectId *string, manualTaskId *string, request *DescribeManualTaskRequest) (_result *DescribeManualTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeManualTaskResponse{}
	_body, _err := client.DescribeManualTaskWithOptions(projectId, manualTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取手动任务实例
//
// @param request - DescribeManualTaskInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeManualTaskInstanceResponse
func (client *Client) DescribeManualTaskInstanceWithOptions(manualTaskInstanceId *string, projectId *string, request *DescribeManualTaskInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeManualTaskInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeManualTaskInstance"),
		Version:     tea.String("2024-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dolphinscheduler/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectId)) + "/manualTaskInstances/" + tea.StringValue(openapiutil.GetEncodeParam(manualTaskInstanceId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeManualTaskInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取手动任务实例
//
// @param request - DescribeManualTaskInstanceRequest
//
// @return DescribeManualTaskInstanceResponse
func (client *Client) DescribeManualTaskInstance(manualTaskInstanceId *string, projectId *string, request *DescribeManualTaskInstanceRequest) (_result *DescribeManualTaskInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeManualTaskInstanceResponse{}
	_body, _err := client.DescribeManualTaskInstanceWithOptions(manualTaskInstanceId, projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取项目详情
//
// @param request - DescribeProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProjectResponse
func (client *Client) DescribeProjectWithOptions(projectId *string, request *DescribeProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeProject"),
		Version:     tea.String("2024-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dolphinscheduler/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取项目详情
//
// @param request - DescribeProjectRequest
//
// @return DescribeProjectResponse
func (client *Client) DescribeProject(projectId *string, request *DescribeProjectRequest) (_result *DescribeProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeProjectResponse{}
	_body, _err := client.DescribeProjectWithOptions(projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询任务定义
//
// @param request - DescribeTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTaskResponse
func (client *Client) DescribeTaskWithOptions(workflowId *string, projectId *string, taskId *string, request *DescribeTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTask"),
		Version:     tea.String("2024-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dolphinscheduler/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectId)) + "/workflows/" + tea.StringValue(openapiutil.GetEncodeParam(workflowId)) + "/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(taskId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询任务定义
//
// @param request - DescribeTaskRequest
//
// @return DescribeTaskResponse
func (client *Client) DescribeTask(workflowId *string, projectId *string, taskId *string, request *DescribeTaskRequest) (_result *DescribeTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeTaskResponse{}
	_body, _err := client.DescribeTaskWithOptions(workflowId, projectId, taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取任务实例
//
// @param request - DescribeTaskInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTaskInstanceResponse
func (client *Client) DescribeTaskInstanceWithOptions(projectId *string, workflowInstanceId *string, taskInstanceId *string, request *DescribeTaskInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeTaskInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTaskInstance"),
		Version:     tea.String("2024-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dolphinscheduler/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectId)) + "/workflows/" + tea.StringValue(openapiutil.GetEncodeParam(workflowInstanceId)) + "/taskInstances/" + tea.StringValue(openapiutil.GetEncodeParam(taskInstanceId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTaskInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务实例
//
// @param request - DescribeTaskInstanceRequest
//
// @return DescribeTaskInstanceResponse
func (client *Client) DescribeTaskInstance(projectId *string, workflowInstanceId *string, taskInstanceId *string, request *DescribeTaskInstanceRequest) (_result *DescribeTaskInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeTaskInstanceResponse{}
	_body, _err := client.DescribeTaskInstanceWithOptions(projectId, workflowInstanceId, taskInstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取工作流详情
//
// @param request - DescribeWorkflowRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWorkflowResponse
func (client *Client) DescribeWorkflowWithOptions(projectId *string, workflowId *string, request *DescribeWorkflowRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeWorkflowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeWorkflow"),
		Version:     tea.String("2024-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dolphinscheduler/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectId)) + "/workflows/" + tea.StringValue(openapiutil.GetEncodeParam(workflowId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeWorkflowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工作流详情
//
// @param request - DescribeWorkflowRequest
//
// @return DescribeWorkflowResponse
func (client *Client) DescribeWorkflow(projectId *string, workflowId *string, request *DescribeWorkflowRequest) (_result *DescribeWorkflowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeWorkflowResponse{}
	_body, _err := client.DescribeWorkflowWithOptions(projectId, workflowId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取工作流实例详情
//
// @param request - DescribeWorkflowInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWorkflowInstanceResponse
func (client *Client) DescribeWorkflowInstanceWithOptions(projectId *string, workflowInstanceId *string, request *DescribeWorkflowInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeWorkflowInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeWorkflowInstance"),
		Version:     tea.String("2024-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dolphinscheduler/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectId)) + "/workflowInstances/" + tea.StringValue(openapiutil.GetEncodeParam(workflowInstanceId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeWorkflowInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工作流实例详情
//
// @param request - DescribeWorkflowInstanceRequest
//
// @return DescribeWorkflowInstanceResponse
func (client *Client) DescribeWorkflowInstance(projectId *string, workflowInstanceId *string, request *DescribeWorkflowInstanceRequest) (_result *DescribeWorkflowInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeWorkflowInstanceResponse{}
	_body, _err := client.DescribeWorkflowInstanceWithOptions(projectId, workflowInstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取手动任务实例列表
//
// @param request - ListManualTaskInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListManualTaskInstancesResponse
func (client *Client) ListManualTaskInstancesWithOptions(projectId *string, request *ListManualTaskInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListManualTaskInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.SearchVal)) {
		query["searchVal"] = request.SearchVal
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListManualTaskInstances"),
		Version:     tea.String("2024-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dolphinscheduler/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectId)) + "/manualTaskInstances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListManualTaskInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取手动任务实例列表
//
// @param request - ListManualTaskInstancesRequest
//
// @return ListManualTaskInstancesResponse
func (client *Client) ListManualTaskInstances(projectId *string, request *ListManualTaskInstancesRequest) (_result *ListManualTaskInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListManualTaskInstancesResponse{}
	_body, _err := client.ListManualTaskInstancesWithOptions(projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询手动任务定义列表
//
// @param request - ListManualTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListManualTasksResponse
func (client *Client) ListManualTasksWithOptions(projectId *string, request *ListManualTasksRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListManualTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.SearchVal)) {
		query["searchVal"] = request.SearchVal
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["taskType"] = request.TaskType
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListManualTasks"),
		Version:     tea.String("2024-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dolphinscheduler/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectId)) + "/manualTasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListManualTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询手动任务定义列表
//
// @param request - ListManualTasksRequest
//
// @return ListManualTasksResponse
func (client *Client) ListManualTasks(projectId *string, request *ListManualTasksRequest) (_result *ListManualTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListManualTasksResponse{}
	_body, _err := client.ListManualTasksWithOptions(projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取项目详情
//
// @param request - ListProjectsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectsResponse
func (client *Client) ListProjectsWithOptions(request *ListProjectsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.SearchVal)) {
		query["searchVal"] = request.SearchVal
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProjects"),
		Version:     tea.String("2024-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dolphinscheduler/v3/projects"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取项目详情
//
// @param request - ListProjectsRequest
//
// @return ListProjectsResponse
func (client *Client) ListProjects(request *ListProjectsRequest) (_result *ListProjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProjectsResponse{}
	_body, _err := client.ListProjectsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取任务实例列表
//
// @param request - ListTaskInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTaskInstancesResponse
func (client *Client) ListTaskInstancesWithOptions(projectId *string, request *ListTaskInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTaskInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.SearchVal)) {
		query["searchVal"] = request.SearchVal
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.WorkflowInstanceId)) {
		query["workflowInstanceId"] = request.WorkflowInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTaskInstances"),
		Version:     tea.String("2024-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dolphinscheduler/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectId)) + "/taskInstances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTaskInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务实例列表
//
// @param request - ListTaskInstancesRequest
//
// @return ListTaskInstancesResponse
func (client *Client) ListTaskInstances(projectId *string, request *ListTaskInstancesRequest) (_result *ListTaskInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTaskInstancesResponse{}
	_body, _err := client.ListTaskInstancesWithOptions(projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询任务定义列表
//
// @param request - ListTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTasksResponse
func (client *Client) ListTasksWithOptions(projectId *string, request *ListTasksRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.SearchVal)) {
		query["searchVal"] = request.SearchVal
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["taskType"] = request.TaskType
	}

	if !tea.BoolValue(util.IsUnset(request.WorkflowId)) {
		query["workflowId"] = request.WorkflowId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTasks"),
		Version:     tea.String("2024-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dolphinscheduler/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectId)) + "/tasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询任务定义列表
//
// @param request - ListTasksRequest
//
// @return ListTasksResponse
func (client *Client) ListTasks(projectId *string, request *ListTasksRequest) (_result *ListTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTasksResponse{}
	_body, _err := client.ListTasksWithOptions(projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取工作流实例列表
//
// @param request - ListWorkflowInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkflowInstancesResponse
func (client *Client) ListWorkflowInstancesWithOptions(projectId *string, request *ListWorkflowInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListWorkflowInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.WorkflowId)) {
		query["workflowId"] = request.WorkflowId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWorkflowInstances"),
		Version:     tea.String("2024-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dolphinscheduler/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectId)) + "/workflowInstances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWorkflowInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工作流实例列表
//
// @param request - ListWorkflowInstancesRequest
//
// @return ListWorkflowInstancesResponse
func (client *Client) ListWorkflowInstances(projectId *string, request *ListWorkflowInstancesRequest) (_result *ListWorkflowInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListWorkflowInstancesResponse{}
	_body, _err := client.ListWorkflowInstancesWithOptions(projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取工作流列表
//
// @param request - ListWorkflowsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkflowsResponse
func (client *Client) ListWorkflowsWithOptions(projectId *string, request *ListWorkflowsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListWorkflowsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.SearchVal)) {
		query["searchVal"] = request.SearchVal
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWorkflows"),
		Version:     tea.String("2024-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dolphinscheduler/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectId)) + "/workflows"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWorkflowsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工作流列表
//
// @param request - ListWorkflowsRequest
//
// @return ListWorkflowsResponse
func (client *Client) ListWorkflows(projectId *string, request *ListWorkflowsRequest) (_result *ListWorkflowsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListWorkflowsResponse{}
	_body, _err := client.ListWorkflowsWithOptions(projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 运行工作流
//
// @param request - RunWorkflowRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunWorkflowResponse
func (client *Client) RunWorkflowWithOptions(projectId *string, request *RunWorkflowRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RunWorkflowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertGroupId)) {
		query["alertGroupId"] = request.AlertGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.AlertStrategy)) {
		query["alertStrategy"] = request.AlertStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.ComplementDependentMode)) {
		query["complementDependentMode"] = request.ComplementDependentMode
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ExecType)) {
		query["execType"] = request.ExecType
	}

	if !tea.BoolValue(util.IsUnset(request.ExpectedParallelismNumber)) {
		query["expectedParallelismNumber"] = request.ExpectedParallelismNumber
	}

	if !tea.BoolValue(util.IsUnset(request.FailureStrategy)) {
		query["failureStrategy"] = request.FailureStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RunMode)) {
		query["runMode"] = request.RunMode
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleTime)) {
		query["scheduleTime"] = request.ScheduleTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartParams)) {
		query["startParams"] = request.StartParams
	}

	if !tea.BoolValue(util.IsUnset(request.WorkflowId)) {
		query["workflowId"] = request.WorkflowId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkflowInstancePriority)) {
		query["workflowInstancePriority"] = request.WorkflowInstancePriority
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RunWorkflow"),
		Version:     tea.String("2024-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dolphinscheduler/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectId)) + "/executors/run-workflow"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RunWorkflowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 运行工作流
//
// @param request - RunWorkflowRequest
//
// @return RunWorkflowResponse
func (client *Client) RunWorkflow(projectId *string, request *RunWorkflowRequest) (_result *RunWorkflowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunWorkflowResponse{}
	_body, _err := client.RunWorkflowWithOptions(projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新工作流
//
// @param request - UpdateWorkflowRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkflowResponse
func (client *Client) UpdateWorkflowWithOptions(projectId *string, workflowId *string, request *UpdateWorkflowRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateWorkflowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertGroupId)) {
		query["alertGroupId"] = request.AlertGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.AlertStrategy)) {
		query["alertStrategy"] = request.AlertStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.CronExpr)) {
		query["cronExpr"] = request.CronExpr
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutionType)) {
		query["executionType"] = request.ExecutionType
	}

	if !tea.BoolValue(util.IsUnset(request.FailureStrategy)) {
		query["failureStrategy"] = request.FailureStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ParentDirectoryId)) {
		query["parentDirectoryId"] = request.ParentDirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleEndTime)) {
		query["scheduleEndTime"] = request.ScheduleEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleStartTime)) {
		query["scheduleStartTime"] = request.ScheduleStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleState)) {
		query["scheduleState"] = request.ScheduleState
	}

	if !tea.BoolValue(util.IsUnset(request.TaskDefinitionJson)) {
		query["taskDefinitionJson"] = request.TaskDefinitionJson
	}

	if !tea.BoolValue(util.IsUnset(request.TaskRelationJson)) {
		query["taskRelationJson"] = request.TaskRelationJson
	}

	if !tea.BoolValue(util.IsUnset(request.TimeZone)) {
		query["timeZone"] = request.TimeZone
	}

	if !tea.BoolValue(util.IsUnset(request.Timeout)) {
		query["timeout"] = request.Timeout
	}

	if !tea.BoolValue(util.IsUnset(request.WorkflowInstancePriority)) {
		query["workflowInstancePriority"] = request.WorkflowInstancePriority
	}

	if !tea.BoolValue(util.IsUnset(request.WorkflowParams)) {
		query["workflowParams"] = request.WorkflowParams
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWorkflow"),
		Version:     tea.String("2024-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dolphinscheduler/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectId)) + "/workflows/" + tea.StringValue(openapiutil.GetEncodeParam(workflowId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWorkflowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新工作流
//
// @param request - UpdateWorkflowRequest
//
// @return UpdateWorkflowResponse
func (client *Client) UpdateWorkflow(projectId *string, workflowId *string, request *UpdateWorkflowRequest) (_result *UpdateWorkflowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateWorkflowResponse{}
	_body, _err := client.UpdateWorkflowWithOptions(projectId, workflowId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
