// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunWorkflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertGroupId(v string) *RunWorkflowRequest
	GetAlertGroupId() *string
	SetAlertStrategy(v string) *RunWorkflowRequest
	GetAlertStrategy() *string
	SetComplementDependentMode(v string) *RunWorkflowRequest
	GetComplementDependentMode() *string
	SetDryRun(v string) *RunWorkflowRequest
	GetDryRun() *string
	SetExecType(v string) *RunWorkflowRequest
	GetExecType() *string
	SetExpectedParallelismNumber(v string) *RunWorkflowRequest
	GetExpectedParallelismNumber() *string
	SetFailureStrategy(v string) *RunWorkflowRequest
	GetFailureStrategy() *string
	SetResourceGroupId(v string) *RunWorkflowRequest
	GetResourceGroupId() *string
	SetRunMode(v string) *RunWorkflowRequest
	GetRunMode() *string
	SetScheduleTime(v string) *RunWorkflowRequest
	GetScheduleTime() *string
	SetStartParams(v string) *RunWorkflowRequest
	GetStartParams() *string
	SetWorkflowId(v string) *RunWorkflowRequest
	GetWorkflowId() *string
	SetWorkflowInstancePriority(v string) *RunWorkflowRequest
	GetWorkflowInstancePriority() *string
	SetWorkspaceId(v string) *RunWorkflowRequest
	GetWorkspaceId() *string
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
	return dara.Prettify(s)
}

func (s RunWorkflowRequest) GoString() string {
	return s.String()
}

func (s *RunWorkflowRequest) GetAlertGroupId() *string {
	return s.AlertGroupId
}

func (s *RunWorkflowRequest) GetAlertStrategy() *string {
	return s.AlertStrategy
}

func (s *RunWorkflowRequest) GetComplementDependentMode() *string {
	return s.ComplementDependentMode
}

func (s *RunWorkflowRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *RunWorkflowRequest) GetExecType() *string {
	return s.ExecType
}

func (s *RunWorkflowRequest) GetExpectedParallelismNumber() *string {
	return s.ExpectedParallelismNumber
}

func (s *RunWorkflowRequest) GetFailureStrategy() *string {
	return s.FailureStrategy
}

func (s *RunWorkflowRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *RunWorkflowRequest) GetRunMode() *string {
	return s.RunMode
}

func (s *RunWorkflowRequest) GetScheduleTime() *string {
	return s.ScheduleTime
}

func (s *RunWorkflowRequest) GetStartParams() *string {
	return s.StartParams
}

func (s *RunWorkflowRequest) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *RunWorkflowRequest) GetWorkflowInstancePriority() *string {
	return s.WorkflowInstancePriority
}

func (s *RunWorkflowRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
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

func (s *RunWorkflowRequest) Validate() error {
	return dara.Validate(s)
}
