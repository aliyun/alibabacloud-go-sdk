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
}

type AddProjectMemberRequest struct {
	// This parameter is required.
	AddCommand *AddProjectMemberRequestAddCommand `json:"AddCommand,omitempty" xml:"AddCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 711833
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
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

func (s *AddProjectMemberRequest) Validate() error {
	if s.AddCommand != nil {
		if err := s.AddCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddProjectMemberRequestAddCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
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
	// This parameter is required.
	RoleList []*int32 `json:"RoleList,omitempty" xml:"RoleList,omitempty" type:"Repeated"`
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

func (s *AddProjectMemberRequestAddCommandUserList) GetRoleList() []*int32 {
	return s.RoleList
}

func (s *AddProjectMemberRequestAddCommandUserList) GetUserId() *string {
	return s.UserId
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
