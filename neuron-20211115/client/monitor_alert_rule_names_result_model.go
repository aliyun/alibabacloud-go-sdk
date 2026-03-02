// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorAlertRuleNamesResult interface {
	dara.Model
	String() string
	GoString() string
	SetRuleNames(v []*string) *MonitorAlertRuleNamesResult
	GetRuleNames() []*string
}

type MonitorAlertRuleNamesResult struct {
	RuleNames []*string `json:"ruleNames,omitempty" xml:"ruleNames,omitempty" type:"Repeated"`
}

func (s MonitorAlertRuleNamesResult) String() string {
	return dara.Prettify(s)
}

func (s MonitorAlertRuleNamesResult) GoString() string {
	return s.String()
}

func (s *MonitorAlertRuleNamesResult) GetRuleNames() []*string {
	return s.RuleNames
}

func (s *MonitorAlertRuleNamesResult) SetRuleNames(v []*string) *MonitorAlertRuleNamesResult {
	s.RuleNames = v
	return s
}

func (s *MonitorAlertRuleNamesResult) Validate() error {
	return dara.Validate(s)
}
