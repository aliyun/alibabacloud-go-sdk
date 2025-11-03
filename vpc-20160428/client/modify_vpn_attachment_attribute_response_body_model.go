// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpnAttachmentAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttachInstanceId(v string) *ModifyVpnAttachmentAttributeResponseBody
	GetAttachInstanceId() *string
	SetAttachType(v string) *ModifyVpnAttachmentAttributeResponseBody
	GetAttachType() *string
	SetCreateTime(v int64) *ModifyVpnAttachmentAttributeResponseBody
	GetCreateTime() *int64
	SetCustomerGatewayId(v string) *ModifyVpnAttachmentAttributeResponseBody
	GetCustomerGatewayId() *string
	SetDescription(v string) *ModifyVpnAttachmentAttributeResponseBody
	GetDescription() *string
	SetEffectImmediately(v bool) *ModifyVpnAttachmentAttributeResponseBody
	GetEffectImmediately() *bool
	SetEnableDpd(v bool) *ModifyVpnAttachmentAttributeResponseBody
	GetEnableDpd() *bool
	SetEnableNatTraversal(v bool) *ModifyVpnAttachmentAttributeResponseBody
	GetEnableNatTraversal() *bool
	SetEnableTunnelsBgp(v bool) *ModifyVpnAttachmentAttributeResponseBody
	GetEnableTunnelsBgp() *bool
	SetIkeConfig(v *ModifyVpnAttachmentAttributeResponseBodyIkeConfig) *ModifyVpnAttachmentAttributeResponseBody
	GetIkeConfig() *ModifyVpnAttachmentAttributeResponseBodyIkeConfig
	SetIpsecConfig(v *ModifyVpnAttachmentAttributeResponseBodyIpsecConfig) *ModifyVpnAttachmentAttributeResponseBody
	GetIpsecConfig() *ModifyVpnAttachmentAttributeResponseBodyIpsecConfig
	SetLocalSubnet(v string) *ModifyVpnAttachmentAttributeResponseBody
	GetLocalSubnet() *string
	SetName(v string) *ModifyVpnAttachmentAttributeResponseBody
	GetName() *string
	SetNetworkType(v string) *ModifyVpnAttachmentAttributeResponseBody
	GetNetworkType() *string
	SetRemoteSubnet(v string) *ModifyVpnAttachmentAttributeResponseBody
	GetRemoteSubnet() *string
	SetRequestId(v string) *ModifyVpnAttachmentAttributeResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *ModifyVpnAttachmentAttributeResponseBody
	GetResourceGroupId() *string
	SetSpec(v string) *ModifyVpnAttachmentAttributeResponseBody
	GetSpec() *string
	SetStatus(v string) *ModifyVpnAttachmentAttributeResponseBody
	GetStatus() *string
	SetTunnelOptionsSpecification(v []*ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification) *ModifyVpnAttachmentAttributeResponseBody
	GetTunnelOptionsSpecification() []*ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification
	SetVcoHealthCheck(v *ModifyVpnAttachmentAttributeResponseBodyVcoHealthCheck) *ModifyVpnAttachmentAttributeResponseBody
	GetVcoHealthCheck() *ModifyVpnAttachmentAttributeResponseBodyVcoHealthCheck
	SetVpnBgpConfig(v *ModifyVpnAttachmentAttributeResponseBodyVpnBgpConfig) *ModifyVpnAttachmentAttributeResponseBody
	GetVpnBgpConfig() *ModifyVpnAttachmentAttributeResponseBodyVpnBgpConfig
	SetVpnConnectionId(v string) *ModifyVpnAttachmentAttributeResponseBody
	GetVpnConnectionId() *string
	SetVpnGatewayId(v string) *ModifyVpnAttachmentAttributeResponseBody
	GetVpnGatewayId() *string
}

type ModifyVpnAttachmentAttributeResponseBody struct {
	// The ID of the Cloud Enterprise Network (CEN) instance to which the transit router associated with the IPsec-VPN connection belongs.
	//
	// example:
	//
	// cen-c2r3m3zxkumoqz****
	AttachInstanceId *string `json:"AttachInstanceId,omitempty" xml:"AttachInstanceId,omitempty"`
	// The type of the resource that is associated with the IPsec-VPN connection. Valid values:
	//
	// 	- **CEN**: The IPsec-VPN connection is associated with a transit router.
	//
	// 	- **VPNGW**: The IPsec-VPN connection is associated with a VPN gateway.
	//
	// 	- **NO_ASSOCIATED**: The IPsec-VPN connection is not associated with any resource.
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
	// 1658201810000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the customer gateway associated with the IPsec-VPN connection.
	//
	// This parameter is returned only for single-tunnel IPsec-VPN connections.
	//
	// example:
	//
	// cgw-p0w2jemrcj5u61un8****
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" xml:"CustomerGatewayId,omitempty"`
	// The description of the IPsec-VPN connection.
	//
	// example:
	//
	// desctest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether IPsec negotiations immediately start after the configuration takes effect. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	EffectImmediately *bool `json:"EffectImmediately,omitempty" xml:"EffectImmediately,omitempty"`
	// Indicates whether the DPD feature is enabled for the IPsec-VPN connection.
	//
	// 	- **true**: The feature is enabled.
	//
	// 	- **false**: The feature is disabled.
	//
	// This parameter is returned only for single-tunnel IPsec-VPN connections.
	//
	// example:
	//
	// true
	EnableDpd *bool `json:"EnableDpd,omitempty" xml:"EnableDpd,omitempty"`
	// Specifies whether to enable NAT traversal for the IPsec-VPN connection.
	//
	// 	- **true**: The feature is enabled.
	//
	// 	- **false**: The feature is disabled.
	//
	// This parameter is returned only for single-tunnel IPsec-VPN connections.
	//
	// example:
	//
	// true
	EnableNatTraversal *bool `json:"EnableNatTraversal,omitempty" xml:"EnableNatTraversal,omitempty"`
	// Specifies whether to enable Border Gateway Protocol (BGP) for tunnels.
	//
	// 	- **true**: The feature is enabled.
	//
	// 	- **false**: The feature is disabled.
	//
	// This parameter is returned only by dual-tunnel IPsec-VPN connections.
	//
	// example:
	//
	// false
	EnableTunnelsBgp *bool `json:"EnableTunnelsBgp,omitempty" xml:"EnableTunnelsBgp,omitempty"`
	// The configuration of Phase 1 negotiations.
	//
	// **IkeConfig*	- parameters are returned only for single-tunnel IPsec-VPN connections.
	IkeConfig *ModifyVpnAttachmentAttributeResponseBodyIkeConfig `json:"IkeConfig,omitempty" xml:"IkeConfig,omitempty" type:"Struct"`
	// The configuration of Phase 2 negotiations.
	//
	// **IpsecConfig*	- parameters are returned only for single-tunnel IPsec-VPN connections.
	IpsecConfig *ModifyVpnAttachmentAttributeResponseBodyIpsecConfig `json:"IpsecConfig,omitempty" xml:"IpsecConfig,omitempty" type:"Struct"`
	// The CIDR block on the Alibaba Cloud side that communicates with the on-premises data center is required, such as CIDR blocks of VPCs.
	//
	// example:
	//
	// 10.1.1.0/24,10.1.2.0/24
	LocalSubnet *string `json:"LocalSubnet,omitempty" xml:"LocalSubnet,omitempty"`
	// The name of the IPsec-VPN connection.
	//
	// example:
	//
	// nametest
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
	// The CIDR block of the on-premises data center that communicates with Alibaba Cloud is required.
	//
	// example:
	//
	// 10.1.3.0/24,10.1.4.0/24
	RemoteSubnet *string `json:"RemoteSubnet,omitempty" xml:"RemoteSubnet,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 35822A84-867F-3936-A2E6-A4C4E3ED11C0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the IPsec-VPN connection belongs.
	//
	// You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation to query resource groups.
	//
	// example:
	//
	// rg-acfmzs372yg****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The bandwidth specification of the IPsec-VPN connection.
	//
	// A value of **M*	- in the response indicates **Mbit/s**.
	//
	// example:
	//
	// 1000M
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
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
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tunnel configurations of the IPsec-VPN connection.
	//
	// **TunnelOptionsSpecification*	- parameters are returned only for dual-tunnel IPsec-VPN connections.
	TunnelOptionsSpecification []*ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification `json:"TunnelOptionsSpecification,omitempty" xml:"TunnelOptionsSpecification,omitempty" type:"Repeated"`
	// The health check configurations of the IPsec-VPN connection.
	//
	// **VcoHealthCheck*	- parameters are returned only for single-tunnel IPsec-VPC connections.
	VcoHealthCheck *ModifyVpnAttachmentAttributeResponseBodyVcoHealthCheck `json:"VcoHealthCheck,omitempty" xml:"VcoHealthCheck,omitempty" type:"Struct"`
	// The BGP configurations of the IPsec-VPN connection.
	//
	// **VpnBgpConfig*	- parameters are returned only for single-tunnel IPsec-VPN connections.
	VpnBgpConfig *ModifyVpnAttachmentAttributeResponseBodyVpnBgpConfig `json:"VpnBgpConfig,omitempty" xml:"VpnBgpConfig,omitempty" type:"Struct"`
	// The ID of the IPsec-VPN connection.
	//
	// example:
	//
	// vco-p0w5112fgnl2ihlmf****
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" xml:"VpnConnectionId,omitempty"`
	// The ID of the VPN gateway that is associated with the IPsec-VPN connection.
	//
	// **vpn-not-exist**: The IPsec-VPN connection is not associated with a VPN Gateway.
	//
	// example:
	//
	// vpn-p0wa1c1018pmeb6cu****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
}

func (s ModifyVpnAttachmentAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnAttachmentAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpnAttachmentAttributeResponseBody) GetAttachInstanceId() *string {
	return s.AttachInstanceId
}

func (s *ModifyVpnAttachmentAttributeResponseBody) GetAttachType() *string {
	return s.AttachType
}

func (s *ModifyVpnAttachmentAttributeResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ModifyVpnAttachmentAttributeResponseBody) GetCustomerGatewayId() *string {
	return s.CustomerGatewayId
}

func (s *ModifyVpnAttachmentAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *ModifyVpnAttachmentAttributeResponseBody) GetEffectImmediately() *bool {
	return s.EffectImmediately
}

func (s *ModifyVpnAttachmentAttributeResponseBody) GetEnableDpd() *bool {
	return s.EnableDpd
}

func (s *ModifyVpnAttachmentAttributeResponseBody) GetEnableNatTraversal() *bool {
	return s.EnableNatTraversal
}

func (s *ModifyVpnAttachmentAttributeResponseBody) GetEnableTunnelsBgp() *bool {
	return s.EnableTunnelsBgp
}

func (s *ModifyVpnAttachmentAttributeResponseBody) GetIkeConfig() *ModifyVpnAttachmentAttributeResponseBodyIkeConfig {
	return s.IkeConfig
}

func (s *ModifyVpnAttachmentAttributeResponseBody) GetIpsecConfig() *ModifyVpnAttachmentAttributeResponseBodyIpsecConfig {
	return s.IpsecConfig
}

func (s *ModifyVpnAttachmentAttributeResponseBody) GetLocalSubnet() *string {
	return s.LocalSubnet
}

func (s *ModifyVpnAttachmentAttributeResponseBody) GetName() *string {
	return s.Name
}

func (s *ModifyVpnAttachmentAttributeResponseBody) GetNetworkType() *string {
	return s.NetworkType
}

func (s *ModifyVpnAttachmentAttributeResponseBody) GetRemoteSubnet() *string {
	return s.RemoteSubnet
}

func (s *ModifyVpnAttachmentAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVpnAttachmentAttributeResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyVpnAttachmentAttributeResponseBody) GetSpec() *string {
	return s.Spec
}

func (s *ModifyVpnAttachmentAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ModifyVpnAttachmentAttributeResponseBody) GetTunnelOptionsSpecification() []*ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification {
	return s.TunnelOptionsSpecification
}

func (s *ModifyVpnAttachmentAttributeResponseBody) GetVcoHealthCheck() *ModifyVpnAttachmentAttributeResponseBodyVcoHealthCheck {
	return s.VcoHealthCheck
}

func (s *ModifyVpnAttachmentAttributeResponseBody) GetVpnBgpConfig() *ModifyVpnAttachmentAttributeResponseBodyVpnBgpConfig {
	return s.VpnBgpConfig
}

func (s *ModifyVpnAttachmentAttributeResponseBody) GetVpnConnectionId() *string {
	return s.VpnConnectionId
}

func (s *ModifyVpnAttachmentAttributeResponseBody) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *ModifyVpnAttachmentAttributeResponseBody) SetAttachInstanceId(v string) *ModifyVpnAttachmentAttributeResponseBody {
	s.AttachInstanceId = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBody) SetAttachType(v string) *ModifyVpnAttachmentAttributeResponseBody {
	s.AttachType = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBody) SetCreateTime(v int64) *ModifyVpnAttachmentAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBody) SetCustomerGatewayId(v string) *ModifyVpnAttachmentAttributeResponseBody {
	s.CustomerGatewayId = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBody) SetDescription(v string) *ModifyVpnAttachmentAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBody) SetEffectImmediately(v bool) *ModifyVpnAttachmentAttributeResponseBody {
	s.EffectImmediately = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBody) SetEnableDpd(v bool) *ModifyVpnAttachmentAttributeResponseBody {
	s.EnableDpd = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBody) SetEnableNatTraversal(v bool) *ModifyVpnAttachmentAttributeResponseBody {
	s.EnableNatTraversal = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBody) SetEnableTunnelsBgp(v bool) *ModifyVpnAttachmentAttributeResponseBody {
	s.EnableTunnelsBgp = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBody) SetIkeConfig(v *ModifyVpnAttachmentAttributeResponseBodyIkeConfig) *ModifyVpnAttachmentAttributeResponseBody {
	s.IkeConfig = v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBody) SetIpsecConfig(v *ModifyVpnAttachmentAttributeResponseBodyIpsecConfig) *ModifyVpnAttachmentAttributeResponseBody {
	s.IpsecConfig = v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBody) SetLocalSubnet(v string) *ModifyVpnAttachmentAttributeResponseBody {
	s.LocalSubnet = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBody) SetName(v string) *ModifyVpnAttachmentAttributeResponseBody {
	s.Name = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBody) SetNetworkType(v string) *ModifyVpnAttachmentAttributeResponseBody {
	s.NetworkType = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBody) SetRemoteSubnet(v string) *ModifyVpnAttachmentAttributeResponseBody {
	s.RemoteSubnet = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBody) SetRequestId(v string) *ModifyVpnAttachmentAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBody) SetResourceGroupId(v string) *ModifyVpnAttachmentAttributeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBody) SetSpec(v string) *ModifyVpnAttachmentAttributeResponseBody {
	s.Spec = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBody) SetStatus(v string) *ModifyVpnAttachmentAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBody) SetTunnelOptionsSpecification(v []*ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification) *ModifyVpnAttachmentAttributeResponseBody {
	s.TunnelOptionsSpecification = v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBody) SetVcoHealthCheck(v *ModifyVpnAttachmentAttributeResponseBodyVcoHealthCheck) *ModifyVpnAttachmentAttributeResponseBody {
	s.VcoHealthCheck = v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBody) SetVpnBgpConfig(v *ModifyVpnAttachmentAttributeResponseBodyVpnBgpConfig) *ModifyVpnAttachmentAttributeResponseBody {
	s.VpnBgpConfig = v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBody) SetVpnConnectionId(v string) *ModifyVpnAttachmentAttributeResponseBody {
	s.VpnConnectionId = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBody) SetVpnGatewayId(v string) *ModifyVpnAttachmentAttributeResponseBody {
	s.VpnGatewayId = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBody) Validate() error {
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
	if s.TunnelOptionsSpecification != nil {
		for _, item := range s.TunnelOptionsSpecification {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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

type ModifyVpnAttachmentAttributeResponseBodyIkeConfig struct {
	// The authentication algorithm that is used in Phase 1 negotiations.
	//
	// example:
	//
	// sha1
	IkeAuthAlg *string `json:"IkeAuthAlg,omitempty" xml:"IkeAuthAlg,omitempty"`
	// The encryption algorithm that is used in Phase 1 negotiations.
	//
	// example:
	//
	// aes
	IkeEncAlg *string `json:"IkeEncAlg,omitempty" xml:"IkeEncAlg,omitempty"`
	// The SA lifetime that is determined by Phase 1 negotiations. Unit: seconds.
	//
	// example:
	//
	// 86400
	IkeLifetime *int64 `json:"IkeLifetime,omitempty" xml:"IkeLifetime,omitempty"`
	// The IKE negotiation mode.
	//
	// 	- **main:*	- This mode offers higher security during negotiations.
	//
	// 	- **aggressive**: This mode is faster with a higher success rate.
	//
	// example:
	//
	// main
	IkeMode *string `json:"IkeMode,omitempty" xml:"IkeMode,omitempty"`
	// The DH key exchange algorithm that is used in Phase 1 negotiations.
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
	// Compared with IKEv1, IKEv2 simplifies the SA negotiation process and provides better support for scenarios with multiple CIDR blocks.
	//
	// example:
	//
	// ikev1
	IkeVersion *string `json:"IkeVersion,omitempty" xml:"IkeVersion,omitempty"`
	// The identifier of the IPsec-VPN connection on the Alibaba Cloud side.
	//
	// example:
	//
	// 47.XX.XX.1
	LocalId *string `json:"LocalId,omitempty" xml:"LocalId,omitempty"`
	// Enter a pre-shared key that is used for identity authentication between Alibaba Cloud and the data center.
	//
	// >  The pre-shared key of the IPsec-VPN connection must be the same as the authentication key of the on-premises data center. Otherwise, connections between the on-premises data center and Alibaba Cloud cannot be established.
	//
	// example:
	//
	// 1234***
	Psk *string `json:"Psk,omitempty" xml:"Psk,omitempty"`
	// The identifier of the IPsec-VPN connection on the data center side.
	//
	// example:
	//
	// 47.XX.XX.2
	RemoteId *string `json:"RemoteId,omitempty" xml:"RemoteId,omitempty"`
}

func (s ModifyVpnAttachmentAttributeResponseBodyIkeConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnAttachmentAttributeResponseBodyIkeConfig) GoString() string {
	return s.String()
}

func (s *ModifyVpnAttachmentAttributeResponseBodyIkeConfig) GetIkeAuthAlg() *string {
	return s.IkeAuthAlg
}

func (s *ModifyVpnAttachmentAttributeResponseBodyIkeConfig) GetIkeEncAlg() *string {
	return s.IkeEncAlg
}

func (s *ModifyVpnAttachmentAttributeResponseBodyIkeConfig) GetIkeLifetime() *int64 {
	return s.IkeLifetime
}

func (s *ModifyVpnAttachmentAttributeResponseBodyIkeConfig) GetIkeMode() *string {
	return s.IkeMode
}

func (s *ModifyVpnAttachmentAttributeResponseBodyIkeConfig) GetIkePfs() *string {
	return s.IkePfs
}

func (s *ModifyVpnAttachmentAttributeResponseBodyIkeConfig) GetIkeVersion() *string {
	return s.IkeVersion
}

func (s *ModifyVpnAttachmentAttributeResponseBodyIkeConfig) GetLocalId() *string {
	return s.LocalId
}

func (s *ModifyVpnAttachmentAttributeResponseBodyIkeConfig) GetPsk() *string {
	return s.Psk
}

func (s *ModifyVpnAttachmentAttributeResponseBodyIkeConfig) GetRemoteId() *string {
	return s.RemoteId
}

func (s *ModifyVpnAttachmentAttributeResponseBodyIkeConfig) SetIkeAuthAlg(v string) *ModifyVpnAttachmentAttributeResponseBodyIkeConfig {
	s.IkeAuthAlg = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyIkeConfig) SetIkeEncAlg(v string) *ModifyVpnAttachmentAttributeResponseBodyIkeConfig {
	s.IkeEncAlg = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyIkeConfig) SetIkeLifetime(v int64) *ModifyVpnAttachmentAttributeResponseBodyIkeConfig {
	s.IkeLifetime = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyIkeConfig) SetIkeMode(v string) *ModifyVpnAttachmentAttributeResponseBodyIkeConfig {
	s.IkeMode = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyIkeConfig) SetIkePfs(v string) *ModifyVpnAttachmentAttributeResponseBodyIkeConfig {
	s.IkePfs = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyIkeConfig) SetIkeVersion(v string) *ModifyVpnAttachmentAttributeResponseBodyIkeConfig {
	s.IkeVersion = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyIkeConfig) SetLocalId(v string) *ModifyVpnAttachmentAttributeResponseBodyIkeConfig {
	s.LocalId = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyIkeConfig) SetPsk(v string) *ModifyVpnAttachmentAttributeResponseBodyIkeConfig {
	s.Psk = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyIkeConfig) SetRemoteId(v string) *ModifyVpnAttachmentAttributeResponseBodyIkeConfig {
	s.RemoteId = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyIkeConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyVpnAttachmentAttributeResponseBodyIpsecConfig struct {
	// The authentication algorithm that is used in Phase 2 negotiations.
	//
	// example:
	//
	// md5
	IpsecAuthAlg *string `json:"IpsecAuthAlg,omitempty" xml:"IpsecAuthAlg,omitempty"`
	// The encryption algorithm that is used in Phase 2 negotiations.
	//
	// example:
	//
	// aes
	IpsecEncAlg *string `json:"IpsecEncAlg,omitempty" xml:"IpsecEncAlg,omitempty"`
	// The SA lifetime that is determined by Phase 2 negotiations. Unit: seconds.
	//
	// example:
	//
	// 86400
	IpsecLifetime *int64 `json:"IpsecLifetime,omitempty" xml:"IpsecLifetime,omitempty"`
	// The DH key exchange algorithm that is used in Phase 2 negotiations.
	//
	// example:
	//
	// group2
	IpsecPfs *string `json:"IpsecPfs,omitempty" xml:"IpsecPfs,omitempty"`
}

func (s ModifyVpnAttachmentAttributeResponseBodyIpsecConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnAttachmentAttributeResponseBodyIpsecConfig) GoString() string {
	return s.String()
}

func (s *ModifyVpnAttachmentAttributeResponseBodyIpsecConfig) GetIpsecAuthAlg() *string {
	return s.IpsecAuthAlg
}

func (s *ModifyVpnAttachmentAttributeResponseBodyIpsecConfig) GetIpsecEncAlg() *string {
	return s.IpsecEncAlg
}

func (s *ModifyVpnAttachmentAttributeResponseBodyIpsecConfig) GetIpsecLifetime() *int64 {
	return s.IpsecLifetime
}

func (s *ModifyVpnAttachmentAttributeResponseBodyIpsecConfig) GetIpsecPfs() *string {
	return s.IpsecPfs
}

func (s *ModifyVpnAttachmentAttributeResponseBodyIpsecConfig) SetIpsecAuthAlg(v string) *ModifyVpnAttachmentAttributeResponseBodyIpsecConfig {
	s.IpsecAuthAlg = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyIpsecConfig) SetIpsecEncAlg(v string) *ModifyVpnAttachmentAttributeResponseBodyIpsecConfig {
	s.IpsecEncAlg = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyIpsecConfig) SetIpsecLifetime(v int64) *ModifyVpnAttachmentAttributeResponseBodyIpsecConfig {
	s.IpsecLifetime = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyIpsecConfig) SetIpsecPfs(v string) *ModifyVpnAttachmentAttributeResponseBodyIpsecConfig {
	s.IpsecPfs = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyIpsecConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification struct {
	// The ID of the customer gateway that is associated with the tunnel.
	//
	// example:
	//
	// cgw-p0w2jemrcj5u61un8****
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" xml:"CustomerGatewayId,omitempty"`
	// Whether the DPD feature is enabled for the tunnel.
	//
	// 	- **true**: The feature is enabled.
	//
	// 	- **false**: The feature is disabled.
	//
	// example:
	//
	// true
	EnableDpd *bool `json:"EnableDpd,omitempty" xml:"EnableDpd,omitempty"`
	// Indicates whether traversal feature is enabled for the tunnel. Valid values:
	//
	// 	- **true**: The feature is enabled.
	//
	// 	- **false**: The feature is disabled.
	//
	// example:
	//
	// true
	EnableNatTraversal *bool `json:"EnableNatTraversal,omitempty" xml:"EnableNatTraversal,omitempty"`
	// The IP address on the Alibaba Cloud side.
	//
	// example:
	//
	// 47.XX.XX.66
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
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
	// The status of the tunnel. Valid values:
	//
	// 	- **active**: The tunnel is active.
	//
	// 	- **updating**: The tunnel is being updated.
	//
	// 	- **deleting:*	- The tunnel is being deleted.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// BGP configuration.
	TunnelBgpConfig *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelBgpConfig `json:"TunnelBgpConfig,omitempty" xml:"TunnelBgpConfig,omitempty" type:"Struct"`
	// The tunnel ID.
	//
	// example:
	//
	// tun-0jod7plwf2a0o9lvu****
	TunnelId *string `json:"TunnelId,omitempty" xml:"TunnelId,omitempty"`
	// The configurations of Phase 1 negotiations.
	TunnelIkeConfig *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIkeConfig `json:"TunnelIkeConfig,omitempty" xml:"TunnelIkeConfig,omitempty" type:"Struct"`
	// The order in which the tunnel was created.
	//
	// 	- **1**: Tunnel 1.
	//
	// 	- **2**: Tunnel 2.
	//
	// example:
	//
	// 1
	TunnelIndex *int32 `json:"TunnelIndex,omitempty" xml:"TunnelIndex,omitempty"`
	// The configurations of Phase 2 negotiations.
	TunnelIpsecConfig *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIpsecConfig `json:"TunnelIpsecConfig,omitempty" xml:"TunnelIpsecConfig,omitempty" type:"Struct"`
}

func (s ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification) GoString() string {
	return s.String()
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification) GetCustomerGatewayId() *string {
	return s.CustomerGatewayId
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification) GetEnableDpd() *bool {
	return s.EnableDpd
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification) GetEnableNatTraversal() *bool {
	return s.EnableNatTraversal
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification) GetInternetIp() *string {
	return s.InternetIp
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification) GetRole() *string {
	return s.Role
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification) GetState() *string {
	return s.State
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification) GetTunnelBgpConfig() *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelBgpConfig {
	return s.TunnelBgpConfig
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification) GetTunnelId() *string {
	return s.TunnelId
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification) GetTunnelIkeConfig() *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIkeConfig {
	return s.TunnelIkeConfig
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification) GetTunnelIndex() *int32 {
	return s.TunnelIndex
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification) GetTunnelIpsecConfig() *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIpsecConfig {
	return s.TunnelIpsecConfig
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification) SetCustomerGatewayId(v string) *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification {
	s.CustomerGatewayId = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification) SetEnableDpd(v bool) *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification {
	s.EnableDpd = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification) SetEnableNatTraversal(v bool) *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification {
	s.EnableNatTraversal = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification) SetInternetIp(v string) *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification {
	s.InternetIp = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification) SetRole(v string) *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification {
	s.Role = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification) SetState(v string) *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification {
	s.State = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification) SetTunnelBgpConfig(v *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelBgpConfig) *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification {
	s.TunnelBgpConfig = v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification) SetTunnelId(v string) *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification {
	s.TunnelId = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification) SetTunnelIkeConfig(v *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIkeConfig) *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification {
	s.TunnelIkeConfig = v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification) SetTunnelIndex(v int32) *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification {
	s.TunnelIndex = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification) SetTunnelIpsecConfig(v *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIpsecConfig) *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification {
	s.TunnelIpsecConfig = v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecification) Validate() error {
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

type ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelBgpConfig struct {
	// The ASN on the Alibaba Cloud side.
	//
	// example:
	//
	// 65530
	LocalAsn *int64 `json:"LocalAsn,omitempty" xml:"LocalAsn,omitempty"`
	// The BGP IP address of the tunnel on the Alibaba Cloud side.
	//
	// example:
	//
	// 169.254.10.1
	LocalBgpIp *string `json:"LocalBgpIp,omitempty" xml:"LocalBgpIp,omitempty"`
	// The ASN of the tunnel peer.
	//
	// example:
	//
	// 65531
	PeerAsn *int64 `json:"PeerAsn,omitempty" xml:"PeerAsn,omitempty"`
	// The BGP IP address of the tunnel peer.
	//
	// example:
	//
	// 169.254.10.2
	PeerBgpIp *string `json:"PeerBgpIp,omitempty" xml:"PeerBgpIp,omitempty"`
	// The BGP CIDR block of the tunnel.
	//
	// example:
	//
	// 169.254.10.0/30
	TunnelCidr *string `json:"TunnelCidr,omitempty" xml:"TunnelCidr,omitempty"`
}

func (s ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelBgpConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelBgpConfig) GoString() string {
	return s.String()
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelBgpConfig) GetLocalAsn() *int64 {
	return s.LocalAsn
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelBgpConfig) GetLocalBgpIp() *string {
	return s.LocalBgpIp
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelBgpConfig) GetPeerAsn() *int64 {
	return s.PeerAsn
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelBgpConfig) GetPeerBgpIp() *string {
	return s.PeerBgpIp
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelBgpConfig) GetTunnelCidr() *string {
	return s.TunnelCidr
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelBgpConfig) SetLocalAsn(v int64) *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelBgpConfig {
	s.LocalAsn = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelBgpConfig) SetLocalBgpIp(v string) *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelBgpConfig {
	s.LocalBgpIp = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelBgpConfig) SetPeerAsn(v int64) *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelBgpConfig {
	s.PeerAsn = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelBgpConfig) SetPeerBgpIp(v string) *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelBgpConfig {
	s.PeerBgpIp = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelBgpConfig) SetTunnelCidr(v string) *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelBgpConfig {
	s.TunnelCidr = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelBgpConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIkeConfig struct {
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
	// The negotiation mode of IKE. Valid values:
	//
	// 	- **main:*	- This mode offers higher security during negotiations.
	//
	// 	- **aggressive**: This mode is faster with a higher success rate.
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
	// example:
	//
	// ikev2
	IkeVersion *string `json:"IkeVersion,omitempty" xml:"IkeVersion,omitempty"`
	// The identifier of the tunnel on the Alibaba Cloud side.
	//
	// example:
	//
	// 47.XX.XX.1
	LocalId *string `json:"LocalId,omitempty" xml:"LocalId,omitempty"`
	// The pre-shared key.
	//
	// example:
	//
	// 123456****
	Psk *string `json:"Psk,omitempty" xml:"Psk,omitempty"`
	// The peer identifier.
	//
	// example:
	//
	// 47.XX.XX.2
	RemoteId *string `json:"RemoteId,omitempty" xml:"RemoteId,omitempty"`
}

func (s ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIkeConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIkeConfig) GoString() string {
	return s.String()
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIkeConfig) GetIkeAuthAlg() *string {
	return s.IkeAuthAlg
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIkeConfig) GetIkeEncAlg() *string {
	return s.IkeEncAlg
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIkeConfig) GetIkeLifetime() *int64 {
	return s.IkeLifetime
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIkeConfig) GetIkeMode() *string {
	return s.IkeMode
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIkeConfig) GetIkePfs() *string {
	return s.IkePfs
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIkeConfig) GetIkeVersion() *string {
	return s.IkeVersion
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIkeConfig) GetLocalId() *string {
	return s.LocalId
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIkeConfig) GetPsk() *string {
	return s.Psk
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIkeConfig) GetRemoteId() *string {
	return s.RemoteId
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIkeConfig) SetIkeAuthAlg(v string) *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkeAuthAlg = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIkeConfig) SetIkeEncAlg(v string) *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkeEncAlg = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIkeConfig) SetIkeLifetime(v int64) *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkeLifetime = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIkeConfig) SetIkeMode(v string) *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkeMode = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIkeConfig) SetIkePfs(v string) *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkePfs = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIkeConfig) SetIkeVersion(v string) *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIkeConfig {
	s.IkeVersion = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIkeConfig) SetLocalId(v string) *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIkeConfig {
	s.LocalId = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIkeConfig) SetPsk(v string) *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIkeConfig {
	s.Psk = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIkeConfig) SetRemoteId(v string) *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIkeConfig {
	s.RemoteId = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIkeConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIpsecConfig struct {
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

func (s ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIpsecConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIpsecConfig) GoString() string {
	return s.String()
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIpsecConfig) GetIpsecAuthAlg() *string {
	return s.IpsecAuthAlg
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIpsecConfig) GetIpsecEncAlg() *string {
	return s.IpsecEncAlg
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIpsecConfig) GetIpsecLifetime() *int64 {
	return s.IpsecLifetime
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIpsecConfig) GetIpsecPfs() *string {
	return s.IpsecPfs
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIpsecConfig) SetIpsecAuthAlg(v string) *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIpsecConfig {
	s.IpsecAuthAlg = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIpsecConfig) SetIpsecEncAlg(v string) *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIpsecConfig {
	s.IpsecEncAlg = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIpsecConfig) SetIpsecLifetime(v int64) *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIpsecConfig {
	s.IpsecLifetime = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIpsecConfig) SetIpsecPfs(v string) *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIpsecConfig {
	s.IpsecPfs = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyTunnelOptionsSpecificationTunnelIpsecConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyVpnAttachmentAttributeResponseBodyVcoHealthCheck struct {
	// The destination IP address that is used for health checks.
	//
	// example:
	//
	// 192.168.1.1
	Dip *string `json:"Dip,omitempty" xml:"Dip,omitempty"`
	// Indicates whether the health check feature is enabled for the IPsec-VPN connection. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The interval between two consecutive health check retries. Unit: seconds.
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
	// The source IP address that is used for health checks.
	//
	// example:
	//
	// 10.1.1.1
	Sip *string `json:"Sip,omitempty" xml:"Sip,omitempty"`
}

func (s ModifyVpnAttachmentAttributeResponseBodyVcoHealthCheck) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnAttachmentAttributeResponseBodyVcoHealthCheck) GoString() string {
	return s.String()
}

func (s *ModifyVpnAttachmentAttributeResponseBodyVcoHealthCheck) GetDip() *string {
	return s.Dip
}

func (s *ModifyVpnAttachmentAttributeResponseBodyVcoHealthCheck) GetEnable() *string {
	return s.Enable
}

func (s *ModifyVpnAttachmentAttributeResponseBodyVcoHealthCheck) GetInterval() *int32 {
	return s.Interval
}

func (s *ModifyVpnAttachmentAttributeResponseBodyVcoHealthCheck) GetPolicy() *string {
	return s.Policy
}

func (s *ModifyVpnAttachmentAttributeResponseBodyVcoHealthCheck) GetRetry() *int32 {
	return s.Retry
}

func (s *ModifyVpnAttachmentAttributeResponseBodyVcoHealthCheck) GetSip() *string {
	return s.Sip
}

func (s *ModifyVpnAttachmentAttributeResponseBodyVcoHealthCheck) SetDip(v string) *ModifyVpnAttachmentAttributeResponseBodyVcoHealthCheck {
	s.Dip = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyVcoHealthCheck) SetEnable(v string) *ModifyVpnAttachmentAttributeResponseBodyVcoHealthCheck {
	s.Enable = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyVcoHealthCheck) SetInterval(v int32) *ModifyVpnAttachmentAttributeResponseBodyVcoHealthCheck {
	s.Interval = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyVcoHealthCheck) SetPolicy(v string) *ModifyVpnAttachmentAttributeResponseBodyVcoHealthCheck {
	s.Policy = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyVcoHealthCheck) SetRetry(v int32) *ModifyVpnAttachmentAttributeResponseBodyVcoHealthCheck {
	s.Retry = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyVcoHealthCheck) SetSip(v string) *ModifyVpnAttachmentAttributeResponseBodyVcoHealthCheck {
	s.Sip = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyVcoHealthCheck) Validate() error {
	return dara.Validate(s)
}

type ModifyVpnAttachmentAttributeResponseBodyVpnBgpConfig struct {
	// Indicates whether BGP is enabled for the IPsec-VPN connection. Valid values:
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
	// 45104
	LocalAsn *int64 `json:"LocalAsn,omitempty" xml:"LocalAsn,omitempty"`
	// The BGP IP address on the Alibaba Cloud side.
	//
	// example:
	//
	// 169.254.11.1
	LocalBgpIp *string `json:"LocalBgpIp,omitempty" xml:"LocalBgpIp,omitempty"`
	// The ASN on the data center side.
	//
	// example:
	//
	// 65535
	PeerAsn *int64 `json:"PeerAsn,omitempty" xml:"PeerAsn,omitempty"`
	// The BGP IP address on the data center side.
	//
	// example:
	//
	// 169.254.11.2
	PeerBgpIp *string `json:"PeerBgpIp,omitempty" xml:"PeerBgpIp,omitempty"`
	// The negotiation state of BGP. Valid values:
	//
	// 	- **success**: normal
	//
	// 	- **false**: abnormal
	//
	// example:
	//
	// false
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The CIDR block of the IPsec tunnel.
	//
	// example:
	//
	// 169.254.11.0/30
	TunnelCidr *string `json:"TunnelCidr,omitempty" xml:"TunnelCidr,omitempty"`
}

func (s ModifyVpnAttachmentAttributeResponseBodyVpnBgpConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnAttachmentAttributeResponseBodyVpnBgpConfig) GoString() string {
	return s.String()
}

func (s *ModifyVpnAttachmentAttributeResponseBodyVpnBgpConfig) GetEnableBgp() *string {
	return s.EnableBgp
}

func (s *ModifyVpnAttachmentAttributeResponseBodyVpnBgpConfig) GetLocalAsn() *int64 {
	return s.LocalAsn
}

func (s *ModifyVpnAttachmentAttributeResponseBodyVpnBgpConfig) GetLocalBgpIp() *string {
	return s.LocalBgpIp
}

func (s *ModifyVpnAttachmentAttributeResponseBodyVpnBgpConfig) GetPeerAsn() *int64 {
	return s.PeerAsn
}

func (s *ModifyVpnAttachmentAttributeResponseBodyVpnBgpConfig) GetPeerBgpIp() *string {
	return s.PeerBgpIp
}

func (s *ModifyVpnAttachmentAttributeResponseBodyVpnBgpConfig) GetStatus() *string {
	return s.Status
}

func (s *ModifyVpnAttachmentAttributeResponseBodyVpnBgpConfig) GetTunnelCidr() *string {
	return s.TunnelCidr
}

func (s *ModifyVpnAttachmentAttributeResponseBodyVpnBgpConfig) SetEnableBgp(v string) *ModifyVpnAttachmentAttributeResponseBodyVpnBgpConfig {
	s.EnableBgp = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyVpnBgpConfig) SetLocalAsn(v int64) *ModifyVpnAttachmentAttributeResponseBodyVpnBgpConfig {
	s.LocalAsn = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyVpnBgpConfig) SetLocalBgpIp(v string) *ModifyVpnAttachmentAttributeResponseBodyVpnBgpConfig {
	s.LocalBgpIp = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyVpnBgpConfig) SetPeerAsn(v int64) *ModifyVpnAttachmentAttributeResponseBodyVpnBgpConfig {
	s.PeerAsn = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyVpnBgpConfig) SetPeerBgpIp(v string) *ModifyVpnAttachmentAttributeResponseBodyVpnBgpConfig {
	s.PeerBgpIp = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyVpnBgpConfig) SetStatus(v string) *ModifyVpnAttachmentAttributeResponseBodyVpnBgpConfig {
	s.Status = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyVpnBgpConfig) SetTunnelCidr(v string) *ModifyVpnAttachmentAttributeResponseBodyVpnBgpConfig {
	s.TunnelCidr = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponseBodyVpnBgpConfig) Validate() error {
	return dara.Validate(s)
}
