// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperatorBasicInfo interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *OperatorBasicInfo
	GetId() *int64
	SetName(v string) *OperatorBasicInfo
	GetName() *string
	SetOid(v string) *OperatorBasicInfo
	GetOid() *string
	SetParam(v *OperatorBasicInfoParam) *OperatorBasicInfo
	GetParam() *OperatorBasicInfoParam
	SetQualityCheckType(v int32) *OperatorBasicInfo
	GetQualityCheckType() *int32
	SetType(v string) *OperatorBasicInfo
	GetType() *string
	SetUserGroup(v string) *OperatorBasicInfo
	GetUserGroup() *string
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
	return dara.Prettify(s)
}

func (s OperatorBasicInfo) GoString() string {
	return s.String()
}

func (s *OperatorBasicInfo) GetId() *int64 {
	return s.Id
}

func (s *OperatorBasicInfo) GetName() *string {
	return s.Name
}

func (s *OperatorBasicInfo) GetOid() *string {
	return s.Oid
}

func (s *OperatorBasicInfo) GetParam() *OperatorBasicInfoParam {
	return s.Param
}

func (s *OperatorBasicInfo) GetQualityCheckType() *int32 {
	return s.QualityCheckType
}

func (s *OperatorBasicInfo) GetType() *string {
	return s.Type
}

func (s *OperatorBasicInfo) GetUserGroup() *string {
	return s.UserGroup
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

func (s *OperatorBasicInfo) Validate() error {
	if s.Param != nil {
		if err := s.Param.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	return dara.Prettify(s)
}

func (s OperatorBasicInfoParam) GoString() string {
	return s.String()
}

func (s *OperatorBasicInfoParam) GetAnswerThreshold() *string {
	return s.AnswerThreshold
}

func (s *OperatorBasicInfoParam) GetAntModelInfo() map[string]*string {
	return s.AntModelInfo
}

func (s *OperatorBasicInfoParam) GetAverage() *bool {
	return s.Average
}

func (s *OperatorBasicInfoParam) GetBeginType() *string {
	return s.BeginType
}

func (s *OperatorBasicInfoParam) GetBotId() *string {
	return s.BotId
}

func (s *OperatorBasicInfoParam) GetCaseSensitive() *bool {
	return s.CaseSensitive
}

func (s *OperatorBasicInfoParam) GetCategoryPathCode() *string {
	return s.CategoryPathCode
}

func (s *OperatorBasicInfoParam) GetCheckFirstSentence() *bool {
	return s.CheckFirstSentence
}

func (s *OperatorBasicInfoParam) GetCheckType() *int32 {
	return s.CheckType
}

func (s *OperatorBasicInfoParam) GetCompareOperator() *string {
	return s.CompareOperator
}

func (s *OperatorBasicInfoParam) GetContextChatMatch() *bool {
	return s.ContextChatMatch
}

func (s *OperatorBasicInfoParam) GetCustomerParam() *JudgeNodeMetaDesc {
	return s.CustomerParam
}

func (s *OperatorBasicInfoParam) GetDelayTime() *int32 {
	return s.DelayTime
}

func (s *OperatorBasicInfoParam) GetDifferentRole() *bool {
	return s.DifferentRole
}

func (s *OperatorBasicInfoParam) GetEndType() *string {
	return s.EndType
}

func (s *OperatorBasicInfoParam) GetExcludes() []*string {
	return s.Excludes
}

func (s *OperatorBasicInfoParam) GetFrom() *int32 {
	return s.From
}

func (s *OperatorBasicInfoParam) GetFromEnd() *bool {
	return s.FromEnd
}

func (s *OperatorBasicInfoParam) GetHitTime() *int32 {
	return s.HitTime
}

func (s *OperatorBasicInfoParam) GetInSentence() *bool {
	return s.InSentence
}

func (s *OperatorBasicInfoParam) GetInterval() *int32 {
	return s.Interval
}

func (s *OperatorBasicInfoParam) GetIntervalEnd() *int32 {
	return s.IntervalEnd
}

func (s *OperatorBasicInfoParam) GetKeywordExtension() *int32 {
	return s.KeywordExtension
}

func (s *OperatorBasicInfoParam) GetKeywordMatchSize() *int32 {
	return s.KeywordMatchSize
}

func (s *OperatorBasicInfoParam) GetKeywords() []*string {
	return s.Keywords
}

func (s *OperatorBasicInfoParam) GetKnowledgeInfo() *string {
	return s.KnowledgeInfo
}

func (s *OperatorBasicInfoParam) GetKnowledgeSentenceNum() *int32 {
	return s.KnowledgeSentenceNum
}

func (s *OperatorBasicInfoParam) GetKnowledgeTargetId() *string {
	return s.KnowledgeTargetId
}

func (s *OperatorBasicInfoParam) GetKnowledgeTargetName() *string {
	return s.KnowledgeTargetName
}

func (s *OperatorBasicInfoParam) GetKnowledgeTargetType() *int32 {
	return s.KnowledgeTargetType
}

func (s *OperatorBasicInfoParam) GetLgfSentences() []*string {
	return s.LgfSentences
}

func (s *OperatorBasicInfoParam) GetMaxEmotionChangeValue() *int32 {
	return s.MaxEmotionChangeValue
}

func (s *OperatorBasicInfoParam) GetMinWordSize() *int32 {
	return s.MinWordSize
}

func (s *OperatorBasicInfoParam) GetNearDialogue() *bool {
	return s.NearDialogue
}

func (s *OperatorBasicInfoParam) GetNotRegex() *string {
	return s.NotRegex
}

func (s *OperatorBasicInfoParam) GetPhrase() *string {
	return s.Phrase
}

func (s *OperatorBasicInfoParam) GetPkey() *string {
	return s.Pkey
}

func (s *OperatorBasicInfoParam) GetPoutputType() *int32 {
	return s.PoutputType
}

func (s *OperatorBasicInfoParam) GetPvalues() []*string {
	return s.Pvalues
}

func (s *OperatorBasicInfoParam) GetQuestionThreshold() *string {
	return s.QuestionThreshold
}

func (s *OperatorBasicInfoParam) GetReferences() []*string {
	return s.References
}

func (s *OperatorBasicInfoParam) GetRegex() *string {
	return s.Regex
}

func (s *OperatorBasicInfoParam) GetRoleId() *int32 {
	return s.RoleId
}

func (s *OperatorBasicInfoParam) GetScore() *int32 {
	return s.Score
}

func (s *OperatorBasicInfoParam) GetSimilarityThreshold() *float64 {
	return s.SimilarityThreshold
}

func (s *OperatorBasicInfoParam) GetSimilarlySentences() []*string {
	return s.SimilarlySentences
}

func (s *OperatorBasicInfoParam) GetSynonyms() map[string][]*string {
	return s.Synonyms
}

func (s *OperatorBasicInfoParam) GetTarget() *int32 {
	return s.Target
}

func (s *OperatorBasicInfoParam) GetTargetRole() *string {
	return s.TargetRole
}

func (s *OperatorBasicInfoParam) GetThreshold() *float32 {
	return s.Threshold
}

func (s *OperatorBasicInfoParam) GetUseEasAlgorithm() *bool {
	return s.UseEasAlgorithm
}

func (s *OperatorBasicInfoParam) GetVelocity() *float64 {
	return s.Velocity
}

func (s *OperatorBasicInfoParam) GetVelocityInMint() *int32 {
	return s.VelocityInMint
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

func (s *OperatorBasicInfoParam) Validate() error {
	if s.CustomerParam != nil {
		if err := s.CustomerParam.Validate(); err != nil {
			return err
		}
	}
	return nil
}
