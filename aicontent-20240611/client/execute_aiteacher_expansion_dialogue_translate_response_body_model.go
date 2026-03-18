// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAITeacherExpansionDialogueTranslateResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v *ExecuteAITeacherExpansionDialogueTranslateResponseBodyData) *ExecuteAITeacherExpansionDialogueTranslateResponseBody
  GetData() *ExecuteAITeacherExpansionDialogueTranslateResponseBodyData 
  SetErrCode(v string) *ExecuteAITeacherExpansionDialogueTranslateResponseBody
  GetErrCode() *string 
  SetErrMessage(v string) *ExecuteAITeacherExpansionDialogueTranslateResponseBody
  GetErrMessage() *string 
  SetHttpStatusCode(v int32) *ExecuteAITeacherExpansionDialogueTranslateResponseBody
  GetHttpStatusCode() *int32 
  SetRequestId(v string) *ExecuteAITeacherExpansionDialogueTranslateResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecuteAITeacherExpansionDialogueTranslateResponseBody
  GetSuccess() *bool 
}

type ExecuteAITeacherExpansionDialogueTranslateResponseBody struct {
  // example:
  // 
  // []
  Data *ExecuteAITeacherExpansionDialogueTranslateResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s ExecuteAITeacherExpansionDialogueTranslateResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueTranslateResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponseBody) GetData() *ExecuteAITeacherExpansionDialogueTranslateResponseBodyData  {
  return s.Data
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponseBody) GetErrCode() *string  {
  return s.ErrCode
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponseBody) GetErrMessage() *string  {
  return s.ErrMessage
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponseBody) SetData(v *ExecuteAITeacherExpansionDialogueTranslateResponseBodyData) *ExecuteAITeacherExpansionDialogueTranslateResponseBody {
  s.Data = v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponseBody) SetErrCode(v string) *ExecuteAITeacherExpansionDialogueTranslateResponseBody {
  s.ErrCode = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponseBody) SetErrMessage(v string) *ExecuteAITeacherExpansionDialogueTranslateResponseBody {
  s.ErrMessage = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponseBody) SetHttpStatusCode(v int32) *ExecuteAITeacherExpansionDialogueTranslateResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponseBody) SetRequestId(v string) *ExecuteAITeacherExpansionDialogueTranslateResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponseBody) SetSuccess(v bool) *ExecuteAITeacherExpansionDialogueTranslateResponseBody {
  s.Success = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteAITeacherExpansionDialogueTranslateResponseBodyData struct {
  // example:
  // 
  // 太好了，谢谢你过来，莉莉。你喜欢吃肉吗？
  Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueTranslateResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueTranslateResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponseBodyData) GetResult() *string  {
  return s.Result
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponseBodyData) SetResult(v string) *ExecuteAITeacherExpansionDialogueTranslateResponseBodyData {
  s.Result = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponseBodyData) Validate() error {
  return dara.Validate(s)
}

