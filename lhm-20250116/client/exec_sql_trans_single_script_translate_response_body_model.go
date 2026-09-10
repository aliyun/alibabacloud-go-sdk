// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecSqlTransSingleScriptTranslateResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v string) *ExecSqlTransSingleScriptTranslateResponseBody
  GetData() *string 
  SetErrCode(v string) *ExecSqlTransSingleScriptTranslateResponseBody
  GetErrCode() *string 
  SetErrMessage(v string) *ExecSqlTransSingleScriptTranslateResponseBody
  GetErrMessage() *string 
  SetRequestId(v string) *ExecSqlTransSingleScriptTranslateResponseBody
  GetRequestId() *string 
  SetSuccess(v string) *ExecSqlTransSingleScriptTranslateResponseBody
  GetSuccess() *string 
}

type ExecSqlTransSingleScriptTranslateResponseBody struct {
  // The business data returned by the operation (in string format). The specific content varies by operation.
  // 
  // example:
  // 
  // demo
  Data *string `json:"data,omitempty" xml:"data,omitempty"`
  // The error code. An empty string is returned if the call is successful.
  // 
  // example:
  // 
  // Success
  ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
  // The error message. An empty string is returned if the call is successful.
  // 
  // example:
  // 
  // success
  ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
  // The request ID, which uniquely identifies the call. Provide this value when troubleshooting issues.
  // 
  // example:
  // 
  // 4C467B38-3910-4477-9B0B-6963D83B4E72
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // Indicates whether the call is successful. A value of true indicates success. A value of false indicates failure. If the call fails, use errCode and errMessage to troubleshoot the issue.
  // 
  // example:
  // 
  // true
  Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExecSqlTransSingleScriptTranslateResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecSqlTransSingleScriptTranslateResponseBody) GoString() string {
  return s.String()
}

func (s *ExecSqlTransSingleScriptTranslateResponseBody) GetData() *string  {
  return s.Data
}

func (s *ExecSqlTransSingleScriptTranslateResponseBody) GetErrCode() *string  {
  return s.ErrCode
}

func (s *ExecSqlTransSingleScriptTranslateResponseBody) GetErrMessage() *string  {
  return s.ErrMessage
}

func (s *ExecSqlTransSingleScriptTranslateResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecSqlTransSingleScriptTranslateResponseBody) GetSuccess() *string  {
  return s.Success
}

func (s *ExecSqlTransSingleScriptTranslateResponseBody) SetData(v string) *ExecSqlTransSingleScriptTranslateResponseBody {
  s.Data = &v
  return s
}

func (s *ExecSqlTransSingleScriptTranslateResponseBody) SetErrCode(v string) *ExecSqlTransSingleScriptTranslateResponseBody {
  s.ErrCode = &v
  return s
}

func (s *ExecSqlTransSingleScriptTranslateResponseBody) SetErrMessage(v string) *ExecSqlTransSingleScriptTranslateResponseBody {
  s.ErrMessage = &v
  return s
}

func (s *ExecSqlTransSingleScriptTranslateResponseBody) SetRequestId(v string) *ExecSqlTransSingleScriptTranslateResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecSqlTransSingleScriptTranslateResponseBody) SetSuccess(v string) *ExecSqlTransSingleScriptTranslateResponseBody {
  s.Success = &v
  return s
}

func (s *ExecSqlTransSingleScriptTranslateResponseBody) Validate() error {
  return dara.Validate(s)
}

