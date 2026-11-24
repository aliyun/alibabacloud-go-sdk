// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iResponseAudioTranscriptText interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *ResponseAudioTranscriptText
  GetType() *string 
  SetEventId(v string) *ResponseAudioTranscriptText
  GetEventId() *string 
  SetResponseId(v string) *ResponseAudioTranscriptText
  GetResponseId() *string 
  SetItemId(v string) *ResponseAudioTranscriptText
  GetItemId() *string 
  SetOutputIndex(v int32) *ResponseAudioTranscriptText
  GetOutputIndex() *int32 
  SetContentIndex(v int32) *ResponseAudioTranscriptText
  GetContentIndex() *int32 
  SetText(v string) *ResponseAudioTranscriptText
  GetText() *string 
  SetStash(v string) *ResponseAudioTranscriptText
  GetStash() *string 
}

type ResponseAudioTranscriptText struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  ResponseId *string `json:"response_id,omitempty" xml:"response_id,omitempty"`
  ItemId *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
  OutputIndex *int32 `json:"output_index,omitempty" xml:"output_index,omitempty"`
  ContentIndex *int32 `json:"content_index,omitempty" xml:"content_index,omitempty"`
  Text *string `json:"text,omitempty" xml:"text,omitempty"`
  Stash *string `json:"stash,omitempty" xml:"stash,omitempty"`
}

func (s ResponseAudioTranscriptText) String() string {
  return dara.Prettify(s)
}

func (s ResponseAudioTranscriptText) GoString() string {
  return s.String()
}

func (s *ResponseAudioTranscriptText) GetType() *string  {
  return s.Type
}

func (s *ResponseAudioTranscriptText) GetEventId() *string  {
  return s.EventId
}

func (s *ResponseAudioTranscriptText) GetResponseId() *string  {
  return s.ResponseId
}

func (s *ResponseAudioTranscriptText) GetItemId() *string  {
  return s.ItemId
}

func (s *ResponseAudioTranscriptText) GetOutputIndex() *int32  {
  return s.OutputIndex
}

func (s *ResponseAudioTranscriptText) GetContentIndex() *int32  {
  return s.ContentIndex
}

func (s *ResponseAudioTranscriptText) GetText() *string  {
  return s.Text
}

func (s *ResponseAudioTranscriptText) GetStash() *string  {
  return s.Stash
}

func (s *ResponseAudioTranscriptText) SetType(v string) *ResponseAudioTranscriptText {
  s.Type = &v
  return s
}

func (s *ResponseAudioTranscriptText) SetEventId(v string) *ResponseAudioTranscriptText {
  s.EventId = &v
  return s
}

func (s *ResponseAudioTranscriptText) SetResponseId(v string) *ResponseAudioTranscriptText {
  s.ResponseId = &v
  return s
}

func (s *ResponseAudioTranscriptText) SetItemId(v string) *ResponseAudioTranscriptText {
  s.ItemId = &v
  return s
}

func (s *ResponseAudioTranscriptText) SetOutputIndex(v int32) *ResponseAudioTranscriptText {
  s.OutputIndex = &v
  return s
}

func (s *ResponseAudioTranscriptText) SetContentIndex(v int32) *ResponseAudioTranscriptText {
  s.ContentIndex = &v
  return s
}

func (s *ResponseAudioTranscriptText) SetText(v string) *ResponseAudioTranscriptText {
  s.Text = &v
  return s
}

func (s *ResponseAudioTranscriptText) SetStash(v string) *ResponseAudioTranscriptText {
  s.Stash = &v
  return s
}

func (s *ResponseAudioTranscriptText) Validate() error {
  return dara.Validate(s)
}

