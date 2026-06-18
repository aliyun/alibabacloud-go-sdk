// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRedirectRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReserveQueryString(v string) *CreateRedirectRuleRequest
	GetReserveQueryString() *string
	SetRule(v string) *CreateRedirectRuleRequest
	GetRule() *string
	SetRuleEnable(v string) *CreateRedirectRuleRequest
	GetRuleEnable() *string
	SetRuleName(v string) *CreateRedirectRuleRequest
	GetRuleName() *string
	SetSequence(v int32) *CreateRedirectRuleRequest
	GetSequence() *int32
	SetSiteId(v int64) *CreateRedirectRuleRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *CreateRedirectRuleRequest
	GetSiteVersion() *int32
	SetStatusCode(v string) *CreateRedirectRuleRequest
	GetStatusCode() *string
	SetTargetUrl(v string) *CreateRedirectRuleRequest
	GetTargetUrl() *string
	SetType(v string) *CreateRedirectRuleRequest
	GetType() *string
}

type CreateRedirectRuleRequest struct {
	// Specifies whether to preserve the query string from the original request. Valid values:
	//
	// - `on`: Preserves the query string.
	//
	// - `off`: Discards the query string.
	//
	// This parameter is required.
	//
	// example:
	//
	// on
	ReserveQueryString *string `json:"ReserveQueryString,omitempty" xml:"ReserveQueryString,omitempty"`
	// The rule content, which is a conditional expression used to match user requests. Do not set this parameter when adding a global configuration. The following use cases are supported:
	//
	// - To match all incoming requests, set the value to `true`.
	//
	// - To match specific requests, set the value to a custom expression, such as `(http.host eq "video.example.com")`.
	//
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether to enable the rule. Do not set this parameter when adding a global configuration. Valid values:
	//
	// - `on`: The rule is enabled.
	//
	// - `off`: The rule is disabled.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name. Do not set this parameter when adding a global configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The execution priority of the rule. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The site ID. To get this value, call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// For sites with configuration version management enabled, specify the version to which this configuration applies.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// The status code that the edge node returns to the client for the redirect. Valid values:
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
	// This parameter is required.
	//
	// example:
	//
	// 301
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// The target URL for the redirect.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://www.exapmle.com/index.html
	TargetUrl *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
	// The redirect type. Valid values:
	//
	// - `static`: Static mode.
	//
	// - `dynamic`: Dynamic mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// static
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateRedirectRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRedirectRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateRedirectRuleRequest) GetReserveQueryString() *string {
	return s.ReserveQueryString
}

func (s *CreateRedirectRuleRequest) GetRule() *string {
	return s.Rule
}

func (s *CreateRedirectRuleRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *CreateRedirectRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateRedirectRuleRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *CreateRedirectRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateRedirectRuleRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *CreateRedirectRuleRequest) GetStatusCode() *string {
	return s.StatusCode
}

func (s *CreateRedirectRuleRequest) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *CreateRedirectRuleRequest) GetType() *string {
	return s.Type
}

func (s *CreateRedirectRuleRequest) SetReserveQueryString(v string) *CreateRedirectRuleRequest {
	s.ReserveQueryString = &v
	return s
}

func (s *CreateRedirectRuleRequest) SetRule(v string) *CreateRedirectRuleRequest {
	s.Rule = &v
	return s
}

func (s *CreateRedirectRuleRequest) SetRuleEnable(v string) *CreateRedirectRuleRequest {
	s.RuleEnable = &v
	return s
}

func (s *CreateRedirectRuleRequest) SetRuleName(v string) *CreateRedirectRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateRedirectRuleRequest) SetSequence(v int32) *CreateRedirectRuleRequest {
	s.Sequence = &v
	return s
}

func (s *CreateRedirectRuleRequest) SetSiteId(v int64) *CreateRedirectRuleRequest {
	s.SiteId = &v
	return s
}

func (s *CreateRedirectRuleRequest) SetSiteVersion(v int32) *CreateRedirectRuleRequest {
	s.SiteVersion = &v
	return s
}

func (s *CreateRedirectRuleRequest) SetStatusCode(v string) *CreateRedirectRuleRequest {
	s.StatusCode = &v
	return s
}

func (s *CreateRedirectRuleRequest) SetTargetUrl(v string) *CreateRedirectRuleRequest {
	s.TargetUrl = &v
	return s
}

func (s *CreateRedirectRuleRequest) SetType(v string) *CreateRedirectRuleRequest {
	s.Type = &v
	return s
}

func (s *CreateRedirectRuleRequest) Validate() error {
	return dara.Validate(s)
}
