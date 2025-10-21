// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelInputContentSyncDetectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModelInputContentSyncDetectResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ModelInputContentSyncDetectResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModelInputContentSyncDetectResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModelInputContentSyncDetectResponseBody
	GetRequestId() *string
	SetRiskResult(v int32) *ModelInputContentSyncDetectResponseBody
	GetRiskResult() *int32
	SetSafeAnswer(v string) *ModelInputContentSyncDetectResponseBody
	GetSafeAnswer() *string
	SetSuccess(v bool) *ModelInputContentSyncDetectResponseBody
	GetSuccess() *bool
	SetTraceInfo(v *ModelInputContentSyncDetectResponseBodyTraceInfo) *ModelInputContentSyncDetectResponseBody
	GetTraceInfo() *ModelInputContentSyncDetectResponseBodyTraceInfo
}

type ModelInputContentSyncDetectResponseBody struct {
	// Result code, 00000 indicates success; others indicate failure.
	//
	// example:
	//
	// 00000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
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
	// Request ID
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 0: No risk
	//
	// 1: Risk exists
	//
	// example:
	//
	// 0
	RiskResult *int32  `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
	SafeAnswer *string `json:"SafeAnswer,omitempty" xml:"SafeAnswer,omitempty"`
	// Whether the operation was successful. true indicates success, false indicates failure.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Inspection result
	TraceInfo *ModelInputContentSyncDetectResponseBodyTraceInfo `json:"TraceInfo,omitempty" xml:"TraceInfo,omitempty" type:"Struct"`
}

func (s ModelInputContentSyncDetectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelInputContentSyncDetectResponseBody) GoString() string {
	return s.String()
}

func (s *ModelInputContentSyncDetectResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModelInputContentSyncDetectResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelInputContentSyncDetectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModelInputContentSyncDetectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelInputContentSyncDetectResponseBody) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *ModelInputContentSyncDetectResponseBody) GetSafeAnswer() *string {
	return s.SafeAnswer
}

func (s *ModelInputContentSyncDetectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelInputContentSyncDetectResponseBody) GetTraceInfo() *ModelInputContentSyncDetectResponseBodyTraceInfo {
	return s.TraceInfo
}

func (s *ModelInputContentSyncDetectResponseBody) SetCode(v string) *ModelInputContentSyncDetectResponseBody {
	s.Code = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBody) SetHttpStatusCode(v int32) *ModelInputContentSyncDetectResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBody) SetMessage(v string) *ModelInputContentSyncDetectResponseBody {
	s.Message = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBody) SetRequestId(v string) *ModelInputContentSyncDetectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBody) SetRiskResult(v int32) *ModelInputContentSyncDetectResponseBody {
	s.RiskResult = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBody) SetSafeAnswer(v string) *ModelInputContentSyncDetectResponseBody {
	s.SafeAnswer = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBody) SetSuccess(v bool) *ModelInputContentSyncDetectResponseBody {
	s.Success = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBody) SetTraceInfo(v *ModelInputContentSyncDetectResponseBodyTraceInfo) *ModelInputContentSyncDetectResponseBody {
	s.TraceInfo = v
	return s
}

func (s *ModelInputContentSyncDetectResponseBody) Validate() error {
	if s.TraceInfo != nil {
		if err := s.TraceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelInputContentSyncDetectResponseBodyTraceInfo struct {
	// Detected keywords
	BlockWord *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWord `json:"BlockWord,omitempty" xml:"BlockWord,omitempty" type:"Struct"`
	// Sensitive topic object list
	DenyTopics *ModelInputContentSyncDetectResponseBodyTraceInfoDenyTopics `json:"DenyTopics,omitempty" xml:"DenyTopics,omitempty" type:"Struct"`
	// HarmfulCategories
	HarmfulCategories *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategories `json:"HarmfulCategories,omitempty" xml:"HarmfulCategories,omitempty" type:"Struct"`
	// Prompt attack information
	PromptAttack  *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttack  `json:"PromptAttack,omitempty" xml:"PromptAttack,omitempty" type:"Struct"`
	SensitiveType *ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveType `json:"SensitiveType,omitempty" xml:"SensitiveType,omitempty" type:"Struct"`
}

func (s ModelInputContentSyncDetectResponseBodyTraceInfo) String() string {
	return dara.Prettify(s)
}

func (s ModelInputContentSyncDetectResponseBodyTraceInfo) GoString() string {
	return s.String()
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfo) GetBlockWord() *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWord {
	return s.BlockWord
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfo) GetDenyTopics() *ModelInputContentSyncDetectResponseBodyTraceInfoDenyTopics {
	return s.DenyTopics
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfo) GetHarmfulCategories() *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategories {
	return s.HarmfulCategories
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfo) GetPromptAttack() *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttack {
	return s.PromptAttack
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfo) GetSensitiveType() *ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveType {
	return s.SensitiveType
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfo) SetBlockWord(v *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWord) *ModelInputContentSyncDetectResponseBodyTraceInfo {
	s.BlockWord = v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfo) SetDenyTopics(v *ModelInputContentSyncDetectResponseBodyTraceInfoDenyTopics) *ModelInputContentSyncDetectResponseBodyTraceInfo {
	s.DenyTopics = v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfo) SetHarmfulCategories(v *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategories) *ModelInputContentSyncDetectResponseBodyTraceInfo {
	s.HarmfulCategories = v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfo) SetPromptAttack(v *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttack) *ModelInputContentSyncDetectResponseBodyTraceInfo {
	s.PromptAttack = v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfo) SetSensitiveType(v *ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveType) *ModelInputContentSyncDetectResponseBodyTraceInfo {
	s.SensitiveType = v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfo) Validate() error {
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
	if s.SensitiveType != nil {
		if err := s.SensitiveType.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelInputContentSyncDetectResponseBodyTraceInfoBlockWord struct {
	// List of keyword detection result objects
	BlockWordGroupInfoList []*ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList `json:"BlockWordGroupInfoList,omitempty" xml:"BlockWordGroupInfoList,omitempty" type:"Repeated"`
	// 0: No risk
	//
	// 1: Risk exists
	//
	// example:
	//
	// 0
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
}

func (s ModelInputContentSyncDetectResponseBodyTraceInfoBlockWord) String() string {
	return dara.Prettify(s)
}

func (s ModelInputContentSyncDetectResponseBodyTraceInfoBlockWord) GoString() string {
	return s.String()
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWord) GetBlockWordGroupInfoList() []*ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList {
	return s.BlockWordGroupInfoList
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWord) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWord) SetBlockWordGroupInfoList(v []*ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList) *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWord {
	s.BlockWordGroupInfoList = v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWord) SetRiskResult(v int32) *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWord {
	s.RiskResult = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWord) Validate() error {
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

type ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList struct {
	// List of keyword detection results
	BlockWordList []*ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList `json:"BlockWordList,omitempty" xml:"BlockWordList,omitempty" type:"Repeated"`
	// Keyword group name
	//
	// example:
	//
	// testGroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList) String() string {
	return dara.Prettify(s)
}

func (s ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList) GoString() string {
	return s.String()
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList) GetBlockWordList() []*ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList {
	return s.BlockWordList
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList) GetGroupName() *string {
	return s.GroupName
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList) SetBlockWordList(v []*ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList {
	s.BlockWordList = v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList) SetGroupName(v string) *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList {
	s.GroupName = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList) Validate() error {
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

type ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList struct {
	// Keyword
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

func (s ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) String() string {
	return dara.Prettify(s)
}

func (s ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) GoString() string {
	return s.String()
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) GetWord() *string {
	return s.Word
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) GetWordLabel() *string {
	return s.WordLabel
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) SetWord(v string) *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList {
	s.Word = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) SetWordLabel(v string) *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList {
	s.WordLabel = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) Validate() error {
	return dara.Validate(s)
}

type ModelInputContentSyncDetectResponseBodyTraceInfoDenyTopics struct {
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
	// 1
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
	// Sensitive topic list
	TopicInfoList []*ModelInputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList `json:"TopicInfoList,omitempty" xml:"TopicInfoList,omitempty" type:"Repeated"`
}

func (s ModelInputContentSyncDetectResponseBodyTraceInfoDenyTopics) String() string {
	return dara.Prettify(s)
}

func (s ModelInputContentSyncDetectResponseBodyTraceInfoDenyTopics) GoString() string {
	return s.String()
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoDenyTopics) GetConfidenceScore() *float64 {
	return s.ConfidenceScore
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoDenyTopics) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoDenyTopics) GetTopicInfoList() []*ModelInputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList {
	return s.TopicInfoList
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoDenyTopics) SetConfidenceScore(v float64) *ModelInputContentSyncDetectResponseBodyTraceInfoDenyTopics {
	s.ConfidenceScore = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoDenyTopics) SetRiskResult(v int32) *ModelInputContentSyncDetectResponseBodyTraceInfoDenyTopics {
	s.RiskResult = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoDenyTopics) SetTopicInfoList(v []*ModelInputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList) *ModelInputContentSyncDetectResponseBodyTraceInfoDenyTopics {
	s.TopicInfoList = v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoDenyTopics) Validate() error {
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

type ModelInputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList struct {
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

func (s ModelInputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList) String() string {
	return dara.Prettify(s)
}

func (s ModelInputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList) GoString() string {
	return s.String()
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList) GetCategoryType() *int32 {
	return s.CategoryType
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList) GetTopicName() *string {
	return s.TopicName
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList) SetCategoryType(v int32) *ModelInputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList {
	s.CategoryType = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList) SetRiskResult(v int32) *ModelInputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList {
	s.RiskResult = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList) SetSecurityLevel(v int32) *ModelInputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList {
	s.SecurityLevel = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList) SetTopicName(v string) *ModelInputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList {
	s.TopicName = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList) Validate() error {
	return dara.Validate(s)
}

type ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategories struct {
	// Confidence score
	//
	// example:
	//
	// 0.0
	ConfidenceScore *float64 `json:"ConfidenceScore,omitempty" xml:"ConfidenceScore,omitempty"`
	// List of harmful category objects
	HarmfulCategoryInfoList []*ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList `json:"HarmfulCategoryInfoList,omitempty" xml:"HarmfulCategoryInfoList,omitempty" type:"Repeated"`
	// 0: No risk
	//
	// 1: Risk exists
	//
	// example:
	//
	// 0
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
}

func (s ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategories) String() string {
	return dara.Prettify(s)
}

func (s ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategories) GoString() string {
	return s.String()
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategories) GetConfidenceScore() *float64 {
	return s.ConfidenceScore
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategories) GetHarmfulCategoryInfoList() []*ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	return s.HarmfulCategoryInfoList
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategories) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategories) SetConfidenceScore(v float64) *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategories {
	s.ConfidenceScore = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategories) SetHarmfulCategoryInfoList(v []*ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategories {
	s.HarmfulCategoryInfoList = v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategories) SetRiskResult(v int32) *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategories {
	s.RiskResult = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategories) Validate() error {
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

type ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList struct {
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
	// Subcategory label
	//
	// example:
	//
	// morality_ethics
	SubCategoryLabel *string `json:"SubCategoryLabel,omitempty" xml:"SubCategoryLabel,omitempty"`
}

func (s ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) String() string {
	return dara.Prettify(s)
}

func (s ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) GoString() string {
	return s.String()
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) GetCategoryLabel() *string {
	return s.CategoryLabel
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) GetCategoryType() *int32 {
	return s.CategoryType
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) GetSubCategoryLabel() *string {
	return s.SubCategoryLabel
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetCategoryLabel(v string) *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.CategoryLabel = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetCategoryType(v int32) *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.CategoryType = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetRiskResult(v int32) *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.RiskResult = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetSecurityLevel(v int32) *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.SecurityLevel = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetSubCategoryLabel(v string) *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.SubCategoryLabel = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) Validate() error {
	return dara.Validate(s)
}

type ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttack struct {
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
	// List of prompt attack objects
	PromptAttackInfoList []*ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList `json:"PromptAttackInfoList,omitempty" xml:"PromptAttackInfoList,omitempty" type:"Repeated"`
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

func (s ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttack) String() string {
	return dara.Prettify(s)
}

func (s ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttack) GoString() string {
	return s.String()
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttack) GetConfidenceScore() *float64 {
	return s.ConfidenceScore
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttack) GetPromptAttackInfo() *string {
	return s.PromptAttackInfo
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttack) GetPromptAttackInfoList() []*ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList {
	return s.PromptAttackInfoList
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttack) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttack) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttack) SetConfidenceScore(v float64) *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttack {
	s.ConfidenceScore = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttack) SetPromptAttackInfo(v string) *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttack {
	s.PromptAttackInfo = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttack) SetPromptAttackInfoList(v []*ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList) *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttack {
	s.PromptAttackInfoList = v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttack) SetRiskResult(v int32) *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttack {
	s.RiskResult = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttack) SetSecurityLevel(v int32) *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttack {
	s.SecurityLevel = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttack) Validate() error {
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

type ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList struct {
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
	// 1
	SecurityLevel *int32 `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
}

func (s ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList) String() string {
	return dara.Prettify(s)
}

func (s ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList) GoString() string {
	return s.String()
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList) GetCategoryLabel() *string {
	return s.CategoryLabel
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList) GetCategoryType() *int32 {
	return s.CategoryType
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList) SetCategoryLabel(v string) *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList {
	s.CategoryLabel = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList) SetCategoryType(v int32) *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList {
	s.CategoryType = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList) SetRiskResult(v int32) *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList {
	s.RiskResult = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList) SetSecurityLevel(v int32) *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList {
	s.SecurityLevel = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList) Validate() error {
	return dara.Validate(s)
}

type ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveType struct {
	MaskedContent         *string                                                                               `json:"MaskedContent,omitempty" xml:"MaskedContent,omitempty"`
	RiskResult            *int32                                                                                `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
	SensitiveTypeInfoList []*ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList `json:"SensitiveTypeInfoList,omitempty" xml:"SensitiveTypeInfoList,omitempty" type:"Repeated"`
}

func (s ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveType) String() string {
	return dara.Prettify(s)
}

func (s ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveType) GoString() string {
	return s.String()
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveType) GetMaskedContent() *string {
	return s.MaskedContent
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveType) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveType) GetSensitiveTypeInfoList() []*ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList {
	return s.SensitiveTypeInfoList
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveType) SetMaskedContent(v string) *ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveType {
	s.MaskedContent = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveType) SetRiskResult(v int32) *ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveType {
	s.RiskResult = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveType) SetSensitiveTypeInfoList(v []*ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList) *ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveType {
	s.SensitiveTypeInfoList = v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveType) Validate() error {
	if s.SensitiveTypeInfoList != nil {
		for _, item := range s.SensitiveTypeInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList struct {
	ActionType        *int32  `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	MaskedContent     *string `json:"MaskedContent,omitempty" xml:"MaskedContent,omitempty"`
	SensitiveCategory *int32  `json:"SensitiveCategory,omitempty" xml:"SensitiveCategory,omitempty"`
	SensitiveContent  *string `json:"SensitiveContent,omitempty" xml:"SensitiveContent,omitempty"`
	SensitiveTypeName *string `json:"SensitiveTypeName,omitempty" xml:"SensitiveTypeName,omitempty"`
}

func (s ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList) String() string {
	return dara.Prettify(s)
}

func (s ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList) GoString() string {
	return s.String()
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList) GetActionType() *int32 {
	return s.ActionType
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList) GetMaskedContent() *string {
	return s.MaskedContent
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList) GetSensitiveCategory() *int32 {
	return s.SensitiveCategory
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList) GetSensitiveContent() *string {
	return s.SensitiveContent
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList) GetSensitiveTypeName() *string {
	return s.SensitiveTypeName
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList) SetActionType(v int32) *ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList {
	s.ActionType = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList) SetMaskedContent(v string) *ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList {
	s.MaskedContent = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList) SetSensitiveCategory(v int32) *ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList {
	s.SensitiveCategory = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList) SetSensitiveContent(v string) *ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList {
	s.SensitiveContent = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList) SetSensitiveTypeName(v string) *ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList {
	s.SensitiveTypeName = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList) Validate() error {
	return dara.Validate(s)
}
