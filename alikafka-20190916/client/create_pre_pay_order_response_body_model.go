// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrePayOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreatePrePayOrderResponseBody
	GetCode() *int32
	SetMessage(v string) *CreatePrePayOrderResponseBody
	GetMessage() *string
	SetOrderId(v string) *CreatePrePayOrderResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreatePrePayOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreatePrePayOrderResponseBody
	GetSuccess() *bool
}

type CreatePrePayOrderResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
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

func (s CreatePrePayOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePrePayOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePrePayOrderResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreatePrePayOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePrePayOrderResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreatePrePayOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePrePayOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreatePrePayOrderResponseBody) SetCode(v int32) *CreatePrePayOrderResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePrePayOrderResponseBody) SetMessage(v string) *CreatePrePayOrderResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePrePayOrderResponseBody) SetOrderId(v string) *CreatePrePayOrderResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreatePrePayOrderResponseBody) SetRequestId(v string) *CreatePrePayOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePrePayOrderResponseBody) SetSuccess(v bool) *CreatePrePayOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CreatePrePayOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
