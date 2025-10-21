// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentSafeModelInstanceId(v int64) *UpdatePolicyShrinkRequest
	GetContentSafeModelInstanceId() *int64
	SetEnableSensitiveInputCheck(v int32) *UpdatePolicyShrinkRequest
	GetEnableSensitiveInputCheck() *int32
	SetEnableSensitiveOutputCheck(v int32) *UpdatePolicyShrinkRequest
	GetEnableSensitiveOutputCheck() *int32
	SetHarmfulCategoryConfigInfoListShrink(v string) *UpdatePolicyShrinkRequest
	GetHarmfulCategoryConfigInfoListShrink() *string
	SetInputSafeAnswer(v string) *UpdatePolicyShrinkRequest
	GetInputSafeAnswer() *string
	SetInputSafeAnswerSwitch(v int32) *UpdatePolicyShrinkRequest
	GetInputSafeAnswerSwitch() *int32
	SetIsSidecarPolicy(v int32) *UpdatePolicyShrinkRequest
	GetIsSidecarPolicy() *int32
	SetOutputSafeAnswer(v string) *UpdatePolicyShrinkRequest
	GetOutputSafeAnswer() *string
	SetOutputSafeAnswerSwitch(v int32) *UpdatePolicyShrinkRequest
	GetOutputSafeAnswerSwitch() *int32
	SetPolicyId(v int64) *UpdatePolicyShrinkRequest
	GetPolicyId() *int64
	SetPolicyName(v string) *UpdatePolicyShrinkRequest
	GetPolicyName() *string
	SetPromptAttackInfoShrink(v string) *UpdatePolicyShrinkRequest
	GetPromptAttackInfoShrink() *string
	SetPromptAttackInfoListShrink(v string) *UpdatePolicyShrinkRequest
	GetPromptAttackInfoListShrink() *string
	SetPromptAttackModelInstanceId(v int64) *UpdatePolicyShrinkRequest
	GetPromptAttackModelInstanceId() *int64
	SetRegionId(v string) *UpdatePolicyShrinkRequest
	GetRegionId() *string
	SetRegularExpressListShrink(v string) *UpdatePolicyShrinkRequest
	GetRegularExpressListShrink() *string
	SetSceneType(v int32) *UpdatePolicyShrinkRequest
	GetSceneType() *int32
	SetSensitiveConfigListShrink(v string) *UpdatePolicyShrinkRequest
	GetSensitiveConfigListShrink() *string
	SetSensitiveTopicListShrink(v string) *UpdatePolicyShrinkRequest
	GetSensitiveTopicListShrink() *string
	SetSensitiveTopicModelInstanceId(v int64) *UpdatePolicyShrinkRequest
	GetSensitiveTopicModelInstanceId() *int64
	SetSensitiveWordListShrink(v string) *UpdatePolicyShrinkRequest
	GetSensitiveWordListShrink() *string
	SetTopicConfigInfoListShrink(v string) *UpdatePolicyShrinkRequest
	GetTopicConfigInfoListShrink() *string
	SetWordGroupInfoListShrink(v string) *UpdatePolicyShrinkRequest
	GetWordGroupInfoListShrink() *string
	SetWorkspaceId(v int64) *UpdatePolicyShrinkRequest
	GetWorkspaceId() *int64
}

type UpdatePolicyShrinkRequest struct {
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
	WorkspaceId             *int64  `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdatePolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdatePolicyShrinkRequest) GetContentSafeModelInstanceId() *int64 {
	return s.ContentSafeModelInstanceId
}

func (s *UpdatePolicyShrinkRequest) GetEnableSensitiveInputCheck() *int32 {
	return s.EnableSensitiveInputCheck
}

func (s *UpdatePolicyShrinkRequest) GetEnableSensitiveOutputCheck() *int32 {
	return s.EnableSensitiveOutputCheck
}

func (s *UpdatePolicyShrinkRequest) GetHarmfulCategoryConfigInfoListShrink() *string {
	return s.HarmfulCategoryConfigInfoListShrink
}

func (s *UpdatePolicyShrinkRequest) GetInputSafeAnswer() *string {
	return s.InputSafeAnswer
}

func (s *UpdatePolicyShrinkRequest) GetInputSafeAnswerSwitch() *int32 {
	return s.InputSafeAnswerSwitch
}

func (s *UpdatePolicyShrinkRequest) GetIsSidecarPolicy() *int32 {
	return s.IsSidecarPolicy
}

func (s *UpdatePolicyShrinkRequest) GetOutputSafeAnswer() *string {
	return s.OutputSafeAnswer
}

func (s *UpdatePolicyShrinkRequest) GetOutputSafeAnswerSwitch() *int32 {
	return s.OutputSafeAnswerSwitch
}

func (s *UpdatePolicyShrinkRequest) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *UpdatePolicyShrinkRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *UpdatePolicyShrinkRequest) GetPromptAttackInfoShrink() *string {
	return s.PromptAttackInfoShrink
}

func (s *UpdatePolicyShrinkRequest) GetPromptAttackInfoListShrink() *string {
	return s.PromptAttackInfoListShrink
}

func (s *UpdatePolicyShrinkRequest) GetPromptAttackModelInstanceId() *int64 {
	return s.PromptAttackModelInstanceId
}

func (s *UpdatePolicyShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdatePolicyShrinkRequest) GetRegularExpressListShrink() *string {
	return s.RegularExpressListShrink
}

func (s *UpdatePolicyShrinkRequest) GetSceneType() *int32 {
	return s.SceneType
}

func (s *UpdatePolicyShrinkRequest) GetSensitiveConfigListShrink() *string {
	return s.SensitiveConfigListShrink
}

func (s *UpdatePolicyShrinkRequest) GetSensitiveTopicListShrink() *string {
	return s.SensitiveTopicListShrink
}

func (s *UpdatePolicyShrinkRequest) GetSensitiveTopicModelInstanceId() *int64 {
	return s.SensitiveTopicModelInstanceId
}

func (s *UpdatePolicyShrinkRequest) GetSensitiveWordListShrink() *string {
	return s.SensitiveWordListShrink
}

func (s *UpdatePolicyShrinkRequest) GetTopicConfigInfoListShrink() *string {
	return s.TopicConfigInfoListShrink
}

func (s *UpdatePolicyShrinkRequest) GetWordGroupInfoListShrink() *string {
	return s.WordGroupInfoListShrink
}

func (s *UpdatePolicyShrinkRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *UpdatePolicyShrinkRequest) SetContentSafeModelInstanceId(v int64) *UpdatePolicyShrinkRequest {
	s.ContentSafeModelInstanceId = &v
	return s
}

func (s *UpdatePolicyShrinkRequest) SetEnableSensitiveInputCheck(v int32) *UpdatePolicyShrinkRequest {
	s.EnableSensitiveInputCheck = &v
	return s
}

func (s *UpdatePolicyShrinkRequest) SetEnableSensitiveOutputCheck(v int32) *UpdatePolicyShrinkRequest {
	s.EnableSensitiveOutputCheck = &v
	return s
}

func (s *UpdatePolicyShrinkRequest) SetHarmfulCategoryConfigInfoListShrink(v string) *UpdatePolicyShrinkRequest {
	s.HarmfulCategoryConfigInfoListShrink = &v
	return s
}

func (s *UpdatePolicyShrinkRequest) SetInputSafeAnswer(v string) *UpdatePolicyShrinkRequest {
	s.InputSafeAnswer = &v
	return s
}

func (s *UpdatePolicyShrinkRequest) SetInputSafeAnswerSwitch(v int32) *UpdatePolicyShrinkRequest {
	s.InputSafeAnswerSwitch = &v
	return s
}

func (s *UpdatePolicyShrinkRequest) SetIsSidecarPolicy(v int32) *UpdatePolicyShrinkRequest {
	s.IsSidecarPolicy = &v
	return s
}

func (s *UpdatePolicyShrinkRequest) SetOutputSafeAnswer(v string) *UpdatePolicyShrinkRequest {
	s.OutputSafeAnswer = &v
	return s
}

func (s *UpdatePolicyShrinkRequest) SetOutputSafeAnswerSwitch(v int32) *UpdatePolicyShrinkRequest {
	s.OutputSafeAnswerSwitch = &v
	return s
}

func (s *UpdatePolicyShrinkRequest) SetPolicyId(v int64) *UpdatePolicyShrinkRequest {
	s.PolicyId = &v
	return s
}

func (s *UpdatePolicyShrinkRequest) SetPolicyName(v string) *UpdatePolicyShrinkRequest {
	s.PolicyName = &v
	return s
}

func (s *UpdatePolicyShrinkRequest) SetPromptAttackInfoShrink(v string) *UpdatePolicyShrinkRequest {
	s.PromptAttackInfoShrink = &v
	return s
}

func (s *UpdatePolicyShrinkRequest) SetPromptAttackInfoListShrink(v string) *UpdatePolicyShrinkRequest {
	s.PromptAttackInfoListShrink = &v
	return s
}

func (s *UpdatePolicyShrinkRequest) SetPromptAttackModelInstanceId(v int64) *UpdatePolicyShrinkRequest {
	s.PromptAttackModelInstanceId = &v
	return s
}

func (s *UpdatePolicyShrinkRequest) SetRegionId(v string) *UpdatePolicyShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdatePolicyShrinkRequest) SetRegularExpressListShrink(v string) *UpdatePolicyShrinkRequest {
	s.RegularExpressListShrink = &v
	return s
}

func (s *UpdatePolicyShrinkRequest) SetSceneType(v int32) *UpdatePolicyShrinkRequest {
	s.SceneType = &v
	return s
}

func (s *UpdatePolicyShrinkRequest) SetSensitiveConfigListShrink(v string) *UpdatePolicyShrinkRequest {
	s.SensitiveConfigListShrink = &v
	return s
}

func (s *UpdatePolicyShrinkRequest) SetSensitiveTopicListShrink(v string) *UpdatePolicyShrinkRequest {
	s.SensitiveTopicListShrink = &v
	return s
}

func (s *UpdatePolicyShrinkRequest) SetSensitiveTopicModelInstanceId(v int64) *UpdatePolicyShrinkRequest {
	s.SensitiveTopicModelInstanceId = &v
	return s
}

func (s *UpdatePolicyShrinkRequest) SetSensitiveWordListShrink(v string) *UpdatePolicyShrinkRequest {
	s.SensitiveWordListShrink = &v
	return s
}

func (s *UpdatePolicyShrinkRequest) SetTopicConfigInfoListShrink(v string) *UpdatePolicyShrinkRequest {
	s.TopicConfigInfoListShrink = &v
	return s
}

func (s *UpdatePolicyShrinkRequest) SetWordGroupInfoListShrink(v string) *UpdatePolicyShrinkRequest {
	s.WordGroupInfoListShrink = &v
	return s
}

func (s *UpdatePolicyShrinkRequest) SetWorkspaceId(v int64) *UpdatePolicyShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdatePolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
