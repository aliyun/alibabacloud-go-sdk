// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcubeUpgradePackageInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetMcubeUpgradePackageInfoRequest
	GetAppId() *string
	SetPackageId(v int64) *GetMcubeUpgradePackageInfoRequest
	GetPackageId() *int64
	SetTenantId(v string) *GetMcubeUpgradePackageInfoRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *GetMcubeUpgradePackageInfoRequest
	GetWorkspaceId() *string
}

type GetMcubeUpgradePackageInfoRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	PackageId   *int64  `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	TenantId    *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetMcubeUpgradePackageInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMcubeUpgradePackageInfoRequest) GoString() string {
	return s.String()
}

func (s *GetMcubeUpgradePackageInfoRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetMcubeUpgradePackageInfoRequest) GetPackageId() *int64 {
	return s.PackageId
}

func (s *GetMcubeUpgradePackageInfoRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetMcubeUpgradePackageInfoRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetMcubeUpgradePackageInfoRequest) SetAppId(v string) *GetMcubeUpgradePackageInfoRequest {
	s.AppId = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoRequest) SetPackageId(v int64) *GetMcubeUpgradePackageInfoRequest {
	s.PackageId = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoRequest) SetTenantId(v string) *GetMcubeUpgradePackageInfoRequest {
	s.TenantId = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoRequest) SetWorkspaceId(v string) *GetMcubeUpgradePackageInfoRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoRequest) Validate() error {
	return dara.Validate(s)
}
