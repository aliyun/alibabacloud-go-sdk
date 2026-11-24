// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iInputAudioBufferClear interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *InputAudioBufferClear
  GetType() *string 
  SetEventId(v string) *InputAudioBufferClear
  GetEventId() *string 
}

type InputAudioBufferClear struct {
  // 无额外业务入参。清理未处理缓冲，不撤回已下发结果；双向仅 manual 模式允许，server_vad 或存在待确认 commit 时拒绝。
  // 
  // This parameter is required.
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // 可省略；提供时必须为非空、非纯空白字符串。
  // 
  // example:
  // 
  // evt_clear_1
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
}

func (s InputAudioBufferClear) String() string {
  return dara.Prettify(s)
}

func (s InputAudioBufferClear) GoString() string {
  return s.String()
}

func (s *InputAudioBufferClear) GetType() *string  {
  return s.Type
}

func (s *InputAudioBufferClear) GetEventId() *string  {
  return s.EventId
}

func (s *InputAudioBufferClear) SetType(v string) *InputAudioBufferClear {
  s.Type = &v
  return s
}

func (s *InputAudioBufferClear) SetEventId(v string) *InputAudioBufferClear {
  s.EventId = &v
  return s
}

func (s *InputAudioBufferClear) Validate() error {
  return dara.Validate(s)
}

