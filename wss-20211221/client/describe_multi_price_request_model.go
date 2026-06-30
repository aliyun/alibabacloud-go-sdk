// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMultiPriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderItems(v []*DescribeMultiPriceRequestOrderItems) *DescribeMultiPriceRequest
	GetOrderItems() []*DescribeMultiPriceRequestOrderItems
	SetOrderType(v string) *DescribeMultiPriceRequest
	GetOrderType() *string
	SetPackageCode(v string) *DescribeMultiPriceRequest
	GetPackageCode() *string
	SetResellerOwnerUid(v int64) *DescribeMultiPriceRequest
	GetResellerOwnerUid() *int64
}

type DescribeMultiPriceRequest struct {
	// The order items.
	OrderItems []*DescribeMultiPriceRequestOrderItems `json:"OrderItems,omitempty" xml:"OrderItems,omitempty" type:"Repeated"`
	// The order type.
	//
	// example:
	//
	// create
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The package code. This parameter is not required for non-package types.
	//
	// example:
	//
	// pacakge
	PackageCode *string `json:"PackageCode,omitempty" xml:"PackageCode,omitempty"`
	// The ID of the user who owns the resource in the reseller model. This parameter is not required in non-reseller mode.
	//
	// example:
	//
	// 182864463481****
	ResellerOwnerUid *int64 `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
}

func (s DescribeMultiPriceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiPriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeMultiPriceRequest) GetOrderItems() []*DescribeMultiPriceRequestOrderItems {
	return s.OrderItems
}

func (s *DescribeMultiPriceRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribeMultiPriceRequest) GetPackageCode() *string {
	return s.PackageCode
}

func (s *DescribeMultiPriceRequest) GetResellerOwnerUid() *int64 {
	return s.ResellerOwnerUid
}

func (s *DescribeMultiPriceRequest) SetOrderItems(v []*DescribeMultiPriceRequestOrderItems) *DescribeMultiPriceRequest {
	s.OrderItems = v
	return s
}

func (s *DescribeMultiPriceRequest) SetOrderType(v string) *DescribeMultiPriceRequest {
	s.OrderType = &v
	return s
}

func (s *DescribeMultiPriceRequest) SetPackageCode(v string) *DescribeMultiPriceRequest {
	s.PackageCode = &v
	return s
}

func (s *DescribeMultiPriceRequest) SetResellerOwnerUid(v int64) *DescribeMultiPriceRequest {
	s.ResellerOwnerUid = &v
	return s
}

func (s *DescribeMultiPriceRequest) Validate() error {
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

type DescribeMultiPriceRequestOrderItems struct {
	// The quantity to purchase.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// A list of components.
	Components []*DescribeMultiPriceRequestOrderItemsComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	Data       *string                                          `json:"Data,omitempty" xml:"Data,omitempty"`
	// A list of instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The subscription period. Valid values:
	//
	// - If `PeriodUnit` is `Year`, the valid values are 1, 2, and 3.
	//
	// - If `PeriodUnit` is `Month`, the valid values are 1, 2, 3, and 6.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription period.
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
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// The resource type.
	//
	// > The value is case-sensitive.
	//
	// example:
	//
	// DurationPackage
	ResourceType     *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SavingPlanPeriod *string `json:"SavingPlanPeriod,omitempty" xml:"SavingPlanPeriod,omitempty"`
}

func (s DescribeMultiPriceRequestOrderItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiPriceRequestOrderItems) GoString() string {
	return s.String()
}

func (s *DescribeMultiPriceRequestOrderItems) GetAmount() *int32 {
	return s.Amount
}

func (s *DescribeMultiPriceRequestOrderItems) GetComponents() []*DescribeMultiPriceRequestOrderItemsComponents {
	return s.Components
}

func (s *DescribeMultiPriceRequestOrderItems) GetData() *string {
	return s.Data
}

func (s *DescribeMultiPriceRequestOrderItems) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeMultiPriceRequestOrderItems) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeMultiPriceRequestOrderItems) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *DescribeMultiPriceRequestOrderItems) GetPromotionId() *string {
	return s.PromotionId
}

func (s *DescribeMultiPriceRequestOrderItems) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *DescribeMultiPriceRequestOrderItems) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeMultiPriceRequestOrderItems) GetSavingPlanPeriod() *string {
	return s.SavingPlanPeriod
}

func (s *DescribeMultiPriceRequestOrderItems) SetAmount(v int32) *DescribeMultiPriceRequestOrderItems {
	s.Amount = &v
	return s
}

func (s *DescribeMultiPriceRequestOrderItems) SetComponents(v []*DescribeMultiPriceRequestOrderItemsComponents) *DescribeMultiPriceRequestOrderItems {
	s.Components = v
	return s
}

func (s *DescribeMultiPriceRequestOrderItems) SetData(v string) *DescribeMultiPriceRequestOrderItems {
	s.Data = &v
	return s
}

func (s *DescribeMultiPriceRequestOrderItems) SetInstanceIds(v []*string) *DescribeMultiPriceRequestOrderItems {
	s.InstanceIds = v
	return s
}

func (s *DescribeMultiPriceRequestOrderItems) SetPeriod(v int32) *DescribeMultiPriceRequestOrderItems {
	s.Period = &v
	return s
}

func (s *DescribeMultiPriceRequestOrderItems) SetPeriodUnit(v string) *DescribeMultiPriceRequestOrderItems {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeMultiPriceRequestOrderItems) SetPromotionId(v string) *DescribeMultiPriceRequestOrderItems {
	s.PromotionId = &v
	return s
}

func (s *DescribeMultiPriceRequestOrderItems) SetResourceIds(v []*string) *DescribeMultiPriceRequestOrderItems {
	s.ResourceIds = v
	return s
}

func (s *DescribeMultiPriceRequestOrderItems) SetResourceType(v string) *DescribeMultiPriceRequestOrderItems {
	s.ResourceType = &v
	return s
}

func (s *DescribeMultiPriceRequestOrderItems) SetSavingPlanPeriod(v string) *DescribeMultiPriceRequestOrderItems {
	s.SavingPlanPeriod = &v
	return s
}

func (s *DescribeMultiPriceRequestOrderItems) Validate() error {
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

type DescribeMultiPriceRequestOrderItemsComponents struct {
	// The key of the component.
	//
	// example:
	//
	// RegionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the component.
	//
	// The following are the keys and their sample or enumerated values for a monthly duration package of the enterprise edition:
	//
	// - `RegionId`: cn-shanghai
	//
	// - `InstanceType`: eds.enterprise_office.4c8g
	//
	// - `DurationType` (in hours): [enum]
	//
	//   - 120
	//
	//   - 250
	//
	// - `OsType`: [enum]
	//
	//   - Windows
	//
	//   - Linux
	//
	// - `RootDiskSize` (in GiB): 80
	//
	// - `RootDiskCategory`: [enum]
	//
	//   - `cloud_efficiency` (Ultra Cloud Disk)
	//
	//   - `cloud_auto` (AutoPL Cloud Disk)
	//
	//   - `cloud_essd` (Enhanced SSD (ESSD), available only for specific instance types)
	//
	// - `RootPerformanceLevel`: [enum]
	//
	//   - PL0
	//
	//   - PL1
	//
	//   - PL2
	//
	//   - PL3
	//
	// - `DataDiskSize` (in GiB): Same as `RootDiskSize`.
	//
	// - `DataDiskCategory`: Same as `RootDiskCategory`.
	//
	// - `DataPerformanceLevel`: Same as `RootPerformanceLevel`.
	//
	// example:
	//
	// cn-shanghai
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeMultiPriceRequestOrderItemsComponents) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiPriceRequestOrderItemsComponents) GoString() string {
	return s.String()
}

func (s *DescribeMultiPriceRequestOrderItemsComponents) GetKey() *string {
	return s.Key
}

func (s *DescribeMultiPriceRequestOrderItemsComponents) GetValue() *string {
	return s.Value
}

func (s *DescribeMultiPriceRequestOrderItemsComponents) SetKey(v string) *DescribeMultiPriceRequestOrderItemsComponents {
	s.Key = &v
	return s
}

func (s *DescribeMultiPriceRequestOrderItemsComponents) SetValue(v string) *DescribeMultiPriceRequestOrderItemsComponents {
	s.Value = &v
	return s
}

func (s *DescribeMultiPriceRequestOrderItemsComponents) Validate() error {
	return dara.Validate(s)
}
