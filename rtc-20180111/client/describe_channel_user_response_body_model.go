// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChannelExist(v bool) *DescribeChannelUserResponseBody
	GetChannelExist() *bool
	SetInChannel(v bool) *DescribeChannelUserResponseBody
	GetInChannel() *bool
	SetRequestId(v string) *DescribeChannelUserResponseBody
	GetRequestId() *string
	SetSessions(v []*DescribeChannelUserResponseBodySessions) *DescribeChannelUserResponseBody
	GetSessions() []*DescribeChannelUserResponseBodySessions
}

type DescribeChannelUserResponseBody struct {
	ChannelExist *bool `json:"ChannelExist,omitempty" xml:"ChannelExist,omitempty"`
	InChannel    *bool `json:"InChannel,omitempty" xml:"InChannel,omitempty"`
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Sessions  []*DescribeChannelUserResponseBodySessions `json:"Sessions,omitempty" xml:"Sessions,omitempty" type:"Repeated"`
}

func (s DescribeChannelUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelUserResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChannelUserResponseBody) GetChannelExist() *bool {
	return s.ChannelExist
}

func (s *DescribeChannelUserResponseBody) GetInChannel() *bool {
	return s.InChannel
}

func (s *DescribeChannelUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeChannelUserResponseBody) GetSessions() []*DescribeChannelUserResponseBodySessions {
	return s.Sessions
}

func (s *DescribeChannelUserResponseBody) SetChannelExist(v bool) *DescribeChannelUserResponseBody {
	s.ChannelExist = &v
	return s
}

func (s *DescribeChannelUserResponseBody) SetInChannel(v bool) *DescribeChannelUserResponseBody {
	s.InChannel = &v
	return s
}

func (s *DescribeChannelUserResponseBody) SetRequestId(v string) *DescribeChannelUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChannelUserResponseBody) SetSessions(v []*DescribeChannelUserResponseBodySessions) *DescribeChannelUserResponseBody {
	s.Sessions = v
	return s
}

func (s *DescribeChannelUserResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeChannelUserResponseBodySessions struct {
	// example:
	//
	// 1557909133
	Joined *int64 `json:"Joined,omitempty" xml:"Joined,omitempty"`
	// example:
	//
	// xa744sxx8rtobgj****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 1811****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeChannelUserResponseBodySessions) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelUserResponseBodySessions) GoString() string {
	return s.String()
}

func (s *DescribeChannelUserResponseBodySessions) GetJoined() *int64 {
	return s.Joined
}

func (s *DescribeChannelUserResponseBodySessions) GetSessionId() *string {
	return s.SessionId
}

func (s *DescribeChannelUserResponseBodySessions) GetUserId() *string {
	return s.UserId
}

func (s *DescribeChannelUserResponseBodySessions) SetJoined(v int64) *DescribeChannelUserResponseBodySessions {
	s.Joined = &v
	return s
}

func (s *DescribeChannelUserResponseBodySessions) SetSessionId(v string) *DescribeChannelUserResponseBodySessions {
	s.SessionId = &v
	return s
}

func (s *DescribeChannelUserResponseBodySessions) SetUserId(v string) *DescribeChannelUserResponseBodySessions {
	s.UserId = &v
	return s
}

func (s *DescribeChannelUserResponseBodySessions) Validate() error {
	return dara.Validate(s)
}
