// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChannelProfile(v int32) *DescribeChannelUsersResponseBody
	GetChannelProfile() *int32
	SetCommTotalNum(v int32) *DescribeChannelUsersResponseBody
	GetCommTotalNum() *int32
	SetInteractiveUserList(v []*string) *DescribeChannelUsersResponseBody
	GetInteractiveUserList() []*string
	SetInteractiveUserNum(v int32) *DescribeChannelUsersResponseBody
	GetInteractiveUserNum() *int32
	SetIsChannelExist(v bool) *DescribeChannelUsersResponseBody
	GetIsChannelExist() *bool
	SetLiveUserList(v []*string) *DescribeChannelUsersResponseBody
	GetLiveUserList() []*string
	SetLiveUserNum(v int32) *DescribeChannelUsersResponseBody
	GetLiveUserNum() *int32
	SetRequestId(v string) *DescribeChannelUsersResponseBody
	GetRequestId() *string
	SetTimestamp(v int32) *DescribeChannelUsersResponseBody
	GetTimestamp() *int32
	SetUserList(v []*string) *DescribeChannelUsersResponseBody
	GetUserList() []*string
}

type DescribeChannelUsersResponseBody struct {
	// example:
	//
	// 1
	ChannelProfile *int32 `json:"ChannelProfile,omitempty" xml:"ChannelProfile,omitempty"`
	// example:
	//
	// 100
	CommTotalNum        *int32    `json:"CommTotalNum,omitempty" xml:"CommTotalNum,omitempty"`
	InteractiveUserList []*string `json:"InteractiveUserList,omitempty" xml:"InteractiveUserList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	InteractiveUserNum *int32 `json:"InteractiveUserNum,omitempty" xml:"InteractiveUserNum,omitempty"`
	// example:
	//
	// true
	IsChannelExist *bool     `json:"IsChannelExist,omitempty" xml:"IsChannelExist,omitempty"`
	LiveUserList   []*string `json:"LiveUserList,omitempty" xml:"LiveUserList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	LiveUserNum *int32 `json:"LiveUserNum,omitempty" xml:"LiveUserNum,omitempty"`
	// example:
	//
	// 6159ba01-6687-4fb2-a831-f0cd8d188648
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1557909133
	Timestamp *int32    `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	UserList  []*string `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s DescribeChannelUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChannelUsersResponseBody) GetChannelProfile() *int32 {
	return s.ChannelProfile
}

func (s *DescribeChannelUsersResponseBody) GetCommTotalNum() *int32 {
	return s.CommTotalNum
}

func (s *DescribeChannelUsersResponseBody) GetInteractiveUserList() []*string {
	return s.InteractiveUserList
}

func (s *DescribeChannelUsersResponseBody) GetInteractiveUserNum() *int32 {
	return s.InteractiveUserNum
}

func (s *DescribeChannelUsersResponseBody) GetIsChannelExist() *bool {
	return s.IsChannelExist
}

func (s *DescribeChannelUsersResponseBody) GetLiveUserList() []*string {
	return s.LiveUserList
}

func (s *DescribeChannelUsersResponseBody) GetLiveUserNum() *int32 {
	return s.LiveUserNum
}

func (s *DescribeChannelUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeChannelUsersResponseBody) GetTimestamp() *int32 {
	return s.Timestamp
}

func (s *DescribeChannelUsersResponseBody) GetUserList() []*string {
	return s.UserList
}

func (s *DescribeChannelUsersResponseBody) SetChannelProfile(v int32) *DescribeChannelUsersResponseBody {
	s.ChannelProfile = &v
	return s
}

func (s *DescribeChannelUsersResponseBody) SetCommTotalNum(v int32) *DescribeChannelUsersResponseBody {
	s.CommTotalNum = &v
	return s
}

func (s *DescribeChannelUsersResponseBody) SetInteractiveUserList(v []*string) *DescribeChannelUsersResponseBody {
	s.InteractiveUserList = v
	return s
}

func (s *DescribeChannelUsersResponseBody) SetInteractiveUserNum(v int32) *DescribeChannelUsersResponseBody {
	s.InteractiveUserNum = &v
	return s
}

func (s *DescribeChannelUsersResponseBody) SetIsChannelExist(v bool) *DescribeChannelUsersResponseBody {
	s.IsChannelExist = &v
	return s
}

func (s *DescribeChannelUsersResponseBody) SetLiveUserList(v []*string) *DescribeChannelUsersResponseBody {
	s.LiveUserList = v
	return s
}

func (s *DescribeChannelUsersResponseBody) SetLiveUserNum(v int32) *DescribeChannelUsersResponseBody {
	s.LiveUserNum = &v
	return s
}

func (s *DescribeChannelUsersResponseBody) SetRequestId(v string) *DescribeChannelUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChannelUsersResponseBody) SetTimestamp(v int32) *DescribeChannelUsersResponseBody {
	s.Timestamp = &v
	return s
}

func (s *DescribeChannelUsersResponseBody) SetUserList(v []*string) *DescribeChannelUsersResponseBody {
	s.UserList = v
	return s
}

func (s *DescribeChannelUsersResponseBody) Validate() error {
	return dara.Validate(s)
}
