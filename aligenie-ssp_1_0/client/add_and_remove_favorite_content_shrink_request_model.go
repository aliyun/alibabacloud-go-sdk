// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAndRemoveFavoriteContentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *AddAndRemoveFavoriteContentShrinkRequest
	GetDeviceInfoShrink() *string
	SetOpenAddAndRemoveFavoriteContentRequestShrink(v string) *AddAndRemoveFavoriteContentShrinkRequest
	GetOpenAddAndRemoveFavoriteContentRequestShrink() *string
	SetUserInfoShrink(v string) *AddAndRemoveFavoriteContentShrinkRequest
	GetUserInfoShrink() *string
}

type AddAndRemoveFavoriteContentShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	OpenAddAndRemoveFavoriteContentRequestShrink *string `json:"OpenAddAndRemoveFavoriteContentRequest,omitempty" xml:"OpenAddAndRemoveFavoriteContentRequest,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s AddAndRemoveFavoriteContentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAndRemoveFavoriteContentShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddAndRemoveFavoriteContentShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *AddAndRemoveFavoriteContentShrinkRequest) GetOpenAddAndRemoveFavoriteContentRequestShrink() *string {
	return s.OpenAddAndRemoveFavoriteContentRequestShrink
}

func (s *AddAndRemoveFavoriteContentShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *AddAndRemoveFavoriteContentShrinkRequest) SetDeviceInfoShrink(v string) *AddAndRemoveFavoriteContentShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *AddAndRemoveFavoriteContentShrinkRequest) SetOpenAddAndRemoveFavoriteContentRequestShrink(v string) *AddAndRemoveFavoriteContentShrinkRequest {
	s.OpenAddAndRemoveFavoriteContentRequestShrink = &v
	return s
}

func (s *AddAndRemoveFavoriteContentShrinkRequest) SetUserInfoShrink(v string) *AddAndRemoveFavoriteContentShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *AddAndRemoveFavoriteContentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
