// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iInputAudioBufferSpeechStopped interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *InputAudioBufferSpeechStopped
  GetType() *string 
  SetEventId(v string) *InputAudioBufferSpeechStopped
  GetEventId() *string 
  SetAudioEndMs(v int64) *InputAudioBufferSpeechStopped
  GetAudioEndMs() *int64 
  SetItemId(v string) *InputAudioBufferSpeechStopped
  GetItemId() *string 
}

type InputAudioBufferSpeechStopped struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  AudioEndMs *int64 `json:"audio_end_ms,omitempty" xml:"audio_end_ms,omitempty"`
  ItemId *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

func (s InputAudioBufferSpeechStopped) String() string {
  return dara.Prettify(s)
}

func (s InputAudioBufferSpeechStopped) GoString() string {
  return s.String()
}

func (s *InputAudioBufferSpeechStopped) GetType() *string  {
  return s.Type
}

func (s *InputAudioBufferSpeechStopped) GetEventId() *string  {
  return s.EventId
}

func (s *InputAudioBufferSpeechStopped) GetAudioEndMs() *int64  {
  return s.AudioEndMs
}

func (s *InputAudioBufferSpeechStopped) GetItemId() *string  {
  return s.ItemId
}

func (s *InputAudioBufferSpeechStopped) SetType(v string) *InputAudioBufferSpeechStopped {
  s.Type = &v
  return s
}

func (s *InputAudioBufferSpeechStopped) SetEventId(v string) *InputAudioBufferSpeechStopped {
  s.EventId = &v
  return s
}

func (s *InputAudioBufferSpeechStopped) SetAudioEndMs(v int64) *InputAudioBufferSpeechStopped {
  s.AudioEndMs = &v
  return s
}

func (s *InputAudioBufferSpeechStopped) SetItemId(v string) *InputAudioBufferSpeechStopped {
  s.ItemId = &v
  return s
}

func (s *InputAudioBufferSpeechStopped) Validate() error {
  return dara.Validate(s)
}

