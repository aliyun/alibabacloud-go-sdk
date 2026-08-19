// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnvironmentVariablesValue interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *EnvironmentVariablesValue
  GetType() *string 
  SetValue(v string) *EnvironmentVariablesValue
  GetValue() *string 
}

type EnvironmentVariablesValue struct {
  // The environment variable type.
  // 
  // Valid values:
  // 
  // - `plain_text`: plain text
  // 
  // - `secret_text`: encrypted text
  // 
  // example:
  // 
  // plain_text
  Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
  // The environment variable value.
  // 
  // example:
  // 
  // value
  Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s EnvironmentVariablesValue) String() string {
  return dara.Prettify(s)
}

func (s EnvironmentVariablesValue) GoString() string {
  return s.String()
}

func (s *EnvironmentVariablesValue) GetType() *string  {
  return s.Type
}

func (s *EnvironmentVariablesValue) GetValue() *string  {
  return s.Value
}

func (s *EnvironmentVariablesValue) SetType(v string) *EnvironmentVariablesValue {
  s.Type = &v
  return s
}

func (s *EnvironmentVariablesValue) SetValue(v string) *EnvironmentVariablesValue {
  s.Value = &v
  return s
}

func (s *EnvironmentVariablesValue) Validate() error {
  return dara.Validate(s)
}

