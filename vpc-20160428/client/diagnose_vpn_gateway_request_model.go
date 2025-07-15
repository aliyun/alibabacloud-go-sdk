// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiagnoseVpnGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DiagnoseVpnGatewayRequest
	GetClientToken() *string
	SetIPsecExtendInfo(v string) *DiagnoseVpnGatewayRequest
	GetIPsecExtendInfo() *string
	SetRegionId(v string) *DiagnoseVpnGatewayRequest
	GetRegionId() *string
	SetResourceId(v string) *DiagnoseVpnGatewayRequest
	GetResourceId() *string
	SetResourceType(v string) *DiagnoseVpnGatewayRequest
	GetResourceType() *string
	SetVpnGatewayId(v string) *DiagnoseVpnGatewayRequest
	GetVpnGatewayId() *string
}

type DiagnoseVpnGatewayRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Check the connectivity of the destination address. Valid values:
	//
	// 	- **PrivateSourceIp**: the source IP address. The source IP address must be on the VPC side.
	//
	// 	- **PrivateDestinationIp**: the destination IP address. The destination IP address must be on the data center side.
	//
	// example:
	//
	// {"PrivateSourceIp":"192.168.1.1","PrivateDestinationIp":"192.168.0.1"}
	IPsecExtendInfo *string `json:"IPsecExtendInfo,omitempty" xml:"IPsecExtendInfo,omitempty"`
	// The region ID of the VPN gateway.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource to be diagnosed.
	//
	// This parameter is required.
	//
	// example:
	//
	// vco-uf66xniofskqtuoz1****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource.
	//
	// Set the value to **Ipsec**, which specifies an IPsec-VPN connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// IPsec
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the VPN gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpn-m5efhj0k1p47ctuhl****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
}

func (s DiagnoseVpnGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s DiagnoseVpnGatewayRequest) GoString() string {
	return s.String()
}

func (s *DiagnoseVpnGatewayRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DiagnoseVpnGatewayRequest) GetIPsecExtendInfo() *string {
	return s.IPsecExtendInfo
}

func (s *DiagnoseVpnGatewayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DiagnoseVpnGatewayRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DiagnoseVpnGatewayRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DiagnoseVpnGatewayRequest) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *DiagnoseVpnGatewayRequest) SetClientToken(v string) *DiagnoseVpnGatewayRequest {
	s.ClientToken = &v
	return s
}

func (s *DiagnoseVpnGatewayRequest) SetIPsecExtendInfo(v string) *DiagnoseVpnGatewayRequest {
	s.IPsecExtendInfo = &v
	return s
}

func (s *DiagnoseVpnGatewayRequest) SetRegionId(v string) *DiagnoseVpnGatewayRequest {
	s.RegionId = &v
	return s
}

func (s *DiagnoseVpnGatewayRequest) SetResourceId(v string) *DiagnoseVpnGatewayRequest {
	s.ResourceId = &v
	return s
}

func (s *DiagnoseVpnGatewayRequest) SetResourceType(v string) *DiagnoseVpnGatewayRequest {
	s.ResourceType = &v
	return s
}

func (s *DiagnoseVpnGatewayRequest) SetVpnGatewayId(v string) *DiagnoseVpnGatewayRequest {
	s.VpnGatewayId = &v
	return s
}

func (s *DiagnoseVpnGatewayRequest) Validate() error {
	return dara.Validate(s)
}
