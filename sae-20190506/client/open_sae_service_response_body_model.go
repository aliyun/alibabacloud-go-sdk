// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenSaeServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *OpenSaeServiceResponseBody
	GetOrderId() *string
	SetRequestId(v string) *OpenSaeServiceResponseBody
	GetRequestId() *string
}

type OpenSaeServiceResponseBody struct {
	// PushEvent
	//
	// example:
	//
	// 20485646575****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// enableWAF
	//
	// example:
	//
	// 559B4247-C41C-4D9E-B866-B55D378B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenSaeServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenSaeServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenSaeServiceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *OpenSaeServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenSaeServiceResponseBody) SetOrderId(v string) *OpenSaeServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenSaeServiceResponseBody) SetRequestId(v string) *OpenSaeServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenSaeServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
