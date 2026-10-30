// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRouteMapRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListRouteMapRulesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListRouteMapRulesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListRouteMapRulesResponseBody
	GetRequestId() *string
	SetRouteMapRules(v []*ListRouteMapRulesResponseBodyRouteMapRules) *ListRouteMapRulesResponseBody
	GetRouteMapRules() []*ListRouteMapRulesResponseBodyRouteMapRules
	SetTotalCount(v int64) *ListRouteMapRulesResponseBody
	GetTotalCount() *int64
}

type ListRouteMapRulesResponseBody struct {
	// The number of entries returned on each page in a paginated query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query. Valid values:
	//
	// - If **NextToken*	- is empty, no next query is available.
	//
	// - If **NextToken*	- has a return value, the value indicates the token for the next query.
	//
	// example:
	//
	// dd20****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EF44D7B7-6248-5A60-8A4C-A92388605A57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of route map rules.
	RouteMapRules []*ListRouteMapRulesResponseBodyRouteMapRules `json:"RouteMapRules,omitempty" xml:"RouteMapRules,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRouteMapRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRouteMapRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRouteMapRulesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRouteMapRulesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRouteMapRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRouteMapRulesResponseBody) GetRouteMapRules() []*ListRouteMapRulesResponseBodyRouteMapRules {
	return s.RouteMapRules
}

func (s *ListRouteMapRulesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListRouteMapRulesResponseBody) SetMaxResults(v int32) *ListRouteMapRulesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListRouteMapRulesResponseBody) SetNextToken(v string) *ListRouteMapRulesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRouteMapRulesResponseBody) SetRequestId(v string) *ListRouteMapRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRouteMapRulesResponseBody) SetRouteMapRules(v []*ListRouteMapRulesResponseBodyRouteMapRules) *ListRouteMapRulesResponseBody {
	s.RouteMapRules = v
	return s
}

func (s *ListRouteMapRulesResponseBody) SetTotalCount(v int64) *ListRouteMapRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRouteMapRulesResponseBody) Validate() error {
	if s.RouteMapRules != nil {
		for _, item := range s.RouteMapRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRouteMapRulesResponseBodyRouteMapRules struct {
	// The description of the rule.
	//
	// example:
	//
	// Matches production prefixes from on-prem and sets a higher local preference.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The match conditions of the route map rule. The system executes the actions defined in **SetActions*	- only when a route meets all the match conditions defined in this parameter.
	MatchConditions *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions `json:"MatchConditions,omitempty" xml:"MatchConditions,omitempty" type:"Struct"`
	// The name of the rule.
	//
	// example:
	//
	// Permit-Prod-Prefixes-Set-LP
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The priority of the rule. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the route map.
	//
	// example:
	//
	// rm-0a1b2c3d4e5f6a7b8
	RouteMapId *string `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
	// The ID of the route map rule.
	//
	// example:
	//
	// rmr-3dm13f66wtctcnydup
	RouteMapRuleId *string `json:"RouteMapRuleId,omitempty" xml:"RouteMapRuleId,omitempty"`
	// The actions to execute for the route map rule. After a route meets all the match conditions defined in **MatchConditions**, the system processes the route based on the actions defined in this parameter.
	SetActions *ListRouteMapRulesResponseBodyRouteMapRulesSetActions `json:"SetActions,omitempty" xml:"SetActions,omitempty" type:"Struct"`
	// The status of the route map rule. Valid values:
	//
	// - **Creating**: The rule is being created.
	//
	// - **Updating**: The rule is being updated.
	//
	// - **Active**: The rule is available.
	//
	// - **Deleting**: The rule is being deleted.
	//
	// - **Deleted**: The rule is deleted.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListRouteMapRulesResponseBodyRouteMapRules) String() string {
	return dara.Prettify(s)
}

func (s ListRouteMapRulesResponseBodyRouteMapRules) GoString() string {
	return s.String()
}

func (s *ListRouteMapRulesResponseBodyRouteMapRules) GetDescription() *string {
	return s.Description
}

func (s *ListRouteMapRulesResponseBodyRouteMapRules) GetMatchConditions() *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions {
	return s.MatchConditions
}

func (s *ListRouteMapRulesResponseBodyRouteMapRules) GetName() *string {
	return s.Name
}

func (s *ListRouteMapRulesResponseBodyRouteMapRules) GetPriority() *int32 {
	return s.Priority
}

func (s *ListRouteMapRulesResponseBodyRouteMapRules) GetRouteMapId() *string {
	return s.RouteMapId
}

func (s *ListRouteMapRulesResponseBodyRouteMapRules) GetRouteMapRuleId() *string {
	return s.RouteMapRuleId
}

func (s *ListRouteMapRulesResponseBodyRouteMapRules) GetSetActions() *ListRouteMapRulesResponseBodyRouteMapRulesSetActions {
	return s.SetActions
}

func (s *ListRouteMapRulesResponseBodyRouteMapRules) GetStatus() *string {
	return s.Status
}

func (s *ListRouteMapRulesResponseBodyRouteMapRules) SetDescription(v string) *ListRouteMapRulesResponseBodyRouteMapRules {
	s.Description = &v
	return s
}

func (s *ListRouteMapRulesResponseBodyRouteMapRules) SetMatchConditions(v *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) *ListRouteMapRulesResponseBodyRouteMapRules {
	s.MatchConditions = v
	return s
}

func (s *ListRouteMapRulesResponseBodyRouteMapRules) SetName(v string) *ListRouteMapRulesResponseBodyRouteMapRules {
	s.Name = &v
	return s
}

func (s *ListRouteMapRulesResponseBodyRouteMapRules) SetPriority(v int32) *ListRouteMapRulesResponseBodyRouteMapRules {
	s.Priority = &v
	return s
}

func (s *ListRouteMapRulesResponseBodyRouteMapRules) SetRouteMapId(v string) *ListRouteMapRulesResponseBodyRouteMapRules {
	s.RouteMapId = &v
	return s
}

func (s *ListRouteMapRulesResponseBodyRouteMapRules) SetRouteMapRuleId(v string) *ListRouteMapRulesResponseBodyRouteMapRules {
	s.RouteMapRuleId = &v
	return s
}

func (s *ListRouteMapRulesResponseBodyRouteMapRules) SetSetActions(v *ListRouteMapRulesResponseBodyRouteMapRulesSetActions) *ListRouteMapRulesResponseBodyRouteMapRules {
	s.SetActions = v
	return s
}

func (s *ListRouteMapRulesResponseBodyRouteMapRules) SetStatus(v string) *ListRouteMapRulesResponseBodyRouteMapRules {
	s.Status = &v
	return s
}

func (s *ListRouteMapRulesResponseBodyRouteMapRules) Validate() error {
	if s.MatchConditions != nil {
		if err := s.MatchConditions.Validate(); err != nil {
			return err
		}
	}
	if s.SetActions != nil {
		if err := s.SetActions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions struct {
	// The route prefix list for fuzzy matching.
	//
	// A match succeeds if the route prefix in the match condition contains the route prefix of the matched route.
	//
	// For example, a policy defined with 10.0.0.0/8 can fuzzy match a route with 10.10.10.0/24.
	//
	// IP address ranges use the CIDR format. Both IPv4 and IPv6 are supported.
	//
	// This parameter and **AddressPrefixesMatch*	- are mutually exclusive and cannot be specified at the same time.
	AddressPrefixesInclude []*string `json:"AddressPrefixesInclude,omitempty" xml:"AddressPrefixesInclude,omitempty" type:"Repeated"`
	// The route prefix list for exact matching.
	//
	// A match succeeds only when the route prefix in the match condition is exactly the same as the route prefix of the matched route.
	//
	// For example, a policy defined with 10.0.0.0/8 can only exactly match a route with 10.0.0.0/8.
	//
	// IP address ranges use the CIDR format. Both IPv4 and IPv6 are supported.
	//
	// This parameter and **AddressPrefixesInclude*	- are mutually exclusive and cannot be specified at the same time.
	AddressPrefixesMatch []*string `json:"AddressPrefixesMatch,omitempty" xml:"AddressPrefixesMatch,omitempty" type:"Repeated"`
	// The list of AS paths for fuzzy matching.
	//
	// A match succeeds if the AS path in the match condition overlaps with the AS path of the matched route.
	//
	// This parameter and **AsPathsMatch*	- are mutually exclusive and cannot be specified at the same time.
	//
	// > Only AS SEQUENCE is supported. AS SET, AS CONFED SEQUENCE, and AS CONFED SET are not supported. This means only a list of AS numbers is supported, and sets and sub-lists are not supported.
	AsPathsInclude []*int64 `json:"AsPathsInclude,omitempty" xml:"AsPathsInclude,omitempty" type:"Repeated"`
	// The list of AS paths for exact matching.
	//
	// A match succeeds only when the AS path in the match condition is exactly the same as the AS path of the matched route.
	//
	// This parameter and **AsPathsInclude*	- are mutually exclusive and cannot be specified at the same time.
	//
	// > Only AS SEQUENCE is supported. AS SET, AS CONFED SEQUENCE, and AS CONFED SET are not supported. This means only a list of AS numbers is supported, and sets and sub-lists are not supported.
	AsPathsMatch []*int64 `json:"AsPathsMatch,omitempty" xml:"AsPathsMatch,omitempty" type:"Repeated"`
	// The community set for fuzzy matching.
	//
	// A match succeeds if the community in the match condition overlaps with the community of the matched route.
	//
	// Each community is in the format of n:m, where n and m range from **1*	- to **65535**. The community must comply with RFC 1997. Large communities (RFC 8092) are not supported.
	//
	// This parameter and **CommunitySetMatch*	- are mutually exclusive and cannot be specified at the same time.
	//
	// > Incorrect community configurations may prevent routes from being advertised to the on-premises data center.
	CommunitySetInclude []*string `json:"CommunitySetInclude,omitempty" xml:"CommunitySetInclude,omitempty" type:"Repeated"`
	// The community set for exact matching.
	//
	// A match succeeds only when the community in the match condition is exactly the same as the community of the matched route.
	//
	// Each community is in the format of n:m, where n and m range from **1*	- to **65535**. The community must comply with RFC 1997. Large communities (RFC 8092) are not supported.
	//
	// This parameter and **CommunitySetInclude*	- are mutually exclusive and cannot be specified at the same time.
	//
	// > Incorrect community configurations may prevent routes from being advertised to the on-premises data center.
	CommunitySetMatch []*string `json:"CommunitySetMatch,omitempty" xml:"CommunitySetMatch,omitempty" type:"Repeated"`
	// The list of destination instance IDs to match. The following instance types are supported:
	//
	// - Virtual private cloud (VPC) instance IDs
	//
	// - Virtual border router (VBR) instance IDs
	//
	// > The list of destination instance IDs takes effect only when the direction of the route map is outbound.
	DestinationInstanceIds []*string `json:"DestinationInstanceIds,omitempty" xml:"DestinationInstanceIds,omitempty" type:"Repeated"`
	// Specifies whether to use the exclude matching mode for the list of destination instance IDs. Valid values:
	//
	// - **false*	- (default): no. A match succeeds when the destination instance ID of the route is in **DestinationInstanceIds.N**.
	//
	// - **true**: yes. A match succeeds when the destination instance ID of the route is not in **DestinationInstanceIds.N**.
	//
	// example:
	//
	// false
	DestinationInstanceIdsReverseMatch *bool `json:"DestinationInstanceIdsReverseMatch,omitempty" xml:"DestinationInstanceIdsReverseMatch,omitempty"`
	// The list of destination instance types to match. The following instance types are supported:
	//
	// - **VPC**: VPC instance.
	//
	// - **VBR**: VBR instance.
	//
	// > The list of destination instance types takes effect only when the direction of the route map is outbound.
	DestinationInstanceTypes []*string `json:"DestinationInstanceTypes,omitempty" xml:"DestinationInstanceTypes,omitempty" type:"Repeated"`
	// The list of destination region IDs to match.
	DestinationRegionIds []*string `json:"DestinationRegionIds,omitempty" xml:"DestinationRegionIds,omitempty" type:"Repeated"`
	// The list of destination route table IDs to match.
	//
	// > The list of destination route table IDs takes effect only when the direction of the route map is outbound.
	DestinationRouteTableIds []*string `json:"DestinationRouteTableIds,omitempty" xml:"DestinationRouteTableIds,omitempty" type:"Repeated"`
	// The IP address type to match. Valid values:
	//
	// - **IPv4**: matches only IPv4 routes.
	//
	// - **IPv6**: matches only IPv6 routes.
	//
	// You can leave this parameter empty to match routes of all types.
	//
	// example:
	//
	// IPv4
	MatchAddressType *string `json:"MatchAddressType,omitempty" xml:"MatchAddressType,omitempty"`
	// The list of route types to match. The following route types are supported:
	//
	// - **System**: system route, which is automatically generated by the system.
	//
	// - **Custom**: custom route, which is manually added by a user.
	//
	// - **BGP**: BGP route, which is propagated through the BGP routable protocol.
	//
	// You can specify multiple route types.
	RouteTypes []*string `json:"RouteTypes,omitempty" xml:"RouteTypes,omitempty" type:"Repeated"`
	// The list of source instance IDs to match. The following instance types are supported:
	//
	// - Virtual private cloud (VPC) instance IDs
	//
	// - Virtual border router (VBR) instance IDs
	SourceInstanceIds []*string `json:"SourceInstanceIds,omitempty" xml:"SourceInstanceIds,omitempty" type:"Repeated"`
	// Specifies whether to use the exclude matching mode for the list of source instance IDs. Valid values:
	//
	// - **false*	- (default): no. A match succeeds when the source instance ID of the route is in **SourceInstanceIds.N**.
	//
	// - **true**: yes. A match succeeds when the source instance ID of the route is not in **SourceInstanceIds.N**.
	//
	// example:
	//
	// false
	SourceInstanceIdsReverseMatch *bool `json:"SourceInstanceIdsReverseMatch,omitempty" xml:"SourceInstanceIdsReverseMatch,omitempty"`
	// The list of source instance types to match. The following instance types are supported:
	//
	// - **VPC**: VPC instance.
	//
	// - **VBR**: VBR instance.
	SourceInstanceTypes []*string `json:"SourceInstanceTypes,omitempty" xml:"SourceInstanceTypes,omitempty" type:"Repeated"`
	// The list of source region IDs to match.
	SourceRegionIds []*string `json:"SourceRegionIds,omitempty" xml:"SourceRegionIds,omitempty" type:"Repeated"`
	// The list of source route table IDs to match.
	SourceRouteTableIds []*string `json:"SourceRouteTableIds,omitempty" xml:"SourceRouteTableIds,omitempty" type:"Repeated"`
}

func (s ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) String() string {
	return dara.Prettify(s)
}

func (s ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) GoString() string {
	return s.String()
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) GetAddressPrefixesInclude() []*string {
	return s.AddressPrefixesInclude
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) GetAddressPrefixesMatch() []*string {
	return s.AddressPrefixesMatch
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) GetAsPathsInclude() []*int64 {
	return s.AsPathsInclude
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) GetAsPathsMatch() []*int64 {
	return s.AsPathsMatch
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) GetCommunitySetInclude() []*string {
	return s.CommunitySetInclude
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) GetCommunitySetMatch() []*string {
	return s.CommunitySetMatch
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) GetDestinationInstanceIds() []*string {
	return s.DestinationInstanceIds
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) GetDestinationInstanceIdsReverseMatch() *bool {
	return s.DestinationInstanceIdsReverseMatch
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) GetDestinationInstanceTypes() []*string {
	return s.DestinationInstanceTypes
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) GetDestinationRegionIds() []*string {
	return s.DestinationRegionIds
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) GetDestinationRouteTableIds() []*string {
	return s.DestinationRouteTableIds
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) GetMatchAddressType() *string {
	return s.MatchAddressType
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) GetRouteTypes() []*string {
	return s.RouteTypes
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) GetSourceInstanceIds() []*string {
	return s.SourceInstanceIds
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) GetSourceInstanceIdsReverseMatch() *bool {
	return s.SourceInstanceIdsReverseMatch
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) GetSourceInstanceTypes() []*string {
	return s.SourceInstanceTypes
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) GetSourceRegionIds() []*string {
	return s.SourceRegionIds
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) GetSourceRouteTableIds() []*string {
	return s.SourceRouteTableIds
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) SetAddressPrefixesInclude(v []*string) *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions {
	s.AddressPrefixesInclude = v
	return s
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) SetAddressPrefixesMatch(v []*string) *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions {
	s.AddressPrefixesMatch = v
	return s
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) SetAsPathsInclude(v []*int64) *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions {
	s.AsPathsInclude = v
	return s
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) SetAsPathsMatch(v []*int64) *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions {
	s.AsPathsMatch = v
	return s
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) SetCommunitySetInclude(v []*string) *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions {
	s.CommunitySetInclude = v
	return s
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) SetCommunitySetMatch(v []*string) *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions {
	s.CommunitySetMatch = v
	return s
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) SetDestinationInstanceIds(v []*string) *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions {
	s.DestinationInstanceIds = v
	return s
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) SetDestinationInstanceIdsReverseMatch(v bool) *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions {
	s.DestinationInstanceIdsReverseMatch = &v
	return s
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) SetDestinationInstanceTypes(v []*string) *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions {
	s.DestinationInstanceTypes = v
	return s
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) SetDestinationRegionIds(v []*string) *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions {
	s.DestinationRegionIds = v
	return s
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) SetDestinationRouteTableIds(v []*string) *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions {
	s.DestinationRouteTableIds = v
	return s
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) SetMatchAddressType(v string) *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions {
	s.MatchAddressType = &v
	return s
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) SetRouteTypes(v []*string) *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions {
	s.RouteTypes = v
	return s
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) SetSourceInstanceIds(v []*string) *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions {
	s.SourceInstanceIds = v
	return s
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) SetSourceInstanceIdsReverseMatch(v bool) *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions {
	s.SourceInstanceIdsReverseMatch = &v
	return s
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) SetSourceInstanceTypes(v []*string) *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions {
	s.SourceInstanceTypes = v
	return s
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) SetSourceRegionIds(v []*string) *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions {
	s.SourceRegionIds = v
	return s
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) SetSourceRouteTableIds(v []*string) *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions {
	s.SourceRouteTableIds = v
	return s
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesMatchConditions) Validate() error {
	return dara.Validate(s)
}

type ListRouteMapRulesResponseBodyRouteMapRulesSetActions struct {
	// The AS path to prepend when the regional gateway receives or advertises routes.
	//
	// This parameter and **AsPathReplace*	- are mutually exclusive and cannot be specified at the same time.
	//
	// This parameter specifies the action to execute after a route meets the match conditions.
	AsPathPrepend []*int64 `json:"AsPathPrepend,omitempty" xml:"AsPathPrepend,omitempty" type:"Repeated"`
	// The AS path to replace when the regional gateway receives or advertises routes. The original AS path of the route is replaced with the list of AS numbers specified in this parameter.
	//
	// This parameter and **AsPathPrepend*	- are mutually exclusive and cannot be specified at the same time.
	//
	// This parameter specifies the action to execute after a route meets the match conditions.
	AsPathReplace []*int64 `json:"AsPathReplace,omitempty" xml:"AsPathReplace,omitempty" type:"Repeated"`
	// The community set to add to the route.
	//
	// Each community is in the format of n:m, where n and m range from **1*	- to **65535**. The community must comply with RFC 1997. Large communities (RFC 8092) are not supported.
	//
	// This parameter and **CommunityReplace*	- are mutually exclusive and cannot be specified at the same time.
	//
	// This parameter specifies the action to execute after a route meets the match conditions.
	//
	// > Incorrect community configurations may prevent routes from being advertised to the on-premises data center.
	CommunityAdd []*string `json:"CommunityAdd,omitempty" xml:"CommunityAdd,omitempty" type:"Repeated"`
	// The community set to replace the original community of the route.
	//
	// Each community is in the format of n:m, where n and m range from **1*	- to **65535**. The community must comply with RFC 1997. Large communities (RFC 8092) are not supported.
	//
	// This parameter and **CommunityAdd*	- are mutually exclusive and cannot be specified at the same time.
	//
	// This parameter specifies the action to execute after a route meets the match conditions.
	//
	// > Incorrect community configurations may prevent routes from being advertised to the on-premises data center.
	CommunityReplace []*string `json:"CommunityReplace,omitempty" xml:"CommunityReplace,omitempty" type:"Repeated"`
	// Policy priority of the associated next route map rule.
	//
	// - You can set policy priority of the associated route map rule only when **RouteAction*	- is set to **Permit**. Only the permitted routes continue to match the associated next route map rule.
	//
	// - Policy priority of the associated next route map rule must be lower than policy priority of the current route map rule.
	//
	// example:
	//
	// 1
	NextPriority *int64 `json:"NextPriority,omitempty" xml:"NextPriority,omitempty"`
	// The policy action after all conditions are met. Valid values:
	//
	// - **Permit**: permits the matched routes.
	//
	// - **Deny**: denies the matched routes.
	//
	// example:
	//
	// Permit
	RouteAction *string `json:"RouteAction,omitempty" xml:"RouteAction,omitempty"`
}

func (s ListRouteMapRulesResponseBodyRouteMapRulesSetActions) String() string {
	return dara.Prettify(s)
}

func (s ListRouteMapRulesResponseBodyRouteMapRulesSetActions) GoString() string {
	return s.String()
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesSetActions) GetAsPathPrepend() []*int64 {
	return s.AsPathPrepend
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesSetActions) GetAsPathReplace() []*int64 {
	return s.AsPathReplace
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesSetActions) GetCommunityAdd() []*string {
	return s.CommunityAdd
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesSetActions) GetCommunityReplace() []*string {
	return s.CommunityReplace
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesSetActions) GetNextPriority() *int64 {
	return s.NextPriority
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesSetActions) GetRouteAction() *string {
	return s.RouteAction
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesSetActions) SetAsPathPrepend(v []*int64) *ListRouteMapRulesResponseBodyRouteMapRulesSetActions {
	s.AsPathPrepend = v
	return s
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesSetActions) SetAsPathReplace(v []*int64) *ListRouteMapRulesResponseBodyRouteMapRulesSetActions {
	s.AsPathReplace = v
	return s
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesSetActions) SetCommunityAdd(v []*string) *ListRouteMapRulesResponseBodyRouteMapRulesSetActions {
	s.CommunityAdd = v
	return s
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesSetActions) SetCommunityReplace(v []*string) *ListRouteMapRulesResponseBodyRouteMapRulesSetActions {
	s.CommunityReplace = v
	return s
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesSetActions) SetNextPriority(v int64) *ListRouteMapRulesResponseBodyRouteMapRulesSetActions {
	s.NextPriority = &v
	return s
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesSetActions) SetRouteAction(v string) *ListRouteMapRulesResponseBodyRouteMapRulesSetActions {
	s.RouteAction = &v
	return s
}

func (s *ListRouteMapRulesResponseBodyRouteMapRulesSetActions) Validate() error {
	return dara.Validate(s)
}
