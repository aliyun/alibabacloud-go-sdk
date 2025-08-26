// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataLakeFunctionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *CreateDataLakeFunctionShrinkRequest
	GetCatalogName() *string
	SetDbName(v string) *CreateDataLakeFunctionShrinkRequest
	GetDbName() *string
	SetFunctionInputShrink(v string) *CreateDataLakeFunctionShrinkRequest
	GetFunctionInputShrink() *string
	SetTid(v int64) *CreateDataLakeFunctionShrinkRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *CreateDataLakeFunctionShrinkRequest
	GetWorkspaceId() *int64
}

type CreateDataLakeFunctionShrinkRequest struct {
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
	// example:
	//
	// 3****
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// example:
	//
	// 12****
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateDataLakeFunctionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataLakeFunctionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDataLakeFunctionShrinkRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *CreateDataLakeFunctionShrinkRequest) GetDbName() *string {
	return s.DbName
}

func (s *CreateDataLakeFunctionShrinkRequest) GetFunctionInputShrink() *string {
	return s.FunctionInputShrink
}

func (s *CreateDataLakeFunctionShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateDataLakeFunctionShrinkRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *CreateDataLakeFunctionShrinkRequest) SetCatalogName(v string) *CreateDataLakeFunctionShrinkRequest {
	s.CatalogName = &v
	return s
}

func (s *CreateDataLakeFunctionShrinkRequest) SetDbName(v string) *CreateDataLakeFunctionShrinkRequest {
	s.DbName = &v
	return s
}

func (s *CreateDataLakeFunctionShrinkRequest) SetFunctionInputShrink(v string) *CreateDataLakeFunctionShrinkRequest {
	s.FunctionInputShrink = &v
	return s
}

func (s *CreateDataLakeFunctionShrinkRequest) SetTid(v int64) *CreateDataLakeFunctionShrinkRequest {
	s.Tid = &v
	return s
}

func (s *CreateDataLakeFunctionShrinkRequest) SetWorkspaceId(v int64) *CreateDataLakeFunctionShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateDataLakeFunctionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
