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
	// Response page.
	//
	// example:
	//
	// 0
	PageId *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	// Response code.
	//
	// example:
	//
	// 400
	ReturnCode *string `json:"ReturnCode,omitempty" xml:"ReturnCode,omitempty"`
	// The content of the rule. A conditional expression is used to match a user request. You do not need to set this parameter when you add global configuration. Use cases:
	//
	// 	- true: Match all incoming requests.
	//
	// 	- Set the value to a custom expression, for example: (http.host eq "video.example.com"): Match the specified request.
	//
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether to enable the rule. Valid values: You do not need to set this parameter when you add global configuration. Valid values:
	//
	// 	- on
	//
	// 	- off
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name. You do not need to set this parameter when you add global configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The order in which the rule is executed. A smaller value gives priority to the rule.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 805864735361584
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version number of the website configurations. You can use this parameter to specify a version of your website to apply the feature settings. By default, version 0 is used.
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
