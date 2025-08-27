// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleFlightItemListItemListShoppingItemMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ModuleFlightItemListItemListShoppingItemMapValue
	GetId() *string
	SetCabinQuantity(v map[string]*ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue) *ModuleFlightItemListItemListShoppingItemMapValue
	GetCabinQuantity() map[string]*ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue
	SetSearchPrice(v *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) *ModuleFlightItemListItemListShoppingItemMapValue
	GetSearchPrice() *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice
	SetSegmentPrice(v map[string]*ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) *ModuleFlightItemListItemListShoppingItemMapValue
	GetSegmentPrice() map[string]*ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue
}

type ModuleFlightItemListItemListShoppingItemMapValue struct {
	// id
	//
	// example:
	//
	// ADT07df0bd9-f803-4a50-8449-f4bd675d9939
	Id            *string                                                                        `json:"id,omitempty" xml:"id,omitempty"`
	CabinQuantity map[string]*ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue `json:"cabin_quantity,omitempty" xml:"cabin_quantity,omitempty"`
	SearchPrice   *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice                   `json:"search_price,omitempty" xml:"search_price,omitempty" type:"Struct"`
	SegmentPrice  map[string]*ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue  `json:"segment_price,omitempty" xml:"segment_price,omitempty"`
}

func (s ModuleFlightItemListItemListShoppingItemMapValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleFlightItemListItemListShoppingItemMapValue) GoString() string {
	return s.String()
}

func (s *ModuleFlightItemListItemListShoppingItemMapValue) GetId() *string {
	return s.Id
}

func (s *ModuleFlightItemListItemListShoppingItemMapValue) GetCabinQuantity() map[string]*ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue {
	return s.CabinQuantity
}

func (s *ModuleFlightItemListItemListShoppingItemMapValue) GetSearchPrice() *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice {
	return s.SearchPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValue) GetSegmentPrice() map[string]*ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue {
	return s.SegmentPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValue) SetId(v string) *ModuleFlightItemListItemListShoppingItemMapValue {
	s.Id = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValue) SetCabinQuantity(v map[string]*ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue) *ModuleFlightItemListItemListShoppingItemMapValue {
	s.CabinQuantity = v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValue) SetSearchPrice(v *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) *ModuleFlightItemListItemListShoppingItemMapValue {
	s.SearchPrice = v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValue) SetSegmentPrice(v map[string]*ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) *ModuleFlightItemListItemListShoppingItemMapValue {
	s.SegmentPrice = v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValue) Validate() error {
	return dara.Validate(s)
}

type ModuleFlightItemListItemListShoppingItemMapValueSearchPrice struct {
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
	MinBeforeControlPriceOfNormal *int32                                                                    `json:"min_before_control_price_of_normal,omitempty" xml:"min_before_control_price_of_normal,omitempty"`
	PriceShowInfo                 *ModuleFlightItemListItemListShoppingItemMapValueSearchPricePriceShowInfo `json:"price_show_info,omitempty" xml:"price_show_info,omitempty" type:"Struct"`
}

func (s ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) String() string {
	return dara.Prettify(s)
}

func (s ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) GoString() string {
	return s.String()
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) GetFloorPrice() *int32 {
	return s.FloorPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) GetTicketPrice() *int32 {
	return s.TicketPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) GetSellPrice() *int32 {
	return s.SellPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) GetOriginalSellPrice() *int32 {
	return s.OriginalSellPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) GetBaseTotalPrice() *int32 {
	return s.BaseTotalPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) GetBeforeControlPrice() *int32 {
	return s.BeforeControlPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) GetTax() *int32 {
	return s.Tax
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) GetSupplyPrice() *int32 {
	return s.SupplyPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) GetBasicCabinPrice() *int32 {
	return s.BasicCabinPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) GetBuildPrice() *int32 {
	return s.BuildPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) GetOilPrice() *int32 {
	return s.OilPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) GetFirstStandardPrice() *int32 {
	return s.FirstStandardPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) GetBusinessStandardPrice() *int32 {
	return s.BusinessStandardPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) GetCommonStandardPrice() *int32 {
	return s.CommonStandardPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) GetInterTicketPrice() *int32 {
	return s.InterTicketPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) GetSubtractedPrice() *int32 {
	return s.SubtractedPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) GetOriginCommonPrice() *int32 {
	return s.OriginCommonPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) GetDynamicPromotionPrice() *int32 {
	return s.DynamicPromotionPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) GetInstallmentNum() *int32 {
	return s.InstallmentNum
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) GetInstallmentPrice() *float64 {
	return s.InstallmentPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) GetCompetitionDynamicPrice() *int32 {
	return s.CompetitionDynamicPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) GetCompetitionPromotionPrice() *int32 {
	return s.CompetitionPromotionPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) GetMinBeforeControlPriceOfNormal() *int32 {
	return s.MinBeforeControlPriceOfNormal
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) GetPriceShowInfo() *ModuleFlightItemListItemListShoppingItemMapValueSearchPricePriceShowInfo {
	return s.PriceShowInfo
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) SetFloorPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice {
	s.FloorPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) SetTicketPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice {
	s.TicketPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) SetSellPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice {
	s.SellPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) SetOriginalSellPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice {
	s.OriginalSellPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) SetBaseTotalPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice {
	s.BaseTotalPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) SetBeforeControlPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice {
	s.BeforeControlPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) SetTax(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice {
	s.Tax = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) SetSupplyPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice {
	s.SupplyPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) SetBasicCabinPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice {
	s.BasicCabinPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) SetBuildPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice {
	s.BuildPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) SetOilPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice {
	s.OilPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) SetFirstStandardPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice {
	s.FirstStandardPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) SetBusinessStandardPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice {
	s.BusinessStandardPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) SetCommonStandardPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice {
	s.CommonStandardPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) SetInterTicketPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice {
	s.InterTicketPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) SetSubtractedPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice {
	s.SubtractedPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) SetOriginCommonPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice {
	s.OriginCommonPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) SetDynamicPromotionPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice {
	s.DynamicPromotionPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) SetInstallmentNum(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice {
	s.InstallmentNum = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) SetInstallmentPrice(v float64) *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice {
	s.InstallmentPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) SetCompetitionDynamicPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice {
	s.CompetitionDynamicPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) SetCompetitionPromotionPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice {
	s.CompetitionPromotionPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) SetMinBeforeControlPriceOfNormal(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice {
	s.MinBeforeControlPriceOfNormal = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) SetPriceShowInfo(v *ModuleFlightItemListItemListShoppingItemMapValueSearchPricePriceShowInfo) *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice {
	s.PriceShowInfo = v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPrice) Validate() error {
	return dara.Validate(s)
}

type ModuleFlightItemListItemListShoppingItemMapValueSearchPricePriceShowInfo struct {
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

func (s ModuleFlightItemListItemListShoppingItemMapValueSearchPricePriceShowInfo) String() string {
	return dara.Prettify(s)
}

func (s ModuleFlightItemListItemListShoppingItemMapValueSearchPricePriceShowInfo) GoString() string {
	return s.String()
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPricePriceShowInfo) GetDiscountInfo() *string {
	return s.DiscountInfo
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPricePriceShowInfo) GetDiscountNum() *float64 {
	return s.DiscountNum
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPricePriceShowInfo) GetShowTicketPrice() *bool {
	return s.ShowTicketPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPricePriceShowInfo) SetDiscountInfo(v string) *ModuleFlightItemListItemListShoppingItemMapValueSearchPricePriceShowInfo {
	s.DiscountInfo = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPricePriceShowInfo) SetDiscountNum(v float64) *ModuleFlightItemListItemListShoppingItemMapValueSearchPricePriceShowInfo {
	s.DiscountNum = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPricePriceShowInfo) SetShowTicketPrice(v bool) *ModuleFlightItemListItemListShoppingItemMapValueSearchPricePriceShowInfo {
	s.ShowTicketPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSearchPricePriceShowInfo) Validate() error {
	return dara.Validate(s)
}
