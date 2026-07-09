// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteQueryRequest interface {
  dara.Model
  String() string
  GoString() string
  SetFrom(v int32) *ExecuteQueryRequest
  GetFrom() *int32 
  SetLength(v int32) *ExecuteQueryRequest
  GetLength() *int32 
  SetMaxOutputLength(v int32) *ExecuteQueryRequest
  GetMaxOutputLength() *int32 
  SetOffset(v int32) *ExecuteQueryRequest
  GetOffset() *int32 
  SetQuery(v string) *ExecuteQueryRequest
  GetQuery() *string 
  SetTo(v int32) *ExecuteQueryRequest
  GetTo() *int32 
  SetType(v string) *ExecuteQueryRequest
  GetType() *string 
}

type ExecuteQueryRequest struct {
  From *int32 `json:"from,omitempty" xml:"from,omitempty"`
  Length *int32 `json:"length,omitempty" xml:"length,omitempty"`
  MaxOutputLength *int32 `json:"maxOutputLength,omitempty" xml:"maxOutputLength,omitempty"`
  Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
  // The query entered by the user.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // select count(*) from product_faq_dataset
  Query *string `json:"query,omitempty" xml:"query,omitempty"`
  To *int32 `json:"to,omitempty" xml:"to,omitempty"`
  // The statement type. Currently, only SQL is supported.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // SQL
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ExecuteQueryRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteQueryRequest) GoString() string {
  return s.String()
}

func (s *ExecuteQueryRequest) GetFrom() *int32  {
  return s.From
}

func (s *ExecuteQueryRequest) GetLength() *int32  {
  return s.Length
}

func (s *ExecuteQueryRequest) GetMaxOutputLength() *int32  {
  return s.MaxOutputLength
}

func (s *ExecuteQueryRequest) GetOffset() *int32  {
  return s.Offset
}

func (s *ExecuteQueryRequest) GetQuery() *string  {
  return s.Query
}

func (s *ExecuteQueryRequest) GetTo() *int32  {
  return s.To
}

func (s *ExecuteQueryRequest) GetType() *string  {
  return s.Type
}

func (s *ExecuteQueryRequest) SetFrom(v int32) *ExecuteQueryRequest {
  s.From = &v
  return s
}

func (s *ExecuteQueryRequest) SetLength(v int32) *ExecuteQueryRequest {
  s.Length = &v
  return s
}

func (s *ExecuteQueryRequest) SetMaxOutputLength(v int32) *ExecuteQueryRequest {
  s.MaxOutputLength = &v
  return s
}

func (s *ExecuteQueryRequest) SetOffset(v int32) *ExecuteQueryRequest {
  s.Offset = &v
  return s
}

func (s *ExecuteQueryRequest) SetQuery(v string) *ExecuteQueryRequest {
  s.Query = &v
  return s
}

func (s *ExecuteQueryRequest) SetTo(v int32) *ExecuteQueryRequest {
  s.To = &v
  return s
}

func (s *ExecuteQueryRequest) SetType(v string) *ExecuteQueryRequest {
  s.Type = &v
  return s
}

func (s *ExecuteQueryRequest) Validate() error {
  return dara.Validate(s)
}

