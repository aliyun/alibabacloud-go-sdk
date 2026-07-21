// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAITeacherExpansionDialogueResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v *ExecuteAITeacherExpansionDialogueResponseBodyData) *ExecuteAITeacherExpansionDialogueResponseBody
  GetData() *ExecuteAITeacherExpansionDialogueResponseBodyData 
  SetErrCode(v string) *ExecuteAITeacherExpansionDialogueResponseBody
  GetErrCode() *string 
  SetErrMessage(v string) *ExecuteAITeacherExpansionDialogueResponseBody
  GetErrMessage() *string 
  SetHttpStatusCode(v int32) *ExecuteAITeacherExpansionDialogueResponseBody
  GetHttpStatusCode() *int32 
  SetRequestId(v string) *ExecuteAITeacherExpansionDialogueResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecuteAITeacherExpansionDialogueResponseBody
  GetSuccess() *bool 
}

type ExecuteAITeacherExpansionDialogueResponseBody struct {
  // The response data.
  // 
  // example:
  // 
  // []
  Data *ExecuteAITeacherExpansionDialogueResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
  // The error code returned if the request fails.
  // 
  // example:
  // 
  // UNKNOWN_ERROR
  ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
  // The error message returned if the request fails.
  // 
  // example:
  // 
  // 未知错误
  ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
  // The HTTP status code.
  // 
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
  // A unique identifier for the request.
  // 
  // example:
  // 
  // xxxx-xxxx-xxxx-xxxxxxxx
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // Indicates whether the request was successful.
  // 
  // example:
  // 
  // true
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueResponseBody) GetData() *ExecuteAITeacherExpansionDialogueResponseBodyData  {
  return s.Data
}

func (s *ExecuteAITeacherExpansionDialogueResponseBody) GetErrCode() *string  {
  return s.ErrCode
}

func (s *ExecuteAITeacherExpansionDialogueResponseBody) GetErrMessage() *string  {
  return s.ErrMessage
}

func (s *ExecuteAITeacherExpansionDialogueResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExecuteAITeacherExpansionDialogueResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteAITeacherExpansionDialogueResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecuteAITeacherExpansionDialogueResponseBody) SetData(v *ExecuteAITeacherExpansionDialogueResponseBodyData) *ExecuteAITeacherExpansionDialogueResponseBody {
  s.Data = v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueResponseBody) SetErrCode(v string) *ExecuteAITeacherExpansionDialogueResponseBody {
  s.ErrCode = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueResponseBody) SetErrMessage(v string) *ExecuteAITeacherExpansionDialogueResponseBody {
  s.ErrMessage = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueResponseBody) SetHttpStatusCode(v int32) *ExecuteAITeacherExpansionDialogueResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueResponseBody) SetRequestId(v string) *ExecuteAITeacherExpansionDialogueResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueResponseBody) SetSuccess(v bool) *ExecuteAITeacherExpansionDialogueResponseBody {
  s.Success = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteAITeacherExpansionDialogueResponseBodyData struct {
  // The AI teacher\\"s response in Chinese.
  // 
  // example:
  // 
  // 1
  ChineseResult *string `json:"chineseResult,omitempty" xml:"chineseResult,omitempty"`
  // The AI teacher\\"s response in English.
  // 
  // example:
  // 
  // 1
  EnglishResult *string `json:"englishResult,omitempty" xml:"englishResult,omitempty"`
  // Indicates whether the dialogue is finished.
  // 
  // example:
  // 
  // true
  IsFinish *bool `json:"isFinish,omitempty" xml:"isFinish,omitempty"`
  // Indicates whether the response is off-topic.
  // 
  // example:
  // 
  // true
  IsOffTopicControl *bool `json:"isOffTopicControl,omitempty" xml:"isOffTopicControl,omitempty"`
  // Indicates whether the response is on-topic.
  // 
  // example:
  // 
  // true
  IsOnTopic *bool `json:"isOnTopic,omitempty" xml:"isOnTopic,omitempty"`
  // The index of the current question.
  // 
  // example:
  // 
  // 2
  QuestionIndex *int32 `json:"questionIndex,omitempty" xml:"questionIndex,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueResponseBodyData) GetChineseResult() *string  {
  return s.ChineseResult
}

func (s *ExecuteAITeacherExpansionDialogueResponseBodyData) GetEnglishResult() *string  {
  return s.EnglishResult
}

func (s *ExecuteAITeacherExpansionDialogueResponseBodyData) GetIsFinish() *bool  {
  return s.IsFinish
}

func (s *ExecuteAITeacherExpansionDialogueResponseBodyData) GetIsOffTopicControl() *bool  {
  return s.IsOffTopicControl
}

func (s *ExecuteAITeacherExpansionDialogueResponseBodyData) GetIsOnTopic() *bool  {
  return s.IsOnTopic
}

func (s *ExecuteAITeacherExpansionDialogueResponseBodyData) GetQuestionIndex() *int32  {
  return s.QuestionIndex
}

func (s *ExecuteAITeacherExpansionDialogueResponseBodyData) SetChineseResult(v string) *ExecuteAITeacherExpansionDialogueResponseBodyData {
  s.ChineseResult = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueResponseBodyData) SetEnglishResult(v string) *ExecuteAITeacherExpansionDialogueResponseBodyData {
  s.EnglishResult = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueResponseBodyData) SetIsFinish(v bool) *ExecuteAITeacherExpansionDialogueResponseBodyData {
  s.IsFinish = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueResponseBodyData) SetIsOffTopicControl(v bool) *ExecuteAITeacherExpansionDialogueResponseBodyData {
  s.IsOffTopicControl = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueResponseBodyData) SetIsOnTopic(v bool) *ExecuteAITeacherExpansionDialogueResponseBodyData {
  s.IsOnTopic = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueResponseBodyData) SetQuestionIndex(v int32) *ExecuteAITeacherExpansionDialogueResponseBodyData {
  s.QuestionIndex = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueResponseBodyData) Validate() error {
  return dara.Validate(s)
}

