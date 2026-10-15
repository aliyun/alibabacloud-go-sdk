// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateInstanceRequest
	GetClientToken() *string
	SetLogistics(v string) *CreateInstanceRequest
	GetLogistics() *string
	SetOwnerId(v int64) *CreateInstanceRequest
	GetOwnerId() *int64
	SetParameter(v []*CreateInstanceRequestParameter) *CreateInstanceRequest
	GetParameter() []*CreateInstanceRequestParameter
	SetPeriod(v int32) *CreateInstanceRequest
	GetPeriod() *int32
	SetPricingCycle(v int64) *CreateInstanceRequest
	GetPricingCycle() *int64
	SetProductCode(v string) *CreateInstanceRequest
	GetProductCode() *string
	SetProductType(v string) *CreateInstanceRequest
	GetProductType() *string
	SetRenewPeriod(v int32) *CreateInstanceRequest
	GetRenewPeriod() *int32
	SetRenewalStatus(v string) *CreateInstanceRequest
	GetRenewalStatus() *string
	SetSubscriptionType(v string) *CreateInstanceRequest
	GetSubscriptionType() *string
}

type CreateInstanceRequest struct {
	// The client idempotency token. The server checks whether a request with the same ClientToken has been made. If yes, the server returns the result of the previous call.
	//
	// example:
	//
	// JASIOFKVNVIXXXXXX
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The logistics address for this order. This parameter is typically valid for orders of physical products.
	//
	// example:
	//
	// {"cityCode":"330100","cityName":"Hangzhou","contactName":"Test","countryCode":"","districtName":"Puyan Street","email":"**@example.com","mobilePhone":"153564848844","phone":"1234567","provCode":"330000","provName":"Zhejiang Province","streetCode":"33010610","streetName":"Zhuantang","zipCode":"0000"}
	Logistics *string `json:"Logistics,omitempty" xml:"Logistics,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The list of modules.
	Parameter []*CreateInstanceRequestParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Repeated"`
	// The subscription period. Unit: months. For yearly subscription products, specify a multiple of 12.
	//
	// example:
	//
	// 12
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription period.
	//
	// - PricingCycle=1: The unit of the subscription period is year.
	//
	// - PricingCycle=2: The unit of the subscription period is month.
	//
	// - PricingCycle=3: The unit of the subscription period is day.
	//
	// Default value: PricingCycle=2.
	//
	// This parameter is applicable only to specific product types (ProductType is ddos_originpre_public_cn, ddosDip, ddoscoo, ddos_originpre_public_intl, ddosDip_intl, or ddoscoo_intl).
	//
	// example:
	//
	// 2
	PricingCycle *int64 `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The product code. You can call the **QueryProductList*	- operation to query the product code.
	//
	// This parameter is required.
	//
	// example:
	//
	// Cdn
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The product type.
	//
	// example:
	//
	// Cdn
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The auto-renewal epoch. Unit: months.
	//
	// > This parameter is required when you set **RenewalStatus*	- to **AutoRenewal*	- in Settings.
	//
	// example:
	//
	// 12
	RenewPeriod *int32 `json:"RenewPeriod,omitempty" xml:"RenewPeriod,omitempty"`
	// The auto-renewal status. Valid values:
	//
	// - AutoRenewal: auto-renewal is enabled.
	//
	// - ManualRenewal: manual renewal.
	//
	// Default value: ManualRenewal.
	//
	// example:
	//
	// ManualRenewal
	RenewalStatus *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
	// The billing method. Valid values:
	//
	// - Subscription: subscription.
	//
	// - PayAsYouGo: pay-as-you-go.
	//
	// This parameter is required.
	//
	// example:
	//
	// Subscription
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateInstanceRequest) GetLogistics() *string {
	return s.Logistics
}

func (s *CreateInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateInstanceRequest) GetParameter() []*CreateInstanceRequestParameter {
	return s.Parameter
}

func (s *CreateInstanceRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateInstanceRequest) GetPricingCycle() *int64 {
	return s.PricingCycle
}

func (s *CreateInstanceRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *CreateInstanceRequest) GetProductType() *string {
	return s.ProductType
}

func (s *CreateInstanceRequest) GetRenewPeriod() *int32 {
	return s.RenewPeriod
}

func (s *CreateInstanceRequest) GetRenewalStatus() *string {
	return s.RenewalStatus
}

func (s *CreateInstanceRequest) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *CreateInstanceRequest) SetClientToken(v string) *CreateInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateInstanceRequest) SetLogistics(v string) *CreateInstanceRequest {
	s.Logistics = &v
	return s
}

func (s *CreateInstanceRequest) SetOwnerId(v int64) *CreateInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateInstanceRequest) SetParameter(v []*CreateInstanceRequestParameter) *CreateInstanceRequest {
	s.Parameter = v
	return s
}

func (s *CreateInstanceRequest) SetPeriod(v int32) *CreateInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateInstanceRequest) SetPricingCycle(v int64) *CreateInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateInstanceRequest) SetProductCode(v string) *CreateInstanceRequest {
	s.ProductCode = &v
	return s
}

func (s *CreateInstanceRequest) SetProductType(v string) *CreateInstanceRequest {
	s.ProductType = &v
	return s
}

func (s *CreateInstanceRequest) SetRenewPeriod(v int32) *CreateInstanceRequest {
	s.RenewPeriod = &v
	return s
}

func (s *CreateInstanceRequest) SetRenewalStatus(v string) *CreateInstanceRequest {
	s.RenewalStatus = &v
	return s
}

func (s *CreateInstanceRequest) SetSubscriptionType(v string) *CreateInstanceRequest {
	s.SubscriptionType = &v
	return s
}

func (s *CreateInstanceRequest) Validate() error {
	if s.Parameter != nil {
		for _, item := range s.Parameter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateInstanceRequestParameter struct {
	// The code of the Nth module property. Valid values of N: 1 to 100. If multiple module properties exist, concatenate them in the format of N=1, 2, 3...N.
	//
	// This parameter is required.
	//
	// example:
	//
	// InstanceType
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The value of the Nth module property. Valid values of N: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// disk
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateInstanceRequestParameter) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestParameter) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestParameter) GetCode() *string {
	return s.Code
}

func (s *CreateInstanceRequestParameter) GetValue() *string {
	return s.Value
}

func (s *CreateInstanceRequestParameter) SetCode(v string) *CreateInstanceRequestParameter {
	s.Code = &v
	return s
}

func (s *CreateInstanceRequestParameter) SetValue(v string) *CreateInstanceRequestParameter {
	s.Value = &v
	return s
}

func (s *CreateInstanceRequestParameter) Validate() error {
	return dara.Validate(s)
}
