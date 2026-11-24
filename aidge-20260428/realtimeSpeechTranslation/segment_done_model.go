// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iSegmentDone interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *SegmentDone
  GetType() *string 
  SetEventId(v string) *SegmentDone
  GetEventId() *string 
  SetSessionId(v string) *SegmentDone
  GetSessionId() *string 
  SetSegmentId(v string) *SegmentDone
  GetSegmentId() *string 
  SetDirection(v string) *SegmentDone
  GetDirection() *string 
  SetSourceLanguage(v string) *SegmentDone
  GetSourceLanguage() *string 
  SetTargetLanguage(v string) *SegmentDone
  GetTargetLanguage() *string 
  SetStartMs(v int64) *SegmentDone
  GetStartMs() *int64 
  SetEndMs(v int64) *SegmentDone
  GetEndMs() *int64 
}

type SegmentDone struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  SessionId *string `json:"session_id,omitempty" xml:"session_id,omitempty"`
  SegmentId *string `json:"segment_id,omitempty" xml:"segment_id,omitempty"`
  Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
  SourceLanguage *string `json:"source_language,omitempty" xml:"source_language,omitempty"`
  TargetLanguage *string `json:"target_language,omitempty" xml:"target_language,omitempty"`
  StartMs *int64 `json:"start_ms,omitempty" xml:"start_ms,omitempty"`
  EndMs *int64 `json:"end_ms,omitempty" xml:"end_ms,omitempty"`
}

func (s SegmentDone) String() string {
  return dara.Prettify(s)
}

func (s SegmentDone) GoString() string {
  return s.String()
}

func (s *SegmentDone) GetType() *string  {
  return s.Type
}

func (s *SegmentDone) GetEventId() *string  {
  return s.EventId
}

func (s *SegmentDone) GetSessionId() *string  {
  return s.SessionId
}

func (s *SegmentDone) GetSegmentId() *string  {
  return s.SegmentId
}

func (s *SegmentDone) GetDirection() *string  {
  return s.Direction
}

func (s *SegmentDone) GetSourceLanguage() *string  {
  return s.SourceLanguage
}

func (s *SegmentDone) GetTargetLanguage() *string  {
  return s.TargetLanguage
}

func (s *SegmentDone) GetStartMs() *int64  {
  return s.StartMs
}

func (s *SegmentDone) GetEndMs() *int64  {
  return s.EndMs
}

func (s *SegmentDone) SetType(v string) *SegmentDone {
  s.Type = &v
  return s
}

func (s *SegmentDone) SetEventId(v string) *SegmentDone {
  s.EventId = &v
  return s
}

func (s *SegmentDone) SetSessionId(v string) *SegmentDone {
  s.SessionId = &v
  return s
}

func (s *SegmentDone) SetSegmentId(v string) *SegmentDone {
  s.SegmentId = &v
  return s
}

func (s *SegmentDone) SetDirection(v string) *SegmentDone {
  s.Direction = &v
  return s
}

func (s *SegmentDone) SetSourceLanguage(v string) *SegmentDone {
  s.SourceLanguage = &v
  return s
}

func (s *SegmentDone) SetTargetLanguage(v string) *SegmentDone {
  s.TargetLanguage = &v
  return s
}

func (s *SegmentDone) SetStartMs(v int64) *SegmentDone {
  s.StartMs = &v
  return s
}

func (s *SegmentDone) SetEndMs(v int64) *SegmentDone {
  s.EndMs = &v
  return s
}

func (s *SegmentDone) Validate() error {
  return dara.Validate(s)
}

