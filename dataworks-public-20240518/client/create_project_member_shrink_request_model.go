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
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace Management page to obtain the ID.
	//
	// This parameter specifies the DataWorks workspace for this API call operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 24054
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The list of workspace role codes. You can call [ListProjectRoles](https://help.aliyun.com/document_detail/2853930.html) to obtain the role codes.
	//
	// This parameter is used to grant workspace roles to the member when adding the member to the workspace.
	//
	// This parameter is required.
	RoleCodesShrink *string `json:"RoleCodes,omitempty" xml:"RoleCodes,omitempty"`
	// The ID of the DataWorks account. You can log on to the [DataWorks console - Management Center](https://dataworks.console.aliyun.com/product/ms_menu), select the workspace to which you want to add a member, go to the Management Center page, and then navigate to the **Tenant Members and Roles*	- page to view the account ID of the user you want to add to the workspace.
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
