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
	ChannelCookie *string `json:"ChannelCookie,omitempty" xml:"ChannelCookie,omitempty"`
	// The items in the order.
	OrderItems []*CreateMultiOrderRequestOrderItems `json:"OrderItems,omitempty" xml:"OrderItems,omitempty" type:"Repeated"`
	// The order type.
	//
	// example:
	//
	// create
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The extended properties.
	Properties       map[string]*string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	ResellerOwnerUid *int64             `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
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
	Components  []*CreateMultiOrderRequestOrderItemsComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	InstanceIds []*string                                      `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
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
