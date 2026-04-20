// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSuperUser(v bool) *CreateUserRequest
	GetSuperUser() *bool
	SetUserName(v string) *CreateUserRequest
	GetUserName() *string
}

type CreateUserRequest struct {
	// example:
	//
	// false
	SuperUser *bool `json:"superUser,omitempty" xml:"superUser,omitempty"`
	// example:
	//
	// p4_2346134
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s CreateUserRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserRequest) GoString() string {
	return s.String()
}

func (s *CreateUserRequest) GetSuperUser() *bool {
	return s.SuperUser
}

func (s *CreateUserRequest) GetUserName() *string {
	return s.UserName
}

func (s *CreateUserRequest) SetSuperUser(v bool) *CreateUserRequest {
	s.SuperUser = &v
	return s
}

func (s *CreateUserRequest) SetUserName(v string) *CreateUserRequest {
	s.UserName = &v
	return s
}

func (s *CreateUserRequest) Validate() error {
	return dara.Validate(s)
}
