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
	SetUserName(v string) *RemoveUserFromGroupRequest
	GetUserName() *string
}

type RemoveUserFromGroupRequest struct {
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

func (s RemoveUserFromGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserFromGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserFromGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *RemoveUserFromGroupRequest) GetUserName() *string {
	return s.UserName
}

func (s *RemoveUserFromGroupRequest) SetGroupName(v string) *RemoveUserFromGroupRequest {
	s.GroupName = &v
	return s
}

func (s *RemoveUserFromGroupRequest) SetUserName(v string) *RemoveUserFromGroupRequest {
	s.UserName = &v
	return s
}

func (s *RemoveUserFromGroupRequest) Validate() error {
	return dara.Validate(s)
}
