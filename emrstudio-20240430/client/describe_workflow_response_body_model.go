// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWorkflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeWorkflowResponseBody
	GetRequestId() *string
	SetSchedule(v *DescribeWorkflowResponseBodySchedule) *DescribeWorkflowResponseBody
	GetSchedule() *DescribeWorkflowResponseBodySchedule
	SetTaskRelations(v []*DescribeWorkflowResponseBodyTaskRelations) *DescribeWorkflowResponseBody
	GetTaskRelations() []*DescribeWorkflowResponseBodyTaskRelations
	SetTasks(v []*DescribeWorkflowResponseBodyTasks) *DescribeWorkflowResponseBody
	GetTasks() []*DescribeWorkflowResponseBodyTasks
	SetWorkflow(v *DescribeWorkflowResponseBodyWorkflow) *DescribeWorkflowResponseBody
	GetWorkflow() *DescribeWorkflowResponseBodyWorkflow
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
	return dara.Prettify(s)
}

func (s DescribeWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWorkflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWorkflowResponseBody) GetSchedule() *DescribeWorkflowResponseBodySchedule {
	return s.Schedule
}

func (s *DescribeWorkflowResponseBody) GetTaskRelations() []*DescribeWorkflowResponseBodyTaskRelations {
	return s.TaskRelations
}

func (s *DescribeWorkflowResponseBody) GetTasks() []*DescribeWorkflowResponseBodyTasks {
	return s.Tasks
}

func (s *DescribeWorkflowResponseBody) GetWorkflow() *DescribeWorkflowResponseBodyWorkflow {
	return s.Workflow
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

func (s *DescribeWorkflowResponseBody) Validate() error {
	if s.Schedule != nil {
		if err := s.Schedule.Validate(); err != nil {
			return err
		}
	}
	if s.TaskRelations != nil {
		for _, item := range s.TaskRelations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tasks != nil {
		for _, item := range s.Tasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Workflow != nil {
		if err := s.Workflow.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	return dara.Prettify(s)
}

func (s DescribeWorkflowResponseBodySchedule) GoString() string {
	return s.String()
}

func (s *DescribeWorkflowResponseBodySchedule) GetAlertGroupId() *string {
	return s.AlertGroupId
}

func (s *DescribeWorkflowResponseBodySchedule) GetAlertStrategy() *string {
	return s.AlertStrategy
}

func (s *DescribeWorkflowResponseBodySchedule) GetCronExpr() *string {
	return s.CronExpr
}

func (s *DescribeWorkflowResponseBodySchedule) GetEmrClusterId() *string {
	return s.EmrClusterId
}

func (s *DescribeWorkflowResponseBodySchedule) GetFailureStrategy() *string {
	return s.FailureStrategy
}

func (s *DescribeWorkflowResponseBodySchedule) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeWorkflowResponseBodySchedule) GetScheduleEndTime() *string {
	return s.ScheduleEndTime
}

func (s *DescribeWorkflowResponseBodySchedule) GetScheduleStartTime() *string {
	return s.ScheduleStartTime
}

func (s *DescribeWorkflowResponseBodySchedule) GetScheduleState() *string {
	return s.ScheduleState
}

func (s *DescribeWorkflowResponseBodySchedule) GetTimeZone() *string {
	return s.TimeZone
}

func (s *DescribeWorkflowResponseBodySchedule) GetWorkflowInstancePriority() *string {
	return s.WorkflowInstancePriority
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

func (s *DescribeWorkflowResponseBodySchedule) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s DescribeWorkflowResponseBodyTaskRelations) GoString() string {
	return s.String()
}

func (s *DescribeWorkflowResponseBodyTaskRelations) GetPostTaskId() *string {
	return s.PostTaskId
}

func (s *DescribeWorkflowResponseBodyTaskRelations) GetPreTaskId() *string {
	return s.PreTaskId
}

func (s *DescribeWorkflowResponseBodyTaskRelations) SetPostTaskId(v string) *DescribeWorkflowResponseBodyTaskRelations {
	s.PostTaskId = &v
	return s
}

func (s *DescribeWorkflowResponseBodyTaskRelations) SetPreTaskId(v string) *DescribeWorkflowResponseBodyTaskRelations {
	s.PreTaskId = &v
	return s
}

func (s *DescribeWorkflowResponseBodyTaskRelations) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s DescribeWorkflowResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DescribeWorkflowResponseBodyTasks) GetDescription() *string {
	return s.Description
}

func (s *DescribeWorkflowResponseBodyTasks) GetName() *string {
	return s.Name
}

func (s *DescribeWorkflowResponseBodyTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeWorkflowResponseBodyTasks) GetVersion() *int32 {
	return s.Version
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

func (s *DescribeWorkflowResponseBodyTasks) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s DescribeWorkflowResponseBodyWorkflow) GoString() string {
	return s.String()
}

func (s *DescribeWorkflowResponseBodyWorkflow) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeWorkflowResponseBodyWorkflow) GetDescription() *string {
	return s.Description
}

func (s *DescribeWorkflowResponseBodyWorkflow) GetExecutionType() *string {
	return s.ExecutionType
}

func (s *DescribeWorkflowResponseBodyWorkflow) GetName() *string {
	return s.Name
}

func (s *DescribeWorkflowResponseBodyWorkflow) GetParentDirectoryId() *string {
	return s.ParentDirectoryId
}

func (s *DescribeWorkflowResponseBodyWorkflow) GetTimeout() *int32 {
	return s.Timeout
}

func (s *DescribeWorkflowResponseBodyWorkflow) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeWorkflowResponseBodyWorkflow) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *DescribeWorkflowResponseBodyWorkflow) GetWorkflowParams() *string {
	return s.WorkflowParams
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

func (s *DescribeWorkflowResponseBodyWorkflow) Validate() error {
	return dara.Validate(s)
}
