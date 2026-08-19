// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDownstreamTaskInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListDownstreamTaskInstancesResponseBodyPagingInfo) *ListDownstreamTaskInstancesResponseBody
	GetPagingInfo() *ListDownstreamTaskInstancesResponseBodyPagingInfo
	SetRequestId(v string) *ListDownstreamTaskInstancesResponseBody
	GetRequestId() *string
}

type ListDownstreamTaskInstancesResponseBody struct {
	// The pagination information.
	PagingInfo *ListDownstreamTaskInstancesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID, which is used to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDownstreamTaskInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDownstreamTaskInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDownstreamTaskInstancesResponseBody) GetPagingInfo() *ListDownstreamTaskInstancesResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListDownstreamTaskInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDownstreamTaskInstancesResponseBody) SetPagingInfo(v *ListDownstreamTaskInstancesResponseBodyPagingInfo) *ListDownstreamTaskInstancesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBody) SetRequestId(v string) *ListDownstreamTaskInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDownstreamTaskInstancesResponseBodyPagingInfo struct {
	// The list of downstream task instances.
	DownstreamTaskInstances []*ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstances `json:"DownstreamTaskInstances,omitempty" xml:"DownstreamTaskInstances,omitempty" type:"Repeated"`
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
	// **[Deprecated]*	- The list of task instances. This field is deprecated. Use DownstreamTaskInstances instead.
	TaskInstances []*ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances `json:"TaskInstances,omitempty" xml:"TaskInstances,omitempty" type:"Repeated"`
	// The total number of records.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDownstreamTaskInstancesResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDownstreamTaskInstancesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfo) GetDownstreamTaskInstances() []*ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstances {
	return s.DownstreamTaskInstances
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfo) GetTaskInstances() []*ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	return s.TaskInstances
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfo) SetDownstreamTaskInstances(v []*ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstances) *ListDownstreamTaskInstancesResponseBodyPagingInfo {
	s.DownstreamTaskInstances = v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfo) SetPageNumber(v int32) *ListDownstreamTaskInstancesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfo) SetPageSize(v int32) *ListDownstreamTaskInstancesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfo) SetTaskInstances(v []*ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) *ListDownstreamTaskInstancesResponseBodyPagingInfo {
	s.TaskInstances = v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfo) SetTotalCount(v int32) *ListDownstreamTaskInstancesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfo) Validate() error {
	if s.DownstreamTaskInstances != nil {
		for _, item := range s.DownstreamTaskInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstances struct {
	// The dependency type.
	//
	// example:
	//
	// Normal
	DependencyType *string `json:"DependencyType,omitempty" xml:"DependencyType,omitempty"`
	// The task instance.
	TaskInstance *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance `json:"TaskInstance,omitempty" xml:"TaskInstance,omitempty" type:"Struct"`
}

func (s ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstances) String() string {
	return dara.Prettify(s)
}

func (s ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstances) GoString() string {
	return s.String()
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstances) GetDependencyType() *string {
	return s.DependencyType
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstances) GetTaskInstance() *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance {
	return s.TaskInstance
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstances) SetDependencyType(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstances {
	s.DependencyType = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstances) SetTaskInstance(v *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstances {
	s.TaskInstance = v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstances) Validate() error {
	if s.TaskInstance != nil {
		if err := s.TaskInstance.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance struct {
	// The baseline ID.
	//
	// example:
	//
	// 1234
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The business date.
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
	// The account ID of the user who created the instance.
	//
	// example:
	//
	// 1000
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The data source information associated with the instance.
	DataSource *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The environment of the target data source. Valid values:
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
	// The unique identifier of the task instance.
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
	// The account ID of the user who last modified the instance.
	//
	// example:
	//
	// 1000
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// The account ID of the node owner.
	//
	// example:
	//
	// 1000
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The period number. Indicates which scheduling cycle of the day the instance is in.
	//
	// example:
	//
	// 1
	PeriodNumber *int32 `json:"PeriodNumber,omitempty" xml:"PeriodNumber,omitempty"`
	// The run priority of the node. Minimum value: 1. Maximum value: 8. A larger value indicates a higher priority. Default value: 1.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The rerun configuration of the node.
	//
	// example:
	//
	// AllAllowed
	RerunMode *string `json:"RerunMode,omitempty" xml:"RerunMode,omitempty"`
	// The current run number. The default value starts from 1.
	//
	// example:
	//
	// 1
	RunNumber *int32 `json:"RunNumber,omitempty" xml:"RunNumber,omitempty"`
	// The runtime information of the instance.
	Runtime *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceRuntime `json:"Runtime,omitempty" xml:"Runtime,omitempty" type:"Struct"`
	// The runtime environment configuration, such as resource group information.
	RuntimeResource *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// The time when the instance started running.
	//
	// example:
	//
	// 1710239005403
	StartedTime *int64 `json:"StartedTime,omitempty" xml:"StartedTime,omitempty"`
	// The run status of the instance.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the corresponding node.
	//
	// example:
	//
	// 1234
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the corresponding node.
	//
	// example:
	//
	// SQL node
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The type of the corresponding node.
	//
	// example:
	//
	// ODPS_SQL
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The timeout period for node execution. Unit: seconds.
	//
	// example:
	//
	// 3600
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The run mode when triggered. This parameter takes effect when TriggerType is set to Scheduler. Valid values:
	//
	// - Pause: paused.
	//
	// - Skip: dry run.
	//
	// - Normal: normal execution.
	//
	// example:
	//
	// Normal
	TriggerRecurrence *string `json:"TriggerRecurrence,omitempty" xml:"TriggerRecurrence,omitempty"`
	// The scheduled trigger time.
	//
	// example:
	//
	// 1710239005403
	TriggerTime *int64 `json:"TriggerTime,omitempty" xml:"TriggerTime,omitempty"`
	// The trigger type.
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
	// The ID of the workflow instance to which the instance belongs.
	//
	// example:
	//
	// 1234
	WorkflowInstanceId *int64 `json:"WorkflowInstanceId,omitempty" xml:"WorkflowInstanceId,omitempty"`
	// The type of the workflow instance to which the instance belongs.
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

func (s ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) String() string {
	return dara.Prettify(s)
}

func (s ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) GoString() string {
	return s.String()
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) GetBizdate() *int64 {
	return s.Bizdate
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) GetDataSource() *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceDataSource {
	return s.DataSource
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) GetDescription() *string {
	return s.Description
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) GetEnvType() *string {
	return s.EnvType
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) GetFinishedTime() *int64 {
	return s.FinishedTime
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) GetId() *int64 {
	return s.Id
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) GetOwner() *string {
	return s.Owner
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) GetPeriodNumber() *int32 {
	return s.PeriodNumber
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) GetPriority() *int32 {
	return s.Priority
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) GetRerunMode() *string {
	return s.RerunMode
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) GetRunNumber() *int32 {
	return s.RunNumber
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) GetRuntime() *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceRuntime {
	return s.Runtime
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) GetRuntimeResource() *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceRuntimeResource {
	return s.RuntimeResource
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) GetStartedTime() *int64 {
	return s.StartedTime
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) GetStatus() *string {
	return s.Status
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) GetTaskName() *string {
	return s.TaskName
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) GetTaskType() *string {
	return s.TaskType
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) GetTimeout() *int32 {
	return s.Timeout
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) GetTriggerRecurrence() *string {
	return s.TriggerRecurrence
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) GetTriggerTime() *int64 {
	return s.TriggerTime
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) GetTriggerType() *string {
	return s.TriggerType
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) GetWorkflowInstanceId() *int64 {
	return s.WorkflowInstanceId
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) GetWorkflowInstanceType() *string {
	return s.WorkflowInstanceType
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) GetWorkflowName() *string {
	return s.WorkflowName
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) SetBaselineId(v int64) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance {
	s.BaselineId = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) SetBizdate(v int64) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance {
	s.Bizdate = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) SetCreateTime(v int64) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance {
	s.CreateTime = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) SetCreateUser(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance {
	s.CreateUser = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) SetDataSource(v *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceDataSource) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance {
	s.DataSource = v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) SetDescription(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance {
	s.Description = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) SetEnvType(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance {
	s.EnvType = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) SetFinishedTime(v int64) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance {
	s.FinishedTime = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) SetId(v int64) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance {
	s.Id = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) SetModifyTime(v int64) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance {
	s.ModifyTime = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) SetModifyUser(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance {
	s.ModifyUser = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) SetOwner(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance {
	s.Owner = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) SetPeriodNumber(v int32) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance {
	s.PeriodNumber = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) SetPriority(v int32) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance {
	s.Priority = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) SetProjectId(v int64) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance {
	s.ProjectId = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) SetRerunMode(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance {
	s.RerunMode = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) SetRunNumber(v int32) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance {
	s.RunNumber = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) SetRuntime(v *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceRuntime) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance {
	s.Runtime = v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) SetRuntimeResource(v *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceRuntimeResource) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance {
	s.RuntimeResource = v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) SetStartedTime(v int64) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance {
	s.StartedTime = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) SetStatus(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance {
	s.Status = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) SetTaskId(v int64) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance {
	s.TaskId = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) SetTaskName(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance {
	s.TaskName = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) SetTaskType(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance {
	s.TaskType = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) SetTimeout(v int32) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance {
	s.Timeout = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) SetTriggerRecurrence(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance {
	s.TriggerRecurrence = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) SetTriggerTime(v int64) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance {
	s.TriggerTime = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) SetTriggerType(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance {
	s.TriggerType = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) SetWorkflowId(v int64) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance {
	s.WorkflowId = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) SetWorkflowInstanceId(v int64) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance {
	s.WorkflowInstanceId = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) SetWorkflowInstanceType(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance {
	s.WorkflowInstanceType = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) SetWorkflowName(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance {
	s.WorkflowName = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstance) Validate() error {
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

type ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceDataSource struct {
	// The data source name.
	//
	// example:
	//
	// mysql_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceDataSource) String() string {
	return dara.Prettify(s)
}

func (s ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceDataSource) GoString() string {
	return s.String()
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceDataSource) GetName() *string {
	return s.Name
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceDataSource) SetName(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceDataSource {
	s.Name = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceDataSource) Validate() error {
	return dara.Validate(s)
}

type ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceRuntime struct {
	// The machine on which the instance runs.
	//
	// example:
	//
	// cn-shanghai.1.2
	Gateway *string `json:"Gateway,omitempty" xml:"Gateway,omitempty"`
	// The unique run ID.
	//
	// example:
	//
	// T3_123
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
}

func (s ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceRuntime) String() string {
	return dara.Prettify(s)
}

func (s ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceRuntime) GoString() string {
	return s.String()
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceRuntime) GetGateway() *string {
	return s.Gateway
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceRuntime) GetProcessId() *string {
	return s.ProcessId
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceRuntime) SetGateway(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceRuntime {
	s.Gateway = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceRuntime) SetProcessId(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceRuntime {
	s.ProcessId = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceRuntime) Validate() error {
	return dara.Validate(s)
}

type ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceRuntimeResource struct {
	// The compute unit (CU) consumption configured for the node.
	//
	// example:
	//
	// 0.25
	Cu *string `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The image ID configured for the node.
	//
	// example:
	//
	// i-xxxxxx
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The identifier of the schedule resource group configured for the node.
	//
	// example:
	//
	// S_res_group_524258031846018_1684XXXXXXXXX
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceRuntimeResource) GoString() string {
	return s.String()
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceRuntimeResource) GetCu() *string {
	return s.Cu
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceRuntimeResource) GetImage() *string {
	return s.Image
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceRuntimeResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceRuntimeResource) SetCu(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceRuntimeResource {
	s.Cu = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceRuntimeResource) SetImage(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceRuntimeResource {
	s.Image = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceRuntimeResource) SetResourceGroupId(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoDownstreamTaskInstancesTaskInstanceRuntimeResource) Validate() error {
	return dara.Validate(s)
}

type ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances struct {
	// The baseline ID.
	//
	// example:
	//
	// 1234
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The business date.
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
	// The account ID of the user who created the instance.
	//
	// example:
	//
	// 1000
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The data source information associated with the instance.
	DataSource *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The environment of the target data source. Valid values:
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
	// The unique identifier of the task instance.
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
	// The account ID of the user who last modified the instance.
	//
	// example:
	//
	// 1000
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// The account ID of the node owner.
	//
	// example:
	//
	// 100
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The period number. Indicates which scheduling cycle of the day the instance is in.
	//
	// example:
	//
	// 1
	PeriodNumber *int32 `json:"PeriodNumber,omitempty" xml:"PeriodNumber,omitempty"`
	// The run priority of the node. Minimum value: 1. Maximum value: 8. A larger value indicates a higher priority. Default value: 1.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- The project environment. This field is deprecated. Use EnvType instead.
	//
	// example:
	//
	// Prod
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The rerun configuration of the node.
	//
	// example:
	//
	// AllAllowed
	RerunMode *string `json:"RerunMode,omitempty" xml:"RerunMode,omitempty"`
	// The current run number. The default value starts from 1.
	//
	// example:
	//
	// 1
	RunNumber *int32 `json:"RunNumber,omitempty" xml:"RunNumber,omitempty"`
	// The runtime information of the instance.
	Runtime *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime `json:"Runtime,omitempty" xml:"Runtime,omitempty" type:"Struct"`
	// The resource group information associated with the instance.
	RuntimeResource *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// The time when the instance started running.
	//
	// example:
	//
	// 1710239005403
	StartedTime *int64 `json:"StartedTime,omitempty" xml:"StartedTime,omitempty"`
	// The run status of the instance.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The dependency type.
	//
	// example:
	//
	// Normal
	StepType *string `json:"StepType,omitempty" xml:"StepType,omitempty"`
	// The ID of the corresponding node.
	//
	// example:
	//
	// 1234
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the corresponding node.
	//
	// example:
	//
	// SQL node
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The type of the corresponding node.
	//
	// example:
	//
	// ODPS_SQL
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The timeout period for node execution. Unit: seconds.
	//
	// example:
	//
	// 3600
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The run mode at the time of triggering. This parameter takes effect when TriggerType is set to Scheduler.
	//
	// example:
	//
	// Normal
	TriggerRecurrence *string `json:"TriggerRecurrence,omitempty" xml:"TriggerRecurrence,omitempty"`
	// The scheduled trigger time.
	//
	// example:
	//
	// 1710239005403
	TriggerTime *int64 `json:"TriggerTime,omitempty" xml:"TriggerTime,omitempty"`
	// The trigger type.
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
	// The ID of the workflow instance to which the instance belongs.
	//
	// example:
	//
	// 1234
	WorkflowInstanceId *int64 `json:"WorkflowInstanceId,omitempty" xml:"WorkflowInstanceId,omitempty"`
	// The type of the workflow instance to which the instance belongs.
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

func (s ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) String() string {
	return dara.Prettify(s)
}

func (s ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GoString() string {
	return s.String()
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetBizdate() *int64 {
	return s.Bizdate
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetDataSource() *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesDataSource {
	return s.DataSource
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetDescription() *string {
	return s.Description
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetEnvType() *string {
	return s.EnvType
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetFinishedTime() *int64 {
	return s.FinishedTime
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetId() *int64 {
	return s.Id
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetOwner() *string {
	return s.Owner
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetPeriodNumber() *int32 {
	return s.PeriodNumber
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetPriority() *int32 {
	return s.Priority
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetRerunMode() *string {
	return s.RerunMode
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetRunNumber() *int32 {
	return s.RunNumber
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetRuntime() *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime {
	return s.Runtime
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetRuntimeResource() *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource {
	return s.RuntimeResource
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetStartedTime() *int64 {
	return s.StartedTime
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetStatus() *string {
	return s.Status
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetStepType() *string {
	return s.StepType
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetTaskName() *string {
	return s.TaskName
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetTaskType() *string {
	return s.TaskType
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetTimeout() *int32 {
	return s.Timeout
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetTriggerRecurrence() *string {
	return s.TriggerRecurrence
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetTriggerTime() *int64 {
	return s.TriggerTime
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetTriggerType() *string {
	return s.TriggerType
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetWorkflowInstanceId() *int64 {
	return s.WorkflowInstanceId
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetWorkflowInstanceType() *string {
	return s.WorkflowInstanceType
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) GetWorkflowName() *string {
	return s.WorkflowName
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetBaselineId(v int64) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.BaselineId = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetBizdate(v int64) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.Bizdate = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetCreateTime(v int64) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.CreateTime = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetCreateUser(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.CreateUser = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetDataSource(v *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesDataSource) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.DataSource = v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetDescription(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.Description = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetEnvType(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.EnvType = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetFinishedTime(v int64) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.FinishedTime = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetId(v int64) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.Id = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetModifyTime(v int64) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.ModifyTime = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetModifyUser(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.ModifyUser = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetOwner(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.Owner = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetPeriodNumber(v int32) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.PeriodNumber = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetPriority(v int32) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.Priority = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetProjectEnv(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.ProjectEnv = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetProjectId(v int64) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.ProjectId = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetRerunMode(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.RerunMode = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetRunNumber(v int32) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.RunNumber = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetRuntime(v *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.Runtime = v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetRuntimeResource(v *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.RuntimeResource = v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetStartedTime(v int64) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.StartedTime = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetStatus(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.Status = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetStepType(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.StepType = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetTaskId(v int64) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.TaskId = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetTaskName(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.TaskName = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetTaskType(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.TaskType = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetTimeout(v int32) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.Timeout = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetTriggerRecurrence(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.TriggerRecurrence = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetTriggerTime(v int64) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.TriggerTime = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetTriggerType(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.TriggerType = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetWorkflowId(v int64) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.WorkflowId = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetWorkflowInstanceId(v int64) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.WorkflowInstanceId = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetWorkflowInstanceType(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.WorkflowInstanceType = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) SetWorkflowName(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances {
	s.WorkflowName = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstances) Validate() error {
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

type ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesDataSource struct {
	// The data source name.
	//
	// example:
	//
	// mysql_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesDataSource) String() string {
	return dara.Prettify(s)
}

func (s ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesDataSource) GoString() string {
	return s.String()
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesDataSource) GetName() *string {
	return s.Name
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesDataSource) SetName(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesDataSource {
	s.Name = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesDataSource) Validate() error {
	return dara.Validate(s)
}

type ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime struct {
	// The machine on which the instance runs.
	//
	// example:
	//
	// cn-shanghai.1.2
	Gateway *string `json:"Gateway,omitempty" xml:"Gateway,omitempty"`
	// The unique run ID.
	//
	// example:
	//
	// T3_123
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
}

func (s ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime) String() string {
	return dara.Prettify(s)
}

func (s ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime) GoString() string {
	return s.String()
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime) GetGateway() *string {
	return s.Gateway
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime) GetProcessId() *string {
	return s.ProcessId
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime) SetGateway(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime {
	s.Gateway = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime) SetProcessId(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime {
	s.ProcessId = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntime) Validate() error {
	return dara.Validate(s)
}

type ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource struct {
	// The compute unit (CU) consumption configured for the node.
	//
	// example:
	//
	// 0.25
	Cu *string `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The image ID configured for the node.
	//
	// example:
	//
	// i-xxxxxx
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The identifier of the schedule resource group configured for the node.
	//
	// example:
	//
	// S_res_group_524258031846018_1684XXXXXXXXX
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource) GoString() string {
	return s.String()
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource) GetCu() *string {
	return s.Cu
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource) GetImage() *string {
	return s.Image
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource) SetCu(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource {
	s.Cu = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource) SetImage(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource {
	s.Image = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource) SetResourceGroupId(v string) *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponseBodyPagingInfoTaskInstancesRuntimeResource) Validate() error {
	return dara.Validate(s)
}
