// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmNotificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNotificationRequestId(v string) *ConfirmNotificationRequest
	GetNotificationRequestId() *string
}

type ConfirmNotificationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// deab3673236843a3b378c7d8d25c5f01
	NotificationRequestId *string `json:"NotificationRequestId,omitempty" xml:"NotificationRequestId,omitempty"`
}

func (s ConfirmNotificationRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfirmNotificationRequest) GoString() string {
	return s.String()
}

func (s *ConfirmNotificationRequest) GetNotificationRequestId() *string {
	return s.NotificationRequestId
}

func (s *ConfirmNotificationRequest) SetNotificationRequestId(v string) *ConfirmNotificationRequest {
	s.NotificationRequestId = &v
	return s
}

func (s *ConfirmNotificationRequest) Validate() error {
	return dara.Validate(s)
}
