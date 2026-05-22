// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomResponseCodeRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*ListCustomResponseCodeRulesResponseBodyConfigs) *ListCustomResponseCodeRulesResponseBody
	GetConfigs() []*ListCustomResponseCodeRulesResponseBodyConfigs
	SetPageNumber(v int32) *ListCustomResponseCodeRulesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCustomResponseCodeRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListCustomResponseCodeRulesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListCustomResponseCodeRulesResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListCustomResponseCodeRulesResponseBody
	GetTotalPage() *int32
}

type ListCustomResponseCodeRulesResponseBody struct {
	// Modifies the list of response code configurations.
	Configs []*ListCustomResponseCodeRulesResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// The number of the returned page. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 500. Valid values: 1 to 500.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 1
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListCustomResponseCodeRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomResponseCodeRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomResponseCodeRulesResponseBody) GetConfigs() []*ListCustomResponseCodeRulesResponseBodyConfigs {
	return s.Configs
}

func (s *ListCustomResponseCodeRulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomResponseCodeRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomResponseCodeRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomResponseCodeRulesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCustomResponseCodeRulesResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListCustomResponseCodeRulesResponseBody) SetConfigs(v []*ListCustomResponseCodeRulesResponseBodyConfigs) *ListCustomResponseCodeRulesResponseBody {
	s.Configs = v
	return s
}

func (s *ListCustomResponseCodeRulesResponseBody) SetPageNumber(v int32) *ListCustomResponseCodeRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCustomResponseCodeRulesResponseBody) SetPageSize(v int32) *ListCustomResponseCodeRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCustomResponseCodeRulesResponseBody) SetRequestId(v string) *ListCustomResponseCodeRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomResponseCodeRulesResponseBody) SetTotalCount(v int32) *ListCustomResponseCodeRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCustomResponseCodeRulesResponseBody) SetTotalPage(v int32) *ListCustomResponseCodeRulesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListCustomResponseCodeRulesResponseBody) Validate() error {
	if s.Configs != nil {
		for _, item := range s.Configs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCustomResponseCodeRulesResponseBodyConfigs struct {
	// The configuration ID.
	//
	// example:
	//
	// 395386449776640
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The configuration type. Valid values:
	//
	// 	- global: global configuration.
	//
	// 	- rule: rule configuration.
	//
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// Response page.
	//
	// example:
	//
	// 0
	PageId *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	// Response code
	//
	// example:
	//
	// 200
	ReturnCode *string `json:"ReturnCode,omitempty" xml:"ReturnCode,omitempty"`
	// The content of the rule. A conditional expression is used to match a user request. You do not need to set this parameter when you add global configuration. Use cases:
	//
	// 	- true: Match all incoming requests.
	//
	// 	- Set the value to a custom expression, for example, (http.host eq "video.example.com"): Match the specified request.
	//
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether to enable the rule. Valid values: You do not need to set this parameter when you add global configuration. Valid values:
	//
	// 	- on
	//
	// 	- off
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name. You do not need to set this parameter when you add global configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The order in which the rule is executed. A smaller value gives priority to the rule.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The version number of the website configurations. You can use this parameter to specify a version of your website to apply the feature settings. By default, version 0 is used.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListCustomResponseCodeRulesResponseBodyConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListCustomResponseCodeRulesResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) GetPageId() *string {
	return s.PageId
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) GetReturnCode() *string {
	return s.ReturnCode
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) GetRule() *string {
	return s.Rule
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) GetRuleName() *string {
	return s.RuleName
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) GetSequence() *int32 {
	return s.Sequence
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) SetConfigId(v int64) *ListCustomResponseCodeRulesResponseBodyConfigs {
	s.ConfigId = &v
	return s
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) SetConfigType(v string) *ListCustomResponseCodeRulesResponseBodyConfigs {
	s.ConfigType = &v
	return s
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) SetPageId(v string) *ListCustomResponseCodeRulesResponseBodyConfigs {
	s.PageId = &v
	return s
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) SetReturnCode(v string) *ListCustomResponseCodeRulesResponseBodyConfigs {
	s.ReturnCode = &v
	return s
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) SetRule(v string) *ListCustomResponseCodeRulesResponseBodyConfigs {
	s.Rule = &v
	return s
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) SetRuleEnable(v string) *ListCustomResponseCodeRulesResponseBodyConfigs {
	s.RuleEnable = &v
	return s
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) SetRuleName(v string) *ListCustomResponseCodeRulesResponseBodyConfigs {
	s.RuleName = &v
	return s
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) SetSequence(v int32) *ListCustomResponseCodeRulesResponseBodyConfigs {
	s.Sequence = &v
	return s
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) SetSiteVersion(v int32) *ListCustomResponseCodeRulesResponseBodyConfigs {
	s.SiteVersion = &v
	return s
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) Validate() error {
	return dara.Validate(s)
}
