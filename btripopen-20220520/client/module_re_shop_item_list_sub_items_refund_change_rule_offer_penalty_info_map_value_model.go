// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetStruct(v bool) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetStruct() *bool
	SetCancelFeeInd(v bool) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetCancelFeeInd() *bool
	SetChangeFeeInd(v bool) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetChangeFeeInd() *bool
	SetUpgradeFeeInd(v bool) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetUpgradeFeeInd() *bool
	SetReissueInd(v bool) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetReissueInd() *bool
	SetPenaltyTypeCode(v int32) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetPenaltyTypeCode() *int32
	SetPenaltyApplyRangeCode(v int32) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetPenaltyApplyRangeCode() *int32
	SetPenaltyChargeTypeCode(v int32) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetPenaltyChargeTypeCode() *int32
	SetFee(v float64) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetFee() *float64
	SetCurrency(v string) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetCurrency() *string
	SetPenaltyPercent(v float64) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetPenaltyPercent() *float64
	SetStartTime(v int32) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetStartTime() *int32
	SetEndTime(v int32) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetEndTime() *int32
	SetTimeUnitCode(v int32) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetTimeUnitCode() *int32
	SetTitle(v string) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetTitle() *string
	SetDepTime(v string) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetDepTime() *string
	SetSegmentNumber(v string) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetSegmentNumber() *string
	SetDescInfos(v map[string]*string) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue
	GetDescInfos() map[string]*string
}

type ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue struct {
	// Indicates whether the rule is applicable.
	//
	// example:
	//
	// true
	Struct *bool `json:"struct,omitempty" xml:"struct,omitempty"`
	// Indicates whether refund is supported.
	//
	// example:
	//
	// false
	CancelFeeInd *bool `json:"cancel_fee_ind,omitempty" xml:"cancel_fee_ind,omitempty"`
	// Indicates whether date change is supported.
	//
	// example:
	//
	// false
	ChangeFeeInd *bool `json:"change_fee_ind,omitempty" xml:"change_fee_ind,omitempty"`
	// Indicates whether upgrade is supported.
	//
	// example:
	//
	// false
	UpgradeFeeInd *bool `json:"upgrade_fee_ind,omitempty" xml:"upgrade_fee_ind,omitempty"`
	// Indicates whether reissue is supported.
	//
	// example:
	//
	// false
	ReissueInd *bool `json:"reissue_ind,omitempty" xml:"reissue_ind,omitempty"`
	// The rule type. Valid values:
	//
	// - 0: Refund fee.
	//
	// - 1: Change fee.
	//
	// - 2: No-show penalty.
	//
	// - 3: Other.
	//
	// - 4: Upgrade fee.
	//
	// - 5: Endorsement.
	//
	// - 6: Deduction for used segments.
	//
	// - 100: Tax refund.
	//
	// example:
	//
	// 1
	PenaltyTypeCode *int32 `json:"penalty_type_code,omitempty" xml:"penalty_type_code,omitempty"`
	// The applicability scope of the rule. Valid values:
	//
	//
	//
	// - 1: All unused.
	//
	// - 2: Partially unused.
	//
	// - 3: Outbound.
	//
	// - 4: Inbound.
	//
	// example:
	//
	// 1
	PenaltyApplyRangeCode *int32 `json:"penalty_apply_range_code,omitempty" xml:"penalty_apply_range_code,omitempty"`
	// The charge method of the rule. Valid values:
	//
	// - 0: Charged per whole trip.
	//
	// - 1: Charged per direction.
	//
	// - 2: Charged per segment.
	//
	// example:
	//
	// 0
	PenaltyChargeTypeCode *int32 `json:"penalty_charge_type_code,omitempty" xml:"penalty_charge_type_code,omitempty"`
	// The fee amount.
	//
	// example:
	//
	// 10
	Fee *float64 `json:"fee,omitempty" xml:"fee,omitempty"`
	// The currency of the fee.
	//
	// example:
	//
	// CNY
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
	// The fee percentage.
	//
	// example:
	//
	// 10
	PenaltyPercent *float64 `json:"penalty_percent,omitempty" xml:"penalty_percent,omitempty"`
	// The start time of the rule time range.
	//
	// example:
	//
	// -720
	StartTime *int32 `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// The end time of the rule time range.
	//
	// example:
	//
	// -168
	EndTime *int32 `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// The time unit. Valid values:
	//
	// - 0: Hours.
	//
	// - 1: Days.
	//
	// example:
	//
	// 1
	TimeUnitCode *int32 `json:"time_unit_code,omitempty" xml:"time_unit_code,omitempty"`
	// The rule title.
	//
	// example:
	//
	// 退票费用
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 起飞时间
	//
	// example:
	//
	// 2025-10-01 01:00:00
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// 航段序号，
	//
	// - OUTBOUND_FIRST("去程第一段")
	//
	// - OUTBOUND_SECOND("去程第二段")
	//
	// - INBOUND_FIRST("回程第一段")
	//
	// - INBOUND_SECOND("回程第二段")
	//
	// example:
	//
	// OUTBOUND_FIRST
	SegmentNumber *string `json:"segment_number,omitempty" xml:"segment_number,omitempty"`
	// 各类非结构化补充说明
	DescInfos map[string]*string `json:"desc_infos,omitempty" xml:"desc_infos,omitempty"`
}

func (s ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GoString() string {
	return s.String()
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetStruct() *bool {
	return s.Struct
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetCancelFeeInd() *bool {
	return s.CancelFeeInd
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetChangeFeeInd() *bool {
	return s.ChangeFeeInd
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetUpgradeFeeInd() *bool {
	return s.UpgradeFeeInd
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetReissueInd() *bool {
	return s.ReissueInd
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetPenaltyTypeCode() *int32 {
	return s.PenaltyTypeCode
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetPenaltyApplyRangeCode() *int32 {
	return s.PenaltyApplyRangeCode
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetPenaltyChargeTypeCode() *int32 {
	return s.PenaltyChargeTypeCode
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetFee() *float64 {
	return s.Fee
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetCurrency() *string {
	return s.Currency
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetPenaltyPercent() *float64 {
	return s.PenaltyPercent
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetStartTime() *int32 {
	return s.StartTime
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetEndTime() *int32 {
	return s.EndTime
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetTimeUnitCode() *int32 {
	return s.TimeUnitCode
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetTitle() *string {
	return s.Title
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetDepTime() *string {
	return s.DepTime
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetSegmentNumber() *string {
	return s.SegmentNumber
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) GetDescInfos() map[string]*string {
	return s.DescInfos
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetStruct(v bool) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.Struct = &v
	return s
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetCancelFeeInd(v bool) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.CancelFeeInd = &v
	return s
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetChangeFeeInd(v bool) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.ChangeFeeInd = &v
	return s
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetUpgradeFeeInd(v bool) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.UpgradeFeeInd = &v
	return s
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetReissueInd(v bool) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.ReissueInd = &v
	return s
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetPenaltyTypeCode(v int32) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.PenaltyTypeCode = &v
	return s
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetPenaltyApplyRangeCode(v int32) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.PenaltyApplyRangeCode = &v
	return s
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetPenaltyChargeTypeCode(v int32) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.PenaltyChargeTypeCode = &v
	return s
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetFee(v float64) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.Fee = &v
	return s
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetCurrency(v string) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.Currency = &v
	return s
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetPenaltyPercent(v float64) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.PenaltyPercent = &v
	return s
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetStartTime(v int32) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.StartTime = &v
	return s
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetEndTime(v int32) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.EndTime = &v
	return s
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetTimeUnitCode(v int32) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.TimeUnitCode = &v
	return s
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetTitle(v string) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.Title = &v
	return s
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetDepTime(v string) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.DepTime = &v
	return s
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetSegmentNumber(v string) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.SegmentNumber = &v
	return s
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) SetDescInfos(v map[string]*string) *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	s.DescInfos = v
	return s
}

func (s *ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) Validate() error {
	return dara.Validate(s)
}
