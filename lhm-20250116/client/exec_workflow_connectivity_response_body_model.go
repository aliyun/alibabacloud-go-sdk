// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecWorkflowConnectivityResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v bool) *ExecWorkflowConnectivityResponseBody
  GetData() *bool 
  SetErrCode(v string) *ExecWorkflowConnectivityResponseBody
  GetErrCode() *string 
  SetErrMessage(v string) *ExecWorkflowConnectivityResponseBody
  GetErrMessage() *string 
  SetRequestId(v string) *ExecWorkflowConnectivityResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecWorkflowConnectivityResponseBody
  GetSuccess() *bool 
}

type ExecWorkflowConnectivityResponseBody struct {
  // The business data flag returned by the operation. The specific meaning varies by operation.
  Data *bool `json:"data,omitempty" xml:"data,omitempty"`
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
  // The request ID. Use this ID to locate and troubleshoot issues with this call.
  // 
  // example:
  // 
  // 4C467B38-3910-4477-9B0B-6963D83B4E72
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // Indicates whether the call is successful. A value of true indicates success. A value of false indicates failure. If the call fails, check errCode and errMessage for details.
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExecWorkflowConnectivityResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecWorkflowConnectivityResponseBody) GoString() string {
  return s.String()
}

func (s *ExecWorkflowConnectivityResponseBody) GetData() *bool  {
  return s.Data
}

func (s *ExecWorkflowConnectivityResponseBody) GetErrCode() *string  {
  return s.ErrCode
}

func (s *ExecWorkflowConnectivityResponseBody) GetErrMessage() *string  {
  return s.ErrMessage
}

func (s *ExecWorkflowConnectivityResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecWorkflowConnectivityResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecWorkflowConnectivityResponseBody) SetData(v bool) *ExecWorkflowConnectivityResponseBody {
  s.Data = &v
  return s
}

func (s *ExecWorkflowConnectivityResponseBody) SetErrCode(v string) *ExecWorkflowConnectivityResponseBody {
  s.ErrCode = &v
  return s
}

func (s *ExecWorkflowConnectivityResponseBody) SetErrMessage(v string) *ExecWorkflowConnectivityResponseBody {
  s.ErrMessage = &v
  return s
}

func (s *ExecWorkflowConnectivityResponseBody) SetRequestId(v string) *ExecWorkflowConnectivityResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecWorkflowConnectivityResponseBody) SetSuccess(v bool) *ExecWorkflowConnectivityResponseBody {
  s.Success = &v
  return s
}

func (s *ExecWorkflowConnectivityResponseBody) Validate() error {
  return dara.Validate(s)
}

