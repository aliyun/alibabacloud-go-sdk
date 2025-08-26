// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataLakeFunctionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *UpdateDataLakeFunctionShrinkRequest
	GetCatalogName() *string
	SetDbName(v string) *UpdateDataLakeFunctionShrinkRequest
	GetDbName() *string
	SetFunctionInputShrink(v string) *UpdateDataLakeFunctionShrinkRequest
	GetFunctionInputShrink() *string
	SetFunctionName(v string) *UpdateDataLakeFunctionShrinkRequest
	GetFunctionName() *string
	SetTid(v int64) *UpdateDataLakeFunctionShrinkRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *UpdateDataLakeFunctionShrinkRequest
	GetWorkspaceId() *int64
}

type UpdateDataLakeFunctionShrinkRequest struct {
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
	FunctionInputShrink *string `json:"FunctionInput,omitempty" xml:"FunctionInput,omitempty"`
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

func (s UpdateDataLakeFunctionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataLakeFunctionShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataLakeFunctionShrinkRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *UpdateDataLakeFunctionShrinkRequest) GetDbName() *string {
	return s.DbName
}

func (s *UpdateDataLakeFunctionShrinkRequest) GetFunctionInputShrink() *string {
	return s.FunctionInputShrink
}

func (s *UpdateDataLakeFunctionShrinkRequest) GetFunctionName() *string {
	return s.FunctionName
}

func (s *UpdateDataLakeFunctionShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateDataLakeFunctionShrinkRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *UpdateDataLakeFunctionShrinkRequest) SetCatalogName(v string) *UpdateDataLakeFunctionShrinkRequest {
	s.CatalogName = &v
	return s
}

func (s *UpdateDataLakeFunctionShrinkRequest) SetDbName(v string) *UpdateDataLakeFunctionShrinkRequest {
	s.DbName = &v
	return s
}

func (s *UpdateDataLakeFunctionShrinkRequest) SetFunctionInputShrink(v string) *UpdateDataLakeFunctionShrinkRequest {
	s.FunctionInputShrink = &v
	return s
}

func (s *UpdateDataLakeFunctionShrinkRequest) SetFunctionName(v string) *UpdateDataLakeFunctionShrinkRequest {
	s.FunctionName = &v
	return s
}

func (s *UpdateDataLakeFunctionShrinkRequest) SetTid(v int64) *UpdateDataLakeFunctionShrinkRequest {
	s.Tid = &v
	return s
}

func (s *UpdateDataLakeFunctionShrinkRequest) SetWorkspaceId(v int64) *UpdateDataLakeFunctionShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateDataLakeFunctionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
