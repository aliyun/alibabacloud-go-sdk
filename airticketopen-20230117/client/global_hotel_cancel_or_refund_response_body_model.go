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
	Data *GlobalHotelCancelOrRefundResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// example:
	//
	// RF202606290001
	RefundOrderId      *string                                                      `json:"RefundOrderId,omitempty" xml:"RefundOrderId,omitempty"`
	TotalPenaltyAmount *GlobalHotelCancelOrRefundResponseBodyDataTotalPenaltyAmount `json:"TotalPenaltyAmount,omitempty" xml:"TotalPenaltyAmount,omitempty" type:"Struct"`
	TotalRefundAmount  *GlobalHotelCancelOrRefundResponseBodyDataTotalRefundAmount  `json:"TotalRefundAmount,omitempty" xml:"TotalRefundAmount,omitempty" type:"Struct"`
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

func (s *GlobalHotelCancelOrRefundResponseBodyData) GetRefundOrderId() *string {
	return s.RefundOrderId
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

func (s *GlobalHotelCancelOrRefundResponseBodyData) SetRefundOrderId(v string) *GlobalHotelCancelOrRefundResponseBodyData {
	s.RefundOrderId = &v
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
