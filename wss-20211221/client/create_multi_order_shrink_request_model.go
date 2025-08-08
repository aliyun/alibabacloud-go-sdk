// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultiOrderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
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
	OrderItems []*CreateMultiOrderShrinkRequestOrderItems `json:"OrderItems,omitempty" xml:"OrderItems,omitempty" type:"Repeated"`
	// example:
	//
	// create
	OrderType        *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	PropertiesShrink *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	ResellerOwnerUid *int64  `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
}

func (s CreateMultiOrderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiOrderShrinkRequest) GoString() string {
	return s.String()
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
	return dara.Validate(s)
}

type CreateMultiOrderShrinkRequestOrderItems struct {
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// example:
	//
	// false
	AutoRenew  *bool                                                `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	BuyChange  *bool                                                `json:"BuyChange,omitempty" xml:"BuyChange,omitempty"`
	Components []*CreateMultiOrderShrinkRequestOrderItemsComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// Year
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	PromotionId *string   `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
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
	return dara.Validate(s)
}

type CreateMultiOrderShrinkRequestOrderItemsComponents struct {
	// example:
	//
	// RegionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
