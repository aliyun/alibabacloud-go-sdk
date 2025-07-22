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
	SetOwnerId(v int64) *DescribeChannelUsersRequest
	GetOwnerId() *int64
}

type DescribeChannelUsersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a2hz****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s *DescribeChannelUsersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeChannelUsersRequest) SetAppId(v string) *DescribeChannelUsersRequest {
	s.AppId = &v
	return s
}

func (s *DescribeChannelUsersRequest) SetChannelId(v string) *DescribeChannelUsersRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeChannelUsersRequest) SetOwnerId(v int64) *DescribeChannelUsersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeChannelUsersRequest) Validate() error {
	return dara.Validate(s)
}
