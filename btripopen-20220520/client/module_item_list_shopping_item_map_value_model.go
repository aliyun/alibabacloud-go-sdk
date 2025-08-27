// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleItemListShoppingItemMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetCabinQuantityList(v []*ModuleItemListShoppingItemMapValueCabinQuantityList) *ModuleItemListShoppingItemMapValue
	GetCabinQuantityList() []*ModuleItemListShoppingItemMapValueCabinQuantityList
	SetSearchPrice(v *ModuleItemListShoppingItemMapValueSearchPrice) *ModuleItemListShoppingItemMapValue
	GetSearchPrice() *ModuleItemListShoppingItemMapValueSearchPrice
	SetSegmentPriceList(v []*ModuleItemListShoppingItemMapValueSegmentPriceList) *ModuleItemListShoppingItemMapValue
	GetSegmentPriceList() []*ModuleItemListShoppingItemMapValueSegmentPriceList
	SetId(v string) *ModuleItemListShoppingItemMapValue
	GetId() *string
	SetCabinQuantity(v map[string]*ModuleItemListShoppingItemMapValueCabinQuantityValue) *ModuleItemListShoppingItemMapValue
	GetCabinQuantity() map[string]*ModuleItemListShoppingItemMapValueCabinQuantityValue
	SetSegmentPrice(v map[string]*ModuleItemListShoppingItemMapValueSegmentPriceValue) *ModuleItemListShoppingItemMapValue
	GetSegmentPrice() map[string]*ModuleItemListShoppingItemMapValueSegmentPriceValue
}

type ModuleItemListShoppingItemMapValue struct {
	CabinQuantityList []*ModuleItemListShoppingItemMapValueCabinQuantityList `json:"cabin_quantity_list,omitempty" xml:"cabin_quantity_list,omitempty" type:"Repeated"`
	SearchPrice       *ModuleItemListShoppingItemMapValueSearchPrice         `json:"search_price,omitempty" xml:"search_price,omitempty" type:"Struct"`
	SegmentPriceList  []*ModuleItemListShoppingItemMapValueSegmentPriceList  `json:"segment_price_list,omitempty" xml:"segment_price_list,omitempty" type:"Repeated"`
	// id
	//
	// example:
	//
	// ADT07df0bd9-f803-4a50-8449-f4bd675d9939
	Id            *string                                                          `json:"id,omitempty" xml:"id,omitempty"`
	CabinQuantity map[string]*ModuleItemListShoppingItemMapValueCabinQuantityValue `json:"cabin_quantity,omitempty" xml:"cabin_quantity,omitempty"`
	SegmentPrice  map[string]*ModuleItemListShoppingItemMapValueSegmentPriceValue  `json:"segment_price,omitempty" xml:"segment_price,omitempty"`
}

func (s ModuleItemListShoppingItemMapValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleItemListShoppingItemMapValue) GoString() string {
	return s.String()
}

func (s *ModuleItemListShoppingItemMapValue) GetCabinQuantityList() []*ModuleItemListShoppingItemMapValueCabinQuantityList {
	return s.CabinQuantityList
}

func (s *ModuleItemListShoppingItemMapValue) GetSearchPrice() *ModuleItemListShoppingItemMapValueSearchPrice {
	return s.SearchPrice
}

func (s *ModuleItemListShoppingItemMapValue) GetSegmentPriceList() []*ModuleItemListShoppingItemMapValueSegmentPriceList {
	return s.SegmentPriceList
}

func (s *ModuleItemListShoppingItemMapValue) GetId() *string {
	return s.Id
}

func (s *ModuleItemListShoppingItemMapValue) GetCabinQuantity() map[string]*ModuleItemListShoppingItemMapValueCabinQuantityValue {
	return s.CabinQuantity
}

func (s *ModuleItemListShoppingItemMapValue) GetSegmentPrice() map[string]*ModuleItemListShoppingItemMapValueSegmentPriceValue {
	return s.SegmentPrice
}

func (s *ModuleItemListShoppingItemMapValue) SetCabinQuantityList(v []*ModuleItemListShoppingItemMapValueCabinQuantityList) *ModuleItemListShoppingItemMapValue {
	s.CabinQuantityList = v
	return s
}

func (s *ModuleItemListShoppingItemMapValue) SetSearchPrice(v *ModuleItemListShoppingItemMapValueSearchPrice) *ModuleItemListShoppingItemMapValue {
	s.SearchPrice = v
	return s
}

func (s *ModuleItemListShoppingItemMapValue) SetSegmentPriceList(v []*ModuleItemListShoppingItemMapValueSegmentPriceList) *ModuleItemListShoppingItemMapValue {
	s.SegmentPriceList = v
	return s
}

func (s *ModuleItemListShoppingItemMapValue) SetId(v string) *ModuleItemListShoppingItemMapValue {
	s.Id = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValue) SetCabinQuantity(v map[string]*ModuleItemListShoppingItemMapValueCabinQuantityValue) *ModuleItemListShoppingItemMapValue {
	s.CabinQuantity = v
	return s
}

func (s *ModuleItemListShoppingItemMapValue) SetSegmentPrice(v map[string]*ModuleItemListShoppingItemMapValueSegmentPriceValue) *ModuleItemListShoppingItemMapValue {
	s.SegmentPrice = v
	return s
}

func (s *ModuleItemListShoppingItemMapValue) Validate() error {
	return dara.Validate(s)
}

type ModuleItemListShoppingItemMapValueCabinQuantityList struct {
	SegmentPosition *ModuleItemListShoppingItemMapValueCabinQuantityListSegmentPosition `json:"segment_position,omitempty" xml:"segment_position,omitempty" type:"Struct"`
	CabinInfo       *ModuleItemListShoppingItemMapValueCabinQuantityListCabinInfo       `json:"cabin_info,omitempty" xml:"cabin_info,omitempty" type:"Struct"`
}

func (s ModuleItemListShoppingItemMapValueCabinQuantityList) String() string {
	return dara.Prettify(s)
}

func (s ModuleItemListShoppingItemMapValueCabinQuantityList) GoString() string {
	return s.String()
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityList) GetSegmentPosition() *ModuleItemListShoppingItemMapValueCabinQuantityListSegmentPosition {
	return s.SegmentPosition
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityList) GetCabinInfo() *ModuleItemListShoppingItemMapValueCabinQuantityListCabinInfo {
	return s.CabinInfo
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityList) SetSegmentPosition(v *ModuleItemListShoppingItemMapValueCabinQuantityListSegmentPosition) *ModuleItemListShoppingItemMapValueCabinQuantityList {
	s.SegmentPosition = v
	return s
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityList) SetCabinInfo(v *ModuleItemListShoppingItemMapValueCabinQuantityListCabinInfo) *ModuleItemListShoppingItemMapValueCabinQuantityList {
	s.CabinInfo = v
	return s
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityList) Validate() error {
	return dara.Validate(s)
}

type ModuleItemListShoppingItemMapValueCabinQuantityListSegmentPosition struct {
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 0
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
}

func (s ModuleItemListShoppingItemMapValueCabinQuantityListSegmentPosition) String() string {
	return dara.Prettify(s)
}

func (s ModuleItemListShoppingItemMapValueCabinQuantityListSegmentPosition) GoString() string {
	return s.String()
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityListSegmentPosition) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityListSegmentPosition) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityListSegmentPosition) SetJourneyIndex(v int32) *ModuleItemListShoppingItemMapValueCabinQuantityListSegmentPosition {
	s.JourneyIndex = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityListSegmentPosition) SetSegmentIndex(v int32) *ModuleItemListShoppingItemMapValueCabinQuantityListSegmentPosition {
	s.SegmentIndex = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityListSegmentPosition) Validate() error {
	return dara.Validate(s)
}

type ModuleItemListShoppingItemMapValueCabinQuantityListCabinInfo struct {
	Cabin          *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	CabinClass     *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	CabinClassName *string `json:"cabin_class_name,omitempty" xml:"cabin_class_name,omitempty"`
	Quantity       *string `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

func (s ModuleItemListShoppingItemMapValueCabinQuantityListCabinInfo) String() string {
	return dara.Prettify(s)
}

func (s ModuleItemListShoppingItemMapValueCabinQuantityListCabinInfo) GoString() string {
	return s.String()
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityListCabinInfo) GetCabin() *string {
	return s.Cabin
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityListCabinInfo) GetCabinClass() *string {
	return s.CabinClass
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityListCabinInfo) GetCabinClassName() *string {
	return s.CabinClassName
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityListCabinInfo) GetQuantity() *string {
	return s.Quantity
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityListCabinInfo) SetCabin(v string) *ModuleItemListShoppingItemMapValueCabinQuantityListCabinInfo {
	s.Cabin = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityListCabinInfo) SetCabinClass(v string) *ModuleItemListShoppingItemMapValueCabinQuantityListCabinInfo {
	s.CabinClass = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityListCabinInfo) SetCabinClassName(v string) *ModuleItemListShoppingItemMapValueCabinQuantityListCabinInfo {
	s.CabinClassName = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityListCabinInfo) SetQuantity(v string) *ModuleItemListShoppingItemMapValueCabinQuantityListCabinInfo {
	s.Quantity = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityListCabinInfo) Validate() error {
	return dara.Validate(s)
}

type ModuleItemListShoppingItemMapValueSearchPrice struct {
	// example:
	//
	// 17400
	FloorPrice *int32 `json:"floor_price,omitempty" xml:"floor_price,omitempty"`
	// example:
	//
	// 121000
	TicketPrice *int32 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// example:
	//
	// 17400
	SellPrice *int32 `json:"sell_price,omitempty" xml:"sell_price,omitempty"`
	// example:
	//
	// 17400
	OriginalSellPrice *int32 `json:"original_sell_price,omitempty" xml:"original_sell_price,omitempty"`
	BaseTotalPrice    *int32 `json:"base_total_price,omitempty" xml:"base_total_price,omitempty"`
	// example:
	//
	// 17400
	BeforeControlPrice *int32 `json:"before_control_price,omitempty" xml:"before_control_price,omitempty"`
	// example:
	//
	// 11000
	Tax         *int32 `json:"tax,omitempty" xml:"tax,omitempty"`
	SupplyPrice *int32 `json:"supply_price,omitempty" xml:"supply_price,omitempty"`
	// example:
	//
	// 242000
	BasicCabinPrice *int32 `json:"basic_cabin_price,omitempty" xml:"basic_cabin_price,omitempty"`
	// example:
	//
	// 5000
	BuildPrice *int32 `json:"build_price,omitempty" xml:"build_price,omitempty"`
	// example:
	//
	// 6000
	OilPrice              *int32 `json:"oil_price,omitempty" xml:"oil_price,omitempty"`
	FirstStandardPrice    *int32 `json:"first_standard_price,omitempty" xml:"first_standard_price,omitempty"`
	BusinessStandardPrice *int32 `json:"business_standard_price,omitempty" xml:"business_standard_price,omitempty"`
	// example:
	//
	// 242000
	CommonStandardPrice *int32 `json:"common_standard_price,omitempty" xml:"common_standard_price,omitempty"`
	// fdPrice
	//
	// example:
	//
	// fdPrice
	InterTicketPrice          *int32   `json:"inter_ticket_price,omitempty" xml:"inter_ticket_price,omitempty"`
	SubtractedPrice           *int32   `json:"subtracted_price,omitempty" xml:"subtracted_price,omitempty"`
	OriginCommonPrice         *int32   `json:"origin_common_price,omitempty" xml:"origin_common_price,omitempty"`
	DynamicPromotionPrice     *int32   `json:"dynamic_promotion_price,omitempty" xml:"dynamic_promotion_price,omitempty"`
	InstallmentNum            *int32   `json:"installment_num,omitempty" xml:"installment_num,omitempty"`
	InstallmentPrice          *float64 `json:"installment_price,omitempty" xml:"installment_price,omitempty"`
	CompetitionDynamicPrice   *int32   `json:"competition_dynamic_price,omitempty" xml:"competition_dynamic_price,omitempty"`
	CompetitionPromotionPrice *int32   `json:"competition_promotion_price,omitempty" xml:"competition_promotion_price,omitempty"`
	// example:
	//
	// 17400
	MinBeforeControlPriceOfNormal *int32                                                      `json:"min_before_control_price_of_normal,omitempty" xml:"min_before_control_price_of_normal,omitempty"`
	PriceShowInfo                 *ModuleItemListShoppingItemMapValueSearchPricePriceShowInfo `json:"price_show_info,omitempty" xml:"price_show_info,omitempty" type:"Struct"`
}

func (s ModuleItemListShoppingItemMapValueSearchPrice) String() string {
	return dara.Prettify(s)
}

func (s ModuleItemListShoppingItemMapValueSearchPrice) GoString() string {
	return s.String()
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) GetFloorPrice() *int32 {
	return s.FloorPrice
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) GetTicketPrice() *int32 {
	return s.TicketPrice
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) GetSellPrice() *int32 {
	return s.SellPrice
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) GetOriginalSellPrice() *int32 {
	return s.OriginalSellPrice
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) GetBaseTotalPrice() *int32 {
	return s.BaseTotalPrice
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) GetBeforeControlPrice() *int32 {
	return s.BeforeControlPrice
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) GetTax() *int32 {
	return s.Tax
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) GetSupplyPrice() *int32 {
	return s.SupplyPrice
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) GetBasicCabinPrice() *int32 {
	return s.BasicCabinPrice
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) GetBuildPrice() *int32 {
	return s.BuildPrice
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) GetOilPrice() *int32 {
	return s.OilPrice
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) GetFirstStandardPrice() *int32 {
	return s.FirstStandardPrice
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) GetBusinessStandardPrice() *int32 {
	return s.BusinessStandardPrice
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) GetCommonStandardPrice() *int32 {
	return s.CommonStandardPrice
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) GetInterTicketPrice() *int32 {
	return s.InterTicketPrice
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) GetSubtractedPrice() *int32 {
	return s.SubtractedPrice
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) GetOriginCommonPrice() *int32 {
	return s.OriginCommonPrice
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) GetDynamicPromotionPrice() *int32 {
	return s.DynamicPromotionPrice
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) GetInstallmentNum() *int32 {
	return s.InstallmentNum
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) GetInstallmentPrice() *float64 {
	return s.InstallmentPrice
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) GetCompetitionDynamicPrice() *int32 {
	return s.CompetitionDynamicPrice
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) GetCompetitionPromotionPrice() *int32 {
	return s.CompetitionPromotionPrice
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) GetMinBeforeControlPriceOfNormal() *int32 {
	return s.MinBeforeControlPriceOfNormal
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) GetPriceShowInfo() *ModuleItemListShoppingItemMapValueSearchPricePriceShowInfo {
	return s.PriceShowInfo
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) SetFloorPrice(v int32) *ModuleItemListShoppingItemMapValueSearchPrice {
	s.FloorPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) SetTicketPrice(v int32) *ModuleItemListShoppingItemMapValueSearchPrice {
	s.TicketPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) SetSellPrice(v int32) *ModuleItemListShoppingItemMapValueSearchPrice {
	s.SellPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) SetOriginalSellPrice(v int32) *ModuleItemListShoppingItemMapValueSearchPrice {
	s.OriginalSellPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) SetBaseTotalPrice(v int32) *ModuleItemListShoppingItemMapValueSearchPrice {
	s.BaseTotalPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) SetBeforeControlPrice(v int32) *ModuleItemListShoppingItemMapValueSearchPrice {
	s.BeforeControlPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) SetTax(v int32) *ModuleItemListShoppingItemMapValueSearchPrice {
	s.Tax = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) SetSupplyPrice(v int32) *ModuleItemListShoppingItemMapValueSearchPrice {
	s.SupplyPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) SetBasicCabinPrice(v int32) *ModuleItemListShoppingItemMapValueSearchPrice {
	s.BasicCabinPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) SetBuildPrice(v int32) *ModuleItemListShoppingItemMapValueSearchPrice {
	s.BuildPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) SetOilPrice(v int32) *ModuleItemListShoppingItemMapValueSearchPrice {
	s.OilPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) SetFirstStandardPrice(v int32) *ModuleItemListShoppingItemMapValueSearchPrice {
	s.FirstStandardPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) SetBusinessStandardPrice(v int32) *ModuleItemListShoppingItemMapValueSearchPrice {
	s.BusinessStandardPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) SetCommonStandardPrice(v int32) *ModuleItemListShoppingItemMapValueSearchPrice {
	s.CommonStandardPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) SetInterTicketPrice(v int32) *ModuleItemListShoppingItemMapValueSearchPrice {
	s.InterTicketPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) SetSubtractedPrice(v int32) *ModuleItemListShoppingItemMapValueSearchPrice {
	s.SubtractedPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) SetOriginCommonPrice(v int32) *ModuleItemListShoppingItemMapValueSearchPrice {
	s.OriginCommonPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) SetDynamicPromotionPrice(v int32) *ModuleItemListShoppingItemMapValueSearchPrice {
	s.DynamicPromotionPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) SetInstallmentNum(v int32) *ModuleItemListShoppingItemMapValueSearchPrice {
	s.InstallmentNum = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) SetInstallmentPrice(v float64) *ModuleItemListShoppingItemMapValueSearchPrice {
	s.InstallmentPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) SetCompetitionDynamicPrice(v int32) *ModuleItemListShoppingItemMapValueSearchPrice {
	s.CompetitionDynamicPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) SetCompetitionPromotionPrice(v int32) *ModuleItemListShoppingItemMapValueSearchPrice {
	s.CompetitionPromotionPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) SetMinBeforeControlPriceOfNormal(v int32) *ModuleItemListShoppingItemMapValueSearchPrice {
	s.MinBeforeControlPriceOfNormal = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) SetPriceShowInfo(v *ModuleItemListShoppingItemMapValueSearchPricePriceShowInfo) *ModuleItemListShoppingItemMapValueSearchPrice {
	s.PriceShowInfo = v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSearchPrice) Validate() error {
	return dara.Validate(s)
}

type ModuleItemListShoppingItemMapValueSearchPricePriceShowInfo struct {
	DiscountInfo *string `json:"discount_info,omitempty" xml:"discount_info,omitempty"`
	// example:
	//
	// 0.8
	DiscountNum *float64 `json:"discount_num,omitempty" xml:"discount_num,omitempty"`
	// example:
	//
	// false
	ShowTicketPrice *bool `json:"show_ticket_price,omitempty" xml:"show_ticket_price,omitempty"`
}

func (s ModuleItemListShoppingItemMapValueSearchPricePriceShowInfo) String() string {
	return dara.Prettify(s)
}

func (s ModuleItemListShoppingItemMapValueSearchPricePriceShowInfo) GoString() string {
	return s.String()
}

func (s *ModuleItemListShoppingItemMapValueSearchPricePriceShowInfo) GetDiscountInfo() *string {
	return s.DiscountInfo
}

func (s *ModuleItemListShoppingItemMapValueSearchPricePriceShowInfo) GetDiscountNum() *float64 {
	return s.DiscountNum
}

func (s *ModuleItemListShoppingItemMapValueSearchPricePriceShowInfo) GetShowTicketPrice() *bool {
	return s.ShowTicketPrice
}

func (s *ModuleItemListShoppingItemMapValueSearchPricePriceShowInfo) SetDiscountInfo(v string) *ModuleItemListShoppingItemMapValueSearchPricePriceShowInfo {
	s.DiscountInfo = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSearchPricePriceShowInfo) SetDiscountNum(v float64) *ModuleItemListShoppingItemMapValueSearchPricePriceShowInfo {
	s.DiscountNum = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSearchPricePriceShowInfo) SetShowTicketPrice(v bool) *ModuleItemListShoppingItemMapValueSearchPricePriceShowInfo {
	s.ShowTicketPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSearchPricePriceShowInfo) Validate() error {
	return dara.Validate(s)
}

type ModuleItemListShoppingItemMapValueSegmentPriceList struct {
	SegmentPosition *ModuleItemListShoppingItemMapValueSegmentPriceListSegmentPosition `json:"segment_position,omitempty" xml:"segment_position,omitempty" type:"Struct"`
	SearchPrice     *ModuleItemListShoppingItemMapValueSegmentPriceListSearchPrice     `json:"search_price,omitempty" xml:"search_price,omitempty" type:"Struct"`
}

func (s ModuleItemListShoppingItemMapValueSegmentPriceList) String() string {
	return dara.Prettify(s)
}

func (s ModuleItemListShoppingItemMapValueSegmentPriceList) GoString() string {
	return s.String()
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceList) GetSegmentPosition() *ModuleItemListShoppingItemMapValueSegmentPriceListSegmentPosition {
	return s.SegmentPosition
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceList) GetSearchPrice() *ModuleItemListShoppingItemMapValueSegmentPriceListSearchPrice {
	return s.SearchPrice
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceList) SetSegmentPosition(v *ModuleItemListShoppingItemMapValueSegmentPriceListSegmentPosition) *ModuleItemListShoppingItemMapValueSegmentPriceList {
	s.SegmentPosition = v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceList) SetSearchPrice(v *ModuleItemListShoppingItemMapValueSegmentPriceListSearchPrice) *ModuleItemListShoppingItemMapValueSegmentPriceList {
	s.SearchPrice = v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceList) Validate() error {
	return dara.Validate(s)
}

type ModuleItemListShoppingItemMapValueSegmentPriceListSegmentPosition struct {
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 0
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
}

func (s ModuleItemListShoppingItemMapValueSegmentPriceListSegmentPosition) String() string {
	return dara.Prettify(s)
}

func (s ModuleItemListShoppingItemMapValueSegmentPriceListSegmentPosition) GoString() string {
	return s.String()
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceListSegmentPosition) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceListSegmentPosition) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceListSegmentPosition) SetJourneyIndex(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceListSegmentPosition {
	s.JourneyIndex = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceListSegmentPosition) SetSegmentIndex(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceListSegmentPosition {
	s.SegmentIndex = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceListSegmentPosition) Validate() error {
	return dara.Validate(s)
}

type ModuleItemListShoppingItemMapValueSegmentPriceListSearchPrice struct {
	// example:
	//
	// 120000
	SellPrice *int32 `json:"sell_price,omitempty" xml:"sell_price,omitempty"`
	// example:
	//
	// 120000
	TicketPrice *int32 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// example:
	//
	// 6000
	Tax *int32 `json:"tax,omitempty" xml:"tax,omitempty"`
}

func (s ModuleItemListShoppingItemMapValueSegmentPriceListSearchPrice) String() string {
	return dara.Prettify(s)
}

func (s ModuleItemListShoppingItemMapValueSegmentPriceListSearchPrice) GoString() string {
	return s.String()
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceListSearchPrice) GetSellPrice() *int32 {
	return s.SellPrice
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceListSearchPrice) GetTicketPrice() *int32 {
	return s.TicketPrice
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceListSearchPrice) GetTax() *int32 {
	return s.Tax
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceListSearchPrice) SetSellPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceListSearchPrice {
	s.SellPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceListSearchPrice) SetTicketPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceListSearchPrice {
	s.TicketPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceListSearchPrice) SetTax(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceListSearchPrice {
	s.Tax = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceListSearchPrice) Validate() error {
	return dara.Validate(s)
}
