// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlertRuleAction interface {
	dara.Model
	String() string
	GoString() string
	SetActions(v []*string) *AlertRuleAction
	GetActions() []*string
}

type AlertRuleAction struct {
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
}

func (s AlertRuleAction) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleAction) GoString() string {
	return s.String()
}

func (s *AlertRuleAction) GetActions() []*string {
	return s.Actions
}

func (s *AlertRuleAction) SetActions(v []*string) *AlertRuleAction {
	s.Actions = v
	return s
}

func (s *AlertRuleAction) Validate() error {
	return dara.Validate(s)
}
