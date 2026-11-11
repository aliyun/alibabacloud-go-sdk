// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateErrorPagesRedirectRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *UpdateErrorPagesRedirectRuleShrinkRequest
	GetConfigId() *int64
	SetErrorPagesRedirectShrink(v string) *UpdateErrorPagesRedirectRuleShrinkRequest
	GetErrorPagesRedirectShrink() *string
	SetRule(v string) *UpdateErrorPagesRedirectRuleShrinkRequest
	GetRule() *string
	SetRuleEnable(v string) *UpdateErrorPagesRedirectRuleShrinkRequest
	GetRuleEnable() *string
	SetRuleName(v string) *UpdateErrorPagesRedirectRuleShrinkRequest
	GetRuleName() *string
	SetSequence(v int32) *UpdateErrorPagesRedirectRuleShrinkRequest
	GetSequence() *int32
	SetSiteId(v int64) *UpdateErrorPagesRedirectRuleShrinkRequest
	GetSiteId() *int64
}

type UpdateErrorPagesRedirectRuleShrinkRequest struct {
	// The configuration ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The error code redirect configuration.
	ErrorPagesRedirectShrink *string `json:"ErrorPagesRedirect,omitempty" xml:"ErrorPagesRedirect,omitempty"`
	// The content of the rule, which uses a conditional expression to match user requests. You do not need to specify this parameter when you add a global configuration. This parameter applies to the following scenarios:
	//
	// - Match all incoming requests: Set the value to true.
	//
	// - Match specific requests: Set the value to a custom expression, such as (http.host eq \\"video.example.com\\").
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether to enable the rule. You do not need to specify this parameter when you add a global configuration. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The name of the rule. You do not need to specify this parameter when you add a global configuration.
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
	// The ID of the site. You can call the [ListSites](~~ListSites~~) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5407498413****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateErrorPagesRedirectRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateErrorPagesRedirectRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateErrorPagesRedirectRuleShrinkRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *UpdateErrorPagesRedirectRuleShrinkRequest) GetErrorPagesRedirectShrink() *string {
	return s.ErrorPagesRedirectShrink
}

func (s *UpdateErrorPagesRedirectRuleShrinkRequest) GetRule() *string {
	return s.Rule
}

func (s *UpdateErrorPagesRedirectRuleShrinkRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *UpdateErrorPagesRedirectRuleShrinkRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateErrorPagesRedirectRuleShrinkRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *UpdateErrorPagesRedirectRuleShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateErrorPagesRedirectRuleShrinkRequest) SetConfigId(v int64) *UpdateErrorPagesRedirectRuleShrinkRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateErrorPagesRedirectRuleShrinkRequest) SetErrorPagesRedirectShrink(v string) *UpdateErrorPagesRedirectRuleShrinkRequest {
	s.ErrorPagesRedirectShrink = &v
	return s
}

func (s *UpdateErrorPagesRedirectRuleShrinkRequest) SetRule(v string) *UpdateErrorPagesRedirectRuleShrinkRequest {
	s.Rule = &v
	return s
}

func (s *UpdateErrorPagesRedirectRuleShrinkRequest) SetRuleEnable(v string) *UpdateErrorPagesRedirectRuleShrinkRequest {
	s.RuleEnable = &v
	return s
}

func (s *UpdateErrorPagesRedirectRuleShrinkRequest) SetRuleName(v string) *UpdateErrorPagesRedirectRuleShrinkRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateErrorPagesRedirectRuleShrinkRequest) SetSequence(v int32) *UpdateErrorPagesRedirectRuleShrinkRequest {
	s.Sequence = &v
	return s
}

func (s *UpdateErrorPagesRedirectRuleShrinkRequest) SetSiteId(v int64) *UpdateErrorPagesRedirectRuleShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateErrorPagesRedirectRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
