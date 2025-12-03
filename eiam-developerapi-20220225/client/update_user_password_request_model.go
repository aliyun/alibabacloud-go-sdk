// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPassword(v string) *UpdateUserPasswordRequest
	GetPassword() *string
}

type UpdateUserPasswordRequest struct {
	// example:
	//
	// xxxx
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
}

func (s UpdateUserPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserPasswordRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserPasswordRequest) GetPassword() *string {
	return s.Password
}

func (s *UpdateUserPasswordRequest) SetPassword(v string) *UpdateUserPasswordRequest {
	s.Password = &v
	return s
}

func (s *UpdateUserPasswordRequest) Validate() error {
	return dara.Validate(s)
}
