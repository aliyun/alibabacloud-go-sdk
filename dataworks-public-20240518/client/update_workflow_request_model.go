// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientUniqueCode(v string) *UpdateWorkflowRequest
	GetClientUniqueCode() *string
	SetDependencies(v []*UpdateWorkflowRequestDependencies) *UpdateWorkflowRequest
	GetDependencies() []*UpdateWorkflowRequestDependencies
	SetDescription(v string) *UpdateWorkflowRequest
	GetDescription() *string
	SetEnvType(v string) *UpdateWorkflowRequest
	GetEnvType() *string
	SetId(v int64) *UpdateWorkflowRequest
	GetId() *int64
	SetInstanceMode(v string) *UpdateWorkflowRequest
	GetInstanceMode() *string
	SetName(v string) *UpdateWorkflowRequest
	GetName() *string
	SetOutputs(v *UpdateWorkflowRequestOutputs) *UpdateWorkflowRequest
	GetOutputs() *UpdateWorkflowRequestOutputs
	SetOwner(v string) *UpdateWorkflowRequest
	GetOwner() *string
	SetParameters(v string) *UpdateWorkflowRequest
	GetParameters() *string
	SetTags(v []*UpdateWorkflowRequestTags) *UpdateWorkflowRequest
	GetTags() []*UpdateWorkflowRequestTags
	SetTasks(v []*UpdateWorkflowRequestTasks) *UpdateWorkflowRequest
	GetTasks() []*UpdateWorkflowRequestTasks
	SetTrigger(v *UpdateWorkflowRequestTrigger) *UpdateWorkflowRequest
	GetTrigger() *UpdateWorkflowRequestTrigger
}

type UpdateWorkflowRequest struct {
	// The unique code of the client. This parameter is used to create a workflow asynchronously and implement the idempotence of the workflow. If you do not specify this parameter when you create the workflow, the system automatically generates a unique code. The unique code is uniquely associated with the workflow ID. If you specify this parameter when you update or delete the workflow, the value of this parameter must be the unique code that is used to create the workflow.
	//
	// example:
	//
	// Workflow_0bc5213917368545132902xxxxxxxx
	ClientUniqueCode *string `json:"ClientUniqueCode,omitempty" xml:"ClientUniqueCode,omitempty"`
	// The dependency information.
	Dependencies []*UpdateWorkflowRequestDependencies `json:"Dependencies,omitempty" xml:"Dependencies,omitempty" type:"Repeated"`
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
	// The workflow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceMode *string `json:"InstanceMode,omitempty" xml:"InstanceMode,omitempty"`
	// The name of the workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// My Workflow
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output information.
	Outputs *UpdateWorkflowRequestOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Struct"`
	// The account ID of the owner.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The parameters.
	//
	// example:
	//
	// para1=$bizdate para2=$[yyyymmdd]
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The tags.
	Tags []*UpdateWorkflowRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The tasks.
	Tasks []*UpdateWorkflowRequestTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// The trigger method.
	//
	// This parameter is required.
	Trigger *UpdateWorkflowRequestTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
}

func (s UpdateWorkflowRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowRequest) GetClientUniqueCode() *string {
	return s.ClientUniqueCode
}

func (s *UpdateWorkflowRequest) GetDependencies() []*UpdateWorkflowRequestDependencies {
	return s.Dependencies
}

func (s *UpdateWorkflowRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateWorkflowRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *UpdateWorkflowRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateWorkflowRequest) GetInstanceMode() *string {
	return s.InstanceMode
}

func (s *UpdateWorkflowRequest) GetName() *string {
	return s.Name
}

func (s *UpdateWorkflowRequest) GetOutputs() *UpdateWorkflowRequestOutputs {
	return s.Outputs
}

func (s *UpdateWorkflowRequest) GetOwner() *string {
	return s.Owner
}

func (s *UpdateWorkflowRequest) GetParameters() *string {
	return s.Parameters
}

func (s *UpdateWorkflowRequest) GetTags() []*UpdateWorkflowRequestTags {
	return s.Tags
}

func (s *UpdateWorkflowRequest) GetTasks() []*UpdateWorkflowRequestTasks {
	return s.Tasks
}

func (s *UpdateWorkflowRequest) GetTrigger() *UpdateWorkflowRequestTrigger {
	return s.Trigger
}

func (s *UpdateWorkflowRequest) SetClientUniqueCode(v string) *UpdateWorkflowRequest {
	s.ClientUniqueCode = &v
	return s
}

func (s *UpdateWorkflowRequest) SetDependencies(v []*UpdateWorkflowRequestDependencies) *UpdateWorkflowRequest {
	s.Dependencies = v
	return s
}

func (s *UpdateWorkflowRequest) SetDescription(v string) *UpdateWorkflowRequest {
	s.Description = &v
	return s
}

func (s *UpdateWorkflowRequest) SetEnvType(v string) *UpdateWorkflowRequest {
	s.EnvType = &v
	return s
}

func (s *UpdateWorkflowRequest) SetId(v int64) *UpdateWorkflowRequest {
	s.Id = &v
	return s
}

func (s *UpdateWorkflowRequest) SetInstanceMode(v string) *UpdateWorkflowRequest {
	s.InstanceMode = &v
	return s
}

func (s *UpdateWorkflowRequest) SetName(v string) *UpdateWorkflowRequest {
	s.Name = &v
	return s
}

func (s *UpdateWorkflowRequest) SetOutputs(v *UpdateWorkflowRequestOutputs) *UpdateWorkflowRequest {
	s.Outputs = v
	return s
}

func (s *UpdateWorkflowRequest) SetOwner(v string) *UpdateWorkflowRequest {
	s.Owner = &v
	return s
}

func (s *UpdateWorkflowRequest) SetParameters(v string) *UpdateWorkflowRequest {
	s.Parameters = &v
	return s
}

func (s *UpdateWorkflowRequest) SetTags(v []*UpdateWorkflowRequestTags) *UpdateWorkflowRequest {
	s.Tags = v
	return s
}

func (s *UpdateWorkflowRequest) SetTasks(v []*UpdateWorkflowRequestTasks) *UpdateWorkflowRequest {
	s.Tasks = v
	return s
}

func (s *UpdateWorkflowRequest) SetTrigger(v *UpdateWorkflowRequestTrigger) *UpdateWorkflowRequest {
	s.Trigger = v
	return s
}

func (s *UpdateWorkflowRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkflowRequestDependencies struct {
	// The dependency type. Valid values:
	//
	// 	- CrossCycleDependsOnChildren: cross-cycle dependency on level-1 descendant nodes
	//
	// 	- CrossCycleDependsOnSelf: cross-cycle dependency on the current node
	//
	// 	- CrossCycleDependsOnOtherNode: cross-cycle dependency on other nodes
	//
	// 	- Normal: same-cycle scheduling dependency
	//
	// This parameter is required.
	//
	// example:
	//
	// Normal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The identifier of the output of the ancestor task. This parameter is returned only if `same-cycle scheduling dependencies` and the node input are configured.
	//
	// example:
	//
	// pre.odps_sql_demo_0
	UpstreamOutput *string `json:"UpstreamOutput,omitempty" xml:"UpstreamOutput,omitempty"`
	// The ancestor task ID. This parameter is returned only if `cross-cycle scheduling dependencies` or `same-cycle scheduling dependencies` and the node input are not configured.
	//
	// example:
	//
	// 1234
	UpstreamTaskId *int64 `json:"UpstreamTaskId,omitempty" xml:"UpstreamTaskId,omitempty"`
}

func (s UpdateWorkflowRequestDependencies) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowRequestDependencies) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowRequestDependencies) GetType() *string {
	return s.Type
}

func (s *UpdateWorkflowRequestDependencies) GetUpstreamOutput() *string {
	return s.UpstreamOutput
}

func (s *UpdateWorkflowRequestDependencies) GetUpstreamTaskId() *int64 {
	return s.UpstreamTaskId
}

func (s *UpdateWorkflowRequestDependencies) SetType(v string) *UpdateWorkflowRequestDependencies {
	s.Type = &v
	return s
}

func (s *UpdateWorkflowRequestDependencies) SetUpstreamOutput(v string) *UpdateWorkflowRequestDependencies {
	s.UpstreamOutput = &v
	return s
}

func (s *UpdateWorkflowRequestDependencies) SetUpstreamTaskId(v int64) *UpdateWorkflowRequestDependencies {
	s.UpstreamTaskId = &v
	return s
}

func (s *UpdateWorkflowRequestDependencies) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkflowRequestOutputs struct {
	// The task outputs.
	TaskOutputs []*UpdateWorkflowRequestOutputsTaskOutputs `json:"TaskOutputs,omitempty" xml:"TaskOutputs,omitempty" type:"Repeated"`
}

func (s UpdateWorkflowRequestOutputs) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowRequestOutputs) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowRequestOutputs) GetTaskOutputs() []*UpdateWorkflowRequestOutputsTaskOutputs {
	return s.TaskOutputs
}

func (s *UpdateWorkflowRequestOutputs) SetTaskOutputs(v []*UpdateWorkflowRequestOutputsTaskOutputs) *UpdateWorkflowRequestOutputs {
	s.TaskOutputs = v
	return s
}

func (s *UpdateWorkflowRequestOutputs) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkflowRequestOutputsTaskOutputs struct {
	// The identifier of the output.
	//
	// example:
	//
	// pre.odps_sql_demo_0
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s UpdateWorkflowRequestOutputsTaskOutputs) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowRequestOutputsTaskOutputs) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowRequestOutputsTaskOutputs) GetOutput() *string {
	return s.Output
}

func (s *UpdateWorkflowRequestOutputsTaskOutputs) SetOutput(v string) *UpdateWorkflowRequestOutputsTaskOutputs {
	s.Output = &v
	return s
}

func (s *UpdateWorkflowRequestOutputsTaskOutputs) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkflowRequestTags struct {
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

func (s UpdateWorkflowRequestTags) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowRequestTags) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowRequestTags) GetKey() *string {
	return s.Key
}

func (s *UpdateWorkflowRequestTags) GetValue() *string {
	return s.Value
}

func (s *UpdateWorkflowRequestTags) SetKey(v string) *UpdateWorkflowRequestTags {
	s.Key = &v
	return s
}

func (s *UpdateWorkflowRequestTags) SetValue(v string) *UpdateWorkflowRequestTags {
	s.Value = &v
	return s
}

func (s *UpdateWorkflowRequestTags) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkflowRequestTasks struct {
	// The baseline ID.
	//
	// example:
	//
	// 1234
	BaseLineId *int64 `json:"BaseLineId,omitempty" xml:"BaseLineId,omitempty"`
	// The unique code of the client. This parameter is used to create a task asynchronously and implement the idempotence of the task. If you do not specify this parameter when you create the workflow, the system automatically generates a unique code. The unique code is uniquely associated with the workflow ID. If you specify this parameter when you update or delete the workflow, the value of this parameter must be the unique code that is used to create the workflow.
	//
	// example:
	//
	// Task_0bc5213917368545132902xxxxxxxx
	ClientUniqueCode *string `json:"ClientUniqueCode,omitempty" xml:"ClientUniqueCode,omitempty"`
	// The information about the associated data source.
	DataSource *UpdateWorkflowRequestTasksDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The dependency information.
	Dependencies []*UpdateWorkflowRequestTasksDependencies `json:"Dependencies,omitempty" xml:"Dependencies,omitempty" type:"Repeated"`
	// The description.
	//
	// example:
	//
	// Test
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
	// The task ID. If you configure this parameter, full update is performed on the task. If you do not configure this parameter, another task is created.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The input information.
	Inputs *UpdateWorkflowRequestTasksInputs `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Struct"`
	// The name of the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// SQL node
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output information.
	Outputs *UpdateWorkflowRequestTasksOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Struct"`
	// The account ID of the owner.
	//
	// This parameter is required.
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
	// This parameter is required.
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
	//
	// This parameter is required.
	RuntimeResource *UpdateWorkflowRequestTasksRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// The script information.
	Script *UpdateWorkflowRequestTasksScript `json:"Script,omitempty" xml:"Script,omitempty" type:"Struct"`
	// The tags.
	Tags []*UpdateWorkflowRequestTasksTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The timeout period of task running. Unit: seconds.
	//
	// example:
	//
	// 3600
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The trigger method.
	//
	// This parameter is required.
	Trigger *UpdateWorkflowRequestTasksTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
	// The type of the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// ODPS_SQL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateWorkflowRequestTasks) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowRequestTasks) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowRequestTasks) GetBaseLineId() *int64 {
	return s.BaseLineId
}

func (s *UpdateWorkflowRequestTasks) GetClientUniqueCode() *string {
	return s.ClientUniqueCode
}

func (s *UpdateWorkflowRequestTasks) GetDataSource() *UpdateWorkflowRequestTasksDataSource {
	return s.DataSource
}

func (s *UpdateWorkflowRequestTasks) GetDependencies() []*UpdateWorkflowRequestTasksDependencies {
	return s.Dependencies
}

func (s *UpdateWorkflowRequestTasks) GetDescription() *string {
	return s.Description
}

func (s *UpdateWorkflowRequestTasks) GetEnvType() *string {
	return s.EnvType
}

func (s *UpdateWorkflowRequestTasks) GetId() *int64 {
	return s.Id
}

func (s *UpdateWorkflowRequestTasks) GetInputs() *UpdateWorkflowRequestTasksInputs {
	return s.Inputs
}

func (s *UpdateWorkflowRequestTasks) GetName() *string {
	return s.Name
}

func (s *UpdateWorkflowRequestTasks) GetOutputs() *UpdateWorkflowRequestTasksOutputs {
	return s.Outputs
}

func (s *UpdateWorkflowRequestTasks) GetOwner() *string {
	return s.Owner
}

func (s *UpdateWorkflowRequestTasks) GetRerunInterval() *int32 {
	return s.RerunInterval
}

func (s *UpdateWorkflowRequestTasks) GetRerunMode() *string {
	return s.RerunMode
}

func (s *UpdateWorkflowRequestTasks) GetRerunTimes() *int32 {
	return s.RerunTimes
}

func (s *UpdateWorkflowRequestTasks) GetRuntimeResource() *UpdateWorkflowRequestTasksRuntimeResource {
	return s.RuntimeResource
}

func (s *UpdateWorkflowRequestTasks) GetScript() *UpdateWorkflowRequestTasksScript {
	return s.Script
}

func (s *UpdateWorkflowRequestTasks) GetTags() []*UpdateWorkflowRequestTasksTags {
	return s.Tags
}

func (s *UpdateWorkflowRequestTasks) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateWorkflowRequestTasks) GetTrigger() *UpdateWorkflowRequestTasksTrigger {
	return s.Trigger
}

func (s *UpdateWorkflowRequestTasks) GetType() *string {
	return s.Type
}

func (s *UpdateWorkflowRequestTasks) SetBaseLineId(v int64) *UpdateWorkflowRequestTasks {
	s.BaseLineId = &v
	return s
}

func (s *UpdateWorkflowRequestTasks) SetClientUniqueCode(v string) *UpdateWorkflowRequestTasks {
	s.ClientUniqueCode = &v
	return s
}

func (s *UpdateWorkflowRequestTasks) SetDataSource(v *UpdateWorkflowRequestTasksDataSource) *UpdateWorkflowRequestTasks {
	s.DataSource = v
	return s
}

func (s *UpdateWorkflowRequestTasks) SetDependencies(v []*UpdateWorkflowRequestTasksDependencies) *UpdateWorkflowRequestTasks {
	s.Dependencies = v
	return s
}

func (s *UpdateWorkflowRequestTasks) SetDescription(v string) *UpdateWorkflowRequestTasks {
	s.Description = &v
	return s
}

func (s *UpdateWorkflowRequestTasks) SetEnvType(v string) *UpdateWorkflowRequestTasks {
	s.EnvType = &v
	return s
}

func (s *UpdateWorkflowRequestTasks) SetId(v int64) *UpdateWorkflowRequestTasks {
	s.Id = &v
	return s
}

func (s *UpdateWorkflowRequestTasks) SetInputs(v *UpdateWorkflowRequestTasksInputs) *UpdateWorkflowRequestTasks {
	s.Inputs = v
	return s
}

func (s *UpdateWorkflowRequestTasks) SetName(v string) *UpdateWorkflowRequestTasks {
	s.Name = &v
	return s
}

func (s *UpdateWorkflowRequestTasks) SetOutputs(v *UpdateWorkflowRequestTasksOutputs) *UpdateWorkflowRequestTasks {
	s.Outputs = v
	return s
}

func (s *UpdateWorkflowRequestTasks) SetOwner(v string) *UpdateWorkflowRequestTasks {
	s.Owner = &v
	return s
}

func (s *UpdateWorkflowRequestTasks) SetRerunInterval(v int32) *UpdateWorkflowRequestTasks {
	s.RerunInterval = &v
	return s
}

func (s *UpdateWorkflowRequestTasks) SetRerunMode(v string) *UpdateWorkflowRequestTasks {
	s.RerunMode = &v
	return s
}

func (s *UpdateWorkflowRequestTasks) SetRerunTimes(v int32) *UpdateWorkflowRequestTasks {
	s.RerunTimes = &v
	return s
}

func (s *UpdateWorkflowRequestTasks) SetRuntimeResource(v *UpdateWorkflowRequestTasksRuntimeResource) *UpdateWorkflowRequestTasks {
	s.RuntimeResource = v
	return s
}

func (s *UpdateWorkflowRequestTasks) SetScript(v *UpdateWorkflowRequestTasksScript) *UpdateWorkflowRequestTasks {
	s.Script = v
	return s
}

func (s *UpdateWorkflowRequestTasks) SetTags(v []*UpdateWorkflowRequestTasksTags) *UpdateWorkflowRequestTasks {
	s.Tags = v
	return s
}

func (s *UpdateWorkflowRequestTasks) SetTimeout(v int32) *UpdateWorkflowRequestTasks {
	s.Timeout = &v
	return s
}

func (s *UpdateWorkflowRequestTasks) SetTrigger(v *UpdateWorkflowRequestTasksTrigger) *UpdateWorkflowRequestTasks {
	s.Trigger = v
	return s
}

func (s *UpdateWorkflowRequestTasks) SetType(v string) *UpdateWorkflowRequestTasks {
	s.Type = &v
	return s
}

func (s *UpdateWorkflowRequestTasks) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkflowRequestTasksDataSource struct {
	// The name of the data source.
	//
	// example:
	//
	// odps_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateWorkflowRequestTasksDataSource) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowRequestTasksDataSource) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowRequestTasksDataSource) GetName() *string {
	return s.Name
}

func (s *UpdateWorkflowRequestTasksDataSource) SetName(v string) *UpdateWorkflowRequestTasksDataSource {
	s.Name = &v
	return s
}

func (s *UpdateWorkflowRequestTasksDataSource) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkflowRequestTasksDependencies struct {
	// The dependency type. Valid values:
	//
	// 	- CrossCycleDependsOnChildren: cross-cycle dependency on level-1 descendant nodes
	//
	// 	- CrossCycleDependsOnSelf: cross-cycle dependency on the current node
	//
	// 	- CrossCycleDependsOnOtherNode: cross-cycle dependency on other nodes
	//
	// 	- Normal: same-cycle scheduling dependency
	//
	// This parameter is required.
	//
	// example:
	//
	// Normal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The identifier of the output of the ancestor task. This parameter is returned only if `same-cycle scheduling dependencies` and the node input are configured.
	//
	// example:
	//
	// pre.odps_sql_demo_0
	UpstreamOutput *string `json:"UpstreamOutput,omitempty" xml:"UpstreamOutput,omitempty"`
	// The ancestor task ID. This parameter is returned only if `cross-cycle scheduling dependencies` or `same-cycle scheduling dependencies` and the node input are not configured.
	//
	// example:
	//
	// 1234
	UpstreamTaskId *int64 `json:"UpstreamTaskId,omitempty" xml:"UpstreamTaskId,omitempty"`
}

func (s UpdateWorkflowRequestTasksDependencies) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowRequestTasksDependencies) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowRequestTasksDependencies) GetType() *string {
	return s.Type
}

func (s *UpdateWorkflowRequestTasksDependencies) GetUpstreamOutput() *string {
	return s.UpstreamOutput
}

func (s *UpdateWorkflowRequestTasksDependencies) GetUpstreamTaskId() *int64 {
	return s.UpstreamTaskId
}

func (s *UpdateWorkflowRequestTasksDependencies) SetType(v string) *UpdateWorkflowRequestTasksDependencies {
	s.Type = &v
	return s
}

func (s *UpdateWorkflowRequestTasksDependencies) SetUpstreamOutput(v string) *UpdateWorkflowRequestTasksDependencies {
	s.UpstreamOutput = &v
	return s
}

func (s *UpdateWorkflowRequestTasksDependencies) SetUpstreamTaskId(v int64) *UpdateWorkflowRequestTasksDependencies {
	s.UpstreamTaskId = &v
	return s
}

func (s *UpdateWorkflowRequestTasksDependencies) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkflowRequestTasksInputs struct {
	// The variables.
	Variables []*UpdateWorkflowRequestTasksInputsVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s UpdateWorkflowRequestTasksInputs) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowRequestTasksInputs) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowRequestTasksInputs) GetVariables() []*UpdateWorkflowRequestTasksInputsVariables {
	return s.Variables
}

func (s *UpdateWorkflowRequestTasksInputs) SetVariables(v []*UpdateWorkflowRequestTasksInputsVariables) *UpdateWorkflowRequestTasksInputs {
	s.Variables = v
	return s
}

func (s *UpdateWorkflowRequestTasksInputs) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkflowRequestTasksInputsVariables struct {
	// The name of the variable.
	//
	// example:
	//
	// key1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type. Valid values:
	//
	// 	- Constant: constant
	//
	// 	- PassThrough: node output
	//
	// 	- System: variable
	//
	// 	- NodeOutput: script output
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

func (s UpdateWorkflowRequestTasksInputsVariables) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowRequestTasksInputsVariables) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowRequestTasksInputsVariables) GetName() *string {
	return s.Name
}

func (s *UpdateWorkflowRequestTasksInputsVariables) GetType() *string {
	return s.Type
}

func (s *UpdateWorkflowRequestTasksInputsVariables) GetValue() *string {
	return s.Value
}

func (s *UpdateWorkflowRequestTasksInputsVariables) SetName(v string) *UpdateWorkflowRequestTasksInputsVariables {
	s.Name = &v
	return s
}

func (s *UpdateWorkflowRequestTasksInputsVariables) SetType(v string) *UpdateWorkflowRequestTasksInputsVariables {
	s.Type = &v
	return s
}

func (s *UpdateWorkflowRequestTasksInputsVariables) SetValue(v string) *UpdateWorkflowRequestTasksInputsVariables {
	s.Value = &v
	return s
}

func (s *UpdateWorkflowRequestTasksInputsVariables) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkflowRequestTasksOutputs struct {
	// The task outputs.
	TaskOutputs []*UpdateWorkflowRequestTasksOutputsTaskOutputs `json:"TaskOutputs,omitempty" xml:"TaskOutputs,omitempty" type:"Repeated"`
	// The variables.
	Variables []*UpdateWorkflowRequestTasksOutputsVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s UpdateWorkflowRequestTasksOutputs) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowRequestTasksOutputs) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowRequestTasksOutputs) GetTaskOutputs() []*UpdateWorkflowRequestTasksOutputsTaskOutputs {
	return s.TaskOutputs
}

func (s *UpdateWorkflowRequestTasksOutputs) GetVariables() []*UpdateWorkflowRequestTasksOutputsVariables {
	return s.Variables
}

func (s *UpdateWorkflowRequestTasksOutputs) SetTaskOutputs(v []*UpdateWorkflowRequestTasksOutputsTaskOutputs) *UpdateWorkflowRequestTasksOutputs {
	s.TaskOutputs = v
	return s
}

func (s *UpdateWorkflowRequestTasksOutputs) SetVariables(v []*UpdateWorkflowRequestTasksOutputsVariables) *UpdateWorkflowRequestTasksOutputs {
	s.Variables = v
	return s
}

func (s *UpdateWorkflowRequestTasksOutputs) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkflowRequestTasksOutputsTaskOutputs struct {
	// The identifier of the output.
	//
	// example:
	//
	// pre.odps_sql_demo_0
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s UpdateWorkflowRequestTasksOutputsTaskOutputs) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowRequestTasksOutputsTaskOutputs) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowRequestTasksOutputsTaskOutputs) GetOutput() *string {
	return s.Output
}

func (s *UpdateWorkflowRequestTasksOutputsTaskOutputs) SetOutput(v string) *UpdateWorkflowRequestTasksOutputsTaskOutputs {
	s.Output = &v
	return s
}

func (s *UpdateWorkflowRequestTasksOutputsTaskOutputs) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkflowRequestTasksOutputsVariables struct {
	// The name of the variable.
	//
	// example:
	//
	// key1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type. Valid values:
	//
	// 	- Constant: constant
	//
	// 	- PassThrough: node output
	//
	// 	- System: variable
	//
	// 	- NodeOutput: script output
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

func (s UpdateWorkflowRequestTasksOutputsVariables) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowRequestTasksOutputsVariables) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowRequestTasksOutputsVariables) GetName() *string {
	return s.Name
}

func (s *UpdateWorkflowRequestTasksOutputsVariables) GetType() *string {
	return s.Type
}

func (s *UpdateWorkflowRequestTasksOutputsVariables) GetValue() *string {
	return s.Value
}

func (s *UpdateWorkflowRequestTasksOutputsVariables) SetName(v string) *UpdateWorkflowRequestTasksOutputsVariables {
	s.Name = &v
	return s
}

func (s *UpdateWorkflowRequestTasksOutputsVariables) SetType(v string) *UpdateWorkflowRequestTasksOutputsVariables {
	s.Type = &v
	return s
}

func (s *UpdateWorkflowRequestTasksOutputsVariables) SetValue(v string) *UpdateWorkflowRequestTasksOutputsVariables {
	s.Value = &v
	return s
}

func (s *UpdateWorkflowRequestTasksOutputsVariables) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkflowRequestTasksRuntimeResource struct {
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
	// This parameter is required.
	//
	// example:
	//
	// S_res_group_524258031846018_1684XXXXXXXXX
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s UpdateWorkflowRequestTasksRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowRequestTasksRuntimeResource) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowRequestTasksRuntimeResource) GetCu() *string {
	return s.Cu
}

func (s *UpdateWorkflowRequestTasksRuntimeResource) GetImage() *string {
	return s.Image
}

func (s *UpdateWorkflowRequestTasksRuntimeResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateWorkflowRequestTasksRuntimeResource) SetCu(v string) *UpdateWorkflowRequestTasksRuntimeResource {
	s.Cu = &v
	return s
}

func (s *UpdateWorkflowRequestTasksRuntimeResource) SetImage(v string) *UpdateWorkflowRequestTasksRuntimeResource {
	s.Image = &v
	return s
}

func (s *UpdateWorkflowRequestTasksRuntimeResource) SetResourceGroupId(v string) *UpdateWorkflowRequestTasksRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateWorkflowRequestTasksRuntimeResource) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkflowRequestTasksScript struct {
	// The script content.
	//
	// example:
	//
	// echo "helloWorld"
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The script parameters.
	//
	// example:
	//
	// para1=$bizdate
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
}

func (s UpdateWorkflowRequestTasksScript) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowRequestTasksScript) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowRequestTasksScript) GetContent() *string {
	return s.Content
}

func (s *UpdateWorkflowRequestTasksScript) GetParameters() *string {
	return s.Parameters
}

func (s *UpdateWorkflowRequestTasksScript) SetContent(v string) *UpdateWorkflowRequestTasksScript {
	s.Content = &v
	return s
}

func (s *UpdateWorkflowRequestTasksScript) SetParameters(v string) *UpdateWorkflowRequestTasksScript {
	s.Parameters = &v
	return s
}

func (s *UpdateWorkflowRequestTasksScript) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkflowRequestTasksTags struct {
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

func (s UpdateWorkflowRequestTasksTags) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowRequestTasksTags) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowRequestTasksTags) GetKey() *string {
	return s.Key
}

func (s *UpdateWorkflowRequestTasksTags) GetValue() *string {
	return s.Value
}

func (s *UpdateWorkflowRequestTasksTags) SetKey(v string) *UpdateWorkflowRequestTasksTags {
	s.Key = &v
	return s
}

func (s *UpdateWorkflowRequestTasksTags) SetValue(v string) *UpdateWorkflowRequestTasksTags {
	s.Value = &v
	return s
}

func (s *UpdateWorkflowRequestTasksTags) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkflowRequestTasksTrigger struct {
	// The running mode of the task after it is triggered. This parameter takes effect only if the Type parameter is set to Scheduler. Valid values:
	//
	// 	- Pause
	//
	// 	- Skip
	//
	// 	- Normal
	//
	// This parameter is required.
	//
	// example:
	//
	// Normal
	Recurrence *string `json:"Recurrence,omitempty" xml:"Recurrence,omitempty"`
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

func (s UpdateWorkflowRequestTasksTrigger) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowRequestTasksTrigger) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowRequestTasksTrigger) GetRecurrence() *string {
	return s.Recurrence
}

func (s *UpdateWorkflowRequestTasksTrigger) GetType() *string {
	return s.Type
}

func (s *UpdateWorkflowRequestTasksTrigger) SetRecurrence(v string) *UpdateWorkflowRequestTasksTrigger {
	s.Recurrence = &v
	return s
}

func (s *UpdateWorkflowRequestTasksTrigger) SetType(v string) *UpdateWorkflowRequestTasksTrigger {
	s.Type = &v
	return s
}

func (s *UpdateWorkflowRequestTasksTrigger) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkflowRequestTrigger struct {
	// The CRON expression. This parameter takes effect only if the Type parameter is set to Scheduler.
	//
	// example:
	//
	// 00 00 00 	- 	- ?
	Cron *string `json:"Cron,omitempty" xml:"Cron,omitempty"`
	// The end time of the time range during which the workflow is periodically scheduled. This parameter takes effect only if the Type parameter is set to Scheduler. The value of this parameter is in the `yyyy-mm-dd hh:mm:ss` format.
	//
	// example:
	//
	// 9999-01-01 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of the time range during which the workflow is periodically scheduled. This parameter takes effect only if the Type parameter is set to Scheduler. The value of this parameter is in the `yyyy-mm-dd hh:mm:ss` format.
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
	// This parameter is required.
	//
	// example:
	//
	// Scheduler
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateWorkflowRequestTrigger) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowRequestTrigger) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowRequestTrigger) GetCron() *string {
	return s.Cron
}

func (s *UpdateWorkflowRequestTrigger) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateWorkflowRequestTrigger) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateWorkflowRequestTrigger) GetType() *string {
	return s.Type
}

func (s *UpdateWorkflowRequestTrigger) SetCron(v string) *UpdateWorkflowRequestTrigger {
	s.Cron = &v
	return s
}

func (s *UpdateWorkflowRequestTrigger) SetEndTime(v string) *UpdateWorkflowRequestTrigger {
	s.EndTime = &v
	return s
}

func (s *UpdateWorkflowRequestTrigger) SetStartTime(v string) *UpdateWorkflowRequestTrigger {
	s.StartTime = &v
	return s
}

func (s *UpdateWorkflowRequestTrigger) SetType(v string) *UpdateWorkflowRequestTrigger {
	s.Type = &v
	return s
}

func (s *UpdateWorkflowRequestTrigger) Validate() error {
	return dara.Validate(s)
}
