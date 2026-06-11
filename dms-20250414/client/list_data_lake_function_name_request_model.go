// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakeFunctionNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *ListDataLakeFunctionNameRequest
	GetCatalogName() *string
	SetDbName(v string) *ListDataLakeFunctionNameRequest
	GetDbName() *string
	SetFunctionNamePattern(v string) *ListDataLakeFunctionNameRequest
	GetFunctionNamePattern() *string
	SetMaxResults(v int32) *ListDataLakeFunctionNameRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataLakeFunctionNameRequest
	GetNextToken() *string
	SetTid(v int64) *ListDataLakeFunctionNameRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *ListDataLakeFunctionNameRequest
	GetWorkspaceId() *int64
}

type ListDataLakeFunctionNameRequest struct {
	// The name of the data catalog.
	//
	// This parameter is required.
	//
	// example:
	//
	// hive
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// A regular expression used to filter function names.
	//
	// example:
	//
	// .*
	FunctionNamePattern *string `json:"FunctionNamePattern,omitempty" xml:"FunctionNamePattern,omitempty"`
	// The page size. The maximum value is 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to page through results. Set this parameter to the \\`NextToken\\` value that is returned in the last response to retrieve the next page of results. You do not need to specify this parameter for the first request.
	//
	// - If **NextToken*	- is empty, there is no subsequent query.
	//
	// - If **NextToken*	- returns a value, it is the token for the next query.
	//
	// example:
	//
	// f056501ada12****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The tenant ID.
	//
	// > Hover over your profile picture in the upper-right corner of the DMS console to obtain the tenant ID. For details, see [View tenant information](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3****
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The ID of the workspace.
	//
	// example:
	//
	// 12****
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDataLakeFunctionNameRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeFunctionNameRequest) GoString() string {
	return s.String()
}

func (s *ListDataLakeFunctionNameRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *ListDataLakeFunctionNameRequest) GetDbName() *string {
	return s.DbName
}

func (s *ListDataLakeFunctionNameRequest) GetFunctionNamePattern() *string {
	return s.FunctionNamePattern
}

func (s *ListDataLakeFunctionNameRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataLakeFunctionNameRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataLakeFunctionNameRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListDataLakeFunctionNameRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *ListDataLakeFunctionNameRequest) SetCatalogName(v string) *ListDataLakeFunctionNameRequest {
	s.CatalogName = &v
	return s
}

func (s *ListDataLakeFunctionNameRequest) SetDbName(v string) *ListDataLakeFunctionNameRequest {
	s.DbName = &v
	return s
}

func (s *ListDataLakeFunctionNameRequest) SetFunctionNamePattern(v string) *ListDataLakeFunctionNameRequest {
	s.FunctionNamePattern = &v
	return s
}

func (s *ListDataLakeFunctionNameRequest) SetMaxResults(v int32) *ListDataLakeFunctionNameRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDataLakeFunctionNameRequest) SetNextToken(v string) *ListDataLakeFunctionNameRequest {
	s.NextToken = &v
	return s
}

func (s *ListDataLakeFunctionNameRequest) SetTid(v int64) *ListDataLakeFunctionNameRequest {
	s.Tid = &v
	return s
}

func (s *ListDataLakeFunctionNameRequest) SetWorkspaceId(v int64) *ListDataLakeFunctionNameRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDataLakeFunctionNameRequest) Validate() error {
	return dara.Validate(s)
}
