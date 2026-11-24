// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iTranslationAudioDelta interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *TranslationAudioDelta
  GetType() *string 
  SetEventId(v string) *TranslationAudioDelta
  GetEventId() *string 
  SetSessionId(v string) *TranslationAudioDelta
  GetSessionId() *string 
  SetSegmentId(v string) *TranslationAudioDelta
  GetSegmentId() *string 
  SetDirection(v string) *TranslationAudioDelta
  GetDirection() *string 
  SetSourceLanguage(v string) *TranslationAudioDelta
  GetSourceLanguage() *string 
  SetTargetLanguage(v string) *TranslationAudioDelta
  GetTargetLanguage() *string 
  SetAudio(v string) *TranslationAudioDelta
  GetAudio() *string 
  SetSequence(v int64) *TranslationAudioDelta
  GetSequence() *int64 
}

type TranslationAudioDelta struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  SessionId *string `json:"session_id,omitempty" xml:"session_id,omitempty"`
  SegmentId *string `json:"segment_id,omitempty" xml:"segment_id,omitempty"`
  Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
  SourceLanguage *string `json:"source_language,omitempty" xml:"source_language,omitempty"`
  TargetLanguage *string `json:"target_language,omitempty" xml:"target_language,omitempty"`
  Audio *string `json:"audio,omitempty" xml:"audio,omitempty"`
  Sequence *int64 `json:"sequence,omitempty" xml:"sequence,omitempty"`
}

func (s TranslationAudioDelta) String() string {
  return dara.Prettify(s)
}

func (s TranslationAudioDelta) GoString() string {
  return s.String()
}

func (s *TranslationAudioDelta) GetType() *string  {
  return s.Type
}

func (s *TranslationAudioDelta) GetEventId() *string  {
  return s.EventId
}

func (s *TranslationAudioDelta) GetSessionId() *string  {
  return s.SessionId
}

func (s *TranslationAudioDelta) GetSegmentId() *string  {
  return s.SegmentId
}

func (s *TranslationAudioDelta) GetDirection() *string  {
  return s.Direction
}

func (s *TranslationAudioDelta) GetSourceLanguage() *string  {
  return s.SourceLanguage
}

func (s *TranslationAudioDelta) GetTargetLanguage() *string  {
  return s.TargetLanguage
}

func (s *TranslationAudioDelta) GetAudio() *string  {
  return s.Audio
}

func (s *TranslationAudioDelta) GetSequence() *int64  {
  return s.Sequence
}

func (s *TranslationAudioDelta) SetType(v string) *TranslationAudioDelta {
  s.Type = &v
  return s
}

func (s *TranslationAudioDelta) SetEventId(v string) *TranslationAudioDelta {
  s.EventId = &v
  return s
}

func (s *TranslationAudioDelta) SetSessionId(v string) *TranslationAudioDelta {
  s.SessionId = &v
  return s
}

func (s *TranslationAudioDelta) SetSegmentId(v string) *TranslationAudioDelta {
  s.SegmentId = &v
  return s
}

func (s *TranslationAudioDelta) SetDirection(v string) *TranslationAudioDelta {
  s.Direction = &v
  return s
}

func (s *TranslationAudioDelta) SetSourceLanguage(v string) *TranslationAudioDelta {
  s.SourceLanguage = &v
  return s
}

func (s *TranslationAudioDelta) SetTargetLanguage(v string) *TranslationAudioDelta {
  s.TargetLanguage = &v
  return s
}

func (s *TranslationAudioDelta) SetAudio(v string) *TranslationAudioDelta {
  s.Audio = &v
  return s
}

func (s *TranslationAudioDelta) SetSequence(v int64) *TranslationAudioDelta {
  s.Sequence = &v
  return s
}

func (s *TranslationAudioDelta) Validate() error {
  return dara.Validate(s)
}

