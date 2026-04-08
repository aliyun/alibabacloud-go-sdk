// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnvironmentConfiguration interface {
  dara.Model
  String() string
  GoString() string
  SetVariables(v []*EnvironmentVariable) *EnvironmentConfiguration
  GetVariables() []*EnvironmentVariable 
}

type EnvironmentConfiguration struct {
  // 环境变量条目列表，每个条目包含 name、value 和可选 description
  // 
  // This parameter is required.
  Variables []*EnvironmentVariable `json:"variables" xml:"variables" type:"Repeated"`
}

func (s EnvironmentConfiguration) String() string {
  return dara.Prettify(s)
}

func (s EnvironmentConfiguration) GoString() string {
  return s.String()
}

func (s *EnvironmentConfiguration) GetVariables() []*EnvironmentVariable  {
  return s.Variables
}

func (s *EnvironmentConfiguration) SetVariables(v []*EnvironmentVariable) *EnvironmentConfiguration {
  s.Variables = v
  return s
}

func (s *EnvironmentConfiguration) Validate() error {
  if s.Variables != nil {
    for _, item := range s.Variables {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

