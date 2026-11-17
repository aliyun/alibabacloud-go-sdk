// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWebFilterRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRules(v []*UpdateWebFilterRulesRequestRules) *UpdateWebFilterRulesRequest
	GetRules() []*UpdateWebFilterRulesRequestRules
}

type UpdateWebFilterRulesRequest struct {
	// The list of web filtering rules to be updated. You can specify up to 20 rules at a time. The rule IDs in the list must be unique.
	//
	// This parameter is required.
	Rules []*UpdateWebFilterRulesRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s UpdateWebFilterRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWebFilterRulesRequest) GoString() string {
	return s.String()
}

func (s *UpdateWebFilterRulesRequest) GetRules() []*UpdateWebFilterRulesRequestRules {
	return s.Rules
}

func (s *UpdateWebFilterRulesRequest) SetRules(v []*UpdateWebFilterRulesRequestRules) *UpdateWebFilterRulesRequest {
	s.Rules = v
	return s
}

func (s *UpdateWebFilterRulesRequest) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateWebFilterRulesRequestRules struct {
	// The list of match conditions for the rule. The list must contain exactly one condition. After the update, URLs within the same template must be unique.
	Conditions []*UpdateWebFilterRulesRequestRulesConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// The action to perform when the rule is matched. Valid values:
	//
	// - **Allow**: allows the request.
	//
	// - **Alert**: monitors the request.
	//
	// - **Deny**: blocks the request.
	//
	// example:
	//
	// Alert
	RuleAction *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	// The ID of the web filtering rule to be updated. The rule must exist within the current Alibaba Cloud account.
	//
	// example:
	//
	// wfr-8c1d2e3f4a5b6c7d
	WebFilterRuleId *string `json:"WebFilterRuleId,omitempty" xml:"WebFilterRuleId,omitempty"`
}

func (s UpdateWebFilterRulesRequestRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateWebFilterRulesRequestRules) GoString() string {
	return s.String()
}

func (s *UpdateWebFilterRulesRequestRules) GetConditions() []*UpdateWebFilterRulesRequestRulesConditions {
	return s.Conditions
}

func (s *UpdateWebFilterRulesRequestRules) GetRuleAction() *string {
	return s.RuleAction
}

func (s *UpdateWebFilterRulesRequestRules) GetWebFilterRuleId() *string {
	return s.WebFilterRuleId
}

func (s *UpdateWebFilterRulesRequestRules) SetConditions(v []*UpdateWebFilterRulesRequestRulesConditions) *UpdateWebFilterRulesRequestRules {
	s.Conditions = v
	return s
}

func (s *UpdateWebFilterRulesRequestRules) SetRuleAction(v string) *UpdateWebFilterRulesRequestRules {
	s.RuleAction = &v
	return s
}

func (s *UpdateWebFilterRulesRequestRules) SetWebFilterRuleId(v string) *UpdateWebFilterRulesRequestRules {
	s.WebFilterRuleId = &v
	return s
}

func (s *UpdateWebFilterRulesRequestRules) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateWebFilterRulesRequestRulesConditions struct {
	// The type of the match condition. Only **URL*	- is supported, which indicates matching by URL.
	//
	// example:
	//
	// URL
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The URL to match. The value cannot be empty and must not exceed 240 bytes in length. Each Chinese character is counted as 3 bytes. The URL must contain at least one forward slash (/), and the first segment must contain letters, digits, or Chinese characters. The URL cannot contain a protocol header, query parameters, or fragment identifiers. A single asterisk (*) matches all URLs. If a wildcard is used, only one wildcard is allowed and it must end with /*. URLs within the same template must be unique.
	//
	// example:
	//
	// www.example.com/login
	Values *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s UpdateWebFilterRulesRequestRulesConditions) String() string {
	return dara.Prettify(s)
}

func (s UpdateWebFilterRulesRequestRulesConditions) GoString() string {
	return s.String()
}

func (s *UpdateWebFilterRulesRequestRulesConditions) GetKey() *string {
	return s.Key
}

func (s *UpdateWebFilterRulesRequestRulesConditions) GetValues() *string {
	return s.Values
}

func (s *UpdateWebFilterRulesRequestRulesConditions) SetKey(v string) *UpdateWebFilterRulesRequestRulesConditions {
	s.Key = &v
	return s
}

func (s *UpdateWebFilterRulesRequestRulesConditions) SetValues(v string) *UpdateWebFilterRulesRequestRulesConditions {
	s.Values = &v
	return s
}

func (s *UpdateWebFilterRulesRequestRulesConditions) Validate() error {
	return dara.Validate(s)
}
