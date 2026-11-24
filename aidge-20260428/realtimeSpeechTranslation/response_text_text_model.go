// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iResponseTextText interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *ResponseTextText
  GetType() *string 
  SetEventId(v string) *ResponseTextText
  GetEventId() *string 
  SetResponseId(v string) *ResponseTextText
  GetResponseId() *string 
  SetItemId(v string) *ResponseTextText
  GetItemId() *string 
  SetOutputIndex(v int32) *ResponseTextText
  GetOutputIndex() *int32 
  SetContentIndex(v int32) *ResponseTextText
  GetContentIndex() *int32 
  SetText(v string) *ResponseTextText
  GetText() *string 
  SetStash(v string) *ResponseTextText
  GetStash() *string 
}

type ResponseTextText struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  ResponseId *string `json:"response_id,omitempty" xml:"response_id,omitempty"`
  ItemId *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
  OutputIndex *int32 `json:"output_index,omitempty" xml:"output_index,omitempty"`
  ContentIndex *int32 `json:"content_index,omitempty" xml:"content_index,omitempty"`
  Text *string `json:"text,omitempty" xml:"text,omitempty"`
  Stash *string `json:"stash,omitempty" xml:"stash,omitempty"`
}

func (s ResponseTextText) String() string {
  return dara.Prettify(s)
}

func (s ResponseTextText) GoString() string {
  return s.String()
}

func (s *ResponseTextText) GetType() *string  {
  return s.Type
}

func (s *ResponseTextText) GetEventId() *string  {
  return s.EventId
}

func (s *ResponseTextText) GetResponseId() *string  {
  return s.ResponseId
}

func (s *ResponseTextText) GetItemId() *string  {
  return s.ItemId
}

func (s *ResponseTextText) GetOutputIndex() *int32  {
  return s.OutputIndex
}

func (s *ResponseTextText) GetContentIndex() *int32  {
  return s.ContentIndex
}

func (s *ResponseTextText) GetText() *string  {
  return s.Text
}

func (s *ResponseTextText) GetStash() *string  {
  return s.Stash
}

func (s *ResponseTextText) SetType(v string) *ResponseTextText {
  s.Type = &v
  return s
}

func (s *ResponseTextText) SetEventId(v string) *ResponseTextText {
  s.EventId = &v
  return s
}

func (s *ResponseTextText) SetResponseId(v string) *ResponseTextText {
  s.ResponseId = &v
  return s
}

func (s *ResponseTextText) SetItemId(v string) *ResponseTextText {
  s.ItemId = &v
  return s
}

func (s *ResponseTextText) SetOutputIndex(v int32) *ResponseTextText {
  s.OutputIndex = &v
  return s
}

func (s *ResponseTextText) SetContentIndex(v int32) *ResponseTextText {
  s.ContentIndex = &v
  return s
}

func (s *ResponseTextText) SetText(v string) *ResponseTextText {
  s.Text = &v
  return s
}

func (s *ResponseTextText) SetStash(v string) *ResponseTextText {
  s.Stash = &v
  return s
}

func (s *ResponseTextText) Validate() error {
  return dara.Validate(s)
}

