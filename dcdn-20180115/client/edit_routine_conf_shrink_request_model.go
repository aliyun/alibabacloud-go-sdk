// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditRoutineConfShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDescription(v string) *EditRoutineConfShrinkRequest
  GetDescription() *string 
  SetEnvConfShrink(v string) *EditRoutineConfShrinkRequest
  GetEnvConfShrink() *string 
  SetName(v string) *EditRoutineConfShrinkRequest
  GetName() *string 
}

type EditRoutineConfShrinkRequest struct {
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
  EnvConfShrink *string `json:"EnvConf,omitempty" xml:"EnvConf,omitempty"`
  // The name of the routine. The name must be unique among the routines that belong to the same Alibaba Cloud account.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // test-slc
  Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s EditRoutineConfShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s EditRoutineConfShrinkRequest) GoString() string {
  return s.String()
}

func (s *EditRoutineConfShrinkRequest) GetDescription() *string  {
  return s.Description
}

func (s *EditRoutineConfShrinkRequest) GetEnvConfShrink() *string  {
  return s.EnvConfShrink
}

func (s *EditRoutineConfShrinkRequest) GetName() *string  {
  return s.Name
}

func (s *EditRoutineConfShrinkRequest) SetDescription(v string) *EditRoutineConfShrinkRequest {
  s.Description = &v
  return s
}

func (s *EditRoutineConfShrinkRequest) SetEnvConfShrink(v string) *EditRoutineConfShrinkRequest {
  s.EnvConfShrink = &v
  return s
}

func (s *EditRoutineConfShrinkRequest) SetName(v string) *EditRoutineConfShrinkRequest {
  s.Name = &v
  return s
}

func (s *EditRoutineConfShrinkRequest) Validate() error {
  return dara.Validate(s)
}

