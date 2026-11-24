// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iLanguageDetected interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *LanguageDetected
  GetType() *string 
  SetEventId(v string) *LanguageDetected
  GetEventId() *string 
  SetSessionId(v string) *LanguageDetected
  GetSessionId() *string 
  SetSegmentId(v string) *LanguageDetected
  GetSegmentId() *string 
  SetDirection(v string) *LanguageDetected
  GetDirection() *string 
  SetSourceLanguage(v string) *LanguageDetected
  GetSourceLanguage() *string 
  SetTargetLanguage(v string) *LanguageDetected
  GetTargetLanguage() *string 
}

type LanguageDetected struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  SessionId *string `json:"session_id,omitempty" xml:"session_id,omitempty"`
  SegmentId *string `json:"segment_id,omitempty" xml:"segment_id,omitempty"`
  Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
  SourceLanguage *string `json:"source_language,omitempty" xml:"source_language,omitempty"`
  TargetLanguage *string `json:"target_language,omitempty" xml:"target_language,omitempty"`
}

func (s LanguageDetected) String() string {
  return dara.Prettify(s)
}

func (s LanguageDetected) GoString() string {
  return s.String()
}

func (s *LanguageDetected) GetType() *string  {
  return s.Type
}

func (s *LanguageDetected) GetEventId() *string  {
  return s.EventId
}

func (s *LanguageDetected) GetSessionId() *string  {
  return s.SessionId
}

func (s *LanguageDetected) GetSegmentId() *string  {
  return s.SegmentId
}

func (s *LanguageDetected) GetDirection() *string  {
  return s.Direction
}

func (s *LanguageDetected) GetSourceLanguage() *string  {
  return s.SourceLanguage
}

func (s *LanguageDetected) GetTargetLanguage() *string  {
  return s.TargetLanguage
}

func (s *LanguageDetected) SetType(v string) *LanguageDetected {
  s.Type = &v
  return s
}

func (s *LanguageDetected) SetEventId(v string) *LanguageDetected {
  s.EventId = &v
  return s
}

func (s *LanguageDetected) SetSessionId(v string) *LanguageDetected {
  s.SessionId = &v
  return s
}

func (s *LanguageDetected) SetSegmentId(v string) *LanguageDetected {
  s.SegmentId = &v
  return s
}

func (s *LanguageDetected) SetDirection(v string) *LanguageDetected {
  s.Direction = &v
  return s
}

func (s *LanguageDetected) SetSourceLanguage(v string) *LanguageDetected {
  s.SourceLanguage = &v
  return s
}

func (s *LanguageDetected) SetTargetLanguage(v string) *LanguageDetected {
  s.TargetLanguage = &v
  return s
}

func (s *LanguageDetected) Validate() error {
  return dara.Validate(s)
}

