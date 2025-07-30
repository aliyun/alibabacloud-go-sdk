// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEmail(v string) *ModifyUserRequest
	GetEmail() *string
	SetEndUserId(v string) *ModifyUserRequest
	GetEndUserId() *string
	SetPhone(v string) *ModifyUserRequest
	GetPhone() *string
}

type ModifyUserRequest struct {
	// The email address of the convenience user. For a user-activated convenience user, the email address or mobile number must be verified. You can choose to verify the email address or the mobile number. For an administrator-activated convenience user, the email address and mobile number can be left empty.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The name of the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// Alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The mobile number of the convenience user. For a user-activated convenience user, the email address or mobile number must be verified. You can choose to verify the email address or the mobile number. For an administrator-activated convenience user, the email address and mobile number can be left empty.
	//
	// >  Accounts created on the International site (alibabacloud.com) do not support mobile number-based authentication.
	//
	// example:
	//
	// 1381111****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s ModifyUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserRequest) GetEmail() *string {
	return s.Email
}

func (s *ModifyUserRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *ModifyUserRequest) GetPhone() *string {
	return s.Phone
}

func (s *ModifyUserRequest) SetEmail(v string) *ModifyUserRequest {
	s.Email = &v
	return s
}

func (s *ModifyUserRequest) SetEndUserId(v string) *ModifyUserRequest {
	s.EndUserId = &v
	return s
}

func (s *ModifyUserRequest) SetPhone(v string) *ModifyUserRequest {
	s.Phone = &v
	return s
}

func (s *ModifyUserRequest) Validate() error {
	return dara.Validate(s)
}
