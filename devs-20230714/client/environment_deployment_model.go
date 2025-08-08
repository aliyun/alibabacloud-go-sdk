// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnvironmentDeployment interface {
  dara.Model
  String() string
  GoString() string
  SetCreatedTime(v string) *EnvironmentDeployment
  GetCreatedTime() *string 
  SetDescription(v string) *EnvironmentDeployment
  GetDescription() *string 
  SetKind(v string) *EnvironmentDeployment
  GetKind() *string 
  SetLabels(v map[string]*string) *EnvironmentDeployment
  GetLabels() map[string]*string 
  SetName(v string) *EnvironmentDeployment
  GetName() *string 
  SetSpec(v *EnvironmentDeploymentSpec) *EnvironmentDeployment
  GetSpec() *EnvironmentDeploymentSpec 
  SetStatus(v *EnvironmentDeploymentStatus) *EnvironmentDeployment
  GetStatus() *EnvironmentDeploymentStatus 
  SetUid(v string) *EnvironmentDeployment
  GetUid() *string 
}

type EnvironmentDeployment struct {
  // example:
  // 
  // 2021-11-19T09:34:38Z
  CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
  // example:
  // 
  // commit by xxx.
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // example:
  // 
  // Deployment
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
  Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // my-deployment
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  Spec *EnvironmentDeploymentSpec `json:"spec,omitempty" xml:"spec,omitempty"`
  Status *EnvironmentDeploymentStatus `json:"status,omitempty" xml:"status,omitempty"`
  // example:
  // 
  // 1455541096***548
  Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s EnvironmentDeployment) String() string {
  return dara.Prettify(s)
}

func (s EnvironmentDeployment) GoString() string {
  return s.String()
}

func (s *EnvironmentDeployment) GetCreatedTime() *string  {
  return s.CreatedTime
}

func (s *EnvironmentDeployment) GetDescription() *string  {
  return s.Description
}

func (s *EnvironmentDeployment) GetKind() *string  {
  return s.Kind
}

func (s *EnvironmentDeployment) GetLabels() map[string]*string  {
  return s.Labels
}

func (s *EnvironmentDeployment) GetName() *string  {
  return s.Name
}

func (s *EnvironmentDeployment) GetSpec() *EnvironmentDeploymentSpec  {
  return s.Spec
}

func (s *EnvironmentDeployment) GetStatus() *EnvironmentDeploymentStatus  {
  return s.Status
}

func (s *EnvironmentDeployment) GetUid() *string  {
  return s.Uid
}

func (s *EnvironmentDeployment) SetCreatedTime(v string) *EnvironmentDeployment {
  s.CreatedTime = &v
  return s
}

func (s *EnvironmentDeployment) SetDescription(v string) *EnvironmentDeployment {
  s.Description = &v
  return s
}

func (s *EnvironmentDeployment) SetKind(v string) *EnvironmentDeployment {
  s.Kind = &v
  return s
}

func (s *EnvironmentDeployment) SetLabels(v map[string]*string) *EnvironmentDeployment {
  s.Labels = v
  return s
}

func (s *EnvironmentDeployment) SetName(v string) *EnvironmentDeployment {
  s.Name = &v
  return s
}

func (s *EnvironmentDeployment) SetSpec(v *EnvironmentDeploymentSpec) *EnvironmentDeployment {
  s.Spec = v
  return s
}

func (s *EnvironmentDeployment) SetStatus(v *EnvironmentDeploymentStatus) *EnvironmentDeployment {
  s.Status = v
  return s
}

func (s *EnvironmentDeployment) SetUid(v string) *EnvironmentDeployment {
  s.Uid = &v
  return s
}

func (s *EnvironmentDeployment) Validate() error {
  return dara.Validate(s)
}

