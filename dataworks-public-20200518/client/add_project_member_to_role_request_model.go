// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddProjectMemberToRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AddProjectMemberToRoleRequest
	GetClientToken() *string
	SetProjectId(v int64) *AddProjectMemberToRoleRequest
	GetProjectId() *int64
	SetRoleCode(v string) *AddProjectMemberToRoleRequest
	GetRoleCode() *string
	SetUserId(v string) *AddProjectMemberToRoleRequest
	GetUserId() *string
}

type AddProjectMemberToRoleRequest struct {
	// The client token that is used to ensure the idempotence of the request. We recommend that you set this parameter to a UUID. This parameter is used to uniquely identify the API operation call.
	//
	// example:
	//
	// 1AFAE64E-D1BE-432B-A9*****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The DataWorks workspace ID. You can call the [ListProjects](https://help.aliyun.com/document_detail/2780068.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 27
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The code of the role that you want to assign to a member of the DataWorks workspace. You can call the [ListProjectRoles](https://help.aliyun.com/document_detail/2780079.html) operation to query the codes of all roles in a DataWorks workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// role_project_guest
	RoleCode *string `json:"RoleCode,omitempty" xml:"RoleCode,omitempty"`
	// The Alibaba Cloud account ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and move the pointer over the profile picture in the upper-right corner to view the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AddProjectMemberToRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s AddProjectMemberToRoleRequest) GoString() string {
	return s.String()
}

func (s *AddProjectMemberToRoleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddProjectMemberToRoleRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *AddProjectMemberToRoleRequest) GetRoleCode() *string {
	return s.RoleCode
}

func (s *AddProjectMemberToRoleRequest) GetUserId() *string {
	return s.UserId
}

func (s *AddProjectMemberToRoleRequest) SetClientToken(v string) *AddProjectMemberToRoleRequest {
	s.ClientToken = &v
	return s
}

func (s *AddProjectMemberToRoleRequest) SetProjectId(v int64) *AddProjectMemberToRoleRequest {
	s.ProjectId = &v
	return s
}

func (s *AddProjectMemberToRoleRequest) SetRoleCode(v string) *AddProjectMemberToRoleRequest {
	s.RoleCode = &v
	return s
}

func (s *AddProjectMemberToRoleRequest) SetUserId(v string) *AddProjectMemberToRoleRequest {
	s.UserId = &v
	return s
}

func (s *AddProjectMemberToRoleRequest) Validate() error {
	return dara.Validate(s)
}
