// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserInfoToMsenceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QueryUserInfoToMsenceRequest
	GetAppId() *string
	SetAuthToken(v string) *QueryUserInfoToMsenceRequest
	GetAuthToken() *string
	SetMiniProgramId(v string) *QueryUserInfoToMsenceRequest
	GetMiniProgramId() *string
	SetPlatformId(v string) *QueryUserInfoToMsenceRequest
	GetPlatformId() *string
	SetTenantId(v string) *QueryUserInfoToMsenceRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *QueryUserInfoToMsenceRequest
	GetWorkspaceId() *string
}

type QueryUserInfoToMsenceRequest struct {
	// example:
	//
	// ALIPUBE5C3F6D091419
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// NPHTGKNR
	AuthToken *string `json:"AuthToken,omitempty" xml:"AuthToken,omitempty"`
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
	// JGBGDUWU
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// preProd
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryUserInfoToMsenceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryUserInfoToMsenceRequest) GoString() string {
	return s.String()
}

func (s *QueryUserInfoToMsenceRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryUserInfoToMsenceRequest) GetAuthToken() *string {
	return s.AuthToken
}

func (s *QueryUserInfoToMsenceRequest) GetMiniProgramId() *string {
	return s.MiniProgramId
}

func (s *QueryUserInfoToMsenceRequest) GetPlatformId() *string {
	return s.PlatformId
}

func (s *QueryUserInfoToMsenceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryUserInfoToMsenceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryUserInfoToMsenceRequest) SetAppId(v string) *QueryUserInfoToMsenceRequest {
	s.AppId = &v
	return s
}

func (s *QueryUserInfoToMsenceRequest) SetAuthToken(v string) *QueryUserInfoToMsenceRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryUserInfoToMsenceRequest) SetMiniProgramId(v string) *QueryUserInfoToMsenceRequest {
	s.MiniProgramId = &v
	return s
}

func (s *QueryUserInfoToMsenceRequest) SetPlatformId(v string) *QueryUserInfoToMsenceRequest {
	s.PlatformId = &v
	return s
}

func (s *QueryUserInfoToMsenceRequest) SetTenantId(v string) *QueryUserInfoToMsenceRequest {
	s.TenantId = &v
	return s
}

func (s *QueryUserInfoToMsenceRequest) SetWorkspaceId(v string) *QueryUserInfoToMsenceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryUserInfoToMsenceRequest) Validate() error {
	return dara.Validate(s)
}
