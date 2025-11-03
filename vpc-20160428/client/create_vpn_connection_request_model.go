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
	// 	- **true*	- (default)
	//
	// 	- **false**
	//
	// example:
	//
	// true
	AutoConfigRoute *bool `json:"AutoConfigRoute,omitempty" xml:"AutoConfigRoute,omitempty"`
	// This parameter is supported if you create an IPsec-VPN connection in single-tunnel mode.
	//
	// BGP configuration:
	//
	// 	- **BgpConfig.EnableBgp**: specifies whether to enable BGP. Valid values: **true*	- and **false*	- (default).
	//
	// 	- **BgpConfig.LocalAsn:*	- the autonomous system number (ASN) on the Alibaba Cloud side. Valid values: **1*	- to **4294967295**. Default value: **45104**.
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
	// > - Before you add BGP configurations, we recommend that you learn about how BGP works and the limits. For more information, see [Configure BGP dynamic routing](https://help.aliyun.com/document_detail/2638220.html).
	//
	// > - We recommend that you use a private ASN to establish BGP connections to Alibaba Cloud. Refer to the relevant documentation for the private ASN range.
	//
	// example:
	//
	// {"EnableBgp":"true","LocalAsn":"45104","TunnelCidr":"169.254.11.0/30","LocalBgpIp":"169.254.11.1"}
	BgpConfig *string `json:"BgpConfig,omitempty" xml:"BgpConfig,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// When you create an IPsec-VPN connection in single-tunnel mode, this parameter is required.
	//
	// The ID of the customer gateway.
	//
	// example:
	//
	// cgw-p0w2jemrcj5u61un8****
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" xml:"CustomerGatewayId,omitempty"`
	DryRun            *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to immediately start IPsec negotiations. Valid values:
	//
	// 	- **true**: immediately starts IPsec negotiations.
	//
	// 	- **false*	- (default): starts IPsec negotiations when inbound traffic is detected.
	//
	// example:
	//
	// false
	EffectImmediately *bool `json:"EffectImmediately,omitempty" xml:"EffectImmediately,omitempty"`
	// This parameter is available if you create an IPsec-VPN connection in single-tunnel mode.
	//
	// Specifies whether to enable the dead peer detection (DPD) feature. Valid values:
	//
	// 	- **true*	- (default) The initiator of the IPsec-VPN connection sends DPD packets to verify the existence and availability of the peer. If no feedback is received from the peer within a specified period of time, the connection fails. ISAKMP SAs and IPsec SAs are deleted. The IPsec tunnel is also deleted.
	//
	// 	- **false**
	//
	// example:
	//
	// true
	EnableDpd *bool `json:"EnableDpd,omitempty" xml:"EnableDpd,omitempty"`
	// This parameter is available if you create an IPsec-VPN connection in single-tunnel mode.
	//
	// Specifies whether to enable NAT traversal. Valid values:
	//
	// 	- **true*	- (default) After NAT traversal is enabled, the initiator does not check the UDP ports during IKE negotiations and can automatically discover NAT gateway devices along the VPN tunnel.
	//
	// 	- **false**
	//
	// example:
	//
	// true
	EnableNatTraversal *bool `json:"EnableNatTraversal,omitempty" xml:"EnableNatTraversal,omitempty"`
	// This parameter is available if you create an IPsec-VPN connection in dual-tunnel mode.
	//
	// Specifies whether to enable the BGP feature for the tunnel. Valid values: **true*	- and **false**. Default value: false.
	//
	// example:
	//
	// true
	EnableTunnelsBgp *bool `json:"EnableTunnelsBgp,omitempty" xml:"EnableTunnelsBgp,omitempty"`
	// This parameter is available if you create an IPsec-VPN connection in single-tunnel mode.
	//
	// The health check configuration:
	//
	// 	- **HealthCheckConfig.enable**: specifies whether to enable health checks. Valid values: **true*	- and **false**. Default value: false.
	//
	// 	- **HealthCheckConfig.dip**: the destination IP address configured for health checks.
	//
	// 	- **HealthCheckConfig.sip:*	- the source IP address that is used for health checks.
	//
	// 	- **HealthCheckConfig.interval**: the time interval of health check retries. Unit: seconds. Default value: **3**.
	//
	// 	- **HealthCheckConfig.retry**: the maximum number of health check retries. Default value: **3**.
	//
	// example:
	//
	// {"enable":"true","dip":"192.168.10.1","sip":"10.10.1.1","interval":"3","retry":"3"}
	HealthCheckConfig *string `json:"HealthCheckConfig,omitempty" xml:"HealthCheckConfig,omitempty"`
	// This parameter is supported if you create an IPsec-VPN connection in single-tunnel mode.
	//
	// The configurations of Phase 1 negotiations:
	//
	// 	- **IkeConfig.Psk**: the pre-shared key that is used for identity authentication between the VPN gateway and the on-premises data center.
	//
	//     	- The key cannot contain spaces. The key must be 1 to 100 characters in length, and can contain digits, letters, and the following special characters: ``~!\\`@#$%^&*()_-+={}[]|;:\\",.<>/?``
	//
	//     	- If you do not specify a pre-shared key, the system randomly generates a 16-bit string as the pre-shared key. You can call the [DescribeVpnConnection](https://help.aliyun.com/document_detail/2526951.html) operation to query the pre-shared key that is automatically generated by the system.
	//
	//         > The pre-shared key of the IPsec-VPN connection must be the same as the authentication key of the on-premises data center. Otherwise, connections between the on-premises data center and the VPN gateway cannot be established.
	//
	// 	- **IkeConfig.IkeVersion**: the version of the Internet Key Exchange (IKE) protocol. Valid values: **ikev1*	- and **ikev2**. Default value: **ikev1**.
	//
	//     Compared with IKEv1, IKEv2 simplifies the security association (SA) negotiation process and provides better support for scenarios with multiple CIDR blocks.
	//
	// 	- **IkeConfig.IkeMode**: the negotiation mode of IKE. Valid values: **main*	- and **aggressive**. Default value: **main**.
	//
	//     	- **main:*	- This mode offers higher security during negotiations.
	//
	//     	- **aggressive:*	- This mode is faster and has a higher success rate.
	//
	// 	- **IkeConfig.IkeEncAlg**: the encryption algorithm that is used in Phase 1 negotiations.
	//
	//     Valid values: **aes**, **aes192**, **aes256**, **des**, and **3des**. Default value: **aes**.
	//
	// 	- **IkeConfig.IkeAuthAlg**: the authentication algorithm that is used in Phase 1 negotiations.
	//
	//     Valid values: **md5**, **sha1**, **sha256**, **sha384**, and **sha512**. Default value: **md5**.
	//
	// 	- **IkeConfig.IkePfs**: the Diffie-Hellman key exchange algorithm that is used in Phase 1 negotiations. Valid values: **group1**, **group2**, **group5**, and **group14**. Default value: **group2**.
	//
	// 	- **IkeConfig.IkeLifetime**: the SA lifetime as a result of Phase 1 negotiations. Unit: seconds Valid values: **0*	- to **86400**. Default value: **86400**.
	//
	// 	- **IkeConfig.LocalId**: the identifier of the VPN gateway. It can be up to 100 characters in length and cannot contain space characters. The default value is the IP address of the VPN gateway.
	//
	// 	- **IkeConfig.RemoteId**: the identifier of the customer gateway. It can be up to 100 characters in length and cannot contain space characters. The default value is the IP address of the customer gateway.
	//
	// example:
	//
	// {"Psk":"1234****","IkeVersion":"ikev1","IkeMode":"main","IkeEncAlg":"aes","IkeAuthAlg":"sha1","IkePfs":"group2","IkeLifetime":86400,"LocalId":"47.XX.XX.1","RemoteId":"47.XX.XX.2"}
	IkeConfig *string `json:"IkeConfig,omitempty" xml:"IkeConfig,omitempty"`
	// This parameter is available if you create an IPsec-VPN connection in single-tunnel mode.
	//
	// The configurations of Phase 2 negotiations:
	//
	// 	- **IpsecConfig.IpsecEncAlg**: the encryption algorithm that is used in Phase 2 negotiations.
	//
	//     Valid values: **aes**, **aes192**, **aes256**, **des**, and **3des**. Default value: **aes**.
	//
	// 	- **IpsecConfig. IpsecAuthAlg**: the authentication algorithm that is used in Phase 2 negotiations.
	//
	//     Valid values: **md5**, **sha1**, **sha256**, **sha384**, and **sha512**. Default value: **md5**.
	//
	// 	- **IpsecConfig. IpsecPfs**: the DH key exchange algorithm that is used in Phase 2 negotiations. Valid values: **disabled**, **group1**, **group2**, **group5**, and **group14**. Default value: **group2**.
	//
	// 	- **IpsecConfig. IpsecLifetime**: the SA lifetime that is determined by Phase 2 negotiations. Unit: seconds. Valid values: **0*	- to **86400**. Default value: **86400**.
	//
	// example:
	//
	// {"IpsecEncAlg":"aes","IpsecAuthAlg":"sha1","IpsecPfs":"group2","IpsecLifetime":86400}
	IpsecConfig *string `json:"IpsecConfig,omitempty" xml:"IpsecConfig,omitempty"`
	// The CIDR block of the virtual private cloud (VPC) that needs to communicate with the on-premises database. The CIDR block is used in Phase 2 negotiations.
	//
	// Separate multiple CIDR blocks with commas (,). Example: 192.168.1.0/24,192.168.2.0/24.
	//
	// The following routing modes are supported:
	//
	// 	- If you set **LocalSubnet*	- and **RemoteSubnet*	- to 0.0.0.0/0, the routing mode of the IPsec-VPN connection is set to Destination Routing Mode.
	//
	// 	- If you set **LocalSubnet*	- and **RemoteSubnet*	- to specific CIDR blocks, the routing mode of the IPsec-VPN connection is set to Protected Data Flows.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.10.1.0/24,10.10.2.0/24
	LocalSubnet *string `json:"LocalSubnet,omitempty" xml:"LocalSubnet,omitempty"`
	// The name of the IPsec-VPN connection.
	//
	// The name must be 1 to 100 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// IPsec
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the IPsec-VPN connection is created. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is available if you create an IPsec-VPN connection in single-tunnel mode.
	//
	// The certificate authority (CA) certificate. If the VPN gateway is of the ShangMi (SM) type, you need to configure a CA certificate for the peer gateway device.
	//
	// 	- If an SM VPN gateway is used to create the IPsec-VPN connection, this parameter is required.
	//
	// 	- If a standard VPN gateway is used to create the IPsec-VPN connection, leave this parameter empty.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- MIIB7zCCAZW***	- -----END CERTIFICATE-----
	RemoteCaCertificate *string `json:"RemoteCaCertificate,omitempty" xml:"RemoteCaCertificate,omitempty"`
	// The CIDR block of the on-premises database that needs to communicate with the VPC. The CIDR block is used in Phase 2 negotiations.
	//
	// Separate multiple CIDR blocks with commas (,). Example: 192.168.3.0/24,192.168.4.0/24.
	//
	// The following routing modes are supported:
	//
	// 	- If you set **LocalSubnet*	- and **RemoteSubnet*	- to 0.0.0.0/0, the routing mode of the IPsec-VPN connection is set to Destination Routing Mode.
	//
	// 	- If you set **LocalSubnet*	- and **RemoteSubnet*	- to specific CIDR blocks, the routing mode of the IPsec-VPN connection is set to Protected Data Flows.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.10.3.0/24,10.10.4.0/24
	RemoteSubnet         *string `json:"RemoteSubnet,omitempty" xml:"RemoteSubnet,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag value.
	//
	// The tag value can be an empty string and cannot exceed 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// Each tag key corresponds to one tag value. You can specify up to 20 tag values in each call.
	Tags []*CreateVpnConnectionRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The tunnel configurations.
	//
	// 	- You can specify parameters in the **TunnelOptionsSpecification*	- array when you create an IPsec-VPN connection in dual tunnel mode.
	//
	// 	- When you create an IPsec-VPN connection in dual tunnel mode, you must add configurations of the active and standby tunnels for the IPsec-VPN connection. Each IPsec-VPN connection supports only one active tunnel and one standby tunnel.
	//
	// if can be null:
	// true
	TunnelOptionsSpecification []*CreateVpnConnectionRequestTunnelOptionsSpecification `json:"TunnelOptionsSpecification,omitempty" xml:"TunnelOptionsSpecification,omitempty" type:"Repeated"`
	// The ID of the VPN gateway.
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
	// The tag key. The tag key cannot be an empty string.
	//
	// It can be at most 64 characters in length, and cannot contain `http://` or `https://`. It cannot start with `aliyun` or `acs:`.
	//
	// You can specify at most 20 tag keys in each call.
	//
	// example:
	//
	// TagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// The tag value can be an empty string and cannot exceed 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// Each tag key corresponds to one tag value. You can specify at most 20 tag values in each call.
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
	// The ID of the customer gateway that is associated with the tunnel.
	//
	// > - This parameter is required when you create an IPsec-VPN connection in dual-tunnel mode.
	//
	// > - You can specify parameters in the **TunnelOptionsSpecification*	- array when you create an IPsec-VPN connection in dual tunnel mode.
	//
	// > - When you create an IPsec-VPN connection in dual tunnel mode, you must add configurations of the active and standby tunnels for the IPsec-VPN connection. Each IPsec-VPN connection supports only one active tunnel and one standby tunnel.
	//
	// example:
	//
	// cgw-p0wy363lucf1uyae8****
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" xml:"CustomerGatewayId,omitempty"`
	// Specifies whether to enable the Dead Peer Detection (DPD) feature for the tunnel. Valid values:
	//
	// 	- **true*	- (default): enables DPD. The initiator of the IPsec-VPN connection sends DPD packets to check the existence and availability of the peer. If no feedback is received from the peer within the specified period of time, the connection fails. In this case, ISAKMP SA and IPsec SA are deleted. The security tunnel is also deleted.
	//
	// 	- **false**: disables DPD. The initiator of the IPsec-VPN connection does not send DPD packets.
	//
	// example:
	//
	// true
	EnableDpd *bool `json:"EnableDpd,omitempty" xml:"EnableDpd,omitempty"`
	// Specifies whether to enable NAT traversal for the tunnel. Valid values:
	//
	// 	- **true*	- (default): enables NAT traversal. After NAT traversal is enabled, the initiator does not check the UDP ports during IKE negotiations and can automatically discover NAT gateway devices along the IPsec-VPN tunnel.
	//
	// 	- **false**: disables NAT traversal.
	//
	// example:
	//
	// true
	EnableNatTraversal *bool `json:"EnableNatTraversal,omitempty" xml:"EnableNatTraversal,omitempty"`
	// If the VPN gateway uses an SM certificate, you need to configure the CA certificate used by the IPsec peer.
	//
	// 	- If the VPN gateway uses an SM certificate, this parameter is required.
	//
	// 	- If the VPN gateway does not use an SM certificate, leave this parameter empty.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- MIIB7zCCAZW***	- -----END CERTIFICATE-----
	RemoteCaCertificate *string `json:"RemoteCaCertificate,omitempty" xml:"RemoteCaCertificate,omitempty"`
	// The role of the tunnel. Valid values:
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
	TunnelBgpConfig *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelBgpConfig `json:"TunnelBgpConfig,omitempty" xml:"TunnelBgpConfig,omitempty" type:"Struct"`
	// The configurations of Phase 1 negotiations.
	TunnelIkeConfig *CreateVpnConnectionRequestTunnelOptionsSpecificationTunnelIkeConfig `json:"TunnelIkeConfig,omitempty" xml:"TunnelIkeConfig,omitempty" type:"Struct"`
	// The configurations of Phase 2 negotiations.
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
	// The autonomous system number (ASN) of the tunnel on the Alibaba Cloud side. Valid values: **1*	- to **4294967295**. Default value: **45104**.
	//
	// > - If you set **EnableTunnelsBgp*	- to **true**, you must set this parameter.
	//
	// > - Before you add BGP configurations, we recommend that you learn about how BGP dynamic routing works and the limits. For more information, see [Configure BGP dynamic routing](https://help.aliyun.com/document_detail/2638220.html).
	//
	// > - We recommend that you use a private ASN to establish BGP connections to Alibaba Cloud. For information about the range of private ASNs, see the relevant documentation.
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
	// The BGP CIDR block of the tunnel. The CIDR block must fall within 169.254.0.0/16 and the mask of the CIDR block must be 30 bits in length. The CIDR block cannot be 169.254.0.0/30, 169.254.1.0/30, 169.254.2.0/30, 169.254.3.0/30, 169.254.4.0/30, 169.254.5.0/30, 169.254.6.0/30, or 169.254.169.252/30.
	//
	// >  The BGP CIDR block of each tunnel must be unique on a VPN gateway.
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
	// The authentication algorithm that is used in Phase 1 negotiations.
	//
	// Valid values: **md5**, **sha1**, **sha256**, **sha384**, and **sha512**. Default value: **md5**.
	//
	// example:
	//
	// md5
	IkeAuthAlg *string `json:"IkeAuthAlg,omitempty" xml:"IkeAuthAlg,omitempty"`
	// The encryption algorithm that is used in Phase 1 negotiations.
	//
	// Valid values: **aes**, **aes192**, **aes256**, **des**, and **3des**. Default value: **aes**.
	//
	// example:
	//
	// aes
	IkeEncAlg *string `json:"IkeEncAlg,omitempty" xml:"IkeEncAlg,omitempty"`
	// The SA lifetime as a result of Phase 1 negotiations. Unit: seconds
	//
	// Valid values: **0*	- to **86400**. Default value: **86400**.
	//
	// example:
	//
	// 86400
	IkeLifetime *int64 `json:"IkeLifetime,omitempty" xml:"IkeLifetime,omitempty"`
	// The negotiation mode of IKE. Valid values: **main*	- and **aggressive**. Default value: **main**.
	//
	// 	- **main:*	- This mode offers higher security during negotiations.
	//
	// 	- **aggressive:*	- This mode is faster and has a higher success rate.
	//
	// example:
	//
	// main
	IkeMode *string `json:"IkeMode,omitempty" xml:"IkeMode,omitempty"`
	// The Diffie-Hellman key exchange algorithm that is used in Phase 1 negotiations. Default value: **group2**.\\
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
	// example:
	//
	// ikev1
	IkeVersion *string `json:"IkeVersion,omitempty" xml:"IkeVersion,omitempty"`
	// The identifier of the tunnel on the Alibaba Cloud side, which is used in Phase 1 negotiations. The identifier cannot exceed 100 characters in length and cannot contain space characters. The default value is the IP address of the tunnel.
	//
	// **LocalId*	- supports fully qualified domain names (FQDNs). If you use an FQDN, we recommend that you set the negotiation mode to **aggressive**.
	//
	// example:
	//
	// 47.21.XX.XX
	LocalId *string `json:"LocalId,omitempty" xml:"LocalId,omitempty"`
	// The pre-shared key that is used for identity authentication between the tunnel and the tunnel peer.
	//
	// 	- The key cannot contain spaces. The key must be 1 to 100 characters in length, and can contain digits, letters, and the following special characters: ``~!\\`@#$%^&*()_-+={}[]|;:\\",.<>/?``
	//
	// 	- If you do not specify a pre-shared key, the system randomly generates a 16-bit string as the pre-shared key. You can call the [DescribeVpnConnection](https://help.aliyun.com/document_detail/2526951.html) operation to query the pre-shared key that is automatically generated by the system.
	//
	// >  The tunnel and the tunnel peer must use the same pre-shared key. Otherwise, the tunnel cannot be established.
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
	// The authentication algorithm that is used in Phase 2 negotiations.
	//
	// Valid values: **md5**, **sha1**, **sha256**, **sha384**, and **sha512**. Default value: **md5**.
	//
	// example:
	//
	// md5
	IpsecAuthAlg *string `json:"IpsecAuthAlg,omitempty" xml:"IpsecAuthAlg,omitempty"`
	// The encryption algorithm that is used in Phase 2 negotiations.
	//
	// Valid values: **aes**, **aes192**, **aes256**, **des**, and **3des**. Default value: **aes**.
	//
	// example:
	//
	// aes
	IpsecEncAlg *string `json:"IpsecEncAlg,omitempty" xml:"IpsecEncAlg,omitempty"`
	// The SA lifetime as a result of Phase 2 negotiations. Unit: seconds
	//
	// Valid values: **0*	- to **86400**. Default value: **86400**.
	//
	// example:
	//
	// 86400
	IpsecLifetime *int64 `json:"IpsecLifetime,omitempty" xml:"IpsecLifetime,omitempty"`
	// The Diffie-Hellman key exchange algorithm that is used in Phase 2 negotiations. Default value: **group2**.
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
