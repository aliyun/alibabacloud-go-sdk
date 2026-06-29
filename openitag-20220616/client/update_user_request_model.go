// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRole(v string) *UpdateUserRequest
	GetRole() *string
	SetUserName(v string) *UpdateUserRequest
	GetUserName() *string
}

type UpdateUserRequest struct {
	// Role. Possible values:
	//
	// - OPERATOR: Annotator.
	//
	// - ADMIN: Administrator.
	//
	// - LEADER: Annotation team leader.
	//
	// This parameter is required.
	//
	// example:
	//
	// ADMIN
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// Username.
	//
	// This parameter is required.
	//
	// example:
	//
	// user1
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s UpdateUserRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserRequest) GetRole() *string {
	return s.Role
}

func (s *UpdateUserRequest) GetUserName() *string {
	return s.UserName
}

func (s *UpdateUserRequest) SetRole(v string) *UpdateUserRequest {
	s.Role = &v
	return s
}

func (s *UpdateUserRequest) SetUserName(v string) *UpdateUserRequest {
	s.UserName = &v
	return s
}

func (s *UpdateUserRequest) Validate() error {
	return dara.Validate(s)
}
