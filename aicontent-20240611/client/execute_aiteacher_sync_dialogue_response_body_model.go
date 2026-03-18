// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAITeacherSyncDialogueResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v *ExecuteAITeacherSyncDialogueResponseBodyData) *ExecuteAITeacherSyncDialogueResponseBody
  GetData() *ExecuteAITeacherSyncDialogueResponseBodyData 
  SetErrCode(v string) *ExecuteAITeacherSyncDialogueResponseBody
  GetErrCode() *string 
  SetErrMessage(v string) *ExecuteAITeacherSyncDialogueResponseBody
  GetErrMessage() *string 
  SetHttpStatusCode(v int32) *ExecuteAITeacherSyncDialogueResponseBody
  GetHttpStatusCode() *int32 
  SetRequestId(v string) *ExecuteAITeacherSyncDialogueResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecuteAITeacherSyncDialogueResponseBody
  GetSuccess() *bool 
}

type ExecuteAITeacherSyncDialogueResponseBody struct {
  // example:
  // 
  // []
  Data *ExecuteAITeacherSyncDialogueResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
  // example:
  // 
  // UNKNOWN_ERROR
  ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
  // example:
  // 
  // 未知错误
  ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
  // example:
  // 
  // xxxx-xxxx-xxxx-xxxxxxxx
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // example:
  // 
  // true
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExecuteAITeacherSyncDialogueResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherSyncDialogueResponseBody) GetData() *ExecuteAITeacherSyncDialogueResponseBodyData  {
  return s.Data
}

func (s *ExecuteAITeacherSyncDialogueResponseBody) GetErrCode() *string  {
  return s.ErrCode
}

func (s *ExecuteAITeacherSyncDialogueResponseBody) GetErrMessage() *string  {
  return s.ErrMessage
}

func (s *ExecuteAITeacherSyncDialogueResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExecuteAITeacherSyncDialogueResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteAITeacherSyncDialogueResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecuteAITeacherSyncDialogueResponseBody) SetData(v *ExecuteAITeacherSyncDialogueResponseBodyData) *ExecuteAITeacherSyncDialogueResponseBody {
  s.Data = v
  return s
}

func (s *ExecuteAITeacherSyncDialogueResponseBody) SetErrCode(v string) *ExecuteAITeacherSyncDialogueResponseBody {
  s.ErrCode = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueResponseBody) SetErrMessage(v string) *ExecuteAITeacherSyncDialogueResponseBody {
  s.ErrMessage = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueResponseBody) SetHttpStatusCode(v int32) *ExecuteAITeacherSyncDialogueResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueResponseBody) SetRequestId(v string) *ExecuteAITeacherSyncDialogueResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueResponseBody) SetSuccess(v bool) *ExecuteAITeacherSyncDialogueResponseBody {
  s.Success = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteAITeacherSyncDialogueResponseBodyData struct {
  // example:
  // 
  // Thanks, Lily. Do you like meat, Lily?
  EnglishResult *string `json:"englishResult,omitempty" xml:"englishResult,omitempty"`
  // example:
  // 
  // true
  IsFinish *bool `json:"isFinish,omitempty" xml:"isFinish,omitempty"`
  // example:
  // 
  // true
  IsOnTopic *bool `json:"isOnTopic,omitempty" xml:"isOnTopic,omitempty"`
  // example:
  // 
  // 2
  QuestionIndex *int32 `json:"questionIndex,omitempty" xml:"questionIndex,omitempty"`
}

func (s ExecuteAITeacherSyncDialogueResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherSyncDialogueResponseBodyData) GetEnglishResult() *string  {
  return s.EnglishResult
}

func (s *ExecuteAITeacherSyncDialogueResponseBodyData) GetIsFinish() *bool  {
  return s.IsFinish
}

func (s *ExecuteAITeacherSyncDialogueResponseBodyData) GetIsOnTopic() *bool  {
  return s.IsOnTopic
}

func (s *ExecuteAITeacherSyncDialogueResponseBodyData) GetQuestionIndex() *int32  {
  return s.QuestionIndex
}

func (s *ExecuteAITeacherSyncDialogueResponseBodyData) SetEnglishResult(v string) *ExecuteAITeacherSyncDialogueResponseBodyData {
  s.EnglishResult = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueResponseBodyData) SetIsFinish(v bool) *ExecuteAITeacherSyncDialogueResponseBodyData {
  s.IsFinish = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueResponseBodyData) SetIsOnTopic(v bool) *ExecuteAITeacherSyncDialogueResponseBodyData {
  s.IsOnTopic = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueResponseBodyData) SetQuestionIndex(v int32) *ExecuteAITeacherSyncDialogueResponseBodyData {
  s.QuestionIndex = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueResponseBodyData) Validate() error {
  return dara.Validate(s)
}

