// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefinition(v string) *CreateFlowRequest
	GetDefinition() *string
	SetDescription(v string) *CreateFlowRequest
	GetDescription() *string
	SetEnvironment(v *CreateFlowRequestEnvironment) *CreateFlowRequest
	GetEnvironment() *CreateFlowRequestEnvironment
	SetExecutionMode(v string) *CreateFlowRequest
	GetExecutionMode() *string
	SetExternalStorageLocation(v string) *CreateFlowRequest
	GetExternalStorageLocation() *string
	SetName(v string) *CreateFlowRequest
	GetName() *string
	SetRoleArn(v string) *CreateFlowRequest
	GetRoleArn() *string
	SetType(v string) *CreateFlowRequest
	GetType() *string
}

type CreateFlowRequest struct {
	// The definition of the workflow. The definition must comply with the flow definition language (FDL) syntax. Considering compatibility, the system supports two flow definition specifications.
	//
	// >  In the preceding flow definition example, Name:my_flow_name is the workflow name, which must be consistent with the input parameter Name
	//
	// This parameter is required.
	//
	// example:
	//
	// version:&nbsp;v1.0<br/>type:&nbsp;flow<br/>steps:<br/>&nbsp;-&nbsp;type:&nbsp;pass<br/>&nbsp;name:&nbsp;mypass
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// The description of the flow.
	//
	// This parameter is required.
	//
	// example:
	//
	// test flow
	Description *string                       `json:"Description,omitempty" xml:"Description,omitempty"`
	Environment *CreateFlowRequestEnvironment `json:"Environment,omitempty" xml:"Environment,omitempty" type:"Struct"`
	// The execution mode. Valid values: Express and Standard. Considering compatibility, an empty string is equivalent to the Standard execution mode.
	//
	// example:
	//
	// Standard
	ExecutionMode *string `json:"ExecutionMode,omitempty" xml:"ExecutionMode,omitempty"`
	// The path of the external storage.
	//
	// example:
	//
	// /path
	ExternalStorageLocation *string `json:"ExternalStorageLocation,omitempty" xml:"ExternalStorageLocation,omitempty"`
	// The name of the flow. The name is unique within the same region and cannot be modified after the flow is created. Set this parameter based on the following rules:
	//
	// 	- The name can contain letters, digits, underscores (_), and hyphens (-).
	//
	// 	- The name must start with a letter or an underscore (_).
	//
	// 	- The name is case-sensitive.
	//
	// 	- The name must be 1 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// flow
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The Alibaba Cloud resource name (ARN) of the authorized role on which the execution of the flow relies. During the execution of the flow, CloudFlow assumes the role to call API operations of relevant services.
	//
	// example:
	//
	// acs:ram:${region}:${accountID}:${role}
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The type of the flow. Set this parameter to **FDL**.
	//
	// This parameter is required.
	//
	// example:
	//
	// FDL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowRequest) GetDefinition() *string {
	return s.Definition
}

func (s *CreateFlowRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateFlowRequest) GetEnvironment() *CreateFlowRequestEnvironment {
	return s.Environment
}

func (s *CreateFlowRequest) GetExecutionMode() *string {
	return s.ExecutionMode
}

func (s *CreateFlowRequest) GetExternalStorageLocation() *string {
	return s.ExternalStorageLocation
}

func (s *CreateFlowRequest) GetName() *string {
	return s.Name
}

func (s *CreateFlowRequest) GetRoleArn() *string {
	return s.RoleArn
}

func (s *CreateFlowRequest) GetType() *string {
	return s.Type
}

func (s *CreateFlowRequest) SetDefinition(v string) *CreateFlowRequest {
	s.Definition = &v
	return s
}

func (s *CreateFlowRequest) SetDescription(v string) *CreateFlowRequest {
	s.Description = &v
	return s
}

func (s *CreateFlowRequest) SetEnvironment(v *CreateFlowRequestEnvironment) *CreateFlowRequest {
	s.Environment = v
	return s
}

func (s *CreateFlowRequest) SetExecutionMode(v string) *CreateFlowRequest {
	s.ExecutionMode = &v
	return s
}

func (s *CreateFlowRequest) SetExternalStorageLocation(v string) *CreateFlowRequest {
	s.ExternalStorageLocation = &v
	return s
}

func (s *CreateFlowRequest) SetName(v string) *CreateFlowRequest {
	s.Name = &v
	return s
}

func (s *CreateFlowRequest) SetRoleArn(v string) *CreateFlowRequest {
	s.RoleArn = &v
	return s
}

func (s *CreateFlowRequest) SetType(v string) *CreateFlowRequest {
	s.Type = &v
	return s
}

func (s *CreateFlowRequest) Validate() error {
	if s.Environment != nil {
		if err := s.Environment.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateFlowRequestEnvironment struct {
	Variables []*CreateFlowRequestEnvironmentVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s CreateFlowRequestEnvironment) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowRequestEnvironment) GoString() string {
	return s.String()
}

func (s *CreateFlowRequestEnvironment) GetVariables() []*CreateFlowRequestEnvironmentVariables {
	return s.Variables
}

func (s *CreateFlowRequestEnvironment) SetVariables(v []*CreateFlowRequestEnvironmentVariables) *CreateFlowRequestEnvironment {
	s.Variables = v
	return s
}

func (s *CreateFlowRequestEnvironment) Validate() error {
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

type CreateFlowRequestEnvironmentVariables struct {
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

func (s CreateFlowRequestEnvironmentVariables) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowRequestEnvironmentVariables) GoString() string {
	return s.String()
}

func (s *CreateFlowRequestEnvironmentVariables) GetDescription() *string {
	return s.Description
}

func (s *CreateFlowRequestEnvironmentVariables) GetName() *string {
	return s.Name
}

func (s *CreateFlowRequestEnvironmentVariables) GetValue() *string {
	return s.Value
}

func (s *CreateFlowRequestEnvironmentVariables) SetDescription(v string) *CreateFlowRequestEnvironmentVariables {
	s.Description = &v
	return s
}

func (s *CreateFlowRequestEnvironmentVariables) SetName(v string) *CreateFlowRequestEnvironmentVariables {
	s.Name = &v
	return s
}

func (s *CreateFlowRequestEnvironmentVariables) SetValue(v string) *CreateFlowRequestEnvironmentVariables {
	s.Value = &v
	return s
}

func (s *CreateFlowRequestEnvironmentVariables) Validate() error {
	return dara.Validate(s)
}
