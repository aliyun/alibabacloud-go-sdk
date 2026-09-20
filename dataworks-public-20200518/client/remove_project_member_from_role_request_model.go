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
	// The ID of the DataWorks workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 27
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The role code of the DataWorks workspace. You can invoke the ListProjectRoles operation to obtain the list of role codes for a project.
	//
	// The default preset roles are as follows:
	//
	// - role_project_owner: Project owner.
	//
	// - role_project_admin: Storage management administrator.
	//
	// - role_project_dev: Developer.
	//
	// - role_project_pe: O&M engineer.
	//
	// - role_project_deploy: Deployment.
	//
	// - role_project_guest: Visitor.
	//
	// - role_project_security: Security administrator.
	//
	// - role_project_tester: Experience user.
	//
	// - role_project_erd: Model designer.
	//
	// This parameter is required.
	//
	// example:
	//
	// role_project_guest
	RoleCode *string `json:"RoleCode,omitempty" xml:"RoleCode,omitempty"`
	// The ID of the user.
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
