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
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// hive
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// default
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// 100g_customer
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 3****
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
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
