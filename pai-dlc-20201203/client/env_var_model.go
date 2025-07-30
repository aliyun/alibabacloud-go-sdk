// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnvVar interface {
  dara.Model
  String() string
  GoString() string
  SetName(v string) *EnvVar
  GetName() *string 
  SetValue(v string) *EnvVar
  GetValue() *string 
}

type EnvVar struct {
  // example:
  // 
  // ENABLE_DEBUG
  Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
  // example:
  // 
  // true
  Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s EnvVar) String() string {
  return dara.Prettify(s)
}

func (s EnvVar) GoString() string {
  return s.String()
}

func (s *EnvVar) GetName() *string  {
  return s.Name
}

func (s *EnvVar) GetValue() *string  {
  return s.Value
}

func (s *EnvVar) SetName(v string) *EnvVar {
  s.Name = &v
  return s
}

func (s *EnvVar) SetValue(v string) *EnvVar {
  s.Value = &v
  return s
}

func (s *EnvVar) Validate() error {
  return dara.Validate(s)
}

