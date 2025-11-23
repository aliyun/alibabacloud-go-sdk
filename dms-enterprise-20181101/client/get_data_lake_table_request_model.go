// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataLakeTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *GetDataLakeTableRequest
	GetCatalogName() *string
	SetDataRegion(v string) *GetDataLakeTableRequest
	GetDataRegion() *string
	SetDbName(v string) *GetDataLakeTableRequest
	GetDbName() *string
	SetName(v string) *GetDataLakeTableRequest
	GetName() *string
	SetTid(v int64) *GetDataLakeTableRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *GetDataLakeTableRequest
	GetWorkspaceId() *int64
}

type GetDataLakeTableRequest struct {
	// The name of the data catalog.
	//
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// hive
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// The region where the data lake resides.
	//
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// cn-hangzhou
	DataRegion *string `json:"DataRegion,omitempty" xml:"DataRegion,omitempty"`
	// The name of the database to which the table belongs.
	//
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// default
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The table name.
	//
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// 100g_customer
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The tenant ID. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 12****
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetDataLakeTableRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataLakeTableRequest) GoString() string {
	return s.String()
}

func (s *GetDataLakeTableRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *GetDataLakeTableRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *GetDataLakeTableRequest) GetDbName() *string {
	return s.DbName
}

func (s *GetDataLakeTableRequest) GetName() *string {
	return s.Name
}

func (s *GetDataLakeTableRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetDataLakeTableRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *GetDataLakeTableRequest) SetCatalogName(v string) *GetDataLakeTableRequest {
	s.CatalogName = &v
	return s
}

func (s *GetDataLakeTableRequest) SetDataRegion(v string) *GetDataLakeTableRequest {
	s.DataRegion = &v
	return s
}

func (s *GetDataLakeTableRequest) SetDbName(v string) *GetDataLakeTableRequest {
	s.DbName = &v
	return s
}

func (s *GetDataLakeTableRequest) SetName(v string) *GetDataLakeTableRequest {
	s.Name = &v
	return s
}

func (s *GetDataLakeTableRequest) SetTid(v int64) *GetDataLakeTableRequest {
	s.Tid = &v
	return s
}

func (s *GetDataLakeTableRequest) SetWorkspaceId(v int64) *GetDataLakeTableRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetDataLakeTableRequest) Validate() error {
	return dara.Validate(s)
}
