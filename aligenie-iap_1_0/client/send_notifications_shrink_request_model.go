// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendNotificationsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *SendNotificationsShrinkRequest
	GetDeviceInfoShrink() *string
	SetNotificationUnicastRequestShrink(v string) *SendNotificationsShrinkRequest
	GetNotificationUnicastRequestShrink() *string
	SetTenantInfoShrink(v string) *SendNotificationsShrinkRequest
	GetTenantInfoShrink() *string
	SetUserInfoShrink(v string) *SendNotificationsShrinkRequest
	GetUserInfoShrink() *string
}

type SendNotificationsShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	NotificationUnicastRequestShrink *string `json:"NotificationUnicastRequest,omitempty" xml:"NotificationUnicastRequest,omitempty"`
	TenantInfoShrink                 *string `json:"TenantInfo,omitempty" xml:"TenantInfo,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s SendNotificationsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SendNotificationsShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendNotificationsShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *SendNotificationsShrinkRequest) GetNotificationUnicastRequestShrink() *string {
	return s.NotificationUnicastRequestShrink
}

func (s *SendNotificationsShrinkRequest) GetTenantInfoShrink() *string {
	return s.TenantInfoShrink
}

func (s *SendNotificationsShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *SendNotificationsShrinkRequest) SetDeviceInfoShrink(v string) *SendNotificationsShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *SendNotificationsShrinkRequest) SetNotificationUnicastRequestShrink(v string) *SendNotificationsShrinkRequest {
	s.NotificationUnicastRequestShrink = &v
	return s
}

func (s *SendNotificationsShrinkRequest) SetTenantInfoShrink(v string) *SendNotificationsShrinkRequest {
	s.TenantInfoShrink = &v
	return s
}

func (s *SendNotificationsShrinkRequest) SetUserInfoShrink(v string) *SendNotificationsShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *SendNotificationsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
