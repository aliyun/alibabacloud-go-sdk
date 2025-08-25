// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDocParserJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnhancementMode(v string) *SubmitDocParserJobShrinkRequest
	GetEnhancementMode() *string
	SetFileName(v string) *SubmitDocParserJobShrinkRequest
	GetFileName() *string
	SetFileNameExtension(v string) *SubmitDocParserJobShrinkRequest
	GetFileNameExtension() *string
	SetFileUrl(v string) *SubmitDocParserJobShrinkRequest
	GetFileUrl() *string
	SetFormulaEnhancement(v bool) *SubmitDocParserJobShrinkRequest
	GetFormulaEnhancement() *bool
	SetLlmEnhancement(v bool) *SubmitDocParserJobShrinkRequest
	GetLlmEnhancement() *bool
	SetMultimediaParametersShrink(v string) *SubmitDocParserJobShrinkRequest
	GetMultimediaParametersShrink() *string
	SetOption(v string) *SubmitDocParserJobShrinkRequest
	GetOption() *string
	SetOssBucket(v string) *SubmitDocParserJobShrinkRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *SubmitDocParserJobShrinkRequest
	GetOssEndpoint() *string
	SetOutputHtmlTable(v bool) *SubmitDocParserJobShrinkRequest
	GetOutputHtmlTable() *bool
	SetPageIndex(v string) *SubmitDocParserJobShrinkRequest
	GetPageIndex() *string
}

type SubmitDocParserJobShrinkRequest struct {
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
	FileUrl                    *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	FormulaEnhancement         *bool   `json:"FormulaEnhancement,omitempty" xml:"FormulaEnhancement,omitempty"`
	LlmEnhancement             *bool   `json:"LlmEnhancement,omitempty" xml:"LlmEnhancement,omitempty"`
	MultimediaParametersShrink *string `json:"MultimediaParameters,omitempty" xml:"MultimediaParameters,omitempty"`
	Option                     *string `json:"Option,omitempty" xml:"Option,omitempty"`
	OssBucket                  *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint                *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	OutputHtmlTable            *bool   `json:"OutputHtmlTable,omitempty" xml:"OutputHtmlTable,omitempty"`
	PageIndex                  *string `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
}

func (s SubmitDocParserJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocParserJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitDocParserJobShrinkRequest) GetEnhancementMode() *string {
	return s.EnhancementMode
}

func (s *SubmitDocParserJobShrinkRequest) GetFileName() *string {
	return s.FileName
}

func (s *SubmitDocParserJobShrinkRequest) GetFileNameExtension() *string {
	return s.FileNameExtension
}

func (s *SubmitDocParserJobShrinkRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *SubmitDocParserJobShrinkRequest) GetFormulaEnhancement() *bool {
	return s.FormulaEnhancement
}

func (s *SubmitDocParserJobShrinkRequest) GetLlmEnhancement() *bool {
	return s.LlmEnhancement
}

func (s *SubmitDocParserJobShrinkRequest) GetMultimediaParametersShrink() *string {
	return s.MultimediaParametersShrink
}

func (s *SubmitDocParserJobShrinkRequest) GetOption() *string {
	return s.Option
}

func (s *SubmitDocParserJobShrinkRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *SubmitDocParserJobShrinkRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *SubmitDocParserJobShrinkRequest) GetOutputHtmlTable() *bool {
	return s.OutputHtmlTable
}

func (s *SubmitDocParserJobShrinkRequest) GetPageIndex() *string {
	return s.PageIndex
}

func (s *SubmitDocParserJobShrinkRequest) SetEnhancementMode(v string) *SubmitDocParserJobShrinkRequest {
	s.EnhancementMode = &v
	return s
}

func (s *SubmitDocParserJobShrinkRequest) SetFileName(v string) *SubmitDocParserJobShrinkRequest {
	s.FileName = &v
	return s
}

func (s *SubmitDocParserJobShrinkRequest) SetFileNameExtension(v string) *SubmitDocParserJobShrinkRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitDocParserJobShrinkRequest) SetFileUrl(v string) *SubmitDocParserJobShrinkRequest {
	s.FileUrl = &v
	return s
}

func (s *SubmitDocParserJobShrinkRequest) SetFormulaEnhancement(v bool) *SubmitDocParserJobShrinkRequest {
	s.FormulaEnhancement = &v
	return s
}

func (s *SubmitDocParserJobShrinkRequest) SetLlmEnhancement(v bool) *SubmitDocParserJobShrinkRequest {
	s.LlmEnhancement = &v
	return s
}

func (s *SubmitDocParserJobShrinkRequest) SetMultimediaParametersShrink(v string) *SubmitDocParserJobShrinkRequest {
	s.MultimediaParametersShrink = &v
	return s
}

func (s *SubmitDocParserJobShrinkRequest) SetOption(v string) *SubmitDocParserJobShrinkRequest {
	s.Option = &v
	return s
}

func (s *SubmitDocParserJobShrinkRequest) SetOssBucket(v string) *SubmitDocParserJobShrinkRequest {
	s.OssBucket = &v
	return s
}

func (s *SubmitDocParserJobShrinkRequest) SetOssEndpoint(v string) *SubmitDocParserJobShrinkRequest {
	s.OssEndpoint = &v
	return s
}

func (s *SubmitDocParserJobShrinkRequest) SetOutputHtmlTable(v bool) *SubmitDocParserJobShrinkRequest {
	s.OutputHtmlTable = &v
	return s
}

func (s *SubmitDocParserJobShrinkRequest) SetPageIndex(v string) *SubmitDocParserJobShrinkRequest {
	s.PageIndex = &v
	return s
}

func (s *SubmitDocParserJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
