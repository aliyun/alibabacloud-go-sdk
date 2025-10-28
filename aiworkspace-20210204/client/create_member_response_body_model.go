// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMembers(v []*CreateMemberResponseBodyMembers) *CreateMemberResponseBody
	GetMembers() []*CreateMemberResponseBodyMembers
	SetRequestId(v string) *CreateMemberResponseBody
	GetRequestId() *string
}

type CreateMemberResponseBody struct {
	// The returned members.
	Members []*CreateMemberResponseBodyMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// DA869D1B-035A-43B2-ACC1-C56681BD9FAA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMemberResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMemberResponseBody) GetMembers() []*CreateMemberResponseBodyMembers {
	return s.Members
}

func (s *CreateMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMemberResponseBody) SetMembers(v []*CreateMemberResponseBodyMembers) *CreateMemberResponseBody {
	s.Members = v
	return s
}

func (s *CreateMemberResponseBody) SetRequestId(v string) *CreateMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMemberResponseBody) Validate() error {
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

type CreateMemberResponseBodyMembers struct {
	// The display name.
	//
	// example:
	//
	// myDisplayName
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The member ID.
	//
	// example:
	//
	// 145883-21513926******88039
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// The list of roles.
	Roles []*string `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// The user ID.
	//
	// example:
	//
	// 21513926******88039
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateMemberResponseBodyMembers) String() string {
	return dara.Prettify(s)
}

func (s CreateMemberResponseBodyMembers) GoString() string {
	return s.String()
}

func (s *CreateMemberResponseBodyMembers) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateMemberResponseBodyMembers) GetMemberId() *string {
	return s.MemberId
}

func (s *CreateMemberResponseBodyMembers) GetRoles() []*string {
	return s.Roles
}

func (s *CreateMemberResponseBodyMembers) GetUserId() *string {
	return s.UserId
}

func (s *CreateMemberResponseBodyMembers) SetDisplayName(v string) *CreateMemberResponseBodyMembers {
	s.DisplayName = &v
	return s
}

func (s *CreateMemberResponseBodyMembers) SetMemberId(v string) *CreateMemberResponseBodyMembers {
	s.MemberId = &v
	return s
}

func (s *CreateMemberResponseBodyMembers) SetRoles(v []*string) *CreateMemberResponseBodyMembers {
	s.Roles = v
	return s
}

func (s *CreateMemberResponseBodyMembers) SetUserId(v string) *CreateMemberResponseBodyMembers {
	s.UserId = &v
	return s
}

func (s *CreateMemberResponseBodyMembers) Validate() error {
	return dara.Validate(s)
}
