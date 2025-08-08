// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcubeUpgradeTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListMcubeUpgradeTasksRequest
	GetAppId() *string
	SetPackageId(v string) *ListMcubeUpgradeTasksRequest
	GetPackageId() *string
	SetTenantId(v string) *ListMcubeUpgradeTasksRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *ListMcubeUpgradeTasksRequest
	GetWorkspaceId() *string
}

type ListMcubeUpgradeTasksRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	PackageId   *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	TenantId    *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListMcubeUpgradeTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeUpgradeTasksRequest) GoString() string {
	return s.String()
}

func (s *ListMcubeUpgradeTasksRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListMcubeUpgradeTasksRequest) GetPackageId() *string {
	return s.PackageId
}

func (s *ListMcubeUpgradeTasksRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListMcubeUpgradeTasksRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListMcubeUpgradeTasksRequest) SetAppId(v string) *ListMcubeUpgradeTasksRequest {
	s.AppId = &v
	return s
}

func (s *ListMcubeUpgradeTasksRequest) SetPackageId(v string) *ListMcubeUpgradeTasksRequest {
	s.PackageId = &v
	return s
}

func (s *ListMcubeUpgradeTasksRequest) SetTenantId(v string) *ListMcubeUpgradeTasksRequest {
	s.TenantId = &v
	return s
}

func (s *ListMcubeUpgradeTasksRequest) SetWorkspaceId(v string) *ListMcubeUpgradeTasksRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListMcubeUpgradeTasksRequest) Validate() error {
	return dara.Validate(s)
}
