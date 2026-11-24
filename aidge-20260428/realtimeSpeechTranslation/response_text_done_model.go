// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iResponseTextDone interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *ResponseTextDone
  GetType() *string 
  SetEventId(v string) *ResponseTextDone
  GetEventId() *string 
  SetResponseId(v string) *ResponseTextDone
  GetResponseId() *string 
  SetItemId(v string) *ResponseTextDone
  GetItemId() *string 
  SetOutputIndex(v int32) *ResponseTextDone
  GetOutputIndex() *int32 
  SetContentIndex(v int32) *ResponseTextDone
  GetContentIndex() *int32 
  SetText(v string) *ResponseTextDone
  GetText() *string 
}

type ResponseTextDone struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  ResponseId *string `json:"response_id,omitempty" xml:"response_id,omitempty"`
  ItemId *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
  OutputIndex *int32 `json:"output_index,omitempty" xml:"output_index,omitempty"`
  ContentIndex *int32 `json:"content_index,omitempty" xml:"content_index,omitempty"`
  Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s ResponseTextDone) String() string {
  return dara.Prettify(s)
}

func (s ResponseTextDone) GoString() string {
  return s.String()
}

func (s *ResponseTextDone) GetType() *string  {
  return s.Type
}

func (s *ResponseTextDone) GetEventId() *string  {
  return s.EventId
}

func (s *ResponseTextDone) GetResponseId() *string  {
  return s.ResponseId
}

func (s *ResponseTextDone) GetItemId() *string  {
  return s.ItemId
}

func (s *ResponseTextDone) GetOutputIndex() *int32  {
  return s.OutputIndex
}

func (s *ResponseTextDone) GetContentIndex() *int32  {
  return s.ContentIndex
}

func (s *ResponseTextDone) GetText() *string  {
  return s.Text
}

func (s *ResponseTextDone) SetType(v string) *ResponseTextDone {
  s.Type = &v
  return s
}

func (s *ResponseTextDone) SetEventId(v string) *ResponseTextDone {
  s.EventId = &v
  return s
}

func (s *ResponseTextDone) SetResponseId(v string) *ResponseTextDone {
  s.ResponseId = &v
  return s
}

func (s *ResponseTextDone) SetItemId(v string) *ResponseTextDone {
  s.ItemId = &v
  return s
}

func (s *ResponseTextDone) SetOutputIndex(v int32) *ResponseTextDone {
  s.OutputIndex = &v
  return s
}

func (s *ResponseTextDone) SetContentIndex(v int32) *ResponseTextDone {
  s.ContentIndex = &v
  return s
}

func (s *ResponseTextDone) SetText(v string) *ResponseTextDone {
  s.Text = &v
  return s
}

func (s *ResponseTextDone) Validate() error {
  return dara.Validate(s)
}

