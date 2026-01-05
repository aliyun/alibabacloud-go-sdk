// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizdate(v int64) *ListTaskInstancesShrinkRequest
	GetBizdate() *int64
	SetFilter(v string) *ListTaskInstancesShrinkRequest
	GetFilter() *string
	SetId(v int64) *ListTaskInstancesShrinkRequest
	GetId() *int64
	SetIdsShrink(v string) *ListTaskInstancesShrinkRequest
	GetIdsShrink() *string
	SetOwner(v string) *ListTaskInstancesShrinkRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListTaskInstancesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTaskInstancesShrinkRequest
	GetPageSize() *int32
	SetProjectEnv(v string) *ListTaskInstancesShrinkRequest
	GetProjectEnv() *string
	SetProjectId(v int64) *ListTaskInstancesShrinkRequest
	GetProjectId() *int64
	SetRuntimeResource(v string) *ListTaskInstancesShrinkRequest
	GetRuntimeResource() *string
	SetSortBy(v string) *ListTaskInstancesShrinkRequest
	GetSortBy() *string
	SetStatus(v string) *ListTaskInstancesShrinkRequest
	GetStatus() *string
	SetTaskId(v int64) *ListTaskInstancesShrinkRequest
	GetTaskId() *int64
	SetTaskIdsShrink(v string) *ListTaskInstancesShrinkRequest
	GetTaskIdsShrink() *string
	SetTaskName(v string) *ListTaskInstancesShrinkRequest
	GetTaskName() *string
	SetTaskType(v string) *ListTaskInstancesShrinkRequest
	GetTaskType() *string
	SetTriggerRecurrence(v string) *ListTaskInstancesShrinkRequest
	GetTriggerRecurrence() *string
	SetTriggerType(v string) *ListTaskInstancesShrinkRequest
	GetTriggerType() *string
	SetUnifiedWorkflowInstanceId(v int64) *ListTaskInstancesShrinkRequest
	GetUnifiedWorkflowInstanceId() *int64
	SetWorkflowId(v int64) *ListTaskInstancesShrinkRequest
	GetWorkflowId() *int64
	SetWorkflowInstanceId(v int64) *ListTaskInstancesShrinkRequest
	GetWorkflowInstanceId() *int64
	SetWorkflowInstanceType(v string) *ListTaskInstancesShrinkRequest
	GetWorkflowInstanceType() *string
}

type ListTaskInstancesShrinkRequest struct {
	// The data timestamp. The value of this parameter is 00:00:00 of the day before the scheduling time of the instance. The value is a UNIX timestamp. Unit: milliseconds. Example: 1743350400000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1710239005403
	Bizdate *int64 `json:"Bizdate,omitempty" xml:"Bizdate,omitempty"`
	// example:
	//
	// {
	//
	//     "startedTimeStart": "1763481600000",
	//
	//     "startedTimeEnd": "1763481600000",
	//
	//     "finishedTimeStart": "1763481600000",
	//
	//     "finishedTimeEnd": "1763481600000",
	//
	//     "createTimeStart": "1763481600000",
	//
	//     "createTimeEnd": "1763481600000"
	//
	// }
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The ID of the instance. The instance may be rerun. If the instance is rerun and you configure this parameter, the system returns the historical information of the instance, including the rerun information. You can use the RunNumber parameter to distinguish each entry in the historical information.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The IDs of the instances. You can query multiple instances at a time by instance ID.
	IdsShrink *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// The account ID of the task owner.
	//
	// example:
	//
	// 1000
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The environment of the workspace. Valid values:
	//
	// 	- Prod: production environment
	//
	// 	- Dev: development environment
	//
	// example:
	//
	// Prod
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
	// The DataWorks workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The information about the resource group. Set this parameter to the identifier of a resource group for scheduling.
	//
	// example:
	//
	// S_res_group_524258031846018_1684XXXXXXXXX
	RuntimeResource *string `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty"`
	// The field used for sorting. Fields such as TriggerTime and StartedTime are supported. The value of this parameter is in the Sort field + Sort by (Desc/Asc) format. By default, results are sorted in ascending order. Valid values:
	//
	// 	- `TriggerTime (Desc/Asc)`
	//
	// 	- `StartedTime (Desc/Asc)`
	//
	// 	- `FinishedTime (Desc/Asc)`
	//
	// 	- `CreateTime (Desc/Asc)`
	//
	// 	- `Id (Desc/Asc)`
	//
	//     Default value: `Id Desc`.
	//
	// example:
	//
	// Id Desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The status of the task instance.
	//
	// 	- `NotRun`: Not started
	//
	// 	- `Running`
	//
	// 	- `Failure`
	//
	// 	- `Success`
	//
	// 	- `WaitTime`: Awaiting scheduled time
	//
	// 	- `WaitResource`: Awaiting resources
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the task for which the instance is generated.
	//
	// example:
	//
	// 1234
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The IDs of the tasks. You can query multiple instances at a time by task ID.
	TaskIdsShrink *string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty"`
	// The name of the task. Fuzzy match is supported.
	//
	// example:
	//
	// SQL node
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The type of the task for which the instance is generated.
	//
	// example:
	//
	// ODPS_SQL
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The running mode of the instance after it is triggered. This parameter takes effect only if the TriggerType parameter is set to Scheduler. Valid values:
	//
	// 	- Pause
	//
	// 	- Skip
	//
	// 	- Normal
	//
	// example:
	//
	// Normal
	TriggerRecurrence *string `json:"TriggerRecurrence,omitempty" xml:"TriggerRecurrence,omitempty"`
	// The trigger type. Valid values:
	//
	// 	- Scheduler: scheduling cycle-based trigger
	//
	// 	- Manual: manual trigger
	//
	// example:
	//
	// Normal
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// example:
	//
	// 1234
	UnifiedWorkflowInstanceId *int64 `json:"UnifiedWorkflowInstanceId,omitempty" xml:"UnifiedWorkflowInstanceId,omitempty"`
	// The ID of the workflow to which the instance belongs.
	//
	// example:
	//
	// 1234
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
	// The workflow instance ID.
	//
	// example:
	//
	// 1234
	WorkflowInstanceId *int64 `json:"WorkflowInstanceId,omitempty" xml:"WorkflowInstanceId,omitempty"`
	// The type of the workflow instance. Valid values:
	//
	// 	- SmokeTest: Testing
	//
	// 	- Manual: Manually triggered node
	//
	// 	- SupplementData: Data backfill
	//
	// 	- ManualWorkflow: Manually triggered workflow
	//
	// 	- Normal: Scheduled execution
	//
	// 	- TriggerWorkflow: Triggered Workflow
	//
	// example:
	//
	// Normal
	WorkflowInstanceType *string `json:"WorkflowInstanceType,omitempty" xml:"WorkflowInstanceType,omitempty"`
}

func (s ListTaskInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTaskInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTaskInstancesShrinkRequest) GetBizdate() *int64 {
	return s.Bizdate
}

func (s *ListTaskInstancesShrinkRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListTaskInstancesShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *ListTaskInstancesShrinkRequest) GetIdsShrink() *string {
	return s.IdsShrink
}

func (s *ListTaskInstancesShrinkRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListTaskInstancesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTaskInstancesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTaskInstancesShrinkRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *ListTaskInstancesShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListTaskInstancesShrinkRequest) GetRuntimeResource() *string {
	return s.RuntimeResource
}

func (s *ListTaskInstancesShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListTaskInstancesShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *ListTaskInstancesShrinkRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ListTaskInstancesShrinkRequest) GetTaskIdsShrink() *string {
	return s.TaskIdsShrink
}

func (s *ListTaskInstancesShrinkRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *ListTaskInstancesShrinkRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *ListTaskInstancesShrinkRequest) GetTriggerRecurrence() *string {
	return s.TriggerRecurrence
}

func (s *ListTaskInstancesShrinkRequest) GetTriggerType() *string {
	return s.TriggerType
}

func (s *ListTaskInstancesShrinkRequest) GetUnifiedWorkflowInstanceId() *int64 {
	return s.UnifiedWorkflowInstanceId
}

func (s *ListTaskInstancesShrinkRequest) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *ListTaskInstancesShrinkRequest) GetWorkflowInstanceId() *int64 {
	return s.WorkflowInstanceId
}

func (s *ListTaskInstancesShrinkRequest) GetWorkflowInstanceType() *string {
	return s.WorkflowInstanceType
}

func (s *ListTaskInstancesShrinkRequest) SetBizdate(v int64) *ListTaskInstancesShrinkRequest {
	s.Bizdate = &v
	return s
}

func (s *ListTaskInstancesShrinkRequest) SetFilter(v string) *ListTaskInstancesShrinkRequest {
	s.Filter = &v
	return s
}

func (s *ListTaskInstancesShrinkRequest) SetId(v int64) *ListTaskInstancesShrinkRequest {
	s.Id = &v
	return s
}

func (s *ListTaskInstancesShrinkRequest) SetIdsShrink(v string) *ListTaskInstancesShrinkRequest {
	s.IdsShrink = &v
	return s
}

func (s *ListTaskInstancesShrinkRequest) SetOwner(v string) *ListTaskInstancesShrinkRequest {
	s.Owner = &v
	return s
}

func (s *ListTaskInstancesShrinkRequest) SetPageNumber(v int32) *ListTaskInstancesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTaskInstancesShrinkRequest) SetPageSize(v int32) *ListTaskInstancesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListTaskInstancesShrinkRequest) SetProjectEnv(v string) *ListTaskInstancesShrinkRequest {
	s.ProjectEnv = &v
	return s
}

func (s *ListTaskInstancesShrinkRequest) SetProjectId(v int64) *ListTaskInstancesShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListTaskInstancesShrinkRequest) SetRuntimeResource(v string) *ListTaskInstancesShrinkRequest {
	s.RuntimeResource = &v
	return s
}

func (s *ListTaskInstancesShrinkRequest) SetSortBy(v string) *ListTaskInstancesShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListTaskInstancesShrinkRequest) SetStatus(v string) *ListTaskInstancesShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListTaskInstancesShrinkRequest) SetTaskId(v int64) *ListTaskInstancesShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *ListTaskInstancesShrinkRequest) SetTaskIdsShrink(v string) *ListTaskInstancesShrinkRequest {
	s.TaskIdsShrink = &v
	return s
}

func (s *ListTaskInstancesShrinkRequest) SetTaskName(v string) *ListTaskInstancesShrinkRequest {
	s.TaskName = &v
	return s
}

func (s *ListTaskInstancesShrinkRequest) SetTaskType(v string) *ListTaskInstancesShrinkRequest {
	s.TaskType = &v
	return s
}

func (s *ListTaskInstancesShrinkRequest) SetTriggerRecurrence(v string) *ListTaskInstancesShrinkRequest {
	s.TriggerRecurrence = &v
	return s
}

func (s *ListTaskInstancesShrinkRequest) SetTriggerType(v string) *ListTaskInstancesShrinkRequest {
	s.TriggerType = &v
	return s
}

func (s *ListTaskInstancesShrinkRequest) SetUnifiedWorkflowInstanceId(v int64) *ListTaskInstancesShrinkRequest {
	s.UnifiedWorkflowInstanceId = &v
	return s
}

func (s *ListTaskInstancesShrinkRequest) SetWorkflowId(v int64) *ListTaskInstancesShrinkRequest {
	s.WorkflowId = &v
	return s
}

func (s *ListTaskInstancesShrinkRequest) SetWorkflowInstanceId(v int64) *ListTaskInstancesShrinkRequest {
	s.WorkflowInstanceId = &v
	return s
}

func (s *ListTaskInstancesShrinkRequest) SetWorkflowInstanceType(v string) *ListTaskInstancesShrinkRequest {
	s.WorkflowInstanceType = &v
	return s
}

func (s *ListTaskInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
