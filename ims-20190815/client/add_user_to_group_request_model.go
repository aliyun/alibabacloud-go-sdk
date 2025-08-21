// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserToGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *AddUserToGroupRequest
	GetGroupName() *string
	SetUserPrincipalName(v string) *AddUserToGroupRequest
	GetUserPrincipalName() *string
}

type AddUserToGroupRequest struct {
	// The name of the RAM user group.
	//
	// example:
	//
	// Test-Team
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The logon name of the RAM user.
	//
	// This parameter is required.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s AddUserToGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AddUserToGroupRequest) GoString() string {
	return s.String()
}

func (s *AddUserToGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *AddUserToGroupRequest) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *AddUserToGroupRequest) SetGroupName(v string) *AddUserToGroupRequest {
	s.GroupName = &v
	return s
}

func (s *AddUserToGroupRequest) SetUserPrincipalName(v string) *AddUserToGroupRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *AddUserToGroupRequest) Validate() error {
	return dara.Validate(s)
}
