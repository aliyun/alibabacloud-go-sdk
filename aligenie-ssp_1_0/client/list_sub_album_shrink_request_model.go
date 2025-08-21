// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubAlbumShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *ListSubAlbumShrinkRequest
	GetDeviceInfoShrink() *string
	SetQuerySubscriptionAlbumRequestShrink(v string) *ListSubAlbumShrinkRequest
	GetQuerySubscriptionAlbumRequestShrink() *string
	SetUserInfoShrink(v string) *ListSubAlbumShrinkRequest
	GetUserInfoShrink() *string
}

type ListSubAlbumShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// request
	QuerySubscriptionAlbumRequestShrink *string `json:"QuerySubscriptionAlbumRequest,omitempty" xml:"QuerySubscriptionAlbumRequest,omitempty"`
	UserInfoShrink                      *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ListSubAlbumShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSubAlbumShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListSubAlbumShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *ListSubAlbumShrinkRequest) GetQuerySubscriptionAlbumRequestShrink() *string {
	return s.QuerySubscriptionAlbumRequestShrink
}

func (s *ListSubAlbumShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *ListSubAlbumShrinkRequest) SetDeviceInfoShrink(v string) *ListSubAlbumShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *ListSubAlbumShrinkRequest) SetQuerySubscriptionAlbumRequestShrink(v string) *ListSubAlbumShrinkRequest {
	s.QuerySubscriptionAlbumRequestShrink = &v
	return s
}

func (s *ListSubAlbumShrinkRequest) SetUserInfoShrink(v string) *ListSubAlbumShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *ListSubAlbumShrinkRequest) Validate() error {
	return dara.Validate(s)
}
