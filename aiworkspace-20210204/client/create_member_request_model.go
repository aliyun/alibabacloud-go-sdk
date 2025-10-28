// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMembers(v []*CreateMemberRequestMembers) *CreateMemberRequest
	GetMembers() []*CreateMemberRequestMembers
}

type CreateMemberRequest struct {
	// The members.
	//
	// This parameter is required.
	Members []*CreateMemberRequestMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
}

func (s CreateMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMemberRequest) GoString() string {
	return s.String()
}

func (s *CreateMemberRequest) GetMembers() []*CreateMemberRequestMembers {
	return s.Members
}

func (s *CreateMemberRequest) SetMembers(v []*CreateMemberRequestMembers) *CreateMemberRequest {
	s.Members = v
	return s
}

func (s *CreateMemberRequest) Validate() error {
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

type CreateMemberRequestMembers struct {
	// The list of roles.
	//
	// This parameter is required.
	Roles []*string `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// The member IDs. Multiple member IDs are separated by commas (,). You can call [ListMembers](https://help.aliyun.com/document_detail/449135.html) to obtain the member IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 21513926******88039
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateMemberRequestMembers) String() string {
	return dara.Prettify(s)
}

func (s CreateMemberRequestMembers) GoString() string {
	return s.String()
}

func (s *CreateMemberRequestMembers) GetRoles() []*string {
	return s.Roles
}

func (s *CreateMemberRequestMembers) GetUserId() *string {
	return s.UserId
}

func (s *CreateMemberRequestMembers) SetRoles(v []*string) *CreateMemberRequestMembers {
	s.Roles = v
	return s
}

func (s *CreateMemberRequestMembers) SetUserId(v string) *CreateMemberRequestMembers {
	s.UserId = &v
	return s
}

func (s *CreateMemberRequestMembers) Validate() error {
	return dara.Validate(s)
}
