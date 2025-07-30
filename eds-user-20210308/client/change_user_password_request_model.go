// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeUserPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndUserId(v string) *ChangeUserPasswordRequest
	GetEndUserId() *string
	SetNewPassword(v string) *ChangeUserPasswordRequest
	GetNewPassword() *string
}

type ChangeUserPasswordRequest struct {
	// example:
	//
	// alice***
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// example:
	//
	// Admin@12***
	NewPassword *string `json:"NewPassword,omitempty" xml:"NewPassword,omitempty"`
}

func (s ChangeUserPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeUserPasswordRequest) GoString() string {
	return s.String()
}

func (s *ChangeUserPasswordRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *ChangeUserPasswordRequest) GetNewPassword() *string {
	return s.NewPassword
}

func (s *ChangeUserPasswordRequest) SetEndUserId(v string) *ChangeUserPasswordRequest {
	s.EndUserId = &v
	return s
}

func (s *ChangeUserPasswordRequest) SetNewPassword(v string) *ChangeUserPasswordRequest {
	s.NewPassword = &v
	return s
}

func (s *ChangeUserPasswordRequest) Validate() error {
	return dara.Validate(s)
}
