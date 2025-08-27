// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleGroupItemShoppingItemMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetCabinQuantityList(v []*ModuleGroupItemShoppingItemMapValueCabinQuantityList) *ModuleGroupItemShoppingItemMapValue
	GetCabinQuantityList() []*ModuleGroupItemShoppingItemMapValueCabinQuantityList
	SetSearchPrice(v *ModuleGroupItemShoppingItemMapValueSearchPrice) *ModuleGroupItemShoppingItemMapValue
	GetSearchPrice() *ModuleGroupItemShoppingItemMapValueSearchPrice
	SetSegmentPriceList(v []*ModuleGroupItemShoppingItemMapValueSegmentPriceList) *ModuleGroupItemShoppingItemMapValue
	GetSegmentPriceList() []*ModuleGroupItemShoppingItemMapValueSegmentPriceList
}

type ModuleGroupItemShoppingItemMapValue struct {
	CabinQuantityList []*ModuleGroupItemShoppingItemMapValueCabinQuantityList `json:"cabin_quantity_list,omitempty" xml:"cabin_quantity_list,omitempty" type:"Repeated"`
	SearchPrice       *ModuleGroupItemShoppingItemMapValueSearchPrice         `json:"search_price,omitempty" xml:"search_price,omitempty" type:"Struct"`
	SegmentPriceList  []*ModuleGroupItemShoppingItemMapValueSegmentPriceList  `json:"segment_price_list,omitempty" xml:"segment_price_list,omitempty" type:"Repeated"`
}

func (s ModuleGroupItemShoppingItemMapValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleGroupItemShoppingItemMapValue) GoString() string {
	return s.String()
}

func (s *ModuleGroupItemShoppingItemMapValue) GetCabinQuantityList() []*ModuleGroupItemShoppingItemMapValueCabinQuantityList {
	return s.CabinQuantityList
}

func (s *ModuleGroupItemShoppingItemMapValue) GetSearchPrice() *ModuleGroupItemShoppingItemMapValueSearchPrice {
	return s.SearchPrice
}

func (s *ModuleGroupItemShoppingItemMapValue) GetSegmentPriceList() []*ModuleGroupItemShoppingItemMapValueSegmentPriceList {
	return s.SegmentPriceList
}

func (s *ModuleGroupItemShoppingItemMapValue) SetCabinQuantityList(v []*ModuleGroupItemShoppingItemMapValueCabinQuantityList) *ModuleGroupItemShoppingItemMapValue {
	s.CabinQuantityList = v
	return s
}

func (s *ModuleGroupItemShoppingItemMapValue) SetSearchPrice(v *ModuleGroupItemShoppingItemMapValueSearchPrice) *ModuleGroupItemShoppingItemMapValue {
	s.SearchPrice = v
	return s
}

func (s *ModuleGroupItemShoppingItemMapValue) SetSegmentPriceList(v []*ModuleGroupItemShoppingItemMapValueSegmentPriceList) *ModuleGroupItemShoppingItemMapValue {
	s.SegmentPriceList = v
	return s
}

func (s *ModuleGroupItemShoppingItemMapValue) Validate() error {
	return dara.Validate(s)
}

type ModuleGroupItemShoppingItemMapValueCabinQuantityList struct {
	SegmentPosition *ModuleGroupItemShoppingItemMapValueCabinQuantityListSegmentPosition `json:"segment_position,omitempty" xml:"segment_position,omitempty" type:"Struct"`
	Cabin           *ModuleGroupItemShoppingItemMapValueCabinQuantityListCabin           `json:"cabin,omitempty" xml:"cabin,omitempty" type:"Struct"`
}

func (s ModuleGroupItemShoppingItemMapValueCabinQuantityList) String() string {
	return dara.Prettify(s)
}

func (s ModuleGroupItemShoppingItemMapValueCabinQuantityList) GoString() string {
	return s.String()
}

func (s *ModuleGroupItemShoppingItemMapValueCabinQuantityList) GetSegmentPosition() *ModuleGroupItemShoppingItemMapValueCabinQuantityListSegmentPosition {
	return s.SegmentPosition
}

func (s *ModuleGroupItemShoppingItemMapValueCabinQuantityList) GetCabin() *ModuleGroupItemShoppingItemMapValueCabinQuantityListCabin {
	return s.Cabin
}

func (s *ModuleGroupItemShoppingItemMapValueCabinQuantityList) SetSegmentPosition(v *ModuleGroupItemShoppingItemMapValueCabinQuantityListSegmentPosition) *ModuleGroupItemShoppingItemMapValueCabinQuantityList {
	s.SegmentPosition = v
	return s
}

func (s *ModuleGroupItemShoppingItemMapValueCabinQuantityList) SetCabin(v *ModuleGroupItemShoppingItemMapValueCabinQuantityListCabin) *ModuleGroupItemShoppingItemMapValueCabinQuantityList {
	s.Cabin = v
	return s
}

func (s *ModuleGroupItemShoppingItemMapValueCabinQuantityList) Validate() error {
	return dara.Validate(s)
}

type ModuleGroupItemShoppingItemMapValueCabinQuantityListSegmentPosition struct {
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 0
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
}

func (s ModuleGroupItemShoppingItemMapValueCabinQuantityListSegmentPosition) String() string {
	return dara.Prettify(s)
}

func (s ModuleGroupItemShoppingItemMapValueCabinQuantityListSegmentPosition) GoString() string {
	return s.String()
}

func (s *ModuleGroupItemShoppingItemMapValueCabinQuantityListSegmentPosition) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *ModuleGroupItemShoppingItemMapValueCabinQuantityListSegmentPosition) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *ModuleGroupItemShoppingItemMapValueCabinQuantityListSegmentPosition) SetJourneyIndex(v int32) *ModuleGroupItemShoppingItemMapValueCabinQuantityListSegmentPosition {
	s.JourneyIndex = &v
	return s
}

func (s *ModuleGroupItemShoppingItemMapValueCabinQuantityListSegmentPosition) SetSegmentIndex(v int32) *ModuleGroupItemShoppingItemMapValueCabinQuantityListSegmentPosition {
	s.SegmentIndex = &v
	return s
}

func (s *ModuleGroupItemShoppingItemMapValueCabinQuantityListSegmentPosition) Validate() error {
	return dara.Validate(s)
}

type ModuleGroupItemShoppingItemMapValueCabinQuantityListCabin struct {
	// example:
	//
	// Y
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// example:
	//
	// Y
	CabinClass     *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	CabinClassName *string `json:"cabin_class_name,omitempty" xml:"cabin_class_name,omitempty"`
	// example:
	//
	// A
	Quantity *string `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

func (s ModuleGroupItemShoppingItemMapValueCabinQuantityListCabin) String() string {
	return dara.Prettify(s)
}

func (s ModuleGroupItemShoppingItemMapValueCabinQuantityListCabin) GoString() string {
	return s.String()
}

func (s *ModuleGroupItemShoppingItemMapValueCabinQuantityListCabin) GetCabin() *string {
	return s.Cabin
}

func (s *ModuleGroupItemShoppingItemMapValueCabinQuantityListCabin) GetCabinClass() *string {
	return s.CabinClass
}

func (s *ModuleGroupItemShoppingItemMapValueCabinQuantityListCabin) GetCabinClassName() *string {
	return s.CabinClassName
}

func (s *ModuleGroupItemShoppingItemMapValueCabinQuantityListCabin) GetQuantity() *string {
	return s.Quantity
}

func (s *ModuleGroupItemShoppingItemMapValueCabinQuantityListCabin) SetCabin(v string) *ModuleGroupItemShoppingItemMapValueCabinQuantityListCabin {
	s.Cabin = &v
	return s
}

func (s *ModuleGroupItemShoppingItemMapValueCabinQuantityListCabin) SetCabinClass(v string) *ModuleGroupItemShoppingItemMapValueCabinQuantityListCabin {
	s.CabinClass = &v
	return s
}

func (s *ModuleGroupItemShoppingItemMapValueCabinQuantityListCabin) SetCabinClassName(v string) *ModuleGroupItemShoppingItemMapValueCabinQuantityListCabin {
	s.CabinClassName = &v
	return s
}

func (s *ModuleGroupItemShoppingItemMapValueCabinQuantityListCabin) SetQuantity(v string) *ModuleGroupItemShoppingItemMapValueCabinQuantityListCabin {
	s.Quantity = &v
	return s
}

func (s *ModuleGroupItemShoppingItemMapValueCabinQuantityListCabin) Validate() error {
	return dara.Validate(s)
}

type ModuleGroupItemShoppingItemMapValueSearchPrice struct {
	// example:
	//
	// 120000
	TicketPrice *int32 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// example:
	//
	// 120000
	SellPrice *int32 `json:"sell_price,omitempty" xml:"sell_price,omitempty"`
	// example:
	//
	// 6000
	Tax *int32 `json:"tax,omitempty" xml:"tax,omitempty"`
}

func (s ModuleGroupItemShoppingItemMapValueSearchPrice) String() string {
	return dara.Prettify(s)
}

func (s ModuleGroupItemShoppingItemMapValueSearchPrice) GoString() string {
	return s.String()
}

func (s *ModuleGroupItemShoppingItemMapValueSearchPrice) GetTicketPrice() *int32 {
	return s.TicketPrice
}

func (s *ModuleGroupItemShoppingItemMapValueSearchPrice) GetSellPrice() *int32 {
	return s.SellPrice
}

func (s *ModuleGroupItemShoppingItemMapValueSearchPrice) GetTax() *int32 {
	return s.Tax
}

func (s *ModuleGroupItemShoppingItemMapValueSearchPrice) SetTicketPrice(v int32) *ModuleGroupItemShoppingItemMapValueSearchPrice {
	s.TicketPrice = &v
	return s
}

func (s *ModuleGroupItemShoppingItemMapValueSearchPrice) SetSellPrice(v int32) *ModuleGroupItemShoppingItemMapValueSearchPrice {
	s.SellPrice = &v
	return s
}

func (s *ModuleGroupItemShoppingItemMapValueSearchPrice) SetTax(v int32) *ModuleGroupItemShoppingItemMapValueSearchPrice {
	s.Tax = &v
	return s
}

func (s *ModuleGroupItemShoppingItemMapValueSearchPrice) Validate() error {
	return dara.Validate(s)
}

type ModuleGroupItemShoppingItemMapValueSegmentPriceList struct {
	SegmentPosition *ModuleGroupItemShoppingItemMapValueSegmentPriceListSegmentPosition `json:"segment_position,omitempty" xml:"segment_position,omitempty" type:"Struct"`
	SearchPrice     *ModuleGroupItemShoppingItemMapValueSegmentPriceListSearchPrice     `json:"search_price,omitempty" xml:"search_price,omitempty" type:"Struct"`
}

func (s ModuleGroupItemShoppingItemMapValueSegmentPriceList) String() string {
	return dara.Prettify(s)
}

func (s ModuleGroupItemShoppingItemMapValueSegmentPriceList) GoString() string {
	return s.String()
}

func (s *ModuleGroupItemShoppingItemMapValueSegmentPriceList) GetSegmentPosition() *ModuleGroupItemShoppingItemMapValueSegmentPriceListSegmentPosition {
	return s.SegmentPosition
}

func (s *ModuleGroupItemShoppingItemMapValueSegmentPriceList) GetSearchPrice() *ModuleGroupItemShoppingItemMapValueSegmentPriceListSearchPrice {
	return s.SearchPrice
}

func (s *ModuleGroupItemShoppingItemMapValueSegmentPriceList) SetSegmentPosition(v *ModuleGroupItemShoppingItemMapValueSegmentPriceListSegmentPosition) *ModuleGroupItemShoppingItemMapValueSegmentPriceList {
	s.SegmentPosition = v
	return s
}

func (s *ModuleGroupItemShoppingItemMapValueSegmentPriceList) SetSearchPrice(v *ModuleGroupItemShoppingItemMapValueSegmentPriceListSearchPrice) *ModuleGroupItemShoppingItemMapValueSegmentPriceList {
	s.SearchPrice = v
	return s
}

func (s *ModuleGroupItemShoppingItemMapValueSegmentPriceList) Validate() error {
	return dara.Validate(s)
}

type ModuleGroupItemShoppingItemMapValueSegmentPriceListSegmentPosition struct {
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 0
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
}

func (s ModuleGroupItemShoppingItemMapValueSegmentPriceListSegmentPosition) String() string {
	return dara.Prettify(s)
}

func (s ModuleGroupItemShoppingItemMapValueSegmentPriceListSegmentPosition) GoString() string {
	return s.String()
}

func (s *ModuleGroupItemShoppingItemMapValueSegmentPriceListSegmentPosition) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *ModuleGroupItemShoppingItemMapValueSegmentPriceListSegmentPosition) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *ModuleGroupItemShoppingItemMapValueSegmentPriceListSegmentPosition) SetJourneyIndex(v int32) *ModuleGroupItemShoppingItemMapValueSegmentPriceListSegmentPosition {
	s.JourneyIndex = &v
	return s
}

func (s *ModuleGroupItemShoppingItemMapValueSegmentPriceListSegmentPosition) SetSegmentIndex(v int32) *ModuleGroupItemShoppingItemMapValueSegmentPriceListSegmentPosition {
	s.SegmentIndex = &v
	return s
}

func (s *ModuleGroupItemShoppingItemMapValueSegmentPriceListSegmentPosition) Validate() error {
	return dara.Validate(s)
}

type ModuleGroupItemShoppingItemMapValueSegmentPriceListSearchPrice struct {
	// example:
	//
	// 120000
	TicketPrice *int32 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// example:
	//
	// 120000
	SellPrice *int32 `json:"sell_price,omitempty" xml:"sell_price,omitempty"`
	// example:
	//
	// 6000
	Tax *int32 `json:"tax,omitempty" xml:"tax,omitempty"`
}

func (s ModuleGroupItemShoppingItemMapValueSegmentPriceListSearchPrice) String() string {
	return dara.Prettify(s)
}

func (s ModuleGroupItemShoppingItemMapValueSegmentPriceListSearchPrice) GoString() string {
	return s.String()
}

func (s *ModuleGroupItemShoppingItemMapValueSegmentPriceListSearchPrice) GetTicketPrice() *int32 {
	return s.TicketPrice
}

func (s *ModuleGroupItemShoppingItemMapValueSegmentPriceListSearchPrice) GetSellPrice() *int32 {
	return s.SellPrice
}

func (s *ModuleGroupItemShoppingItemMapValueSegmentPriceListSearchPrice) GetTax() *int32 {
	return s.Tax
}

func (s *ModuleGroupItemShoppingItemMapValueSegmentPriceListSearchPrice) SetTicketPrice(v int32) *ModuleGroupItemShoppingItemMapValueSegmentPriceListSearchPrice {
	s.TicketPrice = &v
	return s
}

func (s *ModuleGroupItemShoppingItemMapValueSegmentPriceListSearchPrice) SetSellPrice(v int32) *ModuleGroupItemShoppingItemMapValueSegmentPriceListSearchPrice {
	s.SellPrice = &v
	return s
}

func (s *ModuleGroupItemShoppingItemMapValueSegmentPriceListSearchPrice) SetTax(v int32) *ModuleGroupItemShoppingItemMapValueSegmentPriceListSearchPrice {
	s.Tax = &v
	return s
}

func (s *ModuleGroupItemShoppingItemMapValueSegmentPriceListSearchPrice) Validate() error {
	return dara.Validate(s)
}
