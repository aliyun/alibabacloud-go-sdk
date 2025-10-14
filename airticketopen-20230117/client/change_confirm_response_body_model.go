// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeConfirmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ChangeConfirmResponseBody
	GetRequestId() *string
	SetData(v *ChangeConfirmResponseBodyData) *ChangeConfirmResponseBody
	GetData() *ChangeConfirmResponseBodyData
	SetErrorCode(v string) *ChangeConfirmResponseBody
	GetErrorCode() *string
	SetErrorData(v interface{}) *ChangeConfirmResponseBody
	GetErrorData() interface{}
	SetErrorMsg(v string) *ChangeConfirmResponseBody
	GetErrorMsg() *string
	SetStatus(v int32) *ChangeConfirmResponseBody
	GetStatus() *int32
	SetSuccess(v bool) *ChangeConfirmResponseBody
	GetSuccess() *bool
}

type ChangeConfirmResponseBody struct {
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ChangeConfirmResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ChangeConfirmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeConfirmResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeConfirmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeConfirmResponseBody) GetData() *ChangeConfirmResponseBodyData {
	return s.Data
}

func (s *ChangeConfirmResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ChangeConfirmResponseBody) GetErrorData() interface{} {
	return s.ErrorData
}

func (s *ChangeConfirmResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ChangeConfirmResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *ChangeConfirmResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChangeConfirmResponseBody) SetRequestId(v string) *ChangeConfirmResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeConfirmResponseBody) SetData(v *ChangeConfirmResponseBodyData) *ChangeConfirmResponseBody {
	s.Data = v
	return s
}

func (s *ChangeConfirmResponseBody) SetErrorCode(v string) *ChangeConfirmResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ChangeConfirmResponseBody) SetErrorData(v interface{}) *ChangeConfirmResponseBody {
	s.ErrorData = v
	return s
}

func (s *ChangeConfirmResponseBody) SetErrorMsg(v string) *ChangeConfirmResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ChangeConfirmResponseBody) SetStatus(v int32) *ChangeConfirmResponseBody {
	s.Status = &v
	return s
}

func (s *ChangeConfirmResponseBody) SetSuccess(v bool) *ChangeConfirmResponseBody {
	s.Success = &v
	return s
}

func (s *ChangeConfirmResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChangeConfirmResponseBodyData struct {
	// example:
	//
	// 30
	PayAmount *float64 `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
	// example:
	//
	// 1756797933000
	PayTime *int64 `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// example:
	//
	// hkduendkd-2023-dj0
	TransactionNo *string `json:"transaction_no,omitempty" xml:"transaction_no,omitempty"`
}

func (s ChangeConfirmResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ChangeConfirmResponseBodyData) GoString() string {
	return s.String()
}

func (s *ChangeConfirmResponseBodyData) GetPayAmount() *float64 {
	return s.PayAmount
}

func (s *ChangeConfirmResponseBodyData) GetPayTime() *int64 {
	return s.PayTime
}

func (s *ChangeConfirmResponseBodyData) GetTransactionNo() *string {
	return s.TransactionNo
}

func (s *ChangeConfirmResponseBodyData) SetPayAmount(v float64) *ChangeConfirmResponseBodyData {
	s.PayAmount = &v
	return s
}

func (s *ChangeConfirmResponseBodyData) SetPayTime(v int64) *ChangeConfirmResponseBodyData {
	s.PayTime = &v
	return s
}

func (s *ChangeConfirmResponseBodyData) SetTransactionNo(v string) *ChangeConfirmResponseBodyData {
	s.TransactionNo = &v
	return s
}

func (s *ChangeConfirmResponseBodyData) Validate() error {
	return dara.Validate(s)
}
