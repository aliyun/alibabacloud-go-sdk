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
	// example:
	//
	// false
	SuperUser *string `json:"superUser,omitempty" xml:"superUser,omitempty"`
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
