// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iResponseOutputItemDone interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *ResponseOutputItemDone
  GetType() *string 
  SetEventId(v string) *ResponseOutputItemDone
  GetEventId() *string 
  SetResponseId(v string) *ResponseOutputItemDone
  GetResponseId() *string 
  SetOutputIndex(v int32) *ResponseOutputItemDone
  GetOutputIndex() *int32 
  SetItem(v *ResponseOutputItemDoneItem) *ResponseOutputItemDone
  GetItem() *ResponseOutputItemDoneItem 
}

type ResponseOutputItemDone struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  ResponseId *string `json:"response_id,omitempty" xml:"response_id,omitempty"`
  OutputIndex *int32 `json:"output_index,omitempty" xml:"output_index,omitempty"`
  Item *ResponseOutputItemDoneItem `json:"item,omitempty" xml:"item,omitempty" type:"Struct"`
}

func (s ResponseOutputItemDone) String() string {
  return dara.Prettify(s)
}

func (s ResponseOutputItemDone) GoString() string {
  return s.String()
}

func (s *ResponseOutputItemDone) GetType() *string  {
  return s.Type
}

func (s *ResponseOutputItemDone) GetEventId() *string  {
  return s.EventId
}

func (s *ResponseOutputItemDone) GetResponseId() *string  {
  return s.ResponseId
}

func (s *ResponseOutputItemDone) GetOutputIndex() *int32  {
  return s.OutputIndex
}

func (s *ResponseOutputItemDone) GetItem() *ResponseOutputItemDoneItem  {
  return s.Item
}

func (s *ResponseOutputItemDone) SetType(v string) *ResponseOutputItemDone {
  s.Type = &v
  return s
}

func (s *ResponseOutputItemDone) SetEventId(v string) *ResponseOutputItemDone {
  s.EventId = &v
  return s
}

func (s *ResponseOutputItemDone) SetResponseId(v string) *ResponseOutputItemDone {
  s.ResponseId = &v
  return s
}

func (s *ResponseOutputItemDone) SetOutputIndex(v int32) *ResponseOutputItemDone {
  s.OutputIndex = &v
  return s
}

func (s *ResponseOutputItemDone) SetItem(v *ResponseOutputItemDoneItem) *ResponseOutputItemDone {
  s.Item = v
  return s
}

func (s *ResponseOutputItemDone) Validate() error {
  if s.Item != nil {
    if err := s.Item.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ResponseOutputItemDoneItem struct {
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  Object *string `json:"object,omitempty" xml:"object,omitempty"`
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  Role *string `json:"role,omitempty" xml:"role,omitempty"`
  Content []*ResponseOutputItemDoneItemContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s ResponseOutputItemDoneItem) String() string {
  return dara.Prettify(s)
}

func (s ResponseOutputItemDoneItem) GoString() string {
  return s.String()
}

func (s *ResponseOutputItemDoneItem) GetId() *string  {
  return s.Id
}

func (s *ResponseOutputItemDoneItem) GetObject() *string  {
  return s.Object
}

func (s *ResponseOutputItemDoneItem) GetType() *string  {
  return s.Type
}

func (s *ResponseOutputItemDoneItem) GetStatus() *string  {
  return s.Status
}

func (s *ResponseOutputItemDoneItem) GetRole() *string  {
  return s.Role
}

func (s *ResponseOutputItemDoneItem) GetContent() []*ResponseOutputItemDoneItemContent  {
  return s.Content
}

func (s *ResponseOutputItemDoneItem) SetId(v string) *ResponseOutputItemDoneItem {
  s.Id = &v
  return s
}

func (s *ResponseOutputItemDoneItem) SetObject(v string) *ResponseOutputItemDoneItem {
  s.Object = &v
  return s
}

func (s *ResponseOutputItemDoneItem) SetType(v string) *ResponseOutputItemDoneItem {
  s.Type = &v
  return s
}

func (s *ResponseOutputItemDoneItem) SetStatus(v string) *ResponseOutputItemDoneItem {
  s.Status = &v
  return s
}

func (s *ResponseOutputItemDoneItem) SetRole(v string) *ResponseOutputItemDoneItem {
  s.Role = &v
  return s
}

func (s *ResponseOutputItemDoneItem) SetContent(v []*ResponseOutputItemDoneItemContent) *ResponseOutputItemDoneItem {
  s.Content = v
  return s
}

func (s *ResponseOutputItemDoneItem) Validate() error {
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

type ResponseOutputItemDoneItemContent struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  Text *string `json:"text,omitempty" xml:"text,omitempty"`
  Transcript *string `json:"transcript,omitempty" xml:"transcript,omitempty"`
}

func (s ResponseOutputItemDoneItemContent) String() string {
  return dara.Prettify(s)
}

func (s ResponseOutputItemDoneItemContent) GoString() string {
  return s.String()
}

func (s *ResponseOutputItemDoneItemContent) GetType() *string  {
  return s.Type
}

func (s *ResponseOutputItemDoneItemContent) GetText() *string  {
  return s.Text
}

func (s *ResponseOutputItemDoneItemContent) GetTranscript() *string  {
  return s.Transcript
}

func (s *ResponseOutputItemDoneItemContent) SetType(v string) *ResponseOutputItemDoneItemContent {
  s.Type = &v
  return s
}

func (s *ResponseOutputItemDoneItemContent) SetText(v string) *ResponseOutputItemDoneItemContent {
  s.Text = &v
  return s
}

func (s *ResponseOutputItemDoneItemContent) SetTranscript(v string) *ResponseOutputItemDoneItemContent {
  s.Transcript = &v
  return s
}

func (s *ResponseOutputItemDoneItemContent) Validate() error {
  return dara.Validate(s)
}

