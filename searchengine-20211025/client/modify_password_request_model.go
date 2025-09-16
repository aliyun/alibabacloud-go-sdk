// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPassword(v string) *ModifyPasswordRequest
	GetPassword() *string
	SetUsername(v string) *ModifyPasswordRequest
	GetUsername() *string
}

type ModifyPasswordRequest struct {
	// The password.
	//
	// example:
	//
	// ******************************
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// The username.
	//
	// example:
	//
	// "username"
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ModifyPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPasswordRequest) GoString() string {
	return s.String()
}

func (s *ModifyPasswordRequest) GetPassword() *string {
	return s.Password
}

func (s *ModifyPasswordRequest) GetUsername() *string {
	return s.Username
}

func (s *ModifyPasswordRequest) SetPassword(v string) *ModifyPasswordRequest {
	s.Password = &v
	return s
}

func (s *ModifyPasswordRequest) SetUsername(v string) *ModifyPasswordRequest {
	s.Username = &v
	return s
}

func (s *ModifyPasswordRequest) Validate() error {
	return dara.Validate(s)
}
