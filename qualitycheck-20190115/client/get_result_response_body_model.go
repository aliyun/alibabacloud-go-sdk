// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetResultResponseBody
	GetCode() *string
	SetCount(v int32) *GetResultResponseBody
	GetCount() *int32
	SetData(v *GetResultResponseBodyData) *GetResultResponseBody
	GetData() *GetResultResponseBodyData
	SetMessage(v string) *GetResultResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *GetResultResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *GetResultResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetResultResponseBody
	GetRequestId() *string
	SetResultCountId(v string) *GetResultResponseBody
	GetResultCountId() *string
	SetSuccess(v bool) *GetResultResponseBody
	GetSuccess() *bool
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
	return dara.Prettify(s)
}

func (s GetResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetResultResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *GetResultResponseBody) GetData() *GetResultResponseBodyData {
	return s.Data
}

func (s *GetResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetResultResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetResultResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResultResponseBody) GetResultCountId() *string {
	return s.ResultCountId
}

func (s *GetResultResponseBody) GetSuccess() *bool {
	return s.Success
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

func (s *GetResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResultResponseBodyData struct {
	ResultInfo []*GetResultResponseBodyDataResultInfo `json:"ResultInfo,omitempty" xml:"ResultInfo,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyData) GetResultInfo() []*GetResultResponseBodyDataResultInfo {
	return s.ResultInfo
}

func (s *GetResultResponseBodyData) SetResultInfo(v []*GetResultResponseBodyDataResultInfo) *GetResultResponseBodyData {
	s.ResultInfo = v
	return s
}

func (s *GetResultResponseBodyData) Validate() error {
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
	Vid      *string `json:"Vid,omitempty" xml:"Vid,omitempty"`
}

func (s GetResultResponseBodyDataResultInfo) String() string {
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfo) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfo) GetAgent() *GetResultResponseBodyDataResultInfoAgent {
	return s.Agent
}

func (s *GetResultResponseBodyDataResultInfo) GetAsrResult() *GetResultResponseBodyDataResultInfoAsrResult {
	return s.AsrResult
}

func (s *GetResultResponseBodyDataResultInfo) GetAssignmentTime() *string {
	return s.AssignmentTime
}

func (s *GetResultResponseBodyDataResultInfo) GetComments() *string {
	return s.Comments
}

func (s *GetResultResponseBodyDataResultInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetResultResponseBodyDataResultInfo) GetCreateTimeLong() *string {
	return s.CreateTimeLong
}

func (s *GetResultResponseBodyDataResultInfo) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetResultResponseBodyDataResultInfo) GetHitResult() *GetResultResponseBodyDataResultInfoHitResult {
	return s.HitResult
}

func (s *GetResultResponseBodyDataResultInfo) GetHitScore() *GetResultResponseBodyDataResultInfoHitScore {
	return s.HitScore
}

func (s *GetResultResponseBodyDataResultInfo) GetLastDataId() *string {
	return s.LastDataId
}

func (s *GetResultResponseBodyDataResultInfo) GetRecording() *GetResultResponseBodyDataResultInfoRecording {
	return s.Recording
}

func (s *GetResultResponseBodyDataResultInfo) GetResolver() *string {
	return s.Resolver
}

func (s *GetResultResponseBodyDataResultInfo) GetReviewHistoryList() *GetResultResponseBodyDataResultInfoReviewHistoryList {
	return s.ReviewHistoryList
}

func (s *GetResultResponseBodyDataResultInfo) GetReviewResult() *int32 {
	return s.ReviewResult
}

func (s *GetResultResponseBodyDataResultInfo) GetReviewStatus() *int32 {
	return s.ReviewStatus
}

func (s *GetResultResponseBodyDataResultInfo) GetReviewTime() *string {
	return s.ReviewTime
}

func (s *GetResultResponseBodyDataResultInfo) GetReviewTimeLong() *string {
	return s.ReviewTimeLong
}

func (s *GetResultResponseBodyDataResultInfo) GetReviewType() *int32 {
	return s.ReviewType
}

func (s *GetResultResponseBodyDataResultInfo) GetReviewTypeIdList() *GetResultResponseBodyDataResultInfoReviewTypeIdList {
	return s.ReviewTypeIdList
}

func (s *GetResultResponseBodyDataResultInfo) GetReviewer() *string {
	return s.Reviewer
}

func (s *GetResultResponseBodyDataResultInfo) GetSchemeIdList() *GetResultResponseBodyDataResultInfoSchemeIdList {
	return s.SchemeIdList
}

func (s *GetResultResponseBodyDataResultInfo) GetSchemeNameList() *GetResultResponseBodyDataResultInfoSchemeNameList {
	return s.SchemeNameList
}

func (s *GetResultResponseBodyDataResultInfo) GetScore() *int32 {
	return s.Score
}

func (s *GetResultResponseBodyDataResultInfo) GetStatus() *int32 {
	return s.Status
}

func (s *GetResultResponseBodyDataResultInfo) GetTaskId() *string {
	return s.TaskId
}

func (s *GetResultResponseBodyDataResultInfo) GetTaskName() *string {
	return s.TaskName
}

func (s *GetResultResponseBodyDataResultInfo) GetVid() *string {
	return s.Vid
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

func (s *GetResultResponseBodyDataResultInfo) SetVid(v string) *GetResultResponseBodyDataResultInfo {
	s.Vid = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) Validate() error {
	if s.Agent != nil {
		if err := s.Agent.Validate(); err != nil {
			return err
		}
	}
	if s.AsrResult != nil {
		if err := s.AsrResult.Validate(); err != nil {
			return err
		}
	}
	if s.HitResult != nil {
		if err := s.HitResult.Validate(); err != nil {
			return err
		}
	}
	if s.HitScore != nil {
		if err := s.HitScore.Validate(); err != nil {
			return err
		}
	}
	if s.Recording != nil {
		if err := s.Recording.Validate(); err != nil {
			return err
		}
	}
	if s.ReviewHistoryList != nil {
		if err := s.ReviewHistoryList.Validate(); err != nil {
			return err
		}
	}
	if s.ReviewTypeIdList != nil {
		if err := s.ReviewTypeIdList.Validate(); err != nil {
			return err
		}
	}
	if s.SchemeIdList != nil {
		if err := s.SchemeIdList.Validate(); err != nil {
			return err
		}
	}
	if s.SchemeNameList != nil {
		if err := s.SchemeNameList.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoAgent) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoAgent) GetId() *string {
	return s.Id
}

func (s *GetResultResponseBodyDataResultInfoAgent) GetName() *string {
	return s.Name
}

func (s *GetResultResponseBodyDataResultInfoAgent) GetSkillGroup() *string {
	return s.SkillGroup
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

func (s *GetResultResponseBodyDataResultInfoAgent) Validate() error {
	return dara.Validate(s)
}

type GetResultResponseBodyDataResultInfoAsrResult struct {
	AsrResult []*GetResultResponseBodyDataResultInfoAsrResultAsrResult `json:"AsrResult,omitempty" xml:"AsrResult,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoAsrResult) String() string {
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoAsrResult) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoAsrResult) GetAsrResult() []*GetResultResponseBodyDataResultInfoAsrResultAsrResult {
	return s.AsrResult
}

func (s *GetResultResponseBodyDataResultInfoAsrResult) SetAsrResult(v []*GetResultResponseBodyDataResultInfoAsrResultAsrResult) *GetResultResponseBodyDataResultInfoAsrResult {
	s.AsrResult = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoAsrResult) Validate() error {
	if s.AsrResult != nil {
		for _, item := range s.AsrResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
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
	End      *int64  `json:"End,omitempty" xml:"End,omitempty"`
	Identity *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
	Role     *string `json:"Role,omitempty" xml:"Role,omitempty"`
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
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoAsrResultAsrResult) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoAsrResultAsrResult) GetBegin() *int64 {
	return s.Begin
}

func (s *GetResultResponseBodyDataResultInfoAsrResultAsrResult) GetEmotionValue() *int32 {
	return s.EmotionValue
}

func (s *GetResultResponseBodyDataResultInfoAsrResultAsrResult) GetEnd() *int64 {
	return s.End
}

func (s *GetResultResponseBodyDataResultInfoAsrResultAsrResult) GetIdentity() *string {
	return s.Identity
}

func (s *GetResultResponseBodyDataResultInfoAsrResultAsrResult) GetRole() *string {
	return s.Role
}

func (s *GetResultResponseBodyDataResultInfoAsrResultAsrResult) GetSpeechRate() *int32 {
	return s.SpeechRate
}

func (s *GetResultResponseBodyDataResultInfoAsrResultAsrResult) GetWords() *string {
	return s.Words
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

func (s *GetResultResponseBodyDataResultInfoAsrResultAsrResult) SetIdentity(v string) *GetResultResponseBodyDataResultInfoAsrResultAsrResult {
	s.Identity = &v
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

func (s *GetResultResponseBodyDataResultInfoAsrResultAsrResult) Validate() error {
	return dara.Validate(s)
}

type GetResultResponseBodyDataResultInfoHitResult struct {
	HitResult []*GetResultResponseBodyDataResultInfoHitResultHitResult `json:"HitResult,omitempty" xml:"HitResult,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoHitResult) String() string {
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResult) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResult) GetHitResult() []*GetResultResponseBodyDataResultInfoHitResultHitResult {
	return s.HitResult
}

func (s *GetResultResponseBodyDataResultInfoHitResult) SetHitResult(v []*GetResultResponseBodyDataResultInfoHitResultHitResult) *GetResultResponseBodyDataResultInfoHitResult {
	s.HitResult = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResult) Validate() error {
	if s.HitResult != nil {
		for _, item := range s.HitResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetResultResponseBodyDataResultInfoHitResultHitResult struct {
	// example:
	//
	// turn
	ArtificialRule *string                                                          `json:"ArtificialRule,omitempty" xml:"ArtificialRule,omitempty"`
	Conditions     *GetResultResponseBodyDataResultInfoHitResultHitResultConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Struct"`
	// example:
	//
	// 1
	FinalHitResult *string                                                    `json:"FinalHitResult,omitempty" xml:"FinalHitResult,omitempty"`
	Hits           *GetResultResponseBodyDataResultInfoHitResultHitResultHits `json:"Hits,omitempty" xml:"Hits,omitempty" type:"Struct"`
	// example:
	//
	// 1
	MachineHitResult *string `json:"MachineHitResult,omitempty" xml:"MachineHitResult,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	Score         *int32  `json:"Score,omitempty" xml:"Score,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResult) String() string {
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResult) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResult) GetArtificialRule() *string {
	return s.ArtificialRule
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResult) GetConditions() *GetResultResponseBodyDataResultInfoHitResultHitResultConditions {
	return s.Conditions
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResult) GetFinalHitResult() *string {
	return s.FinalHitResult
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResult) GetHits() *GetResultResponseBodyDataResultInfoHitResultHitResultHits {
	return s.Hits
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResult) GetMachineHitResult() *string {
	return s.MachineHitResult
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResult) GetName() *string {
	return s.Name
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResult) GetReviewResult() *int32 {
	return s.ReviewResult
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResult) GetRid() *string {
	return s.Rid
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResult) GetSchemeId() *int64 {
	return s.SchemeId
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResult) GetSchemeVersion() *int64 {
	return s.SchemeVersion
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResult) GetScore() *int32 {
	return s.Score
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResult) GetType() *string {
	return s.Type
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResult) SetArtificialRule(v string) *GetResultResponseBodyDataResultInfoHitResultHitResult {
	s.ArtificialRule = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResult) SetConditions(v *GetResultResponseBodyDataResultInfoHitResultHitResultConditions) *GetResultResponseBodyDataResultInfoHitResultHitResult {
	s.Conditions = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResult) SetFinalHitResult(v string) *GetResultResponseBodyDataResultInfoHitResultHitResult {
	s.FinalHitResult = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResult) SetHits(v *GetResultResponseBodyDataResultInfoHitResultHitResultHits) *GetResultResponseBodyDataResultInfoHitResultHitResult {
	s.Hits = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResult) SetMachineHitResult(v string) *GetResultResponseBodyDataResultInfoHitResultHitResult {
	s.MachineHitResult = &v
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

func (s *GetResultResponseBodyDataResultInfoHitResultHitResult) SetScore(v int32) *GetResultResponseBodyDataResultInfoHitResultHitResult {
	s.Score = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResult) SetType(v string) *GetResultResponseBodyDataResultInfoHitResultHitResult {
	s.Type = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResult) Validate() error {
	if s.Conditions != nil {
		if err := s.Conditions.Validate(); err != nil {
			return err
		}
	}
	if s.Hits != nil {
		if err := s.Hits.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResultResponseBodyDataResultInfoHitResultHitResultConditions struct {
	Conditions []*GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditions) String() string {
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditions) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditions) GetConditions() []*GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditions {
	return s.Conditions
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditions) SetConditions(v []*GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditions) *GetResultResponseBodyDataResultInfoHitResultHitResultConditions {
	s.Conditions = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditions) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
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
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditions) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditions) GetCheckRange() *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRange {
	return s.CheckRange
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditions) GetCid() *string {
	return s.Cid
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditions) GetExclusion() *int32 {
	return s.Exclusion
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditions) GetId() *int64 {
	return s.Id
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditions) GetLambda() *string {
	return s.Lambda
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditions) GetOperators() *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperators {
	return s.Operators
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditions) GetRid() *string {
	return s.Rid
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

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditions) Validate() error {
	if s.CheckRange != nil {
		if err := s.CheckRange.Validate(); err != nil {
			return err
		}
	}
	if s.Operators != nil {
		if err := s.Operators.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRange) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRange) GetAbsolute() *bool {
	return s.Absolute
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRange) GetAllSentencesSatisfy() *bool {
	return s.AllSentencesSatisfy
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRange) GetAnchor() *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeAnchor {
	return s.Anchor
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRange) GetRange() *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeRange {
	return s.Range
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRange) GetRole() *string {
	return s.Role
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRange) GetRoleId() *int32 {
	return s.RoleId
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRange) GetTimeRange() *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeTimeRange {
	return s.TimeRange
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

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRange) Validate() error {
	if s.Anchor != nil {
		if err := s.Anchor.Validate(); err != nil {
			return err
		}
	}
	if s.Range != nil {
		if err := s.Range.Validate(); err != nil {
			return err
		}
	}
	if s.TimeRange != nil {
		if err := s.TimeRange.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeAnchor) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeAnchor) GetCid() *string {
	return s.Cid
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeAnchor) GetHitTime() *int32 {
	return s.HitTime
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeAnchor) GetLocation() *string {
	return s.Location
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

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeAnchor) Validate() error {
	return dara.Validate(s)
}

type GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeRange struct {
	// 对话开始索引
	From *int32 `json:"From,omitempty" xml:"From,omitempty"`
	// 对话结束索引
	To *int32 `json:"To,omitempty" xml:"To,omitempty"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeRange) String() string {
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeRange) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeRange) GetFrom() *int32 {
	return s.From
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeRange) GetTo() *int32 {
	return s.To
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeRange) SetFrom(v int32) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeRange {
	s.From = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeRange) SetTo(v int32) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeRange {
	s.To = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeRange) Validate() error {
	return dara.Validate(s)
}

type GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeTimeRange struct {
	From *int64 `json:"From,omitempty" xml:"From,omitempty"`
	To   *int64 `json:"To,omitempty" xml:"To,omitempty"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeTimeRange) String() string {
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeTimeRange) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeTimeRange) GetFrom() *int64 {
	return s.From
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeTimeRange) GetTo() *int64 {
	return s.To
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeTimeRange) SetFrom(v int64) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeTimeRange {
	s.From = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeTimeRange) SetTo(v int64) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeTimeRange {
	s.To = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsCheckRangeTimeRange) Validate() error {
	return dara.Validate(s)
}

type GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperators struct {
	Operator []*GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperator `json:"Operator,omitempty" xml:"Operator,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperators) String() string {
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperators) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperators) GetOperator() []*GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperator {
	return s.Operator
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperators) SetOperator(v []*GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperator) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperators {
	s.Operator = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperators) Validate() error {
	if s.Operator != nil {
		for _, item := range s.Operator {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
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
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperator) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperator) GetId() *int64 {
	return s.Id
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperator) GetName() *string {
	return s.Name
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperator) GetOid() *string {
	return s.Oid
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperator) GetParam() *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam {
	return s.Param
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperator) GetType() *string {
	return s.Type
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

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperator) Validate() error {
	if s.Param != nil {
		if err := s.Param.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) GetAverage() *bool {
	return s.Average
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) GetBeginType() *string {
	return s.BeginType
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) GetCaseSensitive() *bool {
	return s.CaseSensitive
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) GetCheckFirstSentence() *bool {
	return s.CheckFirstSentence
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) GetCheckType() *int32 {
	return s.CheckType
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) GetCompareOperator() *string {
	return s.CompareOperator
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) GetContextChatMatch() *bool {
	return s.ContextChatMatch
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) GetDelayTime() *int32 {
	return s.DelayTime
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) GetEndType() *string {
	return s.EndType
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) GetExcludes() *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamExcludes {
	return s.Excludes
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) GetFlowNodePrerequisiteParam() *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamFlowNodePrerequisiteParam {
	return s.FlowNodePrerequisiteParam
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) GetFrom() *int32 {
	return s.From
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) GetFromEnd() *bool {
	return s.FromEnd
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) GetHitTime() *int32 {
	return s.HitTime
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) GetInSentence() *bool {
	return s.InSentence
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) GetIntentModelCheckParm() *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParm {
	return s.IntentModelCheckParm
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) GetInterval() *int32 {
	return s.Interval
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) GetIntervalEnd() *int32 {
	return s.IntervalEnd
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) GetKeywordExtension() *int32 {
	return s.KeywordExtension
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) GetKeywordMatchSize() *int32 {
	return s.KeywordMatchSize
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) GetKeywords() *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamKeywords {
	return s.Keywords
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) GetMaxEmotionChangeValue() *int32 {
	return s.MaxEmotionChangeValue
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) GetMinWordSize() *int32 {
	return s.MinWordSize
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) GetNearDialogue() *bool {
	return s.NearDialogue
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) GetNotRegex() *string {
	return s.NotRegex
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) GetPhrase() *string {
	return s.Phrase
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) GetRegex() *string {
	return s.Regex
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) GetTarget() *int32 {
	return s.Target
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) GetThreshold() *float32 {
	return s.Threshold
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

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParam) Validate() error {
	if s.Excludes != nil {
		if err := s.Excludes.Validate(); err != nil {
			return err
		}
	}
	if s.FlowNodePrerequisiteParam != nil {
		if err := s.FlowNodePrerequisiteParam.Validate(); err != nil {
			return err
		}
	}
	if s.IntentModelCheckParm != nil {
		if err := s.IntentModelCheckParm.Validate(); err != nil {
			return err
		}
	}
	if s.Keywords != nil {
		if err := s.Keywords.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamExcludes struct {
	Exclude []*string `json:"Exclude,omitempty" xml:"Exclude,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamExcludes) String() string {
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamExcludes) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamExcludes) GetExclude() []*string {
	return s.Exclude
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamExcludes) SetExclude(v []*string) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamExcludes {
	s.Exclude = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamExcludes) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamFlowNodePrerequisiteParam) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamFlowNodePrerequisiteParam) GetNodeId() *int64 {
	return s.NodeId
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamFlowNodePrerequisiteParam) GetNodeMatchStatus() *int32 {
	return s.NodeMatchStatus
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamFlowNodePrerequisiteParam) GetNodeName() *string {
	return s.NodeName
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

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamFlowNodePrerequisiteParam) Validate() error {
	return dara.Validate(s)
}

type GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParm struct {
	// 引用的意图模型
	Intents *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntents `json:"Intents,omitempty" xml:"Intents,omitempty" type:"Struct"`
	// 模型应用的场景 AGENT:客户场景、CUSTOMER:客服场景 (CUSTOMER: 客户场景, AGENT: 坐席场景)
	ModelScene *string `json:"ModelScene,omitempty" xml:"ModelScene,omitempty"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParm) String() string {
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParm) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParm) GetIntents() *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntents {
	return s.Intents
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParm) GetModelScene() *string {
	return s.ModelScene
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParm) SetIntents(v *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntents) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParm {
	s.Intents = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParm) SetModelScene(v string) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParm {
	s.ModelScene = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParm) Validate() error {
	if s.Intents != nil {
		if err := s.Intents.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntents struct {
	Intent []*GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntentsIntent `json:"Intent,omitempty" xml:"Intent,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntents) String() string {
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntents) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntents) GetIntent() []*GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntentsIntent {
	return s.Intent
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntents) SetIntent(v []*GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntentsIntent) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntents {
	s.Intent = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntents) Validate() error {
	if s.Intent != nil {
		for _, item := range s.Intent {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntentsIntent struct {
	// 意图模型ID
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 意图模型名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntentsIntent) String() string {
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntentsIntent) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntentsIntent) GetId() *int64 {
	return s.Id
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntentsIntent) GetName() *string {
	return s.Name
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntentsIntent) SetId(v int64) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntentsIntent {
	s.Id = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntentsIntent) SetName(v string) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntentsIntent {
	s.Name = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamIntentModelCheckParmIntentsIntent) Validate() error {
	return dara.Validate(s)
}

type GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamKeywords struct {
	Keyword []*string `json:"Keyword,omitempty" xml:"Keyword,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamKeywords) String() string {
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamKeywords) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamKeywords) GetKeyword() []*string {
	return s.Keyword
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamKeywords) SetKeyword(v []*string) *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamKeywords {
	s.Keyword = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultConditionsConditionsOperatorsOperatorParamKeywords) Validate() error {
	return dara.Validate(s)
}

type GetResultResponseBodyDataResultInfoHitResultHitResultHits struct {
	Hit []*GetResultResponseBodyDataResultInfoHitResultHitResultHitsHit `json:"Hit,omitempty" xml:"Hit,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultHits) String() string {
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultHits) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHits) GetHit() []*GetResultResponseBodyDataResultInfoHitResultHitResultHitsHit {
	return s.Hit
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHits) SetHit(v []*GetResultResponseBodyDataResultInfoHitResultHitResultHitsHit) *GetResultResponseBodyDataResultInfoHitResultHitResultHits {
	s.Hit = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHits) Validate() error {
	if s.Hit != nil {
		for _, item := range s.Hit {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetResultResponseBodyDataResultInfoHitResultHitResultHitsHit struct {
	Cid      *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitCid      `json:"Cid,omitempty" xml:"Cid,omitempty" type:"Struct"`
	KeyWords *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWords `json:"KeyWords,omitempty" xml:"KeyWords,omitempty" type:"Struct"`
	Phrase   *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase   `json:"Phrase,omitempty" xml:"Phrase,omitempty" type:"Struct"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultHitsHit) String() string {
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultHitsHit) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHit) GetCid() *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitCid {
	return s.Cid
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHit) GetKeyWords() *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWords {
	return s.KeyWords
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHit) GetPhrase() *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase {
	return s.Phrase
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

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHit) Validate() error {
	if s.Cid != nil {
		if err := s.Cid.Validate(); err != nil {
			return err
		}
	}
	if s.KeyWords != nil {
		if err := s.KeyWords.Validate(); err != nil {
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

type GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitCid struct {
	Cid []*string `json:"Cid,omitempty" xml:"Cid,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitCid) String() string {
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitCid) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitCid) GetCid() []*string {
	return s.Cid
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitCid) SetCid(v []*string) *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitCid {
	s.Cid = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitCid) Validate() error {
	return dara.Validate(s)
}

type GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWords struct {
	KeyWord []*GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWordsKeyWord `json:"KeyWord,omitempty" xml:"KeyWord,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWords) String() string {
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWords) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWords) GetKeyWord() []*GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWordsKeyWord {
	return s.KeyWord
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWords) SetKeyWord(v []*GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWordsKeyWord) *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWords {
	s.KeyWord = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWords) Validate() error {
	if s.KeyWord != nil {
		for _, item := range s.KeyWord {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
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
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWordsKeyWord) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWordsKeyWord) GetCid() *string {
	return s.Cid
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWordsKeyWord) GetFrom() *int32 {
	return s.From
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWordsKeyWord) GetTo() *int32 {
	return s.To
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWordsKeyWord) GetVal() *string {
	return s.Val
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

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWordsKeyWord) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase) GetBegin() *int64 {
	return s.Begin
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase) GetEmotionValue() *int32 {
	return s.EmotionValue
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase) GetEnd() *int32 {
	return s.End
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase) GetRole() *string {
	return s.Role
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase) GetWords() *string {
	return s.Words
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

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase) Validate() error {
	return dara.Validate(s)
}

type GetResultResponseBodyDataResultInfoHitScore struct {
	HitScore []*GetResultResponseBodyDataResultInfoHitScoreHitScore `json:"HitScore,omitempty" xml:"HitScore,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoHitScore) String() string {
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitScore) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitScore) GetHitScore() []*GetResultResponseBodyDataResultInfoHitScoreHitScore {
	return s.HitScore
}

func (s *GetResultResponseBodyDataResultInfoHitScore) SetHitScore(v []*GetResultResponseBodyDataResultInfoHitScoreHitScore) *GetResultResponseBodyDataResultInfoHitScore {
	s.HitScore = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitScore) Validate() error {
	if s.HitScore != nil {
		for _, item := range s.HitScore {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
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
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitScoreHitScore) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitScoreHitScore) GetRuleId() *string {
	return s.RuleId
}

func (s *GetResultResponseBodyDataResultInfoHitScoreHitScore) GetScoreId() *string {
	return s.ScoreId
}

func (s *GetResultResponseBodyDataResultInfoHitScoreHitScore) GetScoreName() *string {
	return s.ScoreName
}

func (s *GetResultResponseBodyDataResultInfoHitScoreHitScore) GetScoreNumber() *string {
	return s.ScoreNumber
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

func (s *GetResultResponseBodyDataResultInfoHitScoreHitScore) Validate() error {
	return dara.Validate(s)
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
	Caller       *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	CustomerName *string `json:"CustomerName,omitempty" xml:"CustomerName,omitempty"`
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
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoRecording) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoRecording) GetBusiness() *string {
	return s.Business
}

func (s *GetResultResponseBodyDataResultInfoRecording) GetCallId() *string {
	return s.CallId
}

func (s *GetResultResponseBodyDataResultInfoRecording) GetCallTime() *string {
	return s.CallTime
}

func (s *GetResultResponseBodyDataResultInfoRecording) GetCallType() *int32 {
	return s.CallType
}

func (s *GetResultResponseBodyDataResultInfoRecording) GetCallee() *string {
	return s.Callee
}

func (s *GetResultResponseBodyDataResultInfoRecording) GetCaller() *string {
	return s.Caller
}

func (s *GetResultResponseBodyDataResultInfoRecording) GetCustomerName() *string {
	return s.CustomerName
}

func (s *GetResultResponseBodyDataResultInfoRecording) GetDataSetName() *string {
	return s.DataSetName
}

func (s *GetResultResponseBodyDataResultInfoRecording) GetDialogueSize() *int32 {
	return s.DialogueSize
}

func (s *GetResultResponseBodyDataResultInfoRecording) GetDuration() *int64 {
	return s.Duration
}

func (s *GetResultResponseBodyDataResultInfoRecording) GetId() *string {
	return s.Id
}

func (s *GetResultResponseBodyDataResultInfoRecording) GetName() *string {
	return s.Name
}

func (s *GetResultResponseBodyDataResultInfoRecording) GetPrimaryId() *string {
	return s.PrimaryId
}

func (s *GetResultResponseBodyDataResultInfoRecording) GetRemark1() *string {
	return s.Remark1
}

func (s *GetResultResponseBodyDataResultInfoRecording) GetRemark10() *string {
	return s.Remark10
}

func (s *GetResultResponseBodyDataResultInfoRecording) GetRemark11() *string {
	return s.Remark11
}

func (s *GetResultResponseBodyDataResultInfoRecording) GetRemark12() *string {
	return s.Remark12
}

func (s *GetResultResponseBodyDataResultInfoRecording) GetRemark13() *string {
	return s.Remark13
}

func (s *GetResultResponseBodyDataResultInfoRecording) GetRemark2() *string {
	return s.Remark2
}

func (s *GetResultResponseBodyDataResultInfoRecording) GetRemark3() *string {
	return s.Remark3
}

func (s *GetResultResponseBodyDataResultInfoRecording) GetRemark4() *string {
	return s.Remark4
}

func (s *GetResultResponseBodyDataResultInfoRecording) GetRemark5() *int64 {
	return s.Remark5
}

func (s *GetResultResponseBodyDataResultInfoRecording) GetRemark6() *string {
	return s.Remark6
}

func (s *GetResultResponseBodyDataResultInfoRecording) GetRemark7() *string {
	return s.Remark7
}

func (s *GetResultResponseBodyDataResultInfoRecording) GetRemark8() *string {
	return s.Remark8
}

func (s *GetResultResponseBodyDataResultInfoRecording) GetRemark9() *string {
	return s.Remark9
}

func (s *GetResultResponseBodyDataResultInfoRecording) GetTaskConfigId() *int64 {
	return s.TaskConfigId
}

func (s *GetResultResponseBodyDataResultInfoRecording) GetTaskConfigName() *string {
	return s.TaskConfigName
}

func (s *GetResultResponseBodyDataResultInfoRecording) GetUrl() *string {
	return s.Url
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

func (s *GetResultResponseBodyDataResultInfoRecording) SetCustomerName(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.CustomerName = &v
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

func (s *GetResultResponseBodyDataResultInfoRecording) Validate() error {
	return dara.Validate(s)
}

type GetResultResponseBodyDataResultInfoReviewHistoryList struct {
	ReviewHistory []*GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory `json:"ReviewHistory,omitempty" xml:"ReviewHistory,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoReviewHistoryList) String() string {
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoReviewHistoryList) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryList) GetReviewHistory() []*GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory {
	return s.ReviewHistory
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryList) SetReviewHistory(v []*GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory) *GetResultResponseBodyDataResultInfoReviewHistoryList {
	s.ReviewHistory = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryList) Validate() error {
	if s.ReviewHistory != nil {
		for _, item := range s.ReviewHistory {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
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
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory) GetComments() *string {
	return s.Comments
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory) GetComplainResult() *int32 {
	return s.ComplainResult
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory) GetOldScore() *int32 {
	return s.OldScore
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory) GetOperator() *int64 {
	return s.Operator
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory) GetOperatorName() *string {
	return s.OperatorName
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory) GetReviewManagerType() *string {
	return s.ReviewManagerType
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory) GetReviewResult() *int32 {
	return s.ReviewResult
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory) GetReviewRightRule() *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRule {
	return s.ReviewRightRule
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory) GetScore() *int32 {
	return s.Score
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory) GetTime() *int64 {
	return s.Time
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory) GetTimeStr() *string {
	return s.TimeStr
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory) GetType() *int32 {
	return s.Type
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

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistory) Validate() error {
	if s.ReviewRightRule != nil {
		if err := s.ReviewRightRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRule struct {
	ReviewRightRule []*GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule `json:"ReviewRightRule,omitempty" xml:"ReviewRightRule,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRule) String() string {
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRule) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRule) GetReviewRightRule() []*GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule {
	return s.ReviewRightRule
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRule) SetReviewRightRule(v []*GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule) *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRule {
	s.ReviewRightRule = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRule) Validate() error {
	if s.ReviewRightRule != nil {
		for _, item := range s.ReviewRightRule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule struct {
	Rid      *int64  `json:"rid,omitempty" xml:"rid,omitempty"`
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
}

func (s GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule) String() string {
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule) GetRid() *int64 {
	return s.Rid
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule) GetRuleName() *string {
	return s.RuleName
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule) SetRid(v int64) *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule {
	s.Rid = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule) SetRuleName(v string) *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule {
	s.RuleName = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoReviewHistoryListReviewHistoryReviewRightRuleReviewRightRule) Validate() error {
	return dara.Validate(s)
}

type GetResultResponseBodyDataResultInfoReviewTypeIdList struct {
	ReviewTypeIdList []*GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdList `json:"ReviewTypeIdList,omitempty" xml:"ReviewTypeIdList,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoReviewTypeIdList) String() string {
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoReviewTypeIdList) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoReviewTypeIdList) GetReviewTypeIdList() []*GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdList {
	return s.ReviewTypeIdList
}

func (s *GetResultResponseBodyDataResultInfoReviewTypeIdList) SetReviewTypeIdList(v []*GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdList) *GetResultResponseBodyDataResultInfoReviewTypeIdList {
	s.ReviewTypeIdList = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoReviewTypeIdList) Validate() error {
	if s.ReviewTypeIdList != nil {
		for _, item := range s.ReviewTypeIdList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdList struct {
	ReviewKeyIdList *GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdListReviewKeyIdList `json:"ReviewKeyIdList,omitempty" xml:"ReviewKeyIdList,omitempty" type:"Struct"`
	ReviewTypeId    *int64                                                                              `json:"ReviewTypeId,omitempty" xml:"ReviewTypeId,omitempty"`
}

func (s GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdList) String() string {
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdList) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdList) GetReviewKeyIdList() *GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdListReviewKeyIdList {
	return s.ReviewKeyIdList
}

func (s *GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdList) GetReviewTypeId() *int64 {
	return s.ReviewTypeId
}

func (s *GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdList) SetReviewKeyIdList(v *GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdListReviewKeyIdList) *GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdList {
	s.ReviewKeyIdList = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdList) SetReviewTypeId(v int64) *GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdList {
	s.ReviewTypeId = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdList) Validate() error {
	if s.ReviewKeyIdList != nil {
		if err := s.ReviewKeyIdList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdListReviewKeyIdList struct {
	ReviewKeyIdList []*int64 `json:"ReviewKeyIdList,omitempty" xml:"ReviewKeyIdList,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdListReviewKeyIdList) String() string {
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdListReviewKeyIdList) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdListReviewKeyIdList) GetReviewKeyIdList() []*int64 {
	return s.ReviewKeyIdList
}

func (s *GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdListReviewKeyIdList) SetReviewKeyIdList(v []*int64) *GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdListReviewKeyIdList {
	s.ReviewKeyIdList = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoReviewTypeIdListReviewTypeIdListReviewKeyIdList) Validate() error {
	return dara.Validate(s)
}

type GetResultResponseBodyDataResultInfoSchemeIdList struct {
	SchemeIdList []*int64 `json:"SchemeIdList,omitempty" xml:"SchemeIdList,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoSchemeIdList) String() string {
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoSchemeIdList) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoSchemeIdList) GetSchemeIdList() []*int64 {
	return s.SchemeIdList
}

func (s *GetResultResponseBodyDataResultInfoSchemeIdList) SetSchemeIdList(v []*int64) *GetResultResponseBodyDataResultInfoSchemeIdList {
	s.SchemeIdList = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoSchemeIdList) Validate() error {
	return dara.Validate(s)
}

type GetResultResponseBodyDataResultInfoSchemeNameList struct {
	SchemeNameList []*string `json:"SchemeNameList,omitempty" xml:"SchemeNameList,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoSchemeNameList) String() string {
	return dara.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoSchemeNameList) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoSchemeNameList) GetSchemeNameList() []*string {
	return s.SchemeNameList
}

func (s *GetResultResponseBodyDataResultInfoSchemeNameList) SetSchemeNameList(v []*string) *GetResultResponseBodyDataResultInfoSchemeNameList {
	s.SchemeNameList = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoSchemeNameList) Validate() error {
	return dara.Validate(s)
}
