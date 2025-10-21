// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEvent interface {
  dara.Model
  String() string
  GoString() string
  SetCreatedAt(v string) *Event
  GetCreatedAt() *string 
  SetDeploymentId(v string) *Event
  GetDeploymentId() *string 
  SetEventId(v string) *Event
  GetEventId() *string 
  SetEventName(v string) *Event
  GetEventName() *string 
  SetJobId(v string) *Event
  GetJobId() *string 
  SetMessage(v string) *Event
  GetMessage() *string 
  SetNamespace(v string) *Event
  GetNamespace() *string 
  SetWorkspace(v string) *Event
  GetWorkspace() *string 
}

type Event struct {
  CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
  // example:
  // 
  // 00000000-0000-0000-0000-000000680003
  DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
  // example:
  // 
  // 00000000-0000-0000-0000-000000000001
  EventId *string `json:"eventId,omitempty" xml:"eventId,omitempty"`
  // example:
  // 
  // STATE_TRANSITION_IS_COMPLETED
  EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
  // example:
  // 
  // 00000000-0000-0000-0000-000000005043
  JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  // example:
  // 
  // default-namespace
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
  // example:
  // 
  // edcef******b4f
  Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s Event) String() string {
  return dara.Prettify(s)
}

func (s Event) GoString() string {
  return s.String()
}

func (s *Event) GetCreatedAt() *string  {
  return s.CreatedAt
}

func (s *Event) GetDeploymentId() *string  {
  return s.DeploymentId
}

func (s *Event) GetEventId() *string  {
  return s.EventId
}

func (s *Event) GetEventName() *string  {
  return s.EventName
}

func (s *Event) GetJobId() *string  {
  return s.JobId
}

func (s *Event) GetMessage() *string  {
  return s.Message
}

func (s *Event) GetNamespace() *string  {
  return s.Namespace
}

func (s *Event) GetWorkspace() *string  {
  return s.Workspace
}

func (s *Event) SetCreatedAt(v string) *Event {
  s.CreatedAt = &v
  return s
}

func (s *Event) SetDeploymentId(v string) *Event {
  s.DeploymentId = &v
  return s
}

func (s *Event) SetEventId(v string) *Event {
  s.EventId = &v
  return s
}

func (s *Event) SetEventName(v string) *Event {
  s.EventName = &v
  return s
}

func (s *Event) SetJobId(v string) *Event {
  s.JobId = &v
  return s
}

func (s *Event) SetMessage(v string) *Event {
  s.Message = &v
  return s
}

func (s *Event) SetNamespace(v string) *Event {
  s.Namespace = &v
  return s
}

func (s *Event) SetWorkspace(v string) *Event {
  s.Workspace = &v
  return s
}

func (s *Event) Validate() error {
  return dara.Validate(s)
}

