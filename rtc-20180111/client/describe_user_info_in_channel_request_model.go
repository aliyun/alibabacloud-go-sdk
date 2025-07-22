// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserInfoInChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeUserInfoInChannelRequest
	GetAppId() *string
	SetChannelId(v string) *DescribeUserInfoInChannelRequest
	GetChannelId() *string
	SetOwnerId(v int64) *DescribeUserInfoInChannelRequest
	GetOwnerId() *int64
	SetUserId(v string) *DescribeUserInfoInChannelRequest
	GetUserId() *string
}

type DescribeUserInfoInChannelRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4eah****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testId
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeUserInfoInChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserInfoInChannelRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserInfoInChannelRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeUserInfoInChannelRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeUserInfoInChannelRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeUserInfoInChannelRequest) GetUserId() *string {
	return s.UserId
}

func (s *DescribeUserInfoInChannelRequest) SetAppId(v string) *DescribeUserInfoInChannelRequest {
	s.AppId = &v
	return s
}

func (s *DescribeUserInfoInChannelRequest) SetChannelId(v string) *DescribeUserInfoInChannelRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeUserInfoInChannelRequest) SetOwnerId(v int64) *DescribeUserInfoInChannelRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUserInfoInChannelRequest) SetUserId(v string) *DescribeUserInfoInChannelRequest {
	s.UserId = &v
	return s
}

func (s *DescribeUserInfoInChannelRequest) Validate() error {
	return dara.Validate(s)
}
