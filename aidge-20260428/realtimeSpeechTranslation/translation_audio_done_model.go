// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iTranslationAudioDone interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *TranslationAudioDone
  GetType() *string 
  SetEventId(v string) *TranslationAudioDone
  GetEventId() *string 
  SetSessionId(v string) *TranslationAudioDone
  GetSessionId() *string 
  SetSegmentId(v string) *TranslationAudioDone
  GetSegmentId() *string 
  SetDirection(v string) *TranslationAudioDone
  GetDirection() *string 
  SetSourceLanguage(v string) *TranslationAudioDone
  GetSourceLanguage() *string 
  SetTargetLanguage(v string) *TranslationAudioDone
  GetTargetLanguage() *string 
  SetFinishReason(v string) *TranslationAudioDone
  GetFinishReason() *string 
}

type TranslationAudioDone struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  SessionId *string `json:"session_id,omitempty" xml:"session_id,omitempty"`
  SegmentId *string `json:"segment_id,omitempty" xml:"segment_id,omitempty"`
  Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
  SourceLanguage *string `json:"source_language,omitempty" xml:"source_language,omitempty"`
  TargetLanguage *string `json:"target_language,omitempty" xml:"target_language,omitempty"`
  FinishReason *string `json:"finish_reason,omitempty" xml:"finish_reason,omitempty"`
}

func (s TranslationAudioDone) String() string {
  return dara.Prettify(s)
}

func (s TranslationAudioDone) GoString() string {
  return s.String()
}

func (s *TranslationAudioDone) GetType() *string  {
  return s.Type
}

func (s *TranslationAudioDone) GetEventId() *string  {
  return s.EventId
}

func (s *TranslationAudioDone) GetSessionId() *string  {
  return s.SessionId
}

func (s *TranslationAudioDone) GetSegmentId() *string  {
  return s.SegmentId
}

func (s *TranslationAudioDone) GetDirection() *string  {
  return s.Direction
}

func (s *TranslationAudioDone) GetSourceLanguage() *string  {
  return s.SourceLanguage
}

func (s *TranslationAudioDone) GetTargetLanguage() *string  {
  return s.TargetLanguage
}

func (s *TranslationAudioDone) GetFinishReason() *string  {
  return s.FinishReason
}

func (s *TranslationAudioDone) SetType(v string) *TranslationAudioDone {
  s.Type = &v
  return s
}

func (s *TranslationAudioDone) SetEventId(v string) *TranslationAudioDone {
  s.EventId = &v
  return s
}

func (s *TranslationAudioDone) SetSessionId(v string) *TranslationAudioDone {
  s.SessionId = &v
  return s
}

func (s *TranslationAudioDone) SetSegmentId(v string) *TranslationAudioDone {
  s.SegmentId = &v
  return s
}

func (s *TranslationAudioDone) SetDirection(v string) *TranslationAudioDone {
  s.Direction = &v
  return s
}

func (s *TranslationAudioDone) SetSourceLanguage(v string) *TranslationAudioDone {
  s.SourceLanguage = &v
  return s
}

func (s *TranslationAudioDone) SetTargetLanguage(v string) *TranslationAudioDone {
  s.TargetLanguage = &v
  return s
}

func (s *TranslationAudioDone) SetFinishReason(v string) *TranslationAudioDone {
  s.FinishReason = &v
  return s
}

func (s *TranslationAudioDone) Validate() error {
  return dara.Validate(s)
}

