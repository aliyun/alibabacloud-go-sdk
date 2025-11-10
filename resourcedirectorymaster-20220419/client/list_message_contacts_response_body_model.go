// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessageContactsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContacts(v []*ListMessageContactsResponseBodyContacts) *ListMessageContactsResponseBody
	GetContacts() []*ListMessageContactsResponseBodyContacts
	SetPageNumber(v int32) *ListMessageContactsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMessageContactsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListMessageContactsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListMessageContactsResponseBody
	GetTotalCount() *int32
}

type ListMessageContactsResponseBody struct {
	// The time when the contact was bound to the objects.
	Contacts []*ListMessageContactsResponseBodyContacts `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 95060F1D-6990-4645-8920-A81D1BBFE992
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMessageContactsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMessageContactsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMessageContactsResponseBody) GetContacts() []*ListMessageContactsResponseBodyContacts {
	return s.Contacts
}

func (s *ListMessageContactsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMessageContactsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMessageContactsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMessageContactsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListMessageContactsResponseBody) SetContacts(v []*ListMessageContactsResponseBodyContacts) *ListMessageContactsResponseBody {
	s.Contacts = v
	return s
}

func (s *ListMessageContactsResponseBody) SetPageNumber(v int32) *ListMessageContactsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListMessageContactsResponseBody) SetPageSize(v int32) *ListMessageContactsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListMessageContactsResponseBody) SetRequestId(v string) *ListMessageContactsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMessageContactsResponseBody) SetTotalCount(v int32) *ListMessageContactsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMessageContactsResponseBody) Validate() error {
	if s.Contacts != nil {
		for _, item := range s.Contacts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMessageContactsResponseBodyContacts struct {
	// The time when the contact was bound to the objects.
	//
	// example:
	//
	// 2023-03-27 17:19:21
	AssociatedDate *string `json:"AssociatedDate,omitempty" xml:"AssociatedDate,omitempty"`
	// The ID of the contact.
	//
	// example:
	//
	// c-qL4HqKONzOM7****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The time when the contact was added.
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
	// - Verifying
	//
	// - Active
	//
	// - Deleting
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

func (s ListMessageContactsResponseBodyContacts) String() string {
	return dara.Prettify(s)
}

func (s ListMessageContactsResponseBodyContacts) GoString() string {
	return s.String()
}

func (s *ListMessageContactsResponseBodyContacts) GetAssociatedDate() *string {
	return s.AssociatedDate
}

func (s *ListMessageContactsResponseBodyContacts) GetContactId() *string {
	return s.ContactId
}

func (s *ListMessageContactsResponseBodyContacts) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListMessageContactsResponseBodyContacts) GetEmailAddress() *string {
	return s.EmailAddress
}

func (s *ListMessageContactsResponseBodyContacts) GetMembers() []*string {
	return s.Members
}

func (s *ListMessageContactsResponseBodyContacts) GetMessageTypes() []*string {
	return s.MessageTypes
}

func (s *ListMessageContactsResponseBodyContacts) GetName() *string {
	return s.Name
}

func (s *ListMessageContactsResponseBodyContacts) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *ListMessageContactsResponseBodyContacts) GetStatus() *string {
	return s.Status
}

func (s *ListMessageContactsResponseBodyContacts) GetTitle() *string {
	return s.Title
}

func (s *ListMessageContactsResponseBodyContacts) SetAssociatedDate(v string) *ListMessageContactsResponseBodyContacts {
	s.AssociatedDate = &v
	return s
}

func (s *ListMessageContactsResponseBodyContacts) SetContactId(v string) *ListMessageContactsResponseBodyContacts {
	s.ContactId = &v
	return s
}

func (s *ListMessageContactsResponseBodyContacts) SetCreateDate(v string) *ListMessageContactsResponseBodyContacts {
	s.CreateDate = &v
	return s
}

func (s *ListMessageContactsResponseBodyContacts) SetEmailAddress(v string) *ListMessageContactsResponseBodyContacts {
	s.EmailAddress = &v
	return s
}

func (s *ListMessageContactsResponseBodyContacts) SetMembers(v []*string) *ListMessageContactsResponseBodyContacts {
	s.Members = v
	return s
}

func (s *ListMessageContactsResponseBodyContacts) SetMessageTypes(v []*string) *ListMessageContactsResponseBodyContacts {
	s.MessageTypes = v
	return s
}

func (s *ListMessageContactsResponseBodyContacts) SetName(v string) *ListMessageContactsResponseBodyContacts {
	s.Name = &v
	return s
}

func (s *ListMessageContactsResponseBodyContacts) SetPhoneNumber(v string) *ListMessageContactsResponseBodyContacts {
	s.PhoneNumber = &v
	return s
}

func (s *ListMessageContactsResponseBodyContacts) SetStatus(v string) *ListMessageContactsResponseBodyContacts {
	s.Status = &v
	return s
}

func (s *ListMessageContactsResponseBodyContacts) SetTitle(v string) *ListMessageContactsResponseBodyContacts {
	s.Title = &v
	return s
}

func (s *ListMessageContactsResponseBodyContacts) Validate() error {
	return dara.Validate(s)
}
