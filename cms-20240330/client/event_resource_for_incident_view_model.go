// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEventResourceForIncidentView interface {
  dara.Model
  String() string
  GoString() string
  SetDomain(v string) *EventResourceForIncidentView
  GetDomain() *string 
  SetEntityId(v string) *EventResourceForIncidentView
  GetEntityId() *string 
  SetEntityType(v string) *EventResourceForIncidentView
  GetEntityType() *string 
  SetProbs(v string) *EventResourceForIncidentView
  GetProbs() *string 
  SetTags(v string) *EventResourceForIncidentView
  GetTags() *string 
}

type EventResourceForIncidentView struct {
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
  EntityId *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
  EntityType *string `json:"entityType,omitempty" xml:"entityType,omitempty"`
  Probs *string `json:"probs,omitempty" xml:"probs,omitempty"`
  Tags *string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s EventResourceForIncidentView) String() string {
  return dara.Prettify(s)
}

func (s EventResourceForIncidentView) GoString() string {
  return s.String()
}

func (s *EventResourceForIncidentView) GetDomain() *string  {
  return s.Domain
}

func (s *EventResourceForIncidentView) GetEntityId() *string  {
  return s.EntityId
}

func (s *EventResourceForIncidentView) GetEntityType() *string  {
  return s.EntityType
}

func (s *EventResourceForIncidentView) GetProbs() *string  {
  return s.Probs
}

func (s *EventResourceForIncidentView) GetTags() *string  {
  return s.Tags
}

func (s *EventResourceForIncidentView) SetDomain(v string) *EventResourceForIncidentView {
  s.Domain = &v
  return s
}

func (s *EventResourceForIncidentView) SetEntityId(v string) *EventResourceForIncidentView {
  s.EntityId = &v
  return s
}

func (s *EventResourceForIncidentView) SetEntityType(v string) *EventResourceForIncidentView {
  s.EntityType = &v
  return s
}

func (s *EventResourceForIncidentView) SetProbs(v string) *EventResourceForIncidentView {
  s.Probs = &v
  return s
}

func (s *EventResourceForIncidentView) SetTags(v string) *EventResourceForIncidentView {
  s.Tags = &v
  return s
}

func (s *EventResourceForIncidentView) Validate() error {
  return dara.Validate(s)
}

