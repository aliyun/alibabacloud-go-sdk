// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMemberInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMemberUid(v string) *DescribeMemberInfoRequest
	GetMemberUid() *string
}

type DescribeMemberInfoRequest struct {
	// example:
	//
	// 150795602499****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
}

func (s DescribeMemberInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMemberInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeMemberInfoRequest) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribeMemberInfoRequest) SetMemberUid(v string) *DescribeMemberInfoRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeMemberInfoRequest) Validate() error {
	return dara.Validate(s)
}
