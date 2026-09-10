// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecDataCheckRunFailedResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v int64) *ExecDataCheckRunFailedResponseBody
  GetData() *int64 
  SetErrCode(v string) *ExecDataCheckRunFailedResponseBody
  GetErrCode() *string 
  SetErrMessage(v string) *ExecDataCheckRunFailedResponseBody
  GetErrMessage() *string 
  SetRequestId(v string) *ExecDataCheckRunFailedResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecDataCheckRunFailedResponseBody
  GetSuccess() *bool 
}

type ExecDataCheckRunFailedResponseBody struct {
  // The ID of the new batch created for this rerun.
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
  // Indicates whether the call is successful. Valid values:
  // 
  // - true: The call is successful.
  // 
  // - false: The call failed. Check errCode and errMessage to identify the cause.
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExecDataCheckRunFailedResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecDataCheckRunFailedResponseBody) GoString() string {
  return s.String()
}

func (s *ExecDataCheckRunFailedResponseBody) GetData() *int64  {
  return s.Data
}

func (s *ExecDataCheckRunFailedResponseBody) GetErrCode() *string  {
  return s.ErrCode
}

func (s *ExecDataCheckRunFailedResponseBody) GetErrMessage() *string  {
  return s.ErrMessage
}

func (s *ExecDataCheckRunFailedResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecDataCheckRunFailedResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecDataCheckRunFailedResponseBody) SetData(v int64) *ExecDataCheckRunFailedResponseBody {
  s.Data = &v
  return s
}

func (s *ExecDataCheckRunFailedResponseBody) SetErrCode(v string) *ExecDataCheckRunFailedResponseBody {
  s.ErrCode = &v
  return s
}

func (s *ExecDataCheckRunFailedResponseBody) SetErrMessage(v string) *ExecDataCheckRunFailedResponseBody {
  s.ErrMessage = &v
  return s
}

func (s *ExecDataCheckRunFailedResponseBody) SetRequestId(v string) *ExecDataCheckRunFailedResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecDataCheckRunFailedResponseBody) SetSuccess(v bool) *ExecDataCheckRunFailedResponseBody {
  s.Success = &v
  return s
}

func (s *ExecDataCheckRunFailedResponseBody) Validate() error {
  return dara.Validate(s)
}

