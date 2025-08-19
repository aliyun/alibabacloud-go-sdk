// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelMessageContactUpdateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v string) *CancelMessageContactUpdateRequest
	GetContactId() *string
	SetEmailAddress(v string) *CancelMessageContactUpdateRequest
	GetEmailAddress() *string
	SetPhoneNumber(v string) *CancelMessageContactUpdateRequest
	GetPhoneNumber() *string
}

type CancelMessageContactUpdateRequest struct {
	// The ID of the contact.
	//
	// example:
	//
	// c-qL4HqKONzOM7****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The email address of the contact.
	//
	// example:
	//
	// someone***@example.com
	EmailAddress *string `json:"EmailAddress,omitempty" xml:"EmailAddress,omitempty"`
	// The mobile phone number of the contact.
	//
	// Specify the mobile phone number in the `<Country code>-<Mobile phone number>` format.
	//
	// example:
	//
	// 86-139****1234
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s CancelMessageContactUpdateRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelMessageContactUpdateRequest) GoString() string {
	return s.String()
}

func (s *CancelMessageContactUpdateRequest) GetContactId() *string {
	return s.ContactId
}

func (s *CancelMessageContactUpdateRequest) GetEmailAddress() *string {
	return s.EmailAddress
}

func (s *CancelMessageContactUpdateRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *CancelMessageContactUpdateRequest) SetContactId(v string) *CancelMessageContactUpdateRequest {
	s.ContactId = &v
	return s
}

func (s *CancelMessageContactUpdateRequest) SetEmailAddress(v string) *CancelMessageContactUpdateRequest {
	s.EmailAddress = &v
	return s
}

func (s *CancelMessageContactUpdateRequest) SetPhoneNumber(v string) *CancelMessageContactUpdateRequest {
	s.PhoneNumber = &v
	return s
}

func (s *CancelMessageContactUpdateRequest) Validate() error {
	return dara.Validate(s)
}
