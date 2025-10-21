// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentSafeModelInstanceId(v int64) *CreatePolicyRequest
	GetContentSafeModelInstanceId() *int64
	SetEnableSensitiveInputCheck(v int32) *CreatePolicyRequest
	GetEnableSensitiveInputCheck() *int32
	SetEnableSensitiveOutputCheck(v int32) *CreatePolicyRequest
	GetEnableSensitiveOutputCheck() *int32
	SetHarmfulCategoryConfigInfoList(v []*CreatePolicyRequestHarmfulCategoryConfigInfoList) *CreatePolicyRequest
	GetHarmfulCategoryConfigInfoList() []*CreatePolicyRequestHarmfulCategoryConfigInfoList
	SetInputSafeAnswer(v string) *CreatePolicyRequest
	GetInputSafeAnswer() *string
	SetInputSafeAnswerSwitch(v int32) *CreatePolicyRequest
	GetInputSafeAnswerSwitch() *int32
	SetIsSidecarPolicy(v int32) *CreatePolicyRequest
	GetIsSidecarPolicy() *int32
	SetOutputSafeAnswer(v string) *CreatePolicyRequest
	GetOutputSafeAnswer() *string
	SetOutputSafeAnswerSwitch(v int32) *CreatePolicyRequest
	GetOutputSafeAnswerSwitch() *int32
	SetPolicyName(v string) *CreatePolicyRequest
	GetPolicyName() *string
	SetPromptAttackInfo(v *CreatePolicyRequestPromptAttackInfo) *CreatePolicyRequest
	GetPromptAttackInfo() *CreatePolicyRequestPromptAttackInfo
	SetPromptAttackInfoList(v []*CreatePolicyRequestPromptAttackInfoList) *CreatePolicyRequest
	GetPromptAttackInfoList() []*CreatePolicyRequestPromptAttackInfoList
	SetPromptAttackModelInstanceId(v int64) *CreatePolicyRequest
	GetPromptAttackModelInstanceId() *int64
	SetRegionId(v string) *CreatePolicyRequest
	GetRegionId() *string
	SetRegularExpressList(v []*CreatePolicyRequestRegularExpressList) *CreatePolicyRequest
	GetRegularExpressList() []*CreatePolicyRequestRegularExpressList
	SetSceneType(v int32) *CreatePolicyRequest
	GetSceneType() *int32
	SetSensitiveConfigList(v []*CreatePolicyRequestSensitiveConfigList) *CreatePolicyRequest
	GetSensitiveConfigList() []*CreatePolicyRequestSensitiveConfigList
	SetSensitiveTopicList(v []*CreatePolicyRequestSensitiveTopicList) *CreatePolicyRequest
	GetSensitiveTopicList() []*CreatePolicyRequestSensitiveTopicList
	SetSensitiveTopicModelInstanceId(v int64) *CreatePolicyRequest
	GetSensitiveTopicModelInstanceId() *int64
	SetSensitiveWordList(v []*CreatePolicyRequestSensitiveWordList) *CreatePolicyRequest
	GetSensitiveWordList() []*CreatePolicyRequestSensitiveWordList
	SetTopicConfigInfoList(v []*CreatePolicyRequestTopicConfigInfoList) *CreatePolicyRequest
	GetTopicConfigInfoList() []*CreatePolicyRequestTopicConfigInfoList
	SetWordGroupInfoList(v []*CreatePolicyRequestWordGroupInfoList) *CreatePolicyRequest
	GetWordGroupInfoList() []*CreatePolicyRequestWordGroupInfoList
	SetWorkspaceId(v int64) *CreatePolicyRequest
	GetWorkspaceId() *int64
}

type CreatePolicyRequest struct {
	ContentSafeModelInstanceId *int64 `json:"ContentSafeModelInstanceId,omitempty" xml:"ContentSafeModelInstanceId,omitempty"`
	EnableSensitiveInputCheck  *int32 `json:"EnableSensitiveInputCheck,omitempty" xml:"EnableSensitiveInputCheck,omitempty"`
	EnableSensitiveOutputCheck *int32 `json:"EnableSensitiveOutputCheck,omitempty" xml:"EnableSensitiveOutputCheck,omitempty"`
	// List of harmful category configurations
	HarmfulCategoryConfigInfoList []*CreatePolicyRequestHarmfulCategoryConfigInfoList `json:"HarmfulCategoryConfigInfoList,omitempty" xml:"HarmfulCategoryConfigInfoList,omitempty" type:"Repeated"`
	InputSafeAnswer               *string                                             `json:"InputSafeAnswer,omitempty" xml:"InputSafeAnswer,omitempty"`
	InputSafeAnswerSwitch         *int32                                              `json:"InputSafeAnswerSwitch,omitempty" xml:"InputSafeAnswerSwitch,omitempty"`
	IsSidecarPolicy               *int32                                              `json:"IsSidecarPolicy,omitempty" xml:"IsSidecarPolicy,omitempty"`
	OutputSafeAnswer              *string                                             `json:"OutputSafeAnswer,omitempty" xml:"OutputSafeAnswer,omitempty"`
	OutputSafeAnswerSwitch        *int32                                              `json:"OutputSafeAnswerSwitch,omitempty" xml:"OutputSafeAnswerSwitch,omitempty"`
	// Detection policy name.
	//
	// example:
	//
	// testPolicy
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// Prompt attack detection result object
	PromptAttackInfo *CreatePolicyRequestPromptAttackInfo `json:"PromptAttackInfo,omitempty" xml:"PromptAttackInfo,omitempty" type:"Struct"`
	// List of prompt attacks
	PromptAttackInfoList        []*CreatePolicyRequestPromptAttackInfoList `json:"PromptAttackInfoList,omitempty" xml:"PromptAttackInfoList,omitempty" type:"Repeated"`
	PromptAttackModelInstanceId *int64                                     `json:"PromptAttackModelInstanceId,omitempty" xml:"PromptAttackModelInstanceId,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId                      *string                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegularExpressList            []*CreatePolicyRequestRegularExpressList  `json:"RegularExpressList,omitempty" xml:"RegularExpressList,omitempty" type:"Repeated"`
	SceneType                     *int32                                    `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
	SensitiveConfigList           []*CreatePolicyRequestSensitiveConfigList `json:"SensitiveConfigList,omitempty" xml:"SensitiveConfigList,omitempty" type:"Repeated"`
	SensitiveTopicList            []*CreatePolicyRequestSensitiveTopicList  `json:"SensitiveTopicList,omitempty" xml:"SensitiveTopicList,omitempty" type:"Repeated"`
	SensitiveTopicModelInstanceId *int64                                    `json:"SensitiveTopicModelInstanceId,omitempty" xml:"SensitiveTopicModelInstanceId,omitempty"`
	SensitiveWordList             []*CreatePolicyRequestSensitiveWordList   `json:"SensitiveWordList,omitempty" xml:"SensitiveWordList,omitempty" type:"Repeated"`
	// List of sensitive topics
	TopicConfigInfoList []*CreatePolicyRequestTopicConfigInfoList `json:"TopicConfigInfoList,omitempty" xml:"TopicConfigInfoList,omitempty" type:"Repeated"`
	// List of keyword group objects
	WordGroupInfoList []*CreatePolicyRequestWordGroupInfoList `json:"WordGroupInfoList,omitempty" xml:"WordGroupInfoList,omitempty" type:"Repeated"`
	// Workspace ID
	//
	// example:
	//
	// 608226
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreatePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyRequest) GetContentSafeModelInstanceId() *int64 {
	return s.ContentSafeModelInstanceId
}

func (s *CreatePolicyRequest) GetEnableSensitiveInputCheck() *int32 {
	return s.EnableSensitiveInputCheck
}

func (s *CreatePolicyRequest) GetEnableSensitiveOutputCheck() *int32 {
	return s.EnableSensitiveOutputCheck
}

func (s *CreatePolicyRequest) GetHarmfulCategoryConfigInfoList() []*CreatePolicyRequestHarmfulCategoryConfigInfoList {
	return s.HarmfulCategoryConfigInfoList
}

func (s *CreatePolicyRequest) GetInputSafeAnswer() *string {
	return s.InputSafeAnswer
}

func (s *CreatePolicyRequest) GetInputSafeAnswerSwitch() *int32 {
	return s.InputSafeAnswerSwitch
}

func (s *CreatePolicyRequest) GetIsSidecarPolicy() *int32 {
	return s.IsSidecarPolicy
}

func (s *CreatePolicyRequest) GetOutputSafeAnswer() *string {
	return s.OutputSafeAnswer
}

func (s *CreatePolicyRequest) GetOutputSafeAnswerSwitch() *int32 {
	return s.OutputSafeAnswerSwitch
}

func (s *CreatePolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *CreatePolicyRequest) GetPromptAttackInfo() *CreatePolicyRequestPromptAttackInfo {
	return s.PromptAttackInfo
}

func (s *CreatePolicyRequest) GetPromptAttackInfoList() []*CreatePolicyRequestPromptAttackInfoList {
	return s.PromptAttackInfoList
}

func (s *CreatePolicyRequest) GetPromptAttackModelInstanceId() *int64 {
	return s.PromptAttackModelInstanceId
}

func (s *CreatePolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePolicyRequest) GetRegularExpressList() []*CreatePolicyRequestRegularExpressList {
	return s.RegularExpressList
}

func (s *CreatePolicyRequest) GetSceneType() *int32 {
	return s.SceneType
}

func (s *CreatePolicyRequest) GetSensitiveConfigList() []*CreatePolicyRequestSensitiveConfigList {
	return s.SensitiveConfigList
}

func (s *CreatePolicyRequest) GetSensitiveTopicList() []*CreatePolicyRequestSensitiveTopicList {
	return s.SensitiveTopicList
}

func (s *CreatePolicyRequest) GetSensitiveTopicModelInstanceId() *int64 {
	return s.SensitiveTopicModelInstanceId
}

func (s *CreatePolicyRequest) GetSensitiveWordList() []*CreatePolicyRequestSensitiveWordList {
	return s.SensitiveWordList
}

func (s *CreatePolicyRequest) GetTopicConfigInfoList() []*CreatePolicyRequestTopicConfigInfoList {
	return s.TopicConfigInfoList
}

func (s *CreatePolicyRequest) GetWordGroupInfoList() []*CreatePolicyRequestWordGroupInfoList {
	return s.WordGroupInfoList
}

func (s *CreatePolicyRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *CreatePolicyRequest) SetContentSafeModelInstanceId(v int64) *CreatePolicyRequest {
	s.ContentSafeModelInstanceId = &v
	return s
}

func (s *CreatePolicyRequest) SetEnableSensitiveInputCheck(v int32) *CreatePolicyRequest {
	s.EnableSensitiveInputCheck = &v
	return s
}

func (s *CreatePolicyRequest) SetEnableSensitiveOutputCheck(v int32) *CreatePolicyRequest {
	s.EnableSensitiveOutputCheck = &v
	return s
}

func (s *CreatePolicyRequest) SetHarmfulCategoryConfigInfoList(v []*CreatePolicyRequestHarmfulCategoryConfigInfoList) *CreatePolicyRequest {
	s.HarmfulCategoryConfigInfoList = v
	return s
}

func (s *CreatePolicyRequest) SetInputSafeAnswer(v string) *CreatePolicyRequest {
	s.InputSafeAnswer = &v
	return s
}

func (s *CreatePolicyRequest) SetInputSafeAnswerSwitch(v int32) *CreatePolicyRequest {
	s.InputSafeAnswerSwitch = &v
	return s
}

func (s *CreatePolicyRequest) SetIsSidecarPolicy(v int32) *CreatePolicyRequest {
	s.IsSidecarPolicy = &v
	return s
}

func (s *CreatePolicyRequest) SetOutputSafeAnswer(v string) *CreatePolicyRequest {
	s.OutputSafeAnswer = &v
	return s
}

func (s *CreatePolicyRequest) SetOutputSafeAnswerSwitch(v int32) *CreatePolicyRequest {
	s.OutputSafeAnswerSwitch = &v
	return s
}

func (s *CreatePolicyRequest) SetPolicyName(v string) *CreatePolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *CreatePolicyRequest) SetPromptAttackInfo(v *CreatePolicyRequestPromptAttackInfo) *CreatePolicyRequest {
	s.PromptAttackInfo = v
	return s
}

func (s *CreatePolicyRequest) SetPromptAttackInfoList(v []*CreatePolicyRequestPromptAttackInfoList) *CreatePolicyRequest {
	s.PromptAttackInfoList = v
	return s
}

func (s *CreatePolicyRequest) SetPromptAttackModelInstanceId(v int64) *CreatePolicyRequest {
	s.PromptAttackModelInstanceId = &v
	return s
}

func (s *CreatePolicyRequest) SetRegionId(v string) *CreatePolicyRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePolicyRequest) SetRegularExpressList(v []*CreatePolicyRequestRegularExpressList) *CreatePolicyRequest {
	s.RegularExpressList = v
	return s
}

func (s *CreatePolicyRequest) SetSceneType(v int32) *CreatePolicyRequest {
	s.SceneType = &v
	return s
}

func (s *CreatePolicyRequest) SetSensitiveConfigList(v []*CreatePolicyRequestSensitiveConfigList) *CreatePolicyRequest {
	s.SensitiveConfigList = v
	return s
}

func (s *CreatePolicyRequest) SetSensitiveTopicList(v []*CreatePolicyRequestSensitiveTopicList) *CreatePolicyRequest {
	s.SensitiveTopicList = v
	return s
}

func (s *CreatePolicyRequest) SetSensitiveTopicModelInstanceId(v int64) *CreatePolicyRequest {
	s.SensitiveTopicModelInstanceId = &v
	return s
}

func (s *CreatePolicyRequest) SetSensitiveWordList(v []*CreatePolicyRequestSensitiveWordList) *CreatePolicyRequest {
	s.SensitiveWordList = v
	return s
}

func (s *CreatePolicyRequest) SetTopicConfigInfoList(v []*CreatePolicyRequestTopicConfigInfoList) *CreatePolicyRequest {
	s.TopicConfigInfoList = v
	return s
}

func (s *CreatePolicyRequest) SetWordGroupInfoList(v []*CreatePolicyRequestWordGroupInfoList) *CreatePolicyRequest {
	s.WordGroupInfoList = v
	return s
}

func (s *CreatePolicyRequest) SetWorkspaceId(v int64) *CreatePolicyRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreatePolicyRequest) Validate() error {
	if s.HarmfulCategoryConfigInfoList != nil {
		for _, item := range s.HarmfulCategoryConfigInfoList {
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
	if s.RegularExpressList != nil {
		for _, item := range s.RegularExpressList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SensitiveConfigList != nil {
		for _, item := range s.SensitiveConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SensitiveTopicList != nil {
		for _, item := range s.SensitiveTopicList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SensitiveWordList != nil {
		for _, item := range s.SensitiveWordList {
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

type CreatePolicyRequestHarmfulCategoryConfigInfoList struct {
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
	// Whether it is enabled
	//
	// 0: Not enabled
	//
	// 1: Enabled
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

func (s CreatePolicyRequestHarmfulCategoryConfigInfoList) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyRequestHarmfulCategoryConfigInfoList) GoString() string {
	return s.String()
}

func (s *CreatePolicyRequestHarmfulCategoryConfigInfoList) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *CreatePolicyRequestHarmfulCategoryConfigInfoList) GetCategoryLabel() *string {
	return s.CategoryLabel
}

func (s *CreatePolicyRequestHarmfulCategoryConfigInfoList) GetCategoryType() *int32 {
	return s.CategoryType
}

func (s *CreatePolicyRequestHarmfulCategoryConfigInfoList) GetInputOutputType() *int32 {
	return s.InputOutputType
}

func (s *CreatePolicyRequestHarmfulCategoryConfigInfoList) GetIsEnabled() *int32 {
	return s.IsEnabled
}

func (s *CreatePolicyRequestHarmfulCategoryConfigInfoList) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *CreatePolicyRequestHarmfulCategoryConfigInfoList) SetCategoryId(v int64) *CreatePolicyRequestHarmfulCategoryConfigInfoList {
	s.CategoryId = &v
	return s
}

func (s *CreatePolicyRequestHarmfulCategoryConfigInfoList) SetCategoryLabel(v string) *CreatePolicyRequestHarmfulCategoryConfigInfoList {
	s.CategoryLabel = &v
	return s
}

func (s *CreatePolicyRequestHarmfulCategoryConfigInfoList) SetCategoryType(v int32) *CreatePolicyRequestHarmfulCategoryConfigInfoList {
	s.CategoryType = &v
	return s
}

func (s *CreatePolicyRequestHarmfulCategoryConfigInfoList) SetInputOutputType(v int32) *CreatePolicyRequestHarmfulCategoryConfigInfoList {
	s.InputOutputType = &v
	return s
}

func (s *CreatePolicyRequestHarmfulCategoryConfigInfoList) SetIsEnabled(v int32) *CreatePolicyRequestHarmfulCategoryConfigInfoList {
	s.IsEnabled = &v
	return s
}

func (s *CreatePolicyRequestHarmfulCategoryConfigInfoList) SetSecurityLevel(v int32) *CreatePolicyRequestHarmfulCategoryConfigInfoList {
	s.SecurityLevel = &v
	return s
}

func (s *CreatePolicyRequestHarmfulCategoryConfigInfoList) Validate() error {
	return dara.Validate(s)
}

type CreatePolicyRequestPromptAttackInfo struct {
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
	// 1
	SecurityLevel *int32 `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
}

func (s CreatePolicyRequestPromptAttackInfo) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyRequestPromptAttackInfo) GoString() string {
	return s.String()
}

func (s *CreatePolicyRequestPromptAttackInfo) GetIsEnabled() *int32 {
	return s.IsEnabled
}

func (s *CreatePolicyRequestPromptAttackInfo) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *CreatePolicyRequestPromptAttackInfo) SetIsEnabled(v int32) *CreatePolicyRequestPromptAttackInfo {
	s.IsEnabled = &v
	return s
}

func (s *CreatePolicyRequestPromptAttackInfo) SetSecurityLevel(v int32) *CreatePolicyRequestPromptAttackInfo {
	s.SecurityLevel = &v
	return s
}

func (s *CreatePolicyRequestPromptAttackInfo) Validate() error {
	return dara.Validate(s)
}

type CreatePolicyRequestPromptAttackInfoList struct {
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
	// 1
	SecurityLevel *int32 `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
}

func (s CreatePolicyRequestPromptAttackInfoList) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyRequestPromptAttackInfoList) GoString() string {
	return s.String()
}

func (s *CreatePolicyRequestPromptAttackInfoList) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *CreatePolicyRequestPromptAttackInfoList) GetCategoryLabel() *string {
	return s.CategoryLabel
}

func (s *CreatePolicyRequestPromptAttackInfoList) GetIsEnabled() *int32 {
	return s.IsEnabled
}

func (s *CreatePolicyRequestPromptAttackInfoList) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *CreatePolicyRequestPromptAttackInfoList) SetCategoryId(v int64) *CreatePolicyRequestPromptAttackInfoList {
	s.CategoryId = &v
	return s
}

func (s *CreatePolicyRequestPromptAttackInfoList) SetCategoryLabel(v string) *CreatePolicyRequestPromptAttackInfoList {
	s.CategoryLabel = &v
	return s
}

func (s *CreatePolicyRequestPromptAttackInfoList) SetIsEnabled(v int32) *CreatePolicyRequestPromptAttackInfoList {
	s.IsEnabled = &v
	return s
}

func (s *CreatePolicyRequestPromptAttackInfoList) SetSecurityLevel(v int32) *CreatePolicyRequestPromptAttackInfoList {
	s.SecurityLevel = &v
	return s
}

func (s *CreatePolicyRequestPromptAttackInfoList) Validate() error {
	return dara.Validate(s)
}

type CreatePolicyRequestRegularExpressList struct {
	ActionType         *int32  `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	RegularExpress     *string `json:"RegularExpress,omitempty" xml:"RegularExpress,omitempty"`
	RegularExpressName *string `json:"RegularExpressName,omitempty" xml:"RegularExpressName,omitempty"`
}

func (s CreatePolicyRequestRegularExpressList) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyRequestRegularExpressList) GoString() string {
	return s.String()
}

func (s *CreatePolicyRequestRegularExpressList) GetActionType() *int32 {
	return s.ActionType
}

func (s *CreatePolicyRequestRegularExpressList) GetRegularExpress() *string {
	return s.RegularExpress
}

func (s *CreatePolicyRequestRegularExpressList) GetRegularExpressName() *string {
	return s.RegularExpressName
}

func (s *CreatePolicyRequestRegularExpressList) SetActionType(v int32) *CreatePolicyRequestRegularExpressList {
	s.ActionType = &v
	return s
}

func (s *CreatePolicyRequestRegularExpressList) SetRegularExpress(v string) *CreatePolicyRequestRegularExpressList {
	s.RegularExpress = &v
	return s
}

func (s *CreatePolicyRequestRegularExpressList) SetRegularExpressName(v string) *CreatePolicyRequestRegularExpressList {
	s.RegularExpressName = &v
	return s
}

func (s *CreatePolicyRequestRegularExpressList) Validate() error {
	return dara.Validate(s)
}

type CreatePolicyRequestSensitiveConfigList struct {
	ActionType        *int32 `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	IsEnabled         *int32 `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
	SensitiveConfigId *int64 `json:"SensitiveConfigId,omitempty" xml:"SensitiveConfigId,omitempty"`
}

func (s CreatePolicyRequestSensitiveConfigList) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyRequestSensitiveConfigList) GoString() string {
	return s.String()
}

func (s *CreatePolicyRequestSensitiveConfigList) GetActionType() *int32 {
	return s.ActionType
}

func (s *CreatePolicyRequestSensitiveConfigList) GetIsEnabled() *int32 {
	return s.IsEnabled
}

func (s *CreatePolicyRequestSensitiveConfigList) GetSensitiveConfigId() *int64 {
	return s.SensitiveConfigId
}

func (s *CreatePolicyRequestSensitiveConfigList) SetActionType(v int32) *CreatePolicyRequestSensitiveConfigList {
	s.ActionType = &v
	return s
}

func (s *CreatePolicyRequestSensitiveConfigList) SetIsEnabled(v int32) *CreatePolicyRequestSensitiveConfigList {
	s.IsEnabled = &v
	return s
}

func (s *CreatePolicyRequestSensitiveConfigList) SetSensitiveConfigId(v int64) *CreatePolicyRequestSensitiveConfigList {
	s.SensitiveConfigId = &v
	return s
}

func (s *CreatePolicyRequestSensitiveConfigList) Validate() error {
	return dara.Validate(s)
}

type CreatePolicyRequestSensitiveTopicList struct {
	CategoryType         *int32                                                       `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	SecurityLevel        *int32                                                       `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	TopicDefinition      *string                                                      `json:"TopicDefinition,omitempty" xml:"TopicDefinition,omitempty"`
	TopicExampleInfoList []*CreatePolicyRequestSensitiveTopicListTopicExampleInfoList `json:"TopicExampleInfoList,omitempty" xml:"TopicExampleInfoList,omitempty" type:"Repeated"`
	TopicName            *string                                                      `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s CreatePolicyRequestSensitiveTopicList) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyRequestSensitiveTopicList) GoString() string {
	return s.String()
}

func (s *CreatePolicyRequestSensitiveTopicList) GetCategoryType() *int32 {
	return s.CategoryType
}

func (s *CreatePolicyRequestSensitiveTopicList) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *CreatePolicyRequestSensitiveTopicList) GetTopicDefinition() *string {
	return s.TopicDefinition
}

func (s *CreatePolicyRequestSensitiveTopicList) GetTopicExampleInfoList() []*CreatePolicyRequestSensitiveTopicListTopicExampleInfoList {
	return s.TopicExampleInfoList
}

func (s *CreatePolicyRequestSensitiveTopicList) GetTopicName() *string {
	return s.TopicName
}

func (s *CreatePolicyRequestSensitiveTopicList) SetCategoryType(v int32) *CreatePolicyRequestSensitiveTopicList {
	s.CategoryType = &v
	return s
}

func (s *CreatePolicyRequestSensitiveTopicList) SetSecurityLevel(v int32) *CreatePolicyRequestSensitiveTopicList {
	s.SecurityLevel = &v
	return s
}

func (s *CreatePolicyRequestSensitiveTopicList) SetTopicDefinition(v string) *CreatePolicyRequestSensitiveTopicList {
	s.TopicDefinition = &v
	return s
}

func (s *CreatePolicyRequestSensitiveTopicList) SetTopicExampleInfoList(v []*CreatePolicyRequestSensitiveTopicListTopicExampleInfoList) *CreatePolicyRequestSensitiveTopicList {
	s.TopicExampleInfoList = v
	return s
}

func (s *CreatePolicyRequestSensitiveTopicList) SetTopicName(v string) *CreatePolicyRequestSensitiveTopicList {
	s.TopicName = &v
	return s
}

func (s *CreatePolicyRequestSensitiveTopicList) Validate() error {
	if s.TopicExampleInfoList != nil {
		for _, item := range s.TopicExampleInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreatePolicyRequestSensitiveTopicListTopicExampleInfoList struct {
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ExampleType *int32  `json:"ExampleType,omitempty" xml:"ExampleType,omitempty"`
}

func (s CreatePolicyRequestSensitiveTopicListTopicExampleInfoList) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyRequestSensitiveTopicListTopicExampleInfoList) GoString() string {
	return s.String()
}

func (s *CreatePolicyRequestSensitiveTopicListTopicExampleInfoList) GetContent() *string {
	return s.Content
}

func (s *CreatePolicyRequestSensitiveTopicListTopicExampleInfoList) GetExampleType() *int32 {
	return s.ExampleType
}

func (s *CreatePolicyRequestSensitiveTopicListTopicExampleInfoList) SetContent(v string) *CreatePolicyRequestSensitiveTopicListTopicExampleInfoList {
	s.Content = &v
	return s
}

func (s *CreatePolicyRequestSensitiveTopicListTopicExampleInfoList) SetExampleType(v int32) *CreatePolicyRequestSensitiveTopicListTopicExampleInfoList {
	s.ExampleType = &v
	return s
}

func (s *CreatePolicyRequestSensitiveTopicListTopicExampleInfoList) Validate() error {
	return dara.Validate(s)
}

type CreatePolicyRequestSensitiveWordList struct {
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Word  *string `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s CreatePolicyRequestSensitiveWordList) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyRequestSensitiveWordList) GoString() string {
	return s.String()
}

func (s *CreatePolicyRequestSensitiveWordList) GetLabel() *string {
	return s.Label
}

func (s *CreatePolicyRequestSensitiveWordList) GetWord() *string {
	return s.Word
}

func (s *CreatePolicyRequestSensitiveWordList) SetLabel(v string) *CreatePolicyRequestSensitiveWordList {
	s.Label = &v
	return s
}

func (s *CreatePolicyRequestSensitiveWordList) SetWord(v string) *CreatePolicyRequestSensitiveWordList {
	s.Word = &v
	return s
}

func (s *CreatePolicyRequestSensitiveWordList) Validate() error {
	return dara.Validate(s)
}

type CreatePolicyRequestTopicConfigInfoList struct {
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
	// 1
	SecurityLevel *int32 `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	// Sensitive topic ID
	//
	// example:
	//
	// 4
	TopicId *int64 `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
	// Topic Name
	//
	// example:
	//
	// Buss.
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s CreatePolicyRequestTopicConfigInfoList) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyRequestTopicConfigInfoList) GoString() string {
	return s.String()
}

func (s *CreatePolicyRequestTopicConfigInfoList) GetCategoryType() *int32 {
	return s.CategoryType
}

func (s *CreatePolicyRequestTopicConfigInfoList) GetInputOutputType() *int32 {
	return s.InputOutputType
}

func (s *CreatePolicyRequestTopicConfigInfoList) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *CreatePolicyRequestTopicConfigInfoList) GetTopicId() *int64 {
	return s.TopicId
}

func (s *CreatePolicyRequestTopicConfigInfoList) GetTopicName() *string {
	return s.TopicName
}

func (s *CreatePolicyRequestTopicConfigInfoList) SetCategoryType(v int32) *CreatePolicyRequestTopicConfigInfoList {
	s.CategoryType = &v
	return s
}

func (s *CreatePolicyRequestTopicConfigInfoList) SetInputOutputType(v int32) *CreatePolicyRequestTopicConfigInfoList {
	s.InputOutputType = &v
	return s
}

func (s *CreatePolicyRequestTopicConfigInfoList) SetSecurityLevel(v int32) *CreatePolicyRequestTopicConfigInfoList {
	s.SecurityLevel = &v
	return s
}

func (s *CreatePolicyRequestTopicConfigInfoList) SetTopicId(v int64) *CreatePolicyRequestTopicConfigInfoList {
	s.TopicId = &v
	return s
}

func (s *CreatePolicyRequestTopicConfigInfoList) SetTopicName(v string) *CreatePolicyRequestTopicConfigInfoList {
	s.TopicName = &v
	return s
}

func (s *CreatePolicyRequestTopicConfigInfoList) Validate() error {
	return dara.Validate(s)
}

type CreatePolicyRequestWordGroupInfoList struct {
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

func (s CreatePolicyRequestWordGroupInfoList) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyRequestWordGroupInfoList) GoString() string {
	return s.String()
}

func (s *CreatePolicyRequestWordGroupInfoList) GetGroupId() *int64 {
	return s.GroupId
}

func (s *CreatePolicyRequestWordGroupInfoList) GetGroupName() *string {
	return s.GroupName
}

func (s *CreatePolicyRequestWordGroupInfoList) GetInputOutputType() *int32 {
	return s.InputOutputType
}

func (s *CreatePolicyRequestWordGroupInfoList) SetGroupId(v int64) *CreatePolicyRequestWordGroupInfoList {
	s.GroupId = &v
	return s
}

func (s *CreatePolicyRequestWordGroupInfoList) SetGroupName(v string) *CreatePolicyRequestWordGroupInfoList {
	s.GroupName = &v
	return s
}

func (s *CreatePolicyRequestWordGroupInfoList) SetInputOutputType(v int32) *CreatePolicyRequestWordGroupInfoList {
	s.InputOutputType = &v
	return s
}

func (s *CreatePolicyRequestWordGroupInfoList) Validate() error {
	return dara.Validate(s)
}
