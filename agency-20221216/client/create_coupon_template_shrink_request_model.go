// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCouponTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreateCouponTemplateShrinkRequest
	GetAcceptLanguage() *string
	SetApplicableProducts(v string) *CreateCouponTemplateShrinkRequest
	GetApplicableProducts() *string
	SetCostBearer(v string) *CreateCouponTemplateShrinkRequest
	GetCostBearer() *string
	SetCouponDescription(v string) *CreateCouponTemplateShrinkRequest
	GetCouponDescription() *string
	SetExpireddate(v string) *CreateCouponTemplateShrinkRequest
	GetExpireddate() *string
	SetLimitPerPerson(v string) *CreateCouponTemplateShrinkRequest
	GetLimitPerPerson() *string
	SetProductTypeShrink(v string) *CreateCouponTemplateShrinkRequest
	GetProductTypeShrink() *string
	SetPurchaseType(v string) *CreateCouponTemplateShrinkRequest
	GetPurchaseType() *string
	SetReasonForApplication(v string) *CreateCouponTemplateShrinkRequest
	GetReasonForApplication() *string
	SetTemplateName(v string) *CreateCouponTemplateShrinkRequest
	GetTemplateName() *string
	SetVailddate(v string) *CreateCouponTemplateShrinkRequest
	GetVailddate() *string
	SetVaildperioddays(v string) *CreateCouponTemplateShrinkRequest
	GetVaildperioddays() *string
	SetValidUntil(v string) *CreateCouponTemplateShrinkRequest
	GetValidUntil() *string
	SetValue(v string) *CreateCouponTemplateShrinkRequest
	GetValue() *string
}

type CreateCouponTemplateShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// All Products
	ApplicableProducts *string `json:"ApplicableProducts,omitempty" xml:"ApplicableProducts,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Partner
	CostBearer        *string `json:"CostBearer,omitempty" xml:"CostBearer,omitempty"`
	CouponDescription *string `json:"CouponDescription,omitempty" xml:"CouponDescription,omitempty"`
	// example:
	//
	// 2024-08-26
	Expireddate *string `json:"Expireddate,omitempty" xml:"Expireddate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Unlimited
	LimitPerPerson    *string `json:"LimitPerPerson,omitempty" xml:"LimitPerPerson,omitempty"`
	ProductTypeShrink *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// ALL
	PurchaseType *string `json:"PurchaseType,omitempty" xml:"PurchaseType,omitempty"`
	// This parameter is required.
	ReasonForApplication *string `json:"ReasonForApplication,omitempty" xml:"ReasonForApplication,omitempty"`
	// This parameter is required.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// 2024-08-26
	Vailddate *string `json:"Vailddate,omitempty" xml:"Vailddate,omitempty"`
	// example:
	//
	// 1
	Vaildperioddays *string `json:"Vaildperioddays,omitempty" xml:"Vaildperioddays,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Validity Duration
	ValidUntil *string `json:"ValidUntil,omitempty" xml:"ValidUntil,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCouponTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCouponTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCouponTemplateShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreateCouponTemplateShrinkRequest) GetApplicableProducts() *string {
	return s.ApplicableProducts
}

func (s *CreateCouponTemplateShrinkRequest) GetCostBearer() *string {
	return s.CostBearer
}

func (s *CreateCouponTemplateShrinkRequest) GetCouponDescription() *string {
	return s.CouponDescription
}

func (s *CreateCouponTemplateShrinkRequest) GetExpireddate() *string {
	return s.Expireddate
}

func (s *CreateCouponTemplateShrinkRequest) GetLimitPerPerson() *string {
	return s.LimitPerPerson
}

func (s *CreateCouponTemplateShrinkRequest) GetProductTypeShrink() *string {
	return s.ProductTypeShrink
}

func (s *CreateCouponTemplateShrinkRequest) GetPurchaseType() *string {
	return s.PurchaseType
}

func (s *CreateCouponTemplateShrinkRequest) GetReasonForApplication() *string {
	return s.ReasonForApplication
}

func (s *CreateCouponTemplateShrinkRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateCouponTemplateShrinkRequest) GetVailddate() *string {
	return s.Vailddate
}

func (s *CreateCouponTemplateShrinkRequest) GetVaildperioddays() *string {
	return s.Vaildperioddays
}

func (s *CreateCouponTemplateShrinkRequest) GetValidUntil() *string {
	return s.ValidUntil
}

func (s *CreateCouponTemplateShrinkRequest) GetValue() *string {
	return s.Value
}

func (s *CreateCouponTemplateShrinkRequest) SetAcceptLanguage(v string) *CreateCouponTemplateShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateCouponTemplateShrinkRequest) SetApplicableProducts(v string) *CreateCouponTemplateShrinkRequest {
	s.ApplicableProducts = &v
	return s
}

func (s *CreateCouponTemplateShrinkRequest) SetCostBearer(v string) *CreateCouponTemplateShrinkRequest {
	s.CostBearer = &v
	return s
}

func (s *CreateCouponTemplateShrinkRequest) SetCouponDescription(v string) *CreateCouponTemplateShrinkRequest {
	s.CouponDescription = &v
	return s
}

func (s *CreateCouponTemplateShrinkRequest) SetExpireddate(v string) *CreateCouponTemplateShrinkRequest {
	s.Expireddate = &v
	return s
}

func (s *CreateCouponTemplateShrinkRequest) SetLimitPerPerson(v string) *CreateCouponTemplateShrinkRequest {
	s.LimitPerPerson = &v
	return s
}

func (s *CreateCouponTemplateShrinkRequest) SetProductTypeShrink(v string) *CreateCouponTemplateShrinkRequest {
	s.ProductTypeShrink = &v
	return s
}

func (s *CreateCouponTemplateShrinkRequest) SetPurchaseType(v string) *CreateCouponTemplateShrinkRequest {
	s.PurchaseType = &v
	return s
}

func (s *CreateCouponTemplateShrinkRequest) SetReasonForApplication(v string) *CreateCouponTemplateShrinkRequest {
	s.ReasonForApplication = &v
	return s
}

func (s *CreateCouponTemplateShrinkRequest) SetTemplateName(v string) *CreateCouponTemplateShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateCouponTemplateShrinkRequest) SetVailddate(v string) *CreateCouponTemplateShrinkRequest {
	s.Vailddate = &v
	return s
}

func (s *CreateCouponTemplateShrinkRequest) SetVaildperioddays(v string) *CreateCouponTemplateShrinkRequest {
	s.Vaildperioddays = &v
	return s
}

func (s *CreateCouponTemplateShrinkRequest) SetValidUntil(v string) *CreateCouponTemplateShrinkRequest {
	s.ValidUntil = &v
	return s
}

func (s *CreateCouponTemplateShrinkRequest) SetValue(v string) *CreateCouponTemplateShrinkRequest {
	s.Value = &v
	return s
}

func (s *CreateCouponTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
