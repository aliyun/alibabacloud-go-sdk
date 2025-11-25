// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMemberInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAdminName(v string) *DescribeMemberInfoResponseBody
	GetAdminName() *string
	SetAdminUid(v string) *DescribeMemberInfoResponseBody
	GetAdminUid() *string
	SetIsMember(v bool) *DescribeMemberInfoResponseBody
	GetIsMember() *bool
	SetMemberUid(v string) *DescribeMemberInfoResponseBody
	GetMemberUid() *string
	SetRequestId(v string) *DescribeMemberInfoResponseBody
	GetRequestId() *string
}

type DescribeMemberInfoResponseBody struct {
	AdminName *string `json:"AdminName,omitempty" xml:"AdminName,omitempty"`
	// example:
	//
	// 164705101205****
	AdminUid *string `json:"AdminUid,omitempty" xml:"AdminUid,omitempty"`
	// example:
	//
	// true
	IsMember *bool `json:"IsMember,omitempty" xml:"IsMember,omitempty"`
	// example:
	//
	// 128720273643****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// example:
	//
	// F2665618-3C41-51A4-8DAF-586FB68****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMemberInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMemberInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMemberInfoResponseBody) GetAdminName() *string {
	return s.AdminName
}

func (s *DescribeMemberInfoResponseBody) GetAdminUid() *string {
	return s.AdminUid
}

func (s *DescribeMemberInfoResponseBody) GetIsMember() *bool {
	return s.IsMember
}

func (s *DescribeMemberInfoResponseBody) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribeMemberInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMemberInfoResponseBody) SetAdminName(v string) *DescribeMemberInfoResponseBody {
	s.AdminName = &v
	return s
}

func (s *DescribeMemberInfoResponseBody) SetAdminUid(v string) *DescribeMemberInfoResponseBody {
	s.AdminUid = &v
	return s
}

func (s *DescribeMemberInfoResponseBody) SetIsMember(v bool) *DescribeMemberInfoResponseBody {
	s.IsMember = &v
	return s
}

func (s *DescribeMemberInfoResponseBody) SetMemberUid(v string) *DescribeMemberInfoResponseBody {
	s.MemberUid = &v
	return s
}

func (s *DescribeMemberInfoResponseBody) SetRequestId(v string) *DescribeMemberInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMemberInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
