// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNetworkPackageBandwidthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *ModifyNetworkPackageBandwidthResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ModifyNetworkPackageBandwidthResponseBody
	GetRequestId() *string
}

type ModifyNetworkPackageBandwidthResponseBody struct {
	// The ID of the order.
	//
	// example:
	//
	// 214552063030752
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNetworkPackageBandwidthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyNetworkPackageBandwidthResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNetworkPackageBandwidthResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyNetworkPackageBandwidthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyNetworkPackageBandwidthResponseBody) SetOrderId(v string) *ModifyNetworkPackageBandwidthResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyNetworkPackageBandwidthResponseBody) SetRequestId(v string) *ModifyNetworkPackageBandwidthResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyNetworkPackageBandwidthResponseBody) Validate() error {
	return dara.Validate(s)
}
