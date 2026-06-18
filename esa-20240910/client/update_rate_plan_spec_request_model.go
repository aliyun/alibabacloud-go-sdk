// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRatePlanSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *UpdateRatePlanSpecRequest
	GetAutoPay() *bool
	SetChargeType(v string) *UpdateRatePlanSpecRequest
	GetChargeType() *string
	SetInstanceId(v string) *UpdateRatePlanSpecRequest
	GetInstanceId() *string
	SetOrderType(v string) *UpdateRatePlanSpecRequest
	GetOrderType() *string
	SetTargetPlanCode(v string) *UpdateRatePlanSpecRequest
	GetTargetPlanCode() *string
	SetTargetPlanName(v string) *UpdateRatePlanSpecRequest
	GetTargetPlanName() *string
}

type UpdateRatePlanSpecRequest struct {
	// Specifies whether to enable automatic payment.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The billing type. Valid values:
	//
	// - PREPAY: Subscription.
	//
	// - POSTPAY: Pay-as-you-go.
	//
	// example:
	//
	// PREPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// xcdn-91fknmb80f0g
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of specification change. Valid values:
	//
	// - UPGRADE: Upgrade.
	//
	// example:
	//
	// UPGRADE
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The code of the target plan.
	//
	// Alibaba Cloud China Website (www.aliyun.com)
	//
	// - Free Edition: entranceplan
	//
	// - Basic Edition: basicplan
	//
	// - Standard Edition: standardplan
	//
	// - Premium Edition: advancedplan
	//
	// Alibaba Cloud International Website (www.alibabacloud.com)
	//
	// - Entrance: entranceplan
	//
	// - Pro: standardplan
	//
	// - Premium: advancedpla.
	//
	// example:
	//
	// entranceplan
	TargetPlanCode *string `json:"TargetPlanCode,omitempty" xml:"TargetPlanCode,omitempty"`
	// The name of the target plan.
	//
	// Alibaba Cloud China Website (www.aliyun.com)
	//
	// - Free Edition: entranceplan
	//
	// - Basic Edition: basic
	//
	// - Standard Edition: medium
	//
	// - Premium Edition: high
	//
	// Alibaba Cloud International Website (www.alibabacloud.com)
	//
	// - Entrance: entranceplan_intl
	//
	// - Pro: basicplan_intl
	//
	// - Premium: vipplan_intl.
	//
	// example:
	//
	// basic
	TargetPlanName *string `json:"TargetPlanName,omitempty" xml:"TargetPlanName,omitempty"`
}

func (s UpdateRatePlanSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRatePlanSpecRequest) GoString() string {
	return s.String()
}

func (s *UpdateRatePlanSpecRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *UpdateRatePlanSpecRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *UpdateRatePlanSpecRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateRatePlanSpecRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *UpdateRatePlanSpecRequest) GetTargetPlanCode() *string {
	return s.TargetPlanCode
}

func (s *UpdateRatePlanSpecRequest) GetTargetPlanName() *string {
	return s.TargetPlanName
}

func (s *UpdateRatePlanSpecRequest) SetAutoPay(v bool) *UpdateRatePlanSpecRequest {
	s.AutoPay = &v
	return s
}

func (s *UpdateRatePlanSpecRequest) SetChargeType(v string) *UpdateRatePlanSpecRequest {
	s.ChargeType = &v
	return s
}

func (s *UpdateRatePlanSpecRequest) SetInstanceId(v string) *UpdateRatePlanSpecRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateRatePlanSpecRequest) SetOrderType(v string) *UpdateRatePlanSpecRequest {
	s.OrderType = &v
	return s
}

func (s *UpdateRatePlanSpecRequest) SetTargetPlanCode(v string) *UpdateRatePlanSpecRequest {
	s.TargetPlanCode = &v
	return s
}

func (s *UpdateRatePlanSpecRequest) SetTargetPlanName(v string) *UpdateRatePlanSpecRequest {
	s.TargetPlanName = &v
	return s
}

func (s *UpdateRatePlanSpecRequest) Validate() error {
	return dara.Validate(s)
}
