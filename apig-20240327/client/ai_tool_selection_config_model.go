// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiToolSelectionConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEnableConditions(v *AiToolSelectionConfigEnableConditions) *AiToolSelectionConfig
	GetEnableConditions() *AiToolSelectionConfigEnableConditions
	SetPluginStatus(v *AiPluginStatus) *AiToolSelectionConfig
	GetPluginStatus() *AiPluginStatus
	SetQueryRewriting(v *AiToolSelectionConfigQueryRewriting) *AiToolSelectionConfig
	GetQueryRewriting() *AiToolSelectionConfigQueryRewriting
	SetToolReranking(v *AiToolSelectionConfigToolReranking) *AiToolSelectionConfig
	GetToolReranking() *AiToolSelectionConfigToolReranking
}

type AiToolSelectionConfig struct {
	// The enable conditions configuration. Controls when the overall feature is triggered.
	EnableConditions *AiToolSelectionConfigEnableConditions `json:"enableConditions,omitempty" xml:"enableConditions,omitempty" type:"Struct"`
	// The plug-in running status.
	//
	// if can be null:
	// true
	PluginStatus *AiPluginStatus `json:"pluginStatus,omitempty" xml:"pluginStatus,omitempty"`
	// The query rewrite configuration. Rewrites user queries before tool reranking to improve matching precision.
	QueryRewriting *AiToolSelectionConfigQueryRewriting `json:"queryRewriting,omitempty" xml:"queryRewriting,omitempty" type:"Struct"`
	// The tool reranking configuration. Uses a model to rank and filter candidate tools.
	ToolReranking *AiToolSelectionConfigToolReranking `json:"toolReranking,omitempty" xml:"toolReranking,omitempty" type:"Struct"`
}

func (s AiToolSelectionConfig) String() string {
	return dara.Prettify(s)
}

func (s AiToolSelectionConfig) GoString() string {
	return s.String()
}

func (s *AiToolSelectionConfig) GetEnableConditions() *AiToolSelectionConfigEnableConditions {
	return s.EnableConditions
}

func (s *AiToolSelectionConfig) GetPluginStatus() *AiPluginStatus {
	return s.PluginStatus
}

func (s *AiToolSelectionConfig) GetQueryRewriting() *AiToolSelectionConfigQueryRewriting {
	return s.QueryRewriting
}

func (s *AiToolSelectionConfig) GetToolReranking() *AiToolSelectionConfigToolReranking {
	return s.ToolReranking
}

func (s *AiToolSelectionConfig) SetEnableConditions(v *AiToolSelectionConfigEnableConditions) *AiToolSelectionConfig {
	s.EnableConditions = v
	return s
}

func (s *AiToolSelectionConfig) SetPluginStatus(v *AiPluginStatus) *AiToolSelectionConfig {
	s.PluginStatus = v
	return s
}

func (s *AiToolSelectionConfig) SetQueryRewriting(v *AiToolSelectionConfigQueryRewriting) *AiToolSelectionConfig {
	s.QueryRewriting = v
	return s
}

func (s *AiToolSelectionConfig) SetToolReranking(v *AiToolSelectionConfigToolReranking) *AiToolSelectionConfig {
	s.ToolReranking = v
	return s
}

func (s *AiToolSelectionConfig) Validate() error {
	if s.EnableConditions != nil {
		if err := s.EnableConditions.Validate(); err != nil {
			return err
		}
	}
	if s.PluginStatus != nil {
		if err := s.PluginStatus.Validate(); err != nil {
			return err
		}
	}
	if s.QueryRewriting != nil {
		if err := s.QueryRewriting.Validate(); err != nil {
			return err
		}
	}
	if s.ToolReranking != nil {
		if err := s.ToolReranking.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AiToolSelectionConfigEnableConditions struct {
	// The tool count threshold.
	//
	// example:
	//
	// 10
	ToolCountThreshold *int32 `json:"toolCountThreshold,omitempty" xml:"toolCountThreshold,omitempty"`
}

func (s AiToolSelectionConfigEnableConditions) String() string {
	return dara.Prettify(s)
}

func (s AiToolSelectionConfigEnableConditions) GoString() string {
	return s.String()
}

func (s *AiToolSelectionConfigEnableConditions) GetToolCountThreshold() *int32 {
	return s.ToolCountThreshold
}

func (s *AiToolSelectionConfigEnableConditions) SetToolCountThreshold(v int32) *AiToolSelectionConfigEnableConditions {
	s.ToolCountThreshold = &v
	return s
}

func (s *AiToolSelectionConfigEnableConditions) Validate() error {
	return dara.Validate(s)
}

type AiToolSelectionConfigQueryRewriting struct {
	// The context selection configuration.
	ContextSelection *AiToolSelectionConfigQueryRewritingContextSelection `json:"contextSelection,omitempty" xml:"contextSelection,omitempty" type:"Struct"`
	// Specifies whether query rewrite is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The fallback strategy.
	//
	// example:
	//
	// skip
	FallbackStrategy *string `json:"fallbackStrategy,omitempty" xml:"fallbackStrategy,omitempty"`
	// The maximum number of output tokens for rewriting.
	//
	// example:
	//
	// 50
	MaxOutputTokens *int32 `json:"maxOutputTokens,omitempty" xml:"maxOutputTokens,omitempty"`
	// The rewriting model service configuration.
	ModelService *AiToolSelectionConfigQueryRewritingModelService `json:"modelService,omitempty" xml:"modelService,omitempty" type:"Struct"`
	// The prompt configuration.
	PromptConfig *AiToolSelectionConfigQueryRewritingPromptConfig `json:"promptConfig,omitempty" xml:"promptConfig,omitempty" type:"Struct"`
	// The trigger condition configuration.
	TriggerConditions *AiToolSelectionConfigQueryRewritingTriggerConditions `json:"triggerConditions,omitempty" xml:"triggerConditions,omitempty" type:"Struct"`
}

func (s AiToolSelectionConfigQueryRewriting) String() string {
	return dara.Prettify(s)
}

func (s AiToolSelectionConfigQueryRewriting) GoString() string {
	return s.String()
}

func (s *AiToolSelectionConfigQueryRewriting) GetContextSelection() *AiToolSelectionConfigQueryRewritingContextSelection {
	return s.ContextSelection
}

func (s *AiToolSelectionConfigQueryRewriting) GetEnabled() *bool {
	return s.Enabled
}

func (s *AiToolSelectionConfigQueryRewriting) GetFallbackStrategy() *string {
	return s.FallbackStrategy
}

func (s *AiToolSelectionConfigQueryRewriting) GetMaxOutputTokens() *int32 {
	return s.MaxOutputTokens
}

func (s *AiToolSelectionConfigQueryRewriting) GetModelService() *AiToolSelectionConfigQueryRewritingModelService {
	return s.ModelService
}

func (s *AiToolSelectionConfigQueryRewriting) GetPromptConfig() *AiToolSelectionConfigQueryRewritingPromptConfig {
	return s.PromptConfig
}

func (s *AiToolSelectionConfigQueryRewriting) GetTriggerConditions() *AiToolSelectionConfigQueryRewritingTriggerConditions {
	return s.TriggerConditions
}

func (s *AiToolSelectionConfigQueryRewriting) SetContextSelection(v *AiToolSelectionConfigQueryRewritingContextSelection) *AiToolSelectionConfigQueryRewriting {
	s.ContextSelection = v
	return s
}

func (s *AiToolSelectionConfigQueryRewriting) SetEnabled(v bool) *AiToolSelectionConfigQueryRewriting {
	s.Enabled = &v
	return s
}

func (s *AiToolSelectionConfigQueryRewriting) SetFallbackStrategy(v string) *AiToolSelectionConfigQueryRewriting {
	s.FallbackStrategy = &v
	return s
}

func (s *AiToolSelectionConfigQueryRewriting) SetMaxOutputTokens(v int32) *AiToolSelectionConfigQueryRewriting {
	s.MaxOutputTokens = &v
	return s
}

func (s *AiToolSelectionConfigQueryRewriting) SetModelService(v *AiToolSelectionConfigQueryRewritingModelService) *AiToolSelectionConfigQueryRewriting {
	s.ModelService = v
	return s
}

func (s *AiToolSelectionConfigQueryRewriting) SetPromptConfig(v *AiToolSelectionConfigQueryRewritingPromptConfig) *AiToolSelectionConfigQueryRewriting {
	s.PromptConfig = v
	return s
}

func (s *AiToolSelectionConfigQueryRewriting) SetTriggerConditions(v *AiToolSelectionConfigQueryRewritingTriggerConditions) *AiToolSelectionConfigQueryRewriting {
	s.TriggerConditions = v
	return s
}

func (s *AiToolSelectionConfigQueryRewriting) Validate() error {
	if s.ContextSelection != nil {
		if err := s.ContextSelection.Validate(); err != nil {
			return err
		}
	}
	if s.ModelService != nil {
		if err := s.ModelService.Validate(); err != nil {
			return err
		}
	}
	if s.PromptConfig != nil {
		if err := s.PromptConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TriggerConditions != nil {
		if err := s.TriggerConditions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AiToolSelectionConfigQueryRewritingContextSelection struct {
	// The context selection method.
	//
	// example:
	//
	// allMessages
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The number of retained messages or characters.
	//
	// example:
	//
	// 5
	Value *int32 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AiToolSelectionConfigQueryRewritingContextSelection) String() string {
	return dara.Prettify(s)
}

func (s AiToolSelectionConfigQueryRewritingContextSelection) GoString() string {
	return s.String()
}

func (s *AiToolSelectionConfigQueryRewritingContextSelection) GetType() *string {
	return s.Type
}

func (s *AiToolSelectionConfigQueryRewritingContextSelection) GetValue() *int32 {
	return s.Value
}

func (s *AiToolSelectionConfigQueryRewritingContextSelection) SetType(v string) *AiToolSelectionConfigQueryRewritingContextSelection {
	s.Type = &v
	return s
}

func (s *AiToolSelectionConfigQueryRewritingContextSelection) SetValue(v int32) *AiToolSelectionConfigQueryRewritingContextSelection {
	s.Value = &v
	return s
}

func (s *AiToolSelectionConfigQueryRewritingContextSelection) Validate() error {
	return dara.Validate(s)
}

type AiToolSelectionConfigQueryRewritingModelService struct {
	// The model name.
	//
	// example:
	//
	// gte-rerank-v2
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// The model service ID.
	//
	// example:
	//
	// svc-xxx
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// The request timeout period, in milliseconds.
	//
	// example:
	//
	// 5000
	TimeoutMillisecond *int32 `json:"timeoutMillisecond,omitempty" xml:"timeoutMillisecond,omitempty"`
}

func (s AiToolSelectionConfigQueryRewritingModelService) String() string {
	return dara.Prettify(s)
}

func (s AiToolSelectionConfigQueryRewritingModelService) GoString() string {
	return s.String()
}

func (s *AiToolSelectionConfigQueryRewritingModelService) GetModelName() *string {
	return s.ModelName
}

func (s *AiToolSelectionConfigQueryRewritingModelService) GetServiceId() *string {
	return s.ServiceId
}

func (s *AiToolSelectionConfigQueryRewritingModelService) GetTimeoutMillisecond() *int32 {
	return s.TimeoutMillisecond
}

func (s *AiToolSelectionConfigQueryRewritingModelService) SetModelName(v string) *AiToolSelectionConfigQueryRewritingModelService {
	s.ModelName = &v
	return s
}

func (s *AiToolSelectionConfigQueryRewritingModelService) SetServiceId(v string) *AiToolSelectionConfigQueryRewritingModelService {
	s.ServiceId = &v
	return s
}

func (s *AiToolSelectionConfigQueryRewritingModelService) SetTimeoutMillisecond(v int32) *AiToolSelectionConfigQueryRewritingModelService {
	s.TimeoutMillisecond = &v
	return s
}

func (s *AiToolSelectionConfigQueryRewritingModelService) Validate() error {
	return dara.Validate(s)
}

type AiToolSelectionConfigQueryRewritingPromptConfig struct {
	// The custom prompt content.
	//
	// example:
	//
	// 请将以下用户问题改写为...
	CustomPrompt *string `json:"customPrompt,omitempty" xml:"customPrompt,omitempty"`
	// The prompt type.
	//
	// example:
	//
	// builtIn
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AiToolSelectionConfigQueryRewritingPromptConfig) String() string {
	return dara.Prettify(s)
}

func (s AiToolSelectionConfigQueryRewritingPromptConfig) GoString() string {
	return s.String()
}

func (s *AiToolSelectionConfigQueryRewritingPromptConfig) GetCustomPrompt() *string {
	return s.CustomPrompt
}

func (s *AiToolSelectionConfigQueryRewritingPromptConfig) GetType() *string {
	return s.Type
}

func (s *AiToolSelectionConfigQueryRewritingPromptConfig) SetCustomPrompt(v string) *AiToolSelectionConfigQueryRewritingPromptConfig {
	s.CustomPrompt = &v
	return s
}

func (s *AiToolSelectionConfigQueryRewritingPromptConfig) SetType(v string) *AiToolSelectionConfigQueryRewritingPromptConfig {
	s.Type = &v
	return s
}

func (s *AiToolSelectionConfigQueryRewritingPromptConfig) Validate() error {
	return dara.Validate(s)
}

type AiToolSelectionConfigQueryRewritingTriggerConditions struct {
	// The number of conversation turns after which rewriting is triggered.
	//
	// example:
	//
	// 1
	MessageCountThreshold *int32 `json:"messageCountThreshold,omitempty" xml:"messageCountThreshold,omitempty"`
}

func (s AiToolSelectionConfigQueryRewritingTriggerConditions) String() string {
	return dara.Prettify(s)
}

func (s AiToolSelectionConfigQueryRewritingTriggerConditions) GoString() string {
	return s.String()
}

func (s *AiToolSelectionConfigQueryRewritingTriggerConditions) GetMessageCountThreshold() *int32 {
	return s.MessageCountThreshold
}

func (s *AiToolSelectionConfigQueryRewritingTriggerConditions) SetMessageCountThreshold(v int32) *AiToolSelectionConfigQueryRewritingTriggerConditions {
	s.MessageCountThreshold = &v
	return s
}

func (s *AiToolSelectionConfigQueryRewritingTriggerConditions) Validate() error {
	return dara.Validate(s)
}

type AiToolSelectionConfigToolReranking struct {
	// The fallback strategy upon failure.
	//
	// example:
	//
	// skip
	FallbackStrategy *string `json:"fallbackStrategy,omitempty" xml:"fallbackStrategy,omitempty"`
	// The filtering method.
	//
	// example:
	//
	// topN
	FilteringMethod *string `json:"filteringMethod,omitempty" xml:"filteringMethod,omitempty"`
	// The reranking model service configuration.
	ModelService *AiToolSelectionConfigToolRerankingModelService `json:"modelService,omitempty" xml:"modelService,omitempty" type:"Struct"`
	// The score threshold.
	//
	// example:
	//
	// 0.5
	ScoreThreshold *float32 `json:"scoreThreshold,omitempty" xml:"scoreThreshold,omitempty"`
	// The retention percentage.
	//
	// example:
	//
	// 50
	TopKPercent *int32 `json:"topKPercent,omitempty" xml:"topKPercent,omitempty"`
	// The retention count.
	//
	// example:
	//
	// 5
	TopNCount *int32 `json:"topNCount,omitempty" xml:"topNCount,omitempty"`
}

func (s AiToolSelectionConfigToolReranking) String() string {
	return dara.Prettify(s)
}

func (s AiToolSelectionConfigToolReranking) GoString() string {
	return s.String()
}

func (s *AiToolSelectionConfigToolReranking) GetFallbackStrategy() *string {
	return s.FallbackStrategy
}

func (s *AiToolSelectionConfigToolReranking) GetFilteringMethod() *string {
	return s.FilteringMethod
}

func (s *AiToolSelectionConfigToolReranking) GetModelService() *AiToolSelectionConfigToolRerankingModelService {
	return s.ModelService
}

func (s *AiToolSelectionConfigToolReranking) GetScoreThreshold() *float32 {
	return s.ScoreThreshold
}

func (s *AiToolSelectionConfigToolReranking) GetTopKPercent() *int32 {
	return s.TopKPercent
}

func (s *AiToolSelectionConfigToolReranking) GetTopNCount() *int32 {
	return s.TopNCount
}

func (s *AiToolSelectionConfigToolReranking) SetFallbackStrategy(v string) *AiToolSelectionConfigToolReranking {
	s.FallbackStrategy = &v
	return s
}

func (s *AiToolSelectionConfigToolReranking) SetFilteringMethod(v string) *AiToolSelectionConfigToolReranking {
	s.FilteringMethod = &v
	return s
}

func (s *AiToolSelectionConfigToolReranking) SetModelService(v *AiToolSelectionConfigToolRerankingModelService) *AiToolSelectionConfigToolReranking {
	s.ModelService = v
	return s
}

func (s *AiToolSelectionConfigToolReranking) SetScoreThreshold(v float32) *AiToolSelectionConfigToolReranking {
	s.ScoreThreshold = &v
	return s
}

func (s *AiToolSelectionConfigToolReranking) SetTopKPercent(v int32) *AiToolSelectionConfigToolReranking {
	s.TopKPercent = &v
	return s
}

func (s *AiToolSelectionConfigToolReranking) SetTopNCount(v int32) *AiToolSelectionConfigToolReranking {
	s.TopNCount = &v
	return s
}

func (s *AiToolSelectionConfigToolReranking) Validate() error {
	if s.ModelService != nil {
		if err := s.ModelService.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AiToolSelectionConfigToolRerankingModelService struct {
	// The model name.
	//
	// example:
	//
	// gte-rerank-v2
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// The model service ID.
	//
	// example:
	//
	// svc-xxx
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// The request timeout period, in milliseconds.
	//
	// example:
	//
	// 5000
	TimeoutMillisecond *int32 `json:"timeoutMillisecond,omitempty" xml:"timeoutMillisecond,omitempty"`
}

func (s AiToolSelectionConfigToolRerankingModelService) String() string {
	return dara.Prettify(s)
}

func (s AiToolSelectionConfigToolRerankingModelService) GoString() string {
	return s.String()
}

func (s *AiToolSelectionConfigToolRerankingModelService) GetModelName() *string {
	return s.ModelName
}

func (s *AiToolSelectionConfigToolRerankingModelService) GetServiceId() *string {
	return s.ServiceId
}

func (s *AiToolSelectionConfigToolRerankingModelService) GetTimeoutMillisecond() *int32 {
	return s.TimeoutMillisecond
}

func (s *AiToolSelectionConfigToolRerankingModelService) SetModelName(v string) *AiToolSelectionConfigToolRerankingModelService {
	s.ModelName = &v
	return s
}

func (s *AiToolSelectionConfigToolRerankingModelService) SetServiceId(v string) *AiToolSelectionConfigToolRerankingModelService {
	s.ServiceId = &v
	return s
}

func (s *AiToolSelectionConfigToolRerankingModelService) SetTimeoutMillisecond(v int32) *AiToolSelectionConfigToolRerankingModelService {
	s.TimeoutMillisecond = &v
	return s
}

func (s *AiToolSelectionConfigToolRerankingModelService) Validate() error {
	return dara.Validate(s)
}
