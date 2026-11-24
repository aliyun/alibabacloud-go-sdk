// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iResponseAudioDelta interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *ResponseAudioDelta
  GetType() *string 
  SetEventId(v string) *ResponseAudioDelta
  GetEventId() *string 
  SetResponseId(v string) *ResponseAudioDelta
  GetResponseId() *string 
  SetItemId(v string) *ResponseAudioDelta
  GetItemId() *string 
  SetOutputIndex(v int32) *ResponseAudioDelta
  GetOutputIndex() *int32 
  SetContentIndex(v int32) *ResponseAudioDelta
  GetContentIndex() *int32 
  SetDelta(v string) *ResponseAudioDelta
  GetDelta() *string 
}

type ResponseAudioDelta struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  ResponseId *string `json:"response_id,omitempty" xml:"response_id,omitempty"`
  ItemId *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
  OutputIndex *int32 `json:"output_index,omitempty" xml:"output_index,omitempty"`
  ContentIndex *int32 `json:"content_index,omitempty" xml:"content_index,omitempty"`
  Delta *string `json:"delta,omitempty" xml:"delta,omitempty"`
}

func (s ResponseAudioDelta) String() string {
  return dara.Prettify(s)
}

func (s ResponseAudioDelta) GoString() string {
  return s.String()
}

func (s *ResponseAudioDelta) GetType() *string  {
  return s.Type
}

func (s *ResponseAudioDelta) GetEventId() *string  {
  return s.EventId
}

func (s *ResponseAudioDelta) GetResponseId() *string  {
  return s.ResponseId
}

func (s *ResponseAudioDelta) GetItemId() *string  {
  return s.ItemId
}

func (s *ResponseAudioDelta) GetOutputIndex() *int32  {
  return s.OutputIndex
}

func (s *ResponseAudioDelta) GetContentIndex() *int32  {
  return s.ContentIndex
}

func (s *ResponseAudioDelta) GetDelta() *string  {
  return s.Delta
}

func (s *ResponseAudioDelta) SetType(v string) *ResponseAudioDelta {
  s.Type = &v
  return s
}

func (s *ResponseAudioDelta) SetEventId(v string) *ResponseAudioDelta {
  s.EventId = &v
  return s
}

func (s *ResponseAudioDelta) SetResponseId(v string) *ResponseAudioDelta {
  s.ResponseId = &v
  return s
}

func (s *ResponseAudioDelta) SetItemId(v string) *ResponseAudioDelta {
  s.ItemId = &v
  return s
}

func (s *ResponseAudioDelta) SetOutputIndex(v int32) *ResponseAudioDelta {
  s.OutputIndex = &v
  return s
}

func (s *ResponseAudioDelta) SetContentIndex(v int32) *ResponseAudioDelta {
  s.ContentIndex = &v
  return s
}

func (s *ResponseAudioDelta) SetDelta(v string) *ResponseAudioDelta {
  s.Delta = &v
  return s
}

func (s *ResponseAudioDelta) Validate() error {
  return dara.Validate(s)
}

