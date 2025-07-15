// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpsecServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreationTime(v string) *CreateIpsecServerResponseBody
	GetCreationTime() *string
	SetIpsecServerId(v string) *CreateIpsecServerResponseBody
	GetIpsecServerId() *string
	SetIpsecServerName(v string) *CreateIpsecServerResponseBody
	GetIpsecServerName() *string
	SetRegionId(v string) *CreateIpsecServerResponseBody
	GetRegionId() *string
	SetRequestId(v string) *CreateIpsecServerResponseBody
	GetRequestId() *string
	SetVpnGatewayId(v string) *CreateIpsecServerResponseBody
	GetVpnGatewayId() *string
}

type CreateIpsecServerResponseBody struct {
	// The time when the IPsec server was created.
	//
	// T is used as a delimiter. Z indicates that the time is in UTC.
	//
	// example:
	//
	// 2021-02-22T03:24:28Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The IPsec server ID.
	//
	// example:
	//
	// iss-bp1jougp8cfsbo8y9****
	IpsecServerId *string `json:"IpsecServerId,omitempty" xml:"IpsecServerId,omitempty"`
	// The IPsec server name.
	//
	// example:
	//
	// test
	IpsecServerName *string `json:"IpsecServerName,omitempty" xml:"IpsecServerName,omitempty"`
	// The ID of the region where the VPN gateway is deployed.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 690A967E-D4CD-4B69-8C78-94FE828BA10B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the VPN gateway.
	//
	// example:
	//
	// vpn-bp17lofy9fd0dnvzv****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
}

func (s CreateIpsecServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIpsecServerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIpsecServerResponseBody) GetCreationTime() *string {
	return s.CreationTime
}

func (s *CreateIpsecServerResponseBody) GetIpsecServerId() *string {
	return s.IpsecServerId
}

func (s *CreateIpsecServerResponseBody) GetIpsecServerName() *string {
	return s.IpsecServerName
}

func (s *CreateIpsecServerResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateIpsecServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIpsecServerResponseBody) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *CreateIpsecServerResponseBody) SetCreationTime(v string) *CreateIpsecServerResponseBody {
	s.CreationTime = &v
	return s
}

func (s *CreateIpsecServerResponseBody) SetIpsecServerId(v string) *CreateIpsecServerResponseBody {
	s.IpsecServerId = &v
	return s
}

func (s *CreateIpsecServerResponseBody) SetIpsecServerName(v string) *CreateIpsecServerResponseBody {
	s.IpsecServerName = &v
	return s
}

func (s *CreateIpsecServerResponseBody) SetRegionId(v string) *CreateIpsecServerResponseBody {
	s.RegionId = &v
	return s
}

func (s *CreateIpsecServerResponseBody) SetRequestId(v string) *CreateIpsecServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIpsecServerResponseBody) SetVpnGatewayId(v string) *CreateIpsecServerResponseBody {
	s.VpnGatewayId = &v
	return s
}

func (s *CreateIpsecServerResponseBody) Validate() error {
	return dara.Validate(s)
}
