// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLoadBalancerInstanceSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *ModifyLoadBalancerInstanceSpecResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *ModifyLoadBalancerInstanceSpecResponseBody
	GetRequestId() *string
}

type ModifyLoadBalancerInstanceSpecResponseBody struct {
	// The order ID of the subscription CLB instance.
	//
	// example:
	//
	// 201429619788910
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLoadBalancerInstanceSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLoadBalancerInstanceSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLoadBalancerInstanceSpecResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ModifyLoadBalancerInstanceSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLoadBalancerInstanceSpecResponseBody) SetOrderId(v int64) *ModifyLoadBalancerInstanceSpecResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyLoadBalancerInstanceSpecResponseBody) SetRequestId(v string) *ModifyLoadBalancerInstanceSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLoadBalancerInstanceSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
