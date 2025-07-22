// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelAllUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeChannelAllUsersRequest
	GetAppId() *string
	SetChannelId(v string) *DescribeChannelAllUsersRequest
	GetChannelId() *string
}

type DescribeChannelAllUsersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
}

func (s DescribeChannelAllUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelAllUsersRequest) GoString() string {
	return s.String()
}

func (s *DescribeChannelAllUsersRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeChannelAllUsersRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeChannelAllUsersRequest) SetAppId(v string) *DescribeChannelAllUsersRequest {
	s.AppId = &v
	return s
}

func (s *DescribeChannelAllUsersRequest) SetChannelId(v string) *DescribeChannelAllUsersRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeChannelAllUsersRequest) Validate() error {
	return dara.Validate(s)
}
