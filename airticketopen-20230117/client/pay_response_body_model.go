// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *PayResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *PayResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *PayResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PayResponseBody
	GetSuccess() *bool
	SetTracerId(v string) *PayResponseBody
	GetTracerId() *string
}

type PayResponseBody struct {
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

func (s PayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PayResponseBody) GoString() string {
	return s.String()
}

func (s *PayResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *PayResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *PayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PayResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PayResponseBody) GetTracerId() *string {
	return s.TracerId
}

func (s *PayResponseBody) SetErrorCode(v string) *PayResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *PayResponseBody) SetErrorMsg(v string) *PayResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *PayResponseBody) SetRequestId(v string) *PayResponseBody {
	s.RequestId = &v
	return s
}

func (s *PayResponseBody) SetSuccess(v bool) *PayResponseBody {
	s.Success = &v
	return s
}

func (s *PayResponseBody) SetTracerId(v string) *PayResponseBody {
	s.TracerId = &v
	return s
}

func (s *PayResponseBody) Validate() error {
	return dara.Validate(s)
}
