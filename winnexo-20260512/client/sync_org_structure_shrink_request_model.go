// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncOrgStructureShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCorpId(v string) *SyncOrgStructureShrinkRequest
	GetCorpId() *string
	SetDepartmentsShrink(v string) *SyncOrgStructureShrinkRequest
	GetDepartmentsShrink() *string
	SetMembersShrink(v string) *SyncOrgStructureShrinkRequest
	GetMembersShrink() *string
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
	// 企业标识（必须与 listAvailableConfigs 返回的 corpId 一致）
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleCorpId
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 部门列表（至少包含一个根部门）
	//
	// This parameter is required.
	DepartmentsShrink *string `json:"departments,omitempty" xml:"departments,omitempty"`
	// 成员列表（syncMembers=true 时必须提供）
	MembersShrink *string `json:"members,omitempty" xml:"members,omitempty"`
	// 平台类型: saml / oauth2 / custom
	//
	// This parameter is required.
	//
	// example:
	//
	// saml
	PlatformType *string `json:"platformType,omitempty" xml:"platformType,omitempty"`
	// SSO 配置 ID（SAML/OAuth2 可选：不传时按 corpId 自动推导；若存在多个 IdP 使用相同 corpId 则必须显式传入，否则报 AMBIGUOUS 错误；custom 不需要）
	//
	// example:
	//
	// exampleSsoSettingsId
	SsoSettingsId *string `json:"ssoSettingsId,omitempty" xml:"ssoSettingsId,omitempty"`
	// 是否同步成员关系（custom 模式强制为 false）
	//
	// example:
	//
	// false
	SyncMembers *bool `json:"syncMembers,omitempty" xml:"syncMembers,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
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

func (s *SyncOrgStructureShrinkRequest) GetCorpId() *string {
	return s.CorpId
}

func (s *SyncOrgStructureShrinkRequest) GetDepartmentsShrink() *string {
	return s.DepartmentsShrink
}

func (s *SyncOrgStructureShrinkRequest) GetMembersShrink() *string {
	return s.MembersShrink
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
