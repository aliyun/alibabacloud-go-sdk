// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCreditPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *CreateCreditPackageRequest
	GetAutoPay() *bool
	SetChannelCookie(v string) *CreateCreditPackageRequest
	GetChannelCookie() *string
	SetCreditAmount(v string) *CreateCreditPackageRequest
	GetCreditAmount() *string
	SetPackageAmount(v string) *CreateCreditPackageRequest
	GetPackageAmount() *string
	SetPeriod(v int32) *CreateCreditPackageRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateCreditPackageRequest
	GetPeriodUnit() *string
	SetPromotionId(v string) *CreateCreditPackageRequest
	GetPromotionId() *string
}

type CreateCreditPackageRequest struct {
	// Specifies whether to enable automatic payment. Valid values:
	//
	// - **true**: enables automatic payment. Make sure that your account balance is sufficient.
	//
	// - **false*	- (default): generates an order without charging your account.
	//
	//
	//
	//
	// > If your payment method has an insufficient balance, set this parameter to false. An unpaid order is generated, and you can log on to the Elastic Cloud Phone console to complete the payment.
	//
	// example:
	//
	// false
	AutoPay       *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	ChannelCookie *string `json:"ChannelCookie,omitempty" xml:"ChannelCookie,omitempty"`
	// The number of credits.
	//
	// example:
	//
	// 1000
	CreditAmount  *string `json:"CreditAmount,omitempty" xml:"CreditAmount,omitempty"`
	PackageAmount *string `json:"PackageAmount,omitempty" xml:"PackageAmount,omitempty"`
	// The duration for which you want to purchase the resource. The unit is specified by PeriodUnit.
	//
	// example:
	//
	// 6
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the duration for which you want to purchase the resource.
	//
	// Valid values:
	//
	// - **Month**: month.
	//
	// - **Year**: year.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The ID of the promotional campaign.
	//
	// example:
	//
	// 50003308011****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
}

func (s CreateCreditPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCreditPackageRequest) GoString() string {
	return s.String()
}

func (s *CreateCreditPackageRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateCreditPackageRequest) GetChannelCookie() *string {
	return s.ChannelCookie
}

func (s *CreateCreditPackageRequest) GetCreditAmount() *string {
	return s.CreditAmount
}

func (s *CreateCreditPackageRequest) GetPackageAmount() *string {
	return s.PackageAmount
}

func (s *CreateCreditPackageRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateCreditPackageRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateCreditPackageRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *CreateCreditPackageRequest) SetAutoPay(v bool) *CreateCreditPackageRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateCreditPackageRequest) SetChannelCookie(v string) *CreateCreditPackageRequest {
	s.ChannelCookie = &v
	return s
}

func (s *CreateCreditPackageRequest) SetCreditAmount(v string) *CreateCreditPackageRequest {
	s.CreditAmount = &v
	return s
}

func (s *CreateCreditPackageRequest) SetPackageAmount(v string) *CreateCreditPackageRequest {
	s.PackageAmount = &v
	return s
}

func (s *CreateCreditPackageRequest) SetPeriod(v int32) *CreateCreditPackageRequest {
	s.Period = &v
	return s
}

func (s *CreateCreditPackageRequest) SetPeriodUnit(v string) *CreateCreditPackageRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateCreditPackageRequest) SetPromotionId(v string) *CreateCreditPackageRequest {
	s.PromotionId = &v
	return s
}

func (s *CreateCreditPackageRequest) Validate() error {
	return dara.Validate(s)
}
