// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRules interface {
	dara.Model
	String() string
	GoString() string
	SetSchedulingRule(v *SchedulingRule) *Rules
	GetSchedulingRule() *SchedulingRule
}

type Rules struct {
	SchedulingRule *SchedulingRule `json:"SchedulingRule,omitempty" xml:"SchedulingRule,omitempty"`
}

func (s Rules) String() string {
	return dara.Prettify(s)
}

func (s Rules) GoString() string {
	return s.String()
}

func (s *Rules) GetSchedulingRule() *SchedulingRule {
	return s.SchedulingRule
}

func (s *Rules) SetSchedulingRule(v *SchedulingRule) *Rules {
	s.SchedulingRule = v
	return s
}

func (s *Rules) Validate() error {
	if s.SchedulingRule != nil {
		if err := s.SchedulingRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}
