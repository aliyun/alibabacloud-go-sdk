// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iInputAudioBufferCommitted interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *InputAudioBufferCommitted
  GetType() *string 
  SetEventId(v string) *InputAudioBufferCommitted
  GetEventId() *string 
  SetSessionId(v string) *InputAudioBufferCommitted
  GetSessionId() *string 
  SetSegmentId(v string) *InputAudioBufferCommitted
  GetSegmentId() *string 
  SetCausedBy(v string) *InputAudioBufferCommitted
  GetCausedBy() *string 
  SetAudioDurationMs(v int64) *InputAudioBufferCommitted
  GetAudioDurationMs() *int64 
  SetItemId(v string) *InputAudioBufferCommitted
  GetItemId() *string 
  SetPreviousItemId(v string) *InputAudioBufferCommitted
  GetPreviousItemId() *string 
}

type InputAudioBufferCommitted struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  SessionId *string `json:"session_id,omitempty" xml:"session_id,omitempty"`
  SegmentId *string `json:"segment_id,omitempty" xml:"segment_id,omitempty"`
  CausedBy *string `json:"caused_by,omitempty" xml:"caused_by,omitempty"`
  AudioDurationMs *int64 `json:"audio_duration_ms,omitempty" xml:"audio_duration_ms,omitempty"`
  ItemId *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
  PreviousItemId *string `json:"previous_item_id,omitempty" xml:"previous_item_id,omitempty"`
}

func (s InputAudioBufferCommitted) String() string {
  return dara.Prettify(s)
}

func (s InputAudioBufferCommitted) GoString() string {
  return s.String()
}

func (s *InputAudioBufferCommitted) GetType() *string  {
  return s.Type
}

func (s *InputAudioBufferCommitted) GetEventId() *string  {
  return s.EventId
}

func (s *InputAudioBufferCommitted) GetSessionId() *string  {
  return s.SessionId
}

func (s *InputAudioBufferCommitted) GetSegmentId() *string  {
  return s.SegmentId
}

func (s *InputAudioBufferCommitted) GetCausedBy() *string  {
  return s.CausedBy
}

func (s *InputAudioBufferCommitted) GetAudioDurationMs() *int64  {
  return s.AudioDurationMs
}

func (s *InputAudioBufferCommitted) GetItemId() *string  {
  return s.ItemId
}

func (s *InputAudioBufferCommitted) GetPreviousItemId() *string  {
  return s.PreviousItemId
}

func (s *InputAudioBufferCommitted) SetType(v string) *InputAudioBufferCommitted {
  s.Type = &v
  return s
}

func (s *InputAudioBufferCommitted) SetEventId(v string) *InputAudioBufferCommitted {
  s.EventId = &v
  return s
}

func (s *InputAudioBufferCommitted) SetSessionId(v string) *InputAudioBufferCommitted {
  s.SessionId = &v
  return s
}

func (s *InputAudioBufferCommitted) SetSegmentId(v string) *InputAudioBufferCommitted {
  s.SegmentId = &v
  return s
}

func (s *InputAudioBufferCommitted) SetCausedBy(v string) *InputAudioBufferCommitted {
  s.CausedBy = &v
  return s
}

func (s *InputAudioBufferCommitted) SetAudioDurationMs(v int64) *InputAudioBufferCommitted {
  s.AudioDurationMs = &v
  return s
}

func (s *InputAudioBufferCommitted) SetItemId(v string) *InputAudioBufferCommitted {
  s.ItemId = &v
  return s
}

func (s *InputAudioBufferCommitted) SetPreviousItemId(v string) *InputAudioBufferCommitted {
  s.PreviousItemId = &v
  return s
}

func (s *InputAudioBufferCommitted) Validate() error {
  return dara.Validate(s)
}

