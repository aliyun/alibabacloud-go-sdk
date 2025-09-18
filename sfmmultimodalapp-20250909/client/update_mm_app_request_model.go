// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateMmAppRequest
	GetAppId() *string
	SetAppName(v string) *UpdateMmAppRequest
	GetAppName() *string
	SetBindingConfig(v *UpdateMmAppRequestBindingConfig) *UpdateMmAppRequest
	GetBindingConfig() *UpdateMmAppRequestBindingConfig
	SetConversationConfig(v *UpdateMmAppRequestConversationConfig) *UpdateMmAppRequest
	GetConversationConfig() *UpdateMmAppRequestConversationConfig
	SetModelConfig(v *UpdateMmAppRequestModelConfig) *UpdateMmAppRequest
	GetModelConfig() *UpdateMmAppRequestModelConfig
	SetPrompt(v string) *UpdateMmAppRequest
	GetPrompt() *string
	SetWorkspaceId(v string) *UpdateMmAppRequest
	GetWorkspaceId() *string
}

type UpdateMmAppRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mm_xxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 多模态应用xxxxx
	AppName            *string                               `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BindingConfig      *UpdateMmAppRequestBindingConfig      `json:"BindingConfig,omitempty" xml:"BindingConfig,omitempty" type:"Struct"`
	ConversationConfig *UpdateMmAppRequestConversationConfig `json:"ConversationConfig,omitempty" xml:"ConversationConfig,omitempty" type:"Struct"`
	ModelConfig        *UpdateMmAppRequestModelConfig        `json:"ModelConfig,omitempty" xml:"ModelConfig,omitempty" type:"Struct"`
	// example:
	//
	// 提示词，不超过8000字符
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// llm-xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateMmAppRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppRequest) GoString() string {
	return s.String()
}

func (s *UpdateMmAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateMmAppRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateMmAppRequest) GetBindingConfig() *UpdateMmAppRequestBindingConfig {
	return s.BindingConfig
}

func (s *UpdateMmAppRequest) GetConversationConfig() *UpdateMmAppRequestConversationConfig {
	return s.ConversationConfig
}

func (s *UpdateMmAppRequest) GetModelConfig() *UpdateMmAppRequestModelConfig {
	return s.ModelConfig
}

func (s *UpdateMmAppRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *UpdateMmAppRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateMmAppRequest) SetAppId(v string) *UpdateMmAppRequest {
	s.AppId = &v
	return s
}

func (s *UpdateMmAppRequest) SetAppName(v string) *UpdateMmAppRequest {
	s.AppName = &v
	return s
}

func (s *UpdateMmAppRequest) SetBindingConfig(v *UpdateMmAppRequestBindingConfig) *UpdateMmAppRequest {
	s.BindingConfig = v
	return s
}

func (s *UpdateMmAppRequest) SetConversationConfig(v *UpdateMmAppRequestConversationConfig) *UpdateMmAppRequest {
	s.ConversationConfig = v
	return s
}

func (s *UpdateMmAppRequest) SetModelConfig(v *UpdateMmAppRequestModelConfig) *UpdateMmAppRequest {
	s.ModelConfig = v
	return s
}

func (s *UpdateMmAppRequest) SetPrompt(v string) *UpdateMmAppRequest {
	s.Prompt = &v
	return s
}

func (s *UpdateMmAppRequest) SetWorkspaceId(v string) *UpdateMmAppRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateMmAppRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateMmAppRequestBindingConfig struct {
	Commands []*UpdateMmAppRequestBindingConfigCommands `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
}

func (s UpdateMmAppRequestBindingConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppRequestBindingConfig) GoString() string {
	return s.String()
}

func (s *UpdateMmAppRequestBindingConfig) GetCommands() []*UpdateMmAppRequestBindingConfigCommands {
	return s.Commands
}

func (s *UpdateMmAppRequestBindingConfig) SetCommands(v []*UpdateMmAppRequestBindingConfigCommands) *UpdateMmAppRequestBindingConfig {
	s.Commands = v
	return s
}

func (s *UpdateMmAppRequestBindingConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateMmAppRequestBindingConfigCommands struct {
	// This parameter is required.
	//
	// example:
	//
	// 724366858658
	DomainCode *string                                         `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	Tools      []*UpdateMmAppRequestBindingConfigCommandsTools `json:"Tools,omitempty" xml:"Tools,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// CUSTOM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateMmAppRequestBindingConfigCommands) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppRequestBindingConfigCommands) GoString() string {
	return s.String()
}

func (s *UpdateMmAppRequestBindingConfigCommands) GetDomainCode() *string {
	return s.DomainCode
}

func (s *UpdateMmAppRequestBindingConfigCommands) GetTools() []*UpdateMmAppRequestBindingConfigCommandsTools {
	return s.Tools
}

func (s *UpdateMmAppRequestBindingConfigCommands) GetType() *string {
	return s.Type
}

func (s *UpdateMmAppRequestBindingConfigCommands) SetDomainCode(v string) *UpdateMmAppRequestBindingConfigCommands {
	s.DomainCode = &v
	return s
}

func (s *UpdateMmAppRequestBindingConfigCommands) SetTools(v []*UpdateMmAppRequestBindingConfigCommandsTools) *UpdateMmAppRequestBindingConfigCommands {
	s.Tools = v
	return s
}

func (s *UpdateMmAppRequestBindingConfigCommands) SetType(v string) *UpdateMmAppRequestBindingConfigCommands {
	s.Type = &v
	return s
}

func (s *UpdateMmAppRequestBindingConfigCommands) Validate() error {
	return dara.Validate(s)
}

type UpdateMmAppRequestBindingConfigCommandsTools struct {
	// This parameter is required.
	//
	// example:
	//
	// 7293782043943
	ToolId *string `json:"ToolId,omitempty" xml:"ToolId,omitempty"`
}

func (s UpdateMmAppRequestBindingConfigCommandsTools) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppRequestBindingConfigCommandsTools) GoString() string {
	return s.String()
}

func (s *UpdateMmAppRequestBindingConfigCommandsTools) GetToolId() *string {
	return s.ToolId
}

func (s *UpdateMmAppRequestBindingConfigCommandsTools) SetToolId(v string) *UpdateMmAppRequestBindingConfigCommandsTools {
	s.ToolId = &v
	return s
}

func (s *UpdateMmAppRequestBindingConfigCommandsTools) Validate() error {
	return dara.Validate(s)
}

type UpdateMmAppRequestConversationConfig struct {
	// example:
	//
	// Paraformer
	AsrModel *string `json:"AsrModel,omitempty" xml:"AsrModel,omitempty"`
	OpenAsr  *bool   `json:"OpenAsr,omitempty" xml:"OpenAsr,omitempty"`
	OpenTts  *bool   `json:"OpenTts,omitempty" xml:"OpenTts,omitempty"`
	// example:
	//
	// cosyvoice-v2
	TtsModel *string `json:"TtsModel,omitempty" xml:"TtsModel,omitempty"`
}

func (s UpdateMmAppRequestConversationConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppRequestConversationConfig) GoString() string {
	return s.String()
}

func (s *UpdateMmAppRequestConversationConfig) GetAsrModel() *string {
	return s.AsrModel
}

func (s *UpdateMmAppRequestConversationConfig) GetOpenAsr() *bool {
	return s.OpenAsr
}

func (s *UpdateMmAppRequestConversationConfig) GetOpenTts() *bool {
	return s.OpenTts
}

func (s *UpdateMmAppRequestConversationConfig) GetTtsModel() *string {
	return s.TtsModel
}

func (s *UpdateMmAppRequestConversationConfig) SetAsrModel(v string) *UpdateMmAppRequestConversationConfig {
	s.AsrModel = &v
	return s
}

func (s *UpdateMmAppRequestConversationConfig) SetOpenAsr(v bool) *UpdateMmAppRequestConversationConfig {
	s.OpenAsr = &v
	return s
}

func (s *UpdateMmAppRequestConversationConfig) SetOpenTts(v bool) *UpdateMmAppRequestConversationConfig {
	s.OpenTts = &v
	return s
}

func (s *UpdateMmAppRequestConversationConfig) SetTtsModel(v string) *UpdateMmAppRequestConversationConfig {
	s.TtsModel = &v
	return s
}

func (s *UpdateMmAppRequestConversationConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateMmAppRequestModelConfig struct {
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
	// qwen-mmh-high-speed
	TextModal *string `json:"TextModal,omitempty" xml:"TextModal,omitempty"`
}

func (s UpdateMmAppRequestModelConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppRequestModelConfig) GoString() string {
	return s.String()
}

func (s *UpdateMmAppRequestModelConfig) GetHistoryLimit() *int32 {
	return s.HistoryLimit
}

func (s *UpdateMmAppRequestModelConfig) GetModelType() *string {
	return s.ModelType
}

func (s *UpdateMmAppRequestModelConfig) GetOpenWebSearch() *bool {
	return s.OpenWebSearch
}

func (s *UpdateMmAppRequestModelConfig) GetTextModal() *string {
	return s.TextModal
}

func (s *UpdateMmAppRequestModelConfig) SetHistoryLimit(v int32) *UpdateMmAppRequestModelConfig {
	s.HistoryLimit = &v
	return s
}

func (s *UpdateMmAppRequestModelConfig) SetModelType(v string) *UpdateMmAppRequestModelConfig {
	s.ModelType = &v
	return s
}

func (s *UpdateMmAppRequestModelConfig) SetOpenWebSearch(v bool) *UpdateMmAppRequestModelConfig {
	s.OpenWebSearch = &v
	return s
}

func (s *UpdateMmAppRequestModelConfig) SetTextModal(v string) *UpdateMmAppRequestModelConfig {
	s.TextModal = &v
	return s
}

func (s *UpdateMmAppRequestModelConfig) Validate() error {
	return dara.Validate(s)
}
