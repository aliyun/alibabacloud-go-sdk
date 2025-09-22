// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iAnalyzeVlRealtimeAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileUrlObject(v io.Reader) *AnalyzeVlRealtimeAdvanceRequest
	GetFileUrlObject() io.Reader
	SetLanguage(v string) *AnalyzeVlRealtimeAdvanceRequest
	GetLanguage() *string
	SetTemplateId(v string) *AnalyzeVlRealtimeAdvanceRequest
	GetTemplateId() *string
}

type AnalyzeVlRealtimeAdvanceRequest struct {
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
	FileUrlObject io.Reader `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
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

func (s AnalyzeVlRealtimeAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeVlRealtimeAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AnalyzeVlRealtimeAdvanceRequest) GetFileUrlObject() io.Reader {
	return s.FileUrlObject
}

func (s *AnalyzeVlRealtimeAdvanceRequest) GetLanguage() *string {
	return s.Language
}

func (s *AnalyzeVlRealtimeAdvanceRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *AnalyzeVlRealtimeAdvanceRequest) SetFileUrlObject(v io.Reader) *AnalyzeVlRealtimeAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *AnalyzeVlRealtimeAdvanceRequest) SetLanguage(v string) *AnalyzeVlRealtimeAdvanceRequest {
	s.Language = &v
	return s
}

func (s *AnalyzeVlRealtimeAdvanceRequest) SetTemplateId(v string) *AnalyzeVlRealtimeAdvanceRequest {
	s.TemplateId = &v
	return s
}

func (s *AnalyzeVlRealtimeAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
