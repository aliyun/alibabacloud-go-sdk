// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataLakeDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *GetDataLakeDatabaseRequest
	GetCatalogName() *string
	SetName(v string) *GetDataLakeDatabaseRequest
	GetName() *string
	SetTid(v int64) *GetDataLakeDatabaseRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *GetDataLakeDatabaseRequest
	GetWorkspaceId() *int64
}

type GetDataLakeDatabaseRequest struct {
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

func (s GetDataLakeDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataLakeDatabaseRequest) GoString() string {
	return s.String()
}

func (s *GetDataLakeDatabaseRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *GetDataLakeDatabaseRequest) GetName() *string {
	return s.Name
}

func (s *GetDataLakeDatabaseRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetDataLakeDatabaseRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *GetDataLakeDatabaseRequest) SetCatalogName(v string) *GetDataLakeDatabaseRequest {
	s.CatalogName = &v
	return s
}

func (s *GetDataLakeDatabaseRequest) SetName(v string) *GetDataLakeDatabaseRequest {
	s.Name = &v
	return s
}

func (s *GetDataLakeDatabaseRequest) SetTid(v int64) *GetDataLakeDatabaseRequest {
	s.Tid = &v
	return s
}

func (s *GetDataLakeDatabaseRequest) SetWorkspaceId(v int64) *GetDataLakeDatabaseRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetDataLakeDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
