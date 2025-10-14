// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketingCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TicketingCheckResponseBody
	GetRequestId() *string
	SetData(v *TicketingCheckResponseBodyData) *TicketingCheckResponseBody
	GetData() *TicketingCheckResponseBodyData
	SetErrorCode(v string) *TicketingCheckResponseBody
	GetErrorCode() *string
	SetErrorData(v interface{}) *TicketingCheckResponseBody
	GetErrorData() interface{}
	SetErrorMsg(v string) *TicketingCheckResponseBody
	GetErrorMsg() *string
	SetStatus(v int32) *TicketingCheckResponseBody
	GetStatus() *int32
	SetSuccess(v bool) *TicketingCheckResponseBody
	GetSuccess() *bool
}

type TicketingCheckResponseBody struct {
	// request ID
	//
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// data
	Data *TicketingCheckResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s TicketingCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TicketingCheckResponseBody) GoString() string {
	return s.String()
}

func (s *TicketingCheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TicketingCheckResponseBody) GetData() *TicketingCheckResponseBodyData {
	return s.Data
}

func (s *TicketingCheckResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *TicketingCheckResponseBody) GetErrorData() interface{} {
	return s.ErrorData
}

func (s *TicketingCheckResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *TicketingCheckResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *TicketingCheckResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TicketingCheckResponseBody) SetRequestId(v string) *TicketingCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *TicketingCheckResponseBody) SetData(v *TicketingCheckResponseBodyData) *TicketingCheckResponseBody {
	s.Data = v
	return s
}

func (s *TicketingCheckResponseBody) SetErrorCode(v string) *TicketingCheckResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TicketingCheckResponseBody) SetErrorData(v interface{}) *TicketingCheckResponseBody {
	s.ErrorData = v
	return s
}

func (s *TicketingCheckResponseBody) SetErrorMsg(v string) *TicketingCheckResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *TicketingCheckResponseBody) SetStatus(v int32) *TicketingCheckResponseBody {
	s.Status = &v
	return s
}

func (s *TicketingCheckResponseBody) SetSuccess(v bool) *TicketingCheckResponseBody {
	s.Success = &v
	return s
}

func (s *TicketingCheckResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TicketingCheckResponseBodyData struct {
	// order number
	//
	// example:
	//
	// 496***2617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
}

func (s TicketingCheckResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TicketingCheckResponseBodyData) GoString() string {
	return s.String()
}

func (s *TicketingCheckResponseBodyData) GetOrderNum() *int64 {
	return s.OrderNum
}

func (s *TicketingCheckResponseBodyData) SetOrderNum(v int64) *TicketingCheckResponseBodyData {
	s.OrderNum = &v
	return s
}

func (s *TicketingCheckResponseBodyData) Validate() error {
	return dara.Validate(s)
}
