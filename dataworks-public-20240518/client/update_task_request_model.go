// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientUniqueCode(v string) *UpdateTaskRequest
	GetClientUniqueCode() *string
	SetDataSource(v *UpdateTaskRequestDataSource) *UpdateTaskRequest
	GetDataSource() *UpdateTaskRequestDataSource
	SetDependencies(v []*UpdateTaskRequestDependencies) *UpdateTaskRequest
	GetDependencies() []*UpdateTaskRequestDependencies
	SetDescription(v string) *UpdateTaskRequest
	GetDescription() *string
	SetEnvType(v string) *UpdateTaskRequest
	GetEnvType() *string
	SetId(v int64) *UpdateTaskRequest
	GetId() *int64
	SetInputs(v *UpdateTaskRequestInputs) *UpdateTaskRequest
	GetInputs() *UpdateTaskRequestInputs
	SetInstanceMode(v string) *UpdateTaskRequest
	GetInstanceMode() *string
	SetName(v string) *UpdateTaskRequest
	GetName() *string
	SetOutputs(v *UpdateTaskRequestOutputs) *UpdateTaskRequest
	GetOutputs() *UpdateTaskRequestOutputs
	SetOwner(v string) *UpdateTaskRequest
	GetOwner() *string
	SetRerunInterval(v int32) *UpdateTaskRequest
	GetRerunInterval() *int32
	SetRerunMode(v string) *UpdateTaskRequest
	GetRerunMode() *string
	SetRerunTimes(v int32) *UpdateTaskRequest
	GetRerunTimes() *int32
	SetRuntimeResource(v *UpdateTaskRequestRuntimeResource) *UpdateTaskRequest
	GetRuntimeResource() *UpdateTaskRequestRuntimeResource
	SetScript(v *UpdateTaskRequestScript) *UpdateTaskRequest
	GetScript() *UpdateTaskRequestScript
	SetTags(v []*UpdateTaskRequestTags) *UpdateTaskRequest
	GetTags() []*UpdateTaskRequestTags
	SetTimeout(v int32) *UpdateTaskRequest
	GetTimeout() *int32
	SetTrigger(v *UpdateTaskRequestTrigger) *UpdateTaskRequest
	GetTrigger() *UpdateTaskRequestTrigger
}

type UpdateTaskRequest struct {
	// The unique code of the client. This code uniquely identifies a task. This parameter is used to create a task asynchronously and implement the idempotence of the task. If you do not specify this parameter when you create the task, the system automatically generates a unique code. The unique code is uniquely associated with the task ID. If you specify this parameter when you update or delete the task, the value of this parameter must be the unique code that is used to create the task.
	//
	// example:
	//
	// Task_0bc5213917368545132902xxxxxxxx
	ClientUniqueCode *string `json:"ClientUniqueCode,omitempty" xml:"ClientUniqueCode,omitempty"`
	// The information about the associated data source.
	DataSource *UpdateTaskRequestDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The dependency information.
	Dependencies []*UpdateTaskRequestDependencies `json:"Dependencies,omitempty" xml:"Dependencies,omitempty" type:"Repeated"`
	// The description of the task.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The project environment.
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
	// This parameter is required.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The input information.
	Inputs *UpdateTaskRequestInputs `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Struct"`
	// The instance generation mode. Valid values:
	//
	// 	- T+1: the next day
	//
	// 	- Immediately
	//
	// example:
	//
	// T+1
	InstanceMode *string `json:"InstanceMode,omitempty" xml:"InstanceMode,omitempty"`
	// Name.
	//
	// example:
	//
	// SQL node
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output information.
	Outputs *UpdateTaskRequestOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Struct"`
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
	// 	- AllDenied: The task cannot be rerun.
	//
	// 	- FailureAllowed: The task can be rerun only after it fails.
	//
	// 	- AllAllowed: The task can always be rerun.
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
	// Runtime environment configurations, such as resource group information.
	RuntimeResource *UpdateTaskRequestRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// The run script information.
	Script *UpdateTaskRequestScript `json:"Script,omitempty" xml:"Script,omitempty" type:"Struct"`
	// The tags.
	Tags []*UpdateTaskRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The timeout period of task running. Unit: seconds.
	//
	// example:
	//
	// 3600
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The triggering method.
	Trigger *UpdateTaskRequestTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
}

func (s UpdateTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskRequest) GetClientUniqueCode() *string {
	return s.ClientUniqueCode
}

func (s *UpdateTaskRequest) GetDataSource() *UpdateTaskRequestDataSource {
	return s.DataSource
}

func (s *UpdateTaskRequest) GetDependencies() []*UpdateTaskRequestDependencies {
	return s.Dependencies
}

func (s *UpdateTaskRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateTaskRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *UpdateTaskRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateTaskRequest) GetInputs() *UpdateTaskRequestInputs {
	return s.Inputs
}

func (s *UpdateTaskRequest) GetInstanceMode() *string {
	return s.InstanceMode
}

func (s *UpdateTaskRequest) GetName() *string {
	return s.Name
}

func (s *UpdateTaskRequest) GetOutputs() *UpdateTaskRequestOutputs {
	return s.Outputs
}

func (s *UpdateTaskRequest) GetOwner() *string {
	return s.Owner
}

func (s *UpdateTaskRequest) GetRerunInterval() *int32 {
	return s.RerunInterval
}

func (s *UpdateTaskRequest) GetRerunMode() *string {
	return s.RerunMode
}

func (s *UpdateTaskRequest) GetRerunTimes() *int32 {
	return s.RerunTimes
}

func (s *UpdateTaskRequest) GetRuntimeResource() *UpdateTaskRequestRuntimeResource {
	return s.RuntimeResource
}

func (s *UpdateTaskRequest) GetScript() *UpdateTaskRequestScript {
	return s.Script
}

func (s *UpdateTaskRequest) GetTags() []*UpdateTaskRequestTags {
	return s.Tags
}

func (s *UpdateTaskRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateTaskRequest) GetTrigger() *UpdateTaskRequestTrigger {
	return s.Trigger
}

func (s *UpdateTaskRequest) SetClientUniqueCode(v string) *UpdateTaskRequest {
	s.ClientUniqueCode = &v
	return s
}

func (s *UpdateTaskRequest) SetDataSource(v *UpdateTaskRequestDataSource) *UpdateTaskRequest {
	s.DataSource = v
	return s
}

func (s *UpdateTaskRequest) SetDependencies(v []*UpdateTaskRequestDependencies) *UpdateTaskRequest {
	s.Dependencies = v
	return s
}

func (s *UpdateTaskRequest) SetDescription(v string) *UpdateTaskRequest {
	s.Description = &v
	return s
}

func (s *UpdateTaskRequest) SetEnvType(v string) *UpdateTaskRequest {
	s.EnvType = &v
	return s
}

func (s *UpdateTaskRequest) SetId(v int64) *UpdateTaskRequest {
	s.Id = &v
	return s
}

func (s *UpdateTaskRequest) SetInputs(v *UpdateTaskRequestInputs) *UpdateTaskRequest {
	s.Inputs = v
	return s
}

func (s *UpdateTaskRequest) SetInstanceMode(v string) *UpdateTaskRequest {
	s.InstanceMode = &v
	return s
}

func (s *UpdateTaskRequest) SetName(v string) *UpdateTaskRequest {
	s.Name = &v
	return s
}

func (s *UpdateTaskRequest) SetOutputs(v *UpdateTaskRequestOutputs) *UpdateTaskRequest {
	s.Outputs = v
	return s
}

func (s *UpdateTaskRequest) SetOwner(v string) *UpdateTaskRequest {
	s.Owner = &v
	return s
}

func (s *UpdateTaskRequest) SetRerunInterval(v int32) *UpdateTaskRequest {
	s.RerunInterval = &v
	return s
}

func (s *UpdateTaskRequest) SetRerunMode(v string) *UpdateTaskRequest {
	s.RerunMode = &v
	return s
}

func (s *UpdateTaskRequest) SetRerunTimes(v int32) *UpdateTaskRequest {
	s.RerunTimes = &v
	return s
}

func (s *UpdateTaskRequest) SetRuntimeResource(v *UpdateTaskRequestRuntimeResource) *UpdateTaskRequest {
	s.RuntimeResource = v
	return s
}

func (s *UpdateTaskRequest) SetScript(v *UpdateTaskRequestScript) *UpdateTaskRequest {
	s.Script = v
	return s
}

func (s *UpdateTaskRequest) SetTags(v []*UpdateTaskRequestTags) *UpdateTaskRequest {
	s.Tags = v
	return s
}

func (s *UpdateTaskRequest) SetTimeout(v int32) *UpdateTaskRequest {
	s.Timeout = &v
	return s
}

func (s *UpdateTaskRequest) SetTrigger(v *UpdateTaskRequestTrigger) *UpdateTaskRequest {
	s.Trigger = v
	return s
}

func (s *UpdateTaskRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateTaskRequestDataSource struct {
	// The name of the data source.
	//
	// example:
	//
	// odps_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateTaskRequestDataSource) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskRequestDataSource) GoString() string {
	return s.String()
}

func (s *UpdateTaskRequestDataSource) GetName() *string {
	return s.Name
}

func (s *UpdateTaskRequestDataSource) SetName(v string) *UpdateTaskRequestDataSource {
	s.Name = &v
	return s
}

func (s *UpdateTaskRequestDataSource) Validate() error {
	return dara.Validate(s)
}

type UpdateTaskRequestDependencies struct {
	// The dependency type. Valid values:
	//
	// 	- CrossCycleDependsOnChildren: Depends on level-1 downstream nodes across cycles
	//
	// 	- CrossCycleDependsOnSelf: Depends on itself across cycles.
	//
	// 	- CrossCycleDependsOnOtherNode: Depends on other nodes across cycles.
	//
	// 	- Normal: Depends on nodes in the same cycle.
	//
	// This parameter is required.
	//
	// example:
	//
	// Normal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The output identifier of the upstream task. (This parameter is returned only if `Normal` is set and the node input is configured.)
	//
	// example:
	//
	// pre.odps_sql_demo_0
	UpstreamOutput *string `json:"UpstreamOutput,omitempty" xml:"UpstreamOutput,omitempty"`
	// The ID of the upstream task. (This parameter is returned only if `Normal` or `CrossCycleDependsOnOtherNode` is set and the node input is not configured.)
	//
	// example:
	//
	// 1234
	UpstreamTaskId *int64 `json:"UpstreamTaskId,omitempty" xml:"UpstreamTaskId,omitempty"`
}

func (s UpdateTaskRequestDependencies) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskRequestDependencies) GoString() string {
	return s.String()
}

func (s *UpdateTaskRequestDependencies) GetType() *string {
	return s.Type
}

func (s *UpdateTaskRequestDependencies) GetUpstreamOutput() *string {
	return s.UpstreamOutput
}

func (s *UpdateTaskRequestDependencies) GetUpstreamTaskId() *int64 {
	return s.UpstreamTaskId
}

func (s *UpdateTaskRequestDependencies) SetType(v string) *UpdateTaskRequestDependencies {
	s.Type = &v
	return s
}

func (s *UpdateTaskRequestDependencies) SetUpstreamOutput(v string) *UpdateTaskRequestDependencies {
	s.UpstreamOutput = &v
	return s
}

func (s *UpdateTaskRequestDependencies) SetUpstreamTaskId(v int64) *UpdateTaskRequestDependencies {
	s.UpstreamTaskId = &v
	return s
}

func (s *UpdateTaskRequestDependencies) Validate() error {
	return dara.Validate(s)
}

type UpdateTaskRequestInputs struct {
	// The variables.
	Variables []*UpdateTaskRequestInputsVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s UpdateTaskRequestInputs) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskRequestInputs) GoString() string {
	return s.String()
}

func (s *UpdateTaskRequestInputs) GetVariables() []*UpdateTaskRequestInputsVariables {
	return s.Variables
}

func (s *UpdateTaskRequestInputs) SetVariables(v []*UpdateTaskRequestInputsVariables) *UpdateTaskRequestInputs {
	s.Variables = v
	return s
}

func (s *UpdateTaskRequestInputs) Validate() error {
	return dara.Validate(s)
}

type UpdateTaskRequestInputsVariables struct {
	// The name of the variable.
	//
	// example:
	//
	// key1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type. Valid values:
	//
	// 	- Constant: constant.
	//
	// 	- PassThrough: node output.
	//
	// 	- System: variable.
	//
	// 	- NodeOutput: script output.
	//
	// This parameter is required.
	//
	// example:
	//
	// Constant
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of the variable.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateTaskRequestInputsVariables) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskRequestInputsVariables) GoString() string {
	return s.String()
}

func (s *UpdateTaskRequestInputsVariables) GetName() *string {
	return s.Name
}

func (s *UpdateTaskRequestInputsVariables) GetType() *string {
	return s.Type
}

func (s *UpdateTaskRequestInputsVariables) GetValue() *string {
	return s.Value
}

func (s *UpdateTaskRequestInputsVariables) SetName(v string) *UpdateTaskRequestInputsVariables {
	s.Name = &v
	return s
}

func (s *UpdateTaskRequestInputsVariables) SetType(v string) *UpdateTaskRequestInputsVariables {
	s.Type = &v
	return s
}

func (s *UpdateTaskRequestInputsVariables) SetValue(v string) *UpdateTaskRequestInputsVariables {
	s.Value = &v
	return s
}

func (s *UpdateTaskRequestInputsVariables) Validate() error {
	return dara.Validate(s)
}

type UpdateTaskRequestOutputs struct {
	// The task outputs.
	TaskOutputs []*UpdateTaskRequestOutputsTaskOutputs `json:"TaskOutputs,omitempty" xml:"TaskOutputs,omitempty" type:"Repeated"`
	// The variables.
	Variables []*UpdateTaskRequestOutputsVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s UpdateTaskRequestOutputs) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskRequestOutputs) GoString() string {
	return s.String()
}

func (s *UpdateTaskRequestOutputs) GetTaskOutputs() []*UpdateTaskRequestOutputsTaskOutputs {
	return s.TaskOutputs
}

func (s *UpdateTaskRequestOutputs) GetVariables() []*UpdateTaskRequestOutputsVariables {
	return s.Variables
}

func (s *UpdateTaskRequestOutputs) SetTaskOutputs(v []*UpdateTaskRequestOutputsTaskOutputs) *UpdateTaskRequestOutputs {
	s.TaskOutputs = v
	return s
}

func (s *UpdateTaskRequestOutputs) SetVariables(v []*UpdateTaskRequestOutputsVariables) *UpdateTaskRequestOutputs {
	s.Variables = v
	return s
}

func (s *UpdateTaskRequestOutputs) Validate() error {
	return dara.Validate(s)
}

type UpdateTaskRequestOutputsTaskOutputs struct {
	// The identifier of the output.
	//
	// example:
	//
	// pre.odps_sql_demo_0
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s UpdateTaskRequestOutputsTaskOutputs) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskRequestOutputsTaskOutputs) GoString() string {
	return s.String()
}

func (s *UpdateTaskRequestOutputsTaskOutputs) GetOutput() *string {
	return s.Output
}

func (s *UpdateTaskRequestOutputsTaskOutputs) SetOutput(v string) *UpdateTaskRequestOutputsTaskOutputs {
	s.Output = &v
	return s
}

func (s *UpdateTaskRequestOutputsTaskOutputs) Validate() error {
	return dara.Validate(s)
}

type UpdateTaskRequestOutputsVariables struct {
	// The name of the variable.
	//
	// example:
	//
	// key1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type. Valid values:
	//
	// 	- Constant: constant.
	//
	// 	- PassThrough: node output.
	//
	// 	- System: variable.
	//
	// 	- NodeOutput: script output.
	//
	// This parameter is required.
	//
	// example:
	//
	// Constant
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of the variable.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateTaskRequestOutputsVariables) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskRequestOutputsVariables) GoString() string {
	return s.String()
}

func (s *UpdateTaskRequestOutputsVariables) GetName() *string {
	return s.Name
}

func (s *UpdateTaskRequestOutputsVariables) GetType() *string {
	return s.Type
}

func (s *UpdateTaskRequestOutputsVariables) GetValue() *string {
	return s.Value
}

func (s *UpdateTaskRequestOutputsVariables) SetName(v string) *UpdateTaskRequestOutputsVariables {
	s.Name = &v
	return s
}

func (s *UpdateTaskRequestOutputsVariables) SetType(v string) *UpdateTaskRequestOutputsVariables {
	s.Type = &v
	return s
}

func (s *UpdateTaskRequestOutputsVariables) SetValue(v string) *UpdateTaskRequestOutputsVariables {
	s.Value = &v
	return s
}

func (s *UpdateTaskRequestOutputsVariables) Validate() error {
	return dara.Validate(s)
}

type UpdateTaskRequestRuntimeResource struct {
	// The default number of compute units (CUs) configured for task running.
	//
	// example:
	//
	// 0.25
	Cu *string `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The image ID used in the task runtime configuration.
	//
	// example:
	//
	// i-xxxxxx
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The identifier of the scheduling resource group used in the task runtime configuration.
	//
	// example:
	//
	// S_res_group_524258031846018_1684XXXXXXXXX
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s UpdateTaskRequestRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskRequestRuntimeResource) GoString() string {
	return s.String()
}

func (s *UpdateTaskRequestRuntimeResource) GetCu() *string {
	return s.Cu
}

func (s *UpdateTaskRequestRuntimeResource) GetImage() *string {
	return s.Image
}

func (s *UpdateTaskRequestRuntimeResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateTaskRequestRuntimeResource) SetCu(v string) *UpdateTaskRequestRuntimeResource {
	s.Cu = &v
	return s
}

func (s *UpdateTaskRequestRuntimeResource) SetImage(v string) *UpdateTaskRequestRuntimeResource {
	s.Image = &v
	return s
}

func (s *UpdateTaskRequestRuntimeResource) SetResourceGroupId(v string) *UpdateTaskRequestRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateTaskRequestRuntimeResource) Validate() error {
	return dara.Validate(s)
}

type UpdateTaskRequestScript struct {
	// Deprecated
	//
	// The script content.
	//
	// example:
	//
	// echo "helloWorld"
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The script parameter list.
	//
	// example:
	//
	// para1=$bizdate
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
}

func (s UpdateTaskRequestScript) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskRequestScript) GoString() string {
	return s.String()
}

func (s *UpdateTaskRequestScript) GetContent() *string {
	return s.Content
}

func (s *UpdateTaskRequestScript) GetParameters() *string {
	return s.Parameters
}

func (s *UpdateTaskRequestScript) SetContent(v string) *UpdateTaskRequestScript {
	s.Content = &v
	return s
}

func (s *UpdateTaskRequestScript) SetParameters(v string) *UpdateTaskRequestScript {
	s.Parameters = &v
	return s
}

func (s *UpdateTaskRequestScript) Validate() error {
	return dara.Validate(s)
}

type UpdateTaskRequestTags struct {
	// The key of a tag.
	//
	// This parameter is required.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of a tag.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateTaskRequestTags) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskRequestTags) GoString() string {
	return s.String()
}

func (s *UpdateTaskRequestTags) GetKey() *string {
	return s.Key
}

func (s *UpdateTaskRequestTags) GetValue() *string {
	return s.Value
}

func (s *UpdateTaskRequestTags) SetKey(v string) *UpdateTaskRequestTags {
	s.Key = &v
	return s
}

func (s *UpdateTaskRequestTags) SetValue(v string) *UpdateTaskRequestTags {
	s.Value = &v
	return s
}

func (s *UpdateTaskRequestTags) Validate() error {
	return dara.Validate(s)
}

type UpdateTaskRequestTrigger struct {
	// The Cron expression. This parameter takes effect only if the Type parameter is set to Scheduler.
	//
	// example:
	//
	// 00 00 00 	- 	- ?
	Cron      *string `json:"Cron,omitempty" xml:"Cron,omitempty"`
	CycleType *string `json:"CycleType,omitempty" xml:"CycleType,omitempty"`
	// The expiration time of periodic triggering. Takes effect only when type is set to Scheduler. The value of this parameter is in the`yyyy-mm-dd hh:mm:ss` format.
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
	// The time when periodic triggering takes effect. This parameter takes effect only if the Type parameter is set to Scheduler. The value of this parameter is in the`yyyy-mm-dd hh:mm:ss` format.
	//
	// example:
	//
	// 1970-01-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The triggering type. Valid values:
	//
	// 	- Scheduler: periodically triggered
	//
	// 	- Manual
	//
	// example:
	//
	// Scheduler
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateTaskRequestTrigger) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskRequestTrigger) GoString() string {
	return s.String()
}

func (s *UpdateTaskRequestTrigger) GetCron() *string {
	return s.Cron
}

func (s *UpdateTaskRequestTrigger) GetCycleType() *string {
	return s.CycleType
}

func (s *UpdateTaskRequestTrigger) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateTaskRequestTrigger) GetRecurrence() *string {
	return s.Recurrence
}

func (s *UpdateTaskRequestTrigger) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateTaskRequestTrigger) GetType() *string {
	return s.Type
}

func (s *UpdateTaskRequestTrigger) SetCron(v string) *UpdateTaskRequestTrigger {
	s.Cron = &v
	return s
}

func (s *UpdateTaskRequestTrigger) SetCycleType(v string) *UpdateTaskRequestTrigger {
	s.CycleType = &v
	return s
}

func (s *UpdateTaskRequestTrigger) SetEndTime(v string) *UpdateTaskRequestTrigger {
	s.EndTime = &v
	return s
}

func (s *UpdateTaskRequestTrigger) SetRecurrence(v string) *UpdateTaskRequestTrigger {
	s.Recurrence = &v
	return s
}

func (s *UpdateTaskRequestTrigger) SetStartTime(v string) *UpdateTaskRequestTrigger {
	s.StartTime = &v
	return s
}

func (s *UpdateTaskRequestTrigger) SetType(v string) *UpdateTaskRequestTrigger {
	s.Type = &v
	return s
}

func (s *UpdateTaskRequestTrigger) Validate() error {
	return dara.Validate(s)
}
