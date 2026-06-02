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
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// example:
	//
	// 1000
	CreditAmount *string `json:"CreditAmount,omitempty" xml:"CreditAmount,omitempty"`
	// example:
	//
	// 6
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
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
