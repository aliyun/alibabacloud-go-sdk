// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurchaseRatePlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int32) *PurchaseRatePlanRequest
	GetAmount() *int32
	SetAutoPay(v bool) *PurchaseRatePlanRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *PurchaseRatePlanRequest
	GetAutoRenew() *bool
	SetChannel(v string) *PurchaseRatePlanRequest
	GetChannel() *string
	SetChargeType(v string) *PurchaseRatePlanRequest
	GetChargeType() *string
	SetCoverage(v string) *PurchaseRatePlanRequest
	GetCoverage() *string
	SetPeriod(v int32) *PurchaseRatePlanRequest
	GetPeriod() *int32
	SetPlanCode(v string) *PurchaseRatePlanRequest
	GetPlanCode() *string
	SetPlanName(v string) *PurchaseRatePlanRequest
	GetPlanName() *string
	SetSiteName(v string) *PurchaseRatePlanRequest
	GetSiteName() *string
	SetType(v string) *PurchaseRatePlanRequest
	GetType() *string
}

type PurchaseRatePlanRequest struct {
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// Specifies whether to enable auto payment.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Auto-renewal:
	//
	// - true: Enable auto-renewal.
	//
	// - false: Disable auto-renewal.
	//
	// example:
	//
	// true
	AutoRenew *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	Channel   *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// The billing method. Valid values:
	//
	// 	- PREPAY: subscription.
	//
	// 	- POSTPAY: pay-as-you-go.
	//
	// example:
	//
	// PREPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The service location. Valid values:
	//
	// 	- domestic: the Chinese mainland.
	//
	// 	- global: global.
	//
	// 	- overseas: outside the Chinese mainland.
	//
	// example:
	//
	// domestic
	Coverage *string `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	// Subscription period (in months).
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// Package code.
	//
	// example:
	//
	// entranceplan
	PlanCode *string `json:"PlanCode,omitempty" xml:"PlanCode,omitempty"`
	// Package name.
	//
	// example:
	//
	// basic
	PlanName *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	// Site name.
	//
	// example:
	//
	// test.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	// The DNS setup option for the website. Valid values:
	//
	// 	- NS
	//
	// 	- CNAME
	//
	// example:
	//
	// CNAME
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PurchaseRatePlanRequest) String() string {
	return dara.Prettify(s)
}

func (s PurchaseRatePlanRequest) GoString() string {
	return s.String()
}

func (s *PurchaseRatePlanRequest) GetAmount() *int32 {
	return s.Amount
}

func (s *PurchaseRatePlanRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *PurchaseRatePlanRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *PurchaseRatePlanRequest) GetChannel() *string {
	return s.Channel
}

func (s *PurchaseRatePlanRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *PurchaseRatePlanRequest) GetCoverage() *string {
	return s.Coverage
}

func (s *PurchaseRatePlanRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *PurchaseRatePlanRequest) GetPlanCode() *string {
	return s.PlanCode
}

func (s *PurchaseRatePlanRequest) GetPlanName() *string {
	return s.PlanName
}

func (s *PurchaseRatePlanRequest) GetSiteName() *string {
	return s.SiteName
}

func (s *PurchaseRatePlanRequest) GetType() *string {
	return s.Type
}

func (s *PurchaseRatePlanRequest) SetAmount(v int32) *PurchaseRatePlanRequest {
	s.Amount = &v
	return s
}

func (s *PurchaseRatePlanRequest) SetAutoPay(v bool) *PurchaseRatePlanRequest {
	s.AutoPay = &v
	return s
}

func (s *PurchaseRatePlanRequest) SetAutoRenew(v bool) *PurchaseRatePlanRequest {
	s.AutoRenew = &v
	return s
}

func (s *PurchaseRatePlanRequest) SetChannel(v string) *PurchaseRatePlanRequest {
	s.Channel = &v
	return s
}

func (s *PurchaseRatePlanRequest) SetChargeType(v string) *PurchaseRatePlanRequest {
	s.ChargeType = &v
	return s
}

func (s *PurchaseRatePlanRequest) SetCoverage(v string) *PurchaseRatePlanRequest {
	s.Coverage = &v
	return s
}

func (s *PurchaseRatePlanRequest) SetPeriod(v int32) *PurchaseRatePlanRequest {
	s.Period = &v
	return s
}

func (s *PurchaseRatePlanRequest) SetPlanCode(v string) *PurchaseRatePlanRequest {
	s.PlanCode = &v
	return s
}

func (s *PurchaseRatePlanRequest) SetPlanName(v string) *PurchaseRatePlanRequest {
	s.PlanName = &v
	return s
}

func (s *PurchaseRatePlanRequest) SetSiteName(v string) *PurchaseRatePlanRequest {
	s.SiteName = &v
	return s
}

func (s *PurchaseRatePlanRequest) SetType(v string) *PurchaseRatePlanRequest {
	s.Type = &v
	return s
}

func (s *PurchaseRatePlanRequest) Validate() error {
	return dara.Validate(s)
}
