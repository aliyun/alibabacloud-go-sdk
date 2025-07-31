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
}

type AddUserGroupMemberShrinkRequest struct {
	AddCommandShrink *string `json:"AddCommand,omitempty" xml:"AddCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
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

func (s *AddUserGroupMemberShrinkRequest) SetAddCommandShrink(v string) *AddUserGroupMemberShrinkRequest {
	s.AddCommandShrink = &v
	return s
}

func (s *AddUserGroupMemberShrinkRequest) SetOpTenantId(v int64) *AddUserGroupMemberShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *AddUserGroupMemberShrinkRequest) Validate() error {
	return dara.Validate(s)
}
