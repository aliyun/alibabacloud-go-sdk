// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserWafRulesetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetUserWafRulesetResponseBody
	GetRequestId() *string
	SetRuleset(v *GetUserWafRulesetResponseBodyRuleset) *GetUserWafRulesetResponseBody
	GetRuleset() *GetUserWafRulesetResponseBodyRuleset
}

type GetUserWafRulesetResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The WAF rule set.
	//
	// example:
	//
	// {}
	Ruleset *GetUserWafRulesetResponseBodyRuleset `json:"Ruleset,omitempty" xml:"Ruleset,omitempty" type:"Struct"`
}

func (s GetUserWafRulesetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserWafRulesetResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserWafRulesetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserWafRulesetResponseBody) GetRuleset() *GetUserWafRulesetResponseBodyRuleset {
	return s.Ruleset
}

func (s *GetUserWafRulesetResponseBody) SetRequestId(v string) *GetUserWafRulesetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserWafRulesetResponseBody) SetRuleset(v *GetUserWafRulesetResponseBodyRuleset) *GetUserWafRulesetResponseBody {
	s.Ruleset = v
	return s
}

func (s *GetUserWafRulesetResponseBody) Validate() error {
	if s.Ruleset != nil {
		if err := s.Ruleset.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserWafRulesetResponseBodyRuleset struct {
	// The description of the WAF rule set.
	//
	// example:
	//
	// example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The expression of the WAF rule set.
	//
	// example:
	//
	// ip.src == 1.1.1.1
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The ID of the WAF rule set.
	//
	// example:
	//
	// 10000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the WAF rule set.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The evaluation phase of the WAF rule set. Valid values:
	//
	// - `http_whitelist`: A whitelist rule.
	//
	// - `http_custom`: A custom rule.
	//
	// - `http_managed`: A managed rule.
	//
	// - `http_anti_scan`: A scan protection rule.
	//
	// - `http_ratelimit`: A rate limiting rule.
	//
	// - `ip_access_rule`: An IP access rule.
	//
	// - `http_bot`: A bot management rule.
	//
	// - `http_security_level_rule`: A security level rule.
	//
	// example:
	//
	// http_custom
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// The position of the WAF rule set.
	//
	// example:
	//
	// 1
	Position *int64 `json:"Position,omitempty" xml:"Position,omitempty"`
	// A list of rules in the WAF rule set.
	//
	// example:
	//
	// [{}]
	Rules []*GetUserWafRulesetResponseBodyRulesetRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The shared configuration of the WAF rule set.
	//
	// example:
	//
	// {}
	Shared *WafBatchRuleShared `json:"Shared,omitempty" xml:"Shared,omitempty"`
	// The status of the WAF rule set.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetUserWafRulesetResponseBodyRuleset) String() string {
	return dara.Prettify(s)
}

func (s GetUserWafRulesetResponseBodyRuleset) GoString() string {
	return s.String()
}

func (s *GetUserWafRulesetResponseBodyRuleset) GetDescription() *string {
	return s.Description
}

func (s *GetUserWafRulesetResponseBodyRuleset) GetExpression() *string {
	return s.Expression
}

func (s *GetUserWafRulesetResponseBodyRuleset) GetId() *int64 {
	return s.Id
}

func (s *GetUserWafRulesetResponseBodyRuleset) GetName() *string {
	return s.Name
}

func (s *GetUserWafRulesetResponseBodyRuleset) GetPhase() *string {
	return s.Phase
}

func (s *GetUserWafRulesetResponseBodyRuleset) GetPosition() *int64 {
	return s.Position
}

func (s *GetUserWafRulesetResponseBodyRuleset) GetRules() []*GetUserWafRulesetResponseBodyRulesetRules {
	return s.Rules
}

func (s *GetUserWafRulesetResponseBodyRuleset) GetShared() *WafBatchRuleShared {
	return s.Shared
}

func (s *GetUserWafRulesetResponseBodyRuleset) GetStatus() *string {
	return s.Status
}

func (s *GetUserWafRulesetResponseBodyRuleset) SetDescription(v string) *GetUserWafRulesetResponseBodyRuleset {
	s.Description = &v
	return s
}

func (s *GetUserWafRulesetResponseBodyRuleset) SetExpression(v string) *GetUserWafRulesetResponseBodyRuleset {
	s.Expression = &v
	return s
}

func (s *GetUserWafRulesetResponseBodyRuleset) SetId(v int64) *GetUserWafRulesetResponseBodyRuleset {
	s.Id = &v
	return s
}

func (s *GetUserWafRulesetResponseBodyRuleset) SetName(v string) *GetUserWafRulesetResponseBodyRuleset {
	s.Name = &v
	return s
}

func (s *GetUserWafRulesetResponseBodyRuleset) SetPhase(v string) *GetUserWafRulesetResponseBodyRuleset {
	s.Phase = &v
	return s
}

func (s *GetUserWafRulesetResponseBodyRuleset) SetPosition(v int64) *GetUserWafRulesetResponseBodyRuleset {
	s.Position = &v
	return s
}

func (s *GetUserWafRulesetResponseBodyRuleset) SetRules(v []*GetUserWafRulesetResponseBodyRulesetRules) *GetUserWafRulesetResponseBodyRuleset {
	s.Rules = v
	return s
}

func (s *GetUserWafRulesetResponseBodyRuleset) SetShared(v *WafBatchRuleShared) *GetUserWafRulesetResponseBodyRuleset {
	s.Shared = v
	return s
}

func (s *GetUserWafRulesetResponseBodyRuleset) SetStatus(v string) *GetUserWafRulesetResponseBodyRuleset {
	s.Status = &v
	return s
}

func (s *GetUserWafRulesetResponseBodyRuleset) Validate() error {
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

type GetUserWafRulesetResponseBodyRulesetRules struct {
	// The action for the rule. Valid values:
	//
	// - `deny`: Blocks the request.
	//
	// - `monitor`: Monitors the request.
	//
	// - `js`: Triggers a JS challenge.
	//
	// - `captcha`: Triggers a CAPTCHA challenge.
	//
	// example:
	//
	// deny
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// A list of WAF rule statistics fields.
	//
	// example:
	//
	// ["http.host"]
	CharacteristicsFields []*string `json:"CharacteristicsFields,omitempty" xml:"CharacteristicsFields,omitempty" type:"Repeated"`
	// The WAF rule configuration.
	//
	// example:
	//
	// {
	//
	//   "Id": 20000001,
	//
	//   "Name": "rule1",
	//
	//   "Expression": "ip.src eq 1.1.1.1",
	//
	//   "Action": "deny"
	//
	// }
	Config *WafRuleConfig `json:"Config,omitempty" xml:"Config,omitempty"`
	// A list of WAF rule match fields.
	//
	// example:
	//
	// ["ip.src"]
	Fields []*string `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// The ID of the WAF rule.
	//
	// example:
	//
	// 20000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the WAF rule.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The evaluation phase of the WAF rule. Valid values:
	//
	// - `http_whitelist`: A whitelist rule.
	//
	// - `http_custom`: A custom rule.
	//
	// - `http_managed`: A managed rule.
	//
	// - `http_anti_scan`: A scan protection rule.
	//
	// - `http_ratelimit`: A rate limiting rule.
	//
	// - `ip_access_rule`: An IP access rule.
	//
	// - `http_bot`: A bot management rule.
	//
	// - `http_security_level_rule`: A security level rule.
	//
	// example:
	//
	// http_custom
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// The position of the WAF rule.
	//
	// example:
	//
	// 1
	Position *int64 `json:"Position,omitempty" xml:"Position,omitempty"`
	// The ID of the WAF rule set.
	//
	// example:
	//
	// 10000001
	RulesetId *int64 `json:"RulesetId,omitempty" xml:"RulesetId,omitempty"`
	// The skip mode for the WAF rule.
	//
	// example:
	//
	// all
	Skip *string `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// The status of the WAF rule.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The phases that the rule skips.
	//
	// example:
	//
	// ["http_custom"]
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The type of the WAF rule.
	//
	// example:
	//
	// http_ratelimit
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The time the WAF rule was last updated.
	//
	// example:
	//
	// 2025-07-07T15:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetUserWafRulesetResponseBodyRulesetRules) String() string {
	return dara.Prettify(s)
}

func (s GetUserWafRulesetResponseBodyRulesetRules) GoString() string {
	return s.String()
}

func (s *GetUserWafRulesetResponseBodyRulesetRules) GetAction() *string {
	return s.Action
}

func (s *GetUserWafRulesetResponseBodyRulesetRules) GetCharacteristicsFields() []*string {
	return s.CharacteristicsFields
}

func (s *GetUserWafRulesetResponseBodyRulesetRules) GetConfig() *WafRuleConfig {
	return s.Config
}

func (s *GetUserWafRulesetResponseBodyRulesetRules) GetFields() []*string {
	return s.Fields
}

func (s *GetUserWafRulesetResponseBodyRulesetRules) GetId() *int64 {
	return s.Id
}

func (s *GetUserWafRulesetResponseBodyRulesetRules) GetName() *string {
	return s.Name
}

func (s *GetUserWafRulesetResponseBodyRulesetRules) GetPhase() *string {
	return s.Phase
}

func (s *GetUserWafRulesetResponseBodyRulesetRules) GetPosition() *int64 {
	return s.Position
}

func (s *GetUserWafRulesetResponseBodyRulesetRules) GetRulesetId() *int64 {
	return s.RulesetId
}

func (s *GetUserWafRulesetResponseBodyRulesetRules) GetSkip() *string {
	return s.Skip
}

func (s *GetUserWafRulesetResponseBodyRulesetRules) GetStatus() *string {
	return s.Status
}

func (s *GetUserWafRulesetResponseBodyRulesetRules) GetTags() []*string {
	return s.Tags
}

func (s *GetUserWafRulesetResponseBodyRulesetRules) GetType() *string {
	return s.Type
}

func (s *GetUserWafRulesetResponseBodyRulesetRules) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetUserWafRulesetResponseBodyRulesetRules) SetAction(v string) *GetUserWafRulesetResponseBodyRulesetRules {
	s.Action = &v
	return s
}

func (s *GetUserWafRulesetResponseBodyRulesetRules) SetCharacteristicsFields(v []*string) *GetUserWafRulesetResponseBodyRulesetRules {
	s.CharacteristicsFields = v
	return s
}

func (s *GetUserWafRulesetResponseBodyRulesetRules) SetConfig(v *WafRuleConfig) *GetUserWafRulesetResponseBodyRulesetRules {
	s.Config = v
	return s
}

func (s *GetUserWafRulesetResponseBodyRulesetRules) SetFields(v []*string) *GetUserWafRulesetResponseBodyRulesetRules {
	s.Fields = v
	return s
}

func (s *GetUserWafRulesetResponseBodyRulesetRules) SetId(v int64) *GetUserWafRulesetResponseBodyRulesetRules {
	s.Id = &v
	return s
}

func (s *GetUserWafRulesetResponseBodyRulesetRules) SetName(v string) *GetUserWafRulesetResponseBodyRulesetRules {
	s.Name = &v
	return s
}

func (s *GetUserWafRulesetResponseBodyRulesetRules) SetPhase(v string) *GetUserWafRulesetResponseBodyRulesetRules {
	s.Phase = &v
	return s
}

func (s *GetUserWafRulesetResponseBodyRulesetRules) SetPosition(v int64) *GetUserWafRulesetResponseBodyRulesetRules {
	s.Position = &v
	return s
}

func (s *GetUserWafRulesetResponseBodyRulesetRules) SetRulesetId(v int64) *GetUserWafRulesetResponseBodyRulesetRules {
	s.RulesetId = &v
	return s
}

func (s *GetUserWafRulesetResponseBodyRulesetRules) SetSkip(v string) *GetUserWafRulesetResponseBodyRulesetRules {
	s.Skip = &v
	return s
}

func (s *GetUserWafRulesetResponseBodyRulesetRules) SetStatus(v string) *GetUserWafRulesetResponseBodyRulesetRules {
	s.Status = &v
	return s
}

func (s *GetUserWafRulesetResponseBodyRulesetRules) SetTags(v []*string) *GetUserWafRulesetResponseBodyRulesetRules {
	s.Tags = v
	return s
}

func (s *GetUserWafRulesetResponseBodyRulesetRules) SetType(v string) *GetUserWafRulesetResponseBodyRulesetRules {
	s.Type = &v
	return s
}

func (s *GetUserWafRulesetResponseBodyRulesetRules) SetUpdateTime(v string) *GetUserWafRulesetResponseBodyRulesetRules {
	s.UpdateTime = &v
	return s
}

func (s *GetUserWafRulesetResponseBodyRulesetRules) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}
