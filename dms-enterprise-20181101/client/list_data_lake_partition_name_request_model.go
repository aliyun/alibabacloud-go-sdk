// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakePartitionNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *ListDataLakePartitionNameRequest
	GetCatalogName() *string
	SetDataRegion(v string) *ListDataLakePartitionNameRequest
	GetDataRegion() *string
	SetDbName(v string) *ListDataLakePartitionNameRequest
	GetDbName() *string
	SetMaxResults(v int32) *ListDataLakePartitionNameRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataLakePartitionNameRequest
	GetNextToken() *string
	SetTableName(v string) *ListDataLakePartitionNameRequest
	GetTableName() *string
	SetTid(v int64) *ListDataLakePartitionNameRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *ListDataLakePartitionNameRequest
	GetWorkspaceId() *int64
}

type ListDataLakePartitionNameRequest struct {
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
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// f056501ada12c1cc
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The name of the table.
	//
	// This parameter is required.
	//
	// example:
	//
	// table_name
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
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

func (s ListDataLakePartitionNameRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakePartitionNameRequest) GoString() string {
	return s.String()
}

func (s *ListDataLakePartitionNameRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *ListDataLakePartitionNameRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *ListDataLakePartitionNameRequest) GetDbName() *string {
	return s.DbName
}

func (s *ListDataLakePartitionNameRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataLakePartitionNameRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataLakePartitionNameRequest) GetTableName() *string {
	return s.TableName
}

func (s *ListDataLakePartitionNameRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListDataLakePartitionNameRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *ListDataLakePartitionNameRequest) SetCatalogName(v string) *ListDataLakePartitionNameRequest {
	s.CatalogName = &v
	return s
}

func (s *ListDataLakePartitionNameRequest) SetDataRegion(v string) *ListDataLakePartitionNameRequest {
	s.DataRegion = &v
	return s
}

func (s *ListDataLakePartitionNameRequest) SetDbName(v string) *ListDataLakePartitionNameRequest {
	s.DbName = &v
	return s
}

func (s *ListDataLakePartitionNameRequest) SetMaxResults(v int32) *ListDataLakePartitionNameRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDataLakePartitionNameRequest) SetNextToken(v string) *ListDataLakePartitionNameRequest {
	s.NextToken = &v
	return s
}

func (s *ListDataLakePartitionNameRequest) SetTableName(v string) *ListDataLakePartitionNameRequest {
	s.TableName = &v
	return s
}

func (s *ListDataLakePartitionNameRequest) SetTid(v int64) *ListDataLakePartitionNameRequest {
	s.Tid = &v
	return s
}

func (s *ListDataLakePartitionNameRequest) SetWorkspaceId(v int64) *ListDataLakePartitionNameRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDataLakePartitionNameRequest) Validate() error {
	return dara.Validate(s)
}
