// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakeTableNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *ListDataLakeTableNameRequest
	GetCatalogName() *string
	SetDataRegion(v string) *ListDataLakeTableNameRequest
	GetDataRegion() *string
	SetDbName(v string) *ListDataLakeTableNameRequest
	GetDbName() *string
	SetMaxResults(v int32) *ListDataLakeTableNameRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataLakeTableNameRequest
	GetNextToken() *string
	SetTableNamePattern(v string) *ListDataLakeTableNameRequest
	GetTableNamePattern() *string
	SetTableType(v string) *ListDataLakeTableNameRequest
	GetTableType() *string
	SetTid(v int64) *ListDataLakeTableNameRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *ListDataLakeTableNameRequest
	GetWorkspaceId() *int64
}

type ListDataLakeTableNameRequest struct {
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
	// 	- If a value of **NextToken*	- is returned, it indicates the token that is used for the next query.
	//
	// example:
	//
	// f056501ada12c1cc
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The regular expression that is used to match the table name.
	//
	// example:
	//
	// .*
	TableNamePattern *string `json:"TableNamePattern,omitempty" xml:"TableNamePattern,omitempty"`
	// The type of the data table. Valid values:
	//
	// 	- MANAGED_TABLE
	//
	// 	- EXTERNAL_TABLE
	//
	// 	- VIRTUAL_VIEW
	//
	// 	- INDEX_TABLE
	//
	// 	- MATERIALIZED_VIEW
	//
	// example:
	//
	// MANAGED_TABLE
	TableType *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
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

func (s ListDataLakeTableNameRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeTableNameRequest) GoString() string {
	return s.String()
}

func (s *ListDataLakeTableNameRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *ListDataLakeTableNameRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *ListDataLakeTableNameRequest) GetDbName() *string {
	return s.DbName
}

func (s *ListDataLakeTableNameRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataLakeTableNameRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataLakeTableNameRequest) GetTableNamePattern() *string {
	return s.TableNamePattern
}

func (s *ListDataLakeTableNameRequest) GetTableType() *string {
	return s.TableType
}

func (s *ListDataLakeTableNameRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListDataLakeTableNameRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *ListDataLakeTableNameRequest) SetCatalogName(v string) *ListDataLakeTableNameRequest {
	s.CatalogName = &v
	return s
}

func (s *ListDataLakeTableNameRequest) SetDataRegion(v string) *ListDataLakeTableNameRequest {
	s.DataRegion = &v
	return s
}

func (s *ListDataLakeTableNameRequest) SetDbName(v string) *ListDataLakeTableNameRequest {
	s.DbName = &v
	return s
}

func (s *ListDataLakeTableNameRequest) SetMaxResults(v int32) *ListDataLakeTableNameRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDataLakeTableNameRequest) SetNextToken(v string) *ListDataLakeTableNameRequest {
	s.NextToken = &v
	return s
}

func (s *ListDataLakeTableNameRequest) SetTableNamePattern(v string) *ListDataLakeTableNameRequest {
	s.TableNamePattern = &v
	return s
}

func (s *ListDataLakeTableNameRequest) SetTableType(v string) *ListDataLakeTableNameRequest {
	s.TableType = &v
	return s
}

func (s *ListDataLakeTableNameRequest) SetTid(v int64) *ListDataLakeTableNameRequest {
	s.Tid = &v
	return s
}

func (s *ListDataLakeTableNameRequest) SetWorkspaceId(v int64) *ListDataLakeTableNameRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDataLakeTableNameRequest) Validate() error {
	return dara.Validate(s)
}
