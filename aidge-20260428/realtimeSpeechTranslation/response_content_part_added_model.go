// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iResponseContentPartAdded interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *ResponseContentPartAdded
  GetType() *string 
  SetEventId(v string) *ResponseContentPartAdded
  GetEventId() *string 
  SetResponseId(v string) *ResponseContentPartAdded
  GetResponseId() *string 
  SetItemId(v string) *ResponseContentPartAdded
  GetItemId() *string 
  SetOutputIndex(v int32) *ResponseContentPartAdded
  GetOutputIndex() *int32 
  SetContentIndex(v int32) *ResponseContentPartAdded
  GetContentIndex() *int32 
  SetPart(v *ResponseContentPartAddedPart) *ResponseContentPartAdded
  GetPart() *ResponseContentPartAddedPart 
}

type ResponseContentPartAdded struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  ResponseId *string `json:"response_id,omitempty" xml:"response_id,omitempty"`
  ItemId *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
  OutputIndex *int32 `json:"output_index,omitempty" xml:"output_index,omitempty"`
  ContentIndex *int32 `json:"content_index,omitempty" xml:"content_index,omitempty"`
  Part *ResponseContentPartAddedPart `json:"part,omitempty" xml:"part,omitempty" type:"Struct"`
}

func (s ResponseContentPartAdded) String() string {
  return dara.Prettify(s)
}

func (s ResponseContentPartAdded) GoString() string {
  return s.String()
}

func (s *ResponseContentPartAdded) GetType() *string  {
  return s.Type
}

func (s *ResponseContentPartAdded) GetEventId() *string  {
  return s.EventId
}

func (s *ResponseContentPartAdded) GetResponseId() *string  {
  return s.ResponseId
}

func (s *ResponseContentPartAdded) GetItemId() *string  {
  return s.ItemId
}

func (s *ResponseContentPartAdded) GetOutputIndex() *int32  {
  return s.OutputIndex
}

func (s *ResponseContentPartAdded) GetContentIndex() *int32  {
  return s.ContentIndex
}

func (s *ResponseContentPartAdded) GetPart() *ResponseContentPartAddedPart  {
  return s.Part
}

func (s *ResponseContentPartAdded) SetType(v string) *ResponseContentPartAdded {
  s.Type = &v
  return s
}

func (s *ResponseContentPartAdded) SetEventId(v string) *ResponseContentPartAdded {
  s.EventId = &v
  return s
}

func (s *ResponseContentPartAdded) SetResponseId(v string) *ResponseContentPartAdded {
  s.ResponseId = &v
  return s
}

func (s *ResponseContentPartAdded) SetItemId(v string) *ResponseContentPartAdded {
  s.ItemId = &v
  return s
}

func (s *ResponseContentPartAdded) SetOutputIndex(v int32) *ResponseContentPartAdded {
  s.OutputIndex = &v
  return s
}

func (s *ResponseContentPartAdded) SetContentIndex(v int32) *ResponseContentPartAdded {
  s.ContentIndex = &v
  return s
}

func (s *ResponseContentPartAdded) SetPart(v *ResponseContentPartAddedPart) *ResponseContentPartAdded {
  s.Part = v
  return s
}

func (s *ResponseContentPartAdded) Validate() error {
  if s.Part != nil {
    if err := s.Part.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ResponseContentPartAddedPart struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s ResponseContentPartAddedPart) String() string {
  return dara.Prettify(s)
}

func (s ResponseContentPartAddedPart) GoString() string {
  return s.String()
}

func (s *ResponseContentPartAddedPart) GetType() *string  {
  return s.Type
}

func (s *ResponseContentPartAddedPart) GetText() *string  {
  return s.Text
}

func (s *ResponseContentPartAddedPart) SetType(v string) *ResponseContentPartAddedPart {
  s.Type = &v
  return s
}

func (s *ResponseContentPartAddedPart) SetText(v string) *ResponseContentPartAddedPart {
  s.Text = &v
  return s
}

func (s *ResponseContentPartAddedPart) Validate() error {
  return dara.Validate(s)
}

