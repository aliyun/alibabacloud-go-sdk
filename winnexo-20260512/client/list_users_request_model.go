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
	// The list of Alibaba Cloud account IDs.
	//
	// example:
	//
	// string_value
	AccountIds []*string `json:"accountIds,omitempty" xml:"accountIds,omitempty" type:"Repeated"`
	// Specifies whether the account is activated.
	//
	//  - **true**: Activated.
	//
	// - **false**: Not activated.
	//
	// example:
	//
	// true
	IsActive *bool `json:"isActive,omitempty" xml:"isActive,omitempty"`
	// The keyword for searching products. Fuzzy match is supported.
	//
	// example:
	//
	// SampleKeyword
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page.
	//
	// > The maximum number of entries per page is 30.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The list of new system role codes (full replacement, at least one role must be included). Valid values: SUPER_ADMIN / SYSTEM_ADMIN / SEMANTIC_ADMIN / SKILL_ADMIN / KB_ADMIN / AGENT_ADMIN / APPLICATION_USER.
	//
	// example:
	//
	// string_value
	RoleCodes []*string `json:"roleCodes,omitempty" xml:"roleCodes,omitempty" type:"Repeated"`
	// The tenant ID. This is a common parameter. The winnexo-cli passes this parameter explicitly by using --tenant-id.
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
