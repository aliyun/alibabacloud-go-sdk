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
	// Specifies whether to automatically pay for the provided products.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The list of products to be purchased. Separate them with commas (,).
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
	// Specifies whether to automatically renew the product.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The billing method. Only POSTPAY is supported.
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The purchase duration. You can use this parameter together with pricingCycle. Only 1 is supported.
	//
	// example:
	//
	// 1
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The properties of the instance.
	//
	// 	- DataWorks_share: [ { "Code": "region", "Value": "cn-shanghai" } ]
	//
	// 	- OSS_share: [ { "Code": "commodity_type", "Value": "oss", "Name": "Object Storage Service" }, { "Code": "ord_time", "Value": "1:Hour", "Name": "1 Hour" } ]
	//
	// 	- PAI_share: None
	//
	// 	- China bid MaxCompute_share: [ { "Code": "region", "Value": "cn-hangzhou" }, { "Code": "odps_specification_type", "Value": "OdpsStandard" }, { "Code": "ord_time", "Value": "1:Hour" } ]
	//
	// 	- International bid MaxCompute_share: [ { "Code": "region", "Value": "cn-hangzhou" }, { "Code": "ord_time", "Value": "1:Hour" } ]
	InstanceProperties []*CreateProductOrdersRequestProductsInstanceProperties `json:"InstanceProperties,omitempty" xml:"InstanceProperties,omitempty" type:"Repeated"`
	// The type of the order. Only BUY is supported.
	//
	// example:
	//
	// BUY
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The billing cycle. Valid values:
	//
	// 	- Month: The price is calculated every month. DataWorks_share only supports Month.
	//
	// 	- Hour: The price is calculated every hour. OSS_share and MaxCompute_share only support Hour.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The product code. Valid values:
	//
	// 	- DataWorks_share: pay-as-you-go DataWorks
	//
	// 	- MaxCompute_share: pay-as-you-go MaxCompute
	//
	// 	- PAI_share: pay-as-you-go PAI.
	//
	// 	- OSS_share: pay-as-you-go OSS
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
	// The property code.
	//
	// example:
	//
	// commodity_type
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The property name.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The property value.
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
