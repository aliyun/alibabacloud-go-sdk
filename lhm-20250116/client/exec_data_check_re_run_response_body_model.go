// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecDataCheckReRunResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v int64) *ExecDataCheckReRunResponseBody
  GetData() *int64 
  SetErrCode(v string) *ExecDataCheckReRunResponseBody
  GetErrCode() *string 
  SetErrMessage(v string) *ExecDataCheckReRunResponseBody
  GetErrMessage() *string 
  SetRequestId(v string) *ExecDataCheckReRunResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecDataCheckReRunResponseBody
  GetSuccess() *bool 
}

type ExecDataCheckReRunResponseBody struct {
  // The ID of the new batch created by the rerun.
  // 
  // example:
  // 
  // 100
  Data *int64 `json:"data,omitempty" xml:"data,omitempty"`
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
  // The request ID, which is used to locate and troubleshoot issues.
  // 
  // example:
  // 
  // 4C467B38-3910-4477-9B0B-6963D83B4E72
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // Indicates whether the call is successful. A value of true indicates success. A value of false indicates failure. If the call fails, use errCode and errMessage to troubleshoot the issue.
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExecDataCheckReRunResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecDataCheckReRunResponseBody) GoString() string {
  return s.String()
}

func (s *ExecDataCheckReRunResponseBody) GetData() *int64  {
  return s.Data
}

func (s *ExecDataCheckReRunResponseBody) GetErrCode() *string  {
  return s.ErrCode
}

func (s *ExecDataCheckReRunResponseBody) GetErrMessage() *string  {
  return s.ErrMessage
}

func (s *ExecDataCheckReRunResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecDataCheckReRunResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecDataCheckReRunResponseBody) SetData(v int64) *ExecDataCheckReRunResponseBody {
  s.Data = &v
  return s
}

func (s *ExecDataCheckReRunResponseBody) SetErrCode(v string) *ExecDataCheckReRunResponseBody {
  s.ErrCode = &v
  return s
}

func (s *ExecDataCheckReRunResponseBody) SetErrMessage(v string) *ExecDataCheckReRunResponseBody {
  s.ErrMessage = &v
  return s
}

func (s *ExecDataCheckReRunResponseBody) SetRequestId(v string) *ExecDataCheckReRunResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecDataCheckReRunResponseBody) SetSuccess(v bool) *ExecDataCheckReRunResponseBody {
  s.Success = &v
  return s
}

func (s *ExecDataCheckReRunResponseBody) Validate() error {
  return dara.Validate(s)
}

