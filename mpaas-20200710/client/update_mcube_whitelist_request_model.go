// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMcubeWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateMcubeWhitelistRequest
	GetAppId() *string
	SetId(v string) *UpdateMcubeWhitelistRequest
	GetId() *string
	SetKeyIds(v string) *UpdateMcubeWhitelistRequest
	GetKeyIds() *string
	SetOnexFlag(v bool) *UpdateMcubeWhitelistRequest
	GetOnexFlag() *bool
	SetOssUrl(v string) *UpdateMcubeWhitelistRequest
	GetOssUrl() *string
	SetTenantId(v string) *UpdateMcubeWhitelistRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *UpdateMcubeWhitelistRequest
	GetWorkspaceId() *string
}

type UpdateMcubeWhitelistRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	Id     *string `json:"Id,omitempty" xml:"Id,omitempty"`
	KeyIds *string `json:"KeyIds,omitempty" xml:"KeyIds,omitempty"`
	// This parameter is required.
	OnexFlag *bool   `json:"OnexFlag,omitempty" xml:"OnexFlag,omitempty"`
	OssUrl   *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateMcubeWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcubeWhitelistRequest) GoString() string {
	return s.String()
}

func (s *UpdateMcubeWhitelistRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateMcubeWhitelistRequest) GetId() *string {
	return s.Id
}

func (s *UpdateMcubeWhitelistRequest) GetKeyIds() *string {
	return s.KeyIds
}

func (s *UpdateMcubeWhitelistRequest) GetOnexFlag() *bool {
	return s.OnexFlag
}

func (s *UpdateMcubeWhitelistRequest) GetOssUrl() *string {
	return s.OssUrl
}

func (s *UpdateMcubeWhitelistRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateMcubeWhitelistRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateMcubeWhitelistRequest) SetAppId(v string) *UpdateMcubeWhitelistRequest {
	s.AppId = &v
	return s
}

func (s *UpdateMcubeWhitelistRequest) SetId(v string) *UpdateMcubeWhitelistRequest {
	s.Id = &v
	return s
}

func (s *UpdateMcubeWhitelistRequest) SetKeyIds(v string) *UpdateMcubeWhitelistRequest {
	s.KeyIds = &v
	return s
}

func (s *UpdateMcubeWhitelistRequest) SetOnexFlag(v bool) *UpdateMcubeWhitelistRequest {
	s.OnexFlag = &v
	return s
}

func (s *UpdateMcubeWhitelistRequest) SetOssUrl(v string) *UpdateMcubeWhitelistRequest {
	s.OssUrl = &v
	return s
}

func (s *UpdateMcubeWhitelistRequest) SetTenantId(v string) *UpdateMcubeWhitelistRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateMcubeWhitelistRequest) SetWorkspaceId(v string) *UpdateMcubeWhitelistRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateMcubeWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
