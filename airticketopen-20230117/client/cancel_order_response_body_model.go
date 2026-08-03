// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CancelOrderResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *CancelOrderResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *CancelOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CancelOrderResponseBody
	GetSuccess() *bool
	SetTracerId(v string) *CancelOrderResponseBody
	GetTracerId() *string
}

type CancelOrderResponseBody struct {
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

func (s CancelOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CancelOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CancelOrderResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *CancelOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelOrderResponseBody) GetTracerId() *string {
	return s.TracerId
}

func (s *CancelOrderResponseBody) SetErrorCode(v string) *CancelOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CancelOrderResponseBody) SetErrorMsg(v string) *CancelOrderResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CancelOrderResponseBody) SetRequestId(v string) *CancelOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelOrderResponseBody) SetSuccess(v bool) *CancelOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CancelOrderResponseBody) SetTracerId(v string) *CancelOrderResponseBody {
	s.TracerId = &v
	return s
}

func (s *CancelOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
