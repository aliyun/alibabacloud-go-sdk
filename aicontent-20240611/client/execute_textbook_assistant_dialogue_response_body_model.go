// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTextbookAssistantDialogueResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v *ExecuteTextbookAssistantDialogueResponseBodyData) *ExecuteTextbookAssistantDialogueResponseBody
  GetData() *ExecuteTextbookAssistantDialogueResponseBodyData 
  SetErrCode(v string) *ExecuteTextbookAssistantDialogueResponseBody
  GetErrCode() *string 
  SetErrMessage(v string) *ExecuteTextbookAssistantDialogueResponseBody
  GetErrMessage() *string 
  SetHttpStatusCode(v string) *ExecuteTextbookAssistantDialogueResponseBody
  GetHttpStatusCode() *string 
  SetRequestId(v string) *ExecuteTextbookAssistantDialogueResponseBody
  GetRequestId() *string 
  SetSuccess(v string) *ExecuteTextbookAssistantDialogueResponseBody
  GetSuccess() *string 
}

type ExecuteTextbookAssistantDialogueResponseBody struct {
  Data *ExecuteTextbookAssistantDialogueResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
  ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
  ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
  HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
  // Id of the request
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExecuteTextbookAssistantDialogueResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantDialogueResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantDialogueResponseBody) GetData() *ExecuteTextbookAssistantDialogueResponseBodyData  {
  return s.Data
}

func (s *ExecuteTextbookAssistantDialogueResponseBody) GetErrCode() *string  {
  return s.ErrCode
}

func (s *ExecuteTextbookAssistantDialogueResponseBody) GetErrMessage() *string  {
  return s.ErrMessage
}

func (s *ExecuteTextbookAssistantDialogueResponseBody) GetHttpStatusCode() *string  {
  return s.HttpStatusCode
}

func (s *ExecuteTextbookAssistantDialogueResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteTextbookAssistantDialogueResponseBody) GetSuccess() *string  {
  return s.Success
}

func (s *ExecuteTextbookAssistantDialogueResponseBody) SetData(v *ExecuteTextbookAssistantDialogueResponseBodyData) *ExecuteTextbookAssistantDialogueResponseBody {
  s.Data = v
  return s
}

func (s *ExecuteTextbookAssistantDialogueResponseBody) SetErrCode(v string) *ExecuteTextbookAssistantDialogueResponseBody {
  s.ErrCode = &v
  return s
}

func (s *ExecuteTextbookAssistantDialogueResponseBody) SetErrMessage(v string) *ExecuteTextbookAssistantDialogueResponseBody {
  s.ErrMessage = &v
  return s
}

func (s *ExecuteTextbookAssistantDialogueResponseBody) SetHttpStatusCode(v string) *ExecuteTextbookAssistantDialogueResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExecuteTextbookAssistantDialogueResponseBody) SetRequestId(v string) *ExecuteTextbookAssistantDialogueResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteTextbookAssistantDialogueResponseBody) SetSuccess(v string) *ExecuteTextbookAssistantDialogueResponseBody {
  s.Success = &v
  return s
}

func (s *ExecuteTextbookAssistantDialogueResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteTextbookAssistantDialogueResponseBodyData struct {
  Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
  ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
  Result *ExecuteTextbookAssistantDialogueResponseBodyDataResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
  User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ExecuteTextbookAssistantDialogueResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantDialogueResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantDialogueResponseBodyData) GetAssistant() *string  {
  return s.Assistant
}

func (s *ExecuteTextbookAssistantDialogueResponseBodyData) GetChatId() *string  {
  return s.ChatId
}

func (s *ExecuteTextbookAssistantDialogueResponseBodyData) GetResult() *ExecuteTextbookAssistantDialogueResponseBodyDataResult  {
  return s.Result
}

func (s *ExecuteTextbookAssistantDialogueResponseBodyData) GetUser() *string  {
  return s.User
}

func (s *ExecuteTextbookAssistantDialogueResponseBodyData) SetAssistant(v string) *ExecuteTextbookAssistantDialogueResponseBodyData {
  s.Assistant = &v
  return s
}

func (s *ExecuteTextbookAssistantDialogueResponseBodyData) SetChatId(v string) *ExecuteTextbookAssistantDialogueResponseBodyData {
  s.ChatId = &v
  return s
}

func (s *ExecuteTextbookAssistantDialogueResponseBodyData) SetResult(v *ExecuteTextbookAssistantDialogueResponseBodyDataResult) *ExecuteTextbookAssistantDialogueResponseBodyData {
  s.Result = v
  return s
}

func (s *ExecuteTextbookAssistantDialogueResponseBodyData) SetUser(v string) *ExecuteTextbookAssistantDialogueResponseBodyData {
  s.User = &v
  return s
}

func (s *ExecuteTextbookAssistantDialogueResponseBodyData) Validate() error {
  if s.Result != nil {
    if err := s.Result.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteTextbookAssistantDialogueResponseBodyDataResult struct {
  ChineseResult *string `json:"chineseResult,omitempty" xml:"chineseResult,omitempty"`
  EnglishResult *string `json:"englishResult,omitempty" xml:"englishResult,omitempty"`
  IsFinish *bool `json:"isFinish,omitempty" xml:"isFinish,omitempty"`
  IsTaskCompleted *bool `json:"isTaskCompleted,omitempty" xml:"isTaskCompleted,omitempty"`
}

func (s ExecuteTextbookAssistantDialogueResponseBodyDataResult) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantDialogueResponseBodyDataResult) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantDialogueResponseBodyDataResult) GetChineseResult() *string  {
  return s.ChineseResult
}

func (s *ExecuteTextbookAssistantDialogueResponseBodyDataResult) GetEnglishResult() *string  {
  return s.EnglishResult
}

func (s *ExecuteTextbookAssistantDialogueResponseBodyDataResult) GetIsFinish() *bool  {
  return s.IsFinish
}

func (s *ExecuteTextbookAssistantDialogueResponseBodyDataResult) GetIsTaskCompleted() *bool  {
  return s.IsTaskCompleted
}

func (s *ExecuteTextbookAssistantDialogueResponseBodyDataResult) SetChineseResult(v string) *ExecuteTextbookAssistantDialogueResponseBodyDataResult {
  s.ChineseResult = &v
  return s
}

func (s *ExecuteTextbookAssistantDialogueResponseBodyDataResult) SetEnglishResult(v string) *ExecuteTextbookAssistantDialogueResponseBodyDataResult {
  s.EnglishResult = &v
  return s
}

func (s *ExecuteTextbookAssistantDialogueResponseBodyDataResult) SetIsFinish(v bool) *ExecuteTextbookAssistantDialogueResponseBodyDataResult {
  s.IsFinish = &v
  return s
}

func (s *ExecuteTextbookAssistantDialogueResponseBodyDataResult) SetIsTaskCompleted(v bool) *ExecuteTextbookAssistantDialogueResponseBodyDataResult {
  s.IsTaskCompleted = &v
  return s
}

func (s *ExecuteTextbookAssistantDialogueResponseBodyDataResult) Validate() error {
  return dara.Validate(s)
}

