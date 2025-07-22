// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelParticipantsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeChannelParticipantsResponseBody
	GetRequestId() *string
	SetTimestamp(v int32) *DescribeChannelParticipantsResponseBody
	GetTimestamp() *int32
	SetTotalNum(v int32) *DescribeChannelParticipantsResponseBody
	GetTotalNum() *int32
	SetTotalPage(v int32) *DescribeChannelParticipantsResponseBody
	GetTotalPage() *int32
	SetUserList(v *DescribeChannelParticipantsResponseBodyUserList) *DescribeChannelParticipantsResponseBody
	GetUserList() *DescribeChannelParticipantsResponseBodyUserList
}

type DescribeChannelParticipantsResponseBody struct {
	// example:
	//
	// 6159ba01-6687-4fb2-a831-f0cd8d188648
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1557909133
	Timestamp *int32 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// example:
	//
	// 3
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// example:
	//
	// 1
	TotalPage *int32                                           `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	UserList  *DescribeChannelParticipantsResponseBodyUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Struct"`
}

func (s DescribeChannelParticipantsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelParticipantsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChannelParticipantsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeChannelParticipantsResponseBody) GetTimestamp() *int32 {
	return s.Timestamp
}

func (s *DescribeChannelParticipantsResponseBody) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *DescribeChannelParticipantsResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeChannelParticipantsResponseBody) GetUserList() *DescribeChannelParticipantsResponseBodyUserList {
	return s.UserList
}

func (s *DescribeChannelParticipantsResponseBody) SetRequestId(v string) *DescribeChannelParticipantsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChannelParticipantsResponseBody) SetTimestamp(v int32) *DescribeChannelParticipantsResponseBody {
	s.Timestamp = &v
	return s
}

func (s *DescribeChannelParticipantsResponseBody) SetTotalNum(v int32) *DescribeChannelParticipantsResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeChannelParticipantsResponseBody) SetTotalPage(v int32) *DescribeChannelParticipantsResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeChannelParticipantsResponseBody) SetUserList(v *DescribeChannelParticipantsResponseBodyUserList) *DescribeChannelParticipantsResponseBody {
	s.UserList = v
	return s
}

func (s *DescribeChannelParticipantsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeChannelParticipantsResponseBodyUserList struct {
	User []*string `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s DescribeChannelParticipantsResponseBodyUserList) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelParticipantsResponseBodyUserList) GoString() string {
	return s.String()
}

func (s *DescribeChannelParticipantsResponseBodyUserList) GetUser() []*string {
	return s.User
}

func (s *DescribeChannelParticipantsResponseBodyUserList) SetUser(v []*string) *DescribeChannelParticipantsResponseBodyUserList {
	s.User = v
	return s
}

func (s *DescribeChannelParticipantsResponseBodyUserList) Validate() error {
	return dara.Validate(s)
}
