// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadUserAppToMsaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UploadUserAppToMsaRequest
	GetAppId() *string
	SetFileUrl(v string) *UploadUserAppToMsaRequest
	GetFileUrl() *string
	SetTenantId(v string) *UploadUserAppToMsaRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *UploadUserAppToMsaRequest
	GetWorkspaceId() *string
}

type UploadUserAppToMsaRequest struct {
	// This parameter is required.
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UploadUserAppToMsaRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadUserAppToMsaRequest) GoString() string {
	return s.String()
}

func (s *UploadUserAppToMsaRequest) GetAppId() *string {
	return s.AppId
}

func (s *UploadUserAppToMsaRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *UploadUserAppToMsaRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UploadUserAppToMsaRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UploadUserAppToMsaRequest) SetAppId(v string) *UploadUserAppToMsaRequest {
	s.AppId = &v
	return s
}

func (s *UploadUserAppToMsaRequest) SetFileUrl(v string) *UploadUserAppToMsaRequest {
	s.FileUrl = &v
	return s
}

func (s *UploadUserAppToMsaRequest) SetTenantId(v string) *UploadUserAppToMsaRequest {
	s.TenantId = &v
	return s
}

func (s *UploadUserAppToMsaRequest) SetWorkspaceId(v string) *UploadUserAppToMsaRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UploadUserAppToMsaRequest) Validate() error {
	return dara.Validate(s)
}
