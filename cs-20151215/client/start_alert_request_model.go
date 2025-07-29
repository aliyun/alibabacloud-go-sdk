// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAlertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertRuleGroupName(v string) *StartAlertRequest
	GetAlertRuleGroupName() *string
	SetAlertRuleName(v string) *StartAlertRequest
	GetAlertRuleName() *string
}

type StartAlertRequest struct {
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

func (s StartAlertRequest) String() string {
	return dara.Prettify(s)
}

func (s StartAlertRequest) GoString() string {
	return s.String()
}

func (s *StartAlertRequest) GetAlertRuleGroupName() *string {
	return s.AlertRuleGroupName
}

func (s *StartAlertRequest) GetAlertRuleName() *string {
	return s.AlertRuleName
}

func (s *StartAlertRequest) SetAlertRuleGroupName(v string) *StartAlertRequest {
	s.AlertRuleGroupName = &v
	return s
}

func (s *StartAlertRequest) SetAlertRuleName(v string) *StartAlertRequest {
	s.AlertRuleName = &v
	return s
}

func (s *StartAlertRequest) Validate() error {
	return dara.Validate(s)
}
