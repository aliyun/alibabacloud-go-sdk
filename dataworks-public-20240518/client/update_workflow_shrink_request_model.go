// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientUniqueCode(v string) *UpdateWorkflowShrinkRequest
	GetClientUniqueCode() *string
	SetDependenciesShrink(v string) *UpdateWorkflowShrinkRequest
	GetDependenciesShrink() *string
	SetDescription(v string) *UpdateWorkflowShrinkRequest
	GetDescription() *string
	SetEnvType(v string) *UpdateWorkflowShrinkRequest
	GetEnvType() *string
	SetId(v int64) *UpdateWorkflowShrinkRequest
	GetId() *int64
	SetInstanceMode(v string) *UpdateWorkflowShrinkRequest
	GetInstanceMode() *string
	SetName(v string) *UpdateWorkflowShrinkRequest
	GetName() *string
	SetOutputsShrink(v string) *UpdateWorkflowShrinkRequest
	GetOutputsShrink() *string
	SetOwner(v string) *UpdateWorkflowShrinkRequest
	GetOwner() *string
	SetParameters(v string) *UpdateWorkflowShrinkRequest
	GetParameters() *string
	SetTagsShrink(v string) *UpdateWorkflowShrinkRequest
	GetTagsShrink() *string
	SetTasksShrink(v string) *UpdateWorkflowShrinkRequest
	GetTasksShrink() *string
	SetTriggerShrink(v string) *UpdateWorkflowShrinkRequest
	GetTriggerShrink() *string
}

type UpdateWorkflowShrinkRequest struct {
	// The client unique code of the workflow, used for asynchronous operations and idempotence. If not specified during creation, the system automatically generates one, and the code is uniquely bound to the resource ID. If this parameter is specified during update or deletion, it must be consistent with the client unique code used during creation.
	//
	// example:
	//
	// Workflow_0bc5213917368545132902xxxxxxxx
	ClientUniqueCode *string `json:"ClientUniqueCode,omitempty" xml:"ClientUniqueCode,omitempty"`
	// The dependency information.
	DependenciesShrink *string `json:"Dependencies,omitempty" xml:"Dependencies,omitempty"`
	// The description.
	//
	// example:
	//
	// test
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
	// The workflow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The instance generation mode. Valid values:
	//
	// - T+1: Instances are generated the next day.
	//
	// - Immediately: Instances are generated immediately. Periodic instances are generated only if the scheduled time of the workflow is at least 10 minutes after the workflow is published. During the full instance generation period (22:00 to 24:00), real-time instance generation is not available. You can submit and publish workflows during this period, but instances are not regenerated after submission.
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
	// My Workflow
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output information.
	OutputsShrink *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// The account ID of the owner.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The parameter list.
	//
	// example:
	//
	// para1=$bizdate para2=$[yyyymmdd]
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The list of workflow tags.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The node list.
	TasksShrink *string `json:"Tasks,omitempty" xml:"Tasks,omitempty"`
	// The trigger configuration.
	//
	// This parameter is required.
	TriggerShrink *string `json:"Trigger,omitempty" xml:"Trigger,omitempty"`
}

func (s UpdateWorkflowShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowShrinkRequest) GetClientUniqueCode() *string {
	return s.ClientUniqueCode
}

func (s *UpdateWorkflowShrinkRequest) GetDependenciesShrink() *string {
	return s.DependenciesShrink
}

func (s *UpdateWorkflowShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateWorkflowShrinkRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *UpdateWorkflowShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateWorkflowShrinkRequest) GetInstanceMode() *string {
	return s.InstanceMode
}

func (s *UpdateWorkflowShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateWorkflowShrinkRequest) GetOutputsShrink() *string {
	return s.OutputsShrink
}

func (s *UpdateWorkflowShrinkRequest) GetOwner() *string {
	return s.Owner
}

func (s *UpdateWorkflowShrinkRequest) GetParameters() *string {
	return s.Parameters
}

func (s *UpdateWorkflowShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *UpdateWorkflowShrinkRequest) GetTasksShrink() *string {
	return s.TasksShrink
}

func (s *UpdateWorkflowShrinkRequest) GetTriggerShrink() *string {
	return s.TriggerShrink
}

func (s *UpdateWorkflowShrinkRequest) SetClientUniqueCode(v string) *UpdateWorkflowShrinkRequest {
	s.ClientUniqueCode = &v
	return s
}

func (s *UpdateWorkflowShrinkRequest) SetDependenciesShrink(v string) *UpdateWorkflowShrinkRequest {
	s.DependenciesShrink = &v
	return s
}

func (s *UpdateWorkflowShrinkRequest) SetDescription(v string) *UpdateWorkflowShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateWorkflowShrinkRequest) SetEnvType(v string) *UpdateWorkflowShrinkRequest {
	s.EnvType = &v
	return s
}

func (s *UpdateWorkflowShrinkRequest) SetId(v int64) *UpdateWorkflowShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateWorkflowShrinkRequest) SetInstanceMode(v string) *UpdateWorkflowShrinkRequest {
	s.InstanceMode = &v
	return s
}

func (s *UpdateWorkflowShrinkRequest) SetName(v string) *UpdateWorkflowShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateWorkflowShrinkRequest) SetOutputsShrink(v string) *UpdateWorkflowShrinkRequest {
	s.OutputsShrink = &v
	return s
}

func (s *UpdateWorkflowShrinkRequest) SetOwner(v string) *UpdateWorkflowShrinkRequest {
	s.Owner = &v
	return s
}

func (s *UpdateWorkflowShrinkRequest) SetParameters(v string) *UpdateWorkflowShrinkRequest {
	s.Parameters = &v
	return s
}

func (s *UpdateWorkflowShrinkRequest) SetTagsShrink(v string) *UpdateWorkflowShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *UpdateWorkflowShrinkRequest) SetTasksShrink(v string) *UpdateWorkflowShrinkRequest {
	s.TasksShrink = &v
	return s
}

func (s *UpdateWorkflowShrinkRequest) SetTriggerShrink(v string) *UpdateWorkflowShrinkRequest {
	s.TriggerShrink = &v
	return s
}

func (s *UpdateWorkflowShrinkRequest) Validate() error {
	return dara.Validate(s)
}
