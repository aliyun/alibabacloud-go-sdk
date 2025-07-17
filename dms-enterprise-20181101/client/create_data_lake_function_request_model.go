// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataLakeFunctionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *CreateDataLakeFunctionRequest
	GetCatalogName() *string
	SetDataRegion(v string) *CreateDataLakeFunctionRequest
	GetDataRegion() *string
	SetDbName(v string) *CreateDataLakeFunctionRequest
	GetDbName() *string
	SetFunctionInput(v *DLFunctionInput) *CreateDataLakeFunctionRequest
	GetFunctionInput() *DLFunctionInput
	SetTid(v int64) *CreateDataLakeFunctionRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *CreateDataLakeFunctionRequest
	GetWorkspaceId() *int64
}

type CreateDataLakeFunctionRequest struct {
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
	FunctionInput *DLFunctionInput `json:"FunctionInput,omitempty" xml:"FunctionInput,omitempty"`
	// example:
	//
	// 3****
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// example:
	//
	// 12****
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateDataLakeFunctionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataLakeFunctionRequest) GoString() string {
	return s.String()
}

func (s *CreateDataLakeFunctionRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *CreateDataLakeFunctionRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *CreateDataLakeFunctionRequest) GetDbName() *string {
	return s.DbName
}

func (s *CreateDataLakeFunctionRequest) GetFunctionInput() *DLFunctionInput {
	return s.FunctionInput
}

func (s *CreateDataLakeFunctionRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateDataLakeFunctionRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *CreateDataLakeFunctionRequest) SetCatalogName(v string) *CreateDataLakeFunctionRequest {
	s.CatalogName = &v
	return s
}

func (s *CreateDataLakeFunctionRequest) SetDataRegion(v string) *CreateDataLakeFunctionRequest {
	s.DataRegion = &v
	return s
}

func (s *CreateDataLakeFunctionRequest) SetDbName(v string) *CreateDataLakeFunctionRequest {
	s.DbName = &v
	return s
}

func (s *CreateDataLakeFunctionRequest) SetFunctionInput(v *DLFunctionInput) *CreateDataLakeFunctionRequest {
	s.FunctionInput = v
	return s
}

func (s *CreateDataLakeFunctionRequest) SetTid(v int64) *CreateDataLakeFunctionRequest {
	s.Tid = &v
	return s
}

func (s *CreateDataLakeFunctionRequest) SetWorkspaceId(v int64) *CreateDataLakeFunctionRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateDataLakeFunctionRequest) Validate() error {
	return dara.Validate(s)
}
