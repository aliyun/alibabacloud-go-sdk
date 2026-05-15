// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteStatementRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDbName(v string) *ExecuteStatementRequest
  GetDbName() *string 
  SetMaxBytes(v int64) *ExecuteStatementRequest
  GetMaxBytes() *int64 
  SetMaxRows(v int64) *ExecuteStatementRequest
  GetMaxRows() *int64 
  SetParameters(v []interface{}) *ExecuteStatementRequest
  GetParameters() []interface{} 
  SetQueryTimeout(v int64) *ExecuteStatementRequest
  GetQueryTimeout() *int64 
  SetSql(v string) *ExecuteStatementRequest
  GetSql() *string 
}

type ExecuteStatementRequest struct {
  DbName *string `json:"dbName,omitempty" xml:"dbName,omitempty"`
  MaxBytes *int64 `json:"maxBytes,omitempty" xml:"maxBytes,omitempty"`
  MaxRows *int64 `json:"maxRows,omitempty" xml:"maxRows,omitempty"`
  Parameters []interface{} `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Repeated"`
  QueryTimeout *int64 `json:"queryTimeout,omitempty" xml:"queryTimeout,omitempty"`
  Sql *string `json:"sql,omitempty" xml:"sql,omitempty"`
}

func (s ExecuteStatementRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteStatementRequest) GoString() string {
  return s.String()
}

func (s *ExecuteStatementRequest) GetDbName() *string  {
  return s.DbName
}

func (s *ExecuteStatementRequest) GetMaxBytes() *int64  {
  return s.MaxBytes
}

func (s *ExecuteStatementRequest) GetMaxRows() *int64  {
  return s.MaxRows
}

func (s *ExecuteStatementRequest) GetParameters() []interface{}  {
  return s.Parameters
}

func (s *ExecuteStatementRequest) GetQueryTimeout() *int64  {
  return s.QueryTimeout
}

func (s *ExecuteStatementRequest) GetSql() *string  {
  return s.Sql
}

func (s *ExecuteStatementRequest) SetDbName(v string) *ExecuteStatementRequest {
  s.DbName = &v
  return s
}

func (s *ExecuteStatementRequest) SetMaxBytes(v int64) *ExecuteStatementRequest {
  s.MaxBytes = &v
  return s
}

func (s *ExecuteStatementRequest) SetMaxRows(v int64) *ExecuteStatementRequest {
  s.MaxRows = &v
  return s
}

func (s *ExecuteStatementRequest) SetParameters(v []interface{}) *ExecuteStatementRequest {
  s.Parameters = v
  return s
}

func (s *ExecuteStatementRequest) SetQueryTimeout(v int64) *ExecuteStatementRequest {
  s.QueryTimeout = &v
  return s
}

func (s *ExecuteStatementRequest) SetSql(v string) *ExecuteStatementRequest {
  s.Sql = &v
  return s
}

func (s *ExecuteStatementRequest) Validate() error {
  return dara.Validate(s)
}

