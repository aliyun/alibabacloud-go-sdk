// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserToDesktopOversoldUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddUserAmount(v int32) *AddUserToDesktopOversoldUserGroupRequest
	GetAddUserAmount() *int32
	SetEndUserId(v string) *AddUserToDesktopOversoldUserGroupRequest
	GetEndUserId() *string
	SetOversoldGroupId(v string) *AddUserToDesktopOversoldUserGroupRequest
	GetOversoldGroupId() *string
	SetUserGroupId(v string) *AddUserToDesktopOversoldUserGroupRequest
	GetUserGroupId() *string
}

type AddUserToDesktopOversoldUserGroupRequest struct {
	AddUserAmount   *int32  `json:"AddUserAmount,omitempty" xml:"AddUserAmount,omitempty"`
	EndUserId       *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	OversoldGroupId *string `json:"OversoldGroupId,omitempty" xml:"OversoldGroupId,omitempty"`
	UserGroupId     *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s AddUserToDesktopOversoldUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AddUserToDesktopOversoldUserGroupRequest) GoString() string {
	return s.String()
}

func (s *AddUserToDesktopOversoldUserGroupRequest) GetAddUserAmount() *int32 {
	return s.AddUserAmount
}

func (s *AddUserToDesktopOversoldUserGroupRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *AddUserToDesktopOversoldUserGroupRequest) GetOversoldGroupId() *string {
	return s.OversoldGroupId
}

func (s *AddUserToDesktopOversoldUserGroupRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *AddUserToDesktopOversoldUserGroupRequest) SetAddUserAmount(v int32) *AddUserToDesktopOversoldUserGroupRequest {
	s.AddUserAmount = &v
	return s
}

func (s *AddUserToDesktopOversoldUserGroupRequest) SetEndUserId(v string) *AddUserToDesktopOversoldUserGroupRequest {
	s.EndUserId = &v
	return s
}

func (s *AddUserToDesktopOversoldUserGroupRequest) SetOversoldGroupId(v string) *AddUserToDesktopOversoldUserGroupRequest {
	s.OversoldGroupId = &v
	return s
}

func (s *AddUserToDesktopOversoldUserGroupRequest) SetUserGroupId(v string) *AddUserToDesktopOversoldUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *AddUserToDesktopOversoldUserGroupRequest) Validate() error {
	return dara.Validate(s)
}
