// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenHbrServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *OpenHbrServiceResponseBody
	GetOrderId() *string
	SetRequestId(v string) *OpenHbrServiceResponseBody
	GetRequestId() *string
}

type OpenHbrServiceResponseBody struct {
	// The ID of the order.
	//
	// example:
	//
	// 215463686160696
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F4A1D5F4-5055-549A-8B25-6DD23311E299
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenHbrServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenHbrServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenHbrServiceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *OpenHbrServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenHbrServiceResponseBody) SetOrderId(v string) *OpenHbrServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenHbrServiceResponseBody) SetRequestId(v string) *OpenHbrServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenHbrServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
