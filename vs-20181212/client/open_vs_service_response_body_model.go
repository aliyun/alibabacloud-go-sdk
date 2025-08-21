// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenVsServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *OpenVsServiceResponseBody
	GetOrderId() *string
	SetRequestId(v string) *OpenVsServiceResponseBody
	GetRequestId() *string
}

type OpenVsServiceResponseBody struct {
	// example:
	//
	// 150275784
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenVsServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenVsServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenVsServiceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *OpenVsServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenVsServiceResponseBody) SetOrderId(v string) *OpenVsServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenVsServiceResponseBody) SetRequestId(v string) *OpenVsServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenVsServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
