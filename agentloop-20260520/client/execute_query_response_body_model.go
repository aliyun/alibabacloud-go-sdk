// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteQueryResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetColumnTypes(v []*string) *ExecuteQueryResponseBody
  GetColumnTypes() []*string 
  SetColumns(v []*string) *ExecuteQueryResponseBody
  GetColumns() []*string 
  SetMeta(v *ExecuteQueryResponseBodyMeta) *ExecuteQueryResponseBody
  GetMeta() *ExecuteQueryResponseBodyMeta 
  SetRequestId(v string) *ExecuteQueryResponseBody
  GetRequestId() *string 
  SetRows(v [][]interface{}) *ExecuteQueryResponseBody
  GetRows() [][]interface{} 
}

type ExecuteQueryResponseBody struct {
  // The result column types.
  ColumnTypes []*string `json:"columnTypes,omitempty" xml:"columnTypes,omitempty" type:"Repeated"`
  // The result column information.
  Columns []*string `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
  // The metadata of the returned data.
  Meta *ExecuteQueryResponseBodyMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Struct"`
  // The request ID.
  // 
  // example:
  // 
  // EB27D183-8F6C-5C5A-A6A3-E0508AF54F78
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // The result rows.
  Rows [][]interface{} `json:"rows,omitempty" xml:"rows,omitempty" type:"Repeated"`
}

func (s ExecuteQueryResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteQueryResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteQueryResponseBody) GetColumnTypes() []*string  {
  return s.ColumnTypes
}

func (s *ExecuteQueryResponseBody) GetColumns() []*string  {
  return s.Columns
}

func (s *ExecuteQueryResponseBody) GetMeta() *ExecuteQueryResponseBodyMeta  {
  return s.Meta
}

func (s *ExecuteQueryResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteQueryResponseBody) GetRows() [][]interface{}  {
  return s.Rows
}

func (s *ExecuteQueryResponseBody) SetColumnTypes(v []*string) *ExecuteQueryResponseBody {
  s.ColumnTypes = v
  return s
}

func (s *ExecuteQueryResponseBody) SetColumns(v []*string) *ExecuteQueryResponseBody {
  s.Columns = v
  return s
}

func (s *ExecuteQueryResponseBody) SetMeta(v *ExecuteQueryResponseBodyMeta) *ExecuteQueryResponseBody {
  s.Meta = v
  return s
}

func (s *ExecuteQueryResponseBody) SetRequestId(v string) *ExecuteQueryResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteQueryResponseBody) SetRows(v [][]interface{}) *ExecuteQueryResponseBody {
  s.Rows = v
  return s
}

func (s *ExecuteQueryResponseBody) Validate() error {
  if s.Meta != nil {
    if err := s.Meta.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteQueryResponseBodyMeta struct {
  // The number of log rows scanned or processed.
  // 
  // example:
  // 
  // 100
  AffectedRows *int32 `json:"affectedRows,omitempty" xml:"affectedRows,omitempty"`
  // The number of log rows returned by this query request.
  // 
  // example:
  // 
  // 1
  Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
  // The time consumed by this execution, in milliseconds.
  // 
  // example:
  // 
  // 15
  ElapsedMillisecond *int64 `json:"elapsedMillisecond,omitempty" xml:"elapsedMillisecond,omitempty"`
  // Indicates whether the query result is complete.
  // 
  // example:
  // 
  // Complete
  Progress *string `json:"progress,omitempty" xml:"progress,omitempty"`
  Truncation *ExecuteQueryResponseBodyMetaTruncation `json:"truncation,omitempty" xml:"truncation,omitempty" type:"Struct"`
}

func (s ExecuteQueryResponseBodyMeta) String() string {
  return dara.Prettify(s)
}

func (s ExecuteQueryResponseBodyMeta) GoString() string {
  return s.String()
}

func (s *ExecuteQueryResponseBodyMeta) GetAffectedRows() *int32  {
  return s.AffectedRows
}

func (s *ExecuteQueryResponseBodyMeta) GetCount() *int32  {
  return s.Count
}

func (s *ExecuteQueryResponseBodyMeta) GetElapsedMillisecond() *int64  {
  return s.ElapsedMillisecond
}

func (s *ExecuteQueryResponseBodyMeta) GetProgress() *string  {
  return s.Progress
}

func (s *ExecuteQueryResponseBodyMeta) GetTruncation() *ExecuteQueryResponseBodyMetaTruncation  {
  return s.Truncation
}

func (s *ExecuteQueryResponseBodyMeta) SetAffectedRows(v int32) *ExecuteQueryResponseBodyMeta {
  s.AffectedRows = &v
  return s
}

func (s *ExecuteQueryResponseBodyMeta) SetCount(v int32) *ExecuteQueryResponseBodyMeta {
  s.Count = &v
  return s
}

func (s *ExecuteQueryResponseBodyMeta) SetElapsedMillisecond(v int64) *ExecuteQueryResponseBodyMeta {
  s.ElapsedMillisecond = &v
  return s
}

func (s *ExecuteQueryResponseBodyMeta) SetProgress(v string) *ExecuteQueryResponseBodyMeta {
  s.Progress = &v
  return s
}

func (s *ExecuteQueryResponseBodyMeta) SetTruncation(v *ExecuteQueryResponseBodyMetaTruncation) *ExecuteQueryResponseBodyMeta {
  s.Truncation = v
  return s
}

func (s *ExecuteQueryResponseBodyMeta) Validate() error {
  if s.Truncation != nil {
    if err := s.Truncation.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteQueryResponseBodyMetaTruncation struct {
  Truncated *bool `json:"truncated,omitempty" xml:"truncated,omitempty"`
  TruncatedColumnIndexes [][]*int32 `json:"truncatedColumnIndexes,omitempty" xml:"truncatedColumnIndexes,omitempty" type:"Repeated"`
}

func (s ExecuteQueryResponseBodyMetaTruncation) String() string {
  return dara.Prettify(s)
}

func (s ExecuteQueryResponseBodyMetaTruncation) GoString() string {
  return s.String()
}

func (s *ExecuteQueryResponseBodyMetaTruncation) GetTruncated() *bool  {
  return s.Truncated
}

func (s *ExecuteQueryResponseBodyMetaTruncation) GetTruncatedColumnIndexes() [][]*int32  {
  return s.TruncatedColumnIndexes
}

func (s *ExecuteQueryResponseBodyMetaTruncation) SetTruncated(v bool) *ExecuteQueryResponseBodyMetaTruncation {
  s.Truncated = &v
  return s
}

func (s *ExecuteQueryResponseBodyMetaTruncation) SetTruncatedColumnIndexes(v [][]*int32) *ExecuteQueryResponseBodyMetaTruncation {
  s.TruncatedColumnIndexes = v
  return s
}

func (s *ExecuteQueryResponseBodyMetaTruncation) Validate() error {
  return dara.Validate(s)
}

