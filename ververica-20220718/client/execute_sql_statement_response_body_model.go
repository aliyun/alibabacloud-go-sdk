// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteSqlStatementResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v *SqlStatementExecuteResult) *ExecuteSqlStatementResponseBody
  GetData() *SqlStatementExecuteResult 
  SetErrorCode(v string) *ExecuteSqlStatementResponseBody
  GetErrorCode() *string 
  SetErrorMessage(v string) *ExecuteSqlStatementResponseBody
  GetErrorMessage() *string 
  SetHttpCode(v int32) *ExecuteSqlStatementResponseBody
  GetHttpCode() *int32 
  SetRequestId(v string) *ExecuteSqlStatementResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecuteSqlStatementResponseBody
  GetSuccess() *bool 
}

type ExecuteSqlStatementResponseBody struct {
  Data *SqlStatementExecuteResult `json:"data,omitempty" xml:"data,omitempty"`
  // example:
  // 
  // ""
  ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
  // example:
  // 
  // ""
  ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
  // example:
  // 
  // 200
  HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
  // example:
  // 
  // CBC799F0-AS7S-1D30-8A4F-882ED4DD****
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // example:
  // 
  // true
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExecuteSqlStatementResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteSqlStatementResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteSqlStatementResponseBody) GetData() *SqlStatementExecuteResult  {
  return s.Data
}

func (s *ExecuteSqlStatementResponseBody) GetErrorCode() *string  {
  return s.ErrorCode
}

func (s *ExecuteSqlStatementResponseBody) GetErrorMessage() *string  {
  return s.ErrorMessage
}

func (s *ExecuteSqlStatementResponseBody) GetHttpCode() *int32  {
  return s.HttpCode
}

func (s *ExecuteSqlStatementResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteSqlStatementResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecuteSqlStatementResponseBody) SetData(v *SqlStatementExecuteResult) *ExecuteSqlStatementResponseBody {
  s.Data = v
  return s
}

func (s *ExecuteSqlStatementResponseBody) SetErrorCode(v string) *ExecuteSqlStatementResponseBody {
  s.ErrorCode = &v
  return s
}

func (s *ExecuteSqlStatementResponseBody) SetErrorMessage(v string) *ExecuteSqlStatementResponseBody {
  s.ErrorMessage = &v
  return s
}

func (s *ExecuteSqlStatementResponseBody) SetHttpCode(v int32) *ExecuteSqlStatementResponseBody {
  s.HttpCode = &v
  return s
}

func (s *ExecuteSqlStatementResponseBody) SetRequestId(v string) *ExecuteSqlStatementResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteSqlStatementResponseBody) SetSuccess(v bool) *ExecuteSqlStatementResponseBody {
  s.Success = &v
  return s
}

func (s *ExecuteSqlStatementResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

