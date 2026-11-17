// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationCustomRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRules(v []*UpdateApplicationCustomRulesRequestRules) *UpdateApplicationCustomRulesRequest
	GetRules() []*UpdateApplicationCustomRulesRequestRules
}

type UpdateApplicationCustomRulesRequest struct {
	// The list of custom application rules to be modified. You can specify up to 20 rules at a time. Rule IDs in the list must be unique.
	//
	// This parameter is required.
	Rules []*UpdateApplicationCustomRulesRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s UpdateApplicationCustomRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationCustomRulesRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationCustomRulesRequest) GetRules() []*UpdateApplicationCustomRulesRequestRules {
	return s.Rules
}

func (s *UpdateApplicationCustomRulesRequest) SetRules(v []*UpdateApplicationCustomRulesRequestRules) *UpdateApplicationCustomRulesRequest {
	s.Rules = v
	return s
}

func (s *UpdateApplicationCustomRulesRequest) Validate() error {
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

type UpdateApplicationCustomRulesRequestRules struct {
	// The list of application IDs to be associated. You can specify up to 200 IDs. The IDs cannot be 0 and must exist in the built-in application signature library. The IDs must not overlap with those in RemoveApplicationIds. The combined list cannot be empty and can contain up to 200 IDs.
	AddApplicationIds []*int64 `json:"AddApplicationIds,omitempty" xml:"AddApplicationIds,omitempty" type:"Repeated"`
	// The ID of the custom application rule to be modified. The rule must exist within the current Alibaba Cloud account.
	//
	// example:
	//
	// acr-7d2e4f6a8b0c1d3e
	ApplicationCustomRuleId *string `json:"ApplicationCustomRuleId,omitempty" xml:"ApplicationCustomRuleId,omitempty"`
	// The name of the custom rule. The name can be up to 64 characters in length and can contain Chinese characters, letters, digits, hyphens (-), underscores (_), and periods (.). If this parameter is not specified or is left empty, the original value is retained.
	//
	// example:
	//
	// CustomRule02
	ApplicationCustomRuleName *string `json:"ApplicationCustomRuleName,omitempty" xml:"ApplicationCustomRuleName,omitempty"`
	// The priority of the rule. This parameter is required. Valid values: integers from 1 to 200. A smaller value indicates a higher priority. Priorities must be unique within the same template. You can swap priorities by deleting and then re-inserting rules.
	//
	// example:
	//
	// 2
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The list of application IDs to be disassociated. You can specify up to 200 IDs. The IDs must not overlap with those in AddApplicationIds. The combined list cannot be empty and can contain up to 200 IDs.
	RemoveApplicationIds []*int64 `json:"RemoveApplicationIds,omitempty" xml:"RemoveApplicationIds,omitempty" type:"Repeated"`
	// The action to perform when the rule is matched. This parameter is required. Valid values:
	//
	// - **Allow**: allows the traffic.
	//
	// - **Alert**: monitors the traffic.
	//
	// - **Deny**: blocks the traffic.
	//
	// example:
	//
	// Alert
	RuleAction *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
}

func (s UpdateApplicationCustomRulesRequestRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationCustomRulesRequestRules) GoString() string {
	return s.String()
}

func (s *UpdateApplicationCustomRulesRequestRules) GetAddApplicationIds() []*int64 {
	return s.AddApplicationIds
}

func (s *UpdateApplicationCustomRulesRequestRules) GetApplicationCustomRuleId() *string {
	return s.ApplicationCustomRuleId
}

func (s *UpdateApplicationCustomRulesRequestRules) GetApplicationCustomRuleName() *string {
	return s.ApplicationCustomRuleName
}

func (s *UpdateApplicationCustomRulesRequestRules) GetPriority() *int64 {
	return s.Priority
}

func (s *UpdateApplicationCustomRulesRequestRules) GetRemoveApplicationIds() []*int64 {
	return s.RemoveApplicationIds
}

func (s *UpdateApplicationCustomRulesRequestRules) GetRuleAction() *string {
	return s.RuleAction
}

func (s *UpdateApplicationCustomRulesRequestRules) SetAddApplicationIds(v []*int64) *UpdateApplicationCustomRulesRequestRules {
	s.AddApplicationIds = v
	return s
}

func (s *UpdateApplicationCustomRulesRequestRules) SetApplicationCustomRuleId(v string) *UpdateApplicationCustomRulesRequestRules {
	s.ApplicationCustomRuleId = &v
	return s
}

func (s *UpdateApplicationCustomRulesRequestRules) SetApplicationCustomRuleName(v string) *UpdateApplicationCustomRulesRequestRules {
	s.ApplicationCustomRuleName = &v
	return s
}

func (s *UpdateApplicationCustomRulesRequestRules) SetPriority(v int64) *UpdateApplicationCustomRulesRequestRules {
	s.Priority = &v
	return s
}

func (s *UpdateApplicationCustomRulesRequestRules) SetRemoveApplicationIds(v []*int64) *UpdateApplicationCustomRulesRequestRules {
	s.RemoveApplicationIds = v
	return s
}

func (s *UpdateApplicationCustomRulesRequestRules) SetRuleAction(v string) *UpdateApplicationCustomRulesRequestRules {
	s.RuleAction = &v
	return s
}

func (s *UpdateApplicationCustomRulesRequestRules) Validate() error {
	return dara.Validate(s)
}
