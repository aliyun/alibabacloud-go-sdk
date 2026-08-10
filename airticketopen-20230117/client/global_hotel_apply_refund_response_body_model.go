// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelApplyRefundResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GlobalHotelApplyRefundResponseBodyData) *GlobalHotelApplyRefundResponseBody
	GetData() *GlobalHotelApplyRefundResponseBodyData
	SetErrorCode(v string) *GlobalHotelApplyRefundResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GlobalHotelApplyRefundResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *GlobalHotelApplyRefundResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GlobalHotelApplyRefundResponseBody
	GetSuccess() *bool
	SetTracerId(v string) *GlobalHotelApplyRefundResponseBody
	GetTracerId() *string
}

type GlobalHotelApplyRefundResponseBody struct {
	// The business data.
	Data *GlobalHotelApplyRefundResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// CreateOrderFailed
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Failed to create order
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The unique request ID.
	//
	// example:
	//
	// 260E4F99-983D-1919-834C-5C42E98E5B2B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// TraceId
	//
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelApplyRefundResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelApplyRefundResponseBody) GoString() string {
	return s.String()
}

func (s *GlobalHotelApplyRefundResponseBody) GetData() *GlobalHotelApplyRefundResponseBodyData {
	return s.Data
}

func (s *GlobalHotelApplyRefundResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GlobalHotelApplyRefundResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GlobalHotelApplyRefundResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GlobalHotelApplyRefundResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GlobalHotelApplyRefundResponseBody) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelApplyRefundResponseBody) SetData(v *GlobalHotelApplyRefundResponseBodyData) *GlobalHotelApplyRefundResponseBody {
	s.Data = v
	return s
}

func (s *GlobalHotelApplyRefundResponseBody) SetErrorCode(v string) *GlobalHotelApplyRefundResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GlobalHotelApplyRefundResponseBody) SetErrorMsg(v string) *GlobalHotelApplyRefundResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GlobalHotelApplyRefundResponseBody) SetRequestId(v string) *GlobalHotelApplyRefundResponseBody {
	s.RequestId = &v
	return s
}

func (s *GlobalHotelApplyRefundResponseBody) SetSuccess(v bool) *GlobalHotelApplyRefundResponseBody {
	s.Success = &v
	return s
}

func (s *GlobalHotelApplyRefundResponseBody) SetTracerId(v string) *GlobalHotelApplyRefundResponseBody {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelApplyRefundResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GlobalHotelApplyRefundResponseBodyData struct {
	// The after-sales refund order ID.
	//
	// example:
	//
	// 100001
	SellRefundOrderId *int64 `json:"SellRefundOrderId,omitempty" xml:"SellRefundOrderId,omitempty"`
	// The total penalty amount.
	TotalPenaltyAmount *GlobalHotelApplyRefundResponseBodyDataTotalPenaltyAmount `json:"TotalPenaltyAmount,omitempty" xml:"TotalPenaltyAmount,omitempty" type:"Struct"`
	// The total refund amount.
	TotalRefundAmount *GlobalHotelApplyRefundResponseBodyDataTotalRefundAmount `json:"TotalRefundAmount,omitempty" xml:"TotalRefundAmount,omitempty" type:"Struct"`
	// TraceId
	//
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelApplyRefundResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelApplyRefundResponseBodyData) GoString() string {
	return s.String()
}

func (s *GlobalHotelApplyRefundResponseBodyData) GetSellRefundOrderId() *int64 {
	return s.SellRefundOrderId
}

func (s *GlobalHotelApplyRefundResponseBodyData) GetTotalPenaltyAmount() *GlobalHotelApplyRefundResponseBodyDataTotalPenaltyAmount {
	return s.TotalPenaltyAmount
}

func (s *GlobalHotelApplyRefundResponseBodyData) GetTotalRefundAmount() *GlobalHotelApplyRefundResponseBodyDataTotalRefundAmount {
	return s.TotalRefundAmount
}

func (s *GlobalHotelApplyRefundResponseBodyData) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelApplyRefundResponseBodyData) SetSellRefundOrderId(v int64) *GlobalHotelApplyRefundResponseBodyData {
	s.SellRefundOrderId = &v
	return s
}

func (s *GlobalHotelApplyRefundResponseBodyData) SetTotalPenaltyAmount(v *GlobalHotelApplyRefundResponseBodyDataTotalPenaltyAmount) *GlobalHotelApplyRefundResponseBodyData {
	s.TotalPenaltyAmount = v
	return s
}

func (s *GlobalHotelApplyRefundResponseBodyData) SetTotalRefundAmount(v *GlobalHotelApplyRefundResponseBodyDataTotalRefundAmount) *GlobalHotelApplyRefundResponseBodyData {
	s.TotalRefundAmount = v
	return s
}

func (s *GlobalHotelApplyRefundResponseBodyData) SetTracerId(v string) *GlobalHotelApplyRefundResponseBodyData {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelApplyRefundResponseBodyData) Validate() error {
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

type GlobalHotelApplyRefundResponseBodyDataTotalPenaltyAmount struct {
	// The amount in the smallest currency unit.
	//
	// example:
	//
	// 10000
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The currency code in ISO 4217 format.
	//
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// TraceId
	//
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelApplyRefundResponseBodyDataTotalPenaltyAmount) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelApplyRefundResponseBodyDataTotalPenaltyAmount) GoString() string {
	return s.String()
}

func (s *GlobalHotelApplyRefundResponseBodyDataTotalPenaltyAmount) GetAmount() *string {
	return s.Amount
}

func (s *GlobalHotelApplyRefundResponseBodyDataTotalPenaltyAmount) GetCurrency() *string {
	return s.Currency
}

func (s *GlobalHotelApplyRefundResponseBodyDataTotalPenaltyAmount) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelApplyRefundResponseBodyDataTotalPenaltyAmount) SetAmount(v string) *GlobalHotelApplyRefundResponseBodyDataTotalPenaltyAmount {
	s.Amount = &v
	return s
}

func (s *GlobalHotelApplyRefundResponseBodyDataTotalPenaltyAmount) SetCurrency(v string) *GlobalHotelApplyRefundResponseBodyDataTotalPenaltyAmount {
	s.Currency = &v
	return s
}

func (s *GlobalHotelApplyRefundResponseBodyDataTotalPenaltyAmount) SetTracerId(v string) *GlobalHotelApplyRefundResponseBodyDataTotalPenaltyAmount {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelApplyRefundResponseBodyDataTotalPenaltyAmount) Validate() error {
	return dara.Validate(s)
}

type GlobalHotelApplyRefundResponseBodyDataTotalRefundAmount struct {
	// The amount in the smallest currency unit.
	//
	// example:
	//
	// 10000
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The currency code in ISO 4217 format.
	//
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// TraceId
	//
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelApplyRefundResponseBodyDataTotalRefundAmount) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelApplyRefundResponseBodyDataTotalRefundAmount) GoString() string {
	return s.String()
}

func (s *GlobalHotelApplyRefundResponseBodyDataTotalRefundAmount) GetAmount() *string {
	return s.Amount
}

func (s *GlobalHotelApplyRefundResponseBodyDataTotalRefundAmount) GetCurrency() *string {
	return s.Currency
}

func (s *GlobalHotelApplyRefundResponseBodyDataTotalRefundAmount) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelApplyRefundResponseBodyDataTotalRefundAmount) SetAmount(v string) *GlobalHotelApplyRefundResponseBodyDataTotalRefundAmount {
	s.Amount = &v
	return s
}

func (s *GlobalHotelApplyRefundResponseBodyDataTotalRefundAmount) SetCurrency(v string) *GlobalHotelApplyRefundResponseBodyDataTotalRefundAmount {
	s.Currency = &v
	return s
}

func (s *GlobalHotelApplyRefundResponseBodyDataTotalRefundAmount) SetTracerId(v string) *GlobalHotelApplyRefundResponseBodyDataTotalRefundAmount {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelApplyRefundResponseBodyDataTotalRefundAmount) Validate() error {
	return dara.Validate(s)
}
