// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResultToReviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetResultToReviewResponseBody
	GetCode() *string
	SetData(v *GetResultToReviewResponseBodyData) *GetResultToReviewResponseBody
	GetData() *GetResultToReviewResponseBodyData
	SetMessage(v string) *GetResultToReviewResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetResultToReviewResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetResultToReviewResponseBody
	GetSuccess() *bool
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
	return dara.Prettify(s)
}

func (s GetResultToReviewResponseBody) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetResultToReviewResponseBody) GetData() *GetResultToReviewResponseBodyData {
	return s.Data
}

func (s *GetResultToReviewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetResultToReviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResultToReviewResponseBody) GetSuccess() *bool {
	return s.Success
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

func (s *GetResultToReviewResponseBody) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s GetResultToReviewResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyData) GetAudioScheme() *string {
	return s.AudioScheme
}

func (s *GetResultToReviewResponseBodyData) GetAudioURL() *string {
	return s.AudioURL
}

func (s *GetResultToReviewResponseBodyData) GetComments() *string {
	return s.Comments
}

func (s *GetResultToReviewResponseBodyData) GetDialogues() *GetResultToReviewResponseBodyDataDialogues {
	return s.Dialogues
}

func (s *GetResultToReviewResponseBodyData) GetFileId() *string {
	return s.FileId
}

func (s *GetResultToReviewResponseBodyData) GetFileMergeName() *string {
	return s.FileMergeName
}

func (s *GetResultToReviewResponseBodyData) GetHitRuleReviewInfoList() *GetResultToReviewResponseBodyDataHitRuleReviewInfoList {
	return s.HitRuleReviewInfoList
}

func (s *GetResultToReviewResponseBodyData) GetManualScoreInfoList() *GetResultToReviewResponseBodyDataManualScoreInfoList {
	return s.ManualScoreInfoList
}

func (s *GetResultToReviewResponseBodyData) GetReviewHistoryList() *GetResultToReviewResponseBodyDataReviewHistoryList {
	return s.ReviewHistoryList
}

func (s *GetResultToReviewResponseBodyData) GetReviewTypeIdList() *GetResultToReviewResponseBodyDataReviewTypeIdList {
	return s.ReviewTypeIdList
}

func (s *GetResultToReviewResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *GetResultToReviewResponseBodyData) GetTotalScore() *int32 {
	return s.TotalScore
}

func (s *GetResultToReviewResponseBodyData) GetVid() *string {
	return s.Vid
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

func (s *GetResultToReviewResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetResultToReviewResponseBodyDataDialogues struct {
	Dialogue []*GetResultToReviewResponseBodyDataDialoguesDialogue `json:"Dialogue,omitempty" xml:"Dialogue,omitempty" type:"Repeated"`
}

func (s GetResultToReviewResponseBodyDataDialogues) String() string {
	return dara.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataDialogues) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataDialogues) GetDialogue() []*GetResultToReviewResponseBodyDataDialoguesDialogue {
	return s.Dialogue
}

func (s *GetResultToReviewResponseBodyDataDialogues) SetDialogue(v []*GetResultToReviewResponseBodyDataDialoguesDialogue) *GetResultToReviewResponseBodyDataDialogues {
	s.Dialogue = v
	return s
}

func (s *GetResultToReviewResponseBodyDataDialogues) Validate() error {
	return dara.Validate(s)
}

type GetResultToReviewResponseBodyDataDialoguesDialogue struct {
	// example:
	//
	// 72000
	Begin *int64 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// example:
	//
	// 2019-10-01 11:12:01
	BeginTime   *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	BeginTimeMs *int64  `json:"BeginTimeMs,omitempty" xml:"BeginTimeMs,omitempty"`
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
	return dara.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataDialoguesDialogue) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataDialoguesDialogue) GetBegin() *int64 {
	return s.Begin
}

func (s *GetResultToReviewResponseBodyDataDialoguesDialogue) GetBeginTime() *string {
	return s.BeginTime
}

func (s *GetResultToReviewResponseBodyDataDialoguesDialogue) GetBeginTimeMs() *int64 {
	return s.BeginTimeMs
}

func (s *GetResultToReviewResponseBodyDataDialoguesDialogue) GetEmotionValue() *int32 {
	return s.EmotionValue
}

func (s *GetResultToReviewResponseBodyDataDialoguesDialogue) GetEnd() *int64 {
	return s.End
}

func (s *GetResultToReviewResponseBodyDataDialoguesDialogue) GetHourMinSec() *string {
	return s.HourMinSec
}

func (s *GetResultToReviewResponseBodyDataDialoguesDialogue) GetIdentity() *string {
	return s.Identity
}

func (s *GetResultToReviewResponseBodyDataDialoguesDialogue) GetRole() *string {
	return s.Role
}

func (s *GetResultToReviewResponseBodyDataDialoguesDialogue) GetSilenceDuration() *int32 {
	return s.SilenceDuration
}

func (s *GetResultToReviewResponseBodyDataDialoguesDialogue) GetSpeechRate() *int32 {
	return s.SpeechRate
}

func (s *GetResultToReviewResponseBodyDataDialoguesDialogue) GetWords() *string {
	return s.Words
}

func (s *GetResultToReviewResponseBodyDataDialoguesDialogue) SetBegin(v int64) *GetResultToReviewResponseBodyDataDialoguesDialogue {
	s.Begin = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataDialoguesDialogue) SetBeginTime(v string) *GetResultToReviewResponseBodyDataDialoguesDialogue {
	s.BeginTime = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataDialoguesDialogue) SetBeginTimeMs(v int64) *GetResultToReviewResponseBodyDataDialoguesDialogue {
	s.BeginTimeMs = &v
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

func (s *GetResultToReviewResponseBodyDataDialoguesDialogue) Validate() error {
	return dara.Validate(s)
}

type GetResultToReviewResponseBodyDataHitRuleReviewInfoList struct {
	HitRuleReviewInfo []*GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo `json:"HitRuleReviewInfo,omitempty" xml:"HitRuleReviewInfo,omitempty" type:"Repeated"`
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoList) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoList) GetHitRuleReviewInfo() []*GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo {
	return s.HitRuleReviewInfo
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoList) SetHitRuleReviewInfo(v []*GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) *GetResultToReviewResponseBodyDataHitRuleReviewInfoList {
	s.HitRuleReviewInfo = v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoList) Validate() error {
	return dara.Validate(s)
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
	MachineHitResult     *int32                                                                                       `json:"MachineHitResult,omitempty" xml:"MachineHitResult,omitempty"`
	ReviewHitResult      *int32                                                                                       `json:"ReviewHitResult,omitempty" xml:"ReviewHitResult,omitempty"`
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
	return dara.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) GetAutoReview() *int32 {
	return s.AutoReview
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) GetComplainHistories() *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistories {
	return s.ComplainHistories
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) GetComplainable() *bool {
	return s.Complainable
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) GetConditionHitInfoList() *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoList {
	return s.ConditionHitInfoList
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) GetMachineHitResult() *int32 {
	return s.MachineHitResult
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) GetReviewHitResult() *int32 {
	return s.ReviewHitResult
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) GetReviewInfo() *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo {
	return s.ReviewInfo
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) GetRid() *int64 {
	return s.Rid
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) GetRuleName() *string {
	return s.RuleName
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) GetScoreId() *int64 {
	return s.ScoreId
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) GetScoreNum() *int32 {
	return s.ScoreNum
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) GetScoreSubId() *int64 {
	return s.ScoreSubId
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) GetScoreSubName() *string {
	return s.ScoreSubName
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

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) SetMachineHitResult(v int32) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo {
	s.MachineHitResult = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) SetReviewHitResult(v int32) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo {
	s.ReviewHitResult = &v
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

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) Validate() error {
	return dara.Validate(s)
}

type GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistories struct {
	ComplainHistories []*GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories `json:"ComplainHistories,omitempty" xml:"ComplainHistories,omitempty" type:"Repeated"`
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistories) String() string {
	return dara.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistories) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistories) GetComplainHistories() []*GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories {
	return s.ComplainHistories
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistories) SetComplainHistories(v []*GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistories {
	s.ComplainHistories = v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistories) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories) GetComments() *string {
	return s.Comments
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories) GetOperationTime() *string {
	return s.OperationTime
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories) GetOperationType() *int32 {
	return s.OperationType
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories) GetOperator() *int64 {
	return s.Operator
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories) GetOperatorName() *string {
	return s.OperatorName
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

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories) Validate() error {
	return dara.Validate(s)
}

type GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoList struct {
	ConditionHitInfo []*GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfo `json:"ConditionHitInfo,omitempty" xml:"ConditionHitInfo,omitempty" type:"Repeated"`
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoList) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoList) GetConditionHitInfo() []*GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfo {
	return s.ConditionHitInfo
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoList) SetConditionHitInfo(v []*GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfo) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoList {
	s.ConditionHitInfo = v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoList) Validate() error {
	return dara.Validate(s)
}

type GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfo struct {
	Cid      *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoCid      `json:"Cid,omitempty" xml:"Cid,omitempty" type:"Struct"`
	KeyWords *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWords `json:"KeyWords,omitempty" xml:"KeyWords,omitempty" type:"Struct"`
	Phrase   *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase   `json:"Phrase,omitempty" xml:"Phrase,omitempty" type:"Struct"`
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfo) String() string {
	return dara.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfo) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfo) GetCid() *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoCid {
	return s.Cid
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfo) GetKeyWords() *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWords {
	return s.KeyWords
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfo) GetPhrase() *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase {
	return s.Phrase
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

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfo) Validate() error {
	return dara.Validate(s)
}

type GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoCid struct {
	Cid []*string `json:"Cid,omitempty" xml:"Cid,omitempty" type:"Repeated"`
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoCid) String() string {
	return dara.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoCid) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoCid) GetCid() []*string {
	return s.Cid
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoCid) SetCid(v []*string) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoCid {
	s.Cid = v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoCid) Validate() error {
	return dara.Validate(s)
}

type GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWords struct {
	KeyWord []*GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord `json:"KeyWord,omitempty" xml:"KeyWord,omitempty" type:"Repeated"`
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWords) String() string {
	return dara.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWords) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWords) GetKeyWord() []*GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord {
	return s.KeyWord
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWords) SetKeyWord(v []*GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWords {
	s.KeyWord = v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWords) Validate() error {
	return dara.Validate(s)
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
	From    *int32 `json:"From,omitempty" xml:"From,omitempty"`
	IsMatch *bool  `json:"IsMatch,omitempty" xml:"IsMatch,omitempty"`
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
	return dara.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord) GetCid() *string {
	return s.Cid
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord) GetCustomizeCode() *string {
	return s.CustomizeCode
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord) GetFrom() *int32 {
	return s.From
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord) GetIsMatch() *bool {
	return s.IsMatch
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord) GetPid() *int32 {
	return s.Pid
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord) GetTid() *string {
	return s.Tid
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord) GetTo() *int32 {
	return s.To
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord) GetVal() *string {
	return s.Val
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

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord) SetIsMatch(v bool) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord {
	s.IsMatch = &v
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

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase) GetBegin() *int64 {
	return s.Begin
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase) GetEmotionValue() *int32 {
	return s.EmotionValue
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase) GetEnd() *int64 {
	return s.End
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase) GetIdentity() *string {
	return s.Identity
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase) GetPid() *int32 {
	return s.Pid
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase) GetRole() *string {
	return s.Role
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase) GetWords() *string {
	return s.Words
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

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo) GetHitId() *string {
	return s.HitId
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo) GetReviewResult() *int32 {
	return s.ReviewResult
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo) GetReviewTime() *string {
	return s.ReviewTime
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo) GetReviewer() *string {
	return s.Reviewer
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo) GetRid() *int64 {
	return s.Rid
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

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo) Validate() error {
	return dara.Validate(s)
}

type GetResultToReviewResponseBodyDataManualScoreInfoList struct {
	ManualScoreInfo []*GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo `json:"ManualScoreInfo,omitempty" xml:"ManualScoreInfo,omitempty" type:"Repeated"`
}

func (s GetResultToReviewResponseBodyDataManualScoreInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataManualScoreInfoList) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoList) GetManualScoreInfo() []*GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo {
	return s.ManualScoreInfo
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoList) SetManualScoreInfo(v []*GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo) *GetResultToReviewResponseBodyDataManualScoreInfoList {
	s.ManualScoreInfo = v
	return s
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoList) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo) GetComplainHistories() *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistories {
	return s.ComplainHistories
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo) GetComplainable() *bool {
	return s.Complainable
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo) GetScoreId() *int64 {
	return s.ScoreId
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo) GetScoreNum() *int32 {
	return s.ScoreNum
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo) GetScoreSubId() *int64 {
	return s.ScoreSubId
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo) GetScoreSubName() *string {
	return s.ScoreSubName
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

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo) Validate() error {
	return dara.Validate(s)
}

type GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistories struct {
	ComplainHistories []*GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories `json:"ComplainHistories,omitempty" xml:"ComplainHistories,omitempty" type:"Repeated"`
}

func (s GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistories) String() string {
	return dara.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistories) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistories) GetComplainHistories() []*GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories {
	return s.ComplainHistories
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistories) SetComplainHistories(v []*GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories) *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistories {
	s.ComplainHistories = v
	return s
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistories) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories) GetComments() *string {
	return s.Comments
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories) GetOperationTime() *string {
	return s.OperationTime
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories) GetOperationType() *int32 {
	return s.OperationType
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories) GetOperator() *int64 {
	return s.Operator
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories) GetOperatorName() *string {
	return s.OperatorName
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

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories) Validate() error {
	return dara.Validate(s)
}

type GetResultToReviewResponseBodyDataReviewHistoryList struct {
	ReviewHistory []*GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory `json:"ReviewHistory,omitempty" xml:"ReviewHistory,omitempty" type:"Repeated"`
}

func (s GetResultToReviewResponseBodyDataReviewHistoryList) String() string {
	return dara.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataReviewHistoryList) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryList) GetReviewHistory() []*GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory {
	return s.ReviewHistory
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryList) SetReviewHistory(v []*GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) *GetResultToReviewResponseBodyDataReviewHistoryList {
	s.ReviewHistory = v
	return s
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryList) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) GetComments() *string {
	return s.Comments
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) GetComplainResult() *int32 {
	return s.ComplainResult
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) GetOldScore() *int32 {
	return s.OldScore
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) GetOperator() *int64 {
	return s.Operator
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) GetOperatorName() *string {
	return s.OperatorName
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) GetReviewManagerType() *string {
	return s.ReviewManagerType
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) GetReviewResult() *int32 {
	return s.ReviewResult
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) GetReviewRightRule() *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRule {
	return s.ReviewRightRule
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) GetScore() *int32 {
	return s.Score
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) GetTime() *int64 {
	return s.Time
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) GetTimeStr() *string {
	return s.TimeStr
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) GetType() *int32 {
	return s.Type
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

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) Validate() error {
	return dara.Validate(s)
}

type GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRule struct {
	ReviewRightRule []*GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule `json:"ReviewRightRule,omitempty" xml:"ReviewRightRule,omitempty" type:"Repeated"`
}

func (s GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRule) String() string {
	return dara.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRule) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRule) GetReviewRightRule() []*GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule {
	return s.ReviewRightRule
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRule) SetReviewRightRule(v []*GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule) *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRule {
	s.ReviewRightRule = v
	return s
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRule) Validate() error {
	return dara.Validate(s)
}

type GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule struct {
	Rid      *int64  `json:"rid,omitempty" xml:"rid,omitempty"`
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
}

func (s GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule) String() string {
	return dara.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule) GetRid() *int64 {
	return s.Rid
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule) GetRuleName() *string {
	return s.RuleName
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule) SetRid(v int64) *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule {
	s.Rid = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule) SetRuleName(v string) *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule {
	s.RuleName = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule) Validate() error {
	return dara.Validate(s)
}

type GetResultToReviewResponseBodyDataReviewTypeIdList struct {
	ReviewTypeIdList []*GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdList `json:"ReviewTypeIdList,omitempty" xml:"ReviewTypeIdList,omitempty" type:"Repeated"`
}

func (s GetResultToReviewResponseBodyDataReviewTypeIdList) String() string {
	return dara.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataReviewTypeIdList) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataReviewTypeIdList) GetReviewTypeIdList() []*GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdList {
	return s.ReviewTypeIdList
}

func (s *GetResultToReviewResponseBodyDataReviewTypeIdList) SetReviewTypeIdList(v []*GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdList) *GetResultToReviewResponseBodyDataReviewTypeIdList {
	s.ReviewTypeIdList = v
	return s
}

func (s *GetResultToReviewResponseBodyDataReviewTypeIdList) Validate() error {
	return dara.Validate(s)
}

type GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdList struct {
	ReviewKeyIdList *GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdListReviewKeyIdList `json:"ReviewKeyIdList,omitempty" xml:"ReviewKeyIdList,omitempty" type:"Struct"`
	ReviewTypeId    *int64                                                                            `json:"ReviewTypeId,omitempty" xml:"ReviewTypeId,omitempty"`
}

func (s GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdList) String() string {
	return dara.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdList) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdList) GetReviewKeyIdList() *GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdListReviewKeyIdList {
	return s.ReviewKeyIdList
}

func (s *GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdList) GetReviewTypeId() *int64 {
	return s.ReviewTypeId
}

func (s *GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdList) SetReviewKeyIdList(v *GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdListReviewKeyIdList) *GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdList {
	s.ReviewKeyIdList = v
	return s
}

func (s *GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdList) SetReviewTypeId(v int64) *GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdList {
	s.ReviewTypeId = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdList) Validate() error {
	return dara.Validate(s)
}

type GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdListReviewKeyIdList struct {
	ReviewKeyIdList []*int64 `json:"ReviewKeyIdList,omitempty" xml:"ReviewKeyIdList,omitempty" type:"Repeated"`
}

func (s GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdListReviewKeyIdList) String() string {
	return dara.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdListReviewKeyIdList) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdListReviewKeyIdList) GetReviewKeyIdList() []*int64 {
	return s.ReviewKeyIdList
}

func (s *GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdListReviewKeyIdList) SetReviewKeyIdList(v []*int64) *GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdListReviewKeyIdList {
	s.ReviewKeyIdList = v
	return s
}

func (s *GetResultToReviewResponseBodyDataReviewTypeIdListReviewTypeIdListReviewKeyIdList) Validate() error {
	return dara.Validate(s)
}
