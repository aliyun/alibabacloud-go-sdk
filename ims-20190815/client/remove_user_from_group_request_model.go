// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserFromGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *RemoveUserFromGroupRequest
	GetGroupName() *string
	SetUserPrincipalName(v string) *RemoveUserFromGroupRequest
	GetUserPrincipalName() *string
}

type RemoveUserFromGroupRequest struct {
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
	// alice@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s RemoveUserFromGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserFromGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserFromGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *RemoveUserFromGroupRequest) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *RemoveUserFromGroupRequest) SetGroupName(v string) *RemoveUserFromGroupRequest {
	s.GroupName = &v
	return s
}

func (s *RemoveUserFromGroupRequest) SetUserPrincipalName(v string) *RemoveUserFromGroupRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *RemoveUserFromGroupRequest) Validate() error {
	return dara.Validate(s)
}
