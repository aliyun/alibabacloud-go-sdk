// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMembers(v []*AssociateMembersResponseBodyMembers) *AssociateMembersResponseBody
	GetMembers() []*AssociateMembersResponseBodyMembers
	SetRequestId(v string) *AssociateMembersResponseBody
	GetRequestId() *string
}

type AssociateMembersResponseBody struct {
	// The time when the contact was bound to the object.
	Members []*AssociateMembersResponseBodyMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 95060F1D-6990-4645-8920-A81D1BBFE992
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateMembersResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateMembersResponseBody) GetMembers() []*AssociateMembersResponseBodyMembers {
	return s.Members
}

func (s *AssociateMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateMembersResponseBody) SetMembers(v []*AssociateMembersResponseBodyMembers) *AssociateMembersResponseBody {
	s.Members = v
	return s
}

func (s *AssociateMembersResponseBody) SetRequestId(v string) *AssociateMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateMembersResponseBody) Validate() error {
	if s.Members != nil {
		for _, item := range s.Members {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AssociateMembersResponseBodyMembers struct {
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
	// The time when the contact was bound to the object.
	//
	// example:
	//
	// 2023-03-27 17:19:21
	ModifyDate *string `json:"ModifyDate,omitempty" xml:"ModifyDate,omitempty"`
}

func (s AssociateMembersResponseBodyMembers) String() string {
	return dara.Prettify(s)
}

func (s AssociateMembersResponseBodyMembers) GoString() string {
	return s.String()
}

func (s *AssociateMembersResponseBodyMembers) GetContactId() *string {
	return s.ContactId
}

func (s *AssociateMembersResponseBodyMembers) GetMemberId() *string {
	return s.MemberId
}

func (s *AssociateMembersResponseBodyMembers) GetModifyDate() *string {
	return s.ModifyDate
}

func (s *AssociateMembersResponseBodyMembers) SetContactId(v string) *AssociateMembersResponseBodyMembers {
	s.ContactId = &v
	return s
}

func (s *AssociateMembersResponseBodyMembers) SetMemberId(v string) *AssociateMembersResponseBodyMembers {
	s.MemberId = &v
	return s
}

func (s *AssociateMembersResponseBodyMembers) SetModifyDate(v string) *AssociateMembersResponseBodyMembers {
	s.ModifyDate = &v
	return s
}

func (s *AssociateMembersResponseBodyMembers) Validate() error {
	return dara.Validate(s)
}
