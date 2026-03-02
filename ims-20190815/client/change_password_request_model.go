// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangePasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNewPassword(v string) *ChangePasswordRequest
	GetNewPassword() *string
	SetOldPassword(v string) *ChangePasswordRequest
	GetOldPassword() *string
}

type ChangePasswordRequest struct {
	// This parameter is required.
	NewPassword *string `json:"NewPassword,omitempty" xml:"NewPassword,omitempty"`
	// This parameter is required.
	OldPassword *string `json:"OldPassword,omitempty" xml:"OldPassword,omitempty"`
}

func (s ChangePasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangePasswordRequest) GoString() string {
	return s.String()
}

func (s *ChangePasswordRequest) GetNewPassword() *string {
	return s.NewPassword
}

func (s *ChangePasswordRequest) GetOldPassword() *string {
	return s.OldPassword
}

func (s *ChangePasswordRequest) SetNewPassword(v string) *ChangePasswordRequest {
	s.NewPassword = &v
	return s
}

func (s *ChangePasswordRequest) SetOldPassword(v string) *ChangePasswordRequest {
	s.OldPassword = &v
	return s
}

func (s *ChangePasswordRequest) Validate() error {
	return dara.Validate(s)
}
