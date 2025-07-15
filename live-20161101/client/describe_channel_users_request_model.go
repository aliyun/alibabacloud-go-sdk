// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeChannelUsersRequest
	GetAppId() *string
	SetChannelId(v string) *DescribeChannelUsersRequest
	GetChannelId() *string
}

type DescribeChannelUsersRequest struct {
	// The application ID. You can specify only one application ID in a request.
	//
	// This parameter is required.
	//
	// example:
	//
	// aec****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The channel ID. You can specify only one channel ID in a request.
	//
	// This parameter is required.
	//
	// example:
	//
	// testId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
}

func (s DescribeChannelUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelUsersRequest) GoString() string {
	return s.String()
}

func (s *DescribeChannelUsersRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeChannelUsersRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeChannelUsersRequest) SetAppId(v string) *DescribeChannelUsersRequest {
	s.AppId = &v
	return s
}

func (s *DescribeChannelUsersRequest) SetChannelId(v string) *DescribeChannelUsersRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeChannelUsersRequest) Validate() error {
	return dara.Validate(s)
}
