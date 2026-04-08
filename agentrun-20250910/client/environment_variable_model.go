// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnvironmentVariable interface {
  dara.Model
  String() string
  GoString() string
  SetDescription(v string) *EnvironmentVariable
  GetDescription() *string 
  SetName(v string) *EnvironmentVariable
  GetName() *string 
  SetValue(v string) *EnvironmentVariable
  GetValue() *string 
}

type EnvironmentVariable struct {
  // 环境变量的描述信息
  // 
  // example:
  // 
  // A configuration key
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // 环境变量的名称
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // MY_KEY
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // 环境变量的值
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // my-value
  Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s EnvironmentVariable) String() string {
  return dara.Prettify(s)
}

func (s EnvironmentVariable) GoString() string {
  return s.String()
}

func (s *EnvironmentVariable) GetDescription() *string  {
  return s.Description
}

func (s *EnvironmentVariable) GetName() *string  {
  return s.Name
}

func (s *EnvironmentVariable) GetValue() *string  {
  return s.Value
}

func (s *EnvironmentVariable) SetDescription(v string) *EnvironmentVariable {
  s.Description = &v
  return s
}

func (s *EnvironmentVariable) SetName(v string) *EnvironmentVariable {
  s.Name = &v
  return s
}

func (s *EnvironmentVariable) SetValue(v string) *EnvironmentVariable {
  s.Value = &v
  return s
}

func (s *EnvironmentVariable) Validate() error {
  return dara.Validate(s)
}

