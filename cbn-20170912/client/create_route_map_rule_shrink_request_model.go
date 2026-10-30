// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRouteMapRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateRouteMapRuleShrinkRequest
	GetClientToken() *string
	SetDescription(v string) *CreateRouteMapRuleShrinkRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateRouteMapRuleShrinkRequest
	GetDryRun() *bool
	SetMatchConditionsShrink(v string) *CreateRouteMapRuleShrinkRequest
	GetMatchConditionsShrink() *string
	SetName(v string) *CreateRouteMapRuleShrinkRequest
	GetName() *string
	SetPriority(v int32) *CreateRouteMapRuleShrinkRequest
	GetPriority() *int32
	SetRouteMapId(v string) *CreateRouteMapRuleShrinkRequest
	GetRouteMapId() *string
	SetSetActionsShrink(v string) *CreateRouteMapRuleShrinkRequest
	GetSetActionsShrink() *string
}

type CreateRouteMapRuleShrinkRequest struct {
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
	MatchConditionsShrink *string `json:"MatchConditions,omitempty" xml:"MatchConditions,omitempty"`
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
	SetActionsShrink *string `json:"SetActions,omitempty" xml:"SetActions,omitempty"`
}

func (s CreateRouteMapRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteMapRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateRouteMapRuleShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateRouteMapRuleShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRouteMapRuleShrinkRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateRouteMapRuleShrinkRequest) GetMatchConditionsShrink() *string {
	return s.MatchConditionsShrink
}

func (s *CreateRouteMapRuleShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateRouteMapRuleShrinkRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateRouteMapRuleShrinkRequest) GetRouteMapId() *string {
	return s.RouteMapId
}

func (s *CreateRouteMapRuleShrinkRequest) GetSetActionsShrink() *string {
	return s.SetActionsShrink
}

func (s *CreateRouteMapRuleShrinkRequest) SetClientToken(v string) *CreateRouteMapRuleShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRouteMapRuleShrinkRequest) SetDescription(v string) *CreateRouteMapRuleShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateRouteMapRuleShrinkRequest) SetDryRun(v bool) *CreateRouteMapRuleShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *CreateRouteMapRuleShrinkRequest) SetMatchConditionsShrink(v string) *CreateRouteMapRuleShrinkRequest {
	s.MatchConditionsShrink = &v
	return s
}

func (s *CreateRouteMapRuleShrinkRequest) SetName(v string) *CreateRouteMapRuleShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateRouteMapRuleShrinkRequest) SetPriority(v int32) *CreateRouteMapRuleShrinkRequest {
	s.Priority = &v
	return s
}

func (s *CreateRouteMapRuleShrinkRequest) SetRouteMapId(v string) *CreateRouteMapRuleShrinkRequest {
	s.RouteMapId = &v
	return s
}

func (s *CreateRouteMapRuleShrinkRequest) SetSetActionsShrink(v string) *CreateRouteMapRuleShrinkRequest {
	s.SetActionsShrink = &v
	return s
}

func (s *CreateRouteMapRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
