// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateParameterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateParameterRequest
	GetDescription() *string
	SetName(v string) *CreateParameterRequest
	GetName() *string
	SetOwner(v string) *CreateParameterRequest
	GetOwner() *string
	SetProjectId(v int64) *CreateParameterRequest
	GetProjectId() *int64
	SetProperties(v []*CreateParameterRequestProperties) *CreateParameterRequest
	GetProperties() []*CreateParameterRequestProperties
	SetScope(v string) *CreateParameterRequest
	GetScope() *string
	SetType(v string) *CreateParameterRequest
	GetType() *string
}

type CreateParameterRequest struct {
	// The description of the parameter.
	//
	// example:
	//
	// This is a test parameter.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameter name. It must be unique within the workspace, be prefixed with `workspace.`, and not exceed 255 characters. The part of the name after the prefix must start with a letter and can contain only letters, digits, and underscores (_).
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
	// The workspace ID. This parameter is required when `Scope` is set to `Project`.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The value configurations for the parameter. A configuration for the production environment is required. If you provide duplicate configurations for an environment, only the first one is used.
	//
	// This parameter is required.
	Properties []*CreateParameterRequestProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Repeated"`
	// The scope of the parameter. The default value is `Project`. No other values are currently supported.
	//
	// example:
	//
	// Project
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The type of the parameter.
	//
	// - `PlainConstant`: plaintext constant.
	//
	// - `SecretConstant`: secret constant.
	//
	// - `Variable`: variable.
	//
	// This parameter is required.
	//
	// example:
	//
	// PlainConstant
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateParameterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateParameterRequest) GoString() string {
	return s.String()
}

func (s *CreateParameterRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateParameterRequest) GetName() *string {
	return s.Name
}

func (s *CreateParameterRequest) GetOwner() *string {
	return s.Owner
}

func (s *CreateParameterRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateParameterRequest) GetProperties() []*CreateParameterRequestProperties {
	return s.Properties
}

func (s *CreateParameterRequest) GetScope() *string {
	return s.Scope
}

func (s *CreateParameterRequest) GetType() *string {
	return s.Type
}

func (s *CreateParameterRequest) SetDescription(v string) *CreateParameterRequest {
	s.Description = &v
	return s
}

func (s *CreateParameterRequest) SetName(v string) *CreateParameterRequest {
	s.Name = &v
	return s
}

func (s *CreateParameterRequest) SetOwner(v string) *CreateParameterRequest {
	s.Owner = &v
	return s
}

func (s *CreateParameterRequest) SetProjectId(v int64) *CreateParameterRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateParameterRequest) SetProperties(v []*CreateParameterRequestProperties) *CreateParameterRequest {
	s.Properties = v
	return s
}

func (s *CreateParameterRequest) SetScope(v string) *CreateParameterRequest {
	s.Scope = &v
	return s
}

func (s *CreateParameterRequest) SetType(v string) *CreateParameterRequest {
	s.Type = &v
	return s
}

func (s *CreateParameterRequest) Validate() error {
	if s.Properties != nil {
		for _, item := range s.Properties {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateParameterRequestProperties struct {
	// The environment.
	//
	// - `Prod`: production environment
	//
	// - `Dev`: development environment
	//
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The value of the parameter. The value can contain Chinese characters, letters, digits, and the following special characters: /, :, ., [, ], ,, \\, \\", ", _, =, ?, space, carriage return, line feed, +, -, \\*, %, &, @, !, $, #, {, and }.
	//
	// example:
	//
	// value123
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateParameterRequestProperties) String() string {
	return dara.Prettify(s)
}

func (s CreateParameterRequestProperties) GoString() string {
	return s.String()
}

func (s *CreateParameterRequestProperties) GetEnvType() *string {
	return s.EnvType
}

func (s *CreateParameterRequestProperties) GetValue() *string {
	return s.Value
}

func (s *CreateParameterRequestProperties) SetEnvType(v string) *CreateParameterRequestProperties {
	s.EnvType = &v
	return s
}

func (s *CreateParameterRequestProperties) SetValue(v string) *CreateParameterRequestProperties {
	s.Value = &v
	return s
}

func (s *CreateParameterRequestProperties) Validate() error {
	return dara.Validate(s)
}
