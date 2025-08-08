// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnvironment interface {
  dara.Model
  String() string
  GoString() string
  SetCreatedTime(v string) *Environment
  GetCreatedTime() *string 
  SetDescription(v string) *Environment
  GetDescription() *string 
  SetGeneration(v int32) *Environment
  GetGeneration() *int32 
  SetKind(v string) *Environment
  GetKind() *string 
  SetLabels(v map[string]*string) *Environment
  GetLabels() map[string]*string 
  SetName(v string) *Environment
  GetName() *string 
  SetProjectName(v string) *Environment
  GetProjectName() *string 
  SetSpec(v *EnvironmentSpec) *Environment
  GetSpec() *EnvironmentSpec 
  SetStatus(v *EnvironmentStatus) *Environment
  GetStatus() *EnvironmentStatus 
  SetUid(v string) *Environment
  GetUid() *string 
}

type Environment struct {
  // example:
  // 
  // 2021-11-19T09:34:38Z
  CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
  // example:
  // 
  // test env
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // example:
  // 
  // 1
  Generation *int32 `json:"generation,omitempty" xml:"generation,omitempty"`
  // example:
  // 
  // Environment
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
  Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // demo-env
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // example:
  // 
  // demo-project
  ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
  // This parameter is required.
  Spec *EnvironmentSpec `json:"spec,omitempty" xml:"spec,omitempty"`
  Status *EnvironmentStatus `json:"status,omitempty" xml:"status,omitempty"`
  // example:
  // 
  // 1455541096***548
  Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s Environment) String() string {
  return dara.Prettify(s)
}

func (s Environment) GoString() string {
  return s.String()
}

func (s *Environment) GetCreatedTime() *string  {
  return s.CreatedTime
}

func (s *Environment) GetDescription() *string  {
  return s.Description
}

func (s *Environment) GetGeneration() *int32  {
  return s.Generation
}

func (s *Environment) GetKind() *string  {
  return s.Kind
}

func (s *Environment) GetLabels() map[string]*string  {
  return s.Labels
}

func (s *Environment) GetName() *string  {
  return s.Name
}

func (s *Environment) GetProjectName() *string  {
  return s.ProjectName
}

func (s *Environment) GetSpec() *EnvironmentSpec  {
  return s.Spec
}

func (s *Environment) GetStatus() *EnvironmentStatus  {
  return s.Status
}

func (s *Environment) GetUid() *string  {
  return s.Uid
}

func (s *Environment) SetCreatedTime(v string) *Environment {
  s.CreatedTime = &v
  return s
}

func (s *Environment) SetDescription(v string) *Environment {
  s.Description = &v
  return s
}

func (s *Environment) SetGeneration(v int32) *Environment {
  s.Generation = &v
  return s
}

func (s *Environment) SetKind(v string) *Environment {
  s.Kind = &v
  return s
}

func (s *Environment) SetLabels(v map[string]*string) *Environment {
  s.Labels = v
  return s
}

func (s *Environment) SetName(v string) *Environment {
  s.Name = &v
  return s
}

func (s *Environment) SetProjectName(v string) *Environment {
  s.ProjectName = &v
  return s
}

func (s *Environment) SetSpec(v *EnvironmentSpec) *Environment {
  s.Spec = v
  return s
}

func (s *Environment) SetStatus(v *EnvironmentStatus) *Environment {
  s.Status = v
  return s
}

func (s *Environment) SetUid(v string) *Environment {
  s.Uid = &v
  return s
}

func (s *Environment) Validate() error {
  return dara.Validate(s)
}

