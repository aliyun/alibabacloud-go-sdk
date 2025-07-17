// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteDataLakePartitionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *BatchDeleteDataLakePartitionsRequest
	GetCatalogName() *string
	SetDataRegion(v string) *BatchDeleteDataLakePartitionsRequest
	GetDataRegion() *string
	SetDbName(v string) *BatchDeleteDataLakePartitionsRequest
	GetDbName() *string
	SetIfExists(v bool) *BatchDeleteDataLakePartitionsRequest
	GetIfExists() *bool
	SetPartitionValuesList(v [][]*string) *BatchDeleteDataLakePartitionsRequest
	GetPartitionValuesList() [][]*string
	SetTableName(v string) *BatchDeleteDataLakePartitionsRequest
	GetTableName() *string
	SetTid(v int64) *BatchDeleteDataLakePartitionsRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *BatchDeleteDataLakePartitionsRequest
	GetWorkspaceId() *int64
}

type BatchDeleteDataLakePartitionsRequest struct {
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
	// true
	IfExists *bool `json:"IfExists,omitempty" xml:"IfExists,omitempty"`
	// This parameter is required.
	PartitionValuesList [][]*string `json:"PartitionValuesList,omitempty" xml:"PartitionValuesList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// table_name
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// example:
	//
	// 3****
	Tid         *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s BatchDeleteDataLakePartitionsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteDataLakePartitionsRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteDataLakePartitionsRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *BatchDeleteDataLakePartitionsRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *BatchDeleteDataLakePartitionsRequest) GetDbName() *string {
	return s.DbName
}

func (s *BatchDeleteDataLakePartitionsRequest) GetIfExists() *bool {
	return s.IfExists
}

func (s *BatchDeleteDataLakePartitionsRequest) GetPartitionValuesList() [][]*string {
	return s.PartitionValuesList
}

func (s *BatchDeleteDataLakePartitionsRequest) GetTableName() *string {
	return s.TableName
}

func (s *BatchDeleteDataLakePartitionsRequest) GetTid() *int64 {
	return s.Tid
}

func (s *BatchDeleteDataLakePartitionsRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *BatchDeleteDataLakePartitionsRequest) SetCatalogName(v string) *BatchDeleteDataLakePartitionsRequest {
	s.CatalogName = &v
	return s
}

func (s *BatchDeleteDataLakePartitionsRequest) SetDataRegion(v string) *BatchDeleteDataLakePartitionsRequest {
	s.DataRegion = &v
	return s
}

func (s *BatchDeleteDataLakePartitionsRequest) SetDbName(v string) *BatchDeleteDataLakePartitionsRequest {
	s.DbName = &v
	return s
}

func (s *BatchDeleteDataLakePartitionsRequest) SetIfExists(v bool) *BatchDeleteDataLakePartitionsRequest {
	s.IfExists = &v
	return s
}

func (s *BatchDeleteDataLakePartitionsRequest) SetPartitionValuesList(v [][]*string) *BatchDeleteDataLakePartitionsRequest {
	s.PartitionValuesList = v
	return s
}

func (s *BatchDeleteDataLakePartitionsRequest) SetTableName(v string) *BatchDeleteDataLakePartitionsRequest {
	s.TableName = &v
	return s
}

func (s *BatchDeleteDataLakePartitionsRequest) SetTid(v int64) *BatchDeleteDataLakePartitionsRequest {
	s.Tid = &v
	return s
}

func (s *BatchDeleteDataLakePartitionsRequest) SetWorkspaceId(v int64) *BatchDeleteDataLakePartitionsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *BatchDeleteDataLakePartitionsRequest) Validate() error {
	return dara.Validate(s)
}
