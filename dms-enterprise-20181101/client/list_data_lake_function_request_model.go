// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakeFunctionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *ListDataLakeFunctionRequest
	GetCatalogName() *string
	SetDataRegion(v string) *ListDataLakeFunctionRequest
	GetDataRegion() *string
	SetDbName(v string) *ListDataLakeFunctionRequest
	GetDbName() *string
	SetFunctionNamePattern(v string) *ListDataLakeFunctionRequest
	GetFunctionNamePattern() *string
	SetMaxResults(v int32) *ListDataLakeFunctionRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataLakeFunctionRequest
	GetNextToken() *string
	SetTid(v int64) *ListDataLakeFunctionRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *ListDataLakeFunctionRequest
	GetWorkspaceId() *int64
}

type ListDataLakeFunctionRequest struct {
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
	// The name of the database.
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

func (s ListDataLakeFunctionRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeFunctionRequest) GoString() string {
	return s.String()
}

func (s *ListDataLakeFunctionRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *ListDataLakeFunctionRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *ListDataLakeFunctionRequest) GetDbName() *string {
	return s.DbName
}

func (s *ListDataLakeFunctionRequest) GetFunctionNamePattern() *string {
	return s.FunctionNamePattern
}

func (s *ListDataLakeFunctionRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataLakeFunctionRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataLakeFunctionRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListDataLakeFunctionRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *ListDataLakeFunctionRequest) SetCatalogName(v string) *ListDataLakeFunctionRequest {
	s.CatalogName = &v
	return s
}

func (s *ListDataLakeFunctionRequest) SetDataRegion(v string) *ListDataLakeFunctionRequest {
	s.DataRegion = &v
	return s
}

func (s *ListDataLakeFunctionRequest) SetDbName(v string) *ListDataLakeFunctionRequest {
	s.DbName = &v
	return s
}

func (s *ListDataLakeFunctionRequest) SetFunctionNamePattern(v string) *ListDataLakeFunctionRequest {
	s.FunctionNamePattern = &v
	return s
}

func (s *ListDataLakeFunctionRequest) SetMaxResults(v int32) *ListDataLakeFunctionRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDataLakeFunctionRequest) SetNextToken(v string) *ListDataLakeFunctionRequest {
	s.NextToken = &v
	return s
}

func (s *ListDataLakeFunctionRequest) SetTid(v int64) *ListDataLakeFunctionRequest {
	s.Tid = &v
	return s
}

func (s *ListDataLakeFunctionRequest) SetWorkspaceId(v int64) *ListDataLakeFunctionRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDataLakeFunctionRequest) Validate() error {
	return dara.Validate(s)
}
