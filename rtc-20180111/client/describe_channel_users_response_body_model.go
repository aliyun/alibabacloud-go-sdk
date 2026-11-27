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
	// The channel mode. Valid values:
	//
	// - **1**: Communication (communication mode).
	//
	// - **2**: Interactive_live (live streaming mode).
	//
	// example:
	//
	// 1
	ChannelProfile *int32 `json:"ChannelProfile,omitempty" xml:"ChannelProfile,omitempty"`
	// The total number of users in the channel in communication mode. If the channel is not in communication mode, this parameter is 0.
	//
	// example:
	//
	// 100
	CommTotalNum *int32 `json:"CommTotalNum,omitempty" xml:"CommTotalNum,omitempty"`
	// The list of interactive user IDs in the channel in live streaming mode. This parameter is empty if the channel is not in live streaming mode.
	InteractiveUserList []*string `json:"InteractiveUserList,omitempty" xml:"InteractiveUserList,omitempty" type:"Repeated"`
	// The total number of interactive users in the channel in live streaming mode. If the channel is not in live streaming mode, this parameter is 0.
	//
	// example:
	//
	// 0
	InteractiveUserNum *int32 `json:"InteractiveUserNum,omitempty" xml:"InteractiveUserNum,omitempty"`
	// Indicates whether the channel exists. Valid values:
	//
	// - **true**: The channel exists.
	//
	// - **false**: The channel does not exist.
	//
	// example:
	//
	// true
	IsChannelExist *bool `json:"IsChannelExist,omitempty" xml:"IsChannelExist,omitempty"`
	// The list of viewer IDs in the channel in live streaming mode. A maximum of 1,000 viewers are returned. This parameter is empty if the channel is not in live streaming mode.
	LiveUserList []*string `json:"LiveUserList,omitempty" xml:"LiveUserList,omitempty" type:"Repeated"`
	// The total number of viewers in the channel in live streaming mode. If the channel is not in live streaming mode, this parameter is 0.
	//
	// example:
	//
	// 0
	LiveUserNum *int32 `json:"LiveUserNum,omitempty" xml:"LiveUserNum,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6159ba01-6687-4fb2-a831-f0cd8d188648
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The timestamp of the current time in UTC.
	//
	// example:
	//
	// 1557909133
	Timestamp *int32 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The list of user IDs in the channel in communication mode. This parameter is empty if the channel is not in communication mode.
	UserList []*string `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
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
