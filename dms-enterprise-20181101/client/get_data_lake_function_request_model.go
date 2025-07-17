// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataLakeFunctionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *GetDataLakeFunctionRequest
	GetCatalogName() *string
	SetDataRegion(v string) *GetDataLakeFunctionRequest
	GetDataRegion() *string
	SetDbName(v string) *GetDataLakeFunctionRequest
	GetDbName() *string
	SetFunctionName(v string) *GetDataLakeFunctionRequest
	GetFunctionName() *string
	SetTid(v int64) *GetDataLakeFunctionRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *GetDataLakeFunctionRequest
	GetWorkspaceId() *int64
}

type GetDataLakeFunctionRequest struct {
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
	// cn-hangzhou
	DataRegion *string `json:"DataRegion,omitempty" xml:"DataRegion,omitempty"`
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
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// example:
	//
	// 12****
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetDataLakeFunctionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataLakeFunctionRequest) GoString() string {
	return s.String()
}

func (s *GetDataLakeFunctionRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *GetDataLakeFunctionRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *GetDataLakeFunctionRequest) GetDbName() *string {
	return s.DbName
}

func (s *GetDataLakeFunctionRequest) GetFunctionName() *string {
	return s.FunctionName
}

func (s *GetDataLakeFunctionRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetDataLakeFunctionRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *GetDataLakeFunctionRequest) SetCatalogName(v string) *GetDataLakeFunctionRequest {
	s.CatalogName = &v
	return s
}

func (s *GetDataLakeFunctionRequest) SetDataRegion(v string) *GetDataLakeFunctionRequest {
	s.DataRegion = &v
	return s
}

func (s *GetDataLakeFunctionRequest) SetDbName(v string) *GetDataLakeFunctionRequest {
	s.DbName = &v
	return s
}

func (s *GetDataLakeFunctionRequest) SetFunctionName(v string) *GetDataLakeFunctionRequest {
	s.FunctionName = &v
	return s
}

func (s *GetDataLakeFunctionRequest) SetTid(v int64) *GetDataLakeFunctionRequest {
	s.Tid = &v
	return s
}

func (s *GetDataLakeFunctionRequest) SetWorkspaceId(v int64) *GetDataLakeFunctionRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetDataLakeFunctionRequest) Validate() error {
	return dara.Validate(s)
}
