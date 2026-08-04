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
	// example:
	//
	// inherit
	AuthConfig *string `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	// example:
	//
	// false
	IncludeAuthorization *bool `json:"includeAuthorization,omitempty" xml:"includeAuthorization,omitempty"`
	// example:
	//
	// true
	IncludeBalance *bool `json:"includeBalance,omitempty" xml:"includeBalance,omitempty"`
	// example:
	//
	// 张三
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// 1
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
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
