// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUpstreamTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListUpstreamTasksResponseBodyPagingInfo) *ListUpstreamTasksResponseBody
	GetPagingInfo() *ListUpstreamTasksResponseBodyPagingInfo
	SetRequestId(v string) *ListUpstreamTasksResponseBody
	GetRequestId() *string
}

type ListUpstreamTasksResponseBody struct {
	// The pagination information.
	PagingInfo *ListUpstreamTasksResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUpstreamTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUpstreamTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListUpstreamTasksResponseBody) GetPagingInfo() *ListUpstreamTasksResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListUpstreamTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUpstreamTasksResponseBody) SetPagingInfo(v *ListUpstreamTasksResponseBodyPagingInfo) *ListUpstreamTasksResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListUpstreamTasksResponseBody) SetRequestId(v string) *ListUpstreamTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUpstreamTasksResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListUpstreamTasksResponseBodyPagingInfo struct {
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
	// The tasks. This parameter is deprecated and replaced by the UpstreamTasks parameter.
	Tasks []*ListUpstreamTasksResponseBodyPagingInfoTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The ancestor tasks.
	UpstreamTasks []*ListUpstreamTasksResponseBodyPagingInfoUpstreamTasks `json:"UpstreamTasks,omitempty" xml:"UpstreamTasks,omitempty" type:"Repeated"`
}

func (s ListUpstreamTasksResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListUpstreamTasksResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListUpstreamTasksResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUpstreamTasksResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUpstreamTasksResponseBodyPagingInfo) GetTasks() []*ListUpstreamTasksResponseBodyPagingInfoTasks {
	return s.Tasks
}

func (s *ListUpstreamTasksResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUpstreamTasksResponseBodyPagingInfo) GetUpstreamTasks() []*ListUpstreamTasksResponseBodyPagingInfoUpstreamTasks {
	return s.UpstreamTasks
}

func (s *ListUpstreamTasksResponseBodyPagingInfo) SetPageNumber(v int32) *ListUpstreamTasksResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfo) SetPageSize(v int32) *ListUpstreamTasksResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfo) SetTasks(v []*ListUpstreamTasksResponseBodyPagingInfoTasks) *ListUpstreamTasksResponseBodyPagingInfo {
	s.Tasks = v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfo) SetTotalCount(v int32) *ListUpstreamTasksResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfo) SetUpstreamTasks(v []*ListUpstreamTasksResponseBodyPagingInfoUpstreamTasks) *ListUpstreamTasksResponseBodyPagingInfo {
	s.UpstreamTasks = v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfo) Validate() error {
	if s.Tasks != nil {
		for _, item := range s.Tasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UpstreamTasks != nil {
		for _, item := range s.UpstreamTasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUpstreamTasksResponseBodyPagingInfoTasks struct {
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
	DataSource *ListUpstreamTasksResponseBodyPagingInfoTasksDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The description of the task.
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
	// Deprecated
	//
	// The environment of the workspace. This parameter is deprecated and replaced by the EnvType parameter.
	//
	// Valid values:
	//
	// 	- Prod
	//
	// 	- Dev
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
	// 	- AllDenied: The task cannot be rerun regardless of whether it is successfully run or fails to run.
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
	RuntimeResource *ListUpstreamTasksResponseBodyPagingInfoTasksRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
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
	// The timeout period of task running. Unit: seconds.
	//
	// example:
	//
	// 3600
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The method to trigger task scheduling.
	Trigger *ListUpstreamTasksResponseBodyPagingInfoTasksTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
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

func (s ListUpstreamTasksResponseBodyPagingInfoTasks) String() string {
	return dara.Prettify(s)
}

func (s ListUpstreamTasksResponseBodyPagingInfoTasks) GoString() string {
	return s.String()
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) GetDataSource() *ListUpstreamTasksResponseBodyPagingInfoTasksDataSource {
	return s.DataSource
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) GetDescription() *string {
	return s.Description
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) GetEnvType() *string {
	return s.EnvType
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) GetId() *int64 {
	return s.Id
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) GetInstanceMode() *string {
	return s.InstanceMode
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) GetName() *string {
	return s.Name
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) GetOwner() *string {
	return s.Owner
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) GetPriority() *int32 {
	return s.Priority
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) GetRerunInterval() *int32 {
	return s.RerunInterval
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) GetRerunMode() *string {
	return s.RerunMode
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) GetRerunTimes() *int32 {
	return s.RerunTimes
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) GetRuntimeResource() *ListUpstreamTasksResponseBodyPagingInfoTasksRuntimeResource {
	return s.RuntimeResource
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) GetStepType() *string {
	return s.StepType
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) GetTimeout() *int32 {
	return s.Timeout
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) GetTrigger() *ListUpstreamTasksResponseBodyPagingInfoTasksTrigger {
	return s.Trigger
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) GetType() *string {
	return s.Type
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) SetBaselineId(v int64) *ListUpstreamTasksResponseBodyPagingInfoTasks {
	s.BaselineId = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) SetCreateTime(v int64) *ListUpstreamTasksResponseBodyPagingInfoTasks {
	s.CreateTime = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) SetCreateUser(v string) *ListUpstreamTasksResponseBodyPagingInfoTasks {
	s.CreateUser = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) SetDataSource(v *ListUpstreamTasksResponseBodyPagingInfoTasksDataSource) *ListUpstreamTasksResponseBodyPagingInfoTasks {
	s.DataSource = v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) SetDescription(v string) *ListUpstreamTasksResponseBodyPagingInfoTasks {
	s.Description = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) SetEnvType(v string) *ListUpstreamTasksResponseBodyPagingInfoTasks {
	s.EnvType = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) SetId(v int64) *ListUpstreamTasksResponseBodyPagingInfoTasks {
	s.Id = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) SetInstanceMode(v string) *ListUpstreamTasksResponseBodyPagingInfoTasks {
	s.InstanceMode = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) SetModifyTime(v int64) *ListUpstreamTasksResponseBodyPagingInfoTasks {
	s.ModifyTime = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) SetModifyUser(v string) *ListUpstreamTasksResponseBodyPagingInfoTasks {
	s.ModifyUser = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) SetName(v string) *ListUpstreamTasksResponseBodyPagingInfoTasks {
	s.Name = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) SetOwner(v string) *ListUpstreamTasksResponseBodyPagingInfoTasks {
	s.Owner = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) SetPriority(v int32) *ListUpstreamTasksResponseBodyPagingInfoTasks {
	s.Priority = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) SetProjectEnv(v string) *ListUpstreamTasksResponseBodyPagingInfoTasks {
	s.ProjectEnv = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) SetProjectId(v int64) *ListUpstreamTasksResponseBodyPagingInfoTasks {
	s.ProjectId = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) SetRerunInterval(v int32) *ListUpstreamTasksResponseBodyPagingInfoTasks {
	s.RerunInterval = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) SetRerunMode(v string) *ListUpstreamTasksResponseBodyPagingInfoTasks {
	s.RerunMode = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) SetRerunTimes(v int32) *ListUpstreamTasksResponseBodyPagingInfoTasks {
	s.RerunTimes = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) SetRuntimeResource(v *ListUpstreamTasksResponseBodyPagingInfoTasksRuntimeResource) *ListUpstreamTasksResponseBodyPagingInfoTasks {
	s.RuntimeResource = v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) SetStepType(v string) *ListUpstreamTasksResponseBodyPagingInfoTasks {
	s.StepType = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) SetTimeout(v int32) *ListUpstreamTasksResponseBodyPagingInfoTasks {
	s.Timeout = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) SetTrigger(v *ListUpstreamTasksResponseBodyPagingInfoTasksTrigger) *ListUpstreamTasksResponseBodyPagingInfoTasks {
	s.Trigger = v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) SetType(v string) *ListUpstreamTasksResponseBodyPagingInfoTasks {
	s.Type = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) SetWorkflowId(v int64) *ListUpstreamTasksResponseBodyPagingInfoTasks {
	s.WorkflowId = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasks) Validate() error {
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

type ListUpstreamTasksResponseBodyPagingInfoTasksDataSource struct {
	// The name of the data source.
	//
	// example:
	//
	// mysql_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListUpstreamTasksResponseBodyPagingInfoTasksDataSource) String() string {
	return dara.Prettify(s)
}

func (s ListUpstreamTasksResponseBodyPagingInfoTasksDataSource) GoString() string {
	return s.String()
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasksDataSource) GetName() *string {
	return s.Name
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasksDataSource) SetName(v string) *ListUpstreamTasksResponseBodyPagingInfoTasksDataSource {
	s.Name = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasksDataSource) Validate() error {
	return dara.Validate(s)
}

type ListUpstreamTasksResponseBodyPagingInfoTasksRuntimeResource struct {
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

func (s ListUpstreamTasksResponseBodyPagingInfoTasksRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s ListUpstreamTasksResponseBodyPagingInfoTasksRuntimeResource) GoString() string {
	return s.String()
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasksRuntimeResource) GetCu() *string {
	return s.Cu
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasksRuntimeResource) GetImage() *string {
	return s.Image
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasksRuntimeResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasksRuntimeResource) SetCu(v string) *ListUpstreamTasksResponseBodyPagingInfoTasksRuntimeResource {
	s.Cu = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasksRuntimeResource) SetImage(v string) *ListUpstreamTasksResponseBodyPagingInfoTasksRuntimeResource {
	s.Image = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasksRuntimeResource) SetResourceGroupId(v string) *ListUpstreamTasksResponseBodyPagingInfoTasksRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasksRuntimeResource) Validate() error {
	return dara.Validate(s)
}

type ListUpstreamTasksResponseBodyPagingInfoTasksTrigger struct {
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
	// The time zone.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
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

func (s ListUpstreamTasksResponseBodyPagingInfoTasksTrigger) String() string {
	return dara.Prettify(s)
}

func (s ListUpstreamTasksResponseBodyPagingInfoTasksTrigger) GoString() string {
	return s.String()
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasksTrigger) GetCron() *string {
	return s.Cron
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasksTrigger) GetEndTime() *string {
	return s.EndTime
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasksTrigger) GetRecurrence() *string {
	return s.Recurrence
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasksTrigger) GetStartTime() *string {
	return s.StartTime
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasksTrigger) GetTimezone() *string {
	return s.Timezone
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasksTrigger) GetType() *string {
	return s.Type
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasksTrigger) SetCron(v string) *ListUpstreamTasksResponseBodyPagingInfoTasksTrigger {
	s.Cron = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasksTrigger) SetEndTime(v string) *ListUpstreamTasksResponseBodyPagingInfoTasksTrigger {
	s.EndTime = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasksTrigger) SetRecurrence(v string) *ListUpstreamTasksResponseBodyPagingInfoTasksTrigger {
	s.Recurrence = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasksTrigger) SetStartTime(v string) *ListUpstreamTasksResponseBodyPagingInfoTasksTrigger {
	s.StartTime = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasksTrigger) SetTimezone(v string) *ListUpstreamTasksResponseBodyPagingInfoTasksTrigger {
	s.Timezone = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasksTrigger) SetType(v string) *ListUpstreamTasksResponseBodyPagingInfoTasksTrigger {
	s.Type = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoTasksTrigger) Validate() error {
	return dara.Validate(s)
}

type ListUpstreamTasksResponseBodyPagingInfoUpstreamTasks struct {
	// The scheduling dependency type. Valid values:
	//
	// 	- Normal: same-cycle scheduling dependency
	//
	// 	- CrossCycle: cross-cycle scheduling dependency
	//
	// example:
	//
	// Normal
	DependencyType *string `json:"DependencyType,omitempty" xml:"DependencyType,omitempty"`
	// The information about the task.
	Task *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Struct"`
}

func (s ListUpstreamTasksResponseBodyPagingInfoUpstreamTasks) String() string {
	return dara.Prettify(s)
}

func (s ListUpstreamTasksResponseBodyPagingInfoUpstreamTasks) GoString() string {
	return s.String()
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasks) GetDependencyType() *string {
	return s.DependencyType
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasks) GetTask() *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask {
	return s.Task
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasks) SetDependencyType(v string) *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasks {
	s.DependencyType = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasks) SetTask(v *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasks {
	s.Task = v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasks) Validate() error {
	if s.Task != nil {
		if err := s.Task.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask struct {
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
	DataSource *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The description of the task.
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
	// The priority of the task. Valid values: 1 to 8.
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
	// The rerun interval. Unit: seconds.
	//
	// example:
	//
	// 60
	RerunInterval *int32 `json:"RerunInterval,omitempty" xml:"RerunInterval,omitempty"`
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
	// The number of times that the task is rerun. This parameter takes effect only if the RerunMode parameter is set to AllAllowed or FailureAllowed.
	//
	// example:
	//
	// 3
	RerunTimes *int32 `json:"RerunTimes,omitempty" xml:"RerunTimes,omitempty"`
	// The configurations of the runtime environment, such as the resource group information.
	RuntimeResource *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// The timeout period of task running. Unit: seconds.
	//
	// example:
	//
	// 3600
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The trigger method.
	Trigger *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
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

func (s ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) String() string {
	return dara.Prettify(s)
}

func (s ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) GoString() string {
	return s.String()
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) GetDataSource() *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskDataSource {
	return s.DataSource
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) GetDescription() *string {
	return s.Description
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) GetEnvType() *string {
	return s.EnvType
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) GetId() *int64 {
	return s.Id
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) GetInstanceMode() *string {
	return s.InstanceMode
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) GetName() *string {
	return s.Name
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) GetOwner() *string {
	return s.Owner
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) GetPriority() *int32 {
	return s.Priority
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) GetRerunInterval() *int32 {
	return s.RerunInterval
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) GetRerunMode() *string {
	return s.RerunMode
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) GetRerunTimes() *int32 {
	return s.RerunTimes
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) GetRuntimeResource() *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskRuntimeResource {
	return s.RuntimeResource
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) GetTimeout() *int32 {
	return s.Timeout
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) GetTrigger() *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskTrigger {
	return s.Trigger
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) GetType() *string {
	return s.Type
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) SetBaselineId(v int64) *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask {
	s.BaselineId = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) SetCreateTime(v int64) *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask {
	s.CreateTime = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) SetCreateUser(v string) *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask {
	s.CreateUser = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) SetDataSource(v *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskDataSource) *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask {
	s.DataSource = v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) SetDescription(v string) *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask {
	s.Description = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) SetEnvType(v string) *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask {
	s.EnvType = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) SetId(v int64) *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask {
	s.Id = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) SetInstanceMode(v string) *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask {
	s.InstanceMode = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) SetModifyTime(v int64) *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask {
	s.ModifyTime = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) SetModifyUser(v string) *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask {
	s.ModifyUser = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) SetName(v string) *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask {
	s.Name = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) SetOwner(v string) *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask {
	s.Owner = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) SetPriority(v int32) *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask {
	s.Priority = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) SetProjectId(v int64) *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask {
	s.ProjectId = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) SetRerunInterval(v int32) *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask {
	s.RerunInterval = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) SetRerunMode(v string) *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask {
	s.RerunMode = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) SetRerunTimes(v int32) *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask {
	s.RerunTimes = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) SetRuntimeResource(v *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskRuntimeResource) *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask {
	s.RuntimeResource = v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) SetTimeout(v int32) *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask {
	s.Timeout = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) SetTrigger(v *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskTrigger) *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask {
	s.Trigger = v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) SetType(v string) *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask {
	s.Type = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) SetWorkflowId(v int64) *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask {
	s.WorkflowId = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTask) Validate() error {
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

type ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskDataSource struct {
	// The name of the data source.
	//
	// example:
	//
	// mysql_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskDataSource) String() string {
	return dara.Prettify(s)
}

func (s ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskDataSource) GoString() string {
	return s.String()
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskDataSource) GetName() *string {
	return s.Name
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskDataSource) SetName(v string) *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskDataSource {
	s.Name = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskDataSource) Validate() error {
	return dara.Validate(s)
}

type ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskRuntimeResource struct {
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

func (s ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskRuntimeResource) GoString() string {
	return s.String()
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskRuntimeResource) GetCu() *string {
	return s.Cu
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskRuntimeResource) GetImage() *string {
	return s.Image
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskRuntimeResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskRuntimeResource) SetCu(v string) *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskRuntimeResource {
	s.Cu = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskRuntimeResource) SetImage(v string) *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskRuntimeResource {
	s.Image = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskRuntimeResource) SetResourceGroupId(v string) *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskRuntimeResource) Validate() error {
	return dara.Validate(s)
}

type ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskTrigger struct {
	// The CRON expression. This parameter takes effect only if the Type parameter is set to Scheduler.
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
	// The running mode of the task after it is triggered. This parameter takes effect only if the Type parameter is set to Scheduler. Valid values:
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
	// The time zone.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	// The trigger type. Valid values:
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

func (s ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskTrigger) String() string {
	return dara.Prettify(s)
}

func (s ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskTrigger) GoString() string {
	return s.String()
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskTrigger) GetCron() *string {
	return s.Cron
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskTrigger) GetEndTime() *string {
	return s.EndTime
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskTrigger) GetRecurrence() *string {
	return s.Recurrence
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskTrigger) GetStartTime() *string {
	return s.StartTime
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskTrigger) GetTimezone() *string {
	return s.Timezone
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskTrigger) GetType() *string {
	return s.Type
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskTrigger) SetCron(v string) *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskTrigger {
	s.Cron = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskTrigger) SetEndTime(v string) *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskTrigger {
	s.EndTime = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskTrigger) SetRecurrence(v string) *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskTrigger {
	s.Recurrence = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskTrigger) SetStartTime(v string) *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskTrigger {
	s.StartTime = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskTrigger) SetTimezone(v string) *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskTrigger {
	s.Timezone = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskTrigger) SetType(v string) *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskTrigger {
	s.Type = &v
	return s
}

func (s *ListUpstreamTasksResponseBodyPagingInfoUpstreamTasksTaskTrigger) Validate() error {
	return dara.Validate(s)
}
