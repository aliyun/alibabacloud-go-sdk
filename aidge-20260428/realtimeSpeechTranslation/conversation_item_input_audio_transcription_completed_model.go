// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iConversationItemInputAudioTranscriptionCompleted interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *ConversationItemInputAudioTranscriptionCompleted
  GetType() *string 
  SetEventId(v string) *ConversationItemInputAudioTranscriptionCompleted
  GetEventId() *string 
  SetItemId(v string) *ConversationItemInputAudioTranscriptionCompleted
  GetItemId() *string 
  SetContentIndex(v int32) *ConversationItemInputAudioTranscriptionCompleted
  GetContentIndex() *int32 
  SetTranscript(v string) *ConversationItemInputAudioTranscriptionCompleted
  GetTranscript() *string 
  SetLanguage(v string) *ConversationItemInputAudioTranscriptionCompleted
  GetLanguage() *string 
  SetEmotion(v string) *ConversationItemInputAudioTranscriptionCompleted
  GetEmotion() *string 
}

type ConversationItemInputAudioTranscriptionCompleted struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  ItemId *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
  ContentIndex *int32 `json:"content_index,omitempty" xml:"content_index,omitempty"`
  Transcript *string `json:"transcript,omitempty" xml:"transcript,omitempty"`
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
  Emotion *string `json:"emotion,omitempty" xml:"emotion,omitempty"`
}

func (s ConversationItemInputAudioTranscriptionCompleted) String() string {
  return dara.Prettify(s)
}

func (s ConversationItemInputAudioTranscriptionCompleted) GoString() string {
  return s.String()
}

func (s *ConversationItemInputAudioTranscriptionCompleted) GetType() *string  {
  return s.Type
}

func (s *ConversationItemInputAudioTranscriptionCompleted) GetEventId() *string  {
  return s.EventId
}

func (s *ConversationItemInputAudioTranscriptionCompleted) GetItemId() *string  {
  return s.ItemId
}

func (s *ConversationItemInputAudioTranscriptionCompleted) GetContentIndex() *int32  {
  return s.ContentIndex
}

func (s *ConversationItemInputAudioTranscriptionCompleted) GetTranscript() *string  {
  return s.Transcript
}

func (s *ConversationItemInputAudioTranscriptionCompleted) GetLanguage() *string  {
  return s.Language
}

func (s *ConversationItemInputAudioTranscriptionCompleted) GetEmotion() *string  {
  return s.Emotion
}

func (s *ConversationItemInputAudioTranscriptionCompleted) SetType(v string) *ConversationItemInputAudioTranscriptionCompleted {
  s.Type = &v
  return s
}

func (s *ConversationItemInputAudioTranscriptionCompleted) SetEventId(v string) *ConversationItemInputAudioTranscriptionCompleted {
  s.EventId = &v
  return s
}

func (s *ConversationItemInputAudioTranscriptionCompleted) SetItemId(v string) *ConversationItemInputAudioTranscriptionCompleted {
  s.ItemId = &v
  return s
}

func (s *ConversationItemInputAudioTranscriptionCompleted) SetContentIndex(v int32) *ConversationItemInputAudioTranscriptionCompleted {
  s.ContentIndex = &v
  return s
}

func (s *ConversationItemInputAudioTranscriptionCompleted) SetTranscript(v string) *ConversationItemInputAudioTranscriptionCompleted {
  s.Transcript = &v
  return s
}

func (s *ConversationItemInputAudioTranscriptionCompleted) SetLanguage(v string) *ConversationItemInputAudioTranscriptionCompleted {
  s.Language = &v
  return s
}

func (s *ConversationItemInputAudioTranscriptionCompleted) SetEmotion(v string) *ConversationItemInputAudioTranscriptionCompleted {
  s.Emotion = &v
  return s
}

func (s *ConversationItemInputAudioTranscriptionCompleted) Validate() error {
  return dara.Validate(s)
}

