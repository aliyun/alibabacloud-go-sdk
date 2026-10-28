// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCouponTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreateCouponTemplateRequest
	GetAcceptLanguage() *string
	SetApplicableProducts(v string) *CreateCouponTemplateRequest
	GetApplicableProducts() *string
	SetCostBearer(v string) *CreateCouponTemplateRequest
	GetCostBearer() *string
	SetCouponDescription(v string) *CreateCouponTemplateRequest
	GetCouponDescription() *string
	SetExpireddate(v string) *CreateCouponTemplateRequest
	GetExpireddate() *string
	SetLimitPerPerson(v string) *CreateCouponTemplateRequest
	GetLimitPerPerson() *string
	SetProductType(v []*string) *CreateCouponTemplateRequest
	GetProductType() []*string
	SetPurchaseType(v string) *CreateCouponTemplateRequest
	GetPurchaseType() *string
	SetReasonForApplication(v string) *CreateCouponTemplateRequest
	GetReasonForApplication() *string
	SetTemplateName(v string) *CreateCouponTemplateRequest
	GetTemplateName() *string
	SetVailddate(v string) *CreateCouponTemplateRequest
	GetVailddate() *string
	SetVaildperioddays(v string) *CreateCouponTemplateRequest
	GetVaildperioddays() *string
	SetValidUntil(v string) *CreateCouponTemplateRequest
	GetValidUntil() *string
	SetValue(v string) *CreateCouponTemplateRequest
	GetValue() *string
}

type CreateCouponTemplateRequest struct {
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
	ProductType []*string `json:"ProductType,omitempty" xml:"ProductType,omitempty" type:"Repeated"`
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

func (s CreateCouponTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCouponTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateCouponTemplateRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreateCouponTemplateRequest) GetApplicableProducts() *string {
	return s.ApplicableProducts
}

func (s *CreateCouponTemplateRequest) GetCostBearer() *string {
	return s.CostBearer
}

func (s *CreateCouponTemplateRequest) GetCouponDescription() *string {
	return s.CouponDescription
}

func (s *CreateCouponTemplateRequest) GetExpireddate() *string {
	return s.Expireddate
}

func (s *CreateCouponTemplateRequest) GetLimitPerPerson() *string {
	return s.LimitPerPerson
}

func (s *CreateCouponTemplateRequest) GetProductType() []*string {
	return s.ProductType
}

func (s *CreateCouponTemplateRequest) GetPurchaseType() *string {
	return s.PurchaseType
}

func (s *CreateCouponTemplateRequest) GetReasonForApplication() *string {
	return s.ReasonForApplication
}

func (s *CreateCouponTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateCouponTemplateRequest) GetVailddate() *string {
	return s.Vailddate
}

func (s *CreateCouponTemplateRequest) GetVaildperioddays() *string {
	return s.Vaildperioddays
}

func (s *CreateCouponTemplateRequest) GetValidUntil() *string {
	return s.ValidUntil
}

func (s *CreateCouponTemplateRequest) GetValue() *string {
	return s.Value
}

func (s *CreateCouponTemplateRequest) SetAcceptLanguage(v string) *CreateCouponTemplateRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetApplicableProducts(v string) *CreateCouponTemplateRequest {
	s.ApplicableProducts = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetCostBearer(v string) *CreateCouponTemplateRequest {
	s.CostBearer = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetCouponDescription(v string) *CreateCouponTemplateRequest {
	s.CouponDescription = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetExpireddate(v string) *CreateCouponTemplateRequest {
	s.Expireddate = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetLimitPerPerson(v string) *CreateCouponTemplateRequest {
	s.LimitPerPerson = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetProductType(v []*string) *CreateCouponTemplateRequest {
	s.ProductType = v
	return s
}

func (s *CreateCouponTemplateRequest) SetPurchaseType(v string) *CreateCouponTemplateRequest {
	s.PurchaseType = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetReasonForApplication(v string) *CreateCouponTemplateRequest {
	s.ReasonForApplication = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetTemplateName(v string) *CreateCouponTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetVailddate(v string) *CreateCouponTemplateRequest {
	s.Vailddate = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetVaildperioddays(v string) *CreateCouponTemplateRequest {
	s.Vaildperioddays = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetValidUntil(v string) *CreateCouponTemplateRequest {
	s.ValidUntil = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetValue(v string) *CreateCouponTemplateRequest {
	s.Value = &v
	return s
}

func (s *CreateCouponTemplateRequest) Validate() error {
	return dara.Validate(s)
}
