// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIpsecServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientIpPool(v string) *UpdateIpsecServerRequest
	GetClientIpPool() *string
	SetClientToken(v string) *UpdateIpsecServerRequest
	GetClientToken() *string
	SetDryRun(v string) *UpdateIpsecServerRequest
	GetDryRun() *string
	SetEffectImmediately(v bool) *UpdateIpsecServerRequest
	GetEffectImmediately() *bool
	SetIkeConfig(v string) *UpdateIpsecServerRequest
	GetIkeConfig() *string
	SetIpsecConfig(v string) *UpdateIpsecServerRequest
	GetIpsecConfig() *string
	SetIpsecServerId(v string) *UpdateIpsecServerRequest
	GetIpsecServerId() *string
	SetIpsecServerName(v string) *UpdateIpsecServerRequest
	GetIpsecServerName() *string
	SetLocalSubnet(v string) *UpdateIpsecServerRequest
	GetLocalSubnet() *string
	SetPsk(v string) *UpdateIpsecServerRequest
	GetPsk() *string
	SetPskEnabled(v bool) *UpdateIpsecServerRequest
	GetPskEnabled() *bool
	SetRegionId(v string) *UpdateIpsecServerRequest
	GetRegionId() *string
}

type UpdateIpsecServerRequest struct {
	// The client CIDR block, which is the CIDR block from which IP addresses are assigned to virtual network interface controllers (NICs) of the client.
	//
	// > The client CIDR block cannot conflict with the VPC CIDR block.
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
	// e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without modifying the IPsec server configuration. The system checks the required parameters, request format, and service limits. If the check fails, the corresponding error message is returned. If the check succeeds, `DryRunOperation` is returned.
	//
	// - **false*	- (default): performs a dry run and sends the request. After the check succeeds, the IPsec server configuration is modified.
	//
	// example:
	//
	// false
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to delete the currently negotiated IPsec tunnel and re-initiate negotiation. Valid values:
	//
	// - **true**: Negotiation is initiated immediately after the configuration is complete.
	//
	// - **false**: Negotiation is initiated when inbound traffic is detected.
	//
	// example:
	//
	// false
	EffectImmediately *bool `json:"EffectImmediately,omitempty" xml:"EffectImmediately,omitempty"`
	// The Phase 1 negotiation parameter settings. Valid values:
	//
	// - **IkeVersion**: The version of the IKE protocol. Valid values: **ikev1*	- or **ikev2**.
	//
	// - **IkeMode**: The negotiation pattern of the IKE version. Default value: **main**.
	//
	// - **IkeEncAlg**: The encryption algorithm used in Phase 1 negotiations. Default value: **aes**.
	//
	// - **IkeAuthAlg**: The authentication algorithm used in Phase 1 negotiations. Default value: **sha1**.
	//
	// - **IkePfs**: The Diffie-Hellman key exchange algorithm used in Phase 1 negotiations. Default value: **group2**.
	//
	// - **IkeLifetime**: The lifetime of the SA negotiated in Phase 1. Unit: seconds. Valid values: **0*	- to **86400**. Default value: **86400**. The parameter specifies the SA epoch.
	//
	// - **LocalId**: The identity of the IPsec server. FQDN and IP address formats are supported.
	//
	// - **RemoteId**: The identity of the peer. FQDN and IP address formats are supported.
	//
	// example:
	//
	// {"IkeVersion":"ikev2","IkeMode":"main","IkeEncAlg":"aes","IkeAuthAlg":"sha1","IkePfs":"group2","IkeLifetime":86400}
	IkeConfig *string `json:"IkeConfig,omitempty" xml:"IkeConfig,omitempty"`
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
	// The ID of the IPsec server.
	//
	// This parameter is required.
	//
	// example:
	//
	// iss-bp1bo3xuvcxo7ixll****
	IpsecServerId *string `json:"IpsecServerId,omitempty" xml:"IpsecServerId,omitempty"`
	// The name of the IPsec server.
	//
	// The name must be 1 to 100 characters in length.
	//
	// example:
	//
	// test
	IpsecServerName *string `json:"IpsecServerName,omitempty" xml:"IpsecServerName,omitempty"`
	// The local CIDR block, which is the CIDR block on the VPC side that needs to communicate with the client CIDR block.
	//
	// Separate multiple CIDR blocks with commas (,). Example: 192.168.1.0/24,192.168.2.0/24.
	//
	// example:
	//
	// 192.168.0.0/24,172.17.0.0/16
	LocalSubnet *string `json:"LocalSubnet,omitempty" xml:"LocalSubnet,omitempty"`
	// The pre-shared key.
	//
	// The pre-shared key is used for identity authentication between the IPsec server and the client. The key must be 1 to 100 characters in length.
	//
	// You can call the [ListIpsecServers](https://help.aliyun.com/document_detail/2794120.html) operation to query the key generated by the system.
	//
	// > The pre-shared key of the IPsec server must be the same as the authentication key of the client. Otherwise, a connection cannot be established between the IPsec server and the client.
	//
	// example:
	//
	// Cfd123****
	Psk *string `json:"Psk,omitempty" xml:"Psk,omitempty"`
	// Specifies whether to enable pre-shared key authentication. Valid values: **true**, which indicates that pre-shared key authentication is enabled.
	//
	// example:
	//
	// true
	PskEnabled *bool `json:"PskEnabled,omitempty" xml:"PskEnabled,omitempty"`
	// The region ID of the IPsec server.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateIpsecServerRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIpsecServerRequest) GoString() string {
	return s.String()
}

func (s *UpdateIpsecServerRequest) GetClientIpPool() *string {
	return s.ClientIpPool
}

func (s *UpdateIpsecServerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateIpsecServerRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *UpdateIpsecServerRequest) GetEffectImmediately() *bool {
	return s.EffectImmediately
}

func (s *UpdateIpsecServerRequest) GetIkeConfig() *string {
	return s.IkeConfig
}

func (s *UpdateIpsecServerRequest) GetIpsecConfig() *string {
	return s.IpsecConfig
}

func (s *UpdateIpsecServerRequest) GetIpsecServerId() *string {
	return s.IpsecServerId
}

func (s *UpdateIpsecServerRequest) GetIpsecServerName() *string {
	return s.IpsecServerName
}

func (s *UpdateIpsecServerRequest) GetLocalSubnet() *string {
	return s.LocalSubnet
}

func (s *UpdateIpsecServerRequest) GetPsk() *string {
	return s.Psk
}

func (s *UpdateIpsecServerRequest) GetPskEnabled() *bool {
	return s.PskEnabled
}

func (s *UpdateIpsecServerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateIpsecServerRequest) SetClientIpPool(v string) *UpdateIpsecServerRequest {
	s.ClientIpPool = &v
	return s
}

func (s *UpdateIpsecServerRequest) SetClientToken(v string) *UpdateIpsecServerRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateIpsecServerRequest) SetDryRun(v string) *UpdateIpsecServerRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateIpsecServerRequest) SetEffectImmediately(v bool) *UpdateIpsecServerRequest {
	s.EffectImmediately = &v
	return s
}

func (s *UpdateIpsecServerRequest) SetIkeConfig(v string) *UpdateIpsecServerRequest {
	s.IkeConfig = &v
	return s
}

func (s *UpdateIpsecServerRequest) SetIpsecConfig(v string) *UpdateIpsecServerRequest {
	s.IpsecConfig = &v
	return s
}

func (s *UpdateIpsecServerRequest) SetIpsecServerId(v string) *UpdateIpsecServerRequest {
	s.IpsecServerId = &v
	return s
}

func (s *UpdateIpsecServerRequest) SetIpsecServerName(v string) *UpdateIpsecServerRequest {
	s.IpsecServerName = &v
	return s
}

func (s *UpdateIpsecServerRequest) SetLocalSubnet(v string) *UpdateIpsecServerRequest {
	s.LocalSubnet = &v
	return s
}

func (s *UpdateIpsecServerRequest) SetPsk(v string) *UpdateIpsecServerRequest {
	s.Psk = &v
	return s
}

func (s *UpdateIpsecServerRequest) SetPskEnabled(v bool) *UpdateIpsecServerRequest {
	s.PskEnabled = &v
	return s
}

func (s *UpdateIpsecServerRequest) SetRegionId(v string) *UpdateIpsecServerRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateIpsecServerRequest) Validate() error {
	return dara.Validate(s)
}
