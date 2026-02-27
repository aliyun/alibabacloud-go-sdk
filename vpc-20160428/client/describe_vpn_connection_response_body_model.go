// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpnConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttachInstanceId(v string) *DescribeVpnConnectionResponseBody
	GetAttachInstanceId() *string
	SetAttachType(v string) *DescribeVpnConnectionResponseBody
	GetAttachType() *string
	SetCreateTime(v int64) *DescribeVpnConnectionResponseBody
	GetCreateTime() *int64
	SetCrossAccountAuthorized(v bool) *DescribeVpnConnectionResponseBody
	GetCrossAccountAuthorized() *bool
	SetCustomerGatewayId(v string) *DescribeVpnConnectionResponseBody
	GetCustomerGatewayId() *string
	SetEffectImmediately(v bool) *DescribeVpnConnectionResponseBody
	GetEffectImmediately() *bool
	SetEnableDpd(v bool) *DescribeVpnConnectionResponseBody
	GetEnableDpd() *bool
	SetEnableNatTraversal(v bool) *DescribeVpnConnectionResponseBody
	GetEnableNatTraversal() *bool
	SetEnableTunnelsBgp(v bool) *DescribeVpnConnectionResponseBody
	GetEnableTunnelsBgp() *bool
	SetIkeConfig(v *DescribeVpnConnectionResponseBodyIkeConfig) *DescribeVpnConnectionResponseBody
	GetIkeConfig() *DescribeVpnConnectionResponseBodyIkeConfig
	SetInternetIp(v string) *DescribeVpnConnectionResponseBody
	GetInternetIp() *string
	SetIpsecConfig(v *DescribeVpnConnectionResponseBodyIpsecConfig) *DescribeVpnConnectionResponseBody
	GetIpsecConfig() *DescribeVpnConnectionResponseBodyIpsecConfig
	SetLocalSubnet(v string) *DescribeVpnConnectionResponseBody
	GetLocalSubnet() *string
	SetName(v string) *DescribeVpnConnectionResponseBody
	GetName() *string
	SetNetworkType(v string) *DescribeVpnConnectionResponseBody
	GetNetworkType() *string
	SetRemoteCaCertificate(v string) *DescribeVpnConnectionResponseBody
	GetRemoteCaCertificate() *string
	SetRemoteSubnet(v string) *DescribeVpnConnectionResponseBody
	GetRemoteSubnet() *string
	SetRequestId(v string) *DescribeVpnConnectionResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeVpnConnectionResponseBody
	GetResourceGroupId() *string
	SetSpec(v string) *DescribeVpnConnectionResponseBody
	GetSpec() *string
	SetState(v string) *DescribeVpnConnectionResponseBody
	GetState() *string
	SetStatus(v string) *DescribeVpnConnectionResponseBody
	GetStatus() *string
	SetTags(v *DescribeVpnConnectionResponseBodyTags) *DescribeVpnConnectionResponseBody
	GetTags() *DescribeVpnConnectionResponseBodyTags
	SetTransitRouterId(v string) *DescribeVpnConnectionResponseBody
	GetTransitRouterId() *string
	SetTransitRouterName(v string) *DescribeVpnConnectionResponseBody
	GetTransitRouterName() *string
	SetTunnelBandwidth(v string) *DescribeVpnConnectionResponseBody
	GetTunnelBandwidth() *string
	SetTunnelOptionsSpecification(v *DescribeVpnConnectionResponseBodyTunnelOptionsSpecification) *DescribeVpnConnectionResponseBody
	GetTunnelOptionsSpecification() *DescribeVpnConnectionResponseBodyTunnelOptionsSpecification
	SetVcoHealthCheck(v *DescribeVpnConnectionResponseBodyVcoHealthCheck) *DescribeVpnConnectionResponseBody
	GetVcoHealthCheck() *DescribeVpnConnectionResponseBodyVcoHealthCheck
	SetVpnBgpConfig(v *DescribeVpnConnectionResponseBodyVpnBgpConfig) *DescribeVpnConnectionResponseBody
	GetVpnBgpConfig() *DescribeVpnConnectionResponseBodyVpnBgpConfig
	SetVpnConnectionId(v string) *DescribeVpnConnectionResponseBody
	GetVpnConnectionId() *string
	SetVpnGatewayId(v string) *DescribeVpnConnectionResponseBody
	GetVpnGatewayId() *string
	SetZoneNo(v string) *DescribeVpnConnectionResponseBody
	GetZoneNo() *string
}

type DescribeVpnConnectionResponseBody struct {
	// The ID of the CEN instance to which the transit router belongs.
	//
	// example:
	//
	// cen-lxxpbpalc776qz****
	AttachInstanceId *string `json:"AttachInstanceId,omitempty" xml:"AttachInstanceId,omitempty"`
	// The type of the resource that is associated with the IPsec-VPN connection. Valid values:
	//
	// 	- **CEN**: indicates that the IPsec-VPN connection is associated with a transit router of a Cloud Enterprise Network (CEN) instance.
	//
	// 	- **NO_ASSOCIATED**: indicates that the IPsec-VPN connection is not associated with any resource.
	//
	// 	- **VPNGW**: indicates that the IPsec-VPN connection is associated with a VPN gateway.
	//
	// example:
	//
	// CEN
	AttachType *string `json:"AttachType,omitempty" xml:"AttachType,omitempty"`
	// The timestamp generated when the IPsec-VPN connection was established. Unit: milliseconds.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1492753817000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the IPsec-VPN connection is associated with a transit router that belongs to another Alibaba Cloud account. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	CrossAccountAuthorized *bool `json:"CrossAccountAuthorized,omitempty" xml:"CrossAccountAuthorized,omitempty"`
	// The ID of the customer gateway associated with the IPsec-VPN connection.
	//
	// example:
	//
	// cgw-bp1mvj4g9kogwwcxk****
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" xml:"CustomerGatewayId,omitempty"`
	// Indicates whether IPsec negotiations immediately start after the configuration takes effect. Valid values:
	//
	// 	- **true**: Negotiations are reinitiated after the configuration is changed.
	//
	// 	- **false**: Negotiations are reinitiated after traffic is detected.
	//
	// example:
	//
	// true
	EffectImmediately *bool `json:"EffectImmediately,omitempty" xml:"EffectImmediately,omitempty"`
	// Indicates whether the dead peer detection (DPD) feature is enabled for the IPsec-VPN connection. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// After you enable the DPD feature, the initiator of the IPsec-VPN connection sends DPD packets to check the existence and availability of the peer. If no response is received from the peer within a specified period of time, the connection fails. Then, the ISAKMP security association (SA), IPsec SA, and IPsec tunnel are deleted.
	//
	// example:
	//
	// true
	EnableDpd *bool `json:"EnableDpd,omitempty" xml:"EnableDpd,omitempty"`
	// Indicates whether NAT traversal is enabled for the IPsec-VPN connection. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// After NAT traversal is enabled, the initiator does not check the UDP ports during IKE negotiations and can automatically discover NAT gateway devices along the IPsec tunnel.
	//
	// example:
	//
	// true
	EnableNatTraversal *bool `json:"EnableNatTraversal,omitempty" xml:"EnableNatTraversal,omitempty"`
	// Indicates whether BGP is enabled for the tunnel. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	EnableTunnelsBgp *bool `json:"EnableTunnelsBgp,omitempty" xml:"EnableTunnelsBgp,omitempty"`
	// The configuration of Phase 1 negotiations.
	IkeConfig *DescribeVpnConnectionResponseBodyIkeConfig `json:"IkeConfig,omitempty" xml:"IkeConfig,omitempty" type:"Struct"`
	// The gateway IP address of the IPsec-VPN connection.
	//
	// example:
	//
	// 47.XX.XX.162
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The configuration of Phase 2 negotiations.
	IpsecConfig *DescribeVpnConnectionResponseBodyIpsecConfig `json:"IpsecConfig,omitempty" xml:"IpsecConfig,omitempty" type:"Struct"`
	// The CIDR block on the Alibaba Cloud side.
	//
	// Multiple CIDR blocks are separated by commas (,).
	//
	// example:
	//
	// 10.0.0.0/8
	LocalSubnet *string `json:"LocalSubnet,omitempty" xml:"LocalSubnet,omitempty"`
	// The name of the IPsec-VPN connection.
	//
	// example:
	//
	// ipsec1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The network type of the IPsec-VPN connection. Valid values:
	//
	// 	- **public**: an encrypted connection over the Internet
	//
	// 	- **private**: an encrypted connection over private networks
	//
	// example:
	//
	// public
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The certificate authority (CA) certificate of the peer.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- MIIB7zCCAZW****
	RemoteCaCertificate *string `json:"RemoteCaCertificate,omitempty" xml:"RemoteCaCertificate,omitempty"`
	// The CIDR block on the data center side.
	//
	// Multiple CIDR blocks are separated by commas (,).
	//
	// example:
	//
	// 192.168.0.0/16
	RemoteSubnet *string `json:"RemoteSubnet,omitempty" xml:"RemoteSubnet,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F2310D45-BCF6-4E2E-9082-B4503844BA4C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the IPsec-VPN connection belongs.
	//
	// You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation to query the resource group information.
	//
	// example:
	//
	// rg-acfmzs372yg****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The bandwidth specification of the IPsec-VPN connection. Unit: **Mbit/s**.
	//
	// example:
	//
	// 1000M
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The association state of the IPsec-VPN connection. Valid values:
	//
	// 	- **active**: The IPsec-VPN connection is associated with a VPN gateway.
	//
	// 	- **init**: The IPsec-VPN connection is not associated with any resource and is being initialized.
	//
	// 	- **attaching**: The IPsec-VPN connection is being associated with a transit router.
	//
	// 	- **attached**: The IPsec-VPN connection is associated with a transit router.
	//
	// 	- **detaching**: The IPsec-VPN connection is being disassociated from a transit router.
	//
	// 	- **financialLocked**: The IPsec-VPN connection is locked due to overdue payments.
	//
	// 	- **provisioning**: The IPsec-VPN connection is being prepared.
	//
	// 	- **updating**: The IPsec-VPN connection is being updated.
	//
	// 	- **Upgrading**: The IPsec-VPN connection is being upgraded.
	//
	// 	- **deleted**: The IPsec-VPN connection is deleted.
	//
	// example:
	//
	// attached
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The state of the IPsec-VPN connection. Valid values:
	//
	// 	- **ike_sa_not_established**: Phase 1 negotiations failed.
	//
	// 	- **ike_sa_established**: Phase 1 negotiations succeeded.
	//
	// 	- **ipsec_sa_not_established**: Phase 2 negotiations failed.
	//
	// 	- **ipsec_sa_established**: Phase 2 negotiations succeeded.
	//
	// example:
	//
	// ike_sa_not_established
	Status *string                                `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   *DescribeVpnConnectionResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The ID of the transit router with which the IPsec-VPN connection is associated.
	//
	// example:
	//
	// tr-p0we2edef9qr44a85****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The name of the transit router.
	//
	// example:
	//
	// nametest
	TransitRouterName          *string                                                      `json:"TransitRouterName,omitempty" xml:"TransitRouterName,omitempty"`
	TunnelBandwidth            *string                                                      `json:"TunnelBandwidth,omitempty" xml:"TunnelBandwidth,omitempty"`
	TunnelOptionsSpecification *DescribeVpnConnectionResponseBodyTunnelOptionsSpecification `json:"TunnelOptionsSpecification,omitempty" xml:"TunnelOptionsSpecification,omitempty" type:"Struct"`
	// The health check information about the IPsec-VPN connection.
	VcoHealthCheck *DescribeVpnConnectionResponseBodyVcoHealthCheck `json:"VcoHealthCheck,omitempty" xml:"VcoHealthCheck,omitempty" type:"Struct"`
	// The Border Gateway Protocol (BGP) configuration of the IPsec-VPN connection.
	VpnBgpConfig *DescribeVpnConnectionResponseBodyVpnBgpConfig `json:"VpnBgpConfig,omitempty" xml:"VpnBgpConfig,omitempty" type:"Struct"`
	// The ID of the IPsec-VPN connection.
	//
	// example:
	//
	// vco-bp1bbi27hojx80nck****
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" xml:"VpnConnectionId,omitempty"`
	// The ID of the VPN gateway.
	//
	// example:
	//
	// vpn-bp1q8bgx4xnkm2ogj****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
	// The ID of the zone where the IPsec-VPN connection is deployed.
	//
	// You can call [DescribeZones](https://help.aliyun.com/document_detail/36064.html) to query zone IDs and mapping between zone IDs and zone names.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneNo *string `json:"ZoneNo,omitempty" xml:"ZoneNo,omitempty"`
}

func (s DescribeVpnConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpnConnectionResponseBody) GetAttachInstanceId() *string {
	return s.AttachInstanceId
}

func (s *DescribeVpnConnectionResponseBody) GetAttachType() *string {
	return s.AttachType
}

func (s *DescribeVpnConnectionResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeVpnConnectionResponseBody) GetCrossAccountAuthorized() *bool {
	return s.CrossAccountAuthorized
}

func (s *DescribeVpnConnectionResponseBody) GetCustomerGatewayId() *string {
	return s.CustomerGatewayId
}

func (s *DescribeVpnConnectionResponseBody) GetEffectImmediately() *bool {
	return s.EffectImmediately
}

func (s *DescribeVpnConnectionResponseBody) GetEnableDpd() *bool {
	return s.EnableDpd
}

func (s *DescribeVpnConnectionResponseBody) GetEnableNatTraversal() *bool {
	return s.EnableNatTraversal
}

func (s *DescribeVpnConnectionResponseBody) GetEnableTunnelsBgp() *bool {
	return s.EnableTunnelsBgp
}

func (s *DescribeVpnConnectionResponseBody) GetIkeConfig() *DescribeVpnConnectionResponseBodyIkeConfig {
	return s.IkeConfig
}

func (s *DescribeVpnConnectionResponseBody) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeVpnConnectionResponseBody) GetIpsecConfig() *DescribeVpnConnectionResponseBodyIpsecConfig {
	return s.IpsecConfig
}

func (s *DescribeVpnConnectionResponseBody) GetLocalSubnet() *string {
	return s.LocalSubnet
}

func (s *DescribeVpnConnectionResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeVpnConnectionResponseBody) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeVpnConnectionResponseBody) GetRemoteCaCertificate() *string {
	return s.RemoteCaCertificate
}

func (s *DescribeVpnConnectionResponseBody) GetRemoteSubnet() *string {
	return s.RemoteSubnet
}

func (s *DescribeVpnConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpnConnectionResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeVpnConnectionResponseBody) GetSpec() *string {
	return s.Spec
}

func (s *DescribeVpnConnectionResponseBody) GetState() *string {
	return s.State
}

func (s *DescribeVpnConnectionResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeVpnConnectionResponseBody) GetTags() *DescribeVpnConnectionResponseBodyTags {
	return s.Tags
}

func (s *DescribeVpnConnectionResponseBody) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *DescribeVpnConnectionResponseBody) GetTransitRouterName() *string {
	return s.TransitRouterName
}

func (s *DescribeVpnConnectionResponseBody) GetTunnelBandwidth() *string {
	return s.TunnelBandwidth
}

func (s *DescribeVpnConnectionResponseBody) GetTunnelOptionsSpecification() *DescribeVpnConnectionResponseBodyTunnelOptionsSpecification {
	return s.TunnelOptionsSpecification
}

func (s *DescribeVpnConnectionResponseBody) GetVcoHealthCheck() *DescribeVpnConnectionResponseBodyVcoHealthCheck {
	return s.VcoHealthCheck
}

func (s *DescribeVpnConnectionResponseBody) GetVpnBgpConfig() *DescribeVpnConnectionResponseBodyVpnBgpConfig {
	return s.VpnBgpConfig
}

func (s *DescribeVpnConnectionResponseBody) GetVpnConnectionId() *string {
	return s.VpnConnectionId
}

func (s *DescribeVpnConnectionResponseBody) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *DescribeVpnConnectionResponseBody) GetZoneNo() *string {
	return s.ZoneNo
}

func (s *DescribeVpnConnectionResponseBody) SetAttachInstanceId(v string) *DescribeVpnConnectionResponseBody {
	s.AttachInstanceId = &v
	return s
}

func (s *DescribeVpnConnectionResponseBody) SetAttachType(v string) *DescribeVpnConnectionResponseBody {
	s.AttachType = &v
	return s
}

func (s *DescribeVpnConnectionResponseBody) SetCreateTime(v int64) *DescribeVpnConnectionResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeVpnConnectionResponseBody) SetCrossAccountAuthorized(v bool) *DescribeVpnConnectionResponseBody {
	s.CrossAccountAuthorized = &v
	return s
}

func (s *DescribeVpnConnectionResponseBody) SetCustomerGatewayId(v string) *DescribeVpnConnectionResponseBody {
	s.CustomerGatewayId = &v
	return s
}

func (s *DescribeVpnConnectionResponseBody) SetEffectImmediately(v bool) *DescribeVpnConnectionResponseBody {
	s.EffectImmediately = &v
	return s
}

func (s *DescribeVpnConnectionResponseBody) SetEnableDpd(v bool) *DescribeVpnConnectionResponseBody {
	s.EnableDpd = &v
	return s
}

func (s *DescribeVpnConnectionResponseBody) SetEnableNatTraversal(v bool) *DescribeVpnConnectionResponseBody {
	s.EnableNatTraversal = &v
	return s
}

func (s *DescribeVpnConnectionResponseBody) SetEnableTunnelsBgp(v bool) *DescribeVpnConnectionResponseBody {
	s.EnableTunnelsBgp = &v
	return s
}

func (s *DescribeVpnConnectionResponseBody) SetIkeConfig(v *DescribeVpnConnectionResponseBodyIkeConfig) *DescribeVpnConnectionResponseBody {
	s.IkeConfig = v
	return s
}

func (s *DescribeVpnConnectionResponseBody) SetInternetIp(v string) *DescribeVpnConnectionResponseBody {
	s.InternetIp = &v
	return s
}

func (s *DescribeVpnConnectionResponseBody) SetIpsecConfig(v *DescribeVpnConnectionResponseBodyIpsecConfig) *DescribeVpnConnectionResponseBody {
	s.IpsecConfig = v
	return s
}

func (s *DescribeVpnConnectionResponseBody) SetLocalSubnet(v string) *DescribeVpnConnectionResponseBody {
	s.LocalSubnet = &v
	return s
}

func (s *DescribeVpnConnectionResponseBody) SetName(v string) *DescribeVpnConnectionResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeVpnConnectionResponseBody) SetNetworkType(v string) *DescribeVpnConnectionResponseBody {
	s.NetworkType = &v
	return s
}

func (s *DescribeVpnConnectionResponseBody) SetRemoteCaCertificate(v string) *DescribeVpnConnectionResponseBody {
	s.RemoteCaCertificate = &v
	return s
}

func (s *DescribeVpnConnectionResponseBody) SetRemoteSubnet(v string) *DescribeVpnConnectionResponseBody {
	s.RemoteSubnet = &v
	return s
}

func (s *DescribeVpnConnectionResponseBody) SetRequestId(v string) *DescribeVpnConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpnConnectionResponseBody) SetResourceGroupId(v string) *DescribeVpnConnectionResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeVpnConnectionResponseBody) SetSpec(v string) *DescribeVpnConnectionResponseBody {
	s.Spec = &v
	return s
}

func (s *DescribeVpnConnectionResponseBody) SetState(v string) *DescribeVpnConnectionResponseBody {
	s.State = &v
	return s
}

func (s *DescribeVpnConnectionResponseBody) SetStatus(v string) *DescribeVpnConnectionResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeVpnConnectionResponseBody) SetTags(v *DescribeVpnConnectionResponseBodyTags) *DescribeVpnConnectionResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeVpnConnectionResponseBody) SetTransitRouterId(v string) *DescribeVpnConnectionResponseBody {
	s.TransitRouterId = &v
	return s
}

func (s *DescribeVpnConnectionResponseBody) SetTransitRouterName(v string) *DescribeVpnConnectionResponseBody {
	s.TransitRouterName = &v
	return s
}

func (s *DescribeVpnConnectionResponseBody) SetTunnelBandwidth(v string) *DescribeVpnConnectionResponseBody {
	s.TunnelBandwidth = &v
	return s
}

func (s *DescribeVpnConnectionResponseBody) SetTunnelOptionsSpecification(v *DescribeVpnConnectionResponseBodyTunnelOptionsSpecification) *DescribeVpnConnectionResponseBody {
	s.TunnelOptionsSpecification = v
	return s
}

func (s *DescribeVpnConnectionResponseBody) SetVcoHealthCheck(v *DescribeVpnConnectionResponseBodyVcoHealthCheck) *DescribeVpnConnectionResponseBody {
	s.VcoHealthCheck = v
	return s
}

func (s *DescribeVpnConnectionResponseBody) SetVpnBgpConfig(v *DescribeVpnConnectionResponseBodyVpnBgpConfig) *DescribeVpnConnectionResponseBody {
	s.VpnBgpConfig = v
	return s
}

func (s *DescribeVpnConnectionResponseBody) SetVpnConnectionId(v string) *DescribeVpnConnectionResponseBody {
	s.VpnConnectionId = &v
	return s
}

func (s *DescribeVpnConnectionResponseBody) SetVpnGatewayId(v string) *DescribeVpnConnectionResponseBody {
	s.VpnGatewayId = &v
	return s
}

func (s *DescribeVpnConnectionResponseBody) SetZoneNo(v string) *DescribeVpnConnectionResponseBody {
	s.ZoneNo = &v
	return s
}

func (s *DescribeVpnConnectionResponseBody) Validate() error {
	if s.IkeConfig != nil {
		if err := s.IkeConfig.Validate(); err != nil {
			return err
		}
	}
	if s.IpsecConfig != nil {
		if err := s.IpsecConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	if s.TunnelOptionsSpecification != nil {
		if err := s.TunnelOptionsSpecification.Validate(); err != nil {
			return err
		}
	}
	if s.VcoHealthCheck != nil {
		if err := s.VcoHealthCheck.Validate(); err != nil {
			return err
		}
	}
	if s.VpnBgpConfig != nil {
		if err := s.VpnBgpConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVpnConnectionResponseBodyIkeConfig struct {
	// The authentication algorithm in the IKE phase.
	//
	// example:
	//
	// sha1
	IkeAuthAlg *string `json:"IkeAuthAlg,omitempty" xml:"IkeAuthAlg,omitempty"`
	// The encryption algorithm in the IKE phase.
	//
	// example:
	//
	// aes
	IkeEncAlg *string `json:"IkeEncAlg,omitempty" xml:"IkeEncAlg,omitempty"`
	// The lifetime in the IKE phase. Unit: seconds.
	//
	// example:
	//
	// 86400
	IkeLifetime *int64 `json:"IkeLifetime,omitempty" xml:"IkeLifetime,omitempty"`
	// The IKE negotiation mode.
	//
	// 	- **main**: This mode offers higher security during negotiations.
	//
	// 	- **aggressive**: This mode is faster and has a higher success rate.
	//
	// example:
	//
	// main
	IkeMode *string `json:"IkeMode,omitempty" xml:"IkeMode,omitempty"`
	// The Diffie-Hellman (DH) group in the IKE phase.
	//
	// example:
	//
	// group2
	IkePfs *string `json:"IkePfs,omitempty" xml:"IkePfs,omitempty"`
	// The version of the IKE protocol.
	//
	// 	- **ikev1**
	//
	// 	- **ikev2**
	//
	// Compared with IKEv1, IKEv2 simplifies the SA negotiation process and is more suitable for scenarios in which multiple CIDR blocks are used.
	//
	// example:
	//
	// ikev1
	IkeVersion *string `json:"IkeVersion,omitempty" xml:"IkeVersion,omitempty"`
	// The identifier of the IPsec-VPN connection on the Alibaba Cloud side.
	//
	// example:
	//
	// 116.28.XX.XX
	LocalId *string `json:"LocalId,omitempty" xml:"LocalId,omitempty"`
	// The pre-shared key.
	//
	// example:
	//
	// pgw6dy****
	Psk *string `json:"Psk,omitempty" xml:"Psk,omitempty"`
	// The identifier of the IPsec-VPN connection on the data center side.
	//
	// example:
	//
	// 139.34.XX.XX
	RemoteId *string `json:"RemoteId,omitempty" xml:"RemoteId,omitempty"`
}

func (s DescribeVpnConnectionResponseBodyIkeConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnConnectionResponseBodyIkeConfig) GoString() string {
	return s.String()
}

func (s *DescribeVpnConnectionResponseBodyIkeConfig) GetIkeAuthAlg() *string {
	return s.IkeAuthAlg
}

func (s *DescribeVpnConnectionResponseBodyIkeConfig) GetIkeEncAlg() *string {
	return s.IkeEncAlg
}

func (s *DescribeVpnConnectionResponseBodyIkeConfig) GetIkeLifetime() *int64 {
	return s.IkeLifetime
}

func (s *DescribeVpnConnectionResponseBodyIkeConfig) GetIkeMode() *string {
	return s.IkeMode
}

func (s *DescribeVpnConnectionResponseBodyIkeConfig) GetIkePfs() *string {
	return s.IkePfs
}

func (s *DescribeVpnConnectionResponseBodyIkeConfig) GetIkeVersion() *string {
	return s.IkeVersion
}

func (s *DescribeVpnConnectionResponseBodyIkeConfig) GetLocalId() *string {
	return s.LocalId
}

func (s *DescribeVpnConnectionResponseBodyIkeConfig) GetPsk() *string {
	return s.Psk
}

func (s *DescribeVpnConnectionResponseBodyIkeConfig) GetRemoteId() *string {
	return s.RemoteId
}

func (s *DescribeVpnConnectionResponseBodyIkeConfig) SetIkeAuthAlg(v string) *DescribeVpnConnectionResponseBodyIkeConfig {
	s.IkeAuthAlg = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyIkeConfig) SetIkeEncAlg(v string) *DescribeVpnConnectionResponseBodyIkeConfig {
	s.IkeEncAlg = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyIkeConfig) SetIkeLifetime(v int64) *DescribeVpnConnectionResponseBodyIkeConfig {
	s.IkeLifetime = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyIkeConfig) SetIkeMode(v string) *DescribeVpnConnectionResponseBodyIkeConfig {
	s.IkeMode = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyIkeConfig) SetIkePfs(v string) *DescribeVpnConnectionResponseBodyIkeConfig {
	s.IkePfs = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyIkeConfig) SetIkeVersion(v string) *DescribeVpnConnectionResponseBodyIkeConfig {
	s.IkeVersion = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyIkeConfig) SetLocalId(v string) *DescribeVpnConnectionResponseBodyIkeConfig {
	s.LocalId = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyIkeConfig) SetPsk(v string) *DescribeVpnConnectionResponseBodyIkeConfig {
	s.Psk = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyIkeConfig) SetRemoteId(v string) *DescribeVpnConnectionResponseBodyIkeConfig {
	s.RemoteId = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyIkeConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeVpnConnectionResponseBodyIpsecConfig struct {
	// The authentication algorithm in the IPsec phase.
	//
	// example:
	//
	// sha1
	IpsecAuthAlg *string `json:"IpsecAuthAlg,omitempty" xml:"IpsecAuthAlg,omitempty"`
	// The encryption algorithm in the IPsec phase.
	//
	// example:
	//
	// aes
	IpsecEncAlg *string `json:"IpsecEncAlg,omitempty" xml:"IpsecEncAlg,omitempty"`
	// The lifetime in the IPsec phase. Unit: seconds.
	//
	// example:
	//
	// 86400
	IpsecLifetime *int64 `json:"IpsecLifetime,omitempty" xml:"IpsecLifetime,omitempty"`
	// The DH group in the IPsec phase.
	//
	// example:
	//
	// group2
	IpsecPfs *string `json:"IpsecPfs,omitempty" xml:"IpsecPfs,omitempty"`
}

func (s DescribeVpnConnectionResponseBodyIpsecConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnConnectionResponseBodyIpsecConfig) GoString() string {
	return s.String()
}

func (s *DescribeVpnConnectionResponseBodyIpsecConfig) GetIpsecAuthAlg() *string {
	return s.IpsecAuthAlg
}

func (s *DescribeVpnConnectionResponseBodyIpsecConfig) GetIpsecEncAlg() *string {
	return s.IpsecEncAlg
}

func (s *DescribeVpnConnectionResponseBodyIpsecConfig) GetIpsecLifetime() *int64 {
	return s.IpsecLifetime
}

func (s *DescribeVpnConnectionResponseBodyIpsecConfig) GetIpsecPfs() *string {
	return s.IpsecPfs
}

func (s *DescribeVpnConnectionResponseBodyIpsecConfig) SetIpsecAuthAlg(v string) *DescribeVpnConnectionResponseBodyIpsecConfig {
	s.IpsecAuthAlg = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyIpsecConfig) SetIpsecEncAlg(v string) *DescribeVpnConnectionResponseBodyIpsecConfig {
	s.IpsecEncAlg = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyIpsecConfig) SetIpsecLifetime(v int64) *DescribeVpnConnectionResponseBodyIpsecConfig {
	s.IpsecLifetime = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyIpsecConfig) SetIpsecPfs(v string) *DescribeVpnConnectionResponseBodyIpsecConfig {
	s.IpsecPfs = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyIpsecConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeVpnConnectionResponseBodyTags struct {
	Tag []*DescribeVpnConnectionResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeVpnConnectionResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnConnectionResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeVpnConnectionResponseBodyTags) GetTag() []*DescribeVpnConnectionResponseBodyTagsTag {
	return s.Tag
}

func (s *DescribeVpnConnectionResponseBodyTags) SetTag(v []*DescribeVpnConnectionResponseBodyTagsTag) *DescribeVpnConnectionResponseBodyTags {
	s.Tag = v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpnConnectionResponseBodyTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVpnConnectionResponseBodyTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnConnectionResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeVpnConnectionResponseBodyTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeVpnConnectionResponseBodyTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeVpnConnectionResponseBodyTagsTag) SetKey(v string) *DescribeVpnConnectionResponseBodyTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTagsTag) SetValue(v string) *DescribeVpnConnectionResponseBodyTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTagsTag) Validate() error {
	return dara.Validate(s)
}

type DescribeVpnConnectionResponseBodyTunnelOptionsSpecification struct {
	TunnelOptions []*DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions `json:"TunnelOptions,omitempty" xml:"TunnelOptions,omitempty" type:"Repeated"`
}

func (s DescribeVpnConnectionResponseBodyTunnelOptionsSpecification) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnConnectionResponseBodyTunnelOptionsSpecification) GoString() string {
	return s.String()
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecification) GetTunnelOptions() []*DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions {
	return s.TunnelOptions
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecification) SetTunnelOptions(v []*DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions) *DescribeVpnConnectionResponseBodyTunnelOptionsSpecification {
	s.TunnelOptions = v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecification) Validate() error {
	if s.TunnelOptions != nil {
		for _, item := range s.TunnelOptions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions struct {
	CustomerGatewayId   *string                                                                                    `json:"CustomerGatewayId,omitempty" xml:"CustomerGatewayId,omitempty"`
	EnableDpd           *string                                                                                    `json:"EnableDpd,omitempty" xml:"EnableDpd,omitempty"`
	EnableNatTraversal  *string                                                                                    `json:"EnableNatTraversal,omitempty" xml:"EnableNatTraversal,omitempty"`
	InternetIp          *string                                                                                    `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	RemoteCaCertificate *string                                                                                    `json:"RemoteCaCertificate,omitempty" xml:"RemoteCaCertificate,omitempty"`
	Role                *string                                                                                    `json:"Role,omitempty" xml:"Role,omitempty"`
	State               *string                                                                                    `json:"State,omitempty" xml:"State,omitempty"`
	Status              *string                                                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	TunnelBgpConfig     *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig   `json:"TunnelBgpConfig,omitempty" xml:"TunnelBgpConfig,omitempty" type:"Struct"`
	TunnelId            *string                                                                                    `json:"TunnelId,omitempty" xml:"TunnelId,omitempty"`
	TunnelIkeConfig     *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig   `json:"TunnelIkeConfig,omitempty" xml:"TunnelIkeConfig,omitempty" type:"Struct"`
	TunnelIndex         *int32                                                                                     `json:"TunnelIndex,omitempty" xml:"TunnelIndex,omitempty"`
	TunnelIpsecConfig   *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig `json:"TunnelIpsecConfig,omitempty" xml:"TunnelIpsecConfig,omitempty" type:"Struct"`
	ZoneNo              *string                                                                                    `json:"ZoneNo,omitempty" xml:"ZoneNo,omitempty"`
}

func (s DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions) GoString() string {
	return s.String()
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions) GetCustomerGatewayId() *string {
	return s.CustomerGatewayId
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions) GetEnableDpd() *string {
	return s.EnableDpd
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions) GetEnableNatTraversal() *string {
	return s.EnableNatTraversal
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions) GetRemoteCaCertificate() *string {
	return s.RemoteCaCertificate
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions) GetRole() *string {
	return s.Role
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions) GetState() *string {
	return s.State
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions) GetStatus() *string {
	return s.Status
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions) GetTunnelBgpConfig() *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig {
	return s.TunnelBgpConfig
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions) GetTunnelId() *string {
	return s.TunnelId
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions) GetTunnelIkeConfig() *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig {
	return s.TunnelIkeConfig
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions) GetTunnelIndex() *int32 {
	return s.TunnelIndex
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions) GetTunnelIpsecConfig() *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig {
	return s.TunnelIpsecConfig
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions) GetZoneNo() *string {
	return s.ZoneNo
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions) SetCustomerGatewayId(v string) *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions {
	s.CustomerGatewayId = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions) SetEnableDpd(v string) *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions {
	s.EnableDpd = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions) SetEnableNatTraversal(v string) *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions {
	s.EnableNatTraversal = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions) SetInternetIp(v string) *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions {
	s.InternetIp = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions) SetRemoteCaCertificate(v string) *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions {
	s.RemoteCaCertificate = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions) SetRole(v string) *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions {
	s.Role = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions) SetState(v string) *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions {
	s.State = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions) SetStatus(v string) *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions {
	s.Status = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions) SetTunnelBgpConfig(v *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions {
	s.TunnelBgpConfig = v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions) SetTunnelId(v string) *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions {
	s.TunnelId = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions) SetTunnelIkeConfig(v *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions {
	s.TunnelIkeConfig = v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions) SetTunnelIndex(v int32) *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions {
	s.TunnelIndex = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions) SetTunnelIpsecConfig(v *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions {
	s.TunnelIpsecConfig = v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions) SetZoneNo(v string) *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions {
	s.ZoneNo = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptions) Validate() error {
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

type DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig struct {
	BgpStatus  *string `json:"BgpStatus,omitempty" xml:"BgpStatus,omitempty"`
	LocalAsn   *string `json:"LocalAsn,omitempty" xml:"LocalAsn,omitempty"`
	LocalBgpIp *string `json:"LocalBgpIp,omitempty" xml:"LocalBgpIp,omitempty"`
	PeerAsn    *string `json:"PeerAsn,omitempty" xml:"PeerAsn,omitempty"`
	PeerBgpIp  *string `json:"PeerBgpIp,omitempty" xml:"PeerBgpIp,omitempty"`
	TunnelCidr *string `json:"TunnelCidr,omitempty" xml:"TunnelCidr,omitempty"`
}

func (s DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) GoString() string {
	return s.String()
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) GetBgpStatus() *string {
	return s.BgpStatus
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) GetLocalAsn() *string {
	return s.LocalAsn
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) GetLocalBgpIp() *string {
	return s.LocalBgpIp
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) GetPeerAsn() *string {
	return s.PeerAsn
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) GetPeerBgpIp() *string {
	return s.PeerBgpIp
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) GetTunnelCidr() *string {
	return s.TunnelCidr
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) SetBgpStatus(v string) *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig {
	s.BgpStatus = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) SetLocalAsn(v string) *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig {
	s.LocalAsn = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) SetLocalBgpIp(v string) *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig {
	s.LocalBgpIp = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) SetPeerAsn(v string) *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig {
	s.PeerAsn = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) SetPeerBgpIp(v string) *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig {
	s.PeerBgpIp = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) SetTunnelCidr(v string) *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig {
	s.TunnelCidr = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig struct {
	IkeAuthAlg  *string `json:"IkeAuthAlg,omitempty" xml:"IkeAuthAlg,omitempty"`
	IkeEncAlg   *string `json:"IkeEncAlg,omitempty" xml:"IkeEncAlg,omitempty"`
	IkeLifetime *string `json:"IkeLifetime,omitempty" xml:"IkeLifetime,omitempty"`
	IkeMode     *string `json:"IkeMode,omitempty" xml:"IkeMode,omitempty"`
	IkePfs      *string `json:"IkePfs,omitempty" xml:"IkePfs,omitempty"`
	IkeVersion  *string `json:"IkeVersion,omitempty" xml:"IkeVersion,omitempty"`
	LocalId     *string `json:"LocalId,omitempty" xml:"LocalId,omitempty"`
	Psk         *string `json:"Psk,omitempty" xml:"Psk,omitempty"`
	RemoteId    *string `json:"RemoteId,omitempty" xml:"RemoteId,omitempty"`
}

func (s DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) GoString() string {
	return s.String()
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) GetIkeAuthAlg() *string {
	return s.IkeAuthAlg
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) GetIkeEncAlg() *string {
	return s.IkeEncAlg
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) GetIkeLifetime() *string {
	return s.IkeLifetime
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) GetIkeMode() *string {
	return s.IkeMode
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) GetIkePfs() *string {
	return s.IkePfs
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) GetIkeVersion() *string {
	return s.IkeVersion
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) GetLocalId() *string {
	return s.LocalId
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) GetPsk() *string {
	return s.Psk
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) GetRemoteId() *string {
	return s.RemoteId
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) SetIkeAuthAlg(v string) *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig {
	s.IkeAuthAlg = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) SetIkeEncAlg(v string) *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig {
	s.IkeEncAlg = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) SetIkeLifetime(v string) *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig {
	s.IkeLifetime = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) SetIkeMode(v string) *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig {
	s.IkeMode = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) SetIkePfs(v string) *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig {
	s.IkePfs = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) SetIkeVersion(v string) *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig {
	s.IkeVersion = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) SetLocalId(v string) *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig {
	s.LocalId = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) SetPsk(v string) *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig {
	s.Psk = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) SetRemoteId(v string) *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig {
	s.RemoteId = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig struct {
	IpsecAuthAlg  *string `json:"IpsecAuthAlg,omitempty" xml:"IpsecAuthAlg,omitempty"`
	IpsecEncAlg   *string `json:"IpsecEncAlg,omitempty" xml:"IpsecEncAlg,omitempty"`
	IpsecLifetime *string `json:"IpsecLifetime,omitempty" xml:"IpsecLifetime,omitempty"`
	IpsecPfs      *string `json:"IpsecPfs,omitempty" xml:"IpsecPfs,omitempty"`
}

func (s DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) GoString() string {
	return s.String()
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) GetIpsecAuthAlg() *string {
	return s.IpsecAuthAlg
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) GetIpsecEncAlg() *string {
	return s.IpsecEncAlg
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) GetIpsecLifetime() *string {
	return s.IpsecLifetime
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) GetIpsecPfs() *string {
	return s.IpsecPfs
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) SetIpsecAuthAlg(v string) *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig {
	s.IpsecAuthAlg = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) SetIpsecEncAlg(v string) *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig {
	s.IpsecEncAlg = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) SetIpsecLifetime(v string) *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig {
	s.IpsecLifetime = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) SetIpsecPfs(v string) *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig {
	s.IpsecPfs = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeVpnConnectionResponseBodyVcoHealthCheck struct {
	// The destination IP address.
	//
	// example:
	//
	// 10.0.0.1
	Dip *string `json:"Dip,omitempty" xml:"Dip,omitempty"`
	// Indicates whether the health check feature is enabled for the IPsec-VPN connection. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// example:
	//
	// true
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The interval between two consecutive health checks. Unit: seconds.
	//
	// example:
	//
	// 3
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// Indicates whether advertised routes are withdrawn when the health check fails. Valid values:
	//
	// 	- **revoke_route**: Advertised routes are withdrawn.
	//
	// 	- **reserve_route**: Advertised routes are not withdrawn.
	//
	// example:
	//
	// revoke_route
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The maximum number of health check retries.
	//
	// example:
	//
	// 3
	Retry *int32 `json:"Retry,omitempty" xml:"Retry,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 192.168.1.1
	Sip *string `json:"Sip,omitempty" xml:"Sip,omitempty"`
	// The state of the health check. Valid values:
	//
	// 	- **failed**
	//
	// 	- **success**: normal
	//
	// example:
	//
	// failed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeVpnConnectionResponseBodyVcoHealthCheck) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnConnectionResponseBodyVcoHealthCheck) GoString() string {
	return s.String()
}

func (s *DescribeVpnConnectionResponseBodyVcoHealthCheck) GetDip() *string {
	return s.Dip
}

func (s *DescribeVpnConnectionResponseBodyVcoHealthCheck) GetEnable() *string {
	return s.Enable
}

func (s *DescribeVpnConnectionResponseBodyVcoHealthCheck) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeVpnConnectionResponseBodyVcoHealthCheck) GetPolicy() *string {
	return s.Policy
}

func (s *DescribeVpnConnectionResponseBodyVcoHealthCheck) GetRetry() *int32 {
	return s.Retry
}

func (s *DescribeVpnConnectionResponseBodyVcoHealthCheck) GetSip() *string {
	return s.Sip
}

func (s *DescribeVpnConnectionResponseBodyVcoHealthCheck) GetStatus() *string {
	return s.Status
}

func (s *DescribeVpnConnectionResponseBodyVcoHealthCheck) SetDip(v string) *DescribeVpnConnectionResponseBodyVcoHealthCheck {
	s.Dip = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyVcoHealthCheck) SetEnable(v string) *DescribeVpnConnectionResponseBodyVcoHealthCheck {
	s.Enable = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyVcoHealthCheck) SetInterval(v int32) *DescribeVpnConnectionResponseBodyVcoHealthCheck {
	s.Interval = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyVcoHealthCheck) SetPolicy(v string) *DescribeVpnConnectionResponseBodyVcoHealthCheck {
	s.Policy = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyVcoHealthCheck) SetRetry(v int32) *DescribeVpnConnectionResponseBodyVcoHealthCheck {
	s.Retry = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyVcoHealthCheck) SetSip(v string) *DescribeVpnConnectionResponseBodyVcoHealthCheck {
	s.Sip = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyVcoHealthCheck) SetStatus(v string) *DescribeVpnConnectionResponseBodyVcoHealthCheck {
	s.Status = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyVcoHealthCheck) Validate() error {
	return dara.Validate(s)
}

type DescribeVpnConnectionResponseBodyVpnBgpConfig struct {
	// The authentication key of the BGP routing protocol.
	//
	// example:
	//
	// AuthKey****
	AuthKey *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	// Indicates whether BGP is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	EnableBgp *string `json:"EnableBgp,omitempty" xml:"EnableBgp,omitempty"`
	// The ASN on the Alibaba Cloud side.
	//
	// example:
	//
	// 65531
	LocalAsn *int64 `json:"LocalAsn,omitempty" xml:"LocalAsn,omitempty"`
	// The BGP IP address on the Alibaba Cloud side.
	//
	// example:
	//
	// 169.254.11.2
	LocalBgpIp *string `json:"LocalBgpIp,omitempty" xml:"LocalBgpIp,omitempty"`
	// The autonomous system number (ASN) of the peer.
	//
	// example:
	//
	// 65530
	PeerAsn *int64 `json:"PeerAsn,omitempty" xml:"PeerAsn,omitempty"`
	// The BGP IP address of the peer.
	//
	// example:
	//
	// 169.254.11.1
	PeerBgpIp *string `json:"PeerBgpIp,omitempty" xml:"PeerBgpIp,omitempty"`
	// The negotiation state of the BGP routing protocol. Valid values:
	//
	// 	- **success**: normal
	//
	// 	- **failed**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The BGP CIDR block of the IPsec-VPN connection. The CIDR block falls within 169.254.0.0/16. The subnet mask of the CIDR block must be 30 bits in length.
	//
	// example:
	//
	// 169.254.11.0/30
	TunnelCidr *string `json:"TunnelCidr,omitempty" xml:"TunnelCidr,omitempty"`
}

func (s DescribeVpnConnectionResponseBodyVpnBgpConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnConnectionResponseBodyVpnBgpConfig) GoString() string {
	return s.String()
}

func (s *DescribeVpnConnectionResponseBodyVpnBgpConfig) GetAuthKey() *string {
	return s.AuthKey
}

func (s *DescribeVpnConnectionResponseBodyVpnBgpConfig) GetEnableBgp() *string {
	return s.EnableBgp
}

func (s *DescribeVpnConnectionResponseBodyVpnBgpConfig) GetLocalAsn() *int64 {
	return s.LocalAsn
}

func (s *DescribeVpnConnectionResponseBodyVpnBgpConfig) GetLocalBgpIp() *string {
	return s.LocalBgpIp
}

func (s *DescribeVpnConnectionResponseBodyVpnBgpConfig) GetPeerAsn() *int64 {
	return s.PeerAsn
}

func (s *DescribeVpnConnectionResponseBodyVpnBgpConfig) GetPeerBgpIp() *string {
	return s.PeerBgpIp
}

func (s *DescribeVpnConnectionResponseBodyVpnBgpConfig) GetStatus() *string {
	return s.Status
}

func (s *DescribeVpnConnectionResponseBodyVpnBgpConfig) GetTunnelCidr() *string {
	return s.TunnelCidr
}

func (s *DescribeVpnConnectionResponseBodyVpnBgpConfig) SetAuthKey(v string) *DescribeVpnConnectionResponseBodyVpnBgpConfig {
	s.AuthKey = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyVpnBgpConfig) SetEnableBgp(v string) *DescribeVpnConnectionResponseBodyVpnBgpConfig {
	s.EnableBgp = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyVpnBgpConfig) SetLocalAsn(v int64) *DescribeVpnConnectionResponseBodyVpnBgpConfig {
	s.LocalAsn = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyVpnBgpConfig) SetLocalBgpIp(v string) *DescribeVpnConnectionResponseBodyVpnBgpConfig {
	s.LocalBgpIp = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyVpnBgpConfig) SetPeerAsn(v int64) *DescribeVpnConnectionResponseBodyVpnBgpConfig {
	s.PeerAsn = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyVpnBgpConfig) SetPeerBgpIp(v string) *DescribeVpnConnectionResponseBodyVpnBgpConfig {
	s.PeerBgpIp = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyVpnBgpConfig) SetStatus(v string) *DescribeVpnConnectionResponseBodyVpnBgpConfig {
	s.Status = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyVpnBgpConfig) SetTunnelCidr(v string) *DescribeVpnConnectionResponseBodyVpnBgpConfig {
	s.TunnelCidr = &v
	return s
}

func (s *DescribeVpnConnectionResponseBodyVpnBgpConfig) Validate() error {
	return dara.Validate(s)
}
