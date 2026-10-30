// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRouteMapRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateRouteMapRuleRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateRouteMapRuleRequest
	GetDescription() *string
	SetDryRun(v bool) *UpdateRouteMapRuleRequest
	GetDryRun() *bool
	SetMatchConditions(v *UpdateRouteMapRuleRequestMatchConditions) *UpdateRouteMapRuleRequest
	GetMatchConditions() *UpdateRouteMapRuleRequestMatchConditions
	SetName(v string) *UpdateRouteMapRuleRequest
	GetName() *string
	SetPriority(v int32) *UpdateRouteMapRuleRequest
	GetPriority() *int32
	SetRouteMapRuleId(v string) *UpdateRouteMapRuleRequest
	GetRouteMapRuleId() *string
	SetSetActions(v *UpdateRouteMapRuleRequestSetActions) *UpdateRouteMapRuleRequest
	GetSetActions() *UpdateRouteMapRuleRequestSetActions
}

type UpdateRouteMapRuleRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that the value is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The detailed description of the rule. The description can be empty or 1 to 256 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// Matches production prefixes from on-prem and sets a higher local preference.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to precheck only this request. Valid values:
	//
	// - **true**: prechecks the request without updating the route map rule. The system checks whether required parameters are set, the request format, and service limits. If the request fails the precheck, the corresponding error is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends the request. If the request passes the precheck, the route map rule is updated.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The match conditions of the route map rule. The system performs the actions defined in **SetActions*	- only when a route meets all the match conditions defined in this parameter.
	//
	// if can be null:
	// false
	MatchConditions *UpdateRouteMapRuleRequestMatchConditions `json:"MatchConditions,omitempty" xml:"MatchConditions,omitempty" type:"Struct"`
	// The name of the rule for easy identification. The name can be empty or 1 to 128 characters in length, and cannot start with http:// or https://.
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
	// The ID of the route map rule to be updated.
	//
	// example:
	//
	// rmr-0a1b2c3d4e5f6a7b8
	RouteMapRuleId *string `json:"RouteMapRuleId,omitempty" xml:"RouteMapRuleId,omitempty"`
	// The actions to be performed for the route map rule. After a route meets all the match conditions defined in **MatchConditions**, the system processes the route based on the actions defined in this parameter.
	//
	// if can be null:
	// false
	SetActions *UpdateRouteMapRuleRequestSetActions `json:"SetActions,omitempty" xml:"SetActions,omitempty" type:"Struct"`
}

func (s UpdateRouteMapRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRouteMapRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateRouteMapRuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateRouteMapRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateRouteMapRuleRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateRouteMapRuleRequest) GetMatchConditions() *UpdateRouteMapRuleRequestMatchConditions {
	return s.MatchConditions
}

func (s *UpdateRouteMapRuleRequest) GetName() *string {
	return s.Name
}

func (s *UpdateRouteMapRuleRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateRouteMapRuleRequest) GetRouteMapRuleId() *string {
	return s.RouteMapRuleId
}

func (s *UpdateRouteMapRuleRequest) GetSetActions() *UpdateRouteMapRuleRequestSetActions {
	return s.SetActions
}

func (s *UpdateRouteMapRuleRequest) SetClientToken(v string) *UpdateRouteMapRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateRouteMapRuleRequest) SetDescription(v string) *UpdateRouteMapRuleRequest {
	s.Description = &v
	return s
}

func (s *UpdateRouteMapRuleRequest) SetDryRun(v bool) *UpdateRouteMapRuleRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateRouteMapRuleRequest) SetMatchConditions(v *UpdateRouteMapRuleRequestMatchConditions) *UpdateRouteMapRuleRequest {
	s.MatchConditions = v
	return s
}

func (s *UpdateRouteMapRuleRequest) SetName(v string) *UpdateRouteMapRuleRequest {
	s.Name = &v
	return s
}

func (s *UpdateRouteMapRuleRequest) SetPriority(v int32) *UpdateRouteMapRuleRequest {
	s.Priority = &v
	return s
}

func (s *UpdateRouteMapRuleRequest) SetRouteMapRuleId(v string) *UpdateRouteMapRuleRequest {
	s.RouteMapRuleId = &v
	return s
}

func (s *UpdateRouteMapRuleRequest) SetSetActions(v *UpdateRouteMapRuleRequestSetActions) *UpdateRouteMapRuleRequest {
	s.SetActions = v
	return s
}

func (s *UpdateRouteMapRuleRequest) Validate() error {
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

type UpdateRouteMapRuleRequestMatchConditions struct {
	// The route prefix list for fuzzy matching. You can specify up to 64 CIDR blocks.
	//
	// A match is successful if the route prefix in the match condition contains the route prefix of the route being matched.
	//
	// For example, a policy defined with 10.0.0.0/8 can fuzzy match the 10.10.10.0/24 route.
	//
	// The CIDR blocks use the CIDR format and support both IPv4 and IPv6.
	//
	// This parameter is mutually exclusive with **AddressPrefixesMatch**. You cannot specify both parameters at the same time.
	AddressPrefixesInclude []*string `json:"AddressPrefixesInclude,omitempty" xml:"AddressPrefixesInclude,omitempty" type:"Repeated"`
	// The route prefix list for exact matching. You can specify up to 64 CIDR blocks.
	//
	// A match is successful only if the route prefix in the match condition is identical to the route prefix of the route being matched.
	//
	// For example, a policy defined with 10.0.0.0/8 can exactly match only the 10.0.0.0/8 route.
	//
	// The CIDR blocks use the CIDR format and support both IPv4 and IPv6.
	//
	// This parameter is mutually exclusive with **AddressPrefixesInclude**. You cannot specify both parameters at the same time.
	AddressPrefixesMatch []*string `json:"AddressPrefixesMatch,omitempty" xml:"AddressPrefixesMatch,omitempty" type:"Repeated"`
	// The list of AS paths for fuzzy match. You can enter up to 64 AS numbers.
	//
	// A route is considered a match if the AS path in the match condition overlaps with the AS path of the matched route.
	//
	// This parameter is mutually exclusive with **AsPathsMatch*	- and cannot be specified at the same time.
	//
	// > Only AS SEQUENCE is supported. AS SET, AS CONFED SEQUENCE, and AS CONFED SET are not supported. This means that only a list of AS numbers is supported, and sets and sublists are not supported.
	AsPathsInclude []*int64 `json:"AsPathsInclude,omitempty" xml:"AsPathsInclude,omitempty" type:"Repeated"`
	// The list of AS paths for exact match. You can enter up to 64 AS numbers.
	//
	// A route is considered a match only when the AS path in the match condition is exactly the same as the AS path of the matched route.
	//
	// This parameter is mutually exclusive with **AsPathsInclude*	- and cannot be specified at the same time.
	//
	// > Only AS SEQUENCE is supported. AS SET, AS CONFED SEQUENCE, and AS CONFED SET are not supported. This means that only a list of AS numbers is supported, and sets and sublists are not supported.
	AsPathsMatch []*int64 `json:"AsPathsMatch,omitempty" xml:"AsPathsMatch,omitempty" type:"Repeated"`
	// The community set for fuzzy match. You can enter up to 64 communities.
	//
	// A route is considered a match if the community in the match condition overlaps with the community of the matched route.
	//
	// Each community is in the format of n:m, where n and m range from **1*	- to **65535**. The community must comply with RFC 1997. Large communities (RFC 8092) are not supported.
	//
	// This parameter is mutually exclusive with **CommunitySetMatch*	- and cannot be specified at the same time.
	//
	// > Incorrect community configurations may prevent routes from being advertised to the on-premises data center.
	CommunitySetInclude []*string `json:"CommunitySetInclude,omitempty" xml:"CommunitySetInclude,omitempty" type:"Repeated"`
	// The community set for exact match. You can enter up to 64 communities.
	//
	// A route is considered a match only when the community in the match condition is exactly the same as the community of the matched route.
	//
	// Each community is in the format of n:m, where n and m range from **1*	- to **65535**. The community must comply with RFC 1997. Large communities (RFC 8092) are not supported.
	//
	// This parameter is mutually exclusive with **CommunitySetInclude*	- and cannot be specified at the same time.
	//
	// > Incorrect community configurations may prevent routes from being advertised to the on-premises data center.
	CommunitySetMatch []*string `json:"CommunitySetMatch,omitempty" xml:"CommunitySetMatch,omitempty" type:"Repeated"`
	// The list of destination instance IDs to be matched. You can enter the following types of instance IDs:
	//
	// - Virtual private cloud (VPC) instance IDs
	//
	// - Virtual border router (VBR) instance IDs
	//
	// You can enter up to 64 instance IDs.
	//
	// > The list of destination instance IDs takes effect only when the direction of the route map is outbound.
	DestinationInstanceIds []*string `json:"DestinationInstanceIds,omitempty" xml:"DestinationInstanceIds,omitempty" type:"Repeated"`
	// Specifies whether to use the exclusion match mode for the list of destination instance IDs. Valid values:
	//
	// - **false*	- (default): no. A route is considered a match if the destination instance ID of the route is in **DestinationInstanceIds.N**.
	//
	// - **true**: yes. A route is considered a match if the destination instance ID of the route is not in **DestinationInstanceIds.N**.
	//
	// example:
	//
	// false
	DestinationInstanceIdsReverseMatch *bool `json:"DestinationInstanceIdsReverseMatch,omitempty" xml:"DestinationInstanceIdsReverseMatch,omitempty"`
	// The list of destination instance types to be matched. The following instance types are supported:
	//
	// - **VPC**: VPC instance.
	//
	// - **VBR**: VBR instance.
	//
	//
	// You can enter multiple instance types.
	//
	// > The list of destination instance types takes effect only when the direction of the route map is outbound.
	DestinationInstanceTypes []*string `json:"DestinationInstanceTypes,omitempty" xml:"DestinationInstanceTypes,omitempty" type:"Repeated"`
	// The list of destination region IDs to be matched. You can enter up to 64 region IDs.
	DestinationRegionIds []*string `json:"DestinationRegionIds,omitempty" xml:"DestinationRegionIds,omitempty" type:"Repeated"`
	// The list of destination route table IDs to be matched. You can enter up to 64 route table IDs.
	//
	// > The list of destination route table IDs takes effect only when the direction of the route map is outbound.
	DestinationRouteTableIds []*string `json:"DestinationRouteTableIds,omitempty" xml:"DestinationRouteTableIds,omitempty" type:"Repeated"`
	// The IP address type of the routes to be matched. Valid values:
	//
	// - **IPv4**: matches only IPv4 routes.
	//
	// - **IPv6**: matches only IPv6 routes.
	//
	// This parameter can be left empty, which indicates that all types of routes are matched.
	//
	// example:
	//
	// IPv4
	MatchAddressType *string `json:"MatchAddressType,omitempty" xml:"MatchAddressType,omitempty"`
	// The list of route types to match. The following route types are supported:
	//
	// - **System**: system route, which is automatically generated by the system.
	//
	// - **Custom**: custom route, which is manually added by users.
	//
	// - **BGP**: BGP route, which is propagated through the BGP routable protocol.
	//
	// You can specify multiple route types.
	RouteTypes []*string `json:"RouteTypes,omitempty" xml:"RouteTypes,omitempty" type:"Repeated"`
	// The list of source instance IDs to be matched. You can enter the following types of instance IDs:
	//
	// - Virtual private cloud (VPC) instance IDs
	//
	// - Virtual border router (VBR) instance IDs
	//
	//
	// You can enter up to 64 instance IDs.
	SourceInstanceIds []*string `json:"SourceInstanceIds,omitempty" xml:"SourceInstanceIds,omitempty" type:"Repeated"`
	// Specifies whether to use the exclusion match mode for the list of source instance IDs. Valid values:
	//
	// - **false*	- (default): no. A route is considered a match if the source instance ID of the route is in **SourceInstanceIds.N**.
	//
	// - **true**: yes. A route is considered a match if the source instance ID of the route is not in **SourceInstanceIds.N**.
	//
	// example:
	//
	// false
	SourceInstanceIdsReverseMatch *bool `json:"SourceInstanceIdsReverseMatch,omitempty" xml:"SourceInstanceIdsReverseMatch,omitempty"`
	// The list of source instance types to be matched. The following instance types are supported:
	//
	// - **VPC**: VPC instance.
	//
	// - **VBR**: VBR instance.
	//
	//
	// You can enter multiple instance types.
	SourceInstanceTypes []*string `json:"SourceInstanceTypes,omitempty" xml:"SourceInstanceTypes,omitempty" type:"Repeated"`
	// The list of source region IDs to be matched. You can enter up to 64 region IDs.
	SourceRegionIds []*string `json:"SourceRegionIds,omitempty" xml:"SourceRegionIds,omitempty" type:"Repeated"`
	// The list of source route table IDs to be matched. You can enter up to 64 route table IDs.
	SourceRouteTableIds []*string `json:"SourceRouteTableIds,omitempty" xml:"SourceRouteTableIds,omitempty" type:"Repeated"`
}

func (s UpdateRouteMapRuleRequestMatchConditions) String() string {
	return dara.Prettify(s)
}

func (s UpdateRouteMapRuleRequestMatchConditions) GoString() string {
	return s.String()
}

func (s *UpdateRouteMapRuleRequestMatchConditions) GetAddressPrefixesInclude() []*string {
	return s.AddressPrefixesInclude
}

func (s *UpdateRouteMapRuleRequestMatchConditions) GetAddressPrefixesMatch() []*string {
	return s.AddressPrefixesMatch
}

func (s *UpdateRouteMapRuleRequestMatchConditions) GetAsPathsInclude() []*int64 {
	return s.AsPathsInclude
}

func (s *UpdateRouteMapRuleRequestMatchConditions) GetAsPathsMatch() []*int64 {
	return s.AsPathsMatch
}

func (s *UpdateRouteMapRuleRequestMatchConditions) GetCommunitySetInclude() []*string {
	return s.CommunitySetInclude
}

func (s *UpdateRouteMapRuleRequestMatchConditions) GetCommunitySetMatch() []*string {
	return s.CommunitySetMatch
}

func (s *UpdateRouteMapRuleRequestMatchConditions) GetDestinationInstanceIds() []*string {
	return s.DestinationInstanceIds
}

func (s *UpdateRouteMapRuleRequestMatchConditions) GetDestinationInstanceIdsReverseMatch() *bool {
	return s.DestinationInstanceIdsReverseMatch
}

func (s *UpdateRouteMapRuleRequestMatchConditions) GetDestinationInstanceTypes() []*string {
	return s.DestinationInstanceTypes
}

func (s *UpdateRouteMapRuleRequestMatchConditions) GetDestinationRegionIds() []*string {
	return s.DestinationRegionIds
}

func (s *UpdateRouteMapRuleRequestMatchConditions) GetDestinationRouteTableIds() []*string {
	return s.DestinationRouteTableIds
}

func (s *UpdateRouteMapRuleRequestMatchConditions) GetMatchAddressType() *string {
	return s.MatchAddressType
}

func (s *UpdateRouteMapRuleRequestMatchConditions) GetRouteTypes() []*string {
	return s.RouteTypes
}

func (s *UpdateRouteMapRuleRequestMatchConditions) GetSourceInstanceIds() []*string {
	return s.SourceInstanceIds
}

func (s *UpdateRouteMapRuleRequestMatchConditions) GetSourceInstanceIdsReverseMatch() *bool {
	return s.SourceInstanceIdsReverseMatch
}

func (s *UpdateRouteMapRuleRequestMatchConditions) GetSourceInstanceTypes() []*string {
	return s.SourceInstanceTypes
}

func (s *UpdateRouteMapRuleRequestMatchConditions) GetSourceRegionIds() []*string {
	return s.SourceRegionIds
}

func (s *UpdateRouteMapRuleRequestMatchConditions) GetSourceRouteTableIds() []*string {
	return s.SourceRouteTableIds
}

func (s *UpdateRouteMapRuleRequestMatchConditions) SetAddressPrefixesInclude(v []*string) *UpdateRouteMapRuleRequestMatchConditions {
	s.AddressPrefixesInclude = v
	return s
}

func (s *UpdateRouteMapRuleRequestMatchConditions) SetAddressPrefixesMatch(v []*string) *UpdateRouteMapRuleRequestMatchConditions {
	s.AddressPrefixesMatch = v
	return s
}

func (s *UpdateRouteMapRuleRequestMatchConditions) SetAsPathsInclude(v []*int64) *UpdateRouteMapRuleRequestMatchConditions {
	s.AsPathsInclude = v
	return s
}

func (s *UpdateRouteMapRuleRequestMatchConditions) SetAsPathsMatch(v []*int64) *UpdateRouteMapRuleRequestMatchConditions {
	s.AsPathsMatch = v
	return s
}

func (s *UpdateRouteMapRuleRequestMatchConditions) SetCommunitySetInclude(v []*string) *UpdateRouteMapRuleRequestMatchConditions {
	s.CommunitySetInclude = v
	return s
}

func (s *UpdateRouteMapRuleRequestMatchConditions) SetCommunitySetMatch(v []*string) *UpdateRouteMapRuleRequestMatchConditions {
	s.CommunitySetMatch = v
	return s
}

func (s *UpdateRouteMapRuleRequestMatchConditions) SetDestinationInstanceIds(v []*string) *UpdateRouteMapRuleRequestMatchConditions {
	s.DestinationInstanceIds = v
	return s
}

func (s *UpdateRouteMapRuleRequestMatchConditions) SetDestinationInstanceIdsReverseMatch(v bool) *UpdateRouteMapRuleRequestMatchConditions {
	s.DestinationInstanceIdsReverseMatch = &v
	return s
}

func (s *UpdateRouteMapRuleRequestMatchConditions) SetDestinationInstanceTypes(v []*string) *UpdateRouteMapRuleRequestMatchConditions {
	s.DestinationInstanceTypes = v
	return s
}

func (s *UpdateRouteMapRuleRequestMatchConditions) SetDestinationRegionIds(v []*string) *UpdateRouteMapRuleRequestMatchConditions {
	s.DestinationRegionIds = v
	return s
}

func (s *UpdateRouteMapRuleRequestMatchConditions) SetDestinationRouteTableIds(v []*string) *UpdateRouteMapRuleRequestMatchConditions {
	s.DestinationRouteTableIds = v
	return s
}

func (s *UpdateRouteMapRuleRequestMatchConditions) SetMatchAddressType(v string) *UpdateRouteMapRuleRequestMatchConditions {
	s.MatchAddressType = &v
	return s
}

func (s *UpdateRouteMapRuleRequestMatchConditions) SetRouteTypes(v []*string) *UpdateRouteMapRuleRequestMatchConditions {
	s.RouteTypes = v
	return s
}

func (s *UpdateRouteMapRuleRequestMatchConditions) SetSourceInstanceIds(v []*string) *UpdateRouteMapRuleRequestMatchConditions {
	s.SourceInstanceIds = v
	return s
}

func (s *UpdateRouteMapRuleRequestMatchConditions) SetSourceInstanceIdsReverseMatch(v bool) *UpdateRouteMapRuleRequestMatchConditions {
	s.SourceInstanceIdsReverseMatch = &v
	return s
}

func (s *UpdateRouteMapRuleRequestMatchConditions) SetSourceInstanceTypes(v []*string) *UpdateRouteMapRuleRequestMatchConditions {
	s.SourceInstanceTypes = v
	return s
}

func (s *UpdateRouteMapRuleRequestMatchConditions) SetSourceRegionIds(v []*string) *UpdateRouteMapRuleRequestMatchConditions {
	s.SourceRegionIds = v
	return s
}

func (s *UpdateRouteMapRuleRequestMatchConditions) SetSourceRouteTableIds(v []*string) *UpdateRouteMapRuleRequestMatchConditions {
	s.SourceRouteTableIds = v
	return s
}

func (s *UpdateRouteMapRuleRequestMatchConditions) Validate() error {
	return dara.Validate(s)
}

type UpdateRouteMapRuleRequestSetActions struct {
	// The AS path to prepend when the regional gateway accepts or executes a publish route entry.
	//
	// This parameter is mutually exclusive with **AsPathReplace**. You cannot specify both parameters at the same time.
	//
	// This parameter specifies the action to execute after a route meets the match conditions.
	AsPathPrepend []*int64 `json:"AsPathPrepend,omitempty" xml:"AsPathPrepend,omitempty" type:"Repeated"`
	// The AS path to use as a replacement when the regional gateway accepts or executes a publish route entry. The original AS path of the route is replaced with the list of AS numbers specified by this parameter.
	//
	// This parameter is mutually exclusive with **AsPathPrepend**. You cannot specify both parameters at the same time.
	//
	// This parameter specifies the action to execute after a route meets the match conditions.
	AsPathReplace []*int64 `json:"AsPathReplace,omitempty" xml:"AsPathReplace,omitempty" type:"Repeated"`
	// The community set to be added to the route.
	//
	// Each community is in the format of n:m, where n and m range from **1*	- to **65535**. The community must comply with RFC 1997. Large communities (RFC 8092) are not supported.
	//
	// This parameter is mutually exclusive with **CommunityReplace*	- and cannot be specified at the same time.
	//
	// This parameter specifies the action to be performed after a route meets the match conditions.
	//
	// > Incorrect community configurations may prevent routes from being advertised to the on-premises data center.
	CommunityAdd []*string `json:"CommunityAdd,omitempty" xml:"CommunityAdd,omitempty" type:"Repeated"`
	// The community set used to replace the original community of the route.
	//
	// Each community is in the format of n:m, where n and m range from **1*	- to **65535**. The community must comply with RFC 1997. Large communities (RFC 8092) are not supported.
	//
	// This parameter is mutually exclusive with **CommunityAdd*	- and cannot be specified at the same time.
	//
	// This parameter specifies the action to be performed after a route meets the match conditions.
	//
	// > Incorrect community configurations may prevent routes from being advertised to the on-premises data center.
	CommunityReplace []*string `json:"CommunityReplace,omitempty" xml:"CommunityReplace,omitempty" type:"Repeated"`
	// Policy priority of the associated next route map rule.
	//
	// - You can set policy priority of the associated route map rule only when **RouteAction*	- is set to **Permit**. Only the permitted routes continue to match the associated next route map rule.
	//
	// - Policy priority of the associated next route map rule must be lower than that of the current route map rule.
	//
	// example:
	//
	// 1
	NextPriority *int64 `json:"NextPriority,omitempty" xml:"NextPriority,omitempty"`
	// The action to be performed after all conditions are met. Valid values:
	//
	// - **Permit**: permits the matched route.
	//
	// - **Deny**: denies the matched route.
	//
	// example:
	//
	// Permit
	RouteAction *string `json:"RouteAction,omitempty" xml:"RouteAction,omitempty"`
}

func (s UpdateRouteMapRuleRequestSetActions) String() string {
	return dara.Prettify(s)
}

func (s UpdateRouteMapRuleRequestSetActions) GoString() string {
	return s.String()
}

func (s *UpdateRouteMapRuleRequestSetActions) GetAsPathPrepend() []*int64 {
	return s.AsPathPrepend
}

func (s *UpdateRouteMapRuleRequestSetActions) GetAsPathReplace() []*int64 {
	return s.AsPathReplace
}

func (s *UpdateRouteMapRuleRequestSetActions) GetCommunityAdd() []*string {
	return s.CommunityAdd
}

func (s *UpdateRouteMapRuleRequestSetActions) GetCommunityReplace() []*string {
	return s.CommunityReplace
}

func (s *UpdateRouteMapRuleRequestSetActions) GetNextPriority() *int64 {
	return s.NextPriority
}

func (s *UpdateRouteMapRuleRequestSetActions) GetRouteAction() *string {
	return s.RouteAction
}

func (s *UpdateRouteMapRuleRequestSetActions) SetAsPathPrepend(v []*int64) *UpdateRouteMapRuleRequestSetActions {
	s.AsPathPrepend = v
	return s
}

func (s *UpdateRouteMapRuleRequestSetActions) SetAsPathReplace(v []*int64) *UpdateRouteMapRuleRequestSetActions {
	s.AsPathReplace = v
	return s
}

func (s *UpdateRouteMapRuleRequestSetActions) SetCommunityAdd(v []*string) *UpdateRouteMapRuleRequestSetActions {
	s.CommunityAdd = v
	return s
}

func (s *UpdateRouteMapRuleRequestSetActions) SetCommunityReplace(v []*string) *UpdateRouteMapRuleRequestSetActions {
	s.CommunityReplace = v
	return s
}

func (s *UpdateRouteMapRuleRequestSetActions) SetNextPriority(v int64) *UpdateRouteMapRuleRequestSetActions {
	s.NextPriority = &v
	return s
}

func (s *UpdateRouteMapRuleRequestSetActions) SetRouteAction(v string) *UpdateRouteMapRuleRequestSetActions {
	s.RouteAction = &v
	return s
}

func (s *UpdateRouteMapRuleRequestSetActions) Validate() error {
	return dara.Validate(s)
}
