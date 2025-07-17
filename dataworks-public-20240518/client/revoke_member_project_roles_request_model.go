// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeMemberProjectRolesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v int64) *RevokeMemberProjectRolesRequest
	GetProjectId() *int64
	SetRoleCodes(v []*string) *RevokeMemberProjectRolesRequest
	GetRoleCodes() []*string
	SetUserId(v string) *RevokeMemberProjectRolesRequest
	GetUserId() *string
}

type RevokeMemberProjectRolesRequest struct {
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://dataworks.console.aliyun.com/workspace/list) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 105149
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The codes of the roles in the workspace. You can call the [ListProjectRoles](https://help.aliyun.com/document_detail/2853930.html) operation to query the codes of all roles in the workspace.
	//
	// You must configure this parameter to specify the roles that you want to revoke from the member in the workspace.
	//
	// This parameter is required.
	RoleCodes []*string `json:"RoleCodes,omitempty" xml:"RoleCodes,omitempty" type:"Repeated"`
	// The ID of the account used by the member in the workspace. You can log on to the [DataWorks console](https://dataworks.console.aliyun.com/product/ms_menu), choose More > Management Center in the left-side navigation pane, select the desired workspace on the Management Center page, and then click Go to Management Center. In the left-side navigation pane of the SettingCenter page, click Tenant Members and Roles. On the Tenant Members and Roles page, view the ID of the account used by the member in the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123422344899
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RevokeMemberProjectRolesRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeMemberProjectRolesRequest) GoString() string {
	return s.String()
}

func (s *RevokeMemberProjectRolesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *RevokeMemberProjectRolesRequest) GetRoleCodes() []*string {
	return s.RoleCodes
}

func (s *RevokeMemberProjectRolesRequest) GetUserId() *string {
	return s.UserId
}

func (s *RevokeMemberProjectRolesRequest) SetProjectId(v int64) *RevokeMemberProjectRolesRequest {
	s.ProjectId = &v
	return s
}

func (s *RevokeMemberProjectRolesRequest) SetRoleCodes(v []*string) *RevokeMemberProjectRolesRequest {
	s.RoleCodes = v
	return s
}

func (s *RevokeMemberProjectRolesRequest) SetUserId(v string) *RevokeMemberProjectRolesRequest {
	s.UserId = &v
	return s
}

func (s *RevokeMemberProjectRolesRequest) Validate() error {
	return dara.Validate(s)
}
