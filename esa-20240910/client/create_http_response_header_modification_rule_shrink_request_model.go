// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpResponseHeaderModificationRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResponseHeaderModificationShrink(v string) *CreateHttpResponseHeaderModificationRuleShrinkRequest
	GetResponseHeaderModificationShrink() *string
	SetRule(v string) *CreateHttpResponseHeaderModificationRuleShrinkRequest
	GetRule() *string
	SetRuleEnable(v string) *CreateHttpResponseHeaderModificationRuleShrinkRequest
	GetRuleEnable() *string
	SetRuleName(v string) *CreateHttpResponseHeaderModificationRuleShrinkRequest
	GetRuleName() *string
	SetSequence(v int32) *CreateHttpResponseHeaderModificationRuleShrinkRequest
	GetSequence() *int32
	SetSiteId(v int64) *CreateHttpResponseHeaderModificationRuleShrinkRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *CreateHttpResponseHeaderModificationRuleShrinkRequest
	GetSiteVersion() *int32
}

type CreateHttpResponseHeaderModificationRuleShrinkRequest struct {
	// The response header modifications. Three operation types are supported: add, delete, and modify.
	//
	// This parameter is required.
	ResponseHeaderModificationShrink *string `json:"ResponseHeaderModification,omitempty" xml:"ResponseHeaderModification,omitempty"`
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
	// The site ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to query the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456******
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version number of the site configuration. For sites with version management enabled, you can use this parameter to specify the site version on which the configuration takes effect. The default value is 0.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s CreateHttpResponseHeaderModificationRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpResponseHeaderModificationRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateHttpResponseHeaderModificationRuleShrinkRequest) GetResponseHeaderModificationShrink() *string {
	return s.ResponseHeaderModificationShrink
}

func (s *CreateHttpResponseHeaderModificationRuleShrinkRequest) GetRule() *string {
	return s.Rule
}

func (s *CreateHttpResponseHeaderModificationRuleShrinkRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *CreateHttpResponseHeaderModificationRuleShrinkRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateHttpResponseHeaderModificationRuleShrinkRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *CreateHttpResponseHeaderModificationRuleShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateHttpResponseHeaderModificationRuleShrinkRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *CreateHttpResponseHeaderModificationRuleShrinkRequest) SetResponseHeaderModificationShrink(v string) *CreateHttpResponseHeaderModificationRuleShrinkRequest {
	s.ResponseHeaderModificationShrink = &v
	return s
}

func (s *CreateHttpResponseHeaderModificationRuleShrinkRequest) SetRule(v string) *CreateHttpResponseHeaderModificationRuleShrinkRequest {
	s.Rule = &v
	return s
}

func (s *CreateHttpResponseHeaderModificationRuleShrinkRequest) SetRuleEnable(v string) *CreateHttpResponseHeaderModificationRuleShrinkRequest {
	s.RuleEnable = &v
	return s
}

func (s *CreateHttpResponseHeaderModificationRuleShrinkRequest) SetRuleName(v string) *CreateHttpResponseHeaderModificationRuleShrinkRequest {
	s.RuleName = &v
	return s
}

func (s *CreateHttpResponseHeaderModificationRuleShrinkRequest) SetSequence(v int32) *CreateHttpResponseHeaderModificationRuleShrinkRequest {
	s.Sequence = &v
	return s
}

func (s *CreateHttpResponseHeaderModificationRuleShrinkRequest) SetSiteId(v int64) *CreateHttpResponseHeaderModificationRuleShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *CreateHttpResponseHeaderModificationRuleShrinkRequest) SetSiteVersion(v int32) *CreateHttpResponseHeaderModificationRuleShrinkRequest {
	s.SiteVersion = &v
	return s
}

func (s *CreateHttpResponseHeaderModificationRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
