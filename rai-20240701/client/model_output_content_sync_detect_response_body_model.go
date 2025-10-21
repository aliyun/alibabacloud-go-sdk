// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelOutputContentSyncDetectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModelOutputContentSyncDetectResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ModelOutputContentSyncDetectResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModelOutputContentSyncDetectResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModelOutputContentSyncDetectResponseBody
	GetRequestId() *string
	SetRiskInfo(v string) *ModelOutputContentSyncDetectResponseBody
	GetRiskInfo() *string
	SetRiskResult(v int32) *ModelOutputContentSyncDetectResponseBody
	GetRiskResult() *int32
	SetSafeAnswer(v string) *ModelOutputContentSyncDetectResponseBody
	GetSafeAnswer() *string
	SetSuccess(v bool) *ModelOutputContentSyncDetectResponseBody
	GetSuccess() *bool
	SetTraceInfo(v *ModelOutputContentSyncDetectResponseBodyTraceInfo) *ModelOutputContentSyncDetectResponseBody
	GetTraceInfo() *ModelOutputContentSyncDetectResponseBodyTraceInfo
}

type ModelOutputContentSyncDetectResponseBody struct {
	// Status code, 00000 indicates success; others indicate failure.
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
	// Risk labels, separated by commas if multiple.
	//
	// example:
	//
	// “”
	RiskInfo *string `json:"RiskInfo,omitempty" xml:"RiskInfo,omitempty"`
	// 0: No risk
	//
	// 1: Risk exists
	//
	// example:
	//
	// 0
	RiskResult *int32  `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
	SafeAnswer *string `json:"SafeAnswer,omitempty" xml:"SafeAnswer,omitempty"`
	// 操作是否成功。true表示成功，false表示失败。
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Inspection result
	TraceInfo *ModelOutputContentSyncDetectResponseBodyTraceInfo `json:"TraceInfo,omitempty" xml:"TraceInfo,omitempty" type:"Struct"`
}

func (s ModelOutputContentSyncDetectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelOutputContentSyncDetectResponseBody) GoString() string {
	return s.String()
}

func (s *ModelOutputContentSyncDetectResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModelOutputContentSyncDetectResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelOutputContentSyncDetectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModelOutputContentSyncDetectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelOutputContentSyncDetectResponseBody) GetRiskInfo() *string {
	return s.RiskInfo
}

func (s *ModelOutputContentSyncDetectResponseBody) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *ModelOutputContentSyncDetectResponseBody) GetSafeAnswer() *string {
	return s.SafeAnswer
}

func (s *ModelOutputContentSyncDetectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelOutputContentSyncDetectResponseBody) GetTraceInfo() *ModelOutputContentSyncDetectResponseBodyTraceInfo {
	return s.TraceInfo
}

func (s *ModelOutputContentSyncDetectResponseBody) SetCode(v string) *ModelOutputContentSyncDetectResponseBody {
	s.Code = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBody) SetHttpStatusCode(v int32) *ModelOutputContentSyncDetectResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBody) SetMessage(v string) *ModelOutputContentSyncDetectResponseBody {
	s.Message = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBody) SetRequestId(v string) *ModelOutputContentSyncDetectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBody) SetRiskInfo(v string) *ModelOutputContentSyncDetectResponseBody {
	s.RiskInfo = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBody) SetRiskResult(v int32) *ModelOutputContentSyncDetectResponseBody {
	s.RiskResult = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBody) SetSafeAnswer(v string) *ModelOutputContentSyncDetectResponseBody {
	s.SafeAnswer = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBody) SetSuccess(v bool) *ModelOutputContentSyncDetectResponseBody {
	s.Success = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBody) SetTraceInfo(v *ModelOutputContentSyncDetectResponseBodyTraceInfo) *ModelOutputContentSyncDetectResponseBody {
	s.TraceInfo = v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBody) Validate() error {
	if s.TraceInfo != nil {
		if err := s.TraceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelOutputContentSyncDetectResponseBodyTraceInfo struct {
	// Detected keywords
	BlockWord *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWord `json:"BlockWord,omitempty" xml:"BlockWord,omitempty" type:"Struct"`
	// Sensitive topic object list
	DenyTopics *ModelOutputContentSyncDetectResponseBodyTraceInfoDenyTopics `json:"DenyTopics,omitempty" xml:"DenyTopics,omitempty" type:"Struct"`
	// HarmfulCategories
	HarmfulCategories *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategories `json:"HarmfulCategories,omitempty" xml:"HarmfulCategories,omitempty" type:"Struct"`
	// Prompt attack information
	PromptAttack  *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttack  `json:"PromptAttack,omitempty" xml:"PromptAttack,omitempty" type:"Struct"`
	SensitiveType *ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveType `json:"SensitiveType,omitempty" xml:"SensitiveType,omitempty" type:"Struct"`
}

func (s ModelOutputContentSyncDetectResponseBodyTraceInfo) String() string {
	return dara.Prettify(s)
}

func (s ModelOutputContentSyncDetectResponseBodyTraceInfo) GoString() string {
	return s.String()
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfo) GetBlockWord() *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWord {
	return s.BlockWord
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfo) GetDenyTopics() *ModelOutputContentSyncDetectResponseBodyTraceInfoDenyTopics {
	return s.DenyTopics
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfo) GetHarmfulCategories() *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategories {
	return s.HarmfulCategories
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfo) GetPromptAttack() *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttack {
	return s.PromptAttack
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfo) GetSensitiveType() *ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveType {
	return s.SensitiveType
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfo) SetBlockWord(v *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWord) *ModelOutputContentSyncDetectResponseBodyTraceInfo {
	s.BlockWord = v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfo) SetDenyTopics(v *ModelOutputContentSyncDetectResponseBodyTraceInfoDenyTopics) *ModelOutputContentSyncDetectResponseBodyTraceInfo {
	s.DenyTopics = v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfo) SetHarmfulCategories(v *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategories) *ModelOutputContentSyncDetectResponseBodyTraceInfo {
	s.HarmfulCategories = v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfo) SetPromptAttack(v *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttack) *ModelOutputContentSyncDetectResponseBodyTraceInfo {
	s.PromptAttack = v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfo) SetSensitiveType(v *ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveType) *ModelOutputContentSyncDetectResponseBodyTraceInfo {
	s.SensitiveType = v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfo) Validate() error {
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

type ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWord struct {
	// List of keyword detection result objects
	BlockWordGroupInfoList []*ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList `json:"BlockWordGroupInfoList,omitempty" xml:"BlockWordGroupInfoList,omitempty" type:"Repeated"`
	// 0: No risk
	//
	// 1: Risk exists
	//
	// example:
	//
	// 0
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
}

func (s ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWord) String() string {
	return dara.Prettify(s)
}

func (s ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWord) GoString() string {
	return s.String()
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWord) GetBlockWordGroupInfoList() []*ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList {
	return s.BlockWordGroupInfoList
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWord) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWord) SetBlockWordGroupInfoList(v []*ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList) *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWord {
	s.BlockWordGroupInfoList = v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWord) SetRiskResult(v int32) *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWord {
	s.RiskResult = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWord) Validate() error {
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

type ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList struct {
	// List of keyword detection result objects
	BlockWordList []*ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList `json:"BlockWordList,omitempty" xml:"BlockWordList,omitempty" type:"Repeated"`
	// Keyword group name
	//
	// example:
	//
	// testGroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList) String() string {
	return dara.Prettify(s)
}

func (s ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList) GoString() string {
	return s.String()
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList) GetBlockWordList() []*ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList {
	return s.BlockWordList
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList) GetGroupName() *string {
	return s.GroupName
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList) SetBlockWordList(v []*ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList {
	s.BlockWordList = v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList) SetGroupName(v string) *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList {
	s.GroupName = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList) Validate() error {
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

type ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList struct {
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

func (s ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) String() string {
	return dara.Prettify(s)
}

func (s ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) GoString() string {
	return s.String()
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) GetWord() *string {
	return s.Word
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) GetWordLabel() *string {
	return s.WordLabel
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) SetWord(v string) *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList {
	s.Word = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) SetWordLabel(v string) *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList {
	s.WordLabel = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) Validate() error {
	return dara.Validate(s)
}

type ModelOutputContentSyncDetectResponseBodyTraceInfoDenyTopics struct {
	// Confidence score
	//
	// example:
	//
	// 0.0
	ConfidenceScore *float64 `json:"ConfidenceScore,omitempty" xml:"ConfidenceScore,omitempty"`
	// 0: No risk
	//
	// 1: Risk present
	//
	// example:
	//
	// 0
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
	// Sensitive topic list
	TopicInfoList []*ModelOutputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList `json:"TopicInfoList,omitempty" xml:"TopicInfoList,omitempty" type:"Repeated"`
}

func (s ModelOutputContentSyncDetectResponseBodyTraceInfoDenyTopics) String() string {
	return dara.Prettify(s)
}

func (s ModelOutputContentSyncDetectResponseBodyTraceInfoDenyTopics) GoString() string {
	return s.String()
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoDenyTopics) GetConfidenceScore() *float64 {
	return s.ConfidenceScore
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoDenyTopics) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoDenyTopics) GetTopicInfoList() []*ModelOutputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList {
	return s.TopicInfoList
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoDenyTopics) SetConfidenceScore(v float64) *ModelOutputContentSyncDetectResponseBodyTraceInfoDenyTopics {
	s.ConfidenceScore = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoDenyTopics) SetRiskResult(v int32) *ModelOutputContentSyncDetectResponseBodyTraceInfoDenyTopics {
	s.RiskResult = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoDenyTopics) SetTopicInfoList(v []*ModelOutputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList) *ModelOutputContentSyncDetectResponseBodyTraceInfoDenyTopics {
	s.TopicInfoList = v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoDenyTopics) Validate() error {
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

type ModelOutputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList struct {
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
	// 1: Risk present
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

func (s ModelOutputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList) String() string {
	return dara.Prettify(s)
}

func (s ModelOutputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList) GoString() string {
	return s.String()
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList) GetCategoryType() *int32 {
	return s.CategoryType
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList) GetTopicName() *string {
	return s.TopicName
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList) SetCategoryType(v int32) *ModelOutputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList {
	s.CategoryType = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList) SetRiskResult(v int32) *ModelOutputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList {
	s.RiskResult = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList) SetSecurityLevel(v int32) *ModelOutputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList {
	s.SecurityLevel = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList) SetTopicName(v string) *ModelOutputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList {
	s.TopicName = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoDenyTopicsTopicInfoList) Validate() error {
	return dara.Validate(s)
}

type ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategories struct {
	// Confidence score
	//
	// example:
	//
	// 0.0
	ConfidenceScore *float64 `json:"ConfidenceScore,omitempty" xml:"ConfidenceScore,omitempty"`
	// List of harmful category objects
	HarmfulCategoryInfoList []*ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList `json:"HarmfulCategoryInfoList,omitempty" xml:"HarmfulCategoryInfoList,omitempty" type:"Repeated"`
	// 0: No risk
	//
	// 1: Risk exists
	//
	// example:
	//
	// 0
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
}

func (s ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategories) String() string {
	return dara.Prettify(s)
}

func (s ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategories) GoString() string {
	return s.String()
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategories) GetConfidenceScore() *float64 {
	return s.ConfidenceScore
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategories) GetHarmfulCategoryInfoList() []*ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	return s.HarmfulCategoryInfoList
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategories) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategories) SetConfidenceScore(v float64) *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategories {
	s.ConfidenceScore = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategories) SetHarmfulCategoryInfoList(v []*ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategories {
	s.HarmfulCategoryInfoList = v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategories) SetRiskResult(v int32) *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategories {
	s.RiskResult = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategories) Validate() error {
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

type ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList struct {
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

func (s ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) String() string {
	return dara.Prettify(s)
}

func (s ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) GoString() string {
	return s.String()
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) GetCategoryLabel() *string {
	return s.CategoryLabel
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) GetCategoryType() *int32 {
	return s.CategoryType
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) GetSubCategoryLabel() *string {
	return s.SubCategoryLabel
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetCategoryLabel(v string) *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.CategoryLabel = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetCategoryType(v int32) *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.CategoryType = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetRiskResult(v int32) *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.RiskResult = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetSecurityLevel(v int32) *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.SecurityLevel = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetSubCategoryLabel(v string) *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.SubCategoryLabel = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) Validate() error {
	return dara.Validate(s)
}

type ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttack struct {
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
	// List of prompt attacks
	PromptAttackInfoList []*ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList `json:"PromptAttackInfoList,omitempty" xml:"PromptAttackInfoList,omitempty" type:"Repeated"`
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

func (s ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttack) String() string {
	return dara.Prettify(s)
}

func (s ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttack) GoString() string {
	return s.String()
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttack) GetConfidenceScore() *float64 {
	return s.ConfidenceScore
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttack) GetPromptAttackInfo() *string {
	return s.PromptAttackInfo
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttack) GetPromptAttackInfoList() []*ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList {
	return s.PromptAttackInfoList
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttack) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttack) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttack) SetConfidenceScore(v float64) *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttack {
	s.ConfidenceScore = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttack) SetPromptAttackInfo(v string) *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttack {
	s.PromptAttackInfo = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttack) SetPromptAttackInfoList(v []*ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList) *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttack {
	s.PromptAttackInfoList = v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttack) SetRiskResult(v int32) *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttack {
	s.RiskResult = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttack) SetSecurityLevel(v int32) *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttack {
	s.SecurityLevel = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttack) Validate() error {
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

type ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList struct {
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
	// 1: Risk present
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

func (s ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList) String() string {
	return dara.Prettify(s)
}

func (s ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList) GoString() string {
	return s.String()
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList) GetCategoryLabel() *string {
	return s.CategoryLabel
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList) GetCategoryType() *int32 {
	return s.CategoryType
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList) SetCategoryLabel(v string) *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList {
	s.CategoryLabel = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList) SetCategoryType(v int32) *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList {
	s.CategoryType = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList) SetRiskResult(v int32) *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList {
	s.RiskResult = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList) SetSecurityLevel(v int32) *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList {
	s.SecurityLevel = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttackPromptAttackInfoList) Validate() error {
	return dara.Validate(s)
}

type ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveType struct {
	MaskedContent         *string                                                                                `json:"MaskedContent,omitempty" xml:"MaskedContent,omitempty"`
	RiskResult            *int32                                                                                 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
	SensitiveTypeInfoList []*ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList `json:"SensitiveTypeInfoList,omitempty" xml:"SensitiveTypeInfoList,omitempty" type:"Repeated"`
}

func (s ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveType) String() string {
	return dara.Prettify(s)
}

func (s ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveType) GoString() string {
	return s.String()
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveType) GetMaskedContent() *string {
	return s.MaskedContent
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveType) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveType) GetSensitiveTypeInfoList() []*ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList {
	return s.SensitiveTypeInfoList
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveType) SetMaskedContent(v string) *ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveType {
	s.MaskedContent = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveType) SetRiskResult(v int32) *ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveType {
	s.RiskResult = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveType) SetSensitiveTypeInfoList(v []*ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList) *ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveType {
	s.SensitiveTypeInfoList = v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveType) Validate() error {
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

type ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList struct {
	ActionType        *int32  `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	MaskedContent     *string `json:"MaskedContent,omitempty" xml:"MaskedContent,omitempty"`
	SensitiveCategory *int32  `json:"SensitiveCategory,omitempty" xml:"SensitiveCategory,omitempty"`
	SensitiveContent  *string `json:"SensitiveContent,omitempty" xml:"SensitiveContent,omitempty"`
	SensitiveTypeName *string `json:"SensitiveTypeName,omitempty" xml:"SensitiveTypeName,omitempty"`
}

func (s ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList) String() string {
	return dara.Prettify(s)
}

func (s ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList) GoString() string {
	return s.String()
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList) GetActionType() *int32 {
	return s.ActionType
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList) GetMaskedContent() *string {
	return s.MaskedContent
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList) GetSensitiveCategory() *int32 {
	return s.SensitiveCategory
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList) GetSensitiveContent() *string {
	return s.SensitiveContent
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList) GetSensitiveTypeName() *string {
	return s.SensitiveTypeName
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList) SetActionType(v int32) *ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList {
	s.ActionType = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList) SetMaskedContent(v string) *ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList {
	s.MaskedContent = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList) SetSensitiveCategory(v int32) *ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList {
	s.SensitiveCategory = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList) SetSensitiveContent(v string) *ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList {
	s.SensitiveContent = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList) SetSensitiveTypeName(v string) *ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList {
	s.SensitiveTypeName = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoSensitiveTypeSensitiveTypeInfoList) Validate() error {
	return dara.Validate(s)
}
