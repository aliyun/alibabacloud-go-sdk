// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSslVpnServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCipher(v string) *CreateSslVpnServerRequest
	GetCipher() *string
	SetClientIpPool(v string) *CreateSslVpnServerRequest
	GetClientIpPool() *string
	SetClientToken(v string) *CreateSslVpnServerRequest
	GetClientToken() *string
	SetCompress(v bool) *CreateSslVpnServerRequest
	GetCompress() *bool
	SetDnsServers(v string) *CreateSslVpnServerRequest
	GetDnsServers() *string
	SetDryRun(v bool) *CreateSslVpnServerRequest
	GetDryRun() *bool
	SetEnableMultiFactorAuth(v bool) *CreateSslVpnServerRequest
	GetEnableMultiFactorAuth() *bool
	SetIDaaSApplicationId(v string) *CreateSslVpnServerRequest
	GetIDaaSApplicationId() *string
	SetIDaaSInstanceId(v string) *CreateSslVpnServerRequest
	GetIDaaSInstanceId() *string
	SetIDaaSRegionId(v string) *CreateSslVpnServerRequest
	GetIDaaSRegionId() *string
	SetLocalSubnet(v string) *CreateSslVpnServerRequest
	GetLocalSubnet() *string
	SetName(v string) *CreateSslVpnServerRequest
	GetName() *string
	SetOwnerAccount(v string) *CreateSslVpnServerRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateSslVpnServerRequest
	GetOwnerId() *int64
	SetPort(v int32) *CreateSslVpnServerRequest
	GetPort() *int32
	SetProto(v string) *CreateSslVpnServerRequest
	GetProto() *string
	SetRegionId(v string) *CreateSslVpnServerRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateSslVpnServerRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateSslVpnServerRequest
	GetResourceOwnerId() *int64
	SetVpnGatewayId(v string) *CreateSslVpnServerRequest
	GetVpnGatewayId() *string
}

type CreateSslVpnServerRequest struct {
	// The encryption algorithm used by the SSL-VPN connection.
	//
	// - If the client uses Tunnelblick or OpenVPN 2.4.0 or later, the SSL server and the client dynamically negotiate the encryption algorithm and preferentially use the encryption algorithm with the highest security level supported by both parties. The encryption algorithm you specify for the SSL server does not take effect.
	//
	// - If the client uses OpenVPN earlier than 2.4.0, the SSL server and the client use the encryption algorithm you specify for the SSL server. The SSL server supports the following encryption algorithms:
	//
	//   - **AES-128-CBC*	- (default): AES-128-CBC algorithm.
	//
	//   - **AES-192-CBC**: AES-192-CBC algorithm.
	//
	//   - **AES-256-CBC**: AES-256-CBC algorithm.
	//
	//   - **none**: no encryption algorithm is used.
	//
	// example:
	//
	// AES-128-CBC
	Cipher *string `json:"Cipher,omitempty" xml:"Cipher,omitempty"`
	// The client CIDR block.
	//
	// The client CIDR block is used to allocate IP addresses to virtual network interface controllers (NICs) of clients. It does not refer to the existing internal network CIDR block of the client.
	//
	// When a client accesses the local network through an SSL-VPN connection, the VPN gateway allocates an IP address from the specified client CIDR block to the client. The client uses the allocated IP address to access cloud resources.
	//
	// When you specify the client CIDR block, make sure that the number of IP addresses in the client CIDR block is at least four times the number of SSL connections supported by the VPN gateway.
	//
	// <details>
	//
	// <summary>Click to view the reason</summary>
	//
	// For example, if you specify 192.168.0.0/24 as the client CIDR block, the system first divides a subnet with a 30-bit subnet mask from the 192.168.0.0/24 CIDR block, such as 192.168.0.4/30, and then allocates one IP address from 192.168.0.4/30 to the client. The remaining three IP addresses are used by the system to ensure network communication. In this case, one client consumes four IP addresses. Therefore, to ensure that all clients can be allocated IP addresses, make sure that the number of IP addresses in the client CIDR block is at least four times the number of SSL connections supported by the VPN gateway.
	//
	// </details>
	//
	// <details>
	//
	// <summary>Click to view unsupported CIDR blocks</summary>
	//
	// - 100.64.0.0~100.127.255.255
	//
	// - 127.0.0.0~127.255.255.255
	//
	// - 169.254.0.0~169.254.255.255
	//
	// - 224.0.0.0~239.255.255.255
	//
	// - 255.0.0.0~255.255.255.255
	//
	// </details>
	//
	// <details>
	//
	// <summary>Click to view the recommended client CIDR blocks for each number of SSL connections</summary>
	//
	// - If the number of SSL connections is 5, the subnet mask of the client CIDR block must be 27 bits or less. For example: 10.0.0.0/27 or 10.0.0.0/26.
	//
	// - If the number of SSL connections is 10, the subnet mask of the client CIDR block must be 26 bits or less. For example: 10.0.0.0/26 or 10.0.0.0/25.
	//
	// - If the number of SSL connections is 20, the subnet mask of the client CIDR block must be 25 bits or less. For example: 10.0.0.0/25 or 10.0.0.0/24.
	//
	// - If the number of SSL connections is 50, the subnet mask of the client CIDR block must be 24 bits or less. For example: 10.0.0.0/24 or 10.0.0.0/23.
	//
	// - If the number of SSL connections is 100, the subnet mask of the client CIDR block must be 23 bits or less. For example: 10.0.0.0/23 or 10.0.0.0/22.
	//
	// - If the number of SSL connections is 200, the subnet mask of the client CIDR block must be 22 bits or less. For example: 10.0.0.0/22 or 10.0.0.0/21.
	//
	// - If the number of SSL connections is 500, the subnet mask of the client CIDR block must be 21 bits or less. For example: 10.0.0.0/21 or 10.0.0.0/20.
	//
	// - If the number of SSL connections is 1000, the subnet mask of the client CIDR block must be 20 bits or less. For example: 10.0.0.0/20 or 10.0.0.0/19.
	//
	// </details>
	//
	// > - The subnet mask of the client CIDR block must be 16 to 29 bits.
	//
	// > - Make sure that the client CIDR block does not overlap with the local CIDR block, the VPC CIDR block, or any routing CIDR block associated with the client terminal.
	//
	// > - When you specify the client CIDR block, use 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16, or their subnets. If you want to specify a public CIDR block as the client CIDR block, set the public CIDR block as a user CIDR block of the virtual private cloud (VPC) to ensure that the VPC can access the public CIDR block. For more information about user CIDR blocks, see [VPC FAQ](https://help.aliyun.com/document_detail/185311.html).
	//
	// > - After the SSL server is created, the system automatically adds a routing entry for the client CIDR block to the route table of the VPC instance. Do not manually add a routing entry for the client CIDR block to the route table of the VPC instance. Otherwise, SSL-VPN connection traffic may be abnormal.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.1.0/24
	ClientIpPool *string `json:"ClientIpPool,omitempty" xml:"ClientIpPool,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The ClientToken value can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-0016e04115b
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to compress communication. Valid values:
	//
	// - **true**: compresses communication.
	//
	// - **false*	- (default): does not compress communication.
	//
	// example:
	//
	// false
	Compress   *bool   `json:"Compress,omitempty" xml:"Compress,omitempty"`
	DnsServers *string `json:"DnsServers,omitempty" xml:"DnsServers,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including required parameters, request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to enable two-factor authentication. If you enable two-factor authentication, you must also configure `IDaaSInstanceId`, `IDaaSRegionId`, and `IDaaSApplicationId`. Valid values:
	//
	// - **true**: enabled.
	//
	// - **false*	- (default): disabled.
	//
	// > - If you use two-factor authentication for the first time, complete [authorization](https://ram.console.aliyun.com/role/authorization?request=%7B%22Services%22%3A%5B%7B%22Service%22%3A%22VPN%22%2C%22Roles%22%3A%5B%7B%22RoleName%22%3A%22AliyunVpnAccessingIdaasRole%22%2C%22TemplateId%22%3A%22IdaasRole%22%7D%5D%7D%5D%2C%22ReturnUrl%22%3A%22https%3A%2F%2Fvpc.console.aliyun.com%2Fsslvpn%2Fcn-shanghai%2Fvpn-servers%22%7D) before creating the SSL server.
	//
	// > - When you create an SSL server in the UAE (Dubai) region, bind an IDaaS EIAM 2.0 instance in the Singapore region to reduce cross-region latency.
	//
	// > - IDaaS EIAM 1.0 instances are no longer available for purchase. If your Alibaba Cloud account has IDaaS EIAM 1.0 instances, you can still bind IDaaS EIAM 1.0 instances after enabling two-factor authentication. If your Alibaba Cloud account does not have IDaaS EIAM 1.0 instances, you can only bind IDaaS EIAM 2.0 instances after enabling two-factor authentication.
	//
	// example:
	//
	// false
	EnableMultiFactorAuth *bool `json:"EnableMultiFactorAuth,omitempty" xml:"EnableMultiFactorAuth,omitempty"`
	// The ID of the IDaaS application.
	//
	// - If you bind an IDaaS EIAM 2.0 instance, enter the IDaaS application ID.
	//
	// - If you bind an IDaaS EIAM 1.0 instance, you do not need to enter the IDaaS application ID.
	//
	// example:
	//
	// app_my6g4qmvnwxzj2f****
	IDaaSApplicationId *string `json:"IDaaSApplicationId,omitempty" xml:"IDaaSApplicationId,omitempty"`
	// The ID of the IDaaS EIAM instance.
	//
	// example:
	//
	// idaas-cn-hangzhou-p****
	IDaaSInstanceId *string `json:"IDaaSInstanceId,omitempty" xml:"IDaaSInstanceId,omitempty"`
	// The region ID of the IDaaS EIAM instance.
	//
	// example:
	//
	// cn-hangzhou
	IDaaSRegionId *string `json:"IDaaSRegionId,omitempty" xml:"IDaaSRegionId,omitempty"`
	// The local CIDR block.
	//
	// The local CIDR block is the CIDR block that the client needs to access through the SSL-VPN connection.
	//
	// The local CIDR block can be the CIDR block of a VPC, the CIDR block of a vSwitch, the CIDR block of an on-premises data center that is connected to the VPC through an Express Connect circuit, or the CIDR block of a cloud service such as Object Storage Service (OSS).
	//
	// The subnet mask of the local CIDR block must be 8 to 32 bits. The following CIDR blocks cannot be specified as the local CIDR block:
	//
	// - 127.0.0.0~127.255.255.255
	//
	// - 169.254.0.0~169.254.255.255
	//
	// - 224.0.0.0~239.255.255.255
	//
	// - 255.0.0.0~255.255.255.255
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.0.0.0/8
	LocalSubnet *string `json:"LocalSubnet,omitempty" xml:"LocalSubnet,omitempty"`
	// The name of the SSL-VPN server.
	//
	// The name must be 1 to 100 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// sslvpnname
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The port used by the SSL-VPN server. Valid values: **1*	- to **65535**. Default value: **1194**.
	//
	// The following ports are not supported: **22**, **2222**, **22222**, **9000**, **9001**, **9002**, **7505**, **80**, **443**, **53**, **68**, **123**, **4510**, **4560**, **500**, and **4500**.
	//
	// example:
	//
	// 1194
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol used by the SSL-VPN server. Valid values:
	//
	// - **TCP*	- (default): TCP protocol.
	//
	// - **UDP**: UDP protocol.
	//
	// example:
	//
	// UDP
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The region ID of the VPN gateway.
	//
	// You can call [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the VPN gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpn-bp1hgim8by0kc9nga****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
}

func (s CreateSslVpnServerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSslVpnServerRequest) GoString() string {
	return s.String()
}

func (s *CreateSslVpnServerRequest) GetCipher() *string {
	return s.Cipher
}

func (s *CreateSslVpnServerRequest) GetClientIpPool() *string {
	return s.ClientIpPool
}

func (s *CreateSslVpnServerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateSslVpnServerRequest) GetCompress() *bool {
	return s.Compress
}

func (s *CreateSslVpnServerRequest) GetDnsServers() *string {
	return s.DnsServers
}

func (s *CreateSslVpnServerRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateSslVpnServerRequest) GetEnableMultiFactorAuth() *bool {
	return s.EnableMultiFactorAuth
}

func (s *CreateSslVpnServerRequest) GetIDaaSApplicationId() *string {
	return s.IDaaSApplicationId
}

func (s *CreateSslVpnServerRequest) GetIDaaSInstanceId() *string {
	return s.IDaaSInstanceId
}

func (s *CreateSslVpnServerRequest) GetIDaaSRegionId() *string {
	return s.IDaaSRegionId
}

func (s *CreateSslVpnServerRequest) GetLocalSubnet() *string {
	return s.LocalSubnet
}

func (s *CreateSslVpnServerRequest) GetName() *string {
	return s.Name
}

func (s *CreateSslVpnServerRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateSslVpnServerRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateSslVpnServerRequest) GetPort() *int32 {
	return s.Port
}

func (s *CreateSslVpnServerRequest) GetProto() *string {
	return s.Proto
}

func (s *CreateSslVpnServerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSslVpnServerRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateSslVpnServerRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateSslVpnServerRequest) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *CreateSslVpnServerRequest) SetCipher(v string) *CreateSslVpnServerRequest {
	s.Cipher = &v
	return s
}

func (s *CreateSslVpnServerRequest) SetClientIpPool(v string) *CreateSslVpnServerRequest {
	s.ClientIpPool = &v
	return s
}

func (s *CreateSslVpnServerRequest) SetClientToken(v string) *CreateSslVpnServerRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSslVpnServerRequest) SetCompress(v bool) *CreateSslVpnServerRequest {
	s.Compress = &v
	return s
}

func (s *CreateSslVpnServerRequest) SetDnsServers(v string) *CreateSslVpnServerRequest {
	s.DnsServers = &v
	return s
}

func (s *CreateSslVpnServerRequest) SetDryRun(v bool) *CreateSslVpnServerRequest {
	s.DryRun = &v
	return s
}

func (s *CreateSslVpnServerRequest) SetEnableMultiFactorAuth(v bool) *CreateSslVpnServerRequest {
	s.EnableMultiFactorAuth = &v
	return s
}

func (s *CreateSslVpnServerRequest) SetIDaaSApplicationId(v string) *CreateSslVpnServerRequest {
	s.IDaaSApplicationId = &v
	return s
}

func (s *CreateSslVpnServerRequest) SetIDaaSInstanceId(v string) *CreateSslVpnServerRequest {
	s.IDaaSInstanceId = &v
	return s
}

func (s *CreateSslVpnServerRequest) SetIDaaSRegionId(v string) *CreateSslVpnServerRequest {
	s.IDaaSRegionId = &v
	return s
}

func (s *CreateSslVpnServerRequest) SetLocalSubnet(v string) *CreateSslVpnServerRequest {
	s.LocalSubnet = &v
	return s
}

func (s *CreateSslVpnServerRequest) SetName(v string) *CreateSslVpnServerRequest {
	s.Name = &v
	return s
}

func (s *CreateSslVpnServerRequest) SetOwnerAccount(v string) *CreateSslVpnServerRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateSslVpnServerRequest) SetOwnerId(v int64) *CreateSslVpnServerRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSslVpnServerRequest) SetPort(v int32) *CreateSslVpnServerRequest {
	s.Port = &v
	return s
}

func (s *CreateSslVpnServerRequest) SetProto(v string) *CreateSslVpnServerRequest {
	s.Proto = &v
	return s
}

func (s *CreateSslVpnServerRequest) SetRegionId(v string) *CreateSslVpnServerRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSslVpnServerRequest) SetResourceOwnerAccount(v string) *CreateSslVpnServerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSslVpnServerRequest) SetResourceOwnerId(v int64) *CreateSslVpnServerRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSslVpnServerRequest) SetVpnGatewayId(v string) *CreateSslVpnServerRequest {
	s.VpnGatewayId = &v
	return s
}

func (s *CreateSslVpnServerRequest) Validate() error {
	return dara.Validate(s)
}
