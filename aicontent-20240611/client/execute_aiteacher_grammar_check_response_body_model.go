// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAITeacherGrammarCheckResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v *ExecuteAITeacherGrammarCheckResponseBodyData) *ExecuteAITeacherGrammarCheckResponseBody
  GetData() *ExecuteAITeacherGrammarCheckResponseBodyData 
  SetErrCode(v string) *ExecuteAITeacherGrammarCheckResponseBody
  GetErrCode() *string 
  SetErrMessage(v string) *ExecuteAITeacherGrammarCheckResponseBody
  GetErrMessage() *string 
  SetHttpStatusCode(v int32) *ExecuteAITeacherGrammarCheckResponseBody
  GetHttpStatusCode() *int32 
  SetRequestId(v string) *ExecuteAITeacherGrammarCheckResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecuteAITeacherGrammarCheckResponseBody
  GetSuccess() *bool 
}

type ExecuteAITeacherGrammarCheckResponseBody struct {
  // The data object that contains the result.
  // 
  // example:
  // 
  // []
  Data *ExecuteAITeacherGrammarCheckResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
  // The error code.
  // 
  // example:
  // 
  // UNKNOWN_ERROR
  ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
  // The error message.
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
  // The request ID.
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

func (s ExecuteAITeacherGrammarCheckResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherGrammarCheckResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherGrammarCheckResponseBody) GetData() *ExecuteAITeacherGrammarCheckResponseBodyData  {
  return s.Data
}

func (s *ExecuteAITeacherGrammarCheckResponseBody) GetErrCode() *string  {
  return s.ErrCode
}

func (s *ExecuteAITeacherGrammarCheckResponseBody) GetErrMessage() *string  {
  return s.ErrMessage
}

func (s *ExecuteAITeacherGrammarCheckResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExecuteAITeacherGrammarCheckResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteAITeacherGrammarCheckResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecuteAITeacherGrammarCheckResponseBody) SetData(v *ExecuteAITeacherGrammarCheckResponseBodyData) *ExecuteAITeacherGrammarCheckResponseBody {
  s.Data = v
  return s
}

func (s *ExecuteAITeacherGrammarCheckResponseBody) SetErrCode(v string) *ExecuteAITeacherGrammarCheckResponseBody {
  s.ErrCode = &v
  return s
}

func (s *ExecuteAITeacherGrammarCheckResponseBody) SetErrMessage(v string) *ExecuteAITeacherGrammarCheckResponseBody {
  s.ErrMessage = &v
  return s
}

func (s *ExecuteAITeacherGrammarCheckResponseBody) SetHttpStatusCode(v int32) *ExecuteAITeacherGrammarCheckResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExecuteAITeacherGrammarCheckResponseBody) SetRequestId(v string) *ExecuteAITeacherGrammarCheckResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteAITeacherGrammarCheckResponseBody) SetSuccess(v bool) *ExecuteAITeacherGrammarCheckResponseBody {
  s.Success = &v
  return s
}

func (s *ExecuteAITeacherGrammarCheckResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteAITeacherGrammarCheckResponseBodyData struct {
  // The error analysis.
  // 
  // example:
  // 
  // 主语 "I" 对应的动词应该是 "am" 而不是 "is"。
  Analysis *string `json:"analysis,omitempty" xml:"analysis,omitempty"`
  // The corrected sentence.
  // 
  // example:
  // 
  // I am good.
  Correction *string `json:"correction,omitempty" xml:"correction,omitempty"`
  // The result of the grammar check.
  // 
  // example:
  // 
  // Has_Error
  CorrectionStatus *string `json:"correctionStatus,omitempty" xml:"correctionStatus,omitempty"`
  // The error reason.
  // 
  // example:
  // 
  // 暂无返回
  ErrorReason *string `json:"errorReason,omitempty" xml:"errorReason,omitempty"`
}

func (s ExecuteAITeacherGrammarCheckResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherGrammarCheckResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherGrammarCheckResponseBodyData) GetAnalysis() *string  {
  return s.Analysis
}

func (s *ExecuteAITeacherGrammarCheckResponseBodyData) GetCorrection() *string  {
  return s.Correction
}

func (s *ExecuteAITeacherGrammarCheckResponseBodyData) GetCorrectionStatus() *string  {
  return s.CorrectionStatus
}

func (s *ExecuteAITeacherGrammarCheckResponseBodyData) GetErrorReason() *string  {
  return s.ErrorReason
}

func (s *ExecuteAITeacherGrammarCheckResponseBodyData) SetAnalysis(v string) *ExecuteAITeacherGrammarCheckResponseBodyData {
  s.Analysis = &v
  return s
}

func (s *ExecuteAITeacherGrammarCheckResponseBodyData) SetCorrection(v string) *ExecuteAITeacherGrammarCheckResponseBodyData {
  s.Correction = &v
  return s
}

func (s *ExecuteAITeacherGrammarCheckResponseBodyData) SetCorrectionStatus(v string) *ExecuteAITeacherGrammarCheckResponseBodyData {
  s.CorrectionStatus = &v
  return s
}

func (s *ExecuteAITeacherGrammarCheckResponseBodyData) SetErrorReason(v string) *ExecuteAITeacherGrammarCheckResponseBodyData {
  s.ErrorReason = &v
  return s
}

func (s *ExecuteAITeacherGrammarCheckResponseBodyData) Validate() error {
  return dara.Validate(s)
}

