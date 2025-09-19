// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCheckFixDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCheckFixDetails(v []*DescribeCheckFixDetailsResponseBodyCheckFixDetails) *DescribeCheckFixDetailsResponseBody
	GetCheckFixDetails() []*DescribeCheckFixDetailsResponseBodyCheckFixDetails
	SetCount(v int32) *DescribeCheckFixDetailsResponseBody
	GetCount() *int32
	SetRequestId(v string) *DescribeCheckFixDetailsResponseBody
	GetRequestId() *string
}

type DescribeCheckFixDetailsResponseBody struct {
	// An array that consists of the parameters.
	CheckFixDetails []*DescribeCheckFixDetailsResponseBodyCheckFixDetails `json:"CheckFixDetails,omitempty" xml:"CheckFixDetails,omitempty" type:"Repeated"`
	// The number of risk items that can be fixed.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 0DBF1E27-98D8-5EC2-9CF3-4A2E26F6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCheckFixDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckFixDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCheckFixDetailsResponseBody) GetCheckFixDetails() []*DescribeCheckFixDetailsResponseBodyCheckFixDetails {
	return s.CheckFixDetails
}

func (s *DescribeCheckFixDetailsResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeCheckFixDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCheckFixDetailsResponseBody) SetCheckFixDetails(v []*DescribeCheckFixDetailsResponseBodyCheckFixDetails) *DescribeCheckFixDetailsResponseBody {
	s.CheckFixDetails = v
	return s
}

func (s *DescribeCheckFixDetailsResponseBody) SetCount(v int32) *DescribeCheckFixDetailsResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeCheckFixDetailsResponseBody) SetRequestId(v string) *DescribeCheckFixDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCheckFixDetailsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCheckFixDetailsResponseBodyCheckFixDetails struct {
	// The detailed description of the risk item.
	//
	// example:
	//
	// Force users not to reuse recently used passwords to reduce the risk of password guessing attacks
	CheckDesc *string `json:"CheckDesc,omitempty" xml:"CheckDesc,omitempty"`
	// The ID of the risk item.
	//
	// example:
	//
	// 58
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The description of the risk item.
	//
	// example:
	//
	// Ensure password reuse is limited
	CheckItem *string `json:"CheckItem,omitempty" xml:"CheckItem,omitempty"`
	// An array consisting of the rules that are supported by the risk item.
	Rules []*DescribeCheckFixDetailsResponseBodyCheckFixDetailsRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeCheckFixDetailsResponseBodyCheckFixDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckFixDetailsResponseBodyCheckFixDetails) GoString() string {
	return s.String()
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetails) GetCheckDesc() *string {
	return s.CheckDesc
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetails) GetCheckId() *int64 {
	return s.CheckId
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetails) GetCheckItem() *string {
	return s.CheckItem
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetails) GetRules() []*DescribeCheckFixDetailsResponseBodyCheckFixDetailsRules {
	return s.Rules
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetails) SetCheckDesc(v string) *DescribeCheckFixDetailsResponseBodyCheckFixDetails {
	s.CheckDesc = &v
	return s
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetails) SetCheckId(v int64) *DescribeCheckFixDetailsResponseBodyCheckFixDetails {
	s.CheckId = &v
	return s
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetails) SetCheckItem(v string) *DescribeCheckFixDetailsResponseBodyCheckFixDetails {
	s.CheckItem = &v
	return s
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetails) SetRules(v []*DescribeCheckFixDetailsResponseBodyCheckFixDetailsRules) *DescribeCheckFixDetailsResponseBodyCheckFixDetails {
	s.Rules = v
	return s
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetails) Validate() error {
	return dara.Validate(s)
}

type DescribeCheckFixDetailsResponseBodyCheckFixDetailsRules struct {
	// The ID of the risk item.
	//
	// example:
	//
	// 58
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The default value of the rule.
	//
	// example:
	//
	// 1
	DefaultValue *int32 `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// Indicates whether the rule is optional. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	Optional *int32 `json:"Optional,omitempty" xml:"Optional,omitempty"`
	// An array that consists of the rule parameters.
	ParamList []*DescribeCheckFixDetailsResponseBodyCheckFixDetailsRulesParamList `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
	// The description of the rule.
	//
	// example:
	//
	// (/etc/system-auth)Force users not to reuse the number of recently used passwords between 5 and 24
	RuleDesc *string `json:"RuleDesc,omitempty" xml:"RuleDesc,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// pwd_reuse.system_auth
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The specified value of the rule parameter.
	//
	// example:
	//
	// 5
	Value *int32 `json:"Value,omitempty" xml:"Value,omitempty"`
	// The name of the variable.
	//
	// example:
	//
	// open
	VarName *string `json:"VarName,omitempty" xml:"VarName,omitempty"`
}

func (s DescribeCheckFixDetailsResponseBodyCheckFixDetailsRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckFixDetailsResponseBodyCheckFixDetailsRules) GoString() string {
	return s.String()
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRules) GetCheckId() *int64 {
	return s.CheckId
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRules) GetDefaultValue() *int32 {
	return s.DefaultValue
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRules) GetOptional() *int32 {
	return s.Optional
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRules) GetParamList() []*DescribeCheckFixDetailsResponseBodyCheckFixDetailsRulesParamList {
	return s.ParamList
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRules) GetRuleDesc() *string {
	return s.RuleDesc
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRules) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRules) GetValue() *int32 {
	return s.Value
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRules) GetVarName() *string {
	return s.VarName
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRules) SetCheckId(v int64) *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRules {
	s.CheckId = &v
	return s
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRules) SetDefaultValue(v int32) *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRules {
	s.DefaultValue = &v
	return s
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRules) SetOptional(v int32) *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRules {
	s.Optional = &v
	return s
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRules) SetParamList(v []*DescribeCheckFixDetailsResponseBodyCheckFixDetailsRulesParamList) *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRules {
	s.ParamList = v
	return s
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRules) SetRuleDesc(v string) *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRules {
	s.RuleDesc = &v
	return s
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRules) SetRuleId(v string) *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRules {
	s.RuleId = &v
	return s
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRules) SetValue(v int32) *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRules {
	s.Value = &v
	return s
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRules) SetVarName(v string) *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRules {
	s.VarName = &v
	return s
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRules) Validate() error {
	return dara.Validate(s)
}

type DescribeCheckFixDetailsResponseBodyCheckFixDetailsRulesParamList struct {
	// The options that can be selected for the rule parameter if the value of the ParamType parameter is 2.
	//
	// example:
	//
	// 0,1,2,3
	EnumValue *string `json:"EnumValue,omitempty" xml:"EnumValue,omitempty"`
	// The maximum value of the rule parameter.
	//
	// example:
	//
	// 24
	MaxValue *int32 `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	// The minimum value of the rule parameter.
	//
	// example:
	//
	// 5
	MinValue *int32 `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	// The default value of the rule parameter.
	//
	// example:
	//
	// 5
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
	// The ID of the rule.
	//
	// example:
	//
	// pwd_reuse.system_auth
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The specified value of the rule parameter.
	//
	// example:
	//
	// 18
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCheckFixDetailsResponseBodyCheckFixDetailsRulesParamList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckFixDetailsResponseBodyCheckFixDetailsRulesParamList) GoString() string {
	return s.String()
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRulesParamList) GetEnumValue() *string {
	return s.EnumValue
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRulesParamList) GetMaxValue() *int32 {
	return s.MaxValue
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRulesParamList) GetMinValue() *int32 {
	return s.MinValue
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRulesParamList) GetParamDefaultValue() *string {
	return s.ParamDefaultValue
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRulesParamList) GetParamDesc() *string {
	return s.ParamDesc
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRulesParamList) GetParamName() *string {
	return s.ParamName
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRulesParamList) GetParamType() *int32 {
	return s.ParamType
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRulesParamList) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRulesParamList) GetValue() *string {
	return s.Value
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRulesParamList) SetEnumValue(v string) *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRulesParamList {
	s.EnumValue = &v
	return s
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRulesParamList) SetMaxValue(v int32) *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRulesParamList {
	s.MaxValue = &v
	return s
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRulesParamList) SetMinValue(v int32) *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRulesParamList {
	s.MinValue = &v
	return s
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRulesParamList) SetParamDefaultValue(v string) *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRulesParamList {
	s.ParamDefaultValue = &v
	return s
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRulesParamList) SetParamDesc(v string) *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRulesParamList {
	s.ParamDesc = &v
	return s
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRulesParamList) SetParamName(v string) *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRulesParamList {
	s.ParamName = &v
	return s
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRulesParamList) SetParamType(v int32) *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRulesParamList {
	s.ParamType = &v
	return s
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRulesParamList) SetRuleId(v string) *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRulesParamList {
	s.RuleId = &v
	return s
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRulesParamList) SetValue(v string) *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRulesParamList {
	s.Value = &v
	return s
}

func (s *DescribeCheckFixDetailsResponseBodyCheckFixDetailsRulesParamList) Validate() error {
	return dara.Validate(s)
}
