// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertGroupId(v string) *CreateWorkflowRequest
	GetAlertGroupId() *string
	SetAlertStrategy(v string) *CreateWorkflowRequest
	GetAlertStrategy() *string
	SetTaskDefinitionJsonValue(v string) *CreateWorkflowRequest
	GetTaskDefinitionJsonValue() *string
	SetTaskRelationJsonValue(v string) *CreateWorkflowRequest
	GetTaskRelationJsonValue() *string
	SetCronExpr(v string) *CreateWorkflowRequest
	GetCronExpr() *string
	SetDescription(v string) *CreateWorkflowRequest
	GetDescription() *string
	SetExecutionType(v string) *CreateWorkflowRequest
	GetExecutionType() *string
	SetFailureStrategy(v string) *CreateWorkflowRequest
	GetFailureStrategy() *string
	SetName(v string) *CreateWorkflowRequest
	GetName() *string
	SetParentDirectoryId(v string) *CreateWorkflowRequest
	GetParentDirectoryId() *string
	SetResourceGroupId(v string) *CreateWorkflowRequest
	GetResourceGroupId() *string
	SetScheduleEndTime(v string) *CreateWorkflowRequest
	GetScheduleEndTime() *string
	SetScheduleStartTime(v string) *CreateWorkflowRequest
	GetScheduleStartTime() *string
	SetScheduleState(v string) *CreateWorkflowRequest
	GetScheduleState() *string
	SetTaskDefinitionJson(v string) *CreateWorkflowRequest
	GetTaskDefinitionJson() *string
	SetTaskRelationJson(v string) *CreateWorkflowRequest
	GetTaskRelationJson() *string
	SetTimeZone(v string) *CreateWorkflowRequest
	GetTimeZone() *string
	SetTimeout(v int32) *CreateWorkflowRequest
	GetTimeout() *int32
	SetWorkflowInstancePriority(v string) *CreateWorkflowRequest
	GetWorkflowInstancePriority() *string
	SetWorkflowParams(v string) *CreateWorkflowRequest
	GetWorkflowParams() *string
	SetWorkspaceId(v string) *CreateWorkflowRequest
	GetWorkspaceId() *string
}

type CreateWorkflowRequest struct {
	// example:
	//
	// ag-v7n2gp3vv3j****
	AlertGroupId *string `json:"alertGroupId,omitempty" xml:"alertGroupId,omitempty"`
	// example:
	//
	// NONE
	AlertStrategy           *string `json:"alertStrategy,omitempty" xml:"alertStrategy,omitempty"`
	TaskDefinitionJsonValue *string `json:"taskDefinitionJsonValue,omitempty" xml:"taskDefinitionJsonValue,omitempty"`
	TaskRelationJsonValue   *string `json:"taskRelationJsonValue,omitempty" xml:"taskRelationJsonValue,omitempty"`
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
	return dara.Prettify(s)
}

func (s CreateWorkflowRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkflowRequest) GetAlertGroupId() *string {
	return s.AlertGroupId
}

func (s *CreateWorkflowRequest) GetAlertStrategy() *string {
	return s.AlertStrategy
}

func (s *CreateWorkflowRequest) GetTaskDefinitionJsonValue() *string {
	return s.TaskDefinitionJsonValue
}

func (s *CreateWorkflowRequest) GetTaskRelationJsonValue() *string {
	return s.TaskRelationJsonValue
}

func (s *CreateWorkflowRequest) GetCronExpr() *string {
	return s.CronExpr
}

func (s *CreateWorkflowRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateWorkflowRequest) GetExecutionType() *string {
	return s.ExecutionType
}

func (s *CreateWorkflowRequest) GetFailureStrategy() *string {
	return s.FailureStrategy
}

func (s *CreateWorkflowRequest) GetName() *string {
	return s.Name
}

func (s *CreateWorkflowRequest) GetParentDirectoryId() *string {
	return s.ParentDirectoryId
}

func (s *CreateWorkflowRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateWorkflowRequest) GetScheduleEndTime() *string {
	return s.ScheduleEndTime
}

func (s *CreateWorkflowRequest) GetScheduleStartTime() *string {
	return s.ScheduleStartTime
}

func (s *CreateWorkflowRequest) GetScheduleState() *string {
	return s.ScheduleState
}

func (s *CreateWorkflowRequest) GetTaskDefinitionJson() *string {
	return s.TaskDefinitionJson
}

func (s *CreateWorkflowRequest) GetTaskRelationJson() *string {
	return s.TaskRelationJson
}

func (s *CreateWorkflowRequest) GetTimeZone() *string {
	return s.TimeZone
}

func (s *CreateWorkflowRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *CreateWorkflowRequest) GetWorkflowInstancePriority() *string {
	return s.WorkflowInstancePriority
}

func (s *CreateWorkflowRequest) GetWorkflowParams() *string {
	return s.WorkflowParams
}

func (s *CreateWorkflowRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateWorkflowRequest) SetAlertGroupId(v string) *CreateWorkflowRequest {
	s.AlertGroupId = &v
	return s
}

func (s *CreateWorkflowRequest) SetAlertStrategy(v string) *CreateWorkflowRequest {
	s.AlertStrategy = &v
	return s
}

func (s *CreateWorkflowRequest) SetTaskDefinitionJsonValue(v string) *CreateWorkflowRequest {
	s.TaskDefinitionJsonValue = &v
	return s
}

func (s *CreateWorkflowRequest) SetTaskRelationJsonValue(v string) *CreateWorkflowRequest {
	s.TaskRelationJsonValue = &v
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

func (s *CreateWorkflowRequest) Validate() error {
	return dara.Validate(s)
}
