// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakePartitionByFilterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *ListDataLakePartitionByFilterRequest
	GetCatalogName() *string
	SetDataRegion(v string) *ListDataLakePartitionByFilterRequest
	GetDataRegion() *string
	SetDbName(v string) *ListDataLakePartitionByFilterRequest
	GetDbName() *string
	SetFilter(v string) *ListDataLakePartitionByFilterRequest
	GetFilter() *string
	SetMaxResults(v int32) *ListDataLakePartitionByFilterRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataLakePartitionByFilterRequest
	GetNextToken() *string
	SetTableName(v string) *ListDataLakePartitionByFilterRequest
	GetTableName() *string
	SetTid(v int64) *ListDataLakePartitionByFilterRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *ListDataLakePartitionByFilterRequest
	GetWorkspaceId() *int64
}

type ListDataLakePartitionByFilterRequest struct {
	// The catalog name.
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
	// The expression of the query condition. The following operators are supported:
	//
	// 	- Comparison operators: =, <>, ! =, <, <=,>, and >=. Example: ds>20240101.
	//
	// 	- Logical operators: AND, OR, and NOT. Example: ds LIKE \\"20240%\\".
	//
	// 	- BETWEEN operator: Specifies a range. Example: ds BETWEEN 20240101 AND 20241201.
	//
	// 	- IN operator: Specifies a set of values. Example: ds IN (20240101, 20240102).
	//
	// This parameter is required.
	//
	// example:
	//
	// ds>20241201
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid value:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// f056501ada12c1cc
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The table name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The ID of the tenant.
	//
	// > To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
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

func (s ListDataLakePartitionByFilterRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakePartitionByFilterRequest) GoString() string {
	return s.String()
}

func (s *ListDataLakePartitionByFilterRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *ListDataLakePartitionByFilterRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *ListDataLakePartitionByFilterRequest) GetDbName() *string {
	return s.DbName
}

func (s *ListDataLakePartitionByFilterRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListDataLakePartitionByFilterRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataLakePartitionByFilterRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataLakePartitionByFilterRequest) GetTableName() *string {
	return s.TableName
}

func (s *ListDataLakePartitionByFilterRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListDataLakePartitionByFilterRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *ListDataLakePartitionByFilterRequest) SetCatalogName(v string) *ListDataLakePartitionByFilterRequest {
	s.CatalogName = &v
	return s
}

func (s *ListDataLakePartitionByFilterRequest) SetDataRegion(v string) *ListDataLakePartitionByFilterRequest {
	s.DataRegion = &v
	return s
}

func (s *ListDataLakePartitionByFilterRequest) SetDbName(v string) *ListDataLakePartitionByFilterRequest {
	s.DbName = &v
	return s
}

func (s *ListDataLakePartitionByFilterRequest) SetFilter(v string) *ListDataLakePartitionByFilterRequest {
	s.Filter = &v
	return s
}

func (s *ListDataLakePartitionByFilterRequest) SetMaxResults(v int32) *ListDataLakePartitionByFilterRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDataLakePartitionByFilterRequest) SetNextToken(v string) *ListDataLakePartitionByFilterRequest {
	s.NextToken = &v
	return s
}

func (s *ListDataLakePartitionByFilterRequest) SetTableName(v string) *ListDataLakePartitionByFilterRequest {
	s.TableName = &v
	return s
}

func (s *ListDataLakePartitionByFilterRequest) SetTid(v int64) *ListDataLakePartitionByFilterRequest {
	s.Tid = &v
	return s
}

func (s *ListDataLakePartitionByFilterRequest) SetWorkspaceId(v int64) *ListDataLakePartitionByFilterRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDataLakePartitionByFilterRequest) Validate() error {
	return dara.Validate(s)
}
