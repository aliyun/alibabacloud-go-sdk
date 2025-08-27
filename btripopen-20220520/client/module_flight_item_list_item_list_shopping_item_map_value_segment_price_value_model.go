// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue interface {
	dara.Model
	String() string
	GoString() string
	SetFloorPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue
	GetFloorPrice() *int32
	SetTicketPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue
	GetTicketPrice() *int32
	SetSellPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue
	GetSellPrice() *int32
	SetOriginalSellPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue
	GetOriginalSellPrice() *int32
	SetBaseTotalPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue
	GetBaseTotalPrice() *int32
	SetBeforeControlPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue
	GetBeforeControlPrice() *int32
	SetTax(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue
	GetTax() *int32
	SetSupplyPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue
	GetSupplyPrice() *int32
	SetBasicCabinPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue
	GetBasicCabinPrice() *int32
	SetBuildPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue
	GetBuildPrice() *int32
	SetOilPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue
	GetOilPrice() *int32
	SetFirstStandardPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue
	GetFirstStandardPrice() *int32
	SetBusinessStandardPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue
	GetBusinessStandardPrice() *int32
	SetCommonStandardPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue
	GetCommonStandardPrice() *int32
	SetInterTicketPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue
	GetInterTicketPrice() *int32
	SetSubtractedPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue
	GetSubtractedPrice() *int32
	SetOriginCommonPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue
	GetOriginCommonPrice() *int32
	SetDynamicPromotionPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue
	GetDynamicPromotionPrice() *int32
	SetInstallmentNum(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue
	GetInstallmentNum() *int32
	SetInstallmentPrice(v float64) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue
	GetInstallmentPrice() *float64
	SetCompetitionDynamicPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue
	GetCompetitionDynamicPrice() *int32
	SetCompetitionPromotionPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue
	GetCompetitionPromotionPrice() *int32
	SetMinBeforeControlPriceOfNormal(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue
	GetMinBeforeControlPriceOfNormal() *int32
	SetPriceShowInfo(v *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue
	GetPriceShowInfo() *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo
}

type ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue struct {
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
	InterTicketPrice              *int32                                                                          `json:"inter_ticket_price,omitempty" xml:"inter_ticket_price,omitempty"`
	SubtractedPrice               *int32                                                                          `json:"subtracted_price,omitempty" xml:"subtracted_price,omitempty"`
	OriginCommonPrice             *int32                                                                          `json:"origin_common_price,omitempty" xml:"origin_common_price,omitempty"`
	DynamicPromotionPrice         *int32                                                                          `json:"dynamic_promotion_price,omitempty" xml:"dynamic_promotion_price,omitempty"`
	InstallmentNum                *int32                                                                          `json:"installment_num,omitempty" xml:"installment_num,omitempty"`
	InstallmentPrice              *float64                                                                        `json:"installment_price,omitempty" xml:"installment_price,omitempty"`
	CompetitionDynamicPrice       *int32                                                                          `json:"competition_dynamic_price,omitempty" xml:"competition_dynamic_price,omitempty"`
	CompetitionPromotionPrice     *int32                                                                          `json:"competition_promotion_price,omitempty" xml:"competition_promotion_price,omitempty"`
	MinBeforeControlPriceOfNormal *int32                                                                          `json:"min_before_control_price_of_normal,omitempty" xml:"min_before_control_price_of_normal,omitempty"`
	PriceShowInfo                 *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo `json:"price_show_info,omitempty" xml:"price_show_info,omitempty" type:"Struct"`
}

func (s ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) GoString() string {
	return s.String()
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) GetFloorPrice() *int32 {
	return s.FloorPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) GetTicketPrice() *int32 {
	return s.TicketPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) GetSellPrice() *int32 {
	return s.SellPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) GetOriginalSellPrice() *int32 {
	return s.OriginalSellPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) GetBaseTotalPrice() *int32 {
	return s.BaseTotalPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) GetBeforeControlPrice() *int32 {
	return s.BeforeControlPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) GetTax() *int32 {
	return s.Tax
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) GetSupplyPrice() *int32 {
	return s.SupplyPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) GetBasicCabinPrice() *int32 {
	return s.BasicCabinPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) GetBuildPrice() *int32 {
	return s.BuildPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) GetOilPrice() *int32 {
	return s.OilPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) GetFirstStandardPrice() *int32 {
	return s.FirstStandardPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) GetBusinessStandardPrice() *int32 {
	return s.BusinessStandardPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) GetCommonStandardPrice() *int32 {
	return s.CommonStandardPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) GetInterTicketPrice() *int32 {
	return s.InterTicketPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) GetSubtractedPrice() *int32 {
	return s.SubtractedPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) GetOriginCommonPrice() *int32 {
	return s.OriginCommonPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) GetDynamicPromotionPrice() *int32 {
	return s.DynamicPromotionPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) GetInstallmentNum() *int32 {
	return s.InstallmentNum
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) GetInstallmentPrice() *float64 {
	return s.InstallmentPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) GetCompetitionDynamicPrice() *int32 {
	return s.CompetitionDynamicPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) GetCompetitionPromotionPrice() *int32 {
	return s.CompetitionPromotionPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) GetMinBeforeControlPriceOfNormal() *int32 {
	return s.MinBeforeControlPriceOfNormal
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) GetPriceShowInfo() *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo {
	return s.PriceShowInfo
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) SetFloorPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue {
	s.FloorPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) SetTicketPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue {
	s.TicketPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) SetSellPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue {
	s.SellPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) SetOriginalSellPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue {
	s.OriginalSellPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) SetBaseTotalPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue {
	s.BaseTotalPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) SetBeforeControlPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue {
	s.BeforeControlPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) SetTax(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue {
	s.Tax = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) SetSupplyPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue {
	s.SupplyPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) SetBasicCabinPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue {
	s.BasicCabinPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) SetBuildPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue {
	s.BuildPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) SetOilPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue {
	s.OilPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) SetFirstStandardPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue {
	s.FirstStandardPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) SetBusinessStandardPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue {
	s.BusinessStandardPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) SetCommonStandardPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue {
	s.CommonStandardPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) SetInterTicketPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue {
	s.InterTicketPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) SetSubtractedPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue {
	s.SubtractedPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) SetOriginCommonPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue {
	s.OriginCommonPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) SetDynamicPromotionPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue {
	s.DynamicPromotionPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) SetInstallmentNum(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue {
	s.InstallmentNum = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) SetInstallmentPrice(v float64) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue {
	s.InstallmentPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) SetCompetitionDynamicPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue {
	s.CompetitionDynamicPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) SetCompetitionPromotionPrice(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue {
	s.CompetitionPromotionPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) SetMinBeforeControlPriceOfNormal(v int32) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue {
	s.MinBeforeControlPriceOfNormal = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) SetPriceShowInfo(v *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue {
	s.PriceShowInfo = v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValue) Validate() error {
	return dara.Validate(s)
}

type ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo struct {
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

func (s ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo) String() string {
	return dara.Prettify(s)
}

func (s ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo) GoString() string {
	return s.String()
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo) GetDiscountInfo() *string {
	return s.DiscountInfo
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo) GetDiscountNum() *float64 {
	return s.DiscountNum
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo) GetShowTicketPrice() *bool {
	return s.ShowTicketPrice
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo) SetDiscountInfo(v string) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo {
	s.DiscountInfo = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo) SetDiscountNum(v float64) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo {
	s.DiscountNum = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo) SetShowTicketPrice(v bool) *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo {
	s.ShowTicketPrice = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo) Validate() error {
	return dara.Validate(s)
}
