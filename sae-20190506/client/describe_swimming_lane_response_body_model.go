// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSwimmingLaneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSwimmingLaneResponseBody
	GetCode() *string
	SetData(v *DescribeSwimmingLaneResponseBodyData) *DescribeSwimmingLaneResponseBody
	GetData() *DescribeSwimmingLaneResponseBodyData
	SetErrorCode(v string) *DescribeSwimmingLaneResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DescribeSwimmingLaneResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSwimmingLaneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeSwimmingLaneResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DescribeSwimmingLaneResponseBody
	GetTraceId() *string
}

type DescribeSwimmingLaneResponseBody struct {
	// The HTTP status code or the error code. Valid values:
	//
	// 	- **2xx**: The request was successful.
	//
	// 	- **3xx**: The request was redirected.
	//
	// 	- **4xx**: The request failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Responses.
	Data *DescribeSwimmingLaneResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The status code. Value values:
	//
	// 	- If the request was successful, **ErrorCode*	- is not returned.
	//
	// 	- If the request failed, **ErrorCode*	- is returned. For more information, see **Error codes*	- in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message. Valid values:
	//
	// 	- The error message returned because the request is normal and **success*	- is returned.
	//
	// 	- If the request is abnormal, the specific exception error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30375C38-F4ED-4135-A0AE-5C75DC7F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: Valid values:
	//
	// 	- **true**: The information was queried.
	//
	// 	- **false**: Query failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the trace. This parameter is used to query the exact call information.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeSwimmingLaneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSwimmingLaneResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSwimmingLaneResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSwimmingLaneResponseBody) GetData() *DescribeSwimmingLaneResponseBodyData {
	return s.Data
}

func (s *DescribeSwimmingLaneResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeSwimmingLaneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSwimmingLaneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSwimmingLaneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSwimmingLaneResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeSwimmingLaneResponseBody) SetCode(v string) *DescribeSwimmingLaneResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSwimmingLaneResponseBody) SetData(v *DescribeSwimmingLaneResponseBodyData) *DescribeSwimmingLaneResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSwimmingLaneResponseBody) SetErrorCode(v string) *DescribeSwimmingLaneResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeSwimmingLaneResponseBody) SetMessage(v string) *DescribeSwimmingLaneResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSwimmingLaneResponseBody) SetRequestId(v string) *DescribeSwimmingLaneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSwimmingLaneResponseBody) SetSuccess(v bool) *DescribeSwimmingLaneResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSwimmingLaneResponseBody) SetTraceId(v string) *DescribeSwimmingLaneResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeSwimmingLaneResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSwimmingLaneResponseBodyData struct {
	// Apply ingress rules.
	AppEntryRule *DescribeSwimmingLaneResponseBodyDataAppEntryRule `json:"AppEntryRule,omitempty" xml:"AppEntryRule,omitempty" type:"Struct"`
	// The apps.
	Apps []*DescribeSwimmingLaneResponseBodyDataApps `json:"Apps,omitempty" xml:"Apps,omitempty" type:"Repeated"`
	// The routing mode for end-to-end canary release.
	//
	// 	- 0: routing based on request content
	//
	// 	- 1: routing based on percentages
	//
	// example:
	//
	// 0
	CanaryModel *int32 `json:"CanaryModel,omitempty" xml:"CanaryModel,omitempty"`
	// Lane status:
	//
	// 	- true: enabled
	//
	// 	- false: disabled
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// Whether the traffic rule is enabled.
	//
	// example:
	//
	// true
	EnableRules *bool `json:"EnableRules,omitempty" xml:"EnableRules,omitempty"`
	// The ID of the lane.
	//
	// example:
	//
	// 9488
	LaneId *int64 `json:"LaneId,omitempty" xml:"LaneId,omitempty"`
	// The name of the lane group.
	//
	// example:
	//
	// mse-test
	LaneName *string `json:"LaneName,omitempty" xml:"LaneName,omitempty"`
	// The label of the lane.
	//
	// example:
	//
	// {"alicloud.service.tag":"g1"}
	LaneTag *string `json:"LaneTag,omitempty" xml:"LaneTag,omitempty"`
	// MSE gateway routes.
	MseGatewayEntryRule *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRule `json:"MseGatewayEntryRule,omitempty" xml:"MseGatewayEntryRule,omitempty" type:"Struct"`
}

func (s DescribeSwimmingLaneResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSwimmingLaneResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSwimmingLaneResponseBodyData) GetAppEntryRule() *DescribeSwimmingLaneResponseBodyDataAppEntryRule {
	return s.AppEntryRule
}

func (s *DescribeSwimmingLaneResponseBodyData) GetApps() []*DescribeSwimmingLaneResponseBodyDataApps {
	return s.Apps
}

func (s *DescribeSwimmingLaneResponseBodyData) GetCanaryModel() *int32 {
	return s.CanaryModel
}

func (s *DescribeSwimmingLaneResponseBodyData) GetEnable() *bool {
	return s.Enable
}

func (s *DescribeSwimmingLaneResponseBodyData) GetEnableRules() *bool {
	return s.EnableRules
}

func (s *DescribeSwimmingLaneResponseBodyData) GetLaneId() *int64 {
	return s.LaneId
}

func (s *DescribeSwimmingLaneResponseBodyData) GetLaneName() *string {
	return s.LaneName
}

func (s *DescribeSwimmingLaneResponseBodyData) GetLaneTag() *string {
	return s.LaneTag
}

func (s *DescribeSwimmingLaneResponseBodyData) GetMseGatewayEntryRule() *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRule {
	return s.MseGatewayEntryRule
}

func (s *DescribeSwimmingLaneResponseBodyData) SetAppEntryRule(v *DescribeSwimmingLaneResponseBodyDataAppEntryRule) *DescribeSwimmingLaneResponseBodyData {
	s.AppEntryRule = v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyData) SetApps(v []*DescribeSwimmingLaneResponseBodyDataApps) *DescribeSwimmingLaneResponseBodyData {
	s.Apps = v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyData) SetCanaryModel(v int32) *DescribeSwimmingLaneResponseBodyData {
	s.CanaryModel = &v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyData) SetEnable(v bool) *DescribeSwimmingLaneResponseBodyData {
	s.Enable = &v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyData) SetEnableRules(v bool) *DescribeSwimmingLaneResponseBodyData {
	s.EnableRules = &v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyData) SetLaneId(v int64) *DescribeSwimmingLaneResponseBodyData {
	s.LaneId = &v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyData) SetLaneName(v string) *DescribeSwimmingLaneResponseBodyData {
	s.LaneName = &v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyData) SetLaneTag(v string) *DescribeSwimmingLaneResponseBodyData {
	s.LaneTag = &v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyData) SetMseGatewayEntryRule(v *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRule) *DescribeSwimmingLaneResponseBodyData {
	s.MseGatewayEntryRule = v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyData) Validate() error {
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

type DescribeSwimmingLaneResponseBodyDataAppEntryRule struct {
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
	// The matching condition.
	Conditions []*DescribeSwimmingLaneResponseBodyDataAppEntryRuleConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// Whether to enable proportional grayscale.
	//
	// example:
	//
	// true
	IndependentPercentageEnable *bool `json:"IndependentPercentageEnable,omitempty" xml:"IndependentPercentageEnable,omitempty"`
	// The request path.
	Paths []*string `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
	// The traffic ratio. Valid values: 0 to 100.
	//
	// example:
	//
	// 50
	Percentage *int32 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	// Traffic matching.
	PercentageByPath map[string]*int32 `json:"PercentageByPath,omitempty" xml:"PercentageByPath,omitempty"`
}

func (s DescribeSwimmingLaneResponseBodyDataAppEntryRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeSwimmingLaneResponseBodyDataAppEntryRule) GoString() string {
	return s.String()
}

func (s *DescribeSwimmingLaneResponseBodyDataAppEntryRule) GetConditionJoiner() *string {
	return s.ConditionJoiner
}

func (s *DescribeSwimmingLaneResponseBodyDataAppEntryRule) GetConditions() []*DescribeSwimmingLaneResponseBodyDataAppEntryRuleConditions {
	return s.Conditions
}

func (s *DescribeSwimmingLaneResponseBodyDataAppEntryRule) GetIndependentPercentageEnable() *bool {
	return s.IndependentPercentageEnable
}

func (s *DescribeSwimmingLaneResponseBodyDataAppEntryRule) GetPaths() []*string {
	return s.Paths
}

func (s *DescribeSwimmingLaneResponseBodyDataAppEntryRule) GetPercentage() *int32 {
	return s.Percentage
}

func (s *DescribeSwimmingLaneResponseBodyDataAppEntryRule) GetPercentageByPath() map[string]*int32 {
	return s.PercentageByPath
}

func (s *DescribeSwimmingLaneResponseBodyDataAppEntryRule) SetConditionJoiner(v string) *DescribeSwimmingLaneResponseBodyDataAppEntryRule {
	s.ConditionJoiner = &v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyDataAppEntryRule) SetConditions(v []*DescribeSwimmingLaneResponseBodyDataAppEntryRuleConditions) *DescribeSwimmingLaneResponseBodyDataAppEntryRule {
	s.Conditions = v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyDataAppEntryRule) SetIndependentPercentageEnable(v bool) *DescribeSwimmingLaneResponseBodyDataAppEntryRule {
	s.IndependentPercentageEnable = &v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyDataAppEntryRule) SetPaths(v []*string) *DescribeSwimmingLaneResponseBodyDataAppEntryRule {
	s.Paths = v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyDataAppEntryRule) SetPercentage(v int32) *DescribeSwimmingLaneResponseBodyDataAppEntryRule {
	s.Percentage = &v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyDataAppEntryRule) SetPercentageByPath(v map[string]*int32) *DescribeSwimmingLaneResponseBodyDataAppEntryRule {
	s.PercentageByPath = v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyDataAppEntryRule) Validate() error {
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

type DescribeSwimmingLaneResponseBodyDataAppEntryRuleConditions struct {
	// The matching rule.
	//
	// example:
	//
	// ==
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// t
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Parameter type.
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
	// The match value of the condition.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeSwimmingLaneResponseBodyDataAppEntryRuleConditions) String() string {
	return dara.Prettify(s)
}

func (s DescribeSwimmingLaneResponseBodyDataAppEntryRuleConditions) GoString() string {
	return s.String()
}

func (s *DescribeSwimmingLaneResponseBodyDataAppEntryRuleConditions) GetCondition() *string {
	return s.Condition
}

func (s *DescribeSwimmingLaneResponseBodyDataAppEntryRuleConditions) GetName() *string {
	return s.Name
}

func (s *DescribeSwimmingLaneResponseBodyDataAppEntryRuleConditions) GetType() *string {
	return s.Type
}

func (s *DescribeSwimmingLaneResponseBodyDataAppEntryRuleConditions) GetValue() *string {
	return s.Value
}

func (s *DescribeSwimmingLaneResponseBodyDataAppEntryRuleConditions) GetValues() []*string {
	return s.Values
}

func (s *DescribeSwimmingLaneResponseBodyDataAppEntryRuleConditions) SetCondition(v string) *DescribeSwimmingLaneResponseBodyDataAppEntryRuleConditions {
	s.Condition = &v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyDataAppEntryRuleConditions) SetName(v string) *DescribeSwimmingLaneResponseBodyDataAppEntryRuleConditions {
	s.Name = &v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyDataAppEntryRuleConditions) SetType(v string) *DescribeSwimmingLaneResponseBodyDataAppEntryRuleConditions {
	s.Type = &v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyDataAppEntryRuleConditions) SetValue(v string) *DescribeSwimmingLaneResponseBodyDataAppEntryRuleConditions {
	s.Value = &v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyDataAppEntryRuleConditions) SetValues(v []*string) *DescribeSwimmingLaneResponseBodyDataAppEntryRuleConditions {
	s.Values = v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyDataAppEntryRuleConditions) Validate() error {
	return dara.Validate(s)
}

type DescribeSwimmingLaneResponseBodyDataApps struct {
	// The ID of the application.
	//
	// example:
	//
	// 6b4c0f64-f679-4580-8105-91eac4******
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ID of the MSE instance.
	//
	// example:
	//
	// mse-cn-53y49******
	MseAppId *string `json:"MseAppId,omitempty" xml:"MseAppId,omitempty"`
	// The name of the MSE instance.
	//
	// example:
	//
	// sae-test
	MseAppName *string `json:"MseAppName,omitempty" xml:"MseAppName,omitempty"`
	// The ID of the MSE namespace.
	//
	// example:
	//
	// 6733e538-d52f-48e6-91a4-192f91******
	MseNamespaceId *string `json:"mseNamespaceId,omitempty" xml:"mseNamespaceId,omitempty"`
}

func (s DescribeSwimmingLaneResponseBodyDataApps) String() string {
	return dara.Prettify(s)
}

func (s DescribeSwimmingLaneResponseBodyDataApps) GoString() string {
	return s.String()
}

func (s *DescribeSwimmingLaneResponseBodyDataApps) GetAppId() *string {
	return s.AppId
}

func (s *DescribeSwimmingLaneResponseBodyDataApps) GetAppName() *string {
	return s.AppName
}

func (s *DescribeSwimmingLaneResponseBodyDataApps) GetMseAppId() *string {
	return s.MseAppId
}

func (s *DescribeSwimmingLaneResponseBodyDataApps) GetMseAppName() *string {
	return s.MseAppName
}

func (s *DescribeSwimmingLaneResponseBodyDataApps) GetMseNamespaceId() *string {
	return s.MseNamespaceId
}

func (s *DescribeSwimmingLaneResponseBodyDataApps) SetAppId(v string) *DescribeSwimmingLaneResponseBodyDataApps {
	s.AppId = &v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyDataApps) SetAppName(v string) *DescribeSwimmingLaneResponseBodyDataApps {
	s.AppName = &v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyDataApps) SetMseAppId(v string) *DescribeSwimmingLaneResponseBodyDataApps {
	s.MseAppId = &v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyDataApps) SetMseAppName(v string) *DescribeSwimmingLaneResponseBodyDataApps {
	s.MseAppName = &v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyDataApps) SetMseNamespaceId(v string) *DescribeSwimmingLaneResponseBodyDataApps {
	s.MseNamespaceId = &v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyDataApps) Validate() error {
	return dara.Validate(s)
}

type DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRule struct {
	// The logical connector between conditions.
	//
	// example:
	//
	// AND
	ConditionJoiner *string `json:"ConditionJoiner,omitempty" xml:"ConditionJoiner,omitempty"`
	// The matching condition.
	Conditions []*DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// Whether to enable proportional grayscale.
	//
	// example:
	//
	// true
	IndependentPercentageEnable *bool `json:"IndependentPercentageEnable,omitempty" xml:"IndependentPercentageEnable,omitempty"`
	// The proportion of path traffic.
	//
	// example:
	//
	// 100
	Percentage *int32 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	// The traffic configuration.
	PercentageByRoute map[string]*int32 `json:"PercentageByRoute,omitempty" xml:"PercentageByRoute,omitempty"`
	// The list of route IDs.
	RouteIds []*int64 `json:"RouteIds,omitempty" xml:"RouteIds,omitempty" type:"Repeated"`
	// The detailed configuration of the routing rule.
	Routes []*DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutes `json:"Routes,omitempty" xml:"Routes,omitempty" type:"Repeated"`
}

func (s DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRule) GoString() string {
	return s.String()
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRule) GetConditionJoiner() *string {
	return s.ConditionJoiner
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRule) GetConditions() []*DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleConditions {
	return s.Conditions
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRule) GetIndependentPercentageEnable() *bool {
	return s.IndependentPercentageEnable
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRule) GetPercentage() *int32 {
	return s.Percentage
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRule) GetPercentageByRoute() map[string]*int32 {
	return s.PercentageByRoute
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRule) GetRouteIds() []*int64 {
	return s.RouteIds
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRule) GetRoutes() []*DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutes {
	return s.Routes
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRule) SetConditionJoiner(v string) *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRule {
	s.ConditionJoiner = &v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRule) SetConditions(v []*DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleConditions) *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRule {
	s.Conditions = v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRule) SetIndependentPercentageEnable(v bool) *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRule {
	s.IndependentPercentageEnable = &v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRule) SetPercentage(v int32) *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRule {
	s.Percentage = &v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRule) SetPercentageByRoute(v map[string]*int32) *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRule {
	s.PercentageByRoute = v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRule) SetRouteIds(v []*int64) *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRule {
	s.RouteIds = v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRule) SetRoutes(v []*DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutes) *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRule {
	s.Routes = v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRule) Validate() error {
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

type DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleConditions struct {
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
	// The type of the parameter.
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

func (s DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleConditions) String() string {
	return dara.Prettify(s)
}

func (s DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleConditions) GoString() string {
	return s.String()
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleConditions) GetCondition() *string {
	return s.Condition
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleConditions) GetName() *string {
	return s.Name
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleConditions) GetType() *string {
	return s.Type
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleConditions) GetValue() *string {
	return s.Value
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleConditions) SetCondition(v string) *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleConditions {
	s.Condition = &v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleConditions) SetName(v string) *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleConditions {
	s.Name = &v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleConditions) SetType(v string) *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleConditions {
	s.Type = &v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleConditions) SetValue(v string) *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleConditions {
	s.Value = &v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleConditions) Validate() error {
	return dara.Validate(s)
}

type DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutes struct {
	// The route ID.
	//
	// example:
	//
	// 9504
	RouteId *int64 `json:"RouteId,omitempty" xml:"RouteId,omitempty"`
	// The name of the route.
	//
	// example:
	//
	// demo
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
	// The routing rule.
	RoutePredicate *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicate `json:"RoutePredicate,omitempty" xml:"RoutePredicate,omitempty" type:"Struct"`
}

func (s DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutes) String() string {
	return dara.Prettify(s)
}

func (s DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutes) GoString() string {
	return s.String()
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutes) GetRouteId() *int64 {
	return s.RouteId
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutes) GetRouteName() *string {
	return s.RouteName
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutes) GetRoutePredicate() *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicate {
	return s.RoutePredicate
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutes) SetRouteId(v int64) *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutes {
	s.RouteId = &v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutes) SetRouteName(v string) *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutes {
	s.RouteName = &v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutes) SetRoutePredicate(v *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicate) *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutes {
	s.RoutePredicate = v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutes) Validate() error {
	if s.RoutePredicate != nil {
		if err := s.RoutePredicate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicate struct {
	// The path matching rule.
	PathPredicate *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicatePathPredicate `json:"PathPredicate,omitempty" xml:"PathPredicate,omitempty" type:"Struct"`
}

func (s DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicate) String() string {
	return dara.Prettify(s)
}

func (s DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicate) GoString() string {
	return s.String()
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicate) GetPathPredicate() *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicatePathPredicate {
	return s.PathPredicate
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicate) SetPathPredicate(v *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicatePathPredicate) *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicate {
	s.PathPredicate = v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicate) Validate() error {
	if s.PathPredicate != nil {
		if err := s.PathPredicate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicatePathPredicate struct {
	// The path.
	//
	// example:
	//
	// /Path
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The matching rule.
	//
	// example:
	//
	// Header
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicatePathPredicate) String() string {
	return dara.Prettify(s)
}

func (s DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicatePathPredicate) GoString() string {
	return s.String()
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicatePathPredicate) GetPath() *string {
	return s.Path
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicatePathPredicate) GetType() *string {
	return s.Type
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicatePathPredicate) SetPath(v string) *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicatePathPredicate {
	s.Path = &v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicatePathPredicate) SetType(v string) *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicatePathPredicate {
	s.Type = &v
	return s
}

func (s *DescribeSwimmingLaneResponseBodyDataMseGatewayEntryRuleRoutesRoutePredicatePathPredicate) Validate() error {
	return dara.Validate(s)
}
