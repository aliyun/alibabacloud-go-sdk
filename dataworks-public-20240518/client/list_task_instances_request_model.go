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
	// The business date. Typically, this is 00:00:00 of the day before the scheduled time of the scheduled instance, in the format of a millisecond-level timestamp, such as 1743350400000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1710239005403
	Bizdate *int64 `json:"Bizdate,omitempty" xml:"Bizdate,omitempty"`
	// The filter conditions. The value must be in JSON format. Multiple filter conditions are connected by the logical AND operator. Supported fields: `startedTimeStart`, `startedTimeEnd`, `finishedTimeStart`, `finishedTimeEnd`, `createTimeStart`, and `createTimeEnd`.
	//
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
	// The task instance ID. Task instances may be rerun. If you specify this parameter, the response includes historical information about reruns. Each piece of historical information is distinguished by RunNumber.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The list of task instance IDs. You can use this parameter to query multiple task instances in batches based on task instance IDs.
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	// The account ID of the task owner.
	//
	// example:
	//
	// 1000
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 500.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The project environment.
	//
	// - Prod: Production.
	//
	// - Dev: Development.
	//
	// example:
	//
	// Prod
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The resource group information. Specify the identifier of the scheduling resource group.
	//
	// example:
	//
	// S_res_group_524258031846018_1684XXXXXXXXX
	RuntimeResource *string `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty"`
	// The sort field. Fields such as scheduled time and start time are supported. The format is "Sort field + Sort order (Desc/Asc)" (Asc can be omitted). The valid values for the sort field are as follows:
	//
	// - `TriggerTime (Desc/Asc)`
	//
	// - `StartedTime (Desc/Asc)`
	//
	// - `FinishedTime (Desc/Asc)`
	//
	// - `CreateTime (Desc/Asc)`
	//
	// - `Id (Desc/Asc)`
	//
	//   Default value: `Id Desc`.
	//
	// example:
	//
	// Id Desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The status of the task instance. Valid values:
	//
	// - `NotRun`: Not run.
	//
	// - `Running`: Running.
	//
	// - `Failure`: Failed.
	//
	// - `Success`: Successful.
	//
	// - `WaitTime`: Waiting for scheduled time.
	//
	// - `WaitResource`: Waiting for resources.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 1234
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The list of task IDs. You can use this parameter to query task instances in batches based on task IDs.
	TaskIds []*int64 `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	// The name of the task. Fuzzy match is supported.
	//
	// example:
	//
	// SQL node
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The task type. For the valid values of TaskType for different nodes, see [DataWorks nodes](https://help.aliyun.com/document_detail/600169.html).
	//
	// example:
	//
	// ODPS_SQL
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The running mode upon triggering. This parameter takes effect only when TriggerType is set to Scheduler.
	//
	// - Pause: Paused.
	//
	// - Skip: Dry run.
	//
	// - Normal: Normal execution.
	//
	// example:
	//
	// Normal
	TriggerRecurrence *string `json:"TriggerRecurrence,omitempty" xml:"TriggerRecurrence,omitempty"`
	// The trigger type.
	//
	// - Scheduler: Triggered by periodic scheduling.
	//
	// - Manual: Triggered manually.
	//
	// example:
	//
	// Scheduler
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// The unified workflow instance ID. The value of this field is the same for all task instances within a specific business date under a single trigger.
	//
	// example:
	//
	// 1234
	UnifiedWorkflowInstanceId *int64 `json:"UnifiedWorkflowInstanceId,omitempty" xml:"UnifiedWorkflowInstanceId,omitempty"`
	// The ID of the workflow to which the task belongs.
	//
	// example:
	//
	// 1234
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
	// The ID of the workflow instance.
	//
	// example:
	//
	// 1234
	WorkflowInstanceId *int64 `json:"WorkflowInstanceId,omitempty" xml:"WorkflowInstanceId,omitempty"`
	// The type of the workflow instance.
	//
	// - SmokeTest: Test.
	//
	// - Manual: Manual task.
	//
	// - SupplementData: Data backfill.
	//
	// - ManualWorkflow: Manual workflow.
	//
	// - Normal: Periodic scheduling.
	//
	// - TriggerWorkflow: Triggered workflow.
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
