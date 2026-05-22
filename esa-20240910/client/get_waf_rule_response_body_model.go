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
	// Rule configuration.
	Config *WafRuleConfig `json:"Config,omitempty" xml:"Config,omitempty"`
	// The ID of the WAF rule, which can be obtained by calling the [ListWafRules](https://help.aliyun.com/document_detail/2878257.html) interface.
	//
	// example:
	//
	// 2000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Rule name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// WAF operation phase.
	//
	// This parameter is required.
	//
	// example:
	//
	// http_custom
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// The position of the rule in the rule set.
	//
	// example:
	//
	// 1
	Position *int64 `json:"Position,omitempty" xml:"Position,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RulesetId *int64  `json:"RulesetId,omitempty" xml:"RulesetId,omitempty"`
	// Rule status.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The last modified time of the rule.
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
