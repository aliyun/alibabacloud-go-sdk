// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleFlightItemListBestPriceItemShoppingItemMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetCabinQuantityList(v []*ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityList) *ModuleFlightItemListBestPriceItemShoppingItemMapValue
	GetCabinQuantityList() []*ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityList
	SetSearchPrice(v *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) *ModuleFlightItemListBestPriceItemShoppingItemMapValue
	GetSearchPrice() *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice
	SetSegmentPriceList(v []*ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceList) *ModuleFlightItemListBestPriceItemShoppingItemMapValue
	GetSegmentPriceList() []*ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceList
	SetId(v string) *ModuleFlightItemListBestPriceItemShoppingItemMapValue
	GetId() *string
	SetCabinQuantity(v map[string]*ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue) *ModuleFlightItemListBestPriceItemShoppingItemMapValue
	GetCabinQuantity() map[string]*ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue
	SetSegmentPrice(v map[string]*ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) *ModuleFlightItemListBestPriceItemShoppingItemMapValue
	GetSegmentPrice() map[string]*ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue
}

type ModuleFlightItemListBestPriceItemShoppingItemMapValue struct {
	CabinQuantityList []*ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityList `json:"cabin_quantity_list,omitempty" xml:"cabin_quantity_list,omitempty" type:"Repeated"`
	SearchPrice       *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice         `json:"search_price,omitempty" xml:"search_price,omitempty" type:"Struct"`
	SegmentPriceList  []*ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceList  `json:"segment_price_list,omitempty" xml:"segment_price_list,omitempty" type:"Repeated"`
	// id
	//
	// example:
	//
	// ADT07df0bd9-f803-4a50-8449-f4bd675d9939
	Id            *string                                                                             `json:"id,omitempty" xml:"id,omitempty"`
	CabinQuantity map[string]*ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue `json:"cabin_quantity,omitempty" xml:"cabin_quantity,omitempty"`
	SegmentPrice  map[string]*ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue  `json:"segment_price,omitempty" xml:"segment_price,omitempty"`
}

func (s ModuleFlightItemListBestPriceItemShoppingItemMapValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleFlightItemListBestPriceItemShoppingItemMapValue) GoString() string {
	return s.String()
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValue) GetCabinQuantityList() []*ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityList {
	return s.CabinQuantityList
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValue) GetSearchPrice() *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice {
	return s.SearchPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValue) GetSegmentPriceList() []*ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceList {
	return s.SegmentPriceList
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValue) GetId() *string {
	return s.Id
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValue) GetCabinQuantity() map[string]*ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue {
	return s.CabinQuantity
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValue) GetSegmentPrice() map[string]*ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue {
	return s.SegmentPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValue) SetCabinQuantityList(v []*ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityList) *ModuleFlightItemListBestPriceItemShoppingItemMapValue {
	s.CabinQuantityList = v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValue) SetSearchPrice(v *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) *ModuleFlightItemListBestPriceItemShoppingItemMapValue {
	s.SearchPrice = v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValue) SetSegmentPriceList(v []*ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceList) *ModuleFlightItemListBestPriceItemShoppingItemMapValue {
	s.SegmentPriceList = v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValue) SetId(v string) *ModuleFlightItemListBestPriceItemShoppingItemMapValue {
	s.Id = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValue) SetCabinQuantity(v map[string]*ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue) *ModuleFlightItemListBestPriceItemShoppingItemMapValue {
	s.CabinQuantity = v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValue) SetSegmentPrice(v map[string]*ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) *ModuleFlightItemListBestPriceItemShoppingItemMapValue {
	s.SegmentPrice = v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValue) Validate() error {
	return dara.Validate(s)
}

type ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityList struct {
	SegmentPosition *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityListSegmentPosition `json:"segment_position,omitempty" xml:"segment_position,omitempty" type:"Struct"`
	CabinInfo       *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo       `json:"cabin_info,omitempty" xml:"cabin_info,omitempty" type:"Struct"`
}

func (s ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityList) String() string {
	return dara.Prettify(s)
}

func (s ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityList) GoString() string {
	return s.String()
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityList) GetSegmentPosition() *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityListSegmentPosition {
	return s.SegmentPosition
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityList) GetCabinInfo() *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo {
	return s.CabinInfo
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityList) SetSegmentPosition(v *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityListSegmentPosition) *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityList {
	s.SegmentPosition = v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityList) SetCabinInfo(v *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo) *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityList {
	s.CabinInfo = v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityList) Validate() error {
	return dara.Validate(s)
}

type ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityListSegmentPosition struct {
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 0
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
}

func (s ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityListSegmentPosition) String() string {
	return dara.Prettify(s)
}

func (s ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityListSegmentPosition) GoString() string {
	return s.String()
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityListSegmentPosition) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityListSegmentPosition) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityListSegmentPosition) SetJourneyIndex(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityListSegmentPosition {
	s.JourneyIndex = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityListSegmentPosition) SetSegmentIndex(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityListSegmentPosition {
	s.SegmentIndex = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityListSegmentPosition) Validate() error {
	return dara.Validate(s)
}

type ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo struct {
	Cabin          *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	CabinClass     *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	CabinClassName *string `json:"cabin_class_name,omitempty" xml:"cabin_class_name,omitempty"`
	Quantity       *string `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

func (s ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo) String() string {
	return dara.Prettify(s)
}

func (s ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo) GoString() string {
	return s.String()
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo) GetCabin() *string {
	return s.Cabin
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo) GetCabinClass() *string {
	return s.CabinClass
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo) GetCabinClassName() *string {
	return s.CabinClassName
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo) GetQuantity() *string {
	return s.Quantity
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo) SetCabin(v string) *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo {
	s.Cabin = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo) SetCabinClass(v string) *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo {
	s.CabinClass = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo) SetCabinClassName(v string) *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo {
	s.CabinClassName = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo) SetQuantity(v string) *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo {
	s.Quantity = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityListCabinInfo) Validate() error {
	return dara.Validate(s)
}

type ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice struct {
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
	MinBeforeControlPriceOfNormal *int32                                                                         `json:"min_before_control_price_of_normal,omitempty" xml:"min_before_control_price_of_normal,omitempty"`
	PriceShowInfo                 *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPricePriceShowInfo `json:"price_show_info,omitempty" xml:"price_show_info,omitempty" type:"Struct"`
}

func (s ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) String() string {
	return dara.Prettify(s)
}

func (s ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) GoString() string {
	return s.String()
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) GetFloorPrice() *int32 {
	return s.FloorPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) GetTicketPrice() *int32 {
	return s.TicketPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) GetSellPrice() *int32 {
	return s.SellPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) GetOriginalSellPrice() *int32 {
	return s.OriginalSellPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) GetBaseTotalPrice() *int32 {
	return s.BaseTotalPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) GetBeforeControlPrice() *int32 {
	return s.BeforeControlPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) GetTax() *int32 {
	return s.Tax
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) GetSupplyPrice() *int32 {
	return s.SupplyPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) GetBasicCabinPrice() *int32 {
	return s.BasicCabinPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) GetBuildPrice() *int32 {
	return s.BuildPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) GetOilPrice() *int32 {
	return s.OilPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) GetFirstStandardPrice() *int32 {
	return s.FirstStandardPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) GetBusinessStandardPrice() *int32 {
	return s.BusinessStandardPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) GetCommonStandardPrice() *int32 {
	return s.CommonStandardPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) GetInterTicketPrice() *int32 {
	return s.InterTicketPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) GetSubtractedPrice() *int32 {
	return s.SubtractedPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) GetOriginCommonPrice() *int32 {
	return s.OriginCommonPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) GetDynamicPromotionPrice() *int32 {
	return s.DynamicPromotionPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) GetInstallmentNum() *int32 {
	return s.InstallmentNum
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) GetInstallmentPrice() *float64 {
	return s.InstallmentPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) GetCompetitionDynamicPrice() *int32 {
	return s.CompetitionDynamicPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) GetCompetitionPromotionPrice() *int32 {
	return s.CompetitionPromotionPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) GetMinBeforeControlPriceOfNormal() *int32 {
	return s.MinBeforeControlPriceOfNormal
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) GetPriceShowInfo() *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPricePriceShowInfo {
	return s.PriceShowInfo
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) SetFloorPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice {
	s.FloorPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) SetTicketPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice {
	s.TicketPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) SetSellPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice {
	s.SellPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) SetOriginalSellPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice {
	s.OriginalSellPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) SetBaseTotalPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice {
	s.BaseTotalPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) SetBeforeControlPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice {
	s.BeforeControlPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) SetTax(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice {
	s.Tax = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) SetSupplyPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice {
	s.SupplyPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) SetBasicCabinPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice {
	s.BasicCabinPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) SetBuildPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice {
	s.BuildPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) SetOilPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice {
	s.OilPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) SetFirstStandardPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice {
	s.FirstStandardPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) SetBusinessStandardPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice {
	s.BusinessStandardPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) SetCommonStandardPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice {
	s.CommonStandardPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) SetInterTicketPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice {
	s.InterTicketPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) SetSubtractedPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice {
	s.SubtractedPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) SetOriginCommonPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice {
	s.OriginCommonPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) SetDynamicPromotionPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice {
	s.DynamicPromotionPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) SetInstallmentNum(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice {
	s.InstallmentNum = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) SetInstallmentPrice(v float64) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice {
	s.InstallmentPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) SetCompetitionDynamicPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice {
	s.CompetitionDynamicPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) SetCompetitionPromotionPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice {
	s.CompetitionPromotionPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) SetMinBeforeControlPriceOfNormal(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice {
	s.MinBeforeControlPriceOfNormal = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) SetPriceShowInfo(v *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPricePriceShowInfo) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice {
	s.PriceShowInfo = v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPrice) Validate() error {
	return dara.Validate(s)
}

type ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPricePriceShowInfo struct {
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

func (s ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPricePriceShowInfo) String() string {
	return dara.Prettify(s)
}

func (s ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPricePriceShowInfo) GoString() string {
	return s.String()
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPricePriceShowInfo) GetDiscountInfo() *string {
	return s.DiscountInfo
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPricePriceShowInfo) GetDiscountNum() *float64 {
	return s.DiscountNum
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPricePriceShowInfo) GetShowTicketPrice() *bool {
	return s.ShowTicketPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPricePriceShowInfo) SetDiscountInfo(v string) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPricePriceShowInfo {
	s.DiscountInfo = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPricePriceShowInfo) SetDiscountNum(v float64) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPricePriceShowInfo {
	s.DiscountNum = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPricePriceShowInfo) SetShowTicketPrice(v bool) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPricePriceShowInfo {
	s.ShowTicketPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSearchPricePriceShowInfo) Validate() error {
	return dara.Validate(s)
}

type ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceList struct {
	SegmentPosition *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceListSegmentPosition `json:"segment_position,omitempty" xml:"segment_position,omitempty" type:"Struct"`
	SearchPrice     *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice     `json:"search_price,omitempty" xml:"search_price,omitempty" type:"Struct"`
}

func (s ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceList) String() string {
	return dara.Prettify(s)
}

func (s ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceList) GoString() string {
	return s.String()
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceList) GetSegmentPosition() *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceListSegmentPosition {
	return s.SegmentPosition
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceList) GetSearchPrice() *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice {
	return s.SearchPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceList) SetSegmentPosition(v *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceListSegmentPosition) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceList {
	s.SegmentPosition = v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceList) SetSearchPrice(v *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceList {
	s.SearchPrice = v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceList) Validate() error {
	return dara.Validate(s)
}

type ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceListSegmentPosition struct {
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 0
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
}

func (s ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceListSegmentPosition) String() string {
	return dara.Prettify(s)
}

func (s ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceListSegmentPosition) GoString() string {
	return s.String()
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceListSegmentPosition) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceListSegmentPosition) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceListSegmentPosition) SetJourneyIndex(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceListSegmentPosition {
	s.JourneyIndex = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceListSegmentPosition) SetSegmentIndex(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceListSegmentPosition {
	s.SegmentIndex = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceListSegmentPosition) Validate() error {
	return dara.Validate(s)
}

type ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice struct {
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

func (s ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice) String() string {
	return dara.Prettify(s)
}

func (s ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice) GoString() string {
	return s.String()
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice) GetSellPrice() *int32 {
	return s.SellPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice) GetTicketPrice() *int32 {
	return s.TicketPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice) GetTax() *int32 {
	return s.Tax
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice) SetSellPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice {
	s.SellPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice) SetTicketPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice {
	s.TicketPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice) SetTax(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice {
	s.Tax = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceListSearchPrice) Validate() error {
	return dara.Validate(s)
}
