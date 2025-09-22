// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVLExtractionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *SubmitVLExtractionTaskRequest
	GetFileName() *string
	SetFileUrl(v string) *SubmitVLExtractionTaskRequest
	GetFileUrl() *string
	SetFolderId(v string) *SubmitVLExtractionTaskRequest
	GetFolderId() *string
	SetTemplateId(v string) *SubmitVLExtractionTaskRequest
	GetTemplateId() *string
}

type SubmitVLExtractionTaskRequest struct {
	// The filename must include the file type suffix.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// Choose one of fileUrl or fileUrlObject:
	//
	// - fileUrl: Use by providing the document URL, for a single document (supports up to 1000 pages and 100MB in size)
	//
	// - fileUrlObject: Use when calling the interface with local file upload, for a single document (supports up to 1000 pages and 100 MB in size)
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
	// - Unique knowledge base folder ID, used when you need to categorize documents and control the scope of online Q&A queries.
	//
	// - The folder ID needs to be obtained from the intelligent document console after logging in.
	//
	// example:
	//
	// xxxxx
	FolderId *string `json:"folderId,omitempty" xml:"folderId,omitempty"`
	// Unique parsing template ID, used to specify the key-value pairs to be extracted from the document. You need to configure the template on the template management page and then obtain the corresponding template ID.
	//
	// example:
	//
	// 572d24k0c95a
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s SubmitVLExtractionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitVLExtractionTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitVLExtractionTaskRequest) GetFileName() *string {
	return s.FileName
}

func (s *SubmitVLExtractionTaskRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *SubmitVLExtractionTaskRequest) GetFolderId() *string {
	return s.FolderId
}

func (s *SubmitVLExtractionTaskRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitVLExtractionTaskRequest) SetFileName(v string) *SubmitVLExtractionTaskRequest {
	s.FileName = &v
	return s
}

func (s *SubmitVLExtractionTaskRequest) SetFileUrl(v string) *SubmitVLExtractionTaskRequest {
	s.FileUrl = &v
	return s
}

func (s *SubmitVLExtractionTaskRequest) SetFolderId(v string) *SubmitVLExtractionTaskRequest {
	s.FolderId = &v
	return s
}

func (s *SubmitVLExtractionTaskRequest) SetTemplateId(v string) *SubmitVLExtractionTaskRequest {
	s.TemplateId = &v
	return s
}

func (s *SubmitVLExtractionTaskRequest) Validate() error {
	return dara.Validate(s)
}
