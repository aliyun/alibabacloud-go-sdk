// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWebFilterRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRulesShrink(v string) *CreateWebFilterRulesShrinkRequest
	GetRulesShrink() *string
	SetWebFilterTemplateId(v string) *CreateWebFilterRulesShrinkRequest
	GetWebFilterTemplateId() *string
}

type CreateWebFilterRulesShrinkRequest struct {
	// The list of web filtering rules. You can specify a maximum of 100 rules at a time, and the total number of rules in the template after creation must not exceed 100.
	//
	// This parameter is required.
	RulesShrink *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	// The ID of the web filtering template to which the web filtering rules belong. The template must exist within the current Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// wft-3f2a9c1d4b5e6f70
	WebFilterTemplateId *string `json:"WebFilterTemplateId,omitempty" xml:"WebFilterTemplateId,omitempty"`
}

func (s CreateWebFilterRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWebFilterRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateWebFilterRulesShrinkRequest) GetRulesShrink() *string {
	return s.RulesShrink
}

func (s *CreateWebFilterRulesShrinkRequest) GetWebFilterTemplateId() *string {
	return s.WebFilterTemplateId
}

func (s *CreateWebFilterRulesShrinkRequest) SetRulesShrink(v string) *CreateWebFilterRulesShrinkRequest {
	s.RulesShrink = &v
	return s
}

func (s *CreateWebFilterRulesShrinkRequest) SetWebFilterTemplateId(v string) *CreateWebFilterRulesShrinkRequest {
	s.WebFilterTemplateId = &v
	return s
}

func (s *CreateWebFilterRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
