// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenEmasServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *OpenEmasServiceResponseBody
	GetOrderId() *string
	SetRequestId(v string) *OpenEmasServiceResponseBody
	GetRequestId() *string
}

type OpenEmasServiceResponseBody struct {
	// example:
	//
	// 20671870151****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 126D4DDD-05A5-49B1-B18C-39C4A929BFB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenEmasServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenEmasServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenEmasServiceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *OpenEmasServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenEmasServiceResponseBody) SetOrderId(v string) *OpenEmasServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenEmasServiceResponseBody) SetRequestId(v string) *OpenEmasServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenEmasServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
