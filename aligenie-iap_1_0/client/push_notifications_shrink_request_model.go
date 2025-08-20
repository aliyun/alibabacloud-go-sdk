// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushNotificationsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNotificationUnicastRequestShrink(v string) *PushNotificationsShrinkRequest
	GetNotificationUnicastRequestShrink() *string
	SetTenantInfoShrink(v string) *PushNotificationsShrinkRequest
	GetTenantInfoShrink() *string
}

type PushNotificationsShrinkRequest struct {
	// This parameter is required.
	NotificationUnicastRequestShrink *string `json:"NotificationUnicastRequest,omitempty" xml:"NotificationUnicastRequest,omitempty"`
	TenantInfoShrink                 *string `json:"TenantInfo,omitempty" xml:"TenantInfo,omitempty"`
}

func (s PushNotificationsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PushNotificationsShrinkRequest) GoString() string {
	return s.String()
}

func (s *PushNotificationsShrinkRequest) GetNotificationUnicastRequestShrink() *string {
	return s.NotificationUnicastRequestShrink
}

func (s *PushNotificationsShrinkRequest) GetTenantInfoShrink() *string {
	return s.TenantInfoShrink
}

func (s *PushNotificationsShrinkRequest) SetNotificationUnicastRequestShrink(v string) *PushNotificationsShrinkRequest {
	s.NotificationUnicastRequestShrink = &v
	return s
}

func (s *PushNotificationsShrinkRequest) SetTenantInfoShrink(v string) *PushNotificationsShrinkRequest {
	s.TenantInfoShrink = &v
	return s
}

func (s *PushNotificationsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
