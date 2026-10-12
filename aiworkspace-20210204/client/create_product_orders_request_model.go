// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProductOrdersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *CreateProductOrdersRequest
	GetAutoPay() *bool
	SetProducts(v []*CreateProductOrdersRequestProducts) *CreateProductOrdersRequest
	GetProducts() []*CreateProductOrdersRequestProducts
}

type CreateProductOrdersRequest struct {
	// Specifies whether automatic payment is applied to all products specified in the Products parameter. Valid values:
	//
	// - true: automatic payment
	//
	// - false: no automatic payment
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The list of products to purchase, separated by commas.
	Products []*CreateProductOrdersRequestProducts `json:"Products,omitempty" xml:"Products,omitempty" type:"Repeated"`
}

func (s CreateProductOrdersRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProductOrdersRequest) GoString() string {
	return s.String()
}

func (s *CreateProductOrdersRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateProductOrdersRequest) GetProducts() []*CreateProductOrdersRequestProducts {
	return s.Products
}

func (s *CreateProductOrdersRequest) SetAutoPay(v bool) *CreateProductOrdersRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateProductOrdersRequest) SetProducts(v []*CreateProductOrdersRequestProducts) *CreateProductOrdersRequest {
	s.Products = v
	return s
}

func (s *CreateProductOrdersRequest) Validate() error {
	if s.Products != nil {
		for _, item := range s.Products {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateProductOrdersRequestProducts struct {
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// - true: Auto-renewal is enabled.
	//
	// - false: Auto-renewal is disabled.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The billing method. Currently, only POSTPAY is supported.
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The purchase duration. This parameter is used together with PricingCycle. Currently, only 1 is supported.
	//
	// example:
	//
	// 1
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The list of instance property information.
	//
	// - DataWorks_share:
	//
	// [ {
	//
	// "Code": "region",
	//
	// "Value": "cn-shanghai"
	//
	// }
	//
	// ]
	//
	// - OSS_share:
	//
	// [ {
	//
	// "Code": "commodity_type",
	//
	// "Value": "oss",
	//
	// "Name": "Object Storage Service (OSS)"
	//
	// },
	//
	// {
	//
	// "Code": "ord_time",
	//
	// "Value": "1:Hour",
	//
	// "Name": "1 hour"
	//
	// }
	//
	// ]
	//
	// - PAI_share: None
	//
	// - MaxCompute_share for the Chinese site:
	//
	// [
	//
	// {
	//
	// "Code": "region",
	//
	// "Value": "cn-hangzhou"
	//
	// },
	//
	// {
	//
	// "Code": "odps_specification_type",
	//
	// "Value": "OdpsStandard"
	//
	// },
	//
	// {
	//
	// "Code": "ord_time",
	//
	// "Value": "1:Hour"
	//
	// }
	//
	// ]
	//
	// - MaxCompute_share for the international site:
	//
	// [
	//
	// {
	//
	// "Code": "region",
	//
	// "Value": "cn-hangzhou"
	//
	// },
	//
	// {
	//
	// "Code": "ord_time",
	//
	// "Value": "1:Hour"
	//
	// }
	//
	// ]
	InstanceProperties []*CreateProductOrdersRequestProductsInstanceProperties `json:"InstanceProperties,omitempty" xml:"InstanceProperties,omitempty" type:"Repeated"`
	// The order type. Currently, only BUY is supported.
	//
	// example:
	//
	// BUY
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The pricing unit. Valid values:
	//
	// - Month: Monthly pricing. Only DataWorks_share supports this value.
	//
	// - Hour: Hourly pricing. Only OSS_share and MaxCompute_share support this value.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The product code. Valid values:
	//
	// - DataWorks_share: Pay-as-you-go DataWorks.
	//
	// - MaxCompute_share: Pay-as-you-go MaxCompute.
	//
	// - PAI_share: Pay-as-you-go PAI.
	//
	// - OSS_share: Pay-as-you-go OSS.
	//
	// example:
	//
	// DataWorks_share
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s CreateProductOrdersRequestProducts) String() string {
	return dara.Prettify(s)
}

func (s CreateProductOrdersRequestProducts) GoString() string {
	return s.String()
}

func (s *CreateProductOrdersRequestProducts) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateProductOrdersRequestProducts) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateProductOrdersRequestProducts) GetDuration() *int64 {
	return s.Duration
}

func (s *CreateProductOrdersRequestProducts) GetInstanceProperties() []*CreateProductOrdersRequestProductsInstanceProperties {
	return s.InstanceProperties
}

func (s *CreateProductOrdersRequestProducts) GetOrderType() *string {
	return s.OrderType
}

func (s *CreateProductOrdersRequestProducts) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CreateProductOrdersRequestProducts) GetProductCode() *string {
	return s.ProductCode
}

func (s *CreateProductOrdersRequestProducts) SetAutoRenew(v bool) *CreateProductOrdersRequestProducts {
	s.AutoRenew = &v
	return s
}

func (s *CreateProductOrdersRequestProducts) SetChargeType(v string) *CreateProductOrdersRequestProducts {
	s.ChargeType = &v
	return s
}

func (s *CreateProductOrdersRequestProducts) SetDuration(v int64) *CreateProductOrdersRequestProducts {
	s.Duration = &v
	return s
}

func (s *CreateProductOrdersRequestProducts) SetInstanceProperties(v []*CreateProductOrdersRequestProductsInstanceProperties) *CreateProductOrdersRequestProducts {
	s.InstanceProperties = v
	return s
}

func (s *CreateProductOrdersRequestProducts) SetOrderType(v string) *CreateProductOrdersRequestProducts {
	s.OrderType = &v
	return s
}

func (s *CreateProductOrdersRequestProducts) SetPricingCycle(v string) *CreateProductOrdersRequestProducts {
	s.PricingCycle = &v
	return s
}

func (s *CreateProductOrdersRequestProducts) SetProductCode(v string) *CreateProductOrdersRequestProducts {
	s.ProductCode = &v
	return s
}

func (s *CreateProductOrdersRequestProducts) Validate() error {
	if s.InstanceProperties != nil {
		for _, item := range s.InstanceProperties {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateProductOrdersRequestProductsInstanceProperties struct {
	// The code of the instance property.
	//
	// example:
	//
	// commodity_type
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The name of the instance property.
	//
	// example:
	//
	// Object Storage Service (OSS)
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the instance property.
	//
	// example:
	//
	// oss
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateProductOrdersRequestProductsInstanceProperties) String() string {
	return dara.Prettify(s)
}

func (s CreateProductOrdersRequestProductsInstanceProperties) GoString() string {
	return s.String()
}

func (s *CreateProductOrdersRequestProductsInstanceProperties) GetCode() *string {
	return s.Code
}

func (s *CreateProductOrdersRequestProductsInstanceProperties) GetName() *string {
	return s.Name
}

func (s *CreateProductOrdersRequestProductsInstanceProperties) GetValue() *string {
	return s.Value
}

func (s *CreateProductOrdersRequestProductsInstanceProperties) SetCode(v string) *CreateProductOrdersRequestProductsInstanceProperties {
	s.Code = &v
	return s
}

func (s *CreateProductOrdersRequestProductsInstanceProperties) SetName(v string) *CreateProductOrdersRequestProductsInstanceProperties {
	s.Name = &v
	return s
}

func (s *CreateProductOrdersRequestProductsInstanceProperties) SetValue(v string) *CreateProductOrdersRequestProductsInstanceProperties {
	s.Value = &v
	return s
}

func (s *CreateProductOrdersRequestProductsInstanceProperties) Validate() error {
	return dara.Validate(s)
}
