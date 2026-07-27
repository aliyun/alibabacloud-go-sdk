// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntityFields interface {
  dara.Model
  String() string
  GoString() string
  SetField(v string) *EntityFields
  GetField() *string 
  SetValue(v string) *EntityFields
  GetValue() *string 
}

type EntityFields struct {
  Field *string `json:"field,omitempty" xml:"field,omitempty"`
  Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s EntityFields) String() string {
  return dara.Prettify(s)
}

func (s EntityFields) GoString() string {
  return s.String()
}

func (s *EntityFields) GetField() *string  {
  return s.Field
}

func (s *EntityFields) GetValue() *string  {
  return s.Value
}

func (s *EntityFields) SetField(v string) *EntityFields {
  s.Field = &v
  return s
}

func (s *EntityFields) SetValue(v string) *EntityFields {
  s.Value = &v
  return s
}

func (s *EntityFields) Validate() error {
  return dara.Validate(s)
}

