// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iResponseAudioDone interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *ResponseAudioDone
  GetType() *string 
  SetEventId(v string) *ResponseAudioDone
  GetEventId() *string 
  SetResponseId(v string) *ResponseAudioDone
  GetResponseId() *string 
  SetItemId(v string) *ResponseAudioDone
  GetItemId() *string 
  SetOutputIndex(v int32) *ResponseAudioDone
  GetOutputIndex() *int32 
  SetContentIndex(v int32) *ResponseAudioDone
  GetContentIndex() *int32 
}

type ResponseAudioDone struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  ResponseId *string `json:"response_id,omitempty" xml:"response_id,omitempty"`
  ItemId *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
  OutputIndex *int32 `json:"output_index,omitempty" xml:"output_index,omitempty"`
  ContentIndex *int32 `json:"content_index,omitempty" xml:"content_index,omitempty"`
}

func (s ResponseAudioDone) String() string {
  return dara.Prettify(s)
}

func (s ResponseAudioDone) GoString() string {
  return s.String()
}

func (s *ResponseAudioDone) GetType() *string  {
  return s.Type
}

func (s *ResponseAudioDone) GetEventId() *string  {
  return s.EventId
}

func (s *ResponseAudioDone) GetResponseId() *string  {
  return s.ResponseId
}

func (s *ResponseAudioDone) GetItemId() *string  {
  return s.ItemId
}

func (s *ResponseAudioDone) GetOutputIndex() *int32  {
  return s.OutputIndex
}

func (s *ResponseAudioDone) GetContentIndex() *int32  {
  return s.ContentIndex
}

func (s *ResponseAudioDone) SetType(v string) *ResponseAudioDone {
  s.Type = &v
  return s
}

func (s *ResponseAudioDone) SetEventId(v string) *ResponseAudioDone {
  s.EventId = &v
  return s
}

func (s *ResponseAudioDone) SetResponseId(v string) *ResponseAudioDone {
  s.ResponseId = &v
  return s
}

func (s *ResponseAudioDone) SetItemId(v string) *ResponseAudioDone {
  s.ItemId = &v
  return s
}

func (s *ResponseAudioDone) SetOutputIndex(v int32) *ResponseAudioDone {
  s.OutputIndex = &v
  return s
}

func (s *ResponseAudioDone) SetContentIndex(v int32) *ResponseAudioDone {
  s.ContentIndex = &v
  return s
}

func (s *ResponseAudioDone) Validate() error {
  return dara.Validate(s)
}

