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
	// This parameter is required.
	//
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// example:
	//
	// 3****
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
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
