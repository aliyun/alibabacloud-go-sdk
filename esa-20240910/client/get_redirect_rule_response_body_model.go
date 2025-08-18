// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRedirectRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetRedirectRuleResponseBody
	GetConfigId() *int64
	SetConfigType(v string) *GetRedirectRuleResponseBody
	GetConfigType() *string
	SetRequestId(v string) *GetRedirectRuleResponseBody
	GetRequestId() *string
	SetReserveQueryString(v string) *GetRedirectRuleResponseBody
	GetReserveQueryString() *string
	SetRule(v string) *GetRedirectRuleResponseBody
	GetRule() *string
	SetRuleEnable(v string) *GetRedirectRuleResponseBody
	GetRuleEnable() *string
	SetRuleName(v string) *GetRedirectRuleResponseBody
	GetRuleName() *string
	SetSequence(v int32) *GetRedirectRuleResponseBody
	GetSequence() *int32
	SetSiteVersion(v int32) *GetRedirectRuleResponseBody
	GetSiteVersion() *int32
	SetStatusCode(v string) *GetRedirectRuleResponseBody
	GetStatusCode() *string
	SetTargetUrl(v string) *GetRedirectRuleResponseBody
	GetTargetUrl() *string
	SetType(v string) *GetRedirectRuleResponseBody
	GetType() *string
}

type GetRedirectRuleResponseBody struct {
	// Configuration ID.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Configuration type. Possible values:
	//
	// - global: Global configuration.
	//
	// - rule: Rule-based configuration.
	//
	// example:
	//
	// rule
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Preserve query string. Possible values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	ReserveQueryString *string `json:"ReserveQueryString,omitempty" xml:"ReserveQueryString,omitempty"`
	// Rule content, using conditional expressions to match user requests. This parameter is not required when adding a global configuration. There are two usage scenarios:
	//
	// - Match all incoming requests: Set the value to true
	//
	// - Match specific requests: Set the value to a custom expression, e.g., (http.host eq \\"video.example.com\\")
	//
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Rule switch. This parameter is not required when adding a global configuration. Possible values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// Rule name. This parameter is not required when adding a global configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Rule execution order. The smaller the value, the higher the priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The version number of the site configuration. For sites with version management enabled, this parameter can specify the effective version of the site, defaulting to version 0.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// Response status code used by the node to respond to the client with the redirect address. Possible values:
	//
	// - 301
	//
	// - 302
	//
	// - 303
	//
	// - 307
	//
	// - 308
	//
	// example:
	//
	// 301
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// Target URL after redirection.
	//
	// example:
	//
	// http://www.exapmle.com/index.html
	TargetUrl *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
	// Redirect type. Possible values:
	//
	// - static: Static mode.
	//
	// - dynamic: Dynamic mode.
	//
	// example:
	//
	// static
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetRedirectRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRedirectRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetRedirectRuleResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetRedirectRuleResponseBody) GetConfigType() *string {
	return s.ConfigType
}

func (s *GetRedirectRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRedirectRuleResponseBody) GetReserveQueryString() *string {
	return s.ReserveQueryString
}

func (s *GetRedirectRuleResponseBody) GetRule() *string {
	return s.Rule
}

func (s *GetRedirectRuleResponseBody) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *GetRedirectRuleResponseBody) GetRuleName() *string {
	return s.RuleName
}

func (s *GetRedirectRuleResponseBody) GetSequence() *int32 {
	return s.Sequence
}

func (s *GetRedirectRuleResponseBody) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *GetRedirectRuleResponseBody) GetStatusCode() *string {
	return s.StatusCode
}

func (s *GetRedirectRuleResponseBody) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *GetRedirectRuleResponseBody) GetType() *string {
	return s.Type
}

func (s *GetRedirectRuleResponseBody) SetConfigId(v int64) *GetRedirectRuleResponseBody {
	s.ConfigId = &v
	return s
}

func (s *GetRedirectRuleResponseBody) SetConfigType(v string) *GetRedirectRuleResponseBody {
	s.ConfigType = &v
	return s
}

func (s *GetRedirectRuleResponseBody) SetRequestId(v string) *GetRedirectRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRedirectRuleResponseBody) SetReserveQueryString(v string) *GetRedirectRuleResponseBody {
	s.ReserveQueryString = &v
	return s
}

func (s *GetRedirectRuleResponseBody) SetRule(v string) *GetRedirectRuleResponseBody {
	s.Rule = &v
	return s
}

func (s *GetRedirectRuleResponseBody) SetRuleEnable(v string) *GetRedirectRuleResponseBody {
	s.RuleEnable = &v
	return s
}

func (s *GetRedirectRuleResponseBody) SetRuleName(v string) *GetRedirectRuleResponseBody {
	s.RuleName = &v
	return s
}

func (s *GetRedirectRuleResponseBody) SetSequence(v int32) *GetRedirectRuleResponseBody {
	s.Sequence = &v
	return s
}

func (s *GetRedirectRuleResponseBody) SetSiteVersion(v int32) *GetRedirectRuleResponseBody {
	s.SiteVersion = &v
	return s
}

func (s *GetRedirectRuleResponseBody) SetStatusCode(v string) *GetRedirectRuleResponseBody {
	s.StatusCode = &v
	return s
}

func (s *GetRedirectRuleResponseBody) SetTargetUrl(v string) *GetRedirectRuleResponseBody {
	s.TargetUrl = &v
	return s
}

func (s *GetRedirectRuleResponseBody) SetType(v string) *GetRedirectRuleResponseBody {
	s.Type = &v
	return s
}

func (s *GetRedirectRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
