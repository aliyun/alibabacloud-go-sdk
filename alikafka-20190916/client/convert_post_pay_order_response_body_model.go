// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertPostPayOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ConvertPostPayOrderResponseBody
	GetCode() *int32
	SetMessage(v string) *ConvertPostPayOrderResponseBody
	GetMessage() *string
	SetOrderId(v string) *ConvertPostPayOrderResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ConvertPostPayOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ConvertPostPayOrderResponseBody
	GetSuccess() *bool
}

type ConvertPostPayOrderResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 20497346575****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 06084011-E093-46F3-A51F-4B19A8AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConvertPostPayOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConvertPostPayOrderResponseBody) GoString() string {
	return s.String()
}

func (s *ConvertPostPayOrderResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ConvertPostPayOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ConvertPostPayOrderResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ConvertPostPayOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConvertPostPayOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ConvertPostPayOrderResponseBody) SetCode(v int32) *ConvertPostPayOrderResponseBody {
	s.Code = &v
	return s
}

func (s *ConvertPostPayOrderResponseBody) SetMessage(v string) *ConvertPostPayOrderResponseBody {
	s.Message = &v
	return s
}

func (s *ConvertPostPayOrderResponseBody) SetOrderId(v string) *ConvertPostPayOrderResponseBody {
	s.OrderId = &v
	return s
}

func (s *ConvertPostPayOrderResponseBody) SetRequestId(v string) *ConvertPostPayOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConvertPostPayOrderResponseBody) SetSuccess(v bool) *ConvertPostPayOrderResponseBody {
	s.Success = &v
	return s
}

func (s *ConvertPostPayOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
