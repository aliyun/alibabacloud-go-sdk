// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateParameterShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateParameterShrinkRequest
	GetDescription() *string
	SetName(v string) *CreateParameterShrinkRequest
	GetName() *string
	SetOwner(v string) *CreateParameterShrinkRequest
	GetOwner() *string
	SetProjectId(v int64) *CreateParameterShrinkRequest
	GetProjectId() *int64
	SetPropertiesShrink(v string) *CreateParameterShrinkRequest
	GetPropertiesShrink() *string
	SetScope(v string) *CreateParameterShrinkRequest
	GetScope() *string
	SetType(v string) *CreateParameterShrinkRequest
	GetType() *string
}

type CreateParameterShrinkRequest struct {
	// The description.
	//
	// example:
	//
	// This is a test parameter
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the parameter. It must be unique within the workspace. It must start with "workspace.". The subsequent content must start with a letter and can contain only letters, underscores, and digits. The total length of the name cannot exceed 255 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// workspace.para
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The account ID of the owner.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The workspace ID. This parameter is required when Scope is set to Project.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The parameter value configurations. This parameter is required for the production environment. If duplicate environment configurations exist in the array, the subsequent configurations do not take effect.
	//
	// This parameter is required.
	PropertiesShrink *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// The scope of the parameter. Default value: Project. Other types are currently not supported.
	//
	// example:
	//
	// Project
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The type. Valid values:
	//
	// - PlainConstant: plaintext constant.
	//
	// - SecretConstant: ciphertext constant.
	//
	// - Variable: variable.
	//
	// This parameter is required.
	//
	// example:
	//
	// PlainConstant
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateParameterShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateParameterShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateParameterShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateParameterShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateParameterShrinkRequest) GetOwner() *string {
	return s.Owner
}

func (s *CreateParameterShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateParameterShrinkRequest) GetPropertiesShrink() *string {
	return s.PropertiesShrink
}

func (s *CreateParameterShrinkRequest) GetScope() *string {
	return s.Scope
}

func (s *CreateParameterShrinkRequest) GetType() *string {
	return s.Type
}

func (s *CreateParameterShrinkRequest) SetDescription(v string) *CreateParameterShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateParameterShrinkRequest) SetName(v string) *CreateParameterShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateParameterShrinkRequest) SetOwner(v string) *CreateParameterShrinkRequest {
	s.Owner = &v
	return s
}

func (s *CreateParameterShrinkRequest) SetProjectId(v int64) *CreateParameterShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateParameterShrinkRequest) SetPropertiesShrink(v string) *CreateParameterShrinkRequest {
	s.PropertiesShrink = &v
	return s
}

func (s *CreateParameterShrinkRequest) SetScope(v string) *CreateParameterShrinkRequest {
	s.Scope = &v
	return s
}

func (s *CreateParameterShrinkRequest) SetType(v string) *CreateParameterShrinkRequest {
	s.Type = &v
	return s
}

func (s *CreateParameterShrinkRequest) Validate() error {
	return dara.Validate(s)
}
