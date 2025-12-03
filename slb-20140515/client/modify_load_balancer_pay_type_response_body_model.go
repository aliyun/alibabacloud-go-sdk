// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLoadBalancerPayTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *ModifyLoadBalancerPayTypeResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *ModifyLoadBalancerPayTypeResponseBody
	GetRequestId() *string
}

type ModifyLoadBalancerPayTypeResponseBody struct {
	// The order ID of the subscription CLB instance.
	//
	// example:
	//
	// 20212961978891
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLoadBalancerPayTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLoadBalancerPayTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLoadBalancerPayTypeResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ModifyLoadBalancerPayTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLoadBalancerPayTypeResponseBody) SetOrderId(v int64) *ModifyLoadBalancerPayTypeResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyLoadBalancerPayTypeResponseBody) SetRequestId(v string) *ModifyLoadBalancerPayTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLoadBalancerPayTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
