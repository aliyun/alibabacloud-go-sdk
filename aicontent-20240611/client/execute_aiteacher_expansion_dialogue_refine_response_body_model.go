// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAITeacherExpansionDialogueRefineResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v *ExecuteAITeacherExpansionDialogueRefineResponseBodyData) *ExecuteAITeacherExpansionDialogueRefineResponseBody
  GetData() *ExecuteAITeacherExpansionDialogueRefineResponseBodyData 
  SetErrCode(v string) *ExecuteAITeacherExpansionDialogueRefineResponseBody
  GetErrCode() *string 
  SetErrMessage(v string) *ExecuteAITeacherExpansionDialogueRefineResponseBody
  GetErrMessage() *string 
  SetHttpStatusCode(v int32) *ExecuteAITeacherExpansionDialogueRefineResponseBody
  GetHttpStatusCode() *int32 
  SetRequestId(v string) *ExecuteAITeacherExpansionDialogueRefineResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecuteAITeacherExpansionDialogueRefineResponseBody
  GetSuccess() *bool 
}

type ExecuteAITeacherExpansionDialogueRefineResponseBody struct {
  // example:
  // 
  // []
  Data *ExecuteAITeacherExpansionDialogueRefineResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s ExecuteAITeacherExpansionDialogueRefineResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueRefineResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponseBody) GetData() *ExecuteAITeacherExpansionDialogueRefineResponseBodyData  {
  return s.Data
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponseBody) GetErrCode() *string  {
  return s.ErrCode
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponseBody) GetErrMessage() *string  {
  return s.ErrMessage
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponseBody) SetData(v *ExecuteAITeacherExpansionDialogueRefineResponseBodyData) *ExecuteAITeacherExpansionDialogueRefineResponseBody {
  s.Data = v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponseBody) SetErrCode(v string) *ExecuteAITeacherExpansionDialogueRefineResponseBody {
  s.ErrCode = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponseBody) SetErrMessage(v string) *ExecuteAITeacherExpansionDialogueRefineResponseBody {
  s.ErrMessage = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponseBody) SetHttpStatusCode(v int32) *ExecuteAITeacherExpansionDialogueRefineResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponseBody) SetRequestId(v string) *ExecuteAITeacherExpansionDialogueRefineResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponseBody) SetSuccess(v bool) *ExecuteAITeacherExpansionDialogueRefineResponseBody {
  s.Success = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteAITeacherExpansionDialogueRefineResponseBodyData struct {
  // example:
  // 
  // Yes, I\\"ll be right there.
  Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueRefineResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueRefineResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponseBodyData) GetResult() *string  {
  return s.Result
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponseBodyData) SetResult(v string) *ExecuteAITeacherExpansionDialogueRefineResponseBodyData {
  s.Result = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponseBodyData) Validate() error {
  return dara.Validate(s)
}

