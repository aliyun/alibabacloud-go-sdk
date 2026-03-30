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
	// The new password that is used to log on to the console.
	//
	// The password must meet the complexity requirements. For more information, see [GetPasswordPolicy](https://help.aliyun.com/document_detail/186691.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// newpassword
	NewPassword *string `json:"NewPassword,omitempty" xml:"NewPassword,omitempty"`
	// The old password that is used to log on to the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// mypassword
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
