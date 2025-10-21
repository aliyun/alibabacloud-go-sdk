// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolicyDefaultOptionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPolicyDefaultOptionsResponseBody
	GetCode() *string
	SetEnableSensitiveInputCheck(v int32) *GetPolicyDefaultOptionsResponseBody
	GetEnableSensitiveInputCheck() *int32
	SetEnableSensitiveOutputCheck(v int32) *GetPolicyDefaultOptionsResponseBody
	GetEnableSensitiveOutputCheck() *int32
	SetHarmfulCategoryInfoList(v []*GetPolicyDefaultOptionsResponseBodyHarmfulCategoryInfoList) *GetPolicyDefaultOptionsResponseBody
	GetHarmfulCategoryInfoList() []*GetPolicyDefaultOptionsResponseBodyHarmfulCategoryInfoList
	SetHttpStatusCode(v int32) *GetPolicyDefaultOptionsResponseBody
	GetHttpStatusCode() *int32
	SetInputSafeAnswer(v string) *GetPolicyDefaultOptionsResponseBody
	GetInputSafeAnswer() *string
	SetInputSafeAnswerSwitch(v int32) *GetPolicyDefaultOptionsResponseBody
	GetInputSafeAnswerSwitch() *int32
	SetMessage(v string) *GetPolicyDefaultOptionsResponseBody
	GetMessage() *string
	SetOutputSafeAnswer(v string) *GetPolicyDefaultOptionsResponseBody
	GetOutputSafeAnswer() *string
	SetOutputSafeAnswerSwitch(v int32) *GetPolicyDefaultOptionsResponseBody
	GetOutputSafeAnswerSwitch() *int32
	SetPromptAttackInfo(v *GetPolicyDefaultOptionsResponseBodyPromptAttackInfo) *GetPolicyDefaultOptionsResponseBody
	GetPromptAttackInfo() *GetPolicyDefaultOptionsResponseBodyPromptAttackInfo
	SetPromptAttackInfoList(v []*GetPolicyDefaultOptionsResponseBodyPromptAttackInfoList) *GetPolicyDefaultOptionsResponseBody
	GetPromptAttackInfoList() []*GetPolicyDefaultOptionsResponseBodyPromptAttackInfoList
	SetRequestId(v string) *GetPolicyDefaultOptionsResponseBody
	GetRequestId() *string
	SetSensitiveDataTypeList(v []*GetPolicyDefaultOptionsResponseBodySensitiveDataTypeList) *GetPolicyDefaultOptionsResponseBody
	GetSensitiveDataTypeList() []*GetPolicyDefaultOptionsResponseBodySensitiveDataTypeList
	SetSuccess(v bool) *GetPolicyDefaultOptionsResponseBody
	GetSuccess() *bool
	SetTopicConfigInfoList(v []*GetPolicyDefaultOptionsResponseBodyTopicConfigInfoList) *GetPolicyDefaultOptionsResponseBody
	GetTopicConfigInfoList() []*GetPolicyDefaultOptionsResponseBodyTopicConfigInfoList
	SetWordGroupInfoList(v []*GetPolicyDefaultOptionsResponseBodyWordGroupInfoList) *GetPolicyDefaultOptionsResponseBody
	GetWordGroupInfoList() []*GetPolicyDefaultOptionsResponseBodyWordGroupInfoList
}

type GetPolicyDefaultOptionsResponseBody struct {
	// Status code, 00000 indicates success; others indicate failure.
	//
	// example:
	//
	// 00000
	Code                       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	EnableSensitiveInputCheck  *int32  `json:"EnableSensitiveInputCheck,omitempty" xml:"EnableSensitiveInputCheck,omitempty"`
	EnableSensitiveOutputCheck *int32  `json:"EnableSensitiveOutputCheck,omitempty" xml:"EnableSensitiveOutputCheck,omitempty"`
	// List of harmful category objects
	HarmfulCategoryInfoList []*GetPolicyDefaultOptionsResponseBodyHarmfulCategoryInfoList `json:"HarmfulCategoryInfoList,omitempty" xml:"HarmfulCategoryInfoList,omitempty" type:"Repeated"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode        *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	InputSafeAnswer       *string `json:"InputSafeAnswer,omitempty" xml:"InputSafeAnswer,omitempty"`
	InputSafeAnswerSwitch *int32  `json:"InputSafeAnswerSwitch,omitempty" xml:"InputSafeAnswerSwitch,omitempty"`
	// Return message.
	//
	// example:
	//
	// ""
	Message                *string `json:"Message,omitempty" xml:"Message,omitempty"`
	OutputSafeAnswer       *string `json:"OutputSafeAnswer,omitempty" xml:"OutputSafeAnswer,omitempty"`
	OutputSafeAnswerSwitch *int32  `json:"OutputSafeAnswerSwitch,omitempty" xml:"OutputSafeAnswerSwitch,omitempty"`
	// Prompt attack detection result object
	PromptAttackInfo *GetPolicyDefaultOptionsResponseBodyPromptAttackInfo `json:"PromptAttackInfo,omitempty" xml:"PromptAttackInfo,omitempty" type:"Struct"`
	// Prompt attack list
	PromptAttackInfoList []*GetPolicyDefaultOptionsResponseBodyPromptAttackInfoList `json:"PromptAttackInfoList,omitempty" xml:"PromptAttackInfoList,omitempty" type:"Repeated"`
	// Request ID
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId             *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SensitiveDataTypeList []*GetPolicyDefaultOptionsResponseBodySensitiveDataTypeList `json:"SensitiveDataTypeList,omitempty" xml:"SensitiveDataTypeList,omitempty" type:"Repeated"`
	// Indicates whether the operation was successful. `true` means success, `false` means failure.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Sensitive topic list
	TopicConfigInfoList []*GetPolicyDefaultOptionsResponseBodyTopicConfigInfoList `json:"TopicConfigInfoList,omitempty" xml:"TopicConfigInfoList,omitempty" type:"Repeated"`
	// List of keyword group objects
	WordGroupInfoList []*GetPolicyDefaultOptionsResponseBodyWordGroupInfoList `json:"WordGroupInfoList,omitempty" xml:"WordGroupInfoList,omitempty" type:"Repeated"`
}

func (s GetPolicyDefaultOptionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyDefaultOptionsResponseBody) GoString() string {
	return s.String()
}

func (s *GetPolicyDefaultOptionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPolicyDefaultOptionsResponseBody) GetEnableSensitiveInputCheck() *int32 {
	return s.EnableSensitiveInputCheck
}

func (s *GetPolicyDefaultOptionsResponseBody) GetEnableSensitiveOutputCheck() *int32 {
	return s.EnableSensitiveOutputCheck
}

func (s *GetPolicyDefaultOptionsResponseBody) GetHarmfulCategoryInfoList() []*GetPolicyDefaultOptionsResponseBodyHarmfulCategoryInfoList {
	return s.HarmfulCategoryInfoList
}

func (s *GetPolicyDefaultOptionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetPolicyDefaultOptionsResponseBody) GetInputSafeAnswer() *string {
	return s.InputSafeAnswer
}

func (s *GetPolicyDefaultOptionsResponseBody) GetInputSafeAnswerSwitch() *int32 {
	return s.InputSafeAnswerSwitch
}

func (s *GetPolicyDefaultOptionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPolicyDefaultOptionsResponseBody) GetOutputSafeAnswer() *string {
	return s.OutputSafeAnswer
}

func (s *GetPolicyDefaultOptionsResponseBody) GetOutputSafeAnswerSwitch() *int32 {
	return s.OutputSafeAnswerSwitch
}

func (s *GetPolicyDefaultOptionsResponseBody) GetPromptAttackInfo() *GetPolicyDefaultOptionsResponseBodyPromptAttackInfo {
	return s.PromptAttackInfo
}

func (s *GetPolicyDefaultOptionsResponseBody) GetPromptAttackInfoList() []*GetPolicyDefaultOptionsResponseBodyPromptAttackInfoList {
	return s.PromptAttackInfoList
}

func (s *GetPolicyDefaultOptionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPolicyDefaultOptionsResponseBody) GetSensitiveDataTypeList() []*GetPolicyDefaultOptionsResponseBodySensitiveDataTypeList {
	return s.SensitiveDataTypeList
}

func (s *GetPolicyDefaultOptionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPolicyDefaultOptionsResponseBody) GetTopicConfigInfoList() []*GetPolicyDefaultOptionsResponseBodyTopicConfigInfoList {
	return s.TopicConfigInfoList
}

func (s *GetPolicyDefaultOptionsResponseBody) GetWordGroupInfoList() []*GetPolicyDefaultOptionsResponseBodyWordGroupInfoList {
	return s.WordGroupInfoList
}

func (s *GetPolicyDefaultOptionsResponseBody) SetCode(v string) *GetPolicyDefaultOptionsResponseBody {
	s.Code = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBody) SetEnableSensitiveInputCheck(v int32) *GetPolicyDefaultOptionsResponseBody {
	s.EnableSensitiveInputCheck = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBody) SetEnableSensitiveOutputCheck(v int32) *GetPolicyDefaultOptionsResponseBody {
	s.EnableSensitiveOutputCheck = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBody) SetHarmfulCategoryInfoList(v []*GetPolicyDefaultOptionsResponseBodyHarmfulCategoryInfoList) *GetPolicyDefaultOptionsResponseBody {
	s.HarmfulCategoryInfoList = v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBody) SetHttpStatusCode(v int32) *GetPolicyDefaultOptionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBody) SetInputSafeAnswer(v string) *GetPolicyDefaultOptionsResponseBody {
	s.InputSafeAnswer = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBody) SetInputSafeAnswerSwitch(v int32) *GetPolicyDefaultOptionsResponseBody {
	s.InputSafeAnswerSwitch = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBody) SetMessage(v string) *GetPolicyDefaultOptionsResponseBody {
	s.Message = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBody) SetOutputSafeAnswer(v string) *GetPolicyDefaultOptionsResponseBody {
	s.OutputSafeAnswer = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBody) SetOutputSafeAnswerSwitch(v int32) *GetPolicyDefaultOptionsResponseBody {
	s.OutputSafeAnswerSwitch = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBody) SetPromptAttackInfo(v *GetPolicyDefaultOptionsResponseBodyPromptAttackInfo) *GetPolicyDefaultOptionsResponseBody {
	s.PromptAttackInfo = v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBody) SetPromptAttackInfoList(v []*GetPolicyDefaultOptionsResponseBodyPromptAttackInfoList) *GetPolicyDefaultOptionsResponseBody {
	s.PromptAttackInfoList = v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBody) SetRequestId(v string) *GetPolicyDefaultOptionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBody) SetSensitiveDataTypeList(v []*GetPolicyDefaultOptionsResponseBodySensitiveDataTypeList) *GetPolicyDefaultOptionsResponseBody {
	s.SensitiveDataTypeList = v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBody) SetSuccess(v bool) *GetPolicyDefaultOptionsResponseBody {
	s.Success = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBody) SetTopicConfigInfoList(v []*GetPolicyDefaultOptionsResponseBodyTopicConfigInfoList) *GetPolicyDefaultOptionsResponseBody {
	s.TopicConfigInfoList = v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBody) SetWordGroupInfoList(v []*GetPolicyDefaultOptionsResponseBodyWordGroupInfoList) *GetPolicyDefaultOptionsResponseBody {
	s.WordGroupInfoList = v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBody) Validate() error {
	if s.HarmfulCategoryInfoList != nil {
		for _, item := range s.HarmfulCategoryInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PromptAttackInfo != nil {
		if err := s.PromptAttackInfo.Validate(); err != nil {
			return err
		}
	}
	if s.PromptAttackInfoList != nil {
		for _, item := range s.PromptAttackInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SensitiveDataTypeList != nil {
		for _, item := range s.SensitiveDataTypeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TopicConfigInfoList != nil {
		for _, item := range s.TopicConfigInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.WordGroupInfoList != nil {
		for _, item := range s.WordGroupInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPolicyDefaultOptionsResponseBodyHarmfulCategoryInfoList struct {
	// Harmful category ID
	//
	// example:
	//
	// 1
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
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
	// Model input/output type
	//
	// 0: Input
	//
	// 1: Output
	//
	// example:
	//
	// 0
	InputOutputType *int32 `json:"InputOutputType,omitempty" xml:"InputOutputType,omitempty"`
	// Harmful category configuration switch
	//
	// 0: Off
	//
	// 1: On
	//
	// example:
	//
	// True
	IsEnabled *int32 `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
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
	// 2
	SecurityLevel *int32 `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
}

func (s GetPolicyDefaultOptionsResponseBodyHarmfulCategoryInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyDefaultOptionsResponseBodyHarmfulCategoryInfoList) GoString() string {
	return s.String()
}

func (s *GetPolicyDefaultOptionsResponseBodyHarmfulCategoryInfoList) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *GetPolicyDefaultOptionsResponseBodyHarmfulCategoryInfoList) GetCategoryLabel() *string {
	return s.CategoryLabel
}

func (s *GetPolicyDefaultOptionsResponseBodyHarmfulCategoryInfoList) GetCategoryType() *int32 {
	return s.CategoryType
}

func (s *GetPolicyDefaultOptionsResponseBodyHarmfulCategoryInfoList) GetInputOutputType() *int32 {
	return s.InputOutputType
}

func (s *GetPolicyDefaultOptionsResponseBodyHarmfulCategoryInfoList) GetIsEnabled() *int32 {
	return s.IsEnabled
}

func (s *GetPolicyDefaultOptionsResponseBodyHarmfulCategoryInfoList) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *GetPolicyDefaultOptionsResponseBodyHarmfulCategoryInfoList) SetCategoryId(v int64) *GetPolicyDefaultOptionsResponseBodyHarmfulCategoryInfoList {
	s.CategoryId = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBodyHarmfulCategoryInfoList) SetCategoryLabel(v string) *GetPolicyDefaultOptionsResponseBodyHarmfulCategoryInfoList {
	s.CategoryLabel = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBodyHarmfulCategoryInfoList) SetCategoryType(v int32) *GetPolicyDefaultOptionsResponseBodyHarmfulCategoryInfoList {
	s.CategoryType = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBodyHarmfulCategoryInfoList) SetInputOutputType(v int32) *GetPolicyDefaultOptionsResponseBodyHarmfulCategoryInfoList {
	s.InputOutputType = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBodyHarmfulCategoryInfoList) SetIsEnabled(v int32) *GetPolicyDefaultOptionsResponseBodyHarmfulCategoryInfoList {
	s.IsEnabled = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBodyHarmfulCategoryInfoList) SetSecurityLevel(v int32) *GetPolicyDefaultOptionsResponseBodyHarmfulCategoryInfoList {
	s.SecurityLevel = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBodyHarmfulCategoryInfoList) Validate() error {
	return dara.Validate(s)
}

type GetPolicyDefaultOptionsResponseBodyPromptAttackInfo struct {
	// Harmful category configuration switch
	//
	// 0: Off
	//
	// 1: On
	//
	// example:
	//
	// 1
	IsEnabled *int32 `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
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

func (s GetPolicyDefaultOptionsResponseBodyPromptAttackInfo) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyDefaultOptionsResponseBodyPromptAttackInfo) GoString() string {
	return s.String()
}

func (s *GetPolicyDefaultOptionsResponseBodyPromptAttackInfo) GetIsEnabled() *int32 {
	return s.IsEnabled
}

func (s *GetPolicyDefaultOptionsResponseBodyPromptAttackInfo) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *GetPolicyDefaultOptionsResponseBodyPromptAttackInfo) SetIsEnabled(v int32) *GetPolicyDefaultOptionsResponseBodyPromptAttackInfo {
	s.IsEnabled = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBodyPromptAttackInfo) SetSecurityLevel(v int32) *GetPolicyDefaultOptionsResponseBodyPromptAttackInfo {
	s.SecurityLevel = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBodyPromptAttackInfo) Validate() error {
	return dara.Validate(s)
}

type GetPolicyDefaultOptionsResponseBodyPromptAttackInfoList struct {
	// Harmful category ID
	//
	// example:
	//
	// 1
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// Category name
	//
	// example:
	//
	// Role Play
	CategoryLabel *string `json:"CategoryLabel,omitempty" xml:"CategoryLabel,omitempty"`
	// Prompt attack configuration switch
	//
	// 0: Off
	//
	// 1: On
	//
	// example:
	//
	// 1
	IsEnabled *int32 `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
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

func (s GetPolicyDefaultOptionsResponseBodyPromptAttackInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyDefaultOptionsResponseBodyPromptAttackInfoList) GoString() string {
	return s.String()
}

func (s *GetPolicyDefaultOptionsResponseBodyPromptAttackInfoList) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *GetPolicyDefaultOptionsResponseBodyPromptAttackInfoList) GetCategoryLabel() *string {
	return s.CategoryLabel
}

func (s *GetPolicyDefaultOptionsResponseBodyPromptAttackInfoList) GetIsEnabled() *int32 {
	return s.IsEnabled
}

func (s *GetPolicyDefaultOptionsResponseBodyPromptAttackInfoList) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *GetPolicyDefaultOptionsResponseBodyPromptAttackInfoList) SetCategoryId(v int64) *GetPolicyDefaultOptionsResponseBodyPromptAttackInfoList {
	s.CategoryId = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBodyPromptAttackInfoList) SetCategoryLabel(v string) *GetPolicyDefaultOptionsResponseBodyPromptAttackInfoList {
	s.CategoryLabel = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBodyPromptAttackInfoList) SetIsEnabled(v int32) *GetPolicyDefaultOptionsResponseBodyPromptAttackInfoList {
	s.IsEnabled = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBodyPromptAttackInfoList) SetSecurityLevel(v int32) *GetPolicyDefaultOptionsResponseBodyPromptAttackInfoList {
	s.SecurityLevel = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBodyPromptAttackInfoList) Validate() error {
	return dara.Validate(s)
}

type GetPolicyDefaultOptionsResponseBodySensitiveDataTypeList struct {
	ActionType        *int32  `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	DataType          *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	Example           *string `json:"Example,omitempty" xml:"Example,omitempty"`
	ExampleProcessed  *string `json:"ExampleProcessed,omitempty" xml:"ExampleProcessed,omitempty"`
	IsEnabled         *int32  `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
	MatchAndReplace   *string `json:"MatchAndReplace,omitempty" xml:"MatchAndReplace,omitempty"`
	SensitiveConfigId *int64  `json:"SensitiveConfigId,omitempty" xml:"SensitiveConfigId,omitempty"`
	SensitiveName     *string `json:"SensitiveName,omitempty" xml:"SensitiveName,omitempty"`
}

func (s GetPolicyDefaultOptionsResponseBodySensitiveDataTypeList) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyDefaultOptionsResponseBodySensitiveDataTypeList) GoString() string {
	return s.String()
}

func (s *GetPolicyDefaultOptionsResponseBodySensitiveDataTypeList) GetActionType() *int32 {
	return s.ActionType
}

func (s *GetPolicyDefaultOptionsResponseBodySensitiveDataTypeList) GetDataType() *string {
	return s.DataType
}

func (s *GetPolicyDefaultOptionsResponseBodySensitiveDataTypeList) GetExample() *string {
	return s.Example
}

func (s *GetPolicyDefaultOptionsResponseBodySensitiveDataTypeList) GetExampleProcessed() *string {
	return s.ExampleProcessed
}

func (s *GetPolicyDefaultOptionsResponseBodySensitiveDataTypeList) GetIsEnabled() *int32 {
	return s.IsEnabled
}

func (s *GetPolicyDefaultOptionsResponseBodySensitiveDataTypeList) GetMatchAndReplace() *string {
	return s.MatchAndReplace
}

func (s *GetPolicyDefaultOptionsResponseBodySensitiveDataTypeList) GetSensitiveConfigId() *int64 {
	return s.SensitiveConfigId
}

func (s *GetPolicyDefaultOptionsResponseBodySensitiveDataTypeList) GetSensitiveName() *string {
	return s.SensitiveName
}

func (s *GetPolicyDefaultOptionsResponseBodySensitiveDataTypeList) SetActionType(v int32) *GetPolicyDefaultOptionsResponseBodySensitiveDataTypeList {
	s.ActionType = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBodySensitiveDataTypeList) SetDataType(v string) *GetPolicyDefaultOptionsResponseBodySensitiveDataTypeList {
	s.DataType = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBodySensitiveDataTypeList) SetExample(v string) *GetPolicyDefaultOptionsResponseBodySensitiveDataTypeList {
	s.Example = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBodySensitiveDataTypeList) SetExampleProcessed(v string) *GetPolicyDefaultOptionsResponseBodySensitiveDataTypeList {
	s.ExampleProcessed = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBodySensitiveDataTypeList) SetIsEnabled(v int32) *GetPolicyDefaultOptionsResponseBodySensitiveDataTypeList {
	s.IsEnabled = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBodySensitiveDataTypeList) SetMatchAndReplace(v string) *GetPolicyDefaultOptionsResponseBodySensitiveDataTypeList {
	s.MatchAndReplace = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBodySensitiveDataTypeList) SetSensitiveConfigId(v int64) *GetPolicyDefaultOptionsResponseBodySensitiveDataTypeList {
	s.SensitiveConfigId = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBodySensitiveDataTypeList) SetSensitiveName(v string) *GetPolicyDefaultOptionsResponseBodySensitiveDataTypeList {
	s.SensitiveName = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBodySensitiveDataTypeList) Validate() error {
	return dara.Validate(s)
}

type GetPolicyDefaultOptionsResponseBodyTopicConfigInfoList struct {
	// 0: Text
	//
	// 1: Image
	//
	// example:
	//
	// 0
	CategoryType *int32 `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// Model input/output type
	//
	// 0: Input
	//
	// 1: Output
	//
	// example:
	//
	// 0
	InputOutputType *int32 `json:"InputOutputType,omitempty" xml:"InputOutputType,omitempty"`
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
	// Sensitive topic ID
	//
	// example:
	//
	// 4
	TopicId *int64 `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
	// Topic name
	//
	// example:
	//
	// Buss.
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s GetPolicyDefaultOptionsResponseBodyTopicConfigInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyDefaultOptionsResponseBodyTopicConfigInfoList) GoString() string {
	return s.String()
}

func (s *GetPolicyDefaultOptionsResponseBodyTopicConfigInfoList) GetCategoryType() *int32 {
	return s.CategoryType
}

func (s *GetPolicyDefaultOptionsResponseBodyTopicConfigInfoList) GetInputOutputType() *int32 {
	return s.InputOutputType
}

func (s *GetPolicyDefaultOptionsResponseBodyTopicConfigInfoList) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *GetPolicyDefaultOptionsResponseBodyTopicConfigInfoList) GetTopicId() *int64 {
	return s.TopicId
}

func (s *GetPolicyDefaultOptionsResponseBodyTopicConfigInfoList) GetTopicName() *string {
	return s.TopicName
}

func (s *GetPolicyDefaultOptionsResponseBodyTopicConfigInfoList) SetCategoryType(v int32) *GetPolicyDefaultOptionsResponseBodyTopicConfigInfoList {
	s.CategoryType = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBodyTopicConfigInfoList) SetInputOutputType(v int32) *GetPolicyDefaultOptionsResponseBodyTopicConfigInfoList {
	s.InputOutputType = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBodyTopicConfigInfoList) SetSecurityLevel(v int32) *GetPolicyDefaultOptionsResponseBodyTopicConfigInfoList {
	s.SecurityLevel = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBodyTopicConfigInfoList) SetTopicId(v int64) *GetPolicyDefaultOptionsResponseBodyTopicConfigInfoList {
	s.TopicId = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBodyTopicConfigInfoList) SetTopicName(v string) *GetPolicyDefaultOptionsResponseBodyTopicConfigInfoList {
	s.TopicName = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBodyTopicConfigInfoList) Validate() error {
	return dara.Validate(s)
}

type GetPolicyDefaultOptionsResponseBodyWordGroupInfoList struct {
	// Keyword group ID.
	//
	// example:
	//
	// 1
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Keyword group name
	//
	// example:
	//
	// testGroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Model input/output type
	//
	// 0: Input
	//
	// 1: Output
	//
	// example:
	//
	// 0
	InputOutputType *int32 `json:"InputOutputType,omitempty" xml:"InputOutputType,omitempty"`
}

func (s GetPolicyDefaultOptionsResponseBodyWordGroupInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyDefaultOptionsResponseBodyWordGroupInfoList) GoString() string {
	return s.String()
}

func (s *GetPolicyDefaultOptionsResponseBodyWordGroupInfoList) GetGroupId() *int64 {
	return s.GroupId
}

func (s *GetPolicyDefaultOptionsResponseBodyWordGroupInfoList) GetGroupName() *string {
	return s.GroupName
}

func (s *GetPolicyDefaultOptionsResponseBodyWordGroupInfoList) GetInputOutputType() *int32 {
	return s.InputOutputType
}

func (s *GetPolicyDefaultOptionsResponseBodyWordGroupInfoList) SetGroupId(v int64) *GetPolicyDefaultOptionsResponseBodyWordGroupInfoList {
	s.GroupId = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBodyWordGroupInfoList) SetGroupName(v string) *GetPolicyDefaultOptionsResponseBodyWordGroupInfoList {
	s.GroupName = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBodyWordGroupInfoList) SetInputOutputType(v int32) *GetPolicyDefaultOptionsResponseBodyWordGroupInfoList {
	s.InputOutputType = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponseBodyWordGroupInfoList) Validate() error {
	return dara.Validate(s)
}
