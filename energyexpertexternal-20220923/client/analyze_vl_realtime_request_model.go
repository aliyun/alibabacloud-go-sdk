// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnalyzeVlRealtimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileUrl(v string) *AnalyzeVlRealtimeRequest
	GetFileUrl() *string
	SetLanguage(v string) *AnalyzeVlRealtimeRequest
	GetLanguage() *string
	SetTemplateId(v string) *AnalyzeVlRealtimeRequest
	GetTemplateId() *string
}

type AnalyzeVlRealtimeRequest struct {
	// Choose one of fileUrl or fileUrlObject:
	//
	// - fileUrl: Use in the form of a document URL, for a single document (supports up to 1000 pages and 100MB)
	//
	// - fileUrlObject: Use when calling the interface with local file upload, for a single document (supports up to 1000 pages and 100 MB)
	//
	// > The relationship between file parsing methods and supported document types
	//
	// > - Long Text RAG: Supports pdf, doc/docx, up to 1000 pages
	//
	// > - Image Processing: Supports pdf, jpg, jpeg, png, bmp
	//
	// > - Long Text Understanding: Supports pdf, doc/docx, xls/xlsx
	//
	// example:
	//
	// fileUrl：https://example.com/example.pdf
	//
	// fileUrlObject：本地文件生成的FileInputStream
	FileUrl *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// Language, parameters that can be passed
	//
	// - zh-CN: Chinese (default)
	//
	// - en-US: English
	//
	// example:
	//
	// zh-CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// A unique parsing template ID used to specify the key-value pairs to be extracted from the document. You need to log in to the template management page, configure the template, and then get the corresponding template ID.
	//
	// example:
	//
	// 572d24k0c95a
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s AnalyzeVlRealtimeRequest) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeVlRealtimeRequest) GoString() string {
	return s.String()
}

func (s *AnalyzeVlRealtimeRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *AnalyzeVlRealtimeRequest) GetLanguage() *string {
	return s.Language
}

func (s *AnalyzeVlRealtimeRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *AnalyzeVlRealtimeRequest) SetFileUrl(v string) *AnalyzeVlRealtimeRequest {
	s.FileUrl = &v
	return s
}

func (s *AnalyzeVlRealtimeRequest) SetLanguage(v string) *AnalyzeVlRealtimeRequest {
	s.Language = &v
	return s
}

func (s *AnalyzeVlRealtimeRequest) SetTemplateId(v string) *AnalyzeVlRealtimeRequest {
	s.TemplateId = &v
	return s
}

func (s *AnalyzeVlRealtimeRequest) Validate() error {
	return dara.Validate(s)
}
