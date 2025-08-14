// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTrafficMarkingPolicyAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddTrafficMatchRules(v []*UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules) *UpdateTrafficMarkingPolicyAttributeRequest
	GetAddTrafficMatchRules() []*UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules
	SetClientToken(v string) *UpdateTrafficMarkingPolicyAttributeRequest
	GetClientToken() *string
	SetDeleteTrafficMatchRules(v []*UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules) *UpdateTrafficMarkingPolicyAttributeRequest
	GetDeleteTrafficMatchRules() []*UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules
	SetDryRun(v bool) *UpdateTrafficMarkingPolicyAttributeRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *UpdateTrafficMarkingPolicyAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateTrafficMarkingPolicyAttributeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateTrafficMarkingPolicyAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateTrafficMarkingPolicyAttributeRequest
	GetResourceOwnerId() *int64
	SetTrafficMarkingPolicyDescription(v string) *UpdateTrafficMarkingPolicyAttributeRequest
	GetTrafficMarkingPolicyDescription() *string
	SetTrafficMarkingPolicyId(v string) *UpdateTrafficMarkingPolicyAttributeRequest
	GetTrafficMarkingPolicyId() *string
	SetTrafficMarkingPolicyName(v string) *UpdateTrafficMarkingPolicyAttributeRequest
	GetTrafficMarkingPolicyName() *string
}

type UpdateTrafficMarkingPolicyAttributeRequest struct {
	// The traffic classification rules to be added to the traffic marking policy.
	//
	// You can add at most 50 traffic classification rules in each call.
	AddTrafficMatchRules []*UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules `json:"AddTrafficMatchRules,omitempty" xml:"AddTrafficMatchRules,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The traffic classification rules to be deleted from the traffic marking policy.
	//
	// >  Specify detailed information about the traffic classification rule, such as the source CIDR block, destination CIDR block, source port, destination port, and DSCP value. If you do not specify sufficient information, the system may fail to match the traffic classification rule that you want to delete.
	DeleteTrafficMatchRules []*UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules `json:"DeleteTrafficMatchRules,omitempty" xml:"DeleteTrafficMatchRules,omitempty" type:"Repeated"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// 	- **true**: preforms a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
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
	// The new description of the traffic marking policy.
	//
	// The description must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The description must start with a letter.
	//
	// example:
	//
	// desctest
	TrafficMarkingPolicyDescription *string `json:"TrafficMarkingPolicyDescription,omitempty" xml:"TrafficMarkingPolicyDescription,omitempty"`
	// The ID of the traffic marking policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// tm-u9nxup5kww5po8****
	TrafficMarkingPolicyId *string `json:"TrafficMarkingPolicyId,omitempty" xml:"TrafficMarkingPolicyId,omitempty"`
	// The new name of the traffic marking policy.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). It must start with a letter.
	//
	// example:
	//
	// nametest
	TrafficMarkingPolicyName *string `json:"TrafficMarkingPolicyName,omitempty" xml:"TrafficMarkingPolicyName,omitempty"`
}

func (s UpdateTrafficMarkingPolicyAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTrafficMarkingPolicyAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateTrafficMarkingPolicyAttributeRequest) GetAddTrafficMatchRules() []*UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules {
	return s.AddTrafficMatchRules
}

func (s *UpdateTrafficMarkingPolicyAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateTrafficMarkingPolicyAttributeRequest) GetDeleteTrafficMatchRules() []*UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules {
	return s.DeleteTrafficMatchRules
}

func (s *UpdateTrafficMarkingPolicyAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateTrafficMarkingPolicyAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateTrafficMarkingPolicyAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateTrafficMarkingPolicyAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateTrafficMarkingPolicyAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateTrafficMarkingPolicyAttributeRequest) GetTrafficMarkingPolicyDescription() *string {
	return s.TrafficMarkingPolicyDescription
}

func (s *UpdateTrafficMarkingPolicyAttributeRequest) GetTrafficMarkingPolicyId() *string {
	return s.TrafficMarkingPolicyId
}

func (s *UpdateTrafficMarkingPolicyAttributeRequest) GetTrafficMarkingPolicyName() *string {
	return s.TrafficMarkingPolicyName
}

func (s *UpdateTrafficMarkingPolicyAttributeRequest) SetAddTrafficMatchRules(v []*UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules) *UpdateTrafficMarkingPolicyAttributeRequest {
	s.AddTrafficMatchRules = v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequest) SetClientToken(v string) *UpdateTrafficMarkingPolicyAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequest) SetDeleteTrafficMatchRules(v []*UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules) *UpdateTrafficMarkingPolicyAttributeRequest {
	s.DeleteTrafficMatchRules = v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequest) SetDryRun(v bool) *UpdateTrafficMarkingPolicyAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequest) SetOwnerAccount(v string) *UpdateTrafficMarkingPolicyAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequest) SetOwnerId(v int64) *UpdateTrafficMarkingPolicyAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequest) SetResourceOwnerAccount(v string) *UpdateTrafficMarkingPolicyAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequest) SetResourceOwnerId(v int64) *UpdateTrafficMarkingPolicyAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequest) SetTrafficMarkingPolicyDescription(v string) *UpdateTrafficMarkingPolicyAttributeRequest {
	s.TrafficMarkingPolicyDescription = &v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequest) SetTrafficMarkingPolicyId(v string) *UpdateTrafficMarkingPolicyAttributeRequest {
	s.TrafficMarkingPolicyId = &v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequest) SetTrafficMarkingPolicyName(v string) *UpdateTrafficMarkingPolicyAttributeRequest {
	s.TrafficMarkingPolicyName = &v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules struct {
	// The address family. Valid values: You can set the value to IPv4 or IPv6, or leave the value empty.
	//
	// example:
	//
	// IPv4
	AddressFamily *string `json:"AddressFamily,omitempty" xml:"AddressFamily,omitempty"`
	// The destination CIDR block of packets. IPv4 and IPv6 addresses are supported.
	//
	// Packets whose destination IP addresses fall into the specified destination CIDR block meet the traffic classification rule. If you do not specify a destination CIDR block, all packets meet the traffic classification rule.
	//
	// You can create up to 50 traffic classification rules in each call. You can specify a destination CIDR block for each traffic classification rule.
	//
	// example:
	//
	// 172.30.0.0/24
	DstCidr *string `json:"DstCidr,omitempty" xml:"DstCidr,omitempty"`
	// The destination port range that is used to match packets. Valid values: **-1*	- and **1*	- to **65535**.
	//
	// Packets whose destination ports fall within the specified destination port range are considered a match. If you do not specify a destination port range, packets are considered a match regardless of the destination port.
	//
	// You can enter up to two port numbers. Take note of the following rules:
	//
	// 	- If you enter only one port number, such as 1, packets whose destination port is 1 match the traffic classification rule. A value of -1 specifies all destination ports.
	//
	// 	- If you enter two port numbers, such as 1 and 200, packets whose destination ports fall into 1 and 200 are considered a match.
	//
	// 	- If you enter two port numbers and one of them is -1, the other port number must also be -1. In this case, all packets meet the traffic classification rule.
	//
	// You can create up to 50 traffic classification rules in each call. You can specify a destination port range for each traffic classification rule.
	DstPortRange []*int32 `json:"DstPortRange,omitempty" xml:"DstPortRange,omitempty" type:"Repeated"`
	// The Differentiated Service Code Point (DSCP) value that is used to match packets. Valid values: **0*	- to **63**.
	//
	// Requests that carry the specified DSCP value are considered a match. If you do not specify a DSCP value, packets are considered a match regardless of the DSCP value.
	//
	// >  The DSCP value that you specify for this parameter is the DSCP value that packets carry before they are transmitted over the inter-region connection.
	//
	// You can create up to 50 traffic classification rules in each call. You can specify a DSCP value for each traffic classification rule.
	//
	// example:
	//
	// 1
	MatchDscp *int32 `json:"MatchDscp,omitempty" xml:"MatchDscp,omitempty"`
	// The protocol that is used to match packets.
	//
	// Traffic classification rules support the following protocols: **HTTP**, **HTTPS**, **TCP**, **UDP**, **SSH**, and **Telnet**. For more information, log on to the [CEN console](https://cen.console.aliyun.com/cen/list).
	//
	// **Some protocols use a fixed port. Click to view the protocols and ports.**
	//
	// 	- If the protocol is **ICMP**, set the destination port to **-1**.
	//
	// 	- If the protocol is **GRE**, set the destination port to **-1**.
	//
	// 	- If the protocol is **SSH**, set the destination port to **22**.
	//
	// 	- If the protocol is **Telnet**, set the destination port to **23**.
	//
	// 	- If the protocol is **HTTP**, set the destination port to **80**.
	//
	// 	- If the protocol is **HTTPS**, set the destination port to **443**.
	//
	// 	- If the protocol is **MS SQL**, set the destination port to **1443**.
	//
	// 	- If the protocol is **Oracle**, set the destination port to **1521**.
	//
	// 	- If the protocol is **Mysql**, set the destination port to **3306**.
	//
	// 	- If the protocol is **RDP**, set the destination port to **3389**.
	//
	// 	- If the protocol is **Postgre SQL**, set the destination port to **5432**.
	//
	// 	- If the protocol is **Redis**, the destination port must be **6379**.
	//
	// You can create up to 50 traffic classification rules in each call. You can specify a protocol for each traffic classification rule.
	//
	// example:
	//
	// UDP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The source CIDR block of packets. IPv4 and IPv6 addresses are supported.
	//
	// Packets whose source IP addresses fall into the specified source CIDR block meet the traffic classification rule. If you do not specify a source CIDR block, all packets meet the traffic classification rule.
	//
	// You can create up to 50 traffic classification rules in each call. You can specify a source CIDR block for each traffic classification rule.
	//
	// example:
	//
	// 10.128.32.0/19
	SrcCidr *string `json:"SrcCidr,omitempty" xml:"SrcCidr,omitempty"`
	// The source port range that is used to match packets. Valid values: **-1*	- and **1*	- to **65535**.
	//
	// The traffic classification rule matches the packets whose source ports fall within the source port range. If you do not specify this parameter, packets are considered a match regardless of the source port.
	//
	// You can enter up to two port numbers. Take note of the following rules:
	//
	// 	- If you enter only one port number, such as 1, packets whose source port is 1 are considered a match. A value of -1 specifies all source ports.
	//
	// 	- If you enter two port numbers, such as 1 and 200, packets whose source ports fall into 1 and 200 are considered a match.
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
	// test1
	TrafficMatchRuleDescription *string `json:"TrafficMatchRuleDescription,omitempty" xml:"TrafficMatchRuleDescription,omitempty"`
	// The name of the traffic classification rule.
	//
	// You can create up to 50 traffic classification rules in each call. You can specify a name for each traffic classification rule.
	//
	// The name can be empty or 1 to 128 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// Guangzhou Testing
	TrafficMatchRuleName *string `json:"TrafficMatchRuleName,omitempty" xml:"TrafficMatchRuleName,omitempty"`
}

func (s UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules) GoString() string {
	return s.String()
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules) GetAddressFamily() *string {
	return s.AddressFamily
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules) GetDstCidr() *string {
	return s.DstCidr
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules) GetDstPortRange() []*int32 {
	return s.DstPortRange
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules) GetMatchDscp() *int32 {
	return s.MatchDscp
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules) GetSrcCidr() *string {
	return s.SrcCidr
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules) GetSrcPortRange() []*int32 {
	return s.SrcPortRange
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules) GetTrafficMatchRuleDescription() *string {
	return s.TrafficMatchRuleDescription
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules) GetTrafficMatchRuleName() *string {
	return s.TrafficMatchRuleName
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules) SetAddressFamily(v string) *UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules {
	s.AddressFamily = &v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules) SetDstCidr(v string) *UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules {
	s.DstCidr = &v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules) SetDstPortRange(v []*int32) *UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules {
	s.DstPortRange = v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules) SetMatchDscp(v int32) *UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules {
	s.MatchDscp = &v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules) SetProtocol(v string) *UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules {
	s.Protocol = &v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules) SetSrcCidr(v string) *UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules {
	s.SrcCidr = &v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules) SetSrcPortRange(v []*int32) *UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules {
	s.SrcPortRange = v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules) SetTrafficMatchRuleDescription(v string) *UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules {
	s.TrafficMatchRuleDescription = &v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules) SetTrafficMatchRuleName(v string) *UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules {
	s.TrafficMatchRuleName = &v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules) Validate() error {
	return dara.Validate(s)
}

type UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules struct {
	// The address family. Valid values: You can set the value to IPv4 or IPv6, or leave the value empty.
	//
	// example:
	//
	// IPv6
	AddressFamily *string `json:"AddressFamily,omitempty" xml:"AddressFamily,omitempty"`
	// The destination CIDR block of packets. IPv4 and IPv6 addresses are supported.
	//
	// example:
	//
	// 192.168.200.3/32
	DstCidr *string `json:"DstCidr,omitempty" xml:"DstCidr,omitempty"`
	// The destination port range that is used to match packets.
	DstPortRange []*int32 `json:"DstPortRange,omitempty" xml:"DstPortRange,omitempty" type:"Repeated"`
	// The DSCP value that is used to match packets.
	//
	// example:
	//
	// 3
	MatchDscp *int32 `json:"MatchDscp,omitempty" xml:"MatchDscp,omitempty"`
	// The protocol that is used to match packets.
	//
	// You can call the [ListTrafficMarkingPolicies](https://help.aliyun.com/document_detail/468322.html) operation to query the details about a traffic classification rule.
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The source CIDR block of packets. IPv4 and IPv6 addresses are supported.
	//
	// example:
	//
	// 10.72.0.0/16
	SrcCidr *string `json:"SrcCidr,omitempty" xml:"SrcCidr,omitempty"`
	// The source port range that is used to match packets.
	SrcPortRange []*int32 `json:"SrcPortRange,omitempty" xml:"SrcPortRange,omitempty" type:"Repeated"`
	// The description of the traffic classification rule.
	//
	// This parameter is optional. If you enter a description, it must be 1 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// Hangzhou-to-Qingdao CAT
	TrafficMatchRuleDescription *string `json:"TrafficMatchRuleDescription,omitempty" xml:"TrafficMatchRuleDescription,omitempty"`
	// The name of the traffic classification rule.
	//
	// The name can be empty or 1 to 128 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// test
	TrafficMatchRuleName *string `json:"TrafficMatchRuleName,omitempty" xml:"TrafficMatchRuleName,omitempty"`
}

func (s UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules) GoString() string {
	return s.String()
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules) GetAddressFamily() *string {
	return s.AddressFamily
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules) GetDstCidr() *string {
	return s.DstCidr
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules) GetDstPortRange() []*int32 {
	return s.DstPortRange
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules) GetMatchDscp() *int32 {
	return s.MatchDscp
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules) GetSrcCidr() *string {
	return s.SrcCidr
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules) GetSrcPortRange() []*int32 {
	return s.SrcPortRange
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules) GetTrafficMatchRuleDescription() *string {
	return s.TrafficMatchRuleDescription
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules) GetTrafficMatchRuleName() *string {
	return s.TrafficMatchRuleName
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules) SetAddressFamily(v string) *UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules {
	s.AddressFamily = &v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules) SetDstCidr(v string) *UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules {
	s.DstCidr = &v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules) SetDstPortRange(v []*int32) *UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules {
	s.DstPortRange = v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules) SetMatchDscp(v int32) *UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules {
	s.MatchDscp = &v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules) SetProtocol(v string) *UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules {
	s.Protocol = &v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules) SetSrcCidr(v string) *UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules {
	s.SrcCidr = &v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules) SetSrcPortRange(v []*int32) *UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules {
	s.SrcPortRange = v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules) SetTrafficMatchRuleDescription(v string) *UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules {
	s.TrafficMatchRuleDescription = &v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules) SetTrafficMatchRuleName(v string) *UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules {
	s.TrafficMatchRuleName = &v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules) Validate() error {
	return dara.Validate(s)
}
