// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRdMemberListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMemberList(v []*AddRdMemberListRequestMemberList) *AddRdMemberListRequest
	GetMemberList() []*AddRdMemberListRequestMemberList
}

type AddRdMemberListRequest struct {
	// The list of the members.
	//
	// This parameter is required.
	MemberList []*AddRdMemberListRequestMemberList `json:"MemberList,omitempty" xml:"MemberList,omitempty" type:"Repeated"`
}

func (s AddRdMemberListRequest) String() string {
	return dara.Prettify(s)
}

func (s AddRdMemberListRequest) GoString() string {
	return s.String()
}

func (s *AddRdMemberListRequest) GetMemberList() []*AddRdMemberListRequestMemberList {
	return s.MemberList
}

func (s *AddRdMemberListRequest) SetMemberList(v []*AddRdMemberListRequestMemberList) *AddRdMemberListRequest {
	s.MemberList = v
	return s
}

func (s *AddRdMemberListRequest) Validate() error {
	if s.MemberList != nil {
		for _, item := range s.MemberList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddRdMemberListRequestMemberList struct {
	// The Alibaba Cloud account ID of the member.
	//
	// example:
	//
	// 19510843762****
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s AddRdMemberListRequestMemberList) String() string {
	return dara.Prettify(s)
}

func (s AddRdMemberListRequestMemberList) GoString() string {
	return s.String()
}

func (s *AddRdMemberListRequestMemberList) GetUid() *string {
	return s.Uid
}

func (s *AddRdMemberListRequestMemberList) SetUid(v string) *AddRdMemberListRequestMemberList {
	s.Uid = &v
	return s
}

func (s *AddRdMemberListRequestMemberList) Validate() error {
	return dara.Validate(s)
}
