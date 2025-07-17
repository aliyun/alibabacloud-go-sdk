// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v int64) *CreateProjectMemberRequest
	GetProjectId() *int64
	SetRoleCodes(v []*string) *CreateProjectMemberRequest
	GetRoleCodes() []*string
	SetUserId(v string) *CreateProjectMemberRequest
	GetUserId() *string
}

type CreateProjectMemberRequest struct {
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
	RoleCodes []*string `json:"RoleCodes,omitempty" xml:"RoleCodes,omitempty" type:"Repeated"`
	// The ID of the account that you want to add to the workspace as a member. You can log on to the [DataWorks console](https://dataworks.console.aliyun.com/product/ms_menu), choose More > Management Center in the left-side navigation pane, select the desired workspace on the Management Center page, and then click Go to Management Center. In the left-side navigation pane of the SettingCenter page, click **Tenant Members and Roles**. On the Tenant Members and Roles page, view the ID of the account that you want to add to the workspace as a member.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123422344899
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateProjectMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectMemberRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectMemberRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateProjectMemberRequest) GetRoleCodes() []*string {
	return s.RoleCodes
}

func (s *CreateProjectMemberRequest) GetUserId() *string {
	return s.UserId
}

func (s *CreateProjectMemberRequest) SetProjectId(v int64) *CreateProjectMemberRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateProjectMemberRequest) SetRoleCodes(v []*string) *CreateProjectMemberRequest {
	s.RoleCodes = v
	return s
}

func (s *CreateProjectMemberRequest) SetUserId(v string) *CreateProjectMemberRequest {
	s.UserId = &v
	return s
}

func (s *CreateProjectMemberRequest) Validate() error {
	return dara.Validate(s)
}
