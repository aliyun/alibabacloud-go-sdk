// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncOrgStructureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCorpId(v string) *SyncOrgStructureRequest
	GetCorpId() *string
	SetDepartments(v []*SyncOrgStructureRequestDepartments) *SyncOrgStructureRequest
	GetDepartments() []*SyncOrgStructureRequestDepartments
	SetMembers(v []*SyncOrgStructureRequestMembers) *SyncOrgStructureRequest
	GetMembers() []*SyncOrgStructureRequestMembers
	SetPlatformType(v string) *SyncOrgStructureRequest
	GetPlatformType() *string
	SetSsoSettingsId(v string) *SyncOrgStructureRequest
	GetSsoSettingsId() *string
	SetSyncMembers(v bool) *SyncOrgStructureRequest
	GetSyncMembers() *bool
	SetTenantId(v string) *SyncOrgStructureRequest
	GetTenantId() *string
}

type SyncOrgStructureRequest struct {
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
	Departments []*SyncOrgStructureRequestDepartments `json:"departments,omitempty" xml:"departments,omitempty" type:"Repeated"`
	// 成员列表（syncMembers=true 时必须提供）
	Members []*SyncOrgStructureRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
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

func (s SyncOrgStructureRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncOrgStructureRequest) GoString() string {
	return s.String()
}

func (s *SyncOrgStructureRequest) GetCorpId() *string {
	return s.CorpId
}

func (s *SyncOrgStructureRequest) GetDepartments() []*SyncOrgStructureRequestDepartments {
	return s.Departments
}

func (s *SyncOrgStructureRequest) GetMembers() []*SyncOrgStructureRequestMembers {
	return s.Members
}

func (s *SyncOrgStructureRequest) GetPlatformType() *string {
	return s.PlatformType
}

func (s *SyncOrgStructureRequest) GetSsoSettingsId() *string {
	return s.SsoSettingsId
}

func (s *SyncOrgStructureRequest) GetSyncMembers() *bool {
	return s.SyncMembers
}

func (s *SyncOrgStructureRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SyncOrgStructureRequest) SetCorpId(v string) *SyncOrgStructureRequest {
	s.CorpId = &v
	return s
}

func (s *SyncOrgStructureRequest) SetDepartments(v []*SyncOrgStructureRequestDepartments) *SyncOrgStructureRequest {
	s.Departments = v
	return s
}

func (s *SyncOrgStructureRequest) SetMembers(v []*SyncOrgStructureRequestMembers) *SyncOrgStructureRequest {
	s.Members = v
	return s
}

func (s *SyncOrgStructureRequest) SetPlatformType(v string) *SyncOrgStructureRequest {
	s.PlatformType = &v
	return s
}

func (s *SyncOrgStructureRequest) SetSsoSettingsId(v string) *SyncOrgStructureRequest {
	s.SsoSettingsId = &v
	return s
}

func (s *SyncOrgStructureRequest) SetSyncMembers(v bool) *SyncOrgStructureRequest {
	s.SyncMembers = &v
	return s
}

func (s *SyncOrgStructureRequest) SetTenantId(v string) *SyncOrgStructureRequest {
	s.TenantId = &v
	return s
}

func (s *SyncOrgStructureRequest) Validate() error {
	if s.Departments != nil {
		for _, item := range s.Departments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Members != nil {
		for _, item := range s.Members {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SyncOrgStructureRequestDepartments struct {
	// 部门 ID（外部标识，客户端自行保证唯一性）
	//
	// example:
	//
	// exampleDeptId
	DeptId *string `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// 部门名称
	//
	// example:
	//
	// string_value
	DeptName *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	// 排序号（数值越小越靠前）
	//
	// example:
	//
	// 1
	Order *int64 `json:"order,omitempty" xml:"order,omitempty"`
	// 父部门 ID（null 表示一级部门/根部门）
	//
	// example:
	//
	// exampleParentDeptId
	ParentDeptId *string `json:"parentDeptId,omitempty" xml:"parentDeptId,omitempty"`
}

func (s SyncOrgStructureRequestDepartments) String() string {
	return dara.Prettify(s)
}

func (s SyncOrgStructureRequestDepartments) GoString() string {
	return s.String()
}

func (s *SyncOrgStructureRequestDepartments) GetDeptId() *string {
	return s.DeptId
}

func (s *SyncOrgStructureRequestDepartments) GetDeptName() *string {
	return s.DeptName
}

func (s *SyncOrgStructureRequestDepartments) GetOrder() *int64 {
	return s.Order
}

func (s *SyncOrgStructureRequestDepartments) GetParentDeptId() *string {
	return s.ParentDeptId
}

func (s *SyncOrgStructureRequestDepartments) SetDeptId(v string) *SyncOrgStructureRequestDepartments {
	s.DeptId = &v
	return s
}

func (s *SyncOrgStructureRequestDepartments) SetDeptName(v string) *SyncOrgStructureRequestDepartments {
	s.DeptName = &v
	return s
}

func (s *SyncOrgStructureRequestDepartments) SetOrder(v int64) *SyncOrgStructureRequestDepartments {
	s.Order = &v
	return s
}

func (s *SyncOrgStructureRequestDepartments) SetParentDeptId(v string) *SyncOrgStructureRequestDepartments {
	s.ParentDeptId = &v
	return s
}

func (s *SyncOrgStructureRequestDepartments) Validate() error {
	return dara.Validate(s)
}

type SyncOrgStructureRequestMembers struct {
	// 用户标识（SAML 场景为邮箱/UPN，需与 rbj_user_account.account_id 匹配）
	//
	// example:
	//
	// exampleAccountId
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// 所属部门 ID（必须与 departments 中的 deptId 对应）
	//
	// example:
	//
	// exampleDeptId
	DeptId *string `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// 用户姓名（展示用，可选）
	//
	// example:
	//
	// 示例名称.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s SyncOrgStructureRequestMembers) String() string {
	return dara.Prettify(s)
}

func (s SyncOrgStructureRequestMembers) GoString() string {
	return s.String()
}

func (s *SyncOrgStructureRequestMembers) GetAccountId() *string {
	return s.AccountId
}

func (s *SyncOrgStructureRequestMembers) GetDeptId() *string {
	return s.DeptId
}

func (s *SyncOrgStructureRequestMembers) GetName() *string {
	return s.Name
}

func (s *SyncOrgStructureRequestMembers) SetAccountId(v string) *SyncOrgStructureRequestMembers {
	s.AccountId = &v
	return s
}

func (s *SyncOrgStructureRequestMembers) SetDeptId(v string) *SyncOrgStructureRequestMembers {
	s.DeptId = &v
	return s
}

func (s *SyncOrgStructureRequestMembers) SetName(v string) *SyncOrgStructureRequestMembers {
	s.Name = &v
	return s
}

func (s *SyncOrgStructureRequestMembers) Validate() error {
	return dara.Validate(s)
}
