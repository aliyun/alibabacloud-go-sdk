// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRdMemberListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMemberList(v []*DeleteRdMemberListRequestMemberList) *DeleteRdMemberListRequest
	GetMemberList() []*DeleteRdMemberListRequestMemberList
}

type DeleteRdMemberListRequest struct {
	// The list of the members.
	//
	// This parameter is required.
	MemberList []*DeleteRdMemberListRequestMemberList `json:"MemberList,omitempty" xml:"MemberList,omitempty" type:"Repeated"`
}

func (s DeleteRdMemberListRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRdMemberListRequest) GoString() string {
	return s.String()
}

func (s *DeleteRdMemberListRequest) GetMemberList() []*DeleteRdMemberListRequestMemberList {
	return s.MemberList
}

func (s *DeleteRdMemberListRequest) SetMemberList(v []*DeleteRdMemberListRequestMemberList) *DeleteRdMemberListRequest {
	s.MemberList = v
	return s
}

func (s *DeleteRdMemberListRequest) Validate() error {
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

type DeleteRdMemberListRequestMemberList struct {
	// The Alibaba Cloud account ID of the member.
	//
	// example:
	//
	// 136548010379****
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s DeleteRdMemberListRequestMemberList) String() string {
	return dara.Prettify(s)
}

func (s DeleteRdMemberListRequestMemberList) GoString() string {
	return s.String()
}

func (s *DeleteRdMemberListRequestMemberList) GetUid() *string {
	return s.Uid
}

func (s *DeleteRdMemberListRequestMemberList) SetUid(v string) *DeleteRdMemberListRequestMemberList {
	s.Uid = &v
	return s
}

func (s *DeleteRdMemberListRequestMemberList) Validate() error {
	return dara.Validate(s)
}
