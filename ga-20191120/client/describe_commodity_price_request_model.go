// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommodityPriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrders(v []*DescribeCommodityPriceRequestOrders) *DescribeCommodityPriceRequest
	GetOrders() []*DescribeCommodityPriceRequestOrders
	SetPromotionOptionNo(v string) *DescribeCommodityPriceRequest
	GetPromotionOptionNo() *string
	SetRegionId(v string) *DescribeCommodityPriceRequest
	GetRegionId() *string
}

type DescribeCommodityPriceRequest struct {
	// The commodity orders.
	//
	// This parameter is required.
	Orders []*DescribeCommodityPriceRequestOrders `json:"Orders,omitempty" xml:"Orders,omitempty" type:"Repeated"`
	// The coupon code.
	//
	// >  This parameter is unavailable on the China site (aliyun.com).
	//
	// example:
	//
	// 50003298014****
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
	// The ID of the region where the Global Accelerator (GA) instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCommodityPriceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommodityPriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeCommodityPriceRequest) GetOrders() []*DescribeCommodityPriceRequestOrders {
	return s.Orders
}

func (s *DescribeCommodityPriceRequest) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *DescribeCommodityPriceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCommodityPriceRequest) SetOrders(v []*DescribeCommodityPriceRequestOrders) *DescribeCommodityPriceRequest {
	s.Orders = v
	return s
}

func (s *DescribeCommodityPriceRequest) SetPromotionOptionNo(v string) *DescribeCommodityPriceRequest {
	s.PromotionOptionNo = &v
	return s
}

func (s *DescribeCommodityPriceRequest) SetRegionId(v string) *DescribeCommodityPriceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCommodityPriceRequest) Validate() error {
	if s.Orders != nil {
		for _, item := range s.Orders {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCommodityPriceRequestOrders struct {
	// The billing method. Set the value to **PREPAY**, which specifies the subscription billing method.
	//
	// example:
	//
	// PREPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The commodity code.
	//
	// Valid values on the China site (aliyun.com):
	//
	// 	- **ga_gapluspre_public_cn**: GA instance.
	//
	// 	- **ga_plusbwppre_public_cn**: basic bandwidth plan.
	//
	// Valid values on the international site (alibabacloud.com):
	//
	// 	- **ga_pluspre_public_intl**: GA instance.
	//
	// 	- **ga_bwppreintl_public_intl:*	- basic bandwidth plan.
	//
	// example:
	//
	// ga_gapluspre_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The information about commodity modules.
	//
	// The information varies based on the commodity module.
	Components []*DescribeCommodityPriceRequestOrdersComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// The subscription duration.
	//
	// 	- Valid values if you set **PricingCycle*	- to **Month**: **1*	- to **9**.
	//
	// 	- Valid values if you set **PricingCycle*	- to **Year**: **1*	- to **3**.
	//
	// example:
	//
	// 1
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The type of the order. Valid values:
	//
	// 	- **BUY**: purchase order.
	//
	// 	- **RENEW**: renewal order.
	//
	// 	- **UPGRADE**: upgrade order.
	//
	// example:
	//
	// BUY
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The billing cycle. Valid values:
	//
	// 	- **Month**
	//
	// 	- **Year**
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The number of instances that you want to purchase.
	//
	// example:
	//
	// 1
	Quantity *int64 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
}

func (s DescribeCommodityPriceRequestOrders) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommodityPriceRequestOrders) GoString() string {
	return s.String()
}

func (s *DescribeCommodityPriceRequestOrders) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeCommodityPriceRequestOrders) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeCommodityPriceRequestOrders) GetComponents() []*DescribeCommodityPriceRequestOrdersComponents {
	return s.Components
}

func (s *DescribeCommodityPriceRequestOrders) GetDuration() *int64 {
	return s.Duration
}

func (s *DescribeCommodityPriceRequestOrders) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribeCommodityPriceRequestOrders) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *DescribeCommodityPriceRequestOrders) GetQuantity() *int64 {
	return s.Quantity
}

func (s *DescribeCommodityPriceRequestOrders) SetChargeType(v string) *DescribeCommodityPriceRequestOrders {
	s.ChargeType = &v
	return s
}

func (s *DescribeCommodityPriceRequestOrders) SetCommodityCode(v string) *DescribeCommodityPriceRequestOrders {
	s.CommodityCode = &v
	return s
}

func (s *DescribeCommodityPriceRequestOrders) SetComponents(v []*DescribeCommodityPriceRequestOrdersComponents) *DescribeCommodityPriceRequestOrders {
	s.Components = v
	return s
}

func (s *DescribeCommodityPriceRequestOrders) SetDuration(v int64) *DescribeCommodityPriceRequestOrders {
	s.Duration = &v
	return s
}

func (s *DescribeCommodityPriceRequestOrders) SetOrderType(v string) *DescribeCommodityPriceRequestOrders {
	s.OrderType = &v
	return s
}

func (s *DescribeCommodityPriceRequestOrders) SetPricingCycle(v string) *DescribeCommodityPriceRequestOrders {
	s.PricingCycle = &v
	return s
}

func (s *DescribeCommodityPriceRequestOrders) SetQuantity(v int64) *DescribeCommodityPriceRequestOrders {
	s.Quantity = &v
	return s
}

func (s *DescribeCommodityPriceRequestOrders) Validate() error {
	if s.Components != nil {
		for _, item := range s.Components {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCommodityPriceRequestOrdersComponents struct {
	// The code of the commodity module.
	//
	// The information varies based on the commodity module. Examples: **instance*	- (GA instance) and **ord_time*	- (subscription duration).
	//
	// example:
	//
	// instance
	ComponentCode *string `json:"ComponentCode,omitempty" xml:"ComponentCode,omitempty"`
	// The attributes of commodity modules.
	//
	// The information varies based on the commodity module.
	Properties []*DescribeCommodityPriceRequestOrdersComponentsProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Repeated"`
}

func (s DescribeCommodityPriceRequestOrdersComponents) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommodityPriceRequestOrdersComponents) GoString() string {
	return s.String()
}

func (s *DescribeCommodityPriceRequestOrdersComponents) GetComponentCode() *string {
	return s.ComponentCode
}

func (s *DescribeCommodityPriceRequestOrdersComponents) GetProperties() []*DescribeCommodityPriceRequestOrdersComponentsProperties {
	return s.Properties
}

func (s *DescribeCommodityPriceRequestOrdersComponents) SetComponentCode(v string) *DescribeCommodityPriceRequestOrdersComponents {
	s.ComponentCode = &v
	return s
}

func (s *DescribeCommodityPriceRequestOrdersComponents) SetProperties(v []*DescribeCommodityPriceRequestOrdersComponentsProperties) *DescribeCommodityPriceRequestOrdersComponents {
	s.Properties = v
	return s
}

func (s *DescribeCommodityPriceRequestOrdersComponents) Validate() error {
	if s.Properties != nil {
		for _, item := range s.Properties {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCommodityPriceRequestOrdersComponentsProperties struct {
	// The code of the attribute of the commodity module.
	//
	// The information varies based on the commodity module. Examples: **instance*	- (GA instance) and **ord_time*	- (subscription duration).
	//
	// example:
	//
	// instance
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The value of the attribute.
	//
	// The information varies based on the commodity module. Examples: **instance_fee*	- (GA instance fee) and **1:Month*	- (one-month subscription).
	//
	// example:
	//
	// instance_fee
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCommodityPriceRequestOrdersComponentsProperties) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommodityPriceRequestOrdersComponentsProperties) GoString() string {
	return s.String()
}

func (s *DescribeCommodityPriceRequestOrdersComponentsProperties) GetCode() *string {
	return s.Code
}

func (s *DescribeCommodityPriceRequestOrdersComponentsProperties) GetValue() *string {
	return s.Value
}

func (s *DescribeCommodityPriceRequestOrdersComponentsProperties) SetCode(v string) *DescribeCommodityPriceRequestOrdersComponentsProperties {
	s.Code = &v
	return s
}

func (s *DescribeCommodityPriceRequestOrdersComponentsProperties) SetValue(v string) *DescribeCommodityPriceRequestOrdersComponentsProperties {
	s.Value = &v
	return s
}

func (s *DescribeCommodityPriceRequestOrdersComponentsProperties) Validate() error {
	return dara.Validate(s)
}
