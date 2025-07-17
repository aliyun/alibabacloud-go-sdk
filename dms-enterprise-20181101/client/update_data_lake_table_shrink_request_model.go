// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataLakeTableShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *UpdateDataLakeTableShrinkRequest
	GetCatalogName() *string
	SetDataRegion(v string) *UpdateDataLakeTableShrinkRequest
	GetDataRegion() *string
	SetDbName(v string) *UpdateDataLakeTableShrinkRequest
	GetDbName() *string
	SetTableInputShrink(v string) *UpdateDataLakeTableShrinkRequest
	GetTableInputShrink() *string
	SetTableName(v string) *UpdateDataLakeTableShrinkRequest
	GetTableName() *string
	SetTid(v int64) *UpdateDataLakeTableShrinkRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *UpdateDataLakeTableShrinkRequest
	GetWorkspaceId() *int64
}

type UpdateDataLakeTableShrinkRequest struct {
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
	TableInputShrink *string `json:"TableInput,omitempty" xml:"TableInput,omitempty"`
	TableName        *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// example:
	//
	// 3***
	Tid         *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateDataLakeTableShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataLakeTableShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataLakeTableShrinkRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *UpdateDataLakeTableShrinkRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *UpdateDataLakeTableShrinkRequest) GetDbName() *string {
	return s.DbName
}

func (s *UpdateDataLakeTableShrinkRequest) GetTableInputShrink() *string {
	return s.TableInputShrink
}

func (s *UpdateDataLakeTableShrinkRequest) GetTableName() *string {
	return s.TableName
}

func (s *UpdateDataLakeTableShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateDataLakeTableShrinkRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *UpdateDataLakeTableShrinkRequest) SetCatalogName(v string) *UpdateDataLakeTableShrinkRequest {
	s.CatalogName = &v
	return s
}

func (s *UpdateDataLakeTableShrinkRequest) SetDataRegion(v string) *UpdateDataLakeTableShrinkRequest {
	s.DataRegion = &v
	return s
}

func (s *UpdateDataLakeTableShrinkRequest) SetDbName(v string) *UpdateDataLakeTableShrinkRequest {
	s.DbName = &v
	return s
}

func (s *UpdateDataLakeTableShrinkRequest) SetTableInputShrink(v string) *UpdateDataLakeTableShrinkRequest {
	s.TableInputShrink = &v
	return s
}

func (s *UpdateDataLakeTableShrinkRequest) SetTableName(v string) *UpdateDataLakeTableShrinkRequest {
	s.TableName = &v
	return s
}

func (s *UpdateDataLakeTableShrinkRequest) SetTid(v int64) *UpdateDataLakeTableShrinkRequest {
	s.Tid = &v
	return s
}

func (s *UpdateDataLakeTableShrinkRequest) SetWorkspaceId(v int64) *UpdateDataLakeTableShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateDataLakeTableShrinkRequest) Validate() error {
	return dara.Validate(s)
}
