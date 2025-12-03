// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLoadBalancerInternetSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *ModifyLoadBalancerInternetSpecResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *ModifyLoadBalancerInternetSpecResponseBody
	GetRequestId() *string
}

type ModifyLoadBalancerInternetSpecResponseBody struct {
	// The order ID of the subscription CLB instance.
	//
	// example:
	//
	// 20142961978****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLoadBalancerInternetSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLoadBalancerInternetSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLoadBalancerInternetSpecResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ModifyLoadBalancerInternetSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLoadBalancerInternetSpecResponseBody) SetOrderId(v int64) *ModifyLoadBalancerInternetSpecResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyLoadBalancerInternetSpecResponseBody) SetRequestId(v string) *ModifyLoadBalancerInternetSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLoadBalancerInternetSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
