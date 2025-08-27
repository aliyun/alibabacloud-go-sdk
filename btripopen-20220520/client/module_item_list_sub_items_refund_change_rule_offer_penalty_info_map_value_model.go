// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetStruct(v bool) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetStruct() *bool
	SetCancelFeeInd(v bool) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetCancelFeeInd() *bool
	SetChangeFeeInd(v bool) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetChangeFeeInd() *bool
	SetUpgradeFeeInd(v bool) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetUpgradeFeeInd() *bool
	SetReissueInd(v bool) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetReissueInd() *bool
	SetPenaltyTypeCode(v int32) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetPenaltyTypeCode() *int32
	SetPenaltyApplyRangeCode(v int32) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetPenaltyApplyRangeCode() *int32
	SetPenaltyChargeTypeCode(v int32) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetPenaltyChargeTypeCode() *int32
	SetFee(v float64) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetFee() *float64
	SetCurrency(v string) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetCurrency() *string
	SetPenaltyPercent(v float64) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetPenaltyPercent() *float64
	SetStartTime(v int32) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetStartTime() *int32
	SetEndTime(v int32) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetEndTime() *int32
	SetTimeUnitCode(v int32) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetTimeUnitCode() *int32
	SetTitle(v string) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetTitle() *string
	SetDepTime(v string) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetDepTime() *string
	SetSegmentNumber(v string) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetSegmentNumber() *string
	SetDescInfos(v map[string]*string) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetDescInfos() map[string]*string
}

type ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue struct {
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
	TimeUnitCode  *int32             `json:"time_unit_code,omitempty" xml:"time_unit_code,omitempty"`
	Title         *string            `json:"title,omitempty" xml:"title,omitempty"`
	DepTime       *string            `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	SegmentNumber *string            `json:"segment_number,omitempty" xml:"segment_number,omitempty"`
	DescInfos     map[string]*string `json:"desc_infos,omitempty" xml:"desc_infos,omitempty"`
}

func (s ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GoString() string {
	return s.String()
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetStruct() *bool {
	return s.Struct
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetCancelFeeInd() *bool {
	return s.CancelFeeInd
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetChangeFeeInd() *bool {
	return s.ChangeFeeInd
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetUpgradeFeeInd() *bool {
	return s.UpgradeFeeInd
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetReissueInd() *bool {
	return s.ReissueInd
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetPenaltyTypeCode() *int32 {
	return s.PenaltyTypeCode
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetPenaltyApplyRangeCode() *int32 {
	return s.PenaltyApplyRangeCode
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetPenaltyChargeTypeCode() *int32 {
	return s.PenaltyChargeTypeCode
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetFee() *float64 {
	return s.Fee
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetCurrency() *string {
	return s.Currency
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetPenaltyPercent() *float64 {
	return s.PenaltyPercent
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetStartTime() *int32 {
	return s.StartTime
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetEndTime() *int32 {
	return s.EndTime
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetTimeUnitCode() *int32 {
	return s.TimeUnitCode
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetTitle() *string {
	return s.Title
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetDepTime() *string {
	return s.DepTime
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetSegmentNumber() *string {
	return s.SegmentNumber
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetDescInfos() map[string]*string {
	return s.DescInfos
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetStruct(v bool) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.Struct = &v
	return s
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetCancelFeeInd(v bool) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.CancelFeeInd = &v
	return s
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetChangeFeeInd(v bool) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.ChangeFeeInd = &v
	return s
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetUpgradeFeeInd(v bool) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.UpgradeFeeInd = &v
	return s
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetReissueInd(v bool) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.ReissueInd = &v
	return s
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetPenaltyTypeCode(v int32) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.PenaltyTypeCode = &v
	return s
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetPenaltyApplyRangeCode(v int32) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.PenaltyApplyRangeCode = &v
	return s
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetPenaltyChargeTypeCode(v int32) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.PenaltyChargeTypeCode = &v
	return s
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetFee(v float64) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.Fee = &v
	return s
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetCurrency(v string) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.Currency = &v
	return s
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetPenaltyPercent(v float64) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.PenaltyPercent = &v
	return s
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetStartTime(v int32) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.StartTime = &v
	return s
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetEndTime(v int32) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.EndTime = &v
	return s
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetTimeUnitCode(v int32) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.TimeUnitCode = &v
	return s
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetTitle(v string) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.Title = &v
	return s
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetDepTime(v string) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.DepTime = &v
	return s
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetSegmentNumber(v string) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.SegmentNumber = &v
	return s
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetDescInfos(v map[string]*string) *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.DescInfos = v
	return s
}

func (s *ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) Validate() error {
	return dara.Validate(s)
}
