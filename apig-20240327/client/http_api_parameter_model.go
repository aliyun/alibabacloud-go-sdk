// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHttpApiParameter interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultValue(v string) *HttpApiParameter
	GetDefaultValue() *string
	SetDescription(v string) *HttpApiParameter
	GetDescription() *string
	SetExampleValue(v string) *HttpApiParameter
	GetExampleValue() *string
	SetName(v string) *HttpApiParameter
	GetName() *string
	SetRequired(v bool) *HttpApiParameter
	GetRequired() *bool
	SetType(v string) *HttpApiParameter
	GetType() *string
}

type HttpApiParameter struct {
	DefaultValue *string `json:"defaultValue,omitempty" xml:"defaultValue,omitempty"`
	Description  *string `json:"description,omitempty" xml:"description,omitempty"`
	ExampleValue *string `json:"exampleValue,omitempty" xml:"exampleValue,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// false
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s HttpApiParameter) String() string {
	return dara.Prettify(s)
}

func (s HttpApiParameter) GoString() string {
	return s.String()
}

func (s *HttpApiParameter) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *HttpApiParameter) GetDescription() *string {
	return s.Description
}

func (s *HttpApiParameter) GetExampleValue() *string {
	return s.ExampleValue
}

func (s *HttpApiParameter) GetName() *string {
	return s.Name
}

func (s *HttpApiParameter) GetRequired() *bool {
	return s.Required
}

func (s *HttpApiParameter) GetType() *string {
	return s.Type
}

func (s *HttpApiParameter) SetDefaultValue(v string) *HttpApiParameter {
	s.DefaultValue = &v
	return s
}

func (s *HttpApiParameter) SetDescription(v string) *HttpApiParameter {
	s.Description = &v
	return s
}

func (s *HttpApiParameter) SetExampleValue(v string) *HttpApiParameter {
	s.ExampleValue = &v
	return s
}

func (s *HttpApiParameter) SetName(v string) *HttpApiParameter {
	s.Name = &v
	return s
}

func (s *HttpApiParameter) SetRequired(v bool) *HttpApiParameter {
	s.Required = &v
	return s
}

func (s *HttpApiParameter) SetType(v string) *HttpApiParameter {
	s.Type = &v
	return s
}

func (s *HttpApiParameter) Validate() error {
	return dara.Validate(s)
}
