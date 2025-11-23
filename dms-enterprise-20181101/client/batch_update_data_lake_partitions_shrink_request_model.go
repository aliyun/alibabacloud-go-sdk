// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateDataLakePartitionsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *BatchUpdateDataLakePartitionsShrinkRequest
	GetCatalogName() *string
	SetDataRegion(v string) *BatchUpdateDataLakePartitionsShrinkRequest
	GetDataRegion() *string
	SetDbName(v string) *BatchUpdateDataLakePartitionsShrinkRequest
	GetDbName() *string
	SetPartitionInputsShrink(v string) *BatchUpdateDataLakePartitionsShrinkRequest
	GetPartitionInputsShrink() *string
	SetTableName(v string) *BatchUpdateDataLakePartitionsShrinkRequest
	GetTableName() *string
	SetTid(v int64) *BatchUpdateDataLakePartitionsShrinkRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *BatchUpdateDataLakePartitionsShrinkRequest
	GetWorkspaceId() *int64
}

type BatchUpdateDataLakePartitionsShrinkRequest struct {
	// The name of the data directory.
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
	// example:
	//
	// cn-hangzhou
	DataRegion *string `json:"DataRegion,omitempty" xml:"DataRegion,omitempty"`
	// The name of the database that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The information about the created partition.
	//
	// This parameter is required.
	PartitionInputsShrink *string `json:"PartitionInputs,omitempty" xml:"PartitionInputs,omitempty"`
	// The name of the table
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

func (s BatchUpdateDataLakePartitionsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateDataLakePartitionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateDataLakePartitionsShrinkRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *BatchUpdateDataLakePartitionsShrinkRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *BatchUpdateDataLakePartitionsShrinkRequest) GetDbName() *string {
	return s.DbName
}

func (s *BatchUpdateDataLakePartitionsShrinkRequest) GetPartitionInputsShrink() *string {
	return s.PartitionInputsShrink
}

func (s *BatchUpdateDataLakePartitionsShrinkRequest) GetTableName() *string {
	return s.TableName
}

func (s *BatchUpdateDataLakePartitionsShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *BatchUpdateDataLakePartitionsShrinkRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *BatchUpdateDataLakePartitionsShrinkRequest) SetCatalogName(v string) *BatchUpdateDataLakePartitionsShrinkRequest {
	s.CatalogName = &v
	return s
}

func (s *BatchUpdateDataLakePartitionsShrinkRequest) SetDataRegion(v string) *BatchUpdateDataLakePartitionsShrinkRequest {
	s.DataRegion = &v
	return s
}

func (s *BatchUpdateDataLakePartitionsShrinkRequest) SetDbName(v string) *BatchUpdateDataLakePartitionsShrinkRequest {
	s.DbName = &v
	return s
}

func (s *BatchUpdateDataLakePartitionsShrinkRequest) SetPartitionInputsShrink(v string) *BatchUpdateDataLakePartitionsShrinkRequest {
	s.PartitionInputsShrink = &v
	return s
}

func (s *BatchUpdateDataLakePartitionsShrinkRequest) SetTableName(v string) *BatchUpdateDataLakePartitionsShrinkRequest {
	s.TableName = &v
	return s
}

func (s *BatchUpdateDataLakePartitionsShrinkRequest) SetTid(v int64) *BatchUpdateDataLakePartitionsShrinkRequest {
	s.Tid = &v
	return s
}

func (s *BatchUpdateDataLakePartitionsShrinkRequest) SetWorkspaceId(v int64) *BatchUpdateDataLakePartitionsShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *BatchUpdateDataLakePartitionsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
