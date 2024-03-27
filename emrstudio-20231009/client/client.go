// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DescribeManualTaskRequest struct {
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
	// 代表创建时间的资源属性字段
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 代表资源一级ID的资源属性字段
	ManualTaskId *string `json:"ManualTaskId,omitempty" xml:"ManualTaskId,omitempty"`
	// 代表资源名称的资源属性字段
	ManualTaskName *string `json:"ManualTaskName,omitempty" xml:"ManualTaskName,omitempty"`
	// 目录ID
	ParentDirectoryId *string `json:"ParentDirectoryId,omitempty" xml:"ParentDirectoryId,omitempty"`
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 资源id列表
	ResourceIds *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// 任务参数
	TaskParams *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty"`
	// 任务类型
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// Id of the request
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
	// EMR集群ID
	EmrClusterId *string `json:"EmrClusterId,omitempty" xml:"EmrClusterId,omitempty"`
	// 结束时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 外部应用ID
	ExternalAppId *string `json:"ExternalAppId,omitempty" xml:"ExternalAppId,omitempty"`
	// 代表资源一级ID的资源属性字段
	ManualTaskInstanceId *string `json:"ManualTaskInstanceId,omitempty" xml:"ManualTaskInstanceId,omitempty"`
	// 代表资源名称的资源属性字段
	ManualTaskInstanceName *string `json:"ManualTaskInstanceName,omitempty" xml:"ManualTaskInstanceName,omitempty"`
	// 资源组ID
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 提交时间
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	// 任务参数
	TaskParams *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty"`
	// 任务类型
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// Id of the request
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
	Code        *int64  `json:"code,omitempty" xml:"code,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	RequestId   *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DescribeProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProjectResponseBody) SetCode(v int64) *DescribeProjectResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeProjectResponseBody) SetDescription(v string) *DescribeProjectResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeProjectResponseBody) SetName(v string) *DescribeProjectResponseBody {
	s.Name = &v
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
	// 代表创建时间的资源属性字段
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 延时执行时间
	DelayTime *int32 `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 失败重试间隔
	FailRetryInterval *int32 `json:"FailRetryInterval,omitempty" xml:"FailRetryInterval,omitempty"`
	// 失败重试次数
	FailRetryTimes *int32 `json:"FailRetryTimes,omitempty" xml:"FailRetryTimes,omitempty"`
	// 运行标志
	Flag *string `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 资源ID列表
	ResourceIds *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// 代表资源一级ID的资源属性字段
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// 代表资源名称的资源属性字段
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// 任务参数
	TaskParams *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty"`
	// 任务优先级
	TaskPriority *string `json:"TaskPriority,omitempty" xml:"TaskPriority,omitempty"`
	// 任务类型
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// 超时时长
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// 超时告警标志
	TimeoutFlag *string `json:"TimeoutFlag,omitempty" xml:"TimeoutFlag,omitempty"`
	// 超时告警标志
	TimeoutNotifyStrategy *string `json:"TimeoutNotifyStrategy,omitempty" xml:"TimeoutNotifyStrategy,omitempty"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 版本
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// Id of the request
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
	// 空跑标识
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// EMR集群ID
	EmrClusterId *string `json:"EmrClusterId,omitempty" xml:"EmrClusterId,omitempty"`
	// 结束时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 外部应用ID
	ExternalAppId *string `json:"ExternalAppId,omitempty" xml:"ExternalAppId,omitempty"`
	// 资源组ID
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// 重试次数
	RetryTimes *int32 `json:"RetryTimes,omitempty" xml:"RetryTimes,omitempty"`
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 提交时间
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// 任务实例ID
	TaskInstanceId *string `json:"TaskInstanceId,omitempty" xml:"TaskInstanceId,omitempty"`
	// 任务实例名称
	TaskInstanceName *string `json:"TaskInstanceName,omitempty" xml:"TaskInstanceName,omitempty"`
	// 任务参数
	TaskParams *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty"`
	// 任务类型
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// 任务版本
	TaskVersion *string `json:"TaskVersion,omitempty" xml:"TaskVersion,omitempty"`
	// 工作流实例ID
	WorkflowInstanceId *string `json:"WorkflowInstanceId,omitempty" xml:"WorkflowInstanceId,omitempty"`
	// Id of the request
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
	WorkspaceId *int64 `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s DescribeWorkflowRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWorkflowRequest) GoString() string {
	return s.String()
}

func (s *DescribeWorkflowRequest) SetWorkspaceId(v int64) *DescribeWorkflowRequest {
	s.WorkspaceId = &v
	return s
}

type DescribeWorkflowResponseBody struct {
	RequestId     *string                                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TaskRelations []*DescribeWorkflowResponseBodyTaskRelations `json:"taskRelations,omitempty" xml:"taskRelations,omitempty" type:"Repeated"`
	Tasks         []*DescribeWorkflowResponseBodyTasks         `json:"tasks,omitempty" xml:"tasks,omitempty" type:"Repeated"`
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

func (s *DescribeWorkflowResponseBody) SetTaskRelations(v []*DescribeWorkflowResponseBodyTaskRelations) *DescribeWorkflowResponseBody {
	s.TaskRelations = v
	return s
}

func (s *DescribeWorkflowResponseBody) SetTasks(v []*DescribeWorkflowResponseBodyTasks) *DescribeWorkflowResponseBody {
	s.Tasks = v
	return s
}

type DescribeWorkflowResponseBodyTaskRelations struct {
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
	PostTaskId      *int64  `json:"postTaskId,omitempty" xml:"postTaskId,omitempty"`
	PostTaskVersion *int32  `json:"postTaskVersion,omitempty" xml:"postTaskVersion,omitempty"`
	PreTaskId       *int64  `json:"preTaskId,omitempty" xml:"preTaskId,omitempty"`
	PreTaskVersion  *int32  `json:"preTaskVersion,omitempty" xml:"preTaskVersion,omitempty"`
}

func (s DescribeWorkflowResponseBodyTaskRelations) String() string {
	return tea.Prettify(s)
}

func (s DescribeWorkflowResponseBodyTaskRelations) GoString() string {
	return s.String()
}

func (s *DescribeWorkflowResponseBodyTaskRelations) SetName(v string) *DescribeWorkflowResponseBodyTaskRelations {
	s.Name = &v
	return s
}

func (s *DescribeWorkflowResponseBodyTaskRelations) SetPostTaskId(v int64) *DescribeWorkflowResponseBodyTaskRelations {
	s.PostTaskId = &v
	return s
}

func (s *DescribeWorkflowResponseBodyTaskRelations) SetPostTaskVersion(v int32) *DescribeWorkflowResponseBodyTaskRelations {
	s.PostTaskVersion = &v
	return s
}

func (s *DescribeWorkflowResponseBodyTaskRelations) SetPreTaskId(v int64) *DescribeWorkflowResponseBodyTaskRelations {
	s.PreTaskId = &v
	return s
}

func (s *DescribeWorkflowResponseBodyTaskRelations) SetPreTaskVersion(v int32) *DescribeWorkflowResponseBodyTaskRelations {
	s.PreTaskVersion = &v
	return s
}

type DescribeWorkflowResponseBodyTasks struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	TaskId      *int64  `json:"taskId,omitempty" xml:"taskId,omitempty"`
	Version     *int32  `json:"version,omitempty" xml:"version,omitempty"`
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

func (s *DescribeWorkflowResponseBodyTasks) SetTaskId(v int64) *DescribeWorkflowResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *DescribeWorkflowResponseBodyTasks) SetVersion(v int32) *DescribeWorkflowResponseBodyTasks {
	s.Version = &v
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
	WorkspaceId *int64 `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s DescribeWorkflowInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWorkflowInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeWorkflowInstanceRequest) SetWorkspaceId(v int64) *DescribeWorkflowInstanceRequest {
	s.WorkspaceId = &v
	return s
}

type DescribeWorkflowInstanceResponseBody struct {
	EmrClusterId     *string `json:"emrClusterId,omitempty" xml:"emrClusterId,omitempty"`
	EndDate          *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	IsComplementData *bool   `json:"isComplementData,omitempty" xml:"isComplementData,omitempty"`
	Name             *string `json:"name,omitempty" xml:"name,omitempty"`
	RequestId        *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ResourceGroupId  *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	RunTimes         *int32  `json:"runTimes,omitempty" xml:"runTimes,omitempty"`
	ScheduleTime     *string `json:"scheduleTime,omitempty" xml:"scheduleTime,omitempty"`
	StartDate        *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	State            *string `json:"state,omitempty" xml:"state,omitempty"`
	Timeout          *int32  `json:"timeout,omitempty" xml:"timeout,omitempty"`
	WorkflowId       *int64  `json:"workflowId,omitempty" xml:"workflowId,omitempty"`
	WorkflowVersion  *int32  `json:"workflowVersion,omitempty" xml:"workflowVersion,omitempty"`
}

func (s DescribeWorkflowInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWorkflowInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWorkflowInstanceResponseBody) SetEmrClusterId(v string) *DescribeWorkflowInstanceResponseBody {
	s.EmrClusterId = &v
	return s
}

func (s *DescribeWorkflowInstanceResponseBody) SetEndDate(v string) *DescribeWorkflowInstanceResponseBody {
	s.EndDate = &v
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

func (s *DescribeWorkflowInstanceResponseBody) SetRunTimes(v int32) *DescribeWorkflowInstanceResponseBody {
	s.RunTimes = &v
	return s
}

func (s *DescribeWorkflowInstanceResponseBody) SetScheduleTime(v string) *DescribeWorkflowInstanceResponseBody {
	s.ScheduleTime = &v
	return s
}

func (s *DescribeWorkflowInstanceResponseBody) SetStartDate(v string) *DescribeWorkflowInstanceResponseBody {
	s.StartDate = &v
	return s
}

func (s *DescribeWorkflowInstanceResponseBody) SetState(v string) *DescribeWorkflowInstanceResponseBody {
	s.State = &v
	return s
}

func (s *DescribeWorkflowInstanceResponseBody) SetTimeout(v int32) *DescribeWorkflowInstanceResponseBody {
	s.Timeout = &v
	return s
}

func (s *DescribeWorkflowInstanceResponseBody) SetWorkflowId(v int64) *DescribeWorkflowInstanceResponseBody {
	s.WorkflowId = &v
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
	EndTime         *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	ExecutionStatus *string `json:"executionStatus,omitempty" xml:"executionStatus,omitempty"`
	MaxResults      *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken       *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	SearchVal       *string `json:"searchVal,omitempty" xml:"searchVal,omitempty"`
	StartTime       *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	WorkspaceId     *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
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

func (s *ListManualTaskInstancesRequest) SetExecutionStatus(v string) *ListManualTaskInstancesRequest {
	s.ExecutionStatus = &v
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

func (s *ListManualTaskInstancesRequest) SetWorkspaceId(v string) *ListManualTaskInstancesRequest {
	s.WorkspaceId = &v
	return s
}

type ListManualTaskInstancesResponseBody struct {
	Data      []*ListManualTaskInstancesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	NextToken *string                                    `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TotalSize *int32  `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
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
	// EMR集群ID
	EmrClusterId *string `json:"EmrClusterId,omitempty" xml:"EmrClusterId,omitempty"`
	// 结束时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 外部应用ID
	ExternalAppId *string `json:"ExternalAppId,omitempty" xml:"ExternalAppId,omitempty"`
	// 代表资源一级ID的资源属性字段
	ManualTaskInstanceId *string `json:"ManualTaskInstanceId,omitempty" xml:"ManualTaskInstanceId,omitempty"`
	// 代表资源名称的资源属性字段
	ManualTaskInstanceName *string `json:"ManualTaskInstanceName,omitempty" xml:"ManualTaskInstanceName,omitempty"`
	// 资源组ID
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 提交时间
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	// 任务参数
	TaskParams *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty"`
	// 任务类型
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
	MaxResults  *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken   *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	SearchVal   *string `json:"searchVal,omitempty" xml:"searchVal,omitempty"`
	TaskType    *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
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
	Data      []*ListManualTasksResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	NextToken *string                            `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TotalSize *int32  `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
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
	// 代表创建时间的资源属性字段
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 代表资源一级ID的资源属性字段
	ManualTaskId *string `json:"ManualTaskId,omitempty" xml:"ManualTaskId,omitempty"`
	// 代表资源名称的资源属性字段
	ManualTaskName *string `json:"ManualTaskName,omitempty" xml:"ManualTaskName,omitempty"`
	// 目录ID
	ParentDirectoryId *string `json:"ParentDirectoryId,omitempty" xml:"ParentDirectoryId,omitempty"`
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 资源id列表
	ResourceIds *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// 任务参数
	TaskParams *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty"`
	// 任务类型
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// 更新时间
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
	MaxResults  *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken   *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	SearchVal   *string `json:"searchVal,omitempty" xml:"searchVal,omitempty"`
	WorkspaceId *int64  `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
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

func (s *ListProjectsRequest) SetWorkspaceId(v int64) *ListProjectsRequest {
	s.WorkspaceId = &v
	return s
}

type ListProjectsResponseBody struct {
	Data      []*ListProjectsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	NextToken *string                         `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RequestId *string                         `json:"requestId,omitempty" xml:"requestId,omitempty"`
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

type ListProjectsResponseBodyData struct {
	Code        *int64  `json:"code,omitempty" xml:"code,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	ProjectId   *int64  `json:"projectId,omitempty" xml:"projectId,omitempty"`
}

func (s ListProjectsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyData) SetCode(v int64) *ListProjectsResponseBodyData {
	s.Code = &v
	return s
}

func (s *ListProjectsResponseBodyData) SetDescription(v string) *ListProjectsResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListProjectsResponseBodyData) SetName(v string) *ListProjectsResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListProjectsResponseBodyData) SetProjectId(v int64) *ListProjectsResponseBodyData {
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
	EndTime            *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	ExecutionStatus    *string `json:"executionStatus,omitempty" xml:"executionStatus,omitempty"`
	MaxResults         *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken          *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	SearchVal          *string `json:"searchVal,omitempty" xml:"searchVal,omitempty"`
	StartTime          *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	WorkflowInstanceId *string `json:"workflowInstanceId,omitempty" xml:"workflowInstanceId,omitempty"`
	WorkspaceId        *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
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

func (s *ListTaskInstancesRequest) SetExecutionStatus(v string) *ListTaskInstancesRequest {
	s.ExecutionStatus = &v
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

func (s *ListTaskInstancesRequest) SetWorkflowInstanceId(v string) *ListTaskInstancesRequest {
	s.WorkflowInstanceId = &v
	return s
}

func (s *ListTaskInstancesRequest) SetWorkspaceId(v string) *ListTaskInstancesRequest {
	s.WorkspaceId = &v
	return s
}

type ListTaskInstancesResponseBody struct {
	Data      []*ListTaskInstancesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	NextToken *string                              `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TotalSize *int32  `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
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
	// 空跑标识
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// EMR集群ID
	EmrClusterId *string `json:"EmrClusterId,omitempty" xml:"EmrClusterId,omitempty"`
	// 结束时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 外部应用ID
	ExternalAppId *string `json:"ExternalAppId,omitempty" xml:"ExternalAppId,omitempty"`
	// 资源组ID
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// 重试次数
	RetryTimes *int32 `json:"RetryTimes,omitempty" xml:"RetryTimes,omitempty"`
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 提交时间
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// 任务实例ID
	TaskInstanceId *string `json:"TaskInstanceId,omitempty" xml:"TaskInstanceId,omitempty"`
	// 任务实例名称
	TaskInstanceName *string `json:"TaskInstanceName,omitempty" xml:"TaskInstanceName,omitempty"`
	// 任务参数
	TaskParams *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty"`
	// 任务类型
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// 任务版本
	TaskVersion *string `json:"TaskVersion,omitempty" xml:"TaskVersion,omitempty"`
	// 工作流实例ID
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
	MaxResults  *string `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken   *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	SearchVal   *string `json:"searchVal,omitempty" xml:"searchVal,omitempty"`
	TaskType    *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	WorkflowId  *string `json:"workflowId,omitempty" xml:"workflowId,omitempty"`
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
	Data      []*ListTasksResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	NextToken *string                      `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TotalSize *int32  `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
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
	// 代表创建时间的资源属性字段
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 延时执行时间
	DelayTime *int32 `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 失败重试间隔
	FailRetryInterval *int32 `json:"FailRetryInterval,omitempty" xml:"FailRetryInterval,omitempty"`
	// 失败重试次数
	FailRetryTimes *int32 `json:"FailRetryTimes,omitempty" xml:"FailRetryTimes,omitempty"`
	// 运行标志
	Flag *string `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 资源ID列表
	ResourceIds *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// 代表资源一级ID的资源属性字段
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// 代表资源名称的资源属性字段
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// 任务参数
	TaskParams *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty"`
	// 任务优先级
	TaskPriority *string `json:"TaskPriority,omitempty" xml:"TaskPriority,omitempty"`
	// 任务类型
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// 超时时长
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// 超时告警标志
	TimeoutFlag *string `json:"TimeoutFlag,omitempty" xml:"TimeoutFlag,omitempty"`
	// 超时告警标志
	TimeoutNotifyStrategy *string `json:"TimeoutNotifyStrategy,omitempty" xml:"TimeoutNotifyStrategy,omitempty"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 版本
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
	EndDate     *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	MaxResults  *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken   *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	StartDate   *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	WorkflowId  *int64  `json:"workflowId,omitempty" xml:"workflowId,omitempty"`
	WorkspaceId *int64  `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListWorkflowInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkflowInstancesRequest) SetEndDate(v string) *ListWorkflowInstancesRequest {
	s.EndDate = &v
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

func (s *ListWorkflowInstancesRequest) SetStartDate(v string) *ListWorkflowInstancesRequest {
	s.StartDate = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetWorkflowId(v int64) *ListWorkflowInstancesRequest {
	s.WorkflowId = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetWorkspaceId(v int64) *ListWorkflowInstancesRequest {
	s.WorkspaceId = &v
	return s
}

type ListWorkflowInstancesResponseBody struct {
	Data      []*ListWorkflowInstancesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	NextToken *string                                  `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TotalSize *int32                                   `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
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
	EndDate            *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	RunTimes           *string `json:"runTimes,omitempty" xml:"runTimes,omitempty"`
	ScheduleTime       *string `json:"scheduleTime,omitempty" xml:"scheduleTime,omitempty"`
	StartDate          *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	State              *string `json:"state,omitempty" xml:"state,omitempty"`
	WorkflowId         *int64  `json:"workflowId,omitempty" xml:"workflowId,omitempty"`
	WorkflowInstanceId *int64  `json:"workflowInstanceId,omitempty" xml:"workflowInstanceId,omitempty"`
	WorkflowVersion    *string `json:"workflowVersion,omitempty" xml:"workflowVersion,omitempty"`
}

func (s ListWorkflowInstancesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWorkflowInstancesResponseBodyData) SetEndDate(v string) *ListWorkflowInstancesResponseBodyData {
	s.EndDate = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyData) SetName(v string) *ListWorkflowInstancesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyData) SetRunTimes(v string) *ListWorkflowInstancesResponseBodyData {
	s.RunTimes = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyData) SetScheduleTime(v string) *ListWorkflowInstancesResponseBodyData {
	s.ScheduleTime = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyData) SetStartDate(v string) *ListWorkflowInstancesResponseBodyData {
	s.StartDate = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyData) SetState(v string) *ListWorkflowInstancesResponseBodyData {
	s.State = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyData) SetWorkflowId(v int64) *ListWorkflowInstancesResponseBodyData {
	s.WorkflowId = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyData) SetWorkflowInstanceId(v int64) *ListWorkflowInstancesResponseBodyData {
	s.WorkflowInstanceId = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyData) SetWorkflowVersion(v string) *ListWorkflowInstancesResponseBodyData {
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
	MaxResults  *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken   *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	SearchVal   *string `json:"searchVal,omitempty" xml:"searchVal,omitempty"`
	WorkspaceId *int64  `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
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

func (s *ListWorkflowsRequest) SetWorkspaceId(v int64) *ListWorkflowsRequest {
	s.WorkspaceId = &v
	return s
}

type ListWorkflowsResponseBody struct {
	Data      []*ListWorkflowsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	NextToken *string                          `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RequestId *string                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TotalSize *int32                           `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
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
	CreateTime   *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description  *string `json:"description,omitempty" xml:"description,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
	ReleaseState *string `json:"releaseState,omitempty" xml:"releaseState,omitempty"`
	UpdateTime   *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	WorkflowId   *int64  `json:"workflowId,omitempty" xml:"workflowId,omitempty"`
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

func (s *ListWorkflowsResponseBodyData) SetReleaseState(v string) *ListWorkflowsResponseBodyData {
	s.ReleaseState = &v
	return s
}

func (s *ListWorkflowsResponseBodyData) SetUpdateTime(v string) *ListWorkflowsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListWorkflowsResponseBodyData) SetWorkflowId(v int64) *ListWorkflowsResponseBodyData {
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
		Version:     tea.String("2023-10-09"),
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
		Version:     tea.String("2023-10-09"),
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

func (client *Client) DescribeProjectWithOptions(code *string, request *DescribeProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeProjectResponse, _err error) {
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
		Version:     tea.String("2023-10-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dolphinscheduler/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(code))),
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

func (client *Client) DescribeProject(code *string, request *DescribeProjectRequest) (_result *DescribeProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeProjectResponse{}
	_body, _err := client.DescribeProjectWithOptions(code, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

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
		Version:     tea.String("2023-10-09"),
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
		Version:     tea.String("2023-10-09"),
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
		Version:     tea.String("2023-10-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dolphinscheduler/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectId)) + "/processDefinitions/" + tea.StringValue(openapiutil.GetEncodeParam(workflowId))),
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
		Version:     tea.String("2023-10-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dolphinscheduler/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectId)) + "/processInstances/" + tea.StringValue(openapiutil.GetEncodeParam(workflowInstanceId))),
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

func (client *Client) ListManualTaskInstancesWithOptions(projectId *string, request *ListManualTaskInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListManualTaskInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutionStatus)) {
		query["executionStatus"] = request.ExecutionStatus
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

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListManualTaskInstances"),
		Version:     tea.String("2023-10-09"),
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
		Version:     tea.String("2023-10-09"),
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
		Version:     tea.String("2023-10-09"),
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

func (client *Client) ListTaskInstancesWithOptions(projectId *string, request *ListTaskInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTaskInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutionStatus)) {
		query["executionStatus"] = request.ExecutionStatus
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
		Version:     tea.String("2023-10-09"),
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
		Version:     tea.String("2023-10-09"),
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

func (client *Client) ListWorkflowInstancesWithOptions(projectId *string, request *ListWorkflowInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListWorkflowInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["endDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["startDate"] = request.StartDate
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
		Version:     tea.String("2023-10-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dolphinscheduler/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectId)) + "/processInstances"),
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
		Version:     tea.String("2023-10-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dolphinscheduler/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectId)) + "/processDefinitions"),
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
