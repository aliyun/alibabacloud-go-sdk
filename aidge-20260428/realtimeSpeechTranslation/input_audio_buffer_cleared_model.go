// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iInputAudioBufferCleared interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *InputAudioBufferCleared
  GetType() *string 
  SetEventId(v string) *InputAudioBufferCleared
  GetEventId() *string 
  SetSessionId(v string) *InputAudioBufferCleared
  GetSessionId() *string 
  SetCausedBy(v string) *InputAudioBufferCleared
  GetCausedBy() *string 
}

type InputAudioBufferCleared struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  SessionId *string `json:"session_id,omitempty" xml:"session_id,omitempty"`
  CausedBy *string `json:"caused_by,omitempty" xml:"caused_by,omitempty"`
}

func (s InputAudioBufferCleared) String() string {
  return dara.Prettify(s)
}

func (s InputAudioBufferCleared) GoString() string {
  return s.String()
}

func (s *InputAudioBufferCleared) GetType() *string  {
  return s.Type
}

func (s *InputAudioBufferCleared) GetEventId() *string  {
  return s.EventId
}

func (s *InputAudioBufferCleared) GetSessionId() *string  {
  return s.SessionId
}

func (s *InputAudioBufferCleared) GetCausedBy() *string  {
  return s.CausedBy
}

func (s *InputAudioBufferCleared) SetType(v string) *InputAudioBufferCleared {
  s.Type = &v
  return s
}

func (s *InputAudioBufferCleared) SetEventId(v string) *InputAudioBufferCleared {
  s.EventId = &v
  return s
}

func (s *InputAudioBufferCleared) SetSessionId(v string) *InputAudioBufferCleared {
  s.SessionId = &v
  return s
}

func (s *InputAudioBufferCleared) SetCausedBy(v string) *InputAudioBufferCleared {
  s.CausedBy = &v
  return s
}

func (s *InputAudioBufferCleared) Validate() error {
  return dara.Validate(s)
}

