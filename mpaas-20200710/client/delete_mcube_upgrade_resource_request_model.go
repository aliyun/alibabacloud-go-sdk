// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcubeUpgradeResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteMcubeUpgradeResourceRequest
	GetAppId() *string
	SetId(v string) *DeleteMcubeUpgradeResourceRequest
	GetId() *string
	SetPlatform(v string) *DeleteMcubeUpgradeResourceRequest
	GetPlatform() *string
	SetTenantId(v string) *DeleteMcubeUpgradeResourceRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *DeleteMcubeUpgradeResourceRequest
	GetWorkspaceId() *string
}

type DeleteMcubeUpgradeResourceRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Platform    *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	TenantId    *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteMcubeUpgradeResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcubeUpgradeResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteMcubeUpgradeResourceRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteMcubeUpgradeResourceRequest) GetId() *string {
	return s.Id
}

func (s *DeleteMcubeUpgradeResourceRequest) GetPlatform() *string {
	return s.Platform
}

func (s *DeleteMcubeUpgradeResourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteMcubeUpgradeResourceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteMcubeUpgradeResourceRequest) SetAppId(v string) *DeleteMcubeUpgradeResourceRequest {
	s.AppId = &v
	return s
}

func (s *DeleteMcubeUpgradeResourceRequest) SetId(v string) *DeleteMcubeUpgradeResourceRequest {
	s.Id = &v
	return s
}

func (s *DeleteMcubeUpgradeResourceRequest) SetPlatform(v string) *DeleteMcubeUpgradeResourceRequest {
	s.Platform = &v
	return s
}

func (s *DeleteMcubeUpgradeResourceRequest) SetTenantId(v string) *DeleteMcubeUpgradeResourceRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteMcubeUpgradeResourceRequest) SetWorkspaceId(v string) *DeleteMcubeUpgradeResourceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteMcubeUpgradeResourceRequest) Validate() error {
	return dara.Validate(s)
}
