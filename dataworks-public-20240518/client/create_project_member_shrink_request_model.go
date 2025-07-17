// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectMemberShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v int64) *CreateProjectMemberShrinkRequest
	GetProjectId() *int64
	SetRoleCodesShrink(v string) *CreateProjectMemberShrinkRequest
	GetRoleCodesShrink() *string
	SetUserId(v string) *CreateProjectMemberShrinkRequest
	GetUserId() *string
}

type CreateProjectMemberShrinkRequest struct {
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 24054
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The codes of the roles in the workspace. You can call the [ListProjectRoles](https://help.aliyun.com/document_detail/2853930.html) operation to query the codes of all roles in the workspace.
	//
	// This parameter specifies the roles that you can assign to a member when you add the member.
	//
	// This parameter is required.
	RoleCodesShrink *string `json:"RoleCodes,omitempty" xml:"RoleCodes,omitempty"`
	// The ID of the account that you want to add to the workspace as a member. You can log on to the [DataWorks console](https://dataworks.console.aliyun.com/product/ms_menu), choose More > Management Center in the left-side navigation pane, select the desired workspace on the Management Center page, and then click Go to Management Center. In the left-side navigation pane of the SettingCenter page, click **Tenant Members and Roles**. On the Tenant Members and Roles page, view the ID of the account that you want to add to the workspace as a member.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123422344899
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateProjectMemberShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectMemberShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectMemberShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateProjectMemberShrinkRequest) GetRoleCodesShrink() *string {
	return s.RoleCodesShrink
}

func (s *CreateProjectMemberShrinkRequest) GetUserId() *string {
	return s.UserId
}

func (s *CreateProjectMemberShrinkRequest) SetProjectId(v int64) *CreateProjectMemberShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateProjectMemberShrinkRequest) SetRoleCodesShrink(v string) *CreateProjectMemberShrinkRequest {
	s.RoleCodesShrink = &v
	return s
}

func (s *CreateProjectMemberShrinkRequest) SetUserId(v string) *CreateProjectMemberShrinkRequest {
	s.UserId = &v
	return s
}

func (s *CreateProjectMemberShrinkRequest) Validate() error {
	return dara.Validate(s)
}
