// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValue
	GetId() *string
	SetCabinQuantity(v map[string]*ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValue
	GetCabinQuantity() map[string]*ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue
	SetSearchPrice(v *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValue
	GetSearchPrice() *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice
	SetSegmentPrice(v map[string]*ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValue
	GetSegmentPrice() map[string]*ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue
}

type ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValue struct {
	// id
	//
	// example:
	//
	// ADT07df0bd9-f803-4a50-8449-f4bd675d9939
	Id            *string                                                                                     `json:"id,omitempty" xml:"id,omitempty"`
	CabinQuantity map[string]*ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue `json:"cabin_quantity,omitempty" xml:"cabin_quantity,omitempty"`
	SearchPrice   *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice                   `json:"search_price,omitempty" xml:"search_price,omitempty" type:"Struct"`
	SegmentPrice  map[string]*ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue  `json:"segment_price,omitempty" xml:"segment_price,omitempty"`
}

func (s ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValue) GoString() string {
	return s.String()
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValue) GetId() *string {
	return s.Id
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValue) GetCabinQuantity() map[string]*ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue {
	return s.CabinQuantity
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValue) GetSearchPrice() *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice {
	return s.SearchPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValue) GetSegmentPrice() map[string]*ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue {
	return s.SegmentPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValue) SetId(v string) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValue {
	s.Id = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValue) SetCabinQuantity(v map[string]*ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValue {
	s.CabinQuantity = v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValue) SetSearchPrice(v *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValue {
	s.SearchPrice = v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValue) SetSegmentPrice(v map[string]*ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValue {
	s.SegmentPrice = v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValue) Validate() error {
	return dara.Validate(s)
}

type ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice struct {
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
	MinBeforeControlPriceOfNormal *int32                                                                                 `json:"min_before_control_price_of_normal,omitempty" xml:"min_before_control_price_of_normal,omitempty"`
	PriceShowInfo                 *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPricePriceShowInfo `json:"price_show_info,omitempty" xml:"price_show_info,omitempty" type:"Struct"`
}

func (s ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) String() string {
	return dara.Prettify(s)
}

func (s ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) GoString() string {
	return s.String()
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) GetFloorPrice() *int32 {
	return s.FloorPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) GetTicketPrice() *int32 {
	return s.TicketPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) GetSellPrice() *int32 {
	return s.SellPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) GetOriginalSellPrice() *int32 {
	return s.OriginalSellPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) GetBaseTotalPrice() *int32 {
	return s.BaseTotalPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) GetBeforeControlPrice() *int32 {
	return s.BeforeControlPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) GetTax() *int32 {
	return s.Tax
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) GetSupplyPrice() *int32 {
	return s.SupplyPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) GetBasicCabinPrice() *int32 {
	return s.BasicCabinPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) GetBuildPrice() *int32 {
	return s.BuildPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) GetOilPrice() *int32 {
	return s.OilPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) GetFirstStandardPrice() *int32 {
	return s.FirstStandardPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) GetBusinessStandardPrice() *int32 {
	return s.BusinessStandardPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) GetCommonStandardPrice() *int32 {
	return s.CommonStandardPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) GetInterTicketPrice() *int32 {
	return s.InterTicketPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) GetSubtractedPrice() *int32 {
	return s.SubtractedPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) GetOriginCommonPrice() *int32 {
	return s.OriginCommonPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) GetDynamicPromotionPrice() *int32 {
	return s.DynamicPromotionPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) GetInstallmentNum() *int32 {
	return s.InstallmentNum
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) GetInstallmentPrice() *float64 {
	return s.InstallmentPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) GetCompetitionDynamicPrice() *int32 {
	return s.CompetitionDynamicPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) GetCompetitionPromotionPrice() *int32 {
	return s.CompetitionPromotionPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) GetMinBeforeControlPriceOfNormal() *int32 {
	return s.MinBeforeControlPriceOfNormal
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) GetPriceShowInfo() *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPricePriceShowInfo {
	return s.PriceShowInfo
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) SetFloorPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice {
	s.FloorPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) SetTicketPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice {
	s.TicketPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) SetSellPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice {
	s.SellPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) SetOriginalSellPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice {
	s.OriginalSellPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) SetBaseTotalPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice {
	s.BaseTotalPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) SetBeforeControlPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice {
	s.BeforeControlPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) SetTax(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice {
	s.Tax = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) SetSupplyPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice {
	s.SupplyPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) SetBasicCabinPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice {
	s.BasicCabinPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) SetBuildPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice {
	s.BuildPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) SetOilPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice {
	s.OilPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) SetFirstStandardPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice {
	s.FirstStandardPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) SetBusinessStandardPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice {
	s.BusinessStandardPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) SetCommonStandardPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice {
	s.CommonStandardPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) SetInterTicketPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice {
	s.InterTicketPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) SetSubtractedPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice {
	s.SubtractedPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) SetOriginCommonPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice {
	s.OriginCommonPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) SetDynamicPromotionPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice {
	s.DynamicPromotionPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) SetInstallmentNum(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice {
	s.InstallmentNum = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) SetInstallmentPrice(v float64) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice {
	s.InstallmentPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) SetCompetitionDynamicPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice {
	s.CompetitionDynamicPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) SetCompetitionPromotionPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice {
	s.CompetitionPromotionPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) SetMinBeforeControlPriceOfNormal(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice {
	s.MinBeforeControlPriceOfNormal = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) SetPriceShowInfo(v *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPricePriceShowInfo) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice {
	s.PriceShowInfo = v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPrice) Validate() error {
	return dara.Validate(s)
}

type ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPricePriceShowInfo struct {
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

func (s ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPricePriceShowInfo) String() string {
	return dara.Prettify(s)
}

func (s ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPricePriceShowInfo) GoString() string {
	return s.String()
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPricePriceShowInfo) GetDiscountInfo() *string {
	return s.DiscountInfo
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPricePriceShowInfo) GetDiscountNum() *float64 {
	return s.DiscountNum
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPricePriceShowInfo) GetShowTicketPrice() *bool {
	return s.ShowTicketPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPricePriceShowInfo) SetDiscountInfo(v string) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPricePriceShowInfo {
	s.DiscountInfo = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPricePriceShowInfo) SetDiscountNum(v float64) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPricePriceShowInfo {
	s.DiscountNum = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPricePriceShowInfo) SetShowTicketPrice(v bool) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPricePriceShowInfo {
	s.ShowTicketPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSearchPricePriceShowInfo) Validate() error {
	return dara.Validate(s)
}
