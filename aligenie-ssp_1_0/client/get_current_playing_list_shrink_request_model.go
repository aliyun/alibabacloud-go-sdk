// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCurrentPlayingListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *GetCurrentPlayingListShrinkRequest
	GetDeviceInfoShrink() *string
	SetOpenQueryPlayListRequestShrink(v string) *GetCurrentPlayingListShrinkRequest
	GetOpenQueryPlayListRequestShrink() *string
	SetUserInfoShrink(v string) *GetCurrentPlayingListShrinkRequest
	GetUserInfoShrink() *string
}

type GetCurrentPlayingListShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	OpenQueryPlayListRequestShrink *string `json:"OpenQueryPlayListRequest,omitempty" xml:"OpenQueryPlayListRequest,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetCurrentPlayingListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCurrentPlayingListShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingListShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *GetCurrentPlayingListShrinkRequest) GetOpenQueryPlayListRequestShrink() *string {
	return s.OpenQueryPlayListRequestShrink
}

func (s *GetCurrentPlayingListShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *GetCurrentPlayingListShrinkRequest) SetDeviceInfoShrink(v string) *GetCurrentPlayingListShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *GetCurrentPlayingListShrinkRequest) SetOpenQueryPlayListRequestShrink(v string) *GetCurrentPlayingListShrinkRequest {
	s.OpenQueryPlayListRequestShrink = &v
	return s
}

func (s *GetCurrentPlayingListShrinkRequest) SetUserInfoShrink(v string) *GetCurrentPlayingListShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *GetCurrentPlayingListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
