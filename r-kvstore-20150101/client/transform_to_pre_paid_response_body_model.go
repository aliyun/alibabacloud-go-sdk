// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransformToPrePaidResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *TransformToPrePaidResponseBody
	GetEndTime() *string
	SetOrderId(v string) *TransformToPrePaidResponseBody
	GetOrderId() *string
	SetRequestId(v string) *TransformToPrePaidResponseBody
	GetRequestId() *string
}

type TransformToPrePaidResponseBody struct {
	// The time when the instance expires after the billing method of the instance is changed from pay-as-you-go to subscription.
	//
	// example:
	//
	// 2019-01-18T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 111111111111111
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 426F1356-B6EF-4DAD-A1C3-DE53B9DAF586
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TransformToPrePaidResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TransformToPrePaidResponseBody) GoString() string {
	return s.String()
}

func (s *TransformToPrePaidResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *TransformToPrePaidResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *TransformToPrePaidResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TransformToPrePaidResponseBody) SetEndTime(v string) *TransformToPrePaidResponseBody {
	s.EndTime = &v
	return s
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
