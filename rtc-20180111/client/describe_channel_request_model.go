// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeChannelRequest
	GetAppId() *string
	SetChannelId(v string) *DescribeChannelRequest
	GetChannelId() *string
}

type DescribeChannelRequest struct {
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

func (s DescribeChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelRequest) GoString() string {
	return s.String()
}

func (s *DescribeChannelRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeChannelRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeChannelRequest) SetAppId(v string) *DescribeChannelRequest {
	s.AppId = &v
	return s
}

func (s *DescribeChannelRequest) SetChannelId(v string) *DescribeChannelRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeChannelRequest) Validate() error {
	return dara.Validate(s)
}
