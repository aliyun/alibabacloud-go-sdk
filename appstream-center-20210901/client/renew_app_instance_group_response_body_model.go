// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewAppInstanceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RenewAppInstanceGroupResponseBody
	GetCode() *string
	SetMessage(v string) *RenewAppInstanceGroupResponseBody
	GetMessage() *string
	SetOrderId(v string) *RenewAppInstanceGroupResponseBody
	GetOrderId() *string
	SetRequestId(v string) *RenewAppInstanceGroupResponseBody
	GetRequestId() *string
}

type RenewAppInstanceGroupResponseBody struct {
	// The error code.
	//
	// example:
	//
	// InvalidParameter.ProductType
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The parameter ProductType is invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 123456****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewAppInstanceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewAppInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RenewAppInstanceGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *RenewAppInstanceGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RenewAppInstanceGroupResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *RenewAppInstanceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewAppInstanceGroupResponseBody) SetCode(v string) *RenewAppInstanceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *RenewAppInstanceGroupResponseBody) SetMessage(v string) *RenewAppInstanceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *RenewAppInstanceGroupResponseBody) SetOrderId(v string) *RenewAppInstanceGroupResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewAppInstanceGroupResponseBody) SetRequestId(v string) *RenewAppInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewAppInstanceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
