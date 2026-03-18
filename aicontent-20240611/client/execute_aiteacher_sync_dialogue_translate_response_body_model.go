// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAITeacherSyncDialogueTranslateResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v *ExecuteAITeacherSyncDialogueTranslateResponseBodyData) *ExecuteAITeacherSyncDialogueTranslateResponseBody
  GetData() *ExecuteAITeacherSyncDialogueTranslateResponseBodyData 
  SetErrCode(v string) *ExecuteAITeacherSyncDialogueTranslateResponseBody
  GetErrCode() *string 
  SetErrMessage(v string) *ExecuteAITeacherSyncDialogueTranslateResponseBody
  GetErrMessage() *string 
  SetHttpStatusCode(v int32) *ExecuteAITeacherSyncDialogueTranslateResponseBody
  GetHttpStatusCode() *int32 
  SetRequestId(v string) *ExecuteAITeacherSyncDialogueTranslateResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecuteAITeacherSyncDialogueTranslateResponseBody
  GetSuccess() *bool 
}

type ExecuteAITeacherSyncDialogueTranslateResponseBody struct {
  // example:
  // 
  // []
  Data *ExecuteAITeacherSyncDialogueTranslateResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s ExecuteAITeacherSyncDialogueTranslateResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueTranslateResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponseBody) GetData() *ExecuteAITeacherSyncDialogueTranslateResponseBodyData  {
  return s.Data
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponseBody) GetErrCode() *string  {
  return s.ErrCode
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponseBody) GetErrMessage() *string  {
  return s.ErrMessage
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponseBody) SetData(v *ExecuteAITeacherSyncDialogueTranslateResponseBodyData) *ExecuteAITeacherSyncDialogueTranslateResponseBody {
  s.Data = v
  return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponseBody) SetErrCode(v string) *ExecuteAITeacherSyncDialogueTranslateResponseBody {
  s.ErrCode = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponseBody) SetErrMessage(v string) *ExecuteAITeacherSyncDialogueTranslateResponseBody {
  s.ErrMessage = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponseBody) SetHttpStatusCode(v int32) *ExecuteAITeacherSyncDialogueTranslateResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponseBody) SetRequestId(v string) *ExecuteAITeacherSyncDialogueTranslateResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponseBody) SetSuccess(v bool) *ExecuteAITeacherSyncDialogueTranslateResponseBody {
  s.Success = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteAITeacherSyncDialogueTranslateResponseBodyData struct {
  // example:
  // 
  // 太好了，谢谢你过来，莉莉。你喜欢吃肉吗？
  Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ExecuteAITeacherSyncDialogueTranslateResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueTranslateResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponseBodyData) GetResult() *string  {
  return s.Result
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponseBodyData) SetResult(v string) *ExecuteAITeacherSyncDialogueTranslateResponseBodyData {
  s.Result = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponseBodyData) Validate() error {
  return dara.Validate(s)
}

