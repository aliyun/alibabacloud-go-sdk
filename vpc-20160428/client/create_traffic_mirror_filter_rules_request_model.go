// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrafficMirrorFilterRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateTrafficMirrorFilterRulesRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateTrafficMirrorFilterRulesRequest
	GetDryRun() *bool
	SetEgressRules(v []*CreateTrafficMirrorFilterRulesRequestEgressRules) *CreateTrafficMirrorFilterRulesRequest
	GetEgressRules() []*CreateTrafficMirrorFilterRulesRequestEgressRules
	SetIngressRules(v []*CreateTrafficMirrorFilterRulesRequestIngressRules) *CreateTrafficMirrorFilterRulesRequest
	GetIngressRules() []*CreateTrafficMirrorFilterRulesRequestIngressRules
	SetOwnerAccount(v string) *CreateTrafficMirrorFilterRulesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateTrafficMirrorFilterRulesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateTrafficMirrorFilterRulesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateTrafficMirrorFilterRulesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateTrafficMirrorFilterRulesRequest
	GetResourceOwnerId() *int64
	SetTrafficMirrorFilterId(v string) *CreateTrafficMirrorFilterRulesRequest
	GetTrafficMirrorFilterId() *string
}

type CreateTrafficMirrorFilterRulesRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, the system uses **RequestId*	- as **ClientToken**. **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to check the request without performing the operation. Valid values:
	//
	// 	- **true**: checks the request without performing the operation. The system checks the required parameters, request format, and limits. If the request fails the precheck, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): sends the request. After the request passes the check, the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The information about the outbound rule.
	EgressRules []*CreateTrafficMirrorFilterRulesRequestEgressRules `json:"EgressRules,omitempty" xml:"EgressRules,omitempty" type:"Repeated"`
	// The information about the inbound rules.
	IngressRules []*CreateTrafficMirrorFilterRulesRequestIngressRules `json:"IngressRules,omitempty" xml:"IngressRules,omitempty" type:"Repeated"`
	OwnerAccount *string                                              `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64                                               `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to which the mirrored traffic belongs.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list. For more information about regions that support traffic mirror, see [Overview of traffic mirror](https://help.aliyun.com/document_detail/207513.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hongkong
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the filter.
	//
	// This parameter is required.
	//
	// example:
	//
	// tmf-j6cmls82xnc86vtpe****
	TrafficMirrorFilterId *string `json:"TrafficMirrorFilterId,omitempty" xml:"TrafficMirrorFilterId,omitempty"`
}

func (s CreateTrafficMirrorFilterRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTrafficMirrorFilterRulesRequest) GoString() string {
	return s.String()
}

func (s *CreateTrafficMirrorFilterRulesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTrafficMirrorFilterRulesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateTrafficMirrorFilterRulesRequest) GetEgressRules() []*CreateTrafficMirrorFilterRulesRequestEgressRules {
	return s.EgressRules
}

func (s *CreateTrafficMirrorFilterRulesRequest) GetIngressRules() []*CreateTrafficMirrorFilterRulesRequestIngressRules {
	return s.IngressRules
}

func (s *CreateTrafficMirrorFilterRulesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateTrafficMirrorFilterRulesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateTrafficMirrorFilterRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTrafficMirrorFilterRulesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateTrafficMirrorFilterRulesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateTrafficMirrorFilterRulesRequest) GetTrafficMirrorFilterId() *string {
	return s.TrafficMirrorFilterId
}

func (s *CreateTrafficMirrorFilterRulesRequest) SetClientToken(v string) *CreateTrafficMirrorFilterRulesRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTrafficMirrorFilterRulesRequest) SetDryRun(v bool) *CreateTrafficMirrorFilterRulesRequest {
	s.DryRun = &v
	return s
}

func (s *CreateTrafficMirrorFilterRulesRequest) SetEgressRules(v []*CreateTrafficMirrorFilterRulesRequestEgressRules) *CreateTrafficMirrorFilterRulesRequest {
	s.EgressRules = v
	return s
}

func (s *CreateTrafficMirrorFilterRulesRequest) SetIngressRules(v []*CreateTrafficMirrorFilterRulesRequestIngressRules) *CreateTrafficMirrorFilterRulesRequest {
	s.IngressRules = v
	return s
}

func (s *CreateTrafficMirrorFilterRulesRequest) SetOwnerAccount(v string) *CreateTrafficMirrorFilterRulesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTrafficMirrorFilterRulesRequest) SetOwnerId(v int64) *CreateTrafficMirrorFilterRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTrafficMirrorFilterRulesRequest) SetRegionId(v string) *CreateTrafficMirrorFilterRulesRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTrafficMirrorFilterRulesRequest) SetResourceOwnerAccount(v string) *CreateTrafficMirrorFilterRulesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTrafficMirrorFilterRulesRequest) SetResourceOwnerId(v int64) *CreateTrafficMirrorFilterRulesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTrafficMirrorFilterRulesRequest) SetTrafficMirrorFilterId(v string) *CreateTrafficMirrorFilterRulesRequest {
	s.TrafficMirrorFilterId = &v
	return s
}

func (s *CreateTrafficMirrorFilterRulesRequest) Validate() error {
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
	return nil
}

type CreateTrafficMirrorFilterRulesRequestEgressRules struct {
	// The collection policy of the outbound rule. Valid values:
	//
	// 	- **accept**: accepts network traffic.
	//
	// 	- **drop**: drops network traffic.
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
	// The destination port range of the outbound traffic. Valid value: **1*	- to **65535**. Separate the first and last port with a forward slash (/). For example **1/200*	- and **80/80**. You cannot set this parameter to \\*\\*-1/-1\\*\\*, which indicates all ports.
	//
	// >  If **EgressRules.N.Protocol*	- is set to **ALL*	- or **ICMP**, you do not need to set this parameter. In this case, all ports are available.
	//
	// example:
	//
	// 22/40
	DestinationPortRange *string `json:"DestinationPortRange,omitempty" xml:"DestinationPortRange,omitempty"`
	// The IP version of the instance. Valid values:
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
	// The protocol that is used by the outbound traffic to be mirrored. Valid values:
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
	// The source port range of the outbound traffic. Valid value: **1*	- to **65535**. Separate the first and last port with a forward slash (/). For example **1/200*	- and **80/80**. You cannot set this parameter to \\*\\*-1/-1\\*\\*, which indicates all ports.
	//
	// >  If **EgressRules.N.Protocol*	- is set to **ALL*	- or **ICMP**, you do not need to set this parameter. In this case, all ports are available.
	//
	// example:
	//
	// 22/40
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
}

func (s CreateTrafficMirrorFilterRulesRequestEgressRules) String() string {
	return dara.Prettify(s)
}

func (s CreateTrafficMirrorFilterRulesRequestEgressRules) GoString() string {
	return s.String()
}

func (s *CreateTrafficMirrorFilterRulesRequestEgressRules) GetAction() *string {
	return s.Action
}

func (s *CreateTrafficMirrorFilterRulesRequestEgressRules) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *CreateTrafficMirrorFilterRulesRequestEgressRules) GetDestinationPortRange() *string {
	return s.DestinationPortRange
}

func (s *CreateTrafficMirrorFilterRulesRequestEgressRules) GetIpVersion() *string {
	return s.IpVersion
}

func (s *CreateTrafficMirrorFilterRulesRequestEgressRules) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateTrafficMirrorFilterRulesRequestEgressRules) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateTrafficMirrorFilterRulesRequestEgressRules) GetSourceCidrBlock() *string {
	return s.SourceCidrBlock
}

func (s *CreateTrafficMirrorFilterRulesRequestEgressRules) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *CreateTrafficMirrorFilterRulesRequestEgressRules) SetAction(v string) *CreateTrafficMirrorFilterRulesRequestEgressRules {
	s.Action = &v
	return s
}

func (s *CreateTrafficMirrorFilterRulesRequestEgressRules) SetDestinationCidrBlock(v string) *CreateTrafficMirrorFilterRulesRequestEgressRules {
	s.DestinationCidrBlock = &v
	return s
}

func (s *CreateTrafficMirrorFilterRulesRequestEgressRules) SetDestinationPortRange(v string) *CreateTrafficMirrorFilterRulesRequestEgressRules {
	s.DestinationPortRange = &v
	return s
}

func (s *CreateTrafficMirrorFilterRulesRequestEgressRules) SetIpVersion(v string) *CreateTrafficMirrorFilterRulesRequestEgressRules {
	s.IpVersion = &v
	return s
}

func (s *CreateTrafficMirrorFilterRulesRequestEgressRules) SetPriority(v int32) *CreateTrafficMirrorFilterRulesRequestEgressRules {
	s.Priority = &v
	return s
}

func (s *CreateTrafficMirrorFilterRulesRequestEgressRules) SetProtocol(v string) *CreateTrafficMirrorFilterRulesRequestEgressRules {
	s.Protocol = &v
	return s
}

func (s *CreateTrafficMirrorFilterRulesRequestEgressRules) SetSourceCidrBlock(v string) *CreateTrafficMirrorFilterRulesRequestEgressRules {
	s.SourceCidrBlock = &v
	return s
}

func (s *CreateTrafficMirrorFilterRulesRequestEgressRules) SetSourcePortRange(v string) *CreateTrafficMirrorFilterRulesRequestEgressRules {
	s.SourcePortRange = &v
	return s
}

func (s *CreateTrafficMirrorFilterRulesRequestEgressRules) Validate() error {
	return dara.Validate(s)
}

type CreateTrafficMirrorFilterRulesRequestIngressRules struct {
	// The policy of the inbound rule. Valid values:
	//
	// 	- **accept**: collects network traffic.
	//
	// 	- **drop**: drops network traffic.
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
	// The destination port range of the inbound traffic. Valid value: **1*	- to **65535**. Separate the first and the last port with a forward slash (/). For example, **1/200*	- or **80/80**.
	//
	// >  If the **IngressRules.N.Protocol*	- parameter is set to **ALL*	- or **ICMP**, you do not need to set this parameter. In this case, all ports are available.
	//
	// example:
	//
	// 80/120
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
	// The priority of the inbound rule. A smaller value indicates a higher priority. The maximum value of **N*	- is **10**. You can configure up to 10 inbound rules for a filter.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The protocol that is used by the inbound traffic to be mirrored. Valid values:
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
	// The source port range of the inbound traffic. Valid value: **1*	- to **65535**. Separate the first and last port with a forward slash (/). For example **1/200*	- and **80/80**. You cannot set this parameter to \\*\\*-1/-1\\*\\*, which indicates all ports.
	//
	// >  If the **IngressRules.N.Protocol*	- parameter is set to **ALL*	- or **ICMP**, you do not need to set this parameter. In this case, all ports are available.
	//
	// example:
	//
	// 80/120
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
}

func (s CreateTrafficMirrorFilterRulesRequestIngressRules) String() string {
	return dara.Prettify(s)
}

func (s CreateTrafficMirrorFilterRulesRequestIngressRules) GoString() string {
	return s.String()
}

func (s *CreateTrafficMirrorFilterRulesRequestIngressRules) GetAction() *string {
	return s.Action
}

func (s *CreateTrafficMirrorFilterRulesRequestIngressRules) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *CreateTrafficMirrorFilterRulesRequestIngressRules) GetDestinationPortRange() *string {
	return s.DestinationPortRange
}

func (s *CreateTrafficMirrorFilterRulesRequestIngressRules) GetIpVersion() *string {
	return s.IpVersion
}

func (s *CreateTrafficMirrorFilterRulesRequestIngressRules) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateTrafficMirrorFilterRulesRequestIngressRules) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateTrafficMirrorFilterRulesRequestIngressRules) GetSourceCidrBlock() *string {
	return s.SourceCidrBlock
}

func (s *CreateTrafficMirrorFilterRulesRequestIngressRules) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *CreateTrafficMirrorFilterRulesRequestIngressRules) SetAction(v string) *CreateTrafficMirrorFilterRulesRequestIngressRules {
	s.Action = &v
	return s
}

func (s *CreateTrafficMirrorFilterRulesRequestIngressRules) SetDestinationCidrBlock(v string) *CreateTrafficMirrorFilterRulesRequestIngressRules {
	s.DestinationCidrBlock = &v
	return s
}

func (s *CreateTrafficMirrorFilterRulesRequestIngressRules) SetDestinationPortRange(v string) *CreateTrafficMirrorFilterRulesRequestIngressRules {
	s.DestinationPortRange = &v
	return s
}

func (s *CreateTrafficMirrorFilterRulesRequestIngressRules) SetIpVersion(v string) *CreateTrafficMirrorFilterRulesRequestIngressRules {
	s.IpVersion = &v
	return s
}

func (s *CreateTrafficMirrorFilterRulesRequestIngressRules) SetPriority(v int32) *CreateTrafficMirrorFilterRulesRequestIngressRules {
	s.Priority = &v
	return s
}

func (s *CreateTrafficMirrorFilterRulesRequestIngressRules) SetProtocol(v string) *CreateTrafficMirrorFilterRulesRequestIngressRules {
	s.Protocol = &v
	return s
}

func (s *CreateTrafficMirrorFilterRulesRequestIngressRules) SetSourceCidrBlock(v string) *CreateTrafficMirrorFilterRulesRequestIngressRules {
	s.SourceCidrBlock = &v
	return s
}

func (s *CreateTrafficMirrorFilterRulesRequestIngressRules) SetSourcePortRange(v string) *CreateTrafficMirrorFilterRulesRequestIngressRules {
	s.SourcePortRange = &v
	return s
}

func (s *CreateTrafficMirrorFilterRulesRequestIngressRules) Validate() error {
	return dara.Validate(s)
}
