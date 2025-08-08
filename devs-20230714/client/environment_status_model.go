// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnvironmentStatus interface {
  dara.Model
  String() string
  GoString() string
  SetLatestEnvironmentDeploymentName(v string) *EnvironmentStatus
  GetLatestEnvironmentDeploymentName() *string 
  SetObservedGeneration(v int64) *EnvironmentStatus
  GetObservedGeneration() *int64 
  SetObservedTime(v string) *EnvironmentStatus
  GetObservedTime() *string 
  SetServicesInstances(v map[string]*ServiceInstance) *EnvironmentStatus
  GetServicesInstances() map[string]*ServiceInstance 
  SetServicesWithPendingChanges(v []*string) *EnvironmentStatus
  GetServicesWithPendingChanges() []*string 
}

type EnvironmentStatus struct {
  LatestEnvironmentDeploymentName *string `json:"latestEnvironmentDeploymentName,omitempty" xml:"latestEnvironmentDeploymentName,omitempty"`
  // example:
  // 
  // 1
  ObservedGeneration *int64 `json:"observedGeneration,omitempty" xml:"observedGeneration,omitempty"`
  // example:
  // 
  // 2021-11-19T09:34:38Z
  ObservedTime *string `json:"observedTime,omitempty" xml:"observedTime,omitempty"`
  ServicesInstances map[string]*ServiceInstance `json:"servicesInstances,omitempty" xml:"servicesInstances,omitempty"`
  ServicesWithPendingChanges []*string `json:"servicesWithPendingChanges,omitempty" xml:"servicesWithPendingChanges,omitempty" type:"Repeated"`
}

func (s EnvironmentStatus) String() string {
  return dara.Prettify(s)
}

func (s EnvironmentStatus) GoString() string {
  return s.String()
}

func (s *EnvironmentStatus) GetLatestEnvironmentDeploymentName() *string  {
  return s.LatestEnvironmentDeploymentName
}

func (s *EnvironmentStatus) GetObservedGeneration() *int64  {
  return s.ObservedGeneration
}

func (s *EnvironmentStatus) GetObservedTime() *string  {
  return s.ObservedTime
}

func (s *EnvironmentStatus) GetServicesInstances() map[string]*ServiceInstance  {
  return s.ServicesInstances
}

func (s *EnvironmentStatus) GetServicesWithPendingChanges() []*string  {
  return s.ServicesWithPendingChanges
}

func (s *EnvironmentStatus) SetLatestEnvironmentDeploymentName(v string) *EnvironmentStatus {
  s.LatestEnvironmentDeploymentName = &v
  return s
}

func (s *EnvironmentStatus) SetObservedGeneration(v int64) *EnvironmentStatus {
  s.ObservedGeneration = &v
  return s
}

func (s *EnvironmentStatus) SetObservedTime(v string) *EnvironmentStatus {
  s.ObservedTime = &v
  return s
}

func (s *EnvironmentStatus) SetServicesInstances(v map[string]*ServiceInstance) *EnvironmentStatus {
  s.ServicesInstances = v
  return s
}

func (s *EnvironmentStatus) SetServicesWithPendingChanges(v []*string) *EnvironmentStatus {
  s.ServicesWithPendingChanges = v
  return s
}

func (s *EnvironmentStatus) Validate() error {
  return dara.Validate(s)
}

