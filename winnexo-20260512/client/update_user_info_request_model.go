// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *UpdateUserInfoRequest
	GetAppToken() *string
	SetAvatar(v string) *UpdateUserInfoRequest
	GetAvatar() *string
	SetLanguagePreference(v string) *UpdateUserInfoRequest
	GetLanguagePreference() *string
	SetName(v string) *UpdateUserInfoRequest
	GetName() *string
	SetOffering(v string) *UpdateUserInfoRequest
	GetOffering() *string
	SetOpsUserCode(v string) *UpdateUserInfoRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *UpdateUserInfoRequest
	GetOpsUserId() *string
	SetProfileRoleInfo(v string) *UpdateUserInfoRequest
	GetProfileRoleInfo() *string
	SetSelfIntroduction(v string) *UpdateUserInfoRequest
	GetSelfIntroduction() *string
	SetTenantId(v string) *UpdateUserInfoRequest
	GetTenantId() *string
}

type UpdateUserInfoRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The profile picture URL.
	//
	// example:
	//
	// https://example.com/avatar.png
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// The language preference: zh-CN or en-US.
	//
	// example:
	//
	// string_value
	LanguagePreference *string `json:"languagePreference,omitempty" xml:"languagePreference,omitempty"`
	// The name.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The user service description. Maximum length: 1000 characters.
	//
	// example:
	//
	// string_value
	Offering *string `json:"offering,omitempty" xml:"offering,omitempty"`
	// 代理用户编码（租户内唯一），仅当应用凭证配置为 source_uid 时使用
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// 代理用户ID（WINNEXO 平台用户ID），仅当应用凭证配置为 winnexo_uid 时使用
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The user role description. This parameter is used when profileRole is set to Others. Maximum length: 100 characters.
	//
	// example:
	//
	// string_value
	ProfileRoleInfo *string `json:"profileRoleInfo,omitempty" xml:"profileRoleInfo,omitempty"`
	// The user self-introduction. Maximum length: 1000 characters.
	//
	// example:
	//
	// string_value
	SelfIntroduction *string `json:"selfIntroduction,omitempty" xml:"selfIntroduction,omitempty"`
	// The ID of the effective tenant.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateUserInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserInfoRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *UpdateUserInfoRequest) GetAvatar() *string {
	return s.Avatar
}

func (s *UpdateUserInfoRequest) GetLanguagePreference() *string {
	return s.LanguagePreference
}

func (s *UpdateUserInfoRequest) GetName() *string {
	return s.Name
}

func (s *UpdateUserInfoRequest) GetOffering() *string {
	return s.Offering
}

func (s *UpdateUserInfoRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *UpdateUserInfoRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *UpdateUserInfoRequest) GetProfileRoleInfo() *string {
	return s.ProfileRoleInfo
}

func (s *UpdateUserInfoRequest) GetSelfIntroduction() *string {
	return s.SelfIntroduction
}

func (s *UpdateUserInfoRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateUserInfoRequest) SetAppToken(v string) *UpdateUserInfoRequest {
	s.AppToken = &v
	return s
}

func (s *UpdateUserInfoRequest) SetAvatar(v string) *UpdateUserInfoRequest {
	s.Avatar = &v
	return s
}

func (s *UpdateUserInfoRequest) SetLanguagePreference(v string) *UpdateUserInfoRequest {
	s.LanguagePreference = &v
	return s
}

func (s *UpdateUserInfoRequest) SetName(v string) *UpdateUserInfoRequest {
	s.Name = &v
	return s
}

func (s *UpdateUserInfoRequest) SetOffering(v string) *UpdateUserInfoRequest {
	s.Offering = &v
	return s
}

func (s *UpdateUserInfoRequest) SetOpsUserCode(v string) *UpdateUserInfoRequest {
	s.OpsUserCode = &v
	return s
}

func (s *UpdateUserInfoRequest) SetOpsUserId(v string) *UpdateUserInfoRequest {
	s.OpsUserId = &v
	return s
}

func (s *UpdateUserInfoRequest) SetProfileRoleInfo(v string) *UpdateUserInfoRequest {
	s.ProfileRoleInfo = &v
	return s
}

func (s *UpdateUserInfoRequest) SetSelfIntroduction(v string) *UpdateUserInfoRequest {
	s.SelfIntroduction = &v
	return s
}

func (s *UpdateUserInfoRequest) SetTenantId(v string) *UpdateUserInfoRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateUserInfoRequest) Validate() error {
	return dara.Validate(s)
}
