// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushDeviceNotificationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantInfoShrink(v string) *PushDeviceNotificationShrinkRequest
	GetTenantInfoShrink() *string
	SetBodyShrink(v string) *PushDeviceNotificationShrinkRequest
	GetBodyShrink() *string
}

type PushDeviceNotificationShrinkRequest struct {
	TenantInfoShrink *string `json:"TenantInfo,omitempty" xml:"TenantInfo,omitempty"`
	BodyShrink       *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushDeviceNotificationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PushDeviceNotificationShrinkRequest) GoString() string {
	return s.String()
}

func (s *PushDeviceNotificationShrinkRequest) GetTenantInfoShrink() *string {
	return s.TenantInfoShrink
}

func (s *PushDeviceNotificationShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *PushDeviceNotificationShrinkRequest) SetTenantInfoShrink(v string) *PushDeviceNotificationShrinkRequest {
	s.TenantInfoShrink = &v
	return s
}

func (s *PushDeviceNotificationShrinkRequest) SetBodyShrink(v string) *PushDeviceNotificationShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *PushDeviceNotificationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
