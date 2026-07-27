// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntityFilters interface {
  dara.Model
  String() string
  GoString() string
  SetField(v string) *EntityFilters
  GetField() *string 
  SetOperator(v string) *EntityFilters
  GetOperator() *string 
  SetValue(v string) *EntityFilters
  GetValue() *string 
}

type EntityFilters struct {
  Field *string `json:"field,omitempty" xml:"field,omitempty"`
  Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
  Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s EntityFilters) String() string {
  return dara.Prettify(s)
}

func (s EntityFilters) GoString() string {
  return s.String()
}

func (s *EntityFilters) GetField() *string  {
  return s.Field
}

func (s *EntityFilters) GetOperator() *string  {
  return s.Operator
}

func (s *EntityFilters) GetValue() *string  {
  return s.Value
}

func (s *EntityFilters) SetField(v string) *EntityFilters {
  s.Field = &v
  return s
}

func (s *EntityFilters) SetOperator(v string) *EntityFilters {
  s.Operator = &v
  return s
}

func (s *EntityFilters) SetValue(v string) *EntityFilters {
  s.Value = &v
  return s
}

func (s *EntityFilters) Validate() error {
  return dara.Validate(s)
}

