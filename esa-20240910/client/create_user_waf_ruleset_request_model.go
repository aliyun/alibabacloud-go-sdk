// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserWafRulesetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateUserWafRulesetRequest
	GetDescription() *string
	SetExpression(v string) *CreateUserWafRulesetRequest
	GetExpression() *string
	SetInstanceId(v string) *CreateUserWafRulesetRequest
	GetInstanceId() *string
	SetName(v string) *CreateUserWafRulesetRequest
	GetName() *string
	SetPhase(v string) *CreateUserWafRulesetRequest
	GetPhase() *string
	SetRules(v []*WafRuleConfig) *CreateUserWafRulesetRequest
	GetRules() []*WafRuleConfig
	SetShared(v *WafBatchRuleShared) *CreateUserWafRulesetRequest
	GetShared() *WafBatchRuleShared
	SetStatus(v string) *CreateUserWafRulesetRequest
	GetStatus() *string
}

type CreateUserWafRulesetRequest struct {
	// The description of the WAF ruleset.
	//
	// example:
	//
	// this is a test ruleset.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The match expression of the WAF ruleset. Rules in this ruleset are evaluated only when a request matches this expression.
	//
	// Examples:
	//
	// - `http.host eq "example.com"` — Only requests with the host example.com enter this ruleset.
	//
	// - `starts_with(http.uri.path, "/api/")` — Only requests with the /api/ prefix enter this ruleset.
	//
	// > The complete expression syntax and available field set are subject to the server-side wirefilter dialect.
	//
	// This parameter is required.
	//
	// example:
	//
	// ip.src == 1.1.1.1
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// esa-site-ads11w
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the WAF ruleset.
	//
	// **Naming suggestion**: Use a combination of letters, digits, and underscores for easy reference. The specific character set, maximum length, and uniqueness constraints are subject to the WAF ruleset service naming conventions.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The phase to which the WAF ruleset belongs. Valid values:
	//
	// - http_whitelist: whitelist rules
	//
	// - http_custom: custom rules
	//
	// - http_managed: managed rules
	//
	// - http_anti_scan: scan protection rules
	//
	// - http_ratelimit: rate limiting rules
	//
	// > Note: The supported fields (Expression match fields, Action values, and others) vary by phase. For more information, refer to the rule configuration documentation for the corresponding phase.
	//
	// This parameter is required.
	//
	// example:
	//
	// http_custom
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// The list of rule configurations in the WAF ruleset. Each element corresponds to a rule.
	//
	// - The field structure of each rule is subject to the `WafRuleConfig` data structure, which includes Expression, Action, Name, and other fields.
	Rules []*WafRuleConfig `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The shared fields across multiple rules in this ruleset, such as a unified Action or Name prefix.
	//
	// > The field structure is subject to the `WafBatchRuleShared` data structure. If you do not need to share properties, you can leave this parameter empty.
	Shared *WafBatchRuleShared `json:"Shared,omitempty" xml:"Shared,omitempty"`
	// The status of the WAF ruleset. Valid values:
	//
	// - on: Enabled. The rules in the ruleset participate in matching and blocking.
	//
	// - off: Disabled. The ruleset is retained but does not participate in matching.
	//
	// > The complete set of valid values is subject to the server-side enum.
	//
	// This parameter is required.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateUserWafRulesetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserWafRulesetRequest) GoString() string {
	return s.String()
}

func (s *CreateUserWafRulesetRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateUserWafRulesetRequest) GetExpression() *string {
	return s.Expression
}

func (s *CreateUserWafRulesetRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateUserWafRulesetRequest) GetName() *string {
	return s.Name
}

func (s *CreateUserWafRulesetRequest) GetPhase() *string {
	return s.Phase
}

func (s *CreateUserWafRulesetRequest) GetRules() []*WafRuleConfig {
	return s.Rules
}

func (s *CreateUserWafRulesetRequest) GetShared() *WafBatchRuleShared {
	return s.Shared
}

func (s *CreateUserWafRulesetRequest) GetStatus() *string {
	return s.Status
}

func (s *CreateUserWafRulesetRequest) SetDescription(v string) *CreateUserWafRulesetRequest {
	s.Description = &v
	return s
}

func (s *CreateUserWafRulesetRequest) SetExpression(v string) *CreateUserWafRulesetRequest {
	s.Expression = &v
	return s
}

func (s *CreateUserWafRulesetRequest) SetInstanceId(v string) *CreateUserWafRulesetRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateUserWafRulesetRequest) SetName(v string) *CreateUserWafRulesetRequest {
	s.Name = &v
	return s
}

func (s *CreateUserWafRulesetRequest) SetPhase(v string) *CreateUserWafRulesetRequest {
	s.Phase = &v
	return s
}

func (s *CreateUserWafRulesetRequest) SetRules(v []*WafRuleConfig) *CreateUserWafRulesetRequest {
	s.Rules = v
	return s
}

func (s *CreateUserWafRulesetRequest) SetShared(v *WafBatchRuleShared) *CreateUserWafRulesetRequest {
	s.Shared = v
	return s
}

func (s *CreateUserWafRulesetRequest) SetStatus(v string) *CreateUserWafRulesetRequest {
	s.Status = &v
	return s
}

func (s *CreateUserWafRulesetRequest) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Shared != nil {
		if err := s.Shared.Validate(); err != nil {
			return err
		}
	}
	return nil
}
