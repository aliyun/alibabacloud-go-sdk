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
	// Specifies whether to automatically publish route entry. Valid values:
	//
	// - **true**: Automatically publishes route entries.
	//
	//
	//
	// - **false**: Does not automatically publish route entries.
	//
	// example:
	//
	// true
	AutoConfigRoute *bool `json:"AutoConfigRoute,omitempty" xml:"AutoConfigRoute,omitempty"`
	// This parameter is supported when you modify a single-tunnel IPsec-VPN connection.
	//
	// The BGP configuration:
	//
	// - **BgpConfig.EnableBgp**: Specifies whether to enable the BGP feature. Valid values: **true*	- and **false**.
	//
	// - **BgpConfig.LocalAsn**: The autonomous system number (ASN) on the Alibaba Cloud side. Valid values: **1*	- to **4294967295**.
	//
	//    The ASN can be entered in the two-segment format: the first 16 bits.the last 16 bits. Each segment is entered in decimal format.
	//
	//     For example, if you enter 123.456, the ASN is 123 × 65536 + 456 = 8061384.
	//
	// - **BgpConfig.TunnelCidr**: The CIDR block of the IPsec tunnel. The CIDR block must be a mask length of 30 within 169.254.0.0/16 and cannot be 169.254.0.0/30, 169.254.1.0/30, 169.254.2.0/30, 169.254.3.0/30, 169.254.4.0/30, 169.254.5.0/30, 169.254.6.0/30, or 169.254.169.252/30.
	//
	//     > The IPsec tunnel CIDR block of each IPsec-VPN connection under a VPN gateway instance must be unique.
	//
	// - **LocalBgpIp**: The BGP address on the Alibaba Cloud side. This address is an IP address within the IPsec tunnel CIDR block.
	//
	// > - Configure this parameter when BGP dynamic routing is enabled on your VPN gateway.
	//
	// >- Before you add BGP configurations, understand the working mechanism and limits of BGP dynamic routing. For more information, see [Configure BGP dynamic routing](https://help.aliyun.com/document_detail/2638220.html).
	//
	// >- Use a private ASN to establish a BGP connection with Alibaba Cloud. Refer to the relevant documentation for the range of private ASNs.
	//
	// example:
	//
	// {"EnableBgp":"true","LocalAsn":"65530","TunnelCidr":"169.254.11.0/30","LocalBgpIp":"169.254.11.1"}
	BgpConfig *string `json:"BgpConfig,omitempty" xml:"BgpConfig,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-0016e04115b
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether the IPsec-VPN connection configuration takes effect immediately. Valid values:
	//
	//
	//
	// - **true**: The system immediately initiates IPsec protocol negotiation after the configuration is complete.
	//
	//
	//
	// - **false**: The system initiates IPsec protocol negotiation only when inbound traffic is detected.
	//
	// example:
	//
	// false
	EffectImmediately *bool `json:"EffectImmediately,omitempty" xml:"EffectImmediately,omitempty"`
	// This parameter is supported when you modify a single-tunnel IPsec-VPN connection.
	//
	// Specifies whether to enable the DPD (Dead Peer Detection) feature. Valid values:
	//
	// - **true**: Enables the DPD feature. The IPsec initiator sends DPD packets to check whether the peer device is alive. If no correct response is received within the specified period of time, the peer is considered disconnected. The ISAKMP SA and the corresponding IPsec SA are deleted, and the security tunnel is also deleted.
	//
	// - **false**: Disables the DPD feature. The IPsec initiator does not send DPD packets.
	//
	// example:
	//
	// true
	EnableDpd *bool `json:"EnableDpd,omitempty" xml:"EnableDpd,omitempty"`
	// This parameter is supported when you modify a single-tunnel IPsec-VPN connection.
	//
	// Specifies whether to enable the NAT traversal feature. Valid values:
	//
	// - **true**: Enables NAT traversal. After NAT traversal is enabled, the IKE negotiation process skips UDP port number verification and can discover NAT gateway devices in the VPN tunnel.
	//
	// - **false**: Disables NAT traversal.
	//
	// example:
	//
	// true
	EnableNatTraversal *bool `json:"EnableNatTraversal,omitempty" xml:"EnableNatTraversal,omitempty"`
	// This parameter is supported when you modify a dual-tunnel IPsec-VPN connection.
	//
	// Specifies whether to enable BGP for the tunnels. Valid values: **true*	- and **false**.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// true
	EnableTunnelsBgp *bool `json:"EnableTunnelsBgp,omitempty" xml:"EnableTunnelsBgp,omitempty"`
	// This parameter is supported when you modify a single-tunnel IPsec-VPN connection.
	//
	// The health check configuration:
	//
	// - **HealthCheckConfig.enable**: Specifies whether to enable health checks. Valid values: **true*	- and **false**.
	//
	// - **HealthCheckConfig.dip**: The destination IP address of the health check.
	//
	// - **HealthCheckConfig.sip**: The source IP address of the health check.
	//
	// - **HealthCheckConfig.interval**: The retry interval of the health check. Unit: seconds.
	//
	// - **HealthCheckConfig.retry**: The number of retries for the health check.
	//
	// example:
	//
	// {"enable":"true","dip":"192.168.1.1","sip":"10.1.1.1","interval":"3","retry":"3"}
	HealthCheckConfig *string `json:"HealthCheckConfig,omitempty" xml:"HealthCheckConfig,omitempty"`
	// This parameter is supported when you modify a single-tunnel IPsec-VPN connection.
	//
	// The Phase 1 negotiation configuration:
	//
	//
	//
	// - **IkeConfig.Psk**: The pre-shared key used for identity authentication between the VPN gateway and the on-premises data center.
	//
	//     - The key must be 1 to 100 characters in length and can contain digits, letters, and the following characters. It cannot contain spaces. ```~!`@#$%^&*()_-+={}[]|;:\\",.<>/?```
	//
	//     - If you do not specify a pre-shared key, the system generates a random 16-character string as the pre-shared key. You can call the [DescribeVpnConnection](https://help.aliyun.com/document_detail/2526951.html) operation to query the pre-shared key that is automatically generated by the system.
	//
	//     > The pre-shared key on the IPsec-VPN connection side must be the same as the authentication key on the on-premises data center side. Otherwise, the connection between the on-premises data center and the VPN gateway cannot be established.
	//
	// - **IkeConfig.IkeVersion**: The version of the IKE protocol. Valid values: **ikev1*	- and **ikev2**.
	//
	//     Compared with IKEv1, IKEv2 simplifies the SA negotiation process and provides better support for multi-CIDR-block scenarios.
	//
	//     <props="china"><ph>If the VPN gateway instance type is Chinese SM-based, only **ikev1*	- is supported for the IKE version.</ph>
	//
	// - **IkeConfig.IkeMode**: The negotiation mode of the IKE version. Valid values: **main*	- and **aggressive**.
	//
	//     - **main**: Main mode. This mode offers high security during negotiation.
	//
	//     - **aggressive**: Aggressive mode. This mode supports fast negotiation and a higher success rate.
	//
	//     <props="china"><ph>If the VPN gateway instance type is Chinese SM-based, only **main*	- is supported for the negotiation mode.</ph>
	//
	// - **IkeConfig.IkeEncAlg**: The encryption algorithm used in Phase 1 negotiation.
	//
	//     <props="intl"><ph>Valid values: **aes**, **aes192**, **aes256**, **des**, and **3des**.</ph>
	//
	//     <props="china"><ph>If the VPN gateway instance type is standard, valid values are **aes**, **aes192**, **aes256**, **des**, and **3des**.</ph>
	//
	//     <props="china"><ph>If the VPN gateway instance type is Chinese SM-based, the only valid value is **sm4**.</ph>
	//
	// - **IkeConfig.IkeAuthAlg**: The authentication algorithm used in Phase 1 negotiation.
	//
	//     <props="intl"><ph>Valid values: **md5**, **sha1**, **sha256**, **sha384**, and **sha512**.</ph>
	//
	//     <props="china"><ph>If the VPN gateway instance type is standard, valid values are **md5**, **sha1**, **sha256**, **sha384**, and **sha512**.</ph>
	//
	//     <props="china"><ph>If the VPN gateway instance type is Chinese SM-based, the only valid value is **sm3**.</ph>
	//
	// - **IkeConfig.IkePfs**: The Diffie-Hellman key exchange algorithm used in Phase 1 negotiation. Valid values: **group1**, **group2**, **group5**, and **group14**.
	//
	// - **IkeConfig.IkeLifetime**: The SA lifetime determined by Phase 1 negotiation. Unit: seconds. Valid values: **0*	- to **86400**.
	//
	// - **IkeConfig.LocalId**: The identifier of the VPN gateway. The identifier can be up to 100 characters in length and cannot contain spaces. The default value is the IP address of the VPN gateway.
	//
	// - **IkeConfig.RemoteId**: The identifier of the customer gateway. The identifier can be up to 100 characters in length and cannot contain spaces. The default value is the IP address of the customer gateway.
	//
	// example:
	//
	// {"Psk":"pgw6dy7d1i8i****","IkeVersion":"ikev1","IkeMode":"main","IkeEncAlg":"aes","IkeAuthAlg":"sha1","IkePfs":"group2","IkeLifetime":86400,"LocalId":"116.64.XX.XX","RemoteId":"139.18.XX.XX"}
	IkeConfig *string `json:"IkeConfig,omitempty" xml:"IkeConfig,omitempty"`
	// This parameter is supported when you modify a single-tunnel IPsec-VPN connection.
	//
	// The Phase 2 negotiation configuration:
	//
	// - **IpsecConfig.IpsecEncAlg**: The encryption algorithm used in Phase 2 negotiation.
	//
	//     <props="intl"><ph>Valid values: **aes**, **aes192**, **aes256**, **des**, and **3des**.</ph>
	//
	//     <props="china"><ph>If the VPN gateway instance type is standard, valid values are **aes**, **aes192**, **aes256**, **des**, and **3des**.</ph>
	//
	//     <props="china"><ph>If the VPN gateway instance type is Chinese SM-based, the only valid value is **sm4**.</ph>
	//
	// - **IpsecConfig. IpsecAuthAlg**: The authentication algorithm used in Phase 2 negotiation.
	//
	//     <props="intl"><ph>Valid values: **md5**, **sha1**, **sha256**, **sha384**, and **sha512**.</ph>
	//
	//     <props="china"><ph>If the VPN gateway instance type is standard, valid values are **md5**, **sha1**, **sha256**, **sha384**, and **sha512**.</ph>
	//
	//     <props="china"><ph>If the VPN gateway instance type is Chinese SM-based, the only valid value is **sm3**.</ph>
	//
	// - **IpsecConfig. IpsecPfs**: The Diffie-Hellman key exchange algorithm used in Phase 2 negotiation for forwarding packets of all protocols. Valid values: **disabled**, **group1**, **group2**, **group5**, and **group14**.
	//
	// - **IpsecConfig. IpsecLifetime**: The SA lifetime determined by Phase 2 negotiation. Unit: seconds. Valid values: **0*	- to **86400**.
	//
	// example:
	//
	// {"IpsecEncAlg":"aes","IpsecAuthAlg":"sha1","IpsecPfs":"group2","IpsecLifetime":86400}
	IpsecConfig *string `json:"IpsecConfig,omitempty" xml:"IpsecConfig,omitempty"`
	// The CIDR block on the VPC side that needs to communicate with the on-premises data center. This CIDR block is used in Phase 2 negotiation.
	//
	// Separate multiple CIDR blocks with commas (,). Example: 192.168.1.0/24,192.168.2.0/24.
	//
	// The following routing modes are supported for IPsec-VPN connections:
	//
	// - If you set both **LocalSubnet*	- and **RemoteSubnet*	- to 0.0.0.0/0, the destination routing mode is used.
	//
	// - If you set both **LocalSubnet*	- and **RemoteSubnet*	- to specific CIDR blocks, the protected data flow mode is used.
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
	// The region ID of the IPsec-VPN connection.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is supported when you modify a single-tunnel IPsec-VPN connection.
	//
	// If the current VPN gateway instance is a Chinese SM-based VPN gateway, you can modify the CA certificate of the peer.
	//
	// If the current VPN gateway instance is a standard VPN gateway, this parameter is not supported.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- MIIB7zCCAZW***	- -----END CERTIFICATE-----
	RemoteCaCertificate *string `json:"RemoteCaCertificate,omitempty" xml:"RemoteCaCertificate,omitempty"`
	// The CIDR block on the on-premises data center side. This CIDR block is used in Phase 2 negotiation.
	//
	// Separate multiple CIDR blocks with commas (,). Example: 192.168.3.0/24,192.168.4.0/24.
	//
	// The following routing modes are supported for IPsec-VPN connections:
	//
	// - If you set both **LocalSubnet*	- and **RemoteSubnet*	- to 0.0.0.0/0, the destination routing mode is used.
	//
	// - If you set both **LocalSubnet*	- and **RemoteSubnet*	- to specific CIDR blocks, the protected data flow mode is used.
	//
	// example:
	//
	// 10.2.1.0/24,10.2.2.0/24
	RemoteSubnet         *string `json:"RemoteSubnet,omitempty" xml:"RemoteSubnet,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tunnel configurations to modify.
	//
	// The parameters under the **TunnelOptionsSpecification*	- array are supported only when you modify a dual-tunnel IPsec-VPN connection. You can modify the configurations of both the active and standby tunnels of the IPsec-VPN connection at the same time.
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

type ModifyVpnConnectionAttributeRequestTunnelOptionsSpecification struct {
	// The instance ID of the customer gateway associated with the tunnel.
	//
	// example:
	//
	// cgw-1nmwbpgrp7ssqm1yn****
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" xml:"CustomerGatewayId,omitempty"`
	// Specifies whether to enable the DPD (Dead Peer Detection) feature for the tunnel. Valid values:
	//
	// - **true**: Enables the DPD feature. The IPsec initiator sends DPD packets to check whether the peer device is alive. If no correct response is received within the specified period of time, the peer is considered disconnected. The ISAKMP SA and the corresponding IPsec SA are deleted, and the security tunnel is also deleted.
	//
	// - **false**: Disables the DPD feature. The IPsec initiator does not send DPD packets.
	//
	// example:
	//
	// true
	EnableDpd *bool `json:"EnableDpd,omitempty" xml:"EnableDpd,omitempty"`
	// Specifies whether to enable the NAT traversal feature for the tunnel. Valid values:
	//
	// - **true**: Enables NAT traversal. After NAT traversal is enabled, the IKE negotiation process skips UDP port number verification and can discover NAT gateway devices in the VPN tunnel.
	//
	// - **false**: Disables NAT traversal.
	//
	// example:
	//
	// true
	EnableNatTraversal *bool `json:"EnableNatTraversal,omitempty" xml:"EnableNatTraversal,omitempty"`
	// If the current VPN gateway instance is a Chinese SM-based VPN gateway, you can modify the CA certificate of the tunnel peer.
	//
	// If the current VPN gateway instance is a standard VPN gateway, this parameter is not supported.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- MIIB7zCCAZW***	- -----END CERTIFICATE-----
	RemoteCaCertificate *string `json:"RemoteCaCertificate,omitempty" xml:"RemoteCaCertificate,omitempty"`
	// The role of the tunnel. Valid values:
	//
	// - **master**: The tunnel is the active tunnel.
	//
	// - **slave**: The tunnel is the standby tunnel.
	//
	// example:
	//
	// master
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The BGP configuration of the tunnel to modify.
	TunnelBgpConfig *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig `json:"TunnelBgpConfig,omitempty" xml:"TunnelBgpConfig,omitempty" type:"Struct"`
	// The tunnel ID.
	//
	// example:
	//
	// tun-opsqc4d97wni27****
	TunnelId *string `json:"TunnelId,omitempty" xml:"TunnelId,omitempty"`
	// The Phase 1 negotiation configuration.
	TunnelIkeConfig *ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig `json:"TunnelIkeConfig,omitempty" xml:"TunnelIkeConfig,omitempty" type:"Struct"`
	// The Phase 2 negotiation configuration.
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

type ModifyVpnConnectionAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig struct {
	// The autonomous system number (ASN) on the tunnel local side (Alibaba Cloud side). Valid values: **1*	- to **4294967295**. Default value: **45104**.
	//
	// > - This parameter can be configured or modified after you enable the BGP dynamic route feature for the IPsec-VPN connection (by setting **EnableTunnelsBgp*	- to **true**).
	//
	// >- Before you add BGP configurations, understand the working mechanism and limits of BGP dynamic route feature. For more information, see [Configure BGP dynamic routing](https://help.aliyun.com/document_detail/2638220.html).
	//
	// >- Use a private ASN to establish a BGP connection with Alibaba Cloud. Refer to the relevant documentation for the range of private ASNs.
	//
	// example:
	//
	// 65530
	LocalAsn *int64 `json:"LocalAsn,omitempty" xml:"LocalAsn,omitempty"`
	// The BGP address on the tunnel local side (Alibaba Cloud side). This address is an IP address within the BGP CIDR block.
	//
	// example:
	//
	// 169.254.10.1
	LocalBgpIp *string `json:"LocalBgpIp,omitempty" xml:"LocalBgpIp,omitempty"`
	// The BGP CIDR block of the tunnel.
	//
	// The CIDR block must be a mask length of 30 within 169.254.0.0/16 and cannot be 169.254.0.0/30, 169.254.1.0/30, 169.254.2.0/30, 169.254.3.0/30, 169.254.4.0/30, 169.254.5.0/30, 169.254.6.0/30, or 169.254.169.252/30.
	//
	// > The BGP CIDR block of each tunnel under a VPN gateway instance must be unique.
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
	// The authentication algorithm used in Phase 1 negotiation.
	//
	// <props="intl"><ph>Valid values: **md5**, **sha1**, **sha256**, **sha384**, and **sha512**.</ph>
	//
	// <props="china"><ph>If the VPN gateway instance type is standard, valid values are **md5**, **sha1**, **sha256**, **sha384**, and **sha512**.</ph>
	//
	// <props="china"><ph>If the VPN gateway instance type is Chinese SM-based, the only valid value is **sm3**.</ph>
	//
	// example:
	//
	// md5
	IkeAuthAlg *string `json:"IkeAuthAlg,omitempty" xml:"IkeAuthAlg,omitempty"`
	// The encryption algorithm used in Phase 1 negotiation.
	//
	// <props="intl"><ph>Valid values: **aes**, **aes192**, **aes256**, **des**, and **3des**.</ph>
	//
	// <props="china"><ph>If the VPN gateway instance type is standard, valid values are **aes**, **aes192**, **aes256**, **des**, and **3des**.</ph>
	//
	// <props="china"><ph>If the VPN gateway instance type is Chinese SM-based, the only valid value is **sm4**.</ph>
	//
	// example:
	//
	// aes
	IkeEncAlg *string `json:"IkeEncAlg,omitempty" xml:"IkeEncAlg,omitempty"`
	// The SA lifetime determined by Phase 1 negotiation. Unit: seconds. Valid values: **0*	- to **86400**.
	//
	// example:
	//
	// 86400
	IkeLifetime *int64 `json:"IkeLifetime,omitempty" xml:"IkeLifetime,omitempty"`
	// The negotiation mode of the IKE version. Valid values:
	//
	// - **main**: Main mode. This mode offers high security during negotiation.
	//
	// - **aggressive**: Aggressive mode. This mode supports fast negotiation and a higher success rate.
	//
	// <props="china"><ph>If the VPN gateway instance type is Chinese SM-based, only **main*	- is supported for the negotiation mode.</ph>
	//
	// example:
	//
	// main
	IkeMode *string `json:"IkeMode,omitempty" xml:"IkeMode,omitempty"`
	// The Diffie-Hellman key exchange algorithm used in Phase 1 negotiation. Valid values: **group1**, **group2**, **group5**, and **group14**.
	//
	// example:
	//
	// group2
	IkePfs *string `json:"IkePfs,omitempty" xml:"IkePfs,omitempty"`
	// The version of the IKE protocol. Valid values: **ikev1*	- and **ikev2**.
	//
	// Compared with IKEv1, IKEv2 simplifies the SA negotiation process and provides better support for multi-CIDR-block scenarios.
	//
	// <props="china"><ph>If the VPN gateway instance type is Chinese SM-based, only **ikev1*	- is supported for the IKE version.</ph>
	//
	// example:
	//
	// ikev1
	IkeVersion *string `json:"IkeVersion,omitempty" xml:"IkeVersion,omitempty"`
	// The identifier of the tunnel local side (Alibaba Cloud side), which is used in Phase 1 negotiation. The identifier can be up to 100 characters in length and cannot contain spaces. The default value is the IP address of the tunnel.
	//
	// **LocalId*	- supports the FQDN format. If you use the FQDN format, set the negotiation mode to **aggressive**.
	//
	// example:
	//
	// 47.21.XX.XX
	LocalId *string `json:"LocalId,omitempty" xml:"LocalId,omitempty"`
	// The pre-shared key used for identity authentication between the tunnel and the tunnel peer.
	//
	// - The key must be 1 to 100 characters in length and can contain digits, letters, and the following characters. It cannot contain spaces. ```~!\\`@#$%^&*()_-+={}[]|;:\\",.<>/?```
	//
	// - If you do not specify a pre-shared key, the system generates a random 16-character string as the pre-shared key. You can call the [DescribeVpnConnection](https://help.aliyun.com/document_detail/2526951.html) operation to query the pre-shared key that is automatically generated by the system.
	//
	// > The pre-shared key of the tunnel must be the same as that of the tunnel peer. Otherwise, the tunnel cannot be established.
	//
	// example:
	//
	// 123456****
	Psk *string `json:"Psk,omitempty" xml:"Psk,omitempty"`
	// The identifier of the tunnel peer, which is used in Phase 1 negotiation. The identifier can be up to 100 characters in length and cannot contain spaces. The default value is the IP address of the customer gateway associated with the tunnel.
	//
	// **RemoteId*	- supports the FQDN format. If you use the FQDN format, set the negotiation mode to **aggressive**.
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
	// The authentication algorithm used in Phase 2 negotiation.
	//
	// <props="intl"><ph>Valid values: **md5**, **sha1**, **sha256**, **sha384**, and **sha512**.</ph>
	//
	// <props="china"><ph>If the VPN gateway instance type is standard, valid values are **md5**, **sha1**, **sha256**, **sha384**, and **sha512**.</ph>
	//
	// <props="china"><ph>If the VPN gateway instance type is Chinese SM-based, the only valid value is **sm3**.</ph>
	//
	// example:
	//
	// md5
	IpsecAuthAlg *string `json:"IpsecAuthAlg,omitempty" xml:"IpsecAuthAlg,omitempty"`
	// The encryption algorithm used in Phase 2 negotiation.
	//
	// <props="intl"><ph>Valid values: **aes**, **aes192**, **aes256**, **des**, and **3des**.</ph>
	//
	// <props="china"><ph>If the VPN gateway instance type is standard, valid values are **aes**, **aes192**, **aes256**, **des**, and **3des**.</ph>
	//
	// <props="china"><ph>If the VPN gateway instance type is Chinese SM-based, the only valid value is **sm4**.</ph>
	//
	// example:
	//
	// aes
	IpsecEncAlg *string `json:"IpsecEncAlg,omitempty" xml:"IpsecEncAlg,omitempty"`
	// The SA lifetime determined by Phase 2 negotiation. Unit: seconds. Valid values: **0*	- to **86400**.
	//
	// example:
	//
	// 86400
	IpsecLifetime *int32 `json:"IpsecLifetime,omitempty" xml:"IpsecLifetime,omitempty"`
	// The Diffie-Hellman key exchange algorithm used in Phase 2 negotiation.
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
