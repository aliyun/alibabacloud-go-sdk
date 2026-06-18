// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAiVoiceAgentDetailNewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryAiVoiceAgentDetailNewResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QueryAiVoiceAgentDetailNewResponseBody
	GetCode() *string
	SetData(v *QueryAiVoiceAgentDetailNewResponseBodyData) *QueryAiVoiceAgentDetailNewResponseBody
	GetData() *QueryAiVoiceAgentDetailNewResponseBodyData
	SetMessage(v string) *QueryAiVoiceAgentDetailNewResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryAiVoiceAgentDetailNewResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAiVoiceAgentDetailNewResponseBody
	GetSuccess() *bool
}

type QueryAiVoiceAgentDetailNewResponseBody struct {
	// The reason for the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detailed data of the agent.
	Data *QueryAiVoiceAgentDetailNewResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is associated with the status code.
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 742C9243-2870-B8D6-0C68-C12345AABBCC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call was successful. Possible values:
	//
	// - **true**: The operation was successful.
	//
	// - **false**: Failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAiVoiceAgentDetailNewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailNewResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailNewResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryAiVoiceAgentDetailNewResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryAiVoiceAgentDetailNewResponseBody) GetData() *QueryAiVoiceAgentDetailNewResponseBodyData {
	return s.Data
}

func (s *QueryAiVoiceAgentDetailNewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAiVoiceAgentDetailNewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAiVoiceAgentDetailNewResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAiVoiceAgentDetailNewResponseBody) SetAccessDeniedDetail(v string) *QueryAiVoiceAgentDetailNewResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBody) SetCode(v string) *QueryAiVoiceAgentDetailNewResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBody) SetData(v *QueryAiVoiceAgentDetailNewResponseBodyData) *QueryAiVoiceAgentDetailNewResponseBody {
	s.Data = v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBody) SetMessage(v string) *QueryAiVoiceAgentDetailNewResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBody) SetRequestId(v string) *QueryAiVoiceAgentDetailNewResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBody) SetSuccess(v bool) *QueryAiVoiceAgentDetailNewResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAiVoiceAgentDetailNewResponseBodyData struct {
	// The voice configuration for intelligent outbound calls.
	AgentCallConfig *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfig `json:"AgentCallConfig,omitempty" xml:"AgentCallConfig,omitempty" type:"Struct"`
	// The business requirement configuration for the agent.
	AgentDemandConfig *QueryAiVoiceAgentDetailNewResponseBodyDataAgentDemandConfig `json:"AgentDemandConfig,omitempty" xml:"AgentDemandConfig,omitempty" type:"Struct"`
	// The agent description.
	//
	// example:
	//
	// 日常测试使用。
	AgentDesc *string `json:"AgentDesc,omitempty" xml:"AgentDesc,omitempty"`
	// The agent ID.
	//
	// example:
	//
	// 3021893791
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The build mode. Valid values:
	//
	// - `0`: prompt-based mode.
	//
	// - `1`: dialog flow mode.
	//
	// example:
	//
	// 1
	AgentMode *int64 `json:"AgentMode,omitempty" xml:"AgentMode,omitempty"`
	// The agent name.
	//
	// example:
	//
	// 测试智能体
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The deployment status. Valid values:
	//
	// - `0`: Inactive (NOT_EFFECT).
	//
	// - `1`: Active (EFFECT).
	//
	// example:
	//
	// 1
	BranchDeployStatus *int64 `json:"BranchDeployStatus,omitempty" xml:"BranchDeployStatus,omitempty"`
	// The branch description.
	//
	// example:
	//
	// 默认分支
	BranchDesc *string `json:"BranchDesc,omitempty" xml:"BranchDesc,omitempty"`
	// The branch ID.
	//
	// example:
	//
	// 18
	BranchId *int64 `json:"BranchId,omitempty" xml:"BranchId,omitempty"`
	// The branch name.
	//
	// example:
	//
	// 默认版本
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	// The knowledge base configuration.
	KnowledgeConfig *QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfig `json:"KnowledgeConfig,omitempty" xml:"KnowledgeConfig,omitempty" type:"Struct"`
	// The call variable configuration.
	PhoneTagConfig []*QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfig `json:"PhoneTagConfig,omitempty" xml:"PhoneTagConfig,omitempty" type:"Repeated"`
	// The scenario.
	//
	// example:
	//
	// 个人客户线索转化
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The call summary configuration.
	SummaryConfig *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfig `json:"SummaryConfig,omitempty" xml:"SummaryConfig,omitempty" type:"Struct"`
	// Version Description
	//
	// example:
	//
	// 测试第一个版本
	VersionDesc *string `json:"VersionDesc,omitempty" xml:"VersionDesc,omitempty"`
	// Version ID.
	//
	// example:
	//
	// 89
	VersionId *int64 `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// Version name.
	//
	// example:
	//
	// 测试版本1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
	// The release status of the version. 0 indicates Unreleased, 1 indicates Released, and 2 indicates Draft.
	//
	// example:
	//
	// 1
	VersionPublishStatus *int64 `json:"VersionPublishStatus,omitempty" xml:"VersionPublishStatus,omitempty"`
	// The most recent release time of the version.
	//
	// example:
	//
	// 2024-01-15 10:30:00
	VersionPublishTime *string `json:"VersionPublishTime,omitempty" xml:"VersionPublishTime,omitempty"`
}

func (s QueryAiVoiceAgentDetailNewResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailNewResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) GetAgentCallConfig() *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfig {
	return s.AgentCallConfig
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) GetAgentDemandConfig() *QueryAiVoiceAgentDetailNewResponseBodyDataAgentDemandConfig {
	return s.AgentDemandConfig
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) GetAgentDesc() *string {
	return s.AgentDesc
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) GetAgentId() *int64 {
	return s.AgentId
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) GetAgentMode() *int64 {
	return s.AgentMode
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) GetAgentName() *string {
	return s.AgentName
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) GetBranchDeployStatus() *int64 {
	return s.BranchDeployStatus
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) GetBranchDesc() *string {
	return s.BranchDesc
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) GetBranchId() *int64 {
	return s.BranchId
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) GetBranchName() *string {
	return s.BranchName
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) GetKnowledgeConfig() *QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfig {
	return s.KnowledgeConfig
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) GetPhoneTagConfig() []*QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfig {
	return s.PhoneTagConfig
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) GetScene() *string {
	return s.Scene
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) GetSummaryConfig() *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfig {
	return s.SummaryConfig
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) GetVersionDesc() *string {
	return s.VersionDesc
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) GetVersionId() *int64 {
	return s.VersionId
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) GetVersionName() *string {
	return s.VersionName
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) GetVersionPublishStatus() *int64 {
	return s.VersionPublishStatus
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) GetVersionPublishTime() *string {
	return s.VersionPublishTime
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) SetAgentCallConfig(v *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfig) *QueryAiVoiceAgentDetailNewResponseBodyData {
	s.AgentCallConfig = v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) SetAgentDemandConfig(v *QueryAiVoiceAgentDetailNewResponseBodyDataAgentDemandConfig) *QueryAiVoiceAgentDetailNewResponseBodyData {
	s.AgentDemandConfig = v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) SetAgentDesc(v string) *QueryAiVoiceAgentDetailNewResponseBodyData {
	s.AgentDesc = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) SetAgentId(v int64) *QueryAiVoiceAgentDetailNewResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) SetAgentMode(v int64) *QueryAiVoiceAgentDetailNewResponseBodyData {
	s.AgentMode = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) SetAgentName(v string) *QueryAiVoiceAgentDetailNewResponseBodyData {
	s.AgentName = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) SetBranchDeployStatus(v int64) *QueryAiVoiceAgentDetailNewResponseBodyData {
	s.BranchDeployStatus = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) SetBranchDesc(v string) *QueryAiVoiceAgentDetailNewResponseBodyData {
	s.BranchDesc = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) SetBranchId(v int64) *QueryAiVoiceAgentDetailNewResponseBodyData {
	s.BranchId = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) SetBranchName(v string) *QueryAiVoiceAgentDetailNewResponseBodyData {
	s.BranchName = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) SetKnowledgeConfig(v *QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfig) *QueryAiVoiceAgentDetailNewResponseBodyData {
	s.KnowledgeConfig = v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) SetPhoneTagConfig(v []*QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfig) *QueryAiVoiceAgentDetailNewResponseBodyData {
	s.PhoneTagConfig = v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) SetScene(v string) *QueryAiVoiceAgentDetailNewResponseBodyData {
	s.Scene = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) SetSummaryConfig(v *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfig) *QueryAiVoiceAgentDetailNewResponseBodyData {
	s.SummaryConfig = v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) SetVersionDesc(v string) *QueryAiVoiceAgentDetailNewResponseBodyData {
	s.VersionDesc = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) SetVersionId(v int64) *QueryAiVoiceAgentDetailNewResponseBodyData {
	s.VersionId = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) SetVersionName(v string) *QueryAiVoiceAgentDetailNewResponseBodyData {
	s.VersionName = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) SetVersionPublishStatus(v int64) *QueryAiVoiceAgentDetailNewResponseBodyData {
	s.VersionPublishStatus = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) SetVersionPublishTime(v string) *QueryAiVoiceAgentDetailNewResponseBodyData {
	s.VersionPublishTime = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyData) Validate() error {
	if s.AgentCallConfig != nil {
		if err := s.AgentCallConfig.Validate(); err != nil {
			return err
		}
	}
	if s.AgentDemandConfig != nil {
		if err := s.AgentDemandConfig.Validate(); err != nil {
			return err
		}
	}
	if s.KnowledgeConfig != nil {
		if err := s.KnowledgeConfig.Validate(); err != nil {
			return err
		}
	}
	if s.PhoneTagConfig != nil {
		for _, item := range s.PhoneTagConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SummaryConfig != nil {
		if err := s.SummaryConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfig struct {
	// The event configuration.
	EventConfig *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigEventConfig `json:"EventConfig,omitempty" xml:"EventConfig,omitempty" type:"Struct"`
	// The prologue.
	//
	// example:
	//
	// 你好，这里是******。
	Prologue *string `json:"Prologue,omitempty" xml:"Prologue,omitempty"`
	// The URL of the audio file for the prologue. This parameter is returned only when `StartWordType` is set to `1`.
	//
	// example:
	//
	// /oss-file-key
	RecordingFile *string `json:"RecordingFile,omitempty" xml:"RecordingFile,omitempty"`
	// The type of the prologue. Valid values: `0` (text) and `1` (recording).
	//
	// example:
	//
	// 1
	StartWordType *int64 `json:"StartWordType,omitempty" xml:"StartWordType,omitempty"`
	// The configuration for transferring the call to a manual agent.
	TransferConfig *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig `json:"TransferConfig,omitempty" xml:"TransferConfig,omitempty" type:"Struct"`
	// The Text-to-Speech (TTS) configuration.
	TtsConfig *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig `json:"TtsConfig,omitempty" xml:"TtsConfig,omitempty" type:"Struct"`
	// The ID of the hotword vocabulary.
	//
	// example:
	//
	// afb2c43**********83e6df30551c11f7
	VocabId *string `json:"VocabId,omitempty" xml:"VocabId,omitempty"`
}

func (s QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfig) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfig) GetEventConfig() *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigEventConfig {
	return s.EventConfig
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfig) GetPrologue() *string {
	return s.Prologue
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfig) GetRecordingFile() *string {
	return s.RecordingFile
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfig) GetStartWordType() *int64 {
	return s.StartWordType
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfig) GetTransferConfig() *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig {
	return s.TransferConfig
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfig) GetTtsConfig() *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig {
	return s.TtsConfig
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfig) GetVocabId() *string {
	return s.VocabId
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfig) SetEventConfig(v *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigEventConfig) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfig {
	s.EventConfig = v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfig) SetPrologue(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfig {
	s.Prologue = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfig) SetRecordingFile(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfig {
	s.RecordingFile = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfig) SetStartWordType(v int64) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfig {
	s.StartWordType = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfig) SetTransferConfig(v *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfig {
	s.TransferConfig = v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfig) SetTtsConfig(v *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfig {
	s.TtsConfig = v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfig) SetVocabId(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfig {
	s.VocabId = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfig) Validate() error {
	if s.EventConfig != nil {
		if err := s.EventConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TransferConfig != nil {
		if err := s.TransferConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TtsConfig != nil {
		if err := s.TtsConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigEventConfig struct {
	// Specifies whether to disconnect the call when an answering machine is detected.
	//
	// example:
	//
	// false
	CallAssistantHangup *bool `json:"CallAssistantHangup,omitempty" xml:"CallAssistantHangup,omitempty"`
	// Specifies whether to enable answering machine detection.
	//
	// example:
	//
	// true
	CallAssistantRecognize *bool `json:"CallAssistantRecognize,omitempty" xml:"CallAssistantRecognize,omitempty"`
	// Specifies whether to wake up the model upon the first mute event.
	//
	// example:
	//
	// true
	MuteActive *bool `json:"MuteActive,omitempty" xml:"MuteActive,omitempty"`
	// The mute duration. Unit: seconds. Valid values: 3 to 15.
	//
	// example:
	//
	// 5
	MuteDuration *int64 `json:"MuteDuration,omitempty" xml:"MuteDuration,omitempty"`
	// The number of consecutive mute events that trigger an automatic disconnection. Valid values: 1 to 5.
	//
	// example:
	//
	// 1
	MuteHangupNum *int64 `json:"MuteHangupNum,omitempty" xml:"MuteHangupNum,omitempty"`
	// The maximum call duration. Unit: seconds. Valid values: 600 to 3600. The call is automatically disconnected if this duration is exceeded.
	//
	// example:
	//
	// 1000
	SessionTimeout *int64 `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
}

func (s QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigEventConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigEventConfig) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigEventConfig) GetCallAssistantHangup() *bool {
	return s.CallAssistantHangup
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigEventConfig) GetCallAssistantRecognize() *bool {
	return s.CallAssistantRecognize
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigEventConfig) GetMuteActive() *bool {
	return s.MuteActive
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigEventConfig) GetMuteDuration() *int64 {
	return s.MuteDuration
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigEventConfig) GetMuteHangupNum() *int64 {
	return s.MuteHangupNum
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigEventConfig) GetSessionTimeout() *int64 {
	return s.SessionTimeout
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigEventConfig) SetCallAssistantHangup(v bool) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigEventConfig {
	s.CallAssistantHangup = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigEventConfig) SetCallAssistantRecognize(v bool) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigEventConfig {
	s.CallAssistantRecognize = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigEventConfig) SetMuteActive(v bool) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigEventConfig {
	s.MuteActive = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigEventConfig) SetMuteDuration(v int64) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigEventConfig {
	s.MuteDuration = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigEventConfig) SetMuteHangupNum(v int64) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigEventConfig {
	s.MuteHangupNum = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigEventConfig) SetSessionTimeout(v int64) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigEventConfig {
	s.SessionTimeout = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigEventConfig) Validate() error {
	return dara.Validate(s)
}

type QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig struct {
	// The called number.
	//
	// example:
	//
	// 123111122222
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// The caller number.
	//
	// example:
	//
	// 123111122222
	CallerNumber *string `json:"CallerNumber,omitempty" xml:"CallerNumber,omitempty"`
	// The type of the caller number.
	//
	// example:
	//
	// 48
	CallerNumberType *int64 `json:"CallerNumberType,omitempty" xml:"CallerNumberType,omitempty"`
	// The destination number for the transfer, such as the phone number of a customer service agent.
	//
	// example:
	//
	// 123111122222
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// The type of the called number.
	//
	// example:
	//
	// 1
	CallingNumberType *int64 `json:"CallingNumberType,omitempty" xml:"CallingNumberType,omitempty"`
	// The customer route code.
	//
	// example:
	//
	// 123
	CustomerRouteCode *string `json:"CustomerRouteCode,omitempty" xml:"CustomerRouteCode,omitempty"`
	// Specifies whether to enable call transfer.
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The supplementary information.
	//
	// example:
	//
	// Additional information for the agent
	ExtraInfo *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// The prompt that is played when the transfer to a manual agent fails.
	//
	// example:
	//
	// 失败说明
	FailureContent *string `json:"FailureContent,omitempty" xml:"FailureContent,omitempty"`
	// The agent route code.
	//
	// example:
	//
	// SEAT_CTI_A_NET
	SeatRouteCode *string `json:"SeatRouteCode,omitempty" xml:"SeatRouteCode,omitempty"`
	// The name of the agent route.
	//
	// example:
	//
	// 人工座席
	SeatRouteName *string `json:"SeatRouteName,omitempty" xml:"SeatRouteName,omitempty"`
	// The business ID.
	//
	// example:
	//
	// 123
	TransferBizId *string `json:"TransferBizId,omitempty" xml:"TransferBizId,omitempty"`
	// The prompt that is played when the call is transferred to a manual agent.
	//
	// example:
	//
	// 提示语
	TransferContent *string `json:"TransferContent,omitempty" xml:"TransferContent,omitempty"`
	// The agent transfer type.
	//
	// example:
	//
	// 73
	TransferType *int64 `json:"TransferType,omitempty" xml:"TransferType,omitempty"`
}

func (s QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig) GetCallerNumber() *string {
	return s.CallerNumber
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig) GetCallerNumberType() *int64 {
	return s.CallerNumberType
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig) GetCallingNumber() *string {
	return s.CallingNumber
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig) GetCallingNumberType() *int64 {
	return s.CallingNumberType
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig) GetCustomerRouteCode() *string {
	return s.CustomerRouteCode
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig) GetFailureContent() *string {
	return s.FailureContent
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig) GetSeatRouteCode() *string {
	return s.SeatRouteCode
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig) GetSeatRouteName() *string {
	return s.SeatRouteName
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig) GetTransferBizId() *string {
	return s.TransferBizId
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig) GetTransferContent() *string {
	return s.TransferContent
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig) GetTransferType() *int64 {
	return s.TransferType
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig) SetCalledNumber(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig {
	s.CalledNumber = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig) SetCallerNumber(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig {
	s.CallerNumber = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig) SetCallerNumberType(v int64) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig {
	s.CallerNumberType = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig) SetCallingNumber(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig {
	s.CallingNumber = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig) SetCallingNumberType(v int64) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig {
	s.CallingNumberType = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig) SetCustomerRouteCode(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig {
	s.CustomerRouteCode = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig) SetEnabled(v bool) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig {
	s.Enabled = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig) SetExtraInfo(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig {
	s.ExtraInfo = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig) SetFailureContent(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig {
	s.FailureContent = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig) SetSeatRouteCode(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig {
	s.SeatRouteCode = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig) SetSeatRouteName(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig {
	s.SeatRouteName = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig) SetTransferBizId(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig {
	s.TransferBizId = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig) SetTransferContent(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig {
	s.TransferContent = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig) SetTransferType(v int64) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig {
	s.TransferType = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTransferConfig) Validate() error {
	return dara.Validate(s)
}

type QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig struct {
	// Specifies whether to enable background sound.
	//
	// example:
	//
	// true
	BackgroundEnabled *bool `json:"BackgroundEnabled,omitempty" xml:"BackgroundEnabled,omitempty"`
	// The ID of the background sound.
	//
	// example:
	//
	// 17
	BackgroundSound *int64 `json:"BackgroundSound,omitempty" xml:"BackgroundSound,omitempty"`
	// The volume of the background sound. Valid values: `0` (low), `1` (medium), and `2` (high).
	//
	// example:
	//
	// 1
	BackgroundVolume *int64 `json:"BackgroundVolume,omitempty" xml:"BackgroundVolume,omitempty"`
	// The account of the third-party voice platform.
	//
	// example:
	//
	// 75
	CustomerAccountId *int64 `json:"CustomerAccountId,omitempty" xml:"CustomerAccountId,omitempty"`
	// Specifies whether to enable audio mixing.
	//
	// example:
	//
	// true
	MixingEnabled *bool `json:"MixingEnabled,omitempty" xml:"MixingEnabled,omitempty"`
	// The ID of the audio mixing template.
	//
	// example:
	//
	// 1
	MixingTemplate *int64 `json:"MixingTemplate,omitempty" xml:"MixingTemplate,omitempty"`
	// The resource ID. This parameter is required only for third-party voices.
	//
	// example:
	//
	// voice-12345
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The speech rate for TTS playback. Valid values: -200 to 200. Default value: 0.
	//
	// example:
	//
	// 0
	TtsSpeed *int64 `json:"TtsSpeed,omitempty" xml:"TtsSpeed,omitempty"`
	// The voice style.
	//
	// example:
	//
	// longxiaoxia_v2p1
	TtsStyle *string `json:"TtsStyle,omitempty" xml:"TtsStyle,omitempty"`
	// The volume of TTS playback. Valid values: 0 to 100.
	//
	// example:
	//
	// 88
	TtsVolume *int64 `json:"TtsVolume,omitempty" xml:"TtsVolume,omitempty"`
	// The voice code.
	//
	// example:
	//
	// V123ABC00
	VoiceCode *string `json:"VoiceCode,omitempty" xml:"VoiceCode,omitempty"`
	// The type of the TTS voice. Valid values:
	//
	// `VOICE_TYPE_SYSTEM`: a system voice.
	//
	// `VOICE_TYPE_CLONE`: a cloned voice.
	//
	// `VOICE_TYPE_DOUBAO`: a Doubao voice.
	//
	// `VOICE_TYPE_MINIMAX`: a Minimax voice.
	//
	// `VOICE_TYPE_OPENTTS`: an open voice.
	//
	// `VOICE_TYPE_BL_CUSTOM`: a high-quality custom cloned voice.
	//
	// example:
	//
	// VOICE_TYPE_SYSTEM
	VoiceType *bool `json:"VoiceType,omitempty" xml:"VoiceType,omitempty"`
}

func (s QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig) GetBackgroundEnabled() *bool {
	return s.BackgroundEnabled
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig) GetBackgroundSound() *int64 {
	return s.BackgroundSound
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig) GetBackgroundVolume() *int64 {
	return s.BackgroundVolume
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig) GetCustomerAccountId() *int64 {
	return s.CustomerAccountId
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig) GetMixingEnabled() *bool {
	return s.MixingEnabled
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig) GetMixingTemplate() *int64 {
	return s.MixingTemplate
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig) GetResourceId() *string {
	return s.ResourceId
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig) GetTtsSpeed() *int64 {
	return s.TtsSpeed
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig) GetTtsStyle() *string {
	return s.TtsStyle
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig) GetTtsVolume() *int64 {
	return s.TtsVolume
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig) GetVoiceCode() *string {
	return s.VoiceCode
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig) GetVoiceType() *bool {
	return s.VoiceType
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig) SetBackgroundEnabled(v bool) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig {
	s.BackgroundEnabled = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig) SetBackgroundSound(v int64) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig {
	s.BackgroundSound = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig) SetBackgroundVolume(v int64) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig {
	s.BackgroundVolume = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig) SetCustomerAccountId(v int64) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig {
	s.CustomerAccountId = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig) SetMixingEnabled(v bool) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig {
	s.MixingEnabled = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig) SetMixingTemplate(v int64) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig {
	s.MixingTemplate = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig) SetResourceId(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig {
	s.ResourceId = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig) SetTtsSpeed(v int64) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig {
	s.TtsSpeed = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig) SetTtsStyle(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig {
	s.TtsStyle = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig) SetTtsVolume(v int64) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig {
	s.TtsVolume = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig) SetVoiceCode(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig {
	s.VoiceCode = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig) SetVoiceType(v bool) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig {
	s.VoiceType = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentCallConfigTtsConfig) Validate() error {
	return dara.Validate(s)
}

type QueryAiVoiceAgentDetailNewResponseBodyDataAgentDemandConfig struct {
	// Specifies whether the agent was built with AI assistance.
	//
	// example:
	//
	// false
	AiGenerate *bool `json:"AiGenerate,omitempty" xml:"AiGenerate,omitempty"`
	// The basic task configuration.
	//
	// example:
	//
	// 进行产品营销，介绍优势特点
	BasicTaskDescription *string `json:"BasicTaskDescription,omitempty" xml:"BasicTaskDescription,omitempty"`
	// The business scenario.
	//
	// example:
	//
	// 1
	BusinessType *int64 `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The core objective.
	//
	// example:
	//
	// 销售线索转化
	CoreTarget *string `json:"CoreTarget,omitempty" xml:"CoreTarget,omitempty"`
	// The system role.
	//
	// example:
	//
	// 销售
	SysRole *string `json:"SysRole,omitempty" xml:"SysRole,omitempty"`
	// The user role.
	//
	// example:
	//
	// 潜在客户
	UserRole *string `json:"UserRole,omitempty" xml:"UserRole,omitempty"`
}

func (s QueryAiVoiceAgentDetailNewResponseBodyDataAgentDemandConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailNewResponseBodyDataAgentDemandConfig) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentDemandConfig) GetAiGenerate() *bool {
	return s.AiGenerate
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentDemandConfig) GetBasicTaskDescription() *string {
	return s.BasicTaskDescription
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentDemandConfig) GetBusinessType() *int64 {
	return s.BusinessType
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentDemandConfig) GetCoreTarget() *string {
	return s.CoreTarget
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentDemandConfig) GetSysRole() *string {
	return s.SysRole
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentDemandConfig) GetUserRole() *string {
	return s.UserRole
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentDemandConfig) SetAiGenerate(v bool) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentDemandConfig {
	s.AiGenerate = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentDemandConfig) SetBasicTaskDescription(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentDemandConfig {
	s.BasicTaskDescription = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentDemandConfig) SetBusinessType(v int64) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentDemandConfig {
	s.BusinessType = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentDemandConfig) SetCoreTarget(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentDemandConfig {
	s.CoreTarget = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentDemandConfig) SetSysRole(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentDemandConfig {
	s.SysRole = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentDemandConfig) SetUserRole(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataAgentDemandConfig {
	s.UserRole = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataAgentDemandConfig) Validate() error {
	return dara.Validate(s)
}

type QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfig struct {
	// The collection of knowledge bases.
	KnowledgeIds []*QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfigKnowledgeIds `json:"KnowledgeIds,omitempty" xml:"KnowledgeIds,omitempty" type:"Repeated"`
	// The Retrieval-Augmented Generation (RAG) configuration.
	RagConfig *QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfigRagConfig `json:"RagConfig,omitempty" xml:"RagConfig,omitempty" type:"Struct"`
}

func (s QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfig) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfig) GetKnowledgeIds() []*QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfigKnowledgeIds {
	return s.KnowledgeIds
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfig) GetRagConfig() *QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfigRagConfig {
	return s.RagConfig
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfig) SetKnowledgeIds(v []*QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfigKnowledgeIds) *QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfig {
	s.KnowledgeIds = v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfig) SetRagConfig(v *QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfigRagConfig) *QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfig {
	s.RagConfig = v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfig) Validate() error {
	if s.KnowledgeIds != nil {
		for _, item := range s.KnowledgeIds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RagConfig != nil {
		if err := s.RagConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfigKnowledgeIds struct {
	// The ID of the knowledge base.
	//
	// example:
	//
	// 68
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// The name of the knowledge base.
	//
	// example:
	//
	// 产品知识库
	KnowledgeName *string `json:"KnowledgeName,omitempty" xml:"KnowledgeName,omitempty"`
}

func (s QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfigKnowledgeIds) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfigKnowledgeIds) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfigKnowledgeIds) GetKnowledgeId() *int64 {
	return s.KnowledgeId
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfigKnowledgeIds) GetKnowledgeName() *string {
	return s.KnowledgeName
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfigKnowledgeIds) SetKnowledgeId(v int64) *QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfigKnowledgeIds {
	s.KnowledgeId = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfigKnowledgeIds) SetKnowledgeName(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfigKnowledgeIds {
	s.KnowledgeName = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfigKnowledgeIds) Validate() error {
	return dara.Validate(s)
}

type QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfigRagConfig struct {
	// RAG retrieval description
	//
	// example:
	//
	// 知识库检索
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Enable RAG retrieval
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfigRagConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfigRagConfig) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfigRagConfig) GetDescription() *string {
	return s.Description
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfigRagConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfigRagConfig) SetDescription(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfigRagConfig {
	s.Description = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfigRagConfig) SetEnabled(v bool) *QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfigRagConfig {
	s.Enabled = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataKnowledgeConfigRagConfig) Validate() error {
	return dara.Validate(s)
}

type QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfig struct {
	// The ID of the call variable.
	//
	// example:
	//
	// 123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The enumerated values for the call variable.
	PhoneTagEnum []*QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfigPhoneTagEnum `json:"PhoneTagEnum,omitempty" xml:"PhoneTagEnum,omitempty" type:"Repeated"`
	// The key of the call variable.
	//
	// example:
	//
	// gender
	PhoneTagKey *string `json:"PhoneTagKey,omitempty" xml:"PhoneTagKey,omitempty"`
	// The description of the call variable.
	//
	// example:
	//
	// 用户性别
	PhoneTagName *string `json:"PhoneTagName,omitempty" xml:"PhoneTagName,omitempty"`
	// Specifies whether the call variable is required.
	//
	// example:
	//
	// false
	PhoneTagRequired *bool `json:"PhoneTagRequired,omitempty" xml:"PhoneTagRequired,omitempty"`
	// The source of the call variable.
	//
	// example:
	//
	// MANUAL_CREATE
	PhoneTagSource *string `json:"PhoneTagSource,omitempty" xml:"PhoneTagSource,omitempty"`
	// The type of the call variable. Valid values: `TEXT` and `ENUM`.
	//
	// example:
	//
	// ENUM
	PhoneTagType *string `json:"PhoneTagType,omitempty" xml:"PhoneTagType,omitempty"`
}

func (s QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfig) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfig) GetId() *string {
	return s.Id
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfig) GetPhoneTagEnum() []*QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfigPhoneTagEnum {
	return s.PhoneTagEnum
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfig) GetPhoneTagKey() *string {
	return s.PhoneTagKey
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfig) GetPhoneTagName() *string {
	return s.PhoneTagName
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfig) GetPhoneTagRequired() *bool {
	return s.PhoneTagRequired
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfig) GetPhoneTagSource() *string {
	return s.PhoneTagSource
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfig) GetPhoneTagType() *string {
	return s.PhoneTagType
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfig) SetId(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfig {
	s.Id = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfig) SetPhoneTagEnum(v []*QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfigPhoneTagEnum) *QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfig {
	s.PhoneTagEnum = v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfig) SetPhoneTagKey(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfig {
	s.PhoneTagKey = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfig) SetPhoneTagName(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfig {
	s.PhoneTagName = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfig) SetPhoneTagRequired(v bool) *QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfig {
	s.PhoneTagRequired = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfig) SetPhoneTagSource(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfig {
	s.PhoneTagSource = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfig) SetPhoneTagType(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfig {
	s.PhoneTagType = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfig) Validate() error {
	if s.PhoneTagEnum != nil {
		for _, item := range s.PhoneTagEnum {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfigPhoneTagEnum struct {
	// The description.
	//
	// example:
	//
	// gender
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The tag ID.
	//
	// example:
	//
	// 123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The tag value.
	//
	// example:
	//
	// male
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfigPhoneTagEnum) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfigPhoneTagEnum) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfigPhoneTagEnum) GetDescription() *string {
	return s.Description
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfigPhoneTagEnum) GetId() *string {
	return s.Id
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfigPhoneTagEnum) GetValue() *string {
	return s.Value
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfigPhoneTagEnum) SetDescription(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfigPhoneTagEnum {
	s.Description = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfigPhoneTagEnum) SetId(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfigPhoneTagEnum {
	s.Id = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfigPhoneTagEnum) SetValue(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfigPhoneTagEnum {
	s.Value = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataPhoneTagConfigPhoneTagEnum) Validate() error {
	return dara.Validate(s)
}

type QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfig struct {
	// The configuration for mapping call results to tags.
	CallResultTagConfig *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigCallResultTagConfig `json:"CallResultTagConfig,omitempty" xml:"CallResultTagConfig,omitempty" type:"Struct"`
	// The main intent.
	MainPurpose *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurpose `json:"MainPurpose,omitempty" xml:"MainPurpose,omitempty" type:"Struct"`
	// The output tag configuration.
	OutputTagConfig []*QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfig `json:"OutputTagConfig,omitempty" xml:"OutputTagConfig,omitempty" type:"Repeated"`
}

func (s QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfig) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfig) GetCallResultTagConfig() *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigCallResultTagConfig {
	return s.CallResultTagConfig
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfig) GetMainPurpose() *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurpose {
	return s.MainPurpose
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfig) GetOutputTagConfig() []*QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfig {
	return s.OutputTagConfig
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfig) SetCallResultTagConfig(v *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigCallResultTagConfig) *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfig {
	s.CallResultTagConfig = v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfig) SetMainPurpose(v *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurpose) *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfig {
	s.MainPurpose = v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfig) SetOutputTagConfig(v []*QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfig) *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfig {
	s.OutputTagConfig = v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfig) Validate() error {
	if s.CallResultTagConfig != nil {
		if err := s.CallResultTagConfig.Validate(); err != nil {
			return err
		}
	}
	if s.MainPurpose != nil {
		if err := s.MainPurpose.Validate(); err != nil {
			return err
		}
	}
	if s.OutputTagConfig != nil {
		for _, item := range s.OutputTagConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigCallResultTagConfig struct {
	// The default fallback tag information.
	DefaultTag *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigCallResultTagConfigDefaultTag `json:"DefaultTag,omitempty" xml:"DefaultTag,omitempty" type:"Struct"`
	// The mapping of call results to tags.
	MappingTag map[string]*string `json:"MappingTag,omitempty" xml:"MappingTag,omitempty"`
}

func (s QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigCallResultTagConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigCallResultTagConfig) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigCallResultTagConfig) GetDefaultTag() *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigCallResultTagConfigDefaultTag {
	return s.DefaultTag
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigCallResultTagConfig) GetMappingTag() map[string]*string {
	return s.MappingTag
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigCallResultTagConfig) SetDefaultTag(v *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigCallResultTagConfigDefaultTag) *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigCallResultTagConfig {
	s.DefaultTag = v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigCallResultTagConfig) SetMappingTag(v map[string]*string) *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigCallResultTagConfig {
	s.MappingTag = v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigCallResultTagConfig) Validate() error {
	if s.DefaultTag != nil {
		if err := s.DefaultTag.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigCallResultTagConfigDefaultTag struct {
	// The description.
	//
	// example:
	//
	// 标签说明
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The tag.
	//
	// example:
	//
	// 标签1
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigCallResultTagConfigDefaultTag) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigCallResultTagConfigDefaultTag) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigCallResultTagConfigDefaultTag) GetDesc() *string {
	return s.Desc
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigCallResultTagConfigDefaultTag) GetTag() *string {
	return s.Tag
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigCallResultTagConfigDefaultTag) SetDesc(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigCallResultTagConfigDefaultTag {
	s.Desc = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigCallResultTagConfigDefaultTag) SetTag(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigCallResultTagConfigDefaultTag {
	s.Tag = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigCallResultTagConfigDefaultTag) Validate() error {
	return dara.Validate(s)
}

type QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurpose struct {
	// The ID of the main intent.
	//
	// example:
	//
	// 123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The description of the main intent.
	//
	// example:
	//
	// Description of the main intent
	MainPurposeDescription *string `json:"MainPurposeDescription,omitempty" xml:"MainPurposeDescription,omitempty"`
	// The enumerated values for the main intent.
	MainPurposeEnum []*QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurposeMainPurposeEnum `json:"MainPurposeEnum,omitempty" xml:"MainPurposeEnum,omitempty" type:"Repeated"`
	// The name of the main intent.
	//
	// example:
	//
	// Customer inquiry handling
	MainPurposeName *string `json:"MainPurposeName,omitempty" xml:"MainPurposeName,omitempty"`
	// The value type of the main intent. Valid values: `TEXT` and `ENUM`.
	//
	// example:
	//
	// ENUM
	MainPurposeType *string `json:"MainPurposeType,omitempty" xml:"MainPurposeType,omitempty"`
}

func (s QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurpose) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurpose) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurpose) GetId() *string {
	return s.Id
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurpose) GetMainPurposeDescription() *string {
	return s.MainPurposeDescription
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurpose) GetMainPurposeEnum() []*QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurposeMainPurposeEnum {
	return s.MainPurposeEnum
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurpose) GetMainPurposeName() *string {
	return s.MainPurposeName
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurpose) GetMainPurposeType() *string {
	return s.MainPurposeType
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurpose) SetId(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurpose {
	s.Id = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurpose) SetMainPurposeDescription(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurpose {
	s.MainPurposeDescription = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurpose) SetMainPurposeEnum(v []*QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurposeMainPurposeEnum) *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurpose {
	s.MainPurposeEnum = v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurpose) SetMainPurposeName(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurpose {
	s.MainPurposeName = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurpose) SetMainPurposeType(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurpose {
	s.MainPurposeType = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurpose) Validate() error {
	if s.MainPurposeEnum != nil {
		for _, item := range s.MainPurposeEnum {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurposeMainPurposeEnum struct {
	// The description of the value.
	//
	// example:
	//
	// Description of the tag value
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The unique ID of the tag.
	//
	// example:
	//
	// tag_12345
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The value.
	//
	// example:
	//
	// Tag value example
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurposeMainPurposeEnum) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurposeMainPurposeEnum) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurposeMainPurposeEnum) GetDescription() *string {
	return s.Description
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurposeMainPurposeEnum) GetId() *string {
	return s.Id
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurposeMainPurposeEnum) GetValue() *string {
	return s.Value
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurposeMainPurposeEnum) SetDescription(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurposeMainPurposeEnum {
	s.Description = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurposeMainPurposeEnum) SetId(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurposeMainPurposeEnum {
	s.Id = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurposeMainPurposeEnum) SetValue(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurposeMainPurposeEnum {
	s.Value = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigMainPurposeMainPurposeEnum) Validate() error {
	return dara.Validate(s)
}

type QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfig struct {
	// The ID of the output tag.
	//
	// example:
	//
	// 123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The description of the output tag.
	//
	// example:
	//
	// 用户购买产品的意愿
	OutputTagDescription *string `json:"OutputTagDescription,omitempty" xml:"OutputTagDescription,omitempty"`
	// The enumerated values for the output tag.
	OutputTagEnum []*QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfigOutputTagEnum `json:"OutputTagEnum,omitempty" xml:"OutputTagEnum,omitempty" type:"Repeated"`
	// The name of the output tag.
	//
	// example:
	//
	// 购买意愿
	OutputTagName *string `json:"OutputTagName,omitempty" xml:"OutputTagName,omitempty"`
	// The value type of the output tag. Valid values: `TEXT` and `ENUM`.
	//
	// example:
	//
	// ENUM
	OutputTagType *string `json:"OutputTagType,omitempty" xml:"OutputTagType,omitempty"`
}

func (s QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfig) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfig) GetId() *string {
	return s.Id
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfig) GetOutputTagDescription() *string {
	return s.OutputTagDescription
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfig) GetOutputTagEnum() []*QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfigOutputTagEnum {
	return s.OutputTagEnum
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfig) GetOutputTagName() *string {
	return s.OutputTagName
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfig) GetOutputTagType() *string {
	return s.OutputTagType
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfig) SetId(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfig {
	s.Id = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfig) SetOutputTagDescription(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfig {
	s.OutputTagDescription = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfig) SetOutputTagEnum(v []*QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfigOutputTagEnum) *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfig {
	s.OutputTagEnum = v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfig) SetOutputTagName(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfig {
	s.OutputTagName = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfig) SetOutputTagType(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfig {
	s.OutputTagType = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfig) Validate() error {
	if s.OutputTagEnum != nil {
		for _, item := range s.OutputTagEnum {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfigOutputTagEnum struct {
	// The tag description.
	//
	// example:
	//
	// 用户有较高购买意愿
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The tag ID.
	//
	// example:
	//
	// 10
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The enumerated value of the tag.
	//
	// example:
	//
	// 高意愿
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfigOutputTagEnum) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfigOutputTagEnum) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfigOutputTagEnum) GetDescription() *string {
	return s.Description
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfigOutputTagEnum) GetId() *string {
	return s.Id
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfigOutputTagEnum) GetValue() *string {
	return s.Value
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfigOutputTagEnum) SetDescription(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfigOutputTagEnum {
	s.Description = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfigOutputTagEnum) SetId(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfigOutputTagEnum {
	s.Id = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfigOutputTagEnum) SetValue(v string) *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfigOutputTagEnum {
	s.Value = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponseBodyDataSummaryConfigOutputTagConfigOutputTagEnum) Validate() error {
	return dara.Validate(s)
}
