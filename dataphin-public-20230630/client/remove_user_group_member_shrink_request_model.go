// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserGroupMemberShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *RemoveUserGroupMemberShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *RemoveUserGroupMemberShrinkRequest
	GetOpUserId() *string
	SetRemoveCommandShrink(v string) *RemoveUserGroupMemberShrinkRequest
	GetRemoveCommandShrink() *string
}

type RemoveUserGroupMemberShrinkRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The command to remove user group members.
	//
	// This parameter is required.
	RemoveCommandShrink *string `json:"RemoveCommand,omitempty" xml:"RemoveCommand,omitempty"`
}

func (s RemoveUserGroupMemberShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserGroupMemberShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserGroupMemberShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *RemoveUserGroupMemberShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *RemoveUserGroupMemberShrinkRequest) GetRemoveCommandShrink() *string {
	return s.RemoveCommandShrink
}

func (s *RemoveUserGroupMemberShrinkRequest) SetOpTenantId(v int64) *RemoveUserGroupMemberShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *RemoveUserGroupMemberShrinkRequest) SetOpUserId(v string) *RemoveUserGroupMemberShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *RemoveUserGroupMemberShrinkRequest) SetRemoveCommandShrink(v string) *RemoveUserGroupMemberShrinkRequest {
	s.RemoveCommandShrink = &v
	return s
}

func (s *RemoveUserGroupMemberShrinkRequest) Validate() error {
	return dara.Validate(s)
}
