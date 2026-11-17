// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpnConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoConfigRoute(v bool) *CreateVpnConnectionRequest
	GetAutoConfigRoute() *bool
	SetBgpConfig(v string) *CreateVpnConnectionRequest
	GetBgpConfig() *string
	SetClientToken(v string) *CreateVpnConnectionRequest
	GetClientToken() *string
	SetCustomerGatewayId(v string) *CreateVpnConnectionRequest
	GetCustomerGatewayId() *string
	SetDryRun(v bool) *CreateVpnConnectionRequest
	GetDryRun() *bool
	SetEffectImmediately(v bool) *CreateVpnConnectionRequest
	GetEffectImmediately() *bool
	SetEnableDpd(v bool) *CreateVpnConnectionRequest
	GetEnableDpd() *bool
	SetEnableNatTraversal(v bool) *CreateVpnConnectionRequest
	GetEnableNatTraversal() *bool
	SetEnableTunnelsBgp(v bool) *CreateVpnConnectionRequest
	GetEnableTunnelsBgp() *bool
	SetHealthCheckConfig(v string) *CreateVpnConnectionRequest
	GetHealthCheckConfig() *string
	SetIkeConfig(v string) *CreateVpnConnectionRequest
	GetIkeConfig() *string
	SetIpsecConfig(v string) *CreateVpnConnectionRequest
	GetIpsecConfig() *string
	SetLocalSubnet(v string) *CreateVpnConnectionRequest
	GetLocalSubnet() *string
	SetName(v string) *CreateVpnConnectionRequest
	GetName() *string
	SetOwnerAccount(v string) *CreateVpnConnectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateVpnConnectionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateVpnConnectionRequest
	GetRegionId() *string
	SetRemoteCaCertificate(v string) *CreateVpnConnectionRequest
	GetRemoteCaCertificate() *string
	SetRemoteSubnet(v string) *CreateVpnConnectionRequest
	GetRemoteSubnet() *string
	SetResourceOwnerAccount(v string) *CreateVpnConnectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateVpnConnectionRequest
	GetResourceOwnerId() *int64
	SetTags(v []*CreateVpnConnectionRequestTags) *CreateVpnConnectionRequest
	GetTags() []*CreateVpnConnectionRequestTags
	SetTunnelOptionsSpecification(v []*CreateVpnConnectionRequestTunnelOptionsSpecification) *CreateVpnConnectionRequest
	GetTunnelOptionsSpecification() []*CreateVpnConnectionRequestTunnelOptionsSpecification
	SetVpnGatewayId(v string) *CreateVpnConnectionRequest
	GetVpnGatewayId() *string
}

type CreateVpnConnectionRequest struct {
	// Specifies whether to automatically configure routes. Valid values:
	//
	// - **true*	- (default): Automatically configures routes.
	//
	// - **false**: Does not automatically configure routes.
	//
	// example:
	//
	// true
	AutoConfigRoute *bool `json:"AutoConfigRoute,omitempty" xml:"AutoConfigRoute,omitempty"`
	// This parameter is supported when you create an IPsec-VPN connection in single-tunnel mode.
	//
	// The BGP configurations:
	//
	// - **BgpConfig.EnableBgp**: Specifies whether to enable the BGP feature. Valid values: **true*	- and **false*	- (default).
	//
	// - **BgpConfig.LocalAsn**: The autonomous system number (ASN) on the Alibaba Cloud side. Valid values: **1*	- to **4294967295**. Default value: **45104**.
	//
	//
	//
	//     You can enter the ASN in the two-segment format: the first 16 bits and the last 16 bits. Each segment must be entered in decimal format.
	//
	//
	//
	//     For example, if you enter 123.456, the ASN is calculated as follows: 123 × 65536 + 456 = 8061384.
	//
	// - **BgpConfig.TunnelCidr**: The CIDR block of the IPsec tunnel. The CIDR block must be within the 169.254.0.0/16 range and have a mask length of 30, and cannot be 169.254.0.0/30, 169.254.1.0/30, 169.254.2.0/30, 169.254.3.0/30, 169.254.4.0/30, 169.254.5.0/30, 169.254.6.0/30, or 169.254.169.252/30.
	//
	//     > The IPsec tunnel CIDR block of each IPsec-VPN connection under the same VPN gateway instance must be unique.
	//
	// - **LocalBgpIp**: The BGP IP address on the Alibaba Cloud side. This IP address must be within the IPsec tunnel CIDR block.
	//
	// >- Before adding BGP configurations, understand the working mechanism and limits of the BGP dynamic routing feature. For more information, refer to [Configure BGP dynamic routing](https://help.aliyun.com/document_detail/2638220.html).
	//
	// >- Use a private ASN to establish a BGP connection with Alibaba Cloud. Refer to relevant documentation for the range of private ASNs.
	//
	// example:
	//
	// {"EnableBgp":"true","LocalAsn":"45104","TunnelCidr":"169.254.11.0/30","LocalBgpIp":"169.254.11.1"}
	BgpConfig *string `json:"BgpConfig,omitempty" xml:"BgpConfig,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is unique.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required when you create an IPsec-VPN connection in single-tunnel mode.
	//
	// The ID of the customer gateway.
	//
	// example:
	//
	// cgw-p0w2jemrcj5u61un8****
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" xml:"CustomerGatewayId,omitempty"`
	// Specifies whether to only precheck the request. Valid values:
	//
	// - **true**: Prechecks the request without creating the IPsec-VPN connection. The system checks whether required parameters are specified, the request format, and service limits. If the request fails the precheck, the corresponding error is returned. If the request passes the precheck, the DryRunOperation error code is returned.
	//
	// - **false*	- (default): Sends a normal request and creates the IPsec-VPN connection after the request passes the precheck.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether the configurations of the IPsec-VPN connection take effect immediately. Valid values:
	//
	//
	//
	// - **true**: Yes. The system immediately initiates IPsec negotiations after the configurations are completed.
	//
	//
	//
	// - **false*	- (default): No. The system initiates IPsec negotiations only when traffic is detected.
	//
	// example:
	//
	// false
	EffectImmediately *bool `json:"EffectImmediately,omitempty" xml:"EffectImmediately,omitempty"`
	// This parameter is supported when you create an IPsec-VPN connection in single-tunnel mode.
	//
	// Specifies whether to enable the dead peer detection (DPD) feature. Valid values:
	//
	// - **true*	- (default): Enables the DPD feature. The IPsec initiator sends DPD packets to check whether the peer device is alive. If no correct response is received within the specified time, the peer is considered disconnected. IPsec deletes the ISAKMP SA and the corresponding IPsec SA, and the secure tunnel is also deleted.
	//
	// - **false**: Disables the DPD feature. The IPsec initiator does not send DPD probe packets.
	//
	// example:
	//
	// true
	EnableDpd *bool `json:"EnableDpd,omitempty" xml:"EnableDpd,omitempty"`
	// This parameter is supported when you create an IPsec-VPN connection in single-tunnel mode.
	//
	// Specifies whether to enable the NAT traversal feature. Valid values:
	//
	// - **true*	- (default): Enables the NAT traversal feature. After this feature is enabled, the IKE negotiation process skips the verification of UDP port numbers and discovers NAT gateway devices in the VPN tunnel.
	//
	// - **false**: Disables the NAT traversal feature.
	//
	// example:
	//
	// true
	EnableNatTraversal *bool `json:"EnableNatTraversal,omitempty" xml:"EnableNatTraversal,omitempty"`
	// This parameter is supported when you create an IPsec-VPN connection in dual-tunnel mode.
	//
	// Specifies whether to enable the BGP feature for the tunnel. Valid values: **true*	- and **false*	- (default).
	//
	// example:
	//
	// true
	EnableTunnelsBgp *bool `json:"EnableTunnelsBgp,omitempty" xml:"EnableTunnelsBgp,omitempty"`
	// This parameter is supported when you create an IPsec-VPN connection in single-tunnel mode.
	//
	// The health check configurations:
	//
	// - **HealthCheckConfig.enable**: Specifies whether to enable health checks. Valid values: **true*	- and **false*	- (default).
	//
	// - **HealthCheckConfig.dip**: The destination IP address for health checks.
	//
	// - **HealthCheckConfig.sip**: The source IP address for health checks.
	//
	// - **HealthCheckConfig.interval**: The retry interval for health checks. Unit: seconds. Default value: **3**.
	//
	// - **HealthCheckConfig.retry**: The maximum number of health check retries. Default value: **3**.
	//
	// example:
	//
	// {"enable":"true","dip":"192.168.10.1","sip":"10.10.1.1","interval":"3","retry":"3"}
	HealthCheckConfig *string `json:"HealthCheckConfig,omitempty" xml:"HealthCheckConfig,omitempty"`
	// This parameter is supported when you create an IPsec-VPN connection in single-tunnel mode.
	//
	// The configurations for Phase 1 negotiations:
	//
	//
	//
	// - **IkeConfig.Psk**: The pre-shared key used for authentication between the VPN gateway and the data center.
	//
	//     - The key must be 1 to 100 characters in length, and can contain digits, letters, and the following special characters: ```~!\\`@#$%^&*()_-+={}[]|;:\\",.<>/?``` The key cannot contain spaces.
	//
	//     - If you do not specify a pre-shared key, the system randomly generates a string as the pre-shared key. You can call the [DescribeVpnConnection](https://help.aliyun.com/document_detail/2526951.html) operation to query the automatically generated pre-shared key.
	//
	//         > The pre-shared key on the IPsec-VPN connection side must be the same as the authentication key on the data center side. Otherwise, the connection between the data center and the VPN gateway cannot be established.
	//
	// - **IkeConfig.IkeVersion**: The version of the IKE protocol. Valid values: **ikev1*	- and **ikev2**. Default value: **ikev1**.
	//
	//     Compared with IKEv1, IKEv2 simplifies the SA negotiation process and provides better support for scenarios with multiple CIDR blocks.
	//
	//    <props="china"><ph>If the VPN gateway instance is of the ShangMi (SM) type, only **ikev1*	- is supported.</ph>
	//
	// - **IkeConfig.IkeMode**: The negotiation mode of the IKE version. Valid values: **main*	- and **aggressive**. Default value: **main**.
	//
	//     - **main**: The main mode. The negotiation process is highly secure.
	//
	//     - **aggressive**: The aggressive mode. The negotiation process is fast and has a high success rate.
	//
	//    <props="china"><ph>If the VPN gateway instance is of the ShangMi (SM) type, only the **main*	- mode is supported.</ph>
	//
	// - **IkeConfig.IkeEncAlg**: The encryption algorithm for Phase 1 negotiations.
	//
	//    <props="intl"><ph>Valid values: **aes**, **aes192**, **aes256**, **des**, and **3des**. Default value: **aes**. </ph>
	//
	//    <props="china">If the VPN gateway instance is of the standard type, valid values are **aes**, **aes192**, **aes256**, **des**, and **3des**. Default value: **aes**.
	//
	//    <props="china"><ph>If the VPN gateway instance is of the ShangMi (SM) type, the value is **sm4*	- (default).</ph>
	//
	// - **IkeConfig.IkeAuthAlg**: The authentication algorithm for Phase 1 negotiations.
	//
	//    <props="intl"><ph>Valid values: **md5**, **sha1**, **sha256**, **sha384**, and **sha512**. Default value: **md5**.</ph>
	//
	//    <props="china"><ph>If the VPN gateway instance is of the standard type, valid values are **md5**, **sha1**, **sha256**, **sha384**, and **sha512**. Default value: **md5**.</ph>
	//
	//    <props="china"><ph>If the VPN gateway instance is of the ShangMi (SM) type, the value is **sm3*	- (default).</ph>
	//
	// - **IkeConfig.IkePfs**: The Diffie-Hellman key exchange algorithm used in Phase 1 negotiations. Valid values: **group1**, **group2**, **group5**, and **group14**. Default value: **group2**.
	//
	// - **IkeConfig.IkeLifetime**: The SA lifecycle negotiated in Phase 1. Unit: seconds. Valid values: **0*	- to **86400**. Default value: **86400**.
	//
	// - **IkeConfig.LocalId**: The identifier of the VPN gateway. The identifier can be up to 100 characters in length and cannot contain spaces. The default value is the IP address of the VPN gateway.
	//
	// - **IkeConfig.RemoteId**: The identifier of the customer gateway. The identifier can be up to 100 characters in length and cannot contain spaces. The default value is the IP address of the customer gateway.
	//
	// example:
	//
	// {"Psk":"1234****","IkeVersion":"ikev1","IkeMode":"main","IkeEncAlg":"aes","IkeAuthAlg":"sha1","IkePfs":"group2","IkeLifetime":86400,"LocalId":"47.XX.XX.1","RemoteId":"47.XX.XX.2"}
	IkeConfig *string `json:"IkeConfig,omitempty" xml:"IkeConfig,omitempty"`
	// This parameter is supported when you create an IPsec-VPN connection in single-tunnel mode.
	//
	// The configurations for Phase 2 negotiations:
	//
	// - **IpsecConfig.IpsecEncAlg**: The encryption algorithm for Phase 2 negotiations.
	//
	//    <props="intl"><ph>Valid values: **aes**, **aes192**, **aes256**, **des**, and **3des**. Default value: **aes**. </ph>
	//
	//    <props="china"><ph>If the VPN gateway instance is of the standard type, valid values are **aes**, **aes192**, **aes256**, **des**, and **3des**. Default value: **aes**.</ph>
	//
	//    <props="china"><ph>If the VPN gateway instance is of the ShangMi (SM) type, the value is **sm4*	- (default).</ph>
	//
	// - **IpsecConfig.IpsecAuthAlg**: The authentication algorithm for Phase 2 negotiations.
	//
	//    <props="intl"><ph>Valid values: **md5**, **sha1**, **sha256**, **sha384**, and **sha512**. Default value: **md5**.</ph>
	//
	//    <props="china"><ph>If the VPN gateway instance is of the standard type, valid values are **md5**, **sha1**, **sha256**, **sha384**, and **sha512**. Default value: **md5**.</ph>
	//
	//    <props="china"><ph>If the VPN gateway instance is of the ShangMi (SM) type, the value is **sm3*	- (default).</ph>
	//
	// - **IpsecConfig.IpsecPfs**: The Diffie-Hellman key exchange algorithm used in Phase 2 negotiations. Valid values: **disabled**, **group1**, **group2**, **group5**, and **group14**. Default value: **group2**.
	//
	// - **IpsecConfig.IpsecLifetime**: The SA lifecycle negotiated in Phase 2. Unit: seconds. Valid values: **0*	- to **86400**. Default value: **86400**.
	//
	// example:
	//
	// {"IpsecEncAlg":"aes","IpsecAuthAlg":"sha1","IpsecPfs":"group2","IpsecLifetime":86400}
	IpsecConfig *string `json:"IpsecConfig,omitempty" xml:"IpsecConfig,omitempty"`
	// The CIDR block on the VPC side that needs to communicate with the data center. This parameter is used for Phase 2 negotiations.
	//
	// Separate multiple CIDR blocks with commas (,). Example: 192.168.1.0/24,192.168.2.0/24.
	//
	// Description of the routing mode of the IPsec-VPN connection:
	//
	// - If both **LocalSubnet*	- and **RemoteSubnet*	- are set to 0.0.0.0/0, the destination routing mode is used.
	//
	// - If both **LocalSubnet*	- and **RemoteSubnet*	- are set to specific CIDR blocks, the interesting traffic mode is used.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.10.1.0/24,10.10.2.0/24
	LocalSubnet *string `json:"LocalSubnet,omitempty" xml:"LocalSubnet,omitempty"`
	// The name of the IPsec-VPN connection.
	//
	// The name must be 1 to 100 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// IPsec
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the IPsec-VPN connection. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is supported when you create an IPsec-VPN connection in single-tunnel mode.
	//
	// If the current VPN gateway instance is a ShangMi (SM) VPN gateway, you must configure the CA certificate of the peer.
	//
	// - For a ShangMi (SM) VPN gateway, this parameter is required when you create an IPsec-VPN connection.
	//
	// - For a standard VPN gateway, this parameter must be left empty.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- MIIB7zCCAZW***	- -----END CERTIFICATE-----
	RemoteCaCertificate *string `json:"RemoteCaCertificate,omitempty" xml:"RemoteCaCertificate,omitempty"`
	// The CIDR block on the data center side that needs to communicate with the VPC. This parameter is used for Phase 2 negotiations.
	//
	// Separate multiple CIDR blocks with commas (,). Example: 192.168.3.0/24,192.168.4.0/24.
	//
	// Description of the routing mode of the IPsec-VPN connection:
	//
	// - If both **LocalSubnet*	- and **RemoteSubnet*	- are set to 0.0.0.0/0, the destination routing mode is used.
	//
	// - If both **LocalSubnet*	- and **RemoteSubnet*	- are set to specific CIDR blocks, the interesting traffic mode is used.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.10.3.0/24,10.10.4.0/24
	RemoteSubnet         *string `json:"RemoteSubnet,omitempty" xml:"RemoteSubnet,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The list of tags to be added to the IPsec-VPN connection.
	//
	// You can add up to 20 tags to an IPsec-VPN connection at a time.
	Tags []*CreateVpnConnectionRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The tunnel configurations.
	//
	// - The parameters in the **TunnelOptionsSpecification*	- array are supported when you create an IPsec-VPN connection in dual-tunnel mode.
	//
	// - When you create an IPsec-VPN connection in dual-tunnel mode, you must add configurations for both the primary and secondary tunnels. Only one primary tunnel and one secondary tunnel can be added to an IPsec-VPN connection.
	//
	// if can be null:
	// true
	TunnelOptionsSpecification []*CreateVpnConnectionRequestTunnelOptionsSpecification `json:"TunnelOptionsSpecification,omitempty" xml:"TunnelOptionsSpecification,omitempty" type:"Repeated"`
	// The instance ID of the VPN gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpn-bp1q8bgx4xnkm****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
}

func (s CreateVpnConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVpnConnectionRequest) GoString() string {
	return s.String()
}

func (s *CreateVpnConnectionRequest) GetAutoConfigRoute() *bool {
	return s.AutoConfigRoute
}

func (s *CreateVpnConnectionRequest) GetBgpConfig() *string {
	return s.BgpConfig
}

func (s *CreateVpnConnectionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateVpnConnectionRequest) GetCustomerGatewayId() *string {
	return s.CustomerGatewayId
}

func (s *CreateVpnConnectionRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateVpnConnectionRequest) GetEffectImmediately() *bool {
	return s.EffectImmediately
}

func (s *CreateVpnConnectionRequest) GetEnableDpd() *bool {
	return s.EnableDpd
}

func (s *CreateVpnConnectionRequest) GetEnableNatTraversal() *bool {
	return s.EnableNatTraversal
}

func (s *CreateVpnConnectionRequest) GetEnableTunnelsBgp() *bool {
	return s.EnableTunnelsBgp
}

func (s *CreateVpnConnectionRequest) GetHealthCheckConfig() *string {
	return s.HealthCheckConfig
}

func (s *CreateVpnConnectionRequest) GetIkeConfig() *string {
	return s.IkeConfig
}

func (s *CreateVpnConnectionRequest) GetIpsecConfig() *string {
	return s.IpsecConfig
}

func (s *CreateVpnConnectionRequest) GetLocalSubnet() *string {
	return s.LocalSubnet
}

func (s *CreateVpnConnectionRequest) GetName() *string {
	return s.Name
}

func (s *CreateVpnConnectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateVpnConnectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateVpnConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateVpnConnectionRequest) GetRemoteCaCertificate() *string {
	return s.RemoteCaCertificate
}

func (s *CreateVpnConnectionRequest) GetRemoteSubnet() *string {
	return s.RemoteSubnet
}

func (s *CreateVpnConnectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateVpnConnectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateVpnConnectionRequest) GetTags() []*CreateVpnConnectionRequestTags {
	return s.Tags
}

func (s *CreateVpnConnectionRequest) GetTunnelOptionsSpecification() []*CreateVpnConnectionRequestTunnelOptionsSpecification {
	return s.TunnelOptionsSpecification
}

func (s *CreateVpnConnectionRequest) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *CreateVpnConnectionRequest) SetAutoConfigRoute(v bool) *CreateVpnConnectionRequest {
	s.AutoConfigRoute = &v
	return s
}

func (s *CreateVpnConnectionRequest) SetBgpConfig(v string) *CreateVpnConnectionRequest {
	s.BgpConfig = &v
	return s
}

func (s *CreateVpnConnectionRequest) SetClientToken(v string) *CreateVpnConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateVpnConnectionRequest) SetCustomerGatewayId(v string) *CreateVpnConnectionRequest {
	s.CustomerGatewayId = &v
	return s
}

func (s *CreateVpnConnectionRequest) SetDryRun(v bool) *CreateVpnConnectionRequest {
	s.DryRun = &v
	return s
}

func (s *CreateVpnConnectionRequest) SetEffectImmediately(v bool) *CreateVpnConnectionRequest {
	s.EffectImmediately = &v
	return s
}

func (s *CreateVpnConnectionRequest) SetEnableDpd(v bool) *CreateVpnConnectionRequest {
	s.EnableDpd = &v
	return s
}

func (s *CreateVpnConnectionRequest) SetEnableNatTraversal(v bool) *CreateVpnConnectionRequest {
	s.EnableNatTraversal = &v
	return s
}

func (s *CreateVpnConnectionRequest) SetEnableTunnelsBgp(v bool) *CreateVpnConnectionRequest {
	s.EnableTunnelsBgp = &v
	return s
}

func (s *CreateVpnConnectionRequest) SetHealthCheckConfig(v string) *CreateVpnConnectionRequest {
	s.HealthCheckConfig = &v
	return s
}

func (s *CreateVpnConnectionRequest) SetIkeConfig(v string) *CreateVpnConnectionRequest {
	s.IkeConfig = &v
	return s
}

func (s *CreateVpnConnectionRequest) SetIpsecConfig(v string) *CreateVpnConnectionRequest {
	s.IpsecConfig = &v
	return s
}

func (s *CreateVpnConnectionRequest) SetLocalSubnet(v string) *CreateVpnConnectionRequest {
	s.LocalSubnet = &v
	return s
}

func (s *CreateVpnConnectionRequest) SetName(v string) *CreateVpnConnectionRequest {
	s.Name = &v
	return s
}

func (s *CreateVpnConnectionRequest) SetOwnerAccount(v string) *CreateVpnConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateVpnConnectionRequest) SetOwnerId(v int64) *CreateVpnConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateVpnConnectionRequest) SetRegionId(v string) *CreateVpnConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVpnConnectionRequest) SetRemoteCaCertificate(v string) *CreateVpnConnectionRequest {
	s.RemoteCaCertificate = &v
	return s
}

func (s *CreateVpnConnectionRequest) SetRemoteSubnet(v string) *CreateVpnConnectionRequest {
	s.RemoteSubnet = &v
	return s
}

func (s *CreateVpnConnectionRequest) SetResourceOwnerAccount(v string) *CreateVpnConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateVpnConnectionRequest) SetResourceOwnerId(v int64) *CreateVpnConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateVpnConnectionRequest) SetTags(v []*CreateVpnConnectionRequestTags) *CreateVpnConnectionRequest {
	s.Tags = v
	return s
}

func (s *CreateVpnConnectionRequest) SetTunnelOptionsSpecification(v []*CreateVpnConnectionRequestTunnelOptionsSpecification) *CreateVpnConnectionRequest {
	s.TunnelOptionsSpecification = v
	return s
}

func (s *CreateVpnConnectionRequest) SetVpnGatewayId(v string) *CreateVpnConnectionRequest {
	s.VpnGatewayId = &v
	return s
}

func (s *CreateVpnConnectionRequest) Validate() error {
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

type CreateVpnConnectionRequestTags struct {
	// The tag key. This parameter cannot be an empty string once specified.
	//
	// The tag key can be up to 64 characters in length, and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// You can specify up to 20 tag keys at a time.
	//
	// example:
	//
	// TagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// The tag value can be up to 128 characters in length and can be an empty string. It cannot start with aliyun or acs:, and cannot contain http:// or https://.
	//
	// Each tag key corresponds to one tag value. You can specify up to 20 tag values at a time.
	//
	// example:
	//
	// TagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateVpnConnectionRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateVpnConnectionRequestTags) GoString() string {
	return s.String()
}

func (s *CreateVpnConnectionRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateVpnConnectionRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateVpnConnectionRequestTags) SetKey(v string) *CreateVpnConnectionRequestTags {
	s.Key = &v
	return s
}

func (s *CreateVpnConnectionRequestTags) SetValue(v string) *CreateVpnConnectionRequestTags {
	s.Value = &v
	return s
}

func (s *CreateVpnConnectionRequestTags) Validate() error {
	return dara.Validate(s)
}

type CreateVpnConnectionRequestTunnelOptionsSpecification struct {
	// The ID of the customer gateway associated with the tunnel.
	//
	// >- This parameter is required when you create an IPsec-VPN connection in dual-tunnel mode.
	//
	// >- The parameters in the **TunnelOptionsSpecification*	- array are supported when you create an IPsec-VPN connection in dual-tunnel mode.
	//
	// >- When you create an IPsec-VPN connection in dual-tunnel mode, you must add configurations for both the primary and secondary tunnels. Only one primary tunnel and one secondary tunnel can be added to an IPsec-VPN connection.
	//
	// example:
	//
	// cgw-p0wy363lucf1uyae8****
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" xml:"CustomerGatewayId,omitempty"`
	// Specifies whether to enable the dead peer detection (DPD) feature for the tunnel. Valid values:
	//
	// - **true*	- (default): Enables the DPD feature. The IPsec initiator sends DPD packets to check whether the peer device is alive. If no correct response is received within the specified time, the peer is considered disconnected. IPsec deletes the ISAKMP SA and the corresponding IPsec SA, and the secure tunnel is also deleted.
	//
	// - **false**: Disables the DPD feature. The IPsec initiator does not send DPD probe packets.
	//
	// example:
	//
	// true
	EnableDpd *bool `json:"EnableDpd,omitempty" xml:"EnableDpd,omitempty"`
	// Specifies whether to enable the NAT traversal feature for the tunnel. Valid values:
	//
	// - **true*	- (default): Enables the NAT traversal feature. After this feature is enabled, the IKE negotiation process skips the verification of UDP port numbers and discovers NAT gateway devices in the tunnel.
	//
	// - **false**: Disables the NAT traversal feature.
	//
	// example:
	//
	// true
	EnableNatTraversal *bool `json:"EnableNatTraversal,omitempty" xml:"EnableNatTraversal,omitempty"`
	// If the current VPN gateway instance is a ShangMi (SM) VPN gateway, you must configure the CA certificate of the peer for the tunnel.
	//
	// - For a ShangMi (SM) VPN gateway, this parameter is required.
	//
	// - For a standard VPN gateway, this parameter must be left empty.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- MIIB7zCCAZW***	- -----END CERTIFICATE-----
	RemoteCaCertificate *string `json:"RemoteCaCertificate,omitempty" xml:"RemoteCaCertificate,omitempty"`
	// The role of the tunnel. Valid values:
	//
	// - **master**: The current tunnel is the primary tunnel.
	//
	// - **slave**: The current tunnel is the secondary tunnel.
	//
	// example:
	//
	// master
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The BGP configurations for the tunnel.
	TunnelBgpConfig *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelBgpConfig `json:"TunnelBgpConfig,omitempty" xml:"TunnelBgpConfig,omitempty" type:"Struct"`
	// The configurations for Phase 1 negotiations.
	TunnelIkeConfig *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIkeConfig `json:"TunnelIkeConfig,omitempty" xml:"TunnelIkeConfig,omitempty" type:"Struct"`
	// The configurations for Phase 2 negotiations.
	TunnelIpsecConfig *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIpsecConfig `json:"TunnelIpsecConfig,omitempty" xml:"TunnelIpsecConfig,omitempty" type:"Struct"`
}

func (s CreateVpnConnectionRequestTunnelOptionsSpecification) String() string {
	return dara.Prettify(s)
}

func (s CreateVpnConnectionRequestTunnelOptionsSpecification) GoString() string {
	return s.String()
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecification) GetCustomerGatewayId() *string {
	return s.CustomerGatewayId
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecification) GetEnableDpd() *bool {
	return s.EnableDpd
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecification) GetEnableNatTraversal() *bool {
	return s.EnableNatTraversal
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecification) GetRemoteCaCertificate() *string {
	return s.RemoteCaCertificate
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecification) GetRole() *string {
	return s.Role
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecification) GetTunnelBgpConfig() *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelBgpConfig {
	return s.TunnelBgpConfig
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecification) GetTunnelIkeConfig() *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIkeConfig {
	return s.TunnelIkeConfig
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecification) GetTunnelIpsecConfig() *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIpsecConfig {
	return s.TunnelIpsecConfig
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecification) SetCustomerGatewayId(v string) *CreateVpnConnectionRequestTunnelOptionsSpecification {
	s.CustomerGatewayId = &v
	return s
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecification) SetEnableDpd(v bool) *CreateVpnConnectionRequestTunnelOptionsSpecification {
	s.EnableDpd = &v
	return s
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecification) SetEnableNatTraversal(v bool) *CreateVpnConnectionRequestTunnelOptionsSpecification {
	s.EnableNatTraversal = &v
	return s
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecification) SetRemoteCaCertificate(v string) *CreateVpnConnectionRequestTunnelOptionsSpecification {
	s.RemoteCaCertificate = &v
	return s
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecification) SetRole(v string) *CreateVpnConnectionRequestTunnelOptionsSpecification {
	s.Role = &v
	return s
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecification) SetTunnelBgpConfig(v *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelBgpConfig) *CreateVpnConnectionRequestTunnelOptionsSpecification {
	s.TunnelBgpConfig = v
	return s
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecification) SetTunnelIkeConfig(v *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIkeConfig) *CreateVpnConnectionRequestTunnelOptionsSpecification {
	s.TunnelIkeConfig = v
	return s
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecification) SetTunnelIpsecConfig(v *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIpsecConfig) *CreateVpnConnectionRequestTunnelOptionsSpecification {
	s.TunnelIpsecConfig = v
	return s
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecification) Validate() error {
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

type CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelBgpConfig struct {
	// The autonomous system number (ASN) on the local side (Alibaba Cloud side) of the tunnel. Valid values: **1*	- to **4294967295**. Default value: **45104**.
	//
	// > - This parameter is required when you enable the BGP feature for the IPsec-VPN connection (that is, when you set the **EnableTunnelsBgp*	- parameter to **true**).
	//
	// >- Before adding BGP configurations, understand the working mechanism and limits of the BGP dynamic routing feature. For more information, refer to [Configure BGP dynamic routing](https://help.aliyun.com/document_detail/2638220.html).
	//
	// >- Use a private ASN to establish a BGP connection with Alibaba Cloud. Refer to relevant documentation for the range of private ASNs.
	//
	// example:
	//
	// 65530
	LocalAsn *int64 `json:"LocalAsn,omitempty" xml:"LocalAsn,omitempty"`
	// The BGP IP address on the local side (Alibaba Cloud side) of the tunnel. This IP address must be within the BGP CIDR block.
	//
	// example:
	//
	// 169.254.10.1
	LocalBgpIp *string `json:"LocalBgpIp,omitempty" xml:"LocalBgpIp,omitempty"`
	// The BGP CIDR block of the tunnel. The CIDR block must be within the 169.254.0.0/16 range and have a mask length of 30, and cannot be 169.254.0.0/30, 169.254.1.0/30, 169.254.2.0/30, 169.254.3.0/30, 169.254.4.0/30, 169.254.5.0/30, 169.254.6.0/30, or 169.254.169.252/30.
	//
	// > The BGP CIDR block of each tunnel under the same VPN gateway instance must be unique.
	//
	// example:
	//
	// 169.254.10.0/30
	TunnelCidr *string `json:"TunnelCidr,omitempty" xml:"TunnelCidr,omitempty"`
}

func (s CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelBgpConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelBgpConfig) GoString() string {
	return s.String()
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelBgpConfig) GetLocalAsn() *int64 {
	return s.LocalAsn
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelBgpConfig) GetLocalBgpIp() *string {
	return s.LocalBgpIp
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelBgpConfig) GetTunnelCidr() *string {
	return s.TunnelCidr
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelBgpConfig) SetLocalAsn(v int64) *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelBgpConfig {
	s.LocalAsn = &v
	return s
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelBgpConfig) SetLocalBgpIp(v string) *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelBgpConfig {
	s.LocalBgpIp = &v
	return s
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelBgpConfig) SetTunnelCidr(v string) *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelBgpConfig {
	s.TunnelCidr = &v
	return s
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelBgpConfig) Validate() error {
	return dara.Validate(s)
}

type CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIkeConfig struct {
	// The authentication algorithm for Phase 1 negotiations.
	//
	// <props="intl"><ph>Valid values: **md5**, **sha1**, **sha256**, **sha384**, and **sha512**. Default value: **md5**.</ph>
	//
	// <props="china"><ph>If the VPN gateway instance is of the standard type, valid values are **md5**, **sha1**, **sha256**, **sha384**, and **sha512**. Default value: **md5**.</ph>
	//
	// <props="china"><ph>If the VPN gateway instance is of the ShangMi (SM) type, the value is **sm3*	- (default).</ph>
	//
	// example:
	//
	// md5
	IkeAuthAlg *string `json:"IkeAuthAlg,omitempty" xml:"IkeAuthAlg,omitempty"`
	// The encryption algorithm for Phase 1 negotiations.
	//
	// <props="intl"><ph>Valid values: **aes**, **aes192**, **aes256**, **des**, and **3des**. Default value: **aes**. </ph>
	//
	// <props="china"><ph>If the VPN gateway instance is of the standard type, valid values are **aes**, **aes192**, **aes256**, **des**, and **3des**. Default value: **aes**.</ph>
	//
	// <props="china"><ph>If the VPN gateway instance is of the ShangMi (SM) type, the value is **sm4*	- (default).</ph>
	//
	// example:
	//
	// aes
	IkeEncAlg *string `json:"IkeEncAlg,omitempty" xml:"IkeEncAlg,omitempty"`
	// The SA lifecycle negotiated in Phase 1. Unit: seconds.
	//
	// Valid values: **0*	- to **86400**. Default value: **86400**.
	//
	// example:
	//
	// 86400
	IkeLifetime *int64 `json:"IkeLifetime,omitempty" xml:"IkeLifetime,omitempty"`
	// The negotiation mode of the IKE version. Valid values: **main*	- and **aggressive**. Default value: **main**.
	//
	// - **main**: The main mode. The negotiation process is highly secure.
	//
	// - **aggressive**: The aggressive mode. The negotiation process is fast and has a high success rate.
	//
	// <props="china"><ph>If the VPN gateway instance is of the ShangMi (SM) type, only the **main*	- mode is supported.</ph>
	//
	// example:
	//
	// main
	IkeMode *string `json:"IkeMode,omitempty" xml:"IkeMode,omitempty"`
	// The Diffie-Hellman key exchange algorithm used in Phase 1 negotiations. Default value: **group2**.
	//
	// Valid values: **group1**, **group2**, **group5**, and **group14**.
	//
	// example:
	//
	// group2
	IkePfs *string `json:"IkePfs,omitempty" xml:"IkePfs,omitempty"`
	// The version of the IKE protocol. Valid values: **ikev1*	- and **ikev2**. Default value: **ikev1**.
	//
	// Compared with IKEv1, IKEv2 simplifies the SA negotiation process and provides better support for scenarios with multiple CIDR blocks.
	//
	//
	//
	// <props="china"><ph>If the VPN gateway instance is of the ShangMi (SM) type, only **ikev1*	- is supported.</ph>
	//
	// example:
	//
	// ikev1
	IkeVersion *string `json:"IkeVersion,omitempty" xml:"IkeVersion,omitempty"`
	// The identifier on the local side (Alibaba Cloud side) of the tunnel, which is used for Phase 1 negotiations. The identifier can be up to 100 characters in length and cannot contain spaces. The default value is the IP address of the tunnel.
	//
	// **LocalId*	- supports the FQDN format. If you use the FQDN format, we recommend that you set the negotiation mode to **aggressive**.
	//
	// example:
	//
	// 47.21.XX.XX
	LocalId *string `json:"LocalId,omitempty" xml:"LocalId,omitempty"`
	// The pre-shared key used for authentication between the tunnel and the peer.
	//
	// - The key must be 1 to 100 characters in length, and can contain digits, letters, and the following special characters: ```~!\\`@#$%^&*()_-+={}[]|;:\\",.<>/?``` The key cannot contain spaces.
	//
	// - If you do not specify a pre-shared key, the system randomly generates a 16-character string as the pre-shared key. You can call the [DescribeVpnConnection](https://help.aliyun.com/document_detail/2526951.html) operation to query the automatically generated pre-shared key.
	//
	// > The pre-shared key of the tunnel must be the same as that of the peer. Otherwise, the system cannot establish the tunnel.
	//
	// example:
	//
	// 123456****
	Psk *string `json:"Psk,omitempty" xml:"Psk,omitempty"`
	// The identifier of the peer, which is used for Phase 1 negotiations. The identifier can be up to 100 characters in length and cannot contain spaces. The default value is the IP address of the customer gateway associated with the tunnel.
	//
	// **RemoteId*	- supports the FQDN format. If you use the FQDN format, we recommend that you set the negotiation mode to **aggressive**.
	//
	// example:
	//
	// 47.42.XX.XX
	RemoteId *string `json:"RemoteId,omitempty" xml:"RemoteId,omitempty"`
}

func (s CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIkeConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIkeConfig) GoString() string {
	return s.String()
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIkeConfig) GetIkeAuthAlg() *string {
	return s.IkeAuthAlg
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIkeConfig) GetIkeEncAlg() *string {
	return s.IkeEncAlg
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIkeConfig) GetIkeLifetime() *int64 {
	return s.IkeLifetime
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIkeConfig) GetIkeMode() *string {
	return s.IkeMode
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIkeConfig) GetIkePfs() *string {
	return s.IkePfs
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIkeConfig) GetIkeVersion() *string {
	return s.IkeVersion
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIkeConfig) GetLocalId() *string {
	return s.LocalId
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIkeConfig) GetPsk() *string {
	return s.Psk
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIkeConfig) GetRemoteId() *string {
	return s.RemoteId
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIkeConfig) SetIkeAuthAlg(v string) *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkeAuthAlg = &v
	return s
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIkeConfig) SetIkeEncAlg(v string) *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkeEncAlg = &v
	return s
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIkeConfig) SetIkeLifetime(v int64) *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkeLifetime = &v
	return s
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIkeConfig) SetIkeMode(v string) *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkeMode = &v
	return s
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIkeConfig) SetIkePfs(v string) *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkePfs = &v
	return s
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIkeConfig) SetIkeVersion(v string) *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkeVersion = &v
	return s
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIkeConfig) SetLocalId(v string) *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.LocalId = &v
	return s
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIkeConfig) SetPsk(v string) *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.Psk = &v
	return s
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIkeConfig) SetRemoteId(v string) *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.RemoteId = &v
	return s
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIkeConfig) Validate() error {
	return dara.Validate(s)
}

type CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIpsecConfig struct {
	// The authentication algorithm for Phase 2 negotiations.
	//
	// <props="intl"><ph>Valid values: **md5**, **sha1**, **sha256**, **sha384**, and **sha512**. Default value: **md5**.</ph>
	//
	// <props="china"><ph>If the VPN gateway instance is of the standard type, valid values are **md5**, **sha1**, **sha256**, **sha384**, and **sha512**. Default value: **md5**.</ph>
	//
	// <props="china"><ph>If the VPN gateway instance is of the ShangMi (SM) type, the value is **sm3*	- (default).</ph>
	//
	// example:
	//
	// md5
	IpsecAuthAlg *string `json:"IpsecAuthAlg,omitempty" xml:"IpsecAuthAlg,omitempty"`
	// The encryption algorithm for Phase 2 negotiations.
	//
	// <props="intl"><ph>Valid values: **aes**, **aes192**, **aes256**, **des**, and **3des**. Default value: **aes**. </ph>
	//
	// <props="china"><ph>If the VPN gateway instance is of the standard type, valid values are **aes**, **aes192**, **aes256**, **des**, and **3des**. Default value: **aes**.</ph>
	//
	// <props="china"><ph>If the VPN gateway instance is of the ShangMi (SM) type, the value is **sm4*	- (default).</ph>
	//
	// example:
	//
	// aes
	IpsecEncAlg *string `json:"IpsecEncAlg,omitempty" xml:"IpsecEncAlg,omitempty"`
	// The SA lifecycle negotiated in Phase 2. Unit: seconds.
	//
	// Valid values: **0*	- to **86400**. Default value: **86400**.
	//
	// example:
	//
	// 86400
	IpsecLifetime *int64 `json:"IpsecLifetime,omitempty" xml:"IpsecLifetime,omitempty"`
	// The Diffie-Hellman key exchange algorithm used in Phase 2 negotiations. Default value: **group2**.
	//
	// Valid values: **disabled**, **group1**, **group2**, **group5**, and **group14**.
	//
	// example:
	//
	// group2
	IpsecPfs *string `json:"IpsecPfs,omitempty" xml:"IpsecPfs,omitempty"`
}

func (s CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIpsecConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIpsecConfig) GoString() string {
	return s.String()
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIpsecConfig) GetIpsecAuthAlg() *string {
	return s.IpsecAuthAlg
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIpsecConfig) GetIpsecEncAlg() *string {
	return s.IpsecEncAlg
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIpsecConfig) GetIpsecLifetime() *int64 {
	return s.IpsecLifetime
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIpsecConfig) GetIpsecPfs() *string {
	return s.IpsecPfs
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIpsecConfig) SetIpsecAuthAlg(v string) *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIpsecConfig {
	s.IpsecAuthAlg = &v
	return s
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIpsecConfig) SetIpsecEncAlg(v string) *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIpsecConfig {
	s.IpsecEncAlg = &v
	return s
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIpsecConfig) SetIpsecLifetime(v int64) *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIpsecConfig {
	s.IpsecLifetime = &v
	return s
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIpsecConfig) SetIpsecPfs(v string) *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIpsecConfig {
	s.IpsecPfs = &v
	return s
}

func (s *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIpsecConfig) Validate() error {
	return dara.Validate(s)
}
