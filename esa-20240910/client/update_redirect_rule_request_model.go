// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRedirectRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *UpdateRedirectRuleRequest
	GetConfigId() *int64
	SetReserveQueryString(v string) *UpdateRedirectRuleRequest
	GetReserveQueryString() *string
	SetRule(v string) *UpdateRedirectRuleRequest
	GetRule() *string
	SetRuleEnable(v string) *UpdateRedirectRuleRequest
	GetRuleEnable() *string
	SetRuleName(v string) *UpdateRedirectRuleRequest
	GetRuleName() *string
	SetSequence(v int32) *UpdateRedirectRuleRequest
	GetSequence() *int32
	SetSiteId(v int64) *UpdateRedirectRuleRequest
	GetSiteId() *int64
	SetStatusCode(v string) *UpdateRedirectRuleRequest
	GetStatusCode() *string
	SetTargetUrl(v string) *UpdateRedirectRuleRequest
	GetTargetUrl() *string
	SetType(v string) *UpdateRedirectRuleRequest
	GetType() *string
}

type UpdateRedirectRuleRequest struct {
	// The ID of the configuration. To get this ID, call the [ListRedirectRules](https://help.aliyun.com/document_detail/2867474.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3528160969****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Specifies whether to preserve the query string of the original request. Valid values:
	//
	// - `on`: Preserves the query string.
	//
	// - `off`: Does not preserve the query string.
	//
	// example:
	//
	// on
	ReserveQueryString *string `json:"ReserveQueryString,omitempty" xml:"ReserveQueryString,omitempty"`
	// The conditional expression that matches user requests. This parameter is not required for a global configuration. Two scenarios are supported:
	//
	// - To match all requests, set the value to `true`.
	//
	// - To match specific requests, use a custom expression, such as `(http.host eq "video.example.com")`.
	//
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether to enable the rule. This parameter is not required for a global configuration. Valid values:
	//
	// - `on`: Enables the rule.
	//
	// - `off`: Disables the rule.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The name of the rule. This parameter is not required for a global configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The execution order of the rule. A lower value indicates a higher priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The ID of the site. To get this ID, call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The response status code that the edge node returns to the client for the redirect. Valid values:
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
	// The URL to which requests are redirected.
	//
	// example:
	//
	// http://www.exapmle.com/index.html
	TargetUrl *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
	// The redirection type. Valid values:
	//
	// - `static`: Static mode.
	//
	// - `dynamic`: Dynamic mode.
	//
	// example:
	//
	// static
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateRedirectRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRedirectRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateRedirectRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *UpdateRedirectRuleRequest) GetReserveQueryString() *string {
	return s.ReserveQueryString
}

func (s *UpdateRedirectRuleRequest) GetRule() *string {
	return s.Rule
}

func (s *UpdateRedirectRuleRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *UpdateRedirectRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateRedirectRuleRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *UpdateRedirectRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateRedirectRuleRequest) GetStatusCode() *string {
	return s.StatusCode
}

func (s *UpdateRedirectRuleRequest) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *UpdateRedirectRuleRequest) GetType() *string {
	return s.Type
}

func (s *UpdateRedirectRuleRequest) SetConfigId(v int64) *UpdateRedirectRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateRedirectRuleRequest) SetReserveQueryString(v string) *UpdateRedirectRuleRequest {
	s.ReserveQueryString = &v
	return s
}

func (s *UpdateRedirectRuleRequest) SetRule(v string) *UpdateRedirectRuleRequest {
	s.Rule = &v
	return s
}

func (s *UpdateRedirectRuleRequest) SetRuleEnable(v string) *UpdateRedirectRuleRequest {
	s.RuleEnable = &v
	return s
}

func (s *UpdateRedirectRuleRequest) SetRuleName(v string) *UpdateRedirectRuleRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateRedirectRuleRequest) SetSequence(v int32) *UpdateRedirectRuleRequest {
	s.Sequence = &v
	return s
}

func (s *UpdateRedirectRuleRequest) SetSiteId(v int64) *UpdateRedirectRuleRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateRedirectRuleRequest) SetStatusCode(v string) *UpdateRedirectRuleRequest {
	s.StatusCode = &v
	return s
}

func (s *UpdateRedirectRuleRequest) SetTargetUrl(v string) *UpdateRedirectRuleRequest {
	s.TargetUrl = &v
	return s
}

func (s *UpdateRedirectRuleRequest) SetType(v string) *UpdateRedirectRuleRequest {
	s.Type = &v
	return s
}

func (s *UpdateRedirectRuleRequest) Validate() error {
	return dara.Validate(s)
}
