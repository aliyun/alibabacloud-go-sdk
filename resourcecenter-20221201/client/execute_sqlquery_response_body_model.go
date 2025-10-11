// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteSQLQueryResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetColumns(v []*ExecuteSQLQueryResponseBodyColumns) *ExecuteSQLQueryResponseBody
  GetColumns() []*ExecuteSQLQueryResponseBodyColumns 
  SetMaxResults(v int32) *ExecuteSQLQueryResponseBody
  GetMaxResults() *int32 
  SetNextToken(v string) *ExecuteSQLQueryResponseBody
  GetNextToken() *string 
  SetRequestId(v string) *ExecuteSQLQueryResponseBody
  GetRequestId() *string 
  SetRows(v []interface{}) *ExecuteSQLQueryResponseBody
  GetRows() []interface{} 
}

type ExecuteSQLQueryResponseBody struct {
  // The columns.
  Columns []*ExecuteSQLQueryResponseBodyColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
  // The number of entries per page.
  // 
  // example:
  // 
  // 1000
  MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
  // A pagination token. It can be used in the next request to retrieve a new page of results.
  // 
  // example:
  // 
  // eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
  NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // D696E6EF-3A6D-5770-801E-4982081FE4D0
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // An array of search results.
  Rows []interface{} `json:"Rows,omitempty" xml:"Rows,omitempty" type:"Repeated"`
}

func (s ExecuteSQLQueryResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteSQLQueryResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteSQLQueryResponseBody) GetColumns() []*ExecuteSQLQueryResponseBodyColumns  {
  return s.Columns
}

func (s *ExecuteSQLQueryResponseBody) GetMaxResults() *int32  {
  return s.MaxResults
}

func (s *ExecuteSQLQueryResponseBody) GetNextToken() *string  {
  return s.NextToken
}

func (s *ExecuteSQLQueryResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteSQLQueryResponseBody) GetRows() []interface{}  {
  return s.Rows
}

func (s *ExecuteSQLQueryResponseBody) SetColumns(v []*ExecuteSQLQueryResponseBodyColumns) *ExecuteSQLQueryResponseBody {
  s.Columns = v
  return s
}

func (s *ExecuteSQLQueryResponseBody) SetMaxResults(v int32) *ExecuteSQLQueryResponseBody {
  s.MaxResults = &v
  return s
}

func (s *ExecuteSQLQueryResponseBody) SetNextToken(v string) *ExecuteSQLQueryResponseBody {
  s.NextToken = &v
  return s
}

func (s *ExecuteSQLQueryResponseBody) SetRequestId(v string) *ExecuteSQLQueryResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteSQLQueryResponseBody) SetRows(v []interface{}) *ExecuteSQLQueryResponseBody {
  s.Rows = v
  return s
}

func (s *ExecuteSQLQueryResponseBody) Validate() error {
  return dara.Validate(s)
}

type ExecuteSQLQueryResponseBodyColumns struct {
  // The name of the column.
  // 
  // example:
  // 
  // resource_id
  Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
  // The type of the column.
  // 
  // example:
  // 
  // varchar
  Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ExecuteSQLQueryResponseBodyColumns) String() string {
  return dara.Prettify(s)
}

func (s ExecuteSQLQueryResponseBodyColumns) GoString() string {
  return s.String()
}

func (s *ExecuteSQLQueryResponseBodyColumns) GetName() *string  {
  return s.Name
}

func (s *ExecuteSQLQueryResponseBodyColumns) GetType() *string  {
  return s.Type
}

func (s *ExecuteSQLQueryResponseBodyColumns) SetName(v string) *ExecuteSQLQueryResponseBodyColumns {
  s.Name = &v
  return s
}

func (s *ExecuteSQLQueryResponseBodyColumns) SetType(v string) *ExecuteSQLQueryResponseBodyColumns {
  s.Type = &v
  return s
}

func (s *ExecuteSQLQueryResponseBodyColumns) Validate() error {
  return dara.Validate(s)
}

