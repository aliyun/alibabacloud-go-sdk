// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iTranslationTextDone interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *TranslationTextDone
  GetType() *string 
  SetEventId(v string) *TranslationTextDone
  GetEventId() *string 
  SetSessionId(v string) *TranslationTextDone
  GetSessionId() *string 
  SetSegmentId(v string) *TranslationTextDone
  GetSegmentId() *string 
  SetDirection(v string) *TranslationTextDone
  GetDirection() *string 
  SetSourceLanguage(v string) *TranslationTextDone
  GetSourceLanguage() *string 
  SetTargetLanguage(v string) *TranslationTextDone
  GetTargetLanguage() *string 
  SetText(v string) *TranslationTextDone
  GetText() *string 
}

type TranslationTextDone struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  SessionId *string `json:"session_id,omitempty" xml:"session_id,omitempty"`
  SegmentId *string `json:"segment_id,omitempty" xml:"segment_id,omitempty"`
  Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
  SourceLanguage *string `json:"source_language,omitempty" xml:"source_language,omitempty"`
  TargetLanguage *string `json:"target_language,omitempty" xml:"target_language,omitempty"`
  Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s TranslationTextDone) String() string {
  return dara.Prettify(s)
}

func (s TranslationTextDone) GoString() string {
  return s.String()
}

func (s *TranslationTextDone) GetType() *string  {
  return s.Type
}

func (s *TranslationTextDone) GetEventId() *string  {
  return s.EventId
}

func (s *TranslationTextDone) GetSessionId() *string  {
  return s.SessionId
}

func (s *TranslationTextDone) GetSegmentId() *string  {
  return s.SegmentId
}

func (s *TranslationTextDone) GetDirection() *string  {
  return s.Direction
}

func (s *TranslationTextDone) GetSourceLanguage() *string  {
  return s.SourceLanguage
}

func (s *TranslationTextDone) GetTargetLanguage() *string  {
  return s.TargetLanguage
}

func (s *TranslationTextDone) GetText() *string  {
  return s.Text
}

func (s *TranslationTextDone) SetType(v string) *TranslationTextDone {
  s.Type = &v
  return s
}

func (s *TranslationTextDone) SetEventId(v string) *TranslationTextDone {
  s.EventId = &v
  return s
}

func (s *TranslationTextDone) SetSessionId(v string) *TranslationTextDone {
  s.SessionId = &v
  return s
}

func (s *TranslationTextDone) SetSegmentId(v string) *TranslationTextDone {
  s.SegmentId = &v
  return s
}

func (s *TranslationTextDone) SetDirection(v string) *TranslationTextDone {
  s.Direction = &v
  return s
}

func (s *TranslationTextDone) SetSourceLanguage(v string) *TranslationTextDone {
  s.SourceLanguage = &v
  return s
}

func (s *TranslationTextDone) SetTargetLanguage(v string) *TranslationTextDone {
  s.TargetLanguage = &v
  return s
}

func (s *TranslationTextDone) SetText(v string) *TranslationTextDone {
  s.Text = &v
  return s
}

func (s *TranslationTextDone) Validate() error {
  return dara.Validate(s)
}

