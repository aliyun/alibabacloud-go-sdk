// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMusicShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *ListMusicShrinkRequest
	GetDeviceInfoShrink() *string
	SetPayloadShrink(v string) *ListMusicShrinkRequest
	GetPayloadShrink() *string
	SetUserInfoShrink(v string) *ListMusicShrinkRequest
	GetUserInfoShrink() *string
}

type ListMusicShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ListMusicShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMusicShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListMusicShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *ListMusicShrinkRequest) GetPayloadShrink() *string {
	return s.PayloadShrink
}

func (s *ListMusicShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *ListMusicShrinkRequest) SetDeviceInfoShrink(v string) *ListMusicShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *ListMusicShrinkRequest) SetPayloadShrink(v string) *ListMusicShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *ListMusicShrinkRequest) SetUserInfoShrink(v string) *ListMusicShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *ListMusicShrinkRequest) Validate() error {
	return dara.Validate(s)
}
