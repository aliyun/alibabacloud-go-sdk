// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpRequestHeaderModificationRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRequestHeaderModificationShrink(v string) *CreateHttpRequestHeaderModificationRuleShrinkRequest
	GetRequestHeaderModificationShrink() *string
	SetRule(v string) *CreateHttpRequestHeaderModificationRuleShrinkRequest
	GetRule() *string
	SetRuleEnable(v string) *CreateHttpRequestHeaderModificationRuleShrinkRequest
	GetRuleEnable() *string
	SetRuleName(v string) *CreateHttpRequestHeaderModificationRuleShrinkRequest
	GetRuleName() *string
	SetSequence(v int32) *CreateHttpRequestHeaderModificationRuleShrinkRequest
	GetSequence() *int32
	SetSiteId(v int64) *CreateHttpRequestHeaderModificationRuleShrinkRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *CreateHttpRequestHeaderModificationRuleShrinkRequest
	GetSiteVersion() *int32
}

type CreateHttpRequestHeaderModificationRuleShrinkRequest struct {
	// The request header modifications. Three operation types are supported: add, delete, and modify.
	//
	// This parameter is required.
	RequestHeaderModificationShrink *string `json:"RequestHeaderModification,omitempty" xml:"RequestHeaderModification,omitempty"`
	// The rule content, which uses conditional expressions to match user requests. This parameter is not required when you add a global configuration. Two scenarios are supported:
	//
	// - Match all incoming requests: Set the value to true.
	//
	// - Match specified requests: Set the value to a custom expression, such as (http.host eq \\"video.example.com\\").
	//
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// The rule switch. This parameter is not required when you add a global configuration. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name. This parameter is not required when you add a global configuration.
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
	// The site ID. You can call the [ListSites](~~ListSites~~) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5407498413****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version number of the site configuration. For sites with version management enabled, you can use this parameter to specify the site version on which the configuration takes effect. The default value is 0.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s CreateHttpRequestHeaderModificationRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpRequestHeaderModificationRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateHttpRequestHeaderModificationRuleShrinkRequest) GetRequestHeaderModificationShrink() *string {
	return s.RequestHeaderModificationShrink
}

func (s *CreateHttpRequestHeaderModificationRuleShrinkRequest) GetRule() *string {
	return s.Rule
}

func (s *CreateHttpRequestHeaderModificationRuleShrinkRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *CreateHttpRequestHeaderModificationRuleShrinkRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateHttpRequestHeaderModificationRuleShrinkRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *CreateHttpRequestHeaderModificationRuleShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateHttpRequestHeaderModificationRuleShrinkRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *CreateHttpRequestHeaderModificationRuleShrinkRequest) SetRequestHeaderModificationShrink(v string) *CreateHttpRequestHeaderModificationRuleShrinkRequest {
	s.RequestHeaderModificationShrink = &v
	return s
}

func (s *CreateHttpRequestHeaderModificationRuleShrinkRequest) SetRule(v string) *CreateHttpRequestHeaderModificationRuleShrinkRequest {
	s.Rule = &v
	return s
}

func (s *CreateHttpRequestHeaderModificationRuleShrinkRequest) SetRuleEnable(v string) *CreateHttpRequestHeaderModificationRuleShrinkRequest {
	s.RuleEnable = &v
	return s
}

func (s *CreateHttpRequestHeaderModificationRuleShrinkRequest) SetRuleName(v string) *CreateHttpRequestHeaderModificationRuleShrinkRequest {
	s.RuleName = &v
	return s
}

func (s *CreateHttpRequestHeaderModificationRuleShrinkRequest) SetSequence(v int32) *CreateHttpRequestHeaderModificationRuleShrinkRequest {
	s.Sequence = &v
	return s
}

func (s *CreateHttpRequestHeaderModificationRuleShrinkRequest) SetSiteId(v int64) *CreateHttpRequestHeaderModificationRuleShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *CreateHttpRequestHeaderModificationRuleShrinkRequest) SetSiteVersion(v int32) *CreateHttpRequestHeaderModificationRuleShrinkRequest {
	s.SiteVersion = &v
	return s
}

func (s *CreateHttpRequestHeaderModificationRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
