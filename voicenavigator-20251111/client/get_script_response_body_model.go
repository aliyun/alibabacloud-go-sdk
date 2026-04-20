// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScriptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetScriptResponseBody
	GetCode() *string
	SetData(v *GetScriptResponseBodyData) *GetScriptResponseBody
	GetData() *GetScriptResponseBodyData
	SetHttpStatusCode(v int32) *GetScriptResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetScriptResponseBody
	GetMessage() *string
	SetParams(v []*string) *GetScriptResponseBody
	GetParams() []*string
	SetRequestId(v string) *GetScriptResponseBody
	GetRequestId() *string
}

type GetScriptResponseBody struct {
	// example:
	//
	// OK
	Code *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetScriptResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Instance af81a389-91f0-4157-8d82-720edd02b66a
	//
	//  does not exist.
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 51E08AA9-8D1F-55F8-84A3-40635E2F0806
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetScriptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetScriptResponseBody) GoString() string {
	return s.String()
}

func (s *GetScriptResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetScriptResponseBody) GetData() *GetScriptResponseBodyData {
	return s.Data
}

func (s *GetScriptResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetScriptResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetScriptResponseBody) GetParams() []*string {
	return s.Params
}

func (s *GetScriptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetScriptResponseBody) SetCode(v string) *GetScriptResponseBody {
	s.Code = &v
	return s
}

func (s *GetScriptResponseBody) SetData(v *GetScriptResponseBodyData) *GetScriptResponseBody {
	s.Data = v
	return s
}

func (s *GetScriptResponseBody) SetHttpStatusCode(v int32) *GetScriptResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetScriptResponseBody) SetMessage(v string) *GetScriptResponseBody {
	s.Message = &v
	return s
}

func (s *GetScriptResponseBody) SetParams(v []*string) *GetScriptResponseBody {
	s.Params = v
	return s
}

func (s *GetScriptResponseBody) SetRequestId(v string) *GetScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScriptResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetScriptResponseBodyData struct {
	// example:
	//
	// 10
	Concurrency *int32 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	// example:
	//
	// 1773228988000
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// test script
	Description  *string                                `json:"Description,omitempty" xml:"Description,omitempty"`
	DraftVersion *GetScriptResponseBodyDataDraftVersion `json:"DraftVersion,omitempty" xml:"DraftVersion,omitempty" type:"Struct"`
	Name         *string                                `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// BEEBOT
	NluEngine        *string                                    `json:"NluEngine,omitempty" xml:"NluEngine,omitempty"`
	PublishedVersion *GetScriptResponseBodyDataPublishedVersion `json:"PublishedVersion,omitempty" xml:"PublishedVersion,omitempty" type:"Struct"`
	// example:
	//
	// 64241e64-190c-45d1-af66-06f51c07b090
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// example:
	//
	// DRAFT
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1773228988000
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s GetScriptResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetScriptResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetScriptResponseBodyData) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *GetScriptResponseBodyData) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *GetScriptResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetScriptResponseBodyData) GetDraftVersion() *GetScriptResponseBodyDataDraftVersion {
	return s.DraftVersion
}

func (s *GetScriptResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetScriptResponseBodyData) GetNluEngine() *string {
	return s.NluEngine
}

func (s *GetScriptResponseBodyData) GetPublishedVersion() *GetScriptResponseBodyDataPublishedVersion {
	return s.PublishedVersion
}

func (s *GetScriptResponseBodyData) GetScriptId() *string {
	return s.ScriptId
}

func (s *GetScriptResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetScriptResponseBodyData) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *GetScriptResponseBodyData) SetConcurrency(v int32) *GetScriptResponseBodyData {
	s.Concurrency = &v
	return s
}

func (s *GetScriptResponseBodyData) SetCreatedTime(v int64) *GetScriptResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *GetScriptResponseBodyData) SetDescription(v string) *GetScriptResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetScriptResponseBodyData) SetDraftVersion(v *GetScriptResponseBodyDataDraftVersion) *GetScriptResponseBodyData {
	s.DraftVersion = v
	return s
}

func (s *GetScriptResponseBodyData) SetName(v string) *GetScriptResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetScriptResponseBodyData) SetNluEngine(v string) *GetScriptResponseBodyData {
	s.NluEngine = &v
	return s
}

func (s *GetScriptResponseBodyData) SetPublishedVersion(v *GetScriptResponseBodyDataPublishedVersion) *GetScriptResponseBodyData {
	s.PublishedVersion = v
	return s
}

func (s *GetScriptResponseBodyData) SetScriptId(v string) *GetScriptResponseBodyData {
	s.ScriptId = &v
	return s
}

func (s *GetScriptResponseBodyData) SetStatus(v string) *GetScriptResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetScriptResponseBodyData) SetUpdatedTime(v int64) *GetScriptResponseBodyData {
	s.UpdatedTime = &v
	return s
}

func (s *GetScriptResponseBodyData) Validate() error {
	if s.DraftVersion != nil {
		if err := s.DraftVersion.Validate(); err != nil {
			return err
		}
	}
	if s.PublishedVersion != nil {
		if err := s.PublishedVersion.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetScriptResponseBodyDataDraftVersion struct {
	InteractionConfig *GetScriptResponseBodyDataDraftVersionInteractionConfig `json:"InteractionConfig,omitempty" xml:"InteractionConfig,omitempty" type:"Struct"`
	LabelConfig       []*GetScriptResponseBodyDataDraftVersionLabelConfig     `json:"LabelConfig,omitempty" xml:"LabelConfig,omitempty" type:"Repeated"`
	ScriptProfile     *GetScriptResponseBodyDataDraftVersionScriptProfile     `json:"ScriptProfile,omitempty" xml:"ScriptProfile,omitempty" type:"Struct"`
	SynthesizerConfig *GetScriptResponseBodyDataDraftVersionSynthesizerConfig `json:"SynthesizerConfig,omitempty" xml:"SynthesizerConfig,omitempty" type:"Struct"`
	TranscriberConfig *GetScriptResponseBodyDataDraftVersionTranscriberConfig `json:"TranscriberConfig,omitempty" xml:"TranscriberConfig,omitempty" type:"Struct"`
	// example:
	//
	// 0c4f978a-73bb-4841-bd84-75c0398edd4e
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetScriptResponseBodyDataDraftVersion) String() string {
	return dara.Prettify(s)
}

func (s GetScriptResponseBodyDataDraftVersion) GoString() string {
	return s.String()
}

func (s *GetScriptResponseBodyDataDraftVersion) GetInteractionConfig() *GetScriptResponseBodyDataDraftVersionInteractionConfig {
	return s.InteractionConfig
}

func (s *GetScriptResponseBodyDataDraftVersion) GetLabelConfig() []*GetScriptResponseBodyDataDraftVersionLabelConfig {
	return s.LabelConfig
}

func (s *GetScriptResponseBodyDataDraftVersion) GetScriptProfile() *GetScriptResponseBodyDataDraftVersionScriptProfile {
	return s.ScriptProfile
}

func (s *GetScriptResponseBodyDataDraftVersion) GetSynthesizerConfig() *GetScriptResponseBodyDataDraftVersionSynthesizerConfig {
	return s.SynthesizerConfig
}

func (s *GetScriptResponseBodyDataDraftVersion) GetTranscriberConfig() *GetScriptResponseBodyDataDraftVersionTranscriberConfig {
	return s.TranscriberConfig
}

func (s *GetScriptResponseBodyDataDraftVersion) GetVersionId() *string {
	return s.VersionId
}

func (s *GetScriptResponseBodyDataDraftVersion) SetInteractionConfig(v *GetScriptResponseBodyDataDraftVersionInteractionConfig) *GetScriptResponseBodyDataDraftVersion {
	s.InteractionConfig = v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersion) SetLabelConfig(v []*GetScriptResponseBodyDataDraftVersionLabelConfig) *GetScriptResponseBodyDataDraftVersion {
	s.LabelConfig = v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersion) SetScriptProfile(v *GetScriptResponseBodyDataDraftVersionScriptProfile) *GetScriptResponseBodyDataDraftVersion {
	s.ScriptProfile = v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersion) SetSynthesizerConfig(v *GetScriptResponseBodyDataDraftVersionSynthesizerConfig) *GetScriptResponseBodyDataDraftVersion {
	s.SynthesizerConfig = v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersion) SetTranscriberConfig(v *GetScriptResponseBodyDataDraftVersionTranscriberConfig) *GetScriptResponseBodyDataDraftVersion {
	s.TranscriberConfig = v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersion) SetVersionId(v string) *GetScriptResponseBodyDataDraftVersion {
	s.VersionId = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersion) Validate() error {
	if s.InteractionConfig != nil {
		if err := s.InteractionConfig.Validate(); err != nil {
			return err
		}
	}
	if s.LabelConfig != nil {
		for _, item := range s.LabelConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ScriptProfile != nil {
		if err := s.ScriptProfile.Validate(); err != nil {
			return err
		}
	}
	if s.SynthesizerConfig != nil {
		if err := s.SynthesizerConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TranscriberConfig != nil {
		if err := s.TranscriberConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetScriptResponseBodyDataDraftVersionInteractionConfig struct {
	// example:
	//
	// office-ambience
	BackgroundMusicId     *string                                                                      `json:"BackgroundMusicId,omitempty" xml:"BackgroundMusicId,omitempty"`
	EndConversationConfig *GetScriptResponseBodyDataDraftVersionInteractionConfigEndConversationConfig `json:"EndConversationConfig,omitempty" xml:"EndConversationConfig,omitempty" type:"Struct"`
	// example:
	//
	// 2000
	InitialGreetingDelayMilliseconds *int32                                                                        `json:"InitialGreetingDelayMilliseconds,omitempty" xml:"InitialGreetingDelayMilliseconds,omitempty"`
	SilenceDetectionConfig           *GetScriptResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig `json:"SilenceDetectionConfig,omitempty" xml:"SilenceDetectionConfig,omitempty" type:"Struct"`
}

func (s GetScriptResponseBodyDataDraftVersionInteractionConfig) String() string {
	return dara.Prettify(s)
}

func (s GetScriptResponseBodyDataDraftVersionInteractionConfig) GoString() string {
	return s.String()
}

func (s *GetScriptResponseBodyDataDraftVersionInteractionConfig) GetBackgroundMusicId() *string {
	return s.BackgroundMusicId
}

func (s *GetScriptResponseBodyDataDraftVersionInteractionConfig) GetEndConversationConfig() *GetScriptResponseBodyDataDraftVersionInteractionConfigEndConversationConfig {
	return s.EndConversationConfig
}

func (s *GetScriptResponseBodyDataDraftVersionInteractionConfig) GetInitialGreetingDelayMilliseconds() *int32 {
	return s.InitialGreetingDelayMilliseconds
}

func (s *GetScriptResponseBodyDataDraftVersionInteractionConfig) GetSilenceDetectionConfig() *GetScriptResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig {
	return s.SilenceDetectionConfig
}

func (s *GetScriptResponseBodyDataDraftVersionInteractionConfig) SetBackgroundMusicId(v string) *GetScriptResponseBodyDataDraftVersionInteractionConfig {
	s.BackgroundMusicId = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionInteractionConfig) SetEndConversationConfig(v *GetScriptResponseBodyDataDraftVersionInteractionConfigEndConversationConfig) *GetScriptResponseBodyDataDraftVersionInteractionConfig {
	s.EndConversationConfig = v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionInteractionConfig) SetInitialGreetingDelayMilliseconds(v int32) *GetScriptResponseBodyDataDraftVersionInteractionConfig {
	s.InitialGreetingDelayMilliseconds = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionInteractionConfig) SetSilenceDetectionConfig(v *GetScriptResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig) *GetScriptResponseBodyDataDraftVersionInteractionConfig {
	s.SilenceDetectionConfig = v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionInteractionConfig) Validate() error {
	if s.EndConversationConfig != nil {
		if err := s.EndConversationConfig.Validate(); err != nil {
			return err
		}
	}
	if s.SilenceDetectionConfig != nil {
		if err := s.SilenceDetectionConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetScriptResponseBodyDataDraftVersionInteractionConfigEndConversationConfig struct {
	// example:
	//
	// 1
	Delay    *int32                                                                                 `json:"Delay,omitempty" xml:"Delay,omitempty"`
	Triggers []*GetScriptResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers `json:"Triggers,omitempty" xml:"Triggers,omitempty" type:"Repeated"`
}

func (s GetScriptResponseBodyDataDraftVersionInteractionConfigEndConversationConfig) String() string {
	return dara.Prettify(s)
}

func (s GetScriptResponseBodyDataDraftVersionInteractionConfigEndConversationConfig) GoString() string {
	return s.String()
}

func (s *GetScriptResponseBodyDataDraftVersionInteractionConfigEndConversationConfig) GetDelay() *int32 {
	return s.Delay
}

func (s *GetScriptResponseBodyDataDraftVersionInteractionConfigEndConversationConfig) GetTriggers() []*GetScriptResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers {
	return s.Triggers
}

func (s *GetScriptResponseBodyDataDraftVersionInteractionConfigEndConversationConfig) SetDelay(v int32) *GetScriptResponseBodyDataDraftVersionInteractionConfigEndConversationConfig {
	s.Delay = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionInteractionConfigEndConversationConfig) SetTriggers(v []*GetScriptResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers) *GetScriptResponseBodyDataDraftVersionInteractionConfigEndConversationConfig {
	s.Triggers = v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionInteractionConfigEndConversationConfig) Validate() error {
	if s.Triggers != nil {
		for _, item := range s.Triggers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetScriptResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers struct {
	ClosingStatement *string   `json:"ClosingStatement,omitempty" xml:"ClosingStatement,omitempty"`
	KeyWords         []*string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty" type:"Repeated"`
	// example:
	//
	// TurnLimit
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// example:
	//
	// 20
	TurnLimit *int32 `json:"TurnLimit,omitempty" xml:"TurnLimit,omitempty"`
}

func (s GetScriptResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers) String() string {
	return dara.Prettify(s)
}

func (s GetScriptResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers) GoString() string {
	return s.String()
}

func (s *GetScriptResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers) GetClosingStatement() *string {
	return s.ClosingStatement
}

func (s *GetScriptResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers) GetKeyWords() []*string {
	return s.KeyWords
}

func (s *GetScriptResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers) GetTriggerType() *string {
	return s.TriggerType
}

func (s *GetScriptResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers) GetTurnLimit() *int32 {
	return s.TurnLimit
}

func (s *GetScriptResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers) SetClosingStatement(v string) *GetScriptResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers {
	s.ClosingStatement = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers) SetKeyWords(v []*string) *GetScriptResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers {
	s.KeyWords = v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers) SetTriggerType(v string) *GetScriptResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers {
	s.TriggerType = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers) SetTurnLimit(v int32) *GetScriptResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers {
	s.TurnLimit = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers) Validate() error {
	return dara.Validate(s)
}

type GetScriptResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig struct {
	// example:
	//
	// 3
	MaxRepeats *int32 `json:"MaxRepeats,omitempty" xml:"MaxRepeats,omitempty"`
	// example:
	//
	// 5000
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s GetScriptResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig) String() string {
	return dara.Prettify(s)
}

func (s GetScriptResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig) GoString() string {
	return s.String()
}

func (s *GetScriptResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig) GetMaxRepeats() *int32 {
	return s.MaxRepeats
}

func (s *GetScriptResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig) GetTimeout() *int32 {
	return s.Timeout
}

func (s *GetScriptResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig) SetMaxRepeats(v int32) *GetScriptResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig {
	s.MaxRepeats = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig) SetTimeout(v int32) *GetScriptResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig {
	s.Timeout = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig) Validate() error {
	return dara.Validate(s)
}

type GetScriptResponseBodyDataDraftVersionLabelConfig struct {
	CandidateValues []*string `json:"CandidateValues,omitempty" xml:"CandidateValues,omitempty" type:"Repeated"`
	Description     *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Name            *string   `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetScriptResponseBodyDataDraftVersionLabelConfig) String() string {
	return dara.Prettify(s)
}

func (s GetScriptResponseBodyDataDraftVersionLabelConfig) GoString() string {
	return s.String()
}

func (s *GetScriptResponseBodyDataDraftVersionLabelConfig) GetCandidateValues() []*string {
	return s.CandidateValues
}

func (s *GetScriptResponseBodyDataDraftVersionLabelConfig) GetDescription() *string {
	return s.Description
}

func (s *GetScriptResponseBodyDataDraftVersionLabelConfig) GetName() *string {
	return s.Name
}

func (s *GetScriptResponseBodyDataDraftVersionLabelConfig) SetCandidateValues(v []*string) *GetScriptResponseBodyDataDraftVersionLabelConfig {
	s.CandidateValues = v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionLabelConfig) SetDescription(v string) *GetScriptResponseBodyDataDraftVersionLabelConfig {
	s.Description = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionLabelConfig) SetName(v string) *GetScriptResponseBodyDataDraftVersionLabelConfig {
	s.Name = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionLabelConfig) Validate() error {
	return dara.Validate(s)
}

type GetScriptResponseBodyDataDraftVersionScriptProfile struct {
	// example:
	//
	// 1309723684579735_p_beebot_public
	AgentKey     *string                                                         `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	AgentProfile *GetScriptResponseBodyDataDraftVersionScriptProfileAgentProfile `json:"AgentProfile,omitempty" xml:"AgentProfile,omitempty" type:"Struct"`
	// example:
	//
	// chatbot-cn-MQuyjjb666
	ChatbotId    *string                                                         `json:"ChatbotId,omitempty" xml:"ChatbotId,omitempty"`
	FunctionMeta *GetScriptResponseBodyDataDraftVersionScriptProfileFunctionMeta `json:"FunctionMeta,omitempty" xml:"FunctionMeta,omitempty" type:"Struct"`
	// example:
	//
	// qwen-plus
	Model            *string                                                             `json:"Model,omitempty" xml:"Model,omitempty"`
	NluAccessProfile *GetScriptResponseBodyDataDraftVersionScriptProfileNluAccessProfile `json:"NluAccessProfile,omitempty" xml:"NluAccessProfile,omitempty" type:"Struct"`
	// example:
	//
	// MANAGED
	NluAccessType *string `json:"NluAccessType,omitempty" xml:"NluAccessType,omitempty"`
	// example:
	//
	// BEEBOT
	NluEngine *string `json:"NluEngine,omitempty" xml:"NluEngine,omitempty"`
	// example:
	//
	// true
	OmniModel *bool `json:"OmniModel,omitempty" xml:"OmniModel,omitempty"`
}

func (s GetScriptResponseBodyDataDraftVersionScriptProfile) String() string {
	return dara.Prettify(s)
}

func (s GetScriptResponseBodyDataDraftVersionScriptProfile) GoString() string {
	return s.String()
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfile) GetAgentKey() *string {
	return s.AgentKey
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfile) GetAgentProfile() *GetScriptResponseBodyDataDraftVersionScriptProfileAgentProfile {
	return s.AgentProfile
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfile) GetChatbotId() *string {
	return s.ChatbotId
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfile) GetFunctionMeta() *GetScriptResponseBodyDataDraftVersionScriptProfileFunctionMeta {
	return s.FunctionMeta
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfile) GetModel() *string {
	return s.Model
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfile) GetNluAccessProfile() *GetScriptResponseBodyDataDraftVersionScriptProfileNluAccessProfile {
	return s.NluAccessProfile
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfile) GetNluAccessType() *string {
	return s.NluAccessType
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfile) GetNluEngine() *string {
	return s.NluEngine
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfile) GetOmniModel() *bool {
	return s.OmniModel
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfile) SetAgentKey(v string) *GetScriptResponseBodyDataDraftVersionScriptProfile {
	s.AgentKey = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfile) SetAgentProfile(v *GetScriptResponseBodyDataDraftVersionScriptProfileAgentProfile) *GetScriptResponseBodyDataDraftVersionScriptProfile {
	s.AgentProfile = v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfile) SetChatbotId(v string) *GetScriptResponseBodyDataDraftVersionScriptProfile {
	s.ChatbotId = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfile) SetFunctionMeta(v *GetScriptResponseBodyDataDraftVersionScriptProfileFunctionMeta) *GetScriptResponseBodyDataDraftVersionScriptProfile {
	s.FunctionMeta = v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfile) SetModel(v string) *GetScriptResponseBodyDataDraftVersionScriptProfile {
	s.Model = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfile) SetNluAccessProfile(v *GetScriptResponseBodyDataDraftVersionScriptProfileNluAccessProfile) *GetScriptResponseBodyDataDraftVersionScriptProfile {
	s.NluAccessProfile = v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfile) SetNluAccessType(v string) *GetScriptResponseBodyDataDraftVersionScriptProfile {
	s.NluAccessType = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfile) SetNluEngine(v string) *GetScriptResponseBodyDataDraftVersionScriptProfile {
	s.NluEngine = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfile) SetOmniModel(v bool) *GetScriptResponseBodyDataDraftVersionScriptProfile {
	s.OmniModel = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfile) Validate() error {
	if s.AgentProfile != nil {
		if err := s.AgentProfile.Validate(); err != nil {
			return err
		}
	}
	if s.FunctionMeta != nil {
		if err := s.FunctionMeta.Validate(); err != nil {
			return err
		}
	}
	if s.NluAccessProfile != nil {
		if err := s.NluAccessProfile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetScriptResponseBodyDataDraftVersionScriptProfileAgentProfile struct {
	PromptsJson *string `json:"PromptsJson,omitempty" xml:"PromptsJson,omitempty"`
	// example:
	//
	// CCC_PROMPTS_DEFAULT
	ScriptProfileTemplateId *string `json:"ScriptProfileTemplateId,omitempty" xml:"ScriptProfileTemplateId,omitempty"`
}

func (s GetScriptResponseBodyDataDraftVersionScriptProfileAgentProfile) String() string {
	return dara.Prettify(s)
}

func (s GetScriptResponseBodyDataDraftVersionScriptProfileAgentProfile) GoString() string {
	return s.String()
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfileAgentProfile) GetPromptsJson() *string {
	return s.PromptsJson
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfileAgentProfile) GetScriptProfileTemplateId() *string {
	return s.ScriptProfileTemplateId
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfileAgentProfile) SetPromptsJson(v string) *GetScriptResponseBodyDataDraftVersionScriptProfileAgentProfile {
	s.PromptsJson = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfileAgentProfile) SetScriptProfileTemplateId(v string) *GetScriptResponseBodyDataDraftVersionScriptProfileAgentProfile {
	s.ScriptProfileTemplateId = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfileAgentProfile) Validate() error {
	return dara.Validate(s)
}

type GetScriptResponseBodyDataDraftVersionScriptProfileFunctionMeta struct {
	// example:
	//
	// 9b752bbb-805a-4d3e-9013-eab5555c3fef
	FunctionId *string `json:"FunctionId,omitempty" xml:"FunctionId,omitempty"`
	// example:
	//
	// my_funciton
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// example:
	//
	// defaultTrigger
	HttpTriggerName *string `json:"HttpTriggerName,omitempty" xml:"HttpTriggerName,omitempty"`
	// example:
	//
	// http://chat-xxxxx-v-yewiundukb.cn-hangzhou-xxx.run
	HttpTriggerUrl *string `json:"HttpTriggerUrl,omitempty" xml:"HttpTriggerUrl,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetScriptResponseBodyDataDraftVersionScriptProfileFunctionMeta) String() string {
	return dara.Prettify(s)
}

func (s GetScriptResponseBodyDataDraftVersionScriptProfileFunctionMeta) GoString() string {
	return s.String()
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfileFunctionMeta) GetFunctionId() *string {
	return s.FunctionId
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfileFunctionMeta) GetFunctionName() *string {
	return s.FunctionName
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfileFunctionMeta) GetHttpTriggerName() *string {
	return s.HttpTriggerName
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfileFunctionMeta) GetHttpTriggerUrl() *string {
	return s.HttpTriggerUrl
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfileFunctionMeta) GetRegionId() *string {
	return s.RegionId
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfileFunctionMeta) SetFunctionId(v string) *GetScriptResponseBodyDataDraftVersionScriptProfileFunctionMeta {
	s.FunctionId = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfileFunctionMeta) SetFunctionName(v string) *GetScriptResponseBodyDataDraftVersionScriptProfileFunctionMeta {
	s.FunctionName = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfileFunctionMeta) SetHttpTriggerName(v string) *GetScriptResponseBodyDataDraftVersionScriptProfileFunctionMeta {
	s.HttpTriggerName = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfileFunctionMeta) SetHttpTriggerUrl(v string) *GetScriptResponseBodyDataDraftVersionScriptProfileFunctionMeta {
	s.HttpTriggerUrl = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfileFunctionMeta) SetRegionId(v string) *GetScriptResponseBodyDataDraftVersionScriptProfileFunctionMeta {
	s.RegionId = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfileFunctionMeta) Validate() error {
	return dara.Validate(s)
}

type GetScriptResponseBodyDataDraftVersionScriptProfileNluAccessProfile struct {
	// example:
	//
	// c2c9baae-9351-4c49-a8cb-6f24a83a8718
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
}

func (s GetScriptResponseBodyDataDraftVersionScriptProfileNluAccessProfile) String() string {
	return dara.Prettify(s)
}

func (s GetScriptResponseBodyDataDraftVersionScriptProfileNluAccessProfile) GoString() string {
	return s.String()
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfileNluAccessProfile) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfileNluAccessProfile) SetAccessProfileId(v string) *GetScriptResponseBodyDataDraftVersionScriptProfileNluAccessProfile {
	s.AccessProfileId = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionScriptProfileNluAccessProfile) Validate() error {
	return dara.Validate(s)
}

type GetScriptResponseBodyDataDraftVersionSynthesizerConfig struct {
	// example:
	//
	// CosyVoice
	Model            *string                                                                 `json:"Model,omitempty" xml:"Model,omitempty"`
	NlsAccessProfile *GetScriptResponseBodyDataDraftVersionSynthesizerConfigNlsAccessProfile `json:"NlsAccessProfile,omitempty" xml:"NlsAccessProfile,omitempty" type:"Struct"`
	// example:
	//
	// MANAGED
	NlsAccessType *string `json:"NlsAccessType,omitempty" xml:"NlsAccessType,omitempty"`
	// example:
	//
	// BAILIAN
	NlsEngine *string `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
	// example:
	//
	// 0
	PitchRate *int32                                                             `json:"PitchRate,omitempty" xml:"PitchRate,omitempty"`
	PronRules []*GetScriptResponseBodyDataDraftVersionSynthesizerConfigPronRules `json:"PronRules,omitempty" xml:"PronRules,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	SpeechRate *int32 `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	// example:
	//
	// longanyang
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
	// example:
	//
	// 50
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s GetScriptResponseBodyDataDraftVersionSynthesizerConfig) String() string {
	return dara.Prettify(s)
}

func (s GetScriptResponseBodyDataDraftVersionSynthesizerConfig) GoString() string {
	return s.String()
}

func (s *GetScriptResponseBodyDataDraftVersionSynthesizerConfig) GetModel() *string {
	return s.Model
}

func (s *GetScriptResponseBodyDataDraftVersionSynthesizerConfig) GetNlsAccessProfile() *GetScriptResponseBodyDataDraftVersionSynthesizerConfigNlsAccessProfile {
	return s.NlsAccessProfile
}

func (s *GetScriptResponseBodyDataDraftVersionSynthesizerConfig) GetNlsAccessType() *string {
	return s.NlsAccessType
}

func (s *GetScriptResponseBodyDataDraftVersionSynthesizerConfig) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *GetScriptResponseBodyDataDraftVersionSynthesizerConfig) GetPitchRate() *int32 {
	return s.PitchRate
}

func (s *GetScriptResponseBodyDataDraftVersionSynthesizerConfig) GetPronRules() []*GetScriptResponseBodyDataDraftVersionSynthesizerConfigPronRules {
	return s.PronRules
}

func (s *GetScriptResponseBodyDataDraftVersionSynthesizerConfig) GetSpeechRate() *int32 {
	return s.SpeechRate
}

func (s *GetScriptResponseBodyDataDraftVersionSynthesizerConfig) GetVoice() *string {
	return s.Voice
}

func (s *GetScriptResponseBodyDataDraftVersionSynthesizerConfig) GetVolume() *int32 {
	return s.Volume
}

func (s *GetScriptResponseBodyDataDraftVersionSynthesizerConfig) SetModel(v string) *GetScriptResponseBodyDataDraftVersionSynthesizerConfig {
	s.Model = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionSynthesizerConfig) SetNlsAccessProfile(v *GetScriptResponseBodyDataDraftVersionSynthesizerConfigNlsAccessProfile) *GetScriptResponseBodyDataDraftVersionSynthesizerConfig {
	s.NlsAccessProfile = v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionSynthesizerConfig) SetNlsAccessType(v string) *GetScriptResponseBodyDataDraftVersionSynthesizerConfig {
	s.NlsAccessType = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionSynthesizerConfig) SetNlsEngine(v string) *GetScriptResponseBodyDataDraftVersionSynthesizerConfig {
	s.NlsEngine = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionSynthesizerConfig) SetPitchRate(v int32) *GetScriptResponseBodyDataDraftVersionSynthesizerConfig {
	s.PitchRate = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionSynthesizerConfig) SetPronRules(v []*GetScriptResponseBodyDataDraftVersionSynthesizerConfigPronRules) *GetScriptResponseBodyDataDraftVersionSynthesizerConfig {
	s.PronRules = v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionSynthesizerConfig) SetSpeechRate(v int32) *GetScriptResponseBodyDataDraftVersionSynthesizerConfig {
	s.SpeechRate = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionSynthesizerConfig) SetVoice(v string) *GetScriptResponseBodyDataDraftVersionSynthesizerConfig {
	s.Voice = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionSynthesizerConfig) SetVolume(v int32) *GetScriptResponseBodyDataDraftVersionSynthesizerConfig {
	s.Volume = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionSynthesizerConfig) Validate() error {
	if s.NlsAccessProfile != nil {
		if err := s.NlsAccessProfile.Validate(); err != nil {
			return err
		}
	}
	if s.PronRules != nil {
		for _, item := range s.PronRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetScriptResponseBodyDataDraftVersionSynthesizerConfigNlsAccessProfile struct {
	// example:
	//
	// 0c4f978a-73bb-4841-bd84-75c0398edd4f
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
}

func (s GetScriptResponseBodyDataDraftVersionSynthesizerConfigNlsAccessProfile) String() string {
	return dara.Prettify(s)
}

func (s GetScriptResponseBodyDataDraftVersionSynthesizerConfigNlsAccessProfile) GoString() string {
	return s.String()
}

func (s *GetScriptResponseBodyDataDraftVersionSynthesizerConfigNlsAccessProfile) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *GetScriptResponseBodyDataDraftVersionSynthesizerConfigNlsAccessProfile) SetAccessProfileId(v string) *GetScriptResponseBodyDataDraftVersionSynthesizerConfigNlsAccessProfile {
	s.AccessProfileId = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionSynthesizerConfigNlsAccessProfile) Validate() error {
	return dara.Validate(s)
}

type GetScriptResponseBodyDataDraftVersionSynthesizerConfigPronRules struct {
	Pattern     *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	Replacement *string `json:"Replacement,omitempty" xml:"Replacement,omitempty"`
}

func (s GetScriptResponseBodyDataDraftVersionSynthesizerConfigPronRules) String() string {
	return dara.Prettify(s)
}

func (s GetScriptResponseBodyDataDraftVersionSynthesizerConfigPronRules) GoString() string {
	return s.String()
}

func (s *GetScriptResponseBodyDataDraftVersionSynthesizerConfigPronRules) GetPattern() *string {
	return s.Pattern
}

func (s *GetScriptResponseBodyDataDraftVersionSynthesizerConfigPronRules) GetReplacement() *string {
	return s.Replacement
}

func (s *GetScriptResponseBodyDataDraftVersionSynthesizerConfigPronRules) SetPattern(v string) *GetScriptResponseBodyDataDraftVersionSynthesizerConfigPronRules {
	s.Pattern = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionSynthesizerConfigPronRules) SetReplacement(v string) *GetScriptResponseBodyDataDraftVersionSynthesizerConfigPronRules {
	s.Replacement = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionSynthesizerConfigPronRules) Validate() error {
	return dara.Validate(s)
}

type GetScriptResponseBodyDataDraftVersionTranscriberConfig struct {
	CorrectionRules []*GetScriptResponseBodyDataDraftVersionTranscriberConfigCorrectionRules `json:"CorrectionRules,omitempty" xml:"CorrectionRules,omitempty" type:"Repeated"`
	// example:
	//
	// 0c4f978a-73bb-4841-bd84-75c0398edd6f
	CustomizationId *string `json:"CustomizationId,omitempty" xml:"CustomizationId,omitempty"`
	// example:
	//
	// 500
	EndSilenceTimeout *int32 `json:"EndSilenceTimeout,omitempty" xml:"EndSilenceTimeout,omitempty"`
	// example:
	//
	// Paraformer
	Model            *string                                                                 `json:"Model,omitempty" xml:"Model,omitempty"`
	NlsAccessProfile *GetScriptResponseBodyDataDraftVersionTranscriberConfigNlsAccessProfile `json:"NlsAccessProfile,omitempty" xml:"NlsAccessProfile,omitempty" type:"Struct"`
	// example:
	//
	// MANAGED
	NlsAccessType *string `json:"NlsAccessType,omitempty" xml:"NlsAccessType,omitempty"`
	// example:
	//
	// BAILIAN
	NlsEngine *string `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
	// example:
	//
	// 0
	SpeechNoiseThreshold *string `json:"SpeechNoiseThreshold,omitempty" xml:"SpeechNoiseThreshold,omitempty"`
	// example:
	//
	// 0c4f978a-73bb-4841-bd84-75c0398edd5f
	VocabularyId *string `json:"VocabularyId,omitempty" xml:"VocabularyId,omitempty"`
}

func (s GetScriptResponseBodyDataDraftVersionTranscriberConfig) String() string {
	return dara.Prettify(s)
}

func (s GetScriptResponseBodyDataDraftVersionTranscriberConfig) GoString() string {
	return s.String()
}

func (s *GetScriptResponseBodyDataDraftVersionTranscriberConfig) GetCorrectionRules() []*GetScriptResponseBodyDataDraftVersionTranscriberConfigCorrectionRules {
	return s.CorrectionRules
}

func (s *GetScriptResponseBodyDataDraftVersionTranscriberConfig) GetCustomizationId() *string {
	return s.CustomizationId
}

func (s *GetScriptResponseBodyDataDraftVersionTranscriberConfig) GetEndSilenceTimeout() *int32 {
	return s.EndSilenceTimeout
}

func (s *GetScriptResponseBodyDataDraftVersionTranscriberConfig) GetModel() *string {
	return s.Model
}

func (s *GetScriptResponseBodyDataDraftVersionTranscriberConfig) GetNlsAccessProfile() *GetScriptResponseBodyDataDraftVersionTranscriberConfigNlsAccessProfile {
	return s.NlsAccessProfile
}

func (s *GetScriptResponseBodyDataDraftVersionTranscriberConfig) GetNlsAccessType() *string {
	return s.NlsAccessType
}

func (s *GetScriptResponseBodyDataDraftVersionTranscriberConfig) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *GetScriptResponseBodyDataDraftVersionTranscriberConfig) GetSpeechNoiseThreshold() *string {
	return s.SpeechNoiseThreshold
}

func (s *GetScriptResponseBodyDataDraftVersionTranscriberConfig) GetVocabularyId() *string {
	return s.VocabularyId
}

func (s *GetScriptResponseBodyDataDraftVersionTranscriberConfig) SetCorrectionRules(v []*GetScriptResponseBodyDataDraftVersionTranscriberConfigCorrectionRules) *GetScriptResponseBodyDataDraftVersionTranscriberConfig {
	s.CorrectionRules = v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionTranscriberConfig) SetCustomizationId(v string) *GetScriptResponseBodyDataDraftVersionTranscriberConfig {
	s.CustomizationId = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionTranscriberConfig) SetEndSilenceTimeout(v int32) *GetScriptResponseBodyDataDraftVersionTranscriberConfig {
	s.EndSilenceTimeout = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionTranscriberConfig) SetModel(v string) *GetScriptResponseBodyDataDraftVersionTranscriberConfig {
	s.Model = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionTranscriberConfig) SetNlsAccessProfile(v *GetScriptResponseBodyDataDraftVersionTranscriberConfigNlsAccessProfile) *GetScriptResponseBodyDataDraftVersionTranscriberConfig {
	s.NlsAccessProfile = v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionTranscriberConfig) SetNlsAccessType(v string) *GetScriptResponseBodyDataDraftVersionTranscriberConfig {
	s.NlsAccessType = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionTranscriberConfig) SetNlsEngine(v string) *GetScriptResponseBodyDataDraftVersionTranscriberConfig {
	s.NlsEngine = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionTranscriberConfig) SetSpeechNoiseThreshold(v string) *GetScriptResponseBodyDataDraftVersionTranscriberConfig {
	s.SpeechNoiseThreshold = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionTranscriberConfig) SetVocabularyId(v string) *GetScriptResponseBodyDataDraftVersionTranscriberConfig {
	s.VocabularyId = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionTranscriberConfig) Validate() error {
	if s.CorrectionRules != nil {
		for _, item := range s.CorrectionRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NlsAccessProfile != nil {
		if err := s.NlsAccessProfile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetScriptResponseBodyDataDraftVersionTranscriberConfigCorrectionRules struct {
	Pattern     *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	Replacement *string `json:"Replacement,omitempty" xml:"Replacement,omitempty"`
}

func (s GetScriptResponseBodyDataDraftVersionTranscriberConfigCorrectionRules) String() string {
	return dara.Prettify(s)
}

func (s GetScriptResponseBodyDataDraftVersionTranscriberConfigCorrectionRules) GoString() string {
	return s.String()
}

func (s *GetScriptResponseBodyDataDraftVersionTranscriberConfigCorrectionRules) GetPattern() *string {
	return s.Pattern
}

func (s *GetScriptResponseBodyDataDraftVersionTranscriberConfigCorrectionRules) GetReplacement() *string {
	return s.Replacement
}

func (s *GetScriptResponseBodyDataDraftVersionTranscriberConfigCorrectionRules) SetPattern(v string) *GetScriptResponseBodyDataDraftVersionTranscriberConfigCorrectionRules {
	s.Pattern = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionTranscriberConfigCorrectionRules) SetReplacement(v string) *GetScriptResponseBodyDataDraftVersionTranscriberConfigCorrectionRules {
	s.Replacement = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionTranscriberConfigCorrectionRules) Validate() error {
	return dara.Validate(s)
}

type GetScriptResponseBodyDataDraftVersionTranscriberConfigNlsAccessProfile struct {
	// example:
	//
	// 0c4f978a-73bb-4841-bd84-75c0398edd4f
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
}

func (s GetScriptResponseBodyDataDraftVersionTranscriberConfigNlsAccessProfile) String() string {
	return dara.Prettify(s)
}

func (s GetScriptResponseBodyDataDraftVersionTranscriberConfigNlsAccessProfile) GoString() string {
	return s.String()
}

func (s *GetScriptResponseBodyDataDraftVersionTranscriberConfigNlsAccessProfile) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *GetScriptResponseBodyDataDraftVersionTranscriberConfigNlsAccessProfile) SetAccessProfileId(v string) *GetScriptResponseBodyDataDraftVersionTranscriberConfigNlsAccessProfile {
	s.AccessProfileId = &v
	return s
}

func (s *GetScriptResponseBodyDataDraftVersionTranscriberConfigNlsAccessProfile) Validate() error {
	return dara.Validate(s)
}

type GetScriptResponseBodyDataPublishedVersion struct {
	InteractionConfig *GetScriptResponseBodyDataPublishedVersionInteractionConfig `json:"InteractionConfig,omitempty" xml:"InteractionConfig,omitempty" type:"Struct"`
	LabelConfig       []*GetScriptResponseBodyDataPublishedVersionLabelConfig     `json:"LabelConfig,omitempty" xml:"LabelConfig,omitempty" type:"Repeated"`
	ScriptProfile     *GetScriptResponseBodyDataPublishedVersionScriptProfile     `json:"ScriptProfile,omitempty" xml:"ScriptProfile,omitempty" type:"Struct"`
	SynthesizerConfig *GetScriptResponseBodyDataPublishedVersionSynthesizerConfig `json:"SynthesizerConfig,omitempty" xml:"SynthesizerConfig,omitempty" type:"Struct"`
	TranscriberConfig *GetScriptResponseBodyDataPublishedVersionTranscriberConfig `json:"TranscriberConfig,omitempty" xml:"TranscriberConfig,omitempty" type:"Struct"`
	// example:
	//
	// 8b77ff09-6a90-4784-8560-fdc2b860dc68
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetScriptResponseBodyDataPublishedVersion) String() string {
	return dara.Prettify(s)
}

func (s GetScriptResponseBodyDataPublishedVersion) GoString() string {
	return s.String()
}

func (s *GetScriptResponseBodyDataPublishedVersion) GetInteractionConfig() *GetScriptResponseBodyDataPublishedVersionInteractionConfig {
	return s.InteractionConfig
}

func (s *GetScriptResponseBodyDataPublishedVersion) GetLabelConfig() []*GetScriptResponseBodyDataPublishedVersionLabelConfig {
	return s.LabelConfig
}

func (s *GetScriptResponseBodyDataPublishedVersion) GetScriptProfile() *GetScriptResponseBodyDataPublishedVersionScriptProfile {
	return s.ScriptProfile
}

func (s *GetScriptResponseBodyDataPublishedVersion) GetSynthesizerConfig() *GetScriptResponseBodyDataPublishedVersionSynthesizerConfig {
	return s.SynthesizerConfig
}

func (s *GetScriptResponseBodyDataPublishedVersion) GetTranscriberConfig() *GetScriptResponseBodyDataPublishedVersionTranscriberConfig {
	return s.TranscriberConfig
}

func (s *GetScriptResponseBodyDataPublishedVersion) GetVersionId() *string {
	return s.VersionId
}

func (s *GetScriptResponseBodyDataPublishedVersion) SetInteractionConfig(v *GetScriptResponseBodyDataPublishedVersionInteractionConfig) *GetScriptResponseBodyDataPublishedVersion {
	s.InteractionConfig = v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersion) SetLabelConfig(v []*GetScriptResponseBodyDataPublishedVersionLabelConfig) *GetScriptResponseBodyDataPublishedVersion {
	s.LabelConfig = v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersion) SetScriptProfile(v *GetScriptResponseBodyDataPublishedVersionScriptProfile) *GetScriptResponseBodyDataPublishedVersion {
	s.ScriptProfile = v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersion) SetSynthesizerConfig(v *GetScriptResponseBodyDataPublishedVersionSynthesizerConfig) *GetScriptResponseBodyDataPublishedVersion {
	s.SynthesizerConfig = v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersion) SetTranscriberConfig(v *GetScriptResponseBodyDataPublishedVersionTranscriberConfig) *GetScriptResponseBodyDataPublishedVersion {
	s.TranscriberConfig = v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersion) SetVersionId(v string) *GetScriptResponseBodyDataPublishedVersion {
	s.VersionId = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersion) Validate() error {
	if s.InteractionConfig != nil {
		if err := s.InteractionConfig.Validate(); err != nil {
			return err
		}
	}
	if s.LabelConfig != nil {
		for _, item := range s.LabelConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ScriptProfile != nil {
		if err := s.ScriptProfile.Validate(); err != nil {
			return err
		}
	}
	if s.SynthesizerConfig != nil {
		if err := s.SynthesizerConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TranscriberConfig != nil {
		if err := s.TranscriberConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetScriptResponseBodyDataPublishedVersionInteractionConfig struct {
	// example:
	//
	// office-ambience
	BackgroundMusicId     *string                                                                          `json:"BackgroundMusicId,omitempty" xml:"BackgroundMusicId,omitempty"`
	EndConversationConfig *GetScriptResponseBodyDataPublishedVersionInteractionConfigEndConversationConfig `json:"EndConversationConfig,omitempty" xml:"EndConversationConfig,omitempty" type:"Struct"`
	// example:
	//
	// 2000
	InitialGreetingDelayMilliseconds *int32                                                                            `json:"InitialGreetingDelayMilliseconds,omitempty" xml:"InitialGreetingDelayMilliseconds,omitempty"`
	SilenceDetectionConfig           *GetScriptResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig `json:"SilenceDetectionConfig,omitempty" xml:"SilenceDetectionConfig,omitempty" type:"Struct"`
}

func (s GetScriptResponseBodyDataPublishedVersionInteractionConfig) String() string {
	return dara.Prettify(s)
}

func (s GetScriptResponseBodyDataPublishedVersionInteractionConfig) GoString() string {
	return s.String()
}

func (s *GetScriptResponseBodyDataPublishedVersionInteractionConfig) GetBackgroundMusicId() *string {
	return s.BackgroundMusicId
}

func (s *GetScriptResponseBodyDataPublishedVersionInteractionConfig) GetEndConversationConfig() *GetScriptResponseBodyDataPublishedVersionInteractionConfigEndConversationConfig {
	return s.EndConversationConfig
}

func (s *GetScriptResponseBodyDataPublishedVersionInteractionConfig) GetInitialGreetingDelayMilliseconds() *int32 {
	return s.InitialGreetingDelayMilliseconds
}

func (s *GetScriptResponseBodyDataPublishedVersionInteractionConfig) GetSilenceDetectionConfig() *GetScriptResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig {
	return s.SilenceDetectionConfig
}

func (s *GetScriptResponseBodyDataPublishedVersionInteractionConfig) SetBackgroundMusicId(v string) *GetScriptResponseBodyDataPublishedVersionInteractionConfig {
	s.BackgroundMusicId = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionInteractionConfig) SetEndConversationConfig(v *GetScriptResponseBodyDataPublishedVersionInteractionConfigEndConversationConfig) *GetScriptResponseBodyDataPublishedVersionInteractionConfig {
	s.EndConversationConfig = v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionInteractionConfig) SetInitialGreetingDelayMilliseconds(v int32) *GetScriptResponseBodyDataPublishedVersionInteractionConfig {
	s.InitialGreetingDelayMilliseconds = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionInteractionConfig) SetSilenceDetectionConfig(v *GetScriptResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig) *GetScriptResponseBodyDataPublishedVersionInteractionConfig {
	s.SilenceDetectionConfig = v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionInteractionConfig) Validate() error {
	if s.EndConversationConfig != nil {
		if err := s.EndConversationConfig.Validate(); err != nil {
			return err
		}
	}
	if s.SilenceDetectionConfig != nil {
		if err := s.SilenceDetectionConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetScriptResponseBodyDataPublishedVersionInteractionConfigEndConversationConfig struct {
	// example:
	//
	// 1
	Delay    *int32                                                                                     `json:"Delay,omitempty" xml:"Delay,omitempty"`
	Triggers []*GetScriptResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers `json:"Triggers,omitempty" xml:"Triggers,omitempty" type:"Repeated"`
}

func (s GetScriptResponseBodyDataPublishedVersionInteractionConfigEndConversationConfig) String() string {
	return dara.Prettify(s)
}

func (s GetScriptResponseBodyDataPublishedVersionInteractionConfigEndConversationConfig) GoString() string {
	return s.String()
}

func (s *GetScriptResponseBodyDataPublishedVersionInteractionConfigEndConversationConfig) GetDelay() *int32 {
	return s.Delay
}

func (s *GetScriptResponseBodyDataPublishedVersionInteractionConfigEndConversationConfig) GetTriggers() []*GetScriptResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers {
	return s.Triggers
}

func (s *GetScriptResponseBodyDataPublishedVersionInteractionConfigEndConversationConfig) SetDelay(v int32) *GetScriptResponseBodyDataPublishedVersionInteractionConfigEndConversationConfig {
	s.Delay = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionInteractionConfigEndConversationConfig) SetTriggers(v []*GetScriptResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers) *GetScriptResponseBodyDataPublishedVersionInteractionConfigEndConversationConfig {
	s.Triggers = v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionInteractionConfigEndConversationConfig) Validate() error {
	if s.Triggers != nil {
		for _, item := range s.Triggers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetScriptResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers struct {
	ClosingStatement *string   `json:"ClosingStatement,omitempty" xml:"ClosingStatement,omitempty"`
	KeyWords         []*string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty" type:"Repeated"`
	// example:
	//
	// TurnLimit
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// example:
	//
	// 20
	TurnLimit *int32 `json:"TurnLimit,omitempty" xml:"TurnLimit,omitempty"`
}

func (s GetScriptResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers) String() string {
	return dara.Prettify(s)
}

func (s GetScriptResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers) GoString() string {
	return s.String()
}

func (s *GetScriptResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers) GetClosingStatement() *string {
	return s.ClosingStatement
}

func (s *GetScriptResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers) GetKeyWords() []*string {
	return s.KeyWords
}

func (s *GetScriptResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers) GetTriggerType() *string {
	return s.TriggerType
}

func (s *GetScriptResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers) GetTurnLimit() *int32 {
	return s.TurnLimit
}

func (s *GetScriptResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers) SetClosingStatement(v string) *GetScriptResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers {
	s.ClosingStatement = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers) SetKeyWords(v []*string) *GetScriptResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers {
	s.KeyWords = v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers) SetTriggerType(v string) *GetScriptResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers {
	s.TriggerType = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers) SetTurnLimit(v int32) *GetScriptResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers {
	s.TurnLimit = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers) Validate() error {
	return dara.Validate(s)
}

type GetScriptResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig struct {
	// example:
	//
	// 3
	MaxRepeats *int32 `json:"MaxRepeats,omitempty" xml:"MaxRepeats,omitempty"`
	// example:
	//
	// 5000
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s GetScriptResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig) String() string {
	return dara.Prettify(s)
}

func (s GetScriptResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig) GoString() string {
	return s.String()
}

func (s *GetScriptResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig) GetMaxRepeats() *int32 {
	return s.MaxRepeats
}

func (s *GetScriptResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig) GetTimeout() *int32 {
	return s.Timeout
}

func (s *GetScriptResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig) SetMaxRepeats(v int32) *GetScriptResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig {
	s.MaxRepeats = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig) SetTimeout(v int32) *GetScriptResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig {
	s.Timeout = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig) Validate() error {
	return dara.Validate(s)
}

type GetScriptResponseBodyDataPublishedVersionLabelConfig struct {
	CandidateValues []*string `json:"CandidateValues,omitempty" xml:"CandidateValues,omitempty" type:"Repeated"`
	Description     *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Name            *string   `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetScriptResponseBodyDataPublishedVersionLabelConfig) String() string {
	return dara.Prettify(s)
}

func (s GetScriptResponseBodyDataPublishedVersionLabelConfig) GoString() string {
	return s.String()
}

func (s *GetScriptResponseBodyDataPublishedVersionLabelConfig) GetCandidateValues() []*string {
	return s.CandidateValues
}

func (s *GetScriptResponseBodyDataPublishedVersionLabelConfig) GetDescription() *string {
	return s.Description
}

func (s *GetScriptResponseBodyDataPublishedVersionLabelConfig) GetName() *string {
	return s.Name
}

func (s *GetScriptResponseBodyDataPublishedVersionLabelConfig) SetCandidateValues(v []*string) *GetScriptResponseBodyDataPublishedVersionLabelConfig {
	s.CandidateValues = v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionLabelConfig) SetDescription(v string) *GetScriptResponseBodyDataPublishedVersionLabelConfig {
	s.Description = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionLabelConfig) SetName(v string) *GetScriptResponseBodyDataPublishedVersionLabelConfig {
	s.Name = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionLabelConfig) Validate() error {
	return dara.Validate(s)
}

type GetScriptResponseBodyDataPublishedVersionScriptProfile struct {
	// example:
	//
	// 1309723684579735_p_beebot_public
	AgentKey     *string                                                             `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	AgentProfile *GetScriptResponseBodyDataPublishedVersionScriptProfileAgentProfile `json:"AgentProfile,omitempty" xml:"AgentProfile,omitempty" type:"Struct"`
	// example:
	//
	// chatbot-cn-MQuyjjb666
	ChatbotId    *string                                                             `json:"ChatbotId,omitempty" xml:"ChatbotId,omitempty"`
	FunctionMeta *GetScriptResponseBodyDataPublishedVersionScriptProfileFunctionMeta `json:"FunctionMeta,omitempty" xml:"FunctionMeta,omitempty" type:"Struct"`
	// example:
	//
	// qwen-plus
	Model            *string                                                                 `json:"Model,omitempty" xml:"Model,omitempty"`
	NluAccessProfile *GetScriptResponseBodyDataPublishedVersionScriptProfileNluAccessProfile `json:"NluAccessProfile,omitempty" xml:"NluAccessProfile,omitempty" type:"Struct"`
	// example:
	//
	// MANAGED
	NluAccessType *string `json:"NluAccessType,omitempty" xml:"NluAccessType,omitempty"`
	// example:
	//
	// BEEBOT
	NluEngine *string `json:"NluEngine,omitempty" xml:"NluEngine,omitempty"`
	// example:
	//
	// true
	OmniModel *bool `json:"OmniModel,omitempty" xml:"OmniModel,omitempty"`
}

func (s GetScriptResponseBodyDataPublishedVersionScriptProfile) String() string {
	return dara.Prettify(s)
}

func (s GetScriptResponseBodyDataPublishedVersionScriptProfile) GoString() string {
	return s.String()
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfile) GetAgentKey() *string {
	return s.AgentKey
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfile) GetAgentProfile() *GetScriptResponseBodyDataPublishedVersionScriptProfileAgentProfile {
	return s.AgentProfile
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfile) GetChatbotId() *string {
	return s.ChatbotId
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfile) GetFunctionMeta() *GetScriptResponseBodyDataPublishedVersionScriptProfileFunctionMeta {
	return s.FunctionMeta
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfile) GetModel() *string {
	return s.Model
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfile) GetNluAccessProfile() *GetScriptResponseBodyDataPublishedVersionScriptProfileNluAccessProfile {
	return s.NluAccessProfile
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfile) GetNluAccessType() *string {
	return s.NluAccessType
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfile) GetNluEngine() *string {
	return s.NluEngine
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfile) GetOmniModel() *bool {
	return s.OmniModel
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfile) SetAgentKey(v string) *GetScriptResponseBodyDataPublishedVersionScriptProfile {
	s.AgentKey = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfile) SetAgentProfile(v *GetScriptResponseBodyDataPublishedVersionScriptProfileAgentProfile) *GetScriptResponseBodyDataPublishedVersionScriptProfile {
	s.AgentProfile = v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfile) SetChatbotId(v string) *GetScriptResponseBodyDataPublishedVersionScriptProfile {
	s.ChatbotId = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfile) SetFunctionMeta(v *GetScriptResponseBodyDataPublishedVersionScriptProfileFunctionMeta) *GetScriptResponseBodyDataPublishedVersionScriptProfile {
	s.FunctionMeta = v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfile) SetModel(v string) *GetScriptResponseBodyDataPublishedVersionScriptProfile {
	s.Model = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfile) SetNluAccessProfile(v *GetScriptResponseBodyDataPublishedVersionScriptProfileNluAccessProfile) *GetScriptResponseBodyDataPublishedVersionScriptProfile {
	s.NluAccessProfile = v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfile) SetNluAccessType(v string) *GetScriptResponseBodyDataPublishedVersionScriptProfile {
	s.NluAccessType = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfile) SetNluEngine(v string) *GetScriptResponseBodyDataPublishedVersionScriptProfile {
	s.NluEngine = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfile) SetOmniModel(v bool) *GetScriptResponseBodyDataPublishedVersionScriptProfile {
	s.OmniModel = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfile) Validate() error {
	if s.AgentProfile != nil {
		if err := s.AgentProfile.Validate(); err != nil {
			return err
		}
	}
	if s.FunctionMeta != nil {
		if err := s.FunctionMeta.Validate(); err != nil {
			return err
		}
	}
	if s.NluAccessProfile != nil {
		if err := s.NluAccessProfile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetScriptResponseBodyDataPublishedVersionScriptProfileAgentProfile struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// test agent
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PromptsJson *string `json:"PromptsJson,omitempty" xml:"PromptsJson,omitempty"`
	// example:
	//
	// CCC_PROMPTS_DEFAULT
	ScriptProfileTemplateId *string `json:"ScriptProfileTemplateId,omitempty" xml:"ScriptProfileTemplateId,omitempty"`
}

func (s GetScriptResponseBodyDataPublishedVersionScriptProfileAgentProfile) String() string {
	return dara.Prettify(s)
}

func (s GetScriptResponseBodyDataPublishedVersionScriptProfileAgentProfile) GoString() string {
	return s.String()
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfileAgentProfile) GetDescription() *string {
	return s.Description
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfileAgentProfile) GetName() *string {
	return s.Name
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfileAgentProfile) GetPromptsJson() *string {
	return s.PromptsJson
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfileAgentProfile) GetScriptProfileTemplateId() *string {
	return s.ScriptProfileTemplateId
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfileAgentProfile) SetDescription(v string) *GetScriptResponseBodyDataPublishedVersionScriptProfileAgentProfile {
	s.Description = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfileAgentProfile) SetName(v string) *GetScriptResponseBodyDataPublishedVersionScriptProfileAgentProfile {
	s.Name = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfileAgentProfile) SetPromptsJson(v string) *GetScriptResponseBodyDataPublishedVersionScriptProfileAgentProfile {
	s.PromptsJson = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfileAgentProfile) SetScriptProfileTemplateId(v string) *GetScriptResponseBodyDataPublishedVersionScriptProfileAgentProfile {
	s.ScriptProfileTemplateId = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfileAgentProfile) Validate() error {
	return dara.Validate(s)
}

type GetScriptResponseBodyDataPublishedVersionScriptProfileFunctionMeta struct {
	// example:
	//
	// 9b752bbb-805a-4d3e-9013-eab5555c3fef
	FunctionId *string `json:"FunctionId,omitempty" xml:"FunctionId,omitempty"`
	// example:
	//
	// my_function
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// example:
	//
	// defaultTrigger
	HttpTriggerName *string `json:"HttpTriggerName,omitempty" xml:"HttpTriggerName,omitempty"`
	// example:
	//
	// http://chat-xxxxx-v-yewiundukb.cn-hangzhou-xxx.run
	HttpTriggerUrl *string `json:"HttpTriggerUrl,omitempty" xml:"HttpTriggerUrl,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetScriptResponseBodyDataPublishedVersionScriptProfileFunctionMeta) String() string {
	return dara.Prettify(s)
}

func (s GetScriptResponseBodyDataPublishedVersionScriptProfileFunctionMeta) GoString() string {
	return s.String()
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfileFunctionMeta) GetFunctionId() *string {
	return s.FunctionId
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfileFunctionMeta) GetFunctionName() *string {
	return s.FunctionName
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfileFunctionMeta) GetHttpTriggerName() *string {
	return s.HttpTriggerName
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfileFunctionMeta) GetHttpTriggerUrl() *string {
	return s.HttpTriggerUrl
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfileFunctionMeta) GetRegionId() *string {
	return s.RegionId
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfileFunctionMeta) SetFunctionId(v string) *GetScriptResponseBodyDataPublishedVersionScriptProfileFunctionMeta {
	s.FunctionId = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfileFunctionMeta) SetFunctionName(v string) *GetScriptResponseBodyDataPublishedVersionScriptProfileFunctionMeta {
	s.FunctionName = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfileFunctionMeta) SetHttpTriggerName(v string) *GetScriptResponseBodyDataPublishedVersionScriptProfileFunctionMeta {
	s.HttpTriggerName = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfileFunctionMeta) SetHttpTriggerUrl(v string) *GetScriptResponseBodyDataPublishedVersionScriptProfileFunctionMeta {
	s.HttpTriggerUrl = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfileFunctionMeta) SetRegionId(v string) *GetScriptResponseBodyDataPublishedVersionScriptProfileFunctionMeta {
	s.RegionId = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfileFunctionMeta) Validate() error {
	return dara.Validate(s)
}

type GetScriptResponseBodyDataPublishedVersionScriptProfileNluAccessProfile struct {
	// example:
	//
	// c2c9baae-9351-4c49-a8cb-6f24a83a8718
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
}

func (s GetScriptResponseBodyDataPublishedVersionScriptProfileNluAccessProfile) String() string {
	return dara.Prettify(s)
}

func (s GetScriptResponseBodyDataPublishedVersionScriptProfileNluAccessProfile) GoString() string {
	return s.String()
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfileNluAccessProfile) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfileNluAccessProfile) SetAccessProfileId(v string) *GetScriptResponseBodyDataPublishedVersionScriptProfileNluAccessProfile {
	s.AccessProfileId = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionScriptProfileNluAccessProfile) Validate() error {
	return dara.Validate(s)
}

type GetScriptResponseBodyDataPublishedVersionSynthesizerConfig struct {
	// example:
	//
	// CosyVoice
	Model            *string                                                                     `json:"Model,omitempty" xml:"Model,omitempty"`
	NlsAccessProfile *GetScriptResponseBodyDataPublishedVersionSynthesizerConfigNlsAccessProfile `json:"NlsAccessProfile,omitempty" xml:"NlsAccessProfile,omitempty" type:"Struct"`
	// example:
	//
	// MANAGED
	NlsAccessType *string `json:"NlsAccessType,omitempty" xml:"NlsAccessType,omitempty"`
	// example:
	//
	// BAILIAN
	NlsEngine *string `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
	// example:
	//
	// 0
	PitchRate *int32                                                                 `json:"PitchRate,omitempty" xml:"PitchRate,omitempty"`
	PronRules []*GetScriptResponseBodyDataPublishedVersionSynthesizerConfigPronRules `json:"PronRules,omitempty" xml:"PronRules,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	SpeechRate *int32 `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	// example:
	//
	// longanyang
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
	// example:
	//
	// 50
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s GetScriptResponseBodyDataPublishedVersionSynthesizerConfig) String() string {
	return dara.Prettify(s)
}

func (s GetScriptResponseBodyDataPublishedVersionSynthesizerConfig) GoString() string {
	return s.String()
}

func (s *GetScriptResponseBodyDataPublishedVersionSynthesizerConfig) GetModel() *string {
	return s.Model
}

func (s *GetScriptResponseBodyDataPublishedVersionSynthesizerConfig) GetNlsAccessProfile() *GetScriptResponseBodyDataPublishedVersionSynthesizerConfigNlsAccessProfile {
	return s.NlsAccessProfile
}

func (s *GetScriptResponseBodyDataPublishedVersionSynthesizerConfig) GetNlsAccessType() *string {
	return s.NlsAccessType
}

func (s *GetScriptResponseBodyDataPublishedVersionSynthesizerConfig) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *GetScriptResponseBodyDataPublishedVersionSynthesizerConfig) GetPitchRate() *int32 {
	return s.PitchRate
}

func (s *GetScriptResponseBodyDataPublishedVersionSynthesizerConfig) GetPronRules() []*GetScriptResponseBodyDataPublishedVersionSynthesizerConfigPronRules {
	return s.PronRules
}

func (s *GetScriptResponseBodyDataPublishedVersionSynthesizerConfig) GetSpeechRate() *int32 {
	return s.SpeechRate
}

func (s *GetScriptResponseBodyDataPublishedVersionSynthesizerConfig) GetVoice() *string {
	return s.Voice
}

func (s *GetScriptResponseBodyDataPublishedVersionSynthesizerConfig) GetVolume() *int32 {
	return s.Volume
}

func (s *GetScriptResponseBodyDataPublishedVersionSynthesizerConfig) SetModel(v string) *GetScriptResponseBodyDataPublishedVersionSynthesizerConfig {
	s.Model = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionSynthesizerConfig) SetNlsAccessProfile(v *GetScriptResponseBodyDataPublishedVersionSynthesizerConfigNlsAccessProfile) *GetScriptResponseBodyDataPublishedVersionSynthesizerConfig {
	s.NlsAccessProfile = v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionSynthesizerConfig) SetNlsAccessType(v string) *GetScriptResponseBodyDataPublishedVersionSynthesizerConfig {
	s.NlsAccessType = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionSynthesizerConfig) SetNlsEngine(v string) *GetScriptResponseBodyDataPublishedVersionSynthesizerConfig {
	s.NlsEngine = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionSynthesizerConfig) SetPitchRate(v int32) *GetScriptResponseBodyDataPublishedVersionSynthesizerConfig {
	s.PitchRate = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionSynthesizerConfig) SetPronRules(v []*GetScriptResponseBodyDataPublishedVersionSynthesizerConfigPronRules) *GetScriptResponseBodyDataPublishedVersionSynthesizerConfig {
	s.PronRules = v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionSynthesizerConfig) SetSpeechRate(v int32) *GetScriptResponseBodyDataPublishedVersionSynthesizerConfig {
	s.SpeechRate = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionSynthesizerConfig) SetVoice(v string) *GetScriptResponseBodyDataPublishedVersionSynthesizerConfig {
	s.Voice = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionSynthesizerConfig) SetVolume(v int32) *GetScriptResponseBodyDataPublishedVersionSynthesizerConfig {
	s.Volume = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionSynthesizerConfig) Validate() error {
	if s.NlsAccessProfile != nil {
		if err := s.NlsAccessProfile.Validate(); err != nil {
			return err
		}
	}
	if s.PronRules != nil {
		for _, item := range s.PronRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetScriptResponseBodyDataPublishedVersionSynthesizerConfigNlsAccessProfile struct {
	// example:
	//
	// c2c9baae-9351-4c49-a8cb-6f24a83a8718
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
}

func (s GetScriptResponseBodyDataPublishedVersionSynthesizerConfigNlsAccessProfile) String() string {
	return dara.Prettify(s)
}

func (s GetScriptResponseBodyDataPublishedVersionSynthesizerConfigNlsAccessProfile) GoString() string {
	return s.String()
}

func (s *GetScriptResponseBodyDataPublishedVersionSynthesizerConfigNlsAccessProfile) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *GetScriptResponseBodyDataPublishedVersionSynthesizerConfigNlsAccessProfile) SetAccessProfileId(v string) *GetScriptResponseBodyDataPublishedVersionSynthesizerConfigNlsAccessProfile {
	s.AccessProfileId = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionSynthesizerConfigNlsAccessProfile) Validate() error {
	return dara.Validate(s)
}

type GetScriptResponseBodyDataPublishedVersionSynthesizerConfigPronRules struct {
	Pattern     *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	Replacement *string `json:"Replacement,omitempty" xml:"Replacement,omitempty"`
}

func (s GetScriptResponseBodyDataPublishedVersionSynthesizerConfigPronRules) String() string {
	return dara.Prettify(s)
}

func (s GetScriptResponseBodyDataPublishedVersionSynthesizerConfigPronRules) GoString() string {
	return s.String()
}

func (s *GetScriptResponseBodyDataPublishedVersionSynthesizerConfigPronRules) GetPattern() *string {
	return s.Pattern
}

func (s *GetScriptResponseBodyDataPublishedVersionSynthesizerConfigPronRules) GetReplacement() *string {
	return s.Replacement
}

func (s *GetScriptResponseBodyDataPublishedVersionSynthesizerConfigPronRules) SetPattern(v string) *GetScriptResponseBodyDataPublishedVersionSynthesizerConfigPronRules {
	s.Pattern = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionSynthesizerConfigPronRules) SetReplacement(v string) *GetScriptResponseBodyDataPublishedVersionSynthesizerConfigPronRules {
	s.Replacement = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionSynthesizerConfigPronRules) Validate() error {
	return dara.Validate(s)
}

type GetScriptResponseBodyDataPublishedVersionTranscriberConfig struct {
	CorrectionRules []*GetScriptResponseBodyDataPublishedVersionTranscriberConfigCorrectionRules `json:"CorrectionRules,omitempty" xml:"CorrectionRules,omitempty" type:"Repeated"`
	// example:
	//
	// cd97223f-42f2-4cd9-95af-e734e2fe1fe4
	CustomizationId *string `json:"CustomizationId,omitempty" xml:"CustomizationId,omitempty"`
	// example:
	//
	// 700
	EndSilenceTimeout *int32 `json:"EndSilenceTimeout,omitempty" xml:"EndSilenceTimeout,omitempty"`
	// example:
	//
	// Paraformer
	Model            *string                                                                     `json:"Model,omitempty" xml:"Model,omitempty"`
	NlsAccessProfile *GetScriptResponseBodyDataPublishedVersionTranscriberConfigNlsAccessProfile `json:"NlsAccessProfile,omitempty" xml:"NlsAccessProfile,omitempty" type:"Struct"`
	// example:
	//
	// MANAGED
	NlsAccessType *string `json:"NlsAccessType,omitempty" xml:"NlsAccessType,omitempty"`
	// example:
	//
	// BAILIAN
	NlsEngine *string `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
	// example:
	//
	// 0
	SpeechNoiseThreshold *string `json:"SpeechNoiseThreshold,omitempty" xml:"SpeechNoiseThreshold,omitempty"`
	// example:
	//
	// cd97223f-42f2-4cd9-95af-e734e2fe1fe3
	VocabularyId *string `json:"VocabularyId,omitempty" xml:"VocabularyId,omitempty"`
}

func (s GetScriptResponseBodyDataPublishedVersionTranscriberConfig) String() string {
	return dara.Prettify(s)
}

func (s GetScriptResponseBodyDataPublishedVersionTranscriberConfig) GoString() string {
	return s.String()
}

func (s *GetScriptResponseBodyDataPublishedVersionTranscriberConfig) GetCorrectionRules() []*GetScriptResponseBodyDataPublishedVersionTranscriberConfigCorrectionRules {
	return s.CorrectionRules
}

func (s *GetScriptResponseBodyDataPublishedVersionTranscriberConfig) GetCustomizationId() *string {
	return s.CustomizationId
}

func (s *GetScriptResponseBodyDataPublishedVersionTranscriberConfig) GetEndSilenceTimeout() *int32 {
	return s.EndSilenceTimeout
}

func (s *GetScriptResponseBodyDataPublishedVersionTranscriberConfig) GetModel() *string {
	return s.Model
}

func (s *GetScriptResponseBodyDataPublishedVersionTranscriberConfig) GetNlsAccessProfile() *GetScriptResponseBodyDataPublishedVersionTranscriberConfigNlsAccessProfile {
	return s.NlsAccessProfile
}

func (s *GetScriptResponseBodyDataPublishedVersionTranscriberConfig) GetNlsAccessType() *string {
	return s.NlsAccessType
}

func (s *GetScriptResponseBodyDataPublishedVersionTranscriberConfig) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *GetScriptResponseBodyDataPublishedVersionTranscriberConfig) GetSpeechNoiseThreshold() *string {
	return s.SpeechNoiseThreshold
}

func (s *GetScriptResponseBodyDataPublishedVersionTranscriberConfig) GetVocabularyId() *string {
	return s.VocabularyId
}

func (s *GetScriptResponseBodyDataPublishedVersionTranscriberConfig) SetCorrectionRules(v []*GetScriptResponseBodyDataPublishedVersionTranscriberConfigCorrectionRules) *GetScriptResponseBodyDataPublishedVersionTranscriberConfig {
	s.CorrectionRules = v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionTranscriberConfig) SetCustomizationId(v string) *GetScriptResponseBodyDataPublishedVersionTranscriberConfig {
	s.CustomizationId = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionTranscriberConfig) SetEndSilenceTimeout(v int32) *GetScriptResponseBodyDataPublishedVersionTranscriberConfig {
	s.EndSilenceTimeout = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionTranscriberConfig) SetModel(v string) *GetScriptResponseBodyDataPublishedVersionTranscriberConfig {
	s.Model = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionTranscriberConfig) SetNlsAccessProfile(v *GetScriptResponseBodyDataPublishedVersionTranscriberConfigNlsAccessProfile) *GetScriptResponseBodyDataPublishedVersionTranscriberConfig {
	s.NlsAccessProfile = v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionTranscriberConfig) SetNlsAccessType(v string) *GetScriptResponseBodyDataPublishedVersionTranscriberConfig {
	s.NlsAccessType = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionTranscriberConfig) SetNlsEngine(v string) *GetScriptResponseBodyDataPublishedVersionTranscriberConfig {
	s.NlsEngine = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionTranscriberConfig) SetSpeechNoiseThreshold(v string) *GetScriptResponseBodyDataPublishedVersionTranscriberConfig {
	s.SpeechNoiseThreshold = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionTranscriberConfig) SetVocabularyId(v string) *GetScriptResponseBodyDataPublishedVersionTranscriberConfig {
	s.VocabularyId = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionTranscriberConfig) Validate() error {
	if s.CorrectionRules != nil {
		for _, item := range s.CorrectionRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NlsAccessProfile != nil {
		if err := s.NlsAccessProfile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetScriptResponseBodyDataPublishedVersionTranscriberConfigCorrectionRules struct {
	Pattern     *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	Replacement *string `json:"Replacement,omitempty" xml:"Replacement,omitempty"`
}

func (s GetScriptResponseBodyDataPublishedVersionTranscriberConfigCorrectionRules) String() string {
	return dara.Prettify(s)
}

func (s GetScriptResponseBodyDataPublishedVersionTranscriberConfigCorrectionRules) GoString() string {
	return s.String()
}

func (s *GetScriptResponseBodyDataPublishedVersionTranscriberConfigCorrectionRules) GetPattern() *string {
	return s.Pattern
}

func (s *GetScriptResponseBodyDataPublishedVersionTranscriberConfigCorrectionRules) GetReplacement() *string {
	return s.Replacement
}

func (s *GetScriptResponseBodyDataPublishedVersionTranscriberConfigCorrectionRules) SetPattern(v string) *GetScriptResponseBodyDataPublishedVersionTranscriberConfigCorrectionRules {
	s.Pattern = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionTranscriberConfigCorrectionRules) SetReplacement(v string) *GetScriptResponseBodyDataPublishedVersionTranscriberConfigCorrectionRules {
	s.Replacement = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionTranscriberConfigCorrectionRules) Validate() error {
	return dara.Validate(s)
}

type GetScriptResponseBodyDataPublishedVersionTranscriberConfigNlsAccessProfile struct {
	// example:
	//
	// c2c9baae-9351-4c49-a8cb-6f24a83a8718
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
}

func (s GetScriptResponseBodyDataPublishedVersionTranscriberConfigNlsAccessProfile) String() string {
	return dara.Prettify(s)
}

func (s GetScriptResponseBodyDataPublishedVersionTranscriberConfigNlsAccessProfile) GoString() string {
	return s.String()
}

func (s *GetScriptResponseBodyDataPublishedVersionTranscriberConfigNlsAccessProfile) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *GetScriptResponseBodyDataPublishedVersionTranscriberConfigNlsAccessProfile) SetAccessProfileId(v string) *GetScriptResponseBodyDataPublishedVersionTranscriberConfigNlsAccessProfile {
	s.AccessProfileId = &v
	return s
}

func (s *GetScriptResponseBodyDataPublishedVersionTranscriberConfigNlsAccessProfile) Validate() error {
	return dara.Validate(s)
}
