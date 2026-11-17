// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWebFilterRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRules(v []*CreateWebFilterRulesRequestRules) *CreateWebFilterRulesRequest
	GetRules() []*CreateWebFilterRulesRequestRules
	SetWebFilterTemplateId(v string) *CreateWebFilterRulesRequest
	GetWebFilterTemplateId() *string
}

type CreateWebFilterRulesRequest struct {
	// The list of web filtering rules. You can specify a maximum of 100 rules at a time, and the total number of rules in the template after creation must not exceed 100.
	//
	// This parameter is required.
	Rules []*CreateWebFilterRulesRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The ID of the web filtering template to which the web filtering rules belong. The template must exist within the current Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// wft-3f2a9c1d4b5e6f70
	WebFilterTemplateId *string `json:"WebFilterTemplateId,omitempty" xml:"WebFilterTemplateId,omitempty"`
}

func (s CreateWebFilterRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWebFilterRulesRequest) GoString() string {
	return s.String()
}

func (s *CreateWebFilterRulesRequest) GetRules() []*CreateWebFilterRulesRequestRules {
	return s.Rules
}

func (s *CreateWebFilterRulesRequest) GetWebFilterTemplateId() *string {
	return s.WebFilterTemplateId
}

func (s *CreateWebFilterRulesRequest) SetRules(v []*CreateWebFilterRulesRequestRules) *CreateWebFilterRulesRequest {
	s.Rules = v
	return s
}

func (s *CreateWebFilterRulesRequest) SetWebFilterTemplateId(v string) *CreateWebFilterRulesRequest {
	s.WebFilterTemplateId = &v
	return s
}

func (s *CreateWebFilterRulesRequest) Validate() error {
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

type CreateWebFilterRulesRequestRules struct {
	// The list of matching conditions for the rule. The list must contain exactly one condition.
	Conditions []*CreateWebFilterRulesRequestRulesConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// The action to perform when a rule is matched. Valid values:
	//
	// - **Allow**: allows the request.
	//
	// - **Alert**: monitors the request.
	//
	// - **Deny**: blocks the request.
	//
	// example:
	//
	// Deny
	RuleAction *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
}

func (s CreateWebFilterRulesRequestRules) String() string {
	return dara.Prettify(s)
}

func (s CreateWebFilterRulesRequestRules) GoString() string {
	return s.String()
}

func (s *CreateWebFilterRulesRequestRules) GetConditions() []*CreateWebFilterRulesRequestRulesConditions {
	return s.Conditions
}

func (s *CreateWebFilterRulesRequestRules) GetRuleAction() *string {
	return s.RuleAction
}

func (s *CreateWebFilterRulesRequestRules) SetConditions(v []*CreateWebFilterRulesRequestRulesConditions) *CreateWebFilterRulesRequestRules {
	s.Conditions = v
	return s
}

func (s *CreateWebFilterRulesRequestRules) SetRuleAction(v string) *CreateWebFilterRulesRequestRules {
	s.RuleAction = &v
	return s
}

func (s *CreateWebFilterRulesRequestRules) Validate() error {
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

type CreateWebFilterRulesRequestRulesConditions struct {
	// The type of the matching condition. Only **URL*	- is supported, which indicates matching by URL.
	//
	// example:
	//
	// URL
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The URL to match. It cannot be empty and must not exceed 240 bytes in length (each Chinese character counts as 3 bytes). It must contain at least one forward slash (/), and the first segment must contain letters, digits, or Chinese characters. It cannot contain a protocol header, query parameters, or fragment identifiers. A single asterisk (*) indicates matching all URLs. When a wildcard is used, a maximum of one wildcard is allowed and it must end with /*. URLs within the same template cannot be duplicated.
	//
	// example:
	//
	// www.example.com/login
	Values *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s CreateWebFilterRulesRequestRulesConditions) String() string {
	return dara.Prettify(s)
}

func (s CreateWebFilterRulesRequestRulesConditions) GoString() string {
	return s.String()
}

func (s *CreateWebFilterRulesRequestRulesConditions) GetKey() *string {
	return s.Key
}

func (s *CreateWebFilterRulesRequestRulesConditions) GetValues() *string {
	return s.Values
}

func (s *CreateWebFilterRulesRequestRulesConditions) SetKey(v string) *CreateWebFilterRulesRequestRulesConditions {
	s.Key = &v
	return s
}

func (s *CreateWebFilterRulesRequestRulesConditions) SetValues(v string) *CreateWebFilterRulesRequestRulesConditions {
	s.Values = &v
	return s
}

func (s *CreateWebFilterRulesRequestRulesConditions) Validate() error {
	return dara.Validate(s)
}
