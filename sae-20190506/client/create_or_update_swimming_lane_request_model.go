// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateSwimmingLaneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppEntryRule(v *CreateOrUpdateSwimmingLaneRequestAppEntryRule) *CreateOrUpdateSwimmingLaneRequest
	GetAppEntryRule() *CreateOrUpdateSwimmingLaneRequestAppEntryRule
	SetCanaryModel(v int32) *CreateOrUpdateSwimmingLaneRequest
	GetCanaryModel() *int32
	SetEnable(v bool) *CreateOrUpdateSwimmingLaneRequest
	GetEnable() *bool
	SetGroupId(v int64) *CreateOrUpdateSwimmingLaneRequest
	GetGroupId() *int64
	SetLaneId(v int64) *CreateOrUpdateSwimmingLaneRequest
	GetLaneId() *int64
	SetLaneName(v string) *CreateOrUpdateSwimmingLaneRequest
	GetLaneName() *string
	SetLaneTag(v string) *CreateOrUpdateSwimmingLaneRequest
	GetLaneTag() *string
	SetMseGatewayEntryRule(v *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRule) *CreateOrUpdateSwimmingLaneRequest
	GetMseGatewayEntryRule() *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRule
	SetNamespaceId(v string) *CreateOrUpdateSwimmingLaneRequest
	GetNamespaceId() *string
}

type CreateOrUpdateSwimmingLaneRequest struct {
	// The route configuration of the gateway.
	//
	// >  This parameter is required if the gateway entry of the lane group is Java.
	AppEntryRule *CreateOrUpdateSwimmingLaneRequestAppEntryRule `json:"AppEntryRule,omitempty" xml:"AppEntryRule,omitempty" type:"Struct"`
	// Full-link Grayscale Mode:
	//
	// 	- 0: The request is routed based on the content of the request.
	//
	// 	- 1: routing based on percentages
	//
	// example:
	//
	// 0
	CanaryModel *int32 `json:"CanaryModel,omitempty" xml:"CanaryModel,omitempty"`
	// Lane Status
	//
	// 	- true: enabled
	//
	// 	- false: disabled
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The ID of the lane group to which the lane belongs.
	//
	// example:
	//
	// b2a8a925-477a-eswa-b823-d5e22500****
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the lane.
	//
	// example:
	//
	// 13857
	LaneId *int64 `json:"LaneId,omitempty" xml:"LaneId,omitempty"`
	// The name of the lane.
	//
	// example:
	//
	// test
	LaneName *string `json:"LaneName,omitempty" xml:"LaneName,omitempty"`
	// The tag of the lane.
	//
	// example:
	//
	// {"alicloud.service.tag":"g1"}
	LaneTag *string `json:"LaneTag,omitempty" xml:"LaneTag,omitempty"`
	// The route configuration of the MSE gateway.
	//
	// >  If the **EntryAppType*	- is set to **apig*	- or **mse-gw**, it is required.
	MseGatewayEntryRule *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRule `json:"MseGatewayEntryRule,omitempty" xml:"MseGatewayEntryRule,omitempty" type:"Struct"`
	// The namespace ID.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s CreateOrUpdateSwimmingLaneRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSwimmingLaneRequest) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSwimmingLaneRequest) GetAppEntryRule() *CreateOrUpdateSwimmingLaneRequestAppEntryRule {
	return s.AppEntryRule
}

func (s *CreateOrUpdateSwimmingLaneRequest) GetCanaryModel() *int32 {
	return s.CanaryModel
}

func (s *CreateOrUpdateSwimmingLaneRequest) GetEnable() *bool {
	return s.Enable
}

func (s *CreateOrUpdateSwimmingLaneRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *CreateOrUpdateSwimmingLaneRequest) GetLaneId() *int64 {
	return s.LaneId
}

func (s *CreateOrUpdateSwimmingLaneRequest) GetLaneName() *string {
	return s.LaneName
}

func (s *CreateOrUpdateSwimmingLaneRequest) GetLaneTag() *string {
	return s.LaneTag
}

func (s *CreateOrUpdateSwimmingLaneRequest) GetMseGatewayEntryRule() *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRule {
	return s.MseGatewayEntryRule
}

func (s *CreateOrUpdateSwimmingLaneRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *CreateOrUpdateSwimmingLaneRequest) SetAppEntryRule(v *CreateOrUpdateSwimmingLaneRequestAppEntryRule) *CreateOrUpdateSwimmingLaneRequest {
	s.AppEntryRule = v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequest) SetCanaryModel(v int32) *CreateOrUpdateSwimmingLaneRequest {
	s.CanaryModel = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequest) SetEnable(v bool) *CreateOrUpdateSwimmingLaneRequest {
	s.Enable = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequest) SetGroupId(v int64) *CreateOrUpdateSwimmingLaneRequest {
	s.GroupId = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequest) SetLaneId(v int64) *CreateOrUpdateSwimmingLaneRequest {
	s.LaneId = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequest) SetLaneName(v string) *CreateOrUpdateSwimmingLaneRequest {
	s.LaneName = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequest) SetLaneTag(v string) *CreateOrUpdateSwimmingLaneRequest {
	s.LaneTag = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequest) SetMseGatewayEntryRule(v *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRule) *CreateOrUpdateSwimmingLaneRequest {
	s.MseGatewayEntryRule = v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequest) SetNamespaceId(v string) *CreateOrUpdateSwimmingLaneRequest {
	s.NamespaceId = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequest) Validate() error {
	if s.AppEntryRule != nil {
		if err := s.AppEntryRule.Validate(); err != nil {
			return err
		}
	}
	if s.MseGatewayEntryRule != nil {
		if err := s.MseGatewayEntryRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateOrUpdateSwimmingLaneRequestAppEntryRule struct {
	// Logical connectors between conditions:
	//
	// 	- AND: All conditions are met at the same time.
	//
	// 	- OR: Any condition is met.
	//
	// example:
	//
	// AND
	ConditionJoiner *string `json:"ConditionJoiner,omitempty" xml:"ConditionJoiner,omitempty"`
	// The conditions that trigger circuit breaking.
	Conditions []*CreateOrUpdateSwimmingLaneRequestAppEntryRuleConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// Whether to enable proportional grayscale.
	//
	// 	- true: enabled. After you enable this parameter, you must configure the PercentageByPath.
	//
	// 	- false: disables the service.
	//
	// example:
	//
	// true
	IndependentPercentageEnable *bool `json:"IndependentPercentageEnable,omitempty" xml:"IndependentPercentageEnable,omitempty"`
	// The matched request path.
	Paths []*string `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
	// The traffic ratio. Valid values: 0 to 100.
	//
	// example:
	//
	// 50
	Percentage *int32 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	// The traffic configuration.
	PercentageByPath map[string]*int32 `json:"PercentageByPath,omitempty" xml:"PercentageByPath,omitempty"`
}

func (s CreateOrUpdateSwimmingLaneRequestAppEntryRule) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSwimmingLaneRequestAppEntryRule) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSwimmingLaneRequestAppEntryRule) GetConditionJoiner() *string {
	return s.ConditionJoiner
}

func (s *CreateOrUpdateSwimmingLaneRequestAppEntryRule) GetConditions() []*CreateOrUpdateSwimmingLaneRequestAppEntryRuleConditions {
	return s.Conditions
}

func (s *CreateOrUpdateSwimmingLaneRequestAppEntryRule) GetIndependentPercentageEnable() *bool {
	return s.IndependentPercentageEnable
}

func (s *CreateOrUpdateSwimmingLaneRequestAppEntryRule) GetPaths() []*string {
	return s.Paths
}

func (s *CreateOrUpdateSwimmingLaneRequestAppEntryRule) GetPercentage() *int32 {
	return s.Percentage
}

func (s *CreateOrUpdateSwimmingLaneRequestAppEntryRule) GetPercentageByPath() map[string]*int32 {
	return s.PercentageByPath
}

func (s *CreateOrUpdateSwimmingLaneRequestAppEntryRule) SetConditionJoiner(v string) *CreateOrUpdateSwimmingLaneRequestAppEntryRule {
	s.ConditionJoiner = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestAppEntryRule) SetConditions(v []*CreateOrUpdateSwimmingLaneRequestAppEntryRuleConditions) *CreateOrUpdateSwimmingLaneRequestAppEntryRule {
	s.Conditions = v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestAppEntryRule) SetIndependentPercentageEnable(v bool) *CreateOrUpdateSwimmingLaneRequestAppEntryRule {
	s.IndependentPercentageEnable = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestAppEntryRule) SetPaths(v []*string) *CreateOrUpdateSwimmingLaneRequestAppEntryRule {
	s.Paths = v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestAppEntryRule) SetPercentage(v int32) *CreateOrUpdateSwimmingLaneRequestAppEntryRule {
	s.Percentage = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestAppEntryRule) SetPercentageByPath(v map[string]*int32) *CreateOrUpdateSwimmingLaneRequestAppEntryRule {
	s.PercentageByPath = v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestAppEntryRule) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateOrUpdateSwimmingLaneRequestAppEntryRuleConditions struct {
	// Matching Rule:
	//
	// 	- The exact match. The condition is met if the traffic value and the condition value are exactly the same.
	//
	// 	- The exact match. The condition is met if the traffic value and the condition value are exactly the same.
	//
	// 	- The inclusive match. The condition is met if the traffic value is included in the specified list.
	//
	// 	- The percentage match. Principle: The condition is met if \\"hash(get(`key`)) % 100 < value\\".
	//
	// 	- Regular match: a regular expression match. The condition is met when the match is based on regular expression rules.
	//
	// example:
	//
	// ==
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The parameter name.
	//
	// example:
	//
	// t
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameter type. Valid values:
	//
	// 	- header
	//
	// 	- param
	//
	// 	- Cookie: forwards requests based on cookies.
	//
	// example:
	//
	// Header
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The match value of the condition.
	//
	// example:
	//
	// g1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateOrUpdateSwimmingLaneRequestAppEntryRuleConditions) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSwimmingLaneRequestAppEntryRuleConditions) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSwimmingLaneRequestAppEntryRuleConditions) GetCondition() *string {
	return s.Condition
}

func (s *CreateOrUpdateSwimmingLaneRequestAppEntryRuleConditions) GetName() *string {
	return s.Name
}

func (s *CreateOrUpdateSwimmingLaneRequestAppEntryRuleConditions) GetType() *string {
	return s.Type
}

func (s *CreateOrUpdateSwimmingLaneRequestAppEntryRuleConditions) GetValue() *string {
	return s.Value
}

func (s *CreateOrUpdateSwimmingLaneRequestAppEntryRuleConditions) SetCondition(v string) *CreateOrUpdateSwimmingLaneRequestAppEntryRuleConditions {
	s.Condition = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestAppEntryRuleConditions) SetName(v string) *CreateOrUpdateSwimmingLaneRequestAppEntryRuleConditions {
	s.Name = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestAppEntryRuleConditions) SetType(v string) *CreateOrUpdateSwimmingLaneRequestAppEntryRuleConditions {
	s.Type = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestAppEntryRuleConditions) SetValue(v string) *CreateOrUpdateSwimmingLaneRequestAppEntryRuleConditions {
	s.Value = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestAppEntryRuleConditions) Validate() error {
	return dara.Validate(s)
}

type CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRule struct {
	// Logical connectors between conditions:
	//
	// 	- AND: All conditions are met at the same time.
	//
	// 	- OR: Any condition is met.
	//
	// example:
	//
	// AND
	ConditionJoiner *string `json:"ConditionJoiner,omitempty" xml:"ConditionJoiner,omitempty"`
	// Routing Condition
	Conditions []*CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRuleConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// Whether to enable proportional grayscale.
	//
	// 	- true: Enabled. After you enable this parameter, you must configure the PercentageByPath.
	//
	// 	- false: Disabled.
	//
	// example:
	//
	// true
	IndependentPercentageEnable *bool `json:"IndependentPercentageEnable,omitempty" xml:"IndependentPercentageEnable,omitempty"`
	// The percentage of traffic replication. Valid values: 0 to 100.
	//
	// example:
	//
	// 100
	Percentage *int32 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	// The traffic configuration.
	PercentageByRoute map[string]*int32 `json:"PercentageByRoute,omitempty" xml:"PercentageByRoute,omitempty"`
	// The ID of the route.
	RouteIds []*int64 `json:"RouteIds,omitempty" xml:"RouteIds,omitempty" type:"Repeated"`
}

func (s CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRule) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRule) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRule) GetConditionJoiner() *string {
	return s.ConditionJoiner
}

func (s *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRule) GetConditions() []*CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRuleConditions {
	return s.Conditions
}

func (s *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRule) GetIndependentPercentageEnable() *bool {
	return s.IndependentPercentageEnable
}

func (s *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRule) GetPercentage() *int32 {
	return s.Percentage
}

func (s *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRule) GetPercentageByRoute() map[string]*int32 {
	return s.PercentageByRoute
}

func (s *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRule) GetRouteIds() []*int64 {
	return s.RouteIds
}

func (s *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRule) SetConditionJoiner(v string) *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRule {
	s.ConditionJoiner = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRule) SetConditions(v []*CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRuleConditions) *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRule {
	s.Conditions = v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRule) SetIndependentPercentageEnable(v bool) *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRule {
	s.IndependentPercentageEnable = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRule) SetPercentage(v int32) *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRule {
	s.Percentage = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRule) SetPercentageByRoute(v map[string]*int32) *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRule {
	s.PercentageByRoute = v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRule) SetRouteIds(v []*int64) *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRule {
	s.RouteIds = v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRule) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRuleConditions struct {
	// The matching rule.
	//
	// 	- \\==: exact match.
	//
	// 	- ! =: exact match.
	//
	// 	- in: contains matches.
	//
	// 	- Percentage: Percentage matching.
	//
	// 	- Regular matching: specifies whether a regular expression is used to search for the original string.
	//
	// example:
	//
	// ==
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The parameter name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameter type. Valid values:
	//
	// 	- header
	//
	// 	- param
	//
	// 	- Cookie: forwards requests based on cookies.
	//
	// example:
	//
	// header
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The match value of the condition.
	//
	// example:
	//
	// g1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRuleConditions) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRuleConditions) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRuleConditions) GetCondition() *string {
	return s.Condition
}

func (s *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRuleConditions) GetName() *string {
	return s.Name
}

func (s *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRuleConditions) GetType() *string {
	return s.Type
}

func (s *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRuleConditions) GetValue() *string {
	return s.Value
}

func (s *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRuleConditions) SetCondition(v string) *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRuleConditions {
	s.Condition = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRuleConditions) SetName(v string) *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRuleConditions {
	s.Name = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRuleConditions) SetType(v string) *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRuleConditions {
	s.Type = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRuleConditions) SetValue(v string) *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRuleConditions {
	s.Value = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRuleConditions) Validate() error {
	return dara.Validate(s)
}
