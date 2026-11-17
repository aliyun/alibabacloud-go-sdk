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
	// You can generate a parameter value on your client to ensure that the value is unique across different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to only perform a dry run for the request. Valid values:
	//
	// - **true**: performs a dry run for the request without creating the traffic mirror filter. The system checks whether required parameters are specified, the request format, and service limits. If the dry run fails, the corresponding error is returned. If the dry run succeeds, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends a normal request. If the request passes the check, a 2xx HTTP status code is returned and the filter is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The information about the outbound rules.
	EgressRules []*CreateTrafficMirrorFilterRequestEgressRules `json:"EgressRules,omitempty" xml:"EgressRules,omitempty" type:"Repeated"`
	// The information about the inbound rules.
	IngressRules []*CreateTrafficMirrorFilterRequestIngressRules `json:"IngressRules,omitempty" xml:"IngressRules,omitempty" type:"Repeated"`
	OwnerAccount *string                                         `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64                                          `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the traffic mirror.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query region IDs. For more information about the regions that support traffic mirroring, see [Traffic mirroring overview](https://help.aliyun.com/document_detail/207513.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hongkong
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the traffic mirroring belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4ph****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags of the resource.
	Tag []*CreateTrafficMirrorFilterRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The description of the traffic mirror filter.
	//
	// The description must be 1 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// this is a filter.
	TrafficMirrorFilterDescription *string `json:"TrafficMirrorFilterDescription,omitempty" xml:"TrafficMirrorFilterDescription,omitempty"`
	// The name of the traffic mirror filter.
	//
	// The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
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
	if s.EgressRules != nil {
		for _, item := range s.EgressRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.IngressRules != nil {
		for _, item := range s.IngressRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type CreateTrafficMirrorFilterRequestEgressRules struct {
	// The collection policy of the outbound rule. Valid values:
	//
	// - **accept**: collects network traffic.
	//
	// - **drop**: does not collect network traffic.
	//
	// example:
	//
	// accept
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The destination CIDR block of the network traffic for the outbound rule.
	//
	// example:
	//
	// 10.0.0.0/24
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The destination port range of the network traffic for the outbound rule. Valid values are **1*	- to **65535**. Separate the start port and the end port with a forward slash (/). Examples: **1/200*	- and **80/80**. You cannot set this parameter only to **-1/-1**, which indicates that all ports are allowed.
	//
	// > If **EgressRules.N.Protocol*	- is set to **ALL*	- or **ICMP**, you do not need to specify this parameter, which indicates that all ports are allowed.
	//
	// example:
	//
	// 22/40
	DestinationPortRange *string `json:"DestinationPortRange,omitempty" xml:"DestinationPortRange,omitempty"`
	// The IP version of the instance. Valid values:
	//
	// - **IPv4**: IPv4.
	//
	// - **IPv6**: IPv6.
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The priority of the outbound rule. A smaller value indicates a higher priority.
	//
	// The maximum value of **N*	- is **10**, which indicates that you can configure up to 10 outbound rules for a filter.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The protocol type of the network traffic to be mirrored for the outbound rule. Valid values:
	//
	// - **ALL**: all protocols.
	//
	// - **ICMP**: Internet Control Message Protocol (ICMP).
	//
	// - **TCP**: Transmission Control Protocol (TCP).
	//
	// - **UDP**: User Datagram Protocol (UDP).
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The source CIDR block of the network traffic for the outbound rule.
	//
	// example:
	//
	// 10.0.0.0/24
	SourceCidrBlock *string `json:"SourceCidrBlock,omitempty" xml:"SourceCidrBlock,omitempty"`
	// The source port range of the network traffic for the outbound rule. Valid values are **1*	- to **65535**. Separate the start port and the end port with a forward slash (/). Examples: **1/200*	- and **80/80**. You cannot set this parameter only to **-1/-1**, which indicates that all ports are allowed.
	//
	// > If **EgressRules.N.Protocol*	- is set to **ALL*	- or **ICMP**, you do not need to specify this parameter, which indicates that all ports are allowed.
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
	// - **accept**: collects network traffic.
	//
	// - **drop**: does not collect network traffic.
	//
	// example:
	//
	// accept
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The destination CIDR block of the network traffic for the inbound rule.
	//
	// example:
	//
	// 10.0.0.0/24
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The destination port range of the network traffic for the inbound rule. Valid values are **1*	- to **65535**. Separate the start port and the end port with a forward slash (/). Examples: **1/200*	- and **80/80**.
	//
	// > If **IngressRules.N.Protocol*	- is set to **ALL*	- or **ICMP**, you do not need to specify this parameter, which indicates that all ports are allowed.
	//
	// example:
	//
	// 80/120
	DestinationPortRange *string `json:"DestinationPortRange,omitempty" xml:"DestinationPortRange,omitempty"`
	// The IP version of the instance. Valid values:
	//
	// - **IPv4**: IPv4.
	//
	// - **IPv6**: IPv6.
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The priority of the inbound rule. A smaller value indicates a higher priority.
	//
	// The maximum value of **N*	- is **10**, which indicates that you can configure up to 10 inbound rules for a filter.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The protocol type of the network traffic to be mirrored for the inbound rule. Valid values:
	//
	// - **ALL**: all protocols.
	//
	// - **ICMP**: Internet Control Message Protocol (ICMP).
	//
	// - **TCP**: Transmission Control Protocol (TCP).
	//
	// - **UDP**: User Datagram Protocol (UDP).
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The source CIDR block of the network traffic for the inbound rule.
	//
	// example:
	//
	// 10.0.0.0/24
	SourceCidrBlock *string `json:"SourceCidrBlock,omitempty" xml:"SourceCidrBlock,omitempty"`
	// The source port range of the network traffic for the inbound rule. Valid values are **1*	- to **65535**. Separate the start port and the end port with a forward slash (/). Examples: **1/200*	- and **80/80**.
	//
	// > If **IngressRules.N.Protocol*	- is set to **ALL*	- or **ICMP**, you do not need to specify this parameter, which indicates that all ports are allowed.
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
	// The tag key of the resource. You can specify up to 20 tag keys. If you specify this parameter, you cannot specify an empty string.
	//
	// A tag key can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. If you specify this parameter, you can specify an empty string.
	//
	// A tag value can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
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
