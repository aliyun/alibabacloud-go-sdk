// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhoneNumberShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *GetPhoneNumberShrinkRequest
	GetDeviceInfoShrink() *string
	SetUserInfoShrink(v string) *GetPhoneNumberShrinkRequest
	GetUserInfoShrink() *string
}

type GetPhoneNumberShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetPhoneNumberShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPhoneNumberShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *GetPhoneNumberShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *GetPhoneNumberShrinkRequest) SetDeviceInfoShrink(v string) *GetPhoneNumberShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *GetPhoneNumberShrinkRequest) SetUserInfoShrink(v string) *GetPhoneNumberShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *GetPhoneNumberShrinkRequest) Validate() error {
	return dara.Validate(s)
}
