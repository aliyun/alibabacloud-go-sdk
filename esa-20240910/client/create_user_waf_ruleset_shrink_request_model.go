// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserWafRulesetShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateUserWafRulesetShrinkRequest
	GetDescription() *string
	SetExpression(v string) *CreateUserWafRulesetShrinkRequest
	GetExpression() *string
	SetInstanceId(v string) *CreateUserWafRulesetShrinkRequest
	GetInstanceId() *string
	SetName(v string) *CreateUserWafRulesetShrinkRequest
	GetName() *string
	SetPhase(v string) *CreateUserWafRulesetShrinkRequest
	GetPhase() *string
	SetRulesShrink(v string) *CreateUserWafRulesetShrinkRequest
	GetRulesShrink() *string
	SetSharedShrink(v string) *CreateUserWafRulesetShrinkRequest
	GetSharedShrink() *string
	SetStatus(v string) *CreateUserWafRulesetShrinkRequest
	GetStatus() *string
}

type CreateUserWafRulesetShrinkRequest struct {
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
	RulesShrink *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	// The shared fields across multiple rules in this ruleset, such as a unified Action or Name prefix.
	//
	// > The field structure is subject to the `WafBatchRuleShared` data structure. If you do not need to share properties, you can leave this parameter empty.
	SharedShrink *string `json:"Shared,omitempty" xml:"Shared,omitempty"`
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

func (s CreateUserWafRulesetShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserWafRulesetShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateUserWafRulesetShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateUserWafRulesetShrinkRequest) GetExpression() *string {
	return s.Expression
}

func (s *CreateUserWafRulesetShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateUserWafRulesetShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateUserWafRulesetShrinkRequest) GetPhase() *string {
	return s.Phase
}

func (s *CreateUserWafRulesetShrinkRequest) GetRulesShrink() *string {
	return s.RulesShrink
}

func (s *CreateUserWafRulesetShrinkRequest) GetSharedShrink() *string {
	return s.SharedShrink
}

func (s *CreateUserWafRulesetShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *CreateUserWafRulesetShrinkRequest) SetDescription(v string) *CreateUserWafRulesetShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateUserWafRulesetShrinkRequest) SetExpression(v string) *CreateUserWafRulesetShrinkRequest {
	s.Expression = &v
	return s
}

func (s *CreateUserWafRulesetShrinkRequest) SetInstanceId(v string) *CreateUserWafRulesetShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateUserWafRulesetShrinkRequest) SetName(v string) *CreateUserWafRulesetShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateUserWafRulesetShrinkRequest) SetPhase(v string) *CreateUserWafRulesetShrinkRequest {
	s.Phase = &v
	return s
}

func (s *CreateUserWafRulesetShrinkRequest) SetRulesShrink(v string) *CreateUserWafRulesetShrinkRequest {
	s.RulesShrink = &v
	return s
}

func (s *CreateUserWafRulesetShrinkRequest) SetSharedShrink(v string) *CreateUserWafRulesetShrinkRequest {
	s.SharedShrink = &v
	return s
}

func (s *CreateUserWafRulesetShrinkRequest) SetStatus(v string) *CreateUserWafRulesetShrinkRequest {
	s.Status = &v
	return s
}

func (s *CreateUserWafRulesetShrinkRequest) Validate() error {
	return dara.Validate(s)
}
