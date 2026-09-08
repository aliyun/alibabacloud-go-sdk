// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTraficMatchRuleToTrafficMarkingPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AddTraficMatchRuleToTrafficMarkingPolicyRequest
	GetClientToken() *string
	SetDryRun(v bool) *AddTraficMatchRuleToTrafficMarkingPolicyRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *AddTraficMatchRuleToTrafficMarkingPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddTraficMatchRuleToTrafficMarkingPolicyRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AddTraficMatchRuleToTrafficMarkingPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddTraficMatchRuleToTrafficMarkingPolicyRequest
	GetResourceOwnerId() *int64
	SetTrafficMarkingPolicyId(v string) *AddTraficMatchRuleToTrafficMarkingPolicyRequest
	GetTrafficMarkingPolicyId() *string
	SetTrafficMatchRules(v []*AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) *AddTraficMatchRuleToTrafficMarkingPolicyRequest
	GetTrafficMatchRules() []*AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules
}

type AddTraficMatchRuleToTrafficMarkingPolicyRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId value as the client token. The RequestId value may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run. The system checks the required parameters, request syntax, and business restrictions without adding a traffic classification rule to the traffic marking policy. If the check fails, the corresponding error is returned. If the check passes, the error code `DryRunOperation` is returned.
	//
	// - **false*	- (default): performs the actual request. After the check passes, a traffic classification rule is directly added to the traffic marking policy.
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
	TrafficMatchRules []*AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules `json:"TrafficMatchRules,omitempty" xml:"TrafficMatchRules,omitempty" type:"Repeated"`
}

func (s AddTraficMatchRuleToTrafficMarkingPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s AddTraficMatchRuleToTrafficMarkingPolicyRequest) GoString() string {
	return s.String()
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequest) GetTrafficMarkingPolicyId() *string {
	return s.TrafficMarkingPolicyId
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequest) GetTrafficMatchRules() []*AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules {
	return s.TrafficMatchRules
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequest) SetClientToken(v string) *AddTraficMatchRuleToTrafficMarkingPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequest) SetDryRun(v bool) *AddTraficMatchRuleToTrafficMarkingPolicyRequest {
	s.DryRun = &v
	return s
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequest) SetOwnerAccount(v string) *AddTraficMatchRuleToTrafficMarkingPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequest) SetOwnerId(v int64) *AddTraficMatchRuleToTrafficMarkingPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequest) SetResourceOwnerAccount(v string) *AddTraficMatchRuleToTrafficMarkingPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequest) SetResourceOwnerId(v int64) *AddTraficMatchRuleToTrafficMarkingPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequest) SetTrafficMarkingPolicyId(v string) *AddTraficMatchRuleToTrafficMarkingPolicyRequest {
	s.TrafficMarkingPolicyId = &v
	return s
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequest) SetTrafficMatchRules(v []*AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) *AddTraficMatchRuleToTrafficMarkingPolicyRequest {
	s.TrafficMatchRules = v
	return s
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequest) Validate() error {
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

type AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules struct {
	// The destination CIDR block of the traffic packet.
	//
	// The traffic classification rule matches traffic whose destination IP address falls within the destination CIDR block. If you do not set this parameter, the traffic classification rule matches traffic with any destination IP address.
	//
	// You can add up to 50 traffic classification rules at a time.
	//
	// example:
	//
	// 10.10.10.0/24
	DstCidr *string `json:"DstCidr,omitempty" xml:"DstCidr,omitempty"`
	// The destination port of the traffic packet. Valid values: **-1*	- and **1*	- to **65535**.
	//
	// The traffic classification rule matches traffic whose destination port number falls within the destination port range. If you do not set this parameter, the traffic classification rule matches traffic with any destination port number.
	//
	// This parameter supports up to two port numbers. The input format is described as follows:
	//
	// - If you enter only one port number, such as 1, the system matches traffic whose destination port is 1 by default.
	//
	// - If you enter two port numbers, such as 1 and 200, the system matches traffic whose destination port falls within the range of 1 to 200 by default.
	//
	// - If you enter two port numbers and one of them is -1, the other port number must also be -1, which indicates that traffic with any destination port is matched.
	//
	// You can add up to 50 traffic classification rules at a time.
	DstPortRange []*int32 `json:"DstPortRange,omitempty" xml:"DstPortRange,omitempty" type:"Repeated"`
	// The Differentiated Services Code Point (DSCP) value of the traffic packet. Valid values: **0*	- to **63**.
	//
	// The traffic classification rule matches traffic that contains the specified DSCP value. If you do not set this parameter, the traffic classification rule matches traffic with any DSCP value.
	//
	// > The DSCP value refers to the DSCP value that the traffic packet already carries before entering the inter-region connection.
	//
	// You can add up to 50 traffic classification rules at a time.
	//
	// example:
	//
	// 5
	MatchDscp *int32 `json:"MatchDscp,omitempty" xml:"MatchDscp,omitempty"`
	// The protocol type of the traffic packet.
	//
	// The traffic classification rule supports matching traffic of multiple protocol types, such as **HTTP**, **HTTPS**, **TCP**, **UDP**, **SSH**, and **Telnet**. For more protocol types, log on to the [Cloud Enterprise Network (CEN) console](https://cen.console.aliyun.com/cen/list).
	//
	// You can add up to 50 traffic classification rules at a time.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The source CIDR block of the traffic packet.
	//
	// The traffic classification rule matches traffic whose source IP address falls within the source CIDR block. If you do not set this parameter, the traffic classification rule matches traffic with any source IP address.
	//
	// You can add up to 50 traffic classification rules at a time.
	//
	// example:
	//
	// 192.168.10.0/24
	SrcCidr *string `json:"SrcCidr,omitempty" xml:"SrcCidr,omitempty"`
	// The source port of the traffic packet. Valid values: **-1*	- and **1*	- to **65535**.
	//
	// The traffic classification rule matches traffic whose source port number falls within the source port range. If you do not set this parameter, the traffic classification rule matches traffic with any source port number.
	//
	// This parameter supports up to two port numbers. The input format is described as follows:
	//
	// - If you enter only one port number, such as 1, the system matches traffic whose source port is 1 by default.
	//
	// - If you enter two port numbers, such as 1 and 200, the system matches traffic whose source port falls within the range of 1 to 200 by default.
	//
	// - If you enter two port numbers and one of them is -1, the other port number must also be -1, which indicates that traffic with any source port is matched.
	//
	// You can add up to 50 traffic classification rules at a time.
	SrcPortRange []*int32 `json:"SrcPortRange,omitempty" xml:"SrcPortRange,omitempty" type:"Repeated"`
	// The description of the traffic classification rule.
	//
	// You can add up to 50 traffic classification rules at a time.
	//
	// example:
	//
	// desctest
	TrafficMatchRuleDescription *string `json:"TrafficMatchRuleDescription,omitempty" xml:"TrafficMatchRuleDescription,omitempty"`
	// The name of the traffic classification rule.
	//
	// You can add up to 50 traffic classification rules at a time.
	//
	// example:
	//
	// nametest
	TrafficMatchRuleName *string `json:"TrafficMatchRuleName,omitempty" xml:"TrafficMatchRuleName,omitempty"`
}

func (s AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) String() string {
	return dara.Prettify(s)
}

func (s AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) GoString() string {
	return s.String()
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) GetDstCidr() *string {
	return s.DstCidr
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) GetDstPortRange() []*int32 {
	return s.DstPortRange
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) GetMatchDscp() *int32 {
	return s.MatchDscp
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) GetProtocol() *string {
	return s.Protocol
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) GetSrcCidr() *string {
	return s.SrcCidr
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) GetSrcPortRange() []*int32 {
	return s.SrcPortRange
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) GetTrafficMatchRuleDescription() *string {
	return s.TrafficMatchRuleDescription
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) GetTrafficMatchRuleName() *string {
	return s.TrafficMatchRuleName
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) SetDstCidr(v string) *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules {
	s.DstCidr = &v
	return s
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) SetDstPortRange(v []*int32) *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules {
	s.DstPortRange = v
	return s
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) SetMatchDscp(v int32) *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules {
	s.MatchDscp = &v
	return s
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) SetProtocol(v string) *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules {
	s.Protocol = &v
	return s
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) SetSrcCidr(v string) *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules {
	s.SrcCidr = &v
	return s
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) SetSrcPortRange(v []*int32) *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules {
	s.SrcPortRange = v
	return s
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) SetTrafficMatchRuleDescription(v string) *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules {
	s.TrafficMatchRuleDescription = &v
	return s
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) SetTrafficMatchRuleName(v string) *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules {
	s.TrafficMatchRuleName = &v
	return s
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) Validate() error {
	return dara.Validate(s)
}
