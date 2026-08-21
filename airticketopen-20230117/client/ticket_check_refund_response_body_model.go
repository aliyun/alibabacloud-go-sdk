// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketCheckRefundResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *TicketCheckRefundResponseBodyData) *TicketCheckRefundResponseBody
	GetData() *TicketCheckRefundResponseBodyData
	SetErrorCode(v string) *TicketCheckRefundResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *TicketCheckRefundResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *TicketCheckRefundResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TicketCheckRefundResponseBody
	GetSuccess() *bool
}

type TicketCheckRefundResponseBody struct {
	Data *TicketCheckRefundResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// DistributorOrderIdInvalid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 分销商订单号不合法
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TicketCheckRefundResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TicketCheckRefundResponseBody) GoString() string {
	return s.String()
}

func (s *TicketCheckRefundResponseBody) GetData() *TicketCheckRefundResponseBodyData {
	return s.Data
}

func (s *TicketCheckRefundResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *TicketCheckRefundResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *TicketCheckRefundResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TicketCheckRefundResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TicketCheckRefundResponseBody) SetData(v *TicketCheckRefundResponseBodyData) *TicketCheckRefundResponseBody {
	s.Data = v
	return s
}

func (s *TicketCheckRefundResponseBody) SetErrorCode(v string) *TicketCheckRefundResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TicketCheckRefundResponseBody) SetErrorMsg(v string) *TicketCheckRefundResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *TicketCheckRefundResponseBody) SetRequestId(v string) *TicketCheckRefundResponseBody {
	s.RequestId = &v
	return s
}

func (s *TicketCheckRefundResponseBody) SetSuccess(v bool) *TicketCheckRefundResponseBody {
	s.Success = &v
	return s
}

func (s *TicketCheckRefundResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TicketCheckRefundResponseBodyData struct {
	// example:
	//
	// true
	CanRefund    *bool                                          `json:"CanRefund,omitempty" xml:"CanRefund,omitempty"`
	RefundAmount *TicketCheckRefundResponseBodyDataRefundAmount `json:"RefundAmount,omitempty" xml:"RefundAmount,omitempty" type:"Struct"`
	RefundRule   *TicketCheckRefundResponseBodyDataRefundRule   `json:"RefundRule,omitempty" xml:"RefundRule,omitempty" type:"Struct"`
}

func (s TicketCheckRefundResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TicketCheckRefundResponseBodyData) GoString() string {
	return s.String()
}

func (s *TicketCheckRefundResponseBodyData) GetCanRefund() *bool {
	return s.CanRefund
}

func (s *TicketCheckRefundResponseBodyData) GetRefundAmount() *TicketCheckRefundResponseBodyDataRefundAmount {
	return s.RefundAmount
}

func (s *TicketCheckRefundResponseBodyData) GetRefundRule() *TicketCheckRefundResponseBodyDataRefundRule {
	return s.RefundRule
}

func (s *TicketCheckRefundResponseBodyData) SetCanRefund(v bool) *TicketCheckRefundResponseBodyData {
	s.CanRefund = &v
	return s
}

func (s *TicketCheckRefundResponseBodyData) SetRefundAmount(v *TicketCheckRefundResponseBodyDataRefundAmount) *TicketCheckRefundResponseBodyData {
	s.RefundAmount = v
	return s
}

func (s *TicketCheckRefundResponseBodyData) SetRefundRule(v *TicketCheckRefundResponseBodyDataRefundRule) *TicketCheckRefundResponseBodyData {
	s.RefundRule = v
	return s
}

func (s *TicketCheckRefundResponseBodyData) Validate() error {
	if s.RefundAmount != nil {
		if err := s.RefundAmount.Validate(); err != nil {
			return err
		}
	}
	if s.RefundRule != nil {
		if err := s.RefundRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TicketCheckRefundResponseBodyDataRefundAmount struct {
	// example:
	//
	// 10000
	Amount *int64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// USD
	CurrencyCode *string `json:"CurrencyCode,omitempty" xml:"CurrencyCode,omitempty"`
}

func (s TicketCheckRefundResponseBodyDataRefundAmount) String() string {
	return dara.Prettify(s)
}

func (s TicketCheckRefundResponseBodyDataRefundAmount) GoString() string {
	return s.String()
}

func (s *TicketCheckRefundResponseBodyDataRefundAmount) GetAmount() *int64 {
	return s.Amount
}

func (s *TicketCheckRefundResponseBodyDataRefundAmount) GetCurrencyCode() *string {
	return s.CurrencyCode
}

func (s *TicketCheckRefundResponseBodyDataRefundAmount) SetAmount(v int64) *TicketCheckRefundResponseBodyDataRefundAmount {
	s.Amount = &v
	return s
}

func (s *TicketCheckRefundResponseBodyDataRefundAmount) SetCurrencyCode(v string) *TicketCheckRefundResponseBodyDataRefundAmount {
	s.CurrencyCode = &v
	return s
}

func (s *TicketCheckRefundResponseBodyDataRefundAmount) Validate() error {
	return dara.Validate(s)
}

type TicketCheckRefundResponseBodyDataRefundRule struct {
	RefundStageRules []*TicketCheckRefundResponseBodyDataRefundRuleRefundStageRules `json:"RefundStageRules,omitempty" xml:"RefundStageRules,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	RefundType *int32 `json:"RefundType,omitempty" xml:"RefundType,omitempty"`
}

func (s TicketCheckRefundResponseBodyDataRefundRule) String() string {
	return dara.Prettify(s)
}

func (s TicketCheckRefundResponseBodyDataRefundRule) GoString() string {
	return s.String()
}

func (s *TicketCheckRefundResponseBodyDataRefundRule) GetRefundStageRules() []*TicketCheckRefundResponseBodyDataRefundRuleRefundStageRules {
	return s.RefundStageRules
}

func (s *TicketCheckRefundResponseBodyDataRefundRule) GetRefundType() *int32 {
	return s.RefundType
}

func (s *TicketCheckRefundResponseBodyDataRefundRule) SetRefundStageRules(v []*TicketCheckRefundResponseBodyDataRefundRuleRefundStageRules) *TicketCheckRefundResponseBodyDataRefundRule {
	s.RefundStageRules = v
	return s
}

func (s *TicketCheckRefundResponseBodyDataRefundRule) SetRefundType(v int32) *TicketCheckRefundResponseBodyDataRefundRule {
	s.RefundType = &v
	return s
}

func (s *TicketCheckRefundResponseBodyDataRefundRule) Validate() error {
	if s.RefundStageRules != nil {
		for _, item := range s.RefundStageRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TicketCheckRefundResponseBodyDataRefundRuleRefundStageRules struct {
	Fee *float64 `json:"Fee,omitempty" xml:"Fee,omitempty"`
	// example:
	//
	// 1
	FeeBase *int32 `json:"FeeBase,omitempty" xml:"FeeBase,omitempty"`
	// example:
	//
	// 1
	FeeType *int32                                                           `json:"FeeType,omitempty" xml:"FeeType,omitempty"`
	From    *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesFrom `json:"From,omitempty" xml:"From,omitempty" type:"Struct"`
	To      *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesTo   `json:"To,omitempty" xml:"To,omitempty" type:"Struct"`
}

func (s TicketCheckRefundResponseBodyDataRefundRuleRefundStageRules) String() string {
	return dara.Prettify(s)
}

func (s TicketCheckRefundResponseBodyDataRefundRuleRefundStageRules) GoString() string {
	return s.String()
}

func (s *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRules) GetFee() *float64 {
	return s.Fee
}

func (s *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRules) GetFeeBase() *int32 {
	return s.FeeBase
}

func (s *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRules) GetFeeType() *int32 {
	return s.FeeType
}

func (s *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRules) GetFrom() *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesFrom {
	return s.From
}

func (s *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRules) GetTo() *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesTo {
	return s.To
}

func (s *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRules) SetFee(v float64) *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRules {
	s.Fee = &v
	return s
}

func (s *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRules) SetFeeBase(v int32) *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRules {
	s.FeeBase = &v
	return s
}

func (s *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRules) SetFeeType(v int32) *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRules {
	s.FeeType = &v
	return s
}

func (s *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRules) SetFrom(v *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesFrom) *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRules {
	s.From = v
	return s
}

func (s *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRules) SetTo(v *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesTo) *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRules {
	s.To = v
	return s
}

func (s *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRules) Validate() error {
	if s.From != nil {
		if err := s.From.Validate(); err != nil {
			return err
		}
	}
	if s.To != nil {
		if err := s.To.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesFrom struct {
	// example:
	//
	// 1
	Anchor *int32 `json:"Anchor,omitempty" xml:"Anchor,omitempty"`
	// example:
	//
	// 2026-01-01
	FixedTime *string `json:"FixedTime,omitempty" xml:"FixedTime,omitempty"`
	// example:
	//
	// 18:00
	OffsetDayOfTime *string `json:"OffsetDayOfTime,omitempty" xml:"OffsetDayOfTime,omitempty"`
	// example:
	//
	// 1
	OffsetUnit *int32 `json:"OffsetUnit,omitempty" xml:"OffsetUnit,omitempty"`
	// example:
	//
	// 1
	OffsetValue *int32 `json:"OffsetValue,omitempty" xml:"OffsetValue,omitempty"`
}

func (s TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesFrom) String() string {
	return dara.Prettify(s)
}

func (s TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesFrom) GoString() string {
	return s.String()
}

func (s *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesFrom) GetAnchor() *int32 {
	return s.Anchor
}

func (s *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesFrom) GetFixedTime() *string {
	return s.FixedTime
}

func (s *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesFrom) GetOffsetDayOfTime() *string {
	return s.OffsetDayOfTime
}

func (s *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesFrom) GetOffsetUnit() *int32 {
	return s.OffsetUnit
}

func (s *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesFrom) GetOffsetValue() *int32 {
	return s.OffsetValue
}

func (s *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesFrom) SetAnchor(v int32) *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesFrom {
	s.Anchor = &v
	return s
}

func (s *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesFrom) SetFixedTime(v string) *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesFrom {
	s.FixedTime = &v
	return s
}

func (s *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesFrom) SetOffsetDayOfTime(v string) *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesFrom {
	s.OffsetDayOfTime = &v
	return s
}

func (s *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesFrom) SetOffsetUnit(v int32) *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesFrom {
	s.OffsetUnit = &v
	return s
}

func (s *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesFrom) SetOffsetValue(v int32) *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesFrom {
	s.OffsetValue = &v
	return s
}

func (s *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesFrom) Validate() error {
	return dara.Validate(s)
}

type TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesTo struct {
	// example:
	//
	// 1
	Anchor *int32 `json:"Anchor,omitempty" xml:"Anchor,omitempty"`
	// example:
	//
	// 2026-01-01
	FixedTime *string `json:"FixedTime,omitempty" xml:"FixedTime,omitempty"`
	// example:
	//
	// 18:00
	OffsetDayOfTime *string `json:"OffsetDayOfTime,omitempty" xml:"OffsetDayOfTime,omitempty"`
	// example:
	//
	// 1
	OffsetUnit *int32 `json:"OffsetUnit,omitempty" xml:"OffsetUnit,omitempty"`
	// example:
	//
	// 1
	OffsetValue *int32 `json:"OffsetValue,omitempty" xml:"OffsetValue,omitempty"`
}

func (s TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesTo) String() string {
	return dara.Prettify(s)
}

func (s TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesTo) GoString() string {
	return s.String()
}

func (s *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesTo) GetAnchor() *int32 {
	return s.Anchor
}

func (s *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesTo) GetFixedTime() *string {
	return s.FixedTime
}

func (s *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesTo) GetOffsetDayOfTime() *string {
	return s.OffsetDayOfTime
}

func (s *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesTo) GetOffsetUnit() *int32 {
	return s.OffsetUnit
}

func (s *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesTo) GetOffsetValue() *int32 {
	return s.OffsetValue
}

func (s *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesTo) SetAnchor(v int32) *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesTo {
	s.Anchor = &v
	return s
}

func (s *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesTo) SetFixedTime(v string) *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesTo {
	s.FixedTime = &v
	return s
}

func (s *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesTo) SetOffsetDayOfTime(v string) *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesTo {
	s.OffsetDayOfTime = &v
	return s
}

func (s *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesTo) SetOffsetUnit(v int32) *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesTo {
	s.OffsetUnit = &v
	return s
}

func (s *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesTo) SetOffsetValue(v int32) *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesTo {
	s.OffsetValue = &v
	return s
}

func (s *TicketCheckRefundResponseBodyDataRefundRuleRefundStageRulesTo) Validate() error {
	return dara.Validate(s)
}
