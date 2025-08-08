// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnvironmentBaseline interface {
  dara.Model
  String() string
  GoString() string
  SetServicesInstances(v map[string]*ServiceInstance) *EnvironmentBaseline
  GetServicesInstances() map[string]*ServiceInstance 
  SetVariables(v map[string]*Variable) *EnvironmentBaseline
  GetVariables() map[string]*Variable 
}

type EnvironmentBaseline struct {
  ServicesInstances map[string]*ServiceInstance `json:"servicesInstances,omitempty" xml:"servicesInstances,omitempty"`
  Variables map[string]*Variable `json:"variables,omitempty" xml:"variables,omitempty"`
}

func (s EnvironmentBaseline) String() string {
  return dara.Prettify(s)
}

func (s EnvironmentBaseline) GoString() string {
  return s.String()
}

func (s *EnvironmentBaseline) GetServicesInstances() map[string]*ServiceInstance  {
  return s.ServicesInstances
}

func (s *EnvironmentBaseline) GetVariables() map[string]*Variable  {
  return s.Variables
}

func (s *EnvironmentBaseline) SetServicesInstances(v map[string]*ServiceInstance) *EnvironmentBaseline {
  s.ServicesInstances = v
  return s
}

func (s *EnvironmentBaseline) SetVariables(v map[string]*Variable) *EnvironmentBaseline {
  s.Variables = v
  return s
}

func (s *EnvironmentBaseline) Validate() error {
  return dara.Validate(s)
}

