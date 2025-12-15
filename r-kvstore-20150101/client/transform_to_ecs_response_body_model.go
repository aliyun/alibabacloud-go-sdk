// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransformToEcsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *TransformToEcsResponseBody
	GetOrderId() *string
	SetRequestId(v string) *TransformToEcsResponseBody
	GetRequestId() *string
}

type TransformToEcsResponseBody struct {
	// The order ID.
	//
	// example:
	//
	// 20905403119****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// DAABAD9B-890F-56C0-806C-6144946594AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TransformToEcsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TransformToEcsResponseBody) GoString() string {
	return s.String()
}

func (s *TransformToEcsResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *TransformToEcsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TransformToEcsResponseBody) SetOrderId(v string) *TransformToEcsResponseBody {
	s.OrderId = &v
	return s
}

func (s *TransformToEcsResponseBody) SetRequestId(v string) *TransformToEcsResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransformToEcsResponseBody) Validate() error {
	return dara.Validate(s)
}
