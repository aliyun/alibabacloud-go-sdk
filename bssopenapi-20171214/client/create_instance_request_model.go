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
	// The client token that is used to ensure the idempotence of the request. The server checks whether a request that uses the same client token has been received. If a request that uses the same client token has been received, the server returns the same request result as the previous request.
	//
	// example:
	//
	// JASIOFKVNVIXXXXXX
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The logistics address of this order. This parameter is generally valid for physical orders.
	//
	// example:
	//
	// {"cityCode":"330100","cityName":"Hangzhou","contactName":"Test","countryCode":"","districtName":"Puyan Street","email":"\\*\\*@example.com","mobilePhone":"153564848844","phone":"1234567","provCode":"330000","provName":"Zhejiang","streetCode":"33010610","streetName":"Zhuantang","zipCode":"0000"}
	Logistics *string `json:"Logistics,omitempty" xml:"Logistics,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The details of the modules.
	Parameter []*CreateInstanceRequestParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Repeated"`
	// The subscription duration. Unit: month. The value must be an integral multiple of 12.
	//
	// >  This parameter is required if you create a subscription instance.
	//
	// example:
	//
	// 12
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The cycle type of the prepaid period
	//
	// - PricingCycle=1 indicates that the unit of the prepaid period is in years;
	//
	// - PricingCycle=2 indicates that the unit of the prepaid period is in months;
	//
	// - PricingCycle=3 indicates that the unit of the prepaid period is in days;
	//
	// - Default value: PricingCycle=2
	//
	// Applicable only to certain product types (ProductType being ddos_originpre_public_cn, ddosDip, ddoscoo, ddos_originpre_public_intl, ddosDip_intl, ddoscoo_intl)
	//
	// example:
	//
	// 2
	PricingCycle *int64 `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The code of the service to which the instance belongs. You can query the service code by calling the **QueryProductList*	- operation or viewing **Codes of Alibaba Cloud Services**.
	//
	// This parameter is required.
	//
	// example:
	//
	// rds
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The type of the service.
	//
	// example:
	//
	// rds
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The auto-renewal period. Unit: month.
	//
	// >  This parameter is required if the **RenewalStatus*	- parameter is set to **AutoRenewal**.
	//
	// example:
	//
	// 12
	RenewPeriod *int32 `json:"RenewPeriod,omitempty" xml:"RenewPeriod,omitempty"`
	// The renewal method. Valid values:
	//
	// 	- AutoRenewal: The instance is automatically renewed.
	//
	// 	- ManualRenewal: The instance is manually renewed.
	//
	// Default value: ManualRenewal.
	//
	// example:
	//
	// ManualRenewal
	RenewalStatus *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
	// The billing method. Valid values:
	//
	// 	- Subscription: the subscription billing method.
	//
	// 	- PayAsYouGo: the pay-as-you-go billing method.
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
	// The code property of the Nth module. Value of N: 1 to 100. If multiple module property parameters are involved, concatenate multiple parameters based on the value of N in sequence.
	//
	// This parameter is required.
	//
	// example:
	//
	// InstanceType
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The value property of the Nth module. Value of N: 1 to 100.
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
