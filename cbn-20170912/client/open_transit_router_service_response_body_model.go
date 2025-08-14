// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenTransitRouterServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *OpenTransitRouterServiceResponseBody
	GetOrderId() *string
	SetRequestId(v string) *OpenTransitRouterServiceResponseBody
	GetRequestId() *string
}

type OpenTransitRouterServiceResponseBody struct {
	// The ID of the order.
	//
	// example:
	//
	// 21370700730****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 866DEBE1-6411-51EC-80D8-975349B9FB4A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenTransitRouterServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenTransitRouterServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenTransitRouterServiceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *OpenTransitRouterServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenTransitRouterServiceResponseBody) SetOrderId(v string) *OpenTransitRouterServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenTransitRouterServiceResponseBody) SetRequestId(v string) *OpenTransitRouterServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenTransitRouterServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
