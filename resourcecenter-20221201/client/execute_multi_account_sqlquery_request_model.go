// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteMultiAccountSQLQueryRequest interface {
  dara.Model
  String() string
  GoString() string
  SetExpression(v string) *ExecuteMultiAccountSQLQueryRequest
  GetExpression() *string 
  SetMaxResults(v int32) *ExecuteMultiAccountSQLQueryRequest
  GetMaxResults() *int32 
  SetNextToken(v string) *ExecuteMultiAccountSQLQueryRequest
  GetNextToken() *string 
  SetScope(v string) *ExecuteMultiAccountSQLQueryRequest
  GetScope() *string 
}

type ExecuteMultiAccountSQLQueryRequest struct {
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
  // The maximum number of entries to return on each page.
  // 
  // Valid values: 1 to 1000.
  // 
  // Default value: 1000.
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
  // The search scope. The value of this parameter can be one of the following items:
  // 
  // 	- ID of a resource directory: Resources within the management account and all members of the resource directory are searched.
  // 
  // 	- ID of the Root folder: Resources within all members in the Root folder and the subfolders of the Root folder are searched.
  // 
  // 	- ID of a folder: Resources within all members in the folder are searched.
  // 
  // 	- ID of a member: Resources within the member are searched.
  // 
  // 	- ID of a member/ID of a Resource group: Resources within the member in the resource group are searched.
  // 
  // For more information about how to obtain the ID of a resource directory, the Root folder, a folder, a member, or a resource group, see [GetResourceDirectory](https://help.aliyun.com/document_detail/159995.html), [ListFoldersForParent](https://help.aliyun.com/document_detail/159997.html), [ListFoldersForParent](https://help.aliyun.com/document_detail/159997.html), [ListAccounts](https://help.aliyun.com/document_detail/160016.html), or [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html).
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // rd-r4****
  Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s ExecuteMultiAccountSQLQueryRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteMultiAccountSQLQueryRequest) GoString() string {
  return s.String()
}

func (s *ExecuteMultiAccountSQLQueryRequest) GetExpression() *string  {
  return s.Expression
}

func (s *ExecuteMultiAccountSQLQueryRequest) GetMaxResults() *int32  {
  return s.MaxResults
}

func (s *ExecuteMultiAccountSQLQueryRequest) GetNextToken() *string  {
  return s.NextToken
}

func (s *ExecuteMultiAccountSQLQueryRequest) GetScope() *string  {
  return s.Scope
}

func (s *ExecuteMultiAccountSQLQueryRequest) SetExpression(v string) *ExecuteMultiAccountSQLQueryRequest {
  s.Expression = &v
  return s
}

func (s *ExecuteMultiAccountSQLQueryRequest) SetMaxResults(v int32) *ExecuteMultiAccountSQLQueryRequest {
  s.MaxResults = &v
  return s
}

func (s *ExecuteMultiAccountSQLQueryRequest) SetNextToken(v string) *ExecuteMultiAccountSQLQueryRequest {
  s.NextToken = &v
  return s
}

func (s *ExecuteMultiAccountSQLQueryRequest) SetScope(v string) *ExecuteMultiAccountSQLQueryRequest {
  s.Scope = &v
  return s
}

func (s *ExecuteMultiAccountSQLQueryRequest) Validate() error {
  return dara.Validate(s)
}

