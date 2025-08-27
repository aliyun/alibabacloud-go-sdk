// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue interface {
	dara.Model
	String() string
	GoString() string
	SetFloorPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue
	GetFloorPrice() *int32
	SetTicketPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue
	GetTicketPrice() *int32
	SetSellPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue
	GetSellPrice() *int32
	SetOriginalSellPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue
	GetOriginalSellPrice() *int32
	SetBaseTotalPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue
	GetBaseTotalPrice() *int32
	SetBeforeControlPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue
	GetBeforeControlPrice() *int32
	SetTax(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue
	GetTax() *int32
	SetSupplyPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue
	GetSupplyPrice() *int32
	SetBasicCabinPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue
	GetBasicCabinPrice() *int32
	SetBuildPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue
	GetBuildPrice() *int32
	SetOilPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue
	GetOilPrice() *int32
	SetFirstStandardPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue
	GetFirstStandardPrice() *int32
	SetBusinessStandardPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue
	GetBusinessStandardPrice() *int32
	SetCommonStandardPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue
	GetCommonStandardPrice() *int32
	SetInterTicketPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue
	GetInterTicketPrice() *int32
	SetSubtractedPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue
	GetSubtractedPrice() *int32
	SetOriginCommonPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue
	GetOriginCommonPrice() *int32
	SetDynamicPromotionPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue
	GetDynamicPromotionPrice() *int32
	SetInstallmentNum(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue
	GetInstallmentNum() *int32
	SetInstallmentPrice(v float64) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue
	GetInstallmentPrice() *float64
	SetCompetitionDynamicPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue
	GetCompetitionDynamicPrice() *int32
	SetCompetitionPromotionPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue
	GetCompetitionPromotionPrice() *int32
	SetMinBeforeControlPriceOfNormal(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue
	GetMinBeforeControlPriceOfNormal() *int32
	SetPriceShowInfo(v *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValuePriceShowInfo) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue
	GetPriceShowInfo() *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValuePriceShowInfo
}

type ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue struct {
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
	InterTicketPrice              *int32                                                                               `json:"inter_ticket_price,omitempty" xml:"inter_ticket_price,omitempty"`
	SubtractedPrice               *int32                                                                               `json:"subtracted_price,omitempty" xml:"subtracted_price,omitempty"`
	OriginCommonPrice             *int32                                                                               `json:"origin_common_price,omitempty" xml:"origin_common_price,omitempty"`
	DynamicPromotionPrice         *int32                                                                               `json:"dynamic_promotion_price,omitempty" xml:"dynamic_promotion_price,omitempty"`
	InstallmentNum                *int32                                                                               `json:"installment_num,omitempty" xml:"installment_num,omitempty"`
	InstallmentPrice              *float64                                                                             `json:"installment_price,omitempty" xml:"installment_price,omitempty"`
	CompetitionDynamicPrice       *int32                                                                               `json:"competition_dynamic_price,omitempty" xml:"competition_dynamic_price,omitempty"`
	CompetitionPromotionPrice     *int32                                                                               `json:"competition_promotion_price,omitempty" xml:"competition_promotion_price,omitempty"`
	MinBeforeControlPriceOfNormal *int32                                                                               `json:"min_before_control_price_of_normal,omitempty" xml:"min_before_control_price_of_normal,omitempty"`
	PriceShowInfo                 *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValuePriceShowInfo `json:"price_show_info,omitempty" xml:"price_show_info,omitempty" type:"Struct"`
}

func (s ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) GoString() string {
	return s.String()
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) GetFloorPrice() *int32 {
	return s.FloorPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) GetTicketPrice() *int32 {
	return s.TicketPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) GetSellPrice() *int32 {
	return s.SellPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) GetOriginalSellPrice() *int32 {
	return s.OriginalSellPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) GetBaseTotalPrice() *int32 {
	return s.BaseTotalPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) GetBeforeControlPrice() *int32 {
	return s.BeforeControlPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) GetTax() *int32 {
	return s.Tax
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) GetSupplyPrice() *int32 {
	return s.SupplyPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) GetBasicCabinPrice() *int32 {
	return s.BasicCabinPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) GetBuildPrice() *int32 {
	return s.BuildPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) GetOilPrice() *int32 {
	return s.OilPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) GetFirstStandardPrice() *int32 {
	return s.FirstStandardPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) GetBusinessStandardPrice() *int32 {
	return s.BusinessStandardPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) GetCommonStandardPrice() *int32 {
	return s.CommonStandardPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) GetInterTicketPrice() *int32 {
	return s.InterTicketPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) GetSubtractedPrice() *int32 {
	return s.SubtractedPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) GetOriginCommonPrice() *int32 {
	return s.OriginCommonPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) GetDynamicPromotionPrice() *int32 {
	return s.DynamicPromotionPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) GetInstallmentNum() *int32 {
	return s.InstallmentNum
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) GetInstallmentPrice() *float64 {
	return s.InstallmentPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) GetCompetitionDynamicPrice() *int32 {
	return s.CompetitionDynamicPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) GetCompetitionPromotionPrice() *int32 {
	return s.CompetitionPromotionPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) GetMinBeforeControlPriceOfNormal() *int32 {
	return s.MinBeforeControlPriceOfNormal
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) GetPriceShowInfo() *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValuePriceShowInfo {
	return s.PriceShowInfo
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) SetFloorPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue {
	s.FloorPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) SetTicketPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue {
	s.TicketPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) SetSellPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue {
	s.SellPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) SetOriginalSellPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue {
	s.OriginalSellPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) SetBaseTotalPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue {
	s.BaseTotalPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) SetBeforeControlPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue {
	s.BeforeControlPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) SetTax(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue {
	s.Tax = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) SetSupplyPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue {
	s.SupplyPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) SetBasicCabinPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue {
	s.BasicCabinPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) SetBuildPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue {
	s.BuildPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) SetOilPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue {
	s.OilPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) SetFirstStandardPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue {
	s.FirstStandardPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) SetBusinessStandardPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue {
	s.BusinessStandardPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) SetCommonStandardPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue {
	s.CommonStandardPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) SetInterTicketPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue {
	s.InterTicketPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) SetSubtractedPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue {
	s.SubtractedPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) SetOriginCommonPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue {
	s.OriginCommonPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) SetDynamicPromotionPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue {
	s.DynamicPromotionPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) SetInstallmentNum(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue {
	s.InstallmentNum = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) SetInstallmentPrice(v float64) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue {
	s.InstallmentPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) SetCompetitionDynamicPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue {
	s.CompetitionDynamicPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) SetCompetitionPromotionPrice(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue {
	s.CompetitionPromotionPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) SetMinBeforeControlPriceOfNormal(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue {
	s.MinBeforeControlPriceOfNormal = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) SetPriceShowInfo(v *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValuePriceShowInfo) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue {
	s.PriceShowInfo = v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValue) Validate() error {
	return dara.Validate(s)
}

type ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValuePriceShowInfo struct {
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

func (s ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValuePriceShowInfo) String() string {
	return dara.Prettify(s)
}

func (s ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValuePriceShowInfo) GoString() string {
	return s.String()
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValuePriceShowInfo) GetDiscountInfo() *string {
	return s.DiscountInfo
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValuePriceShowInfo) GetDiscountNum() *float64 {
	return s.DiscountNum
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValuePriceShowInfo) GetShowTicketPrice() *bool {
	return s.ShowTicketPrice
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValuePriceShowInfo) SetDiscountInfo(v string) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValuePriceShowInfo {
	s.DiscountInfo = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValuePriceShowInfo) SetDiscountNum(v float64) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValuePriceShowInfo {
	s.DiscountNum = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValuePriceShowInfo) SetShowTicketPrice(v bool) *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValuePriceShowInfo {
	s.ShowTicketPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueSegmentPriceValuePriceShowInfo) Validate() error {
	return dara.Validate(s)
}
