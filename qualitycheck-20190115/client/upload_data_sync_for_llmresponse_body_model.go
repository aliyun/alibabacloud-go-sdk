// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadDataSyncForLLMResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UploadDataSyncForLLMResponseBody
	GetCode() *string
	SetData(v *UploadDataSyncForLLMResponseBodyData) *UploadDataSyncForLLMResponseBody
	GetData() *UploadDataSyncForLLMResponseBodyData
	SetMessage(v string) *UploadDataSyncForLLMResponseBody
	GetMessage() *string
	SetRequestId(v string) *UploadDataSyncForLLMResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UploadDataSyncForLLMResponseBody
	GetSuccess() *bool
}

type UploadDataSyncForLLMResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *UploadDataSyncForLLMResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadDataSyncForLLMResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadDataSyncForLLMResponseBody) GoString() string {
	return s.String()
}

func (s *UploadDataSyncForLLMResponseBody) GetCode() *string {
	return s.Code
}

func (s *UploadDataSyncForLLMResponseBody) GetData() *UploadDataSyncForLLMResponseBodyData {
	return s.Data
}

func (s *UploadDataSyncForLLMResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UploadDataSyncForLLMResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadDataSyncForLLMResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UploadDataSyncForLLMResponseBody) SetCode(v string) *UploadDataSyncForLLMResponseBody {
	s.Code = &v
	return s
}

func (s *UploadDataSyncForLLMResponseBody) SetData(v *UploadDataSyncForLLMResponseBodyData) *UploadDataSyncForLLMResponseBody {
	s.Data = v
	return s
}

func (s *UploadDataSyncForLLMResponseBody) SetMessage(v string) *UploadDataSyncForLLMResponseBody {
	s.Message = &v
	return s
}

func (s *UploadDataSyncForLLMResponseBody) SetRequestId(v string) *UploadDataSyncForLLMResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadDataSyncForLLMResponseBody) SetSuccess(v bool) *UploadDataSyncForLLMResponseBody {
	s.Success = &v
	return s
}

func (s *UploadDataSyncForLLMResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UploadDataSyncForLLMResponseBodyData struct {
	ResultInfo []*UploadDataSyncForLLMResponseBodyDataResultInfo `json:"ResultInfo,omitempty" xml:"ResultInfo,omitempty" type:"Repeated"`
}

func (s UploadDataSyncForLLMResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UploadDataSyncForLLMResponseBodyData) GoString() string {
	return s.String()
}

func (s *UploadDataSyncForLLMResponseBodyData) GetResultInfo() []*UploadDataSyncForLLMResponseBodyDataResultInfo {
	return s.ResultInfo
}

func (s *UploadDataSyncForLLMResponseBodyData) SetResultInfo(v []*UploadDataSyncForLLMResponseBodyDataResultInfo) *UploadDataSyncForLLMResponseBodyData {
	s.ResultInfo = v
	return s
}

func (s *UploadDataSyncForLLMResponseBodyData) Validate() error {
	if s.ResultInfo != nil {
		for _, item := range s.ResultInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UploadDataSyncForLLMResponseBodyDataResultInfo struct {
	Rules          *UploadDataSyncForLLMResponseBodyDataResultInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
	Score          *int32                                               `json:"Score,omitempty" xml:"Score,omitempty"`
	TyxmPlusCount  *string                                              `json:"TyxmPlusCount,omitempty" xml:"TyxmPlusCount,omitempty"`
	TyxmTurboCount *string                                              `json:"TyxmTurboCount,omitempty" xml:"TyxmTurboCount,omitempty"`
}

func (s UploadDataSyncForLLMResponseBodyDataResultInfo) String() string {
	return dara.Prettify(s)
}

func (s UploadDataSyncForLLMResponseBodyDataResultInfo) GoString() string {
	return s.String()
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfo) GetRules() *UploadDataSyncForLLMResponseBodyDataResultInfoRules {
	return s.Rules
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfo) GetScore() *int32 {
	return s.Score
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfo) GetTyxmPlusCount() *string {
	return s.TyxmPlusCount
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfo) GetTyxmTurboCount() *string {
	return s.TyxmTurboCount
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfo) SetRules(v *UploadDataSyncForLLMResponseBodyDataResultInfoRules) *UploadDataSyncForLLMResponseBodyDataResultInfo {
	s.Rules = v
	return s
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfo) SetScore(v int32) *UploadDataSyncForLLMResponseBodyDataResultInfo {
	s.Score = &v
	return s
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfo) SetTyxmPlusCount(v string) *UploadDataSyncForLLMResponseBodyDataResultInfo {
	s.TyxmPlusCount = &v
	return s
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfo) SetTyxmTurboCount(v string) *UploadDataSyncForLLMResponseBodyDataResultInfo {
	s.TyxmTurboCount = &v
	return s
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfo) Validate() error {
	if s.Rules != nil {
		if err := s.Rules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UploadDataSyncForLLMResponseBodyDataResultInfoRules struct {
	RuleHitInfo []*UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfo `json:"RuleHitInfo,omitempty" xml:"RuleHitInfo,omitempty" type:"Repeated"`
}

func (s UploadDataSyncForLLMResponseBodyDataResultInfoRules) String() string {
	return dara.Prettify(s)
}

func (s UploadDataSyncForLLMResponseBodyDataResultInfoRules) GoString() string {
	return s.String()
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRules) GetRuleHitInfo() []*UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfo {
	return s.RuleHitInfo
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRules) SetRuleHitInfo(v []*UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfo) *UploadDataSyncForLLMResponseBodyDataResultInfoRules {
	s.RuleHitInfo = v
	return s
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRules) Validate() error {
	if s.RuleHitInfo != nil {
		for _, item := range s.RuleHitInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfo struct {
	ConditionInfo *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoConditionInfo `json:"ConditionInfo,omitempty" xml:"ConditionInfo,omitempty" type:"Struct"`
	Hit           *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHit           `json:"Hit,omitempty" xml:"Hit,omitempty" type:"Struct"`
	Rid           *string                                                                      `json:"Rid,omitempty" xml:"Rid,omitempty"`
	Tid           *string                                                                      `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfo) String() string {
	return dara.Prettify(s)
}

func (s UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfo) GoString() string {
	return s.String()
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfo) GetConditionInfo() *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoConditionInfo {
	return s.ConditionInfo
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfo) GetHit() *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHit {
	return s.Hit
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfo) GetRid() *string {
	return s.Rid
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfo) GetTid() *string {
	return s.Tid
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfo) SetConditionInfo(v *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoConditionInfo) *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfo {
	s.ConditionInfo = v
	return s
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfo) SetHit(v *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHit) *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfo {
	s.Hit = v
	return s
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfo) SetRid(v string) *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfo {
	s.Rid = &v
	return s
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfo) SetTid(v string) *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfo {
	s.Tid = &v
	return s
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfo) Validate() error {
	if s.ConditionInfo != nil {
		if err := s.ConditionInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Hit != nil {
		if err := s.Hit.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoConditionInfo struct {
	ConditionBasicInfo []*UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoConditionInfoConditionBasicInfo `json:"ConditionBasicInfo,omitempty" xml:"ConditionBasicInfo,omitempty" type:"Repeated"`
}

func (s UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoConditionInfo) String() string {
	return dara.Prettify(s)
}

func (s UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoConditionInfo) GoString() string {
	return s.String()
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoConditionInfo) GetConditionBasicInfo() []*UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoConditionInfoConditionBasicInfo {
	return s.ConditionBasicInfo
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoConditionInfo) SetConditionBasicInfo(v []*UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoConditionInfoConditionBasicInfo) *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoConditionInfo {
	s.ConditionBasicInfo = v
	return s
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoConditionInfo) Validate() error {
	if s.ConditionBasicInfo != nil {
		for _, item := range s.ConditionBasicInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoConditionInfoConditionBasicInfo struct {
	ConditionInfoCid *string `json:"ConditionInfoCid,omitempty" xml:"ConditionInfoCid,omitempty"`
}

func (s UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoConditionInfoConditionBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoConditionInfoConditionBasicInfo) GoString() string {
	return s.String()
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoConditionInfoConditionBasicInfo) GetConditionInfoCid() *string {
	return s.ConditionInfoCid
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoConditionInfoConditionBasicInfo) SetConditionInfoCid(v string) *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoConditionInfoConditionBasicInfo {
	s.ConditionInfoCid = &v
	return s
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoConditionInfoConditionBasicInfo) Validate() error {
	return dara.Validate(s)
}

type UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHit struct {
	ConditionHitInfo []*UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo `json:"ConditionHitInfo,omitempty" xml:"ConditionHitInfo,omitempty" type:"Repeated"`
}

func (s UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHit) String() string {
	return dara.Prettify(s)
}

func (s UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHit) GoString() string {
	return s.String()
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHit) GetConditionHitInfo() []*UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo {
	return s.ConditionHitInfo
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHit) SetConditionHitInfo(v []*UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo) *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHit {
	s.ConditionHitInfo = v
	return s
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHit) Validate() error {
	if s.ConditionHitInfo != nil {
		for _, item := range s.ConditionHitInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo struct {
	HitCids     *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitCids     `json:"HitCids,omitempty" xml:"HitCids,omitempty" type:"Struct"`
	HitKeyWords *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWords `json:"HitKeyWords,omitempty" xml:"HitKeyWords,omitempty" type:"Struct"`
	Phrase      *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase      `json:"Phrase,omitempty" xml:"Phrase,omitempty" type:"Struct"`
}

func (s UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo) String() string {
	return dara.Prettify(s)
}

func (s UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo) GoString() string {
	return s.String()
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo) GetHitCids() *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitCids {
	return s.HitCids
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo) GetHitKeyWords() *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWords {
	return s.HitKeyWords
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo) GetPhrase() *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase {
	return s.Phrase
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo) SetHitCids(v *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitCids) *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo {
	s.HitCids = v
	return s
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo) SetHitKeyWords(v *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWords) *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo {
	s.HitKeyWords = v
	return s
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo) SetPhrase(v *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo {
	s.Phrase = v
	return s
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo) Validate() error {
	if s.HitCids != nil {
		if err := s.HitCids.Validate(); err != nil {
			return err
		}
	}
	if s.HitKeyWords != nil {
		if err := s.HitKeyWords.Validate(); err != nil {
			return err
		}
	}
	if s.Phrase != nil {
		if err := s.Phrase.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitCids struct {
	CidItem []*string `json:"CidItem,omitempty" xml:"CidItem,omitempty" type:"Repeated"`
}

func (s UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitCids) String() string {
	return dara.Prettify(s)
}

func (s UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitCids) GoString() string {
	return s.String()
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitCids) GetCidItem() []*string {
	return s.CidItem
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitCids) SetCidItem(v []*string) *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitCids {
	s.CidItem = v
	return s
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitCids) Validate() error {
	return dara.Validate(s)
}

type UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWords struct {
	HitKeyWord []*UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord `json:"HitKeyWord,omitempty" xml:"HitKeyWord,omitempty" type:"Repeated"`
}

func (s UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWords) String() string {
	return dara.Prettify(s)
}

func (s UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWords) GoString() string {
	return s.String()
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWords) GetHitKeyWord() []*UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord {
	return s.HitKeyWord
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWords) SetHitKeyWord(v []*UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWords {
	s.HitKeyWord = v
	return s
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWords) Validate() error {
	if s.HitKeyWord != nil {
		for _, item := range s.HitKeyWord {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord struct {
	From *int32  `json:"From,omitempty" xml:"From,omitempty"`
	Pid  *int32  `json:"Pid,omitempty" xml:"Pid,omitempty"`
	Tid  *string `json:"Tid,omitempty" xml:"Tid,omitempty"`
	To   *int32  `json:"To,omitempty" xml:"To,omitempty"`
	Val  *string `json:"Val,omitempty" xml:"Val,omitempty"`
}

func (s UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) String() string {
	return dara.Prettify(s)
}

func (s UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) GoString() string {
	return s.String()
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) GetFrom() *int32 {
	return s.From
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) GetPid() *int32 {
	return s.Pid
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) GetTid() *string {
	return s.Tid
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) GetTo() *int32 {
	return s.To
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) GetVal() *string {
	return s.Val
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) SetFrom(v int32) *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord {
	s.From = &v
	return s
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) SetPid(v int32) *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord {
	s.Pid = &v
	return s
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) SetTid(v string) *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord {
	s.Tid = &v
	return s
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) SetTo(v int32) *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord {
	s.To = &v
	return s
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) SetVal(v string) *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord {
	s.Val = &v
	return s
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) Validate() error {
	return dara.Validate(s)
}

type UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase struct {
	Begin     *int64  `json:"Begin,omitempty" xml:"Begin,omitempty"`
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	End       *int64  `json:"End,omitempty" xml:"End,omitempty"`
	Identity  *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
	Role      *string `json:"Role,omitempty" xml:"Role,omitempty"`
	Words     *string `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) String() string {
	return dara.Prettify(s)
}

func (s UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) GoString() string {
	return s.String()
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) GetBegin() *int64 {
	return s.Begin
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) GetBeginTime() *string {
	return s.BeginTime
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) GetEnd() *int64 {
	return s.End
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) GetIdentity() *string {
	return s.Identity
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) GetRole() *string {
	return s.Role
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) GetWords() *string {
	return s.Words
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) SetBegin(v int64) *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase {
	s.Begin = &v
	return s
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) SetBeginTime(v string) *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase {
	s.BeginTime = &v
	return s
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) SetEnd(v int64) *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase {
	s.End = &v
	return s
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) SetIdentity(v string) *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase {
	s.Identity = &v
	return s
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) SetRole(v string) *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase {
	s.Role = &v
	return s
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) SetWords(v string) *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase {
	s.Words = &v
	return s
}

func (s *UploadDataSyncForLLMResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) Validate() error {
	return dara.Validate(s)
}
