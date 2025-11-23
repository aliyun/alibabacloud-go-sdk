// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakeTablebaseInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *ListDataLakeTablebaseInfoRequest
	GetCatalogName() *string
	SetDataRegion(v string) *ListDataLakeTablebaseInfoRequest
	GetDataRegion() *string
	SetDbName(v string) *ListDataLakeTablebaseInfoRequest
	GetDbName() *string
	SetPage(v int32) *ListDataLakeTablebaseInfoRequest
	GetPage() *int32
	SetRows(v int32) *ListDataLakeTablebaseInfoRequest
	GetRows() *int32
	SetSearchKey(v string) *ListDataLakeTablebaseInfoRequest
	GetSearchKey() *string
	SetTid(v int64) *ListDataLakeTablebaseInfoRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *ListDataLakeTablebaseInfoRequest
	GetWorkspaceId() *int64
}

type ListDataLakeTablebaseInfoRequest struct {
	// The name of the data catalog to query.
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
	// The name of the database to which the table belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	Rows *int32 `json:"Rows,omitempty" xml:"Rows,omitempty"`
	// The keyword that is used to search for tables.
	//
	// example:
	//
	// test
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The tenant ID. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) operation to query the tenant ID.
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

func (s ListDataLakeTablebaseInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeTablebaseInfoRequest) GoString() string {
	return s.String()
}

func (s *ListDataLakeTablebaseInfoRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *ListDataLakeTablebaseInfoRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *ListDataLakeTablebaseInfoRequest) GetDbName() *string {
	return s.DbName
}

func (s *ListDataLakeTablebaseInfoRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListDataLakeTablebaseInfoRequest) GetRows() *int32 {
	return s.Rows
}

func (s *ListDataLakeTablebaseInfoRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListDataLakeTablebaseInfoRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListDataLakeTablebaseInfoRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *ListDataLakeTablebaseInfoRequest) SetCatalogName(v string) *ListDataLakeTablebaseInfoRequest {
	s.CatalogName = &v
	return s
}

func (s *ListDataLakeTablebaseInfoRequest) SetDataRegion(v string) *ListDataLakeTablebaseInfoRequest {
	s.DataRegion = &v
	return s
}

func (s *ListDataLakeTablebaseInfoRequest) SetDbName(v string) *ListDataLakeTablebaseInfoRequest {
	s.DbName = &v
	return s
}

func (s *ListDataLakeTablebaseInfoRequest) SetPage(v int32) *ListDataLakeTablebaseInfoRequest {
	s.Page = &v
	return s
}

func (s *ListDataLakeTablebaseInfoRequest) SetRows(v int32) *ListDataLakeTablebaseInfoRequest {
	s.Rows = &v
	return s
}

func (s *ListDataLakeTablebaseInfoRequest) SetSearchKey(v string) *ListDataLakeTablebaseInfoRequest {
	s.SearchKey = &v
	return s
}

func (s *ListDataLakeTablebaseInfoRequest) SetTid(v int64) *ListDataLakeTablebaseInfoRequest {
	s.Tid = &v
	return s
}

func (s *ListDataLakeTablebaseInfoRequest) SetWorkspaceId(v int64) *ListDataLakeTablebaseInfoRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDataLakeTablebaseInfoRequest) Validate() error {
	return dara.Validate(s)
}
