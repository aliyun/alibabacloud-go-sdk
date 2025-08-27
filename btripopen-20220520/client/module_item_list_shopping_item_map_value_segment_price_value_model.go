// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleItemListShoppingItemMapValueSegmentPriceValue interface {
	dara.Model
	String() string
	GoString() string
	SetFloorPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue
	GetFloorPrice() *int32
	SetTicketPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue
	GetTicketPrice() *int32
	SetSellPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue
	GetSellPrice() *int32
	SetOriginalSellPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue
	GetOriginalSellPrice() *int32
	SetBaseTotalPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue
	GetBaseTotalPrice() *int32
	SetBeforeControlPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue
	GetBeforeControlPrice() *int32
	SetTax(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue
	GetTax() *int32
	SetSupplyPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue
	GetSupplyPrice() *int32
	SetBasicCabinPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue
	GetBasicCabinPrice() *int32
	SetBuildPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue
	GetBuildPrice() *int32
	SetOilPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue
	GetOilPrice() *int32
	SetFirstStandardPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue
	GetFirstStandardPrice() *int32
	SetBusinessStandardPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue
	GetBusinessStandardPrice() *int32
	SetCommonStandardPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue
	GetCommonStandardPrice() *int32
	SetInterTicketPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue
	GetInterTicketPrice() *int32
	SetSubtractedPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue
	GetSubtractedPrice() *int32
	SetOriginCommonPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue
	GetOriginCommonPrice() *int32
	SetDynamicPromotionPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue
	GetDynamicPromotionPrice() *int32
	SetInstallmentNum(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue
	GetInstallmentNum() *int32
	SetInstallmentPrice(v float64) *ModuleItemListShoppingItemMapValueSegmentPriceValue
	GetInstallmentPrice() *float64
	SetCompetitionDynamicPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue
	GetCompetitionDynamicPrice() *int32
	SetCompetitionPromotionPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue
	GetCompetitionPromotionPrice() *int32
	SetMinBeforeControlPriceOfNormal(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue
	GetMinBeforeControlPriceOfNormal() *int32
	SetPriceShowInfo(v *ModuleItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo) *ModuleItemListShoppingItemMapValueSegmentPriceValue
	GetPriceShowInfo() *ModuleItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo
}

type ModuleItemListShoppingItemMapValueSegmentPriceValue struct {
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
	InterTicketPrice              *int32                                                            `json:"inter_ticket_price,omitempty" xml:"inter_ticket_price,omitempty"`
	SubtractedPrice               *int32                                                            `json:"subtracted_price,omitempty" xml:"subtracted_price,omitempty"`
	OriginCommonPrice             *int32                                                            `json:"origin_common_price,omitempty" xml:"origin_common_price,omitempty"`
	DynamicPromotionPrice         *int32                                                            `json:"dynamic_promotion_price,omitempty" xml:"dynamic_promotion_price,omitempty"`
	InstallmentNum                *int32                                                            `json:"installment_num,omitempty" xml:"installment_num,omitempty"`
	InstallmentPrice              *float64                                                          `json:"installment_price,omitempty" xml:"installment_price,omitempty"`
	CompetitionDynamicPrice       *int32                                                            `json:"competition_dynamic_price,omitempty" xml:"competition_dynamic_price,omitempty"`
	CompetitionPromotionPrice     *int32                                                            `json:"competition_promotion_price,omitempty" xml:"competition_promotion_price,omitempty"`
	MinBeforeControlPriceOfNormal *int32                                                            `json:"min_before_control_price_of_normal,omitempty" xml:"min_before_control_price_of_normal,omitempty"`
	PriceShowInfo                 *ModuleItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo `json:"price_show_info,omitempty" xml:"price_show_info,omitempty" type:"Struct"`
}

func (s ModuleItemListShoppingItemMapValueSegmentPriceValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleItemListShoppingItemMapValueSegmentPriceValue) GoString() string {
	return s.String()
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) GetFloorPrice() *int32 {
	return s.FloorPrice
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) GetTicketPrice() *int32 {
	return s.TicketPrice
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) GetSellPrice() *int32 {
	return s.SellPrice
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) GetOriginalSellPrice() *int32 {
	return s.OriginalSellPrice
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) GetBaseTotalPrice() *int32 {
	return s.BaseTotalPrice
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) GetBeforeControlPrice() *int32 {
	return s.BeforeControlPrice
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) GetTax() *int32 {
	return s.Tax
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) GetSupplyPrice() *int32 {
	return s.SupplyPrice
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) GetBasicCabinPrice() *int32 {
	return s.BasicCabinPrice
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) GetBuildPrice() *int32 {
	return s.BuildPrice
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) GetOilPrice() *int32 {
	return s.OilPrice
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) GetFirstStandardPrice() *int32 {
	return s.FirstStandardPrice
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) GetBusinessStandardPrice() *int32 {
	return s.BusinessStandardPrice
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) GetCommonStandardPrice() *int32 {
	return s.CommonStandardPrice
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) GetInterTicketPrice() *int32 {
	return s.InterTicketPrice
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) GetSubtractedPrice() *int32 {
	return s.SubtractedPrice
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) GetOriginCommonPrice() *int32 {
	return s.OriginCommonPrice
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) GetDynamicPromotionPrice() *int32 {
	return s.DynamicPromotionPrice
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) GetInstallmentNum() *int32 {
	return s.InstallmentNum
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) GetInstallmentPrice() *float64 {
	return s.InstallmentPrice
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) GetCompetitionDynamicPrice() *int32 {
	return s.CompetitionDynamicPrice
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) GetCompetitionPromotionPrice() *int32 {
	return s.CompetitionPromotionPrice
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) GetMinBeforeControlPriceOfNormal() *int32 {
	return s.MinBeforeControlPriceOfNormal
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) GetPriceShowInfo() *ModuleItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo {
	return s.PriceShowInfo
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) SetFloorPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue {
	s.FloorPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) SetTicketPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue {
	s.TicketPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) SetSellPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue {
	s.SellPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) SetOriginalSellPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue {
	s.OriginalSellPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) SetBaseTotalPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue {
	s.BaseTotalPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) SetBeforeControlPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue {
	s.BeforeControlPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) SetTax(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue {
	s.Tax = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) SetSupplyPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue {
	s.SupplyPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) SetBasicCabinPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue {
	s.BasicCabinPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) SetBuildPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue {
	s.BuildPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) SetOilPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue {
	s.OilPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) SetFirstStandardPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue {
	s.FirstStandardPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) SetBusinessStandardPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue {
	s.BusinessStandardPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) SetCommonStandardPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue {
	s.CommonStandardPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) SetInterTicketPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue {
	s.InterTicketPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) SetSubtractedPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue {
	s.SubtractedPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) SetOriginCommonPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue {
	s.OriginCommonPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) SetDynamicPromotionPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue {
	s.DynamicPromotionPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) SetInstallmentNum(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue {
	s.InstallmentNum = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) SetInstallmentPrice(v float64) *ModuleItemListShoppingItemMapValueSegmentPriceValue {
	s.InstallmentPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) SetCompetitionDynamicPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue {
	s.CompetitionDynamicPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) SetCompetitionPromotionPrice(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue {
	s.CompetitionPromotionPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) SetMinBeforeControlPriceOfNormal(v int32) *ModuleItemListShoppingItemMapValueSegmentPriceValue {
	s.MinBeforeControlPriceOfNormal = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) SetPriceShowInfo(v *ModuleItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo) *ModuleItemListShoppingItemMapValueSegmentPriceValue {
	s.PriceShowInfo = v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValue) Validate() error {
	return dara.Validate(s)
}

type ModuleItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo struct {
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

func (s ModuleItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo) String() string {
	return dara.Prettify(s)
}

func (s ModuleItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo) GoString() string {
	return s.String()
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo) GetDiscountInfo() *string {
	return s.DiscountInfo
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo) GetDiscountNum() *float64 {
	return s.DiscountNum
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo) GetShowTicketPrice() *bool {
	return s.ShowTicketPrice
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo) SetDiscountInfo(v string) *ModuleItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo {
	s.DiscountInfo = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo) SetDiscountNum(v float64) *ModuleItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo {
	s.DiscountNum = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo) SetShowTicketPrice(v bool) *ModuleItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo {
	s.ShowTicketPrice = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueSegmentPriceValuePriceShowInfo) Validate() error {
	return dara.Validate(s)
}
