// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleReShopItemListShoppingItemMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetCabinQuantityList(v []*ModuleReShopItemListShoppingItemMapValueCabinQuantityList) *ModuleReShopItemListShoppingItemMapValue
	GetCabinQuantityList() []*ModuleReShopItemListShoppingItemMapValueCabinQuantityList
	SetSearchPrice(v *ModuleReShopItemListShoppingItemMapValueSearchPrice) *ModuleReShopItemListShoppingItemMapValue
	GetSearchPrice() *ModuleReShopItemListShoppingItemMapValueSearchPrice
	SetSegmentPriceList(v []*ModuleReShopItemListShoppingItemMapValueSegmentPriceList) *ModuleReShopItemListShoppingItemMapValue
	GetSegmentPriceList() []*ModuleReShopItemListShoppingItemMapValueSegmentPriceList
}

type ModuleReShopItemListShoppingItemMapValue struct {
	CabinQuantityList []*ModuleReShopItemListShoppingItemMapValueCabinQuantityList `json:"cabin_quantity_list,omitempty" xml:"cabin_quantity_list,omitempty" type:"Repeated"`
	SearchPrice       *ModuleReShopItemListShoppingItemMapValueSearchPrice         `json:"search_price,omitempty" xml:"search_price,omitempty" type:"Struct"`
	SegmentPriceList  []*ModuleReShopItemListShoppingItemMapValueSegmentPriceList  `json:"segment_price_list,omitempty" xml:"segment_price_list,omitempty" type:"Repeated"`
}

func (s ModuleReShopItemListShoppingItemMapValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleReShopItemListShoppingItemMapValue) GoString() string {
	return s.String()
}

func (s *ModuleReShopItemListShoppingItemMapValue) GetCabinQuantityList() []*ModuleReShopItemListShoppingItemMapValueCabinQuantityList {
	return s.CabinQuantityList
}

func (s *ModuleReShopItemListShoppingItemMapValue) GetSearchPrice() *ModuleReShopItemListShoppingItemMapValueSearchPrice {
	return s.SearchPrice
}

func (s *ModuleReShopItemListShoppingItemMapValue) GetSegmentPriceList() []*ModuleReShopItemListShoppingItemMapValueSegmentPriceList {
	return s.SegmentPriceList
}

func (s *ModuleReShopItemListShoppingItemMapValue) SetCabinQuantityList(v []*ModuleReShopItemListShoppingItemMapValueCabinQuantityList) *ModuleReShopItemListShoppingItemMapValue {
	s.CabinQuantityList = v
	return s
}

func (s *ModuleReShopItemListShoppingItemMapValue) SetSearchPrice(v *ModuleReShopItemListShoppingItemMapValueSearchPrice) *ModuleReShopItemListShoppingItemMapValue {
	s.SearchPrice = v
	return s
}

func (s *ModuleReShopItemListShoppingItemMapValue) SetSegmentPriceList(v []*ModuleReShopItemListShoppingItemMapValueSegmentPriceList) *ModuleReShopItemListShoppingItemMapValue {
	s.SegmentPriceList = v
	return s
}

func (s *ModuleReShopItemListShoppingItemMapValue) Validate() error {
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

type ModuleReShopItemListShoppingItemMapValueCabinQuantityList struct {
	SegmentPosition *ModuleReShopItemListShoppingItemMapValueCabinQuantityListSegmentPosition `json:"segment_position,omitempty" xml:"segment_position,omitempty" type:"Struct"`
	CabinInfo       *ModuleReShopItemListShoppingItemMapValueCabinQuantityListCabinInfo       `json:"cabin_info,omitempty" xml:"cabin_info,omitempty" type:"Struct"`
}

func (s ModuleReShopItemListShoppingItemMapValueCabinQuantityList) String() string {
	return dara.Prettify(s)
}

func (s ModuleReShopItemListShoppingItemMapValueCabinQuantityList) GoString() string {
	return s.String()
}

func (s *ModuleReShopItemListShoppingItemMapValueCabinQuantityList) GetSegmentPosition() *ModuleReShopItemListShoppingItemMapValueCabinQuantityListSegmentPosition {
	return s.SegmentPosition
}

func (s *ModuleReShopItemListShoppingItemMapValueCabinQuantityList) GetCabinInfo() *ModuleReShopItemListShoppingItemMapValueCabinQuantityListCabinInfo {
	return s.CabinInfo
}

func (s *ModuleReShopItemListShoppingItemMapValueCabinQuantityList) SetSegmentPosition(v *ModuleReShopItemListShoppingItemMapValueCabinQuantityListSegmentPosition) *ModuleReShopItemListShoppingItemMapValueCabinQuantityList {
	s.SegmentPosition = v
	return s
}

func (s *ModuleReShopItemListShoppingItemMapValueCabinQuantityList) SetCabinInfo(v *ModuleReShopItemListShoppingItemMapValueCabinQuantityListCabinInfo) *ModuleReShopItemListShoppingItemMapValueCabinQuantityList {
	s.CabinInfo = v
	return s
}

func (s *ModuleReShopItemListShoppingItemMapValueCabinQuantityList) Validate() error {
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

type ModuleReShopItemListShoppingItemMapValueCabinQuantityListSegmentPosition struct {
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
}

func (s ModuleReShopItemListShoppingItemMapValueCabinQuantityListSegmentPosition) String() string {
	return dara.Prettify(s)
}

func (s ModuleReShopItemListShoppingItemMapValueCabinQuantityListSegmentPosition) GoString() string {
	return s.String()
}

func (s *ModuleReShopItemListShoppingItemMapValueCabinQuantityListSegmentPosition) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *ModuleReShopItemListShoppingItemMapValueCabinQuantityListSegmentPosition) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *ModuleReShopItemListShoppingItemMapValueCabinQuantityListSegmentPosition) SetJourneyIndex(v int32) *ModuleReShopItemListShoppingItemMapValueCabinQuantityListSegmentPosition {
	s.JourneyIndex = &v
	return s
}

func (s *ModuleReShopItemListShoppingItemMapValueCabinQuantityListSegmentPosition) SetSegmentIndex(v int32) *ModuleReShopItemListShoppingItemMapValueCabinQuantityListSegmentPosition {
	s.SegmentIndex = &v
	return s
}

func (s *ModuleReShopItemListShoppingItemMapValueCabinQuantityListSegmentPosition) Validate() error {
	return dara.Validate(s)
}

type ModuleReShopItemListShoppingItemMapValueCabinQuantityListCabinInfo struct {
	Cabin          *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	CabinClass     *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	CabinClassName *string `json:"cabin_class_name,omitempty" xml:"cabin_class_name,omitempty"`
	Quantity       *string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	CabinClassMemo *string `json:"cabin_class_memo,omitempty" xml:"cabin_class_memo,omitempty"`
	Specification  *string `json:"specification,omitempty" xml:"specification,omitempty"`
}

func (s ModuleReShopItemListShoppingItemMapValueCabinQuantityListCabinInfo) String() string {
	return dara.Prettify(s)
}

func (s ModuleReShopItemListShoppingItemMapValueCabinQuantityListCabinInfo) GoString() string {
	return s.String()
}

func (s *ModuleReShopItemListShoppingItemMapValueCabinQuantityListCabinInfo) GetCabin() *string {
	return s.Cabin
}

func (s *ModuleReShopItemListShoppingItemMapValueCabinQuantityListCabinInfo) GetCabinClass() *string {
	return s.CabinClass
}

func (s *ModuleReShopItemListShoppingItemMapValueCabinQuantityListCabinInfo) GetCabinClassName() *string {
	return s.CabinClassName
}

func (s *ModuleReShopItemListShoppingItemMapValueCabinQuantityListCabinInfo) GetQuantity() *string {
	return s.Quantity
}

func (s *ModuleReShopItemListShoppingItemMapValueCabinQuantityListCabinInfo) GetCabinClassMemo() *string {
	return s.CabinClassMemo
}

func (s *ModuleReShopItemListShoppingItemMapValueCabinQuantityListCabinInfo) GetSpecification() *string {
	return s.Specification
}

func (s *ModuleReShopItemListShoppingItemMapValueCabinQuantityListCabinInfo) SetCabin(v string) *ModuleReShopItemListShoppingItemMapValueCabinQuantityListCabinInfo {
	s.Cabin = &v
	return s
}

func (s *ModuleReShopItemListShoppingItemMapValueCabinQuantityListCabinInfo) SetCabinClass(v string) *ModuleReShopItemListShoppingItemMapValueCabinQuantityListCabinInfo {
	s.CabinClass = &v
	return s
}

func (s *ModuleReShopItemListShoppingItemMapValueCabinQuantityListCabinInfo) SetCabinClassName(v string) *ModuleReShopItemListShoppingItemMapValueCabinQuantityListCabinInfo {
	s.CabinClassName = &v
	return s
}

func (s *ModuleReShopItemListShoppingItemMapValueCabinQuantityListCabinInfo) SetQuantity(v string) *ModuleReShopItemListShoppingItemMapValueCabinQuantityListCabinInfo {
	s.Quantity = &v
	return s
}

func (s *ModuleReShopItemListShoppingItemMapValueCabinQuantityListCabinInfo) SetCabinClassMemo(v string) *ModuleReShopItemListShoppingItemMapValueCabinQuantityListCabinInfo {
	s.CabinClassMemo = &v
	return s
}

func (s *ModuleReShopItemListShoppingItemMapValueCabinQuantityListCabinInfo) SetSpecification(v string) *ModuleReShopItemListShoppingItemMapValueCabinQuantityListCabinInfo {
	s.Specification = &v
	return s
}

func (s *ModuleReShopItemListShoppingItemMapValueCabinQuantityListCabinInfo) Validate() error {
	return dara.Validate(s)
}

type ModuleReShopItemListShoppingItemMapValueSearchPrice struct {
	TotalAmount    *int64  `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	HandlingAmount *int64  `json:"handling_amount,omitempty" xml:"handling_amount,omitempty"`
	UpgradeAmount  *int64  `json:"upgrade_amount,omitempty" xml:"upgrade_amount,omitempty"`
	TaxDiffAmount  *int64  `json:"tax_diff_amount,omitempty" xml:"tax_diff_amount,omitempty"`
	HasPrice       *bool   `json:"has_price,omitempty" xml:"has_price,omitempty"`
	NonPriceText   *string `json:"non_price_text,omitempty" xml:"non_price_text,omitempty"`
}

func (s ModuleReShopItemListShoppingItemMapValueSearchPrice) String() string {
	return dara.Prettify(s)
}

func (s ModuleReShopItemListShoppingItemMapValueSearchPrice) GoString() string {
	return s.String()
}

func (s *ModuleReShopItemListShoppingItemMapValueSearchPrice) GetTotalAmount() *int64 {
	return s.TotalAmount
}

func (s *ModuleReShopItemListShoppingItemMapValueSearchPrice) GetHandlingAmount() *int64 {
	return s.HandlingAmount
}

func (s *ModuleReShopItemListShoppingItemMapValueSearchPrice) GetUpgradeAmount() *int64 {
	return s.UpgradeAmount
}

func (s *ModuleReShopItemListShoppingItemMapValueSearchPrice) GetTaxDiffAmount() *int64 {
	return s.TaxDiffAmount
}

func (s *ModuleReShopItemListShoppingItemMapValueSearchPrice) GetHasPrice() *bool {
	return s.HasPrice
}

func (s *ModuleReShopItemListShoppingItemMapValueSearchPrice) GetNonPriceText() *string {
	return s.NonPriceText
}

func (s *ModuleReShopItemListShoppingItemMapValueSearchPrice) SetTotalAmount(v int64) *ModuleReShopItemListShoppingItemMapValueSearchPrice {
	s.TotalAmount = &v
	return s
}

func (s *ModuleReShopItemListShoppingItemMapValueSearchPrice) SetHandlingAmount(v int64) *ModuleReShopItemListShoppingItemMapValueSearchPrice {
	s.HandlingAmount = &v
	return s
}

func (s *ModuleReShopItemListShoppingItemMapValueSearchPrice) SetUpgradeAmount(v int64) *ModuleReShopItemListShoppingItemMapValueSearchPrice {
	s.UpgradeAmount = &v
	return s
}

func (s *ModuleReShopItemListShoppingItemMapValueSearchPrice) SetTaxDiffAmount(v int64) *ModuleReShopItemListShoppingItemMapValueSearchPrice {
	s.TaxDiffAmount = &v
	return s
}

func (s *ModuleReShopItemListShoppingItemMapValueSearchPrice) SetHasPrice(v bool) *ModuleReShopItemListShoppingItemMapValueSearchPrice {
	s.HasPrice = &v
	return s
}

func (s *ModuleReShopItemListShoppingItemMapValueSearchPrice) SetNonPriceText(v string) *ModuleReShopItemListShoppingItemMapValueSearchPrice {
	s.NonPriceText = &v
	return s
}

func (s *ModuleReShopItemListShoppingItemMapValueSearchPrice) Validate() error {
	return dara.Validate(s)
}

type ModuleReShopItemListShoppingItemMapValueSegmentPriceList struct {
	SegmentPosition *ModuleReShopItemListShoppingItemMapValueSegmentPriceListSegmentPosition `json:"segment_position,omitempty" xml:"segment_position,omitempty" type:"Struct"`
	SearchPrice     *ModuleReShopItemListShoppingItemMapValueSegmentPriceListSearchPrice     `json:"search_price,omitempty" xml:"search_price,omitempty" type:"Struct"`
}

func (s ModuleReShopItemListShoppingItemMapValueSegmentPriceList) String() string {
	return dara.Prettify(s)
}

func (s ModuleReShopItemListShoppingItemMapValueSegmentPriceList) GoString() string {
	return s.String()
}

func (s *ModuleReShopItemListShoppingItemMapValueSegmentPriceList) GetSegmentPosition() *ModuleReShopItemListShoppingItemMapValueSegmentPriceListSegmentPosition {
	return s.SegmentPosition
}

func (s *ModuleReShopItemListShoppingItemMapValueSegmentPriceList) GetSearchPrice() *ModuleReShopItemListShoppingItemMapValueSegmentPriceListSearchPrice {
	return s.SearchPrice
}

func (s *ModuleReShopItemListShoppingItemMapValueSegmentPriceList) SetSegmentPosition(v *ModuleReShopItemListShoppingItemMapValueSegmentPriceListSegmentPosition) *ModuleReShopItemListShoppingItemMapValueSegmentPriceList {
	s.SegmentPosition = v
	return s
}

func (s *ModuleReShopItemListShoppingItemMapValueSegmentPriceList) SetSearchPrice(v *ModuleReShopItemListShoppingItemMapValueSegmentPriceListSearchPrice) *ModuleReShopItemListShoppingItemMapValueSegmentPriceList {
	s.SearchPrice = v
	return s
}

func (s *ModuleReShopItemListShoppingItemMapValueSegmentPriceList) Validate() error {
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

type ModuleReShopItemListShoppingItemMapValueSegmentPriceListSegmentPosition struct {
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
}

func (s ModuleReShopItemListShoppingItemMapValueSegmentPriceListSegmentPosition) String() string {
	return dara.Prettify(s)
}

func (s ModuleReShopItemListShoppingItemMapValueSegmentPriceListSegmentPosition) GoString() string {
	return s.String()
}

func (s *ModuleReShopItemListShoppingItemMapValueSegmentPriceListSegmentPosition) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *ModuleReShopItemListShoppingItemMapValueSegmentPriceListSegmentPosition) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *ModuleReShopItemListShoppingItemMapValueSegmentPriceListSegmentPosition) SetJourneyIndex(v int32) *ModuleReShopItemListShoppingItemMapValueSegmentPriceListSegmentPosition {
	s.JourneyIndex = &v
	return s
}

func (s *ModuleReShopItemListShoppingItemMapValueSegmentPriceListSegmentPosition) SetSegmentIndex(v int32) *ModuleReShopItemListShoppingItemMapValueSegmentPriceListSegmentPosition {
	s.SegmentIndex = &v
	return s
}

func (s *ModuleReShopItemListShoppingItemMapValueSegmentPriceListSegmentPosition) Validate() error {
	return dara.Validate(s)
}

type ModuleReShopItemListShoppingItemMapValueSegmentPriceListSearchPrice struct {
	TotalAmount    *int64  `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	HandlingAmount *int64  `json:"handling_amount,omitempty" xml:"handling_amount,omitempty"`
	UpgradeAmount  *int64  `json:"upgrade_amount,omitempty" xml:"upgrade_amount,omitempty"`
	TaxDiffAmount  *int64  `json:"tax_diff_amount,omitempty" xml:"tax_diff_amount,omitempty"`
	HasPrice       *bool   `json:"has_price,omitempty" xml:"has_price,omitempty"`
	NonPriceText   *string `json:"non_price_text,omitempty" xml:"non_price_text,omitempty"`
}

func (s ModuleReShopItemListShoppingItemMapValueSegmentPriceListSearchPrice) String() string {
	return dara.Prettify(s)
}

func (s ModuleReShopItemListShoppingItemMapValueSegmentPriceListSearchPrice) GoString() string {
	return s.String()
}

func (s *ModuleReShopItemListShoppingItemMapValueSegmentPriceListSearchPrice) GetTotalAmount() *int64 {
	return s.TotalAmount
}

func (s *ModuleReShopItemListShoppingItemMapValueSegmentPriceListSearchPrice) GetHandlingAmount() *int64 {
	return s.HandlingAmount
}

func (s *ModuleReShopItemListShoppingItemMapValueSegmentPriceListSearchPrice) GetUpgradeAmount() *int64 {
	return s.UpgradeAmount
}

func (s *ModuleReShopItemListShoppingItemMapValueSegmentPriceListSearchPrice) GetTaxDiffAmount() *int64 {
	return s.TaxDiffAmount
}

func (s *ModuleReShopItemListShoppingItemMapValueSegmentPriceListSearchPrice) GetHasPrice() *bool {
	return s.HasPrice
}

func (s *ModuleReShopItemListShoppingItemMapValueSegmentPriceListSearchPrice) GetNonPriceText() *string {
	return s.NonPriceText
}

func (s *ModuleReShopItemListShoppingItemMapValueSegmentPriceListSearchPrice) SetTotalAmount(v int64) *ModuleReShopItemListShoppingItemMapValueSegmentPriceListSearchPrice {
	s.TotalAmount = &v
	return s
}

func (s *ModuleReShopItemListShoppingItemMapValueSegmentPriceListSearchPrice) SetHandlingAmount(v int64) *ModuleReShopItemListShoppingItemMapValueSegmentPriceListSearchPrice {
	s.HandlingAmount = &v
	return s
}

func (s *ModuleReShopItemListShoppingItemMapValueSegmentPriceListSearchPrice) SetUpgradeAmount(v int64) *ModuleReShopItemListShoppingItemMapValueSegmentPriceListSearchPrice {
	s.UpgradeAmount = &v
	return s
}

func (s *ModuleReShopItemListShoppingItemMapValueSegmentPriceListSearchPrice) SetTaxDiffAmount(v int64) *ModuleReShopItemListShoppingItemMapValueSegmentPriceListSearchPrice {
	s.TaxDiffAmount = &v
	return s
}

func (s *ModuleReShopItemListShoppingItemMapValueSegmentPriceListSearchPrice) SetHasPrice(v bool) *ModuleReShopItemListShoppingItemMapValueSegmentPriceListSearchPrice {
	s.HasPrice = &v
	return s
}

func (s *ModuleReShopItemListShoppingItemMapValueSegmentPriceListSearchPrice) SetNonPriceText(v string) *ModuleReShopItemListShoppingItemMapValueSegmentPriceListSearchPrice {
	s.NonPriceText = &v
	return s
}

func (s *ModuleReShopItemListShoppingItemMapValueSegmentPriceListSearchPrice) Validate() error {
	return dara.Validate(s)
}
