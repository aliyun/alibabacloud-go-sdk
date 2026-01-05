// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizdate(v int64) *ListTaskInstancesRequest
	GetBizdate() *int64
	SetFilter(v string) *ListTaskInstancesRequest
	GetFilter() *string
	SetId(v int64) *ListTaskInstancesRequest
	GetId() *int64
	SetIds(v []*int64) *ListTaskInstancesRequest
	GetIds() []*int64
	SetOwner(v string) *ListTaskInstancesRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListTaskInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTaskInstancesRequest
	GetPageSize() *int32
	SetProjectEnv(v string) *ListTaskInstancesRequest
	GetProjectEnv() *string
	SetProjectId(v int64) *ListTaskInstancesRequest
	GetProjectId() *int64
	SetRuntimeResource(v string) *ListTaskInstancesRequest
	GetRuntimeResource() *string
	SetSortBy(v string) *ListTaskInstancesRequest
	GetSortBy() *string
	SetStatus(v string) *ListTaskInstancesRequest
	GetStatus() *string
	SetTaskId(v int64) *ListTaskInstancesRequest
	GetTaskId() *int64
	SetTaskIds(v []*int64) *ListTaskInstancesRequest
	GetTaskIds() []*int64
	SetTaskName(v string) *ListTaskInstancesRequest
	GetTaskName() *string
	SetTaskType(v string) *ListTaskInstancesRequest
	GetTaskType() *string
	SetTriggerRecurrence(v string) *ListTaskInstancesRequest
	GetTriggerRecurrence() *string
	SetTriggerType(v string) *ListTaskInstancesRequest
	GetTriggerType() *string
	SetUnifiedWorkflowInstanceId(v int64) *ListTaskInstancesRequest
	GetUnifiedWorkflowInstanceId() *int64
	SetWorkflowId(v int64) *ListTaskInstancesRequest
	GetWorkflowId() *int64
	SetWorkflowInstanceId(v int64) *ListTaskInstancesRequest
	GetWorkflowInstanceId() *int64
	SetWorkflowInstanceType(v string) *ListTaskInstancesRequest
	GetWorkflowInstanceType() *string
}

type ListTaskInstancesRequest struct {
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
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
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
	TaskIds []*int64 `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
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

func (s ListTaskInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTaskInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListTaskInstancesRequest) GetBizdate() *int64 {
	return s.Bizdate
}

func (s *ListTaskInstancesRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListTaskInstancesRequest) GetId() *int64 {
	return s.Id
}

func (s *ListTaskInstancesRequest) GetIds() []*int64 {
	return s.Ids
}

func (s *ListTaskInstancesRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListTaskInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTaskInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTaskInstancesRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *ListTaskInstancesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListTaskInstancesRequest) GetRuntimeResource() *string {
	return s.RuntimeResource
}

func (s *ListTaskInstancesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListTaskInstancesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListTaskInstancesRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ListTaskInstancesRequest) GetTaskIds() []*int64 {
	return s.TaskIds
}

func (s *ListTaskInstancesRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *ListTaskInstancesRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *ListTaskInstancesRequest) GetTriggerRecurrence() *string {
	return s.TriggerRecurrence
}

func (s *ListTaskInstancesRequest) GetTriggerType() *string {
	return s.TriggerType
}

func (s *ListTaskInstancesRequest) GetUnifiedWorkflowInstanceId() *int64 {
	return s.UnifiedWorkflowInstanceId
}

func (s *ListTaskInstancesRequest) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *ListTaskInstancesRequest) GetWorkflowInstanceId() *int64 {
	return s.WorkflowInstanceId
}

func (s *ListTaskInstancesRequest) GetWorkflowInstanceType() *string {
	return s.WorkflowInstanceType
}

func (s *ListTaskInstancesRequest) SetBizdate(v int64) *ListTaskInstancesRequest {
	s.Bizdate = &v
	return s
}

func (s *ListTaskInstancesRequest) SetFilter(v string) *ListTaskInstancesRequest {
	s.Filter = &v
	return s
}

func (s *ListTaskInstancesRequest) SetId(v int64) *ListTaskInstancesRequest {
	s.Id = &v
	return s
}

func (s *ListTaskInstancesRequest) SetIds(v []*int64) *ListTaskInstancesRequest {
	s.Ids = v
	return s
}

func (s *ListTaskInstancesRequest) SetOwner(v string) *ListTaskInstancesRequest {
	s.Owner = &v
	return s
}

func (s *ListTaskInstancesRequest) SetPageNumber(v int32) *ListTaskInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTaskInstancesRequest) SetPageSize(v int32) *ListTaskInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTaskInstancesRequest) SetProjectEnv(v string) *ListTaskInstancesRequest {
	s.ProjectEnv = &v
	return s
}

func (s *ListTaskInstancesRequest) SetProjectId(v int64) *ListTaskInstancesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListTaskInstancesRequest) SetRuntimeResource(v string) *ListTaskInstancesRequest {
	s.RuntimeResource = &v
	return s
}

func (s *ListTaskInstancesRequest) SetSortBy(v string) *ListTaskInstancesRequest {
	s.SortBy = &v
	return s
}

func (s *ListTaskInstancesRequest) SetStatus(v string) *ListTaskInstancesRequest {
	s.Status = &v
	return s
}

func (s *ListTaskInstancesRequest) SetTaskId(v int64) *ListTaskInstancesRequest {
	s.TaskId = &v
	return s
}

func (s *ListTaskInstancesRequest) SetTaskIds(v []*int64) *ListTaskInstancesRequest {
	s.TaskIds = v
	return s
}

func (s *ListTaskInstancesRequest) SetTaskName(v string) *ListTaskInstancesRequest {
	s.TaskName = &v
	return s
}

func (s *ListTaskInstancesRequest) SetTaskType(v string) *ListTaskInstancesRequest {
	s.TaskType = &v
	return s
}

func (s *ListTaskInstancesRequest) SetTriggerRecurrence(v string) *ListTaskInstancesRequest {
	s.TriggerRecurrence = &v
	return s
}

func (s *ListTaskInstancesRequest) SetTriggerType(v string) *ListTaskInstancesRequest {
	s.TriggerType = &v
	return s
}

func (s *ListTaskInstancesRequest) SetUnifiedWorkflowInstanceId(v int64) *ListTaskInstancesRequest {
	s.UnifiedWorkflowInstanceId = &v
	return s
}

func (s *ListTaskInstancesRequest) SetWorkflowId(v int64) *ListTaskInstancesRequest {
	s.WorkflowId = &v
	return s
}

func (s *ListTaskInstancesRequest) SetWorkflowInstanceId(v int64) *ListTaskInstancesRequest {
	s.WorkflowInstanceId = &v
	return s
}

func (s *ListTaskInstancesRequest) SetWorkflowInstanceType(v string) *ListTaskInstancesRequest {
	s.WorkflowInstanceType = &v
	return s
}

func (s *ListTaskInstancesRequest) Validate() error {
	return dara.Validate(s)
}
