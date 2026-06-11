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
	// The name of the data catalog.
	//
	// This parameter is required.
	//
	// example:
	//
	// hive
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The name of the function.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_funciton
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The tenant ID.
	//
	// > Hover over your profile picture in the upper-right corner of the DMS console to obtain the tenant ID. For details, see [View tenant information](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3****
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The workspace ID.
	//
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
