// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRoomCheckOutShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *RoomCheckOutShrinkRequest
	GetDeviceInfoShrink() *string
	SetUserInfoShrink(v string) *RoomCheckOutShrinkRequest
	GetUserInfoShrink() *string
}

type RoomCheckOutShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	UserInfoShrink   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s RoomCheckOutShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RoomCheckOutShrinkRequest) GoString() string {
	return s.String()
}

func (s *RoomCheckOutShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *RoomCheckOutShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *RoomCheckOutShrinkRequest) SetDeviceInfoShrink(v string) *RoomCheckOutShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *RoomCheckOutShrinkRequest) SetUserInfoShrink(v string) *RoomCheckOutShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *RoomCheckOutShrinkRequest) Validate() error {
	return dara.Validate(s)
}
