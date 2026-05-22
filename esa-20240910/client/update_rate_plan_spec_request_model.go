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
	AutoPay        *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	ChargeType     *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OrderType      *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	TargetPlanCode *string `json:"TargetPlanCode,omitempty" xml:"TargetPlanCode,omitempty"`
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
