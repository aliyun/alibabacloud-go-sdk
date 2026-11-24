// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iResponseAudioTranscriptDone interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *ResponseAudioTranscriptDone
  GetType() *string 
  SetEventId(v string) *ResponseAudioTranscriptDone
  GetEventId() *string 
  SetResponseId(v string) *ResponseAudioTranscriptDone
  GetResponseId() *string 
  SetItemId(v string) *ResponseAudioTranscriptDone
  GetItemId() *string 
  SetOutputIndex(v int32) *ResponseAudioTranscriptDone
  GetOutputIndex() *int32 
  SetContentIndex(v int32) *ResponseAudioTranscriptDone
  GetContentIndex() *int32 
  SetTranscript(v string) *ResponseAudioTranscriptDone
  GetTranscript() *string 
}

type ResponseAudioTranscriptDone struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  ResponseId *string `json:"response_id,omitempty" xml:"response_id,omitempty"`
  ItemId *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
  OutputIndex *int32 `json:"output_index,omitempty" xml:"output_index,omitempty"`
  ContentIndex *int32 `json:"content_index,omitempty" xml:"content_index,omitempty"`
  Transcript *string `json:"transcript,omitempty" xml:"transcript,omitempty"`
}

func (s ResponseAudioTranscriptDone) String() string {
  return dara.Prettify(s)
}

func (s ResponseAudioTranscriptDone) GoString() string {
  return s.String()
}

func (s *ResponseAudioTranscriptDone) GetType() *string  {
  return s.Type
}

func (s *ResponseAudioTranscriptDone) GetEventId() *string  {
  return s.EventId
}

func (s *ResponseAudioTranscriptDone) GetResponseId() *string  {
  return s.ResponseId
}

func (s *ResponseAudioTranscriptDone) GetItemId() *string  {
  return s.ItemId
}

func (s *ResponseAudioTranscriptDone) GetOutputIndex() *int32  {
  return s.OutputIndex
}

func (s *ResponseAudioTranscriptDone) GetContentIndex() *int32  {
  return s.ContentIndex
}

func (s *ResponseAudioTranscriptDone) GetTranscript() *string  {
  return s.Transcript
}

func (s *ResponseAudioTranscriptDone) SetType(v string) *ResponseAudioTranscriptDone {
  s.Type = &v
  return s
}

func (s *ResponseAudioTranscriptDone) SetEventId(v string) *ResponseAudioTranscriptDone {
  s.EventId = &v
  return s
}

func (s *ResponseAudioTranscriptDone) SetResponseId(v string) *ResponseAudioTranscriptDone {
  s.ResponseId = &v
  return s
}

func (s *ResponseAudioTranscriptDone) SetItemId(v string) *ResponseAudioTranscriptDone {
  s.ItemId = &v
  return s
}

func (s *ResponseAudioTranscriptDone) SetOutputIndex(v int32) *ResponseAudioTranscriptDone {
  s.OutputIndex = &v
  return s
}

func (s *ResponseAudioTranscriptDone) SetContentIndex(v int32) *ResponseAudioTranscriptDone {
  s.ContentIndex = &v
  return s
}

func (s *ResponseAudioTranscriptDone) SetTranscript(v string) *ResponseAudioTranscriptDone {
  s.Transcript = &v
  return s
}

func (s *ResponseAudioTranscriptDone) Validate() error {
  return dara.Validate(s)
}

