// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpnGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateVpnGatewayResponseBody
	GetName() *string
	SetOrderId(v int64) *CreateVpnGatewayResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *CreateVpnGatewayResponseBody
	GetRequestId() *string
	SetVpnGatewayId(v string) *CreateVpnGatewayResponseBody
	GetVpnGatewayId() *string
}

type CreateVpnGatewayResponseBody struct {
	// The name of the VPN gateway.
	//
	// example:
	//
	// MYVPN
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The order ID.
	//
	// If automatic payment is disabled, you must manually complete the payment for the VPN gateway in the [Alibaba Cloud Management console](https://usercenter2-intl.aliyun.com/billing/#/account/overview).
	//
	// example:
	//
	// 208240895400460
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EB2C156A-41F8-49CC-A756-D55AFC8BFD69
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the VPN gateway.
	//
	// example:
	//
	// vpn-uf68lxhgr7ftbqr3p****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
}

func (s CreateVpnGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVpnGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpnGatewayResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateVpnGatewayResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateVpnGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVpnGatewayResponseBody) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *CreateVpnGatewayResponseBody) SetName(v string) *CreateVpnGatewayResponseBody {
	s.Name = &v
	return s
}

func (s *CreateVpnGatewayResponseBody) SetOrderId(v int64) *CreateVpnGatewayResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateVpnGatewayResponseBody) SetRequestId(v string) *CreateVpnGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVpnGatewayResponseBody) SetVpnGatewayId(v string) *CreateVpnGatewayResponseBody {
	s.VpnGatewayId = &v
	return s
}

func (s *CreateVpnGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
