// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnhancedVpnGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateEnhancedVpnGatewayResponseBody
	GetName() *string
	SetRequestId(v string) *CreateEnhancedVpnGatewayResponseBody
	GetRequestId() *string
	SetVpnGatewayId(v string) *CreateEnhancedVpnGatewayResponseBody
	GetVpnGatewayId() *string
}

type CreateEnhancedVpnGatewayResponseBody struct {
	// example:
	//
	// MYVPN
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// EB2C156A-41F8-49CC-A756-D55AFC8BFD69
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// vpn-uf68lxhgr7ftbqr3p****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
}

func (s CreateEnhancedVpnGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEnhancedVpnGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEnhancedVpnGatewayResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateEnhancedVpnGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEnhancedVpnGatewayResponseBody) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *CreateEnhancedVpnGatewayResponseBody) SetName(v string) *CreateEnhancedVpnGatewayResponseBody {
	s.Name = &v
	return s
}

func (s *CreateEnhancedVpnGatewayResponseBody) SetRequestId(v string) *CreateEnhancedVpnGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEnhancedVpnGatewayResponseBody) SetVpnGatewayId(v string) *CreateEnhancedVpnGatewayResponseBody {
	s.VpnGatewayId = &v
	return s
}

func (s *CreateEnhancedVpnGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
