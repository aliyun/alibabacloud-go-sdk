// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetApplicationResponseBody
	GetCode() *string
	SetData(v *GetApplicationResponseBodyData) *GetApplicationResponseBody
	GetData() *GetApplicationResponseBodyData
	SetHttpStatusCode(v int32) *GetApplicationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetApplicationResponseBody
	GetMessage() *string
	SetParams(v []*string) *GetApplicationResponseBody
	GetParams() []*string
	SetRequestId(v string) *GetApplicationResponseBody
	GetRequestId() *string
}

type GetApplicationResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetApplicationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D771A1B6-3D5F-174A-BEE1-98CE1000D337
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetApplicationResponseBody) GetData() *GetApplicationResponseBodyData {
	return s.Data
}

func (s *GetApplicationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetApplicationResponseBody) GetParams() []*string {
	return s.Params
}

func (s *GetApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApplicationResponseBody) SetCode(v string) *GetApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *GetApplicationResponseBody) SetData(v *GetApplicationResponseBodyData) *GetApplicationResponseBody {
	s.Data = v
	return s
}

func (s *GetApplicationResponseBody) SetHttpStatusCode(v int32) *GetApplicationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetApplicationResponseBody) SetMessage(v string) *GetApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *GetApplicationResponseBody) SetParams(v []*string) *GetApplicationResponseBody {
	s.Params = v
	return s
}

func (s *GetApplicationResponseBody) SetRequestId(v string) *GetApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApplicationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationResponseBodyData struct {
	// The application ID.
	//
	// example:
	//
	// a395011f-a247-400f-bc69-28796749fd52
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The call concurrency, which is the number of calls being made simultaneously.
	//
	// example:
	//
	// 10
	Concurrency *int32 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	// The time when the application was created.
	//
	// example:
	//
	// 1730081561000
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The application description.
	//
	// example:
	//
	// Describe this application
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The draft version configuration content.
	DraftVersion *GetApplicationResponseBodyDataDraftVersion `json:"DraftVersion,omitempty" xml:"DraftVersion,omitempty" type:"Struct"`
	// The application name.
	//
	// example:
	//
	// Test001
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The NLU access method.
	//
	// example:
	//
	// MANAGED
	NluAccessType *string `json:"NluAccessType,omitempty" xml:"NluAccessType,omitempty"`
	// The NLU engine.
	//
	// example:
	//
	// PROMPTS
	NluEngine *string `json:"NluEngine,omitempty" xml:"NluEngine,omitempty"`
	// The published application version.
	PublishedVersion *GetApplicationResponseBodyDataPublishedVersion `json:"PublishedVersion,omitempty" xml:"PublishedVersion,omitempty" type:"Struct"`
	Status           *string                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the application was last modified.
	//
	// example:
	//
	// 1730081561000
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s GetApplicationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyData) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GetApplicationResponseBodyData) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *GetApplicationResponseBodyData) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *GetApplicationResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetApplicationResponseBodyData) GetDraftVersion() *GetApplicationResponseBodyDataDraftVersion {
	return s.DraftVersion
}

func (s *GetApplicationResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetApplicationResponseBodyData) GetNluAccessType() *string {
	return s.NluAccessType
}

func (s *GetApplicationResponseBodyData) GetNluEngine() *string {
	return s.NluEngine
}

func (s *GetApplicationResponseBodyData) GetPublishedVersion() *GetApplicationResponseBodyDataPublishedVersion {
	return s.PublishedVersion
}

func (s *GetApplicationResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetApplicationResponseBodyData) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *GetApplicationResponseBodyData) SetApplicationId(v string) *GetApplicationResponseBodyData {
	s.ApplicationId = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetConcurrency(v int32) *GetApplicationResponseBodyData {
	s.Concurrency = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetCreatedTime(v int64) *GetApplicationResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetDescription(v string) *GetApplicationResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetDraftVersion(v *GetApplicationResponseBodyDataDraftVersion) *GetApplicationResponseBodyData {
	s.DraftVersion = v
	return s
}

func (s *GetApplicationResponseBodyData) SetName(v string) *GetApplicationResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetNluAccessType(v string) *GetApplicationResponseBodyData {
	s.NluAccessType = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetNluEngine(v string) *GetApplicationResponseBodyData {
	s.NluEngine = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetPublishedVersion(v *GetApplicationResponseBodyDataPublishedVersion) *GetApplicationResponseBodyData {
	s.PublishedVersion = v
	return s
}

func (s *GetApplicationResponseBodyData) SetStatus(v string) *GetApplicationResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetUpdatedTime(v int64) *GetApplicationResponseBodyData {
	s.UpdatedTime = &v
	return s
}

func (s *GetApplicationResponseBodyData) Validate() error {
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

type GetApplicationResponseBodyDataDraftVersion struct {
	// The interaction configuration.
	InteractionConfig *GetApplicationResponseBodyDataDraftVersionInteractionConfig `json:"InteractionConfig,omitempty" xml:"InteractionConfig,omitempty" type:"Struct"`
	LabelConfig       []*GetApplicationResponseBodyDataDraftVersionLabelConfig     `json:"LabelConfig,omitempty" xml:"LabelConfig,omitempty" type:"Repeated"`
	// The RAG configuration.
	RagConfig *GetApplicationResponseBodyDataDraftVersionRagConfig `json:"RagConfig,omitempty" xml:"RagConfig,omitempty" type:"Struct"`
	// The application model configuration.
	ScriptProfile *GetApplicationResponseBodyDataDraftVersionScriptProfile `json:"ScriptProfile,omitempty" xml:"ScriptProfile,omitempty" type:"Struct"`
	// The text-to-speech (TTS) configuration.
	SynthesizerConfig *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig `json:"SynthesizerConfig,omitempty" xml:"SynthesizerConfig,omitempty" type:"Struct"`
	// The tool configuration.
	ToolConfig *GetApplicationResponseBodyDataDraftVersionToolConfig `json:"ToolConfig,omitempty" xml:"ToolConfig,omitempty" type:"Struct"`
	// The automatic speech recognition (ASR) configuration.
	TranscriberConfig *GetApplicationResponseBodyDataDraftVersionTranscriberConfig `json:"TranscriberConfig,omitempty" xml:"TranscriberConfig,omitempty" type:"Struct"`
	// The version ID.
	//
	// example:
	//
	// 743219815472857088
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetApplicationResponseBodyDataDraftVersion) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataDraftVersion) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataDraftVersion) GetInteractionConfig() *GetApplicationResponseBodyDataDraftVersionInteractionConfig {
	return s.InteractionConfig
}

func (s *GetApplicationResponseBodyDataDraftVersion) GetLabelConfig() []*GetApplicationResponseBodyDataDraftVersionLabelConfig {
	return s.LabelConfig
}

func (s *GetApplicationResponseBodyDataDraftVersion) GetRagConfig() *GetApplicationResponseBodyDataDraftVersionRagConfig {
	return s.RagConfig
}

func (s *GetApplicationResponseBodyDataDraftVersion) GetScriptProfile() *GetApplicationResponseBodyDataDraftVersionScriptProfile {
	return s.ScriptProfile
}

func (s *GetApplicationResponseBodyDataDraftVersion) GetSynthesizerConfig() *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig {
	return s.SynthesizerConfig
}

func (s *GetApplicationResponseBodyDataDraftVersion) GetToolConfig() *GetApplicationResponseBodyDataDraftVersionToolConfig {
	return s.ToolConfig
}

func (s *GetApplicationResponseBodyDataDraftVersion) GetTranscriberConfig() *GetApplicationResponseBodyDataDraftVersionTranscriberConfig {
	return s.TranscriberConfig
}

func (s *GetApplicationResponseBodyDataDraftVersion) GetVersionId() *string {
	return s.VersionId
}

func (s *GetApplicationResponseBodyDataDraftVersion) SetInteractionConfig(v *GetApplicationResponseBodyDataDraftVersionInteractionConfig) *GetApplicationResponseBodyDataDraftVersion {
	s.InteractionConfig = v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersion) SetLabelConfig(v []*GetApplicationResponseBodyDataDraftVersionLabelConfig) *GetApplicationResponseBodyDataDraftVersion {
	s.LabelConfig = v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersion) SetRagConfig(v *GetApplicationResponseBodyDataDraftVersionRagConfig) *GetApplicationResponseBodyDataDraftVersion {
	s.RagConfig = v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersion) SetScriptProfile(v *GetApplicationResponseBodyDataDraftVersionScriptProfile) *GetApplicationResponseBodyDataDraftVersion {
	s.ScriptProfile = v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersion) SetSynthesizerConfig(v *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) *GetApplicationResponseBodyDataDraftVersion {
	s.SynthesizerConfig = v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersion) SetToolConfig(v *GetApplicationResponseBodyDataDraftVersionToolConfig) *GetApplicationResponseBodyDataDraftVersion {
	s.ToolConfig = v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersion) SetTranscriberConfig(v *GetApplicationResponseBodyDataDraftVersionTranscriberConfig) *GetApplicationResponseBodyDataDraftVersion {
	s.TranscriberConfig = v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersion) SetVersionId(v string) *GetApplicationResponseBodyDataDraftVersion {
	s.VersionId = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersion) Validate() error {
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
	if s.RagConfig != nil {
		if err := s.RagConfig.Validate(); err != nil {
			return err
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
	if s.ToolConfig != nil {
		if err := s.ToolConfig.Validate(); err != nil {
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

type GetApplicationResponseBodyDataDraftVersionInteractionConfig struct {
	BackgroundMusicId                *string                                                                           `json:"BackgroundMusicId,omitempty" xml:"BackgroundMusicId,omitempty"`
	EndConversationConfig            *GetApplicationResponseBodyDataDraftVersionInteractionConfigEndConversationConfig `json:"EndConversationConfig,omitempty" xml:"EndConversationConfig,omitempty" type:"Struct"`
	InitialGreetingDelayMilliseconds *int32                                                                            `json:"InitialGreetingDelayMilliseconds,omitempty" xml:"InitialGreetingDelayMilliseconds,omitempty"`
	// The silence detection configuration.
	SilenceDetectionConfig *GetApplicationResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig `json:"SilenceDetectionConfig,omitempty" xml:"SilenceDetectionConfig,omitempty" type:"Struct"`
}

func (s GetApplicationResponseBodyDataDraftVersionInteractionConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataDraftVersionInteractionConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataDraftVersionInteractionConfig) GetBackgroundMusicId() *string {
	return s.BackgroundMusicId
}

func (s *GetApplicationResponseBodyDataDraftVersionInteractionConfig) GetEndConversationConfig() *GetApplicationResponseBodyDataDraftVersionInteractionConfigEndConversationConfig {
	return s.EndConversationConfig
}

func (s *GetApplicationResponseBodyDataDraftVersionInteractionConfig) GetInitialGreetingDelayMilliseconds() *int32 {
	return s.InitialGreetingDelayMilliseconds
}

func (s *GetApplicationResponseBodyDataDraftVersionInteractionConfig) GetSilenceDetectionConfig() *GetApplicationResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig {
	return s.SilenceDetectionConfig
}

func (s *GetApplicationResponseBodyDataDraftVersionInteractionConfig) SetBackgroundMusicId(v string) *GetApplicationResponseBodyDataDraftVersionInteractionConfig {
	s.BackgroundMusicId = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionInteractionConfig) SetEndConversationConfig(v *GetApplicationResponseBodyDataDraftVersionInteractionConfigEndConversationConfig) *GetApplicationResponseBodyDataDraftVersionInteractionConfig {
	s.EndConversationConfig = v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionInteractionConfig) SetInitialGreetingDelayMilliseconds(v int32) *GetApplicationResponseBodyDataDraftVersionInteractionConfig {
	s.InitialGreetingDelayMilliseconds = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionInteractionConfig) SetSilenceDetectionConfig(v *GetApplicationResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig) *GetApplicationResponseBodyDataDraftVersionInteractionConfig {
	s.SilenceDetectionConfig = v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionInteractionConfig) Validate() error {
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

type GetApplicationResponseBodyDataDraftVersionInteractionConfigEndConversationConfig struct {
	Delay    *int32                                                                                      `json:"Delay,omitempty" xml:"Delay,omitempty"`
	Triggers []*GetApplicationResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers `json:"Triggers,omitempty" xml:"Triggers,omitempty" type:"Repeated"`
}

func (s GetApplicationResponseBodyDataDraftVersionInteractionConfigEndConversationConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataDraftVersionInteractionConfigEndConversationConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataDraftVersionInteractionConfigEndConversationConfig) GetDelay() *int32 {
	return s.Delay
}

func (s *GetApplicationResponseBodyDataDraftVersionInteractionConfigEndConversationConfig) GetTriggers() []*GetApplicationResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers {
	return s.Triggers
}

func (s *GetApplicationResponseBodyDataDraftVersionInteractionConfigEndConversationConfig) SetDelay(v int32) *GetApplicationResponseBodyDataDraftVersionInteractionConfigEndConversationConfig {
	s.Delay = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionInteractionConfigEndConversationConfig) SetTriggers(v []*GetApplicationResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers) *GetApplicationResponseBodyDataDraftVersionInteractionConfigEndConversationConfig {
	s.Triggers = v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionInteractionConfigEndConversationConfig) Validate() error {
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

type GetApplicationResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers struct {
	ClosingStatement *string   `json:"ClosingStatement,omitempty" xml:"ClosingStatement,omitempty"`
	KeyWords         []*string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty" type:"Repeated"`
	TriggerType      *string   `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	TurnLimit        *int32    `json:"TurnLimit,omitempty" xml:"TurnLimit,omitempty"`
}

func (s GetApplicationResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers) GetClosingStatement() *string {
	return s.ClosingStatement
}

func (s *GetApplicationResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers) GetKeyWords() []*string {
	return s.KeyWords
}

func (s *GetApplicationResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers) GetTriggerType() *string {
	return s.TriggerType
}

func (s *GetApplicationResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers) GetTurnLimit() *int32 {
	return s.TurnLimit
}

func (s *GetApplicationResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers) SetClosingStatement(v string) *GetApplicationResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers {
	s.ClosingStatement = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers) SetKeyWords(v []*string) *GetApplicationResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers {
	s.KeyWords = v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers) SetTriggerType(v string) *GetApplicationResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers {
	s.TriggerType = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers) SetTurnLimit(v int32) *GetApplicationResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers {
	s.TurnLimit = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionInteractionConfigEndConversationConfigTriggers) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig struct {
	MaxRepeats *int32 `json:"MaxRepeats,omitempty" xml:"MaxRepeats,omitempty"`
	// The task execution timeout period, in seconds.
	//
	// example:
	//
	// 3
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s GetApplicationResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig) GetMaxRepeats() *int32 {
	return s.MaxRepeats
}

func (s *GetApplicationResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig) GetTimeout() *int32 {
	return s.Timeout
}

func (s *GetApplicationResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig) SetMaxRepeats(v int32) *GetApplicationResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig {
	s.MaxRepeats = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig) SetTimeout(v int32) *GetApplicationResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig {
	s.Timeout = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyDataDraftVersionLabelConfig struct {
	CandidateValues []*string `json:"CandidateValues,omitempty" xml:"CandidateValues,omitempty" type:"Repeated"`
	Description     *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Name            *string   `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetApplicationResponseBodyDataDraftVersionLabelConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataDraftVersionLabelConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataDraftVersionLabelConfig) GetCandidateValues() []*string {
	return s.CandidateValues
}

func (s *GetApplicationResponseBodyDataDraftVersionLabelConfig) GetDescription() *string {
	return s.Description
}

func (s *GetApplicationResponseBodyDataDraftVersionLabelConfig) GetName() *string {
	return s.Name
}

func (s *GetApplicationResponseBodyDataDraftVersionLabelConfig) SetCandidateValues(v []*string) *GetApplicationResponseBodyDataDraftVersionLabelConfig {
	s.CandidateValues = v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionLabelConfig) SetDescription(v string) *GetApplicationResponseBodyDataDraftVersionLabelConfig {
	s.Description = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionLabelConfig) SetName(v string) *GetApplicationResponseBodyDataDraftVersionLabelConfig {
	s.Name = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionLabelConfig) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyDataDraftVersionRagConfig struct {
	// Specifies whether RAG is enabled.
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The list of knowledge base IDs.
	KnowledgeBaseIds []*string `json:"KnowledgeBaseIds,omitempty" xml:"KnowledgeBaseIds,omitempty" type:"Repeated"`
	// The maximum concatenation length of RAG content.
	//
	// example:
	//
	// 2000
	MaxContentLength *int32 `json:"MaxContentLength,omitempty" xml:"MaxContentLength,omitempty"`
	// The RAG engine.
	//
	// example:
	//
	// BAILIAN
	RagEngine *string `json:"RagEngine,omitempty" xml:"RagEngine,omitempty"`
	// The maximum number of data entries to retrieve.
	//
	// example:
	//
	// 5
	TopN *int32 `json:"TopN,omitempty" xml:"TopN,omitempty"`
}

func (s GetApplicationResponseBodyDataDraftVersionRagConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataDraftVersionRagConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataDraftVersionRagConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetApplicationResponseBodyDataDraftVersionRagConfig) GetKnowledgeBaseIds() []*string {
	return s.KnowledgeBaseIds
}

func (s *GetApplicationResponseBodyDataDraftVersionRagConfig) GetMaxContentLength() *int32 {
	return s.MaxContentLength
}

func (s *GetApplicationResponseBodyDataDraftVersionRagConfig) GetRagEngine() *string {
	return s.RagEngine
}

func (s *GetApplicationResponseBodyDataDraftVersionRagConfig) GetTopN() *int32 {
	return s.TopN
}

func (s *GetApplicationResponseBodyDataDraftVersionRagConfig) SetEnabled(v bool) *GetApplicationResponseBodyDataDraftVersionRagConfig {
	s.Enabled = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionRagConfig) SetKnowledgeBaseIds(v []*string) *GetApplicationResponseBodyDataDraftVersionRagConfig {
	s.KnowledgeBaseIds = v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionRagConfig) SetMaxContentLength(v int32) *GetApplicationResponseBodyDataDraftVersionRagConfig {
	s.MaxContentLength = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionRagConfig) SetRagEngine(v string) *GetApplicationResponseBodyDataDraftVersionRagConfig {
	s.RagEngine = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionRagConfig) SetTopN(v int32) *GetApplicationResponseBodyDataDraftVersionRagConfig {
	s.TopN = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionRagConfig) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyDataDraftVersionScriptProfile struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The agent configuration information.
	AgentProfile *GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile `json:"AgentProfile,omitempty" xml:"AgentProfile,omitempty" type:"Struct"`
	ChatbotId    *string                                                              `json:"ChatbotId,omitempty" xml:"ChatbotId,omitempty"`
	FunctionMeta *GetApplicationResponseBodyDataDraftVersionScriptProfileFunctionMeta `json:"FunctionMeta,omitempty" xml:"FunctionMeta,omitempty" type:"Struct"`
	// The model.
	//
	// example:
	//
	// qwen-plus
	Model            *string                                                                  `json:"Model,omitempty" xml:"Model,omitempty"`
	NluAccessProfile *GetApplicationResponseBodyDataDraftVersionScriptProfileNluAccessProfile `json:"NluAccessProfile,omitempty" xml:"NluAccessProfile,omitempty" type:"Struct"`
	NluAccessType    *string                                                                  `json:"NluAccessType,omitempty" xml:"NluAccessType,omitempty"`
	NluEngine        *string                                                                  `json:"NluEngine,omitempty" xml:"NluEngine,omitempty"`
	OmniModel        *bool                                                                    `json:"OmniModel,omitempty" xml:"OmniModel,omitempty"`
	// The probability threshold for nucleus sampling during generation.
	//
	// > - For example, when the value is set to 0.8, only the smallest set of most likely tokens whose cumulative probability is greater than or equal to 0.8 is retained as the candidate set.
	//
	//      	- Valid values: (0, 1.0). A higher value increases randomness in generation. A lower value increases determinism in generation.
	//
	// example:
	//
	// 0.8
	Temperature *string `json:"Temperature,omitempty" xml:"Temperature,omitempty"`
	// Controls the randomness and diversity of model responses.
	//
	// > - Specifically, the temperature value controls the degree of smoothing applied to the probability distribution of each candidate token during text generation. A higher temperature value flattens the probability distribution, allowing more low-probability tokens to be selected, which produces more diverse results. A lower temperature value sharpens the probability distribution, making high-probability tokens more likely to be selected, which produces more deterministic results.
	//
	// > - Valid values: [0, 2). Setting the value to 0 is not recommended because it is meaningless.
	//
	// example:
	//
	// 0.1
	TopP *string `json:"TopP,omitempty" xml:"TopP,omitempty"`
}

func (s GetApplicationResponseBodyDataDraftVersionScriptProfile) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataDraftVersionScriptProfile) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfile) GetAgentKey() *string {
	return s.AgentKey
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfile) GetAgentProfile() *GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile {
	return s.AgentProfile
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfile) GetChatbotId() *string {
	return s.ChatbotId
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfile) GetFunctionMeta() *GetApplicationResponseBodyDataDraftVersionScriptProfileFunctionMeta {
	return s.FunctionMeta
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfile) GetModel() *string {
	return s.Model
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfile) GetNluAccessProfile() *GetApplicationResponseBodyDataDraftVersionScriptProfileNluAccessProfile {
	return s.NluAccessProfile
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfile) GetNluAccessType() *string {
	return s.NluAccessType
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfile) GetNluEngine() *string {
	return s.NluEngine
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfile) GetOmniModel() *bool {
	return s.OmniModel
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfile) GetTemperature() *string {
	return s.Temperature
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfile) GetTopP() *string {
	return s.TopP
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfile) SetAgentKey(v string) *GetApplicationResponseBodyDataDraftVersionScriptProfile {
	s.AgentKey = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfile) SetAgentProfile(v *GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile) *GetApplicationResponseBodyDataDraftVersionScriptProfile {
	s.AgentProfile = v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfile) SetChatbotId(v string) *GetApplicationResponseBodyDataDraftVersionScriptProfile {
	s.ChatbotId = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfile) SetFunctionMeta(v *GetApplicationResponseBodyDataDraftVersionScriptProfileFunctionMeta) *GetApplicationResponseBodyDataDraftVersionScriptProfile {
	s.FunctionMeta = v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfile) SetModel(v string) *GetApplicationResponseBodyDataDraftVersionScriptProfile {
	s.Model = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfile) SetNluAccessProfile(v *GetApplicationResponseBodyDataDraftVersionScriptProfileNluAccessProfile) *GetApplicationResponseBodyDataDraftVersionScriptProfile {
	s.NluAccessProfile = v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfile) SetNluAccessType(v string) *GetApplicationResponseBodyDataDraftVersionScriptProfile {
	s.NluAccessType = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfile) SetNluEngine(v string) *GetApplicationResponseBodyDataDraftVersionScriptProfile {
	s.NluEngine = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfile) SetOmniModel(v bool) *GetApplicationResponseBodyDataDraftVersionScriptProfile {
	s.OmniModel = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfile) SetTemperature(v string) *GetApplicationResponseBodyDataDraftVersionScriptProfile {
	s.Temperature = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfile) SetTopP(v string) *GetApplicationResponseBodyDataDraftVersionScriptProfile {
	s.TopP = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfile) Validate() error {
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

type GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile struct {
	// The agent configuration ID.
	//
	// example:
	//
	// 6a50b67072d44788951de29758432d94
	AgentProfileId *string `json:"AgentProfileId,omitempty" xml:"AgentProfileId,omitempty"`
	// The agent description.
	//
	// example:
	//
	// Chatbot
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The prompt in JSON format.
	//
	// example:
	//
	// {"prompts":"I am a chatbot."}
	PromptsJson *string `json:"PromptsJson,omitempty" xml:"PromptsJson,omitempty"`
	// The application template ID.
	//
	// example:
	//
	// SFM_PROMPTS_DEFAULT
	ScriptProfileTemplateId *string `json:"ScriptProfileTemplateId,omitempty" xml:"ScriptProfileTemplateId,omitempty"`
}

func (s GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile) GetAgentProfileId() *string {
	return s.AgentProfileId
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile) GetDescription() *string {
	return s.Description
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile) GetName() *string {
	return s.Name
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile) GetPromptsJson() *string {
	return s.PromptsJson
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile) GetScriptProfileTemplateId() *string {
	return s.ScriptProfileTemplateId
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile) SetAgentProfileId(v string) *GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile {
	s.AgentProfileId = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile) SetDescription(v string) *GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile {
	s.Description = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile) SetName(v string) *GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile {
	s.Name = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile) SetPromptsJson(v string) *GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile {
	s.PromptsJson = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile) SetScriptProfileTemplateId(v string) *GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile {
	s.ScriptProfileTemplateId = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyDataDraftVersionScriptProfileFunctionMeta struct {
	FunctionId      *string `json:"FunctionId,omitempty" xml:"FunctionId,omitempty"`
	FunctionName    *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	HttpTriggerName *string `json:"HttpTriggerName,omitempty" xml:"HttpTriggerName,omitempty"`
	HttpTriggerUrl  *string `json:"HttpTriggerUrl,omitempty" xml:"HttpTriggerUrl,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetApplicationResponseBodyDataDraftVersionScriptProfileFunctionMeta) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataDraftVersionScriptProfileFunctionMeta) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfileFunctionMeta) GetFunctionId() *string {
	return s.FunctionId
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfileFunctionMeta) GetFunctionName() *string {
	return s.FunctionName
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfileFunctionMeta) GetHttpTriggerName() *string {
	return s.HttpTriggerName
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfileFunctionMeta) GetHttpTriggerUrl() *string {
	return s.HttpTriggerUrl
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfileFunctionMeta) GetRegionId() *string {
	return s.RegionId
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfileFunctionMeta) SetFunctionId(v string) *GetApplicationResponseBodyDataDraftVersionScriptProfileFunctionMeta {
	s.FunctionId = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfileFunctionMeta) SetFunctionName(v string) *GetApplicationResponseBodyDataDraftVersionScriptProfileFunctionMeta {
	s.FunctionName = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfileFunctionMeta) SetHttpTriggerName(v string) *GetApplicationResponseBodyDataDraftVersionScriptProfileFunctionMeta {
	s.HttpTriggerName = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfileFunctionMeta) SetHttpTriggerUrl(v string) *GetApplicationResponseBodyDataDraftVersionScriptProfileFunctionMeta {
	s.HttpTriggerUrl = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfileFunctionMeta) SetRegionId(v string) *GetApplicationResponseBodyDataDraftVersionScriptProfileFunctionMeta {
	s.RegionId = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfileFunctionMeta) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyDataDraftVersionScriptProfileNluAccessProfile struct {
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
}

func (s GetApplicationResponseBodyDataDraftVersionScriptProfileNluAccessProfile) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataDraftVersionScriptProfileNluAccessProfile) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfileNluAccessProfile) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfileNluAccessProfile) SetAccessProfileId(v string) *GetApplicationResponseBodyDataDraftVersionScriptProfileNluAccessProfile {
	s.AccessProfileId = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfileNluAccessProfile) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyDataDraftVersionSynthesizerConfig struct {
	Model            *string                                                                      `json:"Model,omitempty" xml:"Model,omitempty"`
	NlsAccessProfile *GetApplicationResponseBodyDataDraftVersionSynthesizerConfigNlsAccessProfile `json:"NlsAccessProfile,omitempty" xml:"NlsAccessProfile,omitempty" type:"Struct"`
	// The TTS invocation method.
	//
	// example:
	//
	// MANAGED
	NlsAccessType *string `json:"NlsAccessType,omitempty" xml:"NlsAccessType,omitempty"`
	// The TTS engine.
	//
	// example:
	//
	// ALIYUN
	NlsEngine *string `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
	// The pitch rate.
	//
	// > Valid values: -500 to 500.
	//
	// example:
	//
	// 5
	PitchRate *int32                                                                  `json:"PitchRate,omitempty" xml:"PitchRate,omitempty"`
	PronRules []*GetApplicationResponseBodyDataDraftVersionSynthesizerConfigPronRules `json:"PronRules,omitempty" xml:"PronRules,omitempty" type:"Repeated"`
	// The speech rate.
	//
	// > Valid values: -500 to 500.
	//
	// example:
	//
	// 1
	SpeechRate *int32 `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	// The voice.
	//
	// example:
	//
	// aixia
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
	// The volume.
	//
	// example:
	//
	// 50
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) GetModel() *string {
	return s.Model
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) GetNlsAccessProfile() *GetApplicationResponseBodyDataDraftVersionSynthesizerConfigNlsAccessProfile {
	return s.NlsAccessProfile
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) GetNlsAccessType() *string {
	return s.NlsAccessType
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) GetPitchRate() *int32 {
	return s.PitchRate
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) GetPronRules() []*GetApplicationResponseBodyDataDraftVersionSynthesizerConfigPronRules {
	return s.PronRules
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) GetSpeechRate() *int32 {
	return s.SpeechRate
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) GetVoice() *string {
	return s.Voice
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) GetVolume() *int32 {
	return s.Volume
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) SetModel(v string) *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig {
	s.Model = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) SetNlsAccessProfile(v *GetApplicationResponseBodyDataDraftVersionSynthesizerConfigNlsAccessProfile) *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig {
	s.NlsAccessProfile = v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) SetNlsAccessType(v string) *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig {
	s.NlsAccessType = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) SetNlsEngine(v string) *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig {
	s.NlsEngine = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) SetPitchRate(v int32) *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig {
	s.PitchRate = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) SetPronRules(v []*GetApplicationResponseBodyDataDraftVersionSynthesizerConfigPronRules) *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig {
	s.PronRules = v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) SetSpeechRate(v int32) *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig {
	s.SpeechRate = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) SetVoice(v string) *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig {
	s.Voice = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) SetVolume(v int32) *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig {
	s.Volume = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) Validate() error {
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

type GetApplicationResponseBodyDataDraftVersionSynthesizerConfigNlsAccessProfile struct {
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
}

func (s GetApplicationResponseBodyDataDraftVersionSynthesizerConfigNlsAccessProfile) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataDraftVersionSynthesizerConfigNlsAccessProfile) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfigNlsAccessProfile) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfigNlsAccessProfile) SetAccessProfileId(v string) *GetApplicationResponseBodyDataDraftVersionSynthesizerConfigNlsAccessProfile {
	s.AccessProfileId = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfigNlsAccessProfile) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyDataDraftVersionSynthesizerConfigPronRules struct {
	Pattern     *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	Replacement *string `json:"Replacement,omitempty" xml:"Replacement,omitempty"`
}

func (s GetApplicationResponseBodyDataDraftVersionSynthesizerConfigPronRules) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataDraftVersionSynthesizerConfigPronRules) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfigPronRules) GetPattern() *string {
	return s.Pattern
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfigPronRules) GetReplacement() *string {
	return s.Replacement
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfigPronRules) SetPattern(v string) *GetApplicationResponseBodyDataDraftVersionSynthesizerConfigPronRules {
	s.Pattern = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfigPronRules) SetReplacement(v string) *GetApplicationResponseBodyDataDraftVersionSynthesizerConfigPronRules {
	s.Replacement = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfigPronRules) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyDataDraftVersionToolConfig struct {
	// The list of MCP server configurations.
	McpServers []*GetApplicationResponseBodyDataDraftVersionToolConfigMcpServers `json:"McpServers,omitempty" xml:"McpServers,omitempty" type:"Repeated"`
}

func (s GetApplicationResponseBodyDataDraftVersionToolConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataDraftVersionToolConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataDraftVersionToolConfig) GetMcpServers() []*GetApplicationResponseBodyDataDraftVersionToolConfigMcpServers {
	return s.McpServers
}

func (s *GetApplicationResponseBodyDataDraftVersionToolConfig) SetMcpServers(v []*GetApplicationResponseBodyDataDraftVersionToolConfigMcpServers) *GetApplicationResponseBodyDataDraftVersionToolConfig {
	s.McpServers = v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionToolConfig) Validate() error {
	if s.McpServers != nil {
		for _, item := range s.McpServers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetApplicationResponseBodyDataDraftVersionToolConfigMcpServers struct {
	// The base URL.
	//
	// example:
	//
	// https://example.com
	BaseUrl *string `json:"BaseUrl,omitempty" xml:"BaseUrl,omitempty"`
	// The name.
	//
	// example:
	//
	// phone-ai-call
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The SSE endpoint.
	//
	// example:
	//
	// /phone-ai-call/mcp/sse?key=value
	SseEndpoint *string `json:"SseEndpoint,omitempty" xml:"SseEndpoint,omitempty"`
}

func (s GetApplicationResponseBodyDataDraftVersionToolConfigMcpServers) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataDraftVersionToolConfigMcpServers) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataDraftVersionToolConfigMcpServers) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *GetApplicationResponseBodyDataDraftVersionToolConfigMcpServers) GetName() *string {
	return s.Name
}

func (s *GetApplicationResponseBodyDataDraftVersionToolConfigMcpServers) GetSseEndpoint() *string {
	return s.SseEndpoint
}

func (s *GetApplicationResponseBodyDataDraftVersionToolConfigMcpServers) SetBaseUrl(v string) *GetApplicationResponseBodyDataDraftVersionToolConfigMcpServers {
	s.BaseUrl = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionToolConfigMcpServers) SetName(v string) *GetApplicationResponseBodyDataDraftVersionToolConfigMcpServers {
	s.Name = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionToolConfigMcpServers) SetSseEndpoint(v string) *GetApplicationResponseBodyDataDraftVersionToolConfigMcpServers {
	s.SseEndpoint = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionToolConfigMcpServers) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyDataDraftVersionTranscriberConfig struct {
	CorrectionRules   []*GetApplicationResponseBodyDataDraftVersionTranscriberConfigCorrectionRules `json:"CorrectionRules,omitempty" xml:"CorrectionRules,omitempty" type:"Repeated"`
	CustomizationId   *string                                                                       `json:"CustomizationId,omitempty" xml:"CustomizationId,omitempty"`
	EndSilenceTimeout *int32                                                                        `json:"EndSilenceTimeout,omitempty" xml:"EndSilenceTimeout,omitempty"`
	Model             *string                                                                       `json:"Model,omitempty" xml:"Model,omitempty"`
	NlsAccessProfile  *GetApplicationResponseBodyDataDraftVersionTranscriberConfigNlsAccessProfile  `json:"NlsAccessProfile,omitempty" xml:"NlsAccessProfile,omitempty" type:"Struct"`
	// The ASR invocation method.
	//
	// example:
	//
	// MANAGED
	NlsAccessType *string `json:"NlsAccessType,omitempty" xml:"NlsAccessType,omitempty"`
	// The ASR engine.
	//
	// example:
	//
	// ALIYUN
	NlsEngine            *string `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
	SpeechNoiseThreshold *int32  `json:"SpeechNoiseThreshold,omitempty" xml:"SpeechNoiseThreshold,omitempty"`
	VocabularyId         *string `json:"VocabularyId,omitempty" xml:"VocabularyId,omitempty"`
}

func (s GetApplicationResponseBodyDataDraftVersionTranscriberConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataDraftVersionTranscriberConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataDraftVersionTranscriberConfig) GetCorrectionRules() []*GetApplicationResponseBodyDataDraftVersionTranscriberConfigCorrectionRules {
	return s.CorrectionRules
}

func (s *GetApplicationResponseBodyDataDraftVersionTranscriberConfig) GetCustomizationId() *string {
	return s.CustomizationId
}

func (s *GetApplicationResponseBodyDataDraftVersionTranscriberConfig) GetEndSilenceTimeout() *int32 {
	return s.EndSilenceTimeout
}

func (s *GetApplicationResponseBodyDataDraftVersionTranscriberConfig) GetModel() *string {
	return s.Model
}

func (s *GetApplicationResponseBodyDataDraftVersionTranscriberConfig) GetNlsAccessProfile() *GetApplicationResponseBodyDataDraftVersionTranscriberConfigNlsAccessProfile {
	return s.NlsAccessProfile
}

func (s *GetApplicationResponseBodyDataDraftVersionTranscriberConfig) GetNlsAccessType() *string {
	return s.NlsAccessType
}

func (s *GetApplicationResponseBodyDataDraftVersionTranscriberConfig) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *GetApplicationResponseBodyDataDraftVersionTranscriberConfig) GetSpeechNoiseThreshold() *int32 {
	return s.SpeechNoiseThreshold
}

func (s *GetApplicationResponseBodyDataDraftVersionTranscriberConfig) GetVocabularyId() *string {
	return s.VocabularyId
}

func (s *GetApplicationResponseBodyDataDraftVersionTranscriberConfig) SetCorrectionRules(v []*GetApplicationResponseBodyDataDraftVersionTranscriberConfigCorrectionRules) *GetApplicationResponseBodyDataDraftVersionTranscriberConfig {
	s.CorrectionRules = v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionTranscriberConfig) SetCustomizationId(v string) *GetApplicationResponseBodyDataDraftVersionTranscriberConfig {
	s.CustomizationId = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionTranscriberConfig) SetEndSilenceTimeout(v int32) *GetApplicationResponseBodyDataDraftVersionTranscriberConfig {
	s.EndSilenceTimeout = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionTranscriberConfig) SetModel(v string) *GetApplicationResponseBodyDataDraftVersionTranscriberConfig {
	s.Model = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionTranscriberConfig) SetNlsAccessProfile(v *GetApplicationResponseBodyDataDraftVersionTranscriberConfigNlsAccessProfile) *GetApplicationResponseBodyDataDraftVersionTranscriberConfig {
	s.NlsAccessProfile = v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionTranscriberConfig) SetNlsAccessType(v string) *GetApplicationResponseBodyDataDraftVersionTranscriberConfig {
	s.NlsAccessType = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionTranscriberConfig) SetNlsEngine(v string) *GetApplicationResponseBodyDataDraftVersionTranscriberConfig {
	s.NlsEngine = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionTranscriberConfig) SetSpeechNoiseThreshold(v int32) *GetApplicationResponseBodyDataDraftVersionTranscriberConfig {
	s.SpeechNoiseThreshold = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionTranscriberConfig) SetVocabularyId(v string) *GetApplicationResponseBodyDataDraftVersionTranscriberConfig {
	s.VocabularyId = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionTranscriberConfig) Validate() error {
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

type GetApplicationResponseBodyDataDraftVersionTranscriberConfigCorrectionRules struct {
	Pattern     *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	Replacement *string `json:"Replacement,omitempty" xml:"Replacement,omitempty"`
}

func (s GetApplicationResponseBodyDataDraftVersionTranscriberConfigCorrectionRules) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataDraftVersionTranscriberConfigCorrectionRules) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataDraftVersionTranscriberConfigCorrectionRules) GetPattern() *string {
	return s.Pattern
}

func (s *GetApplicationResponseBodyDataDraftVersionTranscriberConfigCorrectionRules) GetReplacement() *string {
	return s.Replacement
}

func (s *GetApplicationResponseBodyDataDraftVersionTranscriberConfigCorrectionRules) SetPattern(v string) *GetApplicationResponseBodyDataDraftVersionTranscriberConfigCorrectionRules {
	s.Pattern = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionTranscriberConfigCorrectionRules) SetReplacement(v string) *GetApplicationResponseBodyDataDraftVersionTranscriberConfigCorrectionRules {
	s.Replacement = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionTranscriberConfigCorrectionRules) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyDataDraftVersionTranscriberConfigNlsAccessProfile struct {
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
}

func (s GetApplicationResponseBodyDataDraftVersionTranscriberConfigNlsAccessProfile) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataDraftVersionTranscriberConfigNlsAccessProfile) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataDraftVersionTranscriberConfigNlsAccessProfile) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *GetApplicationResponseBodyDataDraftVersionTranscriberConfigNlsAccessProfile) SetAccessProfileId(v string) *GetApplicationResponseBodyDataDraftVersionTranscriberConfigNlsAccessProfile {
	s.AccessProfileId = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionTranscriberConfigNlsAccessProfile) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyDataPublishedVersion struct {
	// The interaction configuration.
	InteractionConfig *GetApplicationResponseBodyDataPublishedVersionInteractionConfig `json:"InteractionConfig,omitempty" xml:"InteractionConfig,omitempty" type:"Struct"`
	LabelConfig       []*GetApplicationResponseBodyDataPublishedVersionLabelConfig     `json:"LabelConfig,omitempty" xml:"LabelConfig,omitempty" type:"Repeated"`
	// The RAG configuration.
	RagConfig *GetApplicationResponseBodyDataPublishedVersionRagConfig `json:"RagConfig,omitempty" xml:"RagConfig,omitempty" type:"Struct"`
	// The application model configuration.
	ScriptProfile *GetApplicationResponseBodyDataPublishedVersionScriptProfile `json:"ScriptProfile,omitempty" xml:"ScriptProfile,omitempty" type:"Struct"`
	// The text-to-speech (TTS) configuration.
	SynthesizerConfig *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig `json:"SynthesizerConfig,omitempty" xml:"SynthesizerConfig,omitempty" type:"Struct"`
	// The tool configuration.
	ToolConfig *GetApplicationResponseBodyDataPublishedVersionToolConfig `json:"ToolConfig,omitempty" xml:"ToolConfig,omitempty" type:"Struct"`
	// The automatic speech recognition (ASR) configuration.
	TranscriberConfig *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig `json:"TranscriberConfig,omitempty" xml:"TranscriberConfig,omitempty" type:"Struct"`
	// The version ID.
	//
	// example:
	//
	// 47889c1f-dd3f-4ace-9587-a13a3563e678
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetApplicationResponseBodyDataPublishedVersion) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataPublishedVersion) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataPublishedVersion) GetInteractionConfig() *GetApplicationResponseBodyDataPublishedVersionInteractionConfig {
	return s.InteractionConfig
}

func (s *GetApplicationResponseBodyDataPublishedVersion) GetLabelConfig() []*GetApplicationResponseBodyDataPublishedVersionLabelConfig {
	return s.LabelConfig
}

func (s *GetApplicationResponseBodyDataPublishedVersion) GetRagConfig() *GetApplicationResponseBodyDataPublishedVersionRagConfig {
	return s.RagConfig
}

func (s *GetApplicationResponseBodyDataPublishedVersion) GetScriptProfile() *GetApplicationResponseBodyDataPublishedVersionScriptProfile {
	return s.ScriptProfile
}

func (s *GetApplicationResponseBodyDataPublishedVersion) GetSynthesizerConfig() *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig {
	return s.SynthesizerConfig
}

func (s *GetApplicationResponseBodyDataPublishedVersion) GetToolConfig() *GetApplicationResponseBodyDataPublishedVersionToolConfig {
	return s.ToolConfig
}

func (s *GetApplicationResponseBodyDataPublishedVersion) GetTranscriberConfig() *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig {
	return s.TranscriberConfig
}

func (s *GetApplicationResponseBodyDataPublishedVersion) GetVersionId() *string {
	return s.VersionId
}

func (s *GetApplicationResponseBodyDataPublishedVersion) SetInteractionConfig(v *GetApplicationResponseBodyDataPublishedVersionInteractionConfig) *GetApplicationResponseBodyDataPublishedVersion {
	s.InteractionConfig = v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersion) SetLabelConfig(v []*GetApplicationResponseBodyDataPublishedVersionLabelConfig) *GetApplicationResponseBodyDataPublishedVersion {
	s.LabelConfig = v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersion) SetRagConfig(v *GetApplicationResponseBodyDataPublishedVersionRagConfig) *GetApplicationResponseBodyDataPublishedVersion {
	s.RagConfig = v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersion) SetScriptProfile(v *GetApplicationResponseBodyDataPublishedVersionScriptProfile) *GetApplicationResponseBodyDataPublishedVersion {
	s.ScriptProfile = v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersion) SetSynthesizerConfig(v *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) *GetApplicationResponseBodyDataPublishedVersion {
	s.SynthesizerConfig = v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersion) SetToolConfig(v *GetApplicationResponseBodyDataPublishedVersionToolConfig) *GetApplicationResponseBodyDataPublishedVersion {
	s.ToolConfig = v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersion) SetTranscriberConfig(v *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig) *GetApplicationResponseBodyDataPublishedVersion {
	s.TranscriberConfig = v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersion) SetVersionId(v string) *GetApplicationResponseBodyDataPublishedVersion {
	s.VersionId = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersion) Validate() error {
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
	if s.RagConfig != nil {
		if err := s.RagConfig.Validate(); err != nil {
			return err
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
	if s.ToolConfig != nil {
		if err := s.ToolConfig.Validate(); err != nil {
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

type GetApplicationResponseBodyDataPublishedVersionInteractionConfig struct {
	BackgroundMusicId                *string                                                                               `json:"BackgroundMusicId,omitempty" xml:"BackgroundMusicId,omitempty"`
	EndConversationConfig            *GetApplicationResponseBodyDataPublishedVersionInteractionConfigEndConversationConfig `json:"EndConversationConfig,omitempty" xml:"EndConversationConfig,omitempty" type:"Struct"`
	InitialGreetingDelayMilliseconds *int32                                                                                `json:"InitialGreetingDelayMilliseconds,omitempty" xml:"InitialGreetingDelayMilliseconds,omitempty"`
	// The silence detection configuration.
	SilenceDetectionConfig *GetApplicationResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig `json:"SilenceDetectionConfig,omitempty" xml:"SilenceDetectionConfig,omitempty" type:"Struct"`
}

func (s GetApplicationResponseBodyDataPublishedVersionInteractionConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataPublishedVersionInteractionConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataPublishedVersionInteractionConfig) GetBackgroundMusicId() *string {
	return s.BackgroundMusicId
}

func (s *GetApplicationResponseBodyDataPublishedVersionInteractionConfig) GetEndConversationConfig() *GetApplicationResponseBodyDataPublishedVersionInteractionConfigEndConversationConfig {
	return s.EndConversationConfig
}

func (s *GetApplicationResponseBodyDataPublishedVersionInteractionConfig) GetInitialGreetingDelayMilliseconds() *int32 {
	return s.InitialGreetingDelayMilliseconds
}

func (s *GetApplicationResponseBodyDataPublishedVersionInteractionConfig) GetSilenceDetectionConfig() *GetApplicationResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig {
	return s.SilenceDetectionConfig
}

func (s *GetApplicationResponseBodyDataPublishedVersionInteractionConfig) SetBackgroundMusicId(v string) *GetApplicationResponseBodyDataPublishedVersionInteractionConfig {
	s.BackgroundMusicId = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionInteractionConfig) SetEndConversationConfig(v *GetApplicationResponseBodyDataPublishedVersionInteractionConfigEndConversationConfig) *GetApplicationResponseBodyDataPublishedVersionInteractionConfig {
	s.EndConversationConfig = v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionInteractionConfig) SetInitialGreetingDelayMilliseconds(v int32) *GetApplicationResponseBodyDataPublishedVersionInteractionConfig {
	s.InitialGreetingDelayMilliseconds = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionInteractionConfig) SetSilenceDetectionConfig(v *GetApplicationResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig) *GetApplicationResponseBodyDataPublishedVersionInteractionConfig {
	s.SilenceDetectionConfig = v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionInteractionConfig) Validate() error {
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

type GetApplicationResponseBodyDataPublishedVersionInteractionConfigEndConversationConfig struct {
	Delay    *int32                                                                                          `json:"Delay,omitempty" xml:"Delay,omitempty"`
	Triggers []*GetApplicationResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers `json:"Triggers,omitempty" xml:"Triggers,omitempty" type:"Repeated"`
}

func (s GetApplicationResponseBodyDataPublishedVersionInteractionConfigEndConversationConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataPublishedVersionInteractionConfigEndConversationConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataPublishedVersionInteractionConfigEndConversationConfig) GetDelay() *int32 {
	return s.Delay
}

func (s *GetApplicationResponseBodyDataPublishedVersionInteractionConfigEndConversationConfig) GetTriggers() []*GetApplicationResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers {
	return s.Triggers
}

func (s *GetApplicationResponseBodyDataPublishedVersionInteractionConfigEndConversationConfig) SetDelay(v int32) *GetApplicationResponseBodyDataPublishedVersionInteractionConfigEndConversationConfig {
	s.Delay = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionInteractionConfigEndConversationConfig) SetTriggers(v []*GetApplicationResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers) *GetApplicationResponseBodyDataPublishedVersionInteractionConfigEndConversationConfig {
	s.Triggers = v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionInteractionConfigEndConversationConfig) Validate() error {
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

type GetApplicationResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers struct {
	ClosingStatement *string   `json:"ClosingStatement,omitempty" xml:"ClosingStatement,omitempty"`
	KeyWords         []*string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty" type:"Repeated"`
	TriggerType      *string   `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	TurnLimit        *int32    `json:"TurnLimit,omitempty" xml:"TurnLimit,omitempty"`
}

func (s GetApplicationResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers) GetClosingStatement() *string {
	return s.ClosingStatement
}

func (s *GetApplicationResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers) GetKeyWords() []*string {
	return s.KeyWords
}

func (s *GetApplicationResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers) GetTriggerType() *string {
	return s.TriggerType
}

func (s *GetApplicationResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers) GetTurnLimit() *int32 {
	return s.TurnLimit
}

func (s *GetApplicationResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers) SetClosingStatement(v string) *GetApplicationResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers {
	s.ClosingStatement = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers) SetKeyWords(v []*string) *GetApplicationResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers {
	s.KeyWords = v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers) SetTriggerType(v string) *GetApplicationResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers {
	s.TriggerType = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers) SetTurnLimit(v int32) *GetApplicationResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers {
	s.TurnLimit = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionInteractionConfigEndConversationConfigTriggers) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig struct {
	MaxRepeats *int32 `json:"MaxRepeats,omitempty" xml:"MaxRepeats,omitempty"`
	// The timeout period.
	//
	// example:
	//
	// 30
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s GetApplicationResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig) GetMaxRepeats() *int32 {
	return s.MaxRepeats
}

func (s *GetApplicationResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig) GetTimeout() *int32 {
	return s.Timeout
}

func (s *GetApplicationResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig) SetMaxRepeats(v int32) *GetApplicationResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig {
	s.MaxRepeats = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig) SetTimeout(v int32) *GetApplicationResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig {
	s.Timeout = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyDataPublishedVersionLabelConfig struct {
	CandidateValues []*string `json:"CandidateValues,omitempty" xml:"CandidateValues,omitempty" type:"Repeated"`
	Description     *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Name            *string   `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetApplicationResponseBodyDataPublishedVersionLabelConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataPublishedVersionLabelConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataPublishedVersionLabelConfig) GetCandidateValues() []*string {
	return s.CandidateValues
}

func (s *GetApplicationResponseBodyDataPublishedVersionLabelConfig) GetDescription() *string {
	return s.Description
}

func (s *GetApplicationResponseBodyDataPublishedVersionLabelConfig) GetName() *string {
	return s.Name
}

func (s *GetApplicationResponseBodyDataPublishedVersionLabelConfig) SetCandidateValues(v []*string) *GetApplicationResponseBodyDataPublishedVersionLabelConfig {
	s.CandidateValues = v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionLabelConfig) SetDescription(v string) *GetApplicationResponseBodyDataPublishedVersionLabelConfig {
	s.Description = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionLabelConfig) SetName(v string) *GetApplicationResponseBodyDataPublishedVersionLabelConfig {
	s.Name = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionLabelConfig) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyDataPublishedVersionRagConfig struct {
	// Specifies whether RAG is enabled.
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The list of knowledge base IDs.
	KnowledgeBaseIds []*string `json:"KnowledgeBaseIds,omitempty" xml:"KnowledgeBaseIds,omitempty" type:"Repeated"`
	// The maximum concatenation length of RAG content.
	//
	// example:
	//
	// 2000
	MaxContentLength *int32 `json:"MaxContentLength,omitempty" xml:"MaxContentLength,omitempty"`
	// The RAG engine.
	//
	// example:
	//
	// BAILIAN
	RagEngine *string `json:"RagEngine,omitempty" xml:"RagEngine,omitempty"`
	// The maximum number of data entries to retrieve.
	//
	// example:
	//
	// 5
	TopN *int32 `json:"TopN,omitempty" xml:"TopN,omitempty"`
}

func (s GetApplicationResponseBodyDataPublishedVersionRagConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataPublishedVersionRagConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataPublishedVersionRagConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetApplicationResponseBodyDataPublishedVersionRagConfig) GetKnowledgeBaseIds() []*string {
	return s.KnowledgeBaseIds
}

func (s *GetApplicationResponseBodyDataPublishedVersionRagConfig) GetMaxContentLength() *int32 {
	return s.MaxContentLength
}

func (s *GetApplicationResponseBodyDataPublishedVersionRagConfig) GetRagEngine() *string {
	return s.RagEngine
}

func (s *GetApplicationResponseBodyDataPublishedVersionRagConfig) GetTopN() *int32 {
	return s.TopN
}

func (s *GetApplicationResponseBodyDataPublishedVersionRagConfig) SetEnabled(v bool) *GetApplicationResponseBodyDataPublishedVersionRagConfig {
	s.Enabled = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionRagConfig) SetKnowledgeBaseIds(v []*string) *GetApplicationResponseBodyDataPublishedVersionRagConfig {
	s.KnowledgeBaseIds = v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionRagConfig) SetMaxContentLength(v int32) *GetApplicationResponseBodyDataPublishedVersionRagConfig {
	s.MaxContentLength = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionRagConfig) SetRagEngine(v string) *GetApplicationResponseBodyDataPublishedVersionRagConfig {
	s.RagEngine = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionRagConfig) SetTopN(v int32) *GetApplicationResponseBodyDataPublishedVersionRagConfig {
	s.TopN = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionRagConfig) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyDataPublishedVersionScriptProfile struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The agent configuration information.
	AgentProfile *GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile `json:"AgentProfile,omitempty" xml:"AgentProfile,omitempty" type:"Struct"`
	ChatbotId    *string                                                                  `json:"ChatbotId,omitempty" xml:"ChatbotId,omitempty"`
	FunctionMeta *GetApplicationResponseBodyDataPublishedVersionScriptProfileFunctionMeta `json:"FunctionMeta,omitempty" xml:"FunctionMeta,omitempty" type:"Struct"`
	// The model.
	//
	// example:
	//
	// qwen-plus
	Model            *string                                                                      `json:"Model,omitempty" xml:"Model,omitempty"`
	NluAccessProfile *GetApplicationResponseBodyDataPublishedVersionScriptProfileNluAccessProfile `json:"NluAccessProfile,omitempty" xml:"NluAccessProfile,omitempty" type:"Struct"`
	NluAccessType    *string                                                                      `json:"NluAccessType,omitempty" xml:"NluAccessType,omitempty"`
	NluEngine        *string                                                                      `json:"NluEngine,omitempty" xml:"NluEngine,omitempty"`
	OmniModel        *bool                                                                        `json:"OmniModel,omitempty" xml:"OmniModel,omitempty"`
	// The probability threshold for nucleus sampling during generation.
	//
	// > - For example, when the value is set to 0.8, only the smallest set of most likely tokens whose cumulative probability is greater than or equal to 0.8 is retained as the candidate set.
	//
	//      	- Valid values: (0, 1.0). A higher value increases randomness in generation. A lower value increases determinism in generation.
	//
	// example:
	//
	// 0.8
	Temperature *string `json:"Temperature,omitempty" xml:"Temperature,omitempty"`
	// Controls the randomness and diversity of model responses.
	//
	// > - Specifically, the temperature value controls the degree of smoothing applied to the probability distribution of each candidate token during text generation. A higher temperature value flattens the probability distribution, allowing more low-probability tokens to be selected, which produces more diverse results. A lower temperature value sharpens the probability distribution, making high-probability tokens more likely to be selected, which produces more deterministic results.
	//
	// > - Valid values: [0, 2). Setting the value to 0 is not recommended because it is meaningless.
	//
	// example:
	//
	// 0.1
	TopP *string `json:"TopP,omitempty" xml:"TopP,omitempty"`
}

func (s GetApplicationResponseBodyDataPublishedVersionScriptProfile) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataPublishedVersionScriptProfile) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfile) GetAgentKey() *string {
	return s.AgentKey
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfile) GetAgentProfile() *GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile {
	return s.AgentProfile
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfile) GetChatbotId() *string {
	return s.ChatbotId
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfile) GetFunctionMeta() *GetApplicationResponseBodyDataPublishedVersionScriptProfileFunctionMeta {
	return s.FunctionMeta
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfile) GetModel() *string {
	return s.Model
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfile) GetNluAccessProfile() *GetApplicationResponseBodyDataPublishedVersionScriptProfileNluAccessProfile {
	return s.NluAccessProfile
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfile) GetNluAccessType() *string {
	return s.NluAccessType
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfile) GetNluEngine() *string {
	return s.NluEngine
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfile) GetOmniModel() *bool {
	return s.OmniModel
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfile) GetTemperature() *string {
	return s.Temperature
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfile) GetTopP() *string {
	return s.TopP
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfile) SetAgentKey(v string) *GetApplicationResponseBodyDataPublishedVersionScriptProfile {
	s.AgentKey = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfile) SetAgentProfile(v *GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile) *GetApplicationResponseBodyDataPublishedVersionScriptProfile {
	s.AgentProfile = v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfile) SetChatbotId(v string) *GetApplicationResponseBodyDataPublishedVersionScriptProfile {
	s.ChatbotId = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfile) SetFunctionMeta(v *GetApplicationResponseBodyDataPublishedVersionScriptProfileFunctionMeta) *GetApplicationResponseBodyDataPublishedVersionScriptProfile {
	s.FunctionMeta = v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfile) SetModel(v string) *GetApplicationResponseBodyDataPublishedVersionScriptProfile {
	s.Model = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfile) SetNluAccessProfile(v *GetApplicationResponseBodyDataPublishedVersionScriptProfileNluAccessProfile) *GetApplicationResponseBodyDataPublishedVersionScriptProfile {
	s.NluAccessProfile = v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfile) SetNluAccessType(v string) *GetApplicationResponseBodyDataPublishedVersionScriptProfile {
	s.NluAccessType = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfile) SetNluEngine(v string) *GetApplicationResponseBodyDataPublishedVersionScriptProfile {
	s.NluEngine = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfile) SetOmniModel(v bool) *GetApplicationResponseBodyDataPublishedVersionScriptProfile {
	s.OmniModel = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfile) SetTemperature(v string) *GetApplicationResponseBodyDataPublishedVersionScriptProfile {
	s.Temperature = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfile) SetTopP(v string) *GetApplicationResponseBodyDataPublishedVersionScriptProfile {
	s.TopP = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfile) Validate() error {
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

type GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile struct {
	// The agent configuration ID.
	//
	// example:
	//
	// b97b6822dd624c32b6c2a54d717db718
	AgentProfileId *string `json:"AgentProfileId,omitempty" xml:"AgentProfileId,omitempty"`
	// The agent description.
	//
	// example:
	//
	// I am a chatbot
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The prompt in JSON format.
	//
	// example:
	//
	// {"prompts":"I am a chatbot."}
	PromptsJson *string `json:"PromptsJson,omitempty" xml:"PromptsJson,omitempty"`
	// The agent configuration template ID.
	//
	// example:
	//
	// SFM_PROMPTS_DEFAULT
	ScriptProfileTemplateId *string `json:"ScriptProfileTemplateId,omitempty" xml:"ScriptProfileTemplateId,omitempty"`
}

func (s GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile) GetAgentProfileId() *string {
	return s.AgentProfileId
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile) GetDescription() *string {
	return s.Description
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile) GetName() *string {
	return s.Name
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile) GetPromptsJson() *string {
	return s.PromptsJson
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile) GetScriptProfileTemplateId() *string {
	return s.ScriptProfileTemplateId
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile) SetAgentProfileId(v string) *GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile {
	s.AgentProfileId = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile) SetDescription(v string) *GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile {
	s.Description = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile) SetName(v string) *GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile {
	s.Name = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile) SetPromptsJson(v string) *GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile {
	s.PromptsJson = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile) SetScriptProfileTemplateId(v string) *GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile {
	s.ScriptProfileTemplateId = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyDataPublishedVersionScriptProfileFunctionMeta struct {
	FunctionId      *string `json:"FunctionId,omitempty" xml:"FunctionId,omitempty"`
	FunctionName    *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	HttpTriggerName *string `json:"HttpTriggerName,omitempty" xml:"HttpTriggerName,omitempty"`
	HttpTriggerUrl  *string `json:"HttpTriggerUrl,omitempty" xml:"HttpTriggerUrl,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetApplicationResponseBodyDataPublishedVersionScriptProfileFunctionMeta) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataPublishedVersionScriptProfileFunctionMeta) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfileFunctionMeta) GetFunctionId() *string {
	return s.FunctionId
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfileFunctionMeta) GetFunctionName() *string {
	return s.FunctionName
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfileFunctionMeta) GetHttpTriggerName() *string {
	return s.HttpTriggerName
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfileFunctionMeta) GetHttpTriggerUrl() *string {
	return s.HttpTriggerUrl
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfileFunctionMeta) GetRegionId() *string {
	return s.RegionId
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfileFunctionMeta) SetFunctionId(v string) *GetApplicationResponseBodyDataPublishedVersionScriptProfileFunctionMeta {
	s.FunctionId = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfileFunctionMeta) SetFunctionName(v string) *GetApplicationResponseBodyDataPublishedVersionScriptProfileFunctionMeta {
	s.FunctionName = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfileFunctionMeta) SetHttpTriggerName(v string) *GetApplicationResponseBodyDataPublishedVersionScriptProfileFunctionMeta {
	s.HttpTriggerName = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfileFunctionMeta) SetHttpTriggerUrl(v string) *GetApplicationResponseBodyDataPublishedVersionScriptProfileFunctionMeta {
	s.HttpTriggerUrl = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfileFunctionMeta) SetRegionId(v string) *GetApplicationResponseBodyDataPublishedVersionScriptProfileFunctionMeta {
	s.RegionId = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfileFunctionMeta) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyDataPublishedVersionScriptProfileNluAccessProfile struct {
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
}

func (s GetApplicationResponseBodyDataPublishedVersionScriptProfileNluAccessProfile) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataPublishedVersionScriptProfileNluAccessProfile) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfileNluAccessProfile) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfileNluAccessProfile) SetAccessProfileId(v string) *GetApplicationResponseBodyDataPublishedVersionScriptProfileNluAccessProfile {
	s.AccessProfileId = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfileNluAccessProfile) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig struct {
	Model            *string                                                                          `json:"Model,omitempty" xml:"Model,omitempty"`
	NlsAccessProfile *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfigNlsAccessProfile `json:"NlsAccessProfile,omitempty" xml:"NlsAccessProfile,omitempty" type:"Struct"`
	// The TTS invocation method.
	//
	// example:
	//
	// MANAGED
	NlsAccessType *string `json:"NlsAccessType,omitempty" xml:"NlsAccessType,omitempty"`
	// The TTS engine.
	//
	// example:
	//
	// ALIYUN
	NlsEngine *string `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
	// The pitch rate.
	//
	// > Valid values: -500 to 500.
	//
	// example:
	//
	// 3
	PitchRate *int32                                                                      `json:"PitchRate,omitempty" xml:"PitchRate,omitempty"`
	PronRules []*GetApplicationResponseBodyDataPublishedVersionSynthesizerConfigPronRules `json:"PronRules,omitempty" xml:"PronRules,omitempty" type:"Repeated"`
	// The speech rate.
	//
	// > Valid values: -500 to 500.
	//
	// example:
	//
	// -20
	SpeechRate *int32 `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	// The voice.
	//
	// example:
	//
	// aixia
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
	// The volume.
	//
	// example:
	//
	// 50
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) GetModel() *string {
	return s.Model
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) GetNlsAccessProfile() *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfigNlsAccessProfile {
	return s.NlsAccessProfile
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) GetNlsAccessType() *string {
	return s.NlsAccessType
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) GetPitchRate() *int32 {
	return s.PitchRate
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) GetPronRules() []*GetApplicationResponseBodyDataPublishedVersionSynthesizerConfigPronRules {
	return s.PronRules
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) GetSpeechRate() *int32 {
	return s.SpeechRate
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) GetVoice() *string {
	return s.Voice
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) GetVolume() *int32 {
	return s.Volume
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) SetModel(v string) *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig {
	s.Model = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) SetNlsAccessProfile(v *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfigNlsAccessProfile) *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig {
	s.NlsAccessProfile = v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) SetNlsAccessType(v string) *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig {
	s.NlsAccessType = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) SetNlsEngine(v string) *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig {
	s.NlsEngine = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) SetPitchRate(v int32) *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig {
	s.PitchRate = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) SetPronRules(v []*GetApplicationResponseBodyDataPublishedVersionSynthesizerConfigPronRules) *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig {
	s.PronRules = v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) SetSpeechRate(v int32) *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig {
	s.SpeechRate = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) SetVoice(v string) *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig {
	s.Voice = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) SetVolume(v int32) *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig {
	s.Volume = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) Validate() error {
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

type GetApplicationResponseBodyDataPublishedVersionSynthesizerConfigNlsAccessProfile struct {
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
}

func (s GetApplicationResponseBodyDataPublishedVersionSynthesizerConfigNlsAccessProfile) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataPublishedVersionSynthesizerConfigNlsAccessProfile) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfigNlsAccessProfile) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfigNlsAccessProfile) SetAccessProfileId(v string) *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfigNlsAccessProfile {
	s.AccessProfileId = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfigNlsAccessProfile) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyDataPublishedVersionSynthesizerConfigPronRules struct {
	Pattern     *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	Replacement *string `json:"Replacement,omitempty" xml:"Replacement,omitempty"`
}

func (s GetApplicationResponseBodyDataPublishedVersionSynthesizerConfigPronRules) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataPublishedVersionSynthesizerConfigPronRules) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfigPronRules) GetPattern() *string {
	return s.Pattern
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfigPronRules) GetReplacement() *string {
	return s.Replacement
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfigPronRules) SetPattern(v string) *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfigPronRules {
	s.Pattern = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfigPronRules) SetReplacement(v string) *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfigPronRules {
	s.Replacement = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfigPronRules) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyDataPublishedVersionToolConfig struct {
	// The list of MCP server configurations.
	McpServers []*GetApplicationResponseBodyDataPublishedVersionToolConfigMcpServers `json:"McpServers,omitempty" xml:"McpServers,omitempty" type:"Repeated"`
}

func (s GetApplicationResponseBodyDataPublishedVersionToolConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataPublishedVersionToolConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataPublishedVersionToolConfig) GetMcpServers() []*GetApplicationResponseBodyDataPublishedVersionToolConfigMcpServers {
	return s.McpServers
}

func (s *GetApplicationResponseBodyDataPublishedVersionToolConfig) SetMcpServers(v []*GetApplicationResponseBodyDataPublishedVersionToolConfigMcpServers) *GetApplicationResponseBodyDataPublishedVersionToolConfig {
	s.McpServers = v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionToolConfig) Validate() error {
	if s.McpServers != nil {
		for _, item := range s.McpServers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetApplicationResponseBodyDataPublishedVersionToolConfigMcpServers struct {
	// The base URL.
	//
	// example:
	//
	// https://example.com
	BaseUrl *string `json:"BaseUrl,omitempty" xml:"BaseUrl,omitempty"`
	// The name.
	//
	// example:
	//
	// phone-ai-call
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The SSE endpoint.
	//
	// example:
	//
	// /phone-ai-call/mcp/sse?key=value
	SseEndpoint *string `json:"SseEndpoint,omitempty" xml:"SseEndpoint,omitempty"`
}

func (s GetApplicationResponseBodyDataPublishedVersionToolConfigMcpServers) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataPublishedVersionToolConfigMcpServers) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataPublishedVersionToolConfigMcpServers) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *GetApplicationResponseBodyDataPublishedVersionToolConfigMcpServers) GetName() *string {
	return s.Name
}

func (s *GetApplicationResponseBodyDataPublishedVersionToolConfigMcpServers) GetSseEndpoint() *string {
	return s.SseEndpoint
}

func (s *GetApplicationResponseBodyDataPublishedVersionToolConfigMcpServers) SetBaseUrl(v string) *GetApplicationResponseBodyDataPublishedVersionToolConfigMcpServers {
	s.BaseUrl = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionToolConfigMcpServers) SetName(v string) *GetApplicationResponseBodyDataPublishedVersionToolConfigMcpServers {
	s.Name = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionToolConfigMcpServers) SetSseEndpoint(v string) *GetApplicationResponseBodyDataPublishedVersionToolConfigMcpServers {
	s.SseEndpoint = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionToolConfigMcpServers) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyDataPublishedVersionTranscriberConfig struct {
	CorrectionRules   []*GetApplicationResponseBodyDataPublishedVersionTranscriberConfigCorrectionRules `json:"CorrectionRules,omitempty" xml:"CorrectionRules,omitempty" type:"Repeated"`
	CustomizationId   *string                                                                           `json:"CustomizationId,omitempty" xml:"CustomizationId,omitempty"`
	EndSilenceTimeout *int32                                                                            `json:"EndSilenceTimeout,omitempty" xml:"EndSilenceTimeout,omitempty"`
	Model             *string                                                                           `json:"Model,omitempty" xml:"Model,omitempty"`
	NlsAccessProfile  *GetApplicationResponseBodyDataPublishedVersionTranscriberConfigNlsAccessProfile  `json:"NlsAccessProfile,omitempty" xml:"NlsAccessProfile,omitempty" type:"Struct"`
	// The ASR invocation method.
	//
	// example:
	//
	// MANAGED
	NlsAccessType *string `json:"NlsAccessType,omitempty" xml:"NlsAccessType,omitempty"`
	// The ASR engine.
	//
	// example:
	//
	// ALIYUN
	NlsEngine            *string `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
	SpeechNoiseThreshold *int32  `json:"SpeechNoiseThreshold,omitempty" xml:"SpeechNoiseThreshold,omitempty"`
	VocabularyId         *string `json:"VocabularyId,omitempty" xml:"VocabularyId,omitempty"`
}

func (s GetApplicationResponseBodyDataPublishedVersionTranscriberConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataPublishedVersionTranscriberConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig) GetCorrectionRules() []*GetApplicationResponseBodyDataPublishedVersionTranscriberConfigCorrectionRules {
	return s.CorrectionRules
}

func (s *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig) GetCustomizationId() *string {
	return s.CustomizationId
}

func (s *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig) GetEndSilenceTimeout() *int32 {
	return s.EndSilenceTimeout
}

func (s *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig) GetModel() *string {
	return s.Model
}

func (s *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig) GetNlsAccessProfile() *GetApplicationResponseBodyDataPublishedVersionTranscriberConfigNlsAccessProfile {
	return s.NlsAccessProfile
}

func (s *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig) GetNlsAccessType() *string {
	return s.NlsAccessType
}

func (s *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig) GetSpeechNoiseThreshold() *int32 {
	return s.SpeechNoiseThreshold
}

func (s *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig) GetVocabularyId() *string {
	return s.VocabularyId
}

func (s *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig) SetCorrectionRules(v []*GetApplicationResponseBodyDataPublishedVersionTranscriberConfigCorrectionRules) *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig {
	s.CorrectionRules = v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig) SetCustomizationId(v string) *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig {
	s.CustomizationId = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig) SetEndSilenceTimeout(v int32) *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig {
	s.EndSilenceTimeout = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig) SetModel(v string) *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig {
	s.Model = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig) SetNlsAccessProfile(v *GetApplicationResponseBodyDataPublishedVersionTranscriberConfigNlsAccessProfile) *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig {
	s.NlsAccessProfile = v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig) SetNlsAccessType(v string) *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig {
	s.NlsAccessType = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig) SetNlsEngine(v string) *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig {
	s.NlsEngine = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig) SetSpeechNoiseThreshold(v int32) *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig {
	s.SpeechNoiseThreshold = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig) SetVocabularyId(v string) *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig {
	s.VocabularyId = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig) Validate() error {
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

type GetApplicationResponseBodyDataPublishedVersionTranscriberConfigCorrectionRules struct {
	Pattern     *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	Replacement *string `json:"Replacement,omitempty" xml:"Replacement,omitempty"`
}

func (s GetApplicationResponseBodyDataPublishedVersionTranscriberConfigCorrectionRules) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataPublishedVersionTranscriberConfigCorrectionRules) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataPublishedVersionTranscriberConfigCorrectionRules) GetPattern() *string {
	return s.Pattern
}

func (s *GetApplicationResponseBodyDataPublishedVersionTranscriberConfigCorrectionRules) GetReplacement() *string {
	return s.Replacement
}

func (s *GetApplicationResponseBodyDataPublishedVersionTranscriberConfigCorrectionRules) SetPattern(v string) *GetApplicationResponseBodyDataPublishedVersionTranscriberConfigCorrectionRules {
	s.Pattern = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionTranscriberConfigCorrectionRules) SetReplacement(v string) *GetApplicationResponseBodyDataPublishedVersionTranscriberConfigCorrectionRules {
	s.Replacement = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionTranscriberConfigCorrectionRules) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyDataPublishedVersionTranscriberConfigNlsAccessProfile struct {
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
}

func (s GetApplicationResponseBodyDataPublishedVersionTranscriberConfigNlsAccessProfile) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataPublishedVersionTranscriberConfigNlsAccessProfile) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataPublishedVersionTranscriberConfigNlsAccessProfile) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *GetApplicationResponseBodyDataPublishedVersionTranscriberConfigNlsAccessProfile) SetAccessProfileId(v string) *GetApplicationResponseBodyDataPublishedVersionTranscriberConfigNlsAccessProfile {
	s.AccessProfileId = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionTranscriberConfigNlsAccessProfile) Validate() error {
	return dara.Validate(s)
}
