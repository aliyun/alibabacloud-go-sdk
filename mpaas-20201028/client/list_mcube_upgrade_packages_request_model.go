// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcubeUpgradePackagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListMcubeUpgradePackagesRequest
	GetAppId() *string
	SetPageNum(v int32) *ListMcubeUpgradePackagesRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListMcubeUpgradePackagesRequest
	GetPageSize() *int32
	SetTenantId(v string) *ListMcubeUpgradePackagesRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *ListMcubeUpgradePackagesRequest
	GetWorkspaceId() *string
}

type ListMcubeUpgradePackagesRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	PageNum     *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TenantId    *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListMcubeUpgradePackagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeUpgradePackagesRequest) GoString() string {
	return s.String()
}

func (s *ListMcubeUpgradePackagesRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListMcubeUpgradePackagesRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListMcubeUpgradePackagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMcubeUpgradePackagesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListMcubeUpgradePackagesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListMcubeUpgradePackagesRequest) SetAppId(v string) *ListMcubeUpgradePackagesRequest {
	s.AppId = &v
	return s
}

func (s *ListMcubeUpgradePackagesRequest) SetPageNum(v int32) *ListMcubeUpgradePackagesRequest {
	s.PageNum = &v
	return s
}

func (s *ListMcubeUpgradePackagesRequest) SetPageSize(v int32) *ListMcubeUpgradePackagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListMcubeUpgradePackagesRequest) SetTenantId(v string) *ListMcubeUpgradePackagesRequest {
	s.TenantId = &v
	return s
}

func (s *ListMcubeUpgradePackagesRequest) SetWorkspaceId(v string) *ListMcubeUpgradePackagesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListMcubeUpgradePackagesRequest) Validate() error {
	return dara.Validate(s)
}
