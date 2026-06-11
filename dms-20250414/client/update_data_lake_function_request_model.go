// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataLakeFunctionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *UpdateDataLakeFunctionRequest
	GetCatalogName() *string
	SetDbName(v string) *UpdateDataLakeFunctionRequest
	GetDbName() *string
	SetFunctionInput(v *DLFunctionInput) *UpdateDataLakeFunctionRequest
	GetFunctionInput() *DLFunctionInput
	SetFunctionName(v string) *UpdateDataLakeFunctionRequest
	GetFunctionName() *string
	SetTid(v int64) *UpdateDataLakeFunctionRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *UpdateDataLakeFunctionRequest
	GetWorkspaceId() *int64
}

type UpdateDataLakeFunctionRequest struct {
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
	// The details of the function.
	//
	// This parameter is required.
	FunctionInput *DLFunctionInput `json:"FunctionInput,omitempty" xml:"FunctionInput,omitempty"`
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

func (s UpdateDataLakeFunctionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataLakeFunctionRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataLakeFunctionRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *UpdateDataLakeFunctionRequest) GetDbName() *string {
	return s.DbName
}

func (s *UpdateDataLakeFunctionRequest) GetFunctionInput() *DLFunctionInput {
	return s.FunctionInput
}

func (s *UpdateDataLakeFunctionRequest) GetFunctionName() *string {
	return s.FunctionName
}

func (s *UpdateDataLakeFunctionRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateDataLakeFunctionRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *UpdateDataLakeFunctionRequest) SetCatalogName(v string) *UpdateDataLakeFunctionRequest {
	s.CatalogName = &v
	return s
}

func (s *UpdateDataLakeFunctionRequest) SetDbName(v string) *UpdateDataLakeFunctionRequest {
	s.DbName = &v
	return s
}

func (s *UpdateDataLakeFunctionRequest) SetFunctionInput(v *DLFunctionInput) *UpdateDataLakeFunctionRequest {
	s.FunctionInput = v
	return s
}

func (s *UpdateDataLakeFunctionRequest) SetFunctionName(v string) *UpdateDataLakeFunctionRequest {
	s.FunctionName = &v
	return s
}

func (s *UpdateDataLakeFunctionRequest) SetTid(v int64) *UpdateDataLakeFunctionRequest {
	s.Tid = &v
	return s
}

func (s *UpdateDataLakeFunctionRequest) SetWorkspaceId(v int64) *UpdateDataLakeFunctionRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateDataLakeFunctionRequest) Validate() error {
	if s.FunctionInput != nil {
		if err := s.FunctionInput.Validate(); err != nil {
			return err
		}
	}
	return nil
}
