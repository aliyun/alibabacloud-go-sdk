// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWafRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v *WafRuleConfig) *GetWafRuleResponseBody
	GetConfig() *WafRuleConfig
	SetId(v int64) *GetWafRuleResponseBody
	GetId() *int64
	SetName(v string) *GetWafRuleResponseBody
	GetName() *string
	SetPhase(v string) *GetWafRuleResponseBody
	GetPhase() *string
	SetPosition(v int64) *GetWafRuleResponseBody
	GetPosition() *int64
	SetRequestId(v string) *GetWafRuleResponseBody
	GetRequestId() *string
	SetRulesetId(v int64) *GetWafRuleResponseBody
	GetRulesetId() *int64
	SetStatus(v string) *GetWafRuleResponseBody
	GetStatus() *string
	SetUpdateTime(v string) *GetWafRuleResponseBody
	GetUpdateTime() *string
}

type GetWafRuleResponseBody struct {
	// The rule configuration.
	Config *WafRuleConfig `json:"Config,omitempty" xml:"Config,omitempty"`
	// The ID of the WAF rule. You can get this ID by calling the [ListWafRules](https://help.aliyun.com/document_detail/2878257.html) operation.
	//
	// example:
	//
	// 2000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The execution phase of the WAF rule.
	//
	// - `http_whitelist`: A whitelist rule
	//
	// - `http_custom`: A custom rule
	//
	// - `http_managed`: A managed rule
	//
	// - `http_anti_scan`: A scan protection rule
	//
	// - `http_ratelimit`: A rate limiting rule
	//
	// - `ip_access_rule`: An IP access rule
	//
	// - `http_bot`: A bot management rule
	//
	// - `http_security_level_rule`: A security rule
	//
	// This parameter is required.
	//
	// example:
	//
	// http_custom
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// The position of the rule in the ruleset.
	//
	// example:
	//
	// 1
	Position *int64 `json:"Position,omitempty" xml:"Position,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the WAF ruleset. You can get this ID by calling the [ListWafRulesets](https://help.aliyun.com/document_detail/2878359.html) operation.
	RulesetId *int64 `json:"RulesetId,omitempty" xml:"RulesetId,omitempty"`
	// The status of the rule.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time the rule was last updated.
	//
	// example:
	//
	// 2024-01-01T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetWafRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWafRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetWafRuleResponseBody) GetConfig() *WafRuleConfig {
	return s.Config
}

func (s *GetWafRuleResponseBody) GetId() *int64 {
	return s.Id
}

func (s *GetWafRuleResponseBody) GetName() *string {
	return s.Name
}

func (s *GetWafRuleResponseBody) GetPhase() *string {
	return s.Phase
}

func (s *GetWafRuleResponseBody) GetPosition() *int64 {
	return s.Position
}

func (s *GetWafRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWafRuleResponseBody) GetRulesetId() *int64 {
	return s.RulesetId
}

func (s *GetWafRuleResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetWafRuleResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetWafRuleResponseBody) SetConfig(v *WafRuleConfig) *GetWafRuleResponseBody {
	s.Config = v
	return s
}

func (s *GetWafRuleResponseBody) SetId(v int64) *GetWafRuleResponseBody {
	s.Id = &v
	return s
}

func (s *GetWafRuleResponseBody) SetName(v string) *GetWafRuleResponseBody {
	s.Name = &v
	return s
}

func (s *GetWafRuleResponseBody) SetPhase(v string) *GetWafRuleResponseBody {
	s.Phase = &v
	return s
}

func (s *GetWafRuleResponseBody) SetPosition(v int64) *GetWafRuleResponseBody {
	s.Position = &v
	return s
}

func (s *GetWafRuleResponseBody) SetRequestId(v string) *GetWafRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWafRuleResponseBody) SetRulesetId(v int64) *GetWafRuleResponseBody {
	s.RulesetId = &v
	return s
}

func (s *GetWafRuleResponseBody) SetStatus(v string) *GetWafRuleResponseBody {
	s.Status = &v
	return s
}

func (s *GetWafRuleResponseBody) SetUpdateTime(v string) *GetWafRuleResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetWafRuleResponseBody) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}
