// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllSwimmingLanesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAllSwimmingLanesResponseBody
	GetCode() *string
	SetData(v []*ListAllSwimmingLanesResponseBodyData) *ListAllSwimmingLanesResponseBody
	GetData() []*ListAllSwimmingLanesResponseBodyData
	SetErrorCode(v string) *ListAllSwimmingLanesResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ListAllSwimmingLanesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAllSwimmingLanesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAllSwimmingLanesResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *ListAllSwimmingLanesResponseBody
	GetTraceId() *string
}

type ListAllSwimmingLanesResponseBody struct {
	// The API status code or POP error code. Valid values:
	//
	// - **2xx**: Success.
	//
	// - **3xx**: Redirection.
	//
	// - **4xx**: Client error.
	//
	// - **5xx**: Server error.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data []*ListAllSwimmingLanesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code. Valid values:
	//
	// - If the request is successful, the **ErrorCode*	- parameter is not returned.
	//
	// - If the request fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	//
	// example:
	//
	// InvalidParameter.NotEmpty
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information. Valid values:
	//
	// - If the request is successful, **success*	- is returned.
	//
	// - If the request fails, a specific error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B4D805CA-926D-41B1-8E63-7AD0C1ED****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - **true**: The restart is successful.
	//
	// - **false**: The restart fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID used to query invocation details.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ListAllSwimmingLanesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAllSwimmingLanesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAllSwimmingLanesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAllSwimmingLanesResponseBody) GetData() []*ListAllSwimmingLanesResponseBodyData {
	return s.Data
}

func (s *ListAllSwimmingLanesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListAllSwimmingLanesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAllSwimmingLanesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAllSwimmingLanesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAllSwimmingLanesResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ListAllSwimmingLanesResponseBody) SetCode(v string) *ListAllSwimmingLanesResponseBody {
	s.Code = &v
	return s
}

func (s *ListAllSwimmingLanesResponseBody) SetData(v []*ListAllSwimmingLanesResponseBodyData) *ListAllSwimmingLanesResponseBody {
	s.Data = v
	return s
}

func (s *ListAllSwimmingLanesResponseBody) SetErrorCode(v string) *ListAllSwimmingLanesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAllSwimmingLanesResponseBody) SetMessage(v string) *ListAllSwimmingLanesResponseBody {
	s.Message = &v
	return s
}

func (s *ListAllSwimmingLanesResponseBody) SetRequestId(v string) *ListAllSwimmingLanesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAllSwimmingLanesResponseBody) SetSuccess(v bool) *ListAllSwimmingLanesResponseBody {
	s.Success = &v
	return s
}

func (s *ListAllSwimmingLanesResponseBody) SetTraceId(v string) *ListAllSwimmingLanesResponseBody {
	s.TraceId = &v
	return s
}

func (s *ListAllSwimmingLanesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAllSwimmingLanesResponseBodyData struct {
	// The application entry rule.
	AppEntryRule *ListAllSwimmingLanesResponseBodyDataAppEntryRule `json:"AppEntryRule,omitempty" xml:"AppEntryRule,omitempty" type:"Struct"`
	// The associated applications.
	Apps []*ListAllSwimmingLanesResponseBodyDataApps `json:"Apps,omitempty" xml:"Apps,omitempty" type:"Repeated"`
	// The end-to-end canary release mode. Valid values:
	//
	//
	// - 0: Routes traffic based on request content.
	//
	// - 1: Routes traffic based on proportion.
	//
	// example:
	//
	// 0
	CanaryModel *int32 `json:"CanaryModel,omitempty" xml:"CanaryModel,omitempty"`
	// The traffic lane status. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// Indicates whether traffic rules are enabled.
	//
	// example:
	//
	// true
	EnableRules *bool `json:"EnableRules,omitempty" xml:"EnableRules,omitempty"`
	// The traffic lane ID.
	//
	// example:
	//
	// 16401
	LaneId *int64 `json:"LaneId,omitempty" xml:"LaneId,omitempty"`
	// The traffic lane group name.
	//
	// example:
	//
	// test
	LaneName *string `json:"LaneName,omitempty" xml:"LaneName,omitempty"`
	// The traffic lane tag.
	//
	// example:
	//
	// g1
	LaneTag *string `json:"LaneTag,omitempty" xml:"LaneTag,omitempty"`
	// The MSE gateway route.
	MseGatewayEntryRule *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRule `json:"MseGatewayEntryRule,omitempty" xml:"MseGatewayEntryRule,omitempty" type:"Struct"`
}

func (s ListAllSwimmingLanesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAllSwimmingLanesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAllSwimmingLanesResponseBodyData) GetAppEntryRule() *ListAllSwimmingLanesResponseBodyDataAppEntryRule {
	return s.AppEntryRule
}

func (s *ListAllSwimmingLanesResponseBodyData) GetApps() []*ListAllSwimmingLanesResponseBodyDataApps {
	return s.Apps
}

func (s *ListAllSwimmingLanesResponseBodyData) GetCanaryModel() *int32 {
	return s.CanaryModel
}

func (s *ListAllSwimmingLanesResponseBodyData) GetEnable() *bool {
	return s.Enable
}

func (s *ListAllSwimmingLanesResponseBodyData) GetEnableRules() *bool {
	return s.EnableRules
}

func (s *ListAllSwimmingLanesResponseBodyData) GetLaneId() *int64 {
	return s.LaneId
}

func (s *ListAllSwimmingLanesResponseBodyData) GetLaneName() *string {
	return s.LaneName
}

func (s *ListAllSwimmingLanesResponseBodyData) GetLaneTag() *string {
	return s.LaneTag
}

func (s *ListAllSwimmingLanesResponseBodyData) GetMseGatewayEntryRule() *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRule {
	return s.MseGatewayEntryRule
}

func (s *ListAllSwimmingLanesResponseBodyData) SetAppEntryRule(v *ListAllSwimmingLanesResponseBodyDataAppEntryRule) *ListAllSwimmingLanesResponseBodyData {
	s.AppEntryRule = v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyData) SetApps(v []*ListAllSwimmingLanesResponseBodyDataApps) *ListAllSwimmingLanesResponseBodyData {
	s.Apps = v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyData) SetCanaryModel(v int32) *ListAllSwimmingLanesResponseBodyData {
	s.CanaryModel = &v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyData) SetEnable(v bool) *ListAllSwimmingLanesResponseBodyData {
	s.Enable = &v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyData) SetEnableRules(v bool) *ListAllSwimmingLanesResponseBodyData {
	s.EnableRules = &v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyData) SetLaneId(v int64) *ListAllSwimmingLanesResponseBodyData {
	s.LaneId = &v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyData) SetLaneName(v string) *ListAllSwimmingLanesResponseBodyData {
	s.LaneName = &v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyData) SetLaneTag(v string) *ListAllSwimmingLanesResponseBodyData {
	s.LaneTag = &v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyData) SetMseGatewayEntryRule(v *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRule) *ListAllSwimmingLanesResponseBodyData {
	s.MseGatewayEntryRule = v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyData) Validate() error {
	if s.AppEntryRule != nil {
		if err := s.AppEntryRule.Validate(); err != nil {
			return err
		}
	}
	if s.Apps != nil {
		for _, item := range s.Apps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MseGatewayEntryRule != nil {
		if err := s.MseGatewayEntryRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAllSwimmingLanesResponseBodyDataAppEntryRule struct {
	// The logical connector between conditions. Valid values:
	//
	// - AND: All conditions must be met.
	//
	// - OR: Any condition must be met.
	//
	// example:
	//
	// AND
	ConditionJoiner *string `json:"ConditionJoiner,omitempty" xml:"ConditionJoiner,omitempty"`
	// The matching conditions.
	Conditions []*ListAllSwimmingLanesResponseBodyDataAppEntryRuleConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// Indicates whether proportion-based canary release is enabled.
	//
	// example:
	//
	// true
	IndependentPercentageEnable *bool `json:"IndependentPercentageEnable,omitempty" xml:"IndependentPercentageEnable,omitempty"`
	// The URI of the request.
	Paths []*string `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
	// The traffic proportion used for proportion-based routing. Valid values: 0 to 100.
	//
	// example:
	//
	// 50
	Percentage *int32 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	// The traffic matching.
	PercentageByPath map[string]*int32 `json:"PercentageByPath,omitempty" xml:"PercentageByPath,omitempty"`
}

func (s ListAllSwimmingLanesResponseBodyDataAppEntryRule) String() string {
	return dara.Prettify(s)
}

func (s ListAllSwimmingLanesResponseBodyDataAppEntryRule) GoString() string {
	return s.String()
}

func (s *ListAllSwimmingLanesResponseBodyDataAppEntryRule) GetConditionJoiner() *string {
	return s.ConditionJoiner
}

func (s *ListAllSwimmingLanesResponseBodyDataAppEntryRule) GetConditions() []*ListAllSwimmingLanesResponseBodyDataAppEntryRuleConditions {
	return s.Conditions
}

func (s *ListAllSwimmingLanesResponseBodyDataAppEntryRule) GetIndependentPercentageEnable() *bool {
	return s.IndependentPercentageEnable
}

func (s *ListAllSwimmingLanesResponseBodyDataAppEntryRule) GetPaths() []*string {
	return s.Paths
}

func (s *ListAllSwimmingLanesResponseBodyDataAppEntryRule) GetPercentage() *int32 {
	return s.Percentage
}

func (s *ListAllSwimmingLanesResponseBodyDataAppEntryRule) GetPercentageByPath() map[string]*int32 {
	return s.PercentageByPath
}

func (s *ListAllSwimmingLanesResponseBodyDataAppEntryRule) SetConditionJoiner(v string) *ListAllSwimmingLanesResponseBodyDataAppEntryRule {
	s.ConditionJoiner = &v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyDataAppEntryRule) SetConditions(v []*ListAllSwimmingLanesResponseBodyDataAppEntryRuleConditions) *ListAllSwimmingLanesResponseBodyDataAppEntryRule {
	s.Conditions = v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyDataAppEntryRule) SetIndependentPercentageEnable(v bool) *ListAllSwimmingLanesResponseBodyDataAppEntryRule {
	s.IndependentPercentageEnable = &v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyDataAppEntryRule) SetPaths(v []*string) *ListAllSwimmingLanesResponseBodyDataAppEntryRule {
	s.Paths = v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyDataAppEntryRule) SetPercentage(v int32) *ListAllSwimmingLanesResponseBodyDataAppEntryRule {
	s.Percentage = &v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyDataAppEntryRule) SetPercentageByPath(v map[string]*int32) *ListAllSwimmingLanesResponseBodyDataAppEntryRule {
	s.PercentageByPath = v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyDataAppEntryRule) Validate() error {
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

type ListAllSwimmingLanesResponseBodyDataAppEntryRuleConditions struct {
	// The matching rule.
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
	// The parameter type.
	//
	// example:
	//
	// Header
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The matching value of the condition.
	//
	// example:
	//
	// g1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The matching value of the condition.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListAllSwimmingLanesResponseBodyDataAppEntryRuleConditions) String() string {
	return dara.Prettify(s)
}

func (s ListAllSwimmingLanesResponseBodyDataAppEntryRuleConditions) GoString() string {
	return s.String()
}

func (s *ListAllSwimmingLanesResponseBodyDataAppEntryRuleConditions) GetCondition() *string {
	return s.Condition
}

func (s *ListAllSwimmingLanesResponseBodyDataAppEntryRuleConditions) GetName() *string {
	return s.Name
}

func (s *ListAllSwimmingLanesResponseBodyDataAppEntryRuleConditions) GetType() *string {
	return s.Type
}

func (s *ListAllSwimmingLanesResponseBodyDataAppEntryRuleConditions) GetValue() *string {
	return s.Value
}

func (s *ListAllSwimmingLanesResponseBodyDataAppEntryRuleConditions) GetValues() []*string {
	return s.Values
}

func (s *ListAllSwimmingLanesResponseBodyDataAppEntryRuleConditions) SetCondition(v string) *ListAllSwimmingLanesResponseBodyDataAppEntryRuleConditions {
	s.Condition = &v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyDataAppEntryRuleConditions) SetName(v string) *ListAllSwimmingLanesResponseBodyDataAppEntryRuleConditions {
	s.Name = &v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyDataAppEntryRuleConditions) SetType(v string) *ListAllSwimmingLanesResponseBodyDataAppEntryRuleConditions {
	s.Type = &v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyDataAppEntryRuleConditions) SetValue(v string) *ListAllSwimmingLanesResponseBodyDataAppEntryRuleConditions {
	s.Value = &v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyDataAppEntryRuleConditions) SetValues(v []*string) *ListAllSwimmingLanesResponseBodyDataAppEntryRuleConditions {
	s.Values = v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyDataAppEntryRuleConditions) Validate() error {
	return dara.Validate(s)
}

type ListAllSwimmingLanesResponseBodyDataApps struct {
	// The application ID.
	//
	// example:
	//
	// 8ea0c468-8165-416d-beae-531abb******
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The MSE instance ID.
	//
	// example:
	//
	// mse-cn-53y49******
	MseAppId *string `json:"MseAppId,omitempty" xml:"MseAppId,omitempty"`
	// The MSE instance name.
	//
	// example:
	//
	// sae-test
	MseAppName *string `json:"MseAppName,omitempty" xml:"MseAppName,omitempty"`
	// The ID of the namespace where the MSE instance resides.
	//
	// example:
	//
	// sae-ent
	MseNamespaceId *string `json:"MseNamespaceId,omitempty" xml:"MseNamespaceId,omitempty"`
}

func (s ListAllSwimmingLanesResponseBodyDataApps) String() string {
	return dara.Prettify(s)
}

func (s ListAllSwimmingLanesResponseBodyDataApps) GoString() string {
	return s.String()
}

func (s *ListAllSwimmingLanesResponseBodyDataApps) GetAppId() *string {
	return s.AppId
}

func (s *ListAllSwimmingLanesResponseBodyDataApps) GetAppName() *string {
	return s.AppName
}

func (s *ListAllSwimmingLanesResponseBodyDataApps) GetMseAppId() *string {
	return s.MseAppId
}

func (s *ListAllSwimmingLanesResponseBodyDataApps) GetMseAppName() *string {
	return s.MseAppName
}

func (s *ListAllSwimmingLanesResponseBodyDataApps) GetMseNamespaceId() *string {
	return s.MseNamespaceId
}

func (s *ListAllSwimmingLanesResponseBodyDataApps) SetAppId(v string) *ListAllSwimmingLanesResponseBodyDataApps {
	s.AppId = &v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyDataApps) SetAppName(v string) *ListAllSwimmingLanesResponseBodyDataApps {
	s.AppName = &v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyDataApps) SetMseAppId(v string) *ListAllSwimmingLanesResponseBodyDataApps {
	s.MseAppId = &v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyDataApps) SetMseAppName(v string) *ListAllSwimmingLanesResponseBodyDataApps {
	s.MseAppName = &v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyDataApps) SetMseNamespaceId(v string) *ListAllSwimmingLanesResponseBodyDataApps {
	s.MseNamespaceId = &v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyDataApps) Validate() error {
	return dara.Validate(s)
}

type ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRule struct {
	// The logical connector between conditions.
	//
	// example:
	//
	// AND
	ConditionJoiner *string `json:"ConditionJoiner,omitempty" xml:"ConditionJoiner,omitempty"`
	// The matching conditions.
	Conditions []*ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// Indicates whether proportion-based canary release is enabled.
	//
	// example:
	//
	// true
	IndependentPercentageEnable *bool `json:"IndependentPercentageEnable,omitempty" xml:"IndependentPercentageEnable,omitempty"`
	// The path traffic proportion.
	//
	// example:
	//
	// 100
	Percentage *int32 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	// The traffic configuration.
	PercentageByRoute map[string]*int32 `json:"PercentageByRoute,omitempty" xml:"PercentageByRoute,omitempty"`
	// The list of route IDs.
	RouteIds []*int64 `json:"RouteIds,omitempty" xml:"RouteIds,omitempty" type:"Repeated"`
	// The detailed configuration of the route rule.
	Routes []*ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutes `json:"Routes,omitempty" xml:"Routes,omitempty" type:"Repeated"`
}

func (s ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRule) String() string {
	return dara.Prettify(s)
}

func (s ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRule) GoString() string {
	return s.String()
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRule) GetConditionJoiner() *string {
	return s.ConditionJoiner
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRule) GetConditions() []*ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleConditions {
	return s.Conditions
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRule) GetIndependentPercentageEnable() *bool {
	return s.IndependentPercentageEnable
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRule) GetPercentage() *int32 {
	return s.Percentage
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRule) GetPercentageByRoute() map[string]*int32 {
	return s.PercentageByRoute
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRule) GetRouteIds() []*int64 {
	return s.RouteIds
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRule) GetRoutes() []*ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutes {
	return s.Routes
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRule) SetConditionJoiner(v string) *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRule {
	s.ConditionJoiner = &v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRule) SetConditions(v []*ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleConditions) *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRule {
	s.Conditions = v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRule) SetIndependentPercentageEnable(v bool) *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRule {
	s.IndependentPercentageEnable = &v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRule) SetPercentage(v int32) *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRule {
	s.Percentage = &v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRule) SetPercentageByRoute(v map[string]*int32) *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRule {
	s.PercentageByRoute = v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRule) SetRouteIds(v []*int64) *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRule {
	s.RouteIds = v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRule) SetRoutes(v []*ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutes) *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRule {
	s.Routes = v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRule) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Routes != nil {
		for _, item := range s.Routes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleConditions struct {
	// The matching rule.
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
	// The parameter type.
	//
	// example:
	//
	// Header
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The matching value of the condition.
	//
	// example:
	//
	// g1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleConditions) String() string {
	return dara.Prettify(s)
}

func (s ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleConditions) GoString() string {
	return s.String()
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleConditions) GetCondition() *string {
	return s.Condition
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleConditions) GetName() *string {
	return s.Name
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleConditions) GetType() *string {
	return s.Type
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleConditions) GetValue() *string {
	return s.Value
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleConditions) SetCondition(v string) *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleConditions {
	s.Condition = &v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleConditions) SetName(v string) *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleConditions {
	s.Name = &v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleConditions) SetType(v string) *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleConditions {
	s.Type = &v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleConditions) SetValue(v string) *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleConditions {
	s.Value = &v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleConditions) Validate() error {
	return dara.Validate(s)
}

type ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutes struct {
	// The route ID.
	//
	// example:
	//
	// 9504
	RouteId *int64 `json:"RouteId,omitempty" xml:"RouteId,omitempty"`
	// The route name.
	//
	// example:
	//
	// demo
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
	// The route rule.
	RoutePredicate *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicate `json:"RoutePredicate,omitempty" xml:"RoutePredicate,omitempty" type:"Struct"`
}

func (s ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutes) String() string {
	return dara.Prettify(s)
}

func (s ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutes) GoString() string {
	return s.String()
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutes) GetRouteId() *int64 {
	return s.RouteId
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutes) GetRouteName() *string {
	return s.RouteName
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutes) GetRoutePredicate() *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicate {
	return s.RoutePredicate
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutes) SetRouteId(v int64) *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutes {
	s.RouteId = &v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutes) SetRouteName(v string) *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutes {
	s.RouteName = &v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutes) SetRoutePredicate(v *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicate) *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutes {
	s.RoutePredicate = v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutes) Validate() error {
	if s.RoutePredicate != nil {
		if err := s.RoutePredicate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicate struct {
	// The path matching rule.
	PathPredicate *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicatePathPredicate `json:"PathPredicate,omitempty" xml:"PathPredicate,omitempty" type:"Struct"`
}

func (s ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicate) String() string {
	return dara.Prettify(s)
}

func (s ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicate) GoString() string {
	return s.String()
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicate) GetPathPredicate() *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicatePathPredicate {
	return s.PathPredicate
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicate) SetPathPredicate(v *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicatePathPredicate) *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicate {
	s.PathPredicate = v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicate) Validate() error {
	if s.PathPredicate != nil {
		if err := s.PathPredicate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicatePathPredicate struct {
	// The path.
	//
	// example:
	//
	// /Path
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The matching rule.
	//
	// example:
	//
	// Header
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicatePathPredicate) String() string {
	return dara.Prettify(s)
}

func (s ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicatePathPredicate) GoString() string {
	return s.String()
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicatePathPredicate) GetPath() *string {
	return s.Path
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicatePathPredicate) GetType() *string {
	return s.Type
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicatePathPredicate) SetPath(v string) *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicatePathPredicate {
	s.Path = &v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicatePathPredicate) SetType(v string) *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicatePathPredicate {
	s.Type = &v
	return s
}

func (s *ListAllSwimmingLanesResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicatePathPredicate) Validate() error {
	return dara.Validate(s)
}
