// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentSafeModelInstanceId(v int64) *UpdatePolicyRequest
	GetContentSafeModelInstanceId() *int64
	SetEnableSensitiveInputCheck(v int32) *UpdatePolicyRequest
	GetEnableSensitiveInputCheck() *int32
	SetEnableSensitiveOutputCheck(v int32) *UpdatePolicyRequest
	GetEnableSensitiveOutputCheck() *int32
	SetHarmfulCategoryConfigInfoList(v []*UpdatePolicyRequestHarmfulCategoryConfigInfoList) *UpdatePolicyRequest
	GetHarmfulCategoryConfigInfoList() []*UpdatePolicyRequestHarmfulCategoryConfigInfoList
	SetInputSafeAnswer(v string) *UpdatePolicyRequest
	GetInputSafeAnswer() *string
	SetInputSafeAnswerSwitch(v int32) *UpdatePolicyRequest
	GetInputSafeAnswerSwitch() *int32
	SetIsSidecarPolicy(v int32) *UpdatePolicyRequest
	GetIsSidecarPolicy() *int32
	SetOutputSafeAnswer(v string) *UpdatePolicyRequest
	GetOutputSafeAnswer() *string
	SetOutputSafeAnswerSwitch(v int32) *UpdatePolicyRequest
	GetOutputSafeAnswerSwitch() *int32
	SetPolicyId(v int64) *UpdatePolicyRequest
	GetPolicyId() *int64
	SetPolicyName(v string) *UpdatePolicyRequest
	GetPolicyName() *string
	SetPromptAttackInfo(v *UpdatePolicyRequestPromptAttackInfo) *UpdatePolicyRequest
	GetPromptAttackInfo() *UpdatePolicyRequestPromptAttackInfo
	SetPromptAttackInfoList(v []*UpdatePolicyRequestPromptAttackInfoList) *UpdatePolicyRequest
	GetPromptAttackInfoList() []*UpdatePolicyRequestPromptAttackInfoList
	SetPromptAttackModelInstanceId(v int64) *UpdatePolicyRequest
	GetPromptAttackModelInstanceId() *int64
	SetRegionId(v string) *UpdatePolicyRequest
	GetRegionId() *string
	SetRegularExpressList(v []*UpdatePolicyRequestRegularExpressList) *UpdatePolicyRequest
	GetRegularExpressList() []*UpdatePolicyRequestRegularExpressList
	SetSceneType(v int32) *UpdatePolicyRequest
	GetSceneType() *int32
	SetSensitiveConfigList(v []*UpdatePolicyRequestSensitiveConfigList) *UpdatePolicyRequest
	GetSensitiveConfigList() []*UpdatePolicyRequestSensitiveConfigList
	SetSensitiveTopicList(v []*UpdatePolicyRequestSensitiveTopicList) *UpdatePolicyRequest
	GetSensitiveTopicList() []*UpdatePolicyRequestSensitiveTopicList
	SetSensitiveTopicModelInstanceId(v int64) *UpdatePolicyRequest
	GetSensitiveTopicModelInstanceId() *int64
	SetSensitiveWordList(v []*UpdatePolicyRequestSensitiveWordList) *UpdatePolicyRequest
	GetSensitiveWordList() []*UpdatePolicyRequestSensitiveWordList
	SetTopicConfigInfoList(v []*UpdatePolicyRequestTopicConfigInfoList) *UpdatePolicyRequest
	GetTopicConfigInfoList() []*UpdatePolicyRequestTopicConfigInfoList
	SetWordGroupInfoList(v []*UpdatePolicyRequestWordGroupInfoList) *UpdatePolicyRequest
	GetWordGroupInfoList() []*UpdatePolicyRequestWordGroupInfoList
	SetWorkspaceId(v int64) *UpdatePolicyRequest
	GetWorkspaceId() *int64
}

type UpdatePolicyRequest struct {
	ContentSafeModelInstanceId *int64 `json:"ContentSafeModelInstanceId,omitempty" xml:"ContentSafeModelInstanceId,omitempty"`
	EnableSensitiveInputCheck  *int32 `json:"EnableSensitiveInputCheck,omitempty" xml:"EnableSensitiveInputCheck,omitempty"`
	EnableSensitiveOutputCheck *int32 `json:"EnableSensitiveOutputCheck,omitempty" xml:"EnableSensitiveOutputCheck,omitempty"`
	// List of harmful category configurations
	HarmfulCategoryConfigInfoList []*UpdatePolicyRequestHarmfulCategoryConfigInfoList `json:"HarmfulCategoryConfigInfoList,omitempty" xml:"HarmfulCategoryConfigInfoList,omitempty" type:"Repeated"`
	InputSafeAnswer               *string                                             `json:"InputSafeAnswer,omitempty" xml:"InputSafeAnswer,omitempty"`
	InputSafeAnswerSwitch         *int32                                              `json:"InputSafeAnswerSwitch,omitempty" xml:"InputSafeAnswerSwitch,omitempty"`
	IsSidecarPolicy               *int32                                              `json:"IsSidecarPolicy,omitempty" xml:"IsSidecarPolicy,omitempty"`
	OutputSafeAnswer              *string                                             `json:"OutputSafeAnswer,omitempty" xml:"OutputSafeAnswer,omitempty"`
	OutputSafeAnswerSwitch        *int32                                              `json:"OutputSafeAnswerSwitch,omitempty" xml:"OutputSafeAnswerSwitch,omitempty"`
	// Detection policy ID
	//
	// example:
	//
	// 16
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// Detection policy name.
	//
	// example:
	//
	// testPolicy
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// Prompt attack detection result object
	PromptAttackInfo *UpdatePolicyRequestPromptAttackInfo `json:"PromptAttackInfo,omitempty" xml:"PromptAttackInfo,omitempty" type:"Struct"`
	// List of prompt attacks
	PromptAttackInfoList        []*UpdatePolicyRequestPromptAttackInfoList `json:"PromptAttackInfoList,omitempty" xml:"PromptAttackInfoList,omitempty" type:"Repeated"`
	PromptAttackModelInstanceId *int64                                     `json:"PromptAttackModelInstanceId,omitempty" xml:"PromptAttackModelInstanceId,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId                      *string                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegularExpressList            []*UpdatePolicyRequestRegularExpressList  `json:"RegularExpressList,omitempty" xml:"RegularExpressList,omitempty" type:"Repeated"`
	SceneType                     *int32                                    `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
	SensitiveConfigList           []*UpdatePolicyRequestSensitiveConfigList `json:"SensitiveConfigList,omitempty" xml:"SensitiveConfigList,omitempty" type:"Repeated"`
	SensitiveTopicList            []*UpdatePolicyRequestSensitiveTopicList  `json:"SensitiveTopicList,omitempty" xml:"SensitiveTopicList,omitempty" type:"Repeated"`
	SensitiveTopicModelInstanceId *int64                                    `json:"SensitiveTopicModelInstanceId,omitempty" xml:"SensitiveTopicModelInstanceId,omitempty"`
	SensitiveWordList             []*UpdatePolicyRequestSensitiveWordList   `json:"SensitiveWordList,omitempty" xml:"SensitiveWordList,omitempty" type:"Repeated"`
	// List of sensitive topics
	TopicConfigInfoList []*UpdatePolicyRequestTopicConfigInfoList `json:"TopicConfigInfoList,omitempty" xml:"TopicConfigInfoList,omitempty" type:"Repeated"`
	// List of keyword group objects
	WordGroupInfoList []*UpdatePolicyRequestWordGroupInfoList `json:"WordGroupInfoList,omitempty" xml:"WordGroupInfoList,omitempty" type:"Repeated"`
	WorkspaceId       *int64                                  `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdatePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdatePolicyRequest) GetContentSafeModelInstanceId() *int64 {
	return s.ContentSafeModelInstanceId
}

func (s *UpdatePolicyRequest) GetEnableSensitiveInputCheck() *int32 {
	return s.EnableSensitiveInputCheck
}

func (s *UpdatePolicyRequest) GetEnableSensitiveOutputCheck() *int32 {
	return s.EnableSensitiveOutputCheck
}

func (s *UpdatePolicyRequest) GetHarmfulCategoryConfigInfoList() []*UpdatePolicyRequestHarmfulCategoryConfigInfoList {
	return s.HarmfulCategoryConfigInfoList
}

func (s *UpdatePolicyRequest) GetInputSafeAnswer() *string {
	return s.InputSafeAnswer
}

func (s *UpdatePolicyRequest) GetInputSafeAnswerSwitch() *int32 {
	return s.InputSafeAnswerSwitch
}

func (s *UpdatePolicyRequest) GetIsSidecarPolicy() *int32 {
	return s.IsSidecarPolicy
}

func (s *UpdatePolicyRequest) GetOutputSafeAnswer() *string {
	return s.OutputSafeAnswer
}

func (s *UpdatePolicyRequest) GetOutputSafeAnswerSwitch() *int32 {
	return s.OutputSafeAnswerSwitch
}

func (s *UpdatePolicyRequest) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *UpdatePolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *UpdatePolicyRequest) GetPromptAttackInfo() *UpdatePolicyRequestPromptAttackInfo {
	return s.PromptAttackInfo
}

func (s *UpdatePolicyRequest) GetPromptAttackInfoList() []*UpdatePolicyRequestPromptAttackInfoList {
	return s.PromptAttackInfoList
}

func (s *UpdatePolicyRequest) GetPromptAttackModelInstanceId() *int64 {
	return s.PromptAttackModelInstanceId
}

func (s *UpdatePolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdatePolicyRequest) GetRegularExpressList() []*UpdatePolicyRequestRegularExpressList {
	return s.RegularExpressList
}

func (s *UpdatePolicyRequest) GetSceneType() *int32 {
	return s.SceneType
}

func (s *UpdatePolicyRequest) GetSensitiveConfigList() []*UpdatePolicyRequestSensitiveConfigList {
	return s.SensitiveConfigList
}

func (s *UpdatePolicyRequest) GetSensitiveTopicList() []*UpdatePolicyRequestSensitiveTopicList {
	return s.SensitiveTopicList
}

func (s *UpdatePolicyRequest) GetSensitiveTopicModelInstanceId() *int64 {
	return s.SensitiveTopicModelInstanceId
}

func (s *UpdatePolicyRequest) GetSensitiveWordList() []*UpdatePolicyRequestSensitiveWordList {
	return s.SensitiveWordList
}

func (s *UpdatePolicyRequest) GetTopicConfigInfoList() []*UpdatePolicyRequestTopicConfigInfoList {
	return s.TopicConfigInfoList
}

func (s *UpdatePolicyRequest) GetWordGroupInfoList() []*UpdatePolicyRequestWordGroupInfoList {
	return s.WordGroupInfoList
}

func (s *UpdatePolicyRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *UpdatePolicyRequest) SetContentSafeModelInstanceId(v int64) *UpdatePolicyRequest {
	s.ContentSafeModelInstanceId = &v
	return s
}

func (s *UpdatePolicyRequest) SetEnableSensitiveInputCheck(v int32) *UpdatePolicyRequest {
	s.EnableSensitiveInputCheck = &v
	return s
}

func (s *UpdatePolicyRequest) SetEnableSensitiveOutputCheck(v int32) *UpdatePolicyRequest {
	s.EnableSensitiveOutputCheck = &v
	return s
}

func (s *UpdatePolicyRequest) SetHarmfulCategoryConfigInfoList(v []*UpdatePolicyRequestHarmfulCategoryConfigInfoList) *UpdatePolicyRequest {
	s.HarmfulCategoryConfigInfoList = v
	return s
}

func (s *UpdatePolicyRequest) SetInputSafeAnswer(v string) *UpdatePolicyRequest {
	s.InputSafeAnswer = &v
	return s
}

func (s *UpdatePolicyRequest) SetInputSafeAnswerSwitch(v int32) *UpdatePolicyRequest {
	s.InputSafeAnswerSwitch = &v
	return s
}

func (s *UpdatePolicyRequest) SetIsSidecarPolicy(v int32) *UpdatePolicyRequest {
	s.IsSidecarPolicy = &v
	return s
}

func (s *UpdatePolicyRequest) SetOutputSafeAnswer(v string) *UpdatePolicyRequest {
	s.OutputSafeAnswer = &v
	return s
}

func (s *UpdatePolicyRequest) SetOutputSafeAnswerSwitch(v int32) *UpdatePolicyRequest {
	s.OutputSafeAnswerSwitch = &v
	return s
}

func (s *UpdatePolicyRequest) SetPolicyId(v int64) *UpdatePolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *UpdatePolicyRequest) SetPolicyName(v string) *UpdatePolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *UpdatePolicyRequest) SetPromptAttackInfo(v *UpdatePolicyRequestPromptAttackInfo) *UpdatePolicyRequest {
	s.PromptAttackInfo = v
	return s
}

func (s *UpdatePolicyRequest) SetPromptAttackInfoList(v []*UpdatePolicyRequestPromptAttackInfoList) *UpdatePolicyRequest {
	s.PromptAttackInfoList = v
	return s
}

func (s *UpdatePolicyRequest) SetPromptAttackModelInstanceId(v int64) *UpdatePolicyRequest {
	s.PromptAttackModelInstanceId = &v
	return s
}

func (s *UpdatePolicyRequest) SetRegionId(v string) *UpdatePolicyRequest {
	s.RegionId = &v
	return s
}

func (s *UpdatePolicyRequest) SetRegularExpressList(v []*UpdatePolicyRequestRegularExpressList) *UpdatePolicyRequest {
	s.RegularExpressList = v
	return s
}

func (s *UpdatePolicyRequest) SetSceneType(v int32) *UpdatePolicyRequest {
	s.SceneType = &v
	return s
}

func (s *UpdatePolicyRequest) SetSensitiveConfigList(v []*UpdatePolicyRequestSensitiveConfigList) *UpdatePolicyRequest {
	s.SensitiveConfigList = v
	return s
}

func (s *UpdatePolicyRequest) SetSensitiveTopicList(v []*UpdatePolicyRequestSensitiveTopicList) *UpdatePolicyRequest {
	s.SensitiveTopicList = v
	return s
}

func (s *UpdatePolicyRequest) SetSensitiveTopicModelInstanceId(v int64) *UpdatePolicyRequest {
	s.SensitiveTopicModelInstanceId = &v
	return s
}

func (s *UpdatePolicyRequest) SetSensitiveWordList(v []*UpdatePolicyRequestSensitiveWordList) *UpdatePolicyRequest {
	s.SensitiveWordList = v
	return s
}

func (s *UpdatePolicyRequest) SetTopicConfigInfoList(v []*UpdatePolicyRequestTopicConfigInfoList) *UpdatePolicyRequest {
	s.TopicConfigInfoList = v
	return s
}

func (s *UpdatePolicyRequest) SetWordGroupInfoList(v []*UpdatePolicyRequestWordGroupInfoList) *UpdatePolicyRequest {
	s.WordGroupInfoList = v
	return s
}

func (s *UpdatePolicyRequest) SetWorkspaceId(v int64) *UpdatePolicyRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdatePolicyRequest) Validate() error {
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

type UpdatePolicyRequestHarmfulCategoryConfigInfoList struct {
	// Harmful category configuration ID
	//
	// example:
	//
	// 1
	CategoryConfigId *int64 `json:"CategoryConfigId,omitempty" xml:"CategoryConfigId,omitempty"`
	CategoryId       *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
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
	// 2
	SecurityLevel *int32 `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
}

func (s UpdatePolicyRequestHarmfulCategoryConfigInfoList) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyRequestHarmfulCategoryConfigInfoList) GoString() string {
	return s.String()
}

func (s *UpdatePolicyRequestHarmfulCategoryConfigInfoList) GetCategoryConfigId() *int64 {
	return s.CategoryConfigId
}

func (s *UpdatePolicyRequestHarmfulCategoryConfigInfoList) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *UpdatePolicyRequestHarmfulCategoryConfigInfoList) GetCategoryLabel() *string {
	return s.CategoryLabel
}

func (s *UpdatePolicyRequestHarmfulCategoryConfigInfoList) GetCategoryType() *int32 {
	return s.CategoryType
}

func (s *UpdatePolicyRequestHarmfulCategoryConfigInfoList) GetInputOutputType() *int32 {
	return s.InputOutputType
}

func (s *UpdatePolicyRequestHarmfulCategoryConfigInfoList) GetIsEnabled() *int32 {
	return s.IsEnabled
}

func (s *UpdatePolicyRequestHarmfulCategoryConfigInfoList) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *UpdatePolicyRequestHarmfulCategoryConfigInfoList) SetCategoryConfigId(v int64) *UpdatePolicyRequestHarmfulCategoryConfigInfoList {
	s.CategoryConfigId = &v
	return s
}

func (s *UpdatePolicyRequestHarmfulCategoryConfigInfoList) SetCategoryId(v int64) *UpdatePolicyRequestHarmfulCategoryConfigInfoList {
	s.CategoryId = &v
	return s
}

func (s *UpdatePolicyRequestHarmfulCategoryConfigInfoList) SetCategoryLabel(v string) *UpdatePolicyRequestHarmfulCategoryConfigInfoList {
	s.CategoryLabel = &v
	return s
}

func (s *UpdatePolicyRequestHarmfulCategoryConfigInfoList) SetCategoryType(v int32) *UpdatePolicyRequestHarmfulCategoryConfigInfoList {
	s.CategoryType = &v
	return s
}

func (s *UpdatePolicyRequestHarmfulCategoryConfigInfoList) SetInputOutputType(v int32) *UpdatePolicyRequestHarmfulCategoryConfigInfoList {
	s.InputOutputType = &v
	return s
}

func (s *UpdatePolicyRequestHarmfulCategoryConfigInfoList) SetIsEnabled(v int32) *UpdatePolicyRequestHarmfulCategoryConfigInfoList {
	s.IsEnabled = &v
	return s
}

func (s *UpdatePolicyRequestHarmfulCategoryConfigInfoList) SetSecurityLevel(v int32) *UpdatePolicyRequestHarmfulCategoryConfigInfoList {
	s.SecurityLevel = &v
	return s
}

func (s *UpdatePolicyRequestHarmfulCategoryConfigInfoList) Validate() error {
	return dara.Validate(s)
}

type UpdatePolicyRequestPromptAttackInfo struct {
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
	// 2
	SecurityLevel *int32 `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
}

func (s UpdatePolicyRequestPromptAttackInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyRequestPromptAttackInfo) GoString() string {
	return s.String()
}

func (s *UpdatePolicyRequestPromptAttackInfo) GetIsEnabled() *int32 {
	return s.IsEnabled
}

func (s *UpdatePolicyRequestPromptAttackInfo) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *UpdatePolicyRequestPromptAttackInfo) SetIsEnabled(v int32) *UpdatePolicyRequestPromptAttackInfo {
	s.IsEnabled = &v
	return s
}

func (s *UpdatePolicyRequestPromptAttackInfo) SetSecurityLevel(v int32) *UpdatePolicyRequestPromptAttackInfo {
	s.SecurityLevel = &v
	return s
}

func (s *UpdatePolicyRequestPromptAttackInfo) Validate() error {
	return dara.Validate(s)
}

type UpdatePolicyRequestPromptAttackInfoList struct {
	// Harmful category configuration ID
	//
	// example:
	//
	// 1
	CategoryConfigId *int64 `json:"CategoryConfigId,omitempty" xml:"CategoryConfigId,omitempty"`
	CategoryId       *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
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

func (s UpdatePolicyRequestPromptAttackInfoList) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyRequestPromptAttackInfoList) GoString() string {
	return s.String()
}

func (s *UpdatePolicyRequestPromptAttackInfoList) GetCategoryConfigId() *int64 {
	return s.CategoryConfigId
}

func (s *UpdatePolicyRequestPromptAttackInfoList) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *UpdatePolicyRequestPromptAttackInfoList) GetCategoryLabel() *string {
	return s.CategoryLabel
}

func (s *UpdatePolicyRequestPromptAttackInfoList) GetIsEnabled() *int32 {
	return s.IsEnabled
}

func (s *UpdatePolicyRequestPromptAttackInfoList) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *UpdatePolicyRequestPromptAttackInfoList) SetCategoryConfigId(v int64) *UpdatePolicyRequestPromptAttackInfoList {
	s.CategoryConfigId = &v
	return s
}

func (s *UpdatePolicyRequestPromptAttackInfoList) SetCategoryId(v int64) *UpdatePolicyRequestPromptAttackInfoList {
	s.CategoryId = &v
	return s
}

func (s *UpdatePolicyRequestPromptAttackInfoList) SetCategoryLabel(v string) *UpdatePolicyRequestPromptAttackInfoList {
	s.CategoryLabel = &v
	return s
}

func (s *UpdatePolicyRequestPromptAttackInfoList) SetIsEnabled(v int32) *UpdatePolicyRequestPromptAttackInfoList {
	s.IsEnabled = &v
	return s
}

func (s *UpdatePolicyRequestPromptAttackInfoList) SetSecurityLevel(v int32) *UpdatePolicyRequestPromptAttackInfoList {
	s.SecurityLevel = &v
	return s
}

func (s *UpdatePolicyRequestPromptAttackInfoList) Validate() error {
	return dara.Validate(s)
}

type UpdatePolicyRequestRegularExpressList struct {
	ActionType         *int32  `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	RegularExpress     *string `json:"RegularExpress,omitempty" xml:"RegularExpress,omitempty"`
	RegularExpressName *string `json:"RegularExpressName,omitempty" xml:"RegularExpressName,omitempty"`
}

func (s UpdatePolicyRequestRegularExpressList) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyRequestRegularExpressList) GoString() string {
	return s.String()
}

func (s *UpdatePolicyRequestRegularExpressList) GetActionType() *int32 {
	return s.ActionType
}

func (s *UpdatePolicyRequestRegularExpressList) GetRegularExpress() *string {
	return s.RegularExpress
}

func (s *UpdatePolicyRequestRegularExpressList) GetRegularExpressName() *string {
	return s.RegularExpressName
}

func (s *UpdatePolicyRequestRegularExpressList) SetActionType(v int32) *UpdatePolicyRequestRegularExpressList {
	s.ActionType = &v
	return s
}

func (s *UpdatePolicyRequestRegularExpressList) SetRegularExpress(v string) *UpdatePolicyRequestRegularExpressList {
	s.RegularExpress = &v
	return s
}

func (s *UpdatePolicyRequestRegularExpressList) SetRegularExpressName(v string) *UpdatePolicyRequestRegularExpressList {
	s.RegularExpressName = &v
	return s
}

func (s *UpdatePolicyRequestRegularExpressList) Validate() error {
	return dara.Validate(s)
}

type UpdatePolicyRequestSensitiveConfigList struct {
	ActionType        *int32 `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	IsEnabled         *int32 `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
	SensitiveConfigId *int64 `json:"SensitiveConfigId,omitempty" xml:"SensitiveConfigId,omitempty"`
}

func (s UpdatePolicyRequestSensitiveConfigList) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyRequestSensitiveConfigList) GoString() string {
	return s.String()
}

func (s *UpdatePolicyRequestSensitiveConfigList) GetActionType() *int32 {
	return s.ActionType
}

func (s *UpdatePolicyRequestSensitiveConfigList) GetIsEnabled() *int32 {
	return s.IsEnabled
}

func (s *UpdatePolicyRequestSensitiveConfigList) GetSensitiveConfigId() *int64 {
	return s.SensitiveConfigId
}

func (s *UpdatePolicyRequestSensitiveConfigList) SetActionType(v int32) *UpdatePolicyRequestSensitiveConfigList {
	s.ActionType = &v
	return s
}

func (s *UpdatePolicyRequestSensitiveConfigList) SetIsEnabled(v int32) *UpdatePolicyRequestSensitiveConfigList {
	s.IsEnabled = &v
	return s
}

func (s *UpdatePolicyRequestSensitiveConfigList) SetSensitiveConfigId(v int64) *UpdatePolicyRequestSensitiveConfigList {
	s.SensitiveConfigId = &v
	return s
}

func (s *UpdatePolicyRequestSensitiveConfigList) Validate() error {
	return dara.Validate(s)
}

type UpdatePolicyRequestSensitiveTopicList struct {
	CategoryType         *int32                                                       `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	SecurityLevel        *int32                                                       `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	TopicDefinition      *string                                                      `json:"TopicDefinition,omitempty" xml:"TopicDefinition,omitempty"`
	TopicExampleInfoList []*UpdatePolicyRequestSensitiveTopicListTopicExampleInfoList `json:"TopicExampleInfoList,omitempty" xml:"TopicExampleInfoList,omitempty" type:"Repeated"`
	TopicName            *string                                                      `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s UpdatePolicyRequestSensitiveTopicList) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyRequestSensitiveTopicList) GoString() string {
	return s.String()
}

func (s *UpdatePolicyRequestSensitiveTopicList) GetCategoryType() *int32 {
	return s.CategoryType
}

func (s *UpdatePolicyRequestSensitiveTopicList) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *UpdatePolicyRequestSensitiveTopicList) GetTopicDefinition() *string {
	return s.TopicDefinition
}

func (s *UpdatePolicyRequestSensitiveTopicList) GetTopicExampleInfoList() []*UpdatePolicyRequestSensitiveTopicListTopicExampleInfoList {
	return s.TopicExampleInfoList
}

func (s *UpdatePolicyRequestSensitiveTopicList) GetTopicName() *string {
	return s.TopicName
}

func (s *UpdatePolicyRequestSensitiveTopicList) SetCategoryType(v int32) *UpdatePolicyRequestSensitiveTopicList {
	s.CategoryType = &v
	return s
}

func (s *UpdatePolicyRequestSensitiveTopicList) SetSecurityLevel(v int32) *UpdatePolicyRequestSensitiveTopicList {
	s.SecurityLevel = &v
	return s
}

func (s *UpdatePolicyRequestSensitiveTopicList) SetTopicDefinition(v string) *UpdatePolicyRequestSensitiveTopicList {
	s.TopicDefinition = &v
	return s
}

func (s *UpdatePolicyRequestSensitiveTopicList) SetTopicExampleInfoList(v []*UpdatePolicyRequestSensitiveTopicListTopicExampleInfoList) *UpdatePolicyRequestSensitiveTopicList {
	s.TopicExampleInfoList = v
	return s
}

func (s *UpdatePolicyRequestSensitiveTopicList) SetTopicName(v string) *UpdatePolicyRequestSensitiveTopicList {
	s.TopicName = &v
	return s
}

func (s *UpdatePolicyRequestSensitiveTopicList) Validate() error {
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

type UpdatePolicyRequestSensitiveTopicListTopicExampleInfoList struct {
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ExampleType *int32  `json:"ExampleType,omitempty" xml:"ExampleType,omitempty"`
}

func (s UpdatePolicyRequestSensitiveTopicListTopicExampleInfoList) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyRequestSensitiveTopicListTopicExampleInfoList) GoString() string {
	return s.String()
}

func (s *UpdatePolicyRequestSensitiveTopicListTopicExampleInfoList) GetContent() *string {
	return s.Content
}

func (s *UpdatePolicyRequestSensitiveTopicListTopicExampleInfoList) GetExampleType() *int32 {
	return s.ExampleType
}

func (s *UpdatePolicyRequestSensitiveTopicListTopicExampleInfoList) SetContent(v string) *UpdatePolicyRequestSensitiveTopicListTopicExampleInfoList {
	s.Content = &v
	return s
}

func (s *UpdatePolicyRequestSensitiveTopicListTopicExampleInfoList) SetExampleType(v int32) *UpdatePolicyRequestSensitiveTopicListTopicExampleInfoList {
	s.ExampleType = &v
	return s
}

func (s *UpdatePolicyRequestSensitiveTopicListTopicExampleInfoList) Validate() error {
	return dara.Validate(s)
}

type UpdatePolicyRequestSensitiveWordList struct {
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Word  *string `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s UpdatePolicyRequestSensitiveWordList) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyRequestSensitiveWordList) GoString() string {
	return s.String()
}

func (s *UpdatePolicyRequestSensitiveWordList) GetLabel() *string {
	return s.Label
}

func (s *UpdatePolicyRequestSensitiveWordList) GetWord() *string {
	return s.Word
}

func (s *UpdatePolicyRequestSensitiveWordList) SetLabel(v string) *UpdatePolicyRequestSensitiveWordList {
	s.Label = &v
	return s
}

func (s *UpdatePolicyRequestSensitiveWordList) SetWord(v string) *UpdatePolicyRequestSensitiveWordList {
	s.Word = &v
	return s
}

func (s *UpdatePolicyRequestSensitiveWordList) Validate() error {
	return dara.Validate(s)
}

type UpdatePolicyRequestTopicConfigInfoList struct {
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
	// 3
	TopicId *int64 `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
	// Topic name
	//
	// example:
	//
	// Buss.
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s UpdatePolicyRequestTopicConfigInfoList) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyRequestTopicConfigInfoList) GoString() string {
	return s.String()
}

func (s *UpdatePolicyRequestTopicConfigInfoList) GetCategoryType() *int32 {
	return s.CategoryType
}

func (s *UpdatePolicyRequestTopicConfigInfoList) GetInputOutputType() *int32 {
	return s.InputOutputType
}

func (s *UpdatePolicyRequestTopicConfigInfoList) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *UpdatePolicyRequestTopicConfigInfoList) GetTopicId() *int64 {
	return s.TopicId
}

func (s *UpdatePolicyRequestTopicConfigInfoList) GetTopicName() *string {
	return s.TopicName
}

func (s *UpdatePolicyRequestTopicConfigInfoList) SetCategoryType(v int32) *UpdatePolicyRequestTopicConfigInfoList {
	s.CategoryType = &v
	return s
}

func (s *UpdatePolicyRequestTopicConfigInfoList) SetInputOutputType(v int32) *UpdatePolicyRequestTopicConfigInfoList {
	s.InputOutputType = &v
	return s
}

func (s *UpdatePolicyRequestTopicConfigInfoList) SetSecurityLevel(v int32) *UpdatePolicyRequestTopicConfigInfoList {
	s.SecurityLevel = &v
	return s
}

func (s *UpdatePolicyRequestTopicConfigInfoList) SetTopicId(v int64) *UpdatePolicyRequestTopicConfigInfoList {
	s.TopicId = &v
	return s
}

func (s *UpdatePolicyRequestTopicConfigInfoList) SetTopicName(v string) *UpdatePolicyRequestTopicConfigInfoList {
	s.TopicName = &v
	return s
}

func (s *UpdatePolicyRequestTopicConfigInfoList) Validate() error {
	return dara.Validate(s)
}

type UpdatePolicyRequestWordGroupInfoList struct {
	// Keyword group ID.
	//
	// example:
	//
	// 5fa285a60a9342e097cb7a3491ec00cc
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

func (s UpdatePolicyRequestWordGroupInfoList) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyRequestWordGroupInfoList) GoString() string {
	return s.String()
}

func (s *UpdatePolicyRequestWordGroupInfoList) GetGroupId() *int64 {
	return s.GroupId
}

func (s *UpdatePolicyRequestWordGroupInfoList) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdatePolicyRequestWordGroupInfoList) GetInputOutputType() *int32 {
	return s.InputOutputType
}

func (s *UpdatePolicyRequestWordGroupInfoList) SetGroupId(v int64) *UpdatePolicyRequestWordGroupInfoList {
	s.GroupId = &v
	return s
}

func (s *UpdatePolicyRequestWordGroupInfoList) SetGroupName(v string) *UpdatePolicyRequestWordGroupInfoList {
	s.GroupName = &v
	return s
}

func (s *UpdatePolicyRequestWordGroupInfoList) SetInputOutputType(v int32) *UpdatePolicyRequestWordGroupInfoList {
	s.InputOutputType = &v
	return s
}

func (s *UpdatePolicyRequestWordGroupInfoList) Validate() error {
	return dara.Validate(s)
}
