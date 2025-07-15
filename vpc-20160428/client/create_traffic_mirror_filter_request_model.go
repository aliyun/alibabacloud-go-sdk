// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrafficMirrorFilterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateTrafficMirrorFilterRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateTrafficMirrorFilterRequest
	GetDryRun() *bool
	SetEgressRules(v []*CreateTrafficMirrorFilterRequestEgressRules) *CreateTrafficMirrorFilterRequest
	GetEgressRules() []*CreateTrafficMirrorFilterRequestEgressRules
	SetIngressRules(v []*CreateTrafficMirrorFilterRequestIngressRules) *CreateTrafficMirrorFilterRequest
	GetIngressRules() []*CreateTrafficMirrorFilterRequestIngressRules
	SetOwnerAccount(v string) *CreateTrafficMirrorFilterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateTrafficMirrorFilterRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateTrafficMirrorFilterRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateTrafficMirrorFilterRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateTrafficMirrorFilterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateTrafficMirrorFilterRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreateTrafficMirrorFilterRequestTag) *CreateTrafficMirrorFilterRequest
	GetTag() []*CreateTrafficMirrorFilterRequestTag
	SetTrafficMirrorFilterDescription(v string) *CreateTrafficMirrorFilterRequest
	GetTrafficMirrorFilterDescription() *string
	SetTrafficMirrorFilterName(v string) *CreateTrafficMirrorFilterRequest
	GetTrafficMirrorFilterName() *string
}

type CreateTrafficMirrorFilterRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that the value is unique among all requests. The client token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, the system uses **RequestId*	- as **ClientToken**. **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// 	- **true**: performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false**: performs a dry run and sends the request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed. This is the default value.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The information about the outbound rules.
	EgressRules []*CreateTrafficMirrorFilterRequestEgressRules `json:"EgressRules,omitempty" xml:"EgressRules,omitempty" type:"Repeated"`
	// The information about inbound rules.
	IngressRules []*CreateTrafficMirrorFilterRequestIngressRules `json:"IngressRules,omitempty" xml:"IngressRules,omitempty" type:"Repeated"`
	OwnerAccount *string                                         `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64                                          `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to which the mirrored traffic belongs.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list. For more information about regions that support traffic mirror, see [Overview of traffic mirror](https://help.aliyun.com/document_detail/207513.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hongkong
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the mirrored traffic belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4ph****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag of the resource.
	Tag []*CreateTrafficMirrorFilterRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The description of the filter.
	//
	// The description must be 1 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// this is a filter.
	TrafficMirrorFilterDescription *string `json:"TrafficMirrorFilterDescription,omitempty" xml:"TrafficMirrorFilterDescription,omitempty"`
	// The name of the filter.
	//
	// The name must be 1 to 128 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// abc
	TrafficMirrorFilterName *string `json:"TrafficMirrorFilterName,omitempty" xml:"TrafficMirrorFilterName,omitempty"`
}

func (s CreateTrafficMirrorFilterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTrafficMirrorFilterRequest) GoString() string {
	return s.String()
}

func (s *CreateTrafficMirrorFilterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTrafficMirrorFilterRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateTrafficMirrorFilterRequest) GetEgressRules() []*CreateTrafficMirrorFilterRequestEgressRules {
	return s.EgressRules
}

func (s *CreateTrafficMirrorFilterRequest) GetIngressRules() []*CreateTrafficMirrorFilterRequestIngressRules {
	return s.IngressRules
}

func (s *CreateTrafficMirrorFilterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateTrafficMirrorFilterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateTrafficMirrorFilterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTrafficMirrorFilterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateTrafficMirrorFilterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateTrafficMirrorFilterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateTrafficMirrorFilterRequest) GetTag() []*CreateTrafficMirrorFilterRequestTag {
	return s.Tag
}

func (s *CreateTrafficMirrorFilterRequest) GetTrafficMirrorFilterDescription() *string {
	return s.TrafficMirrorFilterDescription
}

func (s *CreateTrafficMirrorFilterRequest) GetTrafficMirrorFilterName() *string {
	return s.TrafficMirrorFilterName
}

func (s *CreateTrafficMirrorFilterRequest) SetClientToken(v string) *CreateTrafficMirrorFilterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTrafficMirrorFilterRequest) SetDryRun(v bool) *CreateTrafficMirrorFilterRequest {
	s.DryRun = &v
	return s
}

func (s *CreateTrafficMirrorFilterRequest) SetEgressRules(v []*CreateTrafficMirrorFilterRequestEgressRules) *CreateTrafficMirrorFilterRequest {
	s.EgressRules = v
	return s
}

func (s *CreateTrafficMirrorFilterRequest) SetIngressRules(v []*CreateTrafficMirrorFilterRequestIngressRules) *CreateTrafficMirrorFilterRequest {
	s.IngressRules = v
	return s
}

func (s *CreateTrafficMirrorFilterRequest) SetOwnerAccount(v string) *CreateTrafficMirrorFilterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTrafficMirrorFilterRequest) SetOwnerId(v int64) *CreateTrafficMirrorFilterRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTrafficMirrorFilterRequest) SetRegionId(v string) *CreateTrafficMirrorFilterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTrafficMirrorFilterRequest) SetResourceGroupId(v string) *CreateTrafficMirrorFilterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateTrafficMirrorFilterRequest) SetResourceOwnerAccount(v string) *CreateTrafficMirrorFilterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTrafficMirrorFilterRequest) SetResourceOwnerId(v int64) *CreateTrafficMirrorFilterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTrafficMirrorFilterRequest) SetTag(v []*CreateTrafficMirrorFilterRequestTag) *CreateTrafficMirrorFilterRequest {
	s.Tag = v
	return s
}

func (s *CreateTrafficMirrorFilterRequest) SetTrafficMirrorFilterDescription(v string) *CreateTrafficMirrorFilterRequest {
	s.TrafficMirrorFilterDescription = &v
	return s
}

func (s *CreateTrafficMirrorFilterRequest) SetTrafficMirrorFilterName(v string) *CreateTrafficMirrorFilterRequest {
	s.TrafficMirrorFilterName = &v
	return s
}

func (s *CreateTrafficMirrorFilterRequest) Validate() error {
	return dara.Validate(s)
}

type CreateTrafficMirrorFilterRequestEgressRules struct {
	// The collection policy of the outbound rule. Valid values:
	//
	// 	- **accept**: collects the network traffic.
	//
	// 	- **drop**: does not collect the network traffic.
	//
	// example:
	//
	// accept
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The destination CIDR block of the outbound traffic.
	//
	// example:
	//
	// 10.0.0.0/24
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The destination port range of the outbound traffic. Valid values for a port: **1*	- to **65535**. Separate the first port and the last port with a forward slash (/). Examples: **1/200*	- and **80/80**. You cannot set this parameter to only -1/-1, which specifies all ports.
	//
	// >  If **EgressRules.N.Protocol*	- is set to **ALL*	- or **ICMP**, you do not need to specify this parameter. This indicates that all ports are available.
	//
	// example:
	//
	// 22/40
	DestinationPortRange *string `json:"DestinationPortRange,omitempty" xml:"DestinationPortRange,omitempty"`
	// The IP version of the instance. The following value may be returned:
	//
	// 	- **IPv4**: IPv4
	//
	// 	- **IPv6**: IPv6
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The priority of the outbound rule. A smaller value indicates a higher priority. The maximum value of **N*	- is **10**. You can configure up to 10 outbound rules for a filter.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The type of the protocol that is used by the outbound traffic that you want to mirror. Valid values:
	//
	// 	- **ALL**: all protocols
	//
	// 	- **ICMP**: Internet Control Message Protocol.
	//
	// 	- **TCP**: Transmission Control Protocol.
	//
	// 	- **UDP**: User Datagram Protocol.
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The source CIDR block of the outbound traffic.
	//
	// example:
	//
	// 10.0.0.0/24
	SourceCidrBlock *string `json:"SourceCidrBlock,omitempty" xml:"SourceCidrBlock,omitempty"`
	// The source port range of the outbound traffic. Valid values: **1*	- to **65535**. Separate the first port and the last port with a forward slash (/). Examples: **1/200*	- and **80/80**. You cannot set this parameter to only -1/-1, which specifies all ports.
	//
	// >  If **EgressRules.N.Protocol*	- is set to **ALL*	- or **ICMP**, you do not need to specify this parameter. This indicates that all ports are available.
	//
	// example:
	//
	// 22/40
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
}

func (s CreateTrafficMirrorFilterRequestEgressRules) String() string {
	return dara.Prettify(s)
}

func (s CreateTrafficMirrorFilterRequestEgressRules) GoString() string {
	return s.String()
}

func (s *CreateTrafficMirrorFilterRequestEgressRules) GetAction() *string {
	return s.Action
}

func (s *CreateTrafficMirrorFilterRequestEgressRules) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *CreateTrafficMirrorFilterRequestEgressRules) GetDestinationPortRange() *string {
	return s.DestinationPortRange
}

func (s *CreateTrafficMirrorFilterRequestEgressRules) GetIpVersion() *string {
	return s.IpVersion
}

func (s *CreateTrafficMirrorFilterRequestEgressRules) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateTrafficMirrorFilterRequestEgressRules) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateTrafficMirrorFilterRequestEgressRules) GetSourceCidrBlock() *string {
	return s.SourceCidrBlock
}

func (s *CreateTrafficMirrorFilterRequestEgressRules) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *CreateTrafficMirrorFilterRequestEgressRules) SetAction(v string) *CreateTrafficMirrorFilterRequestEgressRules {
	s.Action = &v
	return s
}

func (s *CreateTrafficMirrorFilterRequestEgressRules) SetDestinationCidrBlock(v string) *CreateTrafficMirrorFilterRequestEgressRules {
	s.DestinationCidrBlock = &v
	return s
}

func (s *CreateTrafficMirrorFilterRequestEgressRules) SetDestinationPortRange(v string) *CreateTrafficMirrorFilterRequestEgressRules {
	s.DestinationPortRange = &v
	return s
}

func (s *CreateTrafficMirrorFilterRequestEgressRules) SetIpVersion(v string) *CreateTrafficMirrorFilterRequestEgressRules {
	s.IpVersion = &v
	return s
}

func (s *CreateTrafficMirrorFilterRequestEgressRules) SetPriority(v int32) *CreateTrafficMirrorFilterRequestEgressRules {
	s.Priority = &v
	return s
}

func (s *CreateTrafficMirrorFilterRequestEgressRules) SetProtocol(v string) *CreateTrafficMirrorFilterRequestEgressRules {
	s.Protocol = &v
	return s
}

func (s *CreateTrafficMirrorFilterRequestEgressRules) SetSourceCidrBlock(v string) *CreateTrafficMirrorFilterRequestEgressRules {
	s.SourceCidrBlock = &v
	return s
}

func (s *CreateTrafficMirrorFilterRequestEgressRules) SetSourcePortRange(v string) *CreateTrafficMirrorFilterRequestEgressRules {
	s.SourcePortRange = &v
	return s
}

func (s *CreateTrafficMirrorFilterRequestEgressRules) Validate() error {
	return dara.Validate(s)
}

type CreateTrafficMirrorFilterRequestIngressRules struct {
	// The collection policy of the inbound rule. Valid values:
	//
	// 	- **accept**: collects the network traffic.
	//
	// 	- **drop**: does not collect the network traffic.
	//
	// example:
	//
	// accept
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The destination CIDR block of the inbound traffic.
	//
	// example:
	//
	// 10.0.0.0/24
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The destination port range of the inbound traffic. Valid value: **1*	- to **65535**. Separate the first port and last port with a forward slash (/). For example, **1/200*	- or **80/80**.
	//
	// >  If you set **IngressRules.N.Protocol*	- to **ALL*	- or **ICMP**, you do not need to set this parameter. In this case, all ports are available.
	//
	// example:
	//
	// 80/120
	DestinationPortRange *string `json:"DestinationPortRange,omitempty" xml:"DestinationPortRange,omitempty"`
	// The IP version of the instance. The following value may be returned:
	//
	// 	- **IPv4**
	//
	// 	- **IPv6**
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The priority of the inbound rule. A smaller value indicates a higher priority. The maximum value of **N*	- is **10**. You can configure up to 10 inbound rules for a filter.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The type of the protocol is used by the inbound traffic that you want to mirror. Valid values:
	//
	// 	- **ALL**: all protocols
	//
	// 	- **ICMP**: Internet Control Message Protocol.
	//
	// 	- **TCP**: Transmission Control Protocol.
	//
	// 	- **UDP**: User Datagram Protocol.
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The source CIDR block of the inbound traffic.
	//
	// example:
	//
	// 10.0.0.0/24
	SourceCidrBlock *string `json:"SourceCidrBlock,omitempty" xml:"SourceCidrBlock,omitempty"`
	// The source port range of the inbound traffic. Valid value: **1*	- to **65535**. Separate the first port and last port with a forward slash (/). For example, **1/200*	- or **80/80**.
	//
	// >  If **IngressRules.N.Protocol*	- is set to **ALL*	- or **ICMP**, you do not need to specify this parameter. This indicates that all ports are available.
	//
	// example:
	//
	// 80/120
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
}

func (s CreateTrafficMirrorFilterRequestIngressRules) String() string {
	return dara.Prettify(s)
}

func (s CreateTrafficMirrorFilterRequestIngressRules) GoString() string {
	return s.String()
}

func (s *CreateTrafficMirrorFilterRequestIngressRules) GetAction() *string {
	return s.Action
}

func (s *CreateTrafficMirrorFilterRequestIngressRules) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *CreateTrafficMirrorFilterRequestIngressRules) GetDestinationPortRange() *string {
	return s.DestinationPortRange
}

func (s *CreateTrafficMirrorFilterRequestIngressRules) GetIpVersion() *string {
	return s.IpVersion
}

func (s *CreateTrafficMirrorFilterRequestIngressRules) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateTrafficMirrorFilterRequestIngressRules) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateTrafficMirrorFilterRequestIngressRules) GetSourceCidrBlock() *string {
	return s.SourceCidrBlock
}

func (s *CreateTrafficMirrorFilterRequestIngressRules) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *CreateTrafficMirrorFilterRequestIngressRules) SetAction(v string) *CreateTrafficMirrorFilterRequestIngressRules {
	s.Action = &v
	return s
}

func (s *CreateTrafficMirrorFilterRequestIngressRules) SetDestinationCidrBlock(v string) *CreateTrafficMirrorFilterRequestIngressRules {
	s.DestinationCidrBlock = &v
	return s
}

func (s *CreateTrafficMirrorFilterRequestIngressRules) SetDestinationPortRange(v string) *CreateTrafficMirrorFilterRequestIngressRules {
	s.DestinationPortRange = &v
	return s
}

func (s *CreateTrafficMirrorFilterRequestIngressRules) SetIpVersion(v string) *CreateTrafficMirrorFilterRequestIngressRules {
	s.IpVersion = &v
	return s
}

func (s *CreateTrafficMirrorFilterRequestIngressRules) SetPriority(v int32) *CreateTrafficMirrorFilterRequestIngressRules {
	s.Priority = &v
	return s
}

func (s *CreateTrafficMirrorFilterRequestIngressRules) SetProtocol(v string) *CreateTrafficMirrorFilterRequestIngressRules {
	s.Protocol = &v
	return s
}

func (s *CreateTrafficMirrorFilterRequestIngressRules) SetSourceCidrBlock(v string) *CreateTrafficMirrorFilterRequestIngressRules {
	s.SourceCidrBlock = &v
	return s
}

func (s *CreateTrafficMirrorFilterRequestIngressRules) SetSourcePortRange(v string) *CreateTrafficMirrorFilterRequestIngressRules {
	s.SourcePortRange = &v
	return s
}

func (s *CreateTrafficMirrorFilterRequestIngressRules) Validate() error {
	return dara.Validate(s)
}

type CreateTrafficMirrorFilterRequestTag struct {
	// The tag key. You can specify at most 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length. The tag key cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify at most 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTrafficMirrorFilterRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateTrafficMirrorFilterRequestTag) GoString() string {
	return s.String()
}

func (s *CreateTrafficMirrorFilterRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateTrafficMirrorFilterRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateTrafficMirrorFilterRequestTag) SetKey(v string) *CreateTrafficMirrorFilterRequestTag {
	s.Key = &v
	return s
}

func (s *CreateTrafficMirrorFilterRequestTag) SetValue(v string) *CreateTrafficMirrorFilterRequestTag {
	s.Value = &v
	return s
}

func (s *CreateTrafficMirrorFilterRequestTag) Validate() error {
	return dara.Validate(s)
}
