// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnvironmentChanges interface {
  dara.Model
  String() string
  GoString() string
  SetServices(v map[string]interface{}) *EnvironmentChanges
  GetServices() map[string]interface{} 
}

type EnvironmentChanges struct {
  Services map[string]interface{} `json:"services,omitempty" xml:"services,omitempty"`
}

func (s EnvironmentChanges) String() string {
  return dara.Prettify(s)
}

func (s EnvironmentChanges) GoString() string {
  return s.String()
}

func (s *EnvironmentChanges) GetServices() map[string]interface{}  {
  return s.Services
}

func (s *EnvironmentChanges) SetServices(v map[string]interface{}) *EnvironmentChanges {
  s.Services = v
  return s
}

func (s *EnvironmentChanges) Validate() error {
  return dara.Validate(s)
}

