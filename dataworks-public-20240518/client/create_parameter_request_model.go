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
	// example:
	//
	// 这是一个测试参数
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// workspace.para
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456789
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	Properties []*CreateParameterRequestProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Repeated"`
	// example:
	//
	// Project
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
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
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
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
