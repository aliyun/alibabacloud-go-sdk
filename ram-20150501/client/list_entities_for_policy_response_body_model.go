// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEntitiesForPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroups(v *ListEntitiesForPolicyResponseBodyGroups) *ListEntitiesForPolicyResponseBody
	GetGroups() *ListEntitiesForPolicyResponseBodyGroups
	SetRequestId(v string) *ListEntitiesForPolicyResponseBody
	GetRequestId() *string
	SetRoles(v *ListEntitiesForPolicyResponseBodyRoles) *ListEntitiesForPolicyResponseBody
	GetRoles() *ListEntitiesForPolicyResponseBodyRoles
	SetUsers(v *ListEntitiesForPolicyResponseBodyUsers) *ListEntitiesForPolicyResponseBody
	GetUsers() *ListEntitiesForPolicyResponseBodyUsers
}

type ListEntitiesForPolicyResponseBody struct {
	Groups *ListEntitiesForPolicyResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7B8A4E7D-6CFF-471D-84DF-195A7A241ECB
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Roles     *ListEntitiesForPolicyResponseBodyRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Struct"`
	Users     *ListEntitiesForPolicyResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
}

func (s ListEntitiesForPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEntitiesForPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ListEntitiesForPolicyResponseBody) GetGroups() *ListEntitiesForPolicyResponseBodyGroups {
	return s.Groups
}

func (s *ListEntitiesForPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEntitiesForPolicyResponseBody) GetRoles() *ListEntitiesForPolicyResponseBodyRoles {
	return s.Roles
}

func (s *ListEntitiesForPolicyResponseBody) GetUsers() *ListEntitiesForPolicyResponseBodyUsers {
	return s.Users
}

func (s *ListEntitiesForPolicyResponseBody) SetGroups(v *ListEntitiesForPolicyResponseBodyGroups) *ListEntitiesForPolicyResponseBody {
	s.Groups = v
	return s
}

func (s *ListEntitiesForPolicyResponseBody) SetRequestId(v string) *ListEntitiesForPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEntitiesForPolicyResponseBody) SetRoles(v *ListEntitiesForPolicyResponseBodyRoles) *ListEntitiesForPolicyResponseBody {
	s.Roles = v
	return s
}

func (s *ListEntitiesForPolicyResponseBody) SetUsers(v *ListEntitiesForPolicyResponseBodyUsers) *ListEntitiesForPolicyResponseBody {
	s.Users = v
	return s
}

func (s *ListEntitiesForPolicyResponseBody) Validate() error {
	if s.Groups != nil {
		if err := s.Groups.Validate(); err != nil {
			return err
		}
	}
	if s.Roles != nil {
		if err := s.Roles.Validate(); err != nil {
			return err
		}
	}
	if s.Users != nil {
		if err := s.Users.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListEntitiesForPolicyResponseBodyGroups struct {
	Group []*ListEntitiesForPolicyResponseBodyGroupsGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Repeated"`
}

func (s ListEntitiesForPolicyResponseBodyGroups) String() string {
	return dara.Prettify(s)
}

func (s ListEntitiesForPolicyResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *ListEntitiesForPolicyResponseBodyGroups) GetGroup() []*ListEntitiesForPolicyResponseBodyGroupsGroup {
	return s.Group
}

func (s *ListEntitiesForPolicyResponseBodyGroups) SetGroup(v []*ListEntitiesForPolicyResponseBodyGroupsGroup) *ListEntitiesForPolicyResponseBodyGroups {
	s.Group = v
	return s
}

func (s *ListEntitiesForPolicyResponseBodyGroups) Validate() error {
	if s.Group != nil {
		for _, item := range s.Group {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEntitiesForPolicyResponseBodyGroupsGroup struct {
	AttachDate *string `json:"AttachDate,omitempty" xml:"AttachDate,omitempty"`
	Comments   *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	GroupName  *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s ListEntitiesForPolicyResponseBodyGroupsGroup) String() string {
	return dara.Prettify(s)
}

func (s ListEntitiesForPolicyResponseBodyGroupsGroup) GoString() string {
	return s.String()
}

func (s *ListEntitiesForPolicyResponseBodyGroupsGroup) GetAttachDate() *string {
	return s.AttachDate
}

func (s *ListEntitiesForPolicyResponseBodyGroupsGroup) GetComments() *string {
	return s.Comments
}

func (s *ListEntitiesForPolicyResponseBodyGroupsGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *ListEntitiesForPolicyResponseBodyGroupsGroup) SetAttachDate(v string) *ListEntitiesForPolicyResponseBodyGroupsGroup {
	s.AttachDate = &v
	return s
}

func (s *ListEntitiesForPolicyResponseBodyGroupsGroup) SetComments(v string) *ListEntitiesForPolicyResponseBodyGroupsGroup {
	s.Comments = &v
	return s
}

func (s *ListEntitiesForPolicyResponseBodyGroupsGroup) SetGroupName(v string) *ListEntitiesForPolicyResponseBodyGroupsGroup {
	s.GroupName = &v
	return s
}

func (s *ListEntitiesForPolicyResponseBodyGroupsGroup) Validate() error {
	return dara.Validate(s)
}

type ListEntitiesForPolicyResponseBodyRoles struct {
	Role []*ListEntitiesForPolicyResponseBodyRolesRole `json:"Role,omitempty" xml:"Role,omitempty" type:"Repeated"`
}

func (s ListEntitiesForPolicyResponseBodyRoles) String() string {
	return dara.Prettify(s)
}

func (s ListEntitiesForPolicyResponseBodyRoles) GoString() string {
	return s.String()
}

func (s *ListEntitiesForPolicyResponseBodyRoles) GetRole() []*ListEntitiesForPolicyResponseBodyRolesRole {
	return s.Role
}

func (s *ListEntitiesForPolicyResponseBodyRoles) SetRole(v []*ListEntitiesForPolicyResponseBodyRolesRole) *ListEntitiesForPolicyResponseBodyRoles {
	s.Role = v
	return s
}

func (s *ListEntitiesForPolicyResponseBodyRoles) Validate() error {
	if s.Role != nil {
		for _, item := range s.Role {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEntitiesForPolicyResponseBodyRolesRole struct {
	Arn         *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	AttachDate  *string `json:"AttachDate,omitempty" xml:"AttachDate,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RoleId      *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	RoleName    *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s ListEntitiesForPolicyResponseBodyRolesRole) String() string {
	return dara.Prettify(s)
}

func (s ListEntitiesForPolicyResponseBodyRolesRole) GoString() string {
	return s.String()
}

func (s *ListEntitiesForPolicyResponseBodyRolesRole) GetArn() *string {
	return s.Arn
}

func (s *ListEntitiesForPolicyResponseBodyRolesRole) GetAttachDate() *string {
	return s.AttachDate
}

func (s *ListEntitiesForPolicyResponseBodyRolesRole) GetDescription() *string {
	return s.Description
}

func (s *ListEntitiesForPolicyResponseBodyRolesRole) GetRoleId() *string {
	return s.RoleId
}

func (s *ListEntitiesForPolicyResponseBodyRolesRole) GetRoleName() *string {
	return s.RoleName
}

func (s *ListEntitiesForPolicyResponseBodyRolesRole) SetArn(v string) *ListEntitiesForPolicyResponseBodyRolesRole {
	s.Arn = &v
	return s
}

func (s *ListEntitiesForPolicyResponseBodyRolesRole) SetAttachDate(v string) *ListEntitiesForPolicyResponseBodyRolesRole {
	s.AttachDate = &v
	return s
}

func (s *ListEntitiesForPolicyResponseBodyRolesRole) SetDescription(v string) *ListEntitiesForPolicyResponseBodyRolesRole {
	s.Description = &v
	return s
}

func (s *ListEntitiesForPolicyResponseBodyRolesRole) SetRoleId(v string) *ListEntitiesForPolicyResponseBodyRolesRole {
	s.RoleId = &v
	return s
}

func (s *ListEntitiesForPolicyResponseBodyRolesRole) SetRoleName(v string) *ListEntitiesForPolicyResponseBodyRolesRole {
	s.RoleName = &v
	return s
}

func (s *ListEntitiesForPolicyResponseBodyRolesRole) Validate() error {
	return dara.Validate(s)
}

type ListEntitiesForPolicyResponseBodyUsers struct {
	User []*ListEntitiesForPolicyResponseBodyUsersUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s ListEntitiesForPolicyResponseBodyUsers) String() string {
	return dara.Prettify(s)
}

func (s ListEntitiesForPolicyResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListEntitiesForPolicyResponseBodyUsers) GetUser() []*ListEntitiesForPolicyResponseBodyUsersUser {
	return s.User
}

func (s *ListEntitiesForPolicyResponseBodyUsers) SetUser(v []*ListEntitiesForPolicyResponseBodyUsersUser) *ListEntitiesForPolicyResponseBodyUsers {
	s.User = v
	return s
}

func (s *ListEntitiesForPolicyResponseBodyUsers) Validate() error {
	if s.User != nil {
		for _, item := range s.User {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEntitiesForPolicyResponseBodyUsersUser struct {
	AttachDate  *string `json:"AttachDate,omitempty" xml:"AttachDate,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName    *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListEntitiesForPolicyResponseBodyUsersUser) String() string {
	return dara.Prettify(s)
}

func (s ListEntitiesForPolicyResponseBodyUsersUser) GoString() string {
	return s.String()
}

func (s *ListEntitiesForPolicyResponseBodyUsersUser) GetAttachDate() *string {
	return s.AttachDate
}

func (s *ListEntitiesForPolicyResponseBodyUsersUser) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListEntitiesForPolicyResponseBodyUsersUser) GetUserId() *string {
	return s.UserId
}

func (s *ListEntitiesForPolicyResponseBodyUsersUser) GetUserName() *string {
	return s.UserName
}

func (s *ListEntitiesForPolicyResponseBodyUsersUser) SetAttachDate(v string) *ListEntitiesForPolicyResponseBodyUsersUser {
	s.AttachDate = &v
	return s
}

func (s *ListEntitiesForPolicyResponseBodyUsersUser) SetDisplayName(v string) *ListEntitiesForPolicyResponseBodyUsersUser {
	s.DisplayName = &v
	return s
}

func (s *ListEntitiesForPolicyResponseBodyUsersUser) SetUserId(v string) *ListEntitiesForPolicyResponseBodyUsersUser {
	s.UserId = &v
	return s
}

func (s *ListEntitiesForPolicyResponseBodyUsersUser) SetUserName(v string) *ListEntitiesForPolicyResponseBodyUsersUser {
	s.UserName = &v
	return s
}

func (s *ListEntitiesForPolicyResponseBodyUsersUser) Validate() error {
	return dara.Validate(s)
}
