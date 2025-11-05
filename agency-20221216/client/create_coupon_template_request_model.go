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
	LimitPerPerson *string   `json:"LimitPerPerson,omitempty" xml:"LimitPerPerson,omitempty"`
	ProductType    []*string `json:"ProductType,omitempty" xml:"ProductType,omitempty" type:"Repeated"`
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
