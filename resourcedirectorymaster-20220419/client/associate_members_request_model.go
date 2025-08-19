// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v string) *AssociateMembersRequest
	GetContactId() *string
	SetMembers(v []*string) *AssociateMembersRequest
	GetMembers() []*string
}

type AssociateMembersRequest struct {
	// The ID of the contact.
	//
	// example:
	//
	// c-qL4HqKONzOM7****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The IDs of objects to which you want to bind the contact.
	Members []*string `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
}

func (s AssociateMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateMembersRequest) GoString() string {
	return s.String()
}

func (s *AssociateMembersRequest) GetContactId() *string {
	return s.ContactId
}

func (s *AssociateMembersRequest) GetMembers() []*string {
	return s.Members
}

func (s *AssociateMembersRequest) SetContactId(v string) *AssociateMembersRequest {
	s.ContactId = &v
	return s
}

func (s *AssociateMembersRequest) SetMembers(v []*string) *AssociateMembersRequest {
	s.Members = v
	return s
}

func (s *AssociateMembersRequest) Validate() error {
	return dara.Validate(s)
}
