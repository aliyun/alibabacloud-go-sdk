// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelResponseBody
	GetRequestId() *string
	SetData(v *CancelResponseBodyData) *CancelResponseBody
	GetData() *CancelResponseBodyData
	SetErrorCode(v string) *CancelResponseBody
	GetErrorCode() *string
	SetErrorData(v interface{}) *CancelResponseBody
	GetErrorData() interface{}
	SetErrorMsg(v string) *CancelResponseBody
	GetErrorMsg() *string
	SetStatus(v int32) *CancelResponseBody
	GetStatus() *int32
	SetSuccess(v bool) *CancelResponseBody
	GetSuccess() *bool
}

type CancelResponseBody struct {
	// request ID
	//
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// data
	Data *CancelResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s CancelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelResponseBody) GoString() string {
	return s.String()
}

func (s *CancelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelResponseBody) GetData() *CancelResponseBodyData {
	return s.Data
}

func (s *CancelResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CancelResponseBody) GetErrorData() interface{} {
	return s.ErrorData
}

func (s *CancelResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *CancelResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *CancelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelResponseBody) SetRequestId(v string) *CancelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelResponseBody) SetData(v *CancelResponseBodyData) *CancelResponseBody {
	s.Data = v
	return s
}

func (s *CancelResponseBody) SetErrorCode(v string) *CancelResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CancelResponseBody) SetErrorData(v interface{}) *CancelResponseBody {
	s.ErrorData = v
	return s
}

func (s *CancelResponseBody) SetErrorMsg(v string) *CancelResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CancelResponseBody) SetStatus(v int32) *CancelResponseBody {
	s.Status = &v
	return s
}

func (s *CancelResponseBody) SetSuccess(v bool) *CancelResponseBody {
	s.Success = &v
	return s
}

func (s *CancelResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CancelResponseBodyData struct {
	// order number
	//
	// example:
	//
	// 496***2617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
}

func (s CancelResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CancelResponseBodyData) GoString() string {
	return s.String()
}

func (s *CancelResponseBodyData) GetOrderNum() *int64 {
	return s.OrderNum
}

func (s *CancelResponseBodyData) SetOrderNum(v int64) *CancelResponseBodyData {
	s.OrderNum = &v
	return s
}

func (s *CancelResponseBodyData) Validate() error {
	return dara.Validate(s)
}
