// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateProjectMemberRequest
	GetClientToken() *string
	SetProjectId(v int64) *CreateProjectMemberRequest
	GetProjectId() *int64
	SetRoleCode(v string) *CreateProjectMemberRequest
	GetRoleCode() *string
	SetUserId(v string) *CreateProjectMemberRequest
	GetUserId() *string
}

type CreateProjectMemberRequest struct {
	// The idempotence string. A UUID is recommended.
	//
	// example:
	//
	// 1AFAE64E-D1BE-432B-A9****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the DataWorks workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 27
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The role code. This parameter is optional. If this parameter is specified, the user is also added to the specified role.
	//
	// example:
	//
	// role_project_guest
	RoleCode *string `json:"RoleCode,omitempty" xml:"RoleCode,omitempty"`
	// The ID of the user to be added.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateProjectMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectMemberRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectMemberRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateProjectMemberRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateProjectMemberRequest) GetRoleCode() *string {
	return s.RoleCode
}

func (s *CreateProjectMemberRequest) GetUserId() *string {
	return s.UserId
}

func (s *CreateProjectMemberRequest) SetClientToken(v string) *CreateProjectMemberRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateProjectMemberRequest) SetProjectId(v int64) *CreateProjectMemberRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateProjectMemberRequest) SetRoleCode(v string) *CreateProjectMemberRequest {
	s.RoleCode = &v
	return s
}

func (s *CreateProjectMemberRequest) SetUserId(v string) *CreateProjectMemberRequest {
	s.UserId = &v
	return s
}

func (s *CreateProjectMemberRequest) Validate() error {
	return dara.Validate(s)
}
