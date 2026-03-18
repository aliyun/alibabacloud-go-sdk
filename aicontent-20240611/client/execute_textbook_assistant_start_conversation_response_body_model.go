// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTextbookAssistantStartConversationResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v *ExecuteTextbookAssistantStartConversationResponseBodyData) *ExecuteTextbookAssistantStartConversationResponseBody
  GetData() *ExecuteTextbookAssistantStartConversationResponseBodyData 
  SetErrCode(v string) *ExecuteTextbookAssistantStartConversationResponseBody
  GetErrCode() *string 
  SetErrMessage(v string) *ExecuteTextbookAssistantStartConversationResponseBody
  GetErrMessage() *string 
  SetHttpStatusCode(v int32) *ExecuteTextbookAssistantStartConversationResponseBody
  GetHttpStatusCode() *int32 
  SetRequestId(v string) *ExecuteTextbookAssistantStartConversationResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecuteTextbookAssistantStartConversationResponseBody
  GetSuccess() *bool 
}

type ExecuteTextbookAssistantStartConversationResponseBody struct {
  Data *ExecuteTextbookAssistantStartConversationResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
  // example:
  // 
  // B_USER_NOT_FOUND_EXCEPTION
  ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
  ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // 6F73C114-A76E-51AD-99E3-BC7B941B69E0
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // example:
  // 
  // true
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExecuteTextbookAssistantStartConversationResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantStartConversationResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantStartConversationResponseBody) GetData() *ExecuteTextbookAssistantStartConversationResponseBodyData  {
  return s.Data
}

func (s *ExecuteTextbookAssistantStartConversationResponseBody) GetErrCode() *string  {
  return s.ErrCode
}

func (s *ExecuteTextbookAssistantStartConversationResponseBody) GetErrMessage() *string  {
  return s.ErrMessage
}

func (s *ExecuteTextbookAssistantStartConversationResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExecuteTextbookAssistantStartConversationResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteTextbookAssistantStartConversationResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecuteTextbookAssistantStartConversationResponseBody) SetData(v *ExecuteTextbookAssistantStartConversationResponseBodyData) *ExecuteTextbookAssistantStartConversationResponseBody {
  s.Data = v
  return s
}

func (s *ExecuteTextbookAssistantStartConversationResponseBody) SetErrCode(v string) *ExecuteTextbookAssistantStartConversationResponseBody {
  s.ErrCode = &v
  return s
}

func (s *ExecuteTextbookAssistantStartConversationResponseBody) SetErrMessage(v string) *ExecuteTextbookAssistantStartConversationResponseBody {
  s.ErrMessage = &v
  return s
}

func (s *ExecuteTextbookAssistantStartConversationResponseBody) SetHttpStatusCode(v int32) *ExecuteTextbookAssistantStartConversationResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExecuteTextbookAssistantStartConversationResponseBody) SetRequestId(v string) *ExecuteTextbookAssistantStartConversationResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteTextbookAssistantStartConversationResponseBody) SetSuccess(v bool) *ExecuteTextbookAssistantStartConversationResponseBody {
  s.Success = &v
  return s
}

func (s *ExecuteTextbookAssistantStartConversationResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteTextbookAssistantStartConversationResponseBodyData struct {
  // example:
  // 
  // 6788e0b4b54c5268c1b78638
  Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
  // example:
  // 
  // 6788e0b475a4631ffc626722
  ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
  Result *ExecuteTextbookAssistantStartConversationResponseBodyDataResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
  // example:
  // 
  // 6788e0b45bdfc807f077a5a1
  User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ExecuteTextbookAssistantStartConversationResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantStartConversationResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantStartConversationResponseBodyData) GetAssistant() *string  {
  return s.Assistant
}

func (s *ExecuteTextbookAssistantStartConversationResponseBodyData) GetChatId() *string  {
  return s.ChatId
}

func (s *ExecuteTextbookAssistantStartConversationResponseBodyData) GetResult() *ExecuteTextbookAssistantStartConversationResponseBodyDataResult  {
  return s.Result
}

func (s *ExecuteTextbookAssistantStartConversationResponseBodyData) GetUser() *string  {
  return s.User
}

func (s *ExecuteTextbookAssistantStartConversationResponseBodyData) SetAssistant(v string) *ExecuteTextbookAssistantStartConversationResponseBodyData {
  s.Assistant = &v
  return s
}

func (s *ExecuteTextbookAssistantStartConversationResponseBodyData) SetChatId(v string) *ExecuteTextbookAssistantStartConversationResponseBodyData {
  s.ChatId = &v
  return s
}

func (s *ExecuteTextbookAssistantStartConversationResponseBodyData) SetResult(v *ExecuteTextbookAssistantStartConversationResponseBodyDataResult) *ExecuteTextbookAssistantStartConversationResponseBodyData {
  s.Result = v
  return s
}

func (s *ExecuteTextbookAssistantStartConversationResponseBodyData) SetUser(v string) *ExecuteTextbookAssistantStartConversationResponseBodyData {
  s.User = &v
  return s
}

func (s *ExecuteTextbookAssistantStartConversationResponseBodyData) Validate() error {
  if s.Result != nil {
    if err := s.Result.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteTextbookAssistantStartConversationResponseBodyDataResult struct {
  ChineseResult *string `json:"chineseResult,omitempty" xml:"chineseResult,omitempty"`
  // example:
  // 
  // Good evening! From the book, how does Mike Black introduce himself?
  EnglishResult *string `json:"englishResult,omitempty" xml:"englishResult,omitempty"`
}

func (s ExecuteTextbookAssistantStartConversationResponseBodyDataResult) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantStartConversationResponseBodyDataResult) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantStartConversationResponseBodyDataResult) GetChineseResult() *string  {
  return s.ChineseResult
}

func (s *ExecuteTextbookAssistantStartConversationResponseBodyDataResult) GetEnglishResult() *string  {
  return s.EnglishResult
}

func (s *ExecuteTextbookAssistantStartConversationResponseBodyDataResult) SetChineseResult(v string) *ExecuteTextbookAssistantStartConversationResponseBodyDataResult {
  s.ChineseResult = &v
  return s
}

func (s *ExecuteTextbookAssistantStartConversationResponseBodyDataResult) SetEnglishResult(v string) *ExecuteTextbookAssistantStartConversationResponseBodyDataResult {
  s.EnglishResult = &v
  return s
}

func (s *ExecuteTextbookAssistantStartConversationResponseBodyDataResult) Validate() error {
  return dara.Validate(s)
}

