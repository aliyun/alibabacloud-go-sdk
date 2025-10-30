// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewWuyingServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *RenewWuyingServerResponseBody
	GetOrderId() *string
	SetRequestId(v string) *RenewWuyingServerResponseBody
	GetRequestId() *string
}

type RenewWuyingServerResponseBody struct {
	// The order ID.
	//
	// example:
	//
	// 23977848****97
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewWuyingServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewWuyingServerResponseBody) GoString() string {
	return s.String()
}

func (s *RenewWuyingServerResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *RenewWuyingServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewWuyingServerResponseBody) SetOrderId(v string) *RenewWuyingServerResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewWuyingServerResponseBody) SetRequestId(v string) *RenewWuyingServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewWuyingServerResponseBody) Validate() error {
	return dara.Validate(s)
}
