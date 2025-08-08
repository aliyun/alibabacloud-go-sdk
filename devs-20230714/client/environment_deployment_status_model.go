// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnvironmentDeploymentStatus interface {
  dara.Model
  String() string
  GoString() string
  SetFinishedTime(v string) *EnvironmentDeploymentStatus
  GetFinishedTime() *string 
  SetPhase(v string) *EnvironmentDeploymentStatus
  GetPhase() *string 
  SetPipelineName(v string) *EnvironmentDeploymentStatus
  GetPipelineName() *string 
  SetServiceDeployments(v map[string]*string) *EnvironmentDeploymentStatus
  GetServiceDeployments() map[string]*string 
}

type EnvironmentDeploymentStatus struct {
  FinishedTime *string `json:"finishedTime,omitempty" xml:"finishedTime,omitempty"`
  // example:
  // 
  // Running
  Phase *string `json:"phase,omitempty" xml:"phase,omitempty"`
  // example:
  // 
  // my-pipeline
  PipelineName *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
  ServiceDeployments map[string]*string `json:"serviceDeployments,omitempty" xml:"serviceDeployments,omitempty"`
}

func (s EnvironmentDeploymentStatus) String() string {
  return dara.Prettify(s)
}

func (s EnvironmentDeploymentStatus) GoString() string {
  return s.String()
}

func (s *EnvironmentDeploymentStatus) GetFinishedTime() *string  {
  return s.FinishedTime
}

func (s *EnvironmentDeploymentStatus) GetPhase() *string  {
  return s.Phase
}

func (s *EnvironmentDeploymentStatus) GetPipelineName() *string  {
  return s.PipelineName
}

func (s *EnvironmentDeploymentStatus) GetServiceDeployments() map[string]*string  {
  return s.ServiceDeployments
}

func (s *EnvironmentDeploymentStatus) SetFinishedTime(v string) *EnvironmentDeploymentStatus {
  s.FinishedTime = &v
  return s
}

func (s *EnvironmentDeploymentStatus) SetPhase(v string) *EnvironmentDeploymentStatus {
  s.Phase = &v
  return s
}

func (s *EnvironmentDeploymentStatus) SetPipelineName(v string) *EnvironmentDeploymentStatus {
  s.PipelineName = &v
  return s
}

func (s *EnvironmentDeploymentStatus) SetServiceDeployments(v map[string]*string) *EnvironmentDeploymentStatus {
  s.ServiceDeployments = v
  return s
}

func (s *EnvironmentDeploymentStatus) Validate() error {
  return dara.Validate(s)
}

