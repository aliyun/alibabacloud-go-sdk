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
	// The user group attribute of the user that you want to update. Valid values:
	//
	// 	- users: ordinary permissions, which are suitable for ordinary users that need only to submit and debug jobs.
	//
	// 	- wheel: sudo permissions, which are suitable for administrators who need to manage clusters. In addition to submitting and debugging jobs, you can also run sudo commands to install software and restart nodes.
	//
	// example:
	//
	// users
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The password attribute of the user that you want to update. The password must be 6 to 30 characters in length and must contain three of the following four character types:
	//
	// 	- Uppercase letters
	//
	// 	- Lowercase letters
	//
	// 	- Digits
	//
	// 	- Special characters ()~!@#$%^&\\*-_+=|{}[]:;\\"/<>,.?/
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
