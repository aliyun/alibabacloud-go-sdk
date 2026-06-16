// This file is auto-generated, don't edit it. Thanks.
package multiModalGuardWs

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iAckMessageEvent interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int32) *AckMessageEvent
  GetCode() *int32 
  SetMessage(v string) *AckMessageEvent
  GetMessage() *string 
  SetData(v *AckMessageEventData) *AckMessageEvent
  GetData() *AckMessageEventData 
}

type AckMessageEvent struct {
  Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  Data *AckMessageEventData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s AckMessageEvent) String() string {
  return dara.Prettify(s)
}

func (s AckMessageEvent) GoString() string {
  return s.String()
}

func (s *AckMessageEvent) GetCode() *int32  {
  return s.Code
}

func (s *AckMessageEvent) GetMessage() *string  {
  return s.Message
}

func (s *AckMessageEvent) GetData() *AckMessageEventData  {
  return s.Data
}

func (s *AckMessageEvent) SetCode(v int32) *AckMessageEvent {
  s.Code = &v
  return s
}

func (s *AckMessageEvent) SetMessage(v string) *AckMessageEvent {
  s.Message = &v
  return s
}

func (s *AckMessageEvent) SetData(v *AckMessageEventData) *AckMessageEvent {
  s.Data = v
  return s
}

func (s *AckMessageEvent) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type AckMessageEventData struct {
  Triggered *bool `json:"Triggered,omitempty" xml:"Triggered,omitempty"`
  MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
}

func (s AckMessageEventData) String() string {
  return dara.Prettify(s)
}

func (s AckMessageEventData) GoString() string {
  return s.String()
}

func (s *AckMessageEventData) GetTriggered() *bool  {
  return s.Triggered
}

func (s *AckMessageEventData) GetMsgId() *string  {
  return s.MsgId
}

func (s *AckMessageEventData) SetTriggered(v bool) *AckMessageEventData {
  s.Triggered = &v
  return s
}

func (s *AckMessageEventData) SetMsgId(v string) *AckMessageEventData {
  s.MsgId = &v
  return s
}

func (s *AckMessageEventData) Validate() error {
  return dara.Validate(s)
}

