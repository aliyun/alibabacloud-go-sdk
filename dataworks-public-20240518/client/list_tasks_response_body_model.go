// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListTasksResponseBodyPagingInfo) *ListTasksResponseBody
	GetPagingInfo() *ListTasksResponseBodyPagingInfo
	SetRequestId(v string) *ListTasksResponseBody
	GetRequestId() *string
}

type ListTasksResponseBody struct {
	// The pagination information.
	PagingInfo *ListTasksResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBody) GetPagingInfo() *ListTasksResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTasksResponseBody) SetPagingInfo(v *ListTasksResponseBodyPagingInfo) *ListTasksResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListTasksResponseBody) SetRequestId(v string) *ListTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTasksResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTasksResponseBodyPagingInfo struct {
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
	// The tasks.
	Tasks []*ListTasksResponseBodyPagingInfoTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTasksResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListTasksResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTasksResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTasksResponseBodyPagingInfo) GetTasks() []*ListTasksResponseBodyPagingInfoTasks {
	return s.Tasks
}

func (s *ListTasksResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTasksResponseBodyPagingInfo) SetPageNumber(v int32) *ListTasksResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListTasksResponseBodyPagingInfo) SetPageSize(v int32) *ListTasksResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListTasksResponseBodyPagingInfo) SetTasks(v []*ListTasksResponseBodyPagingInfoTasks) *ListTasksResponseBodyPagingInfo {
	s.Tasks = v
	return s
}

func (s *ListTasksResponseBodyPagingInfo) SetTotalCount(v int32) *ListTasksResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListTasksResponseBodyPagingInfo) Validate() error {
	if s.Tasks != nil {
		for _, item := range s.Tasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTasksResponseBodyPagingInfoTasks struct {
	// The baseline ID.
	//
	// example:
	//
	// 1234
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
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
	DataSource *ListTasksResponseBodyPagingInfoTasksDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The description of the task.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The instance generation mode. Valid values:
	//
	// 	- T+1
	//
	// 	- Immediately
	//
	// example:
	//
	// T+1
	InstanceMode *string `json:"InstanceMode,omitempty" xml:"InstanceMode,omitempty"`
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
	// The name of the task.
	//
	// example:
	//
	// SQL node
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The account ID of the task owner.
	//
	// example:
	//
	// 1000
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The priority of the task. Valid values: 1 to 8. A larger value indicates a higher priority. Default value: 1.
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
	// The workspace ID.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The rerun interval. Unit: seconds.
	//
	// example:
	//
	// 60
	RerunInterval *int32 `json:"RerunInterval,omitempty" xml:"RerunInterval,omitempty"`
	// The rerun mode.
	//
	// Valid values:
	//
	// 	- AllDenied: The task cannot be rerun regardless of whether the task is successfully run or fails to run.
	//
	// 	- FailureAllowed: The task can be rerun only after it fails to run.
	//
	// 	- AllAllowed: The task can be rerun regardless of whether it is successfully run or fails to run.
	//
	// example:
	//
	// AllAllowed
	RerunMode *string `json:"RerunMode,omitempty" xml:"RerunMode,omitempty"`
	// The number of times that the task is rerun. This parameter takes effect only if the RerunMode parameter is set to AllAllowed or FailureAllowed.
	//
	// example:
	//
	// 3
	RerunTimes *int32 `json:"RerunTimes,omitempty" xml:"RerunTimes,omitempty"`
	// The configurations of the runtime environment, such as the resource group information.
	RuntimeResource *ListTasksResponseBodyPagingInfoTasksRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// The list of script parameters.
	//
	// example:
	//
	// para1=$bizdate para2=$[yyyymmdd]
	ScriptParameters *string `json:"ScriptParameters,omitempty" xml:"ScriptParameters,omitempty"`
	// The timeout period of task running. Unit: seconds.
	//
	// example:
	//
	// 3600
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The method to trigger task scheduling.
	Trigger *ListTasksResponseBodyPagingInfoTasksTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
	// The type of the task.
	//
	// example:
	//
	// ODPS_SQL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the workflow to which the task belongs.
	//
	// example:
	//
	// 1234
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s ListTasksResponseBodyPagingInfoTasks) String() string {
	return dara.Prettify(s)
}

func (s ListTasksResponseBodyPagingInfoTasks) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyPagingInfoTasks) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *ListTasksResponseBodyPagingInfoTasks) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListTasksResponseBodyPagingInfoTasks) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListTasksResponseBodyPagingInfoTasks) GetDataSource() *ListTasksResponseBodyPagingInfoTasksDataSource {
	return s.DataSource
}

func (s *ListTasksResponseBodyPagingInfoTasks) GetDescription() *string {
	return s.Description
}

func (s *ListTasksResponseBodyPagingInfoTasks) GetId() *int64 {
	return s.Id
}

func (s *ListTasksResponseBodyPagingInfoTasks) GetInstanceMode() *string {
	return s.InstanceMode
}

func (s *ListTasksResponseBodyPagingInfoTasks) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListTasksResponseBodyPagingInfoTasks) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *ListTasksResponseBodyPagingInfoTasks) GetName() *string {
	return s.Name
}

func (s *ListTasksResponseBodyPagingInfoTasks) GetOwner() *string {
	return s.Owner
}

func (s *ListTasksResponseBodyPagingInfoTasks) GetPriority() *int32 {
	return s.Priority
}

func (s *ListTasksResponseBodyPagingInfoTasks) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *ListTasksResponseBodyPagingInfoTasks) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListTasksResponseBodyPagingInfoTasks) GetRerunInterval() *int32 {
	return s.RerunInterval
}

func (s *ListTasksResponseBodyPagingInfoTasks) GetRerunMode() *string {
	return s.RerunMode
}

func (s *ListTasksResponseBodyPagingInfoTasks) GetRerunTimes() *int32 {
	return s.RerunTimes
}

func (s *ListTasksResponseBodyPagingInfoTasks) GetRuntimeResource() *ListTasksResponseBodyPagingInfoTasksRuntimeResource {
	return s.RuntimeResource
}

func (s *ListTasksResponseBodyPagingInfoTasks) GetScriptParameters() *string {
	return s.ScriptParameters
}

func (s *ListTasksResponseBodyPagingInfoTasks) GetTimeout() *int32 {
	return s.Timeout
}

func (s *ListTasksResponseBodyPagingInfoTasks) GetTrigger() *ListTasksResponseBodyPagingInfoTasksTrigger {
	return s.Trigger
}

func (s *ListTasksResponseBodyPagingInfoTasks) GetType() *string {
	return s.Type
}

func (s *ListTasksResponseBodyPagingInfoTasks) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *ListTasksResponseBodyPagingInfoTasks) SetBaselineId(v int64) *ListTasksResponseBodyPagingInfoTasks {
	s.BaselineId = &v
	return s
}

func (s *ListTasksResponseBodyPagingInfoTasks) SetCreateTime(v int64) *ListTasksResponseBodyPagingInfoTasks {
	s.CreateTime = &v
	return s
}

func (s *ListTasksResponseBodyPagingInfoTasks) SetCreateUser(v string) *ListTasksResponseBodyPagingInfoTasks {
	s.CreateUser = &v
	return s
}

func (s *ListTasksResponseBodyPagingInfoTasks) SetDataSource(v *ListTasksResponseBodyPagingInfoTasksDataSource) *ListTasksResponseBodyPagingInfoTasks {
	s.DataSource = v
	return s
}

func (s *ListTasksResponseBodyPagingInfoTasks) SetDescription(v string) *ListTasksResponseBodyPagingInfoTasks {
	s.Description = &v
	return s
}

func (s *ListTasksResponseBodyPagingInfoTasks) SetId(v int64) *ListTasksResponseBodyPagingInfoTasks {
	s.Id = &v
	return s
}

func (s *ListTasksResponseBodyPagingInfoTasks) SetInstanceMode(v string) *ListTasksResponseBodyPagingInfoTasks {
	s.InstanceMode = &v
	return s
}

func (s *ListTasksResponseBodyPagingInfoTasks) SetModifyTime(v int64) *ListTasksResponseBodyPagingInfoTasks {
	s.ModifyTime = &v
	return s
}

func (s *ListTasksResponseBodyPagingInfoTasks) SetModifyUser(v string) *ListTasksResponseBodyPagingInfoTasks {
	s.ModifyUser = &v
	return s
}

func (s *ListTasksResponseBodyPagingInfoTasks) SetName(v string) *ListTasksResponseBodyPagingInfoTasks {
	s.Name = &v
	return s
}

func (s *ListTasksResponseBodyPagingInfoTasks) SetOwner(v string) *ListTasksResponseBodyPagingInfoTasks {
	s.Owner = &v
	return s
}

func (s *ListTasksResponseBodyPagingInfoTasks) SetPriority(v int32) *ListTasksResponseBodyPagingInfoTasks {
	s.Priority = &v
	return s
}

func (s *ListTasksResponseBodyPagingInfoTasks) SetProjectEnv(v string) *ListTasksResponseBodyPagingInfoTasks {
	s.ProjectEnv = &v
	return s
}

func (s *ListTasksResponseBodyPagingInfoTasks) SetProjectId(v int64) *ListTasksResponseBodyPagingInfoTasks {
	s.ProjectId = &v
	return s
}

func (s *ListTasksResponseBodyPagingInfoTasks) SetRerunInterval(v int32) *ListTasksResponseBodyPagingInfoTasks {
	s.RerunInterval = &v
	return s
}

func (s *ListTasksResponseBodyPagingInfoTasks) SetRerunMode(v string) *ListTasksResponseBodyPagingInfoTasks {
	s.RerunMode = &v
	return s
}

func (s *ListTasksResponseBodyPagingInfoTasks) SetRerunTimes(v int32) *ListTasksResponseBodyPagingInfoTasks {
	s.RerunTimes = &v
	return s
}

func (s *ListTasksResponseBodyPagingInfoTasks) SetRuntimeResource(v *ListTasksResponseBodyPagingInfoTasksRuntimeResource) *ListTasksResponseBodyPagingInfoTasks {
	s.RuntimeResource = v
	return s
}

func (s *ListTasksResponseBodyPagingInfoTasks) SetScriptParameters(v string) *ListTasksResponseBodyPagingInfoTasks {
	s.ScriptParameters = &v
	return s
}

func (s *ListTasksResponseBodyPagingInfoTasks) SetTimeout(v int32) *ListTasksResponseBodyPagingInfoTasks {
	s.Timeout = &v
	return s
}

func (s *ListTasksResponseBodyPagingInfoTasks) SetTrigger(v *ListTasksResponseBodyPagingInfoTasksTrigger) *ListTasksResponseBodyPagingInfoTasks {
	s.Trigger = v
	return s
}

func (s *ListTasksResponseBodyPagingInfoTasks) SetType(v string) *ListTasksResponseBodyPagingInfoTasks {
	s.Type = &v
	return s
}

func (s *ListTasksResponseBodyPagingInfoTasks) SetWorkflowId(v int64) *ListTasksResponseBodyPagingInfoTasks {
	s.WorkflowId = &v
	return s
}

func (s *ListTasksResponseBodyPagingInfoTasks) Validate() error {
	if s.DataSource != nil {
		if err := s.DataSource.Validate(); err != nil {
			return err
		}
	}
	if s.RuntimeResource != nil {
		if err := s.RuntimeResource.Validate(); err != nil {
			return err
		}
	}
	if s.Trigger != nil {
		if err := s.Trigger.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTasksResponseBodyPagingInfoTasksDataSource struct {
	// The name of the data source.
	//
	// example:
	//
	// mysql_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListTasksResponseBodyPagingInfoTasksDataSource) String() string {
	return dara.Prettify(s)
}

func (s ListTasksResponseBodyPagingInfoTasksDataSource) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyPagingInfoTasksDataSource) GetName() *string {
	return s.Name
}

func (s *ListTasksResponseBodyPagingInfoTasksDataSource) SetName(v string) *ListTasksResponseBodyPagingInfoTasksDataSource {
	s.Name = &v
	return s
}

func (s *ListTasksResponseBodyPagingInfoTasksDataSource) Validate() error {
	return dara.Validate(s)
}

type ListTasksResponseBodyPagingInfoTasksRuntimeResource struct {
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

func (s ListTasksResponseBodyPagingInfoTasksRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s ListTasksResponseBodyPagingInfoTasksRuntimeResource) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyPagingInfoTasksRuntimeResource) GetCu() *string {
	return s.Cu
}

func (s *ListTasksResponseBodyPagingInfoTasksRuntimeResource) GetImage() *string {
	return s.Image
}

func (s *ListTasksResponseBodyPagingInfoTasksRuntimeResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListTasksResponseBodyPagingInfoTasksRuntimeResource) SetCu(v string) *ListTasksResponseBodyPagingInfoTasksRuntimeResource {
	s.Cu = &v
	return s
}

func (s *ListTasksResponseBodyPagingInfoTasksRuntimeResource) SetImage(v string) *ListTasksResponseBodyPagingInfoTasksRuntimeResource {
	s.Image = &v
	return s
}

func (s *ListTasksResponseBodyPagingInfoTasksRuntimeResource) SetResourceGroupId(v string) *ListTasksResponseBodyPagingInfoTasksRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTasksResponseBodyPagingInfoTasksRuntimeResource) Validate() error {
	return dara.Validate(s)
}

type ListTasksResponseBodyPagingInfoTasksTrigger struct {
	// The CRON expression of the task. This parameter takes effect only if the Type parameter is set to Scheduler.
	//
	// example:
	//
	// 00 00 00 	- 	- ?
	Cron *string `json:"Cron,omitempty" xml:"Cron,omitempty"`
	// The end time of the time range during which the task is periodically scheduled. This parameter takes effect only if the Type parameter is set to Scheduler.
	//
	// example:
	//
	// 9999-01-01 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The running mode of the task after it is triggered. This parameter takes effect only if the Type parameter is set to Scheduler.
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
	Recurrence *string `json:"Recurrence,omitempty" xml:"Recurrence,omitempty"`
	// The start time of the time range during which the task is periodically scheduled. This parameter takes effect only if the Type parameter is set to Scheduler.
	//
	// example:
	//
	// 1970-01-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListTasksResponseBodyPagingInfoTasksTrigger) String() string {
	return dara.Prettify(s)
}

func (s ListTasksResponseBodyPagingInfoTasksTrigger) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyPagingInfoTasksTrigger) GetCron() *string {
	return s.Cron
}

func (s *ListTasksResponseBodyPagingInfoTasksTrigger) GetEndTime() *string {
	return s.EndTime
}

func (s *ListTasksResponseBodyPagingInfoTasksTrigger) GetRecurrence() *string {
	return s.Recurrence
}

func (s *ListTasksResponseBodyPagingInfoTasksTrigger) GetStartTime() *string {
	return s.StartTime
}

func (s *ListTasksResponseBodyPagingInfoTasksTrigger) GetType() *string {
	return s.Type
}

func (s *ListTasksResponseBodyPagingInfoTasksTrigger) SetCron(v string) *ListTasksResponseBodyPagingInfoTasksTrigger {
	s.Cron = &v
	return s
}

func (s *ListTasksResponseBodyPagingInfoTasksTrigger) SetEndTime(v string) *ListTasksResponseBodyPagingInfoTasksTrigger {
	s.EndTime = &v
	return s
}

func (s *ListTasksResponseBodyPagingInfoTasksTrigger) SetRecurrence(v string) *ListTasksResponseBodyPagingInfoTasksTrigger {
	s.Recurrence = &v
	return s
}

func (s *ListTasksResponseBodyPagingInfoTasksTrigger) SetStartTime(v string) *ListTasksResponseBodyPagingInfoTasksTrigger {
	s.StartTime = &v
	return s
}

func (s *ListTasksResponseBodyPagingInfoTasksTrigger) SetType(v string) *ListTasksResponseBodyPagingInfoTasksTrigger {
	s.Type = &v
	return s
}

func (s *ListTasksResponseBodyPagingInfoTasksTrigger) Validate() error {
	return dara.Validate(s)
}
