// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMessageContactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContact(v *AddMessageContactResponseBodyContact) *AddMessageContactResponseBody
	GetContact() *AddMessageContactResponseBodyContact
	SetRequestId(v string) *AddMessageContactResponseBody
	GetRequestId() *string
}

type AddMessageContactResponseBody struct {
	// The information about the contact.
	Contact *AddMessageContactResponseBodyContact `json:"Contact,omitempty" xml:"Contact,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 2DFCE4C9-04A9-4C83-BB14-FE791275EC53
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddMessageContactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddMessageContactResponseBody) GoString() string {
	return s.String()
}

func (s *AddMessageContactResponseBody) GetContact() *AddMessageContactResponseBodyContact {
	return s.Contact
}

func (s *AddMessageContactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddMessageContactResponseBody) SetContact(v *AddMessageContactResponseBodyContact) *AddMessageContactResponseBody {
	s.Contact = v
	return s
}

func (s *AddMessageContactResponseBody) SetRequestId(v string) *AddMessageContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddMessageContactResponseBody) Validate() error {
	if s.Contact != nil {
		if err := s.Contact.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddMessageContactResponseBodyContact struct {
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
}

func (s AddMessageContactResponseBodyContact) String() string {
	return dara.Prettify(s)
}

func (s AddMessageContactResponseBodyContact) GoString() string {
	return s.String()
}

func (s *AddMessageContactResponseBodyContact) GetContactId() *string {
	return s.ContactId
}

func (s *AddMessageContactResponseBodyContact) GetCreateDate() *string {
	return s.CreateDate
}

func (s *AddMessageContactResponseBodyContact) SetContactId(v string) *AddMessageContactResponseBodyContact {
	s.ContactId = &v
	return s
}

func (s *AddMessageContactResponseBodyContact) SetCreateDate(v string) *AddMessageContactResponseBodyContact {
	s.CreateDate = &v
	return s
}

func (s *AddMessageContactResponseBodyContact) Validate() error {
	return dara.Validate(s)
}
