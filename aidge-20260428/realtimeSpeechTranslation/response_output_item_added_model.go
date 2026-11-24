// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iResponseOutputItemAdded interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *ResponseOutputItemAdded
  GetType() *string 
  SetEventId(v string) *ResponseOutputItemAdded
  GetEventId() *string 
  SetResponseId(v string) *ResponseOutputItemAdded
  GetResponseId() *string 
  SetOutputIndex(v int32) *ResponseOutputItemAdded
  GetOutputIndex() *int32 
  SetItem(v *ResponseOutputItemAddedItem) *ResponseOutputItemAdded
  GetItem() *ResponseOutputItemAddedItem 
}

type ResponseOutputItemAdded struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  ResponseId *string `json:"response_id,omitempty" xml:"response_id,omitempty"`
  OutputIndex *int32 `json:"output_index,omitempty" xml:"output_index,omitempty"`
  Item *ResponseOutputItemAddedItem `json:"item,omitempty" xml:"item,omitempty" type:"Struct"`
}

func (s ResponseOutputItemAdded) String() string {
  return dara.Prettify(s)
}

func (s ResponseOutputItemAdded) GoString() string {
  return s.String()
}

func (s *ResponseOutputItemAdded) GetType() *string  {
  return s.Type
}

func (s *ResponseOutputItemAdded) GetEventId() *string  {
  return s.EventId
}

func (s *ResponseOutputItemAdded) GetResponseId() *string  {
  return s.ResponseId
}

func (s *ResponseOutputItemAdded) GetOutputIndex() *int32  {
  return s.OutputIndex
}

func (s *ResponseOutputItemAdded) GetItem() *ResponseOutputItemAddedItem  {
  return s.Item
}

func (s *ResponseOutputItemAdded) SetType(v string) *ResponseOutputItemAdded {
  s.Type = &v
  return s
}

func (s *ResponseOutputItemAdded) SetEventId(v string) *ResponseOutputItemAdded {
  s.EventId = &v
  return s
}

func (s *ResponseOutputItemAdded) SetResponseId(v string) *ResponseOutputItemAdded {
  s.ResponseId = &v
  return s
}

func (s *ResponseOutputItemAdded) SetOutputIndex(v int32) *ResponseOutputItemAdded {
  s.OutputIndex = &v
  return s
}

func (s *ResponseOutputItemAdded) SetItem(v *ResponseOutputItemAddedItem) *ResponseOutputItemAdded {
  s.Item = v
  return s
}

func (s *ResponseOutputItemAdded) Validate() error {
  if s.Item != nil {
    if err := s.Item.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ResponseOutputItemAddedItem struct {
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  Object *string `json:"object,omitempty" xml:"object,omitempty"`
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  Role *string `json:"role,omitempty" xml:"role,omitempty"`
  Content []*ResponseOutputItemAddedItemContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s ResponseOutputItemAddedItem) String() string {
  return dara.Prettify(s)
}

func (s ResponseOutputItemAddedItem) GoString() string {
  return s.String()
}

func (s *ResponseOutputItemAddedItem) GetId() *string  {
  return s.Id
}

func (s *ResponseOutputItemAddedItem) GetObject() *string  {
  return s.Object
}

func (s *ResponseOutputItemAddedItem) GetType() *string  {
  return s.Type
}

func (s *ResponseOutputItemAddedItem) GetStatus() *string  {
  return s.Status
}

func (s *ResponseOutputItemAddedItem) GetRole() *string  {
  return s.Role
}

func (s *ResponseOutputItemAddedItem) GetContent() []*ResponseOutputItemAddedItemContent  {
  return s.Content
}

func (s *ResponseOutputItemAddedItem) SetId(v string) *ResponseOutputItemAddedItem {
  s.Id = &v
  return s
}

func (s *ResponseOutputItemAddedItem) SetObject(v string) *ResponseOutputItemAddedItem {
  s.Object = &v
  return s
}

func (s *ResponseOutputItemAddedItem) SetType(v string) *ResponseOutputItemAddedItem {
  s.Type = &v
  return s
}

func (s *ResponseOutputItemAddedItem) SetStatus(v string) *ResponseOutputItemAddedItem {
  s.Status = &v
  return s
}

func (s *ResponseOutputItemAddedItem) SetRole(v string) *ResponseOutputItemAddedItem {
  s.Role = &v
  return s
}

func (s *ResponseOutputItemAddedItem) SetContent(v []*ResponseOutputItemAddedItemContent) *ResponseOutputItemAddedItem {
  s.Content = v
  return s
}

func (s *ResponseOutputItemAddedItem) Validate() error {
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

type ResponseOutputItemAddedItemContent struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  Text *string `json:"text,omitempty" xml:"text,omitempty"`
  Transcript *string `json:"transcript,omitempty" xml:"transcript,omitempty"`
}

func (s ResponseOutputItemAddedItemContent) String() string {
  return dara.Prettify(s)
}

func (s ResponseOutputItemAddedItemContent) GoString() string {
  return s.String()
}

func (s *ResponseOutputItemAddedItemContent) GetType() *string  {
  return s.Type
}

func (s *ResponseOutputItemAddedItemContent) GetText() *string  {
  return s.Text
}

func (s *ResponseOutputItemAddedItemContent) GetTranscript() *string  {
  return s.Transcript
}

func (s *ResponseOutputItemAddedItemContent) SetType(v string) *ResponseOutputItemAddedItemContent {
  s.Type = &v
  return s
}

func (s *ResponseOutputItemAddedItemContent) SetText(v string) *ResponseOutputItemAddedItemContent {
  s.Text = &v
  return s
}

func (s *ResponseOutputItemAddedItemContent) SetTranscript(v string) *ResponseOutputItemAddedItemContent {
  s.Transcript = &v
  return s
}

func (s *ResponseOutputItemAddedItemContent) Validate() error {
  return dara.Validate(s)
}

