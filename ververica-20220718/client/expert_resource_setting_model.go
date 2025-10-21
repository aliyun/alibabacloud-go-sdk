// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExpertResourceSetting interface {
  dara.Model
  String() string
  GoString() string
  SetJobmanagerResourceSettingSpec(v *BasicResourceSettingSpec) *ExpertResourceSetting
  GetJobmanagerResourceSettingSpec() *BasicResourceSettingSpec 
  SetResourcePlan(v string) *ExpertResourceSetting
  GetResourcePlan() *string 
}

type ExpertResourceSetting struct {
  JobmanagerResourceSettingSpec *BasicResourceSettingSpec `json:"jobmanagerResourceSettingSpec,omitempty" xml:"jobmanagerResourceSettingSpec,omitempty"`
  ResourcePlan *string `json:"resourcePlan,omitempty" xml:"resourcePlan,omitempty"`
}

func (s ExpertResourceSetting) String() string {
  return dara.Prettify(s)
}

func (s ExpertResourceSetting) GoString() string {
  return s.String()
}

func (s *ExpertResourceSetting) GetJobmanagerResourceSettingSpec() *BasicResourceSettingSpec  {
  return s.JobmanagerResourceSettingSpec
}

func (s *ExpertResourceSetting) GetResourcePlan() *string  {
  return s.ResourcePlan
}

func (s *ExpertResourceSetting) SetJobmanagerResourceSettingSpec(v *BasicResourceSettingSpec) *ExpertResourceSetting {
  s.JobmanagerResourceSettingSpec = v
  return s
}

func (s *ExpertResourceSetting) SetResourcePlan(v string) *ExpertResourceSetting {
  s.ResourcePlan = &v
  return s
}

func (s *ExpertResourceSetting) Validate() error {
  if s.JobmanagerResourceSettingSpec != nil {
    if err := s.JobmanagerResourceSettingSpec.Validate(); err != nil {
      return err
    }
  }
  return nil
}

