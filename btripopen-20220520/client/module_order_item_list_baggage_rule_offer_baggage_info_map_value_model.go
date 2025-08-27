// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetStartCityCode(v string) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue
	GetStartCityCode() *string
	SetEndCityCode(v string) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue
	GetEndCityCode() *string
	SetCarryFreePc(v int32) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue
	GetCarryFreePc() *int32
	SetCarryBagWeight(v int32) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue
	GetCarryBagWeight() *int32
	SetCarryBagSize(v string) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue
	GetCarryBagSize() *string
	SetIsAllCarryBagWeight(v bool) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue
	GetIsAllCarryBagWeight() *bool
	SetTotalPcs(v int64) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue
	GetTotalPcs() *int64
	SetTotalWeight(v int64) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue
	GetTotalWeight() *int64
	SetCarryUnknown(v bool) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue
	GetCarryUnknown() *bool
	SetCarryLength(v int32) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue
	GetCarryLength() *int32
	SetCarryWidth(v int32) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue
	GetCarryWidth() *int32
	SetCarryHeight(v int32) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue
	GetCarryHeight() *int32
	SetCarrySumOfLengthWidthHeight(v int32) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue
	GetCarrySumOfLengthWidthHeight() *int32
	SetFreePcs(v int64) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue
	GetFreePcs() *int64
	SetBaggageWeight(v int64) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue
	GetBaggageWeight() *int64
	SetBaggageUnit(v string) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue
	GetBaggageUnit() *string
	SetBaggageSize(v string) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue
	GetBaggageSize() *string
	SetAllWeight(v bool) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue
	GetAllWeight() *bool
	SetLength(v int32) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue
	GetLength() *int32
	SetWidth(v int32) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue
	GetWidth() *int32
	SetHeight(v int32) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue
	GetHeight() *int32
	SetSumOfLengthWidthHeight(v int32) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue
	GetSumOfLengthWidthHeight() *int32
	SetUnknown(v bool) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue
	GetUnknown() *bool
	SetCnDesc(v string) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue
	GetCnDesc() *string
	SetEnDesc(v string) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue
	GetEnDesc() *string
	SetAttribute(v string) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue
	GetAttribute() *string
	SetBaggagePrice(v int32) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue
	GetBaggagePrice() *int32
	SetCarryOnBaggageTips(v string) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue
	GetCarryOnBaggageTips() *string
}

type ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue struct {
	StartCityCode               *string `json:"start_city_code,omitempty" xml:"start_city_code,omitempty"`
	EndCityCode                 *string `json:"end_city_code,omitempty" xml:"end_city_code,omitempty"`
	CarryFreePc                 *int32  `json:"carry_free_pc,omitempty" xml:"carry_free_pc,omitempty"`
	CarryBagWeight              *int32  `json:"carry_bag_weight,omitempty" xml:"carry_bag_weight,omitempty"`
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

func (s ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) GoString() string {
	return s.String()
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) GetStartCityCode() *string {
	return s.StartCityCode
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) GetEndCityCode() *string {
	return s.EndCityCode
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) GetCarryFreePc() *int32 {
	return s.CarryFreePc
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) GetCarryBagWeight() *int32 {
	return s.CarryBagWeight
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) GetCarryBagSize() *string {
	return s.CarryBagSize
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) GetIsAllCarryBagWeight() *bool {
	return s.IsAllCarryBagWeight
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) GetTotalPcs() *int64 {
	return s.TotalPcs
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) GetTotalWeight() *int64 {
	return s.TotalWeight
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) GetCarryUnknown() *bool {
	return s.CarryUnknown
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) GetCarryLength() *int32 {
	return s.CarryLength
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) GetCarryWidth() *int32 {
	return s.CarryWidth
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) GetCarryHeight() *int32 {
	return s.CarryHeight
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) GetCarrySumOfLengthWidthHeight() *int32 {
	return s.CarrySumOfLengthWidthHeight
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) GetFreePcs() *int64 {
	return s.FreePcs
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) GetBaggageWeight() *int64 {
	return s.BaggageWeight
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) GetBaggageUnit() *string {
	return s.BaggageUnit
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) GetBaggageSize() *string {
	return s.BaggageSize
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) GetAllWeight() *bool {
	return s.AllWeight
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) GetLength() *int32 {
	return s.Length
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) GetWidth() *int32 {
	return s.Width
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) GetHeight() *int32 {
	return s.Height
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) GetSumOfLengthWidthHeight() *int32 {
	return s.SumOfLengthWidthHeight
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) GetUnknown() *bool {
	return s.Unknown
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) GetCnDesc() *string {
	return s.CnDesc
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) GetEnDesc() *string {
	return s.EnDesc
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) GetAttribute() *string {
	return s.Attribute
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) GetBaggagePrice() *int32 {
	return s.BaggagePrice
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) GetCarryOnBaggageTips() *string {
	return s.CarryOnBaggageTips
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) SetStartCityCode(v string) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue {
	s.StartCityCode = &v
	return s
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) SetEndCityCode(v string) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue {
	s.EndCityCode = &v
	return s
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) SetCarryFreePc(v int32) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue {
	s.CarryFreePc = &v
	return s
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) SetCarryBagWeight(v int32) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue {
	s.CarryBagWeight = &v
	return s
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) SetCarryBagSize(v string) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue {
	s.CarryBagSize = &v
	return s
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) SetIsAllCarryBagWeight(v bool) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue {
	s.IsAllCarryBagWeight = &v
	return s
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) SetTotalPcs(v int64) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue {
	s.TotalPcs = &v
	return s
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) SetTotalWeight(v int64) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue {
	s.TotalWeight = &v
	return s
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) SetCarryUnknown(v bool) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue {
	s.CarryUnknown = &v
	return s
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) SetCarryLength(v int32) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue {
	s.CarryLength = &v
	return s
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) SetCarryWidth(v int32) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue {
	s.CarryWidth = &v
	return s
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) SetCarryHeight(v int32) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue {
	s.CarryHeight = &v
	return s
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) SetCarrySumOfLengthWidthHeight(v int32) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue {
	s.CarrySumOfLengthWidthHeight = &v
	return s
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) SetFreePcs(v int64) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue {
	s.FreePcs = &v
	return s
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) SetBaggageWeight(v int64) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue {
	s.BaggageWeight = &v
	return s
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) SetBaggageUnit(v string) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue {
	s.BaggageUnit = &v
	return s
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) SetBaggageSize(v string) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue {
	s.BaggageSize = &v
	return s
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) SetAllWeight(v bool) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue {
	s.AllWeight = &v
	return s
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) SetLength(v int32) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue {
	s.Length = &v
	return s
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) SetWidth(v int32) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue {
	s.Width = &v
	return s
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) SetHeight(v int32) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue {
	s.Height = &v
	return s
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) SetSumOfLengthWidthHeight(v int32) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue {
	s.SumOfLengthWidthHeight = &v
	return s
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) SetUnknown(v bool) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue {
	s.Unknown = &v
	return s
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) SetCnDesc(v string) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue {
	s.CnDesc = &v
	return s
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) SetEnDesc(v string) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue {
	s.EnDesc = &v
	return s
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) SetAttribute(v string) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue {
	s.Attribute = &v
	return s
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) SetBaggagePrice(v int32) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue {
	s.BaggagePrice = &v
	return s
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) SetCarryOnBaggageTips(v string) *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue {
	s.CarryOnBaggageTips = &v
	return s
}

func (s *ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) Validate() error {
	return dara.Validate(s)
}
