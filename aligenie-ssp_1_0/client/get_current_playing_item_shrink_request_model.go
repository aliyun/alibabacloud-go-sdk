// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCurrentPlayingItemShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *GetCurrentPlayingItemShrinkRequest
	GetDeviceInfoShrink() *string
	SetUserInfoShrink(v string) *GetCurrentPlayingItemShrinkRequest
	GetUserInfoShrink() *string
}

type GetCurrentPlayingItemShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetCurrentPlayingItemShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCurrentPlayingItemShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingItemShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *GetCurrentPlayingItemShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *GetCurrentPlayingItemShrinkRequest) SetDeviceInfoShrink(v string) *GetCurrentPlayingItemShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *GetCurrentPlayingItemShrinkRequest) SetUserInfoShrink(v string) *GetCurrentPlayingItemShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *GetCurrentPlayingItemShrinkRequest) Validate() error {
	return dara.Validate(s)
}
