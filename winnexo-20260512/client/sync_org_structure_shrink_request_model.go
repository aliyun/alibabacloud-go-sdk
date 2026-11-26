// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncOrgStructureShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *SyncOrgStructureShrinkRequest
	GetAppToken() *string
	SetCorpId(v string) *SyncOrgStructureShrinkRequest
	GetCorpId() *string
	SetDepartmentsShrink(v string) *SyncOrgStructureShrinkRequest
	GetDepartmentsShrink() *string
	SetMembersShrink(v string) *SyncOrgStructureShrinkRequest
	GetMembersShrink() *string
	SetOpsUserCode(v string) *SyncOrgStructureShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *SyncOrgStructureShrinkRequest
	GetOpsUserId() *string
	SetPlatformType(v string) *SyncOrgStructureShrinkRequest
	GetPlatformType() *string
	SetSsoSettingsId(v string) *SyncOrgStructureShrinkRequest
	GetSsoSettingsId() *string
	SetSyncMembers(v bool) *SyncOrgStructureShrinkRequest
	GetSyncMembers() *bool
	SetTenantId(v string) *SyncOrgStructureShrinkRequest
	GetTenantId() *string
}

type SyncOrgStructureShrinkRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The enterprise identifier. This value must match the corpId returned by listAvailableConfigs.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleCorpId
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// The department list. At least one root department must be included.
	//
	// This parameter is required.
	DepartmentsShrink *string `json:"departments,omitempty" xml:"departments,omitempty"`
	// The member list. This parameter is required when syncMembers is set to true.
	MembersShrink *string `json:"members,omitempty" xml:"members,omitempty"`
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
	// The platform type. Valid values: saml, oauth2, or custom.
	//
	// This parameter is required.
	//
	// example:
	//
	// saml
	PlatformType *string `json:"platformType,omitempty" xml:"platformType,omitempty"`
	// The SSO configuration ID. For SAML/OAuth2, this parameter is optional. If not specified, the value is automatically derived based on corpId. If multiple IdPs use the same corpId, you must explicitly specify this parameter. Otherwise, an AMBIGUOUS error is returned. This parameter is not required for custom.
	//
	// example:
	//
	// exampleSsoSettingsId
	SsoSettingsId *string `json:"ssoSettingsId,omitempty" xml:"ssoSettingsId,omitempty"`
	// Specifies whether to synchronize member relationships. In custom mode, this parameter is forced to false.
	//
	// example:
	//
	// false
	SyncMembers *bool `json:"syncMembers,omitempty" xml:"syncMembers,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s SyncOrgStructureShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncOrgStructureShrinkRequest) GoString() string {
	return s.String()
}

func (s *SyncOrgStructureShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *SyncOrgStructureShrinkRequest) GetCorpId() *string {
	return s.CorpId
}

func (s *SyncOrgStructureShrinkRequest) GetDepartmentsShrink() *string {
	return s.DepartmentsShrink
}

func (s *SyncOrgStructureShrinkRequest) GetMembersShrink() *string {
	return s.MembersShrink
}

func (s *SyncOrgStructureShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *SyncOrgStructureShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *SyncOrgStructureShrinkRequest) GetPlatformType() *string {
	return s.PlatformType
}

func (s *SyncOrgStructureShrinkRequest) GetSsoSettingsId() *string {
	return s.SsoSettingsId
}

func (s *SyncOrgStructureShrinkRequest) GetSyncMembers() *bool {
	return s.SyncMembers
}

func (s *SyncOrgStructureShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SyncOrgStructureShrinkRequest) SetAppToken(v string) *SyncOrgStructureShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *SyncOrgStructureShrinkRequest) SetCorpId(v string) *SyncOrgStructureShrinkRequest {
	s.CorpId = &v
	return s
}

func (s *SyncOrgStructureShrinkRequest) SetDepartmentsShrink(v string) *SyncOrgStructureShrinkRequest {
	s.DepartmentsShrink = &v
	return s
}

func (s *SyncOrgStructureShrinkRequest) SetMembersShrink(v string) *SyncOrgStructureShrinkRequest {
	s.MembersShrink = &v
	return s
}

func (s *SyncOrgStructureShrinkRequest) SetOpsUserCode(v string) *SyncOrgStructureShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *SyncOrgStructureShrinkRequest) SetOpsUserId(v string) *SyncOrgStructureShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *SyncOrgStructureShrinkRequest) SetPlatformType(v string) *SyncOrgStructureShrinkRequest {
	s.PlatformType = &v
	return s
}

func (s *SyncOrgStructureShrinkRequest) SetSsoSettingsId(v string) *SyncOrgStructureShrinkRequest {
	s.SsoSettingsId = &v
	return s
}

func (s *SyncOrgStructureShrinkRequest) SetSyncMembers(v bool) *SyncOrgStructureShrinkRequest {
	s.SyncMembers = &v
	return s
}

func (s *SyncOrgStructureShrinkRequest) SetTenantId(v string) *SyncOrgStructureShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *SyncOrgStructureShrinkRequest) Validate() error {
	return dara.Validate(s)
}
