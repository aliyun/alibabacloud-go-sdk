// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataLakeFunctionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *DeleteDataLakeFunctionRequest
	GetCatalogName() *string
	SetDbName(v string) *DeleteDataLakeFunctionRequest
	GetDbName() *string
	SetFunctionName(v string) *DeleteDataLakeFunctionRequest
	GetFunctionName() *string
	SetTid(v int64) *DeleteDataLakeFunctionRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *DeleteDataLakeFunctionRequest
	GetWorkspaceId() *int64
}

type DeleteDataLakeFunctionRequest struct {
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
	// my_funciton
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// example:
	//
	// 3****
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// example:
	//
	// 12****
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteDataLakeFunctionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataLakeFunctionRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataLakeFunctionRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *DeleteDataLakeFunctionRequest) GetDbName() *string {
	return s.DbName
}

func (s *DeleteDataLakeFunctionRequest) GetFunctionName() *string {
	return s.FunctionName
}

func (s *DeleteDataLakeFunctionRequest) GetTid() *int64 {
	return s.Tid
}

func (s *DeleteDataLakeFunctionRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *DeleteDataLakeFunctionRequest) SetCatalogName(v string) *DeleteDataLakeFunctionRequest {
	s.CatalogName = &v
	return s
}

func (s *DeleteDataLakeFunctionRequest) SetDbName(v string) *DeleteDataLakeFunctionRequest {
	s.DbName = &v
	return s
}

func (s *DeleteDataLakeFunctionRequest) SetFunctionName(v string) *DeleteDataLakeFunctionRequest {
	s.FunctionName = &v
	return s
}

func (s *DeleteDataLakeFunctionRequest) SetTid(v int64) *DeleteDataLakeFunctionRequest {
	s.Tid = &v
	return s
}

func (s *DeleteDataLakeFunctionRequest) SetWorkspaceId(v int64) *DeleteDataLakeFunctionRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteDataLakeFunctionRequest) Validate() error {
	return dara.Validate(s)
}
