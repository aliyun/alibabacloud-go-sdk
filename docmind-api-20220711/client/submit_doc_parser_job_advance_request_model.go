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
	SetLlmEnhancement(v bool) *SubmitDocParserJobAdvanceRequest
	GetLlmEnhancement() *bool
	SetMultimediaParameters(v *SubmitDocParserJobAdvanceRequestMultimediaParameters) *SubmitDocParserJobAdvanceRequest
	GetMultimediaParameters() *SubmitDocParserJobAdvanceRequestMultimediaParameters
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
	EnhancementMode *string `json:"EnhancementMode,omitempty" xml:"EnhancementMode,omitempty"`
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
	LlmEnhancement       *bool                                                 `json:"LlmEnhancement,omitempty" xml:"LlmEnhancement,omitempty"`
	MultimediaParameters *SubmitDocParserJobAdvanceRequestMultimediaParameters `json:"MultimediaParameters,omitempty" xml:"MultimediaParameters,omitempty" type:"Struct"`
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

func (s *SubmitDocParserJobAdvanceRequest) GetLlmEnhancement() *bool {
	return s.LlmEnhancement
}

func (s *SubmitDocParserJobAdvanceRequest) GetMultimediaParameters() *SubmitDocParserJobAdvanceRequestMultimediaParameters {
	return s.MultimediaParameters
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

func (s *SubmitDocParserJobAdvanceRequest) SetLlmEnhancement(v bool) *SubmitDocParserJobAdvanceRequest {
	s.LlmEnhancement = &v
	return s
}

func (s *SubmitDocParserJobAdvanceRequest) SetMultimediaParameters(v *SubmitDocParserJobAdvanceRequestMultimediaParameters) *SubmitDocParserJobAdvanceRequest {
	s.MultimediaParameters = v
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
