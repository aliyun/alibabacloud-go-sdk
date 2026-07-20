// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetCabinQuantityList(v []*ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityList) *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValue
	GetCabinQuantityList() []*ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityList
	SetSearchPrice(v *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValue
	GetSearchPrice() *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice
	SetSegmentPriceList(v []*ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceList) *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValue
	GetSegmentPriceList() []*ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceList
}

type ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValue struct {
	// The cabin remaining inventory per segment.
	CabinQuantityList []*ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityList `json:"cabin_quantity_list,omitempty" xml:"cabin_quantity_list,omitempty" type:"Repeated"`
	// The rebooking search price.
	SearchPrice *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice `json:"search_price,omitempty" xml:"search_price,omitempty" type:"Struct"`
	// The price per segment. This field may not have a value because airline bundled products may not have prices split by segment.
	SegmentPriceList []*ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceList `json:"segment_price_list,omitempty" xml:"segment_price_list,omitempty" type:"Repeated"`
}

func (s ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValue) GoString() string {
	return s.String()
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValue) GetCabinQuantityList() []*ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityList {
	return s.CabinQuantityList
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValue) GetSearchPrice() *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice {
	return s.SearchPrice
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValue) GetSegmentPriceList() []*ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceList {
	return s.SegmentPriceList
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValue) SetCabinQuantityList(v []*ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityList) *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValue {
	s.CabinQuantityList = v
	return s
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValue) SetSearchPrice(v *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValue {
	s.SearchPrice = v
	return s
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValue) SetSegmentPriceList(v []*ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceList) *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValue {
	s.SegmentPriceList = v
	return s
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValue) Validate() error {
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

type ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityList struct {
	// The segment position information, indicating the journey index and segment index within the overall itinerary.
	SegmentPosition *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition `json:"segment_position,omitempty" xml:"segment_position,omitempty" type:"Struct"`
	// The cabin details.
	CabinInfo *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListCabinInfo `json:"cabin_info,omitempty" xml:"cabin_info,omitempty" type:"Struct"`
}

func (s ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityList) String() string {
	return dara.Prettify(s)
}

func (s ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityList) GoString() string {
	return s.String()
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityList) GetSegmentPosition() *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition {
	return s.SegmentPosition
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityList) GetCabinInfo() *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListCabinInfo {
	return s.CabinInfo
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityList) SetSegmentPosition(v *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition) *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityList {
	s.SegmentPosition = v
	return s
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityList) SetCabinInfo(v *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityList {
	s.CabinInfo = v
	return s
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityList) Validate() error {
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

type ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition struct {
	// The journey ordinal number (starting from 0).
	//
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// The segment ordinal number (starting from 0 within the same journey).
	//
	// example:
	//
	// 0
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
}

func (s ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition) String() string {
	return dara.Prettify(s)
}

func (s ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition) GoString() string {
	return s.String()
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition) SetJourneyIndex(v int32) *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition {
	s.JourneyIndex = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition) SetSegmentIndex(v int32) *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition {
	s.SegmentIndex = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition) Validate() error {
	return dara.Validate(s)
}

type ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListCabinInfo struct {
	// The cabin code.
	//
	// example:
	//
	// Y
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// The cabin class. Valid values: F: first class. C: business class. Y: economy class. P: premium economy class.
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
	// The number of remaining seats in the cabin. 0-9: 0 to 9 seats remaining. A: more than 9 seats remaining.
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
	// 无
	Specification *string `json:"specification,omitempty" xml:"specification,omitempty"`
}

func (s ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) String() string {
	return dara.Prettify(s)
}

func (s ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) GoString() string {
	return s.String()
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) GetCabin() *string {
	return s.Cabin
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) GetCabinClass() *string {
	return s.CabinClass
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) GetCabinClassName() *string {
	return s.CabinClassName
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) GetQuantity() *string {
	return s.Quantity
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) GetCabinClassMemo() *string {
	return s.CabinClassMemo
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) GetSpecification() *string {
	return s.Specification
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) SetCabin(v string) *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListCabinInfo {
	s.Cabin = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) SetCabinClass(v string) *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListCabinInfo {
	s.CabinClass = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) SetCabinClassName(v string) *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListCabinInfo {
	s.CabinClassName = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) SetQuantity(v string) *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListCabinInfo {
	s.Quantity = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) SetCabinClassMemo(v string) *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListCabinInfo {
	s.CabinClassMemo = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) SetSpecification(v string) *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListCabinInfo {
	s.Specification = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) Validate() error {
	return dara.Validate(s)
}

type ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice struct {
	// The total amount, in cents.
	//
	// example:
	//
	// 1000
	TotalAmount *int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// The service fee amount, in cents.
	//
	// example:
	//
	// 1000
	HandlingAmount *int64 `json:"handling_amount,omitempty" xml:"handling_amount,omitempty"`
	// The cabin upgrade fee amount, in cents.
	//
	// example:
	//
	// 1000
	UpgradeAmount *int64 `json:"upgrade_amount,omitempty" xml:"upgrade_amount,omitempty"`
	// The tax difference amount, in cents.
	//
	// example:
	//
	// 1000
	TaxDiffAmount *int64 `json:"tax_diff_amount,omitempty" xml:"tax_diff_amount,omitempty"`
	// Indicates whether a direct price is available. Default value: true.
	//
	// example:
	//
	// true
	HasPrice *bool `json:"has_price,omitempty" xml:"has_price,omitempty"`
	// The text prompt displayed when no direct price is available.
	//
	// example:
	//
	// 待服务商确认
	NonPriceText *string `json:"non_price_text,omitempty" xml:"non_price_text,omitempty"`
}

func (s ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) String() string {
	return dara.Prettify(s)
}

func (s ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) GoString() string {
	return s.String()
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) GetTotalAmount() *int64 {
	return s.TotalAmount
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) GetHandlingAmount() *int64 {
	return s.HandlingAmount
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) GetUpgradeAmount() *int64 {
	return s.UpgradeAmount
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) GetTaxDiffAmount() *int64 {
	return s.TaxDiffAmount
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) GetHasPrice() *bool {
	return s.HasPrice
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) GetNonPriceText() *string {
	return s.NonPriceText
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) SetTotalAmount(v int64) *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice {
	s.TotalAmount = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) SetHandlingAmount(v int64) *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice {
	s.HandlingAmount = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) SetUpgradeAmount(v int64) *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice {
	s.UpgradeAmount = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) SetTaxDiffAmount(v int64) *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice {
	s.TaxDiffAmount = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) SetHasPrice(v bool) *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice {
	s.HasPrice = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) SetNonPriceText(v string) *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice {
	s.NonPriceText = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) Validate() error {
	return dara.Validate(s)
}

type ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceList struct {
	// The segment position information, indicating the journey index and segment index within the overall itinerary.
	SegmentPosition *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition `json:"segment_position,omitempty" xml:"segment_position,omitempty" type:"Struct"`
	// The rebooking price.
	SearchPrice *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSearchPrice `json:"search_price,omitempty" xml:"search_price,omitempty" type:"Struct"`
}

func (s ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceList) String() string {
	return dara.Prettify(s)
}

func (s ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceList) GoString() string {
	return s.String()
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceList) GetSegmentPosition() *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition {
	return s.SegmentPosition
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceList) GetSearchPrice() *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSearchPrice {
	return s.SearchPrice
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceList) SetSegmentPosition(v *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition) *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceList {
	s.SegmentPosition = v
	return s
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceList) SetSearchPrice(v *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceList {
	s.SearchPrice = v
	return s
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceList) Validate() error {
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

type ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition struct {
	// The journey ordinal number (starting from 0).
	//
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// The segment ordinal number (starting from 0 within the same journey).
	//
	// example:
	//
	// 0
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
}

func (s ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition) String() string {
	return dara.Prettify(s)
}

func (s ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition) GoString() string {
	return s.String()
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition) SetJourneyIndex(v int32) *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition {
	s.JourneyIndex = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition) SetSegmentIndex(v int32) *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition {
	s.SegmentIndex = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition) Validate() error {
	return dara.Validate(s)
}

type ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSearchPrice struct {
	// The total amount, in cents.
	//
	// example:
	//
	// 1000
	TotalAmount *int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// The service fee amount, in cents.
	//
	// example:
	//
	// 1000
	HandlingAmount *int64 `json:"handling_amount,omitempty" xml:"handling_amount,omitempty"`
	// The cabin upgrade fee amount, in cents.
	//
	// example:
	//
	// 1000
	UpgradeAmount *int64 `json:"upgrade_amount,omitempty" xml:"upgrade_amount,omitempty"`
	// The tax difference amount, in cents.
	//
	// example:
	//
	// 1000
	TaxDiffAmount *int64 `json:"tax_diff_amount,omitempty" xml:"tax_diff_amount,omitempty"`
	// Indicates whether a direct price is available. Default value: true.
	//
	// example:
	//
	// true
	HasPrice *bool `json:"has_price,omitempty" xml:"has_price,omitempty"`
	// The text prompt displayed when no direct price is available.
	//
	// example:
	//
	// 待服务商确认
	NonPriceText *string `json:"non_price_text,omitempty" xml:"non_price_text,omitempty"`
}

func (s ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) String() string {
	return dara.Prettify(s)
}

func (s ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) GoString() string {
	return s.String()
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) GetTotalAmount() *int64 {
	return s.TotalAmount
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) GetHandlingAmount() *int64 {
	return s.HandlingAmount
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) GetUpgradeAmount() *int64 {
	return s.UpgradeAmount
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) GetTaxDiffAmount() *int64 {
	return s.TaxDiffAmount
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) GetHasPrice() *bool {
	return s.HasPrice
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) GetNonPriceText() *string {
	return s.NonPriceText
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) SetTotalAmount(v int64) *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSearchPrice {
	s.TotalAmount = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) SetHandlingAmount(v int64) *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSearchPrice {
	s.HandlingAmount = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) SetUpgradeAmount(v int64) *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSearchPrice {
	s.UpgradeAmount = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) SetTaxDiffAmount(v int64) *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSearchPrice {
	s.TaxDiffAmount = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) SetHasPrice(v bool) *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSearchPrice {
	s.HasPrice = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) SetNonPriceText(v string) *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSearchPrice {
	s.NonPriceText = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) Validate() error {
	return dara.Validate(s)
}
