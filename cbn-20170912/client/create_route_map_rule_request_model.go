// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRouteMapRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateRouteMapRuleRequest
	GetClientToken() *string
	SetDescription(v string) *CreateRouteMapRuleRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateRouteMapRuleRequest
	GetDryRun() *bool
	SetMatchConditions(v *CreateRouteMapRuleRequestMatchConditions) *CreateRouteMapRuleRequest
	GetMatchConditions() *CreateRouteMapRuleRequestMatchConditions
	SetName(v string) *CreateRouteMapRuleRequest
	GetName() *string
	SetPriority(v int32) *CreateRouteMapRuleRequest
	GetPriority() *int32
	SetRouteMapId(v string) *CreateRouteMapRuleRequest
	GetRouteMapId() *string
	SetSetActions(v *CreateRouteMapRuleRequestSetActions) *CreateRouteMapRuleRequest
	GetSetActions() *CreateRouteMapRuleRequestSetActions
}

type CreateRouteMapRuleRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the parameter value on your client. Make sure that the token is unique across different requests. The client token supports only ASCII characters.
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
	// Specifies whether to only precheck the request. Valid values:
	//
	// - **true**: Prechecks the request without creating the route map rule. The system checks whether required parameters are specified, the request format, and service limits. If the request fails the precheck, the corresponding error is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): Sends the request. After the request passes the precheck, the route map rule is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The match conditions of the route map rule. The system executes the actions defined in **SetActions*	- only when a route meets all the match conditions defined in this parameter.
	MatchConditions *CreateRouteMapRuleRequestMatchConditions `json:"MatchConditions,omitempty" xml:"MatchConditions,omitempty" type:"Struct"`
	// The name of the rule for easy identification. The name can be empty or 1 to 128 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// Permit-Prod-Prefixes-Set-LP
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The priority of the rule. A smaller value indicates a higher priority.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the route map.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-0a1b2c3d4e5f6a7b8
	RouteMapId *string `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
	// The actions to execute for the route map rule. After a route meets all the match conditions defined in **MatchConditions**, the system processes the route based on the actions defined in this parameter.
	SetActions *CreateRouteMapRuleRequestSetActions `json:"SetActions,omitempty" xml:"SetActions,omitempty" type:"Struct"`
}

func (s CreateRouteMapRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteMapRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateRouteMapRuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateRouteMapRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRouteMapRuleRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateRouteMapRuleRequest) GetMatchConditions() *CreateRouteMapRuleRequestMatchConditions {
	return s.MatchConditions
}

func (s *CreateRouteMapRuleRequest) GetName() *string {
	return s.Name
}

func (s *CreateRouteMapRuleRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateRouteMapRuleRequest) GetRouteMapId() *string {
	return s.RouteMapId
}

func (s *CreateRouteMapRuleRequest) GetSetActions() *CreateRouteMapRuleRequestSetActions {
	return s.SetActions
}

func (s *CreateRouteMapRuleRequest) SetClientToken(v string) *CreateRouteMapRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRouteMapRuleRequest) SetDescription(v string) *CreateRouteMapRuleRequest {
	s.Description = &v
	return s
}

func (s *CreateRouteMapRuleRequest) SetDryRun(v bool) *CreateRouteMapRuleRequest {
	s.DryRun = &v
	return s
}

func (s *CreateRouteMapRuleRequest) SetMatchConditions(v *CreateRouteMapRuleRequestMatchConditions) *CreateRouteMapRuleRequest {
	s.MatchConditions = v
	return s
}

func (s *CreateRouteMapRuleRequest) SetName(v string) *CreateRouteMapRuleRequest {
	s.Name = &v
	return s
}

func (s *CreateRouteMapRuleRequest) SetPriority(v int32) *CreateRouteMapRuleRequest {
	s.Priority = &v
	return s
}

func (s *CreateRouteMapRuleRequest) SetRouteMapId(v string) *CreateRouteMapRuleRequest {
	s.RouteMapId = &v
	return s
}

func (s *CreateRouteMapRuleRequest) SetSetActions(v *CreateRouteMapRuleRequestSetActions) *CreateRouteMapRuleRequest {
	s.SetActions = v
	return s
}

func (s *CreateRouteMapRuleRequest) Validate() error {
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

type CreateRouteMapRuleRequestMatchConditions struct {
	// The route prefix list for fuzzy match. You can specify up to 64 IP address blocks.
	//
	// A match is successful if the route prefix in the match condition contains the route prefix of the matched route.
	//
	// For example, a policy defined with 10.0.0.0/8 can fuzzy match a route with 10.10.10.0/24.
	//
	// IP address blocks use the CIDR format and support IPv4 and IPv6.
	//
	// This parameter is mutually exclusive with **AddressPrefixesMatch**. You cannot specify both parameters at the same time.
	AddressPrefixesInclude []*string `json:"AddressPrefixesInclude,omitempty" xml:"AddressPrefixesInclude,omitempty" type:"Repeated"`
	// The route prefix list for exact match. You can specify up to 64 IP address blocks.
	//
	// A match is successful only if the route prefix in the match condition is exactly the same as the route prefix of the matched route.
	//
	// For example, a policy defined with 10.0.0.0/8 can only exactly match a route with 10.0.0.0/8.
	//
	// IP address blocks use the CIDR format and support IPv4 and IPv6.
	//
	// This parameter is mutually exclusive with **AddressPrefixesInclude**. You cannot specify both parameters at the same time.
	AddressPrefixesMatch []*string `json:"AddressPrefixesMatch,omitempty" xml:"AddressPrefixesMatch,omitempty" type:"Repeated"`
	// The list of AS paths for fuzzy match. You can specify up to 64 AS numbers.
	//
	// A match is successful if the AS path in the match condition overlaps with the AS path of the matched route.
	//
	// This parameter is mutually exclusive with **AsPathsMatch**. You cannot specify both parameters at the same time.
	//
	// > Only AS SEQUENCE is supported. AS SET, AS CONFED SEQUENCE, and AS CONFED SET are not supported. This means only a list of AS numbers is supported, and sets and sublists are not supported.
	AsPathsInclude []*int64 `json:"AsPathsInclude,omitempty" xml:"AsPathsInclude,omitempty" type:"Repeated"`
	// The list of AS paths for exact match. You can specify up to 64 AS numbers.
	//
	// A match is successful only if the AS path in the match condition is exactly the same as the AS path of the matched route.
	//
	// This parameter is mutually exclusive with **AsPathsInclude**. You cannot specify both parameters at the same time.
	//
	// > Only AS SEQUENCE is supported. AS SET, AS CONFED SEQUENCE, and AS CONFED SET are not supported. This means only a list of AS numbers is supported, and sets and sublists are not supported.
	AsPathsMatch []*int64 `json:"AsPathsMatch,omitempty" xml:"AsPathsMatch,omitempty" type:"Repeated"`
	// The community set for fuzzy match. You can specify up to 64 communities.
	//
	// A match is successful if the community in the match condition overlaps with the community of the matched route.
	//
	// Each community is in the format of n:m, where n and m range from **1*	- to **65535**. The community must comply with RFC 1997. Large communities (RFC 8092) are not supported.
	//
	// This parameter is mutually exclusive with **CommunitySetMatch**. You cannot specify both parameters at the same time.
	//
	// > Incorrect community configurations may prevent routes from being advertised to the on-premises data center.
	CommunitySetInclude []*string `json:"CommunitySetInclude,omitempty" xml:"CommunitySetInclude,omitempty" type:"Repeated"`
	// The community set for exact match. You can specify up to 64 communities.
	//
	// A match is successful only if the community in the match condition is exactly the same as the community of the matched route.
	//
	// Each community is in the format of n:m, where n and m range from **1*	- to **65535**. The community must comply with RFC 1997. Large communities (RFC 8092) are not supported.
	//
	// This parameter is mutually exclusive with **CommunitySetInclude**. You cannot specify both parameters at the same time.
	//
	// > Incorrect community configurations may prevent routes from being advertised to the on-premises data center.
	CommunitySetMatch []*string `json:"CommunitySetMatch,omitempty" xml:"CommunitySetMatch,omitempty" type:"Repeated"`
	// The list of destination instance IDs to match. You can specify the following types of instance IDs:
	//
	// - Virtual Private Cloud (VPC) instance IDs
	//
	// - Virtual Border Router (VBR) instance IDs
	//
	// You can specify up to 64 instance IDs.
	//
	// > The list of destination instance IDs takes effect only when the direction of the route map is outbound.
	DestinationInstanceIds []*string `json:"DestinationInstanceIds,omitempty" xml:"DestinationInstanceIds,omitempty" type:"Repeated"`
	// Specifies whether to use the exclusion match mode for the list of destination instance IDs. Valid values:
	//
	// - **false*	- (default): No. A match is successful if the destination instance ID of the route is in **DestinationInstanceIds.N**.
	//
	// - **true**: Yes. A match is successful if the destination instance ID of the route is not in **DestinationInstanceIds.N**.
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
	//
	// You can specify multiple instance types.
	//
	// > The list of destination instance types takes effect only when the direction of the route map is outbound.
	DestinationInstanceTypes []*string `json:"DestinationInstanceTypes,omitempty" xml:"DestinationInstanceTypes,omitempty" type:"Repeated"`
	// The list of destination region IDs to match. You can specify up to 64 region IDs.
	DestinationRegionIds []*string `json:"DestinationRegionIds,omitempty" xml:"DestinationRegionIds,omitempty" type:"Repeated"`
	// The list of destination route table IDs to match. You can specify up to 64 route table IDs.
	//
	// > The list of destination route table IDs takes effect only when the direction of the route map is outbound.
	DestinationRouteTableIds []*string `json:"DestinationRouteTableIds,omitempty" xml:"DestinationRouteTableIds,omitempty" type:"Repeated"`
	// The IP address type of the routes to match. Valid values:
	//
	// - **IPv4**: Matches only IPv4 routes.
	//
	// - **IPv6**: Matches only IPv6 routes.
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
	// - **Custom**: custom route, which is manually added by the user.
	//
	// - **BGP**: BGP route, which is propagated through the BGP routable protocol.
	//
	// You can specify multiple route types.
	RouteTypes []*string `json:"RouteTypes,omitempty" xml:"RouteTypes,omitempty" type:"Repeated"`
	// The list of source instance IDs to match. You can specify the following types of instance IDs:
	//
	// - Virtual Private Cloud (VPC) instance IDs
	//
	// - Virtual Border Router (VBR) instance IDs
	//
	//
	// You can specify up to 64 instance IDs.
	SourceInstanceIds []*string `json:"SourceInstanceIds,omitempty" xml:"SourceInstanceIds,omitempty" type:"Repeated"`
	// Specifies whether to use the exclusion match mode for the list of source instance IDs. Valid values:
	//
	// - **false*	- (default): No. A match is successful if the source instance ID of the route is in **SourceInstanceIds.N**.
	//
	// - **true**: Yes. A match is successful if the source instance ID of the route is not in **SourceInstanceIds.N**.
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
	//
	//
	// You can specify multiple instance types.
	SourceInstanceTypes []*string `json:"SourceInstanceTypes,omitempty" xml:"SourceInstanceTypes,omitempty" type:"Repeated"`
	// The list of source region IDs to match. You can specify up to 64 region IDs.
	SourceRegionIds []*string `json:"SourceRegionIds,omitempty" xml:"SourceRegionIds,omitempty" type:"Repeated"`
	// The list of source route table IDs to match. You can specify up to 64 route table IDs.
	SourceRouteTableIds []*string `json:"SourceRouteTableIds,omitempty" xml:"SourceRouteTableIds,omitempty" type:"Repeated"`
}

func (s CreateRouteMapRuleRequestMatchConditions) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteMapRuleRequestMatchConditions) GoString() string {
	return s.String()
}

func (s *CreateRouteMapRuleRequestMatchConditions) GetAddressPrefixesInclude() []*string {
	return s.AddressPrefixesInclude
}

func (s *CreateRouteMapRuleRequestMatchConditions) GetAddressPrefixesMatch() []*string {
	return s.AddressPrefixesMatch
}

func (s *CreateRouteMapRuleRequestMatchConditions) GetAsPathsInclude() []*int64 {
	return s.AsPathsInclude
}

func (s *CreateRouteMapRuleRequestMatchConditions) GetAsPathsMatch() []*int64 {
	return s.AsPathsMatch
}

func (s *CreateRouteMapRuleRequestMatchConditions) GetCommunitySetInclude() []*string {
	return s.CommunitySetInclude
}

func (s *CreateRouteMapRuleRequestMatchConditions) GetCommunitySetMatch() []*string {
	return s.CommunitySetMatch
}

func (s *CreateRouteMapRuleRequestMatchConditions) GetDestinationInstanceIds() []*string {
	return s.DestinationInstanceIds
}

func (s *CreateRouteMapRuleRequestMatchConditions) GetDestinationInstanceIdsReverseMatch() *bool {
	return s.DestinationInstanceIdsReverseMatch
}

func (s *CreateRouteMapRuleRequestMatchConditions) GetDestinationInstanceTypes() []*string {
	return s.DestinationInstanceTypes
}

func (s *CreateRouteMapRuleRequestMatchConditions) GetDestinationRegionIds() []*string {
	return s.DestinationRegionIds
}

func (s *CreateRouteMapRuleRequestMatchConditions) GetDestinationRouteTableIds() []*string {
	return s.DestinationRouteTableIds
}

func (s *CreateRouteMapRuleRequestMatchConditions) GetMatchAddressType() *string {
	return s.MatchAddressType
}

func (s *CreateRouteMapRuleRequestMatchConditions) GetRouteTypes() []*string {
	return s.RouteTypes
}

func (s *CreateRouteMapRuleRequestMatchConditions) GetSourceInstanceIds() []*string {
	return s.SourceInstanceIds
}

func (s *CreateRouteMapRuleRequestMatchConditions) GetSourceInstanceIdsReverseMatch() *bool {
	return s.SourceInstanceIdsReverseMatch
}

func (s *CreateRouteMapRuleRequestMatchConditions) GetSourceInstanceTypes() []*string {
	return s.SourceInstanceTypes
}

func (s *CreateRouteMapRuleRequestMatchConditions) GetSourceRegionIds() []*string {
	return s.SourceRegionIds
}

func (s *CreateRouteMapRuleRequestMatchConditions) GetSourceRouteTableIds() []*string {
	return s.SourceRouteTableIds
}

func (s *CreateRouteMapRuleRequestMatchConditions) SetAddressPrefixesInclude(v []*string) *CreateRouteMapRuleRequestMatchConditions {
	s.AddressPrefixesInclude = v
	return s
}

func (s *CreateRouteMapRuleRequestMatchConditions) SetAddressPrefixesMatch(v []*string) *CreateRouteMapRuleRequestMatchConditions {
	s.AddressPrefixesMatch = v
	return s
}

func (s *CreateRouteMapRuleRequestMatchConditions) SetAsPathsInclude(v []*int64) *CreateRouteMapRuleRequestMatchConditions {
	s.AsPathsInclude = v
	return s
}

func (s *CreateRouteMapRuleRequestMatchConditions) SetAsPathsMatch(v []*int64) *CreateRouteMapRuleRequestMatchConditions {
	s.AsPathsMatch = v
	return s
}

func (s *CreateRouteMapRuleRequestMatchConditions) SetCommunitySetInclude(v []*string) *CreateRouteMapRuleRequestMatchConditions {
	s.CommunitySetInclude = v
	return s
}

func (s *CreateRouteMapRuleRequestMatchConditions) SetCommunitySetMatch(v []*string) *CreateRouteMapRuleRequestMatchConditions {
	s.CommunitySetMatch = v
	return s
}

func (s *CreateRouteMapRuleRequestMatchConditions) SetDestinationInstanceIds(v []*string) *CreateRouteMapRuleRequestMatchConditions {
	s.DestinationInstanceIds = v
	return s
}

func (s *CreateRouteMapRuleRequestMatchConditions) SetDestinationInstanceIdsReverseMatch(v bool) *CreateRouteMapRuleRequestMatchConditions {
	s.DestinationInstanceIdsReverseMatch = &v
	return s
}

func (s *CreateRouteMapRuleRequestMatchConditions) SetDestinationInstanceTypes(v []*string) *CreateRouteMapRuleRequestMatchConditions {
	s.DestinationInstanceTypes = v
	return s
}

func (s *CreateRouteMapRuleRequestMatchConditions) SetDestinationRegionIds(v []*string) *CreateRouteMapRuleRequestMatchConditions {
	s.DestinationRegionIds = v
	return s
}

func (s *CreateRouteMapRuleRequestMatchConditions) SetDestinationRouteTableIds(v []*string) *CreateRouteMapRuleRequestMatchConditions {
	s.DestinationRouteTableIds = v
	return s
}

func (s *CreateRouteMapRuleRequestMatchConditions) SetMatchAddressType(v string) *CreateRouteMapRuleRequestMatchConditions {
	s.MatchAddressType = &v
	return s
}

func (s *CreateRouteMapRuleRequestMatchConditions) SetRouteTypes(v []*string) *CreateRouteMapRuleRequestMatchConditions {
	s.RouteTypes = v
	return s
}

func (s *CreateRouteMapRuleRequestMatchConditions) SetSourceInstanceIds(v []*string) *CreateRouteMapRuleRequestMatchConditions {
	s.SourceInstanceIds = v
	return s
}

func (s *CreateRouteMapRuleRequestMatchConditions) SetSourceInstanceIdsReverseMatch(v bool) *CreateRouteMapRuleRequestMatchConditions {
	s.SourceInstanceIdsReverseMatch = &v
	return s
}

func (s *CreateRouteMapRuleRequestMatchConditions) SetSourceInstanceTypes(v []*string) *CreateRouteMapRuleRequestMatchConditions {
	s.SourceInstanceTypes = v
	return s
}

func (s *CreateRouteMapRuleRequestMatchConditions) SetSourceRegionIds(v []*string) *CreateRouteMapRuleRequestMatchConditions {
	s.SourceRegionIds = v
	return s
}

func (s *CreateRouteMapRuleRequestMatchConditions) SetSourceRouteTableIds(v []*string) *CreateRouteMapRuleRequestMatchConditions {
	s.SourceRouteTableIds = v
	return s
}

func (s *CreateRouteMapRuleRequestMatchConditions) Validate() error {
	return dara.Validate(s)
}

type CreateRouteMapRuleRequestSetActions struct {
	// The AS path to prepend when the regional gateway receives a route or executes a publish route entry. You can specify up to 64 AS numbers.
	//
	// This parameter is mutually exclusive with **AsPathReplace**. You cannot specify both parameters at the same time.
	//
	// This parameter specifies the action to execute after a route meets the match conditions.
	AsPathPrepend []*int64 `json:"AsPathPrepend,omitempty" xml:"AsPathPrepend,omitempty" type:"Repeated"`
	// The AS path to replace when the regional gateway receives a route or executes a publish route entry. The original AS path of the route is replaced with the list of AS numbers specified in this parameter. You can specify up to 64 AS numbers.
	//
	// This parameter is mutually exclusive with **AsPathPrepend**. You cannot specify both parameters at the same time.
	//
	// This parameter specifies the action to execute after a route meets the match conditions.
	AsPathReplace []*int64 `json:"AsPathReplace,omitempty" xml:"AsPathReplace,omitempty" type:"Repeated"`
	// The community set to add to the route. You can specify up to 64 communities.
	//
	// Each community is in the format of n:m, where n and m range from **1*	- to **65535**. The community must comply with RFC 1997. Large communities (RFC 8092) are not supported.
	//
	// This parameter is mutually exclusive with **CommunityReplace**. You cannot specify both parameters at the same time.
	//
	// This parameter specifies the action to execute after a route meets the match conditions.
	//
	// > Incorrect community configurations may prevent routes from being advertised to the on-premises data center.
	CommunityAdd []*string `json:"CommunityAdd,omitempty" xml:"CommunityAdd,omitempty" type:"Repeated"`
	// The community set to replace the original community of the route. You can specify up to 64 communities.
	//
	// Each community is in the format of n:m, where n and m range from **1*	- to **65535**. The community must comply with RFC 1997. Large communities (RFC 8092) are not supported.
	//
	// This parameter is mutually exclusive with **CommunityAdd**. You cannot specify both parameters at the same time.
	//
	// This parameter specifies the action to execute after a route meets the match conditions.
	//
	// > Incorrect community configurations may prevent routes from being advertised to the on-premises data center.
	CommunityReplace []*string `json:"CommunityReplace,omitempty" xml:"CommunityReplace,omitempty" type:"Repeated"`
	// Policy priority of the associated next route map rule.
	//
	// - You can set policy priority of the associated route map rule only when **RouteAction*	- is set to **Permit**. Only the routes that are permitted continue to match the associated next route map rule.
	//
	// - Policy priority of the associated next route map rule must be lower than policy priority of the current route map rule.
	//
	// example:
	//
	// 1
	NextPriority *int64 `json:"NextPriority,omitempty" xml:"NextPriority,omitempty"`
	// The policy action after all conditions are met. Valid values:
	//
	// - **Permit**: Permits the matched routes.
	//
	// - **Deny**: Denies the matched routes.
	//
	// example:
	//
	// Permit
	RouteAction *string `json:"RouteAction,omitempty" xml:"RouteAction,omitempty"`
}

func (s CreateRouteMapRuleRequestSetActions) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteMapRuleRequestSetActions) GoString() string {
	return s.String()
}

func (s *CreateRouteMapRuleRequestSetActions) GetAsPathPrepend() []*int64 {
	return s.AsPathPrepend
}

func (s *CreateRouteMapRuleRequestSetActions) GetAsPathReplace() []*int64 {
	return s.AsPathReplace
}

func (s *CreateRouteMapRuleRequestSetActions) GetCommunityAdd() []*string {
	return s.CommunityAdd
}

func (s *CreateRouteMapRuleRequestSetActions) GetCommunityReplace() []*string {
	return s.CommunityReplace
}

func (s *CreateRouteMapRuleRequestSetActions) GetNextPriority() *int64 {
	return s.NextPriority
}

func (s *CreateRouteMapRuleRequestSetActions) GetRouteAction() *string {
	return s.RouteAction
}

func (s *CreateRouteMapRuleRequestSetActions) SetAsPathPrepend(v []*int64) *CreateRouteMapRuleRequestSetActions {
	s.AsPathPrepend = v
	return s
}

func (s *CreateRouteMapRuleRequestSetActions) SetAsPathReplace(v []*int64) *CreateRouteMapRuleRequestSetActions {
	s.AsPathReplace = v
	return s
}

func (s *CreateRouteMapRuleRequestSetActions) SetCommunityAdd(v []*string) *CreateRouteMapRuleRequestSetActions {
	s.CommunityAdd = v
	return s
}

func (s *CreateRouteMapRuleRequestSetActions) SetCommunityReplace(v []*string) *CreateRouteMapRuleRequestSetActions {
	s.CommunityReplace = v
	return s
}

func (s *CreateRouteMapRuleRequestSetActions) SetNextPriority(v int64) *CreateRouteMapRuleRequestSetActions {
	s.NextPriority = &v
	return s
}

func (s *CreateRouteMapRuleRequestSetActions) SetRouteAction(v string) *CreateRouteMapRuleRequestSetActions {
	s.RouteAction = &v
	return s
}

func (s *CreateRouteMapRuleRequestSetActions) Validate() error {
	return dara.Validate(s)
}
