// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationCustomRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationTemplateId(v string) *CreateApplicationCustomRulesRequest
	GetApplicationTemplateId() *string
	SetRules(v []*CreateApplicationCustomRulesRequestRules) *CreateApplicationCustomRulesRequest
	GetRules() []*CreateApplicationCustomRulesRequestRules
}

type CreateApplicationCustomRulesRequest struct {
	// The ID of the application template to which the custom rules belong. The template must exist within the current Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// ati-4f0691c56b8d4f06
	ApplicationTemplateId *string `json:"ApplicationTemplateId,omitempty" xml:"ApplicationTemplateId,omitempty"`
	// The list of custom rules for the application. You can specify up to 100 rules at a time, and the total number of rules in the template after creation must not exceed 100.
	//
	// This parameter is required.
	Rules []*CreateApplicationCustomRulesRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s CreateApplicationCustomRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationCustomRulesRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationCustomRulesRequest) GetApplicationTemplateId() *string {
	return s.ApplicationTemplateId
}

func (s *CreateApplicationCustomRulesRequest) GetRules() []*CreateApplicationCustomRulesRequestRules {
	return s.Rules
}

func (s *CreateApplicationCustomRulesRequest) SetApplicationTemplateId(v string) *CreateApplicationCustomRulesRequest {
	s.ApplicationTemplateId = &v
	return s
}

func (s *CreateApplicationCustomRulesRequest) SetRules(v []*CreateApplicationCustomRulesRequestRules) *CreateApplicationCustomRulesRequest {
	s.Rules = v
	return s
}

func (s *CreateApplicationCustomRulesRequest) Validate() error {
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

type CreateApplicationCustomRulesRequestRules struct {
	// The name of the custom rule. The name cannot be empty and can be up to 64 characters in length. It can contain Chinese characters, English letters, digits, hyphens (-), underscores (_), and periods (.).
	//
	// example:
	//
	// CustomRule01
	ApplicationCustomRuleName *string `json:"ApplicationCustomRuleName,omitempty" xml:"ApplicationCustomRuleName,omitempty"`
	// The list of application IDs associated with the rule. A single rule can be associated with 1 to 200 applications. The IDs cannot be 0, must be unique, and must exist in the built-in application signature library. You can call the ListApplications operation to query the IDs.
	ApplicationIds []*int64 `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// The priority of the rule. Valid values: integers from 1 to 200. A smaller value indicates a higher priority. Priorities must be unique within the same batch and the same template.
	//
	// example:
	//
	// 1
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The action to perform when the rule is matched. Valid values:
	//
	// - **Allow**: allows the traffic.
	//
	// - **Alert**: monitors the traffic.
	//
	// - **Deny**: blocks the traffic.
	//
	// example:
	//
	// Deny
	RuleAction *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
}

func (s CreateApplicationCustomRulesRequestRules) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationCustomRulesRequestRules) GoString() string {
	return s.String()
}

func (s *CreateApplicationCustomRulesRequestRules) GetApplicationCustomRuleName() *string {
	return s.ApplicationCustomRuleName
}

func (s *CreateApplicationCustomRulesRequestRules) GetApplicationIds() []*int64 {
	return s.ApplicationIds
}

func (s *CreateApplicationCustomRulesRequestRules) GetPriority() *int64 {
	return s.Priority
}

func (s *CreateApplicationCustomRulesRequestRules) GetRuleAction() *string {
	return s.RuleAction
}

func (s *CreateApplicationCustomRulesRequestRules) SetApplicationCustomRuleName(v string) *CreateApplicationCustomRulesRequestRules {
	s.ApplicationCustomRuleName = &v
	return s
}

func (s *CreateApplicationCustomRulesRequestRules) SetApplicationIds(v []*int64) *CreateApplicationCustomRulesRequestRules {
	s.ApplicationIds = v
	return s
}

func (s *CreateApplicationCustomRulesRequestRules) SetPriority(v int64) *CreateApplicationCustomRulesRequestRules {
	s.Priority = &v
	return s
}

func (s *CreateApplicationCustomRulesRequestRules) SetRuleAction(v string) *CreateApplicationCustomRulesRequestRules {
	s.RuleAction = &v
	return s
}

func (s *CreateApplicationCustomRulesRequestRules) Validate() error {
	return dara.Validate(s)
}
