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
	SetUserName(v string) *AddUserToGroupRequest
	GetUserName() *string
}

type AddUserToGroupRequest struct {
	// The name of the RAM user group.
	//
	// example:
	//
	// Dev-Team
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The name of the RAM user.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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

func (s *AddUserToGroupRequest) GetUserName() *string {
	return s.UserName
}

func (s *AddUserToGroupRequest) SetGroupName(v string) *AddUserToGroupRequest {
	s.GroupName = &v
	return s
}

func (s *AddUserToGroupRequest) SetUserName(v string) *AddUserToGroupRequest {
	s.UserName = &v
	return s
}

func (s *AddUserToGroupRequest) Validate() error {
	return dara.Validate(s)
}
