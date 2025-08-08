// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnvironmentSnapshot interface {
  dara.Model
  String() string
  GoString() string
  SetServices(v map[string]*ServiceInstance) *EnvironmentSnapshot
  GetServices() map[string]*ServiceInstance 
}

type EnvironmentSnapshot struct {
  Services map[string]*ServiceInstance `json:"services,omitempty" xml:"services,omitempty"`
}

func (s EnvironmentSnapshot) String() string {
  return dara.Prettify(s)
}

func (s EnvironmentSnapshot) GoString() string {
  return s.String()
}

func (s *EnvironmentSnapshot) GetServices() map[string]*ServiceInstance  {
  return s.Services
}

func (s *EnvironmentSnapshot) SetServices(v map[string]*ServiceInstance) *EnvironmentSnapshot {
  s.Services = v
  return s
}

func (s *EnvironmentSnapshot) Validate() error {
  return dara.Validate(s)
}

