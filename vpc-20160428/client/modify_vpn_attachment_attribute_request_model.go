// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpnAttachmentAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoConfigRoute(v bool) *ModifyVpnAttachmentAttributeRequest
	GetAutoConfigRoute() *bool
	SetBgpConfig(v string) *ModifyVpnAttachmentAttributeRequest
	GetBgpConfig() *string
	SetClientToken(v string) *ModifyVpnAttachmentAttributeRequest
	GetClientToken() *string
	SetCustomerGatewayId(v string) *ModifyVpnAttachmentAttributeRequest
	GetCustomerGatewayId() *string
	SetEffectImmediately(v bool) *ModifyVpnAttachmentAttributeRequest
	GetEffectImmediately() *bool
	SetEnableDpd(v bool) *ModifyVpnAttachmentAttributeRequest
	GetEnableDpd() *bool
	SetEnableNatTraversal(v bool) *ModifyVpnAttachmentAttributeRequest
	GetEnableNatTraversal() *bool
	SetEnableTunnelsBgp(v bool) *ModifyVpnAttachmentAttributeRequest
	GetEnableTunnelsBgp() *bool
	SetHealthCheckConfig(v string) *ModifyVpnAttachmentAttributeRequest
	GetHealthCheckConfig() *string
	SetIkeConfig(v string) *ModifyVpnAttachmentAttributeRequest
	GetIkeConfig() *string
	SetIpsecConfig(v string) *ModifyVpnAttachmentAttributeRequest
	GetIpsecConfig() *string
	SetLocalSubnet(v string) *ModifyVpnAttachmentAttributeRequest
	GetLocalSubnet() *string
	SetName(v string) *ModifyVpnAttachmentAttributeRequest
	GetName() *string
	SetNetworkType(v string) *ModifyVpnAttachmentAttributeRequest
	GetNetworkType() *string
	SetOwnerAccount(v string) *ModifyVpnAttachmentAttributeRequest
	GetOwnerAccount() *string
	SetRegionId(v string) *ModifyVpnAttachmentAttributeRequest
	GetRegionId() *string
	SetRemoteCaCert(v string) *ModifyVpnAttachmentAttributeRequest
	GetRemoteCaCert() *string
	SetRemoteSubnet(v string) *ModifyVpnAttachmentAttributeRequest
	GetRemoteSubnet() *string
	SetResourceOwnerAccount(v string) *ModifyVpnAttachmentAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyVpnAttachmentAttributeRequest
	GetResourceOwnerId() *int64
	SetTunnelOptionsSpecification(v []*ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecification) *ModifyVpnAttachmentAttributeRequest
	GetTunnelOptionsSpecification() []*ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecification
	SetVpnConnectionId(v string) *ModifyVpnAttachmentAttributeRequest
	GetVpnConnectionId() *string
}

type ModifyVpnAttachmentAttributeRequest struct {
	// Specifies whether to automatically configure routes. Valid values:
	//
	// - **true**: automatically configures routes.
	//
	// - **false**: does not automatically configure routes.
	//
	// example:
	//
	// true
	AutoConfigRoute *bool `json:"AutoConfigRoute,omitempty" xml:"AutoConfigRoute,omitempty"`
	// This parameter is supported when you modify an IPsec-VPN connection in single-tunnel mode.
	//
	// BGP configuration:
	//
	// - **BgpConfig.EnableBgp**: specifies whether to enable BGP. Valid values:
	//
	//     - **true**: Enable BGP.
	//
	//     - **false**: Disable BGP.
	//
	// - **BgpConfig.LocalAsn**: the autonomous system number (ASN) on the Alibaba Cloud side. Valid values: **1*	- to **4294967295**.
	//
	//     You can enter the ASN in the two-segment notation: the first 16 bits.the last 16 bits. Each segment is entered in decimal format.
	//
	//     For example, if you enter 123.456, the ASN is 123 × 65536 + 456 = 8061384.
	//
	// - **BgpConfig.TunnelCidr**: the CIDR block of the IPsec tunnel. The CIDR block must be a /30 CIDR block within 169.254.0.0/16 and cannot be 169.254.0.0/30, 169.254.1.0/30, 169.254.2.0/30, 169.254.3.0/30, 169.254.4.0/30, 169.254.5.0/30, 169.254.6.0/30, or 169.254.169.252/30.
	//
	// - **LocalBgpIp**: the BGP IP address on the Alibaba Cloud side. This address must be an IP address within the CIDR block of the IPsec tunnel.
	//
	// >- Before you configure BGP, we recommend that you learn about how BGP dynamic routing works and its limits. For more information, see [Configure BGP dynamic routing](https://help.aliyun.com/document_detail/445767.html).
	//
	// >- We recommend that you use a private ASN to establish a BGP connection with Alibaba Cloud. For the range of private ASNs, refer to the relevant documentation.
	//
	// example:
	//
	// {"EnableBgp":"true","LocalAsn":"45104","TunnelCidr":"169.254.11.0/30","LocalBgpIp":"169.254.11.1"}
	BgpConfig *string `json:"BgpConfig,omitempty" xml:"BgpConfig,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-4266****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The customer gateway instance associated with the IPsec-VPN connection.
	//
	// example:
	//
	// cgw-p0w2jemrcj5u61un8****
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" xml:"CustomerGatewayId,omitempty"`
	// Specifies whether the configuration of the IPsec-VPN connection takes effect immediately. Valid values:
	//
	// example:
	//
	// false
	EffectImmediately *bool `json:"EffectImmediately,omitempty" xml:"EffectImmediately,omitempty"`
	// This parameter is supported when you modify a single-tunnel IPsec-VPN connection.
	//
	// Specifies whether to enable the Dead Peer Detection (DPD) feature. Valid values:
	//
	// - **true**: Enables the DPD feature. The IPsec initiator sends DPD packets to check whether the peer device is alive. If no correct response is received within the specified period of time, the peer is considered disconnected. The ISAKMP SA and the corresponding IPsec SA are deleted, and the security tunnel is also deleted.
	//
	// - **false**: Disables the DPD feature. The IPsec initiator does not send DPD probe packets.
	//
	// example:
	//
	// true
	EnableDpd *bool `json:"EnableDpd,omitempty" xml:"EnableDpd,omitempty"`
	// This parameter is supported when you modify an IPsec-VPN connection in single-tunnel mode.
	//
	// Specifies whether to enable NAT traversal. Valid values:
	//
	// - **true**: Enable NAT traversal. After NAT traversal is enabled, the IKE negotiation process skips UDP port number verification and can discover NAT gateway devices in the VPN tunnel.
	//
	// - **false**: Disable NAT traversal.
	//
	// example:
	//
	// true
	EnableNatTraversal *bool `json:"EnableNatTraversal,omitempty" xml:"EnableNatTraversal,omitempty"`
	// This parameter is supported when you modify an IPsec-VPN connection in dual-tunnel mode.
	//
	// Specifies whether to enable BGP for the tunnel. Valid values: **true*	- or **false**.
	//
	// > Before you add BGP configurations, we recommend that you understand the working mechanism and limits of BGP dynamic routing. For more information, see [Configure BGP dynamic routing](https://help.aliyun.com/document_detail/445767.html).
	//
	// if can be null:
	// true
	//
	// example:
	//
	// false
	EnableTunnelsBgp *bool `json:"EnableTunnelsBgp,omitempty" xml:"EnableTunnelsBgp,omitempty"`
	// This parameter is supported when you modify a single-tunnel IPsec-VPN connection.
	//
	// example:
	//
	// {"enable":"true","dip":"192.168.1.1","sip":"10.1.1.1","interval":"3","retry":"3","Policy": "revoke_route"}
	HealthCheckConfig *string `json:"HealthCheckConfig,omitempty" xml:"HealthCheckConfig,omitempty"`
	// This parameter is supported when you modify a single-tunnel IPsec-VPN connection.
	//
	// example:
	//
	// {"Psk":"1234****","IkeVersion":"ikev1","IkeMode":"main","IkeEncAlg":"aes","IkeAuthAlg":"sha1","IkePfs":"group2","IkeLifetime":86400,"LocalId":"47.XX.XX.1","RemoteId":"47.XX.XX.2"}
	IkeConfig *string `json:"IkeConfig,omitempty" xml:"IkeConfig,omitempty"`
	// This parameter is supported when you modify an IPsec-VPN connection in single-tunnel mode.
	//
	// The configuration of Phase 2 negotiation:
	//
	// - **IpsecConfig.IpsecEncAlg**: The encryption algorithm used in Phase 2 negotiation. Valid values: **aes**, **aes192**, **aes256**, **des**, or **3des**.
	//
	// - **IpsecConfig.IpsecAuthAlg**: The authentication algorithm used in Phase 2 negotiation. Valid values: **md5**, **sha1**, **sha256**, **sha384**, or **sha512**.
	//
	// - **IpsecConfig.IpsecPfs**: The Diffie-Hellman key exchange algorithm used in Phase 2 negotiation. Valid values: **disabled**, **group1**, **group2**, **group5**, or **group14**.
	//
	// - **IpsecConfig.IpsecLifetime**: The lifetime of the SA negotiated in Phase 2. Unit: seconds. Valid values: **0*	- to **86400**.
	//
	// example:
	//
	// {"IpsecEncAlg":"aes","IpsecAuthAlg":"sha1","IpsecPfs":"group2","IpsecLifetime":86400}
	IpsecConfig *string `json:"IpsecConfig,omitempty" xml:"IpsecConfig,omitempty"`
	// The CIDR block on the VPC side that needs to communicate with the on-premises data center. This is used for Phase 2 negotiation.
	//
	// Separate multiple CIDR blocks with commas (,). Example: 192.168.1.0/24,192.168.2.0/24.
	//
	// Description of the IPsec-VPN connection routing mode:
	//
	// - If both **LocalSubnet*	- and **RemoteSubnet*	- are set to 0.0.0.0/0, the destination routing mode is used.
	//
	// - If both **LocalSubnet*	- and **RemoteSubnet*	- are set to specific CIDR blocks, the protected data flow mode is used.
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
	// - **public**: The IPsec-VPN connection establishes an encrypted communication channel over the Internet.
	//
	// - **private**: The IPsec-VPN connection establishes an encrypted communication channel over a private network.
	//
	// example:
	//
	// public
	NetworkType  *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	// The region ID of the IPsec-VPN connection.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The CA certificate of the peer.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- MIIB7zCCAZW***	- -----END CERTIFICATE-----
	RemoteCaCert *string `json:"RemoteCaCert,omitempty" xml:"RemoteCaCert,omitempty"`
	// The CIDR block on the on-premises data center side that needs to communicate with the VPC. This parameter is used for Phase 2 negotiation.
	//
	// example:
	//
	// 10.1.3.0/24,10.1.4.0/24
	RemoteSubnet         *string `json:"RemoteSubnet,omitempty" xml:"RemoteSubnet,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tunnel configurations.
	//
	// The parameters in the **TunnelOptionsSpecification*	- array are supported only when you modify an IPsec-VPN connection in dual-tunnel mode. You can modify the configurations of both tunnels of the IPsec-VPN connection at the same time.
	//
	// if can be null:
	// true
	TunnelOptionsSpecification []*ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecification `json:"TunnelOptionsSpecification,omitempty" xml:"TunnelOptionsSpecification,omitempty" type:"Repeated"`
	// The ID of the IPsec-VPN connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// vco-p0w5112fgnl2ihlmf****
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" xml:"VpnConnectionId,omitempty"`
}

func (s ModifyVpnAttachmentAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnAttachmentAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpnAttachmentAttributeRequest) GetAutoConfigRoute() *bool {
	return s.AutoConfigRoute
}

func (s *ModifyVpnAttachmentAttributeRequest) GetBgpConfig() *string {
	return s.BgpConfig
}

func (s *ModifyVpnAttachmentAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyVpnAttachmentAttributeRequest) GetCustomerGatewayId() *string {
	return s.CustomerGatewayId
}

func (s *ModifyVpnAttachmentAttributeRequest) GetEffectImmediately() *bool {
	return s.EffectImmediately
}

func (s *ModifyVpnAttachmentAttributeRequest) GetEnableDpd() *bool {
	return s.EnableDpd
}

func (s *ModifyVpnAttachmentAttributeRequest) GetEnableNatTraversal() *bool {
	return s.EnableNatTraversal
}

func (s *ModifyVpnAttachmentAttributeRequest) GetEnableTunnelsBgp() *bool {
	return s.EnableTunnelsBgp
}

func (s *ModifyVpnAttachmentAttributeRequest) GetHealthCheckConfig() *string {
	return s.HealthCheckConfig
}

func (s *ModifyVpnAttachmentAttributeRequest) GetIkeConfig() *string {
	return s.IkeConfig
}

func (s *ModifyVpnAttachmentAttributeRequest) GetIpsecConfig() *string {
	return s.IpsecConfig
}

func (s *ModifyVpnAttachmentAttributeRequest) GetLocalSubnet() *string {
	return s.LocalSubnet
}

func (s *ModifyVpnAttachmentAttributeRequest) GetName() *string {
	return s.Name
}

func (s *ModifyVpnAttachmentAttributeRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *ModifyVpnAttachmentAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyVpnAttachmentAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyVpnAttachmentAttributeRequest) GetRemoteCaCert() *string {
	return s.RemoteCaCert
}

func (s *ModifyVpnAttachmentAttributeRequest) GetRemoteSubnet() *string {
	return s.RemoteSubnet
}

func (s *ModifyVpnAttachmentAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyVpnAttachmentAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyVpnAttachmentAttributeRequest) GetTunnelOptionsSpecification() []*ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecification {
	return s.TunnelOptionsSpecification
}

func (s *ModifyVpnAttachmentAttributeRequest) GetVpnConnectionId() *string {
	return s.VpnConnectionId
}

func (s *ModifyVpnAttachmentAttributeRequest) SetAutoConfigRoute(v bool) *ModifyVpnAttachmentAttributeRequest {
	s.AutoConfigRoute = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequest) SetBgpConfig(v string) *ModifyVpnAttachmentAttributeRequest {
	s.BgpConfig = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequest) SetClientToken(v string) *ModifyVpnAttachmentAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequest) SetCustomerGatewayId(v string) *ModifyVpnAttachmentAttributeRequest {
	s.CustomerGatewayId = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequest) SetEffectImmediately(v bool) *ModifyVpnAttachmentAttributeRequest {
	s.EffectImmediately = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequest) SetEnableDpd(v bool) *ModifyVpnAttachmentAttributeRequest {
	s.EnableDpd = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequest) SetEnableNatTraversal(v bool) *ModifyVpnAttachmentAttributeRequest {
	s.EnableNatTraversal = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequest) SetEnableTunnelsBgp(v bool) *ModifyVpnAttachmentAttributeRequest {
	s.EnableTunnelsBgp = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequest) SetHealthCheckConfig(v string) *ModifyVpnAttachmentAttributeRequest {
	s.HealthCheckConfig = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequest) SetIkeConfig(v string) *ModifyVpnAttachmentAttributeRequest {
	s.IkeConfig = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequest) SetIpsecConfig(v string) *ModifyVpnAttachmentAttributeRequest {
	s.IpsecConfig = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequest) SetLocalSubnet(v string) *ModifyVpnAttachmentAttributeRequest {
	s.LocalSubnet = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequest) SetName(v string) *ModifyVpnAttachmentAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequest) SetNetworkType(v string) *ModifyVpnAttachmentAttributeRequest {
	s.NetworkType = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequest) SetOwnerAccount(v string) *ModifyVpnAttachmentAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequest) SetRegionId(v string) *ModifyVpnAttachmentAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequest) SetRemoteCaCert(v string) *ModifyVpnAttachmentAttributeRequest {
	s.RemoteCaCert = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequest) SetRemoteSubnet(v string) *ModifyVpnAttachmentAttributeRequest {
	s.RemoteSubnet = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequest) SetResourceOwnerAccount(v string) *ModifyVpnAttachmentAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequest) SetResourceOwnerId(v int64) *ModifyVpnAttachmentAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequest) SetTunnelOptionsSpecification(v []*ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecification) *ModifyVpnAttachmentAttributeRequest {
	s.TunnelOptionsSpecification = v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequest) SetVpnConnectionId(v string) *ModifyVpnAttachmentAttributeRequest {
	s.VpnConnectionId = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequest) Validate() error {
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

type ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecification struct {
	// The ID of the customer gateway associated with the tunnel.
	//
	// example:
	//
	// cgw-p0w2jemrcj5u61un8****
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" xml:"CustomerGatewayId,omitempty"`
	// Specifies whether to enable the Dead Peer Detection (DPD) feature for the tunnel. Valid values:
	//
	// example:
	//
	// true
	EnableDpd *bool `json:"EnableDpd,omitempty" xml:"EnableDpd,omitempty"`
	// Specifies whether to enable NAT traversal for the tunnel. Valid values:
	//
	// - **true**: Enables NAT traversal. After NAT traversal is enabled, the IKE negotiation process skips UDP port number verification and can discover NAT gateway devices along the tunnel.
	//
	// - **false**: Disables NAT traversal.
	//
	// example:
	//
	// true
	EnableNatTraversal *bool `json:"EnableNatTraversal,omitempty" xml:"EnableNatTraversal,omitempty"`
	// Adds BGP configurations for the tunnel.
	//
	// > Configure this parameter after you enable BGP for the IPsec-VPN connection (that is, set **EnableTunnelsBgp*	- to **true**).
	TunnelBgpConfig *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig `json:"TunnelBgpConfig,omitempty" xml:"TunnelBgpConfig,omitempty" type:"Struct"`
	// The tunnel ID.
	//
	// example:
	//
	// tun-0jod7plwf2a0o9lvu****
	TunnelId *string `json:"TunnelId,omitempty" xml:"TunnelId,omitempty"`
	// The Phase 1 negotiation configuration.
	TunnelIkeConfig *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig `json:"TunnelIkeConfig,omitempty" xml:"TunnelIkeConfig,omitempty" type:"Struct"`
	// The creation order of the tunnel.
	//
	// example:
	//
	// 1
	TunnelIndex *int32 `json:"TunnelIndex,omitempty" xml:"TunnelIndex,omitempty"`
	// The Phase 2 negotiation configuration.
	TunnelIpsecConfig *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig `json:"TunnelIpsecConfig,omitempty" xml:"TunnelIpsecConfig,omitempty" type:"Struct"`
}

func (s ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecification) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecification) GoString() string {
	return s.String()
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecification) GetCustomerGatewayId() *string {
	return s.CustomerGatewayId
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecification) GetEnableDpd() *bool {
	return s.EnableDpd
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecification) GetEnableNatTraversal() *bool {
	return s.EnableNatTraversal
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecification) GetTunnelBgpConfig() *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig {
	return s.TunnelBgpConfig
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecification) GetTunnelId() *string {
	return s.TunnelId
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecification) GetTunnelIkeConfig() *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig {
	return s.TunnelIkeConfig
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecification) GetTunnelIndex() *int32 {
	return s.TunnelIndex
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecification) GetTunnelIpsecConfig() *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig {
	return s.TunnelIpsecConfig
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecification) SetCustomerGatewayId(v string) *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecification {
	s.CustomerGatewayId = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecification) SetEnableDpd(v bool) *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecification {
	s.EnableDpd = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecification) SetEnableNatTraversal(v bool) *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecification {
	s.EnableNatTraversal = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecification) SetTunnelBgpConfig(v *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig) *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecification {
	s.TunnelBgpConfig = v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecification) SetTunnelId(v string) *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecification {
	s.TunnelId = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecification) SetTunnelIkeConfig(v *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecification {
	s.TunnelIkeConfig = v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecification) SetTunnelIndex(v int32) *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecification {
	s.TunnelIndex = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecification) SetTunnelIpsecConfig(v *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecification {
	s.TunnelIpsecConfig = v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecification) Validate() error {
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

type ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig struct {
	// The autonomous system number (ASN) on the Alibaba Cloud side of the tunnel. Valid values: **1*	- to **4294967295**. Default value: **45104**.
	//
	// > Use a private ASN to establish a BGP connection with Alibaba Cloud. For the range of private ASNs, refer to the relevant documentation.
	//
	// example:
	//
	// 65530
	LocalAsn *int64 `json:"LocalAsn,omitempty" xml:"LocalAsn,omitempty"`
	// The BGP address on the Alibaba Cloud side. This address is an IP address within the BGP CIDR block.
	//
	// example:
	//
	// 169.254.10.1
	LocalBgpIp *string `json:"LocalBgpIp,omitempty" xml:"LocalBgpIp,omitempty"`
	// The BGP CIDR block of the tunnel. The CIDR block must be a CIDR block with a mask length of 30 within 169.254.0.0/16 and cannot be 169.254.0.0/30, 169.254.1.0/30, 169.254.2.0/30, 169.254.3.0/30, 169.254.4.0/30, 169.254.5.0/30, 169.254.6.0/30, or 169.254.169.252/30.
	//
	// example:
	//
	// 169.254.10.0/30
	TunnelCidr *string `json:"TunnelCidr,omitempty" xml:"TunnelCidr,omitempty"`
}

func (s ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig) GoString() string {
	return s.String()
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig) GetLocalAsn() *int64 {
	return s.LocalAsn
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig) GetLocalBgpIp() *string {
	return s.LocalBgpIp
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig) GetTunnelCidr() *string {
	return s.TunnelCidr
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig) SetLocalAsn(v int64) *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig {
	s.LocalAsn = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig) SetLocalBgpIp(v string) *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig {
	s.LocalBgpIp = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig) SetTunnelCidr(v string) *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig {
	s.TunnelCidr = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig struct {
	// The authentication algorithm for Phase 1 negotiation. Valid values: **md5**, **sha1**, **sha256**, **sha384**, and **sha512**.
	//
	// example:
	//
	// sha1
	IkeAuthAlg *string `json:"IkeAuthAlg,omitempty" xml:"IkeAuthAlg,omitempty"`
	// The encryption algorithm for Phase 1 negotiation. Valid values: **aes**, **aes192**, **aes256**, **des**, and **3des**.
	//
	// example:
	//
	// aes
	IkeEncAlg *string `json:"IkeEncAlg,omitempty" xml:"IkeEncAlg,omitempty"`
	// The lifetime of the SA generated by Phase 1 negotiation. Unit: seconds.
	//
	// example:
	//
	// 86400
	IkeLifetime *int64 `json:"IkeLifetime,omitempty" xml:"IkeLifetime,omitempty"`
	// The negotiation mode of the IKE version. Valid values: **main*	- and **aggressive**.
	//
	// - **main**: Main mode. The negotiation process is highly secure.
	//
	// - **aggressive**: Aggressive mode. The negotiation is fast and has a high success rate.
	//
	// example:
	//
	// main
	IkeMode *string `json:"IkeMode,omitempty" xml:"IkeMode,omitempty"`
	// The Diffie-Hellman key exchange algorithm used in the first-phase negotiation.
	//
	// Valid values: **group1**, **group2**, **group5**, **group14**.
	//
	// example:
	//
	// group2
	IkePfs *string `json:"IkePfs,omitempty" xml:"IkePfs,omitempty"`
	// The version of the IKE protocol. Valid values: **ikev1*	- and **ikev2**.
	//
	// example:
	//
	// ikev2
	IkeVersion *string `json:"IkeVersion,omitempty" xml:"IkeVersion,omitempty"`
	// The identifier on the Alibaba Cloud side for the tunnel, used for Phase 1 negotiation. The value can be up to 100 characters in length and cannot contain spaces.
	//
	// example:
	//
	// 47.XX.XX.1
	LocalId *string `json:"LocalId,omitempty" xml:"LocalId,omitempty"`
	// The pre-shared key, which is used for identity authentication between the tunnel and the tunnel peer.
	//
	// - The key must be 1 to 100 characters in length and can contain digits, uppercase letters, lowercase letters, and the following characters. It cannot contain spaces. ```~!\\`@#$%^&*()_-+={}[]|;:\\",.<>/?```
	//
	// - If you do not specify a pre-shared key, the system randomly generates a 16-character string as the pre-shared key. You can call the [DescribeVpnAttachments](https://help.aliyun.com/document_detail/2526939.html) operation to query the pre-shared key automatically generated by the system.
	//
	// > The pre-shared keys of the tunnel and the tunnel peer must be the same. Otherwise, the tunnel cannot be established.
	//
	// example:
	//
	// 123456****
	Psk *string `json:"Psk,omitempty" xml:"Psk,omitempty"`
	// The identifier of the tunnel peer, used for Phase 1 negotiation. The value can be up to 100 characters in length and cannot contain spaces.
	//
	// example:
	//
	// 47.XX.XX.2
	RemoteId *string `json:"RemoteId,omitempty" xml:"RemoteId,omitempty"`
}

func (s ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) GoString() string {
	return s.String()
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) GetIkeAuthAlg() *string {
	return s.IkeAuthAlg
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) GetIkeEncAlg() *string {
	return s.IkeEncAlg
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) GetIkeLifetime() *int64 {
	return s.IkeLifetime
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) GetIkeMode() *string {
	return s.IkeMode
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) GetIkePfs() *string {
	return s.IkePfs
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) GetIkeVersion() *string {
	return s.IkeVersion
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) GetLocalId() *string {
	return s.LocalId
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) GetPsk() *string {
	return s.Psk
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) GetRemoteId() *string {
	return s.RemoteId
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) SetIkeAuthAlg(v string) *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkeAuthAlg = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) SetIkeEncAlg(v string) *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkeEncAlg = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) SetIkeLifetime(v int64) *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkeLifetime = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) SetIkeMode(v string) *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkeMode = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) SetIkePfs(v string) *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkePfs = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) SetIkeVersion(v string) *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkeVersion = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) SetLocalId(v string) *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.LocalId = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) SetPsk(v string) *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.Psk = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) SetRemoteId(v string) *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.RemoteId = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig struct {
	// The authentication algorithm used in Phase 2 negotiations.
	//
	// Valid values: **md5**, **sha1**, **sha256**, **sha384**, **sha512**.
	//
	// example:
	//
	// sha1
	IpsecAuthAlg *string `json:"IpsecAuthAlg,omitempty" xml:"IpsecAuthAlg,omitempty"`
	// The encryption algorithm for Phase 2 negotiation. Valid values: **aes**, **aes192**, **aes256**, **des**, and **3des**.
	//
	// example:
	//
	// aes
	IpsecEncAlg *string `json:"IpsecEncAlg,omitempty" xml:"IpsecEncAlg,omitempty"`
	// The lifetime of the SA generated by Phase 2 negotiation. Unit: seconds.
	//
	// example:
	//
	// 86400
	IpsecLifetime *int32 `json:"IpsecLifetime,omitempty" xml:"IpsecLifetime,omitempty"`
	// The Diffie-Hellman key exchange algorithm used in the second phase of negotiation.
	//
	// Valid values: **disabled**, **group1**, **group2**, **group5**, **group14**.
	//
	// example:
	//
	// group2
	IpsecPfs *string `json:"IpsecPfs,omitempty" xml:"IpsecPfs,omitempty"`
}

func (s ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) GoString() string {
	return s.String()
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) GetIpsecAuthAlg() *string {
	return s.IpsecAuthAlg
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) GetIpsecEncAlg() *string {
	return s.IpsecEncAlg
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) GetIpsecLifetime() *int32 {
	return s.IpsecLifetime
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) GetIpsecPfs() *string {
	return s.IpsecPfs
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) SetIpsecAuthAlg(v string) *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig {
	s.IpsecAuthAlg = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) SetIpsecEncAlg(v string) *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig {
	s.IpsecEncAlg = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) SetIpsecLifetime(v int32) *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig {
	s.IpsecLifetime = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) SetIpsecPfs(v string) *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig {
	s.IpsecPfs = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) Validate() error {
	return dara.Validate(s)
}
