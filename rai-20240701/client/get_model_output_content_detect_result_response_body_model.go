// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelOutputContentDetectResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetModelOutputContentDetectResultResponseBody
	GetCode() *string
	SetDetectResultList(v []*GetModelOutputContentDetectResultResponseBodyDetectResultList) *GetModelOutputContentDetectResultResponseBody
	GetDetectResultList() []*GetModelOutputContentDetectResultResponseBodyDetectResultList
	SetHttpStatusCode(v int32) *GetModelOutputContentDetectResultResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetModelOutputContentDetectResultResponseBody
	GetMessage() *string
	SetProcessedCount(v int32) *GetModelOutputContentDetectResultResponseBody
	GetProcessedCount() *int32
	SetRequestId(v string) *GetModelOutputContentDetectResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetModelOutputContentDetectResultResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *GetModelOutputContentDetectResultResponseBody
	GetTaskId() *string
	SetTaskStatus(v int32) *GetModelOutputContentDetectResultResponseBody
	GetTaskStatus() *int32
	SetTotalCount(v int32) *GetModelOutputContentDetectResultResponseBody
	GetTotalCount() *int32
}

type GetModelOutputContentDetectResultResponseBody struct {
	// Status code, 00000 indicates success; others indicate failure.
	//
	// example:
	//
	// 00000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// List of detection result objects
	DetectResultList []*GetModelOutputContentDetectResultResponseBodyDetectResultList `json:"DetectResultList,omitempty" xml:"DetectResultList,omitempty" type:"Repeated"`
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
	// Whether the operation was successful. true indicates success, false indicates failure.
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

func (s GetModelOutputContentDetectResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetModelOutputContentDetectResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetModelOutputContentDetectResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetModelOutputContentDetectResultResponseBody) GetDetectResultList() []*GetModelOutputContentDetectResultResponseBodyDetectResultList {
	return s.DetectResultList
}

func (s *GetModelOutputContentDetectResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetModelOutputContentDetectResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetModelOutputContentDetectResultResponseBody) GetProcessedCount() *int32 {
	return s.ProcessedCount
}

func (s *GetModelOutputContentDetectResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetModelOutputContentDetectResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetModelOutputContentDetectResultResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *GetModelOutputContentDetectResultResponseBody) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *GetModelOutputContentDetectResultResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetModelOutputContentDetectResultResponseBody) SetCode(v string) *GetModelOutputContentDetectResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBody) SetDetectResultList(v []*GetModelOutputContentDetectResultResponseBodyDetectResultList) *GetModelOutputContentDetectResultResponseBody {
	s.DetectResultList = v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBody) SetHttpStatusCode(v int32) *GetModelOutputContentDetectResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBody) SetMessage(v string) *GetModelOutputContentDetectResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBody) SetProcessedCount(v int32) *GetModelOutputContentDetectResultResponseBody {
	s.ProcessedCount = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBody) SetRequestId(v string) *GetModelOutputContentDetectResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBody) SetSuccess(v bool) *GetModelOutputContentDetectResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBody) SetTaskId(v string) *GetModelOutputContentDetectResultResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBody) SetTaskStatus(v int32) *GetModelOutputContentDetectResultResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBody) SetTotalCount(v int32) *GetModelOutputContentDetectResultResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBody) Validate() error {
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

type GetModelOutputContentDetectResultResponseBodyDetectResultList struct {
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
	// Inspection results
	TraceInfo *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfo `json:"TraceInfo,omitempty" xml:"TraceInfo,omitempty" type:"Struct"`
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultList) String() string {
	return dara.Prettify(s)
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultList) GoString() string {
	return s.String()
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultList) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultList) GetStatus() *int32 {
	return s.Status
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultList) GetTraceInfo() *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfo {
	return s.TraceInfo
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultList) SetRiskResult(v int32) *GetModelOutputContentDetectResultResponseBodyDetectResultList {
	s.RiskResult = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultList) SetStatus(v int32) *GetModelOutputContentDetectResultResponseBodyDetectResultList {
	s.Status = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultList) SetTraceInfo(v *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfo) *GetModelOutputContentDetectResultResponseBodyDetectResultList {
	s.TraceInfo = v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultList) Validate() error {
	if s.TraceInfo != nil {
		if err := s.TraceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfo struct {
	// Detected keywords
	BlockWord *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord `json:"BlockWord,omitempty" xml:"BlockWord,omitempty" type:"Struct"`
	// Sensitive topic object list
	DenyTopics *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopics `json:"DenyTopics,omitempty" xml:"DenyTopics,omitempty" type:"Struct"`
	// List of harmful category result objects
	HarmfulCategories *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories `json:"HarmfulCategories,omitempty" xml:"HarmfulCategories,omitempty" type:"Struct"`
	// PromptAttack
	PromptAttack *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack `json:"PromptAttack,omitempty" xml:"PromptAttack,omitempty" type:"Struct"`
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfo) GoString() string {
	return s.String()
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfo) GetBlockWord() *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord {
	return s.BlockWord
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfo) GetDenyTopics() *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopics {
	return s.DenyTopics
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfo) GetHarmfulCategories() *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories {
	return s.HarmfulCategories
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfo) GetPromptAttack() *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack {
	return s.PromptAttack
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfo) SetBlockWord(v *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfo {
	s.BlockWord = v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfo) SetDenyTopics(v *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopics) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfo {
	s.DenyTopics = v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfo) SetHarmfulCategories(v *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfo {
	s.HarmfulCategories = v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfo) SetPromptAttack(v *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfo {
	s.PromptAttack = v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfo) Validate() error {
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

type GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord struct {
	// List of keyword detection result objects
	BlockWordGroupInfoList []*GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList `json:"BlockWordGroupInfoList,omitempty" xml:"BlockWordGroupInfoList,omitempty" type:"Repeated"`
	// 0: No risk
	//
	// 1: Risk exists
	//
	// example:
	//
	// 0
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord) String() string {
	return dara.Prettify(s)
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord) GoString() string {
	return s.String()
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord) GetBlockWordGroupInfoList() []*GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList {
	return s.BlockWordGroupInfoList
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord) SetBlockWordGroupInfoList(v []*GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord {
	s.BlockWordGroupInfoList = v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord) SetRiskResult(v int32) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord {
	s.RiskResult = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord) Validate() error {
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

type GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList struct {
	// List of keyword detection results
	BlockWordList []*GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList `json:"BlockWordList,omitempty" xml:"BlockWordList,omitempty" type:"Repeated"`
	// Keyword group name
	//
	// example:
	//
	// testGroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList) GoString() string {
	return s.String()
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList) GetBlockWordList() []*GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList {
	return s.BlockWordList
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList) GetGroupName() *string {
	return s.GroupName
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList) SetBlockWordList(v []*GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList {
	s.BlockWordList = v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList) SetGroupName(v string) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList {
	s.GroupName = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList) Validate() error {
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

type GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList struct {
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

func (s GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) String() string {
	return dara.Prettify(s)
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) GoString() string {
	return s.String()
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) GetWord() *string {
	return s.Word
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) GetWordLabel() *string {
	return s.WordLabel
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) SetWord(v string) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList {
	s.Word = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) SetWordLabel(v string) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList {
	s.WordLabel = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) Validate() error {
	return dara.Validate(s)
}

type GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopics struct {
	// Confidence score
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
	// 0
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
	// List of sensitive topics
	TopicInfoList []*GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList `json:"TopicInfoList,omitempty" xml:"TopicInfoList,omitempty" type:"Repeated"`
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopics) String() string {
	return dara.Prettify(s)
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopics) GoString() string {
	return s.String()
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopics) GetConfidenceScore() *float64 {
	return s.ConfidenceScore
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopics) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopics) GetTopicInfoList() []*GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList {
	return s.TopicInfoList
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopics) SetConfidenceScore(v float64) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopics {
	s.ConfidenceScore = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopics) SetRiskResult(v int32) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopics {
	s.RiskResult = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopics) SetTopicInfoList(v []*GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopics {
	s.TopicInfoList = v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopics) Validate() error {
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

type GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList struct {
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
	// Topic name
	//
	// example:
	//
	// Buss.
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList) GoString() string {
	return s.String()
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList) GetCategoryType() *int32 {
	return s.CategoryType
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList) GetTopicName() *string {
	return s.TopicName
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList) SetCategoryType(v int32) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList {
	s.CategoryType = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList) SetRiskResult(v int32) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList {
	s.RiskResult = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList) SetSecurityLevel(v int32) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList {
	s.SecurityLevel = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList) SetTopicName(v string) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList {
	s.TopicName = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoDenyTopicsTopicInfoList) Validate() error {
	return dara.Validate(s)
}

type GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories struct {
	// Confidence score
	//
	// example:
	//
	// 0.0
	ConfidenceScore *float64 `json:"ConfidenceScore,omitempty" xml:"ConfidenceScore,omitempty"`
	// List of harmful category objects
	HarmfulCategoryInfoList []*GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList `json:"HarmfulCategoryInfoList,omitempty" xml:"HarmfulCategoryInfoList,omitempty" type:"Repeated"`
	// 0: No risk
	//
	// 1: Risk exists
	//
	// example:
	//
	// 0
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories) String() string {
	return dara.Prettify(s)
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories) GoString() string {
	return s.String()
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories) GetConfidenceScore() *float64 {
	return s.ConfidenceScore
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories) GetHarmfulCategoryInfoList() []*GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	return s.HarmfulCategoryInfoList
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories) SetConfidenceScore(v float64) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories {
	s.ConfidenceScore = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories) SetHarmfulCategoryInfoList(v []*GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories {
	s.HarmfulCategoryInfoList = v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories) SetRiskResult(v int32) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories {
	s.RiskResult = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories) Validate() error {
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

type GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList struct {
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
	// 1
	SecurityLevel *int32 `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	// Sub-category label
	//
	// example:
	//
	// morality_ethics
	SubCategoryLabel *string `json:"SubCategoryLabel,omitempty" xml:"SubCategoryLabel,omitempty"`
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) GoString() string {
	return s.String()
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) GetCategoryLabel() *string {
	return s.CategoryLabel
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) GetCategoryType() *int32 {
	return s.CategoryType
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) GetSubCategoryLabel() *string {
	return s.SubCategoryLabel
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetCategoryLabel(v string) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.CategoryLabel = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetCategoryType(v int32) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.CategoryType = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetRiskResult(v int32) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.RiskResult = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetSecurityLevel(v int32) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.SecurityLevel = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetSubCategoryLabel(v string) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.SubCategoryLabel = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) Validate() error {
	return dara.Validate(s)
}

type GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack struct {
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
	PromptAttackInfoList []*GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList `json:"PromptAttackInfoList,omitempty" xml:"PromptAttackInfoList,omitempty" type:"Repeated"`
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
	// 1
	SecurityLevel *int32 `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) String() string {
	return dara.Prettify(s)
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) GoString() string {
	return s.String()
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) GetConfidenceScore() *float64 {
	return s.ConfidenceScore
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) GetPromptAttackInfo() *string {
	return s.PromptAttackInfo
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) GetPromptAttackInfoList() []*GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList {
	return s.PromptAttackInfoList
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) SetConfidenceScore(v float64) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack {
	s.ConfidenceScore = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) SetPromptAttackInfo(v string) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack {
	s.PromptAttackInfo = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) SetPromptAttackInfoList(v []*GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack {
	s.PromptAttackInfoList = v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) SetRiskResult(v int32) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack {
	s.RiskResult = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) SetSecurityLevel(v int32) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack {
	s.SecurityLevel = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) Validate() error {
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

type GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList struct {
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
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList) GoString() string {
	return s.String()
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList) GetCategoryLabel() *string {
	return s.CategoryLabel
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList) GetCategoryType() *int32 {
	return s.CategoryType
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList) SetCategoryLabel(v string) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList {
	s.CategoryLabel = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList) SetCategoryType(v int32) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList {
	s.CategoryType = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList) SetRiskResult(v int32) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList {
	s.RiskResult = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList) SetSecurityLevel(v int32) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList {
	s.SecurityLevel = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttackPromptAttackInfoList) Validate() error {
	return dara.Validate(s)
}
