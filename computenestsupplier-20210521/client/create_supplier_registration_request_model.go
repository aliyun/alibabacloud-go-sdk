// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSupplierRegistrationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactEmail(v string) *CreateSupplierRegistrationRequest
	GetContactEmail() *string
	SetContactNumber(v string) *CreateSupplierRegistrationRequest
	GetContactNumber() *string
	SetContactPerson(v string) *CreateSupplierRegistrationRequest
	GetContactPerson() *string
	SetContactPersonTitle(v string) *CreateSupplierRegistrationRequest
	GetContactPersonTitle() *string
	SetEnableResellerMode(v bool) *CreateSupplierRegistrationRequest
	GetEnableResellerMode() *bool
	SetProductAnnualRevenue(v string) *CreateSupplierRegistrationRequest
	GetProductAnnualRevenue() *string
	SetProductBusiness(v string) *CreateSupplierRegistrationRequest
	GetProductBusiness() *string
	SetProductDeliveryTypes(v []*string) *CreateSupplierRegistrationRequest
	GetProductDeliveryTypes() []*string
	SetProductPublishTime(v string) *CreateSupplierRegistrationRequest
	GetProductPublishTime() *string
	SetProductSellTypes(v []*string) *CreateSupplierRegistrationRequest
	GetProductSellTypes() []*string
	SetRegionId(v string) *CreateSupplierRegistrationRequest
	GetRegionId() *string
	SetResellBusinessDesc(v string) *CreateSupplierRegistrationRequest
	GetResellBusinessDesc() *string
	SetSuggestion(v string) *CreateSupplierRegistrationRequest
	GetSuggestion() *string
	SetSupplierDesc(v string) *CreateSupplierRegistrationRequest
	GetSupplierDesc() *string
	SetSupplierLogo(v string) *CreateSupplierRegistrationRequest
	GetSupplierLogo() *string
	SetSupplierName(v string) *CreateSupplierRegistrationRequest
	GetSupplierName() *string
	SetSupplierNameEn(v string) *CreateSupplierRegistrationRequest
	GetSupplierNameEn() *string
	SetSupplierUrl(v string) *CreateSupplierRegistrationRequest
	GetSupplierUrl() *string
}

type CreateSupplierRegistrationRequest struct {
	// Contact email
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxx@xxx.com
	ContactEmail *string `json:"ContactEmail,omitempty" xml:"ContactEmail,omitempty"`
	// Contact number
	//
	// This parameter is required.
	//
	// example:
	//
	// 186xxxxxxxxx
	ContactNumber *string `json:"ContactNumber,omitempty" xml:"ContactNumber,omitempty"`
	// Contact person
	//
	// This parameter is required.
	//
	// example:
	//
	// Mike
	ContactPerson *string `json:"ContactPerson,omitempty" xml:"ContactPerson,omitempty"`
	// Contact person tiltle
	//
	// This parameter is required.
	//
	// example:
	//
	// CTO
	ContactPersonTitle *string `json:"ContactPersonTitle,omitempty" xml:"ContactPersonTitle,omitempty"`
	// Whether to enable the resell mode
	//
	// example:
	//
	// true
	EnableResellerMode *bool `json:"EnableResellerMode,omitempty" xml:"EnableResellerMode,omitempty"`
	// Annual product revenue
	//
	// example:
	//
	// 1000
	ProductAnnualRevenue *string `json:"ProductAnnualRevenue,omitempty" xml:"ProductAnnualRevenue,omitempty"`
	// The business of product
	//
	// example:
	//
	// AI
	ProductBusiness *string `json:"ProductBusiness,omitempty" xml:"ProductBusiness,omitempty"`
	// Product delivery type
	//
	// This parameter is required.
	ProductDeliveryTypes []*string `json:"ProductDeliveryTypes,omitempty" xml:"ProductDeliveryTypes,omitempty" type:"Repeated"`
	// The publish time of product
	//
	// example:
	//
	// 2020.10.10
	ProductPublishTime *string `json:"ProductPublishTime,omitempty" xml:"ProductPublishTime,omitempty"`
	// Product sell type
	//
	// This parameter is required.
	ProductSellTypes []*string `json:"ProductSellTypes,omitempty" xml:"ProductSellTypes,omitempty" type:"Repeated"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The description of resell business.
	//
	// example:
	//
	// empty
	ResellBusinessDesc *string `json:"ResellBusinessDesc,omitempty" xml:"ResellBusinessDesc,omitempty"`
	// The demands of service providers.
	//
	// example:
	//
	// empty
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The description of service provider.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test supplier
	SupplierDesc *string `json:"SupplierDesc,omitempty" xml:"SupplierDesc,omitempty"`
	// The Logo of service provider.
	//
	// example:
	//
	// http://example.aliyundoc.com/cover/34DB-4F4C-9373-003AA060****.png
	SupplierLogo *string `json:"SupplierLogo,omitempty" xml:"SupplierLogo,omitempty"`
	// The name of the service provider.
	//
	// This parameter is required.
	//
	// example:
	//
	// Alibaba Cloud
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// The english name of the service provider.
	//
	// This parameter is required.
	//
	// example:
	//
	// Alibaba Cloud
	SupplierNameEn *string `json:"SupplierNameEn,omitempty" xml:"SupplierNameEn,omitempty"`
	// The URL of the service provider.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://www.xxx.xxx.cn
	SupplierUrl *string `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
}

func (s CreateSupplierRegistrationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSupplierRegistrationRequest) GoString() string {
	return s.String()
}

func (s *CreateSupplierRegistrationRequest) GetContactEmail() *string {
	return s.ContactEmail
}

func (s *CreateSupplierRegistrationRequest) GetContactNumber() *string {
	return s.ContactNumber
}

func (s *CreateSupplierRegistrationRequest) GetContactPerson() *string {
	return s.ContactPerson
}

func (s *CreateSupplierRegistrationRequest) GetContactPersonTitle() *string {
	return s.ContactPersonTitle
}

func (s *CreateSupplierRegistrationRequest) GetEnableResellerMode() *bool {
	return s.EnableResellerMode
}

func (s *CreateSupplierRegistrationRequest) GetProductAnnualRevenue() *string {
	return s.ProductAnnualRevenue
}

func (s *CreateSupplierRegistrationRequest) GetProductBusiness() *string {
	return s.ProductBusiness
}

func (s *CreateSupplierRegistrationRequest) GetProductDeliveryTypes() []*string {
	return s.ProductDeliveryTypes
}

func (s *CreateSupplierRegistrationRequest) GetProductPublishTime() *string {
	return s.ProductPublishTime
}

func (s *CreateSupplierRegistrationRequest) GetProductSellTypes() []*string {
	return s.ProductSellTypes
}

func (s *CreateSupplierRegistrationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSupplierRegistrationRequest) GetResellBusinessDesc() *string {
	return s.ResellBusinessDesc
}

func (s *CreateSupplierRegistrationRequest) GetSuggestion() *string {
	return s.Suggestion
}

func (s *CreateSupplierRegistrationRequest) GetSupplierDesc() *string {
	return s.SupplierDesc
}

func (s *CreateSupplierRegistrationRequest) GetSupplierLogo() *string {
	return s.SupplierLogo
}

func (s *CreateSupplierRegistrationRequest) GetSupplierName() *string {
	return s.SupplierName
}

func (s *CreateSupplierRegistrationRequest) GetSupplierNameEn() *string {
	return s.SupplierNameEn
}

func (s *CreateSupplierRegistrationRequest) GetSupplierUrl() *string {
	return s.SupplierUrl
}

func (s *CreateSupplierRegistrationRequest) SetContactEmail(v string) *CreateSupplierRegistrationRequest {
	s.ContactEmail = &v
	return s
}

func (s *CreateSupplierRegistrationRequest) SetContactNumber(v string) *CreateSupplierRegistrationRequest {
	s.ContactNumber = &v
	return s
}

func (s *CreateSupplierRegistrationRequest) SetContactPerson(v string) *CreateSupplierRegistrationRequest {
	s.ContactPerson = &v
	return s
}

func (s *CreateSupplierRegistrationRequest) SetContactPersonTitle(v string) *CreateSupplierRegistrationRequest {
	s.ContactPersonTitle = &v
	return s
}

func (s *CreateSupplierRegistrationRequest) SetEnableResellerMode(v bool) *CreateSupplierRegistrationRequest {
	s.EnableResellerMode = &v
	return s
}

func (s *CreateSupplierRegistrationRequest) SetProductAnnualRevenue(v string) *CreateSupplierRegistrationRequest {
	s.ProductAnnualRevenue = &v
	return s
}

func (s *CreateSupplierRegistrationRequest) SetProductBusiness(v string) *CreateSupplierRegistrationRequest {
	s.ProductBusiness = &v
	return s
}

func (s *CreateSupplierRegistrationRequest) SetProductDeliveryTypes(v []*string) *CreateSupplierRegistrationRequest {
	s.ProductDeliveryTypes = v
	return s
}

func (s *CreateSupplierRegistrationRequest) SetProductPublishTime(v string) *CreateSupplierRegistrationRequest {
	s.ProductPublishTime = &v
	return s
}

func (s *CreateSupplierRegistrationRequest) SetProductSellTypes(v []*string) *CreateSupplierRegistrationRequest {
	s.ProductSellTypes = v
	return s
}

func (s *CreateSupplierRegistrationRequest) SetRegionId(v string) *CreateSupplierRegistrationRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSupplierRegistrationRequest) SetResellBusinessDesc(v string) *CreateSupplierRegistrationRequest {
	s.ResellBusinessDesc = &v
	return s
}

func (s *CreateSupplierRegistrationRequest) SetSuggestion(v string) *CreateSupplierRegistrationRequest {
	s.Suggestion = &v
	return s
}

func (s *CreateSupplierRegistrationRequest) SetSupplierDesc(v string) *CreateSupplierRegistrationRequest {
	s.SupplierDesc = &v
	return s
}

func (s *CreateSupplierRegistrationRequest) SetSupplierLogo(v string) *CreateSupplierRegistrationRequest {
	s.SupplierLogo = &v
	return s
}

func (s *CreateSupplierRegistrationRequest) SetSupplierName(v string) *CreateSupplierRegistrationRequest {
	s.SupplierName = &v
	return s
}

func (s *CreateSupplierRegistrationRequest) SetSupplierNameEn(v string) *CreateSupplierRegistrationRequest {
	s.SupplierNameEn = &v
	return s
}

func (s *CreateSupplierRegistrationRequest) SetSupplierUrl(v string) *CreateSupplierRegistrationRequest {
	s.SupplierUrl = &v
	return s
}

func (s *CreateSupplierRegistrationRequest) Validate() error {
	return dara.Validate(s)
}
