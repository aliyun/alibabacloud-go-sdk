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
	// The list of traffic classification rules to add.
	//
	// You can add up to 50 traffic classification rules at a time.
	AddTrafficMatchRules []*UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules `json:"AddTrafficMatchRules,omitempty" xml:"AddTrafficMatchRules,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// The client generates the value of this parameter. Ensure that the value is unique among different requests. The token can be up to 64 ASCII characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The list of traffic classification rules to delete.
	//
	// >Provide as much information as possible for the traffic classification rules, such as the source CIDR block, destination CIDR block, source port, destination port, and DSCP value. Otherwise, the system may fail to locate the target traffic classification rules and will not delete them.
	DeleteTrafficMatchRules []*UpdateTrafficMarkingPolicyAttributeRequestDeleteTrafficMatchRules `json:"DeleteTrafficMatchRules,omitempty" xml:"DeleteTrafficMatchRules,omitempty" type:"Repeated"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run. The system checks the required parameters, request syntax, and business restrictions without modifying the name, description, or traffic classification rules of the traffic marking policy. If the request fails the dry run, the corresponding error message is returned. If the request passes the dry run, the error code `DryRunOperation` is returned.
	//
	// - **false*	- (default): performs a dry run and sends the request. After the request passes the dry run, the name, description, and traffic classification rules of the traffic marking policy are directly modified.
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
	// The description can be empty or 1 to 256 characters in length and cannot start with http:// or https://.
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
	// The name can be empty or 1 to 128 characters in length and cannot start with http:// or https://.
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
	if s.AddTrafficMatchRules != nil {
		for _, item := range s.AddTrafficMatchRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DeleteTrafficMatchRules != nil {
		for _, item := range s.DeleteTrafficMatchRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateTrafficMarkingPolicyAttributeRequestAddTrafficMatchRules struct {
	// The address family. Valid values: IPv4, IPv6, or empty.
	//
	// example:
	//
	// IPv4
	AddressFamily *string `json:"AddressFamily,omitempty" xml:"AddressFamily,omitempty"`
	// The destination CIDR block of the traffic packet. IPv4 and IPv6 addresses are supported.
	//
	// The traffic classification rule matches traffic whose destination IP address falls within the destination CIDR block. If you do not set this parameter, the traffic classification rule matches traffic with any destination IP address.
	//
	// You can add up to 50 traffic classification rules at a time. Each traffic classification rule can specify one destination CIDR block.
	//
	// example:
	//
	// 172.30.0.0/24
	DstCidr *string `json:"DstCidr,omitempty" xml:"DstCidr,omitempty"`
	// The destination port of the traffic packet. Valid values: **-1*	- and **1*	- to **65535**.
	//
	// The traffic classification rule matches traffic whose destination port falls within the destination port range. If you do not set this parameter, the traffic classification rule matches traffic with any destination port.
	//
	// This parameter supports up to two port numbers. The input format is described as follows:
	//
	// - If you enter only one port number, such as 1, the system matches traffic whose destination port is 1 by default. If the value is -1, the system matches traffic with any destination port.
	//
	// - If you enter two port numbers, such as 1 and 200, the system matches traffic whose destination port is in the range of 1 to 200 by default.
	//
	// - If you enter two port numbers and one of them is -1, the other port number must also be -1, which indicates that traffic with any destination port is matched.
	//
	// You can add up to 50 traffic classification rules at a time. Each traffic classification rule can specify one destination port range.
	DstPortRange []*int32 `json:"DstPortRange,omitempty" xml:"DstPortRange,omitempty" type:"Repeated"`
	// The Differentiated Services Code Point (DSCP) value of the traffic packet. Valid values: **0*	- to **63**.
	//
	// The traffic classification rule matches traffic that contains the specified DSCP value. If you do not set this parameter, the traffic classification rule matches traffic with any DSCP value.
	//
	// > The DSCP value refers to the DSCP value that the traffic packet already carries before entering the inter-region connection.
	//
	// You can add up to 50 traffic classification rules at a time. Each traffic classification rule can match one DSCP value.
	//
	// example:
	//
	// 1
	MatchDscp *int32 `json:"MatchDscp,omitempty" xml:"MatchDscp,omitempty"`
	// The protocol type of the traffic packet.
	//
	// The traffic marking policy supports matching traffic of multiple protocol types, such as **HTTP**, **HTTPS**, **TCP**, **UDP**, **SSH**, and **Telnet**. For more protocol types, log on to the [Cloud Enterprise Network (CEN) console](https://cen.console.aliyun.com/cen/list).
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
	// You can add up to 50 traffic classification rules at a time. Each traffic classification rule can match one protocol type.
	//
	// example:
	//
	// UDP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The source CIDR block of the traffic packet. IPv4 and IPv6 addresses are supported.
	//
	// The traffic classification rule matches traffic whose source IP address falls within the source CIDR block. If you do not set this parameter, the traffic classification rule matches traffic with any source IP address.
	//
	// You can add up to 50 traffic classification rules at a time. Each traffic classification rule can match one source CIDR block.
	//
	// example:
	//
	// 10.128.32.0/19
	SrcCidr *string `json:"SrcCidr,omitempty" xml:"SrcCidr,omitempty"`
	// The source port of the traffic packet. Valid values: **-1*	- and **1*	- to **65535**.
	//
	// The traffic classification rule matches traffic whose source port falls within the source port range. If you do not set this parameter, the traffic classification rule matches traffic with any source port.
	//
	// This parameter supports up to two port numbers. The input format is described as follows:
	//
	// - If you enter only one port number, such as 1, the system matches traffic whose source port is 1 by default. If the value is -1, the system matches traffic with any source port.
	//
	// - If you enter two port numbers, such as 1 and 200, the system matches traffic whose source port is in the range of 1 to 200 by default.
	//
	// - If you enter two port numbers and one of them is -1, the other port number must also be -1, which indicates that traffic with any source port is matched.
	//
	// You can add up to 50 traffic classification rules at a time. Each traffic classification rule can specify one source port range.
	SrcPortRange []*int32 `json:"SrcPortRange,omitempty" xml:"SrcPortRange,omitempty" type:"Repeated"`
	// The description of the traffic classification rule.
	//
	// You can add up to 50 traffic classification rules at a time. Each traffic classification rule can have one description.
	//
	// The description can be empty or 1 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// test1
	TrafficMatchRuleDescription *string `json:"TrafficMatchRuleDescription,omitempty" xml:"TrafficMatchRuleDescription,omitempty"`
	// The name of the traffic classification rule.
	//
	// You can add up to 50 traffic classification rules at a time. Each traffic classification rule can have one name.
	//
	// The name can be empty or 1 to 128 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// Guangzhou test
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
	// The address family. Valid values: IPv4, IPv6, or empty.
	//
	// example:
	//
	// IPv6
	AddressFamily *string `json:"AddressFamily,omitempty" xml:"AddressFamily,omitempty"`
	// The destination CIDR block of the traffic packet. IPv4 and IPv6 addresses are supported.
	//
	// example:
	//
	// 192.168.200.3/32
	DstCidr *string `json:"DstCidr,omitempty" xml:"DstCidr,omitempty"`
	// The destination port of the traffic packet.
	DstPortRange []*int32 `json:"DstPortRange,omitempty" xml:"DstPortRange,omitempty" type:"Repeated"`
	// The DSCP value of the traffic packet.
	//
	// example:
	//
	// 3
	MatchDscp *int32 `json:"MatchDscp,omitempty" xml:"MatchDscp,omitempty"`
	// The protocol type of the traffic packet.
	//
	// You can call the [ListTrafficMarkingPolicies](https://help.aliyun.com/document_detail/468322.html) operation to query the details of the traffic classification rules that you want to delete.
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The source CIDR block of the traffic packet. IPv4 and IPv6 addresses are supported.
	//
	// example:
	//
	// 10.72.0.0/16
	SrcCidr *string `json:"SrcCidr,omitempty" xml:"SrcCidr,omitempty"`
	// The source port of the traffic packet.
	SrcPortRange []*int32 `json:"SrcPortRange,omitempty" xml:"SrcPortRange,omitempty" type:"Repeated"`
	// The description of the traffic classification rule.
	//
	// The description can be empty or 1 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// Hangzhou-Qingdao CAT
	TrafficMatchRuleDescription *string `json:"TrafficMatchRuleDescription,omitempty" xml:"TrafficMatchRuleDescription,omitempty"`
	// The name of the traffic classification rule.
	//
	// The name can be empty or 1 to 128 characters in length and cannot start with http:// or https://.
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
