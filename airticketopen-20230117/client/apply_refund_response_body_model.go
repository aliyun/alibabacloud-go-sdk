// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyRefundResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ApplyRefundResponseBodyData) *ApplyRefundResponseBody
	GetData() *ApplyRefundResponseBodyData
	SetErrorCode(v string) *ApplyRefundResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *ApplyRefundResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *ApplyRefundResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ApplyRefundResponseBody
	GetSuccess() *bool
	SetTracerId(v string) *ApplyRefundResponseBody
	GetTracerId() *string
}

type ApplyRefundResponseBody struct {
	Data *ApplyRefundResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// CreateOrderFailed
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 创建订单失败
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 260E4F99-983D-1919-834C-5C42E98E5B2B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s ApplyRefundResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyRefundResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyRefundResponseBody) GetData() *ApplyRefundResponseBodyData {
	return s.Data
}

func (s *ApplyRefundResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ApplyRefundResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ApplyRefundResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyRefundResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ApplyRefundResponseBody) GetTracerId() *string {
	return s.TracerId
}

func (s *ApplyRefundResponseBody) SetData(v *ApplyRefundResponseBodyData) *ApplyRefundResponseBody {
	s.Data = v
	return s
}

func (s *ApplyRefundResponseBody) SetErrorCode(v string) *ApplyRefundResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ApplyRefundResponseBody) SetErrorMsg(v string) *ApplyRefundResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ApplyRefundResponseBody) SetRequestId(v string) *ApplyRefundResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyRefundResponseBody) SetSuccess(v bool) *ApplyRefundResponseBody {
	s.Success = &v
	return s
}

func (s *ApplyRefundResponseBody) SetTracerId(v string) *ApplyRefundResponseBody {
	s.TracerId = &v
	return s
}

func (s *ApplyRefundResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ApplyRefundResponseBodyData struct {
	// example:
	//
	// 100001
	SellRefundOrderId  *int64                                         `json:"SellRefundOrderId,omitempty" xml:"SellRefundOrderId,omitempty"`
	TotalPenaltyAmount *ApplyRefundResponseBodyDataTotalPenaltyAmount `json:"TotalPenaltyAmount,omitempty" xml:"TotalPenaltyAmount,omitempty" type:"Struct"`
	TotalRefundAmount  *ApplyRefundResponseBodyDataTotalRefundAmount  `json:"TotalRefundAmount,omitempty" xml:"TotalRefundAmount,omitempty" type:"Struct"`
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s ApplyRefundResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ApplyRefundResponseBodyData) GoString() string {
	return s.String()
}

func (s *ApplyRefundResponseBodyData) GetSellRefundOrderId() *int64 {
	return s.SellRefundOrderId
}

func (s *ApplyRefundResponseBodyData) GetTotalPenaltyAmount() *ApplyRefundResponseBodyDataTotalPenaltyAmount {
	return s.TotalPenaltyAmount
}

func (s *ApplyRefundResponseBodyData) GetTotalRefundAmount() *ApplyRefundResponseBodyDataTotalRefundAmount {
	return s.TotalRefundAmount
}

func (s *ApplyRefundResponseBodyData) GetTracerId() *string {
	return s.TracerId
}

func (s *ApplyRefundResponseBodyData) SetSellRefundOrderId(v int64) *ApplyRefundResponseBodyData {
	s.SellRefundOrderId = &v
	return s
}

func (s *ApplyRefundResponseBodyData) SetTotalPenaltyAmount(v *ApplyRefundResponseBodyDataTotalPenaltyAmount) *ApplyRefundResponseBodyData {
	s.TotalPenaltyAmount = v
	return s
}

func (s *ApplyRefundResponseBodyData) SetTotalRefundAmount(v *ApplyRefundResponseBodyDataTotalRefundAmount) *ApplyRefundResponseBodyData {
	s.TotalRefundAmount = v
	return s
}

func (s *ApplyRefundResponseBodyData) SetTracerId(v string) *ApplyRefundResponseBodyData {
	s.TracerId = &v
	return s
}

func (s *ApplyRefundResponseBodyData) Validate() error {
	if s.TotalPenaltyAmount != nil {
		if err := s.TotalPenaltyAmount.Validate(); err != nil {
			return err
		}
	}
	if s.TotalRefundAmount != nil {
		if err := s.TotalRefundAmount.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ApplyRefundResponseBodyDataTotalPenaltyAmount struct {
	// example:
	//
	// 10000
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s ApplyRefundResponseBodyDataTotalPenaltyAmount) String() string {
	return dara.Prettify(s)
}

func (s ApplyRefundResponseBodyDataTotalPenaltyAmount) GoString() string {
	return s.String()
}

func (s *ApplyRefundResponseBodyDataTotalPenaltyAmount) GetAmount() *string {
	return s.Amount
}

func (s *ApplyRefundResponseBodyDataTotalPenaltyAmount) GetCurrency() *string {
	return s.Currency
}

func (s *ApplyRefundResponseBodyDataTotalPenaltyAmount) GetTracerId() *string {
	return s.TracerId
}

func (s *ApplyRefundResponseBodyDataTotalPenaltyAmount) SetAmount(v string) *ApplyRefundResponseBodyDataTotalPenaltyAmount {
	s.Amount = &v
	return s
}

func (s *ApplyRefundResponseBodyDataTotalPenaltyAmount) SetCurrency(v string) *ApplyRefundResponseBodyDataTotalPenaltyAmount {
	s.Currency = &v
	return s
}

func (s *ApplyRefundResponseBodyDataTotalPenaltyAmount) SetTracerId(v string) *ApplyRefundResponseBodyDataTotalPenaltyAmount {
	s.TracerId = &v
	return s
}

func (s *ApplyRefundResponseBodyDataTotalPenaltyAmount) Validate() error {
	return dara.Validate(s)
}

type ApplyRefundResponseBodyDataTotalRefundAmount struct {
	// example:
	//
	// 10000
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s ApplyRefundResponseBodyDataTotalRefundAmount) String() string {
	return dara.Prettify(s)
}

func (s ApplyRefundResponseBodyDataTotalRefundAmount) GoString() string {
	return s.String()
}

func (s *ApplyRefundResponseBodyDataTotalRefundAmount) GetAmount() *string {
	return s.Amount
}

func (s *ApplyRefundResponseBodyDataTotalRefundAmount) GetCurrency() *string {
	return s.Currency
}

func (s *ApplyRefundResponseBodyDataTotalRefundAmount) GetTracerId() *string {
	return s.TracerId
}

func (s *ApplyRefundResponseBodyDataTotalRefundAmount) SetAmount(v string) *ApplyRefundResponseBodyDataTotalRefundAmount {
	s.Amount = &v
	return s
}

func (s *ApplyRefundResponseBodyDataTotalRefundAmount) SetCurrency(v string) *ApplyRefundResponseBodyDataTotalRefundAmount {
	s.Currency = &v
	return s
}

func (s *ApplyRefundResponseBodyDataTotalRefundAmount) SetTracerId(v string) *ApplyRefundResponseBodyDataTotalRefundAmount {
	s.TracerId = &v
	return s
}

func (s *ApplyRefundResponseBodyDataTotalRefundAmount) Validate() error {
	return dara.Validate(s)
}
