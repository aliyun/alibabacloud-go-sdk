// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentSafeModelInstanceId(v int64) *CreatePolicyShrinkRequest
	GetContentSafeModelInstanceId() *int64
	SetEnableSensitiveInputCheck(v int32) *CreatePolicyShrinkRequest
	GetEnableSensitiveInputCheck() *int32
	SetEnableSensitiveOutputCheck(v int32) *CreatePolicyShrinkRequest
	GetEnableSensitiveOutputCheck() *int32
	SetHarmfulCategoryConfigInfoListShrink(v string) *CreatePolicyShrinkRequest
	GetHarmfulCategoryConfigInfoListShrink() *string
	SetInputSafeAnswer(v string) *CreatePolicyShrinkRequest
	GetInputSafeAnswer() *string
	SetInputSafeAnswerSwitch(v int32) *CreatePolicyShrinkRequest
	GetInputSafeAnswerSwitch() *int32
	SetIsSidecarPolicy(v int32) *CreatePolicyShrinkRequest
	GetIsSidecarPolicy() *int32
	SetOutputSafeAnswer(v string) *CreatePolicyShrinkRequest
	GetOutputSafeAnswer() *string
	SetOutputSafeAnswerSwitch(v int32) *CreatePolicyShrinkRequest
	GetOutputSafeAnswerSwitch() *int32
	SetPolicyName(v string) *CreatePolicyShrinkRequest
	GetPolicyName() *string
	SetPromptAttackInfoShrink(v string) *CreatePolicyShrinkRequest
	GetPromptAttackInfoShrink() *string
	SetPromptAttackInfoListShrink(v string) *CreatePolicyShrinkRequest
	GetPromptAttackInfoListShrink() *string
	SetPromptAttackModelInstanceId(v int64) *CreatePolicyShrinkRequest
	GetPromptAttackModelInstanceId() *int64
	SetRegionId(v string) *CreatePolicyShrinkRequest
	GetRegionId() *string
	SetRegularExpressListShrink(v string) *CreatePolicyShrinkRequest
	GetRegularExpressListShrink() *string
	SetSceneType(v int32) *CreatePolicyShrinkRequest
	GetSceneType() *int32
	SetSensitiveConfigListShrink(v string) *CreatePolicyShrinkRequest
	GetSensitiveConfigListShrink() *string
	SetSensitiveTopicListShrink(v string) *CreatePolicyShrinkRequest
	GetSensitiveTopicListShrink() *string
	SetSensitiveTopicModelInstanceId(v int64) *CreatePolicyShrinkRequest
	GetSensitiveTopicModelInstanceId() *int64
	SetSensitiveWordListShrink(v string) *CreatePolicyShrinkRequest
	GetSensitiveWordListShrink() *string
	SetTopicConfigInfoListShrink(v string) *CreatePolicyShrinkRequest
	GetTopicConfigInfoListShrink() *string
	SetWordGroupInfoListShrink(v string) *CreatePolicyShrinkRequest
	GetWordGroupInfoListShrink() *string
	SetWorkspaceId(v int64) *CreatePolicyShrinkRequest
	GetWorkspaceId() *int64
}

type CreatePolicyShrinkRequest struct {
	ContentSafeModelInstanceId *int64 `json:"ContentSafeModelInstanceId,omitempty" xml:"ContentSafeModelInstanceId,omitempty"`
	EnableSensitiveInputCheck  *int32 `json:"EnableSensitiveInputCheck,omitempty" xml:"EnableSensitiveInputCheck,omitempty"`
	EnableSensitiveOutputCheck *int32 `json:"EnableSensitiveOutputCheck,omitempty" xml:"EnableSensitiveOutputCheck,omitempty"`
	// List of harmful category configurations
	HarmfulCategoryConfigInfoListShrink *string `json:"HarmfulCategoryConfigInfoList,omitempty" xml:"HarmfulCategoryConfigInfoList,omitempty"`
	InputSafeAnswer                     *string `json:"InputSafeAnswer,omitempty" xml:"InputSafeAnswer,omitempty"`
	InputSafeAnswerSwitch               *int32  `json:"InputSafeAnswerSwitch,omitempty" xml:"InputSafeAnswerSwitch,omitempty"`
	IsSidecarPolicy                     *int32  `json:"IsSidecarPolicy,omitempty" xml:"IsSidecarPolicy,omitempty"`
	OutputSafeAnswer                    *string `json:"OutputSafeAnswer,omitempty" xml:"OutputSafeAnswer,omitempty"`
	OutputSafeAnswerSwitch              *int32  `json:"OutputSafeAnswerSwitch,omitempty" xml:"OutputSafeAnswerSwitch,omitempty"`
	// Detection policy name.
	//
	// example:
	//
	// testPolicy
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// Prompt attack detection result object
	PromptAttackInfoShrink *string `json:"PromptAttackInfo,omitempty" xml:"PromptAttackInfo,omitempty"`
	// List of prompt attacks
	PromptAttackInfoListShrink  *string `json:"PromptAttackInfoList,omitempty" xml:"PromptAttackInfoList,omitempty"`
	PromptAttackModelInstanceId *int64  `json:"PromptAttackModelInstanceId,omitempty" xml:"PromptAttackModelInstanceId,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId                      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegularExpressListShrink      *string `json:"RegularExpressList,omitempty" xml:"RegularExpressList,omitempty"`
	SceneType                     *int32  `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
	SensitiveConfigListShrink     *string `json:"SensitiveConfigList,omitempty" xml:"SensitiveConfigList,omitempty"`
	SensitiveTopicListShrink      *string `json:"SensitiveTopicList,omitempty" xml:"SensitiveTopicList,omitempty"`
	SensitiveTopicModelInstanceId *int64  `json:"SensitiveTopicModelInstanceId,omitempty" xml:"SensitiveTopicModelInstanceId,omitempty"`
	SensitiveWordListShrink       *string `json:"SensitiveWordList,omitempty" xml:"SensitiveWordList,omitempty"`
	// List of sensitive topics
	TopicConfigInfoListShrink *string `json:"TopicConfigInfoList,omitempty" xml:"TopicConfigInfoList,omitempty"`
	// List of keyword group objects
	WordGroupInfoListShrink *string `json:"WordGroupInfoList,omitempty" xml:"WordGroupInfoList,omitempty"`
	// Workspace ID
	//
	// example:
	//
	// 608226
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreatePolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyShrinkRequest) GetContentSafeModelInstanceId() *int64 {
	return s.ContentSafeModelInstanceId
}

func (s *CreatePolicyShrinkRequest) GetEnableSensitiveInputCheck() *int32 {
	return s.EnableSensitiveInputCheck
}

func (s *CreatePolicyShrinkRequest) GetEnableSensitiveOutputCheck() *int32 {
	return s.EnableSensitiveOutputCheck
}

func (s *CreatePolicyShrinkRequest) GetHarmfulCategoryConfigInfoListShrink() *string {
	return s.HarmfulCategoryConfigInfoListShrink
}

func (s *CreatePolicyShrinkRequest) GetInputSafeAnswer() *string {
	return s.InputSafeAnswer
}

func (s *CreatePolicyShrinkRequest) GetInputSafeAnswerSwitch() *int32 {
	return s.InputSafeAnswerSwitch
}

func (s *CreatePolicyShrinkRequest) GetIsSidecarPolicy() *int32 {
	return s.IsSidecarPolicy
}

func (s *CreatePolicyShrinkRequest) GetOutputSafeAnswer() *string {
	return s.OutputSafeAnswer
}

func (s *CreatePolicyShrinkRequest) GetOutputSafeAnswerSwitch() *int32 {
	return s.OutputSafeAnswerSwitch
}

func (s *CreatePolicyShrinkRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *CreatePolicyShrinkRequest) GetPromptAttackInfoShrink() *string {
	return s.PromptAttackInfoShrink
}

func (s *CreatePolicyShrinkRequest) GetPromptAttackInfoListShrink() *string {
	return s.PromptAttackInfoListShrink
}

func (s *CreatePolicyShrinkRequest) GetPromptAttackModelInstanceId() *int64 {
	return s.PromptAttackModelInstanceId
}

func (s *CreatePolicyShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePolicyShrinkRequest) GetRegularExpressListShrink() *string {
	return s.RegularExpressListShrink
}

func (s *CreatePolicyShrinkRequest) GetSceneType() *int32 {
	return s.SceneType
}

func (s *CreatePolicyShrinkRequest) GetSensitiveConfigListShrink() *string {
	return s.SensitiveConfigListShrink
}

func (s *CreatePolicyShrinkRequest) GetSensitiveTopicListShrink() *string {
	return s.SensitiveTopicListShrink
}

func (s *CreatePolicyShrinkRequest) GetSensitiveTopicModelInstanceId() *int64 {
	return s.SensitiveTopicModelInstanceId
}

func (s *CreatePolicyShrinkRequest) GetSensitiveWordListShrink() *string {
	return s.SensitiveWordListShrink
}

func (s *CreatePolicyShrinkRequest) GetTopicConfigInfoListShrink() *string {
	return s.TopicConfigInfoListShrink
}

func (s *CreatePolicyShrinkRequest) GetWordGroupInfoListShrink() *string {
	return s.WordGroupInfoListShrink
}

func (s *CreatePolicyShrinkRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *CreatePolicyShrinkRequest) SetContentSafeModelInstanceId(v int64) *CreatePolicyShrinkRequest {
	s.ContentSafeModelInstanceId = &v
	return s
}

func (s *CreatePolicyShrinkRequest) SetEnableSensitiveInputCheck(v int32) *CreatePolicyShrinkRequest {
	s.EnableSensitiveInputCheck = &v
	return s
}

func (s *CreatePolicyShrinkRequest) SetEnableSensitiveOutputCheck(v int32) *CreatePolicyShrinkRequest {
	s.EnableSensitiveOutputCheck = &v
	return s
}

func (s *CreatePolicyShrinkRequest) SetHarmfulCategoryConfigInfoListShrink(v string) *CreatePolicyShrinkRequest {
	s.HarmfulCategoryConfigInfoListShrink = &v
	return s
}

func (s *CreatePolicyShrinkRequest) SetInputSafeAnswer(v string) *CreatePolicyShrinkRequest {
	s.InputSafeAnswer = &v
	return s
}

func (s *CreatePolicyShrinkRequest) SetInputSafeAnswerSwitch(v int32) *CreatePolicyShrinkRequest {
	s.InputSafeAnswerSwitch = &v
	return s
}

func (s *CreatePolicyShrinkRequest) SetIsSidecarPolicy(v int32) *CreatePolicyShrinkRequest {
	s.IsSidecarPolicy = &v
	return s
}

func (s *CreatePolicyShrinkRequest) SetOutputSafeAnswer(v string) *CreatePolicyShrinkRequest {
	s.OutputSafeAnswer = &v
	return s
}

func (s *CreatePolicyShrinkRequest) SetOutputSafeAnswerSwitch(v int32) *CreatePolicyShrinkRequest {
	s.OutputSafeAnswerSwitch = &v
	return s
}

func (s *CreatePolicyShrinkRequest) SetPolicyName(v string) *CreatePolicyShrinkRequest {
	s.PolicyName = &v
	return s
}

func (s *CreatePolicyShrinkRequest) SetPromptAttackInfoShrink(v string) *CreatePolicyShrinkRequest {
	s.PromptAttackInfoShrink = &v
	return s
}

func (s *CreatePolicyShrinkRequest) SetPromptAttackInfoListShrink(v string) *CreatePolicyShrinkRequest {
	s.PromptAttackInfoListShrink = &v
	return s
}

func (s *CreatePolicyShrinkRequest) SetPromptAttackModelInstanceId(v int64) *CreatePolicyShrinkRequest {
	s.PromptAttackModelInstanceId = &v
	return s
}

func (s *CreatePolicyShrinkRequest) SetRegionId(v string) *CreatePolicyShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePolicyShrinkRequest) SetRegularExpressListShrink(v string) *CreatePolicyShrinkRequest {
	s.RegularExpressListShrink = &v
	return s
}

func (s *CreatePolicyShrinkRequest) SetSceneType(v int32) *CreatePolicyShrinkRequest {
	s.SceneType = &v
	return s
}

func (s *CreatePolicyShrinkRequest) SetSensitiveConfigListShrink(v string) *CreatePolicyShrinkRequest {
	s.SensitiveConfigListShrink = &v
	return s
}

func (s *CreatePolicyShrinkRequest) SetSensitiveTopicListShrink(v string) *CreatePolicyShrinkRequest {
	s.SensitiveTopicListShrink = &v
	return s
}

func (s *CreatePolicyShrinkRequest) SetSensitiveTopicModelInstanceId(v int64) *CreatePolicyShrinkRequest {
	s.SensitiveTopicModelInstanceId = &v
	return s
}

func (s *CreatePolicyShrinkRequest) SetSensitiveWordListShrink(v string) *CreatePolicyShrinkRequest {
	s.SensitiveWordListShrink = &v
	return s
}

func (s *CreatePolicyShrinkRequest) SetTopicConfigInfoListShrink(v string) *CreatePolicyShrinkRequest {
	s.TopicConfigInfoListShrink = &v
	return s
}

func (s *CreatePolicyShrinkRequest) SetWordGroupInfoListShrink(v string) *CreatePolicyShrinkRequest {
	s.WordGroupInfoListShrink = &v
	return s
}

func (s *CreatePolicyShrinkRequest) SetWorkspaceId(v int64) *CreatePolicyShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreatePolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
