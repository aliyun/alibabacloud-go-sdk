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
	// This parameter is required.
	//
	// example:
	//
	// hive
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// default
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ds>20241201
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// f056501ada12****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// example:
	//
	// 3****
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
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
