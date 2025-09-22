// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iSubmitDocParsingTaskAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *SubmitDocParsingTaskAdvanceRequest
	GetFileName() *string
	SetFileUrlObject(v io.Reader) *SubmitDocParsingTaskAdvanceRequest
	GetFileUrlObject() io.Reader
	SetFolderId(v string) *SubmitDocParsingTaskAdvanceRequest
	GetFolderId() *string
	SetNeedAnalyzeImg(v bool) *SubmitDocParsingTaskAdvanceRequest
	GetNeedAnalyzeImg() *bool
}

type SubmitDocParsingTaskAdvanceRequest struct {
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
	FileUrlObject io.Reader `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
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

func (s SubmitDocParsingTaskAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocParsingTaskAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitDocParsingTaskAdvanceRequest) GetFileName() *string {
	return s.FileName
}

func (s *SubmitDocParsingTaskAdvanceRequest) GetFileUrlObject() io.Reader {
	return s.FileUrlObject
}

func (s *SubmitDocParsingTaskAdvanceRequest) GetFolderId() *string {
	return s.FolderId
}

func (s *SubmitDocParsingTaskAdvanceRequest) GetNeedAnalyzeImg() *bool {
	return s.NeedAnalyzeImg
}

func (s *SubmitDocParsingTaskAdvanceRequest) SetFileName(v string) *SubmitDocParsingTaskAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitDocParsingTaskAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitDocParsingTaskAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *SubmitDocParsingTaskAdvanceRequest) SetFolderId(v string) *SubmitDocParsingTaskAdvanceRequest {
	s.FolderId = &v
	return s
}

func (s *SubmitDocParsingTaskAdvanceRequest) SetNeedAnalyzeImg(v bool) *SubmitDocParsingTaskAdvanceRequest {
	s.NeedAnalyzeImg = &v
	return s
}

func (s *SubmitDocParsingTaskAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
