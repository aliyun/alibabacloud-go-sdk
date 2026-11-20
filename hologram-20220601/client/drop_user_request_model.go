// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDropUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSuperUser(v string) *DropUserRequest
	GetSuperUser() *string
	SetUserName(v string) *DropUserRequest
	GetUserName() *string
}

type DropUserRequest struct {
	// Specifies whether the user is a superuser.
	//
	// example:
	//
	// false
	SuperUser *string `json:"superUser,omitempty" xml:"superUser,omitempty"`
	// The username. Only RAM users and STS accounts are supported. If the user is a RAM user, the name must be prefixed with "p4_", such as p4_234253434. If the user is an STS account, the name must be prefixed with "v4_", such as v4_3234134.
	//
	// example:
	//
	// p4_234253434
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s DropUserRequest) String() string {
	return dara.Prettify(s)
}

func (s DropUserRequest) GoString() string {
	return s.String()
}

func (s *DropUserRequest) GetSuperUser() *string {
	return s.SuperUser
}

func (s *DropUserRequest) GetUserName() *string {
	return s.UserName
}

func (s *DropUserRequest) SetSuperUser(v string) *DropUserRequest {
	s.SuperUser = &v
	return s
}

func (s *DropUserRequest) SetUserName(v string) *DropUserRequest {
	s.UserName = &v
	return s
}

func (s *DropUserRequest) Validate() error {
	return dara.Validate(s)
}
