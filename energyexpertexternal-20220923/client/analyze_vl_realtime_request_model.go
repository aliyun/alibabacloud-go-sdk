// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnalyzeVlRealtimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *AnalyzeVlRealtimeRequest
	GetFileName() *string
	SetFileUrl(v string) *AnalyzeVlRealtimeRequest
	GetFileUrl() *string
	SetLanguage(v string) *AnalyzeVlRealtimeRequest
	GetLanguage() *string
	SetTemplateId(v string) *AnalyzeVlRealtimeRequest
	GetTemplateId() *string
}

type AnalyzeVlRealtimeRequest struct {
	// 文件名需带文件类型后缀
	//
	// example:
	//
	// test.png
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// Valid values: fileUrl and fileUrlObject.
	//
	// 	- fileUrl: used as a document URL. A single document with not more than 1,000 pages and whose size does not exceed 100 MB is supported.
	//
	// 	- fileUrlObject: used when the operation is called in local file upload mode. A single document with not more than 1,000 pages and whose size does not exceed 100 MB is supported.
	//
	// > The relationship between file extraction methods and supported document types
	//
	// > - Long text RAG: Supports pdf, doc/docx, xlsx, csv, txt, up to 1000 pages
	//
	// > - Image processing: Supports pdf, jpg, jpeg, png, bmp, jpe, tif, tiff, webp, heic
	//
	// > - Long text understanding: Supports doc/docx, xlsx, pdf, csv, txt
	//
	// example:
	//
	// fileUrl: https://example.com/example.pdf fileUrlObject: FileInputStream generated for a local file
	FileUrl *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// The language, which can be transferred. Valid values:
	//
	// 	- zh-CN (default)
	//
	// 	- en-US
	//
	// example:
	//
	// zh-CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// The unique ID of an extraction template, which is used to specify the content to be extracted from a document. You must log on to the Template Management page to configure the template and then obtain the corresponding template ID.
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

func (s *AnalyzeVlRealtimeRequest) GetFileName() *string {
	return s.FileName
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

func (s *AnalyzeVlRealtimeRequest) SetFileName(v string) *AnalyzeVlRealtimeRequest {
	s.FileName = &v
	return s
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
