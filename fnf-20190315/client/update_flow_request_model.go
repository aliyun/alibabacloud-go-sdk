// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefinition(v string) *UpdateFlowRequest
	GetDefinition() *string
	SetDescription(v string) *UpdateFlowRequest
	GetDescription() *string
	SetEnvironment(v *UpdateFlowRequestEnvironment) *UpdateFlowRequest
	GetEnvironment() *UpdateFlowRequestEnvironment
	SetName(v string) *UpdateFlowRequest
	GetName() *string
	SetRoleArn(v string) *UpdateFlowRequest
	GetRoleArn() *string
	SetType(v string) *UpdateFlowRequest
	GetType() *string
}

type UpdateFlowRequest struct {
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
	Description *string                       `json:"Description,omitempty" xml:"Description,omitempty"`
	Environment *UpdateFlowRequestEnvironment `json:"Environment,omitempty" xml:"Environment,omitempty" type:"Struct"`
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

func (s UpdateFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowRequest) GoString() string {
	return s.String()
}

func (s *UpdateFlowRequest) GetDefinition() *string {
	return s.Definition
}

func (s *UpdateFlowRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateFlowRequest) GetEnvironment() *UpdateFlowRequestEnvironment {
	return s.Environment
}

func (s *UpdateFlowRequest) GetName() *string {
	return s.Name
}

func (s *UpdateFlowRequest) GetRoleArn() *string {
	return s.RoleArn
}

func (s *UpdateFlowRequest) GetType() *string {
	return s.Type
}

func (s *UpdateFlowRequest) SetDefinition(v string) *UpdateFlowRequest {
	s.Definition = &v
	return s
}

func (s *UpdateFlowRequest) SetDescription(v string) *UpdateFlowRequest {
	s.Description = &v
	return s
}

func (s *UpdateFlowRequest) SetEnvironment(v *UpdateFlowRequestEnvironment) *UpdateFlowRequest {
	s.Environment = v
	return s
}

func (s *UpdateFlowRequest) SetName(v string) *UpdateFlowRequest {
	s.Name = &v
	return s
}

func (s *UpdateFlowRequest) SetRoleArn(v string) *UpdateFlowRequest {
	s.RoleArn = &v
	return s
}

func (s *UpdateFlowRequest) SetType(v string) *UpdateFlowRequest {
	s.Type = &v
	return s
}

func (s *UpdateFlowRequest) Validate() error {
	if s.Environment != nil {
		if err := s.Environment.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateFlowRequestEnvironment struct {
	Variables []*UpdateFlowRequestEnvironmentVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s UpdateFlowRequestEnvironment) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowRequestEnvironment) GoString() string {
	return s.String()
}

func (s *UpdateFlowRequestEnvironment) GetVariables() []*UpdateFlowRequestEnvironmentVariables {
	return s.Variables
}

func (s *UpdateFlowRequestEnvironment) SetVariables(v []*UpdateFlowRequestEnvironmentVariables) *UpdateFlowRequestEnvironment {
	s.Variables = v
	return s
}

func (s *UpdateFlowRequestEnvironment) Validate() error {
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

type UpdateFlowRequestEnvironmentVariables struct {
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// key
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateFlowRequestEnvironmentVariables) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowRequestEnvironmentVariables) GoString() string {
	return s.String()
}

func (s *UpdateFlowRequestEnvironmentVariables) GetDescription() *string {
	return s.Description
}

func (s *UpdateFlowRequestEnvironmentVariables) GetName() *string {
	return s.Name
}

func (s *UpdateFlowRequestEnvironmentVariables) GetValue() *string {
	return s.Value
}

func (s *UpdateFlowRequestEnvironmentVariables) SetDescription(v string) *UpdateFlowRequestEnvironmentVariables {
	s.Description = &v
	return s
}

func (s *UpdateFlowRequestEnvironmentVariables) SetName(v string) *UpdateFlowRequestEnvironmentVariables {
	s.Name = &v
	return s
}

func (s *UpdateFlowRequestEnvironmentVariables) SetValue(v string) *UpdateFlowRequestEnvironmentVariables {
	s.Value = &v
	return s
}

func (s *UpdateFlowRequestEnvironmentVariables) Validate() error {
	return dara.Validate(s)
}
