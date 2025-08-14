// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDocParserJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *SubmitDocParserJobRequest
	GetFileName() *string
	SetFileNameExtension(v string) *SubmitDocParserJobRequest
	GetFileNameExtension() *string
	SetFileUrl(v string) *SubmitDocParserJobRequest
	GetFileUrl() *string
	SetFormulaEnhancement(v bool) *SubmitDocParserJobRequest
	GetFormulaEnhancement() *bool
	SetLlmEnhancement(v bool) *SubmitDocParserJobRequest
	GetLlmEnhancement() *bool
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
	FileUrl            *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	FormulaEnhancement *bool   `json:"FormulaEnhancement,omitempty" xml:"FormulaEnhancement,omitempty"`
	LlmEnhancement     *bool   `json:"LlmEnhancement,omitempty" xml:"LlmEnhancement,omitempty"`
	Option             *string `json:"Option,omitempty" xml:"Option,omitempty"`
	OssBucket          *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint        *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	OutputHtmlTable    *bool   `json:"OutputHtmlTable,omitempty" xml:"OutputHtmlTable,omitempty"`
	PageIndex          *string `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
}

func (s SubmitDocParserJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocParserJobRequest) GoString() string {
	return s.String()
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

func (s *SubmitDocParserJobRequest) GetLlmEnhancement() *bool {
	return s.LlmEnhancement
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

func (s *SubmitDocParserJobRequest) SetLlmEnhancement(v bool) *SubmitDocParserJobRequest {
	s.LlmEnhancement = &v
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
