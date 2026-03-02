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
  // The basic resource configuration of the JobManager.
  JobmanagerResourceSettingSpec *BasicResourceSettingSpec `json:"jobmanagerResourceSettingSpec,omitempty" xml:"jobmanagerResourceSettingSpec,omitempty"`
  // The resource configuration plan of the deployment in expert mode.
  // 
  // example:
  // 
  // {\\"ssgProfiles\\":[{\\"name\\":\\"default\\",\\"cpu\\":1.13,\\"heap\\":\\"1 gb\\",\\"offHeap\\":\\"32 mb\\",\\"managed\\":{},\\"extended\\":{}}],\\"nodes\\":[{\\"id\\":1,\\"type\\":\\"StreamExecTableSourceScan\\",\\"desc\\":\\"Source: datagen_source[78]\\",\\"profile\\":{\\"group\\":\\"default\\",\\"parallelism\\":1,\\"maxParallelism\\":32768,\\"minParallelism\\":1}},{\\"id\\":2,\\"type\\":\\"StreamExecSink\\",\\"desc\\":\\"Sink: blackhole_sink[79]\\",\\"profile\\":{\\"group\\":\\"default\\",\\"parallelism\\":1,\\"maxParallelism\\":32768,\\"minParallelism\\":1}}],\\"edges\\":[{\\"source\\":1,\\"target\\":2,\\"mode\\":\\"PIPELINED\\",\\"strategy\\":\\"FORWARD\\"}],\\"vertices\\":{\\"717c7b8afebbfb7137f6f0f99beb2a94\\":[1,2]}}
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

