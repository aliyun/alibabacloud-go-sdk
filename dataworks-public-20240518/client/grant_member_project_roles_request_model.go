// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantMemberProjectRolesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v int64) *GrantMemberProjectRolesRequest
	GetProjectId() *int64
	SetRoleCodes(v []*string) *GrantMemberProjectRolesRequest
	GetRoleCodes() []*string
	SetUserId(v string) *GrantMemberProjectRolesRequest
	GetUserId() *string
}

type GrantMemberProjectRolesRequest struct {
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
	RoleCodes []*string `json:"RoleCodes,omitempty" xml:"RoleCodes,omitempty" type:"Repeated"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://dataworks.console.aliyun.com/product/ms_menu), choose More > Management Center in the left-side navigation pane, select the desired workspace on the Management Center page, and then click Go to Management Center. In the left-side navigation pane of the SettingCenter page, click Tenant Members and Roles. On the Tenant Members and Roles page, view the IDs of the accounts used by the members in the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123422344899
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GrantMemberProjectRolesRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantMemberProjectRolesRequest) GoString() string {
	return s.String()
}

func (s *GrantMemberProjectRolesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GrantMemberProjectRolesRequest) GetRoleCodes() []*string {
	return s.RoleCodes
}

func (s *GrantMemberProjectRolesRequest) GetUserId() *string {
	return s.UserId
}

func (s *GrantMemberProjectRolesRequest) SetProjectId(v int64) *GrantMemberProjectRolesRequest {
	s.ProjectId = &v
	return s
}

func (s *GrantMemberProjectRolesRequest) SetRoleCodes(v []*string) *GrantMemberProjectRolesRequest {
	s.RoleCodes = v
	return s
}

func (s *GrantMemberProjectRolesRequest) SetUserId(v string) *GrantMemberProjectRolesRequest {
	s.UserId = &v
	return s
}

func (s *GrantMemberProjectRolesRequest) Validate() error {
	return dara.Validate(s)
}
