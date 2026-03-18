// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTextbookAssistantSseDialogueResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetAssistant(v string) *ExecuteTextbookAssistantSseDialogueResponseBody
  GetAssistant() *string 
  SetChatId(v string) *ExecuteTextbookAssistantSseDialogueResponseBody
  GetChatId() *string 
  SetData(v *ExecuteTextbookAssistantSseDialogueResponseBodyData) *ExecuteTextbookAssistantSseDialogueResponseBody
  GetData() *ExecuteTextbookAssistantSseDialogueResponseBodyData 
  SetErrCode(v string) *ExecuteTextbookAssistantSseDialogueResponseBody
  GetErrCode() *string 
  SetErrMessage(v string) *ExecuteTextbookAssistantSseDialogueResponseBody
  GetErrMessage() *string 
  SetRequestId(v string) *ExecuteTextbookAssistantSseDialogueResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecuteTextbookAssistantSseDialogueResponseBody
  GetSuccess() *bool 
  SetUser(v string) *ExecuteTextbookAssistantSseDialogueResponseBody
  GetUser() *string 
}

type ExecuteTextbookAssistantSseDialogueResponseBody struct {
  // example:
  // 
  // 67e4c9d95bdfc83cd742ae7c
  Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
  // example:
  // 
  // 67e374acb54c526c95c4fbd4
  ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
  Data *ExecuteTextbookAssistantSseDialogueResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
  // example:
  // 
  // BIZ_ERROR
  ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
  ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // xxxx-xxxx-xxxx-xxxxxxxx
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // example:
  // 
  // true
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
  // example:
  // 
  // 67e4c9d6b54c526c95c53925
  User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ExecuteTextbookAssistantSseDialogueResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantSseDialogueResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantSseDialogueResponseBody) GetAssistant() *string  {
  return s.Assistant
}

func (s *ExecuteTextbookAssistantSseDialogueResponseBody) GetChatId() *string  {
  return s.ChatId
}

func (s *ExecuteTextbookAssistantSseDialogueResponseBody) GetData() *ExecuteTextbookAssistantSseDialogueResponseBodyData  {
  return s.Data
}

func (s *ExecuteTextbookAssistantSseDialogueResponseBody) GetErrCode() *string  {
  return s.ErrCode
}

func (s *ExecuteTextbookAssistantSseDialogueResponseBody) GetErrMessage() *string  {
  return s.ErrMessage
}

func (s *ExecuteTextbookAssistantSseDialogueResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteTextbookAssistantSseDialogueResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecuteTextbookAssistantSseDialogueResponseBody) GetUser() *string  {
  return s.User
}

func (s *ExecuteTextbookAssistantSseDialogueResponseBody) SetAssistant(v string) *ExecuteTextbookAssistantSseDialogueResponseBody {
  s.Assistant = &v
  return s
}

func (s *ExecuteTextbookAssistantSseDialogueResponseBody) SetChatId(v string) *ExecuteTextbookAssistantSseDialogueResponseBody {
  s.ChatId = &v
  return s
}

func (s *ExecuteTextbookAssistantSseDialogueResponseBody) SetData(v *ExecuteTextbookAssistantSseDialogueResponseBodyData) *ExecuteTextbookAssistantSseDialogueResponseBody {
  s.Data = v
  return s
}

func (s *ExecuteTextbookAssistantSseDialogueResponseBody) SetErrCode(v string) *ExecuteTextbookAssistantSseDialogueResponseBody {
  s.ErrCode = &v
  return s
}

func (s *ExecuteTextbookAssistantSseDialogueResponseBody) SetErrMessage(v string) *ExecuteTextbookAssistantSseDialogueResponseBody {
  s.ErrMessage = &v
  return s
}

func (s *ExecuteTextbookAssistantSseDialogueResponseBody) SetRequestId(v string) *ExecuteTextbookAssistantSseDialogueResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteTextbookAssistantSseDialogueResponseBody) SetSuccess(v bool) *ExecuteTextbookAssistantSseDialogueResponseBody {
  s.Success = &v
  return s
}

func (s *ExecuteTextbookAssistantSseDialogueResponseBody) SetUser(v string) *ExecuteTextbookAssistantSseDialogueResponseBody {
  s.User = &v
  return s
}

func (s *ExecuteTextbookAssistantSseDialogueResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteTextbookAssistantSseDialogueResponseBodyData struct {
  // example:
  // 
  // Thanks, Lily. Do you like meat, Lily?
  EnglishResult *string `json:"englishResult,omitempty" xml:"englishResult,omitempty"`
  // example:
  // 
  // true
  IsFinish *bool `json:"isFinish,omitempty" xml:"isFinish,omitempty"`
}

func (s ExecuteTextbookAssistantSseDialogueResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantSseDialogueResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantSseDialogueResponseBodyData) GetEnglishResult() *string  {
  return s.EnglishResult
}

func (s *ExecuteTextbookAssistantSseDialogueResponseBodyData) GetIsFinish() *bool  {
  return s.IsFinish
}

func (s *ExecuteTextbookAssistantSseDialogueResponseBodyData) SetEnglishResult(v string) *ExecuteTextbookAssistantSseDialogueResponseBodyData {
  s.EnglishResult = &v
  return s
}

func (s *ExecuteTextbookAssistantSseDialogueResponseBodyData) SetIsFinish(v bool) *ExecuteTextbookAssistantSseDialogueResponseBodyData {
  s.IsFinish = &v
  return s
}

func (s *ExecuteTextbookAssistantSseDialogueResponseBodyData) Validate() error {
  return dara.Validate(s)
}

