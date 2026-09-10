// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecDataCheckSqlPreviewResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v string) *ExecDataCheckSqlPreviewResponseBody
  GetData() *string 
  SetErrCode(v string) *ExecDataCheckSqlPreviewResponseBody
  GetErrCode() *string 
  SetErrMessage(v string) *ExecDataCheckSqlPreviewResponseBody
  GetErrMessage() *string 
  SetRequestId(v string) *ExecDataCheckSqlPreviewResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecDataCheckSqlPreviewResponseBody
  GetSuccess() *bool 
}

type ExecDataCheckSqlPreviewResponseBody struct {
  // The preview SQL statement.
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
  // The request ID, which is used to locate and troubleshoot issues with this call.
  // 
  // example:
  // 
  // 4C467B38-3910-4477-9B0B-6963D83B4E72
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // Indicates whether the call is successful. Valid values: true and false. If false is returned, use errCode and errMessage to troubleshoot the issue.
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExecDataCheckSqlPreviewResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecDataCheckSqlPreviewResponseBody) GoString() string {
  return s.String()
}

func (s *ExecDataCheckSqlPreviewResponseBody) GetData() *string  {
  return s.Data
}

func (s *ExecDataCheckSqlPreviewResponseBody) GetErrCode() *string  {
  return s.ErrCode
}

func (s *ExecDataCheckSqlPreviewResponseBody) GetErrMessage() *string  {
  return s.ErrMessage
}

func (s *ExecDataCheckSqlPreviewResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecDataCheckSqlPreviewResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecDataCheckSqlPreviewResponseBody) SetData(v string) *ExecDataCheckSqlPreviewResponseBody {
  s.Data = &v
  return s
}

func (s *ExecDataCheckSqlPreviewResponseBody) SetErrCode(v string) *ExecDataCheckSqlPreviewResponseBody {
  s.ErrCode = &v
  return s
}

func (s *ExecDataCheckSqlPreviewResponseBody) SetErrMessage(v string) *ExecDataCheckSqlPreviewResponseBody {
  s.ErrMessage = &v
  return s
}

func (s *ExecDataCheckSqlPreviewResponseBody) SetRequestId(v string) *ExecDataCheckSqlPreviewResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecDataCheckSqlPreviewResponseBody) SetSuccess(v bool) *ExecDataCheckSqlPreviewResponseBody {
  s.Success = &v
  return s
}

func (s *ExecDataCheckSqlPreviewResponseBody) Validate() error {
  return dara.Validate(s)
}

