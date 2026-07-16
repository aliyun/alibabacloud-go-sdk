// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetStartCityCode(v string) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetStartCityCode() *string
	SetEndCityCode(v string) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetEndCityCode() *string
	SetCarryFreePc(v int32) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetCarryFreePc() *int32
	SetCarryBagWeight(v int32) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetCarryBagWeight() *int32
	SetCarryOnWeightUnit(v string) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetCarryOnWeightUnit() *string
	SetCarryBagSize(v string) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetCarryBagSize() *string
	SetIsAllCarryBagWeight(v bool) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetIsAllCarryBagWeight() *bool
	SetTotalPcs(v int64) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetTotalPcs() *int64
	SetTotalWeight(v int64) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetTotalWeight() *int64
	SetCarryUnknown(v bool) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetCarryUnknown() *bool
	SetCarryLength(v int32) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetCarryLength() *int32
	SetCarryWidth(v int32) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetCarryWidth() *int32
	SetCarryHeight(v int32) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetCarryHeight() *int32
	SetCarrySumOfLengthWidthHeight(v int32) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetCarrySumOfLengthWidthHeight() *int32
	SetFreePcs(v int64) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetFreePcs() *int64
	SetBaggageWeight(v int64) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetBaggageWeight() *int64
	SetBaggageUnit(v string) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetBaggageUnit() *string
	SetBaggageSize(v string) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetBaggageSize() *string
	SetAllWeight(v bool) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetAllWeight() *bool
	SetLength(v int32) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetLength() *int32
	SetWidth(v int32) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetWidth() *int32
	SetHeight(v int32) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetHeight() *int32
	SetSumOfLengthWidthHeight(v int32) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetSumOfLengthWidthHeight() *int32
	SetUnknown(v bool) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetUnknown() *bool
	SetCnDesc(v string) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetCnDesc() *string
	SetEnDesc(v string) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetEnDesc() *string
	SetAttribute(v string) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetAttribute() *string
	SetBaggagePrice(v int32) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetBaggagePrice() *int32
	SetCarryOnBaggageTips(v string) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue
	GetCarryOnBaggageTips() *string
}

type ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue struct {
	StartCityCode               *string `json:"start_city_code,omitempty" xml:"start_city_code,omitempty"`
	EndCityCode                 *string `json:"end_city_code,omitempty" xml:"end_city_code,omitempty"`
	CarryFreePc                 *int32  `json:"carry_free_pc,omitempty" xml:"carry_free_pc,omitempty"`
	CarryBagWeight              *int32  `json:"carry_bag_weight,omitempty" xml:"carry_bag_weight,omitempty"`
	CarryOnWeightUnit           *string `json:"carry_on_weight_unit,omitempty" xml:"carry_on_weight_unit,omitempty"`
	CarryBagSize                *string `json:"carry_bag_size,omitempty" xml:"carry_bag_size,omitempty"`
	IsAllCarryBagWeight         *bool   `json:"is_all_carry_bag_weight,omitempty" xml:"is_all_carry_bag_weight,omitempty"`
	TotalPcs                    *int64  `json:"total_pcs,omitempty" xml:"total_pcs,omitempty"`
	TotalWeight                 *int64  `json:"total_weight,omitempty" xml:"total_weight,omitempty"`
	CarryUnknown                *bool   `json:"carry_unknown,omitempty" xml:"carry_unknown,omitempty"`
	CarryLength                 *int32  `json:"carry_length,omitempty" xml:"carry_length,omitempty"`
	CarryWidth                  *int32  `json:"carry_width,omitempty" xml:"carry_width,omitempty"`
	CarryHeight                 *int32  `json:"carry_height,omitempty" xml:"carry_height,omitempty"`
	CarrySumOfLengthWidthHeight *int32  `json:"carry_sum_of_length_width_height,omitempty" xml:"carry_sum_of_length_width_height,omitempty"`
	FreePcs                     *int64  `json:"free_pcs,omitempty" xml:"free_pcs,omitempty"`
	BaggageWeight               *int64  `json:"baggage_weight,omitempty" xml:"baggage_weight,omitempty"`
	BaggageUnit                 *string `json:"baggage_unit,omitempty" xml:"baggage_unit,omitempty"`
	BaggageSize                 *string `json:"baggage_size,omitempty" xml:"baggage_size,omitempty"`
	AllWeight                   *bool   `json:"all_weight,omitempty" xml:"all_weight,omitempty"`
	Length                      *int32  `json:"length,omitempty" xml:"length,omitempty"`
	Width                       *int32  `json:"width,omitempty" xml:"width,omitempty"`
	Height                      *int32  `json:"height,omitempty" xml:"height,omitempty"`
	SumOfLengthWidthHeight      *int32  `json:"sum_of_length_width_height,omitempty" xml:"sum_of_length_width_height,omitempty"`
	Unknown                     *bool   `json:"unknown,omitempty" xml:"unknown,omitempty"`
	CnDesc                      *string `json:"cn_desc,omitempty" xml:"cn_desc,omitempty"`
	EnDesc                      *string `json:"en_desc,omitempty" xml:"en_desc,omitempty"`
	Attribute                   *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	BaggagePrice                *int32  `json:"baggage_price,omitempty" xml:"baggage_price,omitempty"`
	CarryOnBaggageTips          *string `json:"carry_on_baggage_tips,omitempty" xml:"carry_on_baggage_tips,omitempty"`
}

func (s ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GoString() string {
	return s.String()
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetStartCityCode() *string {
	return s.StartCityCode
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetEndCityCode() *string {
	return s.EndCityCode
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetCarryFreePc() *int32 {
	return s.CarryFreePc
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetCarryBagWeight() *int32 {
	return s.CarryBagWeight
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetCarryOnWeightUnit() *string {
	return s.CarryOnWeightUnit
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetCarryBagSize() *string {
	return s.CarryBagSize
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetIsAllCarryBagWeight() *bool {
	return s.IsAllCarryBagWeight
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetTotalPcs() *int64 {
	return s.TotalPcs
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetTotalWeight() *int64 {
	return s.TotalWeight
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetCarryUnknown() *bool {
	return s.CarryUnknown
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetCarryLength() *int32 {
	return s.CarryLength
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetCarryWidth() *int32 {
	return s.CarryWidth
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetCarryHeight() *int32 {
	return s.CarryHeight
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetCarrySumOfLengthWidthHeight() *int32 {
	return s.CarrySumOfLengthWidthHeight
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetFreePcs() *int64 {
	return s.FreePcs
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetBaggageWeight() *int64 {
	return s.BaggageWeight
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetBaggageUnit() *string {
	return s.BaggageUnit
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetBaggageSize() *string {
	return s.BaggageSize
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetAllWeight() *bool {
	return s.AllWeight
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetLength() *int32 {
	return s.Length
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetWidth() *int32 {
	return s.Width
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetHeight() *int32 {
	return s.Height
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetSumOfLengthWidthHeight() *int32 {
	return s.SumOfLengthWidthHeight
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetUnknown() *bool {
	return s.Unknown
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetCnDesc() *string {
	return s.CnDesc
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetEnDesc() *string {
	return s.EnDesc
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetAttribute() *string {
	return s.Attribute
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetBaggagePrice() *int32 {
	return s.BaggagePrice
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) GetCarryOnBaggageTips() *string {
	return s.CarryOnBaggageTips
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetStartCityCode(v string) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.StartCityCode = &v
	return s
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetEndCityCode(v string) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.EndCityCode = &v
	return s
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetCarryFreePc(v int32) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.CarryFreePc = &v
	return s
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetCarryBagWeight(v int32) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.CarryBagWeight = &v
	return s
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetCarryOnWeightUnit(v string) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.CarryOnWeightUnit = &v
	return s
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetCarryBagSize(v string) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.CarryBagSize = &v
	return s
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetIsAllCarryBagWeight(v bool) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.IsAllCarryBagWeight = &v
	return s
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetTotalPcs(v int64) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.TotalPcs = &v
	return s
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetTotalWeight(v int64) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.TotalWeight = &v
	return s
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetCarryUnknown(v bool) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.CarryUnknown = &v
	return s
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetCarryLength(v int32) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.CarryLength = &v
	return s
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetCarryWidth(v int32) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.CarryWidth = &v
	return s
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetCarryHeight(v int32) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.CarryHeight = &v
	return s
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetCarrySumOfLengthWidthHeight(v int32) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.CarrySumOfLengthWidthHeight = &v
	return s
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetFreePcs(v int64) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.FreePcs = &v
	return s
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetBaggageWeight(v int64) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.BaggageWeight = &v
	return s
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetBaggageUnit(v string) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.BaggageUnit = &v
	return s
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetBaggageSize(v string) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.BaggageSize = &v
	return s
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetAllWeight(v bool) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.AllWeight = &v
	return s
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetLength(v int32) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.Length = &v
	return s
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetWidth(v int32) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.Width = &v
	return s
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetHeight(v int32) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.Height = &v
	return s
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetSumOfLengthWidthHeight(v int32) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.SumOfLengthWidthHeight = &v
	return s
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetUnknown(v bool) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.Unknown = &v
	return s
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetCnDesc(v string) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.CnDesc = &v
	return s
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetEnDesc(v string) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.EnDesc = &v
	return s
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetAttribute(v string) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.Attribute = &v
	return s
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetBaggagePrice(v int32) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.BaggagePrice = &v
	return s
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) SetCarryOnBaggageTips(v string) *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	s.CarryOnBaggageTips = &v
	return s
}

func (s *ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) Validate() error {
	return dara.Validate(s)
}
