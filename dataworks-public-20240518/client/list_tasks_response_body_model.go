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
	// The request ID. You can use this ID to locate logs and troubleshoot issues.
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
	// The list of tasks.
	Tasks []*ListTasksResponseBodyPagingInfoTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// The total number of entries.
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
	// The value is a 13-digit number, such as `1710239005403`.
	//
	// example:
	//
	// 1710239005403
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The account ID of the user who created the task.
	//
	// example:
	//
	// 1000
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The associated data source information.
	DataSource *ListTasksResponseBodyPagingInfoTasksDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The unique identifier of the task.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The instance generation mode. Valid values:
	//
	// - T+1: Generated on the next day.
	//
	// - Immediately: Generated immediately.
	//
	// example:
	//
	// T+1
	InstanceMode *string `json:"InstanceMode,omitempty" xml:"InstanceMode,omitempty"`
	// The modification time.
	//
	// The value is a 13-digit number, such as `1710239005403`.
	//
	// example:
	//
	// 1710239005403
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The account ID of the user who modified the task.
	//
	// example:
	//
	// 1000
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// The name.
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
	// The task priority. Valid values: 1 to 8. A larger value indicates a higher priority. Default value: 1.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The project environment.
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
	// The retry time interval, in seconds.
	//
	// example:
	//
	// 60
	RerunInterval *int32 `json:"RerunInterval,omitempty" xml:"RerunInterval,omitempty"`
	// The rerun configuration of the task.
	//
	// example:
	//
	// AllAllowed
	RerunMode *string `json:"RerunMode,omitempty" xml:"RerunMode,omitempty"`
	// The number of retries. This parameter takes effect only when the task is configured to allow reruns.
	//
	// example:
	//
	// 3
	RerunTimes *int32 `json:"RerunTimes,omitempty" xml:"RerunTimes,omitempty"`
	// The runtime environment configuration, such as the resource group information.
	RuntimeResource *ListTasksResponseBodyPagingInfoTasksRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// The list of script parameters.
	//
	// example:
	//
	// para1=$bizdate para2=$[yyyymmdd]
	ScriptParameters *string `json:"ScriptParameters,omitempty" xml:"ScriptParameters,omitempty"`
	// The task execution timeout period, in seconds.
	//
	// example:
	//
	// 3600
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The trigger type of the task.
	Trigger *ListTasksResponseBodyPagingInfoTasksTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
	// The task type.
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
	// The number of compute units (CUs) consumed by the task run configuration.
	//
	// example:
	//
	// 0.25
	Cu *string `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The ID of the image configured for the task runtime.
	//
	// example:
	//
	// i-xxxxxx
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The identifier of the scheduling resource group configured for the task runtime.
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
	// The cron expression. This parameter takes effect only when type is set to Scheduler.
	//
	// example:
	//
	// 00 00 00 	- 	- ?
	Cron *string `json:"Cron,omitempty" xml:"Cron,omitempty"`
	// The expiration time for periodic triggers. This parameter takes effect only when type is set to Scheduler.
	//
	// The format is `yyyy-MM-dd HH:mm:ss`, such as `9999-01-01 00:00:00`. The example does not include a time zone identifier.
	//
	// example:
	//
	// 9999-01-01 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The running mode upon trigger. This parameter takes effect only when type is set to Scheduler.
	//
	// example:
	//
	// Normal
	Recurrence *string `json:"Recurrence,omitempty" xml:"Recurrence,omitempty"`
	// The effective time for periodic triggers. This parameter takes effect only when type is set to Scheduler.
	//
	// The format is `yyyy-MM-dd HH:mm:ss`, such as `1970-01-01 00:00:00`. The example does not include a time zone identifier.
	//
	// example:
	//
	// 1970-01-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The trigger type.
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
