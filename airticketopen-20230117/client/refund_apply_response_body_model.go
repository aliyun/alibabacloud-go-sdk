// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefundApplyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RefundApplyResponseBody
	GetRequestId() *string
	SetData(v *RefundApplyResponseBodyData) *RefundApplyResponseBody
	GetData() *RefundApplyResponseBodyData
	SetErrorCode(v string) *RefundApplyResponseBody
	GetErrorCode() *string
	SetErrorData(v interface{}) *RefundApplyResponseBody
	GetErrorData() interface{}
	SetErrorMsg(v string) *RefundApplyResponseBody
	GetErrorMsg() *string
	SetStatus(v int32) *RefundApplyResponseBody
	GetStatus() *int32
	SetSuccess(v bool) *RefundApplyResponseBody
	GetSuccess() *bool
}

type RefundApplyResponseBody struct {
	// request ID
	//
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// data
	Data *RefundApplyResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// error code
	//
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// error data
	//
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// error message
	//
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// http reqeust has been processed successfully，status code is 200
	//
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// true represents success, false represents failure
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RefundApplyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefundApplyResponseBody) GoString() string {
	return s.String()
}

func (s *RefundApplyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefundApplyResponseBody) GetData() *RefundApplyResponseBodyData {
	return s.Data
}

func (s *RefundApplyResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RefundApplyResponseBody) GetErrorData() interface{} {
	return s.ErrorData
}

func (s *RefundApplyResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *RefundApplyResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *RefundApplyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RefundApplyResponseBody) SetRequestId(v string) *RefundApplyResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefundApplyResponseBody) SetData(v *RefundApplyResponseBodyData) *RefundApplyResponseBody {
	s.Data = v
	return s
}

func (s *RefundApplyResponseBody) SetErrorCode(v string) *RefundApplyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RefundApplyResponseBody) SetErrorData(v interface{}) *RefundApplyResponseBody {
	s.ErrorData = v
	return s
}

func (s *RefundApplyResponseBody) SetErrorMsg(v string) *RefundApplyResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *RefundApplyResponseBody) SetStatus(v int32) *RefundApplyResponseBody {
	s.Status = &v
	return s
}

func (s *RefundApplyResponseBody) SetSuccess(v bool) *RefundApplyResponseBody {
	s.Success = &v
	return s
}

func (s *RefundApplyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RefundApplyResponseBodyData struct {
	// order number
	//
	// example:
	//
	// 4966***617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// refund results
	RefundResults []*RefundApplyResponseBodyDataRefundResults `json:"refund_results,omitempty" xml:"refund_results,omitempty" type:"Repeated"`
}

func (s RefundApplyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RefundApplyResponseBodyData) GoString() string {
	return s.String()
}

func (s *RefundApplyResponseBodyData) GetOrderNum() *int64 {
	return s.OrderNum
}

func (s *RefundApplyResponseBodyData) GetRefundResults() []*RefundApplyResponseBodyDataRefundResults {
	return s.RefundResults
}

func (s *RefundApplyResponseBodyData) SetOrderNum(v int64) *RefundApplyResponseBodyData {
	s.OrderNum = &v
	return s
}

func (s *RefundApplyResponseBodyData) SetRefundResults(v []*RefundApplyResponseBodyDataRefundResults) *RefundApplyResponseBodyData {
	s.RefundResults = v
	return s
}

func (s *RefundApplyResponseBodyData) Validate() error {
	if s.RefundResults != nil {
		for _, item := range s.RefundResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RefundApplyResponseBodyDataRefundResults struct {
	// reason for refund application failure
	//
	// example:
	//
	// desc reason
	FailReason *string `json:"fail_reason,omitempty" xml:"fail_reason,omitempty"`
	// refund order number
	//
	// example:
	//
	// 4966***617202
	RefundOrderNum *int64 `json:"refund_order_num,omitempty" xml:"refund_order_num,omitempty"`
	// passengers of current refund order
	RefundPassengers []*RefundApplyResponseBodyDataRefundResultsRefundPassengers `json:"refund_passengers,omitempty" xml:"refund_passengers,omitempty" type:"Repeated"`
	// refund order status
	//
	// 0: refund order created successfully
	//
	// 1: refund order creation failed
	//
	// example:
	//
	// 0
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s RefundApplyResponseBodyDataRefundResults) String() string {
	return dara.Prettify(s)
}

func (s RefundApplyResponseBodyDataRefundResults) GoString() string {
	return s.String()
}

func (s *RefundApplyResponseBodyDataRefundResults) GetFailReason() *string {
	return s.FailReason
}

func (s *RefundApplyResponseBodyDataRefundResults) GetRefundOrderNum() *int64 {
	return s.RefundOrderNum
}

func (s *RefundApplyResponseBodyDataRefundResults) GetRefundPassengers() []*RefundApplyResponseBodyDataRefundResultsRefundPassengers {
	return s.RefundPassengers
}

func (s *RefundApplyResponseBodyDataRefundResults) GetStatus() *int32 {
	return s.Status
}

func (s *RefundApplyResponseBodyDataRefundResults) SetFailReason(v string) *RefundApplyResponseBodyDataRefundResults {
	s.FailReason = &v
	return s
}

func (s *RefundApplyResponseBodyDataRefundResults) SetRefundOrderNum(v int64) *RefundApplyResponseBodyDataRefundResults {
	s.RefundOrderNum = &v
	return s
}

func (s *RefundApplyResponseBodyDataRefundResults) SetRefundPassengers(v []*RefundApplyResponseBodyDataRefundResultsRefundPassengers) *RefundApplyResponseBodyDataRefundResults {
	s.RefundPassengers = v
	return s
}

func (s *RefundApplyResponseBodyDataRefundResults) SetStatus(v int32) *RefundApplyResponseBodyDataRefundResults {
	s.Status = &v
	return s
}

func (s *RefundApplyResponseBodyDataRefundResults) Validate() error {
	if s.RefundPassengers != nil {
		for _, item := range s.RefundPassengers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RefundApplyResponseBodyDataRefundResultsRefundPassengers struct {
	// credential number
	//
	// example:
	//
	// 411***********4411
	Document *string `json:"document,omitempty" xml:"document,omitempty"`
	// first name
	//
	// example:
	//
	// SAN
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// last name
	//
	// example:
	//
	// ZHANG
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
}

func (s RefundApplyResponseBodyDataRefundResultsRefundPassengers) String() string {
	return dara.Prettify(s)
}

func (s RefundApplyResponseBodyDataRefundResultsRefundPassengers) GoString() string {
	return s.String()
}

func (s *RefundApplyResponseBodyDataRefundResultsRefundPassengers) GetDocument() *string {
	return s.Document
}

func (s *RefundApplyResponseBodyDataRefundResultsRefundPassengers) GetFirstName() *string {
	return s.FirstName
}

func (s *RefundApplyResponseBodyDataRefundResultsRefundPassengers) GetLastName() *string {
	return s.LastName
}

func (s *RefundApplyResponseBodyDataRefundResultsRefundPassengers) SetDocument(v string) *RefundApplyResponseBodyDataRefundResultsRefundPassengers {
	s.Document = &v
	return s
}

func (s *RefundApplyResponseBodyDataRefundResultsRefundPassengers) SetFirstName(v string) *RefundApplyResponseBodyDataRefundResultsRefundPassengers {
	s.FirstName = &v
	return s
}

func (s *RefundApplyResponseBodyDataRefundResultsRefundPassengers) SetLastName(v string) *RefundApplyResponseBodyDataRefundResultsRefundPassengers {
	s.LastName = &v
	return s
}

func (s *RefundApplyResponseBodyDataRefundResultsRefundPassengers) Validate() error {
	return dara.Validate(s)
}
