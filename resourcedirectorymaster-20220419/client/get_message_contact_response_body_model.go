// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessageContactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContact(v *GetMessageContactResponseBodyContact) *GetMessageContactResponseBody
	GetContact() *GetMessageContactResponseBodyContact
	SetRequestId(v string) *GetMessageContactResponseBody
	GetRequestId() *string
}

type GetMessageContactResponseBody struct {
	// The information about the contact.
	Contact *GetMessageContactResponseBodyContact `json:"Contact,omitempty" xml:"Contact,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMessageContactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMessageContactResponseBody) GoString() string {
	return s.String()
}

func (s *GetMessageContactResponseBody) GetContact() *GetMessageContactResponseBodyContact {
	return s.Contact
}

func (s *GetMessageContactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMessageContactResponseBody) SetContact(v *GetMessageContactResponseBodyContact) *GetMessageContactResponseBody {
	s.Contact = v
	return s
}

func (s *GetMessageContactResponseBody) SetRequestId(v string) *GetMessageContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMessageContactResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMessageContactResponseBodyContact struct {
	// The ID of the contact.
	//
	// example:
	//
	// c-qL4HqKONzOM7****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The time when the contact was created.
	//
	// example:
	//
	// 2023-03-27 17:19:21
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The email address of the contact.
	//
	// example:
	//
	// someone***@example.com
	EmailAddress *string `json:"EmailAddress,omitempty" xml:"EmailAddress,omitempty"`
	// The IDs of objects to which the contact is bound.
	Members []*string `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
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
	// example:
	//
	// 86-139****1234
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The status of the contact. Valid values:
	//
	// 	- Verifying
	//
	// 	- Active
	//
	// 	- Deleting
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The job title of the contact.
	//
	// example:
	//
	// TechnicalDirector
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetMessageContactResponseBodyContact) String() string {
	return dara.Prettify(s)
}

func (s GetMessageContactResponseBodyContact) GoString() string {
	return s.String()
}

func (s *GetMessageContactResponseBodyContact) GetContactId() *string {
	return s.ContactId
}

func (s *GetMessageContactResponseBodyContact) GetCreateDate() *string {
	return s.CreateDate
}

func (s *GetMessageContactResponseBodyContact) GetEmailAddress() *string {
	return s.EmailAddress
}

func (s *GetMessageContactResponseBodyContact) GetMembers() []*string {
	return s.Members
}

func (s *GetMessageContactResponseBodyContact) GetMessageTypes() []*string {
	return s.MessageTypes
}

func (s *GetMessageContactResponseBodyContact) GetName() *string {
	return s.Name
}

func (s *GetMessageContactResponseBodyContact) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetMessageContactResponseBodyContact) GetStatus() *string {
	return s.Status
}

func (s *GetMessageContactResponseBodyContact) GetTitle() *string {
	return s.Title
}

func (s *GetMessageContactResponseBodyContact) SetContactId(v string) *GetMessageContactResponseBodyContact {
	s.ContactId = &v
	return s
}

func (s *GetMessageContactResponseBodyContact) SetCreateDate(v string) *GetMessageContactResponseBodyContact {
	s.CreateDate = &v
	return s
}

func (s *GetMessageContactResponseBodyContact) SetEmailAddress(v string) *GetMessageContactResponseBodyContact {
	s.EmailAddress = &v
	return s
}

func (s *GetMessageContactResponseBodyContact) SetMembers(v []*string) *GetMessageContactResponseBodyContact {
	s.Members = v
	return s
}

func (s *GetMessageContactResponseBodyContact) SetMessageTypes(v []*string) *GetMessageContactResponseBodyContact {
	s.MessageTypes = v
	return s
}

func (s *GetMessageContactResponseBodyContact) SetName(v string) *GetMessageContactResponseBodyContact {
	s.Name = &v
	return s
}

func (s *GetMessageContactResponseBodyContact) SetPhoneNumber(v string) *GetMessageContactResponseBodyContact {
	s.PhoneNumber = &v
	return s
}

func (s *GetMessageContactResponseBodyContact) SetStatus(v string) *GetMessageContactResponseBodyContact {
	s.Status = &v
	return s
}

func (s *GetMessageContactResponseBodyContact) SetTitle(v string) *GetMessageContactResponseBodyContact {
	s.Title = &v
	return s
}

func (s *GetMessageContactResponseBodyContact) Validate() error {
	return dara.Validate(s)
}
