// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIndexControlPlayingListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *IndexControlPlayingListShrinkRequest
	GetDeviceInfoShrink() *string
	SetOpenIndexControlRequestShrink(v string) *IndexControlPlayingListShrinkRequest
	GetOpenIndexControlRequestShrink() *string
	SetUserInfoShrink(v string) *IndexControlPlayingListShrinkRequest
	GetUserInfoShrink() *string
}

type IndexControlPlayingListShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	OpenIndexControlRequestShrink *string `json:"OpenIndexControlRequest,omitempty" xml:"OpenIndexControlRequest,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s IndexControlPlayingListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s IndexControlPlayingListShrinkRequest) GoString() string {
	return s.String()
}

func (s *IndexControlPlayingListShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *IndexControlPlayingListShrinkRequest) GetOpenIndexControlRequestShrink() *string {
	return s.OpenIndexControlRequestShrink
}

func (s *IndexControlPlayingListShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *IndexControlPlayingListShrinkRequest) SetDeviceInfoShrink(v string) *IndexControlPlayingListShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *IndexControlPlayingListShrinkRequest) SetOpenIndexControlRequestShrink(v string) *IndexControlPlayingListShrinkRequest {
	s.OpenIndexControlRequestShrink = &v
	return s
}

func (s *IndexControlPlayingListShrinkRequest) SetUserInfoShrink(v string) *IndexControlPlayingListShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *IndexControlPlayingListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
