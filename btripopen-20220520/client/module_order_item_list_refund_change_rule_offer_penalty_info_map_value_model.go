// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetStruct(v bool) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue
	GetStruct() *bool
	SetCancelFeeInd(v bool) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue
	GetCancelFeeInd() *bool
	SetChangeFeeInd(v bool) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue
	GetChangeFeeInd() *bool
	SetUpgradeFeeInd(v bool) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue
	GetUpgradeFeeInd() *bool
	SetReissueInd(v bool) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue
	GetReissueInd() *bool
	SetPenaltyTypeCode(v int32) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue
	GetPenaltyTypeCode() *int32
	SetPenaltyApplyRangeCode(v int32) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue
	GetPenaltyApplyRangeCode() *int32
	SetPenaltyChargeTypeCode(v int32) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue
	GetPenaltyChargeTypeCode() *int32
	SetFee(v float64) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue
	GetFee() *float64
	SetCurrency(v string) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue
	GetCurrency() *string
	SetPenaltyPercent(v float64) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue
	GetPenaltyPercent() *float64
	SetStartTime(v int32) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue
	GetStartTime() *int32
	SetEndTime(v int32) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue
	GetEndTime() *int32
	SetTimeUnitCode(v int32) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue
	GetTimeUnitCode() *int32
	SetTitle(v string) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue
	GetTitle() *string
	SetDepTime(v string) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue
	GetDepTime() *string
	SetSegmentNumber(v string) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue
	GetSegmentNumber() *string
	SetDescInfos(v map[string]*string) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue
	GetDescInfos() map[string]*string
}

type ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue struct {
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
	// false
	ChangeFeeInd *bool `json:"change_fee_ind,omitempty" xml:"change_fee_ind,omitempty"`
	// example:
	//
	// false
	UpgradeFeeInd *bool `json:"upgrade_fee_ind,omitempty" xml:"upgrade_fee_ind,omitempty"`
	// example:
	//
	// false
	ReissueInd *bool `json:"reissue_ind,omitempty" xml:"reissue_ind,omitempty"`
	// example:
	//
	// 0
	PenaltyTypeCode *int32 `json:"penalty_type_code,omitempty" xml:"penalty_type_code,omitempty"`
	// example:
	//
	// 1
	PenaltyApplyRangeCode *int32 `json:"penalty_apply_range_code,omitempty" xml:"penalty_apply_range_code,omitempty"`
	// example:
	//
	// 0
	PenaltyChargeTypeCode *int32 `json:"penalty_charge_type_code,omitempty" xml:"penalty_charge_type_code,omitempty"`
	// example:
	//
	// 1
	Fee *float64 `json:"fee,omitempty" xml:"fee,omitempty"`
	// example:
	//
	// CNY
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
	// example:
	//
	// 0
	PenaltyPercent *float64 `json:"penalty_percent,omitempty" xml:"penalty_percent,omitempty"`
	// example:
	//
	// 1
	StartTime *int32 `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// example:
	//
	// 10
	EndTime *int32 `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// example:
	//
	// 0
	TimeUnitCode  *int32             `json:"time_unit_code,omitempty" xml:"time_unit_code,omitempty"`
	Title         *string            `json:"title,omitempty" xml:"title,omitempty"`
	DepTime       *string            `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	SegmentNumber *string            `json:"segment_number,omitempty" xml:"segment_number,omitempty"`
	DescInfos     map[string]*string `json:"desc_infos,omitempty" xml:"desc_infos,omitempty"`
}

func (s ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) GoString() string {
	return s.String()
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) GetStruct() *bool {
	return s.Struct
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) GetCancelFeeInd() *bool {
	return s.CancelFeeInd
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) GetChangeFeeInd() *bool {
	return s.ChangeFeeInd
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) GetUpgradeFeeInd() *bool {
	return s.UpgradeFeeInd
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) GetReissueInd() *bool {
	return s.ReissueInd
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) GetPenaltyTypeCode() *int32 {
	return s.PenaltyTypeCode
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) GetPenaltyApplyRangeCode() *int32 {
	return s.PenaltyApplyRangeCode
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) GetPenaltyChargeTypeCode() *int32 {
	return s.PenaltyChargeTypeCode
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) GetFee() *float64 {
	return s.Fee
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) GetCurrency() *string {
	return s.Currency
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) GetPenaltyPercent() *float64 {
	return s.PenaltyPercent
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) GetStartTime() *int32 {
	return s.StartTime
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) GetEndTime() *int32 {
	return s.EndTime
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) GetTimeUnitCode() *int32 {
	return s.TimeUnitCode
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) GetTitle() *string {
	return s.Title
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) GetDepTime() *string {
	return s.DepTime
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) GetSegmentNumber() *string {
	return s.SegmentNumber
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) GetDescInfos() map[string]*string {
	return s.DescInfos
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) SetStruct(v bool) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue {
	s.Struct = &v
	return s
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) SetCancelFeeInd(v bool) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue {
	s.CancelFeeInd = &v
	return s
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) SetChangeFeeInd(v bool) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue {
	s.ChangeFeeInd = &v
	return s
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) SetUpgradeFeeInd(v bool) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue {
	s.UpgradeFeeInd = &v
	return s
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) SetReissueInd(v bool) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue {
	s.ReissueInd = &v
	return s
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) SetPenaltyTypeCode(v int32) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue {
	s.PenaltyTypeCode = &v
	return s
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) SetPenaltyApplyRangeCode(v int32) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue {
	s.PenaltyApplyRangeCode = &v
	return s
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) SetPenaltyChargeTypeCode(v int32) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue {
	s.PenaltyChargeTypeCode = &v
	return s
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) SetFee(v float64) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue {
	s.Fee = &v
	return s
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) SetCurrency(v string) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue {
	s.Currency = &v
	return s
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) SetPenaltyPercent(v float64) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue {
	s.PenaltyPercent = &v
	return s
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) SetStartTime(v int32) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue {
	s.StartTime = &v
	return s
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) SetEndTime(v int32) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue {
	s.EndTime = &v
	return s
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) SetTimeUnitCode(v int32) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue {
	s.TimeUnitCode = &v
	return s
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) SetTitle(v string) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue {
	s.Title = &v
	return s
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) SetDepTime(v string) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue {
	s.DepTime = &v
	return s
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) SetSegmentNumber(v string) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue {
	s.SegmentNumber = &v
	return s
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) SetDescInfos(v map[string]*string) *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue {
	s.DescInfos = v
	return s
}

func (s *ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) Validate() error {
	return dara.Validate(s)
}
