// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRouteMapRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateRouteMapRuleShrinkRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateRouteMapRuleShrinkRequest
	GetDescription() *string
	SetDryRun(v bool) *UpdateRouteMapRuleShrinkRequest
	GetDryRun() *bool
	SetMatchConditionsShrink(v string) *UpdateRouteMapRuleShrinkRequest
	GetMatchConditionsShrink() *string
	SetName(v string) *UpdateRouteMapRuleShrinkRequest
	GetName() *string
	SetPriority(v int32) *UpdateRouteMapRuleShrinkRequest
	GetPriority() *int32
	SetRouteMapRuleId(v string) *UpdateRouteMapRuleShrinkRequest
	GetRouteMapRuleId() *string
	SetSetActionsShrink(v string) *UpdateRouteMapRuleShrinkRequest
	GetSetActionsShrink() *string
}

type UpdateRouteMapRuleShrinkRequest struct {
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
	MatchConditionsShrink *string `json:"MatchConditions,omitempty" xml:"MatchConditions,omitempty"`
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
	SetActionsShrink *string `json:"SetActions,omitempty" xml:"SetActions,omitempty"`
}

func (s UpdateRouteMapRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRouteMapRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateRouteMapRuleShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateRouteMapRuleShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateRouteMapRuleShrinkRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateRouteMapRuleShrinkRequest) GetMatchConditionsShrink() *string {
	return s.MatchConditionsShrink
}

func (s *UpdateRouteMapRuleShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateRouteMapRuleShrinkRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateRouteMapRuleShrinkRequest) GetRouteMapRuleId() *string {
	return s.RouteMapRuleId
}

func (s *UpdateRouteMapRuleShrinkRequest) GetSetActionsShrink() *string {
	return s.SetActionsShrink
}

func (s *UpdateRouteMapRuleShrinkRequest) SetClientToken(v string) *UpdateRouteMapRuleShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateRouteMapRuleShrinkRequest) SetDescription(v string) *UpdateRouteMapRuleShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateRouteMapRuleShrinkRequest) SetDryRun(v bool) *UpdateRouteMapRuleShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateRouteMapRuleShrinkRequest) SetMatchConditionsShrink(v string) *UpdateRouteMapRuleShrinkRequest {
	s.MatchConditionsShrink = &v
	return s
}

func (s *UpdateRouteMapRuleShrinkRequest) SetName(v string) *UpdateRouteMapRuleShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateRouteMapRuleShrinkRequest) SetPriority(v int32) *UpdateRouteMapRuleShrinkRequest {
	s.Priority = &v
	return s
}

func (s *UpdateRouteMapRuleShrinkRequest) SetRouteMapRuleId(v string) *UpdateRouteMapRuleShrinkRequest {
	s.RouteMapRuleId = &v
	return s
}

func (s *UpdateRouteMapRuleShrinkRequest) SetSetActionsShrink(v string) *UpdateRouteMapRuleShrinkRequest {
	s.SetActionsShrink = &v
	return s
}

func (s *UpdateRouteMapRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
