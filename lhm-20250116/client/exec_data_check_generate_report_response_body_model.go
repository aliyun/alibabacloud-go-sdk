// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecDataCheckGenerateReportResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetErrCode(v string) *ExecDataCheckGenerateReportResponseBody
  GetErrCode() *string 
  SetErrMessage(v string) *ExecDataCheckGenerateReportResponseBody
  GetErrMessage() *string 
  SetRequestId(v string) *ExecDataCheckGenerateReportResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecDataCheckGenerateReportResponseBody
  GetSuccess() *bool 
}

type ExecDataCheckGenerateReportResponseBody struct {
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
  // The request ID. You can use this ID to locate and troubleshoot issues.
  // 
  // example:
  // 
  // 4C467B38-3910-4477-9B0B-6963D83B4E72
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // Indicates whether the call is successful. A value of true indicates success. A value of false indicates failure. If the call fails, use errCode and errMessage to troubleshoot the issue.
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExecDataCheckGenerateReportResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecDataCheckGenerateReportResponseBody) GoString() string {
  return s.String()
}

func (s *ExecDataCheckGenerateReportResponseBody) GetErrCode() *string  {
  return s.ErrCode
}

func (s *ExecDataCheckGenerateReportResponseBody) GetErrMessage() *string  {
  return s.ErrMessage
}

func (s *ExecDataCheckGenerateReportResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecDataCheckGenerateReportResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecDataCheckGenerateReportResponseBody) SetErrCode(v string) *ExecDataCheckGenerateReportResponseBody {
  s.ErrCode = &v
  return s
}

func (s *ExecDataCheckGenerateReportResponseBody) SetErrMessage(v string) *ExecDataCheckGenerateReportResponseBody {
  s.ErrMessage = &v
  return s
}

func (s *ExecDataCheckGenerateReportResponseBody) SetRequestId(v string) *ExecDataCheckGenerateReportResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecDataCheckGenerateReportResponseBody) SetSuccess(v bool) *ExecDataCheckGenerateReportResponseBody {
  s.Success = &v
  return s
}

func (s *ExecDataCheckGenerateReportResponseBody) Validate() error {
  return dara.Validate(s)
}

