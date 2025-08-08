// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcubeMiniPackagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListMcubeMiniPackagesRequest
	GetAppId() *string
	SetH5Id(v string) *ListMcubeMiniPackagesRequest
	GetH5Id() *string
	SetPackageTypes(v string) *ListMcubeMiniPackagesRequest
	GetPackageTypes() *string
	SetPageNum(v int32) *ListMcubeMiniPackagesRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListMcubeMiniPackagesRequest
	GetPageSize() *int32
	SetTenantId(v string) *ListMcubeMiniPackagesRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *ListMcubeMiniPackagesRequest
	GetWorkspaceId() *string
}

type ListMcubeMiniPackagesRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	H5Id *string `json:"H5Id,omitempty" xml:"H5Id,omitempty"`
	// This parameter is required.
	PackageTypes *string `json:"PackageTypes,omitempty" xml:"PackageTypes,omitempty"`
	PageNum      *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListMcubeMiniPackagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeMiniPackagesRequest) GoString() string {
	return s.String()
}

func (s *ListMcubeMiniPackagesRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListMcubeMiniPackagesRequest) GetH5Id() *string {
	return s.H5Id
}

func (s *ListMcubeMiniPackagesRequest) GetPackageTypes() *string {
	return s.PackageTypes
}

func (s *ListMcubeMiniPackagesRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListMcubeMiniPackagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMcubeMiniPackagesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListMcubeMiniPackagesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListMcubeMiniPackagesRequest) SetAppId(v string) *ListMcubeMiniPackagesRequest {
	s.AppId = &v
	return s
}

func (s *ListMcubeMiniPackagesRequest) SetH5Id(v string) *ListMcubeMiniPackagesRequest {
	s.H5Id = &v
	return s
}

func (s *ListMcubeMiniPackagesRequest) SetPackageTypes(v string) *ListMcubeMiniPackagesRequest {
	s.PackageTypes = &v
	return s
}

func (s *ListMcubeMiniPackagesRequest) SetPageNum(v int32) *ListMcubeMiniPackagesRequest {
	s.PageNum = &v
	return s
}

func (s *ListMcubeMiniPackagesRequest) SetPageSize(v int32) *ListMcubeMiniPackagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListMcubeMiniPackagesRequest) SetTenantId(v string) *ListMcubeMiniPackagesRequest {
	s.TenantId = &v
	return s
}

func (s *ListMcubeMiniPackagesRequest) SetWorkspaceId(v string) *ListMcubeMiniPackagesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListMcubeMiniPackagesRequest) Validate() error {
	return dara.Validate(s)
}
