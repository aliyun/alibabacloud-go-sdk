// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iSessionFinish interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *SessionFinish
  GetType() *string 
  SetEventId(v string) *SessionFinish
  GetEventId() *string 
  SetReason(v string) *SessionFinish
  GetReason() *string 
}

type SessionFinish struct {
  // This parameter is required.
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // 可省略；提供时必须为非空、非纯空白字符串。
  // 
  // example:
  // 
  // evt_finish_1
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  // 双向省略时后端按 completed 排空结果，cancelled 立即取消；clear 进行中拒绝结束。单向客户端省略本字段，等待 SessionFinished 后关闭连接。网关不补默认值。
  // 
  // example:
  // 
  // completed
  Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s SessionFinish) String() string {
  return dara.Prettify(s)
}

func (s SessionFinish) GoString() string {
  return s.String()
}

func (s *SessionFinish) GetType() *string  {
  return s.Type
}

func (s *SessionFinish) GetEventId() *string  {
  return s.EventId
}

func (s *SessionFinish) GetReason() *string  {
  return s.Reason
}

func (s *SessionFinish) SetType(v string) *SessionFinish {
  s.Type = &v
  return s
}

func (s *SessionFinish) SetEventId(v string) *SessionFinish {
  s.EventId = &v
  return s
}

func (s *SessionFinish) SetReason(v string) *SessionFinish {
  s.Reason = &v
  return s
}

func (s *SessionFinish) Validate() error {
  return dara.Validate(s)
}

