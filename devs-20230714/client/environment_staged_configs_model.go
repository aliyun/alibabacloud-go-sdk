// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnvironmentStagedConfigs interface {
  dara.Model
  String() string
  GoString() string
  SetServices(v map[string]*ServiceConfig) *EnvironmentStagedConfigs
  GetServices() map[string]*ServiceConfig 
  SetVariables(v map[string]*Variable) *EnvironmentStagedConfigs
  GetVariables() map[string]*Variable 
}

type EnvironmentStagedConfigs struct {
  Services map[string]*ServiceConfig `json:"services,omitempty" xml:"services,omitempty"`
  Variables map[string]*Variable `json:"variables,omitempty" xml:"variables,omitempty"`
}

func (s EnvironmentStagedConfigs) String() string {
  return dara.Prettify(s)
}

func (s EnvironmentStagedConfigs) GoString() string {
  return s.String()
}

func (s *EnvironmentStagedConfigs) GetServices() map[string]*ServiceConfig  {
  return s.Services
}

func (s *EnvironmentStagedConfigs) GetVariables() map[string]*Variable  {
  return s.Variables
}

func (s *EnvironmentStagedConfigs) SetServices(v map[string]*ServiceConfig) *EnvironmentStagedConfigs {
  s.Services = v
  return s
}

func (s *EnvironmentStagedConfigs) SetVariables(v map[string]*Variable) *EnvironmentStagedConfigs {
  s.Variables = v
  return s
}

func (s *EnvironmentStagedConfigs) Validate() error {
  return dara.Validate(s)
}

