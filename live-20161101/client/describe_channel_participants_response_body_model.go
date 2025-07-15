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
	SetTimes(v int64) *DescribeChannelParticipantsResponseBody
	GetTimes() *int64
	SetTotalNum(v int32) *DescribeChannelParticipantsResponseBody
	GetTotalNum() *int32
	SetTotalPage(v int32) *DescribeChannelParticipantsResponseBody
	GetTotalPage() *int32
	SetUserList(v []*string) *DescribeChannelParticipantsResponseBody
	GetUserList() []*string
}

type DescribeChannelParticipantsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6159ba01-6687-4fb2-a831-f0cd8d188648
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time elapsed until the result was returned. Unit: seconds. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1557909133
	Times *int64 `json:"Times,omitempty" xml:"Times,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 3
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	// The list of user IDs.
	UserList []*string `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
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

func (s *DescribeChannelParticipantsResponseBody) GetTimes() *int64 {
	return s.Times
}

func (s *DescribeChannelParticipantsResponseBody) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *DescribeChannelParticipantsResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeChannelParticipantsResponseBody) GetUserList() []*string {
	return s.UserList
}

func (s *DescribeChannelParticipantsResponseBody) SetRequestId(v string) *DescribeChannelParticipantsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChannelParticipantsResponseBody) SetTimes(v int64) *DescribeChannelParticipantsResponseBody {
	s.Times = &v
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

func (s *DescribeChannelParticipantsResponseBody) SetUserList(v []*string) *DescribeChannelParticipantsResponseBody {
	s.UserList = v
	return s
}

func (s *DescribeChannelParticipantsResponseBody) Validate() error {
	return dara.Validate(s)
}
