// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iSessionFinished interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *SessionFinished
  GetType() *string 
  SetEventId(v string) *SessionFinished
  GetEventId() *string 
  SetSessionId(v string) *SessionFinished
  GetSessionId() *string 
  SetCausedBy(v string) *SessionFinished
  GetCausedBy() *string 
  SetReason(v string) *SessionFinished
  GetReason() *string 
  SetUsageComplete(v bool) *SessionFinished
  GetUsageComplete() *bool 
}

type SessionFinished struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  SessionId *string `json:"session_id,omitempty" xml:"session_id,omitempty"`
  CausedBy *string `json:"caused_by,omitempty" xml:"caused_by,omitempty"`
  Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
  // true 表示结束时用量已完成 MQ 接收与通知写出；超时、取消、异常或用量不完整时为 false，不代表最终扣费结果。兼容旧版本时可省略，网关不补默认值。
  UsageComplete *bool `json:"usage_complete,omitempty" xml:"usage_complete,omitempty"`
}

func (s SessionFinished) String() string {
  return dara.Prettify(s)
}

func (s SessionFinished) GoString() string {
  return s.String()
}

func (s *SessionFinished) GetType() *string  {
  return s.Type
}

func (s *SessionFinished) GetEventId() *string  {
  return s.EventId
}

func (s *SessionFinished) GetSessionId() *string  {
  return s.SessionId
}

func (s *SessionFinished) GetCausedBy() *string  {
  return s.CausedBy
}

func (s *SessionFinished) GetReason() *string  {
  return s.Reason
}

func (s *SessionFinished) GetUsageComplete() *bool  {
  return s.UsageComplete
}

func (s *SessionFinished) SetType(v string) *SessionFinished {
  s.Type = &v
  return s
}

func (s *SessionFinished) SetEventId(v string) *SessionFinished {
  s.EventId = &v
  return s
}

func (s *SessionFinished) SetSessionId(v string) *SessionFinished {
  s.SessionId = &v
  return s
}

func (s *SessionFinished) SetCausedBy(v string) *SessionFinished {
  s.CausedBy = &v
  return s
}

func (s *SessionFinished) SetReason(v string) *SessionFinished {
  s.Reason = &v
  return s
}

func (s *SessionFinished) SetUsageComplete(v bool) *SessionFinished {
  s.UsageComplete = &v
  return s
}

func (s *SessionFinished) Validate() error {
  return dara.Validate(s)
}

