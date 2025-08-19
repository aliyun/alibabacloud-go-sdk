// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMessageContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEmailAddress(v string) *AddMessageContactRequest
	GetEmailAddress() *string
	SetMessageTypes(v []*string) *AddMessageContactRequest
	GetMessageTypes() []*string
	SetName(v string) *AddMessageContactRequest
	GetName() *string
	SetPhoneNumber(v string) *AddMessageContactRequest
	GetPhoneNumber() *string
	SetTitle(v string) *AddMessageContactRequest
	GetTitle() *string
}

type AddMessageContactRequest struct {
	// The email address of the contact.
	//
	// After you specify an email address, you need to call [SendEmailVerificationForMessageContact](~~SendEmailVerificationForMessageContact~~) to send verification information to the email address. After the verification is passed, the email address takes effect.
	//
	// This parameter is required.
	//
	// example:
	//
	// someone***@example.com
	EmailAddress *string `json:"EmailAddress,omitempty" xml:"EmailAddress,omitempty"`
	// The types of messages received by the contact.
	//
	// This parameter is required.
	MessageTypes []*string `json:"MessageTypes,omitempty" xml:"MessageTypes,omitempty" type:"Repeated"`
	// The name of the contact.
	//
	// The name must be unique in your resource directory.
	//
	// The name must be 2 to 12 characters in length and can contain only letters.
	//
	// This parameter is required.
	//
	// example:
	//
	// tom
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The mobile phone number of the contact.
	//
	// Specify the mobile phone number in the `<Country code>-<Mobile phone number>` format.
	//
	// > Only mobile phone numbers in the `86-<Mobile phone number>` format in the Chinese mainland are supported.
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
	// This parameter is required.
	//
	// example:
	//
	// TechnicalDirector
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s AddMessageContactRequest) String() string {
	return dara.Prettify(s)
}

func (s AddMessageContactRequest) GoString() string {
	return s.String()
}

func (s *AddMessageContactRequest) GetEmailAddress() *string {
	return s.EmailAddress
}

func (s *AddMessageContactRequest) GetMessageTypes() []*string {
	return s.MessageTypes
}

func (s *AddMessageContactRequest) GetName() *string {
	return s.Name
}

func (s *AddMessageContactRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *AddMessageContactRequest) GetTitle() *string {
	return s.Title
}

func (s *AddMessageContactRequest) SetEmailAddress(v string) *AddMessageContactRequest {
	s.EmailAddress = &v
	return s
}

func (s *AddMessageContactRequest) SetMessageTypes(v []*string) *AddMessageContactRequest {
	s.MessageTypes = v
	return s
}

func (s *AddMessageContactRequest) SetName(v string) *AddMessageContactRequest {
	s.Name = &v
	return s
}

func (s *AddMessageContactRequest) SetPhoneNumber(v string) *AddMessageContactRequest {
	s.PhoneNumber = &v
	return s
}

func (s *AddMessageContactRequest) SetTitle(v string) *AddMessageContactRequest {
	s.Title = &v
	return s
}

func (s *AddMessageContactRequest) Validate() error {
	return dara.Validate(s)
}
