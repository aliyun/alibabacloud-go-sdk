// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetStartCityCode(v string) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetStartCityCode() *string
	SetEndCityCode(v string) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetEndCityCode() *string
	SetCarryFreePc(v int32) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetCarryFreePc() *int32
	SetCarryBagWeight(v int32) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetCarryBagWeight() *int32
	SetCarryBagSize(v string) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetCarryBagSize() *string
	SetIsAllCarryBagWeight(v bool) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetIsAllCarryBagWeight() *bool
	SetTotalPcs(v int64) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetTotalPcs() *int64
	SetTotalWeight(v int64) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetTotalWeight() *int64
	SetCarryUnknown(v bool) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetCarryUnknown() *bool
	SetCarryLength(v int32) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetCarryLength() *int32
	SetCarryWidth(v int32) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetCarryWidth() *int32
	SetCarryHeight(v int32) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetCarryHeight() *int32
	SetCarrySumOfLengthWidthHeight(v int32) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetCarrySumOfLengthWidthHeight() *int32
	SetFreePcs(v int64) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetFreePcs() *int64
	SetBaggageWeight(v int64) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetBaggageWeight() *int64
	SetBaggageUnit(v string) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetBaggageUnit() *string
	SetBaggageSize(v string) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetBaggageSize() *string
	SetAllWeight(v bool) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetAllWeight() *bool
	SetLength(v int32) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetLength() *int32
	SetWidth(v int32) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetWidth() *int32
	SetHeight(v int32) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetHeight() *int32
	SetSumOfLengthWidthHeight(v int32) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetSumOfLengthWidthHeight() *int32
	SetUnknown(v bool) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetUnknown() *bool
	SetCnDesc(v string) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetCnDesc() *string
	SetEnDesc(v string) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetEnDesc() *string
	SetAttribute(v string) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetAttribute() *string
	SetBaggagePrice(v int32) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetBaggagePrice() *int32
	SetCarryOnBaggageTips(v string) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetCarryOnBaggageTips() *string
}

type ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue struct {
	// example:
	//
	// NNG
	StartCityCode *string `json:"start_city_code,omitempty" xml:"start_city_code,omitempty"`
	// example:
	//
	// SHA
	EndCityCode *string `json:"end_city_code,omitempty" xml:"end_city_code,omitempty"`
	// example:
	//
	// 1
	CarryFreePc *int32 `json:"carry_free_pc,omitempty" xml:"carry_free_pc,omitempty"`
	// example:
	//
	// 20
	CarryBagWeight *int32 `json:"carry_bag_weight,omitempty" xml:"carry_bag_weight,omitempty"`
	// example:
	//
	// 20*40*55CM、三边之和不超过115CM
	CarryBagSize *string `json:"carry_bag_size,omitempty" xml:"carry_bag_size,omitempty"`
	// example:
	//
	// true
	IsAllCarryBagWeight *bool `json:"is_all_carry_bag_weight,omitempty" xml:"is_all_carry_bag_weight,omitempty"`
	// example:
	//
	// 1
	TotalPcs *int64 `json:"total_pcs,omitempty" xml:"total_pcs,omitempty"`
	// example:
	//
	// 20
	TotalWeight *int64 `json:"total_weight,omitempty" xml:"total_weight,omitempty"`
	// example:
	//
	// true
	CarryUnknown *bool `json:"carry_unknown,omitempty" xml:"carry_unknown,omitempty"`
	// example:
	//
	// 20
	CarryLength *int32 `json:"carry_length,omitempty" xml:"carry_length,omitempty"`
	// example:
	//
	// 30
	CarryWidth *int32 `json:"carry_width,omitempty" xml:"carry_width,omitempty"`
	// example:
	//
	// 55
	CarryHeight *int32 `json:"carry_height,omitempty" xml:"carry_height,omitempty"`
	// example:
	//
	// 115
	CarrySumOfLengthWidthHeight *int32 `json:"carry_sum_of_length_width_height,omitempty" xml:"carry_sum_of_length_width_height,omitempty"`
	// example:
	//
	// 1
	FreePcs *int64 `json:"free_pcs,omitempty" xml:"free_pcs,omitempty"`
	// example:
	//
	// 20
	BaggageWeight *int64 `json:"baggage_weight,omitempty" xml:"baggage_weight,omitempty"`
	// example:
	//
	// KG
	BaggageUnit *string `json:"baggage_unit,omitempty" xml:"baggage_unit,omitempty"`
	// example:
	//
	// 长宽高之和≤158CM
	BaggageSize *string `json:"baggage_size,omitempty" xml:"baggage_size,omitempty"`
	// example:
	//
	// true
	AllWeight *bool `json:"all_weight,omitempty" xml:"all_weight,omitempty"`
	// example:
	//
	// 20
	Length *int32 `json:"length,omitempty" xml:"length,omitempty"`
	// example:
	//
	// 30
	Width *int32 `json:"width,omitempty" xml:"width,omitempty"`
	// example:
	//
	// 55
	Height *int32 `json:"height,omitempty" xml:"height,omitempty"`
	// example:
	//
	// 115
	SumOfLengthWidthHeight *int32 `json:"sum_of_length_width_height,omitempty" xml:"sum_of_length_width_height,omitempty"`
	// example:
	//
	// false
	Unknown *bool `json:"unknown,omitempty" xml:"unknown,omitempty"`
	// example:
	//
	// -
	CnDesc *string `json:"cn_desc,omitempty" xml:"cn_desc,omitempty"`
	// example:
	//
	// -
	EnDesc *string `json:"en_desc,omitempty" xml:"en_desc,omitempty"`
	// example:
	//
	// {}
	Attribute *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// example:
	//
	// 20
	BaggagePrice *int32 `json:"baggage_price,omitempty" xml:"baggage_price,omitempty"`
	// example:
	//
	// -
	CarryOnBaggageTips *string `json:"carry_on_baggage_tips,omitempty" xml:"carry_on_baggage_tips,omitempty"`
}

func (s ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GoString() string {
	return s.String()
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetStartCityCode() *string {
	return s.StartCityCode
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetEndCityCode() *string {
	return s.EndCityCode
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetCarryFreePc() *int32 {
	return s.CarryFreePc
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetCarryBagWeight() *int32 {
	return s.CarryBagWeight
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetCarryBagSize() *string {
	return s.CarryBagSize
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetIsAllCarryBagWeight() *bool {
	return s.IsAllCarryBagWeight
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetTotalPcs() *int64 {
	return s.TotalPcs
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetTotalWeight() *int64 {
	return s.TotalWeight
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetCarryUnknown() *bool {
	return s.CarryUnknown
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetCarryLength() *int32 {
	return s.CarryLength
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetCarryWidth() *int32 {
	return s.CarryWidth
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetCarryHeight() *int32 {
	return s.CarryHeight
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetCarrySumOfLengthWidthHeight() *int32 {
	return s.CarrySumOfLengthWidthHeight
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetFreePcs() *int64 {
	return s.FreePcs
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetBaggageWeight() *int64 {
	return s.BaggageWeight
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetBaggageUnit() *string {
	return s.BaggageUnit
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetBaggageSize() *string {
	return s.BaggageSize
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetAllWeight() *bool {
	return s.AllWeight
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetLength() *int32 {
	return s.Length
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetWidth() *int32 {
	return s.Width
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetHeight() *int32 {
	return s.Height
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetSumOfLengthWidthHeight() *int32 {
	return s.SumOfLengthWidthHeight
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetUnknown() *bool {
	return s.Unknown
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetCnDesc() *string {
	return s.CnDesc
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetEnDesc() *string {
	return s.EnDesc
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetAttribute() *string {
	return s.Attribute
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetBaggagePrice() *int32 {
	return s.BaggagePrice
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetCarryOnBaggageTips() *string {
	return s.CarryOnBaggageTips
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetStartCityCode(v string) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.StartCityCode = &v
	return s
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetEndCityCode(v string) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.EndCityCode = &v
	return s
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetCarryFreePc(v int32) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.CarryFreePc = &v
	return s
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetCarryBagWeight(v int32) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.CarryBagWeight = &v
	return s
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetCarryBagSize(v string) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.CarryBagSize = &v
	return s
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetIsAllCarryBagWeight(v bool) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.IsAllCarryBagWeight = &v
	return s
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetTotalPcs(v int64) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.TotalPcs = &v
	return s
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetTotalWeight(v int64) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.TotalWeight = &v
	return s
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetCarryUnknown(v bool) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.CarryUnknown = &v
	return s
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetCarryLength(v int32) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.CarryLength = &v
	return s
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetCarryWidth(v int32) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.CarryWidth = &v
	return s
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetCarryHeight(v int32) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.CarryHeight = &v
	return s
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetCarrySumOfLengthWidthHeight(v int32) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.CarrySumOfLengthWidthHeight = &v
	return s
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetFreePcs(v int64) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.FreePcs = &v
	return s
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetBaggageWeight(v int64) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.BaggageWeight = &v
	return s
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetBaggageUnit(v string) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.BaggageUnit = &v
	return s
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetBaggageSize(v string) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.BaggageSize = &v
	return s
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetAllWeight(v bool) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.AllWeight = &v
	return s
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetLength(v int32) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.Length = &v
	return s
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetWidth(v int32) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.Width = &v
	return s
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetHeight(v int32) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.Height = &v
	return s
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetSumOfLengthWidthHeight(v int32) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.SumOfLengthWidthHeight = &v
	return s
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetUnknown(v bool) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.Unknown = &v
	return s
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetCnDesc(v string) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.CnDesc = &v
	return s
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetEnDesc(v string) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.EnDesc = &v
	return s
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetAttribute(v string) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.Attribute = &v
	return s
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetBaggagePrice(v int32) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.BaggagePrice = &v
	return s
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetCarryOnBaggageTips(v string) *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.CarryOnBaggageTips = &v
	return s
}

func (s *ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) Validate() error {
	return dara.Validate(s)
}
