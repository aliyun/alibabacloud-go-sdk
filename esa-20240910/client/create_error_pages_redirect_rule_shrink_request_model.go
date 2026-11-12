// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateErrorPagesRedirectRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetErrorPagesRedirectShrink(v string) *CreateErrorPagesRedirectRuleShrinkRequest
	GetErrorPagesRedirectShrink() *string
	SetRule(v string) *CreateErrorPagesRedirectRuleShrinkRequest
	GetRule() *string
	SetRuleEnable(v string) *CreateErrorPagesRedirectRuleShrinkRequest
	GetRuleEnable() *string
	SetRuleName(v string) *CreateErrorPagesRedirectRuleShrinkRequest
	GetRuleName() *string
	SetSequence(v int32) *CreateErrorPagesRedirectRuleShrinkRequest
	GetSequence() *int32
	SetSiteId(v int64) *CreateErrorPagesRedirectRuleShrinkRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *CreateErrorPagesRedirectRuleShrinkRequest
	GetSiteVersion() *int32
}

type CreateErrorPagesRedirectRuleShrinkRequest struct {
	// The error code redirect configuration.
	//
	// This parameter is required.
	ErrorPagesRedirectShrink *string `json:"ErrorPagesRedirect,omitempty" xml:"ErrorPagesRedirect,omitempty"`
	// The rule content, which uses conditional expressions to match user requests. You do not need to specify this parameter when adding a global configuration. Two scenarios are supported:
	//
	// - Match all incoming requests: Set the value to true.
	//
	// - Match specific requests: Set the value to a custom expression, such as (http.host eq "video.example.com").
	//
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether to enable the rule. You do not need to specify this parameter when adding a global configuration. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The name of the rule. You do not need to specify this parameter when adding a global configuration.
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
	// The ID of the site. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 478016908379824
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version number of the site configuration. For sites with configuration version management enabled, you can use this parameter to specify the site version to which the configuration applies. Default value: 0.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s CreateErrorPagesRedirectRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateErrorPagesRedirectRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateErrorPagesRedirectRuleShrinkRequest) GetErrorPagesRedirectShrink() *string {
	return s.ErrorPagesRedirectShrink
}

func (s *CreateErrorPagesRedirectRuleShrinkRequest) GetRule() *string {
	return s.Rule
}

func (s *CreateErrorPagesRedirectRuleShrinkRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *CreateErrorPagesRedirectRuleShrinkRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateErrorPagesRedirectRuleShrinkRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *CreateErrorPagesRedirectRuleShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateErrorPagesRedirectRuleShrinkRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *CreateErrorPagesRedirectRuleShrinkRequest) SetErrorPagesRedirectShrink(v string) *CreateErrorPagesRedirectRuleShrinkRequest {
	s.ErrorPagesRedirectShrink = &v
	return s
}

func (s *CreateErrorPagesRedirectRuleShrinkRequest) SetRule(v string) *CreateErrorPagesRedirectRuleShrinkRequest {
	s.Rule = &v
	return s
}

func (s *CreateErrorPagesRedirectRuleShrinkRequest) SetRuleEnable(v string) *CreateErrorPagesRedirectRuleShrinkRequest {
	s.RuleEnable = &v
	return s
}

func (s *CreateErrorPagesRedirectRuleShrinkRequest) SetRuleName(v string) *CreateErrorPagesRedirectRuleShrinkRequest {
	s.RuleName = &v
	return s
}

func (s *CreateErrorPagesRedirectRuleShrinkRequest) SetSequence(v int32) *CreateErrorPagesRedirectRuleShrinkRequest {
	s.Sequence = &v
	return s
}

func (s *CreateErrorPagesRedirectRuleShrinkRequest) SetSiteId(v int64) *CreateErrorPagesRedirectRuleShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *CreateErrorPagesRedirectRuleShrinkRequest) SetSiteVersion(v int32) *CreateErrorPagesRedirectRuleShrinkRequest {
	s.SiteVersion = &v
	return s
}

func (s *CreateErrorPagesRedirectRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
