// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditRoutineConfRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDescription(v string) *EditRoutineConfRequest
  GetDescription() *string 
  SetEnvConf(v map[string]interface{}) *EditRoutineConfRequest
  GetEnvConf() map[string]interface{} 
  SetName(v string) *EditRoutineConfRequest
  GetName() *string 
}

type EditRoutineConfRequest struct {
  // The description of the routine.
  // 
  // example:
  // 
  // the description of this routine
  Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
  // The configurations of the specified environment.
  // 
  // example:
  // 
  // {"staging":{"SpecName":"50ms","AllowedHosts":["test-a.alicdn.com","test-b.alicdn.com"]},"production":{"SpecName":"50ms","AllowedHosts":["test-c.alicdn.com","test-d.alicdn.com"]},"presetCanaryZhejiang":{"SpecName":"100ms","AllowedHosts":["test-e.alicdn.com","test-f.alicdn.com"]},"presetCanaryBeijing":{"SpecName":"5ms","AllowedHosts":["test-g.alicdn.com","test-h.alicdn.com"]},"presetCanaryNotExist":{"SpecName":"5ms","CodeRev":"1622446907645949975","AllowedHosts":["error hosts"]}}
  EnvConf map[string]interface{} `json:"EnvConf,omitempty" xml:"EnvConf,omitempty"`
  // The name of the routine. The name must be unique among the routines that belong to the same Alibaba Cloud account.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // test-slc
  Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s EditRoutineConfRequest) String() string {
  return dara.Prettify(s)
}

func (s EditRoutineConfRequest) GoString() string {
  return s.String()
}

func (s *EditRoutineConfRequest) GetDescription() *string  {
  return s.Description
}

func (s *EditRoutineConfRequest) GetEnvConf() map[string]interface{}  {
  return s.EnvConf
}

func (s *EditRoutineConfRequest) GetName() *string  {
  return s.Name
}

func (s *EditRoutineConfRequest) SetDescription(v string) *EditRoutineConfRequest {
  s.Description = &v
  return s
}

func (s *EditRoutineConfRequest) SetEnvConf(v map[string]interface{}) *EditRoutineConfRequest {
  s.EnvConf = v
  return s
}

func (s *EditRoutineConfRequest) SetName(v string) *EditRoutineConfRequest {
  s.Name = &v
  return s
}

func (s *EditRoutineConfRequest) Validate() error {
  return dara.Validate(s)
}

