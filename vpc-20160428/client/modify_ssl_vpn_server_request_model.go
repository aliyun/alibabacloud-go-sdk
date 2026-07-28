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
	SetDnsServers(v string) *ModifySslVpnServerRequest
	GetDnsServers() *string
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
	// The encryption algorithm used by the SSL-VPN server. Valid values:
	//
	// - **AES-128-CBC*	- (default): AES-128-CBC algorithm.
	//
	// - **AES-192-CBC**: AES-192-CBC algorithm.
	//
	// - **AES-256-CBC**: AES-256-CBC algorithm.
	//
	// - **none**: no encryption algorithm is used.
	//
	// example:
	//
	// AES-128-CBC
	Cipher *string `json:"Cipher,omitempty" xml:"Cipher,omitempty"`
	// The client CIDR block.
	//
	// The client CIDR block is the CIDR block from which IP addresses are assigned to virtual network interface controllers (NICs) of clients. It is not the internal network CIDR block of the client.
	//
	// When a client accesses the local virtual private cloud (VPC) through an SSL-VPN connection, the VPN gateway assigns an IP address from the specified client CIDR block to the client. The client uses the assigned IP address to access cloud resources.
	//
	// When you specify the client CIDR block, make sure that the number of IP addresses in the client CIDR block is at least four times the number of SSL connections supported by the VPN gateway.
	//
	// <details>
	//
	// <summary>Click to view the reason.</summary>
	//
	// For example, if you set 192.168.0.0/24 as the client CIDR block, the system first allocates a subnet with a 30-bit subnet mask from the 192.168.0.0/24 CIDR block, such as 192.168.0.4/30, and then assigns one IP address from 192.168.0.4/30 to the client. The remaining three IP addresses are reserved by the system to ensure network communication. In this case, one client consumes four IP addresses. Therefore, to ensure that all clients can be assigned IP addresses, make sure that the number of IP addresses in the client CIDR block is at least four times the number of SSL connections supported by the VPN gateway.
	//
	// </details>
	//
	// <details>
	//
	// <summary>Click to view unsupported CIDR blocks.</summary>
	//
	// - 100.64.0.0 to 100.127.255.255
	//
	// - 127.0.0.0 to 127.255.255.255
	//
	// - 169.254.0.0 to 169.254.255.255
	//
	// - 224.0.0.0 to 239.255.255.255
	//
	// - 255.0.0.0 to 255.255.255.255
	//
	// </details>
	//
	// <details>
	//
	// <summary>Click to view the recommended client CIDR blocks for each number of SSL connections.</summary>
	//
	// - If the number of SSL connections is 5, the subnet mask of the client CIDR block must be 27 bits or less. For example, 10.0.0.0/27 or 10.0.0.0/26.
	//
	// - If the number of SSL connections is 10, the subnet mask of the client CIDR block must be 26 bits or less. For example, 10.0.0.0/26 or 10.0.0.0/25.
	//
	// - If the number of SSL connections is 20, the subnet mask of the client CIDR block must be 25 bits or less. For example, 10.0.0.0/25 or 10.0.0.0/24.
	//
	// - If the number of SSL connections is 50, the subnet mask of the client CIDR block must be 24 bits or less. For example, 10.0.0.0/24 or 10.0.0.0/23.
	//
	// - If the number of SSL connections is 100, the subnet mask of the client CIDR block must be 23 bits or less. For example, 10.0.0.0/23 or 10.0.0.0/22.
	//
	// - If the number of SSL connections is 200, the subnet mask of the client CIDR block must be 22 bits or less. For example, 10.0.0.0/22 or 10.0.0.0/21.
	//
	// - If the number of SSL connections is 500, the subnet mask of the client CIDR block must be 21 bits or less. For example, 10.0.0.0/21 or 10.0.0.0/20.
	//
	// - If the number of SSL connections is 1000, the subnet mask of the client CIDR block must be 20 bits or less. For example, 10.0.0.0/20 or 10.0.0.0/19.
	//
	// </details>
	//
	// > - The subnet mask of the client CIDR block must be 16 to 29 bits in length.
	//
	// > - Make sure that the client CIDR block does not overlap with the local CIDR block, the VPC CIDR block, or any routing CIDR block associated with the client terminal.
	//
	// > - Use 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16, or their subnets as the client CIDR block. If you want to specify a public CIDR block as the client CIDR block, set the public CIDR block as a user CIDR block of the VPC to ensure that the VPC can access the public CIDR block. For more information about user CIDR blocks, see [VPC FAQ](https://help.aliyun.com/document_detail/185311.html).
	//
	// > - After the SSL server is created, the system automatically adds the routing of the client CIDR block to the route table of the VPC-connected instance. Do not manually add the routing of the client CIDR block to the route table of the VPC-connected instance. Otherwise, SSL-VPN connection traffic may be abnormal.
	//
	// example:
	//
	// 10.30.30.0/24
	ClientIpPool *string `json:"ClientIpPool,omitempty" xml:"ClientIpPool,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-0016e04115b
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to compress communication. Valid values:
	//
	// - **true*	- (default): Communication is compressed.
	//
	// - **false**: Communication is not compressed.
	//
	// example:
	//
	// true
	Compress   *bool   `json:"Compress,omitempty" xml:"Compress,omitempty"`
	DnsServers *string `json:"DnsServers,omitempty" xml:"DnsServers,omitempty"`
	// Specifies whether to perform a dry run, without performing the actual request. Valid values:
	//
	// - **true**: sends a check request without modifying the SSL server configuration. The check items include whether required parameters are specified, the request format, and service limits. If the check fails, the corresponding error is returned. If the check passes, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends a Normal request. If the request passes the check, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to enable two-factor identity authentication. If you enable two-factor identity authentication, you must also configure **IDaaSInstanceId**, **IDaaSRegionId**, and **IDaaSApplicationId**. Valid values:
	//
	// - **true**: enabled.
	//
	// - **false**: not enabled.
	//
	// > - If you use two-factor identity authentication for the first time, complete [authorization](https://ram.console.aliyun.com/role/authorization?request=%7B%22Services%22%3A%5B%7B%22Service%22%3A%22VPN%22%2C%22Roles%22%3A%5B%7B%22RoleName%22%3A%22AliyunVpnAccessingIdaasRole%22%2C%22TemplateId%22%3A%22IdaasRole%22%7D%5D%7D%5D%2C%22ReturnUrl%22%3A%22https%3A%2F%2Fvpc.console.aliyun.com%2Fsslvpn%2Fcn-shanghai%2Fvpn-servers%22%7D) before creating the SSL server.
	//
	// > - When you create an SSL server in the UAE (Dubai) region, bind an IDaaS EIAM 2.0 instance in the Singapore region to reduce cross-region latency.
	//
	// > - IDaaS EIAM 1.0 instances are no longer available for purchase. If your Alibaba Cloud account has existing IDaaS EIAM 1.0 instances, you can still bind IDaaS EIAM 1.0 instances after enabling two-factor identity authentication. If your Alibaba Cloud account does not have IDaaS EIAM 1.0 instances, you can bind only IDaaS EIAM 2.0 instances after enabling two-factor identity authentication.
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
	// The instance ID of the IDaaS EIAM instance.
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
	// The local CIDR block is the CIDR block that the client needs to access through the SSL-VPN connection.
	//
	// The local CIDR block can be the CIDR block of a VPC, the CIDR block of a vSwitch, the CIDR block of an on-premises data center that is connected to the VPC through an Express Connect circuit, or the CIDR block of a cloud service such as Object Storage Service (OSS).
	//
	// The subnet mask of the local CIDR block must be 8 to 32 bits in length. The following CIDR blocks cannot be specified as the local CIDR block:
	//
	// - 127.0.0.0 to 127.255.255.255
	//
	// - 169.254.0.0 to 169.254.255.255
	//
	// - 224.0.0.0 to 239.255.255.255
	//
	// - 255.0.0.0 to 255.255.255.255
	//
	// example:
	//
	// 10.20.20.0/24
	LocalSubnet *string `json:"LocalSubnet,omitempty" xml:"LocalSubnet,omitempty"`
	// The name of the SSL-VPN server.
	//
	// The name must be 1 to 100 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
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
	// The instance ID of the SSL-VPN server.
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

func (s *ModifySslVpnServerRequest) GetDnsServers() *string {
	return s.DnsServers
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

func (s *ModifySslVpnServerRequest) SetDnsServers(v string) *ModifySslVpnServerRequest {
	s.DnsServers = &v
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
