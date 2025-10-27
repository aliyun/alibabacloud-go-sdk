// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefinition(v string) *CreateFlowShrinkRequest
	GetDefinition() *string
	SetDescription(v string) *CreateFlowShrinkRequest
	GetDescription() *string
	SetEnvironmentShrink(v string) *CreateFlowShrinkRequest
	GetEnvironmentShrink() *string
	SetExecutionMode(v string) *CreateFlowShrinkRequest
	GetExecutionMode() *string
	SetExternalStorageLocation(v string) *CreateFlowShrinkRequest
	GetExternalStorageLocation() *string
	SetName(v string) *CreateFlowShrinkRequest
	GetName() *string
	SetRoleArn(v string) *CreateFlowShrinkRequest
	GetRoleArn() *string
	SetType(v string) *CreateFlowShrinkRequest
	GetType() *string
}

type CreateFlowShrinkRequest struct {
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
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EnvironmentShrink *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
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

func (s CreateFlowShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowShrinkRequest) GetDefinition() *string {
	return s.Definition
}

func (s *CreateFlowShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateFlowShrinkRequest) GetEnvironmentShrink() *string {
	return s.EnvironmentShrink
}

func (s *CreateFlowShrinkRequest) GetExecutionMode() *string {
	return s.ExecutionMode
}

func (s *CreateFlowShrinkRequest) GetExternalStorageLocation() *string {
	return s.ExternalStorageLocation
}

func (s *CreateFlowShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateFlowShrinkRequest) GetRoleArn() *string {
	return s.RoleArn
}

func (s *CreateFlowShrinkRequest) GetType() *string {
	return s.Type
}

func (s *CreateFlowShrinkRequest) SetDefinition(v string) *CreateFlowShrinkRequest {
	s.Definition = &v
	return s
}

func (s *CreateFlowShrinkRequest) SetDescription(v string) *CreateFlowShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateFlowShrinkRequest) SetEnvironmentShrink(v string) *CreateFlowShrinkRequest {
	s.EnvironmentShrink = &v
	return s
}

func (s *CreateFlowShrinkRequest) SetExecutionMode(v string) *CreateFlowShrinkRequest {
	s.ExecutionMode = &v
	return s
}

func (s *CreateFlowShrinkRequest) SetExternalStorageLocation(v string) *CreateFlowShrinkRequest {
	s.ExternalStorageLocation = &v
	return s
}

func (s *CreateFlowShrinkRequest) SetName(v string) *CreateFlowShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateFlowShrinkRequest) SetRoleArn(v string) *CreateFlowShrinkRequest {
	s.RoleArn = &v
	return s
}

func (s *CreateFlowShrinkRequest) SetType(v string) *CreateFlowShrinkRequest {
	s.Type = &v
	return s
}

func (s *CreateFlowShrinkRequest) Validate() error {
	return dara.Validate(s)
}
