// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthTokenToMsenceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetAuthTokenToMsenceRequest
	GetAppId() *string
	SetAuthCode(v string) *GetAuthTokenToMsenceRequest
	GetAuthCode() *string
	SetMiniProgramId(v string) *GetAuthTokenToMsenceRequest
	GetMiniProgramId() *string
	SetPlatformId(v string) *GetAuthTokenToMsenceRequest
	GetPlatformId() *string
	SetTenantId(v string) *GetAuthTokenToMsenceRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *GetAuthTokenToMsenceRequest
	GetWorkspaceId() *string
}

type GetAuthTokenToMsenceRequest struct {
	// example:
	//
	// ALIPUBE5C3F6D091419
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// KUAj7XWYYf
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// example:
	//
	// 123321
	MiniProgramId *string `json:"MiniProgramId,omitempty" xml:"MiniProgramId,omitempty"`
	// example:
	//
	// mPaaS_Goosefish
	PlatformId *string `json:"PlatformId,omitempty" xml:"PlatformId,omitempty"`
	// example:
	//
	// NPHTGKNR
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// default
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetAuthTokenToMsenceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAuthTokenToMsenceRequest) GoString() string {
	return s.String()
}

func (s *GetAuthTokenToMsenceRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetAuthTokenToMsenceRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *GetAuthTokenToMsenceRequest) GetMiniProgramId() *string {
	return s.MiniProgramId
}

func (s *GetAuthTokenToMsenceRequest) GetPlatformId() *string {
	return s.PlatformId
}

func (s *GetAuthTokenToMsenceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetAuthTokenToMsenceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetAuthTokenToMsenceRequest) SetAppId(v string) *GetAuthTokenToMsenceRequest {
	s.AppId = &v
	return s
}

func (s *GetAuthTokenToMsenceRequest) SetAuthCode(v string) *GetAuthTokenToMsenceRequest {
	s.AuthCode = &v
	return s
}

func (s *GetAuthTokenToMsenceRequest) SetMiniProgramId(v string) *GetAuthTokenToMsenceRequest {
	s.MiniProgramId = &v
	return s
}

func (s *GetAuthTokenToMsenceRequest) SetPlatformId(v string) *GetAuthTokenToMsenceRequest {
	s.PlatformId = &v
	return s
}

func (s *GetAuthTokenToMsenceRequest) SetTenantId(v string) *GetAuthTokenToMsenceRequest {
	s.TenantId = &v
	return s
}

func (s *GetAuthTokenToMsenceRequest) SetWorkspaceId(v string) *GetAuthTokenToMsenceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetAuthTokenToMsenceRequest) Validate() error {
	return dara.Validate(s)
}
