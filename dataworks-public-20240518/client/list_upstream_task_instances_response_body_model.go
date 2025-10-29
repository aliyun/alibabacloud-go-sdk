// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUpstreamTaskInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListUpstreamTaskInstancesResponseBodyPagingInfo) *ListUpstreamTaskInstancesResponseBody
	GetPagingInfo() *ListUpstreamTaskInstancesResponseBodyPagingInfo
	SetRequestId(v string) *ListUpstreamTaskInstancesResponseBody
	GetRequestId() *string
}

type ListUpstreamTaskInstancesResponseBody struct {
	// The pagination information.
	PagingInfo *ListUpstreamTaskInstancesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUpstreamTaskInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUpstreamTaskInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUpstreamTaskInstancesResponseBody) GetPagingInfo() *ListUpstreamTaskInstancesResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListUpstreamTaskInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUpstreamTaskInstancesResponseBody) SetPagingInfo(v *ListUpstreamTaskInstancesResponseBodyPagingInfo) *ListUpstreamTaskInstancesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBody) SetRequestId(v string) *ListUpstreamTaskInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListUpstreamTaskInstancesResponseBodyPagingInfo struct {
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
	// The instances. This parameter is deprecated and replaced by the UpstreamTaskInstances parameter.
	TaskInstances []*ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances `json:"TaskInstances,omitempty" xml:"TaskInstances,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The ancestor instances.
	UpstreamTaskInstances []*ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstances `json:"UpstreamTaskInstances,omitempty" xml:"UpstreamTaskInstances,omitempty" type:"Repeated"`
}

func (s ListUpstreamTaskInstancesResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListUpstreamTaskInstancesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfo) GetTaskInstances() []*ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	return s.TaskInstances
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfo) GetUpstreamTaskInstances() []*ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstances {
	return s.UpstreamTaskInstances
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfo) SetPageNumber(v int32) *ListUpstreamTaskInstancesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfo) SetPageSize(v int32) *ListUpstreamTaskInstancesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfo) SetTaskInstances(v []*ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) *ListUpstreamTaskInstancesResponseBodyPagingInfo {
	s.TaskInstances = v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfo) SetTotalCount(v int32) *ListUpstreamTaskInstancesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfo) SetUpstreamTaskInstances(v []*ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstances) *ListUpstreamTaskInstancesResponseBodyPagingInfo {
	s.UpstreamTaskInstances = v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfo) Validate() error {
	if s.TaskInstances != nil {
		for _, item := range s.TaskInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UpstreamTaskInstances != nil {
		for _, item := range s.UpstreamTaskInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances struct {
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
	// The account ID of the creator.
	//
	// example:
	//
	// 1000
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The information about the associated data source.
	DataSource *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The environment of the workspace. Valid values:
	//
	// 	- Prod
	//
	// 	- Dev
	//
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
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
	// The account ID of the modifier.
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
	// The sequence number of the period. Indicates which cycle of the day the task instance is in.
	//
	// example:
	//
	// 1
	PeriodNumber *int32 `json:"PeriodNumber,omitempty" xml:"PeriodNumber,omitempty"`
	// The priority of the task. Valid values: 1 to 8. A larger value indicates a higher priority. Default value: 1.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// Deprecated
	//
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
	// The workspace ID.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The rerun mode. Valid values:
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
	// The number of times the instance is run. By default, the value starts from 1.
	//
	// example:
	//
	// 1
	RunNumber *int32 `json:"RunNumber,omitempty" xml:"RunNumber,omitempty"`
	// The runtime information about the instance.
	Runtime *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime `json:"Runtime,omitempty" xml:"Runtime,omitempty" type:"Struct"`
	// The configurations of the runtime environment, such as the resource group information.
	RuntimeResource *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// The time when the instance started to run.
	//
	// example:
	//
	// 1710239005403
	StartedTime *int64 `json:"StartedTime,omitempty" xml:"StartedTime,omitempty"`
	// The status of the instance. Valid values:
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
	// The scheduling dependency type. Valid values:
	//
	// 	- Normal: same-cycle scheduling dependency
	//
	// 	- CrossCycle: cross-cycle scheduling dependency
	//
	// example:
	//
	// Normal
	StepType *string `json:"StepType,omitempty" xml:"StepType,omitempty"`
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
	// example:
	//
	// 1
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
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
	// The scheduling time.
	//
	// example:
	//
	// 1710239005403
	TriggerTime *int64 `json:"TriggerTime,omitempty" xml:"TriggerTime,omitempty"`
	// The trigger type. Valid values:
	//
	// 	- Scheduler: scheduling cycle-based trigger
	//
	// 	- Manual: manual trigger
	//
	// example:
	//
	// Scheduler
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
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

func (s ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) String() string {
	return dara.Prettify(s)
}

func (s ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GoString() string {
	return s.String()
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetBizdate() *int64 {
	return s.Bizdate
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetDataSource() *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesDataSource {
	return s.DataSource
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetDescription() *string {
	return s.Description
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetEnvType() *string {
	return s.EnvType
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetFinishedTime() *int64 {
	return s.FinishedTime
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetId() *int64 {
	return s.Id
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetOwner() *string {
	return s.Owner
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetPeriodNumber() *int32 {
	return s.PeriodNumber
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetPriority() *int32 {
	return s.Priority
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetRerunMode() *string {
	return s.RerunMode
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetRunNumber() *int32 {
	return s.RunNumber
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetRuntime() *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime {
	return s.Runtime
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetRuntimeResource() *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource {
	return s.RuntimeResource
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetStartedTime() *int64 {
	return s.StartedTime
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetStatus() *string {
	return s.Status
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetStepType() *string {
	return s.StepType
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetTaskName() *string {
	return s.TaskName
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetTaskType() *string {
	return s.TaskType
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetTimeout() *int32 {
	return s.Timeout
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetTriggerRecurrence() *string {
	return s.TriggerRecurrence
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetTriggerTime() *int64 {
	return s.TriggerTime
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetTriggerType() *string {
	return s.TriggerType
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetWorkflowInstanceId() *int64 {
	return s.WorkflowInstanceId
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetWorkflowInstanceType() *string {
	return s.WorkflowInstanceType
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetWorkflowName() *string {
	return s.WorkflowName
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetBaselineId(v int64) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.BaselineId = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetBizdate(v int64) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.Bizdate = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetCreateTime(v int64) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.CreateTime = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetCreateUser(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.CreateUser = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetDataSource(v *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesDataSource) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.DataSource = v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetDescription(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.Description = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetEnvType(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.EnvType = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetFinishedTime(v int64) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.FinishedTime = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetId(v int64) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.Id = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetModifyTime(v int64) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.ModifyTime = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetModifyUser(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.ModifyUser = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetOwner(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.Owner = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetPeriodNumber(v int32) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.PeriodNumber = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetPriority(v int32) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.Priority = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetProjectEnv(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.ProjectEnv = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetProjectId(v int64) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.ProjectId = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetRerunMode(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.RerunMode = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetRunNumber(v int32) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.RunNumber = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetRuntime(v *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.Runtime = v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetRuntimeResource(v *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.RuntimeResource = v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetStartedTime(v int64) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.StartedTime = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetStatus(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.Status = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetStepType(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.StepType = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetTaskId(v int64) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.TaskId = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetTaskName(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.TaskName = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetTaskType(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.TaskType = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetTimeout(v int32) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.Timeout = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetTriggerRecurrence(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.TriggerRecurrence = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetTriggerTime(v int64) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.TriggerTime = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetTriggerType(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.TriggerType = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetWorkflowId(v int64) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.WorkflowId = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetWorkflowInstanceId(v int64) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.WorkflowInstanceId = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetWorkflowInstanceType(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.WorkflowInstanceType = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetWorkflowName(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.WorkflowName = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstances) Validate() error {
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

type ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesDataSource struct {
	// The name of the data source.
	//
	// example:
	//
	// mysql_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesDataSource) String() string {
	return dara.Prettify(s)
}

func (s ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesDataSource) GoString() string {
	return s.String()
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesDataSource) GetName() *string {
	return s.Name
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesDataSource) SetName(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesDataSource {
	s.Name = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesDataSource) Validate() error {
	return dara.Validate(s)
}

type ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime struct {
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

func (s ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime) String() string {
	return dara.Prettify(s)
}

func (s ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime) GoString() string {
	return s.String()
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime) GetGateway() *string {
	return s.Gateway
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime) GetProcessId() *string {
	return s.ProcessId
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime) SetGateway(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime {
	s.Gateway = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime) SetProcessId(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime {
	s.ProcessId = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime) Validate() error {
	return dara.Validate(s)
}

type ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource struct {
	// The default number of compute units (CUs) configured for task running.
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

func (s ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource) GoString() string {
	return s.String()
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource) GetCu() *string {
	return s.Cu
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource) GetImage() *string {
	return s.Image
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource) SetCu(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource {
	s.Cu = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource) SetImage(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource {
	s.Image = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource) SetResourceGroupId(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource) Validate() error {
	return dara.Validate(s)
}

type ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstances struct {
	// The scheduling dependency type. Valid values:
	//
	// 	- Normal
	//
	// 	- CrossCycle
	//
	// example:
	//
	// Normal
	DependencyType *string `json:"DependencyType,omitempty" xml:"DependencyType,omitempty"`
	// The information about a task instance.
	TaskInstance *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance `json:"TaskInstance,omitempty" xml:"TaskInstance,omitempty" type:"Struct"`
}

func (s ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstances) String() string {
	return dara.Prettify(s)
}

func (s ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstances) GoString() string {
	return s.String()
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstances) GetDependencyType() *string {
	return s.DependencyType
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstances) GetTaskInstance() *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance {
	return s.TaskInstance
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstances) SetDependencyType(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstances {
	s.DependencyType = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstances) SetTaskInstance(v *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstances {
	s.TaskInstance = v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstances) Validate() error {
	if s.TaskInstance != nil {
		if err := s.TaskInstance.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance struct {
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
	// The account ID of the creator.
	//
	// example:
	//
	// 1000
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The information about the associated data source.
	DataSource *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The environment in which the data source is used. Valid values:
	//
	// 	- Dev
	//
	// 	- Prod
	//
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
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
	// The account ID of the modifier.
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
	// The workspace ID.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The rerun mode.
	//
	// example:
	//
	// AllAllowed
	RerunMode *string `json:"RerunMode,omitempty" xml:"RerunMode,omitempty"`
	// The number of times the instance is run. By default, the value starts from 1.
	//
	// example:
	//
	// 1
	RunNumber *int32 `json:"RunNumber,omitempty" xml:"RunNumber,omitempty"`
	// The runtime information about the instance.
	Runtime *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceRuntime `json:"Runtime,omitempty" xml:"Runtime,omitempty" type:"Struct"`
	// The configurations of the runtime environment, such as the resource group information.
	RuntimeResource *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// The time when the instance started to run.
	//
	// example:
	//
	// 1710239005403
	StartedTime *int64 `json:"StartedTime,omitempty" xml:"StartedTime,omitempty"`
	// The status of the instance. Valid values:
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
	// 	- WaitTrigger: The instance is waiting to be triggered by external scheduling systems.
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
	// 3600
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
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
	// The scheduling time.
	//
	// example:
	//
	// 1710239005403
	TriggerTime *int64 `json:"TriggerTime,omitempty" xml:"TriggerTime,omitempty"`
	// The trigger type. Valid values:
	//
	// 	- Scheduler: scheduling cycle-based trigger
	//
	// 	- Manual: manual trigger
	//
	// example:
	//
	// Scheduler
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
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
	// 	- Normal
	//
	// 	- Manual
	//
	// 	- SmokeTest
	//
	// 	- SupplementData
	//
	// 	- ManualWorkflow
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

func (s ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) String() string {
	return dara.Prettify(s)
}

func (s ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) GoString() string {
	return s.String()
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) GetBizdate() *int64 {
	return s.Bizdate
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) GetDataSource() *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceDataSource {
	return s.DataSource
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) GetDescription() *string {
	return s.Description
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) GetEnvType() *string {
	return s.EnvType
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) GetFinishedTime() *int64 {
	return s.FinishedTime
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) GetId() *int64 {
	return s.Id
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) GetOwner() *string {
	return s.Owner
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) GetPeriodNumber() *int32 {
	return s.PeriodNumber
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) GetPriority() *int32 {
	return s.Priority
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) GetRerunMode() *string {
	return s.RerunMode
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) GetRunNumber() *int32 {
	return s.RunNumber
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) GetRuntime() *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceRuntime {
	return s.Runtime
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) GetRuntimeResource() *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceRuntimeResource {
	return s.RuntimeResource
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) GetStartedTime() *int64 {
	return s.StartedTime
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) GetStatus() *string {
	return s.Status
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) GetTaskName() *string {
	return s.TaskName
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) GetTaskType() *string {
	return s.TaskType
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) GetTimeout() *int32 {
	return s.Timeout
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) GetTriggerRecurrence() *string {
	return s.TriggerRecurrence
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) GetTriggerTime() *int64 {
	return s.TriggerTime
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) GetTriggerType() *string {
	return s.TriggerType
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) GetWorkflowInstanceId() *int64 {
	return s.WorkflowInstanceId
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) GetWorkflowInstanceType() *string {
	return s.WorkflowInstanceType
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) GetWorkflowName() *string {
	return s.WorkflowName
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) SetBaselineId(v int64) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance {
	s.BaselineId = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) SetBizdate(v int64) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance {
	s.Bizdate = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) SetCreateTime(v int64) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance {
	s.CreateTime = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) SetCreateUser(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance {
	s.CreateUser = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) SetDataSource(v *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceDataSource) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance {
	s.DataSource = v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) SetDescription(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance {
	s.Description = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) SetEnvType(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance {
	s.EnvType = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) SetFinishedTime(v int64) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance {
	s.FinishedTime = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) SetId(v int64) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance {
	s.Id = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) SetModifyTime(v int64) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance {
	s.ModifyTime = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) SetModifyUser(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance {
	s.ModifyUser = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) SetOwner(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance {
	s.Owner = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) SetPeriodNumber(v int32) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance {
	s.PeriodNumber = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) SetPriority(v int32) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance {
	s.Priority = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) SetProjectId(v int64) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance {
	s.ProjectId = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) SetRerunMode(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance {
	s.RerunMode = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) SetRunNumber(v int32) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance {
	s.RunNumber = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) SetRuntime(v *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceRuntime) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance {
	s.Runtime = v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) SetRuntimeResource(v *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceRuntimeResource) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance {
	s.RuntimeResource = v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) SetStartedTime(v int64) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance {
	s.StartedTime = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) SetStatus(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance {
	s.Status = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) SetTaskId(v int64) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance {
	s.TaskId = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) SetTaskName(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance {
	s.TaskName = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) SetTaskType(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance {
	s.TaskType = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) SetTimeout(v int32) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance {
	s.Timeout = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) SetTriggerRecurrence(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance {
	s.TriggerRecurrence = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) SetTriggerTime(v int64) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance {
	s.TriggerTime = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) SetTriggerType(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance {
	s.TriggerType = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) SetWorkflowId(v int64) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance {
	s.WorkflowId = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) SetWorkflowInstanceId(v int64) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance {
	s.WorkflowInstanceId = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) SetWorkflowInstanceType(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance {
	s.WorkflowInstanceType = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) SetWorkflowName(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance {
	s.WorkflowName = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstance) Validate() error {
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

type ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceDataSource struct {
	// The name of the data source.
	//
	// example:
	//
	// mysql_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceDataSource) String() string {
	return dara.Prettify(s)
}

func (s ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceDataSource) GoString() string {
	return s.String()
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceDataSource) GetName() *string {
	return s.Name
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceDataSource) SetName(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceDataSource {
	s.Name = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceDataSource) Validate() error {
	return dara.Validate(s)
}

type ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceRuntime struct {
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

func (s ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceRuntime) String() string {
	return dara.Prettify(s)
}

func (s ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceRuntime) GoString() string {
	return s.String()
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceRuntime) GetGateway() *string {
	return s.Gateway
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceRuntime) GetProcessId() *string {
	return s.ProcessId
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceRuntime) SetGateway(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceRuntime {
	s.Gateway = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceRuntime) SetProcessId(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceRuntime {
	s.ProcessId = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceRuntime) Validate() error {
	return dara.Validate(s)
}

type ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceRuntimeResource struct {
	// The default number of compute units (CUs) configured for task running.
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

func (s ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceRuntimeResource) GoString() string {
	return s.String()
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceRuntimeResource) GetCu() *string {
	return s.Cu
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceRuntimeResource) GetImage() *string {
	return s.Image
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceRuntimeResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceRuntimeResource) SetCu(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceRuntimeResource {
	s.Cu = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceRuntimeResource) SetImage(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceRuntimeResource {
	s.Image = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceRuntimeResource) SetResourceGroupId(v string) *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponseBodyPagingInfoUpstreamTaskInstancesTaskInstanceRuntimeResource) Validate() error {
	return dara.Validate(s)
}
