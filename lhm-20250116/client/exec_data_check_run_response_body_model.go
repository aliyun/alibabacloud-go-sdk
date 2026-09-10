// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecDataCheckRunResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetErrCode(v string) *ExecDataCheckRunResponseBody
  GetErrCode() *string 
  SetErrMessage(v string) *ExecDataCheckRunResponseBody
  GetErrMessage() *string 
  SetRequestId(v string) *ExecDataCheckRunResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecDataCheckRunResponseBody
  GetSuccess() *bool 
}

type ExecDataCheckRunResponseBody struct {
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
  // The request ID, which is used to locate and troubleshoot issues with this call.
  // 
  // example:
  // 
  // 4C467B38-3910-4477-9B0B-6963D83B4E72
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // Indicates whether the call is successful. A value of true indicates success. A value of false indicates failure. If the call fails, use errCode and errMessage to troubleshoot the issue.
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExecDataCheckRunResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecDataCheckRunResponseBody) GoString() string {
  return s.String()
}

func (s *ExecDataCheckRunResponseBody) GetErrCode() *string  {
  return s.ErrCode
}

func (s *ExecDataCheckRunResponseBody) GetErrMessage() *string  {
  return s.ErrMessage
}

func (s *ExecDataCheckRunResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecDataCheckRunResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecDataCheckRunResponseBody) SetErrCode(v string) *ExecDataCheckRunResponseBody {
  s.ErrCode = &v
  return s
}

func (s *ExecDataCheckRunResponseBody) SetErrMessage(v string) *ExecDataCheckRunResponseBody {
  s.ErrMessage = &v
  return s
}

func (s *ExecDataCheckRunResponseBody) SetRequestId(v string) *ExecDataCheckRunResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecDataCheckRunResponseBody) SetSuccess(v bool) *ExecDataCheckRunResponseBody {
  s.Success = &v
  return s
}

func (s *ExecDataCheckRunResponseBody) Validate() error {
  return dara.Validate(s)
}

