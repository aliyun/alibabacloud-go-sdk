// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TicketingResponseBody
	GetRequestId() *string
	SetData(v *TicketingResponseBodyData) *TicketingResponseBody
	GetData() *TicketingResponseBodyData
	SetErrorCode(v string) *TicketingResponseBody
	GetErrorCode() *string
	SetErrorData(v interface{}) *TicketingResponseBody
	GetErrorData() interface{}
	SetErrorMsg(v string) *TicketingResponseBody
	GetErrorMsg() *string
	SetStatus(v int32) *TicketingResponseBody
	GetStatus() *int32
	SetSuccess(v bool) *TicketingResponseBody
	GetSuccess() *bool
}

type TicketingResponseBody struct {
	// request ID
	//
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// data
	Data *TicketingResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s TicketingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TicketingResponseBody) GoString() string {
	return s.String()
}

func (s *TicketingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TicketingResponseBody) GetData() *TicketingResponseBodyData {
	return s.Data
}

func (s *TicketingResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *TicketingResponseBody) GetErrorData() interface{} {
	return s.ErrorData
}

func (s *TicketingResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *TicketingResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *TicketingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TicketingResponseBody) SetRequestId(v string) *TicketingResponseBody {
	s.RequestId = &v
	return s
}

func (s *TicketingResponseBody) SetData(v *TicketingResponseBodyData) *TicketingResponseBody {
	s.Data = v
	return s
}

func (s *TicketingResponseBody) SetErrorCode(v string) *TicketingResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TicketingResponseBody) SetErrorData(v interface{}) *TicketingResponseBody {
	s.ErrorData = v
	return s
}

func (s *TicketingResponseBody) SetErrorMsg(v string) *TicketingResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *TicketingResponseBody) SetStatus(v int32) *TicketingResponseBody {
	s.Status = &v
	return s
}

func (s *TicketingResponseBody) SetSuccess(v bool) *TicketingResponseBody {
	s.Success = &v
	return s
}

func (s *TicketingResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TicketingResponseBodyData struct {
	// order number
	//
	// example:
	//
	// 496***2617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// example:
	//
	// 1757404878000
	PayTime *int64 `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// transaction serial number
	//
	// example:
	//
	// hkduendkd-2023-dj0
	TransactionNo *string `json:"transaction_no,omitempty" xml:"transaction_no,omitempty"`
}

func (s TicketingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TicketingResponseBodyData) GoString() string {
	return s.String()
}

func (s *TicketingResponseBodyData) GetOrderNum() *int64 {
	return s.OrderNum
}

func (s *TicketingResponseBodyData) GetPayTime() *int64 {
	return s.PayTime
}

func (s *TicketingResponseBodyData) GetTransactionNo() *string {
	return s.TransactionNo
}

func (s *TicketingResponseBodyData) SetOrderNum(v int64) *TicketingResponseBodyData {
	s.OrderNum = &v
	return s
}

func (s *TicketingResponseBodyData) SetPayTime(v int64) *TicketingResponseBodyData {
	s.PayTime = &v
	return s
}

func (s *TicketingResponseBodyData) SetTransactionNo(v string) *TicketingResponseBodyData {
	s.TransactionNo = &v
	return s
}

func (s *TicketingResponseBodyData) Validate() error {
	return dara.Validate(s)
}
