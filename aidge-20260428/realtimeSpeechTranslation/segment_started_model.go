// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iSegmentStarted interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *SegmentStarted
  GetType() *string 
  SetEventId(v string) *SegmentStarted
  GetEventId() *string 
  SetSessionId(v string) *SegmentStarted
  GetSessionId() *string 
  SetSegmentId(v string) *SegmentStarted
  GetSegmentId() *string 
  SetStartMs(v int64) *SegmentStarted
  GetStartMs() *int64 
}

type SegmentStarted struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  SessionId *string `json:"session_id,omitempty" xml:"session_id,omitempty"`
  SegmentId *string `json:"segment_id,omitempty" xml:"segment_id,omitempty"`
  StartMs *int64 `json:"start_ms,omitempty" xml:"start_ms,omitempty"`
}

func (s SegmentStarted) String() string {
  return dara.Prettify(s)
}

func (s SegmentStarted) GoString() string {
  return s.String()
}

func (s *SegmentStarted) GetType() *string  {
  return s.Type
}

func (s *SegmentStarted) GetEventId() *string  {
  return s.EventId
}

func (s *SegmentStarted) GetSessionId() *string  {
  return s.SessionId
}

func (s *SegmentStarted) GetSegmentId() *string  {
  return s.SegmentId
}

func (s *SegmentStarted) GetStartMs() *int64  {
  return s.StartMs
}

func (s *SegmentStarted) SetType(v string) *SegmentStarted {
  s.Type = &v
  return s
}

func (s *SegmentStarted) SetEventId(v string) *SegmentStarted {
  s.EventId = &v
  return s
}

func (s *SegmentStarted) SetSessionId(v string) *SegmentStarted {
  s.SessionId = &v
  return s
}

func (s *SegmentStarted) SetSegmentId(v string) *SegmentStarted {
  s.SegmentId = &v
  return s
}

func (s *SegmentStarted) SetStartMs(v int64) *SegmentStarted {
  s.StartMs = &v
  return s
}

func (s *SegmentStarted) Validate() error {
  return dara.Validate(s)
}

