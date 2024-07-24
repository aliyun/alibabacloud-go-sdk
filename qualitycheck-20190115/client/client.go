// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type BusinessCategoryBasicInfo struct {
	Bid         *int32  `json:"Bid,omitempty" xml:"Bid,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OriginalId  *int64  `json:"OriginalId,omitempty" xml:"OriginalId,omitempty"`
	ServiceType *int32  `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s BusinessCategoryBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s BusinessCategoryBasicInfo) GoString() string {
	return s.String()
}

func (s *BusinessCategoryBasicInfo) SetBid(v int32) *BusinessCategoryBasicInfo {
	s.Bid = &v
	return s
}

func (s *BusinessCategoryBasicInfo) SetName(v string) *BusinessCategoryBasicInfo {
	s.Name = &v
	return s
}

func (s *BusinessCategoryBasicInfo) SetOriginalId(v int64) *BusinessCategoryBasicInfo {
	s.OriginalId = &v
	return s
}

func (s *BusinessCategoryBasicInfo) SetServiceType(v int32) *BusinessCategoryBasicInfo {
	s.ServiceType = &v
	return s
}

type ConditionBasicInfo struct {
	CheckRange *ConditionBasicInfoCheckRange `json:"Check_range,omitempty" xml:"Check_range,omitempty" type:"Struct"`
	Cid        *string                       `json:"Cid,omitempty" xml:"Cid,omitempty"`
	Exclusion  *int32                        `json:"Exclusion,omitempty" xml:"Exclusion,omitempty"`
	Id         *int64                        `json:"Id,omitempty" xml:"Id,omitempty"`
	Lambda     *string                       `json:"Lambda,omitempty" xml:"Lambda,omitempty"`
	Name       *string                       `json:"Name,omitempty" xml:"Name,omitempty"`
	Operators  []*OperatorBasicInfo          `json:"Operators,omitempty" xml:"Operators,omitempty" type:"Repeated"`
	Rid        *string                       `json:"Rid,omitempty" xml:"Rid,omitempty"`
	UserGroup  *string                       `json:"UserGroup,omitempty" xml:"UserGroup,omitempty"`
}

func (s ConditionBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s ConditionBasicInfo) GoString() string {
	return s.String()
}

func (s *ConditionBasicInfo) SetCheckRange(v *ConditionBasicInfoCheckRange) *ConditionBasicInfo {
	s.CheckRange = v
	return s
}

func (s *ConditionBasicInfo) SetCid(v string) *ConditionBasicInfo {
	s.Cid = &v
	return s
}

func (s *ConditionBasicInfo) SetExclusion(v int32) *ConditionBasicInfo {
	s.Exclusion = &v
	return s
}

func (s *ConditionBasicInfo) SetId(v int64) *ConditionBasicInfo {
	s.Id = &v
	return s
}

func (s *ConditionBasicInfo) SetLambda(v string) *ConditionBasicInfo {
	s.Lambda = &v
	return s
}

func (s *ConditionBasicInfo) SetName(v string) *ConditionBasicInfo {
	s.Name = &v
	return s
}

func (s *ConditionBasicInfo) SetOperators(v []*OperatorBasicInfo) *ConditionBasicInfo {
	s.Operators = v
	return s
}

func (s *ConditionBasicInfo) SetRid(v string) *ConditionBasicInfo {
	s.Rid = &v
	return s
}

func (s *ConditionBasicInfo) SetUserGroup(v string) *ConditionBasicInfo {
	s.UserGroup = &v
	return s
}

type ConditionBasicInfoCheckRange struct {
	Absolute            *bool                               `json:"Absolute,omitempty" xml:"Absolute,omitempty"`
	AllSentencesSatisfy *bool                               `json:"AllSentencesSatisfy,omitempty" xml:"AllSentencesSatisfy,omitempty"`
	Anchor              *ConditionBasicInfoCheckRangeAnchor `json:"Anchor,omitempty" xml:"Anchor,omitempty" type:"Struct"`
	Range               *ConditionBasicInfoCheckRangeRange  `json:"Range,omitempty" xml:"Range,omitempty" type:"Struct"`
	Role                *string                             `json:"Role,omitempty" xml:"Role,omitempty"`
	RoleId              *int32                              `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
}

func (s ConditionBasicInfoCheckRange) String() string {
	return tea.Prettify(s)
}

func (s ConditionBasicInfoCheckRange) GoString() string {
	return s.String()
}

func (s *ConditionBasicInfoCheckRange) SetAbsolute(v bool) *ConditionBasicInfoCheckRange {
	s.Absolute = &v
	return s
}

func (s *ConditionBasicInfoCheckRange) SetAllSentencesSatisfy(v bool) *ConditionBasicInfoCheckRange {
	s.AllSentencesSatisfy = &v
	return s
}

func (s *ConditionBasicInfoCheckRange) SetAnchor(v *ConditionBasicInfoCheckRangeAnchor) *ConditionBasicInfoCheckRange {
	s.Anchor = v
	return s
}

func (s *ConditionBasicInfoCheckRange) SetRange(v *ConditionBasicInfoCheckRangeRange) *ConditionBasicInfoCheckRange {
	s.Range = v
	return s
}

func (s *ConditionBasicInfoCheckRange) SetRole(v string) *ConditionBasicInfoCheckRange {
	s.Role = &v
	return s
}

func (s *ConditionBasicInfoCheckRange) SetRoleId(v int32) *ConditionBasicInfoCheckRange {
	s.RoleId = &v
	return s
}

type ConditionBasicInfoCheckRangeAnchor struct {
	Cid      *string `json:"Cid,omitempty" xml:"Cid,omitempty"`
	HitTime  *int32  `json:"Hit_time,omitempty" xml:"Hit_time,omitempty"`
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s ConditionBasicInfoCheckRangeAnchor) String() string {
	return tea.Prettify(s)
}

func (s ConditionBasicInfoCheckRangeAnchor) GoString() string {
	return s.String()
}

func (s *ConditionBasicInfoCheckRangeAnchor) SetCid(v string) *ConditionBasicInfoCheckRangeAnchor {
	s.Cid = &v
	return s
}

func (s *ConditionBasicInfoCheckRangeAnchor) SetHitTime(v int32) *ConditionBasicInfoCheckRangeAnchor {
	s.HitTime = &v
	return s
}

func (s *ConditionBasicInfoCheckRangeAnchor) SetLocation(v string) *ConditionBasicInfoCheckRangeAnchor {
	s.Location = &v
	return s
}

type ConditionBasicInfoCheckRangeRange struct {
	From *int32 `json:"From,omitempty" xml:"From,omitempty"`
	To   *int32 `json:"To,omitempty" xml:"To,omitempty"`
}

func (s ConditionBasicInfoCheckRangeRange) String() string {
	return tea.Prettify(s)
}

func (s ConditionBasicInfoCheckRangeRange) GoString() string {
	return s.String()
}

func (s *ConditionBasicInfoCheckRangeRange) SetFrom(v int32) *ConditionBasicInfoCheckRangeRange {
	s.From = &v
	return s
}

func (s *ConditionBasicInfoCheckRangeRange) SetTo(v int32) *ConditionBasicInfoCheckRangeRange {
	s.To = &v
	return s
}

type GraphFlowNode struct {
	Conditions    []*ConditionBasicInfo     `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	Content       *string                   `json:"Content,omitempty" xml:"Content,omitempty"`
	Id            *int64                    `json:"Id,omitempty" xml:"Id,omitempty"`
	Index         *int32                    `json:"Index,omitempty" xml:"Index,omitempty"`
	Name          *string                   `json:"Name,omitempty" xml:"Name,omitempty"`
	NextNodes     []*GraphFlowNodeNextNodes `json:"NextNodes,omitempty" xml:"NextNodes,omitempty" type:"Repeated"`
	NodeType      *string                   `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	Properties    *GraphFlowNodeProperties  `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	Rid           *int64                    `json:"Rid,omitempty" xml:"Rid,omitempty"`
	UseConditions *bool                     `json:"UseConditions,omitempty" xml:"UseConditions,omitempty"`
}

func (s GraphFlowNode) String() string {
	return tea.Prettify(s)
}

func (s GraphFlowNode) GoString() string {
	return s.String()
}

func (s *GraphFlowNode) SetConditions(v []*ConditionBasicInfo) *GraphFlowNode {
	s.Conditions = v
	return s
}

func (s *GraphFlowNode) SetContent(v string) *GraphFlowNode {
	s.Content = &v
	return s
}

func (s *GraphFlowNode) SetId(v int64) *GraphFlowNode {
	s.Id = &v
	return s
}

func (s *GraphFlowNode) SetIndex(v int32) *GraphFlowNode {
	s.Index = &v
	return s
}

func (s *GraphFlowNode) SetName(v string) *GraphFlowNode {
	s.Name = &v
	return s
}

func (s *GraphFlowNode) SetNextNodes(v []*GraphFlowNodeNextNodes) *GraphFlowNode {
	s.NextNodes = v
	return s
}

func (s *GraphFlowNode) SetNodeType(v string) *GraphFlowNode {
	s.NodeType = &v
	return s
}

func (s *GraphFlowNode) SetProperties(v *GraphFlowNodeProperties) *GraphFlowNode {
	s.Properties = v
	return s
}

func (s *GraphFlowNode) SetRid(v int64) *GraphFlowNode {
	s.Rid = &v
	return s
}

func (s *GraphFlowNode) SetUseConditions(v bool) *GraphFlowNode {
	s.UseConditions = &v
	return s
}

type GraphFlowNodeNextNodes struct {
	CheckType  *int32    `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	Index      *int32    `json:"Index,omitempty" xml:"Index,omitempty"`
	Lambda     *string   `json:"Lambda,omitempty" xml:"Lambda,omitempty"`
	Name       *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	NextNodeId *int64    `json:"NextNodeId,omitempty" xml:"NextNodeId,omitempty"`
	Triggers   []*string `json:"Triggers,omitempty" xml:"Triggers,omitempty" type:"Repeated"`
}

func (s GraphFlowNodeNextNodes) String() string {
	return tea.Prettify(s)
}

func (s GraphFlowNodeNextNodes) GoString() string {
	return s.String()
}

func (s *GraphFlowNodeNextNodes) SetCheckType(v int32) *GraphFlowNodeNextNodes {
	s.CheckType = &v
	return s
}

func (s *GraphFlowNodeNextNodes) SetIndex(v int32) *GraphFlowNodeNextNodes {
	s.Index = &v
	return s
}

func (s *GraphFlowNodeNextNodes) SetLambda(v string) *GraphFlowNodeNextNodes {
	s.Lambda = &v
	return s
}

func (s *GraphFlowNodeNextNodes) SetName(v string) *GraphFlowNodeNextNodes {
	s.Name = &v
	return s
}

func (s *GraphFlowNodeNextNodes) SetNextNodeId(v int64) *GraphFlowNodeNextNodes {
	s.NextNodeId = &v
	return s
}

func (s *GraphFlowNodeNextNodes) SetTriggers(v []*string) *GraphFlowNodeNextNodes {
	s.Triggers = v
	return s
}

type GraphFlowNodeProperties struct {
	AutoReview       *int32    `json:"AutoReview,omitempty" xml:"AutoReview,omitempty"`
	BranchJudge      *bool     `json:"BranchJudge,omitempty" xml:"BranchJudge,omitempty"`
	CheckMoreSize    *int32    `json:"CheckMoreSize,omitempty" xml:"CheckMoreSize,omitempty"`
	CheckType        *int32    `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	Lambda           *string   `json:"Lambda,omitempty" xml:"Lambda,omitempty"`
	Role             *string   `json:"Role,omitempty" xml:"Role,omitempty"`
	RuleScoreType    *int32    `json:"RuleScoreType,omitempty" xml:"RuleScoreType,omitempty"`
	SayType          *string   `json:"SayType,omitempty" xml:"SayType,omitempty"`
	ScoreNum         *int32    `json:"ScoreNum,omitempty" xml:"ScoreNum,omitempty"`
	ScoreNumType     *int32    `json:"ScoreNumType,omitempty" xml:"ScoreNumType,omitempty"`
	ScoreRuleHitType *int32    `json:"ScoreRuleHitType,omitempty" xml:"ScoreRuleHitType,omitempty"`
	ScoreType        *int32    `json:"ScoreType,omitempty" xml:"ScoreType,omitempty"`
	Triggers         []*string `json:"Triggers,omitempty" xml:"Triggers,omitempty" type:"Repeated"`
	Type             *string   `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GraphFlowNodeProperties) String() string {
	return tea.Prettify(s)
}

func (s GraphFlowNodeProperties) GoString() string {
	return s.String()
}

func (s *GraphFlowNodeProperties) SetAutoReview(v int32) *GraphFlowNodeProperties {
	s.AutoReview = &v
	return s
}

func (s *GraphFlowNodeProperties) SetBranchJudge(v bool) *GraphFlowNodeProperties {
	s.BranchJudge = &v
	return s
}

func (s *GraphFlowNodeProperties) SetCheckMoreSize(v int32) *GraphFlowNodeProperties {
	s.CheckMoreSize = &v
	return s
}

func (s *GraphFlowNodeProperties) SetCheckType(v int32) *GraphFlowNodeProperties {
	s.CheckType = &v
	return s
}

func (s *GraphFlowNodeProperties) SetLambda(v string) *GraphFlowNodeProperties {
	s.Lambda = &v
	return s
}

func (s *GraphFlowNodeProperties) SetRole(v string) *GraphFlowNodeProperties {
	s.Role = &v
	return s
}

func (s *GraphFlowNodeProperties) SetRuleScoreType(v int32) *GraphFlowNodeProperties {
	s.RuleScoreType = &v
	return s
}

func (s *GraphFlowNodeProperties) SetSayType(v string) *GraphFlowNodeProperties {
	s.SayType = &v
	return s
}

func (s *GraphFlowNodeProperties) SetScoreNum(v int32) *GraphFlowNodeProperties {
	s.ScoreNum = &v
	return s
}

func (s *GraphFlowNodeProperties) SetScoreNumType(v int32) *GraphFlowNodeProperties {
	s.ScoreNumType = &v
	return s
}

func (s *GraphFlowNodeProperties) SetScoreRuleHitType(v int32) *GraphFlowNodeProperties {
	s.ScoreRuleHitType = &v
	return s
}

func (s *GraphFlowNodeProperties) SetScoreType(v int32) *GraphFlowNodeProperties {
	s.ScoreType = &v
	return s
}

func (s *GraphFlowNodeProperties) SetTriggers(v []*string) *GraphFlowNodeProperties {
	s.Triggers = v
	return s
}

func (s *GraphFlowNodeProperties) SetType(v string) *GraphFlowNodeProperties {
	s.Type = &v
	return s
}

type JudgeNodeMetaDesc struct {
	ActualValue *string `json:"ActualValue,omitempty" xml:"ActualValue,omitempty"`
	DataType    *int32  `json:"DataType,omitempty" xml:"DataType,omitempty"`
	Field       *string `json:"Field,omitempty" xml:"Field,omitempty"`
	FieldType   *int32  `json:"FieldType,omitempty" xml:"FieldType,omitempty"`
	Symbol      *int32  `json:"Symbol,omitempty" xml:"Symbol,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s JudgeNodeMetaDesc) String() string {
	return tea.Prettify(s)
}

func (s JudgeNodeMetaDesc) GoString() string {
	return s.String()
}

func (s *JudgeNodeMetaDesc) SetActualValue(v string) *JudgeNodeMetaDesc {
	s.ActualValue = &v
	return s
}

func (s *JudgeNodeMetaDesc) SetDataType(v int32) *JudgeNodeMetaDesc {
	s.DataType = &v
	return s
}

func (s *JudgeNodeMetaDesc) SetField(v string) *JudgeNodeMetaDesc {
	s.Field = &v
	return s
}

func (s *JudgeNodeMetaDesc) SetFieldType(v int32) *JudgeNodeMetaDesc {
	s.FieldType = &v
	return s
}

func (s *JudgeNodeMetaDesc) SetSymbol(v int32) *JudgeNodeMetaDesc {
	s.Symbol = &v
	return s
}

func (s *JudgeNodeMetaDesc) SetValue(v string) *JudgeNodeMetaDesc {
	s.Value = &v
	return s
}

type NextNodeSituations struct {
	ConditionGroup []*NextNodeSituationsConditionGroup `json:"ConditionGroup,omitempty" xml:"ConditionGroup,omitempty" type:"Repeated"`
	Type           *string                             `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s NextNodeSituations) String() string {
	return tea.Prettify(s)
}

func (s NextNodeSituations) GoString() string {
	return s.String()
}

func (s *NextNodeSituations) SetConditionGroup(v []*NextNodeSituationsConditionGroup) *NextNodeSituations {
	s.ConditionGroup = v
	return s
}

func (s *NextNodeSituations) SetType(v string) *NextNodeSituations {
	s.Type = &v
	return s
}

type NextNodeSituationsConditionGroup struct {
	Conditions []*JudgeNodeMetaDesc `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	Type       *string              `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s NextNodeSituationsConditionGroup) String() string {
	return tea.Prettify(s)
}

func (s NextNodeSituationsConditionGroup) GoString() string {
	return s.String()
}

func (s *NextNodeSituationsConditionGroup) SetConditions(v []*JudgeNodeMetaDesc) *NextNodeSituationsConditionGroup {
	s.Conditions = v
	return s
}

func (s *NextNodeSituationsConditionGroup) SetType(v string) *NextNodeSituationsConditionGroup {
	s.Type = &v
	return s
}

type OperatorBasicInfo struct {
	Id               *int64                  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name             *string                 `json:"Name,omitempty" xml:"Name,omitempty"`
	Oid              *string                 `json:"Oid,omitempty" xml:"Oid,omitempty"`
	Param            *OperatorBasicInfoParam `json:"Param,omitempty" xml:"Param,omitempty" type:"Struct"`
	QualityCheckType *int32                  `json:"QualityCheckType,omitempty" xml:"QualityCheckType,omitempty"`
	Type             *string                 `json:"Type,omitempty" xml:"Type,omitempty"`
	UserGroup        *string                 `json:"UserGroup,omitempty" xml:"UserGroup,omitempty"`
}

func (s OperatorBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s OperatorBasicInfo) GoString() string {
	return s.String()
}

func (s *OperatorBasicInfo) SetId(v int64) *OperatorBasicInfo {
	s.Id = &v
	return s
}

func (s *OperatorBasicInfo) SetName(v string) *OperatorBasicInfo {
	s.Name = &v
	return s
}

func (s *OperatorBasicInfo) SetOid(v string) *OperatorBasicInfo {
	s.Oid = &v
	return s
}

func (s *OperatorBasicInfo) SetParam(v *OperatorBasicInfoParam) *OperatorBasicInfo {
	s.Param = v
	return s
}

func (s *OperatorBasicInfo) SetQualityCheckType(v int32) *OperatorBasicInfo {
	s.QualityCheckType = &v
	return s
}

func (s *OperatorBasicInfo) SetType(v string) *OperatorBasicInfo {
	s.Type = &v
	return s
}

func (s *OperatorBasicInfo) SetUserGroup(v string) *OperatorBasicInfo {
	s.UserGroup = &v
	return s
}

type OperatorBasicInfoParam struct {
	AnswerThreshold       *string              `json:"AnswerThreshold,omitempty" xml:"AnswerThreshold,omitempty"`
	AntModelInfo          map[string]*string   `json:"AntModelInfo,omitempty" xml:"AntModelInfo,omitempty"`
	Average               *bool                `json:"Average,omitempty" xml:"Average,omitempty"`
	BeginType             *string              `json:"BeginType,omitempty" xml:"BeginType,omitempty"`
	BotId                 *string              `json:"BotId,omitempty" xml:"BotId,omitempty"`
	CaseSensitive         *bool                `json:"Case_sensitive,omitempty" xml:"Case_sensitive,omitempty"`
	CategoryPathCode      *string              `json:"CategoryPathCode,omitempty" xml:"CategoryPathCode,omitempty"`
	CheckFirstSentence    *bool                `json:"CheckFirstSentence,omitempty" xml:"CheckFirstSentence,omitempty"`
	CheckType             *int32               `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	CompareOperator       *string              `json:"CompareOperator,omitempty" xml:"CompareOperator,omitempty"`
	ContextChatMatch      *bool                `json:"ContextChatMatch,omitempty" xml:"ContextChatMatch,omitempty"`
	CustomerParam         *JudgeNodeMetaDesc   `json:"CustomerParam,omitempty" xml:"CustomerParam,omitempty"`
	DelayTime             *int32               `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	DifferentRole         *bool                `json:"Different_role,omitempty" xml:"Different_role,omitempty"`
	EndType               *string              `json:"EndType,omitempty" xml:"EndType,omitempty"`
	Excludes              []*string            `json:"Excludes,omitempty" xml:"Excludes,omitempty" type:"Repeated"`
	From                  *int32               `json:"From,omitempty" xml:"From,omitempty"`
	FromEnd               *bool                `json:"From_end,omitempty" xml:"From_end,omitempty"`
	HitTime               *int32               `json:"Hit_time,omitempty" xml:"Hit_time,omitempty"`
	InSentence            *bool                `json:"In_sentence,omitempty" xml:"In_sentence,omitempty"`
	Interval              *int32               `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IntervalEnd           *int32               `json:"IntervalEnd,omitempty" xml:"IntervalEnd,omitempty"`
	KeywordExtension      *int32               `json:"KeywordExtension,omitempty" xml:"KeywordExtension,omitempty"`
	KeywordMatchSize      *int32               `json:"KeywordMatchSize,omitempty" xml:"KeywordMatchSize,omitempty"`
	Keywords              []*string            `json:"Keywords,omitempty" xml:"Keywords,omitempty" type:"Repeated"`
	KnowledgeInfo         *string              `json:"KnowledgeInfo,omitempty" xml:"KnowledgeInfo,omitempty"`
	KnowledgeSentenceNum  *int32               `json:"KnowledgeSentenceNum,omitempty" xml:"KnowledgeSentenceNum,omitempty"`
	KnowledgeTargetId     *string              `json:"KnowledgeTargetId,omitempty" xml:"KnowledgeTargetId,omitempty"`
	KnowledgeTargetName   *string              `json:"KnowledgeTargetName,omitempty" xml:"KnowledgeTargetName,omitempty"`
	KnowledgeTargetType   *int32               `json:"KnowledgeTargetType,omitempty" xml:"KnowledgeTargetType,omitempty"`
	LgfSentences          []*string            `json:"LgfSentences,omitempty" xml:"LgfSentences,omitempty" type:"Repeated"`
	MaxEmotionChangeValue *int32               `json:"MaxEmotionChangeValue,omitempty" xml:"MaxEmotionChangeValue,omitempty"`
	MinWordSize           *int32               `json:"MinWordSize,omitempty" xml:"MinWordSize,omitempty"`
	NearDialogue          *bool                `json:"Near_dialogue,omitempty" xml:"Near_dialogue,omitempty"`
	NotRegex              *string              `json:"NotRegex,omitempty" xml:"NotRegex,omitempty"`
	Phrase                *string              `json:"Phrase,omitempty" xml:"Phrase,omitempty"`
	Pkey                  *string              `json:"Pkey,omitempty" xml:"Pkey,omitempty"`
	PoutputType           *int32               `json:"Poutput_type,omitempty" xml:"Poutput_type,omitempty"`
	Pvalues               []*string            `json:"Pvalues,omitempty" xml:"Pvalues,omitempty" type:"Repeated"`
	QuestionThreshold     *string              `json:"QuestionThreshold,omitempty" xml:"QuestionThreshold,omitempty"`
	References            []*string            `json:"References,omitempty" xml:"References,omitempty" type:"Repeated"`
	Regex                 *string              `json:"Regex,omitempty" xml:"Regex,omitempty"`
	RoleId                *int32               `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	Score                 *int32               `json:"Score,omitempty" xml:"Score,omitempty"`
	SimilarityThreshold   *float64             `json:"Similarity_threshold,omitempty" xml:"Similarity_threshold,omitempty"`
	SimilarlySentences    []*string            `json:"SimilarlySentences,omitempty" xml:"SimilarlySentences,omitempty" type:"Repeated"`
	Synonyms              map[string][]*string `json:"Synonyms,omitempty" xml:"Synonyms,omitempty"`
	Target                *int32               `json:"Target,omitempty" xml:"Target,omitempty"`
	TargetRole            *string              `json:"Target_role,omitempty" xml:"Target_role,omitempty"`
	Threshold             *float32             `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	UseEasAlgorithm       *bool                `json:"UseEasAlgorithm,omitempty" xml:"UseEasAlgorithm,omitempty"`
	Velocity              *float64             `json:"Velocity,omitempty" xml:"Velocity,omitempty"`
	VelocityInMint        *int32               `json:"VelocityInMint,omitempty" xml:"VelocityInMint,omitempty"`
}

func (s OperatorBasicInfoParam) String() string {
	return tea.Prettify(s)
}

func (s OperatorBasicInfoParam) GoString() string {
	return s.String()
}

func (s *OperatorBasicInfoParam) SetAnswerThreshold(v string) *OperatorBasicInfoParam {
	s.AnswerThreshold = &v
	return s
}

func (s *OperatorBasicInfoParam) SetAntModelInfo(v map[string]*string) *OperatorBasicInfoParam {
	s.AntModelInfo = v
	return s
}

func (s *OperatorBasicInfoParam) SetAverage(v bool) *OperatorBasicInfoParam {
	s.Average = &v
	return s
}

func (s *OperatorBasicInfoParam) SetBeginType(v string) *OperatorBasicInfoParam {
	s.BeginType = &v
	return s
}

func (s *OperatorBasicInfoParam) SetBotId(v string) *OperatorBasicInfoParam {
	s.BotId = &v
	return s
}

func (s *OperatorBasicInfoParam) SetCaseSensitive(v bool) *OperatorBasicInfoParam {
	s.CaseSensitive = &v
	return s
}

func (s *OperatorBasicInfoParam) SetCategoryPathCode(v string) *OperatorBasicInfoParam {
	s.CategoryPathCode = &v
	return s
}

func (s *OperatorBasicInfoParam) SetCheckFirstSentence(v bool) *OperatorBasicInfoParam {
	s.CheckFirstSentence = &v
	return s
}

func (s *OperatorBasicInfoParam) SetCheckType(v int32) *OperatorBasicInfoParam {
	s.CheckType = &v
	return s
}

func (s *OperatorBasicInfoParam) SetCompareOperator(v string) *OperatorBasicInfoParam {
	s.CompareOperator = &v
	return s
}

func (s *OperatorBasicInfoParam) SetContextChatMatch(v bool) *OperatorBasicInfoParam {
	s.ContextChatMatch = &v
	return s
}

func (s *OperatorBasicInfoParam) SetCustomerParam(v *JudgeNodeMetaDesc) *OperatorBasicInfoParam {
	s.CustomerParam = v
	return s
}

func (s *OperatorBasicInfoParam) SetDelayTime(v int32) *OperatorBasicInfoParam {
	s.DelayTime = &v
	return s
}

func (s *OperatorBasicInfoParam) SetDifferentRole(v bool) *OperatorBasicInfoParam {
	s.DifferentRole = &v
	return s
}

func (s *OperatorBasicInfoParam) SetEndType(v string) *OperatorBasicInfoParam {
	s.EndType = &v
	return s
}

func (s *OperatorBasicInfoParam) SetExcludes(v []*string) *OperatorBasicInfoParam {
	s.Excludes = v
	return s
}

func (s *OperatorBasicInfoParam) SetFrom(v int32) *OperatorBasicInfoParam {
	s.From = &v
	return s
}

func (s *OperatorBasicInfoParam) SetFromEnd(v bool) *OperatorBasicInfoParam {
	s.FromEnd = &v
	return s
}

func (s *OperatorBasicInfoParam) SetHitTime(v int32) *OperatorBasicInfoParam {
	s.HitTime = &v
	return s
}

func (s *OperatorBasicInfoParam) SetInSentence(v bool) *OperatorBasicInfoParam {
	s.InSentence = &v
	return s
}

func (s *OperatorBasicInfoParam) SetInterval(v int32) *OperatorBasicInfoParam {
	s.Interval = &v
	return s
}

func (s *OperatorBasicInfoParam) SetIntervalEnd(v int32) *OperatorBasicInfoParam {
	s.IntervalEnd = &v
	return s
}

func (s *OperatorBasicInfoParam) SetKeywordExtension(v int32) *OperatorBasicInfoParam {
	s.KeywordExtension = &v
	return s
}

func (s *OperatorBasicInfoParam) SetKeywordMatchSize(v int32) *OperatorBasicInfoParam {
	s.KeywordMatchSize = &v
	return s
}

func (s *OperatorBasicInfoParam) SetKeywords(v []*string) *OperatorBasicInfoParam {
	s.Keywords = v
	return s
}

func (s *OperatorBasicInfoParam) SetKnowledgeInfo(v string) *OperatorBasicInfoParam {
	s.KnowledgeInfo = &v
	return s
}

func (s *OperatorBasicInfoParam) SetKnowledgeSentenceNum(v int32) *OperatorBasicInfoParam {
	s.KnowledgeSentenceNum = &v
	return s
}

func (s *OperatorBasicInfoParam) SetKnowledgeTargetId(v string) *OperatorBasicInfoParam {
	s.KnowledgeTargetId = &v
	return s
}

func (s *OperatorBasicInfoParam) SetKnowledgeTargetName(v string) *OperatorBasicInfoParam {
	s.KnowledgeTargetName = &v
	return s
}

func (s *OperatorBasicInfoParam) SetKnowledgeTargetType(v int32) *OperatorBasicInfoParam {
	s.KnowledgeTargetType = &v
	return s
}

func (s *OperatorBasicInfoParam) SetLgfSentences(v []*string) *OperatorBasicInfoParam {
	s.LgfSentences = v
	return s
}

func (s *OperatorBasicInfoParam) SetMaxEmotionChangeValue(v int32) *OperatorBasicInfoParam {
	s.MaxEmotionChangeValue = &v
	return s
}

func (s *OperatorBasicInfoParam) SetMinWordSize(v int32) *OperatorBasicInfoParam {
	s.MinWordSize = &v
	return s
}

func (s *OperatorBasicInfoParam) SetNearDialogue(v bool) *OperatorBasicInfoParam {
	s.NearDialogue = &v
	return s
}

func (s *OperatorBasicInfoParam) SetNotRegex(v string) *OperatorBasicInfoParam {
	s.NotRegex = &v
	return s
}

func (s *OperatorBasicInfoParam) SetPhrase(v string) *OperatorBasicInfoParam {
	s.Phrase = &v
	return s
}

func (s *OperatorBasicInfoParam) SetPkey(v string) *OperatorBasicInfoParam {
	s.Pkey = &v
	return s
}

func (s *OperatorBasicInfoParam) SetPoutputType(v int32) *OperatorBasicInfoParam {
	s.PoutputType = &v
	return s
}

func (s *OperatorBasicInfoParam) SetPvalues(v []*string) *OperatorBasicInfoParam {
	s.Pvalues = v
	return s
}

func (s *OperatorBasicInfoParam) SetQuestionThreshold(v string) *OperatorBasicInfoParam {
	s.QuestionThreshold = &v
	return s
}

func (s *OperatorBasicInfoParam) SetReferences(v []*string) *OperatorBasicInfoParam {
	s.References = v
	return s
}

func (s *OperatorBasicInfoParam) SetRegex(v string) *OperatorBasicInfoParam {
	s.Regex = &v
	return s
}

func (s *OperatorBasicInfoParam) SetRoleId(v int32) *OperatorBasicInfoParam {
	s.RoleId = &v
	return s
}

func (s *OperatorBasicInfoParam) SetScore(v int32) *OperatorBasicInfoParam {
	s.Score = &v
	return s
}

func (s *OperatorBasicInfoParam) SetSimilarityThreshold(v float64) *OperatorBasicInfoParam {
	s.SimilarityThreshold = &v
	return s
}

func (s *OperatorBasicInfoParam) SetSimilarlySentences(v []*string) *OperatorBasicInfoParam {
	s.SimilarlySentences = v
	return s
}

func (s *OperatorBasicInfoParam) SetSynonyms(v map[string][]*string) *OperatorBasicInfoParam {
	s.Synonyms = v
	return s
}

func (s *OperatorBasicInfoParam) SetTarget(v int32) *OperatorBasicInfoParam {
	s.Target = &v
	return s
}

func (s *OperatorBasicInfoParam) SetTargetRole(v string) *OperatorBasicInfoParam {
	s.TargetRole = &v
	return s
}

func (s *OperatorBasicInfoParam) SetThreshold(v float32) *OperatorBasicInfoParam {
	s.Threshold = &v
	return s
}

func (s *OperatorBasicInfoParam) SetUseEasAlgorithm(v bool) *OperatorBasicInfoParam {
	s.UseEasAlgorithm = &v
	return s
}

func (s *OperatorBasicInfoParam) SetVelocity(v float64) *OperatorBasicInfoParam {
	s.Velocity = &v
	return s
}

func (s *OperatorBasicInfoParam) SetVelocityInMint(v int32) *OperatorBasicInfoParam {
	s.VelocityInMint = &v
	return s
}

type RuleCountInfo struct {
	AutoReview                    *int32                       `json:"AutoReview,omitempty" xml:"AutoReview,omitempty"`
	BusinessCategoryBasicInfoList []*BusinessCategoryBasicInfo `json:"BusinessCategoryBasicInfoList,omitempty" xml:"BusinessCategoryBasicInfoList,omitempty" type:"Repeated"`
	BusinessCategoryNameList      []*string                    `json:"BusinessCategoryNameList,omitempty" xml:"BusinessCategoryNameList,omitempty" type:"Repeated"`
	BusinessRange                 []*int32                     `json:"BusinessRange,omitempty" xml:"BusinessRange,omitempty" type:"Repeated"`
	CheckNumber                   *int64                       `json:"CheckNumber,omitempty" xml:"CheckNumber,omitempty"`
	Comments                      *string                      `json:"Comments,omitempty" xml:"Comments,omitempty"`
	CreateEmpName                 *string                      `json:"CreateEmpName,omitempty" xml:"CreateEmpName,omitempty"`
	CreateEmpid                   *string                      `json:"CreateEmpid,omitempty" xml:"CreateEmpid,omitempty"`
	CreateTime                    *string                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deny                          *int32                       `json:"Deny,omitempty" xml:"Deny,omitempty"`
	Effective                     *int32                       `json:"Effective,omitempty" xml:"Effective,omitempty"`
	EffectiveEndTime              *string                      `json:"EffectiveEndTime,omitempty" xml:"EffectiveEndTime,omitempty"`
	EffectiveStartTime            *string                      `json:"EffectiveStartTime,omitempty" xml:"EffectiveStartTime,omitempty"`
	EndTime                       *string                      `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FullCycle                     *int32                       `json:"FullCycle,omitempty" xml:"FullCycle,omitempty"`
	GraphFlow                     interface{}                  `json:"GraphFlow,omitempty" xml:"GraphFlow,omitempty"`
	HitNumber                     *int64                       `json:"HitNumber,omitempty" xml:"HitNumber,omitempty"`
	HitRate                       *float32                     `json:"HitRate,omitempty" xml:"HitRate,omitempty"`
	HitRealViolationRate          *float32                     `json:"HitRealViolationRate,omitempty" xml:"HitRealViolationRate,omitempty"`
	IsDelete                      *int32                       `json:"IsDelete,omitempty" xml:"IsDelete,omitempty"`
	IsSelect                      *bool                        `json:"IsSelect,omitempty" xml:"IsSelect,omitempty"`
	JobName                       *string                      `json:"JobName,omitempty" xml:"JobName,omitempty"`
	LastUpdateEmpName             *string                      `json:"LastUpdateEmpName,omitempty" xml:"LastUpdateEmpName,omitempty"`
	LastUpdateEmpid               *string                      `json:"LastUpdateEmpid,omitempty" xml:"LastUpdateEmpid,omitempty"`
	LastUpdateTime                *string                      `json:"LastUpdateTime,omitempty" xml:"LastUpdateTime,omitempty"`
	Name                          *string                      `json:"Name,omitempty" xml:"Name,omitempty"`
	OperationMode                 *int32                       `json:"OperationMode,omitempty" xml:"OperationMode,omitempty"`
	PreReviewNumber               *int64                       `json:"PreReviewNumber,omitempty" xml:"PreReviewNumber,omitempty"`
	ProblemNumber                 *int64                       `json:"ProblemNumber,omitempty" xml:"ProblemNumber,omitempty"`
	QualityCheckType              *int32                       `json:"QualityCheckType,omitempty" xml:"QualityCheckType,omitempty"`
	RealViolationNumber           *int32                       `json:"RealViolationNumber,omitempty" xml:"RealViolationNumber,omitempty"`
	ReviewAccuracyRate            *float32                     `json:"ReviewAccuracyRate,omitempty" xml:"ReviewAccuracyRate,omitempty"`
	ReviewNumber                  *int64                       `json:"ReviewNumber,omitempty" xml:"ReviewNumber,omitempty"`
	ReviewRate                    *float32                     `json:"ReviewRate,omitempty" xml:"ReviewRate,omitempty"`
	ReviewStatusName              *string                      `json:"ReviewStatusName,omitempty" xml:"ReviewStatusName,omitempty"`
	Rid                           *int64                       `json:"Rid,omitempty" xml:"Rid,omitempty"`
	RuleScoreSingleType           *int32                       `json:"RuleScoreSingleType,omitempty" xml:"RuleScoreSingleType,omitempty"`
	RuleScoreType                 *int32                       `json:"RuleScoreType,omitempty" xml:"RuleScoreType,omitempty"`
	RuleType                      *int32                       `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	ScoreSubId                    *int64                       `json:"ScoreSubId,omitempty" xml:"ScoreSubId,omitempty"`
	StartTime                     *string                      `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status                        *int32                       `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetType                    *int32                       `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	Type                          *int32                       `json:"Type,omitempty" xml:"Type,omitempty"`
	TypeName                      *string                      `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
	UnReviewNumber                *int64                       `json:"UnReviewNumber,omitempty" xml:"UnReviewNumber,omitempty"`
	UserGroup                     *string                      `json:"UserGroup,omitempty" xml:"UserGroup,omitempty"`
}

func (s RuleCountInfo) String() string {
	return tea.Prettify(s)
}

func (s RuleCountInfo) GoString() string {
	return s.String()
}

func (s *RuleCountInfo) SetAutoReview(v int32) *RuleCountInfo {
	s.AutoReview = &v
	return s
}

func (s *RuleCountInfo) SetBusinessCategoryBasicInfoList(v []*BusinessCategoryBasicInfo) *RuleCountInfo {
	s.BusinessCategoryBasicInfoList = v
	return s
}

func (s *RuleCountInfo) SetBusinessCategoryNameList(v []*string) *RuleCountInfo {
	s.BusinessCategoryNameList = v
	return s
}

func (s *RuleCountInfo) SetBusinessRange(v []*int32) *RuleCountInfo {
	s.BusinessRange = v
	return s
}

func (s *RuleCountInfo) SetCheckNumber(v int64) *RuleCountInfo {
	s.CheckNumber = &v
	return s
}

func (s *RuleCountInfo) SetComments(v string) *RuleCountInfo {
	s.Comments = &v
	return s
}

func (s *RuleCountInfo) SetCreateEmpName(v string) *RuleCountInfo {
	s.CreateEmpName = &v
	return s
}

func (s *RuleCountInfo) SetCreateEmpid(v string) *RuleCountInfo {
	s.CreateEmpid = &v
	return s
}

func (s *RuleCountInfo) SetCreateTime(v string) *RuleCountInfo {
	s.CreateTime = &v
	return s
}

func (s *RuleCountInfo) SetDeny(v int32) *RuleCountInfo {
	s.Deny = &v
	return s
}

func (s *RuleCountInfo) SetEffective(v int32) *RuleCountInfo {
	s.Effective = &v
	return s
}

func (s *RuleCountInfo) SetEffectiveEndTime(v string) *RuleCountInfo {
	s.EffectiveEndTime = &v
	return s
}

func (s *RuleCountInfo) SetEffectiveStartTime(v string) *RuleCountInfo {
	s.EffectiveStartTime = &v
	return s
}

func (s *RuleCountInfo) SetEndTime(v string) *RuleCountInfo {
	s.EndTime = &v
	return s
}

func (s *RuleCountInfo) SetFullCycle(v int32) *RuleCountInfo {
	s.FullCycle = &v
	return s
}

func (s *RuleCountInfo) SetGraphFlow(v interface{}) *RuleCountInfo {
	s.GraphFlow = v
	return s
}

func (s *RuleCountInfo) SetHitNumber(v int64) *RuleCountInfo {
	s.HitNumber = &v
	return s
}

func (s *RuleCountInfo) SetHitRate(v float32) *RuleCountInfo {
	s.HitRate = &v
	return s
}

func (s *RuleCountInfo) SetHitRealViolationRate(v float32) *RuleCountInfo {
	s.HitRealViolationRate = &v
	return s
}

func (s *RuleCountInfo) SetIsDelete(v int32) *RuleCountInfo {
	s.IsDelete = &v
	return s
}

func (s *RuleCountInfo) SetIsSelect(v bool) *RuleCountInfo {
	s.IsSelect = &v
	return s
}

func (s *RuleCountInfo) SetJobName(v string) *RuleCountInfo {
	s.JobName = &v
	return s
}

func (s *RuleCountInfo) SetLastUpdateEmpName(v string) *RuleCountInfo {
	s.LastUpdateEmpName = &v
	return s
}

func (s *RuleCountInfo) SetLastUpdateEmpid(v string) *RuleCountInfo {
	s.LastUpdateEmpid = &v
	return s
}

func (s *RuleCountInfo) SetLastUpdateTime(v string) *RuleCountInfo {
	s.LastUpdateTime = &v
	return s
}

func (s *RuleCountInfo) SetName(v string) *RuleCountInfo {
	s.Name = &v
	return s
}

func (s *RuleCountInfo) SetOperationMode(v int32) *RuleCountInfo {
	s.OperationMode = &v
	return s
}

func (s *RuleCountInfo) SetPreReviewNumber(v int64) *RuleCountInfo {
	s.PreReviewNumber = &v
	return s
}

func (s *RuleCountInfo) SetProblemNumber(v int64) *RuleCountInfo {
	s.ProblemNumber = &v
	return s
}

func (s *RuleCountInfo) SetQualityCheckType(v int32) *RuleCountInfo {
	s.QualityCheckType = &v
	return s
}

func (s *RuleCountInfo) SetRealViolationNumber(v int32) *RuleCountInfo {
	s.RealViolationNumber = &v
	return s
}

func (s *RuleCountInfo) SetReviewAccuracyRate(v float32) *RuleCountInfo {
	s.ReviewAccuracyRate = &v
	return s
}

func (s *RuleCountInfo) SetReviewNumber(v int64) *RuleCountInfo {
	s.ReviewNumber = &v
	return s
}

func (s *RuleCountInfo) SetReviewRate(v float32) *RuleCountInfo {
	s.ReviewRate = &v
	return s
}

func (s *RuleCountInfo) SetReviewStatusName(v string) *RuleCountInfo {
	s.ReviewStatusName = &v
	return s
}

func (s *RuleCountInfo) SetRid(v int64) *RuleCountInfo {
	s.Rid = &v
	return s
}

func (s *RuleCountInfo) SetRuleScoreSingleType(v int32) *RuleCountInfo {
	s.RuleScoreSingleType = &v
	return s
}

func (s *RuleCountInfo) SetRuleScoreType(v int32) *RuleCountInfo {
	s.RuleScoreType = &v
	return s
}

func (s *RuleCountInfo) SetRuleType(v int32) *RuleCountInfo {
	s.RuleType = &v
	return s
}

func (s *RuleCountInfo) SetScoreSubId(v int64) *RuleCountInfo {
	s.ScoreSubId = &v
	return s
}

func (s *RuleCountInfo) SetStartTime(v string) *RuleCountInfo {
	s.StartTime = &v
	return s
}

func (s *RuleCountInfo) SetStatus(v int32) *RuleCountInfo {
	s.Status = &v
	return s
}

func (s *RuleCountInfo) SetTargetType(v int32) *RuleCountInfo {
	s.TargetType = &v
	return s
}

func (s *RuleCountInfo) SetType(v int32) *RuleCountInfo {
	s.Type = &v
	return s
}

func (s *RuleCountInfo) SetTypeName(v string) *RuleCountInfo {
	s.TypeName = &v
	return s
}

func (s *RuleCountInfo) SetUnReviewNumber(v int64) *RuleCountInfo {
	s.UnReviewNumber = &v
	return s
}

func (s *RuleCountInfo) SetUserGroup(v string) *RuleCountInfo {
	s.UserGroup = &v
	return s
}

type RuleInfo struct {
	AutoReview               *int32              `json:"AutoReview,omitempty" xml:"AutoReview,omitempty"`
	BusinessCategoryNameList []*string           `json:"BusinessCategoryNameList,omitempty" xml:"BusinessCategoryNameList,omitempty" type:"Repeated"`
	CheckType                *int64              `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	Comments                 *string             `json:"Comments,omitempty" xml:"Comments,omitempty"`
	ConfigType               *int32              `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	CreateEmpName            *string             `json:"CreateEmpName,omitempty" xml:"CreateEmpName,omitempty"`
	CreateEmpid              *string             `json:"CreateEmpid,omitempty" xml:"CreateEmpid,omitempty"`
	CreateTime               *string             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deny                     *int32              `json:"Deny,omitempty" xml:"Deny,omitempty"`
	Dialogues                []*RuleTestDialogue `json:"Dialogues,omitempty" xml:"Dialogues,omitempty" type:"Repeated"`
	Effective                *int32              `json:"Effective,omitempty" xml:"Effective,omitempty"`
	EffectiveEndTime         *string             `json:"EffectiveEndTime,omitempty" xml:"EffectiveEndTime,omitempty"`
	EffectiveStartTime       *string             `json:"EffectiveStartTime,omitempty" xml:"EffectiveStartTime,omitempty"`
	EndTime                  *string             `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExternalProperty         *int32              `json:"ExternalProperty,omitempty" xml:"ExternalProperty,omitempty"`
	FullCycle                *int32              `json:"FullCycle,omitempty" xml:"FullCycle,omitempty"`
	GraphFlow                interface{}         `json:"GraphFlow,omitempty" xml:"GraphFlow,omitempty"`
	IsDelete                 *int32              `json:"IsDelete,omitempty" xml:"IsDelete,omitempty"`
	IsOnline                 *int32              `json:"IsOnline,omitempty" xml:"IsOnline,omitempty"`
	Lambda                   *string             `json:"Lambda,omitempty" xml:"Lambda,omitempty"`
	LastUpdateEmpName        *string             `json:"LastUpdateEmpName,omitempty" xml:"LastUpdateEmpName,omitempty"`
	LastUpdateEmpid          *string             `json:"LastUpdateEmpid,omitempty" xml:"LastUpdateEmpid,omitempty"`
	LastUpdateTime           *string             `json:"LastUpdateTime,omitempty" xml:"LastUpdateTime,omitempty"`
	Level                    *int32              `json:"Level,omitempty" xml:"Level,omitempty"`
	Meet                     *int32              `json:"Meet,omitempty" xml:"Meet,omitempty"`
	ModifyType               *int32              `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	Name                     *string             `json:"Name,omitempty" xml:"Name,omitempty"`
	OperationMode            *int32              `json:"OperationMode,omitempty" xml:"OperationMode,omitempty"`
	QualityCheckType         *int32              `json:"QualityCheckType,omitempty" xml:"QualityCheckType,omitempty"`
	Rid                      *string             `json:"Rid,omitempty" xml:"Rid,omitempty"`
	RuleCategoryName         *string             `json:"RuleCategoryName,omitempty" xml:"RuleCategoryName,omitempty"`
	RuleScoreType            *int32              `json:"RuleScoreType,omitempty" xml:"RuleScoreType,omitempty"`
	RuleType                 *int32              `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	SchemeCheckType          *SchemeCheckType    `json:"SchemeCheckType,omitempty" xml:"SchemeCheckType,omitempty"`
	SchemeId                 *int64              `json:"SchemeId,omitempty" xml:"SchemeId,omitempty"`
	SchemeName               *string             `json:"SchemeName,omitempty" xml:"SchemeName,omitempty"`
	SchemeRuleMappingId      *int64              `json:"SchemeRuleMappingId,omitempty" xml:"SchemeRuleMappingId,omitempty"`
	ScoreDeleted             *bool               `json:"ScoreDeleted,omitempty" xml:"ScoreDeleted,omitempty"`
	ScoreId                  *int64              `json:"ScoreId,omitempty" xml:"ScoreId,omitempty"`
	ScoreName                *string             `json:"ScoreName,omitempty" xml:"ScoreName,omitempty"`
	ScoreNum                 *int32              `json:"ScoreNum,omitempty" xml:"ScoreNum,omitempty"`
	ScoreNumType             *int32              `json:"ScoreNumType,omitempty" xml:"ScoreNumType,omitempty"`
	ScoreRuleHitType         *int32              `json:"ScoreRuleHitType,omitempty" xml:"ScoreRuleHitType,omitempty"`
	ScoreSubId               *int64              `json:"ScoreSubId,omitempty" xml:"ScoreSubId,omitempty"`
	ScoreSubName             *string             `json:"ScoreSubName,omitempty" xml:"ScoreSubName,omitempty"`
	ScoreType                *int32              `json:"ScoreType,omitempty" xml:"ScoreType,omitempty"`
	SortIndex                *int32              `json:"SortIndex,omitempty" xml:"SortIndex,omitempty"`
	StartTime                *string             `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status                   *int32              `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetType               *int32              `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	TaskFlowId               *int64              `json:"TaskFlowId,omitempty" xml:"TaskFlowId,omitempty"`
	TaskFlowType             *int32              `json:"TaskFlowType,omitempty" xml:"TaskFlowType,omitempty"`
	Triggers                 []*string           `json:"Triggers,omitempty" xml:"Triggers,omitempty" type:"Repeated"`
	Type                     *int32              `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight                   *string             `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s RuleInfo) String() string {
	return tea.Prettify(s)
}

func (s RuleInfo) GoString() string {
	return s.String()
}

func (s *RuleInfo) SetAutoReview(v int32) *RuleInfo {
	s.AutoReview = &v
	return s
}

func (s *RuleInfo) SetBusinessCategoryNameList(v []*string) *RuleInfo {
	s.BusinessCategoryNameList = v
	return s
}

func (s *RuleInfo) SetCheckType(v int64) *RuleInfo {
	s.CheckType = &v
	return s
}

func (s *RuleInfo) SetComments(v string) *RuleInfo {
	s.Comments = &v
	return s
}

func (s *RuleInfo) SetConfigType(v int32) *RuleInfo {
	s.ConfigType = &v
	return s
}

func (s *RuleInfo) SetCreateEmpName(v string) *RuleInfo {
	s.CreateEmpName = &v
	return s
}

func (s *RuleInfo) SetCreateEmpid(v string) *RuleInfo {
	s.CreateEmpid = &v
	return s
}

func (s *RuleInfo) SetCreateTime(v string) *RuleInfo {
	s.CreateTime = &v
	return s
}

func (s *RuleInfo) SetDeny(v int32) *RuleInfo {
	s.Deny = &v
	return s
}

func (s *RuleInfo) SetDialogues(v []*RuleTestDialogue) *RuleInfo {
	s.Dialogues = v
	return s
}

func (s *RuleInfo) SetEffective(v int32) *RuleInfo {
	s.Effective = &v
	return s
}

func (s *RuleInfo) SetEffectiveEndTime(v string) *RuleInfo {
	s.EffectiveEndTime = &v
	return s
}

func (s *RuleInfo) SetEffectiveStartTime(v string) *RuleInfo {
	s.EffectiveStartTime = &v
	return s
}

func (s *RuleInfo) SetEndTime(v string) *RuleInfo {
	s.EndTime = &v
	return s
}

func (s *RuleInfo) SetExternalProperty(v int32) *RuleInfo {
	s.ExternalProperty = &v
	return s
}

func (s *RuleInfo) SetFullCycle(v int32) *RuleInfo {
	s.FullCycle = &v
	return s
}

func (s *RuleInfo) SetGraphFlow(v interface{}) *RuleInfo {
	s.GraphFlow = v
	return s
}

func (s *RuleInfo) SetIsDelete(v int32) *RuleInfo {
	s.IsDelete = &v
	return s
}

func (s *RuleInfo) SetIsOnline(v int32) *RuleInfo {
	s.IsOnline = &v
	return s
}

func (s *RuleInfo) SetLambda(v string) *RuleInfo {
	s.Lambda = &v
	return s
}

func (s *RuleInfo) SetLastUpdateEmpName(v string) *RuleInfo {
	s.LastUpdateEmpName = &v
	return s
}

func (s *RuleInfo) SetLastUpdateEmpid(v string) *RuleInfo {
	s.LastUpdateEmpid = &v
	return s
}

func (s *RuleInfo) SetLastUpdateTime(v string) *RuleInfo {
	s.LastUpdateTime = &v
	return s
}

func (s *RuleInfo) SetLevel(v int32) *RuleInfo {
	s.Level = &v
	return s
}

func (s *RuleInfo) SetMeet(v int32) *RuleInfo {
	s.Meet = &v
	return s
}

func (s *RuleInfo) SetModifyType(v int32) *RuleInfo {
	s.ModifyType = &v
	return s
}

func (s *RuleInfo) SetName(v string) *RuleInfo {
	s.Name = &v
	return s
}

func (s *RuleInfo) SetOperationMode(v int32) *RuleInfo {
	s.OperationMode = &v
	return s
}

func (s *RuleInfo) SetQualityCheckType(v int32) *RuleInfo {
	s.QualityCheckType = &v
	return s
}

func (s *RuleInfo) SetRid(v string) *RuleInfo {
	s.Rid = &v
	return s
}

func (s *RuleInfo) SetRuleCategoryName(v string) *RuleInfo {
	s.RuleCategoryName = &v
	return s
}

func (s *RuleInfo) SetRuleScoreType(v int32) *RuleInfo {
	s.RuleScoreType = &v
	return s
}

func (s *RuleInfo) SetRuleType(v int32) *RuleInfo {
	s.RuleType = &v
	return s
}

func (s *RuleInfo) SetSchemeCheckType(v *SchemeCheckType) *RuleInfo {
	s.SchemeCheckType = v
	return s
}

func (s *RuleInfo) SetSchemeId(v int64) *RuleInfo {
	s.SchemeId = &v
	return s
}

func (s *RuleInfo) SetSchemeName(v string) *RuleInfo {
	s.SchemeName = &v
	return s
}

func (s *RuleInfo) SetSchemeRuleMappingId(v int64) *RuleInfo {
	s.SchemeRuleMappingId = &v
	return s
}

func (s *RuleInfo) SetScoreDeleted(v bool) *RuleInfo {
	s.ScoreDeleted = &v
	return s
}

func (s *RuleInfo) SetScoreId(v int64) *RuleInfo {
	s.ScoreId = &v
	return s
}

func (s *RuleInfo) SetScoreName(v string) *RuleInfo {
	s.ScoreName = &v
	return s
}

func (s *RuleInfo) SetScoreNum(v int32) *RuleInfo {
	s.ScoreNum = &v
	return s
}

func (s *RuleInfo) SetScoreNumType(v int32) *RuleInfo {
	s.ScoreNumType = &v
	return s
}

func (s *RuleInfo) SetScoreRuleHitType(v int32) *RuleInfo {
	s.ScoreRuleHitType = &v
	return s
}

func (s *RuleInfo) SetScoreSubId(v int64) *RuleInfo {
	s.ScoreSubId = &v
	return s
}

func (s *RuleInfo) SetScoreSubName(v string) *RuleInfo {
	s.ScoreSubName = &v
	return s
}

func (s *RuleInfo) SetScoreType(v int32) *RuleInfo {
	s.ScoreType = &v
	return s
}

func (s *RuleInfo) SetSortIndex(v int32) *RuleInfo {
	s.SortIndex = &v
	return s
}

func (s *RuleInfo) SetStartTime(v string) *RuleInfo {
	s.StartTime = &v
	return s
}

func (s *RuleInfo) SetStatus(v int32) *RuleInfo {
	s.Status = &v
	return s
}

func (s *RuleInfo) SetTargetType(v int32) *RuleInfo {
	s.TargetType = &v
	return s
}

func (s *RuleInfo) SetTaskFlowId(v int64) *RuleInfo {
	s.TaskFlowId = &v
	return s
}

func (s *RuleInfo) SetTaskFlowType(v int32) *RuleInfo {
	s.TaskFlowType = &v
	return s
}

func (s *RuleInfo) SetTriggers(v []*string) *RuleInfo {
	s.Triggers = v
	return s
}

func (s *RuleInfo) SetType(v int32) *RuleInfo {
	s.Type = &v
	return s
}

func (s *RuleInfo) SetWeight(v string) *RuleInfo {
	s.Weight = &v
	return s
}

type RuleTestDialogue struct {
	Content   []*RuleTestDialogueContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	Id        *int64                     `json:"Id,omitempty" xml:"Id,omitempty"`
	Name      *string                    `json:"Name,omitempty" xml:"Name,omitempty"`
	UserGroup *string                    `json:"UserGroup,omitempty" xml:"UserGroup,omitempty"`
}

func (s RuleTestDialogue) String() string {
	return tea.Prettify(s)
}

func (s RuleTestDialogue) GoString() string {
	return s.String()
}

func (s *RuleTestDialogue) SetContent(v []*RuleTestDialogueContent) *RuleTestDialogue {
	s.Content = v
	return s
}

func (s *RuleTestDialogue) SetId(v int64) *RuleTestDialogue {
	s.Id = &v
	return s
}

func (s *RuleTestDialogue) SetName(v string) *RuleTestDialogue {
	s.Name = &v
	return s
}

func (s *RuleTestDialogue) SetUserGroup(v string) *RuleTestDialogue {
	s.UserGroup = &v
	return s
}

type RuleTestDialogueContent struct {
	Begin           *int64  `json:"Begin,omitempty" xml:"Begin,omitempty"`
	BeginTime       *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EmotionValue    *int32  `json:"EmotionValue,omitempty" xml:"EmotionValue,omitempty"`
	End             *int64  `json:"End,omitempty" xml:"End,omitempty"`
	HourMinSec      *string `json:"HourMinSec,omitempty" xml:"HourMinSec,omitempty"`
	Identity        *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
	Role            *string `json:"Role,omitempty" xml:"Role,omitempty"`
	SilenceDuration *int64  `json:"SilenceDuration,omitempty" xml:"SilenceDuration,omitempty"`
	SpeechRate      *int64  `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	Words           *string `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s RuleTestDialogueContent) String() string {
	return tea.Prettify(s)
}

func (s RuleTestDialogueContent) GoString() string {
	return s.String()
}

func (s *RuleTestDialogueContent) SetBegin(v int64) *RuleTestDialogueContent {
	s.Begin = &v
	return s
}

func (s *RuleTestDialogueContent) SetBeginTime(v int64) *RuleTestDialogueContent {
	s.BeginTime = &v
	return s
}

func (s *RuleTestDialogueContent) SetEmotionValue(v int32) *RuleTestDialogueContent {
	s.EmotionValue = &v
	return s
}

func (s *RuleTestDialogueContent) SetEnd(v int64) *RuleTestDialogueContent {
	s.End = &v
	return s
}

func (s *RuleTestDialogueContent) SetHourMinSec(v string) *RuleTestDialogueContent {
	s.HourMinSec = &v
	return s
}

func (s *RuleTestDialogueContent) SetIdentity(v string) *RuleTestDialogueContent {
	s.Identity = &v
	return s
}

func (s *RuleTestDialogueContent) SetRole(v string) *RuleTestDialogueContent {
	s.Role = &v
	return s
}

func (s *RuleTestDialogueContent) SetSilenceDuration(v int64) *RuleTestDialogueContent {
	s.SilenceDuration = &v
	return s
}

func (s *RuleTestDialogueContent) SetSpeechRate(v int64) *RuleTestDialogueContent {
	s.SpeechRate = &v
	return s
}

func (s *RuleTestDialogueContent) SetWords(v string) *RuleTestDialogueContent {
	s.Words = &v
	return s
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
	return tea.Prettify(s)
}

func (s RulesInfo) GoString() string {
	return s.String()
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

type SchemeCheckType struct {
	CheckName             *string                                 `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	CheckType             *int64                                  `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	Enable                *int32                                  `json:"Enable,omitempty" xml:"Enable,omitempty"`
	SchemeId              *int64                                  `json:"SchemeId,omitempty" xml:"SchemeId,omitempty"`
	SchemeScoreInfoList   []*SchemeCheckTypeSchemeScoreInfoList   `json:"SchemeScoreInfoList,omitempty" xml:"SchemeScoreInfoList,omitempty" type:"Repeated"`
	Score                 *int32                                  `json:"Score,omitempty" xml:"Score,omitempty"`
	SourceScore           *int32                                  `json:"SourceScore,omitempty" xml:"SourceScore,omitempty"`
	TaskFlowScoreInfoList []*SchemeCheckTypeTaskFlowScoreInfoList `json:"TaskFlowScoreInfoList,omitempty" xml:"TaskFlowScoreInfoList,omitempty" type:"Repeated"`
}

func (s SchemeCheckType) String() string {
	return tea.Prettify(s)
}

func (s SchemeCheckType) GoString() string {
	return s.String()
}

func (s *SchemeCheckType) SetCheckName(v string) *SchemeCheckType {
	s.CheckName = &v
	return s
}

func (s *SchemeCheckType) SetCheckType(v int64) *SchemeCheckType {
	s.CheckType = &v
	return s
}

func (s *SchemeCheckType) SetEnable(v int32) *SchemeCheckType {
	s.Enable = &v
	return s
}

func (s *SchemeCheckType) SetSchemeId(v int64) *SchemeCheckType {
	s.SchemeId = &v
	return s
}

func (s *SchemeCheckType) SetSchemeScoreInfoList(v []*SchemeCheckTypeSchemeScoreInfoList) *SchemeCheckType {
	s.SchemeScoreInfoList = v
	return s
}

func (s *SchemeCheckType) SetScore(v int32) *SchemeCheckType {
	s.Score = &v
	return s
}

func (s *SchemeCheckType) SetSourceScore(v int32) *SchemeCheckType {
	s.SourceScore = &v
	return s
}

func (s *SchemeCheckType) SetTaskFlowScoreInfoList(v []*SchemeCheckTypeTaskFlowScoreInfoList) *SchemeCheckType {
	s.TaskFlowScoreInfoList = v
	return s
}

type SchemeCheckTypeSchemeScoreInfoList struct {
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Rid              *int64  `json:"Rid,omitempty" xml:"Rid,omitempty"`
	ScoreNum         *int32  `json:"ScoreNum,omitempty" xml:"ScoreNum,omitempty"`
	ScoreNumType     *int32  `json:"ScoreNumType,omitempty" xml:"ScoreNumType,omitempty"`
	ScoreRuleHitType *int32  `json:"ScoreRuleHitType,omitempty" xml:"ScoreRuleHitType,omitempty"`
	ScoreType        *int32  `json:"ScoreType,omitempty" xml:"ScoreType,omitempty"`
	TaskFlowId       *int64  `json:"TaskFlowId,omitempty" xml:"TaskFlowId,omitempty"`
	TaskFlowName     *string `json:"TaskFlowName,omitempty" xml:"TaskFlowName,omitempty"`
}

func (s SchemeCheckTypeSchemeScoreInfoList) String() string {
	return tea.Prettify(s)
}

func (s SchemeCheckTypeSchemeScoreInfoList) GoString() string {
	return s.String()
}

func (s *SchemeCheckTypeSchemeScoreInfoList) SetName(v string) *SchemeCheckTypeSchemeScoreInfoList {
	s.Name = &v
	return s
}

func (s *SchemeCheckTypeSchemeScoreInfoList) SetRid(v int64) *SchemeCheckTypeSchemeScoreInfoList {
	s.Rid = &v
	return s
}

func (s *SchemeCheckTypeSchemeScoreInfoList) SetScoreNum(v int32) *SchemeCheckTypeSchemeScoreInfoList {
	s.ScoreNum = &v
	return s
}

func (s *SchemeCheckTypeSchemeScoreInfoList) SetScoreNumType(v int32) *SchemeCheckTypeSchemeScoreInfoList {
	s.ScoreNumType = &v
	return s
}

func (s *SchemeCheckTypeSchemeScoreInfoList) SetScoreRuleHitType(v int32) *SchemeCheckTypeSchemeScoreInfoList {
	s.ScoreRuleHitType = &v
	return s
}

func (s *SchemeCheckTypeSchemeScoreInfoList) SetScoreType(v int32) *SchemeCheckTypeSchemeScoreInfoList {
	s.ScoreType = &v
	return s
}

func (s *SchemeCheckTypeSchemeScoreInfoList) SetTaskFlowId(v int64) *SchemeCheckTypeSchemeScoreInfoList {
	s.TaskFlowId = &v
	return s
}

func (s *SchemeCheckTypeSchemeScoreInfoList) SetTaskFlowName(v string) *SchemeCheckTypeSchemeScoreInfoList {
	s.TaskFlowName = &v
	return s
}

type SchemeCheckTypeTaskFlowScoreInfoList struct {
	SchemeScoreInfoList []*SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList `json:"SchemeScoreInfoList,omitempty" xml:"SchemeScoreInfoList,omitempty" type:"Repeated"`
	TaskFlowId          *int64                                                     `json:"TaskFlowId,omitempty" xml:"TaskFlowId,omitempty"`
	TaskFlowName        *string                                                    `json:"TaskFlowName,omitempty" xml:"TaskFlowName,omitempty"`
	TaskFlowType        *int32                                                     `json:"TaskFlowType,omitempty" xml:"TaskFlowType,omitempty"`
}

func (s SchemeCheckTypeTaskFlowScoreInfoList) String() string {
	return tea.Prettify(s)
}

func (s SchemeCheckTypeTaskFlowScoreInfoList) GoString() string {
	return s.String()
}

func (s *SchemeCheckTypeTaskFlowScoreInfoList) SetSchemeScoreInfoList(v []*SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList) *SchemeCheckTypeTaskFlowScoreInfoList {
	s.SchemeScoreInfoList = v
	return s
}

func (s *SchemeCheckTypeTaskFlowScoreInfoList) SetTaskFlowId(v int64) *SchemeCheckTypeTaskFlowScoreInfoList {
	s.TaskFlowId = &v
	return s
}

func (s *SchemeCheckTypeTaskFlowScoreInfoList) SetTaskFlowName(v string) *SchemeCheckTypeTaskFlowScoreInfoList {
	s.TaskFlowName = &v
	return s
}

func (s *SchemeCheckTypeTaskFlowScoreInfoList) SetTaskFlowType(v int32) *SchemeCheckTypeTaskFlowScoreInfoList {
	s.TaskFlowType = &v
	return s
}

type SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList struct {
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Rid              *int64  `json:"Rid,omitempty" xml:"Rid,omitempty"`
	ScoreNum         *int32  `json:"ScoreNum,omitempty" xml:"ScoreNum,omitempty"`
	ScoreNumType     *int32  `json:"ScoreNumType,omitempty" xml:"ScoreNumType,omitempty"`
	ScoreRuleHitType *int32  `json:"ScoreRuleHitType,omitempty" xml:"ScoreRuleHitType,omitempty"`
	ScoreType        *int32  `json:"ScoreType,omitempty" xml:"ScoreType,omitempty"`
	TaskFlowId       *int64  `json:"TaskFlowId,omitempty" xml:"TaskFlowId,omitempty"`
	TaskFlowName     *string `json:"TaskFlowName,omitempty" xml:"TaskFlowName,omitempty"`
}

func (s SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList) String() string {
	return tea.Prettify(s)
}

func (s SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList) GoString() string {
	return s.String()
}

func (s *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList) SetName(v string) *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList {
	s.Name = &v
	return s
}

func (s *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList) SetRid(v int64) *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList {
	s.Rid = &v
	return s
}

func (s *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList) SetScoreNum(v int32) *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList {
	s.ScoreNum = &v
	return s
}

func (s *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList) SetScoreNumType(v int32) *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList {
	s.ScoreNumType = &v
	return s
}

func (s *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList) SetScoreRuleHitType(v int32) *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList {
	s.ScoreRuleHitType = &v
	return s
}

func (s *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList) SetScoreType(v int32) *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList {
	s.ScoreType = &v
	return s
}

func (s *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList) SetTaskFlowId(v int64) *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList {
	s.TaskFlowId = &v
	return s
}

func (s *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList) SetTaskFlowName(v string) *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList {
	s.TaskFlowName = &v
	return s
}

type TaskGraphFlow struct {
	FlowRuleScoreType            *int32           `json:"FlowRuleScoreType,omitempty" xml:"FlowRuleScoreType,omitempty"`
	Id                           *int64           `json:"Id,omitempty" xml:"Id,omitempty"`
	Nodes                        []*GraphFlowNode `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	Rid                          *int64           `json:"Rid,omitempty" xml:"Rid,omitempty"`
	RuleName                     *string          `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	ShowProperties               *string          `json:"ShowProperties,omitempty" xml:"ShowProperties,omitempty"`
	SkipWhenFirstSessionNodeMiss *bool            `json:"SkipWhenFirstSessionNodeMiss,omitempty" xml:"SkipWhenFirstSessionNodeMiss,omitempty"`
}

func (s TaskGraphFlow) String() string {
	return tea.Prettify(s)
}

func (s TaskGraphFlow) GoString() string {
	return s.String()
}

func (s *TaskGraphFlow) SetFlowRuleScoreType(v int32) *TaskGraphFlow {
	s.FlowRuleScoreType = &v
	return s
}

func (s *TaskGraphFlow) SetId(v int64) *TaskGraphFlow {
	s.Id = &v
	return s
}

func (s *TaskGraphFlow) SetNodes(v []*GraphFlowNode) *TaskGraphFlow {
	s.Nodes = v
	return s
}

func (s *TaskGraphFlow) SetRid(v int64) *TaskGraphFlow {
	s.Rid = &v
	return s
}

func (s *TaskGraphFlow) SetRuleName(v string) *TaskGraphFlow {
	s.RuleName = &v
	return s
}

func (s *TaskGraphFlow) SetShowProperties(v string) *TaskGraphFlow {
	s.ShowProperties = &v
	return s
}

func (s *TaskGraphFlow) SetSkipWhenFirstSessionNodeMiss(v bool) *TaskGraphFlow {
	s.SkipWhenFirstSessionNodeMiss = &v
	return s
}

type AddBusinessCategoryRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s AddBusinessCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s AddBusinessCategoryRequest) GoString() string {
	return s.String()
}

func (s *AddBusinessCategoryRequest) SetBaseMeAgentId(v int64) *AddBusinessCategoryRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *AddBusinessCategoryRequest) SetJsonStr(v string) *AddBusinessCategoryRequest {
	s.JsonStr = &v
	return s
}

type AddBusinessCategoryResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 348193421
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 76DB5D8C-5BD9-42A7-B527-5AF3A5F83F12
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddBusinessCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddBusinessCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *AddBusinessCategoryResponseBody) SetCode(v string) *AddBusinessCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *AddBusinessCategoryResponseBody) SetData(v string) *AddBusinessCategoryResponseBody {
	s.Data = &v
	return s
}

func (s *AddBusinessCategoryResponseBody) SetMessage(v string) *AddBusinessCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *AddBusinessCategoryResponseBody) SetRequestId(v string) *AddBusinessCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddBusinessCategoryResponseBody) SetSuccess(v bool) *AddBusinessCategoryResponseBody {
	s.Success = &v
	return s
}

type AddBusinessCategoryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddBusinessCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddBusinessCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s AddBusinessCategoryResponse) GoString() string {
	return s.String()
}

func (s *AddBusinessCategoryResponse) SetHeaders(v map[string]*string) *AddBusinessCategoryResponse {
	s.Headers = v
	return s
}

func (s *AddBusinessCategoryResponse) SetStatusCode(v int32) *AddBusinessCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *AddBusinessCategoryResponse) SetBody(v *AddBusinessCategoryResponseBody) *AddBusinessCategoryResponse {
	s.Body = v
	return s
}

type AddRuleCategoryRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s AddRuleCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s AddRuleCategoryRequest) GoString() string {
	return s.String()
}

func (s *AddRuleCategoryRequest) SetBaseMeAgentId(v int64) *AddRuleCategoryRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *AddRuleCategoryRequest) SetJsonStr(v string) *AddRuleCategoryRequest {
	s.JsonStr = &v
	return s
}

type AddRuleCategoryResponseBody struct {
	// example:
	//
	// 200
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *AddRuleCategoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddRuleCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddRuleCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *AddRuleCategoryResponseBody) SetCode(v string) *AddRuleCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *AddRuleCategoryResponseBody) SetData(v *AddRuleCategoryResponseBodyData) *AddRuleCategoryResponseBody {
	s.Data = v
	return s
}

func (s *AddRuleCategoryResponseBody) SetMessage(v string) *AddRuleCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *AddRuleCategoryResponseBody) SetRequestId(v string) *AddRuleCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRuleCategoryResponseBody) SetSuccess(v bool) *AddRuleCategoryResponseBody {
	s.Success = &v
	return s
}

type AddRuleCategoryResponseBodyData struct {
	Select *bool  `json:"Select,omitempty" xml:"Select,omitempty"`
	Type   *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddRuleCategoryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddRuleCategoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddRuleCategoryResponseBodyData) SetSelect(v bool) *AddRuleCategoryResponseBodyData {
	s.Select = &v
	return s
}

func (s *AddRuleCategoryResponseBodyData) SetType(v int32) *AddRuleCategoryResponseBodyData {
	s.Type = &v
	return s
}

type AddRuleCategoryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddRuleCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddRuleCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s AddRuleCategoryResponse) GoString() string {
	return s.String()
}

func (s *AddRuleCategoryResponse) SetHeaders(v map[string]*string) *AddRuleCategoryResponse {
	s.Headers = v
	return s
}

func (s *AddRuleCategoryResponse) SetStatusCode(v int32) *AddRuleCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *AddRuleCategoryResponse) SetBody(v *AddRuleCategoryResponseBody) *AddRuleCategoryResponse {
	s.Body = v
	return s
}

type AddRuleV4Request struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// false
	IsCopy *bool `json:"IsCopy,omitempty" xml:"IsCopy,omitempty"`
	// This parameter is required.
	JsonStrForRule *string `json:"JsonStrForRule,omitempty" xml:"JsonStrForRule,omitempty"`
}

func (s AddRuleV4Request) String() string {
	return tea.Prettify(s)
}

func (s AddRuleV4Request) GoString() string {
	return s.String()
}

func (s *AddRuleV4Request) SetBaseMeAgentId(v int64) *AddRuleV4Request {
	s.BaseMeAgentId = &v
	return s
}

func (s *AddRuleV4Request) SetIsCopy(v bool) *AddRuleV4Request {
	s.IsCopy = &v
	return s
}

func (s *AddRuleV4Request) SetJsonStrForRule(v string) *AddRuleV4Request {
	s.JsonStrForRule = &v
	return s
}

type AddRuleV4ResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *AddRuleV4ResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddRuleV4ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddRuleV4ResponseBody) GoString() string {
	return s.String()
}

func (s *AddRuleV4ResponseBody) SetCode(v string) *AddRuleV4ResponseBody {
	s.Code = &v
	return s
}

func (s *AddRuleV4ResponseBody) SetData(v int64) *AddRuleV4ResponseBody {
	s.Data = &v
	return s
}

func (s *AddRuleV4ResponseBody) SetHttpStatusCode(v int32) *AddRuleV4ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddRuleV4ResponseBody) SetMessage(v string) *AddRuleV4ResponseBody {
	s.Message = &v
	return s
}

func (s *AddRuleV4ResponseBody) SetMessages(v *AddRuleV4ResponseBodyMessages) *AddRuleV4ResponseBody {
	s.Messages = v
	return s
}

func (s *AddRuleV4ResponseBody) SetRequestId(v string) *AddRuleV4ResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRuleV4ResponseBody) SetSuccess(v bool) *AddRuleV4ResponseBody {
	s.Success = &v
	return s
}

type AddRuleV4ResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s AddRuleV4ResponseBodyMessages) String() string {
	return tea.Prettify(s)
}

func (s AddRuleV4ResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *AddRuleV4ResponseBodyMessages) SetMessage(v []*string) *AddRuleV4ResponseBodyMessages {
	s.Message = v
	return s
}

type AddRuleV4Response struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddRuleV4ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddRuleV4Response) String() string {
	return tea.Prettify(s)
}

func (s AddRuleV4Response) GoString() string {
	return s.String()
}

func (s *AddRuleV4Response) SetHeaders(v map[string]*string) *AddRuleV4Response {
	s.Headers = v
	return s
}

func (s *AddRuleV4Response) SetStatusCode(v int32) *AddRuleV4Response {
	s.StatusCode = &v
	return s
}

func (s *AddRuleV4Response) SetBody(v *AddRuleV4ResponseBody) *AddRuleV4Response {
	s.Body = v
	return s
}

type AssignReviewerRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"assignmentList":[{"taskId":"1C21CF1E-2917-4236-A046-20E37B293B69","fileId":"7981b466fd6a4c70a7065da159739a5b"},{"taskId":"0111DDBC-5F10-47E0-B7D4-7175F47D626F","fileId":"1814eeae3cff41e989e31ab547f07561"}],"assignments":[{"reviewer":"255746168704895558"},{"reviewer":"268370362815185444"}]}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s AssignReviewerRequest) String() string {
	return tea.Prettify(s)
}

func (s AssignReviewerRequest) GoString() string {
	return s.String()
}

func (s *AssignReviewerRequest) SetBaseMeAgentId(v int64) *AssignReviewerRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *AssignReviewerRequest) SetJsonStr(v string) *AssignReviewerRequest {
	s.JsonStr = &v
	return s
}

type AssignReviewerResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 82C91484-B2D5-4D2A-A21F-A6D73F4D55C6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AssignReviewerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssignReviewerResponseBody) GoString() string {
	return s.String()
}

func (s *AssignReviewerResponseBody) SetCode(v string) *AssignReviewerResponseBody {
	s.Code = &v
	return s
}

func (s *AssignReviewerResponseBody) SetMessage(v string) *AssignReviewerResponseBody {
	s.Message = &v
	return s
}

func (s *AssignReviewerResponseBody) SetRequestId(v string) *AssignReviewerResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssignReviewerResponseBody) SetSuccess(v bool) *AssignReviewerResponseBody {
	s.Success = &v
	return s
}

type AssignReviewerResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssignReviewerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssignReviewerResponse) String() string {
	return tea.Prettify(s)
}

func (s AssignReviewerResponse) GoString() string {
	return s.String()
}

func (s *AssignReviewerResponse) SetHeaders(v map[string]*string) *AssignReviewerResponse {
	s.Headers = v
	return s
}

func (s *AssignReviewerResponse) SetStatusCode(v int32) *AssignReviewerResponse {
	s.StatusCode = &v
	return s
}

func (s *AssignReviewerResponse) SetBody(v *AssignReviewerResponseBody) *AssignReviewerResponse {
	s.Body = v
	return s
}

type AssignReviewerBySessionGroupRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// {"sessionGroupParam":{"isSchemeData":1,"callStartTime":"2022-09-17 00:00:00","callEndTime":"2022-09-23 23:59:59","schemeTaskConfigId":24},"assignments":[{"reviewer":63,"count":4}],"isSchemeData":1}
	JsonStr *string `json:"jsonStr,omitempty" xml:"jsonStr,omitempty"`
}

func (s AssignReviewerBySessionGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AssignReviewerBySessionGroupRequest) GoString() string {
	return s.String()
}

func (s *AssignReviewerBySessionGroupRequest) SetBaseMeAgentId(v int64) *AssignReviewerBySessionGroupRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *AssignReviewerBySessionGroupRequest) SetJsonStr(v string) *AssignReviewerBySessionGroupRequest {
	s.JsonStr = &v
	return s
}

type AssignReviewerBySessionGroupResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *AssignReviewerBySessionGroupResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// F190ADE9-619A-447D-84E3-7E241A5C428E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AssignReviewerBySessionGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssignReviewerBySessionGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AssignReviewerBySessionGroupResponseBody) SetCode(v string) *AssignReviewerBySessionGroupResponseBody {
	s.Code = &v
	return s
}

func (s *AssignReviewerBySessionGroupResponseBody) SetHttpStatusCode(v int32) *AssignReviewerBySessionGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AssignReviewerBySessionGroupResponseBody) SetMessage(v string) *AssignReviewerBySessionGroupResponseBody {
	s.Message = &v
	return s
}

func (s *AssignReviewerBySessionGroupResponseBody) SetMessages(v *AssignReviewerBySessionGroupResponseBodyMessages) *AssignReviewerBySessionGroupResponseBody {
	s.Messages = v
	return s
}

func (s *AssignReviewerBySessionGroupResponseBody) SetRequestId(v string) *AssignReviewerBySessionGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssignReviewerBySessionGroupResponseBody) SetSuccess(v bool) *AssignReviewerBySessionGroupResponseBody {
	s.Success = &v
	return s
}

type AssignReviewerBySessionGroupResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s AssignReviewerBySessionGroupResponseBodyMessages) String() string {
	return tea.Prettify(s)
}

func (s AssignReviewerBySessionGroupResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *AssignReviewerBySessionGroupResponseBodyMessages) SetMessage(v []*string) *AssignReviewerBySessionGroupResponseBodyMessages {
	s.Message = v
	return s
}

type AssignReviewerBySessionGroupResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssignReviewerBySessionGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssignReviewerBySessionGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AssignReviewerBySessionGroupResponse) GoString() string {
	return s.String()
}

func (s *AssignReviewerBySessionGroupResponse) SetHeaders(v map[string]*string) *AssignReviewerBySessionGroupResponse {
	s.Headers = v
	return s
}

func (s *AssignReviewerBySessionGroupResponse) SetStatusCode(v int32) *AssignReviewerBySessionGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AssignReviewerBySessionGroupResponse) SetBody(v *AssignReviewerBySessionGroupResponseBody) *AssignReviewerBySessionGroupResponse {
	s.Body = v
	return s
}

type BatchSubmitReviewInfoRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// {"isSchemeData":1,"schemeTaskConfigId":334,"sourceDataType":2,"startTime":"2022-08-25 00:00:00","endTime":"2022-09-23 23:59:59","sessionList":[{"taskId":"20220831-F8D7F4DF-0A16-1A1C-BA63-28F203922692","fileId":"20220831-164343"},{"taskId":"20220831-F2A50A72-82C4-1E3F-A1FD-52A662283D25","fileId":"20220831-164343"}]}
	JsonStr *string `json:"jsonStr,omitempty" xml:"jsonStr,omitempty"`
}

func (s BatchSubmitReviewInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchSubmitReviewInfoRequest) GoString() string {
	return s.String()
}

func (s *BatchSubmitReviewInfoRequest) SetBaseMeAgentId(v int64) *BatchSubmitReviewInfoRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *BatchSubmitReviewInfoRequest) SetJsonStr(v string) *BatchSubmitReviewInfoRequest {
	s.JsonStr = &v
	return s
}

type BatchSubmitReviewInfoResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *BatchSubmitReviewInfoResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// F190ADE9-619A-447D-84E3-7E241A5C428E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchSubmitReviewInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchSubmitReviewInfoResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSubmitReviewInfoResponseBody) SetCode(v string) *BatchSubmitReviewInfoResponseBody {
	s.Code = &v
	return s
}

func (s *BatchSubmitReviewInfoResponseBody) SetHttpStatusCode(v int32) *BatchSubmitReviewInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BatchSubmitReviewInfoResponseBody) SetMessage(v string) *BatchSubmitReviewInfoResponseBody {
	s.Message = &v
	return s
}

func (s *BatchSubmitReviewInfoResponseBody) SetMessages(v *BatchSubmitReviewInfoResponseBodyMessages) *BatchSubmitReviewInfoResponseBody {
	s.Messages = v
	return s
}

func (s *BatchSubmitReviewInfoResponseBody) SetRequestId(v string) *BatchSubmitReviewInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchSubmitReviewInfoResponseBody) SetSuccess(v bool) *BatchSubmitReviewInfoResponseBody {
	s.Success = &v
	return s
}

type BatchSubmitReviewInfoResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s BatchSubmitReviewInfoResponseBodyMessages) String() string {
	return tea.Prettify(s)
}

func (s BatchSubmitReviewInfoResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *BatchSubmitReviewInfoResponseBodyMessages) SetMessage(v []*string) *BatchSubmitReviewInfoResponseBodyMessages {
	s.Message = v
	return s
}

type BatchSubmitReviewInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchSubmitReviewInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchSubmitReviewInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchSubmitReviewInfoResponse) GoString() string {
	return s.String()
}

func (s *BatchSubmitReviewInfoResponse) SetHeaders(v map[string]*string) *BatchSubmitReviewInfoResponse {
	s.Headers = v
	return s
}

func (s *BatchSubmitReviewInfoResponse) SetStatusCode(v int32) *BatchSubmitReviewInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchSubmitReviewInfoResponse) SetBody(v *BatchSubmitReviewInfoResponseBody) *BatchSubmitReviewInfoResponse {
	s.Body = v
	return s
}

type CreateAsrVocabRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s CreateAsrVocabRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAsrVocabRequest) GoString() string {
	return s.String()
}

func (s *CreateAsrVocabRequest) SetBaseMeAgentId(v int64) *CreateAsrVocabRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *CreateAsrVocabRequest) SetJsonStr(v string) *CreateAsrVocabRequest {
	s.JsonStr = &v
	return s
}

type CreateAsrVocabResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 71b1795ac8634bd8bdf4d3878480c7c2
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 96138D8D-8D26-4E41-BFF4-77AED1088BBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAsrVocabResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAsrVocabResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAsrVocabResponseBody) SetCode(v string) *CreateAsrVocabResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAsrVocabResponseBody) SetData(v string) *CreateAsrVocabResponseBody {
	s.Data = &v
	return s
}

func (s *CreateAsrVocabResponseBody) SetMessage(v string) *CreateAsrVocabResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAsrVocabResponseBody) SetRequestId(v string) *CreateAsrVocabResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAsrVocabResponseBody) SetSuccess(v bool) *CreateAsrVocabResponseBody {
	s.Success = &v
	return s
}

type CreateAsrVocabResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAsrVocabResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAsrVocabResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAsrVocabResponse) GoString() string {
	return s.String()
}

func (s *CreateAsrVocabResponse) SetHeaders(v map[string]*string) *CreateAsrVocabResponse {
	s.Headers = v
	return s
}

func (s *CreateAsrVocabResponse) SetStatusCode(v int32) *CreateAsrVocabResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAsrVocabResponse) SetBody(v *CreateAsrVocabResponseBody) *CreateAsrVocabResponse {
	s.Body = v
	return s
}

type CreateCheckTypeToSchemeRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64  `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	JsonStr       *string `json:"jsonStr,omitempty" xml:"jsonStr,omitempty"`
}

func (s CreateCheckTypeToSchemeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCheckTypeToSchemeRequest) GoString() string {
	return s.String()
}

func (s *CreateCheckTypeToSchemeRequest) SetBaseMeAgentId(v int64) *CreateCheckTypeToSchemeRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *CreateCheckTypeToSchemeRequest) SetJsonStr(v string) *CreateCheckTypeToSchemeRequest {
	s.JsonStr = &v
	return s
}

type CreateCheckTypeToSchemeResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 5
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *CreateCheckTypeToSchemeResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24F4CE647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCheckTypeToSchemeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCheckTypeToSchemeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCheckTypeToSchemeResponseBody) SetCode(v string) *CreateCheckTypeToSchemeResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCheckTypeToSchemeResponseBody) SetData(v int64) *CreateCheckTypeToSchemeResponseBody {
	s.Data = &v
	return s
}

func (s *CreateCheckTypeToSchemeResponseBody) SetHttpStatusCode(v int32) *CreateCheckTypeToSchemeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateCheckTypeToSchemeResponseBody) SetMessage(v string) *CreateCheckTypeToSchemeResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCheckTypeToSchemeResponseBody) SetMessages(v *CreateCheckTypeToSchemeResponseBodyMessages) *CreateCheckTypeToSchemeResponseBody {
	s.Messages = v
	return s
}

func (s *CreateCheckTypeToSchemeResponseBody) SetRequestId(v string) *CreateCheckTypeToSchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCheckTypeToSchemeResponseBody) SetSuccess(v bool) *CreateCheckTypeToSchemeResponseBody {
	s.Success = &v
	return s
}

type CreateCheckTypeToSchemeResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s CreateCheckTypeToSchemeResponseBodyMessages) String() string {
	return tea.Prettify(s)
}

func (s CreateCheckTypeToSchemeResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *CreateCheckTypeToSchemeResponseBodyMessages) SetMessage(v []*string) *CreateCheckTypeToSchemeResponseBodyMessages {
	s.Message = v
	return s
}

type CreateCheckTypeToSchemeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCheckTypeToSchemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCheckTypeToSchemeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCheckTypeToSchemeResponse) GoString() string {
	return s.String()
}

func (s *CreateCheckTypeToSchemeResponse) SetHeaders(v map[string]*string) *CreateCheckTypeToSchemeResponse {
	s.Headers = v
	return s
}

func (s *CreateCheckTypeToSchemeResponse) SetStatusCode(v int32) *CreateCheckTypeToSchemeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCheckTypeToSchemeResponse) SetBody(v *CreateCheckTypeToSchemeResponseBody) *CreateCheckTypeToSchemeResponse {
	s.Body = v
	return s
}

type CreateQualityCheckSchemeRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"jsonStr,omitempty" xml:"jsonStr,omitempty"`
}

func (s CreateQualityCheckSchemeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateQualityCheckSchemeRequest) GoString() string {
	return s.String()
}

func (s *CreateQualityCheckSchemeRequest) SetBaseMeAgentId(v int64) *CreateQualityCheckSchemeRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *CreateQualityCheckSchemeRequest) SetJsonStr(v string) *CreateQualityCheckSchemeRequest {
	s.JsonStr = &v
	return s
}

type CreateQualityCheckSchemeResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 12
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *CreateQualityCheckSchemeResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// 76DB5D8C-5BD9-42A7-B527-5AF3A5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateQualityCheckSchemeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateQualityCheckSchemeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQualityCheckSchemeResponseBody) SetCode(v string) *CreateQualityCheckSchemeResponseBody {
	s.Code = &v
	return s
}

func (s *CreateQualityCheckSchemeResponseBody) SetData(v int64) *CreateQualityCheckSchemeResponseBody {
	s.Data = &v
	return s
}

func (s *CreateQualityCheckSchemeResponseBody) SetHttpStatusCode(v int32) *CreateQualityCheckSchemeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateQualityCheckSchemeResponseBody) SetMessage(v string) *CreateQualityCheckSchemeResponseBody {
	s.Message = &v
	return s
}

func (s *CreateQualityCheckSchemeResponseBody) SetMessages(v *CreateQualityCheckSchemeResponseBodyMessages) *CreateQualityCheckSchemeResponseBody {
	s.Messages = v
	return s
}

func (s *CreateQualityCheckSchemeResponseBody) SetRequestId(v string) *CreateQualityCheckSchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateQualityCheckSchemeResponseBody) SetSuccess(v bool) *CreateQualityCheckSchemeResponseBody {
	s.Success = &v
	return s
}

type CreateQualityCheckSchemeResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s CreateQualityCheckSchemeResponseBodyMessages) String() string {
	return tea.Prettify(s)
}

func (s CreateQualityCheckSchemeResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *CreateQualityCheckSchemeResponseBodyMessages) SetMessage(v []*string) *CreateQualityCheckSchemeResponseBodyMessages {
	s.Message = v
	return s
}

type CreateQualityCheckSchemeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateQualityCheckSchemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateQualityCheckSchemeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateQualityCheckSchemeResponse) GoString() string {
	return s.String()
}

func (s *CreateQualityCheckSchemeResponse) SetHeaders(v map[string]*string) *CreateQualityCheckSchemeResponse {
	s.Headers = v
	return s
}

func (s *CreateQualityCheckSchemeResponse) SetStatusCode(v int32) *CreateQualityCheckSchemeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateQualityCheckSchemeResponse) SetBody(v *CreateQualityCheckSchemeResponseBody) *CreateQualityCheckSchemeResponse {
	s.Body = v
	return s
}

type CreateSchemeTaskConfigRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64  `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	JsonStr       *string `json:"jsonStr,omitempty" xml:"jsonStr,omitempty"`
}

func (s CreateSchemeTaskConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSchemeTaskConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateSchemeTaskConfigRequest) SetBaseMeAgentId(v int64) *CreateSchemeTaskConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *CreateSchemeTaskConfigRequest) SetJsonStr(v string) *CreateSchemeTaskConfigRequest {
	s.JsonStr = &v
	return s
}

type CreateSchemeTaskConfigResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 22
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *CreateSchemeTaskConfigResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// 3CEA0495-341B-4482-9AD9-8191EF4***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSchemeTaskConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSchemeTaskConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSchemeTaskConfigResponseBody) SetCode(v string) *CreateSchemeTaskConfigResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSchemeTaskConfigResponseBody) SetData(v int64) *CreateSchemeTaskConfigResponseBody {
	s.Data = &v
	return s
}

func (s *CreateSchemeTaskConfigResponseBody) SetHttpStatusCode(v int32) *CreateSchemeTaskConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateSchemeTaskConfigResponseBody) SetMessage(v string) *CreateSchemeTaskConfigResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSchemeTaskConfigResponseBody) SetMessages(v *CreateSchemeTaskConfigResponseBodyMessages) *CreateSchemeTaskConfigResponseBody {
	s.Messages = v
	return s
}

func (s *CreateSchemeTaskConfigResponseBody) SetRequestId(v string) *CreateSchemeTaskConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSchemeTaskConfigResponseBody) SetSuccess(v bool) *CreateSchemeTaskConfigResponseBody {
	s.Success = &v
	return s
}

type CreateSchemeTaskConfigResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s CreateSchemeTaskConfigResponseBodyMessages) String() string {
	return tea.Prettify(s)
}

func (s CreateSchemeTaskConfigResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *CreateSchemeTaskConfigResponseBodyMessages) SetMessage(v []*string) *CreateSchemeTaskConfigResponseBodyMessages {
	s.Message = v
	return s
}

type CreateSchemeTaskConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSchemeTaskConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSchemeTaskConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSchemeTaskConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateSchemeTaskConfigResponse) SetHeaders(v map[string]*string) *CreateSchemeTaskConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateSchemeTaskConfigResponse) SetStatusCode(v int32) *CreateSchemeTaskConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSchemeTaskConfigResponse) SetBody(v *CreateSchemeTaskConfigResponseBody) *CreateSchemeTaskConfigResponse {
	s.Body = v
	return s
}

type CreateSkillGroupConfigRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"skillGroupFrom":0,"qualityCheckType":0,"modelId":746,"name":"test","rid":"2493","vocabId":"267","skillGroupList":[{"skillGroupId":"0903","skillGroupName":"0903"}]}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s CreateSkillGroupConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSkillGroupConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateSkillGroupConfigRequest) SetBaseMeAgentId(v int64) *CreateSkillGroupConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *CreateSkillGroupConfigRequest) SetJsonStr(v string) *CreateSkillGroupConfigRequest {
	s.JsonStr = &v
	return s
}

type CreateSkillGroupConfigResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 223
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3CEA0495-341B-4482-9AD9-8191EF4***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSkillGroupConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSkillGroupConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSkillGroupConfigResponseBody) SetCode(v string) *CreateSkillGroupConfigResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSkillGroupConfigResponseBody) SetData(v int64) *CreateSkillGroupConfigResponseBody {
	s.Data = &v
	return s
}

func (s *CreateSkillGroupConfigResponseBody) SetMessage(v string) *CreateSkillGroupConfigResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSkillGroupConfigResponseBody) SetRequestId(v string) *CreateSkillGroupConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSkillGroupConfigResponseBody) SetSuccess(v bool) *CreateSkillGroupConfigResponseBody {
	s.Success = &v
	return s
}

type CreateSkillGroupConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSkillGroupConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSkillGroupConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSkillGroupConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateSkillGroupConfigResponse) SetHeaders(v map[string]*string) *CreateSkillGroupConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateSkillGroupConfigResponse) SetStatusCode(v int32) *CreateSkillGroupConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSkillGroupConfigResponse) SetBody(v *CreateSkillGroupConfigResponseBody) *CreateSkillGroupConfigResponse {
	s.Body = v
	return s
}

type CreateTaskAssignRuleRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s CreateTaskAssignRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskAssignRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateTaskAssignRuleRequest) SetBaseMeAgentId(v int64) *CreateTaskAssignRuleRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *CreateTaskAssignRuleRequest) SetJsonStr(v string) *CreateTaskAssignRuleRequest {
	s.JsonStr = &v
	return s
}

type CreateTaskAssignRuleResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 54
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24F4CE647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTaskAssignRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskAssignRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTaskAssignRuleResponseBody) SetCode(v string) *CreateTaskAssignRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTaskAssignRuleResponseBody) SetData(v string) *CreateTaskAssignRuleResponseBody {
	s.Data = &v
	return s
}

func (s *CreateTaskAssignRuleResponseBody) SetMessage(v string) *CreateTaskAssignRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTaskAssignRuleResponseBody) SetRequestId(v string) *CreateTaskAssignRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTaskAssignRuleResponseBody) SetSuccess(v bool) *CreateTaskAssignRuleResponseBody {
	s.Success = &v
	return s
}

type CreateTaskAssignRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTaskAssignRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTaskAssignRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskAssignRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateTaskAssignRuleResponse) SetHeaders(v map[string]*string) *CreateTaskAssignRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateTaskAssignRuleResponse) SetStatusCode(v int32) *CreateTaskAssignRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTaskAssignRuleResponse) SetBody(v *CreateTaskAssignRuleResponseBody) *CreateTaskAssignRuleResponse {
	s.Body = v
	return s
}

type CreateUserRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s CreateUserRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserRequest) GoString() string {
	return s.String()
}

func (s *CreateUserRequest) SetBaseMeAgentId(v int64) *CreateUserRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *CreateUserRequest) SetJsonStr(v string) *CreateUserRequest {
	s.JsonStr = &v
	return s
}

type CreateUserResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 82C91484-B2D5-4D2A-A21F-A6D73F4***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBody) SetCode(v string) *CreateUserResponseBody {
	s.Code = &v
	return s
}

func (s *CreateUserResponseBody) SetMessage(v string) *CreateUserResponseBody {
	s.Message = &v
	return s
}

func (s *CreateUserResponseBody) SetRequestId(v string) *CreateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserResponseBody) SetSuccess(v bool) *CreateUserResponseBody {
	s.Success = &v
	return s
}

type CreateUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponse) GoString() string {
	return s.String()
}

func (s *CreateUserResponse) SetHeaders(v map[string]*string) *CreateUserResponse {
	s.Headers = v
	return s
}

func (s *CreateUserResponse) SetStatusCode(v int32) *CreateUserResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserResponse) SetBody(v *CreateUserResponseBody) *CreateUserResponse {
	s.Body = v
	return s
}

type CreateWarningConfigRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s CreateWarningConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWarningConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateWarningConfigRequest) SetBaseMeAgentId(v int64) *CreateWarningConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *CreateWarningConfigRequest) SetJsonStr(v string) *CreateWarningConfigRequest {
	s.JsonStr = &v
	return s
}

type CreateWarningConfigResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 30
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 82C91484-B2D5-4D2A-A21F-A6D73F4D55C6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateWarningConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWarningConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWarningConfigResponseBody) SetCode(v string) *CreateWarningConfigResponseBody {
	s.Code = &v
	return s
}

func (s *CreateWarningConfigResponseBody) SetData(v string) *CreateWarningConfigResponseBody {
	s.Data = &v
	return s
}

func (s *CreateWarningConfigResponseBody) SetMessage(v string) *CreateWarningConfigResponseBody {
	s.Message = &v
	return s
}

func (s *CreateWarningConfigResponseBody) SetRequestId(v string) *CreateWarningConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWarningConfigResponseBody) SetSuccess(v bool) *CreateWarningConfigResponseBody {
	s.Success = &v
	return s
}

type CreateWarningConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWarningConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWarningConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWarningConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateWarningConfigResponse) SetHeaders(v map[string]*string) *CreateWarningConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateWarningConfigResponse) SetStatusCode(v int32) *CreateWarningConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWarningConfigResponse) SetBody(v *CreateWarningConfigResponseBody) *CreateWarningConfigResponse {
	s.Body = v
	return s
}

type CreateWarningStrategyConfigRequest struct {
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s CreateWarningStrategyConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWarningStrategyConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateWarningStrategyConfigRequest) SetBaseMeAgentId(v int64) *CreateWarningStrategyConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *CreateWarningStrategyConfigRequest) SetJsonStr(v string) *CreateWarningStrategyConfigRequest {
	s.JsonStr = &v
	return s
}

type CreateWarningStrategyConfigResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateWarningStrategyConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWarningStrategyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWarningStrategyConfigResponseBody) SetCode(v string) *CreateWarningStrategyConfigResponseBody {
	s.Code = &v
	return s
}

func (s *CreateWarningStrategyConfigResponseBody) SetData(v int64) *CreateWarningStrategyConfigResponseBody {
	s.Data = &v
	return s
}

func (s *CreateWarningStrategyConfigResponseBody) SetMessage(v string) *CreateWarningStrategyConfigResponseBody {
	s.Message = &v
	return s
}

func (s *CreateWarningStrategyConfigResponseBody) SetRequestId(v string) *CreateWarningStrategyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWarningStrategyConfigResponseBody) SetSuccess(v bool) *CreateWarningStrategyConfigResponseBody {
	s.Success = &v
	return s
}

type CreateWarningStrategyConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWarningStrategyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWarningStrategyConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWarningStrategyConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateWarningStrategyConfigResponse) SetHeaders(v map[string]*string) *CreateWarningStrategyConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateWarningStrategyConfigResponse) SetStatusCode(v int32) *CreateWarningStrategyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWarningStrategyConfigResponse) SetBody(v *CreateWarningStrategyConfigResponseBody) *CreateWarningStrategyConfigResponse {
	s.Body = v
	return s
}

type DelRuleCategoryRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DelRuleCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DelRuleCategoryRequest) GoString() string {
	return s.String()
}

func (s *DelRuleCategoryRequest) SetBaseMeAgentId(v int64) *DelRuleCategoryRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *DelRuleCategoryRequest) SetJsonStr(v string) *DelRuleCategoryRequest {
	s.JsonStr = &v
	return s
}

type DelRuleCategoryResponseBody struct {
	// example:
	//
	// 200
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DelRuleCategoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DelRuleCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DelRuleCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *DelRuleCategoryResponseBody) SetCode(v string) *DelRuleCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *DelRuleCategoryResponseBody) SetData(v *DelRuleCategoryResponseBodyData) *DelRuleCategoryResponseBody {
	s.Data = v
	return s
}

func (s *DelRuleCategoryResponseBody) SetMessage(v string) *DelRuleCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *DelRuleCategoryResponseBody) SetRequestId(v string) *DelRuleCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DelRuleCategoryResponseBody) SetSuccess(v bool) *DelRuleCategoryResponseBody {
	s.Success = &v
	return s
}

type DelRuleCategoryResponseBodyData struct {
	// example:
	//
	// false
	Select *bool `json:"Select,omitempty" xml:"Select,omitempty"`
}

func (s DelRuleCategoryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DelRuleCategoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *DelRuleCategoryResponseBodyData) SetSelect(v bool) *DelRuleCategoryResponseBodyData {
	s.Select = &v
	return s
}

type DelRuleCategoryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DelRuleCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DelRuleCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DelRuleCategoryResponse) GoString() string {
	return s.String()
}

func (s *DelRuleCategoryResponse) SetHeaders(v map[string]*string) *DelRuleCategoryResponse {
	s.Headers = v
	return s
}

func (s *DelRuleCategoryResponse) SetStatusCode(v int32) *DelRuleCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DelRuleCategoryResponse) SetBody(v *DelRuleCategoryResponseBody) *DelRuleCategoryResponse {
	s.Body = v
	return s
}

type DeleteAsrVocabRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeleteAsrVocabRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAsrVocabRequest) GoString() string {
	return s.String()
}

func (s *DeleteAsrVocabRequest) SetBaseMeAgentId(v int64) *DeleteAsrVocabRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *DeleteAsrVocabRequest) SetJsonStr(v string) *DeleteAsrVocabRequest {
	s.JsonStr = &v
	return s
}

type DeleteAsrVocabResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 71b1795ac8634bd8bdf4d3878480c7c2
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4987D326-83D9-4A42-B9A5-0B27F9B40539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAsrVocabResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAsrVocabResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAsrVocabResponseBody) SetCode(v string) *DeleteAsrVocabResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAsrVocabResponseBody) SetData(v string) *DeleteAsrVocabResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteAsrVocabResponseBody) SetMessage(v string) *DeleteAsrVocabResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAsrVocabResponseBody) SetRequestId(v string) *DeleteAsrVocabResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAsrVocabResponseBody) SetSuccess(v bool) *DeleteAsrVocabResponseBody {
	s.Success = &v
	return s
}

type DeleteAsrVocabResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAsrVocabResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAsrVocabResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAsrVocabResponse) GoString() string {
	return s.String()
}

func (s *DeleteAsrVocabResponse) SetHeaders(v map[string]*string) *DeleteAsrVocabResponse {
	s.Headers = v
	return s
}

func (s *DeleteAsrVocabResponse) SetStatusCode(v int32) *DeleteAsrVocabResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAsrVocabResponse) SetBody(v *DeleteAsrVocabResponseBody) *DeleteAsrVocabResponse {
	s.Body = v
	return s
}

type DeleteBusinessCategoryRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeleteBusinessCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBusinessCategoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteBusinessCategoryRequest) SetBaseMeAgentId(v int64) *DeleteBusinessCategoryRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *DeleteBusinessCategoryRequest) SetJsonStr(v string) *DeleteBusinessCategoryRequest {
	s.JsonStr = &v
	return s
}

type DeleteBusinessCategoryResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4987D326-83D9-4A42-B9A5-0B27F9B40539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteBusinessCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBusinessCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBusinessCategoryResponseBody) SetCode(v string) *DeleteBusinessCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteBusinessCategoryResponseBody) SetData(v string) *DeleteBusinessCategoryResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteBusinessCategoryResponseBody) SetMessage(v string) *DeleteBusinessCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteBusinessCategoryResponseBody) SetRequestId(v string) *DeleteBusinessCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBusinessCategoryResponseBody) SetSuccess(v bool) *DeleteBusinessCategoryResponseBody {
	s.Success = &v
	return s
}

type DeleteBusinessCategoryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBusinessCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBusinessCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBusinessCategoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteBusinessCategoryResponse) SetHeaders(v map[string]*string) *DeleteBusinessCategoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteBusinessCategoryResponse) SetStatusCode(v int32) *DeleteBusinessCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBusinessCategoryResponse) SetBody(v *DeleteBusinessCategoryResponseBody) *DeleteBusinessCategoryResponse {
	s.Body = v
	return s
}

type DeleteCustomizationConfigRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"modelId":"2412"}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeleteCustomizationConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomizationConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomizationConfigRequest) SetBaseMeAgentId(v int64) *DeleteCustomizationConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *DeleteCustomizationConfigRequest) SetJsonStr(v string) *DeleteCustomizationConfigRequest {
	s.JsonStr = &v
	return s
}

type DeleteCustomizationConfigResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 252
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9987D326-83D9-4A42-B9A5-0B27F9B40539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCustomizationConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomizationConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomizationConfigResponseBody) SetCode(v string) *DeleteCustomizationConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCustomizationConfigResponseBody) SetData(v string) *DeleteCustomizationConfigResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteCustomizationConfigResponseBody) SetMessage(v string) *DeleteCustomizationConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCustomizationConfigResponseBody) SetRequestId(v string) *DeleteCustomizationConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomizationConfigResponseBody) SetSuccess(v bool) *DeleteCustomizationConfigResponseBody {
	s.Success = &v
	return s
}

type DeleteCustomizationConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomizationConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomizationConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomizationConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomizationConfigResponse) SetHeaders(v map[string]*string) *DeleteCustomizationConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomizationConfigResponse) SetStatusCode(v int32) *DeleteCustomizationConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomizationConfigResponse) SetBody(v *DeleteCustomizationConfigResponseBody) *DeleteCustomizationConfigResponse {
	s.Body = v
	return s
}

type DeleteDataSetRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"setId":"234"}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeleteDataSetRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSetRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataSetRequest) SetBaseMeAgentId(v int64) *DeleteDataSetRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *DeleteDataSetRequest) SetJsonStr(v string) *DeleteDataSetRequest {
	s.JsonStr = &v
	return s
}

type DeleteDataSetResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9987D326-83D9-4A42-B9A5-0B27F9B40539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDataSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataSetResponseBody) SetCode(v string) *DeleteDataSetResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDataSetResponseBody) SetMessage(v string) *DeleteDataSetResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDataSetResponseBody) SetRequestId(v string) *DeleteDataSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataSetResponseBody) SetSuccess(v bool) *DeleteDataSetResponseBody {
	s.Success = &v
	return s
}

type DeleteDataSetResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataSetResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSetResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataSetResponse) SetHeaders(v map[string]*string) *DeleteDataSetResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataSetResponse) SetStatusCode(v int32) *DeleteDataSetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataSetResponse) SetBody(v *DeleteDataSetResponseBody) *DeleteDataSetResponse {
	s.Body = v
	return s
}

type DeletePrecisionTaskRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// "{"taskId": "7C1DEF5F-2C18-4D36-99C6*******"}"
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeletePrecisionTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePrecisionTaskRequest) GoString() string {
	return s.String()
}

func (s *DeletePrecisionTaskRequest) SetBaseMeAgentId(v int64) *DeletePrecisionTaskRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *DeletePrecisionTaskRequest) SetJsonStr(v string) *DeletePrecisionTaskRequest {
	s.JsonStr = &v
	return s
}

type DeletePrecisionTaskResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24F4CE647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeletePrecisionTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePrecisionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePrecisionTaskResponseBody) SetCode(v string) *DeletePrecisionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePrecisionTaskResponseBody) SetMessage(v string) *DeletePrecisionTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePrecisionTaskResponseBody) SetRequestId(v string) *DeletePrecisionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePrecisionTaskResponseBody) SetSuccess(v bool) *DeletePrecisionTaskResponseBody {
	s.Success = &v
	return s
}

type DeletePrecisionTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePrecisionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePrecisionTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePrecisionTaskResponse) GoString() string {
	return s.String()
}

func (s *DeletePrecisionTaskResponse) SetHeaders(v map[string]*string) *DeletePrecisionTaskResponse {
	s.Headers = v
	return s
}

func (s *DeletePrecisionTaskResponse) SetStatusCode(v int32) *DeletePrecisionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePrecisionTaskResponse) SetBody(v *DeletePrecisionTaskResponseBody) *DeletePrecisionTaskResponse {
	s.Body = v
	return s
}

type DeleteQualityCheckSchemeRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// {"schemeId":191}
	JsonStr *string `json:"jsonStr,omitempty" xml:"jsonStr,omitempty"`
}

func (s DeleteQualityCheckSchemeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteQualityCheckSchemeRequest) GoString() string {
	return s.String()
}

func (s *DeleteQualityCheckSchemeRequest) SetBaseMeAgentId(v int64) *DeleteQualityCheckSchemeRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *DeleteQualityCheckSchemeRequest) SetJsonStr(v string) *DeleteQualityCheckSchemeRequest {
	s.JsonStr = &v
	return s
}

type DeleteQualityCheckSchemeResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *DeleteQualityCheckSchemeResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteQualityCheckSchemeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteQualityCheckSchemeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQualityCheckSchemeResponseBody) SetCode(v string) *DeleteQualityCheckSchemeResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteQualityCheckSchemeResponseBody) SetHttpStatusCode(v int32) *DeleteQualityCheckSchemeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteQualityCheckSchemeResponseBody) SetMessage(v string) *DeleteQualityCheckSchemeResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteQualityCheckSchemeResponseBody) SetMessages(v *DeleteQualityCheckSchemeResponseBodyMessages) *DeleteQualityCheckSchemeResponseBody {
	s.Messages = v
	return s
}

func (s *DeleteQualityCheckSchemeResponseBody) SetRequestId(v string) *DeleteQualityCheckSchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteQualityCheckSchemeResponseBody) SetSuccess(v bool) *DeleteQualityCheckSchemeResponseBody {
	s.Success = &v
	return s
}

type DeleteQualityCheckSchemeResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s DeleteQualityCheckSchemeResponseBodyMessages) String() string {
	return tea.Prettify(s)
}

func (s DeleteQualityCheckSchemeResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *DeleteQualityCheckSchemeResponseBodyMessages) SetMessage(v []*string) *DeleteQualityCheckSchemeResponseBodyMessages {
	s.Message = v
	return s
}

type DeleteQualityCheckSchemeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteQualityCheckSchemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteQualityCheckSchemeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteQualityCheckSchemeResponse) GoString() string {
	return s.String()
}

func (s *DeleteQualityCheckSchemeResponse) SetHeaders(v map[string]*string) *DeleteQualityCheckSchemeResponse {
	s.Headers = v
	return s
}

func (s *DeleteQualityCheckSchemeResponse) SetStatusCode(v int32) *DeleteQualityCheckSchemeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteQualityCheckSchemeResponse) SetBody(v *DeleteQualityCheckSchemeResponseBody) *DeleteQualityCheckSchemeResponse {
	s.Body = v
	return s
}

type DeleteRuleRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// true
	ForceDelete *bool `json:"ForceDelete,omitempty" xml:"ForceDelete,omitempty"`
	// example:
	//
	// 1
	IsSchemeData *int32 `json:"IsSchemeData,omitempty" xml:"IsSchemeData,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRuleRequest) SetBaseMeAgentId(v int64) *DeleteRuleRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *DeleteRuleRequest) SetForceDelete(v bool) *DeleteRuleRequest {
	s.ForceDelete = &v
	return s
}

func (s *DeleteRuleRequest) SetIsSchemeData(v int32) *DeleteRuleRequest {
	s.IsSchemeData = &v
	return s
}

func (s *DeleteRuleRequest) SetRuleId(v int64) *DeleteRuleRequest {
	s.RuleId = &v
	return s
}

type DeleteRuleResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *DeleteRuleResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRuleResponseBody) SetCode(v string) *DeleteRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRuleResponseBody) SetHttpStatusCode(v int32) *DeleteRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteRuleResponseBody) SetMessage(v string) *DeleteRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteRuleResponseBody) SetMessages(v *DeleteRuleResponseBodyMessages) *DeleteRuleResponseBody {
	s.Messages = v
	return s
}

func (s *DeleteRuleResponseBody) SetRequestId(v string) *DeleteRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRuleResponseBody) SetSuccess(v bool) *DeleteRuleResponseBody {
	s.Success = &v
	return s
}

type DeleteRuleResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s DeleteRuleResponseBodyMessages) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *DeleteRuleResponseBodyMessages) SetMessage(v []*string) *DeleteRuleResponseBodyMessages {
	s.Message = v
	return s
}

type DeleteRuleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteRuleResponse) SetHeaders(v map[string]*string) *DeleteRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteRuleResponse) SetStatusCode(v int32) *DeleteRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRuleResponse) SetBody(v *DeleteRuleResponseBody) *DeleteRuleResponse {
	s.Body = v
	return s
}

type DeleteRuleV4Request struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// false
	ForceDelete *bool `json:"ForceDelete,omitempty" xml:"ForceDelete,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteRuleV4Request) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleV4Request) GoString() string {
	return s.String()
}

func (s *DeleteRuleV4Request) SetBaseMeAgentId(v int64) *DeleteRuleV4Request {
	s.BaseMeAgentId = &v
	return s
}

func (s *DeleteRuleV4Request) SetForceDelete(v bool) *DeleteRuleV4Request {
	s.ForceDelete = &v
	return s
}

func (s *DeleteRuleV4Request) SetRuleId(v int64) *DeleteRuleV4Request {
	s.RuleId = &v
	return s
}

type DeleteRuleV4ResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *DeleteRuleV4ResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// F190ADE9-619A-447D-84E3-7E241A5C428E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteRuleV4ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleV4ResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRuleV4ResponseBody) SetCode(v string) *DeleteRuleV4ResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRuleV4ResponseBody) SetHttpStatusCode(v int32) *DeleteRuleV4ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteRuleV4ResponseBody) SetMessage(v string) *DeleteRuleV4ResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteRuleV4ResponseBody) SetMessages(v *DeleteRuleV4ResponseBodyMessages) *DeleteRuleV4ResponseBody {
	s.Messages = v
	return s
}

func (s *DeleteRuleV4ResponseBody) SetRequestId(v string) *DeleteRuleV4ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRuleV4ResponseBody) SetSuccess(v bool) *DeleteRuleV4ResponseBody {
	s.Success = &v
	return s
}

type DeleteRuleV4ResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s DeleteRuleV4ResponseBodyMessages) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleV4ResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *DeleteRuleV4ResponseBodyMessages) SetMessage(v []*string) *DeleteRuleV4ResponseBodyMessages {
	s.Message = v
	return s
}

type DeleteRuleV4Response struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRuleV4ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRuleV4Response) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleV4Response) GoString() string {
	return s.String()
}

func (s *DeleteRuleV4Response) SetHeaders(v map[string]*string) *DeleteRuleV4Response {
	s.Headers = v
	return s
}

func (s *DeleteRuleV4Response) SetStatusCode(v int32) *DeleteRuleV4Response {
	s.StatusCode = &v
	return s
}

func (s *DeleteRuleV4Response) SetBody(v *DeleteRuleV4ResponseBody) *DeleteRuleV4Response {
	s.Body = v
	return s
}

type DeleteSchemeTaskConfigRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// {"schemeId":"329"}
	JsonStr *string `json:"jsonStr,omitempty" xml:"jsonStr,omitempty"`
}

func (s DeleteSchemeTaskConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSchemeTaskConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteSchemeTaskConfigRequest) SetBaseMeAgentId(v int64) *DeleteSchemeTaskConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *DeleteSchemeTaskConfigRequest) SetJsonStr(v string) *DeleteSchemeTaskConfigRequest {
	s.JsonStr = &v
	return s
}

type DeleteSchemeTaskConfigResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *DeleteSchemeTaskConfigResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// 82C91484-B2D5-4D2A-A21F-A6D73F4D55C6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSchemeTaskConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSchemeTaskConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSchemeTaskConfigResponseBody) SetCode(v string) *DeleteSchemeTaskConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSchemeTaskConfigResponseBody) SetHttpStatusCode(v int32) *DeleteSchemeTaskConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteSchemeTaskConfigResponseBody) SetMessage(v string) *DeleteSchemeTaskConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSchemeTaskConfigResponseBody) SetMessages(v *DeleteSchemeTaskConfigResponseBodyMessages) *DeleteSchemeTaskConfigResponseBody {
	s.Messages = v
	return s
}

func (s *DeleteSchemeTaskConfigResponseBody) SetRequestId(v string) *DeleteSchemeTaskConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSchemeTaskConfigResponseBody) SetSuccess(v bool) *DeleteSchemeTaskConfigResponseBody {
	s.Success = &v
	return s
}

type DeleteSchemeTaskConfigResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s DeleteSchemeTaskConfigResponseBodyMessages) String() string {
	return tea.Prettify(s)
}

func (s DeleteSchemeTaskConfigResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *DeleteSchemeTaskConfigResponseBodyMessages) SetMessage(v []*string) *DeleteSchemeTaskConfigResponseBodyMessages {
	s.Message = v
	return s
}

type DeleteSchemeTaskConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSchemeTaskConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSchemeTaskConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSchemeTaskConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteSchemeTaskConfigResponse) SetHeaders(v map[string]*string) *DeleteSchemeTaskConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteSchemeTaskConfigResponse) SetStatusCode(v int32) *DeleteSchemeTaskConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSchemeTaskConfigResponse) SetBody(v *DeleteSchemeTaskConfigResponseBody) *DeleteSchemeTaskConfigResponse {
	s.Body = v
	return s
}

type DeleteScoreForApiRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeleteScoreForApiRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteScoreForApiRequest) GoString() string {
	return s.String()
}

func (s *DeleteScoreForApiRequest) SetBaseMeAgentId(v int64) *DeleteScoreForApiRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *DeleteScoreForApiRequest) SetJsonStr(v string) *DeleteScoreForApiRequest {
	s.JsonStr = &v
	return s
}

type DeleteScoreForApiResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9987D326-83D9-4A42-B9A5-0B27F9B40539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteScoreForApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteScoreForApiResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScoreForApiResponseBody) SetCode(v string) *DeleteScoreForApiResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteScoreForApiResponseBody) SetMessage(v string) *DeleteScoreForApiResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteScoreForApiResponseBody) SetRequestId(v string) *DeleteScoreForApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteScoreForApiResponseBody) SetSuccess(v bool) *DeleteScoreForApiResponseBody {
	s.Success = &v
	return s
}

type DeleteScoreForApiResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteScoreForApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteScoreForApiResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteScoreForApiResponse) GoString() string {
	return s.String()
}

func (s *DeleteScoreForApiResponse) SetHeaders(v map[string]*string) *DeleteScoreForApiResponse {
	s.Headers = v
	return s
}

func (s *DeleteScoreForApiResponse) SetStatusCode(v int32) *DeleteScoreForApiResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScoreForApiResponse) SetBody(v *DeleteScoreForApiResponseBody) *DeleteScoreForApiResponse {
	s.Body = v
	return s
}

type DeleteSkillGroupConfigRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"id":552}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeleteSkillGroupConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSkillGroupConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteSkillGroupConfigRequest) SetBaseMeAgentId(v int64) *DeleteSkillGroupConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *DeleteSkillGroupConfigRequest) SetJsonStr(v string) *DeleteSkillGroupConfigRequest {
	s.JsonStr = &v
	return s
}

type DeleteSkillGroupConfigResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3CEA0495-341B-4482-9AD9-8191EF4***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSkillGroupConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSkillGroupConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSkillGroupConfigResponseBody) SetCode(v string) *DeleteSkillGroupConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSkillGroupConfigResponseBody) SetMessage(v string) *DeleteSkillGroupConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSkillGroupConfigResponseBody) SetRequestId(v string) *DeleteSkillGroupConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSkillGroupConfigResponseBody) SetSuccess(v bool) *DeleteSkillGroupConfigResponseBody {
	s.Success = &v
	return s
}

type DeleteSkillGroupConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSkillGroupConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSkillGroupConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSkillGroupConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteSkillGroupConfigResponse) SetHeaders(v map[string]*string) *DeleteSkillGroupConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteSkillGroupConfigResponse) SetStatusCode(v int32) *DeleteSkillGroupConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSkillGroupConfigResponse) SetBody(v *DeleteSkillGroupConfigResponseBody) *DeleteSkillGroupConfigResponse {
	s.Body = v
	return s
}

type DeleteSubScoreForApiRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeleteSubScoreForApiRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubScoreForApiRequest) GoString() string {
	return s.String()
}

func (s *DeleteSubScoreForApiRequest) SetBaseMeAgentId(v int64) *DeleteSubScoreForApiRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *DeleteSubScoreForApiRequest) SetJsonStr(v string) *DeleteSubScoreForApiRequest {
	s.JsonStr = &v
	return s
}

type DeleteSubScoreForApiResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9987D326-83D9-4A42-B9A5-0B27F9B43539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSubScoreForApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubScoreForApiResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSubScoreForApiResponseBody) SetCode(v string) *DeleteSubScoreForApiResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSubScoreForApiResponseBody) SetMessage(v string) *DeleteSubScoreForApiResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSubScoreForApiResponseBody) SetRequestId(v string) *DeleteSubScoreForApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSubScoreForApiResponseBody) SetSuccess(v bool) *DeleteSubScoreForApiResponseBody {
	s.Success = &v
	return s
}

type DeleteSubScoreForApiResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSubScoreForApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSubScoreForApiResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubScoreForApiResponse) GoString() string {
	return s.String()
}

func (s *DeleteSubScoreForApiResponse) SetHeaders(v map[string]*string) *DeleteSubScoreForApiResponse {
	s.Headers = v
	return s
}

func (s *DeleteSubScoreForApiResponse) SetStatusCode(v int32) *DeleteSubScoreForApiResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSubScoreForApiResponse) SetBody(v *DeleteSubScoreForApiResponseBody) *DeleteSubScoreForApiResponse {
	s.Body = v
	return s
}

type DeleteTaskAssignRuleRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"ruleId": 24}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeleteTaskAssignRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTaskAssignRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteTaskAssignRuleRequest) SetBaseMeAgentId(v int64) *DeleteTaskAssignRuleRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *DeleteTaskAssignRuleRequest) SetJsonStr(v string) *DeleteTaskAssignRuleRequest {
	s.JsonStr = &v
	return s
}

type DeleteTaskAssignRuleResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24F4CE647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteTaskAssignRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTaskAssignRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTaskAssignRuleResponseBody) SetCode(v string) *DeleteTaskAssignRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTaskAssignRuleResponseBody) SetMessage(v string) *DeleteTaskAssignRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTaskAssignRuleResponseBody) SetRequestId(v string) *DeleteTaskAssignRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTaskAssignRuleResponseBody) SetSuccess(v bool) *DeleteTaskAssignRuleResponseBody {
	s.Success = &v
	return s
}

type DeleteTaskAssignRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTaskAssignRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTaskAssignRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTaskAssignRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteTaskAssignRuleResponse) SetHeaders(v map[string]*string) *DeleteTaskAssignRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteTaskAssignRuleResponse) SetStatusCode(v int32) *DeleteTaskAssignRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTaskAssignRuleResponse) SetBody(v *DeleteTaskAssignRuleResponseBody) *DeleteTaskAssignRuleResponse {
	s.Body = v
	return s
}

type DeleteWarningConfigRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"configId": "31"}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeleteWarningConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWarningConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteWarningConfigRequest) SetBaseMeAgentId(v int64) *DeleteWarningConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *DeleteWarningConfigRequest) SetJsonStr(v string) *DeleteWarningConfigRequest {
	s.JsonStr = &v
	return s
}

type DeleteWarningConfigResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 82C91484-B2D5-4D2A-A21F-A6D73F4D55C6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteWarningConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWarningConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWarningConfigResponseBody) SetCode(v string) *DeleteWarningConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteWarningConfigResponseBody) SetMessage(v string) *DeleteWarningConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteWarningConfigResponseBody) SetRequestId(v string) *DeleteWarningConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWarningConfigResponseBody) SetSuccess(v bool) *DeleteWarningConfigResponseBody {
	s.Success = &v
	return s
}

type DeleteWarningConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWarningConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWarningConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWarningConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteWarningConfigResponse) SetHeaders(v map[string]*string) *DeleteWarningConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteWarningConfigResponse) SetStatusCode(v int32) *DeleteWarningConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWarningConfigResponse) SetBody(v *DeleteWarningConfigResponseBody) *DeleteWarningConfigResponse {
	s.Body = v
	return s
}

type DeleteWarningStrategyConfigRequest struct {
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeleteWarningStrategyConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWarningStrategyConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteWarningStrategyConfigRequest) SetBaseMeAgentId(v int64) *DeleteWarningStrategyConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *DeleteWarningStrategyConfigRequest) SetJsonStr(v string) *DeleteWarningStrategyConfigRequest {
	s.JsonStr = &v
	return s
}

type DeleteWarningStrategyConfigResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteWarningStrategyConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWarningStrategyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWarningStrategyConfigResponseBody) SetCode(v string) *DeleteWarningStrategyConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteWarningStrategyConfigResponseBody) SetData(v string) *DeleteWarningStrategyConfigResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteWarningStrategyConfigResponseBody) SetMessage(v string) *DeleteWarningStrategyConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteWarningStrategyConfigResponseBody) SetRequestId(v string) *DeleteWarningStrategyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWarningStrategyConfigResponseBody) SetSuccess(v bool) *DeleteWarningStrategyConfigResponseBody {
	s.Success = &v
	return s
}

type DeleteWarningStrategyConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWarningStrategyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWarningStrategyConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWarningStrategyConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteWarningStrategyConfigResponse) SetHeaders(v map[string]*string) *DeleteWarningStrategyConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteWarningStrategyConfigResponse) SetStatusCode(v int32) *DeleteWarningStrategyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWarningStrategyConfigResponse) SetBody(v *DeleteWarningStrategyConfigResponseBody) *DeleteWarningStrategyConfigResponse {
	s.Body = v
	return s
}

type GetAsrVocabRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64  `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	JsonStr       *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetAsrVocabRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAsrVocabRequest) GoString() string {
	return s.String()
}

func (s *GetAsrVocabRequest) SetBaseMeAgentId(v int64) *GetAsrVocabRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetAsrVocabRequest) SetJsonStr(v string) *GetAsrVocabRequest {
	s.JsonStr = &v
	return s
}

type GetAsrVocabResponseBody struct {
	// example:
	//
	// 200
	Code *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetAsrVocabResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3CEA0495-341B-4482-9AD9-8191EF4***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAsrVocabResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAsrVocabResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsrVocabResponseBody) SetCode(v string) *GetAsrVocabResponseBody {
	s.Code = &v
	return s
}

func (s *GetAsrVocabResponseBody) SetData(v *GetAsrVocabResponseBodyData) *GetAsrVocabResponseBody {
	s.Data = v
	return s
}

func (s *GetAsrVocabResponseBody) SetMessage(v string) *GetAsrVocabResponseBody {
	s.Message = &v
	return s
}

func (s *GetAsrVocabResponseBody) SetRequestId(v string) *GetAsrVocabResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAsrVocabResponseBody) SetSuccess(v bool) *GetAsrVocabResponseBody {
	s.Success = &v
	return s
}

type GetAsrVocabResponseBodyData struct {
	// example:
	//
	// test
	Name  *string                           `json:"Name,omitempty" xml:"Name,omitempty"`
	Words *GetAsrVocabResponseBodyDataWords `json:"Words,omitempty" xml:"Words,omitempty" type:"Struct"`
}

func (s GetAsrVocabResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAsrVocabResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAsrVocabResponseBodyData) SetName(v string) *GetAsrVocabResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetAsrVocabResponseBodyData) SetWords(v *GetAsrVocabResponseBodyDataWords) *GetAsrVocabResponseBodyData {
	s.Words = v
	return s
}

type GetAsrVocabResponseBodyDataWords struct {
	Word []*GetAsrVocabResponseBodyDataWordsWord `json:"Word,omitempty" xml:"Word,omitempty" type:"Repeated"`
}

func (s GetAsrVocabResponseBodyDataWords) String() string {
	return tea.Prettify(s)
}

func (s GetAsrVocabResponseBodyDataWords) GoString() string {
	return s.String()
}

func (s *GetAsrVocabResponseBodyDataWords) SetWord(v []*GetAsrVocabResponseBodyDataWordsWord) *GetAsrVocabResponseBodyDataWords {
	s.Word = v
	return s
}

type GetAsrVocabResponseBodyDataWordsWord struct {
	// example:
	//
	// 0
	Weight *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
	Word   *string `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s GetAsrVocabResponseBodyDataWordsWord) String() string {
	return tea.Prettify(s)
}

func (s GetAsrVocabResponseBodyDataWordsWord) GoString() string {
	return s.String()
}

func (s *GetAsrVocabResponseBodyDataWordsWord) SetWeight(v int32) *GetAsrVocabResponseBodyDataWordsWord {
	s.Weight = &v
	return s
}

func (s *GetAsrVocabResponseBodyDataWordsWord) SetWord(v string) *GetAsrVocabResponseBodyDataWordsWord {
	s.Word = &v
	return s
}

type GetAsrVocabResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAsrVocabResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAsrVocabResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAsrVocabResponse) GoString() string {
	return s.String()
}

func (s *GetAsrVocabResponse) SetHeaders(v map[string]*string) *GetAsrVocabResponse {
	s.Headers = v
	return s
}

func (s *GetAsrVocabResponse) SetStatusCode(v int32) *GetAsrVocabResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAsrVocabResponse) SetBody(v *GetAsrVocabResponseBody) *GetAsrVocabResponse {
	s.Body = v
	return s
}

type GetBusinessCategoryListRequest struct {
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ""
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetBusinessCategoryListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBusinessCategoryListRequest) GoString() string {
	return s.String()
}

func (s *GetBusinessCategoryListRequest) SetBaseMeAgentId(v int64) *GetBusinessCategoryListRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetBusinessCategoryListRequest) SetJsonStr(v string) *GetBusinessCategoryListRequest {
	s.JsonStr = &v
	return s
}

type GetBusinessCategoryListResponseBody struct {
	// example:
	//
	// 200
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetBusinessCategoryListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A186A419-FDBE-464C-AED4-7121CAC73BF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetBusinessCategoryListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBusinessCategoryListResponseBody) GoString() string {
	return s.String()
}

func (s *GetBusinessCategoryListResponseBody) SetCode(v string) *GetBusinessCategoryListResponseBody {
	s.Code = &v
	return s
}

func (s *GetBusinessCategoryListResponseBody) SetData(v *GetBusinessCategoryListResponseBodyData) *GetBusinessCategoryListResponseBody {
	s.Data = v
	return s
}

func (s *GetBusinessCategoryListResponseBody) SetMessage(v string) *GetBusinessCategoryListResponseBody {
	s.Message = &v
	return s
}

func (s *GetBusinessCategoryListResponseBody) SetRequestId(v string) *GetBusinessCategoryListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBusinessCategoryListResponseBody) SetSuccess(v bool) *GetBusinessCategoryListResponseBody {
	s.Success = &v
	return s
}

type GetBusinessCategoryListResponseBodyData struct {
	BusinessCategoryBasicInfo []*GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo `json:"BusinessCategoryBasicInfo,omitempty" xml:"BusinessCategoryBasicInfo,omitempty" type:"Repeated"`
}

func (s GetBusinessCategoryListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetBusinessCategoryListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBusinessCategoryListResponseBodyData) SetBusinessCategoryBasicInfo(v []*GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo) *GetBusinessCategoryListResponseBodyData {
	s.BusinessCategoryBasicInfo = v
	return s
}

type GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo struct {
	// example:
	//
	// 0
	Bid          *int32  `json:"Bid,omitempty" xml:"Bid,omitempty"`
	BusinessName *string `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	// example:
	//
	// 0
	ServiceType *int32 `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo) GoString() string {
	return s.String()
}

func (s *GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo) SetBid(v int32) *GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo {
	s.Bid = &v
	return s
}

func (s *GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo) SetBusinessName(v string) *GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo {
	s.BusinessName = &v
	return s
}

func (s *GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo) SetServiceType(v int32) *GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo {
	s.ServiceType = &v
	return s
}

type GetBusinessCategoryListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBusinessCategoryListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBusinessCategoryListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBusinessCategoryListResponse) GoString() string {
	return s.String()
}

func (s *GetBusinessCategoryListResponse) SetHeaders(v map[string]*string) *GetBusinessCategoryListResponse {
	s.Headers = v
	return s
}

func (s *GetBusinessCategoryListResponse) SetStatusCode(v int32) *GetBusinessCategoryListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBusinessCategoryListResponse) SetBody(v *GetBusinessCategoryListResponseBody) *GetBusinessCategoryListResponse {
	s.Body = v
	return s
}

type GetCustomizationConfigListRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ""
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetCustomizationConfigListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCustomizationConfigListRequest) GoString() string {
	return s.String()
}

func (s *GetCustomizationConfigListRequest) SetBaseMeAgentId(v int64) *GetCustomizationConfigListRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetCustomizationConfigListRequest) SetJsonStr(v string) *GetCustomizationConfigListRequest {
	s.JsonStr = &v
	return s
}

type GetCustomizationConfigListResponseBody struct {
	// example:
	//
	// 200
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetCustomizationConfigListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24F4CE647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCustomizationConfigListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCustomizationConfigListResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomizationConfigListResponseBody) SetCode(v string) *GetCustomizationConfigListResponseBody {
	s.Code = &v
	return s
}

func (s *GetCustomizationConfigListResponseBody) SetData(v *GetCustomizationConfigListResponseBodyData) *GetCustomizationConfigListResponseBody {
	s.Data = v
	return s
}

func (s *GetCustomizationConfigListResponseBody) SetMessage(v string) *GetCustomizationConfigListResponseBody {
	s.Message = &v
	return s
}

func (s *GetCustomizationConfigListResponseBody) SetRequestId(v string) *GetCustomizationConfigListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomizationConfigListResponseBody) SetSuccess(v bool) *GetCustomizationConfigListResponseBody {
	s.Success = &v
	return s
}

type GetCustomizationConfigListResponseBodyData struct {
	ModelCustomizationDataSetPo []*GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo `json:"ModelCustomizationDataSetPo,omitempty" xml:"ModelCustomizationDataSetPo,omitempty" type:"Repeated"`
}

func (s GetCustomizationConfigListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCustomizationConfigListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCustomizationConfigListResponseBodyData) SetModelCustomizationDataSetPo(v []*GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) *GetCustomizationConfigListResponseBodyData {
	s.ModelCustomizationDataSetPo = v
	return s
}

type GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo struct {
	// example:
	//
	// 2019-01-08
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// cdae396590bb479a9ec40f3476e274fc
	ModeCustomizationId *string `json:"ModeCustomizationId,omitempty" xml:"ModeCustomizationId,omitempty"`
	// example:
	//
	// 1
	ModelId   *int64  `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// example:
	//
	// 5
	ModelStatus *int32 `json:"ModelStatus,omitempty" xml:"ModelStatus,omitempty"`
	// example:
	//
	// 1
	TaskType *int32 `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) String() string {
	return tea.Prettify(s)
}

func (s GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) GoString() string {
	return s.String()
}

func (s *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) SetCreateTime(v string) *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo {
	s.CreateTime = &v
	return s
}

func (s *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) SetModeCustomizationId(v string) *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo {
	s.ModeCustomizationId = &v
	return s
}

func (s *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) SetModelId(v int64) *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo {
	s.ModelId = &v
	return s
}

func (s *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) SetModelName(v string) *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo {
	s.ModelName = &v
	return s
}

func (s *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) SetModelStatus(v int32) *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo {
	s.ModelStatus = &v
	return s
}

func (s *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) SetTaskType(v int32) *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo {
	s.TaskType = &v
	return s
}

type GetCustomizationConfigListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCustomizationConfigListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomizationConfigListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCustomizationConfigListResponse) GoString() string {
	return s.String()
}

func (s *GetCustomizationConfigListResponse) SetHeaders(v map[string]*string) *GetCustomizationConfigListResponse {
	s.Headers = v
	return s
}

func (s *GetCustomizationConfigListResponse) SetStatusCode(v int32) *GetCustomizationConfigListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomizationConfigListResponse) SetBody(v *GetCustomizationConfigListResponseBody) *GetCustomizationConfigListResponse {
	s.Body = v
	return s
}

type GetNextResultToVerifyRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// "{"pageNumber":1,"pageSize":1,"taskId":"593A04C0-E6E9-4CDC-8C9*****","original":1}"
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetNextResultToVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNextResultToVerifyRequest) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyRequest) SetBaseMeAgentId(v int64) *GetNextResultToVerifyRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetNextResultToVerifyRequest) SetJsonStr(v string) *GetNextResultToVerifyRequest {
	s.JsonStr = &v
	return s
}

type GetNextResultToVerifyResponseBody struct {
	// example:
	//
	// 200
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetNextResultToVerifyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24F4CE647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNextResultToVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNextResultToVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyResponseBody) SetCode(v string) *GetNextResultToVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *GetNextResultToVerifyResponseBody) SetData(v *GetNextResultToVerifyResponseBodyData) *GetNextResultToVerifyResponseBody {
	s.Data = v
	return s
}

func (s *GetNextResultToVerifyResponseBody) SetMessage(v string) *GetNextResultToVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *GetNextResultToVerifyResponseBody) SetRequestId(v string) *GetNextResultToVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNextResultToVerifyResponseBody) SetSuccess(v bool) *GetNextResultToVerifyResponseBody {
	s.Success = &v
	return s
}

type GetNextResultToVerifyResponseBodyData struct {
	// example:
	//
	// http
	AudioScheme *string `json:"AudioScheme,omitempty" xml:"AudioScheme,omitempty"`
	// example:
	//
	// sca-bucket.oss-cn-hangzhou.aliyuncs.com/upload_1173636551461420/dateset_1584674455133_SzC/%E4%BA%BA%E5%B7%A5%E6%A0%A1%E9%AA%8C%E6%B5%8B%E8%AF%95-%E6%9F%A5%E5%8C%97%E4%BA%AC%E5%A4%A9%E6%B0%94.wav?Expires=1584847372&amp;OSSAccessKeyId=*****&amp;Signature=HccAKnLOJwoYvzE*********
	AudioURL  *string                                         `json:"AudioURL,omitempty" xml:"AudioURL,omitempty"`
	Dialogues *GetNextResultToVerifyResponseBodyDataDialogues `json:"Dialogues,omitempty" xml:"Dialogues,omitempty" type:"Struct"`
	// example:
	//
	// 23421
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// e790e6c919d84b82b64ee*****
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// xxx.wav
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 23
	IncorrectWords *int32 `json:"IncorrectWords,omitempty" xml:"IncorrectWords,omitempty"`
	// example:
	//
	// 2
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// 0.97079998
	Precision *float32 `json:"Precision,omitempty" xml:"Precision,omitempty"`
	// example:
	//
	// 3
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 2020-03-20T11:26Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// true
	Verified *bool `json:"Verified,omitempty" xml:"Verified,omitempty"`
	// example:
	//
	// 2
	VerifiedCount *int32 `json:"VerifiedCount,omitempty" xml:"VerifiedCount,omitempty"`
}

func (s GetNextResultToVerifyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetNextResultToVerifyResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyResponseBodyData) SetAudioScheme(v string) *GetNextResultToVerifyResponseBodyData {
	s.AudioScheme = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetAudioURL(v string) *GetNextResultToVerifyResponseBodyData {
	s.AudioURL = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetDialogues(v *GetNextResultToVerifyResponseBodyDataDialogues) *GetNextResultToVerifyResponseBodyData {
	s.Dialogues = v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetDuration(v int32) *GetNextResultToVerifyResponseBodyData {
	s.Duration = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetFileId(v string) *GetNextResultToVerifyResponseBodyData {
	s.FileId = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetFileName(v string) *GetNextResultToVerifyResponseBodyData {
	s.FileName = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetIncorrectWords(v int32) *GetNextResultToVerifyResponseBodyData {
	s.IncorrectWords = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetIndex(v int32) *GetNextResultToVerifyResponseBodyData {
	s.Index = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetPrecision(v float32) *GetNextResultToVerifyResponseBodyData {
	s.Precision = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetStatus(v int32) *GetNextResultToVerifyResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetTotalCount(v int32) *GetNextResultToVerifyResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetUpdateTime(v string) *GetNextResultToVerifyResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetVerified(v bool) *GetNextResultToVerifyResponseBodyData {
	s.Verified = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetVerifiedCount(v int32) *GetNextResultToVerifyResponseBodyData {
	s.VerifiedCount = &v
	return s
}

type GetNextResultToVerifyResponseBodyDataDialogues struct {
	Dialogue []*GetNextResultToVerifyResponseBodyDataDialoguesDialogue `json:"Dialogue,omitempty" xml:"Dialogue,omitempty" type:"Repeated"`
}

func (s GetNextResultToVerifyResponseBodyDataDialogues) String() string {
	return tea.Prettify(s)
}

func (s GetNextResultToVerifyResponseBodyDataDialogues) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyResponseBodyDataDialogues) SetDialogue(v []*GetNextResultToVerifyResponseBodyDataDialoguesDialogue) *GetNextResultToVerifyResponseBodyDataDialogues {
	s.Dialogue = v
	return s
}

type GetNextResultToVerifyResponseBodyDataDialoguesDialogue struct {
	// example:
	//
	// 980
	Begin *int64 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// example:
	//
	// XXX
	BeginTime *string                                                       `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	Deltas    *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltas `json:"Deltas,omitempty" xml:"Deltas,omitempty" type:"Struct"`
	// example:
	//
	// 6
	EmotionValue *int32 `json:"EmotionValue,omitempty" xml:"EmotionValue,omitempty"`
	// example:
	//
	// 3422
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// example:
	//
	// 00:00:07
	HourMinSec *string `json:"HourMinSec,omitempty" xml:"HourMinSec,omitempty"`
	Identity   *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
	// example:
	//
	// 2
	IncorrectWords *int32  `json:"IncorrectWords,omitempty" xml:"IncorrectWords,omitempty"`
	Role           *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// 1
	SilenceDuration *int32  `json:"SilenceDuration,omitempty" xml:"SilenceDuration,omitempty"`
	SourceRole      *string `json:"SourceRole,omitempty" xml:"SourceRole,omitempty"`
	SourceWords     *string `json:"SourceWords,omitempty" xml:"SourceWords,omitempty"`
	// example:
	//
	// 332
	SpeechRate *int32  `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	Words      *string `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogue) String() string {
	return tea.Prettify(s)
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogue) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetBegin(v int64) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.Begin = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetBeginTime(v string) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.BeginTime = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetDeltas(v *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltas) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.Deltas = v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetEmotionValue(v int32) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.EmotionValue = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetEnd(v int64) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.End = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetHourMinSec(v string) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.HourMinSec = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetIdentity(v string) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.Identity = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetIncorrectWords(v int32) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.IncorrectWords = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetRole(v string) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.Role = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetSilenceDuration(v int32) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.SilenceDuration = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetSourceRole(v string) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.SourceRole = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetSourceWords(v string) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.SourceWords = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetSpeechRate(v int32) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.SpeechRate = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetWords(v string) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.Words = &v
	return s
}

type GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltas struct {
	Delta []*GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta `json:"Delta,omitempty" xml:"Delta,omitempty" type:"Repeated"`
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltas) String() string {
	return tea.Prettify(s)
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltas) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltas) SetDelta(v []*GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta) *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltas {
	s.Delta = v
	return s
}

type GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta struct {
	Source *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	Target *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// example:
	//
	// CHANGE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta) String() string {
	return tea.Prettify(s)
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta) SetSource(v *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSource) *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta {
	s.Source = v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta) SetTarget(v *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTarget) *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta {
	s.Target = v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta) SetType(v string) *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta {
	s.Type = &v
	return s
}

type GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSource struct {
	Line *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSourceLine `json:"Line,omitempty" xml:"Line,omitempty" type:"Struct"`
	// example:
	//
	// 5
	Position *int32 `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSource) String() string {
	return tea.Prettify(s)
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSource) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSource) SetLine(v *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSourceLine) *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSource {
	s.Line = v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSource) SetPosition(v int32) *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSource {
	s.Position = &v
	return s
}

type GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSourceLine struct {
	Line []*string `json:"Line,omitempty" xml:"Line,omitempty" type:"Repeated"`
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSourceLine) String() string {
	return tea.Prettify(s)
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSourceLine) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSourceLine) SetLine(v []*string) *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSourceLine {
	s.Line = v
	return s
}

type GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTarget struct {
	Line *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTargetLine `json:"Line,omitempty" xml:"Line,omitempty" type:"Struct"`
	// example:
	//
	// 5
	Position *int32 `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTarget) String() string {
	return tea.Prettify(s)
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTarget) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTarget) SetLine(v *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTargetLine) *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTarget {
	s.Line = v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTarget) SetPosition(v int32) *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTarget {
	s.Position = &v
	return s
}

type GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTargetLine struct {
	Line []*string `json:"Line,omitempty" xml:"Line,omitempty" type:"Repeated"`
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTargetLine) String() string {
	return tea.Prettify(s)
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTargetLine) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTargetLine) SetLine(v []*string) *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTargetLine {
	s.Line = v
	return s
}

type GetNextResultToVerifyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNextResultToVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNextResultToVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNextResultToVerifyResponse) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyResponse) SetHeaders(v map[string]*string) *GetNextResultToVerifyResponse {
	s.Headers = v
	return s
}

func (s *GetNextResultToVerifyResponse) SetStatusCode(v int32) *GetNextResultToVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNextResultToVerifyResponse) SetBody(v *GetNextResultToVerifyResponseBody) *GetNextResultToVerifyResponse {
	s.Body = v
	return s
}

type GetPrecisionTaskRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// "{"taskId":"593A04C0-E6E9-4CDC-8C9****"}"
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetPrecisionTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPrecisionTaskRequest) GoString() string {
	return s.String()
}

func (s *GetPrecisionTaskRequest) SetBaseMeAgentId(v int64) *GetPrecisionTaskRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetPrecisionTaskRequest) SetJsonStr(v string) *GetPrecisionTaskRequest {
	s.JsonStr = &v
	return s
}

type GetPrecisionTaskResponseBody struct {
	// example:
	//
	// 200
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetPrecisionTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24F4CE647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPrecisionTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPrecisionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetPrecisionTaskResponseBody) SetCode(v string) *GetPrecisionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetPrecisionTaskResponseBody) SetData(v *GetPrecisionTaskResponseBodyData) *GetPrecisionTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetPrecisionTaskResponseBody) SetMessage(v string) *GetPrecisionTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetPrecisionTaskResponseBody) SetRequestId(v string) *GetPrecisionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPrecisionTaskResponseBody) SetSuccess(v bool) *GetPrecisionTaskResponseBody {
	s.Success = &v
	return s
}

type GetPrecisionTaskResponseBodyData struct {
	// example:
	//
	// 1212
	DataSetId   *int64  `json:"DataSetId,omitempty" xml:"DataSetId,omitempty"`
	DataSetName *string `json:"DataSetName,omitempty" xml:"DataSetName,omitempty"`
	// example:
	//
	// 3423
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 23
	IncorrectWords *int32                                      `json:"IncorrectWords,omitempty" xml:"IncorrectWords,omitempty"`
	Name           *string                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	Precisions     *GetPrecisionTaskResponseBodyDataPrecisions `json:"Precisions,omitempty" xml:"Precisions,omitempty" type:"Struct"`
	// example:
	//
	// 3
	Source *int32 `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 7C1DEF5F-2C18-4D36-99C6-8C27*****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 2020-03-10 20:26:29
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 2
	VerifiedCount *int32 `json:"VerifiedCount,omitempty" xml:"VerifiedCount,omitempty"`
}

func (s GetPrecisionTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPrecisionTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPrecisionTaskResponseBodyData) SetDataSetId(v int64) *GetPrecisionTaskResponseBodyData {
	s.DataSetId = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyData) SetDataSetName(v string) *GetPrecisionTaskResponseBodyData {
	s.DataSetName = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyData) SetDuration(v int32) *GetPrecisionTaskResponseBodyData {
	s.Duration = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyData) SetIncorrectWords(v int32) *GetPrecisionTaskResponseBodyData {
	s.IncorrectWords = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyData) SetName(v string) *GetPrecisionTaskResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyData) SetPrecisions(v *GetPrecisionTaskResponseBodyDataPrecisions) *GetPrecisionTaskResponseBodyData {
	s.Precisions = v
	return s
}

func (s *GetPrecisionTaskResponseBodyData) SetSource(v int32) *GetPrecisionTaskResponseBodyData {
	s.Source = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyData) SetStatus(v int32) *GetPrecisionTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyData) SetTaskId(v string) *GetPrecisionTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyData) SetTotalCount(v int32) *GetPrecisionTaskResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyData) SetUpdateTime(v string) *GetPrecisionTaskResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyData) SetVerifiedCount(v int32) *GetPrecisionTaskResponseBodyData {
	s.VerifiedCount = &v
	return s
}

type GetPrecisionTaskResponseBodyDataPrecisions struct {
	Precision []*GetPrecisionTaskResponseBodyDataPrecisionsPrecision `json:"Precision,omitempty" xml:"Precision,omitempty" type:"Repeated"`
}

func (s GetPrecisionTaskResponseBodyDataPrecisions) String() string {
	return tea.Prettify(s)
}

func (s GetPrecisionTaskResponseBodyDataPrecisions) GoString() string {
	return s.String()
}

func (s *GetPrecisionTaskResponseBodyDataPrecisions) SetPrecision(v []*GetPrecisionTaskResponseBodyDataPrecisionsPrecision) *GetPrecisionTaskResponseBodyDataPrecisions {
	s.Precision = v
	return s
}

type GetPrecisionTaskResponseBodyDataPrecisionsPrecision struct {
	// example:
	//
	// 2311
	ModelId   *int64  `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// example:
	//
	// 0.98
	Precision *float32 `json:"Precision,omitempty" xml:"Precision,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 593A04C0-E6E9-4CDC-8C99-B120C******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetPrecisionTaskResponseBodyDataPrecisionsPrecision) String() string {
	return tea.Prettify(s)
}

func (s GetPrecisionTaskResponseBodyDataPrecisionsPrecision) GoString() string {
	return s.String()
}

func (s *GetPrecisionTaskResponseBodyDataPrecisionsPrecision) SetModelId(v int64) *GetPrecisionTaskResponseBodyDataPrecisionsPrecision {
	s.ModelId = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyDataPrecisionsPrecision) SetModelName(v string) *GetPrecisionTaskResponseBodyDataPrecisionsPrecision {
	s.ModelName = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyDataPrecisionsPrecision) SetPrecision(v float32) *GetPrecisionTaskResponseBodyDataPrecisionsPrecision {
	s.Precision = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyDataPrecisionsPrecision) SetStatus(v int32) *GetPrecisionTaskResponseBodyDataPrecisionsPrecision {
	s.Status = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyDataPrecisionsPrecision) SetTaskId(v string) *GetPrecisionTaskResponseBodyDataPrecisionsPrecision {
	s.TaskId = &v
	return s
}

type GetPrecisionTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPrecisionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPrecisionTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPrecisionTaskResponse) GoString() string {
	return s.String()
}

func (s *GetPrecisionTaskResponse) SetHeaders(v map[string]*string) *GetPrecisionTaskResponse {
	s.Headers = v
	return s
}

func (s *GetPrecisionTaskResponse) SetStatusCode(v int32) *GetPrecisionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPrecisionTaskResponse) SetBody(v *GetPrecisionTaskResponseBody) *GetPrecisionTaskResponse {
	s.Body = v
	return s
}

type GetQualityCheckSchemeRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// {"schemeId":"187","ruleRequireInfos":["BusinessNameInfo","RuleCategory"]}
	JsonStr *string `json:"jsonStr,omitempty" xml:"jsonStr,omitempty"`
}

func (s GetQualityCheckSchemeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQualityCheckSchemeRequest) GoString() string {
	return s.String()
}

func (s *GetQualityCheckSchemeRequest) SetBaseMeAgentId(v int64) *GetQualityCheckSchemeRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetQualityCheckSchemeRequest) SetJsonStr(v string) *GetQualityCheckSchemeRequest {
	s.JsonStr = &v
	return s
}

type GetQualityCheckSchemeResponseBody struct {
	// example:
	//
	// 200
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetQualityCheckSchemeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages []*string `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// example:
	//
	// 96138D8D-8D26-4E41-BFF4-77AED1088BBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQualityCheckSchemeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQualityCheckSchemeResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualityCheckSchemeResponseBody) SetCode(v string) *GetQualityCheckSchemeResponseBody {
	s.Code = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBody) SetData(v *GetQualityCheckSchemeResponseBodyData) *GetQualityCheckSchemeResponseBody {
	s.Data = v
	return s
}

func (s *GetQualityCheckSchemeResponseBody) SetHttpStatusCode(v int32) *GetQualityCheckSchemeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBody) SetMessage(v string) *GetQualityCheckSchemeResponseBody {
	s.Message = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBody) SetMessages(v []*string) *GetQualityCheckSchemeResponseBody {
	s.Messages = v
	return s
}

func (s *GetQualityCheckSchemeResponseBody) SetRequestId(v string) *GetQualityCheckSchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBody) SetSuccess(v bool) *GetQualityCheckSchemeResponseBody {
	s.Success = &v
	return s
}

type GetQualityCheckSchemeResponseBodyData struct {
	// example:
	//
	// 1616113198000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// xxx
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	// example:
	//
	// 1
	DataType    *int32  `json:"DataType,omitempty" xml:"DataType,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// xxx
	Name                *string                                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	RuleIds             []*string                                                   `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Repeated"`
	RuleList            []*RulesInfo                                                `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
	SchemeCheckTypeList []*GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList `json:"SchemeCheckTypeList,omitempty" xml:"SchemeCheckTypeList,omitempty" type:"Repeated"`
	// example:
	//
	// 112**
	SchemeId *int64 `json:"SchemeId,omitempty" xml:"SchemeId,omitempty"`
	// example:
	//
	// 1
	SchemeTemplateId *int64 `json:"SchemeTemplateId,omitempty" xml:"SchemeTemplateId,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1
	TemplateType *int32 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 1616113198000
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// xxx
	UpdateUserName *string `json:"UpdateUserName,omitempty" xml:"UpdateUserName,omitempty"`
	// example:
	//
	// 1616113198000
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetQualityCheckSchemeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetQualityCheckSchemeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQualityCheckSchemeResponseBodyData) SetCreateTime(v string) *GetQualityCheckSchemeResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyData) SetCreateUserName(v string) *GetQualityCheckSchemeResponseBodyData {
	s.CreateUserName = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyData) SetDataType(v int32) *GetQualityCheckSchemeResponseBodyData {
	s.DataType = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyData) SetDescription(v string) *GetQualityCheckSchemeResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyData) SetName(v string) *GetQualityCheckSchemeResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyData) SetRuleIds(v []*string) *GetQualityCheckSchemeResponseBodyData {
	s.RuleIds = v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyData) SetRuleList(v []*RulesInfo) *GetQualityCheckSchemeResponseBodyData {
	s.RuleList = v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyData) SetSchemeCheckTypeList(v []*GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) *GetQualityCheckSchemeResponseBodyData {
	s.SchemeCheckTypeList = v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyData) SetSchemeId(v int64) *GetQualityCheckSchemeResponseBodyData {
	s.SchemeId = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyData) SetSchemeTemplateId(v int64) *GetQualityCheckSchemeResponseBodyData {
	s.SchemeTemplateId = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyData) SetStatus(v int32) *GetQualityCheckSchemeResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyData) SetTemplateType(v int32) *GetQualityCheckSchemeResponseBodyData {
	s.TemplateType = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyData) SetType(v int32) *GetQualityCheckSchemeResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyData) SetUpdateTime(v string) *GetQualityCheckSchemeResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyData) SetUpdateUserName(v string) *GetQualityCheckSchemeResponseBodyData {
	s.UpdateUserName = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyData) SetVersion(v int64) *GetQualityCheckSchemeResponseBodyData {
	s.Version = &v
	return s
}

type GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList struct {
	CheckName *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	// example:
	//
	// 0
	CheckType *int64 `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// example:
	//
	// 1
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// 32
	SchemeId *int64 `json:"SchemeId,omitempty" xml:"SchemeId,omitempty"`
	// example:
	//
	// 20
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// 10
	SourceScore *int32 `json:"SourceScore,omitempty" xml:"SourceScore,omitempty"`
}

func (s GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) String() string {
	return tea.Prettify(s)
}

func (s GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) GoString() string {
	return s.String()
}

func (s *GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) SetCheckName(v string) *GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList {
	s.CheckName = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) SetCheckType(v int64) *GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList {
	s.CheckType = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) SetEnable(v int32) *GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList {
	s.Enable = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) SetSchemeId(v int64) *GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList {
	s.SchemeId = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) SetScore(v int32) *GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList {
	s.Score = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) SetSourceScore(v int32) *GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList {
	s.SourceScore = &v
	return s
}

type GetQualityCheckSchemeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQualityCheckSchemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQualityCheckSchemeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQualityCheckSchemeResponse) GoString() string {
	return s.String()
}

func (s *GetQualityCheckSchemeResponse) SetHeaders(v map[string]*string) *GetQualityCheckSchemeResponse {
	s.Headers = v
	return s
}

func (s *GetQualityCheckSchemeResponse) SetStatusCode(v int32) *GetQualityCheckSchemeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQualityCheckSchemeResponse) SetBody(v *GetQualityCheckSchemeResponseBody) *GetQualityCheckSchemeResponse {
	s.Body = v
	return s
}

type GetResultRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"pageNumber":1,"pageSize":10,"excludeFields":"hitResult.hits, recording.url","requiredFields":"agent,status,errorMessage,reviewStatus,reviewResult,score,taskId,reviewer,resolver,recording.name,recording.duration,hitResult,business","dataType":1,"sourceType":0,"startTime":"2020-06-25 00:00:00","endTime":"2020-07-01 23:59:59"}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResultRequest) GoString() string {
	return s.String()
}

func (s *GetResultRequest) SetBaseMeAgentId(v int64) *GetResultRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetResultRequest) SetJsonStr(v string) *GetResultRequest {
	s.JsonStr = &v
	return s
}

type GetResultResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Count *int32                     `json:"Count,omitempty" xml:"Count,omitempty"`
	Data  *GetResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 3CEA0495-341B-4482-9AD9-8191EF4***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// XXX
	ResultCountId *string `json:"ResultCountId,omitempty" xml:"ResultCountId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetResultResponseBody) SetCode(v string) *GetResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetResultResponseBody) SetCount(v int32) *GetResultResponseBody {
	s.Count = &v
	return s
}

func (s *GetResultResponseBody) SetData(v *GetResultResponseBodyData) *GetResultResponseBody {
	s.Data = v
	return s
}

func (s *GetResultResponseBody) SetMessage(v string) *GetResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetResultResponseBody) SetPageNumber(v int32) *GetResultResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetResultResponseBody) SetPageSize(v int32) *GetResultResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetResultResponseBody) SetRequestId(v string) *GetResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResultResponseBody) SetResultCountId(v string) *GetResultResponseBody {
	s.ResultCountId = &v
	return s
}

func (s *GetResultResponseBody) SetSuccess(v bool) *GetResultResponseBody {
	s.Success = &v
	return s
}

type GetResultResponseBodyData struct {
	ResultInfo []*GetResultResponseBodyDataResultInfo `json:"ResultInfo,omitempty" xml:"ResultInfo,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyData) SetResultInfo(v []*GetResultResponseBodyDataResultInfo) *GetResultResponseBodyData {
	s.ResultInfo = v
	return s
}

type GetResultResponseBodyDataResultInfo struct {
	Agent     *GetResultResponseBodyDataResultInfoAgent     `json:"Agent,omitempty" xml:"Agent,omitempty" type:"Struct"`
	AsrResult *GetResultResponseBodyDataResultInfoAsrResult `json:"AsrResult,omitempty" xml:"AsrResult,omitempty" type:"Struct"`
	// example:
	//
	// 2021-03-02T14:37Z
	AssignmentTime *string `json:"AssignmentTime,omitempty" xml:"AssignmentTime,omitempty"`
	// example:
	//
	// xx
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// example:
	//
	// 2019-07-24T19:31Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1602743090
	CreateTimeLong *string `json:"CreateTimeLong,omitempty" xml:"CreateTimeLong,omitempty"`
	// example:
	//
	// xxx
	ErrorMessage *string                                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	HitResult    *GetResultResponseBodyDataResultInfoHitResult `json:"HitResult,omitempty" xml:"HitResult,omitempty" type:"Struct"`
	HitScore     *GetResultResponseBodyDataResultInfoHitScore  `json:"HitScore,omitempty" xml:"HitScore,omitempty" type:"Struct"`
	// example:
	//
	// 4498420@a_z@93EAADF1-01D3-44BD-8AC9-F57F447EFCE8_1614*****
	LastDataId *string                                       `json:"LastDataId,omitempty" xml:"LastDataId,omitempty"`
	Recording  *GetResultResponseBodyDataResultInfoRecording `json:"Recording,omitempty" xml:"Recording,omitempty" type:"Struct"`
	// example:
	//
	// XXX
	Resolver          *string                                               `json:"Resolver,omitempty" xml:"Resolver,omitempty"`
	ReviewHistoryList *GetResultResponseBodyDataResultInfoReviewHistoryList `json:"ReviewHistoryList,omitempty" xml:"ReviewHistoryList,omitempty" type:"Struct"`
	// example:
	//
	// 0
	ReviewResult *int32 `json:"ReviewResult,omitempty" xml:"ReviewResult,omitempty"`
	// example:
	//
	// 1
	ReviewStatus *int32 `json:"ReviewStatus,omitempty" xml:"ReviewStatus,omitempty"`
	// example:
	//
	// 2019-07-24T19:31Z
	ReviewTime *string `json:"ReviewTime,omitempty" xml:"ReviewTime,omitempty"`
	// example:
	//
	// 1602743090
	ReviewTimeLong *string `json:"ReviewTimeLong,omitempty" xml:"ReviewTimeLong,omitempty"`
	// example:
	//
	// 1
	ReviewType       *int32                                               `json:"ReviewType,omitempty" xml:"ReviewType,omitempty"`
	ReviewTypeIdList *GetResultResponseBodyDataResultInfoReviewTypeIdList `json:"ReviewTypeIdList,omitempty" xml:"ReviewTypeIdList,omitempty" type:"Struct"`
	// example:
	//
	// xxx
	Reviewer       *string                                            `json:"Reviewer,omitempty" xml:"Reviewer,omitempty"`
	SchemeIdList   *GetResultResponseBodyDataResultInfoSchemeIdList   `json:"SchemeIdList,omitempty" xml:"SchemeIdList,omitempty" type:"Struct"`
	SchemeNameList *GetResultResponseBodyDataResultInfoSchemeNameList `json:"SchemeNameList,omitempty" xml:"SchemeNameList,omitempty" type:"Struct"`
	// example:
	//
	// 100
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// A6BEC8D-9A5B-4BE5-8432-4F635E***
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// test
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s GetResultResponseBodyDataResultInfo) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfo) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfo) SetAgent(v *GetResultResponseBodyDataResultInfoAgent) *GetResultResponseBodyDataResultInfo {
	s.Agent = v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetAsrResult(v *GetResultResponseBodyDataResultInfoAsrResult) *GetResultResponseBodyDataResultInfo {
	s.AsrResult = v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetAssignmentTime(v string) *GetResultResponseBodyDataResultInfo {
	s.AssignmentTime = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetComments(v string) *GetResultResponseBodyDataResultInfo {
	s.Comments = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetCreateTime(v string) *GetResultResponseBodyDataResultInfo {
	s.CreateTime = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetCreateTimeLong(v string) *GetResultResponseBodyDataResultInfo {
	s.CreateTimeLong = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetErrorMessage(v string) *GetResultResponseBodyDataResultInfo {
	s.ErrorMessage = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetHitResult(v *GetResultResponseBodyDataResultInfoHitResult) *GetResultResponseBodyDataResultInfo {
	s.HitResult = v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetHitScore(v *GetResultResponseBodyDataResultInfoHitScore) *GetResultResponseBodyDataResultInfo {
	s.HitScore = v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetLastDataId(v string) *GetResultResponseBodyDataResultInfo {
	s.LastDataId = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetRecording(v *GetResultResponseBodyDataResultInfoRecording) *GetResultResponseBodyDataResultInfo {
	s.Recording = v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetResolver(v string) *GetResultResponseBodyDataResultInfo {
	s.Resolver = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetReviewHistoryList(v *GetResultResponseBodyDataResultInfoReviewHistoryList) *GetResultResponseBodyDataResultInfo {
	s.ReviewHistoryList = v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetReviewResult(v int32) *GetResultResponseBodyDataResultInfo {
	s.ReviewResult = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetReviewStatus(v int32) *GetResultResponseBodyDataResultInfo {
	s.ReviewStatus = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetReviewTime(v string) *GetResultResponseBodyDataResultInfo {
	s.ReviewTime = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetReviewTimeLong(v string) *GetResultResponseBodyDataResultInfo {
	s.ReviewTimeLong = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetReviewType(v int32) *GetResultResponseBodyDataResultInfo {
	s.ReviewType = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetReviewTypeIdList(v *GetResultResponseBodyDataResultInfoReviewTypeIdList) *GetResultResponseBodyDataResultInfo {
	s.ReviewTypeIdList = v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetReviewer(v string) *GetResultResponseBodyDataResultInfo {
	s.Reviewer = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetSchemeIdList(v *GetResultResponseBodyDataResultInfoSchemeIdList) *GetResultResponseBodyDataResultInfo {
	s.SchemeIdList = v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetSchemeNameList(v *GetResultResponseBodyDataResultInfoSchemeNameList) *GetResultResponseBodyDataResultInfo {
	s.SchemeNameList = v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetScore(v int32) *GetResultResponseBodyDataResultInfo {
	s.Score = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetStatus(v int32) *GetResultResponseBodyDataResultInfo {
	s.Status = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetTaskId(v string) *GetResultResponseBodyDataResultInfo {
	s.TaskId = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetTaskName(v string) *GetResultResponseBodyDataResultInfo {
	s.TaskName = &v
	return s
}

type GetResultResponseBodyDataResultInfoAgent struct {
	// example:
	//
	// 28240****15643
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SkillGroup *string `json:"SkillGroup,omitempty" xml:"SkillGroup,omitempty"`
}

func (s GetResultResponseBodyDataResultInfoAgent) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoAgent) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoAgent) SetId(v string) *GetResultResponseBodyDataResultInfoAgent {
	s.Id = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoAgent) SetName(v string) *GetResultResponseBodyDataResultInfoAgent {
	s.Name = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoAgent) SetSkillGroup(v string) *GetResultResponseBodyDataResultInfoAgent {
	s.SkillGroup = &v
	return s
}

type GetResultResponseBodyDataResultInfoAsrResult struct {
	AsrResult []*GetResultResponseBodyDataResultInfoAsrResultAsrResult `json:"AsrResult,omitempty" xml:"AsrResult,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoAsrResult) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoAsrResult) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoAsrResult) SetAsrResult(v []*GetResultResponseBodyDataResultInfoAsrResultAsrResult) *GetResultResponseBodyDataResultInfoAsrResult {
	s.AsrResult = v
	return s
}

type GetResultResponseBodyDataResultInfoAsrResultAsrResult struct {
	// example:
	//
	// 10000
	Begin *int64 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// example:
	//
	// 1
	EmotionValue *int32 `json:"EmotionValue,omitempty" xml:"EmotionValue,omitempty"`
	// example:
	//
	// 0
	End  *int64  `json:"End,omitempty" xml:"End,omitempty"`
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// 50
	SpeechRate *int32 `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	// example:
	//
	// xx
	Words *string `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s GetResultResponseBodyDataResultInfoAsrResultAsrResult) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoAsrResultAsrResult) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoAsrResultAsrResult) SetBegin(v int64) *GetResultResponseBodyDataResultInfoAsrResultAsrResult {
	s.Begin = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoAsrResultAsrResult) SetEmotionValue(v int32) *GetResultResponseBodyDataResultInfoAsrResultAsrResult {
	s.EmotionValue = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoAsrResultAsrResult) SetEnd(v int64) *GetResultResponseBodyDataResultInfoAsrResultAsrResult {
	s.End = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoAsrResultAsrResult) SetRole(v string) *GetResultResponseBodyDataResultInfoAsrResultAsrResult {
	s.Role = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoAsrResultAsrResult) SetSpeechRate(v int32) *GetResultResponseBodyDataResultInfoAsrResultAsrResult {
	s.SpeechRate = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoAsrResultAsrResult) SetWords(v string) *GetResultResponseBodyDataResultInfoAsrResultAsrResult {
	s.Words = &v
	return s
}

type GetResultResponseBodyDataResultInfoHitResult struct {
	HitResult []*GetResultResponseBodyDataResultInfoHitResultHitResult `json:"HitResult,omitempty" xml:"HitResult,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoHitResult) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResult) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResult) SetHitResult(v []*GetResultResponseBodyDataResultInfoHitResultHitResult) *GetResultResponseBodyDataResultInfoHitResult {
	s.HitResult = v
	return s
}

type GetResultResponseBodyDataResultInfoHitResultHitResult struct {
	Conditions *GetResultResponseBodyDataResultInfoHitResultHitResultConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Struct"`
	Hits       *GetResultResponseBodyDataResultInfoHitResultHitResultHits       `json:"Hits,omitempty" xml:"Hits,omitempty" type:"Struct"`
	Name       *string                                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0
	ReviewResult *int32 `json:"ReviewResult,omitempty" xml:"ReviewResult,omitempty"`
	// example:
	//
	// 1276
	Rid *string `json:"Rid,omitempty" xml:"Rid,omitempty"`
	// example:
	//
	// 123xx
	SchemeId *int64 `json:"SchemeId,omitempty" xml:"SchemeId,omitempty"`
	// example:
	//
	// 11xx
	SchemeVersion *int64  `json:"SchemeVersion,omitempty" xml:"SchemeVersion,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResult) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResult) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResult) SetConditions(v *GetResultResponseBodyDataResultInfoHitResultHitResultConditions) *GetResultResponseBodyDataResultInfoHitResultHitResult {
	s.Conditions = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResult) SetHits(v *GetResultResponseBodyDataResultInfoHitResultHitResultHits) *GetResultResponseBodyDataResultInfoHitResultHitResult {
	s.Hits = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResult) SetName(v string) *GetResultResponseBodyDataResultInfoHitResultHitResult {
	s.Name = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResult) SetReviewResult(v int32) *GetResultResponseBodyDataResultInfoHitResultHitResult {
	s.ReviewResult = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResult) SetRid(v string) *GetResultResponseBodyDataResultInfoHitResultHitResult {
	s.Rid = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResult) SetSchemeId(v int64) *GetResultResponseBodyDataResultInfoHitResultHitResult {
	s.SchemeId = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResult) SetSchemeVersion(v int64) *GetResultResponseBodyDataResultInfoHitResultHitResult {
	s.SchemeVersion = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResult) SetType(v string) *GetResultResponseBodyDataResultInfoHitResultHitResult {
	s.Type = &v
	return s
}

type GetResultResponseBodyDataResultInfoHitResultHitResultConditions struct {
	Conditions []*GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditions) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditions) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditions) SetConditions(v []*GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditions) *GetResultResponseBodyDataResultInfoHitResultHitResultConditions {
	s.Conditions = v
	return s
}

type GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditions struct {
	// 检测范围
	CheckRange *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRange `json:"Check_range,omitempty" xml:"Check_range,omitempty" type:"Struct"`
	// 条件id，可能是db中的主键，也可能是转换成的a, b, c
	Cid *string `json:"Cid,omitempty" xml:"Cid,omitempty"`
	// 排除
	Exclusion *int32 `json:"Exclusion,omitempty" xml:"Exclusion,omitempty"`
	// 在db中的主键
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Lambda表达式：例如:a&&b
	Lambda *string `json:"Lambda,omitempty" xml:"Lambda,omitempty"`
	// 算子列表
	Operators *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperators `json:"Operators,omitempty" xml:"Operators,omitempty" type:"Struct"`
	// 条件所属的规则id
	Rid *string `json:"Rid,omitempty" xml:"Rid,omitempty"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditions) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditions) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditions) SetCheckRange(v *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRange) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditions {
	s.CheckRange = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditions) SetCid(v string) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditions {
	s.Cid = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditions) SetExclusion(v int32) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditions {
	s.Exclusion = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditions) SetId(v int64) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditions {
	s.Id = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditions) SetLambda(v string) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditions {
	s.Lambda = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditions) SetOperators(v *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperators) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditions {
	s.Operators = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditions) SetRid(v string) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditions {
	s.Rid = &v
	return s
}

type GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRange struct {
	// false: 相对位置; 会结合anchor以及角色来决定句子位置
	Absolute *bool `json:"Absolute,omitempty" xml:"Absolute,omitempty"`
	// true: 每句话都必须满足条件；
	AllSentencesSatisfy *bool `json:"AllSentencesSatisfy,omitempty" xml:"AllSentencesSatisfy,omitempty"`
	// 前置后置条件
	Anchor *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeAnchor `json:"Anchor,omitempty" xml:"Anchor,omitempty" type:"Struct"`
	// 相对范围
	Range *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeRange `json:"Range,omitempty" xml:"Range,omitempty" type:"Struct"`
	// 对应 RoleType.type
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 对应 RoleType.id
	RoleId    *int32                                                                                        `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	TimeRange *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeTimeRange `json:"TimeRange,omitempty" xml:"TimeRange,omitempty" type:"Struct"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRange) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRange) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRange) SetAbsolute(v bool) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRange {
	s.Absolute = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRange) SetAllSentencesSatisfy(v bool) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRange {
	s.AllSentencesSatisfy = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRange) SetAnchor(v *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeAnchor) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRange {
	s.Anchor = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRange) SetRange(v *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeRange) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRange {
	s.Range = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRange) SetRole(v string) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRange {
	s.Role = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRange) SetRoleId(v int32) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRange {
	s.RoleId = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRange) SetTimeRange(v *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeTimeRange) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRange {
	s.TimeRange = v
	return s
}

type GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeAnchor struct {
	// 条件ID
	Cid *string `json:"Cid,omitempty" xml:"Cid,omitempty"`
	// 命中次数
	HitTime *int32 `json:"Hit_time,omitempty" xml:"Hit_time,omitempty"`
	// 位置
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeAnchor) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeAnchor) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeAnchor) SetCid(v string) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeAnchor {
	s.Cid = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeAnchor) SetHitTime(v int32) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeAnchor {
	s.HitTime = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeAnchor) SetLocation(v string) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeAnchor {
	s.Location = &v
	return s
}

type GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeRange struct {
	// 对话开始索引
	From *int32 `json:"From,omitempty" xml:"From,omitempty"`
	// 对话结束索引
	To *int32 `json:"To,omitempty" xml:"To,omitempty"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeRange) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeRange) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeRange) SetFrom(v int32) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeRange {
	s.From = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeRange) SetTo(v int32) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeRange {
	s.To = &v
	return s
}

type GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeTimeRange struct {
	From *int64 `json:"From,omitempty" xml:"From,omitempty"`
	To   *int64 `json:"To,omitempty" xml:"To,omitempty"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeTimeRange) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeTimeRange) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeTimeRange) SetFrom(v int64) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeTimeRange {
	s.From = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeTimeRange) SetTo(v int64) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeTimeRange {
	s.To = &v
	return s
}

type GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperators struct {
	Operator []*GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperator `json:"Operator,omitempty" xml:"Operator,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperators) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperators) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperators) SetOperator(v []*GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperator) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperators {
	s.Operator = v
	return s
}

type GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperator struct {
	// 主键id
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 算子名
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 可能是主键id，也可能是前端生成的id
	Oid *string `json:"Oid,omitempty" xml:"Oid,omitempty"`
	// 算子参数
	Param *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam `json:"Param,omitempty" xml:"Param,omitempty" type:"Struct"`
	// 算子类别
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperator) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperator) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperator) SetId(v int64) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperator {
	s.Id = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperator) SetName(v string) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperator {
	s.Name = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperator) SetOid(v string) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperator {
	s.Oid = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperator) SetParam(v *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperator {
	s.Param = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperator) SetType(v string) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperator {
	s.Type = &v
	return s
}

type GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam struct {
	// 语速检测，是否计算整个对话平均语速，默认false
	Average *bool `json:"Average,omitempty" xml:"Average,omitempty"`
	// 时长算子，时长计算开始类型，录音开始，还是某句对话开始
	BeginType *string `json:"BeginType,omitempty" xml:"BeginType,omitempty"`
	// 区分大小写
	CaseSensitive *bool `json:"Case_sensitive,omitempty" xml:"Case_sensitive,omitempty"`
	// 静音检测：要不要检测第一句话
	CheckFirstSentence *bool `json:"CheckFirstSentence,omitempty" xml:"CheckFirstSentence,omitempty"`
	// 检测方式，1 相邻句能量波动 2 最大能量跨度 默认1
	CheckType *int32 `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// 大于，还是小于，gt/lt
	CompareOperator *string `json:"CompareOperator,omitempty" xml:"CompareOperator,omitempty"`
	// 是否单句话匹配；
	ContextChatMatch *bool `json:"ContextChatMatch,omitempty" xml:"ContextChatMatch,omitempty"`
	// 抢话算子 延时时长
	DelayTime *int32 `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	// 时长算子，时长计算结束类型，录音结束，还是某句对话结束
	EndType *string `json:"EndType,omitempty" xml:"EndType,omitempty"`
	// 上下文重复算子：排除掉某些对话
	Excludes *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamExcludes `json:"Excludes,omitempty" xml:"Excludes,omitempty" type:"Struct"`
	// 流程节点前置条件参数
	FlowNodePrerequisiteParam *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamFlowNodePrerequisiteParam `json:"FlowNodePrerequisiteParam,omitempty" xml:"FlowNodePrerequisiteParam,omitempty" type:"Struct"`
	// 上下文重复算子：检测当前句的前from句是否有重复；0表示前面的所有句
	From *int32 `json:"From,omitempty" xml:"From,omitempty"`
	// from_end
	FromEnd *bool `json:"From_end,omitempty" xml:"From_end,omitempty"`
	// 上下文重复算子：重复几次
	HitTime *int32 `json:"Hit_time,omitempty" xml:"Hit_time,omitempty"`
	// 生效句子， true单个句子，false多个句子
	InSentence *bool `json:"In_sentence,omitempty" xml:"In_sentence,omitempty"`
	// 意图模型检查参数
	IntentModelCheckParm *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParm `json:"IntentModelCheckParm,omitempty" xml:"IntentModelCheckParm,omitempty" type:"Struct"`
	// interval代表区间范围开始
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// intervalEnd 代表区间范围结束
	IntervalEnd *int32 `json:"IntervalEnd,omitempty" xml:"IntervalEnd,omitempty"`
	// 关键字扩展
	KeywordExtension *int32 `json:"KeywordExtension,omitempty" xml:"KeywordExtension,omitempty"`
	// 匹配到的关键字数量
	KeywordMatchSize *int32 `json:"KeywordMatchSize,omitempty" xml:"KeywordMatchSize,omitempty"`
	// 关键词
	Keywords *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamKeywords `json:"Keywords,omitempty" xml:"Keywords,omitempty" type:"Struct"`
	// 能量值变化，默认3, 1~9
	MaxEmotionChangeValue *int32 `json:"MaxEmotionChangeValue,omitempty" xml:"MaxEmotionChangeValue,omitempty"`
	// 句子中最少字数，小于此字数的句子不检查
	MinWordSize *int32 `json:"MinWordSize,omitempty" xml:"MinWordSize,omitempty"`
	// true表示取不同角色相邻的两句话，false表示取不同角色的第一句话比较响应时间（默认）
	NearDialogue *bool `json:"Near_dialogue,omitempty" xml:"Near_dialogue,omitempty"`
	// 排除的正则表达式
	NotRegex *string `json:"NotRegex,omitempty" xml:"NotRegex,omitempty"`
	// 语句
	Phrase *string `json:"Phrase,omitempty" xml:"Phrase,omitempty"`
	// 正则表达式
	Regex *string `json:"Regex,omitempty" xml:"Regex,omitempty"`
	// target
	Target *int32 `json:"Target,omitempty" xml:"Target,omitempty"`
	// 阈值
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) SetAverage(v bool) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam {
	s.Average = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) SetBeginType(v string) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam {
	s.BeginType = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) SetCaseSensitive(v bool) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam {
	s.CaseSensitive = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) SetCheckFirstSentence(v bool) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam {
	s.CheckFirstSentence = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) SetCheckType(v int32) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam {
	s.CheckType = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) SetCompareOperator(v string) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam {
	s.CompareOperator = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) SetContextChatMatch(v bool) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam {
	s.ContextChatMatch = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) SetDelayTime(v int32) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam {
	s.DelayTime = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) SetEndType(v string) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam {
	s.EndType = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) SetExcludes(v *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamExcludes) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam {
	s.Excludes = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) SetFlowNodePrerequisiteParam(v *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamFlowNodePrerequisiteParam) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam {
	s.FlowNodePrerequisiteParam = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) SetFrom(v int32) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam {
	s.From = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) SetFromEnd(v bool) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam {
	s.FromEnd = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) SetHitTime(v int32) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam {
	s.HitTime = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) SetInSentence(v bool) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam {
	s.InSentence = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) SetIntentModelCheckParm(v *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParm) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam {
	s.IntentModelCheckParm = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) SetInterval(v int32) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam {
	s.Interval = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) SetIntervalEnd(v int32) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam {
	s.IntervalEnd = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) SetKeywordExtension(v int32) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam {
	s.KeywordExtension = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) SetKeywordMatchSize(v int32) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam {
	s.KeywordMatchSize = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) SetKeywords(v *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamKeywords) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam {
	s.Keywords = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) SetMaxEmotionChangeValue(v int32) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam {
	s.MaxEmotionChangeValue = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) SetMinWordSize(v int32) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam {
	s.MinWordSize = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) SetNearDialogue(v bool) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam {
	s.NearDialogue = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) SetNotRegex(v string) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam {
	s.NotRegex = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) SetPhrase(v string) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam {
	s.Phrase = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) SetRegex(v string) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam {
	s.Regex = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) SetTarget(v int32) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam {
	s.Target = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) SetThreshold(v float32) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam {
	s.Threshold = &v
	return s
}

type GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamExcludes struct {
	Exclude []*string `json:"Exclude,omitempty" xml:"Exclude,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamExcludes) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamExcludes) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamExcludes) SetExclude(v []*string) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamExcludes {
	s.Exclude = v
	return s
}

type GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamFlowNodePrerequisiteParam struct {
	// 节点id
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// 节点匹配状态。
	NodeMatchStatus *int32 `json:"NodeMatchStatus,omitempty" xml:"NodeMatchStatus,omitempty"`
	// 冗余的节点名称
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamFlowNodePrerequisiteParam) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamFlowNodePrerequisiteParam) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamFlowNodePrerequisiteParam) SetNodeId(v int64) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamFlowNodePrerequisiteParam {
	s.NodeId = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamFlowNodePrerequisiteParam) SetNodeMatchStatus(v int32) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamFlowNodePrerequisiteParam {
	s.NodeMatchStatus = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamFlowNodePrerequisiteParam) SetNodeName(v string) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamFlowNodePrerequisiteParam {
	s.NodeName = &v
	return s
}

type GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParm struct {
	// 引用的意图模型
	Intents *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntents `json:"Intents,omitempty" xml:"Intents,omitempty" type:"Struct"`
	// 模型应用的场景 AGENT:客户场景、CUSTOMER:客服场景 (CUSTOMER: 客户场景, AGENT: 坐席场景)
	ModelScene *string `json:"ModelScene,omitempty" xml:"ModelScene,omitempty"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParm) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParm) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParm) SetIntents(v *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntents) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParm {
	s.Intents = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParm) SetModelScene(v string) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParm {
	s.ModelScene = &v
	return s
}

type GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntents struct {
	Intent []*GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntentsIntent `json:"Intent,omitempty" xml:"Intent,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntents) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntents) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntents) SetIntent(v []*GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntentsIntent) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntents {
	s.Intent = v
	return s
}

type GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntentsIntent struct {
	// 意图模型ID
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 意图模型名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntentsIntent) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntentsIntent) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntentsIntent) SetId(v int64) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntentsIntent {
	s.Id = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntentsIntent) SetName(v string) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntentsIntent {
	s.Name = &v
	return s
}

type GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamKeywords struct {
	Keyword []*string `json:"Keyword,omitempty" xml:"Keyword,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamKeywords) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamKeywords) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamKeywords) SetKeyword(v []*string) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamKeywords {
	s.Keyword = v
	return s
}

type GetResultResponseBodyDataResultInfoHitResultHitResultHits struct {
	Hit []*GetResultResponseBodyDataResultInfoHitResultHitResultHitsHit `json:"Hit,omitempty" xml:"Hit,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultHits) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultHits) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHits) SetHit(v []*GetResultResponseBodyDataResultInfoHitResultHitResultHitsHit) *GetResultResponseBodyDataResultInfoHitResultHitResultHits {
	s.Hit = v
	return s
}

type GetResultResponseBodyDataResultInfoHitResultHitResultHitsHit struct {
	Cid      *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitCid      `json:"Cid,omitempty" xml:"Cid,omitempty" type:"Struct"`
	KeyWords *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWords `json:"KeyWords,omitempty" xml:"KeyWords,omitempty" type:"Struct"`
	Phrase   *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase   `json:"Phrase,omitempty" xml:"Phrase,omitempty" type:"Struct"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultHitsHit) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultHitsHit) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHit) SetCid(v *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitCid) *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHit {
	s.Cid = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHit) SetKeyWords(v *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWords) *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHit {
	s.KeyWords = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHit) SetPhrase(v *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase) *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHit {
	s.Phrase = v
	return s
}

type GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitCid struct {
	Cid []*string `json:"Cid,omitempty" xml:"Cid,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitCid) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitCid) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitCid) SetCid(v []*string) *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitCid {
	s.Cid = v
	return s
}

type GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWords struct {
	KeyWord []*GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWordsKeyWord `json:"KeyWord,omitempty" xml:"KeyWord,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWords) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWords) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWords) SetKeyWord(v []*GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWordsKeyWord) *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWords {
	s.KeyWord = v
	return s
}

type GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWordsKeyWord struct {
	// example:
	//
	// xxxx
	Cid *string `json:"Cid,omitempty" xml:"Cid,omitempty"`
	// example:
	//
	// 1
	From *int32 `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// 2
	To *int32 `json:"To,omitempty" xml:"To,omitempty"`
	// example:
	//
	// test
	Val *string `json:"Val,omitempty" xml:"Val,omitempty"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWordsKeyWord) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWordsKeyWord) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWordsKeyWord) SetCid(v string) *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWordsKeyWord {
	s.Cid = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWordsKeyWord) SetFrom(v int32) *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWordsKeyWord {
	s.From = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWordsKeyWord) SetTo(v int32) *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWordsKeyWord {
	s.To = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWordsKeyWord) SetVal(v string) *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWordsKeyWord {
	s.Val = &v
	return s
}

type GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase struct {
	// example:
	//
	// 300
	Begin *int64 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// example:
	//
	// 0
	EmotionValue *int32 `json:"EmotionValue,omitempty" xml:"EmotionValue,omitempty"`
	// example:
	//
	// 300
	End  *int32  `json:"End,omitempty" xml:"End,omitempty"`
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// xxx
	Words *string `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase) SetBegin(v int64) *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase {
	s.Begin = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase) SetEmotionValue(v int32) *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase {
	s.EmotionValue = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase) SetEnd(v int32) *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase {
	s.End = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase) SetRole(v string) *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase {
	s.Role = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase) SetWords(v string) *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase {
	s.Words = &v
	return s
}

type GetResultResponseBodyDataResultInfoHitScore struct {
	HitScore []*GetResultResponseBodyDataResultInfoHitScoreHitScore `json:"HitScore,omitempty" xml:"HitScore,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoHitScore) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitScore) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitScore) SetHitScore(v []*GetResultResponseBodyDataResultInfoHitScoreHitScore) *GetResultResponseBodyDataResultInfoHitScore {
	s.HitScore = v
	return s
}

type GetResultResponseBodyDataResultInfoHitScoreHitScore struct {
	// example:
	//
	// 123
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// 123456
	ScoreId   *string `json:"ScoreId,omitempty" xml:"ScoreId,omitempty"`
	ScoreName *string `json:"ScoreName,omitempty" xml:"ScoreName,omitempty"`
	// example:
	//
	// -20
	ScoreNumber *string `json:"ScoreNumber,omitempty" xml:"ScoreNumber,omitempty"`
}

func (s GetResultResponseBodyDataResultInfoHitScoreHitScore) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitScoreHitScore) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitScoreHitScore) SetRuleId(v string) *GetResultResponseBodyDataResultInfoHitScoreHitScore {
	s.RuleId = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitScoreHitScore) SetScoreId(v string) *GetResultResponseBodyDataResultInfoHitScoreHitScore {
	s.ScoreId = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitScoreHitScore) SetScoreName(v string) *GetResultResponseBodyDataResultInfoHitScoreHitScore {
	s.ScoreName = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitScoreHitScore) SetScoreNumber(v string) *GetResultResponseBodyDataResultInfoHitScoreHitScore {
	s.ScoreNumber = &v
	return s
}

type GetResultResponseBodyDataResultInfoRecording struct {
	Business *string `json:"Business,omitempty" xml:"Business,omitempty"`
	// example:
	//
	// XXXX
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// example:
	//
	// 1563967699000
	CallTime *string `json:"CallTime,omitempty" xml:"CallTime,omitempty"`
	// example:
	//
	// 1
	CallType *int32 `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// example:
	//
	// 1888888****
	Callee *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	// example:
	//
	// 0108888****
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// example:
	//
	// XXXX
	DataSetName *string `json:"DataSetName,omitempty" xml:"DataSetName,omitempty"`
	// example:
	//
	// 32
	DialogueSize *int32 `json:"DialogueSize,omitempty" xml:"DialogueSize,omitempty"`
	// example:
	//
	// 60
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// XXXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 123456.mkv
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 3437500
	PrimaryId *string `json:"PrimaryId,omitempty" xml:"PrimaryId,omitempty"`
	// example:
	//
	// XXX
	Remark1 *string `json:"Remark1,omitempty" xml:"Remark1,omitempty"`
	// example:
	//
	// XXX
	Remark10 *string `json:"Remark10,omitempty" xml:"Remark10,omitempty"`
	// example:
	//
	// XXX
	Remark11 *string `json:"Remark11,omitempty" xml:"Remark11,omitempty"`
	// example:
	//
	// XXX
	Remark12 *string `json:"Remark12,omitempty" xml:"Remark12,omitempty"`
	// example:
	//
	// XXX
	Remark13 *string `json:"Remark13,omitempty" xml:"Remark13,omitempty"`
	// example:
	//
	// XXX
	Remark2 *string `json:"Remark2,omitempty" xml:"Remark2,omitempty"`
	// example:
	//
	// XXX
	Remark3 *string `json:"Remark3,omitempty" xml:"Remark3,omitempty"`
	// example:
	//
	// XXX
	Remark4 *string `json:"Remark4,omitempty" xml:"Remark4,omitempty"`
	// example:
	//
	// 1232
	Remark5 *int64 `json:"Remark5,omitempty" xml:"Remark5,omitempty"`
	// example:
	//
	// XXX
	Remark6 *string `json:"Remark6,omitempty" xml:"Remark6,omitempty"`
	// example:
	//
	// XXX
	Remark7 *string `json:"Remark7,omitempty" xml:"Remark7,omitempty"`
	// example:
	//
	// XXX
	Remark8 *string `json:"Remark8,omitempty" xml:"Remark8,omitempty"`
	// example:
	//
	// XXX
	Remark9        *string `json:"Remark9,omitempty" xml:"Remark9,omitempty"`
	TaskConfigId   *int64  `json:"TaskConfigId,omitempty" xml:"TaskConfigId,omitempty"`
	TaskConfigName *string `json:"TaskConfigName,omitempty" xml:"TaskConfigName,omitempty"`
	// example:
	//
	// http://aliyun.com/audio.wav
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetResultResponseBodyDataResultInfoRecording) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoRecording) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetBusiness(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Business = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetCallId(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.CallId = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetCallTime(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.CallTime = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetCallType(v int32) *GetResultResponseBodyDataResultInfoRecording {
	s.CallType = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetCallee(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Callee = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetCaller(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Caller = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetDataSetName(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.DataSetName = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetDialogueSize(v int32) *GetResultResponseBodyDataResultInfoRecording {
	s.DialogueSize = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetDuration(v int64) *GetResultResponseBodyDataResultInfoRecording {
	s.Duration = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetId(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Id = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetName(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Name = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetPrimaryId(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.PrimaryId = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetRemark1(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Remark1 = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetRemark10(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Remark10 = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetRemark11(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Remark11 = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetRemark12(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Remark12 = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetRemark13(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Remark13 = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetRemark2(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Remark2 = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetRemark3(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Remark3 = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetRemark4(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Remark4 = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetRemark5(v int64) *GetResultResponseBodyDataResultInfoRecording {
	s.Remark5 = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetRemark6(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Remark6 = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetRemark7(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Remark7 = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetRemark8(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Remark8 = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetRemark9(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Remark9 = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetTaskConfigId(v int64) *GetResultResponseBodyDataResultInfoRecording {
	s.TaskConfigId = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetTaskConfigName(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.TaskConfigName = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetUrl(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Url = &v
	return s
}

type GetResultResponseBodyDataResultInfoReviewHistoryList struct {
	ReviewHistory []*GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory `json:"ReviewHistory,omitempty" xml:"ReviewHistory,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoReviewHistoryList) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoReviewHistoryList) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryList) SetReviewHistory(v []*GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory) *GetResultResponseBodyDataResultInfoReviewHistoryList {
	s.ReviewHistory = v
	return s
}

type GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory struct {
	Comments          *string                                                                           `json:"Comments,omitempty" xml:"Comments,omitempty"`
	ComplainResult    *int32                                                                            `json:"ComplainResult,omitempty" xml:"ComplainResult,omitempty"`
	OldScore          *int32                                                                            `json:"OldScore,omitempty" xml:"OldScore,omitempty"`
	Operator          *int64                                                                            `json:"Operator,omitempty" xml:"Operator,omitempty"`
	OperatorName      *string                                                                           `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	ReviewManagerType *string                                                                           `json:"ReviewManagerType,omitempty" xml:"ReviewManagerType,omitempty"`
	ReviewResult      *int32                                                                            `json:"ReviewResult,omitempty" xml:"ReviewResult,omitempty"`
	ReviewRightRule   *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRule `json:"ReviewRightRule,omitempty" xml:"ReviewRightRule,omitempty" type:"Struct"`
	Score             *int32                                                                            `json:"Score,omitempty" xml:"Score,omitempty"`
	Time              *int64                                                                            `json:"Time,omitempty" xml:"Time,omitempty"`
	TimeStr           *string                                                                           `json:"TimeStr,omitempty" xml:"TimeStr,omitempty"`
	Type              *int32                                                                            `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory) SetComments(v string) *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory {
	s.Comments = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory) SetComplainResult(v int32) *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory {
	s.ComplainResult = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory) SetOldScore(v int32) *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory {
	s.OldScore = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory) SetOperator(v int64) *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory {
	s.Operator = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory) SetOperatorName(v string) *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory {
	s.OperatorName = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory) SetReviewManagerType(v string) *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory {
	s.ReviewManagerType = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory) SetReviewResult(v int32) *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory {
	s.ReviewResult = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory) SetReviewRightRule(v *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRule) *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory {
	s.ReviewRightRule = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory) SetScore(v int32) *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory {
	s.Score = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory) SetTime(v int64) *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory {
	s.Time = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory) SetTimeStr(v string) *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory {
	s.TimeStr = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory) SetType(v int32) *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory {
	s.Type = &v
	return s
}

type GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRule struct {
	ReviewRightRule []*GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule `json:"ReviewRightRule,omitempty" xml:"ReviewRightRule,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRule) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRule) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRule) SetReviewRightRule(v []*GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule) *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRule {
	s.ReviewRightRule = v
	return s
}

type GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule struct {
	Rid      *int64  `json:"rid,omitempty" xml:"rid,omitempty"`
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
}

func (s GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule) SetRid(v int64) *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule {
	s.Rid = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule) SetRuleName(v string) *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule {
	s.RuleName = &v
	return s
}

type GetResultResponseBodyDataResultInfoReviewTypeIdList struct {
	ReviewTypeIdList []*GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdList `json:"ReviewTypeIdList,omitempty" xml:"ReviewTypeIdList,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoReviewTypeIdList) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoReviewTypeIdList) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoReviewTypeIdList) SetReviewTypeIdList(v []*GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdList) *GetResultResponseBodyDataResultInfoReviewTypeIdList {
	s.ReviewTypeIdList = v
	return s
}

type GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdList struct {
	ReviewKeyIdList *GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdListReviewKeyIdList `json:"ReviewKeyIdList,omitempty" xml:"ReviewKeyIdList,omitempty" type:"Struct"`
	ReviewTypeId    *int64                                                                              `json:"ReviewTypeId,omitempty" xml:"ReviewTypeId,omitempty"`
}

func (s GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdList) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdList) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdList) SetReviewKeyIdList(v *GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdListReviewKeyIdList) *GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdList {
	s.ReviewKeyIdList = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdList) SetReviewTypeId(v int64) *GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdList {
	s.ReviewTypeId = &v
	return s
}

type GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdListReviewKeyIdList struct {
	ReviewKeyIdList []*int64 `json:"ReviewKeyIdList,omitempty" xml:"ReviewKeyIdList,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdListReviewKeyIdList) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdListReviewKeyIdList) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdListReviewKeyIdList) SetReviewKeyIdList(v []*int64) *GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdListReviewKeyIdList {
	s.ReviewKeyIdList = v
	return s
}

type GetResultResponseBodyDataResultInfoSchemeIdList struct {
	SchemeIdList []*int64 `json:"SchemeIdList,omitempty" xml:"SchemeIdList,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoSchemeIdList) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoSchemeIdList) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoSchemeIdList) SetSchemeIdList(v []*int64) *GetResultResponseBodyDataResultInfoSchemeIdList {
	s.SchemeIdList = v
	return s
}

type GetResultResponseBodyDataResultInfoSchemeNameList struct {
	SchemeNameList []*string `json:"SchemeNameList,omitempty" xml:"SchemeNameList,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoSchemeNameList) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoSchemeNameList) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoSchemeNameList) SetSchemeNameList(v []*string) *GetResultResponseBodyDataResultInfoSchemeNameList {
	s.SchemeNameList = v
	return s
}

type GetResultResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponse) GoString() string {
	return s.String()
}

func (s *GetResultResponse) SetHeaders(v map[string]*string) *GetResultResponse {
	s.Headers = v
	return s
}

func (s *GetResultResponse) SetStatusCode(v int32) *GetResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResultResponse) SetBody(v *GetResultResponseBody) *GetResultResponse {
	s.Body = v
	return s
}

type GetResultToReviewRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetResultToReviewRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewRequest) GoString() string {
	return s.String()
}

func (s *GetResultToReviewRequest) SetBaseMeAgentId(v int64) *GetResultToReviewRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetResultToReviewRequest) SetJsonStr(v string) *GetResultToReviewRequest {
	s.JsonStr = &v
	return s
}

type GetResultToReviewResponseBody struct {
	// example:
	//
	// 200
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetResultToReviewResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 82C91484-B2D5-4D2A-A21F-A6D73F4***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetResultToReviewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBody) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBody) SetCode(v string) *GetResultToReviewResponseBody {
	s.Code = &v
	return s
}

func (s *GetResultToReviewResponseBody) SetData(v *GetResultToReviewResponseBodyData) *GetResultToReviewResponseBody {
	s.Data = v
	return s
}

func (s *GetResultToReviewResponseBody) SetMessage(v string) *GetResultToReviewResponseBody {
	s.Message = &v
	return s
}

func (s *GetResultToReviewResponseBody) SetRequestId(v string) *GetResultToReviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResultToReviewResponseBody) SetSuccess(v bool) *GetResultToReviewResponseBody {
	s.Success = &v
	return s
}

type GetResultToReviewResponseBodyData struct {
	// example:
	//
	// https
	AudioScheme *string `json:"AudioScheme,omitempty" xml:"AudioScheme,omitempty"`
	// example:
	//
	// sca-ccc-test.oss-cn-beijing.aliyuncs.com/xxxxx
	AudioURL *string `json:"AudioURL,omitempty" xml:"AudioURL,omitempty"`
	// example:
	//
	// xxx
	Comments  *string                                     `json:"Comments,omitempty" xml:"Comments,omitempty"`
	Dialogues *GetResultToReviewResponseBodyDataDialogues `json:"Dialogues,omitempty" xml:"Dialogues,omitempty" type:"Struct"`
	// example:
	//
	// e790e6c919d84b82b64ee*****
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// xxx.wav
	FileMergeName         *string                                                 `json:"FileMergeName,omitempty" xml:"FileMergeName,omitempty"`
	HitRuleReviewInfoList *GetResultToReviewResponseBodyDataHitRuleReviewInfoList `json:"HitRuleReviewInfoList,omitempty" xml:"HitRuleReviewInfoList,omitempty" type:"Struct"`
	ManualScoreInfoList   *GetResultToReviewResponseBodyDataManualScoreInfoList   `json:"ManualScoreInfoList,omitempty" xml:"ManualScoreInfoList,omitempty" type:"Struct"`
	ReviewHistoryList     *GetResultToReviewResponseBodyDataReviewHistoryList     `json:"ReviewHistoryList,omitempty" xml:"ReviewHistoryList,omitempty" type:"Struct"`
	ReviewTypeIdList      *GetResultToReviewResponseBodyDataReviewTypeIdList      `json:"ReviewTypeIdList,omitempty" xml:"ReviewTypeIdList,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 99
	TotalScore *int32 `json:"TotalScore,omitempty" xml:"TotalScore,omitempty"`
	// example:
	//
	// 6fa76916-3ce6-45d8-ac64-01b7f31***
	Vid *string `json:"Vid,omitempty" xml:"Vid,omitempty"`
}

func (s GetResultToReviewResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyData) SetAudioScheme(v string) *GetResultToReviewResponseBodyData {
	s.AudioScheme = &v
	return s
}

func (s *GetResultToReviewResponseBodyData) SetAudioURL(v string) *GetResultToReviewResponseBodyData {
	s.AudioURL = &v
	return s
}

func (s *GetResultToReviewResponseBodyData) SetComments(v string) *GetResultToReviewResponseBodyData {
	s.Comments = &v
	return s
}

func (s *GetResultToReviewResponseBodyData) SetDialogues(v *GetResultToReviewResponseBodyDataDialogues) *GetResultToReviewResponseBodyData {
	s.Dialogues = v
	return s
}

func (s *GetResultToReviewResponseBodyData) SetFileId(v string) *GetResultToReviewResponseBodyData {
	s.FileId = &v
	return s
}

func (s *GetResultToReviewResponseBodyData) SetFileMergeName(v string) *GetResultToReviewResponseBodyData {
	s.FileMergeName = &v
	return s
}

func (s *GetResultToReviewResponseBodyData) SetHitRuleReviewInfoList(v *GetResultToReviewResponseBodyDataHitRuleReviewInfoList) *GetResultToReviewResponseBodyData {
	s.HitRuleReviewInfoList = v
	return s
}

func (s *GetResultToReviewResponseBodyData) SetManualScoreInfoList(v *GetResultToReviewResponseBodyDataManualScoreInfoList) *GetResultToReviewResponseBodyData {
	s.ManualScoreInfoList = v
	return s
}

func (s *GetResultToReviewResponseBodyData) SetReviewHistoryList(v *GetResultToReviewResponseBodyDataReviewHistoryList) *GetResultToReviewResponseBodyData {
	s.ReviewHistoryList = v
	return s
}

func (s *GetResultToReviewResponseBodyData) SetReviewTypeIdList(v *GetResultToReviewResponseBodyDataReviewTypeIdList) *GetResultToReviewResponseBodyData {
	s.ReviewTypeIdList = v
	return s
}

func (s *GetResultToReviewResponseBodyData) SetStatus(v int32) *GetResultToReviewResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetResultToReviewResponseBodyData) SetTotalScore(v int32) *GetResultToReviewResponseBodyData {
	s.TotalScore = &v
	return s
}

func (s *GetResultToReviewResponseBodyData) SetVid(v string) *GetResultToReviewResponseBodyData {
	s.Vid = &v
	return s
}

type GetResultToReviewResponseBodyDataDialogues struct {
	Dialogue []*GetResultToReviewResponseBodyDataDialoguesDialogue `json:"Dialogue,omitempty" xml:"Dialogue,omitempty" type:"Repeated"`
}

func (s GetResultToReviewResponseBodyDataDialogues) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataDialogues) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataDialogues) SetDialogue(v []*GetResultToReviewResponseBodyDataDialoguesDialogue) *GetResultToReviewResponseBodyDataDialogues {
	s.Dialogue = v
	return s
}

type GetResultToReviewResponseBodyDataDialoguesDialogue struct {
	// example:
	//
	// 72000
	Begin *int64 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// example:
	//
	// 2019-10-01 11:12:01
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// 7
	EmotionValue *int32 `json:"EmotionValue,omitempty" xml:"EmotionValue,omitempty"`
	// example:
	//
	// 80000
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// example:
	//
	// 00:08
	HourMinSec *string `json:"HourMinSec,omitempty" xml:"HourMinSec,omitempty"`
	Identity   *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
	Role       *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// 1
	SilenceDuration *int32 `json:"SilenceDuration,omitempty" xml:"SilenceDuration,omitempty"`
	// example:
	//
	// 200
	SpeechRate *int32  `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	Words      *string `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s GetResultToReviewResponseBodyDataDialoguesDialogue) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataDialoguesDialogue) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataDialoguesDialogue) SetBegin(v int64) *GetResultToReviewResponseBodyDataDialoguesDialogue {
	s.Begin = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataDialoguesDialogue) SetBeginTime(v string) *GetResultToReviewResponseBodyDataDialoguesDialogue {
	s.BeginTime = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataDialoguesDialogue) SetEmotionValue(v int32) *GetResultToReviewResponseBodyDataDialoguesDialogue {
	s.EmotionValue = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataDialoguesDialogue) SetEnd(v int64) *GetResultToReviewResponseBodyDataDialoguesDialogue {
	s.End = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataDialoguesDialogue) SetHourMinSec(v string) *GetResultToReviewResponseBodyDataDialoguesDialogue {
	s.HourMinSec = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataDialoguesDialogue) SetIdentity(v string) *GetResultToReviewResponseBodyDataDialoguesDialogue {
	s.Identity = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataDialoguesDialogue) SetRole(v string) *GetResultToReviewResponseBodyDataDialoguesDialogue {
	s.Role = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataDialoguesDialogue) SetSilenceDuration(v int32) *GetResultToReviewResponseBodyDataDialoguesDialogue {
	s.SilenceDuration = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataDialoguesDialogue) SetSpeechRate(v int32) *GetResultToReviewResponseBodyDataDialoguesDialogue {
	s.SpeechRate = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataDialoguesDialogue) SetWords(v string) *GetResultToReviewResponseBodyDataDialoguesDialogue {
	s.Words = &v
	return s
}

type GetResultToReviewResponseBodyDataHitRuleReviewInfoList struct {
	HitRuleReviewInfo []*GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo `json:"HitRuleReviewInfo,omitempty" xml:"HitRuleReviewInfo,omitempty" type:"Repeated"`
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoList) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoList) SetHitRuleReviewInfo(v []*GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) *GetResultToReviewResponseBodyDataHitRuleReviewInfoList {
	s.HitRuleReviewInfo = v
	return s
}

type GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo struct {
	// example:
	//
	// 1
	AutoReview        *int32                                                                                    `json:"AutoReview,omitempty" xml:"AutoReview,omitempty"`
	ComplainHistories *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistories `json:"ComplainHistories,omitempty" xml:"ComplainHistories,omitempty" type:"Struct"`
	// example:
	//
	// true
	Complainable         *bool                                                                                        `json:"Complainable,omitempty" xml:"Complainable,omitempty"`
	ConditionHitInfoList *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoList `json:"ConditionHitInfoList,omitempty" xml:"ConditionHitInfoList,omitempty" type:"Struct"`
	ReviewInfo           *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo           `json:"ReviewInfo,omitempty" xml:"ReviewInfo,omitempty" type:"Struct"`
	// example:
	//
	// 451
	Rid      *int64  `json:"Rid,omitempty" xml:"Rid,omitempty"`
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// xxx
	ScoreId *int64 `json:"ScoreId,omitempty" xml:"ScoreId,omitempty"`
	// example:
	//
	// -10
	ScoreNum *int32 `json:"ScoreNum,omitempty" xml:"ScoreNum,omitempty"`
	// example:
	//
	// xxx
	ScoreSubId *int64 `json:"ScoreSubId,omitempty" xml:"ScoreSubId,omitempty"`
	// example:
	//
	// xxx
	ScoreSubName *string `json:"ScoreSubName,omitempty" xml:"ScoreSubName,omitempty"`
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) SetAutoReview(v int32) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo {
	s.AutoReview = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) SetComplainHistories(v *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistories) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo {
	s.ComplainHistories = v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) SetComplainable(v bool) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo {
	s.Complainable = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) SetConditionHitInfoList(v *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoList) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo {
	s.ConditionHitInfoList = v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) SetReviewInfo(v *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo {
	s.ReviewInfo = v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) SetRid(v int64) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo {
	s.Rid = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) SetRuleName(v string) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo {
	s.RuleName = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) SetScoreId(v int64) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo {
	s.ScoreId = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) SetScoreNum(v int32) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo {
	s.ScoreNum = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) SetScoreSubId(v int64) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo {
	s.ScoreSubId = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) SetScoreSubName(v string) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo {
	s.ScoreSubName = &v
	return s
}

type GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistories struct {
	ComplainHistories []*GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories `json:"ComplainHistories,omitempty" xml:"ComplainHistories,omitempty" type:"Repeated"`
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistories) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistories) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistories) SetComplainHistories(v []*GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistories {
	s.ComplainHistories = v
	return s
}

type GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories struct {
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// example:
	//
	// 2020-10-16T11:13Z
	OperationTime *string `json:"OperationTime,omitempty" xml:"OperationTime,omitempty"`
	// example:
	//
	// 5
	OperationType *int32 `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// example:
	//
	// 123456
	Operator     *int64  `json:"Operator,omitempty" xml:"Operator,omitempty"`
	OperatorName *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories) SetComments(v string) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories {
	s.Comments = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories) SetOperationTime(v string) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories {
	s.OperationTime = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories) SetOperationType(v int32) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories {
	s.OperationType = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories) SetOperator(v int64) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories {
	s.Operator = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories) SetOperatorName(v string) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories {
	s.OperatorName = &v
	return s
}

type GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoList struct {
	ConditionHitInfo []*GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfo `json:"ConditionHitInfo,omitempty" xml:"ConditionHitInfo,omitempty" type:"Repeated"`
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoList) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoList) SetConditionHitInfo(v []*GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfo) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoList {
	s.ConditionHitInfo = v
	return s
}

type GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfo struct {
	Cid      *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoCid      `json:"Cid,omitempty" xml:"Cid,omitempty" type:"Struct"`
	KeyWords *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWords `json:"KeyWords,omitempty" xml:"KeyWords,omitempty" type:"Struct"`
	Phrase   *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase   `json:"Phrase,omitempty" xml:"Phrase,omitempty" type:"Struct"`
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfo) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfo) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfo) SetCid(v *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoCid) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfo {
	s.Cid = v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfo) SetKeyWords(v *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWords) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfo {
	s.KeyWords = v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfo) SetPhrase(v *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfo {
	s.Phrase = v
	return s
}

type GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoCid struct {
	Cid []*string `json:"Cid,omitempty" xml:"Cid,omitempty" type:"Repeated"`
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoCid) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoCid) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoCid) SetCid(v []*string) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoCid {
	s.Cid = v
	return s
}

type GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWords struct {
	KeyWord []*GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord `json:"KeyWord,omitempty" xml:"KeyWord,omitempty" type:"Repeated"`
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWords) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWords) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWords) SetKeyWord(v []*GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWords {
	s.KeyWord = v
	return s
}

type GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord struct {
	// example:
	//
	// 2000
	Cid *string `json:"Cid,omitempty" xml:"Cid,omitempty"`
	// example:
	//
	// xxx
	CustomizeCode *string `json:"CustomizeCode,omitempty" xml:"CustomizeCode,omitempty"`
	// example:
	//
	// 1
	From *int32 `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// 2
	Pid *int32 `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// example:
	//
	// 6fa76916-3ce6-45d8-ac64-01b7f31c7295
	Tid *string `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// example:
	//
	// 3
	To  *int32  `json:"To,omitempty" xml:"To,omitempty"`
	Val *string `json:"Val,omitempty" xml:"Val,omitempty"`
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord) SetCid(v string) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord {
	s.Cid = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord) SetCustomizeCode(v string) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord {
	s.CustomizeCode = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord) SetFrom(v int32) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord {
	s.From = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord) SetPid(v int32) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord {
	s.Pid = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord) SetTid(v string) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord {
	s.Tid = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord) SetTo(v int32) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord {
	s.To = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord) SetVal(v string) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord {
	s.Val = &v
	return s
}

type GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase struct {
	// example:
	//
	// 72000
	Begin *int64 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// example:
	//
	// 7
	EmotionValue *int32 `json:"EmotionValue,omitempty" xml:"EmotionValue,omitempty"`
	// example:
	//
	// 80000
	End      *int64  `json:"End,omitempty" xml:"End,omitempty"`
	Identity *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
	// example:
	//
	// 3
	Pid   *int32  `json:"Pid,omitempty" xml:"Pid,omitempty"`
	Role  *string `json:"Role,omitempty" xml:"Role,omitempty"`
	Words *string `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase) SetBegin(v int64) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase {
	s.Begin = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase) SetEmotionValue(v int32) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase {
	s.EmotionValue = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase) SetEnd(v int64) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase {
	s.End = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase) SetIdentity(v string) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase {
	s.Identity = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase) SetPid(v int32) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase {
	s.Pid = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase) SetRole(v string) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase {
	s.Role = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase) SetWords(v string) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase {
	s.Words = &v
	return s
}

type GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo struct {
	// example:
	//
	// 013c68142fec4f0899fa6ee0exxx
	HitId *string `json:"HitId,omitempty" xml:"HitId,omitempty"`
	// example:
	//
	// 1
	ReviewResult *int32 `json:"ReviewResult,omitempty" xml:"ReviewResult,omitempty"`
	// example:
	//
	// 2019-10-12 17:06:00
	ReviewTime *string `json:"ReviewTime,omitempty" xml:"ReviewTime,omitempty"`
	// example:
	//
	// 123
	Reviewer *string `json:"Reviewer,omitempty" xml:"Reviewer,omitempty"`
	// example:
	//
	// 451
	Rid *int64 `json:"Rid,omitempty" xml:"Rid,omitempty"`
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo) SetHitId(v string) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo {
	s.HitId = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo) SetReviewResult(v int32) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo {
	s.ReviewResult = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo) SetReviewTime(v string) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo {
	s.ReviewTime = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo) SetReviewer(v string) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo {
	s.Reviewer = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo) SetRid(v int64) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo {
	s.Rid = &v
	return s
}

type GetResultToReviewResponseBodyDataManualScoreInfoList struct {
	ManualScoreInfo []*GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo `json:"ManualScoreInfo,omitempty" xml:"ManualScoreInfo,omitempty" type:"Repeated"`
}

func (s GetResultToReviewResponseBodyDataManualScoreInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataManualScoreInfoList) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoList) SetManualScoreInfo(v []*GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo) *GetResultToReviewResponseBodyDataManualScoreInfoList {
	s.ManualScoreInfo = v
	return s
}

type GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo struct {
	ComplainHistories *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistories `json:"ComplainHistories,omitempty" xml:"ComplainHistories,omitempty" type:"Struct"`
	// example:
	//
	// true
	Complainable *bool `json:"Complainable,omitempty" xml:"Complainable,omitempty"`
	// example:
	//
	// xxx
	ScoreId *int64 `json:"ScoreId,omitempty" xml:"ScoreId,omitempty"`
	// example:
	//
	// -10
	ScoreNum *int32 `json:"ScoreNum,omitempty" xml:"ScoreNum,omitempty"`
	// example:
	//
	// xxx
	ScoreSubId   *int64  `json:"ScoreSubId,omitempty" xml:"ScoreSubId,omitempty"`
	ScoreSubName *string `json:"ScoreSubName,omitempty" xml:"ScoreSubName,omitempty"`
}

func (s GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo) SetComplainHistories(v *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistories) *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo {
	s.ComplainHistories = v
	return s
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo) SetComplainable(v bool) *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo {
	s.Complainable = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo) SetScoreId(v int64) *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo {
	s.ScoreId = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo) SetScoreNum(v int32) *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo {
	s.ScoreNum = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo) SetScoreSubId(v int64) *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo {
	s.ScoreSubId = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo) SetScoreSubName(v string) *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo {
	s.ScoreSubName = &v
	return s
}

type GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistories struct {
	ComplainHistories []*GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories `json:"ComplainHistories,omitempty" xml:"ComplainHistories,omitempty" type:"Repeated"`
}

func (s GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistories) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistories) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistories) SetComplainHistories(v []*GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories) *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistories {
	s.ComplainHistories = v
	return s
}

type GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories struct {
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// example:
	//
	// 2020-10-16T11:13Z
	OperationTime *string `json:"OperationTime,omitempty" xml:"OperationTime,omitempty"`
	// example:
	//
	// 5
	OperationType *int32 `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// example:
	//
	// 123456
	Operator     *int64  `json:"Operator,omitempty" xml:"Operator,omitempty"`
	OperatorName *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
}

func (s GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories) SetComments(v string) *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories {
	s.Comments = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories) SetOperationTime(v string) *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories {
	s.OperationTime = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories) SetOperationType(v int32) *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories {
	s.OperationType = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories) SetOperator(v int64) *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories {
	s.Operator = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories) SetOperatorName(v string) *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories {
	s.OperatorName = &v
	return s
}

type GetResultToReviewResponseBodyDataReviewHistoryList struct {
	ReviewHistory []*GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory `json:"ReviewHistory,omitempty" xml:"ReviewHistory,omitempty" type:"Repeated"`
}

func (s GetResultToReviewResponseBodyDataReviewHistoryList) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataReviewHistoryList) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryList) SetReviewHistory(v []*GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) *GetResultToReviewResponseBodyDataReviewHistoryList {
	s.ReviewHistory = v
	return s
}

type GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory struct {
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// example:
	//
	// 1
	ComplainResult *int32 `json:"ComplainResult,omitempty" xml:"ComplainResult,omitempty"`
	// example:
	//
	// 90
	OldScore          *int32  `json:"OldScore,omitempty" xml:"OldScore,omitempty"`
	Operator          *int64  `json:"Operator,omitempty" xml:"Operator,omitempty"`
	OperatorName      *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	ReviewManagerType *string `json:"ReviewManagerType,omitempty" xml:"ReviewManagerType,omitempty"`
	// example:
	//
	// 1
	ReviewResult    *int32                                                                          `json:"ReviewResult,omitempty" xml:"ReviewResult,omitempty"`
	ReviewRightRule *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRule `json:"ReviewRightRule,omitempty" xml:"ReviewRightRule,omitempty" type:"Struct"`
	// example:
	//
	// 95
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	Time  *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// example:
	//
	// 2019-10-28 15:21:00
	TimeStr *string `json:"TimeStr,omitempty" xml:"TimeStr,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) SetComments(v string) *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory {
	s.Comments = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) SetComplainResult(v int32) *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory {
	s.ComplainResult = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) SetOldScore(v int32) *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory {
	s.OldScore = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) SetOperator(v int64) *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory {
	s.Operator = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) SetOperatorName(v string) *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory {
	s.OperatorName = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) SetReviewManagerType(v string) *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory {
	s.ReviewManagerType = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) SetReviewResult(v int32) *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory {
	s.ReviewResult = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) SetReviewRightRule(v *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRule) *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory {
	s.ReviewRightRule = v
	return s
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) SetScore(v int32) *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory {
	s.Score = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) SetTime(v int64) *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory {
	s.Time = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) SetTimeStr(v string) *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory {
	s.TimeStr = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) SetType(v int32) *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory {
	s.Type = &v
	return s
}

type GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRule struct {
	ReviewRightRule []*GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule `json:"ReviewRightRule,omitempty" xml:"ReviewRightRule,omitempty" type:"Repeated"`
}

func (s GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRule) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRule) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRule) SetReviewRightRule(v []*GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule) *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRule {
	s.ReviewRightRule = v
	return s
}

type GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule struct {
	Rid      *int64  `json:"rid,omitempty" xml:"rid,omitempty"`
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
}

func (s GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule) SetRid(v int64) *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule {
	s.Rid = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule) SetRuleName(v string) *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule {
	s.RuleName = &v
	return s
}

type GetResultToReviewResponseBodyDataReviewTypeIdList struct {
	ReviewTypeIdList []*GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdList `json:"ReviewTypeIdList,omitempty" xml:"ReviewTypeIdList,omitempty" type:"Repeated"`
}

func (s GetResultToReviewResponseBodyDataReviewTypeIdList) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataReviewTypeIdList) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataReviewTypeIdList) SetReviewTypeIdList(v []*GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdList) *GetResultToReviewResponseBodyDataReviewTypeIdList {
	s.ReviewTypeIdList = v
	return s
}

type GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdList struct {
	ReviewKeyIdList *GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdListReviewKeyIdList `json:"ReviewKeyIdList,omitempty" xml:"ReviewKeyIdList,omitempty" type:"Struct"`
	ReviewTypeId    *int64                                                                            `json:"ReviewTypeId,omitempty" xml:"ReviewTypeId,omitempty"`
}

func (s GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdList) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdList) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdList) SetReviewKeyIdList(v *GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdListReviewKeyIdList) *GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdList {
	s.ReviewKeyIdList = v
	return s
}

func (s *GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdList) SetReviewTypeId(v int64) *GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdList {
	s.ReviewTypeId = &v
	return s
}

type GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdListReviewKeyIdList struct {
	ReviewKeyIdList []*int64 `json:"ReviewKeyIdList,omitempty" xml:"ReviewKeyIdList,omitempty" type:"Repeated"`
}

func (s GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdListReviewKeyIdList) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdListReviewKeyIdList) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdListReviewKeyIdList) SetReviewKeyIdList(v []*int64) *GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdListReviewKeyIdList {
	s.ReviewKeyIdList = v
	return s
}

type GetResultToReviewResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResultToReviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResultToReviewResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponse) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponse) SetHeaders(v map[string]*string) *GetResultToReviewResponse {
	s.Headers = v
	return s
}

func (s *GetResultToReviewResponse) SetStatusCode(v int32) *GetResultToReviewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResultToReviewResponse) SetBody(v *GetResultToReviewResponseBody) *GetResultToReviewResponse {
	s.Body = v
	return s
}

type GetRuleRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"ruleIds":"123"}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRuleRequest) GoString() string {
	return s.String()
}

func (s *GetRuleRequest) SetBaseMeAgentId(v int64) *GetRuleRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetRuleRequest) SetJsonStr(v string) *GetRuleRequest {
	s.JsonStr = &v
	return s
}

type GetRuleResponseBody struct {
	// example:
	//
	// 200
	Code *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F190ADE9-619A-447D-84E3-7E241A5C428E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBody) SetCode(v string) *GetRuleResponseBody {
	s.Code = &v
	return s
}

func (s *GetRuleResponseBody) SetData(v *GetRuleResponseBodyData) *GetRuleResponseBody {
	s.Data = v
	return s
}

func (s *GetRuleResponseBody) SetMessage(v string) *GetRuleResponseBody {
	s.Message = &v
	return s
}

func (s *GetRuleResponseBody) SetRequestId(v string) *GetRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRuleResponseBody) SetSuccess(v bool) *GetRuleResponseBody {
	s.Success = &v
	return s
}

type GetRuleResponseBodyData struct {
	Rules *GetRuleResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
}

func (s GetRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBodyData) SetRules(v *GetRuleResponseBodyDataRules) *GetRuleResponseBodyData {
	s.Rules = v
	return s
}

type GetRuleResponseBodyDataRules struct {
	RuleInfo []*GetRuleResponseBodyDataRulesRuleInfo `json:"RuleInfo,omitempty" xml:"RuleInfo,omitempty" type:"Repeated"`
}

func (s GetRuleResponseBodyDataRules) String() string {
	return tea.Prettify(s)
}

func (s GetRuleResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBodyDataRules) SetRuleInfo(v []*GetRuleResponseBodyDataRulesRuleInfo) *GetRuleResponseBodyDataRules {
	s.RuleInfo = v
	return s
}

type GetRuleResponseBodyDataRulesRuleInfo struct {
	// example:
	//
	// 1
	AutoReview               *int32                                                        `json:"AutoReview,omitempty" xml:"AutoReview,omitempty"`
	BusinessCategoryNameList *GetRuleResponseBodyDataRulesRuleInfoBusinessCategoryNameList `json:"BusinessCategoryNameList,omitempty" xml:"BusinessCategoryNameList,omitempty" type:"Struct"`
	Comments                 *string                                                       `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// example:
	//
	// 123
	CreateEmpid *string `json:"CreateEmpid,omitempty" xml:"CreateEmpid,omitempty"`
	// example:
	//
	// 2016-08-05 10:37:10
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2016-08-05 10:37:10
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 0
	IsDelete *int32 `json:"IsDelete,omitempty" xml:"IsDelete,omitempty"`
	// example:
	//
	// 1
	IsOnline *int32 `json:"IsOnline,omitempty" xml:"IsOnline,omitempty"`
	// example:
	//
	// 123
	LastUpdateEmpid *string `json:"LastUpdateEmpid,omitempty" xml:"LastUpdateEmpid,omitempty"`
	// example:
	//
	// 2019-10-28 14:23:28
	LastUpdateTime *string `json:"LastUpdateTime,omitempty" xml:"LastUpdateTime,omitempty"`
	// example:
	//
	// demo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 4
	Rid *string `json:"Rid,omitempty" xml:"Rid,omitempty"`
	// example:
	//
	// a && b
	RuleLambda *string `json:"RuleLambda,omitempty" xml:"RuleLambda,omitempty"`
	// example:
	//
	// 1
	RuleScoreType *int32 `json:"RuleScoreType,omitempty" xml:"RuleScoreType,omitempty"`
	// example:
	//
	// 123
	ScoreId   *int32  `json:"ScoreId,omitempty" xml:"ScoreId,omitempty"`
	ScoreName *string `json:"ScoreName,omitempty" xml:"ScoreName,omitempty"`
	// example:
	//
	// 22
	ScoreSubId   *int32  `json:"ScoreSubId,omitempty" xml:"ScoreSubId,omitempty"`
	ScoreSubName *string `json:"ScoreSubName,omitempty" xml:"ScoreSubName,omitempty"`
	// example:
	//
	// 2016-08-05 10:37:10
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 1
	Weight *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s GetRuleResponseBodyDataRulesRuleInfo) String() string {
	return tea.Prettify(s)
}

func (s GetRuleResponseBodyDataRulesRuleInfo) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetAutoReview(v int32) *GetRuleResponseBodyDataRulesRuleInfo {
	s.AutoReview = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetBusinessCategoryNameList(v *GetRuleResponseBodyDataRulesRuleInfoBusinessCategoryNameList) *GetRuleResponseBodyDataRulesRuleInfo {
	s.BusinessCategoryNameList = v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetComments(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.Comments = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetCreateEmpid(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.CreateEmpid = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetCreateTime(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.CreateTime = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetEndTime(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.EndTime = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetIsDelete(v int32) *GetRuleResponseBodyDataRulesRuleInfo {
	s.IsDelete = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetIsOnline(v int32) *GetRuleResponseBodyDataRulesRuleInfo {
	s.IsOnline = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetLastUpdateEmpid(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.LastUpdateEmpid = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetLastUpdateTime(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.LastUpdateTime = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetName(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.Name = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetRid(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.Rid = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetRuleLambda(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.RuleLambda = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetRuleScoreType(v int32) *GetRuleResponseBodyDataRulesRuleInfo {
	s.RuleScoreType = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetScoreId(v int32) *GetRuleResponseBodyDataRulesRuleInfo {
	s.ScoreId = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetScoreName(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.ScoreName = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetScoreSubId(v int32) *GetRuleResponseBodyDataRulesRuleInfo {
	s.ScoreSubId = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetScoreSubName(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.ScoreSubName = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetStartTime(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.StartTime = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetStatus(v int32) *GetRuleResponseBodyDataRulesRuleInfo {
	s.Status = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetType(v int32) *GetRuleResponseBodyDataRulesRuleInfo {
	s.Type = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetWeight(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.Weight = &v
	return s
}

type GetRuleResponseBodyDataRulesRuleInfoBusinessCategoryNameList struct {
	BusinessCategoryNameList []*string `json:"BusinessCategoryNameList,omitempty" xml:"BusinessCategoryNameList,omitempty" type:"Repeated"`
}

func (s GetRuleResponseBodyDataRulesRuleInfoBusinessCategoryNameList) String() string {
	return tea.Prettify(s)
}

func (s GetRuleResponseBodyDataRulesRuleInfoBusinessCategoryNameList) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBodyDataRulesRuleInfoBusinessCategoryNameList) SetBusinessCategoryNameList(v []*string) *GetRuleResponseBodyDataRulesRuleInfoBusinessCategoryNameList {
	s.BusinessCategoryNameList = v
	return s
}

type GetRuleResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRuleResponse) GoString() string {
	return s.String()
}

func (s *GetRuleResponse) SetHeaders(v map[string]*string) *GetRuleResponse {
	s.Headers = v
	return s
}

func (s *GetRuleResponse) SetStatusCode(v int32) *GetRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRuleResponse) SetBody(v *GetRuleResponseBody) *GetRuleResponse {
	s.Body = v
	return s
}

type GetRuleByIdRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 53
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s GetRuleByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRuleByIdRequest) GoString() string {
	return s.String()
}

func (s *GetRuleByIdRequest) SetBaseMeAgentId(v int64) *GetRuleByIdRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetRuleByIdRequest) SetRuleId(v int64) *GetRuleByIdRequest {
	s.RuleId = &v
	return s
}

type GetRuleByIdResponseBody struct {
	// example:
	//
	// 200
	Code *string    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *RulesInfo `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages []*string `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// example:
	//
	// 3CEA0495-341B-4482-9AD9-8191EF4***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRuleByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRuleByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetRuleByIdResponseBody) SetCode(v string) *GetRuleByIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetRuleByIdResponseBody) SetData(v *RulesInfo) *GetRuleByIdResponseBody {
	s.Data = v
	return s
}

func (s *GetRuleByIdResponseBody) SetHttpStatusCode(v int32) *GetRuleByIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetRuleByIdResponseBody) SetMessage(v string) *GetRuleByIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetRuleByIdResponseBody) SetMessages(v []*string) *GetRuleByIdResponseBody {
	s.Messages = v
	return s
}

func (s *GetRuleByIdResponseBody) SetRequestId(v string) *GetRuleByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRuleByIdResponseBody) SetSuccess(v bool) *GetRuleByIdResponseBody {
	s.Success = &v
	return s
}

type GetRuleByIdResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRuleByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRuleByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRuleByIdResponse) GoString() string {
	return s.String()
}

func (s *GetRuleByIdResponse) SetHeaders(v map[string]*string) *GetRuleByIdResponse {
	s.Headers = v
	return s
}

func (s *GetRuleByIdResponse) SetStatusCode(v int32) *GetRuleByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRuleByIdResponse) SetBody(v *GetRuleByIdResponseBody) *GetRuleByIdResponse {
	s.Body = v
	return s
}

type GetRuleCategoryRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ""
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetRuleCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRuleCategoryRequest) GoString() string {
	return s.String()
}

func (s *GetRuleCategoryRequest) SetBaseMeAgentId(v int64) *GetRuleCategoryRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetRuleCategoryRequest) SetJsonStr(v string) *GetRuleCategoryRequest {
	s.JsonStr = &v
	return s
}

type GetRuleCategoryResponseBody struct {
	// example:
	//
	// 200
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetRuleCategoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F190ADE9-619A-447D-84E3-7E241A5C428E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRuleCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRuleCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetRuleCategoryResponseBody) SetCode(v string) *GetRuleCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *GetRuleCategoryResponseBody) SetData(v *GetRuleCategoryResponseBodyData) *GetRuleCategoryResponseBody {
	s.Data = v
	return s
}

func (s *GetRuleCategoryResponseBody) SetMessage(v string) *GetRuleCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *GetRuleCategoryResponseBody) SetRequestId(v string) *GetRuleCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRuleCategoryResponseBody) SetSuccess(v bool) *GetRuleCategoryResponseBody {
	s.Success = &v
	return s
}

type GetRuleCategoryResponseBodyData struct {
	RuleCountInfo []*GetRuleCategoryResponseBodyDataRuleCountInfo `json:"RuleCountInfo,omitempty" xml:"RuleCountInfo,omitempty" type:"Repeated"`
}

func (s GetRuleCategoryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRuleCategoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRuleCategoryResponseBodyData) SetRuleCountInfo(v []*GetRuleCategoryResponseBodyDataRuleCountInfo) *GetRuleCategoryResponseBodyData {
	s.RuleCountInfo = v
	return s
}

type GetRuleCategoryResponseBodyDataRuleCountInfo struct {
	// example:
	//
	// false
	Select *bool `json:"Select,omitempty" xml:"Select,omitempty"`
	// example:
	//
	// 22
	Type     *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s GetRuleCategoryResponseBodyDataRuleCountInfo) String() string {
	return tea.Prettify(s)
}

func (s GetRuleCategoryResponseBodyDataRuleCountInfo) GoString() string {
	return s.String()
}

func (s *GetRuleCategoryResponseBodyDataRuleCountInfo) SetSelect(v bool) *GetRuleCategoryResponseBodyDataRuleCountInfo {
	s.Select = &v
	return s
}

func (s *GetRuleCategoryResponseBodyDataRuleCountInfo) SetType(v int32) *GetRuleCategoryResponseBodyDataRuleCountInfo {
	s.Type = &v
	return s
}

func (s *GetRuleCategoryResponseBodyDataRuleCountInfo) SetTypeName(v string) *GetRuleCategoryResponseBodyDataRuleCountInfo {
	s.TypeName = &v
	return s
}

type GetRuleCategoryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRuleCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRuleCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRuleCategoryResponse) GoString() string {
	return s.String()
}

func (s *GetRuleCategoryResponse) SetHeaders(v map[string]*string) *GetRuleCategoryResponse {
	s.Headers = v
	return s
}

func (s *GetRuleCategoryResponse) SetStatusCode(v int32) *GetRuleCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRuleCategoryResponse) SetBody(v *GetRuleCategoryResponseBody) *GetRuleCategoryResponse {
	s.Body = v
	return s
}

type GetRuleDetailRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"ruleIds":"123"}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetRuleDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailRequest) GoString() string {
	return s.String()
}

func (s *GetRuleDetailRequest) SetBaseMeAgentId(v int64) *GetRuleDetailRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetRuleDetailRequest) SetJsonStr(v string) *GetRuleDetailRequest {
	s.JsonStr = &v
	return s
}

type GetRuleDetailResponseBody struct {
	// example:
	//
	// 200
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetRuleDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4987D326-83D9-4A42-B9A5-0B27F9B40539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRuleDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBody) SetCode(v string) *GetRuleDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetRuleDetailResponseBody) SetData(v *GetRuleDetailResponseBodyData) *GetRuleDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetRuleDetailResponseBody) SetMessage(v string) *GetRuleDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetRuleDetailResponseBody) SetRequestId(v string) *GetRuleDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRuleDetailResponseBody) SetSuccess(v bool) *GetRuleDetailResponseBody {
	s.Success = &v
	return s
}

type GetRuleDetailResponseBodyData struct {
	Conditions *GetRuleDetailResponseBodyDataConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Struct"`
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Rules    *GetRuleDetailResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
}

func (s GetRuleDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyData) SetConditions(v *GetRuleDetailResponseBodyDataConditions) *GetRuleDetailResponseBodyData {
	s.Conditions = v
	return s
}

func (s *GetRuleDetailResponseBodyData) SetCount(v int32) *GetRuleDetailResponseBodyData {
	s.Count = &v
	return s
}

func (s *GetRuleDetailResponseBodyData) SetPageNumber(v int32) *GetRuleDetailResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *GetRuleDetailResponseBodyData) SetPageSize(v int32) *GetRuleDetailResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetRuleDetailResponseBodyData) SetRules(v *GetRuleDetailResponseBodyDataRules) *GetRuleDetailResponseBodyData {
	s.Rules = v
	return s
}

type GetRuleDetailResponseBodyDataConditions struct {
	ConditionBasicInfo []*GetRuleDetailResponseBodyDataConditionsConditionBasicInfo `json:"ConditionBasicInfo,omitempty" xml:"ConditionBasicInfo,omitempty" type:"Repeated"`
}

func (s GetRuleDetailResponseBodyDataConditions) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditions) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditions) SetConditionBasicInfo(v []*GetRuleDetailResponseBodyDataConditionsConditionBasicInfo) *GetRuleDetailResponseBodyDataConditions {
	s.ConditionBasicInfo = v
	return s
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfo struct {
	CheckRange *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange `json:"CheckRange,omitempty" xml:"CheckRange,omitempty" type:"Struct"`
	// example:
	//
	// 7
	ConditionInfoCid *string `json:"ConditionInfoCid,omitempty" xml:"ConditionInfoCid,omitempty"`
	// example:
	//
	// 7
	OperLambda *string                                                             `json:"OperLambda,omitempty" xml:"OperLambda,omitempty"`
	Operators  *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperators `json:"Operators,omitempty" xml:"Operators,omitempty" type:"Struct"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfo) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfo) SetCheckRange(v *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfo {
	s.CheckRange = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfo) SetConditionInfoCid(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfo {
	s.ConditionInfoCid = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfo) SetOperLambda(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfo {
	s.OperLambda = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfo) SetOperators(v *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperators) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfo {
	s.Operators = v
	return s
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange struct {
	// example:
	//
	// true
	Absolute *bool                                                                      `json:"Absolute,omitempty" xml:"Absolute,omitempty"`
	Anchor   *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor `json:"Anchor,omitempty" xml:"Anchor,omitempty" type:"Struct"`
	Range    *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeRange  `json:"Range,omitempty" xml:"Range,omitempty" type:"Struct"`
	Role     *string                                                                    `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange) SetAbsolute(v bool) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange {
	s.Absolute = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange) SetAnchor(v *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange {
	s.Anchor = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange) SetRange(v *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeRange) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange {
	s.Range = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange) SetRole(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange {
	s.Role = &v
	return s
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor struct {
	// example:
	//
	// 7
	AnchorCid *string `json:"AnchorCid,omitempty" xml:"AnchorCid,omitempty"`
	// example:
	//
	// 1
	HitTime *int32 `json:"HitTime,omitempty" xml:"HitTime,omitempty"`
	// example:
	//
	// AFTER
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor) SetAnchorCid(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor {
	s.AnchorCid = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor) SetHitTime(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor {
	s.HitTime = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor) SetLocation(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor {
	s.Location = &v
	return s
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeRange struct {
	// example:
	//
	// 1
	From *int32 `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// 10
	To *int32 `json:"To,omitempty" xml:"To,omitempty"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeRange) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeRange) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeRange) SetFrom(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeRange {
	s.From = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeRange) SetTo(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeRange {
	s.To = &v
	return s
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperators struct {
	OperatorBasicInfo []*GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo `json:"OperatorBasicInfo,omitempty" xml:"OperatorBasicInfo,omitempty" type:"Repeated"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperators) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperators) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperators) SetOperatorBasicInfo(v []*GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperators {
	s.OperatorBasicInfo = v
	return s
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo struct {
	// example:
	//
	// 8
	Oid *string `json:"Oid,omitempty" xml:"Oid,omitempty"`
	// example:
	//
	// operator demo
	OperName *string                                                                                   `json:"OperName,omitempty" xml:"OperName,omitempty"`
	Param    *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam `json:"Param,omitempty" xml:"Param,omitempty" type:"Struct"`
	// example:
	//
	// REGULAR_EXPRESSION
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo) SetOid(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo {
	s.Oid = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo) SetOperName(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo {
	s.OperName = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo) SetParam(v *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo {
	s.Param = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo) SetType(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo {
	s.Type = &v
	return s
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam struct {
	AntModelInfo *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamAntModelInfo `json:"AntModelInfo,omitempty" xml:"AntModelInfo,omitempty" type:"Struct"`
	// example:
	//
	// true
	Average *bool `json:"Average,omitempty" xml:"Average,omitempty"`
	// example:
	//
	// DIALOGUE
	BeginType *string `json:"BeginType,omitempty" xml:"BeginType,omitempty"`
	// example:
	//
	// 1
	CheckType *int32 `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// example:
	//
	// gt
	CompareOperator *string `json:"CompareOperator,omitempty" xml:"CompareOperator,omitempty"`
	// example:
	//
	// true
	ContextChatMatch *bool `json:"ContextChatMatch,omitempty" xml:"ContextChatMatch,omitempty"`
	// example:
	//
	// 1000
	DelayTime *int32 `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	// example:
	//
	// true
	DifferentRole *bool                                                                                             `json:"DifferentRole,omitempty" xml:"DifferentRole,omitempty"`
	Excludes      *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamExcludes `json:"Excludes,omitempty" xml:"Excludes,omitempty" type:"Struct"`
	// example:
	//
	// 3
	From *int32 `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// true
	FromEnd *bool `json:"FromEnd,omitempty" xml:"FromEnd,omitempty"`
	// example:
	//
	// 1
	HitTime *int32 `json:"HitTime,omitempty" xml:"HitTime,omitempty"`
	// example:
	//
	// true
	InSentence *bool `json:"InSentence,omitempty" xml:"InSentence,omitempty"`
	// example:
	//
	// 5000
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// true
	KeywordExtension *bool `json:"KeywordExtension,omitempty" xml:"KeywordExtension,omitempty"`
	// example:
	//
	// 3
	KeywordMatchSize *int32 `json:"KeywordMatchSize,omitempty" xml:"KeywordMatchSize,omitempty"`
	// example:
	//
	// 8
	MaxEmotionChangeValue *int32 `json:"MaxEmotionChangeValue,omitempty" xml:"MaxEmotionChangeValue,omitempty"`
	// example:
	//
	// 4
	MinWordSize  *int32                                                                                                `json:"MinWordSize,omitempty" xml:"MinWordSize,omitempty"`
	NotRegex     *string                                                                                               `json:"NotRegex,omitempty" xml:"NotRegex,omitempty"`
	OperKeyWords *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamOperKeyWords `json:"OperKeyWords,omitempty" xml:"OperKeyWords,omitempty" type:"Struct"`
	// example:
	//
	// xxx
	Phrase     *string                                                                                             `json:"Phrase,omitempty" xml:"Phrase,omitempty"`
	Pvalues    *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamPvalues    `json:"Pvalues,omitempty" xml:"Pvalues,omitempty" type:"Struct"`
	References *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamReferences `json:"References,omitempty" xml:"References,omitempty" type:"Struct"`
	Regex      *string                                                                                             `json:"Regex,omitempty" xml:"Regex,omitempty"`
	// example:
	//
	// 80
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// 90
	SimilarityThreshold *float32                                                                                                    `json:"Similarity_threshold,omitempty" xml:"Similarity_threshold,omitempty"`
	SimilarlySentences  *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamSimilarlySentences `json:"SimilarlySentences,omitempty" xml:"SimilarlySentences,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Target     *int32  `json:"Target,omitempty" xml:"Target,omitempty"`
	TargetRole *string `json:"TargetRole,omitempty" xml:"TargetRole,omitempty"`
	// example:
	//
	// 4
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// example:
	//
	// 280
	VelocityInMint *int32 `json:"VelocityInMint,omitempty" xml:"VelocityInMint,omitempty"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetAntModelInfo(v *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamAntModelInfo) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.AntModelInfo = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetAverage(v bool) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.Average = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetBeginType(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.BeginType = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetCheckType(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.CheckType = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetCompareOperator(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.CompareOperator = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetContextChatMatch(v bool) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.ContextChatMatch = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetDelayTime(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.DelayTime = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetDifferentRole(v bool) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.DifferentRole = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetExcludes(v *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamExcludes) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.Excludes = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetFrom(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.From = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetFromEnd(v bool) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.FromEnd = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetHitTime(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.HitTime = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetInSentence(v bool) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.InSentence = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetInterval(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.Interval = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetKeywordExtension(v bool) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.KeywordExtension = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetKeywordMatchSize(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.KeywordMatchSize = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetMaxEmotionChangeValue(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.MaxEmotionChangeValue = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetMinWordSize(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.MinWordSize = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetNotRegex(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.NotRegex = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetOperKeyWords(v *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamOperKeyWords) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.OperKeyWords = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetPhrase(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.Phrase = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetPvalues(v *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamPvalues) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.Pvalues = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetReferences(v *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamReferences) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.References = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetRegex(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.Regex = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetScore(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.Score = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetSimilarityThreshold(v float32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.SimilarityThreshold = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetSimilarlySentences(v *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamSimilarlySentences) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.SimilarlySentences = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetTarget(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.Target = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetTargetRole(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.TargetRole = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetThreshold(v float32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.Threshold = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetVelocityInMint(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.VelocityInMint = &v
	return s
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamAntModelInfo struct {
	AntModelInfo []*string `json:"AntModelInfo,omitempty" xml:"AntModelInfo,omitempty" type:"Repeated"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamAntModelInfo) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamAntModelInfo) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamAntModelInfo) SetAntModelInfo(v []*string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamAntModelInfo {
	s.AntModelInfo = v
	return s
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamExcludes struct {
	Excludes []*string `json:"Excludes,omitempty" xml:"Excludes,omitempty" type:"Repeated"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamExcludes) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamExcludes) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamExcludes) SetExcludes(v []*string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamExcludes {
	s.Excludes = v
	return s
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamOperKeyWords struct {
	OperKeyWord []*string `json:"OperKeyWord,omitempty" xml:"OperKeyWord,omitempty" type:"Repeated"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamOperKeyWords) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamOperKeyWords) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamOperKeyWords) SetOperKeyWord(v []*string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamOperKeyWords {
	s.OperKeyWord = v
	return s
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamPvalues struct {
	Pvalues []*string `json:"Pvalues,omitempty" xml:"Pvalues,omitempty" type:"Repeated"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamPvalues) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamPvalues) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamPvalues) SetPvalues(v []*string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamPvalues {
	s.Pvalues = v
	return s
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamReferences struct {
	Reference []*string `json:"Reference,omitempty" xml:"Reference,omitempty" type:"Repeated"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamReferences) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamReferences) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamReferences) SetReference(v []*string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamReferences {
	s.Reference = v
	return s
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamSimilarlySentences struct {
	SimilarlySentence []*string `json:"SimilarlySentence,omitempty" xml:"SimilarlySentence,omitempty" type:"Repeated"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamSimilarlySentences) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamSimilarlySentences) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamSimilarlySentences) SetSimilarlySentence(v []*string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamSimilarlySentences {
	s.SimilarlySentence = v
	return s
}

type GetRuleDetailResponseBodyDataRules struct {
	RuleBasicInfo []*GetRuleDetailResponseBodyDataRulesRuleBasicInfo `json:"RuleBasicInfo,omitempty" xml:"RuleBasicInfo,omitempty" type:"Repeated"`
}

func (s GetRuleDetailResponseBodyDataRules) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataRules) SetRuleBasicInfo(v []*GetRuleDetailResponseBodyDataRulesRuleBasicInfo) *GetRuleDetailResponseBodyDataRules {
	s.RuleBasicInfo = v
	return s
}

type GetRuleDetailResponseBodyDataRulesRuleBasicInfo struct {
	BusinessCategories *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategories `json:"BusinessCategories,omitempty" xml:"BusinessCategories,omitempty" type:"Struct"`
	// example:
	//
	// 4
	Rid *string `json:"Rid,omitempty" xml:"Rid,omitempty"`
	// example:
	//
	// 7&&!8
	RuleLambda *string                                                  `json:"RuleLambda,omitempty" xml:"RuleLambda,omitempty"`
	Triggers   *GetRuleDetailResponseBodyDataRulesRuleBasicInfoTriggers `json:"Triggers,omitempty" xml:"Triggers,omitempty" type:"Struct"`
}

func (s GetRuleDetailResponseBodyDataRulesRuleBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataRulesRuleBasicInfo) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfo) SetBusinessCategories(v *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategories) *GetRuleDetailResponseBodyDataRulesRuleBasicInfo {
	s.BusinessCategories = v
	return s
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfo) SetRid(v string) *GetRuleDetailResponseBodyDataRulesRuleBasicInfo {
	s.Rid = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfo) SetRuleLambda(v string) *GetRuleDetailResponseBodyDataRulesRuleBasicInfo {
	s.RuleLambda = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfo) SetTriggers(v *GetRuleDetailResponseBodyDataRulesRuleBasicInfoTriggers) *GetRuleDetailResponseBodyDataRulesRuleBasicInfo {
	s.Triggers = v
	return s
}

type GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategories struct {
	BusinessCategoryBasicInfo []*GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo `json:"BusinessCategoryBasicInfo,omitempty" xml:"BusinessCategoryBasicInfo,omitempty" type:"Repeated"`
}

func (s GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategories) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategories) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategories) SetBusinessCategoryBasicInfo(v []*GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo) *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategories {
	s.BusinessCategoryBasicInfo = v
	return s
}

type GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo struct {
	// example:
	//
	// 264971810
	Bid          *int32  `json:"Bid,omitempty" xml:"Bid,omitempty"`
	BusinessName *string `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	// example:
	//
	// 1
	ServiceType *int32 `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo) SetBid(v int32) *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo {
	s.Bid = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo) SetBusinessName(v string) *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo {
	s.BusinessName = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo) SetServiceType(v int32) *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo {
	s.ServiceType = &v
	return s
}

type GetRuleDetailResponseBodyDataRulesRuleBasicInfoTriggers struct {
	Trigger []*string `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Repeated"`
}

func (s GetRuleDetailResponseBodyDataRulesRuleBasicInfoTriggers) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataRulesRuleBasicInfoTriggers) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfoTriggers) SetTrigger(v []*string) *GetRuleDetailResponseBodyDataRulesRuleBasicInfoTriggers {
	s.Trigger = v
	return s
}

type GetRuleDetailResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRuleDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRuleDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponse) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponse) SetHeaders(v map[string]*string) *GetRuleDetailResponse {
	s.Headers = v
	return s
}

func (s *GetRuleDetailResponse) SetStatusCode(v int32) *GetRuleDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRuleDetailResponse) SetBody(v *GetRuleDetailResponseBody) *GetRuleDetailResponse {
	s.Body = v
	return s
}

type GetRuleV4Request struct {
	// This parameter is required.
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s GetRuleV4Request) String() string {
	return tea.Prettify(s)
}

func (s GetRuleV4Request) GoString() string {
	return s.String()
}

func (s *GetRuleV4Request) SetRuleId(v int64) *GetRuleV4Request {
	s.RuleId = &v
	return s
}

type GetRuleV4ResponseBody struct {
	// example:
	//
	// 200
	Code *string    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *RulesInfo `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages []*string `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRuleV4ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRuleV4ResponseBody) GoString() string {
	return s.String()
}

func (s *GetRuleV4ResponseBody) SetCode(v string) *GetRuleV4ResponseBody {
	s.Code = &v
	return s
}

func (s *GetRuleV4ResponseBody) SetData(v *RulesInfo) *GetRuleV4ResponseBody {
	s.Data = v
	return s
}

func (s *GetRuleV4ResponseBody) SetHttpStatusCode(v int32) *GetRuleV4ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetRuleV4ResponseBody) SetMessage(v string) *GetRuleV4ResponseBody {
	s.Message = &v
	return s
}

func (s *GetRuleV4ResponseBody) SetMessages(v []*string) *GetRuleV4ResponseBody {
	s.Messages = v
	return s
}

func (s *GetRuleV4ResponseBody) SetRequestId(v string) *GetRuleV4ResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRuleV4ResponseBody) SetSuccess(v bool) *GetRuleV4ResponseBody {
	s.Success = &v
	return s
}

type GetRuleV4Response struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRuleV4ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRuleV4Response) String() string {
	return tea.Prettify(s)
}

func (s GetRuleV4Response) GoString() string {
	return s.String()
}

func (s *GetRuleV4Response) SetHeaders(v map[string]*string) *GetRuleV4Response {
	s.Headers = v
	return s
}

func (s *GetRuleV4Response) SetStatusCode(v int32) *GetRuleV4Response {
	s.StatusCode = &v
	return s
}

func (s *GetRuleV4Response) SetBody(v *GetRuleV4ResponseBody) *GetRuleV4Response {
	s.Body = v
	return s
}

type GetRulesCountListRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64  `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	BusinessName  *string `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	// example:
	//
	// 1
	BusinessRange *int32  `json:"BusinessRange,omitempty" xml:"BusinessRange,omitempty"`
	CategoryName  *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// example:
	//
	// true
	CountTotal *bool `json:"CountTotal,omitempty" xml:"CountTotal,omitempty"`
	// example:
	//
	// 63
	CreateEmpid *string `json:"CreateEmpid,omitempty" xml:"CreateEmpid,omitempty"`
	// example:
	//
	// 63
	CreateUserId *int64 `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 2022-10-08 23:59:59
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 63
	LastUpdateEmpid *string `json:"LastUpdateEmpid,omitempty" xml:"LastUpdateEmpid,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize     *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequireInfos []*string `json:"RequireInfos,omitempty" xml:"RequireInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 123
	Rid *int64 `json:"Rid,omitempty" xml:"Rid,omitempty"`
	// example:
	//
	// 123
	RuleIdOrRuleName *string `json:"RuleIdOrRuleName,omitempty" xml:"RuleIdOrRuleName,omitempty"`
	// example:
	//
	// 1
	RuleScoreSingleType *int32 `json:"RuleScoreSingleType,omitempty" xml:"RuleScoreSingleType,omitempty"`
	// example:
	//
	// 1
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// example:
	//
	// 123
	SchemeId *int64 `json:"SchemeId,omitempty" xml:"SchemeId,omitempty"`
	// example:
	//
	// 0
	SourceType *int32 `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// 2022-10-07 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1
	Type     *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
	// example:
	//
	// 2022-10-08 23:59:59
	UpdateEndTime *string `json:"UpdateEndTime,omitempty" xml:"UpdateEndTime,omitempty"`
	// example:
	//
	// 2022-10-07 00:00:00
	UpdateStartTime *string `json:"UpdateStartTime,omitempty" xml:"UpdateStartTime,omitempty"`
	// example:
	//
	// 63
	UpdateUserId *int64 `json:"UpdateUserId,omitempty" xml:"UpdateUserId,omitempty"`
}

func (s GetRulesCountListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRulesCountListRequest) GoString() string {
	return s.String()
}

func (s *GetRulesCountListRequest) SetBaseMeAgentId(v int64) *GetRulesCountListRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetRulesCountListRequest) SetBusinessName(v string) *GetRulesCountListRequest {
	s.BusinessName = &v
	return s
}

func (s *GetRulesCountListRequest) SetBusinessRange(v int32) *GetRulesCountListRequest {
	s.BusinessRange = &v
	return s
}

func (s *GetRulesCountListRequest) SetCategoryName(v string) *GetRulesCountListRequest {
	s.CategoryName = &v
	return s
}

func (s *GetRulesCountListRequest) SetCountTotal(v bool) *GetRulesCountListRequest {
	s.CountTotal = &v
	return s
}

func (s *GetRulesCountListRequest) SetCreateEmpid(v string) *GetRulesCountListRequest {
	s.CreateEmpid = &v
	return s
}

func (s *GetRulesCountListRequest) SetCreateUserId(v int64) *GetRulesCountListRequest {
	s.CreateUserId = &v
	return s
}

func (s *GetRulesCountListRequest) SetCurrentPage(v int32) *GetRulesCountListRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetRulesCountListRequest) SetEndTime(v string) *GetRulesCountListRequest {
	s.EndTime = &v
	return s
}

func (s *GetRulesCountListRequest) SetLastUpdateEmpid(v string) *GetRulesCountListRequest {
	s.LastUpdateEmpid = &v
	return s
}

func (s *GetRulesCountListRequest) SetPageNumber(v int32) *GetRulesCountListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetRulesCountListRequest) SetPageSize(v int32) *GetRulesCountListRequest {
	s.PageSize = &v
	return s
}

func (s *GetRulesCountListRequest) SetRequireInfos(v []*string) *GetRulesCountListRequest {
	s.RequireInfos = v
	return s
}

func (s *GetRulesCountListRequest) SetRid(v int64) *GetRulesCountListRequest {
	s.Rid = &v
	return s
}

func (s *GetRulesCountListRequest) SetRuleIdOrRuleName(v string) *GetRulesCountListRequest {
	s.RuleIdOrRuleName = &v
	return s
}

func (s *GetRulesCountListRequest) SetRuleScoreSingleType(v int32) *GetRulesCountListRequest {
	s.RuleScoreSingleType = &v
	return s
}

func (s *GetRulesCountListRequest) SetRuleType(v int32) *GetRulesCountListRequest {
	s.RuleType = &v
	return s
}

func (s *GetRulesCountListRequest) SetSchemeId(v int64) *GetRulesCountListRequest {
	s.SchemeId = &v
	return s
}

func (s *GetRulesCountListRequest) SetSourceType(v int32) *GetRulesCountListRequest {
	s.SourceType = &v
	return s
}

func (s *GetRulesCountListRequest) SetStartTime(v string) *GetRulesCountListRequest {
	s.StartTime = &v
	return s
}

func (s *GetRulesCountListRequest) SetStatus(v int32) *GetRulesCountListRequest {
	s.Status = &v
	return s
}

func (s *GetRulesCountListRequest) SetType(v int32) *GetRulesCountListRequest {
	s.Type = &v
	return s
}

func (s *GetRulesCountListRequest) SetTypeName(v string) *GetRulesCountListRequest {
	s.TypeName = &v
	return s
}

func (s *GetRulesCountListRequest) SetUpdateEndTime(v string) *GetRulesCountListRequest {
	s.UpdateEndTime = &v
	return s
}

func (s *GetRulesCountListRequest) SetUpdateStartTime(v string) *GetRulesCountListRequest {
	s.UpdateStartTime = &v
	return s
}

func (s *GetRulesCountListRequest) SetUpdateUserId(v int64) *GetRulesCountListRequest {
	s.UpdateUserId = &v
	return s
}

type GetRulesCountListResponseBody struct {
	BusinessType *int32 `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32                             `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Data        *GetRulesCountListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *GetRulesCountListResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 9987D326-83D9-4A42-B9A5-0B27F9B40539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 7
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetRulesCountListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRulesCountListResponseBody) GoString() string {
	return s.String()
}

func (s *GetRulesCountListResponseBody) SetBusinessType(v int32) *GetRulesCountListResponseBody {
	s.BusinessType = &v
	return s
}

func (s *GetRulesCountListResponseBody) SetCode(v string) *GetRulesCountListResponseBody {
	s.Code = &v
	return s
}

func (s *GetRulesCountListResponseBody) SetCount(v int32) *GetRulesCountListResponseBody {
	s.Count = &v
	return s
}

func (s *GetRulesCountListResponseBody) SetCurrentPage(v int32) *GetRulesCountListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *GetRulesCountListResponseBody) SetData(v *GetRulesCountListResponseBodyData) *GetRulesCountListResponseBody {
	s.Data = v
	return s
}

func (s *GetRulesCountListResponseBody) SetHttpStatusCode(v int32) *GetRulesCountListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetRulesCountListResponseBody) SetMessage(v string) *GetRulesCountListResponseBody {
	s.Message = &v
	return s
}

func (s *GetRulesCountListResponseBody) SetMessages(v *GetRulesCountListResponseBodyMessages) *GetRulesCountListResponseBody {
	s.Messages = v
	return s
}

func (s *GetRulesCountListResponseBody) SetPageNumber(v int32) *GetRulesCountListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetRulesCountListResponseBody) SetPageSize(v int32) *GetRulesCountListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetRulesCountListResponseBody) SetRequestId(v string) *GetRulesCountListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRulesCountListResponseBody) SetSuccess(v bool) *GetRulesCountListResponseBody {
	s.Success = &v
	return s
}

func (s *GetRulesCountListResponseBody) SetTotalCount(v int32) *GetRulesCountListResponseBody {
	s.TotalCount = &v
	return s
}

type GetRulesCountListResponseBodyData struct {
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s GetRulesCountListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRulesCountListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRulesCountListResponseBodyData) SetData(v []*string) *GetRulesCountListResponseBodyData {
	s.Data = v
	return s
}

type GetRulesCountListResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s GetRulesCountListResponseBodyMessages) String() string {
	return tea.Prettify(s)
}

func (s GetRulesCountListResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *GetRulesCountListResponseBodyMessages) SetMessage(v []*string) *GetRulesCountListResponseBodyMessages {
	s.Message = v
	return s
}

type GetRulesCountListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRulesCountListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRulesCountListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRulesCountListResponse) GoString() string {
	return s.String()
}

func (s *GetRulesCountListResponse) SetHeaders(v map[string]*string) *GetRulesCountListResponse {
	s.Headers = v
	return s
}

func (s *GetRulesCountListResponse) SetStatusCode(v int32) *GetRulesCountListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRulesCountListResponse) SetBody(v *GetRulesCountListResponseBody) *GetRulesCountListResponse {
	s.Body = v
	return s
}

type GetScoreInfoRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ""
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetScoreInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetScoreInfoRequest) GoString() string {
	return s.String()
}

func (s *GetScoreInfoRequest) SetBaseMeAgentId(v int64) *GetScoreInfoRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetScoreInfoRequest) SetJsonStr(v string) *GetScoreInfoRequest {
	s.JsonStr = &v
	return s
}

type GetScoreInfoResponseBody struct {
	// example:
	//
	// 200
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetScoreInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetScoreInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetScoreInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetScoreInfoResponseBody) SetCode(v string) *GetScoreInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetScoreInfoResponseBody) SetData(v *GetScoreInfoResponseBodyData) *GetScoreInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetScoreInfoResponseBody) SetMessage(v string) *GetScoreInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetScoreInfoResponseBody) SetRequestId(v string) *GetScoreInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScoreInfoResponseBody) SetSuccess(v bool) *GetScoreInfoResponseBody {
	s.Success = &v
	return s
}

type GetScoreInfoResponseBodyData struct {
	ScorePo []*GetScoreInfoResponseBodyDataScorePo `json:"ScorePo,omitempty" xml:"ScorePo,omitempty" type:"Repeated"`
}

func (s GetScoreInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetScoreInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetScoreInfoResponseBodyData) SetScorePo(v []*GetScoreInfoResponseBodyDataScorePo) *GetScoreInfoResponseBodyData {
	s.ScorePo = v
	return s
}

type GetScoreInfoResponseBodyDataScorePo struct {
	// example:
	//
	// 34
	ScoreId    *int32                                         `json:"ScoreId,omitempty" xml:"ScoreId,omitempty"`
	ScoreInfos *GetScoreInfoResponseBodyDataScorePoScoreInfos `json:"ScoreInfos,omitempty" xml:"ScoreInfos,omitempty" type:"Struct"`
	ScoreName  *string                                        `json:"ScoreName,omitempty" xml:"ScoreName,omitempty"`
}

func (s GetScoreInfoResponseBodyDataScorePo) String() string {
	return tea.Prettify(s)
}

func (s GetScoreInfoResponseBodyDataScorePo) GoString() string {
	return s.String()
}

func (s *GetScoreInfoResponseBodyDataScorePo) SetScoreId(v int32) *GetScoreInfoResponseBodyDataScorePo {
	s.ScoreId = &v
	return s
}

func (s *GetScoreInfoResponseBodyDataScorePo) SetScoreInfos(v *GetScoreInfoResponseBodyDataScorePoScoreInfos) *GetScoreInfoResponseBodyDataScorePo {
	s.ScoreInfos = v
	return s
}

func (s *GetScoreInfoResponseBodyDataScorePo) SetScoreName(v string) *GetScoreInfoResponseBodyDataScorePo {
	s.ScoreName = &v
	return s
}

type GetScoreInfoResponseBodyDataScorePoScoreInfos struct {
	ScoreParam []*GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam `json:"ScoreParam,omitempty" xml:"ScoreParam,omitempty" type:"Repeated"`
}

func (s GetScoreInfoResponseBodyDataScorePoScoreInfos) String() string {
	return tea.Prettify(s)
}

func (s GetScoreInfoResponseBodyDataScorePoScoreInfos) GoString() string {
	return s.String()
}

func (s *GetScoreInfoResponseBodyDataScorePoScoreInfos) SetScoreParam(v []*GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam) *GetScoreInfoResponseBodyDataScorePoScoreInfos {
	s.ScoreParam = v
	return s
}

type GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam struct {
	// example:
	//
	// 32
	ScoreNum *int32 `json:"ScoreNum,omitempty" xml:"ScoreNum,omitempty"`
	// example:
	//
	// 3422
	ScoreSubId   *int32  `json:"ScoreSubId,omitempty" xml:"ScoreSubId,omitempty"`
	ScoreSubName *string `json:"ScoreSubName,omitempty" xml:"ScoreSubName,omitempty"`
	// example:
	//
	// 1
	ScoreType *int32 `json:"ScoreType,omitempty" xml:"ScoreType,omitempty"`
}

func (s GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam) String() string {
	return tea.Prettify(s)
}

func (s GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam) GoString() string {
	return s.String()
}

func (s *GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam) SetScoreNum(v int32) *GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam {
	s.ScoreNum = &v
	return s
}

func (s *GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam) SetScoreSubId(v int32) *GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam {
	s.ScoreSubId = &v
	return s
}

func (s *GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam) SetScoreSubName(v string) *GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam {
	s.ScoreSubName = &v
	return s
}

func (s *GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam) SetScoreType(v int32) *GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam {
	s.ScoreType = &v
	return s
}

type GetScoreInfoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetScoreInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScoreInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetScoreInfoResponse) GoString() string {
	return s.String()
}

func (s *GetScoreInfoResponse) SetHeaders(v map[string]*string) *GetScoreInfoResponse {
	s.Headers = v
	return s
}

func (s *GetScoreInfoResponse) SetStatusCode(v int32) *GetScoreInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScoreInfoResponse) SetBody(v *GetScoreInfoResponseBody) *GetScoreInfoResponse {
	s.Body = v
	return s
}

type GetSkillGroupConfigRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetSkillGroupConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupConfigRequest) GoString() string {
	return s.String()
}

func (s *GetSkillGroupConfigRequest) SetBaseMeAgentId(v int64) *GetSkillGroupConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetSkillGroupConfigRequest) SetJsonStr(v string) *GetSkillGroupConfigRequest {
	s.JsonStr = &v
	return s
}

type GetSkillGroupConfigResponseBody struct {
	// example:
	//
	// 200
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetSkillGroupConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3CEA0495-341B-4482-9AD9-8191EF4***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSkillGroupConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetSkillGroupConfigResponseBody) SetCode(v string) *GetSkillGroupConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetSkillGroupConfigResponseBody) SetData(v *GetSkillGroupConfigResponseBodyData) *GetSkillGroupConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetSkillGroupConfigResponseBody) SetMessage(v string) *GetSkillGroupConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetSkillGroupConfigResponseBody) SetRequestId(v string) *GetSkillGroupConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSkillGroupConfigResponseBody) SetSuccess(v bool) *GetSkillGroupConfigResponseBody {
	s.Success = &v
	return s
}

type GetSkillGroupConfigResponseBodyData struct {
	// example:
	//
	// 1
	AllContentQualityCheck *int32 `json:"AllContentQualityCheck,omitempty" xml:"AllContentQualityCheck,omitempty"`
	// example:
	//
	// 223
	AllRids     *string                                         `json:"AllRids,omitempty" xml:"AllRids,omitempty"`
	AllRuleList *GetSkillGroupConfigResponseBodyDataAllRuleList `json:"AllRuleList,omitempty" xml:"AllRuleList,omitempty" type:"Struct"`
	// example:
	//
	// 2020-12-01T15:12Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1212
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1321
	ModelId *int64 `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// example:
	//
	// xxx
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// example:
	//
	// xxx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0
	QualityCheckType *int32 `json:"QualityCheckType,omitempty" xml:"QualityCheckType,omitempty"`
	// example:
	//
	// 2332
	Rid      *string                                      `json:"Rid,omitempty" xml:"Rid,omitempty"`
	RuleList *GetSkillGroupConfigResponseBodyDataRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Struct"`
	// example:
	//
	// 0
	SkillGroupFrom *int32 `json:"SkillGroupFrom,omitempty" xml:"SkillGroupFrom,omitempty"`
	// example:
	//
	// 111
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// example:
	//
	// xxx
	SkillGroupName *string `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 2020-12-01T19:28Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 123
	VocabId *int64 `json:"VocabId,omitempty" xml:"VocabId,omitempty"`
	// example:
	//
	// test
	VocabName *string `json:"VocabName,omitempty" xml:"VocabName,omitempty"`
}

func (s GetSkillGroupConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSkillGroupConfigResponseBodyData) SetAllContentQualityCheck(v int32) *GetSkillGroupConfigResponseBodyData {
	s.AllContentQualityCheck = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetAllRids(v string) *GetSkillGroupConfigResponseBodyData {
	s.AllRids = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetAllRuleList(v *GetSkillGroupConfigResponseBodyDataAllRuleList) *GetSkillGroupConfigResponseBodyData {
	s.AllRuleList = v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetCreateTime(v string) *GetSkillGroupConfigResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetId(v int64) *GetSkillGroupConfigResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetInstanceId(v string) *GetSkillGroupConfigResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetModelId(v int64) *GetSkillGroupConfigResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetModelName(v string) *GetSkillGroupConfigResponseBodyData {
	s.ModelName = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetName(v string) *GetSkillGroupConfigResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetQualityCheckType(v int32) *GetSkillGroupConfigResponseBodyData {
	s.QualityCheckType = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetRid(v string) *GetSkillGroupConfigResponseBodyData {
	s.Rid = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetRuleList(v *GetSkillGroupConfigResponseBodyDataRuleList) *GetSkillGroupConfigResponseBodyData {
	s.RuleList = v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetSkillGroupFrom(v int32) *GetSkillGroupConfigResponseBodyData {
	s.SkillGroupFrom = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetSkillGroupId(v string) *GetSkillGroupConfigResponseBodyData {
	s.SkillGroupId = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetSkillGroupName(v string) *GetSkillGroupConfigResponseBodyData {
	s.SkillGroupName = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetStatus(v int32) *GetSkillGroupConfigResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetType(v int32) *GetSkillGroupConfigResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetUpdateTime(v string) *GetSkillGroupConfigResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetVocabId(v int64) *GetSkillGroupConfigResponseBodyData {
	s.VocabId = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetVocabName(v string) *GetSkillGroupConfigResponseBodyData {
	s.VocabName = &v
	return s
}

type GetSkillGroupConfigResponseBodyDataAllRuleList struct {
	RuleNameInfo []*GetSkillGroupConfigResponseBodyDataAllRuleListRuleNameInfo `json:"RuleNameInfo,omitempty" xml:"RuleNameInfo,omitempty" type:"Repeated"`
}

func (s GetSkillGroupConfigResponseBodyDataAllRuleList) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupConfigResponseBodyDataAllRuleList) GoString() string {
	return s.String()
}

func (s *GetSkillGroupConfigResponseBodyDataAllRuleList) SetRuleNameInfo(v []*GetSkillGroupConfigResponseBodyDataAllRuleListRuleNameInfo) *GetSkillGroupConfigResponseBodyDataAllRuleList {
	s.RuleNameInfo = v
	return s
}

type GetSkillGroupConfigResponseBodyDataAllRuleListRuleNameInfo struct {
	// example:
	//
	// 12
	Rid *int64 `json:"Rid,omitempty" xml:"Rid,omitempty"`
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s GetSkillGroupConfigResponseBodyDataAllRuleListRuleNameInfo) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupConfigResponseBodyDataAllRuleListRuleNameInfo) GoString() string {
	return s.String()
}

func (s *GetSkillGroupConfigResponseBodyDataAllRuleListRuleNameInfo) SetRid(v int64) *GetSkillGroupConfigResponseBodyDataAllRuleListRuleNameInfo {
	s.Rid = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyDataAllRuleListRuleNameInfo) SetRuleName(v string) *GetSkillGroupConfigResponseBodyDataAllRuleListRuleNameInfo {
	s.RuleName = &v
	return s
}

type GetSkillGroupConfigResponseBodyDataRuleList struct {
	RuleNameInfo []*GetSkillGroupConfigResponseBodyDataRuleListRuleNameInfo `json:"RuleNameInfo,omitempty" xml:"RuleNameInfo,omitempty" type:"Repeated"`
}

func (s GetSkillGroupConfigResponseBodyDataRuleList) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupConfigResponseBodyDataRuleList) GoString() string {
	return s.String()
}

func (s *GetSkillGroupConfigResponseBodyDataRuleList) SetRuleNameInfo(v []*GetSkillGroupConfigResponseBodyDataRuleListRuleNameInfo) *GetSkillGroupConfigResponseBodyDataRuleList {
	s.RuleNameInfo = v
	return s
}

type GetSkillGroupConfigResponseBodyDataRuleListRuleNameInfo struct {
	// example:
	//
	// 222
	Rid *int64 `json:"Rid,omitempty" xml:"Rid,omitempty"`
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s GetSkillGroupConfigResponseBodyDataRuleListRuleNameInfo) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupConfigResponseBodyDataRuleListRuleNameInfo) GoString() string {
	return s.String()
}

func (s *GetSkillGroupConfigResponseBodyDataRuleListRuleNameInfo) SetRid(v int64) *GetSkillGroupConfigResponseBodyDataRuleListRuleNameInfo {
	s.Rid = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyDataRuleListRuleNameInfo) SetRuleName(v string) *GetSkillGroupConfigResponseBodyDataRuleListRuleNameInfo {
	s.RuleName = &v
	return s
}

type GetSkillGroupConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSkillGroupConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSkillGroupConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupConfigResponse) GoString() string {
	return s.String()
}

func (s *GetSkillGroupConfigResponse) SetHeaders(v map[string]*string) *GetSkillGroupConfigResponse {
	s.Headers = v
	return s
}

func (s *GetSkillGroupConfigResponse) SetStatusCode(v int32) *GetSkillGroupConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSkillGroupConfigResponse) SetBody(v *GetSkillGroupConfigResponseBody) *GetSkillGroupConfigResponse {
	s.Body = v
	return s
}

type GetSyncResultRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"pageNumber":1,"pageSize":10,"requiredFields":"asrResult,agent,status,errorMessage,reviewStatus,reviewResult,score,taskId,reviewer,resolver,recording.name,recording.duration,recording.url,hitResult,business","startTime":"2020-12-25 00:00:00","endTime":"2020-12-31 23:59:59"}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetSyncResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSyncResultRequest) GoString() string {
	return s.String()
}

func (s *GetSyncResultRequest) SetBaseMeAgentId(v int64) *GetSyncResultRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetSyncResultRequest) SetJsonStr(v string) *GetSyncResultRequest {
	s.JsonStr = &v
	return s
}

type GetSyncResultResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Count *int32                           `json:"Count,omitempty" xml:"Count,omitempty"`
	Data  []*GetSyncResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 76DB5D8C-5BD9-42A7-B527-5AF3A5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// xxx
	ResultCountId *string `json:"ResultCountId,omitempty" xml:"ResultCountId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSyncResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSyncResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetSyncResultResponseBody) SetCode(v string) *GetSyncResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetSyncResultResponseBody) SetCount(v int32) *GetSyncResultResponseBody {
	s.Count = &v
	return s
}

func (s *GetSyncResultResponseBody) SetData(v []*GetSyncResultResponseBodyData) *GetSyncResultResponseBody {
	s.Data = v
	return s
}

func (s *GetSyncResultResponseBody) SetMessage(v string) *GetSyncResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetSyncResultResponseBody) SetPageNumber(v int32) *GetSyncResultResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetSyncResultResponseBody) SetPageSize(v int32) *GetSyncResultResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetSyncResultResponseBody) SetRequestId(v string) *GetSyncResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSyncResultResponseBody) SetResultCountId(v string) *GetSyncResultResponseBody {
	s.ResultCountId = &v
	return s
}

func (s *GetSyncResultResponseBody) SetSuccess(v bool) *GetSyncResultResponseBody {
	s.Success = &v
	return s
}

type GetSyncResultResponseBodyData struct {
	Agent     *GetSyncResultResponseBodyDataAgent       `json:"Agent,omitempty" xml:"Agent,omitempty" type:"Struct"`
	AsrResult []*GetSyncResultResponseBodyDataAsrResult `json:"AsrResult,omitempty" xml:"AsrResult,omitempty" type:"Repeated"`
	// example:
	//
	// xxx
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// example:
	//
	// 2019-07-24T19:31Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// xxxx
	ErrorMessage *string                                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	HitResult    []*GetSyncResultResponseBodyDataHitResult `json:"HitResult,omitempty" xml:"HitResult,omitempty" type:"Repeated"`
	Recording    *GetSyncResultResponseBodyDataRecording   `json:"Recording,omitempty" xml:"Recording,omitempty" type:"Struct"`
	Resolver     *string                                   `json:"Resolver,omitempty" xml:"Resolver,omitempty"`
	// example:
	//
	// 3
	ReviewResult *int32 `json:"ReviewResult,omitempty" xml:"ReviewResult,omitempty"`
	// example:
	//
	// 1
	ReviewStatus *int32  `json:"ReviewStatus,omitempty" xml:"ReviewStatus,omitempty"`
	Reviewer     *string `json:"Reviewer,omitempty" xml:"Reviewer,omitempty"`
	// example:
	//
	// 100
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 20201231de3d34ec-40fa-4a55-8d27-76ea*****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// xxx
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s GetSyncResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSyncResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSyncResultResponseBodyData) SetAgent(v *GetSyncResultResponseBodyDataAgent) *GetSyncResultResponseBodyData {
	s.Agent = v
	return s
}

func (s *GetSyncResultResponseBodyData) SetAsrResult(v []*GetSyncResultResponseBodyDataAsrResult) *GetSyncResultResponseBodyData {
	s.AsrResult = v
	return s
}

func (s *GetSyncResultResponseBodyData) SetComments(v string) *GetSyncResultResponseBodyData {
	s.Comments = &v
	return s
}

func (s *GetSyncResultResponseBodyData) SetCreateTime(v string) *GetSyncResultResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetSyncResultResponseBodyData) SetErrorMessage(v string) *GetSyncResultResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetSyncResultResponseBodyData) SetHitResult(v []*GetSyncResultResponseBodyDataHitResult) *GetSyncResultResponseBodyData {
	s.HitResult = v
	return s
}

func (s *GetSyncResultResponseBodyData) SetRecording(v *GetSyncResultResponseBodyDataRecording) *GetSyncResultResponseBodyData {
	s.Recording = v
	return s
}

func (s *GetSyncResultResponseBodyData) SetResolver(v string) *GetSyncResultResponseBodyData {
	s.Resolver = &v
	return s
}

func (s *GetSyncResultResponseBodyData) SetReviewResult(v int32) *GetSyncResultResponseBodyData {
	s.ReviewResult = &v
	return s
}

func (s *GetSyncResultResponseBodyData) SetReviewStatus(v int32) *GetSyncResultResponseBodyData {
	s.ReviewStatus = &v
	return s
}

func (s *GetSyncResultResponseBodyData) SetReviewer(v string) *GetSyncResultResponseBodyData {
	s.Reviewer = &v
	return s
}

func (s *GetSyncResultResponseBodyData) SetScore(v int32) *GetSyncResultResponseBodyData {
	s.Score = &v
	return s
}

func (s *GetSyncResultResponseBodyData) SetStatus(v int32) *GetSyncResultResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetSyncResultResponseBodyData) SetTaskId(v string) *GetSyncResultResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetSyncResultResponseBodyData) SetTaskName(v string) *GetSyncResultResponseBodyData {
	s.TaskName = &v
	return s
}

type GetSyncResultResponseBodyDataAgent struct {
	// example:
	//
	// 12221
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SkillGroup *string `json:"SkillGroup,omitempty" xml:"SkillGroup,omitempty"`
}

func (s GetSyncResultResponseBodyDataAgent) String() string {
	return tea.Prettify(s)
}

func (s GetSyncResultResponseBodyDataAgent) GoString() string {
	return s.String()
}

func (s *GetSyncResultResponseBodyDataAgent) SetId(v string) *GetSyncResultResponseBodyDataAgent {
	s.Id = &v
	return s
}

func (s *GetSyncResultResponseBodyDataAgent) SetName(v string) *GetSyncResultResponseBodyDataAgent {
	s.Name = &v
	return s
}

func (s *GetSyncResultResponseBodyDataAgent) SetSkillGroup(v string) *GetSyncResultResponseBodyDataAgent {
	s.SkillGroup = &v
	return s
}

type GetSyncResultResponseBodyDataAsrResult struct {
	// example:
	//
	// 340
	Begin *int64 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// example:
	//
	// 6
	EmotionValue *int32 `json:"EmotionValue,omitempty" xml:"EmotionValue,omitempty"`
	// example:
	//
	// 3000
	End  *int64  `json:"End,omitempty" xml:"End,omitempty"`
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// 11
	SilenceDuration *int32 `json:"SilenceDuration,omitempty" xml:"SilenceDuration,omitempty"`
	// example:
	//
	// 221
	SpeechRate *int32  `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	Words      *string `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s GetSyncResultResponseBodyDataAsrResult) String() string {
	return tea.Prettify(s)
}

func (s GetSyncResultResponseBodyDataAsrResult) GoString() string {
	return s.String()
}

func (s *GetSyncResultResponseBodyDataAsrResult) SetBegin(v int64) *GetSyncResultResponseBodyDataAsrResult {
	s.Begin = &v
	return s
}

func (s *GetSyncResultResponseBodyDataAsrResult) SetEmotionValue(v int32) *GetSyncResultResponseBodyDataAsrResult {
	s.EmotionValue = &v
	return s
}

func (s *GetSyncResultResponseBodyDataAsrResult) SetEnd(v int64) *GetSyncResultResponseBodyDataAsrResult {
	s.End = &v
	return s
}

func (s *GetSyncResultResponseBodyDataAsrResult) SetRole(v string) *GetSyncResultResponseBodyDataAsrResult {
	s.Role = &v
	return s
}

func (s *GetSyncResultResponseBodyDataAsrResult) SetSilenceDuration(v int32) *GetSyncResultResponseBodyDataAsrResult {
	s.SilenceDuration = &v
	return s
}

func (s *GetSyncResultResponseBodyDataAsrResult) SetSpeechRate(v int32) *GetSyncResultResponseBodyDataAsrResult {
	s.SpeechRate = &v
	return s
}

func (s *GetSyncResultResponseBodyDataAsrResult) SetWords(v string) *GetSyncResultResponseBodyDataAsrResult {
	s.Words = &v
	return s
}

type GetSyncResultResponseBodyDataHitResult struct {
	Hits []*GetSyncResultResponseBodyDataHitResultHits `json:"Hits,omitempty" xml:"Hits,omitempty" type:"Repeated"`
	Name *string                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	ReviewResult *int32 `json:"ReviewResult,omitempty" xml:"ReviewResult,omitempty"`
	// example:
	//
	// 1211
	Rid *string `json:"Rid,omitempty" xml:"Rid,omitempty"`
	// example:
	//
	// 2
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetSyncResultResponseBodyDataHitResult) String() string {
	return tea.Prettify(s)
}

func (s GetSyncResultResponseBodyDataHitResult) GoString() string {
	return s.String()
}

func (s *GetSyncResultResponseBodyDataHitResult) SetHits(v []*GetSyncResultResponseBodyDataHitResultHits) *GetSyncResultResponseBodyDataHitResult {
	s.Hits = v
	return s
}

func (s *GetSyncResultResponseBodyDataHitResult) SetName(v string) *GetSyncResultResponseBodyDataHitResult {
	s.Name = &v
	return s
}

func (s *GetSyncResultResponseBodyDataHitResult) SetReviewResult(v int32) *GetSyncResultResponseBodyDataHitResult {
	s.ReviewResult = &v
	return s
}

func (s *GetSyncResultResponseBodyDataHitResult) SetRid(v string) *GetSyncResultResponseBodyDataHitResult {
	s.Rid = &v
	return s
}

func (s *GetSyncResultResponseBodyDataHitResult) SetType(v string) *GetSyncResultResponseBodyDataHitResult {
	s.Type = &v
	return s
}

type GetSyncResultResponseBodyDataHitResultHits struct {
	Cid      []*string                                             `json:"Cid,omitempty" xml:"Cid,omitempty" type:"Repeated"`
	KeyWords []*GetSyncResultResponseBodyDataHitResultHitsKeyWords `json:"KeyWords,omitempty" xml:"KeyWords,omitempty" type:"Repeated"`
	Phrase   *GetSyncResultResponseBodyDataHitResultHitsPhrase     `json:"Phrase,omitempty" xml:"Phrase,omitempty" type:"Struct"`
}

func (s GetSyncResultResponseBodyDataHitResultHits) String() string {
	return tea.Prettify(s)
}

func (s GetSyncResultResponseBodyDataHitResultHits) GoString() string {
	return s.String()
}

func (s *GetSyncResultResponseBodyDataHitResultHits) SetCid(v []*string) *GetSyncResultResponseBodyDataHitResultHits {
	s.Cid = v
	return s
}

func (s *GetSyncResultResponseBodyDataHitResultHits) SetKeyWords(v []*GetSyncResultResponseBodyDataHitResultHitsKeyWords) *GetSyncResultResponseBodyDataHitResultHits {
	s.KeyWords = v
	return s
}

func (s *GetSyncResultResponseBodyDataHitResultHits) SetPhrase(v *GetSyncResultResponseBodyDataHitResultHitsPhrase) *GetSyncResultResponseBodyDataHitResultHits {
	s.Phrase = v
	return s
}

type GetSyncResultResponseBodyDataHitResultHitsKeyWords struct {
	// example:
	//
	// 66666
	Cid *string `json:"Cid,omitempty" xml:"Cid,omitempty"`
	// example:
	//
	// 2
	From *int32 `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// 5
	To  *int32  `json:"To,omitempty" xml:"To,omitempty"`
	Val *string `json:"Val,omitempty" xml:"Val,omitempty"`
}

func (s GetSyncResultResponseBodyDataHitResultHitsKeyWords) String() string {
	return tea.Prettify(s)
}

func (s GetSyncResultResponseBodyDataHitResultHitsKeyWords) GoString() string {
	return s.String()
}

func (s *GetSyncResultResponseBodyDataHitResultHitsKeyWords) SetCid(v string) *GetSyncResultResponseBodyDataHitResultHitsKeyWords {
	s.Cid = &v
	return s
}

func (s *GetSyncResultResponseBodyDataHitResultHitsKeyWords) SetFrom(v int32) *GetSyncResultResponseBodyDataHitResultHitsKeyWords {
	s.From = &v
	return s
}

func (s *GetSyncResultResponseBodyDataHitResultHitsKeyWords) SetTo(v int32) *GetSyncResultResponseBodyDataHitResultHitsKeyWords {
	s.To = &v
	return s
}

func (s *GetSyncResultResponseBodyDataHitResultHitsKeyWords) SetVal(v string) *GetSyncResultResponseBodyDataHitResultHitsKeyWords {
	s.Val = &v
	return s
}

type GetSyncResultResponseBodyDataHitResultHitsPhrase struct {
	// example:
	//
	// 440
	Begin *int64 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// example:
	//
	// 6
	EmotionValue *int32 `json:"EmotionValue,omitempty" xml:"EmotionValue,omitempty"`
	// example:
	//
	// 4000
	End  *int32  `json:"End,omitempty" xml:"End,omitempty"`
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// 1
	SilenceDuration *int32 `json:"SilenceDuration,omitempty" xml:"SilenceDuration,omitempty"`
	// example:
	//
	// 234
	SpeechRate *int32  `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	Words      *string `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s GetSyncResultResponseBodyDataHitResultHitsPhrase) String() string {
	return tea.Prettify(s)
}

func (s GetSyncResultResponseBodyDataHitResultHitsPhrase) GoString() string {
	return s.String()
}

func (s *GetSyncResultResponseBodyDataHitResultHitsPhrase) SetBegin(v int64) *GetSyncResultResponseBodyDataHitResultHitsPhrase {
	s.Begin = &v
	return s
}

func (s *GetSyncResultResponseBodyDataHitResultHitsPhrase) SetEmotionValue(v int32) *GetSyncResultResponseBodyDataHitResultHitsPhrase {
	s.EmotionValue = &v
	return s
}

func (s *GetSyncResultResponseBodyDataHitResultHitsPhrase) SetEnd(v int32) *GetSyncResultResponseBodyDataHitResultHitsPhrase {
	s.End = &v
	return s
}

func (s *GetSyncResultResponseBodyDataHitResultHitsPhrase) SetRole(v string) *GetSyncResultResponseBodyDataHitResultHitsPhrase {
	s.Role = &v
	return s
}

func (s *GetSyncResultResponseBodyDataHitResultHitsPhrase) SetSilenceDuration(v int32) *GetSyncResultResponseBodyDataHitResultHitsPhrase {
	s.SilenceDuration = &v
	return s
}

func (s *GetSyncResultResponseBodyDataHitResultHitsPhrase) SetSpeechRate(v int32) *GetSyncResultResponseBodyDataHitResultHitsPhrase {
	s.SpeechRate = &v
	return s
}

func (s *GetSyncResultResponseBodyDataHitResultHitsPhrase) SetWords(v string) *GetSyncResultResponseBodyDataHitResultHitsPhrase {
	s.Words = &v
	return s
}

type GetSyncResultResponseBodyDataRecording struct {
	Business *string `json:"Business,omitempty" xml:"Business,omitempty"`
	// example:
	//
	// xxx
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// example:
	//
	// 1563967699000
	CallTime *string `json:"CallTime,omitempty" xml:"CallTime,omitempty"`
	// example:
	//
	// 1
	CallType *int32 `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// example:
	//
	// 1888888***
	Callee *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	// example:
	//
	// 0108888****
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// example:
	//
	// xxx
	DataSetName *string `json:"DataSetName,omitempty" xml:"DataSetName,omitempty"`
	// example:
	//
	// 232
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 120
	DurationAudio *int64 `json:"DurationAudio,omitempty" xml:"DurationAudio,omitempty"`
	// example:
	//
	// xxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 123123.wav
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// xxxx
	PrimaryId *string `json:"PrimaryId,omitempty" xml:"PrimaryId,omitempty"`
	// example:
	//
	// xxx
	Remark1 *string `json:"Remark1,omitempty" xml:"Remark1,omitempty"`
	// example:
	//
	// xxx
	Remark2 *string `json:"Remark2,omitempty" xml:"Remark2,omitempty"`
	// example:
	//
	// xxx
	Remark3 *string `json:"Remark3,omitempty" xml:"Remark3,omitempty"`
	// example:
	//
	// http://aliyun.com/xxx.wav
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetSyncResultResponseBodyDataRecording) String() string {
	return tea.Prettify(s)
}

func (s GetSyncResultResponseBodyDataRecording) GoString() string {
	return s.String()
}

func (s *GetSyncResultResponseBodyDataRecording) SetBusiness(v string) *GetSyncResultResponseBodyDataRecording {
	s.Business = &v
	return s
}

func (s *GetSyncResultResponseBodyDataRecording) SetCallId(v string) *GetSyncResultResponseBodyDataRecording {
	s.CallId = &v
	return s
}

func (s *GetSyncResultResponseBodyDataRecording) SetCallTime(v string) *GetSyncResultResponseBodyDataRecording {
	s.CallTime = &v
	return s
}

func (s *GetSyncResultResponseBodyDataRecording) SetCallType(v int32) *GetSyncResultResponseBodyDataRecording {
	s.CallType = &v
	return s
}

func (s *GetSyncResultResponseBodyDataRecording) SetCallee(v string) *GetSyncResultResponseBodyDataRecording {
	s.Callee = &v
	return s
}

func (s *GetSyncResultResponseBodyDataRecording) SetCaller(v string) *GetSyncResultResponseBodyDataRecording {
	s.Caller = &v
	return s
}

func (s *GetSyncResultResponseBodyDataRecording) SetDataSetName(v string) *GetSyncResultResponseBodyDataRecording {
	s.DataSetName = &v
	return s
}

func (s *GetSyncResultResponseBodyDataRecording) SetDuration(v int64) *GetSyncResultResponseBodyDataRecording {
	s.Duration = &v
	return s
}

func (s *GetSyncResultResponseBodyDataRecording) SetDurationAudio(v int64) *GetSyncResultResponseBodyDataRecording {
	s.DurationAudio = &v
	return s
}

func (s *GetSyncResultResponseBodyDataRecording) SetId(v string) *GetSyncResultResponseBodyDataRecording {
	s.Id = &v
	return s
}

func (s *GetSyncResultResponseBodyDataRecording) SetName(v string) *GetSyncResultResponseBodyDataRecording {
	s.Name = &v
	return s
}

func (s *GetSyncResultResponseBodyDataRecording) SetPrimaryId(v string) *GetSyncResultResponseBodyDataRecording {
	s.PrimaryId = &v
	return s
}

func (s *GetSyncResultResponseBodyDataRecording) SetRemark1(v string) *GetSyncResultResponseBodyDataRecording {
	s.Remark1 = &v
	return s
}

func (s *GetSyncResultResponseBodyDataRecording) SetRemark2(v string) *GetSyncResultResponseBodyDataRecording {
	s.Remark2 = &v
	return s
}

func (s *GetSyncResultResponseBodyDataRecording) SetRemark3(v string) *GetSyncResultResponseBodyDataRecording {
	s.Remark3 = &v
	return s
}

func (s *GetSyncResultResponseBodyDataRecording) SetUrl(v string) *GetSyncResultResponseBodyDataRecording {
	s.Url = &v
	return s
}

type GetSyncResultResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSyncResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSyncResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSyncResultResponse) GoString() string {
	return s.String()
}

func (s *GetSyncResultResponse) SetHeaders(v map[string]*string) *GetSyncResultResponse {
	s.Headers = v
	return s
}

func (s *GetSyncResultResponse) SetStatusCode(v int32) *GetSyncResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSyncResultResponse) SetBody(v *GetSyncResultResponseBody) *GetSyncResultResponse {
	s.Body = v
	return s
}

type GetWarningStrategyConfigRequest struct {
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetWarningStrategyConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWarningStrategyConfigRequest) GoString() string {
	return s.String()
}

func (s *GetWarningStrategyConfigRequest) SetBaseMeAgentId(v int64) *GetWarningStrategyConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetWarningStrategyConfigRequest) SetJsonStr(v string) *GetWarningStrategyConfigRequest {
	s.JsonStr = &v
	return s
}

type GetWarningStrategyConfigResponseBody struct {
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetWarningStrategyConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetWarningStrategyConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWarningStrategyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetWarningStrategyConfigResponseBody) SetCode(v string) *GetWarningStrategyConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBody) SetData(v *GetWarningStrategyConfigResponseBodyData) *GetWarningStrategyConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetWarningStrategyConfigResponseBody) SetMessage(v string) *GetWarningStrategyConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBody) SetRequestId(v string) *GetWarningStrategyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBody) SetSuccess(v bool) *GetWarningStrategyConfigResponseBody {
	s.Success = &v
	return s
}

type GetWarningStrategyConfigResponseBodyData struct {
	Id                  *int64                                                       `json:"Id,omitempty" xml:"Id,omitempty"`
	IntervalTime        *int64                                                       `json:"IntervalTime,omitempty" xml:"IntervalTime,omitempty"`
	Lambda              *string                                                      `json:"Lambda,omitempty" xml:"Lambda,omitempty"`
	Level               *int64                                                       `json:"Level,omitempty" xml:"Level,omitempty"`
	MaxNumber           *int64                                                       `json:"MaxNumber,omitempty" xml:"MaxNumber,omitempty"`
	Name                *string                                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	WarningStrategyList *GetWarningStrategyConfigResponseBodyDataWarningStrategyList `json:"WarningStrategyList,omitempty" xml:"WarningStrategyList,omitempty" type:"Struct"`
}

func (s GetWarningStrategyConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetWarningStrategyConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWarningStrategyConfigResponseBodyData) SetId(v int64) *GetWarningStrategyConfigResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyData) SetIntervalTime(v int64) *GetWarningStrategyConfigResponseBodyData {
	s.IntervalTime = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyData) SetLambda(v string) *GetWarningStrategyConfigResponseBodyData {
	s.Lambda = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyData) SetLevel(v int64) *GetWarningStrategyConfigResponseBodyData {
	s.Level = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyData) SetMaxNumber(v int64) *GetWarningStrategyConfigResponseBodyData {
	s.MaxNumber = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyData) SetName(v string) *GetWarningStrategyConfigResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyData) SetWarningStrategyList(v *GetWarningStrategyConfigResponseBodyDataWarningStrategyList) *GetWarningStrategyConfigResponseBodyData {
	s.WarningStrategyList = v
	return s
}

type GetWarningStrategyConfigResponseBodyDataWarningStrategyList struct {
	WarningStrategyList []*GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList `json:"warningStrategyList,omitempty" xml:"warningStrategyList,omitempty" type:"Repeated"`
}

func (s GetWarningStrategyConfigResponseBodyDataWarningStrategyList) String() string {
	return tea.Prettify(s)
}

func (s GetWarningStrategyConfigResponseBodyDataWarningStrategyList) GoString() string {
	return s.String()
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyList) SetWarningStrategyList(v []*GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) *GetWarningStrategyConfigResponseBodyDataWarningStrategyList {
	s.WarningStrategyList = v
	return s
}

type GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList struct {
	Code                *string                                                                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Duration            *int64                                                                               `json:"Duration,omitempty" xml:"Duration,omitempty"`
	DurationExpression  *int64                                                                               `json:"DurationExpression,omitempty" xml:"DurationExpression,omitempty"`
	HitNumber           *int64                                                                               `json:"HitNumber,omitempty" xml:"HitNumber,omitempty"`
	HitNumberExpression *int64                                                                               `json:"HitNumberExpression,omitempty" xml:"HitNumberExpression,omitempty"`
	HitRuleList         *string                                                                              `json:"HitRuleList,omitempty" xml:"HitRuleList,omitempty"`
	HitType             *int64                                                                               `json:"HitType,omitempty" xml:"HitType,omitempty"`
	Id                  *int64                                                                               `json:"Id,omitempty" xml:"Id,omitempty"`
	Range               *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyListRange `json:"Range,omitempty" xml:"Range,omitempty" type:"Struct"`
	Status              *int64                                                                               `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) String() string {
	return tea.Prettify(s)
}

func (s GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) GoString() string {
	return s.String()
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) SetCode(v string) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList {
	s.Code = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) SetDuration(v int64) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList {
	s.Duration = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) SetDurationExpression(v int64) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList {
	s.DurationExpression = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) SetHitNumber(v int64) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList {
	s.HitNumber = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) SetHitNumberExpression(v int64) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList {
	s.HitNumberExpression = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) SetHitRuleList(v string) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList {
	s.HitRuleList = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) SetHitType(v int64) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList {
	s.HitType = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) SetId(v int64) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList {
	s.Id = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) SetRange(v *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyListRange) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList {
	s.Range = v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) SetStatus(v int64) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList {
	s.Status = &v
	return s
}

type GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyListRange struct {
	RangeNum *int64 `json:"RangeNum,omitempty" xml:"RangeNum,omitempty"`
	Type     *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyListRange) String() string {
	return tea.Prettify(s)
}

func (s GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyListRange) GoString() string {
	return s.String()
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyListRange) SetRangeNum(v int64) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyListRange {
	s.RangeNum = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyListRange) SetType(v int64) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyListRange {
	s.Type = &v
	return s
}

type GetWarningStrategyConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWarningStrategyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWarningStrategyConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWarningStrategyConfigResponse) GoString() string {
	return s.String()
}

func (s *GetWarningStrategyConfigResponse) SetHeaders(v map[string]*string) *GetWarningStrategyConfigResponse {
	s.Headers = v
	return s
}

func (s *GetWarningStrategyConfigResponse) SetStatusCode(v int32) *GetWarningStrategyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWarningStrategyConfigResponse) SetBody(v *GetWarningStrategyConfigResponseBody) *GetWarningStrategyConfigResponse {
	s.Body = v
	return s
}

type HandleComplaintRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s HandleComplaintRequest) String() string {
	return tea.Prettify(s)
}

func (s HandleComplaintRequest) GoString() string {
	return s.String()
}

func (s *HandleComplaintRequest) SetBaseMeAgentId(v int64) *HandleComplaintRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *HandleComplaintRequest) SetJsonStr(v string) *HandleComplaintRequest {
	s.JsonStr = &v
	return s
}

type HandleComplaintResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9987D326-83D9-4A42-B9A5-0B27F9B40539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s HandleComplaintResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HandleComplaintResponseBody) GoString() string {
	return s.String()
}

func (s *HandleComplaintResponseBody) SetCode(v string) *HandleComplaintResponseBody {
	s.Code = &v
	return s
}

func (s *HandleComplaintResponseBody) SetData(v string) *HandleComplaintResponseBody {
	s.Data = &v
	return s
}

func (s *HandleComplaintResponseBody) SetMessage(v string) *HandleComplaintResponseBody {
	s.Message = &v
	return s
}

func (s *HandleComplaintResponseBody) SetRequestId(v string) *HandleComplaintResponseBody {
	s.RequestId = &v
	return s
}

func (s *HandleComplaintResponseBody) SetSuccess(v bool) *HandleComplaintResponseBody {
	s.Success = &v
	return s
}

type HandleComplaintResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HandleComplaintResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HandleComplaintResponse) String() string {
	return tea.Prettify(s)
}

func (s HandleComplaintResponse) GoString() string {
	return s.String()
}

func (s *HandleComplaintResponse) SetHeaders(v map[string]*string) *HandleComplaintResponse {
	s.Headers = v
	return s
}

func (s *HandleComplaintResponse) SetStatusCode(v int32) *HandleComplaintResponse {
	s.StatusCode = &v
	return s
}

func (s *HandleComplaintResponse) SetBody(v *HandleComplaintResponseBody) *HandleComplaintResponse {
	s.Body = v
	return s
}

type InsertScoreForApiRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s InsertScoreForApiRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertScoreForApiRequest) GoString() string {
	return s.String()
}

func (s *InsertScoreForApiRequest) SetBaseMeAgentId(v int64) *InsertScoreForApiRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *InsertScoreForApiRequest) SetJsonStr(v string) *InsertScoreForApiRequest {
	s.JsonStr = &v
	return s
}

type InsertScoreForApiResponseBody struct {
	// example:
	//
	// 200
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *InsertScoreForApiResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 96138D8D-8D26-4E41-BFF4-77AED1088BBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InsertScoreForApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertScoreForApiResponseBody) GoString() string {
	return s.String()
}

func (s *InsertScoreForApiResponseBody) SetCode(v string) *InsertScoreForApiResponseBody {
	s.Code = &v
	return s
}

func (s *InsertScoreForApiResponseBody) SetData(v *InsertScoreForApiResponseBodyData) *InsertScoreForApiResponseBody {
	s.Data = v
	return s
}

func (s *InsertScoreForApiResponseBody) SetMessage(v string) *InsertScoreForApiResponseBody {
	s.Message = &v
	return s
}

func (s *InsertScoreForApiResponseBody) SetRequestId(v string) *InsertScoreForApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertScoreForApiResponseBody) SetSuccess(v bool) *InsertScoreForApiResponseBody {
	s.Success = &v
	return s
}

type InsertScoreForApiResponseBodyData struct {
	// example:
	//
	// 5728
	ScoreId   *int64  `json:"ScoreId,omitempty" xml:"ScoreId,omitempty"`
	ScoreName *string `json:"ScoreName,omitempty" xml:"ScoreName,omitempty"`
}

func (s InsertScoreForApiResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s InsertScoreForApiResponseBodyData) GoString() string {
	return s.String()
}

func (s *InsertScoreForApiResponseBodyData) SetScoreId(v int64) *InsertScoreForApiResponseBodyData {
	s.ScoreId = &v
	return s
}

func (s *InsertScoreForApiResponseBodyData) SetScoreName(v string) *InsertScoreForApiResponseBodyData {
	s.ScoreName = &v
	return s
}

type InsertScoreForApiResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsertScoreForApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsertScoreForApiResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertScoreForApiResponse) GoString() string {
	return s.String()
}

func (s *InsertScoreForApiResponse) SetHeaders(v map[string]*string) *InsertScoreForApiResponse {
	s.Headers = v
	return s
}

func (s *InsertScoreForApiResponse) SetStatusCode(v int32) *InsertScoreForApiResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertScoreForApiResponse) SetBody(v *InsertScoreForApiResponseBody) *InsertScoreForApiResponse {
	s.Body = v
	return s
}

type InsertSubScoreForApiRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s InsertSubScoreForApiRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertSubScoreForApiRequest) GoString() string {
	return s.String()
}

func (s *InsertSubScoreForApiRequest) SetBaseMeAgentId(v int64) *InsertSubScoreForApiRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *InsertSubScoreForApiRequest) SetJsonStr(v string) *InsertSubScoreForApiRequest {
	s.JsonStr = &v
	return s
}

type InsertSubScoreForApiResponseBody struct {
	// example:
	//
	// 200
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *InsertSubScoreForApiResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D34079C5-AA2F-490E-ADD8-5BFF08AAE207
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InsertSubScoreForApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertSubScoreForApiResponseBody) GoString() string {
	return s.String()
}

func (s *InsertSubScoreForApiResponseBody) SetCode(v string) *InsertSubScoreForApiResponseBody {
	s.Code = &v
	return s
}

func (s *InsertSubScoreForApiResponseBody) SetData(v *InsertSubScoreForApiResponseBodyData) *InsertSubScoreForApiResponseBody {
	s.Data = v
	return s
}

func (s *InsertSubScoreForApiResponseBody) SetMessage(v string) *InsertSubScoreForApiResponseBody {
	s.Message = &v
	return s
}

func (s *InsertSubScoreForApiResponseBody) SetRequestId(v string) *InsertSubScoreForApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertSubScoreForApiResponseBody) SetSuccess(v bool) *InsertSubScoreForApiResponseBody {
	s.Success = &v
	return s
}

type InsertSubScoreForApiResponseBodyData struct {
	// example:
	//
	// 5730
	ScoreSubId   *int64  `json:"ScoreSubId,omitempty" xml:"ScoreSubId,omitempty"`
	ScoreSubName *string `json:"ScoreSubName,omitempty" xml:"ScoreSubName,omitempty"`
}

func (s InsertSubScoreForApiResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s InsertSubScoreForApiResponseBodyData) GoString() string {
	return s.String()
}

func (s *InsertSubScoreForApiResponseBodyData) SetScoreSubId(v int64) *InsertSubScoreForApiResponseBodyData {
	s.ScoreSubId = &v
	return s
}

func (s *InsertSubScoreForApiResponseBodyData) SetScoreSubName(v string) *InsertSubScoreForApiResponseBodyData {
	s.ScoreSubName = &v
	return s
}

type InsertSubScoreForApiResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsertSubScoreForApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsertSubScoreForApiResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertSubScoreForApiResponse) GoString() string {
	return s.String()
}

func (s *InsertSubScoreForApiResponse) SetHeaders(v map[string]*string) *InsertSubScoreForApiResponse {
	s.Headers = v
	return s
}

func (s *InsertSubScoreForApiResponse) SetStatusCode(v int32) *InsertSubScoreForApiResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertSubScoreForApiResponse) SetBody(v *InsertSubScoreForApiResponseBody) *InsertSubScoreForApiResponse {
	s.Body = v
	return s
}

type InvalidRuleRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"ruleIds":[3,4]}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s InvalidRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s InvalidRuleRequest) GoString() string {
	return s.String()
}

func (s *InvalidRuleRequest) SetBaseMeAgentId(v int64) *InvalidRuleRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *InvalidRuleRequest) SetJsonStr(v string) *InvalidRuleRequest {
	s.JsonStr = &v
	return s
}

type InvalidRuleResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9987D326-83D9-4A42-B9A5-0B27F9B40539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InvalidRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InvalidRuleResponseBody) GoString() string {
	return s.String()
}

func (s *InvalidRuleResponseBody) SetCode(v string) *InvalidRuleResponseBody {
	s.Code = &v
	return s
}

func (s *InvalidRuleResponseBody) SetData(v bool) *InvalidRuleResponseBody {
	s.Data = &v
	return s
}

func (s *InvalidRuleResponseBody) SetMessage(v string) *InvalidRuleResponseBody {
	s.Message = &v
	return s
}

func (s *InvalidRuleResponseBody) SetRequestId(v string) *InvalidRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvalidRuleResponseBody) SetSuccess(v bool) *InvalidRuleResponseBody {
	s.Success = &v
	return s
}

type InvalidRuleResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvalidRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvalidRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s InvalidRuleResponse) GoString() string {
	return s.String()
}

func (s *InvalidRuleResponse) SetHeaders(v map[string]*string) *InvalidRuleResponse {
	s.Headers = v
	return s
}

func (s *InvalidRuleResponse) SetStatusCode(v int32) *InvalidRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *InvalidRuleResponse) SetBody(v *InvalidRuleResponseBody) *InvalidRuleResponse {
	s.Body = v
	return s
}

type ListAsrVocabRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"pageSize":1}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s ListAsrVocabRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAsrVocabRequest) GoString() string {
	return s.String()
}

func (s *ListAsrVocabRequest) SetBaseMeAgentId(v int64) *ListAsrVocabRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *ListAsrVocabRequest) SetJsonStr(v string) *ListAsrVocabRequest {
	s.JsonStr = &v
	return s
}

type ListAsrVocabResponseBody struct {
	// example:
	//
	// 200
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListAsrVocabResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 66E1ACB8-17B2-4BE8-8581-954A8EE1324B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAsrVocabResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAsrVocabResponseBody) GoString() string {
	return s.String()
}

func (s *ListAsrVocabResponseBody) SetCode(v string) *ListAsrVocabResponseBody {
	s.Code = &v
	return s
}

func (s *ListAsrVocabResponseBody) SetData(v *ListAsrVocabResponseBodyData) *ListAsrVocabResponseBody {
	s.Data = v
	return s
}

func (s *ListAsrVocabResponseBody) SetMessage(v string) *ListAsrVocabResponseBody {
	s.Message = &v
	return s
}

func (s *ListAsrVocabResponseBody) SetRequestId(v string) *ListAsrVocabResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAsrVocabResponseBody) SetSuccess(v bool) *ListAsrVocabResponseBody {
	s.Success = &v
	return s
}

type ListAsrVocabResponseBodyData struct {
	AsrVocab []*ListAsrVocabResponseBodyDataAsrVocab `json:"AsrVocab,omitempty" xml:"AsrVocab,omitempty" type:"Repeated"`
}

func (s ListAsrVocabResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAsrVocabResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAsrVocabResponseBodyData) SetAsrVocab(v []*ListAsrVocabResponseBodyDataAsrVocab) *ListAsrVocabResponseBodyData {
	s.AsrVocab = v
	return s
}

type ListAsrVocabResponseBodyDataAsrVocab struct {
	// example:
	//
	// 2019-04-15T14:57Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 18
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2019-04-15T14:57Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// a01daaaxxxxxxxxx
	VocabularyId *string `json:"VocabularyId,omitempty" xml:"VocabularyId,omitempty"`
}

func (s ListAsrVocabResponseBodyDataAsrVocab) String() string {
	return tea.Prettify(s)
}

func (s ListAsrVocabResponseBodyDataAsrVocab) GoString() string {
	return s.String()
}

func (s *ListAsrVocabResponseBodyDataAsrVocab) SetCreateTime(v string) *ListAsrVocabResponseBodyDataAsrVocab {
	s.CreateTime = &v
	return s
}

func (s *ListAsrVocabResponseBodyDataAsrVocab) SetId(v string) *ListAsrVocabResponseBodyDataAsrVocab {
	s.Id = &v
	return s
}

func (s *ListAsrVocabResponseBodyDataAsrVocab) SetName(v string) *ListAsrVocabResponseBodyDataAsrVocab {
	s.Name = &v
	return s
}

func (s *ListAsrVocabResponseBodyDataAsrVocab) SetUpdateTime(v string) *ListAsrVocabResponseBodyDataAsrVocab {
	s.UpdateTime = &v
	return s
}

func (s *ListAsrVocabResponseBodyDataAsrVocab) SetVocabularyId(v string) *ListAsrVocabResponseBodyDataAsrVocab {
	s.VocabularyId = &v
	return s
}

type ListAsrVocabResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAsrVocabResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAsrVocabResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAsrVocabResponse) GoString() string {
	return s.String()
}

func (s *ListAsrVocabResponse) SetHeaders(v map[string]*string) *ListAsrVocabResponse {
	s.Headers = v
	return s
}

func (s *ListAsrVocabResponse) SetStatusCode(v int32) *ListAsrVocabResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAsrVocabResponse) SetBody(v *ListAsrVocabResponseBody) *ListAsrVocabResponse {
	s.Body = v
	return s
}

type ListDataSetRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// {"pageNumber":1,"pageSize":10}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s ListDataSetRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataSetRequest) GoString() string {
	return s.String()
}

func (s *ListDataSetRequest) SetBaseMeAgentId(v int64) *ListDataSetRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *ListDataSetRequest) SetJsonStr(v string) *ListDataSetRequest {
	s.JsonStr = &v
	return s
}

type ListDataSetResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 23
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 10
	CurrentPage *int32                       `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Data        *ListDataSetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *ListDataSetResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 96138D8D-8D26-4E41-BFF4-77AED1088BBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataSetResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSetResponseBody) SetCode(v string) *ListDataSetResponseBody {
	s.Code = &v
	return s
}

func (s *ListDataSetResponseBody) SetCount(v int32) *ListDataSetResponseBody {
	s.Count = &v
	return s
}

func (s *ListDataSetResponseBody) SetCurrentPage(v int32) *ListDataSetResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListDataSetResponseBody) SetData(v *ListDataSetResponseBodyData) *ListDataSetResponseBody {
	s.Data = v
	return s
}

func (s *ListDataSetResponseBody) SetHttpStatusCode(v int32) *ListDataSetResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDataSetResponseBody) SetMessage(v string) *ListDataSetResponseBody {
	s.Message = &v
	return s
}

func (s *ListDataSetResponseBody) SetMessages(v *ListDataSetResponseBodyMessages) *ListDataSetResponseBody {
	s.Messages = v
	return s
}

func (s *ListDataSetResponseBody) SetPageNumber(v int32) *ListDataSetResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDataSetResponseBody) SetPageSize(v int32) *ListDataSetResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDataSetResponseBody) SetRequestId(v string) *ListDataSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSetResponseBody) SetSuccess(v bool) *ListDataSetResponseBody {
	s.Success = &v
	return s
}

type ListDataSetResponseBodyData struct {
	Data []*ListDataSetResponseBodyDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s ListDataSetResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDataSetResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataSetResponseBodyData) SetData(v []*ListDataSetResponseBodyDataData) *ListDataSetResponseBodyData {
	s.Data = v
	return s
}

type ListDataSetResponseBodyDataData struct {
	// example:
	//
	// 1
	AutoTranscoding *int32 `json:"AutoTranscoding,omitempty" xml:"AutoTranscoding,omitempty"`
	// example:
	//
	// 0
	ChannelId0 *int32 `json:"ChannelId0,omitempty" xml:"ChannelId0,omitempty"`
	// example:
	//
	// 1
	ChannelId1 *int32 `json:"ChannelId1,omitempty" xml:"ChannelId1,omitempty"`
	// example:
	//
	// 1
	ChannelType *int32 `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// example:
	//
	// 2019-06-20T17:33Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 0
	CreateType *int32 `json:"CreateType,omitempty" xml:"CreateType,omitempty"`
	// example:
	//
	// 3
	DataSetType *int32 `json:"DataSetType,omitempty" xml:"DataSetType,omitempty"`
	// example:
	//
	// 0
	IsDelete *int32 `json:"IsDelete,omitempty" xml:"IsDelete,omitempty"`
	// example:
	//
	// filesFromLocal/ef7ff45c147a4a5e882c925f9a75ac43
	RoleConfigProp *string `json:"RoleConfigProp,omitempty" xml:"RoleConfigProp,omitempty"`
	// example:
	//
	// 1
	RoleConfigStatus *int32 `json:"RoleConfigStatus,omitempty" xml:"RoleConfigStatus,omitempty"`
	// example:
	//
	// xx
	RoleConfigTask *string `json:"RoleConfigTask,omitempty" xml:"RoleConfigTask,omitempty"`
	// example:
	//
	// “”
	SetBucketName *string `json:"SetBucketName,omitempty" xml:"SetBucketName,omitempty"`
	// example:
	//
	// “”
	SetDomain *string `json:"SetDomain,omitempty" xml:"SetDomain,omitempty"`
	// example:
	//
	// “”
	SetFolderName *string `json:"SetFolderName,omitempty" xml:"SetFolderName,omitempty"`
	// example:
	//
	// 1
	SetId   *int64  `json:"SetId,omitempty" xml:"SetId,omitempty"`
	SetName *string `json:"SetName,omitempty" xml:"SetName,omitempty"`
	// example:
	//
	// 1
	SetNumber *int32 `json:"SetNumber,omitempty" xml:"SetNumber,omitempty"`
	// example:
	//
	// “”
	SetRoleArn *string `json:"SetRoleArn,omitempty" xml:"SetRoleArn,omitempty"`
	// example:
	//
	// 1
	SetType *int32 `json:"SetType,omitempty" xml:"SetType,omitempty"`
	// example:
	//
	// 11
	SourceDataType *int32 `json:"SourceDataType,omitempty" xml:"SourceDataType,omitempty"`
	// example:
	//
	// xx
	SubDir *string `json:"SubDir,omitempty" xml:"SubDir,omitempty"`
	// example:
	//
	// 2022-05-10T10:34Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 1
	UserGroup *string `json:"UserGroup,omitempty" xml:"UserGroup,omitempty"`
}

func (s ListDataSetResponseBodyDataData) String() string {
	return tea.Prettify(s)
}

func (s ListDataSetResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *ListDataSetResponseBodyDataData) SetAutoTranscoding(v int32) *ListDataSetResponseBodyDataData {
	s.AutoTranscoding = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetChannelId0(v int32) *ListDataSetResponseBodyDataData {
	s.ChannelId0 = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetChannelId1(v int32) *ListDataSetResponseBodyDataData {
	s.ChannelId1 = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetChannelType(v int32) *ListDataSetResponseBodyDataData {
	s.ChannelType = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetCreateTime(v string) *ListDataSetResponseBodyDataData {
	s.CreateTime = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetCreateType(v int32) *ListDataSetResponseBodyDataData {
	s.CreateType = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetDataSetType(v int32) *ListDataSetResponseBodyDataData {
	s.DataSetType = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetIsDelete(v int32) *ListDataSetResponseBodyDataData {
	s.IsDelete = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetRoleConfigProp(v string) *ListDataSetResponseBodyDataData {
	s.RoleConfigProp = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetRoleConfigStatus(v int32) *ListDataSetResponseBodyDataData {
	s.RoleConfigStatus = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetRoleConfigTask(v string) *ListDataSetResponseBodyDataData {
	s.RoleConfigTask = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetSetBucketName(v string) *ListDataSetResponseBodyDataData {
	s.SetBucketName = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetSetDomain(v string) *ListDataSetResponseBodyDataData {
	s.SetDomain = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetSetFolderName(v string) *ListDataSetResponseBodyDataData {
	s.SetFolderName = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetSetId(v int64) *ListDataSetResponseBodyDataData {
	s.SetId = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetSetName(v string) *ListDataSetResponseBodyDataData {
	s.SetName = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetSetNumber(v int32) *ListDataSetResponseBodyDataData {
	s.SetNumber = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetSetRoleArn(v string) *ListDataSetResponseBodyDataData {
	s.SetRoleArn = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetSetType(v int32) *ListDataSetResponseBodyDataData {
	s.SetType = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetSourceDataType(v int32) *ListDataSetResponseBodyDataData {
	s.SourceDataType = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetSubDir(v string) *ListDataSetResponseBodyDataData {
	s.SubDir = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetUpdateTime(v string) *ListDataSetResponseBodyDataData {
	s.UpdateTime = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetUserGroup(v string) *ListDataSetResponseBodyDataData {
	s.UserGroup = &v
	return s
}

type ListDataSetResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s ListDataSetResponseBodyMessages) String() string {
	return tea.Prettify(s)
}

func (s ListDataSetResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *ListDataSetResponseBodyMessages) SetMessage(v []*string) *ListDataSetResponseBodyMessages {
	s.Message = v
	return s
}

type ListDataSetResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataSetResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataSetResponse) GoString() string {
	return s.String()
}

func (s *ListDataSetResponse) SetHeaders(v map[string]*string) *ListDataSetResponse {
	s.Headers = v
	return s
}

func (s *ListDataSetResponse) SetStatusCode(v int32) *ListDataSetResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataSetResponse) SetBody(v *ListDataSetResponseBody) *ListDataSetResponse {
	s.Body = v
	return s
}

type ListPrecisionTaskRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// "{"pageSize":10,"pageNumber":1}"
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s ListPrecisionTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPrecisionTaskRequest) GoString() string {
	return s.String()
}

func (s *ListPrecisionTaskRequest) SetBaseMeAgentId(v int64) *ListPrecisionTaskRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *ListPrecisionTaskRequest) SetJsonStr(v string) *ListPrecisionTaskRequest {
	s.JsonStr = &v
	return s
}

type ListPrecisionTaskResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 22
	Count *int32                             `json:"Count,omitempty" xml:"Count,omitempty"`
	Data  *ListPrecisionTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24F4CE647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListPrecisionTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPrecisionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrecisionTaskResponseBody) SetCode(v string) *ListPrecisionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ListPrecisionTaskResponseBody) SetCount(v int32) *ListPrecisionTaskResponseBody {
	s.Count = &v
	return s
}

func (s *ListPrecisionTaskResponseBody) SetData(v *ListPrecisionTaskResponseBodyData) *ListPrecisionTaskResponseBody {
	s.Data = v
	return s
}

func (s *ListPrecisionTaskResponseBody) SetMessage(v string) *ListPrecisionTaskResponseBody {
	s.Message = &v
	return s
}

func (s *ListPrecisionTaskResponseBody) SetPageNumber(v int32) *ListPrecisionTaskResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPrecisionTaskResponseBody) SetPageSize(v int32) *ListPrecisionTaskResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPrecisionTaskResponseBody) SetRequestId(v string) *ListPrecisionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPrecisionTaskResponseBody) SetSuccess(v bool) *ListPrecisionTaskResponseBody {
	s.Success = &v
	return s
}

type ListPrecisionTaskResponseBodyData struct {
	PrecisionTask []*ListPrecisionTaskResponseBodyDataPrecisionTask `json:"PrecisionTask,omitempty" xml:"PrecisionTask,omitempty" type:"Repeated"`
}

func (s ListPrecisionTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListPrecisionTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPrecisionTaskResponseBodyData) SetPrecisionTask(v []*ListPrecisionTaskResponseBodyDataPrecisionTask) *ListPrecisionTaskResponseBodyData {
	s.PrecisionTask = v
	return s
}

type ListPrecisionTaskResponseBodyDataPrecisionTask struct {
	// example:
	//
	// 2020-03-10 20:26:29
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1212
	DataSetId   *int64  `json:"DataSetId,omitempty" xml:"DataSetId,omitempty"`
	DataSetName *string `json:"DataSetName,omitempty" xml:"DataSetName,omitempty"`
	// example:
	//
	// 331311
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 32
	IncorrectWords *int32                                                    `json:"IncorrectWords,omitempty" xml:"IncorrectWords,omitempty"`
	Name           *string                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	Precisions     *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisions `json:"Precisions,omitempty" xml:"Precisions,omitempty" type:"Struct"`
	// example:
	//
	// 3
	Source *int32 `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 7C1DEF5F-2C18-4D36-99C6-8C276F781796
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 21
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 2020-03-10 20:26:29
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 3
	VerifiedCount *int32 `json:"VerifiedCount,omitempty" xml:"VerifiedCount,omitempty"`
}

func (s ListPrecisionTaskResponseBodyDataPrecisionTask) String() string {
	return tea.Prettify(s)
}

func (s ListPrecisionTaskResponseBodyDataPrecisionTask) GoString() string {
	return s.String()
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetCreateTime(v string) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.CreateTime = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetDataSetId(v int64) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.DataSetId = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetDataSetName(v string) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.DataSetName = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetDuration(v int32) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.Duration = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetIncorrectWords(v int32) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.IncorrectWords = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetName(v string) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.Name = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetPrecisions(v *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisions) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.Precisions = v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetSource(v int32) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.Source = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetStatus(v int32) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.Status = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetTaskId(v string) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.TaskId = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetTotalCount(v int32) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.TotalCount = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetUpdateTime(v string) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.UpdateTime = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetVerifiedCount(v int32) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.VerifiedCount = &v
	return s
}

type ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisions struct {
	Precision []*ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision `json:"Precision,omitempty" xml:"Precision,omitempty" type:"Repeated"`
}

func (s ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisions) String() string {
	return tea.Prettify(s)
}

func (s ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisions) GoString() string {
	return s.String()
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisions) SetPrecision(v []*ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision) *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisions {
	s.Precision = v
	return s
}

type ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision struct {
	// example:
	//
	// 2020-03-10 20:26:29
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2321
	ModelId   *int64  `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// example:
	//
	// 0.98
	Precision *float32 `json:"Precision,omitempty" xml:"Precision,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 7C1DEF5F-2C18-4D36-99C6-8C276F781796
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision) String() string {
	return tea.Prettify(s)
}

func (s ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision) GoString() string {
	return s.String()
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision) SetCreateTime(v string) *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision {
	s.CreateTime = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision) SetModelId(v int64) *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision {
	s.ModelId = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision) SetModelName(v string) *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision {
	s.ModelName = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision) SetPrecision(v float32) *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision {
	s.Precision = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision) SetStatus(v int32) *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision {
	s.Status = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision) SetTaskId(v string) *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision {
	s.TaskId = &v
	return s
}

type ListPrecisionTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPrecisionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrecisionTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPrecisionTaskResponse) GoString() string {
	return s.String()
}

func (s *ListPrecisionTaskResponse) SetHeaders(v map[string]*string) *ListPrecisionTaskResponse {
	s.Headers = v
	return s
}

func (s *ListPrecisionTaskResponse) SetStatusCode(v int32) *ListPrecisionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrecisionTaskResponse) SetBody(v *ListPrecisionTaskResponseBody) *ListPrecisionTaskResponse {
	s.Body = v
	return s
}

type ListQualityCheckSchemeRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s ListQualityCheckSchemeRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQualityCheckSchemeRequest) GoString() string {
	return s.String()
}

func (s *ListQualityCheckSchemeRequest) SetBaseMeAgentId(v int64) *ListQualityCheckSchemeRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *ListQualityCheckSchemeRequest) SetJsonStr(v string) *ListQualityCheckSchemeRequest {
	s.JsonStr = &v
	return s
}

type ListQualityCheckSchemeResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 22
	Count *int32                                    `json:"Count,omitempty" xml:"Count,omitempty"`
	Data  []*ListQualityCheckSchemeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 82C91484-B2D5-4D2A-A21F-A6D73F4***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// XXX
	ResultCountId *string `json:"ResultCountId,omitempty" xml:"ResultCountId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListQualityCheckSchemeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQualityCheckSchemeResponseBody) GoString() string {
	return s.String()
}

func (s *ListQualityCheckSchemeResponseBody) SetCode(v string) *ListQualityCheckSchemeResponseBody {
	s.Code = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBody) SetCount(v int32) *ListQualityCheckSchemeResponseBody {
	s.Count = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBody) SetData(v []*ListQualityCheckSchemeResponseBodyData) *ListQualityCheckSchemeResponseBody {
	s.Data = v
	return s
}

func (s *ListQualityCheckSchemeResponseBody) SetMessage(v string) *ListQualityCheckSchemeResponseBody {
	s.Message = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBody) SetPageNumber(v int32) *ListQualityCheckSchemeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBody) SetPageSize(v int32) *ListQualityCheckSchemeResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBody) SetRequestId(v string) *ListQualityCheckSchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBody) SetResultCountId(v string) *ListQualityCheckSchemeResponseBody {
	s.ResultCountId = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBody) SetSuccess(v bool) *ListQualityCheckSchemeResponseBody {
	s.Success = &v
	return s
}

type ListQualityCheckSchemeResponseBodyData struct {
	// example:
	//
	// 2022-05-10T09:34Z
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	// example:
	//
	// 1
	DataType    *int32  `json:"DataType,omitempty" xml:"DataType,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// test
	Name                *string                                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	RuleList            []*ListQualityCheckSchemeResponseBodyDataRuleList            `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
	SchemeCheckTypeList []*ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList `json:"SchemeCheckTypeList,omitempty" xml:"SchemeCheckTypeList,omitempty" type:"Repeated"`
	// example:
	//
	// 112**
	SchemeId *int64 `json:"SchemeId,omitempty" xml:"SchemeId,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1
	TemplateType *int32 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 2022-05-10T10:34Z
	UpdateTime     *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateUserName *string `json:"UpdateUserName,omitempty" xml:"UpdateUserName,omitempty"`
	// example:
	//
	// 1
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListQualityCheckSchemeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListQualityCheckSchemeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListQualityCheckSchemeResponseBodyData) SetCreateTime(v string) *ListQualityCheckSchemeResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetCreateUserName(v string) *ListQualityCheckSchemeResponseBodyData {
	s.CreateUserName = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetDataType(v int32) *ListQualityCheckSchemeResponseBodyData {
	s.DataType = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetDescription(v string) *ListQualityCheckSchemeResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetName(v string) *ListQualityCheckSchemeResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetRuleList(v []*ListQualityCheckSchemeResponseBodyDataRuleList) *ListQualityCheckSchemeResponseBodyData {
	s.RuleList = v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetSchemeCheckTypeList(v []*ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) *ListQualityCheckSchemeResponseBodyData {
	s.SchemeCheckTypeList = v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetSchemeId(v int64) *ListQualityCheckSchemeResponseBodyData {
	s.SchemeId = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetStatus(v int32) *ListQualityCheckSchemeResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetTemplateType(v int32) *ListQualityCheckSchemeResponseBodyData {
	s.TemplateType = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetType(v int32) *ListQualityCheckSchemeResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetUpdateTime(v string) *ListQualityCheckSchemeResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetUpdateUserName(v string) *ListQualityCheckSchemeResponseBodyData {
	s.UpdateUserName = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetVersion(v int64) *ListQualityCheckSchemeResponseBodyData {
	s.Version = &v
	return s
}

type ListQualityCheckSchemeResponseBodyDataRuleList struct {
	Rules []*ListQualityCheckSchemeResponseBodyDataRuleListRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s ListQualityCheckSchemeResponseBodyDataRuleList) String() string {
	return tea.Prettify(s)
}

func (s ListQualityCheckSchemeResponseBodyDataRuleList) GoString() string {
	return s.String()
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleList) SetRules(v []*ListQualityCheckSchemeResponseBodyDataRuleListRules) *ListQualityCheckSchemeResponseBodyDataRuleList {
	s.Rules = v
	return s
}

type ListQualityCheckSchemeResponseBodyDataRuleListRules struct {
	// example:
	//
	// 1
	CheckType *int32  `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 12
	Rid *int64 `json:"Rid,omitempty" xml:"Rid,omitempty"`
	// example:
	//
	// 1
	RuleScoreType *int32 `json:"RuleScoreType,omitempty" xml:"RuleScoreType,omitempty"`
	// example:
	//
	// 2
	ScoreNum *int32 `json:"ScoreNum,omitempty" xml:"ScoreNum,omitempty"`
	// example:
	//
	// 0
	ScoreNumType *int32 `json:"ScoreNumType,omitempty" xml:"ScoreNumType,omitempty"`
	// example:
	//
	// 1
	ScoreType *int32 `json:"ScoreType,omitempty" xml:"ScoreType,omitempty"`
	// example:
	//
	// 10
	TargetType *int32 `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListQualityCheckSchemeResponseBodyDataRuleListRules) String() string {
	return tea.Prettify(s)
}

func (s ListQualityCheckSchemeResponseBodyDataRuleListRules) GoString() string {
	return s.String()
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleListRules) SetCheckType(v int32) *ListQualityCheckSchemeResponseBodyDataRuleListRules {
	s.CheckType = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleListRules) SetName(v string) *ListQualityCheckSchemeResponseBodyDataRuleListRules {
	s.Name = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleListRules) SetRid(v int64) *ListQualityCheckSchemeResponseBodyDataRuleListRules {
	s.Rid = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleListRules) SetRuleScoreType(v int32) *ListQualityCheckSchemeResponseBodyDataRuleListRules {
	s.RuleScoreType = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleListRules) SetScoreNum(v int32) *ListQualityCheckSchemeResponseBodyDataRuleListRules {
	s.ScoreNum = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleListRules) SetScoreNumType(v int32) *ListQualityCheckSchemeResponseBodyDataRuleListRules {
	s.ScoreNumType = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleListRules) SetScoreType(v int32) *ListQualityCheckSchemeResponseBodyDataRuleListRules {
	s.ScoreType = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleListRules) SetTargetType(v int32) *ListQualityCheckSchemeResponseBodyDataRuleListRules {
	s.TargetType = &v
	return s
}

type ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList struct {
	CheckName *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	// example:
	//
	// 1
	CheckType *int32 `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// example:
	//
	// 1
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// 20
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// 10
	TargetType *int32 `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) String() string {
	return tea.Prettify(s)
}

func (s ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) GoString() string {
	return s.String()
}

func (s *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) SetCheckName(v string) *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList {
	s.CheckName = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) SetCheckType(v int32) *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList {
	s.CheckType = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) SetEnable(v int32) *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList {
	s.Enable = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) SetScore(v int32) *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList {
	s.Score = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) SetTargetType(v int32) *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList {
	s.TargetType = &v
	return s
}

type ListQualityCheckSchemeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQualityCheckSchemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQualityCheckSchemeResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQualityCheckSchemeResponse) GoString() string {
	return s.String()
}

func (s *ListQualityCheckSchemeResponse) SetHeaders(v map[string]*string) *ListQualityCheckSchemeResponse {
	s.Headers = v
	return s
}

func (s *ListQualityCheckSchemeResponse) SetStatusCode(v int32) *ListQualityCheckSchemeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQualityCheckSchemeResponse) SetBody(v *ListQualityCheckSchemeResponseBody) *ListQualityCheckSchemeResponse {
	s.Body = v
	return s
}

type ListRulesRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"pageNumber":1,"pageSize":10}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s ListRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRulesRequest) GoString() string {
	return s.String()
}

func (s *ListRulesRequest) SetBaseMeAgentId(v int64) *ListRulesRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *ListRulesRequest) SetJsonStr(v string) *ListRulesRequest {
	s.JsonStr = &v
	return s
}

type ListRulesResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 20
	Count *int32                       `json:"Count,omitempty" xml:"Count,omitempty"`
	Data  []*ListRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 82C91484-B2D5-4D2A-A21F-A6D73F**
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBody) SetCode(v string) *ListRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListRulesResponseBody) SetCount(v int32) *ListRulesResponseBody {
	s.Count = &v
	return s
}

func (s *ListRulesResponseBody) SetData(v []*ListRulesResponseBodyData) *ListRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListRulesResponseBody) SetMessage(v string) *ListRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListRulesResponseBody) SetPageNumber(v int32) *ListRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRulesResponseBody) SetPageSize(v int32) *ListRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRulesResponseBody) SetRequestId(v string) *ListRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRulesResponseBody) SetSuccess(v bool) *ListRulesResponseBody {
	s.Success = &v
	return s
}

type ListRulesResponseBodyData struct {
	BusinessCategoryNameList []*string `json:"BusinessCategoryNameList,omitempty" xml:"BusinessCategoryNameList,omitempty" type:"Repeated"`
	Comments                 *string   `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// example:
	//
	// 2020-04-20T20:10Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1234567
	Rid *int64 `json:"Rid,omitempty" xml:"Rid,omitempty"`
	// example:
	//
	// 1
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// example:
	//
	// 1
	Type     *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s ListRulesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyData) SetBusinessCategoryNameList(v []*string) *ListRulesResponseBodyData {
	s.BusinessCategoryNameList = v
	return s
}

func (s *ListRulesResponseBodyData) SetComments(v string) *ListRulesResponseBodyData {
	s.Comments = &v
	return s
}

func (s *ListRulesResponseBodyData) SetCreateTime(v string) *ListRulesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListRulesResponseBodyData) SetName(v string) *ListRulesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListRulesResponseBodyData) SetRid(v int64) *ListRulesResponseBodyData {
	s.Rid = &v
	return s
}

func (s *ListRulesResponseBodyData) SetRuleType(v int32) *ListRulesResponseBodyData {
	s.RuleType = &v
	return s
}

func (s *ListRulesResponseBodyData) SetType(v int32) *ListRulesResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListRulesResponseBodyData) SetTypeName(v string) *ListRulesResponseBodyData {
	s.TypeName = &v
	return s
}

type ListRulesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponse) GoString() string {
	return s.String()
}

func (s *ListRulesResponse) SetHeaders(v map[string]*string) *ListRulesResponse {
	s.Headers = v
	return s
}

func (s *ListRulesResponse) SetStatusCode(v int32) *ListRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRulesResponse) SetBody(v *ListRulesResponseBody) *ListRulesResponse {
	s.Body = v
	return s
}

type ListRulesV4Request struct {
	// baseMeAgentId
	BaseMeAgentId *int64  `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	BusinessName  *string `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	BusinessRange *int32  `json:"BusinessRange,omitempty" xml:"BusinessRange,omitempty"`
	CategoryName  *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// example:
	//
	// false
	CountTotal *bool `json:"CountTotal,omitempty" xml:"CountTotal,omitempty"`
	// example:
	//
	// 1
	CreateEmpid *string `json:"CreateEmpid,omitempty" xml:"CreateEmpid,omitempty"`
	// example:
	//
	// 1
	CreateUserId *int64 `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 2021-11-29 19:11:09
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	LastUpdateEmpid *string `json:"LastUpdateEmpid,omitempty" xml:"LastUpdateEmpid,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize     *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequireInfos []*string `json:"RequireInfos,omitempty" xml:"RequireInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 895EAD5312634F5AA708E3B3FA79662E
	Rid *int64 `json:"Rid,omitempty" xml:"Rid,omitempty"`
	// example:
	//
	// xx
	RuleIdOrRuleName *string `json:"RuleIdOrRuleName,omitempty" xml:"RuleIdOrRuleName,omitempty"`
	// example:
	//
	// 1
	RuleScoreSingleType *int32 `json:"RuleScoreSingleType,omitempty" xml:"RuleScoreSingleType,omitempty"`
	// example:
	//
	// 1
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// example:
	//
	// 1000000090
	SchemeId *int64 `json:"SchemeId,omitempty" xml:"SchemeId,omitempty"`
	// example:
	//
	// 0
	SourceType *int32 `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// 2021-11-29 18:11:09
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1
	Type     *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
	// example:
	//
	// 2021-11-29 18:11:09
	UpdateEndTime *string `json:"UpdateEndTime,omitempty" xml:"UpdateEndTime,omitempty"`
	// example:
	//
	// 2021-11-29 16:11:09
	UpdateStartTime *string `json:"UpdateStartTime,omitempty" xml:"UpdateStartTime,omitempty"`
	// example:
	//
	// 1
	UpdateUserId *int64 `json:"UpdateUserId,omitempty" xml:"UpdateUserId,omitempty"`
}

func (s ListRulesV4Request) String() string {
	return tea.Prettify(s)
}

func (s ListRulesV4Request) GoString() string {
	return s.String()
}

func (s *ListRulesV4Request) SetBaseMeAgentId(v int64) *ListRulesV4Request {
	s.BaseMeAgentId = &v
	return s
}

func (s *ListRulesV4Request) SetBusinessName(v string) *ListRulesV4Request {
	s.BusinessName = &v
	return s
}

func (s *ListRulesV4Request) SetBusinessRange(v int32) *ListRulesV4Request {
	s.BusinessRange = &v
	return s
}

func (s *ListRulesV4Request) SetCategoryName(v string) *ListRulesV4Request {
	s.CategoryName = &v
	return s
}

func (s *ListRulesV4Request) SetCountTotal(v bool) *ListRulesV4Request {
	s.CountTotal = &v
	return s
}

func (s *ListRulesV4Request) SetCreateEmpid(v string) *ListRulesV4Request {
	s.CreateEmpid = &v
	return s
}

func (s *ListRulesV4Request) SetCreateUserId(v int64) *ListRulesV4Request {
	s.CreateUserId = &v
	return s
}

func (s *ListRulesV4Request) SetCurrentPage(v int32) *ListRulesV4Request {
	s.CurrentPage = &v
	return s
}

func (s *ListRulesV4Request) SetEndTime(v string) *ListRulesV4Request {
	s.EndTime = &v
	return s
}

func (s *ListRulesV4Request) SetLastUpdateEmpid(v string) *ListRulesV4Request {
	s.LastUpdateEmpid = &v
	return s
}

func (s *ListRulesV4Request) SetPageNumber(v int32) *ListRulesV4Request {
	s.PageNumber = &v
	return s
}

func (s *ListRulesV4Request) SetPageSize(v int32) *ListRulesV4Request {
	s.PageSize = &v
	return s
}

func (s *ListRulesV4Request) SetRequireInfos(v []*string) *ListRulesV4Request {
	s.RequireInfos = v
	return s
}

func (s *ListRulesV4Request) SetRid(v int64) *ListRulesV4Request {
	s.Rid = &v
	return s
}

func (s *ListRulesV4Request) SetRuleIdOrRuleName(v string) *ListRulesV4Request {
	s.RuleIdOrRuleName = &v
	return s
}

func (s *ListRulesV4Request) SetRuleScoreSingleType(v int32) *ListRulesV4Request {
	s.RuleScoreSingleType = &v
	return s
}

func (s *ListRulesV4Request) SetRuleType(v int32) *ListRulesV4Request {
	s.RuleType = &v
	return s
}

func (s *ListRulesV4Request) SetSchemeId(v int64) *ListRulesV4Request {
	s.SchemeId = &v
	return s
}

func (s *ListRulesV4Request) SetSourceType(v int32) *ListRulesV4Request {
	s.SourceType = &v
	return s
}

func (s *ListRulesV4Request) SetStartTime(v string) *ListRulesV4Request {
	s.StartTime = &v
	return s
}

func (s *ListRulesV4Request) SetStatus(v int32) *ListRulesV4Request {
	s.Status = &v
	return s
}

func (s *ListRulesV4Request) SetType(v int32) *ListRulesV4Request {
	s.Type = &v
	return s
}

func (s *ListRulesV4Request) SetTypeName(v string) *ListRulesV4Request {
	s.TypeName = &v
	return s
}

func (s *ListRulesV4Request) SetUpdateEndTime(v string) *ListRulesV4Request {
	s.UpdateEndTime = &v
	return s
}

func (s *ListRulesV4Request) SetUpdateStartTime(v string) *ListRulesV4Request {
	s.UpdateStartTime = &v
	return s
}

func (s *ListRulesV4Request) SetUpdateUserId(v int64) *ListRulesV4Request {
	s.UpdateUserId = &v
	return s
}

type ListRulesV4ResponseBody struct {
	BusinessType *int32 `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 219
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 10
	CurrentPage *int32           `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Data        []*RuleCountInfo `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages []*string `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 96138D8D-8D26-4E41-BFF4-77AED1088BBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 219
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRulesV4ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRulesV4ResponseBody) GoString() string {
	return s.String()
}

func (s *ListRulesV4ResponseBody) SetBusinessType(v int32) *ListRulesV4ResponseBody {
	s.BusinessType = &v
	return s
}

func (s *ListRulesV4ResponseBody) SetCode(v string) *ListRulesV4ResponseBody {
	s.Code = &v
	return s
}

func (s *ListRulesV4ResponseBody) SetCount(v int32) *ListRulesV4ResponseBody {
	s.Count = &v
	return s
}

func (s *ListRulesV4ResponseBody) SetCurrentPage(v int32) *ListRulesV4ResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListRulesV4ResponseBody) SetData(v []*RuleCountInfo) *ListRulesV4ResponseBody {
	s.Data = v
	return s
}

func (s *ListRulesV4ResponseBody) SetHttpStatusCode(v int32) *ListRulesV4ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListRulesV4ResponseBody) SetMessage(v string) *ListRulesV4ResponseBody {
	s.Message = &v
	return s
}

func (s *ListRulesV4ResponseBody) SetMessages(v []*string) *ListRulesV4ResponseBody {
	s.Messages = v
	return s
}

func (s *ListRulesV4ResponseBody) SetPageNumber(v int32) *ListRulesV4ResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRulesV4ResponseBody) SetPageSize(v int32) *ListRulesV4ResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRulesV4ResponseBody) SetRequestId(v string) *ListRulesV4ResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRulesV4ResponseBody) SetSuccess(v bool) *ListRulesV4ResponseBody {
	s.Success = &v
	return s
}

func (s *ListRulesV4ResponseBody) SetTotalCount(v int32) *ListRulesV4ResponseBody {
	s.TotalCount = &v
	return s
}

type ListRulesV4Response struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRulesV4ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRulesV4Response) String() string {
	return tea.Prettify(s)
}

func (s ListRulesV4Response) GoString() string {
	return s.String()
}

func (s *ListRulesV4Response) SetHeaders(v map[string]*string) *ListRulesV4Response {
	s.Headers = v
	return s
}

func (s *ListRulesV4Response) SetStatusCode(v int32) *ListRulesV4Response {
	s.StatusCode = &v
	return s
}

func (s *ListRulesV4Response) SetBody(v *ListRulesV4ResponseBody) *ListRulesV4Response {
	s.Body = v
	return s
}

type ListSchemeTaskConfigRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// {"pageNumber":1,"pageSize":10,"sourceDataType":"1"}
	JsonStr *string `json:"jsonStr,omitempty" xml:"jsonStr,omitempty"`
}

func (s ListSchemeTaskConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSchemeTaskConfigRequest) GoString() string {
	return s.String()
}

func (s *ListSchemeTaskConfigRequest) SetBaseMeAgentId(v int64) *ListSchemeTaskConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *ListSchemeTaskConfigRequest) SetJsonStr(v string) *ListSchemeTaskConfigRequest {
	s.JsonStr = &v
	return s
}

type ListSchemeTaskConfigResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 22
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32                                `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Data        *ListSchemeTaskConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// xxx
	LastDataId *string `json:"LastDataId,omitempty" xml:"LastDataId,omitempty"`
	// example:
	//
	// successful
	Message  *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *ListSchemeTaskConfigResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 4B0A8DCD-0DDF-5E80-8B9C-0A2F68B3403B
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCountId *string `json:"ResultCountId,omitempty" xml:"ResultCountId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSchemeTaskConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSchemeTaskConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListSchemeTaskConfigResponseBody) SetCode(v string) *ListSchemeTaskConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBody) SetCount(v int32) *ListSchemeTaskConfigResponseBody {
	s.Count = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBody) SetCurrentPage(v int32) *ListSchemeTaskConfigResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBody) SetData(v *ListSchemeTaskConfigResponseBodyData) *ListSchemeTaskConfigResponseBody {
	s.Data = v
	return s
}

func (s *ListSchemeTaskConfigResponseBody) SetHttpStatusCode(v int32) *ListSchemeTaskConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBody) SetLastDataId(v string) *ListSchemeTaskConfigResponseBody {
	s.LastDataId = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBody) SetMessage(v string) *ListSchemeTaskConfigResponseBody {
	s.Message = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBody) SetMessages(v *ListSchemeTaskConfigResponseBodyMessages) *ListSchemeTaskConfigResponseBody {
	s.Messages = v
	return s
}

func (s *ListSchemeTaskConfigResponseBody) SetPageNumber(v int32) *ListSchemeTaskConfigResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBody) SetPageSize(v int32) *ListSchemeTaskConfigResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBody) SetRequestId(v string) *ListSchemeTaskConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBody) SetResultCountId(v string) *ListSchemeTaskConfigResponseBody {
	s.ResultCountId = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBody) SetSuccess(v bool) *ListSchemeTaskConfigResponseBody {
	s.Success = &v
	return s
}

type ListSchemeTaskConfigResponseBodyData struct {
	Data []*ListSchemeTaskConfigResponseBodyDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s ListSchemeTaskConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListSchemeTaskConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSchemeTaskConfigResponseBodyData) SetData(v []*ListSchemeTaskConfigResponseBodyDataData) *ListSchemeTaskConfigResponseBodyData {
	s.Data = v
	return s
}

type ListSchemeTaskConfigResponseBodyDataData struct {
	// example:
	//
	// 2
	AsrTaskPriority *int32 `json:"AsrTaskPriority,omitempty" xml:"AsrTaskPriority,omitempty"`
	// example:
	//
	// 0
	AssignType *int32 `json:"AssignType,omitempty" xml:"AssignType,omitempty"`
	// example:
	//
	// 1650418039000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1
	CreateUser *int64                                              `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	DataConfig *ListSchemeTaskConfigResponseBodyDataDataDataConfig `json:"DataConfig,omitempty" xml:"DataConfig,omitempty" type:"Struct"`
	// example:
	//
	// 100
	FinishRate *float64 `json:"FinishRate,omitempty" xml:"FinishRate,omitempty"`
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 0
	ManualReview *int32 `json:"ManualReview,omitempty" xml:"ManualReview,omitempty"`
	// example:
	//
	// cdae396590b*****ec40f3476e274fc
	ModeCustomizationId *string `json:"ModeCustomizationId,omitempty" xml:"ModeCustomizationId,omitempty"`
	ModelName           *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0
	NumberExecuting *int32 `json:"NumberExecuting,omitempty" xml:"NumberExecuting,omitempty"`
	// example:
	//
	// 0
	NumberFail *int32 `json:"NumberFail,omitempty" xml:"NumberFail,omitempty"`
	// example:
	//
	// 1000
	NumberSuccess *int32 `json:"NumberSuccess,omitempty" xml:"NumberSuccess,omitempty"`
	// example:
	//
	// 1000
	NumberSum    *int32                                                `json:"NumberSum,omitempty" xml:"NumberSum,omitempty"`
	SchemeIdList *ListSchemeTaskConfigResponseBodyDataDataSchemeIdList `json:"SchemeIdList,omitempty" xml:"SchemeIdList,omitempty" type:"Struct"`
	SchemeList   *ListSchemeTaskConfigResponseBodyDataDataSchemeList   `json:"SchemeList,omitempty" xml:"SchemeList,omitempty" type:"Struct"`
	// example:
	//
	// 123
	SchemeTaskConfigId *int64 `json:"SchemeTaskConfigId,omitempty" xml:"SchemeTaskConfigId,omitempty"`
	// example:
	//
	// 2
	SourceDataType *int32 `json:"SourceDataType,omitempty" xml:"SourceDataType,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 1650418039000
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 1
	UpdateUser *int64 `json:"UpdateUser,omitempty" xml:"UpdateUser,omitempty"`
	// example:
	//
	// 1
	UserGroup *string `json:"UserGroup,omitempty" xml:"UserGroup,omitempty"`
	// example:
	//
	// 9f90b3efa2****0a49acec226eafc17
	VocabId   *string `json:"VocabId,omitempty" xml:"VocabId,omitempty"`
	VocabName *string `json:"VocabName,omitempty" xml:"VocabName,omitempty"`
}

func (s ListSchemeTaskConfigResponseBodyDataData) String() string {
	return tea.Prettify(s)
}

func (s ListSchemeTaskConfigResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetAsrTaskPriority(v int32) *ListSchemeTaskConfigResponseBodyDataData {
	s.AsrTaskPriority = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetAssignType(v int32) *ListSchemeTaskConfigResponseBodyDataData {
	s.AssignType = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetCreateTime(v string) *ListSchemeTaskConfigResponseBodyDataData {
	s.CreateTime = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetCreateUser(v int64) *ListSchemeTaskConfigResponseBodyDataData {
	s.CreateUser = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetDataConfig(v *ListSchemeTaskConfigResponseBodyDataDataDataConfig) *ListSchemeTaskConfigResponseBodyDataData {
	s.DataConfig = v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetFinishRate(v float64) *ListSchemeTaskConfigResponseBodyDataData {
	s.FinishRate = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetId(v int64) *ListSchemeTaskConfigResponseBodyDataData {
	s.Id = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetManualReview(v int32) *ListSchemeTaskConfigResponseBodyDataData {
	s.ManualReview = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetModeCustomizationId(v string) *ListSchemeTaskConfigResponseBodyDataData {
	s.ModeCustomizationId = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetModelName(v string) *ListSchemeTaskConfigResponseBodyDataData {
	s.ModelName = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetName(v string) *ListSchemeTaskConfigResponseBodyDataData {
	s.Name = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetNumberExecuting(v int32) *ListSchemeTaskConfigResponseBodyDataData {
	s.NumberExecuting = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetNumberFail(v int32) *ListSchemeTaskConfigResponseBodyDataData {
	s.NumberFail = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetNumberSuccess(v int32) *ListSchemeTaskConfigResponseBodyDataData {
	s.NumberSuccess = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetNumberSum(v int32) *ListSchemeTaskConfigResponseBodyDataData {
	s.NumberSum = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetSchemeIdList(v *ListSchemeTaskConfigResponseBodyDataDataSchemeIdList) *ListSchemeTaskConfigResponseBodyDataData {
	s.SchemeIdList = v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetSchemeList(v *ListSchemeTaskConfigResponseBodyDataDataSchemeList) *ListSchemeTaskConfigResponseBodyDataData {
	s.SchemeList = v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetSchemeTaskConfigId(v int64) *ListSchemeTaskConfigResponseBodyDataData {
	s.SchemeTaskConfigId = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetSourceDataType(v int32) *ListSchemeTaskConfigResponseBodyDataData {
	s.SourceDataType = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetStatus(v int32) *ListSchemeTaskConfigResponseBodyDataData {
	s.Status = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetType(v int32) *ListSchemeTaskConfigResponseBodyDataData {
	s.Type = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetUpdateTime(v string) *ListSchemeTaskConfigResponseBodyDataData {
	s.UpdateTime = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetUpdateUser(v int64) *ListSchemeTaskConfigResponseBodyDataData {
	s.UpdateUser = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetUserGroup(v string) *ListSchemeTaskConfigResponseBodyDataData {
	s.UserGroup = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetVocabId(v string) *ListSchemeTaskConfigResponseBodyDataData {
	s.VocabId = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetVocabName(v string) *ListSchemeTaskConfigResponseBodyDataData {
	s.VocabName = &v
	return s
}

type ListSchemeTaskConfigResponseBodyDataDataDataConfig struct {
	AssignConfigs *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigs `json:"AssignConfigs,omitempty" xml:"AssignConfigs,omitempty" type:"Struct"`
	// example:
	//
	// []
	DataSets *string `json:"DataSets,omitempty" xml:"DataSets,omitempty"`
	// example:
	//
	// 0
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// {}
	ResultParam *string `json:"ResultParam,omitempty" xml:"ResultParam,omitempty"`
}

func (s ListSchemeTaskConfigResponseBodyDataDataDataConfig) String() string {
	return tea.Prettify(s)
}

func (s ListSchemeTaskConfigResponseBodyDataDataDataConfig) GoString() string {
	return s.String()
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfig) SetAssignConfigs(v *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigs) *ListSchemeTaskConfigResponseBodyDataDataDataConfig {
	s.AssignConfigs = v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfig) SetDataSets(v string) *ListSchemeTaskConfigResponseBodyDataDataDataConfig {
	s.DataSets = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfig) SetIndex(v int32) *ListSchemeTaskConfigResponseBodyDataDataDataConfig {
	s.Index = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfig) SetResultParam(v string) *ListSchemeTaskConfigResponseBodyDataDataDataConfig {
	s.ResultParam = &v
	return s
}

type ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigs struct {
	AssignConfig []*ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfig `json:"AssignConfig,omitempty" xml:"AssignConfig,omitempty" type:"Repeated"`
}

func (s ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigs) String() string {
	return tea.Prettify(s)
}

func (s ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigs) GoString() string {
	return s.String()
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigs) SetAssignConfig(v []*ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfig) *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigs {
	s.AssignConfig = v
	return s
}

type ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfig struct {
	AssignConfigContests *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContests `json:"AssignConfigContests,omitempty" xml:"AssignConfigContests,omitempty" type:"Struct"`
}

func (s ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfig) String() string {
	return tea.Prettify(s)
}

func (s ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfig) GoString() string {
	return s.String()
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfig) SetAssignConfigContests(v *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContests) *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfig {
	s.AssignConfigContests = v
	return s
}

type ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContests struct {
	AssignConfigContest []*ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest `json:"AssignConfigContest,omitempty" xml:"AssignConfigContest,omitempty" type:"Repeated"`
}

func (s ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContests) String() string {
	return tea.Prettify(s)
}

func (s ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContests) GoString() string {
	return s.String()
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContests) SetAssignConfigContest(v []*ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest) *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContests {
	s.AssignConfigContest = v
	return s
}

type ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest struct {
	// example:
	//
	// 3
	DataType   *int32                                                                                                                        `json:"DataType,omitempty" xml:"DataType,omitempty"`
	ListObject *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContestListObject `json:"ListObject,omitempty" xml:"ListObject,omitempty" type:"Struct"`
	// example:
	//
	// callStartTime
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 4
	Symbol *int32 `json:"Symbol,omitempty" xml:"Symbol,omitempty"`
	// example:
	//
	// {\\"start\\":\\"2022-09-01 00:00:00\\",\\"end\\":\\"2022-09-30 00:00:00\\"}
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest) String() string {
	return tea.Prettify(s)
}

func (s ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest) GoString() string {
	return s.String()
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest) SetDataType(v int32) *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest {
	s.DataType = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest) SetListObject(v *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContestListObject) *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest {
	s.ListObject = v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest) SetName(v string) *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest {
	s.Name = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest) SetSymbol(v int32) *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest {
	s.Symbol = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest) SetValue(v string) *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest {
	s.Value = &v
	return s
}

type ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContestListObject struct {
	ListObject []interface{} `json:"ListObject,omitempty" xml:"ListObject,omitempty" type:"Repeated"`
}

func (s ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContestListObject) String() string {
	return tea.Prettify(s)
}

func (s ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContestListObject) GoString() string {
	return s.String()
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContestListObject) SetListObject(v []interface{}) *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContestListObject {
	s.ListObject = v
	return s
}

type ListSchemeTaskConfigResponseBodyDataDataSchemeIdList struct {
	SchemeIdList []*int64 `json:"SchemeIdList,omitempty" xml:"SchemeIdList,omitempty" type:"Repeated"`
}

func (s ListSchemeTaskConfigResponseBodyDataDataSchemeIdList) String() string {
	return tea.Prettify(s)
}

func (s ListSchemeTaskConfigResponseBodyDataDataSchemeIdList) GoString() string {
	return s.String()
}

func (s *ListSchemeTaskConfigResponseBodyDataDataSchemeIdList) SetSchemeIdList(v []*int64) *ListSchemeTaskConfigResponseBodyDataDataSchemeIdList {
	s.SchemeIdList = v
	return s
}

type ListSchemeTaskConfigResponseBodyDataDataSchemeList struct {
	SchemeList []*ListSchemeTaskConfigResponseBodyDataDataSchemeListSchemeList `json:"SchemeList,omitempty" xml:"SchemeList,omitempty" type:"Repeated"`
}

func (s ListSchemeTaskConfigResponseBodyDataDataSchemeList) String() string {
	return tea.Prettify(s)
}

func (s ListSchemeTaskConfigResponseBodyDataDataSchemeList) GoString() string {
	return s.String()
}

func (s *ListSchemeTaskConfigResponseBodyDataDataSchemeList) SetSchemeList(v []*ListSchemeTaskConfigResponseBodyDataDataSchemeListSchemeList) *ListSchemeTaskConfigResponseBodyDataDataSchemeList {
	s.SchemeList = v
	return s
}

type ListSchemeTaskConfigResponseBodyDataDataSchemeListSchemeList struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 158
	SchemeId *int64 `json:"SchemeId,omitempty" xml:"SchemeId,omitempty"`
}

func (s ListSchemeTaskConfigResponseBodyDataDataSchemeListSchemeList) String() string {
	return tea.Prettify(s)
}

func (s ListSchemeTaskConfigResponseBodyDataDataSchemeListSchemeList) GoString() string {
	return s.String()
}

func (s *ListSchemeTaskConfigResponseBodyDataDataSchemeListSchemeList) SetName(v string) *ListSchemeTaskConfigResponseBodyDataDataSchemeListSchemeList {
	s.Name = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataDataSchemeListSchemeList) SetSchemeId(v int64) *ListSchemeTaskConfigResponseBodyDataDataSchemeListSchemeList {
	s.SchemeId = &v
	return s
}

type ListSchemeTaskConfigResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s ListSchemeTaskConfigResponseBodyMessages) String() string {
	return tea.Prettify(s)
}

func (s ListSchemeTaskConfigResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *ListSchemeTaskConfigResponseBodyMessages) SetMessage(v []*string) *ListSchemeTaskConfigResponseBodyMessages {
	s.Message = v
	return s
}

type ListSchemeTaskConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSchemeTaskConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSchemeTaskConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSchemeTaskConfigResponse) GoString() string {
	return s.String()
}

func (s *ListSchemeTaskConfigResponse) SetHeaders(v map[string]*string) *ListSchemeTaskConfigResponse {
	s.Headers = v
	return s
}

func (s *ListSchemeTaskConfigResponse) SetStatusCode(v int32) *ListSchemeTaskConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSchemeTaskConfigResponse) SetBody(v *ListSchemeTaskConfigResponseBody) *ListSchemeTaskConfigResponse {
	s.Body = v
	return s
}

type ListSessionGroupRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// {"isSchemeData":1,"pageNumber":1,"pageSize":10,"callStartTime":"2022-09-17 00:00:00","callEndTime":"2022-09-23 23:59:59","schemeTaskConfigId":368}
	JsonStr *string `json:"jsonStr,omitempty" xml:"jsonStr,omitempty"`
}

func (s ListSessionGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSessionGroupRequest) GoString() string {
	return s.String()
}

func (s *ListSessionGroupRequest) SetBaseMeAgentId(v int64) *ListSessionGroupRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *ListSessionGroupRequest) SetJsonStr(v string) *ListSessionGroupRequest {
	s.JsonStr = &v
	return s
}

type ListSessionGroupResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 2228
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32                            `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Data        *ListSessionGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// xxx
	LastDataId *string `json:"LastDataId,omitempty" xml:"LastDataId,omitempty"`
	// example:
	//
	// successful
	Message  *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *ListSessionGroupResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// F190ADE9-619A-447D-84E3-7E241A5C428E
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCountId *string `json:"ResultCountId,omitempty" xml:"ResultCountId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSessionGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSessionGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListSessionGroupResponseBody) SetCode(v string) *ListSessionGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ListSessionGroupResponseBody) SetCount(v int32) *ListSessionGroupResponseBody {
	s.Count = &v
	return s
}

func (s *ListSessionGroupResponseBody) SetCurrentPage(v int32) *ListSessionGroupResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListSessionGroupResponseBody) SetData(v *ListSessionGroupResponseBodyData) *ListSessionGroupResponseBody {
	s.Data = v
	return s
}

func (s *ListSessionGroupResponseBody) SetHttpStatusCode(v int32) *ListSessionGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListSessionGroupResponseBody) SetLastDataId(v string) *ListSessionGroupResponseBody {
	s.LastDataId = &v
	return s
}

func (s *ListSessionGroupResponseBody) SetMessage(v string) *ListSessionGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ListSessionGroupResponseBody) SetMessages(v *ListSessionGroupResponseBodyMessages) *ListSessionGroupResponseBody {
	s.Messages = v
	return s
}

func (s *ListSessionGroupResponseBody) SetPageNumber(v int32) *ListSessionGroupResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSessionGroupResponseBody) SetPageSize(v int32) *ListSessionGroupResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSessionGroupResponseBody) SetRequestId(v string) *ListSessionGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSessionGroupResponseBody) SetResultCountId(v string) *ListSessionGroupResponseBody {
	s.ResultCountId = &v
	return s
}

func (s *ListSessionGroupResponseBody) SetSuccess(v bool) *ListSessionGroupResponseBody {
	s.Success = &v
	return s
}

type ListSessionGroupResponseBodyData struct {
	Data []*ListSessionGroupResponseBodyDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s ListSessionGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListSessionGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSessionGroupResponseBodyData) SetData(v []*ListSessionGroupResponseBodyDataData) *ListSessionGroupResponseBodyData {
	s.Data = v
	return s
}

type ListSessionGroupResponseBodyDataData struct {
	// example:
	//
	// 1
	AssignStatus *int32 `json:"AssignStatus,omitempty" xml:"AssignStatus,omitempty"`
	// example:
	//
	// 2022-09-26 10:09:14
	CallStartTime           *string                                                      `json:"CallStartTime,omitempty" xml:"CallStartTime,omitempty"`
	CallerList              *ListSessionGroupResponseBodyDataDataCallerList              `json:"CallerList,omitempty" xml:"CallerList,omitempty" type:"Struct"`
	CustomerIdList          *ListSessionGroupResponseBodyDataDataCustomerIdList          `json:"CustomerIdList,omitempty" xml:"CustomerIdList,omitempty" type:"Struct"`
	CustomerNameList        *ListSessionGroupResponseBodyDataDataCustomerNameList        `json:"CustomerNameList,omitempty" xml:"CustomerNameList,omitempty" type:"Struct"`
	CustomerServiceIdList   *ListSessionGroupResponseBodyDataDataCustomerServiceIdList   `json:"CustomerServiceIdList,omitempty" xml:"CustomerServiceIdList,omitempty" type:"Struct"`
	CustomerServiceNameList *ListSessionGroupResponseBodyDataDataCustomerServiceNameList `json:"CustomerServiceNameList,omitempty" xml:"CustomerServiceNameList,omitempty" type:"Struct"`
	// example:
	//
	// 1
	HitSessionCount *int32 `json:"HitSessionCount,omitempty" xml:"HitSessionCount,omitempty"`
	// example:
	//
	// 4498420@a_z@93EAADF1-01D3-44BD-8AC9-F57F447EFCE8_1614*****
	LastDataId *string `json:"LastDataId,omitempty" xml:"LastDataId,omitempty"`
	// example:
	//
	// 1
	ReviewStatus *int32                                            `json:"ReviewStatus,omitempty" xml:"ReviewStatus,omitempty"`
	ReviewerList *ListSessionGroupResponseBodyDataDataReviewerList `json:"ReviewerList,omitempty" xml:"ReviewerList,omitempty" type:"Struct"`
	// example:
	//
	// 123
	SchemeTaskConfigId   *int64  `json:"SchemeTaskConfigId,omitempty" xml:"SchemeTaskConfigId,omitempty"`
	SchemeTaskConfigName *string `json:"SchemeTaskConfigName,omitempty" xml:"SchemeTaskConfigName,omitempty"`
	// example:
	//
	// 100
	Score *int64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// 1
	SessionCount *int32 `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	// example:
	//
	// SessionGroupA
	SessionGroupId *string `json:"SessionGroupId,omitempty" xml:"SessionGroupId,omitempty"`
	// example:
	//
	// true
	SessionGroupReviewedOrComplained *bool                                                   `json:"SessionGroupReviewedOrComplained,omitempty" xml:"SessionGroupReviewedOrComplained,omitempty"`
	SkillGroupNameList               *ListSessionGroupResponseBodyDataDataSkillGroupNameList `json:"SkillGroupNameList,omitempty" xml:"SkillGroupNameList,omitempty" type:"Struct"`
}

func (s ListSessionGroupResponseBodyDataData) String() string {
	return tea.Prettify(s)
}

func (s ListSessionGroupResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *ListSessionGroupResponseBodyDataData) SetAssignStatus(v int32) *ListSessionGroupResponseBodyDataData {
	s.AssignStatus = &v
	return s
}

func (s *ListSessionGroupResponseBodyDataData) SetCallStartTime(v string) *ListSessionGroupResponseBodyDataData {
	s.CallStartTime = &v
	return s
}

func (s *ListSessionGroupResponseBodyDataData) SetCallerList(v *ListSessionGroupResponseBodyDataDataCallerList) *ListSessionGroupResponseBodyDataData {
	s.CallerList = v
	return s
}

func (s *ListSessionGroupResponseBodyDataData) SetCustomerIdList(v *ListSessionGroupResponseBodyDataDataCustomerIdList) *ListSessionGroupResponseBodyDataData {
	s.CustomerIdList = v
	return s
}

func (s *ListSessionGroupResponseBodyDataData) SetCustomerNameList(v *ListSessionGroupResponseBodyDataDataCustomerNameList) *ListSessionGroupResponseBodyDataData {
	s.CustomerNameList = v
	return s
}

func (s *ListSessionGroupResponseBodyDataData) SetCustomerServiceIdList(v *ListSessionGroupResponseBodyDataDataCustomerServiceIdList) *ListSessionGroupResponseBodyDataData {
	s.CustomerServiceIdList = v
	return s
}

func (s *ListSessionGroupResponseBodyDataData) SetCustomerServiceNameList(v *ListSessionGroupResponseBodyDataDataCustomerServiceNameList) *ListSessionGroupResponseBodyDataData {
	s.CustomerServiceNameList = v
	return s
}

func (s *ListSessionGroupResponseBodyDataData) SetHitSessionCount(v int32) *ListSessionGroupResponseBodyDataData {
	s.HitSessionCount = &v
	return s
}

func (s *ListSessionGroupResponseBodyDataData) SetLastDataId(v string) *ListSessionGroupResponseBodyDataData {
	s.LastDataId = &v
	return s
}

func (s *ListSessionGroupResponseBodyDataData) SetReviewStatus(v int32) *ListSessionGroupResponseBodyDataData {
	s.ReviewStatus = &v
	return s
}

func (s *ListSessionGroupResponseBodyDataData) SetReviewerList(v *ListSessionGroupResponseBodyDataDataReviewerList) *ListSessionGroupResponseBodyDataData {
	s.ReviewerList = v
	return s
}

func (s *ListSessionGroupResponseBodyDataData) SetSchemeTaskConfigId(v int64) *ListSessionGroupResponseBodyDataData {
	s.SchemeTaskConfigId = &v
	return s
}

func (s *ListSessionGroupResponseBodyDataData) SetSchemeTaskConfigName(v string) *ListSessionGroupResponseBodyDataData {
	s.SchemeTaskConfigName = &v
	return s
}

func (s *ListSessionGroupResponseBodyDataData) SetScore(v int64) *ListSessionGroupResponseBodyDataData {
	s.Score = &v
	return s
}

func (s *ListSessionGroupResponseBodyDataData) SetSessionCount(v int32) *ListSessionGroupResponseBodyDataData {
	s.SessionCount = &v
	return s
}

func (s *ListSessionGroupResponseBodyDataData) SetSessionGroupId(v string) *ListSessionGroupResponseBodyDataData {
	s.SessionGroupId = &v
	return s
}

func (s *ListSessionGroupResponseBodyDataData) SetSessionGroupReviewedOrComplained(v bool) *ListSessionGroupResponseBodyDataData {
	s.SessionGroupReviewedOrComplained = &v
	return s
}

func (s *ListSessionGroupResponseBodyDataData) SetSkillGroupNameList(v *ListSessionGroupResponseBodyDataDataSkillGroupNameList) *ListSessionGroupResponseBodyDataData {
	s.SkillGroupNameList = v
	return s
}

type ListSessionGroupResponseBodyDataDataCallerList struct {
	CallerList []*string `json:"CallerList,omitempty" xml:"CallerList,omitempty" type:"Repeated"`
}

func (s ListSessionGroupResponseBodyDataDataCallerList) String() string {
	return tea.Prettify(s)
}

func (s ListSessionGroupResponseBodyDataDataCallerList) GoString() string {
	return s.String()
}

func (s *ListSessionGroupResponseBodyDataDataCallerList) SetCallerList(v []*string) *ListSessionGroupResponseBodyDataDataCallerList {
	s.CallerList = v
	return s
}

type ListSessionGroupResponseBodyDataDataCustomerIdList struct {
	CustomerIdList []*string `json:"CustomerIdList,omitempty" xml:"CustomerIdList,omitempty" type:"Repeated"`
}

func (s ListSessionGroupResponseBodyDataDataCustomerIdList) String() string {
	return tea.Prettify(s)
}

func (s ListSessionGroupResponseBodyDataDataCustomerIdList) GoString() string {
	return s.String()
}

func (s *ListSessionGroupResponseBodyDataDataCustomerIdList) SetCustomerIdList(v []*string) *ListSessionGroupResponseBodyDataDataCustomerIdList {
	s.CustomerIdList = v
	return s
}

type ListSessionGroupResponseBodyDataDataCustomerNameList struct {
	CustomerNameList []*string `json:"CustomerNameList,omitempty" xml:"CustomerNameList,omitempty" type:"Repeated"`
}

func (s ListSessionGroupResponseBodyDataDataCustomerNameList) String() string {
	return tea.Prettify(s)
}

func (s ListSessionGroupResponseBodyDataDataCustomerNameList) GoString() string {
	return s.String()
}

func (s *ListSessionGroupResponseBodyDataDataCustomerNameList) SetCustomerNameList(v []*string) *ListSessionGroupResponseBodyDataDataCustomerNameList {
	s.CustomerNameList = v
	return s
}

type ListSessionGroupResponseBodyDataDataCustomerServiceIdList struct {
	CustomerServiceIdList []*string `json:"CustomerServiceIdList,omitempty" xml:"CustomerServiceIdList,omitempty" type:"Repeated"`
}

func (s ListSessionGroupResponseBodyDataDataCustomerServiceIdList) String() string {
	return tea.Prettify(s)
}

func (s ListSessionGroupResponseBodyDataDataCustomerServiceIdList) GoString() string {
	return s.String()
}

func (s *ListSessionGroupResponseBodyDataDataCustomerServiceIdList) SetCustomerServiceIdList(v []*string) *ListSessionGroupResponseBodyDataDataCustomerServiceIdList {
	s.CustomerServiceIdList = v
	return s
}

type ListSessionGroupResponseBodyDataDataCustomerServiceNameList struct {
	CustomerServiceNameList []*string `json:"CustomerServiceNameList,omitempty" xml:"CustomerServiceNameList,omitempty" type:"Repeated"`
}

func (s ListSessionGroupResponseBodyDataDataCustomerServiceNameList) String() string {
	return tea.Prettify(s)
}

func (s ListSessionGroupResponseBodyDataDataCustomerServiceNameList) GoString() string {
	return s.String()
}

func (s *ListSessionGroupResponseBodyDataDataCustomerServiceNameList) SetCustomerServiceNameList(v []*string) *ListSessionGroupResponseBodyDataDataCustomerServiceNameList {
	s.CustomerServiceNameList = v
	return s
}

type ListSessionGroupResponseBodyDataDataReviewerList struct {
	ReviewerList []*string `json:"ReviewerList,omitempty" xml:"ReviewerList,omitempty" type:"Repeated"`
}

func (s ListSessionGroupResponseBodyDataDataReviewerList) String() string {
	return tea.Prettify(s)
}

func (s ListSessionGroupResponseBodyDataDataReviewerList) GoString() string {
	return s.String()
}

func (s *ListSessionGroupResponseBodyDataDataReviewerList) SetReviewerList(v []*string) *ListSessionGroupResponseBodyDataDataReviewerList {
	s.ReviewerList = v
	return s
}

type ListSessionGroupResponseBodyDataDataSkillGroupNameList struct {
	SkillGroupNameList []*string `json:"SkillGroupNameList,omitempty" xml:"SkillGroupNameList,omitempty" type:"Repeated"`
}

func (s ListSessionGroupResponseBodyDataDataSkillGroupNameList) String() string {
	return tea.Prettify(s)
}

func (s ListSessionGroupResponseBodyDataDataSkillGroupNameList) GoString() string {
	return s.String()
}

func (s *ListSessionGroupResponseBodyDataDataSkillGroupNameList) SetSkillGroupNameList(v []*string) *ListSessionGroupResponseBodyDataDataSkillGroupNameList {
	s.SkillGroupNameList = v
	return s
}

type ListSessionGroupResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s ListSessionGroupResponseBodyMessages) String() string {
	return tea.Prettify(s)
}

func (s ListSessionGroupResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *ListSessionGroupResponseBodyMessages) SetMessage(v []*string) *ListSessionGroupResponseBodyMessages {
	s.Message = v
	return s
}

type ListSessionGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSessionGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSessionGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSessionGroupResponse) GoString() string {
	return s.String()
}

func (s *ListSessionGroupResponse) SetHeaders(v map[string]*string) *ListSessionGroupResponse {
	s.Headers = v
	return s
}

func (s *ListSessionGroupResponse) SetStatusCode(v int32) *ListSessionGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSessionGroupResponse) SetBody(v *ListSessionGroupResponseBody) *ListSessionGroupResponse {
	s.Body = v
	return s
}

type ListSkillGroupConfigRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"pageNumber":1,"pageSize": 1}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s ListSkillGroupConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupConfigRequest) GoString() string {
	return s.String()
}

func (s *ListSkillGroupConfigRequest) SetBaseMeAgentId(v int64) *ListSkillGroupConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *ListSkillGroupConfigRequest) SetJsonStr(v string) *ListSkillGroupConfigRequest {
	s.JsonStr = &v
	return s
}

type ListSkillGroupConfigResponseBody struct {
	// example:
	//
	// 200
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListSkillGroupConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3CEA0495-341B-4482-9AD9-8191EF4***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSkillGroupConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListSkillGroupConfigResponseBody) SetCode(v string) *ListSkillGroupConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ListSkillGroupConfigResponseBody) SetData(v *ListSkillGroupConfigResponseBodyData) *ListSkillGroupConfigResponseBody {
	s.Data = v
	return s
}

func (s *ListSkillGroupConfigResponseBody) SetMessage(v string) *ListSkillGroupConfigResponseBody {
	s.Message = &v
	return s
}

func (s *ListSkillGroupConfigResponseBody) SetRequestId(v string) *ListSkillGroupConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSkillGroupConfigResponseBody) SetSuccess(v bool) *ListSkillGroupConfigResponseBody {
	s.Success = &v
	return s
}

type ListSkillGroupConfigResponseBodyData struct {
	SkillGroupConfig []*ListSkillGroupConfigResponseBodyDataSkillGroupConfig `json:"SkillGroupConfig,omitempty" xml:"SkillGroupConfig,omitempty" type:"Repeated"`
}

func (s ListSkillGroupConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSkillGroupConfigResponseBodyData) SetSkillGroupConfig(v []*ListSkillGroupConfigResponseBodyDataSkillGroupConfig) *ListSkillGroupConfigResponseBodyData {
	s.SkillGroupConfig = v
	return s
}

type ListSkillGroupConfigResponseBodyDataSkillGroupConfig struct {
	// example:
	//
	// 1
	AllContentQualityCheck *int32 `json:"AllContentQualityCheck,omitempty" xml:"AllContentQualityCheck,omitempty"`
	// example:
	//
	// 223
	AllRids     *string                                                          `json:"AllRids,omitempty" xml:"AllRids,omitempty"`
	AllRuleList *ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleList `json:"AllRuleList,omitempty" xml:"AllRuleList,omitempty" type:"Struct"`
	// example:
	//
	// 2020-12-01T15:12Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 221
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 211
	ModelId *int64 `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// example:
	//
	// xxx
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// example:
	//
	// xxx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0
	QualityCheckType *int32 `json:"QualityCheckType,omitempty" xml:"QualityCheckType,omitempty"`
	// example:
	//
	// 2221
	Rid      *string                                                       `json:"Rid,omitempty" xml:"Rid,omitempty"`
	RuleList *ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Struct"`
	// example:
	//
	// true
	ScreenSwitch *bool `json:"ScreenSwitch,omitempty" xml:"ScreenSwitch,omitempty"`
	// example:
	//
	// 0
	SkillGroupFrom *int32 `json:"SkillGroupFrom,omitempty" xml:"SkillGroupFrom,omitempty"`
	// example:
	//
	// 123
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// example:
	//
	// xxx
	SkillGroupName    *string                                                                `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
	SkillGroupScreens *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreens `json:"SkillGroupScreens,omitempty" xml:"SkillGroupScreens,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 2020-12-01T19:28Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 323
	VocabId *int64 `json:"VocabId,omitempty" xml:"VocabId,omitempty"`
	// example:
	//
	// xxx
	VocabName *string `json:"VocabName,omitempty" xml:"VocabName,omitempty"`
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfig) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfig) GoString() string {
	return s.String()
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetAllContentQualityCheck(v int32) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.AllContentQualityCheck = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetAllRids(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.AllRids = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetAllRuleList(v *ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleList) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.AllRuleList = v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetCreateTime(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.CreateTime = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetId(v int64) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.Id = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetInstanceId(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.InstanceId = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetModelId(v int64) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.ModelId = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetModelName(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.ModelName = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetName(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.Name = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetQualityCheckType(v int32) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.QualityCheckType = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetRid(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.Rid = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetRuleList(v *ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleList) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.RuleList = v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetScreenSwitch(v bool) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.ScreenSwitch = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetSkillGroupFrom(v int32) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.SkillGroupFrom = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetSkillGroupId(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.SkillGroupId = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetSkillGroupName(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.SkillGroupName = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetSkillGroupScreens(v *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreens) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.SkillGroupScreens = v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetStatus(v int32) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.Status = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetType(v int32) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.Type = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetUpdateTime(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.UpdateTime = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetVocabId(v int64) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.VocabId = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetVocabName(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.VocabName = &v
	return s
}

type ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleList struct {
	RuleNameInfo []*ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleListRuleNameInfo `json:"RuleNameInfo,omitempty" xml:"RuleNameInfo,omitempty" type:"Repeated"`
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleList) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleList) GoString() string {
	return s.String()
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleList) SetRuleNameInfo(v []*ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleListRuleNameInfo) *ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleList {
	s.RuleNameInfo = v
	return s
}

type ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleListRuleNameInfo struct {
	// example:
	//
	// 221
	Rid *int64 `json:"Rid,omitempty" xml:"Rid,omitempty"`
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleListRuleNameInfo) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleListRuleNameInfo) GoString() string {
	return s.String()
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleListRuleNameInfo) SetRid(v int64) *ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleListRuleNameInfo {
	s.Rid = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleListRuleNameInfo) SetRuleName(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleListRuleNameInfo {
	s.RuleName = &v
	return s
}

type ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleList struct {
	RuleNameInfo []*ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleListRuleNameInfo `json:"RuleNameInfo,omitempty" xml:"RuleNameInfo,omitempty" type:"Repeated"`
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleList) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleList) GoString() string {
	return s.String()
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleList) SetRuleNameInfo(v []*ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleListRuleNameInfo) *ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleList {
	s.RuleNameInfo = v
	return s
}

type ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleListRuleNameInfo struct {
	// example:
	//
	// 2221
	Rid *int64 `json:"Rid,omitempty" xml:"Rid,omitempty"`
	// example:
	//
	// x\\"x\\"x
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleListRuleNameInfo) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleListRuleNameInfo) GoString() string {
	return s.String()
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleListRuleNameInfo) SetRid(v int64) *ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleListRuleNameInfo {
	s.Rid = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleListRuleNameInfo) SetRuleName(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleListRuleNameInfo {
	s.RuleName = &v
	return s
}

type ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreens struct {
	SkillGroupScreen []*ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen `json:"SkillGroupScreen,omitempty" xml:"SkillGroupScreen,omitempty" type:"Repeated"`
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreens) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreens) GoString() string {
	return s.String()
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreens) SetSkillGroupScreen(v []*ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen) *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreens {
	s.SkillGroupScreen = v
	return s
}

type ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen struct {
	// example:
	//
	// 0
	DataType *int32 `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// example:
	//
	// customerName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	Symbol *int32  `json:"Symbol,omitempty" xml:"Symbol,omitempty"`
	Value  *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen) GoString() string {
	return s.String()
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen) SetDataType(v int32) *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen {
	s.DataType = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen) SetName(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen {
	s.Name = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen) SetSymbol(v int32) *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen {
	s.Symbol = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen) SetValue(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen {
	s.Value = &v
	return s
}

type ListSkillGroupConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSkillGroupConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSkillGroupConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupConfigResponse) GoString() string {
	return s.String()
}

func (s *ListSkillGroupConfigResponse) SetHeaders(v map[string]*string) *ListSkillGroupConfigResponse {
	s.Headers = v
	return s
}

func (s *ListSkillGroupConfigResponse) SetStatusCode(v int32) *ListSkillGroupConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSkillGroupConfigResponse) SetBody(v *ListSkillGroupConfigResponseBody) *ListSkillGroupConfigResponse {
	s.Body = v
	return s
}

type ListTaskAssignRulesRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"pageNumber":1,"pageSize":10}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s ListTaskAssignRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTaskAssignRulesRequest) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesRequest) SetBaseMeAgentId(v int64) *ListTaskAssignRulesRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *ListTaskAssignRulesRequest) SetJsonStr(v string) *ListTaskAssignRulesRequest {
	s.JsonStr = &v
	return s
}

type ListTaskAssignRulesResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 23
	Count *int32                               `json:"Count,omitempty" xml:"Count,omitempty"`
	Data  *ListTaskAssignRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24F4CE647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListTaskAssignRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTaskAssignRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBody) SetCode(v string) *ListTaskAssignRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListTaskAssignRulesResponseBody) SetCount(v int32) *ListTaskAssignRulesResponseBody {
	s.Count = &v
	return s
}

func (s *ListTaskAssignRulesResponseBody) SetData(v *ListTaskAssignRulesResponseBodyData) *ListTaskAssignRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListTaskAssignRulesResponseBody) SetMessage(v string) *ListTaskAssignRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListTaskAssignRulesResponseBody) SetPageNumber(v int32) *ListTaskAssignRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTaskAssignRulesResponseBody) SetPageSize(v int32) *ListTaskAssignRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTaskAssignRulesResponseBody) SetRequestId(v string) *ListTaskAssignRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTaskAssignRulesResponseBody) SetSuccess(v bool) *ListTaskAssignRulesResponseBody {
	s.Success = &v
	return s
}

type ListTaskAssignRulesResponseBodyData struct {
	TaskAssignRuleInfo []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo `json:"TaskAssignRuleInfo,omitempty" xml:"TaskAssignRuleInfo,omitempty" type:"Repeated"`
}

func (s ListTaskAssignRulesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyData) SetTaskAssignRuleInfo(v []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) *ListTaskAssignRulesResponseBodyData {
	s.TaskAssignRuleInfo = v
	return s
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo struct {
	Agents *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgents `json:"Agents,omitempty" xml:"Agents,omitempty" type:"Struct"`
	// example:
	//
	// XX
	AgentsStr *string `json:"AgentsStr,omitempty" xml:"AgentsStr,omitempty"`
	// example:
	//
	// 0
	AssignmentType *int32 `json:"AssignmentType,omitempty" xml:"AssignmentType,omitempty"`
	// example:
	//
	// 39600
	CallTimeEnd *int64 `json:"CallTimeEnd,omitempty" xml:"CallTimeEnd,omitempty"`
	// example:
	//
	// 39600
	CallTimeStart *int64 `json:"CallTimeStart,omitempty" xml:"CallTimeStart,omitempty"`
	// example:
	//
	// 1
	CallType *int32 `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// example:
	//
	// 2019-07-12T14:47Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 400
	DurationMax *int32 `json:"DurationMax,omitempty" xml:"DurationMax,omitempty"`
	// example:
	//
	// 100
	DurationMin *int32 `json:"DurationMin,omitempty" xml:"DurationMin,omitempty"`
	// example:
	//
	// 1
	Enabled *int32 `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// 1
	Priority  *int32                                                          `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Reviewers *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewers `json:"Reviewers,omitempty" xml:"Reviewers,omitempty" type:"Struct"`
	// example:
	//
	// 23
	RuleId       *int64                                                             `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName     *string                                                            `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Rules        *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRules        `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
	SamplingMode *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode `json:"SamplingMode,omitempty" xml:"SamplingMode,omitempty" type:"Struct"`
	SkillGroups  *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroups  `json:"SkillGroups,omitempty" xml:"SkillGroups,omitempty" type:"Struct"`
	// example:
	//
	// XX
	SkillGroupsStr *string `json:"SkillGroupsStr,omitempty" xml:"SkillGroupsStr,omitempty"`
	// example:
	//
	// 2019-07-12T14:47Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) String() string {
	return tea.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetAgents(v *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgents) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.Agents = v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetAgentsStr(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.AgentsStr = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetAssignmentType(v int32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.AssignmentType = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetCallTimeEnd(v int64) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.CallTimeEnd = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetCallTimeStart(v int64) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.CallTimeStart = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetCallType(v int32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.CallType = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetCreateTime(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.CreateTime = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetDurationMax(v int32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.DurationMax = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetDurationMin(v int32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.DurationMin = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetEnabled(v int32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.Enabled = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetPriority(v int32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.Priority = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetReviewers(v *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewers) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.Reviewers = v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetRuleId(v int64) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.RuleId = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetRuleName(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.RuleName = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetRules(v *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRules) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.Rules = v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetSamplingMode(v *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.SamplingMode = v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetSkillGroups(v *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroups) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.SkillGroups = v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetSkillGroupsStr(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.SkillGroupsStr = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetUpdateTime(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.UpdateTime = &v
	return s
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgents struct {
	Agent []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgentsAgent `json:"Agent,omitempty" xml:"Agent,omitempty" type:"Repeated"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgents) String() string {
	return tea.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgents) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgents) SetAgent(v []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgentsAgent) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgents {
	s.Agent = v
	return s
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgentsAgent struct {
	// example:
	//
	// 202526561358712105
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// agent
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgentsAgent) String() string {
	return tea.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgentsAgent) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgentsAgent) SetAgentId(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgentsAgent {
	s.AgentId = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgentsAgent) SetAgentName(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgentsAgent {
	s.AgentName = &v
	return s
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewers struct {
	Reviewer []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewersReviewer `json:"Reviewer,omitempty" xml:"Reviewer,omitempty" type:"Repeated"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewers) String() string {
	return tea.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewers) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewers) SetReviewer(v []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewersReviewer) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewers {
	s.Reviewer = v
	return s
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewersReviewer struct {
	// example:
	//
	// 2337235457340978
	ReviewerId   *string `json:"ReviewerId,omitempty" xml:"ReviewerId,omitempty"`
	ReviewerName *string `json:"ReviewerName,omitempty" xml:"ReviewerName,omitempty"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewersReviewer) String() string {
	return tea.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewersReviewer) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewersReviewer) SetReviewerId(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewersReviewer {
	s.ReviewerId = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewersReviewer) SetReviewerName(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewersReviewer {
	s.ReviewerName = &v
	return s
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRules struct {
	RuleBasicInfo []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRulesRuleBasicInfo `json:"RuleBasicInfo,omitempty" xml:"RuleBasicInfo,omitempty" type:"Repeated"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRules) String() string {
	return tea.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRules) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRules) SetRuleBasicInfo(v []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRulesRuleBasicInfo) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRules {
	s.RuleBasicInfo = v
	return s
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRulesRuleBasicInfo struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2312
	Rid *string `json:"Rid,omitempty" xml:"Rid,omitempty"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRulesRuleBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRulesRuleBasicInfo) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRulesRuleBasicInfo) SetName(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRulesRuleBasicInfo {
	s.Name = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRulesRuleBasicInfo) SetRid(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRulesRuleBasicInfo {
	s.Rid = &v
	return s
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode struct {
	// example:
	//
	// 60
	AnyNumberOfDraws *int32 `json:"AnyNumberOfDraws,omitempty" xml:"AnyNumberOfDraws,omitempty"`
	// example:
	//
	// true
	Designated *bool `json:"Designated,omitempty" xml:"Designated,omitempty"`
	// example:
	//
	// 0
	Dimension *int32 `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// example:
	//
	// 30
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// example:
	//
	// 20
	NumberOfDraws *int32 `json:"NumberOfDraws,omitempty" xml:"NumberOfDraws,omitempty"`
	// example:
	//
	// 0.1
	Proportion *float32 `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
	// example:
	//
	// 5
	RandomInspectionNumber *int32                                                                               `json:"RandomInspectionNumber,omitempty" xml:"RandomInspectionNumber,omitempty"`
	SamplingModeAgents     *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgents `json:"SamplingModeAgents,omitempty" xml:"SamplingModeAgents,omitempty" type:"Struct"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) String() string {
	return tea.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) SetAnyNumberOfDraws(v int32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode {
	s.AnyNumberOfDraws = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) SetDesignated(v bool) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode {
	s.Designated = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) SetDimension(v int32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode {
	s.Dimension = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) SetLimit(v int32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode {
	s.Limit = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) SetNumberOfDraws(v int32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode {
	s.NumberOfDraws = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) SetProportion(v float32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode {
	s.Proportion = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) SetRandomInspectionNumber(v int32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode {
	s.RandomInspectionNumber = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) SetSamplingModeAgents(v *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgents) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode {
	s.SamplingModeAgents = v
	return s
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgents struct {
	SamplingModeAgent []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgentsSamplingModeAgent `json:"SamplingModeAgent,omitempty" xml:"SamplingModeAgent,omitempty" type:"Repeated"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgents) String() string {
	return tea.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgents) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgents) SetSamplingModeAgent(v []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgentsSamplingModeAgent) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgents {
	s.SamplingModeAgent = v
	return s
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgentsSamplingModeAgent struct {
	// example:
	//
	// 123
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// zhangsan
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgentsSamplingModeAgent) String() string {
	return tea.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgentsSamplingModeAgent) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgentsSamplingModeAgent) SetAgentId(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgentsSamplingModeAgent {
	s.AgentId = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgentsSamplingModeAgent) SetAgentName(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgentsSamplingModeAgent {
	s.AgentName = &v
	return s
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroups struct {
	SkillGroup []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroupsSkillGroup `json:"SkillGroup,omitempty" xml:"SkillGroup,omitempty" type:"Repeated"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroups) String() string {
	return tea.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroups) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroups) SetSkillGroup(v []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroupsSkillGroup) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroups {
	s.SkillGroup = v
	return s
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroupsSkillGroup struct {
	// example:
	//
	// XXX
	SkillId   *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroupsSkillGroup) String() string {
	return tea.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroupsSkillGroup) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroupsSkillGroup) SetSkillId(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroupsSkillGroup {
	s.SkillId = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroupsSkillGroup) SetSkillName(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroupsSkillGroup {
	s.SkillName = &v
	return s
}

type ListTaskAssignRulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTaskAssignRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTaskAssignRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTaskAssignRulesResponse) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponse) SetHeaders(v map[string]*string) *ListTaskAssignRulesResponse {
	s.Headers = v
	return s
}

func (s *ListTaskAssignRulesResponse) SetStatusCode(v int32) *ListTaskAssignRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTaskAssignRulesResponse) SetBody(v *ListTaskAssignRulesResponseBody) *ListTaskAssignRulesResponse {
	s.Body = v
	return s
}

type ListUsersRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"pageNumber":1,"pageSize":10}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s ListUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) SetBaseMeAgentId(v int64) *ListUsersRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *ListUsersRequest) SetJsonStr(v string) *ListUsersRequest {
	s.JsonStr = &v
	return s
}

type ListUsersResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 12
	Count *int32                     `json:"Count,omitempty" xml:"Count,omitempty"`
	Data  *ListUsersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24F4CE647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) SetCode(v string) *ListUsersResponseBody {
	s.Code = &v
	return s
}

func (s *ListUsersResponseBody) SetCount(v int32) *ListUsersResponseBody {
	s.Count = &v
	return s
}

func (s *ListUsersResponseBody) SetData(v *ListUsersResponseBodyData) *ListUsersResponseBody {
	s.Data = v
	return s
}

func (s *ListUsersResponseBody) SetMessage(v string) *ListUsersResponseBody {
	s.Message = &v
	return s
}

func (s *ListUsersResponseBody) SetPageNumber(v int32) *ListUsersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListUsersResponseBody) SetPageSize(v int32) *ListUsersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetSuccess(v bool) *ListUsersResponseBody {
	s.Success = &v
	return s
}

type ListUsersResponseBodyData struct {
	User []*ListUsersResponseBodyDataUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyData) SetUser(v []*ListUsersResponseBodyDataUser) *ListUsersResponseBodyData {
	s.User = v
	return s
}

type ListUsersResponseBodyDataUser struct {
	// example:
	//
	// 2951869706989****
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// example:
	//
	// 2020-03-11T16:54Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// XXX
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// xxx
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2
	LoginUserType *int32 `json:"LoginUserType,omitempty" xml:"LoginUserType,omitempty"`
	// example:
	//
	// AGENT
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// example:
	//
	// 2020-03-11T16:54Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// xxx
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListUsersResponseBodyDataUser) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyDataUser) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyDataUser) SetAliUid(v string) *ListUsersResponseBodyDataUser {
	s.AliUid = &v
	return s
}

func (s *ListUsersResponseBodyDataUser) SetCreateTime(v string) *ListUsersResponseBodyDataUser {
	s.CreateTime = &v
	return s
}

func (s *ListUsersResponseBodyDataUser) SetDescription(v string) *ListUsersResponseBodyDataUser {
	s.Description = &v
	return s
}

func (s *ListUsersResponseBodyDataUser) SetDisplayName(v string) *ListUsersResponseBodyDataUser {
	s.DisplayName = &v
	return s
}

func (s *ListUsersResponseBodyDataUser) SetId(v int64) *ListUsersResponseBodyDataUser {
	s.Id = &v
	return s
}

func (s *ListUsersResponseBodyDataUser) SetLoginUserType(v int32) *ListUsersResponseBodyDataUser {
	s.LoginUserType = &v
	return s
}

func (s *ListUsersResponseBodyDataUser) SetRoleName(v string) *ListUsersResponseBodyDataUser {
	s.RoleName = &v
	return s
}

func (s *ListUsersResponseBodyDataUser) SetUpdateTime(v string) *ListUsersResponseBodyDataUser {
	s.UpdateTime = &v
	return s
}

func (s *ListUsersResponseBodyDataUser) SetUserName(v string) *ListUsersResponseBodyDataUser {
	s.UserName = &v
	return s
}

type ListUsersResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponse) GoString() string {
	return s.String()
}

func (s *ListUsersResponse) SetHeaders(v map[string]*string) *ListUsersResponse {
	s.Headers = v
	return s
}

func (s *ListUsersResponse) SetStatusCode(v int32) *ListUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsersResponse) SetBody(v *ListUsersResponseBody) *ListUsersResponse {
	s.Body = v
	return s
}

type ListWarningConfigRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"pageNumber":1,"pageSize":10}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s ListWarningConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWarningConfigRequest) GoString() string {
	return s.String()
}

func (s *ListWarningConfigRequest) SetBaseMeAgentId(v int64) *ListWarningConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *ListWarningConfigRequest) SetJsonStr(v string) *ListWarningConfigRequest {
	s.JsonStr = &v
	return s
}

type ListWarningConfigResponseBody struct {
	// example:
	//
	// 200
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListWarningConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 82C91484-B2D5-4D2A-A21F-A6D73F4D55C6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListWarningConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWarningConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListWarningConfigResponseBody) SetCode(v string) *ListWarningConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ListWarningConfigResponseBody) SetData(v *ListWarningConfigResponseBodyData) *ListWarningConfigResponseBody {
	s.Data = v
	return s
}

func (s *ListWarningConfigResponseBody) SetMessage(v string) *ListWarningConfigResponseBody {
	s.Message = &v
	return s
}

func (s *ListWarningConfigResponseBody) SetRequestId(v string) *ListWarningConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWarningConfigResponseBody) SetSuccess(v bool) *ListWarningConfigResponseBody {
	s.Success = &v
	return s
}

type ListWarningConfigResponseBodyData struct {
	WarningConfigInfo []*ListWarningConfigResponseBodyDataWarningConfigInfo `json:"WarningConfigInfo,omitempty" xml:"WarningConfigInfo,omitempty" type:"Repeated"`
}

func (s ListWarningConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListWarningConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWarningConfigResponseBodyData) SetWarningConfigInfo(v []*ListWarningConfigResponseBodyDataWarningConfigInfo) *ListWarningConfigResponseBodyData {
	s.WarningConfigInfo = v
	return s
}

type ListWarningConfigResponseBodyDataWarningConfigInfo struct {
	Channels *ListWarningConfigResponseBodyDataWarningConfigInfoChannels `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Struct"`
	// example:
	//
	// 32
	ConfigId   *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ConfigName *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	// example:
	//
	// 2019-10-29T15:30Z
	CreateTime *string                                                     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	RidList    *ListWarningConfigResponseBodyDataWarningConfigInfoRidList  `json:"RidList,omitempty" xml:"RidList,omitempty" type:"Struct"`
	RuleList   *ListWarningConfigResponseBodyDataWarningConfigInfoRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2019-10-29T17:24Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfo) String() string {
	return tea.Prettify(s)
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfo) GoString() string {
	return s.String()
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfo) SetChannels(v *ListWarningConfigResponseBodyDataWarningConfigInfoChannels) *ListWarningConfigResponseBodyDataWarningConfigInfo {
	s.Channels = v
	return s
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfo) SetConfigId(v int64) *ListWarningConfigResponseBodyDataWarningConfigInfo {
	s.ConfigId = &v
	return s
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfo) SetConfigName(v string) *ListWarningConfigResponseBodyDataWarningConfigInfo {
	s.ConfigName = &v
	return s
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfo) SetCreateTime(v string) *ListWarningConfigResponseBodyDataWarningConfigInfo {
	s.CreateTime = &v
	return s
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfo) SetRidList(v *ListWarningConfigResponseBodyDataWarningConfigInfoRidList) *ListWarningConfigResponseBodyDataWarningConfigInfo {
	s.RidList = v
	return s
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfo) SetRuleList(v *ListWarningConfigResponseBodyDataWarningConfigInfoRuleList) *ListWarningConfigResponseBodyDataWarningConfigInfo {
	s.RuleList = v
	return s
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfo) SetStatus(v int32) *ListWarningConfigResponseBodyDataWarningConfigInfo {
	s.Status = &v
	return s
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfo) SetUpdateTime(v string) *ListWarningConfigResponseBodyDataWarningConfigInfo {
	s.UpdateTime = &v
	return s
}

type ListWarningConfigResponseBodyDataWarningConfigInfoChannels struct {
	Channel []*ListWarningConfigResponseBodyDataWarningConfigInfoChannelsChannel `json:"Channel,omitempty" xml:"Channel,omitempty" type:"Repeated"`
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfoChannels) String() string {
	return tea.Prettify(s)
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfoChannels) GoString() string {
	return s.String()
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfoChannels) SetChannel(v []*ListWarningConfigResponseBodyDataWarningConfigInfoChannelsChannel) *ListWarningConfigResponseBodyDataWarningConfigInfoChannels {
	s.Channel = v
	return s
}

type ListWarningConfigResponseBodyDataWarningConfigInfoChannelsChannel struct {
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// oapi.dingtalk.com/robot/send?access_token=c55628f700eb9ad2a3ca
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfoChannelsChannel) String() string {
	return tea.Prettify(s)
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfoChannelsChannel) GoString() string {
	return s.String()
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfoChannelsChannel) SetType(v int32) *ListWarningConfigResponseBodyDataWarningConfigInfoChannelsChannel {
	s.Type = &v
	return s
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfoChannelsChannel) SetUrl(v string) *ListWarningConfigResponseBodyDataWarningConfigInfoChannelsChannel {
	s.Url = &v
	return s
}

type ListWarningConfigResponseBodyDataWarningConfigInfoRidList struct {
	RidList []*string `json:"RidList,omitempty" xml:"RidList,omitempty" type:"Repeated"`
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfoRidList) String() string {
	return tea.Prettify(s)
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfoRidList) GoString() string {
	return s.String()
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfoRidList) SetRidList(v []*string) *ListWarningConfigResponseBodyDataWarningConfigInfoRidList {
	s.RidList = v
	return s
}

type ListWarningConfigResponseBodyDataWarningConfigInfoRuleList struct {
	WarningRule []*ListWarningConfigResponseBodyDataWarningConfigInfoRuleListWarningRule `json:"WarningRule,omitempty" xml:"WarningRule,omitempty" type:"Repeated"`
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfoRuleList) String() string {
	return tea.Prettify(s)
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfoRuleList) GoString() string {
	return s.String()
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfoRuleList) SetWarningRule(v []*ListWarningConfigResponseBodyDataWarningConfigInfoRuleListWarningRule) *ListWarningConfigResponseBodyDataWarningConfigInfoRuleList {
	s.WarningRule = v
	return s
}

type ListWarningConfigResponseBodyDataWarningConfigInfoRuleListWarningRule struct {
	// example:
	//
	// 33452
	Rid      *int64  `json:"Rid,omitempty" xml:"Rid,omitempty"`
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfoRuleListWarningRule) String() string {
	return tea.Prettify(s)
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfoRuleListWarningRule) GoString() string {
	return s.String()
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfoRuleListWarningRule) SetRid(v int64) *ListWarningConfigResponseBodyDataWarningConfigInfoRuleListWarningRule {
	s.Rid = &v
	return s
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfoRuleListWarningRule) SetRuleName(v string) *ListWarningConfigResponseBodyDataWarningConfigInfoRuleListWarningRule {
	s.RuleName = &v
	return s
}

type ListWarningConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWarningConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWarningConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWarningConfigResponse) GoString() string {
	return s.String()
}

func (s *ListWarningConfigResponse) SetHeaders(v map[string]*string) *ListWarningConfigResponse {
	s.Headers = v
	return s
}

func (s *ListWarningConfigResponse) SetStatusCode(v int32) *ListWarningConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWarningConfigResponse) SetBody(v *ListWarningConfigResponseBody) *ListWarningConfigResponse {
	s.Body = v
	return s
}

type ListWarningStrategyConfigRequest struct {
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s ListWarningStrategyConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWarningStrategyConfigRequest) GoString() string {
	return s.String()
}

func (s *ListWarningStrategyConfigRequest) SetBaseMeAgentId(v int64) *ListWarningStrategyConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *ListWarningStrategyConfigRequest) SetJsonStr(v string) *ListWarningStrategyConfigRequest {
	s.JsonStr = &v
	return s
}

type ListWarningStrategyConfigResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Count     *int32                                     `json:"Count,omitempty" xml:"Count,omitempty"`
	Data      *ListWarningStrategyConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListWarningStrategyConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWarningStrategyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListWarningStrategyConfigResponseBody) SetCode(v string) *ListWarningStrategyConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ListWarningStrategyConfigResponseBody) SetCount(v int32) *ListWarningStrategyConfigResponseBody {
	s.Count = &v
	return s
}

func (s *ListWarningStrategyConfigResponseBody) SetData(v *ListWarningStrategyConfigResponseBodyData) *ListWarningStrategyConfigResponseBody {
	s.Data = v
	return s
}

func (s *ListWarningStrategyConfigResponseBody) SetMessage(v string) *ListWarningStrategyConfigResponseBody {
	s.Message = &v
	return s
}

func (s *ListWarningStrategyConfigResponseBody) SetRequestId(v string) *ListWarningStrategyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWarningStrategyConfigResponseBody) SetSuccess(v bool) *ListWarningStrategyConfigResponseBody {
	s.Success = &v
	return s
}

type ListWarningStrategyConfigResponseBodyData struct {
	Data []*ListWarningStrategyConfigResponseBodyDataData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s ListWarningStrategyConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListWarningStrategyConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWarningStrategyConfigResponseBodyData) SetData(v []*ListWarningStrategyConfigResponseBodyDataData) *ListWarningStrategyConfigResponseBodyData {
	s.Data = v
	return s
}

type ListWarningStrategyConfigResponseBodyDataData struct {
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	IntervalTime *int64  `json:"IntervalTime,omitempty" xml:"IntervalTime,omitempty"`
	Lambda       *string `json:"Lambda,omitempty" xml:"Lambda,omitempty"`
	Level        *int64  `json:"Level,omitempty" xml:"Level,omitempty"`
	MaxNumber    *int64  `json:"MaxNumber,omitempty" xml:"MaxNumber,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListWarningStrategyConfigResponseBodyDataData) String() string {
	return tea.Prettify(s)
}

func (s ListWarningStrategyConfigResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *ListWarningStrategyConfigResponseBodyDataData) SetId(v int64) *ListWarningStrategyConfigResponseBodyDataData {
	s.Id = &v
	return s
}

func (s *ListWarningStrategyConfigResponseBodyDataData) SetIntervalTime(v int64) *ListWarningStrategyConfigResponseBodyDataData {
	s.IntervalTime = &v
	return s
}

func (s *ListWarningStrategyConfigResponseBodyDataData) SetLambda(v string) *ListWarningStrategyConfigResponseBodyDataData {
	s.Lambda = &v
	return s
}

func (s *ListWarningStrategyConfigResponseBodyDataData) SetLevel(v int64) *ListWarningStrategyConfigResponseBodyDataData {
	s.Level = &v
	return s
}

func (s *ListWarningStrategyConfigResponseBodyDataData) SetMaxNumber(v int64) *ListWarningStrategyConfigResponseBodyDataData {
	s.MaxNumber = &v
	return s
}

func (s *ListWarningStrategyConfigResponseBodyDataData) SetName(v string) *ListWarningStrategyConfigResponseBodyDataData {
	s.Name = &v
	return s
}

type ListWarningStrategyConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWarningStrategyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWarningStrategyConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWarningStrategyConfigResponse) GoString() string {
	return s.String()
}

func (s *ListWarningStrategyConfigResponse) SetHeaders(v map[string]*string) *ListWarningStrategyConfigResponse {
	s.Headers = v
	return s
}

func (s *ListWarningStrategyConfigResponse) SetStatusCode(v int32) *ListWarningStrategyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWarningStrategyConfigResponse) SetBody(v *ListWarningStrategyConfigResponseBody) *ListWarningStrategyConfigResponse {
	s.Body = v
	return s
}

type RevertAssignedSessionRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// {"isSchemeData":1,"searchParam":{"schemeTaskConfigId":1,"sourceDataType":1,"startTime":"2022-09-20 00:00:00","endTime":"2022-09-26 23:59:59"}}
	JsonStr *string `json:"jsonStr,omitempty" xml:"jsonStr,omitempty"`
}

func (s RevertAssignedSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s RevertAssignedSessionRequest) GoString() string {
	return s.String()
}

func (s *RevertAssignedSessionRequest) SetBaseMeAgentId(v int64) *RevertAssignedSessionRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *RevertAssignedSessionRequest) SetJsonStr(v string) *RevertAssignedSessionRequest {
	s.JsonStr = &v
	return s
}

type RevertAssignedSessionResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *RevertAssignedSessionResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// 82C91484-B2D5-4D2A-A21F-A6D73F4D55C6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RevertAssignedSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevertAssignedSessionResponseBody) GoString() string {
	return s.String()
}

func (s *RevertAssignedSessionResponseBody) SetCode(v string) *RevertAssignedSessionResponseBody {
	s.Code = &v
	return s
}

func (s *RevertAssignedSessionResponseBody) SetHttpStatusCode(v int32) *RevertAssignedSessionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RevertAssignedSessionResponseBody) SetMessage(v string) *RevertAssignedSessionResponseBody {
	s.Message = &v
	return s
}

func (s *RevertAssignedSessionResponseBody) SetMessages(v *RevertAssignedSessionResponseBodyMessages) *RevertAssignedSessionResponseBody {
	s.Messages = v
	return s
}

func (s *RevertAssignedSessionResponseBody) SetRequestId(v string) *RevertAssignedSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevertAssignedSessionResponseBody) SetSuccess(v bool) *RevertAssignedSessionResponseBody {
	s.Success = &v
	return s
}

type RevertAssignedSessionResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s RevertAssignedSessionResponseBodyMessages) String() string {
	return tea.Prettify(s)
}

func (s RevertAssignedSessionResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *RevertAssignedSessionResponseBodyMessages) SetMessage(v []*string) *RevertAssignedSessionResponseBodyMessages {
	s.Message = v
	return s
}

type RevertAssignedSessionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevertAssignedSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevertAssignedSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s RevertAssignedSessionResponse) GoString() string {
	return s.String()
}

func (s *RevertAssignedSessionResponse) SetHeaders(v map[string]*string) *RevertAssignedSessionResponse {
	s.Headers = v
	return s
}

func (s *RevertAssignedSessionResponse) SetStatusCode(v int32) *RevertAssignedSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *RevertAssignedSessionResponse) SetBody(v *RevertAssignedSessionResponseBody) *RevertAssignedSessionResponse {
	s.Body = v
	return s
}

type RevertAssignedSessionGroupRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// {"isSchemeData":1,"forceRevertSessionGroup":true,"sessionGroupIdList":["1"]}
	JsonStr *string `json:"jsonStr,omitempty" xml:"jsonStr,omitempty"`
}

func (s RevertAssignedSessionGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RevertAssignedSessionGroupRequest) GoString() string {
	return s.String()
}

func (s *RevertAssignedSessionGroupRequest) SetBaseMeAgentId(v int64) *RevertAssignedSessionGroupRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *RevertAssignedSessionGroupRequest) SetJsonStr(v string) *RevertAssignedSessionGroupRequest {
	s.JsonStr = &v
	return s
}

type RevertAssignedSessionGroupResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *RevertAssignedSessionGroupResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24F4CE647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RevertAssignedSessionGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevertAssignedSessionGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RevertAssignedSessionGroupResponseBody) SetCode(v string) *RevertAssignedSessionGroupResponseBody {
	s.Code = &v
	return s
}

func (s *RevertAssignedSessionGroupResponseBody) SetHttpStatusCode(v int32) *RevertAssignedSessionGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RevertAssignedSessionGroupResponseBody) SetMessage(v string) *RevertAssignedSessionGroupResponseBody {
	s.Message = &v
	return s
}

func (s *RevertAssignedSessionGroupResponseBody) SetMessages(v *RevertAssignedSessionGroupResponseBodyMessages) *RevertAssignedSessionGroupResponseBody {
	s.Messages = v
	return s
}

func (s *RevertAssignedSessionGroupResponseBody) SetRequestId(v string) *RevertAssignedSessionGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevertAssignedSessionGroupResponseBody) SetSuccess(v bool) *RevertAssignedSessionGroupResponseBody {
	s.Success = &v
	return s
}

type RevertAssignedSessionGroupResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s RevertAssignedSessionGroupResponseBodyMessages) String() string {
	return tea.Prettify(s)
}

func (s RevertAssignedSessionGroupResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *RevertAssignedSessionGroupResponseBodyMessages) SetMessage(v []*string) *RevertAssignedSessionGroupResponseBodyMessages {
	s.Message = v
	return s
}

type RevertAssignedSessionGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevertAssignedSessionGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevertAssignedSessionGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RevertAssignedSessionGroupResponse) GoString() string {
	return s.String()
}

func (s *RevertAssignedSessionGroupResponse) SetHeaders(v map[string]*string) *RevertAssignedSessionGroupResponse {
	s.Headers = v
	return s
}

func (s *RevertAssignedSessionGroupResponse) SetStatusCode(v int32) *RevertAssignedSessionGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RevertAssignedSessionGroupResponse) SetBody(v *RevertAssignedSessionGroupResponseBody) *RevertAssignedSessionGroupResponse {
	s.Body = v
	return s
}

type SaveConfigDataSetRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s SaveConfigDataSetRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveConfigDataSetRequest) GoString() string {
	return s.String()
}

func (s *SaveConfigDataSetRequest) SetBaseMeAgentId(v int64) *SaveConfigDataSetRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *SaveConfigDataSetRequest) SetJsonStr(v string) *SaveConfigDataSetRequest {
	s.JsonStr = &v
	return s
}

type SaveConfigDataSetResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveConfigDataSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveConfigDataSetResponseBody) GoString() string {
	return s.String()
}

func (s *SaveConfigDataSetResponseBody) SetCode(v string) *SaveConfigDataSetResponseBody {
	s.Code = &v
	return s
}

func (s *SaveConfigDataSetResponseBody) SetMessage(v string) *SaveConfigDataSetResponseBody {
	s.Message = &v
	return s
}

func (s *SaveConfigDataSetResponseBody) SetRequestId(v string) *SaveConfigDataSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveConfigDataSetResponseBody) SetSuccess(v bool) *SaveConfigDataSetResponseBody {
	s.Success = &v
	return s
}

type SaveConfigDataSetResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveConfigDataSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveConfigDataSetResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveConfigDataSetResponse) GoString() string {
	return s.String()
}

func (s *SaveConfigDataSetResponse) SetHeaders(v map[string]*string) *SaveConfigDataSetResponse {
	s.Headers = v
	return s
}

func (s *SaveConfigDataSetResponse) SetStatusCode(v int32) *SaveConfigDataSetResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveConfigDataSetResponse) SetBody(v *SaveConfigDataSetResponseBody) *SaveConfigDataSetResponse {
	s.Body = v
	return s
}

type SubmitComplaintRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s SubmitComplaintRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitComplaintRequest) GoString() string {
	return s.String()
}

func (s *SubmitComplaintRequest) SetBaseMeAgentId(v int64) *SubmitComplaintRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *SubmitComplaintRequest) SetJsonStr(v string) *SubmitComplaintRequest {
	s.JsonStr = &v
	return s
}

type SubmitComplaintResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 90
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 82C91484-B2D5-4D2A-A21F-A6D73F4D55C6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitComplaintResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitComplaintResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitComplaintResponseBody) SetCode(v string) *SubmitComplaintResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitComplaintResponseBody) SetData(v string) *SubmitComplaintResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitComplaintResponseBody) SetMessage(v string) *SubmitComplaintResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitComplaintResponseBody) SetRequestId(v string) *SubmitComplaintResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitComplaintResponseBody) SetSuccess(v bool) *SubmitComplaintResponseBody {
	s.Success = &v
	return s
}

type SubmitComplaintResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitComplaintResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitComplaintResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitComplaintResponse) GoString() string {
	return s.String()
}

func (s *SubmitComplaintResponse) SetHeaders(v map[string]*string) *SubmitComplaintResponse {
	s.Headers = v
	return s
}

func (s *SubmitComplaintResponse) SetStatusCode(v int32) *SubmitComplaintResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitComplaintResponse) SetBody(v *SubmitComplaintResponseBody) *SubmitComplaintResponse {
	s.Body = v
	return s
}

type SubmitPrecisionTaskRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// "{"name":"test","dataSetId":1865}"
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s SubmitPrecisionTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitPrecisionTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitPrecisionTaskRequest) SetBaseMeAgentId(v int64) *SubmitPrecisionTaskRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *SubmitPrecisionTaskRequest) SetJsonStr(v string) *SubmitPrecisionTaskRequest {
	s.JsonStr = &v
	return s
}

type SubmitPrecisionTaskResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// EA701F66-8CA2-4A79-8E3C-A6F2FA****
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24F4CE647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitPrecisionTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitPrecisionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitPrecisionTaskResponseBody) SetCode(v string) *SubmitPrecisionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitPrecisionTaskResponseBody) SetData(v string) *SubmitPrecisionTaskResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitPrecisionTaskResponseBody) SetMessage(v string) *SubmitPrecisionTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitPrecisionTaskResponseBody) SetRequestId(v string) *SubmitPrecisionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitPrecisionTaskResponseBody) SetSuccess(v bool) *SubmitPrecisionTaskResponseBody {
	s.Success = &v
	return s
}

type SubmitPrecisionTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitPrecisionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitPrecisionTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitPrecisionTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitPrecisionTaskResponse) SetHeaders(v map[string]*string) *SubmitPrecisionTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitPrecisionTaskResponse) SetStatusCode(v int32) *SubmitPrecisionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitPrecisionTaskResponse) SetBody(v *SubmitPrecisionTaskResponseBody) *SubmitPrecisionTaskResponse {
	s.Body = v
	return s
}

type SubmitQualityCheckTaskRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s SubmitQualityCheckTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitQualityCheckTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitQualityCheckTaskRequest) SetBaseMeAgentId(v int64) *SubmitQualityCheckTaskRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *SubmitQualityCheckTaskRequest) SetJsonStr(v string) *SubmitQualityCheckTaskRequest {
	s.JsonStr = &v
	return s
}

type SubmitQualityCheckTaskResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// F6C2B68F-2311-4495-82AC-DAE86C9****
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 00A044A2-D59B-4104-96BA-84060AE8345F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitQualityCheckTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitQualityCheckTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitQualityCheckTaskResponseBody) SetCode(v string) *SubmitQualityCheckTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitQualityCheckTaskResponseBody) SetData(v string) *SubmitQualityCheckTaskResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitQualityCheckTaskResponseBody) SetMessage(v string) *SubmitQualityCheckTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitQualityCheckTaskResponseBody) SetRequestId(v string) *SubmitQualityCheckTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitQualityCheckTaskResponseBody) SetSuccess(v bool) *SubmitQualityCheckTaskResponseBody {
	s.Success = &v
	return s
}

type SubmitQualityCheckTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitQualityCheckTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitQualityCheckTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitQualityCheckTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitQualityCheckTaskResponse) SetHeaders(v map[string]*string) *SubmitQualityCheckTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitQualityCheckTaskResponse) SetStatusCode(v int32) *SubmitQualityCheckTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitQualityCheckTaskResponse) SetBody(v *SubmitQualityCheckTaskResponseBody) *SubmitQualityCheckTaskResponse {
	s.Body = v
	return s
}

type SubmitReviewInfoRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s SubmitReviewInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitReviewInfoRequest) GoString() string {
	return s.String()
}

func (s *SubmitReviewInfoRequest) SetBaseMeAgentId(v int64) *SubmitReviewInfoRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *SubmitReviewInfoRequest) SetJsonStr(v string) *SubmitReviewInfoRequest {
	s.JsonStr = &v
	return s
}

type SubmitReviewInfoResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 95
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9987D326-83D9-4A42-B9A5-0B27F9B40539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitReviewInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitReviewInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitReviewInfoResponseBody) SetCode(v string) *SubmitReviewInfoResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitReviewInfoResponseBody) SetData(v string) *SubmitReviewInfoResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitReviewInfoResponseBody) SetMessage(v string) *SubmitReviewInfoResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitReviewInfoResponseBody) SetRequestId(v string) *SubmitReviewInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitReviewInfoResponseBody) SetSuccess(v bool) *SubmitReviewInfoResponseBody {
	s.Success = &v
	return s
}

type SubmitReviewInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitReviewInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitReviewInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitReviewInfoResponse) GoString() string {
	return s.String()
}

func (s *SubmitReviewInfoResponse) SetHeaders(v map[string]*string) *SubmitReviewInfoResponse {
	s.Headers = v
	return s
}

func (s *SubmitReviewInfoResponse) SetStatusCode(v int32) *SubmitReviewInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitReviewInfoResponse) SetBody(v *SubmitReviewInfoResponseBody) *SubmitReviewInfoResponse {
	s.Body = v
	return s
}

type SyncQualityCheckRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"tid":"20200823-234234","dialogue":"{}"}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s SyncQualityCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncQualityCheckRequest) GoString() string {
	return s.String()
}

func (s *SyncQualityCheckRequest) SetBaseMeAgentId(v int64) *SyncQualityCheckRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *SyncQualityCheckRequest) SetJsonStr(v string) *SyncQualityCheckRequest {
	s.JsonStr = &v
	return s
}

type SyncQualityCheckResponseBody struct {
	// example:
	//
	// 200
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SyncQualityCheckResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 66E1ACB8-17B2-4BE8-8581-954A8*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SyncQualityCheckResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncQualityCheckResponseBody) GoString() string {
	return s.String()
}

func (s *SyncQualityCheckResponseBody) SetCode(v string) *SyncQualityCheckResponseBody {
	s.Code = &v
	return s
}

func (s *SyncQualityCheckResponseBody) SetData(v *SyncQualityCheckResponseBodyData) *SyncQualityCheckResponseBody {
	s.Data = v
	return s
}

func (s *SyncQualityCheckResponseBody) SetMessage(v string) *SyncQualityCheckResponseBody {
	s.Message = &v
	return s
}

func (s *SyncQualityCheckResponseBody) SetRequestId(v string) *SyncQualityCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncQualityCheckResponseBody) SetSuccess(v bool) *SyncQualityCheckResponseBody {
	s.Success = &v
	return s
}

type SyncQualityCheckResponseBodyData struct {
	// example:
	//
	// 1584535485856
	BeginTime *int64                                   `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	Rules     []*SyncQualityCheckResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// 66E1ACB866E1ACB8
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 20200876-66E1ACB8
	Tid *string `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s SyncQualityCheckResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SyncQualityCheckResponseBodyData) GoString() string {
	return s.String()
}

func (s *SyncQualityCheckResponseBodyData) SetBeginTime(v int64) *SyncQualityCheckResponseBodyData {
	s.BeginTime = &v
	return s
}

func (s *SyncQualityCheckResponseBodyData) SetRules(v []*SyncQualityCheckResponseBodyDataRules) *SyncQualityCheckResponseBodyData {
	s.Rules = v
	return s
}

func (s *SyncQualityCheckResponseBodyData) SetScore(v int32) *SyncQualityCheckResponseBodyData {
	s.Score = &v
	return s
}

func (s *SyncQualityCheckResponseBodyData) SetTaskId(v string) *SyncQualityCheckResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SyncQualityCheckResponseBodyData) SetTid(v string) *SyncQualityCheckResponseBodyData {
	s.Tid = &v
	return s
}

type SyncQualityCheckResponseBodyDataRules struct {
	Hit []*SyncQualityCheckResponseBodyDataRulesHit `json:"Hit,omitempty" xml:"Hit,omitempty" type:"Repeated"`
	// example:
	//
	// 232232
	Rid      *string `json:"Rid,omitempty" xml:"Rid,omitempty"`
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s SyncQualityCheckResponseBodyDataRules) String() string {
	return tea.Prettify(s)
}

func (s SyncQualityCheckResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *SyncQualityCheckResponseBodyDataRules) SetHit(v []*SyncQualityCheckResponseBodyDataRulesHit) *SyncQualityCheckResponseBodyDataRules {
	s.Hit = v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRules) SetRid(v string) *SyncQualityCheckResponseBodyDataRules {
	s.Rid = &v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRules) SetRuleName(v string) *SyncQualityCheckResponseBodyDataRules {
	s.RuleName = &v
	return s
}

type SyncQualityCheckResponseBodyDataRulesHit struct {
	HitKeyWords []*SyncQualityCheckResponseBodyDataRulesHitHitKeyWords `json:"HitKeyWords,omitempty" xml:"HitKeyWords,omitempty" type:"Repeated"`
	Phrase      *SyncQualityCheckResponseBodyDataRulesHitPhrase        `json:"Phrase,omitempty" xml:"Phrase,omitempty" type:"Struct"`
}

func (s SyncQualityCheckResponseBodyDataRulesHit) String() string {
	return tea.Prettify(s)
}

func (s SyncQualityCheckResponseBodyDataRulesHit) GoString() string {
	return s.String()
}

func (s *SyncQualityCheckResponseBodyDataRulesHit) SetHitKeyWords(v []*SyncQualityCheckResponseBodyDataRulesHitHitKeyWords) *SyncQualityCheckResponseBodyDataRulesHit {
	s.HitKeyWords = v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRulesHit) SetPhrase(v *SyncQualityCheckResponseBodyDataRulesHitPhrase) *SyncQualityCheckResponseBodyDataRulesHit {
	s.Phrase = v
	return s
}

type SyncQualityCheckResponseBodyDataRulesHitHitKeyWords struct {
	// example:
	//
	// 2312
	Cid *int32 `json:"Cid,omitempty" xml:"Cid,omitempty"`
	// example:
	//
	// 1
	From *int32 `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// 4
	Pid *int32 `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// example:
	//
	// 4
	To  *int32  `json:"To,omitempty" xml:"To,omitempty"`
	Val *string `json:"Val,omitempty" xml:"Val,omitempty"`
}

func (s SyncQualityCheckResponseBodyDataRulesHitHitKeyWords) String() string {
	return tea.Prettify(s)
}

func (s SyncQualityCheckResponseBodyDataRulesHitHitKeyWords) GoString() string {
	return s.String()
}

func (s *SyncQualityCheckResponseBodyDataRulesHitHitKeyWords) SetCid(v int32) *SyncQualityCheckResponseBodyDataRulesHitHitKeyWords {
	s.Cid = &v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRulesHitHitKeyWords) SetFrom(v int32) *SyncQualityCheckResponseBodyDataRulesHitHitKeyWords {
	s.From = &v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRulesHitHitKeyWords) SetPid(v int32) *SyncQualityCheckResponseBodyDataRulesHitHitKeyWords {
	s.Pid = &v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRulesHitHitKeyWords) SetTo(v int32) *SyncQualityCheckResponseBodyDataRulesHitHitKeyWords {
	s.To = &v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRulesHitHitKeyWords) SetVal(v string) *SyncQualityCheckResponseBodyDataRulesHitHitKeyWords {
	s.Val = &v
	return s
}

type SyncQualityCheckResponseBodyDataRulesHitPhrase struct {
	// example:
	//
	// 1230
	Begin *int64 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// example:
	//
	// 6
	EmotionValue *int32 `json:"EmotionValue,omitempty" xml:"EmotionValue,omitempty"`
	// example:
	//
	// 3440
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// example:
	//
	// xxx
	Identity *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
	Role     *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// 123
	SilenceDuration *int32 `json:"SilenceDuration,omitempty" xml:"SilenceDuration,omitempty"`
	// example:
	//
	// 233
	SpeechRate *int32  `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	Words      *string `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s SyncQualityCheckResponseBodyDataRulesHitPhrase) String() string {
	return tea.Prettify(s)
}

func (s SyncQualityCheckResponseBodyDataRulesHitPhrase) GoString() string {
	return s.String()
}

func (s *SyncQualityCheckResponseBodyDataRulesHitPhrase) SetBegin(v int64) *SyncQualityCheckResponseBodyDataRulesHitPhrase {
	s.Begin = &v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRulesHitPhrase) SetEmotionValue(v int32) *SyncQualityCheckResponseBodyDataRulesHitPhrase {
	s.EmotionValue = &v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRulesHitPhrase) SetEnd(v int64) *SyncQualityCheckResponseBodyDataRulesHitPhrase {
	s.End = &v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRulesHitPhrase) SetIdentity(v string) *SyncQualityCheckResponseBodyDataRulesHitPhrase {
	s.Identity = &v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRulesHitPhrase) SetRole(v string) *SyncQualityCheckResponseBodyDataRulesHitPhrase {
	s.Role = &v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRulesHitPhrase) SetSilenceDuration(v int32) *SyncQualityCheckResponseBodyDataRulesHitPhrase {
	s.SilenceDuration = &v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRulesHitPhrase) SetSpeechRate(v int32) *SyncQualityCheckResponseBodyDataRulesHitPhrase {
	s.SpeechRate = &v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRulesHitPhrase) SetWords(v string) *SyncQualityCheckResponseBodyDataRulesHitPhrase {
	s.Words = &v
	return s
}

type SyncQualityCheckResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncQualityCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncQualityCheckResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncQualityCheckResponse) GoString() string {
	return s.String()
}

func (s *SyncQualityCheckResponse) SetHeaders(v map[string]*string) *SyncQualityCheckResponse {
	s.Headers = v
	return s
}

func (s *SyncQualityCheckResponse) SetStatusCode(v int32) *SyncQualityCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncQualityCheckResponse) SetBody(v *SyncQualityCheckResponseBody) *SyncQualityCheckResponse {
	s.Body = v
	return s
}

type TestRuleV4Request struct {
	// example:
	//
	// 1
	IsSchemeData *int32 `json:"IsSchemeData,omitempty" xml:"IsSchemeData,omitempty"`
	// This parameter is required.
	TestJson *string `json:"TestJson,omitempty" xml:"TestJson,omitempty"`
}

func (s TestRuleV4Request) String() string {
	return tea.Prettify(s)
}

func (s TestRuleV4Request) GoString() string {
	return s.String()
}

func (s *TestRuleV4Request) SetIsSchemeData(v int32) *TestRuleV4Request {
	s.IsSchemeData = &v
	return s
}

func (s *TestRuleV4Request) SetTestJson(v string) *TestRuleV4Request {
	s.TestJson = &v
	return s
}

type TestRuleV4ResponseBody struct {
	// example:
	//
	// 200
	Code *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *TestRuleV4ResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 96138D8D-XXXX-4E41-XXXX-77AED1088BBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TestRuleV4ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TestRuleV4ResponseBody) GoString() string {
	return s.String()
}

func (s *TestRuleV4ResponseBody) SetCode(v string) *TestRuleV4ResponseBody {
	s.Code = &v
	return s
}

func (s *TestRuleV4ResponseBody) SetData(v *TestRuleV4ResponseBodyData) *TestRuleV4ResponseBody {
	s.Data = v
	return s
}

func (s *TestRuleV4ResponseBody) SetHttpStatusCode(v int32) *TestRuleV4ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *TestRuleV4ResponseBody) SetMessage(v string) *TestRuleV4ResponseBody {
	s.Message = &v
	return s
}

func (s *TestRuleV4ResponseBody) SetRequestId(v string) *TestRuleV4ResponseBody {
	s.RequestId = &v
	return s
}

func (s *TestRuleV4ResponseBody) SetSuccess(v bool) *TestRuleV4ResponseBody {
	s.Success = &v
	return s
}

type TestRuleV4ResponseBodyData struct {
	HitRuleReviewInfoList   []*TestRuleV4ResponseBodyDataHitRuleReviewInfoList   `json:"HitRuleReviewInfoList,omitempty" xml:"HitRuleReviewInfoList,omitempty" type:"Repeated"`
	HitTaskFlowList         []*TestRuleV4ResponseBodyDataHitTaskFlowList         `json:"HitTaskFlowList,omitempty" xml:"HitTaskFlowList,omitempty" type:"Repeated"`
	UnhitRuleReviewInfoList []*TestRuleV4ResponseBodyDataUnhitRuleReviewInfoList `json:"UnhitRuleReviewInfoList,omitempty" xml:"UnhitRuleReviewInfoList,omitempty" type:"Repeated"`
}

func (s TestRuleV4ResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TestRuleV4ResponseBodyData) GoString() string {
	return s.String()
}

func (s *TestRuleV4ResponseBodyData) SetHitRuleReviewInfoList(v []*TestRuleV4ResponseBodyDataHitRuleReviewInfoList) *TestRuleV4ResponseBodyData {
	s.HitRuleReviewInfoList = v
	return s
}

func (s *TestRuleV4ResponseBodyData) SetHitTaskFlowList(v []*TestRuleV4ResponseBodyDataHitTaskFlowList) *TestRuleV4ResponseBodyData {
	s.HitTaskFlowList = v
	return s
}

func (s *TestRuleV4ResponseBodyData) SetUnhitRuleReviewInfoList(v []*TestRuleV4ResponseBodyDataUnhitRuleReviewInfoList) *TestRuleV4ResponseBodyData {
	s.UnhitRuleReviewInfoList = v
	return s
}

type TestRuleV4ResponseBodyDataHitRuleReviewInfoList struct {
	// example:
	//
	// 1
	BranchHitId          *int64                                                                 `json:"BranchHitId,omitempty" xml:"BranchHitId,omitempty"`
	BranchInfoList       []*TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList       `json:"BranchInfoList,omitempty" xml:"BranchInfoList,omitempty" type:"Repeated"`
	ConditionHitInfoList []*TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoList `json:"ConditionHitInfoList,omitempty" xml:"ConditionHitInfoList,omitempty" type:"Repeated"`
	ConditionInfoList    []*ConditionBasicInfo                                                  `json:"ConditionInfoList,omitempty" xml:"ConditionInfoList,omitempty" type:"Repeated"`
	JudgeNodeName        *string                                                                `json:"JudgeNodeName,omitempty" xml:"JudgeNodeName,omitempty"`
	// example:
	//
	// a&&b
	Lambda *string `json:"Lambda,omitempty" xml:"Lambda,omitempty"`
	// example:
	//
	// true
	Matched *bool `json:"Matched,omitempty" xml:"Matched,omitempty"`
	// example:
	//
	// 0
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// example:
	//
	// 451
	Rid      *int64  `json:"Rid,omitempty" xml:"Rid,omitempty"`
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// 1
	RuleScoreType *int32 `json:"RuleScoreType,omitempty" xml:"RuleScoreType,omitempty"`
	// example:
	//
	// 1
	ScoreNumType *int32 `json:"ScoreNumType,omitempty" xml:"ScoreNumType,omitempty"`
	// example:
	//
	// 1
	TaskFlowId *int64 `json:"TaskFlowId,omitempty" xml:"TaskFlowId,omitempty"`
}

func (s TestRuleV4ResponseBodyDataHitRuleReviewInfoList) String() string {
	return tea.Prettify(s)
}

func (s TestRuleV4ResponseBodyDataHitRuleReviewInfoList) GoString() string {
	return s.String()
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) SetBranchHitId(v int64) *TestRuleV4ResponseBodyDataHitRuleReviewInfoList {
	s.BranchHitId = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) SetBranchInfoList(v []*TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList) *TestRuleV4ResponseBodyDataHitRuleReviewInfoList {
	s.BranchInfoList = v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) SetConditionHitInfoList(v []*TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoList) *TestRuleV4ResponseBodyDataHitRuleReviewInfoList {
	s.ConditionHitInfoList = v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) SetConditionInfoList(v []*ConditionBasicInfo) *TestRuleV4ResponseBodyDataHitRuleReviewInfoList {
	s.ConditionInfoList = v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) SetJudgeNodeName(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoList {
	s.JudgeNodeName = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) SetLambda(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoList {
	s.Lambda = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) SetMatched(v bool) *TestRuleV4ResponseBodyDataHitRuleReviewInfoList {
	s.Matched = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) SetNodeType(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoList {
	s.NodeType = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) SetRid(v int64) *TestRuleV4ResponseBodyDataHitRuleReviewInfoList {
	s.Rid = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) SetRuleName(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoList {
	s.RuleName = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) SetRuleScoreType(v int32) *TestRuleV4ResponseBodyDataHitRuleReviewInfoList {
	s.RuleScoreType = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) SetScoreNumType(v int32) *TestRuleV4ResponseBodyDataHitRuleReviewInfoList {
	s.ScoreNumType = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) SetTaskFlowId(v int64) *TestRuleV4ResponseBodyDataHitRuleReviewInfoList {
	s.TaskFlowId = &v
	return s
}

type TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList struct {
	// example:
	//
	// 0
	CheckType *int32 `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// a&&b
	Lambda *string `json:"Lambda,omitempty" xml:"Lambda,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2
	NextNodeId *int64              `json:"NextNodeId,omitempty" xml:"NextNodeId,omitempty"`
	Situation  *NextNodeSituations `json:"Situation,omitempty" xml:"Situation,omitempty"`
	Triggers   []*string           `json:"Triggers,omitempty" xml:"Triggers,omitempty" type:"Repeated"`
}

func (s TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList) String() string {
	return tea.Prettify(s)
}

func (s TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList) GoString() string {
	return s.String()
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList) SetCheckType(v int32) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList {
	s.CheckType = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList) SetIndex(v int32) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList {
	s.Index = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList) SetLambda(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList {
	s.Lambda = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList) SetName(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList {
	s.Name = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList) SetNextNodeId(v int64) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList {
	s.NextNodeId = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList) SetSituation(v *NextNodeSituations) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList {
	s.Situation = v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList) SetTriggers(v []*string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList {
	s.Triggers = v
	return s
}

type TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoList struct {
	Cid      []*string                                                                      `json:"Cid,omitempty" xml:"Cid,omitempty" type:"Repeated"`
	KeyWords []*TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords `json:"KeyWords,omitempty" xml:"KeyWords,omitempty" type:"Repeated"`
	Phrase   *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase     `json:"Phrase,omitempty" xml:"Phrase,omitempty" type:"Struct"`
}

func (s TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoList) String() string {
	return tea.Prettify(s)
}

func (s TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoList) GoString() string {
	return s.String()
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoList) SetCid(v []*string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoList {
	s.Cid = v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoList) SetKeyWords(v []*TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoList {
	s.KeyWords = v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoList) SetPhrase(v *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoList {
	s.Phrase = v
	return s
}

type TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords struct {
	// example:
	//
	// 4
	Cid           *string `json:"Cid,omitempty" xml:"Cid,omitempty"`
	CustomizeCode *string `json:"CustomizeCode,omitempty" xml:"CustomizeCode,omitempty"`
	// example:
	//
	// 1
	From *int32 `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// 123
	Oid         *string `json:"Oid,omitempty" xml:"Oid,omitempty"`
	OperatorKey *string `json:"OperatorKey,omitempty" xml:"OperatorKey,omitempty"`
	// example:
	//
	// 13
	Pid           *int32  `json:"Pid,omitempty" xml:"Pid,omitempty"`
	SimilarPhrase *string `json:"SimilarPhrase,omitempty" xml:"SimilarPhrase,omitempty"`
	Tid           *string `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// example:
	//
	// 3
	To   *int32  `json:"To,omitempty" xml:"To,omitempty"`
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	Val  *string `json:"Val,omitempty" xml:"Val,omitempty"`
}

func (s TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) String() string {
	return tea.Prettify(s)
}

func (s TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) GoString() string {
	return s.String()
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) SetCid(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords {
	s.Cid = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) SetCustomizeCode(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords {
	s.CustomizeCode = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) SetFrom(v int32) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords {
	s.From = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) SetOid(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords {
	s.Oid = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) SetOperatorKey(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords {
	s.OperatorKey = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) SetPid(v int32) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords {
	s.Pid = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) SetSimilarPhrase(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords {
	s.SimilarPhrase = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) SetTid(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords {
	s.Tid = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) SetTo(v int32) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords {
	s.To = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) SetUuid(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords {
	s.Uuid = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) SetVal(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords {
	s.Val = &v
	return s
}

type TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase struct {
	// example:
	//
	// 72000
	Begin *int64 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// example:
	//
	// 2019-11-25 15:37:16
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// 0
	ChannelId               *int32 `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	EmotionFineGrainedValue *int32 `json:"EmotionFineGrainedValue,omitempty" xml:"EmotionFineGrainedValue,omitempty"`
	// example:
	//
	// 7
	EmotionValue *int32 `json:"EmotionValue,omitempty" xml:"EmotionValue,omitempty"`
	// example:
	//
	// 80000
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// example:
	//
	// 1
	HitStatus *int32 `json:"HitStatus,omitempty" xml:"HitStatus,omitempty"`
	// example:
	//
	// 10:00:00
	HourMinSec *string `json:"HourMinSec,omitempty" xml:"HourMinSec,omitempty"`
	Identity   *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
	// example:
	//
	// 3
	Pid      *int32  `json:"Pid,omitempty" xml:"Pid,omitempty"`
	RenterId *int64  `json:"RenterId,omitempty" xml:"RenterId,omitempty"`
	Role     *string `json:"Role,omitempty" xml:"Role,omitempty"`
	Sid      *int64  `json:"Sid,omitempty" xml:"Sid,omitempty"`
	// example:
	//
	// 1000
	SilenceDuration *int32 `json:"SilenceDuration,omitempty" xml:"SilenceDuration,omitempty"`
	// example:
	//
	// 100
	SpeechRate *int32  `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	Uuid       *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	Words      *string `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) String() string {
	return tea.Prettify(s)
}

func (s TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) GoString() string {
	return s.String()
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) SetBegin(v int64) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase {
	s.Begin = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) SetBeginTime(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase {
	s.BeginTime = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) SetChannelId(v int32) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase {
	s.ChannelId = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) SetEmotionFineGrainedValue(v int32) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase {
	s.EmotionFineGrainedValue = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) SetEmotionValue(v int32) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase {
	s.EmotionValue = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) SetEnd(v int64) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase {
	s.End = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) SetHitStatus(v int32) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase {
	s.HitStatus = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) SetHourMinSec(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase {
	s.HourMinSec = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) SetIdentity(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase {
	s.Identity = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) SetPid(v int32) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase {
	s.Pid = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) SetRenterId(v int64) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase {
	s.RenterId = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) SetRole(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase {
	s.Role = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) SetSid(v int64) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase {
	s.Sid = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) SetSilenceDuration(v int32) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase {
	s.SilenceDuration = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) SetSpeechRate(v int32) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase {
	s.SpeechRate = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) SetUuid(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase {
	s.Uuid = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) SetWords(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase {
	s.Words = &v
	return s
}

type TestRuleV4ResponseBodyDataHitTaskFlowList struct {
	GraphFlow *TaskGraphFlow `json:"GraphFlow,omitempty" xml:"GraphFlow,omitempty"`
	// example:
	//
	// 1
	Rid          *int64 `json:"Rid,omitempty" xml:"Rid,omitempty"`
	TaskFlowType *int32 `json:"TaskFlowType,omitempty" xml:"TaskFlowType,omitempty"`
}

func (s TestRuleV4ResponseBodyDataHitTaskFlowList) String() string {
	return tea.Prettify(s)
}

func (s TestRuleV4ResponseBodyDataHitTaskFlowList) GoString() string {
	return s.String()
}

func (s *TestRuleV4ResponseBodyDataHitTaskFlowList) SetGraphFlow(v *TaskGraphFlow) *TestRuleV4ResponseBodyDataHitTaskFlowList {
	s.GraphFlow = v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitTaskFlowList) SetRid(v int64) *TestRuleV4ResponseBodyDataHitTaskFlowList {
	s.Rid = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitTaskFlowList) SetTaskFlowType(v int32) *TestRuleV4ResponseBodyDataHitTaskFlowList {
	s.TaskFlowType = &v
	return s
}

type TestRuleV4ResponseBodyDataUnhitRuleReviewInfoList struct {
	ConditionInfoList []*ConditionBasicInfo `json:"ConditionInfoList,omitempty" xml:"ConditionInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Matched *bool `json:"Matched,omitempty" xml:"Matched,omitempty"`
	// example:
	//
	// 2
	Rid          *int64 `json:"Rid,omitempty" xml:"Rid,omitempty"`
	TaskFlowType *int32 `json:"TaskFlowType,omitempty" xml:"TaskFlowType,omitempty"`
}

func (s TestRuleV4ResponseBodyDataUnhitRuleReviewInfoList) String() string {
	return tea.Prettify(s)
}

func (s TestRuleV4ResponseBodyDataUnhitRuleReviewInfoList) GoString() string {
	return s.String()
}

func (s *TestRuleV4ResponseBodyDataUnhitRuleReviewInfoList) SetConditionInfoList(v []*ConditionBasicInfo) *TestRuleV4ResponseBodyDataUnhitRuleReviewInfoList {
	s.ConditionInfoList = v
	return s
}

func (s *TestRuleV4ResponseBodyDataUnhitRuleReviewInfoList) SetMatched(v bool) *TestRuleV4ResponseBodyDataUnhitRuleReviewInfoList {
	s.Matched = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataUnhitRuleReviewInfoList) SetRid(v int64) *TestRuleV4ResponseBodyDataUnhitRuleReviewInfoList {
	s.Rid = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataUnhitRuleReviewInfoList) SetTaskFlowType(v int32) *TestRuleV4ResponseBodyDataUnhitRuleReviewInfoList {
	s.TaskFlowType = &v
	return s
}

type TestRuleV4Response struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TestRuleV4ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TestRuleV4Response) String() string {
	return tea.Prettify(s)
}

func (s TestRuleV4Response) GoString() string {
	return s.String()
}

func (s *TestRuleV4Response) SetHeaders(v map[string]*string) *TestRuleV4Response {
	s.Headers = v
	return s
}

func (s *TestRuleV4Response) SetStatusCode(v int32) *TestRuleV4Response {
	s.StatusCode = &v
	return s
}

func (s *TestRuleV4Response) SetBody(v *TestRuleV4ResponseBody) *TestRuleV4Response {
	s.Body = v
	return s
}

type UpdateAsrVocabRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UpdateAsrVocabRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAsrVocabRequest) GoString() string {
	return s.String()
}

func (s *UpdateAsrVocabRequest) SetBaseMeAgentId(v int64) *UpdateAsrVocabRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateAsrVocabRequest) SetJsonStr(v string) *UpdateAsrVocabRequest {
	s.JsonStr = &v
	return s
}

type UpdateAsrVocabResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 71b1795ac8634bd8bdf4d3878480c7c2
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9987D326-83D9-4A42-B9A5-0B27F9B40539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAsrVocabResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAsrVocabResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAsrVocabResponseBody) SetCode(v string) *UpdateAsrVocabResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAsrVocabResponseBody) SetData(v string) *UpdateAsrVocabResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateAsrVocabResponseBody) SetMessage(v string) *UpdateAsrVocabResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAsrVocabResponseBody) SetRequestId(v string) *UpdateAsrVocabResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAsrVocabResponseBody) SetSuccess(v bool) *UpdateAsrVocabResponseBody {
	s.Success = &v
	return s
}

type UpdateAsrVocabResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAsrVocabResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAsrVocabResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAsrVocabResponse) GoString() string {
	return s.String()
}

func (s *UpdateAsrVocabResponse) SetHeaders(v map[string]*string) *UpdateAsrVocabResponse {
	s.Headers = v
	return s
}

func (s *UpdateAsrVocabResponse) SetStatusCode(v int32) *UpdateAsrVocabResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAsrVocabResponse) SetBody(v *UpdateAsrVocabResponseBody) *UpdateAsrVocabResponse {
	s.Body = v
	return s
}

type UpdateCheckTypeToSchemeRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64  `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	JsonStr       *string `json:"jsonStr,omitempty" xml:"jsonStr,omitempty"`
}

func (s UpdateCheckTypeToSchemeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCheckTypeToSchemeRequest) GoString() string {
	return s.String()
}

func (s *UpdateCheckTypeToSchemeRequest) SetBaseMeAgentId(v int64) *UpdateCheckTypeToSchemeRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateCheckTypeToSchemeRequest) SetJsonStr(v string) *UpdateCheckTypeToSchemeRequest {
	s.JsonStr = &v
	return s
}

type UpdateCheckTypeToSchemeResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 4
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *UpdateCheckTypeToSchemeResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// F190ADE9-619A-447D-84E3-7E241A5C428E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCheckTypeToSchemeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCheckTypeToSchemeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCheckTypeToSchemeResponseBody) SetCode(v string) *UpdateCheckTypeToSchemeResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateCheckTypeToSchemeResponseBody) SetData(v int64) *UpdateCheckTypeToSchemeResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateCheckTypeToSchemeResponseBody) SetHttpStatusCode(v int32) *UpdateCheckTypeToSchemeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateCheckTypeToSchemeResponseBody) SetMessage(v string) *UpdateCheckTypeToSchemeResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateCheckTypeToSchemeResponseBody) SetMessages(v *UpdateCheckTypeToSchemeResponseBodyMessages) *UpdateCheckTypeToSchemeResponseBody {
	s.Messages = v
	return s
}

func (s *UpdateCheckTypeToSchemeResponseBody) SetRequestId(v string) *UpdateCheckTypeToSchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCheckTypeToSchemeResponseBody) SetSuccess(v bool) *UpdateCheckTypeToSchemeResponseBody {
	s.Success = &v
	return s
}

type UpdateCheckTypeToSchemeResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s UpdateCheckTypeToSchemeResponseBodyMessages) String() string {
	return tea.Prettify(s)
}

func (s UpdateCheckTypeToSchemeResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *UpdateCheckTypeToSchemeResponseBodyMessages) SetMessage(v []*string) *UpdateCheckTypeToSchemeResponseBodyMessages {
	s.Message = v
	return s
}

type UpdateCheckTypeToSchemeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCheckTypeToSchemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCheckTypeToSchemeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCheckTypeToSchemeResponse) GoString() string {
	return s.String()
}

func (s *UpdateCheckTypeToSchemeResponse) SetHeaders(v map[string]*string) *UpdateCheckTypeToSchemeResponse {
	s.Headers = v
	return s
}

func (s *UpdateCheckTypeToSchemeResponse) SetStatusCode(v int32) *UpdateCheckTypeToSchemeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCheckTypeToSchemeResponse) SetBody(v *UpdateCheckTypeToSchemeResponseBody) *UpdateCheckTypeToSchemeResponse {
	s.Body = v
	return s
}

type UpdateQualityCheckDataRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"taskId":"xxx"}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UpdateQualityCheckDataRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateQualityCheckDataRequest) GoString() string {
	return s.String()
}

func (s *UpdateQualityCheckDataRequest) SetBaseMeAgentId(v int64) *UpdateQualityCheckDataRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateQualityCheckDataRequest) SetJsonStr(v string) *UpdateQualityCheckDataRequest {
	s.JsonStr = &v
	return s
}

type UpdateQualityCheckDataResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateQualityCheckDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateQualityCheckDataResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateQualityCheckDataResponseBody) SetCode(v string) *UpdateQualityCheckDataResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateQualityCheckDataResponseBody) SetMessage(v string) *UpdateQualityCheckDataResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateQualityCheckDataResponseBody) SetRequestId(v string) *UpdateQualityCheckDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateQualityCheckDataResponseBody) SetSuccess(v bool) *UpdateQualityCheckDataResponseBody {
	s.Success = &v
	return s
}

type UpdateQualityCheckDataResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateQualityCheckDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateQualityCheckDataResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateQualityCheckDataResponse) GoString() string {
	return s.String()
}

func (s *UpdateQualityCheckDataResponse) SetHeaders(v map[string]*string) *UpdateQualityCheckDataResponse {
	s.Headers = v
	return s
}

func (s *UpdateQualityCheckDataResponse) SetStatusCode(v int32) *UpdateQualityCheckDataResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateQualityCheckDataResponse) SetBody(v *UpdateQualityCheckDataResponseBody) *UpdateQualityCheckDataResponse {
	s.Body = v
	return s
}

type UpdateQualityCheckSchemeRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64  `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	JsonStr       *string `json:"jsonStr,omitempty" xml:"jsonStr,omitempty"`
}

func (s UpdateQualityCheckSchemeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateQualityCheckSchemeRequest) GoString() string {
	return s.String()
}

func (s *UpdateQualityCheckSchemeRequest) SetBaseMeAgentId(v int64) *UpdateQualityCheckSchemeRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateQualityCheckSchemeRequest) SetJsonStr(v string) *UpdateQualityCheckSchemeRequest {
	s.JsonStr = &v
	return s
}

type UpdateQualityCheckSchemeResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *UpdateQualityCheckSchemeResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// 96138D8D-8D26-4E41-BFF4-77AED1088BBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateQualityCheckSchemeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateQualityCheckSchemeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateQualityCheckSchemeResponseBody) SetCode(v string) *UpdateQualityCheckSchemeResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateQualityCheckSchemeResponseBody) SetHttpStatusCode(v int32) *UpdateQualityCheckSchemeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateQualityCheckSchemeResponseBody) SetMessage(v string) *UpdateQualityCheckSchemeResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateQualityCheckSchemeResponseBody) SetMessages(v *UpdateQualityCheckSchemeResponseBodyMessages) *UpdateQualityCheckSchemeResponseBody {
	s.Messages = v
	return s
}

func (s *UpdateQualityCheckSchemeResponseBody) SetRequestId(v string) *UpdateQualityCheckSchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateQualityCheckSchemeResponseBody) SetSuccess(v bool) *UpdateQualityCheckSchemeResponseBody {
	s.Success = &v
	return s
}

type UpdateQualityCheckSchemeResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s UpdateQualityCheckSchemeResponseBodyMessages) String() string {
	return tea.Prettify(s)
}

func (s UpdateQualityCheckSchemeResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *UpdateQualityCheckSchemeResponseBodyMessages) SetMessage(v []*string) *UpdateQualityCheckSchemeResponseBodyMessages {
	s.Message = v
	return s
}

type UpdateQualityCheckSchemeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateQualityCheckSchemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateQualityCheckSchemeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateQualityCheckSchemeResponse) GoString() string {
	return s.String()
}

func (s *UpdateQualityCheckSchemeResponse) SetHeaders(v map[string]*string) *UpdateQualityCheckSchemeResponse {
	s.Headers = v
	return s
}

func (s *UpdateQualityCheckSchemeResponse) SetStatusCode(v int32) *UpdateQualityCheckSchemeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateQualityCheckSchemeResponse) SetBody(v *UpdateQualityCheckSchemeResponseBody) *UpdateQualityCheckSchemeResponse {
	s.Body = v
	return s
}

type UpdateRuleRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UpdateRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateRuleRequest) SetBaseMeAgentId(v int64) *UpdateRuleRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateRuleRequest) SetJsonStr(v string) *UpdateRuleRequest {
	s.JsonStr = &v
	return s
}

type UpdateRuleResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// xxx
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4987D326-83D9-4A42-B9A5-0B27F9B40539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRuleResponseBody) SetCode(v string) *UpdateRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRuleResponseBody) SetData(v string) *UpdateRuleResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateRuleResponseBody) SetMessage(v string) *UpdateRuleResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateRuleResponseBody) SetRequestId(v string) *UpdateRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRuleResponseBody) SetSuccess(v bool) *UpdateRuleResponseBody {
	s.Success = &v
	return s
}

type UpdateRuleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateRuleResponse) SetHeaders(v map[string]*string) *UpdateRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateRuleResponse) SetStatusCode(v int32) *UpdateRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRuleResponse) SetBody(v *UpdateRuleResponseBody) *UpdateRuleResponse {
	s.Body = v
	return s
}

type UpdateRuleByIdRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// false
	IsCopy *bool `json:"IsCopy,omitempty" xml:"IsCopy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {}
	JsonStrForRule *string `json:"JsonStrForRule,omitempty" xml:"JsonStrForRule,omitempty"`
	// example:
	//
	// 1
	ReturnRelatedSchemes *bool `json:"ReturnRelatedSchemes,omitempty" xml:"ReturnRelatedSchemes,omitempty"`
	// example:
	//
	// 1
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s UpdateRuleByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleByIdRequest) GoString() string {
	return s.String()
}

func (s *UpdateRuleByIdRequest) SetBaseMeAgentId(v int64) *UpdateRuleByIdRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateRuleByIdRequest) SetIsCopy(v bool) *UpdateRuleByIdRequest {
	s.IsCopy = &v
	return s
}

func (s *UpdateRuleByIdRequest) SetJsonStrForRule(v string) *UpdateRuleByIdRequest {
	s.JsonStrForRule = &v
	return s
}

func (s *UpdateRuleByIdRequest) SetReturnRelatedSchemes(v bool) *UpdateRuleByIdRequest {
	s.ReturnRelatedSchemes = &v
	return s
}

func (s *UpdateRuleByIdRequest) SetRuleId(v int64) *UpdateRuleByIdRequest {
	s.RuleId = &v
	return s
}

type UpdateRuleByIdResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *UpdateRuleByIdResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// 82C91484-B2D5-4D2A-A21F-A6D73F4D55C6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateRuleByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleByIdResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRuleByIdResponseBody) SetCode(v string) *UpdateRuleByIdResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRuleByIdResponseBody) SetHttpStatusCode(v int32) *UpdateRuleByIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateRuleByIdResponseBody) SetMessage(v string) *UpdateRuleByIdResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateRuleByIdResponseBody) SetMessages(v *UpdateRuleByIdResponseBodyMessages) *UpdateRuleByIdResponseBody {
	s.Messages = v
	return s
}

func (s *UpdateRuleByIdResponseBody) SetRequestId(v string) *UpdateRuleByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRuleByIdResponseBody) SetSuccess(v bool) *UpdateRuleByIdResponseBody {
	s.Success = &v
	return s
}

type UpdateRuleByIdResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s UpdateRuleByIdResponseBodyMessages) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleByIdResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *UpdateRuleByIdResponseBodyMessages) SetMessage(v []*string) *UpdateRuleByIdResponseBodyMessages {
	s.Message = v
	return s
}

type UpdateRuleByIdResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRuleByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRuleByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleByIdResponse) GoString() string {
	return s.String()
}

func (s *UpdateRuleByIdResponse) SetHeaders(v map[string]*string) *UpdateRuleByIdResponse {
	s.Headers = v
	return s
}

func (s *UpdateRuleByIdResponse) SetStatusCode(v int32) *UpdateRuleByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRuleByIdResponse) SetBody(v *UpdateRuleByIdResponseBody) *UpdateRuleByIdResponse {
	s.Body = v
	return s
}

type UpdateRuleToSchemeRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// {"schemeId":"10","schemeRules":[{"ruleId":229,"checkType":0}]}
	JsonStr *string `json:"jsonStr,omitempty" xml:"jsonStr,omitempty"`
}

func (s UpdateRuleToSchemeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleToSchemeRequest) GoString() string {
	return s.String()
}

func (s *UpdateRuleToSchemeRequest) SetBaseMeAgentId(v int64) *UpdateRuleToSchemeRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateRuleToSchemeRequest) SetJsonStr(v string) *UpdateRuleToSchemeRequest {
	s.JsonStr = &v
	return s
}

type UpdateRuleToSchemeResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 30
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *UpdateRuleToSchemeResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// 9987D326-83D9-4A42-B9A5-0B27F9B40539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateRuleToSchemeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleToSchemeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRuleToSchemeResponseBody) SetCode(v string) *UpdateRuleToSchemeResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRuleToSchemeResponseBody) SetData(v int64) *UpdateRuleToSchemeResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateRuleToSchemeResponseBody) SetHttpStatusCode(v int32) *UpdateRuleToSchemeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateRuleToSchemeResponseBody) SetMessage(v string) *UpdateRuleToSchemeResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateRuleToSchemeResponseBody) SetMessages(v *UpdateRuleToSchemeResponseBodyMessages) *UpdateRuleToSchemeResponseBody {
	s.Messages = v
	return s
}

func (s *UpdateRuleToSchemeResponseBody) SetRequestId(v string) *UpdateRuleToSchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRuleToSchemeResponseBody) SetSuccess(v bool) *UpdateRuleToSchemeResponseBody {
	s.Success = &v
	return s
}

type UpdateRuleToSchemeResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s UpdateRuleToSchemeResponseBodyMessages) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleToSchemeResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *UpdateRuleToSchemeResponseBodyMessages) SetMessage(v []*string) *UpdateRuleToSchemeResponseBodyMessages {
	s.Message = v
	return s
}

type UpdateRuleToSchemeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRuleToSchemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRuleToSchemeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleToSchemeResponse) GoString() string {
	return s.String()
}

func (s *UpdateRuleToSchemeResponse) SetHeaders(v map[string]*string) *UpdateRuleToSchemeResponse {
	s.Headers = v
	return s
}

func (s *UpdateRuleToSchemeResponse) SetStatusCode(v int32) *UpdateRuleToSchemeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRuleToSchemeResponse) SetBody(v *UpdateRuleToSchemeResponseBody) *UpdateRuleToSchemeResponse {
	s.Body = v
	return s
}

type UpdateRuleV4Request struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStrForRule *string `json:"JsonStrForRule,omitempty" xml:"JsonStrForRule,omitempty"`
	// example:
	//
	// 1
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s UpdateRuleV4Request) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleV4Request) GoString() string {
	return s.String()
}

func (s *UpdateRuleV4Request) SetBaseMeAgentId(v int64) *UpdateRuleV4Request {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateRuleV4Request) SetJsonStrForRule(v string) *UpdateRuleV4Request {
	s.JsonStrForRule = &v
	return s
}

func (s *UpdateRuleV4Request) SetRuleId(v int64) *UpdateRuleV4Request {
	s.RuleId = &v
	return s
}

type UpdateRuleV4ResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *UpdateRuleV4ResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateRuleV4ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleV4ResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRuleV4ResponseBody) SetCode(v string) *UpdateRuleV4ResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRuleV4ResponseBody) SetData(v int64) *UpdateRuleV4ResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateRuleV4ResponseBody) SetHttpStatusCode(v int32) *UpdateRuleV4ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateRuleV4ResponseBody) SetMessage(v string) *UpdateRuleV4ResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateRuleV4ResponseBody) SetMessages(v *UpdateRuleV4ResponseBodyMessages) *UpdateRuleV4ResponseBody {
	s.Messages = v
	return s
}

func (s *UpdateRuleV4ResponseBody) SetRequestId(v string) *UpdateRuleV4ResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRuleV4ResponseBody) SetSuccess(v bool) *UpdateRuleV4ResponseBody {
	s.Success = &v
	return s
}

type UpdateRuleV4ResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s UpdateRuleV4ResponseBodyMessages) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleV4ResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *UpdateRuleV4ResponseBodyMessages) SetMessage(v []*string) *UpdateRuleV4ResponseBodyMessages {
	s.Message = v
	return s
}

type UpdateRuleV4Response struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRuleV4ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRuleV4Response) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleV4Response) GoString() string {
	return s.String()
}

func (s *UpdateRuleV4Response) SetHeaders(v map[string]*string) *UpdateRuleV4Response {
	s.Headers = v
	return s
}

func (s *UpdateRuleV4Response) SetStatusCode(v int32) *UpdateRuleV4Response {
	s.StatusCode = &v
	return s
}

func (s *UpdateRuleV4Response) SetBody(v *UpdateRuleV4ResponseBody) *UpdateRuleV4Response {
	s.Body = v
	return s
}

type UpdateSchemeTaskConfigRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64  `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	JsonStr       *string `json:"jsonStr,omitempty" xml:"jsonStr,omitempty"`
}

func (s UpdateSchemeTaskConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSchemeTaskConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateSchemeTaskConfigRequest) SetBaseMeAgentId(v int64) *UpdateSchemeTaskConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateSchemeTaskConfigRequest) SetJsonStr(v string) *UpdateSchemeTaskConfigRequest {
	s.JsonStr = &v
	return s
}

type UpdateSchemeTaskConfigResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *UpdateSchemeTaskConfigResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateSchemeTaskConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSchemeTaskConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSchemeTaskConfigResponseBody) SetCode(v string) *UpdateSchemeTaskConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSchemeTaskConfigResponseBody) SetHttpStatusCode(v int32) *UpdateSchemeTaskConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateSchemeTaskConfigResponseBody) SetMessage(v string) *UpdateSchemeTaskConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSchemeTaskConfigResponseBody) SetMessages(v *UpdateSchemeTaskConfigResponseBodyMessages) *UpdateSchemeTaskConfigResponseBody {
	s.Messages = v
	return s
}

func (s *UpdateSchemeTaskConfigResponseBody) SetRequestId(v string) *UpdateSchemeTaskConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSchemeTaskConfigResponseBody) SetSuccess(v bool) *UpdateSchemeTaskConfigResponseBody {
	s.Success = &v
	return s
}

type UpdateSchemeTaskConfigResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s UpdateSchemeTaskConfigResponseBodyMessages) String() string {
	return tea.Prettify(s)
}

func (s UpdateSchemeTaskConfigResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *UpdateSchemeTaskConfigResponseBodyMessages) SetMessage(v []*string) *UpdateSchemeTaskConfigResponseBodyMessages {
	s.Message = v
	return s
}

type UpdateSchemeTaskConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSchemeTaskConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSchemeTaskConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSchemeTaskConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateSchemeTaskConfigResponse) SetHeaders(v map[string]*string) *UpdateSchemeTaskConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateSchemeTaskConfigResponse) SetStatusCode(v int32) *UpdateSchemeTaskConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSchemeTaskConfigResponse) SetBody(v *UpdateSchemeTaskConfigResponseBody) *UpdateSchemeTaskConfigResponse {
	s.Body = v
	return s
}

type UpdateScoreForApiRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UpdateScoreForApiRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateScoreForApiRequest) GoString() string {
	return s.String()
}

func (s *UpdateScoreForApiRequest) SetBaseMeAgentId(v int64) *UpdateScoreForApiRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateScoreForApiRequest) SetJsonStr(v string) *UpdateScoreForApiRequest {
	s.JsonStr = &v
	return s
}

type UpdateScoreForApiResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6987D326-83D9-4A42-B9A5-0B27F9B40539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateScoreForApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateScoreForApiResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateScoreForApiResponseBody) SetCode(v string) *UpdateScoreForApiResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateScoreForApiResponseBody) SetMessage(v string) *UpdateScoreForApiResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateScoreForApiResponseBody) SetRequestId(v string) *UpdateScoreForApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateScoreForApiResponseBody) SetSuccess(v bool) *UpdateScoreForApiResponseBody {
	s.Success = &v
	return s
}

type UpdateScoreForApiResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateScoreForApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateScoreForApiResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateScoreForApiResponse) GoString() string {
	return s.String()
}

func (s *UpdateScoreForApiResponse) SetHeaders(v map[string]*string) *UpdateScoreForApiResponse {
	s.Headers = v
	return s
}

func (s *UpdateScoreForApiResponse) SetStatusCode(v int32) *UpdateScoreForApiResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateScoreForApiResponse) SetBody(v *UpdateScoreForApiResponseBody) *UpdateScoreForApiResponse {
	s.Body = v
	return s
}

type UpdateSkillGroupConfigRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"skillGroupFrom":0,"name":"test","qualityCheckType":0,"rid":"2493,4098","vocabId":267,"skillGroupList":[{"skillGroupId":"090311","skillGroupName":"09031"}],"id":553}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UpdateSkillGroupConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSkillGroupConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateSkillGroupConfigRequest) SetBaseMeAgentId(v int64) *UpdateSkillGroupConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateSkillGroupConfigRequest) SetJsonStr(v string) *UpdateSkillGroupConfigRequest {
	s.JsonStr = &v
	return s
}

type UpdateSkillGroupConfigResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 38E7E948-0876-4FEE-B0AA-6*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateSkillGroupConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSkillGroupConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSkillGroupConfigResponseBody) SetCode(v string) *UpdateSkillGroupConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSkillGroupConfigResponseBody) SetMessage(v string) *UpdateSkillGroupConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSkillGroupConfigResponseBody) SetRequestId(v string) *UpdateSkillGroupConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSkillGroupConfigResponseBody) SetSuccess(v bool) *UpdateSkillGroupConfigResponseBody {
	s.Success = &v
	return s
}

type UpdateSkillGroupConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSkillGroupConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSkillGroupConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSkillGroupConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateSkillGroupConfigResponse) SetHeaders(v map[string]*string) *UpdateSkillGroupConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateSkillGroupConfigResponse) SetStatusCode(v int32) *UpdateSkillGroupConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSkillGroupConfigResponse) SetBody(v *UpdateSkillGroupConfigResponseBody) *UpdateSkillGroupConfigResponse {
	s.Body = v
	return s
}

type UpdateSubScoreForApiRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UpdateSubScoreForApiRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubScoreForApiRequest) GoString() string {
	return s.String()
}

func (s *UpdateSubScoreForApiRequest) SetBaseMeAgentId(v int64) *UpdateSubScoreForApiRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateSubScoreForApiRequest) SetJsonStr(v string) *UpdateSubScoreForApiRequest {
	s.JsonStr = &v
	return s
}

type UpdateSubScoreForApiResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9987D326-83Q9-4A42-B9A5-0B27F9B40539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateSubScoreForApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubScoreForApiResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSubScoreForApiResponseBody) SetCode(v string) *UpdateSubScoreForApiResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSubScoreForApiResponseBody) SetMessage(v string) *UpdateSubScoreForApiResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSubScoreForApiResponseBody) SetRequestId(v string) *UpdateSubScoreForApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSubScoreForApiResponseBody) SetSuccess(v bool) *UpdateSubScoreForApiResponseBody {
	s.Success = &v
	return s
}

type UpdateSubScoreForApiResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSubScoreForApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSubScoreForApiResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubScoreForApiResponse) GoString() string {
	return s.String()
}

func (s *UpdateSubScoreForApiResponse) SetHeaders(v map[string]*string) *UpdateSubScoreForApiResponse {
	s.Headers = v
	return s
}

func (s *UpdateSubScoreForApiResponse) SetStatusCode(v int32) *UpdateSubScoreForApiResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSubScoreForApiResponse) SetBody(v *UpdateSubScoreForApiResponseBody) *UpdateSubScoreForApiResponse {
	s.Body = v
	return s
}

type UpdateSyncQualityCheckDataRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"tid":"xxx"}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UpdateSyncQualityCheckDataRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSyncQualityCheckDataRequest) GoString() string {
	return s.String()
}

func (s *UpdateSyncQualityCheckDataRequest) SetBaseMeAgentId(v int64) *UpdateSyncQualityCheckDataRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateSyncQualityCheckDataRequest) SetJsonStr(v string) *UpdateSyncQualityCheckDataRequest {
	s.JsonStr = &v
	return s
}

type UpdateSyncQualityCheckDataResponseBody struct {
	// example:
	//
	// 200
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateSyncQualityCheckDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 76DB5D8C-5BD9-42A7-B527-5AF3A5F8***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateSyncQualityCheckDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSyncQualityCheckDataResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSyncQualityCheckDataResponseBody) SetCode(v string) *UpdateSyncQualityCheckDataResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSyncQualityCheckDataResponseBody) SetData(v *UpdateSyncQualityCheckDataResponseBodyData) *UpdateSyncQualityCheckDataResponseBody {
	s.Data = v
	return s
}

func (s *UpdateSyncQualityCheckDataResponseBody) SetMessage(v string) *UpdateSyncQualityCheckDataResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSyncQualityCheckDataResponseBody) SetRequestId(v string) *UpdateSyncQualityCheckDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSyncQualityCheckDataResponseBody) SetSuccess(v bool) *UpdateSyncQualityCheckDataResponseBody {
	s.Success = &v
	return s
}

type UpdateSyncQualityCheckDataResponseBodyData struct {
	// example:
	//
	// 123123D8C-5BD9-42A7-B527-1235F8**
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 20210101-1212121***
	Tid *string `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s UpdateSyncQualityCheckDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateSyncQualityCheckDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateSyncQualityCheckDataResponseBodyData) SetTaskId(v string) *UpdateSyncQualityCheckDataResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *UpdateSyncQualityCheckDataResponseBodyData) SetTid(v string) *UpdateSyncQualityCheckDataResponseBodyData {
	s.Tid = &v
	return s
}

type UpdateSyncQualityCheckDataResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSyncQualityCheckDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSyncQualityCheckDataResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSyncQualityCheckDataResponse) GoString() string {
	return s.String()
}

func (s *UpdateSyncQualityCheckDataResponse) SetHeaders(v map[string]*string) *UpdateSyncQualityCheckDataResponse {
	s.Headers = v
	return s
}

func (s *UpdateSyncQualityCheckDataResponse) SetStatusCode(v int32) *UpdateSyncQualityCheckDataResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSyncQualityCheckDataResponse) SetBody(v *UpdateSyncQualityCheckDataResponseBody) *UpdateSyncQualityCheckDataResponse {
	s.Body = v
	return s
}

type UpdateTaskAssignRuleRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UpdateTaskAssignRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskAssignRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskAssignRuleRequest) SetBaseMeAgentId(v int64) *UpdateTaskAssignRuleRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateTaskAssignRuleRequest) SetJsonStr(v string) *UpdateTaskAssignRuleRequest {
	s.JsonStr = &v
	return s
}

type UpdateTaskAssignRuleResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24F4CE647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTaskAssignRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskAssignRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskAssignRuleResponseBody) SetCode(v string) *UpdateTaskAssignRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTaskAssignRuleResponseBody) SetMessage(v string) *UpdateTaskAssignRuleResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTaskAssignRuleResponseBody) SetRequestId(v string) *UpdateTaskAssignRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTaskAssignRuleResponseBody) SetSuccess(v bool) *UpdateTaskAssignRuleResponseBody {
	s.Success = &v
	return s
}

type UpdateTaskAssignRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTaskAssignRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTaskAssignRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskAssignRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskAssignRuleResponse) SetHeaders(v map[string]*string) *UpdateTaskAssignRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskAssignRuleResponse) SetStatusCode(v int32) *UpdateTaskAssignRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskAssignRuleResponse) SetBody(v *UpdateTaskAssignRuleResponseBody) *UpdateTaskAssignRuleResponse {
	s.Body = v
	return s
}

type UpdateUserRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [{"aliUid":"2951869706989****","roleName":"ADMIN"},{"aliUid":"2557461687048****","roleName":"ADMIN"}]
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UpdateUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserRequest) SetBaseMeAgentId(v int64) *UpdateUserRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateUserRequest) SetJsonStr(v string) *UpdateUserRequest {
	s.JsonStr = &v
	return s
}

type UpdateUserResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24F4CE647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserResponseBody) SetCode(v string) *UpdateUserResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateUserResponseBody) SetMessage(v string) *UpdateUserResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateUserResponseBody) SetRequestId(v string) *UpdateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserResponseBody) SetSuccess(v bool) *UpdateUserResponseBody {
	s.Success = &v
	return s
}

type UpdateUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserResponse) SetHeaders(v map[string]*string) *UpdateUserResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserResponse) SetStatusCode(v int32) *UpdateUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserResponse) SetBody(v *UpdateUserResponseBody) *UpdateUserResponse {
	s.Body = v
	return s
}

type UpdateWarningConfigRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"ridList":[18130],"configName":"0310","channels":[{"type":1,"url":"https://sca.console.aliyun.com/#/warningConfig"}],"configId":29}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UpdateWarningConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWarningConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateWarningConfigRequest) SetBaseMeAgentId(v int64) *UpdateWarningConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateWarningConfigRequest) SetJsonStr(v string) *UpdateWarningConfigRequest {
	s.JsonStr = &v
	return s
}

type UpdateWarningConfigResponseBody struct {
	// example:
	//
	// 200
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24F4CE647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateWarningConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWarningConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWarningConfigResponseBody) SetCode(v string) *UpdateWarningConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateWarningConfigResponseBody) SetMessage(v string) *UpdateWarningConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateWarningConfigResponseBody) SetRequestId(v string) *UpdateWarningConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWarningConfigResponseBody) SetSuccess(v bool) *UpdateWarningConfigResponseBody {
	s.Success = &v
	return s
}

type UpdateWarningConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWarningConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWarningConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWarningConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateWarningConfigResponse) SetHeaders(v map[string]*string) *UpdateWarningConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateWarningConfigResponse) SetStatusCode(v int32) *UpdateWarningConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWarningConfigResponse) SetBody(v *UpdateWarningConfigResponseBody) *UpdateWarningConfigResponse {
	s.Body = v
	return s
}

type UpdateWarningStrategyConfigRequest struct {
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UpdateWarningStrategyConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWarningStrategyConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateWarningStrategyConfigRequest) SetBaseMeAgentId(v int64) *UpdateWarningStrategyConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateWarningStrategyConfigRequest) SetJsonStr(v string) *UpdateWarningStrategyConfigRequest {
	s.JsonStr = &v
	return s
}

type UpdateWarningStrategyConfigResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateWarningStrategyConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWarningStrategyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWarningStrategyConfigResponseBody) SetCode(v string) *UpdateWarningStrategyConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateWarningStrategyConfigResponseBody) SetData(v string) *UpdateWarningStrategyConfigResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateWarningStrategyConfigResponseBody) SetMessage(v string) *UpdateWarningStrategyConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateWarningStrategyConfigResponseBody) SetRequestId(v string) *UpdateWarningStrategyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWarningStrategyConfigResponseBody) SetSuccess(v bool) *UpdateWarningStrategyConfigResponseBody {
	s.Success = &v
	return s
}

type UpdateWarningStrategyConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWarningStrategyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWarningStrategyConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWarningStrategyConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateWarningStrategyConfigResponse) SetHeaders(v map[string]*string) *UpdateWarningStrategyConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateWarningStrategyConfigResponse) SetStatusCode(v int32) *UpdateWarningStrategyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWarningStrategyConfigResponse) SetBody(v *UpdateWarningStrategyConfigResponseBody) *UpdateWarningStrategyConfigResponse {
	s.Body = v
	return s
}

type UploadAudioDataRequest struct {
	// example:
	//
	// 123456
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {“callList”:“xxxxx”}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UploadAudioDataRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadAudioDataRequest) GoString() string {
	return s.String()
}

func (s *UploadAudioDataRequest) SetBaseMeAgentId(v int64) *UploadAudioDataRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UploadAudioDataRequest) SetJsonStr(v string) *UploadAudioDataRequest {
	s.JsonStr = &v
	return s
}

type UploadAudioDataResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 76DB5D8C-5BD9-42A7-B527-5AF3A5***
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 76DB5D8C-5BD9-42A7-B527-5AF3A5F8***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadAudioDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadAudioDataResponseBody) GoString() string {
	return s.String()
}

func (s *UploadAudioDataResponseBody) SetCode(v string) *UploadAudioDataResponseBody {
	s.Code = &v
	return s
}

func (s *UploadAudioDataResponseBody) SetData(v string) *UploadAudioDataResponseBody {
	s.Data = &v
	return s
}

func (s *UploadAudioDataResponseBody) SetMessage(v string) *UploadAudioDataResponseBody {
	s.Message = &v
	return s
}

func (s *UploadAudioDataResponseBody) SetRequestId(v string) *UploadAudioDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadAudioDataResponseBody) SetSuccess(v bool) *UploadAudioDataResponseBody {
	s.Success = &v
	return s
}

type UploadAudioDataResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadAudioDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadAudioDataResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadAudioDataResponse) GoString() string {
	return s.String()
}

func (s *UploadAudioDataResponse) SetHeaders(v map[string]*string) *UploadAudioDataResponse {
	s.Headers = v
	return s
}

func (s *UploadAudioDataResponse) SetStatusCode(v int32) *UploadAudioDataResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadAudioDataResponse) SetBody(v *UploadAudioDataResponseBody) *UploadAudioDataResponse {
	s.Body = v
	return s
}

type UploadDataRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UploadDataRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadDataRequest) GoString() string {
	return s.String()
}

func (s *UploadDataRequest) SetBaseMeAgentId(v int64) *UploadDataRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UploadDataRequest) SetJsonStr(v string) *UploadDataRequest {
	s.JsonStr = &v
	return s
}

type UploadDataResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 6F5934C7-C223-4F0F-BBF3-5B3594***
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6F5934C7-C223-4F0F-BBF3-5B3594****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadDataResponseBody) GoString() string {
	return s.String()
}

func (s *UploadDataResponseBody) SetCode(v string) *UploadDataResponseBody {
	s.Code = &v
	return s
}

func (s *UploadDataResponseBody) SetData(v string) *UploadDataResponseBody {
	s.Data = &v
	return s
}

func (s *UploadDataResponseBody) SetMessage(v string) *UploadDataResponseBody {
	s.Message = &v
	return s
}

func (s *UploadDataResponseBody) SetRequestId(v string) *UploadDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadDataResponseBody) SetSuccess(v bool) *UploadDataResponseBody {
	s.Success = &v
	return s
}

type UploadDataResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadDataResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadDataResponse) GoString() string {
	return s.String()
}

func (s *UploadDataResponse) SetHeaders(v map[string]*string) *UploadDataResponse {
	s.Headers = v
	return s
}

func (s *UploadDataResponse) SetStatusCode(v int32) *UploadDataResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadDataResponse) SetBody(v *UploadDataResponseBody) *UploadDataResponse {
	s.Body = v
	return s
}

type UploadDataSyncRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"tickets":xxx}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UploadDataSyncRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadDataSyncRequest) GoString() string {
	return s.String()
}

func (s *UploadDataSyncRequest) SetBaseMeAgentId(v int64) *UploadDataSyncRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UploadDataSyncRequest) SetJsonStr(v string) *UploadDataSyncRequest {
	s.JsonStr = &v
	return s
}

type UploadDataSyncResponseBody struct {
	// example:
	//
	// 200
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UploadDataSyncResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4987D326-83D9-4A42-B9A5-0B27F9B***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadDataSyncResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadDataSyncResponseBody) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBody) SetCode(v string) *UploadDataSyncResponseBody {
	s.Code = &v
	return s
}

func (s *UploadDataSyncResponseBody) SetData(v *UploadDataSyncResponseBodyData) *UploadDataSyncResponseBody {
	s.Data = v
	return s
}

func (s *UploadDataSyncResponseBody) SetMessage(v string) *UploadDataSyncResponseBody {
	s.Message = &v
	return s
}

func (s *UploadDataSyncResponseBody) SetRequestId(v string) *UploadDataSyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadDataSyncResponseBody) SetSuccess(v bool) *UploadDataSyncResponseBody {
	s.Success = &v
	return s
}

type UploadDataSyncResponseBodyData struct {
	ResultInfo []*UploadDataSyncResponseBodyDataResultInfo `json:"ResultInfo,omitempty" xml:"ResultInfo,omitempty" type:"Repeated"`
}

func (s UploadDataSyncResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UploadDataSyncResponseBodyData) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyData) SetResultInfo(v []*UploadDataSyncResponseBodyDataResultInfo) *UploadDataSyncResponseBodyData {
	s.ResultInfo = v
	return s
}

type UploadDataSyncResponseBodyDataResultInfo struct {
	HandScoreIdList *UploadDataSyncResponseBodyDataResultInfoHandScoreIdList `json:"HandScoreIdList,omitempty" xml:"HandScoreIdList,omitempty" type:"Struct"`
	Rules           *UploadDataSyncResponseBodyDataResultInfoRules           `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
	// example:
	//
	// 100
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s UploadDataSyncResponseBodyDataResultInfo) String() string {
	return tea.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfo) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfo) SetHandScoreIdList(v *UploadDataSyncResponseBodyDataResultInfoHandScoreIdList) *UploadDataSyncResponseBodyDataResultInfo {
	s.HandScoreIdList = v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfo) SetRules(v *UploadDataSyncResponseBodyDataResultInfoRules) *UploadDataSyncResponseBodyDataResultInfo {
	s.Rules = v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfo) SetScore(v int32) *UploadDataSyncResponseBodyDataResultInfo {
	s.Score = &v
	return s
}

type UploadDataSyncResponseBodyDataResultInfoHandScoreIdList struct {
	HandScoreIdList []*string `json:"HandScoreIdList,omitempty" xml:"HandScoreIdList,omitempty" type:"Repeated"`
}

func (s UploadDataSyncResponseBodyDataResultInfoHandScoreIdList) String() string {
	return tea.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfoHandScoreIdList) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfoHandScoreIdList) SetHandScoreIdList(v []*string) *UploadDataSyncResponseBodyDataResultInfoHandScoreIdList {
	s.HandScoreIdList = v
	return s
}

type UploadDataSyncResponseBodyDataResultInfoRules struct {
	RuleHitInfo []*UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfo `json:"RuleHitInfo,omitempty" xml:"RuleHitInfo,omitempty" type:"Repeated"`
}

func (s UploadDataSyncResponseBodyDataResultInfoRules) String() string {
	return tea.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfoRules) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfoRules) SetRuleHitInfo(v []*UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfo) *UploadDataSyncResponseBodyDataResultInfoRules {
	s.RuleHitInfo = v
	return s
}

type UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfo struct {
	ConditionInfo *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfo `json:"ConditionInfo,omitempty" xml:"ConditionInfo,omitempty" type:"Struct"`
	Hit           *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHit           `json:"Hit,omitempty" xml:"Hit,omitempty" type:"Struct"`
	// example:
	//
	// 801
	Rid *string `json:"Rid,omitempty" xml:"Rid,omitempty"`
	// example:
	//
	// 88888888
	Tid *string `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfo) String() string {
	return tea.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfo) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfo) SetConditionInfo(v *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfo) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfo {
	s.ConditionInfo = v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfo) SetHit(v *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHit) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfo {
	s.Hit = v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfo) SetRid(v string) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfo {
	s.Rid = &v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfo) SetTid(v string) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfo {
	s.Tid = &v
	return s
}

type UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfo struct {
	ConditionBasicInfo []*UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfoConditionBasicInfo `json:"ConditionBasicInfo,omitempty" xml:"ConditionBasicInfo,omitempty" type:"Repeated"`
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfo) String() string {
	return tea.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfo) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfo) SetConditionBasicInfo(v []*UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfoConditionBasicInfo) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfo {
	s.ConditionBasicInfo = v
	return s
}

type UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfoConditionBasicInfo struct {
	// example:
	//
	// xxx
	ConditionInfoCid *string `json:"ConditionInfoCid,omitempty" xml:"ConditionInfoCid,omitempty"`
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfoConditionBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfoConditionBasicInfo) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfoConditionBasicInfo) SetConditionInfoCid(v string) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfoConditionBasicInfo {
	s.ConditionInfoCid = &v
	return s
}

type UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHit struct {
	ConditionHitInfo []*UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo `json:"ConditionHitInfo,omitempty" xml:"ConditionHitInfo,omitempty" type:"Repeated"`
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHit) String() string {
	return tea.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHit) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHit) SetConditionHitInfo(v []*UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHit {
	s.ConditionHitInfo = v
	return s
}

type UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo struct {
	HitCids     *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitCids     `json:"HitCids,omitempty" xml:"HitCids,omitempty" type:"Struct"`
	HitKeyWords *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWords `json:"HitKeyWords,omitempty" xml:"HitKeyWords,omitempty" type:"Struct"`
	Phrase      *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase      `json:"Phrase,omitempty" xml:"Phrase,omitempty" type:"Struct"`
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo) String() string {
	return tea.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo) SetHitCids(v *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitCids) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo {
	s.HitCids = v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo) SetHitKeyWords(v *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWords) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo {
	s.HitKeyWords = v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo) SetPhrase(v *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo {
	s.Phrase = v
	return s
}

type UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitCids struct {
	CidItem []*string `json:"CidItem,omitempty" xml:"CidItem,omitempty" type:"Repeated"`
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitCids) String() string {
	return tea.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitCids) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitCids) SetCidItem(v []*string) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitCids {
	s.CidItem = v
	return s
}

type UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWords struct {
	HitKeyWord []*UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord `json:"HitKeyWord,omitempty" xml:"HitKeyWord,omitempty" type:"Repeated"`
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWords) String() string {
	return tea.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWords) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWords) SetHitKeyWord(v []*UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWords {
	s.HitKeyWord = v
	return s
}

type UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord struct {
	// example:
	//
	// 1
	From *int32 `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// 1
	Pid *int32 `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// example:
	//
	// 1
	Tid *string `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// example:
	//
	// 1
	To  *int32  `json:"To,omitempty" xml:"To,omitempty"`
	Val *string `json:"Val,omitempty" xml:"Val,omitempty"`
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) String() string {
	return tea.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) SetFrom(v int32) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord {
	s.From = &v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) SetPid(v int32) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord {
	s.Pid = &v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) SetTid(v string) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord {
	s.Tid = &v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) SetTo(v int32) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord {
	s.To = &v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) SetVal(v string) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord {
	s.Val = &v
	return s
}

type UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase struct {
	// example:
	//
	// 0
	Begin *int64 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// example:
	//
	// 1564574
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// 2090
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// example:
	//
	// 1
	Identity *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
	Role     *string `json:"Role,omitempty" xml:"Role,omitempty"`
	Words    *string `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) String() string {
	return tea.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) SetBegin(v int64) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase {
	s.Begin = &v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) SetBeginTime(v string) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase {
	s.BeginTime = &v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) SetEnd(v int64) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase {
	s.End = &v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) SetIdentity(v string) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase {
	s.Identity = &v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) SetRole(v string) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase {
	s.Role = &v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) SetWords(v string) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase {
	s.Words = &v
	return s
}

type UploadDataSyncResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadDataSyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadDataSyncResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadDataSyncResponse) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponse) SetHeaders(v map[string]*string) *UploadDataSyncResponse {
	s.Headers = v
	return s
}

func (s *UploadDataSyncResponse) SetStatusCode(v int32) *UploadDataSyncResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadDataSyncResponse) SetBody(v *UploadDataSyncResponseBody) *UploadDataSyncResponse {
	s.Body = v
	return s
}

type UploadDataV4Request struct {
	// example:
	//
	// 123456
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UploadDataV4Request) String() string {
	return tea.Prettify(s)
}

func (s UploadDataV4Request) GoString() string {
	return s.String()
}

func (s *UploadDataV4Request) SetBaseMeAgentId(v int64) *UploadDataV4Request {
	s.BaseMeAgentId = &v
	return s
}

func (s *UploadDataV4Request) SetJsonStr(v string) *UploadDataV4Request {
	s.JsonStr = &v
	return s
}

type UploadDataV4ResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 6F5934C7-C223-4F0F-BBF3-5B3594***
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6F5934C7-C223-4F0F-BBF3-5B3594***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadDataV4ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadDataV4ResponseBody) GoString() string {
	return s.String()
}

func (s *UploadDataV4ResponseBody) SetCode(v string) *UploadDataV4ResponseBody {
	s.Code = &v
	return s
}

func (s *UploadDataV4ResponseBody) SetData(v string) *UploadDataV4ResponseBody {
	s.Data = &v
	return s
}

func (s *UploadDataV4ResponseBody) SetMessage(v string) *UploadDataV4ResponseBody {
	s.Message = &v
	return s
}

func (s *UploadDataV4ResponseBody) SetRequestId(v string) *UploadDataV4ResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadDataV4ResponseBody) SetSuccess(v bool) *UploadDataV4ResponseBody {
	s.Success = &v
	return s
}

type UploadDataV4Response struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadDataV4ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadDataV4Response) String() string {
	return tea.Prettify(s)
}

func (s UploadDataV4Response) GoString() string {
	return s.String()
}

func (s *UploadDataV4Response) SetHeaders(v map[string]*string) *UploadDataV4Response {
	s.Headers = v
	return s
}

func (s *UploadDataV4Response) SetStatusCode(v int32) *UploadDataV4Response {
	s.StatusCode = &v
	return s
}

func (s *UploadDataV4Response) SetBody(v *UploadDataV4ResponseBody) *UploadDataV4Response {
	s.Body = v
	return s
}

type UploadRuleRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {“conditions”:“xxxxx”,"rules":"xxxx"}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UploadRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadRuleRequest) GoString() string {
	return s.String()
}

func (s *UploadRuleRequest) SetBaseMeAgentId(v int64) *UploadRuleRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UploadRuleRequest) SetJsonStr(v string) *UploadRuleRequest {
	s.JsonStr = &v
	return s
}

type UploadRuleResponseBody struct {
	// example:
	//
	// 200
	Code *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UploadRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4987D326-83D9-4A42-B9A5-0B27F9B40539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UploadRuleResponseBody) SetCode(v string) *UploadRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UploadRuleResponseBody) SetData(v *UploadRuleResponseBodyData) *UploadRuleResponseBody {
	s.Data = v
	return s
}

func (s *UploadRuleResponseBody) SetMessage(v string) *UploadRuleResponseBody {
	s.Message = &v
	return s
}

func (s *UploadRuleResponseBody) SetRequestId(v string) *UploadRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadRuleResponseBody) SetSuccess(v bool) *UploadRuleResponseBody {
	s.Success = &v
	return s
}

type UploadRuleResponseBodyData struct {
	RidInfo []*string `json:"RidInfo,omitempty" xml:"RidInfo,omitempty" type:"Repeated"`
}

func (s UploadRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UploadRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *UploadRuleResponseBodyData) SetRidInfo(v []*string) *UploadRuleResponseBodyData {
	s.RidInfo = v
	return s
}

type UploadRuleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadRuleResponse) GoString() string {
	return s.String()
}

func (s *UploadRuleResponse) SetHeaders(v map[string]*string) *UploadRuleResponse {
	s.Headers = v
	return s
}

func (s *UploadRuleResponse) SetStatusCode(v int32) *UploadRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadRuleResponse) SetBody(v *UploadRuleResponseBody) *UploadRuleResponse {
	s.Body = v
	return s
}

type VerifyFileRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s VerifyFileRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyFileRequest) GoString() string {
	return s.String()
}

func (s *VerifyFileRequest) SetBaseMeAgentId(v int64) *VerifyFileRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *VerifyFileRequest) SetJsonStr(v string) *VerifyFileRequest {
	s.JsonStr = &v
	return s
}

type VerifyFileResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 0.9485294
	Data *float32 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// s
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24F4CE647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s VerifyFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyFileResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyFileResponseBody) SetCode(v string) *VerifyFileResponseBody {
	s.Code = &v
	return s
}

func (s *VerifyFileResponseBody) SetData(v float32) *VerifyFileResponseBody {
	s.Data = &v
	return s
}

func (s *VerifyFileResponseBody) SetMessage(v string) *VerifyFileResponseBody {
	s.Message = &v
	return s
}

func (s *VerifyFileResponseBody) SetRequestId(v string) *VerifyFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyFileResponseBody) SetSuccess(v bool) *VerifyFileResponseBody {
	s.Success = &v
	return s
}

type VerifyFileResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyFileResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyFileResponse) GoString() string {
	return s.String()
}

func (s *VerifyFileResponse) SetHeaders(v map[string]*string) *VerifyFileResponse {
	s.Headers = v
	return s
}

func (s *VerifyFileResponse) SetStatusCode(v int32) *VerifyFileResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyFileResponse) SetBody(v *VerifyFileResponseBody) *VerifyFileResponse {
	s.Body = v
	return s
}

type VerifySentenceRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s VerifySentenceRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifySentenceRequest) GoString() string {
	return s.String()
}

func (s *VerifySentenceRequest) SetBaseMeAgentId(v int64) *VerifySentenceRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *VerifySentenceRequest) SetJsonStr(v string) *VerifySentenceRequest {
	s.JsonStr = &v
	return s
}

type VerifySentenceResponseBody struct {
	// example:
	//
	// 200
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *VerifySentenceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 2
	IncorrectWords *int32 `json:"IncorrectWords,omitempty" xml:"IncorrectWords,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24F4CE647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 0
	SourceRole *int32 `json:"SourceRole,omitempty" xml:"SourceRole,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1
	TargetRole *int32 `json:"TargetRole,omitempty" xml:"TargetRole,omitempty"`
}

func (s VerifySentenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifySentenceResponseBody) GoString() string {
	return s.String()
}

func (s *VerifySentenceResponseBody) SetCode(v string) *VerifySentenceResponseBody {
	s.Code = &v
	return s
}

func (s *VerifySentenceResponseBody) SetData(v *VerifySentenceResponseBodyData) *VerifySentenceResponseBody {
	s.Data = v
	return s
}

func (s *VerifySentenceResponseBody) SetIncorrectWords(v int32) *VerifySentenceResponseBody {
	s.IncorrectWords = &v
	return s
}

func (s *VerifySentenceResponseBody) SetMessage(v string) *VerifySentenceResponseBody {
	s.Message = &v
	return s
}

func (s *VerifySentenceResponseBody) SetRequestId(v string) *VerifySentenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifySentenceResponseBody) SetSourceRole(v int32) *VerifySentenceResponseBody {
	s.SourceRole = &v
	return s
}

func (s *VerifySentenceResponseBody) SetSuccess(v bool) *VerifySentenceResponseBody {
	s.Success = &v
	return s
}

func (s *VerifySentenceResponseBody) SetTargetRole(v int32) *VerifySentenceResponseBody {
	s.TargetRole = &v
	return s
}

type VerifySentenceResponseBodyData struct {
	Delta []*VerifySentenceResponseBodyDataDelta `json:"Delta,omitempty" xml:"Delta,omitempty" type:"Repeated"`
}

func (s VerifySentenceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s VerifySentenceResponseBodyData) GoString() string {
	return s.String()
}

func (s *VerifySentenceResponseBodyData) SetDelta(v []*VerifySentenceResponseBodyDataDelta) *VerifySentenceResponseBodyData {
	s.Delta = v
	return s
}

type VerifySentenceResponseBodyDataDelta struct {
	Source *VerifySentenceResponseBodyDataDeltaSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	Target *VerifySentenceResponseBodyDataDeltaTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// example:
	//
	// CHANGE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s VerifySentenceResponseBodyDataDelta) String() string {
	return tea.Prettify(s)
}

func (s VerifySentenceResponseBodyDataDelta) GoString() string {
	return s.String()
}

func (s *VerifySentenceResponseBodyDataDelta) SetSource(v *VerifySentenceResponseBodyDataDeltaSource) *VerifySentenceResponseBodyDataDelta {
	s.Source = v
	return s
}

func (s *VerifySentenceResponseBodyDataDelta) SetTarget(v *VerifySentenceResponseBodyDataDeltaTarget) *VerifySentenceResponseBodyDataDelta {
	s.Target = v
	return s
}

func (s *VerifySentenceResponseBodyDataDelta) SetType(v string) *VerifySentenceResponseBodyDataDelta {
	s.Type = &v
	return s
}

type VerifySentenceResponseBodyDataDeltaSource struct {
	Line *VerifySentenceResponseBodyDataDeltaSourceLine `json:"Line,omitempty" xml:"Line,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Position *int32 `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s VerifySentenceResponseBodyDataDeltaSource) String() string {
	return tea.Prettify(s)
}

func (s VerifySentenceResponseBodyDataDeltaSource) GoString() string {
	return s.String()
}

func (s *VerifySentenceResponseBodyDataDeltaSource) SetLine(v *VerifySentenceResponseBodyDataDeltaSourceLine) *VerifySentenceResponseBodyDataDeltaSource {
	s.Line = v
	return s
}

func (s *VerifySentenceResponseBodyDataDeltaSource) SetPosition(v int32) *VerifySentenceResponseBodyDataDeltaSource {
	s.Position = &v
	return s
}

type VerifySentenceResponseBodyDataDeltaSourceLine struct {
	Line []*string `json:"Line,omitempty" xml:"Line,omitempty" type:"Repeated"`
}

func (s VerifySentenceResponseBodyDataDeltaSourceLine) String() string {
	return tea.Prettify(s)
}

func (s VerifySentenceResponseBodyDataDeltaSourceLine) GoString() string {
	return s.String()
}

func (s *VerifySentenceResponseBodyDataDeltaSourceLine) SetLine(v []*string) *VerifySentenceResponseBodyDataDeltaSourceLine {
	s.Line = v
	return s
}

type VerifySentenceResponseBodyDataDeltaTarget struct {
	Line *VerifySentenceResponseBodyDataDeltaTargetLine `json:"Line,omitempty" xml:"Line,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Position *int32 `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s VerifySentenceResponseBodyDataDeltaTarget) String() string {
	return tea.Prettify(s)
}

func (s VerifySentenceResponseBodyDataDeltaTarget) GoString() string {
	return s.String()
}

func (s *VerifySentenceResponseBodyDataDeltaTarget) SetLine(v *VerifySentenceResponseBodyDataDeltaTargetLine) *VerifySentenceResponseBodyDataDeltaTarget {
	s.Line = v
	return s
}

func (s *VerifySentenceResponseBodyDataDeltaTarget) SetPosition(v int32) *VerifySentenceResponseBodyDataDeltaTarget {
	s.Position = &v
	return s
}

type VerifySentenceResponseBodyDataDeltaTargetLine struct {
	Line []*string `json:"Line,omitempty" xml:"Line,omitempty" type:"Repeated"`
}

func (s VerifySentenceResponseBodyDataDeltaTargetLine) String() string {
	return tea.Prettify(s)
}

func (s VerifySentenceResponseBodyDataDeltaTargetLine) GoString() string {
	return s.String()
}

func (s *VerifySentenceResponseBodyDataDeltaTargetLine) SetLine(v []*string) *VerifySentenceResponseBodyDataDeltaTargetLine {
	s.Line = v
	return s
}

type VerifySentenceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifySentenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifySentenceResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifySentenceResponse) GoString() string {
	return s.String()
}

func (s *VerifySentenceResponse) SetHeaders(v map[string]*string) *VerifySentenceResponse {
	s.Headers = v
	return s
}

func (s *VerifySentenceResponse) SetStatusCode(v int32) *VerifySentenceResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifySentenceResponse) SetBody(v *VerifySentenceResponseBody) *VerifySentenceResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("qualitycheck"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AddBusinessCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddBusinessCategoryResponse
func (client *Client) AddBusinessCategoryWithOptions(request *AddBusinessCategoryRequest, runtime *util.RuntimeOptions) (_result *AddBusinessCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddBusinessCategory"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddBusinessCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddBusinessCategoryRequest
//
// @return AddBusinessCategoryResponse
func (client *Client) AddBusinessCategory(request *AddBusinessCategoryRequest) (_result *AddBusinessCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddBusinessCategoryResponse{}
	_body, _err := client.AddBusinessCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AddRuleCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRuleCategoryResponse
func (client *Client) AddRuleCategoryWithOptions(request *AddRuleCategoryRequest, runtime *util.RuntimeOptions) (_result *AddRuleCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddRuleCategory"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddRuleCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddRuleCategoryRequest
//
// @return AddRuleCategoryResponse
func (client *Client) AddRuleCategory(request *AddRuleCategoryRequest) (_result *AddRuleCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddRuleCategoryResponse{}
	_body, _err := client.AddRuleCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// V4创建规则
//
// @param request - AddRuleV4Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRuleV4Response
func (client *Client) AddRuleV4WithOptions(request *AddRuleV4Request, runtime *util.RuntimeOptions) (_result *AddRuleV4Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsCopy)) {
		body["IsCopy"] = request.IsCopy
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStrForRule)) {
		body["JsonStrForRule"] = request.JsonStrForRule
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddRuleV4"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddRuleV4Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// V4创建规则
//
// @param request - AddRuleV4Request
//
// @return AddRuleV4Response
func (client *Client) AddRuleV4(request *AddRuleV4Request) (_result *AddRuleV4Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddRuleV4Response{}
	_body, _err := client.AddRuleV4WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AssignReviewerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssignReviewerResponse
func (client *Client) AssignReviewerWithOptions(request *AssignReviewerRequest, runtime *util.RuntimeOptions) (_result *AssignReviewerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AssignReviewer"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssignReviewerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - AssignReviewerRequest
//
// @return AssignReviewerResponse
func (client *Client) AssignReviewer(request *AssignReviewerRequest) (_result *AssignReviewerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssignReviewerResponse{}
	_body, _err := client.AssignReviewerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 会话组批量分配
//
// @param request - AssignReviewerBySessionGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssignReviewerBySessionGroupResponse
func (client *Client) AssignReviewerBySessionGroupWithOptions(request *AssignReviewerBySessionGroupRequest, runtime *util.RuntimeOptions) (_result *AssignReviewerBySessionGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AssignReviewerBySessionGroup"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssignReviewerBySessionGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 会话组批量分配
//
// @param request - AssignReviewerBySessionGroupRequest
//
// @return AssignReviewerBySessionGroupResponse
func (client *Client) AssignReviewerBySessionGroup(request *AssignReviewerBySessionGroupRequest) (_result *AssignReviewerBySessionGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssignReviewerBySessionGroupResponse{}
	_body, _err := client.AssignReviewerBySessionGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量复核
//
// @param request - BatchSubmitReviewInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchSubmitReviewInfoResponse
func (client *Client) BatchSubmitReviewInfoWithOptions(request *BatchSubmitReviewInfoRequest, runtime *util.RuntimeOptions) (_result *BatchSubmitReviewInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchSubmitReviewInfo"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchSubmitReviewInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量复核
//
// @param request - BatchSubmitReviewInfoRequest
//
// @return BatchSubmitReviewInfoResponse
func (client *Client) BatchSubmitReviewInfo(request *BatchSubmitReviewInfoRequest) (_result *BatchSubmitReviewInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchSubmitReviewInfoResponse{}
	_body, _err := client.BatchSubmitReviewInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateAsrVocabRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAsrVocabResponse
func (client *Client) CreateAsrVocabWithOptions(request *CreateAsrVocabRequest, runtime *util.RuntimeOptions) (_result *CreateAsrVocabResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAsrVocab"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAsrVocabResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateAsrVocabRequest
//
// @return CreateAsrVocabResponse
func (client *Client) CreateAsrVocab(request *CreateAsrVocabRequest) (_result *CreateAsrVocabResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAsrVocabResponse{}
	_body, _err := client.CreateAsrVocabWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建质检方案中的质检维度
//
// @param request - CreateCheckTypeToSchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCheckTypeToSchemeResponse
func (client *Client) CreateCheckTypeToSchemeWithOptions(request *CreateCheckTypeToSchemeRequest, runtime *util.RuntimeOptions) (_result *CreateCheckTypeToSchemeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCheckTypeToScheme"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCheckTypeToSchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建质检方案中的质检维度
//
// @param request - CreateCheckTypeToSchemeRequest
//
// @return CreateCheckTypeToSchemeResponse
func (client *Client) CreateCheckTypeToScheme(request *CreateCheckTypeToSchemeRequest) (_result *CreateCheckTypeToSchemeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCheckTypeToSchemeResponse{}
	_body, _err := client.CreateCheckTypeToSchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增质检方案
//
// @param request - CreateQualityCheckSchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQualityCheckSchemeResponse
func (client *Client) CreateQualityCheckSchemeWithOptions(request *CreateQualityCheckSchemeRequest, runtime *util.RuntimeOptions) (_result *CreateQualityCheckSchemeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateQualityCheckScheme"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateQualityCheckSchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增质检方案
//
// @param request - CreateQualityCheckSchemeRequest
//
// @return CreateQualityCheckSchemeResponse
func (client *Client) CreateQualityCheckScheme(request *CreateQualityCheckSchemeRequest) (_result *CreateQualityCheckSchemeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateQualityCheckSchemeResponse{}
	_body, _err := client.CreateQualityCheckSchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新建质检任务
//
// @param request - CreateSchemeTaskConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSchemeTaskConfigResponse
func (client *Client) CreateSchemeTaskConfigWithOptions(request *CreateSchemeTaskConfigRequest, runtime *util.RuntimeOptions) (_result *CreateSchemeTaskConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSchemeTaskConfig"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSchemeTaskConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新建质检任务
//
// @param request - CreateSchemeTaskConfigRequest
//
// @return CreateSchemeTaskConfigResponse
func (client *Client) CreateSchemeTaskConfig(request *CreateSchemeTaskConfigRequest) (_result *CreateSchemeTaskConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSchemeTaskConfigResponse{}
	_body, _err := client.CreateSchemeTaskConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateSkillGroupConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSkillGroupConfigResponse
func (client *Client) CreateSkillGroupConfigWithOptions(request *CreateSkillGroupConfigRequest, runtime *util.RuntimeOptions) (_result *CreateSkillGroupConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSkillGroupConfig"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSkillGroupConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateSkillGroupConfigRequest
//
// @return CreateSkillGroupConfigResponse
func (client *Client) CreateSkillGroupConfig(request *CreateSkillGroupConfigRequest) (_result *CreateSkillGroupConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSkillGroupConfigResponse{}
	_body, _err := client.CreateSkillGroupConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateTaskAssignRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTaskAssignRuleResponse
func (client *Client) CreateTaskAssignRuleWithOptions(request *CreateTaskAssignRuleRequest, runtime *util.RuntimeOptions) (_result *CreateTaskAssignRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTaskAssignRule"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTaskAssignRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateTaskAssignRuleRequest
//
// @return CreateTaskAssignRuleResponse
func (client *Client) CreateTaskAssignRule(request *CreateTaskAssignRuleRequest) (_result *CreateTaskAssignRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTaskAssignRuleResponse{}
	_body, _err := client.CreateTaskAssignRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI CreateUser is deprecated
//
// @param request - CreateUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserResponse
// Deprecated
func (client *Client) CreateUserWithOptions(request *CreateUserRequest, runtime *util.RuntimeOptions) (_result *CreateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUser"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI CreateUser is deprecated
//
// @param request - CreateUserRequest
//
// @return CreateUserResponse
// Deprecated
func (client *Client) CreateUser(request *CreateUserRequest) (_result *CreateUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUserResponse{}
	_body, _err := client.CreateUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateWarningConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWarningConfigResponse
func (client *Client) CreateWarningConfigWithOptions(request *CreateWarningConfigRequest, runtime *util.RuntimeOptions) (_result *CreateWarningConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWarningConfig"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWarningConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateWarningConfigRequest
//
// @return CreateWarningConfigResponse
func (client *Client) CreateWarningConfig(request *CreateWarningConfigRequest) (_result *CreateWarningConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWarningConfigResponse{}
	_body, _err := client.CreateWarningConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 预警策略-新增
//
// @param request - CreateWarningStrategyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWarningStrategyConfigResponse
func (client *Client) CreateWarningStrategyConfigWithOptions(request *CreateWarningStrategyConfigRequest, runtime *util.RuntimeOptions) (_result *CreateWarningStrategyConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWarningStrategyConfig"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWarningStrategyConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预警策略-新增
//
// @param request - CreateWarningStrategyConfigRequest
//
// @return CreateWarningStrategyConfigResponse
func (client *Client) CreateWarningStrategyConfig(request *CreateWarningStrategyConfigRequest) (_result *CreateWarningStrategyConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWarningStrategyConfigResponse{}
	_body, _err := client.CreateWarningStrategyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DelRuleCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DelRuleCategoryResponse
func (client *Client) DelRuleCategoryWithOptions(request *DelRuleCategoryRequest, runtime *util.RuntimeOptions) (_result *DelRuleCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DelRuleCategory"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DelRuleCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DelRuleCategoryRequest
//
// @return DelRuleCategoryResponse
func (client *Client) DelRuleCategory(request *DelRuleCategoryRequest) (_result *DelRuleCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DelRuleCategoryResponse{}
	_body, _err := client.DelRuleCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteAsrVocabRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAsrVocabResponse
func (client *Client) DeleteAsrVocabWithOptions(request *DeleteAsrVocabRequest, runtime *util.RuntimeOptions) (_result *DeleteAsrVocabResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAsrVocab"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAsrVocabResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteAsrVocabRequest
//
// @return DeleteAsrVocabResponse
func (client *Client) DeleteAsrVocab(request *DeleteAsrVocabRequest) (_result *DeleteAsrVocabResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAsrVocabResponse{}
	_body, _err := client.DeleteAsrVocabWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteBusinessCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBusinessCategoryResponse
func (client *Client) DeleteBusinessCategoryWithOptions(request *DeleteBusinessCategoryRequest, runtime *util.RuntimeOptions) (_result *DeleteBusinessCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBusinessCategory"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteBusinessCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteBusinessCategoryRequest
//
// @return DeleteBusinessCategoryResponse
func (client *Client) DeleteBusinessCategory(request *DeleteBusinessCategoryRequest) (_result *DeleteBusinessCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBusinessCategoryResponse{}
	_body, _err := client.DeleteBusinessCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteCustomizationConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomizationConfigResponse
func (client *Client) DeleteCustomizationConfigWithOptions(request *DeleteCustomizationConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteCustomizationConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCustomizationConfig"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCustomizationConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteCustomizationConfigRequest
//
// @return DeleteCustomizationConfigResponse
func (client *Client) DeleteCustomizationConfig(request *DeleteCustomizationConfigRequest) (_result *DeleteCustomizationConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCustomizationConfigResponse{}
	_body, _err := client.DeleteCustomizationConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteDataSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataSetResponse
func (client *Client) DeleteDataSetWithOptions(request *DeleteDataSetRequest, runtime *util.RuntimeOptions) (_result *DeleteDataSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDataSet"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDataSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteDataSetRequest
//
// @return DeleteDataSetResponse
func (client *Client) DeleteDataSet(request *DeleteDataSetRequest) (_result *DeleteDataSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDataSetResponse{}
	_body, _err := client.DeleteDataSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeletePrecisionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePrecisionTaskResponse
func (client *Client) DeletePrecisionTaskWithOptions(request *DeletePrecisionTaskRequest, runtime *util.RuntimeOptions) (_result *DeletePrecisionTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePrecisionTask"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePrecisionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeletePrecisionTaskRequest
//
// @return DeletePrecisionTaskResponse
func (client *Client) DeletePrecisionTask(request *DeletePrecisionTaskRequest) (_result *DeletePrecisionTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePrecisionTaskResponse{}
	_body, _err := client.DeletePrecisionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除质检方案
//
// @param request - DeleteQualityCheckSchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQualityCheckSchemeResponse
func (client *Client) DeleteQualityCheckSchemeWithOptions(request *DeleteQualityCheckSchemeRequest, runtime *util.RuntimeOptions) (_result *DeleteQualityCheckSchemeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteQualityCheckScheme"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteQualityCheckSchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除质检方案
//
// @param request - DeleteQualityCheckSchemeRequest
//
// @return DeleteQualityCheckSchemeResponse
func (client *Client) DeleteQualityCheckScheme(request *DeleteQualityCheckSchemeRequest) (_result *DeleteQualityCheckSchemeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteQualityCheckSchemeResponse{}
	_body, _err := client.DeleteQualityCheckSchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DeleteRule is deprecated, please use Qualitycheck::2019-01-15::DeleteRuleV4 instead.
//
// Summary:
//
// 删除规则
//
// @param request - DeleteRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRuleResponse
// Deprecated
func (client *Client) DeleteRuleWithOptions(request *DeleteRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ForceDelete)) {
		body["ForceDelete"] = request.ForceDelete
	}

	if !tea.BoolValue(util.IsUnset(request.IsSchemeData)) {
		body["IsSchemeData"] = request.IsSchemeData
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		body["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRule"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DeleteRule is deprecated, please use Qualitycheck::2019-01-15::DeleteRuleV4 instead.
//
// Summary:
//
// 删除规则
//
// @param request - DeleteRuleRequest
//
// @return DeleteRuleResponse
// Deprecated
func (client *Client) DeleteRule(request *DeleteRuleRequest) (_result *DeleteRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRuleResponse{}
	_body, _err := client.DeleteRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// V4删除规则
//
// @param request - DeleteRuleV4Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRuleV4Response
func (client *Client) DeleteRuleV4WithOptions(request *DeleteRuleV4Request, runtime *util.RuntimeOptions) (_result *DeleteRuleV4Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ForceDelete)) {
		body["ForceDelete"] = request.ForceDelete
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		body["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRuleV4"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRuleV4Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// V4删除规则
//
// @param request - DeleteRuleV4Request
//
// @return DeleteRuleV4Response
func (client *Client) DeleteRuleV4(request *DeleteRuleV4Request) (_result *DeleteRuleV4Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRuleV4Response{}
	_body, _err := client.DeleteRuleV4WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除质检任务
//
// @param request - DeleteSchemeTaskConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSchemeTaskConfigResponse
func (client *Client) DeleteSchemeTaskConfigWithOptions(request *DeleteSchemeTaskConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteSchemeTaskConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSchemeTaskConfig"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSchemeTaskConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除质检任务
//
// @param request - DeleteSchemeTaskConfigRequest
//
// @return DeleteSchemeTaskConfigResponse
func (client *Client) DeleteSchemeTaskConfig(request *DeleteSchemeTaskConfigRequest) (_result *DeleteSchemeTaskConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSchemeTaskConfigResponse{}
	_body, _err := client.DeleteSchemeTaskConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DeleteScoreForApi is deprecated
//
// @param request - DeleteScoreForApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScoreForApiResponse
// Deprecated
func (client *Client) DeleteScoreForApiWithOptions(request *DeleteScoreForApiRequest, runtime *util.RuntimeOptions) (_result *DeleteScoreForApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteScoreForApi"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteScoreForApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DeleteScoreForApi is deprecated
//
// @param request - DeleteScoreForApiRequest
//
// @return DeleteScoreForApiResponse
// Deprecated
func (client *Client) DeleteScoreForApi(request *DeleteScoreForApiRequest) (_result *DeleteScoreForApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteScoreForApiResponse{}
	_body, _err := client.DeleteScoreForApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteSkillGroupConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSkillGroupConfigResponse
func (client *Client) DeleteSkillGroupConfigWithOptions(request *DeleteSkillGroupConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteSkillGroupConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSkillGroupConfig"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSkillGroupConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteSkillGroupConfigRequest
//
// @return DeleteSkillGroupConfigResponse
func (client *Client) DeleteSkillGroupConfig(request *DeleteSkillGroupConfigRequest) (_result *DeleteSkillGroupConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSkillGroupConfigResponse{}
	_body, _err := client.DeleteSkillGroupConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DeleteSubScoreForApi is deprecated
//
// @param request - DeleteSubScoreForApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSubScoreForApiResponse
// Deprecated
func (client *Client) DeleteSubScoreForApiWithOptions(request *DeleteSubScoreForApiRequest, runtime *util.RuntimeOptions) (_result *DeleteSubScoreForApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSubScoreForApi"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSubScoreForApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DeleteSubScoreForApi is deprecated
//
// @param request - DeleteSubScoreForApiRequest
//
// @return DeleteSubScoreForApiResponse
// Deprecated
func (client *Client) DeleteSubScoreForApi(request *DeleteSubScoreForApiRequest) (_result *DeleteSubScoreForApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSubScoreForApiResponse{}
	_body, _err := client.DeleteSubScoreForApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteTaskAssignRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTaskAssignRuleResponse
func (client *Client) DeleteTaskAssignRuleWithOptions(request *DeleteTaskAssignRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteTaskAssignRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTaskAssignRule"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTaskAssignRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteTaskAssignRuleRequest
//
// @return DeleteTaskAssignRuleResponse
func (client *Client) DeleteTaskAssignRule(request *DeleteTaskAssignRuleRequest) (_result *DeleteTaskAssignRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTaskAssignRuleResponse{}
	_body, _err := client.DeleteTaskAssignRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteWarningConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWarningConfigResponse
func (client *Client) DeleteWarningConfigWithOptions(request *DeleteWarningConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteWarningConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteWarningConfig"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteWarningConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteWarningConfigRequest
//
// @return DeleteWarningConfigResponse
func (client *Client) DeleteWarningConfig(request *DeleteWarningConfigRequest) (_result *DeleteWarningConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWarningConfigResponse{}
	_body, _err := client.DeleteWarningConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 预警策略-删除
//
// @param request - DeleteWarningStrategyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWarningStrategyConfigResponse
func (client *Client) DeleteWarningStrategyConfigWithOptions(request *DeleteWarningStrategyConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteWarningStrategyConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteWarningStrategyConfig"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteWarningStrategyConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预警策略-删除
//
// @param request - DeleteWarningStrategyConfigRequest
//
// @return DeleteWarningStrategyConfigResponse
func (client *Client) DeleteWarningStrategyConfig(request *DeleteWarningStrategyConfigRequest) (_result *DeleteWarningStrategyConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWarningStrategyConfigResponse{}
	_body, _err := client.DeleteWarningStrategyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetAsrVocabRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAsrVocabResponse
func (client *Client) GetAsrVocabWithOptions(request *GetAsrVocabRequest, runtime *util.RuntimeOptions) (_result *GetAsrVocabResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAsrVocab"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAsrVocabResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetAsrVocabRequest
//
// @return GetAsrVocabResponse
func (client *Client) GetAsrVocab(request *GetAsrVocabRequest) (_result *GetAsrVocabResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAsrVocabResponse{}
	_body, _err := client.GetAsrVocabWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetBusinessCategoryListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBusinessCategoryListResponse
func (client *Client) GetBusinessCategoryListWithOptions(request *GetBusinessCategoryListRequest, runtime *util.RuntimeOptions) (_result *GetBusinessCategoryListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBusinessCategoryList"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBusinessCategoryListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetBusinessCategoryListRequest
//
// @return GetBusinessCategoryListResponse
func (client *Client) GetBusinessCategoryList(request *GetBusinessCategoryListRequest) (_result *GetBusinessCategoryListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBusinessCategoryListResponse{}
	_body, _err := client.GetBusinessCategoryListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// GetCustomizationConfigList HSF_HTTP
//
// @param request - GetCustomizationConfigListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomizationConfigListResponse
func (client *Client) GetCustomizationConfigListWithOptions(request *GetCustomizationConfigListRequest, runtime *util.RuntimeOptions) (_result *GetCustomizationConfigListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCustomizationConfigList"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCustomizationConfigListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// GetCustomizationConfigList HSF_HTTP
//
// @param request - GetCustomizationConfigListRequest
//
// @return GetCustomizationConfigListResponse
func (client *Client) GetCustomizationConfigList(request *GetCustomizationConfigListRequest) (_result *GetCustomizationConfigListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCustomizationConfigListResponse{}
	_body, _err := client.GetCustomizationConfigListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetNextResultToVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNextResultToVerifyResponse
func (client *Client) GetNextResultToVerifyWithOptions(request *GetNextResultToVerifyRequest, runtime *util.RuntimeOptions) (_result *GetNextResultToVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNextResultToVerify"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNextResultToVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetNextResultToVerifyRequest
//
// @return GetNextResultToVerifyResponse
func (client *Client) GetNextResultToVerify(request *GetNextResultToVerifyRequest) (_result *GetNextResultToVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNextResultToVerifyResponse{}
	_body, _err := client.GetNextResultToVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetPrecisionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPrecisionTaskResponse
func (client *Client) GetPrecisionTaskWithOptions(request *GetPrecisionTaskRequest, runtime *util.RuntimeOptions) (_result *GetPrecisionTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPrecisionTask"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPrecisionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetPrecisionTaskRequest
//
// @return GetPrecisionTaskResponse
func (client *Client) GetPrecisionTask(request *GetPrecisionTaskRequest) (_result *GetPrecisionTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPrecisionTaskResponse{}
	_body, _err := client.GetPrecisionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取质检方案
//
// @param request - GetQualityCheckSchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityCheckSchemeResponse
func (client *Client) GetQualityCheckSchemeWithOptions(request *GetQualityCheckSchemeRequest, runtime *util.RuntimeOptions) (_result *GetQualityCheckSchemeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQualityCheckScheme"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetQualityCheckSchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取质检方案
//
// @param request - GetQualityCheckSchemeRequest
//
// @return GetQualityCheckSchemeResponse
func (client *Client) GetQualityCheckScheme(request *GetQualityCheckSchemeRequest) (_result *GetQualityCheckSchemeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetQualityCheckSchemeResponse{}
	_body, _err := client.GetQualityCheckSchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResultResponse
func (client *Client) GetResultWithOptions(request *GetResultRequest, runtime *util.RuntimeOptions) (_result *GetResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResult"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetResultRequest
//
// @return GetResultResponse
func (client *Client) GetResult(request *GetResultRequest) (_result *GetResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResultResponse{}
	_body, _err := client.GetResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetResultToReviewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResultToReviewResponse
func (client *Client) GetResultToReviewWithOptions(request *GetResultToReviewRequest, runtime *util.RuntimeOptions) (_result *GetResultToReviewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResultToReview"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResultToReviewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetResultToReviewRequest
//
// @return GetResultToReviewResponse
func (client *Client) GetResultToReview(request *GetResultToReviewRequest) (_result *GetResultToReviewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResultToReviewResponse{}
	_body, _err := client.GetResultToReviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetRule is deprecated, please use Qualitycheck::2019-01-15::GetRuleV4 instead.
//
// @param request - GetRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRuleResponse
// Deprecated
func (client *Client) GetRuleWithOptions(request *GetRuleRequest, runtime *util.RuntimeOptions) (_result *GetRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRule"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetRule is deprecated, please use Qualitycheck::2019-01-15::GetRuleV4 instead.
//
// @param request - GetRuleRequest
//
// @return GetRuleResponse
// Deprecated
func (client *Client) GetRule(request *GetRuleRequest) (_result *GetRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRuleResponse{}
	_body, _err := client.GetRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetRuleById is deprecated, please use Qualitycheck::2019-01-15::GetRuleV4 instead.
//
// Summary:
//
// 获取规则
//
// @param request - GetRuleByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRuleByIdResponse
// Deprecated
func (client *Client) GetRuleByIdWithOptions(request *GetRuleByIdRequest, runtime *util.RuntimeOptions) (_result *GetRuleByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		body["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRuleById"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRuleByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetRuleById is deprecated, please use Qualitycheck::2019-01-15::GetRuleV4 instead.
//
// Summary:
//
// 获取规则
//
// @param request - GetRuleByIdRequest
//
// @return GetRuleByIdResponse
// Deprecated
func (client *Client) GetRuleById(request *GetRuleByIdRequest) (_result *GetRuleByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRuleByIdResponse{}
	_body, _err := client.GetRuleByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetRuleCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRuleCategoryResponse
func (client *Client) GetRuleCategoryWithOptions(request *GetRuleCategoryRequest, runtime *util.RuntimeOptions) (_result *GetRuleCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRuleCategory"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRuleCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetRuleCategoryRequest
//
// @return GetRuleCategoryResponse
func (client *Client) GetRuleCategory(request *GetRuleCategoryRequest) (_result *GetRuleCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRuleCategoryResponse{}
	_body, _err := client.GetRuleCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetRuleDetail is deprecated, please use Qualitycheck::2019-01-15::GetRuleV4 instead.
//
// @param request - GetRuleDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRuleDetailResponse
// Deprecated
func (client *Client) GetRuleDetailWithOptions(request *GetRuleDetailRequest, runtime *util.RuntimeOptions) (_result *GetRuleDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRuleDetail"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRuleDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetRuleDetail is deprecated, please use Qualitycheck::2019-01-15::GetRuleV4 instead.
//
// @param request - GetRuleDetailRequest
//
// @return GetRuleDetailResponse
// Deprecated
func (client *Client) GetRuleDetail(request *GetRuleDetailRequest) (_result *GetRuleDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRuleDetailResponse{}
	_body, _err := client.GetRuleDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// V4获取规则
//
// @param request - GetRuleV4Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRuleV4Response
func (client *Client) GetRuleV4WithOptions(request *GetRuleV4Request, runtime *util.RuntimeOptions) (_result *GetRuleV4Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		body["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRuleV4"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRuleV4Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// V4获取规则
//
// @param request - GetRuleV4Request
//
// @return GetRuleV4Response
func (client *Client) GetRuleV4(request *GetRuleV4Request) (_result *GetRuleV4Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRuleV4Response{}
	_body, _err := client.GetRuleV4WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获得规则列表
//
// @param request - GetRulesCountListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRulesCountListResponse
func (client *Client) GetRulesCountListWithOptions(request *GetRulesCountListRequest, runtime *util.RuntimeOptions) (_result *GetRulesCountListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessName)) {
		body["BusinessName"] = request.BusinessName
	}

	if !tea.BoolValue(util.IsUnset(request.BusinessRange)) {
		body["BusinessRange"] = request.BusinessRange
	}

	if !tea.BoolValue(util.IsUnset(request.CategoryName)) {
		body["CategoryName"] = request.CategoryName
	}

	if !tea.BoolValue(util.IsUnset(request.CountTotal)) {
		body["CountTotal"] = request.CountTotal
	}

	if !tea.BoolValue(util.IsUnset(request.CreateEmpid)) {
		body["CreateEmpid"] = request.CreateEmpid
	}

	if !tea.BoolValue(util.IsUnset(request.CreateUserId)) {
		body["CreateUserId"] = request.CreateUserId
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.LastUpdateEmpid)) {
		body["LastUpdateEmpid"] = request.LastUpdateEmpid
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RequireInfos)) {
		body["RequireInfos"] = request.RequireInfos
	}

	if !tea.BoolValue(util.IsUnset(request.Rid)) {
		body["Rid"] = request.Rid
	}

	if !tea.BoolValue(util.IsUnset(request.RuleIdOrRuleName)) {
		body["RuleIdOrRuleName"] = request.RuleIdOrRuleName
	}

	if !tea.BoolValue(util.IsUnset(request.RuleScoreSingleType)) {
		body["RuleScoreSingleType"] = request.RuleScoreSingleType
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		body["RuleType"] = request.RuleType
	}

	if !tea.BoolValue(util.IsUnset(request.SchemeId)) {
		body["SchemeId"] = request.SchemeId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		body["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.TypeName)) {
		body["TypeName"] = request.TypeName
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateEndTime)) {
		body["UpdateEndTime"] = request.UpdateEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateStartTime)) {
		body["UpdateStartTime"] = request.UpdateStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateUserId)) {
		body["UpdateUserId"] = request.UpdateUserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRulesCountList"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRulesCountListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获得规则列表
//
// @param request - GetRulesCountListRequest
//
// @return GetRulesCountListResponse
func (client *Client) GetRulesCountList(request *GetRulesCountListRequest) (_result *GetRulesCountListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRulesCountListResponse{}
	_body, _err := client.GetRulesCountListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetScoreInfo is deprecated
//
// @param request - GetScoreInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetScoreInfoResponse
// Deprecated
func (client *Client) GetScoreInfoWithOptions(request *GetScoreInfoRequest, runtime *util.RuntimeOptions) (_result *GetScoreInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetScoreInfo"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetScoreInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetScoreInfo is deprecated
//
// @param request - GetScoreInfoRequest
//
// @return GetScoreInfoResponse
// Deprecated
func (client *Client) GetScoreInfo(request *GetScoreInfoRequest) (_result *GetScoreInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetScoreInfoResponse{}
	_body, _err := client.GetScoreInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetSkillGroupConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSkillGroupConfigResponse
func (client *Client) GetSkillGroupConfigWithOptions(request *GetSkillGroupConfigRequest, runtime *util.RuntimeOptions) (_result *GetSkillGroupConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSkillGroupConfig"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSkillGroupConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetSkillGroupConfigRequest
//
// @return GetSkillGroupConfigResponse
func (client *Client) GetSkillGroupConfig(request *GetSkillGroupConfigRequest) (_result *GetSkillGroupConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSkillGroupConfigResponse{}
	_body, _err := client.GetSkillGroupConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetSyncResult is deprecated, please use Qualitycheck::2019-01-15::GetResult instead.
//
// @param request - GetSyncResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSyncResultResponse
// Deprecated
func (client *Client) GetSyncResultWithOptions(request *GetSyncResultRequest, runtime *util.RuntimeOptions) (_result *GetSyncResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSyncResult"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSyncResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetSyncResult is deprecated, please use Qualitycheck::2019-01-15::GetResult instead.
//
// @param request - GetSyncResultRequest
//
// @return GetSyncResultResponse
// Deprecated
func (client *Client) GetSyncResult(request *GetSyncResultRequest) (_result *GetSyncResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSyncResultResponse{}
	_body, _err := client.GetSyncResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 预警策略-详情
//
// @param request - GetWarningStrategyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWarningStrategyConfigResponse
func (client *Client) GetWarningStrategyConfigWithOptions(request *GetWarningStrategyConfigRequest, runtime *util.RuntimeOptions) (_result *GetWarningStrategyConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWarningStrategyConfig"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWarningStrategyConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预警策略-详情
//
// @param request - GetWarningStrategyConfigRequest
//
// @return GetWarningStrategyConfigResponse
func (client *Client) GetWarningStrategyConfig(request *GetWarningStrategyConfigRequest) (_result *GetWarningStrategyConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWarningStrategyConfigResponse{}
	_body, _err := client.GetWarningStrategyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - HandleComplaintRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HandleComplaintResponse
func (client *Client) HandleComplaintWithOptions(request *HandleComplaintRequest, runtime *util.RuntimeOptions) (_result *HandleComplaintResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("HandleComplaint"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &HandleComplaintResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - HandleComplaintRequest
//
// @return HandleComplaintResponse
func (client *Client) HandleComplaint(request *HandleComplaintRequest) (_result *HandleComplaintResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &HandleComplaintResponse{}
	_body, _err := client.HandleComplaintWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI InsertScoreForApi is deprecated
//
// @param request - InsertScoreForApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsertScoreForApiResponse
// Deprecated
func (client *Client) InsertScoreForApiWithOptions(request *InsertScoreForApiRequest, runtime *util.RuntimeOptions) (_result *InsertScoreForApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InsertScoreForApi"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InsertScoreForApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI InsertScoreForApi is deprecated
//
// @param request - InsertScoreForApiRequest
//
// @return InsertScoreForApiResponse
// Deprecated
func (client *Client) InsertScoreForApi(request *InsertScoreForApiRequest) (_result *InsertScoreForApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InsertScoreForApiResponse{}
	_body, _err := client.InsertScoreForApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI InsertSubScoreForApi is deprecated
//
// @param request - InsertSubScoreForApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsertSubScoreForApiResponse
// Deprecated
func (client *Client) InsertSubScoreForApiWithOptions(request *InsertSubScoreForApiRequest, runtime *util.RuntimeOptions) (_result *InsertSubScoreForApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InsertSubScoreForApi"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InsertSubScoreForApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI InsertSubScoreForApi is deprecated
//
// @param request - InsertSubScoreForApiRequest
//
// @return InsertSubScoreForApiResponse
// Deprecated
func (client *Client) InsertSubScoreForApi(request *InsertSubScoreForApiRequest) (_result *InsertSubScoreForApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InsertSubScoreForApiResponse{}
	_body, _err := client.InsertSubScoreForApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI InvalidRule is deprecated, please use Qualitycheck::2019-01-15::DeleteRuleV4 instead.
//
// @param request - InvalidRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvalidRuleResponse
// Deprecated
func (client *Client) InvalidRuleWithOptions(request *InvalidRuleRequest, runtime *util.RuntimeOptions) (_result *InvalidRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InvalidRule"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InvalidRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI InvalidRule is deprecated, please use Qualitycheck::2019-01-15::DeleteRuleV4 instead.
//
// @param request - InvalidRuleRequest
//
// @return InvalidRuleResponse
// Deprecated
func (client *Client) InvalidRule(request *InvalidRuleRequest) (_result *InvalidRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InvalidRuleResponse{}
	_body, _err := client.InvalidRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListAsrVocabRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAsrVocabResponse
func (client *Client) ListAsrVocabWithOptions(request *ListAsrVocabRequest, runtime *util.RuntimeOptions) (_result *ListAsrVocabResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAsrVocab"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAsrVocabResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListAsrVocabRequest
//
// @return ListAsrVocabResponse
func (client *Client) ListAsrVocab(request *ListAsrVocabRequest) (_result *ListAsrVocabResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAsrVocabResponse{}
	_body, _err := client.ListAsrVocabWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据集列表
//
// @param request - ListDataSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSetResponse
func (client *Client) ListDataSetWithOptions(request *ListDataSetRequest, runtime *util.RuntimeOptions) (_result *ListDataSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDataSet"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDataSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据集列表
//
// @param request - ListDataSetRequest
//
// @return ListDataSetResponse
func (client *Client) ListDataSet(request *ListDataSetRequest) (_result *ListDataSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDataSetResponse{}
	_body, _err := client.ListDataSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListPrecisionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrecisionTaskResponse
func (client *Client) ListPrecisionTaskWithOptions(request *ListPrecisionTaskRequest, runtime *util.RuntimeOptions) (_result *ListPrecisionTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPrecisionTask"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPrecisionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListPrecisionTaskRequest
//
// @return ListPrecisionTaskResponse
func (client *Client) ListPrecisionTask(request *ListPrecisionTaskRequest) (_result *ListPrecisionTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPrecisionTaskResponse{}
	_body, _err := client.ListPrecisionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 质检方案列表
//
// @param request - ListQualityCheckSchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQualityCheckSchemeResponse
func (client *Client) ListQualityCheckSchemeWithOptions(request *ListQualityCheckSchemeRequest, runtime *util.RuntimeOptions) (_result *ListQualityCheckSchemeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQualityCheckScheme"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListQualityCheckSchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 质检方案列表
//
// @param request - ListQualityCheckSchemeRequest
//
// @return ListQualityCheckSchemeResponse
func (client *Client) ListQualityCheckScheme(request *ListQualityCheckSchemeRequest) (_result *ListQualityCheckSchemeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListQualityCheckSchemeResponse{}
	_body, _err := client.ListQualityCheckSchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ListRules is deprecated, please use Qualitycheck::2019-01-15::ListRulesV4 instead.
//
// @param request - ListRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRulesResponse
// Deprecated
func (client *Client) ListRulesWithOptions(request *ListRulesRequest, runtime *util.RuntimeOptions) (_result *ListRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRules"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListRules is deprecated, please use Qualitycheck::2019-01-15::ListRulesV4 instead.
//
// @param request - ListRulesRequest
//
// @return ListRulesResponse
// Deprecated
func (client *Client) ListRules(request *ListRulesRequest) (_result *ListRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRulesResponse{}
	_body, _err := client.ListRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// V4获得规则列表
//
// @param request - ListRulesV4Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRulesV4Response
func (client *Client) ListRulesV4WithOptions(request *ListRulesV4Request, runtime *util.RuntimeOptions) (_result *ListRulesV4Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessName)) {
		body["BusinessName"] = request.BusinessName
	}

	if !tea.BoolValue(util.IsUnset(request.BusinessRange)) {
		body["BusinessRange"] = request.BusinessRange
	}

	if !tea.BoolValue(util.IsUnset(request.CategoryName)) {
		body["CategoryName"] = request.CategoryName
	}

	if !tea.BoolValue(util.IsUnset(request.CountTotal)) {
		body["CountTotal"] = request.CountTotal
	}

	if !tea.BoolValue(util.IsUnset(request.CreateEmpid)) {
		body["CreateEmpid"] = request.CreateEmpid
	}

	if !tea.BoolValue(util.IsUnset(request.CreateUserId)) {
		body["CreateUserId"] = request.CreateUserId
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.LastUpdateEmpid)) {
		body["LastUpdateEmpid"] = request.LastUpdateEmpid
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RequireInfos)) {
		body["RequireInfos"] = request.RequireInfos
	}

	if !tea.BoolValue(util.IsUnset(request.Rid)) {
		body["Rid"] = request.Rid
	}

	if !tea.BoolValue(util.IsUnset(request.RuleIdOrRuleName)) {
		body["RuleIdOrRuleName"] = request.RuleIdOrRuleName
	}

	if !tea.BoolValue(util.IsUnset(request.RuleScoreSingleType)) {
		body["RuleScoreSingleType"] = request.RuleScoreSingleType
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		body["RuleType"] = request.RuleType
	}

	if !tea.BoolValue(util.IsUnset(request.SchemeId)) {
		body["SchemeId"] = request.SchemeId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		body["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.TypeName)) {
		body["TypeName"] = request.TypeName
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateEndTime)) {
		body["UpdateEndTime"] = request.UpdateEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateStartTime)) {
		body["UpdateStartTime"] = request.UpdateStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateUserId)) {
		body["UpdateUserId"] = request.UpdateUserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRulesV4"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRulesV4Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// V4获得规则列表
//
// @param request - ListRulesV4Request
//
// @return ListRulesV4Response
func (client *Client) ListRulesV4(request *ListRulesV4Request) (_result *ListRulesV4Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRulesV4Response{}
	_body, _err := client.ListRulesV4WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取质检任务列表
//
// @param request - ListSchemeTaskConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSchemeTaskConfigResponse
func (client *Client) ListSchemeTaskConfigWithOptions(request *ListSchemeTaskConfigRequest, runtime *util.RuntimeOptions) (_result *ListSchemeTaskConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSchemeTaskConfig"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSchemeTaskConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取质检任务列表
//
// @param request - ListSchemeTaskConfigRequest
//
// @return ListSchemeTaskConfigResponse
func (client *Client) ListSchemeTaskConfig(request *ListSchemeTaskConfigRequest) (_result *ListSchemeTaskConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSchemeTaskConfigResponse{}
	_body, _err := client.ListSchemeTaskConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取会话组列表
//
// @param request - ListSessionGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSessionGroupResponse
func (client *Client) ListSessionGroupWithOptions(request *ListSessionGroupRequest, runtime *util.RuntimeOptions) (_result *ListSessionGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSessionGroup"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSessionGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取会话组列表
//
// @param request - ListSessionGroupRequest
//
// @return ListSessionGroupResponse
func (client *Client) ListSessionGroup(request *ListSessionGroupRequest) (_result *ListSessionGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSessionGroupResponse{}
	_body, _err := client.ListSessionGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListSkillGroupConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSkillGroupConfigResponse
func (client *Client) ListSkillGroupConfigWithOptions(request *ListSkillGroupConfigRequest, runtime *util.RuntimeOptions) (_result *ListSkillGroupConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSkillGroupConfig"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSkillGroupConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListSkillGroupConfigRequest
//
// @return ListSkillGroupConfigResponse
func (client *Client) ListSkillGroupConfig(request *ListSkillGroupConfigRequest) (_result *ListSkillGroupConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSkillGroupConfigResponse{}
	_body, _err := client.ListSkillGroupConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListTaskAssignRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTaskAssignRulesResponse
func (client *Client) ListTaskAssignRulesWithOptions(request *ListTaskAssignRulesRequest, runtime *util.RuntimeOptions) (_result *ListTaskAssignRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTaskAssignRules"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTaskAssignRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListTaskAssignRulesRequest
//
// @return ListTaskAssignRulesResponse
func (client *Client) ListTaskAssignRules(request *ListTaskAssignRulesRequest) (_result *ListTaskAssignRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTaskAssignRulesResponse{}
	_body, _err := client.ListTaskAssignRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersResponse
func (client *Client) ListUsersWithOptions(request *ListUsersRequest, runtime *util.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUsers"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListUsersRequest
//
// @return ListUsersResponse
func (client *Client) ListUsers(request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListWarningConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWarningConfigResponse
func (client *Client) ListWarningConfigWithOptions(request *ListWarningConfigRequest, runtime *util.RuntimeOptions) (_result *ListWarningConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWarningConfig"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWarningConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListWarningConfigRequest
//
// @return ListWarningConfigResponse
func (client *Client) ListWarningConfig(request *ListWarningConfigRequest) (_result *ListWarningConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWarningConfigResponse{}
	_body, _err := client.ListWarningConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 预警策略-列表
//
// @param request - ListWarningStrategyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWarningStrategyConfigResponse
func (client *Client) ListWarningStrategyConfigWithOptions(request *ListWarningStrategyConfigRequest, runtime *util.RuntimeOptions) (_result *ListWarningStrategyConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWarningStrategyConfig"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWarningStrategyConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预警策略-列表
//
// @param request - ListWarningStrategyConfigRequest
//
// @return ListWarningStrategyConfigResponse
func (client *Client) ListWarningStrategyConfig(request *ListWarningStrategyConfigRequest) (_result *ListWarningStrategyConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWarningStrategyConfigResponse{}
	_body, _err := client.ListWarningStrategyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量回收
//
// @param request - RevertAssignedSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevertAssignedSessionResponse
func (client *Client) RevertAssignedSessionWithOptions(request *RevertAssignedSessionRequest, runtime *util.RuntimeOptions) (_result *RevertAssignedSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RevertAssignedSession"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RevertAssignedSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量回收
//
// @param request - RevertAssignedSessionRequest
//
// @return RevertAssignedSessionResponse
func (client *Client) RevertAssignedSession(request *RevertAssignedSessionRequest) (_result *RevertAssignedSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RevertAssignedSessionResponse{}
	_body, _err := client.RevertAssignedSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 会话组批量回收
//
// @param request - RevertAssignedSessionGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevertAssignedSessionGroupResponse
func (client *Client) RevertAssignedSessionGroupWithOptions(request *RevertAssignedSessionGroupRequest, runtime *util.RuntimeOptions) (_result *RevertAssignedSessionGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RevertAssignedSessionGroup"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RevertAssignedSessionGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 会话组批量回收
//
// @param request - RevertAssignedSessionGroupRequest
//
// @return RevertAssignedSessionGroupResponse
func (client *Client) RevertAssignedSessionGroup(request *RevertAssignedSessionGroupRequest) (_result *RevertAssignedSessionGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RevertAssignedSessionGroupResponse{}
	_body, _err := client.RevertAssignedSessionGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SaveConfigDataSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveConfigDataSetResponse
func (client *Client) SaveConfigDataSetWithOptions(request *SaveConfigDataSetRequest, runtime *util.RuntimeOptions) (_result *SaveConfigDataSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveConfigDataSet"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveConfigDataSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - SaveConfigDataSetRequest
//
// @return SaveConfigDataSetResponse
func (client *Client) SaveConfigDataSet(request *SaveConfigDataSetRequest) (_result *SaveConfigDataSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveConfigDataSetResponse{}
	_body, _err := client.SaveConfigDataSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SubmitComplaintRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitComplaintResponse
func (client *Client) SubmitComplaintWithOptions(request *SubmitComplaintRequest, runtime *util.RuntimeOptions) (_result *SubmitComplaintResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitComplaint"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitComplaintResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - SubmitComplaintRequest
//
// @return SubmitComplaintResponse
func (client *Client) SubmitComplaint(request *SubmitComplaintRequest) (_result *SubmitComplaintResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitComplaintResponse{}
	_body, _err := client.SubmitComplaintWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SubmitPrecisionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitPrecisionTaskResponse
func (client *Client) SubmitPrecisionTaskWithOptions(request *SubmitPrecisionTaskRequest, runtime *util.RuntimeOptions) (_result *SubmitPrecisionTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitPrecisionTask"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitPrecisionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - SubmitPrecisionTaskRequest
//
// @return SubmitPrecisionTaskResponse
func (client *Client) SubmitPrecisionTask(request *SubmitPrecisionTaskRequest) (_result *SubmitPrecisionTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitPrecisionTaskResponse{}
	_body, _err := client.SubmitPrecisionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SubmitQualityCheckTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitQualityCheckTaskResponse
func (client *Client) SubmitQualityCheckTaskWithOptions(request *SubmitQualityCheckTaskRequest, runtime *util.RuntimeOptions) (_result *SubmitQualityCheckTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitQualityCheckTask"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitQualityCheckTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - SubmitQualityCheckTaskRequest
//
// @return SubmitQualityCheckTaskResponse
func (client *Client) SubmitQualityCheckTask(request *SubmitQualityCheckTaskRequest) (_result *SubmitQualityCheckTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitQualityCheckTaskResponse{}
	_body, _err := client.SubmitQualityCheckTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SubmitReviewInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitReviewInfoResponse
func (client *Client) SubmitReviewInfoWithOptions(request *SubmitReviewInfoRequest, runtime *util.RuntimeOptions) (_result *SubmitReviewInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitReviewInfo"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitReviewInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - SubmitReviewInfoRequest
//
// @return SubmitReviewInfoResponse
func (client *Client) SubmitReviewInfo(request *SubmitReviewInfoRequest) (_result *SubmitReviewInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitReviewInfoResponse{}
	_body, _err := client.SubmitReviewInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SyncQualityCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncQualityCheckResponse
func (client *Client) SyncQualityCheckWithOptions(request *SyncQualityCheckRequest, runtime *util.RuntimeOptions) (_result *SyncQualityCheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SyncQualityCheck"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncQualityCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - SyncQualityCheckRequest
//
// @return SyncQualityCheckResponse
func (client *Client) SyncQualityCheck(request *SyncQualityCheckRequest) (_result *SyncQualityCheckResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SyncQualityCheckResponse{}
	_body, _err := client.SyncQualityCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 测试规则
//
// @param request - TestRuleV4Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TestRuleV4Response
func (client *Client) TestRuleV4WithOptions(request *TestRuleV4Request, runtime *util.RuntimeOptions) (_result *TestRuleV4Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsSchemeData)) {
		body["IsSchemeData"] = request.IsSchemeData
	}

	if !tea.BoolValue(util.IsUnset(request.TestJson)) {
		body["TestJson"] = request.TestJson
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TestRuleV4"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TestRuleV4Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 测试规则
//
// @param request - TestRuleV4Request
//
// @return TestRuleV4Response
func (client *Client) TestRuleV4(request *TestRuleV4Request) (_result *TestRuleV4Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TestRuleV4Response{}
	_body, _err := client.TestRuleV4WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateAsrVocabRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAsrVocabResponse
func (client *Client) UpdateAsrVocabWithOptions(request *UpdateAsrVocabRequest, runtime *util.RuntimeOptions) (_result *UpdateAsrVocabResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAsrVocab"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAsrVocabResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateAsrVocabRequest
//
// @return UpdateAsrVocabResponse
func (client *Client) UpdateAsrVocab(request *UpdateAsrVocabRequest) (_result *UpdateAsrVocabResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAsrVocabResponse{}
	_body, _err := client.UpdateAsrVocabWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新质检方案中的质检维度
//
// @param request - UpdateCheckTypeToSchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCheckTypeToSchemeResponse
func (client *Client) UpdateCheckTypeToSchemeWithOptions(request *UpdateCheckTypeToSchemeRequest, runtime *util.RuntimeOptions) (_result *UpdateCheckTypeToSchemeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCheckTypeToScheme"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCheckTypeToSchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新质检方案中的质检维度
//
// @param request - UpdateCheckTypeToSchemeRequest
//
// @return UpdateCheckTypeToSchemeResponse
func (client *Client) UpdateCheckTypeToScheme(request *UpdateCheckTypeToSchemeRequest) (_result *UpdateCheckTypeToSchemeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCheckTypeToSchemeResponse{}
	_body, _err := client.UpdateCheckTypeToSchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新会话随录数据
//
// @param request - UpdateQualityCheckDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateQualityCheckDataResponse
func (client *Client) UpdateQualityCheckDataWithOptions(request *UpdateQualityCheckDataRequest, runtime *util.RuntimeOptions) (_result *UpdateQualityCheckDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateQualityCheckData"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateQualityCheckDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新会话随录数据
//
// @param request - UpdateQualityCheckDataRequest
//
// @return UpdateQualityCheckDataResponse
func (client *Client) UpdateQualityCheckData(request *UpdateQualityCheckDataRequest) (_result *UpdateQualityCheckDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateQualityCheckDataResponse{}
	_body, _err := client.UpdateQualityCheckDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新质检方案
//
// @param request - UpdateQualityCheckSchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateQualityCheckSchemeResponse
func (client *Client) UpdateQualityCheckSchemeWithOptions(request *UpdateQualityCheckSchemeRequest, runtime *util.RuntimeOptions) (_result *UpdateQualityCheckSchemeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateQualityCheckScheme"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateQualityCheckSchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新质检方案
//
// @param request - UpdateQualityCheckSchemeRequest
//
// @return UpdateQualityCheckSchemeResponse
func (client *Client) UpdateQualityCheckScheme(request *UpdateQualityCheckSchemeRequest) (_result *UpdateQualityCheckSchemeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateQualityCheckSchemeResponse{}
	_body, _err := client.UpdateQualityCheckSchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI UpdateRule is deprecated, please use Qualitycheck::2019-01-15::UpdateRuleV4 instead.
//
// @param request - UpdateRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRuleResponse
// Deprecated
func (client *Client) UpdateRuleWithOptions(request *UpdateRuleRequest, runtime *util.RuntimeOptions) (_result *UpdateRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRule"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI UpdateRule is deprecated, please use Qualitycheck::2019-01-15::UpdateRuleV4 instead.
//
// @param request - UpdateRuleRequest
//
// @return UpdateRuleResponse
// Deprecated
func (client *Client) UpdateRule(request *UpdateRuleRequest) (_result *UpdateRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRuleResponse{}
	_body, _err := client.UpdateRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI UpdateRuleById is deprecated, please use Qualitycheck::2019-01-15::UpdateRuleV4 instead.
//
// Summary:
//
// 更新规则
//
// @param request - UpdateRuleByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRuleByIdResponse
// Deprecated
func (client *Client) UpdateRuleByIdWithOptions(request *UpdateRuleByIdRequest, runtime *util.RuntimeOptions) (_result *UpdateRuleByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsCopy)) {
		body["IsCopy"] = request.IsCopy
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStrForRule)) {
		body["JsonStrForRule"] = request.JsonStrForRule
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnRelatedSchemes)) {
		body["ReturnRelatedSchemes"] = request.ReturnRelatedSchemes
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		body["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRuleById"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRuleByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI UpdateRuleById is deprecated, please use Qualitycheck::2019-01-15::UpdateRuleV4 instead.
//
// Summary:
//
// 更新规则
//
// @param request - UpdateRuleByIdRequest
//
// @return UpdateRuleByIdResponse
// Deprecated
func (client *Client) UpdateRuleById(request *UpdateRuleByIdRequest) (_result *UpdateRuleByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRuleByIdResponse{}
	_body, _err := client.UpdateRuleByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新质检方案的规则
//
// @param request - UpdateRuleToSchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRuleToSchemeResponse
func (client *Client) UpdateRuleToSchemeWithOptions(request *UpdateRuleToSchemeRequest, runtime *util.RuntimeOptions) (_result *UpdateRuleToSchemeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRuleToScheme"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRuleToSchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新质检方案的规则
//
// @param request - UpdateRuleToSchemeRequest
//
// @return UpdateRuleToSchemeResponse
func (client *Client) UpdateRuleToScheme(request *UpdateRuleToSchemeRequest) (_result *UpdateRuleToSchemeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRuleToSchemeResponse{}
	_body, _err := client.UpdateRuleToSchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// V4更新规则
//
// @param request - UpdateRuleV4Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRuleV4Response
func (client *Client) UpdateRuleV4WithOptions(request *UpdateRuleV4Request, runtime *util.RuntimeOptions) (_result *UpdateRuleV4Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStrForRule)) {
		body["JsonStrForRule"] = request.JsonStrForRule
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		body["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRuleV4"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRuleV4Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// V4更新规则
//
// @param request - UpdateRuleV4Request
//
// @return UpdateRuleV4Response
func (client *Client) UpdateRuleV4(request *UpdateRuleV4Request) (_result *UpdateRuleV4Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRuleV4Response{}
	_body, _err := client.UpdateRuleV4WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建质检任务
//
// @param request - UpdateSchemeTaskConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSchemeTaskConfigResponse
func (client *Client) UpdateSchemeTaskConfigWithOptions(request *UpdateSchemeTaskConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateSchemeTaskConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSchemeTaskConfig"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSchemeTaskConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建质检任务
//
// @param request - UpdateSchemeTaskConfigRequest
//
// @return UpdateSchemeTaskConfigResponse
func (client *Client) UpdateSchemeTaskConfig(request *UpdateSchemeTaskConfigRequest) (_result *UpdateSchemeTaskConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSchemeTaskConfigResponse{}
	_body, _err := client.UpdateSchemeTaskConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI UpdateScoreForApi is deprecated
//
// @param request - UpdateScoreForApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateScoreForApiResponse
// Deprecated
func (client *Client) UpdateScoreForApiWithOptions(request *UpdateScoreForApiRequest, runtime *util.RuntimeOptions) (_result *UpdateScoreForApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateScoreForApi"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateScoreForApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI UpdateScoreForApi is deprecated
//
// @param request - UpdateScoreForApiRequest
//
// @return UpdateScoreForApiResponse
// Deprecated
func (client *Client) UpdateScoreForApi(request *UpdateScoreForApiRequest) (_result *UpdateScoreForApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateScoreForApiResponse{}
	_body, _err := client.UpdateScoreForApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateSkillGroupConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSkillGroupConfigResponse
func (client *Client) UpdateSkillGroupConfigWithOptions(request *UpdateSkillGroupConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateSkillGroupConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSkillGroupConfig"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSkillGroupConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateSkillGroupConfigRequest
//
// @return UpdateSkillGroupConfigResponse
func (client *Client) UpdateSkillGroupConfig(request *UpdateSkillGroupConfigRequest) (_result *UpdateSkillGroupConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSkillGroupConfigResponse{}
	_body, _err := client.UpdateSkillGroupConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI UpdateSubScoreForApi is deprecated
//
// @param request - UpdateSubScoreForApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSubScoreForApiResponse
// Deprecated
func (client *Client) UpdateSubScoreForApiWithOptions(request *UpdateSubScoreForApiRequest, runtime *util.RuntimeOptions) (_result *UpdateSubScoreForApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSubScoreForApi"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSubScoreForApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI UpdateSubScoreForApi is deprecated
//
// @param request - UpdateSubScoreForApiRequest
//
// @return UpdateSubScoreForApiResponse
// Deprecated
func (client *Client) UpdateSubScoreForApi(request *UpdateSubScoreForApiRequest) (_result *UpdateSubScoreForApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSubScoreForApiResponse{}
	_body, _err := client.UpdateSubScoreForApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateSyncQualityCheckDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSyncQualityCheckDataResponse
func (client *Client) UpdateSyncQualityCheckDataWithOptions(request *UpdateSyncQualityCheckDataRequest, runtime *util.RuntimeOptions) (_result *UpdateSyncQualityCheckDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSyncQualityCheckData"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSyncQualityCheckDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateSyncQualityCheckDataRequest
//
// @return UpdateSyncQualityCheckDataResponse
func (client *Client) UpdateSyncQualityCheckData(request *UpdateSyncQualityCheckDataRequest) (_result *UpdateSyncQualityCheckDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSyncQualityCheckDataResponse{}
	_body, _err := client.UpdateSyncQualityCheckDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateTaskAssignRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTaskAssignRuleResponse
func (client *Client) UpdateTaskAssignRuleWithOptions(request *UpdateTaskAssignRuleRequest, runtime *util.RuntimeOptions) (_result *UpdateTaskAssignRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTaskAssignRule"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTaskAssignRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateTaskAssignRuleRequest
//
// @return UpdateTaskAssignRuleResponse
func (client *Client) UpdateTaskAssignRule(request *UpdateTaskAssignRuleRequest) (_result *UpdateTaskAssignRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTaskAssignRuleResponse{}
	_body, _err := client.UpdateTaskAssignRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserResponse
func (client *Client) UpdateUserWithOptions(request *UpdateUserRequest, runtime *util.RuntimeOptions) (_result *UpdateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUser"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateUserRequest
//
// @return UpdateUserResponse
func (client *Client) UpdateUser(request *UpdateUserRequest) (_result *UpdateUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserResponse{}
	_body, _err := client.UpdateUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateWarningConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWarningConfigResponse
func (client *Client) UpdateWarningConfigWithOptions(request *UpdateWarningConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateWarningConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWarningConfig"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWarningConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateWarningConfigRequest
//
// @return UpdateWarningConfigResponse
func (client *Client) UpdateWarningConfig(request *UpdateWarningConfigRequest) (_result *UpdateWarningConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateWarningConfigResponse{}
	_body, _err := client.UpdateWarningConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 预警策略-更新
//
// @param request - UpdateWarningStrategyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWarningStrategyConfigResponse
func (client *Client) UpdateWarningStrategyConfigWithOptions(request *UpdateWarningStrategyConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateWarningStrategyConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWarningStrategyConfig"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWarningStrategyConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预警策略-更新
//
// @param request - UpdateWarningStrategyConfigRequest
//
// @return UpdateWarningStrategyConfigResponse
func (client *Client) UpdateWarningStrategyConfig(request *UpdateWarningStrategyConfigRequest) (_result *UpdateWarningStrategyConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateWarningStrategyConfigResponse{}
	_body, _err := client.UpdateWarningStrategyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UploadAudioDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadAudioDataResponse
func (client *Client) UploadAudioDataWithOptions(request *UploadAudioDataRequest, runtime *util.RuntimeOptions) (_result *UploadAudioDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadAudioData"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadAudioDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UploadAudioDataRequest
//
// @return UploadAudioDataResponse
func (client *Client) UploadAudioData(request *UploadAudioDataRequest) (_result *UploadAudioDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadAudioDataResponse{}
	_body, _err := client.UploadAudioDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI UploadData is deprecated, please use Qualitycheck::2019-01-15::UploadDataV4 instead.
//
// Summary:
//
// 推荐使用UploadDataV4接口,支持更长的JsonStr,但仅支持POST方法.
//
// @param request - UploadDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadDataResponse
// Deprecated
func (client *Client) UploadDataWithOptions(request *UploadDataRequest, runtime *util.RuntimeOptions) (_result *UploadDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadData"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI UploadData is deprecated, please use Qualitycheck::2019-01-15::UploadDataV4 instead.
//
// Summary:
//
// 推荐使用UploadDataV4接口,支持更长的JsonStr,但仅支持POST方法.
//
// @param request - UploadDataRequest
//
// @return UploadDataResponse
// Deprecated
func (client *Client) UploadData(request *UploadDataRequest) (_result *UploadDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadDataResponse{}
	_body, _err := client.UploadDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// http_hsf
//
// @param request - UploadDataSyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadDataSyncResponse
func (client *Client) UploadDataSyncWithOptions(request *UploadDataSyncRequest, runtime *util.RuntimeOptions) (_result *UploadDataSyncResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadDataSync"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadDataSyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// http_hsf
//
// @param request - UploadDataSyncRequest
//
// @return UploadDataSyncResponse
func (client *Client) UploadDataSync(request *UploadDataSyncRequest) (_result *UploadDataSyncResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadDataSyncResponse{}
	_body, _err := client.UploadDataSyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// UploadDataV4
//
// @param request - UploadDataV4Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadDataV4Response
func (client *Client) UploadDataV4WithOptions(request *UploadDataV4Request, runtime *util.RuntimeOptions) (_result *UploadDataV4Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		body["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		body["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadDataV4"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadDataV4Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// UploadDataV4
//
// @param request - UploadDataV4Request
//
// @return UploadDataV4Response
func (client *Client) UploadDataV4(request *UploadDataV4Request) (_result *UploadDataV4Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadDataV4Response{}
	_body, _err := client.UploadDataV4WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UploadRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadRuleResponse
func (client *Client) UploadRuleWithOptions(request *UploadRuleRequest, runtime *util.RuntimeOptions) (_result *UploadRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadRule"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UploadRuleRequest
//
// @return UploadRuleResponse
func (client *Client) UploadRule(request *UploadRuleRequest) (_result *UploadRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadRuleResponse{}
	_body, _err := client.UploadRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - VerifyFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyFileResponse
func (client *Client) VerifyFileWithOptions(request *VerifyFileRequest, runtime *util.RuntimeOptions) (_result *VerifyFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyFile"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - VerifyFileRequest
//
// @return VerifyFileResponse
func (client *Client) VerifyFile(request *VerifyFileRequest) (_result *VerifyFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyFileResponse{}
	_body, _err := client.VerifyFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - VerifySentenceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifySentenceResponse
func (client *Client) VerifySentenceWithOptions(request *VerifySentenceRequest, runtime *util.RuntimeOptions) (_result *VerifySentenceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseMeAgentId)) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifySentence"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifySentenceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - VerifySentenceRequest
//
// @return VerifySentenceResponse
func (client *Client) VerifySentence(request *VerifySentenceRequest) (_result *VerifySentenceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifySentenceResponse{}
	_body, _err := client.VerifySentenceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
