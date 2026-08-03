// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAICoachDebugRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataId(v string) *SubmitAICoachDebugRequest
	GetDataId() *string
	SetDataType(v int64) *SubmitAICoachDebugRequest
	GetDataType() *int64
	SetDeductionRule(v *SubmitAICoachDebugRequestDeductionRule) *SubmitAICoachDebugRequest
	GetDeductionRule() *SubmitAICoachDebugRequestDeductionRule
	SetDialogueList(v []*SubmitAICoachDebugRequestDialogueList) *SubmitAICoachDebugRequest
	GetDialogueList() []*SubmitAICoachDebugRequestDialogueList
	SetExpressiveness(v *SubmitAICoachDebugRequestExpressiveness) *SubmitAICoachDebugRequest
	GetExpressiveness() *SubmitAICoachDebugRequestExpressiveness
	SetPoint(v *SubmitAICoachDebugRequestPoint) *SubmitAICoachDebugRequest
	GetPoint() *SubmitAICoachDebugRequestPoint
}

type SubmitAICoachDebugRequest struct {
	DataId         *string                                  `json:"dataId,omitempty" xml:"dataId,omitempty"`
	DataType       *int64                                   `json:"dataType,omitempty" xml:"dataType,omitempty"`
	DeductionRule  *SubmitAICoachDebugRequestDeductionRule  `json:"deductionRule,omitempty" xml:"deductionRule,omitempty" type:"Struct"`
	DialogueList   []*SubmitAICoachDebugRequestDialogueList `json:"dialogueList,omitempty" xml:"dialogueList,omitempty" type:"Repeated"`
	Expressiveness *SubmitAICoachDebugRequestExpressiveness `json:"expressiveness,omitempty" xml:"expressiveness,omitempty" type:"Struct"`
	Point          *SubmitAICoachDebugRequestPoint          `json:"point,omitempty" xml:"point,omitempty" type:"Struct"`
}

func (s SubmitAICoachDebugRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAICoachDebugRequest) GoString() string {
	return s.String()
}

func (s *SubmitAICoachDebugRequest) GetDataId() *string {
	return s.DataId
}

func (s *SubmitAICoachDebugRequest) GetDataType() *int64 {
	return s.DataType
}

func (s *SubmitAICoachDebugRequest) GetDeductionRule() *SubmitAICoachDebugRequestDeductionRule {
	return s.DeductionRule
}

func (s *SubmitAICoachDebugRequest) GetDialogueList() []*SubmitAICoachDebugRequestDialogueList {
	return s.DialogueList
}

func (s *SubmitAICoachDebugRequest) GetExpressiveness() *SubmitAICoachDebugRequestExpressiveness {
	return s.Expressiveness
}

func (s *SubmitAICoachDebugRequest) GetPoint() *SubmitAICoachDebugRequestPoint {
	return s.Point
}

func (s *SubmitAICoachDebugRequest) SetDataId(v string) *SubmitAICoachDebugRequest {
	s.DataId = &v
	return s
}

func (s *SubmitAICoachDebugRequest) SetDataType(v int64) *SubmitAICoachDebugRequest {
	s.DataType = &v
	return s
}

func (s *SubmitAICoachDebugRequest) SetDeductionRule(v *SubmitAICoachDebugRequestDeductionRule) *SubmitAICoachDebugRequest {
	s.DeductionRule = v
	return s
}

func (s *SubmitAICoachDebugRequest) SetDialogueList(v []*SubmitAICoachDebugRequestDialogueList) *SubmitAICoachDebugRequest {
	s.DialogueList = v
	return s
}

func (s *SubmitAICoachDebugRequest) SetExpressiveness(v *SubmitAICoachDebugRequestExpressiveness) *SubmitAICoachDebugRequest {
	s.Expressiveness = v
	return s
}

func (s *SubmitAICoachDebugRequest) SetPoint(v *SubmitAICoachDebugRequestPoint) *SubmitAICoachDebugRequest {
	s.Point = v
	return s
}

func (s *SubmitAICoachDebugRequest) Validate() error {
	if s.DeductionRule != nil {
		if err := s.DeductionRule.Validate(); err != nil {
			return err
		}
	}
	if s.DialogueList != nil {
		for _, item := range s.DialogueList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Expressiveness != nil {
		if err := s.Expressiveness.Validate(); err != nil {
			return err
		}
	}
	if s.Point != nil {
		if err := s.Point.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitAICoachDebugRequestDeductionRule struct {
	DeductionRuleId *string   `json:"deductionRuleId,omitempty" xml:"deductionRuleId,omitempty"`
	Description     *string   `json:"description,omitempty" xml:"description,omitempty"`
	PunishmentTypes []*string `json:"punishmentTypes,omitempty" xml:"punishmentTypes,omitempty" type:"Repeated"`
	RuleValue       *string   `json:"ruleValue,omitempty" xml:"ruleValue,omitempty"`
	Weight          *int32    `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s SubmitAICoachDebugRequestDeductionRule) String() string {
	return dara.Prettify(s)
}

func (s SubmitAICoachDebugRequestDeductionRule) GoString() string {
	return s.String()
}

func (s *SubmitAICoachDebugRequestDeductionRule) GetDeductionRuleId() *string {
	return s.DeductionRuleId
}

func (s *SubmitAICoachDebugRequestDeductionRule) GetDescription() *string {
	return s.Description
}

func (s *SubmitAICoachDebugRequestDeductionRule) GetPunishmentTypes() []*string {
	return s.PunishmentTypes
}

func (s *SubmitAICoachDebugRequestDeductionRule) GetRuleValue() *string {
	return s.RuleValue
}

func (s *SubmitAICoachDebugRequestDeductionRule) GetWeight() *int32 {
	return s.Weight
}

func (s *SubmitAICoachDebugRequestDeductionRule) SetDeductionRuleId(v string) *SubmitAICoachDebugRequestDeductionRule {
	s.DeductionRuleId = &v
	return s
}

func (s *SubmitAICoachDebugRequestDeductionRule) SetDescription(v string) *SubmitAICoachDebugRequestDeductionRule {
	s.Description = &v
	return s
}

func (s *SubmitAICoachDebugRequestDeductionRule) SetPunishmentTypes(v []*string) *SubmitAICoachDebugRequestDeductionRule {
	s.PunishmentTypes = v
	return s
}

func (s *SubmitAICoachDebugRequestDeductionRule) SetRuleValue(v string) *SubmitAICoachDebugRequestDeductionRule {
	s.RuleValue = &v
	return s
}

func (s *SubmitAICoachDebugRequestDeductionRule) SetWeight(v int32) *SubmitAICoachDebugRequestDeductionRule {
	s.Weight = &v
	return s
}

func (s *SubmitAICoachDebugRequestDeductionRule) Validate() error {
	return dara.Validate(s)
}

type SubmitAICoachDebugRequestDialogueList struct {
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	Role    *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s SubmitAICoachDebugRequestDialogueList) String() string {
	return dara.Prettify(s)
}

func (s SubmitAICoachDebugRequestDialogueList) GoString() string {
	return s.String()
}

func (s *SubmitAICoachDebugRequestDialogueList) GetMessage() *string {
	return s.Message
}

func (s *SubmitAICoachDebugRequestDialogueList) GetRole() *string {
	return s.Role
}

func (s *SubmitAICoachDebugRequestDialogueList) SetMessage(v string) *SubmitAICoachDebugRequestDialogueList {
	s.Message = &v
	return s
}

func (s *SubmitAICoachDebugRequestDialogueList) SetRole(v string) *SubmitAICoachDebugRequestDialogueList {
	s.Role = &v
	return s
}

func (s *SubmitAICoachDebugRequestDialogueList) Validate() error {
	return dara.Validate(s)
}

type SubmitAICoachDebugRequestExpressiveness struct {
	Desc             *string `json:"desc,omitempty" xml:"desc,omitempty"`
	ExpressivenessId *string `json:"expressivenessId,omitempty" xml:"expressivenessId,omitempty"`
	Name             *string `json:"name,omitempty" xml:"name,omitempty"`
	Rule             *string `json:"rule,omitempty" xml:"rule,omitempty"`
}

func (s SubmitAICoachDebugRequestExpressiveness) String() string {
	return dara.Prettify(s)
}

func (s SubmitAICoachDebugRequestExpressiveness) GoString() string {
	return s.String()
}

func (s *SubmitAICoachDebugRequestExpressiveness) GetDesc() *string {
	return s.Desc
}

func (s *SubmitAICoachDebugRequestExpressiveness) GetExpressivenessId() *string {
	return s.ExpressivenessId
}

func (s *SubmitAICoachDebugRequestExpressiveness) GetName() *string {
	return s.Name
}

func (s *SubmitAICoachDebugRequestExpressiveness) GetRule() *string {
	return s.Rule
}

func (s *SubmitAICoachDebugRequestExpressiveness) SetDesc(v string) *SubmitAICoachDebugRequestExpressiveness {
	s.Desc = &v
	return s
}

func (s *SubmitAICoachDebugRequestExpressiveness) SetExpressivenessId(v string) *SubmitAICoachDebugRequestExpressiveness {
	s.ExpressivenessId = &v
	return s
}

func (s *SubmitAICoachDebugRequestExpressiveness) SetName(v string) *SubmitAICoachDebugRequestExpressiveness {
	s.Name = &v
	return s
}

func (s *SubmitAICoachDebugRequestExpressiveness) SetRule(v string) *SubmitAICoachDebugRequestExpressiveness {
	s.Rule = &v
	return s
}

func (s *SubmitAICoachDebugRequestExpressiveness) Validate() error {
	return dara.Validate(s)
}

type SubmitAICoachDebugRequestPoint struct {
	AnswerList     []*SubmitAICoachDebugRequestPointAnswerList `json:"answerList,omitempty" xml:"answerList,omitempty" type:"Repeated"`
	KnowledgeList  []*string                                   `json:"knowledgeList,omitempty" xml:"knowledgeList,omitempty" type:"Repeated"`
	Name           *string                                     `json:"name,omitempty" xml:"name,omitempty"`
	QuestionSample *string                                     `json:"questionSample,omitempty" xml:"questionSample,omitempty"`
	Weight         *int64                                      `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s SubmitAICoachDebugRequestPoint) String() string {
	return dara.Prettify(s)
}

func (s SubmitAICoachDebugRequestPoint) GoString() string {
	return s.String()
}

func (s *SubmitAICoachDebugRequestPoint) GetAnswerList() []*SubmitAICoachDebugRequestPointAnswerList {
	return s.AnswerList
}

func (s *SubmitAICoachDebugRequestPoint) GetKnowledgeList() []*string {
	return s.KnowledgeList
}

func (s *SubmitAICoachDebugRequestPoint) GetName() *string {
	return s.Name
}

func (s *SubmitAICoachDebugRequestPoint) GetQuestionSample() *string {
	return s.QuestionSample
}

func (s *SubmitAICoachDebugRequestPoint) GetWeight() *int64 {
	return s.Weight
}

func (s *SubmitAICoachDebugRequestPoint) SetAnswerList(v []*SubmitAICoachDebugRequestPointAnswerList) *SubmitAICoachDebugRequestPoint {
	s.AnswerList = v
	return s
}

func (s *SubmitAICoachDebugRequestPoint) SetKnowledgeList(v []*string) *SubmitAICoachDebugRequestPoint {
	s.KnowledgeList = v
	return s
}

func (s *SubmitAICoachDebugRequestPoint) SetName(v string) *SubmitAICoachDebugRequestPoint {
	s.Name = &v
	return s
}

func (s *SubmitAICoachDebugRequestPoint) SetQuestionSample(v string) *SubmitAICoachDebugRequestPoint {
	s.QuestionSample = &v
	return s
}

func (s *SubmitAICoachDebugRequestPoint) SetWeight(v int64) *SubmitAICoachDebugRequestPoint {
	s.Weight = &v
	return s
}

func (s *SubmitAICoachDebugRequestPoint) Validate() error {
	if s.AnswerList != nil {
		for _, item := range s.AnswerList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubmitAICoachDebugRequestPointAnswerList struct {
	AnswerValues   []*SubmitAICoachDebugRequestPointAnswerListAnswerValues `json:"answerValues,omitempty" xml:"answerValues,omitempty" type:"Repeated"`
	EnabledKeyword *bool                                                   `json:"enabledKeyword,omitempty" xml:"enabledKeyword,omitempty"`
	NameList       []*string                                               `json:"nameList,omitempty" xml:"nameList,omitempty" type:"Repeated"`
	Operators      *string                                                 `json:"operators,omitempty" xml:"operators,omitempty"`
	Parameters     []*SubmitAICoachDebugRequestPointAnswerListParameters   `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Repeated"`
	Score          *int64                                                  `json:"score,omitempty" xml:"score,omitempty"`
	Type           *string                                                 `json:"type,omitempty" xml:"type,omitempty"`
	Weight         *int64                                                  `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s SubmitAICoachDebugRequestPointAnswerList) String() string {
	return dara.Prettify(s)
}

func (s SubmitAICoachDebugRequestPointAnswerList) GoString() string {
	return s.String()
}

func (s *SubmitAICoachDebugRequestPointAnswerList) GetAnswerValues() []*SubmitAICoachDebugRequestPointAnswerListAnswerValues {
	return s.AnswerValues
}

func (s *SubmitAICoachDebugRequestPointAnswerList) GetEnabledKeyword() *bool {
	return s.EnabledKeyword
}

func (s *SubmitAICoachDebugRequestPointAnswerList) GetNameList() []*string {
	return s.NameList
}

func (s *SubmitAICoachDebugRequestPointAnswerList) GetOperators() *string {
	return s.Operators
}

func (s *SubmitAICoachDebugRequestPointAnswerList) GetParameters() []*SubmitAICoachDebugRequestPointAnswerListParameters {
	return s.Parameters
}

func (s *SubmitAICoachDebugRequestPointAnswerList) GetScore() *int64 {
	return s.Score
}

func (s *SubmitAICoachDebugRequestPointAnswerList) GetType() *string {
	return s.Type
}

func (s *SubmitAICoachDebugRequestPointAnswerList) GetWeight() *int64 {
	return s.Weight
}

func (s *SubmitAICoachDebugRequestPointAnswerList) SetAnswerValues(v []*SubmitAICoachDebugRequestPointAnswerListAnswerValues) *SubmitAICoachDebugRequestPointAnswerList {
	s.AnswerValues = v
	return s
}

func (s *SubmitAICoachDebugRequestPointAnswerList) SetEnabledKeyword(v bool) *SubmitAICoachDebugRequestPointAnswerList {
	s.EnabledKeyword = &v
	return s
}

func (s *SubmitAICoachDebugRequestPointAnswerList) SetNameList(v []*string) *SubmitAICoachDebugRequestPointAnswerList {
	s.NameList = v
	return s
}

func (s *SubmitAICoachDebugRequestPointAnswerList) SetOperators(v string) *SubmitAICoachDebugRequestPointAnswerList {
	s.Operators = &v
	return s
}

func (s *SubmitAICoachDebugRequestPointAnswerList) SetParameters(v []*SubmitAICoachDebugRequestPointAnswerListParameters) *SubmitAICoachDebugRequestPointAnswerList {
	s.Parameters = v
	return s
}

func (s *SubmitAICoachDebugRequestPointAnswerList) SetScore(v int64) *SubmitAICoachDebugRequestPointAnswerList {
	s.Score = &v
	return s
}

func (s *SubmitAICoachDebugRequestPointAnswerList) SetType(v string) *SubmitAICoachDebugRequestPointAnswerList {
	s.Type = &v
	return s
}

func (s *SubmitAICoachDebugRequestPointAnswerList) SetWeight(v int64) *SubmitAICoachDebugRequestPointAnswerList {
	s.Weight = &v
	return s
}

func (s *SubmitAICoachDebugRequestPointAnswerList) Validate() error {
	if s.AnswerValues != nil {
		for _, item := range s.AnswerValues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Parameters != nil {
		for _, item := range s.Parameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubmitAICoachDebugRequestPointAnswerListAnswerValues struct {
	AnswerName    *string                                                              `json:"answerName,omitempty" xml:"answerName,omitempty"`
	AnswerWeight  *int64                                                               `json:"answerWeight,omitempty" xml:"answerWeight,omitempty"`
	KeywordValues []*SubmitAICoachDebugRequestPointAnswerListAnswerValuesKeywordValues `json:"keywordValues,omitempty" xml:"keywordValues,omitempty" type:"Repeated"`
	KeywordWeight *int64                                                               `json:"keywordWeight,omitempty" xml:"keywordWeight,omitempty"`
	ScoringRules  []*SubmitAICoachDebugRequestPointAnswerListAnswerValuesScoringRules  `json:"scoringRules,omitempty" xml:"scoringRules,omitempty" type:"Repeated"`
}

func (s SubmitAICoachDebugRequestPointAnswerListAnswerValues) String() string {
	return dara.Prettify(s)
}

func (s SubmitAICoachDebugRequestPointAnswerListAnswerValues) GoString() string {
	return s.String()
}

func (s *SubmitAICoachDebugRequestPointAnswerListAnswerValues) GetAnswerName() *string {
	return s.AnswerName
}

func (s *SubmitAICoachDebugRequestPointAnswerListAnswerValues) GetAnswerWeight() *int64 {
	return s.AnswerWeight
}

func (s *SubmitAICoachDebugRequestPointAnswerListAnswerValues) GetKeywordValues() []*SubmitAICoachDebugRequestPointAnswerListAnswerValuesKeywordValues {
	return s.KeywordValues
}

func (s *SubmitAICoachDebugRequestPointAnswerListAnswerValues) GetKeywordWeight() *int64 {
	return s.KeywordWeight
}

func (s *SubmitAICoachDebugRequestPointAnswerListAnswerValues) GetScoringRules() []*SubmitAICoachDebugRequestPointAnswerListAnswerValuesScoringRules {
	return s.ScoringRules
}

func (s *SubmitAICoachDebugRequestPointAnswerListAnswerValues) SetAnswerName(v string) *SubmitAICoachDebugRequestPointAnswerListAnswerValues {
	s.AnswerName = &v
	return s
}

func (s *SubmitAICoachDebugRequestPointAnswerListAnswerValues) SetAnswerWeight(v int64) *SubmitAICoachDebugRequestPointAnswerListAnswerValues {
	s.AnswerWeight = &v
	return s
}

func (s *SubmitAICoachDebugRequestPointAnswerListAnswerValues) SetKeywordValues(v []*SubmitAICoachDebugRequestPointAnswerListAnswerValuesKeywordValues) *SubmitAICoachDebugRequestPointAnswerListAnswerValues {
	s.KeywordValues = v
	return s
}

func (s *SubmitAICoachDebugRequestPointAnswerListAnswerValues) SetKeywordWeight(v int64) *SubmitAICoachDebugRequestPointAnswerListAnswerValues {
	s.KeywordWeight = &v
	return s
}

func (s *SubmitAICoachDebugRequestPointAnswerListAnswerValues) SetScoringRules(v []*SubmitAICoachDebugRequestPointAnswerListAnswerValuesScoringRules) *SubmitAICoachDebugRequestPointAnswerListAnswerValues {
	s.ScoringRules = v
	return s
}

func (s *SubmitAICoachDebugRequestPointAnswerListAnswerValues) Validate() error {
	if s.KeywordValues != nil {
		for _, item := range s.KeywordValues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ScoringRules != nil {
		for _, item := range s.ScoringRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubmitAICoachDebugRequestPointAnswerListAnswerValuesKeywordValues struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	Weight *int64  `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s SubmitAICoachDebugRequestPointAnswerListAnswerValuesKeywordValues) String() string {
	return dara.Prettify(s)
}

func (s SubmitAICoachDebugRequestPointAnswerListAnswerValuesKeywordValues) GoString() string {
	return s.String()
}

func (s *SubmitAICoachDebugRequestPointAnswerListAnswerValuesKeywordValues) GetName() *string {
	return s.Name
}

func (s *SubmitAICoachDebugRequestPointAnswerListAnswerValuesKeywordValues) GetWeight() *int64 {
	return s.Weight
}

func (s *SubmitAICoachDebugRequestPointAnswerListAnswerValuesKeywordValues) SetName(v string) *SubmitAICoachDebugRequestPointAnswerListAnswerValuesKeywordValues {
	s.Name = &v
	return s
}

func (s *SubmitAICoachDebugRequestPointAnswerListAnswerValuesKeywordValues) SetWeight(v int64) *SubmitAICoachDebugRequestPointAnswerListAnswerValuesKeywordValues {
	s.Weight = &v
	return s
}

func (s *SubmitAICoachDebugRequestPointAnswerListAnswerValuesKeywordValues) Validate() error {
	return dara.Validate(s)
}

type SubmitAICoachDebugRequestPointAnswerListAnswerValuesScoringRules struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s SubmitAICoachDebugRequestPointAnswerListAnswerValuesScoringRules) String() string {
	return dara.Prettify(s)
}

func (s SubmitAICoachDebugRequestPointAnswerListAnswerValuesScoringRules) GoString() string {
	return s.String()
}

func (s *SubmitAICoachDebugRequestPointAnswerListAnswerValuesScoringRules) GetName() *string {
	return s.Name
}

func (s *SubmitAICoachDebugRequestPointAnswerListAnswerValuesScoringRules) SetName(v string) *SubmitAICoachDebugRequestPointAnswerListAnswerValuesScoringRules {
	s.Name = &v
	return s
}

func (s *SubmitAICoachDebugRequestPointAnswerListAnswerValuesScoringRules) Validate() error {
	return dara.Validate(s)
}

type SubmitAICoachDebugRequestPointAnswerListParameters struct {
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s SubmitAICoachDebugRequestPointAnswerListParameters) String() string {
	return dara.Prettify(s)
}

func (s SubmitAICoachDebugRequestPointAnswerListParameters) GoString() string {
	return s.String()
}

func (s *SubmitAICoachDebugRequestPointAnswerListParameters) GetName() *string {
	return s.Name
}

func (s *SubmitAICoachDebugRequestPointAnswerListParameters) GetValue() *string {
	return s.Value
}

func (s *SubmitAICoachDebugRequestPointAnswerListParameters) SetName(v string) *SubmitAICoachDebugRequestPointAnswerListParameters {
	s.Name = &v
	return s
}

func (s *SubmitAICoachDebugRequestPointAnswerListParameters) SetValue(v string) *SubmitAICoachDebugRequestPointAnswerListParameters {
	s.Value = &v
	return s
}

func (s *SubmitAICoachDebugRequestPointAnswerListParameters) Validate() error {
	return dara.Validate(s)
}
