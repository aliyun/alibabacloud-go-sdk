// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrafficMarkingPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTrafficMarkingPoliciesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListTrafficMarkingPoliciesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTrafficMarkingPoliciesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListTrafficMarkingPoliciesResponseBody
	GetTotalCount() *int32
	SetTrafficMarkingPolicies(v []*ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies) *ListTrafficMarkingPoliciesResponseBody
	GetTrafficMarkingPolicies() []*ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies
}

type ListTrafficMarkingPoliciesResponseBody struct {
	// The number of entries returned on each page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the query.
	//
	// 	- If **NextToken*	- was not returned in the previous query, it indicates that no additional results exist.
	//
	// 	- If **NextToken*	- was returned in the previous query, specify the value to obtain the next set of results.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 699989E4-64A0-5F78-8B93-CDB32D98971F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The information about the traffic marking policy.
	TrafficMarkingPolicies []*ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies `json:"TrafficMarkingPolicies,omitempty" xml:"TrafficMarkingPolicies,omitempty" type:"Repeated"`
}

func (s ListTrafficMarkingPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTrafficMarkingPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrafficMarkingPoliciesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTrafficMarkingPoliciesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTrafficMarkingPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTrafficMarkingPoliciesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTrafficMarkingPoliciesResponseBody) GetTrafficMarkingPolicies() []*ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies {
	return s.TrafficMarkingPolicies
}

func (s *ListTrafficMarkingPoliciesResponseBody) SetMaxResults(v int32) *ListTrafficMarkingPoliciesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBody) SetNextToken(v string) *ListTrafficMarkingPoliciesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBody) SetRequestId(v string) *ListTrafficMarkingPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBody) SetTotalCount(v int32) *ListTrafficMarkingPoliciesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBody) SetTrafficMarkingPolicies(v []*ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies) *ListTrafficMarkingPoliciesResponseBody {
	s.TrafficMarkingPolicies = v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBody) Validate() error {
	if s.TrafficMarkingPolicies != nil {
		for _, item := range s.TrafficMarkingPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies struct {
	// The Differentiated Service Code Point (DSCP) value of the traffic marking policy.
	//
	// example:
	//
	// 5
	MarkingDscp *int32 `json:"MarkingDscp,omitempty" xml:"MarkingDscp,omitempty"`
	// The priority of the traffic marking policy.
	//
	// A lower value indicates a higher priority.
	//
	// example:
	//
	// 5
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The description of the traffic marking policy.
	//
	// example:
	//
	// desctest
	TrafficMarkingPolicyDescription *string `json:"TrafficMarkingPolicyDescription,omitempty" xml:"TrafficMarkingPolicyDescription,omitempty"`
	// The ID of the traffic marking policy.
	//
	// example:
	//
	// tm-iz5egnyitxiroq****
	TrafficMarkingPolicyId *string `json:"TrafficMarkingPolicyId,omitempty" xml:"TrafficMarkingPolicyId,omitempty"`
	// The name of the traffic marking policy.
	//
	// example:
	//
	// nametest
	TrafficMarkingPolicyName *string `json:"TrafficMarkingPolicyName,omitempty" xml:"TrafficMarkingPolicyName,omitempty"`
	// The status of the traffic marking policy. Valid values:
	//
	// 	- **Creating**: The policy is being created.
	//
	// 	- **Active**: The policy is available.
	//
	// 	- **Modifying**: The policy is being modified.
	//
	// 	- **Deleting**: The policy is being deleted.
	//
	// example:
	//
	// Creating
	TrafficMarkingPolicyStatus *string `json:"TrafficMarkingPolicyStatus,omitempty" xml:"TrafficMarkingPolicyStatus,omitempty"`
	// The traffic classification rules.
	TrafficMatchRules []*ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules `json:"TrafficMatchRules,omitempty" xml:"TrafficMatchRules,omitempty" type:"Repeated"`
	// The ID of the transit router.
	//
	// example:
	//
	// tr-ccni***
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies) GoString() string {
	return s.String()
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies) GetMarkingDscp() *int32 {
	return s.MarkingDscp
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies) GetPriority() *int32 {
	return s.Priority
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies) GetTrafficMarkingPolicyDescription() *string {
	return s.TrafficMarkingPolicyDescription
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies) GetTrafficMarkingPolicyId() *string {
	return s.TrafficMarkingPolicyId
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies) GetTrafficMarkingPolicyName() *string {
	return s.TrafficMarkingPolicyName
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies) GetTrafficMarkingPolicyStatus() *string {
	return s.TrafficMarkingPolicyStatus
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies) GetTrafficMatchRules() []*ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules {
	return s.TrafficMatchRules
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies) SetMarkingDscp(v int32) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies {
	s.MarkingDscp = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies) SetPriority(v int32) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies {
	s.Priority = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies) SetTrafficMarkingPolicyDescription(v string) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies {
	s.TrafficMarkingPolicyDescription = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies) SetTrafficMarkingPolicyId(v string) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies {
	s.TrafficMarkingPolicyId = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies) SetTrafficMarkingPolicyName(v string) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies {
	s.TrafficMarkingPolicyName = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies) SetTrafficMarkingPolicyStatus(v string) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies {
	s.TrafficMarkingPolicyStatus = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies) SetTrafficMatchRules(v []*ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies {
	s.TrafficMatchRules = v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies) SetTransitRouterId(v string) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies {
	s.TransitRouterId = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies) Validate() error {
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

type ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules struct {
	// The address family. You can set the value to IPv4 or IPv6, or leave the value empty.
	//
	// example:
	//
	// IPv4
	AddressFamily *string `json:"AddressFamily,omitempty" xml:"AddressFamily,omitempty"`
	// The destination CIDR block of packets. IPv4 and IPv6 addresses are supported.
	//
	// example:
	//
	// 192.168.120.0/24
	DstCidr *string `json:"DstCidr,omitempty" xml:"DstCidr,omitempty"`
	// The destination port range used to match data packets.
	DstPortRange []*int32 `json:"DstPortRange,omitempty" xml:"DstPortRange,omitempty" type:"Repeated"`
	// The DSCP value used to match data packets.
	//
	// >  If the value of the **MatchDscp*	- parameter is -1, data packets are considered a match regardless of the DSCP value.
	//
	// example:
	//
	// 6
	MatchDscp *int32 `json:"MatchDscp,omitempty" xml:"MatchDscp,omitempty"`
	// The protocol that is used to match packets.
	//
	// >  Traffic marking policies support multiple protocols. For more information, see the documentation of CEN.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The source CIDR block of packets. IPv6 and IPv4 addresses are supported.
	//
	// example:
	//
	// 192.168.10.0/24
	SrcCidr *string `json:"SrcCidr,omitempty" xml:"SrcCidr,omitempty"`
	// The source port range used to match data packets.
	SrcPortRange []*int32 `json:"SrcPortRange,omitempty" xml:"SrcPortRange,omitempty" type:"Repeated"`
	// The description of the traffic classification rule.
	//
	// example:
	//
	// desctest
	TrafficMatchRuleDescription *string `json:"TrafficMatchRuleDescription,omitempty" xml:"TrafficMatchRuleDescription,omitempty"`
	// The ID of the traffic classification rule.
	//
	// example:
	//
	// tm-rule-fa9kgq1e90rmhc****
	TrafficMatchRuleId *string `json:"TrafficMatchRuleId,omitempty" xml:"TrafficMatchRuleId,omitempty"`
	// The name of the traffic classification rule.
	//
	// example:
	//
	// nametest
	TrafficMatchRuleName *string `json:"TrafficMatchRuleName,omitempty" xml:"TrafficMatchRuleName,omitempty"`
	// The status of the traffic classification rule. Valid values:
	//
	// 	- **Creating**: The rule is being created.
	//
	// 	- **Active**: The rule is available.
	//
	// 	- **Deleting**: The rule is being deleted.
	//
	// example:
	//
	// Creating
	TrafficMatchRuleStatus *string `json:"TrafficMatchRuleStatus,omitempty" xml:"TrafficMatchRuleStatus,omitempty"`
}

func (s ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) String() string {
	return dara.Prettify(s)
}

func (s ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) GoString() string {
	return s.String()
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) GetAddressFamily() *string {
	return s.AddressFamily
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) GetDstCidr() *string {
	return s.DstCidr
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) GetDstPortRange() []*int32 {
	return s.DstPortRange
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) GetMatchDscp() *int32 {
	return s.MatchDscp
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) GetProtocol() *string {
	return s.Protocol
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) GetSrcCidr() *string {
	return s.SrcCidr
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) GetSrcPortRange() []*int32 {
	return s.SrcPortRange
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) GetTrafficMatchRuleDescription() *string {
	return s.TrafficMatchRuleDescription
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) GetTrafficMatchRuleId() *string {
	return s.TrafficMatchRuleId
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) GetTrafficMatchRuleName() *string {
	return s.TrafficMatchRuleName
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) GetTrafficMatchRuleStatus() *string {
	return s.TrafficMatchRuleStatus
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) SetAddressFamily(v string) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules {
	s.AddressFamily = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) SetDstCidr(v string) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules {
	s.DstCidr = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) SetDstPortRange(v []*int32) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules {
	s.DstPortRange = v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) SetMatchDscp(v int32) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules {
	s.MatchDscp = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) SetProtocol(v string) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules {
	s.Protocol = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) SetSrcCidr(v string) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules {
	s.SrcCidr = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) SetSrcPortRange(v []*int32) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules {
	s.SrcPortRange = v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) SetTrafficMatchRuleDescription(v string) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules {
	s.TrafficMatchRuleDescription = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) SetTrafficMatchRuleId(v string) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules {
	s.TrafficMatchRuleId = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) SetTrafficMatchRuleName(v string) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules {
	s.TrafficMatchRuleName = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) SetTrafficMatchRuleStatus(v string) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules {
	s.TrafficMatchRuleStatus = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) Validate() error {
	return dara.Validate(s)
}
