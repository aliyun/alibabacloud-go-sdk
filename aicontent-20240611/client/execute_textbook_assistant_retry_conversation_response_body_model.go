// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTextbookAssistantRetryConversationResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v *ExecuteTextbookAssistantRetryConversationResponseBodyData) *ExecuteTextbookAssistantRetryConversationResponseBody
  GetData() *ExecuteTextbookAssistantRetryConversationResponseBodyData 
  SetErrCode(v string) *ExecuteTextbookAssistantRetryConversationResponseBody
  GetErrCode() *string 
  SetErrMessage(v string) *ExecuteTextbookAssistantRetryConversationResponseBody
  GetErrMessage() *string 
  SetHttpStatusCode(v int32) *ExecuteTextbookAssistantRetryConversationResponseBody
  GetHttpStatusCode() *int32 
  SetRequestId(v string) *ExecuteTextbookAssistantRetryConversationResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecuteTextbookAssistantRetryConversationResponseBody
  GetSuccess() *bool 
}

type ExecuteTextbookAssistantRetryConversationResponseBody struct {
  Data *ExecuteTextbookAssistantRetryConversationResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
  // example:
  // 
  // 0
  ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
  // example:
  // 
  // null
  ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // 2F2ABF4B-A4F6-5EC7-B287-7EF5B156F1ED
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // example:
  // 
  // true
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExecuteTextbookAssistantRetryConversationResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantRetryConversationResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBody) GetData() *ExecuteTextbookAssistantRetryConversationResponseBodyData  {
  return s.Data
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBody) GetErrCode() *string  {
  return s.ErrCode
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBody) GetErrMessage() *string  {
  return s.ErrMessage
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBody) SetData(v *ExecuteTextbookAssistantRetryConversationResponseBodyData) *ExecuteTextbookAssistantRetryConversationResponseBody {
  s.Data = v
  return s
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBody) SetErrCode(v string) *ExecuteTextbookAssistantRetryConversationResponseBody {
  s.ErrCode = &v
  return s
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBody) SetErrMessage(v string) *ExecuteTextbookAssistantRetryConversationResponseBody {
  s.ErrMessage = &v
  return s
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBody) SetHttpStatusCode(v int32) *ExecuteTextbookAssistantRetryConversationResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBody) SetRequestId(v string) *ExecuteTextbookAssistantRetryConversationResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBody) SetSuccess(v bool) *ExecuteTextbookAssistantRetryConversationResponseBody {
  s.Success = &v
  return s
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteTextbookAssistantRetryConversationResponseBodyData struct {
  // example:
  // 
  // 6788e0b4b54c5268c1b78638
  Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
  // example:
  // 
  // 6788e0b475a4631ffc626722
  ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
  Result *ExecuteTextbookAssistantRetryConversationResponseBodyDataResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
  // example:
  // 
  // 6788e0b45bdfc807f077a5a1
  User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ExecuteTextbookAssistantRetryConversationResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantRetryConversationResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBodyData) GetAssistant() *string  {
  return s.Assistant
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBodyData) GetChatId() *string  {
  return s.ChatId
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBodyData) GetResult() *ExecuteTextbookAssistantRetryConversationResponseBodyDataResult  {
  return s.Result
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBodyData) GetUser() *string  {
  return s.User
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBodyData) SetAssistant(v string) *ExecuteTextbookAssistantRetryConversationResponseBodyData {
  s.Assistant = &v
  return s
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBodyData) SetChatId(v string) *ExecuteTextbookAssistantRetryConversationResponseBodyData {
  s.ChatId = &v
  return s
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBodyData) SetResult(v *ExecuteTextbookAssistantRetryConversationResponseBodyDataResult) *ExecuteTextbookAssistantRetryConversationResponseBodyData {
  s.Result = v
  return s
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBodyData) SetUser(v string) *ExecuteTextbookAssistantRetryConversationResponseBodyData {
  s.User = &v
  return s
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBodyData) Validate() error {
  if s.Result != nil {
    if err := s.Result.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteTextbookAssistantRetryConversationResponseBodyDataResult struct {
  ChineseResult *string `json:"chineseResult,omitempty" xml:"chineseResult,omitempty"`
  // example:
  // 
  // Good evening! From the book, how does Mike Black introduce himself?
  EnglishResult *string `json:"englishResult,omitempty" xml:"englishResult,omitempty"`
}

func (s ExecuteTextbookAssistantRetryConversationResponseBodyDataResult) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantRetryConversationResponseBodyDataResult) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBodyDataResult) GetChineseResult() *string  {
  return s.ChineseResult
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBodyDataResult) GetEnglishResult() *string  {
  return s.EnglishResult
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBodyDataResult) SetChineseResult(v string) *ExecuteTextbookAssistantRetryConversationResponseBodyDataResult {
  s.ChineseResult = &v
  return s
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBodyDataResult) SetEnglishResult(v string) *ExecuteTextbookAssistantRetryConversationResponseBodyDataResult {
  s.EnglishResult = &v
  return s
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBodyDataResult) Validate() error {
  return dara.Validate(s)
}

