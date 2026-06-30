// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSyncResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSyncResultResponseBody
	GetCode() *string
	SetCount(v int32) *GetSyncResultResponseBody
	GetCount() *int32
	SetData(v []*GetSyncResultResponseBodyData) *GetSyncResultResponseBody
	GetData() []*GetSyncResultResponseBodyData
	SetMessage(v string) *GetSyncResultResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *GetSyncResultResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *GetSyncResultResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetSyncResultResponseBody
	GetRequestId() *string
	SetResultCountId(v string) *GetSyncResultResponseBody
	GetResultCountId() *string
	SetSuccess(v bool) *GetSyncResultResponseBody
	GetSuccess() *bool
}

type GetSyncResultResponseBody struct {
	// Result code. A value of 200 indicates success. Any other value indicates failure. The caller can use this field to determine the cause of failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Total number of entries.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// Query result.
	Data []*GetSyncResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Error details when an error occurs; "successful" when the operation succeeded.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Page number
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID
	//
	// example:
	//
	// 76DB5D8C-5BD9-42A7-B527-5AF3A5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Internal field. Ignore it.
	//
	// example:
	//
	// xxx
	ResultCountId *string `json:"ResultCountId,omitempty" xml:"ResultCountId,omitempty"`
	// Indicates whether the request succeeded. The caller can use this field to determine the request status: true indicates success; false or null indicates failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSyncResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSyncResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetSyncResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSyncResultResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *GetSyncResultResponseBody) GetData() []*GetSyncResultResponseBodyData {
	return s.Data
}

func (s *GetSyncResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSyncResultResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetSyncResultResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSyncResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSyncResultResponseBody) GetResultCountId() *string {
	return s.ResultCountId
}

func (s *GetSyncResultResponseBody) GetSuccess() *bool {
	return s.Success
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

func (s *GetSyncResultResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSyncResultResponseBodyData struct {
	// Agent information
	Agent *GetSyncResultResponseBodyDataAgent `json:"Agent,omitempty" xml:"Agent,omitempty" type:"Struct"`
	// Transcription result (dialogue text)
	AsrResult []*GetSyncResultResponseBodyDataAsrResult `json:"AsrResult,omitempty" xml:"AsrResult,omitempty" type:"Repeated"`
	// Review comments.
	//
	// example:
	//
	// xxx
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// Job Creation Time.
	//
	// example:
	//
	// 2019-07-24T19:31Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// When status is neither 0 nor 1, this field indicates the Error Details.
	//
	// example:
	//
	// xxxx
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Rule hit result.
	HitResult []*GetSyncResultResponseBodyDataHitResult `json:"HitResult,omitempty" xml:"HitResult,omitempty" type:"Repeated"`
	// Recording file information
	Recording *GetSyncResultResponseBodyDataRecording `json:"Recording,omitempty" xml:"Recording,omitempty" type:"Struct"`
	// The quality inspector who actually reviewed the task.
	//
	// example:
	//
	// 张三
	Resolver *string `json:"Resolver,omitempty" xml:"Resolver,omitempty"`
	// Review accuracy. Possible values: 0 (fault); 1 (correct); 2 (partially correct); 3 (pending review).
	//
	// example:
	//
	// 3
	ReviewResult *int32 `json:"ReviewResult,omitempty" xml:"ReviewResult,omitempty"`
	// Review status; possible values: 0 (not reviewed); 1 (reviewed).
	//
	// example:
	//
	// 1
	ReviewStatus *int32 `json:"ReviewStatus,omitempty" xml:"ReviewStatus,omitempty"`
	// Username of the assigned quality inspector.
	//
	// example:
	//
	// 张三
	Reviewer *string `json:"Reviewer,omitempty" xml:"Reviewer,omitempty"`
	// Quality inspection score, with a maximum of 100.
	//
	// example:
	//
	// 100
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// Current job status. Possible values: 0 (not completed); 1 (completed). The caller can use this field to determine whether the job is complete. Values other than 0 or 1 indicate an error; see the errorMessage field for Error Details.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Job ID.
	//
	// example:
	//
	// 20201231de3d34ec-40fa-4a55-8d27-76ea*****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Internal field. Ignore it.
	//
	// example:
	//
	// xxx
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s GetSyncResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSyncResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSyncResultResponseBodyData) GetAgent() *GetSyncResultResponseBodyDataAgent {
	return s.Agent
}

func (s *GetSyncResultResponseBodyData) GetAsrResult() []*GetSyncResultResponseBodyDataAsrResult {
	return s.AsrResult
}

func (s *GetSyncResultResponseBodyData) GetComments() *string {
	return s.Comments
}

func (s *GetSyncResultResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetSyncResultResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetSyncResultResponseBodyData) GetHitResult() []*GetSyncResultResponseBodyDataHitResult {
	return s.HitResult
}

func (s *GetSyncResultResponseBodyData) GetRecording() *GetSyncResultResponseBodyDataRecording {
	return s.Recording
}

func (s *GetSyncResultResponseBodyData) GetResolver() *string {
	return s.Resolver
}

func (s *GetSyncResultResponseBodyData) GetReviewResult() *int32 {
	return s.ReviewResult
}

func (s *GetSyncResultResponseBodyData) GetReviewStatus() *int32 {
	return s.ReviewStatus
}

func (s *GetSyncResultResponseBodyData) GetReviewer() *string {
	return s.Reviewer
}

func (s *GetSyncResultResponseBodyData) GetScore() *int32 {
	return s.Score
}

func (s *GetSyncResultResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *GetSyncResultResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetSyncResultResponseBodyData) GetTaskName() *string {
	return s.TaskName
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

func (s *GetSyncResultResponseBodyData) Validate() error {
	if s.Agent != nil {
		if err := s.Agent.Validate(); err != nil {
			return err
		}
	}
	if s.AsrResult != nil {
		for _, item := range s.AsrResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.HitResult != nil {
		for _, item := range s.HitResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Recording != nil {
		if err := s.Recording.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSyncResultResponseBodyDataAgent struct {
	// Agent ID.
	//
	// example:
	//
	// 12221
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Agent name
	//
	// example:
	//
	// 李四
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Skill group name
	//
	// example:
	//
	// 客服组
	SkillGroup *string `json:"SkillGroup,omitempty" xml:"SkillGroup,omitempty"`
}

func (s GetSyncResultResponseBodyDataAgent) String() string {
	return dara.Prettify(s)
}

func (s GetSyncResultResponseBodyDataAgent) GoString() string {
	return s.String()
}

func (s *GetSyncResultResponseBodyDataAgent) GetId() *string {
	return s.Id
}

func (s *GetSyncResultResponseBodyDataAgent) GetName() *string {
	return s.Name
}

func (s *GetSyncResultResponseBodyDataAgent) GetSkillGroup() *string {
	return s.SkillGroup
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

func (s *GetSyncResultResponseBodyDataAgent) Validate() error {
	return dara.Validate(s)
}

type GetSyncResultResponseBodyDataAsrResult struct {
	// The start time of this sentence, which is the offset from the starting point in milliseconds.
	//
	// example:
	//
	// 340
	Begin *int64 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// Emotion intensity value ranging from 1 to 10. A higher value indicates stronger emotion.
	//
	// example:
	//
	// 6
	EmotionValue *int32 `json:"EmotionValue,omitempty" xml:"EmotionValue,omitempty"`
	// The end time of this sentence, which is the offset from the starting point in milliseconds.
	//
	// example:
	//
	// 3000
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// Role in the dialogue content. Possible values: agent, Customer.
	//
	// example:
	//
	// 客服
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// Internal field. Ignore it.
	//
	// example:
	//
	// 11
	SilenceDuration *int32 `json:"SilenceDuration,omitempty" xml:"SilenceDuration,omitempty"`
	// The average speech rate of this sentence, in characters per minute.
	//
	// example:
	//
	// 221
	SpeechRate *int32 `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	// Dialogue content.
	//
	// example:
	//
	// 您好，很高兴为您服务
	Words *string `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s GetSyncResultResponseBodyDataAsrResult) String() string {
	return dara.Prettify(s)
}

func (s GetSyncResultResponseBodyDataAsrResult) GoString() string {
	return s.String()
}

func (s *GetSyncResultResponseBodyDataAsrResult) GetBegin() *int64 {
	return s.Begin
}

func (s *GetSyncResultResponseBodyDataAsrResult) GetEmotionValue() *int32 {
	return s.EmotionValue
}

func (s *GetSyncResultResponseBodyDataAsrResult) GetEnd() *int64 {
	return s.End
}

func (s *GetSyncResultResponseBodyDataAsrResult) GetRole() *string {
	return s.Role
}

func (s *GetSyncResultResponseBodyDataAsrResult) GetSilenceDuration() *int32 {
	return s.SilenceDuration
}

func (s *GetSyncResultResponseBodyDataAsrResult) GetSpeechRate() *int32 {
	return s.SpeechRate
}

func (s *GetSyncResultResponseBodyDataAsrResult) GetWords() *string {
	return s.Words
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

func (s *GetSyncResultResponseBodyDataAsrResult) Validate() error {
	return dara.Validate(s)
}

type GetSyncResultResponseBodyDataHitResult struct {
	// Specific hit location information. At the sentence dimension, returns which condition in the rule was hit and which specific characters triggered the hit within the sentence.
	Hits []*GetSyncResultResponseBodyDataHitResultHits `json:"Hits,omitempty" xml:"Hits,omitempty" type:"Repeated"`
	// Hit rule name.
	//
	// example:
	//
	// 测试规则
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Review accuracy; possible values: 0 (fault); 1 (correct).
	//
	// example:
	//
	// 1
	ReviewResult *int32 `json:"ReviewResult,omitempty" xml:"ReviewResult,omitempty"`
	// Hit rule ID.
	//
	// example:
	//
	// 1211
	Rid *string `json:"Rid,omitempty" xml:"Rid,omitempty"`
	// Rule type associated with the hit rule.
	//
	// example:
	//
	// 2
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetSyncResultResponseBodyDataHitResult) String() string {
	return dara.Prettify(s)
}

func (s GetSyncResultResponseBodyDataHitResult) GoString() string {
	return s.String()
}

func (s *GetSyncResultResponseBodyDataHitResult) GetHits() []*GetSyncResultResponseBodyDataHitResultHits {
	return s.Hits
}

func (s *GetSyncResultResponseBodyDataHitResult) GetName() *string {
	return s.Name
}

func (s *GetSyncResultResponseBodyDataHitResult) GetReviewResult() *int32 {
	return s.ReviewResult
}

func (s *GetSyncResultResponseBodyDataHitResult) GetRid() *string {
	return s.Rid
}

func (s *GetSyncResultResponseBodyDataHitResult) GetType() *string {
	return s.Type
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

func (s *GetSyncResultResponseBodyDataHitResult) Validate() error {
	if s.Hits != nil {
		for _, item := range s.Hits {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSyncResultResponseBodyDataHitResultHits struct {
	// List of hit condition IDs.
	Cid []*string `json:"Cid,omitempty" xml:"Cid,omitempty" type:"Repeated"`
	// Returns the specific characters in the current sentence that hit the rule, which are the keywords to be highlighted.
	KeyWords []*GetSyncResultResponseBodyDataHitResultHitsKeyWords `json:"KeyWords,omitempty" xml:"KeyWords,omitempty" type:"Repeated"`
	// Details of the sentence that hit the current rule.
	Phrase *GetSyncResultResponseBodyDataHitResultHitsPhrase `json:"Phrase,omitempty" xml:"Phrase,omitempty" type:"Struct"`
}

func (s GetSyncResultResponseBodyDataHitResultHits) String() string {
	return dara.Prettify(s)
}

func (s GetSyncResultResponseBodyDataHitResultHits) GoString() string {
	return s.String()
}

func (s *GetSyncResultResponseBodyDataHitResultHits) GetCid() []*string {
	return s.Cid
}

func (s *GetSyncResultResponseBodyDataHitResultHits) GetKeyWords() []*GetSyncResultResponseBodyDataHitResultHitsKeyWords {
	return s.KeyWords
}

func (s *GetSyncResultResponseBodyDataHitResultHits) GetPhrase() *GetSyncResultResponseBodyDataHitResultHitsPhrase {
	return s.Phrase
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

func (s *GetSyncResultResponseBodyDataHitResultHits) Validate() error {
	if s.KeyWords != nil {
		for _, item := range s.KeyWords {
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

type GetSyncResultResponseBodyDataHitResultHitsKeyWords struct {
	// The ID of the condition that was hit.
	//
	// example:
	//
	// 66666
	Cid *string `json:"Cid,omitempty" xml:"Cid,omitempty"`
	// The starting character position (inclusive) of the keyword to be highlighted. The value starts from 0 and can be at most the total number of characters in the sentence minus 1.
	//
	// example:
	//
	// 2
	From *int32 `json:"From,omitempty" xml:"From,omitempty"`
	// The ending character position (exclusive) of the keyword to be highlighted. The maximum value is the total number of characters in the sentence minus 1. For example, in the sentence “不可能给你退货的”, if from=0 and to=3, the highlighted keyword is “不可能”, which consists of three characters.
	//
	// example:
	//
	// 5
	To *int32 `json:"To,omitempty" xml:"To,omitempty"`
	// The exact keyword content.
	//
	// example:
	//
	// 投诉
	Val *string `json:"Val,omitempty" xml:"Val,omitempty"`
}

func (s GetSyncResultResponseBodyDataHitResultHitsKeyWords) String() string {
	return dara.Prettify(s)
}

func (s GetSyncResultResponseBodyDataHitResultHitsKeyWords) GoString() string {
	return s.String()
}

func (s *GetSyncResultResponseBodyDataHitResultHitsKeyWords) GetCid() *string {
	return s.Cid
}

func (s *GetSyncResultResponseBodyDataHitResultHitsKeyWords) GetFrom() *int32 {
	return s.From
}

func (s *GetSyncResultResponseBodyDataHitResultHitsKeyWords) GetTo() *int32 {
	return s.To
}

func (s *GetSyncResultResponseBodyDataHitResultHitsKeyWords) GetVal() *string {
	return s.Val
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

func (s *GetSyncResultResponseBodyDataHitResultHitsKeyWords) Validate() error {
	return dara.Validate(s)
}

type GetSyncResultResponseBodyDataHitResultHitsPhrase struct {
	// The Start Time of this sentence, represented as an offset in milliseconds from the starting point.
	//
	// example:
	//
	// 440
	Begin *int64 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// Emotion intensity value ranging from 1 to 10. A higher value indicates stronger emotion.
	//
	// example:
	//
	// 6
	EmotionValue *int32 `json:"EmotionValue,omitempty" xml:"EmotionValue,omitempty"`
	// The End Time of this sentence, represented as an offset in milliseconds from the starting point.
	//
	// example:
	//
	// 4000
	End *int32 `json:"End,omitempty" xml:"End,omitempty"`
	// The role in the conversation content. Possible values: agent, Customer, System.
	//
	// example:
	//
	// 客服
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// Internal field. Ignore.
	//
	// example:
	//
	// 1
	SilenceDuration *int32 `json:"SilenceDuration,omitempty" xml:"SilenceDuration,omitempty"`
	// The speech rate of this sentence.
	//
	// example:
	//
	// 234
	SpeechRate *int32 `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	// A sentence spoken by this role.
	//
	// example:
	//
	// 我要投诉
	Words *string `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s GetSyncResultResponseBodyDataHitResultHitsPhrase) String() string {
	return dara.Prettify(s)
}

func (s GetSyncResultResponseBodyDataHitResultHitsPhrase) GoString() string {
	return s.String()
}

func (s *GetSyncResultResponseBodyDataHitResultHitsPhrase) GetBegin() *int64 {
	return s.Begin
}

func (s *GetSyncResultResponseBodyDataHitResultHitsPhrase) GetEmotionValue() *int32 {
	return s.EmotionValue
}

func (s *GetSyncResultResponseBodyDataHitResultHitsPhrase) GetEnd() *int32 {
	return s.End
}

func (s *GetSyncResultResponseBodyDataHitResultHitsPhrase) GetRole() *string {
	return s.Role
}

func (s *GetSyncResultResponseBodyDataHitResultHitsPhrase) GetSilenceDuration() *int32 {
	return s.SilenceDuration
}

func (s *GetSyncResultResponseBodyDataHitResultHitsPhrase) GetSpeechRate() *int32 {
	return s.SpeechRate
}

func (s *GetSyncResultResponseBodyDataHitResultHitsPhrase) GetWords() *string {
	return s.Words
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

func (s *GetSyncResultResponseBodyDataHitResultHitsPhrase) Validate() error {
	return dara.Validate(s)
}

type GetSyncResultResponseBodyDataRecording struct {
	// Line-of-business name.
	//
	// example:
	//
	// 客服部
	Business *string `json:"Business,omitempty" xml:"Business,omitempty"`
	// Call ID.
	//
	// example:
	//
	// xxx
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// Recording generation UNIX timestamp, accurate to milliseconds.
	//
	// example:
	//
	// 1563967699000
	CallTime *string `json:"CallTime,omitempty" xml:"CallTime,omitempty"`
	// Call type:
	//
	// - 1: Outgoing call
	//
	// - 3: Incoming call
	//
	// example:
	//
	// 1
	CallType *int32 `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// Callee number.
	//
	// example:
	//
	// 1888888***
	Callee *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	// Caller number.
	//
	// example:
	//
	// 0108888****
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// Internal field. Ignore this.
	//
	// example:
	//
	// xxx
	DataSetName *string `json:"DataSetName,omitempty" xml:"DataSetName,omitempty"`
	// Total number of words in the conversation.
	//
	// example:
	//
	// 232
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Call duration.
	//
	// example:
	//
	// 120
	DurationAudio *int64 `json:"DurationAudio,omitempty" xml:"DurationAudio,omitempty"`
	// File ID, which is the callId in the request parameters. If not specified, a random ID will be generated.
	//
	// example:
	//
	// xxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Recording file name.
	//
	// example:
	//
	// 123123.wav
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Internal field. Ignore it.
	//
	// example:
	//
	// xxxx
	PrimaryId *string `json:"PrimaryId,omitempty" xml:"PrimaryId,omitempty"`
	// Custom data 1.
	//
	// example:
	//
	// xxx
	Remark1 *string `json:"Remark1,omitempty" xml:"Remark1,omitempty"`
	// Custom data 2.
	//
	// example:
	//
	// xxx
	Remark2 *string `json:"Remark2,omitempty" xml:"Remark2,omitempty"`
	// Custom data 3.
	//
	// example:
	//
	// xxx
	Remark3 *string `json:"Remark3,omitempty" xml:"Remark3,omitempty"`
	// Recording file URL, used for playback.
	//
	// example:
	//
	// http://aliyun.com/xxx.wav
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetSyncResultResponseBodyDataRecording) String() string {
	return dara.Prettify(s)
}

func (s GetSyncResultResponseBodyDataRecording) GoString() string {
	return s.String()
}

func (s *GetSyncResultResponseBodyDataRecording) GetBusiness() *string {
	return s.Business
}

func (s *GetSyncResultResponseBodyDataRecording) GetCallId() *string {
	return s.CallId
}

func (s *GetSyncResultResponseBodyDataRecording) GetCallTime() *string {
	return s.CallTime
}

func (s *GetSyncResultResponseBodyDataRecording) GetCallType() *int32 {
	return s.CallType
}

func (s *GetSyncResultResponseBodyDataRecording) GetCallee() *string {
	return s.Callee
}

func (s *GetSyncResultResponseBodyDataRecording) GetCaller() *string {
	return s.Caller
}

func (s *GetSyncResultResponseBodyDataRecording) GetDataSetName() *string {
	return s.DataSetName
}

func (s *GetSyncResultResponseBodyDataRecording) GetDuration() *int64 {
	return s.Duration
}

func (s *GetSyncResultResponseBodyDataRecording) GetDurationAudio() *int64 {
	return s.DurationAudio
}

func (s *GetSyncResultResponseBodyDataRecording) GetId() *string {
	return s.Id
}

func (s *GetSyncResultResponseBodyDataRecording) GetName() *string {
	return s.Name
}

func (s *GetSyncResultResponseBodyDataRecording) GetPrimaryId() *string {
	return s.PrimaryId
}

func (s *GetSyncResultResponseBodyDataRecording) GetRemark1() *string {
	return s.Remark1
}

func (s *GetSyncResultResponseBodyDataRecording) GetRemark2() *string {
	return s.Remark2
}

func (s *GetSyncResultResponseBodyDataRecording) GetRemark3() *string {
	return s.Remark3
}

func (s *GetSyncResultResponseBodyDataRecording) GetUrl() *string {
	return s.Url
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

func (s *GetSyncResultResponseBodyDataRecording) Validate() error {
	return dara.Validate(s)
}
