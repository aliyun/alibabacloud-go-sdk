// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantMemberProjectRolesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v int64) *GrantMemberProjectRolesShrinkRequest
	GetProjectId() *int64
	SetRoleCodesShrink(v string) *GrantMemberProjectRolesShrinkRequest
	GetRoleCodesShrink() *string
	SetUserId(v string) *GrantMemberProjectRolesShrinkRequest
	GetUserId() *string
}

type GrantMemberProjectRolesShrinkRequest struct {
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 105149
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The codes of the roles in the workspace. You can call the [ListProjectRoles](https://help.aliyun.com/document_detail/2853930.html) operation to query the codes of all roles in the workspace.
	//
	// You must configure this parameter to specify the roles that you want to assign to members in the workspace.
	//
	// This parameter is required.
	RoleCodesShrink *string `json:"RoleCodes,omitempty" xml:"RoleCodes,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://dataworks.console.aliyun.com/product/ms_menu), choose More > Management Center in the left-side navigation pane, select the desired workspace on the Management Center page, and then click Go to Management Center. In the left-side navigation pane of the SettingCenter page, click Tenant Members and Roles. On the Tenant Members and Roles page, view the IDs of the accounts used by the members in the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123422344899
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GrantMemberProjectRolesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantMemberProjectRolesShrinkRequest) GoString() string {
	return s.String()
}

func (s *GrantMemberProjectRolesShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GrantMemberProjectRolesShrinkRequest) GetRoleCodesShrink() *string {
	return s.RoleCodesShrink
}

func (s *GrantMemberProjectRolesShrinkRequest) GetUserId() *string {
	return s.UserId
}

func (s *GrantMemberProjectRolesShrinkRequest) SetProjectId(v int64) *GrantMemberProjectRolesShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *GrantMemberProjectRolesShrinkRequest) SetRoleCodesShrink(v string) *GrantMemberProjectRolesShrinkRequest {
	s.RoleCodesShrink = &v
	return s
}

func (s *GrantMemberProjectRolesShrinkRequest) SetUserId(v string) *GrantMemberProjectRolesShrinkRequest {
	s.UserId = &v
	return s
}

func (s *GrantMemberProjectRolesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
