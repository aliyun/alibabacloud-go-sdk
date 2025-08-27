// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue interface {
	dara.Model
	String() string
	GoString() string
	SetFloorPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue
	GetFloorPrice() *int32
	SetTicketPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue
	GetTicketPrice() *int32
	SetSellPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue
	GetSellPrice() *int32
	SetOriginalSellPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue
	GetOriginalSellPrice() *int32
	SetBaseTotalPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue
	GetBaseTotalPrice() *int32
	SetBeforeControlPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue
	GetBeforeControlPrice() *int32
	SetTax(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue
	GetTax() *int32
	SetSupplyPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue
	GetSupplyPrice() *int32
	SetBasicCabinPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue
	GetBasicCabinPrice() *int32
	SetBuildPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue
	GetBuildPrice() *int32
	SetOilPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue
	GetOilPrice() *int32
	SetFirstStandardPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue
	GetFirstStandardPrice() *int32
	SetBusinessStandardPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue
	GetBusinessStandardPrice() *int32
	SetCommonStandardPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue
	GetCommonStandardPrice() *int32
	SetInterTicketPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue
	GetInterTicketPrice() *int32
	SetSubtractedPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue
	GetSubtractedPrice() *int32
	SetOriginCommonPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue
	GetOriginCommonPrice() *int32
	SetDynamicPromotionPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue
	GetDynamicPromotionPrice() *int32
	SetInstallmentNum(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue
	GetInstallmentNum() *int32
	SetInstallmentPrice(v float64) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue
	GetInstallmentPrice() *float64
	SetCompetitionDynamicPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue
	GetCompetitionDynamicPrice() *int32
	SetCompetitionPromotionPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue
	GetCompetitionPromotionPrice() *int32
	SetMinBeforeControlPriceOfNormal(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue
	GetMinBeforeControlPriceOfNormal() *int32
	SetPriceShowInfo(v *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue
	GetPriceShowInfo() *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo
}

type ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue struct {
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
	InterTicketPrice              *int32                                                                                       `json:"inter_ticket_price,omitempty" xml:"inter_ticket_price,omitempty"`
	SubtractedPrice               *int32                                                                                       `json:"subtracted_price,omitempty" xml:"subtracted_price,omitempty"`
	OriginCommonPrice             *int32                                                                                       `json:"origin_common_price,omitempty" xml:"origin_common_price,omitempty"`
	DynamicPromotionPrice         *int32                                                                                       `json:"dynamic_promotion_price,omitempty" xml:"dynamic_promotion_price,omitempty"`
	InstallmentNum                *int32                                                                                       `json:"installment_num,omitempty" xml:"installment_num,omitempty"`
	InstallmentPrice              *float64                                                                                     `json:"installment_price,omitempty" xml:"installment_price,omitempty"`
	CompetitionDynamicPrice       *int32                                                                                       `json:"competition_dynamic_price,omitempty" xml:"competition_dynamic_price,omitempty"`
	CompetitionPromotionPrice     *int32                                                                                       `json:"competition_promotion_price,omitempty" xml:"competition_promotion_price,omitempty"`
	MinBeforeControlPriceOfNormal *int32                                                                                       `json:"min_before_control_price_of_normal,omitempty" xml:"min_before_control_price_of_normal,omitempty"`
	PriceShowInfo                 *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo `json:"price_show_info,omitempty" xml:"price_show_info,omitempty" type:"Struct"`
}

func (s ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) GoString() string {
	return s.String()
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) GetFloorPrice() *int32 {
	return s.FloorPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) GetTicketPrice() *int32 {
	return s.TicketPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) GetSellPrice() *int32 {
	return s.SellPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) GetOriginalSellPrice() *int32 {
	return s.OriginalSellPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) GetBaseTotalPrice() *int32 {
	return s.BaseTotalPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) GetBeforeControlPrice() *int32 {
	return s.BeforeControlPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) GetTax() *int32 {
	return s.Tax
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) GetSupplyPrice() *int32 {
	return s.SupplyPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) GetBasicCabinPrice() *int32 {
	return s.BasicCabinPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) GetBuildPrice() *int32 {
	return s.BuildPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) GetOilPrice() *int32 {
	return s.OilPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) GetFirstStandardPrice() *int32 {
	return s.FirstStandardPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) GetBusinessStandardPrice() *int32 {
	return s.BusinessStandardPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) GetCommonStandardPrice() *int32 {
	return s.CommonStandardPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) GetInterTicketPrice() *int32 {
	return s.InterTicketPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) GetSubtractedPrice() *int32 {
	return s.SubtractedPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) GetOriginCommonPrice() *int32 {
	return s.OriginCommonPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) GetDynamicPromotionPrice() *int32 {
	return s.DynamicPromotionPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) GetInstallmentNum() *int32 {
	return s.InstallmentNum
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) GetInstallmentPrice() *float64 {
	return s.InstallmentPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) GetCompetitionDynamicPrice() *int32 {
	return s.CompetitionDynamicPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) GetCompetitionPromotionPrice() *int32 {
	return s.CompetitionPromotionPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) GetMinBeforeControlPriceOfNormal() *int32 {
	return s.MinBeforeControlPriceOfNormal
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) GetPriceShowInfo() *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo {
	return s.PriceShowInfo
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) SetFloorPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue {
	s.FloorPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) SetTicketPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue {
	s.TicketPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) SetSellPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue {
	s.SellPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) SetOriginalSellPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue {
	s.OriginalSellPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) SetBaseTotalPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue {
	s.BaseTotalPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) SetBeforeControlPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue {
	s.BeforeControlPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) SetTax(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue {
	s.Tax = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) SetSupplyPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue {
	s.SupplyPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) SetBasicCabinPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue {
	s.BasicCabinPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) SetBuildPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue {
	s.BuildPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) SetOilPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue {
	s.OilPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) SetFirstStandardPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue {
	s.FirstStandardPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) SetBusinessStandardPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue {
	s.BusinessStandardPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) SetCommonStandardPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue {
	s.CommonStandardPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) SetInterTicketPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue {
	s.InterTicketPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) SetSubtractedPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue {
	s.SubtractedPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) SetOriginCommonPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue {
	s.OriginCommonPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) SetDynamicPromotionPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue {
	s.DynamicPromotionPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) SetInstallmentNum(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue {
	s.InstallmentNum = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) SetInstallmentPrice(v float64) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue {
	s.InstallmentPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) SetCompetitionDynamicPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue {
	s.CompetitionDynamicPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) SetCompetitionPromotionPrice(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue {
	s.CompetitionPromotionPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) SetMinBeforeControlPriceOfNormal(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue {
	s.MinBeforeControlPriceOfNormal = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) SetPriceShowInfo(v *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue {
	s.PriceShowInfo = v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValue) Validate() error {
	return dara.Validate(s)
}

type ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo struct {
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

func (s ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo) String() string {
	return dara.Prettify(s)
}

func (s ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo) GoString() string {
	return s.String()
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo) GetDiscountInfo() *string {
	return s.DiscountInfo
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo) GetDiscountNum() *float64 {
	return s.DiscountNum
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo) GetShowTicketPrice() *bool {
	return s.ShowTicketPrice
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo) SetDiscountInfo(v string) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo {
	s.DiscountInfo = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo) SetDiscountNum(v float64) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo {
	s.DiscountNum = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo) SetShowTicketPrice(v bool) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo {
	s.ShowTicketPrice = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueSegmentPriceValuePriceShowInfo) Validate() error {
	return dara.Validate(s)
}
