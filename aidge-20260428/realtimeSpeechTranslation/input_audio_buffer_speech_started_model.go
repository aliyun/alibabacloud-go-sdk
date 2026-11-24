// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iInputAudioBufferSpeechStarted interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *InputAudioBufferSpeechStarted
  GetType() *string 
  SetEventId(v string) *InputAudioBufferSpeechStarted
  GetEventId() *string 
  SetAudioStartMs(v int64) *InputAudioBufferSpeechStarted
  GetAudioStartMs() *int64 
  SetItemId(v string) *InputAudioBufferSpeechStarted
  GetItemId() *string 
}

type InputAudioBufferSpeechStarted struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  AudioStartMs *int64 `json:"audio_start_ms,omitempty" xml:"audio_start_ms,omitempty"`
  ItemId *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

func (s InputAudioBufferSpeechStarted) String() string {
  return dara.Prettify(s)
}

func (s InputAudioBufferSpeechStarted) GoString() string {
  return s.String()
}

func (s *InputAudioBufferSpeechStarted) GetType() *string  {
  return s.Type
}

func (s *InputAudioBufferSpeechStarted) GetEventId() *string  {
  return s.EventId
}

func (s *InputAudioBufferSpeechStarted) GetAudioStartMs() *int64  {
  return s.AudioStartMs
}

func (s *InputAudioBufferSpeechStarted) GetItemId() *string  {
  return s.ItemId
}

func (s *InputAudioBufferSpeechStarted) SetType(v string) *InputAudioBufferSpeechStarted {
  s.Type = &v
  return s
}

func (s *InputAudioBufferSpeechStarted) SetEventId(v string) *InputAudioBufferSpeechStarted {
  s.EventId = &v
  return s
}

func (s *InputAudioBufferSpeechStarted) SetAudioStartMs(v int64) *InputAudioBufferSpeechStarted {
  s.AudioStartMs = &v
  return s
}

func (s *InputAudioBufferSpeechStarted) SetItemId(v string) *InputAudioBufferSpeechStarted {
  s.ItemId = &v
  return s
}

func (s *InputAudioBufferSpeechStarted) Validate() error {
  return dara.Validate(s)
}

