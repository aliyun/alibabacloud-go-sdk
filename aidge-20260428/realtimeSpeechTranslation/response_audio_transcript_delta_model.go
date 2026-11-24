// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iResponseAudioTranscriptDelta interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *ResponseAudioTranscriptDelta
  GetType() *string 
  SetEventId(v string) *ResponseAudioTranscriptDelta
  GetEventId() *string 
  SetResponseId(v string) *ResponseAudioTranscriptDelta
  GetResponseId() *string 
  SetItemId(v string) *ResponseAudioTranscriptDelta
  GetItemId() *string 
  SetOutputIndex(v int32) *ResponseAudioTranscriptDelta
  GetOutputIndex() *int32 
  SetContentIndex(v int32) *ResponseAudioTranscriptDelta
  GetContentIndex() *int32 
  SetDelta(v string) *ResponseAudioTranscriptDelta
  GetDelta() *string 
}

type ResponseAudioTranscriptDelta struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  ResponseId *string `json:"response_id,omitempty" xml:"response_id,omitempty"`
  ItemId *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
  OutputIndex *int32 `json:"output_index,omitempty" xml:"output_index,omitempty"`
  ContentIndex *int32 `json:"content_index,omitempty" xml:"content_index,omitempty"`
  Delta *string `json:"delta,omitempty" xml:"delta,omitempty"`
}

func (s ResponseAudioTranscriptDelta) String() string {
  return dara.Prettify(s)
}

func (s ResponseAudioTranscriptDelta) GoString() string {
  return s.String()
}

func (s *ResponseAudioTranscriptDelta) GetType() *string  {
  return s.Type
}

func (s *ResponseAudioTranscriptDelta) GetEventId() *string  {
  return s.EventId
}

func (s *ResponseAudioTranscriptDelta) GetResponseId() *string  {
  return s.ResponseId
}

func (s *ResponseAudioTranscriptDelta) GetItemId() *string  {
  return s.ItemId
}

func (s *ResponseAudioTranscriptDelta) GetOutputIndex() *int32  {
  return s.OutputIndex
}

func (s *ResponseAudioTranscriptDelta) GetContentIndex() *int32  {
  return s.ContentIndex
}

func (s *ResponseAudioTranscriptDelta) GetDelta() *string  {
  return s.Delta
}

func (s *ResponseAudioTranscriptDelta) SetType(v string) *ResponseAudioTranscriptDelta {
  s.Type = &v
  return s
}

func (s *ResponseAudioTranscriptDelta) SetEventId(v string) *ResponseAudioTranscriptDelta {
  s.EventId = &v
  return s
}

func (s *ResponseAudioTranscriptDelta) SetResponseId(v string) *ResponseAudioTranscriptDelta {
  s.ResponseId = &v
  return s
}

func (s *ResponseAudioTranscriptDelta) SetItemId(v string) *ResponseAudioTranscriptDelta {
  s.ItemId = &v
  return s
}

func (s *ResponseAudioTranscriptDelta) SetOutputIndex(v int32) *ResponseAudioTranscriptDelta {
  s.OutputIndex = &v
  return s
}

func (s *ResponseAudioTranscriptDelta) SetContentIndex(v int32) *ResponseAudioTranscriptDelta {
  s.ContentIndex = &v
  return s
}

func (s *ResponseAudioTranscriptDelta) SetDelta(v string) *ResponseAudioTranscriptDelta {
  s.Delta = &v
  return s
}

func (s *ResponseAudioTranscriptDelta) Validate() error {
  return dara.Validate(s)
}

