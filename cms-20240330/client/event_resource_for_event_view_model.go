// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEventResourceForEventView interface {
  dara.Model
  String() string
  GoString() string
  SetEntity(v *EventResourceForEventViewEntity) *EventResourceForEventView
  GetEntity() *EventResourceForEventViewEntity 
  SetTags(v map[string]interface{}) *EventResourceForEventView
  GetTags() map[string]interface{} 
}

type EventResourceForEventView struct {
  Entity *EventResourceForEventViewEntity `json:"entity,omitempty" xml:"entity,omitempty" type:"Struct"`
  Tags map[string]interface{} `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s EventResourceForEventView) String() string {
  return dara.Prettify(s)
}

func (s EventResourceForEventView) GoString() string {
  return s.String()
}

func (s *EventResourceForEventView) GetEntity() *EventResourceForEventViewEntity  {
  return s.Entity
}

func (s *EventResourceForEventView) GetTags() map[string]interface{}  {
  return s.Tags
}

func (s *EventResourceForEventView) SetEntity(v *EventResourceForEventViewEntity) *EventResourceForEventView {
  s.Entity = v
  return s
}

func (s *EventResourceForEventView) SetTags(v map[string]interface{}) *EventResourceForEventView {
  s.Tags = v
  return s
}

func (s *EventResourceForEventView) Validate() error {
  if s.Entity != nil {
    if err := s.Entity.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EventResourceForEventViewEntity struct {
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
  EntityId *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
  EntityType *string `json:"entityType,omitempty" xml:"entityType,omitempty"`
  Prop map[string]interface{} `json:"prop,omitempty" xml:"prop,omitempty"`
}

func (s EventResourceForEventViewEntity) String() string {
  return dara.Prettify(s)
}

func (s EventResourceForEventViewEntity) GoString() string {
  return s.String()
}

func (s *EventResourceForEventViewEntity) GetDomain() *string  {
  return s.Domain
}

func (s *EventResourceForEventViewEntity) GetEntityId() *string  {
  return s.EntityId
}

func (s *EventResourceForEventViewEntity) GetEntityType() *string  {
  return s.EntityType
}

func (s *EventResourceForEventViewEntity) GetProp() map[string]interface{}  {
  return s.Prop
}

func (s *EventResourceForEventViewEntity) SetDomain(v string) *EventResourceForEventViewEntity {
  s.Domain = &v
  return s
}

func (s *EventResourceForEventViewEntity) SetEntityId(v string) *EventResourceForEventViewEntity {
  s.EntityId = &v
  return s
}

func (s *EventResourceForEventViewEntity) SetEntityType(v string) *EventResourceForEventViewEntity {
  s.EntityType = &v
  return s
}

func (s *EventResourceForEventViewEntity) SetProp(v map[string]interface{}) *EventResourceForEventViewEntity {
  s.Prop = v
  return s
}

func (s *EventResourceForEventViewEntity) Validate() error {
  return dara.Validate(s)
}

