// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIdsShrink(v string) *ListUsersShrinkRequest
	GetAccountIdsShrink() *string
	SetIsActive(v bool) *ListUsersShrinkRequest
	GetIsActive() *bool
	SetKeyword(v string) *ListUsersShrinkRequest
	GetKeyword() *string
	SetPage(v int64) *ListUsersShrinkRequest
	GetPage() *int64
	SetPageSize(v int64) *ListUsersShrinkRequest
	GetPageSize() *int64
	SetRoleCodesShrink(v string) *ListUsersShrinkRequest
	GetRoleCodesShrink() *string
	SetTenantId(v string) *ListUsersShrinkRequest
	GetTenantId() *string
}

type ListUsersShrinkRequest struct {
	// 按 WINNEXO 登录账号精确批量查询（多选）；与其他筛选条件取交集。不传或传空列表 [] 均视为不按账号筛选（返回全部符合其他条件的成员）
	//
	// example:
	//
	// string_value
	AccountIdsShrink *string `json:"accountIds,omitempty" xml:"accountIds,omitempty"`
	// 启用/停用状态筛选
	//
	// example:
	//
	// true
	IsActive *bool `json:"isActive,omitempty" xml:"isActive,omitempty"`
	// 搜索关键词（模糊匹配显示名和账号）
	//
	// example:
	//
	// 示例关键词
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 页码（从1开始）
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页数量（最大100）
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 按角色筛选，可选值: SUPER_ADMIN / SYSTEM_ADMIN / SEMANTIC_ADMIN / SKILL_ADMIN / KB_ADMIN / AGENT_ADMIN / APPLICATION_USER
	//
	// example:
	//
	// string_value
	RoleCodesShrink *string `json:"roleCodes,omitempty" xml:"roleCodes,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListUsersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUsersShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListUsersShrinkRequest) GetAccountIdsShrink() *string {
	return s.AccountIdsShrink
}

func (s *ListUsersShrinkRequest) GetIsActive() *bool {
	return s.IsActive
}

func (s *ListUsersShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListUsersShrinkRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListUsersShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListUsersShrinkRequest) GetRoleCodesShrink() *string {
	return s.RoleCodesShrink
}

func (s *ListUsersShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListUsersShrinkRequest) SetAccountIdsShrink(v string) *ListUsersShrinkRequest {
	s.AccountIdsShrink = &v
	return s
}

func (s *ListUsersShrinkRequest) SetIsActive(v bool) *ListUsersShrinkRequest {
	s.IsActive = &v
	return s
}

func (s *ListUsersShrinkRequest) SetKeyword(v string) *ListUsersShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListUsersShrinkRequest) SetPage(v int64) *ListUsersShrinkRequest {
	s.Page = &v
	return s
}

func (s *ListUsersShrinkRequest) SetPageSize(v int64) *ListUsersShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListUsersShrinkRequest) SetRoleCodesShrink(v string) *ListUsersShrinkRequest {
	s.RoleCodesShrink = &v
	return s
}

func (s *ListUsersShrinkRequest) SetTenantId(v string) *ListUsersShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *ListUsersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
