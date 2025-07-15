// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySslVpnServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCipher(v string) *ModifySslVpnServerResponseBody
	GetCipher() *string
	SetClientIpPool(v string) *ModifySslVpnServerResponseBody
	GetClientIpPool() *string
	SetCompress(v bool) *ModifySslVpnServerResponseBody
	GetCompress() *bool
	SetConnections(v int32) *ModifySslVpnServerResponseBody
	GetConnections() *int32
	SetCreateTime(v int64) *ModifySslVpnServerResponseBody
	GetCreateTime() *int64
	SetEnableMultiFactorAuth(v bool) *ModifySslVpnServerResponseBody
	GetEnableMultiFactorAuth() *bool
	SetIDaaSApplicationId(v string) *ModifySslVpnServerResponseBody
	GetIDaaSApplicationId() *string
	SetIDaaSInstanceId(v string) *ModifySslVpnServerResponseBody
	GetIDaaSInstanceId() *string
	SetIDaaSInstanceVersion(v string) *ModifySslVpnServerResponseBody
	GetIDaaSInstanceVersion() *string
	SetInternetIp(v string) *ModifySslVpnServerResponseBody
	GetInternetIp() *string
	SetLocalSubnet(v string) *ModifySslVpnServerResponseBody
	GetLocalSubnet() *string
	SetMaxConnections(v int32) *ModifySslVpnServerResponseBody
	GetMaxConnections() *int32
	SetName(v string) *ModifySslVpnServerResponseBody
	GetName() *string
	SetPort(v int32) *ModifySslVpnServerResponseBody
	GetPort() *int32
	SetProto(v string) *ModifySslVpnServerResponseBody
	GetProto() *string
	SetRegionId(v string) *ModifySslVpnServerResponseBody
	GetRegionId() *string
	SetRequestId(v string) *ModifySslVpnServerResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *ModifySslVpnServerResponseBody
	GetResourceGroupId() *string
	SetSslVpnServerId(v string) *ModifySslVpnServerResponseBody
	GetSslVpnServerId() *string
	SetVpnGatewayId(v string) *ModifySslVpnServerResponseBody
	GetVpnGatewayId() *string
}

type ModifySslVpnServerResponseBody struct {
	// The encryption algorithm.
	//
	// example:
	//
	// AES-128-CBC
	Cipher *string `json:"Cipher,omitempty" xml:"Cipher,omitempty"`
	// The client CIDR block.
	//
	// example:
	//
	// 10.30.30.0/24
	ClientIpPool *string `json:"ClientIpPool,omitempty" xml:"ClientIpPool,omitempty"`
	// Indicates whether data compression is enabled.
	//
	// example:
	//
	// false
	Compress *bool `json:"Compress,omitempty" xml:"Compress,omitempty"`
	// The total number of current connections.
	//
	// example:
	//
	// 0
	Connections *int32 `json:"Connections,omitempty" xml:"Connections,omitempty"`
	// The time when the SSL server was created.
	//
	// example:
	//
	// 1492753580000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether two-factor authentication is enabled.
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	EnableMultiFactorAuth *bool `json:"EnableMultiFactorAuth,omitempty" xml:"EnableMultiFactorAuth,omitempty"`
	// The ID of the IDaaS application.
	//
	// example:
	//
	// app_my6g4qmvnwxzj2f****
	IDaaSApplicationId *string `json:"IDaaSApplicationId,omitempty" xml:"IDaaSApplicationId,omitempty"`
	// The ID of the IDaaS EIAM instance.
	//
	// example:
	//
	// idaas-cn-hangzhou-****
	IDaaSInstanceId *string `json:"IDaaSInstanceId,omitempty" xml:"IDaaSInstanceId,omitempty"`
	// The version of the IDaaS EIAM instance.
	//
	// 	- This parameter is returned only if the SSL server is associated with an IDaaS EIAM 2.0 instance. Only **EIAM 2.0*	- is returned.
	//
	// 	- If the SSL server is associated with an IDaaS EIAM 1.0 instance, no value is returned.
	//
	// example:
	//
	// EIAM 2.0
	IDaaSInstanceVersion *string `json:"IDaaSInstanceVersion,omitempty" xml:"IDaaSInstanceVersion,omitempty"`
	// The public IP address.
	//
	// example:
	//
	// 47.98.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The local CIDR block.
	//
	// example:
	//
	// 10.20.20.0/24
	LocalSubnet *string `json:"LocalSubnet,omitempty" xml:"LocalSubnet,omitempty"`
	// The maximum number of connections.
	//
	// example:
	//
	// 5
	MaxConnections *int32 `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	// The name of the SSL server.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The port that is used by the SSL server.
	//
	// example:
	//
	// 1194
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol that is used by the SSL server.
	//
	// example:
	//
	// UDP
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The ID of the region where the SSL server is created.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DF11D6F6-E35A-41C3-9B20-6FC8A901FE65
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the SSL server belongs.
	//
	// The SSL server and the VPN gateway associated with the SSL server belong to the same resource group. You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation to query resource groups.
	//
	// example:
	//
	// rg-acfmzs372yg****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the SSL server.
	//
	// example:
	//
	// vss-bp18q7hzj6largv4v****
	SslVpnServerId *string `json:"SslVpnServerId,omitempty" xml:"SslVpnServerId,omitempty"`
	// The ID of the VPN gateway.
	//
	// example:
	//
	// vpn-bp1q8bgx4xnkm2ogj****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
}

func (s ModifySslVpnServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySslVpnServerResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySslVpnServerResponseBody) GetCipher() *string {
	return s.Cipher
}

func (s *ModifySslVpnServerResponseBody) GetClientIpPool() *string {
	return s.ClientIpPool
}

func (s *ModifySslVpnServerResponseBody) GetCompress() *bool {
	return s.Compress
}

func (s *ModifySslVpnServerResponseBody) GetConnections() *int32 {
	return s.Connections
}

func (s *ModifySslVpnServerResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ModifySslVpnServerResponseBody) GetEnableMultiFactorAuth() *bool {
	return s.EnableMultiFactorAuth
}

func (s *ModifySslVpnServerResponseBody) GetIDaaSApplicationId() *string {
	return s.IDaaSApplicationId
}

func (s *ModifySslVpnServerResponseBody) GetIDaaSInstanceId() *string {
	return s.IDaaSInstanceId
}

func (s *ModifySslVpnServerResponseBody) GetIDaaSInstanceVersion() *string {
	return s.IDaaSInstanceVersion
}

func (s *ModifySslVpnServerResponseBody) GetInternetIp() *string {
	return s.InternetIp
}

func (s *ModifySslVpnServerResponseBody) GetLocalSubnet() *string {
	return s.LocalSubnet
}

func (s *ModifySslVpnServerResponseBody) GetMaxConnections() *int32 {
	return s.MaxConnections
}

func (s *ModifySslVpnServerResponseBody) GetName() *string {
	return s.Name
}

func (s *ModifySslVpnServerResponseBody) GetPort() *int32 {
	return s.Port
}

func (s *ModifySslVpnServerResponseBody) GetProto() *string {
	return s.Proto
}

func (s *ModifySslVpnServerResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySslVpnServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySslVpnServerResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifySslVpnServerResponseBody) GetSslVpnServerId() *string {
	return s.SslVpnServerId
}

func (s *ModifySslVpnServerResponseBody) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *ModifySslVpnServerResponseBody) SetCipher(v string) *ModifySslVpnServerResponseBody {
	s.Cipher = &v
	return s
}

func (s *ModifySslVpnServerResponseBody) SetClientIpPool(v string) *ModifySslVpnServerResponseBody {
	s.ClientIpPool = &v
	return s
}

func (s *ModifySslVpnServerResponseBody) SetCompress(v bool) *ModifySslVpnServerResponseBody {
	s.Compress = &v
	return s
}

func (s *ModifySslVpnServerResponseBody) SetConnections(v int32) *ModifySslVpnServerResponseBody {
	s.Connections = &v
	return s
}

func (s *ModifySslVpnServerResponseBody) SetCreateTime(v int64) *ModifySslVpnServerResponseBody {
	s.CreateTime = &v
	return s
}

func (s *ModifySslVpnServerResponseBody) SetEnableMultiFactorAuth(v bool) *ModifySslVpnServerResponseBody {
	s.EnableMultiFactorAuth = &v
	return s
}

func (s *ModifySslVpnServerResponseBody) SetIDaaSApplicationId(v string) *ModifySslVpnServerResponseBody {
	s.IDaaSApplicationId = &v
	return s
}

func (s *ModifySslVpnServerResponseBody) SetIDaaSInstanceId(v string) *ModifySslVpnServerResponseBody {
	s.IDaaSInstanceId = &v
	return s
}

func (s *ModifySslVpnServerResponseBody) SetIDaaSInstanceVersion(v string) *ModifySslVpnServerResponseBody {
	s.IDaaSInstanceVersion = &v
	return s
}

func (s *ModifySslVpnServerResponseBody) SetInternetIp(v string) *ModifySslVpnServerResponseBody {
	s.InternetIp = &v
	return s
}

func (s *ModifySslVpnServerResponseBody) SetLocalSubnet(v string) *ModifySslVpnServerResponseBody {
	s.LocalSubnet = &v
	return s
}

func (s *ModifySslVpnServerResponseBody) SetMaxConnections(v int32) *ModifySslVpnServerResponseBody {
	s.MaxConnections = &v
	return s
}

func (s *ModifySslVpnServerResponseBody) SetName(v string) *ModifySslVpnServerResponseBody {
	s.Name = &v
	return s
}

func (s *ModifySslVpnServerResponseBody) SetPort(v int32) *ModifySslVpnServerResponseBody {
	s.Port = &v
	return s
}

func (s *ModifySslVpnServerResponseBody) SetProto(v string) *ModifySslVpnServerResponseBody {
	s.Proto = &v
	return s
}

func (s *ModifySslVpnServerResponseBody) SetRegionId(v string) *ModifySslVpnServerResponseBody {
	s.RegionId = &v
	return s
}

func (s *ModifySslVpnServerResponseBody) SetRequestId(v string) *ModifySslVpnServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySslVpnServerResponseBody) SetResourceGroupId(v string) *ModifySslVpnServerResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifySslVpnServerResponseBody) SetSslVpnServerId(v string) *ModifySslVpnServerResponseBody {
	s.SslVpnServerId = &v
	return s
}

func (s *ModifySslVpnServerResponseBody) SetVpnGatewayId(v string) *ModifySslVpnServerResponseBody {
	s.VpnGatewayId = &v
	return s
}

func (s *ModifySslVpnServerResponseBody) Validate() error {
	return dara.Validate(s)
}
