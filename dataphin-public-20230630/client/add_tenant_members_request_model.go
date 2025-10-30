// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTenantMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddCommand(v *AddTenantMembersRequestAddCommand) *AddTenantMembersRequest
	GetAddCommand() *AddTenantMembersRequestAddCommand
	SetOpTenantId(v int64) *AddTenantMembersRequest
	GetOpTenantId() *int64
}

type AddTenantMembersRequest struct {
	// This parameter is required.
	AddCommand *AddTenantMembersRequestAddCommand `json:"AddCommand,omitempty" xml:"AddCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s AddTenantMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s AddTenantMembersRequest) GoString() string {
	return s.String()
}

func (s *AddTenantMembersRequest) GetAddCommand() *AddTenantMembersRequestAddCommand {
	return s.AddCommand
}

func (s *AddTenantMembersRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *AddTenantMembersRequest) SetAddCommand(v *AddTenantMembersRequestAddCommand) *AddTenantMembersRequest {
	s.AddCommand = v
	return s
}

func (s *AddTenantMembersRequest) SetOpTenantId(v int64) *AddTenantMembersRequest {
	s.OpTenantId = &v
	return s
}

func (s *AddTenantMembersRequest) Validate() error {
	if s.AddCommand != nil {
		if err := s.AddCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddTenantMembersRequestAddCommand struct {
	// This parameter is required.
	UserList []*AddTenantMembersRequestAddCommandUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s AddTenantMembersRequestAddCommand) String() string {
	return dara.Prettify(s)
}

func (s AddTenantMembersRequestAddCommand) GoString() string {
	return s.String()
}

func (s *AddTenantMembersRequestAddCommand) GetUserList() []*AddTenantMembersRequestAddCommandUserList {
	return s.UserList
}

func (s *AddTenantMembersRequestAddCommand) SetUserList(v []*AddTenantMembersRequestAddCommandUserList) *AddTenantMembersRequestAddCommand {
	s.UserList = v
	return s
}

func (s *AddTenantMembersRequestAddCommand) Validate() error {
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

type AddTenantMembersRequestAddCommandUserList struct {
	// example:
	//
	// 1323241
	Id       *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	RoleList []*string `json:"RoleList,omitempty" xml:"RoleList,omitempty" type:"Repeated"`
}

func (s AddTenantMembersRequestAddCommandUserList) String() string {
	return dara.Prettify(s)
}

func (s AddTenantMembersRequestAddCommandUserList) GoString() string {
	return s.String()
}

func (s *AddTenantMembersRequestAddCommandUserList) GetId() *string {
	return s.Id
}

func (s *AddTenantMembersRequestAddCommandUserList) GetRoleList() []*string {
	return s.RoleList
}

func (s *AddTenantMembersRequestAddCommandUserList) SetId(v string) *AddTenantMembersRequestAddCommandUserList {
	s.Id = &v
	return s
}

func (s *AddTenantMembersRequestAddCommandUserList) SetRoleList(v []*string) *AddTenantMembersRequestAddCommandUserList {
	s.RoleList = v
	return s
}

func (s *AddTenantMembersRequestAddCommandUserList) Validate() error {
	return dara.Validate(s)
}
