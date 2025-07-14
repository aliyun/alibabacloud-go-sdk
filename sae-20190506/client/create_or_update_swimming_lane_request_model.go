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
	AppEntryRule *CreateOrUpdateSwimmingLaneRequestAppEntryRule `json:"AppEntryRule,omitempty" xml:"AppEntryRule,omitempty" type:"Struct"`
	// example:
	//
	// 0
	CanaryModel *int32 `json:"CanaryModel,omitempty" xml:"CanaryModel,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// b2a8a925-477a-eswa-b823-d5e22500****
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// 13857
	LaneId *int64 `json:"LaneId,omitempty" xml:"LaneId,omitempty"`
	// example:
	//
	// test
	LaneName *string `json:"LaneName,omitempty" xml:"LaneName,omitempty"`
	// example:
	//
	// {"alicloud.service.tag":"g1"}
	LaneTag             *string                                               `json:"LaneTag,omitempty" xml:"LaneTag,omitempty"`
	MseGatewayEntryRule *CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRule `json:"MseGatewayEntryRule,omitempty" xml:"MseGatewayEntryRule,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type CreateOrUpdateSwimmingLaneRequestAppEntryRule struct {
	// example:
	//
	// AND
	ConditionJoiner *string                                                    `json:"ConditionJoiner,omitempty" xml:"ConditionJoiner,omitempty"`
	Conditions      []*CreateOrUpdateSwimmingLaneRequestAppEntryRuleConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// example:
	//
	// true
	IndependentPercentageEnable *bool     `json:"IndependentPercentageEnable,omitempty" xml:"IndependentPercentageEnable,omitempty"`
	Paths                       []*string `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
	// example:
	//
	// 50
	Percentage       *int32            `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
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
	return dara.Validate(s)
}

type CreateOrUpdateSwimmingLaneRequestAppEntryRuleConditions struct {
	// example:
	//
	// ==
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// example:
	//
	// t
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Header
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	// example:
	//
	// AND
	ConditionJoiner *string                                                           `json:"ConditionJoiner,omitempty" xml:"ConditionJoiner,omitempty"`
	Conditions      []*CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRuleConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// example:
	//
	// true
	IndependentPercentageEnable *bool `json:"IndependentPercentageEnable,omitempty" xml:"IndependentPercentageEnable,omitempty"`
	// example:
	//
	// 100
	Percentage        *int32            `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	PercentageByRoute map[string]*int32 `json:"PercentageByRoute,omitempty" xml:"PercentageByRoute,omitempty"`
	RouteIds          []*int64          `json:"RouteIds,omitempty" xml:"RouteIds,omitempty" type:"Repeated"`
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
	return dara.Validate(s)
}

type CreateOrUpdateSwimmingLaneRequestMseGatewayEntryRuleConditions struct {
	// example:
	//
	// ==
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// header
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
