// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySslVpnServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCipher(v string) *ModifySslVpnServerRequest
	GetCipher() *string
	SetClientIpPool(v string) *ModifySslVpnServerRequest
	GetClientIpPool() *string
	SetClientToken(v string) *ModifySslVpnServerRequest
	GetClientToken() *string
	SetCompress(v bool) *ModifySslVpnServerRequest
	GetCompress() *bool
	SetDryRun(v bool) *ModifySslVpnServerRequest
	GetDryRun() *bool
	SetEnableMultiFactorAuth(v bool) *ModifySslVpnServerRequest
	GetEnableMultiFactorAuth() *bool
	SetIDaaSApplicationId(v string) *ModifySslVpnServerRequest
	GetIDaaSApplicationId() *string
	SetIDaaSInstanceId(v string) *ModifySslVpnServerRequest
	GetIDaaSInstanceId() *string
	SetIDaaSRegionId(v string) *ModifySslVpnServerRequest
	GetIDaaSRegionId() *string
	SetLocalSubnet(v string) *ModifySslVpnServerRequest
	GetLocalSubnet() *string
	SetName(v string) *ModifySslVpnServerRequest
	GetName() *string
	SetOwnerAccount(v string) *ModifySslVpnServerRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySslVpnServerRequest
	GetOwnerId() *int64
	SetPort(v int32) *ModifySslVpnServerRequest
	GetPort() *int32
	SetProto(v string) *ModifySslVpnServerRequest
	GetProto() *string
	SetRegionId(v string) *ModifySslVpnServerRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifySslVpnServerRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySslVpnServerRequest
	GetResourceOwnerId() *int64
	SetSslVpnServerId(v string) *ModifySslVpnServerRequest
	GetSslVpnServerId() *string
}

type ModifySslVpnServerRequest struct {
	// The encryption algorithm that is used in the SSL-VPN connection. Valid values:
	//
	// 	- **AES-128-CBC*	- (default)
	//
	// 	- **AES-192-CBC**
	//
	// 	- **AES-256-CBC**
	//
	// 	- **none**
	//
	// example:
	//
	// AES-128-CBC
	Cipher *string `json:"Cipher,omitempty" xml:"Cipher,omitempty"`
	// The client CIDR block.
	//
	// The CIDR block from which an IP address is allocated to the virtual network interface controller (NIC) of the client, rather than the private CIDR block.
	//
	// If the client accesses the SSL server over an SSL-VPN connection, the VPN gateway assigns an IP address from the specified client CIDR block for the client to access cloud resources.
	//
	// Make sure that the number of IP addresses in the client CIDR block is at least four times the maximum number of SSL-VPN connections supported by the VPN gateway.
	//
	// <details>
	//
	// <summary>Click to view the reason.</summary>
	//
	// For example, if you specify 192.168.0.0/24 as the client CIDR block, the system first divides a subnet CIDR block with a subnet mask of 30 from 192.168.0.0/24, such as 192.168.0.4/30. This subnet provides up to four IP addresses. Then, the system allocates an IP address from 192.168.0.4/30 to the client and uses the other three IP addresses to ensure network communication. In this case, one client consumes four IP addresses. Therefore, to ensure that an IP address is assigned to your client, the number of IP addresses in the client CIDR block must be at least four times the maximum number of SSL-VPN connections supported by the VPN gateway with which the SSL server is associated.
	//
	// </details>
	//
	// <details>
	//
	// <summary>Click to view the CIDR blocks that are not supported.</summary>
	//
	// 	- 100.64.0.0~100.127.255.255
	//
	// 	- 127.0.0.0~127.255.255.255
	//
	// 	- 169.254.0.0~169.254.255.255
	//
	// 	- 224.0.0.0~239.255.255.255
	//
	// 	- 255.0.0.0~255.255.255.255
	//
	// </details>
	//
	// <details>
	//
	// <summary>Click to view the recommended client CIDR blocks for different numbers of SSL-VPN connections.</summary>
	//
	// 	- If the number of SSL-VPN connections is 5, we recommend that you specify a client CIDR block with a subnet mask that is less than or equal to 27 bits in length. Examples: 10.0.0.0/27 and 10.0.0.0/26.
	//
	// 	- If the number of SSL-VPN connections is 10, we recommend that you specify a client CIDR block with a subnet mask that is less than or equal to 26 bits in length. Examples: 10.0.0.0/26 and 10.0.0.0/25.
	//
	// 	- If the number of SSL-VPN connections is 20, we recommend that you specify a client CIDR block with a subnet mask that is less than or equal to 25 bits in length. Examples: 10.0.0.0/25 and 10.0.0.0/24.
	//
	// 	- If the number of SSL-VPN connections is 50, we recommend that you specify a client CIDR block with a subnet mask that is less than or equal to 24 bits in length. Examples: 10.0.0.0/24 and 10.0.0.0/23.
	//
	// 	- If the number of SSL-VPN connections is 100, we recommend that you specify a client CIDR block with a subnet mask that is less than or equal to 23 bits in length. Examples: 10.0.0.0/23 and 10.0.0.0/22.
	//
	// 	- If the number of SSL-VPN connections is 200, we recommend that you specify a client CIDR block with a subnet mask that is less than or equal to 22 bits in length. Examples: 10.0.0.0/22 and 10.0.0.0/21.
	//
	// 	- If the number of SSL-VPN connections is 500, we recommend that you specify a client CIDR block with a subnet mask that is less than or equal to 21 bits in length. Examples: 10.0.0.0/21 and 10.0.0.0/20.
	//
	// 	- If the number of SSL-VPN connections is 1,000, we recommend that you specify a client CIDR block with a subnet mask that is less than or equal to 20 bits in length. Examples: 10.0.0.0/20 and 10.0.0.0/19.
	//
	// </details>
	//
	// > - The subnet mask of the client CIDR block must be 16 to 29 bits in length.
	//
	// > -  Make sure that the client CIDR block does not overlap with the local CIDR block, the VPC CIDR block, or route CIDR blocks associated with the client.
	//
	// > - We recommend that you use 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16, or one of their subnets as the client CIDR block. If you want to specify a public CIDR block as the client CIDR block, you must specify the public CIDR block as the user CIDR block of the virtual private cloud (VPC). This way, the VPC can access the public CIDR block. For more information, see [VPC FAQs](https://help.aliyun.com/document_detail/185311.html).
	//
	// > - After you create an SSL server, the system automatically adds routes that point to the client CIDR block to the VPC route table. Do not manually add routes that point to the client CIDR block. Otherwise, SSL-VPN connections cannot work as expected.
	//
	// example:
	//
	// 10.30.30.0/24
	ClientIpPool *string `json:"ClientIpPool,omitempty" xml:"ClientIpPool,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- is different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-0016e04115b
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to enable data compression. Valid values:
	//
	// 	- **true*	- (default)
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Compress *bool `json:"Compress,omitempty" xml:"Compress,omitempty"`
	DryRun   *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to enable two-factor authentication. To enable two-factor authentication, you need to specify **IDaaSInstanceId**, **IDaaSRegionId**, and **IDaaSApplicationId**. Valid values:
	//
	// 	- **true**: enables the feature.
	//
	// 	- **false**: disables the feature.
	//
	// > -  If you use two-factor authentication for the first time, you must first complete [authorization](https://ram.console.aliyun.com/role/authorization?request=%7B%22Services%22%3A%5B%7B%22Service%22%3A%22VPN%22%2C%22Roles%22%3A%5B%7B%22RoleName%22%3A%22AliyunVpnAccessingIdaasRole%22%2C%22TemplateId%22%3A%22IdaasRole%22%7D%5D%7D%5D%2C%22ReturnUrl%22%3A%22https%3A%2F%2Fvpc.console.aliyun.com%2Fsslvpn%2Fcn-shanghai%2Fvpn-servers%22%7D).
	//
	// > - When you create an SSL server in the UAE (Dubai) region, we recommend that you associate the SSL server with an IDaaS EIAM 2.0 instance in Singapore to reduce latency.
	//
	// > - IDaaS EIAM 1.0 instances are no longer available for purchase. If your Alibaba Cloud account has IDaaS EIAM 1.0 instances, the IDaaS EIAM 1.0 instances can be associated after two-factor authentication is enabled. If your Alibaba Cloud account does not have IDaaS EIAM 1.0 instances, only IDaaS EIAM 2.0 instances can be associated after two-factor authentication is enabled.
	//
	// example:
	//
	// false
	EnableMultiFactorAuth *bool `json:"EnableMultiFactorAuth,omitempty" xml:"EnableMultiFactorAuth,omitempty"`
	// The ID of the IDaaS application.
	//
	// 	- If an IDaaS EIAM 2.0 instance is associated, you need to specify an IDaaS application ID.
	//
	// 	- If an IDaaS EIAM 1.0 instance is associated, you do not need to specify an IDaaS application ID.
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
	// The region ID of the IDaaS EIAM instance.
	//
	// example:
	//
	// cn-hangzhou
	IDaaSRegionId *string `json:"IDaaSRegionId,omitempty" xml:"IDaaSRegionId,omitempty"`
	// The local CIDR block.
	//
	// The CIDR block that your client needs to access by using the SSL-VPN connection.
	//
	// This value can be the CIDR block of a VPC, a vSwitch, a data center that is connected to a VPC by using an Express Connect circuit, or an Alibaba Cloud service such as Object Storage Service (OSS).
	//
	// The subnet mask of the specified local CIDR block must be 8 to 32 bits in length. You cannot specify the following CIDR blocks as the local CIDR blocks:
	//
	// 	- 127.0.0.0~127.255.255.255
	//
	// 	- 169.254.0.0~169.254.255.255
	//
	// 	- 224.0.0.0~239.255.255.255
	//
	// 	- 255.0.0.0~255.255.255.255
	//
	// example:
	//
	// 10.20.20.0/24
	LocalSubnet *string `json:"LocalSubnet,omitempty" xml:"LocalSubnet,omitempty"`
	// The name of the SSL server.
	//
	// The name must be 1 to 100 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The port that is used by the SSL server. Valid values of port numbers: **1*	- to **65535**. Default value: **1194**.
	//
	// The following ports are not supported: **22**, **2222**, **22222**, **9000**, **9001**, **9002**, **7505**, **80**, **443**, **53**, **68**, **123**, **4510**, **4560**, **500**, and **4500**.
	//
	// example:
	//
	// 1194
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol that is used by the SSL server. Valid values:
	//
	// 	- **TCP*	- (default)
	//
	// 	- **UDP**
	//
	// example:
	//
	// UDP
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The region ID of the VPN gateway.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SSL server.
	//
	// This parameter is required.
	//
	// example:
	//
	// vss-bp18q7hzj6largv4v****
	SslVpnServerId *string `json:"SslVpnServerId,omitempty" xml:"SslVpnServerId,omitempty"`
}

func (s ModifySslVpnServerRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySslVpnServerRequest) GoString() string {
	return s.String()
}

func (s *ModifySslVpnServerRequest) GetCipher() *string {
	return s.Cipher
}

func (s *ModifySslVpnServerRequest) GetClientIpPool() *string {
	return s.ClientIpPool
}

func (s *ModifySslVpnServerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifySslVpnServerRequest) GetCompress() *bool {
	return s.Compress
}

func (s *ModifySslVpnServerRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifySslVpnServerRequest) GetEnableMultiFactorAuth() *bool {
	return s.EnableMultiFactorAuth
}

func (s *ModifySslVpnServerRequest) GetIDaaSApplicationId() *string {
	return s.IDaaSApplicationId
}

func (s *ModifySslVpnServerRequest) GetIDaaSInstanceId() *string {
	return s.IDaaSInstanceId
}

func (s *ModifySslVpnServerRequest) GetIDaaSRegionId() *string {
	return s.IDaaSRegionId
}

func (s *ModifySslVpnServerRequest) GetLocalSubnet() *string {
	return s.LocalSubnet
}

func (s *ModifySslVpnServerRequest) GetName() *string {
	return s.Name
}

func (s *ModifySslVpnServerRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySslVpnServerRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySslVpnServerRequest) GetPort() *int32 {
	return s.Port
}

func (s *ModifySslVpnServerRequest) GetProto() *string {
	return s.Proto
}

func (s *ModifySslVpnServerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySslVpnServerRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySslVpnServerRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySslVpnServerRequest) GetSslVpnServerId() *string {
	return s.SslVpnServerId
}

func (s *ModifySslVpnServerRequest) SetCipher(v string) *ModifySslVpnServerRequest {
	s.Cipher = &v
	return s
}

func (s *ModifySslVpnServerRequest) SetClientIpPool(v string) *ModifySslVpnServerRequest {
	s.ClientIpPool = &v
	return s
}

func (s *ModifySslVpnServerRequest) SetClientToken(v string) *ModifySslVpnServerRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifySslVpnServerRequest) SetCompress(v bool) *ModifySslVpnServerRequest {
	s.Compress = &v
	return s
}

func (s *ModifySslVpnServerRequest) SetDryRun(v bool) *ModifySslVpnServerRequest {
	s.DryRun = &v
	return s
}

func (s *ModifySslVpnServerRequest) SetEnableMultiFactorAuth(v bool) *ModifySslVpnServerRequest {
	s.EnableMultiFactorAuth = &v
	return s
}

func (s *ModifySslVpnServerRequest) SetIDaaSApplicationId(v string) *ModifySslVpnServerRequest {
	s.IDaaSApplicationId = &v
	return s
}

func (s *ModifySslVpnServerRequest) SetIDaaSInstanceId(v string) *ModifySslVpnServerRequest {
	s.IDaaSInstanceId = &v
	return s
}

func (s *ModifySslVpnServerRequest) SetIDaaSRegionId(v string) *ModifySslVpnServerRequest {
	s.IDaaSRegionId = &v
	return s
}

func (s *ModifySslVpnServerRequest) SetLocalSubnet(v string) *ModifySslVpnServerRequest {
	s.LocalSubnet = &v
	return s
}

func (s *ModifySslVpnServerRequest) SetName(v string) *ModifySslVpnServerRequest {
	s.Name = &v
	return s
}

func (s *ModifySslVpnServerRequest) SetOwnerAccount(v string) *ModifySslVpnServerRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySslVpnServerRequest) SetOwnerId(v int64) *ModifySslVpnServerRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySslVpnServerRequest) SetPort(v int32) *ModifySslVpnServerRequest {
	s.Port = &v
	return s
}

func (s *ModifySslVpnServerRequest) SetProto(v string) *ModifySslVpnServerRequest {
	s.Proto = &v
	return s
}

func (s *ModifySslVpnServerRequest) SetRegionId(v string) *ModifySslVpnServerRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySslVpnServerRequest) SetResourceOwnerAccount(v string) *ModifySslVpnServerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySslVpnServerRequest) SetResourceOwnerId(v int64) *ModifySslVpnServerRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySslVpnServerRequest) SetSslVpnServerId(v string) *ModifySslVpnServerRequest {
	s.SslVpnServerId = &v
	return s
}

func (s *ModifySslVpnServerRequest) Validate() error {
	return dara.Validate(s)
}
