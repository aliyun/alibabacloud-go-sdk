// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataLakePartitionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *UpdateDataLakePartitionShrinkRequest
	GetCatalogName() *string
	SetDataRegion(v string) *UpdateDataLakePartitionShrinkRequest
	GetDataRegion() *string
	SetDbName(v string) *UpdateDataLakePartitionShrinkRequest
	GetDbName() *string
	SetPartitionInputShrink(v string) *UpdateDataLakePartitionShrinkRequest
	GetPartitionInputShrink() *string
	SetTableName(v string) *UpdateDataLakePartitionShrinkRequest
	GetTableName() *string
	SetTid(v int64) *UpdateDataLakePartitionShrinkRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *UpdateDataLakePartitionShrinkRequest
	GetWorkspaceId() *int64
}

type UpdateDataLakePartitionShrinkRequest struct {
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
	// This parameter is required.
	PartitionInputShrink *string `json:"PartitionInput,omitempty" xml:"PartitionInput,omitempty"`
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

func (s UpdateDataLakePartitionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataLakePartitionShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataLakePartitionShrinkRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *UpdateDataLakePartitionShrinkRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *UpdateDataLakePartitionShrinkRequest) GetDbName() *string {
	return s.DbName
}

func (s *UpdateDataLakePartitionShrinkRequest) GetPartitionInputShrink() *string {
	return s.PartitionInputShrink
}

func (s *UpdateDataLakePartitionShrinkRequest) GetTableName() *string {
	return s.TableName
}

func (s *UpdateDataLakePartitionShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateDataLakePartitionShrinkRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *UpdateDataLakePartitionShrinkRequest) SetCatalogName(v string) *UpdateDataLakePartitionShrinkRequest {
	s.CatalogName = &v
	return s
}

func (s *UpdateDataLakePartitionShrinkRequest) SetDataRegion(v string) *UpdateDataLakePartitionShrinkRequest {
	s.DataRegion = &v
	return s
}

func (s *UpdateDataLakePartitionShrinkRequest) SetDbName(v string) *UpdateDataLakePartitionShrinkRequest {
	s.DbName = &v
	return s
}

func (s *UpdateDataLakePartitionShrinkRequest) SetPartitionInputShrink(v string) *UpdateDataLakePartitionShrinkRequest {
	s.PartitionInputShrink = &v
	return s
}

func (s *UpdateDataLakePartitionShrinkRequest) SetTableName(v string) *UpdateDataLakePartitionShrinkRequest {
	s.TableName = &v
	return s
}

func (s *UpdateDataLakePartitionShrinkRequest) SetTid(v int64) *UpdateDataLakePartitionShrinkRequest {
	s.Tid = &v
	return s
}

func (s *UpdateDataLakePartitionShrinkRequest) SetWorkspaceId(v int64) *UpdateDataLakePartitionShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateDataLakePartitionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
