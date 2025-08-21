// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePlayingListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *CreatePlayingListShrinkRequest
	GetDeviceInfoShrink() *string
	SetOpenCreatePlayingListRequestShrink(v string) *CreatePlayingListShrinkRequest
	GetOpenCreatePlayingListRequestShrink() *string
	SetUserInfoShrink(v string) *CreatePlayingListShrinkRequest
	GetUserInfoShrink() *string
}

type CreatePlayingListShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	OpenCreatePlayingListRequestShrink *string `json:"OpenCreatePlayingListRequest,omitempty" xml:"OpenCreatePlayingListRequest,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s CreatePlayingListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePlayingListShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePlayingListShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *CreatePlayingListShrinkRequest) GetOpenCreatePlayingListRequestShrink() *string {
	return s.OpenCreatePlayingListRequestShrink
}

func (s *CreatePlayingListShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *CreatePlayingListShrinkRequest) SetDeviceInfoShrink(v string) *CreatePlayingListShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *CreatePlayingListShrinkRequest) SetOpenCreatePlayingListRequestShrink(v string) *CreatePlayingListShrinkRequest {
	s.OpenCreatePlayingListRequestShrink = &v
	return s
}

func (s *CreatePlayingListShrinkRequest) SetUserInfoShrink(v string) *CreatePlayingListShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *CreatePlayingListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
