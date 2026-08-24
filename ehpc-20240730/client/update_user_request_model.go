// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UpdateUserRequest
	GetClusterId() *string
	SetGroup(v string) *UpdateUserRequest
	GetGroup() *string
	SetPassword(v string) *UpdateUserRequest
	GetPassword() *string
	SetUserName(v string) *UpdateUserRequest
	GetUserName() *string
}

type UpdateUserRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The user group property to update. Valid values:
	//
	// - users: ordinary permission group. This group is suitable for regular users who only need to submit and debug jobs.
	//
	// - wheel: sudo permission group. This group is suitable for administrators who need cluster management. In addition to submitting and debugging jobs, users in this group can execute sudo commands to install software, restart nodes, and perform other operations.
	//
	// example:
	//
	// users
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The user password property to update. The password must be 8 to 30 characters in length and must contain at least three of the following four character types:
	//
	// - Uppercase letters
	//
	// - Lowercase letters
	//
	// - Digits
	//
	// - Special characters: ()~!@#$%^&*-_+=|{}[]:;\\"/<>,.?/
	//
	// example:
	//
	// 123****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The username.
	//
	// This parameter is required.
	//
	// example:
	//
	// testuser
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s UpdateUserRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateUserRequest) GetGroup() *string {
	return s.Group
}

func (s *UpdateUserRequest) GetPassword() *string {
	return s.Password
}

func (s *UpdateUserRequest) GetUserName() *string {
	return s.UserName
}

func (s *UpdateUserRequest) SetClusterId(v string) *UpdateUserRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateUserRequest) SetGroup(v string) *UpdateUserRequest {
	s.Group = &v
	return s
}

func (s *UpdateUserRequest) SetPassword(v string) *UpdateUserRequest {
	s.Password = &v
	return s
}

func (s *UpdateUserRequest) SetUserName(v string) *UpdateUserRequest {
	s.UserName = &v
	return s
}

func (s *UpdateUserRequest) Validate() error {
	return dara.Validate(s)
}
