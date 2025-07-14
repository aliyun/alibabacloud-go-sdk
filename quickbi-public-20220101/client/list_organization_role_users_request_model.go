// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrganizationRoleUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListOrganizationRoleUsersRequest
	GetKeyword() *string
	SetPageNum(v int32) *ListOrganizationRoleUsersRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListOrganizationRoleUsersRequest
	GetPageSize() *int32
	SetRoleId(v int64) *ListOrganizationRoleUsersRequest
	GetRoleId() *int64
}

type ListOrganizationRoleUsersRequest struct {
	// Keyword for the nickname of the organization member.
	//
	// example:
	//
	// zhangsan
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// Page number.
	//
	// - Default value is 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// Number of items per page.
	//
	// - Default value is 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Organization role ID, including predefined roles and custom roles:
	//
	// - Organization Administrator (predefined role): 111111111
	//
	// - Permission Administrator (predefined role): 111111112
	//
	// - Regular User (predefined role): 111111113
	//
	// - Custom Role: The corresponding role ID for a custom role
	//
	// This parameter is required.
	//
	// example:
	//
	// 111111111
	RoleId *int64 `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
}

func (s ListOrganizationRoleUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationRoleUsersRequest) GoString() string {
	return s.String()
}

func (s *ListOrganizationRoleUsersRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListOrganizationRoleUsersRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListOrganizationRoleUsersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOrganizationRoleUsersRequest) GetRoleId() *int64 {
	return s.RoleId
}

func (s *ListOrganizationRoleUsersRequest) SetKeyword(v string) *ListOrganizationRoleUsersRequest {
	s.Keyword = &v
	return s
}

func (s *ListOrganizationRoleUsersRequest) SetPageNum(v int32) *ListOrganizationRoleUsersRequest {
	s.PageNum = &v
	return s
}

func (s *ListOrganizationRoleUsersRequest) SetPageSize(v int32) *ListOrganizationRoleUsersRequest {
	s.PageSize = &v
	return s
}

func (s *ListOrganizationRoleUsersRequest) SetRoleId(v int64) *ListOrganizationRoleUsersRequest {
	s.RoleId = &v
	return s
}

func (s *ListOrganizationRoleUsersRequest) Validate() error {
	return dara.Validate(s)
}
