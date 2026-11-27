// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeChannelUserRequest
	GetAppId() *string
	SetChannelId(v string) *DescribeChannelUserRequest
	GetChannelId() *string
	SetUserId(v string) *DescribeChannelUserRequest
	GetUserId() *string
}

type DescribeChannelUserRequest struct {
	// The application ID. You can create and query the ID in the console. Only a single ID is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the channel to be queried. Only a single ID is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1811****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeChannelUserRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelUserRequest) GoString() string {
	return s.String()
}

func (s *DescribeChannelUserRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeChannelUserRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeChannelUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *DescribeChannelUserRequest) SetAppId(v string) *DescribeChannelUserRequest {
	s.AppId = &v
	return s
}

func (s *DescribeChannelUserRequest) SetChannelId(v string) *DescribeChannelUserRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeChannelUserRequest) SetUserId(v string) *DescribeChannelUserRequest {
	s.UserId = &v
	return s
}

func (s *DescribeChannelUserRequest) Validate() error {
	return dara.Validate(s)
}
