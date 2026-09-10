// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecDataCheckDownloadReportResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v string) *ExecDataCheckDownloadReportResponseBody
  GetData() *string 
  SetErrCode(v string) *ExecDataCheckDownloadReportResponseBody
  GetErrCode() *string 
  SetErrMessage(v string) *ExecDataCheckDownloadReportResponseBody
  GetErrMessage() *string 
  SetRequestId(v string) *ExecDataCheckDownloadReportResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecDataCheckDownloadReportResponseBody
  GetSuccess() *bool 
}

type ExecDataCheckDownloadReportResponseBody struct {
  // The OSS download link for the report file.
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
  // The request ID. You can use this ID to locate and troubleshoot issues for this call.
  // 
  // example:
  // 
  // 4C467B38-3910-4477-9B0B-6963D83B4E72
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // Indicates whether the call is successful. Valid values: true and false. If false is returned, use errCode and errMessage to troubleshoot the issue.
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExecDataCheckDownloadReportResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecDataCheckDownloadReportResponseBody) GoString() string {
  return s.String()
}

func (s *ExecDataCheckDownloadReportResponseBody) GetData() *string  {
  return s.Data
}

func (s *ExecDataCheckDownloadReportResponseBody) GetErrCode() *string  {
  return s.ErrCode
}

func (s *ExecDataCheckDownloadReportResponseBody) GetErrMessage() *string  {
  return s.ErrMessage
}

func (s *ExecDataCheckDownloadReportResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecDataCheckDownloadReportResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecDataCheckDownloadReportResponseBody) SetData(v string) *ExecDataCheckDownloadReportResponseBody {
  s.Data = &v
  return s
}

func (s *ExecDataCheckDownloadReportResponseBody) SetErrCode(v string) *ExecDataCheckDownloadReportResponseBody {
  s.ErrCode = &v
  return s
}

func (s *ExecDataCheckDownloadReportResponseBody) SetErrMessage(v string) *ExecDataCheckDownloadReportResponseBody {
  s.ErrMessage = &v
  return s
}

func (s *ExecDataCheckDownloadReportResponseBody) SetRequestId(v string) *ExecDataCheckDownloadReportResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecDataCheckDownloadReportResponseBody) SetSuccess(v bool) *ExecDataCheckDownloadReportResponseBody {
  s.Success = &v
  return s
}

func (s *ExecDataCheckDownloadReportResponseBody) Validate() error {
  return dara.Validate(s)
}

