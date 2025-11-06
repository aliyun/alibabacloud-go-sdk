// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListTaskInstancesResponseBodyPagingInfo) *ListTaskInstancesResponseBody
	GetPagingInfo() *ListTaskInstancesResponseBodyPagingInfo
	SetRequestId(v string) *ListTaskInstancesResponseBody
	GetRequestId() *string
}

type ListTaskInstancesResponseBody struct {
	// The pagination information.
	PagingInfo *ListTaskInstancesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTaskInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTaskInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTaskInstancesResponseBody) GetPagingInfo() *ListTaskInstancesResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListTaskInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTaskInstancesResponseBody) SetPagingInfo(v *ListTaskInstancesResponseBodyPagingInfo) *ListTaskInstancesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListTaskInstancesResponseBody) SetRequestId(v string) *ListTaskInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTaskInstancesResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTaskInstancesResponseBodyPagingInfo struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of task instances.
	TaskInstances []*ListTaskInstancesResponseBodyPagingInfoTaskInstances `json:"TaskInstances,omitempty" xml:"TaskInstances,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTaskInstancesResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListTaskInstancesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListTaskInstancesResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTaskInstancesResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTaskInstancesResponseBodyPagingInfo) GetTaskInstances() []*ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	return s.TaskInstances
}

func (s *ListTaskInstancesResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTaskInstancesResponseBodyPagingInfo) SetPageNumber(v int32) *ListTaskInstancesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfo) SetPageSize(v int32) *ListTaskInstancesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfo) SetTaskInstances(v []*ListTaskInstancesResponseBodyPagingInfoTaskInstances) *ListTaskInstancesResponseBodyPagingInfo {
	s.TaskInstances = v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfo) SetTotalCount(v int32) *ListTaskInstancesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfo) Validate() error {
	if s.TaskInstances != nil {
		for _, item := range s.TaskInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTaskInstancesResponseBodyPagingInfoTaskInstances struct {
	// The baseline ID.
	//
	// example:
	//
	// 1234
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The data timestamp.
	//
	// example:
	//
	// 1710239005403
	Bizdate *int64 `json:"Bizdate,omitempty" xml:"Bizdate,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1710239005403
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The account ID of the user who creates the instance.
	//
	// example:
	//
	// 1000
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The information about the associated data source.
	DataSource *ListTaskInstancesResponseBodyPagingInfoTaskInstancesDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the instance finished running.
	//
	// example:
	//
	// 1710239005403
	FinishedTime *int64 `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 1710239005403
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The account ID of the user who modifies the instance.
	//
	// example:
	//
	// 1000
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// The account ID of the task owner.
	//
	// example:
	//
	// 1000
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The sequence number of the cycle. This parameter indicates the cycle of the task instance on the current day.
	//
	// example:
	//
	// 1
	PeriodNumber *int32 `json:"PeriodNumber,omitempty" xml:"PeriodNumber,omitempty"`
	// The priority of the task. Minimum value: 1. Maximum value: 8. A larger value indicates a higher priority. Default value: 1.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The environment of the workspace.
	//
	// Valid values:
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
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The rerun mode
	//
	// Valid values:
	//
	// 	- AllDenied: The task cannot be rerun regardless of whether the task is successfully run or fails to run.
	//
	// 	- FailureAllowed: The task can be rerun only after it fails to run.
	//
	// 	- AllAllowed: The task can be rerun regardless of whether the task is successfully run or fails to run.
	//
	// example:
	//
	// AllAllowed
	RerunMode *string `json:"RerunMode,omitempty" xml:"RerunMode,omitempty"`
	// The number of times the task is run. By default, the value starts from 1.
	//
	// example:
	//
	// 1
	RunNumber *int32 `json:"RunNumber,omitempty" xml:"RunNumber,omitempty"`
	// The runtime information about the instance.
	Runtime *ListTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime `json:"Runtime,omitempty" xml:"Runtime,omitempty" type:"Struct"`
	// The information about the resource group with which the instance is associated.
	RuntimeResource *ListTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// The script parameter list.
	//
	// example:
	//
	// para1=val1 para2=val2
	ScriptParameters *string `json:"ScriptParameters,omitempty" xml:"ScriptParameters,omitempty"`
	// The time when the instance started to run.
	//
	// example:
	//
	// 1710239005403
	StartedTime *int64 `json:"StartedTime,omitempty" xml:"StartedTime,omitempty"`
	// The status of the instance.
	//
	// Valid values:
	//
	// 	- NotRun: The instance is not run.
	//
	// 	- Running: The instance is running.
	//
	// 	- WaitTime: The instance is waiting for the scheduling time to arrive.
	//
	// 	- CheckingCondition: Branch conditions are being checked for the instance.
	//
	// 	- WaitResource: The instance is waiting for resources.
	//
	// 	- Failure: The instance fails to be run.
	//
	// 	- Success: The instance is successfully run.
	//
	// 	- Checking: Data quality is being checked for the instance.
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
	// The name of the task for which the instance is generated.
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
	// The timeout period of task running. Unit: seconds.
	//
	// Note: The value of this parameter is rounded up by hour.
	//
	// example:
	//
	// 1
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The running mode of the instance after it is triggered. This parameter takes effect only if the TriggerType parameter is set to Scheduler.
	//
	// Valid values:
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
	// The scheduling time.
	//
	// example:
	//
	// 1710239005403
	TriggerTime *int64 `json:"TriggerTime,omitempty" xml:"TriggerTime,omitempty"`
	// The trigger type.
	//
	// Valid values:
	//
	// 	- Scheduler: scheduling cycle-based trigger
	//
	// 	- Manual: manual trigger
	//
	// example:
	//
	// Scheduler
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// The timestamp for when it started waiting for resources.
	//
	// example:
	//
	// 1710239005403
	WaitingResourceTime *int64 `json:"WaitingResourceTime,omitempty" xml:"WaitingResourceTime,omitempty"`
	// The timestamp for when it started waiting for the scheduled time.
	//
	// example:
	//
	// 1710239005403
	WaitingTriggerTime *int64 `json:"WaitingTriggerTime,omitempty" xml:"WaitingTriggerTime,omitempty"`
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
	// The type of the workflow instance.
	//
	// Valid values:
	//
	// 	- SmokeTest
	//
	// 	- SupplementData
	//
	// 	- Manual
	//
	// 	- ManualWorkflow
	//
	// 	- Normal
	//
	// 	- ManualFlow
	//
	// example:
	//
	// Normal
	WorkflowInstanceType *string `json:"WorkflowInstanceType,omitempty" xml:"WorkflowInstanceType,omitempty"`
	// The name of the workflow to which the instance belongs.
	//
	// example:
	//
	// Test workflow
	WorkflowName *string `json:"WorkflowName,omitempty" xml:"WorkflowName,omitempty"`
}

func (s ListTaskInstancesResponseBodyPagingInfoTaskInstances) String() string {
	return dara.Prettify(s)
}

func (s ListTaskInstancesResponseBodyPagingInfoTaskInstances) GoString() string {
	return s.String()
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) GetBizdate() *int64 {
	return s.Bizdate
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) GetDataSource() *ListTaskInstancesResponseBodyPagingInfoTaskInstancesDataSource {
	return s.DataSource
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) GetDescription() *string {
	return s.Description
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) GetFinishedTime() *int64 {
	return s.FinishedTime
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) GetId() *int64 {
	return s.Id
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) GetOwner() *string {
	return s.Owner
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) GetPeriodNumber() *int32 {
	return s.PeriodNumber
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) GetPriority() *int32 {
	return s.Priority
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) GetRerunMode() *string {
	return s.RerunMode
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) GetRunNumber() *int32 {
	return s.RunNumber
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) GetRuntime() *ListTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime {
	return s.Runtime
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) GetRuntimeResource() *ListTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource {
	return s.RuntimeResource
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) GetScriptParameters() *string {
	return s.ScriptParameters
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) GetStartedTime() *int64 {
	return s.StartedTime
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) GetStatus() *string {
	return s.Status
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) GetTaskName() *string {
	return s.TaskName
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) GetTaskType() *string {
	return s.TaskType
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) GetTimeout() *int32 {
	return s.Timeout
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) GetTriggerRecurrence() *string {
	return s.TriggerRecurrence
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) GetTriggerTime() *int64 {
	return s.TriggerTime
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) GetTriggerType() *string {
	return s.TriggerType
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) GetWaitingResourceTime() *int64 {
	return s.WaitingResourceTime
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) GetWaitingTriggerTime() *int64 {
	return s.WaitingTriggerTime
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) GetWorkflowInstanceId() *int64 {
	return s.WorkflowInstanceId
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) GetWorkflowInstanceType() *string {
	return s.WorkflowInstanceType
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) GetWorkflowName() *string {
	return s.WorkflowName
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) SetBaselineId(v int64) *ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.BaselineId = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) SetBizdate(v int64) *ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.Bizdate = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) SetCreateTime(v int64) *ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.CreateTime = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) SetCreateUser(v string) *ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.CreateUser = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) SetDataSource(v *ListTaskInstancesResponseBodyPagingInfoTaskInstancesDataSource) *ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.DataSource = v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) SetDescription(v string) *ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.Description = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) SetFinishedTime(v int64) *ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.FinishedTime = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) SetId(v int64) *ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.Id = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) SetModifyTime(v int64) *ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.ModifyTime = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) SetModifyUser(v string) *ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.ModifyUser = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) SetOwner(v string) *ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.Owner = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) SetPeriodNumber(v int32) *ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.PeriodNumber = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) SetPriority(v int32) *ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.Priority = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) SetProjectEnv(v string) *ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.ProjectEnv = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) SetProjectId(v int64) *ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.ProjectId = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) SetRerunMode(v string) *ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.RerunMode = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) SetRunNumber(v int32) *ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.RunNumber = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) SetRuntime(v *ListTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime) *ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.Runtime = v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) SetRuntimeResource(v *ListTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource) *ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.RuntimeResource = v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) SetScriptParameters(v string) *ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.ScriptParameters = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) SetStartedTime(v int64) *ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.StartedTime = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) SetStatus(v string) *ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.Status = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) SetTaskId(v int64) *ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.TaskId = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) SetTaskName(v string) *ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.TaskName = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) SetTaskType(v string) *ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.TaskType = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) SetTimeout(v int32) *ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.Timeout = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) SetTriggerRecurrence(v string) *ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.TriggerRecurrence = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) SetTriggerTime(v int64) *ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.TriggerTime = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) SetTriggerType(v string) *ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.TriggerType = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) SetWaitingResourceTime(v int64) *ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.WaitingResourceTime = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) SetWaitingTriggerTime(v int64) *ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.WaitingTriggerTime = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) SetWorkflowId(v int64) *ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.WorkflowId = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) SetWorkflowInstanceId(v int64) *ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.WorkflowInstanceId = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) SetWorkflowInstanceType(v string) *ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.WorkflowInstanceType = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) SetWorkflowName(v string) *ListTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.WorkflowName = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstances) Validate() error {
	if s.DataSource != nil {
		if err := s.DataSource.Validate(); err != nil {
			return err
		}
	}
	if s.Runtime != nil {
		if err := s.Runtime.Validate(); err != nil {
			return err
		}
	}
	if s.RuntimeResource != nil {
		if err := s.RuntimeResource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTaskInstancesResponseBodyPagingInfoTaskInstancesDataSource struct {
	// The name of the data source.
	//
	// example:
	//
	// mysql_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListTaskInstancesResponseBodyPagingInfoTaskInstancesDataSource) String() string {
	return dara.Prettify(s)
}

func (s ListTaskInstancesResponseBodyPagingInfoTaskInstancesDataSource) GoString() string {
	return s.String()
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstancesDataSource) GetName() *string {
	return s.Name
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstancesDataSource) SetName(v string) *ListTaskInstancesResponseBodyPagingInfoTaskInstancesDataSource {
	s.Name = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstancesDataSource) Validate() error {
	return dara.Validate(s)
}

type ListTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime struct {
	// The host for running.
	//
	// example:
	//
	// cn-shanghai.1.2
	Gateway *string `json:"Gateway,omitempty" xml:"Gateway,omitempty"`
	// The instance run ID.
	//
	// example:
	//
	// T3_123
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
}

func (s ListTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime) String() string {
	return dara.Prettify(s)
}

func (s ListTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime) GoString() string {
	return s.String()
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime) GetGateway() *string {
	return s.Gateway
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime) GetProcessId() *string {
	return s.ProcessId
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime) SetGateway(v string) *ListTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime {
	s.Gateway = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime) SetProcessId(v string) *ListTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime {
	s.ProcessId = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime) Validate() error {
	return dara.Validate(s)
}

type ListTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource struct {
	// The default number of CUs configured for task running.
	//
	// example:
	//
	// 0.25
	Cu *string `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The ID of the image configured for task running.
	//
	// example:
	//
	// i-xxxxxx
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The ID of the resource group for scheduling configured for task running.
	//
	// example:
	//
	// S_res_group_524258031846018_1684XXXXXXXXX
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s ListTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource) GoString() string {
	return s.String()
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource) GetCu() *string {
	return s.Cu
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource) GetImage() *string {
	return s.Image
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource) SetCu(v string) *ListTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource {
	s.Cu = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource) SetImage(v string) *ListTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource {
	s.Image = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource) SetResourceGroupId(v string) *ListTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource) Validate() error {
	return dara.Validate(s)
}
