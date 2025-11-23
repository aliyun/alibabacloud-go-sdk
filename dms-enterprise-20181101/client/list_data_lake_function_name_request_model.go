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
	SetDataRegion(v string) *ListDataLakeFunctionNameRequest
	GetDataRegion() *string
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
	// The region where the data lake resides.
	//
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// cn-hangzhou
	DataRegion *string `json:"DataRegion,omitempty" xml:"DataRegion,omitempty"`
	// The database name.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The regular expression that is used to filter the returned function names.
	//
	// example:
	//
	// .*
	FunctionNamePattern *string `json:"FunctionNamePattern,omitempty" xml:"FunctionNamePattern,omitempty"`
	// The number of records per page. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, there is no next page.
	//
	// 	- If a value of **NextToken*	- is returned, it indicates the token that is used for the next query.
	//
	// example:
	//
	// f056501ada12c1cc
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the tenant.
	//
	// > To view the ID of the tenant, go to the DMS console and move the pointer over the profile picture in the upper-right corner. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The workspace ID.
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

func (s *ListDataLakeFunctionNameRequest) GetDataRegion() *string {
	return s.DataRegion
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

func (s *ListDataLakeFunctionNameRequest) SetDataRegion(v string) *ListDataLakeFunctionNameRequest {
	s.DataRegion = &v
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
