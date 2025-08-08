// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultiOrderRequest interface {
	dara.Model
	String() string
	GoString() string
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
	OrderItems []*CreateMultiOrderRequestOrderItems `json:"OrderItems,omitempty" xml:"OrderItems,omitempty" type:"Repeated"`
	// example:
	//
	// create
	OrderType        *string            `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	Properties       map[string]*string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	ResellerOwnerUid *int64             `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
}

func (s CreateMultiOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiOrderRequest) GoString() string {
	return s.String()
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
	return dara.Validate(s)
}

type CreateMultiOrderRequestOrderItems struct {
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
	AutoRenew  *bool                                          `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	BuyChange  *bool                                          `json:"BuyChange,omitempty" xml:"BuyChange,omitempty"`
	Components []*CreateMultiOrderRequestOrderItemsComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
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
	return dara.Validate(s)
}

type CreateMultiOrderRequestOrderItemsComponents struct {
	// example:
	//
	// RegionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
