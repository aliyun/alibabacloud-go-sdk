// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWafRulesetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetWafRulesetResponseBody
	GetId() *int64
	SetName(v string) *GetWafRulesetResponseBody
	GetName() *string
	SetPhase(v string) *GetWafRulesetResponseBody
	GetPhase() *string
	SetRequestId(v string) *GetWafRulesetResponseBody
	GetRequestId() *string
	SetRules(v []*WafRuleConfig) *GetWafRulesetResponseBody
	GetRules() []*WafRuleConfig
	SetShared(v *WafBatchRuleShared) *GetWafRulesetResponseBody
	GetShared() *WafBatchRuleShared
	SetStatus(v string) *GetWafRulesetResponseBody
	GetStatus() *string
	SetUpdateTime(v string) *GetWafRulesetResponseBody
	GetUpdateTime() *string
}

type GetWafRulesetResponseBody struct {
	// Ruleset ID.
	//
	// example:
	//
	// 10000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Ruleset name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The WAF operation phase applicable to the ruleset.
	//
	// This parameter is required.
	//
	// example:
	//
	// http_bot
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// List of rule configurations in the ruleset.
	Rules []*WafRuleConfig `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// Shared configurations for the rules in the ruleset.
	Shared *WafBatchRuleShared `json:"Shared,omitempty" xml:"Shared,omitempty"`
	// Ruleset status.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The last modified time of the ruleset.
	//
	// example:
	//
	// 2024-01-01T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetWafRulesetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWafRulesetResponseBody) GoString() string {
	return s.String()
}

func (s *GetWafRulesetResponseBody) GetId() *int64 {
	return s.Id
}

func (s *GetWafRulesetResponseBody) GetName() *string {
	return s.Name
}

func (s *GetWafRulesetResponseBody) GetPhase() *string {
	return s.Phase
}

func (s *GetWafRulesetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWafRulesetResponseBody) GetRules() []*WafRuleConfig {
	return s.Rules
}

func (s *GetWafRulesetResponseBody) GetShared() *WafBatchRuleShared {
	return s.Shared
}

func (s *GetWafRulesetResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetWafRulesetResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetWafRulesetResponseBody) SetId(v int64) *GetWafRulesetResponseBody {
	s.Id = &v
	return s
}

func (s *GetWafRulesetResponseBody) SetName(v string) *GetWafRulesetResponseBody {
	s.Name = &v
	return s
}

func (s *GetWafRulesetResponseBody) SetPhase(v string) *GetWafRulesetResponseBody {
	s.Phase = &v
	return s
}

func (s *GetWafRulesetResponseBody) SetRequestId(v string) *GetWafRulesetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWafRulesetResponseBody) SetRules(v []*WafRuleConfig) *GetWafRulesetResponseBody {
	s.Rules = v
	return s
}

func (s *GetWafRulesetResponseBody) SetShared(v *WafBatchRuleShared) *GetWafRulesetResponseBody {
	s.Shared = v
	return s
}

func (s *GetWafRulesetResponseBody) SetStatus(v string) *GetWafRulesetResponseBody {
	s.Status = &v
	return s
}

func (s *GetWafRulesetResponseBody) SetUpdateTime(v string) *GetWafRulesetResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetWafRulesetResponseBody) Validate() error {
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
