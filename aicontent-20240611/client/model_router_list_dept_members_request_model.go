// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterListDeptMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthConfig(v string) *ModelRouterListDeptMembersRequest
	GetAuthConfig() *string
	SetIncludeAuthorization(v bool) *ModelRouterListDeptMembersRequest
	GetIncludeAuthorization() *bool
	SetIncludeBalance(v bool) *ModelRouterListDeptMembersRequest
	GetIncludeBalance() *bool
	SetKeyword(v string) *ModelRouterListDeptMembersRequest
	GetKeyword() *string
	SetModel(v string) *ModelRouterListDeptMembersRequest
	GetModel() *string
	SetPageIndex(v int32) *ModelRouterListDeptMembersRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ModelRouterListDeptMembersRequest
	GetPageSize() *int32
}

type ModelRouterListDeptMembersRequest struct {
	// The authorization configuration filter. Valid values:
	//
	// - inherit: only members that inherit department settings.
	//
	// - custom: only members with custom settings.
	//
	// - Empty: all members.
	//
	// example:
	//
	// inherit
	AuthConfig *string `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	// Specifies whether to include the authorized models and the number of associated keys for the member.
	//
	// example:
	//
	// false
	IncludeAuthorization *bool `json:"includeAuthorization,omitempty" xml:"includeAuthorization,omitempty"`
	// Specifies whether to include the monthly and permanent balance of the member\\"s sub-wallet.
	//
	// example:
	//
	// true
	IncludeBalance *bool `json:"includeBalance,omitempty" xml:"includeBalance,omitempty"`
	// The search keyword.
	//
	// example:
	//
	// John
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// Filters members by the authorized model ID.
	//
	// example:
	//
	// 1
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ModelRouterListDeptMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterListDeptMembersRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterListDeptMembersRequest) GetAuthConfig() *string {
	return s.AuthConfig
}

func (s *ModelRouterListDeptMembersRequest) GetIncludeAuthorization() *bool {
	return s.IncludeAuthorization
}

func (s *ModelRouterListDeptMembersRequest) GetIncludeBalance() *bool {
	return s.IncludeBalance
}

func (s *ModelRouterListDeptMembersRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ModelRouterListDeptMembersRequest) GetModel() *string {
	return s.Model
}

func (s *ModelRouterListDeptMembersRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ModelRouterListDeptMembersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterListDeptMembersRequest) SetAuthConfig(v string) *ModelRouterListDeptMembersRequest {
	s.AuthConfig = &v
	return s
}

func (s *ModelRouterListDeptMembersRequest) SetIncludeAuthorization(v bool) *ModelRouterListDeptMembersRequest {
	s.IncludeAuthorization = &v
	return s
}

func (s *ModelRouterListDeptMembersRequest) SetIncludeBalance(v bool) *ModelRouterListDeptMembersRequest {
	s.IncludeBalance = &v
	return s
}

func (s *ModelRouterListDeptMembersRequest) SetKeyword(v string) *ModelRouterListDeptMembersRequest {
	s.Keyword = &v
	return s
}

func (s *ModelRouterListDeptMembersRequest) SetModel(v string) *ModelRouterListDeptMembersRequest {
	s.Model = &v
	return s
}

func (s *ModelRouterListDeptMembersRequest) SetPageIndex(v int32) *ModelRouterListDeptMembersRequest {
	s.PageIndex = &v
	return s
}

func (s *ModelRouterListDeptMembersRequest) SetPageSize(v int32) *ModelRouterListDeptMembersRequest {
	s.PageSize = &v
	return s
}

func (s *ModelRouterListDeptMembersRequest) Validate() error {
	return dara.Validate(s)
}
