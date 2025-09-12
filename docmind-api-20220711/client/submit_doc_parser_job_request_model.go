// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDocParserJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomOssConfig(v *SubmitDocParserJobRequestCustomOssConfig) *SubmitDocParserJobRequest
	GetCustomOssConfig() *SubmitDocParserJobRequestCustomOssConfig
	SetEnhancementMode(v string) *SubmitDocParserJobRequest
	GetEnhancementMode() *string
	SetFileName(v string) *SubmitDocParserJobRequest
	GetFileName() *string
	SetFileNameExtension(v string) *SubmitDocParserJobRequest
	GetFileNameExtension() *string
	SetFileUrl(v string) *SubmitDocParserJobRequest
	GetFileUrl() *string
	SetFormulaEnhancement(v bool) *SubmitDocParserJobRequest
	GetFormulaEnhancement() *bool
	SetLLMParam(v *SubmitDocParserJobRequestLLMParam) *SubmitDocParserJobRequest
	GetLLMParam() *SubmitDocParserJobRequestLLMParam
	SetLlmEnhancement(v bool) *SubmitDocParserJobRequest
	GetLlmEnhancement() *bool
	SetMultimediaParameters(v *SubmitDocParserJobRequestMultimediaParameters) *SubmitDocParserJobRequest
	GetMultimediaParameters() *SubmitDocParserJobRequestMultimediaParameters
	SetOption(v string) *SubmitDocParserJobRequest
	GetOption() *string
	SetOssBucket(v string) *SubmitDocParserJobRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *SubmitDocParserJobRequest
	GetOssEndpoint() *string
	SetOutputHtmlTable(v bool) *SubmitDocParserJobRequest
	GetOutputHtmlTable() *bool
	SetPageIndex(v string) *SubmitDocParserJobRequest
	GetPageIndex() *string
}

type SubmitDocParserJobRequest struct {
	CustomOssConfig *SubmitDocParserJobRequestCustomOssConfig `json:"CustomOssConfig,omitempty" xml:"CustomOssConfig,omitempty" type:"Struct"`
	EnhancementMode *string                                   `json:"EnhancementMode,omitempty" xml:"EnhancementMode,omitempty"`
	// example:
	//
	// docStructure.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// pdf
	FileNameExtension *string `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	// example:
	//
	// https://gw.alipayobjects.com/os/basement_prod/598b9edf-5287-4065-9e36-464305c60698.pdf
	FileUrl              *string                                        `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	FormulaEnhancement   *bool                                          `json:"FormulaEnhancement,omitempty" xml:"FormulaEnhancement,omitempty"`
	LLMParam             *SubmitDocParserJobRequestLLMParam             `json:"LLMParam,omitempty" xml:"LLMParam,omitempty" type:"Struct"`
	LlmEnhancement       *bool                                          `json:"LlmEnhancement,omitempty" xml:"LlmEnhancement,omitempty"`
	MultimediaParameters *SubmitDocParserJobRequestMultimediaParameters `json:"MultimediaParameters,omitempty" xml:"MultimediaParameters,omitempty" type:"Struct"`
	Option               *string                                        `json:"Option,omitempty" xml:"Option,omitempty"`
	OssBucket            *string                                        `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint          *string                                        `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	OutputHtmlTable      *bool                                          `json:"OutputHtmlTable,omitempty" xml:"OutputHtmlTable,omitempty"`
	PageIndex            *string                                        `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
}

func (s SubmitDocParserJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocParserJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitDocParserJobRequest) GetCustomOssConfig() *SubmitDocParserJobRequestCustomOssConfig {
	return s.CustomOssConfig
}

func (s *SubmitDocParserJobRequest) GetEnhancementMode() *string {
	return s.EnhancementMode
}

func (s *SubmitDocParserJobRequest) GetFileName() *string {
	return s.FileName
}

func (s *SubmitDocParserJobRequest) GetFileNameExtension() *string {
	return s.FileNameExtension
}

func (s *SubmitDocParserJobRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *SubmitDocParserJobRequest) GetFormulaEnhancement() *bool {
	return s.FormulaEnhancement
}

func (s *SubmitDocParserJobRequest) GetLLMParam() *SubmitDocParserJobRequestLLMParam {
	return s.LLMParam
}

func (s *SubmitDocParserJobRequest) GetLlmEnhancement() *bool {
	return s.LlmEnhancement
}

func (s *SubmitDocParserJobRequest) GetMultimediaParameters() *SubmitDocParserJobRequestMultimediaParameters {
	return s.MultimediaParameters
}

func (s *SubmitDocParserJobRequest) GetOption() *string {
	return s.Option
}

func (s *SubmitDocParserJobRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *SubmitDocParserJobRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *SubmitDocParserJobRequest) GetOutputHtmlTable() *bool {
	return s.OutputHtmlTable
}

func (s *SubmitDocParserJobRequest) GetPageIndex() *string {
	return s.PageIndex
}

func (s *SubmitDocParserJobRequest) SetCustomOssConfig(v *SubmitDocParserJobRequestCustomOssConfig) *SubmitDocParserJobRequest {
	s.CustomOssConfig = v
	return s
}

func (s *SubmitDocParserJobRequest) SetEnhancementMode(v string) *SubmitDocParserJobRequest {
	s.EnhancementMode = &v
	return s
}

func (s *SubmitDocParserJobRequest) SetFileName(v string) *SubmitDocParserJobRequest {
	s.FileName = &v
	return s
}

func (s *SubmitDocParserJobRequest) SetFileNameExtension(v string) *SubmitDocParserJobRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitDocParserJobRequest) SetFileUrl(v string) *SubmitDocParserJobRequest {
	s.FileUrl = &v
	return s
}

func (s *SubmitDocParserJobRequest) SetFormulaEnhancement(v bool) *SubmitDocParserJobRequest {
	s.FormulaEnhancement = &v
	return s
}

func (s *SubmitDocParserJobRequest) SetLLMParam(v *SubmitDocParserJobRequestLLMParam) *SubmitDocParserJobRequest {
	s.LLMParam = v
	return s
}

func (s *SubmitDocParserJobRequest) SetLlmEnhancement(v bool) *SubmitDocParserJobRequest {
	s.LlmEnhancement = &v
	return s
}

func (s *SubmitDocParserJobRequest) SetMultimediaParameters(v *SubmitDocParserJobRequestMultimediaParameters) *SubmitDocParserJobRequest {
	s.MultimediaParameters = v
	return s
}

func (s *SubmitDocParserJobRequest) SetOption(v string) *SubmitDocParserJobRequest {
	s.Option = &v
	return s
}

func (s *SubmitDocParserJobRequest) SetOssBucket(v string) *SubmitDocParserJobRequest {
	s.OssBucket = &v
	return s
}

func (s *SubmitDocParserJobRequest) SetOssEndpoint(v string) *SubmitDocParserJobRequest {
	s.OssEndpoint = &v
	return s
}

func (s *SubmitDocParserJobRequest) SetOutputHtmlTable(v bool) *SubmitDocParserJobRequest {
	s.OutputHtmlTable = &v
	return s
}

func (s *SubmitDocParserJobRequest) SetPageIndex(v string) *SubmitDocParserJobRequest {
	s.PageIndex = &v
	return s
}

func (s *SubmitDocParserJobRequest) Validate() error {
	return dara.Validate(s)
}

type SubmitDocParserJobRequestCustomOssConfig struct {
	// example:
	//
	// AccessId
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// example:
	//
	// AccessKeySecret
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	// example:
	//
	// StsToken
	StsToken *string `json:"StsToken,omitempty" xml:"StsToken,omitempty"`
}

func (s SubmitDocParserJobRequestCustomOssConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocParserJobRequestCustomOssConfig) GoString() string {
	return s.String()
}

func (s *SubmitDocParserJobRequestCustomOssConfig) GetAccessId() *string {
	return s.AccessId
}

func (s *SubmitDocParserJobRequestCustomOssConfig) GetAccessKeySecret() *string {
	return s.AccessKeySecret
}

func (s *SubmitDocParserJobRequestCustomOssConfig) GetStsToken() *string {
	return s.StsToken
}

func (s *SubmitDocParserJobRequestCustomOssConfig) SetAccessId(v string) *SubmitDocParserJobRequestCustomOssConfig {
	s.AccessId = &v
	return s
}

func (s *SubmitDocParserJobRequestCustomOssConfig) SetAccessKeySecret(v string) *SubmitDocParserJobRequestCustomOssConfig {
	s.AccessKeySecret = &v
	return s
}

func (s *SubmitDocParserJobRequestCustomOssConfig) SetStsToken(v string) *SubmitDocParserJobRequestCustomOssConfig {
	s.StsToken = &v
	return s
}

func (s *SubmitDocParserJobRequestCustomOssConfig) Validate() error {
	return dara.Validate(s)
}

type SubmitDocParserJobRequestLLMParam struct {
	// example:
	//
	// qwen-vl-ocr-latest
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// Read all the text from the image.
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
}

func (s SubmitDocParserJobRequestLLMParam) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocParserJobRequestLLMParam) GoString() string {
	return s.String()
}

func (s *SubmitDocParserJobRequestLLMParam) GetModel() *string {
	return s.Model
}

func (s *SubmitDocParserJobRequestLLMParam) GetPrompt() *string {
	return s.Prompt
}

func (s *SubmitDocParserJobRequestLLMParam) SetModel(v string) *SubmitDocParserJobRequestLLMParam {
	s.Model = &v
	return s
}

func (s *SubmitDocParserJobRequestLLMParam) SetPrompt(v string) *SubmitDocParserJobRequestLLMParam {
	s.Prompt = &v
	return s
}

func (s *SubmitDocParserJobRequestLLMParam) Validate() error {
	return dara.Validate(s)
}

type SubmitDocParserJobRequestMultimediaParameters struct {
	EnableSynopsisParse *bool   `json:"EnableSynopsisParse,omitempty" xml:"EnableSynopsisParse,omitempty"`
	VlParsePrompt       *string `json:"VlParsePrompt,omitempty" xml:"VlParsePrompt,omitempty"`
}

func (s SubmitDocParserJobRequestMultimediaParameters) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocParserJobRequestMultimediaParameters) GoString() string {
	return s.String()
}

func (s *SubmitDocParserJobRequestMultimediaParameters) GetEnableSynopsisParse() *bool {
	return s.EnableSynopsisParse
}

func (s *SubmitDocParserJobRequestMultimediaParameters) GetVlParsePrompt() *string {
	return s.VlParsePrompt
}

func (s *SubmitDocParserJobRequestMultimediaParameters) SetEnableSynopsisParse(v bool) *SubmitDocParserJobRequestMultimediaParameters {
	s.EnableSynopsisParse = &v
	return s
}

func (s *SubmitDocParserJobRequestMultimediaParameters) SetVlParsePrompt(v string) *SubmitDocParserJobRequestMultimediaParameters {
	s.VlParsePrompt = &v
	return s
}

func (s *SubmitDocParserJobRequestMultimediaParameters) Validate() error {
	return dara.Validate(s)
}
