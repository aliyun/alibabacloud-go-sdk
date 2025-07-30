// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransformInstanceChargeTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *TransformInstanceChargeTypeResponseBody
	GetEndTime() *string
	SetOrderId(v string) *TransformInstanceChargeTypeResponseBody
	GetOrderId() *string
	SetRequestId(v string) *TransformInstanceChargeTypeResponseBody
	GetRequestId() *string
}

type TransformInstanceChargeTypeResponseBody struct {
	// The time when the instance expires.
	//
	// >  A value is returned for this parameter only if the instance was changed from pay-as-you-go to subscription.
	//
	// example:
	//
	// 2021-05-13T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 20881824000****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 82C791FB-8979-489E-853D-706D7743****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TransformInstanceChargeTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TransformInstanceChargeTypeResponseBody) GoString() string {
	return s.String()
}

func (s *TransformInstanceChargeTypeResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *TransformInstanceChargeTypeResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *TransformInstanceChargeTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TransformInstanceChargeTypeResponseBody) SetEndTime(v string) *TransformInstanceChargeTypeResponseBody {
	s.EndTime = &v
	return s
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
