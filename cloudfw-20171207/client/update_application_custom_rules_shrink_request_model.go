// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationCustomRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRulesShrink(v string) *UpdateApplicationCustomRulesShrinkRequest
	GetRulesShrink() *string
}

type UpdateApplicationCustomRulesShrinkRequest struct {
	// The list of custom application rules to be modified. You can specify up to 20 rules at a time. Rule IDs in the list must be unique.
	//
	// This parameter is required.
	RulesShrink *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
}

func (s UpdateApplicationCustomRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationCustomRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationCustomRulesShrinkRequest) GetRulesShrink() *string {
	return s.RulesShrink
}

func (s *UpdateApplicationCustomRulesShrinkRequest) SetRulesShrink(v string) *UpdateApplicationCustomRulesShrinkRequest {
	s.RulesShrink = &v
	return s
}

func (s *UpdateApplicationCustomRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
