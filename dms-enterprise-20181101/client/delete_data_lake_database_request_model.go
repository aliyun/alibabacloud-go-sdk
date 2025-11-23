// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataLakeDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *DeleteDataLakeDatabaseRequest
	GetCatalogName() *string
	SetDataRegion(v string) *DeleteDataLakeDatabaseRequest
	GetDataRegion() *string
	SetDbName(v string) *DeleteDataLakeDatabaseRequest
	GetDbName() *string
	SetTid(v int64) *DeleteDataLakeDatabaseRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *DeleteDataLakeDatabaseRequest
	GetWorkspaceId() *int64
}

type DeleteDataLakeDatabaseRequest struct {
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

func (s DeleteDataLakeDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataLakeDatabaseRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataLakeDatabaseRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *DeleteDataLakeDatabaseRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *DeleteDataLakeDatabaseRequest) GetDbName() *string {
	return s.DbName
}

func (s *DeleteDataLakeDatabaseRequest) GetTid() *int64 {
	return s.Tid
}

func (s *DeleteDataLakeDatabaseRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *DeleteDataLakeDatabaseRequest) SetCatalogName(v string) *DeleteDataLakeDatabaseRequest {
	s.CatalogName = &v
	return s
}

func (s *DeleteDataLakeDatabaseRequest) SetDataRegion(v string) *DeleteDataLakeDatabaseRequest {
	s.DataRegion = &v
	return s
}

func (s *DeleteDataLakeDatabaseRequest) SetDbName(v string) *DeleteDataLakeDatabaseRequest {
	s.DbName = &v
	return s
}

func (s *DeleteDataLakeDatabaseRequest) SetTid(v int64) *DeleteDataLakeDatabaseRequest {
	s.Tid = &v
	return s
}

func (s *DeleteDataLakeDatabaseRequest) SetWorkspaceId(v int64) *DeleteDataLakeDatabaseRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteDataLakeDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
