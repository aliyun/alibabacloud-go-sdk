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
	// table_name
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

func (s ListDataLakePartitionNameRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakePartitionNameRequest) GoString() string {
	return s.String()
}

func (s *ListDataLakePartitionNameRequest) GetCatalogName() *string {
	return s.CatalogName
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
