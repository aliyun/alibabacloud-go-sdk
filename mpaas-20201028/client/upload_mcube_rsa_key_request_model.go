// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadMcubeRsaKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UploadMcubeRsaKeyRequest
	GetAppId() *string
	SetFileUrl(v string) *UploadMcubeRsaKeyRequest
	GetFileUrl() *string
	SetOnexFlag(v bool) *UploadMcubeRsaKeyRequest
	GetOnexFlag() *bool
	SetTenantId(v string) *UploadMcubeRsaKeyRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *UploadMcubeRsaKeyRequest
	GetWorkspaceId() *string
}

type UploadMcubeRsaKeyRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// This parameter is required.
	OnexFlag *bool `json:"OnexFlag,omitempty" xml:"OnexFlag,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UploadMcubeRsaKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadMcubeRsaKeyRequest) GoString() string {
	return s.String()
}

func (s *UploadMcubeRsaKeyRequest) GetAppId() *string {
	return s.AppId
}

func (s *UploadMcubeRsaKeyRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *UploadMcubeRsaKeyRequest) GetOnexFlag() *bool {
	return s.OnexFlag
}

func (s *UploadMcubeRsaKeyRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UploadMcubeRsaKeyRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UploadMcubeRsaKeyRequest) SetAppId(v string) *UploadMcubeRsaKeyRequest {
	s.AppId = &v
	return s
}

func (s *UploadMcubeRsaKeyRequest) SetFileUrl(v string) *UploadMcubeRsaKeyRequest {
	s.FileUrl = &v
	return s
}

func (s *UploadMcubeRsaKeyRequest) SetOnexFlag(v bool) *UploadMcubeRsaKeyRequest {
	s.OnexFlag = &v
	return s
}

func (s *UploadMcubeRsaKeyRequest) SetTenantId(v string) *UploadMcubeRsaKeyRequest {
	s.TenantId = &v
	return s
}

func (s *UploadMcubeRsaKeyRequest) SetWorkspaceId(v string) *UploadMcubeRsaKeyRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UploadMcubeRsaKeyRequest) Validate() error {
	return dara.Validate(s)
}
