// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenPrivateLinkServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *OpenPrivateLinkServiceResponseBody
	GetOrderId() *string
	SetRequestId(v string) *OpenPrivateLinkServiceResponseBody
	GetRequestId() *string
}

type OpenPrivateLinkServiceResponseBody struct {
	// The order ID.
	//
	// example:
	//
	// 3245****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 427688B8-ADFB-4C4E-9D45-EF5C1FD6E23D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenPrivateLinkServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenPrivateLinkServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenPrivateLinkServiceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *OpenPrivateLinkServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenPrivateLinkServiceResponseBody) SetOrderId(v string) *OpenPrivateLinkServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenPrivateLinkServiceResponseBody) SetRequestId(v string) *OpenPrivateLinkServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenPrivateLinkServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
