// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRdMemberListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMemberList(v []*DescribeRdMemberListResponseBodyMemberList) *DescribeRdMemberListResponseBody
	GetMemberList() []*DescribeRdMemberListResponseBodyMemberList
	SetRequestId(v string) *DescribeRdMemberListResponseBody
	GetRequestId() *string
	SetTotal(v int64) *DescribeRdMemberListResponseBody
	GetTotal() *int64
}

type DescribeRdMemberListResponseBody struct {
	// The information about the members.
	MemberList []*DescribeRdMemberListResponseBodyMemberList `json:"MemberList,omitempty" xml:"MemberList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// DC245DEE-9800-5579-BF99-189D6A5BA9FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeRdMemberListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdMemberListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRdMemberListResponseBody) GetMemberList() []*DescribeRdMemberListResponseBodyMemberList {
	return s.MemberList
}

func (s *DescribeRdMemberListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRdMemberListResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeRdMemberListResponseBody) SetMemberList(v []*DescribeRdMemberListResponseBodyMemberList) *DescribeRdMemberListResponseBody {
	s.MemberList = v
	return s
}

func (s *DescribeRdMemberListResponseBody) SetRequestId(v string) *DescribeRdMemberListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRdMemberListResponseBody) SetTotal(v int64) *DescribeRdMemberListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeRdMemberListResponseBody) Validate() error {
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

type DescribeRdMemberListResponseBodyMemberList struct {
	// The creation time.
	//
	// example:
	//
	// 1624954942000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The name of the member.
	//
	// example:
	//
	// test1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The Alibaba Cloud account ID of the member.
	//
	// example:
	//
	// 1960279802016267
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s DescribeRdMemberListResponseBodyMemberList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdMemberListResponseBodyMemberList) GoString() string {
	return s.String()
}

func (s *DescribeRdMemberListResponseBodyMemberList) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeRdMemberListResponseBodyMemberList) GetName() *string {
	return s.Name
}

func (s *DescribeRdMemberListResponseBodyMemberList) GetUid() *string {
	return s.Uid
}

func (s *DescribeRdMemberListResponseBodyMemberList) SetGmtCreate(v int64) *DescribeRdMemberListResponseBodyMemberList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeRdMemberListResponseBodyMemberList) SetName(v string) *DescribeRdMemberListResponseBodyMemberList {
	s.Name = &v
	return s
}

func (s *DescribeRdMemberListResponseBodyMemberList) SetUid(v string) *DescribeRdMemberListResponseBodyMemberList {
	s.Uid = &v
	return s
}

func (s *DescribeRdMemberListResponseBodyMemberList) Validate() error {
	return dara.Validate(s)
}
