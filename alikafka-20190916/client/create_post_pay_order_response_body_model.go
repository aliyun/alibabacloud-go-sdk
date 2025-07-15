// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePostPayOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreatePostPayOrderResponseBody
	GetCode() *int32
	SetMessage(v string) *CreatePostPayOrderResponseBody
	GetMessage() *string
	SetOrderId(v string) *CreatePostPayOrderResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreatePostPayOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreatePostPayOrderResponseBody
	GetSuccess() *bool
}

type CreatePostPayOrderResponseBody struct {
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

func (s CreatePostPayOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePostPayOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePostPayOrderResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreatePostPayOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePostPayOrderResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreatePostPayOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePostPayOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreatePostPayOrderResponseBody) SetCode(v int32) *CreatePostPayOrderResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePostPayOrderResponseBody) SetMessage(v string) *CreatePostPayOrderResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePostPayOrderResponseBody) SetOrderId(v string) *CreatePostPayOrderResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreatePostPayOrderResponseBody) SetRequestId(v string) *CreatePostPayOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePostPayOrderResponseBody) SetSuccess(v bool) *CreatePostPayOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CreatePostPayOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
