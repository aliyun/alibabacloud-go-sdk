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
	// The accepted language.
	//
	// Default value: None.
	//
	// Valid values:
	//
	// - zh-CN: Chinese.
	//
	// - en-US: English.
	//
	// - ja: Japanese.
	//
	// This parameter is required.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The applicable products.
	//
	// - All Products: All products.
	//
	// - Custom: Custom products.
	//
	// This parameter is required.
	//
	// example:
	//
	// All Products
	ApplicableProducts *string `json:"ApplicableProducts,omitempty" xml:"ApplicableProducts,omitempty"`
	// The party that bears the coupon cost.
	//
	// - Partner: Borne by the partner.
	//
	// - Alibaba Cloud: Borne by Alibaba Cloud.
	//
	// This parameter is required.
	//
	// example:
	//
	// Partner
	CostBearer *string `json:"CostBearer,omitempty" xml:"CostBearer,omitempty"`
	// The remark for the coupon.
	//
	// The field length is limited to 100 characters.
	//
	// example:
	//
	// Coupon remark
	CouponDescription *string `json:"CouponDescription,omitempty" xml:"CouponDescription,omitempty"`
	// The end time of the effective period.
	//
	// Format: yyyy-MM-dd.
	//
	// Example: 2024-08-26.
	//
	// This parameter is required when the validity period is set to Custom Timeframe.
	//
	// example:
	//
	// 2024-08-26
	Expireddate *string `json:"Expireddate,omitempty" xml:"Expireddate,omitempty"`
	// The maximum number of coupons that each person can claim.
	//
	// - Limited to 1: Limited to one coupon.
	//
	// - Unlimited: Unlimited.
	//
	// This parameter is required.
	//
	// example:
	//
	// Unlimited
	LimitPerPerson *string `json:"LimitPerPerson,omitempty" xml:"LimitPerPerson,omitempty"`
	// The list of product codes.
	ProductTypeShrink *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The order type.
	//
	// example:
	//
	// ALL
	PurchaseType *string `json:"PurchaseType,omitempty" xml:"PurchaseType,omitempty"`
	// The reason for the template application.
	//
	// The field length is limited to 100 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Reason for template application
	ReasonForApplication *string `json:"ReasonForApplication,omitempty" xml:"ReasonForApplication,omitempty"`
	// The template name.
	//
	// The field length is limited to 30 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Template name
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The start time of the effective period.
	//
	// Format: yyyy-MM-dd.
	//
	// Example: 2024-08-26.
	//
	// This parameter is required when the validity period is set to Custom Timeframe.
	//
	// example:
	//
	// 2024-08-26
	Vailddate *string `json:"Vailddate,omitempty" xml:"Vailddate,omitempty"`
	// The duration of the relative time in days.
	//
	// The value must be a positive integer.
	//
	// This parameter is required when the validity period is set to Validity Duration.
	//
	// example:
	//
	// 1
	Vaildperioddays *string `json:"Vaildperioddays,omitempty" xml:"Vaildperioddays,omitempty"`
	// The validity period of the coupon.
	//
	// - Validity Duration: From the date of issuance.
	//
	// - Custom Timeframe: Custom time period.
	//
	// This parameter is required.
	//
	// example:
	//
	// Validity Duration
	ValidUntil *string `json:"ValidUntil,omitempty" xml:"ValidUntil,omitempty"`
	// The denomination of the coupon.
	//
	// The value must be an integer.
	//
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
