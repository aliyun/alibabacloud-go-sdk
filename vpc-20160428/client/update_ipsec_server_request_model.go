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
	// The client CIDR block from which an IP address is allocated to the virtual network interface controller (NIC) of the client.
	//
	// >  The client CIDR block must not overlap with the CIDR blocks of the VPC.
	//
	// example:
	//
	// 10.0.0.0/24
	ClientIpPool *string `json:"ClientIpPool,omitempty" xml:"ClientIpPool,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate a value, and you must make sure that each request has a unique token value. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the value of **RequestId*	- as the value of **ClientToken**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to only precheck this request. Valid values:
	//
	// 	- **true**: prechecks the request without modifying the configurations of the IPsec server. The system checks the required parameters, request format, and service limits. If the request fails to pass the precheck, an error code is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// 	- **false**: sends the request. This is the default value. If the request passes the precheck, the system modifies the configurations of the IPsec server.
	//
	// example:
	//
	// false
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to delete the negotiated IPsec tunnel and initiate the negotiation again. Valid values:
	//
	// 	- **true**: immediately initiates negotiations after the configuration is complete.
	//
	// 	- **false**: initiates negotiations when inbound traffic is detected.
	//
	// example:
	//
	// false
	EffectImmediately *bool `json:"EffectImmediately,omitempty" xml:"EffectImmediately,omitempty"`
	// The configuration of Phase 1 negotiations. Valid values:
	//
	// 	- **IkeVersion**: The IKE version. Valid values: **ikev1*	- and **ikev2**.
	//
	// 	- **IkeMode**: The IKE negotiation mode. Default value: **main**.
	//
	// 	- **IkeEncAlg**: the encryption algorithm that is used in Phase 1 negotiation. Default value: **aes**.
	//
	// 	- **IkeAuthAlg**: the authentication algorithm that is used in Phase 1 negotiation. Default value: **sha1**.
	//
	// 	- **IkePfs**: The Diffie-Hellman key exchange algorithm that is used in Phase 1 negotiations. Default value: **group2**.
	//
	// 	- **IkeLifetime**: The SA lifetime determined by Phase 1 negotiations. Unit: seconds. Valid values: **0*	- to **86400**. Default value: **86400**.
	//
	// 	- **LocalId**: The identifier of the IPsec server. Only FQDN and IP address formats are supported.
	//
	// 	- **RemoteId**: the peer identifier. Only FQDN and IP address formats are supported.
	//
	// example:
	//
	// {"IkeVersion":"ikev2","IkeMode":"main","IkeEncAlg":"aes","IkeAuthAlg":"sha1","IkePfs":"group2","IkeLifetime":86400}
	IkeConfig *string `json:"IkeConfig,omitempty" xml:"IkeConfig,omitempty"`
	// The configuration of Phase 2 negotiation. Valid values:
	//
	// 	- **IpsecEncAlg**: the encryption algorithm that is used in Phase 2 negotiation. Default value: **aes**.
	//
	// 	- **IpsecAuthAlg**: the authentication algorithm that is used in Phase 2 negotiation. Default value: **sha1**.
	//
	// 	- **IpsecPfs**: forwards packets of all protocols. The Diffie-Hellman key exchange algorithm that is used in Phase 2 negotiation. Default value: **group2**.
	//
	// 	- **IpsecLifetime**: the SA lifetime determined by Phase 2 negotiation. Unit: seconds. Valid values: **0*	- to **86400**. Default value: **86400**.
	//
	// example:
	//
	// {"IpsecEncAlg":"aes","IpsecAuthAlg":"sha1","IpsecPfs":"group2","IpsecLifetime":86400}
	IpsecConfig *string `json:"IpsecConfig,omitempty" xml:"IpsecConfig,omitempty"`
	// The IPsec server ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// iss-bp1bo3xuvcxo7ixll****
	IpsecServerId *string `json:"IpsecServerId,omitempty" xml:"IpsecServerId,omitempty"`
	// The name of the IPsec server.
	//
	// It must be 1 to 100 characters in length.
	//
	// example:
	//
	// test
	IpsecServerName *string `json:"IpsecServerName,omitempty" xml:"IpsecServerName,omitempty"`
	// The local CIDR blocks, which are the CIDR blocks of the virtual private cloud (VPC) for the client to access.
	//
	// Multiple CIDR blocks are separated with commas (,). Example: 192.168.1.0/24,192.168.2.0/24.
	//
	// example:
	//
	// 192.168.0.0/24,172.17.0.0/16
	LocalSubnet *string `json:"LocalSubnet,omitempty" xml:"LocalSubnet,omitempty"`
	// The pre-shared key.
	//
	// The pre-shared key that is used for authentication between the IPsec server and the client. The key must be 1 to 100 characters in length.
	//
	// You can call [ListIpsecServers](https://help.aliyun.com/document_detail/2794120.html) to query keys generated by the system.
	//
	// > The pre-shared key of the IPsec server key must be the same as that of the client. Otherwise, the connection between the IPsec server and the client cannot be established.
	//
	// example:
	//
	// Cfd123****
	Psk *string `json:"Psk,omitempty" xml:"Psk,omitempty"`
	// Specifies whether to enable pre-shared key authentication. If you set the value to **true**, pre-shared key authentication is enabled.
	//
	// example:
	//
	// true
	PskEnabled *bool `json:"PskEnabled,omitempty" xml:"PskEnabled,omitempty"`
	// The ID of the region where the IPsec server is created.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
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
