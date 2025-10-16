// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddInstanceMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMembers(v []*AddInstanceMembersRequestMembers) *AddInstanceMembersRequest
	GetMembers() []*AddInstanceMembersRequestMembers
}

type AddInstanceMembersRequest struct {
	// The members.
	//
	// This parameter is required.
	Members []*AddInstanceMembersRequestMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
}

func (s AddInstanceMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s AddInstanceMembersRequest) GoString() string {
	return s.String()
}

func (s *AddInstanceMembersRequest) GetMembers() []*AddInstanceMembersRequestMembers {
	return s.Members
}

func (s *AddInstanceMembersRequest) SetMembers(v []*AddInstanceMembersRequestMembers) *AddInstanceMembersRequest {
	s.Members = v
	return s
}

func (s *AddInstanceMembersRequest) Validate() error {
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

type AddInstanceMembersRequestMembers struct {
	// The remarks of the member. The value must be 1 to 256 characters in length.
	//
	// example:
	//
	// renewal
	MemberDesc *string `json:"MemberDesc,omitempty" xml:"MemberDesc,omitempty"`
	// The UID of the member. You can add up to 20 members to Cloud Firewall at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 258039427902****
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
}

func (s AddInstanceMembersRequestMembers) String() string {
	return dara.Prettify(s)
}

func (s AddInstanceMembersRequestMembers) GoString() string {
	return s.String()
}

func (s *AddInstanceMembersRequestMembers) GetMemberDesc() *string {
	return s.MemberDesc
}

func (s *AddInstanceMembersRequestMembers) GetMemberUid() *int64 {
	return s.MemberUid
}

func (s *AddInstanceMembersRequestMembers) SetMemberDesc(v string) *AddInstanceMembersRequestMembers {
	s.MemberDesc = &v
	return s
}

func (s *AddInstanceMembersRequestMembers) SetMemberUid(v int64) *AddInstanceMembersRequestMembers {
	s.MemberUid = &v
	return s
}

func (s *AddInstanceMembersRequestMembers) Validate() error {
	return dara.Validate(s)
}
