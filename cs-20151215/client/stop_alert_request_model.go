// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAlertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertRuleGroupName(v string) *StopAlertRequest
	GetAlertRuleGroupName() *string
	SetAlertRuleName(v string) *StopAlertRequest
	GetAlertRuleName() *string
}

type StopAlertRequest struct {
	// The name of the alert rule group.
	//
	// example:
	//
	// sample
	AlertRuleGroupName *string `json:"alert_rule_group_name,omitempty" xml:"alert_rule_group_name,omitempty"`
	// The name of the alert rule.
	//
	// example:
	//
	// sample
	AlertRuleName *string `json:"alert_rule_name,omitempty" xml:"alert_rule_name,omitempty"`
}

func (s StopAlertRequest) String() string {
	return dara.Prettify(s)
}

func (s StopAlertRequest) GoString() string {
	return s.String()
}

func (s *StopAlertRequest) GetAlertRuleGroupName() *string {
	return s.AlertRuleGroupName
}

func (s *StopAlertRequest) GetAlertRuleName() *string {
	return s.AlertRuleName
}

func (s *StopAlertRequest) SetAlertRuleGroupName(v string) *StopAlertRequest {
	s.AlertRuleGroupName = &v
	return s
}

func (s *StopAlertRequest) SetAlertRuleName(v string) *StopAlertRequest {
	s.AlertRuleName = &v
	return s
}

func (s *StopAlertRequest) Validate() error {
	return dara.Validate(s)
}
