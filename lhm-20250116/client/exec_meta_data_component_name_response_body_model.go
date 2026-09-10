// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecMetaDataComponentNameResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v bool) *ExecMetaDataComponentNameResponseBody
  GetData() *bool 
  SetErrCode(v string) *ExecMetaDataComponentNameResponseBody
  GetErrCode() *string 
  SetErrMessage(v string) *ExecMetaDataComponentNameResponseBody
  GetErrMessage() *string 
  SetRequestId(v string) *ExecMetaDataComponentNameResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecMetaDataComponentNameResponseBody
  GetSuccess() *bool 
}

type ExecMetaDataComponentNameResponseBody struct {
  // The name check result. A value of true indicates that a datasource with the same name already exists. A value of false indicates that the name is not in use.
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
  // The request ID, which uniquely identifies this call. Provide this value when troubleshooting issues.
  // 
  // example:
  // 
  // 4C467B38-3910-4477-9B0B-6963D83B4E72
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // Indicates whether the call is successful. A value of true indicates success. A value of false indicates failure. If the call fails, use errCode and errMessage to troubleshoot the issue.
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExecMetaDataComponentNameResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecMetaDataComponentNameResponseBody) GoString() string {
  return s.String()
}

func (s *ExecMetaDataComponentNameResponseBody) GetData() *bool  {
  return s.Data
}

func (s *ExecMetaDataComponentNameResponseBody) GetErrCode() *string  {
  return s.ErrCode
}

func (s *ExecMetaDataComponentNameResponseBody) GetErrMessage() *string  {
  return s.ErrMessage
}

func (s *ExecMetaDataComponentNameResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecMetaDataComponentNameResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecMetaDataComponentNameResponseBody) SetData(v bool) *ExecMetaDataComponentNameResponseBody {
  s.Data = &v
  return s
}

func (s *ExecMetaDataComponentNameResponseBody) SetErrCode(v string) *ExecMetaDataComponentNameResponseBody {
  s.ErrCode = &v
  return s
}

func (s *ExecMetaDataComponentNameResponseBody) SetErrMessage(v string) *ExecMetaDataComponentNameResponseBody {
  s.ErrMessage = &v
  return s
}

func (s *ExecMetaDataComponentNameResponseBody) SetRequestId(v string) *ExecMetaDataComponentNameResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecMetaDataComponentNameResponseBody) SetSuccess(v bool) *ExecMetaDataComponentNameResponseBody {
  s.Success = &v
  return s
}

func (s *ExecMetaDataComponentNameResponseBody) Validate() error {
  return dara.Validate(s)
}

