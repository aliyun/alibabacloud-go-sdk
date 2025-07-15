// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInteractiveUserList(v []*string) *DescribeChannelUsersResponseBody
	GetInteractiveUserList() []*string
	SetInteractiveUserNumber(v int32) *DescribeChannelUsersResponseBody
	GetInteractiveUserNumber() *int32
	SetIsChannelExists(v bool) *DescribeChannelUsersResponseBody
	GetIsChannelExists() *bool
	SetLiveUserList(v []*string) *DescribeChannelUsersResponseBody
	GetLiveUserList() []*string
	SetLiveUserNumber(v int32) *DescribeChannelUsersResponseBody
	GetLiveUserNumber() *int32
	SetRequestId(v string) *DescribeChannelUsersResponseBody
	GetRequestId() *string
	SetTimestamp(v int64) *DescribeChannelUsersResponseBody
	GetTimestamp() *int64
}

type DescribeChannelUsersResponseBody struct {
	// The list of streamers/co-streamers.
	InteractiveUserList []*string `json:"InteractiveUserList,omitempty" xml:"InteractiveUserList,omitempty" type:"Repeated"`
	// The number of co-streamers.
	//
	// example:
	//
	// 1
	InteractiveUserNumber *int32 `json:"InteractiveUserNumber,omitempty" xml:"InteractiveUserNumber,omitempty"`
	// Indicates whether the channel exists. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// > After all users leave the channel, the system requires a few seconds to clear the cache. If you call the operation during this period, the value of this parameter is true, and the value of InteractiveUserNumber and LiveUserNumber is 0.
	//
	// example:
	//
	// true
	IsChannelExists *bool `json:"IsChannelExists,omitempty" xml:"IsChannelExists,omitempty"`
	// The list of viewers.
	LiveUserList []*string `json:"LiveUserList,omitempty" xml:"LiveUserList,omitempty" type:"Repeated"`
	// The number of viewers.
	//
	// example:
	//
	// 1
	LiveUserNumber *int32 `json:"LiveUserNumber,omitempty" xml:"LiveUserNumber,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// AD86847D-1F36-18C8-A995-5EEA340B3202
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The UTC timestamp when the response is returned.
	//
	// example:
	//
	// 1691027655
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeChannelUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChannelUsersResponseBody) GetInteractiveUserList() []*string {
	return s.InteractiveUserList
}

func (s *DescribeChannelUsersResponseBody) GetInteractiveUserNumber() *int32 {
	return s.InteractiveUserNumber
}

func (s *DescribeChannelUsersResponseBody) GetIsChannelExists() *bool {
	return s.IsChannelExists
}

func (s *DescribeChannelUsersResponseBody) GetLiveUserList() []*string {
	return s.LiveUserList
}

func (s *DescribeChannelUsersResponseBody) GetLiveUserNumber() *int32 {
	return s.LiveUserNumber
}

func (s *DescribeChannelUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeChannelUsersResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeChannelUsersResponseBody) SetInteractiveUserList(v []*string) *DescribeChannelUsersResponseBody {
	s.InteractiveUserList = v
	return s
}

func (s *DescribeChannelUsersResponseBody) SetInteractiveUserNumber(v int32) *DescribeChannelUsersResponseBody {
	s.InteractiveUserNumber = &v
	return s
}

func (s *DescribeChannelUsersResponseBody) SetIsChannelExists(v bool) *DescribeChannelUsersResponseBody {
	s.IsChannelExists = &v
	return s
}

func (s *DescribeChannelUsersResponseBody) SetLiveUserList(v []*string) *DescribeChannelUsersResponseBody {
	s.LiveUserList = v
	return s
}

func (s *DescribeChannelUsersResponseBody) SetLiveUserNumber(v int32) *DescribeChannelUsersResponseBody {
	s.LiveUserNumber = &v
	return s
}

func (s *DescribeChannelUsersResponseBody) SetRequestId(v string) *DescribeChannelUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChannelUsersResponseBody) SetTimestamp(v int64) *DescribeChannelUsersResponseBody {
	s.Timestamp = &v
	return s
}

func (s *DescribeChannelUsersResponseBody) Validate() error {
	return dara.Validate(s)
}
