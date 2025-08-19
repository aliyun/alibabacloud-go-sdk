// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendEmailVerificationForMessageContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v string) *SendEmailVerificationForMessageContactRequest
	GetContactId() *string
	SetEmailAddress(v string) *SendEmailVerificationForMessageContactRequest
	GetEmailAddress() *string
}

type SendEmailVerificationForMessageContactRequest struct {
	// The ID of the contact.
	//
	// example:
	//
	// c-5gsZAGt***PGduDF
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The email address of the contact.
	//
	// The specified email address must be the one you specify when you call [AddMessageContact](~~AddMessageContact~~).
	//
	// example:
	//
	// someone***@example.com
	EmailAddress *string `json:"EmailAddress,omitempty" xml:"EmailAddress,omitempty"`
}

func (s SendEmailVerificationForMessageContactRequest) String() string {
	return dara.Prettify(s)
}

func (s SendEmailVerificationForMessageContactRequest) GoString() string {
	return s.String()
}

func (s *SendEmailVerificationForMessageContactRequest) GetContactId() *string {
	return s.ContactId
}

func (s *SendEmailVerificationForMessageContactRequest) GetEmailAddress() *string {
	return s.EmailAddress
}

func (s *SendEmailVerificationForMessageContactRequest) SetContactId(v string) *SendEmailVerificationForMessageContactRequest {
	s.ContactId = &v
	return s
}

func (s *SendEmailVerificationForMessageContactRequest) SetEmailAddress(v string) *SendEmailVerificationForMessageContactRequest {
	s.EmailAddress = &v
	return s
}

func (s *SendEmailVerificationForMessageContactRequest) Validate() error {
	return dara.Validate(s)
}
