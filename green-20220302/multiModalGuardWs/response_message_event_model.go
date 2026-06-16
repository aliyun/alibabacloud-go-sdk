// This file is auto-generated, don't edit it. Thanks.
package multiModalGuardWs

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iResponseMessageEvent interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int32) *ResponseMessageEvent
  GetCode() *int32 
  SetMessage(v string) *ResponseMessageEvent
  GetMessage() *string 
  SetData(v *ResponseMessageEventData) *ResponseMessageEvent
  GetData() *ResponseMessageEventData 
}

type ResponseMessageEvent struct {
  Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  Data *ResponseMessageEventData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ResponseMessageEvent) String() string {
  return dara.Prettify(s)
}

func (s ResponseMessageEvent) GoString() string {
  return s.String()
}

func (s *ResponseMessageEvent) GetCode() *int32  {
  return s.Code
}

func (s *ResponseMessageEvent) GetMessage() *string  {
  return s.Message
}

func (s *ResponseMessageEvent) GetData() *ResponseMessageEventData  {
  return s.Data
}

func (s *ResponseMessageEvent) SetCode(v int32) *ResponseMessageEvent {
  s.Code = &v
  return s
}

func (s *ResponseMessageEvent) SetMessage(v string) *ResponseMessageEvent {
  s.Message = &v
  return s
}

func (s *ResponseMessageEvent) SetData(v *ResponseMessageEventData) *ResponseMessageEvent {
  s.Data = v
  return s
}

func (s *ResponseMessageEvent) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ResponseMessageEventData struct {
  Detail []*ResponseMessageEventDataDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
  Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
  DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
  SeqList []*string `json:"SeqList,omitempty" xml:"SeqList,omitempty" type:"Repeated"`
}

func (s ResponseMessageEventData) String() string {
  return dara.Prettify(s)
}

func (s ResponseMessageEventData) GoString() string {
  return s.String()
}

func (s *ResponseMessageEventData) GetDetail() []*ResponseMessageEventDataDetail  {
  return s.Detail
}

func (s *ResponseMessageEventData) GetSuggestion() *string  {
  return s.Suggestion
}

func (s *ResponseMessageEventData) GetDataId() *string  {
  return s.DataId
}

func (s *ResponseMessageEventData) GetSeqList() []*string  {
  return s.SeqList
}

func (s *ResponseMessageEventData) SetDetail(v []*ResponseMessageEventDataDetail) *ResponseMessageEventData {
  s.Detail = v
  return s
}

func (s *ResponseMessageEventData) SetSuggestion(v string) *ResponseMessageEventData {
  s.Suggestion = &v
  return s
}

func (s *ResponseMessageEventData) SetDataId(v string) *ResponseMessageEventData {
  s.DataId = &v
  return s
}

func (s *ResponseMessageEventData) SetSeqList(v []*string) *ResponseMessageEventData {
  s.SeqList = v
  return s
}

func (s *ResponseMessageEventData) Validate() error {
  if s.Detail != nil {
    for _, item := range s.Detail {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type ResponseMessageEventDataDetail struct {
  Result []*ResponseMessageEventDataDetailResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
  Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
  Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
  Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s ResponseMessageEventDataDetail) String() string {
  return dara.Prettify(s)
}

func (s ResponseMessageEventDataDetail) GoString() string {
  return s.String()
}

func (s *ResponseMessageEventDataDetail) GetResult() []*ResponseMessageEventDataDetailResult  {
  return s.Result
}

func (s *ResponseMessageEventDataDetail) GetType() *string  {
  return s.Type
}

func (s *ResponseMessageEventDataDetail) GetLevel() *string  {
  return s.Level
}

func (s *ResponseMessageEventDataDetail) GetSuggestion() *string  {
  return s.Suggestion
}

func (s *ResponseMessageEventDataDetail) SetResult(v []*ResponseMessageEventDataDetailResult) *ResponseMessageEventDataDetail {
  s.Result = v
  return s
}

func (s *ResponseMessageEventDataDetail) SetType(v string) *ResponseMessageEventDataDetail {
  s.Type = &v
  return s
}

func (s *ResponseMessageEventDataDetail) SetLevel(v string) *ResponseMessageEventDataDetail {
  s.Level = &v
  return s
}

func (s *ResponseMessageEventDataDetail) SetSuggestion(v string) *ResponseMessageEventDataDetail {
  s.Suggestion = &v
  return s
}

func (s *ResponseMessageEventDataDetail) Validate() error {
  if s.Result != nil {
    for _, item := range s.Result {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type ResponseMessageEventDataDetailResult struct {
  Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
  Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
  Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
  Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
  Ext interface{} `json:"Ext,omitempty" xml:"Ext,omitempty"`
}

func (s ResponseMessageEventDataDetailResult) String() string {
  return dara.Prettify(s)
}

func (s ResponseMessageEventDataDetailResult) GoString() string {
  return s.String()
}

func (s *ResponseMessageEventDataDetailResult) GetLabel() *string  {
  return s.Label
}

func (s *ResponseMessageEventDataDetailResult) GetDescription() *string  {
  return s.Description
}

func (s *ResponseMessageEventDataDetailResult) GetConfidence() *float32  {
  return s.Confidence
}

func (s *ResponseMessageEventDataDetailResult) GetLevel() *string  {
  return s.Level
}

func (s *ResponseMessageEventDataDetailResult) GetExt() interface{}  {
  return s.Ext
}

func (s *ResponseMessageEventDataDetailResult) SetLabel(v string) *ResponseMessageEventDataDetailResult {
  s.Label = &v
  return s
}

func (s *ResponseMessageEventDataDetailResult) SetDescription(v string) *ResponseMessageEventDataDetailResult {
  s.Description = &v
  return s
}

func (s *ResponseMessageEventDataDetailResult) SetConfidence(v float32) *ResponseMessageEventDataDetailResult {
  s.Confidence = &v
  return s
}

func (s *ResponseMessageEventDataDetailResult) SetLevel(v string) *ResponseMessageEventDataDetailResult {
  s.Level = &v
  return s
}

func (s *ResponseMessageEventDataDetailResult) SetExt(v interface{}) *ResponseMessageEventDataDetailResult {
  s.Ext = v
  return s
}

func (s *ResponseMessageEventDataDetailResult) Validate() error {
  return dara.Validate(s)
}

