// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTunnelAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerGatewayId(v string) *ModifyTunnelAttributeResponseBody
	GetCustomerGatewayId() *string
	SetEnableDpd(v bool) *ModifyTunnelAttributeResponseBody
	GetEnableDpd() *bool
	SetEnableNatTraversal(v bool) *ModifyTunnelAttributeResponseBody
	GetEnableNatTraversal() *bool
	SetInternetIp(v string) *ModifyTunnelAttributeResponseBody
	GetInternetIp() *string
	SetRemoteCaCertificate(v string) *ModifyTunnelAttributeResponseBody
	GetRemoteCaCertificate() *string
	SetRequestId(v string) *ModifyTunnelAttributeResponseBody
	GetRequestId() *string
	SetRole(v string) *ModifyTunnelAttributeResponseBody
	GetRole() *string
	SetState(v string) *ModifyTunnelAttributeResponseBody
	GetState() *string
	SetTunnelBgpConfig(v *ModifyTunnelAttributeResponseBodyTunnelBgpConfig) *ModifyTunnelAttributeResponseBody
	GetTunnelBgpConfig() *ModifyTunnelAttributeResponseBodyTunnelBgpConfig
	SetTunnelId(v string) *ModifyTunnelAttributeResponseBody
	GetTunnelId() *string
	SetTunnelIkeConfig(v *ModifyTunnelAttributeResponseBodyTunnelIkeConfig) *ModifyTunnelAttributeResponseBody
	GetTunnelIkeConfig() *ModifyTunnelAttributeResponseBodyTunnelIkeConfig
	SetTunnelIpsecConfig(v *ModifyTunnelAttributeResponseBodyTunnelIpsecConfig) *ModifyTunnelAttributeResponseBody
	GetTunnelIpsecConfig() *ModifyTunnelAttributeResponseBodyTunnelIpsecConfig
	SetZoneNo(v string) *ModifyTunnelAttributeResponseBody
	GetZoneNo() *string
}

type ModifyTunnelAttributeResponseBody struct {
	// The ID of the customer gateway associated with the customer gateway.
	//
	// example:
	//
	// cgw-p0wx48ayhrygitm80****
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" xml:"CustomerGatewayId,omitempty"`
	// Indicates whether DPD is enabled. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// example:
	//
	// true
	EnableDpd *bool `json:"EnableDpd,omitempty" xml:"EnableDpd,omitempty"`
	// Indicates whether NAT traversal is enabled. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// example:
	//
	// true
	EnableNatTraversal *bool `json:"EnableNatTraversal,omitempty" xml:"EnableNatTraversal,omitempty"`
	// The tunnel IP address.
	//
	// example:
	//
	// 47.XX.XX.87
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The peer CA certificate when a VPN gateway that uses an SM certificate is used to create the IPsec connection.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- MIIB7zCCAZW***	- -----END CERTIFICATE-----
	RemoteCaCertificate *string `json:"RemoteCaCertificate,omitempty" xml:"RemoteCaCertificate,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E6F36FF0-9544-3AEE-8673-A4647D50064C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tunnel role. Valid values:
	//
	// 	- **master**
	//
	// 	- **slave**
	//
	// example:
	//
	// master
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The tunnel status. Valid values:
	//
	// 	- **active**
	//
	// 	- **updating**
	//
	// 	- **deleting**
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The BGP configuration.
	TunnelBgpConfig *ModifyTunnelAttributeResponseBodyTunnelBgpConfig `json:"TunnelBgpConfig,omitempty" xml:"TunnelBgpConfig,omitempty" type:"Struct"`
	// The tunnel ID.
	//
	// example:
	//
	// tun-gbyz2e070xzo93****
	TunnelId *string `json:"TunnelId,omitempty" xml:"TunnelId,omitempty"`
	// The Phase 1 configuration.
	TunnelIkeConfig *ModifyTunnelAttributeResponseBodyTunnelIkeConfig `json:"TunnelIkeConfig,omitempty" xml:"TunnelIkeConfig,omitempty" type:"Struct"`
	// The configurations of IPsec Phase 2.
	TunnelIpsecConfig *ModifyTunnelAttributeResponseBodyTunnelIpsecConfig `json:"TunnelIpsecConfig,omitempty" xml:"TunnelIpsecConfig,omitempty" type:"Struct"`
	// The tunnel zone.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneNo *string `json:"ZoneNo,omitempty" xml:"ZoneNo,omitempty"`
}

func (s ModifyTunnelAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyTunnelAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTunnelAttributeResponseBody) GetCustomerGatewayId() *string {
	return s.CustomerGatewayId
}

func (s *ModifyTunnelAttributeResponseBody) GetEnableDpd() *bool {
	return s.EnableDpd
}

func (s *ModifyTunnelAttributeResponseBody) GetEnableNatTraversal() *bool {
	return s.EnableNatTraversal
}

func (s *ModifyTunnelAttributeResponseBody) GetInternetIp() *string {
	return s.InternetIp
}

func (s *ModifyTunnelAttributeResponseBody) GetRemoteCaCertificate() *string {
	return s.RemoteCaCertificate
}

func (s *ModifyTunnelAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyTunnelAttributeResponseBody) GetRole() *string {
	return s.Role
}

func (s *ModifyTunnelAttributeResponseBody) GetState() *string {
	return s.State
}

func (s *ModifyTunnelAttributeResponseBody) GetTunnelBgpConfig() *ModifyTunnelAttributeResponseBodyTunnelBgpConfig {
	return s.TunnelBgpConfig
}

func (s *ModifyTunnelAttributeResponseBody) GetTunnelId() *string {
	return s.TunnelId
}

func (s *ModifyTunnelAttributeResponseBody) GetTunnelIkeConfig() *ModifyTunnelAttributeResponseBodyTunnelIkeConfig {
	return s.TunnelIkeConfig
}

func (s *ModifyTunnelAttributeResponseBody) GetTunnelIpsecConfig() *ModifyTunnelAttributeResponseBodyTunnelIpsecConfig {
	return s.TunnelIpsecConfig
}

func (s *ModifyTunnelAttributeResponseBody) GetZoneNo() *string {
	return s.ZoneNo
}

func (s *ModifyTunnelAttributeResponseBody) SetCustomerGatewayId(v string) *ModifyTunnelAttributeResponseBody {
	s.CustomerGatewayId = &v
	return s
}

func (s *ModifyTunnelAttributeResponseBody) SetEnableDpd(v bool) *ModifyTunnelAttributeResponseBody {
	s.EnableDpd = &v
	return s
}

func (s *ModifyTunnelAttributeResponseBody) SetEnableNatTraversal(v bool) *ModifyTunnelAttributeResponseBody {
	s.EnableNatTraversal = &v
	return s
}

func (s *ModifyTunnelAttributeResponseBody) SetInternetIp(v string) *ModifyTunnelAttributeResponseBody {
	s.InternetIp = &v
	return s
}

func (s *ModifyTunnelAttributeResponseBody) SetRemoteCaCertificate(v string) *ModifyTunnelAttributeResponseBody {
	s.RemoteCaCertificate = &v
	return s
}

func (s *ModifyTunnelAttributeResponseBody) SetRequestId(v string) *ModifyTunnelAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTunnelAttributeResponseBody) SetRole(v string) *ModifyTunnelAttributeResponseBody {
	s.Role = &v
	return s
}

func (s *ModifyTunnelAttributeResponseBody) SetState(v string) *ModifyTunnelAttributeResponseBody {
	s.State = &v
	return s
}

func (s *ModifyTunnelAttributeResponseBody) SetTunnelBgpConfig(v *ModifyTunnelAttributeResponseBodyTunnelBgpConfig) *ModifyTunnelAttributeResponseBody {
	s.TunnelBgpConfig = v
	return s
}

func (s *ModifyTunnelAttributeResponseBody) SetTunnelId(v string) *ModifyTunnelAttributeResponseBody {
	s.TunnelId = &v
	return s
}

func (s *ModifyTunnelAttributeResponseBody) SetTunnelIkeConfig(v *ModifyTunnelAttributeResponseBodyTunnelIkeConfig) *ModifyTunnelAttributeResponseBody {
	s.TunnelIkeConfig = v
	return s
}

func (s *ModifyTunnelAttributeResponseBody) SetTunnelIpsecConfig(v *ModifyTunnelAttributeResponseBodyTunnelIpsecConfig) *ModifyTunnelAttributeResponseBody {
	s.TunnelIpsecConfig = v
	return s
}

func (s *ModifyTunnelAttributeResponseBody) SetZoneNo(v string) *ModifyTunnelAttributeResponseBody {
	s.ZoneNo = &v
	return s
}

func (s *ModifyTunnelAttributeResponseBody) Validate() error {
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

type ModifyTunnelAttributeResponseBodyTunnelBgpConfig struct {
	// Indicates whether the BGP feature is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	EnableBgp *bool `json:"EnableBgp,omitempty" xml:"EnableBgp,omitempty"`
	// The local ASN.
	//
	// example:
	//
	// 65530
	LocalAsn *int64 `json:"LocalAsn,omitempty" xml:"LocalAsn,omitempty"`
	// The BGP IP address of the tunnel.
	//
	// example:
	//
	// 169.254.11.1
	LocalBgpIp *string `json:"LocalBgpIp,omitempty" xml:"LocalBgpIp,omitempty"`
	// The peer ASN.
	//
	// example:
	//
	// 65531
	PeerAsn *int64 `json:"PeerAsn,omitempty" xml:"PeerAsn,omitempty"`
	// The BGP IP address of the peer.
	//
	// example:
	//
	// 169.254.11.2
	PeerBgpIp *string `json:"PeerBgpIp,omitempty" xml:"PeerBgpIp,omitempty"`
	// The CIDR block to which the tunnel BGP IP address belongs.
	//
	// example:
	//
	// 169.254.11.0/30
	TunnelCidr *string `json:"TunnelCidr,omitempty" xml:"TunnelCidr,omitempty"`
}

func (s ModifyTunnelAttributeResponseBodyTunnelBgpConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyTunnelAttributeResponseBodyTunnelBgpConfig) GoString() string {
	return s.String()
}

func (s *ModifyTunnelAttributeResponseBodyTunnelBgpConfig) GetEnableBgp() *bool {
	return s.EnableBgp
}

func (s *ModifyTunnelAttributeResponseBodyTunnelBgpConfig) GetLocalAsn() *int64 {
	return s.LocalAsn
}

func (s *ModifyTunnelAttributeResponseBodyTunnelBgpConfig) GetLocalBgpIp() *string {
	return s.LocalBgpIp
}

func (s *ModifyTunnelAttributeResponseBodyTunnelBgpConfig) GetPeerAsn() *int64 {
	return s.PeerAsn
}

func (s *ModifyTunnelAttributeResponseBodyTunnelBgpConfig) GetPeerBgpIp() *string {
	return s.PeerBgpIp
}

func (s *ModifyTunnelAttributeResponseBodyTunnelBgpConfig) GetTunnelCidr() *string {
	return s.TunnelCidr
}

func (s *ModifyTunnelAttributeResponseBodyTunnelBgpConfig) SetEnableBgp(v bool) *ModifyTunnelAttributeResponseBodyTunnelBgpConfig {
	s.EnableBgp = &v
	return s
}

func (s *ModifyTunnelAttributeResponseBodyTunnelBgpConfig) SetLocalAsn(v int64) *ModifyTunnelAttributeResponseBodyTunnelBgpConfig {
	s.LocalAsn = &v
	return s
}

func (s *ModifyTunnelAttributeResponseBodyTunnelBgpConfig) SetLocalBgpIp(v string) *ModifyTunnelAttributeResponseBodyTunnelBgpConfig {
	s.LocalBgpIp = &v
	return s
}

func (s *ModifyTunnelAttributeResponseBodyTunnelBgpConfig) SetPeerAsn(v int64) *ModifyTunnelAttributeResponseBodyTunnelBgpConfig {
	s.PeerAsn = &v
	return s
}

func (s *ModifyTunnelAttributeResponseBodyTunnelBgpConfig) SetPeerBgpIp(v string) *ModifyTunnelAttributeResponseBodyTunnelBgpConfig {
	s.PeerBgpIp = &v
	return s
}

func (s *ModifyTunnelAttributeResponseBodyTunnelBgpConfig) SetTunnelCidr(v string) *ModifyTunnelAttributeResponseBodyTunnelBgpConfig {
	s.TunnelCidr = &v
	return s
}

func (s *ModifyTunnelAttributeResponseBodyTunnelBgpConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyTunnelAttributeResponseBodyTunnelIkeConfig struct {
	// The IKE authentication algorithm.
	//
	// example:
	//
	// sha1
	IkeAuthAlg *string `json:"IkeAuthAlg,omitempty" xml:"IkeAuthAlg,omitempty"`
	// The IKE encryption algorithm.
	//
	// example:
	//
	// aes
	IkeEncAlg *string `json:"IkeEncAlg,omitempty" xml:"IkeEncAlg,omitempty"`
	// The IKE lifetime. Unit: seconds.
	//
	// example:
	//
	// 86400
	IkeLifetime *int64 `json:"IkeLifetime,omitempty" xml:"IkeLifetime,omitempty"`
	// The IKE negotiation mode.
	//
	// 	- **main:*	- This mode offers higher security during negotiations.
	//
	// 	- **aggressive**: This mode is faster and has a higher success rate.
	//
	// example:
	//
	// main
	IkeMode *string `json:"IkeMode,omitempty" xml:"IkeMode,omitempty"`
	// The DH group.
	//
	// example:
	//
	// group2
	IkePfs *string `json:"IkePfs,omitempty" xml:"IkePfs,omitempty"`
	// The IKE version.
	//
	// 	- **ikev1**
	//
	// 	- **ikev2**
	//
	// Compared with IKEv1, IKEv2 simplifies the SA negotiation process and provides better support for scenarios with multiple CIDR blocks.
	//
	// example:
	//
	// ikev2
	IkeVersion *string `json:"IkeVersion,omitempty" xml:"IkeVersion,omitempty"`
	// The tunnel identifier. The identifier supports FQDNs and IP addresses. The default value is the tunnel IP address.
	//
	// example:
	//
	// 47.XX.XX.87
	LocalId *string `json:"LocalId,omitempty" xml:"LocalId,omitempty"`
	// The pre-shared key.
	//
	// example:
	//
	// 123456****
	Psk *string `json:"Psk,omitempty" xml:"Psk,omitempty"`
	// The peer identifier. The identifier supports FQDNs and IP addresses. The default identifier is the IP address of the customer gateway associated with the tunnel.
	//
	// example:
	//
	// 47.XX.XX.207
	RemoteId *string `json:"RemoteId,omitempty" xml:"RemoteId,omitempty"`
}

func (s ModifyTunnelAttributeResponseBodyTunnelIkeConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyTunnelAttributeResponseBodyTunnelIkeConfig) GoString() string {
	return s.String()
}

func (s *ModifyTunnelAttributeResponseBodyTunnelIkeConfig) GetIkeAuthAlg() *string {
	return s.IkeAuthAlg
}

func (s *ModifyTunnelAttributeResponseBodyTunnelIkeConfig) GetIkeEncAlg() *string {
	return s.IkeEncAlg
}

func (s *ModifyTunnelAttributeResponseBodyTunnelIkeConfig) GetIkeLifetime() *int64 {
	return s.IkeLifetime
}

func (s *ModifyTunnelAttributeResponseBodyTunnelIkeConfig) GetIkeMode() *string {
	return s.IkeMode
}

func (s *ModifyTunnelAttributeResponseBodyTunnelIkeConfig) GetIkePfs() *string {
	return s.IkePfs
}

func (s *ModifyTunnelAttributeResponseBodyTunnelIkeConfig) GetIkeVersion() *string {
	return s.IkeVersion
}

func (s *ModifyTunnelAttributeResponseBodyTunnelIkeConfig) GetLocalId() *string {
	return s.LocalId
}

func (s *ModifyTunnelAttributeResponseBodyTunnelIkeConfig) GetPsk() *string {
	return s.Psk
}

func (s *ModifyTunnelAttributeResponseBodyTunnelIkeConfig) GetRemoteId() *string {
	return s.RemoteId
}

func (s *ModifyTunnelAttributeResponseBodyTunnelIkeConfig) SetIkeAuthAlg(v string) *ModifyTunnelAttributeResponseBodyTunnelIkeConfig {
	s.IkeAuthAlg = &v
	return s
}

func (s *ModifyTunnelAttributeResponseBodyTunnelIkeConfig) SetIkeEncAlg(v string) *ModifyTunnelAttributeResponseBodyTunnelIkeConfig {
	s.IkeEncAlg = &v
	return s
}

func (s *ModifyTunnelAttributeResponseBodyTunnelIkeConfig) SetIkeLifetime(v int64) *ModifyTunnelAttributeResponseBodyTunnelIkeConfig {
	s.IkeLifetime = &v
	return s
}

func (s *ModifyTunnelAttributeResponseBodyTunnelIkeConfig) SetIkeMode(v string) *ModifyTunnelAttributeResponseBodyTunnelIkeConfig {
	s.IkeMode = &v
	return s
}

func (s *ModifyTunnelAttributeResponseBodyTunnelIkeConfig) SetIkePfs(v string) *ModifyTunnelAttributeResponseBodyTunnelIkeConfig {
	s.IkePfs = &v
	return s
}

func (s *ModifyTunnelAttributeResponseBodyTunnelIkeConfig) SetIkeVersion(v string) *ModifyTunnelAttributeResponseBodyTunnelIkeConfig {
	s.IkeVersion = &v
	return s
}

func (s *ModifyTunnelAttributeResponseBodyTunnelIkeConfig) SetLocalId(v string) *ModifyTunnelAttributeResponseBodyTunnelIkeConfig {
	s.LocalId = &v
	return s
}

func (s *ModifyTunnelAttributeResponseBodyTunnelIkeConfig) SetPsk(v string) *ModifyTunnelAttributeResponseBodyTunnelIkeConfig {
	s.Psk = &v
	return s
}

func (s *ModifyTunnelAttributeResponseBodyTunnelIkeConfig) SetRemoteId(v string) *ModifyTunnelAttributeResponseBodyTunnelIkeConfig {
	s.RemoteId = &v
	return s
}

func (s *ModifyTunnelAttributeResponseBodyTunnelIkeConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyTunnelAttributeResponseBodyTunnelIpsecConfig struct {
	// The IPsec authentication algorithm.
	//
	// example:
	//
	// sha1
	IpsecAuthAlg *string `json:"IpsecAuthAlg,omitempty" xml:"IpsecAuthAlg,omitempty"`
	// The IPsec encryption algorithm.
	//
	// example:
	//
	// aes
	IpsecEncAlg *string `json:"IpsecEncAlg,omitempty" xml:"IpsecEncAlg,omitempty"`
	// The IPsec lifetime. Unit: seconds.
	//
	// example:
	//
	// 86400
	IpsecLifetime *int64 `json:"IpsecLifetime,omitempty" xml:"IpsecLifetime,omitempty"`
	// The DH group.
	//
	// example:
	//
	// group2
	IpsecPfs *string `json:"IpsecPfs,omitempty" xml:"IpsecPfs,omitempty"`
}

func (s ModifyTunnelAttributeResponseBodyTunnelIpsecConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyTunnelAttributeResponseBodyTunnelIpsecConfig) GoString() string {
	return s.String()
}

func (s *ModifyTunnelAttributeResponseBodyTunnelIpsecConfig) GetIpsecAuthAlg() *string {
	return s.IpsecAuthAlg
}

func (s *ModifyTunnelAttributeResponseBodyTunnelIpsecConfig) GetIpsecEncAlg() *string {
	return s.IpsecEncAlg
}

func (s *ModifyTunnelAttributeResponseBodyTunnelIpsecConfig) GetIpsecLifetime() *int64 {
	return s.IpsecLifetime
}

func (s *ModifyTunnelAttributeResponseBodyTunnelIpsecConfig) GetIpsecPfs() *string {
	return s.IpsecPfs
}

func (s *ModifyTunnelAttributeResponseBodyTunnelIpsecConfig) SetIpsecAuthAlg(v string) *ModifyTunnelAttributeResponseBodyTunnelIpsecConfig {
	s.IpsecAuthAlg = &v
	return s
}

func (s *ModifyTunnelAttributeResponseBodyTunnelIpsecConfig) SetIpsecEncAlg(v string) *ModifyTunnelAttributeResponseBodyTunnelIpsecConfig {
	s.IpsecEncAlg = &v
	return s
}

func (s *ModifyTunnelAttributeResponseBodyTunnelIpsecConfig) SetIpsecLifetime(v int64) *ModifyTunnelAttributeResponseBodyTunnelIpsecConfig {
	s.IpsecLifetime = &v
	return s
}

func (s *ModifyTunnelAttributeResponseBodyTunnelIpsecConfig) SetIpsecPfs(v string) *ModifyTunnelAttributeResponseBodyTunnelIpsecConfig {
	s.IpsecPfs = &v
	return s
}

func (s *ModifyTunnelAttributeResponseBodyTunnelIpsecConfig) Validate() error {
	return dara.Validate(s)
}
