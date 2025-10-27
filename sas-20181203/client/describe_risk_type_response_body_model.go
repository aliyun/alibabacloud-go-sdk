// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRiskTypeResponseBody
	GetRequestId() *string
	SetRiskTypes(v []*DescribeRiskTypeResponseBodyRiskTypes) *DescribeRiskTypeResponseBody
	GetRiskTypes() []*DescribeRiskTypeResponseBodyRiskTypes
}

type DescribeRiskTypeResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// F22037B5-FCE4-5178-A9E7-71798E1F9270
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the information about baseline types.
	RiskTypes []*DescribeRiskTypeResponseBodyRiskTypes `json:"RiskTypes,omitempty" xml:"RiskTypes,omitempty" type:"Repeated"`
}

func (s DescribeRiskTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRiskTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRiskTypeResponseBody) GetRiskTypes() []*DescribeRiskTypeResponseBodyRiskTypes {
	return s.RiskTypes
}

func (s *DescribeRiskTypeResponseBody) SetRequestId(v string) *DescribeRiskTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRiskTypeResponseBody) SetRiskTypes(v []*DescribeRiskTypeResponseBodyRiskTypes) *DescribeRiskTypeResponseBody {
	s.RiskTypes = v
	return s
}

func (s *DescribeRiskTypeResponseBody) Validate() error {
	if s.RiskTypes != nil {
		for _, item := range s.RiskTypes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRiskTypeResponseBodyRiskTypes struct {
	// The alias of the baseline type.
	//
	// example:
	//
	// Redis unauthorized access high exploit vulnerability risk
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The baseline type flag of the current user version. Valid values:
	//
	// - **true**: Have access
	//
	// - **false**: No permissions
	//
	// example:
	//
	// true
	AuthFlag *bool `json:"AuthFlag,omitempty" xml:"AuthFlag,omitempty"`
	// An array that consists of the information about baseline subtypes.
	SubTypes []*DescribeRiskTypeResponseBodyRiskTypesSubTypes `json:"SubTypes,omitempty" xml:"SubTypes,omitempty" type:"Repeated"`
	// The name of the baseline type.
	//
	// example:
	//
	// hc_exploit
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s DescribeRiskTypeResponseBodyRiskTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskTypeResponseBodyRiskTypes) GoString() string {
	return s.String()
}

func (s *DescribeRiskTypeResponseBodyRiskTypes) GetAlias() *string {
	return s.Alias
}

func (s *DescribeRiskTypeResponseBodyRiskTypes) GetAuthFlag() *bool {
	return s.AuthFlag
}

func (s *DescribeRiskTypeResponseBodyRiskTypes) GetSubTypes() []*DescribeRiskTypeResponseBodyRiskTypesSubTypes {
	return s.SubTypes
}

func (s *DescribeRiskTypeResponseBodyRiskTypes) GetTypeName() *string {
	return s.TypeName
}

func (s *DescribeRiskTypeResponseBodyRiskTypes) SetAlias(v string) *DescribeRiskTypeResponseBodyRiskTypes {
	s.Alias = &v
	return s
}

func (s *DescribeRiskTypeResponseBodyRiskTypes) SetAuthFlag(v bool) *DescribeRiskTypeResponseBodyRiskTypes {
	s.AuthFlag = &v
	return s
}

func (s *DescribeRiskTypeResponseBodyRiskTypes) SetSubTypes(v []*DescribeRiskTypeResponseBodyRiskTypesSubTypes) *DescribeRiskTypeResponseBodyRiskTypes {
	s.SubTypes = v
	return s
}

func (s *DescribeRiskTypeResponseBodyRiskTypes) SetTypeName(v string) *DescribeRiskTypeResponseBodyRiskTypes {
	s.TypeName = &v
	return s
}

func (s *DescribeRiskTypeResponseBodyRiskTypes) Validate() error {
	if s.SubTypes != nil {
		for _, item := range s.SubTypes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRiskTypeResponseBodyRiskTypesSubTypes struct {
	// The alias of the baseline subtype.
	//
	// example:
	//
	// Redis unauthorized access high exploit vulnerability risk
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The baseline subtype permission flag of the current user version. Valid values:
	//
	// - **true**: Have access
	//
	// - **false**: No permissions
	//
	// example:
	//
	// true
	AuthFlag *bool `json:"AuthFlag,omitempty" xml:"AuthFlag,omitempty"`
	// An array that consists of the check details about the baseline subtype.
	CheckDetails []*DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetails `json:"CheckDetails,omitempty" xml:"CheckDetails,omitempty" type:"Repeated"`
	// The operating system type of the server. Valid values:
	//
	// - **windows**
	//
	// - **linux**
	//
	// example:
	//
	// linux
	SupportedOs *string `json:"SupportedOs,omitempty" xml:"SupportedOs,omitempty"`
	// The name of the baseline subtype.
	//
	// example:
	//
	// hc_exploit_redis
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s DescribeRiskTypeResponseBodyRiskTypesSubTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskTypeResponseBodyRiskTypesSubTypes) GoString() string {
	return s.String()
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypes) GetAlias() *string {
	return s.Alias
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypes) GetAuthFlag() *bool {
	return s.AuthFlag
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypes) GetCheckDetails() []*DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetails {
	return s.CheckDetails
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypes) GetSupportedOs() *string {
	return s.SupportedOs
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypes) GetTypeName() *string {
	return s.TypeName
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypes) SetAlias(v string) *DescribeRiskTypeResponseBodyRiskTypesSubTypes {
	s.Alias = &v
	return s
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypes) SetAuthFlag(v bool) *DescribeRiskTypeResponseBodyRiskTypesSubTypes {
	s.AuthFlag = &v
	return s
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypes) SetCheckDetails(v []*DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetails) *DescribeRiskTypeResponseBodyRiskTypesSubTypes {
	s.CheckDetails = v
	return s
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypes) SetSupportedOs(v string) *DescribeRiskTypeResponseBodyRiskTypesSubTypes {
	s.SupportedOs = &v
	return s
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypes) SetTypeName(v string) *DescribeRiskTypeResponseBodyRiskTypesSubTypes {
	s.TypeName = &v
	return s
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypes) Validate() error {
	if s.CheckDetails != nil {
		for _, item := range s.CheckDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetails struct {
	// The description of the baseline.
	//
	// example:
	//
	// Set password expiration time, force regular modification of password, reduce password leakage and guess risk.Use non-password login (e.g. key pair) please ignore this item.
	CheckDesc *string `json:"CheckDesc,omitempty" xml:"CheckDesc,omitempty"`
	// The ID of the baseline.
	//
	// example:
	//
	// 1299
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The baseline.
	//
	// example:
	//
	// Ensure password expiration period is set.
	CheckItem *string `json:"CheckItem,omitempty" xml:"CheckItem,omitempty"`
	// An array that consists of the rule details about the baseline.
	Rules []*DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetails) GoString() string {
	return s.String()
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetails) GetCheckDesc() *string {
	return s.CheckDesc
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetails) GetCheckId() *int64 {
	return s.CheckId
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetails) GetCheckItem() *string {
	return s.CheckItem
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetails) GetRules() []*DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRules {
	return s.Rules
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetails) SetCheckDesc(v string) *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetails {
	s.CheckDesc = &v
	return s
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetails) SetCheckId(v int64) *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetails {
	s.CheckId = &v
	return s
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetails) SetCheckItem(v string) *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetails {
	s.CheckItem = &v
	return s
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetails) SetRules(v []*DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRules) *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetails {
	s.Rules = v
	return s
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetails) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRules struct {
	// Indicates whether the baseline can be edited. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 1
	Optional *int32 `json:"Optional,omitempty" xml:"Optional,omitempty"`
	// An array that consists of the parameters in the rule for the baseline.
	ParamList []*DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRulesParamList `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
	// The description of the rule for the baseline.
	//
	// example:
	//
	// Please customize the password expiration time detection standard as
	RuleDesc *string `json:"RuleDesc,omitempty" xml:"RuleDesc,omitempty"`
	// The ID of the rule for the baseline.
	//
	// example:
	//
	// audit.audit_policy.auditpolicychange.cus
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRules) GoString() string {
	return s.String()
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRules) GetOptional() *int32 {
	return s.Optional
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRules) GetParamList() []*DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRulesParamList {
	return s.ParamList
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRules) GetRuleDesc() *string {
	return s.RuleDesc
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRules) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRules) SetOptional(v int32) *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRules {
	s.Optional = &v
	return s
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRules) SetParamList(v []*DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRulesParamList) *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRules {
	s.ParamList = v
	return s
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRules) SetRuleDesc(v string) *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRules {
	s.RuleDesc = &v
	return s
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRules) SetRuleId(v string) *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRules {
	s.RuleId = &v
	return s
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRules) Validate() error {
	if s.ParamList != nil {
		for _, item := range s.ParamList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRulesParamList struct {
	// If the value of paramType is 1, this parameter is empty. If the value of paramType is 2, this parameter provides the options that can be selected for paramType.
	//
	// example:
	//
	// 0,1,2,3
	EnumValue *string `json:"EnumValue,omitempty" xml:"EnumValue,omitempty"`
	// The maximum value of the parameter.
	//
	// example:
	//
	// 999
	MaxValue *int32 `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	// The minimum value of the parameter.
	//
	// example:
	//
	// 1
	MinValue *int32 `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	// The default value of the parameter.
	//
	// example:
	//
	// 7
	ParamDefaultValue *string `json:"ParamDefaultValue,omitempty" xml:"ParamDefaultValue,omitempty"`
	// The description of the parameter.
	//
	// example:
	//
	// The setting value is 0 means no definition, 1 means success, 2 means failure, 3 means success and failure
	ParamDesc *string `json:"ParamDesc,omitempty" xml:"ParamDesc,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// range_val
	ParamName *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
	// The configuration type of the parameter. Valid values:
	//
	// 	- **1**: input
	//
	// 	- **2**: selection
	//
	// example:
	//
	// 1
	ParamType *int32 `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
}

func (s DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRulesParamList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRulesParamList) GoString() string {
	return s.String()
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRulesParamList) GetEnumValue() *string {
	return s.EnumValue
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRulesParamList) GetMaxValue() *int32 {
	return s.MaxValue
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRulesParamList) GetMinValue() *int32 {
	return s.MinValue
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRulesParamList) GetParamDefaultValue() *string {
	return s.ParamDefaultValue
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRulesParamList) GetParamDesc() *string {
	return s.ParamDesc
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRulesParamList) GetParamName() *string {
	return s.ParamName
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRulesParamList) GetParamType() *int32 {
	return s.ParamType
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRulesParamList) SetEnumValue(v string) *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRulesParamList {
	s.EnumValue = &v
	return s
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRulesParamList) SetMaxValue(v int32) *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRulesParamList {
	s.MaxValue = &v
	return s
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRulesParamList) SetMinValue(v int32) *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRulesParamList {
	s.MinValue = &v
	return s
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRulesParamList) SetParamDefaultValue(v string) *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRulesParamList {
	s.ParamDefaultValue = &v
	return s
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRulesParamList) SetParamDesc(v string) *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRulesParamList {
	s.ParamDesc = &v
	return s
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRulesParamList) SetParamName(v string) *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRulesParamList {
	s.ParamName = &v
	return s
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRulesParamList) SetParamType(v int32) *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRulesParamList {
	s.ParamType = &v
	return s
}

func (s *DescribeRiskTypeResponseBodyRiskTypesSubTypesCheckDetailsRulesParamList) Validate() error {
	return dara.Validate(s)
}
