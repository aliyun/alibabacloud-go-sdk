// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteSparkReplStatementRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppId(v string) *ExecuteSparkReplStatementRequest
  GetAppId() *string 
  SetCode(v string) *ExecuteSparkReplStatementRequest
  GetCode() *string 
  SetCodeType(v string) *ExecuteSparkReplStatementRequest
  GetCodeType() *string 
  SetSessionId(v int64) *ExecuteSparkReplStatementRequest
  GetSessionId() *int64 
}

type ExecuteSparkReplStatementRequest struct {
  // The application ID.
  // 
  // >  You can call the [ListSparkApps](https://help.aliyun.com/document_detail/455888.html) operation to query Spark application IDs.
  // 
  // example:
  // 
  // s202411071444hzdvk486d9d2001****
  AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
  // The code that you want to execute.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // print(1+1)
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // The language type of the code. Valid values:
  // 
  // 	- SCALA
  // 
  // 	- PYTHON
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // PYTHON
  CodeType *string `json:"CodeType,omitempty" xml:"CodeType,omitempty"`
  // The ID of the session that you want to use to execute the code.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 123
  SessionId *int64 `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s ExecuteSparkReplStatementRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteSparkReplStatementRequest) GoString() string {
  return s.String()
}

func (s *ExecuteSparkReplStatementRequest) GetAppId() *string  {
  return s.AppId
}

func (s *ExecuteSparkReplStatementRequest) GetCode() *string  {
  return s.Code
}

func (s *ExecuteSparkReplStatementRequest) GetCodeType() *string  {
  return s.CodeType
}

func (s *ExecuteSparkReplStatementRequest) GetSessionId() *int64  {
  return s.SessionId
}

func (s *ExecuteSparkReplStatementRequest) SetAppId(v string) *ExecuteSparkReplStatementRequest {
  s.AppId = &v
  return s
}

func (s *ExecuteSparkReplStatementRequest) SetCode(v string) *ExecuteSparkReplStatementRequest {
  s.Code = &v
  return s
}

func (s *ExecuteSparkReplStatementRequest) SetCodeType(v string) *ExecuteSparkReplStatementRequest {
  s.CodeType = &v
  return s
}

func (s *ExecuteSparkReplStatementRequest) SetSessionId(v int64) *ExecuteSparkReplStatementRequest {
  s.SessionId = &v
  return s
}

func (s *ExecuteSparkReplStatementRequest) Validate() error {
  return dara.Validate(s)
}

