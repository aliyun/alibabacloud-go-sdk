// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEventInfo interface {
  dara.Model
  String() string
  GoString() string
  SetContent(v string) *EventInfo
  GetContent() *string 
  SetId(v string) *EventInfo
  GetId() *string 
  SetPodId(v string) *EventInfo
  GetPodId() *string 
  SetPodUid(v string) *EventInfo
  GetPodUid() *string 
  SetTime(v string) *EventInfo
  GetTime() *string 
}

type EventInfo struct {
  // The text content of the event.
  Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
  // The ID of the event, which is used to obtain the related context.
  Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
  // The ID of the pod to which the event belongs.
  // 
  // example:
  // 
  // dlc*****-master-0
  PodId *string `json:"PodId,omitempty" xml:"PodId,omitempty"`
  // The UID of the pod to which the event belongs.
  // 
  // example:
  // 
  // 94a7cc7c-0033-48b5-85bd-71c63592c268
  PodUid *string `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
  // The timestamp of the event in UTC.
  Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s EventInfo) String() string {
  return dara.Prettify(s)
}

func (s EventInfo) GoString() string {
  return s.String()
}

func (s *EventInfo) GetContent() *string  {
  return s.Content
}

func (s *EventInfo) GetId() *string  {
  return s.Id
}

func (s *EventInfo) GetPodId() *string  {
  return s.PodId
}

func (s *EventInfo) GetPodUid() *string  {
  return s.PodUid
}

func (s *EventInfo) GetTime() *string  {
  return s.Time
}

func (s *EventInfo) SetContent(v string) *EventInfo {
  s.Content = &v
  return s
}

func (s *EventInfo) SetId(v string) *EventInfo {
  s.Id = &v
  return s
}

func (s *EventInfo) SetPodId(v string) *EventInfo {
  s.PodId = &v
  return s
}

func (s *EventInfo) SetPodUid(v string) *EventInfo {
  s.PodUid = &v
  return s
}

func (s *EventInfo) SetTime(v string) *EventInfo {
  s.Time = &v
  return s
}

func (s *EventInfo) Validate() error {
  return dara.Validate(s)
}

