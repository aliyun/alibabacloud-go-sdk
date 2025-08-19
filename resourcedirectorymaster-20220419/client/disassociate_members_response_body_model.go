// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMembers(v []*DisassociateMembersResponseBodyMembers) *DisassociateMembersResponseBody
	GetMembers() []*DisassociateMembersResponseBodyMembers
	SetRequestId(v string) *DisassociateMembersResponseBody
	GetRequestId() *string
}

type DisassociateMembersResponseBody struct {
	// The time when the contact was unbound from the object.
	Members []*DisassociateMembersResponseBodyMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 95060F1D-6990-4645-8920-A81D1BBFE992
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisassociateMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisassociateMembersResponseBody) GoString() string {
	return s.String()
}

func (s *DisassociateMembersResponseBody) GetMembers() []*DisassociateMembersResponseBodyMembers {
	return s.Members
}

func (s *DisassociateMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisassociateMembersResponseBody) SetMembers(v []*DisassociateMembersResponseBodyMembers) *DisassociateMembersResponseBody {
	s.Members = v
	return s
}

func (s *DisassociateMembersResponseBody) SetRequestId(v string) *DisassociateMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisassociateMembersResponseBody) Validate() error {
	return dara.Validate(s)
}

type DisassociateMembersResponseBodyMembers struct {
	// The ID of the contact.
	//
	// example:
	//
	// c-qL4HqKONzOM7****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The ID of the object. Valid values:
	//
	// - ID of the resource directory
	//
	// - ID of the folder
	//
	// - ID of the member
	//
	// example:
	//
	// fd-ZDNPiT****
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// The time when the contact was unbound from the object.
	//
	// example:
	//
	// 2023-03-27 17:19:21
	ModifyDate *string `json:"ModifyDate,omitempty" xml:"ModifyDate,omitempty"`
}

func (s DisassociateMembersResponseBodyMembers) String() string {
	return dara.Prettify(s)
}

func (s DisassociateMembersResponseBodyMembers) GoString() string {
	return s.String()
}

func (s *DisassociateMembersResponseBodyMembers) GetContactId() *string {
	return s.ContactId
}

func (s *DisassociateMembersResponseBodyMembers) GetMemberId() *string {
	return s.MemberId
}

func (s *DisassociateMembersResponseBodyMembers) GetModifyDate() *string {
	return s.ModifyDate
}

func (s *DisassociateMembersResponseBodyMembers) SetContactId(v string) *DisassociateMembersResponseBodyMembers {
	s.ContactId = &v
	return s
}

func (s *DisassociateMembersResponseBodyMembers) SetMemberId(v string) *DisassociateMembersResponseBodyMembers {
	s.MemberId = &v
	return s
}

func (s *DisassociateMembersResponseBodyMembers) SetModifyDate(v string) *DisassociateMembersResponseBodyMembers {
	s.ModifyDate = &v
	return s
}

func (s *DisassociateMembersResponseBodyMembers) Validate() error {
	return dara.Validate(s)
}
