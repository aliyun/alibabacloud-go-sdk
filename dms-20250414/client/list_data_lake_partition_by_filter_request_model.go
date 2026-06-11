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
	// Query conditional expression. Supported operators:
	//
	// - Comparison operators: =, <>, !=, <, <=, >, and >=. For example: \\`ds>20240101\\`.
	//
	// - Logical operators: AND, OR, and NOT. For example: \\`ds LIKE \\"20240%\\"\\`.
	//
	// - The BETWEEN operator, which specifies a range. For example: \\`ds BETWEEN 20240101 AND 20241201\\`.
	//
	// - The IN operator, which specifies a set of values. For example: \\`ds IN (20240101, 20240102)\\`.
	//
	// This parameter is required.
	//
	// example:
	//
	// ds>20241201
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The number of entries per page. The maximum value is 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results. Valid values:
	//
	// - If this parameter is left empty, no more results are returned.
	//
	// - If a value is returned, the value is the token for the next query.
	//
	// example:
	//
	// f056501ada12****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The name of the table.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
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

func (s ListDataLakePartitionByFilterRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakePartitionByFilterRequest) GoString() string {
	return s.String()
}

func (s *ListDataLakePartitionByFilterRequest) GetCatalogName() *string {
	return s.CatalogName
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
