// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlbumIsAddedShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlbumIdListShrink(v string) *ListAlbumIsAddedShrinkRequest
	GetAlbumIdListShrink() *string
	SetDeviceInfoShrink(v string) *ListAlbumIsAddedShrinkRequest
	GetDeviceInfoShrink() *string
	SetUserInfoShrink(v string) *ListAlbumIsAddedShrinkRequest
	GetUserInfoShrink() *string
}

type ListAlbumIsAddedShrinkRequest struct {
	AlbumIdListShrink *string `json:"AlbumIdList,omitempty" xml:"AlbumIdList,omitempty"`
	DeviceInfoShrink  *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	UserInfoShrink    *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ListAlbumIsAddedShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlbumIsAddedShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAlbumIsAddedShrinkRequest) GetAlbumIdListShrink() *string {
	return s.AlbumIdListShrink
}

func (s *ListAlbumIsAddedShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *ListAlbumIsAddedShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *ListAlbumIsAddedShrinkRequest) SetAlbumIdListShrink(v string) *ListAlbumIsAddedShrinkRequest {
	s.AlbumIdListShrink = &v
	return s
}

func (s *ListAlbumIsAddedShrinkRequest) SetDeviceInfoShrink(v string) *ListAlbumIsAddedShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *ListAlbumIsAddedShrinkRequest) SetUserInfoShrink(v string) *ListAlbumIsAddedShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *ListAlbumIsAddedShrinkRequest) Validate() error {
	return dara.Validate(s)
}
