// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveProjectMemberFromRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v int64) *RemoveProjectMemberFromRoleRequest
	GetProjectId() *int64
	SetRoleCode(v string) *RemoveProjectMemberFromRoleRequest
	GetRoleCode() *string
	SetUserId(v string) *RemoveProjectMemberFromRoleRequest
	GetUserId() *string
}

type RemoveProjectMemberFromRoleRequest struct {
	// The DataWorks workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 27
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The code of the role in the DataWorks workspace. You can call the ListProjectRoles operation to query the codes of all roles in a DataWorks workspace. Valid values:
	//
	// 	- role_project_owner: workspace owner
	//
	// 	- role_project_admin: workspace administrator
	//
	// 	- role_project_dev: developer
	//
	// 	- role_project_pe: O\\&M engineer
	//
	// 	- role_project_deploy: deployment expert
	//
	// 	- role_project_guest: visitor
	//
	// 	- role_project_security: security administrator
	//
	// 	- role_project_tester: experiencer
	//
	// 	- role_project_erd: model designer
	//
	// This parameter is required.
	//
	// example:
	//
	// role_project_guest
	RoleCode *string `json:"RoleCode,omitempty" xml:"RoleCode,omitempty"`
	// The user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RemoveProjectMemberFromRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveProjectMemberFromRoleRequest) GoString() string {
	return s.String()
}

func (s *RemoveProjectMemberFromRoleRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *RemoveProjectMemberFromRoleRequest) GetRoleCode() *string {
	return s.RoleCode
}

func (s *RemoveProjectMemberFromRoleRequest) GetUserId() *string {
	return s.UserId
}

func (s *RemoveProjectMemberFromRoleRequest) SetProjectId(v int64) *RemoveProjectMemberFromRoleRequest {
	s.ProjectId = &v
	return s
}

func (s *RemoveProjectMemberFromRoleRequest) SetRoleCode(v string) *RemoveProjectMemberFromRoleRequest {
	s.RoleCode = &v
	return s
}

func (s *RemoveProjectMemberFromRoleRequest) SetUserId(v string) *RemoveProjectMemberFromRoleRequest {
	s.UserId = &v
	return s
}

func (s *RemoveProjectMemberFromRoleRequest) Validate() error {
	return dara.Validate(s)
}
