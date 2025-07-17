// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *BatchUpdateTasksRequest
	GetComment() *string
	SetTasks(v []*BatchUpdateTasksRequestTasks) *BatchUpdateTasksRequest
	GetTasks() []*BatchUpdateTasksRequestTasks
}

type BatchUpdateTasksRequest struct {
	// The remarks.
	//
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The tasks.
	Tasks []*BatchUpdateTasksRequestTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s BatchUpdateTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateTasksRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateTasksRequest) GetComment() *string {
	return s.Comment
}

func (s *BatchUpdateTasksRequest) GetTasks() []*BatchUpdateTasksRequestTasks {
	return s.Tasks
}

func (s *BatchUpdateTasksRequest) SetComment(v string) *BatchUpdateTasksRequest {
	s.Comment = &v
	return s
}

func (s *BatchUpdateTasksRequest) SetTasks(v []*BatchUpdateTasksRequestTasks) *BatchUpdateTasksRequest {
	s.Tasks = v
	return s
}

func (s *BatchUpdateTasksRequest) Validate() error {
	return dara.Validate(s)
}

type BatchUpdateTasksRequestTasks struct {
	// The information about the associated data source.
	DataSource *BatchUpdateTasksRequestTasksDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The environment of the workspace. Valid values:
	//
	// 	- Prod: production environment
	//
	// 	- Dev: development environment
	//
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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
	RuntimeResource *BatchUpdateTasksRequestTasksRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// The tags.
	Tags []*BatchUpdateTasksRequestTasksTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The timeout period of task running. Unit: seconds.
	//
	// example:
	//
	// 3600
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The trigger method.
	Trigger *BatchUpdateTasksRequestTasksTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
}

func (s BatchUpdateTasksRequestTasks) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateTasksRequestTasks) GoString() string {
	return s.String()
}

func (s *BatchUpdateTasksRequestTasks) GetDataSource() *BatchUpdateTasksRequestTasksDataSource {
	return s.DataSource
}

func (s *BatchUpdateTasksRequestTasks) GetDescription() *string {
	return s.Description
}

func (s *BatchUpdateTasksRequestTasks) GetEnvType() *string {
	return s.EnvType
}

func (s *BatchUpdateTasksRequestTasks) GetId() *int64 {
	return s.Id
}

func (s *BatchUpdateTasksRequestTasks) GetName() *string {
	return s.Name
}

func (s *BatchUpdateTasksRequestTasks) GetOwner() *string {
	return s.Owner
}

func (s *BatchUpdateTasksRequestTasks) GetRerunInterval() *int32 {
	return s.RerunInterval
}

func (s *BatchUpdateTasksRequestTasks) GetRerunMode() *string {
	return s.RerunMode
}

func (s *BatchUpdateTasksRequestTasks) GetRerunTimes() *int32 {
	return s.RerunTimes
}

func (s *BatchUpdateTasksRequestTasks) GetRuntimeResource() *BatchUpdateTasksRequestTasksRuntimeResource {
	return s.RuntimeResource
}

func (s *BatchUpdateTasksRequestTasks) GetTags() []*BatchUpdateTasksRequestTasksTags {
	return s.Tags
}

func (s *BatchUpdateTasksRequestTasks) GetTimeout() *int32 {
	return s.Timeout
}

func (s *BatchUpdateTasksRequestTasks) GetTrigger() *BatchUpdateTasksRequestTasksTrigger {
	return s.Trigger
}

func (s *BatchUpdateTasksRequestTasks) SetDataSource(v *BatchUpdateTasksRequestTasksDataSource) *BatchUpdateTasksRequestTasks {
	s.DataSource = v
	return s
}

func (s *BatchUpdateTasksRequestTasks) SetDescription(v string) *BatchUpdateTasksRequestTasks {
	s.Description = &v
	return s
}

func (s *BatchUpdateTasksRequestTasks) SetEnvType(v string) *BatchUpdateTasksRequestTasks {
	s.EnvType = &v
	return s
}

func (s *BatchUpdateTasksRequestTasks) SetId(v int64) *BatchUpdateTasksRequestTasks {
	s.Id = &v
	return s
}

func (s *BatchUpdateTasksRequestTasks) SetName(v string) *BatchUpdateTasksRequestTasks {
	s.Name = &v
	return s
}

func (s *BatchUpdateTasksRequestTasks) SetOwner(v string) *BatchUpdateTasksRequestTasks {
	s.Owner = &v
	return s
}

func (s *BatchUpdateTasksRequestTasks) SetRerunInterval(v int32) *BatchUpdateTasksRequestTasks {
	s.RerunInterval = &v
	return s
}

func (s *BatchUpdateTasksRequestTasks) SetRerunMode(v string) *BatchUpdateTasksRequestTasks {
	s.RerunMode = &v
	return s
}

func (s *BatchUpdateTasksRequestTasks) SetRerunTimes(v int32) *BatchUpdateTasksRequestTasks {
	s.RerunTimes = &v
	return s
}

func (s *BatchUpdateTasksRequestTasks) SetRuntimeResource(v *BatchUpdateTasksRequestTasksRuntimeResource) *BatchUpdateTasksRequestTasks {
	s.RuntimeResource = v
	return s
}

func (s *BatchUpdateTasksRequestTasks) SetTags(v []*BatchUpdateTasksRequestTasksTags) *BatchUpdateTasksRequestTasks {
	s.Tags = v
	return s
}

func (s *BatchUpdateTasksRequestTasks) SetTimeout(v int32) *BatchUpdateTasksRequestTasks {
	s.Timeout = &v
	return s
}

func (s *BatchUpdateTasksRequestTasks) SetTrigger(v *BatchUpdateTasksRequestTasksTrigger) *BatchUpdateTasksRequestTasks {
	s.Trigger = v
	return s
}

func (s *BatchUpdateTasksRequestTasks) Validate() error {
	return dara.Validate(s)
}

type BatchUpdateTasksRequestTasksDataSource struct {
	// The name of the data source.
	//
	// example:
	//
	// odps_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s BatchUpdateTasksRequestTasksDataSource) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateTasksRequestTasksDataSource) GoString() string {
	return s.String()
}

func (s *BatchUpdateTasksRequestTasksDataSource) GetName() *string {
	return s.Name
}

func (s *BatchUpdateTasksRequestTasksDataSource) SetName(v string) *BatchUpdateTasksRequestTasksDataSource {
	s.Name = &v
	return s
}

func (s *BatchUpdateTasksRequestTasksDataSource) Validate() error {
	return dara.Validate(s)
}

type BatchUpdateTasksRequestTasksRuntimeResource struct {
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

func (s BatchUpdateTasksRequestTasksRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateTasksRequestTasksRuntimeResource) GoString() string {
	return s.String()
}

func (s *BatchUpdateTasksRequestTasksRuntimeResource) GetCu() *string {
	return s.Cu
}

func (s *BatchUpdateTasksRequestTasksRuntimeResource) GetImage() *string {
	return s.Image
}

func (s *BatchUpdateTasksRequestTasksRuntimeResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *BatchUpdateTasksRequestTasksRuntimeResource) SetCu(v string) *BatchUpdateTasksRequestTasksRuntimeResource {
	s.Cu = &v
	return s
}

func (s *BatchUpdateTasksRequestTasksRuntimeResource) SetImage(v string) *BatchUpdateTasksRequestTasksRuntimeResource {
	s.Image = &v
	return s
}

func (s *BatchUpdateTasksRequestTasksRuntimeResource) SetResourceGroupId(v string) *BatchUpdateTasksRequestTasksRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

func (s *BatchUpdateTasksRequestTasksRuntimeResource) Validate() error {
	return dara.Validate(s)
}

type BatchUpdateTasksRequestTasksTags struct {
	// The tag key.
	//
	// This parameter is required.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s BatchUpdateTasksRequestTasksTags) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateTasksRequestTasksTags) GoString() string {
	return s.String()
}

func (s *BatchUpdateTasksRequestTasksTags) GetKey() *string {
	return s.Key
}

func (s *BatchUpdateTasksRequestTasksTags) GetValue() *string {
	return s.Value
}

func (s *BatchUpdateTasksRequestTasksTags) SetKey(v string) *BatchUpdateTasksRequestTasksTags {
	s.Key = &v
	return s
}

func (s *BatchUpdateTasksRequestTasksTags) SetValue(v string) *BatchUpdateTasksRequestTasksTags {
	s.Value = &v
	return s
}

func (s *BatchUpdateTasksRequestTasksTags) Validate() error {
	return dara.Validate(s)
}

type BatchUpdateTasksRequestTasksTrigger struct {
	// The CRON expression. This parameter takes effect only if the Type parameter is set to Scheduler.
	//
	// example:
	//
	// 00 00 00 	- 	- ?
	Cron *string `json:"Cron,omitempty" xml:"Cron,omitempty"`
	// The end time of the time range during which the task is periodically scheduled. This parameter takes effect only if the Type parameter is set to Scheduler. The value of this parameter is in the `yyyy-mm-dd hh:mm:ss`.
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
	// The start time of the time range during which the task is periodically scheduled. This parameter takes effect only if the Type parameter is set to Scheduler. The value of this parameter is in the `yyyy-mm-dd hh:mm:ss`.
	//
	// example:
	//
	// 1970-01-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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

func (s BatchUpdateTasksRequestTasksTrigger) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateTasksRequestTasksTrigger) GoString() string {
	return s.String()
}

func (s *BatchUpdateTasksRequestTasksTrigger) GetCron() *string {
	return s.Cron
}

func (s *BatchUpdateTasksRequestTasksTrigger) GetEndTime() *string {
	return s.EndTime
}

func (s *BatchUpdateTasksRequestTasksTrigger) GetRecurrence() *string {
	return s.Recurrence
}

func (s *BatchUpdateTasksRequestTasksTrigger) GetStartTime() *string {
	return s.StartTime
}

func (s *BatchUpdateTasksRequestTasksTrigger) GetType() *string {
	return s.Type
}

func (s *BatchUpdateTasksRequestTasksTrigger) SetCron(v string) *BatchUpdateTasksRequestTasksTrigger {
	s.Cron = &v
	return s
}

func (s *BatchUpdateTasksRequestTasksTrigger) SetEndTime(v string) *BatchUpdateTasksRequestTasksTrigger {
	s.EndTime = &v
	return s
}

func (s *BatchUpdateTasksRequestTasksTrigger) SetRecurrence(v string) *BatchUpdateTasksRequestTasksTrigger {
	s.Recurrence = &v
	return s
}

func (s *BatchUpdateTasksRequestTasksTrigger) SetStartTime(v string) *BatchUpdateTasksRequestTasksTrigger {
	s.StartTime = &v
	return s
}

func (s *BatchUpdateTasksRequestTasksTrigger) SetType(v string) *BatchUpdateTasksRequestTasksTrigger {
	s.Type = &v
	return s
}

func (s *BatchUpdateTasksRequestTasksTrigger) Validate() error {
	return dara.Validate(s)
}
