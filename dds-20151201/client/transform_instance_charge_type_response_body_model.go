// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransformInstanceChargeTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *TransformInstanceChargeTypeResponseBody
	GetOrderId() *string
	SetRequestId(v string) *TransformInstanceChargeTypeResponseBody
	GetRequestId() *string
}

type TransformInstanceChargeTypeResponseBody struct {
	// The order ID.
	//
	// example:
	//
	// 21084641369****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D8F1D721-6439-4257-A89C-F1E8E9C9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TransformInstanceChargeTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TransformInstanceChargeTypeResponseBody) GoString() string {
	return s.String()
}

func (s *TransformInstanceChargeTypeResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *TransformInstanceChargeTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TransformInstanceChargeTypeResponseBody) SetOrderId(v string) *TransformInstanceChargeTypeResponseBody {
	s.OrderId = &v
	return s
}

func (s *TransformInstanceChargeTypeResponseBody) SetRequestId(v string) *TransformInstanceChargeTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransformInstanceChargeTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
