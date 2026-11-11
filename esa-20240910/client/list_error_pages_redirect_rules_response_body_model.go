// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListErrorPagesRedirectRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*ListErrorPagesRedirectRulesResponseBodyConfigs) *ListErrorPagesRedirectRulesResponseBody
	GetConfigs() []*ListErrorPagesRedirectRulesResponseBodyConfigs
	SetPageNumber(v int32) *ListErrorPagesRedirectRulesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListErrorPagesRedirectRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListErrorPagesRedirectRulesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListErrorPagesRedirectRulesResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListErrorPagesRedirectRulesResponseBody
	GetTotalPage() *int32
}

type ListErrorPagesRedirectRulesResponseBody struct {
	// The error code redirection configurations.
	Configs []*ListErrorPagesRedirectRulesResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// The page number. Default value: 1.
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
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
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

func (s ListErrorPagesRedirectRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListErrorPagesRedirectRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListErrorPagesRedirectRulesResponseBody) GetConfigs() []*ListErrorPagesRedirectRulesResponseBodyConfigs {
	return s.Configs
}

func (s *ListErrorPagesRedirectRulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListErrorPagesRedirectRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListErrorPagesRedirectRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListErrorPagesRedirectRulesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListErrorPagesRedirectRulesResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListErrorPagesRedirectRulesResponseBody) SetConfigs(v []*ListErrorPagesRedirectRulesResponseBodyConfigs) *ListErrorPagesRedirectRulesResponseBody {
	s.Configs = v
	return s
}

func (s *ListErrorPagesRedirectRulesResponseBody) SetPageNumber(v int32) *ListErrorPagesRedirectRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListErrorPagesRedirectRulesResponseBody) SetPageSize(v int32) *ListErrorPagesRedirectRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListErrorPagesRedirectRulesResponseBody) SetRequestId(v string) *ListErrorPagesRedirectRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListErrorPagesRedirectRulesResponseBody) SetTotalCount(v int32) *ListErrorPagesRedirectRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListErrorPagesRedirectRulesResponseBody) SetTotalPage(v int32) *ListErrorPagesRedirectRulesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListErrorPagesRedirectRulesResponseBody) Validate() error {
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

type ListErrorPagesRedirectRulesResponseBodyConfigs struct {
	// The configuration ID.
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
	// The error code redirection.
	ErrorPagesRedirect []*ListErrorPagesRedirectRulesResponseBodyConfigsErrorPagesRedirect `json:"ErrorPagesRedirect,omitempty" xml:"ErrorPagesRedirect,omitempty" type:"Repeated"`
	// The rule content. A conditional expression is used to match user requests. You do not need to set this parameter when adding a global configuration. Two scenarios are supported:
	//
	// - Match all incoming requests: Set the value to true.
	//
	// - Match specific requests: Set the value to a custom expression, such as (http.host eq "video.example.com").
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
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
	// The execution sequence of the rule. A smaller value indicates a higher priority.
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

func (s ListErrorPagesRedirectRulesResponseBodyConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListErrorPagesRedirectRulesResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *ListErrorPagesRedirectRulesResponseBodyConfigs) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListErrorPagesRedirectRulesResponseBodyConfigs) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListErrorPagesRedirectRulesResponseBodyConfigs) GetErrorPagesRedirect() []*ListErrorPagesRedirectRulesResponseBodyConfigsErrorPagesRedirect {
	return s.ErrorPagesRedirect
}

func (s *ListErrorPagesRedirectRulesResponseBodyConfigs) GetRule() *string {
	return s.Rule
}

func (s *ListErrorPagesRedirectRulesResponseBodyConfigs) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListErrorPagesRedirectRulesResponseBodyConfigs) GetRuleName() *string {
	return s.RuleName
}

func (s *ListErrorPagesRedirectRulesResponseBodyConfigs) GetSequence() *int32 {
	return s.Sequence
}

func (s *ListErrorPagesRedirectRulesResponseBodyConfigs) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListErrorPagesRedirectRulesResponseBodyConfigs) SetConfigId(v int64) *ListErrorPagesRedirectRulesResponseBodyConfigs {
	s.ConfigId = &v
	return s
}

func (s *ListErrorPagesRedirectRulesResponseBodyConfigs) SetConfigType(v string) *ListErrorPagesRedirectRulesResponseBodyConfigs {
	s.ConfigType = &v
	return s
}

func (s *ListErrorPagesRedirectRulesResponseBodyConfigs) SetErrorPagesRedirect(v []*ListErrorPagesRedirectRulesResponseBodyConfigsErrorPagesRedirect) *ListErrorPagesRedirectRulesResponseBodyConfigs {
	s.ErrorPagesRedirect = v
	return s
}

func (s *ListErrorPagesRedirectRulesResponseBodyConfigs) SetRule(v string) *ListErrorPagesRedirectRulesResponseBodyConfigs {
	s.Rule = &v
	return s
}

func (s *ListErrorPagesRedirectRulesResponseBodyConfigs) SetRuleEnable(v string) *ListErrorPagesRedirectRulesResponseBodyConfigs {
	s.RuleEnable = &v
	return s
}

func (s *ListErrorPagesRedirectRulesResponseBodyConfigs) SetRuleName(v string) *ListErrorPagesRedirectRulesResponseBodyConfigs {
	s.RuleName = &v
	return s
}

func (s *ListErrorPagesRedirectRulesResponseBodyConfigs) SetSequence(v int32) *ListErrorPagesRedirectRulesResponseBodyConfigs {
	s.Sequence = &v
	return s
}

func (s *ListErrorPagesRedirectRulesResponseBodyConfigs) SetSiteVersion(v int32) *ListErrorPagesRedirectRulesResponseBodyConfigs {
	s.SiteVersion = &v
	return s
}

func (s *ListErrorPagesRedirectRulesResponseBodyConfigs) Validate() error {
	if s.ErrorPagesRedirect != nil {
		for _, item := range s.ErrorPagesRedirect {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListErrorPagesRedirectRulesResponseBodyConfigsErrorPagesRedirect struct {
	// The response status code used by the node when responding to the client with a redirect address. Valid values:
	//
	// - 400
	//
	// - 403
	//
	// - 404
	//
	// - 405
	//
	// - 414
	//
	// - 416
	//
	// - 500
	//
	// - 501
	//
	// - 502
	//
	// - 503
	//
	// - 504
	//
	// example:
	//
	// 400
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// The target URL after redirection.
	//
	// example:
	//
	// http://example.com/test
	TargetURL *string `json:"TargetURL,omitempty" xml:"TargetURL,omitempty"`
}

func (s ListErrorPagesRedirectRulesResponseBodyConfigsErrorPagesRedirect) String() string {
	return dara.Prettify(s)
}

func (s ListErrorPagesRedirectRulesResponseBodyConfigsErrorPagesRedirect) GoString() string {
	return s.String()
}

func (s *ListErrorPagesRedirectRulesResponseBodyConfigsErrorPagesRedirect) GetStatusCode() *string {
	return s.StatusCode
}

func (s *ListErrorPagesRedirectRulesResponseBodyConfigsErrorPagesRedirect) GetTargetURL() *string {
	return s.TargetURL
}

func (s *ListErrorPagesRedirectRulesResponseBodyConfigsErrorPagesRedirect) SetStatusCode(v string) *ListErrorPagesRedirectRulesResponseBodyConfigsErrorPagesRedirect {
	s.StatusCode = &v
	return s
}

func (s *ListErrorPagesRedirectRulesResponseBodyConfigsErrorPagesRedirect) SetTargetURL(v string) *ListErrorPagesRedirectRulesResponseBodyConfigsErrorPagesRedirect {
	s.TargetURL = &v
	return s
}

func (s *ListErrorPagesRedirectRulesResponseBodyConfigsErrorPagesRedirect) Validate() error {
	return dara.Validate(s)
}
