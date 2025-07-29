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
	SetInputExample(v string) *RunMarketingInformationWritingRequest
	GetInputExample() *string
	SetModelId(v string) *RunMarketingInformationWritingRequest
	GetModelId() *string
	SetOutputExample(v string) *RunMarketingInformationWritingRequest
	GetOutputExample() *string
	SetSourceMaterial(v string) *RunMarketingInformationWritingRequest
	GetSourceMaterial() *string
	SetWritingType(v string) *RunMarketingInformationWritingRequest
	GetWritingType() *string
}

type RunMarketingInformationWritingRequest struct {
	ApiKey           *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	CustomLimitation *string `json:"customLimitation,omitempty" xml:"customLimitation,omitempty"`
	CustomPrompt     *string `json:"customPrompt,omitempty" xml:"customPrompt,omitempty"`
	InputExample     *string `json:"inputExample,omitempty" xml:"inputExample,omitempty"`
	// example:
	//
	// qwen-max
	//
	// qwen-plus
	ModelId        *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	OutputExample  *string `json:"outputExample,omitempty" xml:"outputExample,omitempty"`
	SourceMaterial *string `json:"sourceMaterial,omitempty" xml:"sourceMaterial,omitempty"`
	WritingType    *string `json:"writingType,omitempty" xml:"writingType,omitempty"`
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

func (s *RunMarketingInformationWritingRequest) GetInputExample() *string {
	return s.InputExample
}

func (s *RunMarketingInformationWritingRequest) GetModelId() *string {
	return s.ModelId
}

func (s *RunMarketingInformationWritingRequest) GetOutputExample() *string {
	return s.OutputExample
}

func (s *RunMarketingInformationWritingRequest) GetSourceMaterial() *string {
	return s.SourceMaterial
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

func (s *RunMarketingInformationWritingRequest) SetInputExample(v string) *RunMarketingInformationWritingRequest {
	s.InputExample = &v
	return s
}

func (s *RunMarketingInformationWritingRequest) SetModelId(v string) *RunMarketingInformationWritingRequest {
	s.ModelId = &v
	return s
}

func (s *RunMarketingInformationWritingRequest) SetOutputExample(v string) *RunMarketingInformationWritingRequest {
	s.OutputExample = &v
	return s
}

func (s *RunMarketingInformationWritingRequest) SetSourceMaterial(v string) *RunMarketingInformationWritingRequest {
	s.SourceMaterial = &v
	return s
}

func (s *RunMarketingInformationWritingRequest) SetWritingType(v string) *RunMarketingInformationWritingRequest {
	s.WritingType = &v
	return s
}

func (s *RunMarketingInformationWritingRequest) Validate() error {
	return dara.Validate(s)
}
