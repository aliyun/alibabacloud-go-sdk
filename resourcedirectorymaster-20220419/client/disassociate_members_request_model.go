// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v string) *DisassociateMembersRequest
	GetContactId() *string
	SetMembers(v []*string) *DisassociateMembersRequest
	GetMembers() []*string
}

type DisassociateMembersRequest struct {
	// The ID of the contact.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-qL4HqKONzOM7****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The IDs of objects from which you want to unbind the contact.
	//
	// This parameter is required.
	Members []*string `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
}

func (s DisassociateMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s DisassociateMembersRequest) GoString() string {
	return s.String()
}

func (s *DisassociateMembersRequest) GetContactId() *string {
	return s.ContactId
}

func (s *DisassociateMembersRequest) GetMembers() []*string {
	return s.Members
}

func (s *DisassociateMembersRequest) SetContactId(v string) *DisassociateMembersRequest {
	s.ContactId = &v
	return s
}

func (s *DisassociateMembersRequest) SetMembers(v []*string) *DisassociateMembersRequest {
	s.Members = v
	return s
}

func (s *DisassociateMembersRequest) Validate() error {
	return dara.Validate(s)
}
