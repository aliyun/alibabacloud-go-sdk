// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableExecuteStatementResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v bool) *EnableExecuteStatementResponseBody
  GetData() *bool 
  SetErrorCode(v string) *EnableExecuteStatementResponseBody
  GetErrorCode() *string 
  SetErrorMessage(v string) *EnableExecuteStatementResponseBody
  GetErrorMessage() *string 
  SetHttpStatusCode(v string) *EnableExecuteStatementResponseBody
  GetHttpStatusCode() *string 
  SetRequestId(v string) *EnableExecuteStatementResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableExecuteStatementResponseBody
  GetSuccess() *bool 
}

type EnableExecuteStatementResponseBody struct {
  Data *bool `json:"data,omitempty" xml:"data,omitempty"`
  // example:
  // 
  // InvalidParameterValue
  ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
  ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
  HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
  // example:
  // 
  // E3F4B2A7-1234-5678-9ABC-DEF012345678
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s EnableExecuteStatementResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableExecuteStatementResponseBody) GoString() string {
  return s.String()
}

func (s *EnableExecuteStatementResponseBody) GetData() *bool  {
  return s.Data
}

func (s *EnableExecuteStatementResponseBody) GetErrorCode() *string  {
  return s.ErrorCode
}

func (s *EnableExecuteStatementResponseBody) GetErrorMessage() *string  {
  return s.ErrorMessage
}

func (s *EnableExecuteStatementResponseBody) GetHttpStatusCode() *string  {
  return s.HttpStatusCode
}

func (s *EnableExecuteStatementResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableExecuteStatementResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableExecuteStatementResponseBody) SetData(v bool) *EnableExecuteStatementResponseBody {
  s.Data = &v
  return s
}

func (s *EnableExecuteStatementResponseBody) SetErrorCode(v string) *EnableExecuteStatementResponseBody {
  s.ErrorCode = &v
  return s
}

func (s *EnableExecuteStatementResponseBody) SetErrorMessage(v string) *EnableExecuteStatementResponseBody {
  s.ErrorMessage = &v
  return s
}

func (s *EnableExecuteStatementResponseBody) SetHttpStatusCode(v string) *EnableExecuteStatementResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *EnableExecuteStatementResponseBody) SetRequestId(v string) *EnableExecuteStatementResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableExecuteStatementResponseBody) SetSuccess(v bool) *EnableExecuteStatementResponseBody {
  s.Success = &v
  return s
}

func (s *EnableExecuteStatementResponseBody) Validate() error {
  return dara.Validate(s)
}

