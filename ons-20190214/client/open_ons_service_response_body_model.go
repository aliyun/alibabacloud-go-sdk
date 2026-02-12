// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenOnsServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *OpenOnsServiceResponseBody
	GetOrderId() *string
	SetRequestId(v string) *OpenOnsServiceResponseBody
	GetRequestId() *string
}

type OpenOnsServiceResponseBody struct {
	// The ID of the order.
	//
	// example:
	//
	// 2068689****0272
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 8C5B4603-8977-4513-AB60-9C3E2F88****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenOnsServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenOnsServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenOnsServiceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *OpenOnsServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenOnsServiceResponseBody) SetOrderId(v string) *OpenOnsServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenOnsServiceResponseBody) SetRequestId(v string) *OpenOnsServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenOnsServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
