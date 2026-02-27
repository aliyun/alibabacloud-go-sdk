// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpnConnectionAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v int64) *ModifyVpnConnectionAttributeResponseBody
	GetCreateTime() *int64
	SetCustomerGatewayId(v string) *ModifyVpnConnectionAttributeResponseBody
	GetCustomerGatewayId() *string
	SetDescription(v string) *ModifyVpnConnectionAttributeResponseBody
	GetDescription() *string
	SetEffectImmediately(v bool) *ModifyVpnConnectionAttributeResponseBody
	GetEffectImmediately() *bool
	SetEnableDpd(v bool) *ModifyVpnConnectionAttributeResponseBody
	GetEnableDpd() *bool
	SetEnableNatTraversal(v bool) *ModifyVpnConnectionAttributeResponseBody
	GetEnableNatTraversal() *bool
	SetEnableTunnelsBgp(v bool) *ModifyVpnConnectionAttributeResponseBody
	GetEnableTunnelsBgp() *bool
	SetIkeConfig(v *ModifyVpnConnectionAttributeResponseBodyIkeConfig) *ModifyVpnConnectionAttributeResponseBody
	GetIkeConfig() *ModifyVpnConnectionAttributeResponseBodyIkeConfig
	SetIpsecConfig(v *ModifyVpnConnectionAttributeResponseBodyIpsecConfig) *ModifyVpnConnectionAttributeResponseBody
	GetIpsecConfig() *ModifyVpnConnectionAttributeResponseBodyIpsecConfig
	SetLocalSubnet(v string) *ModifyVpnConnectionAttributeResponseBody
	GetLocalSubnet() *string
	SetName(v string) *ModifyVpnConnectionAttributeResponseBody
	GetName() *string
	SetRemoteSubnet(v string) *ModifyVpnConnectionAttributeResponseBody
	GetRemoteSubnet() *string
	SetRequestId(v string) *ModifyVpnConnectionAttributeResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *ModifyVpnConnectionAttributeResponseBody
	GetResourceGroupId() *string
	SetTunnelOptionsSpecification(v *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecification) *ModifyVpnConnectionAttributeResponseBody
	GetTunnelOptionsSpecification() *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecification
	SetVcoHealthCheck(v *ModifyVpnConnectionAttributeResponseBodyVcoHealthCheck) *ModifyVpnConnectionAttributeResponseBody
	GetVcoHealthCheck() *ModifyVpnConnectionAttributeResponseBodyVcoHealthCheck
	SetVpnBgpConfig(v *ModifyVpnConnectionAttributeResponseBodyVpnBgpConfig) *ModifyVpnConnectionAttributeResponseBody
	GetVpnBgpConfig() *ModifyVpnConnectionAttributeResponseBodyVpnBgpConfig
	SetVpnConnectionId(v string) *ModifyVpnConnectionAttributeResponseBody
	GetVpnConnectionId() *string
	SetVpnGatewayId(v string) *ModifyVpnConnectionAttributeResponseBody
	GetVpnGatewayId() *string
}

type ModifyVpnConnectionAttributeResponseBody struct {
	// The timestamp generated when the IPsec-VPN connection was established. Unit: milliseconds.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1492753817000
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
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether IPsec negotiations immediately start after the configuration takes effect. Valid values:
	//
	// 	- **true**: IPsec negotiations immediately start after the configuration takes effect.
	//
	// 	- **false**: IPsec negotiations start when inbound traffic is detected.
	//
	// example:
	//
	// false
	EffectImmediately *bool `json:"EffectImmediately,omitempty" xml:"EffectImmediately,omitempty"`
	// Indicates whether the DPD feature is enabled for the IPsec-VPN connection. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// This parameter is returned only for single-tunnel IPsec-VPN connections.
	//
	// example:
	//
	// true
	EnableDpd *bool `json:"EnableDpd,omitempty" xml:"EnableDpd,omitempty"`
	// Indicates whether NAT traversal is enabled for the IPsec-VPN connection. Valid values: Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// This parameter is returned only for single-tunnel IPsec-VPN connections.
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
	// This parameter is returned only by dual-tunnel IPsec-VPN connections.
	//
	// example:
	//
	// true
	EnableTunnelsBgp *bool `json:"EnableTunnelsBgp,omitempty" xml:"EnableTunnelsBgp,omitempty"`
	// The configuration of Phase 1 negotiations.
	//
	// **IkeConfig*	- parameters are returned only for single-tunnel IPsec-VPN connections.
	IkeConfig *ModifyVpnConnectionAttributeResponseBodyIkeConfig `json:"IkeConfig,omitempty" xml:"IkeConfig,omitempty" type:"Struct"`
	// The configuration of Phase 2 negotiations.
	//
	// **IpsecConfig*	- parameters are returned only for single-tunnel IPsec-VPN connections.
	IpsecConfig *ModifyVpnConnectionAttributeResponseBodyIpsecConfig `json:"IpsecConfig,omitempty" xml:"IpsecConfig,omitempty" type:"Struct"`
	// The CIDR block on the VPC side.
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
	// The CIDR block on the data center side.
	//
	// example:
	//
	// 10.2.1.0/24,10.2.2.0/24
	RemoteSubnet *string `json:"RemoteSubnet,omitempty" xml:"RemoteSubnet,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7DB79D0C-5F27-4AB5-995B-79BE55102F90
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the IPsec-VPN connection belongs.
	//
	// The IPsec-VPN connection and the VPN gateway associated with the IPsec-VPN connection belong to the same resource group. You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation to query resource groups.
	//
	// example:
	//
	// rg-acfmzs372yg****
	ResourceGroupId            *string                                                             `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	TunnelOptionsSpecification *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecification `json:"TunnelOptionsSpecification,omitempty" xml:"TunnelOptionsSpecification,omitempty" type:"Struct"`
	// The health check configuration.
	//
	// **VcoHealthCheck*	- parameters are returned only for single-tunnel IPsec-VPN connections.
	VcoHealthCheck *ModifyVpnConnectionAttributeResponseBodyVcoHealthCheck `json:"VcoHealthCheck,omitempty" xml:"VcoHealthCheck,omitempty" type:"Struct"`
	// The BGP configuration.
	//
	// **VpnBgpConfig*	- parameters are returned only for single-tunnel IPsec-VPN connections.
	VpnBgpConfig *ModifyVpnConnectionAttributeResponseBodyVpnBgpConfig `json:"VpnBgpConfig,omitempty" xml:"VpnBgpConfig,omitempty" type:"Struct"`
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
}

func (s ModifyVpnConnectionAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnConnectionAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpnConnectionAttributeResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ModifyVpnConnectionAttributeResponseBody) GetCustomerGatewayId() *string {
	return s.CustomerGatewayId
}

func (s *ModifyVpnConnectionAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *ModifyVpnConnectionAttributeResponseBody) GetEffectImmediately() *bool {
	return s.EffectImmediately
}

func (s *ModifyVpnConnectionAttributeResponseBody) GetEnableDpd() *bool {
	return s.EnableDpd
}

func (s *ModifyVpnConnectionAttributeResponseBody) GetEnableNatTraversal() *bool {
	return s.EnableNatTraversal
}

func (s *ModifyVpnConnectionAttributeResponseBody) GetEnableTunnelsBgp() *bool {
	return s.EnableTunnelsBgp
}

func (s *ModifyVpnConnectionAttributeResponseBody) GetIkeConfig() *ModifyVpnConnectionAttributeResponseBodyIkeConfig {
	return s.IkeConfig
}

func (s *ModifyVpnConnectionAttributeResponseBody) GetIpsecConfig() *ModifyVpnConnectionAttributeResponseBodyIpsecConfig {
	return s.IpsecConfig
}

func (s *ModifyVpnConnectionAttributeResponseBody) GetLocalSubnet() *string {
	return s.LocalSubnet
}

func (s *ModifyVpnConnectionAttributeResponseBody) GetName() *string {
	return s.Name
}

func (s *ModifyVpnConnectionAttributeResponseBody) GetRemoteSubnet() *string {
	return s.RemoteSubnet
}

func (s *ModifyVpnConnectionAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVpnConnectionAttributeResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyVpnConnectionAttributeResponseBody) GetTunnelOptionsSpecification() *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecification {
	return s.TunnelOptionsSpecification
}

func (s *ModifyVpnConnectionAttributeResponseBody) GetVcoHealthCheck() *ModifyVpnConnectionAttributeResponseBodyVcoHealthCheck {
	return s.VcoHealthCheck
}

func (s *ModifyVpnConnectionAttributeResponseBody) GetVpnBgpConfig() *ModifyVpnConnectionAttributeResponseBodyVpnBgpConfig {
	return s.VpnBgpConfig
}

func (s *ModifyVpnConnectionAttributeResponseBody) GetVpnConnectionId() *string {
	return s.VpnConnectionId
}

func (s *ModifyVpnConnectionAttributeResponseBody) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *ModifyVpnConnectionAttributeResponseBody) SetCreateTime(v int64) *ModifyVpnConnectionAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBody) SetCustomerGatewayId(v string) *ModifyVpnConnectionAttributeResponseBody {
	s.CustomerGatewayId = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBody) SetDescription(v string) *ModifyVpnConnectionAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBody) SetEffectImmediately(v bool) *ModifyVpnConnectionAttributeResponseBody {
	s.EffectImmediately = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBody) SetEnableDpd(v bool) *ModifyVpnConnectionAttributeResponseBody {
	s.EnableDpd = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBody) SetEnableNatTraversal(v bool) *ModifyVpnConnectionAttributeResponseBody {
	s.EnableNatTraversal = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBody) SetEnableTunnelsBgp(v bool) *ModifyVpnConnectionAttributeResponseBody {
	s.EnableTunnelsBgp = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBody) SetIkeConfig(v *ModifyVpnConnectionAttributeResponseBodyIkeConfig) *ModifyVpnConnectionAttributeResponseBody {
	s.IkeConfig = v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBody) SetIpsecConfig(v *ModifyVpnConnectionAttributeResponseBodyIpsecConfig) *ModifyVpnConnectionAttributeResponseBody {
	s.IpsecConfig = v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBody) SetLocalSubnet(v string) *ModifyVpnConnectionAttributeResponseBody {
	s.LocalSubnet = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBody) SetName(v string) *ModifyVpnConnectionAttributeResponseBody {
	s.Name = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBody) SetRemoteSubnet(v string) *ModifyVpnConnectionAttributeResponseBody {
	s.RemoteSubnet = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBody) SetRequestId(v string) *ModifyVpnConnectionAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBody) SetResourceGroupId(v string) *ModifyVpnConnectionAttributeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBody) SetTunnelOptionsSpecification(v *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecification) *ModifyVpnConnectionAttributeResponseBody {
	s.TunnelOptionsSpecification = v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBody) SetVcoHealthCheck(v *ModifyVpnConnectionAttributeResponseBodyVcoHealthCheck) *ModifyVpnConnectionAttributeResponseBody {
	s.VcoHealthCheck = v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBody) SetVpnBgpConfig(v *ModifyVpnConnectionAttributeResponseBodyVpnBgpConfig) *ModifyVpnConnectionAttributeResponseBody {
	s.VpnBgpConfig = v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBody) SetVpnConnectionId(v string) *ModifyVpnConnectionAttributeResponseBody {
	s.VpnConnectionId = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBody) SetVpnGatewayId(v string) *ModifyVpnConnectionAttributeResponseBody {
	s.VpnGatewayId = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBody) Validate() error {
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

type ModifyVpnConnectionAttributeResponseBodyIkeConfig struct {
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
	// The DH group in the IKE phase.
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
	// The identifier on the VPC side. The default value is the IP address of the VPN gateway. The value can be an FQDN or an IP address.
	//
	// example:
	//
	// 116.64.XX.XX
	LocalId *string `json:"LocalId,omitempty" xml:"LocalId,omitempty"`
	// The pre-shared key.
	//
	// example:
	//
	// pgw6dy7d1i8i****
	Psk *string `json:"Psk,omitempty" xml:"Psk,omitempty"`
	// The identifier on the data center side. The default value is the IP address of the customer gateway. The value can be a FQDN or an IP address.
	//
	// example:
	//
	// 139.18.XX.XX
	RemoteId *string `json:"RemoteId,omitempty" xml:"RemoteId,omitempty"`
}

func (s ModifyVpnConnectionAttributeResponseBodyIkeConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnConnectionAttributeResponseBodyIkeConfig) GoString() string {
	return s.String()
}

func (s *ModifyVpnConnectionAttributeResponseBodyIkeConfig) GetIkeAuthAlg() *string {
	return s.IkeAuthAlg
}

func (s *ModifyVpnConnectionAttributeResponseBodyIkeConfig) GetIkeEncAlg() *string {
	return s.IkeEncAlg
}

func (s *ModifyVpnConnectionAttributeResponseBodyIkeConfig) GetIkeLifetime() *int64 {
	return s.IkeLifetime
}

func (s *ModifyVpnConnectionAttributeResponseBodyIkeConfig) GetIkeMode() *string {
	return s.IkeMode
}

func (s *ModifyVpnConnectionAttributeResponseBodyIkeConfig) GetIkePfs() *string {
	return s.IkePfs
}

func (s *ModifyVpnConnectionAttributeResponseBodyIkeConfig) GetIkeVersion() *string {
	return s.IkeVersion
}

func (s *ModifyVpnConnectionAttributeResponseBodyIkeConfig) GetLocalId() *string {
	return s.LocalId
}

func (s *ModifyVpnConnectionAttributeResponseBodyIkeConfig) GetPsk() *string {
	return s.Psk
}

func (s *ModifyVpnConnectionAttributeResponseBodyIkeConfig) GetRemoteId() *string {
	return s.RemoteId
}

func (s *ModifyVpnConnectionAttributeResponseBodyIkeConfig) SetIkeAuthAlg(v string) *ModifyVpnConnectionAttributeResponseBodyIkeConfig {
	s.IkeAuthAlg = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyIkeConfig) SetIkeEncAlg(v string) *ModifyVpnConnectionAttributeResponseBodyIkeConfig {
	s.IkeEncAlg = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyIkeConfig) SetIkeLifetime(v int64) *ModifyVpnConnectionAttributeResponseBodyIkeConfig {
	s.IkeLifetime = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyIkeConfig) SetIkeMode(v string) *ModifyVpnConnectionAttributeResponseBodyIkeConfig {
	s.IkeMode = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyIkeConfig) SetIkePfs(v string) *ModifyVpnConnectionAttributeResponseBodyIkeConfig {
	s.IkePfs = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyIkeConfig) SetIkeVersion(v string) *ModifyVpnConnectionAttributeResponseBodyIkeConfig {
	s.IkeVersion = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyIkeConfig) SetLocalId(v string) *ModifyVpnConnectionAttributeResponseBodyIkeConfig {
	s.LocalId = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyIkeConfig) SetPsk(v string) *ModifyVpnConnectionAttributeResponseBodyIkeConfig {
	s.Psk = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyIkeConfig) SetRemoteId(v string) *ModifyVpnConnectionAttributeResponseBodyIkeConfig {
	s.RemoteId = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyIkeConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyVpnConnectionAttributeResponseBodyIpsecConfig struct {
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

func (s ModifyVpnConnectionAttributeResponseBodyIpsecConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnConnectionAttributeResponseBodyIpsecConfig) GoString() string {
	return s.String()
}

func (s *ModifyVpnConnectionAttributeResponseBodyIpsecConfig) GetIpsecAuthAlg() *string {
	return s.IpsecAuthAlg
}

func (s *ModifyVpnConnectionAttributeResponseBodyIpsecConfig) GetIpsecEncAlg() *string {
	return s.IpsecEncAlg
}

func (s *ModifyVpnConnectionAttributeResponseBodyIpsecConfig) GetIpsecLifetime() *int64 {
	return s.IpsecLifetime
}

func (s *ModifyVpnConnectionAttributeResponseBodyIpsecConfig) GetIpsecPfs() *string {
	return s.IpsecPfs
}

func (s *ModifyVpnConnectionAttributeResponseBodyIpsecConfig) SetIpsecAuthAlg(v string) *ModifyVpnConnectionAttributeResponseBodyIpsecConfig {
	s.IpsecAuthAlg = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyIpsecConfig) SetIpsecEncAlg(v string) *ModifyVpnConnectionAttributeResponseBodyIpsecConfig {
	s.IpsecEncAlg = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyIpsecConfig) SetIpsecLifetime(v int64) *ModifyVpnConnectionAttributeResponseBodyIpsecConfig {
	s.IpsecLifetime = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyIpsecConfig) SetIpsecPfs(v string) *ModifyVpnConnectionAttributeResponseBodyIpsecConfig {
	s.IpsecPfs = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyIpsecConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecification struct {
	TunnelOptions []*ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions `json:"TunnelOptions,omitempty" xml:"TunnelOptions,omitempty" type:"Repeated"`
}

func (s ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecification) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecification) GoString() string {
	return s.String()
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecification) GetTunnelOptions() []*ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions {
	return s.TunnelOptions
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecification) SetTunnelOptions(v []*ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions) *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecification {
	s.TunnelOptions = v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecification) Validate() error {
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

type ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions struct {
	CustomerGatewayId   *string                                                                                           `json:"CustomerGatewayId,omitempty" xml:"CustomerGatewayId,omitempty"`
	EnableDpd           *bool                                                                                             `json:"EnableDpd,omitempty" xml:"EnableDpd,omitempty"`
	EnableNatTraversal  *bool                                                                                             `json:"EnableNatTraversal,omitempty" xml:"EnableNatTraversal,omitempty"`
	InternetIp          *string                                                                                           `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	RemoteCaCertificate *string                                                                                           `json:"RemoteCaCertificate,omitempty" xml:"RemoteCaCertificate,omitempty"`
	Role                *string                                                                                           `json:"Role,omitempty" xml:"Role,omitempty"`
	State               *string                                                                                           `json:"State,omitempty" xml:"State,omitempty"`
	TunnelBgpConfig     *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig   `json:"TunnelBgpConfig,omitempty" xml:"TunnelBgpConfig,omitempty" type:"Struct"`
	TunnelId            *string                                                                                           `json:"TunnelId,omitempty" xml:"TunnelId,omitempty"`
	TunnelIkeConfig     *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig   `json:"TunnelIkeConfig,omitempty" xml:"TunnelIkeConfig,omitempty" type:"Struct"`
	TunnelIpsecConfig   *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig `json:"TunnelIpsecConfig,omitempty" xml:"TunnelIpsecConfig,omitempty" type:"Struct"`
	ZoneNo              *string                                                                                           `json:"ZoneNo,omitempty" xml:"ZoneNo,omitempty"`
}

func (s ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions) GoString() string {
	return s.String()
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions) GetCustomerGatewayId() *string {
	return s.CustomerGatewayId
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions) GetEnableDpd() *bool {
	return s.EnableDpd
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions) GetEnableNatTraversal() *bool {
	return s.EnableNatTraversal
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions) GetInternetIp() *string {
	return s.InternetIp
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions) GetRemoteCaCertificate() *string {
	return s.RemoteCaCertificate
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions) GetRole() *string {
	return s.Role
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions) GetState() *string {
	return s.State
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions) GetTunnelBgpConfig() *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig {
	return s.TunnelBgpConfig
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions) GetTunnelId() *string {
	return s.TunnelId
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions) GetTunnelIkeConfig() *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig {
	return s.TunnelIkeConfig
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions) GetTunnelIpsecConfig() *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig {
	return s.TunnelIpsecConfig
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions) GetZoneNo() *string {
	return s.ZoneNo
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions) SetCustomerGatewayId(v string) *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions {
	s.CustomerGatewayId = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions) SetEnableDpd(v bool) *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions {
	s.EnableDpd = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions) SetEnableNatTraversal(v bool) *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions {
	s.EnableNatTraversal = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions) SetInternetIp(v string) *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions {
	s.InternetIp = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions) SetRemoteCaCertificate(v string) *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions {
	s.RemoteCaCertificate = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions) SetRole(v string) *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions {
	s.Role = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions) SetState(v string) *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions {
	s.State = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions) SetTunnelBgpConfig(v *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions {
	s.TunnelBgpConfig = v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions) SetTunnelId(v string) *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions {
	s.TunnelId = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions) SetTunnelIkeConfig(v *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions {
	s.TunnelIkeConfig = v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions) SetTunnelIpsecConfig(v *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions {
	s.TunnelIpsecConfig = v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions) SetZoneNo(v string) *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions {
	s.ZoneNo = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptions) Validate() error {
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

type ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig struct {
	LocalAsn   *int64  `json:"LocalAsn,omitempty" xml:"LocalAsn,omitempty"`
	LocalBgpIp *string `json:"LocalBgpIp,omitempty" xml:"LocalBgpIp,omitempty"`
	PeerAsn    *int64  `json:"PeerAsn,omitempty" xml:"PeerAsn,omitempty"`
	PeerBgpIp  *string `json:"PeerBgpIp,omitempty" xml:"PeerBgpIp,omitempty"`
	TunnelCidr *string `json:"TunnelCidr,omitempty" xml:"TunnelCidr,omitempty"`
}

func (s ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) GoString() string {
	return s.String()
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) GetLocalAsn() *int64 {
	return s.LocalAsn
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) GetLocalBgpIp() *string {
	return s.LocalBgpIp
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) GetPeerAsn() *int64 {
	return s.PeerAsn
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) GetPeerBgpIp() *string {
	return s.PeerBgpIp
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) GetTunnelCidr() *string {
	return s.TunnelCidr
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) SetLocalAsn(v int64) *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig {
	s.LocalAsn = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) SetLocalBgpIp(v string) *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig {
	s.LocalBgpIp = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) SetPeerAsn(v int64) *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig {
	s.PeerAsn = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) SetPeerBgpIp(v string) *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig {
	s.PeerBgpIp = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) SetTunnelCidr(v string) *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig {
	s.TunnelCidr = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig struct {
	IkeAuthAlg  *string `json:"IkeAuthAlg,omitempty" xml:"IkeAuthAlg,omitempty"`
	IkeEncAlg   *string `json:"IkeEncAlg,omitempty" xml:"IkeEncAlg,omitempty"`
	IkeLifetime *int64  `json:"IkeLifetime,omitempty" xml:"IkeLifetime,omitempty"`
	IkeMode     *string `json:"IkeMode,omitempty" xml:"IkeMode,omitempty"`
	IkePfs      *string `json:"IkePfs,omitempty" xml:"IkePfs,omitempty"`
	IkeVersion  *string `json:"IkeVersion,omitempty" xml:"IkeVersion,omitempty"`
	LocalId     *string `json:"LocalId,omitempty" xml:"LocalId,omitempty"`
	Psk         *string `json:"Psk,omitempty" xml:"Psk,omitempty"`
	RemoteId    *string `json:"RemoteId,omitempty" xml:"RemoteId,omitempty"`
}

func (s ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) GoString() string {
	return s.String()
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) GetIkeAuthAlg() *string {
	return s.IkeAuthAlg
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) GetIkeEncAlg() *string {
	return s.IkeEncAlg
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) GetIkeLifetime() *int64 {
	return s.IkeLifetime
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) GetIkeMode() *string {
	return s.IkeMode
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) GetIkePfs() *string {
	return s.IkePfs
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) GetIkeVersion() *string {
	return s.IkeVersion
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) GetLocalId() *string {
	return s.LocalId
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) GetPsk() *string {
	return s.Psk
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) GetRemoteId() *string {
	return s.RemoteId
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) SetIkeAuthAlg(v string) *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig {
	s.IkeAuthAlg = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) SetIkeEncAlg(v string) *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig {
	s.IkeEncAlg = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) SetIkeLifetime(v int64) *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig {
	s.IkeLifetime = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) SetIkeMode(v string) *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig {
	s.IkeMode = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) SetIkePfs(v string) *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig {
	s.IkePfs = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) SetIkeVersion(v string) *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig {
	s.IkeVersion = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) SetLocalId(v string) *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig {
	s.LocalId = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) SetPsk(v string) *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig {
	s.Psk = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) SetRemoteId(v string) *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig {
	s.RemoteId = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig struct {
	IpsecAuthAlg  *string `json:"IpsecAuthAlg,omitempty" xml:"IpsecAuthAlg,omitempty"`
	IpsecEncAlg   *string `json:"IpsecEncAlg,omitempty" xml:"IpsecEncAlg,omitempty"`
	IpsecLifetime *int64  `json:"IpsecLifetime,omitempty" xml:"IpsecLifetime,omitempty"`
	IpsecPfs      *string `json:"IpsecPfs,omitempty" xml:"IpsecPfs,omitempty"`
}

func (s ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) GoString() string {
	return s.String()
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) GetIpsecAuthAlg() *string {
	return s.IpsecAuthAlg
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) GetIpsecEncAlg() *string {
	return s.IpsecEncAlg
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) GetIpsecLifetime() *int64 {
	return s.IpsecLifetime
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) GetIpsecPfs() *string {
	return s.IpsecPfs
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) SetIpsecAuthAlg(v string) *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig {
	s.IpsecAuthAlg = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) SetIpsecEncAlg(v string) *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig {
	s.IpsecEncAlg = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) SetIpsecLifetime(v int64) *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig {
	s.IpsecLifetime = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) SetIpsecPfs(v string) *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig {
	s.IpsecPfs = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyVpnConnectionAttributeResponseBodyVcoHealthCheck struct {
	// The destination IP address.
	//
	// example:
	//
	// 192.168.1.1
	Dip *string `json:"Dip,omitempty" xml:"Dip,omitempty"`
	// Indicates whether the health check feature is enabled for the IPsec-VPN connection.
	//
	// 	- **true**
	//
	// 	- **false**
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

func (s ModifyVpnConnectionAttributeResponseBodyVcoHealthCheck) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnConnectionAttributeResponseBodyVcoHealthCheck) GoString() string {
	return s.String()
}

func (s *ModifyVpnConnectionAttributeResponseBodyVcoHealthCheck) GetDip() *string {
	return s.Dip
}

func (s *ModifyVpnConnectionAttributeResponseBodyVcoHealthCheck) GetEnable() *string {
	return s.Enable
}

func (s *ModifyVpnConnectionAttributeResponseBodyVcoHealthCheck) GetInterval() *int32 {
	return s.Interval
}

func (s *ModifyVpnConnectionAttributeResponseBodyVcoHealthCheck) GetRetry() *int32 {
	return s.Retry
}

func (s *ModifyVpnConnectionAttributeResponseBodyVcoHealthCheck) GetSip() *string {
	return s.Sip
}

func (s *ModifyVpnConnectionAttributeResponseBodyVcoHealthCheck) SetDip(v string) *ModifyVpnConnectionAttributeResponseBodyVcoHealthCheck {
	s.Dip = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyVcoHealthCheck) SetEnable(v string) *ModifyVpnConnectionAttributeResponseBodyVcoHealthCheck {
	s.Enable = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyVcoHealthCheck) SetInterval(v int32) *ModifyVpnConnectionAttributeResponseBodyVcoHealthCheck {
	s.Interval = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyVcoHealthCheck) SetRetry(v int32) *ModifyVpnConnectionAttributeResponseBodyVcoHealthCheck {
	s.Retry = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyVcoHealthCheck) SetSip(v string) *ModifyVpnConnectionAttributeResponseBodyVcoHealthCheck {
	s.Sip = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyVcoHealthCheck) Validate() error {
	return dara.Validate(s)
}

type ModifyVpnConnectionAttributeResponseBodyVpnBgpConfig struct {
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
	// 65530
	LocalAsn *int32 `json:"LocalAsn,omitempty" xml:"LocalAsn,omitempty"`
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
	// 65531
	PeerAsn *int32 `json:"PeerAsn,omitempty" xml:"PeerAsn,omitempty"`
	// The BGP IP address of the data center.
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
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The BGP CIDR block of the IPsec-VPN connection.
	//
	// example:
	//
	// 169.254.11.0/30
	TunnelCidr *string `json:"TunnelCidr,omitempty" xml:"TunnelCidr,omitempty"`
}

func (s ModifyVpnConnectionAttributeResponseBodyVpnBgpConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnConnectionAttributeResponseBodyVpnBgpConfig) GoString() string {
	return s.String()
}

func (s *ModifyVpnConnectionAttributeResponseBodyVpnBgpConfig) GetEnableBgp() *string {
	return s.EnableBgp
}

func (s *ModifyVpnConnectionAttributeResponseBodyVpnBgpConfig) GetLocalAsn() *int32 {
	return s.LocalAsn
}

func (s *ModifyVpnConnectionAttributeResponseBodyVpnBgpConfig) GetLocalBgpIp() *string {
	return s.LocalBgpIp
}

func (s *ModifyVpnConnectionAttributeResponseBodyVpnBgpConfig) GetPeerAsn() *int32 {
	return s.PeerAsn
}

func (s *ModifyVpnConnectionAttributeResponseBodyVpnBgpConfig) GetPeerBgpIp() *string {
	return s.PeerBgpIp
}

func (s *ModifyVpnConnectionAttributeResponseBodyVpnBgpConfig) GetStatus() *string {
	return s.Status
}

func (s *ModifyVpnConnectionAttributeResponseBodyVpnBgpConfig) GetTunnelCidr() *string {
	return s.TunnelCidr
}

func (s *ModifyVpnConnectionAttributeResponseBodyVpnBgpConfig) SetEnableBgp(v string) *ModifyVpnConnectionAttributeResponseBodyVpnBgpConfig {
	s.EnableBgp = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyVpnBgpConfig) SetLocalAsn(v int32) *ModifyVpnConnectionAttributeResponseBodyVpnBgpConfig {
	s.LocalAsn = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyVpnBgpConfig) SetLocalBgpIp(v string) *ModifyVpnConnectionAttributeResponseBodyVpnBgpConfig {
	s.LocalBgpIp = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyVpnBgpConfig) SetPeerAsn(v int32) *ModifyVpnConnectionAttributeResponseBodyVpnBgpConfig {
	s.PeerAsn = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyVpnBgpConfig) SetPeerBgpIp(v string) *ModifyVpnConnectionAttributeResponseBodyVpnBgpConfig {
	s.PeerBgpIp = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyVpnBgpConfig) SetStatus(v string) *ModifyVpnConnectionAttributeResponseBodyVpnBgpConfig {
	s.Status = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyVpnBgpConfig) SetTunnelCidr(v string) *ModifyVpnConnectionAttributeResponseBodyVpnBgpConfig {
	s.TunnelCidr = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponseBodyVpnBgpConfig) Validate() error {
	return dara.Validate(s)
}
