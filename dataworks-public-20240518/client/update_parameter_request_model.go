// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateParameterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateParameterRequest
	GetDescription() *string
	SetId(v int64) *UpdateParameterRequest
	GetId() *int64
	SetOwner(v string) *UpdateParameterRequest
	GetOwner() *string
	SetProperties(v []*UpdateParameterRequestProperties) *UpdateParameterRequest
	GetProperties() []*UpdateParameterRequestProperties
}

type UpdateParameterRequest struct {
	// example:
	//
	// 这是一个测试参数
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 123456789
	Owner      *string                             `json:"Owner,omitempty" xml:"Owner,omitempty"`
	Properties []*UpdateParameterRequestProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Repeated"`
}

func (s UpdateParameterRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateParameterRequest) GoString() string {
	return s.String()
}

func (s *UpdateParameterRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateParameterRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateParameterRequest) GetOwner() *string {
	return s.Owner
}

func (s *UpdateParameterRequest) GetProperties() []*UpdateParameterRequestProperties {
	return s.Properties
}

func (s *UpdateParameterRequest) SetDescription(v string) *UpdateParameterRequest {
	s.Description = &v
	return s
}

func (s *UpdateParameterRequest) SetId(v int64) *UpdateParameterRequest {
	s.Id = &v
	return s
}

func (s *UpdateParameterRequest) SetOwner(v string) *UpdateParameterRequest {
	s.Owner = &v
	return s
}

func (s *UpdateParameterRequest) SetProperties(v []*UpdateParameterRequestProperties) *UpdateParameterRequest {
	s.Properties = v
	return s
}

func (s *UpdateParameterRequest) Validate() error {
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

type UpdateParameterRequestProperties struct {
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// example:
	//
	// value123
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateParameterRequestProperties) String() string {
	return dara.Prettify(s)
}

func (s UpdateParameterRequestProperties) GoString() string {
	return s.String()
}

func (s *UpdateParameterRequestProperties) GetEnvType() *string {
	return s.EnvType
}

func (s *UpdateParameterRequestProperties) GetValue() *string {
	return s.Value
}

func (s *UpdateParameterRequestProperties) SetEnvType(v string) *UpdateParameterRequestProperties {
	s.EnvType = &v
	return s
}

func (s *UpdateParameterRequestProperties) SetValue(v string) *UpdateParameterRequestProperties {
	s.Value = &v
	return s
}

func (s *UpdateParameterRequestProperties) Validate() error {
	return dara.Validate(s)
}
