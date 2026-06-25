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
	SetCreditAmount(v string) *CreateCreditPackageRequest
	GetCreditAmount() *string
	SetPeriod(v int32) *CreateCreditPackageRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateCreditPackageRequest
	GetPeriodUnit() *string
	SetPromotionId(v string) *CreateCreditPackageRequest
	GetPromotionId() *string
}

type CreateCreditPackageRequest struct {
	// Whether to enable auto-payment. Valid values:
	//
	// - **true**: Enables auto-payment. Make sure that your account has a sufficient balance.
	//
	// - **false*	- (Default): Creates an unpaid order.
	//
	// > If your account has an insufficient balance, you can set this parameter to false. This generates an unpaid order. You can then pay for the order in the Wuying Cloud Phone management console.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The number of credits.
	//
	// example:
	//
	// 1000
	CreditAmount *string `json:"CreditAmount,omitempty" xml:"CreditAmount,omitempty"`
	// The subscription duration. The PeriodUnit parameter specifies the unit for the duration.
	//
	// example:
	//
	// 6
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration.
	//
	// Valid values:
	//
	// - **Month**: The period is measured in months.
	//
	// - **Year**: The period is measured in years.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The promotion ID.
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

func (s *CreateCreditPackageRequest) GetCreditAmount() *string {
	return s.CreditAmount
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

func (s *CreateCreditPackageRequest) SetCreditAmount(v string) *CreateCreditPackageRequest {
	s.CreditAmount = &v
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
