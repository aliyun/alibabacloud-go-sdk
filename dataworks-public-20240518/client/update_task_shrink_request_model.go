// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientUniqueCode(v string) *UpdateTaskShrinkRequest
	GetClientUniqueCode() *string
	SetDataSourceShrink(v string) *UpdateTaskShrinkRequest
	GetDataSourceShrink() *string
	SetDependenciesShrink(v string) *UpdateTaskShrinkRequest
	GetDependenciesShrink() *string
	SetDescription(v string) *UpdateTaskShrinkRequest
	GetDescription() *string
	SetEnvType(v string) *UpdateTaskShrinkRequest
	GetEnvType() *string
	SetId(v int64) *UpdateTaskShrinkRequest
	GetId() *int64
	SetInputsShrink(v string) *UpdateTaskShrinkRequest
	GetInputsShrink() *string
	SetInstanceMode(v string) *UpdateTaskShrinkRequest
	GetInstanceMode() *string
	SetName(v string) *UpdateTaskShrinkRequest
	GetName() *string
	SetOutputsShrink(v string) *UpdateTaskShrinkRequest
	GetOutputsShrink() *string
	SetOwner(v string) *UpdateTaskShrinkRequest
	GetOwner() *string
	SetRerunInterval(v int32) *UpdateTaskShrinkRequest
	GetRerunInterval() *int32
	SetRerunMode(v string) *UpdateTaskShrinkRequest
	GetRerunMode() *string
	SetRerunTimes(v int32) *UpdateTaskShrinkRequest
	GetRerunTimes() *int32
	SetRuntimeResourceShrink(v string) *UpdateTaskShrinkRequest
	GetRuntimeResourceShrink() *string
	SetScriptShrink(v string) *UpdateTaskShrinkRequest
	GetScriptShrink() *string
	SetTagsShrink(v string) *UpdateTaskShrinkRequest
	GetTagsShrink() *string
	SetTimeout(v int32) *UpdateTaskShrinkRequest
	GetTimeout() *int32
	SetTriggerShrink(v string) *UpdateTaskShrinkRequest
	GetTriggerShrink() *string
}

type UpdateTaskShrinkRequest struct {
	// The unique code of the client. This code uniquely identifies a task. This parameter is used to create a task asynchronously and implement the idempotence of the task. If you do not specify this parameter when you create the task, the system automatically generates a unique code. The unique code is uniquely associated with the task ID. If you specify this parameter when you update or delete the task, the value of this parameter must be the unique code that is used to create the task.
	//
	// example:
	//
	// Task_0bc5213917368545132902xxxxxxxx
	ClientUniqueCode *string `json:"ClientUniqueCode,omitempty" xml:"ClientUniqueCode,omitempty"`
	// The information about the associated data source.
	DataSourceShrink *string `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	// The dependency information.
	DependenciesShrink *string `json:"Dependencies,omitempty" xml:"Dependencies,omitempty"`
	// The description of the task.
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
	// The input information.
	InputsShrink *string `json:"Inputs,omitempty" xml:"Inputs,omitempty"`
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
	// The name.
	//
	// This parameter is required.
	//
	// example:
	//
	// SQL node
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output information.
	OutputsShrink *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// The account ID of the task owner.
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
	RuntimeResourceShrink *string `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty"`
	// The script information.
	ScriptShrink *string `json:"Script,omitempty" xml:"Script,omitempty"`
	// The tags.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The timeout period of task running. Unit: seconds.
	//
	// example:
	//
	// 3600
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The trigger method.
	//
	// This parameter is required.
	TriggerShrink *string `json:"Trigger,omitempty" xml:"Trigger,omitempty"`
}

func (s UpdateTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskShrinkRequest) GetClientUniqueCode() *string {
	return s.ClientUniqueCode
}

func (s *UpdateTaskShrinkRequest) GetDataSourceShrink() *string {
	return s.DataSourceShrink
}

func (s *UpdateTaskShrinkRequest) GetDependenciesShrink() *string {
	return s.DependenciesShrink
}

func (s *UpdateTaskShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateTaskShrinkRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *UpdateTaskShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateTaskShrinkRequest) GetInputsShrink() *string {
	return s.InputsShrink
}

func (s *UpdateTaskShrinkRequest) GetInstanceMode() *string {
	return s.InstanceMode
}

func (s *UpdateTaskShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateTaskShrinkRequest) GetOutputsShrink() *string {
	return s.OutputsShrink
}

func (s *UpdateTaskShrinkRequest) GetOwner() *string {
	return s.Owner
}

func (s *UpdateTaskShrinkRequest) GetRerunInterval() *int32 {
	return s.RerunInterval
}

func (s *UpdateTaskShrinkRequest) GetRerunMode() *string {
	return s.RerunMode
}

func (s *UpdateTaskShrinkRequest) GetRerunTimes() *int32 {
	return s.RerunTimes
}

func (s *UpdateTaskShrinkRequest) GetRuntimeResourceShrink() *string {
	return s.RuntimeResourceShrink
}

func (s *UpdateTaskShrinkRequest) GetScriptShrink() *string {
	return s.ScriptShrink
}

func (s *UpdateTaskShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *UpdateTaskShrinkRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateTaskShrinkRequest) GetTriggerShrink() *string {
	return s.TriggerShrink
}

func (s *UpdateTaskShrinkRequest) SetClientUniqueCode(v string) *UpdateTaskShrinkRequest {
	s.ClientUniqueCode = &v
	return s
}

func (s *UpdateTaskShrinkRequest) SetDataSourceShrink(v string) *UpdateTaskShrinkRequest {
	s.DataSourceShrink = &v
	return s
}

func (s *UpdateTaskShrinkRequest) SetDependenciesShrink(v string) *UpdateTaskShrinkRequest {
	s.DependenciesShrink = &v
	return s
}

func (s *UpdateTaskShrinkRequest) SetDescription(v string) *UpdateTaskShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateTaskShrinkRequest) SetEnvType(v string) *UpdateTaskShrinkRequest {
	s.EnvType = &v
	return s
}

func (s *UpdateTaskShrinkRequest) SetId(v int64) *UpdateTaskShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateTaskShrinkRequest) SetInputsShrink(v string) *UpdateTaskShrinkRequest {
	s.InputsShrink = &v
	return s
}

func (s *UpdateTaskShrinkRequest) SetInstanceMode(v string) *UpdateTaskShrinkRequest {
	s.InstanceMode = &v
	return s
}

func (s *UpdateTaskShrinkRequest) SetName(v string) *UpdateTaskShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateTaskShrinkRequest) SetOutputsShrink(v string) *UpdateTaskShrinkRequest {
	s.OutputsShrink = &v
	return s
}

func (s *UpdateTaskShrinkRequest) SetOwner(v string) *UpdateTaskShrinkRequest {
	s.Owner = &v
	return s
}

func (s *UpdateTaskShrinkRequest) SetRerunInterval(v int32) *UpdateTaskShrinkRequest {
	s.RerunInterval = &v
	return s
}

func (s *UpdateTaskShrinkRequest) SetRerunMode(v string) *UpdateTaskShrinkRequest {
	s.RerunMode = &v
	return s
}

func (s *UpdateTaskShrinkRequest) SetRerunTimes(v int32) *UpdateTaskShrinkRequest {
	s.RerunTimes = &v
	return s
}

func (s *UpdateTaskShrinkRequest) SetRuntimeResourceShrink(v string) *UpdateTaskShrinkRequest {
	s.RuntimeResourceShrink = &v
	return s
}

func (s *UpdateTaskShrinkRequest) SetScriptShrink(v string) *UpdateTaskShrinkRequest {
	s.ScriptShrink = &v
	return s
}

func (s *UpdateTaskShrinkRequest) SetTagsShrink(v string) *UpdateTaskShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *UpdateTaskShrinkRequest) SetTimeout(v int32) *UpdateTaskShrinkRequest {
	s.Timeout = &v
	return s
}

func (s *UpdateTaskShrinkRequest) SetTriggerShrink(v string) *UpdateTaskShrinkRequest {
	s.TriggerShrink = &v
	return s
}

func (s *UpdateTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
