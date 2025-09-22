// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDocParsingTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *SubmitDocParsingTaskRequest
	GetFileName() *string
	SetFileUrl(v string) *SubmitDocParsingTaskRequest
	GetFileUrl() *string
	SetFolderId(v string) *SubmitDocParsingTaskRequest
	GetFolderId() *string
	SetNeedAnalyzeImg(v bool) *SubmitDocParsingTaskRequest
	GetNeedAnalyzeImg() *bool
}

type SubmitDocParsingTaskRequest struct {
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
	// - fileUrl: Use by providing the document URL, for a single document (supports up to 1000 pages and 100MB in size)
	//
	// - fileUrlObject: Use when calling the interface with local file upload, for a single document (supports up to 1000 pages and 100 MB in size)
	//
	// > The relationship between file parsing methods and supported document types
	//
	// > - Long Text RAG: Supports pdf, doc/docx, supports up to 1000 pages
	//
	// > - Image Processing: Supports pdf, jpg, jpeg, png, bmp
	//
	// > - Long Text Understanding: Supports pdf, doc/docx, xls/xlsx
	//
	// example:
	//
	// fileUrl：https://example.com/example.pdf
	//
	// fileUrlObject：FileInputStream generated from a local file
	FileUrl *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// - Unique knowledge base folder ID, used when categorizing documents and controlling the scope of documents for online Q&A queries.
	//
	// - The folder ID needs to be obtained from the Intelligent Document Console after logging in.
	//
	// example:
	//
	// xxxxx
	FolderId *string `json:"folderId,omitempty" xml:"folderId,omitempty"`
	// Whether to parse image content within the document.
	//
	// example:
	//
	// false
	NeedAnalyzeImg *bool `json:"needAnalyzeImg,omitempty" xml:"needAnalyzeImg,omitempty"`
}

func (s SubmitDocParsingTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocParsingTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitDocParsingTaskRequest) GetFileName() *string {
	return s.FileName
}

func (s *SubmitDocParsingTaskRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *SubmitDocParsingTaskRequest) GetFolderId() *string {
	return s.FolderId
}

func (s *SubmitDocParsingTaskRequest) GetNeedAnalyzeImg() *bool {
	return s.NeedAnalyzeImg
}

func (s *SubmitDocParsingTaskRequest) SetFileName(v string) *SubmitDocParsingTaskRequest {
	s.FileName = &v
	return s
}

func (s *SubmitDocParsingTaskRequest) SetFileUrl(v string) *SubmitDocParsingTaskRequest {
	s.FileUrl = &v
	return s
}

func (s *SubmitDocParsingTaskRequest) SetFolderId(v string) *SubmitDocParsingTaskRequest {
	s.FolderId = &v
	return s
}

func (s *SubmitDocParsingTaskRequest) SetNeedAnalyzeImg(v bool) *SubmitDocParsingTaskRequest {
	s.NeedAnalyzeImg = &v
	return s
}

func (s *SubmitDocParsingTaskRequest) Validate() error {
	return dara.Validate(s)
}
