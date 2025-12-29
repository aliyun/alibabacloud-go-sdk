// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelContactByGenieDeviceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *GetHotelContactByGenieDeviceShrinkRequest
	GetDeviceInfoShrink() *string
	SetUserInfoShrink(v string) *GetHotelContactByGenieDeviceShrinkRequest
	GetUserInfoShrink() *string
}

type GetHotelContactByGenieDeviceShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	UserInfoShrink   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetHotelContactByGenieDeviceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotelContactByGenieDeviceShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetHotelContactByGenieDeviceShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *GetHotelContactByGenieDeviceShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *GetHotelContactByGenieDeviceShrinkRequest) SetDeviceInfoShrink(v string) *GetHotelContactByGenieDeviceShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *GetHotelContactByGenieDeviceShrinkRequest) SetUserInfoShrink(v string) *GetHotelContactByGenieDeviceShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *GetHotelContactByGenieDeviceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
