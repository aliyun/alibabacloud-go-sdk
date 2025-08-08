// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileTokenForUploadToMsaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetFileTokenForUploadToMsaRequest
	GetAppId() *string
	SetOnexFlag(v bool) *GetFileTokenForUploadToMsaRequest
	GetOnexFlag() *bool
	SetTenantId(v string) *GetFileTokenForUploadToMsaRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *GetFileTokenForUploadToMsaRequest
	GetWorkspaceId() *string
}

type GetFileTokenForUploadToMsaRequest struct {
	// This parameter is required.
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OnexFlag *bool   `json:"OnexFlag,omitempty" xml:"OnexFlag,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetFileTokenForUploadToMsaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFileTokenForUploadToMsaRequest) GoString() string {
	return s.String()
}

func (s *GetFileTokenForUploadToMsaRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetFileTokenForUploadToMsaRequest) GetOnexFlag() *bool {
	return s.OnexFlag
}

func (s *GetFileTokenForUploadToMsaRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetFileTokenForUploadToMsaRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetFileTokenForUploadToMsaRequest) SetAppId(v string) *GetFileTokenForUploadToMsaRequest {
	s.AppId = &v
	return s
}

func (s *GetFileTokenForUploadToMsaRequest) SetOnexFlag(v bool) *GetFileTokenForUploadToMsaRequest {
	s.OnexFlag = &v
	return s
}

func (s *GetFileTokenForUploadToMsaRequest) SetTenantId(v string) *GetFileTokenForUploadToMsaRequest {
	s.TenantId = &v
	return s
}

func (s *GetFileTokenForUploadToMsaRequest) SetWorkspaceId(v string) *GetFileTokenForUploadToMsaRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetFileTokenForUploadToMsaRequest) Validate() error {
	return dara.Validate(s)
}
