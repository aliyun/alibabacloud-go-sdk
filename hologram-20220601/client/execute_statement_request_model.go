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
  // The database.
  // 
  // example:
  // 
  // test_db
  DbName *string `json:"dbName,omitempty" xml:"dbName,omitempty"`
  // The maximum amount of data to return in a single request, in bytes. Default value: 10485760 (10 MB). Excess data is truncated and indicated by setting truncated to true.
  // 
  // example:
  // 
  // 1024
  MaxBytes *int64 `json:"maxBytes,omitempty" xml:"maxBytes,omitempty"`
  // The maximum number of rows to return. Default value: 200. Maximum value: 1000. Excess rows are truncated and indicated by the Truncated field in the response.
  // 
  // example:
  // 
  // 300
  MaxRows *int64 `json:"maxRows,omitempty" xml:"maxRows,omitempty"`
  // The list of parameters to attach to the SQL parameterized query, which correspond in sequence to placeholders such as $1 and $2 in the SQL statement. Supported value types include String, Long, Double, Boolean, and null, which represents NULL. Using a parameterized query effectively prevents SQL injection and improves the cache reuse rate of precompiled statements. The parameter list applies to each SQL statement in Statements.
  Parameters []interface{} `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Repeated"`
  // The timeout period for SQL execution, in seconds. Default value: 30. Maximum value: 30. Minimum value: 1. If a timeout occurs, the server actively cancels the query.
  // 
  // example:
  // 
  // 5
  QueryTimeout *int64 `json:"queryTimeout,omitempty" xml:"queryTimeout,omitempty"`
  // The SQL statement to be executed. The maximum length is 16,384 characters. Multiple SQL statements separated by semicolons are supported. If multiple statements are specified, the result of the last statement is returned.
  // 
  // example:
  // 
  // select 	- from test_table limit 10;
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

