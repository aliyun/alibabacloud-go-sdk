// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleFlightItemListItemListSubItemsShoppingItemMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ModuleFlightItemListItemListSubItemsShoppingItemMapValue
	GetId() *string
	SetCabinQuantity(v map[string]*ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue) *ModuleFlightItemListItemListSubItemsShoppingItemMapValue
	GetCabinQuantity() map[string]*ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue
	SetSearchPrice(v *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) *ModuleFlightItemListItemListSubItemsShoppingItemMapValue
	GetSearchPrice() *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice
	SetSegmentPrice(v map[string]*ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) *ModuleFlightItemListItemListSubItemsShoppingItemMapValue
	GetSegmentPrice() map[string]*ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue
}

type ModuleFlightItemListItemListSubItemsShoppingItemMapValue struct {
	// id
	//
	// example:
	//
	// ADT07df0bd9-f803-4a50-8449-f4bd675d9939
	Id            *string                                                                                `json:"id,omitempty" xml:"id,omitempty"`
	CabinQuantity map[string]*ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue `json:"cabin_quantity,omitempty" xml:"cabin_quantity,omitempty"`
	SearchPrice   *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice                   `json:"search_price,omitempty" xml:"search_price,omitempty" type:"Struct"`
	SegmentPrice  map[string]*ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue  `json:"segment_price,omitempty" xml:"segment_price,omitempty"`
}

func (s ModuleFlightItemListItemListSubItemsShoppingItemMapValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleFlightItemListItemListSubItemsShoppingItemMapValue) GoString() string {
	return s.String()
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValue) GetId() *string {
	return s.Id
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValue) GetCabinQuantity() map[string]*ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue {
	return s.CabinQuantity
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValue) GetSearchPrice() *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice {
	return s.SearchPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValue) GetSegmentPrice() map[string]*ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue {
	return s.SegmentPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValue) SetId(v string) *ModuleFlightItemListItemListSubItemsShoppingItemMapValue {
	s.Id = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValue) SetCabinQuantity(v map[string]*ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue) *ModuleFlightItemListItemListSubItemsShoppingItemMapValue {
	s.CabinQuantity = v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValue) SetSearchPrice(v *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) *ModuleFlightItemListItemListSubItemsShoppingItemMapValue {
	s.SearchPrice = v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValue) SetSegmentPrice(v map[string]*ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) *ModuleFlightItemListItemListSubItemsShoppingItemMapValue {
	s.SegmentPrice = v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValue) Validate() error {
	return dara.Validate(s)
}

type ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice struct {
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
	MinBeforeControlPriceOfNormal *int32                                                                            `json:"min_before_control_price_of_normal,omitempty" xml:"min_before_control_price_of_normal,omitempty"`
	PriceShowInfo                 *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPricePriceShowInfo `json:"price_show_info,omitempty" xml:"price_show_info,omitempty" type:"Struct"`
}

func (s ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) String() string {
	return dara.Prettify(s)
}

func (s ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) GoString() string {
	return s.String()
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) GetFloorPrice() *int32 {
	return s.FloorPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) GetTicketPrice() *int32 {
	return s.TicketPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) GetSellPrice() *int32 {
	return s.SellPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) GetOriginalSellPrice() *int32 {
	return s.OriginalSellPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) GetBaseTotalPrice() *int32 {
	return s.BaseTotalPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) GetBeforeControlPrice() *int32 {
	return s.BeforeControlPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) GetTax() *int32 {
	return s.Tax
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) GetSupplyPrice() *int32 {
	return s.SupplyPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) GetBasicCabinPrice() *int32 {
	return s.BasicCabinPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) GetBuildPrice() *int32 {
	return s.BuildPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) GetOilPrice() *int32 {
	return s.OilPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) GetFirstStandardPrice() *int32 {
	return s.FirstStandardPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) GetBusinessStandardPrice() *int32 {
	return s.BusinessStandardPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) GetCommonStandardPrice() *int32 {
	return s.CommonStandardPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) GetInterTicketPrice() *int32 {
	return s.InterTicketPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) GetSubtractedPrice() *int32 {
	return s.SubtractedPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) GetOriginCommonPrice() *int32 {
	return s.OriginCommonPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) GetDynamicPromotionPrice() *int32 {
	return s.DynamicPromotionPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) GetInstallmentNum() *int32 {
	return s.InstallmentNum
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) GetInstallmentPrice() *float64 {
	return s.InstallmentPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) GetCompetitionDynamicPrice() *int32 {
	return s.CompetitionDynamicPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) GetCompetitionPromotionPrice() *int32 {
	return s.CompetitionPromotionPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) GetMinBeforeControlPriceOfNormal() *int32 {
	return s.MinBeforeControlPriceOfNormal
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) GetPriceShowInfo() *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPricePriceShowInfo {
	return s.PriceShowInfo
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) SetFloorPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice {
	s.FloorPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) SetTicketPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice {
	s.TicketPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) SetSellPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice {
	s.SellPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) SetOriginalSellPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice {
	s.OriginalSellPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) SetBaseTotalPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice {
	s.BaseTotalPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) SetBeforeControlPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice {
	s.BeforeControlPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) SetTax(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice {
	s.Tax = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) SetSupplyPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice {
	s.SupplyPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) SetBasicCabinPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice {
	s.BasicCabinPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) SetBuildPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice {
	s.BuildPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) SetOilPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice {
	s.OilPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) SetFirstStandardPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice {
	s.FirstStandardPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) SetBusinessStandardPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice {
	s.BusinessStandardPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) SetCommonStandardPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice {
	s.CommonStandardPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) SetInterTicketPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice {
	s.InterTicketPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) SetSubtractedPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice {
	s.SubtractedPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) SetOriginCommonPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice {
	s.OriginCommonPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) SetDynamicPromotionPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice {
	s.DynamicPromotionPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) SetInstallmentNum(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice {
	s.InstallmentNum = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) SetInstallmentPrice(v float64) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice {
	s.InstallmentPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) SetCompetitionDynamicPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice {
	s.CompetitionDynamicPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) SetCompetitionPromotionPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice {
	s.CompetitionPromotionPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) SetMinBeforeControlPriceOfNormal(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice {
	s.MinBeforeControlPriceOfNormal = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) SetPriceShowInfo(v *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPricePriceShowInfo) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice {
	s.PriceShowInfo = v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPrice) Validate() error {
	return dara.Validate(s)
}

type ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPricePriceShowInfo struct {
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

func (s ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPricePriceShowInfo) String() string {
	return dara.Prettify(s)
}

func (s ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPricePriceShowInfo) GoString() string {
	return s.String()
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPricePriceShowInfo) GetDiscountInfo() *string {
	return s.DiscountInfo
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPricePriceShowInfo) GetDiscountNum() *float64 {
	return s.DiscountNum
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPricePriceShowInfo) GetShowTicketPrice() *bool {
	return s.ShowTicketPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPricePriceShowInfo) SetDiscountInfo(v string) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPricePriceShowInfo {
	s.DiscountInfo = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPricePriceShowInfo) SetDiscountNum(v float64) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPricePriceShowInfo {
	s.DiscountNum = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPricePriceShowInfo) SetShowTicketPrice(v bool) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPricePriceShowInfo {
	s.ShowTicketPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSearchPricePriceShowInfo) Validate() error {
	return dara.Validate(s)
}
