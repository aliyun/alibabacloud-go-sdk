// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserFromDesktopOversoldUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndUserId(v string) *RemoveUserFromDesktopOversoldUserGroupRequest
	GetEndUserId() *string
	SetOversoldGroupId(v string) *RemoveUserFromDesktopOversoldUserGroupRequest
	GetOversoldGroupId() *string
	SetUserDesktopId(v string) *RemoveUserFromDesktopOversoldUserGroupRequest
	GetUserDesktopId() *string
	SetUserGroupId(v string) *RemoveUserFromDesktopOversoldUserGroupRequest
	GetUserGroupId() *string
}

type RemoveUserFromDesktopOversoldUserGroupRequest struct {
	EndUserId       *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	OversoldGroupId *string `json:"OversoldGroupId,omitempty" xml:"OversoldGroupId,omitempty"`
	UserDesktopId   *string `json:"UserDesktopId,omitempty" xml:"UserDesktopId,omitempty"`
	UserGroupId     *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s RemoveUserFromDesktopOversoldUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserFromDesktopOversoldUserGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserFromDesktopOversoldUserGroupRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *RemoveUserFromDesktopOversoldUserGroupRequest) GetOversoldGroupId() *string {
	return s.OversoldGroupId
}

func (s *RemoveUserFromDesktopOversoldUserGroupRequest) GetUserDesktopId() *string {
	return s.UserDesktopId
}

func (s *RemoveUserFromDesktopOversoldUserGroupRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *RemoveUserFromDesktopOversoldUserGroupRequest) SetEndUserId(v string) *RemoveUserFromDesktopOversoldUserGroupRequest {
	s.EndUserId = &v
	return s
}

func (s *RemoveUserFromDesktopOversoldUserGroupRequest) SetOversoldGroupId(v string) *RemoveUserFromDesktopOversoldUserGroupRequest {
	s.OversoldGroupId = &v
	return s
}

func (s *RemoveUserFromDesktopOversoldUserGroupRequest) SetUserDesktopId(v string) *RemoveUserFromDesktopOversoldUserGroupRequest {
	s.UserDesktopId = &v
	return s
}

func (s *RemoveUserFromDesktopOversoldUserGroupRequest) SetUserGroupId(v string) *RemoveUserFromDesktopOversoldUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *RemoveUserFromDesktopOversoldUserGroupRequest) Validate() error {
	return dara.Validate(s)
}
