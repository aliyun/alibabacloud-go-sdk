// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDownstreamTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListDownstreamTasksResponseBodyPagingInfo) *ListDownstreamTasksResponseBody
	GetPagingInfo() *ListDownstreamTasksResponseBodyPagingInfo
	SetRequestId(v string) *ListDownstreamTasksResponseBody
	GetRequestId() *string
}

type ListDownstreamTasksResponseBody struct {
	// The pagination information.
	PagingInfo *ListDownstreamTasksResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDownstreamTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDownstreamTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListDownstreamTasksResponseBody) GetPagingInfo() *ListDownstreamTasksResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListDownstreamTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDownstreamTasksResponseBody) SetPagingInfo(v *ListDownstreamTasksResponseBodyPagingInfo) *ListDownstreamTasksResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDownstreamTasksResponseBody) SetRequestId(v string) *ListDownstreamTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDownstreamTasksResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDownstreamTasksResponseBodyPagingInfo struct {
	// The descendant tasks.
	DownstreamTasks []*ListDownstreamTasksResponseBodyPagingInfoDownstreamTasks `json:"DownstreamTasks,omitempty" xml:"DownstreamTasks,omitempty" type:"Repeated"`
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
	// The tasks. This parameter is deprecated and replaced by the DownstreamTasks parameter.
	Tasks []*ListDownstreamTasksResponseBodyPagingInfoTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDownstreamTasksResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDownstreamTasksResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListDownstreamTasksResponseBodyPagingInfo) GetDownstreamTasks() []*ListDownstreamTasksResponseBodyPagingInfoDownstreamTasks {
	return s.DownstreamTasks
}

func (s *ListDownstreamTasksResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDownstreamTasksResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDownstreamTasksResponseBodyPagingInfo) GetTasks() []*ListDownstreamTasksResponseBodyPagingInfoTasks {
	return s.Tasks
}

func (s *ListDownstreamTasksResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDownstreamTasksResponseBodyPagingInfo) SetDownstreamTasks(v []*ListDownstreamTasksResponseBodyPagingInfoDownstreamTasks) *ListDownstreamTasksResponseBodyPagingInfo {
	s.DownstreamTasks = v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfo) SetPageNumber(v int32) *ListDownstreamTasksResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfo) SetPageSize(v int32) *ListDownstreamTasksResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfo) SetTasks(v []*ListDownstreamTasksResponseBodyPagingInfoTasks) *ListDownstreamTasksResponseBodyPagingInfo {
	s.Tasks = v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfo) SetTotalCount(v int32) *ListDownstreamTasksResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfo) Validate() error {
	return dara.Validate(s)
}

type ListDownstreamTasksResponseBodyPagingInfoDownstreamTasks struct {
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
	Task *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Struct"`
}

func (s ListDownstreamTasksResponseBodyPagingInfoDownstreamTasks) String() string {
	return dara.Prettify(s)
}

func (s ListDownstreamTasksResponseBodyPagingInfoDownstreamTasks) GoString() string {
	return s.String()
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasks) GetDependencyType() *string {
	return s.DependencyType
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasks) GetTask() *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask {
	return s.Task
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasks) SetDependencyType(v string) *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasks {
	s.DependencyType = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasks) SetTask(v *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasks {
	s.Task = v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasks) Validate() error {
	return dara.Validate(s)
}

type ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask struct {
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
	DataSource *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
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
	RuntimeResource *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// The timeout period of task running. Unit: seconds.
	//
	// example:
	//
	// 3600
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The trigger method.
	Trigger *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
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

func (s ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) String() string {
	return dara.Prettify(s)
}

func (s ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) GoString() string {
	return s.String()
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) GetDataSource() *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskDataSource {
	return s.DataSource
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) GetDescription() *string {
	return s.Description
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) GetEnvType() *string {
	return s.EnvType
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) GetId() *int64 {
	return s.Id
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) GetInstanceMode() *string {
	return s.InstanceMode
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) GetName() *string {
	return s.Name
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) GetOwner() *string {
	return s.Owner
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) GetPriority() *int32 {
	return s.Priority
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) GetRerunInterval() *int32 {
	return s.RerunInterval
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) GetRerunMode() *string {
	return s.RerunMode
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) GetRerunTimes() *int32 {
	return s.RerunTimes
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) GetRuntimeResource() *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskRuntimeResource {
	return s.RuntimeResource
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) GetTimeout() *int32 {
	return s.Timeout
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) GetTrigger() *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskTrigger {
	return s.Trigger
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) GetType() *string {
	return s.Type
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) SetBaselineId(v int64) *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask {
	s.BaselineId = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) SetCreateTime(v int64) *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask {
	s.CreateTime = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) SetCreateUser(v string) *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask {
	s.CreateUser = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) SetDataSource(v *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskDataSource) *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask {
	s.DataSource = v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) SetDescription(v string) *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask {
	s.Description = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) SetEnvType(v string) *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask {
	s.EnvType = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) SetId(v int64) *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask {
	s.Id = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) SetInstanceMode(v string) *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask {
	s.InstanceMode = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) SetModifyTime(v int64) *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask {
	s.ModifyTime = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) SetModifyUser(v string) *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask {
	s.ModifyUser = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) SetName(v string) *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask {
	s.Name = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) SetOwner(v string) *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask {
	s.Owner = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) SetPriority(v int32) *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask {
	s.Priority = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) SetProjectId(v int64) *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask {
	s.ProjectId = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) SetRerunInterval(v int32) *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask {
	s.RerunInterval = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) SetRerunMode(v string) *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask {
	s.RerunMode = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) SetRerunTimes(v int32) *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask {
	s.RerunTimes = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) SetRuntimeResource(v *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskRuntimeResource) *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask {
	s.RuntimeResource = v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) SetTimeout(v int32) *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask {
	s.Timeout = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) SetTrigger(v *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskTrigger) *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask {
	s.Trigger = v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) SetType(v string) *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask {
	s.Type = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) SetWorkflowId(v int64) *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask {
	s.WorkflowId = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTask) Validate() error {
	return dara.Validate(s)
}

type ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskDataSource struct {
	// The name of the data source.
	//
	// example:
	//
	// mysql_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskDataSource) String() string {
	return dara.Prettify(s)
}

func (s ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskDataSource) GoString() string {
	return s.String()
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskDataSource) GetName() *string {
	return s.Name
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskDataSource) SetName(v string) *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskDataSource {
	s.Name = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskDataSource) Validate() error {
	return dara.Validate(s)
}

type ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskRuntimeResource struct {
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

func (s ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskRuntimeResource) GoString() string {
	return s.String()
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskRuntimeResource) GetCu() *string {
	return s.Cu
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskRuntimeResource) GetImage() *string {
	return s.Image
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskRuntimeResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskRuntimeResource) SetCu(v string) *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskRuntimeResource {
	s.Cu = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskRuntimeResource) SetImage(v string) *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskRuntimeResource {
	s.Image = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskRuntimeResource) SetResourceGroupId(v string) *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskRuntimeResource) Validate() error {
	return dara.Validate(s)
}

type ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskTrigger struct {
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

func (s ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskTrigger) String() string {
	return dara.Prettify(s)
}

func (s ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskTrigger) GoString() string {
	return s.String()
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskTrigger) GetCron() *string {
	return s.Cron
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskTrigger) GetEndTime() *string {
	return s.EndTime
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskTrigger) GetRecurrence() *string {
	return s.Recurrence
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskTrigger) GetStartTime() *string {
	return s.StartTime
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskTrigger) GetTimezone() *string {
	return s.Timezone
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskTrigger) GetType() *string {
	return s.Type
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskTrigger) SetCron(v string) *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskTrigger {
	s.Cron = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskTrigger) SetEndTime(v string) *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskTrigger {
	s.EndTime = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskTrigger) SetRecurrence(v string) *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskTrigger {
	s.Recurrence = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskTrigger) SetStartTime(v string) *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskTrigger {
	s.StartTime = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskTrigger) SetTimezone(v string) *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskTrigger {
	s.Timezone = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskTrigger) SetType(v string) *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskTrigger {
	s.Type = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoDownstreamTasksTaskTrigger) Validate() error {
	return dara.Validate(s)
}

type ListDownstreamTasksResponseBodyPagingInfoTasks struct {
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
	DataSource *ListDownstreamTasksResponseBodyPagingInfoTasksDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
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
	RuntimeResource *ListDownstreamTasksResponseBodyPagingInfoTasksRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
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
	Trigger *ListDownstreamTasksResponseBodyPagingInfoTasksTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
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

func (s ListDownstreamTasksResponseBodyPagingInfoTasks) String() string {
	return dara.Prettify(s)
}

func (s ListDownstreamTasksResponseBodyPagingInfoTasks) GoString() string {
	return s.String()
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) GetDataSource() *ListDownstreamTasksResponseBodyPagingInfoTasksDataSource {
	return s.DataSource
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) GetDescription() *string {
	return s.Description
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) GetEnvType() *string {
	return s.EnvType
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) GetId() *int64 {
	return s.Id
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) GetInstanceMode() *string {
	return s.InstanceMode
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) GetName() *string {
	return s.Name
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) GetOwner() *string {
	return s.Owner
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) GetPriority() *int32 {
	return s.Priority
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) GetRerunInterval() *int32 {
	return s.RerunInterval
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) GetRerunMode() *string {
	return s.RerunMode
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) GetRerunTimes() *int32 {
	return s.RerunTimes
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) GetRuntimeResource() *ListDownstreamTasksResponseBodyPagingInfoTasksRuntimeResource {
	return s.RuntimeResource
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) GetStepType() *string {
	return s.StepType
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) GetTimeout() *int32 {
	return s.Timeout
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) GetTrigger() *ListDownstreamTasksResponseBodyPagingInfoTasksTrigger {
	return s.Trigger
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) GetType() *string {
	return s.Type
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) SetBaselineId(v int64) *ListDownstreamTasksResponseBodyPagingInfoTasks {
	s.BaselineId = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) SetCreateTime(v int64) *ListDownstreamTasksResponseBodyPagingInfoTasks {
	s.CreateTime = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) SetCreateUser(v string) *ListDownstreamTasksResponseBodyPagingInfoTasks {
	s.CreateUser = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) SetDataSource(v *ListDownstreamTasksResponseBodyPagingInfoTasksDataSource) *ListDownstreamTasksResponseBodyPagingInfoTasks {
	s.DataSource = v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) SetDescription(v string) *ListDownstreamTasksResponseBodyPagingInfoTasks {
	s.Description = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) SetEnvType(v string) *ListDownstreamTasksResponseBodyPagingInfoTasks {
	s.EnvType = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) SetId(v int64) *ListDownstreamTasksResponseBodyPagingInfoTasks {
	s.Id = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) SetInstanceMode(v string) *ListDownstreamTasksResponseBodyPagingInfoTasks {
	s.InstanceMode = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) SetModifyTime(v int64) *ListDownstreamTasksResponseBodyPagingInfoTasks {
	s.ModifyTime = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) SetModifyUser(v string) *ListDownstreamTasksResponseBodyPagingInfoTasks {
	s.ModifyUser = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) SetName(v string) *ListDownstreamTasksResponseBodyPagingInfoTasks {
	s.Name = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) SetOwner(v string) *ListDownstreamTasksResponseBodyPagingInfoTasks {
	s.Owner = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) SetPriority(v int32) *ListDownstreamTasksResponseBodyPagingInfoTasks {
	s.Priority = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) SetProjectEnv(v string) *ListDownstreamTasksResponseBodyPagingInfoTasks {
	s.ProjectEnv = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) SetProjectId(v int64) *ListDownstreamTasksResponseBodyPagingInfoTasks {
	s.ProjectId = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) SetRerunInterval(v int32) *ListDownstreamTasksResponseBodyPagingInfoTasks {
	s.RerunInterval = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) SetRerunMode(v string) *ListDownstreamTasksResponseBodyPagingInfoTasks {
	s.RerunMode = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) SetRerunTimes(v int32) *ListDownstreamTasksResponseBodyPagingInfoTasks {
	s.RerunTimes = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) SetRuntimeResource(v *ListDownstreamTasksResponseBodyPagingInfoTasksRuntimeResource) *ListDownstreamTasksResponseBodyPagingInfoTasks {
	s.RuntimeResource = v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) SetStepType(v string) *ListDownstreamTasksResponseBodyPagingInfoTasks {
	s.StepType = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) SetTimeout(v int32) *ListDownstreamTasksResponseBodyPagingInfoTasks {
	s.Timeout = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) SetTrigger(v *ListDownstreamTasksResponseBodyPagingInfoTasksTrigger) *ListDownstreamTasksResponseBodyPagingInfoTasks {
	s.Trigger = v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) SetType(v string) *ListDownstreamTasksResponseBodyPagingInfoTasks {
	s.Type = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) SetWorkflowId(v int64) *ListDownstreamTasksResponseBodyPagingInfoTasks {
	s.WorkflowId = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasks) Validate() error {
	return dara.Validate(s)
}

type ListDownstreamTasksResponseBodyPagingInfoTasksDataSource struct {
	// The name of the data source.
	//
	// example:
	//
	// mysql_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListDownstreamTasksResponseBodyPagingInfoTasksDataSource) String() string {
	return dara.Prettify(s)
}

func (s ListDownstreamTasksResponseBodyPagingInfoTasksDataSource) GoString() string {
	return s.String()
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasksDataSource) GetName() *string {
	return s.Name
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasksDataSource) SetName(v string) *ListDownstreamTasksResponseBodyPagingInfoTasksDataSource {
	s.Name = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasksDataSource) Validate() error {
	return dara.Validate(s)
}

type ListDownstreamTasksResponseBodyPagingInfoTasksRuntimeResource struct {
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

func (s ListDownstreamTasksResponseBodyPagingInfoTasksRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s ListDownstreamTasksResponseBodyPagingInfoTasksRuntimeResource) GoString() string {
	return s.String()
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasksRuntimeResource) GetCu() *string {
	return s.Cu
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasksRuntimeResource) GetImage() *string {
	return s.Image
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasksRuntimeResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasksRuntimeResource) SetCu(v string) *ListDownstreamTasksResponseBodyPagingInfoTasksRuntimeResource {
	s.Cu = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasksRuntimeResource) SetImage(v string) *ListDownstreamTasksResponseBodyPagingInfoTasksRuntimeResource {
	s.Image = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasksRuntimeResource) SetResourceGroupId(v string) *ListDownstreamTasksResponseBodyPagingInfoTasksRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasksRuntimeResource) Validate() error {
	return dara.Validate(s)
}

type ListDownstreamTasksResponseBodyPagingInfoTasksTrigger struct {
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

func (s ListDownstreamTasksResponseBodyPagingInfoTasksTrigger) String() string {
	return dara.Prettify(s)
}

func (s ListDownstreamTasksResponseBodyPagingInfoTasksTrigger) GoString() string {
	return s.String()
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasksTrigger) GetCron() *string {
	return s.Cron
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasksTrigger) GetEndTime() *string {
	return s.EndTime
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasksTrigger) GetRecurrence() *string {
	return s.Recurrence
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasksTrigger) GetStartTime() *string {
	return s.StartTime
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasksTrigger) GetTimezone() *string {
	return s.Timezone
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasksTrigger) GetType() *string {
	return s.Type
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasksTrigger) SetCron(v string) *ListDownstreamTasksResponseBodyPagingInfoTasksTrigger {
	s.Cron = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasksTrigger) SetEndTime(v string) *ListDownstreamTasksResponseBodyPagingInfoTasksTrigger {
	s.EndTime = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasksTrigger) SetRecurrence(v string) *ListDownstreamTasksResponseBodyPagingInfoTasksTrigger {
	s.Recurrence = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasksTrigger) SetStartTime(v string) *ListDownstreamTasksResponseBodyPagingInfoTasksTrigger {
	s.StartTime = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasksTrigger) SetTimezone(v string) *ListDownstreamTasksResponseBodyPagingInfoTasksTrigger {
	s.Timezone = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasksTrigger) SetType(v string) *ListDownstreamTasksResponseBodyPagingInfoTasksTrigger {
	s.Type = &v
	return s
}

func (s *ListDownstreamTasksResponseBodyPagingInfoTasksTrigger) Validate() error {
	return dara.Validate(s)
}
