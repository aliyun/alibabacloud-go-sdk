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
	// The list of custom response code configurations.
	Configs []*ListCustomResponseCodeRulesResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// Page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size. Default value: 500. Valid values: 1 to 500.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of entries.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Total number of pages.
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
	// Configuration ID.
	//
	// example:
	//
	// 395386449776640
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The configuration type. Valid values:
	//
	// - global: global configuration.
	//
	// - rule: rule configuration.
	//
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// The response page.
	//
	// example:
	//
	// 0
	PageId *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	// The response code.
	//
	// example:
	//
	// 200
	ReturnCode *string `json:"ReturnCode,omitempty" xml:"ReturnCode,omitempty"`
	// The rule content, which uses conditional expressions to match user requests. You do not need to set this parameter when adding a global configuration. Two usage scenarios are supported:
	//
	// - Match all incoming requests: set the value to true.
	//
	// - Match specific requests: set the value to a custom expression, for example: (http.host eq \\"video.example.com\\").
	//
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// The rule switch. You do not need to set this parameter when adding a global configuration. Valid values:
	//
	// - on: enables the rule.
	//
	// - off: disables the rule.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name. You do not need to set this parameter when adding a global configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The execution order of the rule. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The version number of the site configuration. For sites with configuration version management enabled, you can use this parameter to specify the site version for which the configuration takes effect. Default value: 0.
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
