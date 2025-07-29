// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenAckServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *OpenAckServiceResponseBody
	GetOrderId() *string
	SetRequestId(v string) *OpenAckServiceResponseBody
	GetRequestId() *string
}

type OpenAckServiceResponseBody struct {
	// The ID of the order.
	//
	// example:
	//
	// 2067*******0374
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

func (s OpenAckServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenAckServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenAckServiceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *OpenAckServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenAckServiceResponseBody) SetOrderId(v string) *OpenAckServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenAckServiceResponseBody) SetRequestId(v string) *OpenAckServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenAckServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
