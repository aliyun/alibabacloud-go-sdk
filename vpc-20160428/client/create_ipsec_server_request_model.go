// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpsecServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientIpPool(v string) *CreateIpsecServerRequest
	GetClientIpPool() *string
	SetClientToken(v string) *CreateIpsecServerRequest
	GetClientToken() *string
	SetDryRun(v string) *CreateIpsecServerRequest
	GetDryRun() *string
	SetEffectImmediately(v bool) *CreateIpsecServerRequest
	GetEffectImmediately() *bool
	SetIkeConfig(v string) *CreateIpsecServerRequest
	GetIkeConfig() *string
	SetIpSecServerName(v string) *CreateIpsecServerRequest
	GetIpSecServerName() *string
	SetIpsecConfig(v string) *CreateIpsecServerRequest
	GetIpsecConfig() *string
	SetLocalSubnet(v string) *CreateIpsecServerRequest
	GetLocalSubnet() *string
	SetPsk(v string) *CreateIpsecServerRequest
	GetPsk() *string
	SetPskEnabled(v bool) *CreateIpsecServerRequest
	GetPskEnabled() *bool
	SetRegionId(v string) *CreateIpsecServerRequest
	GetRegionId() *string
	SetVpnGatewayId(v string) *CreateIpsecServerRequest
	GetVpnGatewayId() *string
}

type CreateIpsecServerRequest struct {
	// The client CIDR block, which is the address range used to assign IP addresses to virtual network interface controllers (NICs) of clients.
	//
	// > The client CIDR block cannot conflict with the VPC-side CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.0.0.0/24
	ClientIpPool *string `json:"ClientIpPool,omitempty" xml:"ClientIpPool,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// d7d24a21-f4ba-4454-9173-b38****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without creating the IPsec server. The system checks the required parameters, request format, and service limits. If the check fails, the corresponding error message is returned. If the check succeeds, `DryRunOperation` is returned.
	//
	// - **false*	- (default): sends the request. After the request passes the check, the IPsec server is created.
	//
	// example:
	//
	// false
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether the configuration takes effect immediately. Valid values:
	//
	// - **true**: Negotiation starts immediately after the configuration is complete.
	//
	// - **false*	- (default): Negotiation starts when inbound traffic is detected.
	//
	// example:
	//
	// true
	EffectImmediately *bool `json:"EffectImmediately,omitempty" xml:"EffectImmediately,omitempty"`
	// The Phase 1 negotiation parameter settings. Valid values:
	//
	// - **IkeVersion**: The version of the IKE protocol. Valid values: **ikev1*	- and **ikev2**. Default value: **ikev2**.
	//
	// - **IkeMode**: The negotiation pattern of the IKE version. Default value: **main**.
	//
	// - **IkeEncAlg**: The encryption algorithm used in Phase 1 negotiations. Default value: **aes**.
	//
	// - **IkeAuthAlg**: The authentication algorithm used in Phase 1 negotiations. Default value: **sha1**.
	//
	// - **IkePfs**: The Diffie-Hellman key exchange algorithm used in Phase 1 negotiations. Default value: **group2**.
	//
	// - **IkeLifetime**: The epoch of the security association (SA) negotiated in Phase 1. Unit: seconds. Valid values: **0*	- to **86400**. Default value: **86400**.
	//
	// - **LocalId**: The identity of the IPsec server. The FQDN and IP address formats are supported. Default value: the public IP address of the VPN gateway.
	//
	// - **RemoteId**: The identity of the peer. The FQDN and IP address formats are supported. Default value: empty.
	//
	// example:
	//
	// {"IkeVersion":"ikev2","IkeMode":"main","IkeEncAlg":"aes","IkeAuthAlg":"sha1","IkePfs":"group2","IkeLifetime":86400}
	IkeConfig *string `json:"IkeConfig,omitempty" xml:"IkeConfig,omitempty"`
	// The name of the IPsec server.
	//
	// The name must be 1 to 100 characters in length.
	//
	// example:
	//
	// test
	IpSecServerName *string `json:"IpSecServerName,omitempty" xml:"IpSecServerName,omitempty"`
	// The Phase 2 negotiation parameter settings. Valid values:
	//
	// - **IpsecEncAlg**: The encryption algorithm used in Phase 2 negotiations. Default value: **aes**.
	//
	// - **IpsecAuthAlg**: The authentication algorithm used in Phase 2 negotiations. Default value: **sha1**.
	//
	// - **IpsecPfs**: Forward all protocol packets. The Diffie-Hellman key exchange algorithm used in Phase 2 negotiations. Default value: **group2**.
	//
	// - **IpsecLifetime**: The epoch of the SA negotiated in Phase 2. Unit: seconds. Valid values: **0*	- to **86400**. Default value: **86400**.
	//
	// example:
	//
	// {"IpsecEncAlg":"aes","IpsecAuthAlg":"sha1","IpsecPfs":"group2","IpsecLifetime":86400}
	IpsecConfig *string `json:"IpsecConfig,omitempty" xml:"IpsecConfig,omitempty"`
	// The local CIDR block, which is the VPC-side CIDR block that needs to communicate with the client CIDR block.
	//
	// Separate multiple CIDR blocks with commas (,). Example: 192.168.1.0/24,192.168.2.0/24.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.0.0/24
	LocalSubnet *string `json:"LocalSubnet,omitempty" xml:"LocalSubnet,omitempty"`
	// The pre-shared key.
	//
	// The pre-shared key is used for identity authentication between the IPsec server and the client. The key must be 1 to 100 characters in length.
	//
	// If you do not specify a pre-shared key, the system randomly generates a 16-character string as the pre-shared key. You can call the [ListIpsecServers](https://help.aliyun.com/document_detail/2794120.html) operation to query the pre-shared key generated by the system.
	//
	// > The pre-shared key of the IPsec server must be the same as the authentication key of the client. Otherwise, a connection cannot be established between the IPsec server and the client.
	//
	// example:
	//
	// Cfd123****
	Psk *string `json:"Psk,omitempty" xml:"Psk,omitempty"`
	// Specifies whether to enable pre-shared key authentication. Valid values: **true**, which indicates that pre-shared key authentication is enabled.
	//
	// > This parameter is required.
	//
	// example:
	//
	// true
	PskEnabled *bool `json:"PskEnabled,omitempty" xml:"PskEnabled,omitempty"`
	// The region ID of the VPN gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance ID of the VPN gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpn-bp17lofy9fd0dnvzv****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
}

func (s CreateIpsecServerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIpsecServerRequest) GoString() string {
	return s.String()
}

func (s *CreateIpsecServerRequest) GetClientIpPool() *string {
	return s.ClientIpPool
}

func (s *CreateIpsecServerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateIpsecServerRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *CreateIpsecServerRequest) GetEffectImmediately() *bool {
	return s.EffectImmediately
}

func (s *CreateIpsecServerRequest) GetIkeConfig() *string {
	return s.IkeConfig
}

func (s *CreateIpsecServerRequest) GetIpSecServerName() *string {
	return s.IpSecServerName
}

func (s *CreateIpsecServerRequest) GetIpsecConfig() *string {
	return s.IpsecConfig
}

func (s *CreateIpsecServerRequest) GetLocalSubnet() *string {
	return s.LocalSubnet
}

func (s *CreateIpsecServerRequest) GetPsk() *string {
	return s.Psk
}

func (s *CreateIpsecServerRequest) GetPskEnabled() *bool {
	return s.PskEnabled
}

func (s *CreateIpsecServerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateIpsecServerRequest) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *CreateIpsecServerRequest) SetClientIpPool(v string) *CreateIpsecServerRequest {
	s.ClientIpPool = &v
	return s
}

func (s *CreateIpsecServerRequest) SetClientToken(v string) *CreateIpsecServerRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateIpsecServerRequest) SetDryRun(v string) *CreateIpsecServerRequest {
	s.DryRun = &v
	return s
}

func (s *CreateIpsecServerRequest) SetEffectImmediately(v bool) *CreateIpsecServerRequest {
	s.EffectImmediately = &v
	return s
}

func (s *CreateIpsecServerRequest) SetIkeConfig(v string) *CreateIpsecServerRequest {
	s.IkeConfig = &v
	return s
}

func (s *CreateIpsecServerRequest) SetIpSecServerName(v string) *CreateIpsecServerRequest {
	s.IpSecServerName = &v
	return s
}

func (s *CreateIpsecServerRequest) SetIpsecConfig(v string) *CreateIpsecServerRequest {
	s.IpsecConfig = &v
	return s
}

func (s *CreateIpsecServerRequest) SetLocalSubnet(v string) *CreateIpsecServerRequest {
	s.LocalSubnet = &v
	return s
}

func (s *CreateIpsecServerRequest) SetPsk(v string) *CreateIpsecServerRequest {
	s.Psk = &v
	return s
}

func (s *CreateIpsecServerRequest) SetPskEnabled(v bool) *CreateIpsecServerRequest {
	s.PskEnabled = &v
	return s
}

func (s *CreateIpsecServerRequest) SetRegionId(v string) *CreateIpsecServerRequest {
	s.RegionId = &v
	return s
}

func (s *CreateIpsecServerRequest) SetVpnGatewayId(v string) *CreateIpsecServerRequest {
	s.VpnGatewayId = &v
	return s
}

func (s *CreateIpsecServerRequest) Validate() error {
	return dara.Validate(s)
}
