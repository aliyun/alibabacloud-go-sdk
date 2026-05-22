// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomResponseCodeRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *UpdateCustomResponseCodeRuleRequest
	GetConfigId() *int64
	SetPageId(v string) *UpdateCustomResponseCodeRuleRequest
	GetPageId() *string
	SetReturnCode(v string) *UpdateCustomResponseCodeRuleRequest
	GetReturnCode() *string
	SetRule(v string) *UpdateCustomResponseCodeRuleRequest
	GetRule() *string
	SetRuleEnable(v string) *UpdateCustomResponseCodeRuleRequest
	GetRuleEnable() *string
	SetRuleName(v string) *UpdateCustomResponseCodeRuleRequest
	GetRuleName() *string
	SetSequence(v int32) *UpdateCustomResponseCodeRuleRequest
	GetSequence() *int32
	SetSiteId(v int64) *UpdateCustomResponseCodeRuleRequest
	GetSiteId() *int64
}

type UpdateCustomResponseCodeRuleRequest struct {
	// The ID of the configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Response page.
	//
	// example:
	//
	// 0
	PageId *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	// The response code.
	//
	// example:
	//
	// 200
	ReturnCode *string `json:"ReturnCode,omitempty" xml:"ReturnCode,omitempty"`
	// The content of the rule. A conditional expression is used to match a user request. You do not need to set this parameter when you add global configurations. Use cases:
	//
	// 	- true: Match all incoming requests.
	//
	// 	- Set the value to a custom expression, for example, (http.host eq "video.example.com"): Match the specified request.
	//
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether to enable the rule. Valid values: You do not need to set this parameter when you add global configurations. Valid values:
	//
	// 	- on
	//
	// 	- off
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name. You do not need to set this parameter when you add global configurations.
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
	// 437375513708224
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateCustomResponseCodeRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomResponseCodeRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomResponseCodeRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *UpdateCustomResponseCodeRuleRequest) GetPageId() *string {
	return s.PageId
}

func (s *UpdateCustomResponseCodeRuleRequest) GetReturnCode() *string {
	return s.ReturnCode
}

func (s *UpdateCustomResponseCodeRuleRequest) GetRule() *string {
	return s.Rule
}

func (s *UpdateCustomResponseCodeRuleRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *UpdateCustomResponseCodeRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateCustomResponseCodeRuleRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *UpdateCustomResponseCodeRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateCustomResponseCodeRuleRequest) SetConfigId(v int64) *UpdateCustomResponseCodeRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateCustomResponseCodeRuleRequest) SetPageId(v string) *UpdateCustomResponseCodeRuleRequest {
	s.PageId = &v
	return s
}

func (s *UpdateCustomResponseCodeRuleRequest) SetReturnCode(v string) *UpdateCustomResponseCodeRuleRequest {
	s.ReturnCode = &v
	return s
}

func (s *UpdateCustomResponseCodeRuleRequest) SetRule(v string) *UpdateCustomResponseCodeRuleRequest {
	s.Rule = &v
	return s
}

func (s *UpdateCustomResponseCodeRuleRequest) SetRuleEnable(v string) *UpdateCustomResponseCodeRuleRequest {
	s.RuleEnable = &v
	return s
}

func (s *UpdateCustomResponseCodeRuleRequest) SetRuleName(v string) *UpdateCustomResponseCodeRuleRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateCustomResponseCodeRuleRequest) SetSequence(v int32) *UpdateCustomResponseCodeRuleRequest {
	s.Sequence = &v
	return s
}

func (s *UpdateCustomResponseCodeRuleRequest) SetSiteId(v int64) *UpdateCustomResponseCodeRuleRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateCustomResponseCodeRuleRequest) Validate() error {
	return dara.Validate(s)
}
