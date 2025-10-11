// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteMultiAccountSQLQueryResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetColumns(v []*ExecuteMultiAccountSQLQueryResponseBodyColumns) *ExecuteMultiAccountSQLQueryResponseBody
  GetColumns() []*ExecuteMultiAccountSQLQueryResponseBodyColumns 
  SetMaxResults(v int32) *ExecuteMultiAccountSQLQueryResponseBody
  GetMaxResults() *int32 
  SetNextToken(v string) *ExecuteMultiAccountSQLQueryResponseBody
  GetNextToken() *string 
  SetRequestId(v string) *ExecuteMultiAccountSQLQueryResponseBody
  GetRequestId() *string 
  SetRows(v []interface{}) *ExecuteMultiAccountSQLQueryResponseBody
  GetRows() []interface{} 
}

type ExecuteMultiAccountSQLQueryResponseBody struct {
  // The columns.
  Columns []*ExecuteMultiAccountSQLQueryResponseBodyColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
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
  // 44C8A952-D6B0-5BC8-82D5-93BA02E26F2E
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // An array of search results.
  Rows []interface{} `json:"Rows,omitempty" xml:"Rows,omitempty" type:"Repeated"`
}

func (s ExecuteMultiAccountSQLQueryResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteMultiAccountSQLQueryResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteMultiAccountSQLQueryResponseBody) GetColumns() []*ExecuteMultiAccountSQLQueryResponseBodyColumns  {
  return s.Columns
}

func (s *ExecuteMultiAccountSQLQueryResponseBody) GetMaxResults() *int32  {
  return s.MaxResults
}

func (s *ExecuteMultiAccountSQLQueryResponseBody) GetNextToken() *string  {
  return s.NextToken
}

func (s *ExecuteMultiAccountSQLQueryResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteMultiAccountSQLQueryResponseBody) GetRows() []interface{}  {
  return s.Rows
}

func (s *ExecuteMultiAccountSQLQueryResponseBody) SetColumns(v []*ExecuteMultiAccountSQLQueryResponseBodyColumns) *ExecuteMultiAccountSQLQueryResponseBody {
  s.Columns = v
  return s
}

func (s *ExecuteMultiAccountSQLQueryResponseBody) SetMaxResults(v int32) *ExecuteMultiAccountSQLQueryResponseBody {
  s.MaxResults = &v
  return s
}

func (s *ExecuteMultiAccountSQLQueryResponseBody) SetNextToken(v string) *ExecuteMultiAccountSQLQueryResponseBody {
  s.NextToken = &v
  return s
}

func (s *ExecuteMultiAccountSQLQueryResponseBody) SetRequestId(v string) *ExecuteMultiAccountSQLQueryResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteMultiAccountSQLQueryResponseBody) SetRows(v []interface{}) *ExecuteMultiAccountSQLQueryResponseBody {
  s.Rows = v
  return s
}

func (s *ExecuteMultiAccountSQLQueryResponseBody) Validate() error {
  return dara.Validate(s)
}

type ExecuteMultiAccountSQLQueryResponseBodyColumns struct {
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

func (s ExecuteMultiAccountSQLQueryResponseBodyColumns) String() string {
  return dara.Prettify(s)
}

func (s ExecuteMultiAccountSQLQueryResponseBodyColumns) GoString() string {
  return s.String()
}

func (s *ExecuteMultiAccountSQLQueryResponseBodyColumns) GetName() *string  {
  return s.Name
}

func (s *ExecuteMultiAccountSQLQueryResponseBodyColumns) GetType() *string  {
  return s.Type
}

func (s *ExecuteMultiAccountSQLQueryResponseBodyColumns) SetName(v string) *ExecuteMultiAccountSQLQueryResponseBodyColumns {
  s.Name = &v
  return s
}

func (s *ExecuteMultiAccountSQLQueryResponseBodyColumns) SetType(v string) *ExecuteMultiAccountSQLQueryResponseBodyColumns {
  s.Type = &v
  return s
}

func (s *ExecuteMultiAccountSQLQueryResponseBodyColumns) Validate() error {
  return dara.Validate(s)
}

