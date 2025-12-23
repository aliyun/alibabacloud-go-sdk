// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWorkflowInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlertGroupId(v string) *DescribeWorkflowInstanceResponseBody
	GetAlertGroupId() *string
	SetAlertStrategy(v string) *DescribeWorkflowInstanceResponseBody
	GetAlertStrategy() *string
	SetEmrClusterId(v string) *DescribeWorkflowInstanceResponseBody
	GetEmrClusterId() *string
	SetEndTime(v string) *DescribeWorkflowInstanceResponseBody
	GetEndTime() *string
	SetFailureStrategy(v string) *DescribeWorkflowInstanceResponseBody
	GetFailureStrategy() *string
	SetIsComplementData(v bool) *DescribeWorkflowInstanceResponseBody
	GetIsComplementData() *bool
	SetName(v string) *DescribeWorkflowInstanceResponseBody
	GetName() *string
	SetRequestId(v string) *DescribeWorkflowInstanceResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeWorkflowInstanceResponseBody
	GetResourceGroupId() *string
	SetRestartTime(v string) *DescribeWorkflowInstanceResponseBody
	GetRestartTime() *string
	SetRunTimes(v int32) *DescribeWorkflowInstanceResponseBody
	GetRunTimes() *int32
	SetScheduleTime(v string) *DescribeWorkflowInstanceResponseBody
	GetScheduleTime() *string
	SetStartTime(v string) *DescribeWorkflowInstanceResponseBody
	GetStartTime() *string
	SetStatus(v string) *DescribeWorkflowInstanceResponseBody
	GetStatus() *string
	SetTimeout(v int32) *DescribeWorkflowInstanceResponseBody
	GetTimeout() *int32
	SetWorkflowId(v string) *DescribeWorkflowInstanceResponseBody
	GetWorkflowId() *string
	SetWorkflowInstanceId(v string) *DescribeWorkflowInstanceResponseBody
	GetWorkflowInstanceId() *string
	SetWorkflowInstancePriority(v string) *DescribeWorkflowInstanceResponseBody
	GetWorkflowInstancePriority() *string
	SetWorkflowVersion(v int32) *DescribeWorkflowInstanceResponseBody
	GetWorkflowVersion() *int32
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
	return dara.Prettify(s)
}

func (s DescribeWorkflowInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWorkflowInstanceResponseBody) GetAlertGroupId() *string {
	return s.AlertGroupId
}

func (s *DescribeWorkflowInstanceResponseBody) GetAlertStrategy() *string {
	return s.AlertStrategy
}

func (s *DescribeWorkflowInstanceResponseBody) GetEmrClusterId() *string {
	return s.EmrClusterId
}

func (s *DescribeWorkflowInstanceResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeWorkflowInstanceResponseBody) GetFailureStrategy() *string {
	return s.FailureStrategy
}

func (s *DescribeWorkflowInstanceResponseBody) GetIsComplementData() *bool {
	return s.IsComplementData
}

func (s *DescribeWorkflowInstanceResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeWorkflowInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWorkflowInstanceResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeWorkflowInstanceResponseBody) GetRestartTime() *string {
	return s.RestartTime
}

func (s *DescribeWorkflowInstanceResponseBody) GetRunTimes() *int32 {
	return s.RunTimes
}

func (s *DescribeWorkflowInstanceResponseBody) GetScheduleTime() *string {
	return s.ScheduleTime
}

func (s *DescribeWorkflowInstanceResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeWorkflowInstanceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeWorkflowInstanceResponseBody) GetTimeout() *int32 {
	return s.Timeout
}

func (s *DescribeWorkflowInstanceResponseBody) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *DescribeWorkflowInstanceResponseBody) GetWorkflowInstanceId() *string {
	return s.WorkflowInstanceId
}

func (s *DescribeWorkflowInstanceResponseBody) GetWorkflowInstancePriority() *string {
	return s.WorkflowInstancePriority
}

func (s *DescribeWorkflowInstanceResponseBody) GetWorkflowVersion() *int32 {
	return s.WorkflowVersion
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

func (s *DescribeWorkflowInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
