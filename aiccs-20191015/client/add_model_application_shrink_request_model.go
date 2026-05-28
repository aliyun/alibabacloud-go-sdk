// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddModelApplicationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationCps(v int64) *AddModelApplicationShrinkRequest
	GetApplicationCps() *int64
	SetApplicationName(v string) *AddModelApplicationShrinkRequest
	GetApplicationName() *string
	SetCallConnectedTriggerModel(v bool) *AddModelApplicationShrinkRequest
	GetCallConnectedTriggerModel() *bool
	SetDyvmsSceneName(v string) *AddModelApplicationShrinkRequest
	GetDyvmsSceneName() *string
	SetModelCode(v string) *AddModelApplicationShrinkRequest
	GetModelCode() *string
	SetModelVersion(v string) *AddModelApplicationShrinkRequest
	GetModelVersion() *string
	SetMuteActive(v bool) *AddModelApplicationShrinkRequest
	GetMuteActive() *bool
	SetMuteDuration(v int64) *AddModelApplicationShrinkRequest
	GetMuteDuration() *int64
	SetMuteHangupNum(v int64) *AddModelApplicationShrinkRequest
	GetMuteHangupNum() *int64
	SetOwnerId(v int64) *AddModelApplicationShrinkRequest
	GetOwnerId() *int64
	SetPrompt(v string) *AddModelApplicationShrinkRequest
	GetPrompt() *string
	SetQualificationId(v int64) *AddModelApplicationShrinkRequest
	GetQualificationId() *int64
	SetQualificationName(v string) *AddModelApplicationShrinkRequest
	GetQualificationName() *string
	SetRecordingFile(v string) *AddModelApplicationShrinkRequest
	GetRecordingFile() *string
	SetResourceOwnerAccount(v string) *AddModelApplicationShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddModelApplicationShrinkRequest
	GetResourceOwnerId() *int64
	SetSource(v string) *AddModelApplicationShrinkRequest
	GetSource() *string
	SetSpeechContent(v string) *AddModelApplicationShrinkRequest
	GetSpeechContent() *string
	SetSpeechId(v int64) *AddModelApplicationShrinkRequest
	GetSpeechId() *int64
	SetStartWord(v string) *AddModelApplicationShrinkRequest
	GetStartWord() *string
	SetStartWordType(v int64) *AddModelApplicationShrinkRequest
	GetStartWordType() *int64
	SetTtsConfigShrink(v string) *AddModelApplicationShrinkRequest
	GetTtsConfigShrink() *string
	SetUsageDesc(v string) *AddModelApplicationShrinkRequest
	GetUsageDesc() *string
}

type AddModelApplicationShrinkRequest struct {
	// 应用并发请求数
	//
	// This parameter is required.
	//
	// example:
	//
	// 25
	ApplicationCps *int64 `json:"ApplicationCps,omitempty" xml:"ApplicationCps,omitempty"`
	// 模型应用名称
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// example:
	//
	// false
	CallConnectedTriggerModel *bool `json:"CallConnectedTriggerModel,omitempty" xml:"CallConnectedTriggerModel,omitempty"`
	// 场景名称
	//
	// example:
	//
	// 示例值示例值
	DyvmsSceneName *string `json:"DyvmsSceneName,omitempty" xml:"DyvmsSceneName,omitempty"`
	// 模型编码
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	ModelCode *string `json:"ModelCode,omitempty" xml:"ModelCode,omitempty"`
	// 模型版本
	//
	// example:
	//
	// 示例值示例值
	ModelVersion *string `json:"ModelVersion,omitempty" xml:"ModelVersion,omitempty"`
	// 第一个静音是否唤起模型
	//
	// example:
	//
	// false
	MuteActive *bool `json:"MuteActive,omitempty" xml:"MuteActive,omitempty"`
	// 静音时长
	//
	// example:
	//
	// 70
	MuteDuration *int64 `json:"MuteDuration,omitempty" xml:"MuteDuration,omitempty"`
	// 连续多少个静音事件主动挂机
	//
	// example:
	//
	// 5
	MuteHangupNum *int64 `json:"MuteHangupNum,omitempty" xml:"MuteHangupNum,omitempty"`
	OwnerId       *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 提示词
	//
	// example:
	//
	// 示例值示例值示例值
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// 资质ID
	//
	// example:
	//
	// 85
	QualificationId *int64 `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	// 资质名称
	//
	// example:
	//
	// 示例值
	QualificationName *string `json:"QualificationName,omitempty" xml:"QualificationName,omitempty"`
	// example:
	//
	// 示例值示例值
	RecordingFile        *string `json:"RecordingFile,omitempty" xml:"RecordingFile,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 来源
	//
	// example:
	//
	// USER
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// 话术内容
	//
	// example:
	//
	// 示例值示例值示例值
	SpeechContent *string `json:"SpeechContent,omitempty" xml:"SpeechContent,omitempty"`
	// 话束id
	//
	// example:
	//
	// 88
	SpeechId *int64 `json:"SpeechId,omitempty" xml:"SpeechId,omitempty"`
	// 开场白
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	StartWord *string `json:"StartWord,omitempty" xml:"StartWord,omitempty"`
	// example:
	//
	// 1
	StartWordType *int64 `json:"StartWordType,omitempty" xml:"StartWordType,omitempty"`
	// tts配置，包括音色、音量、音速等。
	//
	// This parameter is required.
	TtsConfigShrink *string `json:"TtsConfig,omitempty" xml:"TtsConfig,omitempty"`
	// 用途
	//
	// example:
	//
	// 示例值示例值示例值
	UsageDesc *string `json:"UsageDesc,omitempty" xml:"UsageDesc,omitempty"`
}

func (s AddModelApplicationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddModelApplicationShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddModelApplicationShrinkRequest) GetApplicationCps() *int64 {
	return s.ApplicationCps
}

func (s *AddModelApplicationShrinkRequest) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *AddModelApplicationShrinkRequest) GetCallConnectedTriggerModel() *bool {
	return s.CallConnectedTriggerModel
}

func (s *AddModelApplicationShrinkRequest) GetDyvmsSceneName() *string {
	return s.DyvmsSceneName
}

func (s *AddModelApplicationShrinkRequest) GetModelCode() *string {
	return s.ModelCode
}

func (s *AddModelApplicationShrinkRequest) GetModelVersion() *string {
	return s.ModelVersion
}

func (s *AddModelApplicationShrinkRequest) GetMuteActive() *bool {
	return s.MuteActive
}

func (s *AddModelApplicationShrinkRequest) GetMuteDuration() *int64 {
	return s.MuteDuration
}

func (s *AddModelApplicationShrinkRequest) GetMuteHangupNum() *int64 {
	return s.MuteHangupNum
}

func (s *AddModelApplicationShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddModelApplicationShrinkRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *AddModelApplicationShrinkRequest) GetQualificationId() *int64 {
	return s.QualificationId
}

func (s *AddModelApplicationShrinkRequest) GetQualificationName() *string {
	return s.QualificationName
}

func (s *AddModelApplicationShrinkRequest) GetRecordingFile() *string {
	return s.RecordingFile
}

func (s *AddModelApplicationShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddModelApplicationShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddModelApplicationShrinkRequest) GetSource() *string {
	return s.Source
}

func (s *AddModelApplicationShrinkRequest) GetSpeechContent() *string {
	return s.SpeechContent
}

func (s *AddModelApplicationShrinkRequest) GetSpeechId() *int64 {
	return s.SpeechId
}

func (s *AddModelApplicationShrinkRequest) GetStartWord() *string {
	return s.StartWord
}

func (s *AddModelApplicationShrinkRequest) GetStartWordType() *int64 {
	return s.StartWordType
}

func (s *AddModelApplicationShrinkRequest) GetTtsConfigShrink() *string {
	return s.TtsConfigShrink
}

func (s *AddModelApplicationShrinkRequest) GetUsageDesc() *string {
	return s.UsageDesc
}

func (s *AddModelApplicationShrinkRequest) SetApplicationCps(v int64) *AddModelApplicationShrinkRequest {
	s.ApplicationCps = &v
	return s
}

func (s *AddModelApplicationShrinkRequest) SetApplicationName(v string) *AddModelApplicationShrinkRequest {
	s.ApplicationName = &v
	return s
}

func (s *AddModelApplicationShrinkRequest) SetCallConnectedTriggerModel(v bool) *AddModelApplicationShrinkRequest {
	s.CallConnectedTriggerModel = &v
	return s
}

func (s *AddModelApplicationShrinkRequest) SetDyvmsSceneName(v string) *AddModelApplicationShrinkRequest {
	s.DyvmsSceneName = &v
	return s
}

func (s *AddModelApplicationShrinkRequest) SetModelCode(v string) *AddModelApplicationShrinkRequest {
	s.ModelCode = &v
	return s
}

func (s *AddModelApplicationShrinkRequest) SetModelVersion(v string) *AddModelApplicationShrinkRequest {
	s.ModelVersion = &v
	return s
}

func (s *AddModelApplicationShrinkRequest) SetMuteActive(v bool) *AddModelApplicationShrinkRequest {
	s.MuteActive = &v
	return s
}

func (s *AddModelApplicationShrinkRequest) SetMuteDuration(v int64) *AddModelApplicationShrinkRequest {
	s.MuteDuration = &v
	return s
}

func (s *AddModelApplicationShrinkRequest) SetMuteHangupNum(v int64) *AddModelApplicationShrinkRequest {
	s.MuteHangupNum = &v
	return s
}

func (s *AddModelApplicationShrinkRequest) SetOwnerId(v int64) *AddModelApplicationShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *AddModelApplicationShrinkRequest) SetPrompt(v string) *AddModelApplicationShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *AddModelApplicationShrinkRequest) SetQualificationId(v int64) *AddModelApplicationShrinkRequest {
	s.QualificationId = &v
	return s
}

func (s *AddModelApplicationShrinkRequest) SetQualificationName(v string) *AddModelApplicationShrinkRequest {
	s.QualificationName = &v
	return s
}

func (s *AddModelApplicationShrinkRequest) SetRecordingFile(v string) *AddModelApplicationShrinkRequest {
	s.RecordingFile = &v
	return s
}

func (s *AddModelApplicationShrinkRequest) SetResourceOwnerAccount(v string) *AddModelApplicationShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddModelApplicationShrinkRequest) SetResourceOwnerId(v int64) *AddModelApplicationShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddModelApplicationShrinkRequest) SetSource(v string) *AddModelApplicationShrinkRequest {
	s.Source = &v
	return s
}

func (s *AddModelApplicationShrinkRequest) SetSpeechContent(v string) *AddModelApplicationShrinkRequest {
	s.SpeechContent = &v
	return s
}

func (s *AddModelApplicationShrinkRequest) SetSpeechId(v int64) *AddModelApplicationShrinkRequest {
	s.SpeechId = &v
	return s
}

func (s *AddModelApplicationShrinkRequest) SetStartWord(v string) *AddModelApplicationShrinkRequest {
	s.StartWord = &v
	return s
}

func (s *AddModelApplicationShrinkRequest) SetStartWordType(v int64) *AddModelApplicationShrinkRequest {
	s.StartWordType = &v
	return s
}

func (s *AddModelApplicationShrinkRequest) SetTtsConfigShrink(v string) *AddModelApplicationShrinkRequest {
	s.TtsConfigShrink = &v
	return s
}

func (s *AddModelApplicationShrinkRequest) SetUsageDesc(v string) *AddModelApplicationShrinkRequest {
	s.UsageDesc = &v
	return s
}

func (s *AddModelApplicationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
