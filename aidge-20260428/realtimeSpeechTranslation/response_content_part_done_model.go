// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iResponseContentPartDone interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *ResponseContentPartDone
  GetType() *string 
  SetEventId(v string) *ResponseContentPartDone
  GetEventId() *string 
  SetResponseId(v string) *ResponseContentPartDone
  GetResponseId() *string 
  SetItemId(v string) *ResponseContentPartDone
  GetItemId() *string 
  SetOutputIndex(v int32) *ResponseContentPartDone
  GetOutputIndex() *int32 
  SetContentIndex(v int32) *ResponseContentPartDone
  GetContentIndex() *int32 
  SetPart(v *ResponseContentPartDonePart) *ResponseContentPartDone
  GetPart() *ResponseContentPartDonePart 
}

type ResponseContentPartDone struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  ResponseId *string `json:"response_id,omitempty" xml:"response_id,omitempty"`
  ItemId *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
  OutputIndex *int32 `json:"output_index,omitempty" xml:"output_index,omitempty"`
  ContentIndex *int32 `json:"content_index,omitempty" xml:"content_index,omitempty"`
  Part *ResponseContentPartDonePart `json:"part,omitempty" xml:"part,omitempty" type:"Struct"`
}

func (s ResponseContentPartDone) String() string {
  return dara.Prettify(s)
}

func (s ResponseContentPartDone) GoString() string {
  return s.String()
}

func (s *ResponseContentPartDone) GetType() *string  {
  return s.Type
}

func (s *ResponseContentPartDone) GetEventId() *string  {
  return s.EventId
}

func (s *ResponseContentPartDone) GetResponseId() *string  {
  return s.ResponseId
}

func (s *ResponseContentPartDone) GetItemId() *string  {
  return s.ItemId
}

func (s *ResponseContentPartDone) GetOutputIndex() *int32  {
  return s.OutputIndex
}

func (s *ResponseContentPartDone) GetContentIndex() *int32  {
  return s.ContentIndex
}

func (s *ResponseContentPartDone) GetPart() *ResponseContentPartDonePart  {
  return s.Part
}

func (s *ResponseContentPartDone) SetType(v string) *ResponseContentPartDone {
  s.Type = &v
  return s
}

func (s *ResponseContentPartDone) SetEventId(v string) *ResponseContentPartDone {
  s.EventId = &v
  return s
}

func (s *ResponseContentPartDone) SetResponseId(v string) *ResponseContentPartDone {
  s.ResponseId = &v
  return s
}

func (s *ResponseContentPartDone) SetItemId(v string) *ResponseContentPartDone {
  s.ItemId = &v
  return s
}

func (s *ResponseContentPartDone) SetOutputIndex(v int32) *ResponseContentPartDone {
  s.OutputIndex = &v
  return s
}

func (s *ResponseContentPartDone) SetContentIndex(v int32) *ResponseContentPartDone {
  s.ContentIndex = &v
  return s
}

func (s *ResponseContentPartDone) SetPart(v *ResponseContentPartDonePart) *ResponseContentPartDone {
  s.Part = v
  return s
}

func (s *ResponseContentPartDone) Validate() error {
  if s.Part != nil {
    if err := s.Part.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ResponseContentPartDonePart struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s ResponseContentPartDonePart) String() string {
  return dara.Prettify(s)
}

func (s ResponseContentPartDonePart) GoString() string {
  return s.String()
}

func (s *ResponseContentPartDonePart) GetType() *string  {
  return s.Type
}

func (s *ResponseContentPartDonePart) GetText() *string  {
  return s.Text
}

func (s *ResponseContentPartDonePart) SetType(v string) *ResponseContentPartDonePart {
  s.Type = &v
  return s
}

func (s *ResponseContentPartDonePart) SetText(v string) *ResponseContentPartDonePart {
  s.Text = &v
  return s
}

func (s *ResponseContentPartDonePart) Validate() error {
  return dara.Validate(s)
}

