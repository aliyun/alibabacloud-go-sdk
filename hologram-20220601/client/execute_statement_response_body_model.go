// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteStatementResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v *ExecuteStatementResponseBodyData) *ExecuteStatementResponseBody
  GetData() *ExecuteStatementResponseBodyData 
  SetErrorCode(v string) *ExecuteStatementResponseBody
  GetErrorCode() *string 
  SetErrorMessage(v string) *ExecuteStatementResponseBody
  GetErrorMessage() *string 
  SetHttpStatusCode(v string) *ExecuteStatementResponseBody
  GetHttpStatusCode() *string 
  SetRequestId(v string) *ExecuteStatementResponseBody
  GetRequestId() *string 
  SetSuccess(v string) *ExecuteStatementResponseBody
  GetSuccess() *string 
}

type ExecuteStatementResponseBody struct {
  // Specifies whether OpenAPI SQL execution is allowed. Valid values: true: Allowed. false: Not allowed.
  Data *ExecuteStatementResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
  // The error code. This field is empty if the request is successful.
  // 
  // example:
  // 
  // InvalidParameterValue
  ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
  // The error message. This field is empty if the request is successful.
  // 
  // example:
  // 
  // Invalid parameter value, such as an empty or excessively long SQL statement
  ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
  // The HTTP status code.
  // 
  // example:
  // 
  // 200
  HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // 819A7F0F-2951-540F-BD94-6A41ECF0281F
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // Indicates whether the request is successful.
  // 
  // example:
  // 
  // True
  Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExecuteStatementResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteStatementResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteStatementResponseBody) GetData() *ExecuteStatementResponseBodyData  {
  return s.Data
}

func (s *ExecuteStatementResponseBody) GetErrorCode() *string  {
  return s.ErrorCode
}

func (s *ExecuteStatementResponseBody) GetErrorMessage() *string  {
  return s.ErrorMessage
}

func (s *ExecuteStatementResponseBody) GetHttpStatusCode() *string  {
  return s.HttpStatusCode
}

func (s *ExecuteStatementResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteStatementResponseBody) GetSuccess() *string  {
  return s.Success
}

func (s *ExecuteStatementResponseBody) SetData(v *ExecuteStatementResponseBodyData) *ExecuteStatementResponseBody {
  s.Data = v
  return s
}

func (s *ExecuteStatementResponseBody) SetErrorCode(v string) *ExecuteStatementResponseBody {
  s.ErrorCode = &v
  return s
}

func (s *ExecuteStatementResponseBody) SetErrorMessage(v string) *ExecuteStatementResponseBody {
  s.ErrorMessage = &v
  return s
}

func (s *ExecuteStatementResponseBody) SetHttpStatusCode(v string) *ExecuteStatementResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExecuteStatementResponseBody) SetRequestId(v string) *ExecuteStatementResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteStatementResponseBody) SetSuccess(v string) *ExecuteStatementResponseBody {
  s.Success = &v
  return s
}

func (s *ExecuteStatementResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteStatementResponseBodyData struct {
  // The error code. This field is empty if the request is successful.
  // 
  // example:
  // 
  // InvalidParameterValue
  ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
  // The error message. This field is empty if the request is successful.
  // 
  // example:
  // 
  // Invalid parameter value, such as an empty or excessively long SQL statement
  ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
  // The list of execution results. The length is currently fixed to 1. If multiple SELECT queries are executed, only the result of the last SELECT query is returned.
  Results []*ExecuteStatementResponseBodyDataResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
  // Indicates whether the request is successful.
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExecuteStatementResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExecuteStatementResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExecuteStatementResponseBodyData) GetErrorCode() *string  {
  return s.ErrorCode
}

func (s *ExecuteStatementResponseBodyData) GetErrorMessage() *string  {
  return s.ErrorMessage
}

func (s *ExecuteStatementResponseBodyData) GetResults() []*ExecuteStatementResponseBodyDataResults  {
  return s.Results
}

func (s *ExecuteStatementResponseBodyData) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecuteStatementResponseBodyData) SetErrorCode(v string) *ExecuteStatementResponseBodyData {
  s.ErrorCode = &v
  return s
}

func (s *ExecuteStatementResponseBodyData) SetErrorMessage(v string) *ExecuteStatementResponseBodyData {
  s.ErrorMessage = &v
  return s
}

func (s *ExecuteStatementResponseBodyData) SetResults(v []*ExecuteStatementResponseBodyDataResults) *ExecuteStatementResponseBodyData {
  s.Results = v
  return s
}

func (s *ExecuteStatementResponseBodyData) SetSuccess(v bool) *ExecuteStatementResponseBodyData {
  s.Success = &v
  return s
}

func (s *ExecuteStatementResponseBodyData) Validate() error {
  if s.Results != nil {
    for _, item := range s.Results {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type ExecuteStatementResponseBodyDataResults struct {
  // The **column metadata**.
  ColumnMetadata []*ExecuteStatementResponseBodyDataResultsColumnMetadata `json:"columnMetadata,omitempty" xml:"columnMetadata,omitempty" type:"Repeated"`
  // The number of rows returned.
  // 
  // example:
  // 
  // 25
  Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
  // The error code for the SQL statement.
  // 
  // example:
  // 
  // SQL_ERROR
  ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
  // The error message for the SQL statement.
  // 
  // example:
  // 
  // ERROR: relation \\"non_existent_table\\" does not exist\\n  Position: 15
  ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
  // **Query ID**
  // 
  // example:
  // 
  // E3F4B2A7-1234-5678-9ABC-DEF012345678
  QueryId *string `json:"queryId,omitempty" xml:"queryId,omitempty"`
  // The set of query result rows. Each row is an array of strings. All values are serialized as strings, and NULL values are represented as \\N.
  Records [][]*string `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
  // The SQL statement that is actually executed.
  // 
  // example:
  // 
  // select 	- from test_table limit 10;
  Sql *string `json:"sql,omitempty" xml:"sql,omitempty"`
  // Indicates whether the request is successful.
  // 
  // example:
  // 
  // True
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
  // Indicates whether the result is truncated. The value is true if the result exceeds MaxResults.
  Truncated *bool `json:"truncated,omitempty" xml:"truncated,omitempty"`
  // The number of rows affected by the INSERT, UPDATE, or DELETE statement.
  // 
  // example:
  // 
  // 10
  UpdateCount *int32 `json:"updateCount,omitempty" xml:"updateCount,omitempty"`
}

func (s ExecuteStatementResponseBodyDataResults) String() string {
  return dara.Prettify(s)
}

func (s ExecuteStatementResponseBodyDataResults) GoString() string {
  return s.String()
}

func (s *ExecuteStatementResponseBodyDataResults) GetColumnMetadata() []*ExecuteStatementResponseBodyDataResultsColumnMetadata  {
  return s.ColumnMetadata
}

func (s *ExecuteStatementResponseBodyDataResults) GetCount() *int32  {
  return s.Count
}

func (s *ExecuteStatementResponseBodyDataResults) GetErrorCode() *string  {
  return s.ErrorCode
}

func (s *ExecuteStatementResponseBodyDataResults) GetErrorMessage() *string  {
  return s.ErrorMessage
}

func (s *ExecuteStatementResponseBodyDataResults) GetQueryId() *string  {
  return s.QueryId
}

func (s *ExecuteStatementResponseBodyDataResults) GetRecords() [][]*string  {
  return s.Records
}

func (s *ExecuteStatementResponseBodyDataResults) GetSql() *string  {
  return s.Sql
}

func (s *ExecuteStatementResponseBodyDataResults) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecuteStatementResponseBodyDataResults) GetTruncated() *bool  {
  return s.Truncated
}

func (s *ExecuteStatementResponseBodyDataResults) GetUpdateCount() *int32  {
  return s.UpdateCount
}

func (s *ExecuteStatementResponseBodyDataResults) SetColumnMetadata(v []*ExecuteStatementResponseBodyDataResultsColumnMetadata) *ExecuteStatementResponseBodyDataResults {
  s.ColumnMetadata = v
  return s
}

func (s *ExecuteStatementResponseBodyDataResults) SetCount(v int32) *ExecuteStatementResponseBodyDataResults {
  s.Count = &v
  return s
}

func (s *ExecuteStatementResponseBodyDataResults) SetErrorCode(v string) *ExecuteStatementResponseBodyDataResults {
  s.ErrorCode = &v
  return s
}

func (s *ExecuteStatementResponseBodyDataResults) SetErrorMessage(v string) *ExecuteStatementResponseBodyDataResults {
  s.ErrorMessage = &v
  return s
}

func (s *ExecuteStatementResponseBodyDataResults) SetQueryId(v string) *ExecuteStatementResponseBodyDataResults {
  s.QueryId = &v
  return s
}

func (s *ExecuteStatementResponseBodyDataResults) SetRecords(v [][]*string) *ExecuteStatementResponseBodyDataResults {
  s.Records = v
  return s
}

func (s *ExecuteStatementResponseBodyDataResults) SetSql(v string) *ExecuteStatementResponseBodyDataResults {
  s.Sql = &v
  return s
}

func (s *ExecuteStatementResponseBodyDataResults) SetSuccess(v bool) *ExecuteStatementResponseBodyDataResults {
  s.Success = &v
  return s
}

func (s *ExecuteStatementResponseBodyDataResults) SetTruncated(v bool) *ExecuteStatementResponseBodyDataResults {
  s.Truncated = &v
  return s
}

func (s *ExecuteStatementResponseBodyDataResults) SetUpdateCount(v int32) *ExecuteStatementResponseBodyDataResults {
  s.UpdateCount = &v
  return s
}

func (s *ExecuteStatementResponseBodyDataResults) Validate() error {
  if s.ColumnMetadata != nil {
    for _, item := range s.ColumnMetadata {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type ExecuteStatementResponseBodyDataResultsColumnMetadata struct {
  // The **column name**.
  // 
  // example:
  // 
  // id
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // Indicates whether the column is nullable.
  Nullable *bool `json:"nullable,omitempty" xml:"nullable,omitempty"`
  // The data type, such as int4, text, or bool.
  // 
  // example:
  // 
  // int4
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ExecuteStatementResponseBodyDataResultsColumnMetadata) String() string {
  return dara.Prettify(s)
}

func (s ExecuteStatementResponseBodyDataResultsColumnMetadata) GoString() string {
  return s.String()
}

func (s *ExecuteStatementResponseBodyDataResultsColumnMetadata) GetName() *string  {
  return s.Name
}

func (s *ExecuteStatementResponseBodyDataResultsColumnMetadata) GetNullable() *bool  {
  return s.Nullable
}

func (s *ExecuteStatementResponseBodyDataResultsColumnMetadata) GetType() *string  {
  return s.Type
}

func (s *ExecuteStatementResponseBodyDataResultsColumnMetadata) SetName(v string) *ExecuteStatementResponseBodyDataResultsColumnMetadata {
  s.Name = &v
  return s
}

func (s *ExecuteStatementResponseBodyDataResultsColumnMetadata) SetNullable(v bool) *ExecuteStatementResponseBodyDataResultsColumnMetadata {
  s.Nullable = &v
  return s
}

func (s *ExecuteStatementResponseBodyDataResultsColumnMetadata) SetType(v string) *ExecuteStatementResponseBodyDataResultsColumnMetadata {
  s.Type = &v
  return s
}

func (s *ExecuteStatementResponseBodyDataResultsColumnMetadata) Validate() error {
  return dara.Validate(s)
}

