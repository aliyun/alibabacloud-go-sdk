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
  // The returned data.
  Data *ExecuteTextbookAssistantDialogueResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
  // The error code.
  // 
  // example:
  // 
  // null
  ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
  // The error message.
  // 
  // example:
  // 
  // null
  ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
  // The HTTP status code.
  // 
  // example:
  // 
  // 200
  HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // DBFA232A-1176-50E6-95AE-50F7A62A28AD
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // Indicates whether the request was successful.
  // 
  // example:
  // 
  // true
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
  // The ID of the Textbook Assistant\\"s message.
  // 
  // example:
  // 
  // 6788f4935bdfc807f077a984
  Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
  // The chat ID for this turn.
  // 
  // example:
  // 
  // 6788e0b475a4631ffc626722
  ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
  // The returned data.
  Result *ExecuteTextbookAssistantDialogueResponseBodyDataResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
  // The ID of the user\\"s reply.
  // 
  // example:
  // 
  // 6788f4905bdfc807f077a982
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
  // The Textbook Assistant\\"s reply in Chinese.
  // 
  // example:
  // 
  // 让我们再看一遍课文。迈克说: “我是迈克·布莱克。”你能试着像迈克那样说吗？
  ChineseResult *string `json:"chineseResult,omitempty" xml:"chineseResult,omitempty"`
  // The Textbook Assistant\\"s reply in English.
  // 
  // example:
  // 
  // Let\\"s look at the text again. Mike says, \\"I\\"m Mike Black.\\" Can you try saying it like Mike?
  EnglishResult *string `json:"englishResult,omitempty" xml:"englishResult,omitempty"`
  // Indicates whether the dialogue is finished.
  // 
  // example:
  // 
  // true
  IsFinish *bool `json:"isFinish,omitempty" xml:"isFinish,omitempty"`
  // Indicates whether the task is completed.
  // 
  // example:
  // 
  // true
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

