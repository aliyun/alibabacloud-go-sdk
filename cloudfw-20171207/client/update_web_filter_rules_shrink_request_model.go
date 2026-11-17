// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWebFilterRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRulesShrink(v string) *UpdateWebFilterRulesShrinkRequest
	GetRulesShrink() *string
}

type UpdateWebFilterRulesShrinkRequest struct {
	// The list of web filtering rules to be updated. You can specify up to 20 rules at a time. The rule IDs in the list must be unique.
	//
	// This parameter is required.
	RulesShrink *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
}

func (s UpdateWebFilterRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWebFilterRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateWebFilterRulesShrinkRequest) GetRulesShrink() *string {
	return s.RulesShrink
}

func (s *UpdateWebFilterRulesShrinkRequest) SetRulesShrink(v string) *UpdateWebFilterRulesShrinkRequest {
	s.RulesShrink = &v
	return s
}

func (s *UpdateWebFilterRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
