// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskAsyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientUniqueCode(v string) *UpdateTaskAsyncRequest
	GetClientUniqueCode() *string
	SetDataSource(v *UpdateTaskAsyncRequestDataSource) *UpdateTaskAsyncRequest
	GetDataSource() *UpdateTaskAsyncRequestDataSource
	SetDependencies(v []*UpdateTaskAsyncRequestDependencies) *UpdateTaskAsyncRequest
	GetDependencies() []*UpdateTaskAsyncRequestDependencies
	SetDescription(v string) *UpdateTaskAsyncRequest
	GetDescription() *string
	SetEnvType(v string) *UpdateTaskAsyncRequest
	GetEnvType() *string
	SetId(v int64) *UpdateTaskAsyncRequest
	GetId() *int64
	SetInputs(v *UpdateTaskAsyncRequestInputs) *UpdateTaskAsyncRequest
	GetInputs() *UpdateTaskAsyncRequestInputs
	SetInstanceMode(v string) *UpdateTaskAsyncRequest
	GetInstanceMode() *string
	SetName(v string) *UpdateTaskAsyncRequest
	GetName() *string
	SetOutputs(v *UpdateTaskAsyncRequestOutputs) *UpdateTaskAsyncRequest
	GetOutputs() *UpdateTaskAsyncRequestOutputs
	SetOwner(v string) *UpdateTaskAsyncRequest
	GetOwner() *string
	SetRerunInterval(v int32) *UpdateTaskAsyncRequest
	GetRerunInterval() *int32
	SetRerunMode(v string) *UpdateTaskAsyncRequest
	GetRerunMode() *string
	SetRerunTimes(v int32) *UpdateTaskAsyncRequest
	GetRerunTimes() *int32
	SetRuntimeResource(v *UpdateTaskAsyncRequestRuntimeResource) *UpdateTaskAsyncRequest
	GetRuntimeResource() *UpdateTaskAsyncRequestRuntimeResource
	SetScript(v *UpdateTaskAsyncRequestScript) *UpdateTaskAsyncRequest
	GetScript() *UpdateTaskAsyncRequestScript
	SetTags(v []*UpdateTaskAsyncRequestTags) *UpdateTaskAsyncRequest
	GetTags() []*UpdateTaskAsyncRequestTags
	SetTimeout(v int32) *UpdateTaskAsyncRequest
	GetTimeout() *int32
	SetTrigger(v *UpdateTaskAsyncRequestTrigger) *UpdateTaskAsyncRequest
	GetTrigger() *UpdateTaskAsyncRequestTrigger
}

type UpdateTaskAsyncRequest struct {
	// The client unique code of the node, used to uniquely identify a node. This code is used for asynchronous processing and idempotence. If you do not specify this parameter when creating a node, the system automatically generates a value and binds it to the resource ID. If you specify this parameter when updating or deleting a resource, the value must match the client unique code used when the resource was created.
	//
	// example:
	//
	// Workflow_0bc5213917368545132902xxxxxxxx
	ClientUniqueCode *string `json:"ClientUniqueCode,omitempty" xml:"ClientUniqueCode,omitempty"`
	// The associated data source information.
	DataSource *UpdateTaskAsyncRequestDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The dependency information.
	Dependencies []*UpdateTaskAsyncRequestDependencies `json:"Dependencies,omitempty" xml:"Dependencies,omitempty" type:"Repeated"`
	// The description.
	//
	// example:
	//
	// This is a description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The project environment. Valid values:
	//
	// - Prod: production
	//
	// - Dev: development
	//
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The node ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The input information.
	Inputs *UpdateTaskAsyncRequestInputs `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Struct"`
	// The instance generation mode. Valid values:
	//
	// - T+1: generates instances the next day.
	//
	// - Immediately: generates instances immediately. Note: only periodic instances whose scheduled time is more than 10 minutes after the node publish time are generated normally. During the full-to-instance conversion period (22:00–24:00), real-time instance conversion is not supported. You can submit and publish nodes, but new nodes are not automatically converted to instances.
	//
	// example:
	//
	// T+1
	InstanceMode *string `json:"InstanceMode,omitempty" xml:"InstanceMode,omitempty"`
	// The name.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output information.
	Outputs *UpdateTaskAsyncRequestOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Struct"`
	// The account ID of the node owner. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and hover over the profile picture in the upper-right corner of the top navigation bar to view the account ID. If this parameter is left empty, the Alibaba Cloud account ID of the caller is used by default.
	//
	// example:
	//
	// 1000000000001
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The retry time interval, in milliseconds. The value cannot exceed 1800000.
	//
	// example:
	//
	// 60000
	RerunInterval *int32 `json:"RerunInterval,omitempty" xml:"RerunInterval,omitempty"`
	// The configuration that specifies whether the node can be rerun. Valid values:
	//
	//
	//
	//
	// - AllDenied: the node cannot be rerun regardless of whether it succeeds or fails.
	//
	// - FailureAllowed: the node can be rerun only if it fails.
	//
	// - AllAllowed: the node can be rerun regardless of whether it succeeds or fails.
	//
	// example:
	//
	// AllAllowed
	RerunMode *string `json:"RerunMode,omitempty" xml:"RerunMode,omitempty"`
	// The number of retries. This parameter takes effect only when the node is configured to allow reruns.
	//
	// example:
	//
	// 3
	RerunTimes *int32 `json:"RerunTimes,omitempty" xml:"RerunTimes,omitempty"`
	// The runtime environment configuration, such as the resource group information.
	RuntimeResource *UpdateTaskAsyncRequestRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// The runtime script information.
	Script *UpdateTaskAsyncRequestScript `json:"Script,omitempty" xml:"Script,omitempty" type:"Struct"`
	// The list of data asset tags to attach.
	Tags []*UpdateTaskAsyncRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The timeout period defined in the scheduling configuration.
	//
	// example:
	//
	// 1
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The trigger method of the node.
	Trigger *UpdateTaskAsyncRequestTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
}

func (s UpdateTaskAsyncRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskAsyncRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskAsyncRequest) GetClientUniqueCode() *string {
	return s.ClientUniqueCode
}

func (s *UpdateTaskAsyncRequest) GetDataSource() *UpdateTaskAsyncRequestDataSource {
	return s.DataSource
}

func (s *UpdateTaskAsyncRequest) GetDependencies() []*UpdateTaskAsyncRequestDependencies {
	return s.Dependencies
}

func (s *UpdateTaskAsyncRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateTaskAsyncRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *UpdateTaskAsyncRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateTaskAsyncRequest) GetInputs() *UpdateTaskAsyncRequestInputs {
	return s.Inputs
}

func (s *UpdateTaskAsyncRequest) GetInstanceMode() *string {
	return s.InstanceMode
}

func (s *UpdateTaskAsyncRequest) GetName() *string {
	return s.Name
}

func (s *UpdateTaskAsyncRequest) GetOutputs() *UpdateTaskAsyncRequestOutputs {
	return s.Outputs
}

func (s *UpdateTaskAsyncRequest) GetOwner() *string {
	return s.Owner
}

func (s *UpdateTaskAsyncRequest) GetRerunInterval() *int32 {
	return s.RerunInterval
}

func (s *UpdateTaskAsyncRequest) GetRerunMode() *string {
	return s.RerunMode
}

func (s *UpdateTaskAsyncRequest) GetRerunTimes() *int32 {
	return s.RerunTimes
}

func (s *UpdateTaskAsyncRequest) GetRuntimeResource() *UpdateTaskAsyncRequestRuntimeResource {
	return s.RuntimeResource
}

func (s *UpdateTaskAsyncRequest) GetScript() *UpdateTaskAsyncRequestScript {
	return s.Script
}

func (s *UpdateTaskAsyncRequest) GetTags() []*UpdateTaskAsyncRequestTags {
	return s.Tags
}

func (s *UpdateTaskAsyncRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateTaskAsyncRequest) GetTrigger() *UpdateTaskAsyncRequestTrigger {
	return s.Trigger
}

func (s *UpdateTaskAsyncRequest) SetClientUniqueCode(v string) *UpdateTaskAsyncRequest {
	s.ClientUniqueCode = &v
	return s
}

func (s *UpdateTaskAsyncRequest) SetDataSource(v *UpdateTaskAsyncRequestDataSource) *UpdateTaskAsyncRequest {
	s.DataSource = v
	return s
}

func (s *UpdateTaskAsyncRequest) SetDependencies(v []*UpdateTaskAsyncRequestDependencies) *UpdateTaskAsyncRequest {
	s.Dependencies = v
	return s
}

func (s *UpdateTaskAsyncRequest) SetDescription(v string) *UpdateTaskAsyncRequest {
	s.Description = &v
	return s
}

func (s *UpdateTaskAsyncRequest) SetEnvType(v string) *UpdateTaskAsyncRequest {
	s.EnvType = &v
	return s
}

func (s *UpdateTaskAsyncRequest) SetId(v int64) *UpdateTaskAsyncRequest {
	s.Id = &v
	return s
}

func (s *UpdateTaskAsyncRequest) SetInputs(v *UpdateTaskAsyncRequestInputs) *UpdateTaskAsyncRequest {
	s.Inputs = v
	return s
}

func (s *UpdateTaskAsyncRequest) SetInstanceMode(v string) *UpdateTaskAsyncRequest {
	s.InstanceMode = &v
	return s
}

func (s *UpdateTaskAsyncRequest) SetName(v string) *UpdateTaskAsyncRequest {
	s.Name = &v
	return s
}

func (s *UpdateTaskAsyncRequest) SetOutputs(v *UpdateTaskAsyncRequestOutputs) *UpdateTaskAsyncRequest {
	s.Outputs = v
	return s
}

func (s *UpdateTaskAsyncRequest) SetOwner(v string) *UpdateTaskAsyncRequest {
	s.Owner = &v
	return s
}

func (s *UpdateTaskAsyncRequest) SetRerunInterval(v int32) *UpdateTaskAsyncRequest {
	s.RerunInterval = &v
	return s
}

func (s *UpdateTaskAsyncRequest) SetRerunMode(v string) *UpdateTaskAsyncRequest {
	s.RerunMode = &v
	return s
}

func (s *UpdateTaskAsyncRequest) SetRerunTimes(v int32) *UpdateTaskAsyncRequest {
	s.RerunTimes = &v
	return s
}

func (s *UpdateTaskAsyncRequest) SetRuntimeResource(v *UpdateTaskAsyncRequestRuntimeResource) *UpdateTaskAsyncRequest {
	s.RuntimeResource = v
	return s
}

func (s *UpdateTaskAsyncRequest) SetScript(v *UpdateTaskAsyncRequestScript) *UpdateTaskAsyncRequest {
	s.Script = v
	return s
}

func (s *UpdateTaskAsyncRequest) SetTags(v []*UpdateTaskAsyncRequestTags) *UpdateTaskAsyncRequest {
	s.Tags = v
	return s
}

func (s *UpdateTaskAsyncRequest) SetTimeout(v int32) *UpdateTaskAsyncRequest {
	s.Timeout = &v
	return s
}

func (s *UpdateTaskAsyncRequest) SetTrigger(v *UpdateTaskAsyncRequestTrigger) *UpdateTaskAsyncRequest {
	s.Trigger = v
	return s
}

func (s *UpdateTaskAsyncRequest) Validate() error {
	if s.DataSource != nil {
		if err := s.DataSource.Validate(); err != nil {
			return err
		}
	}
	if s.Dependencies != nil {
		for _, item := range s.Dependencies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Inputs != nil {
		if err := s.Inputs.Validate(); err != nil {
			return err
		}
	}
	if s.Outputs != nil {
		if err := s.Outputs.Validate(); err != nil {
			return err
		}
	}
	if s.RuntimeResource != nil {
		if err := s.RuntimeResource.Validate(); err != nil {
			return err
		}
	}
	if s.Script != nil {
		if err := s.Script.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Trigger != nil {
		if err := s.Trigger.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateTaskAsyncRequestDataSource struct {
	// The data source name.
	//
	// example:
	//
	// odps_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateTaskAsyncRequestDataSource) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskAsyncRequestDataSource) GoString() string {
	return s.String()
}

func (s *UpdateTaskAsyncRequestDataSource) GetName() *string {
	return s.Name
}

func (s *UpdateTaskAsyncRequestDataSource) SetName(v string) *UpdateTaskAsyncRequestDataSource {
	s.Name = &v
	return s
}

func (s *UpdateTaskAsyncRequestDataSource) Validate() error {
	return dara.Validate(s)
}

type UpdateTaskAsyncRequestDependencies struct {
	// The dependency type. Valid values:
	//
	// - CrossCycleDependsOnChildren: cross-cycle dependency on first-level child nodes.
	//
	// - CrossCycleDependsOnSelf: cross-cycle dependency on the node itself.
	//
	// - CrossCycleDependsOnOtherNode: cross-cycle dependency on other nodes.
	//
	// - Normal: same-cycle dependency.
	//
	// This parameter is required.
	//
	// example:
	//
	// Normal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The output identifier of the upstream node. This field is returned for same-cycle dependencies when the input content is configured.
	//
	// example:
	//
	// pre.odps_sql_demo_0
	UpstreamOutput *string `json:"UpstreamOutput,omitempty" xml:"UpstreamOutput,omitempty"`
	// The ID of the upstream node. This field is returned for cross-cycle dependencies on other nodes and for same-cycle dependencies when no input content is configured. It is not returned in other cases.
	//
	// example:
	//
	// 1234
	UpstreamTaskId *int64 `json:"UpstreamTaskId,omitempty" xml:"UpstreamTaskId,omitempty"`
}

func (s UpdateTaskAsyncRequestDependencies) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskAsyncRequestDependencies) GoString() string {
	return s.String()
}

func (s *UpdateTaskAsyncRequestDependencies) GetType() *string {
	return s.Type
}

func (s *UpdateTaskAsyncRequestDependencies) GetUpstreamOutput() *string {
	return s.UpstreamOutput
}

func (s *UpdateTaskAsyncRequestDependencies) GetUpstreamTaskId() *int64 {
	return s.UpstreamTaskId
}

func (s *UpdateTaskAsyncRequestDependencies) SetType(v string) *UpdateTaskAsyncRequestDependencies {
	s.Type = &v
	return s
}

func (s *UpdateTaskAsyncRequestDependencies) SetUpstreamOutput(v string) *UpdateTaskAsyncRequestDependencies {
	s.UpstreamOutput = &v
	return s
}

func (s *UpdateTaskAsyncRequestDependencies) SetUpstreamTaskId(v int64) *UpdateTaskAsyncRequestDependencies {
	s.UpstreamTaskId = &v
	return s
}

func (s *UpdateTaskAsyncRequestDependencies) Validate() error {
	return dara.Validate(s)
}

type UpdateTaskAsyncRequestInputs struct {
	// The list of variable definitions.
	Variables []*UpdateTaskAsyncRequestInputsVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s UpdateTaskAsyncRequestInputs) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskAsyncRequestInputs) GoString() string {
	return s.String()
}

func (s *UpdateTaskAsyncRequestInputs) GetVariables() []*UpdateTaskAsyncRequestInputsVariables {
	return s.Variables
}

func (s *UpdateTaskAsyncRequestInputs) SetVariables(v []*UpdateTaskAsyncRequestInputsVariables) *UpdateTaskAsyncRequestInputs {
	s.Variables = v
	return s
}

func (s *UpdateTaskAsyncRequestInputs) Validate() error {
	if s.Variables != nil {
		for _, item := range s.Variables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateTaskAsyncRequestInputsVariables struct {
	// The variable name.
	//
	// example:
	//
	// key1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The variable type. Valid values:
	//
	// - Constant: constant.
	//
	// - PassThrough: output of a pass-through parameter node.
	//
	// - System: system variable.
	//
	// - NodeOutput: script output.
	//
	// This parameter is required.
	//
	// example:
	//
	// Constant
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The variable value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateTaskAsyncRequestInputsVariables) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskAsyncRequestInputsVariables) GoString() string {
	return s.String()
}

func (s *UpdateTaskAsyncRequestInputsVariables) GetName() *string {
	return s.Name
}

func (s *UpdateTaskAsyncRequestInputsVariables) GetType() *string {
	return s.Type
}

func (s *UpdateTaskAsyncRequestInputsVariables) GetValue() *string {
	return s.Value
}

func (s *UpdateTaskAsyncRequestInputsVariables) SetName(v string) *UpdateTaskAsyncRequestInputsVariables {
	s.Name = &v
	return s
}

func (s *UpdateTaskAsyncRequestInputsVariables) SetType(v string) *UpdateTaskAsyncRequestInputsVariables {
	s.Type = &v
	return s
}

func (s *UpdateTaskAsyncRequestInputsVariables) SetValue(v string) *UpdateTaskAsyncRequestInputsVariables {
	s.Value = &v
	return s
}

func (s *UpdateTaskAsyncRequestInputsVariables) Validate() error {
	return dara.Validate(s)
}

type UpdateTaskAsyncRequestOutputs struct {
	// The list of node output definitions.
	TaskOutputs []*UpdateTaskAsyncRequestOutputsTaskOutputs `json:"TaskOutputs,omitempty" xml:"TaskOutputs,omitempty" type:"Repeated"`
	// The list of variable definitions.
	Variables []*UpdateTaskAsyncRequestOutputsVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s UpdateTaskAsyncRequestOutputs) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskAsyncRequestOutputs) GoString() string {
	return s.String()
}

func (s *UpdateTaskAsyncRequestOutputs) GetTaskOutputs() []*UpdateTaskAsyncRequestOutputsTaskOutputs {
	return s.TaskOutputs
}

func (s *UpdateTaskAsyncRequestOutputs) GetVariables() []*UpdateTaskAsyncRequestOutputsVariables {
	return s.Variables
}

func (s *UpdateTaskAsyncRequestOutputs) SetTaskOutputs(v []*UpdateTaskAsyncRequestOutputsTaskOutputs) *UpdateTaskAsyncRequestOutputs {
	s.TaskOutputs = v
	return s
}

func (s *UpdateTaskAsyncRequestOutputs) SetVariables(v []*UpdateTaskAsyncRequestOutputsVariables) *UpdateTaskAsyncRequestOutputs {
	s.Variables = v
	return s
}

func (s *UpdateTaskAsyncRequestOutputs) Validate() error {
	if s.TaskOutputs != nil {
		for _, item := range s.TaskOutputs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Variables != nil {
		for _, item := range s.Variables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateTaskAsyncRequestOutputsTaskOutputs struct {
	// The output identifier.
	//
	// example:
	//
	// pre.odps_sql_demo_0
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s UpdateTaskAsyncRequestOutputsTaskOutputs) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskAsyncRequestOutputsTaskOutputs) GoString() string {
	return s.String()
}

func (s *UpdateTaskAsyncRequestOutputsTaskOutputs) GetOutput() *string {
	return s.Output
}

func (s *UpdateTaskAsyncRequestOutputsTaskOutputs) SetOutput(v string) *UpdateTaskAsyncRequestOutputsTaskOutputs {
	s.Output = &v
	return s
}

func (s *UpdateTaskAsyncRequestOutputsTaskOutputs) Validate() error {
	return dara.Validate(s)
}

type UpdateTaskAsyncRequestOutputsVariables struct {
	// The variable name.
	//
	// example:
	//
	// key1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The variable type. Valid values:
	//
	// - Constant: constant.
	//
	// - PassThrough: output of a pass-through parameter node.
	//
	// - System: system variable.
	//
	// - NodeOutput: script output.
	//
	// This parameter is required.
	//
	// example:
	//
	// Constant
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The variable value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateTaskAsyncRequestOutputsVariables) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskAsyncRequestOutputsVariables) GoString() string {
	return s.String()
}

func (s *UpdateTaskAsyncRequestOutputsVariables) GetName() *string {
	return s.Name
}

func (s *UpdateTaskAsyncRequestOutputsVariables) GetType() *string {
	return s.Type
}

func (s *UpdateTaskAsyncRequestOutputsVariables) GetValue() *string {
	return s.Value
}

func (s *UpdateTaskAsyncRequestOutputsVariables) SetName(v string) *UpdateTaskAsyncRequestOutputsVariables {
	s.Name = &v
	return s
}

func (s *UpdateTaskAsyncRequestOutputsVariables) SetType(v string) *UpdateTaskAsyncRequestOutputsVariables {
	s.Type = &v
	return s
}

func (s *UpdateTaskAsyncRequestOutputsVariables) SetValue(v string) *UpdateTaskAsyncRequestOutputsVariables {
	s.Value = &v
	return s
}

func (s *UpdateTaskAsyncRequestOutputsVariables) Validate() error {
	return dara.Validate(s)
}

type UpdateTaskAsyncRequestRuntimeResource struct {
	// The CU consumption for the node runtime configuration.
	//
	// example:
	//
	// 0.25
	Cu *string `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The image ID for the node runtime configuration.
	//
	// example:
	//
	// i-xxxxxx
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The identifier of the schedule resource group for the node runtime configuration.
	//
	// example:
	//
	// 63900680
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s UpdateTaskAsyncRequestRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskAsyncRequestRuntimeResource) GoString() string {
	return s.String()
}

func (s *UpdateTaskAsyncRequestRuntimeResource) GetCu() *string {
	return s.Cu
}

func (s *UpdateTaskAsyncRequestRuntimeResource) GetImage() *string {
	return s.Image
}

func (s *UpdateTaskAsyncRequestRuntimeResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateTaskAsyncRequestRuntimeResource) SetCu(v string) *UpdateTaskAsyncRequestRuntimeResource {
	s.Cu = &v
	return s
}

func (s *UpdateTaskAsyncRequestRuntimeResource) SetImage(v string) *UpdateTaskAsyncRequestRuntimeResource {
	s.Image = &v
	return s
}

func (s *UpdateTaskAsyncRequestRuntimeResource) SetResourceGroupId(v string) *UpdateTaskAsyncRequestRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateTaskAsyncRequestRuntimeResource) Validate() error {
	return dara.Validate(s)
}

type UpdateTaskAsyncRequestScript struct {
	// Deprecated
	//
	// The script content.
	//
	// example:
	//
	// echo "helloWorld"
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The list of script parameters.
	//
	// example:
	//
	// para1=$bizdate
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
}

func (s UpdateTaskAsyncRequestScript) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskAsyncRequestScript) GoString() string {
	return s.String()
}

func (s *UpdateTaskAsyncRequestScript) GetContent() *string {
	return s.Content
}

func (s *UpdateTaskAsyncRequestScript) GetParameters() *string {
	return s.Parameters
}

func (s *UpdateTaskAsyncRequestScript) SetContent(v string) *UpdateTaskAsyncRequestScript {
	s.Content = &v
	return s
}

func (s *UpdateTaskAsyncRequestScript) SetParameters(v string) *UpdateTaskAsyncRequestScript {
	s.Parameters = &v
	return s
}

func (s *UpdateTaskAsyncRequestScript) Validate() error {
	return dara.Validate(s)
}

type UpdateTaskAsyncRequestTags struct {
	// The tag key.
	//
	// This parameter is required.
	//
	// example:
	//
	// tagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateTaskAsyncRequestTags) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskAsyncRequestTags) GoString() string {
	return s.String()
}

func (s *UpdateTaskAsyncRequestTags) GetKey() *string {
	return s.Key
}

func (s *UpdateTaskAsyncRequestTags) GetValue() *string {
	return s.Value
}

func (s *UpdateTaskAsyncRequestTags) SetKey(v string) *UpdateTaskAsyncRequestTags {
	s.Key = &v
	return s
}

func (s *UpdateTaskAsyncRequestTags) SetValue(v string) *UpdateTaskAsyncRequestTags {
	s.Value = &v
	return s
}

func (s *UpdateTaskAsyncRequestTags) Validate() error {
	return dara.Validate(s)
}

type UpdateTaskAsyncRequestTrigger struct {
	// The cron expression. This parameter takes effect when Type is set to Scheduler.
	//
	// example:
	//
	// 00 00 00 	- 	- ?
	Cron *string `json:"Cron,omitempty" xml:"Cron,omitempty"`
	// The scheduling cycle type. This parameter takes effect when Type is set to Scheduler and the cron expression specifies a timed scheduling at a specific hour. Default value: Daily. Valid values:
	//
	//
	//
	//
	// - Daily: daily scheduling.
	//
	// - NotDaily: hourly scheduling.
	//
	// example:
	//
	// Daily
	CycleType *string `json:"CycleType,omitempty" xml:"CycleType,omitempty"`
	// The time when the periodic trigger expires. This parameter takes effect when Type is set to Scheduler. The format is `yyyy-mm-dd hh:mm:ss`.
	//
	// example:
	//
	// 9999-01-01 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The run mode when the trigger fires. This parameter takes effect when Type is set to Scheduler. Valid values:
	//
	// - Pause: paused.
	//
	// - Skip: dry run.
	//
	// - Normal: normal run.
	//
	// example:
	//
	// Normal
	Recurrence *string `json:"Recurrence,omitempty" xml:"Recurrence,omitempty"`
	// The time at which the scheduled trigger takes effect. This parameter takes effect when Type is set to Scheduler. The format is `yyyy-mm-dd hh:mm:ss`.
	//
	// example:
	//
	// 1970-01-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The trigger type. Valid values:
	//
	// - Scheduler: scheduled periodic trigger.
	//
	// - Manual: manual trigger.
	//
	// example:
	//
	// BySchedule
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateTaskAsyncRequestTrigger) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskAsyncRequestTrigger) GoString() string {
	return s.String()
}

func (s *UpdateTaskAsyncRequestTrigger) GetCron() *string {
	return s.Cron
}

func (s *UpdateTaskAsyncRequestTrigger) GetCycleType() *string {
	return s.CycleType
}

func (s *UpdateTaskAsyncRequestTrigger) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateTaskAsyncRequestTrigger) GetRecurrence() *string {
	return s.Recurrence
}

func (s *UpdateTaskAsyncRequestTrigger) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateTaskAsyncRequestTrigger) GetType() *string {
	return s.Type
}

func (s *UpdateTaskAsyncRequestTrigger) SetCron(v string) *UpdateTaskAsyncRequestTrigger {
	s.Cron = &v
	return s
}

func (s *UpdateTaskAsyncRequestTrigger) SetCycleType(v string) *UpdateTaskAsyncRequestTrigger {
	s.CycleType = &v
	return s
}

func (s *UpdateTaskAsyncRequestTrigger) SetEndTime(v string) *UpdateTaskAsyncRequestTrigger {
	s.EndTime = &v
	return s
}

func (s *UpdateTaskAsyncRequestTrigger) SetRecurrence(v string) *UpdateTaskAsyncRequestTrigger {
	s.Recurrence = &v
	return s
}

func (s *UpdateTaskAsyncRequestTrigger) SetStartTime(v string) *UpdateTaskAsyncRequestTrigger {
	s.StartTime = &v
	return s
}

func (s *UpdateTaskAsyncRequestTrigger) SetType(v string) *UpdateTaskAsyncRequestTrigger {
	s.Type = &v
	return s
}

func (s *UpdateTaskAsyncRequestTrigger) Validate() error {
	return dara.Validate(s)
}
