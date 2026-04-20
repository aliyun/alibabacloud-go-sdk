// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMmAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppConfig(v *DescribeMmAppResponseBodyAppConfig) *DescribeMmAppResponseBody
	GetAppConfig() *DescribeMmAppResponseBodyAppConfig
	SetAppId(v string) *DescribeMmAppResponseBody
	GetAppId() *string
	SetAppName(v string) *DescribeMmAppResponseBody
	GetAppName() *string
	SetBindingConfig(v *DescribeMmAppResponseBodyBindingConfig) *DescribeMmAppResponseBody
	GetBindingConfig() *DescribeMmAppResponseBodyBindingConfig
	SetConversationConfig(v *DescribeMmAppResponseBodyConversationConfig) *DescribeMmAppResponseBody
	GetConversationConfig() *DescribeMmAppResponseBodyConversationConfig
	SetCreateUserId(v string) *DescribeMmAppResponseBody
	GetCreateUserId() *string
	SetCreateUserName(v string) *DescribeMmAppResponseBody
	GetCreateUserName() *string
	SetGmtCreate(v string) *DescribeMmAppResponseBody
	GetGmtCreate() *string
	SetGmtModified(v string) *DescribeMmAppResponseBody
	GetGmtModified() *string
	SetModelConfig(v *DescribeMmAppResponseBodyModelConfig) *DescribeMmAppResponseBody
	GetModelConfig() *DescribeMmAppResponseBodyModelConfig
	SetModifyUserId(v string) *DescribeMmAppResponseBody
	GetModifyUserId() *string
	SetModifyUserName(v string) *DescribeMmAppResponseBody
	GetModifyUserName() *string
	SetPrompt(v string) *DescribeMmAppResponseBody
	GetPrompt() *string
	SetPublishVersion(v int64) *DescribeMmAppResponseBody
	GetPublishVersion() *int64
	SetRequestId(v string) *DescribeMmAppResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeMmAppResponseBody
	GetStatus() *string
}

type DescribeMmAppResponseBody struct {
	AppConfig *DescribeMmAppResponseBodyAppConfig `json:"AppConfig,omitempty" xml:"AppConfig,omitempty" type:"Struct"`
	// example:
	//
	// mm_xxxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 多模态应用xxxx
	AppName            *string                                      `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BindingConfig      *DescribeMmAppResponseBodyBindingConfig      `json:"BindingConfig,omitempty" xml:"BindingConfig,omitempty" type:"Struct"`
	ConversationConfig *DescribeMmAppResponseBodyConversationConfig `json:"ConversationConfig,omitempty" xml:"ConversationConfig,omitempty" type:"Struct"`
	// example:
	//
	// 243433
	CreateUserId *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// example:
	//
	// xxxx
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	// example:
	//
	// xxx
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// xxx
	GmtModified *string                               `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	ModelConfig *DescribeMmAppResponseBodyModelConfig `json:"ModelConfig,omitempty" xml:"ModelConfig,omitempty" type:"Struct"`
	// example:
	//
	// 56673435
	ModifyUserId *string `json:"ModifyUserId,omitempty" xml:"ModifyUserId,omitempty"`
	// example:
	//
	// xxxx
	ModifyUserName *string `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	// example:
	//
	// 提示词xxxx
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// 1
	PublishVersion *int64 `json:"PublishVersion,omitempty" xml:"PublishVersion,omitempty"`
	// Id of the request
	//
	// example:
	//
	// xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeMmAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMmAppResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMmAppResponseBody) GetAppConfig() *DescribeMmAppResponseBodyAppConfig {
	return s.AppConfig
}

func (s *DescribeMmAppResponseBody) GetAppId() *string {
	return s.AppId
}

func (s *DescribeMmAppResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *DescribeMmAppResponseBody) GetBindingConfig() *DescribeMmAppResponseBodyBindingConfig {
	return s.BindingConfig
}

func (s *DescribeMmAppResponseBody) GetConversationConfig() *DescribeMmAppResponseBodyConversationConfig {
	return s.ConversationConfig
}

func (s *DescribeMmAppResponseBody) GetCreateUserId() *string {
	return s.CreateUserId
}

func (s *DescribeMmAppResponseBody) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *DescribeMmAppResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeMmAppResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeMmAppResponseBody) GetModelConfig() *DescribeMmAppResponseBodyModelConfig {
	return s.ModelConfig
}

func (s *DescribeMmAppResponseBody) GetModifyUserId() *string {
	return s.ModifyUserId
}

func (s *DescribeMmAppResponseBody) GetModifyUserName() *string {
	return s.ModifyUserName
}

func (s *DescribeMmAppResponseBody) GetPrompt() *string {
	return s.Prompt
}

func (s *DescribeMmAppResponseBody) GetPublishVersion() *int64 {
	return s.PublishVersion
}

func (s *DescribeMmAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMmAppResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeMmAppResponseBody) SetAppConfig(v *DescribeMmAppResponseBodyAppConfig) *DescribeMmAppResponseBody {
	s.AppConfig = v
	return s
}

func (s *DescribeMmAppResponseBody) SetAppId(v string) *DescribeMmAppResponseBody {
	s.AppId = &v
	return s
}

func (s *DescribeMmAppResponseBody) SetAppName(v string) *DescribeMmAppResponseBody {
	s.AppName = &v
	return s
}

func (s *DescribeMmAppResponseBody) SetBindingConfig(v *DescribeMmAppResponseBodyBindingConfig) *DescribeMmAppResponseBody {
	s.BindingConfig = v
	return s
}

func (s *DescribeMmAppResponseBody) SetConversationConfig(v *DescribeMmAppResponseBodyConversationConfig) *DescribeMmAppResponseBody {
	s.ConversationConfig = v
	return s
}

func (s *DescribeMmAppResponseBody) SetCreateUserId(v string) *DescribeMmAppResponseBody {
	s.CreateUserId = &v
	return s
}

func (s *DescribeMmAppResponseBody) SetCreateUserName(v string) *DescribeMmAppResponseBody {
	s.CreateUserName = &v
	return s
}

func (s *DescribeMmAppResponseBody) SetGmtCreate(v string) *DescribeMmAppResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *DescribeMmAppResponseBody) SetGmtModified(v string) *DescribeMmAppResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DescribeMmAppResponseBody) SetModelConfig(v *DescribeMmAppResponseBodyModelConfig) *DescribeMmAppResponseBody {
	s.ModelConfig = v
	return s
}

func (s *DescribeMmAppResponseBody) SetModifyUserId(v string) *DescribeMmAppResponseBody {
	s.ModifyUserId = &v
	return s
}

func (s *DescribeMmAppResponseBody) SetModifyUserName(v string) *DescribeMmAppResponseBody {
	s.ModifyUserName = &v
	return s
}

func (s *DescribeMmAppResponseBody) SetPrompt(v string) *DescribeMmAppResponseBody {
	s.Prompt = &v
	return s
}

func (s *DescribeMmAppResponseBody) SetPublishVersion(v int64) *DescribeMmAppResponseBody {
	s.PublishVersion = &v
	return s
}

func (s *DescribeMmAppResponseBody) SetRequestId(v string) *DescribeMmAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMmAppResponseBody) SetStatus(v string) *DescribeMmAppResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeMmAppResponseBody) Validate() error {
	if s.AppConfig != nil {
		if err := s.AppConfig.Validate(); err != nil {
			return err
		}
	}
	if s.BindingConfig != nil {
		if err := s.BindingConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ConversationConfig != nil {
		if err := s.ConversationConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ModelConfig != nil {
		if err := s.ModelConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMmAppResponseBodyAppConfig struct {
	// example:
	//
	// true
	EnableTransition *bool `json:"EnableTransition,omitempty" xml:"EnableTransition,omitempty"`
}

func (s DescribeMmAppResponseBodyAppConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeMmAppResponseBodyAppConfig) GoString() string {
	return s.String()
}

func (s *DescribeMmAppResponseBodyAppConfig) GetEnableTransition() *bool {
	return s.EnableTransition
}

func (s *DescribeMmAppResponseBodyAppConfig) SetEnableTransition(v bool) *DescribeMmAppResponseBodyAppConfig {
	s.EnableTransition = &v
	return s
}

func (s *DescribeMmAppResponseBodyAppConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeMmAppResponseBodyBindingConfig struct {
	Commands  []*DescribeMmAppResponseBodyBindingConfigCommands `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
	Mcps      []*DescribeMmAppResponseBodyBindingConfigMcps     `json:"Mcps,omitempty" xml:"Mcps,omitempty" type:"Repeated"`
	RagConfig *DescribeMmAppResponseBodyBindingConfigRagConfig  `json:"RagConfig,omitempty" xml:"RagConfig,omitempty" type:"Struct"`
}

func (s DescribeMmAppResponseBodyBindingConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeMmAppResponseBodyBindingConfig) GoString() string {
	return s.String()
}

func (s *DescribeMmAppResponseBodyBindingConfig) GetCommands() []*DescribeMmAppResponseBodyBindingConfigCommands {
	return s.Commands
}

func (s *DescribeMmAppResponseBodyBindingConfig) GetMcps() []*DescribeMmAppResponseBodyBindingConfigMcps {
	return s.Mcps
}

func (s *DescribeMmAppResponseBodyBindingConfig) GetRagConfig() *DescribeMmAppResponseBodyBindingConfigRagConfig {
	return s.RagConfig
}

func (s *DescribeMmAppResponseBodyBindingConfig) SetCommands(v []*DescribeMmAppResponseBodyBindingConfigCommands) *DescribeMmAppResponseBodyBindingConfig {
	s.Commands = v
	return s
}

func (s *DescribeMmAppResponseBodyBindingConfig) SetMcps(v []*DescribeMmAppResponseBodyBindingConfigMcps) *DescribeMmAppResponseBodyBindingConfig {
	s.Mcps = v
	return s
}

func (s *DescribeMmAppResponseBodyBindingConfig) SetRagConfig(v *DescribeMmAppResponseBodyBindingConfigRagConfig) *DescribeMmAppResponseBodyBindingConfig {
	s.RagConfig = v
	return s
}

func (s *DescribeMmAppResponseBodyBindingConfig) Validate() error {
	if s.Commands != nil {
		for _, item := range s.Commands {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Mcps != nil {
		for _, item := range s.Mcps {
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

type DescribeMmAppResponseBodyBindingConfigCommands struct {
	// example:
	//
	// xxx
	DomainCode *string   `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	Tools      []*string `json:"Tools,omitempty" xml:"Tools,omitempty" type:"Repeated"`
	// example:
	//
	// BAILIAN
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeMmAppResponseBodyBindingConfigCommands) String() string {
	return dara.Prettify(s)
}

func (s DescribeMmAppResponseBodyBindingConfigCommands) GoString() string {
	return s.String()
}

func (s *DescribeMmAppResponseBodyBindingConfigCommands) GetDomainCode() *string {
	return s.DomainCode
}

func (s *DescribeMmAppResponseBodyBindingConfigCommands) GetTools() []*string {
	return s.Tools
}

func (s *DescribeMmAppResponseBodyBindingConfigCommands) GetType() *string {
	return s.Type
}

func (s *DescribeMmAppResponseBodyBindingConfigCommands) SetDomainCode(v string) *DescribeMmAppResponseBodyBindingConfigCommands {
	s.DomainCode = &v
	return s
}

func (s *DescribeMmAppResponseBodyBindingConfigCommands) SetTools(v []*string) *DescribeMmAppResponseBodyBindingConfigCommands {
	s.Tools = v
	return s
}

func (s *DescribeMmAppResponseBodyBindingConfigCommands) SetType(v string) *DescribeMmAppResponseBodyBindingConfigCommands {
	s.Type = &v
	return s
}

func (s *DescribeMmAppResponseBodyBindingConfigCommands) Validate() error {
	return dara.Validate(s)
}

type DescribeMmAppResponseBodyBindingConfigMcps struct {
	// example:
	//
	// mcp-xxxx
	Code     *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	ToolList []*string `json:"ToolList,omitempty" xml:"ToolList,omitempty" type:"Repeated"`
}

func (s DescribeMmAppResponseBodyBindingConfigMcps) String() string {
	return dara.Prettify(s)
}

func (s DescribeMmAppResponseBodyBindingConfigMcps) GoString() string {
	return s.String()
}

func (s *DescribeMmAppResponseBodyBindingConfigMcps) GetCode() *string {
	return s.Code
}

func (s *DescribeMmAppResponseBodyBindingConfigMcps) GetToolList() []*string {
	return s.ToolList
}

func (s *DescribeMmAppResponseBodyBindingConfigMcps) SetCode(v string) *DescribeMmAppResponseBodyBindingConfigMcps {
	s.Code = &v
	return s
}

func (s *DescribeMmAppResponseBodyBindingConfigMcps) SetToolList(v []*string) *DescribeMmAppResponseBodyBindingConfigMcps {
	s.ToolList = v
	return s
}

func (s *DescribeMmAppResponseBodyBindingConfigMcps) Validate() error {
	return dara.Validate(s)
}

type DescribeMmAppResponseBodyBindingConfigRagConfig struct {
	// example:
	//
	// true
	EnableSearch          *string   `json:"EnableSearch,omitempty" xml:"EnableSearch,omitempty"`
	KnowledgeBaseCodeList []*string `json:"KnowledgeBaseCodeList,omitempty" xml:"KnowledgeBaseCodeList,omitempty" type:"Repeated"`
	// example:
	//
	// top_k
	PromptStrategy *string             `json:"PromptStrategy,omitempty" xml:"PromptStrategy,omitempty"`
	RankWeights    map[string]*float64 `json:"RankWeights,omitempty" xml:"RankWeights,omitempty"`
	// example:
	//
	// 1000
	RetrieveMaxLength *int32 `json:"RetrieveMaxLength,omitempty" xml:"RetrieveMaxLength,omitempty"`
	// example:
	//
	// 5
	TopK *int32 `json:"TopK,omitempty" xml:"TopK,omitempty"`
}

func (s DescribeMmAppResponseBodyBindingConfigRagConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeMmAppResponseBodyBindingConfigRagConfig) GoString() string {
	return s.String()
}

func (s *DescribeMmAppResponseBodyBindingConfigRagConfig) GetEnableSearch() *string {
	return s.EnableSearch
}

func (s *DescribeMmAppResponseBodyBindingConfigRagConfig) GetKnowledgeBaseCodeList() []*string {
	return s.KnowledgeBaseCodeList
}

func (s *DescribeMmAppResponseBodyBindingConfigRagConfig) GetPromptStrategy() *string {
	return s.PromptStrategy
}

func (s *DescribeMmAppResponseBodyBindingConfigRagConfig) GetRankWeights() map[string]*float64 {
	return s.RankWeights
}

func (s *DescribeMmAppResponseBodyBindingConfigRagConfig) GetRetrieveMaxLength() *int32 {
	return s.RetrieveMaxLength
}

func (s *DescribeMmAppResponseBodyBindingConfigRagConfig) GetTopK() *int32 {
	return s.TopK
}

func (s *DescribeMmAppResponseBodyBindingConfigRagConfig) SetEnableSearch(v string) *DescribeMmAppResponseBodyBindingConfigRagConfig {
	s.EnableSearch = &v
	return s
}

func (s *DescribeMmAppResponseBodyBindingConfigRagConfig) SetKnowledgeBaseCodeList(v []*string) *DescribeMmAppResponseBodyBindingConfigRagConfig {
	s.KnowledgeBaseCodeList = v
	return s
}

func (s *DescribeMmAppResponseBodyBindingConfigRagConfig) SetPromptStrategy(v string) *DescribeMmAppResponseBodyBindingConfigRagConfig {
	s.PromptStrategy = &v
	return s
}

func (s *DescribeMmAppResponseBodyBindingConfigRagConfig) SetRankWeights(v map[string]*float64) *DescribeMmAppResponseBodyBindingConfigRagConfig {
	s.RankWeights = v
	return s
}

func (s *DescribeMmAppResponseBodyBindingConfigRagConfig) SetRetrieveMaxLength(v int32) *DescribeMmAppResponseBodyBindingConfigRagConfig {
	s.RetrieveMaxLength = &v
	return s
}

func (s *DescribeMmAppResponseBodyBindingConfigRagConfig) SetTopK(v int32) *DescribeMmAppResponseBodyBindingConfigRagConfig {
	s.TopK = &v
	return s
}

func (s *DescribeMmAppResponseBodyBindingConfigRagConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeMmAppResponseBodyConversationConfig struct {
	// example:
	//
	// xxxx
	AsrModel *string `json:"AsrModel,omitempty" xml:"AsrModel,omitempty"`
	OpenAsr  *bool   `json:"OpenAsr,omitempty" xml:"OpenAsr,omitempty"`
	OpenTts  *bool   `json:"OpenTts,omitempty" xml:"OpenTts,omitempty"`
	// example:
	//
	// xxxx
	TtsModel *string `json:"TtsModel,omitempty" xml:"TtsModel,omitempty"`
}

func (s DescribeMmAppResponseBodyConversationConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeMmAppResponseBodyConversationConfig) GoString() string {
	return s.String()
}

func (s *DescribeMmAppResponseBodyConversationConfig) GetAsrModel() *string {
	return s.AsrModel
}

func (s *DescribeMmAppResponseBodyConversationConfig) GetOpenAsr() *bool {
	return s.OpenAsr
}

func (s *DescribeMmAppResponseBodyConversationConfig) GetOpenTts() *bool {
	return s.OpenTts
}

func (s *DescribeMmAppResponseBodyConversationConfig) GetTtsModel() *string {
	return s.TtsModel
}

func (s *DescribeMmAppResponseBodyConversationConfig) SetAsrModel(v string) *DescribeMmAppResponseBodyConversationConfig {
	s.AsrModel = &v
	return s
}

func (s *DescribeMmAppResponseBodyConversationConfig) SetOpenAsr(v bool) *DescribeMmAppResponseBodyConversationConfig {
	s.OpenAsr = &v
	return s
}

func (s *DescribeMmAppResponseBodyConversationConfig) SetOpenTts(v bool) *DescribeMmAppResponseBodyConversationConfig {
	s.OpenTts = &v
	return s
}

func (s *DescribeMmAppResponseBodyConversationConfig) SetTtsModel(v string) *DescribeMmAppResponseBodyConversationConfig {
	s.TtsModel = &v
	return s
}

func (s *DescribeMmAppResponseBodyConversationConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeMmAppResponseBodyModelConfig struct {
	// example:
	//
	// 5
	HistoryLimit *int32 `json:"HistoryLimit,omitempty" xml:"HistoryLimit,omitempty"`
	// example:
	//
	// MMH
	ModelType     *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	OpenMemory    *bool   `json:"OpenMemory,omitempty" xml:"OpenMemory,omitempty"`
	OpenWebSearch *bool   `json:"OpenWebSearch,omitempty" xml:"OpenWebSearch,omitempty"`
	// example:
	//
	// xxxx
	TextModal *string `json:"TextModal,omitempty" xml:"TextModal,omitempty"`
}

func (s DescribeMmAppResponseBodyModelConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeMmAppResponseBodyModelConfig) GoString() string {
	return s.String()
}

func (s *DescribeMmAppResponseBodyModelConfig) GetHistoryLimit() *int32 {
	return s.HistoryLimit
}

func (s *DescribeMmAppResponseBodyModelConfig) GetModelType() *string {
	return s.ModelType
}

func (s *DescribeMmAppResponseBodyModelConfig) GetOpenMemory() *bool {
	return s.OpenMemory
}

func (s *DescribeMmAppResponseBodyModelConfig) GetOpenWebSearch() *bool {
	return s.OpenWebSearch
}

func (s *DescribeMmAppResponseBodyModelConfig) GetTextModal() *string {
	return s.TextModal
}

func (s *DescribeMmAppResponseBodyModelConfig) SetHistoryLimit(v int32) *DescribeMmAppResponseBodyModelConfig {
	s.HistoryLimit = &v
	return s
}

func (s *DescribeMmAppResponseBodyModelConfig) SetModelType(v string) *DescribeMmAppResponseBodyModelConfig {
	s.ModelType = &v
	return s
}

func (s *DescribeMmAppResponseBodyModelConfig) SetOpenMemory(v bool) *DescribeMmAppResponseBodyModelConfig {
	s.OpenMemory = &v
	return s
}

func (s *DescribeMmAppResponseBodyModelConfig) SetOpenWebSearch(v bool) *DescribeMmAppResponseBodyModelConfig {
	s.OpenWebSearch = &v
	return s
}

func (s *DescribeMmAppResponseBodyModelConfig) SetTextModal(v string) *DescribeMmAppResponseBodyModelConfig {
	s.TextModal = &v
	return s
}

func (s *DescribeMmAppResponseBodyModelConfig) Validate() error {
	return dara.Validate(s)
}
