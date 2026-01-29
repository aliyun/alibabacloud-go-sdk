// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunMarketingInformationWritingShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *RunMarketingInformationWritingShrinkRequest
	GetApiKey() *string
	SetCustomLimitation(v string) *RunMarketingInformationWritingShrinkRequest
	GetCustomLimitation() *string
	SetCustomPrompt(v string) *RunMarketingInformationWritingShrinkRequest
	GetCustomPrompt() *string
	SetExtParametersShrink(v string) *RunMarketingInformationWritingShrinkRequest
	GetExtParametersShrink() *string
	SetGenerateCount(v string) *RunMarketingInformationWritingShrinkRequest
	GetGenerateCount() *string
	SetInputExample(v string) *RunMarketingInformationWritingShrinkRequest
	GetInputExample() *string
	SetKeywords(v string) *RunMarketingInformationWritingShrinkRequest
	GetKeywords() *string
	SetLanguage(v string) *RunMarketingInformationWritingShrinkRequest
	GetLanguage() *string
	SetModelId(v string) *RunMarketingInformationWritingShrinkRequest
	GetModelId() *string
	SetOtherRequirements(v string) *RunMarketingInformationWritingShrinkRequest
	GetOtherRequirements() *string
	SetOutputExample(v string) *RunMarketingInformationWritingShrinkRequest
	GetOutputExample() *string
	SetPrompt(v string) *RunMarketingInformationWritingShrinkRequest
	GetPrompt() *string
	SetSourceMaterial(v string) *RunMarketingInformationWritingShrinkRequest
	GetSourceMaterial() *string
	SetWordCountRange(v string) *RunMarketingInformationWritingShrinkRequest
	GetWordCountRange() *string
	SetWritingType(v string) *RunMarketingInformationWritingShrinkRequest
	GetWritingType() *string
}

type RunMarketingInformationWritingShrinkRequest struct {
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// example:
	//
	// 标题长度不超过30个字符
	CustomLimitation *string `json:"customLimitation,omitempty" xml:"customLimitation,omitempty"`
	// example:
	//
	// 请根据商品特点生成吸引人的标题
	CustomPrompt *string `json:"customPrompt,omitempty" xml:"customPrompt,omitempty"`
	// example:
	//
	// {
	//
	//               "minWordLength": "10",
	//
	//               "maxWordLength": "50",
	//
	//               "enableThinking": "true",
	//
	//               "thinkingBudget": "2000"
	//
	//             }
	ExtParametersShrink *string `json:"extParameters,omitempty" xml:"extParameters,omitempty"`
	// example:
	//
	// 5
	GenerateCount *string `json:"generateCount,omitempty" xml:"generateCount,omitempty"`
	// example:
	//
	// 商品特点：透气、减震、舒适
	InputExample *string `json:"inputExample,omitempty" xml:"inputExample,omitempty"`
	// example:
	//
	// 运动鞋 透气 减震 跑步
	Keywords *string `json:"keywords,omitempty" xml:"keywords,omitempty"`
	// example:
	//
	// en
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// example:
	//
	// qwen-max
	//
	// qwen-plus
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 要求标题简洁有力，突出产品特点
	OtherRequirements *string `json:"otherRequirements,omitempty" xml:"otherRequirements,omitempty"`
	// example:
	//
	// 透气减震运动鞋，舒适跑步首选
	OutputExample *string `json:"outputExample,omitempty" xml:"outputExample,omitempty"`
	// example:
	//
	// 请根据关键词生成吸引人的商品标题
	Prompt *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
	// example:
	//
	// 时尚休闲运动鞋
	SourceMaterial *string `json:"sourceMaterial,omitempty" xml:"sourceMaterial,omitempty"`
	// example:
	//
	// 10-20
	WordCountRange *string `json:"wordCountRange,omitempty" xml:"wordCountRange,omitempty"`
	// example:
	//
	// generateProductTitle
	WritingType *string `json:"writingType,omitempty" xml:"writingType,omitempty"`
}

func (s RunMarketingInformationWritingShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunMarketingInformationWritingShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunMarketingInformationWritingShrinkRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *RunMarketingInformationWritingShrinkRequest) GetCustomLimitation() *string {
	return s.CustomLimitation
}

func (s *RunMarketingInformationWritingShrinkRequest) GetCustomPrompt() *string {
	return s.CustomPrompt
}

func (s *RunMarketingInformationWritingShrinkRequest) GetExtParametersShrink() *string {
	return s.ExtParametersShrink
}

func (s *RunMarketingInformationWritingShrinkRequest) GetGenerateCount() *string {
	return s.GenerateCount
}

func (s *RunMarketingInformationWritingShrinkRequest) GetInputExample() *string {
	return s.InputExample
}

func (s *RunMarketingInformationWritingShrinkRequest) GetKeywords() *string {
	return s.Keywords
}

func (s *RunMarketingInformationWritingShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *RunMarketingInformationWritingShrinkRequest) GetModelId() *string {
	return s.ModelId
}

func (s *RunMarketingInformationWritingShrinkRequest) GetOtherRequirements() *string {
	return s.OtherRequirements
}

func (s *RunMarketingInformationWritingShrinkRequest) GetOutputExample() *string {
	return s.OutputExample
}

func (s *RunMarketingInformationWritingShrinkRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunMarketingInformationWritingShrinkRequest) GetSourceMaterial() *string {
	return s.SourceMaterial
}

func (s *RunMarketingInformationWritingShrinkRequest) GetWordCountRange() *string {
	return s.WordCountRange
}

func (s *RunMarketingInformationWritingShrinkRequest) GetWritingType() *string {
	return s.WritingType
}

func (s *RunMarketingInformationWritingShrinkRequest) SetApiKey(v string) *RunMarketingInformationWritingShrinkRequest {
	s.ApiKey = &v
	return s
}

func (s *RunMarketingInformationWritingShrinkRequest) SetCustomLimitation(v string) *RunMarketingInformationWritingShrinkRequest {
	s.CustomLimitation = &v
	return s
}

func (s *RunMarketingInformationWritingShrinkRequest) SetCustomPrompt(v string) *RunMarketingInformationWritingShrinkRequest {
	s.CustomPrompt = &v
	return s
}

func (s *RunMarketingInformationWritingShrinkRequest) SetExtParametersShrink(v string) *RunMarketingInformationWritingShrinkRequest {
	s.ExtParametersShrink = &v
	return s
}

func (s *RunMarketingInformationWritingShrinkRequest) SetGenerateCount(v string) *RunMarketingInformationWritingShrinkRequest {
	s.GenerateCount = &v
	return s
}

func (s *RunMarketingInformationWritingShrinkRequest) SetInputExample(v string) *RunMarketingInformationWritingShrinkRequest {
	s.InputExample = &v
	return s
}

func (s *RunMarketingInformationWritingShrinkRequest) SetKeywords(v string) *RunMarketingInformationWritingShrinkRequest {
	s.Keywords = &v
	return s
}

func (s *RunMarketingInformationWritingShrinkRequest) SetLanguage(v string) *RunMarketingInformationWritingShrinkRequest {
	s.Language = &v
	return s
}

func (s *RunMarketingInformationWritingShrinkRequest) SetModelId(v string) *RunMarketingInformationWritingShrinkRequest {
	s.ModelId = &v
	return s
}

func (s *RunMarketingInformationWritingShrinkRequest) SetOtherRequirements(v string) *RunMarketingInformationWritingShrinkRequest {
	s.OtherRequirements = &v
	return s
}

func (s *RunMarketingInformationWritingShrinkRequest) SetOutputExample(v string) *RunMarketingInformationWritingShrinkRequest {
	s.OutputExample = &v
	return s
}

func (s *RunMarketingInformationWritingShrinkRequest) SetPrompt(v string) *RunMarketingInformationWritingShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *RunMarketingInformationWritingShrinkRequest) SetSourceMaterial(v string) *RunMarketingInformationWritingShrinkRequest {
	s.SourceMaterial = &v
	return s
}

func (s *RunMarketingInformationWritingShrinkRequest) SetWordCountRange(v string) *RunMarketingInformationWritingShrinkRequest {
	s.WordCountRange = &v
	return s
}

func (s *RunMarketingInformationWritingShrinkRequest) SetWritingType(v string) *RunMarketingInformationWritingShrinkRequest {
	s.WritingType = &v
	return s
}

func (s *RunMarketingInformationWritingShrinkRequest) Validate() error {
	return dara.Validate(s)
}
