// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFieldOption interface {
	dara.Model
	String() string
	GoString() string
	SetFunctionCode(v string) *FieldOption
	GetFunctionCode() *string
	SetIsDefault(v bool) *FieldOption
	GetIsDefault() *bool
	SetName(v string) *FieldOption
	GetName() *string
	SetValue(v string) *FieldOption
	GetValue() *string
}

type FieldOption struct {
	FunctionCode *string `json:"FunctionCode,omitempty" xml:"FunctionCode,omitempty"`
	IsDefault    *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s FieldOption) String() string {
	return dara.Prettify(s)
}

func (s FieldOption) GoString() string {
	return s.String()
}

func (s *FieldOption) GetFunctionCode() *string {
	return s.FunctionCode
}

func (s *FieldOption) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *FieldOption) GetName() *string {
	return s.Name
}

func (s *FieldOption) GetValue() *string {
	return s.Value
}

func (s *FieldOption) SetFunctionCode(v string) *FieldOption {
	s.FunctionCode = &v
	return s
}

func (s *FieldOption) SetIsDefault(v bool) *FieldOption {
	s.IsDefault = &v
	return s
}

func (s *FieldOption) SetName(v string) *FieldOption {
	s.Name = &v
	return s
}

func (s *FieldOption) SetValue(v string) *FieldOption {
	s.Value = &v
	return s
}

func (s *FieldOption) Validate() error {
	return dara.Validate(s)
}
