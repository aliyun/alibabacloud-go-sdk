// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRedirectRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*ListRedirectRulesResponseBodyConfigs) *ListRedirectRulesResponseBody
	GetConfigs() []*ListRedirectRulesResponseBodyConfigs
	SetPageNumber(v int32) *ListRedirectRulesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRedirectRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListRedirectRulesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListRedirectRulesResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListRedirectRulesResponseBody
	GetTotalPage() *int32
}

type ListRedirectRulesResponseBody struct {
	// The redirect configuration list.
	Configs []*ListRedirectRulesResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// The current page number, which is the same as the PageNumber request parameter.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EDBD3EB3-97DA-5465-AEF5-8DCA5DC5E395
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 1
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListRedirectRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRedirectRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRedirectRulesResponseBody) GetConfigs() []*ListRedirectRulesResponseBodyConfigs {
	return s.Configs
}

func (s *ListRedirectRulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRedirectRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRedirectRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRedirectRulesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListRedirectRulesResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListRedirectRulesResponseBody) SetConfigs(v []*ListRedirectRulesResponseBodyConfigs) *ListRedirectRulesResponseBody {
	s.Configs = v
	return s
}

func (s *ListRedirectRulesResponseBody) SetPageNumber(v int32) *ListRedirectRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRedirectRulesResponseBody) SetPageSize(v int32) *ListRedirectRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRedirectRulesResponseBody) SetRequestId(v string) *ListRedirectRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRedirectRulesResponseBody) SetTotalCount(v int32) *ListRedirectRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRedirectRulesResponseBody) SetTotalPage(v int32) *ListRedirectRulesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListRedirectRulesResponseBody) Validate() error {
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

type ListRedirectRulesResponseBodyConfigs struct {
	// The configuration ID.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The configuration type. Valid values:
	//
	// - global: global configuration.
	//
	// - rule: rule configuration.
	//
	// example:
	//
	// rule
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// Specifies whether to preserve the query string. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	ReserveQueryString *string `json:"ReserveQueryString,omitempty" xml:"ReserveQueryString,omitempty"`
	// The rule content, which uses a conditional expression to match user requests. You do not need to set this parameter when adding a global configuration. Two scenarios are supported:
	//
	// - Match all incoming requests: set the value to true.
	//
	// - Match specified requests: set the value to a custom expression, for example, (http.host eq \\"video.example.com\\").
	//
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// The rule switch. You do not need to set this parameter when adding a global configuration. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
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
	// The rule execution order. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The version number of the site configuration. For sites with version management enabled, you can use this parameter to specify the site version for which the configuration takes effect. Default value: 0.
	//
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// The HTTP status code used when the node responds to the client with the redirect address. Valid values:
	//
	// - 301
	//
	// - 302
	//
	// - 303
	//
	// - 307
	//
	// - 308.
	//
	// example:
	//
	// 301
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// The target URL after the redirect.
	//
	// example:
	//
	// http://www.exapmle.com/index.html
	TargetUrl *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
	// The redirect type. Valid values:
	//
	// - static: static pattern.
	//
	// - dynamic: dynamic schema.
	//
	// example:
	//
	// static
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListRedirectRulesResponseBodyConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListRedirectRulesResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *ListRedirectRulesResponseBodyConfigs) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListRedirectRulesResponseBodyConfigs) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListRedirectRulesResponseBodyConfigs) GetReserveQueryString() *string {
	return s.ReserveQueryString
}

func (s *ListRedirectRulesResponseBodyConfigs) GetRule() *string {
	return s.Rule
}

func (s *ListRedirectRulesResponseBodyConfigs) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListRedirectRulesResponseBodyConfigs) GetRuleName() *string {
	return s.RuleName
}

func (s *ListRedirectRulesResponseBodyConfigs) GetSequence() *int32 {
	return s.Sequence
}

func (s *ListRedirectRulesResponseBodyConfigs) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListRedirectRulesResponseBodyConfigs) GetStatusCode() *string {
	return s.StatusCode
}

func (s *ListRedirectRulesResponseBodyConfigs) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *ListRedirectRulesResponseBodyConfigs) GetType() *string {
	return s.Type
}

func (s *ListRedirectRulesResponseBodyConfigs) SetConfigId(v int64) *ListRedirectRulesResponseBodyConfigs {
	s.ConfigId = &v
	return s
}

func (s *ListRedirectRulesResponseBodyConfigs) SetConfigType(v string) *ListRedirectRulesResponseBodyConfigs {
	s.ConfigType = &v
	return s
}

func (s *ListRedirectRulesResponseBodyConfigs) SetReserveQueryString(v string) *ListRedirectRulesResponseBodyConfigs {
	s.ReserveQueryString = &v
	return s
}

func (s *ListRedirectRulesResponseBodyConfigs) SetRule(v string) *ListRedirectRulesResponseBodyConfigs {
	s.Rule = &v
	return s
}

func (s *ListRedirectRulesResponseBodyConfigs) SetRuleEnable(v string) *ListRedirectRulesResponseBodyConfigs {
	s.RuleEnable = &v
	return s
}

func (s *ListRedirectRulesResponseBodyConfigs) SetRuleName(v string) *ListRedirectRulesResponseBodyConfigs {
	s.RuleName = &v
	return s
}

func (s *ListRedirectRulesResponseBodyConfigs) SetSequence(v int32) *ListRedirectRulesResponseBodyConfigs {
	s.Sequence = &v
	return s
}

func (s *ListRedirectRulesResponseBodyConfigs) SetSiteVersion(v int32) *ListRedirectRulesResponseBodyConfigs {
	s.SiteVersion = &v
	return s
}

func (s *ListRedirectRulesResponseBodyConfigs) SetStatusCode(v string) *ListRedirectRulesResponseBodyConfigs {
	s.StatusCode = &v
	return s
}

func (s *ListRedirectRulesResponseBodyConfigs) SetTargetUrl(v string) *ListRedirectRulesResponseBodyConfigs {
	s.TargetUrl = &v
	return s
}

func (s *ListRedirectRulesResponseBodyConfigs) SetType(v string) *ListRedirectRulesResponseBodyConfigs {
	s.Type = &v
	return s
}

func (s *ListRedirectRulesResponseBodyConfigs) Validate() error {
	return dara.Validate(s)
}
