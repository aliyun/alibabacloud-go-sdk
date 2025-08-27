// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue interface {
	dara.Model
	String() string
	GoString() string
	SetFloorPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue
	GetFloorPrice() *int32
	SetTicketPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue
	GetTicketPrice() *int32
	SetSellPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue
	GetSellPrice() *int32
	SetOriginalSellPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue
	GetOriginalSellPrice() *int32
	SetBaseTotalPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue
	GetBaseTotalPrice() *int32
	SetBeforeControlPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue
	GetBeforeControlPrice() *int32
	SetTax(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue
	GetTax() *int32
	SetSupplyPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue
	GetSupplyPrice() *int32
	SetBasicCabinPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue
	GetBasicCabinPrice() *int32
	SetBuildPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue
	GetBuildPrice() *int32
	SetOilPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue
	GetOilPrice() *int32
	SetFirstStandardPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue
	GetFirstStandardPrice() *int32
	SetBusinessStandardPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue
	GetBusinessStandardPrice() *int32
	SetCommonStandardPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue
	GetCommonStandardPrice() *int32
	SetInterTicketPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue
	GetInterTicketPrice() *int32
	SetSubtractedPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue
	GetSubtractedPrice() *int32
	SetOriginCommonPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue
	GetOriginCommonPrice() *int32
	SetDynamicPromotionPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue
	GetDynamicPromotionPrice() *int32
	SetInstallmentNum(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue
	GetInstallmentNum() *int32
	SetInstallmentPrice(v float64) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue
	GetInstallmentPrice() *float64
	SetCompetitionDynamicPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue
	GetCompetitionDynamicPrice() *int32
	SetCompetitionPromotionPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue
	GetCompetitionPromotionPrice() *int32
	SetMinBeforeControlPriceOfNormal(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue
	GetMinBeforeControlPriceOfNormal() *int32
	SetPriceShowInfo(v *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue
	GetPriceShowInfo() *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo
}

type ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue struct {
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
	InterTicketPrice              *int32                                                                                  `json:"inter_ticket_price,omitempty" xml:"inter_ticket_price,omitempty"`
	SubtractedPrice               *int32                                                                                  `json:"subtracted_price,omitempty" xml:"subtracted_price,omitempty"`
	OriginCommonPrice             *int32                                                                                  `json:"origin_common_price,omitempty" xml:"origin_common_price,omitempty"`
	DynamicPromotionPrice         *int32                                                                                  `json:"dynamic_promotion_price,omitempty" xml:"dynamic_promotion_price,omitempty"`
	InstallmentNum                *int32                                                                                  `json:"installment_num,omitempty" xml:"installment_num,omitempty"`
	InstallmentPrice              *float64                                                                                `json:"installment_price,omitempty" xml:"installment_price,omitempty"`
	CompetitionDynamicPrice       *int32                                                                                  `json:"competition_dynamic_price,omitempty" xml:"competition_dynamic_price,omitempty"`
	CompetitionPromotionPrice     *int32                                                                                  `json:"competition_promotion_price,omitempty" xml:"competition_promotion_price,omitempty"`
	MinBeforeControlPriceOfNormal *int32                                                                                  `json:"min_before_control_price_of_normal,omitempty" xml:"min_before_control_price_of_normal,omitempty"`
	PriceShowInfo                 *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo `json:"price_show_info,omitempty" xml:"price_show_info,omitempty" type:"Struct"`
}

func (s ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) GoString() string {
	return s.String()
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) GetFloorPrice() *int32 {
	return s.FloorPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) GetTicketPrice() *int32 {
	return s.TicketPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) GetSellPrice() *int32 {
	return s.SellPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) GetOriginalSellPrice() *int32 {
	return s.OriginalSellPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) GetBaseTotalPrice() *int32 {
	return s.BaseTotalPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) GetBeforeControlPrice() *int32 {
	return s.BeforeControlPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) GetTax() *int32 {
	return s.Tax
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) GetSupplyPrice() *int32 {
	return s.SupplyPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) GetBasicCabinPrice() *int32 {
	return s.BasicCabinPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) GetBuildPrice() *int32 {
	return s.BuildPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) GetOilPrice() *int32 {
	return s.OilPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) GetFirstStandardPrice() *int32 {
	return s.FirstStandardPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) GetBusinessStandardPrice() *int32 {
	return s.BusinessStandardPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) GetCommonStandardPrice() *int32 {
	return s.CommonStandardPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) GetInterTicketPrice() *int32 {
	return s.InterTicketPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) GetSubtractedPrice() *int32 {
	return s.SubtractedPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) GetOriginCommonPrice() *int32 {
	return s.OriginCommonPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) GetDynamicPromotionPrice() *int32 {
	return s.DynamicPromotionPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) GetInstallmentNum() *int32 {
	return s.InstallmentNum
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) GetInstallmentPrice() *float64 {
	return s.InstallmentPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) GetCompetitionDynamicPrice() *int32 {
	return s.CompetitionDynamicPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) GetCompetitionPromotionPrice() *int32 {
	return s.CompetitionPromotionPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) GetMinBeforeControlPriceOfNormal() *int32 {
	return s.MinBeforeControlPriceOfNormal
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) GetPriceShowInfo() *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo {
	return s.PriceShowInfo
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) SetFloorPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue {
	s.FloorPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) SetTicketPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue {
	s.TicketPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) SetSellPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue {
	s.SellPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) SetOriginalSellPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue {
	s.OriginalSellPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) SetBaseTotalPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue {
	s.BaseTotalPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) SetBeforeControlPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue {
	s.BeforeControlPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) SetTax(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue {
	s.Tax = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) SetSupplyPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue {
	s.SupplyPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) SetBasicCabinPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue {
	s.BasicCabinPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) SetBuildPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue {
	s.BuildPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) SetOilPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue {
	s.OilPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) SetFirstStandardPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue {
	s.FirstStandardPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) SetBusinessStandardPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue {
	s.BusinessStandardPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) SetCommonStandardPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue {
	s.CommonStandardPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) SetInterTicketPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue {
	s.InterTicketPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) SetSubtractedPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue {
	s.SubtractedPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) SetOriginCommonPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue {
	s.OriginCommonPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) SetDynamicPromotionPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue {
	s.DynamicPromotionPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) SetInstallmentNum(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue {
	s.InstallmentNum = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) SetInstallmentPrice(v float64) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue {
	s.InstallmentPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) SetCompetitionDynamicPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue {
	s.CompetitionDynamicPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) SetCompetitionPromotionPrice(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue {
	s.CompetitionPromotionPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) SetMinBeforeControlPriceOfNormal(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue {
	s.MinBeforeControlPriceOfNormal = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) SetPriceShowInfo(v *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue {
	s.PriceShowInfo = v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValue) Validate() error {
	return dara.Validate(s)
}

type ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo struct {
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

func (s ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo) String() string {
	return dara.Prettify(s)
}

func (s ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo) GoString() string {
	return s.String()
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo) GetDiscountInfo() *string {
	return s.DiscountInfo
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo) GetDiscountNum() *float64 {
	return s.DiscountNum
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo) GetShowTicketPrice() *bool {
	return s.ShowTicketPrice
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo) SetDiscountInfo(v string) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo {
	s.DiscountInfo = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo) SetDiscountNum(v float64) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo {
	s.DiscountNum = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo) SetShowTicketPrice(v bool) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo {
	s.ShowTicketPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo) Validate() error {
	return dara.Validate(s)
}
