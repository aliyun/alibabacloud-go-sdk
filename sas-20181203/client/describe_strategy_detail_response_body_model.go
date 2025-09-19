// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStrategyDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeStrategyDetailResponseBody
	GetRequestId() *string
	SetStrategy(v *DescribeStrategyDetailResponseBodyStrategy) *DescribeStrategyDetailResponseBody
	GetStrategy() *DescribeStrategyDetailResponseBodyStrategy
}

type DescribeStrategyDetailResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// C5B28F65-9245-5DC1-B3CF-5F2756A756A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the baseline check policy.
	Strategy *DescribeStrategyDetailResponseBodyStrategy `json:"Strategy,omitempty" xml:"Strategy,omitempty" type:"Struct"`
}

func (s DescribeStrategyDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeStrategyDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStrategyDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeStrategyDetailResponseBody) GetStrategy() *DescribeStrategyDetailResponseBodyStrategy {
	return s.Strategy
}

func (s *DescribeStrategyDetailResponseBody) SetRequestId(v string) *DescribeStrategyDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStrategyDetailResponseBody) SetStrategy(v *DescribeStrategyDetailResponseBodyStrategy) *DescribeStrategyDetailResponseBody {
	s.Strategy = v
	return s
}

func (s *DescribeStrategyDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeStrategyDetailResponseBodyStrategy struct {
	// The type of the baseline check policy that you want to query. Valid values:
	//
	// 	- **common**: standard baseline check policy
	//
	// 	- **custom**: custom baseline check policy
	//
	// example:
	//
	// common
	CustomType *string `json:"CustomType,omitempty" xml:"CustomType,omitempty"`
	// The check interval of the policy.
	//
	// example:
	//
	// 3
	CycleDays *int32 `json:"CycleDays,omitempty" xml:"CycleDays,omitempty"`
	// The time period during which the check starts. Valid values:
	//
	// 	- **0**: 00:00 to 06:00
	//
	// 	- **6**: 06:00 to 12:00
	//
	// 	- **12**: 12:00 to 18:00
	//
	// 	- **18**: 18:00 to 24:00
	//
	// example:
	//
	// 0
	CycleStartTime *int32 `json:"CycleStartTime,omitempty" xml:"CycleStartTime,omitempty"`
	// The end time of the check. Specify the time in the HH:mm:ss format.
	//
	// example:
	//
	// 03:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the baseline check policy.
	//
	// example:
	//
	// 123
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the baseline check policy.
	//
	// example:
	//
	// TestStrategy
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The subtype of the baselines.
	//
	// > You can call the [DescribeRiskType](~~DescribeRiskType~~) operation to query the subtypes of baselines.
	//
	// example:
	//
	// hc_nginx_linux,tomcat7,hc_mysql_ali,hc_docker
	RiskSubTypeName *string `json:"RiskSubTypeName,omitempty" xml:"RiskSubTypeName,omitempty"`
	// The information about the whitelist of risk items.
	RiskTypeWhiteListQueryResultList []*DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultList `json:"RiskTypeWhiteListQueryResultList,omitempty" xml:"RiskTypeWhiteListQueryResultList,omitempty" type:"Repeated"`
	// The start time of the check. Specify the time in the HH:mm:ss format.
	//
	// example:
	//
	// 02:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The method that is used to apply the baseline check policy. Valid values:
	//
	// 	- **groupId**: asset groups
	//
	// 	- **uuid**: assets
	//
	// example:
	//
	// groupId
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The type of the baseline check policy. Valid values:
	//
	// 	- **1**: standard policies
	//
	// 	- **2**: custom policies
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeStrategyDetailResponseBodyStrategy) String() string {
	return dara.Prettify(s)
}

func (s DescribeStrategyDetailResponseBodyStrategy) GoString() string {
	return s.String()
}

func (s *DescribeStrategyDetailResponseBodyStrategy) GetCustomType() *string {
	return s.CustomType
}

func (s *DescribeStrategyDetailResponseBodyStrategy) GetCycleDays() *int32 {
	return s.CycleDays
}

func (s *DescribeStrategyDetailResponseBodyStrategy) GetCycleStartTime() *int32 {
	return s.CycleStartTime
}

func (s *DescribeStrategyDetailResponseBodyStrategy) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeStrategyDetailResponseBodyStrategy) GetId() *int32 {
	return s.Id
}

func (s *DescribeStrategyDetailResponseBodyStrategy) GetName() *string {
	return s.Name
}

func (s *DescribeStrategyDetailResponseBodyStrategy) GetRiskSubTypeName() *string {
	return s.RiskSubTypeName
}

func (s *DescribeStrategyDetailResponseBodyStrategy) GetRiskTypeWhiteListQueryResultList() []*DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultList {
	return s.RiskTypeWhiteListQueryResultList
}

func (s *DescribeStrategyDetailResponseBodyStrategy) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeStrategyDetailResponseBodyStrategy) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeStrategyDetailResponseBodyStrategy) GetType() *int32 {
	return s.Type
}

func (s *DescribeStrategyDetailResponseBodyStrategy) SetCustomType(v string) *DescribeStrategyDetailResponseBodyStrategy {
	s.CustomType = &v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategy) SetCycleDays(v int32) *DescribeStrategyDetailResponseBodyStrategy {
	s.CycleDays = &v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategy) SetCycleStartTime(v int32) *DescribeStrategyDetailResponseBodyStrategy {
	s.CycleStartTime = &v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategy) SetEndTime(v string) *DescribeStrategyDetailResponseBodyStrategy {
	s.EndTime = &v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategy) SetId(v int32) *DescribeStrategyDetailResponseBodyStrategy {
	s.Id = &v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategy) SetName(v string) *DescribeStrategyDetailResponseBodyStrategy {
	s.Name = &v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategy) SetRiskSubTypeName(v string) *DescribeStrategyDetailResponseBodyStrategy {
	s.RiskSubTypeName = &v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategy) SetRiskTypeWhiteListQueryResultList(v []*DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultList) *DescribeStrategyDetailResponseBodyStrategy {
	s.RiskTypeWhiteListQueryResultList = v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategy) SetStartTime(v string) *DescribeStrategyDetailResponseBodyStrategy {
	s.StartTime = &v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategy) SetTargetType(v string) *DescribeStrategyDetailResponseBodyStrategy {
	s.TargetType = &v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategy) SetType(v int32) *DescribeStrategyDetailResponseBodyStrategy {
	s.Type = &v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategy) Validate() error {
	return dara.Validate(s)
}

type DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultList struct {
	// The alias of the check item.
	//
	// example:
	//
	// Unauthorized Access
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// Indicates whether the check item is selected. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	On *bool `json:"On,omitempty" xml:"On,omitempty"`
	// The information about sub-check items.
	SubTypes []*DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypes `json:"SubTypes,omitempty" xml:"SubTypes,omitempty" type:"Repeated"`
	// The name of the check item.
	//
	// example:
	//
	// hc_exploit
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultList) String() string {
	return dara.Prettify(s)
}

func (s DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultList) GoString() string {
	return s.String()
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultList) GetAlias() *string {
	return s.Alias
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultList) GetOn() *bool {
	return s.On
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultList) GetSubTypes() []*DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypes {
	return s.SubTypes
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultList) GetTypeName() *string {
	return s.TypeName
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultList) SetAlias(v string) *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultList {
	s.Alias = &v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultList) SetOn(v bool) *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultList {
	s.On = &v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultList) SetSubTypes(v []*DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypes) *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultList {
	s.SubTypes = v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultList) SetTypeName(v string) *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultList {
	s.TypeName = &v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultList) Validate() error {
	return dara.Validate(s)
}

type DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypes struct {
	// The alias of the check item.
	//
	// example:
	//
	// Redis unauthorized access high exploit vulnerability risk
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The details of custom check items.
	CheckDetails []*DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetails `json:"CheckDetails,omitempty" xml:"CheckDetails,omitempty" type:"Repeated"`
	// Indicates whether the sub-check item is selected. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	On *bool `json:"On,omitempty" xml:"On,omitempty"`
	// The operating system type of the server. Valid values:
	//
	// 	- **windows**
	//
	// 	- **linux**
	//
	// example:
	//
	// windows
	SupportedOs *string `json:"SupportedOs,omitempty" xml:"SupportedOs,omitempty"`
	// The type of the sub-check item.
	//
	// example:
	//
	// hc_exploit_redis
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypes) GoString() string {
	return s.String()
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypes) GetAlias() *string {
	return s.Alias
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypes) GetCheckDetails() []*DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetails {
	return s.CheckDetails
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypes) GetOn() *bool {
	return s.On
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypes) GetSupportedOs() *string {
	return s.SupportedOs
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypes) GetTypeName() *string {
	return s.TypeName
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypes) SetAlias(v string) *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypes {
	s.Alias = &v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypes) SetCheckDetails(v []*DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetails) *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypes {
	s.CheckDetails = v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypes) SetOn(v bool) *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypes {
	s.On = &v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypes) SetSupportedOs(v string) *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypes {
	s.SupportedOs = &v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypes) SetTypeName(v string) *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypes {
	s.TypeName = &v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypes) Validate() error {
	return dara.Validate(s)
}

type DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetails struct {
	// The description of the check item.
	//
	// example:
	//
	// Set password expiration time, force regular modification of password, reduce password leakage and guess risk.Use non-password login (e.g. key pair) please ignore this item.
	CheckDesc *string `json:"CheckDesc,omitempty" xml:"CheckDesc,omitempty"`
	// The ID of the check item.
	//
	// example:
	//
	// 206
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The check item.
	//
	// example:
	//
	// Ensure password expiration period is set.
	CheckItem *string `json:"CheckItem,omitempty" xml:"CheckItem,omitempty"`
	// The details of rules.
	Rules []*DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetails) GoString() string {
	return s.String()
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetails) GetCheckDesc() *string {
	return s.CheckDesc
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetails) GetCheckId() *int64 {
	return s.CheckId
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetails) GetCheckItem() *string {
	return s.CheckItem
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetails) GetRules() []*DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRules {
	return s.Rules
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetails) SetCheckDesc(v string) *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetails {
	s.CheckDesc = &v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetails) SetCheckId(v int64) *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetails {
	s.CheckId = &v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetails) SetCheckItem(v string) *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetails {
	s.CheckItem = &v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetails) SetRules(v []*DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRules) *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetails {
	s.Rules = v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetails) Validate() error {
	return dara.Validate(s)
}

type DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRules struct {
	// The default value of the rule.
	//
	// example:
	//
	// 2
	DefaultValue *int32 `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// Indicates whether the rule can be selected. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	Optional *int32 `json:"Optional,omitempty" xml:"Optional,omitempty"`
	// The rule parameters.
	ParamList []*DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRulesParamList `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
	// The description of the rule.
	//
	// example:
	//
	// Please customize the password expiration time detection standard as
	RuleDesc *string `json:"RuleDesc,omitempty" xml:"RuleDesc,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// login_unlock_deny_pam_faillock.must.cus
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRules) GoString() string {
	return s.String()
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRules) GetDefaultValue() *int32 {
	return s.DefaultValue
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRules) GetOptional() *int32 {
	return s.Optional
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRules) GetParamList() []*DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRulesParamList {
	return s.ParamList
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRules) GetRuleDesc() *string {
	return s.RuleDesc
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRules) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRules) SetDefaultValue(v int32) *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRules {
	s.DefaultValue = &v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRules) SetOptional(v int32) *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRules {
	s.Optional = &v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRules) SetParamList(v []*DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRulesParamList) *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRules {
	s.ParamList = v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRules) SetRuleDesc(v string) *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRules {
	s.RuleDesc = &v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRules) SetRuleId(v string) *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRules {
	s.RuleId = &v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRules) Validate() error {
	return dara.Validate(s)
}

type DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRulesParamList struct {
	// The options that can be selected for the rule parameter if the value of ParamType is set to 2.
	//
	// example:
	//
	// 0,1,2,3
	EnumValue *string `json:"EnumValue,omitempty" xml:"EnumValue,omitempty"`
	// The maximum value of the rule parameter.
	//
	// example:
	//
	// 999
	MaxValue *int32 `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	// The minimum value of the rule parameter.
	//
	// example:
	//
	// 1
	MinValue *int32 `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	// The default value of the rule parameter.
	//
	// example:
	//
	// 7
	ParamDefaultValue *string `json:"ParamDefaultValue,omitempty" xml:"ParamDefaultValue,omitempty"`
	// The description of the rule parameter.
	//
	// example:
	//
	// The setting value is 0 means no definition, 1 means success, 2 means failure, 3 means success and failure
	ParamDesc *string `json:"ParamDesc,omitempty" xml:"ParamDesc,omitempty"`
	// The name of the rule parameter.
	//
	// example:
	//
	// range_val
	ParamName *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
	// The type of the rule parameter. Valid values:
	//
	// 	- **1**: input
	//
	// 	- **2**: selection
	//
	// example:
	//
	// 1
	ParamType *int32 `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// The configured value of the rule parameter.
	//
	// example:
	//
	// 7
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRulesParamList) String() string {
	return dara.Prettify(s)
}

func (s DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRulesParamList) GoString() string {
	return s.String()
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRulesParamList) GetEnumValue() *string {
	return s.EnumValue
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRulesParamList) GetMaxValue() *int32 {
	return s.MaxValue
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRulesParamList) GetMinValue() *int32 {
	return s.MinValue
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRulesParamList) GetParamDefaultValue() *string {
	return s.ParamDefaultValue
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRulesParamList) GetParamDesc() *string {
	return s.ParamDesc
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRulesParamList) GetParamName() *string {
	return s.ParamName
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRulesParamList) GetParamType() *int32 {
	return s.ParamType
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRulesParamList) GetValue() *string {
	return s.Value
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRulesParamList) SetEnumValue(v string) *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRulesParamList {
	s.EnumValue = &v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRulesParamList) SetMaxValue(v int32) *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRulesParamList {
	s.MaxValue = &v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRulesParamList) SetMinValue(v int32) *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRulesParamList {
	s.MinValue = &v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRulesParamList) SetParamDefaultValue(v string) *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRulesParamList {
	s.ParamDefaultValue = &v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRulesParamList) SetParamDesc(v string) *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRulesParamList {
	s.ParamDesc = &v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRulesParamList) SetParamName(v string) *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRulesParamList {
	s.ParamName = &v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRulesParamList) SetParamType(v int32) *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRulesParamList {
	s.ParamType = &v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRulesParamList) SetValue(v string) *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRulesParamList {
	s.Value = &v
	return s
}

func (s *DescribeStrategyDetailResponseBodyStrategyRiskTypeWhiteListQueryResultListSubTypesCheckDetailsRulesParamList) Validate() error {
	return dara.Validate(s)
}
