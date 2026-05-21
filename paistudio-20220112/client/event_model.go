// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEvent interface {
  dara.Model
  String() string
  GoString() string
  SetContent(v string) *Event
  GetContent() *string 
  SetEventId(v string) *Event
  GetEventId() *string 
  SetEventType(v string) *Event
  GetEventType() *string 
  SetFunction(v string) *Event
  GetFunction() *string 
  SetIsTruncated(v bool) *Event
  GetIsTruncated() *bool 
  SetObjectId(v string) *Event
  GetObjectId() *string 
  SetObjectType(v string) *Event
  GetObjectType() *string 
  SetTriggerTenantId(v string) *Event
  GetTriggerTenantId() *string 
  SetTriggerTime(v string) *Event
  GetTriggerTime() *string 
  SetTriggerUserId(v string) *Event
  GetTriggerUserId() *string 
  SetTriggeredBy(v string) *Event
  GetTriggeredBy() *string 
}

type Event struct {
  Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
  EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
  EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
  Function *string `json:"Function,omitempty" xml:"Function,omitempty"`
  IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
  ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
  ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
  TriggerTenantId *string `json:"TriggerTenantId,omitempty" xml:"TriggerTenantId,omitempty"`
  TriggerTime *string `json:"TriggerTime,omitempty" xml:"TriggerTime,omitempty"`
  TriggerUserId *string `json:"TriggerUserId,omitempty" xml:"TriggerUserId,omitempty"`
  TriggeredBy *string `json:"TriggeredBy,omitempty" xml:"TriggeredBy,omitempty"`
}

func (s Event) String() string {
  return dara.Prettify(s)
}

func (s Event) GoString() string {
  return s.String()
}

func (s *Event) GetContent() *string  {
  return s.Content
}

func (s *Event) GetEventId() *string  {
  return s.EventId
}

func (s *Event) GetEventType() *string  {
  return s.EventType
}

func (s *Event) GetFunction() *string  {
  return s.Function
}

func (s *Event) GetIsTruncated() *bool  {
  return s.IsTruncated
}

func (s *Event) GetObjectId() *string  {
  return s.ObjectId
}

func (s *Event) GetObjectType() *string  {
  return s.ObjectType
}

func (s *Event) GetTriggerTenantId() *string  {
  return s.TriggerTenantId
}

func (s *Event) GetTriggerTime() *string  {
  return s.TriggerTime
}

func (s *Event) GetTriggerUserId() *string  {
  return s.TriggerUserId
}

func (s *Event) GetTriggeredBy() *string  {
  return s.TriggeredBy
}

func (s *Event) SetContent(v string) *Event {
  s.Content = &v
  return s
}

func (s *Event) SetEventId(v string) *Event {
  s.EventId = &v
  return s
}

func (s *Event) SetEventType(v string) *Event {
  s.EventType = &v
  return s
}

func (s *Event) SetFunction(v string) *Event {
  s.Function = &v
  return s
}

func (s *Event) SetIsTruncated(v bool) *Event {
  s.IsTruncated = &v
  return s
}

func (s *Event) SetObjectId(v string) *Event {
  s.ObjectId = &v
  return s
}

func (s *Event) SetObjectType(v string) *Event {
  s.ObjectType = &v
  return s
}

func (s *Event) SetTriggerTenantId(v string) *Event {
  s.TriggerTenantId = &v
  return s
}

func (s *Event) SetTriggerTime(v string) *Event {
  s.TriggerTime = &v
  return s
}

func (s *Event) SetTriggerUserId(v string) *Event {
  s.TriggerUserId = &v
  return s
}

func (s *Event) SetTriggeredBy(v string) *Event {
  s.TriggeredBy = &v
  return s
}

func (s *Event) Validate() error {
  return dara.Validate(s)
}

