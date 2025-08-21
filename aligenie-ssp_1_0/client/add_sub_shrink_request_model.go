// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSubShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddSubscriptionInfoRequestShrink(v string) *AddSubShrinkRequest
	GetAddSubscriptionInfoRequestShrink() *string
	SetDeviceInfoShrink(v string) *AddSubShrinkRequest
	GetDeviceInfoShrink() *string
	SetUserInfoShrink(v string) *AddSubShrinkRequest
	GetUserInfoShrink() *string
}

type AddSubShrinkRequest struct {
	AddSubscriptionInfoRequestShrink *string `json:"AddSubscriptionInfoRequest,omitempty" xml:"AddSubscriptionInfoRequest,omitempty"`
	DeviceInfoShrink                 *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	UserInfoShrink                   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s AddSubShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddSubShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddSubShrinkRequest) GetAddSubscriptionInfoRequestShrink() *string {
	return s.AddSubscriptionInfoRequestShrink
}

func (s *AddSubShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *AddSubShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *AddSubShrinkRequest) SetAddSubscriptionInfoRequestShrink(v string) *AddSubShrinkRequest {
	s.AddSubscriptionInfoRequestShrink = &v
	return s
}

func (s *AddSubShrinkRequest) SetDeviceInfoShrink(v string) *AddSubShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *AddSubShrinkRequest) SetUserInfoShrink(v string) *AddSubShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *AddSubShrinkRequest) Validate() error {
	return dara.Validate(s)
}
