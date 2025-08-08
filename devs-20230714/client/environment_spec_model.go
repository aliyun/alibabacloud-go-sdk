// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnvironmentSpec interface {
  dara.Model
  String() string
  GoString() string
  SetRoleArn(v string) *EnvironmentSpec
  GetRoleArn() *string 
  SetStagedConfigs(v *EnvironmentStagedConfigs) *EnvironmentSpec
  GetStagedConfigs() *EnvironmentStagedConfigs 
  SetType(v string) *EnvironmentSpec
  GetType() *string 
}

type EnvironmentSpec struct {
  // example:
  // 
  // acs:ram::*******:role/aliyundevsdefaultrole
  RoleArn *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
  StagedConfigs *EnvironmentStagedConfigs `json:"stagedConfigs,omitempty" xml:"stagedConfigs,omitempty"`
  // example:
  // 
  // Testing
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s EnvironmentSpec) String() string {
  return dara.Prettify(s)
}

func (s EnvironmentSpec) GoString() string {
  return s.String()
}

func (s *EnvironmentSpec) GetRoleArn() *string  {
  return s.RoleArn
}

func (s *EnvironmentSpec) GetStagedConfigs() *EnvironmentStagedConfigs  {
  return s.StagedConfigs
}

func (s *EnvironmentSpec) GetType() *string  {
  return s.Type
}

func (s *EnvironmentSpec) SetRoleArn(v string) *EnvironmentSpec {
  s.RoleArn = &v
  return s
}

func (s *EnvironmentSpec) SetStagedConfigs(v *EnvironmentStagedConfigs) *EnvironmentSpec {
  s.StagedConfigs = v
  return s
}

func (s *EnvironmentSpec) SetType(v string) *EnvironmentSpec {
  s.Type = &v
  return s
}

func (s *EnvironmentSpec) Validate() error {
  return dara.Validate(s)
}

