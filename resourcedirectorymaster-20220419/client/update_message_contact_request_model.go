// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMessageContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v string) *UpdateMessageContactRequest
	GetContactId() *string
	SetEmailAddress(v string) *UpdateMessageContactRequest
	GetEmailAddress() *string
	SetMessageTypes(v []*string) *UpdateMessageContactRequest
	GetMessageTypes() []*string
	SetName(v string) *UpdateMessageContactRequest
	GetName() *string
	SetPhoneNumber(v string) *UpdateMessageContactRequest
	GetPhoneNumber() *string
	SetTitle(v string) *UpdateMessageContactRequest
	GetTitle() *string
}

type UpdateMessageContactRequest struct {
	// The ID of the contact.
	//
	// example:
	//
	// c-qL4HqKONzOM7****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The email address of the contact.
	//
	// After you specify an email address, you need to call [SendEmailVerificationForMessageContact](~~SendEmailVerificationForMessageContact~~) to send verification information to the email address. After the verification is passed, the email address takes effect.
	//
	// example:
	//
	// someone***@example.com
	EmailAddress *string `json:"EmailAddress,omitempty" xml:"EmailAddress,omitempty"`
	// The types of messages received by the contact.
	MessageTypes []*string `json:"MessageTypes,omitempty" xml:"MessageTypes,omitempty" type:"Repeated"`
	// The name of the contact.
	//
	// example:
	//
	// tom
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The mobile phone number of the contact.
	//
	// Specify the mobile phone number in the `<Country code>-<Mobile phone number>` format.
	//
	// After you specify a mobile phone number, you need to call [SendPhoneVerificationForMessageContact](~~SendPhoneVerificationForMessageContact~~) to send verification information to the mobile phone number. After the verification is passed, the mobile phone number takes effect.
	//
	// example:
	//
	// 86-139****1234
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The job title of the contact.
	//
	// Valid values:
	//
	// 	- FinanceDirector
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- TechnicalDirector
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- MaintenanceDirector
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- CEO
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- ProjectDirector
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Other
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// TechnicalDirector
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateMessageContactRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMessageContactRequest) GoString() string {
	return s.String()
}

func (s *UpdateMessageContactRequest) GetContactId() *string {
	return s.ContactId
}

func (s *UpdateMessageContactRequest) GetEmailAddress() *string {
	return s.EmailAddress
}

func (s *UpdateMessageContactRequest) GetMessageTypes() []*string {
	return s.MessageTypes
}

func (s *UpdateMessageContactRequest) GetName() *string {
	return s.Name
}

func (s *UpdateMessageContactRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *UpdateMessageContactRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateMessageContactRequest) SetContactId(v string) *UpdateMessageContactRequest {
	s.ContactId = &v
	return s
}

func (s *UpdateMessageContactRequest) SetEmailAddress(v string) *UpdateMessageContactRequest {
	s.EmailAddress = &v
	return s
}

func (s *UpdateMessageContactRequest) SetMessageTypes(v []*string) *UpdateMessageContactRequest {
	s.MessageTypes = v
	return s
}

func (s *UpdateMessageContactRequest) SetName(v string) *UpdateMessageContactRequest {
	s.Name = &v
	return s
}

func (s *UpdateMessageContactRequest) SetPhoneNumber(v string) *UpdateMessageContactRequest {
	s.PhoneNumber = &v
	return s
}

func (s *UpdateMessageContactRequest) SetTitle(v string) *UpdateMessageContactRequest {
	s.Title = &v
	return s
}

func (s *UpdateMessageContactRequest) Validate() error {
	return dara.Validate(s)
}
