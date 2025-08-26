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
	SetDbName(v string) *DeleteDataLakeDatabaseRequest
	GetDbName() *string
	SetTid(v int64) *DeleteDataLakeDatabaseRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *DeleteDataLakeDatabaseRequest
	GetWorkspaceId() *int64
}

type DeleteDataLakeDatabaseRequest struct {
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
	// default
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// example:
	//
	// 3****
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
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
