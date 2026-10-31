// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddProjectMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddCommand(v *AddProjectMemberRequestAddCommand) *AddProjectMemberRequest
	GetAddCommand() *AddProjectMemberRequestAddCommand
	SetId(v int64) *AddProjectMemberRequest
	GetId() *int64
	SetOpTenantId(v int64) *AddProjectMemberRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *AddProjectMemberRequest
	GetOpUserId() *string
}

type AddProjectMemberRequest struct {
	// The command to add members.
	//
	// This parameter is required.
	AddCommand *AddProjectMemberRequestAddCommand `json:"AddCommand,omitempty" xml:"AddCommand,omitempty" type:"Struct"`
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 711833
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The operator user ID.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s AddProjectMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s AddProjectMemberRequest) GoString() string {
	return s.String()
}

func (s *AddProjectMemberRequest) GetAddCommand() *AddProjectMemberRequestAddCommand {
	return s.AddCommand
}

func (s *AddProjectMemberRequest) GetId() *int64 {
	return s.Id
}

func (s *AddProjectMemberRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *AddProjectMemberRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *AddProjectMemberRequest) SetAddCommand(v *AddProjectMemberRequestAddCommand) *AddProjectMemberRequest {
	s.AddCommand = v
	return s
}

func (s *AddProjectMemberRequest) SetId(v int64) *AddProjectMemberRequest {
	s.Id = &v
	return s
}

func (s *AddProjectMemberRequest) SetOpTenantId(v int64) *AddProjectMemberRequest {
	s.OpTenantId = &v
	return s
}

func (s *AddProjectMemberRequest) SetOpUserId(v string) *AddProjectMemberRequest {
	s.OpUserId = &v
	return s
}

func (s *AddProjectMemberRequest) Validate() error {
	if s.AddCommand != nil {
		if err := s.AddCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddProjectMemberRequestAddCommand struct {
	// The environment identifier. Valid values: DEV and PROD.
	//
	// This parameter is required.
	//
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The user members.
	//
	// This parameter is required.
	UserList []*AddProjectMemberRequestAddCommandUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s AddProjectMemberRequestAddCommand) String() string {
	return dara.Prettify(s)
}

func (s AddProjectMemberRequestAddCommand) GoString() string {
	return s.String()
}

func (s *AddProjectMemberRequestAddCommand) GetEnv() *string {
	return s.Env
}

func (s *AddProjectMemberRequestAddCommand) GetUserList() []*AddProjectMemberRequestAddCommandUserList {
	return s.UserList
}

func (s *AddProjectMemberRequestAddCommand) SetEnv(v string) *AddProjectMemberRequestAddCommand {
	s.Env = &v
	return s
}

func (s *AddProjectMemberRequestAddCommand) SetUserList(v []*AddProjectMemberRequestAddCommandUserList) *AddProjectMemberRequestAddCommand {
	s.UserList = v
	return s
}

func (s *AddProjectMemberRequestAddCommand) Validate() error {
	if s.UserList != nil {
		for _, item := range s.UserList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddProjectMemberRequestAddCommandUserList struct {
	// The list of role codes for members. System and custom project roles are supported and can be obtained from ListProjectRoles. Only one role is supported. If this parameter is not empty, it takes precedence over RoleList. You must specify at least one of RoleList and RoleCodeList.
	RoleCodeList []*string `json:"RoleCodeList,omitempty" xml:"RoleCodeList,omitempty" type:"Repeated"`
	// The member role. Valid values: 1 (Project administrator), 2 (Developer), 3 (Guest), 4 (Analyst), and 5 (O&M engineer).
	RoleList []*int32 `json:"RoleList,omitempty" xml:"RoleList,omitempty" type:"Repeated"`
	// The user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30012011
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AddProjectMemberRequestAddCommandUserList) String() string {
	return dara.Prettify(s)
}

func (s AddProjectMemberRequestAddCommandUserList) GoString() string {
	return s.String()
}

func (s *AddProjectMemberRequestAddCommandUserList) GetRoleCodeList() []*string {
	return s.RoleCodeList
}

func (s *AddProjectMemberRequestAddCommandUserList) GetRoleList() []*int32 {
	return s.RoleList
}

func (s *AddProjectMemberRequestAddCommandUserList) GetUserId() *string {
	return s.UserId
}

func (s *AddProjectMemberRequestAddCommandUserList) SetRoleCodeList(v []*string) *AddProjectMemberRequestAddCommandUserList {
	s.RoleCodeList = v
	return s
}

func (s *AddProjectMemberRequestAddCommandUserList) SetRoleList(v []*int32) *AddProjectMemberRequestAddCommandUserList {
	s.RoleList = v
	return s
}

func (s *AddProjectMemberRequestAddCommandUserList) SetUserId(v string) *AddProjectMemberRequestAddCommandUserList {
	s.UserId = &v
	return s
}

func (s *AddProjectMemberRequestAddCommandUserList) Validate() error {
	return dara.Validate(s)
}
