// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelCancelOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GlobalHotelCancelOrderResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GlobalHotelCancelOrderResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *GlobalHotelCancelOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GlobalHotelCancelOrderResponseBody
	GetSuccess() *bool
	SetTracerId(v string) *GlobalHotelCancelOrderResponseBody
	GetTracerId() *string
}

type GlobalHotelCancelOrderResponseBody struct {
	// The error code.
	//
	// example:
	//
	// CreateOrderFailed
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Failed to create the order
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The unique request ID.
	//
	// example:
	//
	// 260E4F99-983D-1919-834C-5C42E98E5B2B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// TraceId
	//
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelCancelOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelCancelOrderResponseBody) GoString() string {
	return s.String()
}

func (s *GlobalHotelCancelOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GlobalHotelCancelOrderResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GlobalHotelCancelOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GlobalHotelCancelOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GlobalHotelCancelOrderResponseBody) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelCancelOrderResponseBody) SetErrorCode(v string) *GlobalHotelCancelOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GlobalHotelCancelOrderResponseBody) SetErrorMsg(v string) *GlobalHotelCancelOrderResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GlobalHotelCancelOrderResponseBody) SetRequestId(v string) *GlobalHotelCancelOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GlobalHotelCancelOrderResponseBody) SetSuccess(v bool) *GlobalHotelCancelOrderResponseBody {
	s.Success = &v
	return s
}

func (s *GlobalHotelCancelOrderResponseBody) SetTracerId(v string) *GlobalHotelCancelOrderResponseBody {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelCancelOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
