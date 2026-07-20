// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleReShopItemListSubItemsShoppingItemMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetCabinQuantityList(v []*ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityList) *ModuleReShopItemListSubItemsShoppingItemMapValue
	GetCabinQuantityList() []*ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityList
	SetSearchPrice(v *ModuleReShopItemListSubItemsShoppingItemMapValueSearchPrice) *ModuleReShopItemListSubItemsShoppingItemMapValue
	GetSearchPrice() *ModuleReShopItemListSubItemsShoppingItemMapValueSearchPrice
	SetSegmentPriceList(v []*ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceList) *ModuleReShopItemListSubItemsShoppingItemMapValue
	GetSegmentPriceList() []*ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceList
}

type ModuleReShopItemListSubItemsShoppingItemMapValue struct {
	// The remaining cabin inventory for each segment.
	CabinQuantityList []*ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityList `json:"cabin_quantity_list,omitempty" xml:"cabin_quantity_list,omitempty" type:"Repeated"`
	// The rebooking search price.
	SearchPrice *ModuleReShopItemListSubItemsShoppingItemMapValueSearchPrice `json:"search_price,omitempty" xml:"search_price,omitempty" type:"Struct"`
	// The price for each segment. This field may not have a value because airline bundled products may not have prices broken down by segment.
	SegmentPriceList []*ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceList `json:"segment_price_list,omitempty" xml:"segment_price_list,omitempty" type:"Repeated"`
}

func (s ModuleReShopItemListSubItemsShoppingItemMapValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleReShopItemListSubItemsShoppingItemMapValue) GoString() string {
	return s.String()
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValue) GetCabinQuantityList() []*ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityList {
	return s.CabinQuantityList
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValue) GetSearchPrice() *ModuleReShopItemListSubItemsShoppingItemMapValueSearchPrice {
	return s.SearchPrice
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValue) GetSegmentPriceList() []*ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceList {
	return s.SegmentPriceList
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValue) SetCabinQuantityList(v []*ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityList) *ModuleReShopItemListSubItemsShoppingItemMapValue {
	s.CabinQuantityList = v
	return s
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValue) SetSearchPrice(v *ModuleReShopItemListSubItemsShoppingItemMapValueSearchPrice) *ModuleReShopItemListSubItemsShoppingItemMapValue {
	s.SearchPrice = v
	return s
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValue) SetSegmentPriceList(v []*ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceList) *ModuleReShopItemListSubItemsShoppingItemMapValue {
	s.SegmentPriceList = v
	return s
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValue) Validate() error {
	if s.CabinQuantityList != nil {
		for _, item := range s.CabinQuantityList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SearchPrice != nil {
		if err := s.SearchPrice.Validate(); err != nil {
			return err
		}
	}
	if s.SegmentPriceList != nil {
		for _, item := range s.SegmentPriceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityList struct {
	// The segment position information, indicating which journey and which segment within the overall itinerary.
	SegmentPosition *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition `json:"segment_position,omitempty" xml:"segment_position,omitempty" type:"Struct"`
	// The detailed cabin information.
	CabinInfo *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo `json:"cabin_info,omitempty" xml:"cabin_info,omitempty" type:"Struct"`
}

func (s ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityList) String() string {
	return dara.Prettify(s)
}

func (s ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityList) GoString() string {
	return s.String()
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityList) GetSegmentPosition() *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition {
	return s.SegmentPosition
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityList) GetCabinInfo() *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo {
	return s.CabinInfo
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityList) SetSegmentPosition(v *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition) *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityList {
	s.SegmentPosition = v
	return s
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityList) SetCabinInfo(v *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityList {
	s.CabinInfo = v
	return s
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityList) Validate() error {
	if s.SegmentPosition != nil {
		if err := s.SegmentPosition.Validate(); err != nil {
			return err
		}
	}
	if s.CabinInfo != nil {
		if err := s.CabinInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition struct {
	// The journey ordinal number, starting from 0.
	//
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// The segment ordinal number, starting from 0 within the same journey.
	//
	// example:
	//
	// 0
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
}

func (s ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition) String() string {
	return dara.Prettify(s)
}

func (s ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition) GoString() string {
	return s.String()
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition) SetJourneyIndex(v int32) *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition {
	s.JourneyIndex = &v
	return s
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition) SetSegmentIndex(v int32) *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition {
	s.SegmentIndex = &v
	return s
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition) Validate() error {
	return dara.Validate(s)
}

type ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo struct {
	// The cabin code.
	//
	// example:
	//
	// Y
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// The cabin class. Valid values:
	//
	// - F: First class.
	//
	// - C: Business class.
	//
	// - Y: Economy class.
	//
	// - P: Premium economy class.
	//
	// example:
	//
	// Y
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// The cabin class name.
	//
	// example:
	//
	// 经济舱
	CabinClassName *string `json:"cabin_class_name,omitempty" xml:"cabin_class_name,omitempty"`
	// The number of remaining seats in the cabin. Valid values: 0-9 (0 to 9 seats remaining) or A (more than 9 seats).
	//
	// example:
	//
	// 8
	Quantity *string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// The cabin class description.
	//
	// example:
	//
	// 经济舱
	CabinClassMemo *string `json:"cabin_class_memo,omitempty" xml:"cabin_class_memo,omitempty"`
	// The domestic special notes.
	//
	// example:
	//
	// 经济舱
	Specification *string `json:"specification,omitempty" xml:"specification,omitempty"`
}

func (s ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) String() string {
	return dara.Prettify(s)
}

func (s ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) GoString() string {
	return s.String()
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) GetCabin() *string {
	return s.Cabin
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) GetCabinClass() *string {
	return s.CabinClass
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) GetCabinClassName() *string {
	return s.CabinClassName
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) GetQuantity() *string {
	return s.Quantity
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) GetCabinClassMemo() *string {
	return s.CabinClassMemo
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) GetSpecification() *string {
	return s.Specification
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) SetCabin(v string) *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo {
	s.Cabin = &v
	return s
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) SetCabinClass(v string) *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo {
	s.CabinClass = &v
	return s
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) SetCabinClassName(v string) *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo {
	s.CabinClassName = &v
	return s
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) SetQuantity(v string) *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo {
	s.Quantity = &v
	return s
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) SetCabinClassMemo(v string) *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo {
	s.CabinClassMemo = &v
	return s
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) SetSpecification(v string) *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo {
	s.Specification = &v
	return s
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) Validate() error {
	return dara.Validate(s)
}

type ModuleReShopItemListSubItemsShoppingItemMapValueSearchPrice struct {
	// The total amount, in cents.
	//
	// example:
	//
	// 1000
	TotalAmount *int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// The service fee amount in cents.
	//
	// example:
	//
	// 100
	HandlingAmount *int64 `json:"handling_amount,omitempty" xml:"handling_amount,omitempty"`
	// The upgrade fee amount, in cents.
	//
	// example:
	//
	// 1000
	UpgradeAmount *int64 `json:"upgrade_amount,omitempty" xml:"upgrade_amount,omitempty"`
	// The tax difference amount, in cents.
	//
	// example:
	//
	// 100
	TaxDiffAmount *int64 `json:"tax_diff_amount,omitempty" xml:"tax_diff_amount,omitempty"`
	// Indicates whether a direct quote is available. Default value: true.
	//
	// example:
	//
	// true
	HasPrice *bool `json:"has_price,omitempty" xml:"has_price,omitempty"`
	// The text prompt displayed when no direct quote is available.
	//
	// example:
	//
	// 待服务商确认
	NonPriceText *string `json:"non_price_text,omitempty" xml:"non_price_text,omitempty"`
}

func (s ModuleReShopItemListSubItemsShoppingItemMapValueSearchPrice) String() string {
	return dara.Prettify(s)
}

func (s ModuleReShopItemListSubItemsShoppingItemMapValueSearchPrice) GoString() string {
	return s.String()
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSearchPrice) GetTotalAmount() *int64 {
	return s.TotalAmount
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSearchPrice) GetHandlingAmount() *int64 {
	return s.HandlingAmount
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSearchPrice) GetUpgradeAmount() *int64 {
	return s.UpgradeAmount
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSearchPrice) GetTaxDiffAmount() *int64 {
	return s.TaxDiffAmount
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSearchPrice) GetHasPrice() *bool {
	return s.HasPrice
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSearchPrice) GetNonPriceText() *string {
	return s.NonPriceText
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSearchPrice) SetTotalAmount(v int64) *ModuleReShopItemListSubItemsShoppingItemMapValueSearchPrice {
	s.TotalAmount = &v
	return s
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSearchPrice) SetHandlingAmount(v int64) *ModuleReShopItemListSubItemsShoppingItemMapValueSearchPrice {
	s.HandlingAmount = &v
	return s
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSearchPrice) SetUpgradeAmount(v int64) *ModuleReShopItemListSubItemsShoppingItemMapValueSearchPrice {
	s.UpgradeAmount = &v
	return s
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSearchPrice) SetTaxDiffAmount(v int64) *ModuleReShopItemListSubItemsShoppingItemMapValueSearchPrice {
	s.TaxDiffAmount = &v
	return s
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSearchPrice) SetHasPrice(v bool) *ModuleReShopItemListSubItemsShoppingItemMapValueSearchPrice {
	s.HasPrice = &v
	return s
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSearchPrice) SetNonPriceText(v string) *ModuleReShopItemListSubItemsShoppingItemMapValueSearchPrice {
	s.NonPriceText = &v
	return s
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSearchPrice) Validate() error {
	return dara.Validate(s)
}

type ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceList struct {
	// The segment position information, indicating which journey and which segment within the overall itinerary.
	SegmentPosition *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition `json:"segment_position,omitempty" xml:"segment_position,omitempty" type:"Struct"`
	// The rebooking quote.
	SearchPrice *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice `json:"search_price,omitempty" xml:"search_price,omitempty" type:"Struct"`
}

func (s ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceList) String() string {
	return dara.Prettify(s)
}

func (s ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceList) GoString() string {
	return s.String()
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceList) GetSegmentPosition() *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition {
	return s.SegmentPosition
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceList) GetSearchPrice() *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice {
	return s.SearchPrice
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceList) SetSegmentPosition(v *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition) *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceList {
	s.SegmentPosition = v
	return s
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceList) SetSearchPrice(v *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceList {
	s.SearchPrice = v
	return s
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceList) Validate() error {
	if s.SegmentPosition != nil {
		if err := s.SegmentPosition.Validate(); err != nil {
			return err
		}
	}
	if s.SearchPrice != nil {
		if err := s.SearchPrice.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition struct {
	// The journey ordinal number, starting from 0.
	//
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// The segment ordinal number, starting from 0 within the same journey.
	//
	// example:
	//
	// 0
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
}

func (s ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition) String() string {
	return dara.Prettify(s)
}

func (s ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition) GoString() string {
	return s.String()
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition) SetJourneyIndex(v int32) *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition {
	s.JourneyIndex = &v
	return s
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition) SetSegmentIndex(v int32) *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition {
	s.SegmentIndex = &v
	return s
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition) Validate() error {
	return dara.Validate(s)
}

type ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice struct {
	// The total amount in cents.
	//
	// example:
	//
	// 1000
	TotalAmount *int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// The service fee amount in cents.
	//
	// example:
	//
	// 1000
	HandlingAmount *int64 `json:"handling_amount,omitempty" xml:"handling_amount,omitempty"`
	// The upgrade fee amount in cents.
	//
	// example:
	//
	// 1000
	UpgradeAmount *int64 `json:"upgrade_amount,omitempty" xml:"upgrade_amount,omitempty"`
	// The tax difference amount in cents.
	//
	// example:
	//
	// 1000
	TaxDiffAmount *int64 `json:"tax_diff_amount,omitempty" xml:"tax_diff_amount,omitempty"`
	// Indicates whether a direct quote is available. Default value: true.
	//
	// example:
	//
	// 1000
	HasPrice *bool `json:"has_price,omitempty" xml:"has_price,omitempty"`
	// The text prompt displayed when no direct quote is available.
	//
	// example:
	//
	// 待服务商确认
	NonPriceText *string `json:"non_price_text,omitempty" xml:"non_price_text,omitempty"`
}

func (s ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) String() string {
	return dara.Prettify(s)
}

func (s ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) GoString() string {
	return s.String()
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) GetTotalAmount() *int64 {
	return s.TotalAmount
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) GetHandlingAmount() *int64 {
	return s.HandlingAmount
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) GetUpgradeAmount() *int64 {
	return s.UpgradeAmount
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) GetTaxDiffAmount() *int64 {
	return s.TaxDiffAmount
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) GetHasPrice() *bool {
	return s.HasPrice
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) GetNonPriceText() *string {
	return s.NonPriceText
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) SetTotalAmount(v int64) *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice {
	s.TotalAmount = &v
	return s
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) SetHandlingAmount(v int64) *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice {
	s.HandlingAmount = &v
	return s
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) SetUpgradeAmount(v int64) *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice {
	s.UpgradeAmount = &v
	return s
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) SetTaxDiffAmount(v int64) *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice {
	s.TaxDiffAmount = &v
	return s
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) SetHasPrice(v bool) *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice {
	s.HasPrice = &v
	return s
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) SetNonPriceText(v string) *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice {
	s.NonPriceText = &v
	return s
}

func (s *ModuleReShopItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) Validate() error {
	return dara.Validate(s)
}
