// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeCancelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ChangeCancelResponseBody
	GetRequestId() *string
	SetData(v interface{}) *ChangeCancelResponseBody
	GetData() interface{}
	SetErrorCode(v string) *ChangeCancelResponseBody
	GetErrorCode() *string
	SetErrorData(v interface{}) *ChangeCancelResponseBody
	GetErrorData() interface{}
	SetErrorMsg(v string) *ChangeCancelResponseBody
	GetErrorMsg() *string
	SetStatus(v int32) *ChangeCancelResponseBody
	GetStatus() *int32
	SetSuccess(v bool) *ChangeCancelResponseBody
	GetSuccess() *bool
}

type ChangeCancelResponseBody struct {
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// null
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
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

func (s ChangeCancelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeCancelResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeCancelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeCancelResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ChangeCancelResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ChangeCancelResponseBody) GetErrorData() interface{} {
	return s.ErrorData
}

func (s *ChangeCancelResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ChangeCancelResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *ChangeCancelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChangeCancelResponseBody) SetRequestId(v string) *ChangeCancelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeCancelResponseBody) SetData(v interface{}) *ChangeCancelResponseBody {
	s.Data = v
	return s
}

func (s *ChangeCancelResponseBody) SetErrorCode(v string) *ChangeCancelResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ChangeCancelResponseBody) SetErrorData(v interface{}) *ChangeCancelResponseBody {
	s.ErrorData = v
	return s
}

func (s *ChangeCancelResponseBody) SetErrorMsg(v string) *ChangeCancelResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ChangeCancelResponseBody) SetStatus(v int32) *ChangeCancelResponseBody {
	s.Status = &v
	return s
}

func (s *ChangeCancelResponseBody) SetSuccess(v bool) *ChangeCancelResponseBody {
	s.Success = &v
	return s
}

func (s *ChangeCancelResponseBody) Validate() error {
	return dara.Validate(s)
}
