// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iSubmitDocumentAnalyzeJobAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnalysisType(v string) *SubmitDocumentAnalyzeJobAdvanceRequest
	GetAnalysisType() *string
	SetFileName(v string) *SubmitDocumentAnalyzeJobAdvanceRequest
	GetFileName() *string
	SetFileUrlObject(v io.Reader) *SubmitDocumentAnalyzeJobAdvanceRequest
	GetFileUrlObject() io.Reader
	SetFolderId(v string) *SubmitDocumentAnalyzeJobAdvanceRequest
	GetFolderId() *string
	SetTemplateId(v string) *SubmitDocumentAnalyzeJobAdvanceRequest
	GetTemplateId() *string
}

type SubmitDocumentAnalyzeJobAdvanceRequest struct {
	// The default extraction method is "doc", with the following optional values:
	//
	// - vl: Image processing
	//
	// - doc: Long text RAG (Retrieval-Augmented Generation)
	//
	// - docUnderstanding: Long text comprehension
	//
	// - recommender: Recommendation type
	//
	// example:
	//
	// doc
	AnalysisType *string `json:"analysisType,omitempty" xml:"analysisType,omitempty"`
	// The filename must include the file type extension.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// Choose one between fileUrl and fileUrlObject:
	//
	// - fileUrl: Use the document URL method for a single document (supports documents with up to 1000 pages and within 100MB).
	//
	// - fileUrlObject: Use when calling the API via local file upload, for a single document (supports documents with up to 1000 pages and
	//
	// within 100MB).
	//
	// > Relationship between file parsing methods and supported document types.
	//
	// >- Long Text RAG: Supports pdf, doc/docx, and up to 1000 pages
	//
	// >- Image Processing: Supports pdf, jpg, jpeg, png, bmp
	//
	// >- Long Text Understanding: Supports pdf, doc/docx, xls/xlsx
	//
	// example:
	//
	// https://example.com/example.pdf
	FileUrlObject io.Reader `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// Unique knowledge base folder ID, used for categorizing documents and controlling the scope of online Q&A queries. If empty, the document will be uploaded to the tenant\\"s root directory.
	//
	// example:
	//
	// folderCode
	FolderId *string `json:"folderId,omitempty" xml:"folderId,omitempty"`
	// The unique extraction template ID is used to specify the key-value pairs to be extracted from the document. You need to log in to the template management page to configure the template and obtain the corresponding template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// templateCode
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s SubmitDocumentAnalyzeJobAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocumentAnalyzeJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitDocumentAnalyzeJobAdvanceRequest) GetAnalysisType() *string {
	return s.AnalysisType
}

func (s *SubmitDocumentAnalyzeJobAdvanceRequest) GetFileName() *string {
	return s.FileName
}

func (s *SubmitDocumentAnalyzeJobAdvanceRequest) GetFileUrlObject() io.Reader {
	return s.FileUrlObject
}

func (s *SubmitDocumentAnalyzeJobAdvanceRequest) GetFolderId() *string {
	return s.FolderId
}

func (s *SubmitDocumentAnalyzeJobAdvanceRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitDocumentAnalyzeJobAdvanceRequest) SetAnalysisType(v string) *SubmitDocumentAnalyzeJobAdvanceRequest {
	s.AnalysisType = &v
	return s
}

func (s *SubmitDocumentAnalyzeJobAdvanceRequest) SetFileName(v string) *SubmitDocumentAnalyzeJobAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitDocumentAnalyzeJobAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitDocumentAnalyzeJobAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *SubmitDocumentAnalyzeJobAdvanceRequest) SetFolderId(v string) *SubmitDocumentAnalyzeJobAdvanceRequest {
	s.FolderId = &v
	return s
}

func (s *SubmitDocumentAnalyzeJobAdvanceRequest) SetTemplateId(v string) *SubmitDocumentAnalyzeJobAdvanceRequest {
	s.TemplateId = &v
	return s
}

func (s *SubmitDocumentAnalyzeJobAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
