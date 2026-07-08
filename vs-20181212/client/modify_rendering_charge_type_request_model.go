// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRenderingChargeTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *ModifyRenderingChargeTypeRequest
	GetAutoRenew() *bool
	SetInstanceBillingCycle(v string) *ModifyRenderingChargeTypeRequest
	GetInstanceBillingCycle() *string
	SetInstanceChargeType(v string) *ModifyRenderingChargeTypeRequest
	GetInstanceChargeType() *string
	SetPeriod(v string) *ModifyRenderingChargeTypeRequest
	GetPeriod() *string
	SetRenderingInstanceId(v string) *ModifyRenderingChargeTypeRequest
	GetRenderingInstanceId() *string
}

type ModifyRenderingChargeTypeRequest struct {
	// > This value is valid only when `InstanceChargeType` is `PrePaid` (subscription).
	//
	// Enable or disable auto-renewal. Valid values:
	//
	// - **true**: Enable.
	//
	// - **false**: Disable.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// > This value is valid only when `InstanceChargeType` is `PostPaid` (pay-as-you-go).
	//
	// Billing type. Valid values:
	//
	// - Hour: Hourly.
	//
	// example:
	//
	// Hour
	InstanceBillingCycle *string `json:"InstanceBillingCycle,omitempty" xml:"InstanceBillingCycle,omitempty"`
	// The target billing method for the instance. Valid values:
	//
	// - PrePaid (default): Subscription.
	//
	// - PostPaid: Pay-as-you-go.
	//
	// This parameter is required.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// > This value is valid only when `InstanceChargeType` is `PrePaid` (subscription).
	//
	// The duration for subscription. Valid values (Note: If you select 12, it converts to one year; other values are in months):
	//
	// - 1 (default)
	//
	// - 2
	//
	// - 3
	//
	// - 4
	//
	// - 5
	//
	// - 6
	//
	// - 7
	//
	// - 8
	//
	// - 9
	//
	// - 12
	//
	// example:
	//
	// 1
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The ID of the Graphic Computing Service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s ModifyRenderingChargeTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRenderingChargeTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyRenderingChargeTypeRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *ModifyRenderingChargeTypeRequest) GetInstanceBillingCycle() *string {
	return s.InstanceBillingCycle
}

func (s *ModifyRenderingChargeTypeRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *ModifyRenderingChargeTypeRequest) GetPeriod() *string {
	return s.Period
}

func (s *ModifyRenderingChargeTypeRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *ModifyRenderingChargeTypeRequest) SetAutoRenew(v bool) *ModifyRenderingChargeTypeRequest {
	s.AutoRenew = &v
	return s
}

func (s *ModifyRenderingChargeTypeRequest) SetInstanceBillingCycle(v string) *ModifyRenderingChargeTypeRequest {
	s.InstanceBillingCycle = &v
	return s
}

func (s *ModifyRenderingChargeTypeRequest) SetInstanceChargeType(v string) *ModifyRenderingChargeTypeRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *ModifyRenderingChargeTypeRequest) SetPeriod(v string) *ModifyRenderingChargeTypeRequest {
	s.Period = &v
	return s
}

func (s *ModifyRenderingChargeTypeRequest) SetRenderingInstanceId(v string) *ModifyRenderingChargeTypeRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *ModifyRenderingChargeTypeRequest) Validate() error {
	return dara.Validate(s)
}
