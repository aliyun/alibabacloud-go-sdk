// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhoneNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPhoneNumber(v string) *GetPhoneNumberResponseBody
	GetPhoneNumber() *string
}

type GetPhoneNumberResponseBody struct {
	// example:
	//
	// 137****
	PhoneNumber *string `json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
}

func (s GetPhoneNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPhoneNumberResponseBody) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberResponseBody) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetPhoneNumberResponseBody) SetPhoneNumber(v string) *GetPhoneNumberResponseBody {
	s.PhoneNumber = &v
	return s
}

func (s *GetPhoneNumberResponseBody) Validate() error {
	return dara.Validate(s)
}
