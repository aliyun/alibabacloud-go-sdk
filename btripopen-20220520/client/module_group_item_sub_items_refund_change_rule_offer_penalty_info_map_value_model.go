// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetStruct(v bool) *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetStruct() *bool
	SetCancelFeeInd(v bool) *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetCancelFeeInd() *bool
	SetChangeFeeInd(v bool) *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetChangeFeeInd() *bool
	SetUpgradeFeeInd(v bool) *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetUpgradeFeeInd() *bool
	SetReissueInd(v bool) *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetReissueInd() *bool
	SetPenaltyTypeCode(v int32) *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetPenaltyTypeCode() *int32
	SetPenaltyApplyRangeCode(v int32) *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetPenaltyApplyRangeCode() *int32
	SetPenaltyChargeTypeCode(v int32) *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetPenaltyChargeTypeCode() *int32
	SetFee(v float64) *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetFee() *float64
	SetCurrency(v string) *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetCurrency() *string
	SetPenaltyPercent(v float64) *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetPenaltyPercent() *float64
	SetStartTime(v int32) *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetStartTime() *int32
	SetEndTime(v int32) *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetEndTime() *int32
	SetTimeUnitCode(v int32) *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetTimeUnitCode() *int32
}

type ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue struct {
	// example:
	//
	// true
	Struct *bool `json:"struct,omitempty" xml:"struct,omitempty"`
	// example:
	//
	// true
	CancelFeeInd *bool `json:"cancel_fee_ind,omitempty" xml:"cancel_fee_ind,omitempty"`
	// example:
	//
	// true
	ChangeFeeInd *bool `json:"change_fee_ind,omitempty" xml:"change_fee_ind,omitempty"`
	// example:
	//
	// true
	UpgradeFeeInd *bool `json:"upgrade_fee_ind,omitempty" xml:"upgrade_fee_ind,omitempty"`
	// example:
	//
	// true
	ReissueInd *bool `json:"reissue_ind,omitempty" xml:"reissue_ind,omitempty"`
	// example:
	//
	// 1
	PenaltyTypeCode *int32 `json:"penalty_type_code,omitempty" xml:"penalty_type_code,omitempty"`
	// example:
	//
	// 0
	PenaltyApplyRangeCode *int32 `json:"penalty_apply_range_code,omitempty" xml:"penalty_apply_range_code,omitempty"`
	// example:
	//
	// 1
	PenaltyChargeTypeCode *int32 `json:"penalty_charge_type_code,omitempty" xml:"penalty_charge_type_code,omitempty"`
	// example:
	//
	// 300
	Fee *float64 `json:"fee,omitempty" xml:"fee,omitempty"`
	// example:
	//
	// CNY
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
	// example:
	//
	// 30
	PenaltyPercent *float64 `json:"penalty_percent,omitempty" xml:"penalty_percent,omitempty"`
	// example:
	//
	// 48
	StartTime *int32 `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// example:
	//
	// 0
	EndTime *int32 `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// example:
	//
	// 0
	TimeUnitCode *int32 `json:"time_unit_code,omitempty" xml:"time_unit_code,omitempty"`
}

func (s ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GoString() string {
	return s.String()
}

func (s *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetStruct() *bool {
	return s.Struct
}

func (s *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetCancelFeeInd() *bool {
	return s.CancelFeeInd
}

func (s *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetChangeFeeInd() *bool {
	return s.ChangeFeeInd
}

func (s *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetUpgradeFeeInd() *bool {
	return s.UpgradeFeeInd
}

func (s *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetReissueInd() *bool {
	return s.ReissueInd
}

func (s *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetPenaltyTypeCode() *int32 {
	return s.PenaltyTypeCode
}

func (s *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetPenaltyApplyRangeCode() *int32 {
	return s.PenaltyApplyRangeCode
}

func (s *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetPenaltyChargeTypeCode() *int32 {
	return s.PenaltyChargeTypeCode
}

func (s *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetFee() *float64 {
	return s.Fee
}

func (s *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetCurrency() *string {
	return s.Currency
}

func (s *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetPenaltyPercent() *float64 {
	return s.PenaltyPercent
}

func (s *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetStartTime() *int32 {
	return s.StartTime
}

func (s *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetEndTime() *int32 {
	return s.EndTime
}

func (s *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetTimeUnitCode() *int32 {
	return s.TimeUnitCode
}

func (s *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetStruct(v bool) *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.Struct = &v
	return s
}

func (s *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetCancelFeeInd(v bool) *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.CancelFeeInd = &v
	return s
}

func (s *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetChangeFeeInd(v bool) *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.ChangeFeeInd = &v
	return s
}

func (s *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetUpgradeFeeInd(v bool) *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.UpgradeFeeInd = &v
	return s
}

func (s *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetReissueInd(v bool) *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.ReissueInd = &v
	return s
}

func (s *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetPenaltyTypeCode(v int32) *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.PenaltyTypeCode = &v
	return s
}

func (s *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetPenaltyApplyRangeCode(v int32) *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.PenaltyApplyRangeCode = &v
	return s
}

func (s *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetPenaltyChargeTypeCode(v int32) *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.PenaltyChargeTypeCode = &v
	return s
}

func (s *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetFee(v float64) *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.Fee = &v
	return s
}

func (s *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetCurrency(v string) *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.Currency = &v
	return s
}

func (s *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetPenaltyPercent(v float64) *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.PenaltyPercent = &v
	return s
}

func (s *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetStartTime(v int32) *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.StartTime = &v
	return s
}

func (s *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetEndTime(v int32) *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.EndTime = &v
	return s
}

func (s *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetTimeUnitCode(v int32) *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.TimeUnitCode = &v
	return s
}

func (s *ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) Validate() error {
	return dara.Validate(s)
}
