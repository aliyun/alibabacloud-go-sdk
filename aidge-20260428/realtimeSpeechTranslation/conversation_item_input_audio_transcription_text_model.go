// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iConversationItemInputAudioTranscriptionText interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *ConversationItemInputAudioTranscriptionText
  GetType() *string 
  SetEventId(v string) *ConversationItemInputAudioTranscriptionText
  GetEventId() *string 
  SetItemId(v string) *ConversationItemInputAudioTranscriptionText
  GetItemId() *string 
  SetContentIndex(v int32) *ConversationItemInputAudioTranscriptionText
  GetContentIndex() *int32 
  SetText(v string) *ConversationItemInputAudioTranscriptionText
  GetText() *string 
  SetStash(v string) *ConversationItemInputAudioTranscriptionText
  GetStash() *string 
  SetLanguage(v string) *ConversationItemInputAudioTranscriptionText
  GetLanguage() *string 
  SetEmotion(v string) *ConversationItemInputAudioTranscriptionText
  GetEmotion() *string 
}

type ConversationItemInputAudioTranscriptionText struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  ItemId *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
  ContentIndex *int32 `json:"content_index,omitempty" xml:"content_index,omitempty"`
  Text *string `json:"text,omitempty" xml:"text,omitempty"`
  Stash *string `json:"stash,omitempty" xml:"stash,omitempty"`
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
  Emotion *string `json:"emotion,omitempty" xml:"emotion,omitempty"`
}

func (s ConversationItemInputAudioTranscriptionText) String() string {
  return dara.Prettify(s)
}

func (s ConversationItemInputAudioTranscriptionText) GoString() string {
  return s.String()
}

func (s *ConversationItemInputAudioTranscriptionText) GetType() *string  {
  return s.Type
}

func (s *ConversationItemInputAudioTranscriptionText) GetEventId() *string  {
  return s.EventId
}

func (s *ConversationItemInputAudioTranscriptionText) GetItemId() *string  {
  return s.ItemId
}

func (s *ConversationItemInputAudioTranscriptionText) GetContentIndex() *int32  {
  return s.ContentIndex
}

func (s *ConversationItemInputAudioTranscriptionText) GetText() *string  {
  return s.Text
}

func (s *ConversationItemInputAudioTranscriptionText) GetStash() *string  {
  return s.Stash
}

func (s *ConversationItemInputAudioTranscriptionText) GetLanguage() *string  {
  return s.Language
}

func (s *ConversationItemInputAudioTranscriptionText) GetEmotion() *string  {
  return s.Emotion
}

func (s *ConversationItemInputAudioTranscriptionText) SetType(v string) *ConversationItemInputAudioTranscriptionText {
  s.Type = &v
  return s
}

func (s *ConversationItemInputAudioTranscriptionText) SetEventId(v string) *ConversationItemInputAudioTranscriptionText {
  s.EventId = &v
  return s
}

func (s *ConversationItemInputAudioTranscriptionText) SetItemId(v string) *ConversationItemInputAudioTranscriptionText {
  s.ItemId = &v
  return s
}

func (s *ConversationItemInputAudioTranscriptionText) SetContentIndex(v int32) *ConversationItemInputAudioTranscriptionText {
  s.ContentIndex = &v
  return s
}

func (s *ConversationItemInputAudioTranscriptionText) SetText(v string) *ConversationItemInputAudioTranscriptionText {
  s.Text = &v
  return s
}

func (s *ConversationItemInputAudioTranscriptionText) SetStash(v string) *ConversationItemInputAudioTranscriptionText {
  s.Stash = &v
  return s
}

func (s *ConversationItemInputAudioTranscriptionText) SetLanguage(v string) *ConversationItemInputAudioTranscriptionText {
  s.Language = &v
  return s
}

func (s *ConversationItemInputAudioTranscriptionText) SetEmotion(v string) *ConversationItemInputAudioTranscriptionText {
  s.Emotion = &v
  return s
}

func (s *ConversationItemInputAudioTranscriptionText) Validate() error {
  return dara.Validate(s)
}

