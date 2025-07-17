// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataLakePartitionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *DeleteDataLakePartitionShrinkRequest
	GetCatalogName() *string
	SetDataRegion(v string) *DeleteDataLakePartitionShrinkRequest
	GetDataRegion() *string
	SetDbName(v string) *DeleteDataLakePartitionShrinkRequest
	GetDbName() *string
	SetIfExists(v bool) *DeleteDataLakePartitionShrinkRequest
	GetIfExists() *bool
	SetPartitionValuesShrink(v string) *DeleteDataLakePartitionShrinkRequest
	GetPartitionValuesShrink() *string
	SetTableName(v string) *DeleteDataLakePartitionShrinkRequest
	GetTableName() *string
	SetTid(v int64) *DeleteDataLakePartitionShrinkRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *DeleteDataLakePartitionShrinkRequest
	GetWorkspaceId() *int64
}

type DeleteDataLakePartitionShrinkRequest struct {
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
	PartitionValuesShrink *string `json:"PartitionValues,omitempty" xml:"PartitionValues,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// table_name
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// example:
	//
	// 3***
	Tid         *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteDataLakePartitionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataLakePartitionShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataLakePartitionShrinkRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *DeleteDataLakePartitionShrinkRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *DeleteDataLakePartitionShrinkRequest) GetDbName() *string {
	return s.DbName
}

func (s *DeleteDataLakePartitionShrinkRequest) GetIfExists() *bool {
	return s.IfExists
}

func (s *DeleteDataLakePartitionShrinkRequest) GetPartitionValuesShrink() *string {
	return s.PartitionValuesShrink
}

func (s *DeleteDataLakePartitionShrinkRequest) GetTableName() *string {
	return s.TableName
}

func (s *DeleteDataLakePartitionShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *DeleteDataLakePartitionShrinkRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *DeleteDataLakePartitionShrinkRequest) SetCatalogName(v string) *DeleteDataLakePartitionShrinkRequest {
	s.CatalogName = &v
	return s
}

func (s *DeleteDataLakePartitionShrinkRequest) SetDataRegion(v string) *DeleteDataLakePartitionShrinkRequest {
	s.DataRegion = &v
	return s
}

func (s *DeleteDataLakePartitionShrinkRequest) SetDbName(v string) *DeleteDataLakePartitionShrinkRequest {
	s.DbName = &v
	return s
}

func (s *DeleteDataLakePartitionShrinkRequest) SetIfExists(v bool) *DeleteDataLakePartitionShrinkRequest {
	s.IfExists = &v
	return s
}

func (s *DeleteDataLakePartitionShrinkRequest) SetPartitionValuesShrink(v string) *DeleteDataLakePartitionShrinkRequest {
	s.PartitionValuesShrink = &v
	return s
}

func (s *DeleteDataLakePartitionShrinkRequest) SetTableName(v string) *DeleteDataLakePartitionShrinkRequest {
	s.TableName = &v
	return s
}

func (s *DeleteDataLakePartitionShrinkRequest) SetTid(v int64) *DeleteDataLakePartitionShrinkRequest {
	s.Tid = &v
	return s
}

func (s *DeleteDataLakePartitionShrinkRequest) SetWorkspaceId(v int64) *DeleteDataLakePartitionShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteDataLakePartitionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
