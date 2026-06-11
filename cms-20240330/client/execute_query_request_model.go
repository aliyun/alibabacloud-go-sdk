// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteQueryRequest interface {
  dara.Model
  String() string
  GoString() string
  SetQuery(v string) *ExecuteQueryRequest
  GetQuery() *string 
  SetType(v string) *ExecuteQueryRequest
  GetType() *string 
}

type ExecuteQueryRequest struct {
  // The query statement to execute.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // SELECT count(1) from "datasetname"
  Query *string `json:"query,omitempty" xml:"query,omitempty"`
  // The query type. Valid values: SQL and SPL.
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

func (s *ExecuteQueryRequest) GetQuery() *string  {
  return s.Query
}

func (s *ExecuteQueryRequest) GetType() *string  {
  return s.Type
}

func (s *ExecuteQueryRequest) SetQuery(v string) *ExecuteQueryRequest {
  s.Query = &v
  return s
}

func (s *ExecuteQueryRequest) SetType(v string) *ExecuteQueryRequest {
  s.Type = &v
  return s
}

func (s *ExecuteQueryRequest) Validate() error {
  return dara.Validate(s)
}

