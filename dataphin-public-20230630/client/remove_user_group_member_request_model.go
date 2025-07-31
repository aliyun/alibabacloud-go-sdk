// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserGroupMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *RemoveUserGroupMemberRequest
	GetOpTenantId() *int64
	SetRemoveCommand(v *RemoveUserGroupMemberRequestRemoveCommand) *RemoveUserGroupMemberRequest
	GetRemoveCommand() *RemoveUserGroupMemberRequestRemoveCommand
}

type RemoveUserGroupMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	RemoveCommand *RemoveUserGroupMemberRequestRemoveCommand `json:"RemoveCommand,omitempty" xml:"RemoveCommand,omitempty" type:"Struct"`
}

func (s RemoveUserGroupMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserGroupMemberRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *RemoveUserGroupMemberRequest) GetRemoveCommand() *RemoveUserGroupMemberRequestRemoveCommand {
	return s.RemoveCommand
}

func (s *RemoveUserGroupMemberRequest) SetOpTenantId(v int64) *RemoveUserGroupMemberRequest {
	s.OpTenantId = &v
	return s
}

func (s *RemoveUserGroupMemberRequest) SetRemoveCommand(v *RemoveUserGroupMemberRequestRemoveCommand) *RemoveUserGroupMemberRequest {
	s.RemoveCommand = v
	return s
}

func (s *RemoveUserGroupMemberRequest) Validate() error {
	return dara.Validate(s)
}

type RemoveUserGroupMemberRequestRemoveCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// 2311
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// This parameter is required.
	UserIdList []*string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty" type:"Repeated"`
}

func (s RemoveUserGroupMemberRequestRemoveCommand) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserGroupMemberRequestRemoveCommand) GoString() string {
	return s.String()
}

func (s *RemoveUserGroupMemberRequestRemoveCommand) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *RemoveUserGroupMemberRequestRemoveCommand) GetUserIdList() []*string {
	return s.UserIdList
}

func (s *RemoveUserGroupMemberRequestRemoveCommand) SetUserGroupId(v string) *RemoveUserGroupMemberRequestRemoveCommand {
	s.UserGroupId = &v
	return s
}

func (s *RemoveUserGroupMemberRequestRemoveCommand) SetUserIdList(v []*string) *RemoveUserGroupMemberRequestRemoveCommand {
	s.UserIdList = v
	return s
}

func (s *RemoveUserGroupMemberRequestRemoveCommand) Validate() error {
	return dara.Validate(s)
}
