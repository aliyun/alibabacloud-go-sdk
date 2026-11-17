// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskAsyncShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientUniqueCode(v string) *UpdateTaskAsyncShrinkRequest
	GetClientUniqueCode() *string
	SetDataSourceShrink(v string) *UpdateTaskAsyncShrinkRequest
	GetDataSourceShrink() *string
	SetDependenciesShrink(v string) *UpdateTaskAsyncShrinkRequest
	GetDependenciesShrink() *string
	SetDescription(v string) *UpdateTaskAsyncShrinkRequest
	GetDescription() *string
	SetEnvType(v string) *UpdateTaskAsyncShrinkRequest
	GetEnvType() *string
	SetId(v int64) *UpdateTaskAsyncShrinkRequest
	GetId() *int64
	SetInputsShrink(v string) *UpdateTaskAsyncShrinkRequest
	GetInputsShrink() *string
	SetInstanceMode(v string) *UpdateTaskAsyncShrinkRequest
	GetInstanceMode() *string
	SetName(v string) *UpdateTaskAsyncShrinkRequest
	GetName() *string
	SetOutputsShrink(v string) *UpdateTaskAsyncShrinkRequest
	GetOutputsShrink() *string
	SetOwner(v string) *UpdateTaskAsyncShrinkRequest
	GetOwner() *string
	SetRerunInterval(v int32) *UpdateTaskAsyncShrinkRequest
	GetRerunInterval() *int32
	SetRerunMode(v string) *UpdateTaskAsyncShrinkRequest
	GetRerunMode() *string
	SetRerunTimes(v int32) *UpdateTaskAsyncShrinkRequest
	GetRerunTimes() *int32
	SetRuntimeResourceShrink(v string) *UpdateTaskAsyncShrinkRequest
	GetRuntimeResourceShrink() *string
	SetScriptShrink(v string) *UpdateTaskAsyncShrinkRequest
	GetScriptShrink() *string
	SetTagsShrink(v string) *UpdateTaskAsyncShrinkRequest
	GetTagsShrink() *string
	SetTimeout(v int32) *UpdateTaskAsyncShrinkRequest
	GetTimeout() *int32
	SetTriggerShrink(v string) *UpdateTaskAsyncShrinkRequest
	GetTriggerShrink() *string
}

type UpdateTaskAsyncShrinkRequest struct {
	// The client unique code of the node, which is used to uniquely identify a node. This parameter is used to implement asynchronous processing and idempotence. If you do not specify this parameter when you create a node, the system automatically generates one. This code is uniquely bound to the resource ID. If you specify this parameter when you update or delete a resource, it must be the same as the client unique code used when the resource was created.
	//
	// example:
	//
	// Workflow_0bc5213917368545132902xxxxxxxx
	ClientUniqueCode *string `json:"ClientUniqueCode,omitempty" xml:"ClientUniqueCode,omitempty"`
	// The associated data source information.
	DataSourceShrink *string `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	// The dependency information.
	DependenciesShrink *string `json:"Dependencies,omitempty" xml:"Dependencies,omitempty"`
	// The description.
	//
	// example:
	//
	// This is a description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The project environment. Valid values:
	//
	// - Prod: production.
	//
	// - Dev: development.
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
	InputsShrink *string `json:"Inputs,omitempty" xml:"Inputs,omitempty"`
	// The instance generation mode. Valid values:
	//
	// - T+1: generates instances on the next day.
	//
	// - Immediately: generates instances immediately. Note: Periodic instances are generated normally only if the scheduled time of the node is at least 10 minutes after the node is published. Real-time instance generation is not supported during the full instance generation period (22:00 to 24:00). You can submit and publish nodes, but new nodes are not automatically converted into instances.
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
	OutputsShrink *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// The Alibaba Cloud account ID of the node owner. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and move the mouse over the profile picture in the upper-right corner of the top menu bar to view the account ID. If this parameter is left empty, the Alibaba Cloud account ID of the invoker is used by default.
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
	// Specifies whether the node can be rerun. Valid values:
	//
	// - AllDenied: The node cannot be rerun regardless of whether it succeeds or fails.
	//
	// - FailureAllowed: The node can be rerun only if it fails.
	//
	// - AllAllowed: The node can be rerun regardless of whether it succeeds or fails.
	//
	// example:
	//
	// AllAllowed
	RerunMode *string `json:"RerunMode,omitempty" xml:"RerunMode,omitempty"`
	// The number of retries. This parameter takes effect when the node is configured to be rerunnable.
	//
	// example:
	//
	// 3
	RerunTimes *int32 `json:"RerunTimes,omitempty" xml:"RerunTimes,omitempty"`
	// The runtime environment configuration, such as the resource group information.
	RuntimeResourceShrink *string `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty"`
	// The information about the script to be run.
	ScriptShrink *string `json:"Script,omitempty" xml:"Script,omitempty"`
	// The list of data asset tags to be bound.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The timeout period defined in the scheduling configuration.
	//
	// example:
	//
	// 1
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The trigger mode of the node.
	TriggerShrink *string `json:"Trigger,omitempty" xml:"Trigger,omitempty"`
}

func (s UpdateTaskAsyncShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskAsyncShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskAsyncShrinkRequest) GetClientUniqueCode() *string {
	return s.ClientUniqueCode
}

func (s *UpdateTaskAsyncShrinkRequest) GetDataSourceShrink() *string {
	return s.DataSourceShrink
}

func (s *UpdateTaskAsyncShrinkRequest) GetDependenciesShrink() *string {
	return s.DependenciesShrink
}

func (s *UpdateTaskAsyncShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateTaskAsyncShrinkRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *UpdateTaskAsyncShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateTaskAsyncShrinkRequest) GetInputsShrink() *string {
	return s.InputsShrink
}

func (s *UpdateTaskAsyncShrinkRequest) GetInstanceMode() *string {
	return s.InstanceMode
}

func (s *UpdateTaskAsyncShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateTaskAsyncShrinkRequest) GetOutputsShrink() *string {
	return s.OutputsShrink
}

func (s *UpdateTaskAsyncShrinkRequest) GetOwner() *string {
	return s.Owner
}

func (s *UpdateTaskAsyncShrinkRequest) GetRerunInterval() *int32 {
	return s.RerunInterval
}

func (s *UpdateTaskAsyncShrinkRequest) GetRerunMode() *string {
	return s.RerunMode
}

func (s *UpdateTaskAsyncShrinkRequest) GetRerunTimes() *int32 {
	return s.RerunTimes
}

func (s *UpdateTaskAsyncShrinkRequest) GetRuntimeResourceShrink() *string {
	return s.RuntimeResourceShrink
}

func (s *UpdateTaskAsyncShrinkRequest) GetScriptShrink() *string {
	return s.ScriptShrink
}

func (s *UpdateTaskAsyncShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *UpdateTaskAsyncShrinkRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateTaskAsyncShrinkRequest) GetTriggerShrink() *string {
	return s.TriggerShrink
}

func (s *UpdateTaskAsyncShrinkRequest) SetClientUniqueCode(v string) *UpdateTaskAsyncShrinkRequest {
	s.ClientUniqueCode = &v
	return s
}

func (s *UpdateTaskAsyncShrinkRequest) SetDataSourceShrink(v string) *UpdateTaskAsyncShrinkRequest {
	s.DataSourceShrink = &v
	return s
}

func (s *UpdateTaskAsyncShrinkRequest) SetDependenciesShrink(v string) *UpdateTaskAsyncShrinkRequest {
	s.DependenciesShrink = &v
	return s
}

func (s *UpdateTaskAsyncShrinkRequest) SetDescription(v string) *UpdateTaskAsyncShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateTaskAsyncShrinkRequest) SetEnvType(v string) *UpdateTaskAsyncShrinkRequest {
	s.EnvType = &v
	return s
}

func (s *UpdateTaskAsyncShrinkRequest) SetId(v int64) *UpdateTaskAsyncShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateTaskAsyncShrinkRequest) SetInputsShrink(v string) *UpdateTaskAsyncShrinkRequest {
	s.InputsShrink = &v
	return s
}

func (s *UpdateTaskAsyncShrinkRequest) SetInstanceMode(v string) *UpdateTaskAsyncShrinkRequest {
	s.InstanceMode = &v
	return s
}

func (s *UpdateTaskAsyncShrinkRequest) SetName(v string) *UpdateTaskAsyncShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateTaskAsyncShrinkRequest) SetOutputsShrink(v string) *UpdateTaskAsyncShrinkRequest {
	s.OutputsShrink = &v
	return s
}

func (s *UpdateTaskAsyncShrinkRequest) SetOwner(v string) *UpdateTaskAsyncShrinkRequest {
	s.Owner = &v
	return s
}

func (s *UpdateTaskAsyncShrinkRequest) SetRerunInterval(v int32) *UpdateTaskAsyncShrinkRequest {
	s.RerunInterval = &v
	return s
}

func (s *UpdateTaskAsyncShrinkRequest) SetRerunMode(v string) *UpdateTaskAsyncShrinkRequest {
	s.RerunMode = &v
	return s
}

func (s *UpdateTaskAsyncShrinkRequest) SetRerunTimes(v int32) *UpdateTaskAsyncShrinkRequest {
	s.RerunTimes = &v
	return s
}

func (s *UpdateTaskAsyncShrinkRequest) SetRuntimeResourceShrink(v string) *UpdateTaskAsyncShrinkRequest {
	s.RuntimeResourceShrink = &v
	return s
}

func (s *UpdateTaskAsyncShrinkRequest) SetScriptShrink(v string) *UpdateTaskAsyncShrinkRequest {
	s.ScriptShrink = &v
	return s
}

func (s *UpdateTaskAsyncShrinkRequest) SetTagsShrink(v string) *UpdateTaskAsyncShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *UpdateTaskAsyncShrinkRequest) SetTimeout(v int32) *UpdateTaskAsyncShrinkRequest {
	s.Timeout = &v
	return s
}

func (s *UpdateTaskAsyncShrinkRequest) SetTriggerShrink(v string) *UpdateTaskAsyncShrinkRequest {
	s.TriggerShrink = &v
	return s
}

func (s *UpdateTaskAsyncShrinkRequest) Validate() error {
	return dara.Validate(s)
}
