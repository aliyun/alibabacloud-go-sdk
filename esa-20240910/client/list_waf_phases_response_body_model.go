// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWafPhasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPhases(v []*ListWafPhasesResponseBodyPhases) *ListWafPhasesResponseBody
	GetPhases() []*ListWafPhasesResponseBodyPhases
	SetRequestId(v string) *ListWafPhasesResponseBody
	GetRequestId() *string
}

type ListWafPhasesResponseBody struct {
	Phases []*ListWafPhasesResponseBodyPhases `json:"Phases,omitempty" xml:"Phases,omitempty" type:"Repeated"`
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListWafPhasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWafPhasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWafPhasesResponseBody) GetPhases() []*ListWafPhasesResponseBodyPhases {
	return s.Phases
}

func (s *ListWafPhasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWafPhasesResponseBody) SetPhases(v []*ListWafPhasesResponseBodyPhases) *ListWafPhasesResponseBody {
	s.Phases = v
	return s
}

func (s *ListWafPhasesResponseBody) SetRequestId(v string) *ListWafPhasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWafPhasesResponseBody) Validate() error {
	if s.Phases != nil {
		for _, item := range s.Phases {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWafPhasesResponseBodyPhases struct {
	Phase    *string                                    `json:"Phase,omitempty" xml:"Phase,omitempty"`
	Rulesets []*ListWafPhasesResponseBodyPhasesRulesets `json:"Rulesets,omitempty" xml:"Rulesets,omitempty" type:"Repeated"`
}

func (s ListWafPhasesResponseBodyPhases) String() string {
	return dara.Prettify(s)
}

func (s ListWafPhasesResponseBodyPhases) GoString() string {
	return s.String()
}

func (s *ListWafPhasesResponseBodyPhases) GetPhase() *string {
	return s.Phase
}

func (s *ListWafPhasesResponseBodyPhases) GetRulesets() []*ListWafPhasesResponseBodyPhasesRulesets {
	return s.Rulesets
}

func (s *ListWafPhasesResponseBodyPhases) SetPhase(v string) *ListWafPhasesResponseBodyPhases {
	s.Phase = &v
	return s
}

func (s *ListWafPhasesResponseBodyPhases) SetRulesets(v []*ListWafPhasesResponseBodyPhasesRulesets) *ListWafPhasesResponseBodyPhases {
	s.Rulesets = v
	return s
}

func (s *ListWafPhasesResponseBodyPhases) Validate() error {
	if s.Rulesets != nil {
		for _, item := range s.Rulesets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWafPhasesResponseBodyPhasesRulesets struct {
	Id     *int64              `json:"Id,omitempty" xml:"Id,omitempty"`
	Name   *string             `json:"Name,omitempty" xml:"Name,omitempty"`
	Rules  []*WafRuleConfig    `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	Shared *WafBatchRuleShared `json:"Shared,omitempty" xml:"Shared,omitempty"`
}

func (s ListWafPhasesResponseBodyPhasesRulesets) String() string {
	return dara.Prettify(s)
}

func (s ListWafPhasesResponseBodyPhasesRulesets) GoString() string {
	return s.String()
}

func (s *ListWafPhasesResponseBodyPhasesRulesets) GetId() *int64 {
	return s.Id
}

func (s *ListWafPhasesResponseBodyPhasesRulesets) GetName() *string {
	return s.Name
}

func (s *ListWafPhasesResponseBodyPhasesRulesets) GetRules() []*WafRuleConfig {
	return s.Rules
}

func (s *ListWafPhasesResponseBodyPhasesRulesets) GetShared() *WafBatchRuleShared {
	return s.Shared
}

func (s *ListWafPhasesResponseBodyPhasesRulesets) SetId(v int64) *ListWafPhasesResponseBodyPhasesRulesets {
	s.Id = &v
	return s
}

func (s *ListWafPhasesResponseBodyPhasesRulesets) SetName(v string) *ListWafPhasesResponseBodyPhasesRulesets {
	s.Name = &v
	return s
}

func (s *ListWafPhasesResponseBodyPhasesRulesets) SetRules(v []*WafRuleConfig) *ListWafPhasesResponseBodyPhasesRulesets {
	s.Rules = v
	return s
}

func (s *ListWafPhasesResponseBodyPhasesRulesets) SetShared(v *WafBatchRuleShared) *ListWafPhasesResponseBodyPhasesRulesets {
	s.Shared = v
	return s
}

func (s *ListWafPhasesResponseBodyPhasesRulesets) Validate() error {
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
