// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenArmsServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *OpenArmsServiceResponseBody
	GetOrderId() *string
	SetRequestId(v string) *OpenArmsServiceResponseBody
	GetRequestId() *string
}

type OpenArmsServiceResponseBody struct {
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenArmsServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenArmsServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenArmsServiceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *OpenArmsServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenArmsServiceResponseBody) SetOrderId(v string) *OpenArmsServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenArmsServiceResponseBody) SetRequestId(v string) *OpenArmsServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenArmsServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
