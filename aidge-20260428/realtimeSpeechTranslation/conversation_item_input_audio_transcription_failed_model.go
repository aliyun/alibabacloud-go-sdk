// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iConversationItemInputAudioTranscriptionFailed interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *ConversationItemInputAudioTranscriptionFailed
  GetType() *string 
  SetEventId(v string) *ConversationItemInputAudioTranscriptionFailed
  GetEventId() *string 
  SetItemId(v string) *ConversationItemInputAudioTranscriptionFailed
  GetItemId() *string 
  SetContentIndex(v int32) *ConversationItemInputAudioTranscriptionFailed
  GetContentIndex() *int32 
  SetError(v *ConversationItemInputAudioTranscriptionFailedError) *ConversationItemInputAudioTranscriptionFailed
  GetError() *ConversationItemInputAudioTranscriptionFailedError 
}

type ConversationItemInputAudioTranscriptionFailed struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  ItemId *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
  ContentIndex *int32 `json:"content_index,omitempty" xml:"content_index,omitempty"`
  Error *ConversationItemInputAudioTranscriptionFailedError `json:"error,omitempty" xml:"error,omitempty" type:"Struct"`
}

func (s ConversationItemInputAudioTranscriptionFailed) String() string {
  return dara.Prettify(s)
}

func (s ConversationItemInputAudioTranscriptionFailed) GoString() string {
  return s.String()
}

func (s *ConversationItemInputAudioTranscriptionFailed) GetType() *string  {
  return s.Type
}

func (s *ConversationItemInputAudioTranscriptionFailed) GetEventId() *string  {
  return s.EventId
}

func (s *ConversationItemInputAudioTranscriptionFailed) GetItemId() *string  {
  return s.ItemId
}

func (s *ConversationItemInputAudioTranscriptionFailed) GetContentIndex() *int32  {
  return s.ContentIndex
}

func (s *ConversationItemInputAudioTranscriptionFailed) GetError() *ConversationItemInputAudioTranscriptionFailedError  {
  return s.Error
}

func (s *ConversationItemInputAudioTranscriptionFailed) SetType(v string) *ConversationItemInputAudioTranscriptionFailed {
  s.Type = &v
  return s
}

func (s *ConversationItemInputAudioTranscriptionFailed) SetEventId(v string) *ConversationItemInputAudioTranscriptionFailed {
  s.EventId = &v
  return s
}

func (s *ConversationItemInputAudioTranscriptionFailed) SetItemId(v string) *ConversationItemInputAudioTranscriptionFailed {
  s.ItemId = &v
  return s
}

func (s *ConversationItemInputAudioTranscriptionFailed) SetContentIndex(v int32) *ConversationItemInputAudioTranscriptionFailed {
  s.ContentIndex = &v
  return s
}

func (s *ConversationItemInputAudioTranscriptionFailed) SetError(v *ConversationItemInputAudioTranscriptionFailedError) *ConversationItemInputAudioTranscriptionFailed {
  s.Error = v
  return s
}

func (s *ConversationItemInputAudioTranscriptionFailed) Validate() error {
  if s.Error != nil {
    if err := s.Error.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ConversationItemInputAudioTranscriptionFailedError struct {
  Code *string `json:"code,omitempty" xml:"code,omitempty"`
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  Param *string `json:"param,omitempty" xml:"param,omitempty"`
}

func (s ConversationItemInputAudioTranscriptionFailedError) String() string {
  return dara.Prettify(s)
}

func (s ConversationItemInputAudioTranscriptionFailedError) GoString() string {
  return s.String()
}

func (s *ConversationItemInputAudioTranscriptionFailedError) GetCode() *string  {
  return s.Code
}

func (s *ConversationItemInputAudioTranscriptionFailedError) GetMessage() *string  {
  return s.Message
}

func (s *ConversationItemInputAudioTranscriptionFailedError) GetParam() *string  {
  return s.Param
}

func (s *ConversationItemInputAudioTranscriptionFailedError) SetCode(v string) *ConversationItemInputAudioTranscriptionFailedError {
  s.Code = &v
  return s
}

func (s *ConversationItemInputAudioTranscriptionFailedError) SetMessage(v string) *ConversationItemInputAudioTranscriptionFailedError {
  s.Message = &v
  return s
}

func (s *ConversationItemInputAudioTranscriptionFailedError) SetParam(v string) *ConversationItemInputAudioTranscriptionFailedError {
  s.Param = &v
  return s
}

func (s *ConversationItemInputAudioTranscriptionFailedError) Validate() error {
  return dara.Validate(s)
}

