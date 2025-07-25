// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRules interface {
	dara.Model
	String() string
	GoString() string
	SetScheduling(v *SchedulingRule) *Rules
	GetScheduling() *SchedulingRule
}

type Rules struct {
	Scheduling *SchedulingRule `json:"Scheduling,omitempty" xml:"Scheduling,omitempty"`
}

func (s Rules) String() string {
	return dara.Prettify(s)
}

func (s Rules) GoString() string {
	return s.String()
}

func (s *Rules) GetScheduling() *SchedulingRule {
	return s.Scheduling
}

func (s *Rules) SetScheduling(v *SchedulingRule) *Rules {
	s.Scheduling = v
	return s
}

func (s *Rules) Validate() error {
	return dara.Validate(s)
}
