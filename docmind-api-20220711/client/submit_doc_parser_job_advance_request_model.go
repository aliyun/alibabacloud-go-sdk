// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iSubmitDocParserJobAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomOssConfig(v *SubmitDocParserJobAdvanceRequestCustomOssConfig) *SubmitDocParserJobAdvanceRequest
	GetCustomOssConfig() *SubmitDocParserJobAdvanceRequestCustomOssConfig
	SetEnhancementMode(v string) *SubmitDocParserJobAdvanceRequest
	GetEnhancementMode() *string
	SetFileName(v string) *SubmitDocParserJobAdvanceRequest
	GetFileName() *string
	SetFileNameExtension(v string) *SubmitDocParserJobAdvanceRequest
	GetFileNameExtension() *string
	SetFileUrlObject(v io.Reader) *SubmitDocParserJobAdvanceRequest
	GetFileUrlObject() io.Reader
	SetFormulaEnhancement(v bool) *SubmitDocParserJobAdvanceRequest
	GetFormulaEnhancement() *bool
	SetLLMParam(v *SubmitDocParserJobAdvanceRequestLLMParam) *SubmitDocParserJobAdvanceRequest
	GetLLMParam() *SubmitDocParserJobAdvanceRequestLLMParam
	SetLlmEnhancement(v bool) *SubmitDocParserJobAdvanceRequest
	GetLlmEnhancement() *bool
	SetMultimediaParameters(v *SubmitDocParserJobAdvanceRequestMultimediaParameters) *SubmitDocParserJobAdvanceRequest
	GetMultimediaParameters() *SubmitDocParserJobAdvanceRequestMultimediaParameters
	SetNeedHeaderFooter(v bool) *SubmitDocParserJobAdvanceRequest
	GetNeedHeaderFooter() *bool
	SetOption(v string) *SubmitDocParserJobAdvanceRequest
	GetOption() *string
	SetOssBucket(v string) *SubmitDocParserJobAdvanceRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *SubmitDocParserJobAdvanceRequest
	GetOssEndpoint() *string
	SetOutputHtmlTable(v bool) *SubmitDocParserJobAdvanceRequest
	GetOutputHtmlTable() *bool
	SetPageIndex(v string) *SubmitDocParserJobAdvanceRequest
	GetPageIndex() *string
}

type SubmitDocParserJobAdvanceRequest struct {
	CustomOssConfig *SubmitDocParserJobAdvanceRequestCustomOssConfig `json:"CustomOssConfig,omitempty" xml:"CustomOssConfig,omitempty" type:"Struct"`
	EnhancementMode *string                                          `json:"EnhancementMode,omitempty" xml:"EnhancementMode,omitempty"`
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
	FileUrlObject        io.Reader                                             `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	FormulaEnhancement   *bool                                                 `json:"FormulaEnhancement,omitempty" xml:"FormulaEnhancement,omitempty"`
	LLMParam             *SubmitDocParserJobAdvanceRequestLLMParam             `json:"LLMParam,omitempty" xml:"LLMParam,omitempty" type:"Struct"`
	LlmEnhancement       *bool                                                 `json:"LlmEnhancement,omitempty" xml:"LlmEnhancement,omitempty"`
	MultimediaParameters *SubmitDocParserJobAdvanceRequestMultimediaParameters `json:"MultimediaParameters,omitempty" xml:"MultimediaParameters,omitempty" type:"Struct"`
	NeedHeaderFooter     *bool                                                 `json:"NeedHeaderFooter,omitempty" xml:"NeedHeaderFooter,omitempty"`
	Option               *string                                               `json:"Option,omitempty" xml:"Option,omitempty"`
	OssBucket            *string                                               `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint          *string                                               `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	OutputHtmlTable      *bool                                                 `json:"OutputHtmlTable,omitempty" xml:"OutputHtmlTable,omitempty"`
	PageIndex            *string                                               `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
}

func (s SubmitDocParserJobAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocParserJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitDocParserJobAdvanceRequest) GetCustomOssConfig() *SubmitDocParserJobAdvanceRequestCustomOssConfig {
	return s.CustomOssConfig
}

func (s *SubmitDocParserJobAdvanceRequest) GetEnhancementMode() *string {
	return s.EnhancementMode
}

func (s *SubmitDocParserJobAdvanceRequest) GetFileName() *string {
	return s.FileName
}

func (s *SubmitDocParserJobAdvanceRequest) GetFileNameExtension() *string {
	return s.FileNameExtension
}

func (s *SubmitDocParserJobAdvanceRequest) GetFileUrlObject() io.Reader {
	return s.FileUrlObject
}

func (s *SubmitDocParserJobAdvanceRequest) GetFormulaEnhancement() *bool {
	return s.FormulaEnhancement
}

func (s *SubmitDocParserJobAdvanceRequest) GetLLMParam() *SubmitDocParserJobAdvanceRequestLLMParam {
	return s.LLMParam
}

func (s *SubmitDocParserJobAdvanceRequest) GetLlmEnhancement() *bool {
	return s.LlmEnhancement
}

func (s *SubmitDocParserJobAdvanceRequest) GetMultimediaParameters() *SubmitDocParserJobAdvanceRequestMultimediaParameters {
	return s.MultimediaParameters
}

func (s *SubmitDocParserJobAdvanceRequest) GetNeedHeaderFooter() *bool {
	return s.NeedHeaderFooter
}

func (s *SubmitDocParserJobAdvanceRequest) GetOption() *string {
	return s.Option
}

func (s *SubmitDocParserJobAdvanceRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *SubmitDocParserJobAdvanceRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *SubmitDocParserJobAdvanceRequest) GetOutputHtmlTable() *bool {
	return s.OutputHtmlTable
}

func (s *SubmitDocParserJobAdvanceRequest) GetPageIndex() *string {
	return s.PageIndex
}

func (s *SubmitDocParserJobAdvanceRequest) SetCustomOssConfig(v *SubmitDocParserJobAdvanceRequestCustomOssConfig) *SubmitDocParserJobAdvanceRequest {
	s.CustomOssConfig = v
	return s
}

func (s *SubmitDocParserJobAdvanceRequest) SetEnhancementMode(v string) *SubmitDocParserJobAdvanceRequest {
	s.EnhancementMode = &v
	return s
}

func (s *SubmitDocParserJobAdvanceRequest) SetFileName(v string) *SubmitDocParserJobAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitDocParserJobAdvanceRequest) SetFileNameExtension(v string) *SubmitDocParserJobAdvanceRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitDocParserJobAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitDocParserJobAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *SubmitDocParserJobAdvanceRequest) SetFormulaEnhancement(v bool) *SubmitDocParserJobAdvanceRequest {
	s.FormulaEnhancement = &v
	return s
}

func (s *SubmitDocParserJobAdvanceRequest) SetLLMParam(v *SubmitDocParserJobAdvanceRequestLLMParam) *SubmitDocParserJobAdvanceRequest {
	s.LLMParam = v
	return s
}

func (s *SubmitDocParserJobAdvanceRequest) SetLlmEnhancement(v bool) *SubmitDocParserJobAdvanceRequest {
	s.LlmEnhancement = &v
	return s
}

func (s *SubmitDocParserJobAdvanceRequest) SetMultimediaParameters(v *SubmitDocParserJobAdvanceRequestMultimediaParameters) *SubmitDocParserJobAdvanceRequest {
	s.MultimediaParameters = v
	return s
}

func (s *SubmitDocParserJobAdvanceRequest) SetNeedHeaderFooter(v bool) *SubmitDocParserJobAdvanceRequest {
	s.NeedHeaderFooter = &v
	return s
}

func (s *SubmitDocParserJobAdvanceRequest) SetOption(v string) *SubmitDocParserJobAdvanceRequest {
	s.Option = &v
	return s
}

func (s *SubmitDocParserJobAdvanceRequest) SetOssBucket(v string) *SubmitDocParserJobAdvanceRequest {
	s.OssBucket = &v
	return s
}

func (s *SubmitDocParserJobAdvanceRequest) SetOssEndpoint(v string) *SubmitDocParserJobAdvanceRequest {
	s.OssEndpoint = &v
	return s
}

func (s *SubmitDocParserJobAdvanceRequest) SetOutputHtmlTable(v bool) *SubmitDocParserJobAdvanceRequest {
	s.OutputHtmlTable = &v
	return s
}

func (s *SubmitDocParserJobAdvanceRequest) SetPageIndex(v string) *SubmitDocParserJobAdvanceRequest {
	s.PageIndex = &v
	return s
}

func (s *SubmitDocParserJobAdvanceRequest) Validate() error {
	if s.CustomOssConfig != nil {
		if err := s.CustomOssConfig.Validate(); err != nil {
			return err
		}
	}
	if s.LLMParam != nil {
		if err := s.LLMParam.Validate(); err != nil {
			return err
		}
	}
	if s.MultimediaParameters != nil {
		if err := s.MultimediaParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitDocParserJobAdvanceRequestCustomOssConfig struct {
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

func (s SubmitDocParserJobAdvanceRequestCustomOssConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocParserJobAdvanceRequestCustomOssConfig) GoString() string {
	return s.String()
}

func (s *SubmitDocParserJobAdvanceRequestCustomOssConfig) GetAccessId() *string {
	return s.AccessId
}

func (s *SubmitDocParserJobAdvanceRequestCustomOssConfig) GetAccessKeySecret() *string {
	return s.AccessKeySecret
}

func (s *SubmitDocParserJobAdvanceRequestCustomOssConfig) GetStsToken() *string {
	return s.StsToken
}

func (s *SubmitDocParserJobAdvanceRequestCustomOssConfig) SetAccessId(v string) *SubmitDocParserJobAdvanceRequestCustomOssConfig {
	s.AccessId = &v
	return s
}

func (s *SubmitDocParserJobAdvanceRequestCustomOssConfig) SetAccessKeySecret(v string) *SubmitDocParserJobAdvanceRequestCustomOssConfig {
	s.AccessKeySecret = &v
	return s
}

func (s *SubmitDocParserJobAdvanceRequestCustomOssConfig) SetStsToken(v string) *SubmitDocParserJobAdvanceRequestCustomOssConfig {
	s.StsToken = &v
	return s
}

func (s *SubmitDocParserJobAdvanceRequestCustomOssConfig) Validate() error {
	return dara.Validate(s)
}

type SubmitDocParserJobAdvanceRequestLLMParam struct {
	// example:
	//
	// qwen-vl-ocr-latest
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// Read all the text from the image.
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
}

func (s SubmitDocParserJobAdvanceRequestLLMParam) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocParserJobAdvanceRequestLLMParam) GoString() string {
	return s.String()
}

func (s *SubmitDocParserJobAdvanceRequestLLMParam) GetModel() *string {
	return s.Model
}

func (s *SubmitDocParserJobAdvanceRequestLLMParam) GetPrompt() *string {
	return s.Prompt
}

func (s *SubmitDocParserJobAdvanceRequestLLMParam) SetModel(v string) *SubmitDocParserJobAdvanceRequestLLMParam {
	s.Model = &v
	return s
}

func (s *SubmitDocParserJobAdvanceRequestLLMParam) SetPrompt(v string) *SubmitDocParserJobAdvanceRequestLLMParam {
	s.Prompt = &v
	return s
}

func (s *SubmitDocParserJobAdvanceRequestLLMParam) Validate() error {
	return dara.Validate(s)
}

type SubmitDocParserJobAdvanceRequestMultimediaParameters struct {
	EnableSynopsisParse *bool   `json:"EnableSynopsisParse,omitempty" xml:"EnableSynopsisParse,omitempty"`
	VlParsePrompt       *string `json:"VlParsePrompt,omitempty" xml:"VlParsePrompt,omitempty"`
}

func (s SubmitDocParserJobAdvanceRequestMultimediaParameters) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocParserJobAdvanceRequestMultimediaParameters) GoString() string {
	return s.String()
}

func (s *SubmitDocParserJobAdvanceRequestMultimediaParameters) GetEnableSynopsisParse() *bool {
	return s.EnableSynopsisParse
}

func (s *SubmitDocParserJobAdvanceRequestMultimediaParameters) GetVlParsePrompt() *string {
	return s.VlParsePrompt
}

func (s *SubmitDocParserJobAdvanceRequestMultimediaParameters) SetEnableSynopsisParse(v bool) *SubmitDocParserJobAdvanceRequestMultimediaParameters {
	s.EnableSynopsisParse = &v
	return s
}

func (s *SubmitDocParserJobAdvanceRequestMultimediaParameters) SetVlParsePrompt(v string) *SubmitDocParserJobAdvanceRequestMultimediaParameters {
	s.VlParsePrompt = &v
	return s
}

func (s *SubmitDocParserJobAdvanceRequestMultimediaParameters) Validate() error {
	return dara.Validate(s)
}
