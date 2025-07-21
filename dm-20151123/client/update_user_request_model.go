// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUser(v *UpdateUserRequestUser) *UpdateUserRequest
	GetUser() *UpdateUserRequestUser
}

type UpdateUserRequest struct {
	// User Information
	User *UpdateUserRequestUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
}

func (s UpdateUserRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserRequest) GetUser() *UpdateUserRequestUser {
	return s.User
}

func (s *UpdateUserRequest) SetUser(v *UpdateUserRequestUser) *UpdateUserRequest {
	s.User = v
	return s
}

func (s *UpdateUserRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateUserRequestUser struct {
	// Whether EventBridge is enabled
	//
	// example:
	//
	// true
	EnableEventbridge *bool `json:"EnableEventbridge,omitempty" xml:"EnableEventbridge,omitempty"`
}

func (s UpdateUserRequestUser) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserRequestUser) GoString() string {
	return s.String()
}

func (s *UpdateUserRequestUser) GetEnableEventbridge() *bool {
	return s.EnableEventbridge
}

func (s *UpdateUserRequestUser) SetEnableEventbridge(v bool) *UpdateUserRequestUser {
	s.EnableEventbridge = &v
	return s
}

func (s *UpdateUserRequestUser) Validate() error {
	return dara.Validate(s)
}
