// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChannel(v *DescribeChannelResponseBodyChannel) *DescribeChannelResponseBody
	GetChannel() *DescribeChannelResponseBodyChannel
	SetChannelExist(v bool) *DescribeChannelResponseBody
	GetChannelExist() *bool
	SetRequestId(v string) *DescribeChannelResponseBody
	GetRequestId() *string
}

type DescribeChannelResponseBody struct {
	// channel
	Channel      *DescribeChannelResponseBodyChannel `json:"Channel,omitempty" xml:"Channel,omitempty" type:"Struct"`
	ChannelExist *bool                               `json:"ChannelExist,omitempty" xml:"ChannelExist,omitempty"`
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChannelResponseBody) GetChannel() *DescribeChannelResponseBodyChannel {
	return s.Channel
}

func (s *DescribeChannelResponseBody) GetChannelExist() *bool {
	return s.ChannelExist
}

func (s *DescribeChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeChannelResponseBody) SetChannel(v *DescribeChannelResponseBodyChannel) *DescribeChannelResponseBody {
	s.Channel = v
	return s
}

func (s *DescribeChannelResponseBody) SetChannelExist(v bool) *DescribeChannelResponseBody {
	s.ChannelExist = &v
	return s
}

func (s *DescribeChannelResponseBody) SetRequestId(v string) *DescribeChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChannelResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeChannelResponseBodyChannel struct {
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// 1557909133
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeChannelResponseBodyChannel) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelResponseBodyChannel) GoString() string {
	return s.String()
}

func (s *DescribeChannelResponseBodyChannel) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeChannelResponseBodyChannel) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeChannelResponseBodyChannel) SetChannelId(v string) *DescribeChannelResponseBodyChannel {
	s.ChannelId = &v
	return s
}

func (s *DescribeChannelResponseBodyChannel) SetStartTime(v int64) *DescribeChannelResponseBodyChannel {
	s.StartTime = &v
	return s
}

func (s *DescribeChannelResponseBodyChannel) Validate() error {
	return dara.Validate(s)
}
