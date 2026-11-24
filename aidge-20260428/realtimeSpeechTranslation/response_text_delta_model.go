// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iResponseTextDelta interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *ResponseTextDelta
  GetType() *string 
  SetEventId(v string) *ResponseTextDelta
  GetEventId() *string 
  SetResponseId(v string) *ResponseTextDelta
  GetResponseId() *string 
  SetItemId(v string) *ResponseTextDelta
  GetItemId() *string 
  SetOutputIndex(v int32) *ResponseTextDelta
  GetOutputIndex() *int32 
  SetContentIndex(v int32) *ResponseTextDelta
  GetContentIndex() *int32 
  SetDelta(v string) *ResponseTextDelta
  GetDelta() *string 
}

type ResponseTextDelta struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  ResponseId *string `json:"response_id,omitempty" xml:"response_id,omitempty"`
  ItemId *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
  OutputIndex *int32 `json:"output_index,omitempty" xml:"output_index,omitempty"`
  ContentIndex *int32 `json:"content_index,omitempty" xml:"content_index,omitempty"`
  Delta *string `json:"delta,omitempty" xml:"delta,omitempty"`
}

func (s ResponseTextDelta) String() string {
  return dara.Prettify(s)
}

func (s ResponseTextDelta) GoString() string {
  return s.String()
}

func (s *ResponseTextDelta) GetType() *string  {
  return s.Type
}

func (s *ResponseTextDelta) GetEventId() *string  {
  return s.EventId
}

func (s *ResponseTextDelta) GetResponseId() *string  {
  return s.ResponseId
}

func (s *ResponseTextDelta) GetItemId() *string  {
  return s.ItemId
}

func (s *ResponseTextDelta) GetOutputIndex() *int32  {
  return s.OutputIndex
}

func (s *ResponseTextDelta) GetContentIndex() *int32  {
  return s.ContentIndex
}

func (s *ResponseTextDelta) GetDelta() *string  {
  return s.Delta
}

func (s *ResponseTextDelta) SetType(v string) *ResponseTextDelta {
  s.Type = &v
  return s
}

func (s *ResponseTextDelta) SetEventId(v string) *ResponseTextDelta {
  s.EventId = &v
  return s
}

func (s *ResponseTextDelta) SetResponseId(v string) *ResponseTextDelta {
  s.ResponseId = &v
  return s
}

func (s *ResponseTextDelta) SetItemId(v string) *ResponseTextDelta {
  s.ItemId = &v
  return s
}

func (s *ResponseTextDelta) SetOutputIndex(v int32) *ResponseTextDelta {
  s.OutputIndex = &v
  return s
}

func (s *ResponseTextDelta) SetContentIndex(v int32) *ResponseTextDelta {
  s.ContentIndex = &v
  return s
}

func (s *ResponseTextDelta) SetDelta(v string) *ResponseTextDelta {
  s.Delta = &v
  return s
}

func (s *ResponseTextDelta) Validate() error {
  return dara.Validate(s)
}

