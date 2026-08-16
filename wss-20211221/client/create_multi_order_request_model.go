// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultiOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelCookie(v string) *CreateMultiOrderRequest
	GetChannelCookie() *string
	SetOrderItems(v []*CreateMultiOrderRequestOrderItems) *CreateMultiOrderRequest
	GetOrderItems() []*CreateMultiOrderRequestOrderItems
	SetOrderType(v string) *CreateMultiOrderRequest
	GetOrderType() *string
	SetProperties(v map[string]*string) *CreateMultiOrderRequest
	GetProperties() map[string]*string
	SetResellerOwnerUid(v int64) *CreateMultiOrderRequest
	GetResellerOwnerUid() *int64
}

type CreateMultiOrderRequest struct {
	// The channel cookie information.
	ChannelCookie *string `json:"ChannelCookie,omitempty" xml:"ChannelCookie,omitempty"`
	// The product information.
	OrderItems []*CreateMultiOrderRequestOrderItems `json:"OrderItems,omitempty" xml:"OrderItems,omitempty" type:"Repeated"`
	// The order type.
	//
	// example:
	//
	// create
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The extended properties.
	Properties map[string]*string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// The UID of the reseller owner.
	ResellerOwnerUid *int64 `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
}

func (s CreateMultiOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateMultiOrderRequest) GetChannelCookie() *string {
	return s.ChannelCookie
}

func (s *CreateMultiOrderRequest) GetOrderItems() []*CreateMultiOrderRequestOrderItems {
	return s.OrderItems
}

func (s *CreateMultiOrderRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *CreateMultiOrderRequest) GetProperties() map[string]*string {
	return s.Properties
}

func (s *CreateMultiOrderRequest) GetResellerOwnerUid() *int64 {
	return s.ResellerOwnerUid
}

func (s *CreateMultiOrderRequest) SetChannelCookie(v string) *CreateMultiOrderRequest {
	s.ChannelCookie = &v
	return s
}

func (s *CreateMultiOrderRequest) SetOrderItems(v []*CreateMultiOrderRequestOrderItems) *CreateMultiOrderRequest {
	s.OrderItems = v
	return s
}

func (s *CreateMultiOrderRequest) SetOrderType(v string) *CreateMultiOrderRequest {
	s.OrderType = &v
	return s
}

func (s *CreateMultiOrderRequest) SetProperties(v map[string]*string) *CreateMultiOrderRequest {
	s.Properties = v
	return s
}

func (s *CreateMultiOrderRequest) SetResellerOwnerUid(v int64) *CreateMultiOrderRequest {
	s.ResellerOwnerUid = &v
	return s
}

func (s *CreateMultiOrderRequest) Validate() error {
	if s.OrderItems != nil {
		for _, item := range s.OrderItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateMultiOrderRequestOrderItems struct {
	// The quantity to purchase.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// Specifies whether to enable automatic payment.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// Specifies whether this is a change purchase.
	BuyChange *bool `json:"BuyChange,omitempty" xml:"BuyChange,omitempty"`
	// The product modules.
	Components []*CreateMultiOrderRequestOrderItemsComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// The list of instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The callback URL after the payment is completed.
	PaidCallBackUrl *string `json:"PaidCallBackUrl,omitempty" xml:"PaidCallBackUrl,omitempty"`
	// The subscription duration. Valid values:
	//
	// - If PeriodUnit is set to Year: 1, 2, 3, and 5.
	//
	// - If PeriodUnit is set to Month: 1, 2, 3, and 6.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration for a subscription instance.
	//
	// > This parameter is required only when the billing method of the instance is subscription. This parameter is case-sensitive. Make sure that the value is spelled correctly.
	//
	// example:
	//
	// Year
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The promotion ID.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The list of resource IDs.
	//
	// > For monthly duration packages, this parameter corresponds to the cloud desktop ID. This parameter is required when OrderType is not set to `create`.
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// The resource type.
	//
	// > This parameter is case-sensitive. Make sure that the value is spelled correctly.
	//
	// This parameter is required.
	//
	// example:
	//
	// DurationPackage
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s CreateMultiOrderRequestOrderItems) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiOrderRequestOrderItems) GoString() string {
	return s.String()
}

func (s *CreateMultiOrderRequestOrderItems) GetAmount() *int32 {
	return s.Amount
}

func (s *CreateMultiOrderRequestOrderItems) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateMultiOrderRequestOrderItems) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateMultiOrderRequestOrderItems) GetBuyChange() *bool {
	return s.BuyChange
}

func (s *CreateMultiOrderRequestOrderItems) GetComponents() []*CreateMultiOrderRequestOrderItemsComponents {
	return s.Components
}

func (s *CreateMultiOrderRequestOrderItems) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *CreateMultiOrderRequestOrderItems) GetPaidCallBackUrl() *string {
	return s.PaidCallBackUrl
}

func (s *CreateMultiOrderRequestOrderItems) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateMultiOrderRequestOrderItems) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateMultiOrderRequestOrderItems) GetPromotionId() *string {
	return s.PromotionId
}

func (s *CreateMultiOrderRequestOrderItems) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *CreateMultiOrderRequestOrderItems) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateMultiOrderRequestOrderItems) SetAmount(v int32) *CreateMultiOrderRequestOrderItems {
	s.Amount = &v
	return s
}

func (s *CreateMultiOrderRequestOrderItems) SetAutoPay(v bool) *CreateMultiOrderRequestOrderItems {
	s.AutoPay = &v
	return s
}

func (s *CreateMultiOrderRequestOrderItems) SetAutoRenew(v bool) *CreateMultiOrderRequestOrderItems {
	s.AutoRenew = &v
	return s
}

func (s *CreateMultiOrderRequestOrderItems) SetBuyChange(v bool) *CreateMultiOrderRequestOrderItems {
	s.BuyChange = &v
	return s
}

func (s *CreateMultiOrderRequestOrderItems) SetComponents(v []*CreateMultiOrderRequestOrderItemsComponents) *CreateMultiOrderRequestOrderItems {
	s.Components = v
	return s
}

func (s *CreateMultiOrderRequestOrderItems) SetInstanceIds(v []*string) *CreateMultiOrderRequestOrderItems {
	s.InstanceIds = v
	return s
}

func (s *CreateMultiOrderRequestOrderItems) SetPaidCallBackUrl(v string) *CreateMultiOrderRequestOrderItems {
	s.PaidCallBackUrl = &v
	return s
}

func (s *CreateMultiOrderRequestOrderItems) SetPeriod(v int32) *CreateMultiOrderRequestOrderItems {
	s.Period = &v
	return s
}

func (s *CreateMultiOrderRequestOrderItems) SetPeriodUnit(v string) *CreateMultiOrderRequestOrderItems {
	s.PeriodUnit = &v
	return s
}

func (s *CreateMultiOrderRequestOrderItems) SetPromotionId(v string) *CreateMultiOrderRequestOrderItems {
	s.PromotionId = &v
	return s
}

func (s *CreateMultiOrderRequestOrderItems) SetResourceIds(v []*string) *CreateMultiOrderRequestOrderItems {
	s.ResourceIds = v
	return s
}

func (s *CreateMultiOrderRequestOrderItems) SetResourceType(v string) *CreateMultiOrderRequestOrderItems {
	s.ResourceType = &v
	return s
}

func (s *CreateMultiOrderRequestOrderItems) Validate() error {
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

type CreateMultiOrderRequestOrderItemsComponents struct {
	// The key of the module.
	//
	// example:
	//
	// RegionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the module.
	//
	// The following section describes the sample values and valid values for each key of the Enterprise Edition monthly duration package:
	//
	// - RegionId: ap-southeast-1
	//
	// - InstanceType: eds.enterprise_office.4c8g
	//
	// - DurationType (hours): Valid values:
	//
	//    - 120
	//
	//    - 250
	//
	// - OsType: Valid values:
	//
	//    - Windows
	//
	//    - Linux
	//
	// - RootDiskSize (GiB): 80
	//
	// - RootDiskCategory: Valid values:
	//
	//    - cloud_efficiency (ultra cloud disk)
	//
	//    - cloud_auto (ultra-fast cloud disk)
	//
	//    - cloud_essd (enhanced standard SSD. Only specific instance types support this value.)
	//
	// - RootPerformanceLevel: Valid values:
	//
	//    - PL0
	//
	//    - PL1
	//
	//    - PL2
	//
	//    - PL3
	//
	// - DataDiskSize (GiB): Valid values are the same as those of RootDiskSize.
	//
	// - DataDiskCategory: Valid values are the same as those of RootDiskCategory.
	//
	// - DataPerformanceLevel: Valid values are the same as those of RootPerformanceLevel.
	//
	// example:
	//
	// cn-shanghai
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateMultiOrderRequestOrderItemsComponents) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiOrderRequestOrderItemsComponents) GoString() string {
	return s.String()
}

func (s *CreateMultiOrderRequestOrderItemsComponents) GetKey() *string {
	return s.Key
}

func (s *CreateMultiOrderRequestOrderItemsComponents) GetValue() *string {
	return s.Value
}

func (s *CreateMultiOrderRequestOrderItemsComponents) SetKey(v string) *CreateMultiOrderRequestOrderItemsComponents {
	s.Key = &v
	return s
}

func (s *CreateMultiOrderRequestOrderItemsComponents) SetValue(v string) *CreateMultiOrderRequestOrderItemsComponents {
	s.Value = &v
	return s
}

func (s *CreateMultiOrderRequestOrderItemsComponents) Validate() error {
	return dara.Validate(s)
}
