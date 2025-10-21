// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelInputContentDetectResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetModelInputContentDetectResultResponseBody
	GetCode() *string
	SetDetectResultList(v []*GetModelInputContentDetectResultResponseBodyDetectResultList) *GetModelInputContentDetectResultResponseBody
	GetDetectResultList() []*GetModelInputContentDetectResultResponseBodyDetectResultList
	SetHttpStatusCode(v int32) *GetModelInputContentDetectResultResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetModelInputContentDetectResultResponseBody
	GetMessage() *string
	SetProcessedCount(v int32) *GetModelInputContentDetectResultResponseBody
	GetProcessedCount() *int32
	SetRequestId(v string) *GetModelInputContentDetectResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetModelInputContentDetectResultResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *GetModelInputContentDetectResultResponseBody
	GetTaskId() *string
	SetTaskStatus(v int32) *GetModelInputContentDetectResultResponseBody
	GetTaskStatus() *int32
	SetTotalCount(v int32) *GetModelInputContentDetectResultResponseBody
	GetTotalCount() *int32
}

type GetModelInputContentDetectResultResponseBody struct {
	// Status code, 00000 indicates success; others indicate failure.
	//
	// example:
	//
	// 00000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Detection result object
	DetectResultList []*GetModelInputContentDetectResultResponseBodyDetectResultList `json:"DetectResultList,omitempty" xml:"DetectResultList,omitempty" type:"Repeated"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Error message.
	//
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Number of processed items in the task.
	//
	// example:
	//
	// 1
	ProcessedCount *int32 `json:"ProcessedCount,omitempty" xml:"ProcessedCount,omitempty"`
	// Request ID
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. true means success, false means failure.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Task ID.
	//
	// example:
	//
	// 5d85cd38-03b2-49fd-86b2-be85c4b13215
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Task processing status:
	//
	// 0: Queued
	//
	// 1: Processing
	//
	// 2: Completed
	//
	// 3: Failed
	//
	// example:
	//
	// 2
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// Total number of items
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetModelInputContentDetectResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetModelInputContentDetectResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetModelInputContentDetectResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetModelInputContentDetectResultResponseBody) GetDetectResultList() []*GetModelInputContentDetectResultResponseBodyDetectResultList {
	return s.DetectResultList
}

func (s *GetModelInputContentDetectResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetModelInputContentDetectResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetModelInputContentDetectResultResponseBody) GetProcessedCount() *int32 {
	return s.ProcessedCount
}

func (s *GetModelInputContentDetectResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetModelInputContentDetectResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetModelInputContentDetectResultResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *GetModelInputContentDetectResultResponseBody) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *GetModelInputContentDetectResultResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetModelInputContentDetectResultResponseBody) SetCode(v string) *GetModelInputContentDetectResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBody) SetDetectResultList(v []*GetModelInputContentDetectResultResponseBodyDetectResultList) *GetModelInputContentDetectResultResponseBody {
	s.DetectResultList = v
	return s
}

func (s *GetModelInputContentDetectResultResponseBody) SetHttpStatusCode(v int32) *GetModelInputContentDetectResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBody) SetMessage(v string) *GetModelInputContentDetectResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBody) SetProcessedCount(v int32) *GetModelInputContentDetectResultResponseBody {
	s.ProcessedCount = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBody) SetRequestId(v string) *GetModelInputContentDetectResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBody) SetSuccess(v bool) *GetModelInputContentDetectResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBody) SetTaskId(v string) *GetModelInputContentDetectResultResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBody) SetTaskStatus(v int32) *GetModelInputContentDetectResultResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBody) SetTotalCount(v int32) *GetModelInputContentDetectResultResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBody) Validate() error {
	if s.DetectResultList != nil {
		for _, item := range s.DetectResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetModelInputContentDetectResultResponseBodyDetectResultList struct {
	// 0: No risk
	//
	// 1: Risk exists
	//
	// example:
	//
	// 0
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
	// 0: Queued
	//
	// 1: Processing
	//
	// 2: Completed
	//
	// 3: Failed
	//
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Inspection result
	TraceInfo *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfo `json:"TraceInfo,omitempty" xml:"TraceInfo,omitempty" type:"Struct"`
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultList) String() string {
	return dara.Prettify(s)
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultList) GoString() string {
	return s.String()
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultList) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultList) GetStatus() *int32 {
	return s.Status
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultList) GetTraceInfo() *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfo {
	return s.TraceInfo
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultList) SetRiskResult(v int32) *GetModelInputContentDetectResultResponseBodyDetectResultList {
	s.RiskResult = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultList) SetStatus(v int32) *GetModelInputContentDetectResultResponseBodyDetectResultList {
	s.Status = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultList) SetTraceInfo(v *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfo) *GetModelInputContentDetectResultResponseBodyDetectResultList {
	s.TraceInfo = v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultList) Validate() error {
	if s.TraceInfo != nil {
		if err := s.TraceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfo struct {
	// Detected keywords
	BlockWord *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord `json:"BlockWord,omitempty" xml:"BlockWord,omitempty" type:"Struct"`
	// Sensitive topic object list
	DenyTopics *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopics `json:"DenyTopics,omitempty" xml:"DenyTopics,omitempty" type:"Struct"`
	// List of harmful category result objects
	HarmfulCategories *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories `json:"HarmfulCategories,omitempty" xml:"HarmfulCategories,omitempty" type:"Struct"`
	// Prompt attack information
	PromptAttack *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack `json:"PromptAttack,omitempty" xml:"PromptAttack,omitempty" type:"Struct"`
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfo) GoString() string {
	return s.String()
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfo) GetBlockWord() *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord {
	return s.BlockWord
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfo) GetDenyTopics() *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopics {
	return s.DenyTopics
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfo) GetHarmfulCategories() *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories {
	return s.HarmfulCategories
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfo) GetPromptAttack() *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack {
	return s.PromptAttack
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfo) SetBlockWord(v *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfo {
	s.BlockWord = v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfo) SetDenyTopics(v *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopics) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfo {
	s.DenyTopics = v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfo) SetHarmfulCategories(v *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfo {
	s.HarmfulCategories = v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfo) SetPromptAttack(v *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfo {
	s.PromptAttack = v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfo) Validate() error {
	if s.BlockWord != nil {
		if err := s.BlockWord.Validate(); err != nil {
			return err
		}
	}
	if s.DenyTopics != nil {
		if err := s.DenyTopics.Validate(); err != nil {
			return err
		}
	}
	if s.HarmfulCategories != nil {
		if err := s.HarmfulCategories.Validate(); err != nil {
			return err
		}
	}
	if s.PromptAttack != nil {
		if err := s.PromptAttack.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord struct {
	// Keyword detection result object list
	BlockWordGroupInfoList []*GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList `json:"BlockWordGroupInfoList,omitempty" xml:"BlockWordGroupInfoList,omitempty" type:"Repeated"`
	// 0: No risk
	//
	// 1: Risk exists
	//
	// example:
	//
	// 0
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord) String() string {
	return dara.Prettify(s)
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord) GoString() string {
	return s.String()
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord) GetBlockWordGroupInfoList() []*GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList {
	return s.BlockWordGroupInfoList
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord) SetBlockWordGroupInfoList(v []*GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord {
	s.BlockWordGroupInfoList = v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord) SetRiskResult(v int32) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord {
	s.RiskResult = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord) Validate() error {
	if s.BlockWordGroupInfoList != nil {
		for _, item := range s.BlockWordGroupInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList struct {
	// Keyword detection result object list
	BlockWordList []*GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList `json:"BlockWordList,omitempty" xml:"BlockWordList,omitempty" type:"Repeated"`
	// Keyword group name
	//
	// example:
	//
	// testGroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList) GoString() string {
	return s.String()
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList) GetBlockWordList() []*GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList {
	return s.BlockWordList
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList) GetGroupName() *string {
	return s.GroupName
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList) SetBlockWordList(v []*GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList {
	s.BlockWordList = v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList) SetGroupName(v string) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList {
	s.GroupName = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList) Validate() error {
	if s.BlockWordList != nil {
		for _, item := range s.BlockWordList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList struct {
	// Word
	//
	// example:
	//
	// testWord
	Word *string `json:"Word,omitempty" xml:"Word,omitempty"`
	// Label
	//
	// example:
	//
	// testLabel
	WordLabel *string `json:"WordLabel,omitempty" xml:"WordLabel,omitempty"`
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) String() string {
	return dara.Prettify(s)
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) GoString() string {
	return s.String()
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) GetWord() *string {
	return s.Word
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) GetWordLabel() *string {
	return s.WordLabel
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) SetWord(v string) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList {
	s.Word = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) SetWordLabel(v string) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList {
	s.WordLabel = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) Validate() error {
	return dara.Validate(s)
}

type GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopics struct {
	// ConfidenceScore
	//
	// example:
	//
	// 0.0
	ConfidenceScore *float64 `json:"ConfidenceScore,omitempty" xml:"ConfidenceScore,omitempty"`
	// 0: No risk
	//
	// 1: Risk exists
	//
	// example:
	//
	// 1
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
	// List of sensitive topics
	TopicInfoList []*GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList `json:"TopicInfoList,omitempty" xml:"TopicInfoList,omitempty" type:"Repeated"`
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopics) String() string {
	return dara.Prettify(s)
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopics) GoString() string {
	return s.String()
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopics) GetConfidenceScore() *float64 {
	return s.ConfidenceScore
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopics) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopics) GetTopicInfoList() []*GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList {
	return s.TopicInfoList
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopics) SetConfidenceScore(v float64) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopics {
	s.ConfidenceScore = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopics) SetRiskResult(v int32) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopics {
	s.RiskResult = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopics) SetTopicInfoList(v []*GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopics {
	s.TopicInfoList = v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopics) Validate() error {
	if s.TopicInfoList != nil {
		for _, item := range s.TopicInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList struct {
	// 0: Text
	//
	// 1: Image
	//
	// example:
	//
	// 0
	CategoryType *int32 `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// 0: No risk
	//
	// 1: Risk exists
	//
	// example:
	//
	// 1
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
	// Security level
	//
	// 0: Low
	//
	// 1: Medium
	//
	// 2: High
	//
	// example:
	//
	// 0
	SecurityLevel *int32 `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	// Topic name
	//
	// example:
	//
	// Buss.
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList) GoString() string {
	return s.String()
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList) GetCategoryType() *int32 {
	return s.CategoryType
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList) GetTopicName() *string {
	return s.TopicName
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList) SetCategoryType(v int32) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList {
	s.CategoryType = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList) SetRiskResult(v int32) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList {
	s.RiskResult = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList) SetSecurityLevel(v int32) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList {
	s.SecurityLevel = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList) SetTopicName(v string) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList {
	s.TopicName = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList) Validate() error {
	return dara.Validate(s)
}

type GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories struct {
	// Confidence score
	//
	// example:
	//
	// 0.0
	ConfidenceScore *float64 `json:"ConfidenceScore,omitempty" xml:"ConfidenceScore,omitempty"`
	// List of harmful category objects
	HarmfulCategoryInfoList []*GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList `json:"HarmfulCategoryInfoList,omitempty" xml:"HarmfulCategoryInfoList,omitempty" type:"Repeated"`
	// 0: No risk
	//
	// 1: Risk exists
	//
	// example:
	//
	// 0
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories) String() string {
	return dara.Prettify(s)
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories) GoString() string {
	return s.String()
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories) GetConfidenceScore() *float64 {
	return s.ConfidenceScore
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories) GetHarmfulCategoryInfoList() []*GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	return s.HarmfulCategoryInfoList
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories) SetConfidenceScore(v float64) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories {
	s.ConfidenceScore = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories) SetHarmfulCategoryInfoList(v []*GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories {
	s.HarmfulCategoryInfoList = v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories) SetRiskResult(v int32) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories {
	s.RiskResult = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories) Validate() error {
	if s.HarmfulCategoryInfoList != nil {
		for _, item := range s.HarmfulCategoryInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList struct {
	// Category name
	//
	// example:
	//
	// Morality
	CategoryLabel *string `json:"CategoryLabel,omitempty" xml:"CategoryLabel,omitempty"`
	// 0: Text
	//
	// 1: Image
	//
	// example:
	//
	// 0
	CategoryType *int32 `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// 0: No risk
	//
	// 1: Risk exists
	//
	// example:
	//
	// 0
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
	// Security level
	//
	// 0: Low
	//
	// 1: Medium
	//
	// 2: High
	//
	// example:
	//
	// 0
	SecurityLevel *int32 `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	// Sub-category label
	//
	// example:
	//
	// morality_ethics
	SubCategoryLabel *string `json:"SubCategoryLabel,omitempty" xml:"SubCategoryLabel,omitempty"`
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) GoString() string {
	return s.String()
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) GetCategoryLabel() *string {
	return s.CategoryLabel
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) GetCategoryType() *int32 {
	return s.CategoryType
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) GetSubCategoryLabel() *string {
	return s.SubCategoryLabel
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetCategoryLabel(v string) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.CategoryLabel = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetCategoryType(v int32) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.CategoryType = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetRiskResult(v int32) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.RiskResult = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetSecurityLevel(v int32) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.SecurityLevel = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetSubCategoryLabel(v string) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.SubCategoryLabel = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) Validate() error {
	return dara.Validate(s)
}

type GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack struct {
	// Confidence score
	//
	// example:
	//
	// 0.0
	ConfidenceScore *float64 `json:"ConfidenceScore,omitempty" xml:"ConfidenceScore,omitempty"`
	// Prompt attack detection result object
	//
	// example:
	//
	// Role Play
	PromptAttackInfo *string `json:"PromptAttackInfo,omitempty" xml:"PromptAttackInfo,omitempty"`
	// Prompt attack list
	PromptAttackInfoList []*GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList `json:"PromptAttackInfoList,omitempty" xml:"PromptAttackInfoList,omitempty" type:"Repeated"`
	// 0: No risk
	//
	// 1: Risk exists
	//
	// example:
	//
	// 0
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
	// Security level
	//
	// 0: Low
	//
	// 1: Medium
	//
	// 2: High
	//
	// example:
	//
	// 0
	SecurityLevel *int32 `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) String() string {
	return dara.Prettify(s)
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) GoString() string {
	return s.String()
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) GetConfidenceScore() *float64 {
	return s.ConfidenceScore
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) GetPromptAttackInfo() *string {
	return s.PromptAttackInfo
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) GetPromptAttackInfoList() []*GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList {
	return s.PromptAttackInfoList
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) SetConfidenceScore(v float64) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack {
	s.ConfidenceScore = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) SetPromptAttackInfo(v string) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack {
	s.PromptAttackInfo = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) SetPromptAttackInfoList(v []*GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack {
	s.PromptAttackInfoList = v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) SetRiskResult(v int32) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack {
	s.RiskResult = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) SetSecurityLevel(v int32) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack {
	s.SecurityLevel = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) Validate() error {
	if s.PromptAttackInfoList != nil {
		for _, item := range s.PromptAttackInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList struct {
	// Category name
	//
	// example:
	//
	// Role Play
	CategoryLabel *string `json:"CategoryLabel,omitempty" xml:"CategoryLabel,omitempty"`
	// 0: Text
	//
	// 1: Image
	//
	// example:
	//
	// 0
	CategoryType *int32 `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// 0: No risk
	//
	// 1: Risk exists
	//
	// example:
	//
	// 0
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
	// Security level
	//
	// 0: Low
	//
	// 1: Medium
	//
	// 2: High
	//
	// example:
	//
	// 0
	SecurityLevel *int32 `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList) GoString() string {
	return s.String()
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList) GetCategoryLabel() *string {
	return s.CategoryLabel
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList) GetCategoryType() *int32 {
	return s.CategoryType
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList) SetCategoryLabel(v string) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList {
	s.CategoryLabel = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList) SetCategoryType(v int32) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList {
	s.CategoryType = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList) SetRiskResult(v int32) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList {
	s.RiskResult = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList) SetSecurityLevel(v int32) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList {
	s.SecurityLevel = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList) Validate() error {
	return dara.Validate(s)
}
