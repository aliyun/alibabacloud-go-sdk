// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleReShopItemListBestPriceItemShoppingItemMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetCabinQuantityList(v []*ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityList) *ModuleReShopItemListBestPriceItemShoppingItemMapValue
	GetCabinQuantityList() []*ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityList
	SetSearchPrice(v *ModuleReShopItemListBestPriceItemShoppingItemMapValueSearchPrice) *ModuleReShopItemListBestPriceItemShoppingItemMapValue
	GetSearchPrice() *ModuleReShopItemListBestPriceItemShoppingItemMapValueSearchPrice
	SetSegmentPriceList(v []*ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceList) *ModuleReShopItemListBestPriceItemShoppingItemMapValue
	GetSegmentPriceList() []*ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceList
}

type ModuleReShopItemListBestPriceItemShoppingItemMapValue struct {
	CabinQuantityList []*ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityList `json:"cabin_quantity_list,omitempty" xml:"cabin_quantity_list,omitempty" type:"Repeated"`
	SearchPrice       *ModuleReShopItemListBestPriceItemShoppingItemMapValueSearchPrice         `json:"search_price,omitempty" xml:"search_price,omitempty" type:"Struct"`
	SegmentPriceList  []*ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceList  `json:"segment_price_list,omitempty" xml:"segment_price_list,omitempty" type:"Repeated"`
}

func (s ModuleReShopItemListBestPriceItemShoppingItemMapValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleReShopItemListBestPriceItemShoppingItemMapValue) GoString() string {
	return s.String()
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValue) GetCabinQuantityList() []*ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityList {
	return s.CabinQuantityList
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValue) GetSearchPrice() *ModuleReShopItemListBestPriceItemShoppingItemMapValueSearchPrice {
	return s.SearchPrice
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValue) GetSegmentPriceList() []*ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceList {
	return s.SegmentPriceList
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValue) SetCabinQuantityList(v []*ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityList) *ModuleReShopItemListBestPriceItemShoppingItemMapValue {
	s.CabinQuantityList = v
	return s
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValue) SetSearchPrice(v *ModuleReShopItemListBestPriceItemShoppingItemMapValueSearchPrice) *ModuleReShopItemListBestPriceItemShoppingItemMapValue {
	s.SearchPrice = v
	return s
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValue) SetSegmentPriceList(v []*ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceList) *ModuleReShopItemListBestPriceItemShoppingItemMapValue {
	s.SegmentPriceList = v
	return s
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValue) Validate() error {
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

type ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityList struct {
	SegmentPosition *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListSegmentPosition `json:"segment_position,omitempty" xml:"segment_position,omitempty" type:"Struct"`
	CabinInfo       *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo       `json:"cabin_info,omitempty" xml:"cabin_info,omitempty" type:"Struct"`
}

func (s ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityList) String() string {
	return dara.Prettify(s)
}

func (s ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityList) GoString() string {
	return s.String()
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityList) GetSegmentPosition() *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListSegmentPosition {
	return s.SegmentPosition
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityList) GetCabinInfo() *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo {
	return s.CabinInfo
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityList) SetSegmentPosition(v *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListSegmentPosition) *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityList {
	s.SegmentPosition = v
	return s
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityList) SetCabinInfo(v *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo) *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityList {
	s.CabinInfo = v
	return s
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityList) Validate() error {
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

type ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListSegmentPosition struct {
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 0
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
}

func (s ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListSegmentPosition) String() string {
	return dara.Prettify(s)
}

func (s ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListSegmentPosition) GoString() string {
	return s.String()
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListSegmentPosition) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListSegmentPosition) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListSegmentPosition) SetJourneyIndex(v int32) *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListSegmentPosition {
	s.JourneyIndex = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListSegmentPosition) SetSegmentIndex(v int32) *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListSegmentPosition {
	s.SegmentIndex = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListSegmentPosition) Validate() error {
	return dara.Validate(s)
}

type ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo struct {
	// example:
	//
	// Y
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// example:
	//
	// Y
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// example:
	//
	// 经济舱
	CabinClassName *string `json:"cabin_class_name,omitempty" xml:"cabin_class_name,omitempty"`
	// example:
	//
	// 8
	Quantity *string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// example:
	//
	// 经济舱
	CabinClassMemo *string `json:"cabin_class_memo,omitempty" xml:"cabin_class_memo,omitempty"`
	// example:
	//
	// 无
	Specification *string `json:"specification,omitempty" xml:"specification,omitempty"`
}

func (s ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo) String() string {
	return dara.Prettify(s)
}

func (s ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo) GoString() string {
	return s.String()
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo) GetCabin() *string {
	return s.Cabin
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo) GetCabinClass() *string {
	return s.CabinClass
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo) GetCabinClassName() *string {
	return s.CabinClassName
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo) GetQuantity() *string {
	return s.Quantity
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo) GetCabinClassMemo() *string {
	return s.CabinClassMemo
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo) GetSpecification() *string {
	return s.Specification
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo) SetCabin(v string) *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo {
	s.Cabin = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo) SetCabinClass(v string) *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo {
	s.CabinClass = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo) SetCabinClassName(v string) *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo {
	s.CabinClassName = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo) SetQuantity(v string) *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo {
	s.Quantity = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo) SetCabinClassMemo(v string) *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo {
	s.CabinClassMemo = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo) SetSpecification(v string) *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo {
	s.Specification = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo) Validate() error {
	return dara.Validate(s)
}

type ModuleReShopItemListBestPriceItemShoppingItemMapValueSearchPrice struct {
	// example:
	//
	// 1000
	TotalAmount *int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// example:
	//
	// 1000
	HandlingAmount *int64 `json:"handling_amount,omitempty" xml:"handling_amount,omitempty"`
	// example:
	//
	// 1000
	UpgradeAmount *int64 `json:"upgrade_amount,omitempty" xml:"upgrade_amount,omitempty"`
	// example:
	//
	// 1000
	TaxDiffAmount *int64 `json:"tax_diff_amount,omitempty" xml:"tax_diff_amount,omitempty"`
	// example:
	//
	// 1000
	HasPrice *bool `json:"has_price,omitempty" xml:"has_price,omitempty"`
	// example:
	//
	// 待服务商确认
	NonPriceText *string `json:"non_price_text,omitempty" xml:"non_price_text,omitempty"`
}

func (s ModuleReShopItemListBestPriceItemShoppingItemMapValueSearchPrice) String() string {
	return dara.Prettify(s)
}

func (s ModuleReShopItemListBestPriceItemShoppingItemMapValueSearchPrice) GoString() string {
	return s.String()
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSearchPrice) GetTotalAmount() *int64 {
	return s.TotalAmount
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSearchPrice) GetHandlingAmount() *int64 {
	return s.HandlingAmount
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSearchPrice) GetUpgradeAmount() *int64 {
	return s.UpgradeAmount
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSearchPrice) GetTaxDiffAmount() *int64 {
	return s.TaxDiffAmount
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSearchPrice) GetHasPrice() *bool {
	return s.HasPrice
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSearchPrice) GetNonPriceText() *string {
	return s.NonPriceText
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSearchPrice) SetTotalAmount(v int64) *ModuleReShopItemListBestPriceItemShoppingItemMapValueSearchPrice {
	s.TotalAmount = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSearchPrice) SetHandlingAmount(v int64) *ModuleReShopItemListBestPriceItemShoppingItemMapValueSearchPrice {
	s.HandlingAmount = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSearchPrice) SetUpgradeAmount(v int64) *ModuleReShopItemListBestPriceItemShoppingItemMapValueSearchPrice {
	s.UpgradeAmount = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSearchPrice) SetTaxDiffAmount(v int64) *ModuleReShopItemListBestPriceItemShoppingItemMapValueSearchPrice {
	s.TaxDiffAmount = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSearchPrice) SetHasPrice(v bool) *ModuleReShopItemListBestPriceItemShoppingItemMapValueSearchPrice {
	s.HasPrice = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSearchPrice) SetNonPriceText(v string) *ModuleReShopItemListBestPriceItemShoppingItemMapValueSearchPrice {
	s.NonPriceText = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSearchPrice) Validate() error {
	return dara.Validate(s)
}

type ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceList struct {
	SegmentPosition *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSegmentPosition `json:"segment_position,omitempty" xml:"segment_position,omitempty" type:"Struct"`
	SearchPrice     *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice     `json:"search_price,omitempty" xml:"search_price,omitempty" type:"Struct"`
}

func (s ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceList) String() string {
	return dara.Prettify(s)
}

func (s ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceList) GoString() string {
	return s.String()
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceList) GetSegmentPosition() *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSegmentPosition {
	return s.SegmentPosition
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceList) GetSearchPrice() *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice {
	return s.SearchPrice
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceList) SetSegmentPosition(v *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSegmentPosition) *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceList {
	s.SegmentPosition = v
	return s
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceList) SetSearchPrice(v *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice) *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceList {
	s.SearchPrice = v
	return s
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceList) Validate() error {
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

type ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSegmentPosition struct {
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 0
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
}

func (s ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSegmentPosition) String() string {
	return dara.Prettify(s)
}

func (s ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSegmentPosition) GoString() string {
	return s.String()
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSegmentPosition) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSegmentPosition) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSegmentPosition) SetJourneyIndex(v int32) *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSegmentPosition {
	s.JourneyIndex = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSegmentPosition) SetSegmentIndex(v int32) *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSegmentPosition {
	s.SegmentIndex = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSegmentPosition) Validate() error {
	return dara.Validate(s)
}

type ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice struct {
	// example:
	//
	// 1000
	TotalAmount *int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// example:
	//
	// 1000
	HandlingAmount *int64 `json:"handling_amount,omitempty" xml:"handling_amount,omitempty"`
	// example:
	//
	// 1000
	UpgradeAmount *int64 `json:"upgrade_amount,omitempty" xml:"upgrade_amount,omitempty"`
	// example:
	//
	// 1000
	TaxDiffAmount *int64 `json:"tax_diff_amount,omitempty" xml:"tax_diff_amount,omitempty"`
	// example:
	//
	// 1000
	HasPrice *bool `json:"has_price,omitempty" xml:"has_price,omitempty"`
	// example:
	//
	// 待服务商确认
	NonPriceText *string `json:"non_price_text,omitempty" xml:"non_price_text,omitempty"`
}

func (s ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice) String() string {
	return dara.Prettify(s)
}

func (s ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice) GoString() string {
	return s.String()
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice) GetTotalAmount() *int64 {
	return s.TotalAmount
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice) GetHandlingAmount() *int64 {
	return s.HandlingAmount
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice) GetUpgradeAmount() *int64 {
	return s.UpgradeAmount
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice) GetTaxDiffAmount() *int64 {
	return s.TaxDiffAmount
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice) GetHasPrice() *bool {
	return s.HasPrice
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice) GetNonPriceText() *string {
	return s.NonPriceText
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice) SetTotalAmount(v int64) *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice {
	s.TotalAmount = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice) SetHandlingAmount(v int64) *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice {
	s.HandlingAmount = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice) SetUpgradeAmount(v int64) *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice {
	s.UpgradeAmount = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice) SetTaxDiffAmount(v int64) *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice {
	s.TaxDiffAmount = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice) SetHasPrice(v bool) *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice {
	s.HasPrice = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice) SetNonPriceText(v string) *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice {
	s.NonPriceText = &v
	return s
}

func (s *ModuleReShopItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice) Validate() error {
	return dara.Validate(s)
}
