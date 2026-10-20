// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpResponseHeaderModificationRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *UpdateHttpResponseHeaderModificationRuleShrinkRequest
	GetConfigId() *int64
	SetResponseHeaderModificationShrink(v string) *UpdateHttpResponseHeaderModificationRuleShrinkRequest
	GetResponseHeaderModificationShrink() *string
	SetRule(v string) *UpdateHttpResponseHeaderModificationRuleShrinkRequest
	GetRule() *string
	SetRuleEnable(v string) *UpdateHttpResponseHeaderModificationRuleShrinkRequest
	GetRuleEnable() *string
	SetRuleName(v string) *UpdateHttpResponseHeaderModificationRuleShrinkRequest
	GetRuleName() *string
	SetSequence(v int32) *UpdateHttpResponseHeaderModificationRuleShrinkRequest
	GetSequence() *int32
	SetSiteId(v int64) *UpdateHttpResponseHeaderModificationRuleShrinkRequest
	GetSiteId() *int64
}

type UpdateHttpResponseHeaderModificationRuleShrinkRequest struct {
	// The configuration ID. You can obtain the ID by calling the [ListHttpResponseHeaderModificationRules](https://help.aliyun.com/document_detail/2867483.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The response header modifications. You can add, delete, or modify response headers.
	ResponseHeaderModificationShrink *string `json:"ResponseHeaderModification,omitempty" xml:"ResponseHeaderModification,omitempty"`
	// The rule content, which uses a conditional expression to match user requests. You do not need to specify this parameter when you add a global configuration. This parameter applies to the following two scenarios:
	//
	// - Match all incoming requests: Set the value to true.
	//
	// - Match specific requests: Set the value to a custom expression, such as (http.host eq "video.example.com").
	//
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// The rule switch. You do not need to specify this parameter when you add a global configuration. Valid values:
	//
	// - on: enabled
	//
	// - off: disabled
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
	// The ID of the site. You can obtain the ID by calling the [ListSites](~~ListSites~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456******
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateHttpResponseHeaderModificationRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpResponseHeaderModificationRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateHttpResponseHeaderModificationRuleShrinkRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *UpdateHttpResponseHeaderModificationRuleShrinkRequest) GetResponseHeaderModificationShrink() *string {
	return s.ResponseHeaderModificationShrink
}

func (s *UpdateHttpResponseHeaderModificationRuleShrinkRequest) GetRule() *string {
	return s.Rule
}

func (s *UpdateHttpResponseHeaderModificationRuleShrinkRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *UpdateHttpResponseHeaderModificationRuleShrinkRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateHttpResponseHeaderModificationRuleShrinkRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *UpdateHttpResponseHeaderModificationRuleShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateHttpResponseHeaderModificationRuleShrinkRequest) SetConfigId(v int64) *UpdateHttpResponseHeaderModificationRuleShrinkRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateHttpResponseHeaderModificationRuleShrinkRequest) SetResponseHeaderModificationShrink(v string) *UpdateHttpResponseHeaderModificationRuleShrinkRequest {
	s.ResponseHeaderModificationShrink = &v
	return s
}

func (s *UpdateHttpResponseHeaderModificationRuleShrinkRequest) SetRule(v string) *UpdateHttpResponseHeaderModificationRuleShrinkRequest {
	s.Rule = &v
	return s
}

func (s *UpdateHttpResponseHeaderModificationRuleShrinkRequest) SetRuleEnable(v string) *UpdateHttpResponseHeaderModificationRuleShrinkRequest {
	s.RuleEnable = &v
	return s
}

func (s *UpdateHttpResponseHeaderModificationRuleShrinkRequest) SetRuleName(v string) *UpdateHttpResponseHeaderModificationRuleShrinkRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateHttpResponseHeaderModificationRuleShrinkRequest) SetSequence(v int32) *UpdateHttpResponseHeaderModificationRuleShrinkRequest {
	s.Sequence = &v
	return s
}

func (s *UpdateHttpResponseHeaderModificationRuleShrinkRequest) SetSiteId(v int64) *UpdateHttpResponseHeaderModificationRuleShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateHttpResponseHeaderModificationRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
