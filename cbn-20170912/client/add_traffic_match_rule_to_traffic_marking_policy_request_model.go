// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTrafficMatchRuleToTrafficMarkingPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AddTrafficMatchRuleToTrafficMarkingPolicyRequest
	GetClientToken() *string
	SetDryRun(v bool) *AddTrafficMatchRuleToTrafficMarkingPolicyRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *AddTrafficMatchRuleToTrafficMarkingPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddTrafficMatchRuleToTrafficMarkingPolicyRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AddTrafficMatchRuleToTrafficMarkingPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddTrafficMatchRuleToTrafficMarkingPolicyRequest
	GetResourceOwnerId() *int64
	SetTrafficMarkingPolicyId(v string) *AddTrafficMatchRuleToTrafficMarkingPolicyRequest
	GetTrafficMarkingPolicyId() *string
	SetTrafficMatchRules(v []*AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) *AddTrafficMatchRuleToTrafficMarkingPolicyRequest
	GetTrafficMatchRules() []*AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules
}

type AddTrafficMatchRuleToTrafficMarkingPolicyRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): performs a dry run and sends the request. If the request passes the dry run, traffic classification rules are added to the traffic marking policy.
	//
	// example:
	//
	// false
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the traffic marking policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// tm-u9nxup5kww5po8****
	TrafficMarkingPolicyId *string `json:"TrafficMarkingPolicyId,omitempty" xml:"TrafficMarkingPolicyId,omitempty"`
	// The list of traffic classification rules.
	//
	// You can add up to 50 traffic classification rules at a time.
	TrafficMatchRules []*AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules `json:"TrafficMatchRules,omitempty" xml:"TrafficMatchRules,omitempty" type:"Repeated"`
}

func (s AddTrafficMatchRuleToTrafficMarkingPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s AddTrafficMatchRuleToTrafficMarkingPolicyRequest) GoString() string {
	return s.String()
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequest) GetTrafficMarkingPolicyId() *string {
	return s.TrafficMarkingPolicyId
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequest) GetTrafficMatchRules() []*AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules {
	return s.TrafficMatchRules
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequest) SetClientToken(v string) *AddTrafficMatchRuleToTrafficMarkingPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequest) SetDryRun(v bool) *AddTrafficMatchRuleToTrafficMarkingPolicyRequest {
	s.DryRun = &v
	return s
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequest) SetOwnerAccount(v string) *AddTrafficMatchRuleToTrafficMarkingPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequest) SetOwnerId(v int64) *AddTrafficMatchRuleToTrafficMarkingPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequest) SetResourceOwnerAccount(v string) *AddTrafficMatchRuleToTrafficMarkingPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequest) SetResourceOwnerId(v int64) *AddTrafficMatchRuleToTrafficMarkingPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequest) SetTrafficMarkingPolicyId(v string) *AddTrafficMatchRuleToTrafficMarkingPolicyRequest {
	s.TrafficMarkingPolicyId = &v
	return s
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequest) SetTrafficMatchRules(v []*AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) *AddTrafficMatchRuleToTrafficMarkingPolicyRequest {
	s.TrafficMatchRules = v
	return s
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequest) Validate() error {
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

type AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules struct {
	// The address type. Valid values: IPv4, IPv6, or empty.
	//
	// example:
	//
	// IPv4
	AddressFamily *string `json:"AddressFamily,omitempty" xml:"AddressFamily,omitempty"`
	// The destination CIDR block of the traffic packet.
	//
	// The traffic classification rule matches traffic whose destination IP address falls within the destination CIDR block. If you do not specify this parameter, the traffic classification rule matches traffic with any destination IP address.
	//
	// example:
	//
	// 10.10.10.0/24
	DstCidr *string `json:"DstCidr,omitempty" xml:"DstCidr,omitempty"`
	// The destination port of the traffic packet. Valid values: **-1*	- and **1*	- to **65535**.
	//
	// The traffic classification rule matches traffic whose destination port falls within the destination port range. If you do not specify this parameter, the traffic classification rule matches traffic with any destination port.
	//
	// You can specify up to 2 port numbers for this parameter. The input format is described as follows:
	//
	// - If you specify only one port number, for example, 1, the system matches traffic whose destination port is 1. If the value is -1, the system matches traffic with any destination port.
	//
	// - If you specify two port numbers, for example, 1 and 200, the system matches traffic whose destination port is in the range of 1 to 200.
	//
	// - If you specify two port numbers and one of them is -1, the other port number must also be -1, which indicates that traffic with any destination port is matched.
	DstPortRange []*int32 `json:"DstPortRange,omitempty" xml:"DstPortRange,omitempty" type:"Repeated"`
	// The Differentiated Services Code Point (DSCP) value of the traffic packet. Valid values: **0*	- to **63**.
	//
	// The traffic classification rule matches traffic that contains the specified DSCP value. If you do not specify this parameter, the traffic classification rule matches traffic with any DSCP value.
	//
	// > The DSCP value refers to the DSCP value that the traffic packet already carries before it enters the inter-region connection.
	//
	// example:
	//
	// 5
	MatchDscp *int32 `json:"MatchDscp,omitempty" xml:"MatchDscp,omitempty"`
	// The protocol type of the traffic packet.
	//
	// The traffic classification rule supports matching traffic of multiple protocol types, such as **HTTP**, **HTTPS**, **TCP**, **UDP**, **SSH**, and **Telnet**. For more protocol types, log on to the [Cloud Enterprise Network (CEN) console](https://cen.console.aliyun.com/cen/list).
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
	// example:
	//
	// HTTP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The source CIDR block of the traffic packet.
	//
	// The traffic classification rule matches traffic whose source IP address falls within the source CIDR block. If you do not specify this parameter, the traffic classification rule matches traffic with any source IP address.
	//
	// example:
	//
	// 192.168.10.0/24
	SrcCidr *string `json:"SrcCidr,omitempty" xml:"SrcCidr,omitempty"`
	// The source port of the traffic packet. Valid values: **-1*	- and **1*	- to **65535**.
	//
	// The traffic classification rule matches traffic whose source port falls within the source port range. If you do not specify this parameter, the traffic classification rule matches traffic with any source port.
	//
	// You can specify up to two port numbers for this parameter. The input format is described as follows:
	//
	// - If you specify only one port number, for example, 1, the system matches traffic whose source port is 1. If the value is -1, the system matches traffic with any source port.
	//
	// - If you specify two port numbers, for example, 1 and 200, the system matches traffic whose source port is in the range of 1 to 200.
	//
	// - If you specify two port numbers and one of them is -1, the other port number must also be -1, which indicates that traffic with any source port is matched.
	SrcPortRange []*int32 `json:"SrcPortRange,omitempty" xml:"SrcPortRange,omitempty" type:"Repeated"`
	// The description of the traffic classification rule.
	//
	// The description can be empty or 1 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// desctest
	TrafficMatchRuleDescription *string `json:"TrafficMatchRuleDescription,omitempty" xml:"TrafficMatchRuleDescription,omitempty"`
	// The name of the traffic classification rule.
	//
	// The name can be empty or 1 to 128 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// nametest
	TrafficMatchRuleName *string `json:"TrafficMatchRuleName,omitempty" xml:"TrafficMatchRuleName,omitempty"`
}

func (s AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) String() string {
	return dara.Prettify(s)
}

func (s AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) GoString() string {
	return s.String()
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) GetAddressFamily() *string {
	return s.AddressFamily
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) GetDstCidr() *string {
	return s.DstCidr
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) GetDstPortRange() []*int32 {
	return s.DstPortRange
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) GetMatchDscp() *int32 {
	return s.MatchDscp
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) GetProtocol() *string {
	return s.Protocol
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) GetSrcCidr() *string {
	return s.SrcCidr
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) GetSrcPortRange() []*int32 {
	return s.SrcPortRange
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) GetTrafficMatchRuleDescription() *string {
	return s.TrafficMatchRuleDescription
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) GetTrafficMatchRuleName() *string {
	return s.TrafficMatchRuleName
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) SetAddressFamily(v string) *AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules {
	s.AddressFamily = &v
	return s
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) SetDstCidr(v string) *AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules {
	s.DstCidr = &v
	return s
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) SetDstPortRange(v []*int32) *AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules {
	s.DstPortRange = v
	return s
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) SetMatchDscp(v int32) *AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules {
	s.MatchDscp = &v
	return s
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) SetProtocol(v string) *AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules {
	s.Protocol = &v
	return s
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) SetSrcCidr(v string) *AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules {
	s.SrcCidr = &v
	return s
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) SetSrcPortRange(v []*int32) *AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules {
	s.SrcPortRange = v
	return s
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) SetTrafficMatchRuleDescription(v string) *AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules {
	s.TrafficMatchRuleDescription = &v
	return s
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) SetTrafficMatchRuleName(v string) *AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules {
	s.TrafficMatchRuleName = &v
	return s
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) Validate() error {
	return dara.Validate(s)
}
