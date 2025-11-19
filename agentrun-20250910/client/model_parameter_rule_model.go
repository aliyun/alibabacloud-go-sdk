// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelParameterRule interface {
	dara.Model
	String() string
	GoString() string
	SetDefault(v interface{}) *ModelParameterRule
	GetDefault() interface{}
	SetMax(v int32) *ModelParameterRule
	GetMax() *int32
	SetMin(v int32) *ModelParameterRule
	GetMin() *int32
	SetName(v string) *ModelParameterRule
	GetName() *string
	SetRequired(v bool) *ModelParameterRule
	GetRequired() *bool
	SetType(v string) *ModelParameterRule
	GetType() *string
}

type ModelParameterRule struct {
	Default  interface{} `json:"default,omitempty" xml:"default,omitempty"`
	Max      *int32      `json:"max,omitempty" xml:"max,omitempty"`
	Min      *int32      `json:"min,omitempty" xml:"min,omitempty"`
	Name     *string     `json:"name,omitempty" xml:"name,omitempty"`
	Required *bool       `json:"required,omitempty" xml:"required,omitempty"`
	Type     *string     `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ModelParameterRule) String() string {
	return dara.Prettify(s)
}

func (s ModelParameterRule) GoString() string {
	return s.String()
}

func (s *ModelParameterRule) GetDefault() interface{} {
	return s.Default
}

func (s *ModelParameterRule) GetMax() *int32 {
	return s.Max
}

func (s *ModelParameterRule) GetMin() *int32 {
	return s.Min
}

func (s *ModelParameterRule) GetName() *string {
	return s.Name
}

func (s *ModelParameterRule) GetRequired() *bool {
	return s.Required
}

func (s *ModelParameterRule) GetType() *string {
	return s.Type
}

func (s *ModelParameterRule) SetDefault(v interface{}) *ModelParameterRule {
	s.Default = v
	return s
}

func (s *ModelParameterRule) SetMax(v int32) *ModelParameterRule {
	s.Max = &v
	return s
}

func (s *ModelParameterRule) SetMin(v int32) *ModelParameterRule {
	s.Min = &v
	return s
}

func (s *ModelParameterRule) SetName(v string) *ModelParameterRule {
	s.Name = &v
	return s
}

func (s *ModelParameterRule) SetRequired(v bool) *ModelParameterRule {
	s.Required = &v
	return s
}

func (s *ModelParameterRule) SetType(v string) *ModelParameterRule {
	s.Type = &v
	return s
}

func (s *ModelParameterRule) Validate() error {
	return dara.Validate(s)
}
