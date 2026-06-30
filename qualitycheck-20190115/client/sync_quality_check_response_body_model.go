// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncQualityCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SyncQualityCheckResponseBody
	GetCode() *string
	SetData(v *SyncQualityCheckResponseBodyData) *SyncQualityCheckResponseBody
	GetData() *SyncQualityCheckResponseBodyData
	SetMessage(v string) *SyncQualityCheckResponseBody
	GetMessage() *string
	SetRequestId(v string) *SyncQualityCheckResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SyncQualityCheckResponseBody
	GetSuccess() *bool
}

type SyncQualityCheckResponseBody struct {
	// Result status code. 200 indicates success. Other values indicate failure. The caller can determine the reason for failure using this field.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned result, including hit information.
	Data *SyncQualityCheckResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Error details if an error occurs. "successful" if successful.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Unique request identifier.
	//
	// example:
	//
	// 66E1ACB8-17B2-4BE8-8581-954A8*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. The caller can use this field to determine if the request succeeded: true for success; false/null for failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SyncQualityCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SyncQualityCheckResponseBody) GoString() string {
	return s.String()
}

func (s *SyncQualityCheckResponseBody) GetCode() *string {
	return s.Code
}

func (s *SyncQualityCheckResponseBody) GetData() *SyncQualityCheckResponseBodyData {
	return s.Data
}

func (s *SyncQualityCheckResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SyncQualityCheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SyncQualityCheckResponseBody) GetSuccess() *bool {
	return s.Success
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

func (s *SyncQualityCheckResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SyncQualityCheckResponseBodyData struct {
	// Time of recording and dialogue occurrence, in milliseconds since January 1, 1970, 00:00:00 UTC (UNIX timestamp in milliseconds, such as 1584535485856).
	//
	// example:
	//
	// 1584535485856
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// List of hit rules. Each item is a rule. Only hit rule information and hit rule location information are returned.
	Rules []*SyncQualityCheckResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// Final score, with a maximum of 100.
	//
	// example:
	//
	// 100
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// Task ID.
	//
	// example:
	//
	// 66E1ACB866E1ACB8
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Unique identifier for the current conversation.
	//
	// example:
	//
	// 20200876-66E1ACB8
	Tid *string `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s SyncQualityCheckResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SyncQualityCheckResponseBodyData) GoString() string {
	return s.String()
}

func (s *SyncQualityCheckResponseBodyData) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *SyncQualityCheckResponseBodyData) GetRules() []*SyncQualityCheckResponseBodyDataRules {
	return s.Rules
}

func (s *SyncQualityCheckResponseBodyData) GetScore() *int32 {
	return s.Score
}

func (s *SyncQualityCheckResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SyncQualityCheckResponseBodyData) GetTid() *string {
	return s.Tid
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

func (s *SyncQualityCheckResponseBodyData) Validate() error {
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

type SyncQualityCheckResponseBodyDataRules struct {
	// List of hit sentences. For this API, if a hit occurs, it is a single data entry.
	Hit []*SyncQualityCheckResponseBodyDataRulesHit `json:"Hit,omitempty" xml:"Hit,omitempty" type:"Repeated"`
	// ID of the hit rule.
	//
	// example:
	//
	// 232232
	Rid *string `json:"Rid,omitempty" xml:"Rid,omitempty"`
	// Rule basic information
	RuleInfoBase *SyncQualityCheckResponseBodyDataRulesRuleInfoBase `json:"RuleInfoBase,omitempty" xml:"RuleInfoBase,omitempty" type:"Struct"`
	// Name of the hit rule.
	//
	// example:
	//
	// 禁用语
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s SyncQualityCheckResponseBodyDataRules) String() string {
	return dara.Prettify(s)
}

func (s SyncQualityCheckResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *SyncQualityCheckResponseBodyDataRules) GetHit() []*SyncQualityCheckResponseBodyDataRulesHit {
	return s.Hit
}

func (s *SyncQualityCheckResponseBodyDataRules) GetRid() *string {
	return s.Rid
}

func (s *SyncQualityCheckResponseBodyDataRules) GetRuleInfoBase() *SyncQualityCheckResponseBodyDataRulesRuleInfoBase {
	return s.RuleInfoBase
}

func (s *SyncQualityCheckResponseBodyDataRules) GetRuleName() *string {
	return s.RuleName
}

func (s *SyncQualityCheckResponseBodyDataRules) SetHit(v []*SyncQualityCheckResponseBodyDataRulesHit) *SyncQualityCheckResponseBodyDataRules {
	s.Hit = v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRules) SetRid(v string) *SyncQualityCheckResponseBodyDataRules {
	s.Rid = &v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRules) SetRuleInfoBase(v *SyncQualityCheckResponseBodyDataRulesRuleInfoBase) *SyncQualityCheckResponseBodyDataRules {
	s.RuleInfoBase = v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRules) SetRuleName(v string) *SyncQualityCheckResponseBodyDataRules {
	s.RuleName = &v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRules) Validate() error {
	if s.Hit != nil {
		for _, item := range s.Hit {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RuleInfoBase != nil {
		if err := s.RuleInfoBase.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SyncQualityCheckResponseBodyDataRulesHit struct {
	// Keywords that met the condition.
	HitKeyWords []*SyncQualityCheckResponseBodyDataRulesHitHitKeyWords `json:"HitKeyWords,omitempty" xml:"HitKeyWords,omitempty" type:"Repeated"`
	// Dialogue content that met the condition.
	Phrase *SyncQualityCheckResponseBodyDataRulesHitPhrase `json:"Phrase,omitempty" xml:"Phrase,omitempty" type:"Struct"`
}

func (s SyncQualityCheckResponseBodyDataRulesHit) String() string {
	return dara.Prettify(s)
}

func (s SyncQualityCheckResponseBodyDataRulesHit) GoString() string {
	return s.String()
}

func (s *SyncQualityCheckResponseBodyDataRulesHit) GetHitKeyWords() []*SyncQualityCheckResponseBodyDataRulesHitHitKeyWords {
	return s.HitKeyWords
}

func (s *SyncQualityCheckResponseBodyDataRulesHit) GetPhrase() *SyncQualityCheckResponseBodyDataRulesHitPhrase {
	return s.Phrase
}

func (s *SyncQualityCheckResponseBodyDataRulesHit) SetHitKeyWords(v []*SyncQualityCheckResponseBodyDataRulesHitHitKeyWords) *SyncQualityCheckResponseBodyDataRulesHit {
	s.HitKeyWords = v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRulesHit) SetPhrase(v *SyncQualityCheckResponseBodyDataRulesHitPhrase) *SyncQualityCheckResponseBodyDataRulesHit {
	s.Phrase = v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRulesHit) Validate() error {
	if s.HitKeyWords != nil {
		for _, item := range s.HitKeyWords {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Phrase != nil {
		if err := s.Phrase.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SyncQualityCheckResponseBodyDataRulesHitHitKeyWords struct {
	// Condition ID of the rule.
	//
	// example:
	//
	// 2312
	Cid *int32 `json:"Cid,omitempty" xml:"Cid,omitempty"`
	// Start position of the keyword.
	//
	// example:
	//
	// 1
	From *int32 `json:"From,omitempty" xml:"From,omitempty"`
	// Index value of the hit sentence in the entire conversation.
	//
	// example:
	//
	// 4
	Pid *int32 `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// End position of the keyword.
	//
	// example:
	//
	// 4
	To *int32 `json:"To,omitempty" xml:"To,omitempty"`
	// Keyword.
	//
	// example:
	//
	// 你好
	Val *string `json:"Val,omitempty" xml:"Val,omitempty"`
}

func (s SyncQualityCheckResponseBodyDataRulesHitHitKeyWords) String() string {
	return dara.Prettify(s)
}

func (s SyncQualityCheckResponseBodyDataRulesHitHitKeyWords) GoString() string {
	return s.String()
}

func (s *SyncQualityCheckResponseBodyDataRulesHitHitKeyWords) GetCid() *int32 {
	return s.Cid
}

func (s *SyncQualityCheckResponseBodyDataRulesHitHitKeyWords) GetFrom() *int32 {
	return s.From
}

func (s *SyncQualityCheckResponseBodyDataRulesHitHitKeyWords) GetPid() *int32 {
	return s.Pid
}

func (s *SyncQualityCheckResponseBodyDataRulesHitHitKeyWords) GetTo() *int32 {
	return s.To
}

func (s *SyncQualityCheckResponseBodyDataRulesHitHitKeyWords) GetVal() *string {
	return s.Val
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

func (s *SyncQualityCheckResponseBodyDataRulesHitHitKeyWords) Validate() error {
	return dara.Validate(s)
}

type SyncQualityCheckResponseBodyDataRulesHitPhrase struct {
	// Start time of this sentence relative to the entire conversation, in milliseconds.
	//
	// example:
	//
	// 1230
	Begin *int64 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// Emotional value of this sentence, 0-10. Higher values indicate stronger emotions.
	//
	// example:
	//
	// 6
	EmotionValue *int32 `json:"EmotionValue,omitempty" xml:"EmotionValue,omitempty"`
	// End time of this sentence relative to the entire conversation, in milliseconds.
	//
	// example:
	//
	// 3440
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// Deprecated field. Ignore it.
	//
	// example:
	//
	// xxx
	Identity *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
	// Role of this sentence. Valid values: customer service representative, customer.
	//
	// example:
	//
	// 客服
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// Deprecated field. Ignore it.
	//
	// example:
	//
	// 123
	SilenceDuration *int32 `json:"SilenceDuration,omitempty" xml:"SilenceDuration,omitempty"`
	// Speech rate of this sentence, in characters per minute.
	//
	// example:
	//
	// 233
	SpeechRate *int32 `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	// Dialogue content.
	//
	// example:
	//
	// 你好请问有什么可以帮您的
	Words *string `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s SyncQualityCheckResponseBodyDataRulesHitPhrase) String() string {
	return dara.Prettify(s)
}

func (s SyncQualityCheckResponseBodyDataRulesHitPhrase) GoString() string {
	return s.String()
}

func (s *SyncQualityCheckResponseBodyDataRulesHitPhrase) GetBegin() *int64 {
	return s.Begin
}

func (s *SyncQualityCheckResponseBodyDataRulesHitPhrase) GetEmotionValue() *int32 {
	return s.EmotionValue
}

func (s *SyncQualityCheckResponseBodyDataRulesHitPhrase) GetEnd() *int64 {
	return s.End
}

func (s *SyncQualityCheckResponseBodyDataRulesHitPhrase) GetIdentity() *string {
	return s.Identity
}

func (s *SyncQualityCheckResponseBodyDataRulesHitPhrase) GetRole() *string {
	return s.Role
}

func (s *SyncQualityCheckResponseBodyDataRulesHitPhrase) GetSilenceDuration() *int32 {
	return s.SilenceDuration
}

func (s *SyncQualityCheckResponseBodyDataRulesHitPhrase) GetSpeechRate() *int32 {
	return s.SpeechRate
}

func (s *SyncQualityCheckResponseBodyDataRulesHitPhrase) GetWords() *string {
	return s.Words
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

func (s *SyncQualityCheckResponseBodyDataRulesHitPhrase) Validate() error {
	return dara.Validate(s)
}

type SyncQualityCheckResponseBodyDataRulesRuleInfoBase struct {
	// Rule remarks
	//
	// example:
	//
	// 邀约客户，客户不同意参加试听
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// Rule importance level
	//
	// example:
	//
	// 2
	Level *int32 `json:"Level,omitempty" xml:"Level,omitempty"`
	// Rule category name
	//
	// 	Notice: The requiredFields parameter must include "ruleInfoBase.ruleCategoryName".
	//
	// example:
	//
	// 服务规范
	RuleCategoryName *string `json:"RuleCategoryName,omitempty" xml:"RuleCategoryName,omitempty"`
	// Score value
	//
	// example:
	//
	// 1
	ScoreNum *int32 `json:"ScoreNum,omitempty" xml:"ScoreNum,omitempty"`
	// Scoring type. 0: bonus/penalty points, 1: one-time score.
	//
	// example:
	//
	// 1
	ScoreNumType *int32 `json:"ScoreNumType,omitempty" xml:"ScoreNumType,omitempty"`
	// 1 for bonus points, 3 for penalty points. Default is 1.
	//
	// example:
	//
	// 1
	ScoreType *int32 `json:"ScoreType,omitempty" xml:"ScoreType,omitempty"`
	// Rule type ID
	//
	// 	Notice: The requiredFields parameter must include "ruleInfoBase".
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SyncQualityCheckResponseBodyDataRulesRuleInfoBase) String() string {
	return dara.Prettify(s)
}

func (s SyncQualityCheckResponseBodyDataRulesRuleInfoBase) GoString() string {
	return s.String()
}

func (s *SyncQualityCheckResponseBodyDataRulesRuleInfoBase) GetComments() *string {
	return s.Comments
}

func (s *SyncQualityCheckResponseBodyDataRulesRuleInfoBase) GetLevel() *int32 {
	return s.Level
}

func (s *SyncQualityCheckResponseBodyDataRulesRuleInfoBase) GetRuleCategoryName() *string {
	return s.RuleCategoryName
}

func (s *SyncQualityCheckResponseBodyDataRulesRuleInfoBase) GetScoreNum() *int32 {
	return s.ScoreNum
}

func (s *SyncQualityCheckResponseBodyDataRulesRuleInfoBase) GetScoreNumType() *int32 {
	return s.ScoreNumType
}

func (s *SyncQualityCheckResponseBodyDataRulesRuleInfoBase) GetScoreType() *int32 {
	return s.ScoreType
}

func (s *SyncQualityCheckResponseBodyDataRulesRuleInfoBase) GetType() *int32 {
	return s.Type
}

func (s *SyncQualityCheckResponseBodyDataRulesRuleInfoBase) SetComments(v string) *SyncQualityCheckResponseBodyDataRulesRuleInfoBase {
	s.Comments = &v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRulesRuleInfoBase) SetLevel(v int32) *SyncQualityCheckResponseBodyDataRulesRuleInfoBase {
	s.Level = &v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRulesRuleInfoBase) SetRuleCategoryName(v string) *SyncQualityCheckResponseBodyDataRulesRuleInfoBase {
	s.RuleCategoryName = &v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRulesRuleInfoBase) SetScoreNum(v int32) *SyncQualityCheckResponseBodyDataRulesRuleInfoBase {
	s.ScoreNum = &v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRulesRuleInfoBase) SetScoreNumType(v int32) *SyncQualityCheckResponseBodyDataRulesRuleInfoBase {
	s.ScoreNumType = &v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRulesRuleInfoBase) SetScoreType(v int32) *SyncQualityCheckResponseBodyDataRulesRuleInfoBase {
	s.ScoreType = &v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRulesRuleInfoBase) SetType(v int32) *SyncQualityCheckResponseBodyDataRulesRuleInfoBase {
	s.Type = &v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRulesRuleInfoBase) Validate() error {
	return dara.Validate(s)
}
