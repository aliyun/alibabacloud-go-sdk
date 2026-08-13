// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIds(v []*string) *ListUsersRequest
	GetAccountIds() []*string
	SetIsActive(v bool) *ListUsersRequest
	GetIsActive() *bool
	SetKeyword(v string) *ListUsersRequest
	GetKeyword() *string
	SetPage(v int64) *ListUsersRequest
	GetPage() *int64
	SetPageSize(v int64) *ListUsersRequest
	GetPageSize() *int64
	SetRoleCodes(v []*string) *ListUsersRequest
	GetRoleCodes() []*string
	SetTenantId(v string) *ListUsersRequest
	GetTenantId() *string
}

type ListUsersRequest struct {
	// 按 WINNEXO 登录账号精确批量查询（多选）；与其他筛选条件取交集。不传或传空列表 [] 均视为不按账号筛选（返回全部符合其他条件的成员）
	//
	// example:
	//
	// string_value
	AccountIds []*string `json:"accountIds,omitempty" xml:"accountIds,omitempty" type:"Repeated"`
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
	RoleCodes []*string `json:"roleCodes,omitempty" xml:"roleCodes,omitempty" type:"Repeated"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) GetAccountIds() []*string {
	return s.AccountIds
}

func (s *ListUsersRequest) GetIsActive() *bool {
	return s.IsActive
}

func (s *ListUsersRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListUsersRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListUsersRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListUsersRequest) GetRoleCodes() []*string {
	return s.RoleCodes
}

func (s *ListUsersRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListUsersRequest) SetAccountIds(v []*string) *ListUsersRequest {
	s.AccountIds = v
	return s
}

func (s *ListUsersRequest) SetIsActive(v bool) *ListUsersRequest {
	s.IsActive = &v
	return s
}

func (s *ListUsersRequest) SetKeyword(v string) *ListUsersRequest {
	s.Keyword = &v
	return s
}

func (s *ListUsersRequest) SetPage(v int64) *ListUsersRequest {
	s.Page = &v
	return s
}

func (s *ListUsersRequest) SetPageSize(v int64) *ListUsersRequest {
	s.PageSize = &v
	return s
}

func (s *ListUsersRequest) SetRoleCodes(v []*string) *ListUsersRequest {
	s.RoleCodes = v
	return s
}

func (s *ListUsersRequest) SetTenantId(v string) *ListUsersRequest {
	s.TenantId = &v
	return s
}

func (s *ListUsersRequest) Validate() error {
	return dara.Validate(s)
}
