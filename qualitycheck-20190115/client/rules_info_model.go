// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRulesInfo interface {
	dara.Model
	String() string
	GoString() string
	SetConditions(v []*ConditionBasicInfo) *RulesInfo
	GetConditions() []*ConditionBasicInfo
	SetCount(v int32) *RulesInfo
	GetCount() *int32
	SetDialogues(v []*RuleTestDialogue) *RulesInfo
	GetDialogues() []*RuleTestDialogue
	SetPageNumber(v int32) *RulesInfo
	GetPageNumber() *int32
	SetPageSize(v int32) *RulesInfo
	GetPageSize() *int32
	SetRules(v []*RuleInfo) *RulesInfo
	GetRules() []*RuleInfo
}

type RulesInfo struct {
	// if can be null:
	// true
	Conditions []*ConditionBasicInfo `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	Count      *int32                `json:"Count,omitempty" xml:"Count,omitempty"`
	// if can be null:
	// true
	Dialogues  []*RuleTestDialogue `json:"Dialogues,omitempty" xml:"Dialogues,omitempty" type:"Repeated"`
	PageNumber *int32              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Rules      []*RuleInfo         `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s RulesInfo) String() string {
	return dara.Prettify(s)
}

func (s RulesInfo) GoString() string {
	return s.String()
}

func (s *RulesInfo) GetConditions() []*ConditionBasicInfo {
	return s.Conditions
}

func (s *RulesInfo) GetCount() *int32 {
	return s.Count
}

func (s *RulesInfo) GetDialogues() []*RuleTestDialogue {
	return s.Dialogues
}

func (s *RulesInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *RulesInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *RulesInfo) GetRules() []*RuleInfo {
	return s.Rules
}

func (s *RulesInfo) SetConditions(v []*ConditionBasicInfo) *RulesInfo {
	s.Conditions = v
	return s
}

func (s *RulesInfo) SetCount(v int32) *RulesInfo {
	s.Count = &v
	return s
}

func (s *RulesInfo) SetDialogues(v []*RuleTestDialogue) *RulesInfo {
	s.Dialogues = v
	return s
}

func (s *RulesInfo) SetPageNumber(v int32) *RulesInfo {
	s.PageNumber = &v
	return s
}

func (s *RulesInfo) SetPageSize(v int32) *RulesInfo {
	s.PageSize = &v
	return s
}

func (s *RulesInfo) SetRules(v []*RuleInfo) *RulesInfo {
	s.Rules = v
	return s
}

func (s *RulesInfo) Validate() error {
	return dara.Validate(s)
}
