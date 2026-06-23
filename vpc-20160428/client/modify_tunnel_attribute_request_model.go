// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTunnelAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyTunnelAttributeRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *ModifyTunnelAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyTunnelAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyTunnelAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyTunnelAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyTunnelAttributeRequest
	GetResourceOwnerId() *int64
	SetTunnelId(v string) *ModifyTunnelAttributeRequest
	GetTunnelId() *string
	SetTunnelOptionsSpecification(v *ModifyTunnelAttributeRequestTunnelOptionsSpecification) *ModifyTunnelAttributeRequest
	GetTunnelOptionsSpecification() *ModifyTunnelAttributeRequestTunnelOptionsSpecification
	SetVpnConnectionId(v string) *ModifyTunnelAttributeRequest
	GetVpnConnectionId() *string
}

type ModifyTunnelAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- as the **ClientToken**. The **RequestId*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-0016e04115b
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the IPsec-VPN connection.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tunnel ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// tun-gbyz2e070xzo93****
	TunnelId *string `json:"TunnelId,omitempty" xml:"TunnelId,omitempty"`
	// The tunnel configuration to modify.
	TunnelOptionsSpecification *ModifyTunnelAttributeRequestTunnelOptionsSpecification `json:"TunnelOptionsSpecification,omitempty" xml:"TunnelOptionsSpecification,omitempty" type:"Struct"`
	// The ID of the IPsec-VPN connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// vco-gw69vm1i71y354****
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" xml:"VpnConnectionId,omitempty"`
}

func (s ModifyTunnelAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTunnelAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyTunnelAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyTunnelAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyTunnelAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyTunnelAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyTunnelAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyTunnelAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyTunnelAttributeRequest) GetTunnelId() *string {
	return s.TunnelId
}

func (s *ModifyTunnelAttributeRequest) GetTunnelOptionsSpecification() *ModifyTunnelAttributeRequestTunnelOptionsSpecification {
	return s.TunnelOptionsSpecification
}

func (s *ModifyTunnelAttributeRequest) GetVpnConnectionId() *string {
	return s.VpnConnectionId
}

func (s *ModifyTunnelAttributeRequest) SetClientToken(v string) *ModifyTunnelAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyTunnelAttributeRequest) SetOwnerAccount(v string) *ModifyTunnelAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyTunnelAttributeRequest) SetOwnerId(v int64) *ModifyTunnelAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyTunnelAttributeRequest) SetRegionId(v string) *ModifyTunnelAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyTunnelAttributeRequest) SetResourceOwnerAccount(v string) *ModifyTunnelAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyTunnelAttributeRequest) SetResourceOwnerId(v int64) *ModifyTunnelAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyTunnelAttributeRequest) SetTunnelId(v string) *ModifyTunnelAttributeRequest {
	s.TunnelId = &v
	return s
}

func (s *ModifyTunnelAttributeRequest) SetTunnelOptionsSpecification(v *ModifyTunnelAttributeRequestTunnelOptionsSpecification) *ModifyTunnelAttributeRequest {
	s.TunnelOptionsSpecification = v
	return s
}

func (s *ModifyTunnelAttributeRequest) SetVpnConnectionId(v string) *ModifyTunnelAttributeRequest {
	s.VpnConnectionId = &v
	return s
}

func (s *ModifyTunnelAttributeRequest) Validate() error {
	if s.TunnelOptionsSpecification != nil {
		if err := s.TunnelOptionsSpecification.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyTunnelAttributeRequestTunnelOptionsSpecification struct {
	// The instance ID of the customer gateway associated with the tunnel.
	//
	// example:
	//
	// cgw-1nmwbpgrp7ssqm1yn****
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" xml:"CustomerGatewayId,omitempty"`
	// Specifies whether to enable the Dead Peer Detection (DPD) feature. Valid values:
	//
	// - **true**: Enabled. The initiator of the IPsec-VPN connection sends DPD packets to check whether the peer is alive. If no correct response is received within the specified period of time, the connection fails. The ISAKMP SA and the corresponding IPsec SA are deleted, and the tunnel is also deleted.
	//
	// - **false**: Disabled. The initiator of the IPsec-VPN connection does not send DPD packets.
	//
	// example:
	//
	// true
	EnableDpd *bool `json:"EnableDpd,omitempty" xml:"EnableDpd,omitempty"`
	// Specifies whether to enable NAT traversal. Valid values:
	//
	// - **true**: Enabled. After NAT traversal is enabled, the verification of the UDP port number is removed during IKE negotiations, and the NAT gateway device in the VPN tunnel can be discovered.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// true
	EnableNatTraversal *bool `json:"EnableNatTraversal,omitempty" xml:"EnableNatTraversal,omitempty"`
	// The CA certificate of the peer when you use an IPsec-VPN connection with a Chinese SM VPN gateway.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- MIIB7zCCAZW***	- -----END CERTIFICATE-----
	RemoteCaCertificate *string `json:"RemoteCaCertificate,omitempty" xml:"RemoteCaCertificate,omitempty"`
	// The BGP configuration of the tunnel to modify.
	//
	// If BGP was not previously enabled for the tunnel, call the [ModifyVpnConnectionAttribute](https://help.aliyun.com/document_detail/120381.html) operation to enable BGP for the tunnel and add the BGP configuration.
	TunnelBgpConfig *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig `json:"TunnelBgpConfig,omitempty" xml:"TunnelBgpConfig,omitempty" type:"Struct"`
	// The IKE phase (Phase 1) configuration of the tunnel to modify.
	TunnelIkeConfig *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig `json:"TunnelIkeConfig,omitempty" xml:"TunnelIkeConfig,omitempty" type:"Struct"`
	// The IPsec phase (Phase 2) configuration of the tunnel to modify.
	TunnelIpsecConfig *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig `json:"TunnelIpsecConfig,omitempty" xml:"TunnelIpsecConfig,omitempty" type:"Struct"`
}

func (s ModifyTunnelAttributeRequestTunnelOptionsSpecification) String() string {
	return dara.Prettify(s)
}

func (s ModifyTunnelAttributeRequestTunnelOptionsSpecification) GoString() string {
	return s.String()
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecification) GetCustomerGatewayId() *string {
	return s.CustomerGatewayId
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecification) GetEnableDpd() *bool {
	return s.EnableDpd
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecification) GetEnableNatTraversal() *bool {
	return s.EnableNatTraversal
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecification) GetRemoteCaCertificate() *string {
	return s.RemoteCaCertificate
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecification) GetTunnelBgpConfig() *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig {
	return s.TunnelBgpConfig
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecification) GetTunnelIkeConfig() *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig {
	return s.TunnelIkeConfig
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecification) GetTunnelIpsecConfig() *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig {
	return s.TunnelIpsecConfig
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecification) SetCustomerGatewayId(v string) *ModifyTunnelAttributeRequestTunnelOptionsSpecification {
	s.CustomerGatewayId = &v
	return s
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecification) SetEnableDpd(v bool) *ModifyTunnelAttributeRequestTunnelOptionsSpecification {
	s.EnableDpd = &v
	return s
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecification) SetEnableNatTraversal(v bool) *ModifyTunnelAttributeRequestTunnelOptionsSpecification {
	s.EnableNatTraversal = &v
	return s
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecification) SetRemoteCaCertificate(v string) *ModifyTunnelAttributeRequestTunnelOptionsSpecification {
	s.RemoteCaCertificate = &v
	return s
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecification) SetTunnelBgpConfig(v *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig) *ModifyTunnelAttributeRequestTunnelOptionsSpecification {
	s.TunnelBgpConfig = v
	return s
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecification) SetTunnelIkeConfig(v *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) *ModifyTunnelAttributeRequestTunnelOptionsSpecification {
	s.TunnelIkeConfig = v
	return s
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecification) SetTunnelIpsecConfig(v *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) *ModifyTunnelAttributeRequestTunnelOptionsSpecification {
	s.TunnelIpsecConfig = v
	return s
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecification) Validate() error {
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

type ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig struct {
	// The autonomous system number (ASN) of the local end of the tunnel. Valid values: **1*	- to **4294967295**.
	//
	// example:
	//
	// 65530
	LocalAsn *int64 `json:"LocalAsn,omitempty" xml:"LocalAsn,omitempty"`
	// The BGP IP address of the local end of the tunnel. The IP address must fall within the **TunnelCidr*	- CIDR block.
	//
	// example:
	//
	// 169.254.11.1
	LocalBgpIp *string `json:"LocalBgpIp,omitempty" xml:"LocalBgpIp,omitempty"`
	// The CIDR block of the BGP IP address on the local end of the tunnel.
	//
	// The CIDR block must be a CIDR block with a mask length of 30 within 169.254.0.0/16 and cannot be 169.254.0.0/30, 169.254.1.0/30, 169.254.2.0/30, 169.254.3.0/30, 169.254.4.0/30, 169.254.5.0/30, 169.254.6.0/30, or 169.254.169.252/30.
	//
	// > The IPsec tunnel CIDR block of each IPsec-VPN connection under a VPN gateway instance must be unique.
	//
	// example:
	//
	// 169.254.11.0/30
	TunnelCidr *string `json:"TunnelCidr,omitempty" xml:"TunnelCidr,omitempty"`
}

func (s ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig) GoString() string {
	return s.String()
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig) GetLocalAsn() *int64 {
	return s.LocalAsn
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig) GetLocalBgpIp() *string {
	return s.LocalBgpIp
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig) GetTunnelCidr() *string {
	return s.TunnelCidr
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig) SetLocalAsn(v int64) *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig {
	s.LocalAsn = &v
	return s
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig) SetLocalBgpIp(v string) *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig {
	s.LocalBgpIp = &v
	return s
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig) SetTunnelCidr(v string) *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig {
	s.TunnelCidr = &v
	return s
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelBgpConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig struct {
	// The authentication algorithm used in Phase 1 negotiations.
	//
	// <props="china">
	//
	// - If the IPsec-VPN connection is associated with a standard VPN gateway, valid values are **md5**, **sha1**, **sha256**, **sha384**, and **sha512**.
	//
	// - If the IPsec-VPN connection is associated with a Chinese SM VPN gateway, the value is **sm3**.
	//
	//
	//
	//
	// <props="intl">
	//
	// Valid values: **md5**, **sha1**, **sha256**, **sha384**, and **sha512**.
	//
	// .
	//
	// example:
	//
	// sha1
	IkeAuthAlg *string `json:"IkeAuthAlg,omitempty" xml:"IkeAuthAlg,omitempty"`
	// The encryption algorithm used in Phase 1 negotiations.
	//
	// <props="china">
	//
	// - If the IPsec-VPN connection is associated with a standard VPN gateway, valid values are **aes**, **aes192**, **aes256**, **des**, and **3des**.
	//
	// - If the IPsec-VPN connection is associated with a Chinese SM VPN gateway, the value is **sm4**.
	//
	//
	//
	// <props="intl">
	//
	// Valid values: **aes**, **aes192**, **aes256**, **des**, and **3des**.
	//
	// .
	//
	// example:
	//
	// aes
	IkeEncAlg *string `json:"IkeEncAlg,omitempty" xml:"IkeEncAlg,omitempty"`
	// The lifetime of the SA negotiated in Phase 1. Unit: seconds. Valid values: **0*	- to **86400**.
	//
	// example:
	//
	// 86400
	IkeLifetime *int64 `json:"IkeLifetime,omitempty" xml:"IkeLifetime,omitempty"`
	// The negotiation mode of the IKE version. Valid values:
	//
	// - **main**: main mode. This mode offers high security during negotiations.
	//
	// - **aggressive**: aggressive mode. This mode supports fast negotiations and a higher success rate.
	//
	// example:
	//
	// main
	IkeMode *string `json:"IkeMode,omitempty" xml:"IkeMode,omitempty"`
	// The Diffie-Hellman key exchange algorithm used in Phase 1 negotiations. Valid values: **group1**, **group2**, **group5**, and **group14**.
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
	// The identifier of the local end of the tunnel. The identifier can be up to 100 characters in length and cannot contain spaces. It supports FQDN and IP formats. Default value: the IP address of the tunnel.
	//
	// example:
	//
	// 47.XX.XX.87
	LocalId *string `json:"LocalId,omitempty" xml:"LocalId,omitempty"`
	// The pre-shared key used for identity authentication between the tunnel and the peer.
	//
	//
	//
	//    - The key must be 1 to 100 characters in length and can contain digits, uppercase letters, lowercase letters, and the following characters. It cannot contain spaces. ```~!`@#$%^&*()_-+={}[]|;:\\",.<>/?```
	//
	//    - If you do not specify a pre-shared key, the system generates a random 16-character string as the pre-shared key. You can call the [DescribeVpnConnection](https://help.aliyun.com/document_detail/120374.html) operation to query the pre-shared key that is automatically generated by the system.
	//
	//    > The pre-shared keys configured on the tunnel and the peer must be the same. Otherwise, the tunnel cannot be established.
	//
	// example:
	//
	// 123456****
	Psk *string `json:"Psk,omitempty" xml:"Psk,omitempty"`
	// The identifier of the peer end of the tunnel. The identifier can be up to 100 characters in length and cannot contain spaces. It supports FQDN and IP formats. Default value: the IP address of the customer gateway instance associated with the tunnel.
	//
	// example:
	//
	// 47.XX.XX.207
	RemoteId *string `json:"RemoteId,omitempty" xml:"RemoteId,omitempty"`
}

func (s ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) GoString() string {
	return s.String()
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) GetIkeAuthAlg() *string {
	return s.IkeAuthAlg
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) GetIkeEncAlg() *string {
	return s.IkeEncAlg
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) GetIkeLifetime() *int64 {
	return s.IkeLifetime
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) GetIkeMode() *string {
	return s.IkeMode
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) GetIkePfs() *string {
	return s.IkePfs
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) GetIkeVersion() *string {
	return s.IkeVersion
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) GetLocalId() *string {
	return s.LocalId
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) GetPsk() *string {
	return s.Psk
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) GetRemoteId() *string {
	return s.RemoteId
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) SetIkeAuthAlg(v string) *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkeAuthAlg = &v
	return s
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) SetIkeEncAlg(v string) *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkeEncAlg = &v
	return s
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) SetIkeLifetime(v int64) *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkeLifetime = &v
	return s
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) SetIkeMode(v string) *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkeMode = &v
	return s
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) SetIkePfs(v string) *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkePfs = &v
	return s
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) SetIkeVersion(v string) *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkeVersion = &v
	return s
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) SetLocalId(v string) *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.LocalId = &v
	return s
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) SetPsk(v string) *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.Psk = &v
	return s
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) SetRemoteId(v string) *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig {
	s.RemoteId = &v
	return s
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIkeConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig struct {
	// The authentication algorithm used in Phase 2 negotiations.
	//
	// <props="china">
	//
	// - If the IPsec-VPN connection is associated with a standard VPN gateway, valid values are **md5**, **sha1**, **sha256**, **sha384**, and **sha512**.
	//
	// - If the IPsec-VPN connection is associated with a Chinese SM VPN gateway, the value is **sm3**.
	//
	//
	//
	// <props="intl">
	//
	// Valid values: **md5**, **sha1**, **sha256**, **sha384**, and **sha512**.
	//
	// .
	//
	// example:
	//
	// sha1
	IpsecAuthAlg *string `json:"IpsecAuthAlg,omitempty" xml:"IpsecAuthAlg,omitempty"`
	// The encryption algorithm used in Phase 2 negotiations.
	//
	// <props="china">
	//
	// - If the IPsec-VPN connection is associated with a standard VPN gateway, valid values are **aes**, **aes192**, **aes256**, **des**, and **3des**.
	//
	// - If the IPsec-VPN connection is associated with a Chinese SM VPN gateway, the value is **sm4**.
	//
	//
	//
	// <props="intl">
	//
	// Valid values: **aes**, **aes192**, **aes256**, **des**, and **3des**.
	//
	// .
	//
	// example:
	//
	// aes
	IpsecEncAlg *string `json:"IpsecEncAlg,omitempty" xml:"IpsecEncAlg,omitempty"`
	// The lifetime of the SA negotiated in Phase 2. Unit: seconds. Valid values: **0*	- to **86400**.
	//
	// example:
	//
	// 86400
	IpsecLifetime *int64 `json:"IpsecLifetime,omitempty" xml:"IpsecLifetime,omitempty"`
	// The Diffie-Hellman key exchange algorithm used in Phase 2 negotiations. Valid values: **disabled**, **group1**, **group2**, **group5**, and **group14**.
	//
	// example:
	//
	// group2
	IpsecPfs *string `json:"IpsecPfs,omitempty" xml:"IpsecPfs,omitempty"`
}

func (s ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) GoString() string {
	return s.String()
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) GetIpsecAuthAlg() *string {
	return s.IpsecAuthAlg
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) GetIpsecEncAlg() *string {
	return s.IpsecEncAlg
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) GetIpsecLifetime() *int64 {
	return s.IpsecLifetime
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) GetIpsecPfs() *string {
	return s.IpsecPfs
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) SetIpsecAuthAlg(v string) *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig {
	s.IpsecAuthAlg = &v
	return s
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) SetIpsecEncAlg(v string) *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig {
	s.IpsecEncAlg = &v
	return s
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) SetIpsecLifetime(v int64) *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig {
	s.IpsecLifetime = &v
	return s
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) SetIpsecPfs(v string) *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig {
	s.IpsecPfs = &v
	return s
}

func (s *ModifyTunnelAttributeRequestTunnelOptionsSpecificationTunnelIpsecConfig) Validate() error {
	return dara.Validate(s)
}
