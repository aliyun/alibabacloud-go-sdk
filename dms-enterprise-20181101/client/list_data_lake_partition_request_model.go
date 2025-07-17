// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakePartitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *ListDataLakePartitionRequest
	GetCatalogName() *string
	SetDataRegion(v string) *ListDataLakePartitionRequest
	GetDataRegion() *string
	SetDbName(v string) *ListDataLakePartitionRequest
	GetDbName() *string
	SetMaxResults(v int32) *ListDataLakePartitionRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataLakePartitionRequest
	GetNextToken() *string
	SetPartNames(v []*string) *ListDataLakePartitionRequest
	GetPartNames() []*string
	SetTableName(v string) *ListDataLakePartitionRequest
	GetTableName() *string
	SetTid(v int64) *ListDataLakePartitionRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *ListDataLakePartitionRequest
	GetWorkspaceId() *int64
}

type ListDataLakePartitionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hive
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// cn-hangzhou
	DataRegion *string `json:"DataRegion,omitempty" xml:"DataRegion,omitempty"`
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
	// f056501ada12c1cc
	NextToken *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PartNames []*string `json:"PartNames,omitempty" xml:"PartNames,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// example:
	//
	// 3***
	Tid         *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDataLakePartitionRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakePartitionRequest) GoString() string {
	return s.String()
}

func (s *ListDataLakePartitionRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *ListDataLakePartitionRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *ListDataLakePartitionRequest) GetDbName() *string {
	return s.DbName
}

func (s *ListDataLakePartitionRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataLakePartitionRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataLakePartitionRequest) GetPartNames() []*string {
	return s.PartNames
}

func (s *ListDataLakePartitionRequest) GetTableName() *string {
	return s.TableName
}

func (s *ListDataLakePartitionRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListDataLakePartitionRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *ListDataLakePartitionRequest) SetCatalogName(v string) *ListDataLakePartitionRequest {
	s.CatalogName = &v
	return s
}

func (s *ListDataLakePartitionRequest) SetDataRegion(v string) *ListDataLakePartitionRequest {
	s.DataRegion = &v
	return s
}

func (s *ListDataLakePartitionRequest) SetDbName(v string) *ListDataLakePartitionRequest {
	s.DbName = &v
	return s
}

func (s *ListDataLakePartitionRequest) SetMaxResults(v int32) *ListDataLakePartitionRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDataLakePartitionRequest) SetNextToken(v string) *ListDataLakePartitionRequest {
	s.NextToken = &v
	return s
}

func (s *ListDataLakePartitionRequest) SetPartNames(v []*string) *ListDataLakePartitionRequest {
	s.PartNames = v
	return s
}

func (s *ListDataLakePartitionRequest) SetTableName(v string) *ListDataLakePartitionRequest {
	s.TableName = &v
	return s
}

func (s *ListDataLakePartitionRequest) SetTid(v int64) *ListDataLakePartitionRequest {
	s.Tid = &v
	return s
}

func (s *ListDataLakePartitionRequest) SetWorkspaceId(v int64) *ListDataLakePartitionRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDataLakePartitionRequest) Validate() error {
	return dara.Validate(s)
}
