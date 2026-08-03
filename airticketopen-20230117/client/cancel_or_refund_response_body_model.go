// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelOrRefundResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CancelOrRefundResponseBodyData) *CancelOrRefundResponseBody
	GetData() *CancelOrRefundResponseBodyData
	SetErrorCode(v string) *CancelOrRefundResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *CancelOrRefundResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *CancelOrRefundResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CancelOrRefundResponseBody
	GetSuccess() *bool
	SetTracerId(v string) *CancelOrRefundResponseBody
	GetTracerId() *string
}

type CancelOrRefundResponseBody struct {
	Data *CancelOrRefundResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s CancelOrRefundResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelOrRefundResponseBody) GoString() string {
	return s.String()
}

func (s *CancelOrRefundResponseBody) GetData() *CancelOrRefundResponseBodyData {
	return s.Data
}

func (s *CancelOrRefundResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CancelOrRefundResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *CancelOrRefundResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelOrRefundResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelOrRefundResponseBody) GetTracerId() *string {
	return s.TracerId
}

func (s *CancelOrRefundResponseBody) SetData(v *CancelOrRefundResponseBodyData) *CancelOrRefundResponseBody {
	s.Data = v
	return s
}

func (s *CancelOrRefundResponseBody) SetErrorCode(v string) *CancelOrRefundResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CancelOrRefundResponseBody) SetErrorMsg(v string) *CancelOrRefundResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CancelOrRefundResponseBody) SetRequestId(v string) *CancelOrRefundResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelOrRefundResponseBody) SetSuccess(v bool) *CancelOrRefundResponseBody {
	s.Success = &v
	return s
}

func (s *CancelOrRefundResponseBody) SetTracerId(v string) *CancelOrRefundResponseBody {
	s.TracerId = &v
	return s
}

func (s *CancelOrRefundResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CancelOrRefundResponseBodyData struct {
	// example:
	//
	// RF202606290001
	RefundOrderId      *string                                           `json:"RefundOrderId,omitempty" xml:"RefundOrderId,omitempty"`
	TotalPenaltyAmount *CancelOrRefundResponseBodyDataTotalPenaltyAmount `json:"TotalPenaltyAmount,omitempty" xml:"TotalPenaltyAmount,omitempty" type:"Struct"`
	TotalRefundAmount  *CancelOrRefundResponseBodyDataTotalRefundAmount  `json:"TotalRefundAmount,omitempty" xml:"TotalRefundAmount,omitempty" type:"Struct"`
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s CancelOrRefundResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CancelOrRefundResponseBodyData) GoString() string {
	return s.String()
}

func (s *CancelOrRefundResponseBodyData) GetRefundOrderId() *string {
	return s.RefundOrderId
}

func (s *CancelOrRefundResponseBodyData) GetTotalPenaltyAmount() *CancelOrRefundResponseBodyDataTotalPenaltyAmount {
	return s.TotalPenaltyAmount
}

func (s *CancelOrRefundResponseBodyData) GetTotalRefundAmount() *CancelOrRefundResponseBodyDataTotalRefundAmount {
	return s.TotalRefundAmount
}

func (s *CancelOrRefundResponseBodyData) GetTracerId() *string {
	return s.TracerId
}

func (s *CancelOrRefundResponseBodyData) SetRefundOrderId(v string) *CancelOrRefundResponseBodyData {
	s.RefundOrderId = &v
	return s
}

func (s *CancelOrRefundResponseBodyData) SetTotalPenaltyAmount(v *CancelOrRefundResponseBodyDataTotalPenaltyAmount) *CancelOrRefundResponseBodyData {
	s.TotalPenaltyAmount = v
	return s
}

func (s *CancelOrRefundResponseBodyData) SetTotalRefundAmount(v *CancelOrRefundResponseBodyDataTotalRefundAmount) *CancelOrRefundResponseBodyData {
	s.TotalRefundAmount = v
	return s
}

func (s *CancelOrRefundResponseBodyData) SetTracerId(v string) *CancelOrRefundResponseBodyData {
	s.TracerId = &v
	return s
}

func (s *CancelOrRefundResponseBodyData) Validate() error {
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

type CancelOrRefundResponseBodyDataTotalPenaltyAmount struct {
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

func (s CancelOrRefundResponseBodyDataTotalPenaltyAmount) String() string {
	return dara.Prettify(s)
}

func (s CancelOrRefundResponseBodyDataTotalPenaltyAmount) GoString() string {
	return s.String()
}

func (s *CancelOrRefundResponseBodyDataTotalPenaltyAmount) GetAmount() *string {
	return s.Amount
}

func (s *CancelOrRefundResponseBodyDataTotalPenaltyAmount) GetCurrency() *string {
	return s.Currency
}

func (s *CancelOrRefundResponseBodyDataTotalPenaltyAmount) GetTracerId() *string {
	return s.TracerId
}

func (s *CancelOrRefundResponseBodyDataTotalPenaltyAmount) SetAmount(v string) *CancelOrRefundResponseBodyDataTotalPenaltyAmount {
	s.Amount = &v
	return s
}

func (s *CancelOrRefundResponseBodyDataTotalPenaltyAmount) SetCurrency(v string) *CancelOrRefundResponseBodyDataTotalPenaltyAmount {
	s.Currency = &v
	return s
}

func (s *CancelOrRefundResponseBodyDataTotalPenaltyAmount) SetTracerId(v string) *CancelOrRefundResponseBodyDataTotalPenaltyAmount {
	s.TracerId = &v
	return s
}

func (s *CancelOrRefundResponseBodyDataTotalPenaltyAmount) Validate() error {
	return dara.Validate(s)
}

type CancelOrRefundResponseBodyDataTotalRefundAmount struct {
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

func (s CancelOrRefundResponseBodyDataTotalRefundAmount) String() string {
	return dara.Prettify(s)
}

func (s CancelOrRefundResponseBodyDataTotalRefundAmount) GoString() string {
	return s.String()
}

func (s *CancelOrRefundResponseBodyDataTotalRefundAmount) GetAmount() *string {
	return s.Amount
}

func (s *CancelOrRefundResponseBodyDataTotalRefundAmount) GetCurrency() *string {
	return s.Currency
}

func (s *CancelOrRefundResponseBodyDataTotalRefundAmount) GetTracerId() *string {
	return s.TracerId
}

func (s *CancelOrRefundResponseBodyDataTotalRefundAmount) SetAmount(v string) *CancelOrRefundResponseBodyDataTotalRefundAmount {
	s.Amount = &v
	return s
}

func (s *CancelOrRefundResponseBodyDataTotalRefundAmount) SetCurrency(v string) *CancelOrRefundResponseBodyDataTotalRefundAmount {
	s.Currency = &v
	return s
}

func (s *CancelOrRefundResponseBodyDataTotalRefundAmount) SetTracerId(v string) *CancelOrRefundResponseBodyDataTotalRefundAmount {
	s.TracerId = &v
	return s
}

func (s *CancelOrRefundResponseBodyDataTotalRefundAmount) Validate() error {
	return dara.Validate(s)
}
