// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenAlimtServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *OpenAlimtServiceResponseBody
	GetOrderId() *string
	SetRequestId(v string) *OpenAlimtServiceResponseBody
	GetRequestId() *string
}

type OpenAlimtServiceResponseBody struct {
	// example:
	//
	// 123456
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// xxxx-xxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenAlimtServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenAlimtServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenAlimtServiceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *OpenAlimtServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenAlimtServiceResponseBody) SetOrderId(v string) *OpenAlimtServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenAlimtServiceResponseBody) SetRequestId(v string) *OpenAlimtServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenAlimtServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
