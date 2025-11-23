// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataLakeTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *DeleteDataLakeTableRequest
	GetCatalogName() *string
	SetDataRegion(v string) *DeleteDataLakeTableRequest
	GetDataRegion() *string
	SetDbName(v string) *DeleteDataLakeTableRequest
	GetDbName() *string
	SetTableName(v string) *DeleteDataLakeTableRequest
	GetTableName() *string
	SetTid(v int64) *DeleteDataLakeTableRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *DeleteDataLakeTableRequest
	GetWorkspaceId() *int64
}

type DeleteDataLakeTableRequest struct {
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
	// The name of the table.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The ID of the tenant.
	//
	// > To view the ID of the tenant, go to the DMS console and move the pointer over the profile picture in the upper-right corner. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
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

func (s DeleteDataLakeTableRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataLakeTableRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataLakeTableRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *DeleteDataLakeTableRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *DeleteDataLakeTableRequest) GetDbName() *string {
	return s.DbName
}

func (s *DeleteDataLakeTableRequest) GetTableName() *string {
	return s.TableName
}

func (s *DeleteDataLakeTableRequest) GetTid() *int64 {
	return s.Tid
}

func (s *DeleteDataLakeTableRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *DeleteDataLakeTableRequest) SetCatalogName(v string) *DeleteDataLakeTableRequest {
	s.CatalogName = &v
	return s
}

func (s *DeleteDataLakeTableRequest) SetDataRegion(v string) *DeleteDataLakeTableRequest {
	s.DataRegion = &v
	return s
}

func (s *DeleteDataLakeTableRequest) SetDbName(v string) *DeleteDataLakeTableRequest {
	s.DbName = &v
	return s
}

func (s *DeleteDataLakeTableRequest) SetTableName(v string) *DeleteDataLakeTableRequest {
	s.TableName = &v
	return s
}

func (s *DeleteDataLakeTableRequest) SetTid(v int64) *DeleteDataLakeTableRequest {
	s.Tid = &v
	return s
}

func (s *DeleteDataLakeTableRequest) SetWorkspaceId(v int64) *DeleteDataLakeTableRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteDataLakeTableRequest) Validate() error {
	return dara.Validate(s)
}
