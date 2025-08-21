// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceByUserIdAndChanelShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelInfoShrink(v string) *ListDeviceByUserIdAndChanelShrinkRequest
	GetChannelInfoShrink() *string
	SetUserInfoShrink(v string) *ListDeviceByUserIdAndChanelShrinkRequest
	GetUserInfoShrink() *string
}

type ListDeviceByUserIdAndChanelShrinkRequest struct {
	// This parameter is required.
	ChannelInfoShrink *string `json:"ChannelInfo,omitempty" xml:"ChannelInfo,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ListDeviceByUserIdAndChanelShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceByUserIdAndChanelShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceByUserIdAndChanelShrinkRequest) GetChannelInfoShrink() *string {
	return s.ChannelInfoShrink
}

func (s *ListDeviceByUserIdAndChanelShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *ListDeviceByUserIdAndChanelShrinkRequest) SetChannelInfoShrink(v string) *ListDeviceByUserIdAndChanelShrinkRequest {
	s.ChannelInfoShrink = &v
	return s
}

func (s *ListDeviceByUserIdAndChanelShrinkRequest) SetUserInfoShrink(v string) *ListDeviceByUserIdAndChanelShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *ListDeviceByUserIdAndChanelShrinkRequest) Validate() error {
	return dara.Validate(s)
}
