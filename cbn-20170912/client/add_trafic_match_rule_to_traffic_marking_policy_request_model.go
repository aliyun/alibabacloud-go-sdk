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
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, ClientToken is set to the value of RequestId. The value of RequestId may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// 	- **true**: performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and sends the request.
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
	// The information about the traffic classification rule.
	//
	// You can specify at most 50 traffic classification rules.
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
	return dara.Validate(s)
}

type AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules struct {
	// The destination CIDR block that is used to match packets.
	//
	// The traffic classification rule matches the packets whose destination IP addresses fall within the specified destination CIDR block. If you do not set this parameter, packets are considered a match regardless of the DSCP value.
	//
	// You can specify at most 50 traffic classification rules.
	//
	// example:
	//
	// 10.10.10.0/24
	DstCidr *string `json:"DstCidr,omitempty" xml:"DstCidr,omitempty"`
	// The destination port range that is used to match packets. Valid values: **-1*	- and **1*	- to **65535**.
	//
	// The traffic classification rule matches the packets whose destination ports fall within the destination port range. If you do not set this parameter, packets are considered a match regardless of the DSCP value.
	//
	// You can specify at most two ports. Take note of the following rules:
	//
	// 	- If you enter only one port number such as 1, the system matches the packets whose destination port is port 1.
	//
	// 	- If you enter two port numbers such as 1 and 200, the system matches the packets whose destination ports fall between 1 and 200.
	//
	// 	- If you enter two port numbers and one of them is -1, the other port number must also be -1. In this case, packets are considered a match regardless of the destination port.
	//
	// You can specify at most 50 traffic classification rules.
	DstPortRange []*int32 `json:"DstPortRange,omitempty" xml:"DstPortRange,omitempty" type:"Repeated"`
	// The differentiated services code point (DSCP) value that is used to match packets. Valid values: **0*	- to **63**.
	//
	// The traffic classification rule matches the packets that contain the specified DSCP value. If you do not set this parameter, packets are considered a match regardless of the DSCP value.
	//
	// >  The DSCP value that you specify for this parameter is the DSCP value that packets carry before they are transmitted over the inter-region connection.
	//
	// You can specify at most 50 traffic classification rules.
	//
	// example:
	//
	// 5
	MatchDscp *int32 `json:"MatchDscp,omitempty" xml:"MatchDscp,omitempty"`
	// The protocol that is used to match packets.
	//
	// Valid values: **HTTP**, **HTTPS**, **TCP**, **UDP**, **SSH**, and **Telnet**. For more information, log on to the [Cloud Enterprise Network (CEN) console](https://cen.console.aliyun.com/cen/list).
	//
	// You can specify at most 50 traffic classification rules.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The source CIDR block that is used to match packets.
	//
	// The traffic classification rule matches the packets whose source IP addresses fall within the specified source CIDR block. If you do not set this parameter, packets are considered a match regardless of the source IP address.
	//
	// You can specify at most 50 traffic classification rules.
	//
	// example:
	//
	// 192.168.10.0/24
	SrcCidr *string `json:"SrcCidr,omitempty" xml:"SrcCidr,omitempty"`
	// The source port range that is used to match packets. Valid values: **-1*	- and **1*	- to **65535**.
	//
	// The traffic classification rule matches the packets whose source ports fall within the source port range. If you do not set this parameter, packets are considered a match regardless of the source port.
	//
	// You can specify at most two ports. Take note of the following rules:
	//
	// 	- If you enter only one port number such as 1, the system matches the packets whose source port is 1.
	//
	// 	- If you enter two port numbers such as 1 and 200, the system matches the packets whose source ports fall between 1 and 200.
	//
	// 	- If you enter two port numbers and one of them is -1, the other port number must also be -1. In this case, packets are considered a match regardless of the source port.
	//
	// You can specify at most 50 traffic classification rules.
	SrcPortRange []*int32 `json:"SrcPortRange,omitempty" xml:"SrcPortRange,omitempty" type:"Repeated"`
	// The description of the traffic classification rule.
	//
	// The description must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The description must start with a letter.
	//
	// You can specify at most 50 traffic classification rules.
	//
	// example:
	//
	// desctest
	TrafficMatchRuleDescription *string `json:"TrafficMatchRuleDescription,omitempty" xml:"TrafficMatchRuleDescription,omitempty"`
	// The name of the traffic classification rule.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). It must start with a letter.
	//
	// You can specify at most 50 traffic classification rules.
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
