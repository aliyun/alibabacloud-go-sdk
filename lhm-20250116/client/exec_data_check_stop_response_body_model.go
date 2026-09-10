// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecDataCheckStopResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetErrCode(v string) *ExecDataCheckStopResponseBody
  GetErrCode() *string 
  SetErrMessage(v string) *ExecDataCheckStopResponseBody
  GetErrMessage() *string 
  SetRequestId(v string) *ExecDataCheckStopResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecDataCheckStopResponseBody
  GetSuccess() *bool 
}

type ExecDataCheckStopResponseBody struct {
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

func (s ExecDataCheckStopResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecDataCheckStopResponseBody) GoString() string {
  return s.String()
}

func (s *ExecDataCheckStopResponseBody) GetErrCode() *string  {
  return s.ErrCode
}

func (s *ExecDataCheckStopResponseBody) GetErrMessage() *string  {
  return s.ErrMessage
}

func (s *ExecDataCheckStopResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecDataCheckStopResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecDataCheckStopResponseBody) SetErrCode(v string) *ExecDataCheckStopResponseBody {
  s.ErrCode = &v
  return s
}

func (s *ExecDataCheckStopResponseBody) SetErrMessage(v string) *ExecDataCheckStopResponseBody {
  s.ErrMessage = &v
  return s
}

func (s *ExecDataCheckStopResponseBody) SetRequestId(v string) *ExecDataCheckStopResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecDataCheckStopResponseBody) SetSuccess(v bool) *ExecDataCheckStopResponseBody {
  s.Success = &v
  return s
}

func (s *ExecDataCheckStopResponseBody) Validate() error {
  return dara.Validate(s)
}

