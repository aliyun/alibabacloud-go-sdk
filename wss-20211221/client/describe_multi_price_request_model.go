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
	// The product information.
	OrderItems []*DescribeMultiPriceRequestOrderItems `json:"OrderItems,omitempty" xml:"OrderItems,omitempty" type:"Repeated"`
	// The order type.
	//
	// example:
	//
	// create
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The package code. You do not need to specify this parameter for non-package types.
	//
	// example:
	//
	// pacakge
	PackageCode *string `json:"PackageCode,omitempty" xml:"PackageCode,omitempty"`
	// The user ID of the resource ownership in reseller pattern. You do not need to specify this parameter in non-reseller pattern.
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
	// The purchase quantity.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The list of product modules.
	Components []*DescribeMultiPriceRequestOrderItemsComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	Data       *string                                          `json:"Data,omitempty" xml:"Data,omitempty"`
	// The list of instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The subscription duration. Valid values:
	//
	// - If PeriodUnit is set to Year: 1, 2, or 3.
	//
	// - If PeriodUnit is set to Month: 1, 2, 3, or 6.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration.
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
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// The resource type.
	//
	// > This parameter is case-sensitive. Make sure that the value is spelled correctly.
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
	// The key of the module.
	//
	// example:
	//
	// RegionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the module.
	//
	// The following example values and valid values are for the Enterprise Edition monthly duration package:
	//
	// - RegionId: cn-shanghai
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
	//    - cloud_efficiency: ultra cloud disk
	//
	//    - cloud_auto: ESSD AutoPL cloud disk
	//
	//    - cloud_essd: enhanced standard SSD. Only specific instance types support this value.
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
	// - DataDiskSize (GiB): same as RootDiskSize
	//
	// - DataDiskCategory: same as RootDiskCategory
	//
	// - DataPerformanceLevel: same as RootPerformanceLevel
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
