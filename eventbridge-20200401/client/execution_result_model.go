// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecutionResult interface {
  dara.Model
  String() string
  GoString() string
  SetIsTruncated(v bool) *ExecutionResult
  GetIsTruncated() *bool 
  SetRowCount(v int32) *ExecutionResult
  GetRowCount() *int32 
  SetRows(v string) *ExecutionResult
  GetRows() *string 
  SetSchema(v []*SchemaColumn) *ExecutionResult
  GetSchema() []*SchemaColumn 
  SetTotalRows(v int32) *ExecutionResult
  GetTotalRows() *int32 
}

type ExecutionResult struct {
  // example:
  // 
  // false
  IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
  // example:
  // 
  // 2
  RowCount *int32 `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
  Rows *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
  Schema []*SchemaColumn `json:"Schema,omitempty" xml:"Schema,omitempty" type:"Repeated"`
  // example:
  // 
  // 2
  TotalRows *int32 `json:"TotalRows,omitempty" xml:"TotalRows,omitempty"`
}

func (s ExecutionResult) String() string {
  return dara.Prettify(s)
}

func (s ExecutionResult) GoString() string {
  return s.String()
}

func (s *ExecutionResult) GetIsTruncated() *bool  {
  return s.IsTruncated
}

func (s *ExecutionResult) GetRowCount() *int32  {
  return s.RowCount
}

func (s *ExecutionResult) GetRows() *string  {
  return s.Rows
}

func (s *ExecutionResult) GetSchema() []*SchemaColumn  {
  return s.Schema
}

func (s *ExecutionResult) GetTotalRows() *int32  {
  return s.TotalRows
}

func (s *ExecutionResult) SetIsTruncated(v bool) *ExecutionResult {
  s.IsTruncated = &v
  return s
}

func (s *ExecutionResult) SetRowCount(v int32) *ExecutionResult {
  s.RowCount = &v
  return s
}

func (s *ExecutionResult) SetRows(v string) *ExecutionResult {
  s.Rows = &v
  return s
}

func (s *ExecutionResult) SetSchema(v []*SchemaColumn) *ExecutionResult {
  s.Schema = v
  return s
}

func (s *ExecutionResult) SetTotalRows(v int32) *ExecutionResult {
  s.TotalRows = &v
  return s
}

func (s *ExecutionResult) Validate() error {
  if s.Schema != nil {
    for _, item := range s.Schema {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

