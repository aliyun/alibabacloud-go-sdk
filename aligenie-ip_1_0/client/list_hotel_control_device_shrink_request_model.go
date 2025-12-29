// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelControlDeviceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserInfoShrink(v string) *ListHotelControlDeviceShrinkRequest
	GetUserInfoShrink() *string
}

type ListHotelControlDeviceShrinkRequest struct {
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ListHotelControlDeviceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHotelControlDeviceShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListHotelControlDeviceShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *ListHotelControlDeviceShrinkRequest) SetUserInfoShrink(v string) *ListHotelControlDeviceShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *ListHotelControlDeviceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
