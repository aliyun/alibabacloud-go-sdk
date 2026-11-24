// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iTranslationTextDelta interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *TranslationTextDelta
  GetType() *string 
  SetEventId(v string) *TranslationTextDelta
  GetEventId() *string 
  SetSessionId(v string) *TranslationTextDelta
  GetSessionId() *string 
  SetSegmentId(v string) *TranslationTextDelta
  GetSegmentId() *string 
  SetDirection(v string) *TranslationTextDelta
  GetDirection() *string 
  SetSourceLanguage(v string) *TranslationTextDelta
  GetSourceLanguage() *string 
  SetTargetLanguage(v string) *TranslationTextDelta
  GetTargetLanguage() *string 
  SetText(v string) *TranslationTextDelta
  GetText() *string 
  SetStash(v string) *TranslationTextDelta
  GetStash() *string 
}

type TranslationTextDelta struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  SessionId *string `json:"session_id,omitempty" xml:"session_id,omitempty"`
  SegmentId *string `json:"segment_id,omitempty" xml:"segment_id,omitempty"`
  Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
  SourceLanguage *string `json:"source_language,omitempty" xml:"source_language,omitempty"`
  TargetLanguage *string `json:"target_language,omitempty" xml:"target_language,omitempty"`
  Text *string `json:"text,omitempty" xml:"text,omitempty"`
  Stash *string `json:"stash,omitempty" xml:"stash,omitempty"`
}

func (s TranslationTextDelta) String() string {
  return dara.Prettify(s)
}

func (s TranslationTextDelta) GoString() string {
  return s.String()
}

func (s *TranslationTextDelta) GetType() *string  {
  return s.Type
}

func (s *TranslationTextDelta) GetEventId() *string  {
  return s.EventId
}

func (s *TranslationTextDelta) GetSessionId() *string  {
  return s.SessionId
}

func (s *TranslationTextDelta) GetSegmentId() *string  {
  return s.SegmentId
}

func (s *TranslationTextDelta) GetDirection() *string  {
  return s.Direction
}

func (s *TranslationTextDelta) GetSourceLanguage() *string  {
  return s.SourceLanguage
}

func (s *TranslationTextDelta) GetTargetLanguage() *string  {
  return s.TargetLanguage
}

func (s *TranslationTextDelta) GetText() *string  {
  return s.Text
}

func (s *TranslationTextDelta) GetStash() *string  {
  return s.Stash
}

func (s *TranslationTextDelta) SetType(v string) *TranslationTextDelta {
  s.Type = &v
  return s
}

func (s *TranslationTextDelta) SetEventId(v string) *TranslationTextDelta {
  s.EventId = &v
  return s
}

func (s *TranslationTextDelta) SetSessionId(v string) *TranslationTextDelta {
  s.SessionId = &v
  return s
}

func (s *TranslationTextDelta) SetSegmentId(v string) *TranslationTextDelta {
  s.SegmentId = &v
  return s
}

func (s *TranslationTextDelta) SetDirection(v string) *TranslationTextDelta {
  s.Direction = &v
  return s
}

func (s *TranslationTextDelta) SetSourceLanguage(v string) *TranslationTextDelta {
  s.SourceLanguage = &v
  return s
}

func (s *TranslationTextDelta) SetTargetLanguage(v string) *TranslationTextDelta {
  s.TargetLanguage = &v
  return s
}

func (s *TranslationTextDelta) SetText(v string) *TranslationTextDelta {
  s.Text = &v
  return s
}

func (s *TranslationTextDelta) SetStash(v string) *TranslationTextDelta {
  s.Stash = &v
  return s
}

func (s *TranslationTextDelta) Validate() error {
  return dara.Validate(s)
}

