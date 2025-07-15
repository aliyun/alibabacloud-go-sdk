// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpsecServersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpsecServers(v []*ListIpsecServersResponseBodyIpsecServers) *ListIpsecServersResponseBody
	GetIpsecServers() []*ListIpsecServersResponseBodyIpsecServers
	SetMaxResults(v int32) *ListIpsecServersResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListIpsecServersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListIpsecServersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListIpsecServersResponseBody
	GetTotalCount() *int32
}

type ListIpsecServersResponseBody struct {
	// The list of IPsec servers.
	IpsecServers []*ListIpsecServersResponseBodyIpsecServers `json:"IpsecServers,omitempty" xml:"IpsecServers,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If no value is returned for **NextToken**, no next queries are sent.
	//
	// 	- If a value is returned for **NextToken**, the value can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B457
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIpsecServersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIpsecServersResponseBody) GoString() string {
	return s.String()
}

func (s *ListIpsecServersResponseBody) GetIpsecServers() []*ListIpsecServersResponseBodyIpsecServers {
	return s.IpsecServers
}

func (s *ListIpsecServersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListIpsecServersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIpsecServersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIpsecServersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListIpsecServersResponseBody) SetIpsecServers(v []*ListIpsecServersResponseBodyIpsecServers) *ListIpsecServersResponseBody {
	s.IpsecServers = v
	return s
}

func (s *ListIpsecServersResponseBody) SetMaxResults(v int32) *ListIpsecServersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListIpsecServersResponseBody) SetNextToken(v string) *ListIpsecServersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIpsecServersResponseBody) SetRequestId(v string) *ListIpsecServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIpsecServersResponseBody) SetTotalCount(v int32) *ListIpsecServersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListIpsecServersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListIpsecServersResponseBodyIpsecServers struct {
	// The client CIDR block. It refers to the CIDR block that is allocated to the virtual interface of the client.
	//
	// example:
	//
	// 10.0.0.0/24
	ClientIpPool *string `json:"ClientIpPool,omitempty" xml:"ClientIpPool,omitempty"`
	// The time when the IPsec server was created.
	//
	// T is used as a delimiter. Z indicates that the time is in UTC.
	//
	// example:
	//
	// 2018-12-03T10:11:55Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Indicates whether the current IPsec tunnel is deleted and negotiations are reinitiated. Valid values:
	//
	// 	- **true**: immediately initiates negotiations after the configuration is completed.
	//
	// 	- **false**: initiates negotiations when inbound traffic is detected.
	//
	// example:
	//
	// false
	EffectImmediately *bool `json:"EffectImmediately,omitempty" xml:"EffectImmediately,omitempty"`
	// The ID of the IDaaS instance.
	//
	// example:
	//
	// idaas-cn-hangzhou-****
	IDaaSInstanceId *string `json:"IDaaSInstanceId,omitempty" xml:"IDaaSInstanceId,omitempty"`
	// The configurations of Phase 1 negotiations.
	IkeConfig *ListIpsecServersResponseBodyIpsecServersIkeConfig `json:"IkeConfig,omitempty" xml:"IkeConfig,omitempty" type:"Struct"`
	// The public IP address of the VPN gateway.
	//
	// example:
	//
	// 47.22.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The configurations of Phase 2 negotiations.
	IpsecConfig *ListIpsecServersResponseBodyIpsecServersIpsecConfig `json:"IpsecConfig,omitempty" xml:"IpsecConfig,omitempty" type:"Struct"`
	// The IPsec server ID.
	//
	// example:
	//
	// iss-bp1bo3xuvcxo7ixll****
	IpsecServerId *string `json:"IpsecServerId,omitempty" xml:"IpsecServerId,omitempty"`
	// The name of the IPsec server.
	//
	// example:
	//
	// test
	IpsecServerName *string `json:"IpsecServerName,omitempty" xml:"IpsecServerName,omitempty"`
	// The local CIDR blocks, which refer to the CIDR blocks on the virtual private cloud (VPC) side.
	//
	// example:
	//
	// 192.168.0.0/16,172.17.0.0/16
	LocalSubnet *string `json:"LocalSubnet,omitempty" xml:"LocalSubnet,omitempty"`
	// The number of SSL-VPN connections supported by the VPN gateway.
	//
	// >  The number of SSL-VPN connections specified in this parameter includes both SSL-VPN and IPsec-VPN connections. For example, you have five SSL-VPN connections and three SSL clients occupy three SSL-VPN connections. In this case, two clients can connect to the IPsec server.
	//
	// example:
	//
	// 5
	MaxConnections *int32 `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	// Indicates whether two-factor authentication is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**: The feature is disabled.
	//
	// example:
	//
	// true
	MultiFactorAuthEnabled *bool `json:"MultiFactorAuthEnabled,omitempty" xml:"MultiFactorAuthEnabled,omitempty"`
	// The number of clients that are connected to the IPsec server.
	//
	// example:
	//
	// 1
	OnlineClientCount *int32 `json:"OnlineClientCount,omitempty" xml:"OnlineClientCount,omitempty"`
	// The pre-shared key.
	//
	// example:
	//
	// pgw6dy7d****
	Psk *string `json:"Psk,omitempty" xml:"Psk,omitempty"`
	// Indicates whether pre-shared key authentication is enabled. Only **true*	- may be returned, which indicates that pre-shared key authentication is enabled.
	//
	// example:
	//
	// true
	PskEnabled *bool `json:"PskEnabled,omitempty" xml:"PskEnabled,omitempty"`
	// The ID of the region where the IPsec server is created.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the IPsec server belongs.
	//
	// You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation to query the resource group information.
	//
	// example:
	//
	// rg-acfmzs372yg****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the VPN gateway.
	//
	// example:
	//
	// vpn-bp1q8bgx4xnkm2ogj****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
}

func (s ListIpsecServersResponseBodyIpsecServers) String() string {
	return dara.Prettify(s)
}

func (s ListIpsecServersResponseBodyIpsecServers) GoString() string {
	return s.String()
}

func (s *ListIpsecServersResponseBodyIpsecServers) GetClientIpPool() *string {
	return s.ClientIpPool
}

func (s *ListIpsecServersResponseBodyIpsecServers) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListIpsecServersResponseBodyIpsecServers) GetEffectImmediately() *bool {
	return s.EffectImmediately
}

func (s *ListIpsecServersResponseBodyIpsecServers) GetIDaaSInstanceId() *string {
	return s.IDaaSInstanceId
}

func (s *ListIpsecServersResponseBodyIpsecServers) GetIkeConfig() *ListIpsecServersResponseBodyIpsecServersIkeConfig {
	return s.IkeConfig
}

func (s *ListIpsecServersResponseBodyIpsecServers) GetInternetIp() *string {
	return s.InternetIp
}

func (s *ListIpsecServersResponseBodyIpsecServers) GetIpsecConfig() *ListIpsecServersResponseBodyIpsecServersIpsecConfig {
	return s.IpsecConfig
}

func (s *ListIpsecServersResponseBodyIpsecServers) GetIpsecServerId() *string {
	return s.IpsecServerId
}

func (s *ListIpsecServersResponseBodyIpsecServers) GetIpsecServerName() *string {
	return s.IpsecServerName
}

func (s *ListIpsecServersResponseBodyIpsecServers) GetLocalSubnet() *string {
	return s.LocalSubnet
}

func (s *ListIpsecServersResponseBodyIpsecServers) GetMaxConnections() *int32 {
	return s.MaxConnections
}

func (s *ListIpsecServersResponseBodyIpsecServers) GetMultiFactorAuthEnabled() *bool {
	return s.MultiFactorAuthEnabled
}

func (s *ListIpsecServersResponseBodyIpsecServers) GetOnlineClientCount() *int32 {
	return s.OnlineClientCount
}

func (s *ListIpsecServersResponseBodyIpsecServers) GetPsk() *string {
	return s.Psk
}

func (s *ListIpsecServersResponseBodyIpsecServers) GetPskEnabled() *bool {
	return s.PskEnabled
}

func (s *ListIpsecServersResponseBodyIpsecServers) GetRegionId() *string {
	return s.RegionId
}

func (s *ListIpsecServersResponseBodyIpsecServers) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListIpsecServersResponseBodyIpsecServers) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *ListIpsecServersResponseBodyIpsecServers) SetClientIpPool(v string) *ListIpsecServersResponseBodyIpsecServers {
	s.ClientIpPool = &v
	return s
}

func (s *ListIpsecServersResponseBodyIpsecServers) SetCreationTime(v string) *ListIpsecServersResponseBodyIpsecServers {
	s.CreationTime = &v
	return s
}

func (s *ListIpsecServersResponseBodyIpsecServers) SetEffectImmediately(v bool) *ListIpsecServersResponseBodyIpsecServers {
	s.EffectImmediately = &v
	return s
}

func (s *ListIpsecServersResponseBodyIpsecServers) SetIDaaSInstanceId(v string) *ListIpsecServersResponseBodyIpsecServers {
	s.IDaaSInstanceId = &v
	return s
}

func (s *ListIpsecServersResponseBodyIpsecServers) SetIkeConfig(v *ListIpsecServersResponseBodyIpsecServersIkeConfig) *ListIpsecServersResponseBodyIpsecServers {
	s.IkeConfig = v
	return s
}

func (s *ListIpsecServersResponseBodyIpsecServers) SetInternetIp(v string) *ListIpsecServersResponseBodyIpsecServers {
	s.InternetIp = &v
	return s
}

func (s *ListIpsecServersResponseBodyIpsecServers) SetIpsecConfig(v *ListIpsecServersResponseBodyIpsecServersIpsecConfig) *ListIpsecServersResponseBodyIpsecServers {
	s.IpsecConfig = v
	return s
}

func (s *ListIpsecServersResponseBodyIpsecServers) SetIpsecServerId(v string) *ListIpsecServersResponseBodyIpsecServers {
	s.IpsecServerId = &v
	return s
}

func (s *ListIpsecServersResponseBodyIpsecServers) SetIpsecServerName(v string) *ListIpsecServersResponseBodyIpsecServers {
	s.IpsecServerName = &v
	return s
}

func (s *ListIpsecServersResponseBodyIpsecServers) SetLocalSubnet(v string) *ListIpsecServersResponseBodyIpsecServers {
	s.LocalSubnet = &v
	return s
}

func (s *ListIpsecServersResponseBodyIpsecServers) SetMaxConnections(v int32) *ListIpsecServersResponseBodyIpsecServers {
	s.MaxConnections = &v
	return s
}

func (s *ListIpsecServersResponseBodyIpsecServers) SetMultiFactorAuthEnabled(v bool) *ListIpsecServersResponseBodyIpsecServers {
	s.MultiFactorAuthEnabled = &v
	return s
}

func (s *ListIpsecServersResponseBodyIpsecServers) SetOnlineClientCount(v int32) *ListIpsecServersResponseBodyIpsecServers {
	s.OnlineClientCount = &v
	return s
}

func (s *ListIpsecServersResponseBodyIpsecServers) SetPsk(v string) *ListIpsecServersResponseBodyIpsecServers {
	s.Psk = &v
	return s
}

func (s *ListIpsecServersResponseBodyIpsecServers) SetPskEnabled(v bool) *ListIpsecServersResponseBodyIpsecServers {
	s.PskEnabled = &v
	return s
}

func (s *ListIpsecServersResponseBodyIpsecServers) SetRegionId(v string) *ListIpsecServersResponseBodyIpsecServers {
	s.RegionId = &v
	return s
}

func (s *ListIpsecServersResponseBodyIpsecServers) SetResourceGroupId(v string) *ListIpsecServersResponseBodyIpsecServers {
	s.ResourceGroupId = &v
	return s
}

func (s *ListIpsecServersResponseBodyIpsecServers) SetVpnGatewayId(v string) *ListIpsecServersResponseBodyIpsecServers {
	s.VpnGatewayId = &v
	return s
}

func (s *ListIpsecServersResponseBodyIpsecServers) Validate() error {
	return dara.Validate(s)
}

type ListIpsecServersResponseBodyIpsecServersIkeConfig struct {
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
	// The IKE negotiation mode. Valid values:
	//
	// **main**: This mode offers higher security during negotiations.
	//
	// example:
	//
	// main
	IkeMode *string `json:"IkeMode,omitempty" xml:"IkeMode,omitempty"`
	// The Diffie-Hellman key exchange algorithm.
	//
	// example:
	//
	// group2
	IkePfs *string `json:"IkePfs,omitempty" xml:"IkePfs,omitempty"`
	// The IKE version.
	//
	// example:
	//
	// ikev2
	IkeVersion *string `json:"IkeVersion,omitempty" xml:"IkeVersion,omitempty"`
	// The ID of the IPsec server. The default value is the public IP address of the VPN gateway. Both FQDNs and IP addresses are supported.
	//
	// example:
	//
	// 116.64.XX.XX
	LocalId *string `json:"LocalId,omitempty" xml:"LocalId,omitempty"`
	// The identifier of the customer gateway. Both fully qualified domain names (FQDNs) and IP addresses are supported. By default, this parameter is empty.
	//
	// example:
	//
	// 139.67.XX.XX
	RemoteId *string `json:"RemoteId,omitempty" xml:"RemoteId,omitempty"`
}

func (s ListIpsecServersResponseBodyIpsecServersIkeConfig) String() string {
	return dara.Prettify(s)
}

func (s ListIpsecServersResponseBodyIpsecServersIkeConfig) GoString() string {
	return s.String()
}

func (s *ListIpsecServersResponseBodyIpsecServersIkeConfig) GetIkeAuthAlg() *string {
	return s.IkeAuthAlg
}

func (s *ListIpsecServersResponseBodyIpsecServersIkeConfig) GetIkeEncAlg() *string {
	return s.IkeEncAlg
}

func (s *ListIpsecServersResponseBodyIpsecServersIkeConfig) GetIkeLifetime() *int64 {
	return s.IkeLifetime
}

func (s *ListIpsecServersResponseBodyIpsecServersIkeConfig) GetIkeMode() *string {
	return s.IkeMode
}

func (s *ListIpsecServersResponseBodyIpsecServersIkeConfig) GetIkePfs() *string {
	return s.IkePfs
}

func (s *ListIpsecServersResponseBodyIpsecServersIkeConfig) GetIkeVersion() *string {
	return s.IkeVersion
}

func (s *ListIpsecServersResponseBodyIpsecServersIkeConfig) GetLocalId() *string {
	return s.LocalId
}

func (s *ListIpsecServersResponseBodyIpsecServersIkeConfig) GetRemoteId() *string {
	return s.RemoteId
}

func (s *ListIpsecServersResponseBodyIpsecServersIkeConfig) SetIkeAuthAlg(v string) *ListIpsecServersResponseBodyIpsecServersIkeConfig {
	s.IkeAuthAlg = &v
	return s
}

func (s *ListIpsecServersResponseBodyIpsecServersIkeConfig) SetIkeEncAlg(v string) *ListIpsecServersResponseBodyIpsecServersIkeConfig {
	s.IkeEncAlg = &v
	return s
}

func (s *ListIpsecServersResponseBodyIpsecServersIkeConfig) SetIkeLifetime(v int64) *ListIpsecServersResponseBodyIpsecServersIkeConfig {
	s.IkeLifetime = &v
	return s
}

func (s *ListIpsecServersResponseBodyIpsecServersIkeConfig) SetIkeMode(v string) *ListIpsecServersResponseBodyIpsecServersIkeConfig {
	s.IkeMode = &v
	return s
}

func (s *ListIpsecServersResponseBodyIpsecServersIkeConfig) SetIkePfs(v string) *ListIpsecServersResponseBodyIpsecServersIkeConfig {
	s.IkePfs = &v
	return s
}

func (s *ListIpsecServersResponseBodyIpsecServersIkeConfig) SetIkeVersion(v string) *ListIpsecServersResponseBodyIpsecServersIkeConfig {
	s.IkeVersion = &v
	return s
}

func (s *ListIpsecServersResponseBodyIpsecServersIkeConfig) SetLocalId(v string) *ListIpsecServersResponseBodyIpsecServersIkeConfig {
	s.LocalId = &v
	return s
}

func (s *ListIpsecServersResponseBodyIpsecServersIkeConfig) SetRemoteId(v string) *ListIpsecServersResponseBodyIpsecServersIkeConfig {
	s.RemoteId = &v
	return s
}

func (s *ListIpsecServersResponseBodyIpsecServersIkeConfig) Validate() error {
	return dara.Validate(s)
}

type ListIpsecServersResponseBodyIpsecServersIpsecConfig struct {
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
	// The Diffie-Hellman key exchange algorithm.
	//
	// example:
	//
	// group2
	IpsecPfs *string `json:"IpsecPfs,omitempty" xml:"IpsecPfs,omitempty"`
}

func (s ListIpsecServersResponseBodyIpsecServersIpsecConfig) String() string {
	return dara.Prettify(s)
}

func (s ListIpsecServersResponseBodyIpsecServersIpsecConfig) GoString() string {
	return s.String()
}

func (s *ListIpsecServersResponseBodyIpsecServersIpsecConfig) GetIpsecAuthAlg() *string {
	return s.IpsecAuthAlg
}

func (s *ListIpsecServersResponseBodyIpsecServersIpsecConfig) GetIpsecEncAlg() *string {
	return s.IpsecEncAlg
}

func (s *ListIpsecServersResponseBodyIpsecServersIpsecConfig) GetIpsecLifetime() *int64 {
	return s.IpsecLifetime
}

func (s *ListIpsecServersResponseBodyIpsecServersIpsecConfig) GetIpsecPfs() *string {
	return s.IpsecPfs
}

func (s *ListIpsecServersResponseBodyIpsecServersIpsecConfig) SetIpsecAuthAlg(v string) *ListIpsecServersResponseBodyIpsecServersIpsecConfig {
	s.IpsecAuthAlg = &v
	return s
}

func (s *ListIpsecServersResponseBodyIpsecServersIpsecConfig) SetIpsecEncAlg(v string) *ListIpsecServersResponseBodyIpsecServersIpsecConfig {
	s.IpsecEncAlg = &v
	return s
}

func (s *ListIpsecServersResponseBodyIpsecServersIpsecConfig) SetIpsecLifetime(v int64) *ListIpsecServersResponseBodyIpsecServersIpsecConfig {
	s.IpsecLifetime = &v
	return s
}

func (s *ListIpsecServersResponseBodyIpsecServersIpsecConfig) SetIpsecPfs(v string) *ListIpsecServersResponseBodyIpsecServersIpsecConfig {
	s.IpsecPfs = &v
	return s
}

func (s *ListIpsecServersResponseBodyIpsecServersIpsecConfig) Validate() error {
	return dara.Validate(s)
}
