// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUserAuthToMsceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CheckUserAuthToMsceneRequest
	GetAppId() *string
	SetAuthToken(v string) *CheckUserAuthToMsceneRequest
	GetAuthToken() *string
	SetMiniProgramId(v string) *CheckUserAuthToMsceneRequest
	GetMiniProgramId() *string
	SetOpenUid(v string) *CheckUserAuthToMsceneRequest
	GetOpenUid() *string
	SetPlatformId(v string) *CheckUserAuthToMsceneRequest
	GetPlatformId() *string
	SetTenantId(v string) *CheckUserAuthToMsceneRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *CheckUserAuthToMsceneRequest
	GetWorkspaceId() *string
}

type CheckUserAuthToMsceneRequest struct {
	// example:
	//
	// ALIPUBE5C3F6D091419
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 111
	AuthToken *string `json:"AuthToken,omitempty" xml:"AuthToken,omitempty"`
	// example:
	//
	// 9999202604200001
	MiniProgramId *string `json:"MiniProgramId,omitempty" xml:"MiniProgramId,omitempty"`
	// example:
	//
	// QSz8YsUpMRiCAv0yKLi0OVGd0lQSg9CQkYI148GNIWE=
	OpenUid *string `json:"OpenUid,omitempty" xml:"OpenUid,omitempty"`
	// example:
	//
	// mPaaS_Goosefish
	PlatformId *string `json:"PlatformId,omitempty" xml:"PlatformId,omitempty"`
	// example:
	//
	// ZXCXMAHQ-zh_CN
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// default
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CheckUserAuthToMsceneRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckUserAuthToMsceneRequest) GoString() string {
	return s.String()
}

func (s *CheckUserAuthToMsceneRequest) GetAppId() *string {
	return s.AppId
}

func (s *CheckUserAuthToMsceneRequest) GetAuthToken() *string {
	return s.AuthToken
}

func (s *CheckUserAuthToMsceneRequest) GetMiniProgramId() *string {
	return s.MiniProgramId
}

func (s *CheckUserAuthToMsceneRequest) GetOpenUid() *string {
	return s.OpenUid
}

func (s *CheckUserAuthToMsceneRequest) GetPlatformId() *string {
	return s.PlatformId
}

func (s *CheckUserAuthToMsceneRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CheckUserAuthToMsceneRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CheckUserAuthToMsceneRequest) SetAppId(v string) *CheckUserAuthToMsceneRequest {
	s.AppId = &v
	return s
}

func (s *CheckUserAuthToMsceneRequest) SetAuthToken(v string) *CheckUserAuthToMsceneRequest {
	s.AuthToken = &v
	return s
}

func (s *CheckUserAuthToMsceneRequest) SetMiniProgramId(v string) *CheckUserAuthToMsceneRequest {
	s.MiniProgramId = &v
	return s
}

func (s *CheckUserAuthToMsceneRequest) SetOpenUid(v string) *CheckUserAuthToMsceneRequest {
	s.OpenUid = &v
	return s
}

func (s *CheckUserAuthToMsceneRequest) SetPlatformId(v string) *CheckUserAuthToMsceneRequest {
	s.PlatformId = &v
	return s
}

func (s *CheckUserAuthToMsceneRequest) SetTenantId(v string) *CheckUserAuthToMsceneRequest {
	s.TenantId = &v
	return s
}

func (s *CheckUserAuthToMsceneRequest) SetWorkspaceId(v string) *CheckUserAuthToMsceneRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CheckUserAuthToMsceneRequest) Validate() error {
	return dara.Validate(s)
}
