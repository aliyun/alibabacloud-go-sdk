// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpnConnectionAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoConfigRoute(v bool) *ModifyVpnConnectionAttributeRequest
	GetAutoConfigRoute() *bool
	SetBgpConfig(v string) *ModifyVpnConnectionAttributeRequest
	GetBgpConfig() *string
	SetClientToken(v string) *ModifyVpnConnectionAttributeRequest
	GetClientToken() *string
	SetEffectImmediately(v bool) *ModifyVpnConnectionAttributeRequest
	GetEffectImmediately() *bool
	SetEnableDpd(v bool) *ModifyVpnConnectionAttributeRequest
	GetEnableDpd() *bool
	SetEnableNatTraversal(v bool) *ModifyVpnConnectionAttributeRequest
	GetEnableNatTraversal() *bool
	SetEnableTunnelsBgp(v bool) *ModifyVpnConnectionAttributeRequest
	GetEnableTunnelsBgp() *bool
	SetHealthCheckConfig(v string) *ModifyVpnConnectionAttributeRequest
	GetHealthCheckConfig() *string
	SetIkeConfig(v string) *ModifyVpnConnectionAttributeRequest
	GetIkeConfig() *string
	SetIpsecConfig(v string) *ModifyVpnConnectionAttributeRequest
	GetIpsecConfig() *string
	SetLocalSubnet(v string) *ModifyVpnConnectionAttributeRequest
	GetLocalSubnet() *string
	SetName(v string) *ModifyVpnConnectionAttributeRequest
	GetName() *string
	SetOwnerAccount(v string) *ModifyVpnConnectionAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyVpnConnectionAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyVpnConnectionAttributeRequest
	GetRegionId() *string
	SetRemoteCaCertificate(v string) *ModifyVpnConnectionAttributeRequest
	GetRemoteCaCertificate() *string
	SetRemoteSubnet(v string) *ModifyVpnConnectionAttributeRequest
	GetRemoteSubnet() *string
	SetResourceOwnerAccount(v string) *ModifyVpnConnectionAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyVpnConnectionAttributeRequest
	GetResourceOwnerId() *int64
	SetTunnelOptionsSpecification(v []*ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification) *ModifyVpnConnectionAttributeRequest
	GetTunnelOptionsSpecification() []*ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification
	SetVpnConnectionId(v string) *ModifyVpnConnectionAttributeRequest
	GetVpnConnectionId() *string
}

type ModifyVpnConnectionAttributeRequest struct {
	// Specifies whether to automatically advertise routes. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	AutoConfigRoute *bool `json:"AutoConfigRoute,omitempty" xml:"AutoConfigRoute,omitempty"`
	// This parameter is supported if you modify the configurations of an IPsec-VPN connection in single-tunnel mode.
	//
	// BGP configuration:
	//
	// 	- **BgpConfig.EnableBgp**: specifies whether to enable BGP. Valid values: **true*	- and **false**.
	//
	// 	- **BgpConfig.LocalAsn:*	- the autonomous system number (ASN) on the Alibaba Cloud side. Valid values: **1*	- to **4294967295**.
	//
	//     You can enter a value in two segments separated by a period (.). Each segment is 16 bits in length. Enter the number in each segment in decimal format.
	//
	//     For example, if you enter 123.456, the ASN is 8061384. The ASN is calculated by using the following formula: 123 × 65536 + 456 = 8061384.
	//
	// 	- **BgpConfig.TunnelCidr**: The CIDR block of the IPsec tunnel. The CIDR block must fall within 169.254.0.0/16 and the mask of the CIDR block must be 30 bits in length. The CIDR block cannot be 169.254.0.0/30, 169.254.1.0/30, 169.254.2.0/30, 169.254.3.0/30, 169.254.4.0/30, 169.254.5.0/30, 169.254.6.0/30, or 169.254.169.252/30.
	//
	//     > The CIDR block of the IPsec tunnel for each IPsec-VPN connection on a VPN gateway must be unique.
	//
	// 	- **LocalBgpIp**: the BGP address on the Alibaba Cloud side. It must be an IP address that falls within the CIDR block of the IPsec tunnel.
	//
	// > - This parameter is required when the VPN gateway has dynamic BGP enabled.
	//
	// > - Before you add BGP configurations, we recommend that you learn about how BGP dynamic routing works and the limits. For more information, see [Configure BGP dynamic routing](https://help.aliyun.com/document_detail/2638220.html).
	//
	// > - We recommend that you use a private ASN to establish BGP connections to Alibaba Cloud. For information about the range of private ASNs, see the relevant documentation.
	//
	// example:
	//
	// {"EnableBgp":"true","LocalAsn":"65530","TunnelCidr":"169.254.11.0/30","LocalBgpIp":"169.254.11.1"}
	BgpConfig *string `json:"BgpConfig,omitempty" xml:"BgpConfig,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the value of **RequestId*	- as the value of **ClientToken**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-0016e04115b
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to immediately start IPsec negotiations after the configuration takes effect. Valid values:
	//
	// 	- **true**: immediately starts IPsec negotiations after the configuration takes effect.
	//
	// 	- **false**: IPsec negotiations start when inbound traffic is detected.
	//
	// example:
	//
	// false
	EffectImmediately *bool `json:"EffectImmediately,omitempty" xml:"EffectImmediately,omitempty"`
	// You can specify this parameter if you modify the configuration of a single-tunnel IPsec-VPN connection.
	//
	// Specifies whether to enable the dead peer detection (DPD) feature. Valid values:
	//
	// 	- **true:**: enables the DPD feature. The initiator of the IPsec-VPN connection sends DPD packets to check the existence and availability of the peer. If no feedback is received from the peer within a specific period of time, the connection fails. Then, the ISAKMP SA, IPsec SA, and IPsec tunnel are deleted.
	//
	// 	- **false**: disables the DPD feature. The initiator of the IPsec-VPN connection does not send DPD packets.
	//
	// example:
	//
	// true
	EnableDpd *bool `json:"EnableDpd,omitempty" xml:"EnableDpd,omitempty"`
	// You can specify this parameter if you modify the configuration of a single-tunnel IPsec-VPN connection.
	//
	// Specifies whether to enable NAT traversal. Valid values:
	//
	// 	- **true*	- After NAT traversal is enabled, the initiator does not check the UDP ports during IKE negotiations and can automatically discover NAT gateway devices along the IPsec tunnel.
	//
	// 	- **false**
	//
	// example:
	//
	// true
	EnableNatTraversal *bool `json:"EnableNatTraversal,omitempty" xml:"EnableNatTraversal,omitempty"`
	// You can specify this parameter if you modify the configuration of a dual-tunnel IPsec-VPN connection.
	//
	// Specifies whether to enable BGP for the tunnel. Valid values: **true*	- and **false**.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// true
	EnableTunnelsBgp *bool `json:"EnableTunnelsBgp,omitempty" xml:"EnableTunnelsBgp,omitempty"`
	// You can specify this parameter if you modify the configuration of a single-tunnel IPsec-VPN connection.
	//
	// The health check configuration:
	//
	// 	- **HealthCheckConfig.enable**: specifies whether to enable health checks. Valid values: **true*	- and **false**.
	//
	// 	- **HealthCheckConfig.dip**: the destination IP address that is used for health checks.
	//
	// 	- **HealthCheckConfig.sip**: the source IP address that is used for health checks.
	//
	// 	- **HealthCheckConfig.interval**: the interval between two consecutive health checks. Unit: seconds.
	//
	// 	- **HealthCheckConfig.retry**: the maximum number of health check retries.
	//
	// example:
	//
	// {"enable":"true","dip":"192.168.1.1","sip":"10.1.1.1","interval":"3","retry":"3"}
	HealthCheckConfig *string `json:"HealthCheckConfig,omitempty" xml:"HealthCheckConfig,omitempty"`
	// This parameter is supported if you modify the configurations of an IPsec-VPN connection in single-tunnel mode.
	//
	// The configurations of Phase 1 negotiations:
	//
	// 	- **IkeConfig.Psk**: The pre-shared key that is used for identity authentication between the VPN gateway and the on-premises data center.
	//
	//     	- The key cannot contain space characters. The key must be 1 to 100 characters in length, and can contain digits, letters, and the following special characters: ``~!`@#$%^&*()_-+={}[]|;:\\",.<>/?``
	//
	//     	- If you do not specify a pre-shared key, the system randomly generates a 16-bit string as the pre-shared key. You can call the [DescribeVpnConnection](https://help.aliyun.com/document_detail/2526951.html) operation to query the pre-shared key that is automatically generated by the system.
	//
	//     > The pre-shared key of the IPsec-VPN connection must be the same as the authentication key of the on-premises data center. Otherwise, connections between the on-premises data center and the VPN gateway cannot be established.
	//
	// 	- **IkeConfig.IkeVersion**: the version of the Internet Key Exchange (IKE) protocol. Valid values: **ikev1*	- and **ikev2**.
	//
	//     Compared with IKEv1, IKEv2 simplifies the security association (SA) negotiation process and provides better support for scenarios with multiple CIDR blocks.
	//
	// 	- **IkeConfig.IkeMode**: the negotiation mode of IKE. Valid values: **main*	- and **aggressive**.
	//
	//     	- **main:*	- This mode offers higher security during negotiations.
	//
	//     	- **aggressive:*	- This mode supports faster negotiations and a higher success rate.
	//
	// 	- **IkeConfig.IkeEncAlg**: the encryption algorithm that is used in Phase 1 negotiations.
	//
	//     Valid values: **aes**, **aes192**, **aes256**, **des**, and **3des**.
	//
	// 	- **IkeConfig.IkeAuthAlg**: the authentication algorithm that is used in Phase 1 negotiations.
	//
	//     Valid values: **md5**, **sha1**, **sha256**, **sha384**, and **sha512**.
	//
	// 	- **IkeConfig.IkePfs**: the Diffie-Hellman key exchange algorithm that is used in Phase 1 negotiations. Valid values: **group1**, **group2**, **group5**, and **group14**.
	//
	// 	- **IkeConfig.IkeLifetime**: the SA lifetime as a result of Phase 1 negotiations. Unit: seconds Valid values: **0 to 86400**.
	//
	// 	- **IkeConfig.LocalId**: the identifier of the VPN gateway. The identifier cannot exceed 100 characters in length and cannot contain space characters. The default value is the IP address of the VPN gateway.
	//
	// 	- **IkeConfig.RemoteId**: the identifier of the customer gateway. The identifier cannot exceed 100 characters in length and cannot contain space characters. The default value is the IP address of the customer gateway.
	//
	// example:
	//
	// {"Psk":"pgw6dy7d1i8i****","IkeVersion":"ikev1","IkeMode":"main","IkeEncAlg":"aes","IkeAuthAlg":"sha1","IkePfs":"group2","IkeLifetime":86400,"LocalId":"116.64.XX.XX","RemoteId":"139.18.XX.XX"}
	IkeConfig *string `json:"IkeConfig,omitempty" xml:"IkeConfig,omitempty"`
	// You can specify this parameter if you modify the configuration of a single-tunnel IPsec-VPN connection.
	//
	// The configuration of Phase 2 negotiations:
	//
	// 	- **IpsecConfig.IpsecEncAlg**: the encryption algorithm that is used in Phase 2 negotiations.
	//
	//     Valid values: **aes**, **aes192**, **aes256**, **des**, and **3des**.
	//
	// 	- **IpsecConfig. IpsecAuthAlg**: the authentication algorithm that is used in Phase 2 negotiations.
	//
	//     Valid values: **md5**, **sha1**, **sha256**, **sha384**, and **sha512**.
	//
	// 	- **IpsecConfig. IpsecPfs**: the DH key exchange algorithm that is used in Phase 1 negotiations. If you specify this parameter, packets of all protocols are forwarded. Valid values: **disabled**, **group1**, **group2**, **group5**, and **group14**.
	//
	// 	- **IpsecConfig. IpsecLifetime:*	- the SA lifetime that is determined by Phase 2 negotiations. Unit: seconds. Valid values: **0 to 86400**.
	//
	// example:
	//
	// {"IpsecEncAlg":"aes","IpsecAuthAlg":"sha1","IpsecPfs":"group2","IpsecLifetime":86400}
	IpsecConfig *string `json:"IpsecConfig,omitempty" xml:"IpsecConfig,omitempty"`
	// The CIDR block used to connect the virtual private cloud (VPC) to the data center. The CIDR block is used in Phase 2 negotiations.
	//
	// Separate multiple CIDR blocks with commas (,). Example: 192.168.1.0/24,192.168.2.0/24.
	//
	// The following routing modes are supported:
	//
	// 	- If you set **LocalSubnet*	- and **RemoteSubnet*	- to 0.0.0.0/0, the routing mode of the IPsec-VPN connection is set to Destination Routing Mode.
	//
	// 	- If you set **LocalSubnet*	- and **RemoteSubnet*	- to specific CIDR blocks, the routing mode of the IPsec-VPN connection is set to Protected Data Flows.
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
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region in which the IPsec-VPN connection is created.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// You can specify this parameter if you modify the configuration of a single-tunnel IPsec-VPN connection.
	//
	// If the VPN gateway uses a ShangMi (SM) certificate, you can modify the CA certificate used by the IPsec peer.
	//
	// If the VPN gateway does not use an SM certificate, you cannot specify this parameter.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- MIIB7zCCAZW***	- -----END CERTIFICATE-----
	RemoteCaCertificate *string `json:"RemoteCaCertificate,omitempty" xml:"RemoteCaCertificate,omitempty"`
	// The CIDR block on the data center side. This CIDR block is used in Phase 2 negotiations.
	//
	// Separate multiple CIDR blocks with commas (,). Example: 192.168.3.0/24,192.168.4.0/24.
	//
	// The following routing modes are supported:
	//
	// 	- If you set **LocalSubnet*	- and **RemoteSubnet*	- to 0.0.0.0/0, the routing mode of the IPsec-VPN connection is set to Destination Routing Mode.
	//
	// 	- If you set **LocalSubnet*	- and **RemoteSubnet*	- to specific CIDR blocks, the routing mode of the IPsec-VPN connection is set to Protected Data Flows.
	//
	// example:
	//
	// 10.2.1.0/24,10.2.2.0/24
	RemoteSubnet         *string `json:"RemoteSubnet,omitempty" xml:"RemoteSubnet,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tunnel configurations.
	//
	// You can specify parameters in the **TunnelOptionsSpecification*	- array when you modify the configurations of an IPsec-VPN connection in dual-tunnel mode. You can modify the configurations of both the active and standby tunnels of the IPsec-VPN connection.
	//
	// if can be null:
	// true
	TunnelOptionsSpecification []*ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification `json:"TunnelOptionsSpecification,omitempty" xml:"TunnelOptionsSpecification,omitempty" type:"Repeated"`
	// The ID of the IPsec-VPN connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// vco-bp1bbi27hojx80nck****
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" xml:"VpnConnectionId,omitempty"`
}

func (s ModifyVpnConnectionAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnConnectionAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpnConnectionAttributeRequest) GetAutoConfigRoute() *bool {
	return s.AutoConfigRoute
}

func (s *ModifyVpnConnectionAttributeRequest) GetBgpConfig() *string {
	return s.BgpConfig
}

func (s *ModifyVpnConnectionAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyVpnConnectionAttributeRequest) GetEffectImmediately() *bool {
	return s.EffectImmediately
}

func (s *ModifyVpnConnectionAttributeRequest) GetEnableDpd() *bool {
	return s.EnableDpd
}

func (s *ModifyVpnConnectionAttributeRequest) GetEnableNatTraversal() *bool {
	return s.EnableNatTraversal
}

func (s *ModifyVpnConnectionAttributeRequest) GetEnableTunnelsBgp() *bool {
	return s.EnableTunnelsBgp
}

func (s *ModifyVpnConnectionAttributeRequest) GetHealthCheckConfig() *string {
	return s.HealthCheckConfig
}

func (s *ModifyVpnConnectionAttributeRequest) GetIkeConfig() *string {
	return s.IkeConfig
}

func (s *ModifyVpnConnectionAttributeRequest) GetIpsecConfig() *string {
	return s.IpsecConfig
}

func (s *ModifyVpnConnectionAttributeRequest) GetLocalSubnet() *string {
	return s.LocalSubnet
}

func (s *ModifyVpnConnectionAttributeRequest) GetName() *string {
	return s.Name
}

func (s *ModifyVpnConnectionAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyVpnConnectionAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyVpnConnectionAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyVpnConnectionAttributeRequest) GetRemoteCaCertificate() *string {
	return s.RemoteCaCertificate
}

func (s *ModifyVpnConnectionAttributeRequest) GetRemoteSubnet() *string {
	return s.RemoteSubnet
}

func (s *ModifyVpnConnectionAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyVpnConnectionAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyVpnConnectionAttributeRequest) GetTunnelOptionsSpecification() []*ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification {
	return s.TunnelOptionsSpecification
}

func (s *ModifyVpnConnectionAttributeRequest) GetVpnConnectionId() *string {
	return s.VpnConnectionId
}

func (s *ModifyVpnConnectionAttributeRequest) SetAutoConfigRoute(v bool) *ModifyVpnConnectionAttributeRequest {
	s.AutoConfigRoute = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequest) SetBgpConfig(v string) *ModifyVpnConnectionAttributeRequest {
	s.BgpConfig = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequest) SetClientToken(v string) *ModifyVpnConnectionAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequest) SetEffectImmediately(v bool) *ModifyVpnConnectionAttributeRequest {
	s.EffectImmediately = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequest) SetEnableDpd(v bool) *ModifyVpnConnectionAttributeRequest {
	s.EnableDpd = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequest) SetEnableNatTraversal(v bool) *ModifyVpnConnectionAttributeRequest {
	s.EnableNatTraversal = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequest) SetEnableTunnelsBgp(v bool) *ModifyVpnConnectionAttributeRequest {
	s.EnableTunnelsBgp = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequest) SetHealthCheckConfig(v string) *ModifyVpnConnectionAttributeRequest {
	s.HealthCheckConfig = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequest) SetIkeConfig(v string) *ModifyVpnConnectionAttributeRequest {
	s.IkeConfig = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequest) SetIpsecConfig(v string) *ModifyVpnConnectionAttributeRequest {
	s.IpsecConfig = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequest) SetLocalSubnet(v string) *ModifyVpnConnectionAttributeRequest {
	s.LocalSubnet = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequest) SetName(v string) *ModifyVpnConnectionAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequest) SetOwnerAccount(v string) *ModifyVpnConnectionAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequest) SetOwnerId(v int64) *ModifyVpnConnectionAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequest) SetRegionId(v string) *ModifyVpnConnectionAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequest) SetRemoteCaCertificate(v string) *ModifyVpnConnectionAttributeRequest {
	s.RemoteCaCertificate = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequest) SetRemoteSubnet(v string) *ModifyVpnConnectionAttributeRequest {
	s.RemoteSubnet = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequest) SetResourceOwnerAccount(v string) *ModifyVpnConnectionAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequest) SetResourceOwnerId(v int64) *ModifyVpnConnectionAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequest) SetTunnelOptionsSpecification(v []*ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification) *ModifyVpnConnectionAttributeRequest {
	s.TunnelOptionsSpecification = v
	return s
}

func (s *ModifyVpnConnectionAttributeRequest) SetVpnConnectionId(v string) *ModifyVpnConnectionAttributeRequest {
	s.VpnConnectionId = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification struct {
	// The ID of the customer gateway associated with the tunnel.
	//
	// example:
	//
	// cgw-1nmwbpgrp7ssqm1yn****
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" xml:"CustomerGatewayId,omitempty"`
	// Specifies whether to enable the Dead Peer Detection (DPD) feature for the tunnel. Valid values:
	//
	// 	- **true**: enables DPD. The initiator of the IPsec-VPN connection sends DPD packets to check the existence and availability of the peer. If no feedback is received from the peer within the specified period of time, the connection fails. In this case, ISAKMP SA and IPsec SA are deleted. The security tunnel is also deleted.
	//
	// 	- **false**: disables DPD. The initiator of the IPsec-VPN connection does not send DPD packets.
	//
	// example:
	//
	// true
	EnableDpd *bool `json:"EnableDpd,omitempty" xml:"EnableDpd,omitempty"`
	// Specifies whether to enable NAT traversal for the tunnel. Valid values:
	//
	// 	- **true**: enables NAT traversal. After NAT traversal is enabled, the initiator does not check the UDP ports during IKE negotiations and can automatically discover NAT gateway devices along the IPsec-VPN tunnel.
	//
	// 	- **false**: disables NAT traversal.
	//
	// example:
	//
	// true
	EnableNatTraversal *bool `json:"EnableNatTraversal,omitempty" xml:"EnableNatTraversal,omitempty"`
	// If the VPN gateway uses an SM certificate, you can modify the CA certificate used by the IPsec peer.
	//
	// If the VPN gateway does not use an SM certificate, this parameter is not supported.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- MIIB7zCCAZW***	- -----END CERTIFICATE-----
	RemoteCaCertificate *string `json:"RemoteCaCertificate,omitempty" xml:"RemoteCaCertificate,omitempty"`
	// The tunnel role. Valid values:
	//
	// 	- **master**: The tunnel is an active tunnel.
	//
	// 	- **slave**: The tunnel is a standby tunnel.
	//
	// example:
	//
	// master
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The Border Gateway Protocol (BGP) configurations of the tunnel.
	TunnelBgpConfig *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig `json:"TunnelBgpConfig,omitempty" xml:"TunnelBgpConfig,omitempty" type:"Struct"`
	// The tunnel ID.
	//
	// example:
	//
	// tun-opsqc4d97wni27****
	TunnelId *string `json:"TunnelId,omitempty" xml:"TunnelId,omitempty"`
	// The configurations of Phase 1 negotiations.
	TunnelIkeConfig *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig `json:"TunnelIkeConfig,omitempty" xml:"TunnelIkeConfig,omitempty" type:"Struct"`
	// The configurations of Phase 2 negotiations.
	TunnelIpsecConfig *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig `json:"TunnelIpsecConfig,omitempty" xml:"TunnelIpsecConfig,omitempty" type:"Struct"`
}

func (s ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification) GoString() string {
	return s.String()
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification) GetCustomerGatewayId() *string {
	return s.CustomerGatewayId
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification) GetEnableDpd() *bool {
	return s.EnableDpd
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification) GetEnableNatTraversal() *bool {
	return s.EnableNatTraversal
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification) GetRemoteCaCertificate() *string {
	return s.RemoteCaCertificate
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification) GetRole() *string {
	return s.Role
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification) GetTunnelBgpConfig() *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig {
	return s.TunnelBgpConfig
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification) GetTunnelId() *string {
	return s.TunnelId
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification) GetTunnelIkeConfig() *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig {
	return s.TunnelIkeConfig
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification) GetTunnelIpsecConfig() *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig {
	return s.TunnelIpsecConfig
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification) SetCustomerGatewayId(v string) *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification {
	s.CustomerGatewayId = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification) SetEnableDpd(v bool) *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification {
	s.EnableDpd = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification) SetEnableNatTraversal(v bool) *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification {
	s.EnableNatTraversal = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification) SetRemoteCaCertificate(v string) *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification {
	s.RemoteCaCertificate = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification) SetRole(v string) *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification {
	s.Role = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification) SetTunnelBgpConfig(v *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig) *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification {
	s.TunnelBgpConfig = v
	return s
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification) SetTunnelId(v string) *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification {
	s.TunnelId = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification) SetTunnelIkeConfig(v *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification {
	s.TunnelIkeConfig = v
	return s
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification) SetTunnelIpsecConfig(v *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification {
	s.TunnelIpsecConfig = v
	return s
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification) Validate() error {
	return dara.Validate(s)
}

type ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig struct {
	// The ASN of the tunnel on the Alibaba Cloud side. Valid values: **1*	- to **4294967295**. Default value: **45104**.
	//
	// >  You can specify this parameter only if **EnableTunnelsBgp*	- is set to **true**.
	//
	// 	- Before you add BGP configurations, we recommend that you learn about how BGP dynamic routing works and the limits. For more information, see [Configure BGP dynamic routing](https://help.aliyun.com/document_detail/2638220.html).
	//
	// 	- We recommend that you use a private ASN to establish BGP connections to Alibaba Cloud. For information about the range of private ASNs, see the relevant documentation.
	//
	// example:
	//
	// 65530
	LocalAsn *int64 `json:"LocalAsn,omitempty" xml:"LocalAsn,omitempty"`
	// The BGP IP address of the tunnel on the Alibaba Cloud side. The address is an IP address that falls within the BGP CIDR block.
	//
	// example:
	//
	// 169.254.10.1
	LocalBgpIp *string `json:"LocalBgpIp,omitempty" xml:"LocalBgpIp,omitempty"`
	// The BGP CIDR block of the tunnel.
	//
	// The CIDR block must fall within 169.254.0.0/16 and the mask of the CIDR block must be 30 bits in length. The CIDR block cannot be 169.254.0.0/30, 169.254.1.0/30, 169.254.2.0/30, 169.254.3.0/30, 169.254.4.0/30, 169.254.5.0/30, 169.254.6.0/30, or 169.254.169.252/30.
	//
	// >  The BGP CIDR block of each tunnel must be unique on a VPN gateway.
	//
	// example:
	//
	// 169.254.10.0/30
	TunnelCidr *string `json:"TunnelCidr,omitempty" xml:"TunnelCidr,omitempty"`
}

func (s ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig) GoString() string {
	return s.String()
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig) GetLocalAsn() *int64 {
	return s.LocalAsn
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig) GetLocalBgpIp() *string {
	return s.LocalBgpIp
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig) GetTunnelCidr() *string {
	return s.TunnelCidr
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig) SetLocalAsn(v int64) *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig {
	s.LocalAsn = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig) SetLocalBgpIp(v string) *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig {
	s.LocalBgpIp = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig) SetTunnelCidr(v string) *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig {
	s.TunnelCidr = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig struct {
	// The authentication algorithm that is used in Phase 1 negotiations.
	//
	// Valid values: **md5**, **sha1**, **sha256**, **sha384**, and **sha512**.
	//
	// example:
	//
	// md5
	IkeAuthAlg *string `json:"IkeAuthAlg,omitempty" xml:"IkeAuthAlg,omitempty"`
	// The encryption algorithm that is used in Phase 1 negotiations.
	//
	// Valid values: **aes**, **aes192**, **aes256**, **des**, and **3des**.
	//
	// example:
	//
	// aes
	IkeEncAlg *string `json:"IkeEncAlg,omitempty" xml:"IkeEncAlg,omitempty"`
	// The SA lifetime as a result of Phase 1 negotiations. Unit: seconds Valid values: **0*	- to **86400**.
	//
	// example:
	//
	// 86400
	IkeLifetime *int64 `json:"IkeLifetime,omitempty" xml:"IkeLifetime,omitempty"`
	// The negotiation mode of IKE. Valid values:
	//
	// 	- **main:*	- This mode offers higher security during negotiations.
	//
	// 	- **aggressive:*	- This mode supports faster negotiations and a higher success rate.
	//
	// example:
	//
	// main
	IkeMode *string `json:"IkeMode,omitempty" xml:"IkeMode,omitempty"`
	// The Diffie-Hellman key exchange algorithm that is used in Phase 1 negotiations. Valid values: **group1**, **group2**, **group5**, and **group14**.
	//
	// example:
	//
	// group2
	IkePfs *string `json:"IkePfs,omitempty" xml:"IkePfs,omitempty"`
	// The version of the IKE protocol. Valid values: **ikev1*	- and **ikev2**.
	//
	// Compared with IKEv1, IKEv2 simplifies the SA negotiation process and provides better support for scenarios with multiple CIDR blocks.
	//
	// example:
	//
	// ikev1
	IkeVersion *string `json:"IkeVersion,omitempty" xml:"IkeVersion,omitempty"`
	// The identifier on the Alibaba Cloud side, which is used in Phase 1 negotiations. The identifier cannot exceed 100 characters in length and cannot contain space characters. The default value is the IP address of the tunnel.
	//
	// **LocalId*	- supports fully qualified domain names (FQDNs). If you use an FQDN, we recommend that you set the negotiation mode to **aggressive**.
	//
	// example:
	//
	// 47.21.XX.XX
	LocalId *string `json:"LocalId,omitempty" xml:"LocalId,omitempty"`
	// The pre-shared key, which is used for identity authentication between the tunnel and the tunnel peer.
	//
	// 	- The key cannot contain space characters. The key must be 1 to 100 characters in length, and can contain digits, letters, and the following special characters: ``~!\\`@#$%^&*()_-+={}[]|;:\\",.<>/?``
	//
	// 	- If you do not specify a pre-shared key, the system randomly generates a 16-bit string as the pre-shared key. You can call the [DescribeVpnConnection](https://help.aliyun.com/document_detail/2526951.html) operation to query the pre-shared key that is automatically generated by the system.
	//
	// >  The tunnel and the tunnel peer must use the same pre-shared key. Otherwise, the tunnel cannot be built.
	//
	// example:
	//
	// 123456****
	Psk *string `json:"Psk,omitempty" xml:"Psk,omitempty"`
	// The identifier of the tunnel peer, which is used in Phase 1 negotiations. The identifier cannot exceed 100 characters in length and cannot contain space characters. The default value is the IP address of the customer gateway that is associated with the tunnel.
	//
	// **RemoteId*	- supports FQDNs. If you use an FQDN, we recommend that you set the negotiation mode to **aggressive**.
	//
	// example:
	//
	// 47.42.XX.XX
	RemoteId *string `json:"RemoteId,omitempty" xml:"RemoteId,omitempty"`
}

func (s ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) GoString() string {
	return s.String()
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) GetIkeAuthAlg() *string {
	return s.IkeAuthAlg
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) GetIkeEncAlg() *string {
	return s.IkeEncAlg
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) GetIkeLifetime() *int64 {
	return s.IkeLifetime
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) GetIkeMode() *string {
	return s.IkeMode
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) GetIkePfs() *string {
	return s.IkePfs
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) GetIkeVersion() *string {
	return s.IkeVersion
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) GetLocalId() *string {
	return s.LocalId
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) GetPsk() *string {
	return s.Psk
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) GetRemoteId() *string {
	return s.RemoteId
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) SetIkeAuthAlg(v string) *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkeAuthAlg = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) SetIkeEncAlg(v string) *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkeEncAlg = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) SetIkeLifetime(v int64) *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkeLifetime = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) SetIkeMode(v string) *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkeMode = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) SetIkePfs(v string) *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkePfs = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) SetIkeVersion(v string) *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkeVersion = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) SetLocalId(v string) *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.LocalId = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) SetPsk(v string) *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.Psk = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) SetRemoteId(v string) *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.RemoteId = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig struct {
	// The authentication algorithm that is used in Phase 2 negotiations.
	//
	// Valid values: **md5**, **sha1**, **sha256**, **sha384**, and **sha512**.
	//
	// example:
	//
	// md5
	IpsecAuthAlg *string `json:"IpsecAuthAlg,omitempty" xml:"IpsecAuthAlg,omitempty"`
	// The encryption algorithm that is used in Phase 2 negotiations.
	//
	// Valid values: **aes**, **aes192**, **aes256**, **des**, and **3des**.
	//
	// example:
	//
	// aes
	IpsecEncAlg *string `json:"IpsecEncAlg,omitempty" xml:"IpsecEncAlg,omitempty"`
	// The SA lifetime as a result of Phase 2 negotiations. Unit: seconds Valid values: **0*	- to **86400**.
	//
	// example:
	//
	// 86400
	IpsecLifetime *int32 `json:"IpsecLifetime,omitempty" xml:"IpsecLifetime,omitempty"`
	// The Diffie-Hellman key exchange algorithm that is used in Phase 2 negotiations.
	//
	// Valid values: **disabled**, **group1**, **group2**, **group5**, and **group14**.
	//
	// example:
	//
	// group2
	IpsecPfs *string `json:"IpsecPfs,omitempty" xml:"IpsecPfs,omitempty"`
}

func (s ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) GoString() string {
	return s.String()
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) GetIpsecAuthAlg() *string {
	return s.IpsecAuthAlg
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) GetIpsecEncAlg() *string {
	return s.IpsecEncAlg
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) GetIpsecLifetime() *int32 {
	return s.IpsecLifetime
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) GetIpsecPfs() *string {
	return s.IpsecPfs
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) SetIpsecAuthAlg(v string) *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig {
	s.IpsecAuthAlg = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) SetIpsecEncAlg(v string) *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig {
	s.IpsecEncAlg = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) SetIpsecLifetime(v int32) *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig {
	s.IpsecLifetime = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) SetIpsecPfs(v string) *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig {
	s.IpsecPfs = &v
	return s
}

func (s *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) Validate() error {
	return dara.Validate(s)
}
