// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteSQLQueryRequest interface {
  dara.Model
  String() string
  GoString() string
  SetExpression(v string) *ExecuteSQLQueryRequest
  GetExpression() *string 
  SetMaxResults(v int32) *ExecuteSQLQueryRequest
  GetMaxResults() *int32 
  SetNextToken(v string) *ExecuteSQLQueryRequest
  GetNextToken() *string 
  SetScope(v string) *ExecuteSQLQueryRequest
  GetScope() *string 
}

type ExecuteSQLQueryRequest struct {
  // The SQL statement to be executed.
  // 
  // The number of characters in the SQL statement must be less than 2,000.
  // 
  // For more information about the SQL syntax, see [Basic SQL syntax](https://help.aliyun.com/document_detail/2539395.html).
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // SELECT 	- FROM resources LIMIT 100;
  Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
  // The number of entries per page.
  // 
  // 	- Valid values: 1 to 1000.
  // 
  // 	- Default value: 1000.
  // 
  // example:
  // 
  // 1000
  MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
  // The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
  // 
  // example:
  // 
  // eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
  NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
  // The search scope.
  // 
  // Set this parameter to the ID of a resource group.
  // 
  // For information about how to obtain the ID of a resource group, see [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html).
  // 
  // example:
  // 
  // rg-acfmzawhxxc****
  Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s ExecuteSQLQueryRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteSQLQueryRequest) GoString() string {
  return s.String()
}

func (s *ExecuteSQLQueryRequest) GetExpression() *string  {
  return s.Expression
}

func (s *ExecuteSQLQueryRequest) GetMaxResults() *int32  {
  return s.MaxResults
}

func (s *ExecuteSQLQueryRequest) GetNextToken() *string  {
  return s.NextToken
}

func (s *ExecuteSQLQueryRequest) GetScope() *string  {
  return s.Scope
}

func (s *ExecuteSQLQueryRequest) SetExpression(v string) *ExecuteSQLQueryRequest {
  s.Expression = &v
  return s
}

func (s *ExecuteSQLQueryRequest) SetMaxResults(v int32) *ExecuteSQLQueryRequest {
  s.MaxResults = &v
  return s
}

func (s *ExecuteSQLQueryRequest) SetNextToken(v string) *ExecuteSQLQueryRequest {
  s.NextToken = &v
  return s
}

func (s *ExecuteSQLQueryRequest) SetScope(v string) *ExecuteSQLQueryRequest {
  s.Scope = &v
  return s
}

func (s *ExecuteSQLQueryRequest) Validate() error {
  return dara.Validate(s)
}

