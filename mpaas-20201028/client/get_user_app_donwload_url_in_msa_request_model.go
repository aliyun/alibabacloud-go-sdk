// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserAppDonwloadUrlInMsaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetUserAppDonwloadUrlInMsaRequest
	GetAppId() *string
	SetId(v int64) *GetUserAppDonwloadUrlInMsaRequest
	GetId() *int64
	SetTenantId(v string) *GetUserAppDonwloadUrlInMsaRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *GetUserAppDonwloadUrlInMsaRequest
	GetWorkspaceId() *string
}

type GetUserAppDonwloadUrlInMsaRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetUserAppDonwloadUrlInMsaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserAppDonwloadUrlInMsaRequest) GoString() string {
	return s.String()
}

func (s *GetUserAppDonwloadUrlInMsaRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetUserAppDonwloadUrlInMsaRequest) GetId() *int64 {
	return s.Id
}

func (s *GetUserAppDonwloadUrlInMsaRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetUserAppDonwloadUrlInMsaRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetUserAppDonwloadUrlInMsaRequest) SetAppId(v string) *GetUserAppDonwloadUrlInMsaRequest {
	s.AppId = &v
	return s
}

func (s *GetUserAppDonwloadUrlInMsaRequest) SetId(v int64) *GetUserAppDonwloadUrlInMsaRequest {
	s.Id = &v
	return s
}

func (s *GetUserAppDonwloadUrlInMsaRequest) SetTenantId(v string) *GetUserAppDonwloadUrlInMsaRequest {
	s.TenantId = &v
	return s
}

func (s *GetUserAppDonwloadUrlInMsaRequest) SetWorkspaceId(v string) *GetUserAppDonwloadUrlInMsaRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetUserAppDonwloadUrlInMsaRequest) Validate() error {
	return dara.Validate(s)
}
