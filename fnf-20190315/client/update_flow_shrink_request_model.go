// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlowShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefinition(v string) *UpdateFlowShrinkRequest
	GetDefinition() *string
	SetDescription(v string) *UpdateFlowShrinkRequest
	GetDescription() *string
	SetEnvironmentShrink(v string) *UpdateFlowShrinkRequest
	GetEnvironmentShrink() *string
	SetName(v string) *UpdateFlowShrinkRequest
	GetName() *string
	SetRoleArn(v string) *UpdateFlowShrinkRequest
	GetRoleArn() *string
	SetType(v string) *UpdateFlowShrinkRequest
	GetType() *string
}

type UpdateFlowShrinkRequest struct {
	// The definition of the workflow. The definition must comply with the flow definition language (FDL) syntax. Considering compatibility, the system supports the two workflow definition specifications.
	//
	// >  In the preceding workflow definition example, Name:my_flow_name is the workflow name, which must be consistent with the input parameter Name
	//
	// example:
	//
	// version: v1.0\\ntype: flow\\nname: test\\nsteps:\\n  - type: pass\\n    name: mypass
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// The description of the flow.
	//
	// example:
	//
	// test definition
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EnvironmentShrink *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// The name of the workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// flow
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The Alibaba Cloud resource name (ARN) of the authorized role on which the execution of the flow relies. During the execution of the flow, the flow execution engine assumes the role to call API operations of relevant services.
	//
	// example:
	//
	// acs:ram::${accountID}:${role}
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The type of the flow. Valid value: **FDL**.
	//
	// example:
	//
	// FDL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateFlowShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateFlowShrinkRequest) GetDefinition() *string {
	return s.Definition
}

func (s *UpdateFlowShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateFlowShrinkRequest) GetEnvironmentShrink() *string {
	return s.EnvironmentShrink
}

func (s *UpdateFlowShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateFlowShrinkRequest) GetRoleArn() *string {
	return s.RoleArn
}

func (s *UpdateFlowShrinkRequest) GetType() *string {
	return s.Type
}

func (s *UpdateFlowShrinkRequest) SetDefinition(v string) *UpdateFlowShrinkRequest {
	s.Definition = &v
	return s
}

func (s *UpdateFlowShrinkRequest) SetDescription(v string) *UpdateFlowShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateFlowShrinkRequest) SetEnvironmentShrink(v string) *UpdateFlowShrinkRequest {
	s.EnvironmentShrink = &v
	return s
}

func (s *UpdateFlowShrinkRequest) SetName(v string) *UpdateFlowShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateFlowShrinkRequest) SetRoleArn(v string) *UpdateFlowShrinkRequest {
	s.RoleArn = &v
	return s
}

func (s *UpdateFlowShrinkRequest) SetType(v string) *UpdateFlowShrinkRequest {
	s.Type = &v
	return s
}

func (s *UpdateFlowShrinkRequest) Validate() error {
	return dara.Validate(s)
}
