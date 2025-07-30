// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadDataSyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UploadDataSyncResponseBody
	GetCode() *string
	SetData(v *UploadDataSyncResponseBodyData) *UploadDataSyncResponseBody
	GetData() *UploadDataSyncResponseBodyData
	SetMessage(v string) *UploadDataSyncResponseBody
	GetMessage() *string
	SetRequestId(v string) *UploadDataSyncResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UploadDataSyncResponseBody
	GetSuccess() *bool
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
	return dara.Prettify(s)
}

func (s UploadDataSyncResponseBody) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBody) GetCode() *string {
	return s.Code
}

func (s *UploadDataSyncResponseBody) GetData() *UploadDataSyncResponseBodyData {
	return s.Data
}

func (s *UploadDataSyncResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UploadDataSyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadDataSyncResponseBody) GetSuccess() *bool {
	return s.Success
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

func (s *UploadDataSyncResponseBody) Validate() error {
	return dara.Validate(s)
}

type UploadDataSyncResponseBodyData struct {
	ResultInfo []*UploadDataSyncResponseBodyDataResultInfo `json:"ResultInfo,omitempty" xml:"ResultInfo,omitempty" type:"Repeated"`
}

func (s UploadDataSyncResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UploadDataSyncResponseBodyData) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyData) GetResultInfo() []*UploadDataSyncResponseBodyDataResultInfo {
	return s.ResultInfo
}

func (s *UploadDataSyncResponseBodyData) SetResultInfo(v []*UploadDataSyncResponseBodyDataResultInfo) *UploadDataSyncResponseBodyData {
	s.ResultInfo = v
	return s
}

func (s *UploadDataSyncResponseBodyData) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfo) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfo) GetHandScoreIdList() *UploadDataSyncResponseBodyDataResultInfoHandScoreIdList {
	return s.HandScoreIdList
}

func (s *UploadDataSyncResponseBodyDataResultInfo) GetRules() *UploadDataSyncResponseBodyDataResultInfoRules {
	return s.Rules
}

func (s *UploadDataSyncResponseBodyDataResultInfo) GetScore() *int32 {
	return s.Score
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

func (s *UploadDataSyncResponseBodyDataResultInfo) Validate() error {
	return dara.Validate(s)
}

type UploadDataSyncResponseBodyDataResultInfoHandScoreIdList struct {
	HandScoreIdList []*string `json:"HandScoreIdList,omitempty" xml:"HandScoreIdList,omitempty" type:"Repeated"`
}

func (s UploadDataSyncResponseBodyDataResultInfoHandScoreIdList) String() string {
	return dara.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfoHandScoreIdList) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfoHandScoreIdList) GetHandScoreIdList() []*string {
	return s.HandScoreIdList
}

func (s *UploadDataSyncResponseBodyDataResultInfoHandScoreIdList) SetHandScoreIdList(v []*string) *UploadDataSyncResponseBodyDataResultInfoHandScoreIdList {
	s.HandScoreIdList = v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfoHandScoreIdList) Validate() error {
	return dara.Validate(s)
}

type UploadDataSyncResponseBodyDataResultInfoRules struct {
	RuleHitInfo []*UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfo `json:"RuleHitInfo,omitempty" xml:"RuleHitInfo,omitempty" type:"Repeated"`
}

func (s UploadDataSyncResponseBodyDataResultInfoRules) String() string {
	return dara.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfoRules) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfoRules) GetRuleHitInfo() []*UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfo {
	return s.RuleHitInfo
}

func (s *UploadDataSyncResponseBodyDataResultInfoRules) SetRuleHitInfo(v []*UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfo) *UploadDataSyncResponseBodyDataResultInfoRules {
	s.RuleHitInfo = v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfoRules) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfo) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfo) GetConditionInfo() *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfo {
	return s.ConditionInfo
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfo) GetHit() *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHit {
	return s.Hit
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfo) GetRid() *string {
	return s.Rid
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfo) GetTid() *string {
	return s.Tid
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

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfo) Validate() error {
	return dara.Validate(s)
}

type UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfo struct {
	ConditionBasicInfo []*UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfoConditionBasicInfo `json:"ConditionBasicInfo,omitempty" xml:"ConditionBasicInfo,omitempty" type:"Repeated"`
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfo) String() string {
	return dara.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfo) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfo) GetConditionBasicInfo() []*UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfoConditionBasicInfo {
	return s.ConditionBasicInfo
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfo) SetConditionBasicInfo(v []*UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfoConditionBasicInfo) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfo {
	s.ConditionBasicInfo = v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfo) Validate() error {
	return dara.Validate(s)
}

type UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfoConditionBasicInfo struct {
	// example:
	//
	// xxx
	ConditionInfoCid *string `json:"ConditionInfoCid,omitempty" xml:"ConditionInfoCid,omitempty"`
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfoConditionBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfoConditionBasicInfo) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfoConditionBasicInfo) GetConditionInfoCid() *string {
	return s.ConditionInfoCid
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfoConditionBasicInfo) SetConditionInfoCid(v string) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfoConditionBasicInfo {
	s.ConditionInfoCid = &v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfoConditionBasicInfo) Validate() error {
	return dara.Validate(s)
}

type UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHit struct {
	ConditionHitInfo []*UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo `json:"ConditionHitInfo,omitempty" xml:"ConditionHitInfo,omitempty" type:"Repeated"`
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHit) String() string {
	return dara.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHit) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHit) GetConditionHitInfo() []*UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo {
	return s.ConditionHitInfo
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHit) SetConditionHitInfo(v []*UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHit {
	s.ConditionHitInfo = v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHit) Validate() error {
	return dara.Validate(s)
}

type UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo struct {
	HitCids     *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitCids     `json:"HitCids,omitempty" xml:"HitCids,omitempty" type:"Struct"`
	HitKeyWords *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWords `json:"HitKeyWords,omitempty" xml:"HitKeyWords,omitempty" type:"Struct"`
	Phrase      *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase      `json:"Phrase,omitempty" xml:"Phrase,omitempty" type:"Struct"`
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo) String() string {
	return dara.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo) GetHitCids() *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitCids {
	return s.HitCids
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo) GetHitKeyWords() *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWords {
	return s.HitKeyWords
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo) GetPhrase() *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase {
	return s.Phrase
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

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo) Validate() error {
	return dara.Validate(s)
}

type UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitCids struct {
	CidItem []*string `json:"CidItem,omitempty" xml:"CidItem,omitempty" type:"Repeated"`
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitCids) String() string {
	return dara.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitCids) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitCids) GetCidItem() []*string {
	return s.CidItem
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitCids) SetCidItem(v []*string) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitCids {
	s.CidItem = v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitCids) Validate() error {
	return dara.Validate(s)
}

type UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWords struct {
	HitKeyWord []*UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord `json:"HitKeyWord,omitempty" xml:"HitKeyWord,omitempty" type:"Repeated"`
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWords) String() string {
	return dara.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWords) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWords) GetHitKeyWord() []*UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord {
	return s.HitKeyWord
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWords) SetHitKeyWord(v []*UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWords {
	s.HitKeyWord = v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWords) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) GetFrom() *int32 {
	return s.From
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) GetPid() *int32 {
	return s.Pid
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) GetTid() *string {
	return s.Tid
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) GetTo() *int32 {
	return s.To
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) GetVal() *string {
	return s.Val
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

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) GetBegin() *int64 {
	return s.Begin
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) GetBeginTime() *string {
	return s.BeginTime
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) GetEnd() *int64 {
	return s.End
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) GetIdentity() *string {
	return s.Identity
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) GetRole() *string {
	return s.Role
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) GetWords() *string {
	return s.Words
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

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) Validate() error {
	return dara.Validate(s)
}
