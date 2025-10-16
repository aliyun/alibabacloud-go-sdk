// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceMemberAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMembers(v []*ModifyInstanceMemberAttributesRequestMembers) *ModifyInstanceMemberAttributesRequest
	GetMembers() []*ModifyInstanceMemberAttributesRequestMembers
}

type ModifyInstanceMemberAttributesRequest struct {
	// The members that to be modified.
	//
	// This parameter is required.
	Members []*ModifyInstanceMemberAttributesRequestMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
}

func (s ModifyInstanceMemberAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceMemberAttributesRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMemberAttributesRequest) GetMembers() []*ModifyInstanceMemberAttributesRequestMembers {
	return s.Members
}

func (s *ModifyInstanceMemberAttributesRequest) SetMembers(v []*ModifyInstanceMemberAttributesRequestMembers) *ModifyInstanceMemberAttributesRequest {
	s.Members = v
	return s
}

func (s *ModifyInstanceMemberAttributesRequest) Validate() error {
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

type ModifyInstanceMemberAttributesRequestMembers struct {
	// The remarks of the member in Cloud Firewall.
	//
	// This parameter is required.
	//
	// example:
	//
	// renewal
	MemberDesc *string `json:"MemberDesc,omitempty" xml:"MemberDesc,omitempty"`
	// The UID of the member in Cloud Firewall.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123412341234****
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
}

func (s ModifyInstanceMemberAttributesRequestMembers) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceMemberAttributesRequestMembers) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMemberAttributesRequestMembers) GetMemberDesc() *string {
	return s.MemberDesc
}

func (s *ModifyInstanceMemberAttributesRequestMembers) GetMemberUid() *int64 {
	return s.MemberUid
}

func (s *ModifyInstanceMemberAttributesRequestMembers) SetMemberDesc(v string) *ModifyInstanceMemberAttributesRequestMembers {
	s.MemberDesc = &v
	return s
}

func (s *ModifyInstanceMemberAttributesRequestMembers) SetMemberUid(v int64) *ModifyInstanceMemberAttributesRequestMembers {
	s.MemberUid = &v
	return s
}

func (s *ModifyInstanceMemberAttributesRequestMembers) Validate() error {
	return dara.Validate(s)
}
