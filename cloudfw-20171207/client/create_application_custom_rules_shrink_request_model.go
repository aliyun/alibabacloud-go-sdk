// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationCustomRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationTemplateId(v string) *CreateApplicationCustomRulesShrinkRequest
	GetApplicationTemplateId() *string
	SetRulesShrink(v string) *CreateApplicationCustomRulesShrinkRequest
	GetRulesShrink() *string
}

type CreateApplicationCustomRulesShrinkRequest struct {
	// The ID of the application template to which the custom rules belong. The template must exist within the current Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// ati-4f0691c56b8d4f06
	ApplicationTemplateId *string `json:"ApplicationTemplateId,omitempty" xml:"ApplicationTemplateId,omitempty"`
	// The list of custom rules for the application. You can specify up to 100 rules at a time, and the total number of rules in the template after creation must not exceed 100.
	//
	// This parameter is required.
	RulesShrink *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
}

func (s CreateApplicationCustomRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationCustomRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationCustomRulesShrinkRequest) GetApplicationTemplateId() *string {
	return s.ApplicationTemplateId
}

func (s *CreateApplicationCustomRulesShrinkRequest) GetRulesShrink() *string {
	return s.RulesShrink
}

func (s *CreateApplicationCustomRulesShrinkRequest) SetApplicationTemplateId(v string) *CreateApplicationCustomRulesShrinkRequest {
	s.ApplicationTemplateId = &v
	return s
}

func (s *CreateApplicationCustomRulesShrinkRequest) SetRulesShrink(v string) *CreateApplicationCustomRulesShrinkRequest {
	s.RulesShrink = &v
	return s
}

func (s *CreateApplicationCustomRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
