// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendPhoneVerificationForMessageContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v string) *SendPhoneVerificationForMessageContactRequest
	GetContactId() *string
	SetPhoneNumber(v string) *SendPhoneVerificationForMessageContactRequest
	GetPhoneNumber() *string
}

type SendPhoneVerificationForMessageContactRequest struct {
	// The ID of the contact.
	//
	// example:
	//
	// c-qL4HqKONzOM7****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The mobile phone number of the contact.
	//
	// The value must be in the `<Country code>-<Mobile phone number>` format.
	//
	// The specified mobile phone number must be the one you specify when you call the [AddMessageContact](~~AddMessageContact~~) operation.
	//
	// example:
	//
	// 86-139****1234
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s SendPhoneVerificationForMessageContactRequest) String() string {
	return dara.Prettify(s)
}

func (s SendPhoneVerificationForMessageContactRequest) GoString() string {
	return s.String()
}

func (s *SendPhoneVerificationForMessageContactRequest) GetContactId() *string {
	return s.ContactId
}

func (s *SendPhoneVerificationForMessageContactRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *SendPhoneVerificationForMessageContactRequest) SetContactId(v string) *SendPhoneVerificationForMessageContactRequest {
	s.ContactId = &v
	return s
}

func (s *SendPhoneVerificationForMessageContactRequest) SetPhoneNumber(v string) *SendPhoneVerificationForMessageContactRequest {
	s.PhoneNumber = &v
	return s
}

func (s *SendPhoneVerificationForMessageContactRequest) Validate() error {
	return dara.Validate(s)
}
