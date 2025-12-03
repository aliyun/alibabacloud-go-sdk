// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdByPhoneNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPhoneNumber(v string) *GetUserIdByPhoneNumberRequest
	GetPhoneNumber() *string
}

type GetUserIdByPhoneNumberRequest struct {
	// The mobile number of the user who owns the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 156xxxxxxx
	PhoneNumber *string `json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
}

func (s GetUserIdByPhoneNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdByPhoneNumberRequest) GoString() string {
	return s.String()
}

func (s *GetUserIdByPhoneNumberRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetUserIdByPhoneNumberRequest) SetPhoneNumber(v string) *GetUserIdByPhoneNumberRequest {
	s.PhoneNumber = &v
	return s
}

func (s *GetUserIdByPhoneNumberRequest) Validate() error {
	return dara.Validate(s)
}
