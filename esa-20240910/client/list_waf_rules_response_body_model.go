// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWafRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceUsage(v int64) *ListWafRulesResponseBody
	GetInstanceUsage() *int64
	SetPageNumber(v int32) *ListWafRulesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListWafRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListWafRulesResponseBody
	GetRequestId() *string
	SetRules(v []*ListWafRulesResponseBodyRules) *ListWafRulesResponseBody
	GetRules() []*ListWafRulesResponseBodyRules
	SetSiteUsage(v int64) *ListWafRulesResponseBody
	GetSiteUsage() *int64
	SetTotalCount(v int64) *ListWafRulesResponseBody
	GetTotalCount() *int64
}

type ListWafRulesResponseBody struct {
	// Number of rules used in this WAF phase for the corresponding instance of the site.
	//
	// example:
	//
	// 10
	InstanceUsage *int64 `json:"InstanceUsage,omitempty" xml:"InstanceUsage,omitempty"`
	// Page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned list of rules.
	Rules []*ListWafRulesResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// Site usage.
	//
	// example:
	//
	// 5
	SiteUsage *int64 `json:"SiteUsage,omitempty" xml:"SiteUsage,omitempty"`
	// Total number of rules after filtering.
	//
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListWafRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWafRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWafRulesResponseBody) GetInstanceUsage() *int64 {
	return s.InstanceUsage
}

func (s *ListWafRulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListWafRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWafRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWafRulesResponseBody) GetRules() []*ListWafRulesResponseBodyRules {
	return s.Rules
}

func (s *ListWafRulesResponseBody) GetSiteUsage() *int64 {
	return s.SiteUsage
}

func (s *ListWafRulesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListWafRulesResponseBody) SetInstanceUsage(v int64) *ListWafRulesResponseBody {
	s.InstanceUsage = &v
	return s
}

func (s *ListWafRulesResponseBody) SetPageNumber(v int32) *ListWafRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListWafRulesResponseBody) SetPageSize(v int32) *ListWafRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListWafRulesResponseBody) SetRequestId(v string) *ListWafRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWafRulesResponseBody) SetRules(v []*ListWafRulesResponseBodyRules) *ListWafRulesResponseBody {
	s.Rules = v
	return s
}

func (s *ListWafRulesResponseBody) SetSiteUsage(v int64) *ListWafRulesResponseBody {
	s.SiteUsage = &v
	return s
}

func (s *ListWafRulesResponseBody) SetTotalCount(v int64) *ListWafRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWafRulesResponseBody) Validate() error {
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

type ListWafRulesResponseBodyRules struct {
	// The action corresponding to the rule.
	//
	// example:
	//
	// deny
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// List of statistical objects for frequency control rules.
	CharacteristicsFields []*string `json:"CharacteristicsFields,omitempty" xml:"CharacteristicsFields,omitempty" type:"Repeated"`
	// Rule configuration.
	Config *WafRuleConfig `json:"Config,omitempty" xml:"Config,omitempty"`
	// List of fields for rule matching
	Fields []*string `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// Rule ID.
	//
	// example:
	//
	// 20000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Rule name.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// WAF phase.
	//
	// example:
	//
	// http_custom
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// Position order of the rule in the corresponding ruleset.
	//
	// example:
	//
	// 1
	Position *int64 `json:"Position,omitempty" xml:"Position,omitempty"`
	// Ruleset ID.
	//
	// example:
	//
	// 10000001
	RulesetId *int64 `json:"RulesetId,omitempty" xml:"RulesetId,omitempty"`
	// Skip attribute for whitelist rules.
	//
	// example:
	//
	// part
	Skip *string `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// Rule status.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// List of WAF phases to be skipped by whitelist rules.
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// Configuration for the effective time of the rule.
	Timer *WafTimer `json:"Timer,omitempty" xml:"Timer,omitempty"`
	// Rule type.
	//
	// example:
	//
	// http_custom
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Modification time.
	//
	// example:
	//
	// 2024-01-01T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListWafRulesResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s ListWafRulesResponseBodyRules) GoString() string {
	return s.String()
}

func (s *ListWafRulesResponseBodyRules) GetAction() *string {
	return s.Action
}

func (s *ListWafRulesResponseBodyRules) GetCharacteristicsFields() []*string {
	return s.CharacteristicsFields
}

func (s *ListWafRulesResponseBodyRules) GetConfig() *WafRuleConfig {
	return s.Config
}

func (s *ListWafRulesResponseBodyRules) GetFields() []*string {
	return s.Fields
}

func (s *ListWafRulesResponseBodyRules) GetId() *int64 {
	return s.Id
}

func (s *ListWafRulesResponseBodyRules) GetName() *string {
	return s.Name
}

func (s *ListWafRulesResponseBodyRules) GetPhase() *string {
	return s.Phase
}

func (s *ListWafRulesResponseBodyRules) GetPosition() *int64 {
	return s.Position
}

func (s *ListWafRulesResponseBodyRules) GetRulesetId() *int64 {
	return s.RulesetId
}

func (s *ListWafRulesResponseBodyRules) GetSkip() *string {
	return s.Skip
}

func (s *ListWafRulesResponseBodyRules) GetStatus() *string {
	return s.Status
}

func (s *ListWafRulesResponseBodyRules) GetTags() []*string {
	return s.Tags
}

func (s *ListWafRulesResponseBodyRules) GetTimer() *WafTimer {
	return s.Timer
}

func (s *ListWafRulesResponseBodyRules) GetType() *string {
	return s.Type
}

func (s *ListWafRulesResponseBodyRules) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListWafRulesResponseBodyRules) SetAction(v string) *ListWafRulesResponseBodyRules {
	s.Action = &v
	return s
}

func (s *ListWafRulesResponseBodyRules) SetCharacteristicsFields(v []*string) *ListWafRulesResponseBodyRules {
	s.CharacteristicsFields = v
	return s
}

func (s *ListWafRulesResponseBodyRules) SetConfig(v *WafRuleConfig) *ListWafRulesResponseBodyRules {
	s.Config = v
	return s
}

func (s *ListWafRulesResponseBodyRules) SetFields(v []*string) *ListWafRulesResponseBodyRules {
	s.Fields = v
	return s
}

func (s *ListWafRulesResponseBodyRules) SetId(v int64) *ListWafRulesResponseBodyRules {
	s.Id = &v
	return s
}

func (s *ListWafRulesResponseBodyRules) SetName(v string) *ListWafRulesResponseBodyRules {
	s.Name = &v
	return s
}

func (s *ListWafRulesResponseBodyRules) SetPhase(v string) *ListWafRulesResponseBodyRules {
	s.Phase = &v
	return s
}

func (s *ListWafRulesResponseBodyRules) SetPosition(v int64) *ListWafRulesResponseBodyRules {
	s.Position = &v
	return s
}

func (s *ListWafRulesResponseBodyRules) SetRulesetId(v int64) *ListWafRulesResponseBodyRules {
	s.RulesetId = &v
	return s
}

func (s *ListWafRulesResponseBodyRules) SetSkip(v string) *ListWafRulesResponseBodyRules {
	s.Skip = &v
	return s
}

func (s *ListWafRulesResponseBodyRules) SetStatus(v string) *ListWafRulesResponseBodyRules {
	s.Status = &v
	return s
}

func (s *ListWafRulesResponseBodyRules) SetTags(v []*string) *ListWafRulesResponseBodyRules {
	s.Tags = v
	return s
}

func (s *ListWafRulesResponseBodyRules) SetTimer(v *WafTimer) *ListWafRulesResponseBodyRules {
	s.Timer = v
	return s
}

func (s *ListWafRulesResponseBodyRules) SetType(v string) *ListWafRulesResponseBodyRules {
	s.Type = &v
	return s
}

func (s *ListWafRulesResponseBodyRules) SetUpdateTime(v string) *ListWafRulesResponseBodyRules {
	s.UpdateTime = &v
	return s
}

func (s *ListWafRulesResponseBodyRules) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	if s.Timer != nil {
		if err := s.Timer.Validate(); err != nil {
			return err
		}
	}
	return nil
}
