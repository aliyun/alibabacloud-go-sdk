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
	// The information about the table.
	//
	// This parameter is required.
	TableInputShrink *string `json:"TableInput,omitempty" xml:"TableInput,omitempty"`
	// The name of the table to update. If you do not want to change the table name, set this parameter to the same value as the Name parameter in TableInput.
	//
	// example:
	//
	// 100g_customer
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

func (s UpdateDataLakeTableShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataLakeTableShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataLakeTableShrinkRequest) GetCatalogName() *string {
	return s.CatalogName
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
