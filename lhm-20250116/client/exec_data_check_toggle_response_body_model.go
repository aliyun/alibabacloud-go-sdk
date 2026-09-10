// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecDataCheckToggleResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetErrCode(v string) *ExecDataCheckToggleResponseBody
  GetErrCode() *string 
  SetErrMessage(v string) *ExecDataCheckToggleResponseBody
  GetErrMessage() *string 
  SetRequestId(v string) *ExecDataCheckToggleResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecDataCheckToggleResponseBody
  GetSuccess() *bool 
}

type ExecDataCheckToggleResponseBody struct {
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
  // Indicates whether the call is successful. A value of true indicates success. A value of false indicates failure. If the call fails, use errCode and errMessage to identify the cause.
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExecDataCheckToggleResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecDataCheckToggleResponseBody) GoString() string {
  return s.String()
}

func (s *ExecDataCheckToggleResponseBody) GetErrCode() *string  {
  return s.ErrCode
}

func (s *ExecDataCheckToggleResponseBody) GetErrMessage() *string  {
  return s.ErrMessage
}

func (s *ExecDataCheckToggleResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecDataCheckToggleResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecDataCheckToggleResponseBody) SetErrCode(v string) *ExecDataCheckToggleResponseBody {
  s.ErrCode = &v
  return s
}

func (s *ExecDataCheckToggleResponseBody) SetErrMessage(v string) *ExecDataCheckToggleResponseBody {
  s.ErrMessage = &v
  return s
}

func (s *ExecDataCheckToggleResponseBody) SetRequestId(v string) *ExecDataCheckToggleResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecDataCheckToggleResponseBody) SetSuccess(v bool) *ExecDataCheckToggleResponseBody {
  s.Success = &v
  return s
}

func (s *ExecDataCheckToggleResponseBody) Validate() error {
  return dara.Validate(s)
}

