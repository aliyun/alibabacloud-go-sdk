// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetCarryFreepc(v int32) *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue
	GetCarryFreepc() *int32
	SetCarryBagWeight(v int32) *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue
	GetCarryBagWeight() *int32
	SetCarryBagSize(v string) *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue
	GetCarryBagSize() *string
	SetIsAllCarryBagWeight(v bool) *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue
	GetIsAllCarryBagWeight() *bool
	SetAirline(v string) *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue
	GetAirline() *string
	SetStartAirport(v string) *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue
	GetStartAirport() *string
	SetEndAirport(v string) *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue
	GetEndAirport() *string
	SetStartCityCode(v string) *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue
	GetStartCityCode() *string
	SetEndCityCode(v string) *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue
	GetEndCityCode() *string
	SetFreePcs(v int64) *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue
	GetFreePcs() *int64
	SetBaggageWeight(v int64) *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue
	GetBaggageWeight() *int64
	SetBaggageUnit(v string) *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue
	GetBaggageUnit() *string
	SetBaggageSize(v string) *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue
	GetBaggageSize() *string
	SetAllWeight(v bool) *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue
	GetAllWeight() *bool
}

type ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue struct {
	// example:
	//
	// 1
	CarryFreepc *int32 `json:"carry_freepc,omitempty" xml:"carry_freepc,omitempty"`
	// example:
	//
	// 10
	CarryBagWeight *int32 `json:"carry_bag_weight,omitempty" xml:"carry_bag_weight,omitempty"`
	// example:
	//
	// 20*20
	CarryBagSize *string `json:"carry_bag_size,omitempty" xml:"carry_bag_size,omitempty"`
	// example:
	//
	// false
	IsAllCarryBagWeight *bool `json:"is_all_carry_bag_weight,omitempty" xml:"is_all_carry_bag_weight,omitempty"`
	// example:
	//
	// CA
	Airline *string `json:"airline,omitempty" xml:"airline,omitempty"`
	// example:
	//
	// BJS
	StartAirport *string `json:"start_airport,omitempty" xml:"start_airport,omitempty"`
	// example:
	//
	// HGH
	EndAirport *string `json:"end_airport,omitempty" xml:"end_airport,omitempty"`
	// example:
	//
	// BJS
	StartCityCode *string `json:"start_city_code,omitempty" xml:"start_city_code,omitempty"`
	// example:
	//
	// HGH
	EndCityCode *string `json:"end_city_code,omitempty" xml:"end_city_code,omitempty"`
	// example:
	//
	// 1
	FreePcs *int64 `json:"free_pcs,omitempty" xml:"free_pcs,omitempty"`
	// example:
	//
	// 30
	BaggageWeight *int64 `json:"baggage_weight,omitempty" xml:"baggage_weight,omitempty"`
	// example:
	//
	// KG
	BaggageUnit *string `json:"baggage_unit,omitempty" xml:"baggage_unit,omitempty"`
	// example:
	//
	// 40*50
	BaggageSize *string `json:"baggage_size,omitempty" xml:"baggage_size,omitempty"`
	// example:
	//
	// false
	AllWeight *bool `json:"all_weight,omitempty" xml:"all_weight,omitempty"`
}

func (s ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue) GoString() string {
	return s.String()
}

func (s *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue) GetCarryFreepc() *int32 {
	return s.CarryFreepc
}

func (s *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue) GetCarryBagWeight() *int32 {
	return s.CarryBagWeight
}

func (s *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue) GetCarryBagSize() *string {
	return s.CarryBagSize
}

func (s *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue) GetIsAllCarryBagWeight() *bool {
	return s.IsAllCarryBagWeight
}

func (s *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue) GetAirline() *string {
	return s.Airline
}

func (s *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue) GetStartAirport() *string {
	return s.StartAirport
}

func (s *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue) GetEndAirport() *string {
	return s.EndAirport
}

func (s *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue) GetStartCityCode() *string {
	return s.StartCityCode
}

func (s *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue) GetEndCityCode() *string {
	return s.EndCityCode
}

func (s *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue) GetFreePcs() *int64 {
	return s.FreePcs
}

func (s *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue) GetBaggageWeight() *int64 {
	return s.BaggageWeight
}

func (s *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue) GetBaggageUnit() *string {
	return s.BaggageUnit
}

func (s *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue) GetBaggageSize() *string {
	return s.BaggageSize
}

func (s *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue) GetAllWeight() *bool {
	return s.AllWeight
}

func (s *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue) SetCarryFreepc(v int32) *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue {
	s.CarryFreepc = &v
	return s
}

func (s *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue) SetCarryBagWeight(v int32) *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue {
	s.CarryBagWeight = &v
	return s
}

func (s *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue) SetCarryBagSize(v string) *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue {
	s.CarryBagSize = &v
	return s
}

func (s *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue) SetIsAllCarryBagWeight(v bool) *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue {
	s.IsAllCarryBagWeight = &v
	return s
}

func (s *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue) SetAirline(v string) *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue {
	s.Airline = &v
	return s
}

func (s *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue) SetStartAirport(v string) *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue {
	s.StartAirport = &v
	return s
}

func (s *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue) SetEndAirport(v string) *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue {
	s.EndAirport = &v
	return s
}

func (s *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue) SetStartCityCode(v string) *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue {
	s.StartCityCode = &v
	return s
}

func (s *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue) SetEndCityCode(v string) *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue {
	s.EndCityCode = &v
	return s
}

func (s *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue) SetFreePcs(v int64) *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue {
	s.FreePcs = &v
	return s
}

func (s *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue) SetBaggageWeight(v int64) *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue {
	s.BaggageWeight = &v
	return s
}

func (s *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue) SetBaggageUnit(v string) *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue {
	s.BaggageUnit = &v
	return s
}

func (s *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue) SetBaggageSize(v string) *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue {
	s.BaggageSize = &v
	return s
}

func (s *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue) SetAllWeight(v bool) *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue {
	s.AllWeight = &v
	return s
}

func (s *ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue) Validate() error {
	return dara.Validate(s)
}
