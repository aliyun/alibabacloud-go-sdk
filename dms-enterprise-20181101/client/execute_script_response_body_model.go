// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteScriptResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetErrorCode(v string) *ExecuteScriptResponseBody
  GetErrorCode() *string 
  SetErrorMessage(v string) *ExecuteScriptResponseBody
  GetErrorMessage() *string 
  SetRequestId(v string) *ExecuteScriptResponseBody
  GetRequestId() *string 
  SetResults(v []*ExecuteScriptResponseBodyResults) *ExecuteScriptResponseBody
  GetResults() []*ExecuteScriptResponseBodyResults 
  SetSuccess(v bool) *ExecuteScriptResponseBody
  GetSuccess() *bool 
}

type ExecuteScriptResponseBody struct {
  // The error code.
  // 
  // example:
  // 
  // UnknownError
  ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
  // The error message about the gateway.
  // 
  // example:
  // 
  // UnknownError
  ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
  // The ID of the request.
  // 
  // example:
  // 
  // FE8EE2F1-4880-46BC-A704-5CF63EAF9A04
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // The results of the SQL statements that are executed, in the format of an array. Each entry in the array indicates the result of an SQL statement.
  Results []*ExecuteScriptResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
  // Indicates whether the request is successful.
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecuteScriptResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteScriptResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteScriptResponseBody) GetErrorCode() *string  {
  return s.ErrorCode
}

func (s *ExecuteScriptResponseBody) GetErrorMessage() *string  {
  return s.ErrorMessage
}

func (s *ExecuteScriptResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteScriptResponseBody) GetResults() []*ExecuteScriptResponseBodyResults  {
  return s.Results
}

func (s *ExecuteScriptResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecuteScriptResponseBody) SetErrorCode(v string) *ExecuteScriptResponseBody {
  s.ErrorCode = &v
  return s
}

func (s *ExecuteScriptResponseBody) SetErrorMessage(v string) *ExecuteScriptResponseBody {
  s.ErrorMessage = &v
  return s
}

func (s *ExecuteScriptResponseBody) SetRequestId(v string) *ExecuteScriptResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteScriptResponseBody) SetResults(v []*ExecuteScriptResponseBodyResults) *ExecuteScriptResponseBody {
  s.Results = v
  return s
}

func (s *ExecuteScriptResponseBody) SetSuccess(v bool) *ExecuteScriptResponseBody {
  s.Success = &v
  return s
}

func (s *ExecuteScriptResponseBody) Validate() error {
  return dara.Validate(s)
}

type ExecuteScriptResponseBodyResults struct {
  // The fields that are queried after the SQL statement is executed.
  ColumnNames []*string `json:"ColumnNames,omitempty" xml:"ColumnNames,omitempty" type:"Repeated"`
  // The error message that is returned if the SQL statement fails to be executed. For example, an error message is returned because the SQL statement is invalid.
  // 
  // example:
  // 
  // UnknownError
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The total number of entries that are returned.
  // 
  // example:
  // 
  // 1
  RowCount *int64 `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
  // The rows that are queried after the SQL statement is executed.
  Rows []map[string]interface{} `json:"Rows,omitempty" xml:"Rows,omitempty" type:"Repeated"`
  // Indicates whether the SQL statement is executed.
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecuteScriptResponseBodyResults) String() string {
  return dara.Prettify(s)
}

func (s ExecuteScriptResponseBodyResults) GoString() string {
  return s.String()
}

func (s *ExecuteScriptResponseBodyResults) GetColumnNames() []*string  {
  return s.ColumnNames
}

func (s *ExecuteScriptResponseBodyResults) GetMessage() *string  {
  return s.Message
}

func (s *ExecuteScriptResponseBodyResults) GetRowCount() *int64  {
  return s.RowCount
}

func (s *ExecuteScriptResponseBodyResults) GetRows() []map[string]interface{}  {
  return s.Rows
}

func (s *ExecuteScriptResponseBodyResults) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecuteScriptResponseBodyResults) SetColumnNames(v []*string) *ExecuteScriptResponseBodyResults {
  s.ColumnNames = v
  return s
}

func (s *ExecuteScriptResponseBodyResults) SetMessage(v string) *ExecuteScriptResponseBodyResults {
  s.Message = &v
  return s
}

func (s *ExecuteScriptResponseBodyResults) SetRowCount(v int64) *ExecuteScriptResponseBodyResults {
  s.RowCount = &v
  return s
}

func (s *ExecuteScriptResponseBodyResults) SetRows(v []map[string]interface{}) *ExecuteScriptResponseBodyResults {
  s.Rows = v
  return s
}

func (s *ExecuteScriptResponseBodyResults) SetSuccess(v bool) *ExecuteScriptResponseBodyResults {
  s.Success = &v
  return s
}

func (s *ExecuteScriptResponseBodyResults) Validate() error {
  return dara.Validate(s)
}

