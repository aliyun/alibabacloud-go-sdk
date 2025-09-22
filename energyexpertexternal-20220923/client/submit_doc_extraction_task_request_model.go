// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDocExtractionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtractType(v string) *SubmitDocExtractionTaskRequest
	GetExtractType() *string
	SetFileName(v string) *SubmitDocExtractionTaskRequest
	GetFileName() *string
	SetFileUrl(v string) *SubmitDocExtractionTaskRequest
	GetFileUrl() *string
	SetFolderId(v string) *SubmitDocExtractionTaskRequest
	GetFolderId() *string
	SetTemplateId(v string) *SubmitDocExtractionTaskRequest
	GetTemplateId() *string
}

type SubmitDocExtractionTaskRequest struct {
	// Document extraction type:
	//
	// Supports rag and long text understanding types, default is rag.
	//
	// example:
	//
	// rag
	ExtractType *string `json:"extractType,omitempty" xml:"extractType,omitempty"`
	// The filename must include the file type extension.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// Choose one of fileUrl or fileUrlObject:
	//
	// - fileUrl: Use by providing the document URL, for a single document (supports up to 1000 pages, 100MB in size)
	//
	// - fileUrlObject: Use when calling the interface with local file upload, for a single document (supports up to 1000 pages, 100 MB in size)
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
	// fileUrl：https://example.com/example.pdf
	//
	// fileUrlObject：FileInputStream generated from a local file
	FileUrl *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// - A unique knowledge base folder ID, used when you need to categorize documents and control the scope of documents for online Q&A queries.
	//
	// - The folder ID needs to be obtained by logging into the intelligent document console.
	//
	// example:
	//
	// xxxxx
	FolderId *string `json:"folderId,omitempty" xml:"folderId,omitempty"`
	// A unique extraction template ID used to specify the content to be extracted from the document. You need to log in to the template management page to configure the template and obtain the corresponding template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 572d24k0c95a
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s SubmitDocExtractionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocExtractionTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitDocExtractionTaskRequest) GetExtractType() *string {
	return s.ExtractType
}

func (s *SubmitDocExtractionTaskRequest) GetFileName() *string {
	return s.FileName
}

func (s *SubmitDocExtractionTaskRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *SubmitDocExtractionTaskRequest) GetFolderId() *string {
	return s.FolderId
}

func (s *SubmitDocExtractionTaskRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitDocExtractionTaskRequest) SetExtractType(v string) *SubmitDocExtractionTaskRequest {
	s.ExtractType = &v
	return s
}

func (s *SubmitDocExtractionTaskRequest) SetFileName(v string) *SubmitDocExtractionTaskRequest {
	s.FileName = &v
	return s
}

func (s *SubmitDocExtractionTaskRequest) SetFileUrl(v string) *SubmitDocExtractionTaskRequest {
	s.FileUrl = &v
	return s
}

func (s *SubmitDocExtractionTaskRequest) SetFolderId(v string) *SubmitDocExtractionTaskRequest {
	s.FolderId = &v
	return s
}

func (s *SubmitDocExtractionTaskRequest) SetTemplateId(v string) *SubmitDocExtractionTaskRequest {
	s.TemplateId = &v
	return s
}

func (s *SubmitDocExtractionTaskRequest) Validate() error {
	return dara.Validate(s)
}
