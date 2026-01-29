// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunMarketingInformationWritingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *RunMarketingInformationWritingRequest
	GetApiKey() *string
	SetCustomLimitation(v string) *RunMarketingInformationWritingRequest
	GetCustomLimitation() *string
	SetCustomPrompt(v string) *RunMarketingInformationWritingRequest
	GetCustomPrompt() *string
	SetExtParameters(v map[string]*string) *RunMarketingInformationWritingRequest
	GetExtParameters() map[string]*string
	SetGenerateCount(v string) *RunMarketingInformationWritingRequest
	GetGenerateCount() *string
	SetInputExample(v string) *RunMarketingInformationWritingRequest
	GetInputExample() *string
	SetKeywords(v string) *RunMarketingInformationWritingRequest
	GetKeywords() *string
	SetLanguage(v string) *RunMarketingInformationWritingRequest
	GetLanguage() *string
	SetModelId(v string) *RunMarketingInformationWritingRequest
	GetModelId() *string
	SetOtherRequirements(v string) *RunMarketingInformationWritingRequest
	GetOtherRequirements() *string
	SetOutputExample(v string) *RunMarketingInformationWritingRequest
	GetOutputExample() *string
	SetPrompt(v string) *RunMarketingInformationWritingRequest
	GetPrompt() *string
	SetSourceMaterial(v string) *RunMarketingInformationWritingRequest
	GetSourceMaterial() *string
	SetWordCountRange(v string) *RunMarketingInformationWritingRequest
	GetWordCountRange() *string
	SetWritingType(v string) *RunMarketingInformationWritingRequest
	GetWritingType() *string
}

type RunMarketingInformationWritingRequest struct {
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
	ExtParameters map[string]*string `json:"extParameters,omitempty" xml:"extParameters,omitempty"`
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

func (s RunMarketingInformationWritingRequest) String() string {
	return dara.Prettify(s)
}

func (s RunMarketingInformationWritingRequest) GoString() string {
	return s.String()
}

func (s *RunMarketingInformationWritingRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *RunMarketingInformationWritingRequest) GetCustomLimitation() *string {
	return s.CustomLimitation
}

func (s *RunMarketingInformationWritingRequest) GetCustomPrompt() *string {
	return s.CustomPrompt
}

func (s *RunMarketingInformationWritingRequest) GetExtParameters() map[string]*string {
	return s.ExtParameters
}

func (s *RunMarketingInformationWritingRequest) GetGenerateCount() *string {
	return s.GenerateCount
}

func (s *RunMarketingInformationWritingRequest) GetInputExample() *string {
	return s.InputExample
}

func (s *RunMarketingInformationWritingRequest) GetKeywords() *string {
	return s.Keywords
}

func (s *RunMarketingInformationWritingRequest) GetLanguage() *string {
	return s.Language
}

func (s *RunMarketingInformationWritingRequest) GetModelId() *string {
	return s.ModelId
}

func (s *RunMarketingInformationWritingRequest) GetOtherRequirements() *string {
	return s.OtherRequirements
}

func (s *RunMarketingInformationWritingRequest) GetOutputExample() *string {
	return s.OutputExample
}

func (s *RunMarketingInformationWritingRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunMarketingInformationWritingRequest) GetSourceMaterial() *string {
	return s.SourceMaterial
}

func (s *RunMarketingInformationWritingRequest) GetWordCountRange() *string {
	return s.WordCountRange
}

func (s *RunMarketingInformationWritingRequest) GetWritingType() *string {
	return s.WritingType
}

func (s *RunMarketingInformationWritingRequest) SetApiKey(v string) *RunMarketingInformationWritingRequest {
	s.ApiKey = &v
	return s
}

func (s *RunMarketingInformationWritingRequest) SetCustomLimitation(v string) *RunMarketingInformationWritingRequest {
	s.CustomLimitation = &v
	return s
}

func (s *RunMarketingInformationWritingRequest) SetCustomPrompt(v string) *RunMarketingInformationWritingRequest {
	s.CustomPrompt = &v
	return s
}

func (s *RunMarketingInformationWritingRequest) SetExtParameters(v map[string]*string) *RunMarketingInformationWritingRequest {
	s.ExtParameters = v
	return s
}

func (s *RunMarketingInformationWritingRequest) SetGenerateCount(v string) *RunMarketingInformationWritingRequest {
	s.GenerateCount = &v
	return s
}

func (s *RunMarketingInformationWritingRequest) SetInputExample(v string) *RunMarketingInformationWritingRequest {
	s.InputExample = &v
	return s
}

func (s *RunMarketingInformationWritingRequest) SetKeywords(v string) *RunMarketingInformationWritingRequest {
	s.Keywords = &v
	return s
}

func (s *RunMarketingInformationWritingRequest) SetLanguage(v string) *RunMarketingInformationWritingRequest {
	s.Language = &v
	return s
}

func (s *RunMarketingInformationWritingRequest) SetModelId(v string) *RunMarketingInformationWritingRequest {
	s.ModelId = &v
	return s
}

func (s *RunMarketingInformationWritingRequest) SetOtherRequirements(v string) *RunMarketingInformationWritingRequest {
	s.OtherRequirements = &v
	return s
}

func (s *RunMarketingInformationWritingRequest) SetOutputExample(v string) *RunMarketingInformationWritingRequest {
	s.OutputExample = &v
	return s
}

func (s *RunMarketingInformationWritingRequest) SetPrompt(v string) *RunMarketingInformationWritingRequest {
	s.Prompt = &v
	return s
}

func (s *RunMarketingInformationWritingRequest) SetSourceMaterial(v string) *RunMarketingInformationWritingRequest {
	s.SourceMaterial = &v
	return s
}

func (s *RunMarketingInformationWritingRequest) SetWordCountRange(v string) *RunMarketingInformationWritingRequest {
	s.WordCountRange = &v
	return s
}

func (s *RunMarketingInformationWritingRequest) SetWritingType(v string) *RunMarketingInformationWritingRequest {
	s.WritingType = &v
	return s
}

func (s *RunMarketingInformationWritingRequest) Validate() error {
	return dara.Validate(s)
}
