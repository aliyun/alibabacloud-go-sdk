// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iSubmitVLExtractionTaskAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *SubmitVLExtractionTaskAdvanceRequest
	GetFileName() *string
	SetFileUrlObject(v io.Reader) *SubmitVLExtractionTaskAdvanceRequest
	GetFileUrlObject() io.Reader
	SetFolderId(v string) *SubmitVLExtractionTaskAdvanceRequest
	GetFolderId() *string
	SetTemplateId(v string) *SubmitVLExtractionTaskAdvanceRequest
	GetTemplateId() *string
}

type SubmitVLExtractionTaskAdvanceRequest struct {
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
	FileUrlObject io.Reader `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
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

func (s SubmitVLExtractionTaskAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitVLExtractionTaskAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitVLExtractionTaskAdvanceRequest) GetFileName() *string {
	return s.FileName
}

func (s *SubmitVLExtractionTaskAdvanceRequest) GetFileUrlObject() io.Reader {
	return s.FileUrlObject
}

func (s *SubmitVLExtractionTaskAdvanceRequest) GetFolderId() *string {
	return s.FolderId
}

func (s *SubmitVLExtractionTaskAdvanceRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitVLExtractionTaskAdvanceRequest) SetFileName(v string) *SubmitVLExtractionTaskAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitVLExtractionTaskAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitVLExtractionTaskAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *SubmitVLExtractionTaskAdvanceRequest) SetFolderId(v string) *SubmitVLExtractionTaskAdvanceRequest {
	s.FolderId = &v
	return s
}

func (s *SubmitVLExtractionTaskAdvanceRequest) SetTemplateId(v string) *SubmitVLExtractionTaskAdvanceRequest {
	s.TemplateId = &v
	return s
}

func (s *SubmitVLExtractionTaskAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
