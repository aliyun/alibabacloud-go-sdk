// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCapabilitiesShrink(v string) *ListModelsShrinkRequest
	GetCapabilitiesShrink() *string
	SetContextWindow(v int64) *ListModelsShrinkRequest
	GetContextWindow() *int64
	SetFeaturesShrink(v string) *ListModelsShrinkRequest
	GetFeaturesShrink() *string
	SetLanguage(v string) *ListModelsShrinkRequest
	GetLanguage() *string
	SetMaxResults(v int64) *ListModelsShrinkRequest
	GetMaxResults() *int64
	SetModel(v string) *ListModelsShrinkRequest
	GetModel() *string
	SetName(v string) *ListModelsShrinkRequest
	GetName() *string
	SetNextToken(v string) *ListModelsShrinkRequest
	GetNextToken() *string
	SetProvidersShrink(v string) *ListModelsShrinkRequest
	GetProvidersShrink() *string
}

type ListModelsShrinkRequest struct {
	// The model type. Valid values:
	//
	// - Reasoning: Deep thinking.
	//
	// - VU: Visual understanding.
	//
	// - IG: Image generation.
	//
	// - VG: Video generation.
	//
	// - ASR: Speech recognition.
	//
	// - TTS: Speech synthesis.
	//
	// - ME: Multimodal embedding.
	//
	// - Realtime-Omni: Real-time omni-modal.
	//
	// - Multimodal-Omni: Omni-modal.
	//
	// - Realtime-Text-to-Speech: Real-time speech synthesis.
	//
	// - TG: Text generation.
	//
	// - TR: Text embedding.
	//
	// - Realtime-ASR: Real-time speech recognition.
	//
	// - Realtime-Audio-Translate: Real-time audio translation.
	//
	// - 3D-generation: 3D generation.
	//
	// - Realtime-Chatting: Real-time voice conversation.
	CapabilitiesShrink *string `json:"capabilities,omitempty" xml:"capabilities,omitempty"`
	// The context window size used to filter models with a context window smaller than this value.
	//
	// example:
	//
	// 10
	ContextWindow *int64 `json:"contextWindow,omitempty" xml:"contextWindow,omitempty"`
	// The model capability. Valid values:
	//
	// - model-experience: Playground.
	//
	// - function-calling: Function calling.
	//
	// - structured-outputs: Structured output.
	//
	// - web-search: Web search.
	//
	// - prefix-completion: Partial mode.
	//
	// - cache: Cache.
	//
	// - batch: Batch inference.
	//
	// - fine-tuning: Model fine-tuning.
	FeaturesShrink *string `json:"features,omitempty" xml:"features,omitempty"`
	// The language. Valid values:
	//
	// - zh-CN: Chinese.
	//
	// - en-US: English.
	//
	// example:
	//
	// zh-CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// The maximum number of results to return. Valid values: 0 to 200.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The model name for exact match.
	//
	// example:
	//
	// qwen-plus
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// The model name for fuzzy match.
	//
	// example:
	//
	// qwen-plus
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The next token.
	//
	// example:
	//
	// lwytFRtLdNk=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The provider. Valid values:
	//
	// - qwen: Alibaba Cloud - Qwen.
	//
	// - zhipu-ai: Zhipu AI.
	//
	// - wan: Alibaba Cloud - Wan.
	//
	// - qwen-domain-model: Alibaba Cloud - Domain Model.
	//
	// - mini-max: MiniMax.
	//
	// - moonshot-ai: Moonshot AI.
	//
	// - deepseek: DeepSeek.
	//
	// - happyhorse: Alibaba Cloud - HappyHorse.
	//
	// - kling: Kling AI.
	//
	// - pixverse: PixVerse.
	//
	// - vidu: Vidu.
	//
	// - tripo: Tripo.
	//
	// - xiaomi: Xiaomi MiMo.
	ProvidersShrink *string `json:"providers,omitempty" xml:"providers,omitempty"`
}

func (s ListModelsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListModelsShrinkRequest) GetCapabilitiesShrink() *string {
	return s.CapabilitiesShrink
}

func (s *ListModelsShrinkRequest) GetContextWindow() *int64 {
	return s.ContextWindow
}

func (s *ListModelsShrinkRequest) GetFeaturesShrink() *string {
	return s.FeaturesShrink
}

func (s *ListModelsShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *ListModelsShrinkRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListModelsShrinkRequest) GetModel() *string {
	return s.Model
}

func (s *ListModelsShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListModelsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListModelsShrinkRequest) GetProvidersShrink() *string {
	return s.ProvidersShrink
}

func (s *ListModelsShrinkRequest) SetCapabilitiesShrink(v string) *ListModelsShrinkRequest {
	s.CapabilitiesShrink = &v
	return s
}

func (s *ListModelsShrinkRequest) SetContextWindow(v int64) *ListModelsShrinkRequest {
	s.ContextWindow = &v
	return s
}

func (s *ListModelsShrinkRequest) SetFeaturesShrink(v string) *ListModelsShrinkRequest {
	s.FeaturesShrink = &v
	return s
}

func (s *ListModelsShrinkRequest) SetLanguage(v string) *ListModelsShrinkRequest {
	s.Language = &v
	return s
}

func (s *ListModelsShrinkRequest) SetMaxResults(v int64) *ListModelsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListModelsShrinkRequest) SetModel(v string) *ListModelsShrinkRequest {
	s.Model = &v
	return s
}

func (s *ListModelsShrinkRequest) SetName(v string) *ListModelsShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListModelsShrinkRequest) SetNextToken(v string) *ListModelsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListModelsShrinkRequest) SetProvidersShrink(v string) *ListModelsShrinkRequest {
	s.ProvidersShrink = &v
	return s
}

func (s *ListModelsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
