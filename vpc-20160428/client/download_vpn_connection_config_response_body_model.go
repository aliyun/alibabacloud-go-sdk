// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadVpnConnectionConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DownloadVpnConnectionConfigResponseBody
	GetRequestId() *string
	SetVpnConnectionConfig(v *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfig) *DownloadVpnConnectionConfigResponseBody
	GetVpnConnectionConfig() *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfig
}

type DownloadVpnConnectionConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0C68048B-0F70-40DA-B8AE-1B79B5CF62E3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The configurations of the peer gateway device.
	VpnConnectionConfig *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfig `json:"VpnConnectionConfig,omitempty" xml:"VpnConnectionConfig,omitempty" type:"Struct"`
}

func (s DownloadVpnConnectionConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DownloadVpnConnectionConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadVpnConnectionConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DownloadVpnConnectionConfigResponseBody) GetVpnConnectionConfig() *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfig {
	return s.VpnConnectionConfig
}

func (s *DownloadVpnConnectionConfigResponseBody) SetRequestId(v string) *DownloadVpnConnectionConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBody) SetVpnConnectionConfig(v *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfig) *DownloadVpnConnectionConfigResponseBody {
	s.VpnConnectionConfig = v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBody) Validate() error {
	if s.VpnConnectionConfig != nil {
		if err := s.VpnConnectionConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DownloadVpnConnectionConfigResponseBodyVpnConnectionConfig struct {
	BgpConfigs *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigBgpConfigs `json:"BgpConfigs,omitempty" xml:"BgpConfigs,omitempty" type:"Struct"`
	// The configurations of Phase 1 negotiations.
	IkeConfig *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIkeConfig `json:"IkeConfig,omitempty" xml:"IkeConfig,omitempty" type:"Struct"`
	// The configurations of Phase 2 negotiations.
	IpsecConfig *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIpsecConfig `json:"IpsecConfig,omitempty" xml:"IpsecConfig,omitempty" type:"Struct"`
	// The identifier of the customer gateway.
	//
	// example:
	//
	// 139.196.XX.XX
	Local *string `json:"Local,omitempty" xml:"Local,omitempty"`
	// The CIDR block on the data center side.
	//
	// example:
	//
	// 10.0.0.0/8
	LocalSubnet *string `json:"LocalSubnet,omitempty" xml:"LocalSubnet,omitempty"`
	// The identifier of the VPN gateway.
	//
	// example:
	//
	// 116.62.XX.XX
	Remote *string `json:"Remote,omitempty" xml:"Remote,omitempty"`
	// The CIDR block on the virtual private cloud (VPC) side.
	//
	// example:
	//
	// 192.168.0.0/16
	RemoteSubnet *string `json:"RemoteSubnet,omitempty" xml:"RemoteSubnet,omitempty"`
	// The tunnel configurations of the peer gateway device.
	//
	// The parameters in **TunnelsConfig*	- are returned only when the IPsec-VPN connection supports the dual-tunnel mode.
	TunnelsConfig *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfig `json:"TunnelsConfig,omitempty" xml:"TunnelsConfig,omitempty" type:"Struct"`
}

func (s DownloadVpnConnectionConfigResponseBodyVpnConnectionConfig) String() string {
	return dara.Prettify(s)
}

func (s DownloadVpnConnectionConfigResponseBodyVpnConnectionConfig) GoString() string {
	return s.String()
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfig) GetBgpConfigs() *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigBgpConfigs {
	return s.BgpConfigs
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfig) GetIkeConfig() *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIkeConfig {
	return s.IkeConfig
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfig) GetIpsecConfig() *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIpsecConfig {
	return s.IpsecConfig
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfig) GetLocal() *string {
	return s.Local
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfig) GetLocalSubnet() *string {
	return s.LocalSubnet
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfig) GetRemote() *string {
	return s.Remote
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfig) GetRemoteSubnet() *string {
	return s.RemoteSubnet
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfig) GetTunnelsConfig() *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfig {
	return s.TunnelsConfig
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfig) SetBgpConfigs(v *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigBgpConfigs) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfig {
	s.BgpConfigs = v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfig) SetIkeConfig(v *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIkeConfig) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfig {
	s.IkeConfig = v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfig) SetIpsecConfig(v *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIpsecConfig) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfig {
	s.IpsecConfig = v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfig) SetLocal(v string) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfig {
	s.Local = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfig) SetLocalSubnet(v string) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfig {
	s.LocalSubnet = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfig) SetRemote(v string) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfig {
	s.Remote = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfig) SetRemoteSubnet(v string) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfig {
	s.RemoteSubnet = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfig) SetTunnelsConfig(v *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfig) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfig {
	s.TunnelsConfig = v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfig) Validate() error {
	if s.BgpConfigs != nil {
		if err := s.BgpConfigs.Validate(); err != nil {
			return err
		}
	}
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
	if s.TunnelsConfig != nil {
		if err := s.TunnelsConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigBgpConfigs struct {
	BgpConfig []*DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigBgpConfigsBgpConfig `json:"BgpConfig,omitempty" xml:"BgpConfig,omitempty" type:"Repeated"`
}

func (s DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigBgpConfigs) String() string {
	return dara.Prettify(s)
}

func (s DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigBgpConfigs) GoString() string {
	return s.String()
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigBgpConfigs) GetBgpConfig() []*DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigBgpConfigsBgpConfig {
	return s.BgpConfig
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigBgpConfigs) SetBgpConfig(v []*DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigBgpConfigsBgpConfig) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigBgpConfigs {
	s.BgpConfig = v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigBgpConfigs) Validate() error {
	if s.BgpConfig != nil {
		for _, item := range s.BgpConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigBgpConfigsBgpConfig struct {
	LocalAsn   *string `json:"LocalAsn,omitempty" xml:"LocalAsn,omitempty"`
	LocalBgpIp *string `json:"LocalBgpIp,omitempty" xml:"LocalBgpIp,omitempty"`
	PeerAsn    *string `json:"PeerAsn,omitempty" xml:"PeerAsn,omitempty"`
	PeerBgpIp  *string `json:"PeerBgpIp,omitempty" xml:"PeerBgpIp,omitempty"`
	TunnelCidr *string `json:"TunnelCidr,omitempty" xml:"TunnelCidr,omitempty"`
	TunnelId   *string `json:"TunnelId,omitempty" xml:"TunnelId,omitempty"`
}

func (s DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigBgpConfigsBgpConfig) String() string {
	return dara.Prettify(s)
}

func (s DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigBgpConfigsBgpConfig) GoString() string {
	return s.String()
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigBgpConfigsBgpConfig) GetLocalAsn() *string {
	return s.LocalAsn
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigBgpConfigsBgpConfig) GetLocalBgpIp() *string {
	return s.LocalBgpIp
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigBgpConfigsBgpConfig) GetPeerAsn() *string {
	return s.PeerAsn
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigBgpConfigsBgpConfig) GetPeerBgpIp() *string {
	return s.PeerBgpIp
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigBgpConfigsBgpConfig) GetTunnelCidr() *string {
	return s.TunnelCidr
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigBgpConfigsBgpConfig) GetTunnelId() *string {
	return s.TunnelId
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigBgpConfigsBgpConfig) SetLocalAsn(v string) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigBgpConfigsBgpConfig {
	s.LocalAsn = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigBgpConfigsBgpConfig) SetLocalBgpIp(v string) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigBgpConfigsBgpConfig {
	s.LocalBgpIp = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigBgpConfigsBgpConfig) SetPeerAsn(v string) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigBgpConfigsBgpConfig {
	s.PeerAsn = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigBgpConfigsBgpConfig) SetPeerBgpIp(v string) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigBgpConfigsBgpConfig {
	s.PeerBgpIp = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigBgpConfigsBgpConfig) SetTunnelCidr(v string) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigBgpConfigsBgpConfig {
	s.TunnelCidr = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigBgpConfigsBgpConfig) SetTunnelId(v string) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigBgpConfigsBgpConfig {
	s.TunnelId = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigBgpConfigsBgpConfig) Validate() error {
	return dara.Validate(s)
}

type DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIkeConfig struct {
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
	// The IKE negotiation mode. Valid values:
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
	// The IKE version.
	//
	// example:
	//
	// ikev1
	IkeVersion *string `json:"IkeVersion,omitempty" xml:"IkeVersion,omitempty"`
	// The identifier of the customer gateway. FQDN and IP formats are supported. The default value is the IP address of the customer gateway.
	//
	// example:
	//
	// 116.62.XX.XX
	LocalId *string `json:"LocalId,omitempty" xml:"LocalId,omitempty"`
	// The pre-shared key.
	//
	// example:
	//
	// pgw6dy7d1i8i****
	Psk *string `json:"Psk,omitempty" xml:"Psk,omitempty"`
	// The identifier of the VPN gateway. FQDN and IP formats are supported. The default value is the IP address of the VPN gateway.
	//
	// example:
	//
	// 139.196.XX.XX
	RemoteId *string `json:"RemoteId,omitempty" xml:"RemoteId,omitempty"`
}

func (s DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIkeConfig) String() string {
	return dara.Prettify(s)
}

func (s DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIkeConfig) GoString() string {
	return s.String()
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIkeConfig) GetIkeAuthAlg() *string {
	return s.IkeAuthAlg
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIkeConfig) GetIkeEncAlg() *string {
	return s.IkeEncAlg
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIkeConfig) GetIkeLifetime() *int64 {
	return s.IkeLifetime
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIkeConfig) GetIkeMode() *string {
	return s.IkeMode
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIkeConfig) GetIkePfs() *string {
	return s.IkePfs
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIkeConfig) GetIkeVersion() *string {
	return s.IkeVersion
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIkeConfig) GetLocalId() *string {
	return s.LocalId
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIkeConfig) GetPsk() *string {
	return s.Psk
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIkeConfig) GetRemoteId() *string {
	return s.RemoteId
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIkeConfig) SetIkeAuthAlg(v string) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIkeConfig {
	s.IkeAuthAlg = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIkeConfig) SetIkeEncAlg(v string) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIkeConfig {
	s.IkeEncAlg = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIkeConfig) SetIkeLifetime(v int64) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIkeConfig {
	s.IkeLifetime = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIkeConfig) SetIkeMode(v string) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIkeConfig {
	s.IkeMode = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIkeConfig) SetIkePfs(v string) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIkeConfig {
	s.IkePfs = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIkeConfig) SetIkeVersion(v string) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIkeConfig {
	s.IkeVersion = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIkeConfig) SetLocalId(v string) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIkeConfig {
	s.LocalId = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIkeConfig) SetPsk(v string) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIkeConfig {
	s.Psk = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIkeConfig) SetRemoteId(v string) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIkeConfig {
	s.RemoteId = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIkeConfig) Validate() error {
	return dara.Validate(s)
}

type DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIpsecConfig struct {
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

func (s DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIpsecConfig) String() string {
	return dara.Prettify(s)
}

func (s DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIpsecConfig) GoString() string {
	return s.String()
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIpsecConfig) GetIpsecAuthAlg() *string {
	return s.IpsecAuthAlg
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIpsecConfig) GetIpsecEncAlg() *string {
	return s.IpsecEncAlg
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIpsecConfig) GetIpsecLifetime() *int64 {
	return s.IpsecLifetime
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIpsecConfig) GetIpsecPfs() *string {
	return s.IpsecPfs
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIpsecConfig) SetIpsecAuthAlg(v string) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIpsecConfig {
	s.IpsecAuthAlg = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIpsecConfig) SetIpsecEncAlg(v string) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIpsecConfig {
	s.IpsecEncAlg = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIpsecConfig) SetIpsecLifetime(v int64) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIpsecConfig {
	s.IpsecLifetime = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIpsecConfig) SetIpsecPfs(v string) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIpsecConfig {
	s.IpsecPfs = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigIpsecConfig) Validate() error {
	return dara.Validate(s)
}

type DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfig struct {
	TunnelConfig []*DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfig `json:"TunnelConfig,omitempty" xml:"TunnelConfig,omitempty" type:"Repeated"`
}

func (s DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfig) String() string {
	return dara.Prettify(s)
}

func (s DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfig) GoString() string {
	return s.String()
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfig) GetTunnelConfig() []*DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfig {
	return s.TunnelConfig
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfig) SetTunnelConfig(v []*DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfig) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfig {
	s.TunnelConfig = v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfig) Validate() error {
	if s.TunnelConfig != nil {
		for _, item := range s.TunnelConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfig struct {
	// The configurations of Phase 1 negotiations.
	IkeConfig *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIkeConfig `json:"IkeConfig,omitempty" xml:"IkeConfig,omitempty" type:"Struct"`
	// The configurations of Phase 2 negotiations.
	IpsecConfig *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIpsecConfig `json:"IpsecConfig,omitempty" xml:"IpsecConfig,omitempty" type:"Struct"`
	// The identifier of the tunnel on the data center side.
	//
	// example:
	//
	// 47.21.XX.XX
	Local *string `json:"Local,omitempty" xml:"Local,omitempty"`
	// The identifier of the tunnel on the Alibaba Cloud side.
	//
	// example:
	//
	// 47.24.XX.XX
	Remote *string `json:"Remote,omitempty" xml:"Remote,omitempty"`
	// The tunnel ID.
	//
	// example:
	//
	// tun-opsqc4d97wni27****
	TunnelId *string `json:"TunnelId,omitempty" xml:"TunnelId,omitempty"`
}

func (s DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfig) String() string {
	return dara.Prettify(s)
}

func (s DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfig) GoString() string {
	return s.String()
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfig) GetIkeConfig() *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIkeConfig {
	return s.IkeConfig
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfig) GetIpsecConfig() *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIpsecConfig {
	return s.IpsecConfig
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfig) GetLocal() *string {
	return s.Local
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfig) GetRemote() *string {
	return s.Remote
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfig) GetTunnelId() *string {
	return s.TunnelId
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfig) SetIkeConfig(v *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIkeConfig) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfig {
	s.IkeConfig = v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfig) SetIpsecConfig(v *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIpsecConfig) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfig {
	s.IpsecConfig = v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfig) SetLocal(v string) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfig {
	s.Local = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfig) SetRemote(v string) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfig {
	s.Remote = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfig) SetTunnelId(v string) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfig {
	s.TunnelId = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfig) Validate() error {
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
	return nil
}

type DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIkeConfig struct {
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
	// The IKE negotiation mode. Valid values:
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
	// The IKE version.
	//
	// example:
	//
	// ikev1
	IkeVersion *string `json:"IkeVersion,omitempty" xml:"IkeVersion,omitempty"`
	// The identifier of the tunnel on the data center side.
	//
	// example:
	//
	// 47.21.XX.XX
	LocalId *string `json:"LocalId,omitempty" xml:"LocalId,omitempty"`
	// The pre-shared key.
	//
	// example:
	//
	// pgw6dy7d1i8i****
	Psk *string `json:"Psk,omitempty" xml:"Psk,omitempty"`
	// The identifier of the tunnel on the Alibaba Cloud side.
	//
	// example:
	//
	// 47.24.XX.XX
	RemoteId *string `json:"RemoteId,omitempty" xml:"RemoteId,omitempty"`
}

func (s DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIkeConfig) String() string {
	return dara.Prettify(s)
}

func (s DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIkeConfig) GoString() string {
	return s.String()
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIkeConfig) GetIkeAuthAlg() *string {
	return s.IkeAuthAlg
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIkeConfig) GetIkeEncAlg() *string {
	return s.IkeEncAlg
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIkeConfig) GetIkeLifetime() *int64 {
	return s.IkeLifetime
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIkeConfig) GetIkeMode() *string {
	return s.IkeMode
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIkeConfig) GetIkePfs() *string {
	return s.IkePfs
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIkeConfig) GetIkeVersion() *string {
	return s.IkeVersion
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIkeConfig) GetLocalId() *string {
	return s.LocalId
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIkeConfig) GetPsk() *string {
	return s.Psk
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIkeConfig) GetRemoteId() *string {
	return s.RemoteId
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIkeConfig) SetIkeAuthAlg(v string) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIkeConfig {
	s.IkeAuthAlg = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIkeConfig) SetIkeEncAlg(v string) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIkeConfig {
	s.IkeEncAlg = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIkeConfig) SetIkeLifetime(v int64) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIkeConfig {
	s.IkeLifetime = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIkeConfig) SetIkeMode(v string) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIkeConfig {
	s.IkeMode = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIkeConfig) SetIkePfs(v string) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIkeConfig {
	s.IkePfs = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIkeConfig) SetIkeVersion(v string) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIkeConfig {
	s.IkeVersion = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIkeConfig) SetLocalId(v string) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIkeConfig {
	s.LocalId = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIkeConfig) SetPsk(v string) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIkeConfig {
	s.Psk = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIkeConfig) SetRemoteId(v string) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIkeConfig {
	s.RemoteId = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIkeConfig) Validate() error {
	return dara.Validate(s)
}

type DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIpsecConfig struct {
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

func (s DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIpsecConfig) String() string {
	return dara.Prettify(s)
}

func (s DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIpsecConfig) GoString() string {
	return s.String()
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIpsecConfig) GetIpsecAuthAlg() *string {
	return s.IpsecAuthAlg
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIpsecConfig) GetIpsecEncAlg() *string {
	return s.IpsecEncAlg
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIpsecConfig) GetIpsecLifetime() *int64 {
	return s.IpsecLifetime
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIpsecConfig) GetIpsecPfs() *string {
	return s.IpsecPfs
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIpsecConfig) SetIpsecAuthAlg(v string) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIpsecConfig {
	s.IpsecAuthAlg = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIpsecConfig) SetIpsecEncAlg(v string) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIpsecConfig {
	s.IpsecEncAlg = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIpsecConfig) SetIpsecLifetime(v int64) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIpsecConfig {
	s.IpsecLifetime = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIpsecConfig) SetIpsecPfs(v string) *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIpsecConfig {
	s.IpsecPfs = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponseBodyVpnConnectionConfigTunnelsConfigTunnelConfigIpsecConfig) Validate() error {
	return dara.Validate(s)
}
