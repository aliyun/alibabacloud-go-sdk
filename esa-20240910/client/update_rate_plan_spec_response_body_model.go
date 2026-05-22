// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRatePlanSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdateRatePlanSpecResponseBody
	GetInstanceId() *string
	SetOrderId(v string) *UpdateRatePlanSpecResponseBody
	GetOrderId() *string
	SetRequestId(v string) *UpdateRatePlanSpecResponseBody
	GetRequestId() *string
}

type UpdateRatePlanSpecResponseBody struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OrderId    *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRatePlanSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRatePlanSpecResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRatePlanSpecResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateRatePlanSpecResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *UpdateRatePlanSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRatePlanSpecResponseBody) SetInstanceId(v string) *UpdateRatePlanSpecResponseBody {
	s.InstanceId = &v
	return s
}

func (s *UpdateRatePlanSpecResponseBody) SetOrderId(v string) *UpdateRatePlanSpecResponseBody {
	s.OrderId = &v
	return s
}

func (s *UpdateRatePlanSpecResponseBody) SetRequestId(v string) *UpdateRatePlanSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRatePlanSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
