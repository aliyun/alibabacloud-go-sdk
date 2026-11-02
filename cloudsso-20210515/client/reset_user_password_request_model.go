// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetUserPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *ResetUserPasswordRequest
	GetDirectoryId() *string
	SetGenerateRandomPassword(v bool) *ResetUserPasswordRequest
	GetGenerateRandomPassword() *bool
	SetPassword(v string) *ResetUserPasswordRequest
	GetPassword() *string
	SetRequirePasswordResetForNextLogin(v bool) *ResetUserPasswordRequest
	GetRequirePasswordResetForNextLogin() *bool
	SetUserId(v string) *ResetUserPasswordRequest
	GetUserId() *string
}

type ResetUserPasswordRequest struct {
	// The directory ID.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// Specifies whether to automatically generate a password. Valid values:
	//
	// - True: The password is automatically generated.
	//
	// - False (default): The password is manually specified.
	//
	// example:
	//
	// True
	GenerateRandomPassword *bool `json:"GenerateRandomPassword,omitempty" xml:"GenerateRandomPassword,omitempty"`
	// The new password.
	//
	// Format: The password must contain at least one uppercase letter, one lowercase letter, one digit, and one special character.
	//
	// Length: 8 to 32 characters.
	//
	// > If GenerateRandomPassword is set to False, you must specify Password.
	//
	// example:
	//
	// uc)XK$?ictf72CKFDy9vtWaFmISl****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Specifies whether the user must reset the password upon the next logon. Valid values:
	//
	// - True: The user must reset the password.
	//
	// - False (default): The user does not need to reset the password.
	//
	// example:
	//
	// False
	RequirePasswordResetForNextLogin *bool `json:"RequirePasswordResetForNextLogin,omitempty" xml:"RequirePasswordResetForNextLogin,omitempty"`
	// The user ID.
	//
	// example:
	//
	// u-00q8wbq42wiltcrk****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ResetUserPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetUserPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetUserPasswordRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ResetUserPasswordRequest) GetGenerateRandomPassword() *bool {
	return s.GenerateRandomPassword
}

func (s *ResetUserPasswordRequest) GetPassword() *string {
	return s.Password
}

func (s *ResetUserPasswordRequest) GetRequirePasswordResetForNextLogin() *bool {
	return s.RequirePasswordResetForNextLogin
}

func (s *ResetUserPasswordRequest) GetUserId() *string {
	return s.UserId
}

func (s *ResetUserPasswordRequest) SetDirectoryId(v string) *ResetUserPasswordRequest {
	s.DirectoryId = &v
	return s
}

func (s *ResetUserPasswordRequest) SetGenerateRandomPassword(v bool) *ResetUserPasswordRequest {
	s.GenerateRandomPassword = &v
	return s
}

func (s *ResetUserPasswordRequest) SetPassword(v string) *ResetUserPasswordRequest {
	s.Password = &v
	return s
}

func (s *ResetUserPasswordRequest) SetRequirePasswordResetForNextLogin(v bool) *ResetUserPasswordRequest {
	s.RequirePasswordResetForNextLogin = &v
	return s
}

func (s *ResetUserPasswordRequest) SetUserId(v string) *ResetUserPasswordRequest {
	s.UserId = &v
	return s
}

func (s *ResetUserPasswordRequest) Validate() error {
	return dara.Validate(s)
}
