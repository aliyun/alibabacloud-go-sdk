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
	// Preserve query string. Value range:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// on
	ReserveQueryString *string `json:"ReserveQueryString,omitempty" xml:"ReserveQueryString,omitempty"`
	// Rule content, using conditional expressions to match user requests. This parameter is not required when adding a global configuration. There are two usage scenarios:
	//
	// - To match all incoming requests: Set the value to true
	//
	// - To match specific requests: Set the value to a custom expression, e.g., (http.host eq \\"video.example.com\\")
	//
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Rule switch. This parameter is not required when adding a global configuration. Value range:
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
	Sequence *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// Version number of the site configuration. For sites with version management enabled, you can use this parameter to specify the version of the site for which the configuration will take effect. The default is version 0.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// Response status code used by the node to respond to the client with the redirect address. Value range:
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
	// Target URL after redirection.
	//
	// This parameter is required.
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
