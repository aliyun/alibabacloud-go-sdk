// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultiOrderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelCookie(v string) *CreateMultiOrderShrinkRequest
	GetChannelCookie() *string
	SetOrderItems(v []*CreateMultiOrderShrinkRequestOrderItems) *CreateMultiOrderShrinkRequest
	GetOrderItems() []*CreateMultiOrderShrinkRequestOrderItems
	SetOrderType(v string) *CreateMultiOrderShrinkRequest
	GetOrderType() *string
	SetPropertiesShrink(v string) *CreateMultiOrderShrinkRequest
	GetPropertiesShrink() *string
	SetResellerOwnerUid(v int64) *CreateMultiOrderShrinkRequest
	GetResellerOwnerUid() *int64
}

type CreateMultiOrderShrinkRequest struct {
	ChannelCookie *string `json:"ChannelCookie,omitempty" xml:"ChannelCookie,omitempty"`
	// The items in the order.
	OrderItems []*CreateMultiOrderShrinkRequestOrderItems `json:"OrderItems,omitempty" xml:"OrderItems,omitempty" type:"Repeated"`
	// The order type.
	//
	// example:
	//
	// create
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The extended properties.
	PropertiesShrink *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	ResellerOwnerUid *int64  `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
}

func (s CreateMultiOrderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateMultiOrderShrinkRequest) GetChannelCookie() *string {
	return s.ChannelCookie
}

func (s *CreateMultiOrderShrinkRequest) GetOrderItems() []*CreateMultiOrderShrinkRequestOrderItems {
	return s.OrderItems
}

func (s *CreateMultiOrderShrinkRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *CreateMultiOrderShrinkRequest) GetPropertiesShrink() *string {
	return s.PropertiesShrink
}

func (s *CreateMultiOrderShrinkRequest) GetResellerOwnerUid() *int64 {
	return s.ResellerOwnerUid
}

func (s *CreateMultiOrderShrinkRequest) SetChannelCookie(v string) *CreateMultiOrderShrinkRequest {
	s.ChannelCookie = &v
	return s
}

func (s *CreateMultiOrderShrinkRequest) SetOrderItems(v []*CreateMultiOrderShrinkRequestOrderItems) *CreateMultiOrderShrinkRequest {
	s.OrderItems = v
	return s
}

func (s *CreateMultiOrderShrinkRequest) SetOrderType(v string) *CreateMultiOrderShrinkRequest {
	s.OrderType = &v
	return s
}

func (s *CreateMultiOrderShrinkRequest) SetPropertiesShrink(v string) *CreateMultiOrderShrinkRequest {
	s.PropertiesShrink = &v
	return s
}

func (s *CreateMultiOrderShrinkRequest) SetResellerOwnerUid(v int64) *CreateMultiOrderShrinkRequest {
	s.ResellerOwnerUid = &v
	return s
}

func (s *CreateMultiOrderShrinkRequest) Validate() error {
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

type CreateMultiOrderShrinkRequestOrderItems struct {
	// The number of resources to purchase.
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
	BuyChange *bool `json:"BuyChange,omitempty" xml:"BuyChange,omitempty"`
	// The components that define the resource.
	Components  []*CreateMultiOrderShrinkRequestOrderItemsComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	InstanceIds []*string                                            `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The subscription period. Valid values:
	//
	// - If `PeriodUnit` is set to `Year`, the valid values are 1, 2, 3, and 5.
	//
	// - If `PeriodUnit` is set to `Month`, the valid values are 1, 2, 3, and 6.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The time unit of the subscription duration.
	//
	// > This parameter is required for prepaid instances and is case-sensitive.
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
	// A list of resource IDs.
	//
	// > For a monthly duration package, this parameter specifies the IDs of the cloud desktops. This parameter is required unless the `OrderType` is `create`.
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// The type of the resource.
	//
	// > This parameter is case-sensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// DurationPackage
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s CreateMultiOrderShrinkRequestOrderItems) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiOrderShrinkRequestOrderItems) GoString() string {
	return s.String()
}

func (s *CreateMultiOrderShrinkRequestOrderItems) GetAmount() *int32 {
	return s.Amount
}

func (s *CreateMultiOrderShrinkRequestOrderItems) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateMultiOrderShrinkRequestOrderItems) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateMultiOrderShrinkRequestOrderItems) GetBuyChange() *bool {
	return s.BuyChange
}

func (s *CreateMultiOrderShrinkRequestOrderItems) GetComponents() []*CreateMultiOrderShrinkRequestOrderItemsComponents {
	return s.Components
}

func (s *CreateMultiOrderShrinkRequestOrderItems) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *CreateMultiOrderShrinkRequestOrderItems) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateMultiOrderShrinkRequestOrderItems) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateMultiOrderShrinkRequestOrderItems) GetPromotionId() *string {
	return s.PromotionId
}

func (s *CreateMultiOrderShrinkRequestOrderItems) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *CreateMultiOrderShrinkRequestOrderItems) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateMultiOrderShrinkRequestOrderItems) SetAmount(v int32) *CreateMultiOrderShrinkRequestOrderItems {
	s.Amount = &v
	return s
}

func (s *CreateMultiOrderShrinkRequestOrderItems) SetAutoPay(v bool) *CreateMultiOrderShrinkRequestOrderItems {
	s.AutoPay = &v
	return s
}

func (s *CreateMultiOrderShrinkRequestOrderItems) SetAutoRenew(v bool) *CreateMultiOrderShrinkRequestOrderItems {
	s.AutoRenew = &v
	return s
}

func (s *CreateMultiOrderShrinkRequestOrderItems) SetBuyChange(v bool) *CreateMultiOrderShrinkRequestOrderItems {
	s.BuyChange = &v
	return s
}

func (s *CreateMultiOrderShrinkRequestOrderItems) SetComponents(v []*CreateMultiOrderShrinkRequestOrderItemsComponents) *CreateMultiOrderShrinkRequestOrderItems {
	s.Components = v
	return s
}

func (s *CreateMultiOrderShrinkRequestOrderItems) SetInstanceIds(v []*string) *CreateMultiOrderShrinkRequestOrderItems {
	s.InstanceIds = v
	return s
}

func (s *CreateMultiOrderShrinkRequestOrderItems) SetPeriod(v int32) *CreateMultiOrderShrinkRequestOrderItems {
	s.Period = &v
	return s
}

func (s *CreateMultiOrderShrinkRequestOrderItems) SetPeriodUnit(v string) *CreateMultiOrderShrinkRequestOrderItems {
	s.PeriodUnit = &v
	return s
}

func (s *CreateMultiOrderShrinkRequestOrderItems) SetPromotionId(v string) *CreateMultiOrderShrinkRequestOrderItems {
	s.PromotionId = &v
	return s
}

func (s *CreateMultiOrderShrinkRequestOrderItems) SetResourceIds(v []*string) *CreateMultiOrderShrinkRequestOrderItems {
	s.ResourceIds = v
	return s
}

func (s *CreateMultiOrderShrinkRequestOrderItems) SetResourceType(v string) *CreateMultiOrderShrinkRequestOrderItems {
	s.ResourceType = &v
	return s
}

func (s *CreateMultiOrderShrinkRequestOrderItems) Validate() error {
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

type CreateMultiOrderShrinkRequestOrderItemsComponents struct {
	// The key of the component.
	//
	// example:
	//
	// RegionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the component.
	//
	// Example and valid values for the keys of a monthly duration package (Enterprise Edition):
	//
	// - RegionId: cn-shanghai
	//
	// - InstanceType: eds.enterprise_office.4c8g
	//
	// - DurationType (in hours): Valid values:
	//
	//   - 120
	//
	//   - 250
	//
	// - OsType: Valid values:
	//
	//   - Windows
	//
	//   - Linux
	//
	// - RootDiskSize (in GiB): 80
	//
	// - RootDiskCategory: Valid values:
	//
	//   - cloud_efficiency (Ultra Disk)
	//
	//   - cloud_auto (ESSD AutoPL Disk)
	//
	//   - `cloud_essd` (Enhanced SSD). This value is supported only by specific instance types.
	//
	// - RootPerformanceLevel: Valid values:
	//
	//   - PL0
	//
	//   - PL1
	//
	//   - PL2
	//
	//   - PL3
	//
	// - DataDiskSize (in GiB): Same as `RootDiskSize`.
	//
	// - DataDiskCategory: Same as `RootDiskCategory`.
	//
	// - DataPerformanceLevel: Same as `RootPerformanceLevel`.
	//
	// example:
	//
	// cn-shanghai
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateMultiOrderShrinkRequestOrderItemsComponents) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiOrderShrinkRequestOrderItemsComponents) GoString() string {
	return s.String()
}

func (s *CreateMultiOrderShrinkRequestOrderItemsComponents) GetKey() *string {
	return s.Key
}

func (s *CreateMultiOrderShrinkRequestOrderItemsComponents) GetValue() *string {
	return s.Value
}

func (s *CreateMultiOrderShrinkRequestOrderItemsComponents) SetKey(v string) *CreateMultiOrderShrinkRequestOrderItemsComponents {
	s.Key = &v
	return s
}

func (s *CreateMultiOrderShrinkRequestOrderItemsComponents) SetValue(v string) *CreateMultiOrderShrinkRequestOrderItemsComponents {
	s.Value = &v
	return s
}

func (s *CreateMultiOrderShrinkRequestOrderItemsComponents) Validate() error {
	return dara.Validate(s)
}
