// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomResponseCodeRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageId(v string) *CreateCustomResponseCodeRuleRequest
	GetPageId() *string
	SetReturnCode(v string) *CreateCustomResponseCodeRuleRequest
	GetReturnCode() *string
	SetRule(v string) *CreateCustomResponseCodeRuleRequest
	GetRule() *string
	SetRuleEnable(v string) *CreateCustomResponseCodeRuleRequest
	GetRuleEnable() *string
	SetRuleName(v string) *CreateCustomResponseCodeRuleRequest
	GetRuleName() *string
	SetSequence(v int32) *CreateCustomResponseCodeRuleRequest
	GetSequence() *int32
	SetSiteId(v int64) *CreateCustomResponseCodeRuleRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *CreateCustomResponseCodeRuleRequest
	GetSiteVersion() *int32
}

type CreateCustomResponseCodeRuleRequest struct {
	// The response page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	PageId *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	// The response code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 400
	ReturnCode *string `json:"ReturnCode,omitempty" xml:"ReturnCode,omitempty"`
	// The rule content. Conditional expressions are used to match user requests. You do not need to set this parameter when adding a global configuration. Two scenarios are supported:
	//
	// - Match all incoming requests: Set the value to true.
	//
	// - Match specified requests: Set the value to a custom expression, such as (http.host eq "video.example.com").
	//
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// The rule switch. You do not need to set this parameter when adding a global configuration. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name. You do not need to set this parameter when adding a global configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The execution order of the rule. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The site ID. You can call [ListSites](https://help.aliyun.com/document_detail/2850189.html) to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 805864735361584
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version number of the site configuration. For sites with version management enabled, you can use this parameter to specify the site version on which the configuration takes effect. Default value: 0.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s CreateCustomResponseCodeRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomResponseCodeRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomResponseCodeRuleRequest) GetPageId() *string {
	return s.PageId
}

func (s *CreateCustomResponseCodeRuleRequest) GetReturnCode() *string {
	return s.ReturnCode
}

func (s *CreateCustomResponseCodeRuleRequest) GetRule() *string {
	return s.Rule
}

func (s *CreateCustomResponseCodeRuleRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *CreateCustomResponseCodeRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateCustomResponseCodeRuleRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *CreateCustomResponseCodeRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateCustomResponseCodeRuleRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *CreateCustomResponseCodeRuleRequest) SetPageId(v string) *CreateCustomResponseCodeRuleRequest {
	s.PageId = &v
	return s
}

func (s *CreateCustomResponseCodeRuleRequest) SetReturnCode(v string) *CreateCustomResponseCodeRuleRequest {
	s.ReturnCode = &v
	return s
}

func (s *CreateCustomResponseCodeRuleRequest) SetRule(v string) *CreateCustomResponseCodeRuleRequest {
	s.Rule = &v
	return s
}

func (s *CreateCustomResponseCodeRuleRequest) SetRuleEnable(v string) *CreateCustomResponseCodeRuleRequest {
	s.RuleEnable = &v
	return s
}

func (s *CreateCustomResponseCodeRuleRequest) SetRuleName(v string) *CreateCustomResponseCodeRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateCustomResponseCodeRuleRequest) SetSequence(v int32) *CreateCustomResponseCodeRuleRequest {
	s.Sequence = &v
	return s
}

func (s *CreateCustomResponseCodeRuleRequest) SetSiteId(v int64) *CreateCustomResponseCodeRuleRequest {
	s.SiteId = &v
	return s
}

func (s *CreateCustomResponseCodeRuleRequest) SetSiteVersion(v int32) *CreateCustomResponseCodeRuleRequest {
	s.SiteVersion = &v
	return s
}

func (s *CreateCustomResponseCodeRuleRequest) Validate() error {
	return dara.Validate(s)
}
