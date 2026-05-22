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
	// Configuration ID. It can be obtained by calling the [ListRedirectRules](https://help.aliyun.com/document_detail/2867474.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3528160969****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Preserve query string. Value range:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	ReserveQueryString *string `json:"ReserveQueryString,omitempty" xml:"ReserveQueryString,omitempty"`
	// Rule content, using conditional expressions to match user requests. This parameter is not required when adding a global configuration. There are two usage scenarios:
	//
	// - Match all incoming requests: Set the value to true
	//
	// - Match specific requests: Set the value to a custom expression, for example: (http.host eq \\"video.example.com\\")
	//
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Rule switch. This parameter is not required when adding a global configuration. Value range:
	//
	// - on: Enable.
	//
	// - off: Disable.
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
	Sequence *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The response status code used by the node to respond with the redirect address to the client. Value range:
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
	// The target URL after redirection.
	//
	// example:
	//
	// http://www.exapmle.com/index.html
	TargetUrl *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
	// Redirect type. Value range:
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
