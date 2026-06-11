// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataLakePartitionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *GetDataLakePartitionShrinkRequest
	GetCatalogName() *string
	SetDbName(v string) *GetDataLakePartitionShrinkRequest
	GetDbName() *string
	SetPartitionValuesShrink(v string) *GetDataLakePartitionShrinkRequest
	GetPartitionValuesShrink() *string
	SetTableName(v string) *GetDataLakePartitionShrinkRequest
	GetTableName() *string
	SetTid(v int64) *GetDataLakePartitionShrinkRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *GetDataLakePartitionShrinkRequest
	GetWorkspaceId() *int64
}

type GetDataLakePartitionShrinkRequest struct {
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
	// A list of partition values.
	//
	// This parameter is required.
	PartitionValuesShrink *string `json:"PartitionValues,omitempty" xml:"PartitionValues,omitempty"`
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
	// The workspace ID.
	//
	// example:
	//
	// 12****
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetDataLakePartitionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataLakePartitionShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDataLakePartitionShrinkRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *GetDataLakePartitionShrinkRequest) GetDbName() *string {
	return s.DbName
}

func (s *GetDataLakePartitionShrinkRequest) GetPartitionValuesShrink() *string {
	return s.PartitionValuesShrink
}

func (s *GetDataLakePartitionShrinkRequest) GetTableName() *string {
	return s.TableName
}

func (s *GetDataLakePartitionShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetDataLakePartitionShrinkRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *GetDataLakePartitionShrinkRequest) SetCatalogName(v string) *GetDataLakePartitionShrinkRequest {
	s.CatalogName = &v
	return s
}

func (s *GetDataLakePartitionShrinkRequest) SetDbName(v string) *GetDataLakePartitionShrinkRequest {
	s.DbName = &v
	return s
}

func (s *GetDataLakePartitionShrinkRequest) SetPartitionValuesShrink(v string) *GetDataLakePartitionShrinkRequest {
	s.PartitionValuesShrink = &v
	return s
}

func (s *GetDataLakePartitionShrinkRequest) SetTableName(v string) *GetDataLakePartitionShrinkRequest {
	s.TableName = &v
	return s
}

func (s *GetDataLakePartitionShrinkRequest) SetTid(v int64) *GetDataLakePartitionShrinkRequest {
	s.Tid = &v
	return s
}

func (s *GetDataLakePartitionShrinkRequest) SetWorkspaceId(v int64) *GetDataLakePartitionShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetDataLakePartitionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
