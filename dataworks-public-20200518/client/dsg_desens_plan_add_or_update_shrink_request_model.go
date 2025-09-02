// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgDesensPlanAddOrUpdateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesensRulesShrink(v string) *DsgDesensPlanAddOrUpdateShrinkRequest
	GetDesensRulesShrink() *string
}

type DsgDesensPlanAddOrUpdateShrinkRequest struct {
	// A collection of data masking rules that you want to add or modify.
	//
	// This parameter is required.
	DesensRulesShrink *string `json:"DesensRules,omitempty" xml:"DesensRules,omitempty"`
}

func (s DsgDesensPlanAddOrUpdateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgDesensPlanAddOrUpdateShrinkRequest) GoString() string {
	return s.String()
}

func (s *DsgDesensPlanAddOrUpdateShrinkRequest) GetDesensRulesShrink() *string {
	return s.DesensRulesShrink
}

func (s *DsgDesensPlanAddOrUpdateShrinkRequest) SetDesensRulesShrink(v string) *DsgDesensPlanAddOrUpdateShrinkRequest {
	s.DesensRulesShrink = &v
	return s
}

func (s *DsgDesensPlanAddOrUpdateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
