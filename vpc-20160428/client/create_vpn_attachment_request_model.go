// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpnAttachmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoConfigRoute(v bool) *CreateVpnAttachmentRequest
	GetAutoConfigRoute() *bool
	SetBgpConfig(v string) *CreateVpnAttachmentRequest
	GetBgpConfig() *string
	SetClientToken(v string) *CreateVpnAttachmentRequest
	GetClientToken() *string
	SetCustomerGatewayId(v string) *CreateVpnAttachmentRequest
	GetCustomerGatewayId() *string
	SetDryRun(v bool) *CreateVpnAttachmentRequest
	GetDryRun() *bool
	SetEffectImmediately(v bool) *CreateVpnAttachmentRequest
	GetEffectImmediately() *bool
	SetEnableDpd(v bool) *CreateVpnAttachmentRequest
	GetEnableDpd() *bool
	SetEnableNatTraversal(v bool) *CreateVpnAttachmentRequest
	GetEnableNatTraversal() *bool
	SetEnableTunnelsBgp(v bool) *CreateVpnAttachmentRequest
	GetEnableTunnelsBgp() *bool
	SetHealthCheckConfig(v string) *CreateVpnAttachmentRequest
	GetHealthCheckConfig() *string
	SetIkeConfig(v string) *CreateVpnAttachmentRequest
	GetIkeConfig() *string
	SetIpsecConfig(v string) *CreateVpnAttachmentRequest
	GetIpsecConfig() *string
	SetLocalSubnet(v string) *CreateVpnAttachmentRequest
	GetLocalSubnet() *string
	SetName(v string) *CreateVpnAttachmentRequest
	GetName() *string
	SetNetworkType(v string) *CreateVpnAttachmentRequest
	GetNetworkType() *string
	SetOwnerAccount(v string) *CreateVpnAttachmentRequest
	GetOwnerAccount() *string
	SetRegionId(v string) *CreateVpnAttachmentRequest
	GetRegionId() *string
	SetRemoteCaCert(v string) *CreateVpnAttachmentRequest
	GetRemoteCaCert() *string
	SetRemoteSubnet(v string) *CreateVpnAttachmentRequest
	GetRemoteSubnet() *string
	SetResourceGroupId(v string) *CreateVpnAttachmentRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateVpnAttachmentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateVpnAttachmentRequest
	GetResourceOwnerId() *int64
	SetTags(v []*CreateVpnAttachmentRequestTags) *CreateVpnAttachmentRequest
	GetTags() []*CreateVpnAttachmentRequestTags
	SetTunnelBandwidth(v string) *CreateVpnAttachmentRequest
	GetTunnelBandwidth() *string
	SetTunnelOptionsSpecification(v []*CreateVpnAttachmentRequestTunnelOptionsSpecification) *CreateVpnAttachmentRequest
	GetTunnelOptionsSpecification() []*CreateVpnAttachmentRequestTunnelOptionsSpecification
}

type CreateVpnAttachmentRequest struct {
	// Specifies whether to automatically configure routes. Valid values:
	//
	// - **true*	- (default): Automatically configure routes.
	//
	// - **false**: Do not automatically configure routes.
	//
	// example:
	//
	// true
	AutoConfigRoute *bool `json:"AutoConfigRoute,omitempty" xml:"AutoConfigRoute,omitempty"`
	// This parameter is supported when you create a single-tunnel IPsec-VPN connection.
	//
	// BGP configuration:
	//
	// - **BgpConfig.EnableBgp**: Specifies whether to enable the BGP feature. Valid values: **true*	- or **false*	- (default).
	//
	// - **BgpConfig.LocalAsn**: The autonomous system number on the Alibaba Cloud side. Valid values: **1*	- to **4294967295**. Default value: **45104**.
	//
	//     You can enter the autonomous system number in the two-segment format: the first 16 bits.the last 16 bits. Enter each segment in decimal format.
	//
	//     For example, if you enter 123.456, the autonomous system number is 123 × 65536 + 456 = 8061384.
	//
	// - **BgpConfig.TunnelCidr**: The IPsec tunnel CIDR block. The CIDR block must fall within 169.254.0.0/16 and have a mask length of 30 bits. The CIDR block cannot be 169.254.0.0/30, 169.254.1.0/30, 169.254.2.0/30, 169.254.3.0/30, 169.254.4.0/30, 169.254.5.0/30, 169.254.6.0/30, or 169.254.169.252/30.
	//
	// - **LocalBgpIp**: The BGP address on the Alibaba Cloud side. This address is an IP address within the IPsec tunnel CIDR block.
	//
	// > - Before you add BGP configurations, learn about how the dynamic routing feature works and its limits. For more information, see [Configure BGP dynamic routing](https://help.aliyun.com/document_detail/445767.html).
	//
	// > - Use a private autonomous system number to establish a BGP connection with Alibaba Cloud. Refer to the relevant documentation for the range of private autonomous system numbers.
	//
	// example:
	//
	// {"EnableBgp":"true","LocalAsn":"45104","TunnelCidr":"169.254.11.0/30","LocalBgpIp":"169.254.11.1"}
	BgpConfig *string `json:"BgpConfig,omitempty" xml:"BgpConfig,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-4266****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The customer gateway ID.
	//
	// > This parameter is required only when you create a single-tunnel IPsec-VPN connection.
	//
	// example:
	//
	// cgw-p0w2jemrcj5u61un8****
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" xml:"CustomerGatewayId,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without creating the IPsec-VPN connection. The system checks the required parameters, request format, and service limits. If the check fails, the corresponding error is returned. If the check succeeds, the error code `DryRunOperation` is returned.
	//
	// - **false*	- (default): performs a dry run and sends the request. If the check succeeds, the IPsec-VPN connection is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether the configuration takes effect immediately. Valid values:
	//
	// - **true**: The system initiates IPsec protocol negotiation immediately after the configuration is complete.
	//
	// - **false*	- (default): The system initiates IPsec protocol negotiation only when traffic enters the IPsec-VPN connection.
	//
	// example:
	//
	// false
	EffectImmediately *bool `json:"EffectImmediately,omitempty" xml:"EffectImmediately,omitempty"`
	// This parameter is supported when you create a single-tunnel IPsec-VPN connection.
	//
	// Specifies whether to enable the Dead Peer Detection (DPD) feature. Valid values:
	//
	// - **true*	- (default): Enables the DPD feature. The IPsec initiator sends DPD packets to check the existence and availability of the peer. If no correct response is received within the specified time, the connection fails. Then, the ISAKMP SA, IPsec SA, and IPsec tunnel are deleted.
	//
	// - **false**: Disables the DPD feature. The IPsec initiator does not send DPD packets.
	//
	// example:
	//
	// true
	EnableDpd *bool `json:"EnableDpd,omitempty" xml:"EnableDpd,omitempty"`
	// This parameter is supported when you create a single-tunnel IPsec-VPN connection.
	//
	// Specifies whether to enable NAT traversal. Valid values:
	//
	// - **true*	- (default): Enables NAT traversal. After NAT traversal is enabled, the verification of the UDP port number is removed during IKE negotiation, and NAT gateway devices along the VPN tunnel can be discovered.
	//
	// - **false**: Disables NAT traversal.
	//
	// example:
	//
	// true
	EnableNatTraversal *bool `json:"EnableNatTraversal,omitempty" xml:"EnableNatTraversal,omitempty"`
	// This parameter is supported when you create a dual-tunnel IPsec-VPN connection.
	//
	// Specifies whether to enable the BGP dynamic routing feature for the tunnels. Valid values: **true*	- or **false*	- (default).
	//
	// > Before you add BGP configurations, learn about how the BGP dynamic routing feature works and its limits. For more information, see [Configure BGP dynamic routing](https://help.aliyun.com/document_detail/445767.html).
	//
	// example:
	//
	// false
	EnableTunnelsBgp *bool `json:"EnableTunnelsBgp,omitempty" xml:"EnableTunnelsBgp,omitempty"`
	// This parameter is supported when you create a single-tunnel IPsec-VPN connection.
	//
	// Health check configuration:
	//
	// - **HealthCheckConfig.enable**: Specifies whether to enable health checks. Valid values: **true*	- or **false*	- (default).
	//
	// - **HealthCheckConfig.dip**: The destination IP address of the health check. Enter an IP address of the on-premises data center that can be accessed from the VPC side through the IPsec-VPN connection.
	//
	// - **HealthCheckConfig.sip**: The source IP address of the health check. Enter an IP address on the VPC side that can be accessed from the on-premises data center through the IPsec-VPN connection.
	//
	// - **HealthCheckConfig.interval**: The retry interval of the health check. Unit: seconds. Default value: **3**.
	//
	// - **HealthCheckConfig.retry**: The number of retry packets sent during the health check. Default value: **3**.
	//
	// - **HealthCheckConfig.Policy**: Specifies whether to withdraw published routes when the health check fails. Valid values:
	//
	//      - **revoke_route*	- (default): Withdraw published routes.
	//
	//      - **reserve_route**: Do not withdraw published routes.
	//
	// example:
	//
	// {"enable":"true","dip":"192.168.1.1","sip":"10.1.1.1","interval":"3","retry":"3","Policy": "revoke_route"}
	HealthCheckConfig *string `json:"HealthCheckConfig,omitempty" xml:"HealthCheckConfig,omitempty"`
	// This parameter is supported when you create a single-tunnel IPsec-VPN connection.
	//
	// Phase 1 negotiation configuration:
	//
	//
	//
	// - **IkeConfig.Psk**: The pre-shared key used for identity authentication between the VPN gateway and the on-premises data center.
	//
	//     - The key must be 1 to 100 characters in length and can contain digits, uppercase and lowercase letters, and the following characters. It cannot contain spaces. ```~!`@#$%^&*()_-+={}[]|;:\\",.<>/?```
	//
	//     - If you do not specify a pre-shared key, the system randomly generates a string as the pre-shared key. You can call the [DescribeVpnConnection](https://help.aliyun.com/document_detail/2526951.html) operation to query the pre-shared key that is automatically generated by the system.
	//
	//     > The pre-shared key on the IPsec-VPN connection side must be the same as the authentication key on the on-premises data center side. Otherwise, the connection between the on-premises data center and the VPN gateway cannot be established.
	//
	// - **IkeConfig.IkeVersion**: The version of the IKE protocol. Valid values: **ikev1*	- or **ikev2**. Default value: **ikev1**.
	//
	// - **IkeConfig.IkeMode**: The negotiation mode. Valid values: **main*	- or **aggressive**. Default value: **main**.
	//
	// - **IkeConfig.IkeEncAlg**: The encryption algorithm used in Phase 1 negotiation. Valid values: **aes**, **aes192**, **aes256**, **des**, or **3des**. Default value: **aes**.
	//
	// - **IkeConfig.IkeAuthAlg**: The authentication algorithm used in Phase 1 negotiation. Valid values: **md5**, **sha1**, **sha256**, **sha384**, **sha512**. Default value: **md5**.
	//
	// - **IkeConfig.IkePfs**: The Diffie-Hellman key exchange algorithm used in Phase 1 negotiation. Valid values: **group1**, **group2**, **group5**, or **group14**. Default value: **group2**.
	//
	// - **IkeConfig.IkeLifetime**: The lifetime of the SA generated in Phase 1 negotiation. Unit: seconds. Valid values: **0*	- to **86400**. Default value: **86400**.
	//
	// - **IkeConfig.LocalId**: The identifier on the Alibaba Cloud side of the IPsec-VPN connection. The identifier is limited to 100 characters in length and cannot contain spaces. The default value is empty.
	//
	// - **IkeConfig.RemoteId**: The identifier on the on-premises data center side of the IPsec-VPN connection. The identifier is limited to 100 characters in length and cannot contain spaces. The default value is the IP address of the customer gateway.
	//
	// example:
	//
	// {"Psk":"1234****","IkeVersion":"ikev1","IkeMode":"main","IkeEncAlg":"aes","IkeAuthAlg":"sha1","IkePfs":"group2","IkeLifetime":86400,"LocalId":"47.XX.XX.1","RemoteId":"47.XX.XX.2"}
	IkeConfig *string `json:"IkeConfig,omitempty" xml:"IkeConfig,omitempty"`
	// This parameter is supported when you create a single-tunnel IPsec-VPN connection.
	//
	// Phase 2 negotiation configuration:
	//
	// - **IpsecConfig.IpsecEncAlg**: The encryption algorithm used in Phase 2 negotiation. Valid values: **aes**, **aes192**, **aes256**, **des**, or **3des**. Default value: **aes**.
	//
	// - **IpsecConfig.IpsecAuthAlg**: The authentication algorithm used in Phase 2 negotiation. Valid values: **md5**, **sha1**, **sha256**, **sha384**, **sha512**. Default value: **md5**.
	//
	// - **IpsecConfig.IpsecPfs**: The Diffie-Hellman key exchange algorithm used in Phase 2 negotiation. Valid values: **disabled**, **group1**, **group2**, **group5**, or **group14**. Default value: **group2**.
	//
	// - **IpsecConfig.IpsecLifetime**: The lifetime of the SA generated in Phase 2 negotiation. Unit: seconds. Valid values: **0*	- to **86400**. Default value: **86400**.
	//
	// example:
	//
	// {"IpsecEncAlg":"aes","IpsecAuthAlg":"sha1","IpsecPfs":"group2","IpsecLifetime":86400}
	IpsecConfig *string `json:"IpsecConfig,omitempty" xml:"IpsecConfig,omitempty"`
	// The CIDR block on the VPC side that needs to communicate with the on-premises data center. This CIDR block is used in Phase 2 negotiation.
	//
	// Separate multiple CIDR blocks with commas (,). Example: 192.168.1.0/24,192.168.2.0/24.
	//
	// Notes on the routing mode of the IPsec-VPN connection:
	//
	// - If both **LocalSubnet*	- and **RemoteSubnet*	- are set to 0.0.0.0/0, the destination routing mode is used.
	//
	// - If both **LocalSubnet*	- and **RemoteSubnet*	- are set to specific CIDR blocks, the policy-based routing mode is used.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.1.1.0/24,10.1.2.0/24
	LocalSubnet *string `json:"LocalSubnet,omitempty" xml:"LocalSubnet,omitempty"`
	// The name of the IPsec-VPN connection.
	//
	// The name must be 1 to 100 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// nametest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The network type of the IPsec-VPN connection. Valid values:
	//
	// - **public*	- (default): public network. The IPsec-VPN connection establishes an encrypted communication channel over the Internet.
	//
	// - **private**: private network. The IPsec-VPN connection establishes an encrypted communication channel over a private network.
	//
	// example:
	//
	// public
	NetworkType  *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	// The region ID of the IPsec-VPN connection.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The CA certificate of the peer.
	//
	// > This parameter is not in effect.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- MIIB7zCCAZW***	- -----END CERTIFICATE-----
	RemoteCaCert *string `json:"RemoteCaCert,omitempty" xml:"RemoteCaCert,omitempty"`
	// The CIDR block on the on-premises data center side that needs to communicate with the VPC. This CIDR block is used in Phase 2 negotiation.
	//
	// Separate multiple CIDR blocks with commas (,). Example: 192.168.3.0/24,192.168.4.0/24.
	//
	// Notes on the routing mode of the IPsec-VPN connection:
	//
	// - If both **LocalSubnet*	- and **RemoteSubnet*	- are set to 0.0.0.0/0, the destination routing mode is used.
	//
	// - If both **LocalSubnet*	- and **RemoteSubnet*	- are set to specific CIDR blocks, the policy-based routing mode is used.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.1.3.0/24,10.1.4.0/24
	RemoteSubnet *string `json:"RemoteSubnet,omitempty" xml:"RemoteSubnet,omitempty"`
	// The ID of the resource group to which the IPsec-VPN connection belongs.
	//
	// - You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation to query the resource group ID.
	//
	// - If you do not specify a resource group ID, the IPsec-VPN connection belongs to the default resource group after it is created.
	//
	// example:
	//
	// rg-acfmzs372yg****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The list of tags to add to the IPsec-VPN connection.
	//
	// You can add up to 20 tags to an IPsec-VPN connection at a time.
	Tags []*CreateVpnAttachmentRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The bandwidth specification of a single VPN tunnel. Valid values:
	//
	// - Standard (default): standard. The default bandwidth is 1 Gbps.
	//
	// - Large: large. The default bandwidth is 3 Gbps.
	//
	// example:
	//
	// Standard
	TunnelBandwidth *string `json:"TunnelBandwidth,omitempty" xml:"TunnelBandwidth,omitempty"`
	// The tunnel configurations.
	//
	// - The parameters in the **TunnelOptionsSpecification*	- array are supported when you create a dual-tunnel IPsec-VPN connection.
	//
	// - When you create a dual-tunnel IPsec-VPN connection, you must add two tunnels at the same time to ensure link redundancy for the IPsec-VPN connection. Only two tunnels can be added to an IPsec-VPN connection.
	//
	// if can be null:
	// true
	TunnelOptionsSpecification []*CreateVpnAttachmentRequestTunnelOptionsSpecification `json:"TunnelOptionsSpecification,omitempty" xml:"TunnelOptionsSpecification,omitempty" type:"Repeated"`
}

func (s CreateVpnAttachmentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVpnAttachmentRequest) GoString() string {
	return s.String()
}

func (s *CreateVpnAttachmentRequest) GetAutoConfigRoute() *bool {
	return s.AutoConfigRoute
}

func (s *CreateVpnAttachmentRequest) GetBgpConfig() *string {
	return s.BgpConfig
}

func (s *CreateVpnAttachmentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateVpnAttachmentRequest) GetCustomerGatewayId() *string {
	return s.CustomerGatewayId
}

func (s *CreateVpnAttachmentRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateVpnAttachmentRequest) GetEffectImmediately() *bool {
	return s.EffectImmediately
}

func (s *CreateVpnAttachmentRequest) GetEnableDpd() *bool {
	return s.EnableDpd
}

func (s *CreateVpnAttachmentRequest) GetEnableNatTraversal() *bool {
	return s.EnableNatTraversal
}

func (s *CreateVpnAttachmentRequest) GetEnableTunnelsBgp() *bool {
	return s.EnableTunnelsBgp
}

func (s *CreateVpnAttachmentRequest) GetHealthCheckConfig() *string {
	return s.HealthCheckConfig
}

func (s *CreateVpnAttachmentRequest) GetIkeConfig() *string {
	return s.IkeConfig
}

func (s *CreateVpnAttachmentRequest) GetIpsecConfig() *string {
	return s.IpsecConfig
}

func (s *CreateVpnAttachmentRequest) GetLocalSubnet() *string {
	return s.LocalSubnet
}

func (s *CreateVpnAttachmentRequest) GetName() *string {
	return s.Name
}

func (s *CreateVpnAttachmentRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateVpnAttachmentRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateVpnAttachmentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateVpnAttachmentRequest) GetRemoteCaCert() *string {
	return s.RemoteCaCert
}

func (s *CreateVpnAttachmentRequest) GetRemoteSubnet() *string {
	return s.RemoteSubnet
}

func (s *CreateVpnAttachmentRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateVpnAttachmentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateVpnAttachmentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateVpnAttachmentRequest) GetTags() []*CreateVpnAttachmentRequestTags {
	return s.Tags
}

func (s *CreateVpnAttachmentRequest) GetTunnelBandwidth() *string {
	return s.TunnelBandwidth
}

func (s *CreateVpnAttachmentRequest) GetTunnelOptionsSpecification() []*CreateVpnAttachmentRequestTunnelOptionsSpecification {
	return s.TunnelOptionsSpecification
}

func (s *CreateVpnAttachmentRequest) SetAutoConfigRoute(v bool) *CreateVpnAttachmentRequest {
	s.AutoConfigRoute = &v
	return s
}

func (s *CreateVpnAttachmentRequest) SetBgpConfig(v string) *CreateVpnAttachmentRequest {
	s.BgpConfig = &v
	return s
}

func (s *CreateVpnAttachmentRequest) SetClientToken(v string) *CreateVpnAttachmentRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateVpnAttachmentRequest) SetCustomerGatewayId(v string) *CreateVpnAttachmentRequest {
	s.CustomerGatewayId = &v
	return s
}

func (s *CreateVpnAttachmentRequest) SetDryRun(v bool) *CreateVpnAttachmentRequest {
	s.DryRun = &v
	return s
}

func (s *CreateVpnAttachmentRequest) SetEffectImmediately(v bool) *CreateVpnAttachmentRequest {
	s.EffectImmediately = &v
	return s
}

func (s *CreateVpnAttachmentRequest) SetEnableDpd(v bool) *CreateVpnAttachmentRequest {
	s.EnableDpd = &v
	return s
}

func (s *CreateVpnAttachmentRequest) SetEnableNatTraversal(v bool) *CreateVpnAttachmentRequest {
	s.EnableNatTraversal = &v
	return s
}

func (s *CreateVpnAttachmentRequest) SetEnableTunnelsBgp(v bool) *CreateVpnAttachmentRequest {
	s.EnableTunnelsBgp = &v
	return s
}

func (s *CreateVpnAttachmentRequest) SetHealthCheckConfig(v string) *CreateVpnAttachmentRequest {
	s.HealthCheckConfig = &v
	return s
}

func (s *CreateVpnAttachmentRequest) SetIkeConfig(v string) *CreateVpnAttachmentRequest {
	s.IkeConfig = &v
	return s
}

func (s *CreateVpnAttachmentRequest) SetIpsecConfig(v string) *CreateVpnAttachmentRequest {
	s.IpsecConfig = &v
	return s
}

func (s *CreateVpnAttachmentRequest) SetLocalSubnet(v string) *CreateVpnAttachmentRequest {
	s.LocalSubnet = &v
	return s
}

func (s *CreateVpnAttachmentRequest) SetName(v string) *CreateVpnAttachmentRequest {
	s.Name = &v
	return s
}

func (s *CreateVpnAttachmentRequest) SetNetworkType(v string) *CreateVpnAttachmentRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateVpnAttachmentRequest) SetOwnerAccount(v string) *CreateVpnAttachmentRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateVpnAttachmentRequest) SetRegionId(v string) *CreateVpnAttachmentRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVpnAttachmentRequest) SetRemoteCaCert(v string) *CreateVpnAttachmentRequest {
	s.RemoteCaCert = &v
	return s
}

func (s *CreateVpnAttachmentRequest) SetRemoteSubnet(v string) *CreateVpnAttachmentRequest {
	s.RemoteSubnet = &v
	return s
}

func (s *CreateVpnAttachmentRequest) SetResourceGroupId(v string) *CreateVpnAttachmentRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateVpnAttachmentRequest) SetResourceOwnerAccount(v string) *CreateVpnAttachmentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateVpnAttachmentRequest) SetResourceOwnerId(v int64) *CreateVpnAttachmentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateVpnAttachmentRequest) SetTags(v []*CreateVpnAttachmentRequestTags) *CreateVpnAttachmentRequest {
	s.Tags = v
	return s
}

func (s *CreateVpnAttachmentRequest) SetTunnelBandwidth(v string) *CreateVpnAttachmentRequest {
	s.TunnelBandwidth = &v
	return s
}

func (s *CreateVpnAttachmentRequest) SetTunnelOptionsSpecification(v []*CreateVpnAttachmentRequestTunnelOptionsSpecification) *CreateVpnAttachmentRequest {
	s.TunnelOptionsSpecification = v
	return s
}

func (s *CreateVpnAttachmentRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TunnelOptionsSpecification != nil {
		for _, item := range s.TunnelOptionsSpecification {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateVpnAttachmentRequestTags struct {
	// The tag key. Once specified, the tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// You can specify up to 20 tag keys at a time.
	//
	// example:
	//
	// TagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// The tag value can be up to 128 characters in length and can be an empty string. It cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// Each tag key corresponds to one tag value. You can specify up to 20 tag values at a time.
	//
	// example:
	//
	// TagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateVpnAttachmentRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateVpnAttachmentRequestTags) GoString() string {
	return s.String()
}

func (s *CreateVpnAttachmentRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateVpnAttachmentRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateVpnAttachmentRequestTags) SetKey(v string) *CreateVpnAttachmentRequestTags {
	s.Key = &v
	return s
}

func (s *CreateVpnAttachmentRequestTags) SetValue(v string) *CreateVpnAttachmentRequestTags {
	s.Value = &v
	return s
}

func (s *CreateVpnAttachmentRequestTags) Validate() error {
	return dara.Validate(s)
}

type CreateVpnAttachmentRequestTunnelOptionsSpecification struct {
	// The customer gateway ID associated with the tunnel.
	//
	// > This parameter is required when you create a dual-tunnel IPsec-VPN connection.
	//
	// example:
	//
	// cgw-p0w2jemrcj5u61un8****
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" xml:"CustomerGatewayId,omitempty"`
	// Specifies whether to enable the Dead Peer Detection (DPD) feature for the tunnel. Valid values:
	//
	// - **true*	- (default): Enables the DPD feature. The IPsec initiator sends DPD packets to check the existence and availability of the peer. If no correct response is received within the specified time, the connection fails. Then, the ISAKMP SA, IPsec SA, and IPsec tunnel are deleted.
	//
	// - **false**: Disables the DPD feature. The IPsec initiator does not send DPD packets.
	//
	// example:
	//
	// true
	EnableDpd *bool `json:"EnableDpd,omitempty" xml:"EnableDpd,omitempty"`
	// Specifies whether to enable NAT traversal for the tunnel. Valid values:
	//
	// - **true*	- (default): Enables NAT traversal. After NAT traversal is enabled, the verification of the UDP port number is removed during IKE negotiation, and NAT gateway devices along the tunnel can be discovered.
	//
	// - **false**: Disables NAT traversal.
	//
	// example:
	//
	// true
	EnableNatTraversal *bool `json:"EnableNatTraversal,omitempty" xml:"EnableNatTraversal,omitempty"`
	// The BGP configuration for the tunnel.
	//
	// > This parameter is required when you enable BGP for the IPsec-VPN connection (set **EnableTunnelsBgp*	- to **true**).
	TunnelBgpConfig *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelBgpConfig `json:"TunnelBgpConfig,omitempty" xml:"TunnelBgpConfig,omitempty" type:"Struct"`
	// The Phase 1 negotiation configuration.
	TunnelIkeConfig *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIkeConfig `json:"TunnelIkeConfig,omitempty" xml:"TunnelIkeConfig,omitempty" type:"Struct"`
	// The order in which the tunnel is created.
	//
	// - **1**: the first tunnel.
	//
	// - **2**: the second tunnel.
	//
	// example:
	//
	// 1
	TunnelIndex *int32 `json:"TunnelIndex,omitempty" xml:"TunnelIndex,omitempty"`
	// The Phase 2 negotiation configuration.
	TunnelIpsecConfig *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIpsecConfig `json:"TunnelIpsecConfig,omitempty" xml:"TunnelIpsecConfig,omitempty" type:"Struct"`
}

func (s CreateVpnAttachmentRequestTunnelOptionsSpecification) String() string {
	return dara.Prettify(s)
}

func (s CreateVpnAttachmentRequestTunnelOptionsSpecification) GoString() string {
	return s.String()
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecification) GetCustomerGatewayId() *string {
	return s.CustomerGatewayId
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecification) GetEnableDpd() *bool {
	return s.EnableDpd
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecification) GetEnableNatTraversal() *bool {
	return s.EnableNatTraversal
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecification) GetTunnelBgpConfig() *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelBgpConfig {
	return s.TunnelBgpConfig
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecification) GetTunnelIkeConfig() *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIkeConfig {
	return s.TunnelIkeConfig
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecification) GetTunnelIndex() *int32 {
	return s.TunnelIndex
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecification) GetTunnelIpsecConfig() *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIpsecConfig {
	return s.TunnelIpsecConfig
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecification) SetCustomerGatewayId(v string) *CreateVpnAttachmentRequestTunnelOptionsSpecification {
	s.CustomerGatewayId = &v
	return s
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecification) SetEnableDpd(v bool) *CreateVpnAttachmentRequestTunnelOptionsSpecification {
	s.EnableDpd = &v
	return s
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecification) SetEnableNatTraversal(v bool) *CreateVpnAttachmentRequestTunnelOptionsSpecification {
	s.EnableNatTraversal = &v
	return s
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecification) SetTunnelBgpConfig(v *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelBgpConfig) *CreateVpnAttachmentRequestTunnelOptionsSpecification {
	s.TunnelBgpConfig = v
	return s
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecification) SetTunnelIkeConfig(v *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIkeConfig) *CreateVpnAttachmentRequestTunnelOptionsSpecification {
	s.TunnelIkeConfig = v
	return s
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecification) SetTunnelIndex(v int32) *CreateVpnAttachmentRequestTunnelOptionsSpecification {
	s.TunnelIndex = &v
	return s
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecification) SetTunnelIpsecConfig(v *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIpsecConfig) *CreateVpnAttachmentRequestTunnelOptionsSpecification {
	s.TunnelIpsecConfig = v
	return s
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecification) Validate() error {
	if s.TunnelBgpConfig != nil {
		if err := s.TunnelBgpConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TunnelIkeConfig != nil {
		if err := s.TunnelIkeConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TunnelIpsecConfig != nil {
		if err := s.TunnelIpsecConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelBgpConfig struct {
	// The autonomous system number on the local end (Alibaba Cloud side) of the tunnel. Valid values: **1*	- to **4294967295**. Default value: **45104**.
	//
	// > We recommend that you use a private autonomous system number to establish a BGP connection with Alibaba Cloud. Refer to the relevant documentation for the range of private autonomous system numbers.
	//
	// example:
	//
	// 65530
	LocalAsn *int64 `json:"LocalAsn,omitempty" xml:"LocalAsn,omitempty"`
	// The BGP address on the local end (Alibaba Cloud side) of the tunnel. This address is an IP address within the BGP CIDR block.
	//
	// example:
	//
	// 169.254.10.1
	LocalBgpIp *string `json:"LocalBgpIp,omitempty" xml:"LocalBgpIp,omitempty"`
	// The BGP CIDR block of the tunnel. The CIDR block must fall within 169.254.0.0/16 and have a mask length of 30 bits. The CIDR block cannot be 169.254.0.0/30, 169.254.1.0/30, 169.254.2.0/30, 169.254.3.0/30, 169.254.4.0/30, 169.254.5.0/30, 169.254.6.0/30, or 169.254.169.252/30.
	//
	// > The tunnel CIDR blocks of the two tunnels under an IPsec-VPN connection must be different.
	//
	// example:
	//
	// 169.254.10.0/30
	TunnelCidr *string `json:"TunnelCidr,omitempty" xml:"TunnelCidr,omitempty"`
}

func (s CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelBgpConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelBgpConfig) GoString() string {
	return s.String()
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelBgpConfig) GetLocalAsn() *int64 {
	return s.LocalAsn
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelBgpConfig) GetLocalBgpIp() *string {
	return s.LocalBgpIp
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelBgpConfig) GetTunnelCidr() *string {
	return s.TunnelCidr
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelBgpConfig) SetLocalAsn(v int64) *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelBgpConfig {
	s.LocalAsn = &v
	return s
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelBgpConfig) SetLocalBgpIp(v string) *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelBgpConfig {
	s.LocalBgpIp = &v
	return s
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelBgpConfig) SetTunnelCidr(v string) *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelBgpConfig {
	s.TunnelCidr = &v
	return s
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelBgpConfig) Validate() error {
	return dara.Validate(s)
}

type CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIkeConfig struct {
	// The authentication algorithm used in Phase 1 negotiation. Valid values: **md5**, **sha1**, **sha256**, **sha384**, **sha512**. Default value: **sha1**.
	//
	// example:
	//
	// sha1
	IkeAuthAlg *string `json:"IkeAuthAlg,omitempty" xml:"IkeAuthAlg,omitempty"`
	// The encryption algorithm used in Phase 1 negotiation. Valid values: **aes**, **aes192**, **aes256**, **des**, or **3des**. Default value: **aes**.
	//
	// example:
	//
	// aes
	IkeEncAlg *string `json:"IkeEncAlg,omitempty" xml:"IkeEncAlg,omitempty"`
	// The lifetime of the SA generated in Phase 1 negotiation. Unit: seconds.
	//
	// Valid values: **0*	- to **86400**. Default value: **86400**.
	//
	// example:
	//
	// 86400
	IkeLifetime *int64 `json:"IkeLifetime,omitempty" xml:"IkeLifetime,omitempty"`
	// The negotiation mode of the IKE version. Valid values: **main*	- or **aggressive**. Default value: **main**.
	//
	// - **main**: main mode. The negotiation process is more secure.
	//
	// - **aggressive**: aggressive mode. The negotiation is faster and has a higher success rate.
	//
	// example:
	//
	// main
	IkeMode *string `json:"IkeMode,omitempty" xml:"IkeMode,omitempty"`
	// The Diffie-Hellman key exchange algorithm used in Phase 1 negotiation. Default value: **group2**.
	//
	// Valid values: **group1**, **group2**, **group5**, **group14**.
	//
	// example:
	//
	// group2
	IkePfs *string `json:"IkePfs,omitempty" xml:"IkePfs,omitempty"`
	// The version of the IKE protocol. Valid values: **ikev1*	- or **ikev2**. Default value: **ikev2**.
	//
	// Compared with IKEv1, IKEv2 simplifies the SA negotiation process and provides better support for multi-CIDR-block scenarios.
	//
	// example:
	//
	// ikev2
	IkeVersion *string `json:"IkeVersion,omitempty" xml:"IkeVersion,omitempty"`
	// The identifier on the local end (Alibaba Cloud side) of the tunnel, which is used in Phase 1 negotiation. The identifier is limited to 100 characters in length and cannot contain spaces. The default value is the IP address of the tunnel.
	//
	// **LocalId*	- supports the FQDN format. If you use the FQDN format, we recommend that you set the negotiation mode to **aggressive**.
	//
	// example:
	//
	// 47.XX.XX.1
	LocalId *string `json:"LocalId,omitempty" xml:"LocalId,omitempty"`
	// The pre-shared key used for identity authentication between the tunnel and the tunnel peer.
	//
	// - The key must be 1 to 100 characters in length and can contain digits, uppercase and lowercase letters, and the following characters. It cannot contain spaces. ```~!\\`@#$%^&*()_-+={}[]|;:\\",.<>/?```
	//
	// - If you do not specify a pre-shared key, the system randomly generates a 16-character string as the pre-shared key. You can call the [DescribeVpnAttachments](https://help.aliyun.com/document_detail/2526939.html) operation to query the pre-shared key that is automatically generated by the system.
	//
	// > The pre-shared keys of the tunnel and the tunnel peer must be the same. Otherwise, the system cannot establish the tunnel.
	//
	// example:
	//
	// 123456****
	Psk *string `json:"Psk,omitempty" xml:"Psk,omitempty"`
	// The identifier of the tunnel peer, which is used in Phase 1 negotiation. The identifier is limited to 100 characters in length and cannot contain spaces. The default value is the IP address of the customer gateway associated with the tunnel.
	//
	// **RemoteId*	- supports the FQDN format. If you use the FQDN format, we recommend that you set the negotiation mode to **aggressive**.
	//
	// example:
	//
	// 47.XX.XX.2
	RemoteId *string `json:"RemoteId,omitempty" xml:"RemoteId,omitempty"`
}

func (s CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIkeConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIkeConfig) GoString() string {
	return s.String()
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIkeConfig) GetIkeAuthAlg() *string {
	return s.IkeAuthAlg
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIkeConfig) GetIkeEncAlg() *string {
	return s.IkeEncAlg
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIkeConfig) GetIkeLifetime() *int64 {
	return s.IkeLifetime
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIkeConfig) GetIkeMode() *string {
	return s.IkeMode
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIkeConfig) GetIkePfs() *string {
	return s.IkePfs
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIkeConfig) GetIkeVersion() *string {
	return s.IkeVersion
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIkeConfig) GetLocalId() *string {
	return s.LocalId
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIkeConfig) GetPsk() *string {
	return s.Psk
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIkeConfig) GetRemoteId() *string {
	return s.RemoteId
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIkeConfig) SetIkeAuthAlg(v string) *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkeAuthAlg = &v
	return s
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIkeConfig) SetIkeEncAlg(v string) *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkeEncAlg = &v
	return s
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIkeConfig) SetIkeLifetime(v int64) *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkeLifetime = &v
	return s
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIkeConfig) SetIkeMode(v string) *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkeMode = &v
	return s
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIkeConfig) SetIkePfs(v string) *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkePfs = &v
	return s
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIkeConfig) SetIkeVersion(v string) *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkeVersion = &v
	return s
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIkeConfig) SetLocalId(v string) *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.LocalId = &v
	return s
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIkeConfig) SetPsk(v string) *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.Psk = &v
	return s
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIkeConfig) SetRemoteId(v string) *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.RemoteId = &v
	return s
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIkeConfig) Validate() error {
	return dara.Validate(s)
}

type CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIpsecConfig struct {
	// The authentication algorithm used in Phase 2 negotiation.
	//
	// Valid values: **md5**, **sha1**, **sha256**, **sha384**, **sha512**. Default value: **sha1**.
	//
	// example:
	//
	// sha1
	IpsecAuthAlg *string `json:"IpsecAuthAlg,omitempty" xml:"IpsecAuthAlg,omitempty"`
	// The encryption algorithm used in Phase 2 negotiation. Valid values: **aes**, **aes192**, **aes256**, **des**, or **3des**. Default value: **aes**.
	//
	// example:
	//
	// aes
	IpsecEncAlg *string `json:"IpsecEncAlg,omitempty" xml:"IpsecEncAlg,omitempty"`
	// The lifetime of the SA generated in Phase 2 negotiation. Unit: seconds.
	//
	// Valid values: **0*	- to **86400**. Default value: **86400**.
	//
	// example:
	//
	// 86400
	IpsecLifetime *int64 `json:"IpsecLifetime,omitempty" xml:"IpsecLifetime,omitempty"`
	// The Diffie-Hellman key exchange algorithm used in Phase 2 negotiation. Default value: **group2**.
	//
	// Valid values: **disabled**, **group1**, **group2**, **group5**, **group14**.
	//
	// example:
	//
	// group2
	IpsecPfs *string `json:"IpsecPfs,omitempty" xml:"IpsecPfs,omitempty"`
}

func (s CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIpsecConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIpsecConfig) GoString() string {
	return s.String()
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIpsecConfig) GetIpsecAuthAlg() *string {
	return s.IpsecAuthAlg
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIpsecConfig) GetIpsecEncAlg() *string {
	return s.IpsecEncAlg
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIpsecConfig) GetIpsecLifetime() *int64 {
	return s.IpsecLifetime
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIpsecConfig) GetIpsecPfs() *string {
	return s.IpsecPfs
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIpsecConfig) SetIpsecAuthAlg(v string) *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIpsecConfig {
	s.IpsecAuthAlg = &v
	return s
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIpsecConfig) SetIpsecEncAlg(v string) *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIpsecConfig {
	s.IpsecEncAlg = &v
	return s
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIpsecConfig) SetIpsecLifetime(v int64) *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIpsecConfig {
	s.IpsecLifetime = &v
	return s
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIpsecConfig) SetIpsecPfs(v string) *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIpsecConfig {
	s.IpsecPfs = &v
	return s
}

func (s *CreateVpnAttachmentRequestTunnelOptionsSpecificationTunnelIpsecConfig) Validate() error {
	return dara.Validate(s)
}
