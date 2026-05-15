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
  Data *ExecuteStatementResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
  // example:
  // 
  // InvalidParameterValue
  ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
  ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
  HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // 819A7F0F-2951-540F-BD94-6A41ECF0281F
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
  // example:
  // 
  // InvalidParameterValue
  ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
  ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
  Results []*ExecuteStatementResponseBodyDataResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
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
  ColumnMetadata []*ExecuteStatementResponseBodyDataResultsColumnMetadata `json:"columnMetadata,omitempty" xml:"columnMetadata,omitempty" type:"Repeated"`
  Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
  ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
  ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
  // **Query ID**
  // 
  // example:
  // 
  // E3F4B2A7-1234-5678-9ABC-DEF012345678
  QueryId *string `json:"queryId,omitempty" xml:"queryId,omitempty"`
  Records [][]*string `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
  Sql *string `json:"sql,omitempty" xml:"sql,omitempty"`
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
  Truncated *bool `json:"truncated,omitempty" xml:"truncated,omitempty"`
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
  // example:
  // 
  // id
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  Nullable *bool `json:"nullable,omitempty" xml:"nullable,omitempty"`
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

