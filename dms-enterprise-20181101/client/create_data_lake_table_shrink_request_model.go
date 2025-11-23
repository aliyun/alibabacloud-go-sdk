// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataLakeTableShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *CreateDataLakeTableShrinkRequest
	GetCatalogName() *string
	SetDataRegion(v string) *CreateDataLakeTableShrinkRequest
	GetDataRegion() *string
	SetDbName(v string) *CreateDataLakeTableShrinkRequest
	GetDbName() *string
	SetTableInputShrink(v string) *CreateDataLakeTableShrinkRequest
	GetTableInputShrink() *string
	SetTid(v int64) *CreateDataLakeTableShrinkRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *CreateDataLakeTableShrinkRequest
	GetWorkspaceId() *int64
}

type CreateDataLakeTableShrinkRequest struct {
	// The name of the data catalog.
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
	// The database name.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The information about the table.
	//
	// This parameter is required.
	TableInputShrink *string `json:"TableInput,omitempty" xml:"TableInput,omitempty"`
	// The ID of the tenant.
	//
	// > You can move the pointer over the profile picture in the upper-right corner of the DMS console to obtain the tenant ID.
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

func (s CreateDataLakeTableShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataLakeTableShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDataLakeTableShrinkRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *CreateDataLakeTableShrinkRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *CreateDataLakeTableShrinkRequest) GetDbName() *string {
	return s.DbName
}

func (s *CreateDataLakeTableShrinkRequest) GetTableInputShrink() *string {
	return s.TableInputShrink
}

func (s *CreateDataLakeTableShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateDataLakeTableShrinkRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *CreateDataLakeTableShrinkRequest) SetCatalogName(v string) *CreateDataLakeTableShrinkRequest {
	s.CatalogName = &v
	return s
}

func (s *CreateDataLakeTableShrinkRequest) SetDataRegion(v string) *CreateDataLakeTableShrinkRequest {
	s.DataRegion = &v
	return s
}

func (s *CreateDataLakeTableShrinkRequest) SetDbName(v string) *CreateDataLakeTableShrinkRequest {
	s.DbName = &v
	return s
}

func (s *CreateDataLakeTableShrinkRequest) SetTableInputShrink(v string) *CreateDataLakeTableShrinkRequest {
	s.TableInputShrink = &v
	return s
}

func (s *CreateDataLakeTableShrinkRequest) SetTid(v int64) *CreateDataLakeTableShrinkRequest {
	s.Tid = &v
	return s
}

func (s *CreateDataLakeTableShrinkRequest) SetWorkspaceId(v int64) *CreateDataLakeTableShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateDataLakeTableShrinkRequest) Validate() error {
	return dara.Validate(s)
}
