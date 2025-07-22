// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelAllUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChannelExist(v bool) *DescribeChannelAllUsersResponseBody
	GetChannelExist() *bool
	SetRequestId(v string) *DescribeChannelAllUsersResponseBody
	GetRequestId() *string
	SetUsers(v []*DescribeChannelAllUsersResponseBodyUsers) *DescribeChannelAllUsersResponseBody
	GetUsers() []*DescribeChannelAllUsersResponseBodyUsers
}

type DescribeChannelAllUsersResponseBody struct {
	ChannelExist *bool `json:"ChannelExist,omitempty" xml:"ChannelExist,omitempty"`
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Users     []*DescribeChannelAllUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s DescribeChannelAllUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelAllUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChannelAllUsersResponseBody) GetChannelExist() *bool {
	return s.ChannelExist
}

func (s *DescribeChannelAllUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeChannelAllUsersResponseBody) GetUsers() []*DescribeChannelAllUsersResponseBodyUsers {
	return s.Users
}

func (s *DescribeChannelAllUsersResponseBody) SetChannelExist(v bool) *DescribeChannelAllUsersResponseBody {
	s.ChannelExist = &v
	return s
}

func (s *DescribeChannelAllUsersResponseBody) SetRequestId(v string) *DescribeChannelAllUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChannelAllUsersResponseBody) SetUsers(v []*DescribeChannelAllUsersResponseBodyUsers) *DescribeChannelAllUsersResponseBody {
	s.Users = v
	return s
}

func (s *DescribeChannelAllUsersResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeChannelAllUsersResponseBodyUsers struct {
	// example:
	//
	// 1811****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeChannelAllUsersResponseBodyUsers) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelAllUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *DescribeChannelAllUsersResponseBodyUsers) GetUserId() *string {
	return s.UserId
}

func (s *DescribeChannelAllUsersResponseBodyUsers) SetUserId(v string) *DescribeChannelAllUsersResponseBodyUsers {
	s.UserId = &v
	return s
}

func (s *DescribeChannelAllUsersResponseBodyUsers) Validate() error {
	return dara.Validate(s)
}
