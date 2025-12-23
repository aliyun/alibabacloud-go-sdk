// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertGroupId(v string) *UpdateWorkflowRequest
	GetAlertGroupId() *string
	SetAlertStrategy(v string) *UpdateWorkflowRequest
	GetAlertStrategy() *string
	SetTaskDefinitionJsonValue(v string) *UpdateWorkflowRequest
	GetTaskDefinitionJsonValue() *string
	SetTaskRelationJsonValue(v string) *UpdateWorkflowRequest
	GetTaskRelationJsonValue() *string
	SetCronExpr(v string) *UpdateWorkflowRequest
	GetCronExpr() *string
	SetDescription(v string) *UpdateWorkflowRequest
	GetDescription() *string
	SetExecutionType(v string) *UpdateWorkflowRequest
	GetExecutionType() *string
	SetFailureStrategy(v string) *UpdateWorkflowRequest
	GetFailureStrategy() *string
	SetName(v string) *UpdateWorkflowRequest
	GetName() *string
	SetParentDirectoryId(v string) *UpdateWorkflowRequest
	GetParentDirectoryId() *string
	SetResourceGroupId(v string) *UpdateWorkflowRequest
	GetResourceGroupId() *string
	SetScheduleEndTime(v string) *UpdateWorkflowRequest
	GetScheduleEndTime() *string
	SetScheduleStartTime(v string) *UpdateWorkflowRequest
	GetScheduleStartTime() *string
	SetScheduleState(v string) *UpdateWorkflowRequest
	GetScheduleState() *string
	SetTaskDefinitionJson(v string) *UpdateWorkflowRequest
	GetTaskDefinitionJson() *string
	SetTaskRelationJson(v string) *UpdateWorkflowRequest
	GetTaskRelationJson() *string
	SetTimeZone(v string) *UpdateWorkflowRequest
	GetTimeZone() *string
	SetTimeout(v int32) *UpdateWorkflowRequest
	GetTimeout() *int32
	SetWorkflowInstancePriority(v string) *UpdateWorkflowRequest
	GetWorkflowInstancePriority() *string
	SetWorkflowParams(v string) *UpdateWorkflowRequest
	GetWorkflowParams() *string
	SetWorkspaceId(v string) *UpdateWorkflowRequest
	GetWorkspaceId() *string
}

type UpdateWorkflowRequest struct {
	// example:
	//
	// ag-n72kong0832****
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
	return dara.Prettify(s)
}

func (s UpdateWorkflowRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowRequest) GetAlertGroupId() *string {
	return s.AlertGroupId
}

func (s *UpdateWorkflowRequest) GetAlertStrategy() *string {
	return s.AlertStrategy
}

func (s *UpdateWorkflowRequest) GetTaskDefinitionJsonValue() *string {
	return s.TaskDefinitionJsonValue
}

func (s *UpdateWorkflowRequest) GetTaskRelationJsonValue() *string {
	return s.TaskRelationJsonValue
}

func (s *UpdateWorkflowRequest) GetCronExpr() *string {
	return s.CronExpr
}

func (s *UpdateWorkflowRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateWorkflowRequest) GetExecutionType() *string {
	return s.ExecutionType
}

func (s *UpdateWorkflowRequest) GetFailureStrategy() *string {
	return s.FailureStrategy
}

func (s *UpdateWorkflowRequest) GetName() *string {
	return s.Name
}

func (s *UpdateWorkflowRequest) GetParentDirectoryId() *string {
	return s.ParentDirectoryId
}

func (s *UpdateWorkflowRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateWorkflowRequest) GetScheduleEndTime() *string {
	return s.ScheduleEndTime
}

func (s *UpdateWorkflowRequest) GetScheduleStartTime() *string {
	return s.ScheduleStartTime
}

func (s *UpdateWorkflowRequest) GetScheduleState() *string {
	return s.ScheduleState
}

func (s *UpdateWorkflowRequest) GetTaskDefinitionJson() *string {
	return s.TaskDefinitionJson
}

func (s *UpdateWorkflowRequest) GetTaskRelationJson() *string {
	return s.TaskRelationJson
}

func (s *UpdateWorkflowRequest) GetTimeZone() *string {
	return s.TimeZone
}

func (s *UpdateWorkflowRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateWorkflowRequest) GetWorkflowInstancePriority() *string {
	return s.WorkflowInstancePriority
}

func (s *UpdateWorkflowRequest) GetWorkflowParams() *string {
	return s.WorkflowParams
}

func (s *UpdateWorkflowRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateWorkflowRequest) SetAlertGroupId(v string) *UpdateWorkflowRequest {
	s.AlertGroupId = &v
	return s
}

func (s *UpdateWorkflowRequest) SetAlertStrategy(v string) *UpdateWorkflowRequest {
	s.AlertStrategy = &v
	return s
}

func (s *UpdateWorkflowRequest) SetTaskDefinitionJsonValue(v string) *UpdateWorkflowRequest {
	s.TaskDefinitionJsonValue = &v
	return s
}

func (s *UpdateWorkflowRequest) SetTaskRelationJsonValue(v string) *UpdateWorkflowRequest {
	s.TaskRelationJsonValue = &v
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

func (s *UpdateWorkflowRequest) Validate() error {
	return dara.Validate(s)
}
