// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolarDBXInstanceNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *UpdatePolarDBXInstanceNodeResponseBody
	GetOrderId() *string
	SetRequestId(v string) *UpdatePolarDBXInstanceNodeResponseBody
	GetRequestId() *string
}

type UpdatePolarDBXInstanceNodeResponseBody struct {
	// example:
	//
	// 20211103105558
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePolarDBXInstanceNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarDBXInstanceNodeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePolarDBXInstanceNodeResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *UpdatePolarDBXInstanceNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePolarDBXInstanceNodeResponseBody) SetOrderId(v string) *UpdatePolarDBXInstanceNodeResponseBody {
	s.OrderId = &v
	return s
}

func (s *UpdatePolarDBXInstanceNodeResponseBody) SetRequestId(v string) *UpdatePolarDBXInstanceNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePolarDBXInstanceNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
