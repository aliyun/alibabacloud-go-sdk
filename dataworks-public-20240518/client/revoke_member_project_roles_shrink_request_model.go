// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeMemberProjectRolesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v int64) *RevokeMemberProjectRolesShrinkRequest
	GetProjectId() *int64
	SetRoleCodesShrink(v string) *RevokeMemberProjectRolesShrinkRequest
	GetRoleCodesShrink() *string
	SetUserId(v string) *RevokeMemberProjectRolesShrinkRequest
	GetUserId() *string
}

type RevokeMemberProjectRolesShrinkRequest struct {
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://dataworks.console.aliyun.com/workspace/list) and go to the workspace settings page to obtain the workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 105149
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The list of workspace role codes. You can call [ListProjectRoles](https://help.aliyun.com/document_detail/2853930.html) to obtain the role codes.
	//
	// This parameter specifies the workspace roles to be revoked by this API call.
	//
	// This parameter is required.
	RoleCodesShrink *string `json:"RoleCodes,omitempty" xml:"RoleCodes,omitempty"`
	// The ID of the DataWorks account. You can log on to the [DataWorks console - Management Center](https://dataworks.console.aliyun.com/product/ms_menu), select the workspace that you want to manage, go to the Tenant Members and Roles page, and view the account ID of the member whose roles you want to revoke.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123422344899
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RevokeMemberProjectRolesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeMemberProjectRolesShrinkRequest) GoString() string {
	return s.String()
}

func (s *RevokeMemberProjectRolesShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *RevokeMemberProjectRolesShrinkRequest) GetRoleCodesShrink() *string {
	return s.RoleCodesShrink
}

func (s *RevokeMemberProjectRolesShrinkRequest) GetUserId() *string {
	return s.UserId
}

func (s *RevokeMemberProjectRolesShrinkRequest) SetProjectId(v int64) *RevokeMemberProjectRolesShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *RevokeMemberProjectRolesShrinkRequest) SetRoleCodesShrink(v string) *RevokeMemberProjectRolesShrinkRequest {
	s.RoleCodesShrink = &v
	return s
}

func (s *RevokeMemberProjectRolesShrinkRequest) SetUserId(v string) *RevokeMemberProjectRolesShrinkRequest {
	s.UserId = &v
	return s
}

func (s *RevokeMemberProjectRolesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
