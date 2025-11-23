// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataLakeCatalogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *GetDataLakeCatalogRequest
	GetCatalogName() *string
	SetDataRegion(v string) *GetDataLakeCatalogRequest
	GetDataRegion() *string
	SetTid(v int64) *GetDataLakeCatalogRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *GetDataLakeCatalogRequest
	GetWorkspaceId() *int64
}

type GetDataLakeCatalogRequest struct {
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

func (s GetDataLakeCatalogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataLakeCatalogRequest) GoString() string {
	return s.String()
}

func (s *GetDataLakeCatalogRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *GetDataLakeCatalogRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *GetDataLakeCatalogRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetDataLakeCatalogRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *GetDataLakeCatalogRequest) SetCatalogName(v string) *GetDataLakeCatalogRequest {
	s.CatalogName = &v
	return s
}

func (s *GetDataLakeCatalogRequest) SetDataRegion(v string) *GetDataLakeCatalogRequest {
	s.DataRegion = &v
	return s
}

func (s *GetDataLakeCatalogRequest) SetTid(v int64) *GetDataLakeCatalogRequest {
	s.Tid = &v
	return s
}

func (s *GetDataLakeCatalogRequest) SetWorkspaceId(v int64) *GetDataLakeCatalogRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetDataLakeCatalogRequest) Validate() error {
	return dara.Validate(s)
}
