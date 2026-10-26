// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspaceRoleUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListWorkspaceRoleUsersRequest
	GetKeyword() *string
	SetPageNum(v int32) *ListWorkspaceRoleUsersRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListWorkspaceRoleUsersRequest
	GetPageSize() *int32
	SetRoleId(v int64) *ListWorkspaceRoleUsersRequest
	GetRoleId() *int64
	SetWorkspaceId(v string) *ListWorkspaceRoleUsersRequest
	GetWorkspaceId() *string
}

type ListWorkspaceRoleUsersRequest struct {
	// Keyword for the user\\"s nickname.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// Current page number for pagination:
	//
	// - Starting value: 1
	//
	// - Default value: 1
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// Number of items per page for pagination:
	//
	// - Default value: 10
	//
	// - Maximum value: 1000
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Workspace role ID, including predefined roles and custom roles:
	//
	// - 25: Workspace Administrator (predefined role)
	//
	// - 26: Developer (predefined role)
	//
	// - 27: Analyst (predefined role)
	//
	// - 30: Viewer (predefined role)
	//
	// - Custom roles: The corresponding role ID for custom roles
	//
	// This parameter is required.
	//
	// example:
	//
	// 25
	RoleId *int64 `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// The ID of the workspace. This parameter is optional. If you do not set this parameter, the roles of all workspaces are returned.
	//
	// example:
	//
	// 726bee5a-****-43e1-9a8e-b550f0120f35
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListWorkspaceRoleUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceRoleUsersRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspaceRoleUsersRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListWorkspaceRoleUsersRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListWorkspaceRoleUsersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWorkspaceRoleUsersRequest) GetRoleId() *int64 {
	return s.RoleId
}

func (s *ListWorkspaceRoleUsersRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListWorkspaceRoleUsersRequest) SetKeyword(v string) *ListWorkspaceRoleUsersRequest {
	s.Keyword = &v
	return s
}

func (s *ListWorkspaceRoleUsersRequest) SetPageNum(v int32) *ListWorkspaceRoleUsersRequest {
	s.PageNum = &v
	return s
}

func (s *ListWorkspaceRoleUsersRequest) SetPageSize(v int32) *ListWorkspaceRoleUsersRequest {
	s.PageSize = &v
	return s
}

func (s *ListWorkspaceRoleUsersRequest) SetRoleId(v int64) *ListWorkspaceRoleUsersRequest {
	s.RoleId = &v
	return s
}

func (s *ListWorkspaceRoleUsersRequest) SetWorkspaceId(v string) *ListWorkspaceRoleUsersRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListWorkspaceRoleUsersRequest) Validate() error {
	return dara.Validate(s)
}
