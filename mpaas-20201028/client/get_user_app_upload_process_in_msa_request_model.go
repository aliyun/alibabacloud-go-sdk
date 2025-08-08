// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserAppUploadProcessInMsaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetUserAppUploadProcessInMsaRequest
	GetAppId() *string
	SetId(v int64) *GetUserAppUploadProcessInMsaRequest
	GetId() *int64
	SetTenantId(v string) *GetUserAppUploadProcessInMsaRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *GetUserAppUploadProcessInMsaRequest
	GetWorkspaceId() *string
}

type GetUserAppUploadProcessInMsaRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetUserAppUploadProcessInMsaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserAppUploadProcessInMsaRequest) GoString() string {
	return s.String()
}

func (s *GetUserAppUploadProcessInMsaRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetUserAppUploadProcessInMsaRequest) GetId() *int64 {
	return s.Id
}

func (s *GetUserAppUploadProcessInMsaRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetUserAppUploadProcessInMsaRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetUserAppUploadProcessInMsaRequest) SetAppId(v string) *GetUserAppUploadProcessInMsaRequest {
	s.AppId = &v
	return s
}

func (s *GetUserAppUploadProcessInMsaRequest) SetId(v int64) *GetUserAppUploadProcessInMsaRequest {
	s.Id = &v
	return s
}

func (s *GetUserAppUploadProcessInMsaRequest) SetTenantId(v string) *GetUserAppUploadProcessInMsaRequest {
	s.TenantId = &v
	return s
}

func (s *GetUserAppUploadProcessInMsaRequest) SetWorkspaceId(v string) *GetUserAppUploadProcessInMsaRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetUserAppUploadProcessInMsaRequest) Validate() error {
	return dara.Validate(s)
}
