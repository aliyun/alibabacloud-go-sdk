// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserGroupMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddCommand(v *AddUserGroupMemberRequestAddCommand) *AddUserGroupMemberRequest
	GetAddCommand() *AddUserGroupMemberRequestAddCommand
	SetOpTenantId(v int64) *AddUserGroupMemberRequest
	GetOpTenantId() *int64
}

type AddUserGroupMemberRequest struct {
	AddCommand *AddUserGroupMemberRequestAddCommand `json:"AddCommand,omitempty" xml:"AddCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s AddUserGroupMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s AddUserGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *AddUserGroupMemberRequest) GetAddCommand() *AddUserGroupMemberRequestAddCommand {
	return s.AddCommand
}

func (s *AddUserGroupMemberRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *AddUserGroupMemberRequest) SetAddCommand(v *AddUserGroupMemberRequestAddCommand) *AddUserGroupMemberRequest {
	s.AddCommand = v
	return s
}

func (s *AddUserGroupMemberRequest) SetOpTenantId(v int64) *AddUserGroupMemberRequest {
	s.OpTenantId = &v
	return s
}

func (s *AddUserGroupMemberRequest) Validate() error {
	return dara.Validate(s)
}

type AddUserGroupMemberRequestAddCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// 132331
	UserGroupId *string   `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	UserIdList  []*string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty" type:"Repeated"`
}

func (s AddUserGroupMemberRequestAddCommand) String() string {
	return dara.Prettify(s)
}

func (s AddUserGroupMemberRequestAddCommand) GoString() string {
	return s.String()
}

func (s *AddUserGroupMemberRequestAddCommand) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *AddUserGroupMemberRequestAddCommand) GetUserIdList() []*string {
	return s.UserIdList
}

func (s *AddUserGroupMemberRequestAddCommand) SetUserGroupId(v string) *AddUserGroupMemberRequestAddCommand {
	s.UserGroupId = &v
	return s
}

func (s *AddUserGroupMemberRequestAddCommand) SetUserIdList(v []*string) *AddUserGroupMemberRequestAddCommand {
	s.UserIdList = v
	return s
}

func (s *AddUserGroupMemberRequestAddCommand) Validate() error {
	return dara.Validate(s)
}
