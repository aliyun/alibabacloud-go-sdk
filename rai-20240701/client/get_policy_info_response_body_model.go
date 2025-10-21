// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolicyInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPolicyInfoResponseBody
	GetCode() *string
	SetContentSafeModelInfo(v *GetPolicyInfoResponseBodyContentSafeModelInfo) *GetPolicyInfoResponseBody
	GetContentSafeModelInfo() *GetPolicyInfoResponseBodyContentSafeModelInfo
	SetEnableSensitiveInputCheck(v int32) *GetPolicyInfoResponseBody
	GetEnableSensitiveInputCheck() *int32
	SetEnableSensitiveOutputCheck(v int32) *GetPolicyInfoResponseBody
	GetEnableSensitiveOutputCheck() *int32
	SetGmtModified(v int64) *GetPolicyInfoResponseBody
	GetGmtModified() *int64
	SetHarmfulCategoryConfigInfoList(v []*GetPolicyInfoResponseBodyHarmfulCategoryConfigInfoList) *GetPolicyInfoResponseBody
	GetHarmfulCategoryConfigInfoList() []*GetPolicyInfoResponseBodyHarmfulCategoryConfigInfoList
	SetHttpStatusCode(v int32) *GetPolicyInfoResponseBody
	GetHttpStatusCode() *int32
	SetInputSafeAnswer(v string) *GetPolicyInfoResponseBody
	GetInputSafeAnswer() *string
	SetInputSafeAnswerSwitch(v int32) *GetPolicyInfoResponseBody
	GetInputSafeAnswerSwitch() *int32
	SetIsSidecarPolicy(v int32) *GetPolicyInfoResponseBody
	GetIsSidecarPolicy() *int32
	SetMessage(v string) *GetPolicyInfoResponseBody
	GetMessage() *string
	SetOutputSafeAnswer(v string) *GetPolicyInfoResponseBody
	GetOutputSafeAnswer() *string
	SetOutputSafeAnswerSwitch(v int32) *GetPolicyInfoResponseBody
	GetOutputSafeAnswerSwitch() *int32
	SetPolicyIdentifier(v string) *GetPolicyInfoResponseBody
	GetPolicyIdentifier() *string
	SetPolicyName(v string) *GetPolicyInfoResponseBody
	GetPolicyName() *string
	SetPromptAttackInfo(v *GetPolicyInfoResponseBodyPromptAttackInfo) *GetPolicyInfoResponseBody
	GetPromptAttackInfo() *GetPolicyInfoResponseBodyPromptAttackInfo
	SetPromptAttackInfoList(v []*GetPolicyInfoResponseBodyPromptAttackInfoList) *GetPolicyInfoResponseBody
	GetPromptAttackInfoList() []*GetPolicyInfoResponseBodyPromptAttackInfoList
	SetPromptAttackModelInfo(v *GetPolicyInfoResponseBodyPromptAttackModelInfo) *GetPolicyInfoResponseBody
	GetPromptAttackModelInfo() *GetPolicyInfoResponseBodyPromptAttackModelInfo
	SetRegularExpressList(v []*GetPolicyInfoResponseBodyRegularExpressList) *GetPolicyInfoResponseBody
	GetRegularExpressList() []*GetPolicyInfoResponseBodyRegularExpressList
	SetRequestId(v string) *GetPolicyInfoResponseBody
	GetRequestId() *string
	SetSceneType(v int32) *GetPolicyInfoResponseBody
	GetSceneType() *int32
	SetSensitiveConfigList(v []*GetPolicyInfoResponseBodySensitiveConfigList) *GetPolicyInfoResponseBody
	GetSensitiveConfigList() []*GetPolicyInfoResponseBodySensitiveConfigList
	SetSensitiveTopicList(v []*GetPolicyInfoResponseBodySensitiveTopicList) *GetPolicyInfoResponseBody
	GetSensitiveTopicList() []*GetPolicyInfoResponseBodySensitiveTopicList
	SetSensitiveTopicModelInfo(v *GetPolicyInfoResponseBodySensitiveTopicModelInfo) *GetPolicyInfoResponseBody
	GetSensitiveTopicModelInfo() *GetPolicyInfoResponseBodySensitiveTopicModelInfo
	SetSensitiveWordList(v []*GetPolicyInfoResponseBodySensitiveWordList) *GetPolicyInfoResponseBody
	GetSensitiveWordList() []*GetPolicyInfoResponseBodySensitiveWordList
	SetSuccess(v bool) *GetPolicyInfoResponseBody
	GetSuccess() *bool
	SetTopicConfigInfoList(v []*GetPolicyInfoResponseBodyTopicConfigInfoList) *GetPolicyInfoResponseBody
	GetTopicConfigInfoList() []*GetPolicyInfoResponseBodyTopicConfigInfoList
	SetWordGroupInfoList(v []*GetPolicyInfoResponseBodyWordGroupInfoList) *GetPolicyInfoResponseBody
	GetWordGroupInfoList() []*GetPolicyInfoResponseBodyWordGroupInfoList
}

type GetPolicyInfoResponseBody struct {
	// Result code, 00000 indicates success; others indicate failure.
	//
	// example:
	//
	// 00000
	Code                       *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	ContentSafeModelInfo       *GetPolicyInfoResponseBodyContentSafeModelInfo `json:"ContentSafeModelInfo,omitempty" xml:"ContentSafeModelInfo,omitempty" type:"Struct"`
	EnableSensitiveInputCheck  *int32                                         `json:"EnableSensitiveInputCheck,omitempty" xml:"EnableSensitiveInputCheck,omitempty"`
	EnableSensitiveOutputCheck *int32                                         `json:"EnableSensitiveOutputCheck,omitempty" xml:"EnableSensitiveOutputCheck,omitempty"`
	// Policy modification time
	//
	// example:
	//
	// 1634122349000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Harmful category configuration list
	HarmfulCategoryConfigInfoList []*GetPolicyInfoResponseBodyHarmfulCategoryConfigInfoList `json:"HarmfulCategoryConfigInfoList,omitempty" xml:"HarmfulCategoryConfigInfoList,omitempty" type:"Repeated"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode        *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	InputSafeAnswer       *string `json:"InputSafeAnswer,omitempty" xml:"InputSafeAnswer,omitempty"`
	InputSafeAnswerSwitch *int32  `json:"InputSafeAnswerSwitch,omitempty" xml:"InputSafeAnswerSwitch,omitempty"`
	IsSidecarPolicy       *int32  `json:"IsSidecarPolicy,omitempty" xml:"IsSidecarPolicy,omitempty"`
	// Error message.
	//
	// example:
	//
	// ""
	Message                *string `json:"Message,omitempty" xml:"Message,omitempty"`
	OutputSafeAnswer       *string `json:"OutputSafeAnswer,omitempty" xml:"OutputSafeAnswer,omitempty"`
	OutputSafeAnswerSwitch *int32  `json:"OutputSafeAnswerSwitch,omitempty" xml:"OutputSafeAnswerSwitch,omitempty"`
	// Policy identifier
	//
	// example:
	//
	// x1bc5xgs4uhx
	PolicyIdentifier *string `json:"PolicyIdentifier,omitempty" xml:"PolicyIdentifier,omitempty"`
	// Detection policy name.
	//
	// example:
	//
	// testPolicy
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// Prompt attack detection result object
	PromptAttackInfo *GetPolicyInfoResponseBodyPromptAttackInfo `json:"PromptAttackInfo,omitempty" xml:"PromptAttackInfo,omitempty" type:"Struct"`
	// Prompt attack list
	PromptAttackInfoList  []*GetPolicyInfoResponseBodyPromptAttackInfoList `json:"PromptAttackInfoList,omitempty" xml:"PromptAttackInfoList,omitempty" type:"Repeated"`
	PromptAttackModelInfo *GetPolicyInfoResponseBodyPromptAttackModelInfo  `json:"PromptAttackModelInfo,omitempty" xml:"PromptAttackModelInfo,omitempty" type:"Struct"`
	RegularExpressList    []*GetPolicyInfoResponseBodyRegularExpressList   `json:"RegularExpressList,omitempty" xml:"RegularExpressList,omitempty" type:"Repeated"`
	// Request ID
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId               *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SceneType               *int32                                            `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
	SensitiveConfigList     []*GetPolicyInfoResponseBodySensitiveConfigList   `json:"SensitiveConfigList,omitempty" xml:"SensitiveConfigList,omitempty" type:"Repeated"`
	SensitiveTopicList      []*GetPolicyInfoResponseBodySensitiveTopicList    `json:"SensitiveTopicList,omitempty" xml:"SensitiveTopicList,omitempty" type:"Repeated"`
	SensitiveTopicModelInfo *GetPolicyInfoResponseBodySensitiveTopicModelInfo `json:"SensitiveTopicModelInfo,omitempty" xml:"SensitiveTopicModelInfo,omitempty" type:"Struct"`
	SensitiveWordList       []*GetPolicyInfoResponseBodySensitiveWordList     `json:"SensitiveWordList,omitempty" xml:"SensitiveWordList,omitempty" type:"Repeated"`
	// Indicates whether the operation was successful. `true` for success, `false` for failure.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Sensitive topic list
	TopicConfigInfoList []*GetPolicyInfoResponseBodyTopicConfigInfoList `json:"TopicConfigInfoList,omitempty" xml:"TopicConfigInfoList,omitempty" type:"Repeated"`
	// Keyword group object list
	WordGroupInfoList []*GetPolicyInfoResponseBodyWordGroupInfoList `json:"WordGroupInfoList,omitempty" xml:"WordGroupInfoList,omitempty" type:"Repeated"`
}

func (s GetPolicyInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetPolicyInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPolicyInfoResponseBody) GetContentSafeModelInfo() *GetPolicyInfoResponseBodyContentSafeModelInfo {
	return s.ContentSafeModelInfo
}

func (s *GetPolicyInfoResponseBody) GetEnableSensitiveInputCheck() *int32 {
	return s.EnableSensitiveInputCheck
}

func (s *GetPolicyInfoResponseBody) GetEnableSensitiveOutputCheck() *int32 {
	return s.EnableSensitiveOutputCheck
}

func (s *GetPolicyInfoResponseBody) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *GetPolicyInfoResponseBody) GetHarmfulCategoryConfigInfoList() []*GetPolicyInfoResponseBodyHarmfulCategoryConfigInfoList {
	return s.HarmfulCategoryConfigInfoList
}

func (s *GetPolicyInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetPolicyInfoResponseBody) GetInputSafeAnswer() *string {
	return s.InputSafeAnswer
}

func (s *GetPolicyInfoResponseBody) GetInputSafeAnswerSwitch() *int32 {
	return s.InputSafeAnswerSwitch
}

func (s *GetPolicyInfoResponseBody) GetIsSidecarPolicy() *int32 {
	return s.IsSidecarPolicy
}

func (s *GetPolicyInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPolicyInfoResponseBody) GetOutputSafeAnswer() *string {
	return s.OutputSafeAnswer
}

func (s *GetPolicyInfoResponseBody) GetOutputSafeAnswerSwitch() *int32 {
	return s.OutputSafeAnswerSwitch
}

func (s *GetPolicyInfoResponseBody) GetPolicyIdentifier() *string {
	return s.PolicyIdentifier
}

func (s *GetPolicyInfoResponseBody) GetPolicyName() *string {
	return s.PolicyName
}

func (s *GetPolicyInfoResponseBody) GetPromptAttackInfo() *GetPolicyInfoResponseBodyPromptAttackInfo {
	return s.PromptAttackInfo
}

func (s *GetPolicyInfoResponseBody) GetPromptAttackInfoList() []*GetPolicyInfoResponseBodyPromptAttackInfoList {
	return s.PromptAttackInfoList
}

func (s *GetPolicyInfoResponseBody) GetPromptAttackModelInfo() *GetPolicyInfoResponseBodyPromptAttackModelInfo {
	return s.PromptAttackModelInfo
}

func (s *GetPolicyInfoResponseBody) GetRegularExpressList() []*GetPolicyInfoResponseBodyRegularExpressList {
	return s.RegularExpressList
}

func (s *GetPolicyInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPolicyInfoResponseBody) GetSceneType() *int32 {
	return s.SceneType
}

func (s *GetPolicyInfoResponseBody) GetSensitiveConfigList() []*GetPolicyInfoResponseBodySensitiveConfigList {
	return s.SensitiveConfigList
}

func (s *GetPolicyInfoResponseBody) GetSensitiveTopicList() []*GetPolicyInfoResponseBodySensitiveTopicList {
	return s.SensitiveTopicList
}

func (s *GetPolicyInfoResponseBody) GetSensitiveTopicModelInfo() *GetPolicyInfoResponseBodySensitiveTopicModelInfo {
	return s.SensitiveTopicModelInfo
}

func (s *GetPolicyInfoResponseBody) GetSensitiveWordList() []*GetPolicyInfoResponseBodySensitiveWordList {
	return s.SensitiveWordList
}

func (s *GetPolicyInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPolicyInfoResponseBody) GetTopicConfigInfoList() []*GetPolicyInfoResponseBodyTopicConfigInfoList {
	return s.TopicConfigInfoList
}

func (s *GetPolicyInfoResponseBody) GetWordGroupInfoList() []*GetPolicyInfoResponseBodyWordGroupInfoList {
	return s.WordGroupInfoList
}

func (s *GetPolicyInfoResponseBody) SetCode(v string) *GetPolicyInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetPolicyInfoResponseBody) SetContentSafeModelInfo(v *GetPolicyInfoResponseBodyContentSafeModelInfo) *GetPolicyInfoResponseBody {
	s.ContentSafeModelInfo = v
	return s
}

func (s *GetPolicyInfoResponseBody) SetEnableSensitiveInputCheck(v int32) *GetPolicyInfoResponseBody {
	s.EnableSensitiveInputCheck = &v
	return s
}

func (s *GetPolicyInfoResponseBody) SetEnableSensitiveOutputCheck(v int32) *GetPolicyInfoResponseBody {
	s.EnableSensitiveOutputCheck = &v
	return s
}

func (s *GetPolicyInfoResponseBody) SetGmtModified(v int64) *GetPolicyInfoResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetPolicyInfoResponseBody) SetHarmfulCategoryConfigInfoList(v []*GetPolicyInfoResponseBodyHarmfulCategoryConfigInfoList) *GetPolicyInfoResponseBody {
	s.HarmfulCategoryConfigInfoList = v
	return s
}

func (s *GetPolicyInfoResponseBody) SetHttpStatusCode(v int32) *GetPolicyInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPolicyInfoResponseBody) SetInputSafeAnswer(v string) *GetPolicyInfoResponseBody {
	s.InputSafeAnswer = &v
	return s
}

func (s *GetPolicyInfoResponseBody) SetInputSafeAnswerSwitch(v int32) *GetPolicyInfoResponseBody {
	s.InputSafeAnswerSwitch = &v
	return s
}

func (s *GetPolicyInfoResponseBody) SetIsSidecarPolicy(v int32) *GetPolicyInfoResponseBody {
	s.IsSidecarPolicy = &v
	return s
}

func (s *GetPolicyInfoResponseBody) SetMessage(v string) *GetPolicyInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetPolicyInfoResponseBody) SetOutputSafeAnswer(v string) *GetPolicyInfoResponseBody {
	s.OutputSafeAnswer = &v
	return s
}

func (s *GetPolicyInfoResponseBody) SetOutputSafeAnswerSwitch(v int32) *GetPolicyInfoResponseBody {
	s.OutputSafeAnswerSwitch = &v
	return s
}

func (s *GetPolicyInfoResponseBody) SetPolicyIdentifier(v string) *GetPolicyInfoResponseBody {
	s.PolicyIdentifier = &v
	return s
}

func (s *GetPolicyInfoResponseBody) SetPolicyName(v string) *GetPolicyInfoResponseBody {
	s.PolicyName = &v
	return s
}

func (s *GetPolicyInfoResponseBody) SetPromptAttackInfo(v *GetPolicyInfoResponseBodyPromptAttackInfo) *GetPolicyInfoResponseBody {
	s.PromptAttackInfo = v
	return s
}

func (s *GetPolicyInfoResponseBody) SetPromptAttackInfoList(v []*GetPolicyInfoResponseBodyPromptAttackInfoList) *GetPolicyInfoResponseBody {
	s.PromptAttackInfoList = v
	return s
}

func (s *GetPolicyInfoResponseBody) SetPromptAttackModelInfo(v *GetPolicyInfoResponseBodyPromptAttackModelInfo) *GetPolicyInfoResponseBody {
	s.PromptAttackModelInfo = v
	return s
}

func (s *GetPolicyInfoResponseBody) SetRegularExpressList(v []*GetPolicyInfoResponseBodyRegularExpressList) *GetPolicyInfoResponseBody {
	s.RegularExpressList = v
	return s
}

func (s *GetPolicyInfoResponseBody) SetRequestId(v string) *GetPolicyInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPolicyInfoResponseBody) SetSceneType(v int32) *GetPolicyInfoResponseBody {
	s.SceneType = &v
	return s
}

func (s *GetPolicyInfoResponseBody) SetSensitiveConfigList(v []*GetPolicyInfoResponseBodySensitiveConfigList) *GetPolicyInfoResponseBody {
	s.SensitiveConfigList = v
	return s
}

func (s *GetPolicyInfoResponseBody) SetSensitiveTopicList(v []*GetPolicyInfoResponseBodySensitiveTopicList) *GetPolicyInfoResponseBody {
	s.SensitiveTopicList = v
	return s
}

func (s *GetPolicyInfoResponseBody) SetSensitiveTopicModelInfo(v *GetPolicyInfoResponseBodySensitiveTopicModelInfo) *GetPolicyInfoResponseBody {
	s.SensitiveTopicModelInfo = v
	return s
}

func (s *GetPolicyInfoResponseBody) SetSensitiveWordList(v []*GetPolicyInfoResponseBodySensitiveWordList) *GetPolicyInfoResponseBody {
	s.SensitiveWordList = v
	return s
}

func (s *GetPolicyInfoResponseBody) SetSuccess(v bool) *GetPolicyInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetPolicyInfoResponseBody) SetTopicConfigInfoList(v []*GetPolicyInfoResponseBodyTopicConfigInfoList) *GetPolicyInfoResponseBody {
	s.TopicConfigInfoList = v
	return s
}

func (s *GetPolicyInfoResponseBody) SetWordGroupInfoList(v []*GetPolicyInfoResponseBodyWordGroupInfoList) *GetPolicyInfoResponseBody {
	s.WordGroupInfoList = v
	return s
}

func (s *GetPolicyInfoResponseBody) Validate() error {
	if s.ContentSafeModelInfo != nil {
		if err := s.ContentSafeModelInfo.Validate(); err != nil {
			return err
		}
	}
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
	if s.PromptAttackModelInfo != nil {
		if err := s.PromptAttackModelInfo.Validate(); err != nil {
			return err
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
	if s.SensitiveTopicModelInfo != nil {
		if err := s.SensitiveTopicModelInfo.Validate(); err != nil {
			return err
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

type GetPolicyInfoResponseBodyContentSafeModelInfo struct {
	EasServiceName  *string `json:"EasServiceName,omitempty" xml:"EasServiceName,omitempty"`
	ModelInstanceId *int64  `json:"ModelInstanceId,omitempty" xml:"ModelInstanceId,omitempty"`
}

func (s GetPolicyInfoResponseBodyContentSafeModelInfo) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyInfoResponseBodyContentSafeModelInfo) GoString() string {
	return s.String()
}

func (s *GetPolicyInfoResponseBodyContentSafeModelInfo) GetEasServiceName() *string {
	return s.EasServiceName
}

func (s *GetPolicyInfoResponseBodyContentSafeModelInfo) GetModelInstanceId() *int64 {
	return s.ModelInstanceId
}

func (s *GetPolicyInfoResponseBodyContentSafeModelInfo) SetEasServiceName(v string) *GetPolicyInfoResponseBodyContentSafeModelInfo {
	s.EasServiceName = &v
	return s
}

func (s *GetPolicyInfoResponseBodyContentSafeModelInfo) SetModelInstanceId(v int64) *GetPolicyInfoResponseBodyContentSafeModelInfo {
	s.ModelInstanceId = &v
	return s
}

func (s *GetPolicyInfoResponseBodyContentSafeModelInfo) Validate() error {
	return dara.Validate(s)
}

type GetPolicyInfoResponseBodyHarmfulCategoryConfigInfoList struct {
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

func (s GetPolicyInfoResponseBodyHarmfulCategoryConfigInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyInfoResponseBodyHarmfulCategoryConfigInfoList) GoString() string {
	return s.String()
}

func (s *GetPolicyInfoResponseBodyHarmfulCategoryConfigInfoList) GetCategoryConfigId() *int64 {
	return s.CategoryConfigId
}

func (s *GetPolicyInfoResponseBodyHarmfulCategoryConfigInfoList) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *GetPolicyInfoResponseBodyHarmfulCategoryConfigInfoList) GetCategoryLabel() *string {
	return s.CategoryLabel
}

func (s *GetPolicyInfoResponseBodyHarmfulCategoryConfigInfoList) GetCategoryType() *int32 {
	return s.CategoryType
}

func (s *GetPolicyInfoResponseBodyHarmfulCategoryConfigInfoList) GetInputOutputType() *int32 {
	return s.InputOutputType
}

func (s *GetPolicyInfoResponseBodyHarmfulCategoryConfigInfoList) GetIsEnabled() *int32 {
	return s.IsEnabled
}

func (s *GetPolicyInfoResponseBodyHarmfulCategoryConfigInfoList) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *GetPolicyInfoResponseBodyHarmfulCategoryConfigInfoList) SetCategoryConfigId(v int64) *GetPolicyInfoResponseBodyHarmfulCategoryConfigInfoList {
	s.CategoryConfigId = &v
	return s
}

func (s *GetPolicyInfoResponseBodyHarmfulCategoryConfigInfoList) SetCategoryId(v int64) *GetPolicyInfoResponseBodyHarmfulCategoryConfigInfoList {
	s.CategoryId = &v
	return s
}

func (s *GetPolicyInfoResponseBodyHarmfulCategoryConfigInfoList) SetCategoryLabel(v string) *GetPolicyInfoResponseBodyHarmfulCategoryConfigInfoList {
	s.CategoryLabel = &v
	return s
}

func (s *GetPolicyInfoResponseBodyHarmfulCategoryConfigInfoList) SetCategoryType(v int32) *GetPolicyInfoResponseBodyHarmfulCategoryConfigInfoList {
	s.CategoryType = &v
	return s
}

func (s *GetPolicyInfoResponseBodyHarmfulCategoryConfigInfoList) SetInputOutputType(v int32) *GetPolicyInfoResponseBodyHarmfulCategoryConfigInfoList {
	s.InputOutputType = &v
	return s
}

func (s *GetPolicyInfoResponseBodyHarmfulCategoryConfigInfoList) SetIsEnabled(v int32) *GetPolicyInfoResponseBodyHarmfulCategoryConfigInfoList {
	s.IsEnabled = &v
	return s
}

func (s *GetPolicyInfoResponseBodyHarmfulCategoryConfigInfoList) SetSecurityLevel(v int32) *GetPolicyInfoResponseBodyHarmfulCategoryConfigInfoList {
	s.SecurityLevel = &v
	return s
}

func (s *GetPolicyInfoResponseBodyHarmfulCategoryConfigInfoList) Validate() error {
	return dara.Validate(s)
}

type GetPolicyInfoResponseBodyPromptAttackInfo struct {
	// Prompt attack configuration switch
	//
	// 0: Off
	//
	// 1: On
	//
	// example:
	//
	// true
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

func (s GetPolicyInfoResponseBodyPromptAttackInfo) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyInfoResponseBodyPromptAttackInfo) GoString() string {
	return s.String()
}

func (s *GetPolicyInfoResponseBodyPromptAttackInfo) GetIsEnabled() *int32 {
	return s.IsEnabled
}

func (s *GetPolicyInfoResponseBodyPromptAttackInfo) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *GetPolicyInfoResponseBodyPromptAttackInfo) SetIsEnabled(v int32) *GetPolicyInfoResponseBodyPromptAttackInfo {
	s.IsEnabled = &v
	return s
}

func (s *GetPolicyInfoResponseBodyPromptAttackInfo) SetSecurityLevel(v int32) *GetPolicyInfoResponseBodyPromptAttackInfo {
	s.SecurityLevel = &v
	return s
}

func (s *GetPolicyInfoResponseBodyPromptAttackInfo) Validate() error {
	return dara.Validate(s)
}

type GetPolicyInfoResponseBodyPromptAttackInfoList struct {
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

func (s GetPolicyInfoResponseBodyPromptAttackInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyInfoResponseBodyPromptAttackInfoList) GoString() string {
	return s.String()
}

func (s *GetPolicyInfoResponseBodyPromptAttackInfoList) GetCategoryConfigId() *int64 {
	return s.CategoryConfigId
}

func (s *GetPolicyInfoResponseBodyPromptAttackInfoList) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *GetPolicyInfoResponseBodyPromptAttackInfoList) GetCategoryLabel() *string {
	return s.CategoryLabel
}

func (s *GetPolicyInfoResponseBodyPromptAttackInfoList) GetIsEnabled() *int32 {
	return s.IsEnabled
}

func (s *GetPolicyInfoResponseBodyPromptAttackInfoList) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *GetPolicyInfoResponseBodyPromptAttackInfoList) SetCategoryConfigId(v int64) *GetPolicyInfoResponseBodyPromptAttackInfoList {
	s.CategoryConfigId = &v
	return s
}

func (s *GetPolicyInfoResponseBodyPromptAttackInfoList) SetCategoryId(v int64) *GetPolicyInfoResponseBodyPromptAttackInfoList {
	s.CategoryId = &v
	return s
}

func (s *GetPolicyInfoResponseBodyPromptAttackInfoList) SetCategoryLabel(v string) *GetPolicyInfoResponseBodyPromptAttackInfoList {
	s.CategoryLabel = &v
	return s
}

func (s *GetPolicyInfoResponseBodyPromptAttackInfoList) SetIsEnabled(v int32) *GetPolicyInfoResponseBodyPromptAttackInfoList {
	s.IsEnabled = &v
	return s
}

func (s *GetPolicyInfoResponseBodyPromptAttackInfoList) SetSecurityLevel(v int32) *GetPolicyInfoResponseBodyPromptAttackInfoList {
	s.SecurityLevel = &v
	return s
}

func (s *GetPolicyInfoResponseBodyPromptAttackInfoList) Validate() error {
	return dara.Validate(s)
}

type GetPolicyInfoResponseBodyPromptAttackModelInfo struct {
	EasServiceName  *string `json:"EasServiceName,omitempty" xml:"EasServiceName,omitempty"`
	ModelInstanceId *int64  `json:"ModelInstanceId,omitempty" xml:"ModelInstanceId,omitempty"`
}

func (s GetPolicyInfoResponseBodyPromptAttackModelInfo) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyInfoResponseBodyPromptAttackModelInfo) GoString() string {
	return s.String()
}

func (s *GetPolicyInfoResponseBodyPromptAttackModelInfo) GetEasServiceName() *string {
	return s.EasServiceName
}

func (s *GetPolicyInfoResponseBodyPromptAttackModelInfo) GetModelInstanceId() *int64 {
	return s.ModelInstanceId
}

func (s *GetPolicyInfoResponseBodyPromptAttackModelInfo) SetEasServiceName(v string) *GetPolicyInfoResponseBodyPromptAttackModelInfo {
	s.EasServiceName = &v
	return s
}

func (s *GetPolicyInfoResponseBodyPromptAttackModelInfo) SetModelInstanceId(v int64) *GetPolicyInfoResponseBodyPromptAttackModelInfo {
	s.ModelInstanceId = &v
	return s
}

func (s *GetPolicyInfoResponseBodyPromptAttackModelInfo) Validate() error {
	return dara.Validate(s)
}

type GetPolicyInfoResponseBodyRegularExpressList struct {
	ActionType         *int32  `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	InputOutputType    *int32  `json:"InputOutputType,omitempty" xml:"InputOutputType,omitempty"`
	IsEnabled          *int32  `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
	MatchAndReplace    *string `json:"MatchAndReplace,omitempty" xml:"MatchAndReplace,omitempty"`
	RegularExpress     *string `json:"RegularExpress,omitempty" xml:"RegularExpress,omitempty"`
	RegularExpressId   *int64  `json:"RegularExpressId,omitempty" xml:"RegularExpressId,omitempty"`
	RegularExpressName *string `json:"RegularExpressName,omitempty" xml:"RegularExpressName,omitempty"`
}

func (s GetPolicyInfoResponseBodyRegularExpressList) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyInfoResponseBodyRegularExpressList) GoString() string {
	return s.String()
}

func (s *GetPolicyInfoResponseBodyRegularExpressList) GetActionType() *int32 {
	return s.ActionType
}

func (s *GetPolicyInfoResponseBodyRegularExpressList) GetInputOutputType() *int32 {
	return s.InputOutputType
}

func (s *GetPolicyInfoResponseBodyRegularExpressList) GetIsEnabled() *int32 {
	return s.IsEnabled
}

func (s *GetPolicyInfoResponseBodyRegularExpressList) GetMatchAndReplace() *string {
	return s.MatchAndReplace
}

func (s *GetPolicyInfoResponseBodyRegularExpressList) GetRegularExpress() *string {
	return s.RegularExpress
}

func (s *GetPolicyInfoResponseBodyRegularExpressList) GetRegularExpressId() *int64 {
	return s.RegularExpressId
}

func (s *GetPolicyInfoResponseBodyRegularExpressList) GetRegularExpressName() *string {
	return s.RegularExpressName
}

func (s *GetPolicyInfoResponseBodyRegularExpressList) SetActionType(v int32) *GetPolicyInfoResponseBodyRegularExpressList {
	s.ActionType = &v
	return s
}

func (s *GetPolicyInfoResponseBodyRegularExpressList) SetInputOutputType(v int32) *GetPolicyInfoResponseBodyRegularExpressList {
	s.InputOutputType = &v
	return s
}

func (s *GetPolicyInfoResponseBodyRegularExpressList) SetIsEnabled(v int32) *GetPolicyInfoResponseBodyRegularExpressList {
	s.IsEnabled = &v
	return s
}

func (s *GetPolicyInfoResponseBodyRegularExpressList) SetMatchAndReplace(v string) *GetPolicyInfoResponseBodyRegularExpressList {
	s.MatchAndReplace = &v
	return s
}

func (s *GetPolicyInfoResponseBodyRegularExpressList) SetRegularExpress(v string) *GetPolicyInfoResponseBodyRegularExpressList {
	s.RegularExpress = &v
	return s
}

func (s *GetPolicyInfoResponseBodyRegularExpressList) SetRegularExpressId(v int64) *GetPolicyInfoResponseBodyRegularExpressList {
	s.RegularExpressId = &v
	return s
}

func (s *GetPolicyInfoResponseBodyRegularExpressList) SetRegularExpressName(v string) *GetPolicyInfoResponseBodyRegularExpressList {
	s.RegularExpressName = &v
	return s
}

func (s *GetPolicyInfoResponseBodyRegularExpressList) Validate() error {
	return dara.Validate(s)
}

type GetPolicyInfoResponseBodySensitiveConfigList struct {
	ActionType        *int32  `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	InputOutputType   *int32  `json:"InputOutputType,omitempty" xml:"InputOutputType,omitempty"`
	IsEnabled         *int32  `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
	MatchAndReplace   *string `json:"MatchAndReplace,omitempty" xml:"MatchAndReplace,omitempty"`
	SensitiveConfigId *int64  `json:"SensitiveConfigId,omitempty" xml:"SensitiveConfigId,omitempty"`
	SensitiveName     *string `json:"SensitiveName,omitempty" xml:"SensitiveName,omitempty"`
}

func (s GetPolicyInfoResponseBodySensitiveConfigList) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyInfoResponseBodySensitiveConfigList) GoString() string {
	return s.String()
}

func (s *GetPolicyInfoResponseBodySensitiveConfigList) GetActionType() *int32 {
	return s.ActionType
}

func (s *GetPolicyInfoResponseBodySensitiveConfigList) GetInputOutputType() *int32 {
	return s.InputOutputType
}

func (s *GetPolicyInfoResponseBodySensitiveConfigList) GetIsEnabled() *int32 {
	return s.IsEnabled
}

func (s *GetPolicyInfoResponseBodySensitiveConfigList) GetMatchAndReplace() *string {
	return s.MatchAndReplace
}

func (s *GetPolicyInfoResponseBodySensitiveConfigList) GetSensitiveConfigId() *int64 {
	return s.SensitiveConfigId
}

func (s *GetPolicyInfoResponseBodySensitiveConfigList) GetSensitiveName() *string {
	return s.SensitiveName
}

func (s *GetPolicyInfoResponseBodySensitiveConfigList) SetActionType(v int32) *GetPolicyInfoResponseBodySensitiveConfigList {
	s.ActionType = &v
	return s
}

func (s *GetPolicyInfoResponseBodySensitiveConfigList) SetInputOutputType(v int32) *GetPolicyInfoResponseBodySensitiveConfigList {
	s.InputOutputType = &v
	return s
}

func (s *GetPolicyInfoResponseBodySensitiveConfigList) SetIsEnabled(v int32) *GetPolicyInfoResponseBodySensitiveConfigList {
	s.IsEnabled = &v
	return s
}

func (s *GetPolicyInfoResponseBodySensitiveConfigList) SetMatchAndReplace(v string) *GetPolicyInfoResponseBodySensitiveConfigList {
	s.MatchAndReplace = &v
	return s
}

func (s *GetPolicyInfoResponseBodySensitiveConfigList) SetSensitiveConfigId(v int64) *GetPolicyInfoResponseBodySensitiveConfigList {
	s.SensitiveConfigId = &v
	return s
}

func (s *GetPolicyInfoResponseBodySensitiveConfigList) SetSensitiveName(v string) *GetPolicyInfoResponseBodySensitiveConfigList {
	s.SensitiveName = &v
	return s
}

func (s *GetPolicyInfoResponseBodySensitiveConfigList) Validate() error {
	return dara.Validate(s)
}

type GetPolicyInfoResponseBodySensitiveTopicList struct {
	CategoryType         *int32                                                             `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	InputOutputType      *int32                                                             `json:"InputOutputType,omitempty" xml:"InputOutputType,omitempty"`
	SecurityLevel        *int32                                                             `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	TopicDefinition      *string                                                            `json:"TopicDefinition,omitempty" xml:"TopicDefinition,omitempty"`
	TopicExampleInfoList []*GetPolicyInfoResponseBodySensitiveTopicListTopicExampleInfoList `json:"TopicExampleInfoList,omitempty" xml:"TopicExampleInfoList,omitempty" type:"Repeated"`
	TopicId              *int64                                                             `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
	TopicName            *string                                                            `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s GetPolicyInfoResponseBodySensitiveTopicList) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyInfoResponseBodySensitiveTopicList) GoString() string {
	return s.String()
}

func (s *GetPolicyInfoResponseBodySensitiveTopicList) GetCategoryType() *int32 {
	return s.CategoryType
}

func (s *GetPolicyInfoResponseBodySensitiveTopicList) GetInputOutputType() *int32 {
	return s.InputOutputType
}

func (s *GetPolicyInfoResponseBodySensitiveTopicList) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *GetPolicyInfoResponseBodySensitiveTopicList) GetTopicDefinition() *string {
	return s.TopicDefinition
}

func (s *GetPolicyInfoResponseBodySensitiveTopicList) GetTopicExampleInfoList() []*GetPolicyInfoResponseBodySensitiveTopicListTopicExampleInfoList {
	return s.TopicExampleInfoList
}

func (s *GetPolicyInfoResponseBodySensitiveTopicList) GetTopicId() *int64 {
	return s.TopicId
}

func (s *GetPolicyInfoResponseBodySensitiveTopicList) GetTopicName() *string {
	return s.TopicName
}

func (s *GetPolicyInfoResponseBodySensitiveTopicList) SetCategoryType(v int32) *GetPolicyInfoResponseBodySensitiveTopicList {
	s.CategoryType = &v
	return s
}

func (s *GetPolicyInfoResponseBodySensitiveTopicList) SetInputOutputType(v int32) *GetPolicyInfoResponseBodySensitiveTopicList {
	s.InputOutputType = &v
	return s
}

func (s *GetPolicyInfoResponseBodySensitiveTopicList) SetSecurityLevel(v int32) *GetPolicyInfoResponseBodySensitiveTopicList {
	s.SecurityLevel = &v
	return s
}

func (s *GetPolicyInfoResponseBodySensitiveTopicList) SetTopicDefinition(v string) *GetPolicyInfoResponseBodySensitiveTopicList {
	s.TopicDefinition = &v
	return s
}

func (s *GetPolicyInfoResponseBodySensitiveTopicList) SetTopicExampleInfoList(v []*GetPolicyInfoResponseBodySensitiveTopicListTopicExampleInfoList) *GetPolicyInfoResponseBodySensitiveTopicList {
	s.TopicExampleInfoList = v
	return s
}

func (s *GetPolicyInfoResponseBodySensitiveTopicList) SetTopicId(v int64) *GetPolicyInfoResponseBodySensitiveTopicList {
	s.TopicId = &v
	return s
}

func (s *GetPolicyInfoResponseBodySensitiveTopicList) SetTopicName(v string) *GetPolicyInfoResponseBodySensitiveTopicList {
	s.TopicName = &v
	return s
}

func (s *GetPolicyInfoResponseBodySensitiveTopicList) Validate() error {
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

type GetPolicyInfoResponseBodySensitiveTopicListTopicExampleInfoList struct {
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ExampleType *int64  `json:"ExampleType,omitempty" xml:"ExampleType,omitempty"`
}

func (s GetPolicyInfoResponseBodySensitiveTopicListTopicExampleInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyInfoResponseBodySensitiveTopicListTopicExampleInfoList) GoString() string {
	return s.String()
}

func (s *GetPolicyInfoResponseBodySensitiveTopicListTopicExampleInfoList) GetContent() *string {
	return s.Content
}

func (s *GetPolicyInfoResponseBodySensitiveTopicListTopicExampleInfoList) GetExampleType() *int64 {
	return s.ExampleType
}

func (s *GetPolicyInfoResponseBodySensitiveTopicListTopicExampleInfoList) SetContent(v string) *GetPolicyInfoResponseBodySensitiveTopicListTopicExampleInfoList {
	s.Content = &v
	return s
}

func (s *GetPolicyInfoResponseBodySensitiveTopicListTopicExampleInfoList) SetExampleType(v int64) *GetPolicyInfoResponseBodySensitiveTopicListTopicExampleInfoList {
	s.ExampleType = &v
	return s
}

func (s *GetPolicyInfoResponseBodySensitiveTopicListTopicExampleInfoList) Validate() error {
	return dara.Validate(s)
}

type GetPolicyInfoResponseBodySensitiveTopicModelInfo struct {
	EasServiceName  *string `json:"EasServiceName,omitempty" xml:"EasServiceName,omitempty"`
	ModelInstanceId *int64  `json:"ModelInstanceId,omitempty" xml:"ModelInstanceId,omitempty"`
}

func (s GetPolicyInfoResponseBodySensitiveTopicModelInfo) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyInfoResponseBodySensitiveTopicModelInfo) GoString() string {
	return s.String()
}

func (s *GetPolicyInfoResponseBodySensitiveTopicModelInfo) GetEasServiceName() *string {
	return s.EasServiceName
}

func (s *GetPolicyInfoResponseBodySensitiveTopicModelInfo) GetModelInstanceId() *int64 {
	return s.ModelInstanceId
}

func (s *GetPolicyInfoResponseBodySensitiveTopicModelInfo) SetEasServiceName(v string) *GetPolicyInfoResponseBodySensitiveTopicModelInfo {
	s.EasServiceName = &v
	return s
}

func (s *GetPolicyInfoResponseBodySensitiveTopicModelInfo) SetModelInstanceId(v int64) *GetPolicyInfoResponseBodySensitiveTopicModelInfo {
	s.ModelInstanceId = &v
	return s
}

func (s *GetPolicyInfoResponseBodySensitiveTopicModelInfo) Validate() error {
	return dara.Validate(s)
}

type GetPolicyInfoResponseBodySensitiveWordList struct {
	InputOutputType *int32  `json:"InputOutputType,omitempty" xml:"InputOutputType,omitempty"`
	Label           *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Word            *string `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s GetPolicyInfoResponseBodySensitiveWordList) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyInfoResponseBodySensitiveWordList) GoString() string {
	return s.String()
}

func (s *GetPolicyInfoResponseBodySensitiveWordList) GetInputOutputType() *int32 {
	return s.InputOutputType
}

func (s *GetPolicyInfoResponseBodySensitiveWordList) GetLabel() *string {
	return s.Label
}

func (s *GetPolicyInfoResponseBodySensitiveWordList) GetWord() *string {
	return s.Word
}

func (s *GetPolicyInfoResponseBodySensitiveWordList) SetInputOutputType(v int32) *GetPolicyInfoResponseBodySensitiveWordList {
	s.InputOutputType = &v
	return s
}

func (s *GetPolicyInfoResponseBodySensitiveWordList) SetLabel(v string) *GetPolicyInfoResponseBodySensitiveWordList {
	s.Label = &v
	return s
}

func (s *GetPolicyInfoResponseBodySensitiveWordList) SetWord(v string) *GetPolicyInfoResponseBodySensitiveWordList {
	s.Word = &v
	return s
}

func (s *GetPolicyInfoResponseBodySensitiveWordList) Validate() error {
	return dara.Validate(s)
}

type GetPolicyInfoResponseBodyTopicConfigInfoList struct {
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

func (s GetPolicyInfoResponseBodyTopicConfigInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyInfoResponseBodyTopicConfigInfoList) GoString() string {
	return s.String()
}

func (s *GetPolicyInfoResponseBodyTopicConfigInfoList) GetCategoryType() *int32 {
	return s.CategoryType
}

func (s *GetPolicyInfoResponseBodyTopicConfigInfoList) GetInputOutputType() *int32 {
	return s.InputOutputType
}

func (s *GetPolicyInfoResponseBodyTopicConfigInfoList) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *GetPolicyInfoResponseBodyTopicConfigInfoList) GetTopicId() *int64 {
	return s.TopicId
}

func (s *GetPolicyInfoResponseBodyTopicConfigInfoList) GetTopicName() *string {
	return s.TopicName
}

func (s *GetPolicyInfoResponseBodyTopicConfigInfoList) SetCategoryType(v int32) *GetPolicyInfoResponseBodyTopicConfigInfoList {
	s.CategoryType = &v
	return s
}

func (s *GetPolicyInfoResponseBodyTopicConfigInfoList) SetInputOutputType(v int32) *GetPolicyInfoResponseBodyTopicConfigInfoList {
	s.InputOutputType = &v
	return s
}

func (s *GetPolicyInfoResponseBodyTopicConfigInfoList) SetSecurityLevel(v int32) *GetPolicyInfoResponseBodyTopicConfigInfoList {
	s.SecurityLevel = &v
	return s
}

func (s *GetPolicyInfoResponseBodyTopicConfigInfoList) SetTopicId(v int64) *GetPolicyInfoResponseBodyTopicConfigInfoList {
	s.TopicId = &v
	return s
}

func (s *GetPolicyInfoResponseBodyTopicConfigInfoList) SetTopicName(v string) *GetPolicyInfoResponseBodyTopicConfigInfoList {
	s.TopicName = &v
	return s
}

func (s *GetPolicyInfoResponseBodyTopicConfigInfoList) Validate() error {
	return dara.Validate(s)
}

type GetPolicyInfoResponseBodyWordGroupInfoList struct {
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

func (s GetPolicyInfoResponseBodyWordGroupInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyInfoResponseBodyWordGroupInfoList) GoString() string {
	return s.String()
}

func (s *GetPolicyInfoResponseBodyWordGroupInfoList) GetGroupId() *int64 {
	return s.GroupId
}

func (s *GetPolicyInfoResponseBodyWordGroupInfoList) GetGroupName() *string {
	return s.GroupName
}

func (s *GetPolicyInfoResponseBodyWordGroupInfoList) GetInputOutputType() *int32 {
	return s.InputOutputType
}

func (s *GetPolicyInfoResponseBodyWordGroupInfoList) SetGroupId(v int64) *GetPolicyInfoResponseBodyWordGroupInfoList {
	s.GroupId = &v
	return s
}

func (s *GetPolicyInfoResponseBodyWordGroupInfoList) SetGroupName(v string) *GetPolicyInfoResponseBodyWordGroupInfoList {
	s.GroupName = &v
	return s
}

func (s *GetPolicyInfoResponseBodyWordGroupInfoList) SetInputOutputType(v int32) *GetPolicyInfoResponseBodyWordGroupInfoList {
	s.InputOutputType = &v
	return s
}

func (s *GetPolicyInfoResponseBodyWordGroupInfoList) Validate() error {
	return dara.Validate(s)
}
