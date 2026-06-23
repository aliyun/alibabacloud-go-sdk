// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProjectMember(v *GetProjectMemberResponseBodyProjectMember) *GetProjectMemberResponseBody
	GetProjectMember() *GetProjectMemberResponseBodyProjectMember
	SetRequestId(v string) *GetProjectMemberResponseBody
	GetRequestId() *string
}

type GetProjectMemberResponseBody struct {
	// The details of the Workspace member.
	ProjectMember *GetProjectMemberResponseBodyProjectMember `json:"ProjectMember,omitempty" xml:"ProjectMember,omitempty" type:"Struct"`
	// The request ID. Use this ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 8abcb91f-d266-4073-b907-2ed670378ed1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetProjectMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProjectMemberResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectMemberResponseBody) GetProjectMember() *GetProjectMemberResponseBodyProjectMember {
	return s.ProjectMember
}

func (s *GetProjectMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProjectMemberResponseBody) SetProjectMember(v *GetProjectMemberResponseBodyProjectMember) *GetProjectMemberResponseBody {
	s.ProjectMember = v
	return s
}

func (s *GetProjectMemberResponseBody) SetRequestId(v string) *GetProjectMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectMemberResponseBody) Validate() error {
	if s.ProjectMember != nil {
		if err := s.ProjectMember.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetProjectMemberResponseBodyProjectMember struct {
	// The ID of the Workspace.
	//
	// example:
	//
	// 88757
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Roles assigned to the Workspace member.
	Roles []*GetProjectMemberResponseBodyProjectMemberRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// The status of the Workspace member.
	//
	// - Normal: The member is active.
	//
	// - Disabled: The member is disabled.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// 123422344899
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The name of the user.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetProjectMemberResponseBodyProjectMember) String() string {
	return dara.Prettify(s)
}

func (s GetProjectMemberResponseBodyProjectMember) GoString() string {
	return s.String()
}

func (s *GetProjectMemberResponseBodyProjectMember) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetProjectMemberResponseBodyProjectMember) GetRoles() []*GetProjectMemberResponseBodyProjectMemberRoles {
	return s.Roles
}

func (s *GetProjectMemberResponseBodyProjectMember) GetStatus() *string {
	return s.Status
}

func (s *GetProjectMemberResponseBodyProjectMember) GetUserId() *string {
	return s.UserId
}

func (s *GetProjectMemberResponseBodyProjectMember) GetUserName() *string {
	return s.UserName
}

func (s *GetProjectMemberResponseBodyProjectMember) SetProjectId(v int64) *GetProjectMemberResponseBodyProjectMember {
	s.ProjectId = &v
	return s
}

func (s *GetProjectMemberResponseBodyProjectMember) SetRoles(v []*GetProjectMemberResponseBodyProjectMemberRoles) *GetProjectMemberResponseBodyProjectMember {
	s.Roles = v
	return s
}

func (s *GetProjectMemberResponseBodyProjectMember) SetStatus(v string) *GetProjectMemberResponseBodyProjectMember {
	s.Status = &v
	return s
}

func (s *GetProjectMemberResponseBodyProjectMember) SetUserId(v string) *GetProjectMemberResponseBodyProjectMember {
	s.UserId = &v
	return s
}

func (s *GetProjectMemberResponseBodyProjectMember) SetUserName(v string) *GetProjectMemberResponseBodyProjectMember {
	s.UserName = &v
	return s
}

func (s *GetProjectMemberResponseBodyProjectMember) Validate() error {
	if s.Roles != nil {
		for _, item := range s.Roles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetProjectMemberResponseBodyProjectMemberRoles struct {
	// The code of the Workspace role.
	//
	// The built-in system roles in a DataWorks Workspace include:
	//
	// - role_project_admin: Workspace Administrator
	//
	// - role_project_dev: Developer
	//
	// - role_project_dg_admin: Data Governance Administrator
	//
	// - role_project_guest: Guest
	//
	// - role_project_security: Security Administrator
	//
	// - role_project_deploy: Deployment
	//
	// - role_project_owner: Workspace Owner
	//
	// - role_project_data_analyst: Data Analyst
	//
	// - role_project_pe: O\\&M (Operations & Maintenance)
	//
	// - role_project_erd: Model Designer
	//
	// example:
	//
	// role_project_guest
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The name of the Workspace role.
	//
	// example:
	//
	// Visitors
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the Workspace role.
	//
	// - UserCustom: A user-defined role.
	//
	// - System: A built-in System Role.
	//
	// example:
	//
	// System
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetProjectMemberResponseBodyProjectMemberRoles) String() string {
	return dara.Prettify(s)
}

func (s GetProjectMemberResponseBodyProjectMemberRoles) GoString() string {
	return s.String()
}

func (s *GetProjectMemberResponseBodyProjectMemberRoles) GetCode() *string {
	return s.Code
}

func (s *GetProjectMemberResponseBodyProjectMemberRoles) GetName() *string {
	return s.Name
}

func (s *GetProjectMemberResponseBodyProjectMemberRoles) GetType() *string {
	return s.Type
}

func (s *GetProjectMemberResponseBodyProjectMemberRoles) SetCode(v string) *GetProjectMemberResponseBodyProjectMemberRoles {
	s.Code = &v
	return s
}

func (s *GetProjectMemberResponseBodyProjectMemberRoles) SetName(v string) *GetProjectMemberResponseBodyProjectMemberRoles {
	s.Name = &v
	return s
}

func (s *GetProjectMemberResponseBodyProjectMemberRoles) SetType(v string) *GetProjectMemberResponseBodyProjectMemberRoles {
	s.Type = &v
	return s
}

func (s *GetProjectMemberResponseBodyProjectMemberRoles) Validate() error {
	return dara.Validate(s)
}
