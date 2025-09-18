// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMmAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *CreateMmAppRequest
	GetAppName() *string
	SetBindingConfig(v *CreateMmAppRequestBindingConfig) *CreateMmAppRequest
	GetBindingConfig() *CreateMmAppRequestBindingConfig
	SetConversationConfig(v *CreateMmAppRequestConversationConfig) *CreateMmAppRequest
	GetConversationConfig() *CreateMmAppRequestConversationConfig
	SetModelConfig(v *CreateMmAppRequestModelConfig) *CreateMmAppRequest
	GetModelConfig() *CreateMmAppRequestModelConfig
	SetPrompt(v string) *CreateMmAppRequest
	GetPrompt() *string
	SetWorkspaceId(v string) *CreateMmAppRequest
	GetWorkspaceId() *string
}

type CreateMmAppRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 多模态xxx
	AppName            *string                               `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BindingConfig      *CreateMmAppRequestBindingConfig      `json:"BindingConfig,omitempty" xml:"BindingConfig,omitempty" type:"Struct"`
	ConversationConfig *CreateMmAppRequestConversationConfig `json:"ConversationConfig,omitempty" xml:"ConversationConfig,omitempty" type:"Struct"`
	ModelConfig        *CreateMmAppRequestModelConfig        `json:"ModelConfig,omitempty" xml:"ModelConfig,omitempty" type:"Struct"`
	// example:
	//
	// 提示词
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateMmAppRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMmAppRequest) GoString() string {
	return s.String()
}

func (s *CreateMmAppRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateMmAppRequest) GetBindingConfig() *CreateMmAppRequestBindingConfig {
	return s.BindingConfig
}

func (s *CreateMmAppRequest) GetConversationConfig() *CreateMmAppRequestConversationConfig {
	return s.ConversationConfig
}

func (s *CreateMmAppRequest) GetModelConfig() *CreateMmAppRequestModelConfig {
	return s.ModelConfig
}

func (s *CreateMmAppRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *CreateMmAppRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateMmAppRequest) SetAppName(v string) *CreateMmAppRequest {
	s.AppName = &v
	return s
}

func (s *CreateMmAppRequest) SetBindingConfig(v *CreateMmAppRequestBindingConfig) *CreateMmAppRequest {
	s.BindingConfig = v
	return s
}

func (s *CreateMmAppRequest) SetConversationConfig(v *CreateMmAppRequestConversationConfig) *CreateMmAppRequest {
	s.ConversationConfig = v
	return s
}

func (s *CreateMmAppRequest) SetModelConfig(v *CreateMmAppRequestModelConfig) *CreateMmAppRequest {
	s.ModelConfig = v
	return s
}

func (s *CreateMmAppRequest) SetPrompt(v string) *CreateMmAppRequest {
	s.Prompt = &v
	return s
}

func (s *CreateMmAppRequest) SetWorkspaceId(v string) *CreateMmAppRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateMmAppRequest) Validate() error {
	return dara.Validate(s)
}

type CreateMmAppRequestBindingConfig struct {
	Commands []*CreateMmAppRequestBindingConfigCommands `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
}

func (s CreateMmAppRequestBindingConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateMmAppRequestBindingConfig) GoString() string {
	return s.String()
}

func (s *CreateMmAppRequestBindingConfig) GetCommands() []*CreateMmAppRequestBindingConfigCommands {
	return s.Commands
}

func (s *CreateMmAppRequestBindingConfig) SetCommands(v []*CreateMmAppRequestBindingConfigCommands) *CreateMmAppRequestBindingConfig {
	s.Commands = v
	return s
}

func (s *CreateMmAppRequestBindingConfig) Validate() error {
	return dara.Validate(s)
}

type CreateMmAppRequestBindingConfigCommands struct {
	// This parameter is required.
	//
	// example:
	//
	// 3686786786
	DomainCode *string                                         `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	Tools      []*CreateMmAppRequestBindingConfigCommandsTools `json:"Tools,omitempty" xml:"Tools,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// CUSTOM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateMmAppRequestBindingConfigCommands) String() string {
	return dara.Prettify(s)
}

func (s CreateMmAppRequestBindingConfigCommands) GoString() string {
	return s.String()
}

func (s *CreateMmAppRequestBindingConfigCommands) GetDomainCode() *string {
	return s.DomainCode
}

func (s *CreateMmAppRequestBindingConfigCommands) GetTools() []*CreateMmAppRequestBindingConfigCommandsTools {
	return s.Tools
}

func (s *CreateMmAppRequestBindingConfigCommands) GetType() *string {
	return s.Type
}

func (s *CreateMmAppRequestBindingConfigCommands) SetDomainCode(v string) *CreateMmAppRequestBindingConfigCommands {
	s.DomainCode = &v
	return s
}

func (s *CreateMmAppRequestBindingConfigCommands) SetTools(v []*CreateMmAppRequestBindingConfigCommandsTools) *CreateMmAppRequestBindingConfigCommands {
	s.Tools = v
	return s
}

func (s *CreateMmAppRequestBindingConfigCommands) SetType(v string) *CreateMmAppRequestBindingConfigCommands {
	s.Type = &v
	return s
}

func (s *CreateMmAppRequestBindingConfigCommands) Validate() error {
	return dara.Validate(s)
}

type CreateMmAppRequestBindingConfigCommandsTools struct {
	// This parameter is required.
	//
	// example:
	//
	// 54645646
	ToolId *string `json:"ToolId,omitempty" xml:"ToolId,omitempty"`
}

func (s CreateMmAppRequestBindingConfigCommandsTools) String() string {
	return dara.Prettify(s)
}

func (s CreateMmAppRequestBindingConfigCommandsTools) GoString() string {
	return s.String()
}

func (s *CreateMmAppRequestBindingConfigCommandsTools) GetToolId() *string {
	return s.ToolId
}

func (s *CreateMmAppRequestBindingConfigCommandsTools) SetToolId(v string) *CreateMmAppRequestBindingConfigCommandsTools {
	s.ToolId = &v
	return s
}

func (s *CreateMmAppRequestBindingConfigCommandsTools) Validate() error {
	return dara.Validate(s)
}

type CreateMmAppRequestConversationConfig struct {
	// example:
	//
	// xxx
	AsrModel *string `json:"AsrModel,omitempty" xml:"AsrModel,omitempty"`
	OpenAsr  *bool   `json:"OpenAsr,omitempty" xml:"OpenAsr,omitempty"`
	OpenTts  *bool   `json:"OpenTts,omitempty" xml:"OpenTts,omitempty"`
	// example:
	//
	// xxx
	TtsModel *string `json:"TtsModel,omitempty" xml:"TtsModel,omitempty"`
}

func (s CreateMmAppRequestConversationConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateMmAppRequestConversationConfig) GoString() string {
	return s.String()
}

func (s *CreateMmAppRequestConversationConfig) GetAsrModel() *string {
	return s.AsrModel
}

func (s *CreateMmAppRequestConversationConfig) GetOpenAsr() *bool {
	return s.OpenAsr
}

func (s *CreateMmAppRequestConversationConfig) GetOpenTts() *bool {
	return s.OpenTts
}

func (s *CreateMmAppRequestConversationConfig) GetTtsModel() *string {
	return s.TtsModel
}

func (s *CreateMmAppRequestConversationConfig) SetAsrModel(v string) *CreateMmAppRequestConversationConfig {
	s.AsrModel = &v
	return s
}

func (s *CreateMmAppRequestConversationConfig) SetOpenAsr(v bool) *CreateMmAppRequestConversationConfig {
	s.OpenAsr = &v
	return s
}

func (s *CreateMmAppRequestConversationConfig) SetOpenTts(v bool) *CreateMmAppRequestConversationConfig {
	s.OpenTts = &v
	return s
}

func (s *CreateMmAppRequestConversationConfig) SetTtsModel(v string) *CreateMmAppRequestConversationConfig {
	s.TtsModel = &v
	return s
}

func (s *CreateMmAppRequestConversationConfig) Validate() error {
	return dara.Validate(s)
}

type CreateMmAppRequestModelConfig struct {
	// example:
	//
	// 5
	HistoryLimit *int32 `json:"HistoryLimit,omitempty" xml:"HistoryLimit,omitempty"`
	// example:
	//
	// MMH
	ModelType     *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	OpenWebSearch *bool   `json:"OpenWebSearch,omitempty" xml:"OpenWebSearch,omitempty"`
	// example:
	//
	// xxxx
	TextModal *string `json:"TextModal,omitempty" xml:"TextModal,omitempty"`
}

func (s CreateMmAppRequestModelConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateMmAppRequestModelConfig) GoString() string {
	return s.String()
}

func (s *CreateMmAppRequestModelConfig) GetHistoryLimit() *int32 {
	return s.HistoryLimit
}

func (s *CreateMmAppRequestModelConfig) GetModelType() *string {
	return s.ModelType
}

func (s *CreateMmAppRequestModelConfig) GetOpenWebSearch() *bool {
	return s.OpenWebSearch
}

func (s *CreateMmAppRequestModelConfig) GetTextModal() *string {
	return s.TextModal
}

func (s *CreateMmAppRequestModelConfig) SetHistoryLimit(v int32) *CreateMmAppRequestModelConfig {
	s.HistoryLimit = &v
	return s
}

func (s *CreateMmAppRequestModelConfig) SetModelType(v string) *CreateMmAppRequestModelConfig {
	s.ModelType = &v
	return s
}

func (s *CreateMmAppRequestModelConfig) SetOpenWebSearch(v bool) *CreateMmAppRequestModelConfig {
	s.OpenWebSearch = &v
	return s
}

func (s *CreateMmAppRequestModelConfig) SetTextModal(v string) *CreateMmAppRequestModelConfig {
	s.TextModal = &v
	return s
}

func (s *CreateMmAppRequestModelConfig) Validate() error {
	return dara.Validate(s)
}
