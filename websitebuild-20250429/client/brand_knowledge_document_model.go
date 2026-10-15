// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBrandKnowledgeDocument interface {
	dara.Model
	String() string
	GoString() string
	SetContentSufficient(v bool) *BrandKnowledgeDocument
	GetContentSufficient() *bool
	SetContentType(v string) *BrandKnowledgeDocument
	GetContentType() *string
	SetCreatedAt(v int64) *BrandKnowledgeDocument
	GetCreatedAt() *int64
	SetDeletable(v bool) *BrandKnowledgeDocument
	GetDeletable() *bool
	SetDisplayTitle(v string) *BrandKnowledgeDocument
	GetDisplayTitle() *string
	SetDocCategory(v string) *BrandKnowledgeDocument
	GetDocCategory() *string
	SetDocumentId(v string) *BrandKnowledgeDocument
	GetDocumentId() *string
	SetErrorCode(v string) *BrandKnowledgeDocument
	GetErrorCode() *string
	SetErrorMessage(v string) *BrandKnowledgeDocument
	GetErrorMessage() *string
	SetFileSize(v int64) *BrandKnowledgeDocument
	GetFileSize() *int64
	SetKeyTopics(v []*string) *BrandKnowledgeDocument
	GetKeyTopics() []*string
	SetOneLineSummary(v string) *BrandKnowledgeDocument
	GetOneLineSummary() *string
	SetOriginalName(v string) *BrandKnowledgeDocument
	GetOriginalName() *string
	SetProjectId(v string) *BrandKnowledgeDocument
	GetProjectId() *string
	SetReadyAt(v int64) *BrandKnowledgeDocument
	GetReadyAt() *int64
	SetStatus(v string) *BrandKnowledgeDocument
	GetStatus() *string
	SetStatusMessage(v string) *BrandKnowledgeDocument
	GetStatusMessage() *string
	SetSummaryErrorCode(v string) *BrandKnowledgeDocument
	GetSummaryErrorCode() *string
	SetSummaryStatus(v string) *BrandKnowledgeDocument
	GetSummaryStatus() *string
	SetUpdatedAt(v int64) *BrandKnowledgeDocument
	GetUpdatedAt() *int64
}

type BrandKnowledgeDocument struct {
	// Specifies whether the content is sufficient to be considered for plan summary candidates.
	ContentSufficient *bool `json:"ContentSufficient,omitempty" xml:"ContentSufficient,omitempty"`
	// The MIME type of the file.
	//
	// example:
	//
	// application/pdf
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The creation time in Unix milliseconds.
	//
	// example:
	//
	// 1735660800000
	CreatedAt *int64 `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// Specifies whether deletion is currently allowed.
	Deletable *bool `json:"Deletable,omitempty" xml:"Deletable,omitempty"`
	// The display title.
	//
	// example:
	//
	// Wanxiaozhi product introduction
	DisplayTitle *string `json:"DisplayTitle,omitempty" xml:"DisplayTitle,omitempty"`
	// The document category.
	//
	// example:
	//
	// PRODUCT
	DocCategory *string `json:"DocCategory,omitempty" xml:"DocCategory,omitempty"`
	// The ID of the brand knowledge document.
	//
	// example:
	//
	// bkd_3f7a1c2e9b7d4a5f8c0e1d2b3a4c5d6e
	DocumentId *string `json:"DocumentId,omitempty" xml:"DocumentId,omitempty"`
	// The stable error code for main processing failures.
	//
	// example:
	//
	// BRAND_KNOWLEDGE_PARSE_FAILED
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The desensitized error message for main processing failures.
	//
	// example:
	//
	// Document parsing failed. Please re-upload
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The file size in bytes.
	//
	// example:
	//
	// 1048576
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The summary keywords.
	KeyTopics []*string `json:"KeyTopics,omitempty" xml:"KeyTopics,omitempty" type:"Repeated"`
	// The one-line summary, which is used only to construct search queries and cannot be used as factual evidence.
	//
	// example:
	//
	// Introduces the product positioning and core capabilities of the Wanxiaozhi AI website building platform. It can be used as a reference for content topic selection
	OneLineSummary *string `json:"OneLineSummary,omitempty" xml:"OneLineSummary,omitempty"`
	// The original file name.
	//
	// example:
	//
	// Wanxiaozhi product introduction.pdf
	OriginalName *string `json:"OriginalName,omitempty" xml:"OriginalName,omitempty"`
	// The ID of the business workspace. For GEO business workspaces of the brand marketing type, this is the appWorkspaceId.
	//
	// example:
	//
	// d4e5f60718293a4b5c6d7e8f90a1b2c3
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The time when the document is ready for reference, in Unix milliseconds.
	//
	// example:
	//
	// 1735660900000
	ReadyAt *int64 `json:"ReadyAt,omitempty" xml:"ReadyAt,omitempty"`
	// The document processing status.
	//
	// example:
	//
	// READY
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The message displayed on the status page.
	//
	// example:
	//
	// Parsing completed and available for content plans
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	// The stable error code for summary failures.
	//
	// example:
	//
	// BRAND_KNOWLEDGE_SUMMARY_TIMEOUT
	SummaryErrorCode *string `json:"SummaryErrorCode,omitempty" xml:"SummaryErrorCode,omitempty"`
	// The summary status.
	//
	// example:
	//
	// SUCCEEDED
	SummaryStatus *string `json:"SummaryStatus,omitempty" xml:"SummaryStatus,omitempty"`
	// The last update time in Unix milliseconds.
	//
	// example:
	//
	// 1735660900000
	UpdatedAt *int64 `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s BrandKnowledgeDocument) String() string {
	return dara.Prettify(s)
}

func (s BrandKnowledgeDocument) GoString() string {
	return s.String()
}

func (s *BrandKnowledgeDocument) GetContentSufficient() *bool {
	return s.ContentSufficient
}

func (s *BrandKnowledgeDocument) GetContentType() *string {
	return s.ContentType
}

func (s *BrandKnowledgeDocument) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *BrandKnowledgeDocument) GetDeletable() *bool {
	return s.Deletable
}

func (s *BrandKnowledgeDocument) GetDisplayTitle() *string {
	return s.DisplayTitle
}

func (s *BrandKnowledgeDocument) GetDocCategory() *string {
	return s.DocCategory
}

func (s *BrandKnowledgeDocument) GetDocumentId() *string {
	return s.DocumentId
}

func (s *BrandKnowledgeDocument) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *BrandKnowledgeDocument) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *BrandKnowledgeDocument) GetFileSize() *int64 {
	return s.FileSize
}

func (s *BrandKnowledgeDocument) GetKeyTopics() []*string {
	return s.KeyTopics
}

func (s *BrandKnowledgeDocument) GetOneLineSummary() *string {
	return s.OneLineSummary
}

func (s *BrandKnowledgeDocument) GetOriginalName() *string {
	return s.OriginalName
}

func (s *BrandKnowledgeDocument) GetProjectId() *string {
	return s.ProjectId
}

func (s *BrandKnowledgeDocument) GetReadyAt() *int64 {
	return s.ReadyAt
}

func (s *BrandKnowledgeDocument) GetStatus() *string {
	return s.Status
}

func (s *BrandKnowledgeDocument) GetStatusMessage() *string {
	return s.StatusMessage
}

func (s *BrandKnowledgeDocument) GetSummaryErrorCode() *string {
	return s.SummaryErrorCode
}

func (s *BrandKnowledgeDocument) GetSummaryStatus() *string {
	return s.SummaryStatus
}

func (s *BrandKnowledgeDocument) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *BrandKnowledgeDocument) SetContentSufficient(v bool) *BrandKnowledgeDocument {
	s.ContentSufficient = &v
	return s
}

func (s *BrandKnowledgeDocument) SetContentType(v string) *BrandKnowledgeDocument {
	s.ContentType = &v
	return s
}

func (s *BrandKnowledgeDocument) SetCreatedAt(v int64) *BrandKnowledgeDocument {
	s.CreatedAt = &v
	return s
}

func (s *BrandKnowledgeDocument) SetDeletable(v bool) *BrandKnowledgeDocument {
	s.Deletable = &v
	return s
}

func (s *BrandKnowledgeDocument) SetDisplayTitle(v string) *BrandKnowledgeDocument {
	s.DisplayTitle = &v
	return s
}

func (s *BrandKnowledgeDocument) SetDocCategory(v string) *BrandKnowledgeDocument {
	s.DocCategory = &v
	return s
}

func (s *BrandKnowledgeDocument) SetDocumentId(v string) *BrandKnowledgeDocument {
	s.DocumentId = &v
	return s
}

func (s *BrandKnowledgeDocument) SetErrorCode(v string) *BrandKnowledgeDocument {
	s.ErrorCode = &v
	return s
}

func (s *BrandKnowledgeDocument) SetErrorMessage(v string) *BrandKnowledgeDocument {
	s.ErrorMessage = &v
	return s
}

func (s *BrandKnowledgeDocument) SetFileSize(v int64) *BrandKnowledgeDocument {
	s.FileSize = &v
	return s
}

func (s *BrandKnowledgeDocument) SetKeyTopics(v []*string) *BrandKnowledgeDocument {
	s.KeyTopics = v
	return s
}

func (s *BrandKnowledgeDocument) SetOneLineSummary(v string) *BrandKnowledgeDocument {
	s.OneLineSummary = &v
	return s
}

func (s *BrandKnowledgeDocument) SetOriginalName(v string) *BrandKnowledgeDocument {
	s.OriginalName = &v
	return s
}

func (s *BrandKnowledgeDocument) SetProjectId(v string) *BrandKnowledgeDocument {
	s.ProjectId = &v
	return s
}

func (s *BrandKnowledgeDocument) SetReadyAt(v int64) *BrandKnowledgeDocument {
	s.ReadyAt = &v
	return s
}

func (s *BrandKnowledgeDocument) SetStatus(v string) *BrandKnowledgeDocument {
	s.Status = &v
	return s
}

func (s *BrandKnowledgeDocument) SetStatusMessage(v string) *BrandKnowledgeDocument {
	s.StatusMessage = &v
	return s
}

func (s *BrandKnowledgeDocument) SetSummaryErrorCode(v string) *BrandKnowledgeDocument {
	s.SummaryErrorCode = &v
	return s
}

func (s *BrandKnowledgeDocument) SetSummaryStatus(v string) *BrandKnowledgeDocument {
	s.SummaryStatus = &v
	return s
}

func (s *BrandKnowledgeDocument) SetUpdatedAt(v int64) *BrandKnowledgeDocument {
	s.UpdatedAt = &v
	return s
}

func (s *BrandKnowledgeDocument) Validate() error {
	return dara.Validate(s)
}
