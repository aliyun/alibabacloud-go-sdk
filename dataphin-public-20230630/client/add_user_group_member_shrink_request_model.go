// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserGroupMemberShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddCommandShrink(v string) *AddUserGroupMemberShrinkRequest
	GetAddCommandShrink() *string
	SetOpTenantId(v int64) *AddUserGroupMemberShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *AddUserGroupMemberShrinkRequest
	GetOpUserId() *string
}

type AddUserGroupMemberShrinkRequest struct {
	// The command to add user group members.
	AddCommandShrink *string `json:"AddCommand,omitempty" xml:"AddCommand,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s AddUserGroupMemberShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddUserGroupMemberShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddUserGroupMemberShrinkRequest) GetAddCommandShrink() *string {
	return s.AddCommandShrink
}

func (s *AddUserGroupMemberShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *AddUserGroupMemberShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *AddUserGroupMemberShrinkRequest) SetAddCommandShrink(v string) *AddUserGroupMemberShrinkRequest {
	s.AddCommandShrink = &v
	return s
}

func (s *AddUserGroupMemberShrinkRequest) SetOpTenantId(v int64) *AddUserGroupMemberShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *AddUserGroupMemberShrinkRequest) SetOpUserId(v string) *AddUserGroupMemberShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *AddUserGroupMemberShrinkRequest) Validate() error {
	return dara.Validate(s)
}
