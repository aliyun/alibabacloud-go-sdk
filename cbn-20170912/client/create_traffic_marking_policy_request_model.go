// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrafficMarkingPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateTrafficMarkingPolicyRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateTrafficMarkingPolicyRequest
	GetDryRun() *bool
	SetMarkingDscp(v int32) *CreateTrafficMarkingPolicyRequest
	GetMarkingDscp() *int32
	SetOwnerAccount(v string) *CreateTrafficMarkingPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateTrafficMarkingPolicyRequest
	GetOwnerId() *int64
	SetPriority(v int32) *CreateTrafficMarkingPolicyRequest
	GetPriority() *int32
	SetResourceOwnerAccount(v string) *CreateTrafficMarkingPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateTrafficMarkingPolicyRequest
	GetResourceOwnerId() *int64
	SetTrafficMarkingPolicyDescription(v string) *CreateTrafficMarkingPolicyRequest
	GetTrafficMarkingPolicyDescription() *string
	SetTrafficMarkingPolicyName(v string) *CreateTrafficMarkingPolicyRequest
	GetTrafficMarkingPolicyName() *string
	SetTrafficMatchRules(v []*CreateTrafficMarkingPolicyRequestTrafficMatchRules) *CreateTrafficMarkingPolicyRequest
	GetTrafficMatchRules() []*CreateTrafficMarkingPolicyRequestTrafficMatchRules
	SetTransitRouterId(v string) *CreateTrafficMarkingPolicyRequest
	GetTransitRouterId() *string
}

type CreateTrafficMarkingPolicyRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among all requests. The client token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, **ClientToken*	- is set to the value of **RequestId**. The value of **RequestId*	- for each API request may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// 	- **true**: performs a dry run. The system checks the required parameters, request format, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and sends the request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The differentiated services code point (DSCP) value to be added to packets that match the traffic classification rule. Valid values: **0*	- to **63**.
	//
	// The DSCP value of each traffic marking policy on a transit router must be unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	MarkingDscp  *int32  `json:"MarkingDscp,omitempty" xml:"MarkingDscp,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The priority value of the traffic marking policy. Valid values: **1*	- to **100**.
	//
	// The priority value of each traffic marking policy on a transit router must be unique. A smaller value specifies a higher priority.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	Priority             *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The description of the traffic marking policy.
	//
	// This parameter is optional. If you enter a description, it must be 1 to 256 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// desctest
	TrafficMarkingPolicyDescription *string `json:"TrafficMarkingPolicyDescription,omitempty" xml:"TrafficMarkingPolicyDescription,omitempty"`
	// The name of the traffic marking policy.
	//
	// The name can be empty or 1 to 128 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// nametest
	TrafficMarkingPolicyName *string `json:"TrafficMarkingPolicyName,omitempty" xml:"TrafficMarkingPolicyName,omitempty"`
	// The traffic classification rules in the traffic marking policy.
	//
	// Data packets that meet the traffic classification rule is assigned the DSCP value of quality of service (QoS) policy.
	//
	// You can create up to 50 traffic classification rules.
	TrafficMatchRules []*CreateTrafficMarkingPolicyRequestTrafficMatchRules `json:"TrafficMatchRules,omitempty" xml:"TrafficMatchRules,omitempty" type:"Repeated"`
	// The ID of the transit router.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-8vbuqeo5h5pu3m01d****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s CreateTrafficMarkingPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTrafficMarkingPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateTrafficMarkingPolicyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTrafficMarkingPolicyRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateTrafficMarkingPolicyRequest) GetMarkingDscp() *int32 {
	return s.MarkingDscp
}

func (s *CreateTrafficMarkingPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateTrafficMarkingPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateTrafficMarkingPolicyRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateTrafficMarkingPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateTrafficMarkingPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateTrafficMarkingPolicyRequest) GetTrafficMarkingPolicyDescription() *string {
	return s.TrafficMarkingPolicyDescription
}

func (s *CreateTrafficMarkingPolicyRequest) GetTrafficMarkingPolicyName() *string {
	return s.TrafficMarkingPolicyName
}

func (s *CreateTrafficMarkingPolicyRequest) GetTrafficMatchRules() []*CreateTrafficMarkingPolicyRequestTrafficMatchRules {
	return s.TrafficMatchRules
}

func (s *CreateTrafficMarkingPolicyRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *CreateTrafficMarkingPolicyRequest) SetClientToken(v string) *CreateTrafficMarkingPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTrafficMarkingPolicyRequest) SetDryRun(v bool) *CreateTrafficMarkingPolicyRequest {
	s.DryRun = &v
	return s
}

func (s *CreateTrafficMarkingPolicyRequest) SetMarkingDscp(v int32) *CreateTrafficMarkingPolicyRequest {
	s.MarkingDscp = &v
	return s
}

func (s *CreateTrafficMarkingPolicyRequest) SetOwnerAccount(v string) *CreateTrafficMarkingPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTrafficMarkingPolicyRequest) SetOwnerId(v int64) *CreateTrafficMarkingPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTrafficMarkingPolicyRequest) SetPriority(v int32) *CreateTrafficMarkingPolicyRequest {
	s.Priority = &v
	return s
}

func (s *CreateTrafficMarkingPolicyRequest) SetResourceOwnerAccount(v string) *CreateTrafficMarkingPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTrafficMarkingPolicyRequest) SetResourceOwnerId(v int64) *CreateTrafficMarkingPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTrafficMarkingPolicyRequest) SetTrafficMarkingPolicyDescription(v string) *CreateTrafficMarkingPolicyRequest {
	s.TrafficMarkingPolicyDescription = &v
	return s
}

func (s *CreateTrafficMarkingPolicyRequest) SetTrafficMarkingPolicyName(v string) *CreateTrafficMarkingPolicyRequest {
	s.TrafficMarkingPolicyName = &v
	return s
}

func (s *CreateTrafficMarkingPolicyRequest) SetTrafficMatchRules(v []*CreateTrafficMarkingPolicyRequestTrafficMatchRules) *CreateTrafficMarkingPolicyRequest {
	s.TrafficMatchRules = v
	return s
}

func (s *CreateTrafficMarkingPolicyRequest) SetTransitRouterId(v string) *CreateTrafficMarkingPolicyRequest {
	s.TransitRouterId = &v
	return s
}

func (s *CreateTrafficMarkingPolicyRequest) Validate() error {
	return dara.Validate(s)
}

type CreateTrafficMarkingPolicyRequestTrafficMatchRules struct {
	// The address family. You can set the value to IPv4 or IPv6, or leave the value empty.
	//
	// example:
	//
	// IPv4
	AddressFamily *string `json:"AddressFamily,omitempty" xml:"AddressFamily,omitempty"`
	// The destination CIDR block of packets. IPv4 and IPv6 addresses are supported.
	//
	// Packets whose destination IP addresses fall into the specified destination CIDR block meet the traffic classification rule. If you do not specify a destination CIDR block, all packets meet the traffic classification rule.
	//
	// You can create up to 50 traffic classification rules in each call You can specify a destination CIDR block for each traffic classification rule.
	//
	// example:
	//
	// 10.10.10.0/24
	DstCidr *string `json:"DstCidr,omitempty" xml:"DstCidr,omitempty"`
	// The destination port range that is used to match packets. Valid values: **-1*	- and **1*	- to **65535**.
	//
	// Packets whose destination ports fall within the destination port range meet the traffic classification rule. If you do not specify destination port range, all packets meet the traffic classification rule.
	//
	// You can enter up to two port numbers. Take note of the following rules:
	//
	// 	- If you enter only one port number, such as 1, packets whose destination port is 1 meet the traffic classification rule. A value of -1 specifies all destination ports.
	//
	// 	- If you enter two port numbers, such as 1 and 200, packets whose destination ports fall into 1 and 200 meet the traffic classification rule.
	//
	// 	- If you enter two port numbers and one of them is -1, the other port number must also be -1. In this case, all packets meet the traffic classification rule.
	//
	// You can create up to 50 traffic classification rules in each call. You can specify a destination port range for each traffic classification rule.
	DstPortRange []*int32 `json:"DstPortRange,omitempty" xml:"DstPortRange,omitempty" type:"Repeated"`
	// The Differentiated Service Code Point (DSCP) value that is used to match packets. Valid values: **0*	- to **63**.
	//
	// Packets that carry the specified DSCP value meet the traffic classification rule. If you do not specify a DSCP value, all packets meet the traffic classification rule.
	//
	// >  The DSCP value that you specify for this parameter is the DSCP value that packets carry before they are transmitted over the inter-region connection.
	//
	// You can create up to 50 traffic classification rules in each call. You can specify a DSCP value for each traffic classification rule.
	//
	// example:
	//
	// 6
	MatchDscp *int32 `json:"MatchDscp,omitempty" xml:"MatchDscp,omitempty"`
	// The protocol that is used to match packets.
	//
	// Traffic classification rules support the following protocols: **HTTP**, **HTTPS**, **TCP**, **UDP**, **SSH**, and **Telnet**. For more information, log on to the [CEN console](https://cen.console.aliyun.com/cen/list).
	//
	// **Some protocols use a fixed port. Click to view the protocols and ports.**
	//
	// 	- If the protocol is **ICMP**, the destination port must be **-1**.
	//
	// 	- If the protocol is **GRE**, the destination port must be **1**.
	//
	// 	- If the protocol is **SSH**, the destination port must be **22**.
	//
	// 	- If the protocol is **Telnet**, the destination port must be **23**.
	//
	// 	- If the protocol is **HTTP**, the destination port must be **80**.
	//
	// 	- If the protocol is **HTTPS**, the destination port must be **443**.
	//
	// 	- If the protocol is **MS SQL**, the destination port must be **1443**.
	//
	// 	- If the protocol is **Oracle**, the destination port must be **1521**.
	//
	// 	- If the protocol is **Mysql**, the destination port must be **3306**.
	//
	// 	- If the protocol is **RDP**, the destination port must be **3389**.
	//
	// 	- If the protocol is **Postgre SQL**, the destination port must be **5432**.
	//
	// 	- If the protocol is **Redis**, the destination port must be **6379**.
	//
	// You can create up to 50 traffic classification rules in each call. You can specify a protocol for each traffic classification rule.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The source CIDR block of packets. IPv6 and IPv4 addresses are supported.
	//
	// Packets whose source IP addresses fall into the specified source CIDR block meet the traffic classification rule. If you do not specify a source CIDR block, all packets meet the traffic classification rule.
	//
	// You can create up to 50 traffic classification rules in each call. You can specify a source CIDR block for each traffic classification rule.
	//
	// example:
	//
	// 192.168.10.0/24
	SrcCidr *string `json:"SrcCidr,omitempty" xml:"SrcCidr,omitempty"`
	// The source port range that is used to match packets. Valid values: **-1*	- and **1*	- to **65535**.
	//
	// Packets whose source ports fall within the source port range meet the traffic classification rule. If you do not specify a source port range, all packets meet the traffic classification rule.
	//
	// You can enter up to two port numbers. Take note of the following rules:
	//
	// 	- If you enter only one port number, such as 1, packets whose source port is 1 meet the traffic classification rule. A value of -1 specifies all source ports.
	//
	// 	- If you enter two port numbers, such as 1 and 200, packets whose source ports fall into 1 and 200 meet the traffic classification rule.
	//
	// 	- If you enter two port numbers and one of them is -1, the other port number must also be -1. In this case, all packets meet the traffic classification rule.
	//
	// You can create up to 50 traffic classification rules in each call. You can specify a source port range for each traffic classification rule.
	SrcPortRange []*int32 `json:"SrcPortRange,omitempty" xml:"SrcPortRange,omitempty" type:"Repeated"`
	// The description of the traffic classification rule.
	//
	// You can create up to 50 traffic classification rules in each call. You can specify a description for each traffic classification rule.
	//
	// This parameter is optional. If you enter a description, it must be 1 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// desctest
	TrafficMatchRuleDescription *string `json:"TrafficMatchRuleDescription,omitempty" xml:"TrafficMatchRuleDescription,omitempty"`
	// The name of the traffic classification rule.
	//
	// You can create up to 50 traffic classification rules in each call. You can specify a name for each traffic classification rule.
	//
	// The name can be empty or 1 to 128 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// nametest
	TrafficMatchRuleName *string `json:"TrafficMatchRuleName,omitempty" xml:"TrafficMatchRuleName,omitempty"`
}

func (s CreateTrafficMarkingPolicyRequestTrafficMatchRules) String() string {
	return dara.Prettify(s)
}

func (s CreateTrafficMarkingPolicyRequestTrafficMatchRules) GoString() string {
	return s.String()
}

func (s *CreateTrafficMarkingPolicyRequestTrafficMatchRules) GetAddressFamily() *string {
	return s.AddressFamily
}

func (s *CreateTrafficMarkingPolicyRequestTrafficMatchRules) GetDstCidr() *string {
	return s.DstCidr
}

func (s *CreateTrafficMarkingPolicyRequestTrafficMatchRules) GetDstPortRange() []*int32 {
	return s.DstPortRange
}

func (s *CreateTrafficMarkingPolicyRequestTrafficMatchRules) GetMatchDscp() *int32 {
	return s.MatchDscp
}

func (s *CreateTrafficMarkingPolicyRequestTrafficMatchRules) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateTrafficMarkingPolicyRequestTrafficMatchRules) GetSrcCidr() *string {
	return s.SrcCidr
}

func (s *CreateTrafficMarkingPolicyRequestTrafficMatchRules) GetSrcPortRange() []*int32 {
	return s.SrcPortRange
}

func (s *CreateTrafficMarkingPolicyRequestTrafficMatchRules) GetTrafficMatchRuleDescription() *string {
	return s.TrafficMatchRuleDescription
}

func (s *CreateTrafficMarkingPolicyRequestTrafficMatchRules) GetTrafficMatchRuleName() *string {
	return s.TrafficMatchRuleName
}

func (s *CreateTrafficMarkingPolicyRequestTrafficMatchRules) SetAddressFamily(v string) *CreateTrafficMarkingPolicyRequestTrafficMatchRules {
	s.AddressFamily = &v
	return s
}

func (s *CreateTrafficMarkingPolicyRequestTrafficMatchRules) SetDstCidr(v string) *CreateTrafficMarkingPolicyRequestTrafficMatchRules {
	s.DstCidr = &v
	return s
}

func (s *CreateTrafficMarkingPolicyRequestTrafficMatchRules) SetDstPortRange(v []*int32) *CreateTrafficMarkingPolicyRequestTrafficMatchRules {
	s.DstPortRange = v
	return s
}

func (s *CreateTrafficMarkingPolicyRequestTrafficMatchRules) SetMatchDscp(v int32) *CreateTrafficMarkingPolicyRequestTrafficMatchRules {
	s.MatchDscp = &v
	return s
}

func (s *CreateTrafficMarkingPolicyRequestTrafficMatchRules) SetProtocol(v string) *CreateTrafficMarkingPolicyRequestTrafficMatchRules {
	s.Protocol = &v
	return s
}

func (s *CreateTrafficMarkingPolicyRequestTrafficMatchRules) SetSrcCidr(v string) *CreateTrafficMarkingPolicyRequestTrafficMatchRules {
	s.SrcCidr = &v
	return s
}

func (s *CreateTrafficMarkingPolicyRequestTrafficMatchRules) SetSrcPortRange(v []*int32) *CreateTrafficMarkingPolicyRequestTrafficMatchRules {
	s.SrcPortRange = v
	return s
}

func (s *CreateTrafficMarkingPolicyRequestTrafficMatchRules) SetTrafficMatchRuleDescription(v string) *CreateTrafficMarkingPolicyRequestTrafficMatchRules {
	s.TrafficMatchRuleDescription = &v
	return s
}

func (s *CreateTrafficMarkingPolicyRequestTrafficMatchRules) SetTrafficMatchRuleName(v string) *CreateTrafficMarkingPolicyRequestTrafficMatchRules {
	s.TrafficMatchRuleName = &v
	return s
}

func (s *CreateTrafficMarkingPolicyRequestTrafficMatchRules) Validate() error {
	return dara.Validate(s)
}
