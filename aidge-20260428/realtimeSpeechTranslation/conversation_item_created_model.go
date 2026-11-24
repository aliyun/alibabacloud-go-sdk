// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iConversationItemCreated interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *ConversationItemCreated
  GetType() *string 
  SetEventId(v string) *ConversationItemCreated
  GetEventId() *string 
  SetPreviousItemId(v string) *ConversationItemCreated
  GetPreviousItemId() *string 
  SetItem(v *ConversationItemCreatedItem) *ConversationItemCreated
  GetItem() *ConversationItemCreatedItem 
}

type ConversationItemCreated struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  PreviousItemId *string `json:"previous_item_id,omitempty" xml:"previous_item_id,omitempty"`
  Item *ConversationItemCreatedItem `json:"item,omitempty" xml:"item,omitempty" type:"Struct"`
}

func (s ConversationItemCreated) String() string {
  return dara.Prettify(s)
}

func (s ConversationItemCreated) GoString() string {
  return s.String()
}

func (s *ConversationItemCreated) GetType() *string  {
  return s.Type
}

func (s *ConversationItemCreated) GetEventId() *string  {
  return s.EventId
}

func (s *ConversationItemCreated) GetPreviousItemId() *string  {
  return s.PreviousItemId
}

func (s *ConversationItemCreated) GetItem() *ConversationItemCreatedItem  {
  return s.Item
}

func (s *ConversationItemCreated) SetType(v string) *ConversationItemCreated {
  s.Type = &v
  return s
}

func (s *ConversationItemCreated) SetEventId(v string) *ConversationItemCreated {
  s.EventId = &v
  return s
}

func (s *ConversationItemCreated) SetPreviousItemId(v string) *ConversationItemCreated {
  s.PreviousItemId = &v
  return s
}

func (s *ConversationItemCreated) SetItem(v *ConversationItemCreatedItem) *ConversationItemCreated {
  s.Item = v
  return s
}

func (s *ConversationItemCreated) Validate() error {
  if s.Item != nil {
    if err := s.Item.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ConversationItemCreatedItem struct {
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  Object *string `json:"object,omitempty" xml:"object,omitempty"`
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  Role *string `json:"role,omitempty" xml:"role,omitempty"`
  Content []*ConversationItemCreatedItemContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s ConversationItemCreatedItem) String() string {
  return dara.Prettify(s)
}

func (s ConversationItemCreatedItem) GoString() string {
  return s.String()
}

func (s *ConversationItemCreatedItem) GetId() *string  {
  return s.Id
}

func (s *ConversationItemCreatedItem) GetObject() *string  {
  return s.Object
}

func (s *ConversationItemCreatedItem) GetType() *string  {
  return s.Type
}

func (s *ConversationItemCreatedItem) GetStatus() *string  {
  return s.Status
}

func (s *ConversationItemCreatedItem) GetRole() *string  {
  return s.Role
}

func (s *ConversationItemCreatedItem) GetContent() []*ConversationItemCreatedItemContent  {
  return s.Content
}

func (s *ConversationItemCreatedItem) SetId(v string) *ConversationItemCreatedItem {
  s.Id = &v
  return s
}

func (s *ConversationItemCreatedItem) SetObject(v string) *ConversationItemCreatedItem {
  s.Object = &v
  return s
}

func (s *ConversationItemCreatedItem) SetType(v string) *ConversationItemCreatedItem {
  s.Type = &v
  return s
}

func (s *ConversationItemCreatedItem) SetStatus(v string) *ConversationItemCreatedItem {
  s.Status = &v
  return s
}

func (s *ConversationItemCreatedItem) SetRole(v string) *ConversationItemCreatedItem {
  s.Role = &v
  return s
}

func (s *ConversationItemCreatedItem) SetContent(v []*ConversationItemCreatedItemContent) *ConversationItemCreatedItem {
  s.Content = v
  return s
}

func (s *ConversationItemCreatedItem) Validate() error {
  if s.Content != nil {
    for _, item := range s.Content {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type ConversationItemCreatedItemContent struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  Text *string `json:"text,omitempty" xml:"text,omitempty"`
  Transcript *string `json:"transcript,omitempty" xml:"transcript,omitempty"`
}

func (s ConversationItemCreatedItemContent) String() string {
  return dara.Prettify(s)
}

func (s ConversationItemCreatedItemContent) GoString() string {
  return s.String()
}

func (s *ConversationItemCreatedItemContent) GetType() *string  {
  return s.Type
}

func (s *ConversationItemCreatedItemContent) GetText() *string  {
  return s.Text
}

func (s *ConversationItemCreatedItemContent) GetTranscript() *string  {
  return s.Transcript
}

func (s *ConversationItemCreatedItemContent) SetType(v string) *ConversationItemCreatedItemContent {
  s.Type = &v
  return s
}

func (s *ConversationItemCreatedItemContent) SetText(v string) *ConversationItemCreatedItemContent {
  s.Text = &v
  return s
}

func (s *ConversationItemCreatedItemContent) SetTranscript(v string) *ConversationItemCreatedItemContent {
  s.Transcript = &v
  return s
}

func (s *ConversationItemCreatedItemContent) Validate() error {
  return dara.Validate(s)
}

