// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleItemListSubItemsShoppingItemMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetCabinQuantityList(v []*ModuleItemListSubItemsShoppingItemMapValueCabinQuantityList) *ModuleItemListSubItemsShoppingItemMapValue
	GetCabinQuantityList() []*ModuleItemListSubItemsShoppingItemMapValueCabinQuantityList
	SetSearchPrice(v *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) *ModuleItemListSubItemsShoppingItemMapValue
	GetSearchPrice() *ModuleItemListSubItemsShoppingItemMapValueSearchPrice
	SetSegmentPriceList(v []*ModuleItemListSubItemsShoppingItemMapValueSegmentPriceList) *ModuleItemListSubItemsShoppingItemMapValue
	GetSegmentPriceList() []*ModuleItemListSubItemsShoppingItemMapValueSegmentPriceList
	SetId(v string) *ModuleItemListSubItemsShoppingItemMapValue
	GetId() *string
	SetCabinQuantity(v map[string]*ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue) *ModuleItemListSubItemsShoppingItemMapValue
	GetCabinQuantity() map[string]*ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue
	SetSegmentPrice(v map[string]*ModuleItemListSubItemsShoppingItemMapValueSegmentPriceValue) *ModuleItemListSubItemsShoppingItemMapValue
	GetSegmentPrice() map[string]*ModuleItemListSubItemsShoppingItemMapValueSegmentPriceValue
}

type ModuleItemListSubItemsShoppingItemMapValue struct {
	CabinQuantityList []*ModuleItemListSubItemsShoppingItemMapValueCabinQuantityList `json:"cabin_quantity_list,omitempty" xml:"cabin_quantity_list,omitempty" type:"Repeated"`
	SearchPrice       *ModuleItemListSubItemsShoppingItemMapValueSearchPrice         `json:"search_price,omitempty" xml:"search_price,omitempty" type:"Struct"`
	SegmentPriceList  []*ModuleItemListSubItemsShoppingItemMapValueSegmentPriceList  `json:"segment_price_list,omitempty" xml:"segment_price_list,omitempty" type:"Repeated"`
	// id
	//
	// example:
	//
	// ADT07df0bd9-f803-4a50-8449-f4bd675d9939
	Id            *string                                                                  `json:"id,omitempty" xml:"id,omitempty"`
	CabinQuantity map[string]*ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue `json:"cabin_quantity,omitempty" xml:"cabin_quantity,omitempty"`
	SegmentPrice  map[string]*ModuleItemListSubItemsShoppingItemMapValueSegmentPriceValue  `json:"segment_price,omitempty" xml:"segment_price,omitempty"`
}

func (s ModuleItemListSubItemsShoppingItemMapValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleItemListSubItemsShoppingItemMapValue) GoString() string {
	return s.String()
}

func (s *ModuleItemListSubItemsShoppingItemMapValue) GetCabinQuantityList() []*ModuleItemListSubItemsShoppingItemMapValueCabinQuantityList {
	return s.CabinQuantityList
}

func (s *ModuleItemListSubItemsShoppingItemMapValue) GetSearchPrice() *ModuleItemListSubItemsShoppingItemMapValueSearchPrice {
	return s.SearchPrice
}

func (s *ModuleItemListSubItemsShoppingItemMapValue) GetSegmentPriceList() []*ModuleItemListSubItemsShoppingItemMapValueSegmentPriceList {
	return s.SegmentPriceList
}

func (s *ModuleItemListSubItemsShoppingItemMapValue) GetId() *string {
	return s.Id
}

func (s *ModuleItemListSubItemsShoppingItemMapValue) GetCabinQuantity() map[string]*ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue {
	return s.CabinQuantity
}

func (s *ModuleItemListSubItemsShoppingItemMapValue) GetSegmentPrice() map[string]*ModuleItemListSubItemsShoppingItemMapValueSegmentPriceValue {
	return s.SegmentPrice
}

func (s *ModuleItemListSubItemsShoppingItemMapValue) SetCabinQuantityList(v []*ModuleItemListSubItemsShoppingItemMapValueCabinQuantityList) *ModuleItemListSubItemsShoppingItemMapValue {
	s.CabinQuantityList = v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValue) SetSearchPrice(v *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) *ModuleItemListSubItemsShoppingItemMapValue {
	s.SearchPrice = v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValue) SetSegmentPriceList(v []*ModuleItemListSubItemsShoppingItemMapValueSegmentPriceList) *ModuleItemListSubItemsShoppingItemMapValue {
	s.SegmentPriceList = v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValue) SetId(v string) *ModuleItemListSubItemsShoppingItemMapValue {
	s.Id = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValue) SetCabinQuantity(v map[string]*ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue) *ModuleItemListSubItemsShoppingItemMapValue {
	s.CabinQuantity = v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValue) SetSegmentPrice(v map[string]*ModuleItemListSubItemsShoppingItemMapValueSegmentPriceValue) *ModuleItemListSubItemsShoppingItemMapValue {
	s.SegmentPrice = v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValue) Validate() error {
	return dara.Validate(s)
}

type ModuleItemListSubItemsShoppingItemMapValueCabinQuantityList struct {
	SegmentPosition *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition `json:"segment_position,omitempty" xml:"segment_position,omitempty" type:"Struct"`
	CabinInfo       *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo       `json:"cabin_info,omitempty" xml:"cabin_info,omitempty" type:"Struct"`
}

func (s ModuleItemListSubItemsShoppingItemMapValueCabinQuantityList) String() string {
	return dara.Prettify(s)
}

func (s ModuleItemListSubItemsShoppingItemMapValueCabinQuantityList) GoString() string {
	return s.String()
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityList) GetSegmentPosition() *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition {
	return s.SegmentPosition
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityList) GetCabinInfo() *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo {
	return s.CabinInfo
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityList) SetSegmentPosition(v *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition) *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityList {
	s.SegmentPosition = v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityList) SetCabinInfo(v *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityList {
	s.CabinInfo = v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityList) Validate() error {
	return dara.Validate(s)
}

type ModuleItemListSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition struct {
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 0
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
}

func (s ModuleItemListSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition) String() string {
	return dara.Prettify(s)
}

func (s ModuleItemListSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition) GoString() string {
	return s.String()
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition) SetJourneyIndex(v int32) *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition {
	s.JourneyIndex = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition) SetSegmentIndex(v int32) *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition {
	s.SegmentIndex = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityListSegmentPosition) Validate() error {
	return dara.Validate(s)
}

type ModuleItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo struct {
	Cabin          *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	CabinClass     *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	CabinClassName *string `json:"cabin_class_name,omitempty" xml:"cabin_class_name,omitempty"`
	Quantity       *string `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

func (s ModuleItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) String() string {
	return dara.Prettify(s)
}

func (s ModuleItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) GoString() string {
	return s.String()
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) GetCabin() *string {
	return s.Cabin
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) GetCabinClass() *string {
	return s.CabinClass
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) GetCabinClassName() *string {
	return s.CabinClassName
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) GetQuantity() *string {
	return s.Quantity
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) SetCabin(v string) *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo {
	s.Cabin = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) SetCabinClass(v string) *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo {
	s.CabinClass = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) SetCabinClassName(v string) *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo {
	s.CabinClassName = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) SetQuantity(v string) *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo {
	s.Quantity = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityListCabinInfo) Validate() error {
	return dara.Validate(s)
}

type ModuleItemListSubItemsShoppingItemMapValueSearchPrice struct {
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
	MinBeforeControlPriceOfNormal *int32                                                              `json:"min_before_control_price_of_normal,omitempty" xml:"min_before_control_price_of_normal,omitempty"`
	PriceShowInfo                 *ModuleItemListSubItemsShoppingItemMapValueSearchPricePriceShowInfo `json:"price_show_info,omitempty" xml:"price_show_info,omitempty" type:"Struct"`
}

func (s ModuleItemListSubItemsShoppingItemMapValueSearchPrice) String() string {
	return dara.Prettify(s)
}

func (s ModuleItemListSubItemsShoppingItemMapValueSearchPrice) GoString() string {
	return s.String()
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) GetFloorPrice() *int32 {
	return s.FloorPrice
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) GetTicketPrice() *int32 {
	return s.TicketPrice
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) GetSellPrice() *int32 {
	return s.SellPrice
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) GetOriginalSellPrice() *int32 {
	return s.OriginalSellPrice
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) GetBaseTotalPrice() *int32 {
	return s.BaseTotalPrice
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) GetBeforeControlPrice() *int32 {
	return s.BeforeControlPrice
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) GetTax() *int32 {
	return s.Tax
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) GetSupplyPrice() *int32 {
	return s.SupplyPrice
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) GetBasicCabinPrice() *int32 {
	return s.BasicCabinPrice
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) GetBuildPrice() *int32 {
	return s.BuildPrice
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) GetOilPrice() *int32 {
	return s.OilPrice
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) GetFirstStandardPrice() *int32 {
	return s.FirstStandardPrice
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) GetBusinessStandardPrice() *int32 {
	return s.BusinessStandardPrice
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) GetCommonStandardPrice() *int32 {
	return s.CommonStandardPrice
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) GetInterTicketPrice() *int32 {
	return s.InterTicketPrice
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) GetSubtractedPrice() *int32 {
	return s.SubtractedPrice
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) GetOriginCommonPrice() *int32 {
	return s.OriginCommonPrice
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) GetDynamicPromotionPrice() *int32 {
	return s.DynamicPromotionPrice
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) GetInstallmentNum() *int32 {
	return s.InstallmentNum
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) GetInstallmentPrice() *float64 {
	return s.InstallmentPrice
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) GetCompetitionDynamicPrice() *int32 {
	return s.CompetitionDynamicPrice
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) GetCompetitionPromotionPrice() *int32 {
	return s.CompetitionPromotionPrice
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) GetMinBeforeControlPriceOfNormal() *int32 {
	return s.MinBeforeControlPriceOfNormal
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) GetPriceShowInfo() *ModuleItemListSubItemsShoppingItemMapValueSearchPricePriceShowInfo {
	return s.PriceShowInfo
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) SetFloorPrice(v int32) *ModuleItemListSubItemsShoppingItemMapValueSearchPrice {
	s.FloorPrice = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) SetTicketPrice(v int32) *ModuleItemListSubItemsShoppingItemMapValueSearchPrice {
	s.TicketPrice = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) SetSellPrice(v int32) *ModuleItemListSubItemsShoppingItemMapValueSearchPrice {
	s.SellPrice = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) SetOriginalSellPrice(v int32) *ModuleItemListSubItemsShoppingItemMapValueSearchPrice {
	s.OriginalSellPrice = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) SetBaseTotalPrice(v int32) *ModuleItemListSubItemsShoppingItemMapValueSearchPrice {
	s.BaseTotalPrice = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) SetBeforeControlPrice(v int32) *ModuleItemListSubItemsShoppingItemMapValueSearchPrice {
	s.BeforeControlPrice = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) SetTax(v int32) *ModuleItemListSubItemsShoppingItemMapValueSearchPrice {
	s.Tax = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) SetSupplyPrice(v int32) *ModuleItemListSubItemsShoppingItemMapValueSearchPrice {
	s.SupplyPrice = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) SetBasicCabinPrice(v int32) *ModuleItemListSubItemsShoppingItemMapValueSearchPrice {
	s.BasicCabinPrice = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) SetBuildPrice(v int32) *ModuleItemListSubItemsShoppingItemMapValueSearchPrice {
	s.BuildPrice = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) SetOilPrice(v int32) *ModuleItemListSubItemsShoppingItemMapValueSearchPrice {
	s.OilPrice = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) SetFirstStandardPrice(v int32) *ModuleItemListSubItemsShoppingItemMapValueSearchPrice {
	s.FirstStandardPrice = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) SetBusinessStandardPrice(v int32) *ModuleItemListSubItemsShoppingItemMapValueSearchPrice {
	s.BusinessStandardPrice = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) SetCommonStandardPrice(v int32) *ModuleItemListSubItemsShoppingItemMapValueSearchPrice {
	s.CommonStandardPrice = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) SetInterTicketPrice(v int32) *ModuleItemListSubItemsShoppingItemMapValueSearchPrice {
	s.InterTicketPrice = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) SetSubtractedPrice(v int32) *ModuleItemListSubItemsShoppingItemMapValueSearchPrice {
	s.SubtractedPrice = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) SetOriginCommonPrice(v int32) *ModuleItemListSubItemsShoppingItemMapValueSearchPrice {
	s.OriginCommonPrice = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) SetDynamicPromotionPrice(v int32) *ModuleItemListSubItemsShoppingItemMapValueSearchPrice {
	s.DynamicPromotionPrice = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) SetInstallmentNum(v int32) *ModuleItemListSubItemsShoppingItemMapValueSearchPrice {
	s.InstallmentNum = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) SetInstallmentPrice(v float64) *ModuleItemListSubItemsShoppingItemMapValueSearchPrice {
	s.InstallmentPrice = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) SetCompetitionDynamicPrice(v int32) *ModuleItemListSubItemsShoppingItemMapValueSearchPrice {
	s.CompetitionDynamicPrice = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) SetCompetitionPromotionPrice(v int32) *ModuleItemListSubItemsShoppingItemMapValueSearchPrice {
	s.CompetitionPromotionPrice = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) SetMinBeforeControlPriceOfNormal(v int32) *ModuleItemListSubItemsShoppingItemMapValueSearchPrice {
	s.MinBeforeControlPriceOfNormal = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) SetPriceShowInfo(v *ModuleItemListSubItemsShoppingItemMapValueSearchPricePriceShowInfo) *ModuleItemListSubItemsShoppingItemMapValueSearchPrice {
	s.PriceShowInfo = v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPrice) Validate() error {
	return dara.Validate(s)
}

type ModuleItemListSubItemsShoppingItemMapValueSearchPricePriceShowInfo struct {
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

func (s ModuleItemListSubItemsShoppingItemMapValueSearchPricePriceShowInfo) String() string {
	return dara.Prettify(s)
}

func (s ModuleItemListSubItemsShoppingItemMapValueSearchPricePriceShowInfo) GoString() string {
	return s.String()
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPricePriceShowInfo) GetDiscountInfo() *string {
	return s.DiscountInfo
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPricePriceShowInfo) GetDiscountNum() *float64 {
	return s.DiscountNum
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPricePriceShowInfo) GetShowTicketPrice() *bool {
	return s.ShowTicketPrice
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPricePriceShowInfo) SetDiscountInfo(v string) *ModuleItemListSubItemsShoppingItemMapValueSearchPricePriceShowInfo {
	s.DiscountInfo = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPricePriceShowInfo) SetDiscountNum(v float64) *ModuleItemListSubItemsShoppingItemMapValueSearchPricePriceShowInfo {
	s.DiscountNum = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPricePriceShowInfo) SetShowTicketPrice(v bool) *ModuleItemListSubItemsShoppingItemMapValueSearchPricePriceShowInfo {
	s.ShowTicketPrice = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSearchPricePriceShowInfo) Validate() error {
	return dara.Validate(s)
}

type ModuleItemListSubItemsShoppingItemMapValueSegmentPriceList struct {
	SegmentPosition *ModuleItemListSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition `json:"segment_position,omitempty" xml:"segment_position,omitempty" type:"Struct"`
	SearchPrice     *ModuleItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice     `json:"search_price,omitempty" xml:"search_price,omitempty" type:"Struct"`
}

func (s ModuleItemListSubItemsShoppingItemMapValueSegmentPriceList) String() string {
	return dara.Prettify(s)
}

func (s ModuleItemListSubItemsShoppingItemMapValueSegmentPriceList) GoString() string {
	return s.String()
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSegmentPriceList) GetSegmentPosition() *ModuleItemListSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition {
	return s.SegmentPosition
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSegmentPriceList) GetSearchPrice() *ModuleItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice {
	return s.SearchPrice
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSegmentPriceList) SetSegmentPosition(v *ModuleItemListSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition) *ModuleItemListSubItemsShoppingItemMapValueSegmentPriceList {
	s.SegmentPosition = v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSegmentPriceList) SetSearchPrice(v *ModuleItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) *ModuleItemListSubItemsShoppingItemMapValueSegmentPriceList {
	s.SearchPrice = v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSegmentPriceList) Validate() error {
	return dara.Validate(s)
}

type ModuleItemListSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition struct {
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 0
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
}

func (s ModuleItemListSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition) String() string {
	return dara.Prettify(s)
}

func (s ModuleItemListSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition) GoString() string {
	return s.String()
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition) SetJourneyIndex(v int32) *ModuleItemListSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition {
	s.JourneyIndex = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition) SetSegmentIndex(v int32) *ModuleItemListSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition {
	s.SegmentIndex = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSegmentPriceListSegmentPosition) Validate() error {
	return dara.Validate(s)
}

type ModuleItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice struct {
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

func (s ModuleItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) String() string {
	return dara.Prettify(s)
}

func (s ModuleItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) GoString() string {
	return s.String()
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) GetSellPrice() *int32 {
	return s.SellPrice
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) GetTicketPrice() *int32 {
	return s.TicketPrice
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) GetTax() *int32 {
	return s.Tax
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) SetSellPrice(v int32) *ModuleItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice {
	s.SellPrice = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) SetTicketPrice(v int32) *ModuleItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice {
	s.TicketPrice = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) SetTax(v int32) *ModuleItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice {
	s.Tax = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueSegmentPriceListSearchPrice) Validate() error {
	return dara.Validate(s)
}
