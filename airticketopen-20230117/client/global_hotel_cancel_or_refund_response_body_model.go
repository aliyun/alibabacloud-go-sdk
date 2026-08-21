// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelCancelOrRefundResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GlobalHotelCancelOrRefundResponseBodyData) *GlobalHotelCancelOrRefundResponseBody
	GetData() *GlobalHotelCancelOrRefundResponseBodyData
	SetErrorCode(v string) *GlobalHotelCancelOrRefundResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GlobalHotelCancelOrRefundResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *GlobalHotelCancelOrRefundResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GlobalHotelCancelOrRefundResponseBody
	GetSuccess() *bool
	SetTracerId(v string) *GlobalHotelCancelOrRefundResponseBody
	GetTracerId() *string
}

type GlobalHotelCancelOrRefundResponseBody struct {
	// The business data.
	Data *GlobalHotelCancelOrRefundResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// Indicates whether the operation is successful.
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

func (s GlobalHotelCancelOrRefundResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelCancelOrRefundResponseBody) GoString() string {
	return s.String()
}

func (s *GlobalHotelCancelOrRefundResponseBody) GetData() *GlobalHotelCancelOrRefundResponseBodyData {
	return s.Data
}

func (s *GlobalHotelCancelOrRefundResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GlobalHotelCancelOrRefundResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GlobalHotelCancelOrRefundResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GlobalHotelCancelOrRefundResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GlobalHotelCancelOrRefundResponseBody) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelCancelOrRefundResponseBody) SetData(v *GlobalHotelCancelOrRefundResponseBodyData) *GlobalHotelCancelOrRefundResponseBody {
	s.Data = v
	return s
}

func (s *GlobalHotelCancelOrRefundResponseBody) SetErrorCode(v string) *GlobalHotelCancelOrRefundResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GlobalHotelCancelOrRefundResponseBody) SetErrorMsg(v string) *GlobalHotelCancelOrRefundResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GlobalHotelCancelOrRefundResponseBody) SetRequestId(v string) *GlobalHotelCancelOrRefundResponseBody {
	s.RequestId = &v
	return s
}

func (s *GlobalHotelCancelOrRefundResponseBody) SetSuccess(v bool) *GlobalHotelCancelOrRefundResponseBody {
	s.Success = &v
	return s
}

func (s *GlobalHotelCancelOrRefundResponseBody) SetTracerId(v string) *GlobalHotelCancelOrRefundResponseBody {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelCancelOrRefundResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GlobalHotelCancelOrRefundResponseBodyData struct {
	// The after-sales refund order number. This parameter is returned when an after-sales refund is processed and can be used for polling the refund status. This parameter is null when a cancellation is processed.
	//
	// example:
	//
	// RF202606290001
	RefundOrderNo *string `json:"RefundOrderNo,omitempty" xml:"RefundOrderNo,omitempty"`
	// The total penalty amount on the sales side. This parameter is returned when an after-sales refund is processed.
	TotalPenaltyAmount *GlobalHotelCancelOrRefundResponseBodyDataTotalPenaltyAmount `json:"TotalPenaltyAmount,omitempty" xml:"TotalPenaltyAmount,omitempty" type:"Struct"`
	// The total refund amount. This parameter is returned when an after-sales refund is processed. The value equals the total sales price minus the total penalty amount.
	TotalRefundAmount *GlobalHotelCancelOrRefundResponseBodyDataTotalRefundAmount `json:"TotalRefundAmount,omitempty" xml:"TotalRefundAmount,omitempty" type:"Struct"`
	// TraceId
	//
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelCancelOrRefundResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelCancelOrRefundResponseBodyData) GoString() string {
	return s.String()
}

func (s *GlobalHotelCancelOrRefundResponseBodyData) GetRefundOrderNo() *string {
	return s.RefundOrderNo
}

func (s *GlobalHotelCancelOrRefundResponseBodyData) GetTotalPenaltyAmount() *GlobalHotelCancelOrRefundResponseBodyDataTotalPenaltyAmount {
	return s.TotalPenaltyAmount
}

func (s *GlobalHotelCancelOrRefundResponseBodyData) GetTotalRefundAmount() *GlobalHotelCancelOrRefundResponseBodyDataTotalRefundAmount {
	return s.TotalRefundAmount
}

func (s *GlobalHotelCancelOrRefundResponseBodyData) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelCancelOrRefundResponseBodyData) SetRefundOrderNo(v string) *GlobalHotelCancelOrRefundResponseBodyData {
	s.RefundOrderNo = &v
	return s
}

func (s *GlobalHotelCancelOrRefundResponseBodyData) SetTotalPenaltyAmount(v *GlobalHotelCancelOrRefundResponseBodyDataTotalPenaltyAmount) *GlobalHotelCancelOrRefundResponseBodyData {
	s.TotalPenaltyAmount = v
	return s
}

func (s *GlobalHotelCancelOrRefundResponseBodyData) SetTotalRefundAmount(v *GlobalHotelCancelOrRefundResponseBodyDataTotalRefundAmount) *GlobalHotelCancelOrRefundResponseBodyData {
	s.TotalRefundAmount = v
	return s
}

func (s *GlobalHotelCancelOrRefundResponseBodyData) SetTracerId(v string) *GlobalHotelCancelOrRefundResponseBodyData {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelCancelOrRefundResponseBodyData) Validate() error {
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

type GlobalHotelCancelOrRefundResponseBodyDataTotalPenaltyAmount struct {
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

func (s GlobalHotelCancelOrRefundResponseBodyDataTotalPenaltyAmount) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelCancelOrRefundResponseBodyDataTotalPenaltyAmount) GoString() string {
	return s.String()
}

func (s *GlobalHotelCancelOrRefundResponseBodyDataTotalPenaltyAmount) GetAmount() *string {
	return s.Amount
}

func (s *GlobalHotelCancelOrRefundResponseBodyDataTotalPenaltyAmount) GetCurrency() *string {
	return s.Currency
}

func (s *GlobalHotelCancelOrRefundResponseBodyDataTotalPenaltyAmount) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelCancelOrRefundResponseBodyDataTotalPenaltyAmount) SetAmount(v string) *GlobalHotelCancelOrRefundResponseBodyDataTotalPenaltyAmount {
	s.Amount = &v
	return s
}

func (s *GlobalHotelCancelOrRefundResponseBodyDataTotalPenaltyAmount) SetCurrency(v string) *GlobalHotelCancelOrRefundResponseBodyDataTotalPenaltyAmount {
	s.Currency = &v
	return s
}

func (s *GlobalHotelCancelOrRefundResponseBodyDataTotalPenaltyAmount) SetTracerId(v string) *GlobalHotelCancelOrRefundResponseBodyDataTotalPenaltyAmount {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelCancelOrRefundResponseBodyDataTotalPenaltyAmount) Validate() error {
	return dara.Validate(s)
}

type GlobalHotelCancelOrRefundResponseBodyDataTotalRefundAmount struct {
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

func (s GlobalHotelCancelOrRefundResponseBodyDataTotalRefundAmount) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelCancelOrRefundResponseBodyDataTotalRefundAmount) GoString() string {
	return s.String()
}

func (s *GlobalHotelCancelOrRefundResponseBodyDataTotalRefundAmount) GetAmount() *string {
	return s.Amount
}

func (s *GlobalHotelCancelOrRefundResponseBodyDataTotalRefundAmount) GetCurrency() *string {
	return s.Currency
}

func (s *GlobalHotelCancelOrRefundResponseBodyDataTotalRefundAmount) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelCancelOrRefundResponseBodyDataTotalRefundAmount) SetAmount(v string) *GlobalHotelCancelOrRefundResponseBodyDataTotalRefundAmount {
	s.Amount = &v
	return s
}

func (s *GlobalHotelCancelOrRefundResponseBodyDataTotalRefundAmount) SetCurrency(v string) *GlobalHotelCancelOrRefundResponseBodyDataTotalRefundAmount {
	s.Currency = &v
	return s
}

func (s *GlobalHotelCancelOrRefundResponseBodyDataTotalRefundAmount) SetTracerId(v string) *GlobalHotelCancelOrRefundResponseBodyDataTotalRefundAmount {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelCancelOrRefundResponseBodyDataTotalRefundAmount) Validate() error {
	return dara.Validate(s)
}
