// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransformToPrePaidResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *TransformToPrePaidResponseBody
	GetOrderId() *string
	SetRequestId(v string) *TransformToPrePaidResponseBody
	GetRequestId() *string
}

type TransformToPrePaidResponseBody struct {
	// The ID of the order.
	//
	// example:
	//
	// 21022019252****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2F42BB4E-461F-5B55-A37C-53B1141C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TransformToPrePaidResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TransformToPrePaidResponseBody) GoString() string {
	return s.String()
}

func (s *TransformToPrePaidResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *TransformToPrePaidResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TransformToPrePaidResponseBody) SetOrderId(v string) *TransformToPrePaidResponseBody {
	s.OrderId = &v
	return s
}

func (s *TransformToPrePaidResponseBody) SetRequestId(v string) *TransformToPrePaidResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransformToPrePaidResponseBody) Validate() error {
	return dara.Validate(s)
}
