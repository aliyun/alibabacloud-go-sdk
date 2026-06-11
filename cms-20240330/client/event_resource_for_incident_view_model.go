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
  // The domain of the resource type.
  // 
  // example:
  // 
  // rum
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
  // The entity ID.
  // 
  // example:
  // 
  // default
  EntityId *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
  // The entity type.
  // 
  // example:
  // 
  // "Instance"
  EntityType *string `json:"entityType,omitempty" xml:"entityType,omitempty"`
  // The string that describes the properties of the resource.
  // 
  // example:
  // 
  // "{"name":"web-server-01","ip":"10.0.0.1"}"
  Probs *string `json:"probs,omitempty" xml:"probs,omitempty"`
  // The tags.
  // 
  // example:
  // 
  // [{\\"value\\":\\"测试\\",\\"key\\":\\"环境\\"}]
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

