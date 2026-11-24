// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iInputAudioBufferCommit interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *InputAudioBufferCommit
  GetType() *string 
  SetEventId(v string) *InputAudioBufferCommit
  GetEventId() *string 
}

type InputAudioBufferCommit struct {
  // 无额外业务入参。Manual 模式提交非空缓冲；双向须等待上一提交确认，server_vad 模式拒绝显式 commit。
  // 
  // This parameter is required.
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // 可省略；提供时必须为非空、非纯空白字符串。
  // 
  // example:
  // 
  // evt_commit_1
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
}

func (s InputAudioBufferCommit) String() string {
  return dara.Prettify(s)
}

func (s InputAudioBufferCommit) GoString() string {
  return s.String()
}

func (s *InputAudioBufferCommit) GetType() *string  {
  return s.Type
}

func (s *InputAudioBufferCommit) GetEventId() *string  {
  return s.EventId
}

func (s *InputAudioBufferCommit) SetType(v string) *InputAudioBufferCommit {
  s.Type = &v
  return s
}

func (s *InputAudioBufferCommit) SetEventId(v string) *InputAudioBufferCommit {
  s.EventId = &v
  return s
}

func (s *InputAudioBufferCommit) Validate() error {
  return dara.Validate(s)
}

