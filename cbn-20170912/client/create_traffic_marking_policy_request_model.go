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
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run. The system checks the required parameters, request syntax, and business restrictions. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): performs a dry run and sends the request. If the request passes the dry run, the traffic marking policy is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The DSCP value to be added to traffic packets that match the traffic classification rules. Valid values: **0*	- to **63**.
	//
	// The DSCP value of each traffic marking policy under a transit router instance must be unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	MarkingDscp  *int32  `json:"MarkingDscp,omitempty" xml:"MarkingDscp,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The priority of the traffic marking policy. Valid values: **1*	- to **100**.
	//
	// The priority of each traffic marking policy under a transit router instance must be unique. A smaller value indicates a higher priority.
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
	// The description can be empty or 1 to 256 characters in length, and cannot start with http:// or https://.
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
	// The list of traffic classification rules for the traffic marking policy.
	//
	// Traffic packets that match the traffic classification rules are marked with the DSCP value of the traffic marking policy.
	//
	// You can create up to 50 traffic classification rules at a time.
	TrafficMatchRules []*CreateTrafficMarkingPolicyRequestTrafficMatchRules `json:"TrafficMatchRules,omitempty" xml:"TrafficMatchRules,omitempty" type:"Repeated"`
	// The instance ID of the forward routing transit router.
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
	if s.TrafficMatchRules != nil {
		for _, item := range s.TrafficMatchRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateTrafficMarkingPolicyRequestTrafficMatchRules struct {
	// The address type. Valid values: IPv4, IPv6, or empty.
	//
	// example:
	//
	// IPv4
	AddressFamily *string `json:"AddressFamily,omitempty" xml:"AddressFamily,omitempty"`
	// The destination CIDR block of traffic packets. IPv4 and IPv6 addresses are supported.
	//
	// The traffic classification rule matches traffic whose destination IP address falls within the destination CIDR block. If you do not set this parameter, the traffic classification rule matches traffic with any destination IP address.
	//
	// You can create up to 50 traffic classification rules at a time, and each traffic classification rule can specify one destination CIDR block.
	//
	// example:
	//
	// 10.10.10.0/24
	DstCidr *string `json:"DstCidr,omitempty" xml:"DstCidr,omitempty"`
	// The destination port of traffic packets. Valid values: **-1*	- and **1*	- to **65535**.
	//
	// The traffic classification rule matches traffic whose destination port falls within the destination port range. If you do not set this parameter, the traffic classification rule matches traffic with any destination port.
	//
	// This parameter supports up to two port numbers. The input format is described as follows:
	//
	// - If you enter only one port number, such as 1, the system matches traffic whose destination port is 1 by default. If the value is -1, the system matches traffic with any destination port.
	//
	// - If you enter two port numbers, such as 1 and 200, the system matches traffic whose destination port is in the range of 1 to 200 by default.
	//
	// - If you enter two port numbers and one of them is -1, the other port number must also be -1, which indicates matching traffic with any destination port.
	//
	// You can create up to 50 traffic classification rules at a time, and each traffic classification rule can specify one destination port range.
	DstPortRange []*int32 `json:"DstPortRange,omitempty" xml:"DstPortRange,omitempty" type:"Repeated"`
	// The DSCP value of traffic packets. Valid values: **0*	- to **63**.
	//
	// The traffic classification rule matches traffic that contains the specified DSCP value. If you do not set this parameter, the traffic classification rule matches traffic with any DSCP value.
	//
	// > The DSCP value refers to the DSCP value that the traffic packets already carry before entering the inter-region connection.
	//
	// You can create up to 50 traffic classification rules at a time, and each traffic classification rule can match one DSCP value.
	//
	// example:
	//
	// 6
	MatchDscp *int32 `json:"MatchDscp,omitempty" xml:"MatchDscp,omitempty"`
	// The protocol type of traffic packets.
	//
	// The traffic marking policy supports matching traffic of multiple protocol types such as **HTTP**, **HTTPS**, **TCP**, **UDP**, **SSH**, and **Telnet**. For more protocol types, log on to the [Cloud Enterprise Network (CEN) console](https://cen.console.aliyun.com/cen/list).
	//
	// <details>
	//
	// <summary>Some protocols have fixed ports. Click to view port details.</summary>
	//
	// - If the protocol type is **ICMP**, the destination port must be set to **-1**.
	//
	// - If the protocol type is **GRE**, the destination port must be set to **-1**.
	//
	// - If the protocol type is **SSH**, the destination port must be set to **22**.
	//
	// - If the protocol type is **Telnet**, the destination port must be set to **23**.
	//
	// - If the protocol type is **HTTP**, the destination port must be set to **80**.
	//
	// - If the protocol type is **HTTPS**, the destination port must be set to **443**.
	//
	// - If the protocol type is **MS SQL**, the destination port must be set to **1443**.
	//
	// - If the protocol type is **Oracle**, the destination port must be set to **1521**.
	//
	// - If the protocol type is **Mysql**, the destination port must be set to **3306**.
	//
	// - If the protocol type is **RDP**, the destination port must be set to **3389**.
	//
	// - If the protocol type is **Postgre SQL**, the destination port must be set to **5432**.
	//
	// - If the protocol type is **Redis**, the destination port must be set to **6379**.
	//
	// </details>
	//
	// You can create up to 50 traffic classification rules at a time, and each traffic classification rule can match one protocol type.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The source CIDR block of traffic packets. IPv6 and IPv4 addresses are supported.
	//
	// The traffic classification rule matches traffic whose source IP address falls within the source CIDR block. If you do not set this parameter, the traffic classification rule matches traffic with any source IP address.
	//
	// You can create up to 50 traffic classification rules at a time, and each traffic classification rule can match one source CIDR block.
	//
	// example:
	//
	// 192.168.10.0/24
	SrcCidr *string `json:"SrcCidr,omitempty" xml:"SrcCidr,omitempty"`
	// The source port of traffic packets. Valid values: **-1*	- and **1*	- to **65535**.
	//
	// The traffic classification rule matches traffic whose source port falls within the source port range. If you do not set this parameter, the traffic classification rule matches traffic with any source port.
	//
	// This parameter supports up to two port numbers. The input format is described as follows:
	//
	// - If you enter only one port number, such as 1, the system matches traffic whose source port is 1 by default. If the value is -1, the system matches traffic with any source port.
	//
	// - If you enter two port numbers, such as 1 and 200, the system matches traffic whose source port is in the range of 1 to 200 by default.
	//
	// - If you enter two port numbers and one of them is -1, the other port number must also be -1, which indicates matching traffic with any source port.
	//
	// You can create up to 50 traffic classification rules at a time, and each traffic classification rule can specify one source port range.
	SrcPortRange []*int32 `json:"SrcPortRange,omitempty" xml:"SrcPortRange,omitempty" type:"Repeated"`
	// The description of the traffic classification rule.
	//
	// You can create up to 50 traffic classification rules at a time, and each traffic classification rule can have one description.
	//
	// The description can be empty or 1 to 256 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// desctest
	TrafficMatchRuleDescription *string `json:"TrafficMatchRuleDescription,omitempty" xml:"TrafficMatchRuleDescription,omitempty"`
	// The name of the traffic classification rule.
	//
	// You can create up to 50 traffic classification rules at a time, and each traffic classification rule can have one name.
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
